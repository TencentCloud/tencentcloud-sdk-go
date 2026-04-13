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

	// FailedOperation.CreateTable
	FAILEDOPERATION_CREATETABLE = "FailedOperation.CreateTable"

	// 数据不存在
	FAILEDOPERATION_DATASOURCENOTEXIST = "FailedOperation.DataSourceNotExist"

	// 数据库建立链接失败。
	FAILEDOPERATION_DATABASECONNECTERROR = "FailedOperation.DatabaseConnectError"

	// 执行SQL出错。
	FAILEDOPERATION_DATABASEEXECSQLERROR = "FailedOperation.DatabaseExecSqlError"

	// 数据库元信息异常。
	FAILEDOPERATION_DATABASESCHEMAERROR = "FailedOperation.DatabaseSchemaError"

	// FailedOperation.DuplicatedData
	FAILEDOPERATION_DUPLICATEDDATA = "FailedOperation.DuplicatedData"

	// 数据库链接点为空。
	FAILEDOPERATION_EMPTYDATABASEENDPOINT = "FailedOperation.EmptyDatabaseEndpoint"

	// FailedOperation.FlexdbResourceOverdue
	FAILEDOPERATION_FLEXDBRESOURCEOVERDUE = "FailedOperation.FlexdbResourceOverdue"

	// Instance status does not match the required status for this operation.
	FAILEDOPERATION_INSTANCESTATUSCONFLICT = "FailedOperation.InstanceStatusConflict"

	// 无效上下文
	FAILEDOPERATION_INVALIDCONTEXT = "FailedOperation.InvalidContext"

	// FailedOperation.ListTable
	FAILEDOPERATION_LISTTABLE = "FailedOperation.ListTable"

	// 网络异常
	FAILEDOPERATION_NETWORKERROR = "FailedOperation.NetworkError"

	// Failed to connect to PostgreSQL instance.
	FAILEDOPERATION_PGCONNECTERROR = "FailedOperation.PGConnectError"

	// Execute SQL error.
	FAILEDOPERATION_PGEXECUTESQLERROR = "FailedOperation.PGExecuteSqlError"

	// 查询异常
	FAILEDOPERATION_QUERYERROR = "FailedOperation.QueryError"

	// 查询语句解析错误
	FAILEDOPERATION_SYNTAXERROR = "FailedOperation.SyntaxError"

	// Instance is resuming, please try connecting again.
	FAILEDOPERATION_TDSQLPAUSED = "FailedOperation.TdsqlPaused"

	// FailedOperation.Timeout
	FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"

	// Topic隔离
	FAILEDOPERATION_TOPICISOLATED = "FailedOperation.TopicIsolated"

	// 内部错误。
	INTERNALERROR = "InternalError"

	// 数据库错误。
	INTERNALERROR_DATABASE = "InternalError.Database"

	// 系统内部异常。
	INTERNALERROR_SYS_ERR = "InternalError.SYS_ERR"

	// 参数错误。
	INVALIDPARAMETER = "InvalidParameter"

	// 接口名非法。
	INVALIDPARAMETER_ACTION = "InvalidParameter.Action"

	// 证书验证失败
	INVALIDPARAMETER_CERTVERIFYFAILED = "InvalidParameter.CertVerifyFailed"

	// 环境ID非法。
	INVALIDPARAMETER_ENVID = "InvalidParameter.EnvId"

	// HTTP访问服务没有ICP备案
	INVALIDPARAMETER_HTTPSERVICEDOMAINNOTICP = "InvalidParameter.HTTPServiceDomainNotICP"

	// 请求参数错误。
	INVALIDPARAMETER_INVALID_PARAM = "InvalidParameter.INVALID_PARAM"

	// 服务不存在。
	INVALIDPARAMETER_SERVICENOTEXIST = "InvalidParameter.ServiceNotExist"

	// 参数取值错误。
	INVALIDPARAMETERVALUE = "InvalidParameterValue"

	// InvalidParameterValue.InvalidDoc
	INVALIDPARAMETERVALUE_INVALIDDOC = "InvalidParameterValue.InvalidDoc"

	// 超过配额限制。
	LIMITEXCEEDED = "LimitExceeded"

	// 并发请求超过配额限制。
	LIMITEXCEEDED_CONCURRENT = "LimitExceeded.Concurrent"

	// 命名空间超过配额。
	LIMITEXCEEDED_ERRNAMESPACEMAXLIMIT = "LimitExceeded.ErrNamespaceMaxLimit"

	// 镜像容器超过配额。
	LIMITEXCEEDED_ERRREPOMAXLIMIT = "LimitExceeded.ErrRepoMaxLimit"

	// HTTP访问服务域名超过限制
	LIMITEXCEEDED_HTTPSERVICEDOMAIN = "LimitExceeded.HTTPServiceDomain"

	// HTTP访问服务路由超过上限
	LIMITEXCEEDED_HTTPSERVICEROUTE = "LimitExceeded.HTTPServiceRoute"

	// LimitExceeded.NoValidConnection
	LIMITEXCEEDED_NOVALIDCONNECTION = "LimitExceeded.NoValidConnection"

	// LimitExceeded.OutOfIndexQuota
	LIMITEXCEEDED_OUTOFINDEXQUOTA = "LimitExceeded.OutOfIndexQuota"

	// LimitExceeded.OutOfReadRequestQuota
	LIMITEXCEEDED_OUTOFREADREQUESTQUOTA = "LimitExceeded.OutOfReadRequestQuota"

	// 超过响应大小限制
	LIMITEXCEEDED_OUTOFRESULTSIZELIMIT = "LimitExceeded.OutOfResultSizeLimit"

	// LimitExceeded.OutOfTableQuota
	LIMITEXCEEDED_OUTOFTABLEQUOTA = "LimitExceeded.OutOfTableQuota"

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

	// 域名在黑名单中，无法创建
	OPERATIONDENIED_HTTPSERVICEDOMAININBLACKLIST = "OperationDenied.HTTPServiceDomainInBlacklist"

	// 非内部账号禁止操作
	OPERATIONDENIED_NONINTERNALACCOUNT = "OperationDenied.NonInternalAccount"

	// 操作失败：资源被冻结。
	OPERATIONDENIED_RESOURCEFROZEN = "OperationDenied.ResourceFrozen"

	// 资源被占用。
	RESOURCEINUSE = "ResourceInUse"

	// 云存储目前后后台任务正在执行，请稍后再重试。
	RESOURCEINUSE_FSACLJOBUNDONE = "ResourceInUse.FsACLJobUnDone"

	// HTTP访问服务域名已经存在
	RESOURCEINUSE_HTTPSERVICEDOMAIN = "ResourceInUse.HTTPServiceDomain"

	// HTTP访问服务路由已存在
	RESOURCEINUSE_HTTPSERVICEROUTE = "ResourceInUse.HTTPServiceRoute"

	// ResourceInUse.IndexCreating
	RESOURCEINUSE_INDEXCREATING = "ResourceInUse.IndexCreating"

	// 环境下日志资源已存在，无需重新开通日志资源
	RESOURCEINUSE_LOGEXIST = "ResourceInUse.LogExist"

	// 资源不足。
	RESOURCEINSUFFICIENT = "ResourceInsufficient"

	// 资源不存在。
	RESOURCENOTFOUND = "ResourceNotFound"

	// 连接器未找到,请创建连接器或检查连接器参数是否正确
	RESOURCENOTFOUND_CONNECTOR = "ResourceNotFound.Connector"

	// HTTP访问服务域名不存在
	RESOURCENOTFOUND_HTTPSERVICEDOMAIN = "ResourceNotFound.HTTPServiceDomain"

	// 数据库实例不存在。
	RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"

	// Database role not found.
	RESOURCENOTFOUND_ROLENOTFOUND = "ResourceNotFound.RoleNotFound"

	// 请求的云托管服务未找到
	RESOURCENOTFOUND_SERVERNOTFOUND = "ResourceNotFound.ServerNotFound"

	// 表未找到,请创建表或检查表名参数是否正确
	RESOURCENOTFOUND_TABLE = "ResourceNotFound.Table"

	// 表不存在。
	RESOURCENOTFOUND_TABLENOTFOUND = "ResourceNotFound.TableNotFound"

	// 主题不存在
	RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"

	// 用户不存在
	RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"

	// 用户不存在。
	RESOURCENOTFOUND_USERNOTEXISTS = "ResourceNotFound.UserNotExists"

	// 请求参数中的云托管版本未找到
	RESOURCENOTFOUND_VERSIONNOTFOUND = "ResourceNotFound.VersionNotFound"

	// 资源不可用。
	RESOURCEUNAVAILABLE = "ResourceUnavailable"

	// 索引对应的字段已经存在
	RESOURCEUNAVAILABLE_INDEXOPTIONSCONFLICT = "ResourceUnavailable.IndexOptionsConflict"

	// 当前发票余额不足，无法退费。
	RESOURCEUNAVAILABLE_INVOICEAMOUNTLACK = "ResourceUnavailable.InvoiceAmountLack"

	// MongoDB集群已隔离,由于集群已被隔离写入被禁止,请跳转MongoDB控制台查看详情
	RESOURCEUNAVAILABLE_MONGOISOLATED = "ResourceUnavailable.MongoIsolated"

	// 资源被封禁
	RESOURCEUNAVAILABLE_RESOURCEBANNED = "ResourceUnavailable.ResourceBanned"

	// ResourceUnavailable.ResourceExist
	RESOURCEUNAVAILABLE_RESOURCEEXIST = "ResourceUnavailable.ResourceExist"

	// 资源已到期
	RESOURCEUNAVAILABLE_RESOURCEEXPIRED = "ResourceUnavailable.ResourceExpired"

	// 资源已冻结
	RESOURCEUNAVAILABLE_RESOURCEFROZEN = "ResourceUnavailable.ResourceFrozen"

	// 资源已隔离
	RESOURCEUNAVAILABLE_RESOURCEISOLATED = "ResourceUnavailable.ResourceIsolated"

	// 资源过期。
	RESOURCEUNAVAILABLE_RESOURCEOVERDUE = "ResourceUnavailable.ResourceOverdue"

	// 资源售罄。
	RESOURCESSOLDOUT = "ResourcesSoldOut"

	// 未授权操作。
	UNAUTHORIZEDOPERATION = "UnauthorizedOperation"

	// 未知参数错误。
	UNKNOWNPARAMETER = "UnknownParameter"

	// 操作不支持。
	UNSUPPORTEDOPERATION = "UnsupportedOperation"

	// 表数量超过限制。
	UNSUPPORTEDOPERATION_TOOMANYTABLES = "UnsupportedOperation.TooManyTables"
)
