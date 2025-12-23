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

package v20220901

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

type APIResource struct {
	// 资源 ID。
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 资源名称。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// API 资源关联的 API 服务 ID 列表。
	APIServiceIds []*string `json:"APIServiceIds,omitnil,omitempty" name:"APIServiceIds"`

	// 资源路径。
	Path *string `json:"Path,omitnil,omitempty" name:"Path"`

	// 请求方法列表。支持以下取值：GET, POST, PUT, HEAD, PATCH, OPTIONS, DELETE。
	Methods []*string `json:"Methods,omitnil,omitempty" name:"Methods"`

	// 请求内容匹配规则的具体内容，需符合表达式语法，详细规范参见产品文档。
	RequestConstraint *string `json:"RequestConstraint,omitnil,omitempty" name:"RequestConstraint"`
}

type APIService struct {
	// API 服务 ID。
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// API 服务名称。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 基础路径。
	BasePath *string `json:"BasePath,omitnil,omitempty" name:"BasePath"`
}

type AccelerateMainland struct {
	// 是否开启中国大陆加速优化配置，取值有：
	// <li>on：开启；</li>
	// <li>off：关闭。</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`
}

type AccelerateMainlandParameters struct {
	// 中国大陆加速优化配置开关，取值有：
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
	// <li>init：未生效，待激活站点。</li>
	DomainStatus *string `json:"DomainStatus,omitnil,omitempty" name:"DomainStatus"`

	// CNAME 地址。
	Cname *string `json:"Cname,omitnil,omitempty" name:"Cname"`

	// IPv6 状态，取值有：
	// <li>follow：遵循站点IPv6配置；</li>
	// <li>on：开启状态；</li>
	// <li>off：关闭状态。</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	IPv6Status *string `json:"IPv6Status,omitnil,omitempty" name:"IPv6Status"`

	// 加速域名归属权验证状态，取值有： 
	// <li>pending：待验证；</li>
	// <li>finished：已完成验证。</li>	
	// 注意：此字段可能返回 null，表示取不到有效值。
	IdentificationStatus *string `json:"IdentificationStatus,omitnil,omitempty" name:"IdentificationStatus"`

	// 加速域名需进行归属权验证才能继续提供服务时，该对象会携带对应验证方式所需要的信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OwnershipVerification *OwnershipVerification `json:"OwnershipVerification,omitnil,omitempty" name:"OwnershipVerification"`

	// 源站信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OriginDetail *OriginDetail `json:"OriginDetail,omitnil,omitempty" name:"OriginDetail"`

	// 回源协议，取值有：
	// <li>FOLLOW：协议跟随；</li>
	// <li>HTTP：HTTP协议回源；</li>
	// <li>HTTPS：HTTPS协议回源。</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	OriginProtocol *string `json:"OriginProtocol,omitnil,omitempty" name:"OriginProtocol"`

	// HTTP 回源端口。
	// 注意：此字段可能返回 null，表示取不到有效值。
	HttpOriginPort *uint64 `json:"HttpOriginPort,omitnil,omitempty" name:"HttpOriginPort"`

	// HTTPS 回源端口。
	// 注意：此字段可能返回 null，表示取不到有效值。
	HttpsOriginPort *uint64 `json:"HttpsOriginPort,omitnil,omitempty" name:"HttpsOriginPort"`

	// 加速域名证书信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Certificate *AccelerationDomainCertificate `json:"Certificate,omitnil,omitempty" name:"Certificate"`

	// 创建时间。
	CreatedOn *string `json:"CreatedOn,omitnil,omitempty" name:"CreatedOn"`

	// 修改时间。
	ModifiedOn *string `json:"ModifiedOn,omitnil,omitempty" name:"ModifiedOn"`
}

type AccelerationDomainCertificate struct {
	// 配置服务端证书的模式，取值有： <ul><li>disable：不配置服务端证书；</li> <li>eofreecert：通过自动验证申请免费证书并部署。验证方式详见：[申请免费证书支持的验证方式](https://cloud.tencent.com/document/product/1552/90437) - 在 NS 或者 DNSPod 托管接入模式下，仅支持自动验证的方式申请免费证书。 - 当免费证书申请失败时会导致证书部署失败，您可以通过<a href = 'https://cloud.tencent.com/document/product/1552/124806'>检查免费证书申请结果</a>接口获取申请失败原因。</li><li>eofreecert_manual：部署 DNS 委派验证或者文件验证申请的免费证书。在部署免费证书前，您需要触发<a href = 'https://cloud.tencent.com/document/product/1552/124807'>申请免费证书</a>接口申请免费证书。在免费证书申请成功后，你可以通过该枚举值对免费证书进行部署；</li> <ul><li>注意：在对免费证书部署时，需要保证当前已存在申请成功的免费证书。您可以通过<a href = 'https://cloud.tencent.com/document/product/1552/124806'>检查免费证书申请结果</a>接口检查当前是否已存在申请成功的免费证书。</li> </ul> <li>sslcert：配置 SSL 托管服务端证书。</li></ul>
	Mode *string `json:"Mode,omitnil,omitempty" name:"Mode"`

	// 服务端证书列表，相关证书部署在 EO 的入口侧。
	// 注意：此字段可能返回 null，表示取不到有效值。
	List []*CertificateInfo `json:"List,omitnil,omitempty" name:"List"`

	// 在边缘双向认证场景下，该字段为客户端的 CA 证书，部署在 EO 节点内，用于 EO 节点认证客户端证书。
	ClientCertInfo *MutualTLS `json:"ClientCertInfo,omitnil,omitempty" name:"ClientCertInfo"`

	// 用于分别开启/关闭回源双向认证和源站证书校验。回源双向认证的证书用于 EO 回源时携带，源站可选择校验该证书用于确保请求来源于受信任的 EO 节点。源站证书校验开启时，证书配置用于 EO 节点校验源站证书是否可信。
	UpstreamCertInfo *UpstreamCertInfo `json:"UpstreamCertInfo,omitnil,omitempty" name:"UpstreamCertInfo"`
}

type AccessURLRedirectParameters struct {
	// 状态码，取值为 301、302、303、307、308 之一。
	StatusCode *int64 `json:"StatusCode,omitnil,omitempty" name:"StatusCode"`

	// 目标请求协议，取值有：
	// <li>http：目标请求协议 HTTP；</li>
	// <li>https：目标请求协议 HTTPS；</li>
	// <li>follow：跟随请求。</li>
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// 目标 HostName 。
	// 注意：此字段可能返回 null，表示取不到有效值。
	HostName *HostName `json:"HostName,omitnil,omitempty" name:"HostName"`

	// 目标路径。
	// 注意：此字段可能返回 null，表示取不到有效值。
	URLPath *URLPath `json:"URLPath,omitnil,omitempty" name:"URLPath"`

	// 携带查询参数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	QueryString *AccessURLRedirectQueryString `json:"QueryString,omitnil,omitempty" name:"QueryString"`
}

type AccessURLRedirectQueryString struct {
	// 执行动作，取值有：
	// <li>full：全部保留；</li>
	// <li>ignore：全部忽略。</li>
	Action *string `json:"Action,omitnil,omitempty" name:"Action"`
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
	// <li> 单连接下载限速（ResponseSpeedLimit）；</li>
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

type AdaptiveFrequencyControl struct {
	// 自适应频控是否开启。取值有：<li>on：开启；</li><li>off：关闭。</li>
	Enabled *string `json:"Enabled,omitnil,omitempty" name:"Enabled"`

	// 自适应频控的限制等级，当 Enabled 为 on 时，此字段必填。取值有：<li>Loose：宽松；</li><li>Moderate：适中；</li><li>Strict：严格。</li>
	Sensitivity *string `json:"Sensitivity,omitnil,omitempty" name:"Sensitivity"`

	// 自适应频控的处置方式，当 Enabled 为 on 时，此字段必填。SecurityAction 的 Name 取值支持：<li>Monitor：观察；</li><li>Deny：拦截；</li><li>Challenge：挑战，其中ChallengeActionParameters.Name仅支持JSChallenge。</li>
	Action *SecurityAction `json:"Action,omitnil,omitempty" name:"Action"`
}

type Addresses struct {
	// IPv4 网段列表。
	IPv4 []*string `json:"IPv4,omitnil,omitempty" name:"IPv4"`

	// IPv6 网段列表。
	IPv6 []*string `json:"IPv6,omitnil,omitempty" name:"IPv6"`
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

type AllowActionParameters struct {
	// 最小延迟响应时间，当配置为 0s 时，表示不延迟直接响应。支持的单位有：<li>s：秒，取值范围 0～5。</li>
	MinDelayTime *string `json:"MinDelayTime,omitnil,omitempty" name:"MinDelayTime"`

	// 最大延迟响应时间，支持的单位有：<li>s：秒，取值范围 5～10。</li>
	MaxDelayTime *string `json:"MaxDelayTime,omitnil,omitempty" name:"MaxDelayTime"`
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
	SessionPersistTime *uint64 `json:"SessionPersistTime,omitnil,omitempty" name:"SessionPersistTime"`

	// 源站端口，支持格式：
	// <li>单端口，如：80。</li>
	// <li>端口段：81-82，表示81，82两个端口。</li>
	OriginPort *string `json:"OriginPort,omitnil,omitempty" name:"OriginPort"`

	// 规则标签。
	RuleTag *string `json:"RuleTag,omitnil,omitempty" name:"RuleTag"`
}

// Predefined struct for user
type ApplyFreeCertificateRequestParams struct {
	// 站点ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 申请免费证书的目标域名。
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 申请免费证书的验证方式，详细验证方式说明参考[免费证书申请方式说明文档](https://cloud.tencent.com/document/product/1552/90437) ，相关取值有：
	// <li>http_challenge：HTTP 访问文件验证方式，通过 HTTP 访问域名指定 URL 获取文件信息以完成免费证书申请验证；</li>
	// <li>dns_challenge：DNS 委派验证方式，通过添加指定的主机记录解析指向 EdgeOne 以完成免费证书申请验证。</li>
	// 注意：在触发本接口后，你需要根据返回的验证信息，完成验证内容配置。配置完成后，还需要通过<a href = 'https://cloud.tencent.com/document/product/1552/124806'>检查免费证书申请结果</a>接口进行验证，验证通过后，即可申请成功。在免费证书申请成功后，你可以调用<a href = 'https://cloud.tencent.com/document/product/1552/80764'>配置域名证书</a>接口为当前域名部署免费证书。
	VerificationMethod *string `json:"VerificationMethod,omitnil,omitempty" name:"VerificationMethod"`
}

type ApplyFreeCertificateRequest struct {
	*tchttp.BaseRequest
	
	// 站点ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 申请免费证书的目标域名。
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 申请免费证书的验证方式，详细验证方式说明参考[免费证书申请方式说明文档](https://cloud.tencent.com/document/product/1552/90437) ，相关取值有：
	// <li>http_challenge：HTTP 访问文件验证方式，通过 HTTP 访问域名指定 URL 获取文件信息以完成免费证书申请验证；</li>
	// <li>dns_challenge：DNS 委派验证方式，通过添加指定的主机记录解析指向 EdgeOne 以完成免费证书申请验证。</li>
	// 注意：在触发本接口后，你需要根据返回的验证信息，完成验证内容配置。配置完成后，还需要通过<a href = 'https://cloud.tencent.com/document/product/1552/124806'>检查免费证书申请结果</a>接口进行验证，验证通过后，即可申请成功。在免费证书申请成功后，你可以调用<a href = 'https://cloud.tencent.com/document/product/1552/80764'>配置域名证书</a>接口为当前域名部署免费证书。
	VerificationMethod *string `json:"VerificationMethod,omitnil,omitempty" name:"VerificationMethod"`
}

func (r *ApplyFreeCertificateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ApplyFreeCertificateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "Domain")
	delete(f, "VerificationMethod")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ApplyFreeCertificateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ApplyFreeCertificateResponseParams struct {
	// 当 VerificationMethod 为 dns_challenge 时，域名申请免费证书的相关验证信息。
	DnsVerification *DnsVerification `json:"DnsVerification,omitnil,omitempty" name:"DnsVerification"`

	// 当 VerificationMethod 为 http_challenge 时，域名申请免费证书的相关验证信息。
	FileVerification *FileVerification `json:"FileVerification,omitnil,omitempty" name:"FileVerification"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ApplyFreeCertificateResponse struct {
	*tchttp.BaseResponse
	Response *ApplyFreeCertificateResponseParams `json:"Response"`
}

func (r *ApplyFreeCertificateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ApplyFreeCertificateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AscriptionInfo struct {
	// 主机记录。
	Subdomain *string `json:"Subdomain,omitnil,omitempty" name:"Subdomain"`

	// 记录类型。
	RecordType *string `json:"RecordType,omitnil,omitempty" name:"RecordType"`

	// 记录值。
	RecordValue *string `json:"RecordValue,omitnil,omitempty" name:"RecordValue"`
}

type AudioTemplateInfo struct {
	// 音频流的编码格式。可选值为：
	// <li>libfdk_aac。</li>
	Codec *string `json:"Codec,omitnil,omitempty" name:"Codec"`

	// 音频通道数，可选值：<li>2：双通道。</li>默认值：2。
	AudioChannel *uint64 `json:"AudioChannel,omitnil,omitempty" name:"AudioChannel"`
}

type AuthenticationParameters struct {
	// 鉴权类型。取值有：
	// 
	// <li>TypeA：鉴权方式 A 类型，具体含义请参考 [鉴权方式 A](https://cloud.tencent.com/document/product/1552/109329)；</li>
	// <li>TypeB：鉴权方式 B 类型，具体含义请参考 [鉴权方式 B](https://cloud.tencent.com/document/product/1552/109330)；</li>
	// <li>TypeC：鉴权方式 C 类型，具体含义请参考 [鉴权方式 C](https://cloud.tencent.com/document/product/1552/109331)；</li>
	// <li>TypeD：鉴权方式 D 类型，具体含义请参考 [鉴权方式 D](https://cloud.tencent.com/document/product/1552/109332)；</li>
	// <li>TypeVOD：鉴权方式 V 类型，具体含义请参考 [鉴权方式 V](https://cloud.tencent.com/document/product/1552/109333)。</li>
	AuthType *string `json:"AuthType,omitnil,omitempty" name:"AuthType"`

	// 主鉴权密钥，由 6～40 位大小写英文字母或数字组成，不能包含 " 和 $。
	SecretKey *string `json:"SecretKey,omitnil,omitempty" name:"SecretKey"`

	// 鉴权 URL 的有效时长，单位为秒，取值：1～630720000。用于判断客户端访问请求是否过期：
	// <li>若当前时间超过 “timestamp + 有效时长” 时间，则为过期请求，直接返回 403。</li>
	// <li>若当前时间未超过 “timestamp + 有效时长” 时间，则请求未过期，继续校验 md5 字符串。</li>注意：当 AuthType 为 TypeA、TypeB、TypeC、TypeD 之一时，此字段必填。
	Timeout *int64 `json:"Timeout,omitnil,omitempty" name:"Timeout"`

	// 备鉴权密钥，由 6～40 位大小写英文字母或数字组成，不能包含 " 和 $。
	BackupSecretKey *string `json:"BackupSecretKey,omitnil,omitempty" name:"BackupSecretKey"`

	// 鉴权参数名称，节点将校验此参数名对应的值。由 1～100 位大小写字母、数字或下划线组成。<br>注意：当 AuthType 为 TypeA、TypeD 之一时，此字段必填。
	AuthParam *string `json:"AuthParam,omitnil,omitempty" name:"AuthParam"`

	// 鉴权时间戳，和 AuthParam 字段的值不能相同。<br>注意：当 AuthType 为 TypeD 时，此字段必填。
	TimeParam *string `json:"TimeParam,omitnil,omitempty" name:"TimeParam"`

	// 鉴权时间格式，取值有：
	// <li>dec：十进制；</li>
	// <li>hex：十六进制。</li>注意：当 AuthType 为 TypeD 时，此字段必填。默认为 hex。
	TimeFormat *string `json:"TimeFormat,omitnil,omitempty" name:"TimeFormat"`
}

type BandwidthAbuseDefense struct {
	// 流量防盗刷（仅适用中国大陆地区）是否开启。取值有：<li>on：开启；</li><li>off：关闭。</li>
	Enabled *string `json:"Enabled,omitnil,omitempty" name:"Enabled"`

	// 流量防盗刷（仅适用中国大陆地区）的处置方式，当 Enabled 为 on 时，此字段必填。SecurityAction 的 Name 取值支持：<li>Monitor：观察；</li><li>Deny：拦截；</li><li>Challenge：挑战，其中ChallengeActionParameters.Name仅支持JSChallenge。</li>
	Action *SecurityAction `json:"Action,omitnil,omitempty" name:"Action"`
}

type BasicBotSettings struct {
	// 客户端 IP 的来源 IDC 配置，用于处置来自 IDC（数据中心） 的客户端 IP 的访问请求。此类来源请求不是由移动端或浏览器端直接访问。
	SourceIDC *SourceIDC `json:"SourceIDC,omitnil,omitempty" name:"SourceIDC"`

	// 搜索引擎爬虫配置，用于处置来自搜索引擎爬虫的请求。此类请求的 IP、User-Agent 或 rDNS 结果匹配已知搜索引擎爬虫。
	SearchEngineBots *SearchEngineBots `json:"SearchEngineBots,omitnil,omitempty" name:"SearchEngineBots"`

	// 商业或开源工具 UA 特征配置（原 UA 特征规则），用于处置来自已知商业工具或开源工具的访问请求。此类请求的 User-Agent 头部符合已知商业或开源工具特征。
	KnownBotCategories *KnownBotCategories `json:"KnownBotCategories,omitnil,omitempty" name:"KnownBotCategories"`

	// IP 威胁情报库（原客户端画像分析）配置，用于处置近期访问行为具有特定风险特征的客户端 IP。
	IPReputation *IPReputation `json:"IPReputation,omitnil,omitempty" name:"IPReputation"`

	// Bot 智能分析的具体配置。
	BotIntelligence *BotIntelligence `json:"BotIntelligence,omitnil,omitempty" name:"BotIntelligence"`
}

type BillingData struct {
	// 数据时间戳。
	Time *string `json:"Time,omitnil,omitempty" name:"Time"`

	// 数值。
	Value *uint64 `json:"Value,omitnil,omitempty" name:"Value"`

	// 数据点所属站点 ID。若使用内容标识符功能，则该项值为内容标识符。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 数据点所属域名。
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`

	// 数据点所属四层代理实例 ID。
	ProxyId *string `json:"ProxyId,omitnil,omitempty" name:"ProxyId"`

	// 数据点所属计费大区 ID。计费大区以实际服务用户客户端的 EdgeOne 节点所在区域为准。取值有：<li>CH：中国大陆境内</li><li>AF：非洲</li><li>AS1：亚太一区</li><li>AS2：亚太二区</li><li>AS3：亚太三区</li><li>EU：欧洲</li><li>MidEast：中东</li><li>NA：北美</li><li> SA：南美</li>
	RegionId *string `json:"RegionId,omitnil,omitempty" name:"RegionId"`
}

type BillingDataFilter struct {
	// 参数名称。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 参数值。
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type BindDomainInfo struct {
	// 域名。
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 域名所属的站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 绑定状态，取值有: 
	// <li>process：绑定中；</li>
	// <li>online：绑定成功；</li>
	// <li>fail：绑定失败。</li>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`
}

// Predefined struct for user
type BindSecurityTemplateToEntityRequestParams struct {
	// 需要绑定或解绑的策略模板所属站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 绑定至策略模板（或者从策略模板解绑）的域名列表。
	Entities []*string `json:"Entities,omitnil,omitempty" name:"Entities"`

	// 绑定或解绑操作选项，取值有：
	// <li>bind：绑定域名至策略模板。</li>
	// <li>unbind-keep-policy：将域名从策略模板解绑，解绑时保留当前策略。</li>
	// <li>unbind-use-default：将域名从策略模板解绑，并使用默认空白策略。</li>注意：解绑操作当前仅支持单个域名解绑。即：当 Operate 参数取值为 unbind-keep-policy 或 unbind-use-default 时，Entities 参数列表仅支持填写一个域名。
	Operate *string `json:"Operate,omitnil,omitempty" name:"Operate"`

	// 指定绑定或解绑的策略模板 ID 或站点全局策略
	// <li>如需绑定至策略模板，或从策略模板解绑，请指定策略模板 ID。</li>
	// <li>如需绑定至站点全局策略，或从站点全局策略解绑，请使用 @ZoneLevel@domain 参数值。</li>
	// 
	// 注意：解绑后，域名将使用独立策略，并单独计算规则配额，请确保解绑前套餐规则配额充足。
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// 传入域名如果已经绑定了策略模板（含站点级防护策略），通过设置该参数表示是否替换该模板，默认值为 true。支持下列取值：<li>true： 替换域名当前绑定的模板。</li><li>false：不替换域名当前绑定的模板。</li>注意：当设置为 false 时，若传入域名已经绑定策略模板，API  将返回错误；站点级防护策略也为一种策略模板。
	OverWrite *bool `json:"OverWrite,omitnil,omitempty" name:"OverWrite"`
}

type BindSecurityTemplateToEntityRequest struct {
	*tchttp.BaseRequest
	
	// 需要绑定或解绑的策略模板所属站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 绑定至策略模板（或者从策略模板解绑）的域名列表。
	Entities []*string `json:"Entities,omitnil,omitempty" name:"Entities"`

	// 绑定或解绑操作选项，取值有：
	// <li>bind：绑定域名至策略模板。</li>
	// <li>unbind-keep-policy：将域名从策略模板解绑，解绑时保留当前策略。</li>
	// <li>unbind-use-default：将域名从策略模板解绑，并使用默认空白策略。</li>注意：解绑操作当前仅支持单个域名解绑。即：当 Operate 参数取值为 unbind-keep-policy 或 unbind-use-default 时，Entities 参数列表仅支持填写一个域名。
	Operate *string `json:"Operate,omitnil,omitempty" name:"Operate"`

	// 指定绑定或解绑的策略模板 ID 或站点全局策略
	// <li>如需绑定至策略模板，或从策略模板解绑，请指定策略模板 ID。</li>
	// <li>如需绑定至站点全局策略，或从站点全局策略解绑，请使用 @ZoneLevel@domain 参数值。</li>
	// 
	// 注意：解绑后，域名将使用独立策略，并单独计算规则配额，请确保解绑前套餐规则配额充足。
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// 传入域名如果已经绑定了策略模板（含站点级防护策略），通过设置该参数表示是否替换该模板，默认值为 true。支持下列取值：<li>true： 替换域名当前绑定的模板。</li><li>false：不替换域名当前绑定的模板。</li>注意：当设置为 false 时，若传入域名已经绑定策略模板，API  将返回错误；站点级防护策略也为一种策略模板。
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

type BlockIPActionParameters struct {
	// 封禁 IP 的惩罚时长。支持的单位有：<li>s：秒，取值范围1～120；</li><li>m：分，取值范围1～120；</li><li>h：小时，取值范围1～48。</li>
	Duration *string `json:"Duration,omitnil,omitempty" name:"Duration"`
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
	IntelligenceRule *IntelligenceRule `json:"IntelligenceRule,omitnil,omitempty" name:"IntelligenceRule"`

	// Bot自定义规则。如果为null，默认使用历史配置。
	BotUserRules []*BotUserRule `json:"BotUserRules,omitnil,omitempty" name:"BotUserRules"`

	// Bot主动特征识别规则。
	AlgDetectRule []*AlgDetectRule `json:"AlgDetectRule,omitnil,omitempty" name:"AlgDetectRule"`

	// Bot托管定制策略，入参可不填，仅出参使用。
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
	Percent *uint64 `json:"Percent,omitnil,omitempty" name:"Percent"`
}

type BotIntelligence struct {
	// 基于客户端和请求特征，将请求来源分为人类来源请求、合法 Bot 请求、疑似 Bot 请求和高风险 Bot 请求，并提供请求处置选项。
	BotRatings *BotRatings `json:"BotRatings,omitnil,omitempty" name:"BotRatings"`

	// Bot 智能分析的具体配置开关。取值有：
	// 
	// on：开启；
	// off：关闭。
	Enabled *string `json:"Enabled,omitnil,omitempty" name:"Enabled"`
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
	TransManagedIds []*int64 `json:"TransManagedIds,omitnil,omitempty" name:"TransManagedIds"`

	// JS挑战的规则ID。默认所有规则不配置JS挑战。
	AlgManagedIds []*int64 `json:"AlgManagedIds,omitnil,omitempty" name:"AlgManagedIds"`

	// 数字验证码的规则ID。默认所有规则不配置数字验证码。
	CapManagedIds []*int64 `json:"CapManagedIds,omitnil,omitempty" name:"CapManagedIds"`

	// 观察的规则ID。默认所有规则不配置观察。
	MonManagedIds []*int64 `json:"MonManagedIds,omitnil,omitempty" name:"MonManagedIds"`

	// 拦截的规则ID。默认所有规则不配置拦截。
	DropManagedIds []*int64 `json:"DropManagedIds,omitnil,omitempty" name:"DropManagedIds"`
}

type BotManagement struct {
	// Bot 管理是否开启。取值有：<li>on：开启；</li><li>off：关闭。</li>
	Enabled *string `json:"Enabled,omitnil,omitempty" name:"Enabled"`

	// Bot 管理的自定义规则，组合各类爬虫和请求行为特征，精准定义 Bot 并配置定制化处置方式。
	CustomRules *BotManagementCustomRules `json:"CustomRules,omitnil,omitempty" name:"CustomRules"`

	// Bot 管理的基础配置，对策略关联的所有域名生效。可以通过 CustomRules 进行精细化定制。
	BasicBotSettings *BasicBotSettings `json:"BasicBotSettings,omitnil,omitempty" name:"BasicBotSettings"`

	// 客户端认证规则的定义列表。该功能内测中，如需使用，请提工单。
	ClientAttestationRules *ClientAttestationRules `json:"ClientAttestationRules,omitnil,omitempty" name:"ClientAttestationRules"`

	// 配置浏览器伪造识别规则（原主动特征识别规则）。设置注入 JavaScript 的响应页面范围，浏览器校验选项，以及对非浏览器客户端的处置方式。
	BrowserImpersonationDetection *BrowserImpersonationDetection `json:"BrowserImpersonationDetection,omitnil,omitempty" name:"BrowserImpersonationDetection"`
}

type BotManagementActionOverrides struct {
	// Bot 规则组下的具体项，用于改写此单条规则项配置的内容，Ids 所对应的具体信息请参考 DescribeBotManagedRules 接口返回的信息。
	Ids []*string `json:"Ids,omitnil,omitempty" name:"Ids"`

	// Ids 中指定 Bot 规则项的处置动作。 SecurityAction 的 Name 取值支持：<li>Deny：拦截；</li><li>Monitor：观察；</li><li>Disabled：未启用，不启用指定规则；</li><li>Challenge：挑战，其中 ChallengeActionParameters 中的 ChallengeOption 支持 JSChallenge 和 ManagedChallenge；</li><li>Allow：放行（仅限Bot基础特征管理）。</li>
	Action *SecurityAction `json:"Action,omitnil,omitempty" name:"Action"`
}

type BotManagementCustomRule struct {
	// Bot 自定义规则的 ID。<br>通过规则 ID 可支持不同的规则配置操作：<br> <li> <b>增加</b>新规则：ID 为空或不指定 ID 参数；</li><li><b>修改</b>已有规则：指定需要更新/修改的规则 ID；</li><li><b>删除</b>已有规则：BotManagementCustomRules 参数中，Rules 列表中未包含的已有规则将被删除。</li>
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// Bot 自定义规则的名称。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Bot 自定义规则是否开启。取值有：<li>on：开启；</li><li>off：关闭。</li>
	Enabled *string `json:"Enabled,omitnil,omitempty" name:"Enabled"`

	// Bot 自定义规则的优先级，范围是 1 ~ 100，默认为 50。
	Priority *int64 `json:"Priority,omitnil,omitempty" name:"Priority"`

	// Bot 自定义规则的具体内容，需符合表达式语法，详细规范参见产品文档。
	Condition *string `json:"Condition,omitnil,omitempty" name:"Condition"`

	// Bot 自定义规则的处置方式。取值有：<li>Monitor：观察；</li><li>Deny：拦截，其中 DenyActionParameters.Name 支持 Deny 和 ReturnCustomPage；</li><li>Challenge：挑战，其中 ChallengeActionParameters.Name 支持 JSChallenge 和 ManagedChallenge；</li><li>Redirect：重定向至 URL。</li>
	Action []*SecurityWeightedAction `json:"Action,omitnil,omitempty" name:"Action"`
}

type BotManagementCustomRules struct {
	// Bot 自定义规则的列表。使用 ModifySecurityPolicy 修改 Web 防护配置时： <br> <li>  若未指定 SecurityPolicy.BotManagement.CustomRules 中的 Rules 参数，或 Rules 参数长度为零：清空所有 Bot 自定义规则配置。</li> <li> 若 SecurityPolicy.BotManagement 参数中，未指定 CustomRules 参数值：保持已有 Bot 自定义规则配置，不做修改。</li>
	Rules []*BotManagementCustomRule `json:"Rules,omitnil,omitempty" name:"Rules"`
}

type BotPortraitRule struct {
	// 本功能的开关，取值有：
	// <li>on：开启；</li>
	// <li>off：关闭。</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// 本规则的ID。仅出参使用。
	RuleID *int64 `json:"RuleID,omitnil,omitempty" name:"RuleID"`

	// JS挑战的规则ID。默认所有规则不配置JS挑战。
	AlgManagedIds []*int64 `json:"AlgManagedIds,omitnil,omitempty" name:"AlgManagedIds"`

	// 数字验证码的规则ID。默认所有规则不配置数字验证码。
	CapManagedIds []*int64 `json:"CapManagedIds,omitnil,omitempty" name:"CapManagedIds"`

	// 观察的规则ID。默认所有规则不配置观察。
	MonManagedIds []*int64 `json:"MonManagedIds,omitnil,omitempty" name:"MonManagedIds"`

	// 拦截的规则ID。默认所有规则不配置拦截。
	DropManagedIds []*int64 `json:"DropManagedIds,omitnil,omitempty" name:"DropManagedIds"`
}

type BotRatings struct {
	// 恶意 Bot 请求的执行动作。 SecurityAction 的 Name 取值支持：<li>Deny：拦截；</li><li>Monitor：观察；</li><li>Allow：放行；</li><li>Challenge：挑战，其中 ChallengeActionParameters 中的 ChallengeOption 支持 JSChallenge 和 ManagedChallenge。</li>
	HighRiskBotRequestsAction *SecurityAction `json:"HighRiskBotRequestsAction,omitnil,omitempty" name:"HighRiskBotRequestsAction"`

	// 疑似 Bot 请求的执行动作。 SecurityAction 的 Name 取值支持：<li>Deny：拦截；</li><li>Monitor：观察；</li><li>Allow：放行；</li><li>Challenge：挑战，其中 ChallengeActionParameters 中的 ChallengeOption 支持 JSChallenge 和 ManagedChallenge。</li>
	LikelyBotRequestsAction *SecurityAction `json:"LikelyBotRequestsAction,omitnil,omitempty" name:"LikelyBotRequestsAction"`

	// 友好 Bot 请求的执行动作。 SecurityAction 的 Name 取值支持：<li>Deny：拦截；</li><li>Monitor：观察；</li><li>Allow：放行；</li><li>Challenge：挑战，其中 ChallengeActionParameters 中的 ChallengeOption 支持 JSChallenge 和 ManagedChallenge。</li>
	VerifiedBotRequestsAction *SecurityAction `json:"VerifiedBotRequestsAction,omitnil,omitempty" name:"VerifiedBotRequestsAction"`

	// 正常 Bot 请求的执行动作。 SecurityAction 的 Name 取值支持：<li>Allow：放行。</li>
	HumanRequestsAction *SecurityAction `json:"HumanRequestsAction,omitnil,omitempty" name:"HumanRequestsAction"`
}

type BotSessionValidation struct {
	// 是否更新 Cookie 并校验。取值有：<li>on：更新 Cookie 并校验；</li><li>off：仅校验。</li>
	IssueNewBotSessionCookie *string `json:"IssueNewBotSessionCookie,omitnil,omitempty" name:"IssueNewBotSessionCookie"`

	// 更新 Cookie 并校验时的触发阈值，仅当 IssueNewBotSessionCookie 为 on 时有效。
	MaxNewSessionTriggerConfig *MaxNewSessionTriggerConfig `json:"MaxNewSessionTriggerConfig,omitnil,omitempty" name:"MaxNewSessionTriggerConfig"`

	// 未携带 Cookie 或 Cookie 已过期的执行动作。 SecurityAction 的 Name 取值支持：<li>Deny：拦截，其中 DenyActionParameters 中支持 Stall 配置；</li><li>Monitor：观察；</li><li>Allow：等待后响应，其中 AllowActionParameters 需要 MinDelayTime 和 MaxDelayTime 配置。</li>
	SessionExpiredAction *SecurityAction `json:"SessionExpiredAction,omitnil,omitempty" name:"SessionExpiredAction"`

	// 不合法 Cookie 的执行动作。 SecurityAction 的 Name 取值支持：<li>Deny：拦截，其中 DenyActionParameters 中支持 Stall 配置；</li><li>Monitor：观察；</li><li>Allow：等待后响应，其中 AllowActionParameters 需要 MinDelayTime 和 MaxDelayTime 配置。</li>
	SessionInvalidAction *SecurityAction `json:"SessionInvalidAction,omitnil,omitempty" name:"SessionInvalidAction"`

	// 会话速率和周期特征校验的具体配置。
	SessionRateControl *SessionRateControl `json:"SessionRateControl,omitnil,omitempty" name:"SessionRateControl"`
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

type BrowserImpersonationDetection struct {
	// 浏览器伪造识别规则的列表。使用 ModifySecurityPolicy 修改 Web 防护配置时： <br> <li>  若未指定 SecurityPolicy.BotManagement.BrowserImpersonationDetection 中的 Rules 参数，或 Rules 参数长度为零： 清空所有浏览器伪造识别规则配置。</li> <li> 若 SecurityPolicy.BotManagement 参数中，未指定 BrowserImpersonationDetection 参数值： 保持已有浏览器伪造识别规则配置，不做修改。</li>
	Rules []*BrowserImpersonationDetectionRule `json:"Rules,omitnil,omitempty" name:"Rules"`
}

type BrowserImpersonationDetectionAction struct {
	// Cookie 校验和会话跟踪配置。
	BotSessionValidation *BotSessionValidation `json:"BotSessionValidation,omitnil,omitempty" name:"BotSessionValidation"`

	// 客户端行为校验配置。
	ClientBehaviorDetection *ClientBehaviorDetection `json:"ClientBehaviorDetection,omitnil,omitempty" name:"ClientBehaviorDetection"`
}

type BrowserImpersonationDetectionRule struct {
	// 浏览器伪造识别规则的 ID。<br>通过规则 ID 可支持不同的规则配置操作：<br> <li> <b>增加</b>新规则：ID 为空或不指定 ID 参数；</li><li><b>修改</b>已有规则：指定需要更新/修改的规则 ID；</li><li><b>删除</b>已有规则：BrowserImpersonationDetection 参数中，Rules 列表中未包含的已有规则将被删除。</li>
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 浏览器伪造识别规则的名称。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 浏览器伪造识别规则是否开启。取值有：<li>on：开启；</li><li>off：关闭。</li>
	Enabled *string `json:"Enabled,omitnil,omitempty" name:"Enabled"`

	// 浏览器伪造识别规则的具体内容，其中仅支持请求方式（Method）、请求路径（Path）和请求 URL 的配置，需符合表达式语法，详细规范参见产品文档。
	Condition *string `json:"Condition,omitnil,omitempty" name:"Condition"`

	// 浏览器伪造识别规则的处置方式，包括 Cookie 校验和会话跟踪配置以及客户端行为校验配置。
	Action *BrowserImpersonationDetectionAction `json:"Action,omitnil,omitempty" name:"Action"`
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

type CNAMEDetail struct {
	// 是否伪站点，取值有：
	// <li> 0：非伪站点；</li>
	// <li> 1：伪站点。</li>
	IsFake *int64 `json:"IsFake,omitnil,omitempty" name:"IsFake"`

	// 归属权验证信息。详情请参考 [站点/域名归属权验证](https://cloud.tencent.com/document/product/1552/70789) 。
	OwnershipVerification *OwnershipVerification `json:"OwnershipVerification,omitnil,omitempty" name:"OwnershipVerification"`
}

type Cache struct {
	// 缓存配置开关，取值有：
	// <li>on：开启；</li>
	// <li>off：关闭。</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// 缓存过期时间设置。
	// 单位为秒，最大可设置为 365 天。
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

type CacheConfigCustomTime struct {
	// 自定义缓存时间开关，取值有：
	// <li>on：开启；</li>
	// <li>off：关闭。</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// 自定义缓存时间数值，单位为秒，取值：0-315360000。<br>注意：当 Switch 为 on 时，此字段必填；当 Switch 为 off 时，无需填写此字段，若填写则不生效。
	CacheTime *int64 `json:"CacheTime,omitnil,omitempty" name:"CacheTime"`
}

type CacheConfigParameters struct {
	// 遵循源站缓存配置。FollowOrigin、NoCache、CustomTime 最多只能配置一个 Switch 为 on。
	// 注意：此字段可能返回 null，表示取不到有效值。
	FollowOrigin *FollowOrigin `json:"FollowOrigin,omitnil,omitempty" name:"FollowOrigin"`

	// 不缓存配置。FollowOrigin、NoCache、CustomTime 最多只能配置一个 Switch 为 on。
	// 注意：此字段可能返回 null，表示取不到有效值。
	NoCache *NoCache `json:"NoCache,omitnil,omitempty" name:"NoCache"`

	// 自定义缓存时间配置。FollowOrigin、NoCache、CustomTime 最多只能配置一个 Switch 为 on。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CustomTime *CacheConfigCustomTime `json:"CustomTime,omitnil,omitempty" name:"CustomTime"`
}

type CacheKey struct {
	// 是否开启全路径缓存，取值有：
	// <li>on：开启全路径缓存（即关闭参数忽略）；</li>
	// <li>off：关闭全路径缓存（即开启参数忽略）。</li>
	FullUrlCache *string `json:"FullUrlCache,omitnil,omitempty" name:"FullUrlCache"`

	// 是否忽略大小写缓存，取值有：
	// <li>on：忽略；</li>
	// <li>off：不忽略。</li>
	IgnoreCase *string `json:"IgnoreCase,omitnil,omitempty" name:"IgnoreCase"`

	// CacheKey 中包含请求参数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	QueryString *QueryString `json:"QueryString,omitnil,omitempty" name:"QueryString"`
}

type CacheKeyConfigParameters struct {
	// 是否开启全路径缓存，取值有：
	// <li>on：开启全路径缓存（即关闭参数忽略）；</li>
	// <li>off：关闭全路径缓存（即开启参数忽略）。</li>
	FullURLCache *string `json:"FullURLCache,omitnil,omitempty" name:"FullURLCache"`

	// 是否忽略大小写缓存，取值有：
	// <li>on：忽略；</li>
	// <li>off：不忽略。</li>
	IgnoreCase *string `json:"IgnoreCase,omitnil,omitempty" name:"IgnoreCase"`

	// 查询字符串保留配置参数。此字段和 FullURLCache 必须同时设置，但不能同为 on。
	QueryString *CacheKeyQueryString `json:"QueryString,omitnil,omitempty" name:"QueryString"`
}

type CacheKeyCookie struct {
	// 功能开关，取值有：
	// <li>on：开启；</li>
	// <li>off：关闭。</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// 缓存动作，取值有：
	// <li>full：全部保留；</li>
	// <li> ignore：全部忽略；</li>
	// <li> includeCustom：保留指定参数；</li>
	// <li>excludeCustom：忽略指定参数。</li>注意：当 Switch 为 on 时，此字段必填；当 Switch 为 off 时，无需填写此字段，若填写则不生效。
	Action *string `json:"Action,omitnil,omitempty" name:"Action"`

	// 自定义 Cache Key Cookie 名称列表。<br>注意：当 Action 为 includeCustom 或 excludeCustom 时，此字段必填；当 Action 为 full 或 ignore 时，无需填写此字段，若填写则不生效。
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`
}

type CacheKeyHeader struct {
	// 功能开关，取值有：
	// <li>on：开启；</li>
	// <li>off：关闭。</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// 自定义 Cache Key HTTP 请求头列表。<br>注意：当 Switch 为 on 时，此字段必填；当 Switch 为 off 时，无需填写此字段，若填写则不生效。
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`
}

type CacheKeyParameters struct {
	// 查询字符串全部保留开关，取值有：
	// <li>on：开启；</li>
	// <li>off：关闭。</li>注意：FullURLCache、IgnoreCase、Header、Scheme、Cookie 至少设置一个配置。此字段和 QueryString.Switch 必须同时设置，但不能同为 on。
	FullURLCache *string `json:"FullURLCache,omitnil,omitempty" name:"FullURLCache"`

	// 查询字符串保留配置参数。此字段和 FullURLCache 必须同时设置，但不能同为 on。
	// 注意：此字段可能返回 null，表示取不到有效值。
	QueryString *CacheKeyQueryString `json:"QueryString,omitnil,omitempty" name:"QueryString"`

	// 忽略大小写开关，取值有：
	// <li>on：开启；</li>
	// <li>off：关闭。</li>注意：FullURLCache、IgnoreCase、Header、Scheme、Cookie 至少设置一个配置。
	IgnoreCase *string `json:"IgnoreCase,omitnil,omitempty" name:"IgnoreCase"`

	// HTTP 请求头配置参数。FullURLCache、IgnoreCase、Header、Scheme、Cookie 至少设置一个配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Header *CacheKeyHeader `json:"Header,omitnil,omitempty" name:"Header"`

	// 请求协议开关，取值有：
	// <li>on：开启；</li>
	// <li>off：关闭。</li>注意：FullURLCache、IgnoreCase、Header、Scheme、Cookie 至少设置一个配置。
	Scheme *string `json:"Scheme,omitnil,omitempty" name:"Scheme"`

	// Cookie 配置参数。FullURLCache、IgnoreCase、Header、Scheme、Cookie 至少设置一个配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Cookie *CacheKeyCookie `json:"Cookie,omitnil,omitempty" name:"Cookie"`
}

type CacheKeyQueryString struct {
	// 查询字符串保留/忽略指定参数开关，取值有：
	// <li>on：开启；</li>
	// <li>off：关闭。</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// 查询字符串保留/忽略指定参数动作。取值有：
	// <li>includeCustom：表示保留部分参数；</li>
	// <li>excludeCustom：表示忽略部分参数。</li>注意：当 Switch 为 on 时，此字段必填；当 Switch 为 off 时，无需填写此字段，若填写则不生效。
	Action *string `json:"Action,omitnil,omitempty" name:"Action"`

	// 查询字符串中需保留/忽略的参数名列表。<br>注意：当 Switch 为 on 时，此字段必填；当 Switch 为 off 时，无需填写此字段，若填写则不生效。
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`
}

type CacheParameters struct {
	// 缓存遵循源站。不填表示不设置该配置，FollowOrigin、NoCache、CustomTime 最多只能配置一个 Switch 为 on。
	// 注意：此字段可能返回 null，表示取不到有效值。
	FollowOrigin *FollowOrigin `json:"FollowOrigin,omitnil,omitempty" name:"FollowOrigin"`

	// 不缓存。不填表示不设置该配置，FollowOrigin、NoCache、CustomTime 最多只能配置一个 Switch 为 on。
	// 注意：此字段可能返回 null，表示取不到有效值。
	NoCache *NoCache `json:"NoCache,omitnil,omitempty" name:"NoCache"`

	// 自定义缓存时间。不填表示不设置该配置，FollowOrigin、NoCache、CustomTime 最多只能配置一个 Switch 为 on。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CustomTime *CustomTime `json:"CustomTime,omitnil,omitempty" name:"CustomTime"`
}

type CachePrefresh struct {
	// 缓存预刷新配置开关，取值有：
	// <li>on：开启；</li>
	// <li>off：关闭。</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// 缓存预刷新百分比，取值范围：1-99。
	Percent *int64 `json:"Percent,omitnil,omitempty" name:"Percent"`
}

type CachePrefreshParameters struct {
	// 缓存预刷新开关，取值有：
	// <li>on：开启；</li>
	// <li>off：关闭。</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// 预刷新时间设置为节点缓存时间的百分比数值，取值：1～99。<br>注意：当 Switch 为 on 时，此字段必填；当 Switch 为 off 时，无需填写此字段，若填写则不生效。
	CacheTimePercent *int64 `json:"CacheTimePercent,omitnil,omitempty" name:"CacheTimePercent"`
}

type CacheTag struct {
	// 待清除缓存的域名列表。
	Domains []*string `json:"Domains,omitnil,omitempty" name:"Domains"`
}

type CertificateInfo struct {
	// 证书 ID。来源于 SSL 侧，您可以前往 [SSL 证书列表](https://console.cloud.tencent.com/ssl) 查看 CertId。
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

type ChallengeActionParameters struct {
	// 安全执行的具体挑战动作。取值有：<li> InterstitialChallenge：插页式挑战；</li><li> InlineChallenge：内嵌式挑战；</li><li> JSChallenge：JavaScript 挑战；</li><li> ManagedChallenge：托管挑战。</li>
	ChallengeOption *string `json:"ChallengeOption,omitnil,omitempty" name:"ChallengeOption"`

	// 重复挑战的时间间隔，当 Name 为 InterstitialChallenge/InlineChallenge 时，该字段必填。默认值为 300s。支持的单位有：<li>s：秒，取值范围1～60；</li><li>m：分，取值范围1～60；</li><li>h：小时，取值范围1～24。</li>
	Interval *string `json:"Interval,omitnil,omitempty" name:"Interval"`

	// 客户端认证方式 ID 。当 Name 为 InterstitialChallenge/InlineChallenge 时，该字段必填。
	AttesterId *string `json:"AttesterId,omitnil,omitempty" name:"AttesterId"`
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

// Predefined struct for user
type CheckFreeCertificateVerificationRequestParams struct {
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 加速域名，该域名为[申请免费证书](https://cloud.tencent.com/document/product/1552/124807)时使用的域名。
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`
}

type CheckFreeCertificateVerificationRequest struct {
	*tchttp.BaseRequest
	
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 加速域名，该域名为[申请免费证书](https://cloud.tencent.com/document/product/1552/124807)时使用的域名。
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`
}

func (r *CheckFreeCertificateVerificationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckFreeCertificateVerificationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "Domain")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CheckFreeCertificateVerificationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CheckFreeCertificateVerificationResponseParams struct {
	// 免费证书申请成功时，该证书颁发给的域名。
	// 注意：一个域名只允许申请一本免费证书， 如果已经有泛域名申请了免费证书的情况下，其子域名会匹配使用该泛域名证书。
	CommonName *string `json:"CommonName,omitnil,omitempty" name:"CommonName"`

	// 免费证书申请成功时，该证书使用的签名算法，当前仅支持 RSA 2048。
	SignatureAlgorithm *string `json:"SignatureAlgorithm,omitnil,omitempty" name:"SignatureAlgorithm"`

	// 免费证书申请成功时，该证书的过期时间。时间为世界标准时间（UTC）， 遵循 ISO 8601 标准的日期和时间格式。
	ExpireTime *string `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CheckFreeCertificateVerificationResponse struct {
	*tchttp.BaseResponse
	Response *CheckFreeCertificateVerificationResponseParams `json:"Response"`
}

func (r *CheckFreeCertificateVerificationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckFreeCertificateVerificationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CheckRegionHealthStatus struct {
	// 健康检查区域，ISO-3166-1 两位字母代码。
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 单健康检查区域下探测源站的健康状态，取值有：
	// <li>Healthy：健康；</li>
	// <li>Unhealthy：不健康；</li>
	// <li> Undetected：未探测到数据。</li>说明：单健康检查区域下所有源站为健康，则状态为健康，否则为不健康。
	Healthy *string `json:"Healthy,omitnil,omitempty" name:"Healthy"`

	// 源站健康状态。
	OriginHealthStatus []*OriginHealthStatus `json:"OriginHealthStatus,omitnil,omitempty" name:"OriginHealthStatus"`
}

type ClientAttestationRule struct {
	// 客户端认证规则的 ID。<br>通过规则 ID 可支持不同的规则配置操作：<br> <li> <b>增加</b>新规则：ID 为空或不指定 ID 参数；</li><li> <b>修改</b>已有规则：指定需要更新/修改的规则 ID；</li><li> <b>删除</b>已有规则：BotManagement 参数中，ClientAttestationRule 列表中未包含的已有规则将被删除。</li>
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 客户端认证规则的名称。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 规则是否开启。取值有：<li>on：开启；</li><li>off：关闭。</li>
	Enabled *string `json:"Enabled,omitnil,omitempty" name:"Enabled"`

	// 规则的优先级，数值越小越优先执行，范围是 0 ~ 100，默认为 0。
	Priority *uint64 `json:"Priority,omitnil,omitempty" name:"Priority"`

	// 规则的具体内容，需符合表达式语法，详细规范参见产品文档。
	Condition *string `json:"Condition,omitnil,omitempty" name:"Condition"`

	// 客户端认证选项 ID。
	AttesterId *string `json:"AttesterId,omitnil,omitempty" name:"AttesterId"`

	// 客户端设备配置。若 ClientAttestationRules 参数中，未指定 DeviceProfiles 参数值：保持已有客户端设备配置，不做修改。
	DeviceProfiles []*DeviceProfile `json:"DeviceProfiles,omitnil,omitempty" name:"DeviceProfiles"`

	// 客户端认证未通过的处置方式。SecurityAction 的 Name 取值支持：<li>Deny：拦截；</li><li>Monitor：观察；</li><li>Redirect：重定向；</li><li>Challenge：挑战。</li>默认值为 Monitor。
	InvalidAttestationAction *SecurityAction `json:"InvalidAttestationAction,omitnil,omitempty" name:"InvalidAttestationAction"`
}

type ClientAttestationRules struct {
	// 客户端认证的列表。使用 ModifySecurityPolicy 修改 Web 防护配置时：<li>  若未指定 SecurityPolicy.BotManagement.ClientAttestationRules 中的 Rules 参数，或 Rules 参数长度为零：清空所有客户端认证规则配置。</li> <li> 若 SecurityPolicy.BotManagement 参数中，未指定 ClientAttestationRules 参数值：保持已有客户端认证规则配置，不做修改。</li>
	Rules []*ClientAttestationRule `json:"Rules,omitnil,omitempty" name:"Rules"`
}

type ClientAttester struct {
	// 认证选项 ID。
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 认证选项名称。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 认证规则类型。仅出参返回，取值有：
	// <li>PRESET: 系统预置规则，仅允许修改 AttesterDuration；</li>
	// <li>CUSTOM: 用户自定义规则。</li>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 认证方法。取值有：
	// <li>TC-RCE: 使用风险识别 RCE 进行认证；</li>
	// <li>TC-CAPTCHA: 使用天御验证码进行认证。</li>
	AttesterSource *string `json:"AttesterSource,omitnil,omitempty" name:"AttesterSource"`

	// 认证有效时间。默认为 60s，支持的单位有：
	// <li>s：秒，取值范围 60～43200；</li>
	// <li>m：分，取值范围 1～720；</li>
	// <li>h：小时，取值范围 1～12。</li>
	AttesterDuration *string `json:"AttesterDuration,omitnil,omitempty" name:"AttesterDuration"`

	// TC-RCE 认证的配置信息。
	// <li>当 AttesterSource 参数值为 TC-RCE 时，此字段必填。</li>
	TCRCEOption *TCRCEOption `json:"TCRCEOption,omitnil,omitempty" name:"TCRCEOption"`

	// TC-CAPTCHA 认证的配置信息。
	// <li>当 AttesterSource 参数值为 TC-CAPTCHA 时，此字段必填。</li>
	TCCaptchaOption *TCCaptchaOption `json:"TCCaptchaOption,omitnil,omitempty" name:"TCCaptchaOption"`
}

type ClientBehaviorDetection struct {
	// 工作量证明校验强度。取值有：<li>low：低；</li><li>medium：中；</li><li>high：高。</li>
	CryptoChallengeIntensity *string `json:"CryptoChallengeIntensity,omitnil,omitempty" name:"CryptoChallengeIntensity"`

	// 客户端行为校验的执行方式。取值有：<li>0ms：立即执行；</li><li>100ms：延迟 100ms 执行；</li><li>200ms：延迟 200ms 执行；</li><li>300ms：延迟 300ms 执行；</li><li>400ms：延迟 400ms 执行；</li><li>500ms：延迟 500ms 执行；</li><li>600ms：延迟 600ms 执行；</li><li>700ms：延迟 700ms 执行；</li><li>800ms：延迟 800ms 执行；</li><li>900ms：延迟 900ms 执行；</li><li>1000ms：延迟 1000ms 执行。</li> 
	CryptoChallengeDelayBefore *string `json:"CryptoChallengeDelayBefore,omitnil,omitempty" name:"CryptoChallengeDelayBefore"`

	// 触发阈值统计的时间窗口，取值有：<li>5s：5 秒内；</li><li>10s：10 秒内；</li><li>15s：15 秒内；</li><li>30s：30 秒内；</li><li>60s：60 秒内；</li><li>5m：5 分钟内；</li><li>10m：10 分钟内；</li><li>30m：30 分钟内；</li><li>60m：60 分钟内。</li> 
	MaxChallengeCountInterval *string `json:"MaxChallengeCountInterval,omitnil,omitempty" name:"MaxChallengeCountInterval"`

	// 触发阈值统计的累计次数，取值范围 1 ~ 100000000。
	MaxChallengeCountThreshold *int64 `json:"MaxChallengeCountThreshold,omitnil,omitempty" name:"MaxChallengeCountThreshold"`

	// 客户端未启用 JS（未完成检测）时的执行动作。 SecurityAction 的 Name 取值支持：<li>Deny：拦截，其中 DenyActionParameters 中支持 Stall 配置；</li><li>Monitor：观察；</li><li>Allow：等待后响应，其中 AllowActionParameters 需要 MinDelayTime 和 MaxDelayTime 配置。</li>
	ChallengeNotFinishedAction *SecurityAction `json:"ChallengeNotFinishedAction,omitnil,omitempty" name:"ChallengeNotFinishedAction"`

	// 客户端检测超时的执行动作。 SecurityAction 的 Name 取值支持：<li>Deny：拦截，其中 DenyActionParameters 中支持 Stall 配置；</li><li>Monitor：观察；</li><li>Allow：等待后响应，其中 AllowActionParameters 需要 MinDelayTime 和 MaxDelayTime 配置。</li>
	ChallengeTimeoutAction *SecurityAction `json:"ChallengeTimeoutAction,omitnil,omitempty" name:"ChallengeTimeoutAction"`

	// Bot 客户端的执行动作。 SecurityAction 的 Name 取值支持：<li>Deny：拦截，其中 DenyActionParameters 中支持 Stall 配置；</li><li>Monitor：观察；</li><li>Allow：等待后响应，其中 AllowActionParameters 需要 MinDelayTime 和 MaxDelayTime 配置。</li>
	BotClientAction *SecurityAction `json:"BotClientAction,omitnil,omitempty" name:"BotClientAction"`
}

type ClientFiltering struct {
	// 智能客户端过滤是否开启。取值有：<li>on：开启；</li><li>off：关闭。</li>
	Enabled *string `json:"Enabled,omitnil,omitempty" name:"Enabled"`

	// 智能客户端过滤的处置方式，当 Enabled 为 on 时，此字段必填。SecurityAction 的 Name 取值支持：<li>Monitor：观察；</li><li>Deny：拦截；</li><li>Challenge：挑战，其中ChallengeActionParameters.Name仅支持JSChallenge。</li>
	Action *SecurityAction `json:"Action,omitnil,omitempty" name:"Action"`
}

type ClientIPCountryParameters struct {
	// 配置开关，取值有：
	// <li>on：开启；</li>
	// <li>off：关闭。</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// 存放客户端 IP 所属地域信息的请求头名称，当 Switch=on 时有效。为空则使用默认值：EO-Client-IPCountry。
	HeaderName *string `json:"HeaderName,omitnil,omitempty" name:"HeaderName"`
}

type ClientIPHeaderParameters struct {
	// 配置开关，取值有：
	// <li>on：开启；</li>
	// <li>off：关闭。</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// 回源时，存放客户端 IP 的请求头名称。当 Switch 为 on 时，该参数必填。该参数不允许填写 X-Forwarded-For。
	HeaderName *string `json:"HeaderName,omitnil,omitempty" name:"HeaderName"`
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

	// 回源时，存放客户端 IP 的请求头名称。当 Switch 为 on 时，该参数必填。该参数不允许填写 X-Forwarded-For。
	HeaderName *string `json:"HeaderName,omitnil,omitempty" name:"HeaderName"`
}

type CnameStatus struct {
	// 记录名称。
	RecordName *string `json:"RecordName,omitnil,omitempty" name:"RecordName"`

	// CNAME 地址。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Cname *string `json:"Cname,omitnil,omitempty" name:"Cname"`

	// CNAME 状态信息，取值有：
	// <li>active：生效；</li>
	// <li>moved：不生效；</li>
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
	Algorithms []*string `json:"Algorithms,omitnil,omitempty" name:"Algorithms"`
}

type CompressionParameters struct {
	// 智能压缩配置开关，取值有：
	// <li>on：开启；</li>
	// <li>off：关闭。</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// 支持的压缩算法列表。当 Switch 为 on 时，此字段必填，否则此字段不生效。取值有：
	// <li>brotli：brotli 算法；</li>
	// <li>gzip：gzip 算法。</li>
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
type ConfirmMultiPathGatewayOriginACLRequestParams struct {
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 网关 ID。
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// 回源 IP 版本号。
	OriginACLVersion *int64 `json:"OriginACLVersion,omitnil,omitempty" name:"OriginACLVersion"`
}

type ConfirmMultiPathGatewayOriginACLRequest struct {
	*tchttp.BaseRequest
	
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 网关 ID。
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// 回源 IP 版本号。
	OriginACLVersion *int64 `json:"OriginACLVersion,omitnil,omitempty" name:"OriginACLVersion"`
}

func (r *ConfirmMultiPathGatewayOriginACLRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ConfirmMultiPathGatewayOriginACLRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "GatewayId")
	delete(f, "OriginACLVersion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ConfirmMultiPathGatewayOriginACLRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ConfirmMultiPathGatewayOriginACLResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ConfirmMultiPathGatewayOriginACLResponse struct {
	*tchttp.BaseResponse
	Response *ConfirmMultiPathGatewayOriginACLResponseParams `json:"Response"`
}

func (r *ConfirmMultiPathGatewayOriginACLResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ConfirmMultiPathGatewayOriginACLResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ConfirmOriginACLUpdateRequestParams struct {
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`
}

type ConfirmOriginACLUpdateRequest struct {
	*tchttp.BaseRequest
	
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`
}

func (r *ConfirmOriginACLUpdateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ConfirmOriginACLUpdateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ConfirmOriginACLUpdateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ConfirmOriginACLUpdateResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ConfirmOriginACLUpdateResponse struct {
	*tchttp.BaseResponse
	Response *ConfirmOriginACLUpdateResponseParams `json:"Response"`
}

func (r *ConfirmOriginACLUpdateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ConfirmOriginACLUpdateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ContentCompressionParameters struct {
	// 内容压缩配置开关，取值有：
	// <li>on：开启；</li>
	// <li>off：关闭。</li>
	// 当 Switch 为 on 时，将同时支持 brotli 和 gzip 压缩算法。
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`
}

type ContentIdentifier struct {
	// 内容标识符 ID。
	ContentId *string `json:"ContentId,omitnil,omitempty" name:"ContentId"`

	// 内容标识符描述。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 被规则引擎引用的次数。
	ReferenceCount *int64 `json:"ReferenceCount,omitnil,omitempty" name:"ReferenceCount"`

	// 绑定的套餐 ID。
	PlanId *string `json:"PlanId,omitnil,omitempty" name:"PlanId"`

	// 绑定的标签。
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 内容标识符状态，取值有：
	// <li> active：已生效； </li>
	// <li> deleted：已删除。</li>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 创建时间，时间为世界标准时间（UTC）， 遵循 ISO 8601 标准的日期和时间格式。
	CreatedOn *string `json:"CreatedOn,omitnil,omitempty" name:"CreatedOn"`

	// 最新一次更新时间，时间为世界标准时间（UTC）， 遵循 ISO 8601 标准的日期和时间格式。
	ModifiedOn *string `json:"ModifiedOn,omitnil,omitempty" name:"ModifiedOn"`

	// 删除时间，状态非 deleted 时候为空；时间为世界标准时间（UTC）， 遵循 ISO 8601 标准的日期和时间格式。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeletedOn *string `json:"DeletedOn,omitnil,omitempty" name:"DeletedOn"`
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
type CreateContentIdentifierRequestParams struct {
	// 内容标识符的描述，长度限制不超过 20 个字符。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 待绑定的目标套餐 ID，仅限企业版可用。<li>当您账号下已存在套餐时，需要先前往 [套餐管理](https://console.cloud.tencent.com/edgeone/package) 获取套餐 ID，直接将内容标识符绑定至该套餐；</li><li>若您当前没有可绑定的套餐时，请先购买企业版套餐。</li>
	PlanId *string `json:"PlanId,omitnil,omitempty" name:"PlanId"`

	// 标签。该参数用于对内容标识符进行分权限管控。您需要先前往 [标签控制台](https://console.cloud.tencent.com/tag/taglist) 创建标签才可以在此处传入对应的标签键和标签值。
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type CreateContentIdentifierRequest struct {
	*tchttp.BaseRequest
	
	// 内容标识符的描述，长度限制不超过 20 个字符。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 待绑定的目标套餐 ID，仅限企业版可用。<li>当您账号下已存在套餐时，需要先前往 [套餐管理](https://console.cloud.tencent.com/edgeone/package) 获取套餐 ID，直接将内容标识符绑定至该套餐；</li><li>若您当前没有可绑定的套餐时，请先购买企业版套餐。</li>
	PlanId *string `json:"PlanId,omitnil,omitempty" name:"PlanId"`

	// 标签。该参数用于对内容标识符进行分权限管控。您需要先前往 [标签控制台](https://console.cloud.tencent.com/tag/taglist) 创建标签才可以在此处传入对应的标签键和标签值。
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

func (r *CreateContentIdentifierRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateContentIdentifierRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Description")
	delete(f, "PlanId")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateContentIdentifierRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateContentIdentifierResponseParams struct {
	// 生成的内容标识符 ID。创建完成之后您可以前往规则引擎在一定匹配条件下「设置内容标识符」。
	ContentId *string `json:"ContentId,omitnil,omitempty" name:"ContentId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateContentIdentifierResponse struct {
	*tchttp.BaseResponse
	Response *CreateContentIdentifierResponseParams `json:"Response"`
}

func (r *CreateContentIdentifierResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateContentIdentifierResponse) FromJsonString(s string) error {
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
type CreateDnsRecordRequestParams struct {
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// DNS 记录名，如果是中文、韩文、日文域名，需要转换为 punycode 后输入。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// DNS 记录类型，取值有：<li>A：将域名指向一个外网 IPv4 地址，如 8.8.8.8；</li><li>AAAA：将域名指向一个外网 IPv6 地址；</li><li>MX：用于邮箱服务器。存在多条 MX 记录时，优先级越低越优先；</li><li>CNAME：将域名指向另一个域名，再由该域名解析出最终 IP 地址；</li><li>TXT：对域名进行标识和说明，常用于域名验证和 SPF 记录（反垃圾邮件）；</li><li>NS：如果需要将子域名交给其他 DNS 服务商解析，则需要添加 NS 记录。根域名无法添加 NS 记录；</li><li>CAA：指定可为本站点颁发证书的 CA；</li><li>SRV：标识某台服务器使用了某个服务，常见于微软系统的目录管理。</li>
	// 不同的记录类型呢例如 SRV、CAA 记录对主机记录名称、记录值格式有不同的要求，各记录类型的详细说明介绍和格式示例请参考：[解析记录类型介绍](https://cloud.tencent.com/document/product/1552/90453#2f681022-91ab-4a9e-ac3d-0a6c454d954e)。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// DNS 记录内容，根据 Type 值填入与之相对应的内容，如果是中文、韩文、日文域名，需要转换为 punycode 后输入。
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`

	// DNS 记录解析线路，不指定默认为 Default，表示默认解析线路，代表全部地域生效。
	// 
	// - 解析线路配置仅适用于当 Type（DNS 记录类型）为 A、AAAA、CNAME 时。
	// - 解析线路配置仅适用于标准版、企业版套餐使用，取值请参考：[解析线路及对应代码枚举](https://cloud.tencent.com/document/product/1552/112542)。
	Location *string `json:"Location,omitnil,omitempty" name:"Location"`

	// 缓存时间，用户可指定值范围 60~86400，数值越小，修改记录各地生效时间越快，默认为 300，单位：秒。
	TTL *int64 `json:"TTL,omitnil,omitempty" name:"TTL"`

	// DNS 记录权重，用户可指定值范围 -1~100，设置为 0 时表示不解析，不指定默认为 -1，表示不设置权重。权重配置仅适用于当 Type（DNS 记录类型）为 A、AAAA、CNAME 时。<br>注意：同一个子域名下，相同解析线路的不同 DNS 记录，应保持同时设置权重或者同时都不设置权重。
	Weight *int64 `json:"Weight,omitnil,omitempty" name:"Weight"`

	// MX 记录优先级，该参数仅在当 Type（DNS 记录类型）为 MX 时生效，值越小优先级越高，用户可指定值范围0~50，不指定默认为0。
	Priority *int64 `json:"Priority,omitnil,omitempty" name:"Priority"`
}

type CreateDnsRecordRequest struct {
	*tchttp.BaseRequest
	
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// DNS 记录名，如果是中文、韩文、日文域名，需要转换为 punycode 后输入。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// DNS 记录类型，取值有：<li>A：将域名指向一个外网 IPv4 地址，如 8.8.8.8；</li><li>AAAA：将域名指向一个外网 IPv6 地址；</li><li>MX：用于邮箱服务器。存在多条 MX 记录时，优先级越低越优先；</li><li>CNAME：将域名指向另一个域名，再由该域名解析出最终 IP 地址；</li><li>TXT：对域名进行标识和说明，常用于域名验证和 SPF 记录（反垃圾邮件）；</li><li>NS：如果需要将子域名交给其他 DNS 服务商解析，则需要添加 NS 记录。根域名无法添加 NS 记录；</li><li>CAA：指定可为本站点颁发证书的 CA；</li><li>SRV：标识某台服务器使用了某个服务，常见于微软系统的目录管理。</li>
	// 不同的记录类型呢例如 SRV、CAA 记录对主机记录名称、记录值格式有不同的要求，各记录类型的详细说明介绍和格式示例请参考：[解析记录类型介绍](https://cloud.tencent.com/document/product/1552/90453#2f681022-91ab-4a9e-ac3d-0a6c454d954e)。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// DNS 记录内容，根据 Type 值填入与之相对应的内容，如果是中文、韩文、日文域名，需要转换为 punycode 后输入。
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`

	// DNS 记录解析线路，不指定默认为 Default，表示默认解析线路，代表全部地域生效。
	// 
	// - 解析线路配置仅适用于当 Type（DNS 记录类型）为 A、AAAA、CNAME 时。
	// - 解析线路配置仅适用于标准版、企业版套餐使用，取值请参考：[解析线路及对应代码枚举](https://cloud.tencent.com/document/product/1552/112542)。
	Location *string `json:"Location,omitnil,omitempty" name:"Location"`

	// 缓存时间，用户可指定值范围 60~86400，数值越小，修改记录各地生效时间越快，默认为 300，单位：秒。
	TTL *int64 `json:"TTL,omitnil,omitempty" name:"TTL"`

	// DNS 记录权重，用户可指定值范围 -1~100，设置为 0 时表示不解析，不指定默认为 -1，表示不设置权重。权重配置仅适用于当 Type（DNS 记录类型）为 A、AAAA、CNAME 时。<br>注意：同一个子域名下，相同解析线路的不同 DNS 记录，应保持同时设置权重或者同时都不设置权重。
	Weight *int64 `json:"Weight,omitnil,omitempty" name:"Weight"`

	// MX 记录优先级，该参数仅在当 Type（DNS 记录类型）为 MX 时生效，值越小优先级越高，用户可指定值范围0~50，不指定默认为0。
	Priority *int64 `json:"Priority,omitnil,omitempty" name:"Priority"`
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
	delete(f, "Name")
	delete(f, "Type")
	delete(f, "Content")
	delete(f, "Location")
	delete(f, "TTL")
	delete(f, "Weight")
	delete(f, "Priority")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDnsRecordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDnsRecordResponseParams struct {
	// DNS 记录 ID。
	RecordId *string `json:"RecordId,omitnil,omitempty" name:"RecordId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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

	// 函数选择配置类型：
	// <li> direct：直接指定执行函数；</li>
	// <li> weight：基于权重比选择函数；</li>
	// <li> region：基于客户端 IP 的国家/地区选择函数。</li>
	// 不填时默认为 direct 。
	TriggerType *string `json:"TriggerType,omitnil,omitempty" name:"TriggerType"`

	// 指定执行的函数 ID。当 TriggerType 为 direct 或 TriggerType 不填时生效。
	FunctionId *string `json:"FunctionId,omitnil,omitempty" name:"FunctionId"`

	// 基于客户端 IP 国家/地区的函数选择配置，当 TriggerType 为 region 时生效且 RegionMappingSelections 必填。RegionMappingSelections 中至少包含一项 Regions 为 Default 的配置。
	RegionMappingSelections []*FunctionRegionSelection `json:"RegionMappingSelections,omitnil,omitempty" name:"RegionMappingSelections"`

	// 基于权重的函数选择配置，当 TriggerType 为 weight 时生效且 WeightedSelections 必填。WeightedSelections 中的所有权重之和需要为100。
	WeightedSelections []*FunctionWeightedSelection `json:"WeightedSelections,omitnil,omitempty" name:"WeightedSelections"`

	// 规则描述，最大支持 60 个字符。
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`
}

type CreateFunctionRuleRequest struct {
	*tchttp.BaseRequest
	
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 规则条件列表，相同触发规则的不同条件匹配项之间为或关系。
	FunctionRuleConditions []*FunctionRuleCondition `json:"FunctionRuleConditions,omitnil,omitempty" name:"FunctionRuleConditions"`

	// 函数选择配置类型：
	// <li> direct：直接指定执行函数；</li>
	// <li> weight：基于权重比选择函数；</li>
	// <li> region：基于客户端 IP 的国家/地区选择函数。</li>
	// 不填时默认为 direct 。
	TriggerType *string `json:"TriggerType,omitnil,omitempty" name:"TriggerType"`

	// 指定执行的函数 ID。当 TriggerType 为 direct 或 TriggerType 不填时生效。
	FunctionId *string `json:"FunctionId,omitnil,omitempty" name:"FunctionId"`

	// 基于客户端 IP 国家/地区的函数选择配置，当 TriggerType 为 region 时生效且 RegionMappingSelections 必填。RegionMappingSelections 中至少包含一项 Regions 为 Default 的配置。
	RegionMappingSelections []*FunctionRegionSelection `json:"RegionMappingSelections,omitnil,omitempty" name:"RegionMappingSelections"`

	// 基于权重的函数选择配置，当 TriggerType 为 weight 时生效且 WeightedSelections 必填。WeightedSelections 中的所有权重之和需要为100。
	WeightedSelections []*FunctionWeightedSelection `json:"WeightedSelections,omitnil,omitempty" name:"WeightedSelections"`

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
	delete(f, "TriggerType")
	delete(f, "FunctionId")
	delete(f, "RegionMappingSelections")
	delete(f, "WeightedSelections")
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
type CreateJustInTimeTranscodeTemplateRequestParams struct {
	// 站点ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 即时转码模板名称，长度限制：64 个字符。
	TemplateName *string `json:"TemplateName,omitnil,omitempty" name:"TemplateName"`

	// 模板描述信息，长度限制：256 个字符。默认为空。
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// 启用视频流开关，取值：
	// <li>on：开启；</li>
	// <li>off：关闭。</li>默认值：on。
	VideoStreamSwitch *string `json:"VideoStreamSwitch,omitnil,omitempty" name:"VideoStreamSwitch"`

	// 启用音频流开关，取值：
	// <li>on：开启；</li>
	// <li>off：关闭。</li>默认值：on。
	AudioStreamSwitch *string `json:"AudioStreamSwitch,omitnil,omitempty" name:"AudioStreamSwitch"`

	// 视频流配置参数，当 VideoStreamSwitch 为 on，该字段必填。
	VideoTemplate *VideoTemplateInfo `json:"VideoTemplate,omitnil,omitempty" name:"VideoTemplate"`

	// 音频流配置参数，当 AudioStreamSwitch 为 on，该字段必填。
	AudioTemplate *AudioTemplateInfo `json:"AudioTemplate,omitnil,omitempty" name:"AudioTemplate"`
}

type CreateJustInTimeTranscodeTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 站点ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 即时转码模板名称，长度限制：64 个字符。
	TemplateName *string `json:"TemplateName,omitnil,omitempty" name:"TemplateName"`

	// 模板描述信息，长度限制：256 个字符。默认为空。
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// 启用视频流开关，取值：
	// <li>on：开启；</li>
	// <li>off：关闭。</li>默认值：on。
	VideoStreamSwitch *string `json:"VideoStreamSwitch,omitnil,omitempty" name:"VideoStreamSwitch"`

	// 启用音频流开关，取值：
	// <li>on：开启；</li>
	// <li>off：关闭。</li>默认值：on。
	AudioStreamSwitch *string `json:"AudioStreamSwitch,omitnil,omitempty" name:"AudioStreamSwitch"`

	// 视频流配置参数，当 VideoStreamSwitch 为 on，该字段必填。
	VideoTemplate *VideoTemplateInfo `json:"VideoTemplate,omitnil,omitempty" name:"VideoTemplate"`

	// 音频流配置参数，当 AudioStreamSwitch 为 on，该字段必填。
	AudioTemplate *AudioTemplateInfo `json:"AudioTemplate,omitnil,omitempty" name:"AudioTemplate"`
}

func (r *CreateJustInTimeTranscodeTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateJustInTimeTranscodeTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "TemplateName")
	delete(f, "Comment")
	delete(f, "VideoStreamSwitch")
	delete(f, "AudioStreamSwitch")
	delete(f, "VideoTemplate")
	delete(f, "AudioTemplate")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateJustInTimeTranscodeTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateJustInTimeTranscodeTemplateResponseParams struct {
	// 即时转码模板唯一标识，用于即时转码 URL 拼接。
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateJustInTimeTranscodeTemplateResponse struct {
	*tchttp.BaseResponse
	Response *CreateJustInTimeTranscodeTemplateResponseParams `json:"Response"`
}

func (r *CreateJustInTimeTranscodeTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateJustInTimeTranscodeTemplateResponse) FromJsonString(s string) error {
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
type CreateL7AccRulesRequestParams struct {
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 规则内容。
	Rules []*RuleEngineItem `json:"Rules,omitnil,omitempty" name:"Rules"`
}

type CreateL7AccRulesRequest struct {
	*tchttp.BaseRequest
	
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 规则内容。
	Rules []*RuleEngineItem `json:"Rules,omitnil,omitempty" name:"Rules"`
}

func (r *CreateL7AccRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateL7AccRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "Rules")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateL7AccRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateL7AccRulesResponseParams struct {
	// 规则 ID 列表。
	RuleIds []*string `json:"RuleIds,omitnil,omitempty" name:"RuleIds"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateL7AccRulesResponse struct {
	*tchttp.BaseResponse
	Response *CreateL7AccRulesResponseParams `json:"Response"`
}

func (r *CreateL7AccRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateL7AccRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateLoadBalancerRequestParams struct {
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 实例名称，可输入 1-200 个字符，允许字符为 a-z，A-Z，0-9，_，-。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 实例类型，取值有：
	// <li>HTTP：HTTP 专用型，支持添加 HTTP 专用型和通用型源站组，仅支持被站点加速相关服务引用（如域名服务和规则引擎）；</li>
	// <li>GENERAL：通用型，仅支持添加通用型源站组，能被站点加速服务（如域名服务和规则引擎）和四层代理引用。</li>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 源站组列表及其对应的容灾调度优先级。详情请参考 [快速创建负载均衡实例](https://cloud.tencent.com/document/product/1552/104223) 中的示例场景。
	OriginGroups []*OriginGroupInLoadBalancer `json:"OriginGroups,omitnil,omitempty" name:"OriginGroups"`

	// 健康检查策略。详情请参考 [健康检查策略介绍](https://cloud.tencent.com/document/product/1552/104228)。不填写时，默认为不启用健康检查。
	HealthChecker *HealthChecker `json:"HealthChecker,omitnil,omitempty" name:"HealthChecker"`

	// 源站组间的流量调度策略，取值有：
	// <li>Pritory：按优先级顺序进行故障转移。</li>默认值为 Pritory。
	SteeringPolicy *string `json:"SteeringPolicy,omitnil,omitempty" name:"SteeringPolicy"`

	// 实际访问某源站失败时的请求重试策略，详情请参考 [请求重试策略介绍](https://cloud.tencent.com/document/product/1552/104227)，取值有：
	// <li>OtherOriginGroup：单次请求失败后，请求优先重试下一优先级源站组；</li>
	// <li>OtherRecordInOriginGroup：单次请求失败后，请求优先重试同源站组内的其他源站。</li>默认值为 OtherRecordInOriginGroup。
	FailoverPolicy *string `json:"FailoverPolicy,omitnil,omitempty" name:"FailoverPolicy"`
}

type CreateLoadBalancerRequest struct {
	*tchttp.BaseRequest
	
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 实例名称，可输入 1-200 个字符，允许字符为 a-z，A-Z，0-9，_，-。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 实例类型，取值有：
	// <li>HTTP：HTTP 专用型，支持添加 HTTP 专用型和通用型源站组，仅支持被站点加速相关服务引用（如域名服务和规则引擎）；</li>
	// <li>GENERAL：通用型，仅支持添加通用型源站组，能被站点加速服务（如域名服务和规则引擎）和四层代理引用。</li>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 源站组列表及其对应的容灾调度优先级。详情请参考 [快速创建负载均衡实例](https://cloud.tencent.com/document/product/1552/104223) 中的示例场景。
	OriginGroups []*OriginGroupInLoadBalancer `json:"OriginGroups,omitnil,omitempty" name:"OriginGroups"`

	// 健康检查策略。详情请参考 [健康检查策略介绍](https://cloud.tencent.com/document/product/1552/104228)。不填写时，默认为不启用健康检查。
	HealthChecker *HealthChecker `json:"HealthChecker,omitnil,omitempty" name:"HealthChecker"`

	// 源站组间的流量调度策略，取值有：
	// <li>Pritory：按优先级顺序进行故障转移。</li>默认值为 Pritory。
	SteeringPolicy *string `json:"SteeringPolicy,omitnil,omitempty" name:"SteeringPolicy"`

	// 实际访问某源站失败时的请求重试策略，详情请参考 [请求重试策略介绍](https://cloud.tencent.com/document/product/1552/104227)，取值有：
	// <li>OtherOriginGroup：单次请求失败后，请求优先重试下一优先级源站组；</li>
	// <li>OtherRecordInOriginGroup：单次请求失败后，请求优先重试同源站组内的其他源站。</li>默认值为 OtherRecordInOriginGroup。
	FailoverPolicy *string `json:"FailoverPolicy,omitnil,omitempty" name:"FailoverPolicy"`
}

func (r *CreateLoadBalancerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateLoadBalancerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "Name")
	delete(f, "Type")
	delete(f, "OriginGroups")
	delete(f, "HealthChecker")
	delete(f, "SteeringPolicy")
	delete(f, "FailoverPolicy")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateLoadBalancerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateLoadBalancerResponseParams struct {
	// 负载均衡实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateLoadBalancerResponse struct {
	*tchttp.BaseResponse
	Response *CreateLoadBalancerResponseParams `json:"Response"`
}

func (r *CreateLoadBalancerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateLoadBalancerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateMultiPathGatewayLineRequestParams struct {
	// 站点 ID 。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 多通道安全网关 ID 。
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// 线路类型，取值有：
	//  <li>direct ：直连线路，不支持修改和删除。</li> <li>proxy ：EdgeOne 四层代理线路，支持修改实例 ID 和规则 ID，不支持删除。</li> <li>custom ：自定义线路，支持修改、删除实例 ID 和规则 ID。</li>
	LineType *string `json:"LineType,omitnil,omitempty" name:"LineType"`

	// 线路地址，格式为 ip:port。
	LineAddress *string `json:"LineAddress,omitnil,omitempty" name:"LineAddress"`

	// 四层代理实例 ID，当线路类型 LineType 取值为 proxy（EdgeOne 四层代理）必传，可由接口 [DescribeL4Proxy](https://cloud.tencent.com/document/api/1552/103413) 获取。
	ProxyId *string `json:"ProxyId,omitnil,omitempty" name:"ProxyId"`

	// 转发规则 ID ，当线路类型 LineType 取值为 proxy（EdgeOne 四层代理）必传，可以从接口 [DescribeL4ProxyRules](https://cloud.tencent.com/document/api/1552/103412) 获取。
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`
}

type CreateMultiPathGatewayLineRequest struct {
	*tchttp.BaseRequest
	
	// 站点 ID 。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 多通道安全网关 ID 。
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// 线路类型，取值有：
	//  <li>direct ：直连线路，不支持修改和删除。</li> <li>proxy ：EdgeOne 四层代理线路，支持修改实例 ID 和规则 ID，不支持删除。</li> <li>custom ：自定义线路，支持修改、删除实例 ID 和规则 ID。</li>
	LineType *string `json:"LineType,omitnil,omitempty" name:"LineType"`

	// 线路地址，格式为 ip:port。
	LineAddress *string `json:"LineAddress,omitnil,omitempty" name:"LineAddress"`

	// 四层代理实例 ID，当线路类型 LineType 取值为 proxy（EdgeOne 四层代理）必传，可由接口 [DescribeL4Proxy](https://cloud.tencent.com/document/api/1552/103413) 获取。
	ProxyId *string `json:"ProxyId,omitnil,omitempty" name:"ProxyId"`

	// 转发规则 ID ，当线路类型 LineType 取值为 proxy（EdgeOne 四层代理）必传，可以从接口 [DescribeL4ProxyRules](https://cloud.tencent.com/document/api/1552/103412) 获取。
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`
}

func (r *CreateMultiPathGatewayLineRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateMultiPathGatewayLineRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "GatewayId")
	delete(f, "LineType")
	delete(f, "LineAddress")
	delete(f, "ProxyId")
	delete(f, "RuleId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateMultiPathGatewayLineRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateMultiPathGatewayLineResponseParams struct {
	// 线路 ID ， 取值有:
	// <li> line-1： EdgeOne 四层代理线路，支持修改实例和规则，不支持删除；</li>
	// <li> line-2 及以上：EdgeOne 四层代理线路或者自定义线路，支持修改、删除实例和规则。</li>
	LineId *string `json:"LineId,omitnil,omitempty" name:"LineId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateMultiPathGatewayLineResponse struct {
	*tchttp.BaseResponse
	Response *CreateMultiPathGatewayLineResponseParams `json:"Response"`
}

func (r *CreateMultiPathGatewayLineResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateMultiPathGatewayLineResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateMultiPathGatewayRequestParams struct {
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 网关类型，取值有：
	// <li> cloud：云上网关，腾讯云创建和管理的网关；</li>
	// <li> private：自有网关，用户部署的私有网关。</li>
	GatewayType *string `json:"GatewayType,omitnil,omitempty" name:"GatewayType"`

	// 网关名称，16 个字符以内，可用字符（a-z,A-Z,0-9,-,_）。
	GatewayName *string `json:"GatewayName,omitnil,omitempty" name:"GatewayName"`

	// 网关端口，范围 1～65535（除去 8888 ）。
	GatewayPort *int64 `json:"GatewayPort,omitnil,omitempty" name:"GatewayPort"`

	// 网关地域，GatewayType 取值为 cloud（云上网关）必填。可以从接口 DescribeMultiPathGatewayRegions 获取 RegionId 列表。
	RegionId *string `json:"RegionId,omitnil,omitempty" name:"RegionId"`

	// 网关地址，GatewayType 取值为 private（自有网关）必填，使用该地址时，请确保该地址已录入腾讯云多通道安全加速网关系统。如未录入，需要在本接口调用前通过工单或者联系架构师把网关 IP 地址提前录入腾讯云多通道安全加速网关系统。
	GatewayIP *string `json:"GatewayIP,omitnil,omitempty" name:"GatewayIP"`
}

type CreateMultiPathGatewayRequest struct {
	*tchttp.BaseRequest
	
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 网关类型，取值有：
	// <li> cloud：云上网关，腾讯云创建和管理的网关；</li>
	// <li> private：自有网关，用户部署的私有网关。</li>
	GatewayType *string `json:"GatewayType,omitnil,omitempty" name:"GatewayType"`

	// 网关名称，16 个字符以内，可用字符（a-z,A-Z,0-9,-,_）。
	GatewayName *string `json:"GatewayName,omitnil,omitempty" name:"GatewayName"`

	// 网关端口，范围 1～65535（除去 8888 ）。
	GatewayPort *int64 `json:"GatewayPort,omitnil,omitempty" name:"GatewayPort"`

	// 网关地域，GatewayType 取值为 cloud（云上网关）必填。可以从接口 DescribeMultiPathGatewayRegions 获取 RegionId 列表。
	RegionId *string `json:"RegionId,omitnil,omitempty" name:"RegionId"`

	// 网关地址，GatewayType 取值为 private（自有网关）必填，使用该地址时，请确保该地址已录入腾讯云多通道安全加速网关系统。如未录入，需要在本接口调用前通过工单或者联系架构师把网关 IP 地址提前录入腾讯云多通道安全加速网关系统。
	GatewayIP *string `json:"GatewayIP,omitnil,omitempty" name:"GatewayIP"`
}

func (r *CreateMultiPathGatewayRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateMultiPathGatewayRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "GatewayType")
	delete(f, "GatewayName")
	delete(f, "GatewayPort")
	delete(f, "RegionId")
	delete(f, "GatewayIP")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateMultiPathGatewayRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateMultiPathGatewayResponseParams struct {
	// 网关 ID。
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateMultiPathGatewayResponse struct {
	*tchttp.BaseResponse
	Response *CreateMultiPathGatewayResponseParams `json:"Response"`
}

func (r *CreateMultiPathGatewayResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateMultiPathGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateMultiPathGatewaySecretKeyRequestParams struct {
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 多通道安全加速网关接入密钥，base64字符串，编码前字符串长度为 32-48 个字符，非必填，不填系统自动生成，可通过接口 DescribeMultiPathGatewaySecretKey 查询。
	SecretKey *string `json:"SecretKey,omitnil,omitempty" name:"SecretKey"`
}

type CreateMultiPathGatewaySecretKeyRequest struct {
	*tchttp.BaseRequest
	
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 多通道安全加速网关接入密钥，base64字符串，编码前字符串长度为 32-48 个字符，非必填，不填系统自动生成，可通过接口 DescribeMultiPathGatewaySecretKey 查询。
	SecretKey *string `json:"SecretKey,omitnil,omitempty" name:"SecretKey"`
}

func (r *CreateMultiPathGatewaySecretKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateMultiPathGatewaySecretKeyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "SecretKey")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateMultiPathGatewaySecretKeyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateMultiPathGatewaySecretKeyResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateMultiPathGatewaySecretKeyResponse struct {
	*tchttp.BaseResponse
	Response *CreateMultiPathGatewaySecretKeyResponseParams `json:"Response"`
}

func (r *CreateMultiPathGatewaySecretKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateMultiPathGatewaySecretKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateOriginGroupRequestParams struct {
	// 站点 ID
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 源站记录信息，此参数必填。
	Records []*OriginRecord `json:"Records,omitnil,omitempty" name:"Records"`

	// 源站组名称，可输入1 - 200个字符，允许的字符为 a - z, A - Z, 0 - 9, _, - 。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 源站组类型，此参数必填，取值有：
	// <li>GENERAL：通用型源站组，仅支持添加 IP/域名 源站，可以被域名服务、规则引擎、四层代理、通用型负载均衡、HTTP 专用型负载均衡引用；</li>
	// <li>HTTP： HTTP 专用型源站组，支持添加 IP/域名、对象存储源站作为源站，无法被四层代理引用，仅支持被添加加速域名、规则引擎-修改源站、HTTP 专用型负载均衡引用。</li>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 回源 Host Header，仅 Type = HTTP 时传入生效，规则引擎修改 Host Header 配置优先级高于源站组的 Host Header。
	HostHeader *string `json:"HostHeader,omitnil,omitempty" name:"HostHeader"`
}

type CreateOriginGroupRequest struct {
	*tchttp.BaseRequest
	
	// 站点 ID
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 源站记录信息，此参数必填。
	Records []*OriginRecord `json:"Records,omitnil,omitempty" name:"Records"`

	// 源站组名称，可输入1 - 200个字符，允许的字符为 a - z, A - Z, 0 - 9, _, - 。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 源站组类型，此参数必填，取值有：
	// <li>GENERAL：通用型源站组，仅支持添加 IP/域名 源站，可以被域名服务、规则引擎、四层代理、通用型负载均衡、HTTP 专用型负载均衡引用；</li>
	// <li>HTTP： HTTP 专用型源站组，支持添加 IP/域名、对象存储源站作为源站，无法被四层代理引用，仅支持被添加加速域名、规则引擎-修改源站、HTTP 专用型负载均衡引用。</li>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

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
	delete(f, "Records")
	delete(f, "Name")
	delete(f, "Type")
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
	// 若您希望快速提交不同站点下的 Targets Url，可以将其填写为 *，但前提是调用该 API 的账号必须具备主账号下全部站点资源的权限。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 要预热的资源列表，必填。每个元素格式类似如下:
	// http://www.example.com/example.txt。
	// 注意：提交任务数受计费套餐配额限制，请查看 [EO计费套餐](https://cloud.tencent.com/document/product/1552/77380)。
	Targets []*string `json:"Targets,omitnil,omitempty" name:"Targets"`

	// 预热模式，取值有：
	// <li>default：默认模式，即预热到中间层；</li>
	// <li>edge：边缘预热模式，即预热到边缘和中间层。</li>不填写时，默认值为 default。
	// 注意事项：
	// 1.预热至边缘产生的边缘层流量，会计入计费流量；
	// 2.边缘预热默认分配单独的预热额度 1000 条/天，不消费常规预热额度。
	// 说明：
	// 该参数为白名单功能，如有需要，请联系腾讯云工程师处理。
	Mode *string `json:"Mode,omitnil,omitempty" name:"Mode"`

	// 是否对url进行encode，若内容含有非 ASCII 字符集的字符，请开启此开关进行编码转换（编码规则遵循 RFC3986）。
	//
	// Deprecated: EncodeUrl is deprecated.
	EncodeUrl *bool `json:"EncodeUrl,omitnil,omitempty" name:"EncodeUrl"`

	// 若需要携带 HTTP 头部信息预热，可入参该参数，否则放空即可。
	Headers []*Header `json:"Headers,omitnil,omitempty" name:"Headers"`

	// 媒体分片预热控制，取值有：
	// <li>on：开启分片预热，预热描述文件，并递归解析描述文件分片进行预热；</li>
	// <li>off：仅预热提交的描述文件；</li>不填写时，默认值为 off。
	// 注意事项：
	// 1. 支持的描述文件为 M3U8，对应分片为 TS；
	// 2. 要求描述文件能正常请求，并按行业标准描述分片路径；
	// 3. 递归解析深度不超过 3 层；
	// 4. 解析获取的分片会正常累加每日预热用量，当用量超出配额时，会静默处理，不再执行预热。
	// 说明：
	// 该参数为白名单功能，如有需要，请联系腾讯云工程师处理。
	PrefetchMediaSegments *string `json:"PrefetchMediaSegments,omitnil,omitempty" name:"PrefetchMediaSegments"`
}

type CreatePrefetchTaskRequest struct {
	*tchttp.BaseRequest
	
	// 站点 ID。
	// 若您希望快速提交不同站点下的 Targets Url，可以将其填写为 *，但前提是调用该 API 的账号必须具备主账号下全部站点资源的权限。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 要预热的资源列表，必填。每个元素格式类似如下:
	// http://www.example.com/example.txt。
	// 注意：提交任务数受计费套餐配额限制，请查看 [EO计费套餐](https://cloud.tencent.com/document/product/1552/77380)。
	Targets []*string `json:"Targets,omitnil,omitempty" name:"Targets"`

	// 预热模式，取值有：
	// <li>default：默认模式，即预热到中间层；</li>
	// <li>edge：边缘预热模式，即预热到边缘和中间层。</li>不填写时，默认值为 default。
	// 注意事项：
	// 1.预热至边缘产生的边缘层流量，会计入计费流量；
	// 2.边缘预热默认分配单独的预热额度 1000 条/天，不消费常规预热额度。
	// 说明：
	// 该参数为白名单功能，如有需要，请联系腾讯云工程师处理。
	Mode *string `json:"Mode,omitnil,omitempty" name:"Mode"`

	// 是否对url进行encode，若内容含有非 ASCII 字符集的字符，请开启此开关进行编码转换（编码规则遵循 RFC3986）。
	EncodeUrl *bool `json:"EncodeUrl,omitnil,omitempty" name:"EncodeUrl"`

	// 若需要携带 HTTP 头部信息预热，可入参该参数，否则放空即可。
	Headers []*Header `json:"Headers,omitnil,omitempty" name:"Headers"`

	// 媒体分片预热控制，取值有：
	// <li>on：开启分片预热，预热描述文件，并递归解析描述文件分片进行预热；</li>
	// <li>off：仅预热提交的描述文件；</li>不填写时，默认值为 off。
	// 注意事项：
	// 1. 支持的描述文件为 M3U8，对应分片为 TS；
	// 2. 要求描述文件能正常请求，并按行业标准描述分片路径；
	// 3. 递归解析深度不超过 3 层；
	// 4. 解析获取的分片会正常累加每日预热用量，当用量超出配额时，会静默处理，不再执行预热。
	// 说明：
	// 该参数为白名单功能，如有需要，请联系腾讯云工程师处理。
	PrefetchMediaSegments *string `json:"PrefetchMediaSegments,omitnil,omitempty" name:"PrefetchMediaSegments"`
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
	delete(f, "Mode")
	delete(f, "EncodeUrl")
	delete(f, "Headers")
	delete(f, "PrefetchMediaSegments")
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
	// 若您希望快速提交不同站点下的 Targets Url，可以将其填写为 *，但前提是调用该 API 的账号必须具备主账号下全部站点资源的权限。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 节点缓存清除类型，取值有：
	// <li>purge_url：URL刷新；</li>
	// <li>purge_prefix：目录刷新；</li>
	// <li>purge_host：Hostname 刷新；</li>
	// <li>purge_all：站点下全部缓存刷新；</li>
	// <li>purge_cache_tag：cache-tag 刷新。</li>缓存清除类型详情请查看[清除缓存](https://cloud.tencent.com/document/product/1552/70759)。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 节点缓存清除方法，针对目录刷新、Hostname 刷新以及刷新全部缓存类型有效，取值有：<li> invalidate：仅刷新目录下产生了更新的资源；</li><li> delete：无论目录下资源是否更新都刷新节点资源。</li>默认值： invalidate。
	Method *string `json:"Method,omitnil,omitempty" name:"Method"`

	// 需清除缓存的资源列表，如 https://www.example.com/example.jpg，必须携带协议信息。更多元素格式依据清除缓存类型而定，可参考下方接口调用示例。<li>单次提交的任务数受计费套餐配额限制，请查看 [EO 计费套餐](https://cloud.tencent.com/document/product/1552/77380)。</li>
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
	// 若您希望快速提交不同站点下的 Targets Url，可以将其填写为 *，但前提是调用该 API 的账号必须具备主账号下全部站点资源的权限。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 节点缓存清除类型，取值有：
	// <li>purge_url：URL刷新；</li>
	// <li>purge_prefix：目录刷新；</li>
	// <li>purge_host：Hostname 刷新；</li>
	// <li>purge_all：站点下全部缓存刷新；</li>
	// <li>purge_cache_tag：cache-tag 刷新。</li>缓存清除类型详情请查看[清除缓存](https://cloud.tencent.com/document/product/1552/70759)。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 节点缓存清除方法，针对目录刷新、Hostname 刷新以及刷新全部缓存类型有效，取值有：<li> invalidate：仅刷新目录下产生了更新的资源；</li><li> delete：无论目录下资源是否更新都刷新节点资源。</li>默认值： invalidate。
	Method *string `json:"Method,omitnil,omitempty" name:"Method"`

	// 需清除缓存的资源列表，如 https://www.example.com/example.jpg，必须携带协议信息。更多元素格式依据清除缓存类型而定，可参考下方接口调用示例。<li>单次提交的任务数受计费套餐配额限制，请查看 [EO 计费套餐](https://cloud.tencent.com/document/product/1552/77380)。</li>
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
	// <li>s3：推送到 AWS S3 兼容存储桶地址；</li>
	// <li>log_analysis：推送到 EdgeOne 日志分析，该任务类型仅支持“站点加速日志”这一数据投递类型。</li>
	TaskType *string `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// 实时日志投递任务对应的实体列表。取值示例如下：
	// <li>七层域名：domain.example.com</li>
	// <li>四层代理实例：sid-2s69eb5wcms7</li>
	// <li>边缘函数实例：test-zone-2mxigizoh9l9-1257626257</li>
	EntityList []*string `json:"EntityList,omitnil,omitempty" name:"EntityList"`

	// 数据投递类型，取值有：
	// <li>domain：站点加速日志；</li>
	// <li>application：四层代理日志；</li>
	// <li>function：边缘函数运行日志；</li>
	// <li>web-rateLiming：速率限制和 CC 攻击防护日志；</li>
	// <li>web-attack：托管规则日志；</li>
	// <li>web-rule：自定义规则日志；</li>
	// <li>web-bot：Bot管理日志。</li>
	LogType *string `json:"LogType,omitnil,omitempty" name:"LogType"`

	// 数据投递区域，取值有：
	// <li>mainland：中国大陆境内；</li>
	// <li>overseas：全球（不含中国大陆）。</li>
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 投递的预设字段列表。取值参考：
	// <li>[站点加速日志（七层访问日志）](https://cloud.tencent.com/document/product/1552/105791)</li>
	// <li>[四层代理日志](https://cloud.tencent.com/document/product/1552/105792)</li>
	// <li>[边缘函数运行日志](https://cloud.tencent.com/document/product/1552/115585)</li>
	Fields []*string `json:"Fields,omitnil,omitempty" name:"Fields"`

	// 投递的自定义字段列表，支持在 HTTP 请求头、响应头、Cookie、请求正文中提取指定内容。自定义字段名称不能重复，且最多不能超过 200 个字段。单个实时日志推送任务最多添加 5 个请求正文类型的自定义字段。目前仅站点加速日志（LogType=domain）支持添加自定义字段。
	CustomFields []*CustomField `json:"CustomFields,omitnil,omitempty" name:"CustomFields"`

	// 日志投递的过滤条件，不填表示投递全量日志。
	DeliveryConditions []*DeliveryCondition `json:"DeliveryConditions,omitnil,omitempty" name:"DeliveryConditions"`

	// 采样比例，采用千分制，取值范围为1-1000，例如：填写 605 表示采样比例为 60.5%。不填表示采样比例为 100%。
	Sample *uint64 `json:"Sample,omitnil,omitempty" name:"Sample"`

	// 日志投递的输出格式。不填表示为默认格式，默认格式逻辑如下：
	// <li>当 TaskType 取值为 custom_endpoint 时，默认格式为多个 JSON 对象组成的数组，每个 JSON 对象为一条日志；</li>
	// <li>当 TaskType 取值为 s3 时，默认格式为 JSON Lines；</li>特别地，当 TaskType 取值为 cls 或 log_analysis 时，LogFormat.FormatType 的值只能为 json，且 LogFormat 中其他参数将被忽略，建议不传 LogFormat。
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
	// <li>s3：推送到 AWS S3 兼容存储桶地址；</li>
	// <li>log_analysis：推送到 EdgeOne 日志分析，该任务类型仅支持“站点加速日志”这一数据投递类型。</li>
	TaskType *string `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// 实时日志投递任务对应的实体列表。取值示例如下：
	// <li>七层域名：domain.example.com</li>
	// <li>四层代理实例：sid-2s69eb5wcms7</li>
	// <li>边缘函数实例：test-zone-2mxigizoh9l9-1257626257</li>
	EntityList []*string `json:"EntityList,omitnil,omitempty" name:"EntityList"`

	// 数据投递类型，取值有：
	// <li>domain：站点加速日志；</li>
	// <li>application：四层代理日志；</li>
	// <li>function：边缘函数运行日志；</li>
	// <li>web-rateLiming：速率限制和 CC 攻击防护日志；</li>
	// <li>web-attack：托管规则日志；</li>
	// <li>web-rule：自定义规则日志；</li>
	// <li>web-bot：Bot管理日志。</li>
	LogType *string `json:"LogType,omitnil,omitempty" name:"LogType"`

	// 数据投递区域，取值有：
	// <li>mainland：中国大陆境内；</li>
	// <li>overseas：全球（不含中国大陆）。</li>
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 投递的预设字段列表。取值参考：
	// <li>[站点加速日志（七层访问日志）](https://cloud.tencent.com/document/product/1552/105791)</li>
	// <li>[四层代理日志](https://cloud.tencent.com/document/product/1552/105792)</li>
	// <li>[边缘函数运行日志](https://cloud.tencent.com/document/product/1552/115585)</li>
	Fields []*string `json:"Fields,omitnil,omitempty" name:"Fields"`

	// 投递的自定义字段列表，支持在 HTTP 请求头、响应头、Cookie、请求正文中提取指定内容。自定义字段名称不能重复，且最多不能超过 200 个字段。单个实时日志推送任务最多添加 5 个请求正文类型的自定义字段。目前仅站点加速日志（LogType=domain）支持添加自定义字段。
	CustomFields []*CustomField `json:"CustomFields,omitnil,omitempty" name:"CustomFields"`

	// 日志投递的过滤条件，不填表示投递全量日志。
	DeliveryConditions []*DeliveryCondition `json:"DeliveryConditions,omitnil,omitempty" name:"DeliveryConditions"`

	// 采样比例，采用千分制，取值范围为1-1000，例如：填写 605 表示采样比例为 60.5%。不填表示采样比例为 100%。
	Sample *uint64 `json:"Sample,omitnil,omitempty" name:"Sample"`

	// 日志投递的输出格式。不填表示为默认格式，默认格式逻辑如下：
	// <li>当 TaskType 取值为 custom_endpoint 时，默认格式为多个 JSON 对象组成的数组，每个 JSON 对象为一条日志；</li>
	// <li>当 TaskType 取值为 s3 时，默认格式为 JSON Lines；</li>特别地，当 TaskType 取值为 cls 或 log_analysis 时，LogFormat.FormatType 的值只能为 json，且 LogFormat 中其他参数将被忽略，建议不传 LogFormat。
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
type CreateSecurityAPIResourceRequestParams struct {
	// 站点 ID。	
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	//  API 资源列表。
	APIResources []*APIResource `json:"APIResources,omitnil,omitempty" name:"APIResources"`
}

type CreateSecurityAPIResourceRequest struct {
	*tchttp.BaseRequest
	
	// 站点 ID。	
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	//  API 资源列表。
	APIResources []*APIResource `json:"APIResources,omitnil,omitempty" name:"APIResources"`
}

func (r *CreateSecurityAPIResourceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSecurityAPIResourceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "APIResources")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSecurityAPIResourceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSecurityAPIResourceResponseParams struct {
	// API 资源 ID 列表。
	APIResourceIds []*string `json:"APIResourceIds,omitnil,omitempty" name:"APIResourceIds"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateSecurityAPIResourceResponse struct {
	*tchttp.BaseResponse
	Response *CreateSecurityAPIResourceResponseParams `json:"Response"`
}

func (r *CreateSecurityAPIResourceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSecurityAPIResourceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSecurityAPIServiceRequestParams struct {
	// 站点 ID。	
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	//  API 服务列表。
	APIServices []*APIService `json:"APIServices,omitnil,omitempty" name:"APIServices"`
}

type CreateSecurityAPIServiceRequest struct {
	*tchttp.BaseRequest
	
	// 站点 ID。	
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	//  API 服务列表。
	APIServices []*APIService `json:"APIServices,omitnil,omitempty" name:"APIServices"`
}

func (r *CreateSecurityAPIServiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSecurityAPIServiceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "APIServices")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSecurityAPIServiceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSecurityAPIServiceResponseParams struct {
	// API 服务 ID 列表。
	APIServiceIds []*string `json:"APIServiceIds,omitnil,omitempty" name:"APIServiceIds"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateSecurityAPIServiceResponse struct {
	*tchttp.BaseResponse
	Response *CreateSecurityAPIServiceResponseParams `json:"Response"`
}

func (r *CreateSecurityAPIServiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSecurityAPIServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSecurityClientAttesterRequestParams struct {
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 认证选项列表。
	ClientAttesters []*ClientAttester `json:"ClientAttesters,omitnil,omitempty" name:"ClientAttesters"`
}

type CreateSecurityClientAttesterRequest struct {
	*tchttp.BaseRequest
	
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 认证选项列表。
	ClientAttesters []*ClientAttester `json:"ClientAttesters,omitnil,omitempty" name:"ClientAttesters"`
}

func (r *CreateSecurityClientAttesterRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSecurityClientAttesterRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "ClientAttesters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSecurityClientAttesterRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSecurityClientAttesterResponseParams struct {
	// 认证选项 ID 列表。
	ClientAttesterIds []*string `json:"ClientAttesterIds,omitnil,omitempty" name:"ClientAttesterIds"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateSecurityClientAttesterResponse struct {
	*tchttp.BaseResponse
	Response *CreateSecurityClientAttesterResponseParams `json:"Response"`
}

func (r *CreateSecurityClientAttesterResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSecurityClientAttesterResponse) FromJsonString(s string) error {
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
type CreateSecurityJSInjectionRuleRequestParams struct {
	// 站点 ID。	
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// JavaScript 注入规则列表。
	JSInjectionRules []*JSInjectionRule `json:"JSInjectionRules,omitnil,omitempty" name:"JSInjectionRules"`
}

type CreateSecurityJSInjectionRuleRequest struct {
	*tchttp.BaseRequest
	
	// 站点 ID。	
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// JavaScript 注入规则列表。
	JSInjectionRules []*JSInjectionRule `json:"JSInjectionRules,omitnil,omitempty" name:"JSInjectionRules"`
}

func (r *CreateSecurityJSInjectionRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSecurityJSInjectionRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "JSInjectionRules")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSecurityJSInjectionRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSecurityJSInjectionRuleResponseParams struct {
	// JavaScript 注入规则 ID 列表。
	JSInjectionRuleIds []*string `json:"JSInjectionRuleIds,omitnil,omitempty" name:"JSInjectionRuleIds"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateSecurityJSInjectionRuleResponse struct {
	*tchttp.BaseResponse
	Response *CreateSecurityJSInjectionRuleResponseParams `json:"Response"`
}

func (r *CreateSecurityJSInjectionRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSecurityJSInjectionRuleResponse) FromJsonString(s string) error {
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
type CreateWebSecurityTemplateRequestParams struct {
	// 站点 ID。该参数明确策略模板在访问权限上归属的站点。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 策略模板名称。由中文、英文、数字和下划线组成，不能以下划线开头，且长度不能超过 32 个字符。
	TemplateName *string `json:"TemplateName,omitnil,omitempty" name:"TemplateName"`

	// 安全策略模板配置内容，字段为空时生成默认配置。目前支持 Web 防护模块中的例外规则、自定义规则、速率限制规则和托管规则配置，通过表达式语法对安全策略进行配置。 Bot 管理规则配置暂不支持，正在开发中。
	SecurityPolicy *SecurityPolicy `json:"SecurityPolicy,omitnil,omitempty" name:"SecurityPolicy"`
}

type CreateWebSecurityTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 站点 ID。该参数明确策略模板在访问权限上归属的站点。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 策略模板名称。由中文、英文、数字和下划线组成，不能以下划线开头，且长度不能超过 32 个字符。
	TemplateName *string `json:"TemplateName,omitnil,omitempty" name:"TemplateName"`

	// 安全策略模板配置内容，字段为空时生成默认配置。目前支持 Web 防护模块中的例外规则、自定义规则、速率限制规则和托管规则配置，通过表达式语法对安全策略进行配置。 Bot 管理规则配置暂不支持，正在开发中。
	SecurityPolicy *SecurityPolicy `json:"SecurityPolicy,omitnil,omitempty" name:"SecurityPolicy"`
}

func (r *CreateWebSecurityTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateWebSecurityTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "TemplateName")
	delete(f, "SecurityPolicy")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateWebSecurityTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateWebSecurityTemplateResponseParams struct {
	// 策略模板 ID。
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateWebSecurityTemplateResponse struct {
	*tchttp.BaseResponse
	Response *CreateWebSecurityTemplateResponseParams `json:"Response"`
}

func (r *CreateWebSecurityTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateWebSecurityTemplateResponse) FromJsonString(s string) error {
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

	// 待绑定的目标套餐 ID。当您账号下已存在套餐时，可以填写此参数，直接将站点绑定至该套餐。若您当前没有可绑定的套餐时，可通过 [CreatePlan](https://cloud.tencent.com/document/product/1552/105771) 购买套餐。
	// 注意：如果不填写此参数，将创建一个处于“init”状态的站点，该站点为未激活状态，并不会显示在控制台上。您可以通过访问 [BindZoneToPlan](https://cloud.tencent.com/document/product/1552/83042) 来绑定套餐并激活站点，激活后站点可以正常提供服务。
	// 
	PlanId *string `json:"PlanId,omitnil,omitempty" name:"PlanId"`

	// 同名站点标识。限制输入数字、英文、"." 、"-" 和 "_"，长度 200 个字符以内。详情参考 [同名站点标识](https://cloud.tencent.com/document/product/1552/70202)，无此使用场景时，该字段保留为空即可。
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

	// 待绑定的目标套餐 ID。当您账号下已存在套餐时，可以填写此参数，直接将站点绑定至该套餐。若您当前没有可绑定的套餐时，可通过 [CreatePlan](https://cloud.tencent.com/document/product/1552/105771) 购买套餐。
	// 注意：如果不填写此参数，将创建一个处于“init”状态的站点，该站点为未激活状态，并不会显示在控制台上。您可以通过访问 [BindZoneToPlan](https://cloud.tencent.com/document/product/1552/83042) 来绑定套餐并激活站点，激活后站点可以正常提供服务。
	// 
	PlanId *string `json:"PlanId,omitnil,omitempty" name:"PlanId"`

	// 同名站点标识。限制输入数字、英文、"." 、"-" 和 "_"，长度 200 个字符以内。详情参考 [同名站点标识](https://cloud.tencent.com/document/product/1552/70202)，无此使用场景时，该字段保留为空即可。
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

type CurrentOriginACL struct {
	// 回源 IP 网段详情。
	// 注意：此字段可能返回 null，表示取不到有效值。
	EntireAddresses *Addresses `json:"EntireAddresses,omitnil,omitempty" name:"EntireAddresses"`

	// 版本号。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Version *string `json:"Version,omitnil,omitempty" name:"Version"`

	// 版本生效时间，时间是北京时间 UTC+8， 遵循 ISO 8601 标准的日期和时间格式。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ActiveTime *string `json:"ActiveTime,omitnil,omitempty" name:"ActiveTime"`

	// 本参数用于记录当前版本生效前是否完成「我已更新至最新回源 IP 网段」的确认。取值有：
	// <li>true：版本生效时，已完成更新至最新回源 IP 的确认；</li>
	// <li>false：版本生效时，仍未完成已更新至最新回源 IP 的确认，回源 IP 网段由后台强制更新至最新版本。</li>注意：本参数返回 false 时，请及时确认您的源站防火墙配置是否已更新至最新的回源 IP 网段，以避免出现回源失败。
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsPlaned *string `json:"IsPlaned,omitnil,omitempty" name:"IsPlaned"`
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
	// 自定义日志字段类型。从 HTTP 请求和响应中的指定位置提取数据，取值有：
	// <li>ReqHeader：从 HTTP 请求头中提取指定字段值；</li>
	// <li>RspHeader：从 HTTP 响应头中提取指定字段值；</li>
	// <li>Cookie: 从 Cookie 中提取指定字段值；</li>
	// <li>ReqBody: 从 HTTP 请求正文中通过 Google RE2 正则表达式提取指定内容。</li>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 根据字段类型（Name）填入字段值的定义。需要区分大小写。
	// <li>当字段类型为 ReqHeader、RspHeader、Cookie 时，填入需要提取值的参数名称，例如：Accept-Language。可输入 1-100 个字符，允许的字符开头为字母，中间为字母、数字、-，结尾为字母、数字；</li>
	// <li>当字段类型为 ReqBody 时，填入 Google RE2 正则表达式，正则表达式长度上限为 4KB。</li>
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`

	// 是否投递该字段，不填表示不投递此字段。
	Enabled *bool `json:"Enabled,omitnil,omitempty" name:"Enabled"`
}

type CustomRule struct {
	// 自定义规则的名称。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 自定义规则的具体内容，需符合表达式语法，详细规范参见产品文档。
	Condition *string `json:"Condition,omitnil,omitempty" name:"Condition"`

	// 自定义规则的执行动作。	SecurityAction 的 Name 取值支持：<li>Deny：拦截；</li><li>Monitor：观察；</li><li>ReturnCustomPage：使用指定页面拦截；</li><li>Redirect：重定向至 URL；</li><li>BlockIP：IP 封禁；</li><li>JSChallenge：JavaScript 挑战；</li><li>ManagedChallenge：托管挑战；</li><li>Allow：放行。</li>
	Action *SecurityAction `json:"Action,omitnil,omitempty" name:"Action"`

	// 自定义规则是否开启。取值有：<li>on：开启</li><li>off：关闭</li>
	Enabled *string `json:"Enabled,omitnil,omitempty" name:"Enabled"`

	// 自定义规则的 ID。<br>通过规则 ID 可支持不同的规则配置操作：<br> - 增加新规则：ID 为空或不指定 ID 参数；<br> - 修改已有规则：指定需要更新/修改的规则 ID；<br> - 删除已有规则：CustomRules 参数中，Rules 列表中未包含的已有规则将被删除。
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 自定义规则的类型。取值有：<li>BasicAccessRule：基础访问管控；</li><li>PreciseMatchRule：精准匹配规则，默认；</li><li>ManagedAccessRule：专家定制规则，仅出参。</li><br/>默认为PreciseMatchRule。
	RuleType *string `json:"RuleType,omitnil,omitempty" name:"RuleType"`

	// 自定义规则的优先级，范围是 0 ~ 100，默认为 0，仅支持精准匹配规则（PreciseMatchRule）。
	Priority *int64 `json:"Priority,omitnil,omitempty" name:"Priority"`
}

type CustomRules struct {
	// 自定义规则的定义列表。<br>使用 ModifySecurityPolicy 修改 Web 防护配置时: <br> -  若未指定 Rules 参数，或 Rules 参数长度为零：清空所有自定义规则配置。<br> - 若 SecurityPolicy 参数中，未指定 CustomRules 参数值：保持已有自定义规则配置，不做修改。
	Rules []*CustomRule `json:"Rules,omitnil,omitempty" name:"Rules"`
}

type CustomTime struct {
	// 自定义缓存时间开关，取值有：
	// <li>on：开启；</li>
	// <li>off：关闭。</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// 忽略源站 CacheControl 开关，取值有：
	// <li>on：开启；</li>
	// <li>off：关闭。</li>注意：当 Switch 为 on 时，此字段必填；当 Switch 为 off 时，无需填写此字段，若填写则不生效。
	IgnoreCacheControl *string `json:"IgnoreCacheControl,omitnil,omitempty" name:"IgnoreCacheControl"`

	// 自定义缓存时间数值，单位为秒，取值：0～315360000。<br>注意：当 Switch 为 on 时，此字段必填；当 Switch 为 off 时，无需填写此字段，若填写则不生效。
	CacheTime *int64 `json:"CacheTime,omitnil,omitempty" name:"CacheTime"`
}

type CustomizedHeader struct {
	// 自定义头部 Key。
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 自定义头部 Value。
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
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

type DDoSProtection struct {
	// 指定独立 DDoS 的防护范围。取值为：
	// <li> protect_all_domains：独立 DDoS 防护对站点内全部域名生效，新接入域名自动开启独立 DDoS 防护，入参为 protect_all_domains 时，入参 DomainDDoSProtections 不作处理；</li>
	// <li> protect_specified_domains：仅对指定域名生效，具体范围可通过 DomainDDoSProtection 参数指定。</li>
	ProtectionOption *string `json:"ProtectionOption,omitnil,omitempty" name:"ProtectionOption"`

	// 域名的独立 DDoS 防护配置。在入参场景中：
	// <li> 当 ProtectionOption 保持为 protect_specified_domains 时：未填写的域名维持原有独立 DDoS 防护配置不变，显式指定的域名​按传入参数更新；</li>
	// <li> 当 ProtectionOption 由 protect_all_domains 切换为 protect_specified_domains 时：若 DomainDDoSProtections 传空，停用站点下全部域名的独立 DDoS 防护；若 DomainDDoSProtections 不为空，参数中指定的域名停用或保持独立 DDoS 防护，其余未列出的域名统一停用独立 DDoS 防护。</li>
	DomainDDoSProtections []*DomainDDoSProtection `json:"DomainDDoSProtections,omitnil,omitempty" name:"DomainDDoSProtections"`

	// 共享 CNAME 的独立 DDoS 防护配置。仅作为出参使用。
	SharedCNAMEDDoSProtections []*DomainDDoSProtection `json:"SharedCNAMEDDoSProtections,omitnil,omitempty" name:"SharedCNAMEDDoSProtections"`
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

type DNSPodDetail struct {
	// 是否伪站点，取值有：
	// <li> 0：非伪站点；</li>
	// <li> 1：伪站点。</li>
	IsFake *int64 `json:"IsFake,omitnil,omitempty" name:"IsFake"`
}

type DefaultServerCertInfo struct {
	// 服务器证书 ID。
	CertId *string `json:"CertId,omitnil,omitempty" name:"CertId"`

	// 证书备注名。
	Alias *string `json:"Alias,omitnil,omitempty" name:"Alias"`

	// 证书类型，取值有：
	// <li>default: 默认证书;</li>
	// <li>upload:用户上传;</li>
	// <li>managed:腾讯云托管。</li>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 证书过期时间。
	ExpireTime *string `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`

	// 证书生效时间。
	EffectiveTime *string `json:"EffectiveTime,omitnil,omitempty" name:"EffectiveTime"`

	// 证书公用名。
	CommonName *string `json:"CommonName,omitnil,omitempty" name:"CommonName"`

	// 证书SAN域名。
	SubjectAltName []*string `json:"SubjectAltName,omitnil,omitempty" name:"SubjectAltName"`

	// 部署状态，取值有：
	// <li>processing: 部署中；</li>
	// <li>deployed: 已部署；</li>
	// <li>failed: 部署失败。</li>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Status为失败时,此字段返回失败原因。
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 证书算法。
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
type DeleteContentIdentifierRequestParams struct {
	// 内容标识符 ID。
	ContentId *string `json:"ContentId,omitnil,omitempty" name:"ContentId"`
}

type DeleteContentIdentifierRequest struct {
	*tchttp.BaseRequest
	
	// 内容标识符 ID。
	ContentId *string `json:"ContentId,omitnil,omitempty" name:"ContentId"`
}

func (r *DeleteContentIdentifierRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteContentIdentifierRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ContentId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteContentIdentifierRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteContentIdentifierResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteContentIdentifierResponse struct {
	*tchttp.BaseResponse
	Response *DeleteContentIdentifierResponseParams `json:"Response"`
}

func (r *DeleteContentIdentifierResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteContentIdentifierResponse) FromJsonString(s string) error {
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
type DeleteDnsRecordsRequestParams struct {
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 待删除的 DNS 记录 ID 列表，上限：1000。
	RecordIds []*string `json:"RecordIds,omitnil,omitempty" name:"RecordIds"`
}

type DeleteDnsRecordsRequest struct {
	*tchttp.BaseRequest
	
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 待删除的 DNS 记录 ID 列表，上限：1000。
	RecordIds []*string `json:"RecordIds,omitnil,omitempty" name:"RecordIds"`
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
	delete(f, "RecordIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteDnsRecordsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDnsRecordsResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DeleteJustInTimeTranscodeTemplatesRequestParams struct {
	// 站点ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 需删除的即时转码模板唯一标识数组，长度限制：100。
	TemplateIds []*string `json:"TemplateIds,omitnil,omitempty" name:"TemplateIds"`
}

type DeleteJustInTimeTranscodeTemplatesRequest struct {
	*tchttp.BaseRequest
	
	// 站点ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 需删除的即时转码模板唯一标识数组，长度限制：100。
	TemplateIds []*string `json:"TemplateIds,omitnil,omitempty" name:"TemplateIds"`
}

func (r *DeleteJustInTimeTranscodeTemplatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteJustInTimeTranscodeTemplatesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "TemplateIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteJustInTimeTranscodeTemplatesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteJustInTimeTranscodeTemplatesResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteJustInTimeTranscodeTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *DeleteJustInTimeTranscodeTemplatesResponseParams `json:"Response"`
}

func (r *DeleteJustInTimeTranscodeTemplatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteJustInTimeTranscodeTemplatesResponse) FromJsonString(s string) error {
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
type DeleteL7AccRulesRequestParams struct {
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 需要删除的规则 ID 列表。您可以通过 DescribeL7AccRules 获取 Ruleid。
	RuleIds []*string `json:"RuleIds,omitnil,omitempty" name:"RuleIds"`
}

type DeleteL7AccRulesRequest struct {
	*tchttp.BaseRequest
	
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 需要删除的规则 ID 列表。您可以通过 DescribeL7AccRules 获取 Ruleid。
	RuleIds []*string `json:"RuleIds,omitnil,omitempty" name:"RuleIds"`
}

func (r *DeleteL7AccRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteL7AccRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "RuleIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteL7AccRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteL7AccRulesResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteL7AccRulesResponse struct {
	*tchttp.BaseResponse
	Response *DeleteL7AccRulesResponseParams `json:"Response"`
}

func (r *DeleteL7AccRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteL7AccRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteLoadBalancerRequestParams struct {
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 负载均衡实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DeleteLoadBalancerRequest struct {
	*tchttp.BaseRequest
	
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 负载均衡实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DeleteLoadBalancerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteLoadBalancerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteLoadBalancerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteLoadBalancerResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteLoadBalancerResponse struct {
	*tchttp.BaseResponse
	Response *DeleteLoadBalancerResponseParams `json:"Response"`
}

func (r *DeleteLoadBalancerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteLoadBalancerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteMultiPathGatewayLineRequestParams struct {
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 网关 ID。
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// 线路 ID。
	LineId *string `json:"LineId,omitnil,omitempty" name:"LineId"`
}

type DeleteMultiPathGatewayLineRequest struct {
	*tchttp.BaseRequest
	
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 网关 ID。
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// 线路 ID。
	LineId *string `json:"LineId,omitnil,omitempty" name:"LineId"`
}

func (r *DeleteMultiPathGatewayLineRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteMultiPathGatewayLineRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "GatewayId")
	delete(f, "LineId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteMultiPathGatewayLineRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteMultiPathGatewayLineResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteMultiPathGatewayLineResponse struct {
	*tchttp.BaseResponse
	Response *DeleteMultiPathGatewayLineResponseParams `json:"Response"`
}

func (r *DeleteMultiPathGatewayLineResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteMultiPathGatewayLineResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteMultiPathGatewayRequestParams struct {
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 网关 ID。
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`
}

type DeleteMultiPathGatewayRequest struct {
	*tchttp.BaseRequest
	
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 网关 ID。
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`
}

func (r *DeleteMultiPathGatewayRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteMultiPathGatewayRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "GatewayId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteMultiPathGatewayRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteMultiPathGatewayResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteMultiPathGatewayResponse struct {
	*tchttp.BaseResponse
	Response *DeleteMultiPathGatewayResponseParams `json:"Response"`
}

func (r *DeleteMultiPathGatewayResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteMultiPathGatewayResponse) FromJsonString(s string) error {
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
type DeleteSecurityAPIResourceRequestParams struct {
	// 站点 ID。	
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 需要删除的 API 资源 ID 列表。
	APIResourceIds []*string `json:"APIResourceIds,omitnil,omitempty" name:"APIResourceIds"`
}

type DeleteSecurityAPIResourceRequest struct {
	*tchttp.BaseRequest
	
	// 站点 ID。	
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 需要删除的 API 资源 ID 列表。
	APIResourceIds []*string `json:"APIResourceIds,omitnil,omitempty" name:"APIResourceIds"`
}

func (r *DeleteSecurityAPIResourceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSecurityAPIResourceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "APIResourceIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteSecurityAPIResourceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSecurityAPIResourceResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteSecurityAPIResourceResponse struct {
	*tchttp.BaseResponse
	Response *DeleteSecurityAPIResourceResponseParams `json:"Response"`
}

func (r *DeleteSecurityAPIResourceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSecurityAPIResourceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSecurityAPIServiceRequestParams struct {
	// 站点 ID。	
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// API 服务 ID 列表。
	APIServiceIds []*string `json:"APIServiceIds,omitnil,omitempty" name:"APIServiceIds"`
}

type DeleteSecurityAPIServiceRequest struct {
	*tchttp.BaseRequest
	
	// 站点 ID。	
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// API 服务 ID 列表。
	APIServiceIds []*string `json:"APIServiceIds,omitnil,omitempty" name:"APIServiceIds"`
}

func (r *DeleteSecurityAPIServiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSecurityAPIServiceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "APIServiceIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteSecurityAPIServiceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSecurityAPIServiceResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteSecurityAPIServiceResponse struct {
	*tchttp.BaseResponse
	Response *DeleteSecurityAPIServiceResponseParams `json:"Response"`
}

func (r *DeleteSecurityAPIServiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSecurityAPIServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSecurityClientAttesterRequestParams struct {
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 客户端认证选项 ID。
	ClientAttesterIds []*string `json:"ClientAttesterIds,omitnil,omitempty" name:"ClientAttesterIds"`
}

type DeleteSecurityClientAttesterRequest struct {
	*tchttp.BaseRequest
	
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 客户端认证选项 ID。
	ClientAttesterIds []*string `json:"ClientAttesterIds,omitnil,omitempty" name:"ClientAttesterIds"`
}

func (r *DeleteSecurityClientAttesterRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSecurityClientAttesterRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "ClientAttesterIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteSecurityClientAttesterRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSecurityClientAttesterResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteSecurityClientAttesterResponse struct {
	*tchttp.BaseResponse
	Response *DeleteSecurityClientAttesterResponseParams `json:"Response"`
}

func (r *DeleteSecurityClientAttesterResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSecurityClientAttesterResponse) FromJsonString(s string) error {
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
type DeleteSecurityJSInjectionRuleRequestParams struct {
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// JavaScript 注入规则 ID 列表。
	JSInjectionRuleIds []*string `json:"JSInjectionRuleIds,omitnil,omitempty" name:"JSInjectionRuleIds"`
}

type DeleteSecurityJSInjectionRuleRequest struct {
	*tchttp.BaseRequest
	
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// JavaScript 注入规则 ID 列表。
	JSInjectionRuleIds []*string `json:"JSInjectionRuleIds,omitnil,omitempty" name:"JSInjectionRuleIds"`
}

func (r *DeleteSecurityJSInjectionRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSecurityJSInjectionRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "JSInjectionRuleIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteSecurityJSInjectionRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSecurityJSInjectionRuleResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteSecurityJSInjectionRuleResponse struct {
	*tchttp.BaseResponse
	Response *DeleteSecurityJSInjectionRuleResponseParams `json:"Response"`
}

func (r *DeleteSecurityJSInjectionRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSecurityJSInjectionRuleResponse) FromJsonString(s string) error {
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
type DeleteWebSecurityTemplateRequestParams struct {
	// 站点 ID。需要传入目标策略模板在访问权限上归属的站点，可使用 DescribeWebSecurityTemplates 接口查询策略模板归属的站点。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 策略模板 ID。
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`
}

type DeleteWebSecurityTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 站点 ID。需要传入目标策略模板在访问权限上归属的站点，可使用 DescribeWebSecurityTemplates 接口查询策略模板归属的站点。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 策略模板 ID。
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`
}

func (r *DeleteWebSecurityTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteWebSecurityTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "TemplateId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteWebSecurityTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteWebSecurityTemplateResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteWebSecurityTemplateResponse struct {
	*tchttp.BaseResponse
	Response *DeleteWebSecurityTemplateResponseParams `json:"Response"`
}

func (r *DeleteWebSecurityTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteWebSecurityTemplateResponse) FromJsonString(s string) error {
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

type DenyActionParameters struct {
	// 是否对来源 IP 延长封禁。取值有：
	// <li>on：开启；</li>
	// <li>off：关闭。</li>
	// 启用后，对触发规则的客户端 IP 持续拦截。当启用该选项时，必须同时指定 BlockIpDuration 参数。
	// 注意：该选项不可与 ReturnCustomPage 或 Stall 选项同时启用。
	BlockIp *string `json:"BlockIp,omitnil,omitempty" name:"BlockIp"`

	// 当 BlockIP 为 on 时IP 的封禁时长。
	BlockIpDuration *string `json:"BlockIpDuration,omitnil,omitempty" name:"BlockIpDuration"`

	// 是否使用自定义页面。取值有：
	// <li>on：开启；</li>
	// <li>off：关闭。</li>
	// 启用后，使用自定义页面内容拦截（响应）请求，当启用该选项时，必须同时指定 ResponseCode 和 ErrorPageId 参数。
	// 注意：该选项不可与 BlockIp 或 Stall 选项同时启用。
	ReturnCustomPage *string `json:"ReturnCustomPage,omitnil,omitempty" name:"ReturnCustomPage"`

	// 自定义页面的状态码。
	ResponseCode *string `json:"ResponseCode,omitnil,omitempty" name:"ResponseCode"`

	// 自定义页面的PageId。
	ErrorPageId *string `json:"ErrorPageId,omitnil,omitempty" name:"ErrorPageId"`

	// 是否对请求来源挂起不予处理。取值有：
	// <li>on：开启；</li>
	// <li>off：关闭。</li>
	// 启用后，不再响应当前连接会话内请求，且不会主动断开连接。用于爬虫对抗时，消耗客户端连接资源。
	// 注意：该选项不可与 BlockIp 或 ReturnCustomPage 选项同时启用。
	Stall *string `json:"Stall,omitnil,omitempty" name:"Stall"`
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

	// 结束时间。查询时间范围（`EndTime` - `StartTime`）需小于等于 31 天。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 站点 ID 集合，此参数必填。最多传入 100 个站点 ID。若需查询腾讯云主账号下所有站点数据，请用 `*` 代替，查询账号级别数据需具备本接口全部站点资源权限。
	ZoneIds []*string `json:"ZoneIds,omitnil,omitempty" name:"ZoneIds"`

	// 指标列表，取值如下：
	// <b>四/七层加速流量：</b><li>acc_flux: 内容加速流量，单位为 Byte；</li><li>smt_flux: 智能加速流量，单位为 Byte；</li><li>l4_flux: 四层加速流量，单位为 Byte；</li><li>sec_flux: 独立防护流量，单位为 Byte；</li><li>zxctg_flux: 中国大陆网络优化流量，单位为 Byte。</li><br><b>四/七层加速带宽:</b><li>acc_bandwidth: 内容加速带宽，单位为 bps；</li><li>smt_bandwidth: 智能加速带宽，单位为 bps；</li><li>l4_bandwidth: 四层加速带宽，单位为 bps；</li><li>sec_bandwidth: 独立防护带宽，单位为 bps；</li><li>zxctg_bandwidth: 中国大陆网络优化带宽，单位为 bps。</li><br><b>HTTP/HTTPS 安全请求数：</b><li>sec_request_clean: HTTP/HTTPS 请求，单位为次。</li><b><br>增值服务用量：</b><li>smt_request_clean: 智能加速请求，单位为次；</li><li>quic_request: QUIC 请求，单位为次；</li><li>bot_request_clean: Bot 请求，单位为次；</li><li>cls_count: 实时日志推送条数，单位为条；</li><li>ddos_bandwidth: 弹性 DDoS 防护带宽，单位为 bps。</li><br><b>边缘计算用量：</b><li>edgefunction_request：边缘函数请求数，单位为次；</li><li>edgefunction_cpu_time：边缘函数CPU处理时间，单位为毫秒。</li>
	// <b>媒体处理用量：</b><li>total_transcode：所有规格音频，视频即时转码，转封装时长，单位为秒；</li><li>remux：转封装时长，单位为秒；</li><li>transcode_audio：音频转码时长，单位为秒；</li><li>transcode_H264_SD：H.264 编码方式的标清视频（短边 <= 480 px）时长，单位为秒；</li><li>transcode_H264_HD：H.264 编码方式的高清视频（短边 <= 720 px）时长，单位为秒；</li><li>transcode_H264_FHD：H.264 编码方式的全高清视频（短边 <= 1080 px）时长，单位为秒；</li><li>transcode_H264_2K：H.264 编码方式的 2K 视频（短边 <= 1440 px）时长，单位为秒。</li>
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
	// 说明：相同 `Type` 的 `BillingDataFilter` 之间为“或”关系，不同 `Type` 的 `BillingDataFilter` 之间为“且”关系。
	Filters []*BillingDataFilter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 分组聚合维度。最多允许同时按照两种维度进行分组。取值如下：  <li>zone-id：按照站点 ID 进行分组，若使用了内容标识符功能，则优先按照内容标识符分组；<br></li><li>host：按照域名进行分组；<br></li> <li>proxy-id：按照四层代理实例 ID 进行分组；<br></li> <li>region-id：按照计费大区进行分组。</li> 
	GroupBy []*string `json:"GroupBy,omitnil,omitempty" name:"GroupBy"`
}

type DescribeBillingDataRequest struct {
	*tchttp.BaseRequest
	
	// 起始时间。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间。查询时间范围（`EndTime` - `StartTime`）需小于等于 31 天。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 站点 ID 集合，此参数必填。最多传入 100 个站点 ID。若需查询腾讯云主账号下所有站点数据，请用 `*` 代替，查询账号级别数据需具备本接口全部站点资源权限。
	ZoneIds []*string `json:"ZoneIds,omitnil,omitempty" name:"ZoneIds"`

	// 指标列表，取值如下：
	// <b>四/七层加速流量：</b><li>acc_flux: 内容加速流量，单位为 Byte；</li><li>smt_flux: 智能加速流量，单位为 Byte；</li><li>l4_flux: 四层加速流量，单位为 Byte；</li><li>sec_flux: 独立防护流量，单位为 Byte；</li><li>zxctg_flux: 中国大陆网络优化流量，单位为 Byte。</li><br><b>四/七层加速带宽:</b><li>acc_bandwidth: 内容加速带宽，单位为 bps；</li><li>smt_bandwidth: 智能加速带宽，单位为 bps；</li><li>l4_bandwidth: 四层加速带宽，单位为 bps；</li><li>sec_bandwidth: 独立防护带宽，单位为 bps；</li><li>zxctg_bandwidth: 中国大陆网络优化带宽，单位为 bps。</li><br><b>HTTP/HTTPS 安全请求数：</b><li>sec_request_clean: HTTP/HTTPS 请求，单位为次。</li><b><br>增值服务用量：</b><li>smt_request_clean: 智能加速请求，单位为次；</li><li>quic_request: QUIC 请求，单位为次；</li><li>bot_request_clean: Bot 请求，单位为次；</li><li>cls_count: 实时日志推送条数，单位为条；</li><li>ddos_bandwidth: 弹性 DDoS 防护带宽，单位为 bps。</li><br><b>边缘计算用量：</b><li>edgefunction_request：边缘函数请求数，单位为次；</li><li>edgefunction_cpu_time：边缘函数CPU处理时间，单位为毫秒。</li>
	// <b>媒体处理用量：</b><li>total_transcode：所有规格音频，视频即时转码，转封装时长，单位为秒；</li><li>remux：转封装时长，单位为秒；</li><li>transcode_audio：音频转码时长，单位为秒；</li><li>transcode_H264_SD：H.264 编码方式的标清视频（短边 <= 480 px）时长，单位为秒；</li><li>transcode_H264_HD：H.264 编码方式的高清视频（短边 <= 720 px）时长，单位为秒；</li><li>transcode_H264_FHD：H.264 编码方式的全高清视频（短边 <= 1080 px）时长，单位为秒；</li><li>transcode_H264_2K：H.264 编码方式的 2K 视频（短边 <= 1440 px）时长，单位为秒。</li>
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
	// 说明：相同 `Type` 的 `BillingDataFilter` 之间为“或”关系，不同 `Type` 的 `BillingDataFilter` 之间为“且”关系。
	Filters []*BillingDataFilter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 分组聚合维度。最多允许同时按照两种维度进行分组。取值如下：  <li>zone-id：按照站点 ID 进行分组，若使用了内容标识符功能，则优先按照内容标识符分组；<br></li><li>host：按照域名进行分组；<br></li> <li>proxy-id：按照四层代理实例 ID 进行分组；<br></li> <li>region-id：按照计费大区进行分组。</li> 
	GroupBy []*string `json:"GroupBy,omitnil,omitempty" name:"GroupBy"`
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
	delete(f, "GroupBy")
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
type DescribeContentIdentifiersRequestParams struct {
	// 分页查询偏移量。默认值：0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页查询限制数目。默认值：20，最大值：100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 过滤条件，Filters 的上限为 20，Filters.Values 的上限为 20。该参数不填写时，默认返回当前 AppId 下有权限的内容标识符。详细的过滤条件如下：<li>description：按照内容标识符描述批量进行过滤；例如：test；</li><li>content-id：按照内容标识符 ID 批量进行过滤；例如：eocontent-2noz78a8ev6k；</li><li>tag-key：按照标签键进行过滤；</li> <li>tag-value： 按照标签值进行过滤；</li><li>status：按照内容标识符状态进行过滤，取值有：active：生效中；deleted：已删除。</li>仅支持按照 description 模糊查询，其余字段需要精准查询。
	Filters []*AdvancedFilter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeContentIdentifiersRequest struct {
	*tchttp.BaseRequest
	
	// 分页查询偏移量。默认值：0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页查询限制数目。默认值：20，最大值：100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 过滤条件，Filters 的上限为 20，Filters.Values 的上限为 20。该参数不填写时，默认返回当前 AppId 下有权限的内容标识符。详细的过滤条件如下：<li>description：按照内容标识符描述批量进行过滤；例如：test；</li><li>content-id：按照内容标识符 ID 批量进行过滤；例如：eocontent-2noz78a8ev6k；</li><li>tag-key：按照标签键进行过滤；</li> <li>tag-value： 按照标签值进行过滤；</li><li>status：按照内容标识符状态进行过滤，取值有：active：生效中；deleted：已删除。</li>仅支持按照 description 模糊查询，其余字段需要精准查询。
	Filters []*AdvancedFilter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribeContentIdentifiersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeContentIdentifiersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeContentIdentifiersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeContentIdentifiersResponseParams struct {
	// 符合过滤条件的内容标识符总数。	
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 内容标识符详细内容列表。
	ContentIdentifiers []*ContentIdentifier `json:"ContentIdentifiers,omitnil,omitempty" name:"ContentIdentifiers"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeContentIdentifiersResponse struct {
	*tchttp.BaseResponse
	Response *DescribeContentIdentifiersResponseParams `json:"Response"`
}

func (r *DescribeContentIdentifiersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeContentIdentifiersResponse) FromJsonString(s string) error {
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

	// 结束时间。查询时间范围（`EndTime` - `StartTime`）需小于等于 31 天。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 统计指标列表，取值有：
	// <li>ddos_attackMaxBandwidth：攻击带宽峰值；</li>
	// <li>ddos_attackMaxPackageRate：攻击包速率峰值 ；</li>
	// <li>ddos_attackBandwidth：攻击带宽曲线；</li>
	// <li>ddos_attackPackageRate：攻击包速率曲线。</li>
	MetricNames []*string `json:"MetricNames,omitnil,omitempty" name:"MetricNames"`

	// 站点 ID 集合，此参数将于2024年05月30日后由可选改为必填，详见公告：[【腾讯云 EdgeOne】云 API 变更通知](https://cloud.tencent.com/document/product/1552/104902)。最多传入 100 个站点 ID。若需查询腾讯云主账号下所有站点数据，请用 `*` 代替，查询账号级别数据需具备本接口全部站点资源权限。
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

	// 结束时间。查询时间范围（`EndTime` - `StartTime`）需小于等于 31 天。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 统计指标列表，取值有：
	// <li>ddos_attackMaxBandwidth：攻击带宽峰值；</li>
	// <li>ddos_attackMaxPackageRate：攻击包速率峰值 ；</li>
	// <li>ddos_attackBandwidth：攻击带宽曲线；</li>
	// <li>ddos_attackPackageRate：攻击包速率曲线。</li>
	MetricNames []*string `json:"MetricNames,omitnil,omitempty" name:"MetricNames"`

	// 站点 ID 集合，此参数将于2024年05月30日后由可选改为必填，详见公告：[【腾讯云 EdgeOne】云 API 变更通知](https://cloud.tencent.com/document/product/1552/104902)。最多传入 100 个站点 ID。若需查询腾讯云主账号下所有站点数据，请用 `*` 代替，查询账号级别数据需具备本接口全部站点资源权限。
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

	// 结束时间。查询时间范围（`EndTime` - `StartTime`）需小于等于 31 天。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// ddos策略组集合，不填默认选择全部策略。
	PolicyIds []*int64 `json:"PolicyIds,omitnil,omitempty" name:"PolicyIds"`

	// 站点 ID 集合，此参数将于2024年05月30日后由可选改为必填，详见公告：[【腾讯云 EdgeOne】云 API 变更通知](https://cloud.tencent.com/document/product/1552/104902)。最多传入 100 个站点 ID。若需查询腾讯云主账号下所有站点数据，请用 `*` 代替，查询账号级别数据需具备本接口全部站点资源权限。
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

	// 结束时间。查询时间范围（`EndTime` - `StartTime`）需小于等于 31 天。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// ddos策略组集合，不填默认选择全部策略。
	PolicyIds []*int64 `json:"PolicyIds,omitnil,omitempty" name:"PolicyIds"`

	// 站点 ID 集合，此参数将于2024年05月30日后由可选改为必填，详见公告：[【腾讯云 EdgeOne】云 API 变更通知](https://cloud.tencent.com/document/product/1552/104902)。最多传入 100 个站点 ID。若需查询腾讯云主账号下所有站点数据，请用 `*` 代替，查询账号级别数据需具备本接口全部站点资源权限。
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

	// 结束时间。查询时间范围（`EndTime` - `StartTime`）需小于等于 31 天。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 查询的统计指标，取值有：
	// <li>ddos_attackFlux_protocol：按各协议的攻击流量排行；</li>
	// <li>ddos_attackPackageNum_protocol：按各协议的攻击包量排行；</li>
	// <li>ddos_attackNum_attackType：按各攻击类型的攻击数量排行；</li>
	// <li>ddos_attackNum_sregion：按攻击源地区的攻击数量排行；</li>
	// <li>ddos_attackFlux_sip：按攻击源IP的攻击数量排行；</li>
	// <li>ddos_attackFlux_sregion：按攻击源地区的攻击数量排行。</li>
	MetricName *string `json:"MetricName,omitnil,omitempty" name:"MetricName"`

	// 站点 ID 集合，此参数将于2024年05月30日后由可选改为必填，详见公告：[【腾讯云 EdgeOne】云 API 变更通知](https://cloud.tencent.com/document/product/1552/104902)。最多传入 100 个站点 ID。若需查询腾讯云主账号下所有站点数据，请用 `*` 代替，查询账号级别数据需具备本接口全部站点资源权限。
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

	// 结束时间。查询时间范围（`EndTime` - `StartTime`）需小于等于 31 天。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 查询的统计指标，取值有：
	// <li>ddos_attackFlux_protocol：按各协议的攻击流量排行；</li>
	// <li>ddos_attackPackageNum_protocol：按各协议的攻击包量排行；</li>
	// <li>ddos_attackNum_attackType：按各攻击类型的攻击数量排行；</li>
	// <li>ddos_attackNum_sregion：按攻击源地区的攻击数量排行；</li>
	// <li>ddos_attackFlux_sip：按攻击源IP的攻击数量排行；</li>
	// <li>ddos_attackFlux_sregion：按攻击源地区的攻击数量排行。</li>
	MetricName *string `json:"MetricName,omitnil,omitempty" name:"MetricName"`

	// 站点 ID 集合，此参数将于2024年05月30日后由可选改为必填，详见公告：[【腾讯云 EdgeOne】云 API 变更通知](https://cloud.tencent.com/document/product/1552/104902)。最多传入 100 个站点 ID。若需查询腾讯云主账号下所有站点数据，请用 `*` 代替，查询账号级别数据需具备本接口全部站点资源权限。
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
type DescribeDDoSProtectionRequestParams struct {
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`
}

type DescribeDDoSProtectionRequest struct {
	*tchttp.BaseRequest
	
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`
}

func (r *DescribeDDoSProtectionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDDoSProtectionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDDoSProtectionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDDoSProtectionResponseParams struct {
	// 独立 DDoS 防护配置。用于控制独立 DDoS 防护的生效范围。
	DDoSProtection *DDoSProtection `json:"DDoSProtection,omitnil,omitempty" name:"DDoSProtection"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDDoSProtectionResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDDoSProtectionResponseParams `json:"Response"`
}

func (r *DescribeDDoSProtectionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDDoSProtectionResponse) FromJsonString(s string) error {
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
type DescribeDnsRecordsRequestParams struct {
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 分页查询偏移量，默认为 0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页查询限制数目，默认值：20，上限：1000。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 过滤条件，Filters.Values 的上限为20。详细的过滤条件如下：<li>id： 按照 DNS 记录 ID 进行过滤，支持模糊查询；</li><li>name：按照 DNS 记录名称进行过滤，支持模糊查询；</li><li>content：按照 DNS 记录内容进行过滤，支持模糊查询；</li><li>type：按照 DNS 记录类型进行过滤，不支持模糊查询。可选项：<br>   A：将域名指向一个外网 IPv4 地址，如 8.8.8.8；<br>   AAAA：将域名指向一个外网 IPv6 地址；<br>   CNAME：将域名指向另一个域名，再由该域名解析出最终 IP 地址；<br>   TXT：对域名进行标识和说明，常用于域名验证和 SPF 记录（反垃圾邮件）；<br>   NS：如果需要将子域名交给其他 DNS 服务商解析，则需要添加 NS 记录。根域名无法添加 NS 记录；<br>   CAA：指定可为本站点颁发证书的 CA；<br>   SRV：标识某台服务器使用了某个服务，常见于微软系统的目录管理；<br>   MX：指定收件人邮件服务器。</li><li>ttl：按照解析生效时间进行过滤，不支持模糊查询。</li>
	Filters []*AdvancedFilter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 排序依据，取值有：<li>content：DNS 记录内容；</li><li>created-on：DNS 记录创建时间；</li><li>name：DNS 记录名称；</li><li>ttl：缓存时间；</li><li>type：DNS 记录类型。</li>默认根据 type, name 属性组合排序。
	SortBy *string `json:"SortBy,omitnil,omitempty" name:"SortBy"`

	// 列表排序方式，取值有：<li>asc：升序排列；</li><li>desc：降序排列。</li>默认值为 asc。
	SortOrder *string `json:"SortOrder,omitnil,omitempty" name:"SortOrder"`

	// 匹配方式，取值有：<li>all：返回匹配所有查询条件的记录；</li><li>any：返回匹配任意一个查询条件的记录。</li>默认值为 all。
	Match *string `json:"Match,omitnil,omitempty" name:"Match"`
}

type DescribeDnsRecordsRequest struct {
	*tchttp.BaseRequest
	
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 分页查询偏移量，默认为 0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页查询限制数目，默认值：20，上限：1000。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 过滤条件，Filters.Values 的上限为20。详细的过滤条件如下：<li>id： 按照 DNS 记录 ID 进行过滤，支持模糊查询；</li><li>name：按照 DNS 记录名称进行过滤，支持模糊查询；</li><li>content：按照 DNS 记录内容进行过滤，支持模糊查询；</li><li>type：按照 DNS 记录类型进行过滤，不支持模糊查询。可选项：<br>   A：将域名指向一个外网 IPv4 地址，如 8.8.8.8；<br>   AAAA：将域名指向一个外网 IPv6 地址；<br>   CNAME：将域名指向另一个域名，再由该域名解析出最终 IP 地址；<br>   TXT：对域名进行标识和说明，常用于域名验证和 SPF 记录（反垃圾邮件）；<br>   NS：如果需要将子域名交给其他 DNS 服务商解析，则需要添加 NS 记录。根域名无法添加 NS 记录；<br>   CAA：指定可为本站点颁发证书的 CA；<br>   SRV：标识某台服务器使用了某个服务，常见于微软系统的目录管理；<br>   MX：指定收件人邮件服务器。</li><li>ttl：按照解析生效时间进行过滤，不支持模糊查询。</li>
	Filters []*AdvancedFilter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 排序依据，取值有：<li>content：DNS 记录内容；</li><li>created-on：DNS 记录创建时间；</li><li>name：DNS 记录名称；</li><li>ttl：缓存时间；</li><li>type：DNS 记录类型。</li>默认根据 type, name 属性组合排序。
	SortBy *string `json:"SortBy,omitnil,omitempty" name:"SortBy"`

	// 列表排序方式，取值有：<li>asc：升序排列；</li><li>desc：降序排列。</li>默认值为 asc。
	SortOrder *string `json:"SortOrder,omitnil,omitempty" name:"SortOrder"`

	// 匹配方式，取值有：<li>all：返回匹配所有查询条件的记录；</li><li>any：返回匹配任意一个查询条件的记录。</li>默认值为 all。
	Match *string `json:"Match,omitnil,omitempty" name:"Match"`
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
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	delete(f, "SortBy")
	delete(f, "SortOrder")
	delete(f, "Match")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDnsRecordsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDnsRecordsResponseParams struct {
	// DNS 记录总数。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// DNS 记录列表。
	DnsRecords []*DnsRecord `json:"DnsRecords,omitnil,omitempty" name:"DnsRecords"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DescribeJustInTimeTranscodeTemplatesRequestParams struct {
	// 站点ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 过滤条件，Filters 的上限为 20，Filters.Values 的上限为 20。该参数不填写时，默认返回当前 ZoneId 下有权限的即时转码模板。详细的过滤条件如下：<li>template-name：按照模板名批量进行过滤。例如：mytemplate；</li><li>template-type：按照模板类型批量进行过滤。例如：preset 或 custom。</li><li>template-id：按照模板 ID 批量进行过滤。例如：C1LZ7982VgTpYhJ7M。</li>默认为空。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 排序字段，取值有：<li>createTime：模板创建时间。</li>默认值为：createTime。
	SortBy *string `json:"SortBy,omitnil,omitempty" name:"SortBy"`

	// 排序方式，取值有：<li>asc：升序方式；</li><li>desc：降序方式。</li>默认值为：desc。
	SortOrder *string `json:"SortOrder,omitnil,omitempty" name:"SortOrder"`

	// 分页偏移量，默认值：0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回记录条数，默认值：20，最大值：1000。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeJustInTimeTranscodeTemplatesRequest struct {
	*tchttp.BaseRequest
	
	// 站点ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 过滤条件，Filters 的上限为 20，Filters.Values 的上限为 20。该参数不填写时，默认返回当前 ZoneId 下有权限的即时转码模板。详细的过滤条件如下：<li>template-name：按照模板名批量进行过滤。例如：mytemplate；</li><li>template-type：按照模板类型批量进行过滤。例如：preset 或 custom。</li><li>template-id：按照模板 ID 批量进行过滤。例如：C1LZ7982VgTpYhJ7M。</li>默认为空。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 排序字段，取值有：<li>createTime：模板创建时间。</li>默认值为：createTime。
	SortBy *string `json:"SortBy,omitnil,omitempty" name:"SortBy"`

	// 排序方式，取值有：<li>asc：升序方式；</li><li>desc：降序方式。</li>默认值为：desc。
	SortOrder *string `json:"SortOrder,omitnil,omitempty" name:"SortOrder"`

	// 分页偏移量，默认值：0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回记录条数，默认值：20，最大值：1000。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeJustInTimeTranscodeTemplatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeJustInTimeTranscodeTemplatesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "Filters")
	delete(f, "SortBy")
	delete(f, "SortOrder")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeJustInTimeTranscodeTemplatesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeJustInTimeTranscodeTemplatesResponseParams struct {
	// 符合过滤条件的记录总数。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 模板详情列表。
	TemplateSet []*JustInTimeTranscodeTemplate `json:"TemplateSet,omitnil,omitempty" name:"TemplateSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeJustInTimeTranscodeTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeJustInTimeTranscodeTemplatesResponseParams `json:"Response"`
}

func (r *DescribeJustInTimeTranscodeTemplatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeJustInTimeTranscodeTemplatesResponse) FromJsonString(s string) error {
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

	// 过滤条件，Filters.Values的上限为20。不填写时返回当前四层实例下所有的规则信息，详细的过滤条件如下：  <li>rule-id：按照规则 ID 对四层代理实例下的规则进行过滤。规则 ID 形如：rule-31vv7qig0vjy；</li> <li>rule-tag：按照规则标签对四层代理实例下的规则进行过滤。</li>
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

	// 过滤条件，Filters.Values的上限为20。不填写时返回当前四层实例下所有的规则信息，详细的过滤条件如下：  <li>rule-id：按照规则 ID 对四层代理实例下的规则进行过滤。规则 ID 形如：rule-31vv7qig0vjy；</li> <li>rule-tag：按照规则标签对四层代理实例下的规则进行过滤。</li>
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
type DescribeL7AccRulesRequestParams struct {
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 过滤条件，Filters.Values 的上限为 20，不填写此参数时默认按顺序返回站点下的规则。详细的过滤条件如下：
	// <li>rule-id：按照规则 ID 进行过滤。</li>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 分页查询限制数目，默认值：20，上限：1000。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页查询偏移量，默认为 0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type DescribeL7AccRulesRequest struct {
	*tchttp.BaseRequest
	
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 过滤条件，Filters.Values 的上限为 20，不填写此参数时默认按顺序返回站点下的规则。详细的过滤条件如下：
	// <li>rule-id：按照规则 ID 进行过滤。</li>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 分页查询限制数目，默认值：20，上限：1000。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页查询偏移量，默认为 0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

func (r *DescribeL7AccRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeL7AccRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "Filters")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeL7AccRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeL7AccRulesResponseParams struct {
	// 规则总数。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 规则列表，规则按照从上到下的顺序执行，详情参考[规则生效优先级](https://cloud.tencent.com/document/product/1552/70901#.E4.BC.98.E5.85.88.E7.BA.A7)。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Rules []*RuleEngineItem `json:"Rules,omitnil,omitempty" name:"Rules"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeL7AccRulesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeL7AccRulesResponseParams `json:"Response"`
}

func (r *DescribeL7AccRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeL7AccRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeL7AccSettingRequestParams struct {
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`
}

type DescribeL7AccSettingRequest struct {
	*tchttp.BaseRequest
	
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`
}

func (r *DescribeL7AccSettingRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeL7AccSettingRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeL7AccSettingRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeL7AccSettingResponseParams struct {
	// 站点加速全局配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ZoneSetting *ZoneConfigParameters `json:"ZoneSetting,omitnil,omitempty" name:"ZoneSetting"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeL7AccSettingResponse struct {
	*tchttp.BaseResponse
	Response *DescribeL7AccSettingResponseParams `json:"Response"`
}

func (r *DescribeL7AccSettingResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeL7AccSettingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLoadBalancerListRequestParams struct {
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 分页查询偏移量，默认为 0。	
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页查询限制数目，默认值：20，最大值：100。	
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 过滤条件，Filters.Values 的上限为 20。该参数不填写时，返回当前 zone-id 下所有负载均衡实例信息。详细的过滤条件如下：
	// <li>InstanceName：按照负载均衡实例名称进行过滤；</li>
	// <li>InstanceId：按照负载均衡实例 ID 进行过滤。</li>  
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeLoadBalancerListRequest struct {
	*tchttp.BaseRequest
	
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 分页查询偏移量，默认为 0。	
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页查询限制数目，默认值：20，最大值：100。	
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 过滤条件，Filters.Values 的上限为 20。该参数不填写时，返回当前 zone-id 下所有负载均衡实例信息。详细的过滤条件如下：
	// <li>InstanceName：按照负载均衡实例名称进行过滤；</li>
	// <li>InstanceId：按照负载均衡实例 ID 进行过滤。</li>  
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribeLoadBalancerListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLoadBalancerListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLoadBalancerListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLoadBalancerListResponseParams struct {
	// 负载均衡实例总数。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 负载均衡实例列表。
	LoadBalancerList []*LoadBalancer `json:"LoadBalancerList,omitnil,omitempty" name:"LoadBalancerList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeLoadBalancerListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLoadBalancerListResponseParams `json:"Response"`
}

func (r *DescribeLoadBalancerListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLoadBalancerListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMultiPathGatewayLineRequestParams struct {
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 网关 ID。
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// 线路 ID。
	LineId *string `json:"LineId,omitnil,omitempty" name:"LineId"`
}

type DescribeMultiPathGatewayLineRequest struct {
	*tchttp.BaseRequest
	
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 网关 ID。
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// 线路 ID。
	LineId *string `json:"LineId,omitnil,omitempty" name:"LineId"`
}

func (r *DescribeMultiPathGatewayLineRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMultiPathGatewayLineRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "GatewayId")
	delete(f, "LineId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMultiPathGatewayLineRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMultiPathGatewayLineResponseParams struct {
	// 线路信息。
	Line *MultiPathGatewayLine `json:"Line,omitnil,omitempty" name:"Line"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeMultiPathGatewayLineResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMultiPathGatewayLineResponseParams `json:"Response"`
}

func (r *DescribeMultiPathGatewayLineResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMultiPathGatewayLineResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMultiPathGatewayOriginACLRequestParams struct {
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 网关 ID。
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`
}

type DescribeMultiPathGatewayOriginACLRequest struct {
	*tchttp.BaseRequest
	
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 网关 ID。
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`
}

func (r *DescribeMultiPathGatewayOriginACLRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMultiPathGatewayOriginACLRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "GatewayId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMultiPathGatewayOriginACLRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMultiPathGatewayOriginACLResponseParams struct {
	// 多通道网关实例与回源 IP 网段的绑定关系详情。
	MultiPathGatewayOriginACLInfo *MultiPathGatewayOriginACLInfo `json:"MultiPathGatewayOriginACLInfo,omitnil,omitempty" name:"MultiPathGatewayOriginACLInfo"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeMultiPathGatewayOriginACLResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMultiPathGatewayOriginACLResponseParams `json:"Response"`
}

func (r *DescribeMultiPathGatewayOriginACLResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMultiPathGatewayOriginACLResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMultiPathGatewayRegionsRequestParams struct {
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`
}

type DescribeMultiPathGatewayRegionsRequest struct {
	*tchttp.BaseRequest
	
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`
}

func (r *DescribeMultiPathGatewayRegionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMultiPathGatewayRegionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMultiPathGatewayRegionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMultiPathGatewayRegionsResponseParams struct {
	// 网关可用地域列表。
	GatewayRegions []*GatewayRegion `json:"GatewayRegions,omitnil,omitempty" name:"GatewayRegions"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeMultiPathGatewayRegionsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMultiPathGatewayRegionsResponseParams `json:"Response"`
}

func (r *DescribeMultiPathGatewayRegionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMultiPathGatewayRegionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMultiPathGatewayRequestParams struct {
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 网关 ID。
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`
}

type DescribeMultiPathGatewayRequest struct {
	*tchttp.BaseRequest
	
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 网关 ID。
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`
}

func (r *DescribeMultiPathGatewayRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMultiPathGatewayRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "GatewayId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMultiPathGatewayRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMultiPathGatewayResponseParams struct {
	// 网关详情。
	GatewayDetail *MultiPathGateway `json:"GatewayDetail,omitnil,omitempty" name:"GatewayDetail"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeMultiPathGatewayResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMultiPathGatewayResponseParams `json:"Response"`
}

func (r *DescribeMultiPathGatewayResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMultiPathGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMultiPathGatewaySecretKeyRequestParams struct {
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`
}

type DescribeMultiPathGatewaySecretKeyRequest struct {
	*tchttp.BaseRequest
	
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`
}

func (r *DescribeMultiPathGatewaySecretKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMultiPathGatewaySecretKeyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMultiPathGatewaySecretKeyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMultiPathGatewaySecretKeyResponseParams struct {
	// 接入密钥。
	SecretKey *string `json:"SecretKey,omitnil,omitempty" name:"SecretKey"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeMultiPathGatewaySecretKeyResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMultiPathGatewaySecretKeyResponseParams `json:"Response"`
}

func (r *DescribeMultiPathGatewaySecretKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMultiPathGatewaySecretKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMultiPathGatewaysRequestParams struct {
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 分页查询偏移量。默认值：0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页查询限制数目。默认值：20，最大值：1000。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 网关列表的过滤字段，该参数不填写时，返回当前 appid 下所有网关信息，详细的过滤条件如下：
	// <li> gateway-type：按照网关类型进行过滤，支持取值 cloud 和 private，分别代表过滤云上网关和自由网关；</li>
	// <li> keyword：按照网关名的关键字进行过滤。</li>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeMultiPathGatewaysRequest struct {
	*tchttp.BaseRequest
	
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 分页查询偏移量。默认值：0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页查询限制数目。默认值：20，最大值：1000。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 网关列表的过滤字段，该参数不填写时，返回当前 appid 下所有网关信息，详细的过滤条件如下：
	// <li> gateway-type：按照网关类型进行过滤，支持取值 cloud 和 private，分别代表过滤云上网关和自由网关；</li>
	// <li> keyword：按照网关名的关键字进行过滤。</li>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribeMultiPathGatewaysRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMultiPathGatewaysRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMultiPathGatewaysRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMultiPathGatewaysResponseParams struct {
	// 网关详情。
	Gateways []*MultiPathGateway `json:"Gateways,omitnil,omitempty" name:"Gateways"`

	// 总条数。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeMultiPathGatewaysResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMultiPathGatewaysResponseParams `json:"Response"`
}

func (r *DescribeMultiPathGatewaysResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMultiPathGatewaysResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOriginACLRequestParams struct {
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`
}

type DescribeOriginACLRequest struct {
	*tchttp.BaseRequest
	
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`
}

func (r *DescribeOriginACLRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOriginACLRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeOriginACLRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOriginACLResponseParams struct {
	// 七层加速域名/四层代理实例与回源 IP 网段的绑定关系详情。
	OriginACLInfo *OriginACLInfo `json:"OriginACLInfo,omitnil,omitempty" name:"OriginACLInfo"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeOriginACLResponse struct {
	*tchttp.BaseResponse
	Response *DescribeOriginACLResponseParams `json:"Response"`
}

func (r *DescribeOriginACLResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOriginACLResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOriginGroupHealthStatusRequestParams struct {
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 负载均衡实例 ID。
	LBInstanceId *string `json:"LBInstanceId,omitnil,omitempty" name:"LBInstanceId"`

	// 源站组 ID。不填写时默认获取负载均衡下所有源站组的健康状态。
	OriginGroupIds []*string `json:"OriginGroupIds,omitnil,omitempty" name:"OriginGroupIds"`
}

type DescribeOriginGroupHealthStatusRequest struct {
	*tchttp.BaseRequest
	
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 负载均衡实例 ID。
	LBInstanceId *string `json:"LBInstanceId,omitnil,omitempty" name:"LBInstanceId"`

	// 源站组 ID。不填写时默认获取负载均衡下所有源站组的健康状态。
	OriginGroupIds []*string `json:"OriginGroupIds,omitnil,omitempty" name:"OriginGroupIds"`
}

func (r *DescribeOriginGroupHealthStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOriginGroupHealthStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "LBInstanceId")
	delete(f, "OriginGroupIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeOriginGroupHealthStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOriginGroupHealthStatusResponseParams struct {
	// 源站组下源站的健康状态。
	OriginGroupHealthStatusList []*OriginGroupHealthStatusDetail `json:"OriginGroupHealthStatusList,omitnil,omitempty" name:"OriginGroupHealthStatusList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeOriginGroupHealthStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribeOriginGroupHealthStatusResponseParams `json:"Response"`
}

func (r *DescribeOriginGroupHealthStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOriginGroupHealthStatusResponse) FromJsonString(s string) error {
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
	// <li>origin-group-id：按照源站组 ID 进行过滤，不支持模糊查询。源站组 ID 形如：origin-2ccgtb24-7dc5-46s2-9r3e-95825d53dwe3a；</li><li>origin-group-name： 按照源站组名称进行过滤，使用模糊查询时，仅支持填写一个源站组名称。</li>
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
	// <li>origin-group-id：按照源站组 ID 进行过滤，不支持模糊查询。源站组 ID 形如：origin-2ccgtb24-7dc5-46s2-9r3e-95825d53dwe3a；</li><li>origin-group-name： 按照源站组名称进行过滤，使用模糊查询时，仅支持填写一个源站组名称。</li>
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
type DescribePlansRequestParams struct {
	// 过滤条件，Filters.Values 的上限为 20。详细的过滤条件如下：<li>plan-type<br>  按照【<strong>套餐类型</strong>】进行过滤。<br>  可选的类型有：<br>  plan-trial：试用版套餐；<br>  plan-personal：个人版套餐；<br>  plan-basic：基础版套餐； <br>  plan-standard：标准版套餐； <br>  plan-enterprise：企业版套餐。 </li><li>plan-id<br>  按照【<strong>套餐 ID</strong>】进行过滤。套餐 ID 形如：edgeone-268z103ob0sx。</li><li>area<br>  按照【<strong>套餐加速地域</strong>】进行过滤。</li>  服务区域，可选的类型有：<br>  mainland: 中国大陆；<br>  overseas: 全球（不包括中国大陆)；<br>  global: 全球（包括中国大陆)。<br><li>status<br>  按照【<strong>套餐状态</strong>】进行过滤。<br>  可选的状态有：<br>  normal：正常状态；<br>  expiring-soon：即将过期；<br>  expired：已到期;<br>  isolated：已隔离。</li>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 排序字段，取值有：
	// <li> enable-time：生效时间；</li>
	// <li> expire-time：过期时间。</li>不填写使用默认值 enable-time。
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// 排序方向，取值有：
	// <li>asc：从小到大排序；</li>
	// <li>desc：从大到小排序。</li>不填写使用默认值 desc。
	Direction *string `json:"Direction,omitnil,omitempty" name:"Direction"`

	// 分页查询限制数目。默认值：20，最大值：200。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页查询偏移量。默认值：0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type DescribePlansRequest struct {
	*tchttp.BaseRequest
	
	// 过滤条件，Filters.Values 的上限为 20。详细的过滤条件如下：<li>plan-type<br>  按照【<strong>套餐类型</strong>】进行过滤。<br>  可选的类型有：<br>  plan-trial：试用版套餐；<br>  plan-personal：个人版套餐；<br>  plan-basic：基础版套餐； <br>  plan-standard：标准版套餐； <br>  plan-enterprise：企业版套餐。 </li><li>plan-id<br>  按照【<strong>套餐 ID</strong>】进行过滤。套餐 ID 形如：edgeone-268z103ob0sx。</li><li>area<br>  按照【<strong>套餐加速地域</strong>】进行过滤。</li>  服务区域，可选的类型有：<br>  mainland: 中国大陆；<br>  overseas: 全球（不包括中国大陆)；<br>  global: 全球（包括中国大陆)。<br><li>status<br>  按照【<strong>套餐状态</strong>】进行过滤。<br>  可选的状态有：<br>  normal：正常状态；<br>  expiring-soon：即将过期；<br>  expired：已到期;<br>  isolated：已隔离。</li>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 排序字段，取值有：
	// <li> enable-time：生效时间；</li>
	// <li> expire-time：过期时间。</li>不填写使用默认值 enable-time。
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// 排序方向，取值有：
	// <li>asc：从小到大排序；</li>
	// <li>desc：从大到小排序。</li>不填写使用默认值 desc。
	Direction *string `json:"Direction,omitnil,omitempty" name:"Direction"`

	// 分页查询限制数目。默认值：20，最大值：200。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页查询偏移量。默认值：0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

func (r *DescribePlansRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePlansRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Order")
	delete(f, "Direction")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePlansRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePlansResponseParams struct {
	// 符合条件的套餐个数。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 套餐信息列表。
	Plans []*Plan `json:"Plans,omitnil,omitempty" name:"Plans"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribePlansResponse struct {
	*tchttp.BaseResponse
	Response *DescribePlansResponseParams `json:"Response"`
}

func (r *DescribePlansResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePlansResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePrefetchOriginLimitRequestParams struct {
	// 站点ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 分页查询偏移量，默认为 0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页查询限制数目，默认值：20，上限：100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 过滤条件，Filters.Values 的上限为 20。详细的过滤条件如下：
	// <li>domain-name：按照域名过滤。domain-name 形如：www.qq.com，不支持模糊查询；</li>
	// <li>area：按照限制加速区域过滤，不支持模糊查询。可选项：<br> Overseas：全球可用区（不含中国大陆）；<br> MainlandChina：中国大陆可用区。</li>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribePrefetchOriginLimitRequest struct {
	*tchttp.BaseRequest
	
	// 站点ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 分页查询偏移量，默认为 0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页查询限制数目，默认值：20，上限：100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 过滤条件，Filters.Values 的上限为 20。详细的过滤条件如下：
	// <li>domain-name：按照域名过滤。domain-name 形如：www.qq.com，不支持模糊查询；</li>
	// <li>area：按照限制加速区域过滤，不支持模糊查询。可选项：<br> Overseas：全球可用区（不含中国大陆）；<br> MainlandChina：中国大陆可用区。</li>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribePrefetchOriginLimitRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrefetchOriginLimitRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePrefetchOriginLimitRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePrefetchOriginLimitResponseParams struct {
	// 回源限速限制总数。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 回源限速限制详情List。
	Limits []*PrefetchOriginLimit `json:"Limits,omitnil,omitempty" name:"Limits"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribePrefetchOriginLimitResponse struct {
	*tchttp.BaseResponse
	Response *DescribePrefetchOriginLimitResponseParams `json:"Response"`
}

func (r *DescribePrefetchOriginLimitResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrefetchOriginLimitResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePrefetchTasksRequestParams struct {
	// 站点ID。此参数将于2024年05月30日后由可选改为必填，详见公告：[【腾讯云 EdgeOne】云 API 变更通知](https://cloud.tencent.com/document/product/1552/104902)。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 查询起始时间，时间与 job-id 必填一个。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 查询结束时间，时间与 job-id 必填一个。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 分页查询偏移量，默认为 0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页查询限制数目，默认值：20，上限：1000。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 过滤条件，Filters.Values 的上限为 20。详细的过滤条件如下：<li>job-id：按照任务 ID 进行过滤。job-id 形如：1379afjk91u32h，暂不支持多值，不支持模糊查询；</li><li>target：按照目标资源信息进行过滤。target 形如：http://www.qq.com/1.txt，暂不支持多值，不支持模糊查询；</li><li>domains：按照域名行过滤。domains 形如：www.qq.com，不支持模糊查询；</li><li>statuses：按照任务状态进行过滤，不支持模糊查询。可选项：<br>   processing：处理中<br>   success：成功<br>   failed：失败<br>   timeout：超时<br>   canceled：已取消<br>   invalid：无效。即源站响应非 2xx 状态码，请检查源站服务。</li>
	Filters []*AdvancedFilter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribePrefetchTasksRequest struct {
	*tchttp.BaseRequest
	
	// 站点ID。此参数将于2024年05月30日后由可选改为必填，详见公告：[【腾讯云 EdgeOne】云 API 变更通知](https://cloud.tencent.com/document/product/1552/104902)。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 查询起始时间，时间与 job-id 必填一个。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 查询结束时间，时间与 job-id 必填一个。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 分页查询偏移量，默认为 0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页查询限制数目，默认值：20，上限：1000。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 过滤条件，Filters.Values 的上限为 20。详细的过滤条件如下：<li>job-id：按照任务 ID 进行过滤。job-id 形如：1379afjk91u32h，暂不支持多值，不支持模糊查询；</li><li>target：按照目标资源信息进行过滤。target 形如：http://www.qq.com/1.txt，暂不支持多值，不支持模糊查询；</li><li>domains：按照域名行过滤。domains 形如：www.qq.com，不支持模糊查询；</li><li>statuses：按照任务状态进行过滤，不支持模糊查询。可选项：<br>   processing：处理中<br>   success：成功<br>   failed：失败<br>   timeout：超时<br>   canceled：已取消<br>   invalid：无效。即源站响应非 2xx 状态码，请检查源站服务。</li>
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
	// 站点 ID。此参数将于2024年05月30日后由可选改为必填，详见公告：[【腾讯云 EdgeOne】云 API 变更通知](https://cloud.tencent.com/document/product/1552/104902)。
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
	// <li>statuses：按照任务状态进行过滤，不支持模糊查询。可选项：<br>   processing：处理中<br>   success：成功<br>   failed：失败<br>   timeout：超时<br>   canceled：已取消</li>
	// <li>type：按照清除缓存类型进行过滤，暂不支持多值，不支持模糊查询。可选项：<br>   purge_url：URL<br>   purge_prefix：前缀<br>   purge_all：全部缓存内容<br>   purge_host：Hostname<br>   purge_cache_tag：CacheTag</li>
	Filters []*AdvancedFilter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribePurgeTasksRequest struct {
	*tchttp.BaseRequest
	
	// 站点 ID。此参数将于2024年05月30日后由可选改为必填，详见公告：[【腾讯云 EdgeOne】云 API 变更通知](https://cloud.tencent.com/document/product/1552/104902)。
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
	// <li>statuses：按照任务状态进行过滤，不支持模糊查询。可选项：<br>   processing：处理中<br>   success：成功<br>   failed：失败<br>   timeout：超时<br>   canceled：已取消</li>
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
	// <li>task-type：按照实时日志投递任务类型进行过滤。不支持模糊查询。可选项如下：<br>   cls: 推送到腾讯云 CLS；<br>   custom_endpoint：推送到自定义 HTTP(S) 地址；<br>   s3：推送到 AWS S3 兼容存储桶地址；<br>   log_analysis：推送到 EdgeOne 日志分析。</li>
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
	// <li>task-type：按照实时日志投递任务类型进行过滤。不支持模糊查询。可选项如下：<br>   cls: 推送到腾讯云 CLS；<br>   custom_endpoint：推送到自定义 HTTP(S) 地址；<br>   s3：推送到 AWS S3 兼容存储桶地址；<br>   log_analysis：推送到 EdgeOne 日志分析。</li>
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
	// <li>rule-id：按照规则 ID 进行过滤。</li>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeRulesRequest struct {
	*tchttp.BaseRequest
	
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 过滤条件，Filters.Values的上限为20。详细的过滤条件如下：
	// <li>rule-id：按照规则 ID 进行过滤。</li>
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
type DescribeSecurityAPIResourceRequestParams struct {
	// 站点 ID。	
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 分页查询限制数目。默认值：20，最大值：100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页查询偏移量。默认值：0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type DescribeSecurityAPIResourceRequest struct {
	*tchttp.BaseRequest
	
	// 站点 ID。	
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 分页查询限制数目。默认值：20，最大值：100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页查询偏移量。默认值：0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

func (r *DescribeSecurityAPIResourceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecurityAPIResourceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSecurityAPIResourceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSecurityAPIResourceResponseParams struct {
	// API 资源总数量。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// API 资源列表。	
	APIResources []*APIResource `json:"APIResources,omitnil,omitempty" name:"APIResources"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSecurityAPIResourceResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSecurityAPIResourceResponseParams `json:"Response"`
}

func (r *DescribeSecurityAPIResourceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecurityAPIResourceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSecurityAPIServiceRequestParams struct {
	// 站点 ID。	
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 分页查询限制数目。默认值：20，最大值：100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页查询偏移量。默认值：0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type DescribeSecurityAPIServiceRequest struct {
	*tchttp.BaseRequest
	
	// 站点 ID。	
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 分页查询限制数目。默认值：20，最大值：100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页查询偏移量。默认值：0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

func (r *DescribeSecurityAPIServiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecurityAPIServiceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSecurityAPIServiceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSecurityAPIServiceResponseParams struct {
	// API 服务总数量。	
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// API 服务列表。	
	APIServices []*APIService `json:"APIServices,omitnil,omitempty" name:"APIServices"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSecurityAPIServiceResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSecurityAPIServiceResponseParams `json:"Response"`
}

func (r *DescribeSecurityAPIServiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecurityAPIServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSecurityClientAttesterRequestParams struct {
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 分页查询限制数目。默认值：20，最大值：100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页查询偏移量。默认值：0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type DescribeSecurityClientAttesterRequest struct {
	*tchttp.BaseRequest
	
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 分页查询限制数目。默认值：20，最大值：100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页查询偏移量。默认值：0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

func (r *DescribeSecurityClientAttesterRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecurityClientAttesterRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSecurityClientAttesterRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSecurityClientAttesterResponseParams struct {
	// 认证选项总数量。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 认证选项列表。
	ClientAttesters []*ClientAttester `json:"ClientAttesters,omitnil,omitempty" name:"ClientAttesters"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSecurityClientAttesterResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSecurityClientAttesterResponseParams `json:"Response"`
}

func (r *DescribeSecurityClientAttesterResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecurityClientAttesterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSecurityIPGroupContentRequestParams struct {
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// IP 组 ID。
	GroupId *int64 `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 分页查询限制数目。默认值：2000，最大值：100000。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页查询偏移量。默认值：0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type DescribeSecurityIPGroupContentRequest struct {
	*tchttp.BaseRequest
	
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// IP 组 ID。
	GroupId *int64 `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 分页查询限制数目。默认值：2000，最大值：100000。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页查询偏移量。默认值：0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

func (r *DescribeSecurityIPGroupContentRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecurityIPGroupContentRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "GroupId")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSecurityIPGroupContentRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSecurityIPGroupContentResponseParams struct {
	// IP 组中正在生效的 IP 或网段个数。
	IPTotalCount *int64 `json:"IPTotalCount,omitnil,omitempty" name:"IPTotalCount"`

	// 满足查询条件的 IP 或网段列表。受 Limit 和 Offset 参数限制。
	IPList []*string `json:"IPList,omitnil,omitempty" name:"IPList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSecurityIPGroupContentResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSecurityIPGroupContentResponseParams `json:"Response"`
}

func (r *DescribeSecurityIPGroupContentResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecurityIPGroupContentResponse) FromJsonString(s string) error {
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

	// 指定安全 IP 组 ID。 <li>提供该参数时，仅查询指定 ID 的安全 IP 组配置；</li> <li>不传递参数时，返回站点下所有安全 IP 组信息。</li>
	GroupIds []*int64 `json:"GroupIds,omitnil,omitempty" name:"GroupIds"`
}

type DescribeSecurityIPGroupRequest struct {
	*tchttp.BaseRequest
	
	// 站点 ID ，用于指定查询的站点范围。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 指定安全 IP 组 ID。 <li>提供该参数时，仅查询指定 ID 的安全 IP 组配置；</li> <li>不传递参数时，返回站点下所有安全 IP 组信息。</li>
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
	// 安全 IP 组的详细配置信息。包含每个安全 IP 组的 ID 、名称、IP / 网段总数量、 IP / 网段列表信息和过期时间信息。
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
type DescribeSecurityJSInjectionRuleRequestParams struct {
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 分页查询限制数目。默认值：20，最大值：100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页查询偏移量。默认值：0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type DescribeSecurityJSInjectionRuleRequest struct {
	*tchttp.BaseRequest
	
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 分页查询限制数目。默认值：20，最大值：100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页查询偏移量。默认值：0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

func (r *DescribeSecurityJSInjectionRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecurityJSInjectionRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSecurityJSInjectionRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSecurityJSInjectionRuleResponseParams struct {
	// JavaScript 注入规则总数量。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// JavaScript 注入规则列表。
	JSInjectionRules []*JSInjectionRule `json:"JSInjectionRules,omitnil,omitempty" name:"JSInjectionRules"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSecurityJSInjectionRuleResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSecurityJSInjectionRuleResponseParams `json:"Response"`
}

func (r *DescribeSecurityJSInjectionRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecurityJSInjectionRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSecurityPolicyRequestParams struct {
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 安全策略类型，可使用以下参数值进行查询： <li>ZoneDefaultPolicy：用于指定查询站点级策略；</li><li>Template：用于指定查询策略模板，需要同时指定 TemplateId 参数；</li><li>Host：用于指定查询域名级策略（注意：当使用域名来指定域名服务策略时，仅支持已经应用了域名级策略的域名服务或者策略模板）。</li>	
	Entity *string `json:"Entity,omitnil,omitempty" name:"Entity"`

	// 指定策略模板 ID。当 Entity 参数值为 Template 时，使用本参数指定策略模板的 ID 查询模板配置。
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// 指定域名。当 Entity 参数值为 Host 时，使用本参数指定的域名级策略查询域名配置，例如：使用 www.example.com ，配置该域名的域名级策略。
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`
}

type DescribeSecurityPolicyRequest struct {
	*tchttp.BaseRequest
	
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 安全策略类型，可使用以下参数值进行查询： <li>ZoneDefaultPolicy：用于指定查询站点级策略；</li><li>Template：用于指定查询策略模板，需要同时指定 TemplateId 参数；</li><li>Host：用于指定查询域名级策略（注意：当使用域名来指定域名服务策略时，仅支持已经应用了域名级策略的域名服务或者策略模板）。</li>	
	Entity *string `json:"Entity,omitnil,omitempty" name:"Entity"`

	// 指定策略模板 ID。当 Entity 参数值为 Template 时，使用本参数指定策略模板的 ID 查询模板配置。
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// 指定域名。当 Entity 参数值为 Host 时，使用本参数指定的域名级策略查询域名配置，例如：使用 www.example.com ，配置该域名的域名级策略。
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`
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
	delete(f, "TemplateId")
	delete(f, "Host")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSecurityPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSecurityPolicyResponseParams struct {
	// 安全策略配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SecurityPolicy *SecurityPolicy `json:"SecurityPolicy,omitnil,omitempty" name:"SecurityPolicy"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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

	// 结束时间。查询时间范围（`EndTime` - `StartTime`）需小于等于 31 天。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 查询指标，取值有：
	// <li>l4Flow_connections: 访问并发连接数；</li>
	// <li>l4Flow_flux: 访问总流量；</li>
	// <li>l4Flow_inFlux: 访问入流量；</li>
	// <li>l4Flow_outFlux: 访问出流量；</li>
	// <li>l4Flow_inBandwidth: 访问入向带宽峰值；</li>
	// <li>l4Flow_outBandwidth: 访问出向带宽峰值。</li>
	MetricNames []*string `json:"MetricNames,omitnil,omitempty" name:"MetricNames"`

	// 站点ID，此参数将于2024年05月30日后由可选改为必填，详见公告：[【腾讯云 EdgeOne】云 API 变更通知](https://cloud.tencent.com/document/product/1552/104902)。
	// 最多传入 100 个站点 ID。若需查询腾讯云主账号下所有站点数据，请用 `*` 代替，查询账号级别数据需具备本接口全部站点资源权限。
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

	// 数据归属地区。该参数已废弃。请在 Filters.country 中按客户端地域过滤数据。
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`
}

type DescribeTimingL4DataRequest struct {
	*tchttp.BaseRequest
	
	// 开始时间。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间。查询时间范围（`EndTime` - `StartTime`）需小于等于 31 天。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 查询指标，取值有：
	// <li>l4Flow_connections: 访问并发连接数；</li>
	// <li>l4Flow_flux: 访问总流量；</li>
	// <li>l4Flow_inFlux: 访问入流量；</li>
	// <li>l4Flow_outFlux: 访问出流量；</li>
	// <li>l4Flow_inBandwidth: 访问入向带宽峰值；</li>
	// <li>l4Flow_outBandwidth: 访问出向带宽峰值。</li>
	MetricNames []*string `json:"MetricNames,omitnil,omitempty" name:"MetricNames"`

	// 站点ID，此参数将于2024年05月30日后由可选改为必填，详见公告：[【腾讯云 EdgeOne】云 API 变更通知](https://cloud.tencent.com/document/product/1552/104902)。
	// 最多传入 100 个站点 ID。若需查询腾讯云主账号下所有站点数据，请用 `*` 代替，查询账号级别数据需具备本接口全部站点资源权限。
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

	// 数据归属地区。该参数已废弃。请在 Filters.country 中按客户端地域过滤数据。
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

	// 结束时间。查询时间范围（`EndTime` - `StartTime`）需小于等于 31 天。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 指标列表，取值有:
	// <li>l7Flow_outFlux: L7 EdgeOne 响应流量，单位：Byte；</li>
	// <li>l7Flow_inFlux: L7 客户端请求流量，单位：Byte；</li>
	// <li>l7Flow_flux: L7 访问总流量（EdgeOne 响应+客户端请求），单位：Byte；</li>
	// <li>l7Flow_outBandwidth: L7 EdgeOne 响应带宽，单位：bps；</li>
	// <li>l7Flow_inBandwidth：L7 客户端请求带宽，单位：bps；</li>
	// <li>l7Flow_bandwidth：L7 访问总带宽（EdgeOne 响应+客户端请求），单位：bps；</li>
	// <li>l7Flow_request: L7 访问请求数，单位：次；</li>
	// <li> l7Flow_avgResponseTime: L7 访问平均响应耗时，单位：ms；</li>
	// <li> l7Flow_avgFirstByteResponseTime: L7 访问平均首字节响应耗时，单位：ms。</li>
	MetricNames []*string `json:"MetricNames,omitnil,omitempty" name:"MetricNames"`

	// 站点 ID 集合，此参数将于2024年05月30日后由可选改为必填，详见公告：[【腾讯云 EdgeOne】云 API 变更通知](https://cloud.tencent.com/document/product/1552/104902)。最多传入 100 个站点 ID。若需查询腾讯云主账号下所有站点数据，请用 `*` 代替，查询账号级别数据需具备本接口全部站点资源权限。
	ZoneIds []*string `json:"ZoneIds,omitnil,omitempty" name:"ZoneIds"`

	// 查询时间粒度，取值有：
	// <li>min: 1分钟；</li>
	// <li>5min: 5分钟；</li>
	// <li>hour: 1小时；</li>
	// <li>day: 1天。</li>不填将根据开始时间跟结束时间的间距自动推算粒度，具体为：2 小时范围内以 min 粒度查询，2 天范围内以 5min 粒度查询，7 天范围内以 hour 粒度查询，超过 7 天以 day 粒度查询。
	Interval *string `json:"Interval,omitnil,omitempty" name:"Interval"`

	// 筛选数据时使用的过滤条件，取值参考 [指标分析筛选条件说明](https://cloud.tencent.com/document/product/1552/98219#1aaf1150-55a4-4b4d-b103-3a8317ac7945) 中针对 L7 访问流量、带宽、请求数的可用筛选项。
	// 如需限定站点或内容标识符，请在 `ZoneIds.N` 参数中另行传入对应的值。
	Filters []*QueryCondition `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 数据归属地区。该参数已废弃。请在 `Filters.country` 中按客户端地域过滤数据。
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`
}

type DescribeTimingL7AnalysisDataRequest struct {
	*tchttp.BaseRequest
	
	// 开始时间。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间。查询时间范围（`EndTime` - `StartTime`）需小于等于 31 天。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 指标列表，取值有:
	// <li>l7Flow_outFlux: L7 EdgeOne 响应流量，单位：Byte；</li>
	// <li>l7Flow_inFlux: L7 客户端请求流量，单位：Byte；</li>
	// <li>l7Flow_flux: L7 访问总流量（EdgeOne 响应+客户端请求），单位：Byte；</li>
	// <li>l7Flow_outBandwidth: L7 EdgeOne 响应带宽，单位：bps；</li>
	// <li>l7Flow_inBandwidth：L7 客户端请求带宽，单位：bps；</li>
	// <li>l7Flow_bandwidth：L7 访问总带宽（EdgeOne 响应+客户端请求），单位：bps；</li>
	// <li>l7Flow_request: L7 访问请求数，单位：次；</li>
	// <li> l7Flow_avgResponseTime: L7 访问平均响应耗时，单位：ms；</li>
	// <li> l7Flow_avgFirstByteResponseTime: L7 访问平均首字节响应耗时，单位：ms。</li>
	MetricNames []*string `json:"MetricNames,omitnil,omitempty" name:"MetricNames"`

	// 站点 ID 集合，此参数将于2024年05月30日后由可选改为必填，详见公告：[【腾讯云 EdgeOne】云 API 变更通知](https://cloud.tencent.com/document/product/1552/104902)。最多传入 100 个站点 ID。若需查询腾讯云主账号下所有站点数据，请用 `*` 代替，查询账号级别数据需具备本接口全部站点资源权限。
	ZoneIds []*string `json:"ZoneIds,omitnil,omitempty" name:"ZoneIds"`

	// 查询时间粒度，取值有：
	// <li>min: 1分钟；</li>
	// <li>5min: 5分钟；</li>
	// <li>hour: 1小时；</li>
	// <li>day: 1天。</li>不填将根据开始时间跟结束时间的间距自动推算粒度，具体为：2 小时范围内以 min 粒度查询，2 天范围内以 5min 粒度查询，7 天范围内以 hour 粒度查询，超过 7 天以 day 粒度查询。
	Interval *string `json:"Interval,omitnil,omitempty" name:"Interval"`

	// 筛选数据时使用的过滤条件，取值参考 [指标分析筛选条件说明](https://cloud.tencent.com/document/product/1552/98219#1aaf1150-55a4-4b4d-b103-3a8317ac7945) 中针对 L7 访问流量、带宽、请求数的可用筛选项。
	// 如需限定站点或内容标识符，请在 `ZoneIds.N` 参数中另行传入对应的值。
	Filters []*QueryCondition `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 数据归属地区。该参数已废弃。请在 `Filters.country` 中按客户端地域过滤数据。
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
type DescribeTimingL7OriginPullDataRequestParams struct {
	// 开始时间。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间。查询时间范围（`EndTime` - `StartTime`）需小于等于 31 天。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 指标列表，取值有:
	// <li>l7Flow_outFlux_hy: EdgeOne 节点至源站方向的请求流量，单位：Byte；</li>
	// <li>l7Flow_outBandwidth_hy: EdgeOne 节点至源站方向的请求带宽，单位：bps；</li>
	// <li>l7Flow_request_hy: EdgeOne 节点至源站方向的请求数，单位：次。</li>
	// <li>l7Flow_inFlux_hy: 源站至 EdgeOne 节点方向的响应流量，单位：Byte；</li>
	// <li>l7Flow_inBandwidth_hy: 源站至 EdgeOne 节点方向的响应带宽，单位：bps；</li>
	MetricNames []*string `json:"MetricNames,omitnil,omitempty" name:"MetricNames"`

	// 站点 ID 集合，此参数必填。最多传入 100 个站点 ID。若需查询腾讯云主账号下所有站点数据，请用 `*` 代替，查询账号级别数据需具备本接口全部站点资源权限。
	ZoneIds []*string `json:"ZoneIds,omitnil,omitempty" name:"ZoneIds"`

	// 查询时间粒度，取值有：
	// <li>min: 1分钟；</li>
	// <li>5min: 5分钟；</li>
	// <li>hour: 1小时；</li>
	// <li>day: 1天。</li>不填将根据开始时间跟结束时间的间距自动推算粒度，具体为：2 小时范围内以 min 粒度查询，2 天范围内以 5min 粒度查询，7 天范围内以 hour 粒度查询，超过 7 天以 day 粒度查询。
	Interval *string `json:"Interval,omitnil,omitempty" name:"Interval"`

	// 过滤条件，详细的过滤条件如下：
	// <li>domain：客户端请求的域名。若按泛域名接入 EdgeOne，则数据中记录为泛域名，而不是具体域名。</li>
	Filters []*QueryCondition `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeTimingL7OriginPullDataRequest struct {
	*tchttp.BaseRequest
	
	// 开始时间。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间。查询时间范围（`EndTime` - `StartTime`）需小于等于 31 天。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 指标列表，取值有:
	// <li>l7Flow_outFlux_hy: EdgeOne 节点至源站方向的请求流量，单位：Byte；</li>
	// <li>l7Flow_outBandwidth_hy: EdgeOne 节点至源站方向的请求带宽，单位：bps；</li>
	// <li>l7Flow_request_hy: EdgeOne 节点至源站方向的请求数，单位：次。</li>
	// <li>l7Flow_inFlux_hy: 源站至 EdgeOne 节点方向的响应流量，单位：Byte；</li>
	// <li>l7Flow_inBandwidth_hy: 源站至 EdgeOne 节点方向的响应带宽，单位：bps；</li>
	MetricNames []*string `json:"MetricNames,omitnil,omitempty" name:"MetricNames"`

	// 站点 ID 集合，此参数必填。最多传入 100 个站点 ID。若需查询腾讯云主账号下所有站点数据，请用 `*` 代替，查询账号级别数据需具备本接口全部站点资源权限。
	ZoneIds []*string `json:"ZoneIds,omitnil,omitempty" name:"ZoneIds"`

	// 查询时间粒度，取值有：
	// <li>min: 1分钟；</li>
	// <li>5min: 5分钟；</li>
	// <li>hour: 1小时；</li>
	// <li>day: 1天。</li>不填将根据开始时间跟结束时间的间距自动推算粒度，具体为：2 小时范围内以 min 粒度查询，2 天范围内以 5min 粒度查询，7 天范围内以 hour 粒度查询，超过 7 天以 day 粒度查询。
	Interval *string `json:"Interval,omitnil,omitempty" name:"Interval"`

	// 过滤条件，详细的过滤条件如下：
	// <li>domain：客户端请求的域名。若按泛域名接入 EdgeOne，则数据中记录为泛域名，而不是具体域名。</li>
	Filters []*QueryCondition `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribeTimingL7OriginPullDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTimingL7OriginPullDataRequest) FromJsonString(s string) error {
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTimingL7OriginPullDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTimingL7OriginPullDataResponseParams struct {
	// 查询结果的总条数。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 回源时序数据列表。
	TimingDataRecords []*TimingDataRecord `json:"TimingDataRecords,omitnil,omitempty" name:"TimingDataRecords"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTimingL7OriginPullDataResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTimingL7OriginPullDataResponseParams `json:"Response"`
}

func (r *DescribeTimingL7OriginPullDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTimingL7OriginPullDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTopL7AnalysisDataRequestParams struct {
	// 开始时间。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间。查询时间范围（`EndTime` - `StartTime`）需小于等于 31 天。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 查询的指标，取值有：
	// <li> l7Flow_outFlux_country：按国家/地区维度统计 L7 EdgeOne 响应流量指标；</li>
	// <li> l7Flow_outFlux_province：按中国大陆境内省份维度统计 L7 EdgeOne 响应流量指标；</li>
	// <li> l7Flow_outFlux_statusCode：按状态码维度统计 L7 EdgeOne 响应流量指标；</li>
	// <li> l7Flow_outFlux_domain：按域名维度统计 L7 EdgeOne 响应流量指标；</li>
	// <li> l7Flow_outFlux_url：按 URL Path 维度统计 L7 EdgeOne 响应流量指标; </li>
	// <li> l7Flow_outFlux_resourceType：按资源类型维度统计 L7 EdgeOne 响应流量指标；</li>
	// <li> l7Flow_outFlux_sip：按客户端 IP 维度统计 L7 EdgeOne 响应流量指标；</li>
	// <li> l7Flow_outFlux_referers：按 Referer 维度统计 L7 EdgeOne 响应流量指标；</li>
	// <li> l7Flow_outFlux_ua_device：按设备类型维度统计 L7 EdgeOne 响应流量指标; </li>
	// <li> l7Flow_outFlux_ua_browser：按浏览器类型维度统计 L7 EdgeOne 响应流量指标；</li>
	// <li> l7Flow_outFlux_ua_os：按操作系统类型维度统计 L7 EdgeOne 响应流量指标；</li>
	// <li> l7Flow_outFlux_ua：按 User-Agent 维度统计 L7 EdgeOne 响应流量指标；</li>
	// <li> l7Flow_request_country：按国家/地区维度统计 L7 访问请求数指标；</li>
	// <li> l7Flow_request_province：按中国大陆境内省份维度统计 L7 访问请求数指标；</li>
	// <li> l7Flow_request_statusCode：按状态码维度统计 L7 访问请求数指标；</li>
	// <li> l7Flow_request_domain：按域名维度统计 L7 访问请求数指标；</li>
	// <li> l7Flow_request_url：按 URL Path 维度统计 L7 访问请求数指标; </li>
	// <li> l7Flow_request_resourceType：按资源类型维度统计 L7 访问请求数指标；</li>
	// <li> l7Flow_request_sip：按客户端 IP 维度统计 L7 访问请求数指标；</li>
	// <li> l7Flow_request_referers：按 Referer 维度统计 L7 访问请求数指标；</li>
	// <li> l7Flow_request_ua_device：按设备类型维度统计 L7 访问请求数指标; </li>
	// <li> l7Flow_request_ua_browser：按浏览器类型维度统计 L7 访问请求数指标；</li>
	// <li> l7Flow_request_ua_os：按操作系统类型维度统计 L7 访问请求数指标；</li>
	// <li> l7Flow_request_ua：按 User-Agent 维度统计 L7 访问请求数指标。</li>
	MetricName *string `json:"MetricName,omitnil,omitempty" name:"MetricName"`

	// 站点 ID 集合，此参数将于2024年05月30日后由可选改为必填，详见公告：[【腾讯云 EdgeOne】云 API 变更通知](https://cloud.tencent.com/document/product/1552/104902)。最多传入 100 个站点 ID。若需查询腾讯云主账号下所有站点数据，请用 `*` 代替，查询账号级别数据需具备本接口全部站点资源权限。
	ZoneIds []*string `json:"ZoneIds,omitnil,omitempty" name:"ZoneIds"`

	// 查询前多少个 top 数据，最大值为1000。不填默认为10，表示查询 top10 的数据。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 筛选数据时使用的过滤条件，取值参考 [指标分析筛选条件说明](https://cloud.tencent.com/document/product/1552/98219#1aaf1150-55a4-4b4d-b103-3a8317ac7945) 中针对 L7 访问流量、带宽、请求数的可用筛选项。
	// 如需限定站点或内容标识符，请在 `ZoneIds.N` 参数中另行传入对应的值。
	Filters []*QueryCondition `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 查询时间粒度，该参数无效，待废弃。
	Interval *string `json:"Interval,omitnil,omitempty" name:"Interval"`

	// 数据归属地区。该参数已废弃。请在 `Filters.country` 中按客户端地域过滤数据。
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`
}

type DescribeTopL7AnalysisDataRequest struct {
	*tchttp.BaseRequest
	
	// 开始时间。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间。查询时间范围（`EndTime` - `StartTime`）需小于等于 31 天。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 查询的指标，取值有：
	// <li> l7Flow_outFlux_country：按国家/地区维度统计 L7 EdgeOne 响应流量指标；</li>
	// <li> l7Flow_outFlux_province：按中国大陆境内省份维度统计 L7 EdgeOne 响应流量指标；</li>
	// <li> l7Flow_outFlux_statusCode：按状态码维度统计 L7 EdgeOne 响应流量指标；</li>
	// <li> l7Flow_outFlux_domain：按域名维度统计 L7 EdgeOne 响应流量指标；</li>
	// <li> l7Flow_outFlux_url：按 URL Path 维度统计 L7 EdgeOne 响应流量指标; </li>
	// <li> l7Flow_outFlux_resourceType：按资源类型维度统计 L7 EdgeOne 响应流量指标；</li>
	// <li> l7Flow_outFlux_sip：按客户端 IP 维度统计 L7 EdgeOne 响应流量指标；</li>
	// <li> l7Flow_outFlux_referers：按 Referer 维度统计 L7 EdgeOne 响应流量指标；</li>
	// <li> l7Flow_outFlux_ua_device：按设备类型维度统计 L7 EdgeOne 响应流量指标; </li>
	// <li> l7Flow_outFlux_ua_browser：按浏览器类型维度统计 L7 EdgeOne 响应流量指标；</li>
	// <li> l7Flow_outFlux_ua_os：按操作系统类型维度统计 L7 EdgeOne 响应流量指标；</li>
	// <li> l7Flow_outFlux_ua：按 User-Agent 维度统计 L7 EdgeOne 响应流量指标；</li>
	// <li> l7Flow_request_country：按国家/地区维度统计 L7 访问请求数指标；</li>
	// <li> l7Flow_request_province：按中国大陆境内省份维度统计 L7 访问请求数指标；</li>
	// <li> l7Flow_request_statusCode：按状态码维度统计 L7 访问请求数指标；</li>
	// <li> l7Flow_request_domain：按域名维度统计 L7 访问请求数指标；</li>
	// <li> l7Flow_request_url：按 URL Path 维度统计 L7 访问请求数指标; </li>
	// <li> l7Flow_request_resourceType：按资源类型维度统计 L7 访问请求数指标；</li>
	// <li> l7Flow_request_sip：按客户端 IP 维度统计 L7 访问请求数指标；</li>
	// <li> l7Flow_request_referers：按 Referer 维度统计 L7 访问请求数指标；</li>
	// <li> l7Flow_request_ua_device：按设备类型维度统计 L7 访问请求数指标; </li>
	// <li> l7Flow_request_ua_browser：按浏览器类型维度统计 L7 访问请求数指标；</li>
	// <li> l7Flow_request_ua_os：按操作系统类型维度统计 L7 访问请求数指标；</li>
	// <li> l7Flow_request_ua：按 User-Agent 维度统计 L7 访问请求数指标。</li>
	MetricName *string `json:"MetricName,omitnil,omitempty" name:"MetricName"`

	// 站点 ID 集合，此参数将于2024年05月30日后由可选改为必填，详见公告：[【腾讯云 EdgeOne】云 API 变更通知](https://cloud.tencent.com/document/product/1552/104902)。最多传入 100 个站点 ID。若需查询腾讯云主账号下所有站点数据，请用 `*` 代替，查询账号级别数据需具备本接口全部站点资源权限。
	ZoneIds []*string `json:"ZoneIds,omitnil,omitempty" name:"ZoneIds"`

	// 查询前多少个 top 数据，最大值为1000。不填默认为10，表示查询 top10 的数据。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 筛选数据时使用的过滤条件，取值参考 [指标分析筛选条件说明](https://cloud.tencent.com/document/product/1552/98219#1aaf1150-55a4-4b4d-b103-3a8317ac7945) 中针对 L7 访问流量、带宽、请求数的可用筛选项。
	// 如需限定站点或内容标识符，请在 `ZoneIds.N` 参数中另行传入对应的值。
	Filters []*QueryCondition `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 查询时间粒度，该参数无效，待废弃。
	Interval *string `json:"Interval,omitnil,omitempty" name:"Interval"`

	// 数据归属地区。该参数已废弃。请在 `Filters.country` 中按客户端地域过滤数据。
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

	// 七层访问数据按照 MetricName 指定统计维度的前 topN 数据列表。
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
type DescribeWebSecurityTemplateRequestParams struct {
	// 站点 ID。需要传入目标策略模板在访问权限上归属的站点，可使用 DescribeWebSecurityTemplates 接口查询策略模板归属的站点。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 策略模板 ID。
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`
}

type DescribeWebSecurityTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 站点 ID。需要传入目标策略模板在访问权限上归属的站点，可使用 DescribeWebSecurityTemplates 接口查询策略模板归属的站点。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 策略模板 ID。
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`
}

func (r *DescribeWebSecurityTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWebSecurityTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "TemplateId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeWebSecurityTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWebSecurityTemplateResponseParams struct {
	// 安全策略模板配置内容，Bot 配置暂不支持，正在开发中。
	SecurityPolicy *SecurityPolicy `json:"SecurityPolicy,omitnil,omitempty" name:"SecurityPolicy"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeWebSecurityTemplateResponse struct {
	*tchttp.BaseResponse
	Response *DescribeWebSecurityTemplateResponseParams `json:"Response"`
}

func (r *DescribeWebSecurityTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWebSecurityTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWebSecurityTemplatesRequestParams struct {
	// 站点 ID 列表。单次查询最多传入 100 个站点。
	ZoneIds []*string `json:"ZoneIds,omitnil,omitempty" name:"ZoneIds"`
}

type DescribeWebSecurityTemplatesRequest struct {
	*tchttp.BaseRequest
	
	// 站点 ID 列表。单次查询最多传入 100 个站点。
	ZoneIds []*string `json:"ZoneIds,omitnil,omitempty" name:"ZoneIds"`
}

func (r *DescribeWebSecurityTemplatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWebSecurityTemplatesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeWebSecurityTemplatesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWebSecurityTemplatesResponseParams struct {
	// 策略模板总数。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 策略模板列表。
	SecurityPolicyTemplates []*SecurityPolicyTemplateInfo `json:"SecurityPolicyTemplates,omitnil,omitempty" name:"SecurityPolicyTemplates"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeWebSecurityTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeWebSecurityTemplatesResponseParams `json:"Response"`
}

func (r *DescribeWebSecurityTemplatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWebSecurityTemplatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeZoneConfigImportResultRequestParams struct {
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 表示需要查询结果的导入配置任务 Id，导入任务 Id 仅支持查询最近 7 天的导入任务。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

type DescribeZoneConfigImportResultRequest struct {
	*tchttp.BaseRequest
	
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 表示需要查询结果的导入配置任务 Id，导入任务 Id 仅支持查询最近 7 天的导入任务。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

func (r *DescribeZoneConfigImportResultRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeZoneConfigImportResultRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeZoneConfigImportResultRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeZoneConfigImportResultResponseParams struct {
	// 本次导入任务的导入状态。取值有：  <li>success：表示配置项导入成功；</li> <li>failure：表示配置项导入失败；</li> <li>doing：表示配置项正在导入中。</li>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 本次导入任务的状态的提示信息。当配置项导入失败时，可通过本字段查看失败原因。
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 本次导入任务的配置内容。
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`

	// 本次导入任务的开始时间。
	ImportTime *string `json:"ImportTime,omitnil,omitempty" name:"ImportTime"`

	// 本次导入任务的结束时间。
	FinishTime *string `json:"FinishTime,omitnil,omitempty" name:"FinishTime"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeZoneConfigImportResultResponse struct {
	*tchttp.BaseResponse
	Response *DescribeZoneConfigImportResultResponseParams `json:"Response"`
}

func (r *DescribeZoneConfigImportResultResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeZoneConfigImportResultResponse) FromJsonString(s string) error {
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
	// <li>zone-name：按照站点名称进行过滤；</li><li>zone-type：按照站点类型进行过滤。可选项：<br>   full：NS 接入类型；<br>   partial：CNAME 接入类型；<br>   partialComposite：无域名接入类型；<br>   dnsPodAccess：DNSPod 托管接入类型；<br>   pages：Pages 类型。</li><li>zone-id：按照站点 ID 进行过滤，站点 ID 形如：zone-2noz78a8ev6k；</li><li>status：按照站点状态进行过滤。可选项：<br>   active：NS 已切换；<br>   pending：NS 待切换；<br>   deleted：已删除。</li><li>tag-key：按照标签键进行过滤；</li><li>tag-value： 按照标签值进行过滤；</li><li>alias-zone-name： 按照同名站点标识进行过滤。</li>模糊查询时支持过滤字段名为 zone-name 或 alias-zone-name。
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
	// <li>zone-name：按照站点名称进行过滤；</li><li>zone-type：按照站点类型进行过滤。可选项：<br>   full：NS 接入类型；<br>   partial：CNAME 接入类型；<br>   partialComposite：无域名接入类型；<br>   dnsPodAccess：DNSPod 托管接入类型；<br>   pages：Pages 类型。</li><li>zone-id：按照站点 ID 进行过滤，站点 ID 形如：zone-2noz78a8ev6k；</li><li>status：按照站点状态进行过滤。可选项：<br>   active：NS 已切换；<br>   pending：NS 待切换；<br>   deleted：已删除。</li><li>tag-key：按照标签键进行过滤；</li><li>tag-value： 按照标签值进行过滤；</li><li>alias-zone-name： 按照同名站点标识进行过滤。</li>模糊查询时支持过滤字段名为 zone-name 或 alias-zone-name。
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

	// 站点列表详情。
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

type DetectLengthLimitCondition struct {
	// 匹配条件的参数名称，取值有：
	// <li>body_depth：请求正文包部分的检测深度。</li>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 匹配条件的参数值，取值与 Name 成对使用。
	// 当 Name 值为 body_depth 时， Values 只支持传入单个值，取值有：
	// <li>10KB；</li>
	// <li>64KB；</li>
	// <li>128KB。</li>
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`
}

type DetectLengthLimitConfig struct {
	// 检测长度限制的规则列表。
	DetectLengthLimitRules []*DetectLengthLimitRule `json:"DetectLengthLimitRules,omitnil,omitempty" name:"DetectLengthLimitRules"`
}

type DetectLengthLimitRule struct {
	// 规则Id。仅出参使用。
	RuleId *uint64 `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// 规则名称。仅出参使用。
	RuleName *string `json:"RuleName,omitnil,omitempty" name:"RuleName"`

	// 规则描述，仅出参使用。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 规则配置条件。仅出参使用。
	Conditions []*DetectLengthLimitCondition `json:"Conditions,omitnil,omitempty" name:"Conditions"`

	// 处置方式，取值有：
	// <li>skip：当请求正文数据超过 Conditions 出参中 body_depth 设置的检测深度时，跳过所有请求正文内容的检测；</li>
	// <li>scan：仅检测 Conditions 出参中 body_depth 设置的检测深度，对超出部分的请求正文内容直接截断处理，超出部分的请求正文不会经过安全检测。</li>仅出参使用。
	Action *string `json:"Action,omitnil,omitempty" name:"Action"`
}

type DeviceProfile struct {
	// 客户端设备类型。取值有：<li>iOS；</li><li>Android；</li><li>WebView。</li>
	ClientType *string `json:"ClientType,omitnil,omitempty" name:"ClientType"`

	// 判定请求为高风险的最低值，取值范围为 1～99。数值越大请求风险越高越接近 Bot 客户端发起的请求。默认值为 50，对应含义 51～100 为高风险。
	HighRiskMinScore *uint64 `json:"HighRiskMinScore,omitnil,omitempty" name:"HighRiskMinScore"`

	// 高风险请求的处置方式。SecurityAction 的 Name 取值支持：<li>Deny：拦截；</li><li>Monitor：观察；</li><li>Redirect：重定向；</li><li>Challenge：挑战。</li>默认值为 Monitor。
	HighRiskRequestAction *SecurityAction `json:"HighRiskRequestAction,omitnil,omitempty" name:"HighRiskRequestAction"`

	// 判定请求为中风险的最低值，取值范围为 1～99。数值越大请求风险越高越接近 Bot 客户端发起的请求。默认值为 15，对应含义 16～50 为中风险。
	MediumRiskMinScore *uint64 `json:"MediumRiskMinScore,omitnil,omitempty" name:"MediumRiskMinScore"`

	// 中风险请求的处置方式。SecurityAction 的 Name 取值支持：<li>Deny：拦截；</li><li>Monitor：观察；</li><li>Redirect：重定向；</li><li>Challenge：挑战。</li>默认值为 Monitor。
	MediumRiskRequestAction *SecurityAction `json:"MediumRiskRequestAction,omitnil,omitempty" name:"MediumRiskRequestAction"`
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

// Predefined struct for user
type DisableOriginACLRequestParams struct {
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`
}

type DisableOriginACLRequest struct {
	*tchttp.BaseRequest
	
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`
}

func (r *DisableOriginACLRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisableOriginACLRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DisableOriginACLRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisableOriginACLResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DisableOriginACLResponse struct {
	*tchttp.BaseResponse
	Response *DisableOriginACLResponseParams `json:"Response"`
}

func (r *DisableOriginACLResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisableOriginACLResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DnsRecord struct {
	// 站点 ID。<br>注意：ZoneId 仅做出参使用，在 ModifyDnsRecords 不可作为入参使用，如有传此参数，会忽略。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// DNS 记录 ID。
	RecordId *string `json:"RecordId,omitnil,omitempty" name:"RecordId"`

	// DNS 记录名。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// DNS 记录类型，取值有：
	// <li>A：将域名指向一个外网 IPv4 地址，如 8.8.8.8；</li>
	// <li>AAAA：将域名指向一个外网 IPv6 地址；</li>
	// <li>MX：用于邮箱服务器。存在多条 MX 记录时，优先级越低越优先；</li>
	// <li>CNAME：将域名指向另一个域名，再由该域名解析出最终 IP 地址；</li>
	// <li>TXT：对域名进行标识和说明，常用于域名验证和 SPF 记录（反垃圾邮件）；</li>
	// <li>NS：如果需要将子域名交给其他 DNS 服务商解析，则需要添加 NS 记录。根域名无法添加 NS 记录；</li>
	// <li>CAA：指定可为本站点颁发证书的 CA；</li>
	// <li>SRV：标识某台服务器使用了某个服务，常见于微软系统的目录管理。</li>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// DNS 记录解析线路，不指定默认为 Default，表示默认解析线路，代表全部地域生效。<br>解析线路配置仅适用于当 Type（DNS 记录类型）为 A、AAAA、CNAME 时。<br>取值请参考：[解析线路及对应代码枚举](https://cloud.tencent.com/document/product/1552/112542)。
	Location *string `json:"Location,omitnil,omitempty" name:"Location"`

	// DNS 记录内容。根据 Type 值填入与之相对应的内容。
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`

	// 缓存时间，取值范围 60~86400，数值越小，修改记录各地生效时间越快，单位：秒。
	TTL *int64 `json:"TTL,omitnil,omitempty" name:"TTL"`

	// DNS 记录权重，取值范围 -1~100，为 -1 时表示不分配权重，为 0 时表示不解析。权重配置仅适用于当 Type（DNS 记录类型）为 A、AAAA、CNAME 时。
	Weight *int64 `json:"Weight,omitnil,omitempty" name:"Weight"`

	// MX 记录优先级，取值范围 0~50，数值越小越优先。
	Priority *int64 `json:"Priority,omitnil,omitempty" name:"Priority"`

	// DNS 记录解析状态，取值有：<li>enable：已生效；</li><li>disable：已停用。</li>注意：Status 仅做出参使用，在 ModifyDnsRecords 不可作为入参使用，如有传此参数，会忽略。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 创建时间。<br>注意：CreatedOn 仅做出参使用，在 ModifyDnsRecords 不可作为入参使用，如有传此参数，会忽略。
	CreatedOn *string `json:"CreatedOn,omitnil,omitempty" name:"CreatedOn"`

	// 修改时间。<br>注意：ModifiedOn 仅做出参使用，在 ModifyDnsRecords 不可作为入参使用，如有传此参数，会忽略。
	ModifiedOn *string `json:"ModifiedOn,omitnil,omitempty" name:"ModifiedOn"`
}

type DnsVerification struct {
	// 主机记录。
	Subdomain *string `json:"Subdomain,omitnil,omitempty" name:"Subdomain"`

	// 记录类型。
	RecordType *string `json:"RecordType,omitnil,omitempty" name:"RecordType"`

	// 记录值。
	RecordValue *string `json:"RecordValue,omitnil,omitempty" name:"RecordValue"`
}

type DomainDDoSProtection struct {
	// 域名。
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 域名的独立 DDoS 开关，取值为：
	// <li> on：已开启；</li>
	// <li> off：已关闭。</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`
}

// Predefined struct for user
type DownloadL4LogsRequestParams struct {
	// 开始时间。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 站点 ID 集合，此参数将于2024年05月30日后由可选改为必填，详见公告：[【腾讯云 EdgeOne】云 API 变更通知](https://cloud.tencent.com/document/product/1552/104902)。
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

	// 站点 ID 集合，此参数将于2024年05月30日后由可选改为必填，详见公告：[【腾讯云 EdgeOne】云 API 变更通知](https://cloud.tencent.com/document/product/1552/104902)。
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

	// 站点ID集合，此参数将于2024年05月30日后由可选改为必填，详见公告：[【腾讯云 EdgeOne】云 API 变更通知](https://cloud.tencent.com/document/product/1552/104902)。
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

	// 站点ID集合，此参数将于2024年05月30日后由可选改为必填，详见公告：[【腾讯云 EdgeOne】云 API 变更通知](https://cloud.tencent.com/document/product/1552/104902)。
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
	WafDropPageDetail *DropPageDetail `json:"WafDropPageDetail,omitnil,omitempty" name:"WafDropPageDetail"`

	// 自定义页面的拦截页面配置。如果为null，默认使用历史配置。
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

// Predefined struct for user
type EnableOriginACLRequestParams struct {
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 七层加速域名开启源站防护的模式。
	// <li>all：针对站点下的所有七层加速域名开启。</li>
	// <li>specific：针对站点下指定的七层加速域名开启。</li>当参数为空时，默认为 specific。
	L7EnableMode *string `json:"L7EnableMode,omitnil,omitempty" name:"L7EnableMode"`

	// 开启源站防护的七层加速域名列表，仅当参数 L7EnableMode 为 specific 时生效。L7EnableMode 为 all 时，请保留此参数为空。单次最大仅支持填写 200 个七层加速域名。
	L7Hosts []*string `json:"L7Hosts,omitnil,omitempty" name:"L7Hosts"`

	// 四层代理实例开启源站防护的模式。
	// <li>all：针对站点下的所有四层代理实例开启。</li>
	// <li>specific：针对站点下指定的四层代理实例开启。</li>当参数为空时，默认为 specific。
	L4EnableMode *string `json:"L4EnableMode,omitnil,omitempty" name:"L4EnableMode"`

	// 开启源站防护的四层代理实例列表，仅当参数 L4EnableMode 为 specific 时生效。L4EnableMode 为 all 时，请保留此参数为空。单次最大仅支持填写 100 个四层代理实例。
	L4ProxyIds []*string `json:"L4ProxyIds,omitnil,omitempty" name:"L4ProxyIds"`
}

type EnableOriginACLRequest struct {
	*tchttp.BaseRequest
	
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 七层加速域名开启源站防护的模式。
	// <li>all：针对站点下的所有七层加速域名开启。</li>
	// <li>specific：针对站点下指定的七层加速域名开启。</li>当参数为空时，默认为 specific。
	L7EnableMode *string `json:"L7EnableMode,omitnil,omitempty" name:"L7EnableMode"`

	// 开启源站防护的七层加速域名列表，仅当参数 L7EnableMode 为 specific 时生效。L7EnableMode 为 all 时，请保留此参数为空。单次最大仅支持填写 200 个七层加速域名。
	L7Hosts []*string `json:"L7Hosts,omitnil,omitempty" name:"L7Hosts"`

	// 四层代理实例开启源站防护的模式。
	// <li>all：针对站点下的所有四层代理实例开启。</li>
	// <li>specific：针对站点下指定的四层代理实例开启。</li>当参数为空时，默认为 specific。
	L4EnableMode *string `json:"L4EnableMode,omitnil,omitempty" name:"L4EnableMode"`

	// 开启源站防护的四层代理实例列表，仅当参数 L4EnableMode 为 specific 时生效。L4EnableMode 为 all 时，请保留此参数为空。单次最大仅支持填写 100 个四层代理实例。
	L4ProxyIds []*string `json:"L4ProxyIds,omitnil,omitempty" name:"L4ProxyIds"`
}

func (r *EnableOriginACLRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnableOriginACLRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "L7EnableMode")
	delete(f, "L7Hosts")
	delete(f, "L4EnableMode")
	delete(f, "L4ProxyIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "EnableOriginACLRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EnableOriginACLResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type EnableOriginACLResponse struct {
	*tchttp.BaseResponse
	Response *EnableOriginACLResponseParams `json:"Response"`
}

func (r *EnableOriginACLResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnableOriginACLResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
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

type ErrorPage struct {
	// 状态码。支持范围为 400、403、404、405、414、416、451、500、501、502、503、504。
	StatusCode *int64 `json:"StatusCode,omitnil,omitempty" name:"StatusCode"`

	// 重定向 URL，需要为完整跳转路径，如 https://www.test.com/error.html。
	RedirectURL *string `json:"RedirectURL,omitnil,omitempty" name:"RedirectURL"`
}

type ErrorPageParameters struct {
	// 自定义错误页面配置列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorPageParams []*ErrorPage `json:"ErrorPageParams,omitnil,omitempty" name:"ErrorPageParams"`
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
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 匹配条件。
	ExceptUserRuleConditions []*ExceptUserRuleCondition `json:"ExceptUserRuleConditions,omitnil,omitempty" name:"ExceptUserRuleConditions"`

	// 规则生效的范围。
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
	Modules []*string `json:"Modules,omitnil,omitempty" name:"Modules"`

	// 跳过部分规则ID的例外规则详情。如果为null，默认使用历史配置。
	PartialModules []*PartialModule `json:"PartialModules,omitnil,omitempty" name:"PartialModules"`

	// 跳过具体字段不去扫描的例外规则详情。如果为null，默认使用历史配置。
	SkipConditions []*SkipCondition `json:"SkipConditions,omitnil,omitempty" name:"SkipConditions"`
}

type ExceptionRule struct {
	// 例外规则的 ID。<br>通过规则 ID 可支持不同的规则配置操作：<br> <li> <b>增加</b>新规则：ID 为空或不指定 ID 参数；</li><li> <b>修改</b>已有规则：指定需要更新/修改的规则 ID；</li><li> <b>删除</b>已有规则：ExceptionRules 参数中，Rules 列表中未包含的已有规则将被删除。</li>
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 例外规则的名称。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 例外规则的具体内容，需符合表达式语法，详细规范参见产品文档。
	Condition *string `json:"Condition,omitnil,omitempty" name:"Condition"`

	// 例外规则执行选项，取值有：<li>WebSecurityModules: 指定例外规则的安全防护模块。</li><li>ManagedRules：指定托管规则。</li>
	SkipScope *string `json:"SkipScope,omitnil,omitempty" name:"SkipScope"`

	// 跳过请求的具体类型，取值有：<li>SkipOnAllRequestFields: 跳过所有请求；</li><li>SkipOnSpecifiedRequestFields: 跳过指定请求字段。</li>仅当 SkipScope 为 ManagedRules 时有效。
	SkipOption *string `json:"SkipOption,omitnil,omitempty" name:"SkipOption"`

	// 指定例外规则的安全防护模块，仅当 SkipScope 为 WebSecurityModules 时有效。取值有：<li>websec-mod-managed-rules：托管规则；</li><li>websec-mod-rate-limiting：速率限制；</li><li>websec-mod-custom-rules：自定义规则；</li><li>websec-mod-adaptive-control：自适应频控、智能客户端过滤、慢速攻击防护、流量盗刷防护；</li><li>websec-mod-bot：Bot管理。</li>
	WebSecurityModulesForException []*string `json:"WebSecurityModulesForException,omitnil,omitempty" name:"WebSecurityModulesForException"`

	// 指定例外规则的具体托管规则，仅当 SkipScope 为 ManagedRules 时有效，且此时不能指定 ManagedRuleGroupsForException 。
	ManagedRulesForException []*string `json:"ManagedRulesForException,omitnil,omitempty" name:"ManagedRulesForException"`

	// 指定例外规则的托管规则组，仅当 SkipScope 为 ManagedRules 时有效，且此时不能指定 ManagedRulesForException 。
	ManagedRuleGroupsForException []*string `json:"ManagedRuleGroupsForException,omitnil,omitempty" name:"ManagedRuleGroupsForException"`

	// 指定例外规则跳过指定请求字段的具体配置，仅当 SkipScope 为 ManagedRules 并且 SkipOption 为 SkipOnSpecifiedRequestFields 时有效。
	RequestFieldsForException []*RequestFieldsForException `json:"RequestFieldsForException,omitnil,omitempty" name:"RequestFieldsForException"`

	// 例外规则是否开启。取值有：<li>on：开启</li><li>off：关闭</li>
	Enabled *string `json:"Enabled,omitnil,omitempty" name:"Enabled"`
}

type ExceptionRules struct {
	// 例外规则的定义列表。使用 ModifySecurityPolicy 修改 Web 防护配置时: <li>若未指定 Rules 参数，或 Rules 参数长度为零：清空所有例外规则配置。</li><li>若 SecurityPolicy 参数中，未指定 ExceptionRules 参数值：保持已有例外规则配置，不做修改。</li>
	Rules []*ExceptionRule `json:"Rules,omitnil,omitempty" name:"Rules"`
}

// Predefined struct for user
type ExportZoneConfigRequestParams struct {
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 导出配置项的类型列表，不填表示导出所有类型的配置，当前支持的取值有：<li>L7AccelerationConfig：表示导出七层加速配置，对应控制台「站点加速-全局加速配置」和「站点加速-规则引擎」。</li>
	// 需注意：后续支持导出的类型会随着迭代增加，导出所有类型时需要注意导出文件大小，建议使用时指定需要导出的配置类型，以便控制请求响应包负载大小。
	Types []*string `json:"Types,omitnil,omitempty" name:"Types"`
}

type ExportZoneConfigRequest struct {
	*tchttp.BaseRequest
	
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 导出配置项的类型列表，不填表示导出所有类型的配置，当前支持的取值有：<li>L7AccelerationConfig：表示导出七层加速配置，对应控制台「站点加速-全局加速配置」和「站点加速-规则引擎」。</li>
	// 需注意：后续支持导出的类型会随着迭代增加，导出所有类型时需要注意导出文件大小，建议使用时指定需要导出的配置类型，以便控制请求响应包负载大小。
	Types []*string `json:"Types,omitnil,omitempty" name:"Types"`
}

func (r *ExportZoneConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportZoneConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "Types")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ExportZoneConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExportZoneConfigResponseParams struct {
	// 导出的配置的具体内容。以 JSON 格式返回，按照 UTF-8 方式进行编码。配置内容可参考下方示例。
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ExportZoneConfigResponse struct {
	*tchttp.BaseResponse
	Response *ExportZoneConfigResponseParams `json:"Response"`
}

func (r *ExportZoneConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportZoneConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
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
	// EdgeOne 后台服务器将通过 http://{Host}{URL Path} 的格式（例如 http://www.example.com/.well-known/teo-verification/z12h416twn.txt）获取文件验证信息。其中，本字段为您需要创建的 URL Path 部分，Host 为当前加速域名。
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
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// 首段包的统计时长，单位是秒，即期望首段包的统计时长是多少，默认5秒。
	StatTime *uint64 `json:"StatTime,omitnil,omitempty" name:"StatTime"`
}

type FollowOrigin struct {
	// 遵循源站配置开关，取值有：
	// <li>on：开启；</li>
	// <li>off：关闭。</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// 源站未返回 Cache-Control 头时，缓存/不缓存开关。当 Switch 为 on 时，此字段必填，当 Switch 为 off 时，无需填写此字段，若填写则不生效。取值有：
	// <li>on：缓存；</li>
	// <li>off：不缓存。</li>
	DefaultCache *string `json:"DefaultCache,omitnil,omitempty" name:"DefaultCache"`

	// 源站未返回 Cache-Control 头时，使用/不使用默认缓存策略开关。当 DefaultCache 为 on 时，此字段必填，否则此字段不生效；当 DefaultCacheTime 不为 0 时，此字段必须为 off。取值有：
	// <li>on：使用默认缓存策略；</li>
	// <li>off：不使用默认缓存策略。</li>
	DefaultCacheStrategy *string `json:"DefaultCacheStrategy,omitnil,omitempty" name:"DefaultCacheStrategy"`

	// 源站未返回 Cache-Control 头时，表示默认的缓存时间，单位为秒，取值：0-315360000。当 DefaultCache 为 on 时，此字段必填，否则此字段不生效；当 DefaultCacheStrategy 为 on 时， 此字段必须为 0。
	DefaultCacheTime *int64 `json:"DefaultCacheTime,omitnil,omitempty" name:"DefaultCacheTime"`
}

type ForceRedirect struct {
	// 访问强制跳转配置开关，取值有：
	// <li>on：开启；</li>
	// <li>off：关闭。</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// 重定向状态码，取值有：
	// <li>301：301跳转；</li>
	// <li>302：302跳转。</li>
	RedirectStatusCode *int64 `json:"RedirectStatusCode,omitnil,omitempty" name:"RedirectStatusCode"`
}

type ForceRedirectHTTPSParameters struct {
	// 访问强制跳转配置开关，取值有：
	// <li>on：开启；</li>
	// <li>off：关闭。</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// 重定向状态码。当 Switch 为 on 时，此字段必填，否则此字段不生效。取值有：
	// <li>301：301跳转；</li>
	// <li>302：302跳转。</li>
	RedirectStatusCode *int64 `json:"RedirectStatusCode,omitnil,omitempty" name:"RedirectStatusCode"`
}

type FrequentScanningProtection struct {
	// 高频扫描防护规则是否开启。取值有：<li>on：开启，高频扫描防护规则生效；</li><li>off：关闭，高频扫描防护规则不生效。</li>	
	Enabled *string `json:"Enabled,omitnil,omitempty" name:"Enabled"`

	// 高频扫描防护的处置动作。 当 Enabled 为 on 时，此字段必填。SecurityAction 的 Name 取值支持：<li>Deny：拦截，响应拦截页面；</li><li>Monitor：观察，不处理请求记录安全事件到日志中；</li><li>JSChallenge：JavaScript 挑战，响应 JavaScript 挑战页面。</li>
	Action *SecurityAction `json:"Action,omitnil,omitempty" name:"Action"`

	// 请求统计的匹配方式，当 Enabled 为 on 时，此字段必填。取值有：<li>http.request.xff_header_ip：客户端 IP（优先匹配 XFF 头部）；</li><li>http.request.ip：客户端 IP。</li> 
	CountBy *string `json:"CountBy,omitnil,omitempty" name:"CountBy"`

	// 此参数指定高频扫描防护的阈值，即在 CountingPeriod 所设置时间范围内命中「配置为拦截」的托管规则时的累计拦截次数，取值范围 1 ~ 4294967294，例如 100，当超过此统计值时，后续请求将触发 Action 所设置的处置动作。当 Enabled 为 on 时，此字段必填。
	BlockThreshold *int64 `json:"BlockThreshold,omitnil,omitempty" name:"BlockThreshold"`

	// 此参数指定高频扫描防护所统计的时间窗口，即命中「配置为拦截」的托管规则的请求的统计时间窗口，取值 5 ~ 1800，单位仅支持秒（s），例如 5s。 当 Enabled 为 on 时，此字段必填。
	CountingPeriod *string `json:"CountingPeriod,omitnil,omitempty" name:"CountingPeriod"`

	// 此参数指定高频扫描防护 Action 参数所设置处置动作的持续时长，取值范围 60 ~ 86400，单位仅支持秒（s），例如 60s。当 Enabled 为 on 时，此字段必填。
	ActionDuration *string `json:"ActionDuration,omitnil,omitempty" name:"ActionDuration"`
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

type FunctionRegionSelection struct {
	// 函数 ID 。
	FunctionId *string `json:"FunctionId,omitnil,omitempty" name:"FunctionId"`

	// 国家/地区列表。示例值：CN：中国，CN.GD：中国广东。取值请参考：[国家/地区及对应代码枚举](https://cloud.tencent.com/document/product/1552/112542)。
	Regions []*string `json:"Regions,omitnil,omitempty" name:"Regions"`
}

type FunctionRule struct {
	// 规则ID。
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// 规则条件列表，列表项之间为或关系。
	FunctionRuleConditions []*FunctionRuleCondition `json:"FunctionRuleConditions,omitnil,omitempty" name:"FunctionRuleConditions"`

	// 函数选择配置类型：
	// <li> direct：直接指定执行函数；</li>
	// <li> weight：基于权重比选择函数；</li>
	// <li> region：基于客户端 IP 的国家/地区选择函数。</li>
	TriggerType *string `json:"TriggerType,omitnil,omitempty" name:"TriggerType"`

	// 指定执行的函数 ID。当 TriggerType 为 direct 时有效。
	FunctionId *string `json:"FunctionId,omitnil,omitempty" name:"FunctionId"`

	// 指定执行的函数名称。
	FunctionName *string `json:"FunctionName,omitnil,omitempty" name:"FunctionName"`

	// 基于客户端 IP 国家/地区的函数选择配置。
	RegionMappingSelections []*FunctionRegionSelection `json:"RegionMappingSelections,omitnil,omitempty" name:"RegionMappingSelections"`

	// 基于权重的函数选择配置。
	WeightedSelections []*FunctionWeightedSelection `json:"WeightedSelections,omitnil,omitempty" name:"WeightedSelections"`

	// 函数触发规则优先级，数值越大，优先级越高。
	Priority *int64 `json:"Priority,omitnil,omitempty" name:"Priority"`

	// 规则描述。
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 创建时间。时间为世界标准时间（UTC）， 遵循 ISO 8601 标准的日期和时间格式。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 更新时间。时间为世界标准时间（UTC）， 遵循 ISO 8601 标准的日期和时间格式。
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`
}

type FunctionRuleCondition struct {
	// 边缘函数触发规则条件，该列表内所有项全部满足即判断该条件满足。
	RuleConditions []*RuleCondition `json:"RuleConditions,omitnil,omitempty" name:"RuleConditions"`
}

type FunctionWeightedSelection struct {
	// 函数 ID 。
	FunctionId *string `json:"FunctionId,omitnil,omitempty" name:"FunctionId"`

	// 选中权重。取值范围0-100，所有的权重之和需要为100。
	// 选中概率计算方式为：
	// weight/100。例如设置了两个函数 A 和 B ，其中 A 的权重为30，那么 B 的权重必须为70，最终选中 A 的概率为30%，选中 B 的概率为70%。
	Weight *uint64 `json:"Weight,omitnil,omitempty" name:"Weight"`
}

type GatewayRegion struct {
	// 地域 ID 。
	RegionId *string `json:"RegionId,omitnil,omitempty" name:"RegionId"`

	// 中文地域名称。
	CNName *string `json:"CNName,omitnil,omitempty" name:"CNName"`

	// 英文地域名称。
	ENName *string `json:"ENName,omitnil,omitempty" name:"ENName"`
}

type Grpc struct {
	// 是否开启 Grpc 配置，取值有：
	// <li>on：开启；</li>
	// <li>off：关闭。</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`
}

type GrpcParameters struct {
	// gRPC 配置开关，取值有：
	// <li>on：开启；</li>
	// <li>off：关闭。</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`
}

type HSTSParameters struct {
	// HSTS 配置开关，取值有：
	// <li>on：开启；</li>
	// <li>off：关闭。</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// 缓存 HSTS 头部时间，单位为秒，取值：1-31536000。<br>注意：当 Switch 为 on 时，此字段必填；当 Switch 为 off 时，无需填写此字段，若填写则不生效。
	Timeout *int64 `json:"Timeout,omitnil,omitempty" name:"Timeout"`

	// 是否允许其他子域名继承相同的 HSTS 头部，取值有：
	// <li>on：允许其他子域名继承相同的 HSTS 头部；</li>
	// <li>off：不允许其他子域名继承相同的 HSTS 头部。</li>注意：当 Switch 为 on 时，此字段必填；当 Switch 为 off 时，无需填写此字段，若填写则不生效。
	IncludeSubDomains *string `json:"IncludeSubDomains,omitnil,omitempty" name:"IncludeSubDomains"`

	// 是否允许浏览器预加载 HSTS 头部，取值有：
	// <li>on：允许浏览器预加载 HSTS 头部；</li>
	// <li>off：不允许浏览器预加载 HSTS 头部。</li>注意：当 Switch 为 on 时，此字段必填；当 Switch 为 off 时，无需填写此字段，若填写则不生效。
	Preload *string `json:"Preload,omitnil,omitempty" name:"Preload"`
}

type HTTP2Parameters struct {
	// HTTP2 接入配置开关，取值有：
	// <li>on：开启；</li>
	// <li>off：关闭。</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`
}

type HTTPResponseParameters struct {
	// 响应状态码。支持 2XX、4XX、5XX，不包括 499、514、101、301、302、303、509、520-599。
	StatusCode *int64 `json:"StatusCode,omitnil,omitempty" name:"StatusCode"`

	// 响应页面 ID。
	ResponsePage *string `json:"ResponsePage,omitnil,omitempty" name:"ResponsePage"`
}

type HTTPUpstreamTimeoutParameters struct {
	// HTTP 应答超时时间，单位为秒，取值：5～600。
	ResponseTimeout *int64 `json:"ResponseTimeout,omitnil,omitempty" name:"ResponseTimeout"`
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

type HeaderAction struct {
	// HTTP 头部设置方式。取值有：
	// <li>set：设置。变更指定头部参数的取值为设置后的值；</li>
	// <li>del：删除。删除指定的头部参数；</li>
	// <li>add：增加。增加指定的头部参数。</li>
	Action *string `json:"Action,omitnil,omitempty" name:"Action"`

	// HTTP 头部名称。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// HTTP 头部值。当 Action 取值为 set 或者 add 时，该参数必填；当 Action 取值为 del 时，该参数无需填写。
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type HealthChecker struct {
	// 健康检查策略，取值有：
	// <li>HTTP；</li>
	// <li>HTTPS；</li>
	// <li>TCP；</li>
	// <li>UDP；</li>
	// <li>ICMP Ping；</li>
	// <li>NoCheck。</li>
	// 注意：NoCheck 表示不启用健康检查策略。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 检查端口。当 Type=HTTP 或 Type=HTTPS 或 Type=TCP 或 Type=UDP 时为必填。
	Port *uint64 `json:"Port,omitnil,omitempty" name:"Port"`

	// 检查频率，表示多久发起一次健康检查任务，单位为秒。可取值有：30，60，180，300 或 600。
	Interval *uint64 `json:"Interval,omitnil,omitempty" name:"Interval"`

	// 每一次健康检查的超时时间，若健康检查消耗时间大于此值，则检查结果判定为”不健康“， 单位为秒，默认值为 5s，取值必须小于 Interval。
	Timeout *uint64 `json:"Timeout,omitnil,omitempty" name:"Timeout"`

	// 健康阈值，表示连续几次健康检查结果为"健康"，则判断源站为"健康"，单位为次，默认 3 次，最小取值 1 次。
	HealthThreshold *uint64 `json:"HealthThreshold,omitnil,omitempty" name:"HealthThreshold"`

	// 不健康阈值，表示连续几次健康检查结果为"不健康"，则判断源站为"不健康"，单位为次，默认 2 次。
	CriticalThreshold *uint64 `json:"CriticalThreshold,omitnil,omitempty" name:"CriticalThreshold"`

	// 该参数仅当 Type=HTTP 或 Type=HTTPS 时有效，表示探测路径，需要填写完整的 host/path，不包含协议部分，例如：www.example.com/test。
	Path *string `json:"Path,omitnil,omitempty" name:"Path"`

	// 该参数仅当 Type=HTTP 或 Type=HTTPS 时有效，表示请求方法，取值有：
	// <li>GET；</li>
	// <li>HEAD。</li>
	Method *string `json:"Method,omitnil,omitempty" name:"Method"`

	// 该参数仅当 Type=HTTP 或 Type=HTTPS 时有效，表示探测节点向源站发起健康检查时，响应哪些状态码可用于认定探测结果为健康。
	ExpectedCodes []*string `json:"ExpectedCodes,omitnil,omitempty" name:"ExpectedCodes"`

	// 该参数仅当 Type=HTTP 或 Type=HTTPS 时有效，表示探测请求携带的自定义  HTTP 请求头，至多可配置 10 个。
	Headers []*CustomizedHeader `json:"Headers,omitnil,omitempty" name:"Headers"`

	// 该参数仅当 Type=HTTP 或 Type=HTTPS 时有效，表示是否启用遵循 301/302 重定向。启用后，301/302 默认为"健康"的状态码，默认跳转 3 次。
	FollowRedirect *string `json:"FollowRedirect,omitnil,omitempty" name:"FollowRedirect"`

	// 该参数仅当 Type=UDP 时有效，表示健康检查发送的内容。只允许 ASCII 可见字符，最大长度限制 500 个字符。
	SendContext *string `json:"SendContext,omitnil,omitempty" name:"SendContext"`

	// 该参数仅当 Type=UDP 时有效，表示健康检查期望源站返回结果。只允许 ASCII 可见字符，最大长度限制 500 个字符。
	RecvContext *string `json:"RecvContext,omitnil,omitempty" name:"RecvContext"`
}

type HostHeaderParameters struct {
	// 执行动作，取值有：
	// <li>followOrigin：跟随源站域名；</li>
	// <li>custom：自定义。</li>
	Action *string `json:"Action,omitnil,omitempty" name:"Action"`

	// Host Header 重写，需要填写完整域名。<br>注意：当 Switch 为 on 时，此字段必填；当 Switch 为 off 时，无需填写此字段，若填写则不生效。
	ServerName *string `json:"ServerName,omitnil,omitempty" name:"ServerName"`
}

type HostName struct {
	// 目标 HostName 配置，取值有：
	// <li>follow：跟随请求；</li>
	// <li>custom：自定义。</li>
	Action *string `json:"Action,omitnil,omitempty" name:"Action"`

	// 目标 HostName 自定义取值，最大长度 1024。<br>注意：当 Action 为 custom 时，此字段必填；当 Action 为 follow 时，此字段不生效。
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type Hsts struct {
	// 是否开启，取值有：
	// <li>on：开启；</li>
	// <li>off：关闭。</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// MaxAge 数值。单位为秒，最大值为1天。
	MaxAge *int64 `json:"MaxAge,omitnil,omitempty" name:"MaxAge"`

	// 是否包含子域名，取值有：
	// <li>on：开启；</li>
	// <li>off：关闭。</li>
	IncludeSubDomains *string `json:"IncludeSubDomains,omitnil,omitempty" name:"IncludeSubDomains"`

	// 是否开启预加载，取值有：
	// <li>on：开启；</li>
	// <li>off：关闭。</li>
	Preload *string `json:"Preload,omitnil,omitempty" name:"Preload"`
}

type HttpDDoSProtection struct {
	// 自适应频控的具体配置。
	AdaptiveFrequencyControl *AdaptiveFrequencyControl `json:"AdaptiveFrequencyControl,omitnil,omitempty" name:"AdaptiveFrequencyControl"`

	// 智能客户端过滤的具体配置。
	ClientFiltering *ClientFiltering `json:"ClientFiltering,omitnil,omitempty" name:"ClientFiltering"`

	// 流量防盗刷的具体配置。
	BandwidthAbuseDefense *BandwidthAbuseDefense `json:"BandwidthAbuseDefense,omitnil,omitempty" name:"BandwidthAbuseDefense"`

	// 慢速攻击防护的具体配置。
	SlowAttackDefense *SlowAttackDefense `json:"SlowAttackDefense,omitnil,omitempty" name:"SlowAttackDefense"`
}

type Https struct {
	// http2 配置开关，取值有：
	// <li>on：开启；</li>
	// <li>off：关闭。</li>
	Http2 *string `json:"Http2,omitnil,omitempty" name:"Http2"`

	// OCSP 配置开关，取值有：
	// <li>on：开启；</li>
	// <li>off：关闭。</li>
	OcspStapling *string `json:"OcspStapling,omitnil,omitempty" name:"OcspStapling"`

	// Tls 版本设置，取值有：
	// <li>TLSv1：TLSv1版本；</li>
	// <li>TLSv1.1：TLSv1.1版本；</li>
	// <li>TLSv1.2：TLSv1.2版本；</li>
	// <li>TLSv1.3：TLSv1.3版本。</li>修改时必须开启连续的版本。
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
	ApplyType *string `json:"ApplyType,omitnil,omitempty" name:"ApplyType"`

	// 密码套件，取值有：
	// <li>loose-v2023：提供高兼容性，安全性一般，支持 TLS 1.0-1.3 密码套件；</li>
	// <li>general-v2023：提供较高兼容性，安全性中等，支持 TLS 1.2-1.3 密码套件；</li>
	// <li>strict-v2023：提供高安全性能，禁用所有含不安全隐患的加密套件，支持 TLS 1.2-1.3 密码套件。</li>
	CipherSuite *string `json:"CipherSuite,omitnil,omitempty" name:"CipherSuite"`
}

type IPExpireInfo struct {
	// 定时过期时间，遵循 ISO 8601 标准的日期和时间格式。例如 "2022-01-01T00:00:00+08:00"。
	ExpireTime *string `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`

	// IP 列表。仅支持 IP  及 IP 网段。
	IPList []*string `json:"IPList,omitnil,omitempty" name:"IPList"`
}

type IPGroup struct {
	// 组 Id，创建时填 0 即可。
	GroupId *int64 `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 组名称。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// IP 组内容，仅支持 IP 及 IP 网段。
	Content []*string `json:"Content,omitnil,omitempty" name:"Content"`

	// IP 组中正在生效的 IP 或网段个数。作为出参时有效，作为入参时无需填写该字段。
	IPTotalCount *int64 `json:"IPTotalCount,omitnil,omitempty" name:"IPTotalCount"`

	// IP 定时过期信息。
	// 作为入参，用于为指定的 IP 地址或网段配置定时过期时间。
	// 作为出参，包含以下两类信息：
	// <li>当前未到期的定时过期信息：尚未触发的过期配置。</li>
	// <li>一周内已到期的定时过期信息：已触发的过期配置。</li>
	IPExpireInfo []*IPExpireInfo `json:"IPExpireInfo,omitnil,omitempty" name:"IPExpireInfo"`
}

type IPRegionInfo struct {
	// IP 地址，IPV4 或 IPV6。
	IP *string `json:"IP,omitnil,omitempty" name:"IP"`

	// IP 是否属于 EdgeOne 节点，取值有：
	// <li>yes：该 IP 属于 EdgeOne 节点；</li>
	// <li>no：该 IP 不属于 EdgeOne 节点。</li>
	IsEdgeOneIP *string `json:"IsEdgeOneIP,omitnil,omitempty" name:"IsEdgeOneIP"`
}

type IPReputation struct {
	// IP 情报库（原客户端画像分析）。取值有：<li>on：开启；</li><li>off：关闭。</li>
	Enabled *string `json:"Enabled,omitnil,omitempty" name:"Enabled"`

	// IP 情报库（原客户端画像分析）的具体配置内容。
	IPReputationGroup *IPReputationGroup `json:"IPReputationGroup,omitnil,omitempty" name:"IPReputationGroup"`
}

type IPReputationGroup struct {
	// IP 情报库（原客户端画像分析）的执行动作。SecurityAction 的 Name 取值支持：<li>Deny：拦截；</li><li>Monitor：观察；</li><li>Disabled：未启用，不启用指定规则；</li><li>Challenge：挑战，其中 ChallengeActionParameters 中的 ChallengeOption 支持 JSChallenge 和 ManagedChallenge。</li>
	BaseAction *SecurityAction `json:"BaseAction,omitnil,omitempty" name:"BaseAction"`

	// IP 情报库（原客户端画像分析）的具体配置，用于覆盖 BaseAction 中的默认配置。其中 BotManagementActionOverrides 的 Ids 中可以填写：<li>IPREP_WEB_AND_DDOS_ATTACKERS_LOW：网络攻击 - 一般置信度；</li><li>IPREP_WEB_AND_DDOS_ATTACKERS_MID：网络攻击 - 中等置信度；</li><li>IPREP_WEB_AND_DDOS_ATTACKERS_HIGH：网络攻击 - 高置信度；</li><li>IPREP_PROXIES_AND_ANONYMIZERS_LOW：网络代理 - 一般置信度；</li><li>IPREP_PROXIES_AND_ANONYMIZERS_MID：网络代理 - 中等置信度；</li><li>IPREP_PROXIES_AND_ANONYMIZERS_HIGH：网络代理 - 高置信度；</li><li>IPREP_SCANNING_TOOLS_LOW：扫描器 - 一般置信度；</li><li>IPREP_SCANNING_TOOLS_MID：扫描器 - 中等置信度；</li><li>IPREP_SCANNING_TOOLS_HIGH：扫描器 - 高置信度；</li><li>IPREP_ATO_ATTACKERS_LOW：账号接管攻击 - 一般置信度；</li><li>IPREP_ATO_ATTACKERS_MID：账号接管攻击 - 中等置信度；</li><li>IPREP_ATO_ATTACKERS_HIGH：账号接管攻击 - 高置信度；</li><li>IPREP_WEB_SCRAPERS_AND_TRAFFIC_BOTS_LOW：恶意 BOT - 一般置信度；</li><li>IPREP_WEB_SCRAPERS_AND_TRAFFIC_BOTS_MID：恶意 BOT - 中等置信度；</li><li>IPREP_WEB_SCRAPERS_AND_TRAFFIC_BOTS_HIGH：恶意 BOT - 高置信度。</li>
	BotManagementActionOverrides []*BotManagementActionOverrides `json:"BotManagementActionOverrides,omitnil,omitempty" name:"BotManagementActionOverrides"`
}

type IPWhitelist struct {
	// IPv4列表。
	IPv4 []*string `json:"IPv4,omitnil,omitempty" name:"IPv4"`

	// IPv6列表。
	IPv6 []*string `json:"IPv6,omitnil,omitempty" name:"IPv6"`
}

type IPv6Parameters struct {
	// IPv6 访问功能配置，取值有：
	// <li>on：开启 IPv6 访问功能；</li>
	// <li>off：关闭 IPv6 访问功能。</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`
}

type Identification struct {
	// 站点名称。
	ZoneName *string `json:"ZoneName,omitnil,omitempty" name:"ZoneName"`

	// 验证子域名。验证站点时，该值为空。验证子域名是为具体子域名。
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 验证状态，取值有：
	// <li> pending：验证中；</li>
	// <li> finished：验证完成。</li>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 站点归属权校验：Dns校验信息。
	Ascription *AscriptionInfo `json:"Ascription,omitnil,omitempty" name:"Ascription"`

	// 域名当前的 NS 记录。
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
type ImportZoneConfigRequestParams struct {
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 待导入的配置内容。要求采用 JSON 格式，按照 UTF-8 方式进行编码。配置内容可通过站点配置导出接口（ExportZoneConfig）获取。您可以单独导入「站点加速-全局加速配置」或「站点加速-规则引擎」，传入对应的字段即可，详情可以参考下方示例。
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`
}

type ImportZoneConfigRequest struct {
	*tchttp.BaseRequest
	
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 待导入的配置内容。要求采用 JSON 格式，按照 UTF-8 方式进行编码。配置内容可通过站点配置导出接口（ExportZoneConfig）获取。您可以单独导入「站点加速-全局加速配置」或「站点加速-规则引擎」，传入对应的字段即可，详情可以参考下方示例。
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`
}

func (r *ImportZoneConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ImportZoneConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "Content")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ImportZoneConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ImportZoneConfigResponseParams struct {
	// 表示该次导入配置的任务 Id，通过查询站点配置导入结果接口（DescribeZoneConfigImportResult）获取本次导入任务执行的结果。注意：导入任务 Id 仅支持查询最近 7 天的导入任务。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ImportZoneConfigResponse struct {
	*tchttp.BaseResponse
	Response *ImportZoneConfigResponseParams `json:"Response"`
}

func (r *ImportZoneConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ImportZoneConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
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
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// 规则详情。
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
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// 基础管控规则。如果为null，默认使用历史配置。
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

type JITVideoProcess struct {
	// 视频即时处理配置开关，取值有：
	// <li>on：开启；</li>
	// <li>off：关闭。</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`
}

type JSInjectionRule struct {
	// 规则 ID。
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// 规则名称。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 规则优先级，数值越小越优先执行，范围是 0 ~ 100，默认为 0。
	Priority *int64 `json:"Priority,omitnil,omitempty" name:"Priority"`

	// 匹配条件内容。需符合表达式语法，详细规范参见产品文档。
	Condition *string `json:"Condition,omitnil,omitempty" name:"Condition"`

	// JavaScript 注入选项。默认值为 run-attestations，取值有：
	// <li> no-injection: 不注入 JavaScript;</li>
	// <li> inject-sdk-only: 注入当前支持的所有认证方式的 SDK，当前支持：TC-RCE 和 TC-CAPTCHA。注意：若需执行认证检测，请配置挑战规则。</li>
	InjectJS *string `json:"InjectJS,omitnil,omitempty" name:"InjectJS"`
}

type JustInTimeTranscodeTemplate struct {
	// 即时转码模板唯一标识。
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// 转码模板名称。
	TemplateName *string `json:"TemplateName,omitnil,omitempty" name:"TemplateName"`

	// 模板描述信息。
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// 模板类型，取值：<li>preset：系统预置模板；</li><li>custom：用户自定义模板。</li>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 启用视频流开关，取值：<li>on：开启；</li><li>off：关闭。</li>
	VideoStreamSwitch *string `json:"VideoStreamSwitch,omitnil,omitempty" name:"VideoStreamSwitch"`

	// 启用音频流开关，取值：<li>on：开启；</li><li>off：关闭。</li>
	AudioStreamSwitch *string `json:"AudioStreamSwitch,omitnil,omitempty" name:"AudioStreamSwitch"`

	// 视频流配置参数，仅当 VideoStreamSwitch 为 on，该字段有效。
	VideoTemplate *VideoTemplateInfo `json:"VideoTemplate,omitnil,omitempty" name:"VideoTemplate"`

	// 音频流配置参数，仅当 AudioStreamSwitch 为 on，该字段有效。
	AudioTemplate *AudioTemplateInfo `json:"AudioTemplate,omitnil,omitempty" name:"AudioTemplate"`

	// 模板创建时间，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#I)。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 模板最后修改时间，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#I)。
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`
}

type KnownBotCategories struct {
	// 来自已知商业工具或开源工具的访问请求的处置方式。 SecurityAction 的 Name 取值支持：<li>Deny：拦截；</li><li>Monitor：观察；</li><li>Disabled：未启用，不启用指定规则；</li><li>Challenge：挑战，其中 ChallengeActionParameters 中的 ChallengeOption 支持 JSChallenge 和 ManagedChallenge；</li><li>Allow：放行（待废弃）。</li> 
	BaseAction *SecurityAction `json:"BaseAction,omitnil,omitempty" name:"BaseAction"`

	// 指定已知商业工具或开源工具的访问请求的处置方式。
	BotManagementActionOverrides []*BotManagementActionOverrides `json:"BotManagementActionOverrides,omitnil,omitempty" name:"BotManagementActionOverrides"`
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

type LoadBalancer struct {
	// 实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 实例名称，可输入 1-200 个字符，允许字符为 a-z，A-Z，0-9，_，-。	
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 实例类型，取值有：
	// <li>HTTP：HTTP 专用型，支持添加 HTTP 专用型和通用型源站组，仅支持被站点加速相关服务引用（如域名服务和规则引擎）；</li>
	// <li>GENERAL：通用型，仅支持添加通用型源站组，能被站点加速服务（如域名服务和规则引擎）和四层代理引用。</li>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 健康检查策略。详情请参考 [健康检查策略介绍](https://cloud.tencent.com/document/product/1552/104228)。
	HealthChecker *HealthChecker `json:"HealthChecker,omitnil,omitempty" name:"HealthChecker"`

	// 源站组间的流量调度策略，取值有：
	// <li>Pritory：按优先级顺序进行故障转移 。</li>
	SteeringPolicy *string `json:"SteeringPolicy,omitnil,omitempty" name:"SteeringPolicy"`

	// 实际访问某源站失败时的请求重试策略，详情请参考 [请求重试策略介绍](https://cloud.tencent.com/document/product/1552/104227)，取值有：
	// <li>OtherOriginGroup：单次请求失败后，请求优先重试下一优先级源站组；</li>
	// <li>OtherRecordInOriginGroup：单次请求失败后，请求优先重试同源站组内的其他源站。</li>
	FailoverPolicy *string `json:"FailoverPolicy,omitnil,omitempty" name:"FailoverPolicy"`

	// 源站组健康状态。
	OriginGroupHealthStatus []*OriginGroupHealthStatus `json:"OriginGroupHealthStatus,omitnil,omitempty" name:"OriginGroupHealthStatus"`

	// 负载均衡状态，取值有：
	// <li>Pending：部署中；</li>
	// <li>Deleting：删除中；</li>
	// <li>Running：已生效。</li>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 该负载均衡实例绑的四层代理实例的列表。
	L4UsedList []*string `json:"L4UsedList,omitnil,omitempty" name:"L4UsedList"`

	// 该负载均衡实例绑定的七层域名列表。
	L7UsedList []*string `json:"L7UsedList,omitnil,omitempty" name:"L7UsedList"`

	// 负载均衡被引用实例的列表。
	References []*OriginGroupReference `json:"References,omitnil,omitempty" name:"References"`
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

type ManagedRuleAction struct {
	// 托管规则组下的具体项，用于改写此单条规则项配置的内容，具体参考产品文档。	
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// RuleId 中指定托管规则项的处置动作。 SecurityAction 的 Name 取值支持：<li>Deny：拦截，响应拦截页面；</li><li>Monitor：观察，不处理请求记录安全事件到日志中；</li><li>Disabled：未启用，不扫描请求跳过该规则。</li>
	Action *SecurityAction `json:"Action,omitnil,omitempty" name:"Action"`
}

type ManagedRuleAutoUpdate struct {
	// 是否开启自动更新至最新版本。取值有：<li>on：开启</li><li>off：关闭</li>
	AutoUpdateToLatestVersion *string `json:"AutoUpdateToLatestVersion,omitnil,omitempty" name:"AutoUpdateToLatestVersion"`

	// 当前使用的版本，格式符合ISO 8601标准，如2023-12-21T12:00:32Z，默认为空，仅出参。
	RulesetVersion *string `json:"RulesetVersion,omitnil,omitempty" name:"RulesetVersion"`
}

type ManagedRuleDetail struct {
	// 托管规则Id。
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// 托管规则的防护级别。取值有：<li>low：低风险，此规则风险较低，适用于非常严格控制环境下的访问场景，该等级规则可能造成较多的误报；</li><li>medium：中风险，表示此条规则风险正常，适用较为严格的防护场景；</li><li>high：高风险，表示此条规则风险较高，大多数场景不会产生误报；</li><li>extreme：超高风险，表示此条规则风险极高，基本不会产生误报；</li>
	RiskLevel *string `json:"RiskLevel,omitnil,omitempty" name:"RiskLevel"`

	// 规则描述。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 规则标签。部分类型的规则不存在标签。
	Tags []*string `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 规则所属版本。
	RuleVersion *string `json:"RuleVersion,omitnil,omitempty" name:"RuleVersion"`
}

type ManagedRuleGroup struct {
	// 托管规则的组名称，未指定配置的规则分组将按照默认配置处理，GroupId 的具体取值参考产品文档。
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 托管规则组的防护级别。取值有：<li>loose：宽松，只包含超高风险规则，此时需配置Action，且RuleActions配置无效；</li><li>normal：正常，包含超高风险和高风险规则，此时需配置Action，且RuleActions配置无效；</li><li>strict：严格，包含超高风险、高风险和中风险规则，此时需配置Action，且RuleActions配置无效；</li><li>extreme：超严格，包含超高风险、高风险、中风险和低风险规则，此时需配置Action，且RuleActions配置无效；</li><li>custom：自定义，精细化策略，按单条规则配置处置方式，此时Action字段无效，使用RuleActions配置单条规则的精细化策略。</li>	
	SensitivityLevel *string `json:"SensitivityLevel,omitnil,omitempty" name:"SensitivityLevel"`

	// 托管规则组的处置动作。SecurityAction 的 Name 取值支持：<li>Deny：拦截，响应拦截页面；</li><li>Monitor：观察，不处理请求记录安全事件到日志中；</li><li>Disabled：未启用，不扫描请求跳过该规则。</li>
	Action *SecurityAction `json:"Action,omitnil,omitempty" name:"Action"`

	// 托管规则组下规则项的具体配置，仅在 SensitivityLevel 为 custom 时配置生效。
	RuleActions []*ManagedRuleAction `json:"RuleActions,omitnil,omitempty" name:"RuleActions"`

	// 托管规则组信息，仅出参。	
	MetaData *ManagedRuleGroupMeta `json:"MetaData,omitnil,omitempty" name:"MetaData"`
}

type ManagedRuleGroupMeta struct {
	// 托管规则组描述，仅出参。
	GroupDetail *string `json:"GroupDetail,omitnil,omitempty" name:"GroupDetail"`

	// 托管规则组名称，仅出参。
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// 当前托管规则组下的所有子规则信息，仅出参。
	RuleDetails []*ManagedRuleDetail `json:"RuleDetails,omitnil,omitempty" name:"RuleDetails"`
}

type ManagedRules struct {
	// 托管规则是否开启。取值有：<li>on：开启，所有托管规则按配置生效；</li><li>off：关闭，所有托管规则不生效。</li>
	Enabled *string `json:"Enabled,omitnil,omitempty" name:"Enabled"`

	// 评估模式是否开启，仅在 Enabled 参数为 on 时有效。取值有：<li>on：开启，表示所有托管规则以观察模式生效；</li><li>off：关闭，表示所有托管规则以实际配置生效。</li>
	DetectionOnly *string `json:"DetectionOnly,omitnil,omitempty" name:"DetectionOnly"`

	// 托管规则语义分析选项是否开启，仅在 Enabled 参数为 on 时有效。取值有：<li>on：开启，对请求进行语义分析后进行处理；</li><li>off：关闭，对请求不进行语义分析，直接进行处理。</li> <br/>默认为 off。
	SemanticAnalysis *string `json:"SemanticAnalysis,omitnil,omitempty" name:"SemanticAnalysis"`

	// 托管规则自动更新选项。
	AutoUpdate *ManagedRuleAutoUpdate `json:"AutoUpdate,omitnil,omitempty" name:"AutoUpdate"`

	// 托管规则组的配置。如果此结构传空数组或 GroupId 未包含在列表内将按照默认方式处理。
	ManagedRuleGroups []*ManagedRuleGroup `json:"ManagedRuleGroups,omitnil,omitempty" name:"ManagedRuleGroups"`

	// 高频扫描防护配置选项，当某一访客的请求频繁命中「配置为拦截」的托管规则时，在一段时间内封禁该访客所有请求。
	FrequentScanningProtection *FrequentScanningProtection `json:"FrequentScanningProtection,omitnil,omitempty" name:"FrequentScanningProtection"`
}

type MaxAge struct {
	// 是否遵循源站，取值有：
	// <li>on：遵循源站，忽略MaxAge 时间设置；</li>
	// <li>off：不遵循源站，使用MaxAge 时间设置。</li>
	FollowOrigin *string `json:"FollowOrigin,omitnil,omitempty" name:"FollowOrigin"`

	// MaxAge 时间设置，单位为秒，取值：0～315360000。
	// 注意：时间为0，即不缓存。
	MaxAgeTime *int64 `json:"MaxAgeTime,omitnil,omitempty" name:"MaxAgeTime"`
}

type MaxAgeParameters struct {
	// 遵循源站 Cache-Control 开关，取值有：
	// <li>on：遵循源站，忽略 CacheTime 时间设置；</li>
	// <li>off：不遵循源站，使用 CacheTime 时间设置。</li>
	FollowOrigin *string `json:"FollowOrigin,omitnil,omitempty" name:"FollowOrigin"`

	// 自定义缓存时间数值，单位为秒，取值：0～315360000。<br>注意：当 FollowOrigin 为 off 时，表示不遵循源站，使用 CacheTime 设置缓存时间，否则此字段不生效。
	CacheTime *int64 `json:"CacheTime,omitnil,omitempty" name:"CacheTime"`
}

type MaxNewSessionTriggerConfig struct {
	// 触发阈值统计的时间窗口，取值有：<li>5s：5 秒内；</li><li>10s：10 秒内；</li><li>15s：15 秒内；</li><li>30s：30 秒内；</li><li>60s：60 秒内；</li><li>5m：5 分钟内；</li><li>10m：10 分钟内；</li><li>30m：30 分钟内；</li><li>60m：60 分钟内。</li> 
	MaxNewSessionCountInterval *string `json:"MaxNewSessionCountInterval,omitnil,omitempty" name:"MaxNewSessionCountInterval"`

	// 触发阈值统计的累计次数，取值范围 1 ~ 100000000。
	MaxNewSessionCountThreshold *int64 `json:"MaxNewSessionCountThreshold,omitnil,omitempty" name:"MaxNewSessionCountThreshold"`
}

type MinimalRequestBodyTransferRate struct {
	// 正文传输最小速率阈值，单位仅支持bps。
	MinimalAvgTransferRateThreshold *string `json:"MinimalAvgTransferRateThreshold,omitnil,omitempty" name:"MinimalAvgTransferRateThreshold"`

	// 正文传输最小速率统计时间范围，取值有：<li>10s：10秒；</li><li>30s：30秒；</li><li>60s：60秒；</li><li>120s：120秒。</li> 
	CountingPeriod *string `json:"CountingPeriod,omitnil,omitempty" name:"CountingPeriod"`

	// 正文传输最小速率阈值是否开启。取值有：<li>on：开启；</li><li>off：关闭。</li>
	Enabled *string `json:"Enabled,omitnil,omitempty" name:"Enabled"`
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
type ModifyContentIdentifierRequestParams struct {
	// 内容标识符 ID。
	ContentId *string `json:"ContentId,omitnil,omitempty" name:"ContentId"`

	// 内容标识符描述，长度限制不超过 20 个字符。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type ModifyContentIdentifierRequest struct {
	*tchttp.BaseRequest
	
	// 内容标识符 ID。
	ContentId *string `json:"ContentId,omitnil,omitempty" name:"ContentId"`

	// 内容标识符描述，长度限制不超过 20 个字符。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

func (r *ModifyContentIdentifierRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyContentIdentifierRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ContentId")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyContentIdentifierRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyContentIdentifierResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyContentIdentifierResponse struct {
	*tchttp.BaseResponse
	Response *ModifyContentIdentifierResponseParams `json:"Response"`
}

func (r *ModifyContentIdentifierResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyContentIdentifierResponse) FromJsonString(s string) error {
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
type ModifyDDoSProtectionRequestParams struct {
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 独立 DDoS 防护配置。
	DDoSProtection *DDoSProtection `json:"DDoSProtection,omitnil,omitempty" name:"DDoSProtection"`
}

type ModifyDDoSProtectionRequest struct {
	*tchttp.BaseRequest
	
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 独立 DDoS 防护配置。
	DDoSProtection *DDoSProtection `json:"DDoSProtection,omitnil,omitempty" name:"DDoSProtection"`
}

func (r *ModifyDDoSProtectionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDDoSProtectionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "DDoSProtection")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDDoSProtectionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDDoSProtectionResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyDDoSProtectionResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDDoSProtectionResponseParams `json:"Response"`
}

func (r *ModifyDDoSProtectionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDDoSProtectionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDnsRecordsRequestParams struct {
	// 站点 ID 。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// DNS 记录修改数据列表，一次最多修改100条。
	DnsRecords []*DnsRecord `json:"DnsRecords,omitnil,omitempty" name:"DnsRecords"`
}

type ModifyDnsRecordsRequest struct {
	*tchttp.BaseRequest
	
	// 站点 ID 。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// DNS 记录修改数据列表，一次最多修改100条。
	DnsRecords []*DnsRecord `json:"DnsRecords,omitnil,omitempty" name:"DnsRecords"`
}

func (r *ModifyDnsRecordsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDnsRecordsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "DnsRecords")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDnsRecordsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDnsRecordsResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyDnsRecordsResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDnsRecordsResponseParams `json:"Response"`
}

func (r *ModifyDnsRecordsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDnsRecordsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDnsRecordsStatusRequestParams struct {
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 待启用的 DNS 记录 ID 列表，上限：200。<br>注意：同个 DNS 记录 ID 不能同时存在于 RecordsToEnable 和 RecordsToDisable。
	RecordsToEnable []*string `json:"RecordsToEnable,omitnil,omitempty" name:"RecordsToEnable"`

	// 待停用的 DNS 记录 ID 列表，上限：200。<br>注意：同个 DNS 记录 ID 不能同时存在于 RecordsToEnable 和 RecordsToDisable。
	RecordsToDisable []*string `json:"RecordsToDisable,omitnil,omitempty" name:"RecordsToDisable"`
}

type ModifyDnsRecordsStatusRequest struct {
	*tchttp.BaseRequest
	
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 待启用的 DNS 记录 ID 列表，上限：200。<br>注意：同个 DNS 记录 ID 不能同时存在于 RecordsToEnable 和 RecordsToDisable。
	RecordsToEnable []*string `json:"RecordsToEnable,omitnil,omitempty" name:"RecordsToEnable"`

	// 待停用的 DNS 记录 ID 列表，上限：200。<br>注意：同个 DNS 记录 ID 不能同时存在于 RecordsToEnable 和 RecordsToDisable。
	RecordsToDisable []*string `json:"RecordsToDisable,omitnil,omitempty" name:"RecordsToDisable"`
}

func (r *ModifyDnsRecordsStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDnsRecordsStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "RecordsToEnable")
	delete(f, "RecordsToDisable")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDnsRecordsStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDnsRecordsStatusResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyDnsRecordsStatusResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDnsRecordsStatusResponseParams `json:"Response"`
}

func (r *ModifyDnsRecordsStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDnsRecordsStatusResponse) FromJsonString(s string) error {
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

	// 规则 ID。您可以先通过 DescribeFunctionRules 接口来获取需要修改的规则的 RuleId，然后传入修改后的规则内容，原规则内容会被覆盖式更新。
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// 规则条件列表，相同触发规则的不同条件匹配项之间为或关系，不填写保持原有配置。
	FunctionRuleConditions []*FunctionRuleCondition `json:"FunctionRuleConditions,omitnil,omitempty" name:"FunctionRuleConditions"`

	// 函数选择配置类型：
	// <li> direct：直接指定执行函数；</li>
	// <li> weight：基于权重比选择函数；</li>
	// <li> region：基于客户端 IP 的国家/地区选择函数。</li>
	// 不填时默认为 direct 。
	TriggerType *string `json:"TriggerType,omitnil,omitempty" name:"TriggerType"`

	// 指定执行的函数 ID。当 TriggerType 为 direct 或 TriggerType 不填时生效。
	FunctionId *string `json:"FunctionId,omitnil,omitempty" name:"FunctionId"`

	// 基于客户端 IP 国家/地区的函数选择配置，当 TriggerType 为 region 时生效且 RegionMappingSelections 必填。RegionMappingSelections 中至少包含一项 Regions 为 Default 的配置。
	RegionMappingSelections []*FunctionRegionSelection `json:"RegionMappingSelections,omitnil,omitempty" name:"RegionMappingSelections"`

	// 基于权重的函数选择配置，当 TriggerType 为 weight 时生效且 WeightedSelections 必填。WeightedSelections 中的所有权重之和需要为100。
	WeightedSelections []*FunctionWeightedSelection `json:"WeightedSelections,omitnil,omitempty" name:"WeightedSelections"`

	// 规则描述，最大支持 60 个字符，不填写保持原有配置。
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`
}

type ModifyFunctionRuleRequest struct {
	*tchttp.BaseRequest
	
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 规则 ID。您可以先通过 DescribeFunctionRules 接口来获取需要修改的规则的 RuleId，然后传入修改后的规则内容，原规则内容会被覆盖式更新。
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// 规则条件列表，相同触发规则的不同条件匹配项之间为或关系，不填写保持原有配置。
	FunctionRuleConditions []*FunctionRuleCondition `json:"FunctionRuleConditions,omitnil,omitempty" name:"FunctionRuleConditions"`

	// 函数选择配置类型：
	// <li> direct：直接指定执行函数；</li>
	// <li> weight：基于权重比选择函数；</li>
	// <li> region：基于客户端 IP 的国家/地区选择函数。</li>
	// 不填时默认为 direct 。
	TriggerType *string `json:"TriggerType,omitnil,omitempty" name:"TriggerType"`

	// 指定执行的函数 ID。当 TriggerType 为 direct 或 TriggerType 不填时生效。
	FunctionId *string `json:"FunctionId,omitnil,omitempty" name:"FunctionId"`

	// 基于客户端 IP 国家/地区的函数选择配置，当 TriggerType 为 region 时生效且 RegionMappingSelections 必填。RegionMappingSelections 中至少包含一项 Regions 为 Default 的配置。
	RegionMappingSelections []*FunctionRegionSelection `json:"RegionMappingSelections,omitnil,omitempty" name:"RegionMappingSelections"`

	// 基于权重的函数选择配置，当 TriggerType 为 weight 时生效且 WeightedSelections 必填。WeightedSelections 中的所有权重之和需要为100。
	WeightedSelections []*FunctionWeightedSelection `json:"WeightedSelections,omitnil,omitempty" name:"WeightedSelections"`

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
	delete(f, "TriggerType")
	delete(f, "FunctionId")
	delete(f, "RegionMappingSelections")
	delete(f, "WeightedSelections")
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
	// <ul><li>disable：不配置服务端证书；</li>
	// <li>eofreecert：通过自动验证申请免费证书并部署。验证方式详见：[申请免费证书支持的验证方式](https://cloud.tencent.com/document/product/1552/90437)
	// 
	// - 在 NS 或者 DNSPod 托管接入模式下，仅支持自动验证的方式申请免费证书。
	// - 当免费证书申请失败时会导致证书部署失败，您可以通过<a href = 'https://cloud.tencent.com/document/product/1552/124806'>检查免费证书申请结果</a>接口获取申请失败原因。</li>
	// </ul><li>eofreecert_manual：部署 DNS 委派验证或者文件验证申请的免费证书。在部署免费证书前，您需要触发<a href = 'https://cloud.tencent.com/document/product/1552/124807'>申请免费证书</a>接口申请免费证书。在免费证书申请成功后，你可以通过该枚举值对免费证书进行部署；</li>
	// <ul><li>注意：在对免费证书部署时，需要保证当前已存在申请成功的免费证书。您可以通过<a href = 'https://cloud.tencent.com/document/product/1552/124806'>检查免费证书申请结果</a>接口检查当前是否已存在申请成功的免费证书。</li>
	// </ul><li>sslcert：配置 SSL 托管服务端证书。</li>
	Mode *string `json:"Mode,omitnil,omitempty" name:"Mode"`

	// SSL 证书配置，本参数仅在 mode 为 sslcert 时生效，传入对应证书的 CertId 即可。您可以前往 [SSL 证书列表](https://console.cloud.tencent.com/ssl) 查看 CertId。
	ServerCertInfo []*ServerCertInfo `json:"ServerCertInfo,omitnil,omitempty" name:"ServerCertInfo"`

	// 托管类型，取值有：
	// <li>none：不托管EO；</li>
	// <li>apply：托管EO</li>
	// 不填，默认取值为none。
	//
	// Deprecated: ApplyType is deprecated.
	ApplyType *string `json:"ApplyType,omitnil,omitempty" name:"ApplyType"`

	// 在边缘双向认证场景下，该字段为客户端的 CA 证书，部署在 EO 节点内，用于客户端对 EO 节点进行认证。默认关闭，不填写表示保持原有配置。
	ClientCertInfo *MutualTLS `json:"ClientCertInfo,omitnil,omitempty" name:"ClientCertInfo"`

	// 用于分别开启/关闭回源双向认证和源站证书校验。默认关闭，不填写表示保持原有配置。回源双向认证配置当前为白名单内侧中，如需使用，请[联系我们](https://cloud.tencent.com/online-service)。
	UpstreamCertInfo *UpstreamCertInfo `json:"UpstreamCertInfo,omitnil,omitempty" name:"UpstreamCertInfo"`
}

type ModifyHostsCertificateRequest struct {
	*tchttp.BaseRequest
	
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 需要修改证书配置的加速域名。
	Hosts []*string `json:"Hosts,omitnil,omitempty" name:"Hosts"`

	// 配置服务端证书的模式，取值有：
	// <ul><li>disable：不配置服务端证书；</li>
	// <li>eofreecert：通过自动验证申请免费证书并部署。验证方式详见：[申请免费证书支持的验证方式](https://cloud.tencent.com/document/product/1552/90437)
	// 
	// - 在 NS 或者 DNSPod 托管接入模式下，仅支持自动验证的方式申请免费证书。
	// - 当免费证书申请失败时会导致证书部署失败，您可以通过<a href = 'https://cloud.tencent.com/document/product/1552/124806'>检查免费证书申请结果</a>接口获取申请失败原因。</li>
	// </ul><li>eofreecert_manual：部署 DNS 委派验证或者文件验证申请的免费证书。在部署免费证书前，您需要触发<a href = 'https://cloud.tencent.com/document/product/1552/124807'>申请免费证书</a>接口申请免费证书。在免费证书申请成功后，你可以通过该枚举值对免费证书进行部署；</li>
	// <ul><li>注意：在对免费证书部署时，需要保证当前已存在申请成功的免费证书。您可以通过<a href = 'https://cloud.tencent.com/document/product/1552/124806'>检查免费证书申请结果</a>接口检查当前是否已存在申请成功的免费证书。</li>
	// </ul><li>sslcert：配置 SSL 托管服务端证书。</li>
	Mode *string `json:"Mode,omitnil,omitempty" name:"Mode"`

	// SSL 证书配置，本参数仅在 mode 为 sslcert 时生效，传入对应证书的 CertId 即可。您可以前往 [SSL 证书列表](https://console.cloud.tencent.com/ssl) 查看 CertId。
	ServerCertInfo []*ServerCertInfo `json:"ServerCertInfo,omitnil,omitempty" name:"ServerCertInfo"`

	// 托管类型，取值有：
	// <li>none：不托管EO；</li>
	// <li>apply：托管EO</li>
	// 不填，默认取值为none。
	ApplyType *string `json:"ApplyType,omitnil,omitempty" name:"ApplyType"`

	// 在边缘双向认证场景下，该字段为客户端的 CA 证书，部署在 EO 节点内，用于客户端对 EO 节点进行认证。默认关闭，不填写表示保持原有配置。
	ClientCertInfo *MutualTLS `json:"ClientCertInfo,omitnil,omitempty" name:"ClientCertInfo"`

	// 用于分别开启/关闭回源双向认证和源站证书校验。默认关闭，不填写表示保持原有配置。回源双向认证配置当前为白名单内侧中，如需使用，请[联系我们](https://cloud.tencent.com/online-service)。
	UpstreamCertInfo *UpstreamCertInfo `json:"UpstreamCertInfo,omitnil,omitempty" name:"UpstreamCertInfo"`
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
	delete(f, "UpstreamCertInfo")
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
type ModifyL7AccRulePriorityRequestParams struct {
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 站点 ID 下完整的规则 ID 列表，规则 ID 列表可以通过 [查询七层加速规则](https://cloud.tencent.com/document/product/1552/115820) 获取，最终优先级顺序将调整成规则 ID 列表的顺序，从前往后依次执行。
	RuleIds []*string `json:"RuleIds,omitnil,omitempty" name:"RuleIds"`
}

type ModifyL7AccRulePriorityRequest struct {
	*tchttp.BaseRequest
	
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 站点 ID 下完整的规则 ID 列表，规则 ID 列表可以通过 [查询七层加速规则](https://cloud.tencent.com/document/product/1552/115820) 获取，最终优先级顺序将调整成规则 ID 列表的顺序，从前往后依次执行。
	RuleIds []*string `json:"RuleIds,omitnil,omitempty" name:"RuleIds"`
}

func (r *ModifyL7AccRulePriorityRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyL7AccRulePriorityRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "RuleIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyL7AccRulePriorityRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyL7AccRulePriorityResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyL7AccRulePriorityResponse struct {
	*tchttp.BaseResponse
	Response *ModifyL7AccRulePriorityResponseParams `json:"Response"`
}

func (r *ModifyL7AccRulePriorityResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyL7AccRulePriorityResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyL7AccRuleRequestParams struct {
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 需要修改的规则。您可以先通过 DescribeL7AccRules 接口来获取需要修改的规则的 Ruleid，然后传入修改后的规则内容，原规则内容会被覆盖式更新。
	Rule *RuleEngineItem `json:"Rule,omitnil,omitempty" name:"Rule"`
}

type ModifyL7AccRuleRequest struct {
	*tchttp.BaseRequest
	
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 需要修改的规则。您可以先通过 DescribeL7AccRules 接口来获取需要修改的规则的 Ruleid，然后传入修改后的规则内容，原规则内容会被覆盖式更新。
	Rule *RuleEngineItem `json:"Rule,omitnil,omitempty" name:"Rule"`
}

func (r *ModifyL7AccRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyL7AccRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "Rule")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyL7AccRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyL7AccRuleResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyL7AccRuleResponse struct {
	*tchttp.BaseResponse
	Response *ModifyL7AccRuleResponseParams `json:"Response"`
}

func (r *ModifyL7AccRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyL7AccRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyL7AccSettingRequestParams struct {
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 站点加速全局配置，该参数中的配置会对站点下的所有域名生效。您只需直接修改所需的配置，未传入的其他配置将保持原有状态。
	// 
	ZoneConfig *ZoneConfig `json:"ZoneConfig,omitnil,omitempty" name:"ZoneConfig"`
}

type ModifyL7AccSettingRequest struct {
	*tchttp.BaseRequest
	
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 站点加速全局配置，该参数中的配置会对站点下的所有域名生效。您只需直接修改所需的配置，未传入的其他配置将保持原有状态。
	// 
	ZoneConfig *ZoneConfig `json:"ZoneConfig,omitnil,omitempty" name:"ZoneConfig"`
}

func (r *ModifyL7AccSettingRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyL7AccSettingRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "ZoneConfig")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyL7AccSettingRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyL7AccSettingResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyL7AccSettingResponse struct {
	*tchttp.BaseResponse
	Response *ModifyL7AccSettingResponseParams `json:"Response"`
}

func (r *ModifyL7AccSettingResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyL7AccSettingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyLoadBalancerRequestParams struct {
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 负载均衡实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 实例名称，可输入 1-200 个字符，允许字符为 a-z，A-Z，0-9，_，-。不填写表示维持原有配置。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 源站组列表及其对应的容灾调度优先级。详情请参考 [快速创建负载均衡实例](https://cloud.tencent.com/document/product/1552/104223) 中的示例场景。不填写表示维持原有配置。
	OriginGroups []*OriginGroupInLoadBalancer `json:"OriginGroups,omitnil,omitempty" name:"OriginGroups"`

	// 健康检查策略。详情请参考 [健康检查策略介绍](https://cloud.tencent.com/document/product/1552/104228)。不填写表示维持原有配置。
	HealthChecker *HealthChecker `json:"HealthChecker,omitnil,omitempty" name:"HealthChecker"`

	// 源站组间的流量调度策略，取值有：
	// <li>Pritory：按优先级顺序进行故障转移 。</li>不填写表示维持原有配置。
	SteeringPolicy *string `json:"SteeringPolicy,omitnil,omitempty" name:"SteeringPolicy"`

	// 实际访问某源站失败时的请求重试策略，详情请参考 [请求重试策略介绍](https://cloud.tencent.com/document/product/1552/104227)，取值有：
	// <li>OtherOriginGroup：单次请求失败后，请求优先重试下一优先级源站组；</li>
	// <li>OtherRecordInOriginGroup：单次请求失败后，请求优先重试同源站组内的其他源站。</li>不填写表示维持原有配置。
	FailoverPolicy *string `json:"FailoverPolicy,omitnil,omitempty" name:"FailoverPolicy"`
}

type ModifyLoadBalancerRequest struct {
	*tchttp.BaseRequest
	
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 负载均衡实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 实例名称，可输入 1-200 个字符，允许字符为 a-z，A-Z，0-9，_，-。不填写表示维持原有配置。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 源站组列表及其对应的容灾调度优先级。详情请参考 [快速创建负载均衡实例](https://cloud.tencent.com/document/product/1552/104223) 中的示例场景。不填写表示维持原有配置。
	OriginGroups []*OriginGroupInLoadBalancer `json:"OriginGroups,omitnil,omitempty" name:"OriginGroups"`

	// 健康检查策略。详情请参考 [健康检查策略介绍](https://cloud.tencent.com/document/product/1552/104228)。不填写表示维持原有配置。
	HealthChecker *HealthChecker `json:"HealthChecker,omitnil,omitempty" name:"HealthChecker"`

	// 源站组间的流量调度策略，取值有：
	// <li>Pritory：按优先级顺序进行故障转移 。</li>不填写表示维持原有配置。
	SteeringPolicy *string `json:"SteeringPolicy,omitnil,omitempty" name:"SteeringPolicy"`

	// 实际访问某源站失败时的请求重试策略，详情请参考 [请求重试策略介绍](https://cloud.tencent.com/document/product/1552/104227)，取值有：
	// <li>OtherOriginGroup：单次请求失败后，请求优先重试下一优先级源站组；</li>
	// <li>OtherRecordInOriginGroup：单次请求失败后，请求优先重试同源站组内的其他源站。</li>不填写表示维持原有配置。
	FailoverPolicy *string `json:"FailoverPolicy,omitnil,omitempty" name:"FailoverPolicy"`
}

func (r *ModifyLoadBalancerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLoadBalancerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "InstanceId")
	delete(f, "Name")
	delete(f, "OriginGroups")
	delete(f, "HealthChecker")
	delete(f, "SteeringPolicy")
	delete(f, "FailoverPolicy")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyLoadBalancerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyLoadBalancerResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyLoadBalancerResponse struct {
	*tchttp.BaseResponse
	Response *ModifyLoadBalancerResponseParams `json:"Response"`
}

func (r *ModifyLoadBalancerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLoadBalancerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyMultiPathGatewayLineRequestParams struct {
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 多通道安全加速网关 ID 。
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// 线路 ID ， 取值有:
	// <li> line-1： EdgeOne 四层代理线路，支持修改实例和规则，不支持删除；</li>
	// <li> line-2 及以上：EdgeOne 四层代理线路或者自定义线路，支持修改、删除实例和规则。</li>
	LineId *string `json:"LineId,omitnil,omitempty" name:"LineId"`

	// 线路类型，取值有： 
	// <li>proxy ：EdgeOne 四层代理线路，支持修改实例和规则，不支持删除；</li> 
	// <li>custom ：自定义线路，支持编辑、删除实例和规则。</li>
	LineType *string `json:"LineType,omitnil,omitempty" name:"LineType"`

	// 线路地址，格式为 host:port，直连线路（ LineType 取值为 direct ）不允许修改，其余类型支持修改。
	LineAddress *string `json:"LineAddress,omitnil,omitempty" name:"LineAddress"`

	// 四层代理实例 ID  ，当线路类型 LineType  取值为 proxy（EdgeOne 四层代理）可传入，进行修改。
	ProxyId *string `json:"ProxyId,omitnil,omitempty" name:"ProxyId"`

	// 转发规则 ID ，当线路类型 LineType 取值为 proxy（EdgeOne 四层代理）可传入，进行修改。
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`
}

type ModifyMultiPathGatewayLineRequest struct {
	*tchttp.BaseRequest
	
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 多通道安全加速网关 ID 。
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// 线路 ID ， 取值有:
	// <li> line-1： EdgeOne 四层代理线路，支持修改实例和规则，不支持删除；</li>
	// <li> line-2 及以上：EdgeOne 四层代理线路或者自定义线路，支持修改、删除实例和规则。</li>
	LineId *string `json:"LineId,omitnil,omitempty" name:"LineId"`

	// 线路类型，取值有： 
	// <li>proxy ：EdgeOne 四层代理线路，支持修改实例和规则，不支持删除；</li> 
	// <li>custom ：自定义线路，支持编辑、删除实例和规则。</li>
	LineType *string `json:"LineType,omitnil,omitempty" name:"LineType"`

	// 线路地址，格式为 host:port，直连线路（ LineType 取值为 direct ）不允许修改，其余类型支持修改。
	LineAddress *string `json:"LineAddress,omitnil,omitempty" name:"LineAddress"`

	// 四层代理实例 ID  ，当线路类型 LineType  取值为 proxy（EdgeOne 四层代理）可传入，进行修改。
	ProxyId *string `json:"ProxyId,omitnil,omitempty" name:"ProxyId"`

	// 转发规则 ID ，当线路类型 LineType 取值为 proxy（EdgeOne 四层代理）可传入，进行修改。
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`
}

func (r *ModifyMultiPathGatewayLineRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyMultiPathGatewayLineRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "GatewayId")
	delete(f, "LineId")
	delete(f, "LineType")
	delete(f, "LineAddress")
	delete(f, "ProxyId")
	delete(f, "RuleId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyMultiPathGatewayLineRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyMultiPathGatewayLineResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyMultiPathGatewayLineResponse struct {
	*tchttp.BaseResponse
	Response *ModifyMultiPathGatewayLineResponseParams `json:"Response"`
}

func (r *ModifyMultiPathGatewayLineResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyMultiPathGatewayLineResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyMultiPathGatewayRequestParams struct {
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 网关 ID。
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// 网关名称，16 个字符以内，可用字符（a-z,A-Z,0-9,-,_）。
	GatewayName *string `json:"GatewayName,omitnil,omitempty" name:"GatewayName"`

	// 网关地址，GatewayType 取值为 private（自有网关）可填入进行修改，使用该地址时，请确保该地址已录入腾讯云多通道安全加速网关系统。如未录入，需要在本接口调用前通过工单或者联系架构师把网关 IP 地址提前录入腾讯云多通道安全加速网关系统。
	GatewayIP *string `json:"GatewayIP,omitnil,omitempty" name:"GatewayIP"`

	// 网关端口，范围 1～65535（除去 8888 ），只支持修改 GatewayType 取值为 private 的自有网关。
	GatewayPort *int64 `json:"GatewayPort,omitnil,omitempty" name:"GatewayPort"`
}

type ModifyMultiPathGatewayRequest struct {
	*tchttp.BaseRequest
	
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 网关 ID。
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// 网关名称，16 个字符以内，可用字符（a-z,A-Z,0-9,-,_）。
	GatewayName *string `json:"GatewayName,omitnil,omitempty" name:"GatewayName"`

	// 网关地址，GatewayType 取值为 private（自有网关）可填入进行修改，使用该地址时，请确保该地址已录入腾讯云多通道安全加速网关系统。如未录入，需要在本接口调用前通过工单或者联系架构师把网关 IP 地址提前录入腾讯云多通道安全加速网关系统。
	GatewayIP *string `json:"GatewayIP,omitnil,omitempty" name:"GatewayIP"`

	// 网关端口，范围 1～65535（除去 8888 ），只支持修改 GatewayType 取值为 private 的自有网关。
	GatewayPort *int64 `json:"GatewayPort,omitnil,omitempty" name:"GatewayPort"`
}

func (r *ModifyMultiPathGatewayRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyMultiPathGatewayRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "GatewayId")
	delete(f, "GatewayName")
	delete(f, "GatewayIP")
	delete(f, "GatewayPort")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyMultiPathGatewayRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyMultiPathGatewayResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyMultiPathGatewayResponse struct {
	*tchttp.BaseResponse
	Response *ModifyMultiPathGatewayResponseParams `json:"Response"`
}

func (r *ModifyMultiPathGatewayResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyMultiPathGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyMultiPathGatewaySecretKeyRequestParams struct {
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 多通道安全加速网关接入密钥，base64 字符串，编码前字符串长度为 32-48 个字符。
	SecretKey *string `json:"SecretKey,omitnil,omitempty" name:"SecretKey"`
}

type ModifyMultiPathGatewaySecretKeyRequest struct {
	*tchttp.BaseRequest
	
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 多通道安全加速网关接入密钥，base64 字符串，编码前字符串长度为 32-48 个字符。
	SecretKey *string `json:"SecretKey,omitnil,omitempty" name:"SecretKey"`
}

func (r *ModifyMultiPathGatewaySecretKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyMultiPathGatewaySecretKeyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "SecretKey")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyMultiPathGatewaySecretKeyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyMultiPathGatewaySecretKeyResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyMultiPathGatewaySecretKeyResponse struct {
	*tchttp.BaseResponse
	Response *ModifyMultiPathGatewaySecretKeyResponseParams `json:"Response"`
}

func (r *ModifyMultiPathGatewaySecretKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyMultiPathGatewaySecretKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyMultiPathGatewayStatusRequestParams struct {
	// 网关 ID。
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 修改网关的启用停用状态，取值有：<li> offline：停用；</li><li> online：启用。</li>
	GatewayStatus *string `json:"GatewayStatus,omitnil,omitempty" name:"GatewayStatus"`
}

type ModifyMultiPathGatewayStatusRequest struct {
	*tchttp.BaseRequest
	
	// 网关 ID。
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 修改网关的启用停用状态，取值有：<li> offline：停用；</li><li> online：启用。</li>
	GatewayStatus *string `json:"GatewayStatus,omitnil,omitempty" name:"GatewayStatus"`
}

func (r *ModifyMultiPathGatewayStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyMultiPathGatewayStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GatewayId")
	delete(f, "ZoneId")
	delete(f, "GatewayStatus")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyMultiPathGatewayStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyMultiPathGatewayStatusResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyMultiPathGatewayStatusResponse struct {
	*tchttp.BaseResponse
	Response *ModifyMultiPathGatewayStatusResponseParams `json:"Response"`
}

func (r *ModifyMultiPathGatewayStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyMultiPathGatewayStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyOriginACLRequestParams struct {
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 需要启用/关闭特定回源 IP 网段回源的实例。
	OriginACLEntities []*OriginACLEntity `json:"OriginACLEntities,omitnil,omitempty" name:"OriginACLEntities"`
}

type ModifyOriginACLRequest struct {
	*tchttp.BaseRequest
	
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 需要启用/关闭特定回源 IP 网段回源的实例。
	OriginACLEntities []*OriginACLEntity `json:"OriginACLEntities,omitnil,omitempty" name:"OriginACLEntities"`
}

func (r *ModifyOriginACLRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyOriginACLRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "OriginACLEntities")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyOriginACLRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyOriginACLResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyOriginACLResponse struct {
	*tchttp.BaseResponse
	Response *ModifyOriginACLResponseParams `json:"Response"`
}

func (r *ModifyOriginACLResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyOriginACLResponse) FromJsonString(s string) error {
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

type ModifyOriginParameters struct {
	// 源站类型。取值有：
	// <li>IPDomain：IPV4、IPV6 或域名类型源站；</li>
	// <li>OriginGroup：源站组类型源站；</li>
	// <li>LoadBalance：负载均衡，该功能内测中，如需使用，请提工单或联系智能客服；</li>
	// <li>COS：腾讯云 COS 对象存储源站；</li>
	// <li>AWSS3：支持 AWS S3 协议的所有对象存储源站。</li>
	OriginType *string `json:"OriginType,omitnil,omitempty" name:"OriginType"`

	// 源站地址，根据 OriginType 的取值分为以下情况：
	// <li>当 OriginType = IPDomain 时，该参数请填写 IPV4、IPV6 地址或域名；</li>
	// <li>当 OriginType = COS 时，该参数请填写 COS 桶的访问域名；</li>
	// <li>当 OriginType = AWSS3，该参数请填写 S3 桶的访问域名；</li>
	// <li>当 OriginType = OriginGroup 时，该参数请填写源站组 ID；</li>
	// <li>当 OriginType = LoadBalance 时，该参数请填写负载均衡实例 ID，该功能当前仅白名单开放。</li>
	Origin *string `json:"Origin,omitnil,omitempty" name:"Origin"`

	// 回源协议配置。当 OriginType 取值为 IPDomain、OriginGroup、LoadBalance 时该参数必填。取值有：
	// <li>http：使用 HTTP 协议；</li>
	// <li>https：使用 HTTPS 协议；</li>
	// <li>follow：协议跟随。</li>
	OriginProtocol *string `json:"OriginProtocol,omitnil,omitempty" name:"OriginProtocol"`

	// HTTP 回源端口，取值范围 1～65535。当回源协议 OriginProtocol 为 http 或者 follow 时该参数必填。
	HTTPOriginPort *int64 `json:"HTTPOriginPort,omitnil,omitempty" name:"HTTPOriginPort"`

	// HTTPS 回源端口，取值范围 1～65535。当回源协议 OriginProtocol 为 https 或者 follow 时该参数必填。
	HTTPSOriginPort *int64 `json:"HTTPSOriginPort,omitnil,omitempty" name:"HTTPSOriginPort"`

	// 指定是否允许访问私有对象存储源站，当源站类型 OriginType = COS 或 AWSS3 时该参数必填，取值有：
	// <li>on：使用私有鉴权；</li>
	// <li>off：不使用私有鉴权。</li>
	PrivateAccess *string `json:"PrivateAccess,omitnil,omitempty" name:"PrivateAccess"`

	// 私有鉴权使用参数，该参数仅当 OriginType = AWSS3 且 PrivateAccess = on 时会生效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PrivateParameters *OriginPrivateParameters `json:"PrivateParameters,omitnil,omitempty" name:"PrivateParameters"`
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
type ModifyPrefetchOriginLimitRequestParams struct {
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 加速域名。
	DomainName *string `json:"DomainName,omitnil,omitempty" name:"DomainName"`

	// 回源限速限制的加速区域。
	// 预热时，该加速区域将会受到配置的Bandwidth值限制。取值有：
	// <li>Overseas：全球可用区（不含中国大陆）；</li>
	// <li>MainlandChina：中国大陆可用区。</li>
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 回源限速带宽。
	// 预热时回到源站的带宽上限值，取值范围 100 - 100,000，单位 Mbps。
	Bandwidth *int64 `json:"Bandwidth,omitnil,omitempty" name:"Bandwidth"`

	// 回源限速限制控制开关。
	// 用于启用/删除本条回源限速限制，取值有：
	// <li>on：启用限制；</li>
	// <li>off：删除限制。</li>
	Enabled *string `json:"Enabled,omitnil,omitempty" name:"Enabled"`
}

type ModifyPrefetchOriginLimitRequest struct {
	*tchttp.BaseRequest
	
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 加速域名。
	DomainName *string `json:"DomainName,omitnil,omitempty" name:"DomainName"`

	// 回源限速限制的加速区域。
	// 预热时，该加速区域将会受到配置的Bandwidth值限制。取值有：
	// <li>Overseas：全球可用区（不含中国大陆）；</li>
	// <li>MainlandChina：中国大陆可用区。</li>
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 回源限速带宽。
	// 预热时回到源站的带宽上限值，取值范围 100 - 100,000，单位 Mbps。
	Bandwidth *int64 `json:"Bandwidth,omitnil,omitempty" name:"Bandwidth"`

	// 回源限速限制控制开关。
	// 用于启用/删除本条回源限速限制，取值有：
	// <li>on：启用限制；</li>
	// <li>off：删除限制。</li>
	Enabled *string `json:"Enabled,omitnil,omitempty" name:"Enabled"`
}

func (r *ModifyPrefetchOriginLimitRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPrefetchOriginLimitRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "DomainName")
	delete(f, "Area")
	delete(f, "Bandwidth")
	delete(f, "Enabled")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyPrefetchOriginLimitRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyPrefetchOriginLimitResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyPrefetchOriginLimitResponse struct {
	*tchttp.BaseResponse
	Response *ModifyPrefetchOriginLimitResponseParams `json:"Response"`
}

func (r *ModifyPrefetchOriginLimitResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPrefetchOriginLimitResponse) FromJsonString(s string) error {
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

	// 投递的自定义字段列表，支持在 HTTP 请求头、响应头、Cookie、请求正文中提取指定内容。不填保持原有配置。自定义字段名称不能重复，且最多不能超过 200 个字段。单个实时日志推送任务最多添加 5 个请求正文类型的自定义字段。目前仅站点加速日志（LogType=domain）支持添加自定义字段。
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

	// 投递的自定义字段列表，支持在 HTTP 请求头、响应头、Cookie、请求正文中提取指定内容。不填保持原有配置。自定义字段名称不能重复，且最多不能超过 200 个字段。单个实时日志推送任务最多添加 5 个请求正文类型的自定义字段。目前仅站点加速日志（LogType=domain）支持添加自定义字段。
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

type ModifyRequestHeaderParameters struct {
	// HTTP 头部设置规则列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	HeaderActions []*HeaderAction `json:"HeaderActions,omitnil,omitempty" name:"HeaderActions"`
}

type ModifyResponseHeaderParameters struct {
	// HTTP 回源头部规则列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	HeaderActions []*HeaderAction `json:"HeaderActions,omitnil,omitempty" name:"HeaderActions"`
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
type ModifySecurityAPIResourceRequestParams struct {
	// 站点 ID。	
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// API 资源列表。
	APIResources []*APIResource `json:"APIResources,omitnil,omitempty" name:"APIResources"`
}

type ModifySecurityAPIResourceRequest struct {
	*tchttp.BaseRequest
	
	// 站点 ID。	
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// API 资源列表。
	APIResources []*APIResource `json:"APIResources,omitnil,omitempty" name:"APIResources"`
}

func (r *ModifySecurityAPIResourceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySecurityAPIResourceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "APIResources")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifySecurityAPIResourceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySecurityAPIResourceResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifySecurityAPIResourceResponse struct {
	*tchttp.BaseResponse
	Response *ModifySecurityAPIResourceResponseParams `json:"Response"`
}

func (r *ModifySecurityAPIResourceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySecurityAPIResourceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySecurityAPIServiceRequestParams struct {
	// 站点 ID。	
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// API 服务列表。
	APIServices []*APIService `json:"APIServices,omitnil,omitempty" name:"APIServices"`
}

type ModifySecurityAPIServiceRequest struct {
	*tchttp.BaseRequest
	
	// 站点 ID。	
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// API 服务列表。
	APIServices []*APIService `json:"APIServices,omitnil,omitempty" name:"APIServices"`
}

func (r *ModifySecurityAPIServiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySecurityAPIServiceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "APIServices")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifySecurityAPIServiceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySecurityAPIServiceResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifySecurityAPIServiceResponse struct {
	*tchttp.BaseResponse
	Response *ModifySecurityAPIServiceResponseParams `json:"Response"`
}

func (r *ModifySecurityAPIServiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySecurityAPIServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySecurityClientAttesterRequestParams struct {
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 认证选项列表。
	ClientAttesters []*ClientAttester `json:"ClientAttesters,omitnil,omitempty" name:"ClientAttesters"`
}

type ModifySecurityClientAttesterRequest struct {
	*tchttp.BaseRequest
	
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 认证选项列表。
	ClientAttesters []*ClientAttester `json:"ClientAttesters,omitnil,omitempty" name:"ClientAttesters"`
}

func (r *ModifySecurityClientAttesterRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySecurityClientAttesterRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "ClientAttesters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifySecurityClientAttesterRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySecurityClientAttesterResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifySecurityClientAttesterResponse struct {
	*tchttp.BaseResponse
	Response *ModifySecurityClientAttesterResponseParams `json:"Response"`
}

func (r *ModifySecurityClientAttesterResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySecurityClientAttesterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySecurityIPGroupRequestParams struct {
	// 站点 Id。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// IP 组配置。
	IPGroup *IPGroup `json:"IPGroup,omitnil,omitempty" name:"IPGroup"`

	// 操作类型，取值有：<li> append: 向 IPGroup 中添加新的 IP 地址或设置定时过期时间；</li><li>  remove: 从 IPGroup 中删除指定的 IP 地址或其定时过期时间；</li><li>  update: 完全替换 IPGroup 中 Content 或 ExpireInfo 的内容，并且可以修改 IPGroup 的名称。</li>    使用 append 操作时注意：   <li> 为 IP 或网段添加定时过期时间时，必须晚于当前时间。如果该 IP 或网段在组中不存在，必须同时在 Content 参数中添加该 IP 或网段。若该 IP 或网段已存在过期时间，则新时间将覆盖原有时间。</li>  使用 remove 操作时注意： <li> 删除 IP 或网段时，相关的未过期的定时过期时间也会被删除；</li> <li> 删除定时过期时间时，仅能删除当前未过期的时间。</li>  使用 update 操作时注意： <li> 替换 Content 内容时，不在 Content 中的 IP 或网段的未过期时间会被删除；</li> <li> 替换 IPExpireInfo 内容时，IPExpireInfo 中的 IP 或网段必须在 Content 中或在 IP 组中存在。</li>
	Mode *string `json:"Mode,omitnil,omitempty" name:"Mode"`
}

type ModifySecurityIPGroupRequest struct {
	*tchttp.BaseRequest
	
	// 站点 Id。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// IP 组配置。
	IPGroup *IPGroup `json:"IPGroup,omitnil,omitempty" name:"IPGroup"`

	// 操作类型，取值有：<li> append: 向 IPGroup 中添加新的 IP 地址或设置定时过期时间；</li><li>  remove: 从 IPGroup 中删除指定的 IP 地址或其定时过期时间；</li><li>  update: 完全替换 IPGroup 中 Content 或 ExpireInfo 的内容，并且可以修改 IPGroup 的名称。</li>    使用 append 操作时注意：   <li> 为 IP 或网段添加定时过期时间时，必须晚于当前时间。如果该 IP 或网段在组中不存在，必须同时在 Content 参数中添加该 IP 或网段。若该 IP 或网段已存在过期时间，则新时间将覆盖原有时间。</li>  使用 remove 操作时注意： <li> 删除 IP 或网段时，相关的未过期的定时过期时间也会被删除；</li> <li> 删除定时过期时间时，仅能删除当前未过期的时间。</li>  使用 update 操作时注意： <li> 替换 Content 内容时，不在 Content 中的 IP 或网段的未过期时间会被删除；</li> <li> 替换 IPExpireInfo 内容时，IPExpireInfo 中的 IP 或网段必须在 Content 中或在 IP 组中存在。</li>
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
type ModifySecurityJSInjectionRuleRequestParams struct {
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// JavaScript 注入规则列表。
	JSInjectionRules []*JSInjectionRule `json:"JSInjectionRules,omitnil,omitempty" name:"JSInjectionRules"`
}

type ModifySecurityJSInjectionRuleRequest struct {
	*tchttp.BaseRequest
	
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// JavaScript 注入规则列表。
	JSInjectionRules []*JSInjectionRule `json:"JSInjectionRules,omitnil,omitempty" name:"JSInjectionRules"`
}

func (r *ModifySecurityJSInjectionRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySecurityJSInjectionRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "JSInjectionRules")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifySecurityJSInjectionRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySecurityJSInjectionRuleResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifySecurityJSInjectionRuleResponse struct {
	*tchttp.BaseResponse
	Response *ModifySecurityJSInjectionRuleResponseParams `json:"Response"`
}

func (r *ModifySecurityJSInjectionRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySecurityJSInjectionRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySecurityPolicyRequestParams struct {
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 安全策略配置。<li>当 SecurityPolicy 参数中的 ExceptionRules 被设置时，SecurityConfig 参数中的 ExceptConfig 将被忽略；</li><li>当 SecurityPolicy 参数中的 CustomRules 被设置时，SecurityConfig 参数中的 AclConfig、 IpTableConfig 将被忽略；</li><li>当 SecurityPolicy 参数中的 HttpDDoSProtection 和 RateLimitingRules 被设置时，SecurityConfig 参数中的 RateLimitConfig 将被忽略；</li><li>当 SecurityPolicy 参数中的 ManagedRule 被设置时，SecurityConfig 参数中的 WafConfig 将被忽略；</li><li>当 SecurityPolicy 参数中的 BotManagement 被设置时，SecurityConfig 参数中的 BotConfig 将被忽略；</li><li>对于例外规则、自定义规则、速率限制、托管规则以及 Bot 管理策略配置建议使用 SecurityPolicy 参数进行设置。</li>
	SecurityConfig *SecurityConfig `json:"SecurityConfig,omitnil,omitempty" name:"SecurityConfig"`

	// 安全策略配置。对 Web 例外规则、防护自定义策略、速率规则、托管规则和 Bot 管理配置建议使用，支持表达式语法对安全策略进行配置。
	SecurityPolicy *SecurityPolicy `json:"SecurityPolicy,omitnil,omitempty" name:"SecurityPolicy"`

	// 安全策略类型，可使用以下参数值： <li>ZoneDefaultPolicy：用于指定站点级策略；</li><li>Template：用于指定策略模板，需要同时指定 TemplateId 参数；</li><li>Host：用于指定域名级策略（注意：当使用域名来指定域名服务策略时，仅支持已经应用了域名级策略的域名服务或者策略模板）。</li>
	Entity *string `json:"Entity,omitnil,omitempty" name:"Entity"`

	// 指定域名。当 Entity 参数值为 Host 时，使用本参数指定的域名级策略，例如：使用 www.example.com ，配置该域名的域名级策略。
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`

	// 指定策略模板 ID。当 Entity 参数值为 Template 时，使用本参数指定策略模板的 ID。
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`
}

type ModifySecurityPolicyRequest struct {
	*tchttp.BaseRequest
	
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 安全策略配置。<li>当 SecurityPolicy 参数中的 ExceptionRules 被设置时，SecurityConfig 参数中的 ExceptConfig 将被忽略；</li><li>当 SecurityPolicy 参数中的 CustomRules 被设置时，SecurityConfig 参数中的 AclConfig、 IpTableConfig 将被忽略；</li><li>当 SecurityPolicy 参数中的 HttpDDoSProtection 和 RateLimitingRules 被设置时，SecurityConfig 参数中的 RateLimitConfig 将被忽略；</li><li>当 SecurityPolicy 参数中的 ManagedRule 被设置时，SecurityConfig 参数中的 WafConfig 将被忽略；</li><li>当 SecurityPolicy 参数中的 BotManagement 被设置时，SecurityConfig 参数中的 BotConfig 将被忽略；</li><li>对于例外规则、自定义规则、速率限制、托管规则以及 Bot 管理策略配置建议使用 SecurityPolicy 参数进行设置。</li>
	SecurityConfig *SecurityConfig `json:"SecurityConfig,omitnil,omitempty" name:"SecurityConfig"`

	// 安全策略配置。对 Web 例外规则、防护自定义策略、速率规则、托管规则和 Bot 管理配置建议使用，支持表达式语法对安全策略进行配置。
	SecurityPolicy *SecurityPolicy `json:"SecurityPolicy,omitnil,omitempty" name:"SecurityPolicy"`

	// 安全策略类型，可使用以下参数值： <li>ZoneDefaultPolicy：用于指定站点级策略；</li><li>Template：用于指定策略模板，需要同时指定 TemplateId 参数；</li><li>Host：用于指定域名级策略（注意：当使用域名来指定域名服务策略时，仅支持已经应用了域名级策略的域名服务或者策略模板）。</li>
	Entity *string `json:"Entity,omitnil,omitempty" name:"Entity"`

	// 指定域名。当 Entity 参数值为 Host 时，使用本参数指定的域名级策略，例如：使用 www.example.com ，配置该域名的域名级策略。
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`

	// 指定策略模板 ID。当 Entity 参数值为 Template 时，使用本参数指定策略模板的 ID。
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
	delete(f, "SecurityPolicy")
	delete(f, "Entity")
	delete(f, "Host")
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
type ModifyWebSecurityTemplateRequestParams struct {
	// 站点 ID。需要传入目标策略模板在访问权限上归属的站点，可使用 DescribeWebSecurityTemplates 接口查询策略模板归属的站点。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 策略模板 ID。
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// 修改后的策略模板名称。由中文、英文、数字和下划线组成，不能以下划线开头，且长度不能超过32个字符。字段为空时则不修改。
	TemplateName *string `json:"TemplateName,omitnil,omitempty" name:"TemplateName"`

	// 安全策略模板配置内容。值为空时不修改；没有传入的子模块结构不会被修改。目前支持 Web 防护模块中的例外规则、自定义规则、速率限制规则和托管规则配置，通过表达式语法对安全策略进行配置。 Bot 管理规则配置暂不支持，正在开发中。
	// 特别说明：当入参某个子模块结构时，请确保携带所有需要保留的规则内容，未传入规则内容视为删除。
	SecurityPolicy *SecurityPolicy `json:"SecurityPolicy,omitnil,omitempty" name:"SecurityPolicy"`
}

type ModifyWebSecurityTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 站点 ID。需要传入目标策略模板在访问权限上归属的站点，可使用 DescribeWebSecurityTemplates 接口查询策略模板归属的站点。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 策略模板 ID。
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// 修改后的策略模板名称。由中文、英文、数字和下划线组成，不能以下划线开头，且长度不能超过32个字符。字段为空时则不修改。
	TemplateName *string `json:"TemplateName,omitnil,omitempty" name:"TemplateName"`

	// 安全策略模板配置内容。值为空时不修改；没有传入的子模块结构不会被修改。目前支持 Web 防护模块中的例外规则、自定义规则、速率限制规则和托管规则配置，通过表达式语法对安全策略进行配置。 Bot 管理规则配置暂不支持，正在开发中。
	// 特别说明：当入参某个子模块结构时，请确保携带所有需要保留的规则内容，未传入规则内容视为删除。
	SecurityPolicy *SecurityPolicy `json:"SecurityPolicy,omitnil,omitempty" name:"SecurityPolicy"`
}

func (r *ModifyWebSecurityTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyWebSecurityTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "TemplateId")
	delete(f, "TemplateName")
	delete(f, "SecurityPolicy")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyWebSecurityTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyWebSecurityTemplateResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyWebSecurityTemplateResponse struct {
	*tchttp.BaseResponse
	Response *ModifyWebSecurityTemplateResponseParams `json:"Response"`
}

func (r *ModifyWebSecurityTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyWebSecurityTemplateResponse) FromJsonString(s string) error {
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

	// 同名站点标识。限制输入数字、英文、"." 、"-" 和 "_"，长度 200 个字符以内。
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

	// 同名站点标识。限制输入数字、英文、"." 、"-" 和 "_"，长度 200 个字符以内。
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

	// 网络错误日志记录配置。不填写表示保持原有配置。
	NetworkErrorLogging *NetworkErrorLogging `json:"NetworkErrorLogging,omitnil,omitempty" name:"NetworkErrorLogging"`

	// 图片优化配置。
	// 不填写表示关闭。
	ImageOptimize *ImageOptimize `json:"ImageOptimize,omitnil,omitempty" name:"ImageOptimize"`

	// 标准 Debug 配置。
	StandardDebug *StandardDebug `json:"StandardDebug,omitnil,omitempty" name:"StandardDebug"`

	// 视频即时处理配置。不填写表示保持原有配置。
	JITVideoProcess *JITVideoProcess `json:"JITVideoProcess,omitnil,omitempty" name:"JITVideoProcess"`
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

	// 网络错误日志记录配置。不填写表示保持原有配置。
	NetworkErrorLogging *NetworkErrorLogging `json:"NetworkErrorLogging,omitnil,omitempty" name:"NetworkErrorLogging"`

	// 图片优化配置。
	// 不填写表示关闭。
	ImageOptimize *ImageOptimize `json:"ImageOptimize,omitnil,omitempty" name:"ImageOptimize"`

	// 标准 Debug 配置。
	StandardDebug *StandardDebug `json:"StandardDebug,omitnil,omitempty" name:"StandardDebug"`

	// 视频即时处理配置。不填写表示保持原有配置。
	JITVideoProcess *JITVideoProcess `json:"JITVideoProcess,omitnil,omitempty" name:"JITVideoProcess"`
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
	delete(f, "NetworkErrorLogging")
	delete(f, "ImageOptimize")
	delete(f, "StandardDebug")
	delete(f, "JITVideoProcess")
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

type MultiPathGateway struct {
	// 网关 ID。
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// 网关名。
	GatewayName *string `json:"GatewayName,omitnil,omitempty" name:"GatewayName"`

	// 网关类型，取值有：
	// <li> cloud：云上网关，腾讯云创建和管理的网关。</li>
	// <li> private：自有网关，用户部署的私有网关。</li>
	GatewayType *string `json:"GatewayType,omitnil,omitempty" name:"GatewayType"`

	// 网关端口，范围 1～65535（除去 8888 ）。
	GatewayPort *int64 `json:"GatewayPort,omitnil,omitempty" name:"GatewayPort"`

	// 网关状态，取值有：
	// <li> creating : 创建中；</li>
	// <li> online : 在线；</li>
	// <li> offline : 离线；</li>
	// <li> disable : 已停用。</li>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 网关 IP， 格式为 IPv4。
	GatewayIP *string `json:"GatewayIP,omitnil,omitempty" name:"GatewayIP"`

	// 网关地域 Id，可以从接口 DescribeMultiPathGatewayRegions 获取 RegionId 列表。
	RegionId *string `json:"RegionId,omitnil,omitempty" name:"RegionId"`

	// 线路信息，当查询网关信息详情 DescribeMultiPathGateway 的时候会返回，当查询网关列表 DescribeMultiPathGateways 的时候不会返回。
	Lines []*MultiPathGatewayLine `json:"Lines,omitnil,omitempty" name:"Lines"`

	// 网关回源 IP 列表发生了变化是否需要重新确认，取值有：<li>true：回源 IP 列表发生了变化，需要确认；</li><li>false：回源 IP 列表未发生变化，无需确认。</li>
	NeedConfirm *string `json:"NeedConfirm,omitnil,omitempty" name:"NeedConfirm"`
}

type MultiPathGatewayCurrentOriginACL struct {
	// 回源 IP 网段详情。
	EntireAddresses *Addresses `json:"EntireAddresses,omitnil,omitempty" name:"EntireAddresses"`

	// 版本号。
	Version *int64 `json:"Version,omitnil,omitempty" name:"Version"`

	// 本参数用于记录当前版本生效前是否完成「我已更新至最新回源 IP 网段」的确认。取值有：
	// <li>true：已完成更新至最新回源 IP 的确认；</li>
	// <li>false：未完成更新至最新回源 IP 的确认；</li>
	// 注意：本参数返回 false 时，请及时确认您的源站防火墙配置是否已更新至最新的回源 IP 网段，以避免出现回源失败。
	IsPlaned *string `json:"IsPlaned,omitnil,omitempty" name:"IsPlaned"`
}

type MultiPathGatewayLine struct {
	// 线路 ID ， 其中 line-0 和 line-1 为系统内置线路 ID，取值有:
	// <li> line-0：直连线路，不支持添加、编辑和删除；</li>
	// <li> line-1： EdgeOne 四层代理线路，支持修改实例和规则，不支持删除；</li>
	// <li> line-2 及以上：EdgeOne 四层代理线路或者自定义线路，支持修改、删除实例和规则。</li>
	LineId *string `json:"LineId,omitnil,omitempty" name:"LineId"`

	// 线路类型，取值有：
	// <li>direct ：直连线路，不支持编辑、不支持删除；</li>
	// <li>proxy ：EdgeOne 四层代理线路，支持编辑修改实例和规则，不支持删除；</li>
	// <li>custom ：自定义线路，支持编辑、支持删除。</li>
	LineType *string `json:"LineType,omitnil,omitempty" name:"LineType"`

	// 线路地址，格式为 host:port 。
	LineAddress *string `json:"LineAddress,omitnil,omitempty" name:"LineAddress"`

	// 四层代理实例 ID  ，当线路类型 LineType 取值为 proxy（EdgeOne 四层代理）返回。
	ProxyId *string `json:"ProxyId,omitnil,omitempty" name:"ProxyId"`

	// 转发规则 ID ，当线路类型 LineType 取值为 proxy（EdgeOne 四层代理）返回。
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`
}

type MultiPathGatewayNextOriginACL struct {
	// 版本号。
	Version *int64 `json:"Version,omitnil,omitempty" name:"Version"`

	// 回源 IP 网段详情。
	EntireAddresses *Addresses `json:"EntireAddresses,omitnil,omitempty" name:"EntireAddresses"`

	// 最新回源 IP 网段相较于 MultiPathGatewayCurrentOrginACL 中回源 IP 网段新增的部分。
	AddedAddresses *Addresses `json:"AddedAddresses,omitnil,omitempty" name:"AddedAddresses"`

	// 最新回源 IP 网段相较于 MultiPathGatewayCurrentOrginACL 中回源 IP 网段删减的部分。
	RemovedAddresses *Addresses `json:"RemovedAddresses,omitnil,omitempty" name:"RemovedAddresses"`

	// 最新回源 IP 网段相较于 MultiPathGatewayCurrentOrginACL 中回源 IP 网段无变化的部分。
	NoChangeAddresses *Addresses `json:"NoChangeAddresses,omitnil,omitempty" name:"NoChangeAddresses"`
}

type MultiPathGatewayOriginACLInfo struct {
	// 当前生效的回源 IP 网段。
	MultiPathGatewayCurrentOriginACL *MultiPathGatewayCurrentOriginACL `json:"MultiPathGatewayCurrentOriginACL,omitnil,omitempty" name:"MultiPathGatewayCurrentOriginACL"`

	// 当回源 IP 网段发生更新时，该字段会返回下一个版本将要生效的回源 IP 网段，包含与当前回源 IP 网段的对比。无更新时该字段为空。
	MultiPathGatewayNextOriginACL *MultiPathGatewayNextOriginACL `json:"MultiPathGatewayNextOriginACL,omitnil,omitempty" name:"MultiPathGatewayNextOriginACL"`
}

type MutualTLS struct {
	// 双向认证配置开关，取值有：
	// <li>on：开启；</li>
	// <li>off：关闭。</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// 双向认证证书列表。
	// 注意：MutualTLS 在 ModifyHostsCertificate 作为入参使用时，该参数传入对应证书的 CertId 即可。您可以前往 [SSL 证书列表](https://console.cloud.tencent.com/ssl) 查看 CertId。
	CertInfos []*CertificateInfo `json:"CertInfos,omitnil,omitempty" name:"CertInfos"`
}

type NSDetail struct {
	// 是否开启 CNAME 加速，取值有：
	// <li> enabled：开启；</li>
	// <li> disabled：关闭。</li>
	CnameSpeedUp *string `json:"CnameSpeedUp,omitnil,omitempty" name:"CnameSpeedUp"`

	// 是否存在同名站点，取值有：
	// <li> 0：不存在同名站点；</li>
	// <li> 1：已存在同名站点。</li>
	IsFake *int64 `json:"IsFake,omitnil,omitempty" name:"IsFake"`

	// 归属权验证信息。针对 NS 接入类型的站点，将当前的 NS 服务器切换至腾讯云 EdgeOne 指定的 NS 服务器，即视为通过归属权验证。详情请参考 [站点/域名归属权验证](https://cloud.tencent.com/document/product/1552/70789) 。
	OwnershipVerification *OwnershipVerification `json:"OwnershipVerification,omitnil,omitempty" name:"OwnershipVerification"`

	// 由 EdgeOne 检测到的站点当前正在使用的 NS 服务器列表。
	OriginalNameServers []*string `json:"OriginalNameServers,omitnil,omitempty" name:"OriginalNameServers"`

	// 腾讯云 EdgeOne 分配的 NS 服务器列表。需要将当前站点 NS 服务器指向该地址，站点才能生效。
	NameServers []*string `json:"NameServers,omitnil,omitempty" name:"NameServers"`

	// 用户自定义 NS 服务器域名信息。如果启用了自定义 NS 服务，需要在域名注册厂商内将 NS 指向该地址。
	VanityNameServers *VanityNameServers `json:"VanityNameServers,omitnil,omitempty" name:"VanityNameServers"`

	// 用户自定义 NS 服务器对应的 IP 地址信息。
	VanityNameServersIps []*VanityNameServersIps `json:"VanityNameServersIps,omitnil,omitempty" name:"VanityNameServersIps"`
}

type NetworkErrorLogging struct {
	// 是否开启网络错误日志记录配置，取值有：
	// <li>on：开启；</li>
	// <li>off：关闭。</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`
}

type NetworkErrorLoggingParameters struct {
	// 网络错误日志记录配置开关，取值有：
	// <li>on：开启；</li>
	// <li>off：关闭。</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`
}

type NextOriginACL struct {
	// 版本号。
	Version *string `json:"Version,omitnil,omitempty" name:"Version"`

	// 版本生效时间，时间是北京时间 UTC+8， 遵循 ISO 8601 标准的日期和时间格式。
	PlannedActiveTime *string `json:"PlannedActiveTime,omitnil,omitempty" name:"PlannedActiveTime"`

	// 回源 IP 网段详情。
	EntireAddresses *Addresses `json:"EntireAddresses,omitnil,omitempty" name:"EntireAddresses"`

	// 最新回源 IP 网段相较于 CurrentOrginACL 中回源 IP 网段新增的部分。
	AddedAddresses *Addresses `json:"AddedAddresses,omitnil,omitempty" name:"AddedAddresses"`

	// 最新回源 IP 网段相较于 CurrentOrginACL 中回源 IP 网段删减的部分。
	RemovedAddresses *Addresses `json:"RemovedAddresses,omitnil,omitempty" name:"RemovedAddresses"`

	// 最新回源 IP 网段相较于 CurrentOrginACL 中回源 IP 网段无变化的部分。
	NoChangeAddresses *Addresses `json:"NoChangeAddresses,omitnil,omitempty" name:"NoChangeAddresses"`
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

type OCSPStaplingParameters struct {
	// OCSP 装订配置开关，取值有：
	// <li>on：开启；</li>
	// <li>off：关闭。</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`
}

type OfflineCache struct {
	// 离线缓存是否开启，取值有：
	// <li>on：开启；</li>
	// <li>off：关闭。</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`
}

type OfflineCacheParameters struct {
	// 离线缓存开关，取值有：
	// <li>on：开启；</li>
	// <li>off：关闭。</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`
}

type Origin struct {
	// 主源站列表。
	Origins []*string `json:"Origins,omitnil,omitempty" name:"Origins"`

	// 备源站列表。
	BackupOrigins []*string `json:"BackupOrigins,omitnil,omitempty" name:"BackupOrigins"`

	// 回源协议配置，取值有：
	// <li>http：强制 http 回源；</li>
	// <li>follow：协议跟随回源；</li>
	// <li>https：强制 https 回源。</li>
	OriginPullProtocol *string `json:"OriginPullProtocol,omitnil,omitempty" name:"OriginPullProtocol"`

	// 源站为腾讯云 COS 时，是否为私有访问 bucket，取值有：
	// <li>on：私有访问；</li>
	// <li>off：公共访问。</li>
	CosPrivateAccess *string `json:"CosPrivateAccess,omitnil,omitempty" name:"CosPrivateAccess"`
}

type OriginACLEntity struct {
	// 实例类型，取值有：
	// - l7：七层加速域名；
	// - l4：四层代理实例。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 实例详情，取值有：
	// - 当 Type = l7 时，请填写七层加速域名；
	// - 当 Type = l4 时，请填写四层代理实例 ID。
	Instances []*string `json:"Instances,omitnil,omitempty" name:"Instances"`

	// 操作模式，取值有：
	// <li>enable：启用；</li>
	// <li>disable：停用。</li>
	OperationMode *string `json:"OperationMode,omitnil,omitempty" name:"OperationMode"`
}

type OriginACLInfo struct {
	// 启用了特定回源 IP 网段回源的七层加速域名列表。源站防护未开启时为空。
	L7Hosts []*string `json:"L7Hosts,omitnil,omitempty" name:"L7Hosts"`

	// 启用了特定回源 IP 网段回源的四层代理实例列表。源站防护未开启时为空。
	L4ProxyIds []*string `json:"L4ProxyIds,omitnil,omitempty" name:"L4ProxyIds"`

	// 当前生效的回源 IP 网段。源站防护未开启时为空。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CurrentOriginACL *CurrentOriginACL `json:"CurrentOriginACL,omitnil,omitempty" name:"CurrentOriginACL"`

	// 当回源 IP 网段发生更新时，该字段会返回下一个版本将要生效的回源 IP 网段，包含与当前回源 IP 网段的对比。无更新或者源站防护未开启时该字段为空。
	// 注意：此字段可能返回 null，表示取不到有效值。
	NextOriginACL *NextOriginACL `json:"NextOriginACL,omitnil,omitempty" name:"NextOriginACL"`

	// 源站防护状态，取值有：
	// <li>online：已生效；</li>
	// <li>offline：已停用；</li>
	// <li>updating: 配置部署中。</li>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`
}

type OriginAuthenticationParameters struct {
	// 回源鉴权请求属性。
	RequestProperties []*OriginAuthenticationRequestProperties `json:"RequestProperties,omitnil,omitempty" name:"RequestProperties"`
}

type OriginAuthenticationRequestProperties struct {
	// 设置回源鉴权参数类型，取值有：<li>QueryString：表示设置回源鉴权参数类型为查询字符串；</li><li>Header：表示设置回源鉴权参数类型为请求头。</li>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 设置回源鉴权类型对应的参数名称。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 设置回源鉴权类型对应的参数值。
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type OriginCertificateVerify struct {
	// 源站证书校验模式。取值有：<li>disable:禁用源站证书校验。</li><li>custom_ca:使用指定受信任 CA 证书校验。</li>
	VerificationMode *string `json:"VerificationMode,omitnil,omitempty" name:"VerificationMode"`

	// 指定受信任的 CA 证书列表，源站证书需要由该 CA 签发才能校验通过。 注意：仅当 VerificationMode 为 custom_ca 时，需要传入该参数，指定受信任的CA证书信息。
	// OriginCertificateVerify 在 ModifyHostsCertificate 作为入参使用时，该参数传入对应证书的 CertId 即可。您可以前往 [SSL 证书列表](https://console.cloud.tencent.com/ssl) 查看 CertId。	
	CustomCACerts []*CertificateInfo `json:"CustomCACerts,omitnil,omitempty" name:"CustomCACerts"`
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

	// 当前配置的回源 HOST 头。
	HostHeader *string `json:"HostHeader,omitnil,omitempty" name:"HostHeader"`

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

	// 云点播回源范围，当 OriginType = VOD 时该参数会返回值。取值有:<li>all：当前源站对应的云点播应用内所有文件，默认值为 all；</li> <li>bucket：当前源站对应的云点播应用下指定某一个存储桶内的文件。通过参数 VodBucketId 来指定存储桶。
	// </li>
	VodOriginScope *string `json:"VodOriginScope,omitnil,omitempty" name:"VodOriginScope"`

	// 云点播存储桶 ID，该参数当 OriginType = VOD 且 VodOriginScope = bucket 时必填。数据来源：云点播专业版应用下存储桶的存储 ID 。
	VodBucketId *string `json:"VodBucketId,omitnil,omitempty" name:"VodBucketId"`
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
	HostHeader *string `json:"HostHeader,omitnil,omitempty" name:"HostHeader"`
}

type OriginGroupHealthStatus struct {
	// 源站组 ID。
	OriginGroupID *string `json:"OriginGroupID,omitnil,omitempty" name:"OriginGroupID"`

	// 源站组名。
	OriginGroupName *string `json:"OriginGroupName,omitnil,omitempty" name:"OriginGroupName"`

	// 源站组类型，取值有：
	// <li>HTTP：HTTP 专用型；</li>
	// <li>GENERAL：通用型。</li>
	OriginType *string `json:"OriginType,omitnil,omitempty" name:"OriginType"`

	// 优先级。
	Priority *string `json:"Priority,omitnil,omitempty" name:"Priority"`

	// 源站组里各源站的健康状态。
	OriginHealthStatus []*OriginHealthStatus `json:"OriginHealthStatus,omitnil,omitempty" name:"OriginHealthStatus"`
}

type OriginGroupHealthStatusDetail struct {
	// 源站组 ID。
	OriginGroupId *string `json:"OriginGroupId,omitnil,omitempty" name:"OriginGroupId"`

	// 根据所有探测区域的结果综合决策出来的源站组下各个源站的健康状态。超过一半的地域判定该源站不健康，则对应状态为不健康，否则为健康。
	OriginHealthStatus []*OriginHealthStatus `json:"OriginHealthStatus,omitnil,omitempty" name:"OriginHealthStatus"`

	// 各个健康检查区域下源站的健康状态。
	CheckRegionHealthStatus []*CheckRegionHealthStatus `json:"CheckRegionHealthStatus,omitnil,omitempty" name:"CheckRegionHealthStatus"`
}

type OriginGroupInLoadBalancer struct {
	// 优先级，填写格式为 "priority_" + "数字"，最高优先级为 "priority_1"。参考取值有：
	// <li>priority_1：第一优先级；</li>
	// <li>priority_2：第二优先级；</li>
	// <li>priority_3：第三优先级。</li>其他优先级可以将数字递增，最多可以递增至 "priority_10"。
	Priority *string `json:"Priority,omitnil,omitempty" name:"Priority"`

	// 源站组 ID。
	OriginGroupId *string `json:"OriginGroupId,omitnil,omitempty" name:"OriginGroupId"`
}

type OriginGroupReference struct {
	// 引用服务类型，取值有：
	// <li>acceleration-domain: 加速域名；</li>
	// <li>rule-engine: 规则引擎；</li>
	// <li>load-balancer: 负载均衡；</li>
	// <li>application-proxy: 四层代理。</li>
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// 引用类型的实例ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 引用类型的实例名称。
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 引用站点ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 引用站点名称。
	ZoneName *string `json:"ZoneName,omitnil,omitempty" name:"ZoneName"`

	// 引用站点别名。
	AliasZoneName *string `json:"AliasZoneName,omitnil,omitempty" name:"AliasZoneName"`
}

type OriginHealthStatus struct {
	// 源站。
	Origin *string `json:"Origin,omitnil,omitempty" name:"Origin"`

	// 源站健康状态，取值有：
	// <li>Healthy：健康；</li>
	// <li>Unhealthy：不健康；</li>
	// <li>Undetected：未探测到数据。</li>
	Healthy *string `json:"Healthy,omitnil,omitempty" name:"Healthy"`
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

	// 自定义回源 HOST 头，该参数仅当 OriginType=IP_DOMAIN 时生效。
	// 如果 OriginType=COS 或 AWS_S3 时，回源 HOST 头将与源站域名保持一致。
	// 如果OriginType=ORIGIN_GROUP 时，回源 HOST 头遵循源站组内配置，如果没有配置则默认为加速域名。
	// 如果 OriginType=VOD 或 SPACE 时，无需配置该头部，按对应的回源域名生效。
	HostHeader *string `json:"HostHeader,omitnil,omitempty" name:"HostHeader"`

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

	// 云点播回源范围，该参数当 OriginType = VOD 时生效。取值有：<li>all：当前源站对应的云点播应用内所有文件，默认值为 all；</li><li>bucket：当前源站对应的云点播应用下指定某一个存储桶内的文件。通过参数 VodBucketId 来指定存储桶。
	// </li>
	VodOriginScope *string `json:"VodOriginScope,omitnil,omitempty" name:"VodOriginScope"`

	// VOD 存储桶 ID，该参数当 OriginType = VOD 且 VodOriginScope = bucket 时必填。数据来源：云点播专业版应用下存储桶的存储 ID 。
	VodBucketId *string `json:"VodBucketId,omitnil,omitempty" name:"VodBucketId"`
}

type OriginPrivateParameters struct {
	// 鉴权参数 Access Key ID。
	AccessKeyId *string `json:"AccessKeyId,omitnil,omitempty" name:"AccessKeyId"`

	// 鉴权参数 Secret Access Key。
	SecretAccessKey *string `json:"SecretAccessKey,omitnil,omitempty" name:"SecretAccessKey"`

	// 鉴权版本。取值有：
	// <li>v2：v2版本；</li>
	// <li>v4：v4版本。</li>
	SignatureVersion *string `json:"SignatureVersion,omitnil,omitempty" name:"SignatureVersion"`

	// 存储桶地域。
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`
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

type OriginPullProtocolParameters struct {
	// 回源协议配置，取值有：
	// <li>http：使用 HTTP 协议回源；</li>
	// <li>https：使用 HTTPS 协议回源；</li>
	// <li>follow：协议跟随。</li>
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`
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

	// 【源站权重】：用于控制流量分配优先级的参数，取值范围：0-100（整数）：<li>空值：不设置权重，系统按默认策略调度；</li><li>0 值：明确设置权重为0，流量将不会分配到该源站，注意事项：必须确保至少有一个源站的权重值大于0；</li><li>正常值：数值越大分配流量越多 ；</li>
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
	// CNAME 、无域名接入时，使用 DNS 解析验证时所需的信息。详情参考 [站点/域名归属权验证
	// ](https://cloud.tencent.com/document/product/1552/70789#7af6ecf8-afca-4e35-8811-b5797ed1bde5)。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DnsVerification *DnsVerification `json:"DnsVerification,omitnil,omitempty" name:"DnsVerification"`

	// CNAME 、无域名接入时，使用文件验证时所需的信息。详情参考 [站点/域名归属权验证
	// ](https://cloud.tencent.com/document/product/1552/70789#7af6ecf8-afca-4e35-8811-b5797ed1bde5)。
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileVerification *FileVerification `json:"FileVerification,omitnil,omitempty" name:"FileVerification"`

	// NS 接入，切换 DNS 服务器所需的信息。详情参考 [修改 DNS 服务器](https://cloud.tencent.com/document/product/1552/90452)。
	// 注意：此字段可能返回 null，表示取不到有效值。
	NsVerification *NsVerification `json:"NsVerification,omitnil,omitempty" name:"NsVerification"`
}

type PartialModule struct {
	// 模块名称，取值为：
	// <li>managed-rule：托管规则 Id；</li>
	// <li>managed-group：托管规则组；</li>
	// <li>waf：待废弃，托管规则。</li>
	Module *string `json:"Module,omitnil,omitempty" name:"Module"`

	// 模块下的需要例外的具体规则ID列表。
	Include []*int64 `json:"Include,omitnil,omitempty" name:"Include"`
}

type Plan struct {
	// 套餐类型。取值有：
	// <li>plan-trial: 试用版套餐；</li>
	// <li>plan-personal: 个人版套餐；</li>
	// <li>plan-basic: 基础版套餐；</li>
	// <li>plan-standard: 标准版套餐；</li>
	// <li>plan-enterprise-v2: 企业版套餐；</li>
	// <li>plan-enterprise-model-a: 企业版 Model A 套餐。</li>
	// <li>plan-enterprise: 旧企业版套餐。</li>
	PlanType *string `json:"PlanType,omitnil,omitempty" name:"PlanType"`

	// 套餐 ID。形如 edgeone-2y041pblwaxe。
	PlanId *string `json:"PlanId,omitnil,omitempty" name:"PlanId"`

	// 服务区域，取值有：
	// <li>mainland: 中国大陆；</li>
	// <li>overseas: 全球（不包括中国大陆）；</li>
	// <li>global: 全球（包括中国大陆）。</li>
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 自动续费开关。取值有：
	// <li>true: 已开启自动续费；</li>
	// <li>false: 未开启自动续费。</li>
	AutoRenewal *bool `json:"AutoRenewal,omitnil,omitempty" name:"AutoRenewal"`

	// 套餐状态，取值有：
	// <li>normal：正常状态；</li>
	// <li>expiring-soon：即将到期状态；</li>
	// <li>expired：到期状态；</li>
	// <li>isolated：隔离状态；</li>
	// <li>overdue-isolated：欠费隔离状态。</li>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 付费类型，取值有：
	// <li>0: 后付费；</li>
	// <li>1: 预付费。</li>
	PayMode *int64 `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// 套餐绑定的站点信息，包括站点id和站点名称，站点状态。
	ZonesInfo []*ZoneInfo `json:"ZonesInfo,omitnil,omitempty" name:"ZonesInfo"`

	// 套餐内智能加速请求数规格，单位：次。
	SmartRequestCapacity *int64 `json:"SmartRequestCapacity,omitnil,omitempty" name:"SmartRequestCapacity"`

	// 套餐内VAU规格，单位：个。
	VAUCapacity *int64 `json:"VAUCapacity,omitnil,omitempty" name:"VAUCapacity"`

	// 套餐内内容加速流量规格，单位：字节。
	AccTrafficCapacity *int64 `json:"AccTrafficCapacity,omitnil,omitempty" name:"AccTrafficCapacity"`

	// 套餐内智能加速流量规格，单位：字节。
	SmartTrafficCapacity *int64 `json:"SmartTrafficCapacity,omitnil,omitempty" name:"SmartTrafficCapacity"`

	// 套餐内DDoS防护流量规格，单位：字节。
	DDoSTrafficCapacity *int64 `json:"DDoSTrafficCapacity,omitnil,omitempty" name:"DDoSTrafficCapacity"`

	// 套餐内安全流量规格，单位：字节。
	SecTrafficCapacity *int64 `json:"SecTrafficCapacity,omitnil,omitempty" name:"SecTrafficCapacity"`

	// 套餐内安全请求数规格，单位：次。
	SecRequestCapacity *int64 `json:"SecRequestCapacity,omitnil,omitempty" name:"SecRequestCapacity"`

	// 套餐内四层加速流量规格，单位：字节。
	L4TrafficCapacity *int64 `json:"L4TrafficCapacity,omitnil,omitempty" name:"L4TrafficCapacity"`

	// 套餐内中国大陆网络优化流量规格，单位：字节。
	CrossMLCTrafficCapacity *int64 `json:"CrossMLCTrafficCapacity,omitnil,omitempty" name:"CrossMLCTrafficCapacity"`

	// 套餐是否允许绑定新站点，取值有：
	// <li>true: 允许绑定新站点；</li>
	// <li>false: 不允许绑定新站点。</li>
	Bindable *string `json:"Bindable,omitnil,omitempty" name:"Bindable"`

	// 套餐生效时间。
	EnabledTime *string `json:"EnabledTime,omitnil,omitempty" name:"EnabledTime"`

	// 套餐过期时间。
	ExpiredTime *string `json:"ExpiredTime,omitnil,omitempty" name:"ExpiredTime"`

	// 套餐所支持的功能，取值有：<li>ContentAcceleration：内容加速功能；</li><li>SmartAcceleration：智能加速功能；</li><li>L4：四层加速功能；</li><li>Waf：高级 Web 防护；</li><li>QUIC：QUIC功能；</li><li>CrossMLC：中国大陆网络优化功能；</li><li>ProcessMedia：媒体处理功能；</li><li>L4DDoS：四层DDoS防护功能；</li>L7DDoS功能只会出现以下所有规格中的一项<li>L7DDoS.CM30G；七层DDoS防护功能-中国大陆30G保底带宽规格；</li><li>L7DDoS.CM60G；七层DDoS防护功能-中国大陆60G保底带宽规格；</li><li>L7DDoS.CM100G；七层DDoS防护功能-中国大陆100G保底带宽规格；</li><li>L7DDoS.Anycast300G；七层DDoS防护功能-中国大陆以外Anycast300G保底带宽规格；</li><li>L7DDoS.AnycastUnlimited；七层DDoS防护功能-中国大陆以外Anycast无上限全力防护规格；</li><li>L7DDoS.CM30G_Anycast300G；七层DDoS防护功能-中国大陆30G保底带宽规格，中国大陆以外Anycast300G保底带宽规格；</li><li>L7DDoS.CM60G_Anycast300G；七层DDoS防护功能-中国大陆60G保底带宽规格，中国大陆以外Anycast300G保底带宽规格；</li><li>L7DDoS.CM100G_Anycast300G；七层DDoS防护功能-中国大陆100G保底带宽规格，中国大陆以外Anycast300G保底带宽规格；</li><li>L7DDoS.CM30G_AnycastUnlimited；七层DDoS防护功能-中国大陆30G保底带宽规格，中国大陆以外Anycast无上限全力防护规格；</li><li>L7DDoS.CM60G_AnycastUnlimited；七层DDoS防护功能-中国大陆60G保底带宽规格，中国大陆以外Anycast无上限全力防护规格；</li><li>L7DDoS.CM100G_AnycastUnlimited；七层DDoS防护功能-中国大陆100G保底带宽规格，中国大陆以外Anycast无上限全力防护规格；</li>
	Features []*string `json:"Features,omitnil,omitempty" name:"Features"`
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

	// 最大限制，该字段仅在 Switch 为 on 时生效，取值在 1MB 和 800MB 之间，单位字节。
	MaxSize *int64 `json:"MaxSize,omitnil,omitempty" name:"MaxSize"`
}

type PostMaxSizeParameters struct {
	// 是否开启 POST 请求上传文件限制，单位为 Byte，平台默认为限制为 32 * 2<sup>20</sup> Byte，取值有：<li>on：开启限制；</li><li>off：关闭限制。</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// POST 请求上传文件流式传输最大限制，该字段仅在 Switch 为 on 时生效，取值在 1MB 和 800MB 之间，单位字节。
	MaxSize *int64 `json:"MaxSize,omitnil,omitempty" name:"MaxSize"`
}

type PrefetchOriginLimit struct {
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 加速域名。
	DomainName *string `json:"DomainName,omitnil,omitempty" name:"DomainName"`

	// 回源限速限制的加速区域。
	// 预热时，该加速区域将会受到配置的Bandwidth值限制。取值有：
	// <li>Overseas：全球可用区（不含中国大陆）；</li>
	// <li>MainlandChina：中国大陆可用区。</li>
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 回源限速带宽。
	// 预热时回到源站的带宽上限值，取值范围 100 - 100,000，单位 Mbps。
	Bandwidth *int64 `json:"Bandwidth,omitnil,omitempty" name:"Bandwidth"`

	// 回源限速限制创建的时间。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 回源限速限制更新的时间。
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`
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

type QUICParameters struct {
	// QUIC 配置开关，取值有：
	// <li>on：开启；</li>
	// <li>off：关闭。</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`
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
	Action *string `json:"Action,omitnil,omitempty" name:"Action"`

	// 使用/排除的url参数数组。
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

type RangeOriginPullParameters struct {
	// 分片回源开关，取值有：
	// <li>on：开启；</li>
	// <li>off：关闭。</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`
}

type RateLimitConfig struct {
	// 开关，取值有：
	// <li>on：开启；</li>
	// <li>off：关闭。</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// 速率限制-用户规则列表。如果为null，默认使用历史配置。
	RateLimitUserRules []*RateLimitUserRule `json:"RateLimitUserRules,omitnil,omitempty" name:"RateLimitUserRules"`

	// 速率限制模板功能。如果为null，默认使用历史配置。
	RateLimitTemplate *RateLimitTemplate `json:"RateLimitTemplate,omitnil,omitempty" name:"RateLimitTemplate"`

	// 智能客户端过滤。如果为null，默认使用历史配置。
	RateLimitIntelligence *RateLimitIntelligence `json:"RateLimitIntelligence,omitnil,omitempty" name:"RateLimitIntelligence"`

	// 速率限制-托管定制规则。如果为null，默认使用历史配置。
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
	Mode *string `json:"Mode,omitnil,omitempty" name:"Mode"`

	// 唯一id。
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// 模板处置方式，取值有：
	// <li>alg：JavaScript挑战；</li>
	// <li>monitor：观察。</li>
	Action *string `json:"Action,omitnil,omitempty" name:"Action"`

	// 惩罚时间，取值范围0-2天，单位秒。
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

type RateLimitingRule struct {
	// 精准速率限制的 ID。<br>通过规则 ID 可支持不同的规则配置操作：<br> <li> <b>增加</b>新规则：ID 为空或不指定 ID 参数；</li><li><b>修改</b>已有规则：指定需要更新/修改的规则 ID；</li><li><b>删除</b>已有规则：RateLimitingRules 参数中，Rules 列表中未包含的已有规则将被删除。</li>
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 精准速率限制的名称。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 精准速率限制的具体内容，需符合表达式语法，详细规范参见产品文档。
	Condition *string `json:"Condition,omitnil,omitempty" name:"Condition"`

	// 速率阈值请求特征的匹配方式， 当 Enabled 为 on 时，此字段必填。<br /><br />当条件有多个时，将组合多个条件共同进行统计计算，条件最多不可超过5条。取值有：<br/><li><b>http.request.ip</b>：客户端 IP；</li><li><b>http.request.xff_header_ip</b>：客户端 IP（优先匹配 XFF 头部）；</li><li><b>http.request.uri.path</b>：请求的访问路径；</li><li><b>http.request.cookies['session']</b>：名称为session的Cookie，其中session可替换为自己指定的参数；</li><li><b>http.request.headers['user-agent']</b>：名称为user-agent的HTTP头部，其中user-agent可替换为自己指定的参数；</li><li><b>http.request.ja3</b>：请求的JA3指纹；</li><li><b>http.request.uri.query['test']</b>：名称为test的URL查询参数，其中test可替换为自己指定的参数。</li> 
	CountBy []*string `json:"CountBy,omitnil,omitempty" name:"CountBy"`

	// 精准速率限制在时间范围内的累计拦截次数，取值范围 1 ~ 100000。
	MaxRequestThreshold *int64 `json:"MaxRequestThreshold,omitnil,omitempty" name:"MaxRequestThreshold"`

	// 统计的时间窗口，取值有：<li>1s：1秒；</li><li>5s：5秒；</li><li>10s：10秒；</li><li>20s：20秒；</li><li>30s：30秒；</li><li>40s：40秒；</li><li>50s：50秒；</li><li>1m：1分钟；</li><li>2m：2分钟；</li><li>5m：5分钟；</li><li>10m：10分钟；</li><li>1h：1小时。</li> 
	CountingPeriod *string `json:"CountingPeriod,omitnil,omitempty" name:"CountingPeriod"`

	// Action 动作的持续时长，单位仅支持：<li>s：秒，取值 1 ~ 120；</li><li>m：分钟，取值 1 ~ 120；</li><li>h：小时，取值 1 ~ 48；</li><li>d：天，取值 1 ~ 30。</li>
	ActionDuration *string `json:"ActionDuration,omitnil,omitempty" name:"ActionDuration"`

	// 精准速率限制的处置方式。取值有：<li>Monitor：观察；</li><li>Deny：拦截，其中DenyActionParameters.Name支持Deny和ReturnCustomPage；</li><li>Challenge：挑战，其中ChallengeActionParameters.Name支持JSChallenge和ManagedChallenge；</li><li>Redirect：重定向至URL；</li>
	Action *SecurityAction `json:"Action,omitnil,omitempty" name:"Action"`

	// 精准速率限制的优先级，范围是 0 ~ 100，默认为 0。
	Priority *int64 `json:"Priority,omitnil,omitempty" name:"Priority"`

	// 精准速率限制规则是否开启。取值有：<li>on：开启；</li><li>off：关闭。</li>
	Enabled *string `json:"Enabled,omitnil,omitempty" name:"Enabled"`
}

type RateLimitingRules struct {
	// 精准速率限制的定义列表。使用 ModifySecurityPolicy 修改 Web 防护配置时: <br> <li>  若未指定 Rules 参数，或 Rules 参数长度为零：清空所有精准速率限制配置。</li> <li> 若 SecurityPolicy 参数中，未指定 RateLimitingRules 参数值：保持已有自定义规则配置，不做修改。</li>
	Rules []*RateLimitingRule `json:"Rules,omitnil,omitempty" name:"Rules"`
}

type RealtimeLogDeliveryTask struct {
	// 实时日志投递任务 ID。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 实时日志投递任务的名称。
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`

	// 实时日志投递任务的状态，取值有： <li>enabled: 已启用；</li> <li>disabled: 已停用；</li><li>deleted: 异常删除状态，请检查目的地腾讯云 CLS 日志集/日志主题是否已被删除。</li>
	DeliveryStatus *string `json:"DeliveryStatus,omitnil,omitempty" name:"DeliveryStatus"`

	// 实时日志投递任务类型，取值有： <li>cls: 推送到腾讯云 CLS；</li> <li>custom_endpoint：推送到自定义 HTTP(S) 地址；</li> <li>s3：推送到 AWS S3 兼容存储桶地址；</li><li>log_analysis：推送到 EdgeOne 日志分析。</li>
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

type RedirectActionParameters struct {
	// 重定向的URL。
	URL *string `json:"URL,omitnil,omitempty" name:"URL"`
}

// Predefined struct for user
type RefreshMultiPathGatewaySecretKeyRequestParams struct {
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`
}

type RefreshMultiPathGatewaySecretKeyRequest struct {
	*tchttp.BaseRequest
	
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`
}

func (r *RefreshMultiPathGatewaySecretKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RefreshMultiPathGatewaySecretKeyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RefreshMultiPathGatewaySecretKeyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RefreshMultiPathGatewaySecretKeyResponseParams struct {
	// 多通道安全加速网关接入密钥。
	SecretKey *string `json:"SecretKey,omitnil,omitempty" name:"SecretKey"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RefreshMultiPathGatewaySecretKeyResponse struct {
	*tchttp.BaseResponse
	Response *RefreshMultiPathGatewaySecretKeyResponseParams `json:"Response"`
}

func (r *RefreshMultiPathGatewaySecretKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RefreshMultiPathGatewaySecretKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
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

type RequestBodyTransferTimeout struct {
	// 正文传输超时时长，取值 5 ~ 120，单位仅支持秒（s）。
	IdleTimeout *string `json:"IdleTimeout,omitnil,omitempty" name:"IdleTimeout"`

	// 正文传输超时时长是否开启。取值有：<li>on：开启；</li><li>off：关闭。</li>
	Enabled *string `json:"Enabled,omitnil,omitempty" name:"Enabled"`
}

type RequestFieldsForException struct {
	// 跳过的具体字段。取值支持：<br/>
	// <li>body.json：JSON 请求内容；此时 Condition 支持 key、value,  TargetField 支持 key、value，例如 { "Scope": "body.json",  "Condition": "", "TargetField": "key" }，表示 JSON 请求内容所有参数跳过 WAF 扫描；</li>
	// <li style="margin-top:5px">cookie：Cookie；此时 Condition 支持 key、value,  TargetField 支持 key、value，例如 { "Scope": "cookie",  "Condition": "${key} in ['account-id'] and ${value} like ['prefix-*']", "TargetField": "value" }，表示 Cookie 参数名称等于account-id 并且参数值通配符匹配 prefix-* 跳过 WAF 扫描；</li>
	// <li style="margin-top:5px">header：HTTP 头部参数；此时 Condition 支持 key、value,  TargetField 支持 key、value，例如 { "Scope": "header",  "Condition": "${key} like ['x-auth-*']", "TargetField": "value" }，表示 header 参数名称通配符匹配 x-auth-* 跳过 WAF 扫描；</li>
	// <li style="margin-top:5px">uri.query：URL 编码内容/查询参数；此时 Condition 支持 key、value,  TargetField 支持 key、value，例如 { "Scope": "uri.query",  "Condition": "${key} in ['action'] and ${value} in ['upload', 'delete']", "TargetField": "value" }，表示 URL 编码内容/查询参数的参数名称等于 action 并且参数值等于 upload 或 delete 跳过 WAF 扫描；</li>
	// <li style="margin-top:5px">uri：请求路径URI；此时 Condition 必须为空， TargetField 支持 query、path、fullpath，例如 { "Scope": "uri",  "Condition": "", "TargetField": "query" }，表示请求路径 URI 仅查询参数跳过 WAF 扫描；</li>
	// <li style="margin-top:5px">body：请求正文内容。此时 Condition 必须为空， TargetField 支持 fullbody、multipart，例如 { "Scope": "body",  "Condition": "", "TargetField": "fullbody" }，表示请求正文内容为完整请求正文跳过 WAF 扫描；</li>
	Scope *string `json:"Scope,omitnil,omitempty" name:"Scope"`

	// 跳过的具体字段的表达式，需要符合表达式语法。<br />
	// Condition  支持表达式配置语法：<li> 按规则的匹配条件表达式语法编写，支持引用 key、value。</li><li> 支持 in、like 操作符，以及 and 逻辑组合。</li>
	// 例如：<li>${key} in ['x-trace-id']：参数名称等于x-trace-id。</li><li>${key} in ['x-trace-id'] and ${value} like ['Bearer *']：参数名称等于x-trace-id并且参数值通配符匹配Bearer *。</li>
	Condition *string `json:"Condition,omitnil,omitempty" name:"Condition"`

	// Scope 参数使用不同取值时，TargetField 表达式中支持的值如下：
	// <li> body.json：支持 key、value</li>
	// <li> cookie：支持 key、value</li>
	// <li> header：支持 key、value</li>
	// <li> uri.query：支持 key、value</li>
	// <li> uri：支持 path、query、fullpath</li>
	// <li> body：支持 fullbody、multipart</li>
	TargetField *string `json:"TargetField,omitnil,omitempty" name:"TargetField"`
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
	// <li>mainland：中国大陆境内；</li>
	// <li>overseas：中国大陆境外。</li>
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

type ResponseSpeedLimitParameters struct {
	// 下载限速模式，取值有：
	// <li>LimitUponDownload：全过程下载限速；</li>
	// <li>LimitAfterSpecificBytesDownloaded：全速下载特定字节后开始限速；</li>
	// <li>LimitAfterSpecificSecondsDownloaded：全速下载特定时间后开始限速。</li>
	Mode *string `json:"Mode,omitnil,omitempty" name:"Mode"`

	// 限速值，指定限速大小，填写含单位的数值或变量。当前支持单位有：KB/s。
	MaxSpeed *string `json:"MaxSpeed,omitnil,omitempty" name:"MaxSpeed"`

	// 限速开始值，可以为下载大小或指定时长，填写含单位的数值或变量，指定下载大小或指定时长。
	// 
	// - 当Mode 取值为 LimitAfterSpecificBytesDownloaded 时，单位取值有： KB；
	// 
	// - 当Mode 取值为 LimitAfterSpecificSecondsDownloaded 时，单位取值有： s。
	StartAt *string `json:"StartAt,omitnil,omitempty" name:"StartAt"`
}

type ReturnCustomPageActionParameters struct {
	// 响应状态码。
	ResponseCode *string `json:"ResponseCode,omitnil,omitempty" name:"ResponseCode"`

	// 响应的自定义页面ID。
	ErrorPageId *string `json:"ErrorPageId,omitnil,omitempty" name:"ErrorPageId"`
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

type RuleBranch struct {
	// [匹配条件](https://cloud.tencent.com/document/product/1552/90438#33f65828-c6c6-4b66-a011-25a20b548d5d)。
	Condition *string `json:"Condition,omitnil,omitempty" name:"Condition"`

	// [操作](https://cloud.tencent.com/document/product/1552/90438#c7bd7e02-9247-4a72-b0e4-11c27cadb198)。<br>注意：Actions 和 SubRules 不可同时为空。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Actions []*RuleEngineAction `json:"Actions,omitnil,omitempty" name:"Actions"`

	// 子规则列表。此列表中时存在多条规则，按照从上往下的顺序依次执行。<br>注意：SubRules 和 Actions 不可同时为空。且当前只支持填写一层 SubRules。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubRules []*RuleEngineSubRule `json:"SubRules,omitnil,omitempty" name:"SubRules"`
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

	// 匹配类型，取值有： <li> filename：文件名； </li> <li> extension：文件后缀； </li> <li> host：HOST； </li> <li> full_url：URL Full，当前站点下完整 URL 路径，必须包含 HTTP 协议，Host 和 路径； </li> <li> url：URL Path，当前站点下 URL 路径的请求； </li><li>client_country：客户端国家/地区；</li> <li> query_string：查询字符串，当前站点下请求 URL 的查询字符串； </li> <li> request_header：HTTP 请求头部。 </li><li> client_ip：客户端 IP。 </li><li> request_protocol：请求协议。 </li><li> request_method：HTTP 请求方法。 </li>
	Target *string `json:"Target,omitnil,omitempty" name:"Target"`

	// 对应匹配类型的参数值，仅在匹配类型为查询字符串或HTTP请求头并且运算符取值为存在或不存在时允许传空数组，对应匹配类型有：
	// <li> 文件后缀：jpg、txt 等文件后缀；</li>
	// <li> 文件名称：例如 foo.jpg 中的 foo；</li>
	// <li> 全部（站点任意请求）：all；</li>
	// <li> HOST：当前站点下的 host ，例如www.maxx55.com；</li>
	// <li> URL Path：当前站点下 URL 路径的请求，例如：/example；</li>
	// <li> URL Full：当前站点下完整 URL 请求，必须包含 HTTP 协议，Host 和 路径，例如：https://www.maxx55.cn/example；</li>
	// <li> 客户端国家/地区：符合 ISO3166 标准的国家/地区标识；</li>
	// <li> 查询字符串: 当前站点下 URL 请求中查询字符串的参数值，例如 lang=cn&version=1 中的 cn 和 1； </li>
	// <li> HTTP 请求头: HTTP 请求头部字段值，例如 Accept-Language:zh-CN,zh;q=0.9中的zh-CN,zh;q=0.9 ；</li>
	// <li> 客户端 IP: 当前请求携带的客户端请求 IP，支持 IPv4/IPv6, 支持 IP 段； </li>
	// <li> 请求协议: 当前请求的协议，取值范围为：HTTP、HTTPS；</li>
	// <li> HTTP 请求方法: 当前请求的方法，取值范围为：GET、HEAD、POST、PUT、DELETE、TRACE、CONNECT、OPTIONS、PATCH、COPY、LOCK、MKCOL、MOVE、PROPFIND、PROPPATCH、UNLOCK。 </li>
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

type RuleEngineAction struct {
	// 操作名称。名称需要与参数结构体对应，例如 Name=Cache，则 CacheParameters 必填。
	// <li>Cache：节点缓存 TTL；</li>
	// <li>CacheKey：自定义 Cache Key；</li>
	// <li>CachePrefresh：缓存预刷新；</li>
	// <li>AccessURLRedirect：访问 URL 重定向；</li>
	// <li>UpstreamURLRewrite：回源 URL 重写；</li>
	// <li>QUIC：QUIC；</li>
	// <li>WebSocket：WebSocket；</li>
	// <li>Authentication：Token 鉴权；</li>
	// <li>MaxAge：浏览器缓存 TTL；</li>
	// <li>StatusCodeCache：状态码缓存 TTL；</li>
	// <li>OfflineCache：离线缓存；</li>
	// <li>SmartRouting：智能加速；</li>
	// <li>RangeOriginPull：分片回源 ；</li>
	// <li>UpstreamHTTP2：HTTP2 回源；</li>
	// <li>HostHeader：Host Header 重写；</li>
	// <li>ForceRedirectHTTPS：访问协议强制 HTTPS 跳转配置；</li>
	// <li>OriginPullProtocol：回源 HTTPS；</li>
	// <li>Compression：智能压缩配置；</li>
	// <li>HSTS：HSTS；</li>
	// <li>ClientIPHeader：存储客户端请求 IP 的头部信息配置；</li>
	// <li>OCSPStapling：OCSP 装订；</li>
	// <li>HTTP2：HTTP2 接入；</li>
	// <li>PostMaxSize：POST 请求上传文件流式传输最大限制配置；</li>
	// <li>ClientIPCountry：回源时携带客户端 IP 所属地域信息；</li>
	// <li>UpstreamFollowRedirect：回源跟随重定向参数配置；</li>
	// <li>UpstreamRequest：回源请求参数；</li>
	// <li>TLSConfig：SSL/TLS 安全；</li>
	// <li>ModifyOrigin：修改源站；</li>
	// <li>HTTPUpstreamTimeout：七层回源超时配置；</li>
	// <li>HttpResponse：HTTP 应答；</li>
	// <li>ErrorPage：自定义错误页面；</li>
	// <li>ModifyResponseHeader：修改 HTTP 节点响应头；</li>
	// <li>ModifyRequestHeader：修改 HTTP 节点请求头；</li>
	// <li>ResponseSpeedLimit：单连接下载限速；</li>
	// <li>SetContentIdentifier：设置内容标识符；</li>
	// <li>Vary：Vary 特性配置；</li>
	// <li>ContentCompression：内容压缩配置；</li>
	// <li>OriginAuthentication：回源鉴权配置。</li>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 节点缓存 TTL 配置参数，当 Name 取值为 Cache 时，该参数必填。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CacheParameters *CacheParameters `json:"CacheParameters,omitnil,omitempty" name:"CacheParameters"`

	// 自定义 Cache Key 配置参数，当 Name 取值为 CacheKey 时，该参数必填。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CacheKeyParameters *CacheKeyParameters `json:"CacheKeyParameters,omitnil,omitempty" name:"CacheKeyParameters"`

	// 缓存预刷新配置参数，当 Name 取值为 CachePrefresh 时，该参数必填。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CachePrefreshParameters *CachePrefreshParameters `json:"CachePrefreshParameters,omitnil,omitempty" name:"CachePrefreshParameters"`

	// 访问 URL 重定向配置参数，当 Name 取值为 AccessURLRedirect 时，该参数必填。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AccessURLRedirectParameters *AccessURLRedirectParameters `json:"AccessURLRedirectParameters,omitnil,omitempty" name:"AccessURLRedirectParameters"`

	// 回源 URL 重写配置参数，当 Name 取值为 UpstreamURLRewrite 时，该参数必填。
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpstreamURLRewriteParameters *UpstreamURLRewriteParameters `json:"UpstreamURLRewriteParameters,omitnil,omitempty" name:"UpstreamURLRewriteParameters"`

	// QUIC 配置参数，当 Name 取值为 QUIC 时，该参数必填。
	// 注意：此字段可能返回 null，表示取不到有效值。
	QUICParameters *QUICParameters `json:"QUICParameters,omitnil,omitempty" name:"QUICParameters"`

	// WebSocket 配置参数，当 Name 取值为 WebSocket 时，该参数必填。
	// 注意：此字段可能返回 null，表示取不到有效值。
	WebSocketParameters *WebSocketParameters `json:"WebSocketParameters,omitnil,omitempty" name:"WebSocketParameters"`

	// Token 鉴权配置参数，当 Name 取值为 Authentication 时，该参数必填。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AuthenticationParameters *AuthenticationParameters `json:"AuthenticationParameters,omitnil,omitempty" name:"AuthenticationParameters"`

	// 浏览器缓存 TTL 配置参数，当 Name 取值为 MaxAge 时，该参数必填。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxAgeParameters *MaxAgeParameters `json:"MaxAgeParameters,omitnil,omitempty" name:"MaxAgeParameters"`

	// 状态码缓存 TTL 配置参数，当 Name 取值为 StatusCodeCache 时，该参数必填。
	// 注意：此字段可能返回 null，表示取不到有效值。
	StatusCodeCacheParameters *StatusCodeCacheParameters `json:"StatusCodeCacheParameters,omitnil,omitempty" name:"StatusCodeCacheParameters"`

	// 离线缓存配置参数，当 Name 取值为 OfflineCache 时，该参数必填。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OfflineCacheParameters *OfflineCacheParameters `json:"OfflineCacheParameters,omitnil,omitempty" name:"OfflineCacheParameters"`

	// 智能加速配置参数，当 Name 取值为 SmartRouting 时，该参数必填。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SmartRoutingParameters *SmartRoutingParameters `json:"SmartRoutingParameters,omitnil,omitempty" name:"SmartRoutingParameters"`

	// 分片回源配置参数，当 Name 取值为 RangeOriginPull 时，该参数必填。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RangeOriginPullParameters *RangeOriginPullParameters `json:"RangeOriginPullParameters,omitnil,omitempty" name:"RangeOriginPullParameters"`

	// HTTP2 回源配置参数，当 Name 取值为 UpstreamHTTP2 时，该参数必填。
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpstreamHTTP2Parameters *UpstreamHTTP2Parameters `json:"UpstreamHTTP2Parameters,omitnil,omitempty" name:"UpstreamHTTP2Parameters"`

	// Host Header 重写配置参数，当 Name 取值为 HostHeader 时，该参数必填。
	// 注意：此字段可能返回 null，表示取不到有效值。
	HostHeaderParameters *HostHeaderParameters `json:"HostHeaderParameters,omitnil,omitempty" name:"HostHeaderParameters"`

	// 访问协议强制 HTTPS 跳转配置，当 Name 取值为 ForceRedirectHTTPS 时，该参数必填。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ForceRedirectHTTPSParameters *ForceRedirectHTTPSParameters `json:"ForceRedirectHTTPSParameters,omitnil,omitempty" name:"ForceRedirectHTTPSParameters"`

	// 回源 HTTPS 配置参数，当 Name 取值为 OriginPullProtocol 时，该参数必填。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OriginPullProtocolParameters *OriginPullProtocolParameters `json:"OriginPullProtocolParameters,omitnil,omitempty" name:"OriginPullProtocolParameters"`

	// 智能压缩配置，当 Name 取值为 Compression 时，该参数必填。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CompressionParameters *CompressionParameters `json:"CompressionParameters,omitnil,omitempty" name:"CompressionParameters"`

	// HSTS 配置参数，当 Name 取值为 HSTS 时，该参数必填。
	// 注意：此字段可能返回 null，表示取不到有效值。
	HSTSParameters *HSTSParameters `json:"HSTSParameters,omitnil,omitempty" name:"HSTSParameters"`

	// 存储客户端请求 IP 的头部信息配置，当 Name 取值为 ClientIPHeader 时，该参数必填。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClientIPHeaderParameters *ClientIPHeaderParameters `json:"ClientIPHeaderParameters,omitnil,omitempty" name:"ClientIPHeaderParameters"`

	// OCSP 装订配置参数，当 Name 取值为 OCSPStapling 时，该参数必填。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OCSPStaplingParameters *OCSPStaplingParameters `json:"OCSPStaplingParameters,omitnil,omitempty" name:"OCSPStaplingParameters"`

	// HTTP2 接入配置参数，当 Name 取值为 HTTP2 时，该参数必填。
	// 注意：此字段可能返回 null，表示取不到有效值。
	HTTP2Parameters *HTTP2Parameters `json:"HTTP2Parameters,omitnil,omitempty" name:"HTTP2Parameters"`

	// POST 请求上传文件流式传输最大限制配置，当 Name 取值为 PostMaxSize 时，该参数必填。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PostMaxSizeParameters *PostMaxSizeParameters `json:"PostMaxSizeParameters,omitnil,omitempty" name:"PostMaxSizeParameters"`

	// 回源时携带客户端 IP 所属地域信息配置参数，当 Name 取值为 ClientIPCountry 时，该参数必填。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClientIPCountryParameters *ClientIPCountryParameters `json:"ClientIPCountryParameters,omitnil,omitempty" name:"ClientIPCountryParameters"`

	// 回源跟随重定向参数配置，当 Name 取值为 UpstreamFollowRedirect 时，该参数必填。
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpstreamFollowRedirectParameters *UpstreamFollowRedirectParameters `json:"UpstreamFollowRedirectParameters,omitnil,omitempty" name:"UpstreamFollowRedirectParameters"`

	// 回源请求参数配置参数，当 Name 取值为 UpstreamRequest 时，该参数必填。
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpstreamRequestParameters *UpstreamRequestParameters `json:"UpstreamRequestParameters,omitnil,omitempty" name:"UpstreamRequestParameters"`

	// SSL/TLS 安全配置参数，当 Name 取值为 TLSConfig 时，该参数必填。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TLSConfigParameters *TLSConfigParameters `json:"TLSConfigParameters,omitnil,omitempty" name:"TLSConfigParameters"`

	// 修改源站配置参数，当 Name 取值为 ModifyOrigin 时，该参数必填。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModifyOriginParameters *ModifyOriginParameters `json:"ModifyOriginParameters,omitnil,omitempty" name:"ModifyOriginParameters"`

	// 七层回源超时配置，当 Name 取值为 HTTPUpstreamTimeout 时，该参数必填。
	// 注意：此字段可能返回 null，表示取不到有效值。
	HTTPUpstreamTimeoutParameters *HTTPUpstreamTimeoutParameters `json:"HTTPUpstreamTimeoutParameters,omitnil,omitempty" name:"HTTPUpstreamTimeoutParameters"`

	// HTTP 应答配置参数，当 Name 取值为 HttpResponse 时，该参数必填。
	// 注意：此字段可能返回 null，表示取不到有效值。
	HttpResponseParameters *HTTPResponseParameters `json:"HttpResponseParameters,omitnil,omitempty" name:"HttpResponseParameters"`

	// 自定义错误页面配置参数，当 Name 取值为 ErrorPage 时，该参数必填。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorPageParameters *ErrorPageParameters `json:"ErrorPageParameters,omitnil,omitempty" name:"ErrorPageParameters"`

	// 修改 HTTP 节点响应头配置参数，当 Name 取值为 ModifyResponseHeader 时，该参数必填。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModifyResponseHeaderParameters *ModifyResponseHeaderParameters `json:"ModifyResponseHeaderParameters,omitnil,omitempty" name:"ModifyResponseHeaderParameters"`

	// 修改 HTTP 节点请求头配置参数，当 Name 取值为 ModifyRequestHeader 时，该参数必填。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModifyRequestHeaderParameters *ModifyRequestHeaderParameters `json:"ModifyRequestHeaderParameters,omitnil,omitempty" name:"ModifyRequestHeaderParameters"`

	// 单连接下载限速配置参数，当 Name 取值为 ResponseSpeedLimit 时，该参数必填。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResponseSpeedLimitParameters *ResponseSpeedLimitParameters `json:"ResponseSpeedLimitParameters,omitnil,omitempty" name:"ResponseSpeedLimitParameters"`

	// 内容标识配置参数，当 Name 取值为 SetContentIdentifier 时，该参数必填。
	// 
	// 注意：此字段可能返回 null，表示取不到有效值。
	SetContentIdentifierParameters *SetContentIdentifierParameters `json:"SetContentIdentifierParameters,omitnil,omitempty" name:"SetContentIdentifierParameters"`

	// Vary 特性配置参数，当 Name 取值为 Vary 时，该参数必填。
	VaryParameters *VaryParameters `json:"VaryParameters,omitnil,omitempty" name:"VaryParameters"`

	// 内容压缩配置参数，当 Name 取值为 ContentCompression 时，该参数必填。该参数为白名单功能，如有需要，请联系腾讯云工程师处理。
	ContentCompressionParameters *ContentCompressionParameters `json:"ContentCompressionParameters,omitnil,omitempty" name:"ContentCompressionParameters"`

	// 回源鉴权配置参数，当 Name 取值为 OriginAuthentication 时，该参数必填。该参数为白名单功能，如有需要，请联系腾讯云工程师处理。
	OriginAuthenticationParameters *OriginAuthenticationParameters `json:"OriginAuthenticationParameters,omitnil,omitempty" name:"OriginAuthenticationParameters"`
}

type RuleEngineItem struct {
	// 规则状态。取值有：<li> enable: 启用； </li><li> disable: 未启用。</li>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 规则 ID。规则的唯一性标识，当调用 ModifyL7AccRules 时，该参数必填。
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// 规则名称。名称长度限制不超过 255 个字符。
	RuleName *string `json:"RuleName,omitnil,omitempty" name:"RuleName"`

	// 规则注释。可以填写多个注释。
	Description []*string `json:"Description,omitnil,omitempty" name:"Description"`

	// 子规则分支。此列表当前只支持填写一项规则，多填无效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Branches []*RuleBranch `json:"Branches,omitnil,omitempty" name:"Branches"`

	// 规则优先级。仅作为出参使用。
	RulePriority *int64 `json:"RulePriority,omitnil,omitempty" name:"RulePriority"`
}

type RuleEngineSubRule struct {
	// 子规则分支
	// 注意：此字段可能返回 null，表示取不到有效值。
	Branches []*RuleBranch `json:"Branches,omitnil,omitempty" name:"Branches"`

	// 规则注释。
	Description []*string `json:"Description,omitnil,omitempty" name:"Description"`
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
	// <li> CUSTOM_STRING：参数值用户自定义，字符串类型。</li>注意：当参数类型为 OBJECT 类型时，请注意参考 [示例2 参数为 OBJECT 类型的创建](https://cloud.tencent.com/document/product/1552/80622#.E7.A4.BA.E4.BE.8B2-.E4.BF.AE.E6.94.B9.E6.BA.90.E7.AB.99.E4.B8.BAIP.E5.9F.9F.E5.90.8D)
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

type SearchEngineBots struct {
	// 来自搜索引擎爬虫的请求的执行动作。 SecurityAction 的 Name 取值支持：<li>Deny：拦截；</li><li>Monitor：观察；</li><li>Disabled：未启用，不启用指定规则；</li><li>Challenge：挑战，其中 ChallengeActionParameters 中的 ChallengeOption 支持 JSChallenge 和 ManagedChallenge；</li><li>Allow：放行（待废弃）。</li> 
	BaseAction *SecurityAction `json:"BaseAction,omitnil,omitempty" name:"BaseAction"`

	// 指定搜索引擎爬虫请求的处置方式。
	BotManagementActionOverrides []*BotManagementActionOverrides `json:"BotManagementActionOverrides,omitnil,omitempty" name:"BotManagementActionOverrides"`
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

type SecurityAction struct {
	// 安全执行的具体动作。取值有：
	// <li>Deny：拦截，阻止请求访问站点资源；</li>
	// <li>Monitor：观察，仅记录日志；</li>
	// <li>Redirect：重定向至 URL；</li>
	// <li>Disabled：未启用，不启用指定规则；</li>
	// <li>Allow：允许访问，但延迟处理请求；</li>
	// <li>Challenge：挑战，响应挑战内容；</li>
	// <li>Trans：放行，允许请求直接访问站点资源；</li>
	// <li>BlockIP：待废弃，IP 封禁；</li>
	// <li>ReturnCustomPage：待废弃，使用指定页面拦截；</li>
	// <li>JSChallenge：待废弃，JavaScript 挑战；</li>
	// <li>ManagedChallenge：待废弃，托管挑战。</li>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 当 Name 为 Deny 时的附加参数。
	DenyActionParameters *DenyActionParameters `json:"DenyActionParameters,omitnil,omitempty" name:"DenyActionParameters"`

	// 当 Name 为 Redirect 时的附加参数。
	RedirectActionParameters *RedirectActionParameters `json:"RedirectActionParameters,omitnil,omitempty" name:"RedirectActionParameters"`

	// 当 Name 为 Allow 时的附加参数。
	AllowActionParameters *AllowActionParameters `json:"AllowActionParameters,omitnil,omitempty" name:"AllowActionParameters"`

	// 当 Name 为 Challenge 时的附加参数。
	ChallengeActionParameters *ChallengeActionParameters `json:"ChallengeActionParameters,omitnil,omitempty" name:"ChallengeActionParameters"`

	// 待废弃，当 Name 为 BlockIP 时的附加参数。
	BlockIPActionParameters *BlockIPActionParameters `json:"BlockIPActionParameters,omitnil,omitempty" name:"BlockIPActionParameters"`

	// 待废弃，当 Name 为 ReturnCustomPage 时的附加参数。
	ReturnCustomPageActionParameters *ReturnCustomPageActionParameters `json:"ReturnCustomPageActionParameters,omitnil,omitempty" name:"ReturnCustomPageActionParameters"`
}

type SecurityConfig struct {
	// 托管规则。如果入参为空或不填，默认使用历史配置。
	WafConfig *WafConfig `json:"WafConfig,omitnil,omitempty" name:"WafConfig"`

	// 速率限制。如果入参为空或不填，默认使用历史配置。
	RateLimitConfig *RateLimitConfig `json:"RateLimitConfig,omitnil,omitempty" name:"RateLimitConfig"`

	// 自定义规则。如果入参为空或不填，默认使用历史配置。
	AclConfig *AclConfig `json:"AclConfig,omitnil,omitempty" name:"AclConfig"`

	// Bot配置。如果入参为空或不填，默认使用历史配置。
	BotConfig *BotConfig `json:"BotConfig,omitnil,omitempty" name:"BotConfig"`

	// 七层防护总开关。如果入参为空或不填，默认使用历史配置。
	SwitchConfig *SwitchConfig `json:"SwitchConfig,omitnil,omitempty" name:"SwitchConfig"`

	// 基础访问管控。如果入参为空或不填，默认使用历史配置。
	IpTableConfig *IpTableConfig `json:"IpTableConfig,omitnil,omitempty" name:"IpTableConfig"`

	// 例外规则配置。如果入参为空或不填，默认使用历史配置。
	ExceptConfig *ExceptConfig `json:"ExceptConfig,omitnil,omitempty" name:"ExceptConfig"`

	// 自定义拦截页面配置。如果入参为空或不填，默认使用历史配置。
	DropPageConfig *DropPageConfig `json:"DropPageConfig,omitnil,omitempty" name:"DropPageConfig"`

	// 模板配置。此处仅出参数使用。
	TemplateConfig *TemplateConfig `json:"TemplateConfig,omitnil,omitempty" name:"TemplateConfig"`

	// 慢速攻击配置。如果入参为空或不填，默认使用历史配置。
	SlowPostConfig *SlowPostConfig `json:"SlowPostConfig,omitnil,omitempty" name:"SlowPostConfig"`

	// 检测长度限制配置。仅出参使用。
	DetectLengthLimitConfig *DetectLengthLimitConfig `json:"DetectLengthLimitConfig,omitnil,omitempty" name:"DetectLengthLimitConfig"`
}

type SecurityPolicy struct {
	// 自定义规则配置。
	CustomRules *CustomRules `json:"CustomRules,omitnil,omitempty" name:"CustomRules"`

	// 托管规则配置。
	ManagedRules *ManagedRules `json:"ManagedRules,omitnil,omitempty" name:"ManagedRules"`

	// HTTP DDOS 防护配置。
	HttpDDoSProtection *HttpDDoSProtection `json:"HttpDDoSProtection,omitnil,omitempty" name:"HttpDDoSProtection"`

	// 速率限制规则配置。
	RateLimitingRules *RateLimitingRules `json:"RateLimitingRules,omitnil,omitempty" name:"RateLimitingRules"`

	// 例外规则配置。
	ExceptionRules *ExceptionRules `json:"ExceptionRules,omitnil,omitempty" name:"ExceptionRules"`

	// Bot 管理配置。
	BotManagement *BotManagement `json:"BotManagement,omitnil,omitempty" name:"BotManagement"`
}

type SecurityPolicyTemplateInfo struct {
	// 策略模板所属的站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 策略模板 ID。
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// 策略模板名称。
	TemplateName *string `json:"TemplateName,omitnil,omitempty" name:"TemplateName"`

	// 策略模板绑定的域名信息。
	BindDomains []*BindDomainInfo `json:"BindDomains,omitnil,omitempty" name:"BindDomains"`
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

type SecurityWeightedAction struct {
	// Bot 自定义规则的处置方式。取值有：<li>Allow：放行，其中 AllowActionParameters 支持 MinDelayTime 和 MaxDelayTime 配置；</li><li>Deny：拦截，其中 DenyActionParameters 中支持 BlockIp、ReturnCustomPage 和 Stall 配置；</li><li>Monitor：观察；</li><li>Challenge：挑战，其中 ChallengeActionParameters.ChallengeOption 支持 JSChallenge 和 ManagedChallenge；</li><li>Redirect：重定向至URL。</li>
	SecurityAction *SecurityAction `json:"SecurityAction,omitnil,omitempty" name:"SecurityAction"`

	// 当前 SecurityAction 的权重，仅支持 10 ~ 100 且必须为 10 的倍数，其中 Weight 参数全部相加须等于 100。
	Weight *int64 `json:"Weight,omitnil,omitempty" name:"Weight"`
}

type ServerCertInfo struct {
	// 服务器证书 ID。来源于 SSL 侧，您可以前往 [SSL 证书列表](https://console.cloud.tencent.com/ssl) 查看 CertId。
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

	// 证书归属域名名称。
	CommonName *string `json:"CommonName,omitnil,omitempty" name:"CommonName"`
}

type SessionRateControl struct {
	// 会话速率和周期特征校验配置是否开启。取值有：<li>on：启用</li><li>off：关闭</li>
	Enabled *string `json:"Enabled,omitnil,omitempty" name:"Enabled"`

	// 会话速率和周期特征校验高风险的执行动作。 SecurityAction 的 Name 取值支持：<li>Deny：拦截，其中 DenyActionParameters 中支持 Stall 配置；</li><li>Monitor：观察；</li><li>Allow：等待后响应，其中 AllowActionParameters 需要 MinDelayTime 和 MaxDelayTime 配置。</li>
	HighRateSessionAction *SecurityAction `json:"HighRateSessionAction,omitnil,omitempty" name:"HighRateSessionAction"`

	// 会话速率和周期特征校验中风险的执行动作。 SecurityAction 的 Name 取值支持：<li>Deny：拦截，其中 DenyActionParameters 中支持 Stall 配置；</li><li>Monitor：观察；</li><li>Allow：等待后响应，其中 AllowActionParameters 需要 MinDelayTime 和 MaxDelayTime 配置。</li>
	MidRateSessionAction *SecurityAction `json:"MidRateSessionAction,omitnil,omitempty" name:"MidRateSessionAction"`

	// 会话速率和周期特征校验低风险的执行动作。 SecurityAction 的 Name 取值支持：<li>Deny：拦截，其中 DenyActionParameters 中支持 Stall 配置；</li><li>Monitor：观察；</li><li>Allow：等待后响应，其中 AllowActionParameters 需要 MinDelayTime 和 MaxDelayTime 配置。</li>
	LowRateSessionAction *SecurityAction `json:"LowRateSessionAction,omitnil,omitempty" name:"LowRateSessionAction"`
}

type SetContentIdentifierParameters struct {
	// 内容标识id
	ContentIdentifier *string `json:"ContentIdentifier,omitnil,omitempty" name:"ContentIdentifier"`
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
	MatchFrom []*string `json:"MatchFrom,omitnil,omitempty" name:"MatchFrom"`

	// 匹配Content所使用的匹配方式，取值为：
	// <li>equal：精准匹配，等于；</li>
	// <li>wildcard：通配符匹配，支持 * 通配。</li>
	MatchContentType *string `json:"MatchContentType,omitnil,omitempty" name:"MatchContentType"`

	// 匹配Value的值。
	MatchContent []*string `json:"MatchContent,omitnil,omitempty" name:"MatchContent"`
}

type SlowAttackDefense struct {
	// 慢速攻击防护是否开启。取值有：<li>on：开启；</li><li>off：关闭。</li>
	Enabled *string `json:"Enabled,omitnil,omitempty" name:"Enabled"`

	// 慢速攻击防护的处置方式，当 Enabled 为 on 时，此字段必填。SecurityAction 的 Name 取值支持：<li>Monitor：观察；</li><li>Deny：拦截；</li>
	Action *SecurityAction `json:"Action,omitnil,omitempty" name:"Action"`

	// 正文传输最小速率阈值的具体配置，当 Enabled 为 on 时，此字段必填。
	MinimalRequestBodyTransferRate *MinimalRequestBodyTransferRate `json:"MinimalRequestBodyTransferRate,omitnil,omitempty" name:"MinimalRequestBodyTransferRate"`

	// 正文传输超时时长的具体配置，当 Enabled 为 on 时，此字段必填。
	RequestBodyTransferTimeout *RequestBodyTransferTimeout `json:"RequestBodyTransferTimeout,omitnil,omitempty" name:"RequestBodyTransferTimeout"`
}

type SlowPostConfig struct {
	// 开关，取值有：
	// <li>on：开启；</li>
	// <li>off：关闭。</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// 首包配置。
	FirstPartConfig *FirstPartConfig `json:"FirstPartConfig,omitnil,omitempty" name:"FirstPartConfig"`

	// 基础配置。
	SlowRateConfig *SlowRateConfig `json:"SlowRateConfig,omitnil,omitempty" name:"SlowRateConfig"`

	// 慢速攻击的处置动作，取值有：
	// <li>monitor：观察；</li>
	// <li>drop：拦截。</li>
	Action *string `json:"Action,omitnil,omitempty" name:"Action"`

	// 本规则的Id。
	RuleId *uint64 `json:"RuleId,omitnil,omitempty" name:"RuleId"`
}

type SlowRateConfig struct {
	// 开关，取值有：
	// <li>on：开启；</li>
	// <li>off：关闭。</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// 统计的间隔，单位是秒，即在首段包传输结束后，将数据传输轴按照本参数切分，每个分片独立计算慢速攻击。
	Interval *uint64 `json:"Interval,omitnil,omitempty" name:"Interval"`

	// 统计时应用的速率阈值，单位是bps，即如果本分片中的传输速率没达到本参数的值，则判定为慢速攻击，应用慢速攻击的处置方式。
	Threshold *uint64 `json:"Threshold,omitnil,omitempty" name:"Threshold"`
}

type SmartRouting struct {
	// 智能加速配置开关，取值有：
	// <li>on：开启；</li>
	// <li>off：关闭。</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`
}

type SmartRoutingParameters struct {
	// 智能加速配置开关，取值有：
	// <li>on：开启；</li>
	// <li>off：关闭。</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`
}

type SourceIDC struct {
	// 来自指定 IDC 请求的处置方式。 SecurityAction 的 Name 取值支持：<li>Deny：拦截；</li><li>Monitor：观察；</li><li>Disabled：未启用，不启用指定规则；</li><li>Challenge：挑战，其中 ChallengeActionParameters 中的 ChallengeOption 支持 JSChallenge 和 ManagedChallenge；</li><li>Allow：放行（待废弃）。</li>
	BaseAction *SecurityAction `json:"BaseAction,omitnil,omitempty" name:"BaseAction"`

	// 指定 IDC 请求的处置方式。
	BotManagementActionOverrides []*BotManagementActionOverrides `json:"BotManagementActionOverrides,omitnil,omitempty" name:"BotManagementActionOverrides"`
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

type StandardDebugParameters struct {
	// Debug 功能开关，取值有：
	// <li>on：开启；</li>
	// <li>off：关闭。</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// 允许的客户端来源。支持填写 IPv4 以及 IPv6 的 IP 网段。0.0.0.0/0 表示允许所有 IPv4 客户端进行调试；::/0 表示允许所有 IPv6 客户端进行调试；不能填写 127.0.0.1。<br>注意：当 Switch 字段为 on 时，此字段必填，且填写个数为 1～100；当 Switch 为 off 时，无需填写此字段，若填写则不生效。
	AllowClientIPList []*string `json:"AllowClientIPList,omitnil,omitempty" name:"AllowClientIPList"`

	// Debug 功能到期时间。超出设置的时间，则功能失效。<br>注意：当 Switch 为 on 时，此字段必填；当 Switch 为 off 时，无需填写此字段，若填写则不生效。
	Expires *string `json:"Expires,omitnil,omitempty" name:"Expires"`
}

type StatusCodeCacheParam struct {
	// 状态码，取值为 400、 401、403、 404、 405、 407、 414、 500、 501、 502、 503、 504、 509、 514 之一。
	StatusCode *int64 `json:"StatusCode,omitnil,omitempty" name:"StatusCode"`

	// 缓存时间数值，单位为秒，取值：0～31536000。
	CacheTime *int64 `json:"CacheTime,omitnil,omitempty" name:"CacheTime"`
}

type StatusCodeCacheParameters struct {
	// 状态码缓存 TTL 。
	// 注意：此字段可能返回 null，表示取不到有效值。
	StatusCodeCacheParams []*StatusCodeCacheParam `json:"StatusCodeCacheParams,omitnil,omitempty" name:"StatusCodeCacheParams"`
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

type TCCaptchaOption struct {
	// CaptchaAppId 信息。
	CaptchaAppId *string `json:"CaptchaAppId,omitnil,omitempty" name:"CaptchaAppId"`

	// AppSecretKey 信息。
	AppSecretKey *string `json:"AppSecretKey,omitnil,omitempty" name:"AppSecretKey"`
}

type TCRCEOption struct {
	// Channel 信息。
	Channel *string `json:"Channel,omitnil,omitempty" name:"Channel"`
}

type TLSConfigParameters struct {
	// TLS 版本。至少填写一个，如果是多个时，需要为连续版本号，例如：开启 TLS1、1.1、1.2 和 1.3，不可仅开启 1 和 1.2 而关闭 1.1。取值有：<li>TLSv1：TLSv1 版本；</li><li>TLSv1.1：TLSv1.1 版本；</li><li>TLSv1.2：TLSv1.2 版本；</li><li>TLSv1.3：TLSv1.3 版本。</li>
	Version []*string `json:"Version,omitnil,omitempty" name:"Version"`

	// 密码套件。详细介绍请参考 [TLS 版本及密码套件说明](https://cloud.tencent.com/document/product/1552/86545)。取值有：<li>loose-v2023：loose-v2023 密码套件；</li><li>general-v2023：general-v2023 密码套件；</li><li>strict-v2023：strict-v2023 密码套件。</li>
	CipherSuite *string `json:"CipherSuite,omitnil,omitempty" name:"CipherSuite"`
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
	Method *string `json:"Method,omitnil,omitempty" name:"Method"`

	// 状态。取值有：
	// <li>processing：处理中；</li>
	// <li>success：成功；</li>
	// <li>failed：失败；</li>
	// <li>timeout：超时；</li>
	// <li>canceled：已取消。</li>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 任务创建时间。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 任务完成时间。
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 刷新、预热失败类型。取值有：
	// <li>taskFailed：任务失败；</li>
	// <li>quotaExceeded：配额超限；</li>
	// <li>downloadManifestFailed：下载描述文件失败；</li>
	// <li>accessDenied：访问被拒绝。</li>
	// <li>originPullFailed：回源失败。</li>
	FailType *string `json:"FailType,omitnil,omitempty" name:"FailType"`

	// 清除缓存、预热缓存的失败原因描述。
	FailMessage *string `json:"FailMessage,omitnil,omitempty" name:"FailMessage"`
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

type URLPath struct {
	// 执行动作，取值有：
	// <li>follow：跟随请求；</li>
	// <li>custom：自定义；</li>
	// <li>regex：正则匹配。</li>
	Action *string `json:"Action,omitnil,omitempty" name:"Action"`

	// 正则匹配的表达式，长度范围为 1～1024。<br>注意：当 Action 为 regex 时，此字段必填；当 Action 为 follow 或 custom 时，无需填写此字段，若填写则不生效。
	Regex *string `json:"Regex,omitnil,omitempty" name:"Regex"`

	// 重定向的目标URL，长度范围为 1～1024。<br>注意：当 Action 为 regex 或 custom 时，此字段必填；当 Action 为 follow 时，无需填写此字段，若填写则不生效。
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
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

type UpstreamCertInfo struct {
	// 在回源双向认证场景下，该字段为 EO 节点回源时携带的证书（包含公钥、私钥即可），部署在 EO 节点，用于源站对 EO 节点进行认证。在作为入参使用时，不填写表示保持原有配置。
	UpstreamMutualTLS *MutualTLS `json:"UpstreamMutualTLS,omitnil,omitempty" name:"UpstreamMutualTLS"`

	// 在源站证书校验场景下，该字段为 EO 节点回源时用于验证的 CA 证书，部署在 EO 节点，用于 EO 节点对服务端证书进行认证。在作为入参使用时，不填写表示保持原有配置。
	UpstreamCertificateVerify *OriginCertificateVerify `json:"UpstreamCertificateVerify,omitnil,omitempty" name:"UpstreamCertificateVerify"`
}

type UpstreamFollowRedirectParameters struct {
	// 回源跟随重定向配置开关，取值有：
	// <li>on：开启；</li>
	// <li>off：关闭。</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// 最大重定向次数。取值为 1-5。
	// 注意：当 Switch 为 on 时，此字段必填；当 Switch 为 off 时，无需填写此字段，若填写则不生效。
	MaxTimes *int64 `json:"MaxTimes,omitnil,omitempty" name:"MaxTimes"`
}

type UpstreamHTTP2Parameters struct {
	// HTTP2 回源配置开关，取值有：
	// <li>on：开启；</li>
	// <li>off：关闭。</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`
}

type UpstreamHttp2 struct {
	// http2 回源配置开关，取值有：
	// <li>on：开启；</li>
	// <li>off：关闭。</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`
}

type UpstreamRequestCookie struct {
	// 回源请求参数 Cookie 配置开关，取值有：
	// <li>on：开启；</li>
	// <li>off：关闭。</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// 回源请求参数 Cookie 模式。当 Switch 为 on 时，该参数必填。取值有：
	// <li>full：表示全部保留；</li>
	// <li>ignore：表示全部忽略；</li>
	// <li>includeCustom：表示保留部分参数；</li>
	// <li>excludeCustom：表示忽略部分参数。</li>
	Action *string `json:"Action,omitnil,omitempty" name:"Action"`

	// 指定参数值。仅当查询字符串模式 Action 为 includeCustom 或者 excludeCustom 时该参数生效，用于指定需要保留或者忽略的参数。最大支持 10 个参数。
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`
}

type UpstreamRequestParameters struct {
	// 查询字符串配置。可选配置项，不填表示不配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	QueryString *UpstreamRequestQueryString `json:"QueryString,omitnil,omitempty" name:"QueryString"`

	// Cookie 配置。可选配置项，不填表示不配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Cookie *UpstreamRequestCookie `json:"Cookie,omitnil,omitempty" name:"Cookie"`
}

type UpstreamRequestQueryString struct {
	// 回源请求参数查询字符串配置开关，取值有：
	// <li>on：开启；</li>
	// <li>off：关闭。</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// 查询字符串模式。当 Switch 为 on 时，该参数必填。取值有：
	// <li>full：全部保留；</li>
	// <li>ignore：全部忽略；</li>
	// <li>includeCustom：保留部分参数；</li>
	// <li>excludeCustom：忽略部分参数。</li>
	Action *string `json:"Action,omitnil,omitempty" name:"Action"`

	// 指定参数值。仅当查询字符串模式 Action 为 includeCustom 或者 excludeCustom 时该参数生效，用于指定需要保留或者忽略的参数。最大支持 10 个参数。
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`
}

type UpstreamURLRewriteParameters struct {
	// 回源 URL 重写类型。仅支持填写 Path。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 回源 URL 重写动作。取值有：
	// <li><b>replace</b>：指替换完整路径。用于将完整的请求 URL Path 替换为指定路径。
	// </li>
	// <li><b>addPrefix</b>：指增加路径前缀。用于增加指定路径前缀至请求 URL Path。
	// </li>
	// <li><b>rmvPrefix</b>：指移除路径前缀。用于移除请求 URL Path 的指定路径前缀。
	// </li>
	// <li><b>regexReplace</b>：指正则替换完整路径。用于通过 Google RE2 正则表达式匹配和替换完整路径。
	// </li>
	Action *string `json:"Action,omitnil,omitempty" name:"Action"`

	// 回源 URL 重写值。需要满足 URL Path 规范，且保证重写后的 Path 以 / 开头，以防止回源 URL 的 Host 被修改，长度范围为 1～1024。当 Action 为 addPrefix 时，不能以 / 结尾；当 Action 为 rmvPrefix 时，不能存在 *；当 Action 为 regexReplace 时，支持用 $NUM 引用正则捕获组，其中 NUM 代表组编号，如 $1，最多支持 $9。
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`

	// 回源 URL 重写用于正则替换匹配完整路径的正则表达式。需要满足 Google RE2 规范，长度范围为 1～1024。当 Action 为 regexReplace 时，此字段必填，否则无需填写此字段。
	Regex *string `json:"Regex,omitnil,omitempty" name:"Regex"`
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

type VaryParameters struct {
	// Vary 特性配置开关，取值有：
	// <li>on：开启；</li>
	// <li>off：关闭。</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`
}

// Predefined struct for user
type VerifyOwnershipRequestParams struct {
	// 站点域名或者站点下的加速域名。
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`
}

type VerifyOwnershipRequest struct {
	*tchttp.BaseRequest
	
	// 站点域名或者站点下的加速域名。
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

type VideoTemplateInfo struct {
	// 视频流的编码格式，可选值：<li>H.264: 使用 H.264 编码；</li><li>H.265: 使用 H.265 编码。</li>
	Codec *string `json:"Codec,omitnil,omitempty" name:"Codec"`

	// 视频帧率，取值范围：[0, 30]，单位：Hz。
	// 取值为 0，表示帧率和原始视频保持一致，但最大不超过 30。
	// 默认值：0。
	Fps *float64 `json:"Fps,omitnil,omitempty" name:"Fps"`

	// 视频流的码率，取值范围：0 和 [128, 10000]，单位：kbps。
	// 取值为 0，表示自动根据视频画面和质量选择视频码率。
	// 默认值：0。
	Bitrate *uint64 `json:"Bitrate,omitnil,omitempty" name:"Bitrate"`

	// 分辨率自适应，可选值：<li>open：开启，此时，Width 代表视频的长边，Height 表示视频的短边；</li><li>close：关闭，此时，Width 代表视频的宽度，Height 表示视频的高度。</li>默认值：open。
	ResolutionAdaptive *string `json:"ResolutionAdaptive,omitnil,omitempty" name:"ResolutionAdaptive"`

	// 视频流宽度（或长边）的最大值，取值范围：0 和 [128, 1920]，单位：px。<li>当 Width、Height 均为 0，则分辨率同源；</li><li>当 Width 为 0，Height 非 0，则 Width 按比例缩放；</li><li>当 Width 非 0，Height 为 0，则 Height 按比例缩放；</li><li>当 Width、Height 均非 0，则分辨率按用户指定。</li>默认值：0。
	Width *uint64 `json:"Width,omitnil,omitempty" name:"Width"`

	// 视频流高度（或短边）的最大值，取值范围：0 和 [128, 1080]，单位：px。<li>当 Width、Height 均为 0，则分辨率同源；</li><li>当 Width 为 0，Height 非 0，则 Width 按比例缩放；</li><li>当 Width 非 0，Height 为 0，则 Height 按比例缩放；</li><li>当 Width、Height 均非 0，则分辨率按用户指定。</li>默认值：0。
	Height *uint64 `json:"Height,omitnil,omitempty" name:"Height"`

	// 填充方式，当视频流配置宽高参数与原始视频的宽高比不一致时，对转码的处理方式，即为“填充”。可选填充方式：<li> stretch：拉伸，对每一帧进行拉伸，填满整个画面，可能导致转码后的视频被“压扁”或者“拉长”。</li><li>black：留黑，保持视频宽高比不变，边缘剩余部分使用黑色填充。</li><li>white：留白，保持视频宽高比不变，边缘剩余部分使用白色填充。</li><li>gauss：高斯模糊，保持视频宽高比不变，边缘剩余部分使用高斯模糊填充。</li>默认值：black 。
	FillType *string `json:"FillType,omitnil,omitempty" name:"FillType"`
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

type WebSocketParameters struct {
	// WebSocket 超时时间配置开关，取值有：
	// <li>on：使用 Timeout 作为 WebSocket 超时时间；</li>
	// <li>off：平台仍支持 WebSocket 连接，此时使用系统默认的 15 秒为超时时间。</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// 超时时间，单位为秒，最大超时时间 120 秒。<br>注意：当 Switch 为 on 时，此字段必填，否则此字段不生效。
	Timeout *int64 `json:"Timeout,omitnil,omitempty" name:"Timeout"`
}

type Zone struct {
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 站点名称。
	ZoneName *string `json:"ZoneName,omitnil,omitempty" name:"ZoneName"`

	// 同名站点标识。允许输入数字、英文、"." 、"-" 和 "_" 组合，长度 200 个字符以内。
	AliasZoneName *string `json:"AliasZoneName,omitnil,omitempty" name:"AliasZoneName"`

	// 站点加速区域，取值有：
	// <li> global：全球可用区；</li>
	// <li> mainland：中国大陆可用区；</li>
	// <li> overseas：全球可用区（不含中国大陆）。</li>
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 站点接入类型，取值有：
	// <li> full：NS 接入类型；</li>
	// <li> partial：CNAME 接入类型；</li>
	// <li> noDomainAccess：无域名接入类型；</li>
	// <li>dnsPodAccess：DNSPod 托管类型，该类型要求您的域名已托管在腾讯云 DNSPod；</li>
	// <li> pages：Pages 类型。</li>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 站点关联的标签。
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 计费资源列表。
	Resources []*Resource `json:"Resources,omitnil,omitempty" name:"Resources"`

	// NS 类型站点详情。仅当 Type = full 时返回值。
	NSDetail *NSDetail `json:"NSDetail,omitnil,omitempty" name:"NSDetail"`

	// CNAME 类型站点详情。仅当 Type = partial 时返回值。
	CNAMEDetail *CNAMEDetail `json:"CNAMEDetail,omitnil,omitempty" name:"CNAMEDetail"`

	// DNSPod 托管类型站点详情。仅当 Type = dnsPodAccess 时返回值。
	DNSPodDetail *DNSPodDetail `json:"DNSPodDetail,omitnil,omitempty" name:"DNSPodDetail"`

	// 站点创建时间。
	CreatedOn *string `json:"CreatedOn,omitnil,omitempty" name:"CreatedOn"`

	// 站点修改时间。
	ModifiedOn *string `json:"ModifiedOn,omitnil,omitempty" name:"ModifiedOn"`

	// 站点状态，取值有：
	// <li> active：NS 已切换； </li>
	// <li> pending：NS 未切换；</li>
	// <li> moved：NS 已切走；</li>
	// <li> deactivated：被封禁。 </li>
	// <li> initializing：待绑定套餐。 </li>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// CNAME 接入状态，取值有：
	// <li> finished：站点已验证；</li>
	// <li> pending：站点验证中。</li>
	CnameStatus *string `json:"CnameStatus,omitnil,omitempty" name:"CnameStatus"`

	// 展示状态，取值有：
	// <li> active：已启用；</li>
	// <li> inactive：未生效；</li>
	// <li> paused：已停用。</li>
	ActiveStatus *string `json:"ActiveStatus,omitnil,omitempty" name:"ActiveStatus"`

	// 锁定状态，取值有：<li> enable：正常，允许进行修改操作；</li><li> disable：锁定中，不允许进行修改操作；</li><li> plan_migrate：套餐迁移中，不允许进行修改操作。</li>
	LockStatus *string `json:"LockStatus,omitnil,omitempty" name:"LockStatus"`

	// 站点是否关闭。
	Paused *bool `json:"Paused,omitnil,omitempty" name:"Paused"`

	// 是否伪站点（该字段为历史保留字段，已不再维护，请根据站点类型参考对应字段），取值有：
	// <li> 0：非伪站点；</li>
	// <li> 1：伪站点。</li>
	IsFake *int64 `json:"IsFake,omitnil,omitempty" name:"IsFake"`

	// 是否开启 CNAME 加速（该字段为历史保留字段，已不再维护，请根据站点类型参考对应字段），取值有：
	// <li> enabled：开启；</li>
	// <li> disabled：关闭。</li>
	CnameSpeedUp *string `json:"CnameSpeedUp,omitnil,omitempty" name:"CnameSpeedUp"`

	// 归属权验证信息。（该字段为历史保留字段，已不再维护，请根据站点类型参考对应字段）
	// 注意：此字段可能返回 null，表示取不到有效值。
	OwnershipVerification *OwnershipVerification `json:"OwnershipVerification,omitnil,omitempty" name:"OwnershipVerification"`

	// 站点当前使用的 NS 列表。（该字段为历史保留字段，已不再维护，请根据站点类型参考对应字段）
	OriginalNameServers []*string `json:"OriginalNameServers,omitnil,omitempty" name:"OriginalNameServers"`

	// 腾讯云分配的 NS 列表。（该字段为历史保留字段，已不再维护，请根据站点类型参考对应字段）
	NameServers []*string `json:"NameServers,omitnil,omitempty" name:"NameServers"`

	// 用户自定义 NS 信息。（该字段为历史保留字段，已不再维护，请根据站点类型参考对应字段）
	// 注意：此字段可能返回 null，表示取不到有效值。
	VanityNameServers *VanityNameServers `json:"VanityNameServers,omitnil,omitempty" name:"VanityNameServers"`

	// 用户自定义 NS IP 信息。（该字段为历史保留字段，已不再维护，请根据站点类型参考对应字段）
	// 注意：此字段可能返回 null，表示取不到有效值。
	VanityNameServersIps []*VanityNameServersIps `json:"VanityNameServersIps,omitnil,omitempty" name:"VanityNameServersIps"`
}

type ZoneConfig struct {
	// 智能加速配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SmartRouting *SmartRoutingParameters `json:"SmartRouting,omitnil,omitempty" name:"SmartRouting"`

	// 缓存过期时间配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Cache *CacheConfigParameters `json:"Cache,omitnil,omitempty" name:"Cache"`

	// 浏览器缓存配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxAge *MaxAgeParameters `json:"MaxAge,omitnil,omitempty" name:"MaxAge"`

	// 节点缓存键配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CacheKey *CacheKeyConfigParameters `json:"CacheKey,omitnil,omitempty" name:"CacheKey"`

	// 缓存预刷新配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CachePrefresh *CachePrefreshParameters `json:"CachePrefresh,omitnil,omitempty" name:"CachePrefresh"`

	// 离线缓存配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OfflineCache *OfflineCacheParameters `json:"OfflineCache,omitnil,omitempty" name:"OfflineCache"`

	// 智能压缩配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Compression *CompressionParameters `json:"Compression,omitnil,omitempty" name:"Compression"`

	// 访问协议强制 HTTPS 跳转配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ForceRedirectHTTPS *ForceRedirectHTTPSParameters `json:"ForceRedirectHTTPS,omitnil,omitempty" name:"ForceRedirectHTTPS"`

	// HSTS 相关配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	HSTS *HSTSParameters `json:"HSTS,omitnil,omitempty" name:"HSTS"`

	// TLS 相关配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TLSConfig *TLSConfigParameters `json:"TLSConfig,omitnil,omitempty" name:"TLSConfig"`

	// OCSP 装订配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OCSPStapling *OCSPStaplingParameters `json:"OCSPStapling,omitnil,omitempty" name:"OCSPStapling"`

	// HTTP2 相关配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	HTTP2 *HTTP2Parameters `json:"HTTP2,omitnil,omitempty" name:"HTTP2"`

	// QUIC 访问配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	QUIC *QUICParameters `json:"QUIC,omitnil,omitempty" name:"QUIC"`

	// HTTP2 回源配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpstreamHTTP2 *UpstreamHTTP2Parameters `json:"UpstreamHTTP2,omitnil,omitempty" name:"UpstreamHTTP2"`

	// IPv6 访问配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	IPv6 *IPv6Parameters `json:"IPv6,omitnil,omitempty" name:"IPv6"`

	// WebSocket 配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	WebSocket *WebSocketParameters `json:"WebSocket,omitnil,omitempty" name:"WebSocket"`

	// POST 请求传输配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PostMaxSize *PostMaxSizeParameters `json:"PostMaxSize,omitnil,omitempty" name:"PostMaxSize"`

	// 客户端 IP 回源请求头配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClientIPHeader *ClientIPHeaderParameters `json:"ClientIPHeader,omitnil,omitempty" name:"ClientIPHeader"`

	// 回源时是否携带客户端 IP 所属地域信息的配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClientIPCountry *ClientIPCountryParameters `json:"ClientIPCountry,omitnil,omitempty" name:"ClientIPCountry"`

	// gRPC 协议支持配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Grpc *GrpcParameters `json:"Grpc,omitnil,omitempty" name:"Grpc"`

	// 网络错误日志记录配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	NetworkErrorLogging *NetworkErrorLoggingParameters `json:"NetworkErrorLogging,omitnil,omitempty" name:"NetworkErrorLogging"`

	// 中国大陆加速优化配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AccelerateMainland *AccelerateMainlandParameters `json:"AccelerateMainland,omitnil,omitempty" name:"AccelerateMainland"`

	// 标准 Debug 配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	StandardDebug *StandardDebugParameters `json:"StandardDebug,omitnil,omitempty" name:"StandardDebug"`
}

type ZoneConfigParameters struct {
	// 站点名称。
	ZoneName *string `json:"ZoneName,omitnil,omitempty" name:"ZoneName"`

	// 站点配置信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ZoneConfig *ZoneConfig `json:"ZoneConfig,omitnil,omitempty" name:"ZoneConfig"`
}

type ZoneInfo struct {
	// 站点id。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 站点名称。
	ZoneName *string `json:"ZoneName,omitnil,omitempty" name:"ZoneName"`

	// 站点是否停用。取值有：<li>false：非停用；</li>
	// <li>true：停用。</li>
	Paused *bool `json:"Paused,omitnil,omitempty" name:"Paused"`
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

	// 网络错误日志记录配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	NetworkErrorLogging *NetworkErrorLogging `json:"NetworkErrorLogging,omitnil,omitempty" name:"NetworkErrorLogging"`

	// 图片优化相关配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ImageOptimize *ImageOptimize `json:"ImageOptimize,omitnil,omitempty" name:"ImageOptimize"`

	// 中国大陆加速优化配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AccelerateMainland *AccelerateMainland `json:"AccelerateMainland,omitnil,omitempty" name:"AccelerateMainland"`

	// 标准 Debug 配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	StandardDebug *StandardDebug `json:"StandardDebug,omitnil,omitempty" name:"StandardDebug"`

	// 视频即时处理配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	JITVideoProcess *JITVideoProcess `json:"JITVideoProcess,omitnil,omitempty" name:"JITVideoProcess"`
}