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

	// 内部错误。
	INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"

	// 参数错误。
	INVALIDPARAMETER = "InvalidParameter"

	// 域名不存在或不属于该账号。
	INVALIDPARAMETER_DOMAINNOTFOUND = "InvalidParameter.DomainNotFound"

	// 证书信息错误。
	INVALIDPARAMETER_INVALIDCERTINFO = "InvalidParameter.InvalidCertInfo"

	// 参数错误。
	INVALIDPARAMETER_PARAMETERERROR = "InvalidParameter.ParameterError"

	// 域名配置错误。
	INVALIDPARAMETER_SETTINGINVALIDPARAM = "InvalidParameter.SettingInvalidParam"

	// 资源存在错误。
	INVALIDPARAMETER_TARGET = "InvalidParameter.Target"

	// 任务无法生成。
	INVALIDPARAMETER_TASKNOTGENERATED = "InvalidParameter.TaskNotGenerated"

	// 文件上传链接存在问题。
	INVALIDPARAMETER_UPLOADURL = "InvalidParameter.UploadUrl"

	// 与已经添加的记录冲突。
	INVALIDPARAMETERVALUE_CONFLICTRECORD = "InvalidParameterValue.ConflictRecord"

	// DNS 记录与 LB 记录冲突。
	INVALIDPARAMETERVALUE_CONFLICTWITHLBRECORD = "InvalidParameterValue.ConflictWithLBRecord"

	// DNS 记录与 NS 记录冲突。
	INVALIDPARAMETERVALUE_CONFLICTWITHNSRECORD = "InvalidParameterValue.ConflictWithNSRecord"

	// DNS 记录内容错误。
	INVALIDPARAMETERVALUE_INVALIDDNSCONTENT = "InvalidParameterValue.InvalidDNSContent"

	// DNS 记录名称错误。
	INVALIDPARAMETERVALUE_INVALIDDNSNAME = "InvalidParameterValue.InvalidDNSName"

	// DNS 代理域名源站错误。
	INVALIDPARAMETERVALUE_INVALIDPROXYORIGIN = "InvalidParameterValue.InvalidProxyOrigin"

	// 记录已存在。
	INVALIDPARAMETERVALUE_RECORDALREADYEXISTS = "InvalidParameterValue.RecordAlreadyExists"

	// 记录不允许添加。
	INVALIDPARAMETERVALUE_RECORDNOTALLOWED = "InvalidParameterValue.RecordNotAllowed"

	// 本次提交的资源数超过上限。
	LIMITEXCEEDED_BATCHQUOTA = "LimitExceeded.BatchQuota"

	// 当天提交的资源数超过上限。
	LIMITEXCEEDED_DAILYQUOTA = "LimitExceeded.DailyQuota"

	// 操作被拒绝。
	OPERATIONDENIED = "OperationDenied"

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
