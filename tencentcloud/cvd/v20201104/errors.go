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

package v20201104

const (
	// 此产品的特有错误码

	// CAM签名/鉴权错误。
	AUTHFAILURE = "AuthFailure"

	// DryRun 操作，代表请求将会是成功的，只是多传了 DryRun 参数。
	DRYRUNOPERATION = "DryRunOperation"

	// 操作失败。
	FAILEDOPERATION = "FailedOperation"

	// 订单校验没有通过。
	FAILEDOPERATION_DEALCHECKFAILED = "FailedOperation.DealCheckFailed"

	// 实例使用中，必须以抢占式发起登录
	FAILEDOPERATION_INSTANCEINUSE = "FailedOperation.InstanceInUse"

	// 未查询到相关实例，请检查 实例 ID 入参
	FAILEDOPERATION_INSTANCENOTFOUND = "FailedOperation.InstanceNotFound"

	// 账户余额不足。
	FAILEDOPERATION_INSUFFICIENTBALANCE = "FailedOperation.InsufficientBalance"

	// 内部模块调用错误，稍后重试
	FAILEDOPERATION_INTERNALMODULEERROR = "FailedOperation.InternalModuleError"

	// 当前镜像不能适用全部重置的机器。
	FAILEDOPERATION_INVALIDIMAGESIZE = "FailedOperation.InvalidImageSize"

	// 实例绑定的用户不正确，请检查相关属性
	FAILEDOPERATION_INVALIDINSTANCEBOUNDUSER = "FailedOperation.InvalidInstanceBoundUser"

	// 实例区域属性错误，请检查相关属性
	FAILEDOPERATION_INVALIDINSTANCEREGION = "FailedOperation.InvalidInstanceRegion"

	// 实例当前状态不允许执行相关指令
	FAILEDOPERATION_INVALIDINSTANCESTATUSOPERATION = "FailedOperation.InvalidInstanceStatusOperation"

	// 请确认所有机器都处于运行状态。
	FAILEDOPERATION_INVALIDSTATUS = "FailedOperation.InvalidStatus"

	// 当前区域不支持
	FAILEDOPERATION_NOSUPPORTINSTANCEREGION = "FailedOperation.NoSupportInstanceRegion"

	// 库存不足
	FAILEDOPERATION_RESOURCEINSUFFICIENT = "FailedOperation.ResourceInsufficient"

	// 信令执行时发生错误,建议重试。
	FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"

	// 未知错误，请稍后重试
	FAILEDOPERATION_UNKNOWNERROR = "FailedOperation.UnknownError"

	// 当前用户未绑定相关资源
	FAILEDOPERATION_USERRESOURCENOTFOUND = "FailedOperation.UserResourceNotFound"

	// 内部错误。
	INTERNALERROR = "InternalError"

	// 无效的过滤器
	INVALIDFILTER = "InvalidFilter"

	// 参数错误。
	INVALIDPARAMETER = "InvalidParameter"

	// 脚本任务不是运行态
	INVALIDPARAMETER_COMMANDNOTRUNNING = "InvalidParameter.CommandNotRunning"

	// 搜索不到用户
	INVALIDPARAMETER_USERKEYWORD = "InvalidParameter.UserKeyword"

	// 参数取值错误。
	INVALIDPARAMETERVALUE = "InvalidParameterValue"

	// 硬盘相关的参数错误。
	INVALIDPARAMETERVALUE_DISK = "InvalidParameterValue.Disk"

	// 镜像相关的参数错误。
	INVALIDPARAMETERVALUE_IMAGE = "InvalidParameterValue.Image"

	// 实例ID参数错误。
	INVALIDPARAMETERVALUE_INSTANCEID = "InvalidParameterValue.InstanceId"

	// 参数取值错误，超过单次允许操作的最大数量。
	INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"

	// 资源不足
	INVALIDPARAMETERVALUE_RESOURCENOTENOUGH = "InvalidParameterValue.ResourceNotEnough"

	// 超过配额限制。
	LIMITEXCEEDED = "LimitExceeded"

	// 超出自定义镜像配额。
	LIMITEXCEEDED_IMAGEQUOTA = "LimitExceeded.ImageQuota"

	// 超过允许的配置数量。
	LIMITEXCEEDED_INSTANCEQUOTA = "LimitExceeded.InstanceQuota"

	// 并行执行的操作过多。
	LIMITEXCEEDED_PARALLELOPERATION = "LimitExceeded.ParallelOperation"

	// 可用策略已达上限。
	LIMITEXCEEDED_TOOMANYPOLICIES = "LimitExceeded.TooManyPolicies"

	// 缺少参数错误。
	MISSINGPARAMETER = "MissingParameter"

	// 操作被拒绝。
	OPERATIONDENIED = "OperationDenied"

	// 资源被占用。
	RESOURCEINUSE = "ResourceInUse"

	// 已存在相同的策略规则。
	RESOURCEINUSE_POLICYEXISTS = "ResourceInUse.PolicyExists"

	// 资源不足。
	RESOURCEINSUFFICIENT = "ResourceInsufficient"

	// 指定类型实例库存不足。
	RESOURCEINSUFFICIENT_SPECIFIEDINSTANCETYPE = "ResourceInsufficient.SpecifiedInstanceType"

	// 资源不存在。
	RESOURCENOTFOUND = "ResourceNotFound"

	// 未授权操作。
	UNAUTHORIZEDOPERATION = "UnauthorizedOperation"

	// 尚未开通该项服务，请先开通服务。
	UNAUTHORIZEDOPERATION_SERVICENOTACTIVATED = "UnauthorizedOperation.ServiceNotActivated"

	// 未知参数错误。
	UNKNOWNPARAMETER = "UnknownParameter"

	// 操作不支持。
	UNSUPPORTEDOPERATION = "UnsupportedOperation"

	// 检查桌面助手的状态未通过。
	UNSUPPORTEDOPERATION_ASSISTANTSTATE = "UnsupportedOperation.AssistantState"

	// 由于硬盘的过大不支持该操作。
	UNSUPPORTEDOPERATION_DISKSIZE = "UnsupportedOperation.DiskSize"

	// 由于镜像状态的原因不支持该操作。
	UNSUPPORTEDOPERATION_IMAGESTATE = "UnsupportedOperation.ImageState"

	// 由于镜像类型的原因不支持该操作。
	UNSUPPORTEDOPERATION_IMAGETYPE = "UnsupportedOperation.ImageType"

	// 实例当前的状态下不支持该操作。
	UNSUPPORTEDOPERATION_INSTANCESTATE = "UnsupportedOperation.InstanceState"

	// 不支持关机不收费特性。
	UNSUPPORTEDOPERATION_STOPPEDMODESTOPCHARGING = "UnsupportedOperation.StoppedModeStopCharging"
)
