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

const (
	// 此产品的特有错误码

	// CAM签名/鉴权错误。
	AUTHFAILURE = "AuthFailure"

	// DryRun 操作，代表请求将会是成功的，只是多传了 DryRun 参数。
	DRYRUNOPERATION = "DryRunOperation"

	// 操作失败。
	FAILEDOPERATION = "FailedOperation"

	// 证书不存在。
	FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"

	// 站点状态不正确。
	FAILEDOPERATION_INVALIDZONESTATUS = "FailedOperation.InvalidZoneStatus"

	// 内部错误。
	INTERNALERROR = "InternalError"

	// 后台处理出错。
	INTERNALERROR_BACKENDERROR = "InternalError.BackendError"

	// 数据库错误。
	INTERNALERROR_DBERROR = "InternalError.DBError"

	// 获取配置失败。
	INTERNALERROR_DOMAINCONFIG = "InternalError.DomainConfig"

	// 上传链接生成失败。
	INTERNALERROR_FAILEDTOGENERATEURL = "InternalError.FailedToGenerateUrl"

	// 获取角色失败。
	INTERNALERROR_GETROLEERROR = "InternalError.GetRoleError"

	// 后端服务器发生未知错误。
	INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"

	// 配额系统处理失败。
	INTERNALERROR_QUOTASYSTEM = "InternalError.QuotaSystem"

	// 后端服务路由地址错误。
	INTERNALERROR_ROUTEERROR = "InternalError.RouteError"

	// 内部错误-系统错误。
	INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"

	// 参数错误。
	INVALIDPARAMETER = "InvalidParameter"

	// 操作频繁，请稍后重试。
	INVALIDPARAMETER_ACTIONINPROGRESS = "InvalidParameter.ActionInProgress"

	// 查询字符串规则超过了限制。
	INVALIDPARAMETER_CACHEKEYQUERYSTRINGTOOMANYVALUE = "InvalidParameter.CacheKeyQueryStringTooManyValue"

	// HTTPS证书链错误。
	INVALIDPARAMETER_CERTCHAINERROR = "InvalidParameter.CertChainError"

	// 证书错误。
	INVALIDPARAMETER_CERTCHECKERROR = "InvalidParameter.CertCheckError"

	// 证书错误。
	INVALIDPARAMETER_CERTCOMPLETEERROR = "InvalidParameter.CertCompleteError"

	// 证书错误。
	INVALIDPARAMETER_CERTFORMATERROR = "InvalidParameter.CertFormatError"

	// HTTPS证书已过期。
	INVALIDPARAMETER_CERTISEXPIRED = "InvalidParameter.CertIsExpired"

	// 证书错误。
	INVALIDPARAMETER_CERTNOCN = "InvalidParameter.CertNoCn"

	// 无效的HTTPS证书。
	INVALIDPARAMETER_CERTNOINFO = "InvalidParameter.CertNoInfo"

	// HTTPS证书和域名不匹配。
	INVALIDPARAMETER_CERTNOTMATCHDOMAIN = "InvalidParameter.CertNotMatchDomain"

	// HTTPS证书和密钥不匹配。
	INVALIDPARAMETER_CERTNOTMATCHKEY = "InvalidParameter.CertNotMatchKey"

	// 证书错误。
	INVALIDPARAMETER_CERTNOTPEM = "InvalidParameter.CertNotPem"

	// 内部错误。
	INVALIDPARAMETER_CERTSYSTEMERROR = "InvalidParameter.CertSystemError"

	// HTTPS证书即将过期。
	INVALIDPARAMETER_CERTTOEXPIRE = "InvalidParameter.CertToExpire"

	// 证书错误。
	INVALIDPARAMETER_CERTTOOSHORTKEYSIZE = "InvalidParameter.CertTooShortKeySize"

	// 证书错误。
	INVALIDPARAMETER_CERTUNSUPPORTEDTYPE = "InvalidParameter.CertUnsupportedType"

	// CNAME模式下无法申请泛域名证书。
	INVALIDPARAMETER_CNAMEWILDHOSTNOTALLOWAPPLYCERTIFICATE = "InvalidParameter.CnameWildHostNotAllowApplyCertificate"

	// 域名不存在或不属于该账号。
	INVALIDPARAMETER_DOMAINNOTFOUND = "InvalidParameter.DomainNotFound"

	// 当前域名已开启流量调度功能。
	INVALIDPARAMETER_DOMAINONTRAFFICSCHEDULING = "InvalidParameter.DomainOnTrafficScheduling"

	// 操作不支持条件。
	INVALIDPARAMETER_ERRACTIONUNSUPPORTTARGET = "InvalidParameter.ErrActionUnsupportTarget"

	// 非法操作。
	INVALIDPARAMETER_ERRINVALIDACTION = "InvalidParameter.ErrInvalidAction"

	// 非法操作-非法参数。
	INVALIDPARAMETER_ERRINVALIDACTIONPARAM = "InvalidParameter.ErrInvalidActionParam"

	// 非法操作-非法参数-无效action。
	INVALIDPARAMETER_ERRINVALIDACTIONPARAMACTION = "InvalidParameter.ErrInvalidActionParamAction"

	// 非法操作-非法参数-无效参数值类型。
	INVALIDPARAMETER_ERRINVALIDACTIONPARAMBADVALUETYPE = "InvalidParameter.ErrInvalidActionParamBadValueType"

	// 非法操作-非法参数-参数名重复。
	INVALIDPARAMETER_ERRINVALIDACTIONPARAMDUPLICATENAME = "InvalidParameter.ErrInvalidActionParamDuplicateName"

	// 非法操作-非法参数-无效参数值类型。
	INVALIDPARAMETER_ERRINVALIDACTIONPARAMNAME = "InvalidParameter.ErrInvalidActionParamName"

	// 非法操作-非法参数-参数值数量超出限制。
	INVALIDPARAMETER_ERRINVALIDACTIONPARAMTOOMANYVALUES = "InvalidParameter.ErrInvalidActionParamTooManyValues"

	// 非法操作-非法类型。
	INVALIDPARAMETER_ERRINVALIDACTIONTYPE = "InvalidParameter.ErrInvalidActionType"

	// 非法条件-非法忽略大小写。
	INVALIDPARAMETER_ERRINVALIDCONDITIONIGNORECASE = "InvalidParameter.ErrInvalidConditionIgnoreCase"

	// 非法条件-非法参数名称-无效的参数名称。
	INVALIDPARAMETER_ERRINVALIDCONDITIONNAMEBADNAME = "InvalidParameter.ErrInvalidConditionNameBadName"

	// 非法条件-非法参数名称-匹配类型不支持参数名称。
	INVALIDPARAMETER_ERRINVALIDCONDITIONNAMETARGETNOTSUPPORTNAME = "InvalidParameter.ErrInvalidConditionNameTargetNotSupportName"

	// 非法条件-非法参数值-无效的正则表达式。
	INVALIDPARAMETER_ERRINVALIDCONDITIONVALUEBADREGULAR = "InvalidParameter.ErrInvalidConditionValueBadRegular"

	// 非法条件-非法参数值-无效的url。
	INVALIDPARAMETER_ERRINVALIDCONDITIONVALUEBADURL = "InvalidParameter.ErrInvalidConditionValueBadUrl"

	// 非法条件-非法参数值-无效的参数值。
	INVALIDPARAMETER_ERRINVALIDCONDITIONVALUEBADVALUE = "InvalidParameter.ErrInvalidConditionValueBadValue"

	// 非法条件-非法参数值-无效的参数值-文件名不应包含文件后缀。
	INVALIDPARAMETER_ERRINVALIDCONDITIONVALUEBADVALUECONTAINFILENAMEEXTENSION = "InvalidParameter.ErrInvalidConditionValueBadValueContainFileNameExtension"

	// 非法条件-非法参数值-参数值长度超出限制。
	INVALIDPARAMETER_ERRINVALIDCONDITIONVALUETOOLONGVALUE = "InvalidParameter.ErrInvalidConditionValueTooLongValue"

	// 非法条件-非法参数值-正则表达式数量超出限制。
	INVALIDPARAMETER_ERRINVALIDCONDITIONVALUETOOMANYREGULAR = "InvalidParameter.ErrInvalidConditionValueTooManyRegular"

	// 非法条件-非法参数值-通配符数量超出限制。
	INVALIDPARAMETER_ERRINVALIDCONDITIONVALUETOOMANYWILDCARD = "InvalidParameter.ErrInvalidConditionValueTooManyWildcard"

	// 非法条件-非法参数值-参数值数量为0。
	INVALIDPARAMETER_ERRINVALIDCONDITIONVALUEZEROLENGTH = "InvalidParameter.ErrInvalidConditionValueZeroLength"

	// 开启 Grpc 协议支持需要同时开启 HTTP/2 协议支持。
	INVALIDPARAMETER_GRPCREQUIREHTTP2 = "InvalidParameter.GrpcRequireHttp2"

	// 域名不存在。
	INVALIDPARAMETER_HOSTNOTFOUND = "InvalidParameter.HostNotFound"

	// CNAME 未切换或者源站未路由到EO服务器。
	INVALIDPARAMETER_HOSTSTATUSNOTALLOWAPPLYCERTIFICATE = "InvalidParameter.HostStatusNotAllowApplyCertificate"

	// 参数错误。
	INVALIDPARAMETER_INVALIDACCELERATETYPE = "InvalidParameter.InvalidAccelerateType"

	// 无效的token鉴权。
	INVALIDPARAMETER_INVALIDAUTHENTICATION = "InvalidParameter.InvalidAuthentication"

	// 无效的token鉴权密钥。
	INVALIDPARAMETER_INVALIDAUTHENTICATIONTYPESECRETKEY = "InvalidParameter.InvalidAuthenticationTypeSecretKey"

	// 无效的token鉴权参数。
	INVALIDPARAMETER_INVALIDAUTHENTICATIONTYPESIGNPARAM = "InvalidParameter.InvalidAuthenticationTypeSignParam"

	// 无效的备源回源Host。
	INVALIDPARAMETER_INVALIDBACKUPSERVERNAME = "InvalidParameter.InvalidBackupServerName"

	// 无效的节点缓存。
	INVALIDPARAMETER_INVALIDCACHECONFIGCACHE = "InvalidParameter.InvalidCacheConfigCache"

	// 无效的节点缓存，遵循源站行为。
	INVALIDPARAMETER_INVALIDCACHECONFIGFOLLOWORIGIN = "InvalidParameter.InvalidCacheConfigFollowOrigin"

	// 无效的缓存键。
	INVALIDPARAMETER_INVALIDCACHEKEY = "InvalidParameter.InvalidCacheKey"

	// 无效的查询字符串。
	INVALIDPARAMETER_INVALIDCACHEKEYQUERYSTRINGVALUE = "InvalidParameter.InvalidCacheKeyQueryStringValue"

	// 无效的节点缓存。
	INVALIDPARAMETER_INVALIDCACHEONLYONSWITCH = "InvalidParameter.InvalidCacheOnlyOnSwitch"

	// 无效的节点缓存时间。
	INVALIDPARAMETER_INVALIDCACHETIME = "InvalidParameter.InvalidCacheTime"

	// 证书信息错误。
	INVALIDPARAMETER_INVALIDCERTINFO = "InvalidParameter.InvalidCertInfo"

	// 无效的客户端IP请求头。
	INVALIDPARAMETER_INVALIDCLIENTIPHEADERNAME = "InvalidParameter.InvalidClientIpHeaderName"

	// 无效的智能加速。
	INVALIDPARAMETER_INVALIDDYNAMICROUTINE = "InvalidParameter.InvalidDynamicRoutine"

	// 套餐包不支持智能加速配置。
	INVALIDPARAMETER_INVALIDDYNAMICROUTINEBILLING = "InvalidParameter.InvalidDynamicRoutineBilling"

	// 无效的自定义错误页面。
	INVALIDPARAMETER_INVALIDERRORPAGEREDIRECTURL = "InvalidParameter.InvalidErrorPageRedirectUrl"

	// 无效的HTTPS。
	INVALIDPARAMETER_INVALIDHTTPS = "InvalidParameter.InvalidHttps"

	// 无效的HTTPS HSTS。
	INVALIDPARAMETER_INVALIDHTTPSHSTSMAXAGE = "InvalidParameter.InvalidHttpsHstsMaxAge"

	// 无效的HTTPS TLS版本。
	INVALIDPARAMETER_INVALIDHTTPSTLSVERSION = "InvalidParameter.InvalidHttpsTlsVersion"

	// 无效的Ipv6开关配置。
	INVALIDPARAMETER_INVALIDIPV6SWITCH = "InvalidParameter.InvalidIpv6Switch"

	// 无效的源站。
	INVALIDPARAMETER_INVALIDORIGIN = "InvalidParameter.InvalidOrigin"

	// 参数错误。
	INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"

	// 套餐包不支持最大上传大小。
	INVALIDPARAMETER_INVALIDPOSTMAXSIZEBILLING = "InvalidParameter.InvalidPostMaxSizeBilling"

	// 无效的最大上传大小。
	INVALIDPARAMETER_INVALIDPOSTSIZEVALUE = "InvalidParameter.InvalidPostSizeValue"

	// 套餐包不支持Quic配置。
	INVALIDPARAMETER_INVALIDQUICBILLING = "InvalidParameter.InvalidQuicBilling"

	// 无效的请求头header。
	INVALIDPARAMETER_INVALIDREQUESTHEADERNAME = "InvalidParameter.InvalidRequestHeaderName"

	// x-forwarded-for 请求头配置无效
	INVALIDPARAMETER_INVALIDREQUESTHEADERNAMEXFF = "InvalidParameter.InvalidRequestHeaderNameXff"

	// 无效的请求头header。
	INVALIDPARAMETER_INVALIDREQUESTHEADERVALUE = "InvalidParameter.InvalidRequestHeaderValue"

	// 无套餐包。
	INVALIDPARAMETER_INVALIDRESOURCEIDBILLING = "InvalidParameter.InvalidResourceIdBilling"

	// 无效的响应头header。
	INVALIDPARAMETER_INVALIDRESPONSEHEADERNAME = "InvalidParameter.InvalidResponseHeaderName"

	// 无效的响应头header。
	INVALIDPARAMETER_INVALIDRESPONSEHEADERVALUE = "InvalidParameter.InvalidResponseHeaderValue"

	// 无效的规则引擎配置。
	INVALIDPARAMETER_INVALIDRULEENGINE = "InvalidParameter.InvalidRuleEngine"

	// 无效的规则引擎操作。
	INVALIDPARAMETER_INVALIDRULEENGINEACTION = "InvalidParameter.InvalidRuleEngineAction"

	// 规则不存在。
	INVALIDPARAMETER_INVALIDRULEENGINENOTFOUND = "InvalidParameter.InvalidRuleEngineNotFound"

	// 无效的规则引擎条件。
	INVALIDPARAMETER_INVALIDRULEENGINETARGET = "InvalidParameter.InvalidRuleEngineTarget"

	// 无效的规则引擎文件后缀条件。
	INVALIDPARAMETER_INVALIDRULEENGINETARGETSEXTENSION = "InvalidParameter.InvalidRuleEngineTargetsExtension"

	// 无效的规则引擎URL条件。
	INVALIDPARAMETER_INVALIDRULEENGINETARGETSURL = "InvalidParameter.InvalidRuleEngineTargetsUrl"

	// 无效的回源Host。
	INVALIDPARAMETER_INVALIDSERVERNAME = "InvalidParameter.InvalidServerName"

	// URL重写的目标HOST无效。
	INVALIDPARAMETER_INVALIDURLREDIRECTHOST = "InvalidParameter.InvalidUrlRedirectHost"

	// URL重写的目标URL无效。
	INVALIDPARAMETER_INVALIDURLREDIRECTURL = "InvalidParameter.InvalidUrlRedirectUrl"

	// 无效的WebSocket。
	INVALIDPARAMETER_INVALIDWEBSOCKETTIMEOUT = "InvalidParameter.InvalidWebSocketTimeout"

	// 无效的缓存键。
	INVALIDPARAMETER_KEYRULESINVALIDQUERYSTRINGVALUE = "InvalidParameter.KeyRulesInvalidQueryStringValue"

	// 参数长度超过限制。
	INVALIDPARAMETER_LENGTHEXCEEDSLIMIT = "InvalidParameter.LengthExceedsLimit"

	// 源站是内网IP。
	INVALIDPARAMETER_ORIGINISINNERIP = "InvalidParameter.OriginIsInnerIp"

	// 修改源站操作中源站组Id必填。
	INVALIDPARAMETER_ORIGINORIGINGROUPIDISREQUIRED = "InvalidParameter.OriginOriginGroupIdIsRequired"

	// 参数错误。
	INVALIDPARAMETER_PARAMETERERROR = "InvalidParameter.ParameterError"

	// 套餐不存在。
	INVALIDPARAMETER_PLANNOTFOUND = "InvalidParameter.PlanNotFound"

	// 最大上传大小超出限制
	INVALIDPARAMETER_POSTMAXSIZELIMITEXCEEDED = "InvalidParameter.PostMaxSizeLimitExceeded"

	// 参数错误
	INVALIDPARAMETER_SECURITY = "InvalidParameter.Security"

	// 配置项参数错误。
	INVALIDPARAMETER_SETTINGINVALIDPARAM = "InvalidParameter.SettingInvalidParam"

	// 资源存在错误。
	INVALIDPARAMETER_TARGET = "InvalidParameter.Target"

	// 任务无法生成。
	INVALIDPARAMETER_TASKNOTGENERATED = "InvalidParameter.TaskNotGenerated"

	// 内部错误。
	INVALIDPARAMETER_TASKSYSTEMERROR = "InvalidParameter.TaskSystemError"

	// 文件上传链接存在问题。
	INVALIDPARAMETER_UPLOADURL = "InvalidParameter.UploadUrl"

	// 站点已被绑定。
	INVALIDPARAMETER_ZONEHASBEENBOUND = "InvalidParameter.ZoneHasBeenBound"

	// 站点不存在。
	INVALIDPARAMETER_ZONENOTFOUND = "InvalidParameter.ZoneNotFound"

	// 参数取值错误。
	INVALIDPARAMETERVALUE = "InvalidParameterValue"

	// DNS 记录与 DNSSEC 功能冲突。
	INVALIDPARAMETERVALUE_CONFLICTWITHDNSSEC = "InvalidParameterValue.ConflictWithDNSSEC"

	// 主机记录与记录值不能取值相同。
	INVALIDPARAMETERVALUE_CONTENTSAMEASNAME = "InvalidParameterValue.ContentSameAsName"

	// 入参中的域名与站点参数不匹配，请更正后重试。
	INVALIDPARAMETERVALUE_DOMAINNOTMATCHZONE = "InvalidParameterValue.DomainNotMatchZone"

	// DNS 记录内容错误。
	INVALIDPARAMETERVALUE_INVALIDDNSCONTENT = "InvalidParameterValue.InvalidDNSContent"

	// DNS 记录名称错误。
	INVALIDPARAMETERVALUE_INVALIDDNSNAME = "InvalidParameterValue.InvalidDNSName"

	// 加速域名状态不符合要求。
	INVALIDPARAMETERVALUE_INVALIDDOMAINSTATUS = "InvalidParameterValue.InvalidDomainStatus"

	// 该同名站点标识已被占用，请重新输入。
	INVALIDPARAMETERVALUE_ZONESAMEASNAME = "InvalidParameterValue.ZoneSameAsName"

	// 超过配额限制。
	LIMITEXCEEDED = "LimitExceeded"

	// 本次提交的资源数超过上限。
	LIMITEXCEEDED_BATCHQUOTA = "LimitExceeded.BatchQuota"

	// 当天提交的资源数超过上限。
	LIMITEXCEEDED_DAILYQUOTA = "LimitExceeded.DailyQuota"

	// 查询时间范围超出限制。
	LIMITEXCEEDED_QUERYTIMELIMITEXCEEDED = "LimitExceeded.QueryTimeLimitExceeded"

	// 单位时间内接口请求频率达到限制。
	LIMITEXCEEDED_RATELIMITEXCEEDED = "LimitExceeded.RateLimitExceeded"

	// 操作被拒绝。
	OPERATIONDENIED = "OperationDenied"

	// 域名被封禁，暂时无法操作。
	OPERATIONDENIED_DOMAINISBLOCKED = "OperationDenied.DomainIsBlocked"

	// 域名尚未备案。
	OPERATIONDENIED_DOMAINNOICP = "OperationDenied.DomainNoICP"

	// 4层代理资源处于封禁中，禁止操作。
	OPERATIONDENIED_L4PROXYINBANNEDSTATUS = "OperationDenied.L4ProxyInBannedStatus"

	// 绑定4层实例有处于非运行中的状态，禁止操作。
	OPERATIONDENIED_L4STATUSNOTINONLINE = "OperationDenied.L4StatusNotInOnline"

	// 回源白名单已经是最新版本，无需更新。
	OPERATIONDENIED_LATESTVERSIONNOW = "OperationDenied.LatestVersionNow"

	// 已存在多个Cname接入站点，不允许切换至NS。
	OPERATIONDENIED_MULTIPLECNAMEZONE = "OperationDenied.MultipleCnameZone"

	// NS接入模式不支持域名流量调度功能。
	OPERATIONDENIED_NSNOTALLOWTRAFFICSTRATEGY = "OperationDenied.NSNotAllowTrafficStrategy"

	// DNS 记录不允许添加。
	OPERATIONDENIED_RECORDISFORBIDDEN = "OperationDenied.RecordIsForbidden"

	// 资源被占用。
	RESOURCEINUSE = "ResourceInUse"

	// 资源被本账号别称域名占用。
	RESOURCEINUSE_ALIASDOMAIN = "ResourceInUse.AliasDomain"

	// 资源被本账号Cname接入占用。
	RESOURCEINUSE_CNAME = "ResourceInUse.Cname"

	// Dns资源被占用。
	RESOURCEINUSE_DNS = "ResourceInUse.Dns"

	// 该域名正在解析中，如果您需要开启加速，请前往 DNS 记录页开启加速。
	RESOURCEINUSE_DNSRECORD = "ResourceInUse.DnsRecord"

	// 已存在相同的别称域名。
	RESOURCEINUSE_DUPLICATENAME = "ResourceInUse.DuplicateName"

	// 资源已被泛域名占用。
	RESOURCEINUSE_GENERICHOST = "ResourceInUse.GenericHost"

	// 资源被本账号的子域名占用。
	RESOURCEINUSE_HOST = "ResourceInUse.Host"

	// 资源被本账号NS接入占用。
	RESOURCEINUSE_NS = "ResourceInUse.NS"

	// 资源被其他用户接入。
	RESOURCEINUSE_OTHERS = "ResourceInUse.Others"

	// 资源被其他账号别称域名占用。
	RESOURCEINUSE_OTHERSALIASDOMAIN = "ResourceInUse.OthersAliasDomain"

	// 资源被其他账号Cname接入占用。
	RESOURCEINUSE_OTHERSCNAME = "ResourceInUse.OthersCname"

	// 资源被其他账号的子域名占用。
	RESOURCEINUSE_OTHERSHOST = "ResourceInUse.OthersHost"

	// 资源被其他账号NS接入占用。
	RESOURCEINUSE_OTHERSNS = "ResourceInUse.OthersNS"

	// 资源被本账号和其他账号同时Cname接入占用。
	RESOURCEINUSE_SELFANDOTHERSCNAME = "ResourceInUse.SelfAndOthersCname"

	// 别称域名已站点接入。
	RESOURCEINUSE_ZONE = "ResourceInUse.Zone"

	// 资源不足。
	RESOURCEINSUFFICIENT = "ResourceInsufficient"

	// 资源不存在。
	RESOURCENOTFOUND = "ResourceNotFound"

	// 资源不可用。
	RESOURCEUNAVAILABLE = "ResourceUnavailable"

	// 证书不存在或未授权。
	RESOURCEUNAVAILABLE_CERTNOTFOUND = "ResourceUnavailable.CertNotFound"

	// 当前域名已接入EdgeOne，如确认需要接入到当前账号，请进行域名找回。
	RESOURCEUNAVAILABLE_DOMAINALREADYEXISTS = "ResourceUnavailable.DomainAlreadyExists"

	// 请求的加速域名不存在，请更正后重试。
	RESOURCEUNAVAILABLE_DOMAINNOTFOUND = "ResourceUnavailable.DomainNotFound"

	// 域名不存在或未开启代理。
	RESOURCEUNAVAILABLE_HOSTNOTFOUND = "ResourceUnavailable.HostNotFound"

	// 未拉取到已开启代理的zone信息。
	RESOURCEUNAVAILABLE_PROXYZONENOTFOUND = "ResourceUnavailable.ProxyZoneNotFound"

	// 站点不存在或不属于该账号。
	RESOURCEUNAVAILABLE_ZONENOTFOUND = "ResourceUnavailable.ZoneNotFound"

	// 资源售卖火爆，已售罄，正在加紧补货中，当前无法新增域名，请您耐心等待。
	RESOURCESSOLDOUT_L7LACKOFRESOURCES = "ResourcesSoldOut.L7LackOfResources"

	// 未授权操作。
	UNAUTHORIZEDOPERATION = "UnauthorizedOperation"

	// Cam 未授权。
	UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"

	// 鉴权错误。
	UNAUTHORIZEDOPERATION_DOMAINEMPTY = "UnauthorizedOperation.DomainEmpty"

	// 子账户没有操作权限，请添加权限后继续操作。
	UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"

	// 后端服务器发生未知错误。
	UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"

	// 操作不支持。
	UNSUPPORTEDOPERATION = "UnsupportedOperation"
)
