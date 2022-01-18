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

package v20201002

const (
	// 此产品的特有错误码

	// 操作失败。
	FAILEDOPERATION = "FailedOperation"

	// 附件太大，请参考单个附件以及附件总量大小限制。
	FAILEDOPERATION_ATTACHCONTENTTOOLARGE = "FailedOperation.AttachContentToolarge"

	// 邮件地址在黑名单中。
	FAILEDOPERATION_EMAILADDRINBLACKLIST = "FailedOperation.EmailAddrInBlacklist"

	// 邮件内容太大，请减少发送邮件内容。
	FAILEDOPERATION_EMAILCONTENTTOOLARGE = "FailedOperation.EmailContentToolarge"

	// 超出当日总量发送限制。
	FAILEDOPERATION_EXCEEDSENDLIMIT = "FailedOperation.ExceedSendLimit"

	// 超出最大模版数量限制。
	FAILEDOPERATION_EXCEEDTEMPLATELIMIT = "FailedOperation.ExceedTemplateLimit"

	// 触发频率控制，短时间内对同一地址发送过多邮件。
	FAILEDOPERATION_FREQUENCYLIMIT = "FailedOperation.FrequencyLimit"

	// 拒信率过高，被临时block。
	FAILEDOPERATION_HIGHREJECTIONRATE = "FailedOperation.HighRejectionRate"

	// 邮箱地址错误。
	FAILEDOPERATION_INCORRECTEMAIL = "FailedOperation.IncorrectEmail"

	// 发件人格式不正确。请参考文档示例填写。
	FAILEDOPERATION_INCORRECTSENDER = "FailedOperation.IncorrectSender"

	// 余额不足，账号欠费等。
	FAILEDOPERATION_INSUFFICIENTBALANCE = "FailedOperation.InsufficientBalance"

	// 套餐包额度不足。
	FAILEDOPERATION_INSUFFICIENTQUOTA = "FailedOperation.InsufficientQuota"

	// 不支持的附件名称，请检查附件名称是否含有特殊字符，详情请参考附录附件说明。
	FAILEDOPERATION_INVALIDATTACHNAME = "FailedOperation.InvalidAttachName"

	// 超过查询限制，Limit最大支持100。
	FAILEDOPERATION_INVALIDLIMIT = "FailedOperation.InvalidLimit"

	// 模板ID无效或者不可用。
	FAILEDOPERATION_INVALIDTEMPLATEID = "FailedOperation.InvalidTemplateID"

	// 缺少发信内容（TemplateData和Simple不能同时为空)。
	FAILEDOPERATION_MISSINGEMAILCONTENT = "FailedOperation.MissingEmailContent"

	// 没有附件发送权限，请检查。
	FAILEDOPERATION_NOATTACHPERMISSION = "FailedOperation.NoAttachPermission"

	// 发件sender没有经过认证，无法发送。
	FAILEDOPERATION_NOTAUTHENTICATEDSENDER = "FailedOperation.NotAuthenticatedSender"

	// 不支持查询该日期邮件记录，仅支持查询90天之内的数据。
	FAILEDOPERATION_NOTSUPPORTDATE = "FailedOperation.NotSupportDate"

	// 协议检查错误，请检查协议是否正确。
	FAILEDOPERATION_PROTOCOLCHECKERR = "FailedOperation.ProtocolCheckErr"

	// 您的发送遇到问题，请联系腾讯云提交工单沟通原因并解决。
	FAILEDOPERATION_SENDEMAILERR = "FailedOperation.SendEmailErr"

	// 临时性错误，请求未生效，请重试。
	FAILEDOPERATION_SERVICENOTAVAILABLE = "FailedOperation.ServiceNotAvailable"

	// 模版内容太大，请减少模版内容。
	FAILEDOPERATION_TEMPLATECONTENTTOOLARGE = "FailedOperation.TemplateContentToolarge"

	// 因触发了某些规则导致临时Block。
	FAILEDOPERATION_TEMPORARYBLOCKED = "FailedOperation.TemporaryBlocked"

	// 附件数量太多，单封邮件最多支持10个附件。
	FAILEDOPERATION_TOOMANYATTACHMENTS = "FailedOperation.TooManyAttachments"

	// 收件人数太多，最多支持同时发送50人。
	FAILEDOPERATION_TOOMANYRECIPIENTS = "FailedOperation.TooManyRecipients"

	// 不支持的邮箱类型。
	FAILEDOPERATION_UNSUPPORTMAILTYPE = "FailedOperation.UnsupportMailType"

	// 未开通自定义发送权限，必须使用模版发送。
	FAILEDOPERATION_WITHOUTPERMISSION = "FailedOperation.WithOutPermission"

	// TemplateData字段格式不正确，请保持为json格式。
	FAILEDOPERATION_WRONGCONTENTJSON = "FailedOperation.WrongContentJson"

	// 内部错误。
	INTERNALERROR = "InternalError"

	// 参数错误。
	INVALIDPARAMETER = "InvalidParameter"

	// 参数取值错误。
	INVALIDPARAMETERVALUE = "InvalidParameterValue"

	// 附件内容错误，请检查base64内容是否正确。
	INVALIDPARAMETERVALUE_ATTACHCONTENTISWRONG = "InvalidParameterValue.AttachContentIsWrong"

	// 开始时间不能早于当前时间。
	INVALIDPARAMETERVALUE_BEGINTIMEBEFORENOW = "InvalidParameterValue.BeginTimeBeforeNow"

	// 该域名已被其他腾讯云账号创建。
	INVALIDPARAMETERVALUE_CREATEDBYOTHER = "InvalidParameterValue.CreatedByOther"

	// 收件或者发件地址为空，请检查邮件地址。
	INVALIDPARAMETERVALUE_EMAILADDRESSISNULL = "InvalidParameterValue.EmailAddressIsNULL"

	// 邮件内容错误，请检查TEXT/HTML base64是否正确。
	INVALIDPARAMETERVALUE_EMAILCONTENTISWRONG = "InvalidParameterValue.EmailContentIsWrong"

	// 邮件地址不合法，请检查邮件地址格式。
	INVALIDPARAMETERVALUE_ILLEGALEMAILADDRESS = "InvalidParameterValue.IllegalEmailAddress"

	// 域名取值错误。
	INVALIDPARAMETERVALUE_INVALIDEMAILIDENTITY = "InvalidParameterValue.InvalidEmailIdentity"

	// 您没有这个发件地址，请检查是否存在。
	INVALIDPARAMETERVALUE_NOSUCHSENDER = "InvalidParameterValue.NoSuchSender"

	// 该域名还没有创建，请先新建发信域名。
	INVALIDPARAMETERVALUE_NOTEXISTDOMAIN = "InvalidParameterValue.NotExistDomain"

	// 收件人列表描述不合法，描述不可包含特殊字符，字符长度范围:1~300。
	INVALIDPARAMETERVALUE_RECEIVERDESCILLEGAL = "InvalidParameterValue.ReceiverDescIllegal"

	// 收件人列表名字不合法，请检查字符内容及大小。
	INVALIDPARAMETERVALUE_RECEIVERNAMEILLEGAL = "InvalidParameterValue.ReceiverNameIllegal"

	// 该域名已经创建过了，请勿重复创建。
	INVALIDPARAMETERVALUE_REPEATCREATION = "InvalidParameterValue.RepeatCreation"

	// 该邮件地址已经创建过了，请勿重复创建。
	INVALIDPARAMETERVALUE_REPEATEMAILADDRESS = "InvalidParameterValue.RepeatEmailAddress"

	// 收件人列表名称重复。
	INVALIDPARAMETERVALUE_REPEATRECEIVERNAME = "InvalidParameterValue.RepeatReceiverName"

	// 模版资源为空。
	INVALIDPARAMETERVALUE_TEMPLATECONTENTISNULL = "InvalidParameterValue.TemplateContentIsNULL"

	// 模版内容错误，请检查base64内容是否正确。
	INVALIDPARAMETERVALUE_TEMPLATECONTENTISWRONG = "InvalidParameterValue.TemplateContentIsWrong"

	// 变量设置必须为json格式。
	INVALIDPARAMETERVALUE_TEMPLATEDATAERROR = "InvalidParameterValue.TemplateDataError"

	// 模版名字不合法，请检查字符内容及大小。
	INVALIDPARAMETERVALUE_TEMPLATENAMEILLEGAL = "InvalidParameterValue.TemplateNameIllegal"

	// 模版名字不能为空。
	INVALIDPARAMETERVALUE_TEMPLATENAMEISNULL = "InvalidParameterValue.TemplateNameIsNULL"

	// 不存在该模版，请先创建模版。
	INVALIDPARAMETERVALUE_TEMPLATENOTEXIST = "InvalidParameterValue.TemplateNotExist"

	// 搜索日期错误，请检查日期及格式是否有效。
	INVALIDPARAMETERVALUE_WRONGDATE = "InvalidParameterValue.WrongDate"

	// 超过配额限制。
	LIMITEXCEEDED = "LimitExceeded"

	// 收件人地址数量超过限制。
	LIMITEXCEEDED_EXCEEDRECEIVERDETAILLIMIT = "LimitExceeded.ExceedReceiverDetailLimit"

	// 收件人列表数量超过限制。
	LIMITEXCEEDED_EXCEEDRECEIVERLIMIT = "LimitExceeded.ExceedReceiverLimit"

	// 缺少参数错误。
	MISSINGPARAMETER = "MissingParameter"

	// 周期发送任务需选择开始时间和间隔时间。
	MISSINGPARAMETER_CYCLEPARAMNECESSARY = "MissingParameter.CycleParamNecessary"

	// 收件人地址必须传。
	MISSINGPARAMETER_EMAILSNECESSARY = "MissingParameter.EmailsNecessary"

	// 收件人列表Id必传。
	MISSINGPARAMETER_RECEIVERIDNECESSARY = "MissingParameter.ReceiverIdNecessary"

	// 发件地址、收件人列表,发件模板必须传。
	MISSINGPARAMETER_SENDPARAMNECESSARY = "MissingParameter.SendParamNecessary"

	// 定时发送任务需选择开始时间。
	MISSINGPARAMETER_TIMEDPARAMNECESSARY = "MissingParameter.TimedParamNecessary"

	// 操作被拒绝。
	OPERATIONDENIED = "OperationDenied"

	// 域名验证失败，请检查域名是否通过验证。
	OPERATIONDENIED_DOMAINNOTVERIFIED = "OperationDenied.DomainNotVerified"

	// 超出最大发信域名限制。
	OPERATIONDENIED_EXCEEDDOMAINLIMIT = "OperationDenied.ExceedDomainLimit"

	// 超出最大发信地址限制。
	OPERATIONDENIED_EXCEEDSENDERLIMIT = "OperationDenied.ExceedSenderLimit"

	// 收件人列表正在被操作，请稍候操作。
	OPERATIONDENIED_RECEIVERISOPERATING = "OperationDenied.ReceiverIsOperating"

	// 收件人列表不存在。
	OPERATIONDENIED_RECEIVERNOTEXIST = "OperationDenied.ReceiverNotExist"

	// 收件人列表空或状态不是上传完成。
	OPERATIONDENIED_RECEIVERSTATUSERROR = "OperationDenied.ReceiverStatusError"

	// 发信地址不存在或者状态不是通过状态。
	OPERATIONDENIED_SENDADDRESSSTATUSERROR = "OperationDenied.SendAddressStatusError"

	// 发信模板不存在或者状态不是审核通过状态。
	OPERATIONDENIED_TEMPLATESTATUSERROR = "OperationDenied.TemplateStatusError"

	// 请求的次数超过了频率限制。
	REQUESTLIMITEXCEEDED = "RequestLimitExceeded"

	// 资源不足。
	RESOURCEINSUFFICIENT = "ResourceInsufficient"

	// 资源不存在。
	RESOURCENOTFOUND = "ResourceNotFound"

	// 资源不可用。
	RESOURCEUNAVAILABLE = "ResourceUnavailable"

	// 未知参数错误。
	UNKNOWNPARAMETER = "UnknownParameter"

	// 操作不支持。
	UNSUPPORTEDOPERATION = "UnsupportedOperation"
)
