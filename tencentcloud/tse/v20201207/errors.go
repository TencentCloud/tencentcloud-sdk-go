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

package v20201207

const (
	// 此产品的特有错误码

	// 未授权操作错误。
	AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"

	// 操作失败。
	FAILEDOPERATION = "FailedOperation"

	// 创建内部错误。
	INTERNALERROR_CREATEERROR = "InternalError.CreateError"

	// 获取凭证失败。
	INTERNALERROR_GETCREDENTIAL = "InternalError.GetCredential"

	// 角色获取错误。
	INTERNALERROR_GETROLEERROR = "InternalError.GetRoleError"

	// 内部服务调用异常。
	INTERNALERROR_IOERROR = "InternalError.IOError"

	// 服务内部错误。
	INTERNALERROR_INTERNALERROR = "InternalError.InternalError"

	// 操作失败。
	INTERNALERROR_OPERATIONFAILED = "InternalError.OperationFailed"

	// 查询内部错误。
	INTERNALERROR_QUERYERROR = "InternalError.QueryError"

	// TKE相关操作失败。
	INTERNALERROR_TKEFAILURE = "InternalError.TKEFailure"

	// 标签操作失败。
	INTERNALERROR_TAGFAILURE = "InternalError.TagFailure"

	// 未知错误。
	INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"

	// 更新内部错误。
	INTERNALERROR_UPDATEERROR = "InternalError.UpdateError"

	// 访问VPC内部错误。
	INTERNALERROR_VPCFAILURE = "InternalError.VPCFailure"

	// 旧实例不支持此操作。
	INVALIDPARAMETERVALUE_ACTION = "InvalidParameterValue.Action"

	// 请求格式不正确。
	INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"

	// 无效请求参数导致创建失败。
	INVALIDPARAMETERVALUE_CREATEERROR = "InvalidParameterValue.CreateError"

	// 无效的参数值。
	INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"

	// 无效请求参数导致操作失败。
	INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"

	// 无效请求参数，查询失败。
	INVALIDPARAMETERVALUE_QUERYERROR = "InvalidParameterValue.QueryError"

	// 无效请求参数导致更新失败。
	INVALIDPARAMETERVALUE_UPDATEERROR = "InvalidParameterValue.UpdateError"

	// 超过配额限制。
	LIMITEXCEEDED = "LimitExceeded"

	// 缺失参数导致创建失败。
	MISSINGPARAMETER_CREATEERROR = "MissingParameter.CreateError"

	// 缺失参数导致更新失败。
	MISSINGPARAMETER_UPDATEERROR = "MissingParameter.UpdateError"

	// 操作被拒绝。
	OPERATIONDENIED = "OperationDenied"

	// 资源不存在。
	RESOURCENOTFOUND = "ResourceNotFound"

	// 实例不存在。
	RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"

	// 资源不存在。
	RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"

	// 未授权操作。
	UNAUTHORIZEDOPERATION = "UnauthorizedOperation"

	// cam认证失败。
	UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"

	// 子账号缺少passRole权限。
	UNAUTHORIZEDOPERATION_CAMPASSROLENOTEXIST = "UnauthorizedOperation.CamPassRoleNotExist"
)
