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
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

type AccelerateMainland struct {
	// 是否开启中国大陆加速优化配置，取值有：
	// <li>on：开启；</li>
	// <li>off：关闭。</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`
}

type AccelerateType struct {
	// 加速开关。取值范围：
	// <li> on：打开;</li>
	// <li>off：关闭。</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`
}

type AccelerationDomain struct {
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 加速域名名称。
	DomainName *string `json:"DomainName,omitnil,omitempty" name:"DomainName"`

	// 加速域名状态，取值有：
	// <li>online：已生效；</li>
	// <li>process：部署中；</li>
	// <li>offline：已停用；</li>
	// <li>forbidden：已封禁；</li>
	// <li>init：未生效，待激活站点；</li>
	DomainStatus *string `json:"DomainStatus,omitnil,omitempty" name:"DomainStatus"`

	// 源站信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OriginDetail *OriginDetail `json:"OriginDetail,omitnil,omitempty" name:"OriginDetail"`

	// 回源协议，取值有：
	// <li>FOLLOW: 协议跟随；</li>
	// <li>HTTP: HTTP协议回源；</li>
	// <li>HTTPS: HTTPS协议回源。</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	OriginProtocol *string `json:"OriginProtocol,omitnil,omitempty" name:"OriginProtocol"`

	// 域名证书信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Certificate *AccelerationDomainCertificate `json:"Certificate,omitnil,omitempty" name:"Certificate"`

	// HTTP回源端口。
	// 注意：此字段可能返回 null，表示取不到有效值。
	HttpOriginPort *uint64 `json:"HttpOriginPort,omitnil,omitempty" name:"HttpOriginPort"`

	// HTTPS回源端口。
	// 注意：此字段可能返回 null，表示取不到有效值。
	HttpsOriginPort *uint64 `json:"HttpsOriginPort,omitnil,omitempty" name:"HttpsOriginPort"`

	// IPv6状态，取值有：
	// <li>follow：遵循站点IPv6配置；</li>
	// <li>on：开启状态；</li>
	// <li>off：关闭状态。</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	IPv6Status *string `json:"IPv6Status,omitnil,omitempty" name:"IPv6Status"`

	// CNAME 地址。
	Cname *string `json:"Cname,omitnil,omitempty" name:"Cname"`

	// 加速域名归属权验证状态，取值有： <li>pending：待验证；</li> <li>finished：已完成验证。</li>	
	// 注意：此字段可能返回 null，表示取不到有效值。
	IdentificationStatus *string `json:"IdentificationStatus,omitnil,omitempty" name:"IdentificationStatus"`

	// 创建时间。
	CreatedOn *string `json:"CreatedOn,omitnil,omitempty" name:"CreatedOn"`

	// 修改时间。
	ModifiedOn *string `json:"ModifiedOn,omitnil,omitempty" name:"ModifiedOn"`

	// 当域名需要进行归属权验证才能继续提供服务时，该对象会携带对应验证方式所需要的信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OwnershipVerification *OwnershipVerification `json:"OwnershipVerification,omitnil,omitempty" name:"OwnershipVerification"`
}

type AccelerationDomainCertificate struct {
	// 配置证书的模式，取值有： <li>disable：不配置证书；</li> <li>eofreecert：配置 EdgeOne 免费证书；</li> <li>sslcert：配置 SSL 证书。</li>
	Mode *string `json:"Mode,omitnil,omitempty" name:"Mode"`

	// 服务端证书列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	List []*CertificateInfo `json:"List,omitnil,omitempty" name:"List"`

	// 边缘双向认证配置。
	ClientCertInfo *MutualTLS `json:"ClientCertInfo,omitnil,omitempty" name:"ClientCertInfo"`
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
	// <li>app_proto：应用层协议；</li>
	// <li>sip_proto：网络层协议；</li>
	// <li>uabot：UA 特征规则，仅bot自定义规则可用；</li>
	// <li>idcid：IDC 规则，仅bot自定义规则可用；</li>
	// <li>sipbot：搜索引擎规则，仅bot自定义规则可用；</li>
	// <li>portrait：画像分析，仅bot自定义规则可用；</li>
	// <li>header_seq：请求头顺序，仅bot自定义规则可用；</li>
	// <li>hdr：请求正文，仅Web防护自定义规则可用。</li>
	MatchFrom *string `json:"MatchFrom,omitnil,omitempty" name:"MatchFrom"`

	// 匹配字符串。当 MatchFrom 为 header 时，可以填入 header 的 key 作为参数。
	MatchParam *string `json:"MatchParam,omitnil,omitempty" name:"MatchParam"`

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
	Operator *string `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 匹配内容。
	MatchContent *string `json:"MatchContent,omitnil,omitempty" name:"MatchContent"`
}

type AclConfig struct {
	// 开关，取值有：
	// <li> on：开启；</li>
	// <li> off：关闭。</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// 用户自定义规则。
	AclUserRules []*AclUserRule `json:"AclUserRules,omitnil,omitempty" name:"AclUserRules"`

	// 托管定制规则
	// 注意：此字段可能返回 null，表示取不到有效值。
	Customizes []*AclUserRule `json:"Customizes,omitnil,omitempty" name:"Customizes"`
}

type AclUserRule struct {
	// 规则名。
	RuleName *string `json:"RuleName,omitnil,omitempty" name:"RuleName"`

	// 处罚动作，取值有：
	// <li>trans：放行；</li>
	// <li>drop：拦截；</li>
	// <li>monitor：观察；</li>
	// <li>ban：IP 封禁；</li>
	// <li>redirect：重定向；</li>
	// <li>page：指定页面；</li>
	// <li>alg：JavaScript 挑战。</li>
	Action *string `json:"Action,omitnil,omitempty" name:"Action"`

	// 规则状态，取值有：
	// <li>on：生效；</li>
	// <li>off：失效。</li>
	RuleStatus *string `json:"RuleStatus,omitnil,omitempty" name:"RuleStatus"`

	// 自定义规则。
	AclConditions []*AclCondition `json:"AclConditions,omitnil,omitempty" name:"AclConditions"`

	// 规则优先级，取值范围0-100。
	RulePriority *int64 `json:"RulePriority,omitnil,omitempty" name:"RulePriority"`

	// 规则 Id。仅出参使用。
	RuleID *int64 `json:"RuleID,omitnil,omitempty" name:"RuleID"`

	// 更新时间。仅出参使用。
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// ip 封禁的惩罚时间。Action 是 ban 时必填，且不能为空，取值范围0-2天。
	PunishTime *int64 `json:"PunishTime,omitnil,omitempty" name:"PunishTime"`

	// ip 封禁的惩罚时间单位，取值有：
	// <li>second：秒；</li>
	// <li>minutes：分；</li>
	// <li>hour：小时。</li>默认为 second。
	PunishTimeUnit *string `json:"PunishTimeUnit,omitnil,omitempty" name:"PunishTimeUnit"`

	// 自定义返回页面的名称。Action 是 page 时必填，且不能为空。	
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 自定义返回页面的实例 Id。默认为0，代表使用系统默认拦截页面。该参数已废弃。
	PageId *int64 `json:"PageId,omitnil,omitempty" name:"PageId"`

	// 自定义响应 Id。该 Id 可通过查询自定义错误页列表接口获取。默认值为default，使用系统默认页面。Action 是 page 时必填，且不能为空。	
	CustomResponseId *string `json:"CustomResponseId,omitnil,omitempty" name:"CustomResponseId"`

	// 自定义返回页面的响应码。Action 是 page 时必填，且不能为空，取值: 100~600，不支持 3xx 响应码。默认值：567。
	ResponseCode *int64 `json:"ResponseCode,omitnil,omitempty" name:"ResponseCode"`

	// 重定向时候的地址。Action 是 redirect 时必填，且不能为空。	
	RedirectUrl *string `json:"RedirectUrl,omitnil,omitempty" name:"RedirectUrl"`
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
	// <li> SslTlsSecureConf；</li>
	// <li> OcspStapling；</li>
	// <li> HTTP/2 访问（Http2）；</li>
	// <li> 回源跟随重定向(UpstreamFollowRedirect)；</li>
	// <li> 修改源站(Origin)。</li>
	// <li> 七层回源超时(HTTPUpstreamTimeout)。</li>
	// <li> Http应答（HttpResponse）。</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	NormalAction *NormalAction `json:"NormalAction,omitnil,omitempty" name:"NormalAction"`

	// 带有请求头/响应头的功能操作，选择该类型的功能项有：
	// <li> 修改 HTTP 请求头（RequestHeader）；</li>
	// <li> 修改HTTP响应头（ResponseHeader）。</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	RewriteAction *RewriteAction `json:"RewriteAction,omitnil,omitempty" name:"RewriteAction"`

	// 带有状态码的功能操作，选择该类型的功能项有：
	// <li> 自定义错误页面（ErrorPage）；</li>
	// <li> 状态码缓存 TTL（StatusCodeCache）。</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	CodeAction *CodeAction `json:"CodeAction,omitnil,omitempty" name:"CodeAction"`
}

type AdvancedFilter struct {
	// 需要过滤的字段。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 字段的过滤值。
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`

	// 是否启用模糊查询。
	Fuzzy *bool `json:"Fuzzy,omitnil,omitempty" name:"Fuzzy"`
}

type AiRule struct {
	// AI规则引擎状态，取值有：
	// <li> smart_status_close：关闭；</li>
	// <li> smart_status_open：拦截处置；</li>
	// <li> smart_status_observe：观察处置。</li>
	Mode *string `json:"Mode,omitnil,omitempty" name:"Mode"`
}

type AlgDetectJS struct {
	// 操作名称。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 工作量证明 (proof_Of-Work)校验强度，默认low，取值有：
	// <li>low：低；</li>
	// <li>middle：中；</li>
	// <li>high：高。</li>
	WorkLevel *string `json:"WorkLevel,omitnil,omitempty" name:"WorkLevel"`

	// 执行方式，js延迟执行的时间。单位为ms，默认500，取值：0～1000。
	ExecuteMode *int64 `json:"ExecuteMode,omitnil,omitempty" name:"ExecuteMode"`

	// 客户端末启用JS（末完成检测）统计周期。单位为秒，默认10，取值：5～3600。
	InvalidStatTime *int64 `json:"InvalidStatTime,omitnil,omitempty" name:"InvalidStatTime"`

	// 客户端末启用JS（末完成检测）触发阈值。单位为次，默认300，取值：1～100000000。
	InvalidThreshold *int64 `json:"InvalidThreshold,omitnil,omitempty" name:"InvalidThreshold"`

	// Bot主动特征识别客户端行为校验结果。
	AlgDetectResults []*AlgDetectResult `json:"AlgDetectResults,omitnil,omitempty" name:"AlgDetectResults"`
}

type AlgDetectResult struct {
	// 校验结果，取值有：
	// <li>invalid：不合法Cookie；</li>
	// <li>cookie_empty：末携带Cookie或Cookie己过期；</li>
	// <li>js_empty：客户端末启用JS（末完成检测）；</li>
	// <li>low：会话速率和周期特征校验低风险；</li>
	// <li>middle：会话速率和周期特征校验中风险；</li>
	// <li>high：会话速率和周期特征校验高风险；</li>
	// <li>timeout：检测超时时长；</li>
	// <li>not_browser：不合法浏览器；</li>
	// <li>is_bot：Bot客户端。</li>
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// 处罚动作，取值有：
	// <li>drop：拦截；</li>
	// <li>monitor：观察；</li>
	// <li>silence：静默；</li>
	// <li>shortdelay：（短时间）等待后响应；</li>
	// <li>longdelay：（长时间）等待后响应。</li>
	Action *string `json:"Action,omitnil,omitempty" name:"Action"`
}

type AlgDetectRule struct {
	// 规则id。
	RuleID *int64 `json:"RuleID,omitnil,omitempty" name:"RuleID"`

	// 规则名。
	RuleName *string `json:"RuleName,omitnil,omitempty" name:"RuleName"`

	// 规则开关。
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// 自定义规则。
	AlgConditions []*AclCondition `json:"AlgConditions,omitnil,omitempty" name:"AlgConditions"`

	// Cookie校验和会话行为分析。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AlgDetectSession *AlgDetectSession `json:"AlgDetectSession,omitnil,omitempty" name:"AlgDetectSession"`

	// 客户端行为校验。
	AlgDetectJS []*AlgDetectJS `json:"AlgDetectJS,omitnil,omitempty" name:"AlgDetectJS"`

	// 更新时间。仅出参使用。
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`
}

type AlgDetectSession struct {
	// 操作名称。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 校验方式，默认update_detect，取值有：
	// <li>detect：仅校验；</li>
	// <li>update_detect：更新Cookie并校验。</li>
	DetectMode *string `json:"DetectMode,omitnil,omitempty" name:"DetectMode"`

	// 会话速率和周期特征校验开关，默认off，取值有：
	// <li>off：关闭；</li>
	// <li>on：打开。</li>
	SessionAnalyzeSwitch *string `json:"SessionAnalyzeSwitch,omitnil,omitempty" name:"SessionAnalyzeSwitch"`

	// 校验结果为未携带Cookie或Cookie已过期的统计周期。单位为秒，默认10，取值：5～3600。
	InvalidStatTime *int64 `json:"InvalidStatTime,omitnil,omitempty" name:"InvalidStatTime"`

	// 校验结果为未携带Cookie或Cookie已过期的触发阈值。单位为次，默认300，取值：1～100000000。
	InvalidThreshold *int64 `json:"InvalidThreshold,omitnil,omitempty" name:"InvalidThreshold"`

	// Cookie校验校验结果。
	AlgDetectResults []*AlgDetectResult `json:"AlgDetectResults,omitnil,omitempty" name:"AlgDetectResults"`

	// 会话速率和周期特征校验结果。
	SessionBehaviors []*AlgDetectResult `json:"SessionBehaviors,omitnil,omitempty" name:"SessionBehaviors"`
}

type AliasDomain struct {
	// 别称域名名称。
	AliasName *string `json:"AliasName,omitnil,omitempty" name:"AliasName"`

	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 目标域名名称。
	TargetName *string `json:"TargetName,omitnil,omitempty" name:"TargetName"`

	// 别称域名状态，取值有：
	// <li> active：已生效； </li>
	// <li> pending：部署中；</li>
	// <li> conflict：被找回。 </li>
	// <li> stop：已停用；</li>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 封禁模式，取值有：
	// <li> 0：未封禁； </li>
	// <li> 11：合规封禁；</li>
	// <li> 14：未备案封禁。</li>
	ForbidMode *int64 `json:"ForbidMode,omitnil,omitempty" name:"ForbidMode"`

	// 别称域名创建时间。
	CreatedOn *string `json:"CreatedOn,omitnil,omitempty" name:"CreatedOn"`

	// 别称域名修改时间。
	ModifiedOn *string `json:"ModifiedOn,omitnil,omitempty" name:"ModifiedOn"`
}

type ApplicationProxy struct {
	// 站点ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 站点名称。
	ZoneName *string `json:"ZoneName,omitnil,omitempty" name:"ZoneName"`

	// 代理ID。
	ProxyId *string `json:"ProxyId,omitnil,omitempty" name:"ProxyId"`

	// 当ProxyType=hostname时，表示域名或子域名；
	// 当ProxyType=instance时，表示代理名称。
	ProxyName *string `json:"ProxyName,omitnil,omitempty" name:"ProxyName"`

	// 四层代理模式，取值有：
	// <li>hostname：表示子域名模式；</li>
	// <li>instance：表示实例模式。</li>
	ProxyType *string `json:"ProxyType,omitnil,omitempty" name:"ProxyType"`

	// 调度模式，取值有：
	// <li>ip：表示Anycast IP调度；</li>
	// <li>domain：表示CNAME调度。</li>
	PlatType *string `json:"PlatType,omitnil,omitempty" name:"PlatType"`

	// 加速区域，取值有：
	// <li>mainland：中国大陆境内;</li>
	// <li>overseas：全球（不含中国大陆）。</li>
	// 默认值：overseas
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 是否开启安全，取值有：
	// <li>0：关闭安全；</li>
	// <li>1：开启安全。</li>
	SecurityType *int64 `json:"SecurityType,omitnil,omitempty" name:"SecurityType"`

	// 是否开启加速，取值有：
	// <li>0：关闭加速；</li>
	// <li>1：开启加速。</li>
	AccelerateType *int64 `json:"AccelerateType,omitnil,omitempty" name:"AccelerateType"`

	// 会话保持时间。
	SessionPersistTime *uint64 `json:"SessionPersistTime,omitnil,omitempty" name:"SessionPersistTime"`

	// 状态，取值有：
	// <li>online：启用；</li>
	// <li>offline：停用；</li>
	// <li>progress：部署中；</li>
	// <li>stopping：停用中；</li>
	// <li>fail：部署失败/停用失败。</li>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 封禁状态，取值有：
	// <li>banned：已封禁;</li>
	// <li>banning：封禁中；</li>
	// <li>recover：已解封；</li>
	// <li>recovering：解封禁中。</li>
	BanStatus *string `json:"BanStatus,omitnil,omitempty" name:"BanStatus"`

	// 调度信息。
	ScheduleValue []*string `json:"ScheduleValue,omitnil,omitempty" name:"ScheduleValue"`

	// 当ProxyType=hostname时：
	// 表示代理加速唯一标识。
	HostId *string `json:"HostId,omitnil,omitempty" name:"HostId"`

	// Ipv6访问配置。
	Ipv6 *Ipv6 `json:"Ipv6,omitnil,omitempty" name:"Ipv6"`

	// 更新时间。
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 规则列表。
	ApplicationProxyRules []*ApplicationProxyRule `json:"ApplicationProxyRules,omitnil,omitempty" name:"ApplicationProxyRules"`

	// 中国大陆加速优化配置。
	AccelerateMainland *AccelerateMainland `json:"AccelerateMainland,omitnil,omitempty" name:"AccelerateMainland"`
}

type ApplicationProxyRule struct {
	// 协议，取值有：
	// <li>TCP：TCP协议；</li>
	// <li>UDP：UDP协议。</li>
	Proto *string `json:"Proto,omitnil,omitempty" name:"Proto"`

	// 端口，支持格式：
	// <li>单个端口，如：80。</li>
	// <li>端口段，如：81-82。表示81，82两个端口。</li>
	// 注意：一条规则最多可填写20个端口。
	Port []*string `json:"Port,omitnil,omitempty" name:"Port"`

	// 源站类型，取值有：
	// <li>custom：手动添加；</li>
	// <li>loadbalancer：负载均衡；</li>
	// <li>origins：源站组。</li>
	OriginType *string `json:"OriginType,omitnil,omitempty" name:"OriginType"`

	// 源站信息：
	// <li>当 OriginType 为 custom 时，表示一个或多个源站，如`["8.8.8.8","9.9.9.9"]` 或 `OriginValue=["test.com"]`；</li>
	// <li>当 OriginType 为 loadbalancer 时，表示一个负载均衡，如`["lb-xdffsfasdfs"]`；</li>
	// <li>当 OriginType 为 origins 时，要求有且仅有一个元素，表示源站组ID，如`["origin-537f5b41-162a-11ed-abaa-525400c5da15"]`。</li>
	OriginValue []*string `json:"OriginValue,omitnil,omitempty" name:"OriginValue"`

	// 规则ID。
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// 状态，取值有：
	// <li>online：启用；</li>
	// <li>offline：停用；</li>
	// <li>progress：部署中；</li>
	// <li>stopping：停用中；</li>
	// <li>fail：部署失败/停用失败。</li>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 传递客户端IP，取值有：
	// <li>TOA：TOA（仅Proto=TCP时可选）；</li>
	// <li>PPV1：Proxy Protocol传递，协议版本V1（仅Proto=TCP时可选）；</li>
	// <li>PPV2：Proxy Protocol传递，协议版本V2；</li>
	// <li>OFF：不传递。</li>默认值：OFF。
	ForwardClientIp *string `json:"ForwardClientIp,omitnil,omitempty" name:"ForwardClientIp"`

	// 是否开启会话保持，取值有：
	// <li>true：开启；</li>
	// <li>false：关闭。</li>默认值：false。
	SessionPersist *bool `json:"SessionPersist,omitnil,omitempty" name:"SessionPersist"`

	// 会话保持的时间，只有当SessionPersist为true时，该值才会生效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SessionPersistTime *uint64 `json:"SessionPersistTime,omitnil,omitempty" name:"SessionPersistTime"`

	// 源站端口，支持格式：
	// <li>单端口，如：80。</li>
	// <li>端口段：81-82，表示81，82两个端口。</li>
	OriginPort *string `json:"OriginPort,omitnil,omitempty" name:"OriginPort"`

	// 规则标签。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleTag *string `json:"RuleTag,omitnil,omitempty" name:"RuleTag"`
}

type AscriptionInfo struct {
	// 主机记录。
	Subdomain *string `json:"Subdomain,omitnil,omitempty" name:"Subdomain"`

	// 记录类型。
	RecordType *string `json:"RecordType,omitnil,omitempty" name:"RecordType"`

	// 记录值。
	RecordValue *string `json:"RecordValue,omitnil,omitempty" name:"RecordValue"`
}

type BillingData struct {
	// 时间。
	Time *string `json:"Time,omitnil,omitempty" name:"Time"`

	// 数值。
	Value *uint64 `json:"Value,omitnil,omitempty" name:"Value"`
}

type BillingDataFilter struct {
	// 参数名称。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 参数值。
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

// Predefined struct for user
type BindSecurityTemplateToEntityRequestParams struct {
	// 需要绑定或解绑的策略模板所属站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 绑定至策略模板（或者从策略模板解绑）的域名列表。
	Entities []*string `json:"Entities,omitnil,omitempty" name:"Entities"`

	// 绑定或解绑操作选项，取值有：
	// <li>bind：绑定域名至策略模板；</li>
	// <li>unbind-keep-policy：将域名从策略模板解绑，解绑时保留当前策略；</li>
	// <li>unbind-use-default：将域名从策略模板解绑，并使用默认空白策略。</li>注意：解绑操作当前仅支持单个域名解绑。即：当 Operate 参数取值为 unbind-keep-policy 或 unbind-use-default 时，Entities 参数列表仅支持填写一个域名。
	Operate *string `json:"Operate,omitnil,omitempty" name:"Operate"`

	// 指定绑定或解绑的策略模板 ID 或站点全局策略
	// - 如需绑定至策略模板，或从策略模板解绑，请指定策略模板 ID。
	// - 如需绑定至站点全局策略，或从站点全局策略解绑，请使用 @ZoneLevel@domain 参数值。
	// 
	// 注意：解绑后，域名将使用独立策略，并单独计算规则配额，请确保解绑前套餐规则配额充足。
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// 如指定的域名已经绑定了策略模板，是否替换该模板。支持下列取值：
	// <li>true： 替换域名当前绑定的模板；</li>
	// <li>false：不替换域名当前绑定的模板。</li>注意：当选择不替换已有策略模板时，若指定域名已经绑定策略模板，API 将返回错误。
	OverWrite *bool `json:"OverWrite,omitnil,omitempty" name:"OverWrite"`
}

type BindSecurityTemplateToEntityRequest struct {
	*tchttp.BaseRequest
	
	// 需要绑定或解绑的策略模板所属站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 绑定至策略模板（或者从策略模板解绑）的域名列表。
	Entities []*string `json:"Entities,omitnil,omitempty" name:"Entities"`

	// 绑定或解绑操作选项，取值有：
	// <li>bind：绑定域名至策略模板；</li>
	// <li>unbind-keep-policy：将域名从策略模板解绑，解绑时保留当前策略；</li>
	// <li>unbind-use-default：将域名从策略模板解绑，并使用默认空白策略。</li>注意：解绑操作当前仅支持单个域名解绑。即：当 Operate 参数取值为 unbind-keep-policy 或 unbind-use-default 时，Entities 参数列表仅支持填写一个域名。
	Operate *string `json:"Operate,omitnil,omitempty" name:"Operate"`

	// 指定绑定或解绑的策略模板 ID 或站点全局策略
	// - 如需绑定至策略模板，或从策略模板解绑，请指定策略模板 ID。
	// - 如需绑定至站点全局策略，或从站点全局策略解绑，请使用 @ZoneLevel@domain 参数值。
	// 
	// 注意：解绑后，域名将使用独立策略，并单独计算规则配额，请确保解绑前套餐规则配额充足。
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// 如指定的域名已经绑定了策略模板，是否替换该模板。支持下列取值：
	// <li>true： 替换域名当前绑定的模板；</li>
	// <li>false：不替换域名当前绑定的模板。</li>注意：当选择不替换已有策略模板时，若指定域名已经绑定策略模板，API 将返回错误。
	OverWrite *bool `json:"OverWrite,omitnil,omitempty" name:"OverWrite"`
}

func (r *BindSecurityTemplateToEntityRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindSecurityTemplateToEntityRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "Entities")
	delete(f, "Operate")
	delete(f, "TemplateId")
	delete(f, "OverWrite")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BindSecurityTemplateToEntityRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BindSecurityTemplateToEntityResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type BindSecurityTemplateToEntityResponse struct {
	*tchttp.BaseResponse
	Response *BindSecurityTemplateToEntityResponseParams `json:"Response"`
}

func (r *BindSecurityTemplateToEntityResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindSecurityTemplateToEntityResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BindSharedCNAMEMap struct {
	// 需要绑定或解绑的共享 CNAME。
	SharedCNAME *string `json:"SharedCNAME,omitnil,omitempty" name:"SharedCNAME"`

	// 加速域名，可传递多个，最多20个。
	DomainNames []*string `json:"DomainNames,omitnil,omitempty" name:"DomainNames"`
}

// Predefined struct for user
type BindSharedCNAMERequestParams struct {
	// 加速域名所属站点 ID。	
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 绑定类型，取值有：
	// <li>bind：绑定；</li>
	// <li>unbind：解绑。</li>
	BindType *string `json:"BindType,omitnil,omitempty" name:"BindType"`

	// 接入域名与共享 CNAME 的绑定关系。
	BindSharedCNAMEMaps []*BindSharedCNAMEMap `json:"BindSharedCNAMEMaps,omitnil,omitempty" name:"BindSharedCNAMEMaps"`
}

type BindSharedCNAMERequest struct {
	*tchttp.BaseRequest
	
	// 加速域名所属站点 ID。	
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 绑定类型，取值有：
	// <li>bind：绑定；</li>
	// <li>unbind：解绑。</li>
	BindType *string `json:"BindType,omitnil,omitempty" name:"BindType"`

	// 接入域名与共享 CNAME 的绑定关系。
	BindSharedCNAMEMaps []*BindSharedCNAMEMap `json:"BindSharedCNAMEMaps,omitnil,omitempty" name:"BindSharedCNAMEMaps"`
}

func (r *BindSharedCNAMERequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindSharedCNAMERequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "BindType")
	delete(f, "BindSharedCNAMEMaps")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BindSharedCNAMERequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BindSharedCNAMEResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type BindSharedCNAMEResponse struct {
	*tchttp.BaseResponse
	Response *BindSharedCNAMEResponseParams `json:"Response"`
}

func (r *BindSharedCNAMEResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindSharedCNAMEResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BindZoneToPlanRequestParams struct {
	// 未绑定套餐的站点ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 待绑定的目标套餐ID。
	PlanId *string `json:"PlanId,omitnil,omitempty" name:"PlanId"`
}

type BindZoneToPlanRequest struct {
	*tchttp.BaseRequest
	
	// 未绑定套餐的站点ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 待绑定的目标套餐ID。
	PlanId *string `json:"PlanId,omitnil,omitempty" name:"PlanId"`
}

func (r *BindZoneToPlanRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindZoneToPlanRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "PlanId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BindZoneToPlanRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BindZoneToPlanResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type BindZoneToPlanResponse struct {
	*tchttp.BaseResponse
	Response *BindZoneToPlanResponseParams `json:"Response"`
}

func (r *BindZoneToPlanResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindZoneToPlanResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BotConfig struct {
	// bot开关，取值有：
	// <li>on：开启；</li>
	// <li>off：关闭。</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// 通用详细基础规则。如果为null，默认使用历史配置。
	BotManagedRule *BotManagedRule `json:"BotManagedRule,omitnil,omitempty" name:"BotManagedRule"`

	// 用户画像规则。如果为null，默认使用历史配置。
	BotPortraitRule *BotPortraitRule `json:"BotPortraitRule,omitnil,omitempty" name:"BotPortraitRule"`

	// Bot智能分析。如果为null，默认使用历史配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	IntelligenceRule *IntelligenceRule `json:"IntelligenceRule,omitnil,omitempty" name:"IntelligenceRule"`

	// Bot自定义规则。如果为null，默认使用历史配置。
	BotUserRules []*BotUserRule `json:"BotUserRules,omitnil,omitempty" name:"BotUserRules"`

	// Bot主动特征识别规则。
	AlgDetectRule []*AlgDetectRule `json:"AlgDetectRule,omitnil,omitempty" name:"AlgDetectRule"`

	// Bot托管定制策略，入参可不填，仅出参使用。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Customizes []*BotUserRule `json:"Customizes,omitnil,omitempty" name:"Customizes"`
}

type BotExtendAction struct {
	// 处置动作，取值有：
	// <li>monitor：观察；</li>
	// <li>alg：JavaScript挑战；</li>
	// <li>captcha：托管挑战；</li>
	// <li>random：随机，按照ExtendActions分配处置动作和比例；</li>
	// <li>silence：静默；</li>
	// <li>shortdelay：短时响应；</li>
	// <li>longdelay：长时响应。</li>
	Action *string `json:"Action,omitnil,omitempty" name:"Action"`

	// 处置方式的触发概率，范围0-100。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Percent *uint64 `json:"Percent,omitnil,omitempty" name:"Percent"`
}

type BotManagedRule struct {
	// 触发规则后的处置方式，取值有：
	// <li>drop：拦截；</li>
	// <li>trans：放行；</li>
	// <li>alg：Javascript挑战；</li>
	// <li>monitor：观察。</li>
	Action *string `json:"Action,omitnil,omitempty" name:"Action"`

	// 本规则的ID。仅出参使用。
	RuleID *int64 `json:"RuleID,omitnil,omitempty" name:"RuleID"`

	// 放行的规则ID。默认所有规则不配置放行。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TransManagedIds []*int64 `json:"TransManagedIds,omitnil,omitempty" name:"TransManagedIds"`

	// JS挑战的规则ID。默认所有规则不配置JS挑战。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AlgManagedIds []*int64 `json:"AlgManagedIds,omitnil,omitempty" name:"AlgManagedIds"`

	// 数字验证码的规则ID。默认所有规则不配置数字验证码。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CapManagedIds []*int64 `json:"CapManagedIds,omitnil,omitempty" name:"CapManagedIds"`

	// 观察的规则ID。默认所有规则不配置观察。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MonManagedIds []*int64 `json:"MonManagedIds,omitnil,omitempty" name:"MonManagedIds"`

	// 拦截的规则ID。默认所有规则不配置拦截。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DropManagedIds []*int64 `json:"DropManagedIds,omitnil,omitempty" name:"DropManagedIds"`
}

type BotPortraitRule struct {
	// 本功能的开关，取值有：
	// <li>on：开启；</li>
	// <li>off：关闭。</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// 本规则的ID。仅出参使用。
	RuleID *int64 `json:"RuleID,omitnil,omitempty" name:"RuleID"`

	// JS挑战的规则ID。默认所有规则不配置JS挑战。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AlgManagedIds []*int64 `json:"AlgManagedIds,omitnil,omitempty" name:"AlgManagedIds"`

	// 数字验证码的规则ID。默认所有规则不配置数字验证码。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CapManagedIds []*int64 `json:"CapManagedIds,omitnil,omitempty" name:"CapManagedIds"`

	// 观察的规则ID。默认所有规则不配置观察。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MonManagedIds []*int64 `json:"MonManagedIds,omitnil,omitempty" name:"MonManagedIds"`

	// 拦截的规则ID。默认所有规则不配置拦截。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DropManagedIds []*int64 `json:"DropManagedIds,omitnil,omitempty" name:"DropManagedIds"`
}

type BotUserRule struct {
	// 规则名，只能以英文字符，数字，下划线组合，且不能以下划线开头。
	RuleName *string `json:"RuleName,omitnil,omitempty" name:"RuleName"`

	// 处置动作，取值有：
	// <li>drop：拦截；</li>
	// <li>monitor：观察；</li>
	// <li>trans：放行；</li>
	// <li>redirect：重定向；</li>
	// <li>page：指定页面；</li>
	// <li>alg：JavaScript 挑战；</li>
	// <li>captcha：托管挑战；</li>
	// <li>random：随机处置；</li>
	// <li>silence：静默；</li>
	// <li>shortdelay：短时响应；</li>
	// <li>longdelay：长时响应。</li>
	Action *string `json:"Action,omitnil,omitempty" name:"Action"`

	// 规则状态，取值有：
	// <li>on：生效；</li>
	// <li>off：不生效。</li>默认 on 生效。
	RuleStatus *string `json:"RuleStatus,omitnil,omitempty" name:"RuleStatus"`

	// 规则详情。
	AclConditions []*AclCondition `json:"AclConditions,omitnil,omitempty" name:"AclConditions"`

	// 规则权重，取值范围0-100。
	RulePriority *int64 `json:"RulePriority,omitnil,omitempty" name:"RulePriority"`

	// 规则 Id。仅出参使用。
	RuleID *int64 `json:"RuleID,omitnil,omitempty" name:"RuleID"`

	// 随机处置的处置方式及占比，非随机处置可不填暂不支持。
	ExtendActions []*BotExtendAction `json:"ExtendActions,omitnil,omitempty" name:"ExtendActions"`

	// 过滤词，取值有：
	// <li>sip：客户端 ip。</li>
	// 默认为空字符串。
	FreqFields []*string `json:"FreqFields,omitnil,omitempty" name:"FreqFields"`

	// 更新时间。仅出参使用。
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 统计范围。取值有：
	// <li>source_to_eo：（响应）源站到 EdgeOne；</li>
	// <li>client_to_eo：（请求）客户端到 EdgeOne。</li>
	// 默认为 source_to_eo。
	FreqScope []*string `json:"FreqScope,omitnil,omitempty" name:"FreqScope"`

	// 自定义返回页面的名称。Action 是 page 时必填，且不能为空。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 自定义响应 Id。该 Id 可通过查询自定义错误页列表接口获取。默认值为default，使用系统默认页面。Action 是 page 时必填，且不能为空。	
	CustomResponseId *string `json:"CustomResponseId,omitnil,omitempty" name:"CustomResponseId"`

	// 自定义返回页面的响应码。Action 是 page 时必填，且不能为空，取值: 100~600，不支持 3xx 响应码。默认值：567。
	ResponseCode *int64 `json:"ResponseCode,omitnil,omitempty" name:"ResponseCode"`

	// 重定向时候的地址。Action 是 redirect 时必填，且不能为空。
	RedirectUrl *string `json:"RedirectUrl,omitnil,omitempty" name:"RedirectUrl"`
}

type CC struct {
	// Waf开关，取值为：
	// <li> on：开启；</li>
	// <li> off：关闭。</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// 策略ID。
	PolicyId *int64 `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`
}

type CLSTopic struct {
	// 腾讯云 CLS 日志集 ID。	
	LogSetId *string `json:"LogSetId,omitnil,omitempty" name:"LogSetId"`

	// 腾讯云 CLS 日志主题 ID。
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// 腾讯云 CLS 日志集所在的地域。
	LogSetRegion *string `json:"LogSetRegion,omitnil,omitempty" name:"LogSetRegion"`
}

type Cache struct {
	// 缓存配置开关，取值有：
	// <li>on：开启；</li>
	// <li>off：关闭。</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// 缓存过期时间设置。
	// 单位为秒，最大可设置为 365 天。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CacheTime *int64 `json:"CacheTime,omitnil,omitempty" name:"CacheTime"`

	// 是否开启强制缓存，取值有：
	// <li>on：开启；</li>
	// <li>off：关闭。</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	//
	// Deprecated: IgnoreCacheControl is deprecated.
	IgnoreCacheControl *string `json:"IgnoreCacheControl,omitnil,omitempty" name:"IgnoreCacheControl"`
}

type CacheConfig struct {
	// 缓存配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Cache *Cache `json:"Cache,omitnil,omitempty" name:"Cache"`

	// 不缓存配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	NoCache *NoCache `json:"NoCache,omitnil,omitempty" name:"NoCache"`

	// 遵循源站配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	FollowOrigin *FollowOrigin `json:"FollowOrigin,omitnil,omitempty" name:"FollowOrigin"`
}

type CacheKey struct {
	// 是否开启全路径缓存，取值有：
	// <li>on：开启全路径缓存（即关闭参数忽略）；</li>
	// <li>off：关闭全路径缓存（即开启参数忽略）。</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	FullUrlCache *string `json:"FullUrlCache,omitnil,omitempty" name:"FullUrlCache"`

	// 是否忽略大小写缓存，取值有：
	// <li>on：忽略；</li>
	// <li>off：不忽略。</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	IgnoreCase *string `json:"IgnoreCase,omitnil,omitempty" name:"IgnoreCase"`

	// CacheKey 中包含请求参数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	QueryString *QueryString `json:"QueryString,omitnil,omitempty" name:"QueryString"`
}

type CachePrefresh struct {
	// 缓存预刷新配置开关，取值有：
	// <li>on：开启；</li>
	// <li>off：关闭。</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// 缓存预刷新百分比，取值范围：1-99。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Percent *int64 `json:"Percent,omitnil,omitempty" name:"Percent"`
}

type CacheTag struct {
	// 待清除缓存的域名列表。
	Domains []*string `json:"Domains,omitnil,omitempty" name:"Domains"`
}

type CertificateInfo struct {
	// 证书 ID。
	CertId *string `json:"CertId,omitnil,omitempty" name:"CertId"`

	// 证书备注名。
	Alias *string `json:"Alias,omitnil,omitempty" name:"Alias"`

	// 证书类型，取值有：
	// <li>default：默认证书；</li>
	// <li>upload：用户上传；</li>
	// <li>managed：腾讯云托管。</li>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 证书过期时间。
	ExpireTime *string `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`

	// 证书部署时间。
	DeployTime *string `json:"DeployTime,omitnil,omitempty" name:"DeployTime"`

	// 签名算法。
	SignAlgo *string `json:"SignAlgo,omitnil,omitempty" name:"SignAlgo"`

	// 证书状态，取值有：
	// <li>deployed：已部署；</li>
	// <li>processing：部署中；</li>
	// <li>applying：申请中；</li>
	// <li>failed：申请失败；</li>
	// <li>issued：绑定失败。</li>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`
}

// Predefined struct for user
type CheckCnameStatusRequestParams struct {
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 加速域名列表。
	RecordNames []*string `json:"RecordNames,omitnil,omitempty" name:"RecordNames"`
}

type CheckCnameStatusRequest struct {
	*tchttp.BaseRequest
	
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 加速域名列表。
	RecordNames []*string `json:"RecordNames,omitnil,omitempty" name:"RecordNames"`
}

func (r *CheckCnameStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckCnameStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "RecordNames")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CheckCnameStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CheckCnameStatusResponseParams struct {
	// 加速域名 CNAME 状态信息列表。
	CnameStatus []*CnameStatus `json:"CnameStatus,omitnil,omitempty" name:"CnameStatus"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CheckCnameStatusResponse struct {
	*tchttp.BaseResponse
	Response *CheckCnameStatusResponseParams `json:"Response"`
}

func (r *CheckCnameStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckCnameStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ClientIpCountry struct {
	// 配置开关，取值有：
	// <li>on：开启；</li>
	// <li>off：关闭。</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// 存放客户端 IP 所属地域信息的请求头名称，当 Switch=on 时有效。
	// 为空则使用默认值：EO-Client-IPCountry。
	HeaderName *string `json:"HeaderName,omitnil,omitempty" name:"HeaderName"`
}

type ClientIpHeader struct {
	// 配置开关，取值有：
	// <li>on：开启；</li>
	// <li>off：关闭。</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// 回源时，存放客户端 IP 的请求头名称。
	// 为空则使用默认值：X-Forwarded-IP。
	// 注意：此字段可能返回 null，表示取不到有效值。
	HeaderName *string `json:"HeaderName,omitnil,omitempty" name:"HeaderName"`
}

type CnameStatus struct {
	// 记录名称。
	RecordName *string `json:"RecordName,omitnil,omitempty" name:"RecordName"`

	// CNAME 地址。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Cname *string `json:"Cname,omitnil,omitempty" name:"Cname"`

	// Cname状态信息，取值有：
	// <li>active：生效；</li>
	// <li>moved：不生效。</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`
}

type CodeAction struct {
	// 功能名称，功能名称填写规范可调用接口 [查询规则引擎的设置参数](https://cloud.tencent.com/document/product/1552/80618) 查看。
	Action *string `json:"Action,omitnil,omitempty" name:"Action"`

	// 操作参数。
	Parameters []*RuleCodeActionParams `json:"Parameters,omitnil,omitempty" name:"Parameters"`
}

type Compression struct {
	// 智能压缩配置开关，取值有：
	// <li>on：开启；</li>
	// <li>off：关闭。</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// 支持的压缩算法列表，取值有：
	// <li>brotli：brotli算法；</li>
	// <li>gzip：gzip算法。</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Algorithms []*string `json:"Algorithms,omitnil,omitempty" name:"Algorithms"`
}

type ConfigGroupVersionInfo struct {
	// 版本 ID。
	VersionId *string `json:"VersionId,omitnil,omitempty" name:"VersionId"`

	// 版本号。
	VersionNumber *string `json:"VersionNumber,omitnil,omitempty" name:"VersionNumber"`

	// 配置组 ID。
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 配置组类型。取值有：
	// <li>l7_acceleration ：七层加速配置组。</li>
	// <li>edge_functions ：边缘函数配置组。</li>
	GroupType *string `json:"GroupType,omitnil,omitempty" name:"GroupType"`

	// 版本描述。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 版本状态，取值有：
	// <li>creating：创建中；</li>
	// <li>inactive：未生效；</li>
	// <li>active：已生效。</li>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 版本创建时间。时间为世界标准时间（UTC）， 遵循 ISO 8601 标准的日期和时间格式。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`
}

// Predefined struct for user
type CreateAccelerationDomainRequestParams struct {
	// 加速域名所属站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 加速域名。
	DomainName *string `json:"DomainName,omitnil,omitempty" name:"DomainName"`

	// 源站信息。
	OriginInfo *OriginInfo `json:"OriginInfo,omitnil,omitempty" name:"OriginInfo"`

	// 回源协议，取值有：
	// <li>FOLLOW: 协议跟随；</li>
	// <li>HTTP: HTTP协议回源；</li>
	// <li>HTTPS: HTTPS协议回源。</li>
	// <li>不填默认为： FOLLOW。</li>
	OriginProtocol *string `json:"OriginProtocol,omitnil,omitempty" name:"OriginProtocol"`

	// HTTP回源端口，取值为1-65535，当OriginProtocol=FOLLOW/HTTP时生效, 不填默认为80。
	HttpOriginPort *uint64 `json:"HttpOriginPort,omitnil,omitempty" name:"HttpOriginPort"`

	// HTTPS回源端口，取值为1-65535，当OriginProtocol=FOLLOW/HTTPS时生效，不填默认为443。
	HttpsOriginPort *uint64 `json:"HttpsOriginPort,omitnil,omitempty" name:"HttpsOriginPort"`

	// IPv6状态，取值有：
	// <li>follow：遵循站点IPv6配置；</li>
	// <li>on：开启状态；</li>
	// <li>off：关闭状态。</li>
	// <li>不填默认为：follow。</li>
	IPv6Status *string `json:"IPv6Status,omitnil,omitempty" name:"IPv6Status"`
}

type CreateAccelerationDomainRequest struct {
	*tchttp.BaseRequest
	
	// 加速域名所属站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 加速域名。
	DomainName *string `json:"DomainName,omitnil,omitempty" name:"DomainName"`

	// 源站信息。
	OriginInfo *OriginInfo `json:"OriginInfo,omitnil,omitempty" name:"OriginInfo"`

	// 回源协议，取值有：
	// <li>FOLLOW: 协议跟随；</li>
	// <li>HTTP: HTTP协议回源；</li>
	// <li>HTTPS: HTTPS协议回源。</li>
	// <li>不填默认为： FOLLOW。</li>
	OriginProtocol *string `json:"OriginProtocol,omitnil,omitempty" name:"OriginProtocol"`

	// HTTP回源端口，取值为1-65535，当OriginProtocol=FOLLOW/HTTP时生效, 不填默认为80。
	HttpOriginPort *uint64 `json:"HttpOriginPort,omitnil,omitempty" name:"HttpOriginPort"`

	// HTTPS回源端口，取值为1-65535，当OriginProtocol=FOLLOW/HTTPS时生效，不填默认为443。
	HttpsOriginPort *uint64 `json:"HttpsOriginPort,omitnil,omitempty" name:"HttpsOriginPort"`

	// IPv6状态，取值有：
	// <li>follow：遵循站点IPv6配置；</li>
	// <li>on：开启状态；</li>
	// <li>off：关闭状态。</li>
	// <li>不填默认为：follow。</li>
	IPv6Status *string `json:"IPv6Status,omitnil,omitempty" name:"IPv6Status"`
}

func (r *CreateAccelerationDomainRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAccelerationDomainRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "DomainName")
	delete(f, "OriginInfo")
	delete(f, "OriginProtocol")
	delete(f, "HttpOriginPort")
	delete(f, "HttpsOriginPort")
	delete(f, "IPv6Status")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAccelerationDomainRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAccelerationDomainResponseParams struct {
	// 当您的站点未进行归属权验证时，您可通过该参数返回的信息单独对域名进行归属权校验。详情参考 [站点/域名归属权验证](https://cloud.tencent.com/document/product/1552/70789)。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OwnershipVerification *OwnershipVerification `json:"OwnershipVerification,omitnil,omitempty" name:"OwnershipVerification"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateAccelerationDomainResponse struct {
	*tchttp.BaseResponse
	Response *CreateAccelerationDomainResponseParams `json:"Response"`
}

func (r *CreateAccelerationDomainResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAccelerationDomainResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAliasDomainRequestParams struct {
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 别称域名名称。
	AliasName *string `json:"AliasName,omitnil,omitempty" name:"AliasName"`

	// 目标域名名称。
	TargetName *string `json:"TargetName,omitnil,omitempty" name:"TargetName"`

	// 证书配置，取值有：
	// <li> none：不配置；</li>
	// <li> hosting：SSL托管证书。</li>默认取值为 none。
	CertType *string `json:"CertType,omitnil,omitempty" name:"CertType"`

	// 当 CertType 取值为 hosting 时需填入相应证书 ID。
	CertId []*string `json:"CertId,omitnil,omitempty" name:"CertId"`
}

type CreateAliasDomainRequest struct {
	*tchttp.BaseRequest
	
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 别称域名名称。
	AliasName *string `json:"AliasName,omitnil,omitempty" name:"AliasName"`

	// 目标域名名称。
	TargetName *string `json:"TargetName,omitnil,omitempty" name:"TargetName"`

	// 证书配置，取值有：
	// <li> none：不配置；</li>
	// <li> hosting：SSL托管证书。</li>默认取值为 none。
	CertType *string `json:"CertType,omitnil,omitempty" name:"CertType"`

	// 当 CertType 取值为 hosting 时需填入相应证书 ID。
	CertId []*string `json:"CertId,omitnil,omitempty" name:"CertId"`
}

func (r *CreateAliasDomainRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAliasDomainRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "AliasName")
	delete(f, "TargetName")
	delete(f, "CertType")
	delete(f, "CertId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAliasDomainRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAliasDomainResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateAliasDomainResponse struct {
	*tchttp.BaseResponse
	Response *CreateAliasDomainResponseParams `json:"Response"`
}

func (r *CreateAliasDomainResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAliasDomainResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateApplicationProxyRequestParams struct {
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 当 ProxyType=hostname 时，表示域名或子域名；
	// 当 ProxyType=instance 时，表示代理名称。
	ProxyName *string `json:"ProxyName,omitnil,omitempty" name:"ProxyName"`

	// 调度模式，取值有：
	// <li>ip：表示Anycast IP调度；</li>
	// <li>domain：表示CNAME调度。</li>
	PlatType *string `json:"PlatType,omitnil,omitempty" name:"PlatType"`

	// 是否开启安全，取值有：
	// <li>0：关闭安全；</li>
	// <li>1：开启安全。</li>
	SecurityType *int64 `json:"SecurityType,omitnil,omitempty" name:"SecurityType"`

	// 是否开启加速，取值有：
	// <li>0：关闭加速；</li>
	// <li>1：开启加速。</li>
	AccelerateType *int64 `json:"AccelerateType,omitnil,omitempty" name:"AccelerateType"`

	// 四层代理模式，取值有： <li>instance：表示实例模式。</li>不填写使用默认值instance。
	ProxyType *string `json:"ProxyType,omitnil,omitempty" name:"ProxyType"`

	// 会话保持时间，取值范围：30-3600，单位：秒。
	// 不填写使用默认值600。
	SessionPersistTime *uint64 `json:"SessionPersistTime,omitnil,omitempty" name:"SessionPersistTime"`

	// Ipv6 访问配置。
	// 不填写表示关闭 Ipv6 访问。
	Ipv6 *Ipv6 `json:"Ipv6,omitnil,omitempty" name:"Ipv6"`

	// 规则详细信息。
	// 不填写则不创建规则。
	ApplicationProxyRules []*ApplicationProxyRule `json:"ApplicationProxyRules,omitnil,omitempty" name:"ApplicationProxyRules"`

	// 中国大陆加速优化配置。不填写表示关闭中国大陆加速优化。
	AccelerateMainland *AccelerateMainland `json:"AccelerateMainland,omitnil,omitempty" name:"AccelerateMainland"`
}

type CreateApplicationProxyRequest struct {
	*tchttp.BaseRequest
	
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 当 ProxyType=hostname 时，表示域名或子域名；
	// 当 ProxyType=instance 时，表示代理名称。
	ProxyName *string `json:"ProxyName,omitnil,omitempty" name:"ProxyName"`

	// 调度模式，取值有：
	// <li>ip：表示Anycast IP调度；</li>
	// <li>domain：表示CNAME调度。</li>
	PlatType *string `json:"PlatType,omitnil,omitempty" name:"PlatType"`

	// 是否开启安全，取值有：
	// <li>0：关闭安全；</li>
	// <li>1：开启安全。</li>
	SecurityType *int64 `json:"SecurityType,omitnil,omitempty" name:"SecurityType"`

	// 是否开启加速，取值有：
	// <li>0：关闭加速；</li>
	// <li>1：开启加速。</li>
	AccelerateType *int64 `json:"AccelerateType,omitnil,omitempty" name:"AccelerateType"`

	// 四层代理模式，取值有： <li>instance：表示实例模式。</li>不填写使用默认值instance。
	ProxyType *string `json:"ProxyType,omitnil,omitempty" name:"ProxyType"`

	// 会话保持时间，取值范围：30-3600，单位：秒。
	// 不填写使用默认值600。
	SessionPersistTime *uint64 `json:"SessionPersistTime,omitnil,omitempty" name:"SessionPersistTime"`

	// Ipv6 访问配置。
	// 不填写表示关闭 Ipv6 访问。
	Ipv6 *Ipv6 `json:"Ipv6,omitnil,omitempty" name:"Ipv6"`

	// 规则详细信息。
	// 不填写则不创建规则。
	ApplicationProxyRules []*ApplicationProxyRule `json:"ApplicationProxyRules,omitnil,omitempty" name:"ApplicationProxyRules"`

	// 中国大陆加速优化配置。不填写表示关闭中国大陆加速优化。
	AccelerateMainland *AccelerateMainland `json:"AccelerateMainland,omitnil,omitempty" name:"AccelerateMainland"`
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
	delete(f, "AccelerateMainland")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateApplicationProxyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateApplicationProxyResponseParams struct {
	// 新增的四层代理应用ID。
	ProxyId *string `json:"ProxyId,omitnil,omitempty" name:"ProxyId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 代理ID。
	ProxyId *string `json:"ProxyId,omitnil,omitempty" name:"ProxyId"`

	// 协议，取值有：
	// <li>TCP：TCP协议；</li>
	// <li>UDP：UDP协议。</li>
	Proto *string `json:"Proto,omitnil,omitempty" name:"Proto"`

	// 端口，支持格式：
	// <li>80：80端口；</li>
	// <li>81-90：81至90端口。</li>
	Port []*string `json:"Port,omitnil,omitempty" name:"Port"`

	// 源站类型，取值有：
	// <li>custom：手动添加；</li>
	// <li>loadbalancer：负载均衡；</li>
	// <li>origins：源站组。</li>
	OriginType *string `json:"OriginType,omitnil,omitempty" name:"OriginType"`

	// 源站信息：
	// <li>当 OriginType 为 custom 时，表示一个或多个源站，如`["8.8.8.8","9.9.9.9"]` 或 `OriginValue=["test.com"]`；</li>
	// <li>当 OriginType 为 loadbalancer 时，表示一个负载均衡，如`["lb-xdffsfasdfs"]`；</li>
	// <li>当 OriginType 为 origins 时，要求有且仅有一个元素，表示源站组ID，如`["origin-537f5b41-162a-11ed-abaa-525400c5da15"]`。</li>
	OriginValue []*string `json:"OriginValue,omitnil,omitempty" name:"OriginValue"`

	// 传递客户端IP，取值有：
	// <li>TOA：TOA（仅Proto=TCP时可选）；</li>
	// <li>PPV1：Proxy Protocol传递，协议版本V1（仅Proto=TCP时可选）；</li>
	// <li>PPV2：Proxy Protocol传递，协议版本V2；</li>
	// <li>OFF：不传递。</li>默认值：OFF。
	ForwardClientIp *string `json:"ForwardClientIp,omitnil,omitempty" name:"ForwardClientIp"`

	// 是否开启会话保持，取值有：
	// <li>true：开启；</li>
	// <li>false：关闭。</li>默认值：false。
	SessionPersist *bool `json:"SessionPersist,omitnil,omitempty" name:"SessionPersist"`

	// 会话保持的时间，只有当SessionPersist为true时，该值才会生效。
	SessionPersistTime *uint64 `json:"SessionPersistTime,omitnil,omitempty" name:"SessionPersistTime"`

	// 源站端口，支持格式：
	// <li>单端口：80；</li>
	// <li>端口段：81-90，81至90端口。</li>
	OriginPort *string `json:"OriginPort,omitnil,omitempty" name:"OriginPort"`

	// 规则标签。默认值为空字符串。
	RuleTag *string `json:"RuleTag,omitnil,omitempty" name:"RuleTag"`
}

type CreateApplicationProxyRuleRequest struct {
	*tchttp.BaseRequest
	
	// 站点ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 代理ID。
	ProxyId *string `json:"ProxyId,omitnil,omitempty" name:"ProxyId"`

	// 协议，取值有：
	// <li>TCP：TCP协议；</li>
	// <li>UDP：UDP协议。</li>
	Proto *string `json:"Proto,omitnil,omitempty" name:"Proto"`

	// 端口，支持格式：
	// <li>80：80端口；</li>
	// <li>81-90：81至90端口。</li>
	Port []*string `json:"Port,omitnil,omitempty" name:"Port"`

	// 源站类型，取值有：
	// <li>custom：手动添加；</li>
	// <li>loadbalancer：负载均衡；</li>
	// <li>origins：源站组。</li>
	OriginType *string `json:"OriginType,omitnil,omitempty" name:"OriginType"`

	// 源站信息：
	// <li>当 OriginType 为 custom 时，表示一个或多个源站，如`["8.8.8.8","9.9.9.9"]` 或 `OriginValue=["test.com"]`；</li>
	// <li>当 OriginType 为 loadbalancer 时，表示一个负载均衡，如`["lb-xdffsfasdfs"]`；</li>
	// <li>当 OriginType 为 origins 时，要求有且仅有一个元素，表示源站组ID，如`["origin-537f5b41-162a-11ed-abaa-525400c5da15"]`。</li>
	OriginValue []*string `json:"OriginValue,omitnil,omitempty" name:"OriginValue"`

	// 传递客户端IP，取值有：
	// <li>TOA：TOA（仅Proto=TCP时可选）；</li>
	// <li>PPV1：Proxy Protocol传递，协议版本V1（仅Proto=TCP时可选）；</li>
	// <li>PPV2：Proxy Protocol传递，协议版本V2；</li>
	// <li>OFF：不传递。</li>默认值：OFF。
	ForwardClientIp *string `json:"ForwardClientIp,omitnil,omitempty" name:"ForwardClientIp"`

	// 是否开启会话保持，取值有：
	// <li>true：开启；</li>
	// <li>false：关闭。</li>默认值：false。
	SessionPersist *bool `json:"SessionPersist,omitnil,omitempty" name:"SessionPersist"`

	// 会话保持的时间，只有当SessionPersist为true时，该值才会生效。
	SessionPersistTime *uint64 `json:"SessionPersistTime,omitnil,omitempty" name:"SessionPersistTime"`

	// 源站端口，支持格式：
	// <li>单端口：80；</li>
	// <li>端口段：81-90，81至90端口。</li>
	OriginPort *string `json:"OriginPort,omitnil,omitempty" name:"OriginPort"`

	// 规则标签。默认值为空字符串。
	RuleTag *string `json:"RuleTag,omitnil,omitempty" name:"RuleTag"`
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
	delete(f, "SessionPersistTime")
	delete(f, "OriginPort")
	delete(f, "RuleTag")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateApplicationProxyRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateApplicationProxyRuleResponseParams struct {
	// 规则ID
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type CreateCLSIndexRequestParams struct {
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 实时日志投递任务 ID。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

type CreateCLSIndexRequest struct {
	*tchttp.BaseRequest
	
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 实时日志投递任务 ID。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

func (r *CreateCLSIndexRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCLSIndexRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCLSIndexRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCLSIndexResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateCLSIndexResponse struct {
	*tchttp.BaseResponse
	Response *CreateCLSIndexResponseParams `json:"Response"`
}

func (r *CreateCLSIndexResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCLSIndexResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateConfigGroupVersionRequestParams struct {
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 待新建版本的配置组 ID。
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 待导入的配置内容。要求采用 JSON 格式，按照 UTF-8 方式进行编码。配置文件内容可参考下方示例。
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`

	// 版本描述，可输入最大长度为 50 个字符，可以通过本字段填写该版本的使用场景等。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type CreateConfigGroupVersionRequest struct {
	*tchttp.BaseRequest
	
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 待新建版本的配置组 ID。
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 待导入的配置内容。要求采用 JSON 格式，按照 UTF-8 方式进行编码。配置文件内容可参考下方示例。
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`

	// 版本描述，可输入最大长度为 50 个字符，可以通过本字段填写该版本的使用场景等。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

func (r *CreateConfigGroupVersionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateConfigGroupVersionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "GroupId")
	delete(f, "Content")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateConfigGroupVersionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateConfigGroupVersionResponseParams struct {
	// 版本 ID。
	VersionId *string `json:"VersionId,omitnil,omitempty" name:"VersionId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateConfigGroupVersionResponse struct {
	*tchttp.BaseResponse
	Response *CreateConfigGroupVersionResponseParams `json:"Response"`
}

func (r *CreateConfigGroupVersionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateConfigGroupVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCustomizeErrorPageRequestParams struct {
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 自定义错误页名称，名称为 2-30 个字符。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 自定义错误页面类型，取值有：<li>text/html； </li><li>application/json；</li><li>text/plain；</li><li>text/xml。</li>
	ContentType *string `json:"ContentType,omitnil,omitempty" name:"ContentType"`

	// 自定义错误页面描述，描述不超过 60 个字符。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 自定义错误页面内容，内容不超过 2KB。
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`
}

type CreateCustomizeErrorPageRequest struct {
	*tchttp.BaseRequest
	
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 自定义错误页名称，名称为 2-30 个字符。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 自定义错误页面类型，取值有：<li>text/html； </li><li>application/json；</li><li>text/plain；</li><li>text/xml。</li>
	ContentType *string `json:"ContentType,omitnil,omitempty" name:"ContentType"`

	// 自定义错误页面描述，描述不超过 60 个字符。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 自定义错误页面内容，内容不超过 2KB。
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`
}

func (r *CreateCustomizeErrorPageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCustomizeErrorPageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "Name")
	delete(f, "ContentType")
	delete(f, "Description")
	delete(f, "Content")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCustomizeErrorPageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCustomizeErrorPageResponseParams struct {
	// 页面 ID。
	PageId *string `json:"PageId,omitnil,omitempty" name:"PageId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateCustomizeErrorPageResponse struct {
	*tchttp.BaseResponse
	Response *CreateCustomizeErrorPageResponseParams `json:"Response"`
}

func (r *CreateCustomizeErrorPageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCustomizeErrorPageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateFunctionRequestParams struct {
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 函数名称，只能包含小写字母、数字、连字符，以数字或字母开头，以数字或字母结尾，最大支持 30 个字符。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 函数内容，当前仅支持 JavaScript 代码，最大支持 5MB 大小。
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`

	// 函数描述，最大支持 60 个字符。
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`
}

type CreateFunctionRequest struct {
	*tchttp.BaseRequest
	
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 函数名称，只能包含小写字母、数字、连字符，以数字或字母开头，以数字或字母结尾，最大支持 30 个字符。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 函数内容，当前仅支持 JavaScript 代码，最大支持 5MB 大小。
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`

	// 函数描述，最大支持 60 个字符。
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`
}

func (r *CreateFunctionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateFunctionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "Name")
	delete(f, "Content")
	delete(f, "Remark")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateFunctionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateFunctionResponseParams struct {
	// 函数 ID。
	FunctionId *string `json:"FunctionId,omitnil,omitempty" name:"FunctionId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateFunctionResponse struct {
	*tchttp.BaseResponse
	Response *CreateFunctionResponseParams `json:"Response"`
}

func (r *CreateFunctionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateFunctionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateFunctionRuleRequestParams struct {
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 规则条件列表，相同触发规则的不同条件匹配项之间为或关系。
	FunctionRuleConditions []*FunctionRuleCondition `json:"FunctionRuleConditions,omitnil,omitempty" name:"FunctionRuleConditions"`

	// 函数 ID，命中触发规则条件后执行的函数。
	FunctionId *string `json:"FunctionId,omitnil,omitempty" name:"FunctionId"`

	// 规则描述，最大支持 60 个字符。
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`
}

type CreateFunctionRuleRequest struct {
	*tchttp.BaseRequest
	
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 规则条件列表，相同触发规则的不同条件匹配项之间为或关系。
	FunctionRuleConditions []*FunctionRuleCondition `json:"FunctionRuleConditions,omitnil,omitempty" name:"FunctionRuleConditions"`

	// 函数 ID，命中触发规则条件后执行的函数。
	FunctionId *string `json:"FunctionId,omitnil,omitempty" name:"FunctionId"`

	// 规则描述，最大支持 60 个字符。
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`
}

func (r *CreateFunctionRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateFunctionRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "FunctionRuleConditions")
	delete(f, "FunctionId")
	delete(f, "Remark")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateFunctionRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateFunctionRuleResponseParams struct {
	// 规则 ID。
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateFunctionRuleResponse struct {
	*tchttp.BaseResponse
	Response *CreateFunctionRuleResponseParams `json:"Response"`
}

func (r *CreateFunctionRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateFunctionRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateL4ProxyRequestParams struct {
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 四层代理实例名称，可输入 1-50 个字符，允许的字符为 a-z、0-9、-，且 - 不能单独注册或连续使用，不能放在开头或结尾。创建完成后不支持修改。
	ProxyName *string `json:"ProxyName,omitnil,omitempty" name:"ProxyName"`

	// 四层代理实例加速区域。
	// <li>mainland：中国大陆可用区；</li>
	// <li>overseas：全球可用区（不含中国大陆）；</li>
	// <li>global：全球可用区。</li>
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 是否开启 IPv6 访问，不填写时默认为 off。该配置仅在部分加速区域和安全防护配置下支持开启，详情请参考 [新建四层代理实例](https://cloud.tencent.com/document/product/1552/90025) 。取值为：
	// <li>on：开启；</li>
	// <li>off：关闭。</li>
	// 
	Ipv6 *string `json:"Ipv6,omitnil,omitempty" name:"Ipv6"`

	// 是否开启固定 IP，不填写时默认为 off。该配置仅在部分加速区域和安全防护配置下支持开启，详情请参考 [新建四层代理实例](https://cloud.tencent.com/document/product/1552/90025) 。取值为：
	// <li>on：开启；</li>
	// <li>off：关闭。</li>
	StaticIp *string `json:"StaticIp,omitnil,omitempty" name:"StaticIp"`

	// 是否开启中国大陆网络优化，不填写时默认为 off。该配置仅在部分加速区域和安全防护配置下支持开启，详情请参考 [新建四层代理实例](https://cloud.tencent.com/document/product/1552/90025) 。取值为：
	// <li>on：开启；</li>
	// <li>off：关闭。</li>
	AccelerateMainland *string `json:"AccelerateMainland,omitnil,omitempty" name:"AccelerateMainland"`

	// L3/L4 DDoS 防护配置，不填写时默认使用平台默认防护选项。详情参考 [独立 DDoS 防护](https://cloud.tencent.com/document/product/1552/95994)。
	DDosProtectionConfig *DDosProtectionConfig `json:"DDosProtectionConfig,omitnil,omitempty" name:"DDosProtectionConfig"`
}

type CreateL4ProxyRequest struct {
	*tchttp.BaseRequest
	
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 四层代理实例名称，可输入 1-50 个字符，允许的字符为 a-z、0-9、-，且 - 不能单独注册或连续使用，不能放在开头或结尾。创建完成后不支持修改。
	ProxyName *string `json:"ProxyName,omitnil,omitempty" name:"ProxyName"`

	// 四层代理实例加速区域。
	// <li>mainland：中国大陆可用区；</li>
	// <li>overseas：全球可用区（不含中国大陆）；</li>
	// <li>global：全球可用区。</li>
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 是否开启 IPv6 访问，不填写时默认为 off。该配置仅在部分加速区域和安全防护配置下支持开启，详情请参考 [新建四层代理实例](https://cloud.tencent.com/document/product/1552/90025) 。取值为：
	// <li>on：开启；</li>
	// <li>off：关闭。</li>
	// 
	Ipv6 *string `json:"Ipv6,omitnil,omitempty" name:"Ipv6"`

	// 是否开启固定 IP，不填写时默认为 off。该配置仅在部分加速区域和安全防护配置下支持开启，详情请参考 [新建四层代理实例](https://cloud.tencent.com/document/product/1552/90025) 。取值为：
	// <li>on：开启；</li>
	// <li>off：关闭。</li>
	StaticIp *string `json:"StaticIp,omitnil,omitempty" name:"StaticIp"`

	// 是否开启中国大陆网络优化，不填写时默认为 off。该配置仅在部分加速区域和安全防护配置下支持开启，详情请参考 [新建四层代理实例](https://cloud.tencent.com/document/product/1552/90025) 。取值为：
	// <li>on：开启；</li>
	// <li>off：关闭。</li>
	AccelerateMainland *string `json:"AccelerateMainland,omitnil,omitempty" name:"AccelerateMainland"`

	// L3/L4 DDoS 防护配置，不填写时默认使用平台默认防护选项。详情参考 [独立 DDoS 防护](https://cloud.tencent.com/document/product/1552/95994)。
	DDosProtectionConfig *DDosProtectionConfig `json:"DDosProtectionConfig,omitnil,omitempty" name:"DDosProtectionConfig"`
}

func (r *CreateL4ProxyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateL4ProxyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "ProxyName")
	delete(f, "Area")
	delete(f, "Ipv6")
	delete(f, "StaticIp")
	delete(f, "AccelerateMainland")
	delete(f, "DDosProtectionConfig")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateL4ProxyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateL4ProxyResponseParams struct {
	// 四层实例 ID。
	ProxyId *string `json:"ProxyId,omitnil,omitempty" name:"ProxyId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateL4ProxyResponse struct {
	*tchttp.BaseResponse
	Response *CreateL4ProxyResponseParams `json:"Response"`
}

func (r *CreateL4ProxyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateL4ProxyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateL4ProxyRulesRequestParams struct {
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 四层代理实例 ID。
	ProxyId *string `json:"ProxyId,omitnil,omitempty" name:"ProxyId"`

	// 转发规则列表。单次最多支持 200 条转发规则。
	// 注意：L4ProxyRule 在此处使用时，Protocol、PortRange、OriginType、OriginValue、OriginPortRange 为必填字段；ClientIPPassThroughMode、SessionPersist、SessionPersistTime、RuleTag 均为选填字段；RuleId、Status 请勿填写。
	L4ProxyRules []*L4ProxyRule `json:"L4ProxyRules,omitnil,omitempty" name:"L4ProxyRules"`
}

type CreateL4ProxyRulesRequest struct {
	*tchttp.BaseRequest
	
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 四层代理实例 ID。
	ProxyId *string `json:"ProxyId,omitnil,omitempty" name:"ProxyId"`

	// 转发规则列表。单次最多支持 200 条转发规则。
	// 注意：L4ProxyRule 在此处使用时，Protocol、PortRange、OriginType、OriginValue、OriginPortRange 为必填字段；ClientIPPassThroughMode、SessionPersist、SessionPersistTime、RuleTag 均为选填字段；RuleId、Status 请勿填写。
	L4ProxyRules []*L4ProxyRule `json:"L4ProxyRules,omitnil,omitempty" name:"L4ProxyRules"`
}

func (r *CreateL4ProxyRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateL4ProxyRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "ProxyId")
	delete(f, "L4ProxyRules")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateL4ProxyRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateL4ProxyRulesResponseParams struct {
	// 新增转发规则的 ID，以数组的形式返回。
	L4ProxyRuleIds []*string `json:"L4ProxyRuleIds,omitnil,omitempty" name:"L4ProxyRuleIds"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateL4ProxyRulesResponse struct {
	*tchttp.BaseResponse
	Response *CreateL4ProxyRulesResponseParams `json:"Response"`
}

func (r *CreateL4ProxyRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateL4ProxyRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateOriginGroupRequestParams struct {
	// 站点 ID
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 源站组名称，可输入1 - 200个字符，允许的字符为 a - z, A - Z, 0 - 9, _, - 。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 源站组类型，此参数必填，取值有：
	// <li>GENERAL：通用型源站组，仅支持添加 IP/域名 源站，可以被域名服务、规则引擎、四层代理、通用型负载均衡、HTTP 专用型负载均衡引用；</li>
	// <li>HTTP： HTTP 专用型源站组，支持添加 IP/域名、对象存储源站作为源站，无法被四层代理引用，仅支持被添加加速域名、规则引擎-修改源站、HTTP 专用型负载均衡引用。</li>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 源站记录信息，此参数必填。
	Records []*OriginRecord `json:"Records,omitnil,omitempty" name:"Records"`

	// 回源 Host Header，仅 Type = HTTP 时传入生效，规则引擎修改 Host Header 配置优先级高于源站组的 Host Header。
	HostHeader *string `json:"HostHeader,omitnil,omitempty" name:"HostHeader"`
}

type CreateOriginGroupRequest struct {
	*tchttp.BaseRequest
	
	// 站点 ID
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 源站组名称，可输入1 - 200个字符，允许的字符为 a - z, A - Z, 0 - 9, _, - 。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 源站组类型，此参数必填，取值有：
	// <li>GENERAL：通用型源站组，仅支持添加 IP/域名 源站，可以被域名服务、规则引擎、四层代理、通用型负载均衡、HTTP 专用型负载均衡引用；</li>
	// <li>HTTP： HTTP 专用型源站组，支持添加 IP/域名、对象存储源站作为源站，无法被四层代理引用，仅支持被添加加速域名、规则引擎-修改源站、HTTP 专用型负载均衡引用。</li>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 源站记录信息，此参数必填。
	Records []*OriginRecord `json:"Records,omitnil,omitempty" name:"Records"`

	// 回源 Host Header，仅 Type = HTTP 时传入生效，规则引擎修改 Host Header 配置优先级高于源站组的 Host Header。
	HostHeader *string `json:"HostHeader,omitnil,omitempty" name:"HostHeader"`
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
	delete(f, "Name")
	delete(f, "Type")
	delete(f, "Records")
	delete(f, "HostHeader")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateOriginGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateOriginGroupResponseParams struct {
	// 源站组ID。
	OriginGroupId *string `json:"OriginGroupId,omitnil,omitempty" name:"OriginGroupId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 所要购买套餐的类型，取值有：
	// <li> sta: 全球内容分发网络（不包括中国大陆）标准版套餐； </li>
	// <li> sta_with_bot: 全球内容分发网络（不包括中国大陆）标准版套餐附带bot管理；</li>
	// <li> sta_cm: 中国大陆内容分发网络标准版套餐； </li>
	// <li> sta_cm_with_bot: 中国大陆内容分发网络标准版套餐附带bot管理；</li>
	// <li> sta_global ：全球内容分发网络（包括中国大陆）标准版套餐； </li>
	// <li> sta_global_with_bot ：全球内容分发网络（包括中国大陆）标准版套餐附带bot管理；</li>
	// <li> ent: 全球内容分发网络（不包括中国大陆）企业版套餐； </li>
	// <li> ent_with_bot: 全球内容分发网络（不包括中国大陆）企业版套餐附带bot管理；</li>
	// <li> ent_cm: 中国大陆内容分发网络企业版套餐； </li>
	// <li> ent_cm_with_bot: 中国大陆内容分发网络企业版套餐附带bot管理。</li>
	// <li> ent_global ：全球内容分发网络（包括中国大陆）企业版套餐； </li>
	// <li> ent_global_with_bot ：全球内容分发网络（包括中国大陆）企业版套餐附带bot管理。</li>当前账户可购买套餐类型请以<a href="https://cloud.tencent.com/document/product/1552/80606">DescribeAvailablePlans</a>返回为准。
	PlanType *string `json:"PlanType,omitnil,omitempty" name:"PlanType"`
}

type CreatePlanForZoneRequest struct {
	*tchttp.BaseRequest
	
	// 站点ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 所要购买套餐的类型，取值有：
	// <li> sta: 全球内容分发网络（不包括中国大陆）标准版套餐； </li>
	// <li> sta_with_bot: 全球内容分发网络（不包括中国大陆）标准版套餐附带bot管理；</li>
	// <li> sta_cm: 中国大陆内容分发网络标准版套餐； </li>
	// <li> sta_cm_with_bot: 中国大陆内容分发网络标准版套餐附带bot管理；</li>
	// <li> sta_global ：全球内容分发网络（包括中国大陆）标准版套餐； </li>
	// <li> sta_global_with_bot ：全球内容分发网络（包括中国大陆）标准版套餐附带bot管理；</li>
	// <li> ent: 全球内容分发网络（不包括中国大陆）企业版套餐； </li>
	// <li> ent_with_bot: 全球内容分发网络（不包括中国大陆）企业版套餐附带bot管理；</li>
	// <li> ent_cm: 中国大陆内容分发网络企业版套餐； </li>
	// <li> ent_cm_with_bot: 中国大陆内容分发网络企业版套餐附带bot管理。</li>
	// <li> ent_global ：全球内容分发网络（包括中国大陆）企业版套餐； </li>
	// <li> ent_global_with_bot ：全球内容分发网络（包括中国大陆）企业版套餐附带bot管理。</li>当前账户可购买套餐类型请以<a href="https://cloud.tencent.com/document/product/1552/80606">DescribeAvailablePlans</a>返回为准。
	PlanType *string `json:"PlanType,omitnil,omitempty" name:"PlanType"`
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
	ResourceNames []*string `json:"ResourceNames,omitnil,omitempty" name:"ResourceNames"`

	// 购买的订单号列表。
	DealNames []*string `json:"DealNames,omitnil,omitempty" name:"DealNames"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type CreatePlanRequestParams struct {
	// 订阅的套餐类型，取值有：<li> personal：个人版套餐，预付费套餐；</li><li> basic：基础版套餐，预付费套餐；</li><li> standard：标准版套餐，预付费套餐；</li><li> enterprise：企业版套餐，后付费套餐。</li>订阅预付费套餐时，请确保账号内有足够余额，余额不足会产生一个待支付的订单。
	// 计费概述参考 [Edgeone计费概述](https://cloud.tencent.com/document/product/1552/94156)
	// 不同套餐区别参考 [Edgeone计费套餐选型对比](https://cloud.tencent.com/document/product/1552/94165)
	PlanType *string `json:"PlanType,omitnil,omitempty" name:"PlanType"`

	// 是否自动使用代金券，取值有：<li> true：是；</li><li> false：否。</li>该参数仅在 PlanType 为 personal, basic, standard 时有效。
	// 不填写使用默认值 false。
	AutoUseVoucher *string `json:"AutoUseVoucher,omitnil,omitempty" name:"AutoUseVoucher"`

	// 订阅预付费套餐参数，PlanType 为 personal, basic, standard 时，可以选填该参数，用于传入套餐的订阅时长和是否开启自动续费。
	// 不填该参数时，默认开通套餐时长为 1 个月，不开启自动续费。
	PrepaidPlanParam *PrepaidPlanParam `json:"PrepaidPlanParam,omitnil,omitempty" name:"PrepaidPlanParam"`
}

type CreatePlanRequest struct {
	*tchttp.BaseRequest
	
	// 订阅的套餐类型，取值有：<li> personal：个人版套餐，预付费套餐；</li><li> basic：基础版套餐，预付费套餐；</li><li> standard：标准版套餐，预付费套餐；</li><li> enterprise：企业版套餐，后付费套餐。</li>订阅预付费套餐时，请确保账号内有足够余额，余额不足会产生一个待支付的订单。
	// 计费概述参考 [Edgeone计费概述](https://cloud.tencent.com/document/product/1552/94156)
	// 不同套餐区别参考 [Edgeone计费套餐选型对比](https://cloud.tencent.com/document/product/1552/94165)
	PlanType *string `json:"PlanType,omitnil,omitempty" name:"PlanType"`

	// 是否自动使用代金券，取值有：<li> true：是；</li><li> false：否。</li>该参数仅在 PlanType 为 personal, basic, standard 时有效。
	// 不填写使用默认值 false。
	AutoUseVoucher *string `json:"AutoUseVoucher,omitnil,omitempty" name:"AutoUseVoucher"`

	// 订阅预付费套餐参数，PlanType 为 personal, basic, standard 时，可以选填该参数，用于传入套餐的订阅时长和是否开启自动续费。
	// 不填该参数时，默认开通套餐时长为 1 个月，不开启自动续费。
	PrepaidPlanParam *PrepaidPlanParam `json:"PrepaidPlanParam,omitnil,omitempty" name:"PrepaidPlanParam"`
}

func (r *CreatePlanRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePlanRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PlanType")
	delete(f, "AutoUseVoucher")
	delete(f, "PrepaidPlanParam")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreatePlanRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePlanResponseParams struct {
	// 套餐 ID，形如 edgeone-2unuvzjmmn2q。
	PlanId *string `json:"PlanId,omitnil,omitempty" name:"PlanId"`

	// 订单号。
	DealName *string `json:"DealName,omitnil,omitempty" name:"DealName"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreatePlanResponse struct {
	*tchttp.BaseResponse
	Response *CreatePlanResponseParams `json:"Response"`
}

func (r *CreatePlanResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePlanResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePrefetchTaskRequestParams struct {
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 要预热的资源列表，每个元素格式类似如下:
	// http://www.example.com/example.txt。参数值当前必填。
	// 注意：提交任务数受计费套餐配额限制，请查看 [EO计费套餐](https://cloud.tencent.com/document/product/1552/77380)。
	Targets []*string `json:"Targets,omitnil,omitempty" name:"Targets"`

	// 是否对url进行encode，若内容含有非 ASCII 字符集的字符，请开启此开关进行编码转换（编码规则遵循 RFC3986）。
	EncodeUrl *bool `json:"EncodeUrl,omitnil,omitempty" name:"EncodeUrl"`

	// 附带的http头部信息。
	Headers []*Header `json:"Headers,omitnil,omitempty" name:"Headers"`
}

type CreatePrefetchTaskRequest struct {
	*tchttp.BaseRequest
	
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 要预热的资源列表，每个元素格式类似如下:
	// http://www.example.com/example.txt。参数值当前必填。
	// 注意：提交任务数受计费套餐配额限制，请查看 [EO计费套餐](https://cloud.tencent.com/document/product/1552/77380)。
	Targets []*string `json:"Targets,omitnil,omitempty" name:"Targets"`

	// 是否对url进行encode，若内容含有非 ASCII 字符集的字符，请开启此开关进行编码转换（编码规则遵循 RFC3986）。
	EncodeUrl *bool `json:"EncodeUrl,omitnil,omitempty" name:"EncodeUrl"`

	// 附带的http头部信息。
	Headers []*Header `json:"Headers,omitnil,omitempty" name:"Headers"`
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
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// 失败的任务列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	FailedList []*FailReason `json:"FailedList,omitnil,omitempty" name:"FailedList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 节点缓存清除类型，取值有：
	// <li>purge_url：URL刷新；</li>
	// <li>purge_prefix：目录刷新；</li>
	// <li>purge_host：Hostname 刷新；</li>
	// <li>purge_all：站点下全部缓存刷新；</li>
	// <li>purge_cache_tag：cache-tag 刷新。</li>缓存清除类型详情请查看[清除缓存](https://cloud.tencent.com/document/product/1552/70759)。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 节点缓存清除方法，针对目录刷新、Hostname刷新以及刷新全部缓存类型有效，取值有：<li> invalidate：仅刷新目录下产生了更新的资源；</li><li> delete：无论目录下资源是否更新都刷新节点资源。</li>默认值： invalidate。
	Method *string `json:"Method,omitnil,omitempty" name:"Method"`

	// 要清除缓存的资源列表。每个元素格式依据清除缓存类型而定，可参考接口示例。<li>单次提交的任务数受计费套餐配额限制，请查看 [EO计费套餐](https://cloud.tencent.com/document/product/1552/77380)。</li>
	Targets []*string `json:"Targets,omitnil,omitempty" name:"Targets"`

	// 若有编码转换，仅清除编码转换后匹配的资源。
	// 若内容含有非 ASCII 字符集的字符，请开启此开关进行编码转换（编码规则遵循 RFC3986）。
	//
	// Deprecated: EncodeUrl is deprecated.
	EncodeUrl *bool `json:"EncodeUrl,omitnil,omitempty" name:"EncodeUrl"`

	// 节点缓存清除类型取值为 purge_cache_tag 时附带的信息。
	CacheTag *CacheTag `json:"CacheTag,omitnil,omitempty" name:"CacheTag"`
}

type CreatePurgeTaskRequest struct {
	*tchttp.BaseRequest
	
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 节点缓存清除类型，取值有：
	// <li>purge_url：URL刷新；</li>
	// <li>purge_prefix：目录刷新；</li>
	// <li>purge_host：Hostname 刷新；</li>
	// <li>purge_all：站点下全部缓存刷新；</li>
	// <li>purge_cache_tag：cache-tag 刷新。</li>缓存清除类型详情请查看[清除缓存](https://cloud.tencent.com/document/product/1552/70759)。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 节点缓存清除方法，针对目录刷新、Hostname刷新以及刷新全部缓存类型有效，取值有：<li> invalidate：仅刷新目录下产生了更新的资源；</li><li> delete：无论目录下资源是否更新都刷新节点资源。</li>默认值： invalidate。
	Method *string `json:"Method,omitnil,omitempty" name:"Method"`

	// 要清除缓存的资源列表。每个元素格式依据清除缓存类型而定，可参考接口示例。<li>单次提交的任务数受计费套餐配额限制，请查看 [EO计费套餐](https://cloud.tencent.com/document/product/1552/77380)。</li>
	Targets []*string `json:"Targets,omitnil,omitempty" name:"Targets"`

	// 若有编码转换，仅清除编码转换后匹配的资源。
	// 若内容含有非 ASCII 字符集的字符，请开启此开关进行编码转换（编码规则遵循 RFC3986）。
	EncodeUrl *bool `json:"EncodeUrl,omitnil,omitempty" name:"EncodeUrl"`

	// 节点缓存清除类型取值为 purge_cache_tag 时附带的信息。
	CacheTag *CacheTag `json:"CacheTag,omitnil,omitempty" name:"CacheTag"`
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
	delete(f, "Method")
	delete(f, "Targets")
	delete(f, "EncodeUrl")
	delete(f, "CacheTag")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreatePurgeTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePurgeTaskResponseParams struct {
	// 任务 ID。
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// 失败的任务列表及原因。
	// 注意：此字段可能返回 null，表示取不到有效值。
	FailedList []*FailReason `json:"FailedList,omitnil,omitempty" name:"FailedList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type CreateRealtimeLogDeliveryTaskRequestParams struct {
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 实时日志投递任务的名称，格式为数字、英文、-和_组合，最多 200 个字符。
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`

	// 实时日志投递任务类型，取值有：
	// <li>cls: 推送到腾讯云 CLS；</li>
	// <li>custom_endpoint：推送到自定义 HTTP(S) 地址；</li>
	// <li>s3：推送到 AWS S3 兼容存储桶地址。</li>
	TaskType *string `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// 实时日志投递任务对应的实体（七层域名或者四层代理实例）列表。取值示例如下：
	// <li>七层域名：domain.example.com；</li>
	// <li>四层代理实例：sid-2s69eb5wcms7。</li>
	EntityList []*string `json:"EntityList,omitnil,omitempty" name:"EntityList"`

	// 数据投递类型，取值有：
	// <li>domain：站点加速日志；</li>
	// <li>application：四层代理日志；</li>
	// <li>web-rateLiming：速率限制和 CC 攻击防护日志；</li>
	// <li>web-attack：托管规则日志；</li>
	// <li>web-rule：自定义规则日志；</li>
	// <li>web-bot：Bot管理日志。</li>
	LogType *string `json:"LogType,omitnil,omitempty" name:"LogType"`

	// 数据投递区域，取值有：
	// <li>mainland：中国大陆境内；</li>
	// <li>overseas：全球（不含中国大陆）。</li>
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 投递的预设字段列表。
	Fields []*string `json:"Fields,omitnil,omitempty" name:"Fields"`

	// 投递的自定义字段列表，支持在 HTTP 请求头、响应头、Cookie 中提取指定字段值。自定义字段名称不能重复，且最多不能超过 200 个字段。
	CustomFields []*CustomField `json:"CustomFields,omitnil,omitempty" name:"CustomFields"`

	// 日志投递的过滤条件，不填表示投递全量日志。
	DeliveryConditions []*DeliveryCondition `json:"DeliveryConditions,omitnil,omitempty" name:"DeliveryConditions"`

	// 采样比例，采用千分制，取值范围为1-1000，例如：填写 605 表示采样比例为 60.5%。不填表示采样比例为 100%。
	Sample *uint64 `json:"Sample,omitnil,omitempty" name:"Sample"`

	// 日志投递的输出格式。不填表示为默认格式，默认格式逻辑如下：
	// <li>当 TaskType 取值为 custom_endpoint 时，默认格式为多个 JSON 对象组成的数组，每个 JSON 对象为一条日志；</li>
	// <li>当 TaskType 取值为 s3 时，默认格式为 JSON Lines；</li>特别地，当 TaskType 取值为 cls 时，LogFormat.FormatType 的值只能为 json，且 LogFormat 中其他参数将被忽略，建议不传 LogFormat。
	LogFormat *LogFormat `json:"LogFormat,omitnil,omitempty" name:"LogFormat"`

	// CLS 的配置信息。当 TaskType 取值为 cls 时，该参数必填。
	CLS *CLSTopic `json:"CLS,omitnil,omitempty" name:"CLS"`

	// 自定义 HTTP 服务的配置信息。当 TaskType 取值为 custom_endpoint 时，该参数必填。
	CustomEndpoint *CustomEndpoint `json:"CustomEndpoint,omitnil,omitempty" name:"CustomEndpoint"`

	// AWS S3 兼容存储桶的配置信息。当 TaskType 取值为 s3 时，该参数必填。
	S3 *S3 `json:"S3,omitnil,omitempty" name:"S3"`
}

type CreateRealtimeLogDeliveryTaskRequest struct {
	*tchttp.BaseRequest
	
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 实时日志投递任务的名称，格式为数字、英文、-和_组合，最多 200 个字符。
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`

	// 实时日志投递任务类型，取值有：
	// <li>cls: 推送到腾讯云 CLS；</li>
	// <li>custom_endpoint：推送到自定义 HTTP(S) 地址；</li>
	// <li>s3：推送到 AWS S3 兼容存储桶地址。</li>
	TaskType *string `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// 实时日志投递任务对应的实体（七层域名或者四层代理实例）列表。取值示例如下：
	// <li>七层域名：domain.example.com；</li>
	// <li>四层代理实例：sid-2s69eb5wcms7。</li>
	EntityList []*string `json:"EntityList,omitnil,omitempty" name:"EntityList"`

	// 数据投递类型，取值有：
	// <li>domain：站点加速日志；</li>
	// <li>application：四层代理日志；</li>
	// <li>web-rateLiming：速率限制和 CC 攻击防护日志；</li>
	// <li>web-attack：托管规则日志；</li>
	// <li>web-rule：自定义规则日志；</li>
	// <li>web-bot：Bot管理日志。</li>
	LogType *string `json:"LogType,omitnil,omitempty" name:"LogType"`

	// 数据投递区域，取值有：
	// <li>mainland：中国大陆境内；</li>
	// <li>overseas：全球（不含中国大陆）。</li>
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 投递的预设字段列表。
	Fields []*string `json:"Fields,omitnil,omitempty" name:"Fields"`

	// 投递的自定义字段列表，支持在 HTTP 请求头、响应头、Cookie 中提取指定字段值。自定义字段名称不能重复，且最多不能超过 200 个字段。
	CustomFields []*CustomField `json:"CustomFields,omitnil,omitempty" name:"CustomFields"`

	// 日志投递的过滤条件，不填表示投递全量日志。
	DeliveryConditions []*DeliveryCondition `json:"DeliveryConditions,omitnil,omitempty" name:"DeliveryConditions"`

	// 采样比例，采用千分制，取值范围为1-1000，例如：填写 605 表示采样比例为 60.5%。不填表示采样比例为 100%。
	Sample *uint64 `json:"Sample,omitnil,omitempty" name:"Sample"`

	// 日志投递的输出格式。不填表示为默认格式，默认格式逻辑如下：
	// <li>当 TaskType 取值为 custom_endpoint 时，默认格式为多个 JSON 对象组成的数组，每个 JSON 对象为一条日志；</li>
	// <li>当 TaskType 取值为 s3 时，默认格式为 JSON Lines；</li>特别地，当 TaskType 取值为 cls 时，LogFormat.FormatType 的值只能为 json，且 LogFormat 中其他参数将被忽略，建议不传 LogFormat。
	LogFormat *LogFormat `json:"LogFormat,omitnil,omitempty" name:"LogFormat"`

	// CLS 的配置信息。当 TaskType 取值为 cls 时，该参数必填。
	CLS *CLSTopic `json:"CLS,omitnil,omitempty" name:"CLS"`

	// 自定义 HTTP 服务的配置信息。当 TaskType 取值为 custom_endpoint 时，该参数必填。
	CustomEndpoint *CustomEndpoint `json:"CustomEndpoint,omitnil,omitempty" name:"CustomEndpoint"`

	// AWS S3 兼容存储桶的配置信息。当 TaskType 取值为 s3 时，该参数必填。
	S3 *S3 `json:"S3,omitnil,omitempty" name:"S3"`
}

func (r *CreateRealtimeLogDeliveryTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRealtimeLogDeliveryTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "TaskName")
	delete(f, "TaskType")
	delete(f, "EntityList")
	delete(f, "LogType")
	delete(f, "Area")
	delete(f, "Fields")
	delete(f, "CustomFields")
	delete(f, "DeliveryConditions")
	delete(f, "Sample")
	delete(f, "LogFormat")
	delete(f, "CLS")
	delete(f, "CustomEndpoint")
	delete(f, "S3")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateRealtimeLogDeliveryTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateRealtimeLogDeliveryTaskResponseParams struct {
	// 创建成功的任务ID。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateRealtimeLogDeliveryTaskResponse struct {
	*tchttp.BaseResponse
	Response *CreateRealtimeLogDeliveryTaskResponseParams `json:"Response"`
}

func (r *CreateRealtimeLogDeliveryTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRealtimeLogDeliveryTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateRuleRequestParams struct {
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 规则名称，名称字符串长度 1～255。
	RuleName *string `json:"RuleName,omitnil,omitempty" name:"RuleName"`

	// 规则状态，取值有：
	// <li> enable: 启用； </li>
	// <li> disable: 未启用。</li>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 规则内容。
	Rules []*Rule `json:"Rules,omitnil,omitempty" name:"Rules"`

	// 规则标签。
	Tags []*string `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type CreateRuleRequest struct {
	*tchttp.BaseRequest
	
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 规则名称，名称字符串长度 1～255。
	RuleName *string `json:"RuleName,omitnil,omitempty" name:"RuleName"`

	// 规则状态，取值有：
	// <li> enable: 启用； </li>
	// <li> disable: 未启用。</li>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 规则内容。
	Rules []*Rule `json:"Rules,omitnil,omitempty" name:"Rules"`

	// 规则标签。
	Tags []*string `json:"Tags,omitnil,omitempty" name:"Tags"`
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
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateRuleResponseParams struct {
	// 规则 ID。
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type CreateSecurityIPGroupRequestParams struct {
	// 站点 Id。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// IP 组信息。
	IPGroup *IPGroup `json:"IPGroup,omitnil,omitempty" name:"IPGroup"`
}

type CreateSecurityIPGroupRequest struct {
	*tchttp.BaseRequest
	
	// 站点 Id。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// IP 组信息。
	IPGroup *IPGroup `json:"IPGroup,omitnil,omitempty" name:"IPGroup"`
}

func (r *CreateSecurityIPGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSecurityIPGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "IPGroup")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSecurityIPGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSecurityIPGroupResponseParams struct {
	// IP 组 Id。
	GroupId *int64 `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateSecurityIPGroupResponse struct {
	*tchttp.BaseResponse
	Response *CreateSecurityIPGroupResponseParams `json:"Response"`
}

func (r *CreateSecurityIPGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSecurityIPGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSharedCNAMERequestParams struct {
	// 共享 CNAME 所属站点的 ID。	
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 共享 CNAME 前缀。请输入合法的域名前缀，例如"test-api"、"test-api.com"，限制输入 50 个字符。
	// 
	// 共享 CNAME 完整格式为：`<自定义前缀>+<zoneid中的12位随机字符串>+share.dnse[0-5].com`。
	// 
	// 例如前缀传入 example.com，EO 会为您创建共享 CNAME：example.com.sai2ig51kaa5.share.dnse2.com。
	SharedCNAMEPrefix *string `json:"SharedCNAMEPrefix,omitnil,omitempty" name:"SharedCNAMEPrefix"`

	// 描述。可输入 1-50 个任意字符。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type CreateSharedCNAMERequest struct {
	*tchttp.BaseRequest
	
	// 共享 CNAME 所属站点的 ID。	
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 共享 CNAME 前缀。请输入合法的域名前缀，例如"test-api"、"test-api.com"，限制输入 50 个字符。
	// 
	// 共享 CNAME 完整格式为：`<自定义前缀>+<zoneid中的12位随机字符串>+share.dnse[0-5].com`。
	// 
	// 例如前缀传入 example.com，EO 会为您创建共享 CNAME：example.com.sai2ig51kaa5.share.dnse2.com。
	SharedCNAMEPrefix *string `json:"SharedCNAMEPrefix,omitnil,omitempty" name:"SharedCNAMEPrefix"`

	// 描述。可输入 1-50 个任意字符。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

func (r *CreateSharedCNAMERequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSharedCNAMERequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "SharedCNAMEPrefix")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSharedCNAMERequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSharedCNAMEResponseParams struct {
	// 共享 CNAME。格式为：`<自定义前缀>+<ZoneId中的12位随机字符串>+share.dnse[0-5].com`。
	SharedCNAME *string `json:"SharedCNAME,omitnil,omitempty" name:"SharedCNAME"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateSharedCNAMEResponse struct {
	*tchttp.BaseResponse
	Response *CreateSharedCNAMEResponseParams `json:"Response"`
}

func (r *CreateSharedCNAMEResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSharedCNAMEResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateZoneRequestParams struct {
	// 站点接入类型。该参数取值如下，不填写时默认为 partial：
	// <li>partial：CNAME 接入；</li>
	// <li>full：NS 接入；</li>
	// <li>noDomainAccess：无域名接入；</li>
	// <li>dnsPodAccess：DNSPod 托管接入，该接入模式要求您的域名已托管在 DNSPod 内。</li>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 站点名称。CNAME/NS 接入的时，请传入二级域名（example.com）作为站点名称；无域名接入时，该值请保留为空。
	ZoneName *string `json:"ZoneName,omitnil,omitempty" name:"ZoneName"`

	// Type 取值为 partial/full 时，七层域名的加速区域。以下为该参数取值，不填写时该值默认为 overseas。Type 取值为 noDomainAccess 时该值请保留为空：
	// <li> global: 全球可用区；</li>
	// <li> mainland: 中国大陆可用区；</li>
	// <li> overseas: 全球可用区（不含中国大陆）。</li>
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 待绑定的目标套餐 ID。当您账号下已存在套餐时，可以填写此参数，直接将站点绑定至该套餐。若您当前没有可绑定的套餐时，请前往控制台购买套餐完成站点创建。
	PlanId *string `json:"PlanId,omitnil,omitempty" name:"PlanId"`

	// 同名站点标识。限制输入数字、英文、- 和 _ 组合，长度 20 个字符以内。详情参考 [同名站点标识](https://cloud.tencent.com/document/product/1552/70202)，无此使用场景时，该字段保留为空即可。
	AliasZoneName *string `json:"AliasZoneName,omitnil,omitempty" name:"AliasZoneName"`

	// 标签。该参数用于对站点进行分权限管控、分账。需要先前往 [标签控制台](https://console.cloud.tencent.com/tag/taglist) 创建对应的标签才可以在此处传入对应的标签键和标签值。
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 是否允许重复接入。
	// <li> true：允许重复接入；</li>
	// <li> false：不允许重复接入。</li>不填写使用默认值false。
	//
	// Deprecated: AllowDuplicates is deprecated.
	AllowDuplicates *bool `json:"AllowDuplicates,omitnil,omitempty" name:"AllowDuplicates"`

	// 是否跳过站点现有的DNS记录扫描。默认值：false。
	//
	// Deprecated: JumpStart is deprecated.
	JumpStart *bool `json:"JumpStart,omitnil,omitempty" name:"JumpStart"`
}

type CreateZoneRequest struct {
	*tchttp.BaseRequest
	
	// 站点接入类型。该参数取值如下，不填写时默认为 partial：
	// <li>partial：CNAME 接入；</li>
	// <li>full：NS 接入；</li>
	// <li>noDomainAccess：无域名接入；</li>
	// <li>dnsPodAccess：DNSPod 托管接入，该接入模式要求您的域名已托管在 DNSPod 内。</li>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 站点名称。CNAME/NS 接入的时，请传入二级域名（example.com）作为站点名称；无域名接入时，该值请保留为空。
	ZoneName *string `json:"ZoneName,omitnil,omitempty" name:"ZoneName"`

	// Type 取值为 partial/full 时，七层域名的加速区域。以下为该参数取值，不填写时该值默认为 overseas。Type 取值为 noDomainAccess 时该值请保留为空：
	// <li> global: 全球可用区；</li>
	// <li> mainland: 中国大陆可用区；</li>
	// <li> overseas: 全球可用区（不含中国大陆）。</li>
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 待绑定的目标套餐 ID。当您账号下已存在套餐时，可以填写此参数，直接将站点绑定至该套餐。若您当前没有可绑定的套餐时，请前往控制台购买套餐完成站点创建。
	PlanId *string `json:"PlanId,omitnil,omitempty" name:"PlanId"`

	// 同名站点标识。限制输入数字、英文、- 和 _ 组合，长度 20 个字符以内。详情参考 [同名站点标识](https://cloud.tencent.com/document/product/1552/70202)，无此使用场景时，该字段保留为空即可。
	AliasZoneName *string `json:"AliasZoneName,omitnil,omitempty" name:"AliasZoneName"`

	// 标签。该参数用于对站点进行分权限管控、分账。需要先前往 [标签控制台](https://console.cloud.tencent.com/tag/taglist) 创建对应的标签才可以在此处传入对应的标签键和标签值。
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 是否允许重复接入。
	// <li> true：允许重复接入；</li>
	// <li> false：不允许重复接入。</li>不填写使用默认值false。
	AllowDuplicates *bool `json:"AllowDuplicates,omitnil,omitempty" name:"AllowDuplicates"`

	// 是否跳过站点现有的DNS记录扫描。默认值：false。
	JumpStart *bool `json:"JumpStart,omitnil,omitempty" name:"JumpStart"`
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
	delete(f, "Type")
	delete(f, "ZoneName")
	delete(f, "Area")
	delete(f, "PlanId")
	delete(f, "AliasZoneName")
	delete(f, "Tags")
	delete(f, "AllowDuplicates")
	delete(f, "JumpStart")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateZoneRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateZoneResponseParams struct {
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 站点归属权验证信息。站点完成创建后，您还需要完成归属权校验，站点才能正常服务。
	// 
	// Type = partial 时，您需要参考 [站点/域名归属权验证](https://cloud.tencent.com/document/product/1552/70789) 前往您的域名解析服务商添加 TXT 记录或者前往根域名服务器添加文件，再调用接口 [VerifyOwnership]() 完成验证；
	// 
	// Type = full 时，您需要参考 [修改 DNS 服务器](https://cloud.tencent.com/document/product/1552/90452) 切换 DNS 服务器即可，可通过接口 [VerifyOwnership]() 查询 DNS 是否切换成功；
	// 
	// Type = noDomainAccess 时，该值为空，不需要进行任何操作。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OwnershipVerification *OwnershipVerification `json:"OwnershipVerification,omitnil,omitempty" name:"OwnershipVerification"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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

type CustomEndpoint struct {
	// 实时日志投递的自定义 HTTP 接口地址，暂仅支持 HTTP/HTTPS 协议。
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 填写自定义的 SecretId 用于生成加密签名，如果源站需要鉴权此参数必填。
	AccessId *string `json:"AccessId,omitnil,omitempty" name:"AccessId"`

	// 填写自定义的 SecretKey 用于生成加密签名，如果源站需要鉴权此参数必填。
	AccessKey *string `json:"AccessKey,omitnil,omitempty" name:"AccessKey"`

	// 数据压缩类型，取值有: <li> gzip：使用 gzip 方式压缩。</li>不填表示不启用压缩。
	CompressType *string `json:"CompressType,omitnil,omitempty" name:"CompressType"`

	// POST 请求投递日志时，使用的应用层协议类型，取值有： 
	// <li>http：HTTP 协议；</li>
	// <li>https：HTTPS 协议。</li>如果不填默认根据填写的 URL 地址解析出协议类型。	
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// 投递日志时携带的自定义请求头。若您填写的头部名称为 Content-Type 等 EdgeOne 日志推送默认携带的头部，那么您填写的头部值将覆盖默认值。头部值引用单个变量${batchSize}，以获取每次 POST 请求中包含的日志条数。
	Headers []*Header `json:"Headers,omitnil,omitempty" name:"Headers"`
}

type CustomErrorPage struct {
	// 自定义错误页面 ID。
	PageId *string `json:"PageId,omitnil,omitempty" name:"PageId"`

	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 自定义错误页面名称。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 自定义错误页面类型。
	ContentType *string `json:"ContentType,omitnil,omitempty" name:"ContentType"`

	// 自定义错误页面描述。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 自定义错误页面内容。
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`

	// 自定义错误页面引用。
	References []*ErrorPageReference `json:"References,omitnil,omitempty" name:"References"`
}

type CustomField struct {
	// 从 HTTP 请求和响应中的指定位置提取数据，取值有：
	// <li>ReqHeader：从 HTTP 请求头中提取指定字段值；</li>
	// <li>RspHeader：从 HTTP 响应头中提取指定字段值；</li>
	// <li>Cookie: 从 Cookie 中提取指定字段值。</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 需要提取值的参数名称，例如：Accept-Language。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`

	// 是否投递该字段，不填表示不投递此字段。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Enabled *bool `json:"Enabled,omitnil,omitempty" name:"Enabled"`
}

type DDoS struct {
	// 开关，取值有：
	// <li>on：开启；</li>
	// <li>off：关闭。</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`
}

type DDoSAttackEvent struct {
	// 事件ID。
	EventId *string `json:"EventId,omitnil,omitempty" name:"EventId"`

	// 攻击类型(对应交互事件名称)。
	AttackType *string `json:"AttackType,omitnil,omitempty" name:"AttackType"`

	// 攻击状态。
	AttackStatus *int64 `json:"AttackStatus,omitnil,omitempty" name:"AttackStatus"`

	// 攻击最大带宽。
	AttackMaxBandWidth *int64 `json:"AttackMaxBandWidth,omitnil,omitempty" name:"AttackMaxBandWidth"`

	// 攻击包速率峰值。
	AttackPacketMaxRate *int64 `json:"AttackPacketMaxRate,omitnil,omitempty" name:"AttackPacketMaxRate"`

	// 攻击开始时间，单位为s。
	AttackStartTime *int64 `json:"AttackStartTime,omitnil,omitempty" name:"AttackStartTime"`

	// 攻击结束时间，单位为s。
	AttackEndTime *int64 `json:"AttackEndTime,omitnil,omitempty" name:"AttackEndTime"`

	// DDoS策略组ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PolicyId *int64 `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`

	// 站点ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 攻击事件所属地区，取值有：
	// <li>overseas：全球（除中国大陆地区）数据；</li>
	// <li>mainland：中国大陆地区数据。</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 封禁解封信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DDoSBlockData []*DDoSBlockData `json:"DDoSBlockData,omitnil,omitempty" name:"DDoSBlockData"`
}

type DDoSBlockData struct {
	// 开始时间，采用unix时间戳。
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间，采用unix时间戳, 为0表示还处于封禁中。
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 封禁受影响区域。
	BlockArea *string `json:"BlockArea,omitnil,omitempty" name:"BlockArea"`
}

type DDosProtectionConfig struct {
	// 中国大陆地区独立 DDoS 防护的规格。详情请参考 [独立 DDoS 防护相关费用](https://cloud.tencent.com/document/product/1552/94162)
	// <li>PLATFORM：平台默认防护，即不开启独立 DDoS 防护；</li>
	// <li>BASE30_MAX300：开启独立 DDoS 防护，提供 30 Gbps 保底防护带宽以及 300 Gbps 弹性防护带宽；</li>
	// <li>BASE60_MAX600：开启独立 DDoS 防护，提供 60 Gbps 保底防护带宽以及 600 Gbps 弹性防护带宽。</li>不填写参数时，取默认值 PLATFORM。
	LevelMainland *string `json:"LevelMainland,omitnil,omitempty" name:"LevelMainland"`

	// 中国大陆地区独立 DDoS 防护的弹性防护带宽配置。
	// 仅当开启中国大陆区域独立 DDos 防护时有效（详见 LevelMainland 参数配置），且取值范围有如下限制：
	// <li>开启中国大陆地区独立 DDoS 防护，使用 30 Gbps 保底防护带宽规格时（ LevelMainland 参数值为 BASE30_MAX300 ）：有效取值范围为 30 至 300，单位为 Gbps；</li>
	// <li>开启中国大陆地区独立 DDoS 防护，使用 60 Gbps 保底防护带宽规格时（ LevelMainland 参数值为 BASE60_MAX600 ）：有效取值范围为 60 至 600，单位为 Gbps；</li>
	// <li>使用平台默认防护（ LevelMainland 参数值为 PLATFORM ）：不支持配置，本参数值无效。</li>
	MaxBandwidthMainland *uint64 `json:"MaxBandwidthMainland,omitnil,omitempty" name:"MaxBandwidthMainland"`

	// 全球（除中国大陆以外）地区独立 DDoS 防护的规格。
	// <li>PLATFORM：平台默认防护，即不开启独立 DDoS 防护；</li>
	// <li>ANYCAST300：开启独立 DDoS 防护，提供 300 Gbps 防护带宽；</li>
	// <li>ANYCAST_ALLIN：开启独立 DDoS 防护，使用全部可用防护资源进行防护。</li>不填写参数时，取默认值 PLATFORM。
	LevelOverseas *string `json:"LevelOverseas,omitnil,omitempty" name:"LevelOverseas"`
}

type DefaultServerCertInfo struct {
	// 服务器证书 ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CertId *string `json:"CertId,omitnil,omitempty" name:"CertId"`

	// 证书备注名。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Alias *string `json:"Alias,omitnil,omitempty" name:"Alias"`

	// 证书类型，取值有：
	// <li>default: 默认证书;</li>
	// <li>upload:用户上传;</li>
	// <li>managed:腾讯云托管。</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 证书过期时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExpireTime *string `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`

	// 证书生效时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	EffectiveTime *string `json:"EffectiveTime,omitnil,omitempty" name:"EffectiveTime"`

	// 证书公用名。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CommonName *string `json:"CommonName,omitnil,omitempty" name:"CommonName"`

	// 证书SAN域名。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubjectAltName []*string `json:"SubjectAltName,omitnil,omitempty" name:"SubjectAltName"`

	// 部署状态，取值有：
	// <li>processing: 部署中；</li>
	// <li>deployed: 已部署；</li>
	// <li>failed: 部署失败。</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Status为失败时,此字段返回失败原因。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 证书算法。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SignAlgo *string `json:"SignAlgo,omitnil,omitempty" name:"SignAlgo"`
}

// Predefined struct for user
type DeleteAccelerationDomainsRequestParams struct {
	// 加速域名所属站点ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 需要删除的加速域名ID列表。
	DomainNames []*string `json:"DomainNames,omitnil,omitempty" name:"DomainNames"`

	// 是否强制删除。当域名存在关联资源（如马甲域名、流量调度功能）时，是否强制删除该域名，取值有：
	// <li> true：删除该域名及所有关联资源；</li>
	// <li> false：当该加速域名存在关联资源时，不允许删除。</li>不填写，默认值为：false。
	Force *bool `json:"Force,omitnil,omitempty" name:"Force"`
}

type DeleteAccelerationDomainsRequest struct {
	*tchttp.BaseRequest
	
	// 加速域名所属站点ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 需要删除的加速域名ID列表。
	DomainNames []*string `json:"DomainNames,omitnil,omitempty" name:"DomainNames"`

	// 是否强制删除。当域名存在关联资源（如马甲域名、流量调度功能）时，是否强制删除该域名，取值有：
	// <li> true：删除该域名及所有关联资源；</li>
	// <li> false：当该加速域名存在关联资源时，不允许删除。</li>不填写，默认值为：false。
	Force *bool `json:"Force,omitnil,omitempty" name:"Force"`
}

func (r *DeleteAccelerationDomainsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAccelerationDomainsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "DomainNames")
	delete(f, "Force")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAccelerationDomainsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAccelerationDomainsResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteAccelerationDomainsResponse struct {
	*tchttp.BaseResponse
	Response *DeleteAccelerationDomainsResponseParams `json:"Response"`
}

func (r *DeleteAccelerationDomainsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAccelerationDomainsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAliasDomainRequestParams struct {
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 待删除别称域名名称。如果为空，则不执行删除操作。
	AliasNames []*string `json:"AliasNames,omitnil,omitempty" name:"AliasNames"`
}

type DeleteAliasDomainRequest struct {
	*tchttp.BaseRequest
	
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 待删除别称域名名称。如果为空，则不执行删除操作。
	AliasNames []*string `json:"AliasNames,omitnil,omitempty" name:"AliasNames"`
}

func (r *DeleteAliasDomainRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAliasDomainRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "AliasNames")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAliasDomainRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAliasDomainResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteAliasDomainResponse struct {
	*tchttp.BaseResponse
	Response *DeleteAliasDomainResponseParams `json:"Response"`
}

func (r *DeleteAliasDomainResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAliasDomainResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteApplicationProxyRequestParams struct {
	// 站点ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 代理ID。
	ProxyId *string `json:"ProxyId,omitnil,omitempty" name:"ProxyId"`
}

type DeleteApplicationProxyRequest struct {
	*tchttp.BaseRequest
	
	// 站点ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 代理ID。
	ProxyId *string `json:"ProxyId,omitnil,omitempty" name:"ProxyId"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 代理ID。
	ProxyId *string `json:"ProxyId,omitnil,omitempty" name:"ProxyId"`

	// 规则ID。
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`
}

type DeleteApplicationProxyRuleRequest struct {
	*tchttp.BaseRequest
	
	// 站点ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 代理ID。
	ProxyId *string `json:"ProxyId,omitnil,omitempty" name:"ProxyId"`

	// 规则ID。
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DeleteCustomErrorPageRequestParams struct {
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 自定义页面 ID。
	PageId *string `json:"PageId,omitnil,omitempty" name:"PageId"`
}

type DeleteCustomErrorPageRequest struct {
	*tchttp.BaseRequest
	
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 自定义页面 ID。
	PageId *string `json:"PageId,omitnil,omitempty" name:"PageId"`
}

func (r *DeleteCustomErrorPageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCustomErrorPageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "PageId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteCustomErrorPageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCustomErrorPageResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteCustomErrorPageResponse struct {
	*tchttp.BaseResponse
	Response *DeleteCustomErrorPageResponseParams `json:"Response"`
}

func (r *DeleteCustomErrorPageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCustomErrorPageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteFunctionRequestParams struct {
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 函数 ID。
	FunctionId *string `json:"FunctionId,omitnil,omitempty" name:"FunctionId"`
}

type DeleteFunctionRequest struct {
	*tchttp.BaseRequest
	
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 函数 ID。
	FunctionId *string `json:"FunctionId,omitnil,omitempty" name:"FunctionId"`
}

func (r *DeleteFunctionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteFunctionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "FunctionId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteFunctionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteFunctionResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteFunctionResponse struct {
	*tchttp.BaseResponse
	Response *DeleteFunctionResponseParams `json:"Response"`
}

func (r *DeleteFunctionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteFunctionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteFunctionRulesRequestParams struct {
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 规则 ID 列表。
	RuleIds []*string `json:"RuleIds,omitnil,omitempty" name:"RuleIds"`
}

type DeleteFunctionRulesRequest struct {
	*tchttp.BaseRequest
	
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 规则 ID 列表。
	RuleIds []*string `json:"RuleIds,omitnil,omitempty" name:"RuleIds"`
}

func (r *DeleteFunctionRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteFunctionRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "RuleIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteFunctionRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteFunctionRulesResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteFunctionRulesResponse struct {
	*tchttp.BaseResponse
	Response *DeleteFunctionRulesResponseParams `json:"Response"`
}

func (r *DeleteFunctionRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteFunctionRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteL4ProxyRequestParams struct {
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 四层代理实例 ID。
	ProxyId *string `json:"ProxyId,omitnil,omitempty" name:"ProxyId"`
}

type DeleteL4ProxyRequest struct {
	*tchttp.BaseRequest
	
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 四层代理实例 ID。
	ProxyId *string `json:"ProxyId,omitnil,omitempty" name:"ProxyId"`
}

func (r *DeleteL4ProxyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteL4ProxyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "ProxyId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteL4ProxyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteL4ProxyResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteL4ProxyResponse struct {
	*tchttp.BaseResponse
	Response *DeleteL4ProxyResponseParams `json:"Response"`
}

func (r *DeleteL4ProxyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteL4ProxyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteL4ProxyRulesRequestParams struct {
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 四层代理实例 ID。
	ProxyId *string `json:"ProxyId,omitnil,omitempty" name:"ProxyId"`

	// 转发规则 ID 列表。单次最多支持 200 条转发规则。
	RuleIds []*string `json:"RuleIds,omitnil,omitempty" name:"RuleIds"`
}

type DeleteL4ProxyRulesRequest struct {
	*tchttp.BaseRequest
	
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 四层代理实例 ID。
	ProxyId *string `json:"ProxyId,omitnil,omitempty" name:"ProxyId"`

	// 转发规则 ID 列表。单次最多支持 200 条转发规则。
	RuleIds []*string `json:"RuleIds,omitnil,omitempty" name:"RuleIds"`
}

func (r *DeleteL4ProxyRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteL4ProxyRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "ProxyId")
	delete(f, "RuleIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteL4ProxyRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteL4ProxyRulesResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteL4ProxyRulesResponse struct {
	*tchttp.BaseResponse
	Response *DeleteL4ProxyRulesResponseParams `json:"Response"`
}

func (r *DeleteL4ProxyRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteL4ProxyRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteOriginGroupRequestParams struct {
	// 站点 ID
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 源站组 ID，此参数必填。
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`
}

type DeleteOriginGroupRequest struct {
	*tchttp.BaseRequest
	
	// 站点 ID
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 源站组 ID，此参数必填。
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`
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
	delete(f, "GroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteOriginGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteOriginGroupResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DeleteRealtimeLogDeliveryTaskRequestParams struct {
	// 站点 ID。	
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 实时日志投递任务 ID。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

type DeleteRealtimeLogDeliveryTaskRequest struct {
	*tchttp.BaseRequest
	
	// 站点 ID。	
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 实时日志投递任务 ID。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

func (r *DeleteRealtimeLogDeliveryTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRealtimeLogDeliveryTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteRealtimeLogDeliveryTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRealtimeLogDeliveryTaskResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteRealtimeLogDeliveryTaskResponse struct {
	*tchttp.BaseResponse
	Response *DeleteRealtimeLogDeliveryTaskResponseParams `json:"Response"`
}

func (r *DeleteRealtimeLogDeliveryTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRealtimeLogDeliveryTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRulesRequestParams struct {
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 指定删除的规则 ID 列表。
	RuleIds []*string `json:"RuleIds,omitnil,omitempty" name:"RuleIds"`
}

type DeleteRulesRequest struct {
	*tchttp.BaseRequest
	
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 指定删除的规则 ID 列表。
	RuleIds []*string `json:"RuleIds,omitnil,omitempty" name:"RuleIds"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DeleteSecurityIPGroupRequestParams struct {
	// 站点 Id。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// IP 组 Id。
	GroupId *int64 `json:"GroupId,omitnil,omitempty" name:"GroupId"`
}

type DeleteSecurityIPGroupRequest struct {
	*tchttp.BaseRequest
	
	// 站点 Id。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// IP 组 Id。
	GroupId *int64 `json:"GroupId,omitnil,omitempty" name:"GroupId"`
}

func (r *DeleteSecurityIPGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSecurityIPGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "GroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteSecurityIPGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSecurityIPGroupResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteSecurityIPGroupResponse struct {
	*tchttp.BaseResponse
	Response *DeleteSecurityIPGroupResponseParams `json:"Response"`
}

func (r *DeleteSecurityIPGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSecurityIPGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSharedCNAMERequestParams struct {
	// 共享 CNAME 所属站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 需要删除的共享 CNAME。
	SharedCNAME *string `json:"SharedCNAME,omitnil,omitempty" name:"SharedCNAME"`
}

type DeleteSharedCNAMERequest struct {
	*tchttp.BaseRequest
	
	// 共享 CNAME 所属站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 需要删除的共享 CNAME。
	SharedCNAME *string `json:"SharedCNAME,omitnil,omitempty" name:"SharedCNAME"`
}

func (r *DeleteSharedCNAMERequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSharedCNAMERequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "SharedCNAME")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteSharedCNAMERequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSharedCNAMEResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteSharedCNAMEResponse struct {
	*tchttp.BaseResponse
	Response *DeleteSharedCNAMEResponseParams `json:"Response"`
}

func (r *DeleteSharedCNAMEResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSharedCNAMEResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteZoneRequestParams struct {
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`
}

type DeleteZoneRequest struct {
	*tchttp.BaseRequest
	
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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

type DeliveryCondition struct {
	// 日志过滤条件，详细的过滤条件如下：
	// <li>EdgeResponseStatusCode：按照 EdgeOne 节点响应返回给客户端的状态码进行过滤。<br>   支持运算符：equal、great、less、great_equal、less_equal<br>   取值范围：任意大于等于 0 的整数</li>
	// <li>OriginResponseStatusCode：按照源站响应状态码进行过滤。<br>   支持运算符：equal、great、less、great_equal、less_equal<br>   取值范围：任意大于等于 -1 的整数</li>
	// <li>SecurityAction：按照请求命中安全规则后的最终处置动作进行过滤。<br>   支持运算符：equal<br>   可选项如下：<br>   -：未知/未命中<br>   Monitor：观察<br>   JSChallenge：JavaScript 挑战<br>   Deny：拦截<br>   Allow：放行<br>   BlockIP：IP 封禁<br>   Redirect：重定向<br>   ReturnCustomPage：返回自定义页面<br>   ManagedChallenge：托管挑战<br>   Silence：静默<br>   LongDelay：长时间等待后响应<br>   ShortDelay：短时间等待后响应</li>
	// <li>SecurityModule：按照最终处置请求的安全模块名称进行过滤。<br>   支持运算符：equal<br>   可选项如下：<br>   -：未知/未命中<br>   CustomRule：Web防护 - 自定义规则<br>   RateLimitingCustomRule：Web防护 - 速率限制规则<br>   ManagedRule：Web防护 - 托管规则<br>   L7DDoS：Web防护 - CC攻击防护<br>   BotManagement：Bot管理 - Bot基础管理<br>   BotClientReputation：Bot管理 - 客户端画像分析<br>   BotBehaviorAnalysis：Bot管理 - Bot智能分析<br>   BotCustomRule：Bot管理 - 自定义Bot规则<br>   BotActiveDetection：Bot管理 - 主动特征识别</li>
	Conditions []*QueryCondition `json:"Conditions,omitnil,omitempty" name:"Conditions"`
}

// Predefined struct for user
type DeployConfigGroupVersionRequestParams struct {
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 环境 ID。请填写版本需要发布到的环境 ID。
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 需要发布的版本信息。可以同时变更多个不同配置组的版本，每个配置组一次仅支持变更一个版本。
	ConfigGroupVersionInfos []*ConfigGroupVersionInfo `json:"ConfigGroupVersionInfos,omitnil,omitempty" name:"ConfigGroupVersionInfos"`

	// 变更说明。用于描述此次变更的内容、原因，最大支持 100 个字符。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type DeployConfigGroupVersionRequest struct {
	*tchttp.BaseRequest
	
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 环境 ID。请填写版本需要发布到的环境 ID。
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 需要发布的版本信息。可以同时变更多个不同配置组的版本，每个配置组一次仅支持变更一个版本。
	ConfigGroupVersionInfos []*ConfigGroupVersionInfo `json:"ConfigGroupVersionInfos,omitnil,omitempty" name:"ConfigGroupVersionInfos"`

	// 变更说明。用于描述此次变更的内容、原因，最大支持 100 个字符。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

func (r *DeployConfigGroupVersionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeployConfigGroupVersionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "EnvId")
	delete(f, "ConfigGroupVersionInfos")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeployConfigGroupVersionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeployConfigGroupVersionResponseParams struct {
	// 发布记录 ID。
	RecordId *string `json:"RecordId,omitnil,omitempty" name:"RecordId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeployConfigGroupVersionResponse struct {
	*tchttp.BaseResponse
	Response *DeployConfigGroupVersionResponseParams `json:"Response"`
}

func (r *DeployConfigGroupVersionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeployConfigGroupVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeployRecord struct {
	// 发布版本的详细信息。
	ConfigGroupVersionInfos []*ConfigGroupVersionInfo `json:"ConfigGroupVersionInfos,omitnil,omitempty" name:"ConfigGroupVersionInfos"`

	// 发布时间。时间为世界标准时间（UTC）， 遵循 ISO 8601 标准的日期和时间格式。
	DeployTime *string `json:"DeployTime,omitnil,omitempty" name:"DeployTime"`

	// 发布状态，取值有：
	// <li> deploying ：发布中；</li>
	// <li>failure ：发布失败；</li>
	// <li>success： 发布成功。</li>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 发布结果信息。
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 发布记录 ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RecordId *string `json:"RecordId,omitnil,omitempty" name:"RecordId"`

	// 变更说明。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

// Predefined struct for user
type DescribeAccelerationDomainsRequestParams struct {
	// 加速域名所属站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 分页查询偏移量，默认为 0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页查询限制数目，默认值：20，上限：200。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 过滤条件，Filters.Values 的上限为 20。该参数不填写时，返回当前 zone-id 下所有域名信息。详细的过滤条件如下：
	// <li>domain-name：按照加速域名进行过滤；</li>
	// <li>origin-type：按照源站类型进行过滤；</li>
	// <li>origin：按照主源站地址进行过滤；</li>
	// <li>backup-origin： 按照备用源站地址进行过滤；</li>
	// <li>domain-cname：按照 CNAME 进行过滤；</li>
	// <li>share-cname：按照共享 CNAME 进行过滤；</li>
	Filters []*AdvancedFilter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 可根据该字段对返回结果进行排序，取值有：
	// <li>created_on：加速域名创建时间；</li>
	// <li>domain-name：加速域名。</li>不填写时，默认对返回结果按照 domain-name 排序。
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// 排序方向，如果是字段值为数字，则根据数字大小排序；如果字段值为文本，则根据 ascill 码的大小排序。取值有：
	// <li>asc：升序排列；</li>
	// <li>desc：降序排列。</li>不填写使用默认值 asc。
	Direction *string `json:"Direction,omitnil,omitempty" name:"Direction"`

	// 匹配方式，取值有：
	// <li>all：返回匹配所有查询条件的加速域名；</li>
	// <li>any：返回匹配任意一个查询条件的加速域名。</li>不填写时默认值为 all。
	Match *string `json:"Match,omitnil,omitempty" name:"Match"`
}

type DescribeAccelerationDomainsRequest struct {
	*tchttp.BaseRequest
	
	// 加速域名所属站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 分页查询偏移量，默认为 0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页查询限制数目，默认值：20，上限：200。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 过滤条件，Filters.Values 的上限为 20。该参数不填写时，返回当前 zone-id 下所有域名信息。详细的过滤条件如下：
	// <li>domain-name：按照加速域名进行过滤；</li>
	// <li>origin-type：按照源站类型进行过滤；</li>
	// <li>origin：按照主源站地址进行过滤；</li>
	// <li>backup-origin： 按照备用源站地址进行过滤；</li>
	// <li>domain-cname：按照 CNAME 进行过滤；</li>
	// <li>share-cname：按照共享 CNAME 进行过滤；</li>
	Filters []*AdvancedFilter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 可根据该字段对返回结果进行排序，取值有：
	// <li>created_on：加速域名创建时间；</li>
	// <li>domain-name：加速域名。</li>不填写时，默认对返回结果按照 domain-name 排序。
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// 排序方向，如果是字段值为数字，则根据数字大小排序；如果字段值为文本，则根据 ascill 码的大小排序。取值有：
	// <li>asc：升序排列；</li>
	// <li>desc：降序排列。</li>不填写使用默认值 asc。
	Direction *string `json:"Direction,omitnil,omitempty" name:"Direction"`

	// 匹配方式，取值有：
	// <li>all：返回匹配所有查询条件的加速域名；</li>
	// <li>any：返回匹配任意一个查询条件的加速域名。</li>不填写时默认值为 all。
	Match *string `json:"Match,omitnil,omitempty" name:"Match"`
}

func (r *DescribeAccelerationDomainsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccelerationDomainsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	delete(f, "Order")
	delete(f, "Direction")
	delete(f, "Match")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAccelerationDomainsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAccelerationDomainsResponseParams struct {
	// 符合查询条件的加速域名个数。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 符合查询条件的所有加速域名的信息。
	AccelerationDomains []*AccelerationDomain `json:"AccelerationDomains,omitnil,omitempty" name:"AccelerationDomains"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAccelerationDomainsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAccelerationDomainsResponseParams `json:"Response"`
}

func (r *DescribeAccelerationDomainsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccelerationDomainsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAliasDomainsRequestParams struct {
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 分页查询偏移量。默认值：0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页查询限制数目。默认值：20，最大值：1000。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 过滤条件，Filters.Values的上限为20。详细的过滤条件如下：
	// <li>target-name：目标域名名称进行过滤；</li>
	// <li>alias-name：按照别称域名名称进行过滤。</li>模糊查询时仅支持过滤字段名为 alias-name。
	Filters []*AdvancedFilter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeAliasDomainsRequest struct {
	*tchttp.BaseRequest
	
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 分页查询偏移量。默认值：0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页查询限制数目。默认值：20，最大值：1000。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 过滤条件，Filters.Values的上限为20。详细的过滤条件如下：
	// <li>target-name：目标域名名称进行过滤；</li>
	// <li>alias-name：按照别称域名名称进行过滤。</li>模糊查询时仅支持过滤字段名为 alias-name。
	Filters []*AdvancedFilter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribeAliasDomainsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAliasDomainsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAliasDomainsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAliasDomainsResponseParams struct {
	// 符合条件的别称域名个数。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 别称域名详细信息列表。
	AliasDomains []*AliasDomain `json:"AliasDomains,omitnil,omitempty" name:"AliasDomains"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAliasDomainsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAliasDomainsResponseParams `json:"Response"`
}

func (r *DescribeAliasDomainsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAliasDomainsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApplicationProxiesRequestParams struct {
	// 站点ID。该参数必填。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 过滤条件，Filters.Values的上限为20。详细的过滤条件如下：<li>proxy-id<br>   按照【<strong>代理ID</strong>】进行过滤。代理ID形如：proxy-ev2sawbwfd。<br>   类型：String<br>   必选：否</li><li>zone-id<br>   按照【<strong>站点ID</strong>】进行过滤。站点ID形如：zone-vawer2vadg。<br>   类型：String<br>   必选：否</li><li>rule-tag<br>   按照【<strong>规则标签</strong>】对应用代理下的规则进行过滤。规则标签形如：rule-service-1。<br>   类型：String<br>   必选：否</li>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 分页查询偏移量。默认为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页查询限制数目。默认值：20，最大值：1000。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeApplicationProxiesRequest struct {
	*tchttp.BaseRequest
	
	// 站点ID。该参数必填。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 过滤条件，Filters.Values的上限为20。详细的过滤条件如下：<li>proxy-id<br>   按照【<strong>代理ID</strong>】进行过滤。代理ID形如：proxy-ev2sawbwfd。<br>   类型：String<br>   必选：否</li><li>zone-id<br>   按照【<strong>站点ID</strong>】进行过滤。站点ID形如：zone-vawer2vadg。<br>   类型：String<br>   必选：否</li><li>rule-tag<br>   按照【<strong>规则标签</strong>】对应用代理下的规则进行过滤。规则标签形如：rule-service-1。<br>   类型：String<br>   必选：否</li>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 分页查询偏移量。默认为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页查询限制数目。默认值：20，最大值：1000。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
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
	delete(f, "ZoneId")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeApplicationProxiesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApplicationProxiesResponseParams struct {
	// 应用代理列表。
	ApplicationProxies []*ApplicationProxy `json:"ApplicationProxies,omitnil,omitempty" name:"ApplicationProxies"`

	// 记录总数。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	PlanInfo []*PlanInfo `json:"PlanInfo,omitnil,omitempty" name:"PlanInfo"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 站点 ID 集合，此参数必填。
	ZoneIds []*string `json:"ZoneIds,omitnil,omitempty" name:"ZoneIds"`

	// 指标列表，取值有：
	// <li>acc_flux: 内容加速流量，单位为 Byte；</li>
	// <li>smt_flux: 智能加速流量，单位为 Byte；</li>
	// <li>l4_flux: 四层加速流量，单位为 Byte；</li>
	// <li>sec_flux: 独立防护流量，单位为 Byte；</li>
	// <li>zxctg_flux: 中国大陆网络优化流量，单位为 Byte；</li>
	// <li>acc_bandwidth: 内容加速带宽，单位为 bps；</li>
	// <li>smt_bandwidth: 智能加速带宽，单位为 bps；</li>
	// <li>l4_bandwidth: 四层加速带宽，单位为 bps；</li>
	// <li>sec_bandwidth: 独立防护带宽，单位为 bps；</li>
	// <li>zxctg_bandwidth: 中国大陆网络优化带宽，单位为 bps；</li>
	// <li>sec_request_clean: HTTP/HTTPS 请求，单位为次；</li>
	// <li>smt_request_clean: 智能加速请求，单位为次；</li>
	// <li>quic_request: QUIC 请求，单位为次；</li>
	// <li>bot_request_clean: Bot 请求，单位为次；</li>
	// <li>cls_count: 实时日志推送条数，单位为条；</li>
	// <li>ddos_bandwidth: 弹性 DDoS 防护带宽，单位为 bps。</li>
	MetricName *string `json:"MetricName,omitnil,omitempty" name:"MetricName"`

	// 查询时间粒度，取值有：
	// <li>5min：5 分钟粒度；</li>
	// <li>hour：1 小时粒度；</li>
	// <li>day：1 天粒度。</li>
	Interval *string `json:"Interval,omitnil,omitempty" name:"Interval"`

	// 过滤条件，详细的过滤条件取值如下：
	// <li>host：按照域名进行过滤。示例值：test.example.com。<br></li>
	// <li>proxy-id：按照四层代理实例 ID 进行过滤。示例值：sid-2rugn89bkla9。<br></li>
	// <li>region-id：按照计费大区进行过滤。可选项如下：<br>  CH：中国大陆境内<br>  AF：非洲<br>  AS1：亚太一区<br>  AS2：亚太二区<br>  AS3：亚太三区<br>  EU：欧洲<br>  MidEast：中东<br>  NA：北美<br>  SA：南美</li>
	Filters []*BillingDataFilter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeBillingDataRequest struct {
	*tchttp.BaseRequest
	
	// 起始时间。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 站点 ID 集合，此参数必填。
	ZoneIds []*string `json:"ZoneIds,omitnil,omitempty" name:"ZoneIds"`

	// 指标列表，取值有：
	// <li>acc_flux: 内容加速流量，单位为 Byte；</li>
	// <li>smt_flux: 智能加速流量，单位为 Byte；</li>
	// <li>l4_flux: 四层加速流量，单位为 Byte；</li>
	// <li>sec_flux: 独立防护流量，单位为 Byte；</li>
	// <li>zxctg_flux: 中国大陆网络优化流量，单位为 Byte；</li>
	// <li>acc_bandwidth: 内容加速带宽，单位为 bps；</li>
	// <li>smt_bandwidth: 智能加速带宽，单位为 bps；</li>
	// <li>l4_bandwidth: 四层加速带宽，单位为 bps；</li>
	// <li>sec_bandwidth: 独立防护带宽，单位为 bps；</li>
	// <li>zxctg_bandwidth: 中国大陆网络优化带宽，单位为 bps；</li>
	// <li>sec_request_clean: HTTP/HTTPS 请求，单位为次；</li>
	// <li>smt_request_clean: 智能加速请求，单位为次；</li>
	// <li>quic_request: QUIC 请求，单位为次；</li>
	// <li>bot_request_clean: Bot 请求，单位为次；</li>
	// <li>cls_count: 实时日志推送条数，单位为条；</li>
	// <li>ddos_bandwidth: 弹性 DDoS 防护带宽，单位为 bps。</li>
	MetricName *string `json:"MetricName,omitnil,omitempty" name:"MetricName"`

	// 查询时间粒度，取值有：
	// <li>5min：5 分钟粒度；</li>
	// <li>hour：1 小时粒度；</li>
	// <li>day：1 天粒度。</li>
	Interval *string `json:"Interval,omitnil,omitempty" name:"Interval"`

	// 过滤条件，详细的过滤条件取值如下：
	// <li>host：按照域名进行过滤。示例值：test.example.com。<br></li>
	// <li>proxy-id：按照四层代理实例 ID 进行过滤。示例值：sid-2rugn89bkla9。<br></li>
	// <li>region-id：按照计费大区进行过滤。可选项如下：<br>  CH：中国大陆境内<br>  AF：非洲<br>  AS1：亚太一区<br>  AS2：亚太二区<br>  AS3：亚太三区<br>  EU：欧洲<br>  MidEast：中东<br>  NA：北美<br>  SA：南美</li>
	Filters []*BillingDataFilter `json:"Filters,omitnil,omitempty" name:"Filters"`
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
	delete(f, "ZoneIds")
	delete(f, "MetricName")
	delete(f, "Interval")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBillingDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBillingDataResponseParams struct {
	// 数据点列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*BillingData `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DescribeConfigGroupVersionDetailRequestParams struct {
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 版本 ID。
	VersionId *string `json:"VersionId,omitnil,omitempty" name:"VersionId"`
}

type DescribeConfigGroupVersionDetailRequest struct {
	*tchttp.BaseRequest
	
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 版本 ID。
	VersionId *string `json:"VersionId,omitnil,omitempty" name:"VersionId"`
}

func (r *DescribeConfigGroupVersionDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeConfigGroupVersionDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "VersionId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeConfigGroupVersionDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeConfigGroupVersionDetailResponseParams struct {
	// 版本信息。
	ConfigGroupVersionInfo *ConfigGroupVersionInfo `json:"ConfigGroupVersionInfo,omitnil,omitempty" name:"ConfigGroupVersionInfo"`

	// 版本文件的内容。以 JSON 格式返回。
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeConfigGroupVersionDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeConfigGroupVersionDetailResponseParams `json:"Response"`
}

func (r *DescribeConfigGroupVersionDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeConfigGroupVersionDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeConfigGroupVersionsRequestParams struct {
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 配置组 ID。
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 过滤条件，Filters.Values 的上限为 20，该参数不填写时，返回所选配置组下的所有版本信息。详细的过滤条件如下：
	// <li>version-id：按照版本 ID 进行过滤；</li>
	Filters []*AdvancedFilter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 分页查询偏移量。默认值为 0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页查询限制数目。默认值为 20，最大值为 100。 
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeConfigGroupVersionsRequest struct {
	*tchttp.BaseRequest
	
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 配置组 ID。
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 过滤条件，Filters.Values 的上限为 20，该参数不填写时，返回所选配置组下的所有版本信息。详细的过滤条件如下：
	// <li>version-id：按照版本 ID 进行过滤；</li>
	Filters []*AdvancedFilter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 分页查询偏移量。默认值为 0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页查询限制数目。默认值为 20，最大值为 100。 
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeConfigGroupVersionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeConfigGroupVersionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "GroupId")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeConfigGroupVersionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeConfigGroupVersionsResponseParams struct {
	// 版本总数。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 版本信息列表。
	ConfigGroupVersionInfos []*ConfigGroupVersionInfo `json:"ConfigGroupVersionInfos,omitnil,omitempty" name:"ConfigGroupVersionInfos"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeConfigGroupVersionsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeConfigGroupVersionsResponseParams `json:"Response"`
}

func (r *DescribeConfigGroupVersionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeConfigGroupVersionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeContentQuotaRequestParams struct {
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`
}

type DescribeContentQuotaRequest struct {
	*tchttp.BaseRequest
	
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`
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
	PurgeQuota []*Quota `json:"PurgeQuota,omitnil,omitempty" name:"PurgeQuota"`

	// 预热相关配额。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PrefetchQuota []*Quota `json:"PrefetchQuota,omitnil,omitempty" name:"PrefetchQuota"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DescribeCustomErrorPagesRequestParams struct {
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 过滤条件，Filters.Values 的上限为20，详细的过滤条件Name值如下：
	// <li>page-id： 按照页面 ID 进行过滤；</li>
	// <li>name： 按照页面名称进行过滤；</li>
	// <li>description：按照页面描述过滤；</li>
	// <li>content-type：按照页面类型过滤。</li>
	Filters []*AdvancedFilter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 分页查询偏移量。默认值：0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页查询限制数目。默认值：20，最大值：1000。 
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeCustomErrorPagesRequest struct {
	*tchttp.BaseRequest
	
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 过滤条件，Filters.Values 的上限为20，详细的过滤条件Name值如下：
	// <li>page-id： 按照页面 ID 进行过滤；</li>
	// <li>name： 按照页面名称进行过滤；</li>
	// <li>description：按照页面描述过滤；</li>
	// <li>content-type：按照页面类型过滤。</li>
	Filters []*AdvancedFilter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 分页查询偏移量。默认值：0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页查询限制数目。默认值：20，最大值：1000。 
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeCustomErrorPagesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCustomErrorPagesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCustomErrorPagesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCustomErrorPagesResponseParams struct {
	// 自定义错误页面总数。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 自定义错误页面数据列表。
	ErrorPages []*CustomErrorPage `json:"ErrorPages,omitnil,omitempty" name:"ErrorPages"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCustomErrorPagesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCustomErrorPagesResponseParams `json:"Response"`
}

func (r *DescribeCustomErrorPagesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCustomErrorPagesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDDoSAttackDataRequestParams struct {
	// 开始时间。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 统计指标列表，取值有：
	// <li>ddos_attackMaxBandwidth：攻击带宽峰值；</li>
	// <li>ddos_attackMaxPackageRate：攻击包速率峰值 ；</li>
	// <li>ddos_attackBandwidth：攻击带宽曲线；</li>
	// <li>ddos_attackPackageRate：攻击包速率曲线。</li>
	MetricNames []*string `json:"MetricNames,omitnil,omitempty" name:"MetricNames"`

	// 站点集合，此参数必填。
	ZoneIds []*string `json:"ZoneIds,omitnil,omitempty" name:"ZoneIds"`

	// DDoS策略组ID列表，不填默认选择全部策略ID。
	PolicyIds []*int64 `json:"PolicyIds,omitnil,omitempty" name:"PolicyIds"`

	// 查询时间粒度，取值有：
	// <li>min：1分钟；</li>
	// <li>5min：5分钟；</li>
	// <li>hour：1小时；</li>
	// <li>day：1天。</li>不填将根据开始时间与结束时间的间隔自动推算粒度，具体为：1小时范围内以min粒度查询，2天范围内以5min粒度查询，7天范围内以hour粒度查询，超过7天以day粒度查询。
	Interval *string `json:"Interval,omitnil,omitempty" name:"Interval"`

	// 数据归属地区，取值有：
	// <li>overseas：全球（除中国大陆地区）数据；</li>
	// <li>mainland：中国大陆地区数据；</li>
	// <li>global：全球数据。</li>不填默认取值为global。
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`
}

type DescribeDDoSAttackDataRequest struct {
	*tchttp.BaseRequest
	
	// 开始时间。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 统计指标列表，取值有：
	// <li>ddos_attackMaxBandwidth：攻击带宽峰值；</li>
	// <li>ddos_attackMaxPackageRate：攻击包速率峰值 ；</li>
	// <li>ddos_attackBandwidth：攻击带宽曲线；</li>
	// <li>ddos_attackPackageRate：攻击包速率曲线。</li>
	MetricNames []*string `json:"MetricNames,omitnil,omitempty" name:"MetricNames"`

	// 站点集合，此参数必填。
	ZoneIds []*string `json:"ZoneIds,omitnil,omitempty" name:"ZoneIds"`

	// DDoS策略组ID列表，不填默认选择全部策略ID。
	PolicyIds []*int64 `json:"PolicyIds,omitnil,omitempty" name:"PolicyIds"`

	// 查询时间粒度，取值有：
	// <li>min：1分钟；</li>
	// <li>5min：5分钟；</li>
	// <li>hour：1小时；</li>
	// <li>day：1天。</li>不填将根据开始时间与结束时间的间隔自动推算粒度，具体为：1小时范围内以min粒度查询，2天范围内以5min粒度查询，7天范围内以hour粒度查询，超过7天以day粒度查询。
	Interval *string `json:"Interval,omitnil,omitempty" name:"Interval"`

	// 数据归属地区，取值有：
	// <li>overseas：全球（除中国大陆地区）数据；</li>
	// <li>mainland：中国大陆地区数据；</li>
	// <li>global：全球数据。</li>不填默认取值为global。
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`
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
	delete(f, "ZoneIds")
	delete(f, "PolicyIds")
	delete(f, "Interval")
	delete(f, "Area")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDDoSAttackDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDDoSAttackDataResponseParams struct {
	// 查询结果的总条数。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// DDoS攻击数据内容列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*SecEntry `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DescribeDDoSAttackEventRequestParams struct {
	// 开始时间，时间范围为 30 天。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间，时间范围为 30 天。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// ddos策略组集合，不填默认选择全部策略。
	PolicyIds []*int64 `json:"PolicyIds,omitnil,omitempty" name:"PolicyIds"`

	// 站点集合，此参数必填。
	ZoneIds []*string `json:"ZoneIds,omitnil,omitempty" name:"ZoneIds"`

	// 分页查询的限制数目，默认值为20，最大查询条目为1000。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页的偏移量，默认值为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 展示攻击详情的参数，若填false，默认只返回攻击次数，不返回攻击详情；若填true，返回攻击详情。
	ShowDetail *bool `json:"ShowDetail,omitnil,omitempty" name:"ShowDetail"`

	// 数据归属地区，取值有：
	// <li>overseas：全球（除中国大陆地区）数据；</li>
	// <li>mainland：中国大陆地区数据；</li>
	// <li>global：全球数据；</li>不填默认取值为global。
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 排序字段，取值有：
	// <li>MaxBandWidth：带宽峰值；</li>
	// <li>AttackStartTime：攻击开始时间。</li>不填默认值为：AttackStartTime。
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// 排序方式，取值有：
	// <li>asc：升序方式；</li>
	// <li>desc：降序方式。</li>不填默认值为：desc。
	OrderType *string `json:"OrderType,omitnil,omitempty" name:"OrderType"`
}

type DescribeDDoSAttackEventRequest struct {
	*tchttp.BaseRequest
	
	// 开始时间，时间范围为 30 天。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间，时间范围为 30 天。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// ddos策略组集合，不填默认选择全部策略。
	PolicyIds []*int64 `json:"PolicyIds,omitnil,omitempty" name:"PolicyIds"`

	// 站点集合，此参数必填。
	ZoneIds []*string `json:"ZoneIds,omitnil,omitempty" name:"ZoneIds"`

	// 分页查询的限制数目，默认值为20，最大查询条目为1000。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页的偏移量，默认值为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 展示攻击详情的参数，若填false，默认只返回攻击次数，不返回攻击详情；若填true，返回攻击详情。
	ShowDetail *bool `json:"ShowDetail,omitnil,omitempty" name:"ShowDetail"`

	// 数据归属地区，取值有：
	// <li>overseas：全球（除中国大陆地区）数据；</li>
	// <li>mainland：中国大陆地区数据；</li>
	// <li>global：全球数据；</li>不填默认取值为global。
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 排序字段，取值有：
	// <li>MaxBandWidth：带宽峰值；</li>
	// <li>AttackStartTime：攻击开始时间。</li>不填默认值为：AttackStartTime。
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// 排序方式，取值有：
	// <li>asc：升序方式；</li>
	// <li>desc：降序方式。</li>不填默认值为：desc。
	OrderType *string `json:"OrderType,omitnil,omitempty" name:"OrderType"`
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
	delete(f, "PolicyIds")
	delete(f, "ZoneIds")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "ShowDetail")
	delete(f, "Area")
	delete(f, "OrderBy")
	delete(f, "OrderType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDDoSAttackEventRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDDoSAttackEventResponseParams struct {
	// DDOS攻击事件数据列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*DDoSAttackEvent `json:"Data,omitnil,omitempty" name:"Data"`

	// 查询结果的总条数。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DescribeDDoSAttackTopDataRequestParams struct {
	// 开始时间。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 查询的统计指标，取值有：
	// <li>ddos_attackFlux_protocol：按各协议的攻击流量排行；</li>
	// <li>ddos_attackPackageNum_protocol：按各协议的攻击包量排行；</li>
	// <li>ddos_attackNum_attackType：按各攻击类型的攻击数量排行；</li>
	// <li>ddos_attackNum_sregion：按攻击源地区的攻击数量排行；</li>
	// <li>ddos_attackFlux_sip：按攻击源IP的攻击数量排行；</li>
	// <li>ddos_attackFlux_sregion：按攻击源地区的攻击数量排行。</li>
	MetricName *string `json:"MetricName,omitnil,omitempty" name:"MetricName"`

	// 站点ID集合，此参数必填。
	ZoneIds []*string `json:"ZoneIds,omitnil,omitempty" name:"ZoneIds"`

	// DDoS策略组ID集合，不填默认选择全部策略ID。
	PolicyIds []*int64 `json:"PolicyIds,omitnil,omitempty" name:"PolicyIds"`

	// 攻击类型，取值有：
	// <li>flood：洪泛攻击；</li>
	// <li>icmpFlood：icmp洪泛攻击；</li>
	// <li>all：所有的攻击类型。</li>不填默认为all，表示查询全部攻击类型。
	AttackType *string `json:"AttackType,omitnil,omitempty" name:"AttackType"`

	// 协议类型，取值有：
	// <li>tcp：tcp协议；</li>
	// <li>udp：udp协议；</li>
	// <li>all：所有的协议类型。</li>不填默认为all，表示查询所有协议。
	ProtocolType *string `json:"ProtocolType,omitnil,omitempty" name:"ProtocolType"`

	// 端口号。
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`

	// 查询前多少个数据，不填默认默认为10， 表示查询前top 10的数据。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 数据归属地区，取值有：
	// <li>overseas：全球（除中国大陆地区）数据；</li>
	// <li>mainland：中国大陆地区数据。</li>不填将根据用户所在地智能选择地区。
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`
}

type DescribeDDoSAttackTopDataRequest struct {
	*tchttp.BaseRequest
	
	// 开始时间。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 查询的统计指标，取值有：
	// <li>ddos_attackFlux_protocol：按各协议的攻击流量排行；</li>
	// <li>ddos_attackPackageNum_protocol：按各协议的攻击包量排行；</li>
	// <li>ddos_attackNum_attackType：按各攻击类型的攻击数量排行；</li>
	// <li>ddos_attackNum_sregion：按攻击源地区的攻击数量排行；</li>
	// <li>ddos_attackFlux_sip：按攻击源IP的攻击数量排行；</li>
	// <li>ddos_attackFlux_sregion：按攻击源地区的攻击数量排行。</li>
	MetricName *string `json:"MetricName,omitnil,omitempty" name:"MetricName"`

	// 站点ID集合，此参数必填。
	ZoneIds []*string `json:"ZoneIds,omitnil,omitempty" name:"ZoneIds"`

	// DDoS策略组ID集合，不填默认选择全部策略ID。
	PolicyIds []*int64 `json:"PolicyIds,omitnil,omitempty" name:"PolicyIds"`

	// 攻击类型，取值有：
	// <li>flood：洪泛攻击；</li>
	// <li>icmpFlood：icmp洪泛攻击；</li>
	// <li>all：所有的攻击类型。</li>不填默认为all，表示查询全部攻击类型。
	AttackType *string `json:"AttackType,omitnil,omitempty" name:"AttackType"`

	// 协议类型，取值有：
	// <li>tcp：tcp协议；</li>
	// <li>udp：udp协议；</li>
	// <li>all：所有的协议类型。</li>不填默认为all，表示查询所有协议。
	ProtocolType *string `json:"ProtocolType,omitnil,omitempty" name:"ProtocolType"`

	// 端口号。
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`

	// 查询前多少个数据，不填默认默认为10， 表示查询前top 10的数据。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 数据归属地区，取值有：
	// <li>overseas：全球（除中国大陆地区）数据；</li>
	// <li>mainland：中国大陆地区数据。</li>不填将根据用户所在地智能选择地区。
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`
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
	Data []*TopEntry `json:"Data,omitnil,omitempty" name:"Data"`

	// 查询结果的总条数。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DescribeDefaultCertificatesRequestParams struct {
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 过滤条件，Filters.Values的上限为5。详细的过滤条件如下：
	// <li>zone-id<br>   按照【<strong>站点ID</strong>】进行过滤。站点ID形如：zone-xxx。<br>   类型：String<br>   必选：是 </li>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 分页查询偏移量。默认值：0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页查询限制数目。默认值：20，最大值：100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeDefaultCertificatesRequest struct {
	*tchttp.BaseRequest
	
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 过滤条件，Filters.Values的上限为5。详细的过滤条件如下：
	// <li>zone-id<br>   按照【<strong>站点ID</strong>】进行过滤。站点ID形如：zone-xxx。<br>   类型：String<br>   必选：是 </li>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 分页查询偏移量。默认值：0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页查询限制数目。默认值：20，最大值：100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
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
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 默认证书列表。
	DefaultServerCertInfo []*DefaultServerCertInfo `json:"DefaultServerCertInfo,omitnil,omitempty" name:"DefaultServerCertInfo"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DescribeDeployHistoryRequestParams struct {
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 环境 ID。
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 过滤条件，Filters.Values 的上限为 20，详细的过滤条件如下：
	// <li>record-id：按照发布记录 ID 进行过滤进行过滤。</li>
	Filters []*AdvancedFilter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeDeployHistoryRequest struct {
	*tchttp.BaseRequest
	
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 环境 ID。
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 过滤条件，Filters.Values 的上限为 20，详细的过滤条件如下：
	// <li>record-id：按照发布记录 ID 进行过滤进行过滤。</li>
	Filters []*AdvancedFilter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribeDeployHistoryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeployHistoryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "EnvId")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDeployHistoryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDeployHistoryResponseParams struct {
	// 发布记录总数。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 发布记录详情。
	Records []*DeployRecord `json:"Records,omitnil,omitempty" name:"Records"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDeployHistoryResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDeployHistoryResponseParams `json:"Response"`
}

func (r *DescribeDeployHistoryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeployHistoryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEnvironmentsRequestParams struct {
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`
}

type DescribeEnvironmentsRequest struct {
	*tchttp.BaseRequest
	
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`
}

func (r *DescribeEnvironmentsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEnvironmentsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeEnvironmentsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEnvironmentsResponseParams struct {
	// 环境总数。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 环境列表。
	EnvInfos []*EnvInfo `json:"EnvInfos,omitnil,omitempty" name:"EnvInfos"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeEnvironmentsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeEnvironmentsResponseParams `json:"Response"`
}

func (r *DescribeEnvironmentsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEnvironmentsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFunctionRulesRequestParams struct {
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 过滤条件列表，多个条件为且关系，Filters.Values 的上限为 20。详细的过滤条件如下：
	// <li>rule-id：按照【规则 ID】进行精确匹配。</li>
	// <li>function-id：按照【函数 ID】进行精确匹配。</li>
	// <li>remark：按照【规则描述】进行模糊匹配。</li>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeFunctionRulesRequest struct {
	*tchttp.BaseRequest
	
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 过滤条件列表，多个条件为且关系，Filters.Values 的上限为 20。详细的过滤条件如下：
	// <li>rule-id：按照【规则 ID】进行精确匹配。</li>
	// <li>function-id：按照【函数 ID】进行精确匹配。</li>
	// <li>remark：按照【规则描述】进行模糊匹配。</li>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribeFunctionRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFunctionRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeFunctionRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFunctionRulesResponseParams struct {
	// 规则详情列表。
	FunctionRules []*FunctionRule `json:"FunctionRules,omitnil,omitempty" name:"FunctionRules"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeFunctionRulesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeFunctionRulesResponseParams `json:"Response"`
}

func (r *DescribeFunctionRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFunctionRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFunctionRuntimeEnvironmentRequestParams struct {
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 函数 ID。
	FunctionId *string `json:"FunctionId,omitnil,omitempty" name:"FunctionId"`
}

type DescribeFunctionRuntimeEnvironmentRequest struct {
	*tchttp.BaseRequest
	
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 函数 ID。
	FunctionId *string `json:"FunctionId,omitnil,omitempty" name:"FunctionId"`
}

func (r *DescribeFunctionRuntimeEnvironmentRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFunctionRuntimeEnvironmentRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "FunctionId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeFunctionRuntimeEnvironmentRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFunctionRuntimeEnvironmentResponseParams struct {
	// 环境变量列表。
	EnvironmentVariables []*FunctionEnvironmentVariable `json:"EnvironmentVariables,omitnil,omitempty" name:"EnvironmentVariables"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeFunctionRuntimeEnvironmentResponse struct {
	*tchttp.BaseResponse
	Response *DescribeFunctionRuntimeEnvironmentResponseParams `json:"Response"`
}

func (r *DescribeFunctionRuntimeEnvironmentResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFunctionRuntimeEnvironmentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFunctionsRequestParams struct {
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 按照函数 ID 列表过滤。
	FunctionIds []*string `json:"FunctionIds,omitnil,omitempty" name:"FunctionIds"`

	// 过滤条件列表，多个条件为且关系，Filters.Values 的上限为 20。详细的过滤条件如下：
	// <li>name：按照【函数名称】进行模糊匹配。</li>
	// <li>remark：按照【函数描述】进行模糊匹配。</li>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 分页查询偏移量。默认值：0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页查询限制数目。默认值：20，最大值：200。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeFunctionsRequest struct {
	*tchttp.BaseRequest
	
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 按照函数 ID 列表过滤。
	FunctionIds []*string `json:"FunctionIds,omitnil,omitempty" name:"FunctionIds"`

	// 过滤条件列表，多个条件为且关系，Filters.Values 的上限为 20。详细的过滤条件如下：
	// <li>name：按照【函数名称】进行模糊匹配。</li>
	// <li>remark：按照【函数描述】进行模糊匹配。</li>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 分页查询偏移量。默认值：0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页查询限制数目。默认值：20，最大值：200。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeFunctionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFunctionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "FunctionIds")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeFunctionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFunctionsResponseParams struct {
	// 符合查询条件的函数总数。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 符合查询条件的所有函数信息。
	Functions []*Function `json:"Functions,omitnil,omitempty" name:"Functions"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeFunctionsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeFunctionsResponseParams `json:"Response"`
}

func (r *DescribeFunctionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFunctionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeHostsSettingRequestParams struct {
	// 站点ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 分页查询偏移量。默认值： 0，最小值：0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页查询限制数目。默认值： 100，最大值：1000。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 过滤条件，Filters.Values的上限为20。详细的过滤条件如下：
	// <li>host：按照域名进行过滤。</li>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeHostsSettingRequest struct {
	*tchttp.BaseRequest
	
	// 站点ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 分页查询偏移量。默认值： 0，最小值：0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页查询限制数目。默认值： 100，最大值：1000。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 过滤条件，Filters.Values的上限为20。详细的过滤条件如下：
	// <li>host：按照域名进行过滤。</li>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
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
	DetailHosts []*DetailHost `json:"DetailHosts,omitnil,omitempty" name:"DetailHosts"`

	// 域名数量。
	TotalNumber *int64 `json:"TotalNumber,omitnil,omitempty" name:"TotalNumber"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DescribeIPRegionRequestParams struct {
	// 待查询的 IP 列表，支持 IPV4 和 IPV6，最大可查询 100 条。
	IPs []*string `json:"IPs,omitnil,omitempty" name:"IPs"`
}

type DescribeIPRegionRequest struct {
	*tchttp.BaseRequest
	
	// 待查询的 IP 列表，支持 IPV4 和 IPV6，最大可查询 100 条。
	IPs []*string `json:"IPs,omitnil,omitempty" name:"IPs"`
}

func (r *DescribeIPRegionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIPRegionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "IPs")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeIPRegionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIPRegionResponseParams struct {
	// IP 归属信息列表。
	IPRegionInfo []*IPRegionInfo `json:"IPRegionInfo,omitnil,omitempty" name:"IPRegionInfo"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeIPRegionResponse struct {
	*tchttp.BaseResponse
	Response *DescribeIPRegionResponseParams `json:"Response"`
}

func (r *DescribeIPRegionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIPRegionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIdentificationsRequestParams struct {
	// 过滤条件，Filters.Values的上限为20。详细的过滤条件如下：
	// <li>zone-name：按照站点名称进行过滤。</li>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 分页查询偏移量。默认值：0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页查询限制数目。默认值：20，最大值：1000。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeIdentificationsRequest struct {
	*tchttp.BaseRequest
	
	// 过滤条件，Filters.Values的上限为20。详细的过滤条件如下：
	// <li>zone-name：按照站点名称进行过滤。</li>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 分页查询偏移量。默认值：0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页查询限制数目。默认值：20，最大值：1000。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
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
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 站点验证信息列表。
	Identifications []*Identification `json:"Identifications,omitnil,omitempty" name:"Identifications"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DescribeL4ProxyRequestParams struct {
	// 四层代理实例所属站点的 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 分页查询偏移量，不填写时默认为 0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页查询限制数目。默认值：20，最大值：1000。	
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 过滤条件，Filters.Values 的上限为 20。该参数不填写时，返回当前 zone-id 下所有四层代理实例信息。详细的过滤条件如下：
	// <li>proxy-id：按照四层代理实例 ID 进行过滤；</li>
	// <li>ddos-protection-type：按照安全防护类型进行过滤。</li>
	// 
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeL4ProxyRequest struct {
	*tchttp.BaseRequest
	
	// 四层代理实例所属站点的 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 分页查询偏移量，不填写时默认为 0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页查询限制数目。默认值：20，最大值：1000。	
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 过滤条件，Filters.Values 的上限为 20。该参数不填写时，返回当前 zone-id 下所有四层代理实例信息。详细的过滤条件如下：
	// <li>proxy-id：按照四层代理实例 ID 进行过滤；</li>
	// <li>ddos-protection-type：按照安全防护类型进行过滤。</li>
	// 
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribeL4ProxyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeL4ProxyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeL4ProxyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeL4ProxyResponseParams struct {
	// 四层代理实例的数量。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 四层代理实例列表。
	L4Proxies []*L4Proxy `json:"L4Proxies,omitnil,omitempty" name:"L4Proxies"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeL4ProxyResponse struct {
	*tchttp.BaseResponse
	Response *DescribeL4ProxyResponseParams `json:"Response"`
}

func (r *DescribeL4ProxyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeL4ProxyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeL4ProxyRulesRequestParams struct {
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 四层代理实例 ID。
	ProxyId *string `json:"ProxyId,omitnil,omitempty" name:"ProxyId"`

	// 分页查询偏移量，不填写时默认为 0。	
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页查询限制数目。默认值：20，最大值：1000。	
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 过滤条件，Filters.Values的上限为20。不填写时返回当前四层实例下所有的规则信息，详细的过滤条件如下： 
	//  <li>rule-tag：按照规则标签对四层代理实例下的规则进行过滤。</li>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeL4ProxyRulesRequest struct {
	*tchttp.BaseRequest
	
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 四层代理实例 ID。
	ProxyId *string `json:"ProxyId,omitnil,omitempty" name:"ProxyId"`

	// 分页查询偏移量，不填写时默认为 0。	
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页查询限制数目。默认值：20，最大值：1000。	
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 过滤条件，Filters.Values的上限为20。不填写时返回当前四层实例下所有的规则信息，详细的过滤条件如下： 
	//  <li>rule-tag：按照规则标签对四层代理实例下的规则进行过滤。</li>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribeL4ProxyRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeL4ProxyRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "ProxyId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeL4ProxyRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeL4ProxyRulesResponseParams struct {
	// 转发规则总数。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 转发规则列表。	
	L4ProxyRules []*L4ProxyRule `json:"L4ProxyRules,omitnil,omitempty" name:"L4ProxyRules"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeL4ProxyRulesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeL4ProxyRulesResponseParams `json:"Response"`
}

func (r *DescribeL4ProxyRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeL4ProxyRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOriginGroupRequestParams struct {
	// 站点ID，此参数必填。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 分页查询偏移量，不填默认为0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页查询限制数目，不填默认为20，取值：1-1000。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 过滤条件，Filters.Values的上限为20。详细的过滤条件如下：
	// <li>origin-group-id<br>   按照【<strong>源站组ID</strong>】进行过滤。源站组ID形如：origin-2ccgtb24-7dc5-46s2-9r3e-95825d53dwe3a<br>   模糊查询：不支持</li><li>origin-group-name<br>   按照【<strong>源站组名称</strong>】进行过滤<br>   模糊查询：支持。使用模糊查询时，仅支持填写一个源站组名称</li>
	Filters []*AdvancedFilter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeOriginGroupRequest struct {
	*tchttp.BaseRequest
	
	// 站点ID，此参数必填。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 分页查询偏移量，不填默认为0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页查询限制数目，不填默认为20，取值：1-1000。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 过滤条件，Filters.Values的上限为20。详细的过滤条件如下：
	// <li>origin-group-id<br>   按照【<strong>源站组ID</strong>】进行过滤。源站组ID形如：origin-2ccgtb24-7dc5-46s2-9r3e-95825d53dwe3a<br>   模糊查询：不支持</li><li>origin-group-name<br>   按照【<strong>源站组名称</strong>】进行过滤<br>   模糊查询：支持。使用模糊查询时，仅支持填写一个源站组名称</li>
	Filters []*AdvancedFilter `json:"Filters,omitnil,omitempty" name:"Filters"`
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
	delete(f, "ZoneId")
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
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 源站组信息。
	OriginGroups []*OriginGroup `json:"OriginGroups,omitnil,omitempty" name:"OriginGroups"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DescribeOriginProtectionRequestParams struct {
	// 查询的站点ID集合。该参数必填。
	ZoneIds []*string `json:"ZoneIds,omitnil,omitempty" name:"ZoneIds"`

	// 过滤条件，Filters.Values的上限为20。详细的过滤条件如下：
	// <li>need-update<br>   按照【<strong>站点是否需要更新源站防护IP白名单</strong>】进行过滤。<br>   类型：String<br>   必选：否<br>   可选项：<br>   true：需要更新<br>   false：无需更新<br></li>
	// <li>plan-support<br>   按照【<strong>站点套餐是否支持源站防护</strong>】进行过滤。<br>   类型：String<br>   必选：否<br>   可选项：<br>   true：支持<br>   false：不支持<br></li>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 分页查询偏移量，默认为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页查询限制数目。默认值：20，最大值：1000。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeOriginProtectionRequest struct {
	*tchttp.BaseRequest
	
	// 查询的站点ID集合。该参数必填。
	ZoneIds []*string `json:"ZoneIds,omitnil,omitempty" name:"ZoneIds"`

	// 过滤条件，Filters.Values的上限为20。详细的过滤条件如下：
	// <li>need-update<br>   按照【<strong>站点是否需要更新源站防护IP白名单</strong>】进行过滤。<br>   类型：String<br>   必选：否<br>   可选项：<br>   true：需要更新<br>   false：无需更新<br></li>
	// <li>plan-support<br>   按照【<strong>站点套餐是否支持源站防护</strong>】进行过滤。<br>   类型：String<br>   必选：否<br>   可选项：<br>   true：支持<br>   false：不支持<br></li>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 分页查询偏移量，默认为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页查询限制数目。默认值：20，最大值：1000。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeOriginProtectionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOriginProtectionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneIds")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeOriginProtectionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOriginProtectionResponseParams struct {
	// 源站防护信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OriginProtectionInfo []*OriginProtectionInfo `json:"OriginProtectionInfo,omitnil,omitempty" name:"OriginProtectionInfo"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeOriginProtectionResponse struct {
	*tchttp.BaseResponse
	Response *DescribeOriginProtectionResponseParams `json:"Response"`
}

func (r *DescribeOriginProtectionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOriginProtectionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOverviewL7DataRequestParams struct {
	// 开始时间。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 查询的指标，取值有：
	// <li>l7Flow_outFlux: Edegone响应流量；</li>
	// <li>l7Flow_inFlux: Edgeone请求流量；</li>
	// <li>l7Flow_outBandwidth: Edegone响应带宽；</li>
	// <li>l7Flow_inBandwidth: Edegone请求带宽；</li>
	// <li>l7Flow_hit_outFlux: 缓存命中流量；</li>
	// <li>l7Flow_request: 访问请求数；</li>
	// <li>l7Flow_flux: 访问请求上行+下行流量；</li>
	// <li>l7Flow_bandwidth：访问请求上行+下行带宽。</li>
	MetricNames []*string `json:"MetricNames,omitnil,omitempty" name:"MetricNames"`

	// 站点 ID 集合，此参数必填。
	ZoneIds []*string `json:"ZoneIds,omitnil,omitempty" name:"ZoneIds"`

	// 查询的域名集合，此参数已经废弃。
	Domains []*string `json:"Domains,omitnil,omitempty" name:"Domains"`

	// 查询的协议类型，取值有：
	// <li>http: http协议；</li>
	// <li>https: https协议；</li>
	// <li>http2: http2协议；</li>
	// <li>all:  所有协议。</li>不填默认为all，此参数暂未生效。
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// 查询时间粒度，取值有：
	// <li>min：1分钟；</li>
	// <li>5min：5分钟；</li>
	// <li>hour：1小时；</li>
	// <li>day：1天。</li>不填将根据开始时间跟结束时间的间距自动推算粒度，具体为：1小时范围内以min粒度查询，2天范围内以5min粒度查询，7天范围内以hour粒度查询，超过7天以day粒度查询。
	Interval *string `json:"Interval,omitnil,omitempty" name:"Interval"`

	// 过滤条件，详细的过滤条件Key值如下：
	// <li>socket<br>   按照【<strong>HTTP协议类型</strong>】进行过滤。<br>   对应的Value可选项如下：<br>   HTTP：HTTP 协议；<br>   HTTPS：HTTPS协议；<br>   QUIC：QUIC协议。</li>
	// <li>domain<br>   按照【<strong>域名</strong>】进行过滤。</li>
	// <li>tagKey<br>   按照【<strong>标签Key</strong>】进行过滤。</li>
	// <li>tagValue<br>   按照【<strong>标签Value</strong>】进行过滤。</li>
	Filters []*QueryCondition `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 数据归属地区，取值有：
	// <li>overseas：全球（除中国大陆地区）数据；</li>
	// <li>mainland：中国大陆地区数据；</li>
	// <li>global：全球数据。</li>不填默认取值为global。
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`
}

type DescribeOverviewL7DataRequest struct {
	*tchttp.BaseRequest
	
	// 开始时间。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 查询的指标，取值有：
	// <li>l7Flow_outFlux: Edegone响应流量；</li>
	// <li>l7Flow_inFlux: Edgeone请求流量；</li>
	// <li>l7Flow_outBandwidth: Edegone响应带宽；</li>
	// <li>l7Flow_inBandwidth: Edegone请求带宽；</li>
	// <li>l7Flow_hit_outFlux: 缓存命中流量；</li>
	// <li>l7Flow_request: 访问请求数；</li>
	// <li>l7Flow_flux: 访问请求上行+下行流量；</li>
	// <li>l7Flow_bandwidth：访问请求上行+下行带宽。</li>
	MetricNames []*string `json:"MetricNames,omitnil,omitempty" name:"MetricNames"`

	// 站点 ID 集合，此参数必填。
	ZoneIds []*string `json:"ZoneIds,omitnil,omitempty" name:"ZoneIds"`

	// 查询的域名集合，此参数已经废弃。
	Domains []*string `json:"Domains,omitnil,omitempty" name:"Domains"`

	// 查询的协议类型，取值有：
	// <li>http: http协议；</li>
	// <li>https: https协议；</li>
	// <li>http2: http2协议；</li>
	// <li>all:  所有协议。</li>不填默认为all，此参数暂未生效。
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// 查询时间粒度，取值有：
	// <li>min：1分钟；</li>
	// <li>5min：5分钟；</li>
	// <li>hour：1小时；</li>
	// <li>day：1天。</li>不填将根据开始时间跟结束时间的间距自动推算粒度，具体为：1小时范围内以min粒度查询，2天范围内以5min粒度查询，7天范围内以hour粒度查询，超过7天以day粒度查询。
	Interval *string `json:"Interval,omitnil,omitempty" name:"Interval"`

	// 过滤条件，详细的过滤条件Key值如下：
	// <li>socket<br>   按照【<strong>HTTP协议类型</strong>】进行过滤。<br>   对应的Value可选项如下：<br>   HTTP：HTTP 协议；<br>   HTTPS：HTTPS协议；<br>   QUIC：QUIC协议。</li>
	// <li>domain<br>   按照【<strong>域名</strong>】进行过滤。</li>
	// <li>tagKey<br>   按照【<strong>标签Key</strong>】进行过滤。</li>
	// <li>tagValue<br>   按照【<strong>标签Value</strong>】进行过滤。</li>
	Filters []*QueryCondition `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 数据归属地区，取值有：
	// <li>overseas：全球（除中国大陆地区）数据；</li>
	// <li>mainland：中国大陆地区数据；</li>
	// <li>global：全球数据。</li>不填默认取值为global。
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`
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
	delete(f, "Filters")
	delete(f, "Area")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeOverviewL7DataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOverviewL7DataResponseParams struct {
	// 查询结果的总条数。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 七层监控类时序流量数据列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*TimingDataRecord `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	// 站点ID。该参数必填。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 查询起始时间，时间与 job-id 必填一个。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 查询结束时间，时间与 job-id 必填一个。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 分页查询偏移量，默认为 0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页查询限制数目，默认值：20，上限：1000。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 过滤条件，Filters.Values 的上限为 20。详细的过滤条件如下：<li>job-id：按照任务 ID 进行过滤。job-id 形如：1379afjk91u32h，暂不支持多值，不支持模糊查询；</li><li>target：按照目标资源信息进行过滤。target 形如：http://www.qq.com/1.txt，暂不支持多值，不支持模糊查询；</li><li>domains：按照域名行过滤。domains 形如：www.qq.com，不支持模糊查询；</li><li>statuses：按照任务状态进行过滤，不支持模糊查询。可选项：<br>   processing：处理中<br>   success：成功<br>   failed：失败<br>   timeout：超时<br>   invalid：无效。即源站响应非 2xx 状态码，请检查源站服务。</li>
	Filters []*AdvancedFilter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribePrefetchTasksRequest struct {
	*tchttp.BaseRequest
	
	// 站点ID。该参数必填。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 查询起始时间，时间与 job-id 必填一个。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 查询结束时间，时间与 job-id 必填一个。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 分页查询偏移量，默认为 0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页查询限制数目，默认值：20，上限：1000。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 过滤条件，Filters.Values 的上限为 20。详细的过滤条件如下：<li>job-id：按照任务 ID 进行过滤。job-id 形如：1379afjk91u32h，暂不支持多值，不支持模糊查询；</li><li>target：按照目标资源信息进行过滤。target 形如：http://www.qq.com/1.txt，暂不支持多值，不支持模糊查询；</li><li>domains：按照域名行过滤。domains 形如：www.qq.com，不支持模糊查询；</li><li>statuses：按照任务状态进行过滤，不支持模糊查询。可选项：<br>   processing：处理中<br>   success：成功<br>   failed：失败<br>   timeout：超时<br>   invalid：无效。即源站响应非 2xx 状态码，请检查源站服务。</li>
	Filters []*AdvancedFilter `json:"Filters,omitnil,omitempty" name:"Filters"`
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
	delete(f, "ZoneId")
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
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 任务结果列表。
	Tasks []*Task `json:"Tasks,omitnil,omitempty" name:"Tasks"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	// 站点 ID。该参数必填。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 查询起始时间，时间与 job-id 必填一个。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 查询结束时间，时间与 job-id 必填一个。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 分页查询偏移量，默认为 0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页查询限制数目，默认值：20，最大值：1000。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 过滤条件，Filters.Values的上限为20。详细的过滤条件如下：
	// <li>job-id：按照任务 ID 进行过滤。job-id 形如：1379afjk91u32h，暂不支持多值，不支持模糊查询；</li>
	// <li>target：按照目标资源信息进行过滤，target 形如：http://www.qq.com/1.txt 或者 tag1，暂不支持多值，支持模糊查询；</li>
	// <li>domains：按照域名进行过滤，形如：www.qq.com，不支持模糊查询；</li>
	// <li>statuses：按照任务状态进行过滤，不支持模糊查询。可选项：<br>   processing：处理中<br>   success：成功<br>   failed：失败<br>   timeout：超时</li>
	// <li>type：按照清除缓存类型进行过滤，暂不支持多值，不支持模糊查询。可选项：<br>   purge_url：URL<br>   purge_prefix：前缀<br>   purge_all：全部缓存内容<br>   purge_host：Hostname<br>   purge_cache_tag：CacheTag</li>
	Filters []*AdvancedFilter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribePurgeTasksRequest struct {
	*tchttp.BaseRequest
	
	// 站点 ID。该参数必填。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 查询起始时间，时间与 job-id 必填一个。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 查询结束时间，时间与 job-id 必填一个。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 分页查询偏移量，默认为 0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页查询限制数目，默认值：20，最大值：1000。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 过滤条件，Filters.Values的上限为20。详细的过滤条件如下：
	// <li>job-id：按照任务 ID 进行过滤。job-id 形如：1379afjk91u32h，暂不支持多值，不支持模糊查询；</li>
	// <li>target：按照目标资源信息进行过滤，target 形如：http://www.qq.com/1.txt 或者 tag1，暂不支持多值，支持模糊查询；</li>
	// <li>domains：按照域名进行过滤，形如：www.qq.com，不支持模糊查询；</li>
	// <li>statuses：按照任务状态进行过滤，不支持模糊查询。可选项：<br>   processing：处理中<br>   success：成功<br>   failed：失败<br>   timeout：超时</li>
	// <li>type：按照清除缓存类型进行过滤，暂不支持多值，不支持模糊查询。可选项：<br>   purge_url：URL<br>   purge_prefix：前缀<br>   purge_all：全部缓存内容<br>   purge_host：Hostname<br>   purge_cache_tag：CacheTag</li>
	Filters []*AdvancedFilter `json:"Filters,omitnil,omitempty" name:"Filters"`
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
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 任务结果列表。
	Tasks []*Task `json:"Tasks,omitnil,omitempty" name:"Tasks"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DescribeRealtimeLogDeliveryTasksRequestParams struct {
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 分页查询偏移量。默认值：0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页查询限制数目。默认值：20，最大值：1000。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 过滤条件，Filters.Values 的上限为 20。该参数不填写时，返回当前 zone-id 下所有实时日志投递任务信息。详细的过滤条件如下：
	// <li>task-id：按照实时日志投递任务 ID进行过滤。不支持模糊查询。</li>
	// <li>task-name：按照实时日志投递任务名称进行过滤。支持模糊查询，使用模糊查询时，仅支持填写一个实时日志投递任务名称。</li>
	// <li>entity-list：按照实时日志投递任务对应的实体进行过滤。不支持模糊查询。示例值：domain.example.com 或者 sid-2s69eb5wcms7。</li>
	// <li>task-type：按照实时日志投递任务类型进行过滤。不支持模糊查询。可选项如下：<br>   cls: 推送到腾讯云 CLS；<br>   custom_endpoint：推送到自定义 HTTP(S) 地址；<br>   s3：推送到 AWS S3 兼容存储桶地址。</li>
	Filters []*AdvancedFilter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeRealtimeLogDeliveryTasksRequest struct {
	*tchttp.BaseRequest
	
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 分页查询偏移量。默认值：0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页查询限制数目。默认值：20，最大值：1000。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 过滤条件，Filters.Values 的上限为 20。该参数不填写时，返回当前 zone-id 下所有实时日志投递任务信息。详细的过滤条件如下：
	// <li>task-id：按照实时日志投递任务 ID进行过滤。不支持模糊查询。</li>
	// <li>task-name：按照实时日志投递任务名称进行过滤。支持模糊查询，使用模糊查询时，仅支持填写一个实时日志投递任务名称。</li>
	// <li>entity-list：按照实时日志投递任务对应的实体进行过滤。不支持模糊查询。示例值：domain.example.com 或者 sid-2s69eb5wcms7。</li>
	// <li>task-type：按照实时日志投递任务类型进行过滤。不支持模糊查询。可选项如下：<br>   cls: 推送到腾讯云 CLS；<br>   custom_endpoint：推送到自定义 HTTP(S) 地址；<br>   s3：推送到 AWS S3 兼容存储桶地址。</li>
	Filters []*AdvancedFilter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribeRealtimeLogDeliveryTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRealtimeLogDeliveryTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRealtimeLogDeliveryTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRealtimeLogDeliveryTasksResponseParams struct {
	// 符合查询条件的实时日志投递任务个数。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 符合查询条件的所有实时日志投递任务列表。
	RealtimeLogDeliveryTasks []*RealtimeLogDeliveryTask `json:"RealtimeLogDeliveryTasks,omitnil,omitempty" name:"RealtimeLogDeliveryTasks"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRealtimeLogDeliveryTasksResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRealtimeLogDeliveryTasksResponseParams `json:"Response"`
}

func (r *DescribeRealtimeLogDeliveryTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRealtimeLogDeliveryTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRulesRequestParams struct {
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 过滤条件，Filters.Values的上限为20。详细的过滤条件如下：
	// <li>rule-id<br>   按照【<strong>规则ID</strong>】进行过滤。<br>   类型：string<br>   必选：否</li>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeRulesRequest struct {
	*tchttp.BaseRequest
	
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 过滤条件，Filters.Values的上限为20。详细的过滤条件如下：
	// <li>rule-id<br>   按照【<strong>规则ID</strong>】进行过滤。<br>   类型：string<br>   必选：否</li>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
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
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 规则列表，按规则执行顺序从先往后排序。
	RuleItems []*RuleItem `json:"RuleItems,omitnil,omitempty" name:"RuleItems"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Actions []*RulesSettingAction `json:"Actions,omitnil,omitempty" name:"Actions"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DescribeSecurityIPGroupInfoRequestParams struct {
	// 站点的 ID ，用于指定查询的站点范围。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 单次返回的最大条目数。默认值为 20 ，最大查询条目为 1000 。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页查询的起始条目偏移量。默认值为 0 。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type DescribeSecurityIPGroupInfoRequest struct {
	*tchttp.BaseRequest
	
	// 站点的 ID ，用于指定查询的站点范围。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 单次返回的最大条目数。默认值为 20 ，最大查询条目为 1000 。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页查询的起始条目偏移量。默认值为 0 。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

func (r *DescribeSecurityIPGroupInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecurityIPGroupInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSecurityIPGroupInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSecurityIPGroupInfoResponseParams struct {
	// 返回的满足条件的 IP 组数量。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// IP 组的详细配置信息。包含每个 IP 组的 ID 、名称和 IP /网段列表信息。
	IPGroups []*IPGroup `json:"IPGroups,omitnil,omitempty" name:"IPGroups"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSecurityIPGroupInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSecurityIPGroupInfoResponseParams `json:"Response"`
}

func (r *DescribeSecurityIPGroupInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecurityIPGroupInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSecurityIPGroupRequestParams struct {
	// 站点 ID ，用于指定查询的站点范围。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 指定安全 IP 组 ID。
	// <li>提供该参数时，仅查询指定 ID 的安全 IP 组配置；</li>
	// <li>不传递参数时，返回站点下所有安全 IP 组信息。</li>
	GroupIds []*int64 `json:"GroupIds,omitnil,omitempty" name:"GroupIds"`
}

type DescribeSecurityIPGroupRequest struct {
	*tchttp.BaseRequest
	
	// 站点 ID ，用于指定查询的站点范围。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 指定安全 IP 组 ID。
	// <li>提供该参数时，仅查询指定 ID 的安全 IP 组配置；</li>
	// <li>不传递参数时，返回站点下所有安全 IP 组信息。</li>
	GroupIds []*int64 `json:"GroupIds,omitnil,omitempty" name:"GroupIds"`
}

func (r *DescribeSecurityIPGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecurityIPGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "GroupIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSecurityIPGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSecurityIPGroupResponseParams struct {
	// 安全 IP 组的详细配置信息。包含每个安全 IP 组的 ID 、名称和 IP / 网段列表信息。
	IPGroups []*IPGroup `json:"IPGroups,omitnil,omitempty" name:"IPGroups"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSecurityIPGroupResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSecurityIPGroupResponseParams `json:"Response"`
}

func (r *DescribeSecurityIPGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecurityIPGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSecurityTemplateBindingsRequestParams struct {
	// 要查询的站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 要查询的策略模板 ID。
	TemplateId []*string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`
}

type DescribeSecurityTemplateBindingsRequest struct {
	*tchttp.BaseRequest
	
	// 要查询的站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 要查询的策略模板 ID。
	TemplateId []*string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`
}

func (r *DescribeSecurityTemplateBindingsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecurityTemplateBindingsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "TemplateId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSecurityTemplateBindingsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSecurityTemplateBindingsResponseParams struct {
	// 指定策略模板的绑定关系列表。
	// 
	// 当某个站点中的域名包含在指定策略模板的绑定关系中时，绑定关系列表 `TemplateScope` 中会包含该站点的 `ZoneId`，和该站点下的和该策略模板有关的域名绑定关系。
	// 
	// 注意：当没有任何域名正在绑定或已经绑定到指定策略模板时，绑定关系为空。即：返回结构体中，`TemplateScope` 数组长度为 0。
	// 
	// 绑定关系中，同一域名可能在 `EntityStatus` 列表中重复出现，并标记为不同 `Status` 。例如，正在被绑定到其他策略模板的域名，会同时标记为 `online` 和 `pending` 。
	SecurityTemplate []*SecurityTemplateBinding `json:"SecurityTemplate,omitnil,omitempty" name:"SecurityTemplate"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSecurityTemplateBindingsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSecurityTemplateBindingsResponseParams `json:"Response"`
}

func (r *DescribeSecurityTemplateBindingsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecurityTemplateBindingsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTimingL4DataRequestParams struct {
	// 开始时间。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 查询指标，取值有：
	// <li>l4Flow_connections: 访问连接数；</li>
	// <li>l4Flow_flux: 访问总流量；</li>
	// <li>l4Flow_inFlux: 访问入流量；</li>
	// <li>l4Flow_outFlux: 访问出流量。</li>
	MetricNames []*string `json:"MetricNames,omitnil,omitempty" name:"MetricNames"`

	// 站点 ID 集合，此参数必填。
	ZoneIds []*string `json:"ZoneIds,omitnil,omitempty" name:"ZoneIds"`

	// 四层实例列表, 不填表示选择全部实例。
	ProxyIds []*string `json:"ProxyIds,omitnil,omitempty" name:"ProxyIds"`

	// 查询时间粒度，取值有：
	// <li>min: 1分钟 ；</li>
	// <li>5min: 5分钟 ；</li>
	// <li>hour: 1小时 ；</li>
	// <li>day: 1天 。</li>不填将根据开始时间跟结束时间的间距自动推算粒度，具体为：1小时范围内以min粒度查询，2天范围内以5min粒度查询，7天范围内以hour粒度查询，超过7天以day粒度查询。
	Interval *string `json:"Interval,omitnil,omitempty" name:"Interval"`

	// 过滤条件，详细的过滤条件Key值如下：
	// <li>ruleId：按照转发规则 ID 进行过滤。</li>
	// <li>proxyId：按照四层代理实例 ID 进行过滤。</li>
	Filters []*QueryCondition `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 数据归属地区，取值有：
	// <li>overseas：全球（除中国大陆地区）数据；</li>
	// <li>mainland：中国大陆地区数据；</li>
	// <li>global：全球数据。</li>不填默认取值为global。
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`
}

type DescribeTimingL4DataRequest struct {
	*tchttp.BaseRequest
	
	// 开始时间。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 查询指标，取值有：
	// <li>l4Flow_connections: 访问连接数；</li>
	// <li>l4Flow_flux: 访问总流量；</li>
	// <li>l4Flow_inFlux: 访问入流量；</li>
	// <li>l4Flow_outFlux: 访问出流量。</li>
	MetricNames []*string `json:"MetricNames,omitnil,omitempty" name:"MetricNames"`

	// 站点 ID 集合，此参数必填。
	ZoneIds []*string `json:"ZoneIds,omitnil,omitempty" name:"ZoneIds"`

	// 四层实例列表, 不填表示选择全部实例。
	ProxyIds []*string `json:"ProxyIds,omitnil,omitempty" name:"ProxyIds"`

	// 查询时间粒度，取值有：
	// <li>min: 1分钟 ；</li>
	// <li>5min: 5分钟 ；</li>
	// <li>hour: 1小时 ；</li>
	// <li>day: 1天 。</li>不填将根据开始时间跟结束时间的间距自动推算粒度，具体为：1小时范围内以min粒度查询，2天范围内以5min粒度查询，7天范围内以hour粒度查询，超过7天以day粒度查询。
	Interval *string `json:"Interval,omitnil,omitempty" name:"Interval"`

	// 过滤条件，详细的过滤条件Key值如下：
	// <li>ruleId：按照转发规则 ID 进行过滤。</li>
	// <li>proxyId：按照四层代理实例 ID 进行过滤。</li>
	Filters []*QueryCondition `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 数据归属地区，取值有：
	// <li>overseas：全球（除中国大陆地区）数据；</li>
	// <li>mainland：中国大陆地区数据；</li>
	// <li>global：全球数据。</li>不填默认取值为global。
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`
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
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 四层时序流量数据列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*TimingDataRecord `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 指标列表，取值有:
	// <li>l7Flow_outFlux: Edgeone 响应流量；</li>
	// <li>l7Flow_inFlux: Edgeone 请求流量；</li>
	// <li>l7Flow_outBandwidth: Edgeone 响应带宽；</li>
	// <li>l7Flow_inBandwidth：Edgeone 请求带宽；</li>
	// <li>l7Flow_request: 访问请求数；</li>
	// <li>l7Flow_flux: 访问请求上行+下行流量；</li>
	// <li>l7Flow_bandwidth：访问请求上行+下行带宽。</li>
	MetricNames []*string `json:"MetricNames,omitnil,omitempty" name:"MetricNames"`

	// 站点 ID 集合, 此参数必填。
	ZoneIds []*string `json:"ZoneIds,omitnil,omitempty" name:"ZoneIds"`

	// 查询时间粒度，取值有：
	// <li>min: 1分钟；</li>
	// <li>5min: 5分钟；</li>
	// <li>hour: 1小时；</li>
	// <li>day: 1天。</li>不填将根据开始时间跟结束时间的间距自动推算粒度，具体为：1小时范围内以min粒度查询，2天范围内以5min粒度查询，7天范围内以hour粒度查询，超过7天以day粒度查询。
	Interval *string `json:"Interval,omitnil,omitempty" name:"Interval"`

	// 过滤条件，详细的过滤条件Key值如下：
	// <li>country<br>   按照【<strong>国家/地区</strong>】进行过滤，国家/地区遵循 <a href="https://baike.baidu.com/item/ISO%203166-1/5269555">ISO 3166</a> 规范。</li>
	// <li>province<br>   按照【<strong>省份</strong>】进行过滤，此参数只支持服务区域为中国大陆。</li>
	// <li>isp<br>   按照【<strong>运营商</strong>】进行过滤，此参数只支持服务区域为中国大陆。<br>   对应的Value可选项如下：<br>   2：中国电信；<br>   26：中国联通；<br>   1046：中国移动；<br>   3947：中国铁通；<br>   38：教育网；<br>   43：长城宽带；<br>   0：其他运营商。</li>
	// <li>domain<br>   按照【<strong>子域名</strong>】进行过滤，子域名形如： test.example.com。</li>
	// <li>url<br>   按照【<strong>URL Path</strong>】进行过滤，URL Path形如：/content或/content/test.jpg。<br>   若只填写url参数，则最多可查询近30天的数据；<br>   若同时填写url+Zonelds参数，则支持的查询数据范围为套餐支持的<a href="https://cloud.tencent.com/document/product/1552/77380#edgeone-.E5.A5.97.E9.A4.90">数据分析最大查询范围</a>与30天两者中的较小值。</li>
	// <li>referer<br>   按照【<strong>Referer头信息</strong>】进行过滤, Referer形如：example.com。<br>   若只填写referer参数，则最多可查询近30天的数据；<br>   若同时填写referer+Zonelds参数，则支持的查询数据范围为套餐支持的<a href="https://cloud.tencent.com/document/product/1552/77380#edgeone-.E5.A5.97.E9.A4.90">数据分析最大查询范围</a>与30天两者中的较小值。</li>
	// <li>resourceType<br>   按照【<strong>资源类型</strong>】进行过滤，资源类型一般是文件后缀，形如: .jpg, .css。<br>   若只填写resourceType参数，则最多可查询近30天的数据；<br>   若同时填写resourceType+Zonelds参数，则支持的查询数据范围为套餐支持的<a href="https://cloud.tencent.com/document/product/1552/77380#edgeone-.E5.A5.97.E9.A4.90">数据分析最大查询范围</a>与30天两者中的较小值。</li>
	// <li>protocol<br>   按照【<strong>HTTP协议版本</strong>】进行过滤。<br>   对应的Value可选项如下：<br>   HTTP/1.0：HTTP 1.0；<br>   HTTP/1.1：HTTP 1.1；<br>   HTTP/2.0：HTTP 2.0；<br>   HTTP/3.0：HTTP 3.0；<br>   WebSocket：WebSocket。</li>
	// <li>socket<br>   按照【<strong>HTTP协议类型</strong>】进行过滤。<br>   对应的Value可选项如下：<br>   HTTP：HTTP 协议；<br>   HTTPS：HTTPS协议；<br>   QUIC：QUIC协议。</li>
	// <li>statusCode<br>   按照【<strong>状态码</strong>】进行过滤。<br>   若只填写statusCode参数，则最多可查询近30天的数据；<br>   若同时填写statusCode+Zonelds参数，则支持的查询数据范围为套餐支持的<a href="https://cloud.tencent.com/document/product/1552/77380#edgeone-.E5.A5.97.E9.A4.90">数据分析最大查询范围</a>与30天两者中的较小值。<br>   对应的Value可选项如下：<br>   1XX：1xx类型的状态码；<br>   100：100状态码；<br>   101：101状态码；<br>   102：102状态码；<br>   2XX：2xx类型的状态码；<br>   200：200状态码；<br>   201：201状态码；<br>   202：202状态码；<br>   203：203状态码；<br>   204：204状态码；<br>   205：205状态码；<br>   206：206状态码；<br>   207：207状态码；<br>   3XX：3xx类型的状态码；<br>   300：300状态码；<br>   301：301状态码；<br>   302：302状态码；<br>   303：303状态码；<br>   304：304状态码；<br>   305：305状态码；<br>   307：307状态码；<br>   4XX：4xx类型的状态码；<br>   400：400状态码；<br>   401：401状态码；<br>   402：402状态码；<br>   403：403状态码；<br>   404：404状态码；<br>   405：405状态码；<br>   406：406状态码；<br>   407：407状态码；<br>   408：408状态码；<br>   409：409状态码；<br>   410：410状态码；<br>   411：411状态码；<br>   412：412状态码；<br>   412：413状态码；<br>   414：414状态码；<br>   415：415状态码；<br>   416：416状态码；<br>   417：417状态码；<br>   422：422状态码；<br>   423：423状态码；<br>   424：424状态码；<br>   426：426状态码；<br>   451：451状态码；<br>   5XX：5xx类型的状态码；<br>   500：500状态码；<br>   501：501状态码；<br>   502：502状态码；<br>   503：503状态码；<br>   504：504状态码；<br>   505：505状态码；<br>   506：506状态码；<br>   507：507状态码；<br>   510：510状态码；<br>   514：514状态码；<br>   544：544状态码。</li>
	// <li>browserType<br>   按照【<strong>浏览器类型</strong>】进行过滤。<br>   若只填写browserType参数，则最多可查询近30天的数据；<br>   若同时填写browserType+Zonelds参数，则支持的查询数据范围为套餐支持的<a href="https://cloud.tencent.com/document/product/1552/77380#edgeone-.E5.A5.97.E9.A4.90">数据分析最大查询范围</a>与30天两者中的较小值。<br>   对应Value的可选项如下：<br>   Firefox：Firefox浏览器；<br>   Chrome：Chrome浏览器；<br>   Safari：Safari浏览器；<br>   Other：其他浏览器类型；<br>   Empty：浏览器类型为空；<br>   Bot：搜索引擎爬虫；<br>   MicrosoftEdge：MicrosoftEdge浏览器；<br>   IE：IE浏览器；<br>   Opera：Opera浏览器；<br>   QQBrowser：QQ浏览器；<br>   LBBrowser：LB浏览器；<br>   MaxthonBrowser：Maxthon浏览器；<br>   SouGouBrowser：搜狗浏览器；<br>   BIDUBrowser：百度浏览器；<br>   TaoBrowser：淘浏览器；<br>   UBrowser：UC浏览器。</li>
	// <li>deviceType<br>   按照【<strong>设备类型</strong>】进行过滤。<br>   若只填写deviceType参数，则最多可查询近30天的数据；<br>   若同时填写deviceType+Zonelds参数，则支持的查询数据范围为套餐支持的<a href="https://cloud.tencent.com/document/product/1552/77380#edgeone-.E5.A5.97.E9.A4.90">数据分析最大查询范围</a>与30天两者中的较小值。<br>   对应Value的可选项如下：<br>   TV：TV设备；<br>   Tablet：Tablet设备；<br>   Mobile：Mobile设备；<br>   Desktop：Desktop设备；<br>   Other：其他设备类型；<br>   Empty：设备类型为空。</li>
	// <li>operatingSystemType<br>   按照【<strong>操作系统类型</strong>】进行过滤。<br>   若只填写operatingSystemType参数，则最多可查询近30天的数据；<br>   若同时填写operatingSystemType+Zonelds参数，则支持的查询数据范围为套餐支持的<a href="https://cloud.tencent.com/document/product/1552/77380#edgeone-.E5.A5.97.E9.A4.90">数据分析最大查询范围</a>与30天两者中的较小值。<br>   对应Value的可选项如下：<br>   Linux：Linux操作系统；<br>   MacOS：MacOs操作系统；<br>   Android：Android操作系统；<br>   IOS：IOS操作系统；<br>   Windows：Windows操作系统；<br>   NetBSD：NetBSD；<br>   ChromiumOS：ChromiumOS；<br>   Bot：搜索引擎爬虫；<br>   Other：其他类型的操作系统；<br>   Empty：操作系统为空。</li>
	// <li>tlsVersion<br>   按照【<strong>TLS版本</strong>】进行过滤。<br>   若只填写tlsVersion参数，则最多可查询近30天的数据；<br>   若同时填写tlsVersion+Zonelds参数，则支持的查询数据范围为套餐支持的<a href="https://cloud.tencent.com/document/product/1552/77380#edgeone-.E5.A5.97.E9.A4.90">数据分析最大查询范围</a>与30天两者中的较小值。<br>   对应Value的可选项如下：<br>   TLS1.0：TLS 1.0；<br>   TLS1.1：TLS 1.1；<br>   TLS1.2：TLS 1.2；<br>   TLS1.3：TLS 1.3。</li>
	// <li>ipVersion<br>   按照【<strong>IP版本</strong>】进行过滤。<br>   对应Value的可选项如下：<br>   4：Ipv4；<br>   6：Ipv6。</li>
	// <li>tagKey<br>   按照【<strong>标签Key</strong>】进行过滤。</li>
	// <li>tagValue<br>   按照【<strong>标签Value</strong>】进行过滤。</li>
	Filters []*QueryCondition `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 数据归属地区，取值有：
	// <li>overseas：全球（除中国大陆地区）数据；</li>
	// <li>mainland：中国大陆地区数据；</li>
	// <li>global：全球数据。</li>不填默认取值为global。
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`
}

type DescribeTimingL7AnalysisDataRequest struct {
	*tchttp.BaseRequest
	
	// 开始时间。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 指标列表，取值有:
	// <li>l7Flow_outFlux: Edgeone 响应流量；</li>
	// <li>l7Flow_inFlux: Edgeone 请求流量；</li>
	// <li>l7Flow_outBandwidth: Edgeone 响应带宽；</li>
	// <li>l7Flow_inBandwidth：Edgeone 请求带宽；</li>
	// <li>l7Flow_request: 访问请求数；</li>
	// <li>l7Flow_flux: 访问请求上行+下行流量；</li>
	// <li>l7Flow_bandwidth：访问请求上行+下行带宽。</li>
	MetricNames []*string `json:"MetricNames,omitnil,omitempty" name:"MetricNames"`

	// 站点 ID 集合, 此参数必填。
	ZoneIds []*string `json:"ZoneIds,omitnil,omitempty" name:"ZoneIds"`

	// 查询时间粒度，取值有：
	// <li>min: 1分钟；</li>
	// <li>5min: 5分钟；</li>
	// <li>hour: 1小时；</li>
	// <li>day: 1天。</li>不填将根据开始时间跟结束时间的间距自动推算粒度，具体为：1小时范围内以min粒度查询，2天范围内以5min粒度查询，7天范围内以hour粒度查询，超过7天以day粒度查询。
	Interval *string `json:"Interval,omitnil,omitempty" name:"Interval"`

	// 过滤条件，详细的过滤条件Key值如下：
	// <li>country<br>   按照【<strong>国家/地区</strong>】进行过滤，国家/地区遵循 <a href="https://baike.baidu.com/item/ISO%203166-1/5269555">ISO 3166</a> 规范。</li>
	// <li>province<br>   按照【<strong>省份</strong>】进行过滤，此参数只支持服务区域为中国大陆。</li>
	// <li>isp<br>   按照【<strong>运营商</strong>】进行过滤，此参数只支持服务区域为中国大陆。<br>   对应的Value可选项如下：<br>   2：中国电信；<br>   26：中国联通；<br>   1046：中国移动；<br>   3947：中国铁通；<br>   38：教育网；<br>   43：长城宽带；<br>   0：其他运营商。</li>
	// <li>domain<br>   按照【<strong>子域名</strong>】进行过滤，子域名形如： test.example.com。</li>
	// <li>url<br>   按照【<strong>URL Path</strong>】进行过滤，URL Path形如：/content或/content/test.jpg。<br>   若只填写url参数，则最多可查询近30天的数据；<br>   若同时填写url+Zonelds参数，则支持的查询数据范围为套餐支持的<a href="https://cloud.tencent.com/document/product/1552/77380#edgeone-.E5.A5.97.E9.A4.90">数据分析最大查询范围</a>与30天两者中的较小值。</li>
	// <li>referer<br>   按照【<strong>Referer头信息</strong>】进行过滤, Referer形如：example.com。<br>   若只填写referer参数，则最多可查询近30天的数据；<br>   若同时填写referer+Zonelds参数，则支持的查询数据范围为套餐支持的<a href="https://cloud.tencent.com/document/product/1552/77380#edgeone-.E5.A5.97.E9.A4.90">数据分析最大查询范围</a>与30天两者中的较小值。</li>
	// <li>resourceType<br>   按照【<strong>资源类型</strong>】进行过滤，资源类型一般是文件后缀，形如: .jpg, .css。<br>   若只填写resourceType参数，则最多可查询近30天的数据；<br>   若同时填写resourceType+Zonelds参数，则支持的查询数据范围为套餐支持的<a href="https://cloud.tencent.com/document/product/1552/77380#edgeone-.E5.A5.97.E9.A4.90">数据分析最大查询范围</a>与30天两者中的较小值。</li>
	// <li>protocol<br>   按照【<strong>HTTP协议版本</strong>】进行过滤。<br>   对应的Value可选项如下：<br>   HTTP/1.0：HTTP 1.0；<br>   HTTP/1.1：HTTP 1.1；<br>   HTTP/2.0：HTTP 2.0；<br>   HTTP/3.0：HTTP 3.0；<br>   WebSocket：WebSocket。</li>
	// <li>socket<br>   按照【<strong>HTTP协议类型</strong>】进行过滤。<br>   对应的Value可选项如下：<br>   HTTP：HTTP 协议；<br>   HTTPS：HTTPS协议；<br>   QUIC：QUIC协议。</li>
	// <li>statusCode<br>   按照【<strong>状态码</strong>】进行过滤。<br>   若只填写statusCode参数，则最多可查询近30天的数据；<br>   若同时填写statusCode+Zonelds参数，则支持的查询数据范围为套餐支持的<a href="https://cloud.tencent.com/document/product/1552/77380#edgeone-.E5.A5.97.E9.A4.90">数据分析最大查询范围</a>与30天两者中的较小值。<br>   对应的Value可选项如下：<br>   1XX：1xx类型的状态码；<br>   100：100状态码；<br>   101：101状态码；<br>   102：102状态码；<br>   2XX：2xx类型的状态码；<br>   200：200状态码；<br>   201：201状态码；<br>   202：202状态码；<br>   203：203状态码；<br>   204：204状态码；<br>   205：205状态码；<br>   206：206状态码；<br>   207：207状态码；<br>   3XX：3xx类型的状态码；<br>   300：300状态码；<br>   301：301状态码；<br>   302：302状态码；<br>   303：303状态码；<br>   304：304状态码；<br>   305：305状态码；<br>   307：307状态码；<br>   4XX：4xx类型的状态码；<br>   400：400状态码；<br>   401：401状态码；<br>   402：402状态码；<br>   403：403状态码；<br>   404：404状态码；<br>   405：405状态码；<br>   406：406状态码；<br>   407：407状态码；<br>   408：408状态码；<br>   409：409状态码；<br>   410：410状态码；<br>   411：411状态码；<br>   412：412状态码；<br>   412：413状态码；<br>   414：414状态码；<br>   415：415状态码；<br>   416：416状态码；<br>   417：417状态码；<br>   422：422状态码；<br>   423：423状态码；<br>   424：424状态码；<br>   426：426状态码；<br>   451：451状态码；<br>   5XX：5xx类型的状态码；<br>   500：500状态码；<br>   501：501状态码；<br>   502：502状态码；<br>   503：503状态码；<br>   504：504状态码；<br>   505：505状态码；<br>   506：506状态码；<br>   507：507状态码；<br>   510：510状态码；<br>   514：514状态码；<br>   544：544状态码。</li>
	// <li>browserType<br>   按照【<strong>浏览器类型</strong>】进行过滤。<br>   若只填写browserType参数，则最多可查询近30天的数据；<br>   若同时填写browserType+Zonelds参数，则支持的查询数据范围为套餐支持的<a href="https://cloud.tencent.com/document/product/1552/77380#edgeone-.E5.A5.97.E9.A4.90">数据分析最大查询范围</a>与30天两者中的较小值。<br>   对应Value的可选项如下：<br>   Firefox：Firefox浏览器；<br>   Chrome：Chrome浏览器；<br>   Safari：Safari浏览器；<br>   Other：其他浏览器类型；<br>   Empty：浏览器类型为空；<br>   Bot：搜索引擎爬虫；<br>   MicrosoftEdge：MicrosoftEdge浏览器；<br>   IE：IE浏览器；<br>   Opera：Opera浏览器；<br>   QQBrowser：QQ浏览器；<br>   LBBrowser：LB浏览器；<br>   MaxthonBrowser：Maxthon浏览器；<br>   SouGouBrowser：搜狗浏览器；<br>   BIDUBrowser：百度浏览器；<br>   TaoBrowser：淘浏览器；<br>   UBrowser：UC浏览器。</li>
	// <li>deviceType<br>   按照【<strong>设备类型</strong>】进行过滤。<br>   若只填写deviceType参数，则最多可查询近30天的数据；<br>   若同时填写deviceType+Zonelds参数，则支持的查询数据范围为套餐支持的<a href="https://cloud.tencent.com/document/product/1552/77380#edgeone-.E5.A5.97.E9.A4.90">数据分析最大查询范围</a>与30天两者中的较小值。<br>   对应Value的可选项如下：<br>   TV：TV设备；<br>   Tablet：Tablet设备；<br>   Mobile：Mobile设备；<br>   Desktop：Desktop设备；<br>   Other：其他设备类型；<br>   Empty：设备类型为空。</li>
	// <li>operatingSystemType<br>   按照【<strong>操作系统类型</strong>】进行过滤。<br>   若只填写operatingSystemType参数，则最多可查询近30天的数据；<br>   若同时填写operatingSystemType+Zonelds参数，则支持的查询数据范围为套餐支持的<a href="https://cloud.tencent.com/document/product/1552/77380#edgeone-.E5.A5.97.E9.A4.90">数据分析最大查询范围</a>与30天两者中的较小值。<br>   对应Value的可选项如下：<br>   Linux：Linux操作系统；<br>   MacOS：MacOs操作系统；<br>   Android：Android操作系统；<br>   IOS：IOS操作系统；<br>   Windows：Windows操作系统；<br>   NetBSD：NetBSD；<br>   ChromiumOS：ChromiumOS；<br>   Bot：搜索引擎爬虫；<br>   Other：其他类型的操作系统；<br>   Empty：操作系统为空。</li>
	// <li>tlsVersion<br>   按照【<strong>TLS版本</strong>】进行过滤。<br>   若只填写tlsVersion参数，则最多可查询近30天的数据；<br>   若同时填写tlsVersion+Zonelds参数，则支持的查询数据范围为套餐支持的<a href="https://cloud.tencent.com/document/product/1552/77380#edgeone-.E5.A5.97.E9.A4.90">数据分析最大查询范围</a>与30天两者中的较小值。<br>   对应Value的可选项如下：<br>   TLS1.0：TLS 1.0；<br>   TLS1.1：TLS 1.1；<br>   TLS1.2：TLS 1.2；<br>   TLS1.3：TLS 1.3。</li>
	// <li>ipVersion<br>   按照【<strong>IP版本</strong>】进行过滤。<br>   对应Value的可选项如下：<br>   4：Ipv4；<br>   6：Ipv6。</li>
	// <li>tagKey<br>   按照【<strong>标签Key</strong>】进行过滤。</li>
	// <li>tagValue<br>   按照【<strong>标签Value</strong>】进行过滤。</li>
	Filters []*QueryCondition `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 数据归属地区，取值有：
	// <li>overseas：全球（除中国大陆地区）数据；</li>
	// <li>mainland：中国大陆地区数据；</li>
	// <li>global：全球数据。</li>不填默认取值为global。
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`
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
	// 查询结果的总条数。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 时序流量数据列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*TimingDataRecord `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 查询的指标，取值有：
	// <li>l7Cache_outFlux：响应流量；</li>
	// <li>l7Cache_request：响应请求数；</li>
	// <li> l7Cache_outBandwidth：响应带宽。</li>
	MetricNames []*string `json:"MetricNames,omitnil,omitempty" name:"MetricNames"`

	// 站点 ID 集合，此参数必填。
	ZoneIds []*string `json:"ZoneIds,omitnil,omitempty" name:"ZoneIds"`

	// 过滤条件，详细的过滤条件如下：
	// <li>domain<br>   按照【<strong>子域名</strong>】进行过滤，子域名形如： test.example.com。<br>   类型：String<br>   必选：否</li>
	// <li>url<br>   按照【<strong>URL</strong>】进行过滤，此参数只支持30天的时间范围，URL形如：/content。<br>   类型：String<br>   必选：否</li>
	// <li>resourceType<br>   按照【<strong>资源类型</strong>】进行过滤，此参数只支持30天的时间范围，资源类型形如：jpg，png。<br>   类型：String<br>   必选：否</li>
	// <li>cacheType<br>   按照【<strong>缓存类型</strong>】进行过滤。<br>   类型：String<br>   必选：否<br>   可选项：<br>   hit：命中缓存；<br>   dynamic：资源不可缓存；<br>   miss：未命中缓存。</li>
	// <li>statusCode<br>   按照【<strong>状态码</strong>】进行过滤，此参数只支持30天的时间范围。<br>   类型：String<br>   必选：否<br>   可选项：<br>   1XX：1xx类型的状态码；<br>   100：100状态码；<br>   101：101状态码；<br>   102：102状态码；<br>   2XX：2xx类型的状态码；<br>   200：200状态码；<br>   201：201状态码；<br>   202：202状态码；<br>   203：203状态码；<br>   204：204状态码；<br>   100：100状态码；<br>   206：206状态码；<br>   207：207状态码；<br>   3XX：3xx类型的状态码；<br>   300：300状态码；<br>   301：301状态码；<br>   302：302状态码；<br>   303：303状态码；<br>   304：304状态码；<br>   305：305状态码；<br>   307：307状态码；<br>   4XX：4xx类型的状态码；<br>   400：400状态码；<br>   401：401状态码；<br>   402：402状态码；<br>   403：403状态码；<br>   404：404状态码；<br>   405：405状态码；<br>   406：406状态码；<br>   407：407状态码；<br>   408：408状态码；<br>   409：409状态码；<br>   410：410状态码；<br>   411：411状态码；<br>   412：412状态码；<br>   412：413状态码；<br>   414：414状态码；<br>   415：415状态码；<br>   416：416状态码；<br>   417：417状态码；<br>   422：422状态码；<br>   423：423状态码；<br>   424：424状态码；<br>   426：426状态码；<br>   451：451状态码；<br>   5XX：5xx类型的状态码；<br>   500：500状态码；<br>   501：501状态码；<br>   502：502状态码；<br>   503：503状态码；<br>   504：504状态码；<br>   505：505状态码；<br>   506：506状态码；<br>   507：507状态码；<br>   510：510状态码；<br>   514：514状态码；<br>   544：544状态码。</li>
	// <li>tagKey<br>   按照【<strong>标签Key</strong>】进行过滤。<br>   类型：String<br>   必选：否</li>
	// <li>tagValue<br>   按照【<strong>标签Value</strong>】进行过滤。<br>   类型：String<br>   必选：否</li>
	Filters []*QueryCondition `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 查询时间粒度，可选的值有：
	// <li>min：1分钟的时间粒度；</li>
	// <li>5min：5分钟的时间粒度；</li>
	// <li>hour：1小时的时间粒度；</li>
	// <li>day：1天的时间粒度。</li>不填将根据开始时间跟结束时间的间距自动推算粒度，具体为：一小时范围内以min粒度查询，两天范围内以5min粒度查询，七天范围内以hour粒度查询，超过七天以day粒度查询。
	Interval *string `json:"Interval,omitnil,omitempty" name:"Interval"`

	// 数据归属地区，取值有：
	// <li>overseas：全球（除中国大陆地区）数据；</li>
	// <li>mainland：中国大陆地区数据；</li>
	// <li>global：全球数据。</li>不填默认取值为global。
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`
}

type DescribeTimingL7CacheDataRequest struct {
	*tchttp.BaseRequest
	
	// 开始时间。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 查询的指标，取值有：
	// <li>l7Cache_outFlux：响应流量；</li>
	// <li>l7Cache_request：响应请求数；</li>
	// <li> l7Cache_outBandwidth：响应带宽。</li>
	MetricNames []*string `json:"MetricNames,omitnil,omitempty" name:"MetricNames"`

	// 站点 ID 集合，此参数必填。
	ZoneIds []*string `json:"ZoneIds,omitnil,omitempty" name:"ZoneIds"`

	// 过滤条件，详细的过滤条件如下：
	// <li>domain<br>   按照【<strong>子域名</strong>】进行过滤，子域名形如： test.example.com。<br>   类型：String<br>   必选：否</li>
	// <li>url<br>   按照【<strong>URL</strong>】进行过滤，此参数只支持30天的时间范围，URL形如：/content。<br>   类型：String<br>   必选：否</li>
	// <li>resourceType<br>   按照【<strong>资源类型</strong>】进行过滤，此参数只支持30天的时间范围，资源类型形如：jpg，png。<br>   类型：String<br>   必选：否</li>
	// <li>cacheType<br>   按照【<strong>缓存类型</strong>】进行过滤。<br>   类型：String<br>   必选：否<br>   可选项：<br>   hit：命中缓存；<br>   dynamic：资源不可缓存；<br>   miss：未命中缓存。</li>
	// <li>statusCode<br>   按照【<strong>状态码</strong>】进行过滤，此参数只支持30天的时间范围。<br>   类型：String<br>   必选：否<br>   可选项：<br>   1XX：1xx类型的状态码；<br>   100：100状态码；<br>   101：101状态码；<br>   102：102状态码；<br>   2XX：2xx类型的状态码；<br>   200：200状态码；<br>   201：201状态码；<br>   202：202状态码；<br>   203：203状态码；<br>   204：204状态码；<br>   100：100状态码；<br>   206：206状态码；<br>   207：207状态码；<br>   3XX：3xx类型的状态码；<br>   300：300状态码；<br>   301：301状态码；<br>   302：302状态码；<br>   303：303状态码；<br>   304：304状态码；<br>   305：305状态码；<br>   307：307状态码；<br>   4XX：4xx类型的状态码；<br>   400：400状态码；<br>   401：401状态码；<br>   402：402状态码；<br>   403：403状态码；<br>   404：404状态码；<br>   405：405状态码；<br>   406：406状态码；<br>   407：407状态码；<br>   408：408状态码；<br>   409：409状态码；<br>   410：410状态码；<br>   411：411状态码；<br>   412：412状态码；<br>   412：413状态码；<br>   414：414状态码；<br>   415：415状态码；<br>   416：416状态码；<br>   417：417状态码；<br>   422：422状态码；<br>   423：423状态码；<br>   424：424状态码；<br>   426：426状态码；<br>   451：451状态码；<br>   5XX：5xx类型的状态码；<br>   500：500状态码；<br>   501：501状态码；<br>   502：502状态码；<br>   503：503状态码；<br>   504：504状态码；<br>   505：505状态码；<br>   506：506状态码；<br>   507：507状态码；<br>   510：510状态码；<br>   514：514状态码；<br>   544：544状态码。</li>
	// <li>tagKey<br>   按照【<strong>标签Key</strong>】进行过滤。<br>   类型：String<br>   必选：否</li>
	// <li>tagValue<br>   按照【<strong>标签Value</strong>】进行过滤。<br>   类型：String<br>   必选：否</li>
	Filters []*QueryCondition `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 查询时间粒度，可选的值有：
	// <li>min：1分钟的时间粒度；</li>
	// <li>5min：5分钟的时间粒度；</li>
	// <li>hour：1小时的时间粒度；</li>
	// <li>day：1天的时间粒度。</li>不填将根据开始时间跟结束时间的间距自动推算粒度，具体为：一小时范围内以min粒度查询，两天范围内以5min粒度查询，七天范围内以hour粒度查询，超过七天以day粒度查询。
	Interval *string `json:"Interval,omitnil,omitempty" name:"Interval"`

	// 数据归属地区，取值有：
	// <li>overseas：全球（除中国大陆地区）数据；</li>
	// <li>mainland：中国大陆地区数据；</li>
	// <li>global：全球数据。</li>不填默认取值为global。
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`
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
	// 查询结果的总条数。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 七层缓存分析时序类流量数据列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*TimingDataRecord `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 查询的指标，取值有：
	// <li> l7Flow_outFlux_country：按国家/地区维度统计流量指标；</li>
	// <li> l7Flow_outFlux_statusCode：按状态码维度统计流量指标；</li>
	// <li> l7Flow_outFlux_domain：按域名维度统计流量指标；</li>
	// <li> l7Flow_outFlux_url：按URL维度统计流量指标; </li>
	// <li> l7Flow_outFlux_resourceType：按资源类型维度统计流量指标；</li>
	// <li> l7Flow_outFlux_sip：按客户端的源IP维度统计流量指标；</li>
	// <li> l7Flow_outFlux_referers：按refer信息维度统计流量指标；</li>
	// <li> l7Flow_outFlux_ua_device：按设备类型维度统计流量指标; </li>
	// <li> l7Flow_outFlux_ua_browser：按浏览器类型维度统计流量指标；</li>
	// <li> l7Flow_outFlux_us_os：按操作系统类型维度统计流量指标；</li>
	// <li> l7Flow_request_country：按国家/地区维度统计请求数指标；</li>
	// <li> l7Flow_request_statusCode：按状态码维度统计请求数指标；</li>
	// <li> l7Flow_request_domain：按域名维度统计请求数指标；</li>
	// <li> l7Flow_request_url：按URL维度统计请求数指标; </li>
	// <li> l7Flow_request_resourceType：按资源类型维度统计请求数指标；</li>
	// <li> l7Flow_request_sip：按客户端的源IP维度统计请求数指标；</li>
	// <li> l7Flow_request_referer：按refer信息维度统计请求数指标；</li>
	// <li> l7Flow_request_ua_device：按设备类型维度统计请求数指标; </li>
	// <li> l7Flow_request_ua_browser：按浏览器类型维度统计请求数指标；</li>
	// <li> l7Flow_request_us_os：按操作系统类型维度统计请求数指标。</li>
	MetricName *string `json:"MetricName,omitnil,omitempty" name:"MetricName"`

	// 站点 ID 集合，此参数必填。
	ZoneIds []*string `json:"ZoneIds,omitnil,omitempty" name:"ZoneIds"`

	// 查询前多少个数据，最大值为1000，不填默认为10， 表示查询前top10的数据。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 过滤条件，详细的过滤条件Key值如下：
	// <li>country<br>   按照【<strong>国家/地区</strong>】进行过滤，国家/地区遵循 <a href="https://baike.baidu.com/item/ISO%203166-1/5269555">ISO 3166</a> 规范。</li>
	// <li>province<br>   按照【<strong>省份</strong>】进行过滤，此参数只支持服务区域为中国大陆。</li>
	// <li>isp<br>   按照【<strong>运营商</strong>】进行过滤，此参数只支持服务区域为中国大陆。<br>   对应的Value可选项如下：<br>   2：中国电信；<br>   26：中国联通；<br>   1046：中国移动；<br>   3947：中国铁通；<br>   38：教育网；<br>   43：长城宽带；<br>   0：其他运营商。</li>
	// <li>domain<br>   按照【<strong>子域名</strong>】进行过滤，子域名形如： test.example.com。</li>
	// <li>url<br>   按照【<strong>URL Path</strong>】进行过滤，URL Path形如：/content或/content/test.jpg。<br>   若只填写url参数，则最多可查询近30天的数据；<br>   若同时填写url+Zonelds参数，则支持的查询数据范围为套餐支持的<a href="https://cloud.tencent.com/document/product/1552/77380#edgeone-.E5.A5.97.E9.A4.90">数据分析最大查询范围</a>与30天两者中的较小值。</li>
	// <li>referer<br>   按照【<strong>Referer头信息</strong>】进行过滤, Referer形如：example.com。<br>   若只填写referer参数，则最多可查询近30天的数据；<br>   若同时填写referer+Zonelds参数，则支持的查询数据范围为套餐支持的<a href="https://cloud.tencent.com/document/product/1552/77380#edgeone-.E5.A5.97.E9.A4.90">数据分析最大查询范围</a>与30天两者中的较小值。</li>
	// <li>resourceType<br>   按照【<strong>资源类型</strong>】进行过滤，资源类型一般是文件后缀，形如: .jpg, .css。<br>   若只填写resourceType参数，则最多可查询近30天的数据；<br>   若同时填写resourceType+Zonelds参数，则支持的查询数据范围为套餐支持的<a href="https://cloud.tencent.com/document/product/1552/77380#edgeone-.E5.A5.97.E9.A4.90">数据分析最大查询范围</a>与30天两者中的较小值。</li>
	// <li>protocol<br>   按照【<strong>HTTP协议版本</strong>】进行过滤。<br>   对应的Value可选项如下：<br>   HTTP/1.0：HTTP 1.0；<br>   HTTP/1.1：HTTP 1.1；<br>   HTTP/2.0：HTTP 2.0；<br>   HTTP/3.0：HTTP 3.0；<br>   WebSocket：WebSocket。</li>
	// <li>socket<br>   按照【<strong>HTTP协议类型</strong>】进行过滤。<br>   对应的Value可选项如下：<br>   HTTP：HTTP 协议；<br>   HTTPS：HTTPS协议；<br>   QUIC：QUIC协议。</li>
	// <li>statusCode<br>   按照【<strong>状态码</strong>】进行过滤。<br>   若只填写statusCode参数，则最多可查询近30天的数据；<br>   若同时填写statusCode+Zonelds参数，则支持的查询数据范围为套餐支持的<a href="https://cloud.tencent.com/document/product/1552/77380#edgeone-.E5.A5.97.E9.A4.90">数据分析最大查询范围</a>与30天两者中的较小值。<br>   对应的Value可选项如下：<br>   1XX：1xx类型的状态码；<br>   100：100状态码；<br>   101：101状态码；<br>   102：102状态码；<br>   2XX：2xx类型的状态码；<br>   200：200状态码；<br>   201：201状态码；<br>   202：202状态码；<br>   203：203状态码；<br>   204：204状态码；<br>   205：205状态码；<br>   206：206状态码；<br>   207：207状态码；<br>   3XX：3xx类型的状态码；<br>   300：300状态码；<br>   301：301状态码；<br>   302：302状态码；<br>   303：303状态码；<br>   304：304状态码；<br>   305：305状态码；<br>   307：307状态码；<br>   4XX：4xx类型的状态码；<br>   400：400状态码；<br>   401：401状态码；<br>   402：402状态码；<br>   403：403状态码；<br>   404：404状态码；<br>   405：405状态码；<br>   406：406状态码；<br>   407：407状态码；<br>   408：408状态码；<br>   409：409状态码；<br>   410：410状态码；<br>   411：411状态码；<br>   412：412状态码；<br>   412：413状态码；<br>   414：414状态码；<br>   415：415状态码；<br>   416：416状态码；<br>   417：417状态码；<br>   422：422状态码；<br>   423：423状态码；<br>   424：424状态码；<br>   426：426状态码；<br>   451：451状态码；<br>   5XX：5xx类型的状态码；<br>   500：500状态码；<br>   501：501状态码；<br>   502：502状态码；<br>   503：503状态码；<br>   504：504状态码；<br>   505：505状态码；<br>   506：506状态码；<br>   507：507状态码；<br>   510：510状态码；<br>   514：514状态码；<br>   544：544状态码。</li>
	// <li>browserType<br>   按照【<strong>浏览器类型</strong>】进行过滤。<br>   若只填写browserType参数，则最多可查询近30天的数据；<br>   若同时填写browserType+Zonelds参数，则支持的查询数据范围为套餐支持的<a href="https://cloud.tencent.com/document/product/1552/77380#edgeone-.E5.A5.97.E9.A4.90">数据分析最大查询范围</a>与30天两者中的较小值。<br>   对应Value的可选项如下：<br>   Firefox：Firefox浏览器；<br>   Chrome：Chrome浏览器；<br>   Safari：Safari浏览器；<br>   Other：其他浏览器类型；<br>   Empty：浏览器类型为空；<br>   Bot：搜索引擎爬虫；<br>   MicrosoftEdge：MicrosoftEdge浏览器；<br>   IE：IE浏览器；<br>   Opera：Opera浏览器；<br>   QQBrowser：QQ浏览器；<br>   LBBrowser：LB浏览器；<br>   MaxthonBrowser：Maxthon浏览器；<br>   SouGouBrowser：搜狗浏览器；<br>   BIDUBrowser：百度浏览器；<br>   TaoBrowser：淘浏览器；<br>   UBrowser：UC浏览器。</li>
	// <li>deviceType<br>   按照【<strong>设备类型</strong>】进行过滤。<br>   若只填写deviceType参数，则最多可查询近30天的数据；<br>   若同时填写deviceType+Zonelds参数，则支持的查询数据范围为套餐支持的<a href="https://cloud.tencent.com/document/product/1552/77380#edgeone-.E5.A5.97.E9.A4.90">数据分析最大查询范围</a>与30天两者中的较小值。<br>   对应Value的可选项如下：<br>   TV：TV设备；<br>   Tablet：Tablet设备；<br>   Mobile：Mobile设备；<br>   Desktop：Desktop设备；<br>   Other：其他设备类型；<br>   Empty：设备类型为空。</li>
	// <li>operatingSystemType<br>   按照【<strong>操作系统类型</strong>】进行过滤。<br>   若只填写operatingSystemType参数，则最多可查询近30天的数据；<br>   若同时填写operatingSystemType+Zonelds参数，则支持的查询数据范围为套餐支持的<a href="https://cloud.tencent.com/document/product/1552/77380#edgeone-.E5.A5.97.E9.A4.90">数据分析最大查询范围</a>与30天两者中的较小值。<br>   对应Value的可选项如下：<br>   Linux：Linux操作系统；<br>   MacOS：MacOs操作系统；<br>   Android：Android操作系统；<br>   IOS：IOS操作系统；<br>   Windows：Windows操作系统；<br>   NetBSD：NetBSD；<br>   ChromiumOS：ChromiumOS；<br>   Bot：搜索引擎爬虫；<br>   Other：其他类型的操作系统；<br>   Empty：操作系统为空。</li>
	// <li>tlsVersion<br>   按照【<strong>TLS版本</strong>】进行过滤。<br>   若只填写tlsVersion参数，则最多可查询近30天的数据；<br>   若同时填写tlsVersion+Zonelds参数，则支持的查询数据范围为套餐支持的<a href="https://cloud.tencent.com/document/product/1552/77380#edgeone-.E5.A5.97.E9.A4.90">数据分析最大查询范围</a>与30天两者中的较小值。<br>   对应Value的可选项如下：<br>   TLS1.0：TLS 1.0；<br>   TLS1.1：TLS 1.1；<br>   TLS1.2：TLS 1.2；<br>   TLS1.3：TLS 1.3。</li>
	// <li>ipVersion<br>   按照【<strong>IP版本</strong>】进行过滤。<br>   对应Value的可选项如下：<br>   4：Ipv4；<br>   6：Ipv6。</li>
	// <li>tagKey<br>   按照【<strong>标签Key</strong>】进行过滤。</li>
	// <li>tagValue<br>   按照【<strong>标签Value</strong>】进行过滤。</li>
	Filters []*QueryCondition `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 查询时间粒度，取值有：
	// <li>min：1分钟；</li>
	// <li>5min：5分钟；</li>
	// <li>hour：1小时；</li>
	// <li>day：1天。</li>不填将根据开始时间跟结束时间的间距自动推算粒度，具体为：一小时范围内以min粒度查询，两天范围内以5min粒度查询，七天范围内以hour粒度查询，超过七天以day粒度查询。
	Interval *string `json:"Interval,omitnil,omitempty" name:"Interval"`

	// 数据归属地区，取值有：
	// <li>overseas：全球（除中国大陆地区）数据；</li>
	// <li>mainland：中国大陆地区数据；</li>
	// <li>global：全球数据。</li>不填默认取值为global。
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`
}

type DescribeTopL7AnalysisDataRequest struct {
	*tchttp.BaseRequest
	
	// 开始时间。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 查询的指标，取值有：
	// <li> l7Flow_outFlux_country：按国家/地区维度统计流量指标；</li>
	// <li> l7Flow_outFlux_statusCode：按状态码维度统计流量指标；</li>
	// <li> l7Flow_outFlux_domain：按域名维度统计流量指标；</li>
	// <li> l7Flow_outFlux_url：按URL维度统计流量指标; </li>
	// <li> l7Flow_outFlux_resourceType：按资源类型维度统计流量指标；</li>
	// <li> l7Flow_outFlux_sip：按客户端的源IP维度统计流量指标；</li>
	// <li> l7Flow_outFlux_referers：按refer信息维度统计流量指标；</li>
	// <li> l7Flow_outFlux_ua_device：按设备类型维度统计流量指标; </li>
	// <li> l7Flow_outFlux_ua_browser：按浏览器类型维度统计流量指标；</li>
	// <li> l7Flow_outFlux_us_os：按操作系统类型维度统计流量指标；</li>
	// <li> l7Flow_request_country：按国家/地区维度统计请求数指标；</li>
	// <li> l7Flow_request_statusCode：按状态码维度统计请求数指标；</li>
	// <li> l7Flow_request_domain：按域名维度统计请求数指标；</li>
	// <li> l7Flow_request_url：按URL维度统计请求数指标; </li>
	// <li> l7Flow_request_resourceType：按资源类型维度统计请求数指标；</li>
	// <li> l7Flow_request_sip：按客户端的源IP维度统计请求数指标；</li>
	// <li> l7Flow_request_referer：按refer信息维度统计请求数指标；</li>
	// <li> l7Flow_request_ua_device：按设备类型维度统计请求数指标; </li>
	// <li> l7Flow_request_ua_browser：按浏览器类型维度统计请求数指标；</li>
	// <li> l7Flow_request_us_os：按操作系统类型维度统计请求数指标。</li>
	MetricName *string `json:"MetricName,omitnil,omitempty" name:"MetricName"`

	// 站点 ID 集合，此参数必填。
	ZoneIds []*string `json:"ZoneIds,omitnil,omitempty" name:"ZoneIds"`

	// 查询前多少个数据，最大值为1000，不填默认为10， 表示查询前top10的数据。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 过滤条件，详细的过滤条件Key值如下：
	// <li>country<br>   按照【<strong>国家/地区</strong>】进行过滤，国家/地区遵循 <a href="https://baike.baidu.com/item/ISO%203166-1/5269555">ISO 3166</a> 规范。</li>
	// <li>province<br>   按照【<strong>省份</strong>】进行过滤，此参数只支持服务区域为中国大陆。</li>
	// <li>isp<br>   按照【<strong>运营商</strong>】进行过滤，此参数只支持服务区域为中国大陆。<br>   对应的Value可选项如下：<br>   2：中国电信；<br>   26：中国联通；<br>   1046：中国移动；<br>   3947：中国铁通；<br>   38：教育网；<br>   43：长城宽带；<br>   0：其他运营商。</li>
	// <li>domain<br>   按照【<strong>子域名</strong>】进行过滤，子域名形如： test.example.com。</li>
	// <li>url<br>   按照【<strong>URL Path</strong>】进行过滤，URL Path形如：/content或/content/test.jpg。<br>   若只填写url参数，则最多可查询近30天的数据；<br>   若同时填写url+Zonelds参数，则支持的查询数据范围为套餐支持的<a href="https://cloud.tencent.com/document/product/1552/77380#edgeone-.E5.A5.97.E9.A4.90">数据分析最大查询范围</a>与30天两者中的较小值。</li>
	// <li>referer<br>   按照【<strong>Referer头信息</strong>】进行过滤, Referer形如：example.com。<br>   若只填写referer参数，则最多可查询近30天的数据；<br>   若同时填写referer+Zonelds参数，则支持的查询数据范围为套餐支持的<a href="https://cloud.tencent.com/document/product/1552/77380#edgeone-.E5.A5.97.E9.A4.90">数据分析最大查询范围</a>与30天两者中的较小值。</li>
	// <li>resourceType<br>   按照【<strong>资源类型</strong>】进行过滤，资源类型一般是文件后缀，形如: .jpg, .css。<br>   若只填写resourceType参数，则最多可查询近30天的数据；<br>   若同时填写resourceType+Zonelds参数，则支持的查询数据范围为套餐支持的<a href="https://cloud.tencent.com/document/product/1552/77380#edgeone-.E5.A5.97.E9.A4.90">数据分析最大查询范围</a>与30天两者中的较小值。</li>
	// <li>protocol<br>   按照【<strong>HTTP协议版本</strong>】进行过滤。<br>   对应的Value可选项如下：<br>   HTTP/1.0：HTTP 1.0；<br>   HTTP/1.1：HTTP 1.1；<br>   HTTP/2.0：HTTP 2.0；<br>   HTTP/3.0：HTTP 3.0；<br>   WebSocket：WebSocket。</li>
	// <li>socket<br>   按照【<strong>HTTP协议类型</strong>】进行过滤。<br>   对应的Value可选项如下：<br>   HTTP：HTTP 协议；<br>   HTTPS：HTTPS协议；<br>   QUIC：QUIC协议。</li>
	// <li>statusCode<br>   按照【<strong>状态码</strong>】进行过滤。<br>   若只填写statusCode参数，则最多可查询近30天的数据；<br>   若同时填写statusCode+Zonelds参数，则支持的查询数据范围为套餐支持的<a href="https://cloud.tencent.com/document/product/1552/77380#edgeone-.E5.A5.97.E9.A4.90">数据分析最大查询范围</a>与30天两者中的较小值。<br>   对应的Value可选项如下：<br>   1XX：1xx类型的状态码；<br>   100：100状态码；<br>   101：101状态码；<br>   102：102状态码；<br>   2XX：2xx类型的状态码；<br>   200：200状态码；<br>   201：201状态码；<br>   202：202状态码；<br>   203：203状态码；<br>   204：204状态码；<br>   205：205状态码；<br>   206：206状态码；<br>   207：207状态码；<br>   3XX：3xx类型的状态码；<br>   300：300状态码；<br>   301：301状态码；<br>   302：302状态码；<br>   303：303状态码；<br>   304：304状态码；<br>   305：305状态码；<br>   307：307状态码；<br>   4XX：4xx类型的状态码；<br>   400：400状态码；<br>   401：401状态码；<br>   402：402状态码；<br>   403：403状态码；<br>   404：404状态码；<br>   405：405状态码；<br>   406：406状态码；<br>   407：407状态码；<br>   408：408状态码；<br>   409：409状态码；<br>   410：410状态码；<br>   411：411状态码；<br>   412：412状态码；<br>   412：413状态码；<br>   414：414状态码；<br>   415：415状态码；<br>   416：416状态码；<br>   417：417状态码；<br>   422：422状态码；<br>   423：423状态码；<br>   424：424状态码；<br>   426：426状态码；<br>   451：451状态码；<br>   5XX：5xx类型的状态码；<br>   500：500状态码；<br>   501：501状态码；<br>   502：502状态码；<br>   503：503状态码；<br>   504：504状态码；<br>   505：505状态码；<br>   506：506状态码；<br>   507：507状态码；<br>   510：510状态码；<br>   514：514状态码；<br>   544：544状态码。</li>
	// <li>browserType<br>   按照【<strong>浏览器类型</strong>】进行过滤。<br>   若只填写browserType参数，则最多可查询近30天的数据；<br>   若同时填写browserType+Zonelds参数，则支持的查询数据范围为套餐支持的<a href="https://cloud.tencent.com/document/product/1552/77380#edgeone-.E5.A5.97.E9.A4.90">数据分析最大查询范围</a>与30天两者中的较小值。<br>   对应Value的可选项如下：<br>   Firefox：Firefox浏览器；<br>   Chrome：Chrome浏览器；<br>   Safari：Safari浏览器；<br>   Other：其他浏览器类型；<br>   Empty：浏览器类型为空；<br>   Bot：搜索引擎爬虫；<br>   MicrosoftEdge：MicrosoftEdge浏览器；<br>   IE：IE浏览器；<br>   Opera：Opera浏览器；<br>   QQBrowser：QQ浏览器；<br>   LBBrowser：LB浏览器；<br>   MaxthonBrowser：Maxthon浏览器；<br>   SouGouBrowser：搜狗浏览器；<br>   BIDUBrowser：百度浏览器；<br>   TaoBrowser：淘浏览器；<br>   UBrowser：UC浏览器。</li>
	// <li>deviceType<br>   按照【<strong>设备类型</strong>】进行过滤。<br>   若只填写deviceType参数，则最多可查询近30天的数据；<br>   若同时填写deviceType+Zonelds参数，则支持的查询数据范围为套餐支持的<a href="https://cloud.tencent.com/document/product/1552/77380#edgeone-.E5.A5.97.E9.A4.90">数据分析最大查询范围</a>与30天两者中的较小值。<br>   对应Value的可选项如下：<br>   TV：TV设备；<br>   Tablet：Tablet设备；<br>   Mobile：Mobile设备；<br>   Desktop：Desktop设备；<br>   Other：其他设备类型；<br>   Empty：设备类型为空。</li>
	// <li>operatingSystemType<br>   按照【<strong>操作系统类型</strong>】进行过滤。<br>   若只填写operatingSystemType参数，则最多可查询近30天的数据；<br>   若同时填写operatingSystemType+Zonelds参数，则支持的查询数据范围为套餐支持的<a href="https://cloud.tencent.com/document/product/1552/77380#edgeone-.E5.A5.97.E9.A4.90">数据分析最大查询范围</a>与30天两者中的较小值。<br>   对应Value的可选项如下：<br>   Linux：Linux操作系统；<br>   MacOS：MacOs操作系统；<br>   Android：Android操作系统；<br>   IOS：IOS操作系统；<br>   Windows：Windows操作系统；<br>   NetBSD：NetBSD；<br>   ChromiumOS：ChromiumOS；<br>   Bot：搜索引擎爬虫；<br>   Other：其他类型的操作系统；<br>   Empty：操作系统为空。</li>
	// <li>tlsVersion<br>   按照【<strong>TLS版本</strong>】进行过滤。<br>   若只填写tlsVersion参数，则最多可查询近30天的数据；<br>   若同时填写tlsVersion+Zonelds参数，则支持的查询数据范围为套餐支持的<a href="https://cloud.tencent.com/document/product/1552/77380#edgeone-.E5.A5.97.E9.A4.90">数据分析最大查询范围</a>与30天两者中的较小值。<br>   对应Value的可选项如下：<br>   TLS1.0：TLS 1.0；<br>   TLS1.1：TLS 1.1；<br>   TLS1.2：TLS 1.2；<br>   TLS1.3：TLS 1.3。</li>
	// <li>ipVersion<br>   按照【<strong>IP版本</strong>】进行过滤。<br>   对应Value的可选项如下：<br>   4：Ipv4；<br>   6：Ipv6。</li>
	// <li>tagKey<br>   按照【<strong>标签Key</strong>】进行过滤。</li>
	// <li>tagValue<br>   按照【<strong>标签Value</strong>】进行过滤。</li>
	Filters []*QueryCondition `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 查询时间粒度，取值有：
	// <li>min：1分钟；</li>
	// <li>5min：5分钟；</li>
	// <li>hour：1小时；</li>
	// <li>day：1天。</li>不填将根据开始时间跟结束时间的间距自动推算粒度，具体为：一小时范围内以min粒度查询，两天范围内以5min粒度查询，七天范围内以hour粒度查询，超过七天以day粒度查询。
	Interval *string `json:"Interval,omitnil,omitempty" name:"Interval"`

	// 数据归属地区，取值有：
	// <li>overseas：全球（除中国大陆地区）数据；</li>
	// <li>mainland：中国大陆地区数据；</li>
	// <li>global：全球数据。</li>不填默认取值为global。
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`
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
	// 查询结果的总条数。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 七层流量前topN数据列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*TopDataRecord `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 查询的指标，取值有：
	// <li> l7Cache_outFlux_domain：host/域名；</li>
	// <li> l7Cache_outFlux_url：url地址；</li>
	// <li> l7Cache_outFlux_resourceType：资源类型；</li>
	// <li> l7Cache_outFlux_statusCode：状态码。</li>
	MetricName *string `json:"MetricName,omitnil,omitempty" name:"MetricName"`

	// 站点 ID 集合，此参数必填。
	ZoneIds []*string `json:"ZoneIds,omitnil,omitempty" name:"ZoneIds"`

	// 查询前多少个数据，最大值为1000，不填默认为10， 表示查询前top 10的数据。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 过滤条件，详细的过滤条件如下：
	// <li>domain<br>   按照【<strong>子域名</strong>】进行过滤，子域名形如： test.example.com。<br>   类型：String<br>   必选：否</li>
	// <li>url<br>   按照【<strong>URL</strong>】进行过滤，此参数只支持30天的时间范围，URL形如：/content。<br>   类型：String<br>   必选：否</li>
	// <li>resourceType<br>   按照【<strong>资源类型</strong>】进行过滤，此参数只支持30天的时间范围，资源类型形如：jpg，png。<br>   类型：String<br>   必选：否</li>
	// <li>cacheType<br>   按照【<strong>缓存类型</strong>】进行过滤。<br>   类型：String<br>   必选：否<br>   可选项：<br>   hit：命中缓存；<br>   dynamic：资源不可缓存；<br>   miss：未命中缓存。</li>
	// <li>statusCode<br>   按照【<strong>状态码</strong>】进行过滤，此参数只支持30天的时间范围。<br>   类型：String<br>   必选：否<br>   可选项：<br>   1XX：1xx类型的状态码；<br>   100：100状态码；<br>   101：101状态码；<br>   102：102状态码；<br>   2XX：2xx类型的状态码；<br>   200：200状态码；<br>   201：201状态码；<br>   202：202状态码；<br>   203：203状态码；<br>   204：204状态码；<br>   100：100状态码；<br>   206：206状态码；<br>   207：207状态码；<br>   3XX：3xx类型的状态码；<br>   300：300状态码；<br>   301：301状态码；<br>   302：302状态码；<br>   303：303状态码；<br>   304：304状态码；<br>   305：305状态码；<br>   307：307状态码；<br>   4XX：4xx类型的状态码；<br>   400：400状态码；<br>   401：401状态码；<br>   402：402状态码；<br>   403：403状态码；<br>   404：404状态码；<br>   405：405状态码；<br>   406：406状态码；<br>   407：407状态码；<br>   408：408状态码；<br>   409：409状态码；<br>   410：410状态码；<br>   411：411状态码；<br>   412：412状态码；<br>   412：413状态码；<br>   414：414状态码；<br>   415：415状态码；<br>   416：416状态码；<br>   417：417状态码；<br>   422：422状态码；<br>   423：423状态码；<br>   424：424状态码；<br>   426：426状态码；<br>   451：451状态码；<br>   5XX：5xx类型的状态码；<br>   500：500状态码；<br>   501：501状态码；<br>   502：502状态码；<br>   503：503状态码；<br>   504：504状态码；<br>   505：505状态码；<br>   506：506状态码；<br>   507：507状态码；<br>   510：510状态码；<br>   514：514状态码；<br>   544：544状态码。</li>
	// <li>tagKey<br>   按照【<strong>标签Key</strong>】进行过滤。<br>   类型：String<br>   必选：否</li>
	// <li>tagValue<br>   按照【<strong>标签Value</strong>】进行过滤。<br>   类型：String<br>   必选：否</li>
	Filters []*QueryCondition `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 查询时间粒度，取值有：
	// <li>min: 1分钟；</li>
	// <li>5min: 5分钟；</li>
	// <li>hour: 1小时；</li>
	// <li>day: 1天。</li>不填将根据开始时间跟结束时间的间距自动推算粒度，具体为：一小时范围内以min粒度查询，两天范围内以5min粒度查询，七天范围内以hour粒度查询，超过七天以day粒度查询。
	Interval *string `json:"Interval,omitnil,omitempty" name:"Interval"`

	// 数据归属地区，取值有：
	// <li>overseas：全球（除中国大陆地区）数据；</li>
	// <li>mainland：中国大陆地区数据；</li>
	// <li>global：全球数据。</li>不填默认取值为global。
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`
}

type DescribeTopL7CacheDataRequest struct {
	*tchttp.BaseRequest
	
	// 开始时间。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 查询的指标，取值有：
	// <li> l7Cache_outFlux_domain：host/域名；</li>
	// <li> l7Cache_outFlux_url：url地址；</li>
	// <li> l7Cache_outFlux_resourceType：资源类型；</li>
	// <li> l7Cache_outFlux_statusCode：状态码。</li>
	MetricName *string `json:"MetricName,omitnil,omitempty" name:"MetricName"`

	// 站点 ID 集合，此参数必填。
	ZoneIds []*string `json:"ZoneIds,omitnil,omitempty" name:"ZoneIds"`

	// 查询前多少个数据，最大值为1000，不填默认为10， 表示查询前top 10的数据。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 过滤条件，详细的过滤条件如下：
	// <li>domain<br>   按照【<strong>子域名</strong>】进行过滤，子域名形如： test.example.com。<br>   类型：String<br>   必选：否</li>
	// <li>url<br>   按照【<strong>URL</strong>】进行过滤，此参数只支持30天的时间范围，URL形如：/content。<br>   类型：String<br>   必选：否</li>
	// <li>resourceType<br>   按照【<strong>资源类型</strong>】进行过滤，此参数只支持30天的时间范围，资源类型形如：jpg，png。<br>   类型：String<br>   必选：否</li>
	// <li>cacheType<br>   按照【<strong>缓存类型</strong>】进行过滤。<br>   类型：String<br>   必选：否<br>   可选项：<br>   hit：命中缓存；<br>   dynamic：资源不可缓存；<br>   miss：未命中缓存。</li>
	// <li>statusCode<br>   按照【<strong>状态码</strong>】进行过滤，此参数只支持30天的时间范围。<br>   类型：String<br>   必选：否<br>   可选项：<br>   1XX：1xx类型的状态码；<br>   100：100状态码；<br>   101：101状态码；<br>   102：102状态码；<br>   2XX：2xx类型的状态码；<br>   200：200状态码；<br>   201：201状态码；<br>   202：202状态码；<br>   203：203状态码；<br>   204：204状态码；<br>   100：100状态码；<br>   206：206状态码；<br>   207：207状态码；<br>   3XX：3xx类型的状态码；<br>   300：300状态码；<br>   301：301状态码；<br>   302：302状态码；<br>   303：303状态码；<br>   304：304状态码；<br>   305：305状态码；<br>   307：307状态码；<br>   4XX：4xx类型的状态码；<br>   400：400状态码；<br>   401：401状态码；<br>   402：402状态码；<br>   403：403状态码；<br>   404：404状态码；<br>   405：405状态码；<br>   406：406状态码；<br>   407：407状态码；<br>   408：408状态码；<br>   409：409状态码；<br>   410：410状态码；<br>   411：411状态码；<br>   412：412状态码；<br>   412：413状态码；<br>   414：414状态码；<br>   415：415状态码；<br>   416：416状态码；<br>   417：417状态码；<br>   422：422状态码；<br>   423：423状态码；<br>   424：424状态码；<br>   426：426状态码；<br>   451：451状态码；<br>   5XX：5xx类型的状态码；<br>   500：500状态码；<br>   501：501状态码；<br>   502：502状态码；<br>   503：503状态码；<br>   504：504状态码；<br>   505：505状态码；<br>   506：506状态码；<br>   507：507状态码；<br>   510：510状态码；<br>   514：514状态码；<br>   544：544状态码。</li>
	// <li>tagKey<br>   按照【<strong>标签Key</strong>】进行过滤。<br>   类型：String<br>   必选：否</li>
	// <li>tagValue<br>   按照【<strong>标签Value</strong>】进行过滤。<br>   类型：String<br>   必选：否</li>
	Filters []*QueryCondition `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 查询时间粒度，取值有：
	// <li>min: 1分钟；</li>
	// <li>5min: 5分钟；</li>
	// <li>hour: 1小时；</li>
	// <li>day: 1天。</li>不填将根据开始时间跟结束时间的间距自动推算粒度，具体为：一小时范围内以min粒度查询，两天范围内以5min粒度查询，七天范围内以hour粒度查询，超过七天以day粒度查询。
	Interval *string `json:"Interval,omitnil,omitempty" name:"Interval"`

	// 数据归属地区，取值有：
	// <li>overseas：全球（除中国大陆地区）数据；</li>
	// <li>mainland：中国大陆地区数据；</li>
	// <li>global：全球数据。</li>不填默认取值为global。
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`
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
	// 查询结果的总条数。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 七层缓存TopN流量数据列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*TopDataRecord `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DescribeZoneSettingRequestParams struct {
	// 站点ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`
}

type DescribeZoneSettingRequest struct {
	*tchttp.BaseRequest
	
	// 站点ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`
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
	ZoneSetting *ZoneSetting `json:"ZoneSetting,omitnil,omitempty" name:"ZoneSetting"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页查询限制数目。默认值：20，最大值：100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 过滤条件，Filters.Values 的上限为 20。该参数不填写时，返回当前 appid 下有权限的所有站点信息。详细的过滤条件如下：
	// <li>zone-name：按照站点名称进行过滤；</li><li>zone-id：按照站点 ID进行过滤。站点 ID 形如：zone-2noz78a8ev6k；</li><li>status：按照站点状态进行过滤；</li><li>tag-key：按照标签键进行过滤；</li><li>tag-value： 按照标签值进行过滤。</li>模糊查询时仅支持过滤字段名为 zone-name。
	Filters []*AdvancedFilter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 可根据该字段对返回结果进行排序，取值有：
	// <li> type：接入类型；</li>
	// <li> area：加速区域；</li>
	// <li> create-time：创建时间；</li>
	// <li> zone-name：站点名称；</li>
	// <li> use-time：最近使用时间；</li>
	// <li> active-status：生效状态。</li>不填写时对返回结果默认按照 create-time 排序。
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// 排序方向，如果是字段值为数字，则根据数字大小排序；如果字段值为文本，则根据 ascill 码的大小排序。取值有：
	// <li> asc：从小到大排序；</li>
	// <li> desc：从大到小排序。</li>不填写使用默认值 desc。
	Direction *string `json:"Direction,omitnil,omitempty" name:"Direction"`
}

type DescribeZonesRequest struct {
	*tchttp.BaseRequest
	
	// 分页查询偏移量。默认值：0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页查询限制数目。默认值：20，最大值：100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 过滤条件，Filters.Values 的上限为 20。该参数不填写时，返回当前 appid 下有权限的所有站点信息。详细的过滤条件如下：
	// <li>zone-name：按照站点名称进行过滤；</li><li>zone-id：按照站点 ID进行过滤。站点 ID 形如：zone-2noz78a8ev6k；</li><li>status：按照站点状态进行过滤；</li><li>tag-key：按照标签键进行过滤；</li><li>tag-value： 按照标签值进行过滤。</li>模糊查询时仅支持过滤字段名为 zone-name。
	Filters []*AdvancedFilter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 可根据该字段对返回结果进行排序，取值有：
	// <li> type：接入类型；</li>
	// <li> area：加速区域；</li>
	// <li> create-time：创建时间；</li>
	// <li> zone-name：站点名称；</li>
	// <li> use-time：最近使用时间；</li>
	// <li> active-status：生效状态。</li>不填写时对返回结果默认按照 create-time 排序。
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// 排序方向，如果是字段值为数字，则根据数字大小排序；如果字段值为文本，则根据 ascill 码的大小排序。取值有：
	// <li> asc：从小到大排序；</li>
	// <li> desc：从大到小排序。</li>不填写使用默认值 desc。
	Direction *string `json:"Direction,omitnil,omitempty" name:"Direction"`
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
	delete(f, "Order")
	delete(f, "Direction")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeZonesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeZonesResponseParams struct {
	// 符合条件的站点个数。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 站点详细信息。
	Zones []*Zone `json:"Zones,omitnil,omitempty" name:"Zones"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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

// Predefined struct for user
type DestroyPlanRequestParams struct {
	// 套餐 ID，形如 edgeone-2wdo315m2y4c。
	PlanId *string `json:"PlanId,omitnil,omitempty" name:"PlanId"`
}

type DestroyPlanRequest struct {
	*tchttp.BaseRequest
	
	// 套餐 ID，形如 edgeone-2wdo315m2y4c。
	PlanId *string `json:"PlanId,omitnil,omitempty" name:"PlanId"`
}

func (r *DestroyPlanRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DestroyPlanRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PlanId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DestroyPlanRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DestroyPlanResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DestroyPlanResponse struct {
	*tchttp.BaseResponse
	Response *DestroyPlanResponseParams `json:"Response"`
}

func (r *DestroyPlanResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DestroyPlanResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DetailHost struct {
	// 站点ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 加速服务状态，取值为：
	// <li> process：部署中；</li>
	// <li> online：已启动；</li>
	// <li> offline：已关闭。</li>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 域名。
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`

	// 站点名称。
	ZoneName *string `json:"ZoneName,omitnil,omitempty" name:"ZoneName"`

	// 分配的Cname域名
	Cname *string `json:"Cname,omitnil,omitempty" name:"Cname"`

	// 资源ID。
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 实例ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 锁状态。
	Lock *int64 `json:"Lock,omitnil,omitempty" name:"Lock"`

	// 域名状态类型。
	Mode *int64 `json:"Mode,omitnil,omitempty" name:"Mode"`

	// 域名加速地域，取值有：
	// <li> global：全球；</li>
	// <li> mainland：中国大陆；</li>
	// <li> overseas：境外区域。</li>
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 加速类型配置项。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AccelerateType *AccelerateType `json:"AccelerateType,omitnil,omitempty" name:"AccelerateType"`

	// Https配置项。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Https *Https `json:"Https,omitnil,omitempty" name:"Https"`

	// 缓存配置项。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CacheConfig *CacheConfig `json:"CacheConfig,omitnil,omitempty" name:"CacheConfig"`

	// 源站配置项。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Origin *Origin `json:"Origin,omitnil,omitempty" name:"Origin"`

	// 安全类型。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SecurityType *SecurityType `json:"SecurityType,omitnil,omitempty" name:"SecurityType"`

	// 缓存键配置项。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CacheKey *CacheKey `json:"CacheKey,omitnil,omitempty" name:"CacheKey"`

	// 智能压缩配置项。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Compression *Compression `json:"Compression,omitnil,omitempty" name:"Compression"`

	// Waf防护配置项。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Waf *Waf `json:"Waf,omitnil,omitempty" name:"Waf"`

	// CC防护配置项。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CC *CC `json:"CC,omitnil,omitempty" name:"CC"`

	// DDoS防护配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DDoS *DDoS `json:"DDoS,omitnil,omitempty" name:"DDoS"`

	// 智能路由配置项。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SmartRouting *SmartRouting `json:"SmartRouting,omitnil,omitempty" name:"SmartRouting"`

	// Ipv6访问配置项。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Ipv6 *Ipv6 `json:"Ipv6,omitnil,omitempty" name:"Ipv6"`

	// 回源时是否携带客户端IP所属地域信息的配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClientIpCountry *ClientIpCountry `json:"ClientIpCountry,omitnil,omitempty" name:"ClientIpCountry"`
}

type DiffIPWhitelist struct {
	// 最新IP白名单列表。
	LatestIPWhitelist *IPWhitelist `json:"LatestIPWhitelist,omitnil,omitempty" name:"LatestIPWhitelist"`

	// 最新IP白名单列表相比于当前IP白名单列表，新增部分。
	AddedIPWhitelist *IPWhitelist `json:"AddedIPWhitelist,omitnil,omitempty" name:"AddedIPWhitelist"`

	// 最新IP白名单列表相比于当前IP白名单列表，删减部分。
	RemovedIPWhitelist *IPWhitelist `json:"RemovedIPWhitelist,omitnil,omitempty" name:"RemovedIPWhitelist"`

	// 最新IP白名单列表相比于当前IP白名单列表，不变部分。
	NoChangeIPWhitelist *IPWhitelist `json:"NoChangeIPWhitelist,omitnil,omitempty" name:"NoChangeIPWhitelist"`
}

type DnsVerification struct {
	// 主机记录。
	Subdomain *string `json:"Subdomain,omitnil,omitempty" name:"Subdomain"`

	// 记录类型。
	RecordType *string `json:"RecordType,omitnil,omitempty" name:"RecordType"`

	// 记录值。
	RecordValue *string `json:"RecordValue,omitnil,omitempty" name:"RecordValue"`
}

// Predefined struct for user
type DownloadL4LogsRequestParams struct {
	// 开始时间。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 站点 ID 集合，此参数必填。
	ZoneIds []*string `json:"ZoneIds,omitnil,omitempty" name:"ZoneIds"`

	// 四层实例 ID 集合。
	ProxyIds []*string `json:"ProxyIds,omitnil,omitempty" name:"ProxyIds"`

	// 分页查询的限制数目，默认值为 20，最大查询条目为 300。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页的偏移量，默认值为 0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type DownloadL4LogsRequest struct {
	*tchttp.BaseRequest
	
	// 开始时间。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 站点 ID 集合，此参数必填。
	ZoneIds []*string `json:"ZoneIds,omitnil,omitempty" name:"ZoneIds"`

	// 四层实例 ID 集合。
	ProxyIds []*string `json:"ProxyIds,omitnil,omitempty" name:"ProxyIds"`

	// 分页查询的限制数目，默认值为 20，最大查询条目为 300。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页的偏移量，默认值为 0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
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
	// 查询结果的总条数。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 四层离线日志数据列表。
	Data []*L4OfflineLog `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 站点ID集合，此参数必填。
	ZoneIds []*string `json:"ZoneIds,omitnil,omitempty" name:"ZoneIds"`

	// 子域名集合，不填默认选择全部子域名。
	Domains []*string `json:"Domains,omitnil,omitempty" name:"Domains"`

	// 分页查询的限制数目，默认值为 20，最大查询条目为 300。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页的偏移量，默认值为 0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type DownloadL7LogsRequest struct {
	*tchttp.BaseRequest
	
	// 开始时间。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 站点ID集合，此参数必填。
	ZoneIds []*string `json:"ZoneIds,omitnil,omitempty" name:"ZoneIds"`

	// 子域名集合，不填默认选择全部子域名。
	Domains []*string `json:"Domains,omitnil,omitempty" name:"Domains"`

	// 分页查询的限制数目，默认值为 20，最大查询条目为 300。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页的偏移量，默认值为 0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
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
	// 查询结果的总条数。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 七层离线日志数据列表。
	Data []*L7OfflineLog `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// Waf(托管规则)模块的拦截页面配置。如果为null，默认使用历史配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	WafDropPageDetail *DropPageDetail `json:"WafDropPageDetail,omitnil,omitempty" name:"WafDropPageDetail"`

	// 自定义页面的拦截页面配置。如果为null，默认使用历史配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AclDropPageDetail *DropPageDetail `json:"AclDropPageDetail,omitnil,omitempty" name:"AclDropPageDetail"`
}

type DropPageDetail struct {
	// 拦截页面的唯一 Id。系统默认包含一个自带拦截页面，Id 值为0。
	// 该 Id 可通过创建拦截页面接口进行上传获取。如传入0，代表使用系统默认拦截页面。该参数已废弃。
	PageId *int64 `json:"PageId,omitnil,omitempty" name:"PageId"`

	// 拦截页面的 HTTP 状态码。状态码取值：100～600，不支持 3xx 状态码。托管规则拦截页面默认：566，安全防护（除托管规则外）拦截页面默认：567.
	StatusCode *int64 `json:"StatusCode,omitnil,omitempty" name:"StatusCode"`

	// 页面文件名或 url。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 页面的类型，取值有：
	// <li>page：指定页面。</li>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 自定义响应 Id。该 Id 可通过查询自定义错误页列表接口获取。默认值为default，使用系统默认页面。Type 类型是 page 时必填，且不能为空。
	CustomResponseId *string `json:"CustomResponseId,omitnil,omitempty" name:"CustomResponseId"`
}

type EntityStatus struct {
	// 实例名，现在只有子域名。
	Entity *string `json:"Entity,omitnil,omitempty" name:"Entity"`

	// 实例配置下发状态，取值有：
	// <li>online：配置已生效；</li><li>fail：配置失败；</li><li> process：配置下发中。</li>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 实例配置下发信息提示。
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`
}

type EnvInfo struct {
	// 环境 ID。
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 环境类型，取值有：
	// <li>production: 生产环境；</li><li>staging: 测试环境。</li>
	EnvType *string `json:"EnvType,omitnil,omitempty" name:"EnvType"`

	// 环境状态，取值有：
	// <li>creating：创建中；</li>
	// <li>running：稳定运行中，可进行版本变更；</li>
	// <li>version_deploying：版本部署中，不能进行新的变更。</li>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 当前环境的配置生效范围：
	// <li>当 EnvType 取值为 production 时，该参数值为 ["ALL"]，代表全网生效；</li>
	// <li>当 EnvType 取值为 staging 时，会返回测试节点 IP，可用于绑定 host 测试。</li>
	Scope []*string `json:"Scope,omitnil,omitempty" name:"Scope"`

	// 当前环境中各配置组实际生效的版本，根据 Status 的取值有以下两种情况：
	// <li>当 Status 取值为 version_deploying 时，本字段返回的值为执行变更动作之前生效的版本，即新版本部署期间，实际生效的版本为执行变更动作之前的版本；</li>
	// <li>当 Status 取值为 running 时，本字段返回的值即为当前实际生效的版本。</li>
	CurrentConfigGroupVersionInfos []*ConfigGroupVersionInfo `json:"CurrentConfigGroupVersionInfos,omitnil,omitempty" name:"CurrentConfigGroupVersionInfos"`

	// 创建时间。时间为世界标准时间（UTC）， 遵循 ISO 8601 标准的日期和时间格式。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 更新时间。时间为世界标准时间（UTC）， 遵循 ISO 8601 标准的日期和时间格式。
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`
}

type ErrorPageReference struct {
	// 引用的业务 ID，如自定义拦截规则 ID。
	BusinessId *string `json:"BusinessId,omitnil,omitempty" name:"BusinessId"`
}

type ExceptConfig struct {
	// 配置开关，取值有：
	// <li>on：开启；</li>
	// <li>off：关闭。</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// 例外规则详情。如果为null，默认使用历史配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExceptUserRules []*ExceptUserRule `json:"ExceptUserRules,omitnil,omitempty" name:"ExceptUserRules"`
}

type ExceptUserRule struct {
	// 规则名称，不可使用中文。
	RuleName *string `json:"RuleName,omitnil,omitempty" name:"RuleName"`

	// 规则的处置方式，当前仅支持skip：跳过全部托管规则。
	Action *string `json:"Action,omitnil,omitempty" name:"Action"`

	// 规则生效状态，取值有：
	// <li>on：生效；</li>
	// <li>off：失效。</li>
	RuleStatus *string `json:"RuleStatus,omitnil,omitempty" name:"RuleStatus"`

	// 规则ID。仅出参使用。默认由底层生成。
	RuleID *int64 `json:"RuleID,omitnil,omitempty" name:"RuleID"`

	// 更新时间，如果为null，默认由底层按当前时间生成。
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 匹配条件。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExceptUserRuleConditions []*ExceptUserRuleCondition `json:"ExceptUserRuleConditions,omitnil,omitempty" name:"ExceptUserRuleConditions"`

	// 规则生效的范围。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExceptUserRuleScope *ExceptUserRuleScope `json:"ExceptUserRuleScope,omitnil,omitempty" name:"ExceptUserRuleScope"`

	// 优先级，取值范围0-100。如果为null，默认由底层设置为0。
	RulePriority *int64 `json:"RulePriority,omitnil,omitempty" name:"RulePriority"`
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
	MatchFrom *string `json:"MatchFrom,omitnil,omitempty" name:"MatchFrom"`

	// 匹配项的参数。仅当 MatchFrom 为 header 时，可以使用本参数，值可填入 header 的 key 作为参数。
	MatchParam *string `json:"MatchParam,omitnil,omitempty" name:"MatchParam"`

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
	Operator *string `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 匹配值。
	MatchContent *string `json:"MatchContent,omitnil,omitempty" name:"MatchContent"`
}

type ExceptUserRuleScope struct {
	// 例外规则类型。其中complete模式代表全量数据进行例外，partial模式代表可选择指定模块指定字段进行例外，该字段取值有：
	// <li>complete：完全跳过模式；</li>
	// <li>partial：部分跳过模式。</li>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 生效的模块，该字段取值有：
	// <li>waf：托管规则；</li>
	// <li>rate：速率限制；</li>
	// <li>acl：自定义规则；</li>
	// <li>cc：cc攻击防护；</li>
	// <li>bot：Bot防护。</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Modules []*string `json:"Modules,omitnil,omitempty" name:"Modules"`

	// 跳过部分规则ID的例外规则详情。如果为null，默认使用历史配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PartialModules []*PartialModule `json:"PartialModules,omitnil,omitempty" name:"PartialModules"`

	// 跳过具体字段不去扫描的例外规则详情。如果为null，默认使用历史配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SkipConditions []*SkipCondition `json:"SkipConditions,omitnil,omitempty" name:"SkipConditions"`
}

type FailReason struct {
	// 失败原因。
	Reason *string `json:"Reason,omitnil,omitempty" name:"Reason"`

	// 处理失败的资源列表。
	Targets []*string `json:"Targets,omitnil,omitempty" name:"Targets"`
}

type FileAscriptionInfo struct {
	// 文件校验目录。
	IdentifyPath *string `json:"IdentifyPath,omitnil,omitempty" name:"IdentifyPath"`

	// 文件校验内容。
	IdentifyContent *string `json:"IdentifyContent,omitnil,omitempty" name:"IdentifyContent"`
}

type FileVerification struct {
	// EdgeOne 后台服务器将通过 Scheme + Host + URL Path 的格式（例如 https://www.example.com/.well-known/teo-verification/z12h416twn.txt）获取文件验证信息。该字段为您需要创建的 URL Path 部分。
	Path *string `json:"Path,omitnil,omitempty" name:"Path"`

	// 验证文件的内容。该字段的内容需要您填写至 Path 字段返回的 txt 文件中。
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`
}

type Filter struct {
	// 需要过滤的字段。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 字段的过滤值。
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`
}

type FirstPartConfig struct {
	// 开关，取值有：
	// <li>on：开启；</li>
	// <li>off：关闭。</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// 首段包的统计时长，单位是秒，即期望首段包的统计时长是多少，默认5秒。
	// 注意：此字段可能返回 null，表示取不到有效值。
	StatTime *uint64 `json:"StatTime,omitnil,omitempty" name:"StatTime"`
}

type FollowOrigin struct {
	// 遵循源站配置开关，取值有：
	// <li>on：开启；</li>
	// <li>off：关闭。</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// 源站未返回 Cache-Control 头时, 设置默认的缓存时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	DefaultCacheTime *int64 `json:"DefaultCacheTime,omitnil,omitempty" name:"DefaultCacheTime"`

	// 源站未返回 Cache-Control 头时, 设置缓存/不缓存
	// 注意：此字段可能返回 null，表示取不到有效值。
	DefaultCache *string `json:"DefaultCache,omitnil,omitempty" name:"DefaultCache"`

	// 源站未返回 Cache-Control 头时, 使用/不使用默认缓存策略
	// 注意：此字段可能返回 null，表示取不到有效值。
	DefaultCacheStrategy *string `json:"DefaultCacheStrategy,omitnil,omitempty" name:"DefaultCacheStrategy"`
}

type ForceRedirect struct {
	// 访问强制跳转配置开关，取值有：
	// <li>on：开启；</li>
	// <li>off：关闭。</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// 重定向状态码，取值有：
	// <li>301：301跳转；</li>
	// <li>302：302跳转。</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	RedirectStatusCode *int64 `json:"RedirectStatusCode,omitnil,omitempty" name:"RedirectStatusCode"`
}

type Function struct {
	// 函数 ID。
	FunctionId *string `json:"FunctionId,omitnil,omitempty" name:"FunctionId"`

	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 函数名字。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 函数描述。
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 函数内容。
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`

	// 函数默认域名。
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 创建时间。时间为世界标准时间（UTC）， 遵循 ISO 8601 标准的日期和时间格式。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 修改时间。时间为世界标准时间（UTC）， 遵循 ISO 8601 标准的日期和时间格式。
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`
}

type FunctionEnvironmentVariable struct {
	// 变量的名称，限制只能包含大小写字母、数字，特殊字符仅支持 @ . - _ ，最大 64 个字节，不支持重复。
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 变量的值，限制最大 5000 字节，默认值为空。
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`

	// 变量的类型，取值有：
	// <li>string：字符串类型；</li>
	// <li>json：json 对象类型。</li>默认值为：string。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

type FunctionRule struct {
	// 规则ID。
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// 规则条件列表，列表项之间为或关系。
	FunctionRuleConditions []*FunctionRuleCondition `json:"FunctionRuleConditions,omitnil,omitempty" name:"FunctionRuleConditions"`

	// 函数 ID，命中触发规则条件后执行的函数。
	FunctionId *string `json:"FunctionId,omitnil,omitempty" name:"FunctionId"`

	// 规则描述。
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 函数名称。
	FunctionName *string `json:"FunctionName,omitnil,omitempty" name:"FunctionName"`

	// 函数触发规则优先级，数值越大，优先级越高。
	Priority *int64 `json:"Priority,omitnil,omitempty" name:"Priority"`

	// 创建时间。时间为世界标准时间（UTC）， 遵循 ISO 8601 标准的日期和时间格式。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 更新时间。时间为世界标准时间（UTC）， 遵循 ISO 8601 标准的日期和时间格式。
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`
}

type FunctionRuleCondition struct {
	// 边缘函数触发规则条件，该列表内所有项全部满足即判断该条件满足。
	RuleConditions []*RuleCondition `json:"RuleConditions,omitnil,omitempty" name:"RuleConditions"`
}

type Grpc struct {
	// 是否开启 Grpc 配置，取值有：
	// <li>on：开启；</li>
	// <li>off：关闭。</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`
}

// Predefined struct for user
type HandleFunctionRuntimeEnvironmentRequestParams struct {
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 函数 ID。
	FunctionId *string `json:"FunctionId,omitnil,omitempty" name:"FunctionId"`

	// 操作类型，取值有：
	// <li>setEnvironmentVariable：设置环境变量，当环境变量存在时为修改行为，否则为添加行为；</li>
	// <li>deleteEnvironmentVariable：删除环境变量变量；</li>
	// <li>clearEnvironmentVariable：清空环境变量变量；</li>
	// <li>resetEnvironmentVariable：重置环境变量变量。</li>
	Operation *string `json:"Operation,omitnil,omitempty" name:"Operation"`

	// 环境变量列表，函数运行环境最多支持 64 个变量。当 Operation 取值为 setEnvironmentVariable、deleteEnvironmentVariable、resetEnvironmentVariable 时必填。
	EnvironmentVariables []*FunctionEnvironmentVariable `json:"EnvironmentVariables,omitnil,omitempty" name:"EnvironmentVariables"`
}

type HandleFunctionRuntimeEnvironmentRequest struct {
	*tchttp.BaseRequest
	
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 函数 ID。
	FunctionId *string `json:"FunctionId,omitnil,omitempty" name:"FunctionId"`

	// 操作类型，取值有：
	// <li>setEnvironmentVariable：设置环境变量，当环境变量存在时为修改行为，否则为添加行为；</li>
	// <li>deleteEnvironmentVariable：删除环境变量变量；</li>
	// <li>clearEnvironmentVariable：清空环境变量变量；</li>
	// <li>resetEnvironmentVariable：重置环境变量变量。</li>
	Operation *string `json:"Operation,omitnil,omitempty" name:"Operation"`

	// 环境变量列表，函数运行环境最多支持 64 个变量。当 Operation 取值为 setEnvironmentVariable、deleteEnvironmentVariable、resetEnvironmentVariable 时必填。
	EnvironmentVariables []*FunctionEnvironmentVariable `json:"EnvironmentVariables,omitnil,omitempty" name:"EnvironmentVariables"`
}

func (r *HandleFunctionRuntimeEnvironmentRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *HandleFunctionRuntimeEnvironmentRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "FunctionId")
	delete(f, "Operation")
	delete(f, "EnvironmentVariables")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "HandleFunctionRuntimeEnvironmentRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type HandleFunctionRuntimeEnvironmentResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type HandleFunctionRuntimeEnvironmentResponse struct {
	*tchttp.BaseResponse
	Response *HandleFunctionRuntimeEnvironmentResponseParams `json:"Response"`
}

func (r *HandleFunctionRuntimeEnvironmentResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *HandleFunctionRuntimeEnvironmentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Header struct {
	// HTTP头部名称。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// HTTP头部值。
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type Hsts struct {
	// 是否开启，取值有：
	// <li>on：开启；</li>
	// <li>off：关闭。</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// MaxAge 数值。单位为秒，最大值为1天。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxAge *int64 `json:"MaxAge,omitnil,omitempty" name:"MaxAge"`

	// 是否包含子域名，取值有：
	// <li>on：开启；</li>
	// <li>off：关闭。</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	IncludeSubDomains *string `json:"IncludeSubDomains,omitnil,omitempty" name:"IncludeSubDomains"`

	// 是否开启预加载，取值有：
	// <li>on：开启；</li>
	// <li>off：关闭。</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Preload *string `json:"Preload,omitnil,omitempty" name:"Preload"`
}

type Https struct {
	// http2 配置开关，取值有：
	// <li>on：开启；</li>
	// <li>off：关闭。</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Http2 *string `json:"Http2,omitnil,omitempty" name:"Http2"`

	// OCSP 配置开关，取值有：
	// <li>on：开启；</li>
	// <li>off：关闭。</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	OcspStapling *string `json:"OcspStapling,omitnil,omitempty" name:"OcspStapling"`

	// Tls 版本设置，取值有：
	// <li>TLSv1：TLSv1版本；</li>
	// <li>TLSV1.1：TLSv1.1版本；</li>
	// <li>TLSV1.2：TLSv1.2版本；</li>
	// <li>TLSv1.3：TLSv1.3版本。</li>修改时必须开启连续的版本。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TlsVersion []*string `json:"TlsVersion,omitnil,omitempty" name:"TlsVersion"`

	// HSTS 配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Hsts *Hsts `json:"Hsts,omitnil,omitempty" name:"Hsts"`

	// 证书配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CertInfo []*ServerCertInfo `json:"CertInfo,omitnil,omitempty" name:"CertInfo"`

	// 申请类型，取值有：
	// <li>apply：托管EdgeOne；</li>
	// <li>none：不托管EdgeOne。</li>不填，默认取值为none。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplyType *string `json:"ApplyType,omitnil,omitempty" name:"ApplyType"`

	// 密码套件，取值有：
	// <li>loose-v2023：提供高兼容性，安全性一般，支持 TLS 1.0-1.3 密码套件；</li>
	// <li>general-v2023：提供较高兼容性，安全性中等，支持 TLS 1.2-1.3 密码套件；</li>
	// <li>strict-v2023：提供高安全性能，禁用所有含不安全隐患的加密套件，支持 TLS 1.2-1.3 密码套件。</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	CipherSuite *string `json:"CipherSuite,omitnil,omitempty" name:"CipherSuite"`
}

type IPGroup struct {
	// 组 Id，创建时填 0 即可。
	GroupId *int64 `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 组名称。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// IP 组内容，仅支持 IP 及 IP 掩码。
	Content []*string `json:"Content,omitnil,omitempty" name:"Content"`
}

type IPRegionInfo struct {
	// IP 地址，IPV4 或 IPV6。
	IP *string `json:"IP,omitnil,omitempty" name:"IP"`

	// IP 是否属于 EdgeOne 节点，取值有：
	// <li>yes：该 IP 属于 EdgeOne 节点；</li>
	// <li>no：该 IP 不属于 EdgeOne 节点。</li>
	IsEdgeOneIP *string `json:"IsEdgeOneIP,omitnil,omitempty" name:"IsEdgeOneIP"`
}

type IPWhitelist struct {
	// IPv4列表。
	IPv4 []*string `json:"IPv4,omitnil,omitempty" name:"IPv4"`

	// IPv6列表。
	IPv6 []*string `json:"IPv6,omitnil,omitempty" name:"IPv6"`
}

type Identification struct {
	// 站点名称。
	ZoneName *string `json:"ZoneName,omitnil,omitempty" name:"ZoneName"`

	// 验证子域名。验证站点时，该值为空。验证子域名是为具体子域名。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 验证状态，取值有：
	// <li> pending：验证中；</li>
	// <li> finished：验证完成。</li>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 站点归属权校验：Dns校验信息。
	Ascription *AscriptionInfo `json:"Ascription,omitnil,omitempty" name:"Ascription"`

	// 域名当前的 NS 记录。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OriginalNameServers []*string `json:"OriginalNameServers,omitnil,omitempty" name:"OriginalNameServers"`

	// 站点归属权校验：文件校验信息。
	FileAscription *FileAscriptionInfo `json:"FileAscription,omitnil,omitempty" name:"FileAscription"`
}

// Predefined struct for user
type IdentifyZoneRequestParams struct {
	// 站点名称。
	ZoneName *string `json:"ZoneName,omitnil,omitempty" name:"ZoneName"`

	// 站点下的子域名。如果验证站点下的子域名，则传该值，否则为空。
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`
}

type IdentifyZoneRequest struct {
	*tchttp.BaseRequest
	
	// 站点名称。
	ZoneName *string `json:"ZoneName,omitnil,omitempty" name:"ZoneName"`

	// 站点下的子域名。如果验证站点下的子域名，则传该值，否则为空。
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`
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
	delete(f, "Domain")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "IdentifyZoneRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type IdentifyZoneResponseParams struct {
	// 站点归属校验：Dns校验信息。
	Ascription *AscriptionInfo `json:"Ascription,omitnil,omitempty" name:"Ascription"`

	// 站点归属权校验：文件校验信息。
	FileAscription *FileAscriptionInfo `json:"FileAscription,omitnil,omitempty" name:"FileAscription"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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

type ImageOptimize struct {
	// 开关，取值有：
	// <li>on：开启；</li>
	// <li>off：关闭。</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`
}

// Predefined struct for user
type IncreasePlanQuotaRequestParams struct {
	// 套餐 ID, 形如 edgeone-2unuvzjmmn2q。
	PlanId *string `json:"PlanId,omitnil,omitempty" name:"PlanId"`

	// 新增的套餐配额类型，取值有：<li> site：站点数；</li><li> precise_access_control_rule：Web 防护 - 自定义规则 - 精准匹配策略的规则配额；</li><li> rate_limiting_rule：Web 防护 - 速率限制 - 精准速率限制模块的规则配额。</li>
	QuotaType *string `json:"QuotaType,omitnil,omitempty" name:"QuotaType"`

	// 新增的配额个数。单次新增的配额个数上限为 100。
	QuotaNumber *int64 `json:"QuotaNumber,omitnil,omitempty" name:"QuotaNumber"`
}

type IncreasePlanQuotaRequest struct {
	*tchttp.BaseRequest
	
	// 套餐 ID, 形如 edgeone-2unuvzjmmn2q。
	PlanId *string `json:"PlanId,omitnil,omitempty" name:"PlanId"`

	// 新增的套餐配额类型，取值有：<li> site：站点数；</li><li> precise_access_control_rule：Web 防护 - 自定义规则 - 精准匹配策略的规则配额；</li><li> rate_limiting_rule：Web 防护 - 速率限制 - 精准速率限制模块的规则配额。</li>
	QuotaType *string `json:"QuotaType,omitnil,omitempty" name:"QuotaType"`

	// 新增的配额个数。单次新增的配额个数上限为 100。
	QuotaNumber *int64 `json:"QuotaNumber,omitnil,omitempty" name:"QuotaNumber"`
}

func (r *IncreasePlanQuotaRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *IncreasePlanQuotaRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PlanId")
	delete(f, "QuotaType")
	delete(f, "QuotaNumber")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "IncreasePlanQuotaRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type IncreasePlanQuotaResponseParams struct {
	// 订单号。
	DealName *string `json:"DealName,omitnil,omitempty" name:"DealName"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type IncreasePlanQuotaResponse struct {
	*tchttp.BaseResponse
	Response *IncreasePlanQuotaResponseParams `json:"Response"`
}

func (r *IncreasePlanQuotaResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *IncreasePlanQuotaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type IntelligenceRule struct {
	// 开关，取值有：
	// <li>on：开启；</li>
	// <li>off：关闭。</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// 规则详情。
	// 注意：此字段可能返回 null，表示取不到有效值。
	IntelligenceRuleItems []*IntelligenceRuleItem `json:"IntelligenceRuleItems,omitnil,omitempty" name:"IntelligenceRuleItems"`
}

type IntelligenceRuleItem struct {
	// 智能分析标签，取值有：
	// <li>evil_bot：恶意bot；</li>
	// <li>suspect_bot：疑似bot；</li>
	// <li>good_bot：良好bot；</li>
	// <li>normal：正常请求。</li>
	Label *string `json:"Label,omitnil,omitempty" name:"Label"`

	// 触发智能分析标签对应的处置方式，取值有：
	// <li>drop：拦截；</li>
	// <li>trans：放行；</li>
	// <li>alg：Javascript挑战；</li>
	// <li>captcha：数字验证码；</li>
	// <li>monitor：观察。</li>
	Action *string `json:"Action,omitnil,omitempty" name:"Action"`
}

type IpTableConfig struct {
	// 开关，取值有：
	// <li>on：开启；</li>
	// <li>off：关闭；</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// 基础管控规则。如果为null，默认使用历史配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	IpTableRules []*IpTableRule `json:"IpTableRules,omitnil,omitempty" name:"IpTableRules"`
}

type IpTableRule struct {
	// 动作，取值有：
	// <li> drop：拦截；</li>
	// <li> trans：放行；</li>
	// <li> monitor：观察。</li>
	Action *string `json:"Action,omitnil,omitempty" name:"Action"`

	// 根据类型匹配，取值有：
	// <li>ip：客户端 IP 进行匹配；</li>
	// <li>area：客户端 IP 所属地区匹配；</li>
	// <li>asn：客户端所属的自治系统进行匹配；</li>
	// <li>referer：请求头 Referer 进行匹配；</li>
	// <li>ua：请求头 User-Agent 进行匹配；</li>
	// <li>url：请求 URL 进行匹配。</li>
	MatchFrom *string `json:"MatchFrom,omitnil,omitempty" name:"MatchFrom"`

	// 规则的匹配方式。取值有：
	// <li> match：匹配，适用于 MatchFrom 为 ip；</li>
	// <li> not_match：不匹配，适用于 MatchFrom 为 ip；</li>
	// <li> include_area：地域包含，适用于 MatchFrom 为 area；</li>
	// <li> not_include_area：地域不包含，适用于 MatchFrom 为 area；</li>
	// <li> asn_match：ASN 包含，适用于 MatchFrom 为 asn；</li>
	// <li> asn_not_match：ASN 不包含，适用于 MatchFrom 为 asn；</li>
	// <li> equal：等于，适用于 MatchFrom 为 ua , referer；</li>
	// <li> not_equal：不等于，适用于 MatchFrom 为 ua , referer；</li>
	// <li> include：通配符匹配，适用于 MatchFrom 为 ua , referer , url；</li>
	// <li> not_include：通配符不匹配，适用于 MatchFrom 为 ua , referer；</li>
	// <li> is_emty：配置内容为空，适用于 MatchFrom 为 ua , referer；</li>
	// <li> not_exists：配置内容不存在，适用于 MatchFrom 为 ua , referer。</li>
	Operator *string `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 规则id。仅出参使用。
	RuleID *int64 `json:"RuleID,omitnil,omitempty" name:"RuleID"`

	// 更新时间。仅出参使用。
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 规则启用状态。取值有：
	// <li> on：启用；</li>
	// <li> off：未启用。</li>
	// 当入参缺省时，按 on 取值。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 规则名。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleName *string `json:"RuleName,omitnil,omitempty" name:"RuleName"`

	// 匹配内容。支持多值输入。
	// <li>当输入多个匹配值时，请使用英文逗号分隔；</li>
	// <li>当 MatchFrom 为 ua 时，不支持多值输入；</li>
	// <li>当 Operator 为 is_empty 或 not_exists 时，本字段入参值无效。</li>
	MatchContent *string `json:"MatchContent,omitnil,omitempty" name:"MatchContent"`
}

type Ipv6 struct {
	// Ipv6 访问功能配置，取值有：
	// <li>on：开启Ipv6访问功能；</li>
	// <li>off：关闭Ipv6访问功能。</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`
}

type L4OfflineLog struct {
	// 四层代理实例 ID。
	ProxyId *string `json:"ProxyId,omitnil,omitempty" name:"ProxyId"`

	// 日志所属区域，取值有：
	// <li>mainland：中国大陆境内;</li>
	// <li>overseas：全球（不含中国大陆）。</li>
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 离线日志数据包名。
	LogPacketName *string `json:"LogPacketName,omitnil,omitempty" name:"LogPacketName"`

	// 离线日志下载地址。
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 日志打包时间，此参数已经废弃。
	LogTime *int64 `json:"LogTime,omitnil,omitempty" name:"LogTime"`

	// 日志打包开始时间。
	LogStartTime *string `json:"LogStartTime,omitnil,omitempty" name:"LogStartTime"`

	// 日志打包结束时间。
	LogEndTime *string `json:"LogEndTime,omitnil,omitempty" name:"LogEndTime"`

	// 日志大小，单位为 Byte。
	Size *int64 `json:"Size,omitnil,omitempty" name:"Size"`
}

type L4Proxy struct {
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 四层代理实例 ID。
	ProxyId *string `json:"ProxyId,omitnil,omitempty" name:"ProxyId"`

	// 四层代理实例名称。
	ProxyName *string `json:"ProxyName,omitnil,omitempty" name:"ProxyName"`

	// 四层代理实例的加速区域。 
	// <li>mainland：中国大陆可用区；</li>
	// <li>overseas： 全球可用区（不含中国大陆）；</li>
	//  <li>global：全球可用区。</li>	
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 接入 CNAME。
	Cname *string `json:"Cname,omitnil,omitempty" name:"Cname"`

	// 开启固定 IP 后，该值会返回对应的接入 IP；未开启时，该值为空。
	Ips []*string `json:"Ips,omitnil,omitempty" name:"Ips"`

	// 四层代理实例状态。
	// <li>online：已启用；</li>
	// <li>offline：已停用；</li>
	// <li>progress：部署中；</li>	
	// <li>stopping：停用中；</li>
	// <li>banned：已封禁；</li>
	// <li>fail：部署失败/停用失败。</li>	
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 是否开启 IPv6 访问。 
	// <li>on：开启；</li> 
	// <li>off：关闭。</li>
	Ipv6 *string `json:"Ipv6,omitnil,omitempty" name:"Ipv6"`

	// 是否开启固定 IP。
	//  <li>on：开启；</li> <li>off：关闭。</li>
	StaticIp *string `json:"StaticIp,omitnil,omitempty" name:"StaticIp"`

	// 是否开启中国大陆网络优化。
	//  <li>on：开启</li> <li>off：关闭</li>
	AccelerateMainland *string `json:"AccelerateMainland,omitnil,omitempty" name:"AccelerateMainland"`

	// 安全防护配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DDosProtectionConfig *DDosProtectionConfig `json:"DDosProtectionConfig,omitnil,omitempty" name:"DDosProtectionConfig"`

	// 四层代理实例下的转发规则数量。
	L4ProxyRuleCount *int64 `json:"L4ProxyRuleCount,omitnil,omitempty" name:"L4ProxyRuleCount"`

	// 最新变更时间。
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`
}

type L4ProxyRemoteAuth struct {
	// 四层远程鉴权开关，取值有：
	// <li>on：表示开启;</li>
	// <li>off：表示关闭。</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// 远程鉴权服务地址，格式为: domain/ip:port。例：example.auth.com:8888
	Address *string `json:"Address,omitnil,omitempty" name:"Address"`

	// 远程鉴权服务不可访问后，经过四层转发规则默认回源行为，取值有：
	// <li>reject：表示进行拦截，拒绝访问;</li>
	// <li>allow：表示允许通过。</li>
	ServerFaultyBehavior *string `json:"ServerFaultyBehavior,omitnil,omitempty" name:"ServerFaultyBehavior"`
}

type L4ProxyRule struct {
	// 转发规则 ID。
	// 注意：L4ProxyRule 在 CreateL4ProxyRules 作为入参使用时，该参数请勿填写；在 ModifyL4ProxyRules 作为入参使用时，该参数必填。
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// 转发协议。取值有：
	// <li>TCP：TCP 协议；</li>
	// <li>UDP：UDP 协议。</li>
	// 注意：L4ProxyRule 在 CreateL4ProxyRules 作为入参使用时，该参数必填；在 ModifyL4ProxyRules 作为入参使用时，该参数选填，不填写时表示不修改。
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// 转发端口，支持按照以下形式填写：
	// <li>单端口，如：80；</li>
	// <li>端口段，如：81-85。表示 81、82、83、84、85 五个端口。</li>
	// 注意：L4ProxyRule 在 CreateL4ProxyRules 作为入参使用时，该参数必填；在 ModifyL4ProxyRules 作为入参使用时，该参数选填，不填写时表示不修改。
	PortRange []*string `json:"PortRange,omitnil,omitempty" name:"PortRange"`

	// 源站类型，取值有：
	// <li>IP_DOMAIN：IP/域名源站；</li>
	// <li>ORIGIN_GROUP：源站组；</li>
	// <li>LB：负载均衡，当前仅白名单开放。</li>
	// 注意：L4ProxyRule 在 CreateL4ProxyRules 作为入参使用时，该参数必填；在 ModifyL4ProxyRules 作为入参使用时，该参数选填，不填写时表示不修改。
	OriginType *string `json:"OriginType,omitnil,omitempty" name:"OriginType"`

	// 源站地址：
	// <li>当 OriginType 为 IP_DOMAIN 时，填写 IP 或域名，如 8.8.8.8 或 test.com ；</li>
	// <li>当 OriginType 为 ORIGIN_GROUP 时，填写源站组 ID，如 og-537y24vf5b41；</li>
	// <li>当 OriginType 为 LB 时，填写负载均衡实例 ID，如 lb-2qwk30xf7s9g。</li>
	// 注意：L4ProxyRule 在 CreateL4ProxyRules 作为入参使用时，该参数必填；在 ModifyL4ProxyRules 作为入参使用时，该参数选填，不填写时表示不修改。
	OriginValue []*string `json:"OriginValue,omitnil,omitempty" name:"OriginValue"`

	// 源站端口，支持按照以下形式填写：
	// <li>单端口，如：80；</li>
	// <li>端口段，如：81-85，表示 81、82、83、84、85 五个端口。填写端口段时，则需要与转发端口段长度保持一致，例如转发端口：80-90，则转发端口：90-100。</li>
	// 注意：L4ProxyRule 在 CreateL4ProxyRules 作为入参使用时，该参数必填；在 ModifyL4ProxyRules 作为入参使用时，该参数选填，不填写时表示不修改。
	OriginPortRange *string `json:"OriginPortRange,omitnil,omitempty" name:"OriginPortRange"`

	// 传递客户端 IP 的形式，取值有：
	// <li>TOA：TOA（仅 Protocol = TCP 时可选）；</li> 
	// <li>PPV1：Proxy Protocol 传递，协议版本 V1（仅 Protocol = TCP 时可选）；</li>
	// <li>PPV2：Proxy Protocol 传递，协议版本 V2；</li> 
	// <li>SPP：Simple Proxy Protocol 传递，（仅 Protocol = UDP 时可选）；</li> 
	// <li>OFF：不传递。</li>
	// 注意：L4ProxyRule 在 CreateL4ProxyRules 作为入参使用时，该参数选填，不填写时默认为 OFF；在 ModifyL4ProxyRules 作为入参使用时，该参数选填，不填写表示不修改。
	ClientIPPassThroughMode *string `json:"ClientIPPassThroughMode,omitnil,omitempty" name:"ClientIPPassThroughMode"`

	// 是否开启会话保持，取值有：
	// <li>on：开启；</li>
	// <li>off：关闭。</li>
	// 注意：L4ProxyRule 在 CreateL4ProxyRules 作为入参使用时，该参数选填，不填写时默认为 off；在 ModifyL4ProxyRules 作为入参使用时，该参数选填，不填写表示不修改。
	SessionPersist *string `json:"SessionPersist,omitnil,omitempty" name:"SessionPersist"`

	// 会话保持时间，取值范围为 30-3600，单位为秒。
	// 注意：L4ProxyRule 在 CreateL4ProxyRules 作为入参使用时，该参数选填，仅当 SessionPersist = on 时，该值才会生效，且当 SessionPersist = on ，该值不填写默认为 3600；在 ModifyL4ProxyRules 作为入参使用时，该参数选填，不填写表示不修改。
	SessionPersistTime *uint64 `json:"SessionPersistTime,omitnil,omitempty" name:"SessionPersistTime"`

	// 规则标签。可输入1-50 个任意字符。
	// 注意：L4ProxyRule 在 CreateL4ProxyRules 作为入参使用时，该参数选填；在 ModifyL4ProxyRules 作为入参使用时，该参数选填，不填写表示不修改。
	RuleTag *string `json:"RuleTag,omitnil,omitempty" name:"RuleTag"`

	// 规则状态，取值有：
	// <li>online：已启用；</li>
	// <li>offline：已停用；</li>
	// <li>progress：部署中；</li>
	// <li>stopping：停用中；</li>
	// <li>fail：部署失败/停用失败。</li>
	// 注意：L4ProxyRule 在 CreateL4ProxyRules、ModifyL4ProxyRules 作为入参使用时，该参数请勿填写。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// BuID。
	BuId *string `json:"BuId,omitnil,omitempty" name:"BuId"`

	// 远程鉴权信息。
	// 注意：RemoteAuth 在 CreateL4ProxyRules 或 ModifyL4ProxyRules 不可作为入参使用，如有传此参数，会忽略。在 DescribeL4ProxyRules 返回为空时，表示没有开启远程鉴权。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RemoteAuth *L4ProxyRemoteAuth `json:"RemoteAuth,omitnil,omitempty" name:"RemoteAuth"`
}

type L7OfflineLog struct {
	// 离线日志域名。
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 日志所属区域，取值有：
	// <li>mainland：中国大陆境内; </li>
	// <li>overseas：全球（不含中国大陆）。</li>
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 离线日志数据包名。	
	LogPacketName *string `json:"LogPacketName,omitnil,omitempty" name:"LogPacketName"`

	// 离线日志下载地址。	
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 日志打包时间，此参数已经废弃。
	LogTime *int64 `json:"LogTime,omitnil,omitempty" name:"LogTime"`

	// 日志打包开始时间。
	LogStartTime *string `json:"LogStartTime,omitnil,omitempty" name:"LogStartTime"`

	// 日志打包结束时间。
	LogEndTime *string `json:"LogEndTime,omitnil,omitempty" name:"LogEndTime"`

	// 日志原始大小，单位 Byte。
	Size *int64 `json:"Size,omitnil,omitempty" name:"Size"`
}

type LogFormat struct {
	// 日志投递的预设输出格式类型，取值有：
	// <li>json：使用预设日志输出格式 JSON Lines，单条日志中的字段以键值对方式呈现；</li>
	// <li>csv：使用预设日志输出格式 csv，单条日志中仅呈现字段值，不呈现字段名称。</li>
	FormatType *string `json:"FormatType,omitnil,omitempty" name:"FormatType"`

	// 在每个日志投递批次之前添加的字符串。每个日志投递批次可能包含多条日志记录。
	BatchPrefix *string `json:"BatchPrefix,omitnil,omitempty" name:"BatchPrefix"`

	// 在每个日志投递批次后附加的字符串。
	BatchSuffix *string `json:"BatchSuffix,omitnil,omitempty" name:"BatchSuffix"`

	// 在每条日志记录之前添加的字符串。
	RecordPrefix *string `json:"RecordPrefix,omitnil,omitempty" name:"RecordPrefix"`

	// 在每条日志记录后附加的字符串。
	RecordSuffix *string `json:"RecordSuffix,omitnil,omitempty" name:"RecordSuffix"`

	// 插入日志记录之间作为分隔符的字符串，取值有：
	// <li>\n：换行符；</li>
	// <li>\t：制表符；</li>
	// <li>，：半角逗号。</li>
	RecordDelimiter *string `json:"RecordDelimiter,omitnil,omitempty" name:"RecordDelimiter"`

	// 单条日志记录内，插入字段之间作为分隔符的字符串，取值有：
	// <li>\t：制表符；</li>
	// <li>，：半角逗号；</li>
	// <li>;：半角分号。</li>
	FieldDelimiter *string `json:"FieldDelimiter,omitnil,omitempty" name:"FieldDelimiter"`
}

type MaxAge struct {
	// 是否遵循源站，取值有：
	// <li>on：遵循源站，忽略MaxAge 时间设置；</li>
	// <li>off：不遵循源站，使用MaxAge 时间设置。</li>
	FollowOrigin *string `json:"FollowOrigin,omitnil,omitempty" name:"FollowOrigin"`

	// MaxAge 时间设置，单位秒，最大365天。
	// 注意：时间为0，即不缓存。
	MaxAgeTime *int64 `json:"MaxAgeTime,omitnil,omitempty" name:"MaxAgeTime"`
}

// Predefined struct for user
type ModifyAccelerationDomainRequestParams struct {
	// 加速域名所属站点ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 加速域名名称。
	DomainName *string `json:"DomainName,omitnil,omitempty" name:"DomainName"`

	// 源站信息。
	OriginInfo *OriginInfo `json:"OriginInfo,omitnil,omitempty" name:"OriginInfo"`

	// 回源协议，取值有：
	// <li>FOLLOW: 协议跟随；</li>
	// <li>HTTP: HTTP协议回源；</li>
	// <li>HTTPS: HTTPS协议回源。</li>
	// <li>不填保持原有配置。</li>
	OriginProtocol *string `json:"OriginProtocol,omitnil,omitempty" name:"OriginProtocol"`

	// HTTP回源端口，取值为1-65535，当OriginProtocol=FOLLOW/HTTP时生效, 不填保持原有配置。
	HttpOriginPort *uint64 `json:"HttpOriginPort,omitnil,omitempty" name:"HttpOriginPort"`

	// HTTPS回源端口，取值为1-65535，当OriginProtocol=FOLLOW/HTTPS时生效，不填保持原有配置。
	HttpsOriginPort *uint64 `json:"HttpsOriginPort,omitnil,omitempty" name:"HttpsOriginPort"`

	// IPv6状态，取值有：
	// <li>follow：遵循站点IPv6配置；</li>
	// <li>on：开启状态；</li>
	// <li>off：关闭状态。</li>
	// <li>不填保持原有配置。</li>
	IPv6Status *string `json:"IPv6Status,omitnil,omitempty" name:"IPv6Status"`
}

type ModifyAccelerationDomainRequest struct {
	*tchttp.BaseRequest
	
	// 加速域名所属站点ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 加速域名名称。
	DomainName *string `json:"DomainName,omitnil,omitempty" name:"DomainName"`

	// 源站信息。
	OriginInfo *OriginInfo `json:"OriginInfo,omitnil,omitempty" name:"OriginInfo"`

	// 回源协议，取值有：
	// <li>FOLLOW: 协议跟随；</li>
	// <li>HTTP: HTTP协议回源；</li>
	// <li>HTTPS: HTTPS协议回源。</li>
	// <li>不填保持原有配置。</li>
	OriginProtocol *string `json:"OriginProtocol,omitnil,omitempty" name:"OriginProtocol"`

	// HTTP回源端口，取值为1-65535，当OriginProtocol=FOLLOW/HTTP时生效, 不填保持原有配置。
	HttpOriginPort *uint64 `json:"HttpOriginPort,omitnil,omitempty" name:"HttpOriginPort"`

	// HTTPS回源端口，取值为1-65535，当OriginProtocol=FOLLOW/HTTPS时生效，不填保持原有配置。
	HttpsOriginPort *uint64 `json:"HttpsOriginPort,omitnil,omitempty" name:"HttpsOriginPort"`

	// IPv6状态，取值有：
	// <li>follow：遵循站点IPv6配置；</li>
	// <li>on：开启状态；</li>
	// <li>off：关闭状态。</li>
	// <li>不填保持原有配置。</li>
	IPv6Status *string `json:"IPv6Status,omitnil,omitempty" name:"IPv6Status"`
}

func (r *ModifyAccelerationDomainRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAccelerationDomainRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "DomainName")
	delete(f, "OriginInfo")
	delete(f, "OriginProtocol")
	delete(f, "HttpOriginPort")
	delete(f, "HttpsOriginPort")
	delete(f, "IPv6Status")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAccelerationDomainRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAccelerationDomainResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyAccelerationDomainResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAccelerationDomainResponseParams `json:"Response"`
}

func (r *ModifyAccelerationDomainResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAccelerationDomainResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAccelerationDomainStatusesRequestParams struct {
	// 加速域名所属站点ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 要执行状态变更的加速域名列表。
	DomainNames []*string `json:"DomainNames,omitnil,omitempty" name:"DomainNames"`

	// 加速域名状态，取值有：
	// <li>online：启用；</li>
	// <li>offline：停用。</li>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 是否强制停用。当域名存在关联资源（如马甲域名、流量调度功能）时，是否强制停用该域名，取值有：
	// <li> true：停用该域名及所有关联资源；</li>
	// <li> false：当该加速域名存在关联资源时，不允许停用。</li>不填写，默认值为：false。
	Force *bool `json:"Force,omitnil,omitempty" name:"Force"`
}

type ModifyAccelerationDomainStatusesRequest struct {
	*tchttp.BaseRequest
	
	// 加速域名所属站点ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 要执行状态变更的加速域名列表。
	DomainNames []*string `json:"DomainNames,omitnil,omitempty" name:"DomainNames"`

	// 加速域名状态，取值有：
	// <li>online：启用；</li>
	// <li>offline：停用。</li>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 是否强制停用。当域名存在关联资源（如马甲域名、流量调度功能）时，是否强制停用该域名，取值有：
	// <li> true：停用该域名及所有关联资源；</li>
	// <li> false：当该加速域名存在关联资源时，不允许停用。</li>不填写，默认值为：false。
	Force *bool `json:"Force,omitnil,omitempty" name:"Force"`
}

func (r *ModifyAccelerationDomainStatusesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAccelerationDomainStatusesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "DomainNames")
	delete(f, "Status")
	delete(f, "Force")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAccelerationDomainStatusesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAccelerationDomainStatusesResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyAccelerationDomainStatusesResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAccelerationDomainStatusesResponseParams `json:"Response"`
}

func (r *ModifyAccelerationDomainStatusesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAccelerationDomainStatusesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAliasDomainRequestParams struct {
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 别称域名名称。
	AliasName *string `json:"AliasName,omitnil,omitempty" name:"AliasName"`

	// 目标域名名称。
	TargetName *string `json:"TargetName,omitnil,omitempty" name:"TargetName"`

	// 证书配置，取值有：
	// <li> none：不配置；</li>
	// <li> hosting：SSL托管证书；</li>
	// <li> apply：申请免费证书。</li>不填写保持原有配置。
	CertType *string `json:"CertType,omitnil,omitempty" name:"CertType"`

	// 当 CertType 取值为 hosting 时填入相应证书 ID。
	CertId []*string `json:"CertId,omitnil,omitempty" name:"CertId"`
}

type ModifyAliasDomainRequest struct {
	*tchttp.BaseRequest
	
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 别称域名名称。
	AliasName *string `json:"AliasName,omitnil,omitempty" name:"AliasName"`

	// 目标域名名称。
	TargetName *string `json:"TargetName,omitnil,omitempty" name:"TargetName"`

	// 证书配置，取值有：
	// <li> none：不配置；</li>
	// <li> hosting：SSL托管证书；</li>
	// <li> apply：申请免费证书。</li>不填写保持原有配置。
	CertType *string `json:"CertType,omitnil,omitempty" name:"CertType"`

	// 当 CertType 取值为 hosting 时填入相应证书 ID。
	CertId []*string `json:"CertId,omitnil,omitempty" name:"CertId"`
}

func (r *ModifyAliasDomainRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAliasDomainRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "AliasName")
	delete(f, "TargetName")
	delete(f, "CertType")
	delete(f, "CertId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAliasDomainRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAliasDomainResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyAliasDomainResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAliasDomainResponseParams `json:"Response"`
}

func (r *ModifyAliasDomainResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAliasDomainResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAliasDomainStatusRequestParams struct {
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 别称域名状态，取值有：
	// <li> false：开启别称域名；</li>
	// <li> true：关闭别称域名。</li>
	Paused *bool `json:"Paused,omitnil,omitempty" name:"Paused"`

	// 待修改状态的别称域名名称。如果为空，则不执行修改状态操作。
	AliasNames []*string `json:"AliasNames,omitnil,omitempty" name:"AliasNames"`
}

type ModifyAliasDomainStatusRequest struct {
	*tchttp.BaseRequest
	
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 别称域名状态，取值有：
	// <li> false：开启别称域名；</li>
	// <li> true：关闭别称域名。</li>
	Paused *bool `json:"Paused,omitnil,omitempty" name:"Paused"`

	// 待修改状态的别称域名名称。如果为空，则不执行修改状态操作。
	AliasNames []*string `json:"AliasNames,omitnil,omitempty" name:"AliasNames"`
}

func (r *ModifyAliasDomainStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAliasDomainStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "Paused")
	delete(f, "AliasNames")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAliasDomainStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAliasDomainStatusResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyAliasDomainStatusResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAliasDomainStatusResponseParams `json:"Response"`
}

func (r *ModifyAliasDomainStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAliasDomainStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyApplicationProxyRequestParams struct {
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 代理 ID。
	ProxyId *string `json:"ProxyId,omitnil,omitempty" name:"ProxyId"`

	// 当 ProxyType=hostname 时，表示域名或子域名；
	// 当 ProxyType=instance 时，表示代理名称。
	ProxyName *string `json:"ProxyName,omitnil,omitempty" name:"ProxyName"`

	// 会话保持时间，取值范围：30-3600，单位：秒。
	// 不填写保持原有配置。
	SessionPersistTime *uint64 `json:"SessionPersistTime,omitnil,omitempty" name:"SessionPersistTime"`

	// 四层代理模式，取值有：
	// <li>instance：表示实例模式。</li>不填写使用默认值instance。
	ProxyType *string `json:"ProxyType,omitnil,omitempty" name:"ProxyType"`

	// Ipv6 访问配置，不填写保持原有配置。
	Ipv6 *Ipv6 `json:"Ipv6,omitnil,omitempty" name:"Ipv6"`

	// 中国大陆加速优化配置。 不填写表示保持原有配置。
	AccelerateMainland *AccelerateMainland `json:"AccelerateMainland,omitnil,omitempty" name:"AccelerateMainland"`
}

type ModifyApplicationProxyRequest struct {
	*tchttp.BaseRequest
	
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 代理 ID。
	ProxyId *string `json:"ProxyId,omitnil,omitempty" name:"ProxyId"`

	// 当 ProxyType=hostname 时，表示域名或子域名；
	// 当 ProxyType=instance 时，表示代理名称。
	ProxyName *string `json:"ProxyName,omitnil,omitempty" name:"ProxyName"`

	// 会话保持时间，取值范围：30-3600，单位：秒。
	// 不填写保持原有配置。
	SessionPersistTime *uint64 `json:"SessionPersistTime,omitnil,omitempty" name:"SessionPersistTime"`

	// 四层代理模式，取值有：
	// <li>instance：表示实例模式。</li>不填写使用默认值instance。
	ProxyType *string `json:"ProxyType,omitnil,omitempty" name:"ProxyType"`

	// Ipv6 访问配置，不填写保持原有配置。
	Ipv6 *Ipv6 `json:"Ipv6,omitnil,omitempty" name:"Ipv6"`

	// 中国大陆加速优化配置。 不填写表示保持原有配置。
	AccelerateMainland *AccelerateMainland `json:"AccelerateMainland,omitnil,omitempty" name:"AccelerateMainland"`
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
	delete(f, "AccelerateMainland")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyApplicationProxyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyApplicationProxyResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 代理ID。
	ProxyId *string `json:"ProxyId,omitnil,omitempty" name:"ProxyId"`

	// 规则ID。
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// 源站类型，取值有：
	// <li>custom：手动添加；</li>
	// <li>origins：源站组。</li>
	OriginType *string `json:"OriginType,omitnil,omitempty" name:"OriginType"`

	// 端口，支持格式：
	// <li>80：80端口；</li>
	// <li>81-90：81至90端口。</li>
	Port []*string `json:"Port,omitnil,omitempty" name:"Port"`

	// 协议，取值有：
	// <li>TCP：TCP协议；</li>
	// <li>UDP：UDP协议。</li>不填保持原有值。
	Proto *string `json:"Proto,omitnil,omitempty" name:"Proto"`

	// 源站信息：
	// <li>当 OriginType 为 custom 时，表示一个或多个源站，如`["8.8.8.8","9.9.9.9"]` 或 `OriginValue=["test.com"]`；</li>
	// <li>当 OriginType 为 origins 时，要求有且仅有一个元素，表示源站组ID，如`["origin-537f5b41-162a-11ed-abaa-525400c5da15"]`。</li>
	// 
	// 不填保持原有值。
	OriginValue []*string `json:"OriginValue,omitnil,omitempty" name:"OriginValue"`

	// 传递客户端IP，取值有：
	// <li>TOA：TOA（仅Proto=TCP时可选）；</li>
	// <li>PPV1：Proxy Protocol传递，协议版本V1（仅Proto=TCP时可选）；</li>
	// <li>PPV2：Proxy Protocol传递，协议版本V2；</li>
	// <li>OFF：不传递。</li>不填保持原有值。
	ForwardClientIp *string `json:"ForwardClientIp,omitnil,omitempty" name:"ForwardClientIp"`

	// 是否开启会话保持，取值有：
	// <li>true：开启；</li>
	// <li>false：关闭。</li>不填为false。
	SessionPersist *bool `json:"SessionPersist,omitnil,omitempty" name:"SessionPersist"`

	// 会话保持的时间，只有当SessionPersist为true时，该值才会生效。
	SessionPersistTime *uint64 `json:"SessionPersistTime,omitnil,omitempty" name:"SessionPersistTime"`

	// 源站端口，支持格式：
	// <li>单端口：80；</li>
	// <li>端口段：81-90，81至90端口。</li>
	OriginPort *string `json:"OriginPort,omitnil,omitempty" name:"OriginPort"`

	// 规则标签。不填保持原有值。
	RuleTag *string `json:"RuleTag,omitnil,omitempty" name:"RuleTag"`
}

type ModifyApplicationProxyRuleRequest struct {
	*tchttp.BaseRequest
	
	// 站点ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 代理ID。
	ProxyId *string `json:"ProxyId,omitnil,omitempty" name:"ProxyId"`

	// 规则ID。
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// 源站类型，取值有：
	// <li>custom：手动添加；</li>
	// <li>origins：源站组。</li>
	OriginType *string `json:"OriginType,omitnil,omitempty" name:"OriginType"`

	// 端口，支持格式：
	// <li>80：80端口；</li>
	// <li>81-90：81至90端口。</li>
	Port []*string `json:"Port,omitnil,omitempty" name:"Port"`

	// 协议，取值有：
	// <li>TCP：TCP协议；</li>
	// <li>UDP：UDP协议。</li>不填保持原有值。
	Proto *string `json:"Proto,omitnil,omitempty" name:"Proto"`

	// 源站信息：
	// <li>当 OriginType 为 custom 时，表示一个或多个源站，如`["8.8.8.8","9.9.9.9"]` 或 `OriginValue=["test.com"]`；</li>
	// <li>当 OriginType 为 origins 时，要求有且仅有一个元素，表示源站组ID，如`["origin-537f5b41-162a-11ed-abaa-525400c5da15"]`。</li>
	// 
	// 不填保持原有值。
	OriginValue []*string `json:"OriginValue,omitnil,omitempty" name:"OriginValue"`

	// 传递客户端IP，取值有：
	// <li>TOA：TOA（仅Proto=TCP时可选）；</li>
	// <li>PPV1：Proxy Protocol传递，协议版本V1（仅Proto=TCP时可选）；</li>
	// <li>PPV2：Proxy Protocol传递，协议版本V2；</li>
	// <li>OFF：不传递。</li>不填保持原有值。
	ForwardClientIp *string `json:"ForwardClientIp,omitnil,omitempty" name:"ForwardClientIp"`

	// 是否开启会话保持，取值有：
	// <li>true：开启；</li>
	// <li>false：关闭。</li>不填为false。
	SessionPersist *bool `json:"SessionPersist,omitnil,omitempty" name:"SessionPersist"`

	// 会话保持的时间，只有当SessionPersist为true时，该值才会生效。
	SessionPersistTime *uint64 `json:"SessionPersistTime,omitnil,omitempty" name:"SessionPersistTime"`

	// 源站端口，支持格式：
	// <li>单端口：80；</li>
	// <li>端口段：81-90，81至90端口。</li>
	OriginPort *string `json:"OriginPort,omitnil,omitempty" name:"OriginPort"`

	// 规则标签。不填保持原有值。
	RuleTag *string `json:"RuleTag,omitnil,omitempty" name:"RuleTag"`
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
	delete(f, "SessionPersistTime")
	delete(f, "OriginPort")
	delete(f, "RuleTag")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyApplicationProxyRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyApplicationProxyRuleResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 代理ID。
	ProxyId *string `json:"ProxyId,omitnil,omitempty" name:"ProxyId"`

	// 规则ID。
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// 状态，取值有：
	// <li>offline: 停用；</li>
	// <li>online: 启用。</li>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`
}

type ModifyApplicationProxyRuleStatusRequest struct {
	*tchttp.BaseRequest
	
	// 站点ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 代理ID。
	ProxyId *string `json:"ProxyId,omitnil,omitempty" name:"ProxyId"`

	// 规则ID。
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// 状态，取值有：
	// <li>offline: 停用；</li>
	// <li>online: 启用。</li>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 代理ID。
	ProxyId *string `json:"ProxyId,omitnil,omitempty" name:"ProxyId"`

	// 状态，取值有：
	// <li>offline: 停用；</li>
	// <li>online: 启用。</li>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`
}

type ModifyApplicationProxyStatusRequest struct {
	*tchttp.BaseRequest
	
	// 站点ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 代理ID。
	ProxyId *string `json:"ProxyId,omitnil,omitempty" name:"ProxyId"`

	// 状态，取值有：
	// <li>offline: 停用；</li>
	// <li>online: 启用。</li>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type ModifyCustomErrorPageRequestParams struct {
	// 自定义错误页面 ID。
	PageId *string `json:"PageId,omitnil,omitempty" name:"PageId"`

	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 自定义错误页名称，名称为2 - 60个字符。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 自定义错误页描述，描述内容不超过60个字符。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 自定义错误页面类型，取值有：<li>text/html。 </li><li>application/json。</li><li>plain/text。</li><li>text/xml。</li>
	ContentType *string `json:"ContentType,omitnil,omitempty" name:"ContentType"`

	// 自定义错误页面内容。内容不超过 2KB。
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`
}

type ModifyCustomErrorPageRequest struct {
	*tchttp.BaseRequest
	
	// 自定义错误页面 ID。
	PageId *string `json:"PageId,omitnil,omitempty" name:"PageId"`

	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 自定义错误页名称，名称为2 - 60个字符。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 自定义错误页描述，描述内容不超过60个字符。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 自定义错误页面类型，取值有：<li>text/html。 </li><li>application/json。</li><li>plain/text。</li><li>text/xml。</li>
	ContentType *string `json:"ContentType,omitnil,omitempty" name:"ContentType"`

	// 自定义错误页面内容。内容不超过 2KB。
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`
}

func (r *ModifyCustomErrorPageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCustomErrorPageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PageId")
	delete(f, "ZoneId")
	delete(f, "Name")
	delete(f, "Description")
	delete(f, "ContentType")
	delete(f, "Content")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyCustomErrorPageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCustomErrorPageResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyCustomErrorPageResponse struct {
	*tchttp.BaseResponse
	Response *ModifyCustomErrorPageResponseParams `json:"Response"`
}

func (r *ModifyCustomErrorPageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCustomErrorPageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyFunctionRequestParams struct {
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 函数 ID。
	FunctionId *string `json:"FunctionId,omitnil,omitempty" name:"FunctionId"`

	// 函数描述，最大支持 60 个字符，不填写保持原有配置。
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 函数内容，当前仅支持 JavaScript 代码，最大支持 5MB 大小，不填写保持原有配置。
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`
}

type ModifyFunctionRequest struct {
	*tchttp.BaseRequest
	
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 函数 ID。
	FunctionId *string `json:"FunctionId,omitnil,omitempty" name:"FunctionId"`

	// 函数描述，最大支持 60 个字符，不填写保持原有配置。
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 函数内容，当前仅支持 JavaScript 代码，最大支持 5MB 大小，不填写保持原有配置。
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`
}

func (r *ModifyFunctionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyFunctionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "FunctionId")
	delete(f, "Remark")
	delete(f, "Content")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyFunctionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyFunctionResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyFunctionResponse struct {
	*tchttp.BaseResponse
	Response *ModifyFunctionResponseParams `json:"Response"`
}

func (r *ModifyFunctionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyFunctionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyFunctionRulePriorityRequestParams struct {
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 规则 ID 列表，必须填入调整优先级后的所有规则 ID，多条规则执行顺序依次从上往下，不填写保持原优先级顺序。
	RuleIds []*string `json:"RuleIds,omitnil,omitempty" name:"RuleIds"`
}

type ModifyFunctionRulePriorityRequest struct {
	*tchttp.BaseRequest
	
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 规则 ID 列表，必须填入调整优先级后的所有规则 ID，多条规则执行顺序依次从上往下，不填写保持原优先级顺序。
	RuleIds []*string `json:"RuleIds,omitnil,omitempty" name:"RuleIds"`
}

func (r *ModifyFunctionRulePriorityRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyFunctionRulePriorityRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "RuleIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyFunctionRulePriorityRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyFunctionRulePriorityResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyFunctionRulePriorityResponse struct {
	*tchttp.BaseResponse
	Response *ModifyFunctionRulePriorityResponseParams `json:"Response"`
}

func (r *ModifyFunctionRulePriorityResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyFunctionRulePriorityResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyFunctionRuleRequestParams struct {
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 规则 ID。
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// 规则条件列表，相同触发规则的不同条件匹配项之间为或关系，不填写保持原有配置。
	FunctionRuleConditions []*FunctionRuleCondition `json:"FunctionRuleConditions,omitnil,omitempty" name:"FunctionRuleConditions"`

	// 函数 ID，命中触发规则条件后执行的函数，不填写保持原有配置。
	FunctionId *string `json:"FunctionId,omitnil,omitempty" name:"FunctionId"`

	// 规则描述，最大支持 60 个字符，不填写保持原有配置。
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`
}

type ModifyFunctionRuleRequest struct {
	*tchttp.BaseRequest
	
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 规则 ID。
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// 规则条件列表，相同触发规则的不同条件匹配项之间为或关系，不填写保持原有配置。
	FunctionRuleConditions []*FunctionRuleCondition `json:"FunctionRuleConditions,omitnil,omitempty" name:"FunctionRuleConditions"`

	// 函数 ID，命中触发规则条件后执行的函数，不填写保持原有配置。
	FunctionId *string `json:"FunctionId,omitnil,omitempty" name:"FunctionId"`

	// 规则描述，最大支持 60 个字符，不填写保持原有配置。
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`
}

func (r *ModifyFunctionRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyFunctionRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "RuleId")
	delete(f, "FunctionRuleConditions")
	delete(f, "FunctionId")
	delete(f, "Remark")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyFunctionRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyFunctionRuleResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyFunctionRuleResponse struct {
	*tchttp.BaseResponse
	Response *ModifyFunctionRuleResponseParams `json:"Response"`
}

func (r *ModifyFunctionRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyFunctionRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyHostsCertificateRequestParams struct {
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 需要修改证书配置的加速域名。
	Hosts []*string `json:"Hosts,omitnil,omitempty" name:"Hosts"`

	// 配置服务端证书的模式，取值有：
	// <li>disable：不配置服务端证书；</li>
	// <li>eofreecert：配置 EdgeOne 免费服务端证书；</li>
	// <li>sslcert：配置 SSL 托管服务端证书；</li>
	// 不填写表示服务端证书保持原有配置。
	Mode *string `json:"Mode,omitnil,omitempty" name:"Mode"`

	// SSL 证书配置，本参数仅在 mode 为 sslcert 时生效，传入对应证书的 CertId 即可。您可以前往 [SSL 证书列表](https://console.cloud.tencent.com/certoverview) 查看 CertId。
	ServerCertInfo []*ServerCertInfo `json:"ServerCertInfo,omitnil,omitempty" name:"ServerCertInfo"`

	// 托管类型，取值有：
	// <li>none：不托管EO；</li>
	// <li>apply：托管EO</li>
	// 不填，默认取值为none。
	//
	// Deprecated: ApplyType is deprecated.
	ApplyType *string `json:"ApplyType,omitnil,omitempty" name:"ApplyType"`

	// 边缘双向认证配置。
	// 不填写表示边缘双向认证保持原有配置。
	ClientCertInfo *MutualTLS `json:"ClientCertInfo,omitnil,omitempty" name:"ClientCertInfo"`
}

type ModifyHostsCertificateRequest struct {
	*tchttp.BaseRequest
	
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 需要修改证书配置的加速域名。
	Hosts []*string `json:"Hosts,omitnil,omitempty" name:"Hosts"`

	// 配置服务端证书的模式，取值有：
	// <li>disable：不配置服务端证书；</li>
	// <li>eofreecert：配置 EdgeOne 免费服务端证书；</li>
	// <li>sslcert：配置 SSL 托管服务端证书；</li>
	// 不填写表示服务端证书保持原有配置。
	Mode *string `json:"Mode,omitnil,omitempty" name:"Mode"`

	// SSL 证书配置，本参数仅在 mode 为 sslcert 时生效，传入对应证书的 CertId 即可。您可以前往 [SSL 证书列表](https://console.cloud.tencent.com/certoverview) 查看 CertId。
	ServerCertInfo []*ServerCertInfo `json:"ServerCertInfo,omitnil,omitempty" name:"ServerCertInfo"`

	// 托管类型，取值有：
	// <li>none：不托管EO；</li>
	// <li>apply：托管EO</li>
	// 不填，默认取值为none。
	ApplyType *string `json:"ApplyType,omitnil,omitempty" name:"ApplyType"`

	// 边缘双向认证配置。
	// 不填写表示边缘双向认证保持原有配置。
	ClientCertInfo *MutualTLS `json:"ClientCertInfo,omitnil,omitempty" name:"ClientCertInfo"`
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
	delete(f, "Mode")
	delete(f, "ServerCertInfo")
	delete(f, "ApplyType")
	delete(f, "ClientCertInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyHostsCertificateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyHostsCertificateResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type ModifyL4ProxyRequestParams struct {
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 应用 ID。
	ProxyId *string `json:"ProxyId,omitnil,omitempty" name:"ProxyId"`

	// 是否开启 IPv6 访问。 不填该参数时，表示不修改该配置。该配置仅在部分加速区域和安全防护配置下支持开启，详情请参考 [新建四层代理实例](https://cloud.tencent.com/document/product/1552/90025) 。取值为：
	// <li>on：开启；</li> 
	// <li>off：关闭。</li>
	Ipv6 *string `json:"Ipv6,omitnil,omitempty" name:"Ipv6"`

	// 是否开启中国大陆网络优化。不填该参数时，表示不修改该配置。该配置仅在部分加速区域和安全防护配置下支持开启，详情请参考 [新建四层代理实例](https://cloud.tencent.com/document/product/1552/90025) 。取值为：
	// <li>on：开启；</li> 
	// <li>off：关闭。</li>
	AccelerateMainland *string `json:"AccelerateMainland,omitnil,omitempty" name:"AccelerateMainland"`
}

type ModifyL4ProxyRequest struct {
	*tchttp.BaseRequest
	
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 应用 ID。
	ProxyId *string `json:"ProxyId,omitnil,omitempty" name:"ProxyId"`

	// 是否开启 IPv6 访问。 不填该参数时，表示不修改该配置。该配置仅在部分加速区域和安全防护配置下支持开启，详情请参考 [新建四层代理实例](https://cloud.tencent.com/document/product/1552/90025) 。取值为：
	// <li>on：开启；</li> 
	// <li>off：关闭。</li>
	Ipv6 *string `json:"Ipv6,omitnil,omitempty" name:"Ipv6"`

	// 是否开启中国大陆网络优化。不填该参数时，表示不修改该配置。该配置仅在部分加速区域和安全防护配置下支持开启，详情请参考 [新建四层代理实例](https://cloud.tencent.com/document/product/1552/90025) 。取值为：
	// <li>on：开启；</li> 
	// <li>off：关闭。</li>
	AccelerateMainland *string `json:"AccelerateMainland,omitnil,omitempty" name:"AccelerateMainland"`
}

func (r *ModifyL4ProxyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyL4ProxyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "ProxyId")
	delete(f, "Ipv6")
	delete(f, "AccelerateMainland")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyL4ProxyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyL4ProxyResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyL4ProxyResponse struct {
	*tchttp.BaseResponse
	Response *ModifyL4ProxyResponseParams `json:"Response"`
}

func (r *ModifyL4ProxyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyL4ProxyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyL4ProxyRulesRequestParams struct {
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 四层代理实例 ID。
	ProxyId *string `json:"ProxyId,omitnil,omitempty" name:"ProxyId"`

	// 转发规则列表。单次最多支持 200 条转发规则。
	// 注意：L4ProxyRule 在此处使用时，RuleId 为必填字段；Protocol、PortRange、OriginType、OriginValue、OriginPortRange、ClientIPPassThroughMode、SessionPersist、SessionPersistTime、RuleTag 均为选填字段，不填写表示不修改；Status 请勿填写。
	L4ProxyRules []*L4ProxyRule `json:"L4ProxyRules,omitnil,omitempty" name:"L4ProxyRules"`
}

type ModifyL4ProxyRulesRequest struct {
	*tchttp.BaseRequest
	
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 四层代理实例 ID。
	ProxyId *string `json:"ProxyId,omitnil,omitempty" name:"ProxyId"`

	// 转发规则列表。单次最多支持 200 条转发规则。
	// 注意：L4ProxyRule 在此处使用时，RuleId 为必填字段；Protocol、PortRange、OriginType、OriginValue、OriginPortRange、ClientIPPassThroughMode、SessionPersist、SessionPersistTime、RuleTag 均为选填字段，不填写表示不修改；Status 请勿填写。
	L4ProxyRules []*L4ProxyRule `json:"L4ProxyRules,omitnil,omitempty" name:"L4ProxyRules"`
}

func (r *ModifyL4ProxyRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyL4ProxyRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "ProxyId")
	delete(f, "L4ProxyRules")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyL4ProxyRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyL4ProxyRulesResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyL4ProxyRulesResponse struct {
	*tchttp.BaseResponse
	Response *ModifyL4ProxyRulesResponseParams `json:"Response"`
}

func (r *ModifyL4ProxyRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyL4ProxyRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyL4ProxyRulesStatusRequestParams struct {
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 四层代理实例 ID。
	ProxyId *string `json:"ProxyId,omitnil,omitempty" name:"ProxyId"`

	// 转发规则 ID 列表。单次最多支持 200 条转发规则。
	RuleIds []*string `json:"RuleIds,omitnil,omitempty" name:"RuleIds"`

	// 转发规则状态，取值有：
	// <li>online：启用；</li>
	// <li>offline：停用。</li>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`
}

type ModifyL4ProxyRulesStatusRequest struct {
	*tchttp.BaseRequest
	
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 四层代理实例 ID。
	ProxyId *string `json:"ProxyId,omitnil,omitempty" name:"ProxyId"`

	// 转发规则 ID 列表。单次最多支持 200 条转发规则。
	RuleIds []*string `json:"RuleIds,omitnil,omitempty" name:"RuleIds"`

	// 转发规则状态，取值有：
	// <li>online：启用；</li>
	// <li>offline：停用。</li>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`
}

func (r *ModifyL4ProxyRulesStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyL4ProxyRulesStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "ProxyId")
	delete(f, "RuleIds")
	delete(f, "Status")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyL4ProxyRulesStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyL4ProxyRulesStatusResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyL4ProxyRulesStatusResponse struct {
	*tchttp.BaseResponse
	Response *ModifyL4ProxyRulesStatusResponseParams `json:"Response"`
}

func (r *ModifyL4ProxyRulesStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyL4ProxyRulesStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyL4ProxyStatusRequestParams struct {
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 四层代理实例 ID。
	ProxyId *string `json:"ProxyId,omitnil,omitempty" name:"ProxyId"`

	// 四层代理实例状态，取值有：
	// <li>online：启用；</li>
	// <li>offline：停用。</li>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`
}

type ModifyL4ProxyStatusRequest struct {
	*tchttp.BaseRequest
	
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 四层代理实例 ID。
	ProxyId *string `json:"ProxyId,omitnil,omitempty" name:"ProxyId"`

	// 四层代理实例状态，取值有：
	// <li>online：启用；</li>
	// <li>offline：停用。</li>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`
}

func (r *ModifyL4ProxyStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyL4ProxyStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "ProxyId")
	delete(f, "Status")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyL4ProxyStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyL4ProxyStatusResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyL4ProxyStatusResponse struct {
	*tchttp.BaseResponse
	Response *ModifyL4ProxyStatusResponseParams `json:"Response"`
}

func (r *ModifyL4ProxyStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyL4ProxyStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyOriginGroupRequestParams struct {
	// 站点 ID
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 源站组 ID，此参数必填。
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 源站组名称，不填保持原有配置，可输入1 - 200个字符，允许的字符为 a - z, A - Z, 0 - 9, _, - 。	
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 源站组类型，取值有：
	// <li>GENERAL：通用型源站组，仅支持添加 IP/域名 源站，可以被域名服务、规则引擎、四层代理、通用型负载均衡引用；</li>
	// <li>HTTP： HTTP专用型源站组，支持添加 IP/域名、对象存储源站，无法被四层代理引用。</li>不填保持原有配置。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 源站记录信息，不填保持原有配置。
	Records []*OriginRecord `json:"Records,omitnil,omitempty" name:"Records"`

	// 回源 Host Header，仅 Type = HTTP 时生效， 不填或者填空表示不配置回源Host，规则引擎修改 Host Header 配置优先级高于源站组的 Host Header。
	HostHeader *string `json:"HostHeader,omitnil,omitempty" name:"HostHeader"`
}

type ModifyOriginGroupRequest struct {
	*tchttp.BaseRequest
	
	// 站点 ID
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 源站组 ID，此参数必填。
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 源站组名称，不填保持原有配置，可输入1 - 200个字符，允许的字符为 a - z, A - Z, 0 - 9, _, - 。	
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 源站组类型，取值有：
	// <li>GENERAL：通用型源站组，仅支持添加 IP/域名 源站，可以被域名服务、规则引擎、四层代理、通用型负载均衡引用；</li>
	// <li>HTTP： HTTP专用型源站组，支持添加 IP/域名、对象存储源站，无法被四层代理引用。</li>不填保持原有配置。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 源站记录信息，不填保持原有配置。
	Records []*OriginRecord `json:"Records,omitnil,omitempty" name:"Records"`

	// 回源 Host Header，仅 Type = HTTP 时生效， 不填或者填空表示不配置回源Host，规则引擎修改 Host Header 配置优先级高于源站组的 Host Header。
	HostHeader *string `json:"HostHeader,omitnil,omitempty" name:"HostHeader"`
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
	delete(f, "GroupId")
	delete(f, "Name")
	delete(f, "Type")
	delete(f, "Records")
	delete(f, "HostHeader")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyOriginGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyOriginGroupResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type ModifyPlanRequestParams struct {
	// 套餐 ID，形如 edgeone-2unuvzjmmn2q。
	PlanId *string `json:"PlanId,omitnil,omitempty" name:"PlanId"`

	// 预付费套餐自动续费配置。若开启了自动续费，则会在套餐到期前一天自动续费，仅支持个人版，基础版，标准版套餐。不填写表示保持原有配置。
	RenewFlag *RenewFlag `json:"RenewFlag,omitnil,omitempty" name:"RenewFlag"`
}

type ModifyPlanRequest struct {
	*tchttp.BaseRequest
	
	// 套餐 ID，形如 edgeone-2unuvzjmmn2q。
	PlanId *string `json:"PlanId,omitnil,omitempty" name:"PlanId"`

	// 预付费套餐自动续费配置。若开启了自动续费，则会在套餐到期前一天自动续费，仅支持个人版，基础版，标准版套餐。不填写表示保持原有配置。
	RenewFlag *RenewFlag `json:"RenewFlag,omitnil,omitempty" name:"RenewFlag"`
}

func (r *ModifyPlanRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPlanRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PlanId")
	delete(f, "RenewFlag")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyPlanRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyPlanResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyPlanResponse struct {
	*tchttp.BaseResponse
	Response *ModifyPlanResponseParams `json:"Response"`
}

func (r *ModifyPlanResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPlanResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyRealtimeLogDeliveryTaskRequestParams struct {
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 实时日志投递任务 ID。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 实时日志投递任务的名称，格式为数字、英文、-和_组合，最多 200 个字符。不填保持原有配置。
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`

	// 实时日志投递任务的状态，取值有：
	// <li>enabled: 启用；</li>
	// <li>disabled: 停用。</li>不填保持原有配置。
	DeliveryStatus *string `json:"DeliveryStatus,omitnil,omitempty" name:"DeliveryStatus"`

	// 实时日志投递任务对应的实体（七层域名或者四层代理实例）列表。取值示例如下：
	// <li>七层域名：domain.example.com；</li>
	// <li>四层代理实例：sid-2s69eb5wcms7。</li>不填保持原有配置。
	EntityList []*string `json:"EntityList,omitnil,omitempty" name:"EntityList"`

	// 投递的预设字段列表。不填保持原有配置。
	Fields []*string `json:"Fields,omitnil,omitempty" name:"Fields"`

	// 投递的自定义字段列表，支持在 HTTP 请求头、响应头、Cookie 中提取指定字段值。自定义字段名称不能重复，且最多不能超过 200 个字段。不填保持原有配置。
	CustomFields []*CustomField `json:"CustomFields,omitnil,omitempty" name:"CustomFields"`

	// 日志投递的过滤条件。不填表示投递全量日志。
	DeliveryConditions []*DeliveryCondition `json:"DeliveryConditions,omitnil,omitempty" name:"DeliveryConditions"`

	// 采样比例，采用千分制，取值范围为1-1000，例如：填写 605 表示采样比例为 60.5%。不填保持原有配置。
	Sample *uint64 `json:"Sample,omitnil,omitempty" name:"Sample"`

	// 日志投递的输出格式。不填保持原有配置。
	// 特别地，当 TaskType 取值为 cls 时，LogFormat.FormatType 的值只能为 json，且 LogFormat 中其他参数将被忽略，建议不传 LogFormat。
	LogFormat *LogFormat `json:"LogFormat,omitnil,omitempty" name:"LogFormat"`

	// 自定义 HTTP 服务的配置信息，不填保持原有配置。 
	CustomEndpoint *CustomEndpoint `json:"CustomEndpoint,omitnil,omitempty" name:"CustomEndpoint"`

	// AWS S3 兼容存储桶的配置信息，不填保持原有配置。
	S3 *S3 `json:"S3,omitnil,omitempty" name:"S3"`
}

type ModifyRealtimeLogDeliveryTaskRequest struct {
	*tchttp.BaseRequest
	
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 实时日志投递任务 ID。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 实时日志投递任务的名称，格式为数字、英文、-和_组合，最多 200 个字符。不填保持原有配置。
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`

	// 实时日志投递任务的状态，取值有：
	// <li>enabled: 启用；</li>
	// <li>disabled: 停用。</li>不填保持原有配置。
	DeliveryStatus *string `json:"DeliveryStatus,omitnil,omitempty" name:"DeliveryStatus"`

	// 实时日志投递任务对应的实体（七层域名或者四层代理实例）列表。取值示例如下：
	// <li>七层域名：domain.example.com；</li>
	// <li>四层代理实例：sid-2s69eb5wcms7。</li>不填保持原有配置。
	EntityList []*string `json:"EntityList,omitnil,omitempty" name:"EntityList"`

	// 投递的预设字段列表。不填保持原有配置。
	Fields []*string `json:"Fields,omitnil,omitempty" name:"Fields"`

	// 投递的自定义字段列表，支持在 HTTP 请求头、响应头、Cookie 中提取指定字段值。自定义字段名称不能重复，且最多不能超过 200 个字段。不填保持原有配置。
	CustomFields []*CustomField `json:"CustomFields,omitnil,omitempty" name:"CustomFields"`

	// 日志投递的过滤条件。不填表示投递全量日志。
	DeliveryConditions []*DeliveryCondition `json:"DeliveryConditions,omitnil,omitempty" name:"DeliveryConditions"`

	// 采样比例，采用千分制，取值范围为1-1000，例如：填写 605 表示采样比例为 60.5%。不填保持原有配置。
	Sample *uint64 `json:"Sample,omitnil,omitempty" name:"Sample"`

	// 日志投递的输出格式。不填保持原有配置。
	// 特别地，当 TaskType 取值为 cls 时，LogFormat.FormatType 的值只能为 json，且 LogFormat 中其他参数将被忽略，建议不传 LogFormat。
	LogFormat *LogFormat `json:"LogFormat,omitnil,omitempty" name:"LogFormat"`

	// 自定义 HTTP 服务的配置信息，不填保持原有配置。 
	CustomEndpoint *CustomEndpoint `json:"CustomEndpoint,omitnil,omitempty" name:"CustomEndpoint"`

	// AWS S3 兼容存储桶的配置信息，不填保持原有配置。
	S3 *S3 `json:"S3,omitnil,omitempty" name:"S3"`
}

func (r *ModifyRealtimeLogDeliveryTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRealtimeLogDeliveryTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "TaskId")
	delete(f, "TaskName")
	delete(f, "DeliveryStatus")
	delete(f, "EntityList")
	delete(f, "Fields")
	delete(f, "CustomFields")
	delete(f, "DeliveryConditions")
	delete(f, "Sample")
	delete(f, "LogFormat")
	delete(f, "CustomEndpoint")
	delete(f, "S3")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyRealtimeLogDeliveryTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyRealtimeLogDeliveryTaskResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyRealtimeLogDeliveryTaskResponse struct {
	*tchttp.BaseResponse
	Response *ModifyRealtimeLogDeliveryTaskResponseParams `json:"Response"`
}

func (r *ModifyRealtimeLogDeliveryTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRealtimeLogDeliveryTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyRuleRequestParams struct {
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 规则名称，字符串名称长度 1~255。
	RuleName *string `json:"RuleName,omitnil,omitempty" name:"RuleName"`

	// 规则内容。
	Rules []*Rule `json:"Rules,omitnil,omitempty" name:"Rules"`

	// 规则 ID。
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// 规则状态，取值有：
	// <li> enable: 启用； </li>
	// <li> disable: 未启用。</li>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 规则标签。
	Tags []*string `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type ModifyRuleRequest struct {
	*tchttp.BaseRequest
	
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 规则名称，字符串名称长度 1~255。
	RuleName *string `json:"RuleName,omitnil,omitempty" name:"RuleName"`

	// 规则内容。
	Rules []*Rule `json:"Rules,omitnil,omitempty" name:"Rules"`

	// 规则 ID。
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// 规则状态，取值有：
	// <li> enable: 启用； </li>
	// <li> disable: 未启用。</li>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 规则标签。
	Tags []*string `json:"Tags,omitnil,omitempty" name:"Tags"`
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
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyRuleResponseParams struct {
	// 规则 ID。
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type ModifySecurityIPGroupRequestParams struct {
	// 站点 Id。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// IP 组配置。
	IPGroup *IPGroup `json:"IPGroup,omitnil,omitempty" name:"IPGroup"`

	// 操作类型，取值有：
	// <li> append: 向 IPGroup 中追加 Content 参数中内容；</li>
	// <li> remove: 从 IPGroup 中删除 Content 参数中内容；</li>
	// <li> update: 全量替换 IPGroup 内容，并可修改 IPGroup 名称。 </li>
	Mode *string `json:"Mode,omitnil,omitempty" name:"Mode"`
}

type ModifySecurityIPGroupRequest struct {
	*tchttp.BaseRequest
	
	// 站点 Id。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// IP 组配置。
	IPGroup *IPGroup `json:"IPGroup,omitnil,omitempty" name:"IPGroup"`

	// 操作类型，取值有：
	// <li> append: 向 IPGroup 中追加 Content 参数中内容；</li>
	// <li> remove: 从 IPGroup 中删除 Content 参数中内容；</li>
	// <li> update: 全量替换 IPGroup 内容，并可修改 IPGroup 名称。 </li>
	Mode *string `json:"Mode,omitnil,omitempty" name:"Mode"`
}

func (r *ModifySecurityIPGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySecurityIPGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "IPGroup")
	delete(f, "Mode")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifySecurityIPGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySecurityIPGroupResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifySecurityIPGroupResponse struct {
	*tchttp.BaseResponse
	Response *ModifySecurityIPGroupResponseParams `json:"Response"`
}

func (r *ModifySecurityIPGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySecurityIPGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySecurityPolicyRequestParams struct {
	// 站点Id。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 安全配置。
	SecurityConfig *SecurityConfig `json:"SecurityConfig,omitnil,omitempty" name:"SecurityConfig"`

	// 子域名/应用名。
	// 
	// 注意：当同时指定本参数和 TemplateId 参数时，本参数不生效。请勿同时指定本参数和 TemplateId 参数。
	Entity *string `json:"Entity,omitnil,omitempty" name:"Entity"`

	// 指定模板策略 ID，或指定站点全局策略。
	// - 如需配置策略模板，请指定策略模板 ID。
	// - 如需配置站点全局策略，请使用 @ZoneLevel@Domain 参数值
	// 
	// 注意：当使用本参数时，Entity 参数不生效。请勿同时使用本参数和 Entity 参数。
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`
}

type ModifySecurityPolicyRequest struct {
	*tchttp.BaseRequest
	
	// 站点Id。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 安全配置。
	SecurityConfig *SecurityConfig `json:"SecurityConfig,omitnil,omitempty" name:"SecurityConfig"`

	// 子域名/应用名。
	// 
	// 注意：当同时指定本参数和 TemplateId 参数时，本参数不生效。请勿同时指定本参数和 TemplateId 参数。
	Entity *string `json:"Entity,omitnil,omitempty" name:"Entity"`

	// 指定模板策略 ID，或指定站点全局策略。
	// - 如需配置策略模板，请指定策略模板 ID。
	// - 如需配置站点全局策略，请使用 @ZoneLevel@Domain 参数值
	// 
	// 注意：当使用本参数时，Entity 参数不生效。请勿同时使用本参数和 Entity 参数。
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`
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
	delete(f, "SecurityConfig")
	delete(f, "Entity")
	delete(f, "TemplateId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifySecurityPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySecurityPolicyResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type ModifyZoneRequestParams struct {
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 站点接入方式，取值有：
	// <li>full：NS 接入；</li>
	// <li>partial：CNAME 接入，如果站点当前是无域名接入，仅支持切换到 CNAME 接入；</li>
	// <li>dnsPodAccess：DNSPod 托管接入，该接入模式要求您的域名已托管在 DNSPod 内。</li>不填写保持原有配置。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 自定义站点信息，以替代系统默认分配的名称服务器。不填写保持原有配置。当站点是无域名接入方式时不允许传此参数。
	VanityNameServers *VanityNameServers `json:"VanityNameServers,omitnil,omitempty" name:"VanityNameServers"`

	// 站点别名。数字、英文、-和_组合，限制20个字符。
	AliasZoneName *string `json:"AliasZoneName,omitnil,omitempty" name:"AliasZoneName"`

	// 站点接入地域，取值有：
	// <li> global：全球；</li>
	// <li> mainland：中国大陆；</li>
	// <li> overseas：境外区域。</li>当站点是无域名接入方式时，不允许传此参数。
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 站点名称。仅当站点由无域名接入方式切换到CNAME接入方式的场景下有效。
	ZoneName *string `json:"ZoneName,omitnil,omitempty" name:"ZoneName"`
}

type ModifyZoneRequest struct {
	*tchttp.BaseRequest
	
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 站点接入方式，取值有：
	// <li>full：NS 接入；</li>
	// <li>partial：CNAME 接入，如果站点当前是无域名接入，仅支持切换到 CNAME 接入；</li>
	// <li>dnsPodAccess：DNSPod 托管接入，该接入模式要求您的域名已托管在 DNSPod 内。</li>不填写保持原有配置。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 自定义站点信息，以替代系统默认分配的名称服务器。不填写保持原有配置。当站点是无域名接入方式时不允许传此参数。
	VanityNameServers *VanityNameServers `json:"VanityNameServers,omitnil,omitempty" name:"VanityNameServers"`

	// 站点别名。数字、英文、-和_组合，限制20个字符。
	AliasZoneName *string `json:"AliasZoneName,omitnil,omitempty" name:"AliasZoneName"`

	// 站点接入地域，取值有：
	// <li> global：全球；</li>
	// <li> mainland：中国大陆；</li>
	// <li> overseas：境外区域。</li>当站点是无域名接入方式时，不允许传此参数。
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 站点名称。仅当站点由无域名接入方式切换到CNAME接入方式的场景下有效。
	ZoneName *string `json:"ZoneName,omitnil,omitempty" name:"ZoneName"`
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
	delete(f, "AliasZoneName")
	delete(f, "Area")
	delete(f, "ZoneName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyZoneRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyZoneResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	// 待变更的站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 缓存过期时间配置。
	// 不填写表示保持原有配置。
	CacheConfig *CacheConfig `json:"CacheConfig,omitnil,omitempty" name:"CacheConfig"`

	// 节点缓存键配置。
	// 不填写表示保持原有配置。
	CacheKey *CacheKey `json:"CacheKey,omitnil,omitempty" name:"CacheKey"`

	// 浏览器缓存配置。
	// 不填写表示保持原有配置。
	MaxAge *MaxAge `json:"MaxAge,omitnil,omitempty" name:"MaxAge"`

	// 离线缓存配置。
	// 不填写表示保持原有配置。
	OfflineCache *OfflineCache `json:"OfflineCache,omitnil,omitempty" name:"OfflineCache"`

	// Quic 访问配置。
	// 不填写表示保持原有配置。
	Quic *Quic `json:"Quic,omitnil,omitempty" name:"Quic"`

	// Post 请求传输配置。
	// 不填写表示保持原有配置。
	PostMaxSize *PostMaxSize `json:"PostMaxSize,omitnil,omitempty" name:"PostMaxSize"`

	// 智能压缩配置。
	// 不填写表示保持原有配置。
	Compression *Compression `json:"Compression,omitnil,omitempty" name:"Compression"`

	// Http2 回源配置。
	// 不填写表示保持原有配置。
	UpstreamHttp2 *UpstreamHttp2 `json:"UpstreamHttp2,omitnil,omitempty" name:"UpstreamHttp2"`

	// 访问协议强制 Https 跳转配置。
	// 不填写表示保持原有配置。
	ForceRedirect *ForceRedirect `json:"ForceRedirect,omitnil,omitempty" name:"ForceRedirect"`

	// Https 加速配置。
	// 不填写表示保持原有配置。
	Https *Https `json:"Https,omitnil,omitempty" name:"Https"`

	// 源站配置。
	// 不填写表示保持原有配置。
	Origin *Origin `json:"Origin,omitnil,omitempty" name:"Origin"`

	// 智能加速配置。
	// 不填写表示保持原有配置。
	SmartRouting *SmartRouting `json:"SmartRouting,omitnil,omitempty" name:"SmartRouting"`

	// WebSocket 配置。
	// 不填写表示保持原有配置。
	WebSocket *WebSocket `json:"WebSocket,omitnil,omitempty" name:"WebSocket"`

	// 客户端 IP 回源请求头配置。
	// 不填写表示保持原有配置。
	ClientIpHeader *ClientIpHeader `json:"ClientIpHeader,omitnil,omitempty" name:"ClientIpHeader"`

	// 缓存预刷新配置。
	// 不填写表示保持原有配置。
	CachePrefresh *CachePrefresh `json:"CachePrefresh,omitnil,omitempty" name:"CachePrefresh"`

	// Ipv6 访问配置。
	// 不填写表示保持原有配置。
	Ipv6 *Ipv6 `json:"Ipv6,omitnil,omitempty" name:"Ipv6"`

	// 回源时是否携带客户端 IP 所属地域信息的配置。
	// 不填写表示保持原有配置。
	ClientIpCountry *ClientIpCountry `json:"ClientIpCountry,omitnil,omitempty" name:"ClientIpCountry"`

	// Grpc 协议支持配置。
	// 不填写表示保持原有配置。
	Grpc *Grpc `json:"Grpc,omitnil,omitempty" name:"Grpc"`

	// 图片优化配置。
	// 不填写表示关闭。
	ImageOptimize *ImageOptimize `json:"ImageOptimize,omitnil,omitempty" name:"ImageOptimize"`

	// 标准 Debug 配置。
	StandardDebug *StandardDebug `json:"StandardDebug,omitnil,omitempty" name:"StandardDebug"`
}

type ModifyZoneSettingRequest struct {
	*tchttp.BaseRequest
	
	// 待变更的站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 缓存过期时间配置。
	// 不填写表示保持原有配置。
	CacheConfig *CacheConfig `json:"CacheConfig,omitnil,omitempty" name:"CacheConfig"`

	// 节点缓存键配置。
	// 不填写表示保持原有配置。
	CacheKey *CacheKey `json:"CacheKey,omitnil,omitempty" name:"CacheKey"`

	// 浏览器缓存配置。
	// 不填写表示保持原有配置。
	MaxAge *MaxAge `json:"MaxAge,omitnil,omitempty" name:"MaxAge"`

	// 离线缓存配置。
	// 不填写表示保持原有配置。
	OfflineCache *OfflineCache `json:"OfflineCache,omitnil,omitempty" name:"OfflineCache"`

	// Quic 访问配置。
	// 不填写表示保持原有配置。
	Quic *Quic `json:"Quic,omitnil,omitempty" name:"Quic"`

	// Post 请求传输配置。
	// 不填写表示保持原有配置。
	PostMaxSize *PostMaxSize `json:"PostMaxSize,omitnil,omitempty" name:"PostMaxSize"`

	// 智能压缩配置。
	// 不填写表示保持原有配置。
	Compression *Compression `json:"Compression,omitnil,omitempty" name:"Compression"`

	// Http2 回源配置。
	// 不填写表示保持原有配置。
	UpstreamHttp2 *UpstreamHttp2 `json:"UpstreamHttp2,omitnil,omitempty" name:"UpstreamHttp2"`

	// 访问协议强制 Https 跳转配置。
	// 不填写表示保持原有配置。
	ForceRedirect *ForceRedirect `json:"ForceRedirect,omitnil,omitempty" name:"ForceRedirect"`

	// Https 加速配置。
	// 不填写表示保持原有配置。
	Https *Https `json:"Https,omitnil,omitempty" name:"Https"`

	// 源站配置。
	// 不填写表示保持原有配置。
	Origin *Origin `json:"Origin,omitnil,omitempty" name:"Origin"`

	// 智能加速配置。
	// 不填写表示保持原有配置。
	SmartRouting *SmartRouting `json:"SmartRouting,omitnil,omitempty" name:"SmartRouting"`

	// WebSocket 配置。
	// 不填写表示保持原有配置。
	WebSocket *WebSocket `json:"WebSocket,omitnil,omitempty" name:"WebSocket"`

	// 客户端 IP 回源请求头配置。
	// 不填写表示保持原有配置。
	ClientIpHeader *ClientIpHeader `json:"ClientIpHeader,omitnil,omitempty" name:"ClientIpHeader"`

	// 缓存预刷新配置。
	// 不填写表示保持原有配置。
	CachePrefresh *CachePrefresh `json:"CachePrefresh,omitnil,omitempty" name:"CachePrefresh"`

	// Ipv6 访问配置。
	// 不填写表示保持原有配置。
	Ipv6 *Ipv6 `json:"Ipv6,omitnil,omitempty" name:"Ipv6"`

	// 回源时是否携带客户端 IP 所属地域信息的配置。
	// 不填写表示保持原有配置。
	ClientIpCountry *ClientIpCountry `json:"ClientIpCountry,omitnil,omitempty" name:"ClientIpCountry"`

	// Grpc 协议支持配置。
	// 不填写表示保持原有配置。
	Grpc *Grpc `json:"Grpc,omitnil,omitempty" name:"Grpc"`

	// 图片优化配置。
	// 不填写表示关闭。
	ImageOptimize *ImageOptimize `json:"ImageOptimize,omitnil,omitempty" name:"ImageOptimize"`

	// 标准 Debug 配置。
	StandardDebug *StandardDebug `json:"StandardDebug,omitnil,omitempty" name:"StandardDebug"`
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
	delete(f, "ClientIpCountry")
	delete(f, "Grpc")
	delete(f, "ImageOptimize")
	delete(f, "StandardDebug")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyZoneSettingRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyZoneSettingResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 站点状态，取值有：
	// <li> false：开启站点；</li>
	// <li> true：关闭站点。</li>
	Paused *bool `json:"Paused,omitnil,omitempty" name:"Paused"`
}

type ModifyZoneStatusRequest struct {
	*tchttp.BaseRequest
	
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 站点状态，取值有：
	// <li> false：开启站点；</li>
	// <li> true：关闭站点。</li>
	Paused *bool `json:"Paused,omitnil,omitempty" name:"Paused"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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

type MutualTLS struct {
	// 双向认证配置开关，取值有：
	// <li>on：开启；</li>
	// <li>off：关闭。</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// 双向认证证书列表。
	// 注意：MutualTLS 在 ModifyHostsCertificate 作为入参使用时，该参数传入对应证书的 CertId 即可。您可以前往 [SSL 证书列表](https://console.cloud.tencent.com/certoverview) 查看 CertId。
	CertInfos []*CertificateInfo `json:"CertInfos,omitnil,omitempty" name:"CertInfos"`
}

type NoCache struct {
	// 不缓存配置开关，取值有：
	// <li>on：开启；</li>
	// <li>off：关闭。</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`
}

type NormalAction struct {
	// 功能名称，功能名称填写规范可调用接口 [查询规则引擎的设置参数](https://cloud.tencent.com/document/product/1552/80618) 查看。
	Action *string `json:"Action,omitnil,omitempty" name:"Action"`

	// 参数。
	Parameters []*RuleNormalActionParams `json:"Parameters,omitnil,omitempty" name:"Parameters"`
}

type NsVerification struct {
	// NS 接入时，分配给用户的 DNS 服务器地址，需要将域名的 NameServer 切换至该地址。
	NameServers []*string `json:"NameServers,omitnil,omitempty" name:"NameServers"`
}

type OfflineCache struct {
	// 离线缓存是否开启，取值有：
	// <li>on：开启；</li>
	// <li>off：关闭。</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`
}

type Origin struct {
	// 主源站列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Origins []*string `json:"Origins,omitnil,omitempty" name:"Origins"`

	// 备源站列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	BackupOrigins []*string `json:"BackupOrigins,omitnil,omitempty" name:"BackupOrigins"`

	// 回源协议配置，取值有：
	// <li>http：强制 http 回源；</li>
	// <li>follow：协议跟随回源；</li>
	// <li>https：强制 https 回源。</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	OriginPullProtocol *string `json:"OriginPullProtocol,omitnil,omitempty" name:"OriginPullProtocol"`

	// 源站为腾讯云 COS 时，是否为私有访问 bucket，取值有：
	// <li>on：私有访问；</li>
	// <li>off：公共访问。</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	CosPrivateAccess *string `json:"CosPrivateAccess,omitnil,omitempty" name:"CosPrivateAccess"`
}

type OriginDetail struct {
	// 源站类型，取值有：
	// <li>IP_DOMAIN：IPV4、IPV6 或域名类型源站；</li>
	// <li>COS：腾讯云 COS 对象存储源站；</li>
	// <li>AWS_S3：AWS S3 对象存储源站；</li>
	// <li>ORIGIN_GROUP：源站组类型源站；</li>
	// <li>VOD：云点播；</li>
	// <li>SPACE：源站卸载，当前仅白名单开放；</li>
	// <li>LB：负载均衡，当前仅白名单开放。</li>
	OriginType *string `json:"OriginType,omitnil,omitempty" name:"OriginType"`

	// 源站地址，根据 OriginType 的取值分为以下情况：
	// <li>当 OriginType = IP_DOMAIN 时，该参数为 IPv4、IPv6 地址或域名；</li>
	// <li>当 OriginType = COS 时，该参数为 COS 桶的访问域名；</li>
	// <li>当 OriginType = AWS_S3，该参数为 S3 桶的访问域名；</li>
	// <li>当 OriginType = ORIGIN_GROUP 时，该参数为源站组 ID；</li>
	// <li>当 OriginType = VOD 时，该参数请填写云点播应用 ID ；</li>
	Origin *string `json:"Origin,omitnil,omitempty" name:"Origin"`

	// 备用源站组 ID，该参数仅在 OriginType = ORIGIN_GROUP 且配置了备源站组时会生效。
	BackupOrigin *string `json:"BackupOrigin,omitnil,omitempty" name:"BackupOrigin"`

	// 主源源站组名称，当 OriginType = ORIGIN_GROUP 时该参数会返回值。
	OriginGroupName *string `json:"OriginGroupName,omitnil,omitempty" name:"OriginGroupName"`

	// 备用源站组名称，该参数仅当 OriginType = ORIGIN_GROUP 且配置了备用源站组时会生效。
	BackOriginGroupName *string `json:"BackOriginGroupName,omitnil,omitempty" name:"BackOriginGroupName"`

	// 指定是否允许访问私有对象存储源站，该参数仅当源站类型OriginType = COS 或 AWS_S3 时会生效，取值有：
	// <li>on：使用私有鉴权；</li>
	// <li>off：不使用私有鉴权。</li>
	// 不填写，默认值为 off。
	PrivateAccess *string `json:"PrivateAccess,omitnil,omitempty" name:"PrivateAccess"`

	// 私有鉴权使用参数，该参数仅当源站类型 PrivateAccess = on 时会生效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PrivateParameters []*PrivateParameter `json:"PrivateParameters,omitnil,omitempty" name:"PrivateParameters"`

	// MO 子应用 ID
	//
	// Deprecated: VodeoSubAppId is deprecated.
	VodeoSubAppId *int64 `json:"VodeoSubAppId,omitnil,omitempty" name:"VodeoSubAppId"`

	// MO 分发范围，取值有： <li>All：全部</li> <li>Bucket：存储桶</li>
	//
	// Deprecated: VodeoDistributionRange is deprecated.
	VodeoDistributionRange *string `json:"VodeoDistributionRange,omitnil,omitempty" name:"VodeoDistributionRange"`

	// MO 存储桶 ID，分发范围(DistributionRange)为存储桶(Bucket)时必填
	//
	// Deprecated: VodeoBucketId is deprecated.
	VodeoBucketId *string `json:"VodeoBucketId,omitnil,omitempty" name:"VodeoBucketId"`
}

type OriginGroup struct {
	// 源站组ID。
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 源站组名称。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 源站组类型，取值有：
	// <li>GENERAL：通用型源站组；</li>
	// <li>HTTP： HTTP专用型源站组。</li>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 源站记录信息。
	Records []*OriginRecord `json:"Records,omitnil,omitempty" name:"Records"`

	// 源站组被引用实例列表。	
	References []*OriginGroupReference `json:"References,omitnil,omitempty" name:"References"`

	// 源站组创建时间。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 源站组更新时间。
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 回源Host Header。
	// 注意：此字段可能返回 null，表示取不到有效值。
	HostHeader *string `json:"HostHeader,omitnil,omitempty" name:"HostHeader"`
}

type OriginGroupReference struct {
	// 引用服务类型，取值有：
	// <li>AccelerationDomain: 加速域名；</li>
	// <li>RuleEngine: 规则引擎；</li>
	// <li>Loadbalance: 负载均衡；</li>
	// <li>ApplicationProxy: 四层代理。</li>
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// 引用类型的实例ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 应用类型的实例名称。
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`
}

type OriginInfo struct {
	// 源站类型，取值有：
	// <li>IP_DOMAIN：IPV4、IPV6 或域名类型源站；</li>
	// <li>COS：腾讯云 COS 对象存储源站；</li>
	// <li>AWS_S3：AWS S3 对象存储源站；</li>
	// <li>ORIGIN_GROUP：源站组类型源站；</li>
	//  <li>VOD：云点播；</li>
	// <li>SPACE：源站卸载，当前仅白名单开放；</li>
	// <li>LB：负载均衡，当前仅白名单开放。</li>
	OriginType *string `json:"OriginType,omitnil,omitempty" name:"OriginType"`

	// 源站地址，根据 OriginType 的取值分为以下情况：
	// <li>当 OriginType = IP_DOMAIN 时，该参数请填写 IPv4、IPv6 地址或域名；</li>
	// <li>当 OriginType = COS 时，该参数请填写 COS 桶的访问域名；</li>
	// <li>当 OriginType = AWS_S3，该参数请填写 S3 桶的访问域名；</li>
	// <li>当 OriginType = ORIGIN_GROUP 时，该参数请填写源站组 ID；</li>
	// <li>当 OriginType = VOD 时，该参数请填写云点播应用 ID ；</li>
	// <li>当 OriginType = LB 时，该参数请填写负载均衡实例 ID，该功能当前仅白名单开放；</li>
	// <li>当 OriginType = SPACE 时，该参数请填写源站卸载空间 ID，该功能当前仅白名单开放。</li>
	Origin *string `json:"Origin,omitnil,omitempty" name:"Origin"`

	// 备用源站组 ID，该参数仅在 OriginType = ORIGIN_GROUP 时生效，该字段为旧版能力，调用后控制台无法进行配置修改，如需使用请提交工单咨询。
	BackupOrigin *string `json:"BackupOrigin,omitnil,omitempty" name:"BackupOrigin"`

	// 指定是否允许访问私有对象存储源站，该参数仅当源站类型 OriginType = COS 或 AWS_S3 时会生效，取值有：
	// <li>on：使用私有鉴权；</li>
	// <li>off：不使用私有鉴权。</li>
	// 不填写时，默认值为off。
	PrivateAccess *string `json:"PrivateAccess,omitnil,omitempty" name:"PrivateAccess"`

	// 私有鉴权使用参数，该参数仅当源站类型 PrivateAccess = on 时会生效。
	PrivateParameters []*PrivateParameter `json:"PrivateParameters,omitnil,omitempty" name:"PrivateParameters"`

	// VODEO 子应用 ID。该参数当 OriginType = VODEO 时必填。
	//
	// Deprecated: VodeoSubAppId is deprecated.
	VodeoSubAppId *int64 `json:"VodeoSubAppId,omitnil,omitempty" name:"VodeoSubAppId"`

	// VODEO 分发范围，该参数当 OriginType = VODEO 时必填。取值有： 
	// <li>All：当前应用下所有存储桶；</li> 
	// <li>Bucket：指定的某一个存储桶。</li>
	//
	// Deprecated: VodeoDistributionRange is deprecated.
	VodeoDistributionRange *string `json:"VodeoDistributionRange,omitnil,omitempty" name:"VodeoDistributionRange"`

	// VODEO 存储桶 ID，该参数当 OriginType = VODEO 且 VodeoDistributionRange = Bucket 时必填。
	//
	// Deprecated: VodeoBucketId is deprecated.
	VodeoBucketId *string `json:"VodeoBucketId,omitnil,omitempty" name:"VodeoBucketId"`
}

type OriginProtectionInfo struct {
	// 站点ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 域名列表。
	Hosts []*string `json:"Hosts,omitnil,omitempty" name:"Hosts"`

	// 代理ID列表。
	ProxyIds []*string `json:"ProxyIds,omitnil,omitempty" name:"ProxyIds"`

	// 当前版本的IP白名单。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CurrentIPWhitelist *IPWhitelist `json:"CurrentIPWhitelist,omitnil,omitempty" name:"CurrentIPWhitelist"`

	// 该站点是否需要更新源站白名单，取值有：
	// <li>true ：需要更新IP白名单 ；</li>
	// <li>false ：无需更新IP白名单。</li>
	NeedUpdate *bool `json:"NeedUpdate,omitnil,omitempty" name:"NeedUpdate"`

	// 源站防护状态，取值有：
	// <li>online ：源站防护启用中 ；</li>
	// <li>offline ：源站防护已停用 ；</li>
	// <li>nonactivate ：源站防护未激活，仅在从未使用过源站防护功能的站点调用中返回。</li>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 站点套餐是否支持源站防护，取值有：
	// <li>true ：支持 ；</li>
	// <li>false ：不支持。</li>
	PlanSupport *bool `json:"PlanSupport,omitnil,omitempty" name:"PlanSupport"`

	// 最新IP白名单与当前IP白名单的对比。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiffIPWhitelist *DiffIPWhitelist `json:"DiffIPWhitelist,omitnil,omitempty" name:"DiffIPWhitelist"`
}

type OriginRecord struct {
	// 源站记录值，不包含端口信息，可以为：IPv4，IPv6，域名格式。
	Record *string `json:"Record,omitnil,omitempty" name:"Record"`

	// 源站类型，取值有：
	// <li>IP_DOMAIN：IPV4、IPV6、域名类型源站；</li>
	// <li>COS：COS源。</li>
	// <li>AWS_S3：AWS S3对象存储源站。</li>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 源站记录ID。
	RecordId *string `json:"RecordId,omitnil,omitempty" name:"RecordId"`

	// 源站权重，取值为0-100, 不填表示不设置权重，由系统自由调度，填0表示权重为0, 流量将不会调度到此源站。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Weight *uint64 `json:"Weight,omitnil,omitempty" name:"Weight"`

	// 是否私有鉴权，当源站类型 RecordType=COS/AWS_S3 时生效，取值有：
	// <li>true：使用私有鉴权；</li>
	// <li>false：不使用私有鉴权。</li>不填写，默认值为：false。
	Private *bool `json:"Private,omitnil,omitempty" name:"Private"`

	// 私有鉴权参数，当源站类型Private=true时有效。
	PrivateParameters []*PrivateParameter `json:"PrivateParameters,omitnil,omitempty" name:"PrivateParameters"`
}

type OwnershipVerification struct {
	// CNAME 接入，使用 DNS 解析验证时所需的信息。详情参考 [站点/域名归属权验证
	// ](https://cloud.tencent.com/document/product/1552/70789#7af6ecf8-afca-4e35-8811-b5797ed1bde5)。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DnsVerification *DnsVerification `json:"DnsVerification,omitnil,omitempty" name:"DnsVerification"`

	// CNAME 接入，使用文件验证时所需的信息。详情参考 [站点/域名归属权验证
	// ](https://cloud.tencent.com/document/product/1552/70789#7af6ecf8-afca-4e35-8811-b5797ed1bde5)。
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileVerification *FileVerification `json:"FileVerification,omitnil,omitempty" name:"FileVerification"`

	// NS 接入，切换 DNS 服务器所需的信息。详情参考 [修改 DNS 服务器](https://cloud.tencent.com/document/product/1552/90452)。
	// 注意：此字段可能返回 null，表示取不到有效值。
	NsVerification *NsVerification `json:"NsVerification,omitnil,omitempty" name:"NsVerification"`
}

type PartialModule struct {
	// 模块名称，取值为：
	// <li>waf：托管规则。</li>
	Module *string `json:"Module,omitnil,omitempty" name:"Module"`

	// 模块下的需要例外的具体规则ID列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Include []*int64 `json:"Include,omitnil,omitempty" name:"Include"`
}

type PlanInfo struct {
	// 结算货币类型，取值有：
	// <li> CNY ：人民币结算； </li>
	// <li> USD ：美元结算。</li>
	Currency *string `json:"Currency,omitnil,omitempty" name:"Currency"`

	// 套餐所含流量，该流量数值为安全加速流量，内容加速流量和智能加速流量的总和（单位：字节）。
	Flux *uint64 `json:"Flux,omitnil,omitempty" name:"Flux"`

	// 结算周期，取值有：
	// <li> y ：按年结算； </li>
	// <li> m ：按月结算；</li>
	// <li> h ：按小时结算； </li>
	// <li> M ：按分钟结算；</li>
	// <li> s ：按秒结算。 </li>
	Frequency *string `json:"Frequency,omitnil,omitempty" name:"Frequency"`

	// 套餐类型，取值有：
	// <li> sta ：全球内容分发网络（不包括中国大陆）标准版套餐； </li>
	// <li> sta_with_bot ：全球内容分发网络（不包括中国大陆）标准版套餐附带bot管理；</li>
	// <li> sta_cm ：中国大陆内容分发网络标准版套餐； </li>
	// <li> sta_cm_with_bot ：中国大陆内容分发网络标准版套餐附带bot管理；</li>
	// <li> sta_global ：全球内容分发网络（包括中国大陆）标准版套餐； </li>
	// <li> sta_global_with_bot ：全球内容分发网络（包括中国大陆）标准版套餐附带bot管理；</li>
	// <li> ent ：全球内容分发网络（不包括中国大陆）企业版套餐； </li>
	// <li> ent_with_bot ： 全球内容分发网络（不包括中国大陆）企业版套餐附带bot管理；</li>
	// <li> ent_cm ：中国大陆内容分发网络企业版套餐； </li>
	// <li> ent_cm_with_bot ：中国大陆内容分发网络企业版套餐附带bot管理；</li>
	// <li> ent_global ：全球内容分发网络（包括中国大陆）企业版套餐； </li>
	// <li> ent_global_with_bot ：全球内容分发网络（包括中国大陆）企业版套餐附带bot管理。</li>
	PlanType *string `json:"PlanType,omitnil,omitempty" name:"PlanType"`

	// 套餐价格（单位：分）。
	Price *float64 `json:"Price,omitnil,omitempty" name:"Price"`

	// 套餐所含请求次数，该请求次数为安全加速请求次数。（单位：次）。
	Request *uint64 `json:"Request,omitnil,omitempty" name:"Request"`

	// 套餐所能绑定的站点个数。
	SiteNumber *uint64 `json:"SiteNumber,omitnil,omitempty" name:"SiteNumber"`

	// 套餐加速区域类型，取值有：
	// <li> mainland ：中国大陆； </li>
	// <li> overseas ：全球（不包括中国大陆）；</li>
	// <li> global ：全球（包括中国大陆）。 </li>
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`
}

type PostMaxSize struct {
	// 是否开启 POST 请求上传文件限制，平台默认为限制为32MB，取值有：
	// <li>on：开启限制；</li>
	// <li>off：关闭限制。</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// 最大限制，取值在1MB和500MB之间。单位字节。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxSize *int64 `json:"MaxSize,omitnil,omitempty" name:"MaxSize"`
}

type PrepaidPlanParam struct {
	// 订阅预付费套餐的周期，单位：月，取值有：1，2，3，4，5，6，7，8，9，10，11，12，24，36。
	// 
	// 不填写使用默认值 1。
	Period *int64 `json:"Period,omitnil,omitempty" name:"Period"`

	// 预付费套餐的自动续费标志，取值有：
	// <li> on：开启自动续费；</li>
	// <li> off：不开启自动续费。</li>
	// 不填写使用默认值 off，自动续费时，默认续费1个月。
	RenewFlag *string `json:"RenewFlag,omitnil,omitempty" name:"RenewFlag"`
}

type PrivateParameter struct {
	// 私有鉴权参数名称，取值有：
	// <li>AccessKeyId：鉴权参数 Access Key ID；</li>
	// <li>SecretAccessKey：鉴权参数 Secret Access Key；</li>
	// <li>SignatureVersion：鉴权版本，v2 或者 v4；</li>
	// <li>Region：存储桶地域。</li>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 私有鉴权参数值。
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type QueryCondition struct {
	// 筛选条件的key。
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 查询条件操作符，操作类型有：
	// <li>equals: 等于；</li>
	// <li>notEquals: 不等于；</li>
	// <li>include: 包含；</li>
	// <li>notInclude: 不包含; </li>
	// <li>startWith: 开始的值是value；</li>
	// <li>notStartWith: 不以value的值开始；</li>
	// <li>endWith: 结尾是value值；</li>
	// <li>notEndWith: 不以value的值结尾。</li>
	Operator *string `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 筛选条件的值。
	Value []*string `json:"Value,omitnil,omitempty" name:"Value"`
}

type QueryString struct {
	// CacheKey是否由QueryString组成，取值有：
	// <li>on：是；</li>
	// <li>off：否。</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// CacheKey使用QueryString的方式，取值有：
	// <li>includeCustom：使用部分url参数；</li>
	// <li>excludeCustom：排除部分url参数。</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Action *string `json:"Action,omitnil,omitempty" name:"Action"`

	// 使用/排除的url参数数组。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value []*string `json:"Value,omitnil,omitempty" name:"Value"`
}

type Quic struct {
	// 是否开启 Quic 配置，取值有：
	// <li>on：开启；</li>
	// <li>off：关闭。</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`
}

type Quota struct {
	// 单次批量提交配额上限。
	Batch *int64 `json:"Batch,omitnil,omitempty" name:"Batch"`

	// 每日提交配额上限。
	Daily *int64 `json:"Daily,omitnil,omitempty" name:"Daily"`

	// 每日剩余的可提交配额。
	DailyAvailable *int64 `json:"DailyAvailable,omitnil,omitempty" name:"DailyAvailable"`

	// 刷新预热缓存类型，取值有：
	// <li> purge_prefix：按前缀刷新；</li>
	// <li> purge_url：按URL刷新；</li>
	// <li> purge_host：按Hostname刷新；</li>
	// <li> purge_all：刷新全部缓存内容；</li>
	// <li> purge_cache_tag：按CacheTag刷新；</li><li> prefetch_url：按URL预热。</li>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

type RateLimitConfig struct {
	// 开关，取值有：
	// <li>on：开启；</li>
	// <li>off：关闭。</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// 速率限制-用户规则列表。如果为null，默认使用历史配置。
	RateLimitUserRules []*RateLimitUserRule `json:"RateLimitUserRules,omitnil,omitempty" name:"RateLimitUserRules"`

	// 速率限制模板功能。如果为null，默认使用历史配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RateLimitTemplate *RateLimitTemplate `json:"RateLimitTemplate,omitnil,omitempty" name:"RateLimitTemplate"`

	// 智能客户端过滤。如果为null，默认使用历史配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RateLimitIntelligence *RateLimitIntelligence `json:"RateLimitIntelligence,omitnil,omitempty" name:"RateLimitIntelligence"`

	// 速率限制-托管定制规则。如果为null，默认使用历史配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RateLimitCustomizes []*RateLimitUserRule `json:"RateLimitCustomizes,omitnil,omitempty" name:"RateLimitCustomizes"`
}

type RateLimitIntelligence struct {
	// 功能开关，取值有：
	// <li>on：开启；</li>
	// <li>off：关闭。</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// 执行动作，取值有：
	// <li>monitor：观察；</li>
	// <li>alg：挑战。</li>
	Action *string `json:"Action,omitnil,omitempty" name:"Action"`

	// 规则id，仅出参使用。
	RuleId *int64 `json:"RuleId,omitnil,omitempty" name:"RuleId"`
}

type RateLimitTemplate struct {
	// 模板等级名称，取值有：
	// <li>sup_loose：超级宽松；</li>
	// <li>loose：宽松；</li>
	// <li>emergency：紧急；</li>
	// <li>normal：适中；</li>
	// <li>strict：严格；</li>
	// <li>close：关闭，仅精准速率限制生效。</li>
	Mode *string `json:"Mode,omitnil,omitempty" name:"Mode"`

	// 模板处置方式，取值有：
	// <li>alg：JavaScript挑战；</li>
	// <li>monitor：观察。</li>不填写默认取alg。
	Action *string `json:"Action,omitnil,omitempty" name:"Action"`

	// 模板值详情。仅出参返回。
	RateLimitTemplateDetail *RateLimitTemplateDetail `json:"RateLimitTemplateDetail,omitnil,omitempty" name:"RateLimitTemplateDetail"`
}

type RateLimitTemplateDetail struct {
	// 模板等级名称，取值有：
	// <li>sup_loose：超级宽松；</li>
	// <li>loose：宽松；</li>
	// <li>emergency：紧急；</li>
	// <li>normal：适中；</li>
	// <li>strict：严格；</li>
	// <li>close：关闭，仅精准速率限制生效。</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Mode *string `json:"Mode,omitnil,omitempty" name:"Mode"`

	// 唯一id。
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// 模板处置方式，取值有：
	// <li>alg：JavaScript挑战；</li>
	// <li>monitor：观察。</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Action *string `json:"Action,omitnil,omitempty" name:"Action"`

	// 惩罚时间，取值范围0-2天，单位秒。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PunishTime *int64 `json:"PunishTime,omitnil,omitempty" name:"PunishTime"`

	// 统计阈值，单位是次，取值范围0-4294967294。
	Threshold *int64 `json:"Threshold,omitnil,omitempty" name:"Threshold"`

	// 统计周期，取值范围0-120秒。
	Period *int64 `json:"Period,omitnil,omitempty" name:"Period"`
}

type RateLimitUserRule struct {
	// 速率限制统计阈值，单位是次，取值范围0-4294967294。
	Threshold *int64 `json:"Threshold,omitnil,omitempty" name:"Threshold"`

	// 速率限制统计时间，取值范围 10/20/30/40/50/60 单位是秒。
	Period *int64 `json:"Period,omitnil,omitempty" name:"Period"`

	// 规则名，只能以英文字符，数字，下划线组合，且不能以下划线开头。
	RuleName *string `json:"RuleName,omitnil,omitempty" name:"RuleName"`

	// 处置动作，取值有： <li>monitor：观察；</li> <li>drop：拦截；</li><li> redirect：重定向；</li><li> page：指定页面；</li><li>alg：JavaScript 挑战。</li>	
	Action *string `json:"Action,omitnil,omitempty" name:"Action"`

	// 惩罚时长，0-2天。
	PunishTime *int64 `json:"PunishTime,omitnil,omitempty" name:"PunishTime"`

	// 处罚时长单位，取值有：
	// <li>second：秒；</li>
	// <li>minutes：分钟；</li>
	// <li>hour：小时。</li>
	PunishTimeUnit *string `json:"PunishTimeUnit,omitnil,omitempty" name:"PunishTimeUnit"`

	// 规则状态，取值有：
	// <li>on：生效；</li>
	// <li>off：不生效。</li>默认 on 生效。
	RuleStatus *string `json:"RuleStatus,omitnil,omitempty" name:"RuleStatus"`

	// 规则详情。
	AclConditions []*AclCondition `json:"AclConditions,omitnil,omitempty" name:"AclConditions"`

	// 规则权重，取值范围0-100。
	RulePriority *int64 `json:"RulePriority,omitnil,omitempty" name:"RulePriority"`

	// 规则 Id。仅出参使用。
	RuleID *int64 `json:"RuleID,omitnil,omitempty" name:"RuleID"`

	// 过滤词，取值有：
	// <li>sip：客户端 ip。</li>
	// 默认为空字符串。
	FreqFields []*string `json:"FreqFields,omitnil,omitempty" name:"FreqFields"`

	// 更新时间。仅出参使用。修改时默认为当前时间。
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 统计范围。取值有：
	// <li>source_to_eo：（响应）源站到  EdgeOne；</li>
	// <li>client_to_eo：（请求）客户端到  EdgeOne。</li>
	// 默认为 source_to_eo。
	FreqScope []*string `json:"FreqScope,omitnil,omitempty" name:"FreqScope"`

	// 自定义返回页面的名称。Action 是 page 时必填，且不能为空。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 自定义响应 Id。该 Id 可通过查询自定义错误页列表接口获取。默认值为default，使用系统默认页面。Action 是 page 时必填，且不能为空。	
	CustomResponseId *string `json:"CustomResponseId,omitnil,omitempty" name:"CustomResponseId"`

	// 自定义返回页面的响应码。Action 是 page 时必填，且不能为空，取值: 100~600，不支持 3xx 响应码。默认值：567。
	ResponseCode *int64 `json:"ResponseCode,omitnil,omitempty" name:"ResponseCode"`

	// 重定向时候的地址。Action 是 redirect 时必填，且不能为空。
	RedirectUrl *string `json:"RedirectUrl,omitnil,omitempty" name:"RedirectUrl"`
}

type RealtimeLogDeliveryTask struct {
	// 实时日志投递任务 ID。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 实时日志投递任务的名称。
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`

	// 实时日志投递任务的状态，取值有： <li>enabled: 已启用；</li> <li>disabled: 已停用；</li><li>deleted: 异常删除状态，请检查目的地腾讯云 CLS 日志集/日志主题是否已被删除。</li>
	DeliveryStatus *string `json:"DeliveryStatus,omitnil,omitempty" name:"DeliveryStatus"`

	// 实时日志投递任务类型，取值有： <li>cls: 推送到腾讯云 CLS；</li> <li>custom_endpoint：推送到自定义 HTTP(S) 地址；</li> <li>s3：推送到 AWS S3 兼容存储桶地址。</li>
	TaskType *string `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// 实时日志投递任务对应的实体（七层域名或者四层代理实例）列表。取值示例如下： <li>七层域名：domain.example.com；</li> <li>四层代理实例：sid-2s69eb5wcms7。</li>	
	EntityList []*string `json:"EntityList,omitnil,omitempty" name:"EntityList"`

	// 数据投递类型，取值有： <li>domain：站点加速日志；</li> <li>application：四层代理日志；</li> <li>web-rateLiming：速率限制和 CC 攻击防护日志；</li> <li>web-attack：托管规则日志；</li> <li>web-rule：自定义规则日志；</li> <li>web-bot：Bot管理日志。</li>
	LogType *string `json:"LogType,omitnil,omitempty" name:"LogType"`

	// 数据投递区域，取值有： <li>mainland：中国大陆境内；</li> <li>overseas：全球（不含中国大陆）。</li>
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 投递的预设字段列表。
	Fields []*string `json:"Fields,omitnil,omitempty" name:"Fields"`

	// 投递的自定义字段列表。
	CustomFields []*CustomField `json:"CustomFields,omitnil,omitempty" name:"CustomFields"`

	// 日志投递的过滤条件。
	DeliveryConditions []*DeliveryCondition `json:"DeliveryConditions,omitnil,omitempty" name:"DeliveryConditions"`

	// 采样比例，采用千分制，取值范围为1-1000，例如：605 表示采样比例为 60.5%。
	Sample *uint64 `json:"Sample,omitnil,omitempty" name:"Sample"`

	// 日志投递的输出格式。出参为 null 时表示为默认格式，默认格式逻辑如下：
	// <li>当 TaskType 取值为 custom_endpoint 时，默认格式为多个 JSON 对象组成的数组，每个 JSON 对象为一条日志；</li>
	// <li>当 TaskType 取值为 s3 时，默认格式为 JSON Lines。</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	LogFormat *LogFormat `json:"LogFormat,omitnil,omitempty" name:"LogFormat"`

	// CLS 的配置信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CLS *CLSTopic `json:"CLS,omitnil,omitempty" name:"CLS"`

	// 自定义 HTTP 服务的配置信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CustomEndpoint *CustomEndpoint `json:"CustomEndpoint,omitnil,omitempty" name:"CustomEndpoint"`

	// AWS S3 兼容存储桶的配置信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	S3 *S3 `json:"S3,omitnil,omitempty" name:"S3"`

	// 创建时间。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 更新时间。
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`
}

type RenewFlag struct {
	// 预付费套餐的自动续费标志，取值有：
	// <li> on：开启自动续费；</li>
	// <li> off：不开启自动续费。</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`
}

// Predefined struct for user
type RenewPlanRequestParams struct {
	// 套餐 ID，形如 edgeone-2unuvzjmmn2q。
	PlanId *string `json:"PlanId,omitnil,omitempty" name:"PlanId"`

	// 续费套餐的时长，单位：月，取值有：1，2，3，4，5，6，7，8，9，10，11，12，24，36。
	Period *int64 `json:"Period,omitnil,omitempty" name:"Period"`

	// 是否自动使用代金券，取值有：<li> true：是；</li><li> false：否。</li>不填写使用默认值 false。
	AutoUseVoucher *string `json:"AutoUseVoucher,omitnil,omitempty" name:"AutoUseVoucher"`
}

type RenewPlanRequest struct {
	*tchttp.BaseRequest
	
	// 套餐 ID，形如 edgeone-2unuvzjmmn2q。
	PlanId *string `json:"PlanId,omitnil,omitempty" name:"PlanId"`

	// 续费套餐的时长，单位：月，取值有：1，2，3，4，5，6，7，8，9，10，11，12，24，36。
	Period *int64 `json:"Period,omitnil,omitempty" name:"Period"`

	// 是否自动使用代金券，取值有：<li> true：是；</li><li> false：否。</li>不填写使用默认值 false。
	AutoUseVoucher *string `json:"AutoUseVoucher,omitnil,omitempty" name:"AutoUseVoucher"`
}

func (r *RenewPlanRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RenewPlanRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PlanId")
	delete(f, "Period")
	delete(f, "AutoUseVoucher")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RenewPlanRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RenewPlanResponseParams struct {
	// 订单号。
	DealName *string `json:"DealName,omitnil,omitempty" name:"DealName"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RenewPlanResponse struct {
	*tchttp.BaseResponse
	Response *RenewPlanResponseParams `json:"Response"`
}

func (r *RenewPlanResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RenewPlanResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Resource struct {
	// 资源 ID。
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 付费模式，取值有：
	// <li>0：后付费。</li>
	PayMode *int64 `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// 创建时间。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 生效时间。
	EnableTime *string `json:"EnableTime,omitnil,omitempty" name:"EnableTime"`

	// 失效时间。
	ExpireTime *string `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`

	// 套餐状态，取值有：
	// <li>normal：正常；</li>
	// <li>isolated：隔离；</li>
	// <li>destroyed：销毁。</li>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 询价参数。
	Sv []*Sv `json:"Sv,omitnil,omitempty" name:"Sv"`

	// 是否自动续费，取值有：
	// <li>0：默认状态；</li>
	// <li>1：自动续费；</li>
	// <li>2：不自动续费。</li>
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitnil,omitempty" name:"AutoRenewFlag"`

	// 套餐关联资源 ID。
	PlanId *string `json:"PlanId,omitnil,omitempty" name:"PlanId"`

	// 地域，取值有：
	// <li>mainland：国内；</li>
	// <li>overseas：海外。</li>
	// <li>global：全球。</li>
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 资源类型，取值有：
	// <li>plan：套餐类型；</li>
	// <li>pay-as-you-go：后付费类型。</li>
	// <li>value-added：增值服务类型。</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Group *string `json:"Group,omitnil,omitempty" name:"Group"`

	// 当前资源绑定的站点数量。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ZoneNumber *int64 `json:"ZoneNumber,omitnil,omitempty" name:"ZoneNumber"`

	// 资源标记类型，取值有：
	// <li>vodeo：vodeo资源。</li>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

type RewriteAction struct {
	// 功能名称，功能名称填写规范可调用接口 [查询规则引擎的设置参数](https://cloud.tencent.com/document/product/1552/80618) 查看。
	Action *string `json:"Action,omitnil,omitempty" name:"Action"`

	// 参数。
	Parameters []*RuleRewriteActionParams `json:"Parameters,omitnil,omitempty" name:"Parameters"`
}

type Rule struct {
	// 执行功能判断条件。
	// 注意：满足该数组内任意一项条件，功能即可执行。
	Conditions []*RuleAndConditions `json:"Conditions,omitnil,omitempty" name:"Conditions"`

	// 执行的功能。注意：Actions 和 SubRules 不可都为空
	Actions []*Action `json:"Actions,omitnil,omitempty" name:"Actions"`

	// 嵌套规则。注意：SubRules 和 Actions 不可都为空
	SubRules []*SubRuleItem `json:"SubRules,omitnil,omitempty" name:"SubRules"`
}

type RuleAndConditions struct {
	// 规则引擎条件，该数组内所有项全部满足即判断该条件满足。
	Conditions []*RuleCondition `json:"Conditions,omitnil,omitempty" name:"Conditions"`
}

type RuleChoicePropertiesItem struct {
	// 参数名称。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 参数值类型。
	// <li> CHOICE：参数值只能在 ChoicesValue 中选择； </li>
	// <li> TOGGLE：参数值为开关类型，可在 ChoicesValue 中选择；</li>
	// <li> CUSTOM_NUM：参数值用户自定义，整型类型；</li>
	// <li> CUSTOM_STRING：参数值用户自定义，字符串类型。</li>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 参数值的可选值。
	// 注意：若参数值为用户自定义则该数组为空数组。
	ChoicesValue []*string `json:"ChoicesValue,omitnil,omitempty" name:"ChoicesValue"`

	// 数值参数的最小值，非数值参数或 Min 和 Max 值都为 0 则此项无意义。
	Min *int64 `json:"Min,omitnil,omitempty" name:"Min"`

	// 数值参数的最大值，非数值参数或 Min 和 Max 值都为 0 则此项无意义。
	Max *int64 `json:"Max,omitnil,omitempty" name:"Max"`

	// 参数值是否支持多选或者填写多个。
	IsMultiple *bool `json:"IsMultiple,omitnil,omitempty" name:"IsMultiple"`

	// 是否允许为空。
	IsAllowEmpty *bool `json:"IsAllowEmpty,omitnil,omitempty" name:"IsAllowEmpty"`

	// 特殊参数。
	// <li> 为 NULL：RuleAction 选择 NormalAction；</li>
	// <li> 成员参数 Id 为 Action：RuleAction 选择 RewirteAction；</li>
	// <li> 成员参数 Id 为 StatusCode：RuleAction 选择 CodeAction。</li>
	ExtraParameter *RuleExtraParameter `json:"ExtraParameter,omitnil,omitempty" name:"ExtraParameter"`
}

type RuleCodeActionParams struct {
	// 状态 Code。
	StatusCode *int64 `json:"StatusCode,omitnil,omitempty" name:"StatusCode"`

	// 参数名称，参数填写规范可调用接口 [查询规则引擎的设置参数](https://cloud.tencent.com/document/product/1552/80618) 查看。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 参数值。
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`
}

type RuleCondition struct {
	// 运算符，取值有：
	// <li> equal: 等于； </li>
	// <li> notequal: 不等于；</li>
	// <li> exist: 存在； </li>
	// <li> notexist: 不存在。</li>
	Operator *string `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 匹配类型，取值有： <li> filename：文件名； </li> <li> extension：文件后缀； </li> <li> host：HOST； </li> <li> full_url：URL Full，当前站点下完整 URL 路径，必须包含 HTTP 协议，Host 和 路径； </li> <li> url：URL Path，当前站点下 URL 路径的请求； </li><li>client_country：客户端国家/地区；</li> <li> query_string：查询字符串，当前站点下请求URL的查询字符串； </li> <li> request_header：HTTP请求头部。 </li><li> client_ip：客户端 IP。 </li>
	Target *string `json:"Target,omitnil,omitempty" name:"Target"`

	// 对应匹配类型的参数值，仅在匹配类型为查询字符串或HTTP请求头并且运算符取值为存在或不存在时允许传空数组，对应匹配类型有：
	// <li> 文件后缀：jpg、txt等文件后缀；</li>
	// <li> 文件名称：例如 foo.jpg 中的 foo；</li>
	// <li> 全部（站点任意请求）： all； </li>
	// <li> HOST：当前站点下的 host ，例如www.maxx55.com；</li>
	// <li> URL Path：当前站点下 URL 路径的请求，例如：/example；</li>
	// <li> URL Full：当前站点下完整 URL 请求，必须包含 HTTP 协议，Host 和 路径，例如：https://www.maxx55.cn/example；</li>
	// <li> 客户端国家/地区：符合ISO3166标准的国家/地区标识；</li>
	// <li> 查询字符串: 当前站点下URL请求中查询字符串的参数值，例如lang=cn&version=1中的cn和1； </li>
	// <li> HTTP 请求头: HTTP请求头部字段值，例如Accept-Language:zh-CN,zh;q=0.9中的zh-CN,zh;q=0.9。 </li>
	// <li> 客户端 IP: 当前请求携带的客户端请求IP，支持IPv4 IPv6, 支持IP段。 </li>
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`

	// 是否忽略参数值的大小写，默认值为 false。
	IgnoreCase *bool `json:"IgnoreCase,omitnil,omitempty" name:"IgnoreCase"`

	// 对应匹配类型的参数名称，在 Target 值为以下取值时有效，有效时值不能为空：
	// <li> query_string（查询字符串）: 当前站点下URL请求中查询字符串的参数名称，例如lang=cn&version=1中的lang和version； </li>
	// <li> request_header（HTTP 请求头）: HTTP请求头部字段名，例如Accept-Language:zh-CN,zh;q=0.9中的Accept-Language。 </li>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 是否忽略参数名称的大小写，默认值为 false。
	// 注意：此字段可能返回 null，表示取不到有效值。
	//
	// Deprecated: IgnoreNameCase is deprecated.
	IgnoreNameCase *bool `json:"IgnoreNameCase,omitnil,omitempty" name:"IgnoreNameCase"`
}

type RuleExtraParameter struct {
	// 参数名，取值有：
	// <li> Action：修改 HTTP 头部所需参数，RuleAction 选择 RewirteAction；</li>
	// <li> StatusCode：状态码相关功能所需参数，RuleAction 选择 CodeAction。</li>
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 参数值类型。
	// <li> CHOICE：参数值只能在 Values 中选择； </li>
	// <li> CUSTOM_NUM：参数值用户自定义，整型类型；</li>
	// <li> CUSTOM_STRING：参数值用户自定义，字符串类型。</li>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 可选参数值。
	// 注意：当 Id 的值为 StatusCode 时数组中的值为整型，填写参数值时请填写字符串的整型数值。
	Choices []*string `json:"Choices,omitnil,omitempty" name:"Choices"`
}

type RuleItem struct {
	// 规则ID。
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// 规则名称，名称字符串长度 1~255。
	RuleName *string `json:"RuleName,omitnil,omitempty" name:"RuleName"`

	// 规则状态，取值有:
	// <li> enable: 启用； </li>
	// <li> disable: 未启用。 </li>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 规则内容。
	Rules []*Rule `json:"Rules,omitnil,omitempty" name:"Rules"`

	// 规则优先级, 值越大优先级越高，最小为 1。
	RulePriority *int64 `json:"RulePriority,omitnil,omitempty" name:"RulePriority"`

	// 规则标签。
	Tags []*string `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type RuleNormalActionParams struct {
	// 参数名称，参数填写规范可调用接口 [查询规则引擎的设置参数](https://cloud.tencent.com/document/product/1552/80618) 查看。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 参数值。
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`
}

type RuleRewriteActionParams struct {
	// 功能参数名称，参数填写规范可调用接口 [查询规则引擎的设置参数](https://cloud.tencent.com/document/product/1552/80618) 查看。现在只有三种取值：
	// <li> add：添加 HTTP 头部；</li>
	// <li> set：重写 HTTP 头部；</li>
	// <li> del：删除 HTTP 头部。</li>
	Action *string `json:"Action,omitnil,omitempty" name:"Action"`

	// 参数名称。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 参数值。
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`
}

type RulesProperties struct {
	// 值为参数名称。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 数值参数的最小值，非数值参数或 Min 和 Max 值都为 0 则此项无意义。
	Min *int64 `json:"Min,omitnil,omitempty" name:"Min"`

	// 参数值的可选值。
	// 注意：若参数值为用户自定义则该数组为空数组。
	ChoicesValue []*string `json:"ChoicesValue,omitnil,omitempty" name:"ChoicesValue"`

	// 参数值类型。
	// <li> CHOICE：参数值只能在 ChoicesValue 中选择； </li>
	// <li> TOGGLE：参数值为开关类型，可在 ChoicesValue 中选择；</li>
	// <li> OBJECT：参数值为对象类型，ChoiceProperties 为改对象类型关联的属性；</li>
	// <li> CUSTOM_NUM：参数值用户自定义，整型类型；</li>
	// <li> CUSTOM_STRING：参数值用户自定义，字符串类型。</li>注意：当参数类型为 OBJECT 类型时，请注意参考 [示例2 参数为 OBJECT 类型的创建](https://cloud.tencent.com/document/product/1552/80622#.E7.A4.BA.E4.BE.8B2-.E5.8F.82.E6.95.B0.E4.B8.BA-OBJECT-.E7.B1.BB.E5.9E.8B.E7.9A.84.E5.88.9B.E5.BB.BA)
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 数值参数的最大值，非数值参数或 Min 和 Max 值都为 0 则此项无意义。
	Max *int64 `json:"Max,omitnil,omitempty" name:"Max"`

	// 参数值是否支持多选或者填写多个。
	IsMultiple *bool `json:"IsMultiple,omitnil,omitempty" name:"IsMultiple"`

	// 是否允许为空。
	IsAllowEmpty *bool `json:"IsAllowEmpty,omitnil,omitempty" name:"IsAllowEmpty"`

	// 该参数对应的关联配置参数，属于调用接口的必填参数。
	// 注意：如果可选参数无特殊新增参数则该数组为空数组。
	ChoiceProperties []*RuleChoicePropertiesItem `json:"ChoiceProperties,omitnil,omitempty" name:"ChoiceProperties"`

	// <li> 为 NULL：无特殊参数，RuleAction 选择 NormalAction；</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExtraParameter *RuleExtraParameter `json:"ExtraParameter,omitnil,omitempty" name:"ExtraParameter"`
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
	Action *string `json:"Action,omitnil,omitempty" name:"Action"`

	// 参数信息。
	Properties []*RulesProperties `json:"Properties,omitnil,omitempty" name:"Properties"`
}

type S3 struct {
	// 不包含存储桶名称或路径的 URL，例如：`https://storage.googleapis.com`、`https://s3.ap-northeast-2.amazonaws.com`、`https://cos.ap-nanjing.myqcloud.com`。
	Endpoint *string `json:"Endpoint,omitnil,omitempty" name:"Endpoint"`

	// 存储桶所在的地域，例如：`ap-northeast-2`。
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 存储桶名称和日志存储目录，例如：`your_bucket_name/EO-logs/`。如果存储桶中无此目录则会自动创建。
	Bucket *string `json:"Bucket,omitnil,omitempty" name:"Bucket"`

	// 访问存储桶使用的 Access Key ID。
	AccessId *string `json:"AccessId,omitnil,omitempty" name:"AccessId"`

	// 访问存储桶使用的 secret key。
	AccessKey *string `json:"AccessKey,omitnil,omitempty" name:"AccessKey"`

	// 数据压缩类型，取值有: <li> gzip：gzip压缩。</li>不填表示不启用压缩。
	CompressType *string `json:"CompressType,omitnil,omitempty" name:"CompressType"`
}

type SecEntry struct {
	// 查询维度值。
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 查询维度下详细数据。
	Value []*SecEntryValue `json:"Value,omitnil,omitempty" name:"Value"`
}

type SecEntryValue struct {
	// 指标名称。
	Metric *string `json:"Metric,omitnil,omitempty" name:"Metric"`

	// 时序数据详情。
	Detail []*TimingDataItem `json:"Detail,omitnil,omitempty" name:"Detail"`

	// 最大值。
	Max *int64 `json:"Max,omitnil,omitempty" name:"Max"`

	// 平均值。
	Avg *float64 `json:"Avg,omitnil,omitempty" name:"Avg"`

	// 数据总和。
	Sum *float64 `json:"Sum,omitnil,omitempty" name:"Sum"`
}

type SecurityConfig struct {
	// 托管规则。如果入参为空或不填，默认使用历史配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	WafConfig *WafConfig `json:"WafConfig,omitnil,omitempty" name:"WafConfig"`

	// 速率限制。如果入参为空或不填，默认使用历史配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RateLimitConfig *RateLimitConfig `json:"RateLimitConfig,omitnil,omitempty" name:"RateLimitConfig"`

	// 自定义规则。如果入参为空或不填，默认使用历史配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AclConfig *AclConfig `json:"AclConfig,omitnil,omitempty" name:"AclConfig"`

	// Bot配置。如果入参为空或不填，默认使用历史配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	BotConfig *BotConfig `json:"BotConfig,omitnil,omitempty" name:"BotConfig"`

	// 七层防护总开关。如果入参为空或不填，默认使用历史配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SwitchConfig *SwitchConfig `json:"SwitchConfig,omitnil,omitempty" name:"SwitchConfig"`

	// 基础访问管控。如果入参为空或不填，默认使用历史配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	IpTableConfig *IpTableConfig `json:"IpTableConfig,omitnil,omitempty" name:"IpTableConfig"`

	// 例外规则配置。如果入参为空或不填，默认使用历史配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExceptConfig *ExceptConfig `json:"ExceptConfig,omitnil,omitempty" name:"ExceptConfig"`

	// 自定义拦截页面配置。如果入参为空或不填，默认使用历史配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DropPageConfig *DropPageConfig `json:"DropPageConfig,omitnil,omitempty" name:"DropPageConfig"`

	// 模板配置。此处仅出参数使用。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TemplateConfig *TemplateConfig `json:"TemplateConfig,omitnil,omitempty" name:"TemplateConfig"`

	// 慢速攻击配置。如果入参为空或不填，默认使用历史配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SlowPostConfig *SlowPostConfig `json:"SlowPostConfig,omitnil,omitempty" name:"SlowPostConfig"`
}

type SecurityTemplateBinding struct {
	// 模板ID
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// 模板绑定状态。
	TemplateScope []*TemplateScope `json:"TemplateScope,omitnil,omitempty" name:"TemplateScope"`
}

type SecurityType struct {
	// 安全类型开关，取值为：
	// <li> on：开启；</li>
	// <li> off：关闭。</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`
}

type ServerCertInfo struct {
	// 服务器证书 ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CertId *string `json:"CertId,omitnil,omitempty" name:"CertId"`

	// 证书备注名。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Alias *string `json:"Alias,omitnil,omitempty" name:"Alias"`

	// 证书类型，取值有：
	// <li>default：默认证书；</li>
	// <li>upload：用户上传；</li>
	// <li>managed：腾讯云托管。</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 证书过期时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExpireTime *string `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`

	// 证书部署时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeployTime *string `json:"DeployTime,omitnil,omitempty" name:"DeployTime"`

	// 签名算法。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SignAlgo *string `json:"SignAlgo,omitnil,omitempty" name:"SignAlgo"`

	// 证书归属域名名称。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CommonName *string `json:"CommonName,omitnil,omitempty" name:"CommonName"`
}

type SkipCondition struct {
	// 例外跳过类型，取值为：
	// <li>header_fields：HTTP请求Header；</li>
	// <li>cookie：HTTP请求Cookie；</li>
	// <li>query_string：HTTP请求URL中的Query参数；</li>
	// <li>uri：HTTP请求URI；</li>
	// <li>body_raw：HTTP请求Body；</li>
	// <li>body_json： JSON格式的HTTP Body。</li>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 选择跳过的字段，取值为：
	// <li>args：uri 下选择 query 参数: ?name1=jack&age=12；</li>
	// <li>path：uri 下选择部分路径：/path/to/resource.jpg；</li>
	// <li>full：uri 下选择完整路径：example.com/path/to/resource.jpg?name1=jack&age=12；</li>
	// <li>upload_filename：分段文件名，即分段传输文件时；</li>
	// <li>keys：所有的Key；</li>
	// <li>values：匹配Key对应的值；</li>
	// <li>key_value：匹配Key及匹配Value。</li>
	Selector *string `json:"Selector,omitnil,omitempty" name:"Selector"`

	// 匹配Key所使用的匹配方式，取值为：
	// <li>equal：精准匹配，等于；</li>
	// <li>wildcard：通配符匹配，支持 * 通配。</li>
	MatchFromType *string `json:"MatchFromType,omitnil,omitempty" name:"MatchFromType"`

	// 匹配Key的值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MatchFrom []*string `json:"MatchFrom,omitnil,omitempty" name:"MatchFrom"`

	// 匹配Content所使用的匹配方式，取值为：
	// <li>equal：精准匹配，等于；</li>
	// <li>wildcard：通配符匹配，支持 * 通配。</li>
	MatchContentType *string `json:"MatchContentType,omitnil,omitempty" name:"MatchContentType"`

	// 匹配Value的值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MatchContent []*string `json:"MatchContent,omitnil,omitempty" name:"MatchContent"`
}

type SlowPostConfig struct {
	// 开关，取值有：
	// <li>on：开启；</li>
	// <li>off：关闭。</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// 首包配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	FirstPartConfig *FirstPartConfig `json:"FirstPartConfig,omitnil,omitempty" name:"FirstPartConfig"`

	// 基础配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SlowRateConfig *SlowRateConfig `json:"SlowRateConfig,omitnil,omitempty" name:"SlowRateConfig"`

	// 慢速攻击的处置动作，取值有：
	// <li>monitor：观察；</li>
	// <li>drop：拦截。</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Action *string `json:"Action,omitnil,omitempty" name:"Action"`

	// 本规则的Id。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleId *uint64 `json:"RuleId,omitnil,omitempty" name:"RuleId"`
}

type SlowRateConfig struct {
	// 开关，取值有：
	// <li>on：开启；</li>
	// <li>off：关闭。</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// 统计的间隔，单位是秒，即在首段包传输结束后，将数据传输轴按照本参数切分，每个分片独立计算慢速攻击。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Interval *uint64 `json:"Interval,omitnil,omitempty" name:"Interval"`

	// 统计时应用的速率阈值，单位是bps，即如果本分片中的传输速率没达到本参数的值，则判定为慢速攻击，应用慢速攻击的处置方式。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Threshold *uint64 `json:"Threshold,omitnil,omitempty" name:"Threshold"`
}

type SmartRouting struct {
	// 智能加速配置开关，取值有：
	// <li>on：开启；</li>
	// <li>off：关闭。</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`
}

type StandardDebug struct {
	// Debug 功能开关，取值有：
	// <li>on：开启；</li>
	// <li>off：关闭。</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// 允许的客户端来源。支持填写 IPv4 以及 IPv6 的 IP/IP 段。0.0.0.0/0 表示允许所有 IPv4 客户端进行调试，::/0 表示允许所有 IPv6 客户端进行调试。
	AllowClientIPList []*string `json:"AllowClientIPList,omitnil,omitempty" name:"AllowClientIPList"`

	// Debug 功能到期时间。超出设置的时间，则功能失效。
	ExpireTime *string `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`
}

type SubRule struct {
	// 执行功能判断条件。
	// 注意：满足该数组内任意一项条件，功能即可执行。
	Conditions []*RuleAndConditions `json:"Conditions,omitnil,omitempty" name:"Conditions"`

	// 执行的功能。
	Actions []*Action `json:"Actions,omitnil,omitempty" name:"Actions"`
}

type SubRuleItem struct {
	// 嵌套规则信息。
	Rules []*SubRule `json:"Rules,omitnil,omitempty" name:"Rules"`

	// 规则标签。
	Tags []*string `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type Sv struct {
	// 询价参数键。
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 询价参数值。
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`

	// 询价参数映射的配额，取值有：
	// <li>zone：站点数；</li>
	// <li>custom-rule：自定义规则数；</li>
	// <li>rate-limiting-rule：速率限制规则数；</li>
	// <li>l4-proxy-instance：四层代理实例数。</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Pack *string `json:"Pack,omitnil,omitempty" name:"Pack"`

	// 询价参数映射的四层代理实例Id。
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 询价参数对应的防护等级。
	// 取值有： <li> cm_30G：中国大陆加速区域保底防护30Gbps；</li><li> cm_60G：中国大陆加速区域保底防护60Gbps；</li><li> cm_100G：中国大陆加速区域保底防护100Gbps；</li><li> anycast_300G：全球加速区域（除中国大陆）Anycast联防300Gbps；</li><li> anycast_unlimited：全球加速区域（除中国大陆）Anycast无上限全力防护；</li><li> cm_30G_anycast_300G：中国大陆加速区域保底防护30Gbps，全球加速区域（除中国大陆）Anycast联防300Gbps；</li><li> cm_30G_anycast_unlimited：中国大陆加速区域保底防护30Gbps，全球加速区域（除中国大陆）Anycast无上限全力防护；</li><li> cm_60G_anycast_300G：中国大陆加速区域保底防护60Gbps，全球加速区域（除中国大陆）Anycast联防300Gbps；</li><li> cm_60G_anycast_unlimited：中国大陆加速区域保底防护60Gbps，全球加速区域（除中国大陆）Anycast无上限全力防护；</li><li> cm_100G_anycast_300G：中国大陆加速区域保底防护100Gbps，全球加速区域（除中国大陆）Anycast联防300Gbps；</li><li> cm_100G_anycast_unlimited：中国大陆加速区域保底防护100Gbps，全球加速区域（除中国大陆）Anycast无上限全力防护。</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProtectionSpecs *string `json:"ProtectionSpecs,omitnil,omitempty" name:"ProtectionSpecs"`
}

type SwitchConfig struct {
	// Web类型的安全总开关，取值有：
	// <li>on：开启；</li>
	// <li>off：关闭。</li>不影响DDoS与Bot的开关。
	WebSwitch *string `json:"WebSwitch,omitnil,omitempty" name:"WebSwitch"`
}

type Tag struct {
	// 标签键。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagKey *string `json:"TagKey,omitnil,omitempty" name:"TagKey"`

	// 标签值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagValue *string `json:"TagValue,omitnil,omitempty" name:"TagValue"`
}

type Task struct {
	// 任务 ID。
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// 资源。
	Target *string `json:"Target,omitnil,omitempty" name:"Target"`

	// 任务类型。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 节点缓存清除方法，取值有：
	// <li>invalidate：标记过期，用户请求时触发回源校验，即发送带有 If-None-Match 和 If-Modified-Since 头部的 HTTP 条件请求。若源站响应 200，则节点会回源拉取新的资源并更新缓存；若源站响应 304，则节点不会更新缓存；</li>
	// <li>delete：直接删除节点缓存，用户请求时触发回源拉取资源。</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Method *string `json:"Method,omitnil,omitempty" name:"Method"`

	// 状态。取值有：
	// <li>processing：处理中；</li>
	// <li>success：成功；</li>
	// <li> failed：失败；</li>
	// <li>timeout：超时。</li>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 任务创建时间。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 任务完成时间。
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`
}

type TemplateConfig struct {
	// 模板ID。
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// 模板名称。
	TemplateName *string `json:"TemplateName,omitnil,omitempty" name:"TemplateName"`
}

type TemplateScope struct {
	// 站点ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 实例状态列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	EntityStatus []*EntityStatus `json:"EntityStatus,omitnil,omitempty" name:"EntityStatus"`
}

type TimingDataItem struct {
	// 返回数据对应时间点，采用 unix 秒级时间戳。
	Timestamp *int64 `json:"Timestamp,omitnil,omitempty" name:"Timestamp"`

	// 具体数值。
	Value *int64 `json:"Value,omitnil,omitempty" name:"Value"`
}

type TimingDataRecord struct {
	// 查询维度值。
	TypeKey *string `json:"TypeKey,omitnil,omitempty" name:"TypeKey"`

	// 详细时序数据。
	TypeValue []*TimingTypeValue `json:"TypeValue,omitnil,omitempty" name:"TypeValue"`
}

type TimingTypeValue struct {
	// 数据和。
	Sum *int64 `json:"Sum,omitnil,omitempty" name:"Sum"`

	// 最大值。
	Max *int64 `json:"Max,omitnil,omitempty" name:"Max"`

	// 平均值。
	Avg *int64 `json:"Avg,omitnil,omitempty" name:"Avg"`

	// 指标名。
	MetricName *string `json:"MetricName,omitnil,omitempty" name:"MetricName"`

	// 详细数据。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Detail []*TimingDataItem `json:"Detail,omitnil,omitempty" name:"Detail"`
}

type TopDataRecord struct {
	// 查询维度值。
	TypeKey *string `json:"TypeKey,omitnil,omitempty" name:"TypeKey"`

	// top数据排行。
	DetailData []*TopDetailData `json:"DetailData,omitnil,omitempty" name:"DetailData"`
}

type TopDetailData struct {
	// 字段名。
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 字段值。
	Value *int64 `json:"Value,omitnil,omitempty" name:"Value"`
}

type TopEntry struct {
	// top查询维度值。
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 查询具体数据。
	Value []*TopEntryValue `json:"Value,omitnil,omitempty" name:"Value"`
}

type TopEntryValue struct {
	// 排序实体名。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 排序实体数量。
	Count *int64 `json:"Count,omitnil,omitempty" name:"Count"`
}

// Predefined struct for user
type UpgradePlanRequestParams struct {
	// 套餐 ID，形如 edgeone-2unuvzjmmn2q。
	PlanId *string `json:"PlanId,omitnil,omitempty" name:"PlanId"`

	// 需要升级到的目标套餐版本，取值有：<li> basic：基础版套餐；</li><li> standard：标准版套餐。</li>
	PlanType *string `json:"PlanType,omitnil,omitempty" name:"PlanType"`

	// 是否自动使用代金券，取值有：<li> true：是；</li><li> false：否。</li>不填写使用默认值 false。
	AutoUseVoucher *string `json:"AutoUseVoucher,omitnil,omitempty" name:"AutoUseVoucher"`
}

type UpgradePlanRequest struct {
	*tchttp.BaseRequest
	
	// 套餐 ID，形如 edgeone-2unuvzjmmn2q。
	PlanId *string `json:"PlanId,omitnil,omitempty" name:"PlanId"`

	// 需要升级到的目标套餐版本，取值有：<li> basic：基础版套餐；</li><li> standard：标准版套餐。</li>
	PlanType *string `json:"PlanType,omitnil,omitempty" name:"PlanType"`

	// 是否自动使用代金券，取值有：<li> true：是；</li><li> false：否。</li>不填写使用默认值 false。
	AutoUseVoucher *string `json:"AutoUseVoucher,omitnil,omitempty" name:"AutoUseVoucher"`
}

func (r *UpgradePlanRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpgradePlanRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PlanId")
	delete(f, "PlanType")
	delete(f, "AutoUseVoucher")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpgradePlanRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpgradePlanResponseParams struct {
	// 订单号。
	DealName *string `json:"DealName,omitnil,omitempty" name:"DealName"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpgradePlanResponse struct {
	*tchttp.BaseResponse
	Response *UpgradePlanResponseParams `json:"Response"`
}

func (r *UpgradePlanResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpgradePlanResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpstreamHttp2 struct {
	// http2 回源配置开关，取值有：
	// <li>on：开启；</li>
	// <li>off：关闭。</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`
}

type VanityNameServers struct {
	// 自定义 ns 开关，取值有：
	// <li> on：开启；</li>
	// <li> off：关闭。</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// 自定义 ns 列表。
	Servers []*string `json:"Servers,omitnil,omitempty" name:"Servers"`
}

type VanityNameServersIps struct {
	// 自定义名字服务器名称。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 自定义名字服务器 IPv4 地址。
	IPv4 *string `json:"IPv4,omitnil,omitempty" name:"IPv4"`
}

// Predefined struct for user
type VerifyOwnershipRequestParams struct {
	// 站点或者加速域名。
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`
}

type VerifyOwnershipRequest struct {
	*tchttp.BaseRequest
	
	// 站点或者加速域名。
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`
}

func (r *VerifyOwnershipRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *VerifyOwnershipRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "VerifyOwnershipRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type VerifyOwnershipResponseParams struct {
	// 归属权验证结果。
	// <li>success：验证成功；</li>
	// <li>fail：验证失败。</li>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 当验证结果为不通过时，该字段会返回原因，协助您排查问题。
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type VerifyOwnershipResponse struct {
	*tchttp.BaseResponse
	Response *VerifyOwnershipResponseParams `json:"Response"`
}

func (r *VerifyOwnershipResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *VerifyOwnershipResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Waf struct {
	// Waf开关，取值为：
	// <li> on：开启；</li>
	// <li> off：关闭。</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// 策略ID。
	PolicyId *int64 `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`
}

type WafConfig struct {
	// WafConfig开关，取值有：
	// <li> on：开启；</li>
	// <li> off：关闭。</li>开关仅与配置是否生效有关，即使为off（关闭），也可以正常修改配置的内容。
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// 上一次设置的防护级别，取值有：
	// <li> loose：宽松；</li>
	// <li> normal：正常；</li>
	// <li> strict：严格；</li>
	// <li> stricter：超严格；</li>
	// <li> custom：自定义。</li>
	Level *string `json:"Level,omitnil,omitempty" name:"Level"`

	// 全局WAF模式，取值有：
	// <li> block：阻断（全局阻断，但可对详细规则配置观察）；</li>
	// <li> observe：观察（无论详细规则配置什么，都为观察）。</li>
	Mode *string `json:"Mode,omitnil,omitempty" name:"Mode"`

	// 托管规则详细配置。如果为null，默认使用历史配置。
	WafRule *WafRule `json:"WafRule,omitnil,omitempty" name:"WafRule"`

	// AI规则引擎防护配置。如果为null，默认使用历史配置。
	AiRule *AiRule `json:"AiRule,omitnil,omitempty" name:"AiRule"`
}

type WafRule struct {
	// 托管规则开关，取值有：
	// <li> on：开启；</li>
	// <li> off：关闭。</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// 黑名单ID列表，将规则ID加入本参数列表中代表该ID关闭，即该规则ID不再生效。
	BlockRuleIDs []*int64 `json:"BlockRuleIDs,omitnil,omitempty" name:"BlockRuleIDs"`

	// 观察模式ID列表，将规则ID加入本参数列表中代表该ID使用观察模式生效，即该规则ID进入观察模式。
	ObserveRuleIDs []*int64 `json:"ObserveRuleIDs,omitnil,omitempty" name:"ObserveRuleIDs"`
}

type WebSocket struct {
	// WebSocket 超时时间配置开关，取值有：
	// <li>on：使用Timeout作为WebSocket超时时间；</li>
	// <li>off：平台仍支持WebSocket连接，此时使用系统默认的15秒为超时时间。</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// 超时时间，单位为秒，最大超时时间120秒。
	Timeout *int64 `json:"Timeout,omitnil,omitempty" name:"Timeout"`
}

type Zone struct {
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 站点名称。
	ZoneName *string `json:"ZoneName,omitnil,omitempty" name:"ZoneName"`

	// 站点当前使用的 NS 列表。
	OriginalNameServers []*string `json:"OriginalNameServers,omitnil,omitempty" name:"OriginalNameServers"`

	// 腾讯云分配的 NS 列表。
	NameServers []*string `json:"NameServers,omitnil,omitempty" name:"NameServers"`

	// 站点状态，取值有：
	// <li> active：NS 已切换； </li>
	// <li> pending：NS 未切换；</li>
	// <li> moved：NS 已切走；</li>
	// <li> deactivated：被封禁。 </li>
	// <li> initializing：待绑定套餐。 </li>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 站点接入方式，取值有：
	// <li> full：NS 接入；</li>
	// <li> partial：CNAME 接入；</li>
	// <li> noDomainAccess：无域名接入；</li>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 站点是否关闭。
	Paused *bool `json:"Paused,omitnil,omitempty" name:"Paused"`

	// 是否开启 CNAME 加速，取值有：
	// <li> enabled：开启；</li>
	// <li> disabled：关闭。</li>
	CnameSpeedUp *string `json:"CnameSpeedUp,omitnil,omitempty" name:"CnameSpeedUp"`

	// CNAME 接入状态，取值有：
	// <li> finished：站点已验证；</li>
	// <li> pending：站点验证中。</li>
	CnameStatus *string `json:"CnameStatus,omitnil,omitempty" name:"CnameStatus"`

	// 资源标签列表。
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 计费资源列表。
	Resources []*Resource `json:"Resources,omitnil,omitempty" name:"Resources"`

	// 站点创建时间。
	CreatedOn *string `json:"CreatedOn,omitnil,omitempty" name:"CreatedOn"`

	// 站点修改时间。
	ModifiedOn *string `json:"ModifiedOn,omitnil,omitempty" name:"ModifiedOn"`

	// 站点接入地域，取值有：
	// <li> global：全球；</li>
	// <li> mainland：中国大陆；</li>
	// <li> overseas：境外区域。</li>
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 用户自定义 NS 信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	VanityNameServers *VanityNameServers `json:"VanityNameServers,omitnil,omitempty" name:"VanityNameServers"`

	// 用户自定义 NS IP 信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	VanityNameServersIps []*VanityNameServersIps `json:"VanityNameServersIps,omitnil,omitempty" name:"VanityNameServersIps"`

	// 展示状态，取值有：
	// <li> active：已启用；</li>
	// <li> inactive：未生效；</li>
	// <li> paused：已停用。</li>
	ActiveStatus *string `json:"ActiveStatus,omitnil,omitempty" name:"ActiveStatus"`

	// 站点别名。数字、英文、-和_组合，限制20个字符。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AliasZoneName *string `json:"AliasZoneName,omitnil,omitempty" name:"AliasZoneName"`

	// 是否伪站点，取值有：
	// <li> 0：非伪站点；</li>
	// <li> 1：伪站点。</li>
	IsFake *int64 `json:"IsFake,omitnil,omitempty" name:"IsFake"`

	// 锁定状态，取值有：<li> enable：正常，允许进行修改操作；</li><li> disable：锁定中，不允许进行修改操作；</li><li> plan_migrate：套餐迁移中，不允许进行修改操作。</li>
	LockStatus *string `json:"LockStatus,omitnil,omitempty" name:"LockStatus"`

	// 归属权验证信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OwnershipVerification *OwnershipVerification `json:"OwnershipVerification,omitnil,omitempty" name:"OwnershipVerification"`
}

type ZoneSetting struct {
	// 站点名称。
	ZoneName *string `json:"ZoneName,omitnil,omitempty" name:"ZoneName"`

	// 站点加速区域信息，取值有：
	// <li> mainland：中国境内加速；</li>
	// <li> overseas：中国境外加速。</li>
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 节点缓存键配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CacheKey *CacheKey `json:"CacheKey,omitnil,omitempty" name:"CacheKey"`

	// Quic访问配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Quic *Quic `json:"Quic,omitnil,omitempty" name:"Quic"`

	// POST请求传输配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PostMaxSize *PostMaxSize `json:"PostMaxSize,omitnil,omitempty" name:"PostMaxSize"`

	// 智能压缩配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Compression *Compression `json:"Compression,omitnil,omitempty" name:"Compression"`

	// Http2回源配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpstreamHttp2 *UpstreamHttp2 `json:"UpstreamHttp2,omitnil,omitempty" name:"UpstreamHttp2"`

	// 访问协议强制Https跳转配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ForceRedirect *ForceRedirect `json:"ForceRedirect,omitnil,omitempty" name:"ForceRedirect"`

	// 缓存过期时间配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CacheConfig *CacheConfig `json:"CacheConfig,omitnil,omitempty" name:"CacheConfig"`

	// 源站配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Origin *Origin `json:"Origin,omitnil,omitempty" name:"Origin"`

	// 智能加速配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SmartRouting *SmartRouting `json:"SmartRouting,omitnil,omitempty" name:"SmartRouting"`

	// 浏览器缓存配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxAge *MaxAge `json:"MaxAge,omitnil,omitempty" name:"MaxAge"`

	// 离线缓存配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OfflineCache *OfflineCache `json:"OfflineCache,omitnil,omitempty" name:"OfflineCache"`

	// WebSocket配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	WebSocket *WebSocket `json:"WebSocket,omitnil,omitempty" name:"WebSocket"`

	// 客户端IP回源请求头配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClientIpHeader *ClientIpHeader `json:"ClientIpHeader,omitnil,omitempty" name:"ClientIpHeader"`

	// 缓存预刷新配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CachePrefresh *CachePrefresh `json:"CachePrefresh,omitnil,omitempty" name:"CachePrefresh"`

	// Ipv6访问配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Ipv6 *Ipv6 `json:"Ipv6,omitnil,omitempty" name:"Ipv6"`

	// Https 加速配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Https *Https `json:"Https,omitnil,omitempty" name:"Https"`

	// 回源时是否携带客户端IP所属地域信息的配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClientIpCountry *ClientIpCountry `json:"ClientIpCountry,omitnil,omitempty" name:"ClientIpCountry"`

	// Grpc协议支持配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Grpc *Grpc `json:"Grpc,omitnil,omitempty" name:"Grpc"`

	// 图片优化相关配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ImageOptimize *ImageOptimize `json:"ImageOptimize,omitnil,omitempty" name:"ImageOptimize"`

	// 中国大陆加速优化配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AccelerateMainland *AccelerateMainland `json:"AccelerateMainland,omitnil,omitempty" name:"AccelerateMainland"`

	// 标准 Debug 配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	StandardDebug *StandardDebug `json:"StandardDebug,omitnil,omitempty" name:"StandardDebug"`
}