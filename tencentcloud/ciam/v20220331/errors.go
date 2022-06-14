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

package v20220331

const (
	// 此产品的特有错误码

	// CAM签名/鉴权错误。
	AUTHFAILURE = "AuthFailure"

	// DryRun 操作，代表请求将会是成功的，只是多传了 DryRun 参数。
	DRYRUNOPERATION = "DryRunOperation"

	// 操作失败。
	FAILEDOPERATION = "FailedOperation"

	// 认证源名称已存在。
	FAILEDOPERATION_AUTHSOURCENAMEALREADYEXISTS = "FailedOperation.AuthSourceNameAlreadyExists"

	// 认证源不存在。
	FAILEDOPERATION_AUTHSOURCENOTFOUND = "FailedOperation.AuthSourceNotFound"

	// 无效租户。
	FAILEDOPERATION_INVALIDTENANT = "FailedOperation.InvalidTenant"

	// 无效用户池。
	FAILEDOPERATION_INVALIDUSERSTORE = "FailedOperation.InvalidUserStore"

	// 配额超限，如有需求请联系客服人员。
	FAILEDOPERATION_QUOTALIMITEXCEEDED = "FailedOperation.QuotaLimitExceeded"

	// 用户已被冻结。
	FAILEDOPERATION_USERISFREEZE = "FailedOperation.UserIsFreeze"

	// 用户不存在。
	FAILEDOPERATION_USERNOTFOUND = "FailedOperation.UserNotFound"

	// 属性标识已存在。
	FAILEDOPERATION_USERPROPERTYCODEALREADYEXISTS = "FailedOperation.UserPropertyCodeAlreadyExists"

	// 用户属性不存在。
	FAILEDOPERATION_USERPROPERTYNOTFOUND = "FailedOperation.UserPropertyNotFound"

	// 已存在相同域名。
	FAILEDOPERATION_USERSTOREDOMAINALREADYEXISTS = "FailedOperation.UserStoreDomainAlreadyExists"

	// 内部错误。
	INTERNALERROR = "InternalError"

	// 判断存在用户内部请求错误，请稍后重试。
	INTERNALERROR_JUDGEUSEREXISTEXCEPTION = "InternalError.JudgeUserExistException"

	// 未知错误。
	INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"

	// 参数错误。
	INVALIDPARAMETER = "InvalidParameter"

	// 请求参数不合法。
	INVALIDPARAMETER_PARAMETERILLEGAL = "InvalidParameter.ParameterIllegal"

	// 参数取值错误。
	INVALIDPARAMETERVALUE = "InvalidParameterValue"

	// 超过配额限制。
	LIMITEXCEEDED = "LimitExceeded"

	// 缺少参数错误。
	MISSINGPARAMETER = "MissingParameter"

	// 操作被拒绝。
	OPERATIONDENIED = "OperationDenied"

	// 请求的次数超过了频率限制。
	REQUESTLIMITEXCEEDED = "RequestLimitExceeded"

	// 请求过于频繁，请稍后再试。
	REQUESTLIMITEXCEEDED_FREQUENTREQUEST = "RequestLimitExceeded.FrequentRequest"

	// 请求重复，请稍后再试。
	REQUESTLIMITEXCEEDED_REPEATREQUEST = "RequestLimitExceeded.RepeatRequest"

	// 资源被占用。
	RESOURCEINUSE = "ResourceInUse"

	// 资源不足。
	RESOURCEINSUFFICIENT = "ResourceInsufficient"

	// 资源不存在。
	RESOURCENOTFOUND = "ResourceNotFound"

	// 资源不可用。
	RESOURCEUNAVAILABLE = "ResourceUnavailable"

	// 资源售罄。
	RESOURCESSOLDOUT = "ResourcesSoldOut"

	// 未授权操作。
	UNAUTHORIZEDOPERATION = "UnauthorizedOperation"

	// 服务未开通。
	UNAUTHORIZEDOPERATION_TENANTNOTACTIVATED = "UnauthorizedOperation.TenantNotActivated"

	// 未知参数错误。
	UNKNOWNPARAMETER = "UnknownParameter"

	// 操作不支持。
	UNSUPPORTEDOPERATION = "UnsupportedOperation"

	// 认证源已启用，不允许删除。
	UNSUPPORTEDOPERATION_ENABLEDAUTHSOURCECANNOTBEDELETED = "UnsupportedOperation.EnabledAuthSourceCanNotBeDeleted"

	// 认证源已启用，不允许更新。
	UNSUPPORTEDOPERATION_ENABLEDAUTHSOURCECANNOTBEUPDATED = "UnsupportedOperation.EnabledAuthSourceCanNotBeUpdated"

	// 内置属性不支持该操作。
	UNSUPPORTEDOPERATION_INTERNALUSERPROPERTY = "UnsupportedOperation.InternalUserProperty"

	// 属性被应用关联，不允许删除。
	UNSUPPORTEDOPERATION_LINKEDAPPUSERPROPERTYCANNOTBEDELETED = "UnsupportedOperation.LinkedAppUserPropertyCanNotBeDeleted"

	// 认证源被应用关联，不允许关闭。
	UNSUPPORTEDOPERATION_LINKEDAUTHSOURCECANNOTBECLOSED = "UnsupportedOperation.LinkedAuthSourceCanNotBeClosed"

	// 认证源被应用关联，不允许删除。
	UNSUPPORTEDOPERATION_LINKEDAUTHSOURCECANNOTBEDELETED = "UnsupportedOperation.LinkedAuthSourceCanNotBeDeleted"

	// 属性被认证源关联，不允许删除。
	UNSUPPORTEDOPERATION_LINKEDAUTHSOURCEUSERPROPERTYCANNOTBEDELETED = "UnsupportedOperation.LinkedAuthSourceUserPropertyCanNotBeDeleted"

	// 属性被其它模块关联，不允许删除。
	UNSUPPORTEDOPERATION_LINKEDUSERPROPERTYCANNOTBEDELETED = "UnsupportedOperation.LinkedUserPropertyCanNotBeDeleted"

	// 不允许使用平台域名后缀作为自有域名后缀。
	UNSUPPORTEDOPERATION_PLATFORMDOMAINSUFFIXCANNOTBEUSED = "UnsupportedOperation.PlatformDomainSuffixCanNotBeUsed"

	// 用户池下已存在用户，用户属性不允许被删除。
	UNSUPPORTEDOPERATION_WHENUSEREXISTUSERPROPERTYCANNOTBEDELETED = "UnsupportedOperation.WhenUserExistUserPropertyCanNotBeDeleted"
)
