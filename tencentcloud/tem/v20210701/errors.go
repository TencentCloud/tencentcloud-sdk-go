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

	// 创建apm资源失败。
	INTERNALERROR_CREATEAPMRESOURCEERROR = "InternalError.CreateApmResourceError"

	// 底层集群创建失败。
	INTERNALERROR_CREATEEKSCLUSTERERROR = "InternalError.CreateEksClusterError"

	// 创建服务失败。
	INTERNALERROR_CREATESERVICEERROR = "InternalError.CreateServiceError"

	// 服务器繁忙,请稍后再试。
	INTERNALERROR_DEFAULTINTERNALERROR = "InternalError.DefaultInternalError"

	// 版本部署调用失败。
	INTERNALERROR_DEPLOYVERSIONERROR = "InternalError.DeployVersionError"

	// 查询实例信息失败。
	INTERNALERROR_DESCRIBERUNPODLISTERROR = "InternalError.DescribeRunPodListError"

	// 重启失败。
	INTERNALERROR_RESTARTAPPLICATIONERROR = "InternalError.RestartApplicationError"

	// 版本号格式非法。
	INVALIDPARAMETERVALUE_INVALIDDEPLOYVERSION = "InvalidParameterValue.InvalidDeployVersion"

	// 环境重复。
	INVALIDPARAMETERVALUE_NAMESPACEDUPLICATEERROR = "InvalidParameterValue.NamespaceDuplicateError"

	// 环境创建失败，达到上限。
	INVALIDPARAMETERVALUE_NAMESPACEREACHMAXIMUM = "InvalidParameterValue.NamespaceReachMaximum"

	// 应用存在正在运行的实例。
	INVALIDPARAMETERVALUE_SERVICEFOUNDRUNNINGVERSION = "InvalidParameterValue.ServiceFoundRunningVersion"

	// 应用名已存在。
	INVALIDPARAMETERVALUE_SERVICENAMEDUPLICATEERROR = "InvalidParameterValue.ServiceNameDuplicateError"

	// 非 JAVA 应用不支持链路追踪特性。
	INVALIDPARAMETERVALUE_TRAITSTRACINGNOTSUPPORTED = "InvalidParameterValue.TraitsTracingNotSupported"

	// 版本必须小写。
	INVALIDPARAMETERVALUE_VERSIONLOWERCASE = "InvalidParameterValue.VersionLowerCase"

	// 版本的路由流量不为0。
	INVALIDPARAMETERVALUE_VERSIONROUTERATENOTZERO = "InvalidParameterValue.VersionRouteRateNotZero"

	// 环境ID不能为空。
	MISSINGPARAMETER_NAMESPACEIDNULL = "MissingParameter.NamespaceIdNull"

	// 目标微服务已离线。
	RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"

	// 找不到应用。
	RESOURCENOTFOUND_SERVICENOTFOUND = "ResourceNotFound.ServiceNotFound"

	// 找不到运行的服务实例。
	RESOURCENOTFOUND_SERVICERUNNINGVERSIONNOTFOUND = "ResourceNotFound.ServiceRunningVersionNotFound"

	// 找不到版本对应的环境。
	RESOURCENOTFOUND_VERSIONNAMESPACENOTFOUND = "ResourceNotFound.VersionNamespaceNotFound"

	// 等待组件安装。
	RESOURCEUNAVAILABLE_WAITFORKRUISE = "ResourceUnavailable.WaitForKruise"

	// 未授权。
	UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
)
