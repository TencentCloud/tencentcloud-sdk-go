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

package v20201028

const (
	// 此产品的特有错误码

	// CAM签名/鉴权错误。
	AUTHFAILURE = "AuthFailure"

	// 调用 CVM 失败。
	FAILEDOPERATION_CVMERROR = "FailedOperation.CVMError"

	// 调用 Lighthouse 失败。
	FAILEDOPERATION_LIGHTHOUSEERROR = "FailedOperation.LighthouseError"

	// 内部错误。
	INTERNALERROR = "InternalError"

	// 参数错误。
	INVALIDPARAMETER = "InvalidParameter"

	// 参数冲突。
	INVALIDPARAMETER_CONFLICTPARAMETER = "InvalidParameter.ConflictParameter"

	// 参数取值错误。
	INVALIDPARAMETERVALUE = "InvalidParameterValue"

	// Agent不支持此命令类型。
	INVALIDPARAMETERVALUE_AGENTUNSUPPORTEDCOMMANDTYPE = "InvalidParameterValue.AgentUnsupportedCommandType"

	// Command 内容无效。
	INVALIDPARAMETERVALUE_COMMANDCONTENTINVALID = "InvalidParameterValue.CommandContentInvalid"

	// Command 名称重复。
	INVALIDPARAMETERVALUE_COMMANDNAMEDUPLICATED = "InvalidParameterValue.CommandNameDuplicated"

	// 实例类型不一致。
	INVALIDPARAMETERVALUE_INCONSISTENTINSTANCE = "InvalidParameterValue.InconsistentInstance"

	// 实例ID与执行活动无关。
	INVALIDPARAMETERVALUE_INSTANCEISNOTRELATEDTOINVOCATION = "InvalidParameterValue.InstanceIsNotRelatedToInvocation"

	// CommandId 无效。
	INVALIDPARAMETERVALUE_INVALIDCOMMANDID = "InvalidParameterValue.InvalidCommandId"

	// Command 名称无效。
	INVALIDPARAMETERVALUE_INVALIDCOMMANDNAME = "InvalidParameterValue.InvalidCommandName"

	// 命令内容无效。
	INVALIDPARAMETERVALUE_INVALIDCONTENT = "InvalidParameterValue.InvalidContent"

	// Crontab 表达式无效。
	INVALIDPARAMETERVALUE_INVALIDCRONEXPRESSION = "InvalidParameterValue.InvalidCronExpression"

	// Filter 无效。
	INVALIDPARAMETERVALUE_INVALIDFILTER = "InvalidParameterValue.InvalidFilter"

	// 实例ID无效。
	INVALIDPARAMETERVALUE_INVALIDINSTANCEID = "InvalidParameterValue.InvalidInstanceId"

	// 不合法的执行活动ID。
	INVALIDPARAMETERVALUE_INVALIDINVOCATIONID = "InvalidParameterValue.InvalidInvocationId"

	// 不合法的执行任务ID。
	INVALIDPARAMETERVALUE_INVALIDINVOCATIONTASKID = "InvalidParameterValue.InvalidInvocationTaskId"

	// InvokerId 无效。
	INVALIDPARAMETERVALUE_INVALIDINVOKERID = "InvalidParameterValue.InvalidInvokerId"

	// OutputCOSBucketUrl 无效。
	INVALIDPARAMETERVALUE_INVALIDOUTPUTCOSBUCKETURL = "InvalidParameterValue.InvalidOutputCOSBucketUrl"

	// OutputCOSKeyPrefix 无效。
	INVALIDPARAMETERVALUE_INVALIDOUTPUTCOSKEYPREFIX = "InvalidParameterValue.InvalidOutputCOSKeyPrefix"

	// 命令执行路径不合法。
	INVALIDPARAMETERVALUE_INVALIDWORKINGDIRECTORY = "InvalidParameterValue.InvalidWorkingDirectory"

	// 已启用自定义参数功能，但缺失自定义参数信息。
	INVALIDPARAMETERVALUE_LACKOFPARAMETERINFO = "InvalidParameterValue.LackOfParameterInfo"

	// 未提供 Parameters 信息。
	INVALIDPARAMETERVALUE_LACKOFPARAMETERS = "InvalidParameterValue.LackOfParameters"

	// 超过参数限制。
	INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"

	// 未启用自定义参数功能。
	INVALIDPARAMETERVALUE_PARAMETERDISABLED = "InvalidParameterValue.ParameterDisabled"

	// 参数为非法 json string 格式。
	INVALIDPARAMETERVALUE_PARAMETERINVALIDJSONFORMAT = "InvalidParameterValue.ParameterInvalidJsonFormat"

	// 参数 Key 包含非法字符。
	INVALIDPARAMETERVALUE_PARAMETERKEYCONTAINSINVALIDCHAR = "InvalidParameterValue.ParameterKeyContainsInvalidChar"

	// 参数 Key 重复。
	INVALIDPARAMETERVALUE_PARAMETERKEYDUPLICATED = "InvalidParameterValue.ParameterKeyDuplicated"

	// 参数 Key 过长。
	INVALIDPARAMETERVALUE_PARAMETERKEYLENEXCEEDED = "InvalidParameterValue.ParameterKeyLenExceeded"

	// 参数数目过多。
	INVALIDPARAMETERVALUE_PARAMETERNUMBEREXCEEDED = "InvalidParameterValue.ParameterNumberExceeded"

	// 参数 Value 非 string 类型。
	INVALIDPARAMETERVALUE_PARAMETERVALUENOTSTRING = "InvalidParameterValue.ParameterValueNotString"

	// 参数取值范围不合法。
	INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"

	// 未启用自定义参数功能。
	INVALIDPARAMETERVALUE_SUPPORTPARAMETERSONLYIFENABLEPARAMETER = "InvalidParameterValue.SupportParametersOnlyIfEnableParameter"

	// 长度超过限制。
	INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"

	// 超过配额限制。
	LIMITEXCEEDED = "LimitExceeded"

	// 缺少参数错误。
	MISSINGPARAMETER = "MissingParameter"

	// 资源不存在。
	RESOURCENOTFOUND = "ResourceNotFound"

	// 命令不存在。
	RESOURCENOTFOUND_COMMANDNOTFOUND = "ResourceNotFound.CommandNotFound"

	// 实例不存在。
	RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"

	// 执行活动未找到。
	RESOURCENOTFOUND_INVOCATIONNOTFOUND = "ResourceNotFound.InvocationNotFound"

	// Agent 未安装。
	RESOURCEUNAVAILABLE_AGENTNOTINSTALLED = "ResourceUnavailable.AgentNotInstalled"

	// Agent 不在线。
	RESOURCEUNAVAILABLE_AGENTSTATUSNOTONLINE = "ResourceUnavailable.AgentStatusNotOnline"

	// 命令正在执行中。
	RESOURCEUNAVAILABLE_COMMANDINEXECUTING = "ResourceUnavailable.CommandInExecuting"

	// 命令已关联执行器。
	RESOURCEUNAVAILABLE_COMMANDININVOKER = "ResourceUnavailable.CommandInInvoker"

	// 实例未处于运行中。
	RESOURCEUNAVAILABLE_INSTANCESTATENOTRUNNING = "ResourceUnavailable.InstanceStateNotRunning"

	// CAM鉴权失败。
	UNAUTHORIZEDOPERATION_CAMAUTHFAILED = "UnauthorizedOperation.CamAuthFailed"

	// Token 无效。
	UNAUTHORIZEDOPERATION_INVALIDTOKEN = "UnauthorizedOperation.InvalidToken"

	// Multi-Factor Authentication(MFA) 过期。
	UNAUTHORIZEDOPERATION_MFAEXPIRED = "UnauthorizedOperation.MFAExpired"

	// Multi-Factor Authentication(MFA) 不存在。
	UNAUTHORIZEDOPERATION_MFANOTFOUND = "UnauthorizedOperation.MFANotFound"

	// 未知参数错误。
	UNKNOWNPARAMETER = "UnknownParameter"
)
