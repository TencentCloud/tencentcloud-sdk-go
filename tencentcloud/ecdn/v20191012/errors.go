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

package v20191012

const (
	// 此产品的特有错误码

	// 域名配置更新操作失败，请重试或联系客服人员解决。
	FAILEDOPERATION_ECDNCONFIGERROR = "FailedOperation.EcdnConfigError"

	// CAM鉴权错误，请稍后重试。
	INTERNALERROR_CAMSYSTEMERROR = "InternalError.CamSystemError"

	// 内部配置服务错误，请重试或联系客服人员解决。
	INTERNALERROR_ECDNCONFIGERROR = "InternalError.EcdnConfigError"

	// 内部数据错误，请联系腾讯云工程师进一步排查。
	INTERNALERROR_ECDNDBERROR = "InternalError.EcdnDbError"

	// 系统内部错误，请联系客户人员解决或稍后重试。
	INTERNALERROR_ECDNQUERYSYSTEMERROR = "InternalError.EcdnQuerySystemError"

	// 系统错误，请联系腾讯云工程师进一步排查。
	INTERNALERROR_ECDNSYSTEMERROR = "InternalError.EcdnSystemError"

	// 内部服务错误，请联系腾讯云工程师进一步排查。
	INTERNALERROR_ERROR = "InternalError.Error"

	// 后端服务错误,请稍后重试 。
	INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"

	// 无法解析证书信息。
	INVALIDPARAMETER_ECDNCERTNOCERTINFO = "InvalidParameter.EcdnCertNoCertInfo"

	// 缓存配置不合法 。
	INVALIDPARAMETER_ECDNCONFIGINVALIDCACHE = "InvalidParameter.EcdnConfigInvalidCache"

	// 内部接口错误，请联系腾讯云工程师进一步排查。
	INVALIDPARAMETER_ECDNINTERFACEERROR = "InvalidParameter.EcdnInterfaceError"

	// 统计粒度不合法，请参考文档中统计分析示例。
	INVALIDPARAMETER_ECDNINVALIDPARAMINTERVAL = "InvalidParameter.EcdnInvalidParamInterval"

	// 参数错误，请参考文档中示例参数填充。
	INVALIDPARAMETER_ECDNPARAMERROR = "InvalidParameter.EcdnParamError"

	// 日期不合法，请参考文档中日期示例。
	INVALIDPARAMETER_ECDNSTATINVALIDDATE = "InvalidParameter.EcdnStatInvalidDate"

	// 统计类型不合法，请参考文档中统计分析示例。
	INVALIDPARAMETER_ECDNSTATINVALIDMETRIC = "InvalidParameter.EcdnStatInvalidMetric"

	// 参数错误。
	INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"

	// 域名操作过于频繁。
	LIMITEXCEEDED_ECDNDOMAINOPTOOOFTEN = "LimitExceeded.EcdnDomainOpTooOften"

	// 账号下无此域名，请确认后重试。
	RESOURCENOTFOUND_ECDNDOMAINNOTEXISTS = "ResourceNotFound.EcdnDomainNotExists"

	// 账号下无此域名，请确认后重试。
	RESOURCENOTFOUND_ECDNHOSTNOTEXISTS = "ResourceNotFound.EcdnHostNotExists"

	// 项目不存在。
	RESOURCENOTFOUND_ECDNPROJECTNOTEXISTS = "ResourceNotFound.EcdnProjectNotExists"

	// 未开通ECDN服务，请开通后使用此接口。
	RESOURCENOTFOUND_ECDNUSERNOTEXISTS = "ResourceNotFound.EcdnUserNotExists"

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

	// ECDN子账号加速域名未授权。
	UNAUTHORIZEDOPERATION_ECDNDOMAINUNAUTHORIZED = "UnauthorizedOperation.EcdnDomainUnauthorized"

	// ECDN子账号加速域名未授权。
	UNAUTHORIZEDOPERATION_ECDNHOSTUNAUTHORIZED = "UnauthorizedOperation.EcdnHostUnauthorized"

	// 子账号没有授权域名权限，请授权后重试。
	UNAUTHORIZEDOPERATION_ECDNNODOMAINUNAUTHORIZED = "UnauthorizedOperation.EcdnNoDomainUnauthorized"

	// 子账号项目未授权。
	UNAUTHORIZEDOPERATION_ECDNPROJECTUNAUTHORIZED = "UnauthorizedOperation.EcdnProjectUnauthorized"

	// 非内测白名单用户，无该功能使用权限。
	UNAUTHORIZEDOPERATION_ECDNUSERNOWHITELIST = "UnauthorizedOperation.EcdnUserNoWhitelist"

	// ECDN 子账号cam未授权。
	UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"

	// 操作过于频繁，请稍后重试 。
	UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"

	// ECDN 子账号项目未授权。
	UNAUTHORIZEDOPERATION_PROJECTNOPERMISSION = "UnauthorizedOperation.ProjectNoPermission"

	// ECDN 子账号项目未授权。
	UNAUTHORIZEDOPERATION_PROJECTSNOPERMISSION = "UnauthorizedOperation.ProjectsNoPermission"

	// 未知错误,请稍后重试 。
	UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
)
