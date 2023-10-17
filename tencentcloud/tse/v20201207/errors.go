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

	// 调用CLS服务失败
	FAILEDOPERATION_CLS = "FailedOperation.Cls"

	// 操作失败。
	FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"

	// 操作失败，内部错误。
	FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"

	// 超过购买实例的最大数量。
	FAILEDOPERATION_RESOURCE = "FailedOperation.Resource"

	// 获取临时密钥失败
	FAILEDOPERATION_ROLE = "FailedOperation.Role"

	// 调用VPC服务失败
	FAILEDOPERATION_VPC = "FailedOperation.Vpc"

	// 内部错误。
	INTERNALERROR = "InternalError"

	// 创建内部错误。
	INTERNALERROR_CREATEERROR = "InternalError.CreateError"

	// 获取凭证失败。
	INTERNALERROR_GETCREDENTIAL = "InternalError.GetCredential"

	// 角色获取错误。
	INTERNALERROR_GETROLEERROR = "InternalError.GetRoleError"

	// 状态码错误。
	INTERNALERROR_HTTPSTATUSCODEERROR = "InternalError.HttpStatusCodeError"

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

	// 参数取值错误。
	INVALIDPARAMETERVALUE = "InvalidParameterValue"

	// 旧实例不支持此操作。
	INVALIDPARAMETERVALUE_ACTION = "InvalidParameterValue.Action"

	// 请求格式不正确。
	INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"

	// 无效请求参数导致创建失败。
	INVALIDPARAMETERVALUE_CREATEERROR = "InvalidParameterValue.CreateError"

	// 无效的描述信息。
	INVALIDPARAMETERVALUE_DESCRIPTION = "InvalidParameterValue.Description"

	// 网关ID无效
	INVALIDPARAMETERVALUE_GATEWAYID = "InvalidParameterValue.GatewayId"

	// 无效的参数值。
	INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"

	// 无效的名称。
	INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"

	// 无效请求参数导致操作失败。
	INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"

	// 无效请求参数，查询失败。
	INVALIDPARAMETERVALUE_QUERYERROR = "InvalidParameterValue.QueryError"

	// 无效的Region。
	INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"

	// 资源已经存在。
	INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"

	// 网关规格参数内容错误
	INVALIDPARAMETERVALUE_SPECIFICATION = "InvalidParameterValue.Specification"

	// 网关类型参数内容错误
	INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"

	// 无效请求参数导致更新失败。
	INVALIDPARAMETERVALUE_UPDATEERROR = "InvalidParameterValue.UpdateError"

	// 超过配额限制。
	LIMITEXCEEDED = "LimitExceeded"

	// 网关证书域名总和超出限制
	LIMITEXCEEDED_LBDOMAINS = "LimitExceeded.LBDomains"

	// 参数超过限制。
	LIMITEXCEEDED_LIMITEXCEEDED = "LimitExceeded.LimitExceeded"

	// 缺少参数错误。
	MISSINGPARAMETER = "MissingParameter"

	// 缺失参数导致创建失败。
	MISSINGPARAMETER_CREATEERROR = "MissingParameter.CreateError"

	// 缺少参数。
	MISSINGPARAMETER_MISSPARAMETER = "MissingParameter.MissParameter"

	// 缺失参数导致更新失败。
	MISSINGPARAMETER_UPDATEERROR = "MissingParameter.UpdateError"

	// 操作被拒绝。
	OPERATIONDENIED = "OperationDenied"

	// 不允许的操作。
	OPERATIONDENIED_OPERATIONDENIED = "OperationDenied.OperationDenied"

	// 资源不存在。
	RESOURCENOTFOUND = "ResourceNotFound"

	// 资源不存在不允许操作。
	RESOURCENOTFOUND_FORBIDDEN = "ResourceNotFound.Forbidden"

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

	// 当前CLS产品未开通
	UNAUTHORIZEDOPERATION_CLSNOTACTIVATED = "UnauthorizedOperation.ClsNotActivated"

	// Uin未授权
	UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"

	// 未授权的操作。
	UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"

	// 操作不支持。
	UNSUPPORTEDOPERATION = "UnsupportedOperation"
)
