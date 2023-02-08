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

package v20220308

const (
	// 此产品的特有错误码

	// 实例安全组信息添加失败。
	FAILEDOPERATION_ADDINSTANCEINFOFAILED = "FailedOperation.AddInstanceInfoFailed"

	// 绑定安全组失败。
	FAILEDOPERATION_ASSOCIATESECURITYGROUPSFAILED = "FailedOperation.AssociateSecurityGroupsFailed"

	// 实例安全组信息清除失败。
	FAILEDOPERATION_CLEARINSTANCEINFOFAILED = "FailedOperation.ClearInstanceInfoFailed"

	// 解绑安全组失败。
	FAILEDOPERATION_DISASSOCIATESECURITYGROUPSFAILED = "FailedOperation.DisassociateSecurityGroupsFailed"

	// 当前DTS状态下不允许该操作。
	FAILEDOPERATION_DTSSTATUSABNORMAL = "FailedOperation.DtsStatusAbnormal"

	// 获取安全组详情失败。
	FAILEDOPERATION_GETSECURITYGROUPDETAILFAILED = "FailedOperation.GetSecurityGroupDetailFailed"

	// 设置规则失败。
	FAILEDOPERATION_SETRULELOCATIONFAILED = "FailedOperation.SetRuleLocationFailed"

	// 内部系统错误，和业务无关。
	FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"

	// 实例不支持该接口。
	FAILEDOPERATION_UNSUPPORTERROR = "FailedOperation.UnSupportError"

	// 未知类型操作错误。
	FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"

	// 实例安全组信息更新失败。
	FAILEDOPERATION_UPDATEINSTANCEINFOFAILED = "FailedOperation.UpdateInstanceInfoFailed"

	// 更新安全组失败。
	FAILEDOPERATION_UPDATESECURITYGROUPSFAILED = "FailedOperation.UpdateSecurityGroupsFailed"

	// cam鉴权错误。
	INTERNALERROR_CAMAUTHOSSRESPONSERETURNCODEERROR = "InternalError.CamAuthOssResponseReturnCodeError"

	// 统一的 DB 操作错误，可以是 update insert select..。
	INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"

	// 执行Http请求失败。
	INTERNALERROR_EXECHTTPREQUESTERROR = "InternalError.ExecHttpRequestError"

	// 无操作权限。
	INTERNALERROR_INSTANCEOPERATEPERMISSIONERROR = "InternalError.InstanceOperatePermissionError"

	// 内部错误。
	INTERNALERROR_INTERNALERROR = "InternalError.InternalError"

	// 获取实例列表出错。
	INTERNALERROR_LISTINSTANCESERROR = "InternalError.ListInstancesError"

	// 网络错误。
	INTERNALERROR_NETWORKERR = "InternalError.NetWorkErr"

	// 参数错误。
	INVALIDPARAMETER = "InvalidParameter"

	// 参数为空。
	INVALIDPARAMETER_EMPTYPARAM = "InvalidParameter.EmptyParam"

	// 非法参数错误。
	INVALIDPARAMETER_ILLEGALPARAMETERERROR = "InvalidParameter.IllegalParameterError"

	// 业务参数错误。
	INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"

	// 参数错误，不支持操作。
	INVALIDPARAMETER_NOTSUPPORTED = "InvalidParameter.NotSupported"

	// 时间格式或者范围不符合要求。
	INVALIDPARAMETER_PERIOD = "InvalidParameter.Period"

	// 接口没有cam权限。
	INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"

	// 不支持的实例规格。
	INVALIDPARAMETER_SPECNOTSUPPORTED = "InvalidParameter.SpecNotSupported"

	// 密码为空。
	INVALIDPARAMETERVALUE_PASSWORDEMPTY = "InvalidParameterValue.PasswordEmpty"

	// 密码校验出错，密码错误。
	INVALIDPARAMETERVALUE_PASSWORDERROR = "InvalidParameterValue.PasswordError"

	// 腾讯集团内部账号禁止使用免密实例。
	INVALIDPARAMETERVALUE_PASSWORDFREEDENIED = "InvalidParameterValue.PasswordFreeDenied"

	// 设置密码时，MC 传入的 old password 与先前设定密码不同。
	INVALIDPARAMETERVALUE_PASSWORDRULEERROR = "InvalidParameterValue.PasswordRuleError"

	// 实例类型不支持。
	INVALIDPARAMETERVALUE_UNSUPPORTEDTYPE = "InvalidParameterValue.UnSupportedType"

	// weekday输入无效数据。
	INVALIDPARAMETERVALUE_WEEKDAYSISINVALID = "InvalidParameterValue.WeekDaysIsInvalid"

	// 请求的容量不在售卖规格中。
	LIMITEXCEEDED_INVALIDSIZE = "LimitExceeded.InvalidSize"

	// 资源不足。
	RESOURCEINSUFFICIENT = "ResourceInsufficient"

	// 找不到对应实例
	RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"

	// 找不到该实例。
	RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"

	// 调用后端接口失败。
	RESOURCEUNAVAILABLE_CALLOSSERROR = "ResourceUnavailable.CallOssError"

	// 实例配置错误。
	RESOURCEUNAVAILABLE_INSTANCECONFERROR = "ResourceUnavailable.InstanceConfError"

	// 实例已经被其它流程锁定。
	RESOURCEUNAVAILABLE_INSTANCELOCKEDERROR = "ResourceUnavailable.InstanceLockedError"

	// 实例状态错误。
	RESOURCEUNAVAILABLE_INSTANCESTATEERROR = "ResourceUnavailable.InstanceStateError"

	// 实例状态异常，不能执行对应流程。
	RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"

	// 实例状态异常，不能执行对应操作。
	RESOURCEUNAVAILABLE_INSTANCESTATUSERROR = "ResourceUnavailable.InstanceStatusError"

	// 未授权操作。
	UNAUTHORIZEDOPERATION = "UnauthorizedOperation"

	// 用户不在白名单中。
	UNAUTHORIZEDOPERATION_USERNOTINWHITELIST = "UnauthorizedOperation.UserNotInWhiteList"

	// 操作不支持。
	UNSUPPORTEDOPERATION = "UnsupportedOperation"

	// 自动续费标识错误。
	UNSUPPORTEDOPERATION_ISAUTORENEWERROR = "UnsupportedOperation.IsAutoRenewError"
)
