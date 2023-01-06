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

package v20180813

const (
	// 此产品的特有错误码

	// 密钥不合法。
	AUTHFAILURE_ACCESSKEYILLEGAL = "AuthFailure.AccessKeyIllegal"

	// DB错误。
	INTERNALERROR_DBERROR = "InternalError.DbError"

	// 加密失败。
	INTERNALERROR_ENCRYPTERROR = "InternalError.EncryptError"

	// 获取appid错误。
	INTERNALERROR_GETAPPIDERROR = "InternalError.GetAppIdError"

	// 获取角色失败。
	INTERNALERROR_GETROLEERROR = "InternalError.GetRoleError"

	// 获取token失败。
	INTERNALERROR_GETSEEDTOKENERROR = "InternalError.GetSeedTokenError"

	// 角色非法。
	INTERNALERROR_ILLEGALROLE = "InternalError.IllegalRole"

	// pb打包失败。
	INTERNALERROR_PBSERIALIZEERROR = "InternalError.PbSerializeError"

	// 系统内部错误，如网络错误。
	INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"

	// 未知错误。
	INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"

	// 不支持该类型密钥。
	INVALIDPARAMETER_ACCESSKEYNOTSUPPORT = "InvalidParameter.AccessKeyNotSupport"

	// 账号不存在或不可用。
	INVALIDPARAMETER_ACCOUNTNOTAVALIABLE = "InvalidParameter.AccountNotAvaliable"

	// 扩展策略过大。
	INVALIDPARAMETER_EXTENDSTRATEGYOVERSIZE = "InvalidParameter.ExtendStrategyOverSize"

	// 越权访问资源。
	INVALIDPARAMETER_GRANTOTHERRESOURCE = "InvalidParameter.GrantOtherResource"

	// 超过频率限制。
	INVALIDPARAMETER_OVERLIMIT = "InvalidParameter.OverLimit"

	// 过期时间超过阈值。
	INVALIDPARAMETER_OVERTIMEERROR = "InvalidParameter.OverTimeError"

	// 参数错误。
	INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"

	// 策略过长。
	INVALIDPARAMETER_POLICYTOOLONG = "InvalidParameter.PolicyTooLong"

	// 策略资源六段式错误。
	INVALIDPARAMETER_RESOUCEERROR = "InvalidParameter.ResouceError"

	// 策略语法错误。
	INVALIDPARAMETER_STRATEGYFORMATERROR = "InvalidParameter.StrategyFormatError"

	// 非法策略。
	INVALIDPARAMETER_STRATEGYINVALID = "InvalidParameter.StrategyInvalid"

	// 临时Code无效。
	INVALIDPARAMETER_TEMPCODENOTAVALIABLE = "InvalidParameter.TempCodeNotAvaliable"

	// WebIdentityToken参数错误。
	INVALIDPARAMETER_WEBIDENTITYTOKENERROR = "InvalidParameter.WebIdentityTokenError"

	// 请求的次数超过了频率限制。
	REQUESTLIMITEXCEEDED = "RequestLimitExceeded"

	// 账号对应的角色不存在。
	RESOURCENOTFOUND_ROLENOTFOUND = "ResourceNotFound.RoleNotFound"

	// 未授权操作。
	UNAUTHORIZEDOPERATION = "UnauthorizedOperation"

	// 操作不支持。
	UNSUPPORTEDOPERATION = "UnsupportedOperation"
)
