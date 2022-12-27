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

package v20210111

const (
	// 此产品的特有错误码

	// 短信内容中含有敏感词，请联系 [腾讯云短信小助手](https://cloud.tencent.com/document/product/382/3773#.E6.8A.80.E6.9C.AF.E4.BA.A4.E6.B5.81)。
	FAILEDOPERATION_CONTAINSENSITIVEWORD = "FailedOperation.ContainSensitiveWord"

	// 请求包解析失败，通常情况下是由于没有遵守 API 接口说明规范导致的，请参考 [请求包体解析1004错误详解](https://cloud.tencent.com/document/product/382/9558#.E8.BF.94.E5.9B.9E1004.E9.94.99.E8.AF.AF.E5.A6.82.E4.BD.95.E5.A4.84.E7.90.86.EF.BC.9F)。
	FAILEDOPERATION_FAILRESOLVEPACKET = "FailedOperation.FailResolvePacket"

	// 个人用户不能申请营销短信。
	FAILEDOPERATION_FORBIDADDMARKETINGTEMPLATES = "FailedOperation.ForbidAddMarketingTemplates"

	// 套餐包余量不足，请 [购买套餐包](https://buy.cloud.tencent.com/sms)。
	FAILEDOPERATION_INSUFFICIENTBALANCEINSMSPACKAGE = "FailedOperation.InsufficientBalanceInSmsPackage"

	// 解析请求包体时候失败。
	FAILEDOPERATION_JSONPARSEFAIL = "FailedOperation.JsonParseFail"

	// 营销短信发送时间限制，为避免骚扰用户，营销短信只允许在8点到22点发送。
	FAILEDOPERATION_MARKETINGSENDTIMECONSTRAINT = "FailedOperation.MarketingSendTimeConstraint"

	// 没有申请签名之前，无法申请模板，请根据 [创建签名](https://cloud.tencent.com/document/product/382/37794#.E5.88.9B.E5.BB.BA.E7.AD.BE.E5.90.8D) 申请完成之后再次申请。
	FAILEDOPERATION_MISSINGSIGNATURE = "FailedOperation.MissingSignature"

	// 无法识别签名，请确认是否已有签名通过申请，一般是签名未通过申请，可以查看 [签名审核](https://cloud.tencent.com/document/product/382/37745)。
	FAILEDOPERATION_MISSINGSIGNATURELIST = "FailedOperation.MissingSignatureList"

	// 此签名 ID 未提交申请或不存在，不能进行修改操作，请检查您的 SignId 是否填写正确。
	FAILEDOPERATION_MISSINGSIGNATURETOMODIFY = "FailedOperation.MissingSignatureToModify"

	// 无法识别模板，请确认是否已有模板通过申请，一般是模板未通过申请，可以查看 [模板审核](https://cloud.tencent.com/document/product/382/37745)。
	FAILEDOPERATION_MISSINGTEMPLATELIST = "FailedOperation.MissingTemplateList"

	// 此模板 ID 未提交申请或不存在，不能进行修改操作，请检查您的 TemplateId是否填写正确。
	FAILEDOPERATION_MISSINGTEMPLATETOMODIFY = "FailedOperation.MissingTemplateToModify"

	// 非企业认证无法使用签名及模板相关接口，您可以[ 变更实名认证模式](https://cloud.tencent.com/document/product/378/34075)，变更为企业认证用户后，约1小时左右生效。
	FAILEDOPERATION_NOTENTERPRISECERTIFICATION = "FailedOperation.NotEnterpriseCertification"

	// 其他错误，一般是由于参数携带不符合要求导致，请参考API接口说明，如有需要请联系 [腾讯云短信小助手](https://cloud.tencent.com/document/product/382/3773#.E6.8A.80.E6.9C.AF.E4.BA.A4.E6.B5.81)。
	FAILEDOPERATION_OTHERERROR = "FailedOperation.OtherError"

	// 未知错误，如有需要请联系 [腾讯云短信小助手](https://cloud.tencent.com/document/product/382/3773#.E6.8A.80.E6.9C.AF.E4.BA.A4.E6.B5.81)。
	FAILEDOPERATION_PARAMETERSOTHERERROR = "FailedOperation.ParametersOtherError"

	// 手机号在免打扰名单库中，通常是用户退订或者命中运营商免打扰名单导致的，可联系 [腾讯云短信小助手](https://cloud.tencent.com/document/product/382/3773#.E6.8A.80.E6.9C.AF.E4.BA.A4.E6.B5.81) 解决。
	FAILEDOPERATION_PHONENUMBERINBLACKLIST = "FailedOperation.PhoneNumberInBlacklist"

	// 号码解析失败，请检查号码是否符合 E.164 标准。
	FAILEDOPERATION_PHONENUMBERPARSEFAIL = "FailedOperation.PhoneNumberParseFail"

	// 非主账号无法使用拉取模板列表功能。您可以使用主账号下云 API 密钥来调用接口。
	FAILEDOPERATION_PROHIBITSUBACCOUNTUSE = "FailedOperation.ProhibitSubAccountUse"

	// 签名 ID 不存在。
	FAILEDOPERATION_SIGNIDNOTEXIST = "FailedOperation.SignIdNotExist"

	// 签名个数达到最大值。
	FAILEDOPERATION_SIGNNUMBERLIMIT = "FailedOperation.SignNumberLimit"

	// 签名未审批或格式错误。（1）可登录 [短信控制台](https://console.cloud.tencent.com/smsv2)，核查签名是否已审批并且审批通过；（2）核查是否符合格式规范，签名只能由中英文、数字组成，要求2 - 12个字，若存在疑问可联系 [腾讯云短信小助手](https://cloud.tencent.com/document/product/382/3773#.E6.8A.80.E6.9C.AF.E4.BA.A4.E6.B5.81)。
	FAILEDOPERATION_SIGNATUREINCORRECTORUNAPPROVED = "FailedOperation.SignatureIncorrectOrUnapproved"

	// 此模板已经通过审核，无法再次进行修改。
	FAILEDOPERATION_TEMPLATEALREADYPASSEDCHECK = "FailedOperation.TemplateAlreadyPassedCheck"

	// 模板 ID 不存在。
	FAILEDOPERATION_TEMPLATEIDNOTEXIST = "FailedOperation.TemplateIdNotExist"

	// 模板未审批或内容不匹配。（1）可登录 [短信控制台](https://console.cloud.tencent.com/smsv2)，核查模板是否已审批并审批通过；（2）核查是否符合 [格式规范](https://cloud.tencent.com/document/product/382/9558#.E8.BF.94.E5.9B.9E1014.E9.94.99.E8.AF.AF.E5.A6.82.E4.BD.95.E5.A4.84.E7.90.86.EF.BC.9F)，若存在疑问可联系 [腾讯云短信小助手](https://cloud.tencent.com/document/product/382/3773#.E6.8A.80.E6.9C.AF.E4.BA.A4.E6.B5.81)。
	FAILEDOPERATION_TEMPLATEINCORRECTORUNAPPROVED = "FailedOperation.TemplateIncorrectOrUnapproved"

	// 模板个数达到最大值。
	FAILEDOPERATION_TEMPLATENUMBERLIMIT = "FailedOperation.TemplateNumberLimit"

	// 请求内容与审核通过的模板内容不匹配。请检查请求中模板参数的个数是否与申请的模板一致。若存在疑问可联系 [腾讯云短信小助手](https://cloud.tencent.com/document/product/382/3773#.E6.8A.80.E6.9C.AF.E4.BA.A4.E6.B5.81)。
	FAILEDOPERATION_TEMPLATEPARAMSETNOTMATCHAPPROVEDTEMPLATE = "FailedOperation.TemplateParamSetNotMatchApprovedTemplate"

	// 模板未审批或不存在。可登录 [短信控制台](https://console.cloud.tencent.com/smsv2)，核查模板是否已审批并审批通过。若存在疑问可联系 [腾讯云短信小助手](https://cloud.tencent.com/document/product/382/3773#.E6.8A.80.E6.9C.AF.E4.BA.A4.E6.B5.81)。
	FAILEDOPERATION_TEMPLATEUNAPPROVEDORNOTEXIST = "FailedOperation.TemplateUnapprovedOrNotExist"

	// 解析用户参数失败，可联系 [腾讯云短信小助手](https://cloud.tencent.com/document/product/382/3773#.E6.8A.80.E6.9C.AF.E4.BA.A4.E6.B5.81)。
	INTERNALERROR_JSONPARSEFAIL = "InternalError.JsonParseFail"

	// 其他错误，请联系 [腾讯云短信小助手](https://cloud.tencent.com/document/product/382/3773#.E6.8A.80.E6.9C.AF.E4.BA.A4.E6.B5.81) 并提供失败手机号。
	INTERNALERROR_OTHERERROR = "InternalError.OtherError"

	// 解析运营商包体失败，可联系 [sms helper](https://cloud.tencent.com/document/product/382/3773#.E6.8A.80.E6.9C.AF.E4.BA.A4.E6.B5.81) 。
	INTERNALERROR_PARSEBACKENDRESPONSEFAIL = "InternalError.ParseBackendResponseFail"

	// 请求发起时间不正常，通常是由于您的服务器时间与腾讯云服务器时间差异超过10分钟导致的，请核对服务器时间及 API 接口中的时间字段是否正常。
	INTERNALERROR_REQUESTTIMEEXCEPTION = "InternalError.RequestTimeException"

	// 不存在该 RESTAPI 接口，请核查 REST API 接口说明。
	INTERNALERROR_RESTAPIINTERFACENOTEXIST = "InternalError.RestApiInterfaceNotExist"

	// 接口超时或短信收发包超时，请检查您的网络是否有波动，或联系 [腾讯云短信小助手](https://cloud.tencent.com/document/product/382/3773#.E6.8A.80.E6.9C.AF.E4.BA.A4.E6.B5.81) 解决。
	INTERNALERROR_SENDANDRECVFAIL = "InternalError.SendAndRecvFail"

	// 后端包体中请求包体没有 Sig 字段或 Sig 为空。
	INTERNALERROR_SIGFIELDMISSING = "InternalError.SigFieldMissing"

	// 后端校验 Sig 失败。
	INTERNALERROR_SIGVERIFICATIONFAIL = "InternalError.SigVerificationFail"

	// 请求下发短信超时，请参考 [60008错误详解](https://cloud.tencent.com/document/product/382/9558#.E8.BF.94.E5.9B.9E60008.E9.94.99.E8.AF.AF.E5.A6.82.E4.BD.95.E5.A4.84.E7.90.86.EF.BC.9F)。
	INTERNALERROR_TIMEOUT = "InternalError.Timeout"

	// 未知错误类型。
	INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"

	// 参数错误。
	INVALIDPARAMETER = "InvalidParameter"

	// 账号与应用id不匹配。
	INVALIDPARAMETER_APPIDANDBIZID = "InvalidParameter.AppidAndBizId"

	// 存在敏感词。
	INVALIDPARAMETER_DIRTYWORDFOUND = "InvalidParameter.DirtyWordFound"

	// 参数有误，如有需要请联系 [腾讯云短信小助手](https://cloud.tencent.com/document/product/382/3773#.E6.8A.80.E6.9C.AF.E4.BA.A4.E6.B5.81)。
	INVALIDPARAMETER_INVALIDPARAMETERS = "InvalidParameter.InvalidParameters"

	// 参数 BeginTime 校验失败。
	INVALIDPARAMETERVALUE_BEGINTIMEVERIFYFAIL = "InvalidParameterValue.BeginTimeVerifyFail"

	// 请求的短信内容太长，短信长度规则请参考 [国内短信内容长度计算规则](https://cloud.tencent.com/document/product/382/18058)。
	INVALIDPARAMETERVALUE_CONTENTLENGTHLIMIT = "InvalidParameterValue.ContentLengthLimit"

	// 参数 EndTime 校验失败。
	INVALIDPARAMETERVALUE_ENDTIMEVERIFYFAIL = "InvalidParameterValue.EndTimeVerifyFail"

	// 上传的转码图片格式错误，请参照 API 接口说明中对改字段的说明，如有需要请联系 [腾讯云短信小助手](https://cloud.tencent.com/document/product/382/3773#.E6.8A.80.E6.9C.AF.E4.BA.A4.E6.B5.81)。
	INVALIDPARAMETERVALUE_IMAGEINVALID = "InvalidParameterValue.ImageInvalid"

	// 手机号格式错误。
	INVALIDPARAMETERVALUE_INCORRECTPHONENUMBER = "InvalidParameterValue.IncorrectPhoneNumber"

	// DocumentType 字段校验错误，请参照 API 接口说明中对改字段的说明，如有需要请联系 [腾讯云短信小助手](https://cloud.tencent.com/document/product/382/3773#.E6.8A.80.E6.9C.AF.E4.BA.A4.E6.B5.81)。
	INVALIDPARAMETERVALUE_INVALIDDOCUMENTTYPE = "InvalidParameterValue.InvalidDocumentType"

	// International 字段校验错误，请参照 API 接口说明中对改字段的说明，如有需要请联系 [腾讯云短信小助手](https://cloud.tencent.com/document/product/382/3773#.E6.8A.80.E6.9C.AF.E4.BA.A4.E6.B5.81)。
	INVALIDPARAMETERVALUE_INVALIDINTERNATIONAL = "InvalidParameterValue.InvalidInternational"

	// SignPurpose 字段校验错误，请参照 API 接口说明中对改字段的说明，如有需要请联系 [腾讯云短信小助手](https://cloud.tencent.com/document/product/382/3773#.E6.8A.80.E6.9C.AF.E4.BA.A4.E6.B5.81)。
	INVALIDPARAMETERVALUE_INVALIDSIGNPURPOSE = "InvalidParameterValue.InvalidSignPurpose"

	// 无效的拉取起始/截止时间，具体原因可能是请求的 SendDateTime 大于 EndDateTime。
	INVALIDPARAMETERVALUE_INVALIDSTARTTIME = "InvalidParameterValue.InvalidStartTime"

	// 模板格式错误，请参考[正文模板审核标准](https://cloud.tencent.com/document/product/382/39023)。
	INVALIDPARAMETERVALUE_INVALIDTEMPLATEFORMAT = "InvalidParameterValue.InvalidTemplateFormat"

	// UsedMethod 字段校验错误，请参照 API 接口说明中对改字段的说明，如有需要请联系 [腾讯云短信小助手](https://cloud.tencent.com/document/product/382/3773#.E6.8A.80.E6.9C.AF.E4.BA.A4.E6.B5.81)。
	INVALIDPARAMETERVALUE_INVALIDUSEDMETHOD = "InvalidParameterValue.InvalidUsedMethod"

	// 参数 Limit 校验失败。
	INVALIDPARAMETERVALUE_LIMITVERIFYFAIL = "InvalidParameterValue.LimitVerifyFail"

	// 参数 Offset 校验失败。
	INVALIDPARAMETERVALUE_OFFSETVERIFYFAIL = "InvalidParameterValue.OffsetVerifyFail"

	// 禁止在模板变量中使用 URL。
	INVALIDPARAMETERVALUE_PROHIBITEDUSEURLINTEMPLATEPARAMETER = "InvalidParameterValue.ProhibitedUseUrlInTemplateParameter"

	// SdkAppId 不存在。
	INVALIDPARAMETERVALUE_SDKAPPIDNOTEXIST = "InvalidParameterValue.SdkAppIdNotExist"

	// 此签名已经通过审核，无法再次进行修改。
	INVALIDPARAMETERVALUE_SIGNALREADYPASSEDCHECK = "InvalidParameterValue.SignAlreadyPassedCheck"

	// 已存在相同的待审核签名。
	INVALIDPARAMETERVALUE_SIGNEXISTANDUNAPPROVED = "InvalidParameterValue.SignExistAndUnapproved"

	// 签名内容长度过长。
	INVALIDPARAMETERVALUE_SIGNNAMELENGTHTOOLONG = "InvalidParameterValue.SignNameLengthTooLong"

	// 验证码模板参数格式错误，验证码类模板，模板变量只能传入0 - 6位（包括6位）纯数字。
	INVALIDPARAMETERVALUE_TEMPLATEPARAMETERFORMATERROR = "InvalidParameterValue.TemplateParameterFormatError"

	// 单个模板变量字符数超过12个，企业认证用户不限制单个变量值字数，您可以 [变更实名认证模式](https://cloud.tencent.com/document/product/378/34075)，变更为企业认证用户后，该限制变更约1小时左右生效。
	INVALIDPARAMETERVALUE_TEMPLATEPARAMETERLENGTHLIMIT = "InvalidParameterValue.TemplateParameterLengthLimit"

	// 模板内容存在敏感词，请参考[正文模板审核标准](https://cloud.tencent.com/document/product/382/39023)。
	INVALIDPARAMETERVALUE_TEMPLATEWITHDIRTYWORDS = "InvalidParameterValue.TemplateWithDirtyWords"

	// 业务短信国家/地区日下发条数超过设定的上限，可自行到控制台应用管理>基础配置下调整国际港澳台短信发送限制。
	LIMITEXCEEDED_APPCOUNTRYORREGIONDAILYLIMIT = "LimitExceeded.AppCountryOrRegionDailyLimit"

	// 业务短信国家/地区不在国际港澳台短信发送限制设置的列表中而禁发，可自行到控制台应用管理>基础配置下调整国际港澳台短信发送限制。
	LIMITEXCEEDED_APPCOUNTRYORREGIONINBLACKLIST = "LimitExceeded.AppCountryOrRegionInBlacklist"

	// 业务短信日下发条数超过设定的上限 ，可自行到控制台调整短信频率限制策略。
	LIMITEXCEEDED_APPDAILYLIMIT = "LimitExceeded.AppDailyLimit"

	// 业务短信国际/港澳台日下发条数超过设定的上限，可自行到控制台应用管理>基础配置下调整发送总量阈值。
	LIMITEXCEEDED_APPGLOBALDAILYLIMIT = "LimitExceeded.AppGlobalDailyLimit"

	// 业务短信中国大陆日下发条数超过设定的上限，可自行到控制台应用管理>基础配置下调整发送总量阈值。
	LIMITEXCEEDED_APPMAINLANDCHINADAILYLIMIT = "LimitExceeded.AppMainlandChinaDailyLimit"

	// 短信日下发条数超过设定的上限 (国际/港澳台)，如需调整限制，可联系 [腾讯云短信小助手](https://cloud.tencent.com/document/product/382/3773)。
	LIMITEXCEEDED_DAILYLIMIT = "LimitExceeded.DailyLimit"

	// 下发短信命中了频率限制策略，可自行到控制台调整短信频率限制策略，如有其他需求请联系 [腾讯云短信小助手](https://cloud.tencent.com/document/product/382/3773#.E6.8A.80.E6.9C.AF.E4.BA.A4.E6.B5.81)。
	LIMITEXCEEDED_DELIVERYFREQUENCYLIMIT = "LimitExceeded.DeliveryFrequencyLimit"

	// 调用接口单次提交的手机号个数超过200个，请遵守 API 接口输入参数 PhoneNumberSet 描述。
	LIMITEXCEEDED_PHONENUMBERCOUNTLIMIT = "LimitExceeded.PhoneNumberCountLimit"

	// 单个手机号日下发短信条数超过设定的上限，可自行到控制台调整短信频率限制策略。
	LIMITEXCEEDED_PHONENUMBERDAILYLIMIT = "LimitExceeded.PhoneNumberDailyLimit"

	// 单个手机号1小时内下发短信条数超过设定的上限，可自行到控制台调整短信频率限制策略。
	LIMITEXCEEDED_PHONENUMBERONEHOURLIMIT = "LimitExceeded.PhoneNumberOneHourLimit"

	// 单个手机号下发相同内容超过设定的上限，可自行到控制台调整短信频率限制策略。
	LIMITEXCEEDED_PHONENUMBERSAMECONTENTDAILYLIMIT = "LimitExceeded.PhoneNumberSameContentDailyLimit"

	// 单个手机号30秒内下发短信条数超过设定的上限，可自行到控制台调整短信频率限制策略。
	LIMITEXCEEDED_PHONENUMBERTHIRTYSECONDLIMIT = "LimitExceeded.PhoneNumberThirtySecondLimit"

	// 缺少参数错误。
	MISSINGPARAMETER = "MissingParameter"

	// 传入的号码列表为空，请确认您的参数中是否传入号码。
	MISSINGPARAMETER_EMPTYPHONENUMBERSET = "MissingParameter.EmptyPhoneNumberSet"

	// 个人用户没有发营销短信的权限，请参考 [权益区别](https://cloud.tencent.com/document/product/382/13444)。
	UNAUTHORIZEDOPERATION_INDIVIDUALUSERMARKETINGSMSPERMISSIONDENY = "UnauthorizedOperation.IndividualUserMarketingSmsPermissionDeny"

	// 请求 IP 不在白名单中，您配置了校验请求来源 IP，但是检测到当前请求 IP 不在配置列表中，如有需要请联系 [腾讯云短信小助手](https://cloud.tencent.com/document/product/382/3773#.E6.8A.80.E6.9C.AF.E4.BA.A4.E6.B5.81)。
	UNAUTHORIZEDOPERATION_REQUESTIPNOTINWHITELIST = "UnauthorizedOperation.RequestIpNotInWhitelist"

	// 请求没有权限，请联系 [腾讯云短信小助手](https://cloud.tencent.com/document/product/382/3773#.E6.8A.80.E6.9C.AF.E4.BA.A4.E6.B5.81)。
	UNAUTHORIZEDOPERATION_REQUESTPERMISSIONDENY = "UnauthorizedOperation.RequestPermissionDeny"

	// 此 SdkAppId 禁止提供服务，如有需要请联系 [腾讯云短信小助手](https://cloud.tencent.com/document/product/382/3773#.E6.8A.80.E6.9C.AF.E4.BA.A4.E6.B5.81)。
	UNAUTHORIZEDOPERATION_SDKAPPIDISDISABLED = "UnauthorizedOperation.SdkAppIdIsDisabled"

	// 欠费被停止服务，可自行登录腾讯云充值来缴清欠款。
	UNAUTHORIZEDOPERATION_SERIVCESUSPENDDUETOARREARS = "UnauthorizedOperation.SerivceSuspendDueToArrears"

	// SmsSdkAppId 校验失败，请检查 [SmsSdkAppId](https://console.cloud.tencent.com/smsv2/app-manage) 是否属于 [云API密钥](https://console.cloud.tencent.com/cam/capi) 的关联账户。
	UNAUTHORIZEDOPERATION_SMSSDKAPPIDVERIFYFAIL = "UnauthorizedOperation.SmsSdkAppIdVerifyFail"

	// 不支持该请求。
	UNSUPPORTEDOPERATION_ = "UnsupportedOperation."

	// 国内短信模板不支持发送国际/港澳台手机号。发送国际/港澳台手机号请使用国际/港澳台短信正文模板。
	UNSUPPORTEDOPERATION_CHINESEMAINLANDTEMPLATETOGLOBALPHONE = "UnsupportedOperation.ChineseMainlandTemplateToGlobalPhone"

	// 群发请求里既有国内手机号也有国际手机号。请排查是否存在（1）使用国内签名或模板却发送短信到国际手机号；（2）使用国际签名或模板却发送短信到国内手机号。
	UNSUPPORTEDOPERATION_CONTAINDOMESTICANDINTERNATIONALPHONENUMBER = "UnsupportedOperation.ContainDomesticAndInternationalPhoneNumber"

	// 国际/港澳台短信模板不支持发送国内手机号。发送国内手机号请使用国内短信正文模板。
	UNSUPPORTEDOPERATION_GLOBALTEMPLATETOCHINESEMAINLANDPHONE = "UnsupportedOperation.GlobalTemplateToChineseMainlandPhone"

	// 不支持该地区短信下发。
	UNSUPPORTEDOPERATION_UNSUPORTEDREGION = "UnsupportedOperation.UnsuportedRegion"
)
