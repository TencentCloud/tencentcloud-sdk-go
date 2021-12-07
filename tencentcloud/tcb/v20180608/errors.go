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

package v20180608

const (
	// 此产品的特有错误码

	// CAM签名/鉴权错误。
	AUTHFAILURE = "AuthFailure"

	// 您没有查看该资源的权限。
	AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"

	// 操作失败。
	FAILEDOPERATION = "FailedOperation"

	// 部分失败（有一部分操作失败）。
	FAILEDOPERATION_PARTIALFAILURE = "FailedOperation.PartialFailure"

	// 内部错误。
	INTERNALERROR = "InternalError"

	// 数据库错误。
	INTERNALERROR_DATABASE = "InternalError.Database"

	// 系统失败。
	INTERNALERROR_SYSTEMFAIL = "InternalError.SystemFail"

	// 服务超时。
	INTERNALERROR_TIMEOUT = "InternalError.Timeout"

	// 参数错误。
	INVALIDPARAMETER = "InvalidParameter"

	// 接口名非法。
	INVALIDPARAMETER_ACTION = "InvalidParameter.Action"

	// 环境ID非法。
	INVALIDPARAMETER_ENVID = "InvalidParameter.EnvId"

	// 对应资源不存在。
	INVALIDPARAMETER_RESOURCENOTEXISTS = "InvalidParameter.ResourceNotExists"

	// 没有操作权限。
	INVALIDPARAMETER_SERVICEEVIL = "InvalidParameter.ServiceEvil"

	// 服务不存在。
	INVALIDPARAMETER_SERVICENOTEXIST = "InvalidParameter.ServiceNotExist"

	// 参数取值错误。
	INVALIDPARAMETERVALUE = "InvalidParameterValue"

	// 超过配额限制。
	LIMITEXCEEDED = "LimitExceeded"

	// 并发请求超过配额限制。
	LIMITEXCEEDED_CONCURRENT = "LimitExceeded.Concurrent"

	// 命名空间超过配额。
	LIMITEXCEEDED_ERRNAMESPACEMAXLIMIT = "LimitExceeded.ErrNamespaceMaxLimit"

	// 镜像容器超过配额。
	LIMITEXCEEDED_ERRREPOMAXLIMIT = "LimitExceeded.ErrRepoMaxLimit"

	// 请求次数超过配额限制。
	LIMITEXCEEDED_REQUEST = "LimitExceeded.Request"

	// 缺少参数错误。
	MISSINGPARAMETER = "MissingParameter"

	// 缺少必要参数。
	MISSINGPARAMETER_PARAM = "MissingParameter.Param"

	// 操作被拒绝。
	OPERATIONDENIED = "OperationDenied"

	// 操作失败：资源被冻结。
	OPERATIONDENIED_RESOURCEFROZEN = "OperationDenied.ResourceFrozen"

	// 资源被占用。
	RESOURCEINUSE = "ResourceInUse"

	// 资源不足。
	RESOURCEINSUFFICIENT = "ResourceInsufficient"

	// 资源不存在。
	RESOURCENOTFOUND = "ResourceNotFound"

	// 用户不存在。
	RESOURCENOTFOUND_USERNOTEXISTS = "ResourceNotFound.UserNotExists"

	// 资源不可用。
	RESOURCEUNAVAILABLE = "ResourceUnavailable"

	// 资源不可用-余额不足。
	RESOURCEUNAVAILABLE_BALANCENOTENOUGH = "ResourceUnavailable.BalanceNotEnough"

	// 资源不可用，CDN冻结。
	RESOURCEUNAVAILABLE_CDNFREEZED = "ResourceUnavailable.CDNFreezed"

	// 当前发票余额不足，无法退费。
	RESOURCEUNAVAILABLE_INVOICEAMOUNTLACK = "ResourceUnavailable.InvoiceAmountLack"

	// 资源过期。
	RESOURCEUNAVAILABLE_RESOURCEOVERDUE = "ResourceUnavailable.ResourceOverdue"

	// 未授权操作。
	UNAUTHORIZEDOPERATION = "UnauthorizedOperation"

	// 外部代码仓库未授权。
	UNAUTHORIZEDOPERATION_CODEOAUTHUNAUTHORIZED = "UnauthorizedOperation.CodeOAuthUnauthorized"

	// 未知参数错误。
	UNKNOWNPARAMETER = "UnknownParameter"

	// 操作不支持。
	UNSUPPORTEDOPERATION = "UnsupportedOperation"

	// 有正在进行中的任务。
	UNSUPPORTEDOPERATION_TASKEXISTED = "UnsupportedOperation.TaskExisted"
)
