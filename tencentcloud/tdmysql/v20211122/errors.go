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

	// 白名单鉴权错误。
	AUTHFAILURE_UINWHITELISTCHECKERROR = "AuthFailure.UinWhiteListCheckError"

	// CAM鉴权失败。
	AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"

	// 绑定标签失败。
	DRYRUNOPERATION_ADDTAGDRYRUNERROR = "DryRunOperation.AddTagDryrunError"

	// FailedOperation.AddInstanceInfoFailed
	FAILEDOPERATION_ADDINSTANCEINFOFAILED = "FailedOperation.AddInstanceInfoFailed"

	// FailedOperation.ApplyVipFailed
	FAILEDOPERATION_APPLYVIPFAILED = "FailedOperation.ApplyVipFailed"

	// FailedOperation.AuthNoStrategy
	FAILEDOPERATION_AUTHNOSTRATEGY = "FailedOperation.AuthNoStrategy"

	// 当前实例版本不支持该接口。
	FAILEDOPERATION_CHECKSUPPORTACTIONERROR = "FailedOperation.CheckSupportActionError"

	// FailedOperation.ClearInstanceInfoFailed
	FAILEDOPERATION_CLEARINSTANCEINFOFAILED = "FailedOperation.ClearInstanceInfoFailed"

	// 创建任务失败。
	FAILEDOPERATION_CREATEFLOWERROR = "FailedOperation.CreateFlowError"

	// 添加实例数据失败。
	FAILEDOPERATION_DBCREATEINSTANCEERROR = "FailedOperation.DBCreateInstanceError"

	// 查询实例数据失败。
	FAILEDOPERATION_DBQUERYINSTANCEERROR = "FailedOperation.DBQueryInstanceError"

	// 查询用户权限错误。
	FAILEDOPERATION_DBQUERYPRIVILEGESERROR = "FailedOperation.DBQueryPrivilegesError"

	// 查询tag绑定实例数据失败。
	FAILEDOPERATION_DBQUERYTAGERROR = "FailedOperation.DBQueryTagError"

	// 查询用户数据失败。
	FAILEDOPERATION_DBQUERYUSERERROR = "FailedOperation.DBQueryUserError"

	// 查询售卖可用区失败。
	FAILEDOPERATION_DBQUERYZONEERROR = "FailedOperation.DBQueryZoneError"

	// 更新实例名称数据失败。
	FAILEDOPERATION_DBUPDATEINSTANCEERROR = "FailedOperation.DBUpdateInstanceError"

	// 更新维护窗口配置失败
	FAILEDOPERATION_DBUPSERTMAINTENANCEWINDOWERROR = "FailedOperation.DBUpsertMaintenanceWindowError"

	// DB数量超出限制。
	FAILEDOPERATION_DBCOUNTLIMITERROR = "FailedOperation.DbCountLimitError"

	// 查询全量备份列表错误。
	FAILEDOPERATION_DESCRIBEFULLBACKUPLISTERROR = "FailedOperation.DescribeFullBackupListError"

	// FailedOperation.DisassociateSecurityGroupsFailed
	FAILEDOPERATION_DISASSOCIATESECURITYGROUPSFAILED = "FailedOperation.DisassociateSecurityGroupsFailed"

	// 实例横向扩容失败。
	FAILEDOPERATION_EXPANDINSTANCEERROR = "FailedOperation.ExpandInstanceError"

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

	// 查询用户失败。
	FAILEDOPERATION_QUERYUSERERROR = "FailedOperation.QueryUserError"

	// FailedOperation.SetRuleLocationFailed
	FAILEDOPERATION_SETRULELOCATIONFAILED = "FailedOperation.SetRuleLocationFailed"

	// table数量超出限制。
	FAILEDOPERATION_TABLECOUNTLIMITERROR = "FailedOperation.TableCountLimitError"

	// FailedOperation.UpdateInstanceInfoFailed
	FAILEDOPERATION_UPDATEINSTANCEINFOFAILED = "FailedOperation.UpdateInstanceInfoFailed"

	// 更新权限错误。
	FAILEDOPERATION_UPDATEPRIVILEGESERROR = "FailedOperation.UpdatePrivilegesError"

	// FailedOperation.WanStatusAbnormal
	FAILEDOPERATION_WANSTATUSABNORMAL = "FailedOperation.WanStatusAbnormal"

	// InternalError.CheckVipStatusFailed
	INTERNALERROR_CHECKVIPSTATUSFAILED = "InternalError.CheckVipStatusFailed"

	// 数据库访问错误。
	INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"

	// InternalError.DbOperationFailed
	INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"

	// 查询数据库对象失败
	INTERNALERROR_DESCRIBEDBOBJECTSERROR = "InternalError.DescribeDBObjectsError"

	// InternalError.GetSecurityGroupDetailFailed
	INTERNALERROR_GETSECURITYGROUPDETAILFAILED = "InternalError.GetSecurityGroupDetailFailed"

	// InternalError.GetSubnetFailed
	INTERNALERROR_GETSUBNETFAILED = "InternalError.GetSubnetFailed"

	// InternalError.GetVpcFailed
	INTERNALERROR_GETVPCFAILED = "InternalError.GetVpcFailed"

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

	// InternalError.VpcOperationFailed
	INTERNALERROR_VPCOPERATIONFAILED = "InternalError.VpcOperationFailed"

	// 参数错误。
	INVALIDPARAMETER = "InvalidParameter"

	// 创建实例参数错误。
	INVALIDPARAMETER_CREATEINSTANCEPARAMERROR = "InvalidParameter.CreateInstanceParamError"

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

	// InvalidParameter.SubnetUnavailable
	INVALIDPARAMETER_SUBNETUNAVAILABLE = "InvalidParameter.SubnetUnavailable"

	// InvalidParameter.VipNotInSubnet
	INVALIDPARAMETER_VIPNOTINSUBNET = "InvalidParameter.VipNotInSubnet"

	// InvalidParameter.VipUsed
	INVALIDPARAMETER_VIPUSED = "InvalidParameter.VipUsed"

	// 参数取值错误。
	INVALIDPARAMETERVALUE = "InvalidParameterValue"

	// 参数值无效或超出范围。
	INVALIDPARAMETERVALUE_BIZINVALIDPARAMETERVALUEERROR = "InvalidParameterValue.BizInvalidParameterValueError"

	// 查询磁盘大小错误。
	INVALIDPARAMETERVALUE_CHECKCLONEINSTANCEDISKERROR = "InvalidParameterValue.CheckCloneInstanceDiskError"

	// 磁盘大小校验错误。
	INVALIDPARAMETERVALUE_CHECKDISKERROR = "InvalidParameterValue.CheckDiskError"

	// host校验失败。
	INVALIDPARAMETERVALUE_CHECKHOSTERROR = "InvalidParameterValue.CheckHostError"

	// 实例版本校验错误。
	INVALIDPARAMETERVALUE_CHECKINSTANCEVERSIONERROR = "InvalidParameterValue.CheckInstanceVersionError"

	// 名称校验失败。
	INVALIDPARAMETERVALUE_CHECKNAMEERROR = "InvalidParameterValue.CheckNameError"

	// 实例规格校验错误。
	INVALIDPARAMETERVALUE_CHECKSPECERROR = "InvalidParameterValue.CheckSpecError"

	// vpc校验错误。
	INVALIDPARAMETERVALUE_CHECKVPCERROR = "InvalidParameterValue.CheckVpcError"

	// 实例过滤参数错误。
	INVALIDPARAMETERVALUE_INSTANCEFILTERKEYERROR = "InvalidParameterValue.InstanceFilterKeyError"

	// 参数错误。
	INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"

	// 维护窗口参数校验失败
	INVALIDPARAMETERVALUE_MAINTENANCEWINDOWPARAMERROR = "InvalidParameterValue.MaintenanceWindowParamError"

	// 创建实例副本数不符合规范。
	INVALIDPARAMETERVALUE_NODEREPLICASINVALIDERROR = "InvalidParameterValue.NodeReplicasInvalidError"

	// Vport范围错误。
	INVALIDPARAMETERVALUE_VPORTRANGEERROR = "InvalidParameterValue.VportRangeError"

	// 绑定tag数超限。
	LIMITEXCEEDED_ADDTAGCOUNTERROR = "LimitExceeded.AddTagCountError"

	// 批量创建实例数超过限制个数。
	LIMITEXCEEDED_OUTOFINSTANCECNTLIMITERROR = "LimitExceeded.OutOfInstanceCntLimitError"

	// 拉取实例超过限制个数。
	LIMITEXCEEDED_OUTOFINSTANCECOUNTLIMITERROR = "LimitExceeded.OutOfInstanceCountLimitError"

	// 规格超过限制。
	LIMITEXCEEDED_OUTOFLIMITERROR = "LimitExceeded.OutOfLimitError"

	// 副本数超过限制。
	LIMITEXCEEDED_OUTOFNODEREPLICASLIMITERROR = "LimitExceeded.OutOfNodeReplicasLimitError"

	// 规格超过限制。
	LIMITEXCEEDED_OUTOFSPECLIMITERROR = "LimitExceeded.OutOfSpecLimitError"

	// 缺少参数错误
	MISSINGPARAMETER_MISSINGPARAMETERERROR = "MissingParameter.MissingParameterError"

	// 操作被拒绝。
	OPERATIONDENIED = "OperationDenied"

	// 暂不支持实例规格缩容操作。
	OPERATIONDENIED_CHECKDECREASESPECERROR = "OperationDenied.CheckDecreaseSpecError"

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

	// 获取最大节点数错误。
	UNSUPPORTEDOPERATION_GETMAXNODENUMERROR = "UnsupportedOperation.GetMaxNodeNumError"
)
