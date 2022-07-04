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

package v20210601

const (
	// 此产品的特有错误码

	// 操作失败。
	FAILEDOPERATION = "FailedOperation"

	// 鉴权失败。
	FAILEDOPERATION_AUTHENTICATIONFAILED = "FailedOperation.AuthenticationFailed"

	// 接口处理超时。
	FAILEDOPERATION_INNERLOGICTIMEOUT = "FailedOperation.InnerLogicTimeOut"

	// 依赖服务错误。
	FAILEDOPERATION_METACOMPILERERROR = "FailedOperation.MetaCompilerError"

	// 不支持的操作类型。
	FAILEDOPERATION_UNSUPPORTEDOPERATIONTYPE = "FailedOperation.UnSupportedOperationType"

	// 检查运行时失败。
	INTERNALERROR_CHECKRUNTIMEFAILED = "InternalError.CheckRuntimeFailed"

	// 获取运行时应用数量失败。
	INTERNALERROR_COUNTRUNTIMEINSTANCESFAILED = "InternalError.CountRuntimeInstancesFailed"

	// 数据库内部错误。
	INTERNALERROR_DATABASEERROR = "InternalError.DatabaseError"

	// 获取运行时应用列表失败。
	INTERNALERROR_LISTRUNTIMEINSTANCESFAILED = "InternalError.ListRuntimeInstancesFailed"

	// 获取运行时列表失败。
	INTERNALERROR_LISTRUNTIMESFAILED = "InternalError.ListRuntimesFailed"

	// 依赖模块metacompiler错误。
	INTERNALERROR_METACOMPILERERROR = "InternalError.MetaCompilerError"

	// rpc调用异常。
	INTERNALERROR_RPCPILOTSERVERERROR = "InternalError.RpcPilotServerError"

	// 指标查询参数错误。
	INVALIDPARAMETERVALUE_INVALIDRUNTIMEMETRICSEARCHCONDITION = "InvalidParameterValue.InvalidRuntimeMetricSearchCondition"

	// 共享运行时不支持此操作。
	INVALIDPARAMETERVALUE_NOTSUPPORTEDACTIONFORPUBLICRUNTIME = "InvalidParameterValue.NotSupportedActionForPublicRuntime"

	// 不支持的Zone。
	INVALIDPARAMETERVALUE_PILOTZONENOTSUPPORTED = "InvalidParameterValue.PilotZoneNotSupported"

	// 运行时已删除。
	INVALIDPARAMETERVALUE_RUNTIMEALREADYDELETED = "InvalidParameterValue.RuntimeAlreadyDeleted"

	// 运行时不存在。
	INVALIDPARAMETERVALUE_RUNTIMEIDNOTEXIST = "InvalidParameterValue.RuntimeIdNotExist"

	// 当前指标不支持查询百分比。
	INVALIDPARAMETERVALUE_RUNTIMEMETRICRATENOTSUPPORT = "InvalidParameterValue.RuntimeMetricRateNotSupport"

	// 运行时命名空间不合法。
	INVALIDPARAMETERVALUE_RUNTIMENAMESPACEINVALID = "InvalidParameterValue.RuntimeNamespaceInvalid"

	// 运行时地域不存在。
	INVALIDPARAMETERVALUE_RUNTIMEZONENOTEXISTED = "InvalidParameterValue.RuntimeZoneNotExisted"

	// 资源不可用。
	RESOURCEUNAVAILABLE = "ResourceUnavailable"
)
