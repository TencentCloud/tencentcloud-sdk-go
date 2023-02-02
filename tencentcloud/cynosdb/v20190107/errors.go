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

package v20190107

const (
	// 此产品的特有错误码

	// CAM签名/鉴权错误。
	AUTHFAILURE = "AuthFailure"

	// 批量查询失败。
	FAILEDOPERATION_BATCHGETINSTANCEERROR = "FailedOperation.BatchGetInstanceError"

	// 获取权限失败，请稍后重试。如果持续不成功，请联系客服进行处理。
	FAILEDOPERATION_CAMCHECKRESOURCEERROR = "FailedOperation.CamCheckResourceError"

	// 鉴权失败，请稍后重试。如果持续不成功，请联系客服进行处理。
	FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"

	// 创建并支付订单失败。
	FAILEDOPERATION_CREATEORDER = "FailedOperation.CreateOrder"

	// 数据库访问失败，请稍后重试。如果持续不成功，请联系客服进行处理。
	FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"

	// 创建流程失败，请稍后重试。如果持续不成功，请联系客服进行处理。
	FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"

	// Flow不存在。
	FAILEDOPERATION_FLOWNOTFOUNDERROR = "FailedOperation.FlowNotFoundError"

	// 获取备份策略失败，请稍后重试。如果持续不成功，请联系客服进行处理。
	FAILEDOPERATION_GETBACKUPSTRATEGYERROR = "FailedOperation.GetBackupStrategyError"

	// 账号余额不足。
	FAILEDOPERATION_INSUFFICIENTBALANCE = "FailedOperation.InsufficientBalance"

	// 操作失败，请稍后重试。如果持续不成功，请联系客服进行处理。
	FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"

	// 规格信息查询失败。
	FAILEDOPERATION_QUERYSPECBYSPECCODEERROR = "FailedOperation.QuerySpecBySpecCodeError"

	// 创建并支付订单失败，请稍后重试。如果持续不成功，请联系客服进行处理。
	FAILEDOPERATION_TRADECREATEORDERERROR = "FailedOperation.TradeCreateOrderError"

	// 内部错误。
	INTERNALERROR = "InternalError"

	// 查询数据库失败。
	INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"

	// 获取安全组信息失败。
	INTERNALERROR_GETSECURITYGROUPDETAILFAILED = "InternalError.GetSecurityGroupDetailFailed"

	// 获取子网失败。
	INTERNALERROR_GETSUBNETFAILED = "InternalError.GetSubnetFailed"

	// 获取VPC失败。
	INTERNALERROR_GETVPCFAILED = "InternalError.GetVpcFailed"

	// http请求执行异常。
	INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"

	// 安全组查询实例失败。
	INTERNALERROR_LISTINSTANCEFAILED = "InternalError.ListInstanceFailed"

	// 操作外网失败。
	INTERNALERROR_OPERATEWANFAIL = "InternalError.OperateWanFail"

	// 操作不支持。
	INTERNALERROR_OPERATIONNOTSUPPORT = "InternalError.OperationNotSupport"

	// 查询数据库失败。
	INTERNALERROR_QUERYDATABASEFAILED = "InternalError.QueryDatabaseFailed"

	// 系统内部错误。
	INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"

	// 未知的内部错误。
	INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"

	// 参数错误。
	INVALIDPARAMETER = "InvalidParameter"

	// 未找到该接口。
	INVALIDPARAMETER_CONTROLLERNOTFOUNDERROR = "InvalidParameter.ControllerNotFoundError"

	// 参数传参错误。
	INVALIDPARAMETER_EXCEPTIONPARAM = "InvalidParameter.ExceptionParam"

	// 未知参数，请检查参数是否有效。
	INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"

	// 当前实例不可隔离。
	INVALIDPARAMETER_ISOLATENOTALLOWED = "InvalidParameter.IsolateNotAllowed"

	// 参数取值错误。
	INVALIDPARAMETERVALUE = "InvalidParameterValue"

	// 账号已存在。
	INVALIDPARAMETERVALUE_ACCOUNTALREADYEXISTERROR = "InvalidParameterValue.AccountAlreadyExistError"

	// 账号已存在。
	INVALIDPARAMETERVALUE_ACCOUNTEXIST = "InvalidParameterValue.AccountExist"

	// 实例不存在账号。
	INVALIDPARAMETERVALUE_ACCOUNTNOTEXISTERROR = "InvalidParameterValue.AccountNotExistError"

	// 未查询到集群。
	INVALIDPARAMETERVALUE_CLUSTERNOTFOUND = "InvalidParameterValue.ClusterNotFound"

	// 不支持的实例类型。
	INVALIDPARAMETERVALUE_DBTYPENOTFOUND = "InvalidParameterValue.DBTypeNotFound"

	// 集群类型必须为serverless。
	INVALIDPARAMETERVALUE_DBMODENOTSERVERLESSERROR = "InvalidParameterValue.DbModeNotServerlessError"

	// 未查询到订单id。
	INVALIDPARAMETERVALUE_DEALNAMENOTFOUND = "InvalidParameterValue.DealNameNotFound"

	// 任务流ID不存在。
	INVALIDPARAMETERVALUE_FLOWNOTFOUND = "InvalidParameterValue.FlowNotFound"

	// 实例名称字符非法。
	INVALIDPARAMETERVALUE_ILLEGALINSTANCENAME = "InvalidParameterValue.IllegalInstanceName"

	// 无效的排序字段。
	INVALIDPARAMETERVALUE_ILLEGALORDERBY = "InvalidParameterValue.IllegalOrderBy"

	// 密码不符合要求。
	INVALIDPARAMETERVALUE_ILLEGALPASSWORD = "InvalidParameterValue.IllegalPassword"

	// 实例不存在。
	INVALIDPARAMETERVALUE_INSTANCENOTFOUND = "InvalidParameterValue.InstanceNotFound"

	// 内置账号不允许操作。
	INVALIDPARAMETERVALUE_INTERNALACCOUNT = "InvalidParameterValue.InternalAccount"

	// 实例版本非法。
	INVALIDPARAMETERVALUE_INVALIDDBVERSION = "InvalidParameterValue.InvalidDBVersion"

	// 参数值无效。
	INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"

	// 密码不合规。
	INVALIDPARAMETERVALUE_INVALIDPASSWORDVALUEERROR = "InvalidParameterValue.InvalidPasswordValueError"

	// 无效的地域。
	INVALIDPARAMETERVALUE_INVALIDREGIONIDERROR = "InvalidParameterValue.InvalidRegionIdError"

	// 实例规格非法。
	INVALIDPARAMETERVALUE_INVALIDSPEC = "InvalidParameterValue.InvalidSpec"

	// 无效的可用区。
	INVALIDPARAMETERVALUE_INVALIDZONEIDERROR = "InvalidParameterValue.InvalidZoneIdError"

	// 该两个参数不可以同时设定。
	INVALIDPARAMETERVALUE_PARAMBOTHSETERROR = "InvalidParameterValue.ParamBothSetError"

	// 参数错误。
	INVALIDPARAMETERVALUE_PARAMERROR = "InvalidParameterValue.ParamError"

	// 参数值无效，超过允许范围。
	INVALIDPARAMETERVALUE_PARAMETEROUTRANGEERROR = "InvalidParameterValue.ParameterOutRangeError"

	// 预付费类型不允许该操作。
	INVALIDPARAMETERVALUE_PREPAYPAYMODEERROR = "InvalidParameterValue.PrePayPayModeError"

	// 所选地域和可用区不可用。
	INVALIDPARAMETERVALUE_REGIONZONEUNAVAILABLE = "InvalidParameterValue.RegionZoneUnavailable"

	// 未找到相关存储pool。
	INVALIDPARAMETERVALUE_STORAGEPOOLNOTFOUND = "InvalidParameterValue.StoragePoolNotFound"

	// 找不到所选子网。
	INVALIDPARAMETERVALUE_SUBNETNOTFOUND = "InvalidParameterValue.SubnetNotFound"

	// 找不到所选VPC网络。
	INVALIDPARAMETERVALUE_VPCNOTFOUND = "InvalidParameterValue.VpcNotFound"

	// 超过配额限制。
	LIMITEXCEEDED = "LimitExceeded"

	// 集群中节点数超过最大限制。
	LIMITEXCEEDED_CLUSTERINSTANCELIMIT = "LimitExceeded.ClusterInstanceLimit"

	// 用户实例个数超出限制。
	LIMITEXCEEDED_USERINSTANCELIMIT = "LimitExceeded.UserInstanceLimit"

	// 缺少参数错误。
	MISSINGPARAMETER = "MissingParameter"

	// 权限校验失败。
	OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"

	// 由于操作禁止，当前集群不允许该操作。
	OPERATIONDENIED_CLUSTEROPNOTALLOWEDERROR = "OperationDenied.ClusterOpNotAllowedError"

	// 集群当前状态不允许该操作。
	OPERATIONDENIED_CLUSTERSTATUSDENIEDERROR = "OperationDenied.ClusterStatusDeniedError"

	// 实例当前状态不允许该操作。
	OPERATIONDENIED_INSTANCESTATUSDENIEDERROR = "OperationDenied.InstanceStatusDeniedError"

	// 实例未处于运行状态
	OPERATIONDENIED_INSTANCESTATUSLIMITERROR = "OperationDenied.InstanceStatusLimitError"

	// 账号余额不足。
	OPERATIONDENIED_INSUFFICIENTBALANCEERROR = "OperationDenied.InsufficientBalanceError"

	// serverless集群当前状态不允许该操作。
	OPERATIONDENIED_SERVERLESSCLUSTERSTATUSDENIED = "OperationDenied.ServerlessClusterStatusDenied"

	// serverless实例当前状态不允许该操作。
	OPERATIONDENIED_SERVERLESSINSTANCESTATUSDENIED = "OperationDenied.ServerlessInstanceStatusDenied"

	// 任务冲突检查不通过。
	OPERATIONDENIED_TASKCONFLICTERROR = "OperationDenied.TaskConflictError"

	// 用户未进行实名认证，请先进行实名认证才可购买。
	OPERATIONDENIED_USERNOTAUTHENTICATEDERROR = "OperationDenied.UserNotAuthenticatedError"

	// Serverless不支持该版本。
	OPERATIONDENIED_VERSIONNOTSUPPORTERROR = "OperationDenied.VersionNotSupportError"

	// 集群不存在。
	RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"

	// 实例不存在。
	RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"

	// 参数对应资源不存在。
	RESOURCENOTFOUND_RESOURCEERROR = "ResourceNotFound.ResourceError"

	// 锁定实例失败，暂时不可操作。
	RESOURCEUNAVAILABLE_INSTANCELOCKFAIL = "ResourceUnavailable.InstanceLockFail"

	// 实例状态异常，暂时不可操作。
	RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"

	// 非实名用户禁止购买。
	UNAUTHORIZEDOPERATION_NOTREALNAMEACCOUNT = "UnauthorizedOperation.NotRealNameAccount"

	// CAM鉴权不通过。
	UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
)
