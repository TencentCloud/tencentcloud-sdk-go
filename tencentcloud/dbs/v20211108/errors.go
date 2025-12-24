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

package v20211108

const (
	// 此产品的特有错误码

	// 权限不足
	AUTHFAILURE_UNAUTHORIZEDOPERATIONERROR = "AuthFailure.UnauthorizedOperationError"

	// 创建备份失败
	FAILEDOPERATION_CREATEBACKUPERR = "FailedOperation.CreateBackupErr"

	// 内部错误。
	INTERNALERROR = "InternalError"

	// 数据库访问错误。
	INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"

	// 内部请求失败错误
	INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"

	// 参数错误。
	INVALIDPARAMETER = "InvalidParameter"

	// 参数取值错误。
	INVALIDPARAMETERVALUE = "InvalidParameterValue"

	// 操作被拒绝。
	OPERATIONDENIED = "OperationDenied"

	// 不允许的操作
	OPERATIONDENIED_BIZOPERATIONDENIEDERROR = "OperationDenied.BizOperationDeniedError"

	// 资源不存在。
	RESOURCENOTFOUND = "ResourceNotFound"

	// 未授权操作。
	UNAUTHORIZEDOPERATION = "UnauthorizedOperation"

	// 操作不支持。
	UNSUPPORTEDOPERATION = "UnsupportedOperation"
)
