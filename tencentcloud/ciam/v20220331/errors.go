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

	// 从账号已经被融合过，不允许再次融合。
	FAILEDOPERATION_ACCOUNTALREADYBELINKED = "FailedOperation.AccountAlreadyBeLinked"

	// 用户属性格式异常。
	FAILEDOPERATION_ATTRIBUTEFORMATERROR = "FailedOperation.AttributeFormatError"

	// 已存在2个有效的任务，请等待其中一部分完成后重试。
	FAILEDOPERATION_DATAFLOWTOOMANYREQUESTS = "FailedOperation.DataFlowTooManyRequests"

	// 该邮箱已被用户绑定。
	FAILEDOPERATION_EMAILALREADYEXISTS = "FailedOperation.EmailAlreadyExists"

	// 邮箱地址不能为空。
	FAILEDOPERATION_EMAILISNULL = "FailedOperation.EmailIsNull"

	// 导入用户时，用户为空。
	FAILEDOPERATION_IMPORTUSERISEMPTY = "FailedOperation.ImportUserIsEmpty"

	// indexedAttribute最长512字符。
	FAILEDOPERATION_INDEXEDATTRIBUTETOOLONG = "FailedOperation.IndexedAttributeTooLong"

	// 无效租户。
	FAILEDOPERATION_INVALIDTENANT = "FailedOperation.InvalidTenant"

	// 状态枚举值错误，请参考接口文档。
	FAILEDOPERATION_INVALIDUSERSTATUSENUM = "FailedOperation.InvalidUserStatusEnum"

	// 无效用户池。
	FAILEDOPERATION_INVALIDUSERSTORE = "FailedOperation.InvalidUserStore"

	// 密码不能为空。
	FAILEDOPERATION_PASSWORDISNULL = "FailedOperation.PasswordIsNull"

	// 该手机号已被用户绑定。
	FAILEDOPERATION_PHONENUMBERALREADYEXISTS = "FailedOperation.PhoneNumberAlreadyExists"

	// 电话号码不能为空。
	FAILEDOPERATION_PHONENUMBERISNULL = "FailedOperation.PhoneNumberIsNull"

	// 主用户不存在。
	FAILEDOPERATION_PRIMARYUSERNOTFOUND = "FailedOperation.PrimaryUserNotFound"

	// 查询条件必须包含id, userName, phoneNumber, email, userGroup, alipayUserId, wechatOpenId,wechatUnionId, qqOpenId, qqUnionId,lastModifiedDate, indexedAttribute1, indexedAttribute2,indexedAttribute3, indexedAttribute4, indexedAttribute5重的一个或者多个。
	FAILEDOPERATION_QUERYUSERSPARAMETERMUSTINWHITELIST = "FailedOperation.QueryUsersParameterMustInWhitelist"

	// 查询用户条件不能有重复的条件。
	FAILEDOPERATION_QUERYUSERSPARAMETERREPEAT = "FailedOperation.QueryUsersParameterRepeat"

	// Sort条件的PropertyKey必须在id, userName, phoneNumber, email, userGroup, alipayUserId, wechatOpenId,wechatUnionId, qqOpenId, qqUnionId,lastModifiedDate, indexedAttribute1, indexedAttribute2,indexedAttribute3, indexedAttribute4, indexedAttribute5中。
	FAILEDOPERATION_QUERYUSERSSORTPARAMETERMUSTINWHITELIST = "FailedOperation.QueryUsersSortParameterMustInWhitelist"

	// 配额超限，如有需求请联系客服人员。
	FAILEDOPERATION_QUOTALIMITEXCEEDED = "FailedOperation.QuotaLimitExceeded"

	// 从用户不存在。
	FAILEDOPERATION_SECONDARYUSERNOTFOUND = "FailedOperation.SecondaryUserNotFound"

	// 用户组不存在。
	FAILEDOPERATION_USERGROUPNOTFOUND = "FailedOperation.UserGroupNotFound"

	// 用户已被冻结。
	FAILEDOPERATION_USERISFREEZE = "FailedOperation.UserIsFreeze"

	// 用户名称已存在。
	FAILEDOPERATION_USERNAMEALREADYEXISTS = "FailedOperation.UserNameAlreadyExists"

	// 用户名不能为空。
	FAILEDOPERATION_USERNAMEISNULL = "FailedOperation.UserNameIsNull"

	// 用户不存在。
	FAILEDOPERATION_USERNOTFOUND = "FailedOperation.UserNotFound"

	// 更新状态必传。
	FAILEDOPERATION_USERSTATUSREQUIRED = "FailedOperation.UserStatusRequired"

	// 内部错误。
	INTERNALERROR = "InternalError"

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
)
