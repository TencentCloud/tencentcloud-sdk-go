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

const (
	// 此产品的特有错误码

	// DryRun 操作，代表请求将会是成功的，只是多传了 DryRun 参数。
	DRYRUNOPERATION = "DryRunOperation"

	// 操作失败。
	FAILEDOPERATION = "FailedOperation"

	// 证书不存在。
	FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"

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

	// 无效的节点缓存。
	INVALIDPARAMETER_INVALIDCACHEONLYONSWITCH = "InvalidParameter.InvalidCacheOnlyOnSwitch"

	// 证书信息错误。
	INVALIDPARAMETER_INVALIDCERTINFO = "InvalidParameter.InvalidCertInfo"

	// 无效的客户端IP请求头。
	INVALIDPARAMETER_INVALIDCLIENTIPHEADERNAME = "InvalidParameter.InvalidClientIpHeaderName"

	// 套餐包不支持智能加速配置。
	INVALIDPARAMETER_INVALIDDYNAMICROUTINEBILLING = "InvalidParameter.InvalidDynamicRoutineBilling"

	// 不合法的过滤字段。
	INVALIDPARAMETER_INVALIDFILTERNAME = "InvalidParameter.InvalidFilterName"

	// 无效的强制HTTPS跳转。
	INVALIDPARAMETER_INVALIDFORCEREDIRECTTYPE = "InvalidParameter.InvalidForceRedirectType"

	// 无效的源站。
	INVALIDPARAMETER_INVALIDORIGIN = "InvalidParameter.InvalidOrigin"

	// 套餐包不支持最大上传大小。
	INVALIDPARAMETER_INVALIDPOSTMAXSIZEBILLING = "InvalidParameter.InvalidPostMaxSizeBilling"

	// 无效的最大上传大小。
	INVALIDPARAMETER_INVALIDPOSTSIZEVALUE = "InvalidParameter.InvalidPostSizeValue"

	// 无效的URL重写。
	INVALIDPARAMETER_INVALIDREDIRECTURLCAPTURE = "InvalidParameter.InvalidRedirectUrlCapture"

	// 无效的请求头header。
	INVALIDPARAMETER_INVALIDREQUESTHEADERNAME = "InvalidParameter.InvalidRequestHeaderName"

	// 无效的请求头header。
	INVALIDPARAMETER_INVALIDREQUESTHEADERVALUE = "InvalidParameter.InvalidRequestHeaderValue"

	// 无套餐包。
	INVALIDPARAMETER_INVALIDRESOURCEIDBILLING = "InvalidParameter.InvalidResourceIdBilling"

	// 无效的响应头header。
	INVALIDPARAMETER_INVALIDRESPONSEHEADERVALUE = "InvalidParameter.InvalidResponseHeaderValue"

	// 无效的规则引擎操作。
	INVALIDPARAMETER_INVALIDRULEENGINEACTION = "InvalidParameter.InvalidRuleEngineAction"

	// 无效的规则引擎条件。
	INVALIDPARAMETER_INVALIDRULEENGINETARGET = "InvalidParameter.InvalidRuleEngineTarget"

	// 无效的规则引擎文件后缀条件。
	INVALIDPARAMETER_INVALIDRULEENGINETARGETSEXTENSION = "InvalidParameter.InvalidRuleEngineTargetsExtension"

	// URL重写的目标URL无效。
	INVALIDPARAMETER_INVALIDURLREDIRECTURL = "InvalidParameter.InvalidUrlRedirectUrl"

	// 无效的WebSocket。
	INVALIDPARAMETER_INVALIDWEBSOCKETTIMEOUT = "InvalidParameter.InvalidWebSocketTimeout"

	// 参数错误。
	INVALIDPARAMETER_PARAMETERERROR = "InvalidParameter.ParameterError"

	// 参数错误
	INVALIDPARAMETER_SECURITY = "InvalidParameter.Security"

	// 配置项参数错误。
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

	// SRV 记录名称错误。
	INVALIDPARAMETERVALUE_INVALIDSRVNAME = "InvalidParameterValue.InvalidSRVName"

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

	// 操作被拒绝。
	OPERATIONDENIED = "OperationDenied"

	// 域名尚未备案。
	OPERATIONDENIED_DOMAINNOICP = "OperationDenied.DomainNoICP"

	// 资源被占用。
	RESOURCEINUSE = "ResourceInUse"

	// 资源被其他用户接入。
	RESOURCEINUSE_OTHERS = "ResourceInUse.Others"

	// 资源不足。
	RESOURCEINSUFFICIENT = "ResourceInsufficient"

	// 资源不存在。
	RESOURCENOTFOUND = "ResourceNotFound"

	// 资源不可用。
	RESOURCEUNAVAILABLE = "ResourceUnavailable"

	// 证书不存在或未授权。
	RESOURCEUNAVAILABLE_CERTNOTFOUND = "ResourceUnavailable.CertNotFound"

	// 域名不存在或未开启代理。
	RESOURCEUNAVAILABLE_HOSTNOTFOUND = "ResourceUnavailable.HostNotFound"

	// 站点不存在或不属于该账号。
	RESOURCEUNAVAILABLE_ZONENOTFOUND = "ResourceUnavailable.ZoneNotFound"

	// Cam 未授权。
	UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"

	// 鉴权错误。
	UNAUTHORIZEDOPERATION_DOMAINEMPTY = "UnauthorizedOperation.DomainEmpty"

	// 子账户没有操作权限，请添加权限后继续操作。
	UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
)
