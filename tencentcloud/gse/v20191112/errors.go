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

package v20191112

const (
	// 此产品的特有错误码

	// CAM签名/鉴权错误。
	AUTHFAILURE = "AuthFailure"

	// 资源未授权。
	AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"

	// 操作失败。
	FAILEDOPERATION = "FailedOperation"

	// 内部错误。
	INTERNALERROR = "InternalError"

	// 参数错误。
	INVALIDPARAMETER = "InvalidParameter"

	// 参数取值错误。
	INVALIDPARAMETERVALUE = "InvalidParameterValue"

	// 超过配额限制。
	LIMITEXCEEDED = "LimitExceeded"

	// 超过云联网配额限制。
	LIMITEXCEEDED_CCNLIMITEXCEEDED = "LimitExceeded.CcnLimitExceeded"

	// 服务器舰队个数超限。
	LIMITEXCEEDED_FLEETLIMITEXCEEDED = "LimitExceeded.FleetLimitExceeded"

	// 服务器个数超限。
	LIMITEXCEEDED_INSTANCELIMITEXCEEDED = "LimitExceeded.InstanceLimitExceeded"

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

	// 云联网资源不存在。
	RESOURCENOTFOUND_CCNRESOURCENOTFOUND = "ResourceNotFound.CcnResourceNotFound"

	// 资源不可用。
	RESOURCEUNAVAILABLE = "ResourceUnavailable"

	// 未授权操作。
	UNAUTHORIZEDOPERATION = "UnauthorizedOperation"

	// 没有开通服务。
	UNAUTHORIZEDOPERATION_SERVICENOTOPENED = "UnauthorizedOperation.ServiceNotOpened"

	// 未知参数错误。
	UNKNOWNPARAMETER = "UnknownParameter"

	// 操作不支持。
	UNSUPPORTEDOPERATION = "UnsupportedOperation"

	// 云联网重复关联。
	UNSUPPORTEDOPERATION_CCNATTACHED = "UnsupportedOperation.CcnAttached"

	// 账号ID不存在。
	UNSUPPORTEDOPERATION_UINNOTFOUND = "UnsupportedOperation.UinNotFound"

	// 不支持跨境打通。
	UNSUPPORTEDOPERATION_UNABLECROSSBORDER = "UnsupportedOperation.UnableCrossBorder"
)
