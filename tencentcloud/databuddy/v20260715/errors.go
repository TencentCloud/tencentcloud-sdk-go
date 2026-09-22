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

package v20260715

const (
	// 此产品的特有错误码

	// CAM签名/鉴权错误。
	AUTHFAILURE = "AuthFailure"

	// 操作失败。
	FAILEDOPERATION = "FailedOperation"

	// 调用 <serviceName> 服务的接口 <apiName> 失败：<message>。
	FAILEDOPERATION_CALLTHIRDPARTAPIERROR = "FailedOperation.CallThirdPartApiError"

	// 创建工作流失败
	FAILEDOPERATION_CREATEWORKFLOWFAILED = "FailedOperation.CreateWorkflowFailed"

	// 存在活跃的工作流运行，无法操作
	FAILEDOPERATION_EXISTWORKFLOWEXECUTIONS = "FailedOperation.ExistWorkflowExecutions"

	// 标签数量已达到最大允许限制
	FAILEDOPERATION_LABELCOUNTLIMIT = "FailedOperation.LabelCountLimit"

	// 无需操作的工作流运行
	FAILEDOPERATION_NOWORKFLOWEXECUTIONNEEDOPERATE = "FailedOperation.NoWorkflowExecutionNeedOperate"

	// 重跑工作流失败
	FAILEDOPERATION_RERUNWORKFLOWFAIL = "FailedOperation.RerunWorkflowFail"

	// 运行工作流失败
	FAILEDOPERATION_RUNWORKFLOWFAIL = "FailedOperation.RunWorkflowFail"

	// 运行工作流返回的执行ID为空
	FAILEDOPERATION_RUNWORKFLOWFAILEXECUTIONIDEMPTY = "FailedOperation.RunWorkflowFailExecutionIdEmpty"

	// 更新工作流失败
	FAILEDOPERATION_UPDATEWORKFLOWFAILED = "FailedOperation.UpdateWorkflowFailed"

	// 该工作流为通过bundle包部署的内容，请勿编辑此工作流
	FAILEDOPERATION_WORKFLOWBUNDLENOPERMISSION = "FailedOperation.WorkflowBundleNoPermission"

	// 工作流数量超过10000上限
	FAILEDOPERATION_WORKFLOWCOUNTLIMIT = "FailedOperation.WorkflowCountLimit"

	// 获取工作流名称分布式锁失败
	FAILEDOPERATION_WORKFLOWCREATELOCKACQUIREFAILED = "FailedOperation.WorkflowCreateLockAcquireFailed"

	// 工作流运行已被删除
	FAILEDOPERATION_WORKFLOWEXECUTIONHASBEDELETE = "FailedOperation.WorkflowExecutionHasBeDelete"

	// 无操作权限
	FAILEDOPERATION_WORKFLOWNOPERMISSION = "FailedOperation.WorkflowNoPermission"

	// 内部错误。
	INTERNALERROR = "InternalError"

	// 未知错误
	INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"

	// 参数错误。
	INVALIDPARAMETER = "InvalidParameter"

	// 无效参数
	INVALIDPARAMETER_INVALIDPARAMSERROR = "InvalidParameter.InvalidParamsError"

	// 参数取值错误。
	INVALIDPARAMETERVALUE = "InvalidParameterValue"

	// 同一工作流内存在重名任务
	INVALIDPARAMETERVALUE_DUPLICATETASKNAMEERROR = "InvalidParameterValue.DuplicateTaskNameError"

	// 参数错误
	INVALIDPARAMETERVALUE_INVALIDPARAMETER = "InvalidParameterValue.InvalidParameter"

	// 工作流列表筛选字段Key值错误
	INVALIDPARAMETERVALUE_LISTWORKFLOWFILTERPARAMERROR = "InvalidParameterValue.ListWorkflowFilterParamError"

	// LoopDataArray 数组元素数量超过了最大允许值
	INVALIDPARAMETERVALUE_LOOPDATAARRAYELEMENTCOUNTLIMIT = "InvalidParameterValue.LoopDataArrayElementCountLimit"

	// LoopDataArray 值既不是 JSON 数组字面量，也不是 {{var}} 占位符表达式
	INVALIDPARAMETERVALUE_LOOPDATAARRAYJSONINVALID = "InvalidParameterValue.LoopDataArrayJsonInvalid"

	// LoopDataArray 值解析成功但顶层节点不是 JSON 数组
	INVALIDPARAMETERVALUE_LOOPDATAARRAYNOTJSONARRAY = "InvalidParameterValue.LoopDataArrayNotJsonArray"

	// LoopDataArray 值既不是 JSON 数组字面量，也不是 {{var}} 占位符表达式
	INVALIDPARAMETERVALUE_LOOPDATAARRAYNOTJSONARRAYLITERAL = "InvalidParameterValue.LoopDataArrayNotJsonArrayLiteral"

	// FOR_EACH 任务的 LoopDataArray 字段为空或仅含空白字符
	INVALIDPARAMETERVALUE_LOOPDATAARRAYVALUEBLANK = "InvalidParameterValue.LoopDataArrayValueBlank"

	// FOR_EACH 任务的 LoopDataArray 字段字符串长度超过了最大允许值
	INVALIDPARAMETERVALUE_LOOPDATAARRAYVALUELENGTHLIMIT = "InvalidParameterValue.LoopDataArrayValueLengthLimit"

	// LoopDataArray 变量表达式不满足占位符语法（大括号需成对配对且不嵌套）
	INVALIDPARAMETERVALUE_LOOPDATAARRAYVARIABLEEXPRESSIONINVALID = "InvalidParameterValue.LoopDataArrayVariableExpressionInvalid"

	// 参数 <parameter> 不能为空。
	INVALIDPARAMETERVALUE_PARAMBLANKERROR = "InvalidParameterValue.ParamBlankError"

	// 参数 <parameter> 不符合要求：<message>
	INVALIDPARAMETERVALUE_PARAMILLEGALERROR = "InvalidParameterValue.ParamIllegalError"

	// 参数 <parameter> 不能为 null。
	INVALIDPARAMETERVALUE_PARAMNULLERROR = "InvalidParameterValue.ParamNullError"

	// 引用对象不存在或已被删除
	INVALIDPARAMETERVALUE_TASKHOOKVALIDATIONFAILED = "InvalidParameterValue.TaskHookValidationFailed"

	// 任务名包含非法字符
	INVALIDPARAMETERVALUE_TASKNAMECONTAINSILLEGALCHARACTERSERROR = "InvalidParameterValue.TaskNameContainsIllegalCharactersError"

	// 任务名超过128字符限制
	INVALIDPARAMETERVALUE_TASKNAMEEXCEEDSLIMITERROR = "InvalidParameterValue.TaskNameExceedsLimitError"

	// 请移除该任务类型的资源组配置
	INVALIDPARAMETERVALUE_TASKTYPENOTSUPPORTRESOURCEGROUP = "InvalidParameterValue.TaskTypeNotSupportResourceGroup"

	// 任务类型属性值在入参时为必填，但是并没有传递
	INVALIDPARAMETERVALUE_TASKTYPEPROPERTYKEYVALUEREQUESTREQUIREDERROR = "InvalidParameterValue.TaskTypePropertyKeyValueRequestRequiredError"

	// 结束时间不能早于开始时间
	INVALIDPARAMETERVALUE_WORKFLOWENDTIMELESSSTARTTIME = "InvalidParameterValue.WorkflowEndTimeLessStartTime"

	// 工作流名称重复
	INVALIDPARAMETERVALUE_WORKFLOWNAMEEXISTS = "InvalidParameterValue.WorkflowNameExists"

	// 工作流名称不合法
	INVALIDPARAMETERVALUE_WORKFLOWNAMEINVALID = "InvalidParameterValue.WorkflowNameInvalid"

	// 查询时间范围超过60天限制
	INVALIDPARAMETERVALUE_WORKFLOWQUERYENDTIMEANDSTARTTIMEEXCEED = "InvalidParameterValue.WorkflowQueryEndTimeAndStartTimeExceed"

	// 工作流调度开始时间不能晚于结束时间
	INVALIDPARAMETERVALUE_WORKFLOWSTARTTIMEAFTERENDTIMEERROR = "InvalidParameterValue.WorkflowStartTimeAfterEndTimeError"

	// 请检查工作流触发器高级配置，修正非法的 JSON 内容
	INVALIDPARAMETERVALUE_WORKFLOWTRIGGERADVANCEDCONFIGERROR = "InvalidParameterValue.WorkflowTriggerAdvancedConfigError"

	// 缺少参数错误。
	MISSINGPARAMETER = "MissingParameter"

	// 缺少必填参数
	MISSINGPARAMETER_MISSINGPARAMETER = "MissingParameter.MissingParameter"

	// 操作被拒绝。
	OPERATIONDENIED = "OperationDenied"

	// 地域错误
	REGIONERROR = "RegionError"

	// 资源不存在。
	RESOURCENOTFOUND = "ResourceNotFound"

	// 资源不存在或已被删除
	RESOURCENOTFOUND_ONEFLOWRESOURCENOEXISTERROR = "ResourceNotFound.OneFlowResourceNoExistError"

	// 任务运行不存在
	RESOURCENOTFOUND_TASKEXECUTIONNOTEXIST = "ResourceNotFound.TaskExecutionNotExist"

	// 工作流运行不存在
	RESOURCENOTFOUND_WORKFLOWEXECUTIONNOTEXIST = "ResourceNotFound.WorkflowExecutionNotExist"

	// 工作流不存在
	RESOURCENOTFOUND_WORKFLOWNOTEXIST = "ResourceNotFound.WorkflowNotExist"

	// 根据传入的工作流ID未找到工作流
	RESOURCENOTFOUND_WORKFLOWNOTFOUND = "ResourceNotFound.WorkflowNotFound"

	// 工作流下面的任务不存在
	RESOURCENOTFOUND_WORKFLOWTASKNOTEXIST = "ResourceNotFound.WorkflowTaskNotExist"

	// 工作流的调度配置未找到
	RESOURCENOTFOUND_WORKFLOWTRIGGERNOTFOUND = "ResourceNotFound.WorkflowTriggerNotFound"

	// 未知参数错误。
	UNKNOWNPARAMETER = "UnknownParameter"

	// 工作流运行已到达终态，不可终止
	UNSUPPORTEDOPERATION_WORKFLOWEXECUTIONHASREACHEDFINALSTATECANNOTBESTOPPED = "UnsupportedOperation.WorkflowExecutionHasReachedFinalStateCannotBeStopped"
)
