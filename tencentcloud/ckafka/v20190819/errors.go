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

package v20190819

const (
	// 此产品的特有错误码

	// CAM签名/鉴权错误。
	AUTHFAILURE = "AuthFailure"

	// DryRun 操作，代表请求将会是成功的，只是多传了 DryRun 参数。
	DRYRUNOPERATION = "DryRunOperation"

	// 操作失败。
	FAILEDOPERATION = "FailedOperation"

	// 内部错误。
	INTERNALERROR = "InternalError"

	// 参数错误。
	INVALIDPARAMETER = "InvalidParameter"

	// 参数取值错误。
	INVALIDPARAMETERVALUE = "InvalidParameterValue"

	// 实例不存在。
	INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"

	// 参数不允许为空。
	INVALIDPARAMETERVALUE_NOTALLOWEDEMPTY = "InvalidParameterValue.NotAllowedEmpty"

	// 已存在相同参数。
	INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"

	// 无效的子网id。
	INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"

	// 子网不属于zone。
	INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"

	// 无效的 Vpc Id。
	INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"

	// Action参数取值错误。
	INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"

	// zone不支持。
	INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"

	// 超过配额限制。
	LIMITEXCEEDED = "LimitExceeded"

	// 路由数超过限制。
	LIMITEXCEEDED_ROUTEOVERLIMIT = "LimitExceeded.RouteOverLimit"

	// SASL路由超过限制。
	LIMITEXCEEDED_ROUTESASLOVERLIMIT = "LimitExceeded.RouteSASLOverLimit"

	// 缺少参数错误。
	MISSINGPARAMETER = "MissingParameter"

	// 操作被拒绝。
	OPERATIONDENIED = "OperationDenied"

	// 任务资源暂停。
	OPERATIONDENIED_RESOURCETASKPAUSED = "OperationDenied.ResourceTaskPaused"

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

	// 未授权操作。
	UNAUTHORIZEDOPERATION = "UnauthorizedOperation"

	// 未知参数错误。
	UNKNOWNPARAMETER = "UnknownParameter"

	// 操作不支持。
	UNSUPPORTEDOPERATION = "UnsupportedOperation"

	// 批量删除实例限制。
	UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"

	// Oss拒绝该操作。
	UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
)
