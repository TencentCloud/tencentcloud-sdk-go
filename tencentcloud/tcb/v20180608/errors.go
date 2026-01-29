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

package v20180608

const (
	// 此产品的特有错误码

	// CAM签名/鉴权错误。
	AUTHFAILURE = "AuthFailure"

	// 您没有查看该资源的权限。
	AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"

	// 操作失败。
	FAILEDOPERATION = "FailedOperation"

	// 数据库建立链接失败。
	FAILEDOPERATION_DATABASECONNECTERROR = "FailedOperation.DatabaseConnectError"

	// 执行SQL出错。
	FAILEDOPERATION_DATABASEEXECSQLERROR = "FailedOperation.DatabaseExecSqlError"

	// 数据库元信息异常。
	FAILEDOPERATION_DATABASESCHEMAERROR = "FailedOperation.DatabaseSchemaError"

	// 数据库链接点为空。
	FAILEDOPERATION_EMPTYDATABASEENDPOINT = "FailedOperation.EmptyDatabaseEndpoint"

	// 操作失败，后台依赖平台错误。
	FAILEDOPERATION_PLATFORMERROR = "FailedOperation.PlatformError"

	// Instance is resuming, please try connecting again.
	FAILEDOPERATION_TDSQLPAUSED = "FailedOperation.TdsqlPaused"

	// 内部错误。
	INTERNALERROR = "InternalError"

	// 数据库错误。
	INTERNALERROR_DATABASE = "InternalError.Database"

	// 系统内部异常。
	INTERNALERROR_SYS_ERR = "InternalError.SYS_ERR"

	// 参数错误。
	INVALIDPARAMETER = "InvalidParameter"

	// API已经创建。
	INVALIDPARAMETER_APICREATED = "InvalidParameter.APICreated"

	// 接口名非法。
	INVALIDPARAMETER_ACTION = "InvalidParameter.Action"

	// 环境ID非法。
	INVALIDPARAMETER_ENVID = "InvalidParameter.EnvId"

	// 请求参数错误。
	INVALIDPARAMETER_INVALID_PARAM = "InvalidParameter.INVALID_PARAM"

	// 路径已存在。
	INVALIDPARAMETER_PATHEXIST = "InvalidParameter.PathExist"

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

	// 免费套餐拒绝此操作
	OPERATIONDENIED_FREEPACKAGEDENIED = "OperationDenied.FreePackageDenied"

	// 操作失败：资源被冻结。
	OPERATIONDENIED_RESOURCEFROZEN = "OperationDenied.ResourceFrozen"

	// 资源被占用。
	RESOURCEINUSE = "ResourceInUse"

	// 云存储目前后后台任务正在执行，请稍后再重试。
	RESOURCEINUSE_FSACLJOBUNDONE = "ResourceInUse.FsACLJobUnDone"

	// 资源不足。
	RESOURCEINSUFFICIENT = "ResourceInsufficient"

	// 资源不存在。
	RESOURCENOTFOUND = "ResourceNotFound"

	// 数据库实例不存在。
	RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"

	// 请求的云托管服务未找到
	RESOURCENOTFOUND_SERVERNOTFOUND = "ResourceNotFound.ServerNotFound"

	// 表不存在。
	RESOURCENOTFOUND_TABLENOTFOUND = "ResourceNotFound.TableNotFound"

	// 用户不存在。
	RESOURCENOTFOUND_USERNOTEXISTS = "ResourceNotFound.UserNotExists"

	// 请求参数中的云托管版本未找到
	RESOURCENOTFOUND_VERSIONNOTFOUND = "ResourceNotFound.VersionNotFound"

	// 资源不可用。
	RESOURCEUNAVAILABLE = "ResourceUnavailable"

	// 资源不可用-余额不足。
	RESOURCEUNAVAILABLE_BALANCENOTENOUGH = "ResourceUnavailable.BalanceNotEnough"

	// 资源不可用，CDN冻结。
	RESOURCEUNAVAILABLE_CDNFREEZED = "ResourceUnavailable.CDNFreezed"

	// 当前发票余额不足，无法退费。
	RESOURCEUNAVAILABLE_INVOICEAMOUNTLACK = "ResourceUnavailable.InvoiceAmountLack"

	// 云项目oAuth授权失效（即RefreshToken过期）。
	RESOURCEUNAVAILABLE_REFRESHTOKENEXPIRED = "ResourceUnavailable.RefreshTokenExpired"

	// 资源被封禁
	RESOURCEUNAVAILABLE_RESOURCEBANNED = "ResourceUnavailable.ResourceBanned"

	// 资源已冻结
	RESOURCEUNAVAILABLE_RESOURCEFROZEN = "ResourceUnavailable.ResourceFrozen"

	// 资源已隔离
	RESOURCEUNAVAILABLE_RESOURCEISOLATED = "ResourceUnavailable.ResourceIsolated"

	// 资源过期。
	RESOURCEUNAVAILABLE_RESOURCEOVERDUE = "ResourceUnavailable.ResourceOverdue"

	// 资源售罄。后付费套餐已不再支持购买。
	RESOURCESSOLDOUT_POSTPAYPACKAGENOTAVAILABLE = "ResourcesSoldOut.PostpayPackageNotAvailable"

	// 未授权操作。
	UNAUTHORIZEDOPERATION = "UnauthorizedOperation"

	// 外部代码仓库未授权。
	UNAUTHORIZEDOPERATION_CODEOAUTHUNAUTHORIZED = "UnauthorizedOperation.CodeOAuthUnauthorized"

	// 操作不支持。
	UNSUPPORTEDOPERATION = "UnsupportedOperation"

	// 表数量超过限制。
	UNSUPPORTEDOPERATION_TOOMANYTABLES = "UnsupportedOperation.TooManyTables"
)
