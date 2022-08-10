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

package v20210701

const (
	// 此产品的特有错误码

	// 请求响应超时。
	INTERNALERROR_ACTIONREADTIMEOUT = "InternalError.ActionReadTimeout"

	// 添加子网/虚拟节点异常。
	INTERNALERROR_ADDNEWNODEERROR = "InternalError.AddNewNodeError"

	// 创建apm资源失败。
	INTERNALERROR_CREATEAPMRESOURCEERROR = "InternalError.CreateApmResourceError"

	// 创建配置失败。
	INTERNALERROR_CREATECONFIGDATAERROR = "InternalError.CreateConfigDataError"

	// 底层集群创建失败。
	INTERNALERROR_CREATEEKSCLUSTERERROR = "InternalError.CreateEksClusterError"

	// 创建服务失败。
	INTERNALERROR_CREATESERVICEERROR = "InternalError.CreateServiceError"

	// 服务器繁忙,请稍后再试。
	INTERNALERROR_DEFAULTINTERNALERROR = "InternalError.DefaultInternalError"

	// 删除 ingress 失败。
	INTERNALERROR_DELETEINGRESSERROR = "InternalError.DeleteIngressError"

	// 删除应用失败。
	INTERNALERROR_DELETESERVICEERROR = "InternalError.DeleteServiceError"

	// 版本部署调用失败。
	INTERNALERROR_DEPLOYVERSIONERROR = "InternalError.DeployVersionError"

	// 查询配置列表失败。
	INTERNALERROR_DESCRIBECONFIGDATALISTERROR = "InternalError.DescribeConfigDataListError"

	// 查询 ingress 失败。
	INTERNALERROR_DESCRIBEINGRESSERROR = "InternalError.DescribeIngressError"

	// 查询实例信息失败。
	INTERNALERROR_DESCRIBERUNPODLISTERROR = "InternalError.DescribeRunPodListError"

	// 查询service失败。
	INTERNALERROR_DESCRIBESERVICEERROR = "InternalError.DescribeServiceError"

	// 查询服务关联的 ingress 失败。
	INTERNALERROR_DESCRIBESERVICEINGRESSERROR = "InternalError.DescribeServiceIngressError"

	// 查询service列表失败。
	INTERNALERROR_DESCRIBESERVICELISTERROR = "InternalError.DescribeServiceListError"

	// 修改配置失败。
	INTERNALERROR_MODIFYCONFIGDATAERROR = "InternalError.ModifyConfigDataError"

	// 重启失败。
	INTERNALERROR_RESTARTAPPLICATIONERROR = "InternalError.RestartApplicationError"

	// 停止应用失败。
	INTERNALERROR_STOPAPPLICATIONERROR = "InternalError.StopApplicationError"

	// 更新 ingress 失败。
	INTERNALERROR_UPDATEINGRESSERROR = "InternalError.UpdateIngressError"

	// 只支持绑定一种弹性伸缩。
	INVALIDPARAMETERVALUE_AUTOSCALERLARGERTHANONE = "InvalidParameterValue.AutoScalerLargerThanOne"

	// 配置已存在。
	INVALIDPARAMETERVALUE_CONFIGDATAALREADYEXIST = "InvalidParameterValue.ConfigDataAlreadyExist"

	// 配置不合法。
	INVALIDPARAMETERVALUE_CONFIGDATAINVALID = "InvalidParameterValue.ConfigDataInvalid"

	// 定时弹性伸缩目标实例数不合法。
	INVALIDPARAMETERVALUE_CRONHPAREPLICASINVALID = "InvalidParameterValue.CronHpaReplicasInvalid"

	// 弹性伸缩指标不合法。
	INVALIDPARAMETERVALUE_HPAMETRICSINVALID = "InvalidParameterValue.HpaMetricsInvalid"

	// 版本号格式非法。
	INVALIDPARAMETERVALUE_INVALIDDEPLOYVERSION = "InvalidParameterValue.InvalidDeployVersion"

	// 应用名格式非法。
	INVALIDPARAMETERVALUE_INVALIDSERVICENAME = "InvalidParameterValue.InvalidServiceName"

	// 环境重复。
	INVALIDPARAMETERVALUE_NAMESPACEDUPLICATEERROR = "InvalidParameterValue.NamespaceDuplicateError"

	// 命名空间不属于用户。
	INVALIDPARAMETERVALUE_NAMESPACENOTBELONGTOAPPID = "InvalidParameterValue.NamespaceNotBelongToAppid"

	// 环境不存在。
	INVALIDPARAMETERVALUE_NAMESPACENOTFOUND = "InvalidParameterValue.NamespaceNotFound"

	// 环境创建失败，达到上限。
	INVALIDPARAMETERVALUE_NAMESPACEREACHMAXIMUM = "InvalidParameterValue.NamespaceReachMaximum"

	// 环境资源创建失败，达到上限。
	INVALIDPARAMETERVALUE_NAMESPACERESOURCEREACHMAXIMUM = "InvalidParameterValue.NamespaceResourceReachMaximum"

	// 公有镜像参数错误。
	INVALIDPARAMETERVALUE_PUBLICREPOTYPEPARAMETERERROR = "InvalidParameterValue.PublicRepoTypeParameterError"

	// 应用存在正在运行的实例。
	INVALIDPARAMETERVALUE_SERVICEFOUNDRUNNINGVERSION = "InvalidParameterValue.ServiceFoundRunningVersion"

	// 服务名必须小写。
	INVALIDPARAMETERVALUE_SERVICELOWERCASE = "InvalidParameterValue.ServiceLowerCase"

	// 应用名已存在。
	INVALIDPARAMETERVALUE_SERVICENAMEDUPLICATEERROR = "InvalidParameterValue.ServiceNameDuplicateError"

	// 应用不属于此账户。
	INVALIDPARAMETERVALUE_SERVICENOTBELONGTOAPPID = "InvalidParameterValue.ServiceNotBelongToAppid"

	// 实例创建失败，达到上限。
	INVALIDPARAMETERVALUE_SERVICEPODREACHMAXIMUM = "InvalidParameterValue.ServicePodReachMaximum"

	// 应用创建失败，达到上限。
	INVALIDPARAMETERVALUE_SERVICEREACHMAXIMUM = "InvalidParameterValue.ServiceReachMaximum"

	// 不是合法的TEM ID。
	INVALIDPARAMETERVALUE_TEMIDINVALID = "InvalidParameterValue.TemIdInvalid"

	// 非 JAVA 应用不支持链路追踪特性。
	INVALIDPARAMETERVALUE_TRAITSTRACINGNOTSUPPORTED = "InvalidParameterValue.TraitsTracingNotSupported"

	// 版本必须小写。
	INVALIDPARAMETERVALUE_VERSIONLOWERCASE = "InvalidParameterValue.VersionLowerCase"

	// 版本的路由流量不为0。
	INVALIDPARAMETERVALUE_VERSIONROUTERATENOTZERO = "InvalidParameterValue.VersionRouteRateNotZero"

	// 部署版本不能为空。
	MISSINGPARAMETER_DEPLOYVERSIONNULL = "MissingParameter.DeployVersionNull"

	// 环境ID不能为空。
	MISSINGPARAMETER_NAMESPACEIDNULL = "MissingParameter.NamespaceIdNull"

	// 包名不能为空。
	MISSINGPARAMETER_PKGNAMENULL = "MissingParameter.PkgNameNull"

	// 弹性规则ID不能为空。
	MISSINGPARAMETER_SCALERIDNULL = "MissingParameter.ScalerIdNull"

	// 服务ID不能为空。
	MISSINGPARAMETER_SERVICEIDNULL = "MissingParameter.ServiceIdNull"

	// 镜像仓库还未就绪。
	MISSINGPARAMETER_SVCREPONOTREADY = "MissingParameter.SvcRepoNotReady"

	// 账户余额不足。
	OPERATIONDENIED_BALANCENOTENOUGH = "OperationDenied.BalanceNotEnough"

	// 账号欠费状态下不支持该操作，请冲正后重试。
	OPERATIONDENIED_RESOURCEISOLATED = "OperationDenied.ResourceIsolated"

	// 你操作的资源已被其他操作占用，请稍后重试。
	RESOURCEINUSE_RESOURCEALREADYLOCKED = "ResourceInUse.ResourceAlreadyLocked"

	// 资源已绑定。
	RESOURCEINUSE_RESOURCEALREADYUSED = "ResourceInUse.ResourceAlreadyUsed"

	// 目标微服务已离线。
	RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"

	// 环境不存在。
	RESOURCENOTFOUND_NAMESPACENOTFOUND = "ResourceNotFound.NamespaceNotFound"

	// 找不到应用。
	RESOURCENOTFOUND_SERVICENOTFOUND = "ResourceNotFound.ServiceNotFound"

	// 找不到运行的服务实例。
	RESOURCENOTFOUND_SERVICERUNNINGVERSIONNOTFOUND = "ResourceNotFound.ServiceRunningVersionNotFound"

	// 找不到版本对应的环境。
	RESOURCENOTFOUND_VERSIONNAMESPACENOTFOUND = "ResourceNotFound.VersionNamespaceNotFound"

	// 找不到版本对应的应用。
	RESOURCENOTFOUND_VERSIONSERVICENOTFOUND = "ResourceNotFound.VersionServiceNotFound"

	// 应用已停止。
	RESOURCEUNAVAILABLE_APPLICATIONSTOPPED = "ResourceUnavailable.ApplicationStopped"

	// 等待组件安装。
	RESOURCEUNAVAILABLE_WAITFORKRUISE = "ResourceUnavailable.WaitForKruise"

	// 未授权。
	UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"

	// 不支持的ACTION。
	UNSUPPORTEDOPERATION_UNSUPPORTACTION = "UnsupportedOperation.UnsupportAction"
)
