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

package v20211122

const (
	// 此产品的特有错误码

	// CAM签名/鉴权错误。
	AUTHFAILURE = "AuthFailure"

	// CAM鉴权错误。
	AUTHFAILURE_CAMAUTHERROR = "AuthFailure.CamAuthError"

	// CAM鉴权请求参数检查失败。
	AUTHFAILURE_CHECKCAMAUTHERROR = "AuthFailure.CheckCamAuthError"

	// CAM鉴权失败。
	AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"

	// FailedOperation.AddInstanceInfoFailed
	FAILEDOPERATION_ADDINSTANCEINFOFAILED = "FailedOperation.AddInstanceInfoFailed"

	// FailedOperation.AuthNoStrategy
	FAILEDOPERATION_AUTHNOSTRATEGY = "FailedOperation.AuthNoStrategy"

	// FailedOperation.ClearInstanceInfoFailed
	FAILEDOPERATION_CLEARINSTANCEINFOFAILED = "FailedOperation.ClearInstanceInfoFailed"

	// 创建任务失败。
	FAILEDOPERATION_CREATEFLOWERROR = "FailedOperation.CreateFlowError"

	// 查询实例数据失败。
	FAILEDOPERATION_DBQUERYINSTANCEERROR = "FailedOperation.DBQueryInstanceError"

	// 更新实例名称数据失败。
	FAILEDOPERATION_DBUPDATEINSTANCEERROR = "FailedOperation.DBUpdateInstanceError"

	// FailedOperation.DisassociateSecurityGroupsFailed
	FAILEDOPERATION_DISASSOCIATESECURITYGROUPSFAILED = "FailedOperation.DisassociateSecurityGroupsFailed"

	// FailedOperation.GetSecurityGroupDetailFailed
	FAILEDOPERATION_GETSECURITYGROUPDETAILFAILED = "FailedOperation.GetSecurityGroupDetailFailed"

	// 实例隔离失败。
	FAILEDOPERATION_ISOLATEINSTANCEERROR = "FailedOperation.IsolateInstanceError"

	// 锁定实例失败。
	FAILEDOPERATION_LOCKINSTANCEERROR = "FailedOperation.LockInstanceError"

	// 修改备份策略失败
	FAILEDOPERATION_MODIFYBACKUPPOLICYERR = "FailedOperation.ModifyBackupPolicyErr"

	// 修改备份策略失败
	FAILEDOPERATION_MODIFYPOLICYERR = "FailedOperation.ModifyPolicyErr"

	// oss查询参数错误。
	FAILEDOPERATION_OSSGETVARIABLESERROR = "FailedOperation.OssGetVariablesError"

	// oss修改参数错误。
	FAILEDOPERATION_OSSMODIFYVARIABLESERROR = "FailedOperation.OssModifyVariablesError"

	// 查询数据库错误。
	FAILEDOPERATION_QUERYDBERROR = "FailedOperation.QueryDBError"

	// FailedOperation.SetRuleLocationFailed
	FAILEDOPERATION_SETRULELOCATIONFAILED = "FailedOperation.SetRuleLocationFailed"

	// FailedOperation.UpdateInstanceInfoFailed
	FAILEDOPERATION_UPDATEINSTANCEINFOFAILED = "FailedOperation.UpdateInstanceInfoFailed"

	// 数据库访问错误。
	INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"

	// 查询数据库对象失败
	INTERNALERROR_DESCRIBEDBOBJECTSERROR = "InternalError.DescribeDBObjectsError"

	// InternalError.GetSecurityGroupDetailFailed
	INTERNALERROR_GETSECURITYGROUPDETAILFAILED = "InternalError.GetSecurityGroupDetailFailed"

	// InternalError.InstanceOperatePermissionError
	INTERNALERROR_INSTANCEOPERATEPERMISSIONERROR = "InternalError.InstanceOperatePermissionError"

	// InternalError.InstanceSGOverLimitError
	INTERNALERROR_INSTANCESGOVERLIMITERROR = "InternalError.InstanceSGOverLimitError"

	// InternalError.ListInstanceRespResourceCountNotMatchError
	INTERNALERROR_LISTINSTANCERESPRESOURCECOUNTNOTMATCHERROR = "InternalError.ListInstanceRespResourceCountNotMatchError"

	// InternalError.ListInstancesError
	INTERNALERROR_LISTINSTANCESERROR = "InternalError.ListInstancesError"

	// InternalError.QueryDatabaseFailed
	INTERNALERROR_QUERYDATABASEFAILED = "InternalError.QueryDatabaseFailed"

	// InternalError.ReadDatabaseFailed
	INTERNALERROR_READDATABASEFAILED = "InternalError.ReadDatabaseFailed"

	// InternalError.RouteNotFound
	INTERNALERROR_ROUTENOTFOUND = "InternalError.RouteNotFound"

	// 找不到路由，是否参数传递不对
	INTERNALERROR_ROUTERNOTFOUND = "InternalError.RouterNotFound"

	// InternalError.SetSvcLocationFailed
	INTERNALERROR_SETSVCLOCATIONFAILED = "InternalError.SetSvcLocationFailed"

	// 参数错误。
	INVALIDPARAMETER = "InvalidParameter"

	// 备份策略参数不合法
	INVALIDPARAMETER_ILLEGALBACKUPPOLICYPARAMSERR = "InvalidParameter.IllegalBackupPolicyParamsErr"

	// InvalidParameter.IllegalParameterError
	INVALIDPARAMETER_ILLEGALPARAMETERERROR = "InvalidParameter.IllegalParameterError"

	// 参数校验错误
	INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"

	// InvalidParameter.InstanceNotFound
	INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"

	// 参数校验失败
	INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"

	// InvalidParameter.PermissionDenied
	INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"

	// 参数取值错误。
	INVALIDPARAMETERVALUE = "InvalidParameterValue"

	// 参数值无效或超出范围。
	INVALIDPARAMETERVALUE_BIZINVALIDPARAMETERVALUEERROR = "InvalidParameterValue.BizInvalidParameterValueError"

	// 实例版本校验错误。
	INVALIDPARAMETERVALUE_CHECKINSTANCEVERSIONERROR = "InvalidParameterValue.CheckInstanceVersionError"

	// 名称校验失败。
	INVALIDPARAMETERVALUE_CHECKNAMEERROR = "InvalidParameterValue.CheckNameError"

	// 实例规格校验错误。
	INVALIDPARAMETERVALUE_CHECKSPECERROR = "InvalidParameterValue.CheckSpecError"

	// 参数错误。
	INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"

	// 缺少参数错误
	MISSINGPARAMETER_MISSINGPARAMETERERROR = "MissingParameter.MissingParameterError"

	// 操作被拒绝。
	OPERATIONDENIED = "OperationDenied"

	// 创建备份任务个数超过阈值
	OPERATIONDENIED_CREATEBACKUPTASKTHRESHOLDERR = "OperationDenied.CreateBackupTaskThresholdErr"

	// 不允许删除正在运行的备份任务
	OPERATIONDENIED_DELETERUNNINGBACKUPTASKERR = "OperationDenied.DeleteRunningBackupTaskErr"

	// 实例状态错误。
	OPERATIONDENIED_INSTANCESTATUSERROR = "OperationDenied.InstanceStatusError"

	// 单个实例每天最多进行手动备份的配额
	OPERATIONDENIED_MANUALBACKUPQUOTAPERDAYEXCEEDEDERR = "OperationDenied.ManualBackupQuotaPerDayExceededErr"

	// 备份数量已经超过配额，不能再创建备份。
	OPERATIONDENIED_MANUALBACKUPSETQUOTAEXCEEDEDERR = "OperationDenied.ManualBackupSetQuotaExceededErr"

	// 实例资源找不到
	RESOURCENOTFOUND_BIZRESOURCENOTFOUNDERROR = "ResourceNotFound.BizResourceNotFoundError"

	// ResourceNotFound.InstanceNotFound
	RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"

	// ResourceNotFound.ProductConfigNotExistedError
	RESOURCENOTFOUND_PRODUCTCONFIGNOTEXISTEDERROR = "ResourceNotFound.ProductConfigNotExistedError"

	// ResourceUnavailable.InstanceStatusAbnormal
	RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"

	// ResourceUnavailable.SGCheckFail
	RESOURCEUNAVAILABLE_SGCHECKFAIL = "ResourceUnavailable.SGCheckFail"

	// UnauthorizedOperation.PermissionDenied
	UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
)
