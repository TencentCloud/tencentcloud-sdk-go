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

	// 创建日志集失败，请检查日志集名是否已存在。
	FAILEDOPERATION_CREATECLSLOGSETFAILED = "FailedOperation.CreateClsLogSetFailed"

	// 创建日志主题任务失败，请检查日志主题名或任务名是否已存在。
	FAILEDOPERATION_CREATECLSLOGTOPICTASKFAILED = "FailedOperation.CreateClsLogTopicTaskFailed"

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

	// 域名不存在或不属于该账号。
	INVALIDPARAMETER_DOMAINNOTFOUND = "InvalidParameter.DomainNotFound"

	// 当前域名已开启流量调度功能。
	INVALIDPARAMETER_DOMAINONTRAFFICSCHEDULING = "InvalidParameter.DomainOnTrafficScheduling"

	// 非法操作-非法参数。
	INVALIDPARAMETER_ERRINVALIDACTIONPARAM = "InvalidParameter.ErrInvalidActionParam"

	// 非法操作-非法参数-参数值数量超出限制。
	INVALIDPARAMETER_ERRINVALIDACTIONPARAMTOOMANYVALUES = "InvalidParameter.ErrInvalidActionParamTooManyValues"

	// 非法条件-非法参数名称-匹配类型不支持参数名称。
	INVALIDPARAMETER_ERRINVALIDCONDITIONNAMETARGETNOTSUPPORTNAME = "InvalidParameter.ErrInvalidConditionNameTargetNotSupportName"

	// 非法条件-非法参数值-无效的参数值。
	INVALIDPARAMETER_ERRINVALIDCONDITIONVALUEBADVALUE = "InvalidParameter.ErrInvalidConditionValueBadValue"

	// 非法条件-非法参数值-参数值长度超出限制。
	INVALIDPARAMETER_ERRINVALIDCONDITIONVALUETOOLONGVALUE = "InvalidParameter.ErrInvalidConditionValueTooLongValue"

	// 域名不存在。
	INVALIDPARAMETER_HOSTNOTFOUND = "InvalidParameter.HostNotFound"

	// 无效的token鉴权参数。
	INVALIDPARAMETER_INVALIDAUTHENTICATIONTYPESIGNPARAM = "InvalidParameter.InvalidAuthenticationTypeSignParam"

	// 无效的查询字符串。
	INVALIDPARAMETER_INVALIDCACHEKEYQUERYSTRINGVALUE = "InvalidParameter.InvalidCacheKeyQueryStringValue"

	// 无效的节点缓存。
	INVALIDPARAMETER_INVALIDCACHEONLYONSWITCH = "InvalidParameter.InvalidCacheOnlyOnSwitch"

	// 证书信息错误。
	INVALIDPARAMETER_INVALIDCERTINFO = "InvalidParameter.InvalidCertInfo"

	// 无效的客户端IP请求头。
	INVALIDPARAMETER_INVALIDCLIENTIPHEADERNAME = "InvalidParameter.InvalidClientIpHeaderName"

	// 无效的智能加速。
	INVALIDPARAMETER_INVALIDDYNAMICROUTINE = "InvalidParameter.InvalidDynamicRoutine"

	// 无效的自定义错误页面。
	INVALIDPARAMETER_INVALIDERRORPAGEREDIRECTURL = "InvalidParameter.InvalidErrorPageRedirectUrl"

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

	// 无效的请求头header。
	INVALIDPARAMETER_INVALIDREQUESTHEADERNAME = "InvalidParameter.InvalidRequestHeaderName"

	// 无套餐包。
	INVALIDPARAMETER_INVALIDRESOURCEIDBILLING = "InvalidParameter.InvalidResourceIdBilling"

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

	// URL重写的目标HOST无效。
	INVALIDPARAMETER_INVALIDURLREDIRECTHOST = "InvalidParameter.InvalidUrlRedirectHost"

	// URL重写的目标URL无效。
	INVALIDPARAMETER_INVALIDURLREDIRECTURL = "InvalidParameter.InvalidUrlRedirectUrl"

	// 无效的WebSocket。
	INVALIDPARAMETER_INVALIDWEBSOCKETTIMEOUT = "InvalidParameter.InvalidWebSocketTimeout"

	// 无效的缓存键。
	INVALIDPARAMETER_KEYRULESINVALIDQUERYSTRINGVALUE = "InvalidParameter.KeyRulesInvalidQueryStringValue"

	// 参数错误。
	INVALIDPARAMETER_PARAMETERERROR = "InvalidParameter.ParameterError"

	// 参数错误
	INVALIDPARAMETER_SECURITY = "InvalidParameter.Security"

	// 参数错误-setting非法参数。
	INVALIDPARAMETER_SETTINGINVALIDPARAM = "InvalidParameter.SettingInvalidParam"

	// 资源存在错误。
	INVALIDPARAMETER_TARGET = "InvalidParameter.Target"

	// 任务无法生成。
	INVALIDPARAMETER_TASKNOTGENERATED = "InvalidParameter.TaskNotGenerated"

	// 文件上传链接存在问题。
	INVALIDPARAMETER_UPLOADURL = "InvalidParameter.UploadUrl"

	// 站点不存在。
	INVALIDPARAMETER_ZONENOTFOUND = "InvalidParameter.ZoneNotFound"

	// 与已经添加的记录冲突。
	INVALIDPARAMETERVALUE_CONFLICTRECORD = "InvalidParameterValue.ConflictRecord"

	// DNS 记录与 DNSSEC 功能冲突。
	INVALIDPARAMETERVALUE_CONFLICTWITHDNSSEC = "InvalidParameterValue.ConflictWithDNSSEC"

	// DNS 记录与 LB 记录冲突。
	INVALIDPARAMETERVALUE_CONFLICTWITHLBRECORD = "InvalidParameterValue.ConflictWithLBRecord"

	// DNS 记录与 NS 记录冲突。
	INVALIDPARAMETERVALUE_CONFLICTWITHNSRECORD = "InvalidParameterValue.ConflictWithNSRecord"

	// DNS 记录内容错误。
	INVALIDPARAMETERVALUE_INVALIDDNSCONTENT = "InvalidParameterValue.InvalidDNSContent"

	// DNS 记录名称错误。
	INVALIDPARAMETERVALUE_INVALIDDNSNAME = "InvalidParameterValue.InvalidDNSName"

	// DNS 代理域名错误。
	INVALIDPARAMETERVALUE_INVALIDPROXYNAME = "InvalidParameterValue.InvalidProxyName"

	// DNS 代理域名源站错误。
	INVALIDPARAMETERVALUE_INVALIDPROXYORIGIN = "InvalidParameterValue.InvalidProxyOrigin"

	// 记录已存在。
	INVALIDPARAMETERVALUE_RECORDALREADYEXISTS = "InvalidParameterValue.RecordAlreadyExists"

	// 记录不允许添加。
	INVALIDPARAMETERVALUE_RECORDNOTALLOWED = "InvalidParameterValue.RecordNotAllowed"

	// 超过配额限制。
	LIMITEXCEEDED = "LimitExceeded"

	// 本次提交的资源数超过上限。
	LIMITEXCEEDED_BATCHQUOTA = "LimitExceeded.BatchQuota"

	// 当天提交的资源数超过上限。
	LIMITEXCEEDED_DAILYQUOTA = "LimitExceeded.DailyQuota"

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

	// 已存在多个Cname接入站点，不允许切换至NS。
	OPERATIONDENIED_MULTIPLECNAMEZONE = "OperationDenied.MultipleCnameZone"

	// NS接入模式不支持域名流量调度功能。
	OPERATIONDENIED_NSNOTALLOWTRAFFICSTRATEGY = "OperationDenied.NSNotAllowTrafficStrategy"

	// 资源被占用。
	RESOURCEINUSE = "ResourceInUse"

	// 资源被本账号别称域名占用。
	RESOURCEINUSE_ALIASDOMAIN = "ResourceInUse.AliasDomain"

	// 资源被本账号Cname接入占用。
	RESOURCEINUSE_CNAME = "ResourceInUse.Cname"

	// Dns资源被占用。
	RESOURCEINUSE_DNS = "ResourceInUse.Dns"

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

	// 资源不足。
	RESOURCEINSUFFICIENT = "ResourceInsufficient"

	// 资源不存在。
	RESOURCENOTFOUND = "ResourceNotFound"

	// 资源不可用。
	RESOURCEUNAVAILABLE = "ResourceUnavailable"

	// 无剩余可创建新任务的域名。
	RESOURCEUNAVAILABLE_AVAILABLEDOMAINNOTFOUND = "ResourceUnavailable.AvailableDomainNotFound"

	// 证书不存在或未授权。
	RESOURCEUNAVAILABLE_CERTNOTFOUND = "ResourceUnavailable.CertNotFound"

	// 域名不存在或未开启代理。
	RESOURCEUNAVAILABLE_HOSTNOTFOUND = "ResourceUnavailable.HostNotFound"

	// 未拉取到已开启代理的zone信息。
	RESOURCEUNAVAILABLE_PROXYZONENOTFOUND = "ResourceUnavailable.ProxyZoneNotFound"

	// 站点不存在或不属于该账号。
	RESOURCEUNAVAILABLE_ZONENOTFOUND = "ResourceUnavailable.ZoneNotFound"

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

	// 未知参数错误。
	UNKNOWNPARAMETER = "UnknownParameter"

	// 操作不支持。
	UNSUPPORTEDOPERATION = "UnsupportedOperation"
)
