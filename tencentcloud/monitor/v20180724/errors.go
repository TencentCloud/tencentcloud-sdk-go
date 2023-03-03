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

	// 访问鉴权失败。
	AUTHFAILURE_ACCESSCAMFAIL = "AuthFailure.AccessCAMFail"

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

	// agent当前状态不允许该操作。
	FAILEDOPERATION_AGENTNOTALLOWED = "FailedOperation.AgentNotAllowed"

	// Agent版本不支持该操作，请升级Agent。
	FAILEDOPERATION_AGENTVERSIONNOTSUPPORTED = "FailedOperation.AgentVersionNotSupported"

	// 该实例上有正在运行的Agent。
	FAILEDOPERATION_AGENTSNOTINUNINSTALLSTAGE = "FailedOperation.AgentsNotInUninstallStage"

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

	// FailedOperation.ComponentClientCommon
	FAILEDOPERATION_COMPONENTCLIENTCOMMON = "FailedOperation.ComponentClientCommon"

	// FailedOperation.ComponentClientHttp
	FAILEDOPERATION_COMPONENTCLIENTHTTP = "FailedOperation.ComponentClientHttp"

	// FailedOperation.ComponentClientUnpack
	FAILEDOPERATION_COMPONENTCLIENTUNPACK = "FailedOperation.ComponentClientUnpack"

	// 创建实例失败。
	FAILEDOPERATION_CREATEINSTANCE = "FailedOperation.CreateInstance"

	// 该账号创建实例被限制。
	FAILEDOPERATION_CREATEINSTANCELIMITED = "FailedOperation.CreateInstanceLimited"

	// 数据表字段不存在。
	FAILEDOPERATION_DATACOLUMNNOTFOUND = "FailedOperation.DataColumnNotFound"

	// 数据查询失败。
	FAILEDOPERATION_DATAQUERYFAILED = "FailedOperation.DataQueryFailed"

	// 数据表不存在。
	FAILEDOPERATION_DATATABLENOTFOUND = "FailedOperation.DataTableNotFound"

	// FailedOperation.Db
	FAILEDOPERATION_DB = "FailedOperation.Db"

	// 数据库查询失败。
	FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"

	// 创建数据库记录失败。
	FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"

	// 数据库记录删除失败。
	FAILEDOPERATION_DBRECORDDELETEFAILED = "FailedOperation.DbRecordDeleteFailed"

	// FailedOperation.DbRecordNotFound
	FAILEDOPERATION_DBRECORDNOTFOUND = "FailedOperation.DbRecordNotFound"

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

	// 后端服务超时
	FAILEDOPERATION_DOHTTPTRANSFERFAILED = "FailedOperation.DoHTTPTransferFailed"

	// 查询分析数据失败。
	FAILEDOPERATION_DRUIDQUERYFAILED = "FailedOperation.DruidQueryFailed"

	// druid表不存在。
	FAILEDOPERATION_DRUIDTABLENOTFOUND = "FailedOperation.DruidTableNotFound"

	// 名字重复。
	FAILEDOPERATION_DUPLICATENAME = "FailedOperation.DuplicateName"

	// 未开通。
	FAILEDOPERATION_ERRNOTOPEN = "FailedOperation.ErrNotOpen"

	// 欠费。
	FAILEDOPERATION_ERROWED = "FailedOperation.ErrOwed"

	// 生成资源ID错误。
	FAILEDOPERATION_GENERATEINSTANCEIDFAILED = "FailedOperation.GenerateInstanceIDFailed"

	// 实例不存在。
	FAILEDOPERATION_INSTANCENOTFOUND = "FailedOperation.InstanceNotFound"

	// 实例没有运行。
	FAILEDOPERATION_INSTANCENOTRUNNING = "FailedOperation.InstanceNotRunning"

	// 内部服务错误。
	FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"

	// FailedOperation.KubeClientConf
	FAILEDOPERATION_KUBECLIENTCONF = "FailedOperation.KubeClientConf"

	// FailedOperation.KubeCommon
	FAILEDOPERATION_KUBECOMMON = "FailedOperation.KubeCommon"

	// 地区不可用。
	FAILEDOPERATION_REGIONUNAVAILABLE = "FailedOperation.RegionUnavailable"

	// 资源已经存在。
	FAILEDOPERATION_RESOURCEEXIST = "FailedOperation.ResourceExist"

	// 资源不存在。
	FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"

	// 资源正在被操作。
	FAILEDOPERATION_RESOURCEOPERATING = "FailedOperation.ResourceOperating"

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

	// 可用区不可用。
	FAILEDOPERATION_ZONEUNAVAILABLE = "FailedOperation.ZoneUnavailable"

	// 内部错误。
	INTERNALERROR = "InternalError"

	// 回调出错。
	INTERNALERROR_CALLBACKFAIL = "InternalError.CallbackFail"

	// InternalError.Db
	INTERNALERROR_DB = "InternalError.Db"

	// InternalError.DbRecordNotFound
	INTERNALERROR_DBRECORDNOTFOUND = "InternalError.DbRecordNotFound"

	// 依赖的其他api出错。
	INTERNALERROR_DEPENDSAPI = "InternalError.DependsApi"

	// 依赖的db出错。
	INTERNALERROR_DEPENDSDB = "InternalError.DependsDb"

	// 依赖的mq出错。
	INTERNALERROR_DEPENDSMQ = "InternalError.DependsMq"

	// 执行超时。
	INTERNALERROR_EXETIMEOUT = "InternalError.ExeTimeout"

	// InternalError.Param
	INTERNALERROR_PARAM = "InternalError.Param"

	// 系统错误。
	INTERNALERROR_SYSTEM = "InternalError.System"

	// InternalError.TaskNotFound
	INTERNALERROR_TASKNOTFOUND = "InternalError.TaskNotFound"

	// 任务结果解析错误。
	INTERNALERROR_TASKRESULTFORMAT = "InternalError.TaskResultFormat"

	// InternalError.UnexpectedInternal
	INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"

	// 参数错误。
	INVALIDPARAMETER = "InvalidParameter"

	// InvalidParameter.ClusterNotFound
	INVALIDPARAMETER_CLUSTERNOTFOUND = "InvalidParameter.ClusterNotFound"

	// 重复提交任务。
	INVALIDPARAMETER_DUPTASK = "InvalidParameter.DupTask"

	// 参数错误。
	INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"

	// 参数错误。
	INVALIDPARAMETER_INVALIDPARAMETERPARAM = "InvalidParameter.InvalidParameterParam"

	// 缺少平台配置。
	INVALIDPARAMETER_MISSAKSK = "InvalidParameter.MissAKSK"

	// InvalidParameter.Param
	INVALIDPARAMETER_PARAM = "InvalidParameter.Param"

	// 参数错误。
	INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"

	// InvalidParameter.PromClusterNotFound
	INVALIDPARAMETER_PROMCLUSTERNOTFOUND = "InvalidParameter.PromClusterNotFound"

	// InvalidParameter.PromInstanceNotFound
	INVALIDPARAMETER_PROMINSTANCENOTFOUND = "InvalidParameter.PromInstanceNotFound"

	// InvalidParameter.ResourceNotFound
	INVALIDPARAMETER_RESOURCENOTFOUND = "InvalidParameter.ResourceNotFound"

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

	// ResourceInUse.ResourceExistAlready
	RESOURCEINUSE_RESOURCEEXISTALREADY = "ResourceInUse.ResourceExistAlready"

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
