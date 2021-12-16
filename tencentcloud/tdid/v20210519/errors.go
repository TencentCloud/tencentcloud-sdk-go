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

package v20210519

const (
	// 此产品的特有错误码

	// CAM签名/鉴权错误。
	AUTHFAILURE = "AuthFailure"

	// 无访问权限。
	FAILEDOPERATION_INVALIDAUTH = "FailedOperation.InvalidAuth"

	// 操作失败。
	FAILEDOPERATION_OPERATIONEXCEPTION = "FailedOperation.OperationException"

	// 内部错误。
	INTERNALERROR = "InternalError"

	// 服务器异常。
	INTERNALERROR_SERVEREXCEPTION = "InternalError.ServerException"

	// 服务内部错误。
	INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"

	// 服务异常。
	INTERNALERROR_SERVICEPANIC = "InternalError.ServicePanic"

	// 未知错误。
	INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"

	// 参数错误。
	INVALIDPARAMETER = "InvalidParameter"

	// 参数取值错误。
	INVALIDPARAMETERVALUE = "InvalidParameterValue"

	// 参数错误。
	INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"

	// 资源不存在。
	RESOURCENOTFOUND = "ResourceNotFound"

	// 操作不支持。
	UNSUPPORTEDOPERATION = "UnsupportedOperation"
)
