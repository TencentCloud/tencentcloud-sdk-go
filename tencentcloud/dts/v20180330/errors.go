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

package v20180330

const (
	// 此产品的特有错误码

	// 认证失败。
	AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"

	// 鉴权失败，当前用户不允许执行该操作。
	AUTHFAILURE_UNAUTHORIZEDOPERATIONERROR = "AuthFailure.UnauthorizedOperationError"

	// 操作失败。
	FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"

	// 禁止该操作。
	FAILEDOPERATION_NOTALLOWOPERATION = "FailedOperation.NotAllowOperation"

	// 代理出错。
	FAILEDOPERATION_PROXYERROR = "FailedOperation.ProxyError"

	// 任务开始失败。
	FAILEDOPERATION_STARTJOBFAILED = "FailedOperation.StartJobFailed"

	// 当前状态冲突，不能执行该操作。
	FAILEDOPERATION_STATUSINCONFLICT = "FailedOperation.StatusInConflict"

	// 当前状态不允许操作。
	FAILEDOPERATION_STATUSINCONFLICTERROR = "FailedOperation.StatusInConflictError"

	// 内部错误。
	INTERNALERROR = "InternalError"

	// 添加异步任务失败。
	INTERNALERROR_ADDTASKERROR = "InternalError.AddTaskError"

	// 内部调度系统错误。
	INTERNALERROR_CELERYERROR = "InternalError.CeleryError"

	// cgw系统错误。
	INTERNALERROR_CGWSYSTEMERROR = "InternalError.CgwSystemError"

	// 迁移平台数据库访问失败。
	INTERNALERROR_DATABASEERROR = "InternalError.DatabaseError"

	// 迁移任务冲突。
	INTERNALERROR_DUPLICATEJOB = "InternalError.DuplicateJob"

	// 内部错误。
	INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"

	// http请求访问出错。
	INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"

	// 内部组件访问错误。
	INTERNALERROR_INTERNALINNERCOMMONERROR = "InternalError.InternalInnerCommonError"

	// 调用计费服务失败。
	INTERNALERROR_INTERNALTRADEERROR = "InternalError.InternalTradeError"

	// 锁冲突。
	INTERNALERROR_LOCKERROR = "InternalError.LockError"

	// 通信协议错误。
	INTERNALERROR_PROTOCOLERROR = "InternalError.ProtocolError"

	// 内部错误。
	INTERNALERROR_UNDEFINEDERROR = "InternalError.UndefinedError"

	// 未知的内部错误。
	INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"

	// 参数错误。
	INVALIDPARAMETER = "InvalidParameter"

	// 参数值错误。
	INVALIDPARAMETER_BIZINVALIDPARAMETERVALUEERROR = "InvalidParameter.BizInvalidParameterValueError"

	// 实例不存在。
	INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"

	// ip地址错误。
	INVALIDPARAMETER_INVALIDIPADDRESS = "InvalidParameter.InvalidIpAddress"

	// 参数无效。
	INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"

	// 业务参数错误。
	INVALIDPARAMETERVALUE_BIZINVALIDPARAMETERVALUEERROR = "InvalidParameterValue.BizInvalidParameterValueError"

	// 数据转换错误。
	INVALIDPARAMETERVALUE_DATACONVERTERROR = "InvalidParameterValue.DataConvertError"

	// 参数值错误。
	INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"

	// 非法参数。
	INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"

	// 闲置迁移任务数目超过限制。
	LIMITEXCEEDED_MAXUNUSEDJOBS = "LimitExceeded.MaxUnusedJobs"

	// 操作被拒绝。
	OPERATIONDENIED = "OperationDenied"

	// 当前操作不满足条件被阻止。
	OPERATIONDENIED_BIZOPERATIONDENIEDERROR = "OperationDenied.BizOperationDeniedError"

	// 任务操作失败。
	OPERATIONDENIED_JOBOPERATIONDENIEDERROR = "OperationDenied.JobOperationDeniedError"

	// 数据迁移服务不支持当前迁移类型。
	OPERATIONDENIED_MIGRATESERVICESUPPORTERROR = "OperationDenied.MigrateServiceSupportError"

	// 用户余额不足。
	OPERATIONDENIED_NOTENOUGHMONEYERROR = "OperationDenied.NotEnoughMoneyError"

	// 该操作不能执行。
	OPERATIONDENIED_OPERATIONDENIED = "OperationDenied.OperationDenied"

	// 操作被拒绝。
	OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"

	// 请求次数超限。
	REQUESTLIMITEXCEEDED_REQUESTLIMITEXCEEDEDERROR = "RequestLimitExceeded.RequestLimitExceededError"

	// 资源不存在。
	RESOURCENOTFOUND = "ResourceNotFound"

	// 未找到资源。
	RESOURCENOTFOUND_BIZRESOURCENOTFOUNDERROR = "ResourceNotFound.BizResourceNotFoundError"

	// 迁移任务不存在。
	RESOURCENOTFOUND_JOBNOTEXIST = "ResourceNotFound.JobNotExist"

	// 实例未找到。
	RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"

	// 资源未找到。
	RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"

	// 认证失败，没有足够权限。
	UNAUTHORIZEDOPERATION_NOTENOUGHPRIVILEGES = "UnauthorizedOperation.NotEnoughPrivileges"

	// 未授权的操作。
	UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"

	// 操作不支持。
	UNSUPPORTEDOPERATION = "UnsupportedOperation"

	// 自研上云用户创建迁移任务时未指定标签，需要补齐“运营部门”、“运营产品”、“负责人”这三类标签。
	UNSUPPORTEDOPERATION_INTRANETUSERNOTTAGGEDERROR = "UnsupportedOperation.IntraNetUserNotTaggedError"
)
