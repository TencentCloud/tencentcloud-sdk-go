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

package v20230418

const (
	// 此产品的特有错误码

	// CAM签名/鉴权错误。
	AUTHFAILURE = "AuthFailure"

	// 操作失败。
	FAILEDOPERATION = "FailedOperation"

	// FailedOperation.AuthFailed
	FAILEDOPERATION_AUTHFAILED = "FailedOperation.AuthFailed"

	// FailedOperation.ConnectionFailed
	FAILEDOPERATION_CONNECTIONFAILED = "FailedOperation.ConnectionFailed"

	// FailedOperation.DataNotFound
	FAILEDOPERATION_DATANOTFOUND = "FailedOperation.DataNotFound"

	// FailedOperation.DuplicateData
	FAILEDOPERATION_DUPLICATEDATA = "FailedOperation.DuplicateData"

	// 非专业版资源
	FAILEDOPERATION_RESOURCENOTPRO = "FailedOperation.ResourceNotPro"

	// FailedOperation.SearchFailed
	FAILEDOPERATION_SEARCHFAILED = "FailedOperation.SearchFailed"

	// FailedOperation.TooFrequent
	FAILEDOPERATION_TOOFREQUENT = "FailedOperation.TooFrequent"

	// 用户已过期
	FAILEDOPERATION_USEREXPIRED = "FailedOperation.UserExpired"

	// 用户未生效
	FAILEDOPERATION_USERINVALID = "FailedOperation.UserInvalid"

	// 用户已锁定
	FAILEDOPERATION_USERLOCKED = "FailedOperation.UserLocked"

	// 用户不存在
	FAILEDOPERATION_USERNOTEXIST = "FailedOperation.UserNotExist"

	// FailedOperation.VPCDeployed
	FAILEDOPERATION_VPCDEPLOYED = "FailedOperation.VPCDeployed"

	// 内部错误。
	INTERNALERROR = "InternalError"

	// 无效的过滤器
	INVALIDFILTER = "InvalidFilter"

	// 参数错误。
	INVALIDPARAMETER = "InvalidParameter"

	// InvalidParameter.ResourceId
	INVALIDPARAMETER_RESOURCEID = "InvalidParameter.ResourceId"

	// 参数取值错误。
	INVALIDPARAMETERVALUE = "InvalidParameterValue"

	// 超过配额限制。
	LIMITEXCEEDED = "LimitExceeded"

	// 缺少参数错误。
	MISSINGPARAMETER = "MissingParameter"

	// 操作被拒绝。
	OPERATIONDENIED = "OperationDenied"

	// 请求的次数超过了频率限制。
	REQUESTLIMITEXCEEDED = "RequestLimitExceeded"

	// 资源被占用。
	RESOURCEINUSE = "ResourceInUse"

	// 资源不足。
	RESOURCEINSUFFICIENT = "ResourceInsufficient"

	// 资源不存在。
	RESOURCENOTFOUND = "ResourceNotFound"

	// 资源不可用。
	RESOURCEUNAVAILABLE = "ResourceUnavailable"

	// 资源售罄。
	RESOURCESSOLDOUT = "ResourcesSoldOut"

	// 未授权操作。
	UNAUTHORIZEDOPERATION = "UnauthorizedOperation"

	// UnauthorizedOperation.NoPermission
	UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"

	// 未知参数错误。
	UNKNOWNPARAMETER = "UnknownParameter"

	// 操作不支持。
	UNSUPPORTEDOPERATION = "UnsupportedOperation"
)
