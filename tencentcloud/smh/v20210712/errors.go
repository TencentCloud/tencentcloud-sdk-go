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

package v20210712

const (
	// 此产品的特有错误码

	// 内部错误。
	INTERNALERROR = "InternalError"

	// 媒体库资源删除失败
	INTERNALERROR_MODIFYRESOURCESTATUSFAIL = "InternalError.ModifyResourceStatusFail"

	// 后付费开通失败
	INTERNALERROR_POSTPAIDFAIL = "InternalError.PostpaidFail"

	// 发送短信验证码时发生错误。
	INTERNALERROR_SENDSMS = "InternalError.SendSms"

	// AccountName 和 AccountPassword 必须同时指定或同时不指定。
	INVALIDPARAMETER_ONEOFACCOUNTNAMEANDACCOUNTPASSWORDISEMPTY = "InvalidParameter.OneOfAccountNameAndAccountPasswordIsEmpty"

	// CountryCode 和 PhoneNumber 必须同时指定或同时不指定。
	INVALIDPARAMETER_ONEOFCOUNTRYCODEANDPHONENUMBERISEMPTY = "InvalidParameter.OneOfCountryCodeAndPhoneNumberIsEmpty"

	// 参数取值错误。
	INVALIDPARAMETERVALUE = "InvalidParameterValue"

	// AccountName 过长。
	INVALIDPARAMETERVALUE_ACCOUNTNAMETOOLONG = "InvalidParameterValue.AccountNameTooLong"

	// AccountPassword 无效。
	INVALIDPARAMETERVALUE_ACCOUNTPASSWORD = "InvalidParameterValue.AccountPassword"

	// AccountUserId 过长。
	INVALIDPARAMETERVALUE_ACCOUNTUSERIDTOOLONG = "InvalidParameterValue.AccountUserIdTooLong"

	// 存储桶名称无效。
	INVALIDPARAMETERVALUE_BUCKETNAMEINVALID = "InvalidParameterValue.BucketNameInvalid"

	// 存储桶名称不属于当前主账号。
	INVALIDPARAMETERVALUE_BUCKETNAMENOTBELONGYOU = "InvalidParameterValue.BucketNameNotBelongYou"

	// 存储桶不存在或不在指定地域。
	INVALIDPARAMETERVALUE_BUCKETNOTFOUND = "InvalidParameterValue.BucketNotFound"

	// 服务已不支持自选桶
	INVALIDPARAMETERVALUE_BUCKETNOTSUPPORT = "InvalidParameterValue.BucketNotSupport"

	// 存储桶所在地域无效。
	INVALIDPARAMETERVALUE_BUCKETREGIONINVALID = "InvalidParameterValue.BucketRegionInvalid"

	// 指定的存储类型无效。
	INVALIDPARAMETERVALUE_COSSTORAGECLASS = "InvalidParameterValue.CosStorageClass"

	// 指定的存储桶未开启智能分层存储，请先开通存储桶的智能分层存储功能。
	INVALIDPARAMETERVALUE_COSSTORAGECLASSINTELLIGENTTIERING = "InvalidParameterValue.CosStorageClassIntelligentTiering"

	// 指定的国家代码无效。
	INVALIDPARAMETERVALUE_COUNTRYCODE = "InvalidParameterValue.CountryCode"

	// CountryCode 过长。
	INVALIDPARAMETERVALUE_COUNTRYCODETOOLONG = "InvalidParameterValue.CountryCodeTooLong"

	// DestroyTime 无效。
	INVALIDPARAMETERVALUE_DESTROYTIME = "InvalidParameterValue.DestroyTime"

	// AccountName 重复。
	INVALIDPARAMETERVALUE_DUPLICATEACCOUNTNAME = "InvalidParameterValue.DuplicateAccountName"

	// AccountUserId 重复。
	INVALIDPARAMETERVALUE_DUPLICATEACCOUNTUSERID = "InvalidParameterValue.DuplicateAccountUserId"

	// Email 重复。
	INVALIDPARAMETERVALUE_DUPLICATEEMAIL = "InvalidParameterValue.DuplicateEmail"

	// CountryCode 和 PhoneNumber 重复。
	INVALIDPARAMETERVALUE_DUPLICATEUSERPHONENUMBER = "InvalidParameterValue.DuplicateUserPhoneNumber"

	// Email 过长。
	INVALIDPARAMETERVALUE_EMAILTOOLONG = "InvalidParameterValue.EmailTooLong"

	// 指定的 Filter 无效。
	INVALIDPARAMETERVALUE_FILTER = "InvalidParameterValue.Filter"

	// AccountName 无效。
	INVALIDPARAMETERVALUE_INVALIDACCOUNTNAME = "InvalidParameterValue.InvalidAccountName"

	// AccountUserId 无效。
	INVALIDPARAMETERVALUE_INVALIDACCOUNTUSERID = "InvalidParameterValue.InvalidAccountUserId"

	// CountryCode 无效。
	INVALIDPARAMETERVALUE_INVALIDCOUNTRYCODE = "InvalidParameterValue.InvalidCountryCode"

	// Email 无效。
	INVALIDPARAMETERVALUE_INVALIDEMAIL = "InvalidParameterValue.InvalidEmail"

	// PhoneNumber 无效。
	INVALIDPARAMETERVALUE_INVALIDPHONENUMBER = "InvalidParameterValue.InvalidPhoneNumber"

	// IsolateTime 无效。
	INVALIDPARAMETERVALUE_ISOLATETIME = "InvalidParameterValue.IsolateTime"

	// 参数值数量或字符数超过限制。
	INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"

	// 指定的手机号码与当前的相同。
	INVALIDPARAMETERVALUE_NOTMODIFIED = "InvalidParameterValue.NotModified"

	// 指定的手机号码非该企业的超级管理员。
	INVALIDPARAMETERVALUE_NOTSUPERADMIN = "InvalidParameterValue.NotSuperAdmin"

	// 指定的手机号码无效。
	INVALIDPARAMETERVALUE_PHONENUMBER = "InvalidParameterValue.PhoneNumber"

	// PhoneNumber 过长。
	INVALIDPARAMETERVALUE_PHONENUMBERTOOLONG = "InvalidParameterValue.PhoneNumberTooLong"

	// 指定的 Role 不支持。
	INVALIDPARAMETERVALUE_ROLENOTSUPPORT = "InvalidParameterValue.RoleNotSupport"

	// 指定了太多的过滤器，过滤器最多为 100 个。
	INVALIDPARAMETERVALUE_TOOMANYFILTERS = "InvalidParameterValue.TooManyFilters"

	// 已达到用户数量限制，请先升级产品规格。
	LIMITEXCEEDED_USERLIMIT = "LimitExceeded.UserLimit"

	// 发送至指定手机号码的短信验证码已达到频率限制。
	REQUESTLIMITEXCEEDED_SENDSMS = "RequestLimitExceeded.SendSms"

	// 多租户空间媒体库中存在正在使用的租户空间。
	RESOURCEINUSE_MULTISPACE = "ResourceInUse.MultiSpace"

	// 媒体库不存在或不属于当前账号。
	RESOURCENOTFOUND_LIBRARY = "ResourceNotFound.Library"

	// 官方云盘实例不存在或不属于当前账号。
	RESOURCENOTFOUND_OFFICIALINSTANCE = "ResourceNotFound.OfficialInstance"

	// 指定的用户不存在。
	RESOURCENOTFOUND_USER = "ResourceNotFound.User"

	// 未授予 cam:PassRole 权限。
	UNAUTHORIZEDOPERATION_PASSROLE = "UnauthorizedOperation.PassRole"

	// 未授予 SMH 服务相关角色。
	UNAUTHORIZEDOPERATION_SERVICELINKEDROLE = "UnauthorizedOperation.ServiceLinkedRole"

	// 短信验证码错误或已过期。
	UNAUTHORIZEDOPERATION_SMSCODE = "UnauthorizedOperation.SmsCode"

	// 短信验证码验证次数超限，请重新发送短信验证码。
	UNAUTHORIZEDOPERATION_SMSCODEEXCEEDED = "UnauthorizedOperation.SmsCodeExceeded"

	// 账户余额不足
	UNSUPPORTEDOPERATION_BALANCELESS = "UnsupportedOperation.BalanceLess"

	// 指定的目的不受支持。
	UNSUPPORTEDOPERATION_PURPOSE = "UnsupportedOperation.Purpose"
)
