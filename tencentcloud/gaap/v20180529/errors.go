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

package v20180529

const (
	// 此产品的特有错误码

	// 非企业认证客户。
	AUTHFAILURE_NOTENTERPRISEAUTHORIZATION = "AuthFailure.NotEnterpriseAuthorization"

	// 操作失败。
	FAILEDOPERATION = "FailedOperation"

	// 账户余额不足，无法创建该通道。
	FAILEDOPERATION_ACCOUNTBALANCEINSUFFICIENT = "FailedOperation.AccountBalanceInsufficient"

	// 操作正在执行中，请勿重复操作。
	FAILEDOPERATION_ACTIONISDOING = "FailedOperation.ActionIsDoing"

	// 接口操作太频繁，请稍后重试。
	FAILEDOPERATION_ACTIONOPERATETOOQUICKLY = "FailedOperation.ActionOperateTooQuickly"

	// 该批通道归属于不同的通道组，无法批量操作。
	FAILEDOPERATION_BELONGDIFFERENTGROUP = "FailedOperation.BelongDifferentGroup"

	// 证书正在使用中，无法操作。
	FAILEDOPERATION_CERTIFICATEISUSING = "FailedOperation.CertificateIsUsing"

	// 通道组中存在通道，无法删除。
	FAILEDOPERATION_DELETEPROXYGROUPPROXYREMAINED = "FailedOperation.DeleteProxyGroupProxyRemained"

	// 该域名在监听器下已使用。
	FAILEDOPERATION_DOMAINALREADYEXISTED = "FailedOperation.DomainAlreadyExisted"

	// 域名状态为非运行状态，无法操作。
	FAILEDOPERATION_DOMAINSTATUSNOTINRUNNING = "FailedOperation.DomainStatusNotInRunning"

	// 重复的请求，请检查ClientToken的值。
	FAILEDOPERATION_DUPLICATEDREQUEST = "FailedOperation.DuplicatedRequest"

	// 通道组状态为非运行状态，无法操作。
	FAILEDOPERATION_GROUPSTATUSNOTINRUNING = "FailedOperation.GroupStatusNotInRuning"

	// 通道状态为非运行状态，无法操作。
	FAILEDOPERATION_INSTANCESTATUSNOTINRUNING = "FailedOperation.InstanceStatusNotInRuning"

	// 无效的监听器协议。
	FAILEDOPERATION_INVALIDLISTENERPROTOCOL = "FailedOperation.InvalidListenerProtocol"

	// 监听器数量超过限制。
	FAILEDOPERATION_LIMITNUMOFLISTENER = "FailedOperation.LimitNumofListener"

	// 通道组下通道数量超过限制。
	FAILEDOPERATION_LIMITNUMOFPROXIESINGROUP = "FailedOperation.LimitNumofProxiesInGroup"

	// 转发规则数超过限制。
	FAILEDOPERATION_LIMITNUMOFRULES = "FailedOperation.LimitNumofRules"

	// 绑定源站数量超过限制。
	FAILEDOPERATION_LIMITREALSERVERNUM = "FailedOperation.LimitRealServerNum"

	// 监听器正在操作中，请勿重复操作。
	FAILEDOPERATION_LISTENERHASTASK = "FailedOperation.ListenerHasTask"

	// 监听器当前状态无法支持该操作。
	FAILEDOPERATION_LISTENERSTATUSERROR = "FailedOperation.ListenerStatusError"

	// 仅支持Version2.0的通道。
	FAILEDOPERATION_NOTSUPPORTOLDVERSIONPROXY = "FailedOperation.NotSupportOldVersionProxy"

	// 该操作不支持通道组。
	FAILEDOPERATION_NOTSUPPORTPROXYGROUP = "FailedOperation.NotSupportProxyGroup"

	// 不支持配置变更。
	FAILEDOPERATION_NOTSUPPORTSCALAR = "FailedOperation.NotSupportScalar"

	// 单次操作端口过多，超过数量上限。
	FAILEDOPERATION_OPERATELIMITNUMOFLISTENER = "FailedOperation.OperateLimitNumofListener"

	// 安全策略已关闭，请勿重复操作。
	FAILEDOPERATION_PROXYSECURITYALREADYCLOSE = "FailedOperation.ProxySecurityAlreadyClose"

	// 安全策略已开启，请勿重复操作。
	FAILEDOPERATION_PROXYSECURITYALREADYOPEN = "FailedOperation.ProxySecurityAlreadyOpen"

	// 添加规则失败：禁止拒绝默认的通道访问规则。
	FAILEDOPERATION_PROXYSECURITYPOLICYDEFAULTRULE = "FailedOperation.ProxySecurityPolicyDefaultRule"

	// 添加规则失败：安全防护访问规则重复。
	FAILEDOPERATION_PROXYSECURITYPOLICYDUPLICATEDRULE = "FailedOperation.ProxySecurityPolicyDuplicatedRule"

	// 安全策略已经存在，请勿重复创建。
	FAILEDOPERATION_PROXYSECURITYPOLICYEXISTED = "FailedOperation.ProxySecurityPolicyExisted"

	// 线路售罄或资源不足，请提工单申请。
	FAILEDOPERATION_PROXYSELLOUT = "FailedOperation.ProxySellOut"

	// 通道处于非运行状态，不能添加监听器。
	FAILEDOPERATION_PROXYSTATUSNOTINRUNING = "FailedOperation.ProxyStatusNotInRuning"

	// 不支持该版本通道。
	FAILEDOPERATION_PROXYVERSIONNOTSUPPORT = "FailedOperation.ProxyVersionNotSupport"

	// 已经绑定了源站，无法删除。
	FAILEDOPERATION_REALSERVERALREADYBOUND = "FailedOperation.RealServerAlreadyBound"

	// 源站不归属于该项目。
	FAILEDOPERATION_REALSERVERNOTINPROJECT = "FailedOperation.RealServerNotInProject"

	// 该资源不可访问。
	FAILEDOPERATION_RESOURCECANNOTACCESS = "FailedOperation.ResourceCanNotAccess"

	// 资源升级中。
	FAILEDOPERATION_RESOURCEUPGRADING = "FailedOperation.ResourceUpgrading"

	// 规则已经存在。
	FAILEDOPERATION_RULEALREADYEXISTED = "FailedOperation.RuleAlreadyExisted"

	// 用户未认证。
	FAILEDOPERATION_USERNOTAUTHENTICATED = "FailedOperation.UserNotAuthenticated"

	// 内部错误。
	INTERNALERROR = "InternalError"

	// 参数错误。
	INVALIDPARAMETER = "InvalidParameter"

	// 参数取值错误。
	INVALIDPARAMETERVALUE = "InvalidParameterValue"

	// Https证书与域名不匹配。
	INVALIDPARAMETERVALUE_CERTIFICATENOTMATCHDOMAIN = "InvalidParameterValue.CertificateNotMatchDomain"

	// 域名在ICP黑名单内。
	INVALIDPARAMETERVALUE_DOMAININICPBLACKLIST = "InvalidParameterValue.DomainInIcpBlacklist"

	// 域名未备案。
	INVALIDPARAMETERVALUE_DOMAINNOTREGISTER = "InvalidParameterValue.DomainNotRegister"

	// RealServer已存在。
	INVALIDPARAMETERVALUE_DUPLICATERS = "InvalidParameterValue.DuplicateRS"

	// 监听器端口已存在。
	INVALIDPARAMETERVALUE_DUPLICATEDLISTENERPORT = "InvalidParameterValue.DuplicatedListenerPort"

	// 无法同时开启所设置的特性。
	INVALIDPARAMETERVALUE_FEATURECONFLICT = "InvalidParameterValue.FeatureConflict"

	// header黑名单限制。
	INVALIDPARAMETERVALUE_HITBLACKLIST = "InvalidParameterValue.HitBlacklist"

	// 带宽值不在可选范围内。
	INVALIDPARAMETERVALUE_INVALIDBANDWIDTH = "InvalidParameterValue.InvalidBandwidth"

	// 解析失败，请检查证书内容格式。
	INVALIDPARAMETERVALUE_INVALIDCERTIFICATECONTENT = "InvalidParameterValue.InvalidCertificateContent"

	// 证书不可用。
	INVALIDPARAMETERVALUE_INVALIDCERTIFICATEID = "InvalidParameterValue.InvalidCertificateId"

	// 解析失败，请检查证书密钥格式。
	INVALIDPARAMETERVALUE_INVALIDCERTIFICATEKEY = "InvalidParameterValue.InvalidCertificateKey"

	// 并发量值不在可选范围内。
	INVALIDPARAMETERVALUE_INVALIDCONCURRENCY = "InvalidParameterValue.InvalidConcurrency"

	// 该监听器端口不可用。
	INVALIDPARAMETERVALUE_INVALIDLISTENERPORT = "InvalidParameterValue.InvalidListenerPort"

	// 未找到或无权限访问该标签
	INVALIDPARAMETERVALUE_INVALIDTAGS = "InvalidParameterValue.InvalidTags"

	// 项目不属于该用户。
	INVALIDPARAMETERVALUE_PROJECTIDNOTBELONG = "InvalidParameterValue.ProjectIdNotBelong"

	// 该通道组下无法支持该通道所需的特性。
	INVALIDPARAMETERVALUE_PROXYANDGROUPFEATURECONFLICT = "InvalidParameterValue.ProxyAndGroupFeatureConflict"

	// 该地区不支持通道所设置的特性。
	INVALIDPARAMETERVALUE_PROXYANDREGIONFEATURECONFLICT = "InvalidParameterValue.ProxyAndRegionFeatureConflict"

	// 源站不属于该用户。
	INVALIDPARAMETERVALUE_REALSERVERNOTBELONG = "InvalidParameterValue.RealServerNotBelong"

	// 未找到或无权限访问的加速区域。
	INVALIDPARAMETERVALUE_UNKNOWNACCESSREGION = "InvalidParameterValue.UnknownAccessRegion"

	// 未找到或无权限访问的源站区域。
	INVALIDPARAMETERVALUE_UNKNOWNDESTREGION = "InvalidParameterValue.UnknownDestRegion"

	// 超过配额限制。
	LIMITEXCEEDED = "LimitExceeded"

	// 缺少参数错误。
	MISSINGPARAMETER = "MissingParameter"

	// 资源不存在。
	RESOURCENOTFOUND = "ResourceNotFound"

	// 资源不可用。
	RESOURCEUNAVAILABLE = "ResourceUnavailable"

	// 账号下存在违规资源，详情请查看站内信或邮件。
	RESOURCEUNAVAILABLE_ACCOUNTVIOLATION = "ResourceUnavailable.AccountViolation"

	// 未授权操作。
	UNAUTHORIZEDOPERATION = "UnauthorizedOperation"

	// 跨境通道的联通带宽处于隔离中。
	UNAUTHORIZEDOPERATION_CROSSBORDERINISOLATING = "UnauthorizedOperation.CrossBorderInIsolating"

	// 未知参数错误。
	UNKNOWNPARAMETER = "UnknownParameter"

	// 操作不支持。
	UNSUPPORTEDOPERATION = "UnsupportedOperation"
)
