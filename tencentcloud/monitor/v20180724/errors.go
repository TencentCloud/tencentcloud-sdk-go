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

package v20180724

const (
	// 此产品的特有错误码

	// CAM签名/鉴权错误。
	AUTHFAILURE = "AuthFailure"

	// 请求未授权。请参考 CAM 文档对鉴权的说明。
	AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"

	// DryRun 操作，代表请求将会是成功的，只是多传了 DryRun 参数。
	DRYRUNOPERATION = "DryRunOperation"

	// 操作失败。
	FAILEDOPERATION = "FailedOperation"

	// 访问STS失败。
	FAILEDOPERATION_ACCESSSTSFAIL = "FailedOperation.AccessSTSFail"

	// 访问用户TKE集群失败。
	FAILEDOPERATION_ACCESSTKEFAIL = "FailedOperation.AccessTKEFail"

	// 访问标签服务失败。
	FAILEDOPERATION_ACCESSTAGFAIL = "FailedOperation.AccessTagFail"

	// Agent版本不支持该操作，请升级Agent。
	FAILEDOPERATION_AGENTVERSIONNOTSUPPORTED = "FailedOperation.AgentVersionNotSupported"

	// 删除过滤条件失败。
	FAILEDOPERATION_ALERTFILTERRULEDELETEFAILED = "FailedOperation.AlertFilterRuleDeleteFailed"

	// 创建告警策略失败。
	FAILEDOPERATION_ALERTPOLICYCREATEFAILED = "FailedOperation.AlertPolicyCreateFailed"

	// 告警策略删除失败。
	FAILEDOPERATION_ALERTPOLICYDELETEFAILED = "FailedOperation.AlertPolicyDeleteFailed"

	// 告警策略查询失败。
	FAILEDOPERATION_ALERTPOLICYDESCRIBEFAILED = "FailedOperation.AlertPolicyDescribeFailed"

	// 告警策略修改失败。
	FAILEDOPERATION_ALERTPOLICYMODIFYFAILED = "FailedOperation.AlertPolicyModifyFailed"

	// 删除触发条件失败。
	FAILEDOPERATION_ALERTTRIGGERRULEDELETEFAILED = "FailedOperation.AlertTriggerRuleDeleteFailed"

	// Yaml 格式不正确。
	FAILEDOPERATION_BADYAMLFORMAT = "FailedOperation.BadYamlFormat"

	// 数据表字段不存在。
	FAILEDOPERATION_DATACOLUMNNOTFOUND = "FailedOperation.DataColumnNotFound"

	// 数据查询失败。
	FAILEDOPERATION_DATAQUERYFAILED = "FailedOperation.DataQueryFailed"

	// 数据表不存在。
	FAILEDOPERATION_DATATABLENOTFOUND = "FailedOperation.DataTableNotFound"

	// 数据库查询失败。
	FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"

	// 创建数据库记录失败。
	FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"

	// 数据库记录删除失败。
	FAILEDOPERATION_DBRECORDDELETEFAILED = "FailedOperation.DbRecordDeleteFailed"

	// 数据库记录更新失败。
	FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"

	// 数据库事务开始失败。
	FAILEDOPERATION_DBTRANSACTIONBEGINFAILED = "FailedOperation.DbTransactionBeginFailed"

	// 数据库事务提交失败。
	FAILEDOPERATION_DBTRANSACTIONCOMMITFAILED = "FailedOperation.DbTransactionCommitFailed"

	// 请求维度查询服务失败。
	FAILEDOPERATION_DIMQUERYREQUESTFAILED = "FailedOperation.DimQueryRequestFailed"

	// 被除数为0。
	FAILEDOPERATION_DIVISIONBYZERO = "FailedOperation.DivisionByZero"

	// 查询分析数据失败。
	FAILEDOPERATION_DRUIDQUERYFAILED = "FailedOperation.DruidQueryFailed"

	// druid表不存在。
	FAILEDOPERATION_DRUIDTABLENOTFOUND = "FailedOperation.DruidTableNotFound"

	// 名字重复。
	FAILEDOPERATION_DUPLICATENAME = "FailedOperation.DuplicateName"

	// 生成资源ID错误。
	FAILEDOPERATION_GENERATEINSTANCEIDFAILED = "FailedOperation.GenerateInstanceIDFailed"

	// 实例没有运行。
	FAILEDOPERATION_INSTANCENOTRUNNING = "FailedOperation.InstanceNotRunning"

	// 资源已经存在。
	FAILEDOPERATION_RESOURCEEXIST = "FailedOperation.ResourceExist"

	// 资源不存在。
	FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"

	// 发送授权请求失败。
	FAILEDOPERATION_SENDREQUEST = "FailedOperation.SendRequest"

	// 服务未启用，开通服务后方可使用。
	FAILEDOPERATION_SERVICENOTENABLED = "FailedOperation.ServiceNotEnabled"

	// 没有访问TKE权限。
	FAILEDOPERATION_TKECLIENTAUTHFAIL = "FailedOperation.TKEClientAuthFail"

	// TKE的endpoint不可访问。
	FAILEDOPERATION_TKEENDPOINTSTATUSERROR = "FailedOperation.TKEEndpointStatusError"

	// 更新TKE资源时出现冲突。
	FAILEDOPERATION_TKERESOURCECONFLICT = "FailedOperation.TKEResourceConflict"

	// 内部错误。
	INTERNALERROR = "InternalError"

	// 回调出错。
	INTERNALERROR_CALLBACKFAIL = "InternalError.CallbackFail"

	// 依赖的其他api出错。
	INTERNALERROR_DEPENDSAPI = "InternalError.DependsApi"

	// 依赖的db出错。
	INTERNALERROR_DEPENDSDB = "InternalError.DependsDb"

	// 依赖的mq出错。
	INTERNALERROR_DEPENDSMQ = "InternalError.DependsMq"

	// 执行超时。
	INTERNALERROR_EXETIMEOUT = "InternalError.ExeTimeout"

	// 系统错误。
	INTERNALERROR_SYSTEM = "InternalError.System"

	// 任务结果解析错误。
	INTERNALERROR_TASKRESULTFORMAT = "InternalError.TaskResultFormat"

	// 参数错误。
	INVALIDPARAMETER = "InvalidParameter"

	// 重复提交任务。
	INVALIDPARAMETER_DUPTASK = "InvalidParameter.DupTask"

	// 参数错误。
	INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"

	// 参数错误。
	INVALIDPARAMETER_INVALIDPARAMETERPARAM = "InvalidParameter.InvalidParameterParam"

	// 缺少平台配置。
	INVALIDPARAMETER_MISSAKSK = "InvalidParameter.MissAKSK"

	// 参数错误。
	INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"

	// 平台配置错误。
	INVALIDPARAMETER_SECRETIDORSECRETKEYERROR = "InvalidParameter.SecretIdOrSecretKeyError"

	// 这个产品还不支持扫描。
	INVALIDPARAMETER_UNSUPPORTEDPRODUCT = "InvalidParameter.UnsupportedProduct"

	// 参数取值错误。
	INVALIDPARAMETERVALUE = "InvalidParameterValue"

	// dashboard 名重复。
	INVALIDPARAMETERVALUE_DASHBOARDNAMEEXISTS = "InvalidParameterValue.DashboardNameExists"

	// 版本不匹配。
	INVALIDPARAMETERVALUE_VERSIONMISMATCH = "InvalidParameterValue.VersionMismatch"

	// 超过配额限制。
	LIMITEXCEEDED = "LimitExceeded"

	// 指标数量达到配额限制，禁止含有未注册指标的请求。
	LIMITEXCEEDED_METRICQUOTAEXCEEDED = "LimitExceeded.MetricQuotaExceeded"

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

	// 任务不存在。
	RESOURCENOTFOUND_NOTEXISTTASK = "ResourceNotFound.NotExistTask"

	// 资源不可用。
	RESOURCEUNAVAILABLE = "ResourceUnavailable"

	// 资源售罄。
	RESOURCESSOLDOUT = "ResourcesSoldOut"

	// 未授权操作。
	UNAUTHORIZEDOPERATION = "UnauthorizedOperation"

	// 未知参数错误。
	UNKNOWNPARAMETER = "UnknownParameter"

	// 操作不支持。
	UNSUPPORTEDOPERATION = "UnsupportedOperation"
)
