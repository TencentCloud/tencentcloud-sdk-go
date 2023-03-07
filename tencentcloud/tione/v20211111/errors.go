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

package v20211111

const (
	// 此产品的特有错误码

	// CAM签名/鉴权错误。
	AUTHFAILURE = "AuthFailure"

	// CAM系统异常。
	AUTHFAILURE_CAMEXCEPTION = "AuthFailure.CamException"

	// 没有权限。
	AUTHFAILURE_NOPERMISSION = "AuthFailure.NoPermission"

	// 未授权操作。
	AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"

	// DryRun 操作，代表请求将会是成功的，只是多传了 DryRun 参数。
	DRYRUNOPERATION = "DryRunOperation"

	// 操作失败。
	FAILEDOPERATION = "FailedOperation"

	// API网关访问失败，请重试。
	FAILEDOPERATION_APIGATEWAYQUERYFAILED = "FailedOperation.ApiGatewayQueryFailed"

	// 计费模块访问失败。
	FAILEDOPERATION_BILLINGQUERYFAILED = "FailedOperation.BillingQueryFailed"

	// 绑定标签失败。
	FAILEDOPERATION_BINDINGTAGSFAILED = "FailedOperation.BindingTagsFailed"

	// CAM内部错误。
	FAILEDOPERATION_CAMFAILURE = "FailedOperation.CAMFailure"

	// 调用集群失败。
	FAILEDOPERATION_CALLCLUSTERFAIL = "FailedOperation.CallClusterFail"

	// 尚未开通CLS日志服务，请开前往开通。
	FAILEDOPERATION_CLSSERVICENOTACTIVED = "FailedOperation.ClsServiceNotActived"

	// 集群访问失败。
	FAILEDOPERATION_CLUSTERQUERYFAILED = "FailedOperation.ClusterQueryFailed"

	// 启动实例失败。
	FAILEDOPERATION_CREATEJOBINSTANCEFAILED = "FailedOperation.CreateJobInstanceFailed"

	// cos client 内部错误。
	FAILEDOPERATION_DCCOSCLIENTERR = "FailedOperation.DCCosClientErr"

	// 创建内部异步任务失败。
	FAILEDOPERATION_DCCREATEASYNCTASKERROR = "FailedOperation.DCCreateAsyncTaskError"

	// 创建cos client 失败。
	FAILEDOPERATION_DCCREATEUSERCOSCLIENTERR = "FailedOperation.DCCreateUserCosClientErr"

	// 数据标注 rpc 内部错误。
	FAILEDOPERATION_DCDATAANNOTATIONRPCERR = "FailedOperation.DCDataAnnotationRpcErr"

	// 数据仓库 rpc 内部错误。
	FAILEDOPERATION_DCDATAREPORPCERR = "FailedOperation.DCDatarepoRpcErr"

	// 数据集操作超过上限。
	FAILEDOPERATION_DCDATASETEXCEEDLIMIT = "FailedOperation.DCDatasetExceedLimit"

	// 数据集状态未恢复。
	FAILEDOPERATION_DCDATASETSTATUSNOTREADY = "FailedOperation.DCDatasetStatusNotReady"

	// 获取用户临时秘钥失败。
	FAILEDOPERATION_DCGETUSERTEMPORARYSECRETERR = "FailedOperation.DCGetUserTemporarySecretErr"

	// 数据序列化错误。
	FAILEDOPERATION_DCMARSHALDATAERR = "FailedOperation.DCMarshalDataErr"

	// 数据集获取文件内容异常。
	FAILEDOPERATION_DCQUERYDATASETCONTENTERR = "FailedOperation.DCQueryDatasetContentErr"

	// 数据反序列化错误。
	FAILEDOPERATION_DCUNMARSHALDATAERR = "FailedOperation.DCUnmarshalDataErr"

	// 数据集操作不支持。
	FAILEDOPERATION_DCUNSUPPORTEDOPERATION = "FailedOperation.DCUnsupportedOperation"

	// 名称重复。
	FAILEDOPERATION_DUPLICATENAME = "FailedOperation.DuplicateName"

	// 数据库执行错误。
	FAILEDOPERATION_EXECDATABASEFAIL = "FailedOperation.ExecDatabaseFail"

	// 标签操作失败。
	FAILEDOPERATION_EXECTAGFAIL = "FailedOperation.ExecTagFail"

	// 白名单免费配额不足。
	FAILEDOPERATION_INSUFFICIENTWHITELISTQUOTA = "FailedOperation.InsufficientWhitelistQuota"

	// 操作失败，用户类型异常。
	FAILEDOPERATION_INVALIDUSERTYPE = "FailedOperation.InvalidUserType"

	// 密钥管理系统服务未开通，请先开通腾讯云密钥管理系统服务。
	FAILEDOPERATION_KMSKEYNOTOPEN = "FailedOperation.KmsKeyNotOpen"

	// 移动模型目录失败。
	FAILEDOPERATION_MOVEMODELDIRFAILED = "FailedOperation.MoveModelDirFailed"

	// 没有空闲免费桶。
	FAILEDOPERATION_NOFREEBUCKET = "FailedOperation.NoFreeBucket"

	// 没有权限。
	FAILEDOPERATION_NOPERMISSION = "FailedOperation.NoPermission"

	// 操作不允许。
	FAILEDOPERATION_NOTALLOW = "FailedOperation.NotAllow"

	// 请求正在处理中，请稍候再试。
	FAILEDOPERATION_PROCESSING = "FailedOperation.Processing"

	// 查询资源标签失败。
	FAILEDOPERATION_QUERYBINDINGTAGSFAILED = "FailedOperation.QueryBindingTagsFailed"

	// 数据库查询错误。
	FAILEDOPERATION_QUERYDATABASEFAIL = "FailedOperation.QueryDatabaseFail"

	// 根据标签查询资源失败。
	FAILEDOPERATION_QUERYMODELSBYTAGSFAILED = "FailedOperation.QueryModelsByTagsFailed"

	// 查询计费价格失败。
	FAILEDOPERATION_QUERYPRICEFAILED = "FailedOperation.QueryPriceFailed"

	// 查询计费项失败。
	FAILEDOPERATION_QUERYSPECSFAILED = "FailedOperation.QuerySpecsFailed"

	// 查询标签服务失败。
	FAILEDOPERATION_QUERYTAGFAIL = "FailedOperation.QueryTagFail"

	// 记录不存在。
	FAILEDOPERATION_RECORDNOTFOUND = "FailedOperation.RecordNotFound"

	// 存储库有绑定的实例，请先删除绑定的实例。
	FAILEDOPERATION_REPOBINDBYINSTANCE = "FailedOperation.RepoBindByInstance"

	// 密钥服务访问失败，请重试。
	FAILEDOPERATION_STSQUERYFAILED = "FailedOperation.StsQueryFailed"

	// 实例启动失败。
	FAILEDOPERATION_TIMEDOUT = "FailedOperation.Timedout"

	// 未知的实例规格。
	FAILEDOPERATION_UNKNOWNINSTANCETYPE = "FailedOperation.UnknownInstanceType"

	// 数据解析失败。
	FAILEDOPERATION_UNMARSHALDATA = "FailedOperation.UnmarshalData"

	// 内部错误。
	INTERNALERROR = "InternalError"

	// 冻结失败。
	INTERNALERROR_FREEZEBILLFAILED = "InternalError.FreezeBillFailed"

	// 余额不足。
	INTERNALERROR_INSUFFICIENTBALANCE = "InternalError.InsufficientBalance"

	// 没有权限。
	INTERNALERROR_NOPERMISSION = "InternalError.NoPermission"

	// 操作不允许。
	INTERNALERROR_NOTALLOW = "InternalError.NotAllow"

	// 查询标签失败。
	INTERNALERROR_QUERYBINDINGTAGSFAILED = "InternalError.QueryBindingTagsFailed"

	// 获取HDFS存储信息失败。
	INTERNALERROR_QUERYHDFSINFOFAILED = "InternalError.QueryHDFSInfoFailed"

	// 查询预付费资源组详情失败。
	INTERNALERROR_QUERYRESOURCEGROUPFAILED = "InternalError.QueryResourceGroupFailed"

	// 停止任务失败。
	INTERNALERROR_STOPJOBINSTANCEFAILED = "InternalError.StopJobInstanceFailed"

	// 参数错误。
	INVALIDPARAMETER = "InvalidParameter"

	// 模型包不符合规范。
	INVALIDPARAMETER_MODELFILEINVALID = "InvalidParameter.ModelFileInvalid"

	// 无效的接口。
	INVALIDPARAMETER_TGWINVALIDINTERFACE = "InvalidParameter.TgwInvalidInterface"

	// 无效的请求包体。
	INVALIDPARAMETER_TGWINVALIDREQUESTBODY = "InvalidParameter.TgwInvalidRequestBody"

	// 请求参数校验失败。
	INVALIDPARAMETER_VALIDATEERROR = "InvalidParameter.ValidateError"

	// 参数取值错误。
	INVALIDPARAMETERVALUE = "InvalidParameterValue"

	// 名称不合法。
	INVALIDPARAMETERVALUE_BADNAME = "InvalidParameterValue.BadName"

	// 请设置日志集、日志主题ID。
	INVALIDPARAMETERVALUE_CLSCONFIGREQUIRED = "InvalidParameterValue.ClsConfigRequired"

	// 存储库不存在。
	INVALIDPARAMETERVALUE_CODEREPONOTFOUND = "InvalidParameterValue.CodeRepoNotFound"

	// 不支持的标注类型。
	INVALIDPARAMETERVALUE_DCANNOTATIONTYPE = "InvalidParameterValue.DCAnnotationType"

	// 存储桶参数错误。
	INVALIDPARAMETERVALUE_DCCOSPATHINFO = "InvalidParameterValue.DCCosPathInfo"

	// 数据集标注状态不匹配。
	INVALIDPARAMETERVALUE_DCDATASETANNOTATIONNOTMATCH = "InvalidParameterValue.DCDatasetAnnotationNotMatch"

	// 数据集Id不存在。
	INVALIDPARAMETERVALUE_DCDATASETIDNOTEXIST = "InvalidParameterValue.DCDatasetIdNotExist"

	// 数据集重名已存在。
	INVALIDPARAMETERVALUE_DCDATASETNAMEEXIST = "InvalidParameterValue.DCDatasetNameExist"

	// 不支持的数据集类型。
	INVALIDPARAMETERVALUE_DCDATASETTYPE = "InvalidParameterValue.DCDatasetType"

	// 不支持的过滤参数。
	INVALIDPARAMETERVALUE_DCFILTERVALUES = "InvalidParameterValue.DCFilterValues"

	// 数据集数量超过限制。
	INVALIDPARAMETERVALUE_DATASETNUMLIMITEXCEEDED = "InvalidParameterValue.DatasetNumLimitExceeded"

	// 实例名称冲突，请更换名称后重试。
	INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"

	// 训练框架对应的版本不支持，请阅读文档查看TIONE目前支持的框架和版本。
	INVALIDPARAMETERVALUE_FRAMEWORKVERSIONNOTSUPPORT = "InvalidParameterValue.FrameworkVersionNotSupport"

	// 训练任务镜像不存在。
	INVALIDPARAMETERVALUE_IMAGENOTFOUND = "InvalidParameterValue.ImageNotFound"

	// 无效的过滤器。
	INVALIDPARAMETERVALUE_INVALIDFILTER = "InvalidParameterValue.InvalidFilter"

	// 参数值数量超过限制。
	INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"

	// 操作不允许。
	INVALIDPARAMETERVALUE_NOTALLOW = "InvalidParameterValue.NotAllow"

	// 超过配额限制。
	LIMITEXCEEDED = "LimitExceeded"

	// 缺少参数错误。
	MISSINGPARAMETER = "MissingParameter"

	// 操作被拒绝。
	OPERATIONDENIED = "OperationDenied"

	// 余额不足，创建/更新失败。
	OPERATIONDENIED_BALANCEINSUFFICIENT = "OperationDenied.BalanceInsufficient"

	// 后付费资源售罄。
	OPERATIONDENIED_BILLINGSTATUSRESOURCEINSUFFICIENT = "OperationDenied.BillingStatusResourceInsufficient"

	// 觅影资源包余额不足，请先充值。
	OPERATIONDENIED_MIYINGBALANCEINSUFFICIENT = "OperationDenied.MIYINGBalanceInsufficient"

	// 网段不合法。
	OPERATIONDENIED_NETWORKCIDRILLEGAL = "OperationDenied.NetworkCidrIllegal"

	// 预付费资源组余量不足。
	OPERATIONDENIED_RESOURCEGROUPINSUFFICIENT = "OperationDenied.ResourceGroupInsufficient"

	// 白名单免费配额不足。
	OPERATIONDENIED_WHITELISTQUOTAEXCEED = "OperationDenied.WhitelistQuotaExceed"

	// 资源被占用。
	RESOURCEINUSE = "ResourceInUse"

	// 资源不足。
	RESOURCEINSUFFICIENT = "ResourceInsufficient"

	// 资源不存在。
	RESOURCENOTFOUND = "ResourceNotFound"

	// 没有该模型。
	RESOURCENOTFOUND_NOMODEL = "ResourceNotFound.NoModel"

	// 资源不可用。
	RESOURCEUNAVAILABLE = "ResourceUnavailable"

	// 资源售罄。
	RESOURCESSOLDOUT = "ResourcesSoldOut"

	// 未授权操作。
	UNAUTHORIZEDOPERATION = "UnauthorizedOperation"

	// 资源不属于当前登陆用户主账号，无权限访问。
	UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"

	// 未知参数错误。
	UNKNOWNPARAMETER = "UnknownParameter"

	// 操作不支持。
	UNSUPPORTEDOPERATION = "UnsupportedOperation"
)
