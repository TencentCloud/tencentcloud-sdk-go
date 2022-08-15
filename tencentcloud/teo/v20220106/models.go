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

type ACLCondition struct {
	// 匹配字段
	MatchFrom *string `json:"MatchFrom,omitempty" name:"MatchFrom"`

	// 匹配字符串
	MatchParam *string `json:"MatchParam,omitempty" name:"MatchParam"`

	// 匹配关系
	Operator *string `json:"Operator,omitempty" name:"Operator"`

	// 匹配内容
	MatchContent *string `json:"MatchContent,omitempty" name:"MatchContent"`
}

type ACLUserRule struct {
	// 规则名
	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`

	// 动作
	Action *string `json:"Action,omitempty" name:"Action"`

	// 状态
	RuleStatus *string `json:"RuleStatus,omitempty" name:"RuleStatus"`

	// ACL规则
	Conditions []*ACLCondition `json:"Conditions,omitempty" name:"Conditions"`

	// 规则优先级
	RulePriority *int64 `json:"RulePriority,omitempty" name:"RulePriority"`

	// 规则id
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleID *int64 `json:"RuleID,omitempty" name:"RuleID"`

	// 更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// ip封禁的惩罚时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	PunishTime *int64 `json:"PunishTime,omitempty" name:"PunishTime"`

	// ip封禁的惩罚时间单位
	// 注意：此字段可能返回 null，表示取不到有效值。
	PunishTimeUnit *string `json:"PunishTimeUnit,omitempty" name:"PunishTimeUnit"`

	// 自定义返回页面的名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 自定义返回页面的实例id
	// 注意：此字段可能返回 null，表示取不到有效值。
	PageId *int64 `json:"PageId,omitempty" name:"PageId"`

	// 重定向时候的地址，必须为本用户接入的站点子域名
	// 注意：此字段可能返回 null，表示取不到有效值。
	RedirectUrl *string `json:"RedirectUrl,omitempty" name:"RedirectUrl"`

	// 重定向时候的返回码
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResponseCode *int64 `json:"ResponseCode,omitempty" name:"ResponseCode"`
}

type AclConfig struct {
	// 开关
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// ACL用户规则
	UserRules []*ACLUserRule `json:"UserRules,omitempty" name:"UserRules"`
}

type AiRule struct {
	// smart_status_close-关闭；smart_status_open-拦截处置；
	// smart_status_observe-观察处置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Mode *string `json:"Mode,omitempty" name:"Mode"`
}

type ApplicationProxy struct {
	// 代理ID。
	ProxyId *string `json:"ProxyId,omitempty" name:"ProxyId"`

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

	// 字段已经废弃。
	ForwardClientIp *string `json:"ForwardClientIp,omitempty" name:"ForwardClientIp"`

	// 字段已经废弃。
	SessionPersist *bool `json:"SessionPersist,omitempty" name:"SessionPersist"`

	// 规则列表。
	Rule []*ApplicationProxyRule `json:"Rule,omitempty" name:"Rule"`

	// 状态，取值有：
	// <li>online：启用；</li>
	// <li>offline：停用；</li>
	// <li>progress：部署中；</li>
	// <li>stopping：停用中；</li>
	// <li>fail：部署失败/停用失败。</li>
	Status *string `json:"Status,omitempty" name:"Status"`

	// 调度信息。
	ScheduleValue []*string `json:"ScheduleValue,omitempty" name:"ScheduleValue"`

	// 更新时间。
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 站点ID。
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 站点名称。
	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`

	// 会话保持时间。
	SessionPersistTime *uint64 `json:"SessionPersistTime,omitempty" name:"SessionPersistTime"`

	// 四层代理模式，取值有：
	// <li>hostname：表示子域名模式；</li>
	// <li>instance：表示实例模式。</li>
	ProxyType *string `json:"ProxyType,omitempty" name:"ProxyType"`

	// 当ProxyType=hostname时：
	// 表示代理加速唯一标识。
	HostId *string `json:"HostId,omitempty" name:"HostId"`

	// Ipv6访问配置。
	Ipv6 *Ipv6Access `json:"Ipv6,omitempty" name:"Ipv6"`
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
	// 注意：此字段可能返回 null，表示取不到有效值。
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
	// <li>OFF：不传递。</li>
	ForwardClientIp *string `json:"ForwardClientIp,omitempty" name:"ForwardClientIp"`

	// 是否开启会话保持，取值有：
	// <li>true：开启；</li>
	// <li>false：关闭。</li>
	SessionPersist *bool `json:"SessionPersist,omitempty" name:"SessionPersist"`
}

type BotConfig struct {
	// bot开关
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// 预置规则
	ManagedRule *BotManagedRule `json:"ManagedRule,omitempty" name:"ManagedRule"`

	// 保留
	UaBotRule *BotManagedRule `json:"UaBotRule,omitempty" name:"UaBotRule"`

	// 保留
	IspBotRule *BotManagedRule `json:"IspBotRule,omitempty" name:"IspBotRule"`

	// 用户画像规则
	PortraitRule *BotPortraitRule `json:"PortraitRule,omitempty" name:"PortraitRule"`

	// Bot智能分析
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

	// 当前该字段无效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AttackType *string `json:"AttackType,omitempty" name:"AttackType"`

	// 请求方法。
	RequestMethod *string `json:"RequestMethod,omitempty" name:"RequestMethod"`

	// 攻击内容。
	AttackContent *string `json:"AttackContent,omitempty" name:"AttackContent"`

	// 当前该字段无效 。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RiskLevel *string `json:"RiskLevel,omitempty" name:"RiskLevel"`

	// 当前该字段无效 。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleId *uint64 `json:"RuleId,omitempty" name:"RuleId"`

	// IP所在国家iso-3166中alpha-2编码，编码信息请参考[ISO-3166](https://git.woa.com/edgeone/iso-3166/blob/master/all/all.json)。
	SipCountryCode *string `json:"SipCountryCode,omitempty" name:"SipCountryCode"`

	// 请求（事件）ID。
	EventId *string `json:"EventId,omitempty" name:"EventId"`

	// 该字段当前无效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DisposalMethod *string `json:"DisposalMethod,omitempty" name:"DisposalMethod"`

	// 该字段当前无效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	HttpLog *string `json:"HttpLog,omitempty" name:"HttpLog"`

	// user agent。
	Ua *string `json:"Ua,omitempty" name:"Ua"`

	// 该字段当前无效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DetectionMethod *string `json:"DetectionMethod,omitempty" name:"DetectionMethod"`

	// 该字段当前无效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Confidence *string `json:"Confidence,omitempty" name:"Confidence"`

	// 该字段当前无效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Maliciousness *string `json:"Maliciousness,omitempty" name:"Maliciousness"`

	// 规则相关信息列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleDetailList []*SecRuleRelatedInfo `json:"RuleDetailList,omitempty" name:"RuleDetailList"`

	// Bot标签。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Label *string `json:"Label,omitempty" name:"Label"`
}

type BotLogData struct {
	// Bot攻击日志数据集合。
	// 注意：此字段可能返回 null，表示取不到有效值。
	List []*BotLog `json:"List,omitempty" name:"List"`

	// 分页拉取的起始页号。最小值：1。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PageNo *int64 `json:"PageNo,omitempty" name:"PageNo"`

	// 分页拉取的最大返回结果数。最大值：1000。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`

	// 总页数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Pages *int64 `json:"Pages,omitempty" name:"Pages"`

	// 总条数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalSize *int64 `json:"TotalSize,omitempty" name:"TotalSize"`
}

type BotManagedRule struct {
	// 想开启的规则id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ManagedIds []*int64 `json:"ManagedIds,omitempty" name:"ManagedIds"`

	// 本规则的id
	RuleID *int64 `json:"RuleID,omitempty" name:"RuleID"`

	// drop/trans/monitor/alg
	Action *string `json:"Action,omitempty" name:"Action"`

	// ip封禁的惩罚时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	PunishTime *int64 `json:"PunishTime,omitempty" name:"PunishTime"`

	// 单位
	// 注意：此字段可能返回 null，表示取不到有效值。
	PunishTimeUnit *string `json:"PunishTimeUnit,omitempty" name:"PunishTimeUnit"`

	// 自定义返回页面的名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 自定义返回页面的实例id
	// 注意：此字段可能返回 null，表示取不到有效值。
	PageId *int64 `json:"PageId,omitempty" name:"PageId"`

	// 重定向时候的地址，必须为本用户接入的站点子域名，使用URLENCODE
	// 注意：此字段可能返回 null，表示取不到有效值。
	RedirectUrl *string `json:"RedirectUrl,omitempty" name:"RedirectUrl"`

	// 重定向时候的返回码
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResponseCode *int64 `json:"ResponseCode,omitempty" name:"ResponseCode"`

	// 放行的规则ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TransManagedIds []*int64 `json:"TransManagedIds,omitempty" name:"TransManagedIds"`

	// JS挑战的规则ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	AlgManagedIds []*int64 `json:"AlgManagedIds,omitempty" name:"AlgManagedIds"`

	// 数字验证码的规则ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	CapManagedIds []*int64 `json:"CapManagedIds,omitempty" name:"CapManagedIds"`

	// 观察的规则ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	MonManagedIds []*int64 `json:"MonManagedIds,omitempty" name:"MonManagedIds"`

	// 拦截的规则ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	DropManagedIds []*int64 `json:"DropManagedIds,omitempty" name:"DropManagedIds"`
}

type BotManagedRuleDetail struct {
	// 规则ID
	RuleId *int64 `json:"RuleId,omitempty" name:"RuleId"`

	// 规则描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// 规则分类
	RuleTypeName *string `json:"RuleTypeName,omitempty" name:"RuleTypeName"`

	// 该规则开启/关闭
	Status *string `json:"Status,omitempty" name:"Status"`
}

type BotPortraitRule struct {
	// 本规则的id
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleID *int64 `json:"RuleID,omitempty" name:"RuleID"`

	// JS挑战的规则ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	AlgManagedIds []*int64 `json:"AlgManagedIds,omitempty" name:"AlgManagedIds"`

	// 数字验证码的规则ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	CapManagedIds []*int64 `json:"CapManagedIds,omitempty" name:"CapManagedIds"`

	// 观察的规则ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	MonManagedIds []*int64 `json:"MonManagedIds,omitempty" name:"MonManagedIds"`

	// 拦截的规则ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	DropManagedIds []*int64 `json:"DropManagedIds,omitempty" name:"DropManagedIds"`

	// 本功能的开关
	// 注意：此字段可能返回 null，表示取不到有效值。
	Switch *string `json:"Switch,omitempty" name:"Switch"`
}

type CCInterceptEvent struct {
	// 客户端ip
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClientIp *string `json:"ClientIp,omitempty" name:"ClientIp"`

	// 拦截次数/min
	// 注意：此字段可能返回 null，表示取不到有效值。
	InterceptNum *int64 `json:"InterceptNum,omitempty" name:"InterceptNum"`

	// 速拦截时间，分钟时间/min,单位为s
	InterceptTime *int64 `json:"InterceptTime,omitempty" name:"InterceptTime"`
}

type CCInterceptEventData struct {
	// 攻击事件数据集合
	// 注意：此字段可能返回 null，表示取不到有效值。
	List []*CCInterceptEvent `json:"List,omitempty" name:"List"`

	// 当前页
	// 注意：此字段可能返回 null，表示取不到有效值。
	PageNo *int64 `json:"PageNo,omitempty" name:"PageNo"`

	// 每页展示条数
	// 注意：此字段可能返回 null，表示取不到有效值。
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`

	// 总页数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Pages *int64 `json:"Pages,omitempty" name:"Pages"`

	// 总条数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalSize *int64 `json:"TotalSize,omitempty" name:"TotalSize"`
}

type CCLog struct {
	// 攻击请求时间，采用unix秒级时间戳。
	AttackTime *uint64 `json:"AttackTime,omitempty" name:"AttackTime"`

	// 客户端ip。
	AttackSip *string `json:"AttackSip,omitempty" name:"AttackSip"`

	// 受攻击域名。
	AttackDomain *string `json:"AttackDomain,omitempty" name:"AttackDomain"`

	// URI。
	RequestUri *string `json:"RequestUri,omitempty" name:"RequestUri"`

	// 命中次数。
	HitCount *uint64 `json:"HitCount,omitempty" name:"HitCount"`

	// IP所在国家iso-3166中alpha-2编码，编码信息请参考[ISO-3166](https://git.woa.com/edgeone/iso-3166/blob/master/all/all.json)。
	SipCountryCode *string `json:"SipCountryCode,omitempty" name:"SipCountryCode"`

	// 请求（事件）ID。
	EventId *string `json:"EventId,omitempty" name:"EventId"`

	// 当前该字段已废弃。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DisposalMethod *string `json:"DisposalMethod,omitempty" name:"DisposalMethod"`

	// 当前该字段已废弃。
	// 注意：此字段可能返回 null，表示取不到有效值。
	HttpLog *string `json:"HttpLog,omitempty" name:"HttpLog"`

	// 当前该字段已废弃。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleId *uint64 `json:"RuleId,omitempty" name:"RuleId"`

	// 当前该字段已废弃。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RiskLevel *string `json:"RiskLevel,omitempty" name:"RiskLevel"`

	// User Agent，仅自定义规则日志中存在。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Ua *string `json:"Ua,omitempty" name:"Ua"`

	// 请求方法，仅自定义规则日志中存在。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RequestMethod *string `json:"RequestMethod,omitempty" name:"RequestMethod"`

	// 规则信息列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleDetailList []*SecRuleRelatedInfo `json:"RuleDetailList,omitempty" name:"RuleDetailList"`
}

type CCLogData struct {
	// CC拦截日志数据集合。
	// 注意：此字段可能返回 null，表示取不到有效值。
	List []*CCLog `json:"List,omitempty" name:"List"`

	// 分页拉取的起始页号。最小值：1。
	PageNo *int64 `json:"PageNo,omitempty" name:"PageNo"`

	// 分页拉取的最大返回结果数。最大值：1000。
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`

	// 总页数。
	Pages *int64 `json:"Pages,omitempty" name:"Pages"`

	// 总条数。
	TotalSize *int64 `json:"TotalSize,omitempty" name:"TotalSize"`
}

type CacheConfig struct {
	// 缓存配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Cache *CacheConfigCache `json:"Cache,omitempty" name:"Cache"`

	// 不缓存配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	NoCache *CacheConfigNoCache `json:"NoCache,omitempty" name:"NoCache"`

	// 遵循源站配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	FollowOrigin *CacheConfigFollowOrigin `json:"FollowOrigin,omitempty" name:"FollowOrigin"`
}

type CacheConfigCache struct {
	// 缓存配置开关，取值有：
	// <li>on：开启；</li>
	// <li>off：关闭。</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
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

type CacheConfigFollowOrigin struct {
	// 遵循源站配置开关，取值有：
	// <li>on：开启；</li>
	// <li>off：关闭。</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Switch *string `json:"Switch,omitempty" name:"Switch"`
}

type CacheConfigNoCache struct {
	// 不缓存配置开关，取值有：
	// <li>on：开启；</li>
	// <li>off：关闭。</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Switch *string `json:"Switch,omitempty" name:"Switch"`
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

// Predefined struct for user
type CheckCertificateRequestParams struct {
	// 证书
	Certificate *string `json:"Certificate,omitempty" name:"Certificate"`

	// 私钥
	PrivateKey *string `json:"PrivateKey,omitempty" name:"PrivateKey"`
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

type ClientIp struct {
	// 配置开关，取值有：
	// <li>on：开启；</li>
	// <li>off：关闭。</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// 回源时，存放客户端IP的请求头名称。
	// 为空则使用默认值：X-Forwarded-IP。
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

	// 站点名称。
	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`

	// 规则详细信息。
	Rule []*ApplicationProxyRule `json:"Rule,omitempty" name:"Rule"`

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

	// 字段已经废弃。
	SessionPersist *bool `json:"SessionPersist,omitempty" name:"SessionPersist"`

	// 字段已经废弃。
	ForwardClientIp *string `json:"ForwardClientIp,omitempty" name:"ForwardClientIp"`

	// 四层代理模式，取值有：
	// <li>hostname：表示子域名模式；</li>
	// <li>instance：表示实例模式。</li>不填写使用默认值instance。
	ProxyType *string `json:"ProxyType,omitempty" name:"ProxyType"`

	// 会话保持时间，取值范围：30-3600，单位：秒。
	// 不填写使用默认值600。
	SessionPersistTime *uint64 `json:"SessionPersistTime,omitempty" name:"SessionPersistTime"`

	// Ipv6访问配置。
	// 不填写表示关闭Ipv6访问。
	Ipv6 *Ipv6Access `json:"Ipv6,omitempty" name:"Ipv6"`
}

type CreateApplicationProxyRequest struct {
	*tchttp.BaseRequest
	
	// 站点ID。
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 站点名称。
	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`

	// 规则详细信息。
	Rule []*ApplicationProxyRule `json:"Rule,omitempty" name:"Rule"`

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

	// 字段已经废弃。
	SessionPersist *bool `json:"SessionPersist,omitempty" name:"SessionPersist"`

	// 字段已经废弃。
	ForwardClientIp *string `json:"ForwardClientIp,omitempty" name:"ForwardClientIp"`

	// 四层代理模式，取值有：
	// <li>hostname：表示子域名模式；</li>
	// <li>instance：表示实例模式。</li>不填写使用默认值instance。
	ProxyType *string `json:"ProxyType,omitempty" name:"ProxyType"`

	// 会话保持时间，取值范围：30-3600，单位：秒。
	// 不填写使用默认值600。
	SessionPersistTime *uint64 `json:"SessionPersistTime,omitempty" name:"SessionPersistTime"`

	// Ipv6访问配置。
	// 不填写表示关闭Ipv6访问。
	Ipv6 *Ipv6Access `json:"Ipv6,omitempty" name:"Ipv6"`
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
	delete(f, "Rule")
	delete(f, "ProxyName")
	delete(f, "PlatType")
	delete(f, "SecurityType")
	delete(f, "AccelerateType")
	delete(f, "SessionPersist")
	delete(f, "ForwardClientIp")
	delete(f, "ProxyType")
	delete(f, "SessionPersistTime")
	delete(f, "Ipv6")
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
	OriginType *string `json:"OriginType,omitempty" name:"OriginType"`

	// 源站信息：
	// 当OriginType=custom时，表示多个：
	// IP:端口
	// 域名:端口
	// 当OriginType=origins时，包含一个元素，表示源站组ID
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
	OriginType *string `json:"OriginType,omitempty" name:"OriginType"`

	// 源站信息：
	// 当OriginType=custom时，表示多个：
	// IP:端口
	// 域名:端口
	// 当OriginType=origins时，包含一个元素，表示源站组ID
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
type CreateApplicationProxyRulesRequestParams struct {
	// 站点ID
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 代理ID
	ProxyId *string `json:"ProxyId,omitempty" name:"ProxyId"`

	// 规则列表
	Rule []*ApplicationProxyRule `json:"Rule,omitempty" name:"Rule"`
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

// Predefined struct for user
type CreateApplicationProxyRulesResponseParams struct {
	// 新增的规则ID数组
	RuleId []*string `json:"RuleId,omitempty" name:"RuleId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateApplicationProxyRulesResponse struct {
	*tchttp.BaseResponse
	Response *CreateApplicationProxyRulesResponseParams `json:"Response"`
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

// Predefined struct for user
type CreateCustomErrorPageRequestParams struct {
	// zone的id
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 具体所属实体
	Entity *string `json:"Entity,omitempty" name:"Entity"`

	// 自定义页面的文件名
	Name *string `json:"Name,omitempty" name:"Name"`

	// 自定义页面的内容
	Content *string `json:"Content,omitempty" name:"Content"`
}

type CreateCustomErrorPageRequest struct {
	*tchttp.BaseRequest
	
	// zone的id
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 具体所属实体
	Entity *string `json:"Entity,omitempty" name:"Entity"`

	// 自定义页面的文件名
	Name *string `json:"Name,omitempty" name:"Name"`

	// 自定义页面的内容
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
	// 自定义页面上传后的唯一id
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

// Predefined struct for user
type CreateDnsRecordResponseParams struct {
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
type CreateLoadBalancingRequestParams struct {
	// 站点ID
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 子域名
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

type CreateLoadBalancingRequest struct {
	*tchttp.BaseRequest
	
	// 站点ID
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 子域名
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

// Predefined struct for user
type CreateLoadBalancingResponseParams struct {
	// 负载均衡ID
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
type CreateOriginGroupRequestParams struct {
	// 源站组名称
	OriginName *string `json:"OriginName,omitempty" name:"OriginName"`

	// 配置类型，当OriginType=self 时，需要填写：
	// area: 按区域配置
	// weight: 按权重配置
	// 当OriginType=third_party/cos 时，不需要填写
	Type *string `json:"Type,omitempty" name:"Type"`

	// 源站记录
	Record []*OriginRecord `json:"Record,omitempty" name:"Record"`

	// 站点ID
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 源站类型
	// self：自有源站
	// third_party：第三方源站
	// cos：腾讯云COS源站
	OriginType *string `json:"OriginType,omitempty" name:"OriginType"`
}

type CreateOriginGroupRequest struct {
	*tchttp.BaseRequest
	
	// 源站组名称
	OriginName *string `json:"OriginName,omitempty" name:"OriginName"`

	// 配置类型，当OriginType=self 时，需要填写：
	// area: 按区域配置
	// weight: 按权重配置
	// 当OriginType=third_party/cos 时，不需要填写
	Type *string `json:"Type,omitempty" name:"Type"`

	// 源站记录
	Record []*OriginRecord `json:"Record,omitempty" name:"Record"`

	// 站点ID
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 源站类型
	// self：自有源站
	// third_party：第三方源站
	// cos：腾讯云COS源站
	OriginType *string `json:"OriginType,omitempty" name:"OriginType"`
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
	delete(f, "OriginName")
	delete(f, "Type")
	delete(f, "Record")
	delete(f, "ZoneId")
	delete(f, "OriginType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateOriginGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateOriginGroupResponseParams struct {
	// 新增的源站组ID
	OriginId *string `json:"OriginId,omitempty" name:"OriginId"`

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
type CreatePrefetchTaskRequestParams struct {
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

// Predefined struct for user
type CreatePrefetchTaskResponseParams struct {
	// 任务ID
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// 失败的任务列表
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

// Predefined struct for user
type CreatePurgeTaskResponseParams struct {
	// 任务ID
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// 失败的任务列表及原因
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
type CreateZoneRequestParams struct {
	// 站点名字
	Name *string `json:"Name,omitempty" name:"Name"`

	// 接入方式，默认full
	// - full NS接入
	// - partial CNAME接入
	Type *string `json:"Type,omitempty" name:"Type"`

	// 是否跳过站点历史解析记录扫描
	JumpStart *bool `json:"JumpStart,omitempty" name:"JumpStart"`

	// 资源标签
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
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

	// 资源标签
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
	delete(f, "Name")
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

type DDoSAcl struct {
	// 目的端口end
	DportEnd *int64 `json:"DportEnd,omitempty" name:"DportEnd"`

	// 目的端口start
	DportStart *int64 `json:"DportStart,omitempty" name:"DportStart"`

	// 源端口end
	SportEnd *int64 `json:"SportEnd,omitempty" name:"SportEnd"`

	// 源端口start
	SportStart *int64 `json:"SportStart,omitempty" name:"SportStart"`

	// 协议 'tcp', 'udp', 'all'
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// 动作  drop-丢弃,；transmit-放行； forward-继续防护
	Action *string `json:"Action,omitempty" name:"Action"`

	// 是否为系统配置 0-人工配置；1-系统配置
	Default *int64 `json:"Default,omitempty" name:"Default"`
}

type DDoSAntiPly struct {
	// tcp协议封禁 on-开；off-关
	DropTcp *string `json:"DropTcp,omitempty" name:"DropTcp"`

	// udp协议封禁 on-开；off-关
	DropUdp *string `json:"DropUdp,omitempty" name:"DropUdp"`

	// icmp协议封禁 on-开；off-关
	DropIcmp *string `json:"DropIcmp,omitempty" name:"DropIcmp"`

	// 其他协议封禁 on-开；off-关
	DropOther *string `json:"DropOther,omitempty" name:"DropOther"`

	// 源每秒新建数限制  0-4294967295
	SourceCreateLimit *int64 `json:"SourceCreateLimit,omitempty" name:"SourceCreateLimit"`

	// 源并发连接控制 0-4294967295
	SourceConnectLimit *int64 `json:"SourceConnectLimit,omitempty" name:"SourceConnectLimit"`

	// 目的每秒新建数限制 0-4294967295
	DestinationCreateLimit *int64 `json:"DestinationCreateLimit,omitempty" name:"DestinationCreateLimit"`

	// 目的端口的并发连接控制 0-4294967295
	DestinationConnectLimit *int64 `json:"DestinationConnectLimit,omitempty" name:"DestinationConnectLimit"`

	// 异常连接数阈值  0-4294967295
	AbnormalConnectNum *int64 `json:"AbnormalConnectNum,omitempty" name:"AbnormalConnectNum"`

	// syn占比异常阈值 0-100
	AbnormalSynRatio *int64 `json:"AbnormalSynRatio,omitempty" name:"AbnormalSynRatio"`

	// syn个数异常阈值 0-65535
	AbnormalSynNum *int64 `json:"AbnormalSynNum,omitempty" name:"AbnormalSynNum"`

	// 连接超时检测 0-65535
	ConnectTimeout *int64 `json:"ConnectTimeout,omitempty" name:"ConnectTimeout"`

	// 空连接防护开启 0-1
	EmptyConnectProtect *string `json:"EmptyConnectProtect,omitempty" name:"EmptyConnectProtect"`

	// UDP分片开关；off-关闭，on-开启
	// 注意：此字段可能返回 null，表示取不到有效值。
	UdpShard *string `json:"UdpShard,omitempty" name:"UdpShard"`
}

type DDoSApplication struct {
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

type DDoSConfig struct {
	// 开关
	Switch *string `json:"Switch,omitempty" name:"Switch"`
}

type DDoSFeaturesFilter struct {
	// 动作 drop-丢弃；transmit-放行；drop_block-丢弃并拉黑；forward-继续防护
	Action *string `json:"Action,omitempty" name:"Action"`

	// 深度值1
	Depth *int64 `json:"Depth,omitempty" name:"Depth"`

	// 深度值2
	Depth2 *int64 `json:"Depth2,omitempty" name:"Depth2"`

	// 目标端口结束
	DportEnd *int64 `json:"DportEnd,omitempty" name:"DportEnd"`

	// 目标端口开始
	DportStart *int64 `json:"DportStart,omitempty" name:"DportStart"`

	// 取非判断1
	IsNot *int64 `json:"IsNot,omitempty" name:"IsNot"`

	// 取非判断2
	IsNot2 *int64 `json:"IsNot2,omitempty" name:"IsNot2"`

	// 多特征关系（单特征时(none)，第二特征相关配置可不填） none；and；or
	MatchLogic *string `json:"MatchLogic,omitempty" name:"MatchLogic"`

	// 匹配方式1 pcre-正则匹配, sunday-字符串匹配
	MatchType *string `json:"MatchType,omitempty" name:"MatchType"`

	// 匹配方式2 pcre-正则匹配, sunday-字符串匹配
	MatchType2 *string `json:"MatchType2,omitempty" name:"MatchType2"`

	// 偏移量1
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 偏移量2
	Offset2 *int64 `json:"Offset2,omitempty" name:"Offset2"`

	// 最大包长
	PacketMax *int64 `json:"PacketMax,omitempty" name:"PacketMax"`

	// 最小包长
	PacketMin *int64 `json:"PacketMin,omitempty" name:"PacketMin"`

	// 协议 tcp；udp；icmp；all
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// 源端口结束
	SportEnd *int64 `json:"SportEnd,omitempty" name:"SportEnd"`

	// 源端口开始
	SportStart *int64 `json:"SportStart,omitempty" name:"SportStart"`

	// 匹配字符串1
	Str *string `json:"Str,omitempty" name:"Str"`

	// 匹配字符串2
	Str2 *string `json:"Str2,omitempty" name:"Str2"`

	// 匹配开始层级，层级参考计算机网络结构 begin_l5, no_match, begin_l3, begin_l4
	MatchBegin *string `json:"MatchBegin,omitempty" name:"MatchBegin"`

	// 匹配开始层级，层级参考计算机网络结构 begin_l5, no_match, begin_l3, begin_l4
	MatchBegin2 *string `json:"MatchBegin2,omitempty" name:"MatchBegin2"`
}

type DDoSGeoIp struct {
	// 地域信息，ID参考接口DescribeSecurityPolicyRegions
	// 注意：此字段可能返回 null，表示取不到有效值。
	RegionId []*int64 `json:"RegionId,omitempty" name:"RegionId"`

	// 区域封禁清空标识
	// 注意：此字段可能返回 null，表示取不到有效值。
	Switch *string `json:"Switch,omitempty" name:"Switch"`
}

type DDoSStatusInfo struct {
	// 不支持，填off
	AiStatus *string `json:"AiStatus,omitempty" name:"AiStatus"`

	// 用户appid
	Appid *string `json:"Appid,omitempty" name:"Appid"`

	// 策略等级 low, middle, high
	PlyLevel *string `json:"PlyLevel,omitempty" name:"PlyLevel"`
}

type DDoSUserAllowBlockIP struct {
	// 用户ip
	// 注意：此字段可能返回 null，表示取不到有效值。
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// 掩码
	// 注意：此字段可能返回 null，表示取不到有效值。
	Mask *int64 `json:"Mask,omitempty" name:"Mask"`

	// 类型 block-丢弃；allow-允许
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitempty" name:"Type"`

	// 时间戳
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *int64 `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 用户ip范围截止
	// 注意：此字段可能返回 null，表示取不到有效值。
	Ip2 *string `json:"Ip2,omitempty" name:"Ip2"`

	// 掩码截止范围
	// 注意：此字段可能返回 null，表示取不到有效值。
	Mask2 *int64 `json:"Mask2,omitempty" name:"Mask2"`
}

type DDosAttackEvent struct {
	// ddos 策略组id
	// 注意：此字段可能返回 null，表示取不到有效值。
	PolicyId *int64 `json:"PolicyId,omitempty" name:"PolicyId"`

	// 攻击类型(对应交互事件名称)
	// 注意：此字段可能返回 null，表示取不到有效值。
	AttackType *string `json:"AttackType,omitempty" name:"AttackType"`

	// 攻击状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	AttackStatus *int64 `json:"AttackStatus,omitempty" name:"AttackStatus"`

	// 攻击最大带宽
	// 注意：此字段可能返回 null，表示取不到有效值。
	AttackMaxBandWidth *int64 `json:"AttackMaxBandWidth,omitempty" name:"AttackMaxBandWidth"`

	// 攻击包速率峰值
	// 注意：此字段可能返回 null，表示取不到有效值。
	AttackPacketMaxRate *int64 `json:"AttackPacketMaxRate,omitempty" name:"AttackPacketMaxRate"`

	// 攻击开始时间 单位为s
	// 注意：此字段可能返回 null，表示取不到有效值。
	AttackStartTime *int64 `json:"AttackStartTime,omitempty" name:"AttackStartTime"`

	// 攻击结束时间 单位为s
	// 注意：此字段可能返回 null，表示取不到有效值。
	AttackEndTime *int64 `json:"AttackEndTime,omitempty" name:"AttackEndTime"`

	// 事件ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	EventId *string `json:"EventId,omitempty" name:"EventId"`

	// 站点id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`
}

type DDosAttackEventData struct {
	// 攻击事件数据集合。
	// 注意：此字段可能返回 null，表示取不到有效值。
	List []*DDosAttackEvent `json:"List,omitempty" name:"List"`

	// 分页拉取的起始页号。最小值：1。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PageNo *int64 `json:"PageNo,omitempty" name:"PageNo"`

	// 分页拉取的最大返回结果数。最大值：1000。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`

	// 总页数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Pages *int64 `json:"Pages,omitempty" name:"Pages"`

	// 总条数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalSize *int64 `json:"TotalSize,omitempty" name:"TotalSize"`
}

type DDosAttackEventDetailData struct {
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

	// ddos 策略组id。
	PolicyId *int64 `json:"PolicyId,omitempty" name:"PolicyId"`
}

type DDosAttackSourceEvent struct {
	// 攻击源ip。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AttackSourceIp *string `json:"AttackSourceIp,omitempty" name:"AttackSourceIp"`

	// 地区（国家）。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AttackRegion *string `json:"AttackRegion,omitempty" name:"AttackRegion"`

	// 累计攻击流量。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AttackFlow *uint64 `json:"AttackFlow,omitempty" name:"AttackFlow"`

	// 累计攻击包量。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AttackPacketNum *uint64 `json:"AttackPacketNum,omitempty" name:"AttackPacketNum"`
}

type DDosAttackSourceEventData struct {
	// DDos攻击源数据集合。
	// 注意：此字段可能返回 null，表示取不到有效值。
	List []*DDosAttackSourceEvent `json:"List,omitempty" name:"List"`

	// 分页拉取的起始页号。最小值：1。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PageNo *int64 `json:"PageNo,omitempty" name:"PageNo"`

	// 分页拉取的最大返回结果数。最大值：1000。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`

	// 总页数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Pages *int64 `json:"Pages,omitempty" name:"Pages"`

	// 总条数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalSize *int64 `json:"TotalSize,omitempty" name:"TotalSize"`
}

type DDosMajorAttackEvent struct {
	// ddos 策略组id。
	PolicyId *int64 `json:"PolicyId,omitempty" name:"PolicyId"`

	// 攻击最大带宽。
	AttackMaxBandWidth *int64 `json:"AttackMaxBandWidth,omitempty" name:"AttackMaxBandWidth"`

	// 攻击请求时间，采用unix秒级时间戳。
	AttackTime *int64 `json:"AttackTime,omitempty" name:"AttackTime"`
}

type DDosMajorAttackEventData struct {
	// DDosMajorAttackEvent ddos 攻击事件。
	// 注意：此字段可能返回 null，表示取不到有效值。
	List []*DDosMajorAttackEvent `json:"List,omitempty" name:"List"`

	// 分页拉取的起始页号。最小值：1。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PageNo *int64 `json:"PageNo,omitempty" name:"PageNo"`

	// 分页拉取的最大返回结果数。最大值：1000。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`

	// 总页数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Pages *int64 `json:"Pages,omitempty" name:"Pages"`

	// 总条数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalSize *int64 `json:"TotalSize,omitempty" name:"TotalSize"`
}

type DataItem struct {
	// 时间
	Time *string `json:"Time,omitempty" name:"Time"`

	// 数值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *uint64 `json:"Value,omitempty" name:"Value"`
}

type DdosAcls struct {
	// 开关 off清空规则标识
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// 端口过了详细参数
	Acl []*DDoSAcl `json:"Acl,omitempty" name:"Acl"`
}

type DdosAllowBlock struct {
	// 开关标识防护是否清空
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// 黑白名单数组
	UserAllowBlockIp []*DDoSUserAllowBlockIP `json:"UserAllowBlockIp,omitempty" name:"UserAllowBlockIp"`
}

type DdosPacketFilter struct {
	// 特征过滤清空标识，off清空处理
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// 特征过滤数组
	PacketFilter []*DDoSFeaturesFilter `json:"PacketFilter,omitempty" name:"PacketFilter"`
}

type DdosRule struct {
	// DDoS防护等级
	// 注意：此字段可能返回 null，表示取不到有效值。
	DdosStatusInfo *DDoSStatusInfo `json:"DdosStatusInfo,omitempty" name:"DdosStatusInfo"`

	// DDoS地域封禁
	// 注意：此字段可能返回 null，表示取不到有效值。
	DdosGeoIp *DDoSGeoIp `json:"DdosGeoIp,omitempty" name:"DdosGeoIp"`

	// DDoS黑白名单
	// 注意：此字段可能返回 null，表示取不到有效值。
	DdosAllowBlock *DdosAllowBlock `json:"DdosAllowBlock,omitempty" name:"DdosAllowBlock"`

	// DDoS 协议封禁+连接防护
	// 注意：此字段可能返回 null，表示取不到有效值。
	DdosAntiPly *DDoSAntiPly `json:"DdosAntiPly,omitempty" name:"DdosAntiPly"`

	// DDoS特征过滤
	// 注意：此字段可能返回 null，表示取不到有效值。
	DdosPacketFilter *DdosPacketFilter `json:"DdosPacketFilter,omitempty" name:"DdosPacketFilter"`

	// DDoS端口过滤
	// 注意：此字段可能返回 null，表示取不到有效值。
	DdosAcl *DdosAcls `json:"DdosAcl,omitempty" name:"DdosAcl"`

	// DDoS开关 on-开启；off-关闭
	// 注意：此字段可能返回 null，表示取不到有效值。
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// UDP分片功能是否支持，off-不支持，on-支持
	// 注意：此字段可能返回 null，表示取不到有效值。
	UdpShardOpen *string `json:"UdpShardOpen,omitempty" name:"UdpShardOpen"`
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

// Predefined struct for user
type DeleteApplicationProxyRequestParams struct {
	// 站点ID
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 代理ID
	ProxyId *string `json:"ProxyId,omitempty" name:"ProxyId"`
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

// Predefined struct for user
type DeleteApplicationProxyResponseParams struct {
	// 代理ID
	ProxyId *string `json:"ProxyId,omitempty" name:"ProxyId"`

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
	// 站点ID
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 代理ID
	ProxyId *string `json:"ProxyId,omitempty" name:"ProxyId"`

	// 规则ID
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`
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

// Predefined struct for user
type DeleteApplicationProxyRuleResponseParams struct {
	// 规则ID
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

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
	// 站点 ID
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 记录 ID
	Ids []*string `json:"Ids,omitempty" name:"Ids"`
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

// Predefined struct for user
type DeleteDnsRecordsResponseParams struct {
	// 记录 ID
	Ids []*string `json:"Ids,omitempty" name:"Ids"`

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
	// 站点ID
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 负载均衡ID
	LoadBalancingId *string `json:"LoadBalancingId,omitempty" name:"LoadBalancingId"`
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

// Predefined struct for user
type DeleteLoadBalancingResponseParams struct {
	// 负载均衡ID
	LoadBalancingId *string `json:"LoadBalancingId,omitempty" name:"LoadBalancingId"`

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
type DeleteOriginGroupRequestParams struct {
	// 源站组ID
	OriginId *string `json:"OriginId,omitempty" name:"OriginId"`

	// 站点ID
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`
}

type DeleteOriginGroupRequest struct {
	*tchttp.BaseRequest
	
	// 源站组ID
	OriginId *string `json:"OriginId,omitempty" name:"OriginId"`

	// 站点ID
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`
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
	delete(f, "OriginId")
	delete(f, "ZoneId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteOriginGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteOriginGroupResponseParams struct {
	// 源站组ID
	OriginId *string `json:"OriginId,omitempty" name:"OriginId"`

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
type DeleteZoneRequestParams struct {
	// 站点 ID
	Id *string `json:"Id,omitempty" name:"Id"`
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

// Predefined struct for user
type DeleteZoneResponseParams struct {
	// 站点 ID
	Id *string `json:"Id,omitempty" name:"Id"`

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
type DescribeApplicationProxyDetailRequestParams struct {
	// 站点ID。
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 实例ID。
	ProxyId *string `json:"ProxyId,omitempty" name:"ProxyId"`
}

type DescribeApplicationProxyDetailRequest struct {
	*tchttp.BaseRequest
	
	// 站点ID。
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 实例ID。
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

// Predefined struct for user
type DescribeApplicationProxyDetailResponseParams struct {
	// 实例ID。
	ProxyId *string `json:"ProxyId,omitempty" name:"ProxyId"`

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

	// 字段已经废弃。
	ForwardClientIp *string `json:"ForwardClientIp,omitempty" name:"ForwardClientIp"`

	// 字段已经废弃。
	SessionPersist *bool `json:"SessionPersist,omitempty" name:"SessionPersist"`

	// 规则列表。
	Rule []*ApplicationProxyRule `json:"Rule,omitempty" name:"Rule"`

	// 状态，取值有：
	// <li>online：启用；</li>
	// <li>offline：停用；</li>
	// <li>progress：部署中。</li>
	Status *string `json:"Status,omitempty" name:"Status"`

	// 调度信息。
	ScheduleValue []*string `json:"ScheduleValue,omitempty" name:"ScheduleValue"`

	// 更新时间。
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 站点ID。
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 站点名称。
	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`

	// 会话保持时间。
	SessionPersistTime *uint64 `json:"SessionPersistTime,omitempty" name:"SessionPersistTime"`

	// 四层代理模式，取值有：
	// <li>hostname：表示子域名模式；</li>
	// <li>instance：表示实例模式。</li>
	ProxyType *string `json:"ProxyType,omitempty" name:"ProxyType"`

	// 当ProxyType=hostname时：
	// 表示代理加速唯一标识。
	HostId *string `json:"HostId,omitempty" name:"HostId"`

	// IPv6访问配置。
	Ipv6 *Ipv6Access `json:"Ipv6,omitempty" name:"Ipv6"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeApplicationProxyDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeApplicationProxyDetailResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeApplicationProxyRequestParams struct {
	// 站点ID。
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 分页查询偏移量，默认为0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页查询限制数目，默认为10，最大可设置为1000。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 代理ID。
	// 当ProxyId为空时，表示查询站点下所有应用代理的列表。
	ProxyId *string `json:"ProxyId,omitempty" name:"ProxyId"`
}

type DescribeApplicationProxyRequest struct {
	*tchttp.BaseRequest
	
	// 站点ID。
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 分页查询偏移量，默认为0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页查询限制数目，默认为10，最大可设置为1000。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 代理ID。
	// 当ProxyId为空时，表示查询站点下所有应用代理的列表。
	ProxyId *string `json:"ProxyId,omitempty" name:"ProxyId"`
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
	delete(f, "ProxyId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeApplicationProxyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApplicationProxyResponseParams struct {
	// 应用代理列表。
	Data []*ApplicationProxy `json:"Data,omitempty" name:"Data"`

	// 记录总数。
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 字段已废弃。
	Quota *int64 `json:"Quota,omitempty" name:"Quota"`

	// 当ProxyId为空时，表示套餐内PlatType为ip的Anycast IP的实例数量。
	IpCount *uint64 `json:"IpCount,omitempty" name:"IpCount"`

	// 当ProxyId为空时，表示套餐内PlatType为domain的CNAME的实例数量。
	DomainCount *uint64 `json:"DomainCount,omitempty" name:"DomainCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeApplicationProxyResponse struct {
	*tchttp.BaseResponse
	Response *DescribeApplicationProxyResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeBotLogRequestParams struct {
	// 起始时间。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 分页拉取的最大返回结果数。最大值：1000。
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`

	// 分页拉取的起始页号。最小值：1。
	PageNo *int64 `json:"PageNo,omitempty" name:"PageNo"`

	// 站点集合，不填默认查询所有站点。
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// 域名集合，不填默认查询所有子域名。
	Domains []*string `json:"Domains,omitempty" name:"Domains"`

	// 筛选条件，取值有：
	// <li>action ：执行动作（处置方式）；</li>
	// <li>sipCountryCode ：ip所在国家 ；</li>
	// <li>attackIp ：攻击ip ；</li>
	// <li>ruleId ：规则id ；</li>
	// <li>eventId ：事件id ；</li>
	// <li>ua ：用户代理 ；</li>
	// <li>requestMethod ：请求方法 ；</li>
	// <li>uri ：统一资源标识符 。</li>
	QueryCondition []*QueryCondition `json:"QueryCondition,omitempty" name:"QueryCondition"`
}

type DescribeBotLogRequest struct {
	*tchttp.BaseRequest
	
	// 起始时间。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 分页拉取的最大返回结果数。最大值：1000。
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`

	// 分页拉取的起始页号。最小值：1。
	PageNo *int64 `json:"PageNo,omitempty" name:"PageNo"`

	// 站点集合，不填默认查询所有站点。
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// 域名集合，不填默认查询所有子域名。
	Domains []*string `json:"Domains,omitempty" name:"Domains"`

	// 筛选条件，取值有：
	// <li>action ：执行动作（处置方式）；</li>
	// <li>sipCountryCode ：ip所在国家 ；</li>
	// <li>attackIp ：攻击ip ；</li>
	// <li>ruleId ：规则id ；</li>
	// <li>eventId ：事件id ；</li>
	// <li>ua ：用户代理 ；</li>
	// <li>requestMethod ：请求方法 ；</li>
	// <li>uri ：统一资源标识符 。</li>
	QueryCondition []*QueryCondition `json:"QueryCondition,omitempty" name:"QueryCondition"`
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
	delete(f, "PageSize")
	delete(f, "PageNo")
	delete(f, "ZoneIds")
	delete(f, "Domains")
	delete(f, "QueryCondition")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBotLogRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBotLogResponseParams struct {
	// Bot攻击数据内容。
	Data *BotLogData `json:"Data,omitempty" name:"Data"`

	// 请求响应状态，取值有：
	// <li>1 ：失败 ；</li>
	// <li>0 ：成功 。</li>
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 请求响应信息。
	Msg *string `json:"Msg,omitempty" name:"Msg"`

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
	// 一级域名
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 子域名/应用名
	Entity *string `json:"Entity,omitempty" name:"Entity"`

	// 页数
	Page *int64 `json:"Page,omitempty" name:"Page"`

	// 每页数量
	PerPage *int64 `json:"PerPage,omitempty" name:"PerPage"`

	// idcid/sipbot/uabot规则类型，空代表拉取全部
	RuleType *string `json:"RuleType,omitempty" name:"RuleType"`
}

type DescribeBotManagedRulesRequest struct {
	*tchttp.BaseRequest
	
	// 一级域名
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 子域名/应用名
	Entity *string `json:"Entity,omitempty" name:"Entity"`

	// 页数
	Page *int64 `json:"Page,omitempty" name:"Page"`

	// 每页数量
	PerPage *int64 `json:"PerPage,omitempty" name:"PerPage"`

	// idcid/sipbot/uabot规则类型，空代表拉取全部
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
	delete(f, "Page")
	delete(f, "PerPage")
	delete(f, "RuleType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBotManagedRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBotManagedRulesResponseParams struct {
	// 本次返回的规则数
	Count *int64 `json:"Count,omitempty" name:"Count"`

	// Bot规则
	Rules []*BotManagedRuleDetail `json:"Rules,omitempty" name:"Rules"`

	// 总规则数
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
type DescribeCnameStatusRequestParams struct {
	// 站点 ID
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 域名列表
	Names []*string `json:"Names,omitempty" name:"Names"`
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

// Predefined struct for user
type DescribeCnameStatusResponseParams struct {
	// 状态列表
	Status []*CnameStatus `json:"Status,omitempty" name:"Status"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeCnameStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCnameStatusResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeDDoSPolicyRequestParams struct {
	// 策略组id
	PolicyId *int64 `json:"PolicyId,omitempty" name:"PolicyId"`

	// 一级域名zone
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`
}

type DescribeDDoSPolicyRequest struct {
	*tchttp.BaseRequest
	
	// 策略组id
	PolicyId *int64 `json:"PolicyId,omitempty" name:"PolicyId"`

	// 一级域名zone
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`
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
	delete(f, "PolicyId")
	delete(f, "ZoneId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDDoSPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDDoSPolicyResponseParams struct {
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
type DescribeDDosAttackDataRequestParams struct {
	// 开始时间。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 统计指标列表，取值有：
	// <li>ddos_attackMaxBandwidth ：攻击带宽峰值 ；</li>
	// <li>ddos_attackMaxPackageRate：攻击包速率峰值  ；</li>
	// <li>ddos_attackBandwidth ：攻击带宽曲线 ；</li>
	// <li>ddos_attackPackageRate ：攻击包速率曲线 。</li>
	MetricNames []*string `json:"MetricNames,omitempty" name:"MetricNames"`

	// 站点id列表，不填默认选择全部站点。
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// ddos策略组id列表，不填默认选择全部策略id。
	PolicyIds []*int64 `json:"PolicyIds,omitempty" name:"PolicyIds"`

	// 端口号。
	Port *int64 `json:"Port,omitempty" name:"Port"`

	// 协议类型，取值有：
	// <li>tcp ；</li>
	// <li>udp ；</li>
	// <li>all 。</li>
	ProtocolType *string `json:"ProtocolType,omitempty" name:"ProtocolType"`

	// 攻击类型，取值有：
	// <li>flood ；</li>
	// <li>icmpFlood ；</li>
	// <li>all 。</li>
	AttackType *string `json:"AttackType,omitempty" name:"AttackType"`

	// 查询时间粒度，取值有：
	// <li>min ：1分钟 ；</li>
	// <li>5min ：5分钟 ；</li>
	// <li>hour ：1小时 ；</li>
	// <li>day ：1天 。</li>
	Interval *string `json:"Interval,omitempty" name:"Interval"`
}

type DescribeDDosAttackDataRequest struct {
	*tchttp.BaseRequest
	
	// 开始时间。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 统计指标列表，取值有：
	// <li>ddos_attackMaxBandwidth ：攻击带宽峰值 ；</li>
	// <li>ddos_attackMaxPackageRate：攻击包速率峰值  ；</li>
	// <li>ddos_attackBandwidth ：攻击带宽曲线 ；</li>
	// <li>ddos_attackPackageRate ：攻击包速率曲线 。</li>
	MetricNames []*string `json:"MetricNames,omitempty" name:"MetricNames"`

	// 站点id列表，不填默认选择全部站点。
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// ddos策略组id列表，不填默认选择全部策略id。
	PolicyIds []*int64 `json:"PolicyIds,omitempty" name:"PolicyIds"`

	// 端口号。
	Port *int64 `json:"Port,omitempty" name:"Port"`

	// 协议类型，取值有：
	// <li>tcp ；</li>
	// <li>udp ；</li>
	// <li>all 。</li>
	ProtocolType *string `json:"ProtocolType,omitempty" name:"ProtocolType"`

	// 攻击类型，取值有：
	// <li>flood ；</li>
	// <li>icmpFlood ；</li>
	// <li>all 。</li>
	AttackType *string `json:"AttackType,omitempty" name:"AttackType"`

	// 查询时间粒度，取值有：
	// <li>min ：1分钟 ；</li>
	// <li>5min ：5分钟 ；</li>
	// <li>hour ：1小时 ；</li>
	// <li>day ：1天 。</li>
	Interval *string `json:"Interval,omitempty" name:"Interval"`
}

func (r *DescribeDDosAttackDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDDosAttackDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "MetricNames")
	delete(f, "ZoneIds")
	delete(f, "PolicyIds")
	delete(f, "Port")
	delete(f, "ProtocolType")
	delete(f, "AttackType")
	delete(f, "Interval")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDDosAttackDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDDosAttackDataResponseParams struct {
	// DDos攻击数据内容。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*SecEntry `json:"Data,omitempty" name:"Data"`

	// 请求响应状态，取值有：
	// <li>1 ：失败 ；</li>
	// <li>0 ：成功 。</li>
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 请求响应信息。
	Msg *string `json:"Msg,omitempty" name:"Msg"`

	// 查询时间粒度，取值有：
	// <li>min ：1分钟 ；</li>
	// <li>5min ：5分钟 ；</li>
	// <li>hour ：1小时 ；</li>
	// <li>day ：1天 。</li>
	Interval *string `json:"Interval,omitempty" name:"Interval"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDDosAttackDataResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDDosAttackDataResponseParams `json:"Response"`
}

func (r *DescribeDDosAttackDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDDosAttackDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDDosAttackEventDetailRequestParams struct {
	// 事件id。
	EventId *string `json:"EventId,omitempty" name:"EventId"`
}

type DescribeDDosAttackEventDetailRequest struct {
	*tchttp.BaseRequest
	
	// 事件id。
	EventId *string `json:"EventId,omitempty" name:"EventId"`
}

func (r *DescribeDDosAttackEventDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDDosAttackEventDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EventId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDDosAttackEventDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDDosAttackEventDetailResponseParams struct {
	// DDos攻击事件详情。
	Data *DDosAttackEventDetailData `json:"Data,omitempty" name:"Data"`

	// 请求响应状态，取值有：
	// <li>1 ：失败 ；</li>
	// <li>0 ：成功 。</li>
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 请求响应信息。
	Msg *string `json:"Msg,omitempty" name:"Msg"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDDosAttackEventDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDDosAttackEventDetailResponseParams `json:"Response"`
}

func (r *DescribeDDosAttackEventDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDDosAttackEventDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDDosAttackEventRequestParams struct {
	// 开始时间。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 分页拉取的最大返回结果数。最大值：1000。
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`

	// 分页拉取的起始页号。最小值：1。
	PageNo *int64 `json:"PageNo,omitempty" name:"PageNo"`

	// ddos策略组id列表，不填默认选择全部策略Id。
	PolicyIds []*int64 `json:"PolicyIds,omitempty" name:"PolicyIds"`

	// 站点id列表，不填默认选择全部站点。
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// 协议类型，取值有：
	// <li>tcp ；</li>
	// <li>udp ；</li>
	// <li>all 。</li>
	ProtocolType *string `json:"ProtocolType,omitempty" name:"ProtocolType"`

	// 是否展示详情，取值有：
	// <li>Y ：展示 ；</li>
	// <li>N ：不展示 。</li>默认为Y。
	IsShowDetail *string `json:"IsShowDetail,omitempty" name:"IsShowDetail"`
}

type DescribeDDosAttackEventRequest struct {
	*tchttp.BaseRequest
	
	// 开始时间。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 分页拉取的最大返回结果数。最大值：1000。
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`

	// 分页拉取的起始页号。最小值：1。
	PageNo *int64 `json:"PageNo,omitempty" name:"PageNo"`

	// ddos策略组id列表，不填默认选择全部策略Id。
	PolicyIds []*int64 `json:"PolicyIds,omitempty" name:"PolicyIds"`

	// 站点id列表，不填默认选择全部站点。
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// 协议类型，取值有：
	// <li>tcp ；</li>
	// <li>udp ；</li>
	// <li>all 。</li>
	ProtocolType *string `json:"ProtocolType,omitempty" name:"ProtocolType"`

	// 是否展示详情，取值有：
	// <li>Y ：展示 ；</li>
	// <li>N ：不展示 。</li>默认为Y。
	IsShowDetail *string `json:"IsShowDetail,omitempty" name:"IsShowDetail"`
}

func (r *DescribeDDosAttackEventRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDDosAttackEventRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "PageSize")
	delete(f, "PageNo")
	delete(f, "PolicyIds")
	delete(f, "ZoneIds")
	delete(f, "ProtocolType")
	delete(f, "IsShowDetail")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDDosAttackEventRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDDosAttackEventResponseParams struct {
	// DDos攻击事件数据。
	Data *DDosAttackEventData `json:"Data,omitempty" name:"Data"`

	// 请求响应状态，取值有：
	// <li>1 ：失败 ；</li>
	// <li>0 ：成功 。</li>
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 请求响应信息。
	Msg *string `json:"Msg,omitempty" name:"Msg"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDDosAttackEventResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDDosAttackEventResponseParams `json:"Response"`
}

func (r *DescribeDDosAttackEventResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDDosAttackEventResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDDosAttackSourceEventRequestParams struct {
	// 开始时间。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 分页拉取的最大返回结果数。最大值：1000。
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`

	// 分页拉取的起始页号。最小值：1。
	PageNo *int64 `json:"PageNo,omitempty" name:"PageNo"`

	// ddos策略组id 集合，不填默认选择全部策略id。
	PolicyIds []*int64 `json:"PolicyIds,omitempty" name:"PolicyIds"`

	// 站点集合，不填默认选择全部站点。
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// 协议类型，取值有：
	// <li>tcp ；</li>
	// <li>udp ；</li>
	// <li>all 。</li>
	ProtocolType *string `json:"ProtocolType,omitempty" name:"ProtocolType"`
}

type DescribeDDosAttackSourceEventRequest struct {
	*tchttp.BaseRequest
	
	// 开始时间。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 分页拉取的最大返回结果数。最大值：1000。
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`

	// 分页拉取的起始页号。最小值：1。
	PageNo *int64 `json:"PageNo,omitempty" name:"PageNo"`

	// ddos策略组id 集合，不填默认选择全部策略id。
	PolicyIds []*int64 `json:"PolicyIds,omitempty" name:"PolicyIds"`

	// 站点集合，不填默认选择全部站点。
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// 协议类型，取值有：
	// <li>tcp ；</li>
	// <li>udp ；</li>
	// <li>all 。</li>
	ProtocolType *string `json:"ProtocolType,omitempty" name:"ProtocolType"`
}

func (r *DescribeDDosAttackSourceEventRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDDosAttackSourceEventRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "PageSize")
	delete(f, "PageNo")
	delete(f, "PolicyIds")
	delete(f, "ZoneIds")
	delete(f, "ProtocolType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDDosAttackSourceEventRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDDosAttackSourceEventResponseParams struct {
	// DDos攻击源数据。
	Data *DDosAttackSourceEventData `json:"Data,omitempty" name:"Data"`

	// 请求响应状态，取值有：
	// <li>1 ：失败 ；</li>
	// <li>0 ：成功 。</li>
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 请求响应信息。
	Msg *string `json:"Msg,omitempty" name:"Msg"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDDosAttackSourceEventResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDDosAttackSourceEventResponseParams `json:"Response"`
}

func (r *DescribeDDosAttackSourceEventResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDDosAttackSourceEventResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDDosAttackTopDataRequestParams struct {
	// 开始时间。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 统计指标列表，取值有：
	// <li>ddos_attackFlux_protocol ：攻击总流量协议类型分布排行 ；</li>
	// <li>ddos_attackPackageNum_protocol ：攻击总包量协议类型分布排行 ；</li>
	// <li>ddos_attackNum_attackType ：攻击总次数攻击类型分布排行 ；</li>
	// <li>ddos_attackNum_sregion ：攻击总次数攻击源地区分布排行 ；</li>
	// <li>ddos_attackFlux_sip ：攻击总流量攻击源ip分布排行 ；</li>
	// <li>ddos_attackFlux_sregion ：攻击总流量攻击源地区分布排行 。</li>
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`

	// 查询前多少个，传值为0返回全量。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 站点id集合，不填默认选择全部站点。
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// ddos策略组id 集合，不填默认选择全部策略id。
	PolicyIds []*int64 `json:"PolicyIds,omitempty" name:"PolicyIds"`

	// 端口号。
	Port *int64 `json:"Port,omitempty" name:"Port"`

	// 协议类型，取值有：
	// <li>tcp ；</li>
	// <li>udp ；</li>
	// <li>all 。</li>
	ProtocolType *string `json:"ProtocolType,omitempty" name:"ProtocolType"`

	// 攻击类型，取值有：
	// <li>flood ；</li>
	// <li>icmpFlood ；</li>
	// <li>all 。</li>
	AttackType *string `json:"AttackType,omitempty" name:"AttackType"`
}

type DescribeDDosAttackTopDataRequest struct {
	*tchttp.BaseRequest
	
	// 开始时间。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 统计指标列表，取值有：
	// <li>ddos_attackFlux_protocol ：攻击总流量协议类型分布排行 ；</li>
	// <li>ddos_attackPackageNum_protocol ：攻击总包量协议类型分布排行 ；</li>
	// <li>ddos_attackNum_attackType ：攻击总次数攻击类型分布排行 ；</li>
	// <li>ddos_attackNum_sregion ：攻击总次数攻击源地区分布排行 ；</li>
	// <li>ddos_attackFlux_sip ：攻击总流量攻击源ip分布排行 ；</li>
	// <li>ddos_attackFlux_sregion ：攻击总流量攻击源地区分布排行 。</li>
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`

	// 查询前多少个，传值为0返回全量。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 站点id集合，不填默认选择全部站点。
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// ddos策略组id 集合，不填默认选择全部策略id。
	PolicyIds []*int64 `json:"PolicyIds,omitempty" name:"PolicyIds"`

	// 端口号。
	Port *int64 `json:"Port,omitempty" name:"Port"`

	// 协议类型，取值有：
	// <li>tcp ；</li>
	// <li>udp ；</li>
	// <li>all 。</li>
	ProtocolType *string `json:"ProtocolType,omitempty" name:"ProtocolType"`

	// 攻击类型，取值有：
	// <li>flood ；</li>
	// <li>icmpFlood ；</li>
	// <li>all 。</li>
	AttackType *string `json:"AttackType,omitempty" name:"AttackType"`
}

func (r *DescribeDDosAttackTopDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDDosAttackTopDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "MetricName")
	delete(f, "Limit")
	delete(f, "ZoneIds")
	delete(f, "PolicyIds")
	delete(f, "Port")
	delete(f, "ProtocolType")
	delete(f, "AttackType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDDosAttackTopDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDDosAttackTopDataResponseParams struct {
	// top数据内容
	Data []*TopNEntry `json:"Data,omitempty" name:"Data"`

	// 请求响应状态，取值有：
	// <li>1 ：失败 ；</li>
	// <li>0 ：成功 。</li>
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 请求响应消息。
	Msg *string `json:"Msg,omitempty" name:"Msg"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDDosAttackTopDataResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDDosAttackTopDataResponseParams `json:"Response"`
}

func (r *DescribeDDosAttackTopDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDDosAttackTopDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDDosMajorAttackEventRequestParams struct {
	// 开始时间。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 分页拉取的最大返回结果数。最大值：1000。
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`

	// 分页拉取的起始页号。最小值：1。
	PageNo *int64 `json:"PageNo,omitempty" name:"PageNo"`

	// ddos 策略组id集合，不填默认选择全部策略id。
	PolicyIds []*int64 `json:"PolicyIds,omitempty" name:"PolicyIds"`

	// 协议类型，取值有：
	// <li>tcp ；</li>
	// <li>udp ；</li>
	// <li>all 。</li>
	ProtocolType *string `json:"ProtocolType,omitempty" name:"ProtocolType"`

	// 站点id列表，不填默认选择全部站点。
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`
}

type DescribeDDosMajorAttackEventRequest struct {
	*tchttp.BaseRequest
	
	// 开始时间。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 分页拉取的最大返回结果数。最大值：1000。
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`

	// 分页拉取的起始页号。最小值：1。
	PageNo *int64 `json:"PageNo,omitempty" name:"PageNo"`

	// ddos 策略组id集合，不填默认选择全部策略id。
	PolicyIds []*int64 `json:"PolicyIds,omitempty" name:"PolicyIds"`

	// 协议类型，取值有：
	// <li>tcp ；</li>
	// <li>udp ；</li>
	// <li>all 。</li>
	ProtocolType *string `json:"ProtocolType,omitempty" name:"ProtocolType"`

	// 站点id列表，不填默认选择全部站点。
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`
}

func (r *DescribeDDosMajorAttackEventRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDDosMajorAttackEventRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "PageSize")
	delete(f, "PageNo")
	delete(f, "PolicyIds")
	delete(f, "ProtocolType")
	delete(f, "ZoneIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDDosMajorAttackEventRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDDosMajorAttackEventResponseParams struct {
	// DDos查询主攻击事件。
	Data *DDosMajorAttackEventData `json:"Data,omitempty" name:"Data"`

	// 请求响应状态，取值有：
	// <li>1 ：失败 ；</li>
	// <li>0 ：成功 。</li>
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 请求响应消息。
	Msg *string `json:"Msg,omitempty" name:"Msg"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDDosMajorAttackEventResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDDosMajorAttackEventResponseParams `json:"Response"`
}

func (r *DescribeDDosMajorAttackEventResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDDosMajorAttackEventResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDefaultCertificatesRequestParams struct {
	// Zone ID
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`
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

// Predefined struct for user
type DescribeDefaultCertificatesResponseParams struct {
	// 证书总数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 默认证书列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	CertInfo []*DefaultServerCertInfo `json:"CertInfo,omitempty" name:"CertInfo"`

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

// Predefined struct for user
type DescribeDnsDataResponseParams struct {
	// 统计曲线数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*DataItem `json:"Data,omitempty" name:"Data"`

	// 时间粒度
	// 注意：此字段可能返回 null，表示取不到有效值。
	Interval *string `json:"Interval,omitempty" name:"Interval"`

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

// Predefined struct for user
type DescribeDnsRecordsResponseParams struct {
	// 总数，用于分页查询
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// DNS 记录列表
	Records []*DnsRecord `json:"Records,omitempty" name:"Records"`

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
	// 站点 ID
	Id *string `json:"Id,omitempty" name:"Id"`
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

// Predefined struct for user
type DescribeDnssecResponseParams struct {
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
type DescribeHostsCertificateRequestParams struct {
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

// Predefined struct for user
type DescribeHostsCertificateResponseParams struct {
	// 总数，用于分页查询
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 域名证书配置列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Hosts []*HostCertSetting `json:"Hosts,omitempty" name:"Hosts"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeHostsCertificateResponse struct {
	*tchttp.BaseResponse
	Response *DescribeHostsCertificateResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeHostsSettingRequestParams struct {
	// 站点ID
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 分页查询偏移量，默认为 0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页查询限制数目，默认为 100，最大可设置为 1000
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 指定域名查询
	Hosts []*string `json:"Hosts,omitempty" name:"Hosts"`
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

// Predefined struct for user
type DescribeHostsSettingResponseParams struct {
	// 域名列表
	Hosts []*DetailHost `json:"Hosts,omitempty" name:"Hosts"`

	// 域名数量
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
type DescribeIdentificationRequestParams struct {
	// 站点名称
	Name *string `json:"Name,omitempty" name:"Name"`
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

// Predefined struct for user
type DescribeIdentificationResponseParams struct {
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
}

type DescribeIdentificationResponse struct {
	*tchttp.BaseResponse
	Response *DescribeIdentificationResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeLoadBalancingDetailRequestParams struct {
	// 站点ID
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 负载均衡ID
	LoadBalancingId *string `json:"LoadBalancingId,omitempty" name:"LoadBalancingId"`
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

// Predefined struct for user
type DescribeLoadBalancingDetailResponseParams struct {
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
}

type DescribeLoadBalancingDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLoadBalancingDetailResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeLoadBalancingRequestParams struct {
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

// Predefined struct for user
type DescribeLoadBalancingResponseParams struct {
	// 记录总数
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 负载均衡信息
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
type DescribeOriginGroupDetailRequestParams struct {
	// 源站组ID
	OriginId *string `json:"OriginId,omitempty" name:"OriginId"`

	// 站点ID
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`
}

type DescribeOriginGroupDetailRequest struct {
	*tchttp.BaseRequest
	
	// 源站组ID
	OriginId *string `json:"OriginId,omitempty" name:"OriginId"`

	// 站点ID
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`
}

func (r *DescribeOriginGroupDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOriginGroupDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "OriginId")
	delete(f, "ZoneId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeOriginGroupDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOriginGroupDetailResponseParams struct {
	// 源站组ID
	OriginId *string `json:"OriginId,omitempty" name:"OriginId"`

	// 源站组名称
	OriginName *string `json:"OriginName,omitempty" name:"OriginName"`

	// 源站组配置类型
	// area：表示按照Record记录中的Area字段进行按客户端IP所在区域回源。
	// weight：表示按照Record记录中的Weight字段进行按权重回源。
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

	// 当前源站组是否被四层代理使用。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationProxyUsed *bool `json:"ApplicationProxyUsed,omitempty" name:"ApplicationProxyUsed"`

	// 当前源站组是否被负载均衡使用。
	// 注意：此字段可能返回 null，表示取不到有效值。
	LoadBalancingUsed *bool `json:"LoadBalancingUsed,omitempty" name:"LoadBalancingUsed"`

	// 使用当前源站组的负载均衡的类型：
	// none：未被使用
	// dns_only：被仅DNS类型负载均衡使用
	// proxied：被代理加速类型负载均衡使用
	// both：同时被仅DNS和代理加速类型负载均衡使用
	// 注意：此字段可能返回 null，表示取不到有效值。
	LoadBalancingUsedType *string `json:"LoadBalancingUsedType,omitempty" name:"LoadBalancingUsedType"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeOriginGroupDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeOriginGroupDetailResponseParams `json:"Response"`
}

func (r *DescribeOriginGroupDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOriginGroupDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOriginGroupRequestParams struct {
	// 分页参数Offset
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 分页参数Limit
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 过滤参数
	Filters []*OriginFilter `json:"Filters,omitempty" name:"Filters"`

	// 站点ID
	// 不填写获取所有站点源站组
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`
}

type DescribeOriginGroupRequest struct {
	*tchttp.BaseRequest
	
	// 分页参数Offset
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 分页参数Limit
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 过滤参数
	Filters []*OriginFilter `json:"Filters,omitempty" name:"Filters"`

	// 站点ID
	// 不填写获取所有站点源站组
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`
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
	delete(f, "ZoneId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeOriginGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOriginGroupResponseParams struct {
	// 源站组信息
	Data []*OriginGroup `json:"Data,omitempty" name:"Data"`

	// 记录总数
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

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
	// RFC3339格式，客户端时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// RFC3339格式，客户端时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 指标列表，支持的指标
	// l7Flow_outFlux: 访问流量
	// l7Flow_request: 访问请求数
	// l7Flow_outBandwidth: 访问带宽
	MetricNames []*string `json:"MetricNames,omitempty" name:"MetricNames"`

	// 时间间隔，选填{min, 5min, hour, day, week}
	Interval *string `json:"Interval,omitempty" name:"Interval"`

	// ZoneId列表，仅在zone/domain维度下查询时该参数有效
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// Domain列表，仅在domain维度下查询时该参数有效
	Domains []*string `json:"Domains,omitempty" name:"Domains"`

	// 协议类型， 选填{http,http2,https,all}
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
}

type DescribeOverviewL7DataRequest struct {
	*tchttp.BaseRequest
	
	// RFC3339格式，客户端时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// RFC3339格式，客户端时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 指标列表，支持的指标
	// l7Flow_outFlux: 访问流量
	// l7Flow_request: 访问请求数
	// l7Flow_outBandwidth: 访问带宽
	MetricNames []*string `json:"MetricNames,omitempty" name:"MetricNames"`

	// 时间间隔，选填{min, 5min, hour, day, week}
	Interval *string `json:"Interval,omitempty" name:"Interval"`

	// ZoneId列表，仅在zone/domain维度下查询时该参数有效
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// Domain列表，仅在domain维度下查询时该参数有效
	Domains []*string `json:"Domains,omitempty" name:"Domains"`

	// 协议类型， 选填{http,http2,https,all}
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
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
	delete(f, "Interval")
	delete(f, "ZoneIds")
	delete(f, "Domains")
	delete(f, "Protocol")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeOverviewL7DataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOverviewL7DataResponseParams struct {
	// 查询维度
	Type *string `json:"Type,omitempty" name:"Type"`

	// 时间间隔
	Interval *string `json:"Interval,omitempty" name:"Interval"`

	// 详细数据
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

// Predefined struct for user
type DescribePrefetchTasksResponseParams struct {
	// 该查询条件总共条目数
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 任务结果列表
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

// Predefined struct for user
type DescribePurgeTasksResponseParams struct {
	// 该查询条件总共条目数
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 任务结果列表
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
type DescribeSecurityPolicyListRequestParams struct {
	// 一级域名
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`
}

type DescribeSecurityPolicyListRequest struct {
	*tchttp.BaseRequest
	
	// 一级域名
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
	// 防护资源列表
	Entities []*SecurityEntity `json:"Entities,omitempty" name:"Entities"`

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
type DescribeSecurityPolicyManagedRulesIdRequestParams struct {
	// 规则id集合
	RuleId []*int64 `json:"RuleId,omitempty" name:"RuleId"`
}

type DescribeSecurityPolicyManagedRulesIdRequest struct {
	*tchttp.BaseRequest
	
	// 规则id集合
	RuleId []*int64 `json:"RuleId,omitempty" name:"RuleId"`
}

func (r *DescribeSecurityPolicyManagedRulesIdRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecurityPolicyManagedRulesIdRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RuleId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSecurityPolicyManagedRulesIdRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSecurityPolicyManagedRulesIdResponseParams struct {
	// 返回总数
	Total *int64 `json:"Total,omitempty" name:"Total"`

	// 门神规则
	Rules []*ManagedRule `json:"Rules,omitempty" name:"Rules"`

	// 返回总数
	Count *int64 `json:"Count,omitempty" name:"Count"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeSecurityPolicyManagedRulesIdResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSecurityPolicyManagedRulesIdResponseParams `json:"Response"`
}

func (r *DescribeSecurityPolicyManagedRulesIdResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecurityPolicyManagedRulesIdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSecurityPolicyManagedRulesRequestParams struct {
	// 一级域名
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 子域名/应用名
	Entity *string `json:"Entity,omitempty" name:"Entity"`

	// 页数
	Page *int64 `json:"Page,omitempty" name:"Page"`

	// 每页数量
	PerPage *int64 `json:"PerPage,omitempty" name:"PerPage"`
}

type DescribeSecurityPolicyManagedRulesRequest struct {
	*tchttp.BaseRequest
	
	// 一级域名
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 子域名/应用名
	Entity *string `json:"Entity,omitempty" name:"Entity"`

	// 页数
	Page *int64 `json:"Page,omitempty" name:"Page"`

	// 每页数量
	PerPage *int64 `json:"PerPage,omitempty" name:"PerPage"`
}

func (r *DescribeSecurityPolicyManagedRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecurityPolicyManagedRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "Entity")
	delete(f, "Page")
	delete(f, "PerPage")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSecurityPolicyManagedRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSecurityPolicyManagedRulesResponseParams struct {
	// 本次返回的规则数
	Count *int64 `json:"Count,omitempty" name:"Count"`

	// 门神规则
	Rules []*ManagedRule `json:"Rules,omitempty" name:"Rules"`

	// 总规则数
	Total *int64 `json:"Total,omitempty" name:"Total"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeSecurityPolicyManagedRulesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSecurityPolicyManagedRulesResponseParams `json:"Response"`
}

func (r *DescribeSecurityPolicyManagedRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecurityPolicyManagedRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSecurityPolicyRegionsRequestParams struct {

}

type DescribeSecurityPolicyRegionsRequest struct {
	*tchttp.BaseRequest
	
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
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSecurityPolicyRegionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSecurityPolicyRegionsResponseParams struct {
	// 总数
	Count *int64 `json:"Count,omitempty" name:"Count"`

	// 地域信息
	GeoIp []*GeoIp `json:"GeoIp,omitempty" name:"GeoIp"`

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
	// 一级域名
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 二级域名
	Entity *string `json:"Entity,omitempty" name:"Entity"`
}

type DescribeSecurityPolicyRequest struct {
	*tchttp.BaseRequest
	
	// 一级域名
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 二级域名
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
	// 用户id
	AppId *int64 `json:"AppId,omitempty" name:"AppId"`

	// 一级域名
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 二级域名
	Entity *string `json:"Entity,omitempty" name:"Entity"`

	// 安全配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Config *SecurityConfig `json:"Config,omitempty" name:"Config"`

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
	// 一级域名
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 子域名/应用名
	Entity *string `json:"Entity,omitempty" name:"Entity"`
}

type DescribeSecurityPortraitRulesRequest struct {
	*tchttp.BaseRequest
	
	// 一级域名
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 子域名/应用名
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
	// 本次返回的规则数
	Count *int64 `json:"Count,omitempty" name:"Count"`

	// Bot用户画像规则
	Rules []*PortraitManagedRuleDetail `json:"Rules,omitempty" name:"Rules"`

	// 总规则数
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
type DescribeTimingL4DataRequestParams struct {
	// RFC3339格式，客户端时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// RFC3339格式，客户端时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 支持的指标：
	// l4Flow_connections: 访问连接数
	// l4Flow_flux: 访问总流量
	// l4Flow_inFlux: 访问入流量
	// l4Flow_outFlux: 访问出流量
	MetricNames []*string `json:"MetricNames,omitempty" name:"MetricNames"`

	// 站点id列表
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// 该字段已废弃，请使用ProxyIds字段
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// 该字段当前无效
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// 时间间隔，选填{min, 5min, hour, day}
	Interval *string `json:"Interval,omitempty" name:"Interval"`

	// 该字段当前无效，请使用Filter筛选
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// 支持的 Filter：proxyd,ruleId
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 四层实例列表
	ProxyIds []*string `json:"ProxyIds,omitempty" name:"ProxyIds"`
}

type DescribeTimingL4DataRequest struct {
	*tchttp.BaseRequest
	
	// RFC3339格式，客户端时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// RFC3339格式，客户端时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 支持的指标：
	// l4Flow_connections: 访问连接数
	// l4Flow_flux: 访问总流量
	// l4Flow_inFlux: 访问入流量
	// l4Flow_outFlux: 访问出流量
	MetricNames []*string `json:"MetricNames,omitempty" name:"MetricNames"`

	// 站点id列表
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// 该字段已废弃，请使用ProxyIds字段
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// 该字段当前无效
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// 时间间隔，选填{min, 5min, hour, day}
	Interval *string `json:"Interval,omitempty" name:"Interval"`

	// 该字段当前无效，请使用Filter筛选
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// 支持的 Filter：proxyd,ruleId
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 四层实例列表
	ProxyIds []*string `json:"ProxyIds,omitempty" name:"ProxyIds"`
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
	delete(f, "InstanceIds")
	delete(f, "Protocol")
	delete(f, "Interval")
	delete(f, "RuleId")
	delete(f, "Filters")
	delete(f, "ProxyIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTimingL4DataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTimingL4DataResponseParams struct {
	// 查询维度
	Type *string `json:"Type,omitempty" name:"Type"`

	// 时间间隔
	Interval *string `json:"Interval,omitempty" name:"Interval"`

	// 详细数据
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
	// RFC3339标准，客户端时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// RFC3339标准，客户端时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 指标列表，支持的指标
	// l7Flow_outFlux: 访问流量
	// l7Flow_request: 访问请求数
	// l7Flow_outBandwidth: 访问带宽
	MetricNames []*string `json:"MetricNames,omitempty" name:"MetricNames"`

	// 时间间隔，选填{min, 5min, hour, day, week}
	Interval *string `json:"Interval,omitempty" name:"Interval"`

	// ZoneId数组
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// 筛选条件
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

type DescribeTimingL7AnalysisDataRequest struct {
	*tchttp.BaseRequest
	
	// RFC3339标准，客户端时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// RFC3339标准，客户端时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 指标列表，支持的指标
	// l7Flow_outFlux: 访问流量
	// l7Flow_request: 访问请求数
	// l7Flow_outBandwidth: 访问带宽
	MetricNames []*string `json:"MetricNames,omitempty" name:"MetricNames"`

	// 时间间隔，选填{min, 5min, hour, day, week}
	Interval *string `json:"Interval,omitempty" name:"Interval"`

	// ZoneId数组
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// 筛选条件
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
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
	delete(f, "Interval")
	delete(f, "ZoneIds")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTimingL7AnalysisDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTimingL7AnalysisDataResponseParams struct {
	// 详细数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*TimingDataRecord `json:"Data,omitempty" name:"Data"`

	// 查询维度
	Type *string `json:"Type,omitempty" name:"Type"`

	// 时间间隔
	Interval *string `json:"Interval,omitempty" name:"Interval"`

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
	// RFC3339标准，客户端时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// RFC3339标准，客户端时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 时序类访问流量指标列表，支持的指标
	// l7Cache_outFlux: 访问流量
	// l7Cache_request: 访问请求数
	MetricNames []*string `json:"MetricNames,omitempty" name:"MetricNames"`

	// 时间间隔，选填{min, 5min, hour, day, week}
	Interval *string `json:"Interval,omitempty" name:"Interval"`

	// 站点id列表
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// 筛选条件，筛选EO/源站响应如下：
	// EO响应：{Key: "cacheType", Value: ["hit"], Operator: "equals"}；
	// 源站响应：{Key: "cacheType", Value: ["miss", "dynamic"], Operator: "equals"}
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

type DescribeTimingL7CacheDataRequest struct {
	*tchttp.BaseRequest
	
	// RFC3339标准，客户端时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// RFC3339标准，客户端时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 时序类访问流量指标列表，支持的指标
	// l7Cache_outFlux: 访问流量
	// l7Cache_request: 访问请求数
	MetricNames []*string `json:"MetricNames,omitempty" name:"MetricNames"`

	// 时间间隔，选填{min, 5min, hour, day, week}
	Interval *string `json:"Interval,omitempty" name:"Interval"`

	// 站点id列表
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// 筛选条件，筛选EO/源站响应如下：
	// EO响应：{Key: "cacheType", Value: ["hit"], Operator: "equals"}；
	// 源站响应：{Key: "cacheType", Value: ["miss", "dynamic"], Operator: "equals"}
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
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
	delete(f, "Interval")
	delete(f, "ZoneIds")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTimingL7CacheDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTimingL7CacheDataResponseParams struct {
	// 详细数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*TimingDataRecord `json:"Data,omitempty" name:"Data"`

	// 查询维度
	Type *string `json:"Type,omitempty" name:"Type"`

	// 时间间隔
	Interval *string `json:"Interval,omitempty" name:"Interval"`

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
	// RFC3339标准，客户端时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// RFC3339标准，客户端时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 时序类访问流量指标
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`

	// topN,填0时返回全量数据
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 时间间隔，选填{min, 5min, hour, day, week}
	Interval *string `json:"Interval,omitempty" name:"Interval"`

	// ZoneId数组
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// 筛选条件
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

type DescribeTopL7AnalysisDataRequest struct {
	*tchttp.BaseRequest
	
	// RFC3339标准，客户端时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// RFC3339标准，客户端时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 时序类访问流量指标
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`

	// topN,填0时返回全量数据
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 时间间隔，选填{min, 5min, hour, day, week}
	Interval *string `json:"Interval,omitempty" name:"Interval"`

	// ZoneId数组
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// 筛选条件
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
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
	delete(f, "Limit")
	delete(f, "Interval")
	delete(f, "ZoneIds")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTopL7AnalysisDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTopL7AnalysisDataResponseParams struct {
	// top详细数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*TopDataRecord `json:"Data,omitempty" name:"Data"`

	// 查询维度
	Type *string `json:"Type,omitempty" name:"Type"`

	// 查询指标
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`

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
	// RFC3339标准，客户端时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// RFC3339标准，客户端时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 时序类访问流量指标
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`

	// topN,填0时返回全量数据
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 时间间隔，选填{min, 5min, hour, day, week}
	Interval *string `json:"Interval,omitempty" name:"Interval"`

	// ZoneId数组
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// 筛选条件
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

type DescribeTopL7CacheDataRequest struct {
	*tchttp.BaseRequest
	
	// RFC3339标准，客户端时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// RFC3339标准，客户端时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 时序类访问流量指标
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`

	// topN,填0时返回全量数据
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 时间间隔，选填{min, 5min, hour, day, week}
	Interval *string `json:"Interval,omitempty" name:"Interval"`

	// ZoneId数组
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// 筛选条件
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
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
	delete(f, "Limit")
	delete(f, "Interval")
	delete(f, "ZoneIds")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTopL7CacheDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTopL7CacheDataResponseParams struct {
	// top详细数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*TopDataRecord `json:"Data,omitempty" name:"Data"`

	// 查询维度
	Type *string `json:"Type,omitempty" name:"Type"`

	// 查询指标
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`

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
type DescribeWebManagedRulesAttackEventsRequestParams struct {
	// 开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 条数
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`

	// 当前页
	PageNo *int64 `json:"PageNo,omitempty" name:"PageNo"`

	// ddos策略组id列表
	PolicyIds []*int64 `json:"PolicyIds,omitempty" name:"PolicyIds"`

	// 站点集合
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// 子域名列表
	Domains []*string `json:"Domains,omitempty" name:"Domains"`

	// 选填{Y、N},默认为Y；Y：展示，N：不展示
	IsShowDetail *string `json:"IsShowDetail,omitempty" name:"IsShowDetail"`
}

type DescribeWebManagedRulesAttackEventsRequest struct {
	*tchttp.BaseRequest
	
	// 开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 条数
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`

	// 当前页
	PageNo *int64 `json:"PageNo,omitempty" name:"PageNo"`

	// ddos策略组id列表
	PolicyIds []*int64 `json:"PolicyIds,omitempty" name:"PolicyIds"`

	// 站点集合
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// 子域名列表
	Domains []*string `json:"Domains,omitempty" name:"Domains"`

	// 选填{Y、N},默认为Y；Y：展示，N：不展示
	IsShowDetail *string `json:"IsShowDetail,omitempty" name:"IsShowDetail"`
}

func (r *DescribeWebManagedRulesAttackEventsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWebManagedRulesAttackEventsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "PageSize")
	delete(f, "PageNo")
	delete(f, "PolicyIds")
	delete(f, "ZoneIds")
	delete(f, "Domains")
	delete(f, "IsShowDetail")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeWebManagedRulesAttackEventsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWebManagedRulesAttackEventsResponseParams struct {
	// Web攻击事件数据
	Data *WebEventData `json:"Data,omitempty" name:"Data"`

	// 状态，1:失败，0:成功
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 返回数据
	Msg *string `json:"Msg,omitempty" name:"Msg"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeWebManagedRulesAttackEventsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeWebManagedRulesAttackEventsResponseParams `json:"Response"`
}

func (r *DescribeWebManagedRulesAttackEventsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWebManagedRulesAttackEventsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWebManagedRulesDataRequestParams struct {
	// 开始时间，RFC3339格式。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间，RFC3339格式。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 统计指标列表，取值有：
	// <li>waf_interceptNum ：waf拦截次数 。</li>
	MetricNames []*string `json:"MetricNames,omitempty" name:"MetricNames"`

	// 站点id列表，不填默认选择全部站点。
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// 子域名列表，不填默认选择子域名。
	Domains []*string `json:"Domains,omitempty" name:"Domains"`

	// 该字段已废弃，请勿传。
	ProtocolType *string `json:"ProtocolType,omitempty" name:"ProtocolType"`

	// 该字段已废弃，请勿传。
	AttackType *string `json:"AttackType,omitempty" name:"AttackType"`

	// 查询时间粒度，取值有：
	// <li>min ：1分钟 ；</li>
	// <li>5min ：5分钟 ；</li>
	// <li>hour ：1小时 ；</li>
	// <li>day ：1天 。</li>
	Interval *string `json:"Interval,omitempty" name:"Interval"`

	// 筛选条件，取值有：
	// <li>action ：执行动作 。</li>
	QueryCondition []*QueryCondition `json:"QueryCondition,omitempty" name:"QueryCondition"`
}

type DescribeWebManagedRulesDataRequest struct {
	*tchttp.BaseRequest
	
	// 开始时间，RFC3339格式。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间，RFC3339格式。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 统计指标列表，取值有：
	// <li>waf_interceptNum ：waf拦截次数 。</li>
	MetricNames []*string `json:"MetricNames,omitempty" name:"MetricNames"`

	// 站点id列表，不填默认选择全部站点。
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// 子域名列表，不填默认选择子域名。
	Domains []*string `json:"Domains,omitempty" name:"Domains"`

	// 该字段已废弃，请勿传。
	ProtocolType *string `json:"ProtocolType,omitempty" name:"ProtocolType"`

	// 该字段已废弃，请勿传。
	AttackType *string `json:"AttackType,omitempty" name:"AttackType"`

	// 查询时间粒度，取值有：
	// <li>min ：1分钟 ；</li>
	// <li>5min ：5分钟 ；</li>
	// <li>hour ：1小时 ；</li>
	// <li>day ：1天 。</li>
	Interval *string `json:"Interval,omitempty" name:"Interval"`

	// 筛选条件，取值有：
	// <li>action ：执行动作 。</li>
	QueryCondition []*QueryCondition `json:"QueryCondition,omitempty" name:"QueryCondition"`
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
	delete(f, "ProtocolType")
	delete(f, "AttackType")
	delete(f, "Interval")
	delete(f, "QueryCondition")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeWebManagedRulesDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWebManagedRulesDataResponseParams struct {
	// Web攻击日志实体。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*SecEntry `json:"Data,omitempty" name:"Data"`

	// 请求响应状态，取值有：
	// <li>1 ：失败 ；</li>
	// <li>0 ：成功 。</li>
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 请求响应消息。
	Msg *string `json:"Msg,omitempty" name:"Msg"`

	// 查询时间粒度，取值有：
	// <li>min ：1分钟 ；</li>
	// <li>5min ：5分钟 ；</li>
	// <li>hour ：1小时 ；</li>
	// <li>day ：1天 。</li>
	Interval *string `json:"Interval,omitempty" name:"Interval"`

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
type DescribeWebManagedRulesLogRequestParams struct {
	// 起始时间。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 分页拉取的最大返回结果数。最大值：1000。
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`

	// 分页拉取的起始页号。最小值：1。
	PageNo *int64 `json:"PageNo,omitempty" name:"PageNo"`

	// 站点集合，不填默认选择全部站点。
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// 域名集合，不填默认选择全部子域名。
	Domains []*string `json:"Domains,omitempty" name:"Domains"`

	// 筛选条件，取值有：
	// <li>attackType ：攻击类型 ；</li>
	// <li>riskLevel ：风险等级 ；</li>
	// <li>action ：执行动作（处置方式） ；</li>
	// <li>ruleId ：规则id ；</li>
	// <li>sipCountryCode ：ip所在国家 ；</li>
	// <li>attackIp ：攻击ip ；</li>
	// <li>oriDomain ：被攻击的子域名 ；</li>
	// <li>eventId ：事件id ；</li>
	// <li>ua ：用户代理 ；</li>
	// <li>requestMethod ：请求方法 ；</li>
	// <li>uri ：统一资源标识符 。</li>
	QueryCondition []*QueryCondition `json:"QueryCondition,omitempty" name:"QueryCondition"`
}

type DescribeWebManagedRulesLogRequest struct {
	*tchttp.BaseRequest
	
	// 起始时间。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 分页拉取的最大返回结果数。最大值：1000。
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`

	// 分页拉取的起始页号。最小值：1。
	PageNo *int64 `json:"PageNo,omitempty" name:"PageNo"`

	// 站点集合，不填默认选择全部站点。
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// 域名集合，不填默认选择全部子域名。
	Domains []*string `json:"Domains,omitempty" name:"Domains"`

	// 筛选条件，取值有：
	// <li>attackType ：攻击类型 ；</li>
	// <li>riskLevel ：风险等级 ；</li>
	// <li>action ：执行动作（处置方式） ；</li>
	// <li>ruleId ：规则id ；</li>
	// <li>sipCountryCode ：ip所在国家 ；</li>
	// <li>attackIp ：攻击ip ；</li>
	// <li>oriDomain ：被攻击的子域名 ；</li>
	// <li>eventId ：事件id ；</li>
	// <li>ua ：用户代理 ；</li>
	// <li>requestMethod ：请求方法 ；</li>
	// <li>uri ：统一资源标识符 。</li>
	QueryCondition []*QueryCondition `json:"QueryCondition,omitempty" name:"QueryCondition"`
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
	delete(f, "PageSize")
	delete(f, "PageNo")
	delete(f, "ZoneIds")
	delete(f, "Domains")
	delete(f, "QueryCondition")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeWebManagedRulesLogRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWebManagedRulesLogResponseParams struct {
	// web攻击日志数据内容。
	Data *WebLogData `json:"Data,omitempty" name:"Data"`

	// 请求响应状态，取值有：
	// <li>1 ：失败 ；</li>
	// <li>0 ：成功 。</li>
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 请求响应信息。
	Msg *string `json:"Msg,omitempty" name:"Msg"`

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
type DescribeWebManagedRulesTopDataRequestParams struct {
	// 开始时间。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 统计指标列表，取值有：
	// <li>waf_requestNum_url ：url请求数排行 ；</li>
	// <li>waf_requestNum_cip：客户端ip请求数排行 ；</li>
	// <li>waf_cipRequestNum_region ：客户端区域请求数排行 。</li>
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`

	// 查询前多少个，传值为0返回全量。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 站点id列表，不填默认选择全部站点。
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// 该字段已废弃，请勿传。
	PolicyIds []*int64 `json:"PolicyIds,omitempty" name:"PolicyIds"`

	// 该字段已废弃，请勿传。
	Port *int64 `json:"Port,omitempty" name:"Port"`

	// 该字段已废弃，请勿传。
	ProtocolType *string `json:"ProtocolType,omitempty" name:"ProtocolType"`

	// 该字段已废弃，请勿传。
	AttackType *string `json:"AttackType,omitempty" name:"AttackType"`

	// 域名列表，不填默认选择全部子域名。
	Domains []*string `json:"Domains,omitempty" name:"Domains"`

	// 查询时间粒度，取值有：
	// <li>min ：1分钟 ；</li>
	// <li>5min ：5分钟 ；</li>
	// <li>hour ：1小时 ；</li>
	// <li>day ：1天 。</li>
	Interval *string `json:"Interval,omitempty" name:"Interval"`

	// 筛选条件，取值有：
	// <li>action ：执行动作 。</li>
	QueryCondition []*QueryCondition `json:"QueryCondition,omitempty" name:"QueryCondition"`
}

type DescribeWebManagedRulesTopDataRequest struct {
	*tchttp.BaseRequest
	
	// 开始时间。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 统计指标列表，取值有：
	// <li>waf_requestNum_url ：url请求数排行 ；</li>
	// <li>waf_requestNum_cip：客户端ip请求数排行 ；</li>
	// <li>waf_cipRequestNum_region ：客户端区域请求数排行 。</li>
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`

	// 查询前多少个，传值为0返回全量。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 站点id列表，不填默认选择全部站点。
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// 该字段已废弃，请勿传。
	PolicyIds []*int64 `json:"PolicyIds,omitempty" name:"PolicyIds"`

	// 该字段已废弃，请勿传。
	Port *int64 `json:"Port,omitempty" name:"Port"`

	// 该字段已废弃，请勿传。
	ProtocolType *string `json:"ProtocolType,omitempty" name:"ProtocolType"`

	// 该字段已废弃，请勿传。
	AttackType *string `json:"AttackType,omitempty" name:"AttackType"`

	// 域名列表，不填默认选择全部子域名。
	Domains []*string `json:"Domains,omitempty" name:"Domains"`

	// 查询时间粒度，取值有：
	// <li>min ：1分钟 ；</li>
	// <li>5min ：5分钟 ；</li>
	// <li>hour ：1小时 ；</li>
	// <li>day ：1天 。</li>
	Interval *string `json:"Interval,omitempty" name:"Interval"`

	// 筛选条件，取值有：
	// <li>action ：执行动作 。</li>
	QueryCondition []*QueryCondition `json:"QueryCondition,omitempty" name:"QueryCondition"`
}

func (r *DescribeWebManagedRulesTopDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWebManagedRulesTopDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "MetricName")
	delete(f, "Limit")
	delete(f, "ZoneIds")
	delete(f, "PolicyIds")
	delete(f, "Port")
	delete(f, "ProtocolType")
	delete(f, "AttackType")
	delete(f, "Domains")
	delete(f, "Interval")
	delete(f, "QueryCondition")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeWebManagedRulesTopDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWebManagedRulesTopDataResponseParams struct {
	// top数据内容。
	Data []*TopNEntry `json:"Data,omitempty" name:"Data"`

	// 请求响应状态，取值有：
	// <li>1 ：失败 ；</li>
	// <li>0 ：成功 。</li>
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 请求响应消息。
	Msg *string `json:"Msg,omitempty" name:"Msg"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeWebManagedRulesTopDataResponse struct {
	*tchttp.BaseResponse
	Response *DescribeWebManagedRulesTopDataResponseParams `json:"Response"`
}

func (r *DescribeWebManagedRulesTopDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWebManagedRulesTopDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWebProtectionAttackEventsRequestParams struct {
	// 开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 条数
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`

	// 当前页
	PageNo *int64 `json:"PageNo,omitempty" name:"PageNo"`

	// 域名
	Domains []*string `json:"Domains,omitempty" name:"Domains"`

	// 站点集合
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`
}

type DescribeWebProtectionAttackEventsRequest struct {
	*tchttp.BaseRequest
	
	// 开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 条数
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`

	// 当前页
	PageNo *int64 `json:"PageNo,omitempty" name:"PageNo"`

	// 域名
	Domains []*string `json:"Domains,omitempty" name:"Domains"`

	// 站点集合
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`
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
	delete(f, "PageSize")
	delete(f, "PageNo")
	delete(f, "Domains")
	delete(f, "ZoneIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeWebProtectionAttackEventsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWebProtectionAttackEventsResponseParams struct {
	// DDos攻击事件数据
	Data *CCInterceptEventData `json:"Data,omitempty" name:"Data"`

	// 状态，1:失败，0:成功
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 返回消息
	Msg *string `json:"Msg,omitempty" name:"Msg"`

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
type DescribeWebProtectionDataRequestParams struct {
	// 开始时间，RFC3339格式。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间，RFC3339格式。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 统计指标列表，取值有：
	// <li>ccRate_interceptNum ：速率限制规则限制次数 ；</li>
	// <li>ccAcl_interceptNum ：自定义规则拦截次数 。</li>
	MetricNames []*string `json:"MetricNames,omitempty" name:"MetricNames"`

	// 站点id列表，不填默认选择全部站点。
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// 子域名列表，不填默认选择全部子域名。
	Domains []*string `json:"Domains,omitempty" name:"Domains"`

	// 该字段已废弃，请勿传。
	ProtocolType *string `json:"ProtocolType,omitempty" name:"ProtocolType"`

	// 该字段已废弃，请勿传。
	AttackType *string `json:"AttackType,omitempty" name:"AttackType"`

	// 查询时间粒度，取值有：
	// <li>min ：1分钟 ；</li>
	// <li>5min ：5分钟 ；</li>
	// <li>hour ：1小时 ；</li>
	// <li>day ：1天 。</li>
	Interval *string `json:"Interval,omitempty" name:"Interval"`

	// 筛选条件，取值有：
	// <li>action ：执行动作 。</li>
	QueryCondition []*QueryCondition `json:"QueryCondition,omitempty" name:"QueryCondition"`
}

type DescribeWebProtectionDataRequest struct {
	*tchttp.BaseRequest
	
	// 开始时间，RFC3339格式。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间，RFC3339格式。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 统计指标列表，取值有：
	// <li>ccRate_interceptNum ：速率限制规则限制次数 ；</li>
	// <li>ccAcl_interceptNum ：自定义规则拦截次数 。</li>
	MetricNames []*string `json:"MetricNames,omitempty" name:"MetricNames"`

	// 站点id列表，不填默认选择全部站点。
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// 子域名列表，不填默认选择全部子域名。
	Domains []*string `json:"Domains,omitempty" name:"Domains"`

	// 该字段已废弃，请勿传。
	ProtocolType *string `json:"ProtocolType,omitempty" name:"ProtocolType"`

	// 该字段已废弃，请勿传。
	AttackType *string `json:"AttackType,omitempty" name:"AttackType"`

	// 查询时间粒度，取值有：
	// <li>min ：1分钟 ；</li>
	// <li>5min ：5分钟 ；</li>
	// <li>hour ：1小时 ；</li>
	// <li>day ：1天 。</li>
	Interval *string `json:"Interval,omitempty" name:"Interval"`

	// 筛选条件，取值有：
	// <li>action ：执行动作 。</li>
	QueryCondition []*QueryCondition `json:"QueryCondition,omitempty" name:"QueryCondition"`
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
	delete(f, "ProtocolType")
	delete(f, "AttackType")
	delete(f, "Interval")
	delete(f, "QueryCondition")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeWebProtectionDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWebProtectionDataResponseParams struct {
	// 数据详情。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*SecEntry `json:"Data,omitempty" name:"Data"`

	// 请求响应状态，取值有：
	// <li>1 ：失败 ；</li>
	// <li>0 ：成功 。</li>
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 请求响应消息。
	Msg *string `json:"Msg,omitempty" name:"Msg"`

	// 查询时间粒度，取值有：
	// <li>min ：1分钟 ；</li>
	// <li>5min ：5分钟 ；</li>
	// <li>hour ：1小时 ；</li>
	// <li>day ：1天 。</li>
	Interval *string `json:"Interval,omitempty" name:"Interval"`

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
type DescribeWebProtectionLogRequestParams struct {
	// 起始时间。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 分页拉取的最大返回结果数。最大值：1000。
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`

	// 分页拉取的起始页号。最小值：1。
	PageNo *int64 `json:"PageNo,omitempty" name:"PageNo"`

	// 站点集合，不填默认查询所有站点。
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// 域名集合，不填默认查询所有域名。
	Domains []*string `json:"Domains,omitempty" name:"Domains"`

	// 筛选条件。
	// 限速规则日志中取值有：
	// <li>action ：执行动作（处置方式）；</li>
	// <li>ruleId ：规则id ；</li>
	// <li>oriDomain ：被攻击的子域名 ；</li>
	// <li>attackIp ：攻击ip 。</li>
	// 自定义规则日志中取值有：
	// <li>action ：执行动作（处置方式）；</li>
	// <li>ruleId ：规则id ；</li>
	// <li>oriDomain ：被攻击的子域名 ；</li>
	// <li>attackIp ：攻击ip ；</li>
	// <li>eventId ：事件id ；</li>
	// <li>ua ：用户代理 ；</li>
	// <li>requestMethod ：请求方法 ；</li>
	// <li>uri ：统一资源标识符 。</li>
	QueryCondition []*QueryCondition `json:"QueryCondition,omitempty" name:"QueryCondition"`

	// 日志类型，取值有：
	// <li>rate ：限速日志 ；</li>
	// <li>acl ：自定义规则日志 。</li>不填默认为rate。
	EntityType *string `json:"EntityType,omitempty" name:"EntityType"`
}

type DescribeWebProtectionLogRequest struct {
	*tchttp.BaseRequest
	
	// 起始时间。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 分页拉取的最大返回结果数。最大值：1000。
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`

	// 分页拉取的起始页号。最小值：1。
	PageNo *int64 `json:"PageNo,omitempty" name:"PageNo"`

	// 站点集合，不填默认查询所有站点。
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// 域名集合，不填默认查询所有域名。
	Domains []*string `json:"Domains,omitempty" name:"Domains"`

	// 筛选条件。
	// 限速规则日志中取值有：
	// <li>action ：执行动作（处置方式）；</li>
	// <li>ruleId ：规则id ；</li>
	// <li>oriDomain ：被攻击的子域名 ；</li>
	// <li>attackIp ：攻击ip 。</li>
	// 自定义规则日志中取值有：
	// <li>action ：执行动作（处置方式）；</li>
	// <li>ruleId ：规则id ；</li>
	// <li>oriDomain ：被攻击的子域名 ；</li>
	// <li>attackIp ：攻击ip ；</li>
	// <li>eventId ：事件id ；</li>
	// <li>ua ：用户代理 ；</li>
	// <li>requestMethod ：请求方法 ；</li>
	// <li>uri ：统一资源标识符 。</li>
	QueryCondition []*QueryCondition `json:"QueryCondition,omitempty" name:"QueryCondition"`

	// 日志类型，取值有：
	// <li>rate ：限速日志 ；</li>
	// <li>acl ：自定义规则日志 。</li>不填默认为rate。
	EntityType *string `json:"EntityType,omitempty" name:"EntityType"`
}

func (r *DescribeWebProtectionLogRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWebProtectionLogRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "PageSize")
	delete(f, "PageNo")
	delete(f, "ZoneIds")
	delete(f, "Domains")
	delete(f, "QueryCondition")
	delete(f, "EntityType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeWebProtectionLogRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWebProtectionLogResponseParams struct {
	// 限速拦截数据内容。
	Data *CCLogData `json:"Data,omitempty" name:"Data"`

	// 请求响应状态，取值有：
	// <li>1 ：失败 ；</li>
	// <li>0 ：成功 。</li>
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 请求响应信息。
	Msg *string `json:"Msg,omitempty" name:"Msg"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeWebProtectionLogResponse struct {
	*tchttp.BaseResponse
	Response *DescribeWebProtectionLogResponseParams `json:"Response"`
}

func (r *DescribeWebProtectionLogResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWebProtectionLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeZoneDDoSPolicyRequestParams struct {
	// 一级域名id
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`
}

type DescribeZoneDDoSPolicyRequest struct {
	*tchttp.BaseRequest
	
	// 一级域名id
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
	// 用户appid
	AppId *int64 `json:"AppId,omitempty" name:"AppId"`

	// 防护分区
	ShieldAreas []*ShieldArea `json:"ShieldAreas,omitempty" name:"ShieldAreas"`

	// 所有子域名信息，包含安全加速/内容加速
	// 注意：此字段可能返回 null，表示取不到有效值。
	Domains []*DDoSApplication `json:"Domains,omitempty" name:"Domains"`

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
type DescribeZoneDetailsRequestParams struct {
	// 站点 ID
	Id *string `json:"Id,omitempty" name:"Id"`
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

// Predefined struct for user
type DescribeZoneDetailsResponseParams struct {
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

	// 资源标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`

	// 计费资源
	// 注意：此字段可能返回 null，表示取不到有效值。
	Resources []*Resource `json:"Resources,omitempty" name:"Resources"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeZoneDetailsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeZoneDetailsResponseParams `json:"Response"`
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
	// 站点ID。
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 站点名称。
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 缓存过期时间配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Cache *CacheConfig `json:"Cache,omitempty" name:"Cache"`

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

	// Https 加速配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Https *Https `json:"Https,omitempty" name:"Https"`

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
	ClientIpHeader *ClientIp `json:"ClientIpHeader,omitempty" name:"ClientIpHeader"`

	// 缓存预刷新配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CachePrefresh *CachePrefresh `json:"CachePrefresh,omitempty" name:"CachePrefresh"`

	// Ipv6访问配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Ipv6 *Ipv6Access `json:"Ipv6,omitempty" name:"Ipv6"`

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
	// 分页参数，页偏移
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页参数，每页返回的站点个数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 查询条件过滤器，复杂类型
	Filters []*ZoneFilter `json:"Filters,omitempty" name:"Filters"`
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

// Predefined struct for user
type DescribeZonesResponseParams struct {
	// 符合条件的站点数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 站点详细信息列表
	// 注意：此字段可能返回 null，表示取不到有效值。
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

	// 域名是否开启了负载均衡，四层代理，安全
	// - lb 负载均衡
	// - security 安全
	// - l4 四层代理
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

// Predefined struct for user
type DownloadL7LogsRequestParams struct {
	// 起始时间(需严格按照RFC3339标准传参)
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间(需严格按照RFC3339标准传参)
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 每页展示条数
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`

	// 当前页
	PageNo *int64 `json:"PageNo,omitempty" name:"PageNo"`

	// 站点名集合
	Zones []*string `json:"Zones,omitempty" name:"Zones"`

	// 子域名集合
	Domains []*string `json:"Domains,omitempty" name:"Domains"`
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

	// 站点名集合
	Zones []*string `json:"Zones,omitempty" name:"Zones"`

	// 子域名集合
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

// Predefined struct for user
type DownloadL7LogsResponseParams struct {
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

type FailReason struct {
	// 失败原因
	Reason *string `json:"Reason,omitempty" name:"Reason"`

	// 处理失败的资源列表。
	// 该列表元素来源于输入参数中的Targets，因此格式和入参中的Targets保持一致
	Targets []*string `json:"Targets,omitempty" name:"Targets"`
}

type Filter struct {
	// 筛选维度
	Key *string `json:"Key,omitempty" name:"Key"`

	// 操作符
	Operator *string `json:"Operator,omitempty" name:"Operator"`

	// 筛选维度值
	Value []*string `json:"Value,omitempty" name:"Value"`
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
	// 地域ID
	RegionId *int64 `json:"RegionId,omitempty" name:"RegionId"`

	// 国家名
	Country *string `json:"Country,omitempty" name:"Country"`

	// 洲
	Continent *string `json:"Continent,omitempty" name:"Continent"`

	// 国家英文名
	CountryEn *string `json:"CountryEn,omitempty" name:"CountryEn"`

	// 洲
	ContinentEn *string `json:"ContinentEn,omitempty" name:"ContinentEn"`
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
}

// Predefined struct for user
type IdentifyZoneRequestParams struct {
	// 站点名称
	Name *string `json:"Name,omitempty" name:"Name"`
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

// Predefined struct for user
type IdentifyZoneResponseParams struct {
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

// Predefined struct for user
type ImportDnsRecordsRequestParams struct {
	// 站点 ID
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 文件内容
	File *string `json:"File,omitempty" name:"File"`
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

// Predefined struct for user
type ImportDnsRecordsResponseParams struct {
	// 记录 ID
	Ids []*string `json:"Ids,omitempty" name:"Ids"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ImportDnsRecordsResponse struct {
	*tchttp.BaseResponse
	Response *ImportDnsRecordsResponseParams `json:"Response"`
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

type IntelligenceRule struct {
	// 开关
	// 注意：此字段可能返回 null，表示取不到有效值。
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// 规则详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	Items []*IntelligenceRuleItem `json:"Items,omitempty" name:"Items"`
}

type IntelligenceRuleItem struct {
	// 恶意BOT
	// 注意：此字段可能返回 null，表示取不到有效值。
	Label *string `json:"Label,omitempty" name:"Label"`

	// 动作
	// 注意：此字段可能返回 null，表示取不到有效值。
	Action *string `json:"Action,omitempty" name:"Action"`
}

type IpTableConfig struct {
	// 开关
	// 注意：此字段可能返回 null，表示取不到有效值。
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// []
	// 注意：此字段可能返回 null，表示取不到有效值。
	Rules []*IpTableRule `json:"Rules,omitempty" name:"Rules"`
}

type IpTableRule struct {
	// 动作: drop拦截，trans放行，monitor观察
	// 注意：此字段可能返回 null，表示取不到有效值。
	Action *string `json:"Action,omitempty" name:"Action"`

	// 根据类型匹配：ip(根据ip), area(根据区域)
	// 注意：此字段可能返回 null，表示取不到有效值。
	MatchFrom *string `json:"MatchFrom,omitempty" name:"MatchFrom"`

	// 匹配内容
	// 注意：此字段可能返回 null，表示取不到有效值。
	MatchContent *string `json:"MatchContent,omitempty" name:"MatchContent"`

	// 规则id
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleID *int64 `json:"RuleID,omitempty" name:"RuleID"`

	// 更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type Ipv6Access struct {
	// Ipv6访问功能配置，取值有：
	// <li>on：开启Ipv6访问功能；</li>
	// <li>off：关闭Ipv6访问功能。</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`
}

type L7OfflineLog struct {
	// 日志打包开始时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	LogTime *int64 `json:"LogTime,omitempty" name:"LogTime"`

	// 子域名
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

type ManagedRule struct {
	// 规则id
	RuleId *int64 `json:"RuleId,omitempty" name:"RuleId"`

	// 规则描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// 规则类型名
	RuleTypeName *string `json:"RuleTypeName,omitempty" name:"RuleTypeName"`

	// 策略规则防护等级
	RuleLevelDesc *string `json:"RuleLevelDesc,omitempty" name:"RuleLevelDesc"`

	// 更新时间
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 规则当前状态  block, allow
	Status *string `json:"Status,omitempty" name:"Status"`

	// 规则标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleTags []*string `json:"RuleTags,omitempty" name:"RuleTags"`

	// 规则类型详细描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleTypeDesc *string `json:"RuleTypeDesc,omitempty" name:"RuleTypeDesc"`

	// 规则类型id
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleTypeId *int64 `json:"RuleTypeId,omitempty" name:"RuleTypeId"`
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
type ModifyApplicationProxyRequestParams struct {
	// 站点ID。
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 代理ID。
	ProxyId *string `json:"ProxyId,omitempty" name:"ProxyId"`

	// 当ProxyType=hostname时，表示域名或子域名；
	// 当ProxyType=instance时，表示代理名称。
	ProxyName *string `json:"ProxyName,omitempty" name:"ProxyName"`

	// 参数已经废弃。
	ForwardClientIp *string `json:"ForwardClientIp,omitempty" name:"ForwardClientIp"`

	// 参数已经废弃。
	SessionPersist *bool `json:"SessionPersist,omitempty" name:"SessionPersist"`

	// 会话保持时间，不填写保持原有配置。取值范围：30-3600，单位：秒。
	SessionPersistTime *uint64 `json:"SessionPersistTime,omitempty" name:"SessionPersistTime"`

	// 四层代理模式，取值有：
	// <li>hostname：表示子域名模式；</li>
	// <li>instance：表示实例模式。</li>不填写保持原有配置。
	ProxyType *string `json:"ProxyType,omitempty" name:"ProxyType"`

	// Ipv6访问配置，不填写保持原有配置。
	Ipv6 *Ipv6Access `json:"Ipv6,omitempty" name:"Ipv6"`
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

	// 参数已经废弃。
	ForwardClientIp *string `json:"ForwardClientIp,omitempty" name:"ForwardClientIp"`

	// 参数已经废弃。
	SessionPersist *bool `json:"SessionPersist,omitempty" name:"SessionPersist"`

	// 会话保持时间，不填写保持原有配置。取值范围：30-3600，单位：秒。
	SessionPersistTime *uint64 `json:"SessionPersistTime,omitempty" name:"SessionPersistTime"`

	// 四层代理模式，取值有：
	// <li>hostname：表示子域名模式；</li>
	// <li>instance：表示实例模式。</li>不填写保持原有配置。
	ProxyType *string `json:"ProxyType,omitempty" name:"ProxyType"`

	// Ipv6访问配置，不填写保持原有配置。
	Ipv6 *Ipv6Access `json:"Ipv6,omitempty" name:"Ipv6"`
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
	delete(f, "Ipv6")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyApplicationProxyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyApplicationProxyResponseParams struct {
	// 代理ID。
	ProxyId *string `json:"ProxyId,omitempty" name:"ProxyId"`

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
	OriginType *string `json:"OriginType,omitempty" name:"OriginType"`

	// 源站信息：
	// 当OriginType=custom时，表示一个或多个源站，如：
	// OriginValue=["8.8.8.8:80","9.9.9.9:80"]
	// OriginValue=["test.com:80"]
	// 
	// 当OriginType=origins时，包含一个元素，表示源站组ID，如：
	// OriginValue=["origin-xxx"]
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
	OriginType *string `json:"OriginType,omitempty" name:"OriginType"`

	// 源站信息：
	// 当OriginType=custom时，表示一个或多个源站，如：
	// OriginValue=["8.8.8.8:80","9.9.9.9:80"]
	// OriginValue=["test.com:80"]
	// 
	// 当OriginType=origins时，包含一个元素，表示源站组ID，如：
	// OriginValue=["origin-xxx"]
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

// Predefined struct for user
type ModifyApplicationProxyRuleResponseParams struct {
	// 规则ID
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

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
	// 站点ID
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 代理ID
	ProxyId *string `json:"ProxyId,omitempty" name:"ProxyId"`

	// 规则ID
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// 状态
	// offline: 停用
	// online: 启用
	Status *string `json:"Status,omitempty" name:"Status"`
}

type ModifyApplicationProxyRuleStatusRequest struct {
	*tchttp.BaseRequest
	
	// 站点ID
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 代理ID
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

// Predefined struct for user
type ModifyApplicationProxyRuleStatusResponseParams struct {
	// 规则ID
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

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
	// 站点ID
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 代理ID
	ProxyId *string `json:"ProxyId,omitempty" name:"ProxyId"`

	// 状态
	// offline: 停用
	// online: 启用
	Status *string `json:"Status,omitempty" name:"Status"`
}

type ModifyApplicationProxyStatusRequest struct {
	*tchttp.BaseRequest
	
	// 站点ID
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 代理ID
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

// Predefined struct for user
type ModifyApplicationProxyStatusResponseParams struct {
	// 代理ID
	ProxyId *string `json:"ProxyId,omitempty" name:"ProxyId"`

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
	// 站点id
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 二级域名
	Host *string `json:"Host,omitempty" name:"Host"`

	// 加速开关 on-开启加速；off-关闭加速（AccelerateType：on，SecurityType：on，安全加速，未开防护增强；AccelerateType：off，SecurityType：on，安全加速，开启防护增强；AccelerateType：on，SecurityType：off，内容加速，未开防护增强）
	AccelerateType *string `json:"AccelerateType,omitempty" name:"AccelerateType"`

	// 策略id
	PolicyId *int64 `json:"PolicyId,omitempty" name:"PolicyId"`

	// 安全开关 on-开启安全；off-关闭安全（AccelerateType：on，SecurityType：on，安全加速，未开防护增强；AccelerateType：off，SecurityType：on，安全加速，开启防护增强；AccelerateType：on，SecurityType：off，内容加速，未开防护增强）
	SecurityType *string `json:"SecurityType,omitempty" name:"SecurityType"`
}

type ModifyDDoSPolicyHostRequest struct {
	*tchttp.BaseRequest
	
	// 站点id
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 二级域名
	Host *string `json:"Host,omitempty" name:"Host"`

	// 加速开关 on-开启加速；off-关闭加速（AccelerateType：on，SecurityType：on，安全加速，未开防护增强；AccelerateType：off，SecurityType：on，安全加速，开启防护增强；AccelerateType：on，SecurityType：off，内容加速，未开防护增强）
	AccelerateType *string `json:"AccelerateType,omitempty" name:"AccelerateType"`

	// 策略id
	PolicyId *int64 `json:"PolicyId,omitempty" name:"PolicyId"`

	// 安全开关 on-开启安全；off-关闭安全（AccelerateType：on，SecurityType：on，安全加速，未开防护增强；AccelerateType：off，SecurityType：on，安全加速，开启防护增强；AccelerateType：on，SecurityType：off，内容加速，未开防护增强）
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
	// 修改成功的host
	Host *string `json:"Host,omitempty" name:"Host"`

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
	// 策略组ID
	PolicyId *int64 `json:"PolicyId,omitempty" name:"PolicyId"`

	// 一级域名
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// DDoS具体防护配置
	DdosRule *DdosRule `json:"DdosRule,omitempty" name:"DdosRule"`
}

type ModifyDDoSPolicyRequest struct {
	*tchttp.BaseRequest
	
	// 策略组ID
	PolicyId *int64 `json:"PolicyId,omitempty" name:"PolicyId"`

	// 一级域名
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// DDoS具体防护配置
	DdosRule *DdosRule `json:"DdosRule,omitempty" name:"DdosRule"`
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
	delete(f, "DdosRule")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDDoSPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDDoSPolicyResponseParams struct {
	// 策略组ID
	PolicyId *int64 `json:"PolicyId,omitempty" name:"PolicyId"`

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

// Predefined struct for user
type ModifyDnsRecordResponseParams struct {
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
	// 站点 ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// DNSSEC 状态
	// - enabled 开启
	// - disabled 关闭
	Status *string `json:"Status,omitempty" name:"Status"`
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

// Predefined struct for user
type ModifyDnssecResponseParams struct {
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
	// Zone ID
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 本次变更的域名
	Hosts []*string `json:"Hosts,omitempty" name:"Hosts"`

	// 证书信息, 只需要传入 CertId 即可, 如果为空, 则使用默认证书
	CertInfo []*ServerCertInfo `json:"CertInfo,omitempty" name:"CertInfo"`
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

// Predefined struct for user
type ModifyLoadBalancingResponseParams struct {
	// 负载均衡ID
	LoadBalancingId *string `json:"LoadBalancingId,omitempty" name:"LoadBalancingId"`

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
	// 站点ID
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 负载均衡ID
	LoadBalancingId *string `json:"LoadBalancingId,omitempty" name:"LoadBalancingId"`

	// 状态
	// online: 启用
	// offline: 停用
	Status *string `json:"Status,omitempty" name:"Status"`
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

// Predefined struct for user
type ModifyLoadBalancingStatusResponseParams struct {
	// 负载均衡ID
	LoadBalancingId *string `json:"LoadBalancingId,omitempty" name:"LoadBalancingId"`

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
type ModifyOriginGroupRequestParams struct {
	// 源站组ID
	OriginId *string `json:"OriginId,omitempty" name:"OriginId"`

	// 源站组名称
	OriginName *string `json:"OriginName,omitempty" name:"OriginName"`

	// 配置类型，当OriginType=self 时，需要填写：
	// area: 按区域配置
	// weight: 按权重配置
	// 当OriginType=third_party/cos 时，不需要填写
	Type *string `json:"Type,omitempty" name:"Type"`

	// 源站记录
	Record []*OriginRecord `json:"Record,omitempty" name:"Record"`

	// 站点ID
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 源站类型
	// self：自有源站
	// third_party：第三方源站
	// cos：腾讯云COS源站
	OriginType *string `json:"OriginType,omitempty" name:"OriginType"`
}

type ModifyOriginGroupRequest struct {
	*tchttp.BaseRequest
	
	// 源站组ID
	OriginId *string `json:"OriginId,omitempty" name:"OriginId"`

	// 源站组名称
	OriginName *string `json:"OriginName,omitempty" name:"OriginName"`

	// 配置类型，当OriginType=self 时，需要填写：
	// area: 按区域配置
	// weight: 按权重配置
	// 当OriginType=third_party/cos 时，不需要填写
	Type *string `json:"Type,omitempty" name:"Type"`

	// 源站记录
	Record []*OriginRecord `json:"Record,omitempty" name:"Record"`

	// 站点ID
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 源站类型
	// self：自有源站
	// third_party：第三方源站
	// cos：腾讯云COS源站
	OriginType *string `json:"OriginType,omitempty" name:"OriginType"`
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
	delete(f, "OriginId")
	delete(f, "OriginName")
	delete(f, "Type")
	delete(f, "Record")
	delete(f, "ZoneId")
	delete(f, "OriginType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyOriginGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyOriginGroupResponseParams struct {
	// 源站组ID
	OriginId *string `json:"OriginId,omitempty" name:"OriginId"`

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
type ModifySecurityPolicyRequestParams struct {
	// 一级域名
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 二级域名/应用名
	Entity *string `json:"Entity,omitempty" name:"Entity"`

	// 安全配置
	Config *SecurityConfig `json:"Config,omitempty" name:"Config"`
}

type ModifySecurityPolicyRequest struct {
	*tchttp.BaseRequest
	
	// 一级域名
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 二级域名/应用名
	Entity *string `json:"Entity,omitempty" name:"Entity"`

	// 安全配置
	Config *SecurityConfig `json:"Config,omitempty" name:"Config"`
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
	delete(f, "Config")
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
type ModifyZoneCnameSpeedUpRequestParams struct {
	// 站点 ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// CNAME 加速状态
	// - enabled 开启
	// - disabled 关闭
	Status *string `json:"Status,omitempty" name:"Status"`
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

// Predefined struct for user
type ModifyZoneCnameSpeedUpResponseParams struct {
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
	// 站点 ID，用于唯一标识站点信息
	Id *string `json:"Id,omitempty" name:"Id"`

	// 站点接入方式
	// - full NS 接入
	// - partial CNAME 接入
	Type *string `json:"Type,omitempty" name:"Type"`

	// 自定义站点信息
	VanityNameServers *VanityNameServers `json:"VanityNameServers,omitempty" name:"VanityNameServers"`
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

// Predefined struct for user
type ModifyZoneResponseParams struct {
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
	Cache *CacheConfig `json:"Cache,omitempty" name:"Cache"`

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
	ClientIpHeader *ClientIp `json:"ClientIpHeader,omitempty" name:"ClientIpHeader"`

	// 缓存预刷新配置。
	// 不填写表示保持原有配置。
	CachePrefresh *CachePrefresh `json:"CachePrefresh,omitempty" name:"CachePrefresh"`

	// Ipv6访问配置。
	// 不填写表示保持原有配置。
	Ipv6 *Ipv6Access `json:"Ipv6,omitempty" name:"Ipv6"`
}

type ModifyZoneSettingRequest struct {
	*tchttp.BaseRequest
	
	// 待变更的站点ID。
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 缓存过期时间配置。
	// 不填写表示保持原有配置。
	Cache *CacheConfig `json:"Cache,omitempty" name:"Cache"`

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
	ClientIpHeader *ClientIp `json:"ClientIpHeader,omitempty" name:"ClientIpHeader"`

	// 缓存预刷新配置。
	// 不填写表示保持原有配置。
	CachePrefresh *CachePrefresh `json:"CachePrefresh,omitempty" name:"CachePrefresh"`

	// Ipv6访问配置。
	// 不填写表示保持原有配置。
	Ipv6 *Ipv6Access `json:"Ipv6,omitempty" name:"Ipv6"`
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
	delete(f, "CachePrefresh")
	delete(f, "Ipv6")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyZoneSettingRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyZoneSettingResponseParams struct {
	// 站点ID。
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

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
	// 站点 ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 站点状态
	// - false 开启站点
	// - true 关闭站点
	Paused *bool `json:"Paused,omitempty" name:"Paused"`
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

// Predefined struct for user
type ModifyZoneStatusResponseParams struct {
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

type OfflineCache struct {
	// 离线缓存是否开启，取值有：
	// <li>on：开启；</li>
	// <li>off：关闭。</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Switch *string `json:"Switch,omitempty" name:"Switch"`
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

type OriginCheckOriginStatus struct {
	// healthy: 健康，unhealthy: 不健康，process: 探测中
	Status *string `json:"Status,omitempty" name:"Status"`

	// host列表，源站组不健康时存在值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Host []*string `json:"Host,omitempty" name:"Host"`
}

type OriginFilter struct {
	// 要过滤的字段，支持：name
	Name *string `json:"Name,omitempty" name:"Name"`

	// 要过滤的值
	Value *string `json:"Value,omitempty" name:"Value"`
}

type OriginGroup struct {
	// 源站组ID
	OriginId *string `json:"OriginId,omitempty" name:"OriginId"`

	// 源站组名称
	OriginName *string `json:"OriginName,omitempty" name:"OriginName"`

	// 源站组配置类型
	// area：表示按照Record记录中的Area字段进行按客户端IP所在区域回源。
	// weight：表示按照Record记录中的Weight字段进行按权重回源。
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

	// 当前源站组是否被四层代理使用。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationProxyUsed *bool `json:"ApplicationProxyUsed,omitempty" name:"ApplicationProxyUsed"`

	// 当前源站组是否被负载均衡使用。
	// 注意：此字段可能返回 null，表示取不到有效值。
	LoadBalancingUsed *bool `json:"LoadBalancingUsed,omitempty" name:"LoadBalancingUsed"`

	// 源站状态信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *OriginCheckOriginStatus `json:"Status,omitempty" name:"Status"`

	// 使用当前源站组的负载均衡的类型：
	// none：未被使用
	// dns_only：被仅DNS类型负载均衡使用
	// proxied：被代理加速类型负载均衡使用
	// both：同时被仅DNS和代理加速类型负载均衡使用
	// 注意：此字段可能返回 null，表示取不到有效值。
	LoadBalancingUsedType *string `json:"LoadBalancingUsedType,omitempty" name:"LoadBalancingUsedType"`
}

type OriginRecord struct {
	// 记录值
	Record *string `json:"Record,omitempty" name:"Record"`

	// 当源站配置类型Type=area时，表示区域
	// 为空表示默认区域
	Area []*string `json:"Area,omitempty" name:"Area"`

	// 当源站配置类型Type=weight时，表示权重
	// 取值范围为[1-100]
	// 源站组内多个源站权重总和应为100。
	// 当源站配置类型Type=proto，表示权重
	// 取值范围为[1-100]
	// 源站组内Proto相同的多个源站权重总和应为100。
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

	// 当源站配置类型Type=proto时，表示客户端请求协议，取值：http/https
	// 注意：此字段可能返回 null，表示取不到有效值。
	Proto *string `json:"Proto,omitempty" name:"Proto"`
}

type OriginRecordPrivateParameter struct {
	// 私有鉴权参数名称：
	// "AccessKeyId"：Access Key ID
	// "SecretAccessKey"：Secret Access Key
	Name *string `json:"Name,omitempty" name:"Name"`

	// 私有鉴权参数数值
	Value *string `json:"Value,omitempty" name:"Value"`
}

type PortraitManagedRuleDetail struct {
	// 规则唯一id
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleId *int64 `json:"RuleId,omitempty" name:"RuleId"`

	// 规则的描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 规则所属类型的名字, botdb(用户画像)
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleTypeName *string `json:"RuleTypeName,omitempty" name:"RuleTypeName"`

	// 规则内的功能分类Id(扫描器，Bot行为等)
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClassificationId *int64 `json:"ClassificationId,omitempty" name:"ClassificationId"`

	// 规则当前所属动作状态(block, alg, ...)
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

type QueryCondition struct {
	// 维度
	Key *string `json:"Key,omitempty" name:"Key"`

	// 操作符
	Operator *string `json:"Operator,omitempty" name:"Operator"`

	// 维度值
	Value []*string `json:"Value,omitempty" name:"Value"`
}

type QueryString struct {
	// CacheKey是否由QueryString组成，取值有：
	// <li>on：是；</li>
	// <li>off：否。</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
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

type RateLimitConfig struct {
	// 开关
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// 用户规则
	UserRules []*RateLimitUserRule `json:"UserRules,omitempty" name:"UserRules"`

	// 默认模板
	// 注意：此字段可能返回 null，表示取不到有效值。
	Template *RateLimitTemplate `json:"Template,omitempty" name:"Template"`

	// 智能客户端过滤
	// 注意：此字段可能返回 null，表示取不到有效值。
	Intelligence *RateLimitIntelligence `json:"Intelligence,omitempty" name:"Intelligence"`
}

type RateLimitIntelligence struct {
	// 功能开关
	// 注意：此字段可能返回 null，表示取不到有效值。
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// 执行动作 monitor(观察), alg(挑战)
	// 注意：此字段可能返回 null，表示取不到有效值。
	Action *string `json:"Action,omitempty" name:"Action"`
}

type RateLimitTemplate struct {
	// 模板名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Mode *string `json:"Mode,omitempty" name:"Mode"`

	// 模板值详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	Detail *RateLimitTemplateDetail `json:"Detail,omitempty" name:"Detail"`
}

type RateLimitTemplateDetail struct {
	// 模板名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Mode *string `json:"Mode,omitempty" name:"Mode"`

	// 唯一id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ID *int64 `json:"ID,omitempty" name:"ID"`

	// 处置动作
	// 注意：此字段可能返回 null，表示取不到有效值。
	Action *string `json:"Action,omitempty" name:"Action"`

	// 惩罚时间，秒
	// 注意：此字段可能返回 null，表示取不到有效值。
	PunishTime *int64 `json:"PunishTime,omitempty" name:"PunishTime"`

	// 阈值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Threshold *int64 `json:"Threshold,omitempty" name:"Threshold"`

	// 统计周期
	// 注意：此字段可能返回 null，表示取不到有效值。
	Period *int64 `json:"Period,omitempty" name:"Period"`
}

type RateLimitUserRule struct {
	// RateLimit统计阈值
	Threshold *int64 `json:"Threshold,omitempty" name:"Threshold"`

	// RateLimit统计时间
	Period *int64 `json:"Period,omitempty" name:"Period"`

	// 规则名
	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`

	// 动作：monitor(观察), drop(拦截)
	Action *string `json:"Action,omitempty" name:"Action"`

	// 惩罚时长
	PunishTime *int64 `json:"PunishTime,omitempty" name:"PunishTime"`

	// 处罚时长单位，second
	PunishTimeUnit *string `json:"PunishTimeUnit,omitempty" name:"PunishTimeUnit"`

	// 规则状态
	RuleStatus *string `json:"RuleStatus,omitempty" name:"RuleStatus"`

	// 规则
	Conditions []*ACLCondition `json:"Conditions,omitempty" name:"Conditions"`

	// 规则权重
	RulePriority *int64 `json:"RulePriority,omitempty" name:"RulePriority"`

	// 规则id
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleID *int64 `json:"RuleID,omitempty" name:"RuleID"`

	// 过滤词
	// 注意：此字段可能返回 null，表示取不到有效值。
	FreqFields []*string `json:"FreqFields,omitempty" name:"FreqFields"`

	// 更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

// Predefined struct for user
type ReclaimZoneRequestParams struct {
	// 站点名称
	Name *string `json:"Name,omitempty" name:"Name"`
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

// Predefined struct for user
type ReclaimZoneResponseParams struct {
	// 站点名称
	Name *string `json:"Name,omitempty" name:"Name"`

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
	// 资源 ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 付费模式
	// 0 为后付费
	// 1 为预付费
	PayMode *int64 `json:"PayMode,omitempty" name:"PayMode"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 生效时间
	EnableTime *string `json:"EnableTime,omitempty" name:"EnableTime"`

	// 失效时间
	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// 套餐状态
	Status *string `json:"Status,omitempty" name:"Status"`

	// 询价参数
	Sv []*Sv `json:"Sv,omitempty" name:"Sv"`

	// 是否自动续费
	// 0 表示默认状态
	// 1 表示自动续费
	// 2 表示不自动续费
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitempty" name:"AutoRenewFlag"`

	// 套餐关联资源ID
	PlanId *string `json:"PlanId,omitempty" name:"PlanId"`
}

// Predefined struct for user
type ScanDnsRecordsRequestParams struct {
	// 站点 ID
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`
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

// Predefined struct for user
type ScanDnsRecordsResponseParams struct {
	// 扫描状态
	// - doing 扫描中
	// - done 扫描完成
	Status *string `json:"Status,omitempty" name:"Status"`

	// 扫描后添加的记录数
	RecordsAdded *int64 `json:"RecordsAdded,omitempty" name:"RecordsAdded"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ScanDnsRecordsResponse struct {
	*tchttp.BaseResponse
	Response *ScanDnsRecordsResponseParams `json:"Response"`
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

type SecEntry struct {
	// 查询维度值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Key *string `json:"Key,omitempty" name:"Key"`

	// 查询维度下详细数据。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value []*SecEntryValue `json:"Value,omitempty" name:"Value"`
}

type SecEntryValue struct {
	// 指标名称。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Metric *string `json:"Metric,omitempty" name:"Metric"`

	// 时序数据详情。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Detail []*TimingDataItem `json:"Detail,omitempty" name:"Detail"`

	// 最大值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Max *int64 `json:"Max,omitempty" name:"Max"`

	// 平均值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Avg *float64 `json:"Avg,omitempty" name:"Avg"`

	// 数据总和。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Sum *float64 `json:"Sum,omitempty" name:"Sum"`
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
	// 门神配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	WafConfig *WafConfig `json:"WafConfig,omitempty" name:"WafConfig"`

	// RateLimit配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	RateLimitConfig *RateLimitConfig `json:"RateLimitConfig,omitempty" name:"RateLimitConfig"`

	// DDoS配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	DdosConfig *DDoSConfig `json:"DdosConfig,omitempty" name:"DdosConfig"`

	// ACL配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	AclConfig *AclConfig `json:"AclConfig,omitempty" name:"AclConfig"`

	// Bot配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	BotConfig *BotConfig `json:"BotConfig,omitempty" name:"BotConfig"`

	// 总开关
	// 注意：此字段可能返回 null，表示取不到有效值。
	SwitchConfig *SwitchConfig `json:"SwitchConfig,omitempty" name:"SwitchConfig"`

	// IP黑白名单
	// 注意：此字段可能返回 null，表示取不到有效值。
	IpTableConfig *IpTableConfig `json:"IpTableConfig,omitempty" name:"IpTableConfig"`
}

type SecurityEntity struct {
	// 用户appid
	AppId *int64 `json:"AppId,omitempty" name:"AppId"`

	// 一级域名
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 二级域名
	Entity *string `json:"Entity,omitempty" name:"Entity"`

	// 类型 domain/application
	EntityType *string `json:"EntityType,omitempty" name:"EntityType"`
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

type ShieldArea struct {
	// 一级域名id
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 策略id
	PolicyId *int64 `json:"PolicyId,omitempty" name:"PolicyId"`

	// 防护类型 domain/application
	Type *string `json:"Type,omitempty" name:"Type"`

	// 四层应用名
	// 注意：此字段可能返回 null，表示取不到有效值。
	EntityName *string `json:"EntityName,omitempty" name:"EntityName"`

	// 7层域名参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Application []*DDoSApplication `json:"Application,omitempty" name:"Application"`

	// 四层tcp转发规则数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TcpNum *int64 `json:"TcpNum,omitempty" name:"TcpNum"`

	// 四层udp转发规则数
	// 注意：此字段可能返回 null，表示取不到有效值。
	UdpNum *int64 `json:"UdpNum,omitempty" name:"UdpNum"`

	// 实例名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Entity *string `json:"Entity,omitempty" name:"Entity"`

	// 是否为共享资源客户，注意共享资源用户不可以切换代理模式，true-是；false-否
	// 注意：此字段可能返回 null，表示取不到有效值。
	Share *bool `json:"Share,omitempty" name:"Share"`
}

type SmartRouting struct {
	// 智能加速配置开关，取值有：
	// <li>on：开启；</li>
	// <li>off：关闭。</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`
}

type Sv struct {
	// 询价参数 key
	Key *string `json:"Key,omitempty" name:"Key"`

	// 询价参数 value
	Value *string `json:"Value,omitempty" name:"Value"`
}

type SwitchConfig struct {
	// Web类型的安全总开关：Web基础防护，自定义规则，速率限制
	WebSwitch *string `json:"WebSwitch,omitempty" name:"WebSwitch"`
}

type Tag struct {
	// 标签键
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// 标签值
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`
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

type TimingDataItem struct {
	// 返回数据对应时间点，采用unix秒级时间戳
	// 注意：此字段可能返回 null，表示取不到有效值。
	Timestamp *int64 `json:"Timestamp,omitempty" name:"Timestamp"`

	// 具体数值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *int64 `json:"Value,omitempty" name:"Value"`
}

type TimingDataRecord struct {
	// 查询维度值
	TypeKey *string `json:"TypeKey,omitempty" name:"TypeKey"`

	// 详细时序数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	TypeValue []*TimingTypeValue `json:"TypeValue,omitempty" name:"TypeValue"`
}

type TimingTypeValue struct {
	// 数据和
	// 注意：此字段可能返回 null，表示取不到有效值。
	Sum *int64 `json:"Sum,omitempty" name:"Sum"`

	// 最大
	// 注意：此字段可能返回 null，表示取不到有效值。
	Max *int64 `json:"Max,omitempty" name:"Max"`

	// 平均
	// 注意：此字段可能返回 null，表示取不到有效值。
	Avg *int64 `json:"Avg,omitempty" name:"Avg"`

	// 指标名
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`

	// 废弃字段，即将下线，请使用Detail字段
	// 注意：此字段可能返回 null，表示取不到有效值。
	DetailData []*int64 `json:"DetailData,omitempty" name:"DetailData"`

	// 详细数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	Detail []*TimingDataItem `json:"Detail,omitempty" name:"Detail"`
}

type TopDataRecord struct {
	// 查询维度值
	TypeKey *string `json:"TypeKey,omitempty" name:"TypeKey"`

	// top数据排行
	// 注意：此字段可能返回 null，表示取不到有效值。
	DetailData []*TopDetailData `json:"DetailData,omitempty" name:"DetailData"`
}

type TopDetailData struct {
	// 字段名
	Key *string `json:"Key,omitempty" name:"Key"`

	// 字段值
	Value *int64 `json:"Value,omitempty" name:"Value"`
}

type TopNEntry struct {
	// top查询维度值。
	Key *string `json:"Key,omitempty" name:"Key"`

	// 查询具体数据。
	Value []*TopNEntryValue `json:"Value,omitempty" name:"Value"`
}

type TopNEntryValue struct {
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

type WafConfig struct {
	// 开关
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// 防护级别，loose/normal/strict/stricter/custom
	Level *string `json:"Level,omitempty" name:"Level"`

	// 模式 block-阻断；observe-观察模式；close-关闭
	Mode *string `json:"Mode,omitempty" name:"Mode"`

	// 托管规则黑白名单
	WafRules *WafRule `json:"WafRules,omitempty" name:"WafRules"`

	// AI规则引擎防护
	// 注意：此字段可能返回 null，表示取不到有效值。
	AiRule *AiRule `json:"AiRule,omitempty" name:"AiRule"`
}

type WafRule struct {
	// 黑名单，ID参考接口 DescribeSecurityPolicyManagedRules
	BlockRuleIDs []*int64 `json:"BlockRuleIDs,omitempty" name:"BlockRuleIDs"`

	// 托管规则 开关
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// 观察模式，ID参考接口 DescribeSecurityPolicyManagedRules
	// 注意：此字段可能返回 null，表示取不到有效值。
	ObserveRuleIDs []*int64 `json:"ObserveRuleIDs,omitempty" name:"ObserveRuleIDs"`
}

type WebAttackEvent struct {
	// 客户端ip
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClientIp *string `json:"ClientIp,omitempty" name:"ClientIp"`

	// 攻击URL
	// 注意：此字段可能返回 null，表示取不到有效值。
	AttackUrl *string `json:"AttackUrl,omitempty" name:"AttackUrl"`

	// 攻击时间 单位为s
	// 注意：此字段可能返回 null，表示取不到有效值。
	AttackTime *int64 `json:"AttackTime,omitempty" name:"AttackTime"`
}

type WebEventData struct {
	// 攻击事件数据集合
	// 注意：此字段可能返回 null，表示取不到有效值。
	List []*WebAttackEvent `json:"List,omitempty" name:"List"`

	// 当前页
	// 注意：此字段可能返回 null，表示取不到有效值。
	PageNo *int64 `json:"PageNo,omitempty" name:"PageNo"`

	// 每页展示条数
	// 注意：此字段可能返回 null，表示取不到有效值。
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`

	// 总页数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Pages *int64 `json:"Pages,omitempty" name:"Pages"`

	// 总条数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalSize *int64 `json:"TotalSize,omitempty" name:"TotalSize"`
}

type WebLogData struct {
	// 分组数据。
	// 注意：此字段可能返回 null，表示取不到有效值。
	List []*WebLogs `json:"List,omitempty" name:"List"`

	// 分页拉取的起始页号。最小值：1。
	PageNo *int64 `json:"PageNo,omitempty" name:"PageNo"`

	// 分页拉取的最大返回结果数。最大值：1000。
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`

	// 总页数。
	Pages *int64 `json:"Pages,omitempty" name:"Pages"`

	// 总条数。
	TotalSize *int64 `json:"TotalSize,omitempty" name:"TotalSize"`
}

type WebLogs struct {
	// 该字段已废弃。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AttackContent *string `json:"AttackContent,omitempty" name:"AttackContent"`

	// 攻击源（客户端）Ip。
	AttackIp *string `json:"AttackIp,omitempty" name:"AttackIp"`

	// 该字段已废弃。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AttackType *string `json:"AttackType,omitempty" name:"AttackType"`

	// 受攻击子域名。
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 该字段已废弃。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Msuuid *string `json:"Msuuid,omitempty" name:"Msuuid"`

	// 该字段已废弃。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RequestMethod *string `json:"RequestMethod,omitempty" name:"RequestMethod"`

	// URI
	RequestUri *string `json:"RequestUri,omitempty" name:"RequestUri"`

	// 该字段已废弃。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RiskLevel *string `json:"RiskLevel,omitempty" name:"RiskLevel"`

	// 该字段已废弃。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleId *uint64 `json:"RuleId,omitempty" name:"RuleId"`

	// IP所在国家iso-3166中alpha-2编码，编码信息请参考[ISO-3166](https://git.woa.com/edgeone/iso-3166/blob/master/all/all.json)
	SipCountryCode *string `json:"SipCountryCode,omitempty" name:"SipCountryCode"`

	// 请求（事件）ID。
	EventId *string `json:"EventId,omitempty" name:"EventId"`

	// 该字段已废弃。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DisposalMethod *string `json:"DisposalMethod,omitempty" name:"DisposalMethod"`

	// http log。
	HttpLog *string `json:"HttpLog,omitempty" name:"HttpLog"`

	// 该字段已废弃。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Ua *string `json:"Ua,omitempty" name:"Ua"`

	// 攻击时间，采用unix秒级时间戳。
	AttackTime *uint64 `json:"AttackTime,omitempty" name:"AttackTime"`

	// 规则相关信息列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleDetailList []*SecRuleRelatedInfo `json:"RuleDetailList,omitempty" name:"RuleDetailList"`
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

	// 资源标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`

	// 计费资源
	// 注意：此字段可能返回 null，表示取不到有效值。
	Resources []*Resource `json:"Resources,omitempty" name:"Resources"`

	// 是否开启cname加速
	// - enabled 开启
	// - disabled 关闭
	// 注意：此字段可能返回 null，表示取不到有效值。
	CnameSpeedUp *string `json:"CnameSpeedUp,omitempty" name:"CnameSpeedUp"`
}

type ZoneFilter struct {
	// 过滤字段名，支持的列表如下：
	// - name: 站点名。
	// - status: 站点状态
	// - tagKey: 标签键
	// - tagValue: 标签值
	Name *string `json:"Name,omitempty" name:"Name"`

	// 过滤字段值
	Values []*string `json:"Values,omitempty" name:"Values"`

	// 是否启用模糊查询，仅支持过滤字段名为name。模糊查询时，Values长度最大为1
	Fuzzy *bool `json:"Fuzzy,omitempty" name:"Fuzzy"`
}