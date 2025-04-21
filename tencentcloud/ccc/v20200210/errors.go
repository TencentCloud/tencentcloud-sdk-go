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

package v20200210

const (
	// 此产品的特有错误码

	// 无可用的AI并发许可，请去控制台进行购买
	FAILEDOPERATION_AICONCURRENTNOPACKAGE = "FailedOperation.AIConcurrentNoPackage"

	// AI并发超频
	FAILEDOPERATION_AICONCURRENTOVERFREQUENCY = "FailedOperation.AIConcurrentOverFrequency"

	// 外呼失败。
	FAILEDOPERATION_CALLOUTFAILED = "FailedOperation.CallOutFailed"

	// 高风险用户，请勿呼叫
	FAILEDOPERATION_CALLEEISBLACKUSER = "FailedOperation.CalleeIsBlackUser"

	// 被叫号码外呼受限。
	FAILEDOPERATION_CALLEEISLIMITED = "FailedOperation.CalleeIsLimited"

	// 主叫号码外呼超频。
	FAILEDOPERATION_CALLEROVERFREQUENCY = "FailedOperation.CallerOverFrequency"

	// 触发默认被叫规则，呼叫盲区
	FAILEDOPERATION_CALLOUTRULEBLINDAREA = "FailedOperation.CalloutRuleBlindArea"

	// 触发默认外呼规则，被叫一段时间内呼叫量
	FAILEDOPERATION_CALLOUTRULEMAXCALLCOUNTCALLEEINTERVALTIME = "FailedOperation.CalloutRuleMaxCallCountCalleeIntervalTime"

	// 触发默认外呼规则，被叫每日最大呼叫量
	FAILEDOPERATION_CALLOUTRULEMAXCALLCOUNTCALLEEPERDAYAPPID = "FailedOperation.CalloutRuleMaxCallCountCalleePerDayAppID"

	// 触发默认外呼规则，不在可外呼时间
	FAILEDOPERATION_CALLOUTRULENOTWORKTIME = "FailedOperation.CalloutRuleNotWorkTime"

	// 当前号码状态不能被修改。
	FAILEDOPERATION_CURSTATENOTALLOWMODIFY = "FailedOperation.CurStateNotAllowModify"

	// 重复账号。
	FAILEDOPERATION_DUPLICATEDACCOUNT = "FailedOperation.DuplicatedAccount"

	// 无可用的外呼号码。
	FAILEDOPERATION_NOCALLOUTNUMBER = "FailedOperation.NoCallOutNumber"

	// 权限不足。
	FAILEDOPERATION_PERMISSIONDENIED = "FailedOperation.PermissionDenied"

	// 坐席处于忙碌状态。
	FAILEDOPERATION_SEATSTATUSBUSY = "FailedOperation.SeatStatusBusy"

	// 会话不存在
	FAILEDOPERATION_SESSIONNOTEXISTS = "FailedOperation.SessionNotExists"

	// 上传文件个数超过限制
	FAILEDOPERATION_UPLOADFILEOVERFLOW = "FailedOperation.UploadFileOverflow"

	// 内部错误。
	INTERNALERROR = "InternalError"

	// 内部数据库访问失败。
	INTERNALERROR_DBERROR = "InternalError.DBError"

	// 参数错误。
	INVALIDPARAMETER = "InvalidParameter"

	// 重复的地址
	INVALIDPARAMETER_DUPLICATEADDRESS = "InvalidParameter.DuplicateAddress"

	// 重复的号码
	INVALIDPARAMETER_DUPLICATEPHONENUMBER = "InvalidParameter.DuplicatePhoneNumber"

	// 重复的SIP账号
	INVALIDPARAMETER_DUPLICATESIPACCOUNT = "InvalidParameter.DuplicateSipAccount"

	// 非法的地址
	INVALIDPARAMETER_ILLEGALADDRESS = "InvalidParameter.IllegalAddress"

	// 非法的号码
	INVALIDPARAMETER_ILLEGALPHONENUMBER = "InvalidParameter.IllegalPhoneNumber"

	// 实例不存在。
	INVALIDPARAMETER_INSTANCENOTEXIST = "InvalidParameter.InstanceNotExist"

	// 无效的地址
	INVALIDPARAMETER_INVALIDADDRESS = "InvalidParameter.InvalidAddress"

	// 无效的IP信息
	INVALIDPARAMETER_INVALIDIP = "InvalidParameter.InvalidIP"

	// 无效的号码
	INVALIDPARAMETER_INVALIDPHONENUMBER = "InvalidParameter.InvalidPhoneNumber"

	// 无效的端口信息
	INVALIDPARAMETER_INVALIDPORT = "InvalidParameter.InvalidPort"

	// 密码不合法(长度大于等于八位，必须包含大小写字母以及数字)
	INVALIDPARAMETER_SIPACCOUNTPASSWORDFORMAT = "InvalidParameter.SipAccountPasswordFormat"

	// 用户名不合法(只能包含A-Z、a-z、以及数字)
	INVALIDPARAMETER_SIPACCOUNTUSERFORMAT = "InvalidParameter.SipAccountUserFormat"

	// SIP通道仍在使用中
	INVALIDPARAMETER_SIPTRUNKINUSED = "InvalidParameter.SipTrunkInUsed"

	// 未找到SIP通道信息
	INVALIDPARAMETER_SIPTRUNKNOTFOUND = "InvalidParameter.SipTrunkNotFound"

	// 参数取值错误。
	INVALIDPARAMETERVALUE = "InvalidParameterValue"

	// 账号不存在。
	INVALIDPARAMETERVALUE_ACCOUNTNOTEXIST = "InvalidParameterValue.AccountNotExist"

	// 审批单号不存在。
	INVALIDPARAMETERVALUE_APPLYIDNOTEXIST = "InvalidParameterValue.ApplyIDNotExist"

	// 公司信息已经存在。
	INVALIDPARAMETERVALUE_COMPANYEXIST = "InvalidParameterValue.CompanyExist"

	// 审核单状态错误。
	INVALIDPARAMETERVALUE_ERRORAPPLYSTATUS = "InvalidParameterValue.ErrorApplyStatus"

	// 使用体验账号的智能体不支持此类操作
	INVALIDPARAMETERVALUE_EXPERIENCEACCOUNT = "InvalidParameterValue.ExperienceAccount"

	// 实例不存在。
	INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"

	// 号码状态无效。
	INVALIDPARAMETERVALUE_PHONENUMINVALID = "InvalidParameterValue.PhoneNumInvalid"

	// 号码已经绑定别的账号。
	INVALIDPARAMETERVALUE_PHONENUMISBOUNDOTHERACCOUNT = "InvalidParameterValue.PhoneNumIsBoundOtherAccount"

	// 查询记录不存在。
	INVALIDPARAMETERVALUE_RECORDNOTEXIST = "InvalidParameterValue.RecordNotExist"

	// 同振技能组人数超上限2，请确认并修改。
	INVALIDPARAMETERVALUE_RINGALLGROUPMEMBEROVERFLOW = "InvalidParameterValue.RingAllGroupMemberOverflow"

	// 技能组错误。
	INVALIDPARAMETERVALUE_SKILLGROUPERROR = "InvalidParameterValue.SkillGroupError"

	// 技能组已存在。
	INVALIDPARAMETERVALUE_SKILLGROUPEXIST = "InvalidParameterValue.SkillGroupExist"

	// 待审核单已经超限，暂时不能再提交
	INVALIDPARAMETERVALUE_WAITINGAPPROVALOVERFLOW = "InvalidParameterValue.WaitingApprovalOverflow"

	// 超过配额限制。
	LIMITEXCEEDED = "LimitExceeded"

	// 套餐包耗尽
	LIMITEXCEEDED_BASEPACKAGEEXPIRED = "LimitExceeded.BasePackageExpired"

	// 超出数量限制。
	LIMITEXCEEDED_OUTOFCOUNTLIMIT = "LimitExceeded.OutOfCountLimit"

	// 缺少参数错误。
	MISSINGPARAMETER = "MissingParameter"

	// 不在白名单中。
	OPERATIONDENIED_NOTINWHITELIST = "OperationDenied.NotInWhiteList"

	// 账号已被禁用。
	OPERATIONDENIED_UINDISABLED = "OperationDenied.UinDisabled"

	// 资源不存在。
	RESOURCENOTFOUND = "ResourceNotFound"

	// 资源不可用。
	RESOURCEUNAVAILABLE = "ResourceUnavailable"

	// 未知参数错误。
	UNKNOWNPARAMETER = "UnknownParameter"

	// 操作不支持。
	UNSUPPORTEDOPERATION = "UnsupportedOperation"
)
