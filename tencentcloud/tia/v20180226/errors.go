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

package v20180226

const (
	// 此产品的特有错误码

	// 资源已存在。
	FAILEDOPERATION_ALREADYEXISTS = "FailedOperation.AlreadyExists"

	// 操作超时。
	FAILEDOPERATION_TIMEOUT = "FailedOperation.TimeOut"

	// 内部错误。
	INTERNALERROR = "InternalError"

	// 内部系统错误。
	INTERNALERROR_SYSTEM = "InternalError.System"

	// 参数错误。
	INVALIDPARAMETER = "InvalidParameter"

	// 参数取值错误。
	INVALIDPARAMETERVALUE = "InvalidParameterValue"

	// Code传入错误。
	INVALIDPARAMETERVALUE_CODE = "InvalidParameterValue.Code"

	// Description传入错误。
	INVALIDPARAMETERVALUE_DESCRIPTION = "InvalidParameterValue.Description"

	// Environment传入错误。
	INVALIDPARAMETERVALUE_ENVIRONMENT = "InvalidParameterValue.Environment"

	// 函数不存在。
	INVALIDPARAMETERVALUE_FUNCTIONNAME = "InvalidParameterValue.FunctionName"

	// Handler传入错误。
	INVALIDPARAMETERVALUE_HANDLER = "InvalidParameterValue.Handler"

	// Order传入错误。
	INVALIDPARAMETERVALUE_ORDER = "InvalidParameterValue.Order"

	// Orderby传入错误。
	INVALIDPARAMETERVALUE_ORDERBY = "InvalidParameterValue.Orderby"

	// Runtime传入错误。
	INVALIDPARAMETERVALUE_RUNTIME = "InvalidParameterValue.Runtime"

	// 函数数量超出最大限制。
	LIMITEXCEEDED_FUNCTION = "LimitExceeded.Function"

	// 内存超出最大限制。
	LIMITEXCEEDED_MEMORY = "LimitExceeded.Memory"

	// Timeout超出最大限制。
	LIMITEXCEEDED_TIMEOUT = "LimitExceeded.Timeout"

	// Code没有传入。
	MISSINGPARAMETER_CODE = "MissingParameter.Code"

	// FunctionName已存在。
	RESOURCEINUSE_FUNCTIONNAME = "ResourceInUse.FunctionName"

	// 资源不存在。
	RESOURCENOTFOUND = "ResourceNotFound"

	// 函数不存在。
	RESOURCENOTFOUND_FUNCTIONNAME = "ResourceNotFound.FunctionName"

	// 资源不可用。
	RESOURCEUNAVAILABLE = "ResourceUnavailable"

	// 未授权操作。
	UNAUTHORIZEDOPERATION = "UnauthorizedOperation"

	// CAM鉴权失败。
	UNAUTHORIZEDOPERATION_CAM = "UnauthorizedOperation.CAM"

	// Region错误。
	UNAUTHORIZEDOPERATION_REGION = "UnauthorizedOperation.Region"

	// 集群版本过低。
	UNSUPPORTEDOPERATION_UNSUPPORTEDVERSION = "UnsupportedOperation.UnsupportedVersion"
)
