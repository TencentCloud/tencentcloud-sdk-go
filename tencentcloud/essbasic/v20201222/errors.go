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

package v20201222

const (
	// 此产品的特有错误码

	// CAM签名/鉴权错误。
	AUTHFAILURE = "AuthFailure"

	// DryRun 操作，代表请求将会是成功的，只是多传了 DryRun 参数。
	DRYRUNOPERATION = "DryRunOperation"

	// 操作失败。
	FAILEDOPERATION = "FailedOperation"

	// 帐号已存在并实名。
	FAILEDOPERATION_ACCOUNTALREADYEXISTS = "FailedOperation.AccountAlreadyExists"

	// 实名认证错误。
	FAILEDOPERATION_ACCOUNTVERIFYFAIL = "FailedOperation.AccountVerifyFail"

	// 鉴权失败。
	FAILEDOPERATION_AUTHFAIL = "FailedOperation.AuthFail"

	// 加锁失败。
	FAILEDOPERATION_DLOCKFAILED = "FailedOperation.DLockFailed"

	// 已绑定其它手机号码或手机号码已被其它终端(微信)绑定。
	FAILEDOPERATION_ERRBINDINGREPEATED = "FailedOperation.ErrBindingRepeated"

	// 生成企业印章失败。
	FAILEDOPERATION_GENERATEORGSEAL = "FailedOperation.GenerateOrgSeal"

	// 生成个人印章失败。
	FAILEDOPERATION_GENERATEUSERSEAL = "FailedOperation.GenerateUserSeal"

	// 无角色。
	FAILEDOPERATION_NOROLE = "FailedOperation.NoRole"

	// 注册的OpenId已存在。
	FAILEDOPERATION_OPENIDALREADYEXISTS = "FailedOperation.OpenIdAlreadyExists"

	// 注册的企业证件号码已存在。
	FAILEDOPERATION_ORGIDCARDNUMBERALREADYEXISTS = "FailedOperation.OrgIdCardNumberAlreadyExists"

	// 请求的次数超过了频率限制。
	FAILEDOPERATION_REQUESTLIMITEXCEEDED = "FailedOperation.RequestLimitExceeded"

	// 今日验证码发送量已超出限制，请联系工作人员处理。
	FAILEDOPERATION_REQUESTLIMITEXCEEDED1D = "FailedOperation.RequestLimitExceeded1D"

	// 本小时验证码发送数量超出限制，请稍后重试。
	FAILEDOPERATION_REQUESTLIMITEXCEEDED1H = "FailedOperation.RequestLimitExceeded1H"

	// 短信发送频率超出限制，请等待一分钟后重试。
	FAILEDOPERATION_REQUESTLIMITEXCEEDED30S = "FailedOperation.RequestLimitExceeded30S"

	// 内部错误。
	INTERNALERROR = "InternalError"

	// 其他API错误。
	INTERNALERROR_API = "InternalError.Api"

	// 缓存错误。
	INTERNALERROR_CACHE = "InternalError.Cache"

	// 回调错误。
	INTERNALERROR_CALLBACK = "InternalError.Callback"

	// 数据库错误。
	INTERNALERROR_DB = "InternalError.Db"

	// 解密错误。
	INTERNALERROR_DECRYPTION = "InternalError.Decryption"

	// 加密错误。
	INTERNALERROR_ENCRYPTION = "InternalError.Encryption"

	// 生成唯一ID错误。
	INTERNALERROR_GENERATEID = "InternalError.GenerateId"

	// MQ错误。
	INTERNALERROR_MQ = "InternalError.Mq"

	// Pdf错误。
	INTERNALERROR_PDF = "InternalError.Pdf"

	// 序列化错误。
	INTERNALERROR_SERIALIZE = "InternalError.Serialize"

	// 存储错误。
	INTERNALERROR_STORAGE = "InternalError.Storage"

	// 第三方错误。
	INTERNALERROR_THIRDPARTY = "InternalError.ThirdParty"

	// 反序列化错误。
	INTERNALERROR_UNSERIALIZE = "InternalError.UnSerialize"

	// 参数错误。
	INVALIDPARAMETER = "InvalidParameter"

	// 重复添加签署人。
	INVALIDPARAMETER_BIZAPPROVERALREADYEXISTS = "InvalidParameter.BizApproverAlreadyExists"

	// 无效的意愿确认票据。
	INVALIDPARAMETER_INVALIDVERIFYRESULT = "InvalidParameter.InvalidVerifyResult"

	// 参数取值错误。
	INVALIDPARAMETERVALUE = "InvalidParameterValue"

	// 超过配额限制。
	LIMITEXCEEDED = "LimitExceeded"

	// 缺少参数错误。
	MISSINGPARAMETER = "MissingParameter"

	// 操作被拒绝。
	OPERATIONDENIED = "OperationDenied"

	// 应用号已被禁止。
	OPERATIONDENIED_BANNEDAPPLICATION = "OperationDenied.BannedApplication"

	// 没有API权限。
	OPERATIONDENIED_NOAPIAUTH = "OperationDenied.NoApiAuth"

	// 未通过补充实名。
	OPERATIONDENIED_NOEXTRAVERIFY = "OperationDenied.NoExtraVerify"

	// 未通过个人实名。
	OPERATIONDENIED_NOIDENTITYVERIFY = "OperationDenied.NoIdentityVerify"

	// 未在腾讯云实名打款。
	OPERATIONDENIED_NOPAYMENTVERIFY = "OperationDenied.NoPaymentVerify"

	// 没有登录态需要登录。
	OPERATIONDENIED_NOSESSION = "OperationDenied.NoSession"

	// 未在腾讯云实名。
	OPERATIONDENIED_NOVERIFY = "OperationDenied.NoVerify"

	// 非企业主账号。
	OPERATIONDENIED_NOTOWNERUIN = "OperationDenied.NotOwnerUin"

	// 用户与企业不对应。
	OPERATIONDENIED_USERNOTINORGANIZATION = "OperationDenied.UserNotInOrganization"

	// 请求的次数超过了频率限制。
	REQUESTLIMITEXCEEDED = "RequestLimitExceeded"

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

	// 未知参数错误。
	UNKNOWNPARAMETER = "UnknownParameter"

	// 操作不支持。
	UNSUPPORTEDOPERATION = "UnsupportedOperation"
)
