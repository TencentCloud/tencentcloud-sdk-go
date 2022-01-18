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

package v20191012

const (
	// 此产品的特有错误码

	// 域名配置更新操作失败，请重试或联系客服人员解决。
	FAILEDOPERATION_ECDNCONFIGERROR = "FailedOperation.EcdnConfigError"

	// 内部错误。
	INTERNALERROR = "InternalError"

	// 获取用户信息失败，请联系腾讯云工程师进一步排查。
	INTERNALERROR_ACCOUNTSYSTEMERROR = "InternalError.AccountSystemError"

	// 数据查询错误，请联系腾讯云工程师进一步排查。
	INTERNALERROR_DATASYSTEMERROR = "InternalError.DataSystemError"

	// 内部配置服务错误，请重试或联系客服人员解决。
	INTERNALERROR_ECDNCONFIGERROR = "InternalError.EcdnConfigError"

	// 内部数据错误，请联系腾讯云工程师进一步排查。
	INTERNALERROR_ECDNDBERROR = "InternalError.EcdnDbError"

	// 系统错误，请联系腾讯云工程师进一步排查。
	INTERNALERROR_ECDNSYSTEMERROR = "InternalError.EcdnSystemError"

	// 后端服务错误,请稍后重试 。
	INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"

	// 标签键不存在。
	INVALIDPARAMETER_ECDNCAMTAGKEYNOTEXIST = "InvalidParameter.EcdnCamTagKeyNotExist"

	// 无法解析证书信息。
	INVALIDPARAMETER_ECDNCERTNOCERTINFO = "InvalidParameter.EcdnCertNoCertInfo"

	// 缓存配置不合法 。
	INVALIDPARAMETER_ECDNCONFIGINVALIDCACHE = "InvalidParameter.EcdnConfigInvalidCache"

	// 域名状态不合法。
	INVALIDPARAMETER_ECDNDOMAININVALIDSTATUS = "InvalidParameter.EcdnDomainInvalidStatus"

	// 内部接口错误，请联系腾讯云工程师进一步排查。
	INVALIDPARAMETER_ECDNINTERFACEERROR = "InvalidParameter.EcdnInterfaceError"

	// 非法Area参数，请参考文档中示例参数填充。
	INVALIDPARAMETER_ECDNINVALIDPARAMAREA = "InvalidParameter.EcdnInvalidParamArea"

	// 统计粒度不合法，请参考文档中统计分析示例。
	INVALIDPARAMETER_ECDNINVALIDPARAMINTERVAL = "InvalidParameter.EcdnInvalidParamInterval"

	// 参数错误，请参考文档中示例参数填充。
	INVALIDPARAMETER_ECDNPARAMERROR = "InvalidParameter.EcdnParamError"

	// 刷新不支持泛域名。
	INVALIDPARAMETER_ECDNPURGEWILDCARDNOTALLOWED = "InvalidParameter.EcdnPurgeWildcardNotAllowed"

	// 该域名绑定的标签键数量过多。
	INVALIDPARAMETER_ECDNRESOURCEMANYTAGKEY = "InvalidParameter.EcdnResourceManyTagKey"

	// 日期不合法，请参考文档中日期示例。
	INVALIDPARAMETER_ECDNSTATINVALIDDATE = "InvalidParameter.EcdnStatInvalidDate"

	// 统计类型不合法，请参考文档中统计分析示例。
	INVALIDPARAMETER_ECDNSTATINVALIDMETRIC = "InvalidParameter.EcdnStatInvalidMetric"

	// 标签键不合法。
	INVALIDPARAMETER_ECDNTAGKEYINVALID = "InvalidParameter.EcdnTagKeyInvalid"

	// 标签键不存在。
	INVALIDPARAMETER_ECDNTAGKEYNOTEXIST = "InvalidParameter.EcdnTagKeyNotExist"

	// 标签键下的值数量过多。
	INVALIDPARAMETER_ECDNTAGKEYTOOMANYVALUE = "InvalidParameter.EcdnTagKeyTooManyValue"

	// 标签值不合法。
	INVALIDPARAMETER_ECDNTAGVALUEINVALID = "InvalidParameter.EcdnTagValueInvalid"

	// URL 超过限制长度。
	INVALIDPARAMETER_ECDNURLEXCEEDLENGTH = "InvalidParameter.EcdnUrlExceedLength"

	// 该用户下标签键数量过多。
	INVALIDPARAMETER_ECDNUSERTOOMANYTAGKEY = "InvalidParameter.EcdnUserTooManyTagKey"

	// 参数错误。
	INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"

	// 域名操作过于频繁。
	LIMITEXCEEDED_ECDNDOMAINOPTOOOFTEN = "LimitExceeded.EcdnDomainOpTooOften"

	// 刷新的目录数量超过单次限制。
	LIMITEXCEEDED_ECDNPURGEPATHEXCEEDBATCHLIMIT = "LimitExceeded.EcdnPurgePathExceedBatchLimit"

	// 刷新的目录数量超过每日限制。
	LIMITEXCEEDED_ECDNPURGEPATHEXCEEDDAYLIMIT = "LimitExceeded.EcdnPurgePathExceedDayLimit"

	// 刷新的Url数量超过单次限制。
	LIMITEXCEEDED_ECDNPURGEURLEXCEEDBATCHLIMIT = "LimitExceeded.EcdnPurgeUrlExceedBatchLimit"

	// 刷新的Url数量超过每日限额。
	LIMITEXCEEDED_ECDNPURGEURLEXCEEDDAYLIMIT = "LimitExceeded.EcdnPurgeUrlExceedDayLimit"

	// 接入域名数超出限制。
	LIMITEXCEEDED_ECDNUSERTOOMANYDOMAINS = "LimitExceeded.EcdnUserTooManyDomains"

	// 域名已存在。
	RESOURCEINUSE_ECDNDOMAINEXISTS = "ResourceInUse.EcdnDomainExists"

	// ECDN资源正在被操作中。
	RESOURCEINUSE_ECDNOPINPROGRESS = "ResourceInUse.EcdnOpInProgress"

	// 账号下无此域名，请确认后重试。
	RESOURCENOTFOUND_ECDNDOMAINNOTEXISTS = "ResourceNotFound.EcdnDomainNotExists"

	// 账号下无此域名，请确认后重试。
	RESOURCENOTFOUND_ECDNHOSTNOTEXISTS = "ResourceNotFound.EcdnHostNotExists"

	// 项目不存在。
	RESOURCENOTFOUND_ECDNPROJECTNOTEXISTS = "ResourceNotFound.EcdnProjectNotExists"

	// 未开通ECDN服务，请开通后使用此接口。
	RESOURCENOTFOUND_ECDNUSERNOTEXISTS = "ResourceNotFound.EcdnUserNotExists"

	// 域名已锁定,请联系腾讯云工程师 。
	RESOURCEUNAVAILABLE_ECDNDOMAINISLOCKED = "ResourceUnavailable.EcdnDomainIsLocked"

	// 域名未下线，请检查后重试。
	RESOURCEUNAVAILABLE_ECDNDOMAINISNOTOFFLINE = "ResourceUnavailable.EcdnDomainIsNotOffline"

	// 域名已下线，请检查后重试。
	RESOURCEUNAVAILABLE_ECDNDOMAINISNOTONLINE = "ResourceUnavailable.EcdnDomainIsNotOnline"

	// 子账号禁止查询整体数据。
	UNAUTHORIZEDOPERATION_CDNACCOUNTUNAUTHORIZED = "UnauthorizedOperation.CdnAccountUnauthorized"

	// 子账号未配置cam策略。
	UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"

	// ECDN子账号加速域名未授权。
	UNAUTHORIZEDOPERATION_CDNDOMAINUNAUTHORIZED = "UnauthorizedOperation.CdnDomainUnauthorized"

	// ECDN子账号加速域名未授权。
	UNAUTHORIZEDOPERATION_CDNHOSTUNAUTHORIZED = "UnauthorizedOperation.CdnHostUnauthorized"

	// 子账号没有授权域名权限，请授权后重试。
	UNAUTHORIZEDOPERATION_CDNNODOMAINUNAUTHORIZED = "UnauthorizedOperation.CdnNoDomainUnauthorized"

	// 子账号项目未授权。
	UNAUTHORIZEDOPERATION_CDNPROJECTUNAUTHORIZED = "UnauthorizedOperation.CdnProjectUnauthorized"

	// ECDN 子账号加速域名未授权。
	UNAUTHORIZEDOPERATION_DOMAINNOPERMISSION = "UnauthorizedOperation.DomainNoPermission"

	// ECDN 子账号加速域名未授权。
	UNAUTHORIZEDOPERATION_DOMAINSNOPERMISSION = "UnauthorizedOperation.DomainsNoPermission"

	// 子账号禁止查询整体数据。
	UNAUTHORIZEDOPERATION_ECDNACCOUNTUNAUTHORIZED = "UnauthorizedOperation.EcdnAccountUnauthorized"

	// 子账号未配置cam策略。
	UNAUTHORIZEDOPERATION_ECDNCAMUNAUTHORIZED = "UnauthorizedOperation.EcdnCamUnauthorized"

	// 域名解析未进行验证。
	UNAUTHORIZEDOPERATION_ECDNDOMAINRECORDNOTVERIFIED = "UnauthorizedOperation.EcdnDomainRecordNotVerified"

	// ECDN子账号加速域名未授权。
	UNAUTHORIZEDOPERATION_ECDNDOMAINUNAUTHORIZED = "UnauthorizedOperation.EcdnDomainUnauthorized"

	// 该域名属于其他账号，您没有权限接入。
	UNAUTHORIZEDOPERATION_ECDNHOSTISOWNEDBYOTHER = "UnauthorizedOperation.EcdnHostIsOwnedByOther"

	// ECDN子账号加速域名未授权。
	UNAUTHORIZEDOPERATION_ECDNHOSTUNAUTHORIZED = "UnauthorizedOperation.EcdnHostUnauthorized"

	// 请前往CDN控制台进行操作 。
	UNAUTHORIZEDOPERATION_ECDNMIGRATEDCDN = "UnauthorizedOperation.EcdnMigratedCdn"

	// 子账号没有授权域名权限，请授权后重试。
	UNAUTHORIZEDOPERATION_ECDNNODOMAINUNAUTHORIZED = "UnauthorizedOperation.EcdnNoDomainUnauthorized"

	// 子账号项目未授权。
	UNAUTHORIZEDOPERATION_ECDNPROJECTUNAUTHORIZED = "UnauthorizedOperation.EcdnProjectUnauthorized"

	// 加速服务已停服，请重启加速服务后重试。
	UNAUTHORIZEDOPERATION_ECDNUSERISSUSPENDED = "UnauthorizedOperation.EcdnUserIsSuspended"

	// 非内测白名单用户，无该功能使用权限。
	UNAUTHORIZEDOPERATION_ECDNUSERNOWHITELIST = "UnauthorizedOperation.EcdnUserNoWhitelist"

	// ECDN 子账号cam未授权。
	UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"

	// ECDN 子账号项目未授权。
	UNAUTHORIZEDOPERATION_PROJECTNOPERMISSION = "UnauthorizedOperation.ProjectNoPermission"

	// ECDN 子账号项目未授权。
	UNAUTHORIZEDOPERATION_PROJECTSNOPERMISSION = "UnauthorizedOperation.ProjectsNoPermission"

	// 未知错误,请稍后重试 。
	UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
)
