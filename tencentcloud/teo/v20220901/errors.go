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

	// DryRun 操作，代表请求将会是成功的，只是多传了 DryRun 参数。
	DRYRUNOPERATION = "DryRunOperation"

	// 操作失败。
	FAILEDOPERATION = "FailedOperation"

	// 证书已过期，暂不支持下发过期证书。
	FAILEDOPERATION_CERTIFICATEHASEXPIRED = "FailedOperation.CertificateHasExpired"

	// 证书不存在。
	FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"

	// 配置文件Condition表达式语法错误。
	FAILEDOPERATION_CONFIGCONDITIONSYNTAXERROR = "FailedOperation.ConfigConditionSyntaxError"

	// 无法识别的Condition匹配类型。
	FAILEDOPERATION_CONFIGCONDITIONUNKNOWNTARGET = "FailedOperation.ConfigConditionUnknownTarget"

	// 配置文件Condition表达式值的项不能为空。
	FAILEDOPERATION_CONFIGCONDITIONVALUEEMPTYERROR = "FailedOperation.ConfigConditionValueEmptyError"

	// 配置文件存在类型不匹配的字段。
	FAILEDOPERATION_CONFIGFIELDTYPEERROR = "FailedOperation.ConfigFieldTypeError"

	// 配置文件存在语法错误。
	FAILEDOPERATION_CONFIGFORMATERROR = "FailedOperation.ConfigFormatError"

	// 配置文件内容格式错误，无法解析。
	FAILEDOPERATION_CONFIGMALFORMEDCONTENT = "FailedOperation.ConfigMalformedContent"

	// 配置文件参数校验错误。
	FAILEDOPERATION_CONFIGPARAMVALIDATEERRORS = "FailedOperation.ConfigParamValidateErrors"

	// 配置文件无法识别的字段：存在拼写错误，或者该字段所在的层级出错。
	FAILEDOPERATION_CONFIGUNKNOWNFIELD = "FailedOperation.ConfigUnknownField"

	// 当前不支持该配置文件版本。
	FAILEDOPERATION_CONFIGUNSUPPORTEDFORMATVERSION = "FailedOperation.ConfigUnsupportedFormatVersion"

	// 创建日志集失败，请检查日志集名是否已存在。
	FAILEDOPERATION_CREATECLSLOGSETFAILED = "FailedOperation.CreateClsLogSetFailed"

	// 创建日志主题任务失败，请检查日志主题名或任务名是否已存在。
	FAILEDOPERATION_CREATECLSLOGTOPICTASKFAILED = "FailedOperation.CreateClsLogTopicTaskFailed"

	// 创建自定义推送任务认证失败, 请检查推送地址是否正确。
	FAILEDOPERATION_CREATELOGTOPICTASKAUTHFAILURE = "FailedOperation.CreateLogTopicTaskAuthFailure"

	// 有其他任务正在部署中，请稍后重试。
	FAILEDOPERATION_FUNCTIONDEPLOYING = "FailedOperation.FunctionDeploying"

	// 账户余额不足。
	FAILEDOPERATION_INSUFFICIENTACCOUNTBALANCE = "FailedOperation.InsufficientAccountBalance"

	// 站点状态不正确。
	FAILEDOPERATION_INVALIDZONESTATUS = "FailedOperation.InvalidZoneStatus"

	// content缺少必带的配置块。
	FAILEDOPERATION_MISSINGCONFIGCHUNK = "FailedOperation.MissingConfigChunk"

	// 操作失败。
	FAILEDOPERATION_MODIFYFAILED = "FailedOperation.ModifyFailed"

	// 实时日志认证失败
	FAILEDOPERATION_REALTIMELOGAUTHFAILURE = "FailedOperation.RealtimeLogAuthFailure"

	// 实时日志推送任务不存在
	FAILEDOPERATION_REALTIMELOGNOTFOUND = "FailedOperation.RealtimeLogNotFound"

	// 该站点下函数规则操作冲突。
	FAILEDOPERATION_RULEOPERATIONCONFLICT = "FailedOperation.RuleOperationConflict"

	// 未知的配置组类型。
	FAILEDOPERATION_UNKNOWNCONFIGGROUPTYPE = "FailedOperation.UnknownConfigGroupType"

	// 内部错误。
	INTERNALERROR = "InternalError"

	// 后台处理出错。
	INTERNALERROR_BACKENDERROR = "InternalError.BackendError"

	// 配置已被锁定，请解除配置锁定之后在重试。
	INTERNALERROR_CONFIGLOCKED = "InternalError.ConfigLocked"

	// 数据库错误。
	INTERNALERROR_DBERROR = "InternalError.DBError"

	// 获取配置失败。
	INTERNALERROR_DOMAINCONFIG = "InternalError.DomainConfig"

	// 调用 DNSPod 失败，请稍后重试，若无法解决，请联系智能客服或提交工单。
	INTERNALERROR_FAILEDTOCALLDNSPOD = "InternalError.FailedToCallDNSPod"

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

	// 未知错误。
	INTERNALERROR_UNKNOWERROR = "InternalError.UnknowError"

	// 参数错误。
	INVALIDPARAMETER = "InvalidParameter"

	// 正则表达式非标准RE2格式。
	INVALIDPARAMETER_ACCESSREDIRECTREGEXERROR = "InvalidParameter.AccessRedirectRegexError"

	// 操作频繁，请稍后重试。
	INVALIDPARAMETER_ACTIONINPROGRESS = "InvalidParameter.ActionInProgress"

	// 别称域名不支持配置国密证书。
	INVALIDPARAMETER_ALIASDOMAINNOTSUPPORTSMCERT = "InvalidParameter.AliasDomainNotSupportSMCert"

	// 函数内容存在语法错误。
	INVALIDPARAMETER_BADCONTENT = "InvalidParameter.BadContent"

	// 函数名称不符合命名规范。
	INVALIDPARAMETER_BADFUNCTIONNAME = "InvalidParameter.BadFunctionName"

	// 无效的查询字符串。
	INVALIDPARAMETER_CACHEKEYQUERYSTRINGREQUIRESFULLURLCACHEOFF = "InvalidParameter.CacheKeyQueryStringRequiresFullUrlCacheOff"

	// 查询字符串规则超过了限制。
	INVALIDPARAMETER_CACHEKEYQUERYSTRINGTOOMANYVALUE = "InvalidParameter.CacheKeyQueryStringTooManyValue"

	// HTTPS证书和域名不匹配。
	INVALIDPARAMETER_CERTNOTMATCHDOMAIN = "InvalidParameter.CertNotMatchDomain"

	// 内部错误。
	INVALIDPARAMETER_CERTSYSTEMERROR = "InvalidParameter.CertSystemError"

	// HTTPS证书即将过期。
	INVALIDPARAMETER_CERTTOEXPIRE = "InvalidParameter.CertToExpire"

	// 证书错误。
	INVALIDPARAMETER_CERTTOOSHORTKEYSIZE = "InvalidParameter.CertTooShortKeySize"

	// IPv6 访问与客户端 IP 地理位置功能冲突。
	INVALIDPARAMETER_CLIENTIPCOUNTRYCONFLICTSWITHIPV6 = "InvalidParameter.ClientIpCountryConflictsWithIpv6"

	// CNAME模式下无法申请泛域名证书。
	INVALIDPARAMETER_CNAMEWILDHOSTNOTALLOWAPPLYCERTIFICATE = "InvalidParameter.CnameWildHostNotAllowApplyCertificate"

	// 无效的压缩算法。
	INVALIDPARAMETER_COMPRESSIONINVALIDALGORITHMS = "InvalidParameter.CompressionInvalidAlgorithms"

	// 源站不能和域名一致。
	INVALIDPARAMETER_CONFLICTHOSTORIGIN = "InvalidParameter.ConflictHostOrigin"

	// 函数内容超过大小限制。
	INVALIDPARAMETER_CONTENTEXCEEDSLIMIT = "InvalidParameter.ContentExceedsLimit"

	// 域名不存在或不属于该账号。
	INVALIDPARAMETER_DOMAINNOTFOUND = "InvalidParameter.DomainNotFound"

	// 当前域名已开启流量调度功能。
	INVALIDPARAMETER_DOMAINONTRAFFICSCHEDULING = "InvalidParameter.DomainOnTrafficScheduling"

	// 重复规则。
	INVALIDPARAMETER_DUPLICATERULE = "InvalidParameter.DuplicateRule"

	// 操作不支持条件。
	INVALIDPARAMETER_ERRACTIONUNSUPPORTTARGET = "InvalidParameter.ErrActionUnsupportTarget"

	// 非法操作。
	INVALIDPARAMETER_ERRINVALIDACTION = "InvalidParameter.ErrInvalidAction"

	// 非法操作-操作配置重复。
	INVALIDPARAMETER_ERRINVALIDACTIONDUPLICATEACTION = "InvalidParameter.ErrInvalidActionDuplicateAction"

	// 无效的规则引擎操作，源站IP不支持内网IP或回环地址。
	INVALIDPARAMETER_ERRINVALIDACTIONORIGINPRIVATEADDRESS = "InvalidParameter.ErrInvalidActionOriginPrivateAddress"

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

	// 非法操作-非法参数-无效参数值。
	INVALIDPARAMETER_ERRINVALIDACTIONPARAMVALUE = "InvalidParameter.ErrInvalidActionParamValue"

	// 非法操作-非法类型。
	INVALIDPARAMETER_ERRINVALIDACTIONTYPE = "InvalidParameter.ErrInvalidActionType"

	// 非法条件。
	INVALIDPARAMETER_ERRINVALIDCONDITION = "InvalidParameter.ErrInvalidCondition"

	// 修改源站操作不能仅配置host匹配类型。
	INVALIDPARAMETER_ERRINVALIDCONDITIONCANNOTONLYCONTAINHOSTWHENMODIFYORIGINACTIONCONFIGURED = "InvalidParameter.ErrInvalidConditionCannotOnlyContainHostWhenModifyOriginActionConfigured"

	// 修改源站操作仅支持配置一个host匹配类型。
	INVALIDPARAMETER_ERRINVALIDCONDITIONHOSTTOOMANYWHENMODIFYORIGINACTIONCONFIGURED = "InvalidParameter.ErrInvalidConditionHostTooManyWhenModifyOriginActionConfigured"

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

	// 非法条件-非法参数值-参数值数量超出限制。
	INVALIDPARAMETER_ERRINVALIDCONDITIONVALUETOOMANYVALUES = "InvalidParameter.ErrInvalidConditionValueTooManyValues"

	// 非法条件-非法参数值-通配符数量超出限制。
	INVALIDPARAMETER_ERRINVALIDCONDITIONVALUETOOMANYWILDCARD = "InvalidParameter.ErrInvalidConditionValueTooManyWildcard"

	// 非法条件-非法参数值-参数值数量为0。
	INVALIDPARAMETER_ERRINVALIDCONDITIONVALUEZEROLENGTH = "InvalidParameter.ErrInvalidConditionValueZeroLength"

	// 修改源站操作不支持ELSE。
	INVALIDPARAMETER_ERRINVALIDELSEWHENMODIFYORIGINACTIONCONFIGURED = "InvalidParameter.ErrInvalidElseWhenModifyOriginActionConfigured"

	// 条件为空。
	INVALIDPARAMETER_ERRNILCONDITION = "InvalidParameter.ErrNilCondition"

	// 函数名称和本账号下其他函数冲突。
	INVALIDPARAMETER_FUNCTIONNAMECONFLICT = "InvalidParameter.FunctionNameConflict"

	// 开启 Grpc 协议支持需要同时开启 HTTP/2 协议支持。
	INVALIDPARAMETER_GRPCREQUIREHTTP2 = "InvalidParameter.GrpcRequireHttp2"

	// 回源Host错误。
	INVALIDPARAMETER_HOSTHEADERINVALID = "InvalidParameter.HostHeaderInvalid"

	// 域名不存在。
	INVALIDPARAMETER_HOSTNOTFOUND = "InvalidParameter.HostNotFound"

	// CNAME 未切换或者源站未路由到EO服务器。
	INVALIDPARAMETER_HOSTSTATUSNOTALLOWAPPLYCERTIFICATE = "InvalidParameter.HostStatusNotAllowApplyCertificate"

	// 参数错误。
	INVALIDPARAMETER_INVALIDACCELERATETYPE = "InvalidParameter.InvalidAccelerateType"

	// 无效的token鉴权。
	INVALIDPARAMETER_INVALIDAUTHENTICATION = "InvalidParameter.InvalidAuthentication"

	// 无效的token鉴权过期时间。
	INVALIDPARAMETER_INVALIDAUTHENTICATIONTYPEEXPIRETIME = "InvalidParameter.InvalidAuthenticationTypeExpireTime"

	// 无效的token鉴权密钥。
	INVALIDPARAMETER_INVALIDAUTHENTICATIONTYPESECRETKEY = "InvalidParameter.InvalidAuthenticationTypeSecretKey"

	// 无效的token鉴权参数。
	INVALIDPARAMETER_INVALIDAUTHENTICATIONTYPESIGNPARAM = "InvalidParameter.InvalidAuthenticationTypeSignParam"

	// 无效的token鉴权时间格式。
	INVALIDPARAMETER_INVALIDAUTHENTICATIONTYPETIMEFORMAT = "InvalidParameter.InvalidAuthenticationTypeTimeFormat"

	// 无效的token鉴权时间参数。
	INVALIDPARAMETER_INVALIDAUTHENTICATIONTYPETIMEPARAM = "InvalidParameter.InvalidAuthenticationTypeTimeParam"

	// 自动使用代金券格式不正确，请输入正确的自动使用代金券格式。
	INVALIDPARAMETER_INVALIDAUTOUSEVOUCHER = "InvalidParameter.InvalidAutoUseVoucher"

	// 无效的第三方对象存储。
	INVALIDPARAMETER_INVALIDAWSPRIVATEACCESS = "InvalidParameter.InvalidAwsPrivateAccess"

	// 请正确填写地域作为第三方对象存储私有访问参数。
	INVALIDPARAMETER_INVALIDAWSREGION = "InvalidParameter.InvalidAwsRegion"

	// 无效的第三方对象存储。
	INVALIDPARAMETER_INVALIDAWSSECRETKEY = "InvalidParameter.InvalidAwsSecretKey"

	// 无效的备源回源Host。
	INVALIDPARAMETER_INVALIDBACKUPSERVERNAME = "InvalidParameter.InvalidBackupServerName"

	// 无效的节点缓存。
	INVALIDPARAMETER_INVALIDCACHECONFIGCACHE = "InvalidParameter.InvalidCacheConfigCache"

	// 无效的节点缓存，遵循源站行为。
	INVALIDPARAMETER_INVALIDCACHECONFIGFOLLOWORIGIN = "InvalidParameter.InvalidCacheConfigFollowOrigin"

	// 无效的缓存键。
	INVALIDPARAMETER_INVALIDCACHEKEY = "InvalidParameter.InvalidCacheKey"

	// 无效的自定义Cache Key Cookie无效。
	INVALIDPARAMETER_INVALIDCACHEKEYCOOKIE = "InvalidParameter.InvalidCacheKeyCookie"

	// 无效的缓存键忽略大小写。
	INVALIDPARAMETER_INVALIDCACHEKEYIGNORECASE = "InvalidParameter.InvalidCacheKeyIgnoreCase"

	// 无效的查询字符串。
	INVALIDPARAMETER_INVALIDCACHEKEYQUERYSTRINGACTION = "InvalidParameter.InvalidCacheKeyQueryStringAction"

	// 无效的查询字符串。
	INVALIDPARAMETER_INVALIDCACHEKEYQUERYSTRINGVALUE = "InvalidParameter.InvalidCacheKeyQueryStringValue"

	// 无效的自定义Cache Key Scheme无效。
	INVALIDPARAMETER_INVALIDCACHEKEYSCHEME = "InvalidParameter.InvalidCacheKeyScheme"

	// 无效的节点缓存。
	INVALIDPARAMETER_INVALIDCACHEONLYONSWITCH = "InvalidParameter.InvalidCacheOnlyOnSwitch"

	// 无效的节点缓存时间。
	INVALIDPARAMETER_INVALIDCACHETIME = "InvalidParameter.InvalidCacheTime"

	// 证书信息错误。
	INVALIDPARAMETER_INVALIDCERTINFO = "InvalidParameter.InvalidCertInfo"

	// 无效的客户端IP地理位置配置，HeaderName由1-100个字母或数字组成的，不能以"-"开头或结尾。
	INVALIDPARAMETER_INVALIDCLIENTIPCOUNTRYHEADERNAME = "InvalidParameter.InvalidClientIpCountryHeaderName"

	// 无效的客户端IP请求头。
	INVALIDPARAMETER_INVALIDCLIENTIPHEADERNAME = "InvalidParameter.InvalidClientIpHeaderName"

	// 分区域回源的源站无效。
	INVALIDPARAMETER_INVALIDCLIENTIPORIGIN = "InvalidParameter.InvalidClientIpOrigin"

	// 不合法的条件。
	INVALIDPARAMETER_INVALIDCONDITIONS = "InvalidParameter.InvalidConditions"

	// 无效的回源配置，源站类型为对象存储时，源站地址必须为对象存储域名。
	INVALIDPARAMETER_INVALIDCOSDOMAIN = "InvalidParameter.InvalidCosDomain"

	// 无效的智能加速。
	INVALIDPARAMETER_INVALIDDYNAMICROUTINE = "InvalidParameter.InvalidDynamicRoutine"

	// 套餐包不支持智能加速配置。
	INVALIDPARAMETER_INVALIDDYNAMICROUTINEBILLING = "InvalidParameter.InvalidDynamicRoutineBilling"

	// 无效的自定义错误页面。
	INVALIDPARAMETER_INVALIDERRORPAGE = "InvalidParameter.InvalidErrorPage"

	// 无效的自定义错误页面。
	INVALIDPARAMETER_INVALIDERRORPAGEREDIRECTURL = "InvalidParameter.InvalidErrorPageRedirectUrl"

	// 不合法的过滤字段。
	INVALIDPARAMETER_INVALIDFILTERNAME = "InvalidParameter.InvalidFilterName"

	// 无效的强制HTTPS跳转。
	INVALIDPARAMETER_INVALIDFORCEREDIRECTTYPE = "InvalidParameter.InvalidForceRedirectType"

	// 无效的HTTPS。
	INVALIDPARAMETER_INVALIDHTTPS = "InvalidParameter.InvalidHttps"

	// 无效的HTTPS证书。
	INVALIDPARAMETER_INVALIDHTTPSCERTINFO = "InvalidParameter.InvalidHttpsCertInfo"

	// 加密套件与TLS版本不匹配。
	INVALIDPARAMETER_INVALIDHTTPSCIPHERSUITEANDTLSVERSION = "InvalidParameter.InvalidHttpsCipherSuiteAndTlsVersion"

	// 无效的HTTPS HSTS。
	INVALIDPARAMETER_INVALIDHTTPSHSTSMAXAGE = "InvalidParameter.InvalidHttpsHstsMaxAge"

	// 无效的HTTPS TLS版本。
	INVALIDPARAMETER_INVALIDHTTPSTLSVERSION = "InvalidParameter.InvalidHttpsTlsVersion"

	// 无效时间间隔。取值应为[min 5min hour day]。
	INVALIDPARAMETER_INVALIDINTERVAL = "InvalidParameter.InvalidInterval"

	// 无效的Ipv6开关配置。
	INVALIDPARAMETER_INVALIDIPV6SWITCH = "InvalidParameter.InvalidIpv6Switch"

	// 日志输出格式的字段分隔符不正确
	INVALIDPARAMETER_INVALIDLOGFORMATFIELDDELIMITER = "InvalidParameter.InvalidLogFormatFieldDelimiter"

	// 日志输出格式类型不正确
	INVALIDPARAMETER_INVALIDLOGFORMATFORMATTYPE = "InvalidParameter.InvalidLogFormatFormatType"

	// 日志输出格式的日志记录分隔符不正确
	INVALIDPARAMETER_INVALIDLOGFORMATRECORDDELIMITER = "InvalidParameter.InvalidLogFormatRecordDelimiter"

	// 无效的浏览器缓存。
	INVALIDPARAMETER_INVALIDMAXAGEFOLLOWORIGIN = "InvalidParameter.InvalidMaxAgeFollowOrigin"

	// 无效的浏览器缓存。
	INVALIDPARAMETER_INVALIDMAXAGETIME = "InvalidParameter.InvalidMaxAgeTime"

	// 无效查询维度。
	INVALIDPARAMETER_INVALIDMETRIC = "InvalidParameter.InvalidMetric"

	// 无效的源站。
	INVALIDPARAMETER_INVALIDORIGIN = "InvalidParameter.InvalidOrigin"

	// 源站组类型错误。
	INVALIDPARAMETER_INVALIDORIGINGROUPTYPE = "InvalidParameter.InvalidOriginGroupType"

	// 不支持填写内网IP/回环地址作为源站地址
	INVALIDPARAMETER_INVALIDORIGINIP = "InvalidParameter.InvalidOriginIp"

	// 源站类型错误。
	INVALIDPARAMETER_INVALIDORIGINTYPE = "InvalidParameter.InvalidOriginType"

	// 源站错误或不存在。
	INVALIDPARAMETER_INVALIDORIGINVALUE = "InvalidParameter.InvalidOriginValue"

	// 参数错误。
	INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"

	// 套餐周期格式不正确，请输入正确的套餐周期格式。
	INVALIDPARAMETER_INVALIDPERIOD = "InvalidParameter.InvalidPeriod"

	// 套餐类型格式不正确，请输入正确的套餐类型格式。
	INVALIDPARAMETER_INVALIDPLANTYPE = "InvalidParameter.InvalidPlanType"

	// 套餐包不支持最大上传大小。
	INVALIDPARAMETER_INVALIDPOSTMAXSIZEBILLING = "InvalidParameter.InvalidPostMaxSizeBilling"

	// 无效的最大上传大小。
	INVALIDPARAMETER_INVALIDPOSTSIZEVALUE = "InvalidParameter.InvalidPostSizeValue"

	// 请填写AccessKeyId、SecretAccessKey作为第三方对象存储私有访问参数。
	INVALIDPARAMETER_INVALIDPRIVATEACCESSPARAMS = "InvalidParameter.InvalidPrivateAccessParams"

	// 请指定PrivateAccess参数值为on/off。
	INVALIDPARAMETER_INVALIDPRIVATEACCESSSWITCH = "InvalidParameter.InvalidPrivateAccessSwitch"

	// 套餐包不支持Quic配置。
	INVALIDPARAMETER_INVALIDQUICBILLING = "InvalidParameter.InvalidQuicBilling"

	// 配额数量格式不正确，请输入正确的配额数量格式。
	INVALIDPARAMETER_INVALIDQUOTANUMBER = "InvalidParameter.InvalidQuotaNumber"

	// 配额类型格式不正确，请输入正确的配额类型格式。
	INVALIDPARAMETER_INVALIDQUOTATYPE = "InvalidParameter.InvalidQuotaType"

	// 无效的分片回源。
	INVALIDPARAMETER_INVALIDRANGEORIGINPULL = "InvalidParameter.InvalidRangeOriginPull"

	// 自动续费标志格式不正确，请输入正确的自动续费标签格式。
	INVALIDPARAMETER_INVALIDRENEWFLAG = "InvalidParameter.InvalidRenewFlag"

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

	// 规则协议错误。（TCP/UDP）。
	INVALIDPARAMETER_INVALIDRULEPROTO = "InvalidParameter.InvalidRuleProto"

	// 无效的回源Host。
	INVALIDPARAMETER_INVALIDSERVERNAME = "InvalidParameter.InvalidServerName"

	// edgeone的debug配置无效。
	INVALIDPARAMETER_INVALIDSTANDARDDEBUG = "InvalidParameter.InvalidStandardDebug"

	// 无效的客户端ip或ip段。
	INVALIDPARAMETER_INVALIDSTANDARDDEBUGCLIENTIP = "InvalidParameter.InvalidStandardDebugClientIp"

	// AllowClientIPList 参数必填，支持 IPv4 及 IPv6 网段。0.0.0.0/0 表示允许所有 IPv4 客户端进行调试，::/0 表示允许所有 IPv6 客户端进行调试。
	INVALIDPARAMETER_INVALIDSTANDARDDEBUGEMPTYLIST = "InvalidParameter.InvalidStandardDebugEmptyList"

	// 有效期超过限制。
	INVALIDPARAMETER_INVALIDSTANDARDDEBUGEXPIRETIMELIMIT = "InvalidParameter.InvalidStandardDebugExpireTimeLimit"

	// 无效的回源请求参数设置-无效查询字符串值。
	INVALIDPARAMETER_INVALIDUPSTREAMREQUESTQUERYSTRINGVALUE = "InvalidParameter.InvalidUpstreamRequestQueryStringValue"

	// 无效的URL重写。
	INVALIDPARAMETER_INVALIDURLREDIRECT = "InvalidParameter.InvalidUrlRedirect"

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

	// 修改源站操作中负载均衡实例Id必填。
	INVALIDPARAMETER_LOADBALANCEINSTANCEIDISREQUIRED = "InvalidParameter.LoadBalanceInstanceIdIsRequired"

	// 修改参数缺失。
	INVALIDPARAMETER_MODIFYPARAMETERSMISSING = "InvalidParameter.ModifyParametersMissing"

	// 不支持智能路由
	INVALIDPARAMETER_MULTIPLYLAYERNOTSUPPORTSMARTROUTING = "InvalidParameter.MultiplyLayerNotSupportSmartRouting"

	// 操作配置存在不支持的预设变量。
	INVALIDPARAMETER_NOTSUPPORTTHISPRESET = "InvalidParameter.NotSupportThisPreset"

	// 域名处于直接回源架构，需要保持智能加速功能的开启。
	INVALIDPARAMETER_OCDIRECTORIGINREQUIRESSMARTROUTING = "InvalidParameter.OCDirectOriginRequiresSmartRouting"

	// 源站是内网IP。
	INVALIDPARAMETER_ORIGINISINNERIP = "InvalidParameter.OriginIsInnerIp"

	// 四层代理禁止IP域名混填。
	INVALIDPARAMETER_ORIGINL4RECORDIPV4MIXDOMAIN = "InvalidParameter.OriginL4RecordIPV4MixDomain"

	// 四层代理禁止使用多域名源站。
	INVALIDPARAMETER_ORIGINL4RECORDMULTIDOMAIN = "InvalidParameter.OriginL4RecordMultiDomain"

	// 源站组名称已经存在。
	INVALIDPARAMETER_ORIGINNAMEEXISTS = "InvalidParameter.OriginNameExists"

	// 修改源站操作中源站组Id必填。
	INVALIDPARAMETER_ORIGINORIGINGROUPIDISREQUIRED = "InvalidParameter.OriginOriginGroupIdIsRequired"

	// 修改源站操作中回源协议必填。
	INVALIDPARAMETER_ORIGINPULLPROTOCOLISREQUIRED = "InvalidParameter.OriginPullProtocolIsRequired"

	// 源站格式错误。
	INVALIDPARAMETER_ORIGINRECORDFORMATERROR = "InvalidParameter.OriginRecordFormatError"

	// 权重取值范围：0 - 100。
	INVALIDPARAMETER_ORIGINRECORDWEIGHTVALUE = "InvalidParameter.OriginRecordWeightValue"

	// 密钥格式错误。
	INVALIDPARAMETER_ORIGINTHIRDPARTYPARAMFORMATERROR = "InvalidParameter.OriginThirdPartyParamFormatError"

	// 参数错误: 无效 "结束时间", 不在允许的查询范围内: [开始时间, 开始+ 7天]
	INVALIDPARAMETER_PARAMETERERROR = "InvalidParameter.ParameterError"

	// 套餐不存在。
	INVALIDPARAMETER_PLANNOTFOUND = "InvalidParameter.PlanNotFound"

	// 最大上传大小超出限制
	INVALIDPARAMETER_POSTMAXSIZELIMITEXCEEDED = "InvalidParameter.PostMaxSizeLimitExceeded"

	// 实例名称重复。
	INVALIDPARAMETER_PROXYNAMEDUPLICATING = "InvalidParameter.ProxyNameDuplicating"

	// 实例名称可输入1-50个字符，允许的字符为a-z、0-9、-，- 不能单独注册或连续使用，不能放在开头或结尾。
	INVALIDPARAMETER_PROXYNAMENOTMATCHED = "InvalidParameter.ProxyNameNotMatched"

	// 推送实例已经创建
	INVALIDPARAMETER_REALTIMELOGENTITYALREADYCREATED = "InvalidParameter.RealtimeLogEntityAlreadyCreated"

	// 日志推送地域不合法
	INVALIDPARAMETER_REALTIMELOGINVALIDDELIVERYAREA = "InvalidParameter.RealtimeLogInvalidDeliveryArea"

	// 日志推送类型不合法
	INVALIDPARAMETER_REALTIMELOGINVALIDLOGTYPE = "InvalidParameter.RealtimeLogInvalidLogType"

	// 实时日志数据投递类型不合法
	INVALIDPARAMETER_REALTIMELOGINVALIDTASKTYPE = "InvalidParameter.RealtimeLogInvalidTaskType"

	// 实时日志推送任务数据超过了限制
	INVALIDPARAMETER_REALTIMELOGNUMSEXCEEDLIMIT = "InvalidParameter.RealtimeLogNumsExceedLimit"

	// 无效的响应头header。
	INVALIDPARAMETER_RESPONSEHEADERCACHECONTROLNOTALLOWDELETE = "InvalidParameter.ResponseHeaderCacheControlNotAllowDelete"

	// 规则源站信息格式错误。
	INVALIDPARAMETER_RULEORIGINFORMATERROR = "InvalidParameter.RuleOriginFormatError"

	// 规则源站不支持多域名。
	INVALIDPARAMETER_RULEORIGINMULTIDOMAIN = "InvalidParameter.RuleOriginMultiDomain"

	// 规则源站端口应为整数。
	INVALIDPARAMETER_RULEORIGINPORTINTEGER = "InvalidParameter.RuleOriginPortInteger"

	// 规则源站错误。
	INVALIDPARAMETER_RULEORIGINVALUEERROR = "InvalidParameter.RuleOriginValueError"

	// 规则端口重复。
	INVALIDPARAMETER_RULEPORTDUPLICATING = "InvalidParameter.RulePortDuplicating"

	// 规则端口段错误。
	INVALIDPARAMETER_RULEPORTGROUP = "InvalidParameter.RulePortGroup"

	// 规则端口必须为整数。
	INVALIDPARAMETER_RULEPORTINTEGER = "InvalidParameter.RulePortInteger"

	// 参数错误
	INVALIDPARAMETER_SECURITY = "InvalidParameter.Security"

	// 配置项参数错误。
	INVALIDPARAMETER_SETTINGINVALIDPARAM = "InvalidParameter.SettingInvalidParam"

	// Shield Space 未绑定源站，请先绑定源站后重试。
	INVALIDPARAMETER_SPACENOTBINDORIGIN = "InvalidParameter.SpaceNotBindOrigin"

	// 状态码缓存的状态码无效。
	INVALIDPARAMETER_STATUSCODECACHEINVALIDSTATUSCODE = "InvalidParameter.StatusCodeCacheInvalidStatusCode"

	// 资源存在错误。
	INVALIDPARAMETER_TARGET = "InvalidParameter.Target"

	// 任务无法生成。
	INVALIDPARAMETER_TASKNOTGENERATED = "InvalidParameter.TaskNotGenerated"

	// 内部错误。
	INVALIDPARAMETER_TASKSYSTEMERROR = "InvalidParameter.TaskSystemError"

	// HTTPS的TLS版本不连续。
	INVALIDPARAMETER_TLSVERSIONNOTINSEQUENCE = "InvalidParameter.TlsVersionNotInSequence"

	// 过滤值过多。
	INVALIDPARAMETER_TOOMANYFILTERVALUES = "InvalidParameter.TooManyFilterValues"

	// 过滤项过多。
	INVALIDPARAMETER_TOOMANYFILTERS = "InvalidParameter.TooManyFilters"

	// 文件上传链接存在问题。
	INVALIDPARAMETER_UPLOADURL = "InvalidParameter.UploadUrl"

	// 站点已被绑定。
	INVALIDPARAMETER_ZONEHASBEENBOUND = "InvalidParameter.ZoneHasBeenBound"

	// 站点未绑定套餐。
	INVALIDPARAMETER_ZONEHASNOTBEENBOUNDTOPLAN = "InvalidParameter.ZoneHasNotBeenBoundToPlan"

	// 站点升级中，暂不支持进行变更操作，请稍后再试。
	INVALIDPARAMETER_ZONEISGRAYPUBLISHING = "InvalidParameter.ZoneIsGrayPublishing"

	// 无域名接入站点切换到CNAME接入类型时站点名称必传。
	INVALIDPARAMETER_ZONENAMEISREQUIRED = "InvalidParameter.ZoneNameIsRequired"

	// 站点不存在。
	INVALIDPARAMETER_ZONENOTFOUND = "InvalidParameter.ZoneNotFound"

	// 参数取值错误。
	INVALIDPARAMETERVALUE = "InvalidParameterValue"

	// 该站点域名已被禁用。
	INVALIDPARAMETERVALUE_ACCESSBLACKLIST = "InvalidParameterValue.AccessBlacklist"

	// 边缘双向认证配置中的客户端证书必须是CA证书。
	INVALIDPARAMETERVALUE_CERTIFICATEVERIFYCLIENTMUSTCA = "InvalidParameterValue.CertificateVerifyClientMustCa"

	// 边缘双向认证配置至少需要配置一本证书。
	INVALIDPARAMETERVALUE_CERTIFICATEVERIFYCLIENTNEEDCERT = "InvalidParameterValue.CertificateVerifyClientNeedCert"

	// 边缘双向认证配置中的客户端 CA 证书最多允许配置20本。
	INVALIDPARAMETERVALUE_CLIENTCERTINFOQUOTALIMIT = "InvalidParameterValue.ClientCertInfoQuotaLimit"

	// 与已经添加的记录冲突。
	INVALIDPARAMETERVALUE_CONFLICTRECORD = "InvalidParameterValue.ConflictRecord"

	// DNS 记录与 DNSSEC 功能冲突。
	INVALIDPARAMETERVALUE_CONFLICTWITHDNSSEC = "InvalidParameterValue.ConflictWithDNSSEC"

	// DNS 记录与 NS 记录冲突。
	INVALIDPARAMETERVALUE_CONFLICTWITHNSRECORD = "InvalidParameterValue.ConflictWithNSRecord"

	// 主机记录与记录值不能取值相同。
	INVALIDPARAMETERVALUE_CONTENTSAMEASNAME = "InvalidParameterValue.ContentSameAsName"

	// 页面内容与Content-Type不匹配。
	INVALIDPARAMETERVALUE_CONTENTTYPENOTMATCH = "InvalidParameterValue.ContentTypeNotMatch"

	// 入参中的域名与站点参数不匹配，请更正后重试。
	INVALIDPARAMETERVALUE_DOMAINNOTMATCHZONE = "InvalidParameterValue.DomainNotMatchZone"

	// 格式不匹配。
	INVALIDPARAMETERVALUE_FORMATMISMATCH = "InvalidParameterValue.FormatMismatch"

	// 配置参数格式不匹配。
	INVALIDPARAMETERVALUE_GENERALMISMATCH = "InvalidParameterValue.GeneralMismatch"

	// 包含无效的值。
	INVALIDPARAMETERVALUE_INCLUDEINVALIDVALUE = "InvalidParameterValue.IncludeInvalidValue"

	// 别称域名名称不合法，别称域名应该由数字、英文字母、连词符组成，且连词符不能位于开头和结尾处。
	INVALIDPARAMETERVALUE_INVALIDALIASDOMAINNAME = "InvalidParameterValue.InvalidAliasDomainName"

	// 无效的别称域名后缀（该域名为内部接入域名，暂不支持作为别称域名接入）。
	INVALIDPARAMETERVALUE_INVALIDALIASNAMESUFFIX = "InvalidParameterValue.InvalidAliasNameSuffix"

	// DNS 记录内容错误。
	INVALIDPARAMETERVALUE_INVALIDDNSCONTENT = "InvalidParameterValue.InvalidDNSContent"

	// DNS 记录名称错误。
	INVALIDPARAMETERVALUE_INVALIDDNSNAME = "InvalidParameterValue.InvalidDNSName"

	// 加速域名名称不合法，加速域名应该由数字、英文字母、连词符组成，且连词符不能位于开头和结尾处。
	INVALIDPARAMETERVALUE_INVALIDDOMAINNAME = "InvalidParameterValue.InvalidDomainName"

	// 加速域名状态不符合要求。
	INVALIDPARAMETERVALUE_INVALIDDOMAINSTATUS = "InvalidParameterValue.InvalidDomainStatus"

	// DNS 代理域名源站错误。
	INVALIDPARAMETERVALUE_INVALIDPROXYORIGIN = "InvalidParameterValue.InvalidProxyOrigin"

	// 标签值存在不合法字符。
	INVALIDPARAMETERVALUE_INVALIDTAGVALUE = "InvalidParameterValue.InvalidTagValue"

	// 缺少必要配置参数。
	INVALIDPARAMETERVALUE_MISSINGNECESSARYPARAM = "InvalidParameterValue.MissingNecessaryParam"

	// 不支持接入泛域名 CNAME
	INVALIDPARAMETERVALUE_NOTALLOWEDWILDCARDSHAREDCNAME = "InvalidParameterValue.NotAllowedWildcardSharedCNAME"

	// 值不在要求的集合内。
	INVALIDPARAMETERVALUE_NOTINENUMERATION = "InvalidParameterValue.NotInEnumeration"

	// 值不在指定范围。
	INVALIDPARAMETERVALUE_NOTWITHINRANGE = "InvalidParameterValue.NotWithinRange"

	// 指定的源站组不存在。
	INVALIDPARAMETERVALUE_ORIGINGROUPNOTEXISTS = "InvalidParameterValue.OriginGroupNotExists"

	// 页面名称已存在。
	INVALIDPARAMETERVALUE_PAGENAMEALREADYEXIST = "InvalidParameterValue.PageNameAlreadyExist"

	// 不符合指定的正则表达式。
	INVALIDPARAMETERVALUE_REGEXMISMATCH = "InvalidParameterValue.RegExMismatch"

	// 已开启双向认证，客户端使用 RSA 或者 ECC 算法证书时，HTTPS 证书必须也配置有相同算法证书。
	INVALIDPARAMETERVALUE_SERVERCERTINFONEEDCONTAINRSAORECC = "InvalidParameterValue.ServerCertInfoNeedContainRSAorECC"

	// 已开启双向认证，客户端使用国密算法 CA 证书时，HTTPS 证书必须也配置有国密证书。
	INVALIDPARAMETERVALUE_SERVERCERTINFONEEDCONTAINSM2 = "InvalidParameterValue.ServerCertInfoNeedContainSM2"

	// 请输入合法的共享 CNAME 前缀，最大支持50个字符。
	INVALIDPARAMETERVALUE_SHAREDCNAMEPREFIXNOTMATCH = "InvalidParameterValue.SharedCNAMEPrefixNotMatch"

	// 暂不支持当前域名后缀接入，如您需要使用请联系我们。
	INVALIDPARAMETERVALUE_TOPLEVELDOMAINNOTSUPPORT = "InvalidParameterValue.TopLevelDomainNotSupport"

	// 配置项错误。
	INVALIDPARAMETERVALUE_UNRECOGNIZABLEVALUE = "InvalidParameterValue.UnrecognizableValue"

	// 站点名称格式不正确，请输入正确的域名格式。
	INVALIDPARAMETERVALUE_ZONENAMEINVALID = "InvalidParameterValue.ZoneNameInvalid"

	// 暂不支持 punycode 接入。
	INVALIDPARAMETERVALUE_ZONENAMENOTSUPPORTPUNYCODE = "InvalidParameterValue.ZoneNameNotSupportPunyCode"

	// 站点不支持以子域名接入，请以二级域名作为站点接入。
	INVALIDPARAMETERVALUE_ZONENAMENOTSUPPORTSUBDOMAIN = "InvalidParameterValue.ZoneNameNotSupportSubDomain"

	// 该同名站点标识已被占用，请重新输入。
	INVALIDPARAMETERVALUE_ZONESAMEASNAME = "InvalidParameterValue.ZoneSameAsName"

	// 超过配额限制。
	LIMITEXCEEDED = "LimitExceeded"

	// 本次提交的资源数超过上限。
	LIMITEXCEEDED_BATCHQUOTA = "LimitExceeded.BatchQuota"

	// 实时日志自定义字段正则类型字段数量超过限制
	LIMITEXCEEDED_CUSTOMLOGFIELDREGEXLIMITEXCEEDED = "LimitExceeded.CustomLogFieldRegexLimitExceeded"

	// 当天提交的资源数超过上限。
	LIMITEXCEEDED_DAILYQUOTA = "LimitExceeded.DailyQuota"

	// 函数数量达到限制。
	LIMITEXCEEDED_FUNCTIONLIMITEXCEEDED = "LimitExceeded.FunctionLimitExceeded"

	// 计费套餐不支持。
	LIMITEXCEEDED_PACKNOTALLOW = "LimitExceeded.PackNotAllow"

	// 导入的规则数量超限制。
	LIMITEXCEEDED_PROXYRULESLIMITEXCEEDED = "LimitExceeded.ProxyRulesLimitExceeded"

	// 查询时间范围超出限制。
	LIMITEXCEEDED_QUERYTIMELIMITEXCEEDED = "LimitExceeded.QueryTimeLimitExceeded"

	// 单位时间内接口请求频率达到限制。
	LIMITEXCEEDED_RATELIMITEXCEEDED = "LimitExceeded.RateLimitExceeded"

	// 规则数量达到限制。
	LIMITEXCEEDED_RULELIMITEXCEEDED = "LimitExceeded.RuleLimitExceeded"

	// 超出功能限制。
	LIMITEXCEEDED_SECURITY = "LimitExceeded.Security"

	// 用户实例数量限制。
	LIMITEXCEEDED_USERQUOTALIMITED = "LimitExceeded.UserQuotaLimited"

	// 套餐可绑定的站点数量超过配额。
	LIMITEXCEEDED_ZONEBINDPLAN = "LimitExceeded.ZoneBindPlan"

	// 缺少参数错误。
	MISSINGPARAMETER = "MissingParameter"

	// 操作被拒绝。
	OPERATIONDENIED = "OperationDenied"

	// 功能内测中，请联系商务开启「中国大陆网络优化(国际加速)」功能。
	OPERATIONDENIED_ACCELERATEMAINLANDDISABLE = "OperationDenied.AccelerateMainlandDisable"

	// 中国大陆加速与IPv6冲突，不能同时配置。
	OPERATIONDENIED_ACCELERATEMAINLANDIPV6CONFLICT = "OperationDenied.AccelerateMainlandIpv6Conflict"

	// 当前站点下存在域名功能与中国大陆网络优化冲突，不能同时配置。
	OPERATIONDENIED_ACCELERATEMAINLANDMULTIPLYLAYERCONFLICT = "OperationDenied.AccelerateMainlandMultiplyLayerConflict"

	// 引用源站组的7层域名服务正在部署中，请稍后再编辑。
	OPERATIONDENIED_ACCELERATIONDOMAINSTATUSNOTINONLINE = "OperationDenied.AccelerationDomainStatusNotInOnline"

	// 合规封禁中。
	OPERATIONDENIED_COMPLIANCEFORBIDDEN = "OperationDenied.ComplianceForbidden"

	// 配置已被锁定，请解除配置锁定之后在重试。
	OPERATIONDENIED_CONFIGLOCKED = "OperationDenied.ConfigLocked"

	// TEO_QCSLinkedRoleInDnspodAccessEO 角色未进行授权，请授权后重试。
	OPERATIONDENIED_DNSPODUNAUTHORIZEDROLEOPERATION = "OperationDenied.DNSPodUnauthorizedRoleOperation"

	// 删除站点时预检查未通过。
	OPERATIONDENIED_DELETEZONEPRECHECKFAILED = "OperationDenied.DeleteZonePreCheckFailed"

	// 站点停用未完毕，请稍后再试。
	OPERATIONDENIED_DISABLEZONENOTCOMPLETED = "OperationDenied.DisableZoneNotCompleted"

	// 有域名在共享cname组内，不可切换接入类型。
	OPERATIONDENIED_DOMAININSHARECNAMEGROUP = "OperationDenied.DomainInShareCnameGroup"

	// 域名被封禁，暂时无法操作。
	OPERATIONDENIED_DOMAINISBLOCKED = "OperationDenied.DomainIsBlocked"

	// 域名尚未备案。
	OPERATIONDENIED_DOMAINNOICP = "OperationDenied.DomainNoICP"

	// 站点下有域名时不允许修改服务区域。
	OPERATIONDENIED_DOMAINNUMBERISNOTZERO = "OperationDenied.DomainNumberIsNotZero"

	// 站点内有域名处于非稳态，稳态包括：“在线” 和 “离线” 两个状态。
	OPERATIONDENIED_DOMAINSTATUSUNSTABLE = "OperationDenied.DomainStatusUnstable"

	// 企业版套餐不支持自动续费。
	OPERATIONDENIED_ENTERPRISEPLANAUTORENEWUNSUPPORTED = "OperationDenied.EnterprisePlanAutoRenewUnsupported"

	// 企业版套餐不支持续费。
	OPERATIONDENIED_ENTERPRISEPLANRENEWUNSUPPORTED = "OperationDenied.EnterprisePlanRenewUnsupported"

	// 企业版套餐不支持升级。
	OPERATIONDENIED_ENTERPRISEPLANUPGRADEUNSUPPORTED = "OperationDenied.EnterprisePlanUpgradeUnsupported"

	// 站点环境未准备好。
	OPERATIONDENIED_ENVNOTREADY = "OperationDenied.EnvNotReady"

	// 站点处于停用状态，请开启后重试。
	OPERATIONDENIED_ERRZONEISALREADYPAUSED = "OperationDenied.ErrZoneIsAlreadyPaused"

	// 开启高防时必须保证安全是开启状态。
	OPERATIONDENIED_INVALIDADVANCEDDEFENSESECURITYTYPE = "OperationDenied.InvalidAdvancedDefenseSecurityType"

	// 开启高防必须保证站点加速区域是国内。
	OPERATIONDENIED_INVALIDADVANCEDDEFENSEZONEAREA = "OperationDenied.InvalidAdvancedDefenseZoneArea"

	// 独立DDoS防护与IPv6冲突，不能同时配置。
	OPERATIONDENIED_IPV6ADVANCEDCONFLICT = "OperationDenied.Ipv6AdvancedConflict"

	// ipv6功能和固定ip无法同时开启。
	OPERATIONDENIED_IPV6STATICIPCONFLICT = "OperationDenied.Ipv6StaticIpConflict"

	// 四层实例资源售卖火爆，已售罄，正在加紧补货中，当前无法新增四层代理，请您耐心等待。
	OPERATIONDENIED_L4LACKOFRESOURCES = "OperationDenied.L4LackOfResources"

	// 四层端口资源售卖火爆，已售罄，正在加紧补货中，当前无法新增四层代理或添加规则，请您耐心等待。
	OPERATIONDENIED_L4PORTLACKOFRESOURCES = "OperationDenied.L4PortLackOfResources"

	// 4层代理资源处于封禁中，禁止操作。
	OPERATIONDENIED_L4PROXYINBANNEDSTATUS = "OperationDenied.L4ProxyInBannedStatus"

	// 4层通道处于关闭状态，不允许添加规则。
	OPERATIONDENIED_L4PROXYINOFFLINESTATUS = "OperationDenied.L4ProxyInOfflineStatus"

	// 实例处于部署中，无法操作。
	OPERATIONDENIED_L4PROXYINPROCESSSTATUS = "OperationDenied.L4ProxyInProcessStatus"

	// 存在四层代理实例处于部署中状态，暂不支持停用站点。
	OPERATIONDENIED_L4PROXYINPROGRESSSTATUS = "OperationDenied.L4ProxyInProgressStatus"

	// 存在四层代理实例处于停用中状态，暂不支持停用站点。
	OPERATIONDENIED_L4PROXYINSTOPPINGSTATUS = "OperationDenied.L4ProxyInStoppingStatus"

	// 绑定4层实例有处于非运行中的状态，禁止操作。
	OPERATIONDENIED_L4STATUSNOTINONLINE = "OperationDenied.L4StatusNotInOnline"

	// 存在加速域名处于部署中状态，暂不支持停用站点。
	OPERATIONDENIED_L7HOSTINPROCESSSTATUS = "OperationDenied.L7HostInProcessStatus"

	// 绑定负载均衡有处于非运行中的状态，禁止操作。
	OPERATIONDENIED_LOADBALANCESTATUSNOTINONLINE = "OperationDenied.LoadBalanceStatusNotInOnline"

	// 站点状态不支持操作负载均衡。
	OPERATIONDENIED_LOADBALANCINGZONEISNOTACTIVE = "OperationDenied.LoadBalancingZoneIsNotActive"

	// 非海外独立防护，不能开启ipv6。
	OPERATIONDENIED_MSGIPV6ADVANCEDCONFLICT = "OperationDenied.MsgIpv6AdvancedConflict"

	// 已存在多个Cname接入站点，不允许切换至NS。
	OPERATIONDENIED_MULTIPLECNAMEZONE = "OperationDenied.MultipleCnameZone"

	// NS接入模式不支持域名流量调度功能。
	OPERATIONDENIED_NSNOTALLOWTRAFFICSTRATEGY = "OperationDenied.NSNotAllowTrafficStrategy"

	// 无域名接入站点仅可以切换到CNAME接入类型。
	OPERATIONDENIED_NODOMAINACCESSZONEONLYALLOWMODIFIEDTOCNAME = "OperationDenied.NoDomainAccessZoneOnlyAllowModifiedToCNAME"

	// 无域名接入站点不支持除切换为CNAME接入类型以外的任何修改动作。
	OPERATIONDENIED_NODOMAINACCESSZONEONLYSUPPORTMODIFYTYPE = "OperationDenied.NoDomainAccessZoneOnlySupportModifyType"

	// 用户不在版本管理的白名单内。
	OPERATIONDENIED_NOTINVERSIONCONTROLWHITELIST = "OperationDenied.NotInVersionControlWhiteList"

	// 加速域名使用中，无法删除。
	OPERATIONDENIED_ORIGINGROUPACCELERATIONDOMAINUSED = "OperationDenied.OriginGroupAccelerationDomainUsed"

	// 四层代理使用中，无法删除。
	OPERATIONDENIED_ORIGINGROUPL4USED = "OperationDenied.OriginGroupL4Used"

	// 负载均衡使用中，无法删除。
	OPERATIONDENIED_ORIGINGROUPLBUSED = "OperationDenied.OriginGroupLBUsed"

	// 规则引擎使用中，无法删除。
	OPERATIONDENIED_ORIGINGROUPRULEENGINEUSED = "OperationDenied.OriginGroupRuleEngineUsed"

	// 归属权校验未通过，请先完成站点归属权校验。
	OPERATIONDENIED_OWNERSHIPVERIFICATIONNOTPASSED = "OperationDenied.OwnershipVerificationNotPassed"

	// 套餐不支持降级。
	OPERATIONDENIED_PLANDOWNGRADENOTALLOWED = "OperationDenied.PlanDowngradeNotAllowed"

	// 套餐已过期。
	OPERATIONDENIED_PLANHASBEENEXPIRED = "OperationDenied.PlanHasBeenExpired"

	// 套餐已被隔离。
	OPERATIONDENIED_PLANHASBEENISOLATED = "OperationDenied.PlanHasBeenIsolated"

	// 该套餐不支持增购套餐配额。
	OPERATIONDENIED_PLANINCREASEPLANQUOTAUNSUPPORTED = "OperationDenied.PlanIncreasePlanQuotaUnsupported"

	// 套餐不支持变更站点的服务区域。
	OPERATIONDENIED_PLANNOTSUPPORTMODIFYZONEAREA = "OperationDenied.PlanNotSupportModifyZoneArea"

	// Anycast IP调度模式不支持开启中国大陆加速。
	OPERATIONDENIED_PLATTYPEIPACCELERATEMAINLANDNOTSUPPORT = "OperationDenied.PlatTypeIPAccelerateMainlandNotSupport"

	// 如需创建企业版套餐，请联系商务人员。
	OPERATIONDENIED_PLEASECONTACTBUSINESSPERSONNEL = "OperationDenied.PleaseContactBusinessPersonnel"

	// DNS 记录不允许添加。
	OPERATIONDENIED_RECORDISFORBIDDEN = "OperationDenied.RecordIsForbidden"

	// 计费资源操作中，请稍后重试。
	OPERATIONDENIED_RESOURCEHASBEENLOCKED = "OperationDenied.ResourceHasBeenLocked"

	// 当前有互相排斥的请求操作并行发起，请稍后重试。
	OPERATIONDENIED_RESOURCELOCKEDTEMPORARY = "OperationDenied.ResourceLockedTemporary"

	// 绑定在共享 CNAME 中的域名不允许变更 中国大陆网络优化（国际加速）访问，如果您需要单独变更，请先将域名从共享 CNAME 中解绑。
	OPERATIONDENIED_SHAREDCNAMEUNSUPPORTEDACCELERATEMAINLAND = "OperationDenied.SharedCNAMEUnsupportedAccelerateMainland"

	// 绑定在共享 CNAME 中的域名不允许变更 IPv6 访问，如果您需要单独变更，请先将域名从共享 CNAME 中解绑。
	OPERATIONDENIED_SHAREDCNAMEUNSUPPORTEDIPV6 = "OperationDenied.SharedCNAMEUnsupportedIPv6"

	// 该实例地域无法开启固定IP。
	OPERATIONDENIED_STATICIPAREACONFLICT = "OperationDenied.StaticIpAreaConflict"

	// 存在使用中的测试版本，请将测试版本发布现网或者回滚测试版本再重试。
	OPERATIONDENIED_VERSIONCONTROLISGRAYING = "OperationDenied.VersionControlIsGraying"

	// 当前处于版本管理模式下，该操作不允许。
	OPERATIONDENIED_VERSIONCONTROLLOCKED = "OperationDenied.VersionControlLocked"

	// 站点工作模式不属于版本管理模式。
	OPERATIONDENIED_WORKMODENOTINVERSIONCONTROL = "OperationDenied.WorkModeNotInVersionControl"

	// 共享CNAME已被其他站点绑定，请先解绑才能删除站点
	OPERATIONDENIED_ZONEISBINDINGSHAREDCNAME = "OperationDenied.ZoneIsBindingSharedCNAME"

	// 站点存在自定义错误页面引用，请先解绑。
	OPERATIONDENIED_ZONEISREFERENCECUSTOMERRORPAGE = "OperationDenied.ZoneIsReferenceCustomErrorPage"

	// 站点未启用。
	OPERATIONDENIED_ZONENOTACTIVE = "OperationDenied.ZoneNotActive"

	// 资源被占用。
	RESOURCEINUSE = "ResourceInUse"

	// 资源被本账号别称域名占用。
	RESOURCEINUSE_ALIASDOMAIN = "ResourceInUse.AliasDomain"

	// 当前已存在相同的别称域名，不支持重复添加
	RESOURCEINUSE_ALIASNAME = "ResourceInUse.AliasName"

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

	// 该共享 CNAME 已绑定加速域名，禁止删除。如您需要删除，请先将对应域名解绑。
	RESOURCEINUSE_SHAREDCNAME = "ResourceInUse.SharedCNAME"

	// 别称域名已站点接入。
	RESOURCEINUSE_ZONE = "ResourceInUse.Zone"

	// 资源不足。
	RESOURCEINSUFFICIENT = "ResourceInsufficient"

	// 资源不存在。
	RESOURCENOTFOUND = "ResourceNotFound"

	// DNSPod 未正常接入该域名，请在 DNSPod 接入后重试。
	RESOURCENOTFOUND_DNSPODDOMAINNOTINACCOUNT = "ResourceNotFound.DNSPodDomainNotInAccount"

	// 最大上传大小额度未配置
	RESOURCENOTFOUND_POSTMAXSIZEQUOTANOTFOUND = "ResourceNotFound.PostMaxSizeQuotaNotFound"

	// 配置组版本不存在，请检查后重试。
	RESOURCENOTFOUND_VERSIONNOTFOUND = "ResourceNotFound.VersionNotFound"

	// 资源不可用。
	RESOURCEUNAVAILABLE = "ResourceUnavailable"

	// 证书不存在或未授权。
	RESOURCEUNAVAILABLE_CERTNOTFOUND = "ResourceUnavailable.CertNotFound"

	// 当前域名已接入EdgeOne，如确认需要接入到当前账号，请进行域名找回。
	RESOURCEUNAVAILABLE_DOMAINALREADYEXISTS = "ResourceUnavailable.DomainAlreadyExists"

	// 请求的加速域名不存在，请更正后重试。
	RESOURCEUNAVAILABLE_DOMAINNOTFOUND = "ResourceUnavailable.DomainNotFound"

	// 函数不存在或不属于该账号。
	RESOURCEUNAVAILABLE_FUNCTIONNOTFOUND = "ResourceUnavailable.FunctionNotFound"

	// 域名不存在或未开启代理。
	RESOURCEUNAVAILABLE_HOSTNOTFOUND = "ResourceUnavailable.HostNotFound"

	// 规则不存在或不属于该账号。
	RESOURCEUNAVAILABLE_RULENOTFOUND = "ResourceUnavailable.RuleNotFound"

	// 该共享 CNAME 已被占用，请重新输入。
	RESOURCEUNAVAILABLE_SHAREDCNAMEALREADYEXISTS = "ResourceUnavailable.SharedCNAMEAlreadyExists"

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

	// 别称域名不支持目标域名源站类型为对象存储。
	UNSUPPORTEDOPERATION_TARGETNAMEORIGINTYPECOS = "UnsupportedOperation.TargetNameOriginTypeCos"
)
