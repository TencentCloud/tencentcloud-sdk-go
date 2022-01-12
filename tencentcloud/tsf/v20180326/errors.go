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

package v20180326

const (
	// 此产品的特有错误码

	// API元数据解析失败。
	FAILEDOPERATION_APIMETAPARSEFAILED = "FailedOperation.ApiMetaParseFailed"

	// 创建应用，获取ES鉴权信息失败。
	FAILEDOPERATION_APPLICATIONCREATEESATUHERROR = "FailedOperation.ApplicationCreateEsAtuhError"

	// 应用查询失败。
	FAILEDOPERATION_APPLICATIONQUERYFAILED = "FailedOperation.ApplicationQueryFailed"

	// 创建集群，开通VPC网络权限失败。
	FAILEDOPERATION_CLUSTERCREATEVPCFAIL = "FailedOperation.ClusterCreateVpcFail"

	// 查询集群失败。
	FAILEDOPERATION_CLUSTERQUERYFAILED = "FailedOperation.ClusterQueryFailed"

	// 应用查询失败。
	FAILEDOPERATION_CONFIGAPPLICATIONQUERYFAILED = "FailedOperation.ConfigApplicationQueryFailed"

	// 配置项创建失败。
	FAILEDOPERATION_CONFIGCREATEFAILED = "FailedOperation.ConfigCreateFailed"

	// 部署组查询失败。
	FAILEDOPERATION_CONFIGGROUPQUERYFAILED = "FailedOperation.ConfigGroupQueryFailed"

	// 命名空间查询失败。
	FAILEDOPERATION_CONFIGNAMESPACEQUERYFAILED = "FailedOperation.ConfigNamespaceQueryFailed"

	// 配置项查询失败。
	FAILEDOPERATION_CONFIGQUERYFAILED = "FailedOperation.ConfigQueryFailed"

	// 配置项发布信息查询失败。
	FAILEDOPERATION_CONFIGRELEASEQUERYFAILED = "FailedOperation.ConfigReleaseQueryFailed"

	// 部署组处于运行状态，无法启动。
	FAILEDOPERATION_CONTAINERGROUPGROUPHASRUN = "FailedOperation.ContainergroupGroupHasrun"

	// 部署组处于停止状态，无法执行此操作。
	FAILEDOPERATION_CONTAINERGROUPGROUPHASSTOP = "FailedOperation.ContainergroupGroupHasstop"

	// 健康检查配置失败。
	FAILEDOPERATION_CVMCAEMASTERHEALTHCHECKCONFIGERROR = "FailedOperation.CvmCaeMasterHealthCheckConfigError"

	// 远端访问错误: %s。
	FAILEDOPERATION_GATEWAYREMOTECALLERROR = "FailedOperation.GatewayRemoteCallError"

	// 命名空间下存在部署组。
	FAILEDOPERATION_GROUPEXISTS = "FailedOperation.GroupExists"

	// 部署组查询失败。
	FAILEDOPERATION_GROUPQUERYFAILD = "FailedOperation.GroupQueryFaild"

	// 查询机器实例部分失败。
	FAILEDOPERATION_INSTANCEQUERYFAILED = "FailedOperation.InstanceQueryFailed"

	// 重装系统失败，请稍后重试。若无法解决，请联系智能客服或提交工单。
	FAILEDOPERATION_INSTANCERESETERROR = "FailedOperation.InstanceResetError"

	// 重装系统，请求超时。
	FAILEDOPERATION_INSTANCERESETTIMEOUT = "FailedOperation.InstanceResetTimeout"

	// 机器实例更新失败。
	FAILEDOPERATION_INSTANCEUPDATEFAILED = "FailedOperation.InstanceUpdateFailed"

	// 泳道从consul删除失败。
	FAILEDOPERATION_LANEINFODELETECONSULFAILED = "FailedOperation.LaneInfoDeleteConsulFailed"

	// 新增关联部署组不能为空。
	FAILEDOPERATION_LANEINFOGROUPNOTEMPTY = "FailedOperation.LaneInfoGroupNotEmpty"

	// 泳道同步到consul失败。
	FAILEDOPERATION_LANEINFORELEASECONSULFAILED = "FailedOperation.LaneInfoReleaseConsulFailed"

	// 泳道发布到mesh失败。
	FAILEDOPERATION_LANEINFORELEASEMESHFAILED = "FailedOperation.LaneInfoReleaseMeshFailed"

	// 全链路灰度规则启用失败。
	FAILEDOPERATION_LANERULEENABLECONSULFAILED = "FailedOperation.LaneRuleEnableConsulFailed"

	// 用户全链路灰度规则最大100条。
	FAILEDOPERATION_LANERULEMAXLIMIT = "FailedOperation.LaneRuleMaxLimit"

	// 无法创建命名空间。
	FAILEDOPERATION_NAMESPACECREATEFAILED = "FailedOperation.NamespaceCreateFailed"

	// 命名空间查询失败。
	FAILEDOPERATION_NAMESPACEQUERYFAILED = "FailedOperation.NamespaceQueryFailed"

	// 访问配置中心失败。
	FAILEDOPERATION_RATELIMITCONSULERROR = "FailedOperation.RatelimitConsulError"

	// 服务数据库入库失败。
	FAILEDOPERATION_SERVICEINSERTFAILED = "FailedOperation.ServiceInsertFailed"

	// 服务查询失败。
	FAILEDOPERATION_SERVICEQUERYFAILED = "FailedOperation.ServiceQueryFailed"

	// 任务创建异常。
	FAILEDOPERATION_TASKCREATEERROR = "FailedOperation.TaskCreateError"

	// 任务删除异常。
	FAILEDOPERATION_TASKDELETEERROR = "FailedOperation.TaskDeleteError"

	// 操作失败。
	FAILEDOPERATION_TASKOPERATIONFAILED = "FailedOperation.TaskOperationFailed"

	// 禁止操作。
	FAILEDOPERATION_TASKOPERATIONFORBIDDEN = "FailedOperation.TaskOperationForbidden"

	// 任务下发异常。
	FAILEDOPERATION_TASKPUSHERROR = "FailedOperation.TaskPushError"

	// 任务查询异常。
	FAILEDOPERATION_TASKQUERYERROR = "FailedOperation.TaskQueryError"

	// 停止任务失败。
	FAILEDOPERATION_TASKTERMINATEFAILED = "FailedOperation.TaskTerminateFailed"

	// 任务更新异常。
	FAILEDOPERATION_TASKUPDATEERROR = "FailedOperation.TaskUpdateError"

	// TKE 集群创建失败，%s。
	FAILEDOPERATION_TKECLUSTERCREATEFAILED = "FailedOperation.TkeClusterCreateFailed"

	// TKE 集群查询失败。
	FAILEDOPERATION_TKECLUSTERQUERYFAILED = "FailedOperation.TkeClusterQueryFailed"

	// resource服务错误。
	FAILEDOPERATION_TSFASRESOURCESERVERERROR = "FailedOperation.TsfAsResourceServerError"

	// TSF权限模块异常，请联系系统管理员。。
	FAILEDOPERATION_TSFPRIVILEGEERROR = "FailedOperation.TsfPrivilegeError"

	// 应用操作请求MASTER FEIGN失败。
	INTERNALERROR_APPLICATIONMASTERFEIGNERROR = "InternalError.ApplicationMasterFeignError"

	// 应用操作请求MASTER 操作失败。
	INTERNALERROR_APPLICATIONMASTERNUKNOWNERROR = "InternalError.ApplicationMasterNuknownError"

	// 删除应用程序包请求仓库失败。
	INTERNALERROR_APPLICATIONREPODELETEPKG = "InternalError.ApplicationRepoDeletePkg"

	// 创建应用初始化tsf-scalable请求失败。
	INTERNALERROR_APPLICATIONSCALABLEINITERROR = "InternalError.ApplicationScalableInitError"

	// TSF云API调用申请角色临时凭证调用请求失败。
	INTERNALERROR_CAMROLEREQUESTERROR = "InternalError.CamRoleRequestError"

	// TSF云API请求调用失败。
	INTERNALERROR_CLOUDAPIPROXYERROR = "InternalError.CloudApiProxyError"

	// 集群通用错误。
	INTERNALERROR_CLUSTERCOMMONERROR = "InternalError.ClusterCommonError"

	// 虚拟机集群请求MASTER FEIGN失败。
	INTERNALERROR_CLUSTERMASTERFEIGNERROR = "InternalError.ClusterMasterFeignError"

	// 无法找到部署组，或相应集群/命名空间/应用的权限不足。
	INTERNALERROR_CLUSTERNOTEXISTORPRIVILEGEERROR = "InternalError.ClusterNotExistOrPrivilegeError"

	// 配置发布失败：配置中心服务器处理失败。
	INTERNALERROR_CONSULSERVERERROR = "InternalError.ConsulServerError"

	// 访问TKE服务失败。
	INTERNALERROR_CONTAINERGROUPKUBERNETEAPIINVOKEERROR = "InternalError.ContainergroupKuberneteApiInvokeError"

	// 连接TKE服务失败。
	INTERNALERROR_CONTAINERGROUPKUBERNETECONNECTERROR = "InternalError.ContainergroupKuberneteConnectError"

	// 容器应用SQL错误。
	INTERNALERROR_CONTAINERGROUPSQLFAILED = "InternalError.ContainergroupSqlFailed"

	// 容器平台集群不可用，当前状态 %s。
	INTERNALERROR_CPCLUSTERUNAVAILABLE = "InternalError.CpClusterUnavailable"

	// 命令下放失败。
	INTERNALERROR_CVMCAEMASTERDISPATCHERROR = "InternalError.CvmCaeMasterDispatchError"

	// TSF MASTER 内部执行错误。
	INTERNALERROR_CVMCAEMASTERINTERNALERROR = "InternalError.CvmCaeMasterInternalError"

	// MASTER通道查询失败。
	INTERNALERROR_CVMCAEMASTERNONALIVE = "InternalError.CvmCaeMasterNonAlive"

	// 网关通用异常:%s。
	INTERNALERROR_GATEWAYCOMMONERROR = "InternalError.GatewayCommonError"

	// 数据一致性异常:%s。
	INTERNALERROR_GATEWAYCONSISTENCYERROR = "InternalError.GatewayConsistencyError"

	// 配置中心访问异常。
	INTERNALERROR_GATEWAYCONSULERROR = "InternalError.GatewayConsulError"

	// 网关数据异常。
	INTERNALERROR_GATEWAYDBERROR = "InternalError.GatewayDbError"

	// 部署组通用异常。
	INTERNALERROR_GROUPCOMMONERROR = "InternalError.GroupCommonError"

	// 部署组操作请求MASTER 操作失败。
	INTERNALERROR_GROUPMASTERNUKNOWNERROR = "InternalError.GroupMasterNuknownError"

	// tcr仓库绑定失败。
	INTERNALERROR_IMAGEREPOTCRBINDERROR = "InternalError.ImagerepoTcrBindError"

	// TSF节点管理通用错误信息。
	INTERNALERROR_INSTANCECOMMONERROR = "InternalError.InstanceCommonError"

	// 创建kubernetes命名空间失败。
	INTERNALERROR_KUBERNETESAPICREATENAMESPACESERROR = "InternalError.KubernetesApiCreateNamespacesError"

	// 创建kubernetes秘钥失败。
	INTERNALERROR_KUBERNETESAPICREATESECRETERROR = "InternalError.KubernetesApiCreateSecretError"

	// kubernetes api 调用失败。
	INTERNALERROR_KUBERNETESCALLERROR = "InternalError.KubernetesCallError"

	// 远程调用失败。
	INTERNALERROR_REMOTESERVICECALLERROR = "InternalError.RemoteServiceCallError"

	// 仓库内部错误。
	INTERNALERROR_RUNTIMEERROR = "InternalError.RuntimeError"

	// IN子句中超过1000个候选项。
	INTERNALERROR_SQLTOOMANYINITEM = "InternalError.SqlTooManyInItem"

	// 任务内部异常。
	INTERNALERROR_TASKINTERNALERROR = "InternalError.TaskInternalError"

	// 调用 TKE 接口失败，%s。
	INTERNALERROR_TKEAPIFAILEDOPERATION = "InternalError.TkeApiFailedOperation"

	// TSF应用性能管理ES客户端响应状态异常。
	INTERNALERROR_TSFAPMESRESPONSESTATUSEXCEPTION = "InternalError.TsfApmEsResponseStatusException"

	// [%s]模块未处理异常。。
	INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"

	// 参数错误。
	INVALIDPARAMETER = "InvalidParameter"

	// [%s]模块接口[%s]请求不正确（400 BAD REQUEST）。。
	INVALIDPARAMETER_BADREQUEST = "InvalidParameter.BadRequest"

	// TSF MASTER 实例状态异常。
	INVALIDPARAMETER_CVMCAEMASTERUNKNOWNINSTANCESTATUS = "InvalidParameter.CvmCaeMasterUnknownInstanceStatus"

	// 参数错误。
	INVALIDPARAMETER_KUBERNETESPARAMERROR = "InvalidParameter.KubernetesParamError"

	// 已经绑定灰度规则，无法删除。
	INVALIDPARAMETER_LANEINFOALREADYUSED = "InvalidParameter.LaneInfoAlreadyUsed"

	// 存在同名的泳道。
	INVALIDPARAMETER_LANEINFONAMEALREADYUSED = "InvalidParameter.LaneInfoNameAlreadyUsed"

	// 泳道名称格式有误。
	INVALIDPARAMETER_LANEINFONAMEINVALID = "InvalidParameter.LaneInfoNameInvalid"

	// 泳道名称不能为空。
	INVALIDPARAMETER_LANEINFONAMENOTEMPTY = "InvalidParameter.LaneInfoNameNotEmpty"

	// 泳道名称不能超过60个字符。
	INVALIDPARAMETER_LANEINFONAMETOOLONG = "InvalidParameter.LaneInfoNameTooLong"

	// 泳道不存在。
	INVALIDPARAMETER_LANEINFONOTEXIST = "InvalidParameter.LaneInfoNotExist"

	// 泳道没有设置任何入口应用。
	INVALIDPARAMETER_LANEINFONOTEXISTENTRANCE = "InvalidParameter.LaneInfoNotExistEntrance"

	// 泳道备注不能超过200个字符。
	INVALIDPARAMETER_LANEINFOREMARKTOOLONG = "InvalidParameter.LaneInfoRemarkTooLong"

	// 泳道规则中的泳道不存在。
	INVALIDPARAMETER_LANERULEINFONOTEXIST = "InvalidParameter.LaneRuleInfoNotExist"

	// 存在同名的泳道规则名称。
	INVALIDPARAMETER_LANERULENAMEALREADYUSED = "InvalidParameter.LaneRuleNameAlreadyUsed"

	// 泳道规则名称格式有误。
	INVALIDPARAMETER_LANERULENAMEINVALID = "InvalidParameter.LaneRuleNameInvalid"

	// 泳道规则名称不能为空。
	INVALIDPARAMETER_LANERULENAMENOTEMPTY = "InvalidParameter.LaneRuleNameNotEmpty"

	// 泳道规则名称不能超过60个字符。
	INVALIDPARAMETER_LANERULENAMETOOLONG = "InvalidParameter.LaneRuleNameTooLong"

	// 泳道规则不存在。
	INVALIDPARAMETER_LANERULENOTEXIST = "InvalidParameter.LaneRuleNotExist"

	// 泳道规则备注不能超过200个字符。
	INVALIDPARAMETER_LANERULEREMARKTOOLONG = "InvalidParameter.LaneRuleRemarkTooLong"

	// 泳道规则标签名不能为空。
	INVALIDPARAMETER_LANERULETAGNAMENOTEMPTY = "InvalidParameter.LaneRuleTagNameNotEmpty"

	// 泳道规则标签名不能超过32个字符。
	INVALIDPARAMETER_LANERULETAGNAMETOOLONG = "InvalidParameter.LaneRuleTagNameTooLong"

	// 泳道规则必须设置至少一个标签。
	INVALIDPARAMETER_LANERULETAGNOTEMPTY = "InvalidParameter.LaneRuleTagNotEmpty"

	// 泳道规则标签值不能超过128个字符。
	INVALIDPARAMETER_LANERULETAGVALUETOOLONG = "InvalidParameter.LaneRuleTagValueTooLong"

	// 泳道规则总标签值不能超过200个字符。
	INVALIDPARAMETER_LANERULETAGVALUETOTALTOOLONG = "InvalidParameter.LaneRuleTagValueTotalTooLong"

	// 包正在被使用，请先解除占用。
	INVALIDPARAMETER_PACKAGEINUSE = "InvalidParameter.PackageInUse"

	// 参数错误。
	INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"

	// 请求参数有误。
	INVALIDPARAMETER_REPOPACKAGEPARAMERROR = "InvalidParameter.RepoPackageParamError"

	// 仓库中存在软件包，请先删除软件包。
	INVALIDPARAMETER_REPOSITORYNOTEMPTY = "InvalidParameter.RepositoryNotEmpty"

	// TSF应用性能管理业务日志配置与应用关联参数错误。
	INVALIDPARAMETER_TSFAPMBUSILOGCFGAPPRELATIONPARAMERROR = "InvalidParameter.TsfApmBusiLogCfgAppRelationParamError"

	// TSF应用性能管理业务日志搜索请求参数错误。
	INVALIDPARAMETER_TSFAPMBUSILOGSEARCHREQUESTPARAMERROR = "InvalidParameter.TsfApmBusiLogSearchRequestParamError"

	// TSF应用性能管理标准输出日志搜索请求参数错误。
	INVALIDPARAMETER_TSFAPMSTDOUTSEARCHREQUESTPARAMERROR = "InvalidParameter.TsfApmStdoutSearchRequestParamError"

	// 仓库批量删除包数量超过单次允许上限。
	INVALIDPARAMETER_UPPERDELETELIMIT = "InvalidParameter.UpperDeleteLimit"

	// 参数取值错误。
	INVALIDPARAMETERVALUE = "InvalidParameterValue"

	// 所属应用ID不能为空。
	INVALIDPARAMETERVALUE_APPLICATIONIDNULL = "InvalidParameterValue.ApplicationIdNull"

	// 无效的微服务类型。
	INVALIDPARAMETERVALUE_APPLICATIONMICROTYPEINVALID = "InvalidParameterValue.ApplicationMicroTypeInvalid"

	// 应用名称已存在，请更换其他名称。
	INVALIDPARAMETERVALUE_APPLICATIONNAMEEXIST = "InvalidParameterValue.ApplicationNameExist"

	// 应用名称不能为空。
	INVALIDPARAMETERVALUE_APPLICATIONNAMENULL = "InvalidParameterValue.ApplicationNameNull"

	// 应用名称格式不正确,只能包含小写字母、数字及分隔符("_"、"-")，且不能以分隔符开头或结尾。
	INVALIDPARAMETERVALUE_APPLICATIONNAMEREGXINVALID = "InvalidParameterValue.ApplicationNameRegxInvalid"

	// 无法获取应用。
	INVALIDPARAMETERVALUE_APPLICATIONNOTEXISTS = "InvalidParameterValue.ApplicationNotExists"

	// 无效的应用分页参数。
	INVALIDPARAMETERVALUE_APPLICATIONPAGELIMITINVALID = "InvalidParameterValue.ApplicationPageLimitInvalid"

	// 无效的应用类型。
	INVALIDPARAMETERVALUE_APPLICATIONTYPEINVALID = "InvalidParameterValue.ApplicationTypeInvalid"

	// 与同VPC其它集群CIDR冲突。
	INVALIDPARAMETERVALUE_CLUSTERCIDRCONFLICT = "InvalidParameterValue.ClusterCidrConflict"

	// 集群命名已存在，请更换其他名称。
	INVALIDPARAMETERVALUE_CLUSTERNAMEEXIST = "InvalidParameterValue.ClusterNameExist"

	// 集群命名不能为空。
	INVALIDPARAMETERVALUE_CLUSTERNAMEREQUIRED = "InvalidParameterValue.ClusterNameRequired"

	// 创建集群，无效的地域字段。
	INVALIDPARAMETERVALUE_CLUSTERREGIONINVALID = "InvalidParameterValue.ClusterRegionInvalid"

	// 非法集群类型。
	INVALIDPARAMETERVALUE_CLUSTERTYPEINVALID = "InvalidParameterValue.ClusterTypeInvalid"

	// 创建集群，无效的可用区字段。
	INVALIDPARAMETERVALUE_CLUSTERZONEINVALID = "InvalidParameterValue.ClusterZoneInvalid"

	// 配置项已经发布过。
	INVALIDPARAMETERVALUE_CONFIGALREADYRELEASED = "InvalidParameterValue.ConfigAlreadyReleased"

	// 配置项已存在。
	INVALIDPARAMETERVALUE_CONFIGEXISTS = "InvalidParameterValue.ConfigExists"

	// 配置项和部署组所属应用不一致。
	INVALIDPARAMETERVALUE_CONFIGGROUPAPPLICATIONIDNOTMATCH = "InvalidParameterValue.ConfigGroupApplicationIdNotMatch"

	// 配置项名称不合规。
	INVALIDPARAMETERVALUE_CONFIGNAMEINVALID = "InvalidParameterValue.ConfigNameInvalid"

	// 无法获取配置项或无权限访问。
	INVALIDPARAMETERVALUE_CONFIGNOTEXISTSORPERMISSIONDENIED = "InvalidParameterValue.ConfigNotExistsOrPermissionDenied"

	// 无法获取配置项发布信息。
	INVALIDPARAMETERVALUE_CONFIGRELEASENOTEXISTS = "InvalidParameterValue.ConfigReleaseNotExists"

	// 配置格式不符合YAML要求。
	INVALIDPARAMETERVALUE_CONFIGVALUEFORMATINVALID = "InvalidParameterValue.ConfigValueFormatInvalid"

	// 配置项值内容大小（%s）超过限制。
	INVALIDPARAMETERVALUE_CONFIGVALUETOOLONG = "InvalidParameterValue.ConfigValueTooLong"

	// 配置项版本描述不合规。
	INVALIDPARAMETERVALUE_CONFIGVERSIONDESCINVALID = "InvalidParameterValue.ConfigVersionDescInvalid"

	// 配置项版本不合规。
	INVALIDPARAMETERVALUE_CONFIGVERSIONINVALID = "InvalidParameterValue.ConfigVersionInvalid"

	// 该镜像被占用中。
	INVALIDPARAMETERVALUE_CONTAINERGROUPIMAGETAGISINUSE = "InvalidParameterValue.ContainerGroupImageTagIsInUse"

	// 服务访问方式不能为空。
	INVALIDPARAMETERVALUE_CONTAINERGROUPACCESSTYPENULL = "InvalidParameterValue.ContainergroupAccesstypeNull"

	// 所属应用ID不能为空。
	INVALIDPARAMETERVALUE_CONTAINERGROUPAPPLICATIONIDNULL = "InvalidParameterValue.ContainergroupApplicationIdNull"

	// 集群 CPU 资源不足。
	INVALIDPARAMETERVALUE_CONTAINERGROUPCPULIMITOVER = "InvalidParameterValue.ContainergroupCpulimitOver"

	// 部署组ID不能为空。
	INVALIDPARAMETERVALUE_CONTAINERGROUPGROUPIDNULL = "InvalidParameterValue.ContainergroupGroupidNull"

	// 部署组名不能大于60个字符。
	INVALIDPARAMETERVALUE_CONTAINERGROUPGROUPNAMELEGNTH = "InvalidParameterValue.ContainergroupGroupnameLegnth"

	// 部署组名称格式不正确,只能包含小写字母、数字及分隔符("-"),且必须以小写字母开头，数字或小写字母结尾。
	INVALIDPARAMETERVALUE_CONTAINERGROUPGROUPNAMEREGEXMATCHFALSE = "InvalidParameterValue.ContainergroupGroupnameRegexMatchFalse"

	// 实例数量不能为空或不合法。
	INVALIDPARAMETERVALUE_CONTAINERGROUPINSTANCENUMINVALID = "InvalidParameterValue.ContainergroupInstanceNumInvalid"

	// CPU limit 和 request 不能同时为空。
	INVALIDPARAMETERVALUE_CONTAINERGROUPINVALIDCPUINFO = "InvalidParameterValue.ContainergroupInvalidCpuInfo"

	// 内存 limit 和 request 不能同时为空。
	INVALIDPARAMETERVALUE_CONTAINERGROUPINVALIDMEMINFO = "InvalidParameterValue.ContainergroupInvalidMemInfo"

	// 集群内存资源不足。
	INVALIDPARAMETERVALUE_CONTAINERGROUPMEMLIMITOVER = "InvalidParameterValue.ContainergroupMemlimitOver"

	// 主机端口值非法。
	INVALIDPARAMETERVALUE_CONTAINERGROUPNODEPORTINVALID = "InvalidParameterValue.ContainergroupNodePortInvalid"

	// 服务端口值非法。
	INVALIDPARAMETERVALUE_CONTAINERGROUPPORTINVALID = "InvalidParameterValue.ContainergroupPortInvalid"

	// 服务端口不允许重复映射。
	INVALIDPARAMETERVALUE_CONTAINERGROUPPORTSREPEAT = "InvalidParameterValue.ContainergroupPortsRepeat"

	// 协议值非法,限定:TCP/UDP。
	INVALIDPARAMETERVALUE_CONTAINERGROUPPROTOCOLINVALID = "InvalidParameterValue.ContainergroupProtocolInvalid"

	// 协议端口不能为空。
	INVALIDPARAMETERVALUE_CONTAINERGROUPPROTOCOLPORTSNULL = "InvalidParameterValue.ContainergroupProtocolPortsNull"

	// 镜像仓库名与应用名不匹配。
	INVALIDPARAMETERVALUE_CONTAINERGROUPREPONAMEINVALID = "InvalidParameterValue.ContainergroupReponameInvalid"

	// agent 容器资源值非法 , %s。
	INVALIDPARAMETERVALUE_CONTAINERGROUPRESOURCEAGENTVALUEINVALID = "InvalidParameterValue.ContainergroupResourceAgentValueInvalid"

	// 容器端口不允许重复映射。
	INVALIDPARAMETERVALUE_CONTAINERGROUPTARGETPORTSREPEAT = "InvalidParameterValue.ContainergroupTargetPortsRepeat"

	// 容器端口不能为空。
	INVALIDPARAMETERVALUE_CONTAINERGROUPTARGETPORTNULL = "InvalidParameterValue.ContainergroupTargetportNull"

	// 更新间隔不能为空或者数值非法。
	INVALIDPARAMETERVALUE_CONTAINERGROUPUPDATEIVLINVALID = "InvalidParameterValue.ContainergroupUpdateivlInvalid"

	// 找不到业务容器。
	INVALIDPARAMETERVALUE_CONTAINERGROUPYAMLUSERCONTAINERNOTFOUND = "InvalidParameterValue.ContainergroupYamlUserContainerNotFound"

	// TSF MASTER 正在执行任务，请等待任务执行完成再下发新任务。
	INVALIDPARAMETERVALUE_CVMCAEMASTERAGENTBUSY = "InvalidParameterValue.CvmCaeMasterAgentBusy"

	// 无可用实例。
	INVALIDPARAMETERVALUE_CVMCAEMASTERAGENTNOTFOUND = "InvalidParameterValue.CvmCaeMasterAgentNotFound"

	// TSF MASTER 部署组中无云主机。
	INVALIDPARAMETERVALUE_CVMCAEMASTERGROUPNOAGENT = "InvalidParameterValue.CvmCaeMasterGroupNoAgent"

	// 部署组不存在。
	INVALIDPARAMETERVALUE_DEPLOYGROUPNOTEXISTS = "InvalidParameterValue.DeployGroupNotExists"

	// 文件配置项已经发布。
	INVALIDPARAMETERVALUE_FILECONFIGALREADYRELEASED = "InvalidParameterValue.FileConfigAlreadyReleased"

	// 文件配置项已存在。
	INVALIDPARAMETERVALUE_FILECONFIGEXISTS = "InvalidParameterValue.FileConfigExists"

	// 配置文件路径重复。
	INVALIDPARAMETERVALUE_FILECONFIGEXISTSPATH = "InvalidParameterValue.FileConfigExistsPath"

	// 其他用户已发布此配置文件路径。
	INVALIDPARAMETERVALUE_FILECONFIGEXISTSPATHOTHER = "InvalidParameterValue.FileConfigExistsPathOther"

	// 文件配置项文件路径不合规。
	INVALIDPARAMETERVALUE_FILECONFIGFILEPATHINVALID = "InvalidParameterValue.FileConfigFilePathInvalid"

	// 文件配置项名称不合规。
	INVALIDPARAMETERVALUE_FILECONFIGNAMEINVALID = "InvalidParameterValue.FileConfigNameInvalid"

	// 同一部署组禁止配置文件重复(文件路径+文件名)。
	INVALIDPARAMETERVALUE_FILECONFIGPATHEXISTS = "InvalidParameterValue.FileConfigPathExists"

	// 文件配置项版本描述不合规。
	INVALIDPARAMETERVALUE_FILECONFIGVERSIONDESCINVALID = "InvalidParameterValue.FileConfigVersionDescInvalid"

	// 请求参数异常:%s。
	INVALIDPARAMETERVALUE_GATEWAYPARAMETERERROR = "InvalidParameterValue.GatewayParameterError"

	// 无效请求参数:%s。
	INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"

	// 全局命名空间已经存在，只能创建一个全局命名空间。
	INVALIDPARAMETERVALUE_GLOBALNAMESPACENAMEEXIST = "InvalidParameterValue.GlobalNamespaceNameExist"

	// 部署相关请求参数校验失败。
	INVALIDPARAMETERVALUE_GROUPBATCHPARAMETERINVALID = "InvalidParameterValue.GroupBatchParameterInvalid"

	// 部署组的集群未绑定该命名空间。
	INVALIDPARAMETERVALUE_GROUPCLUSTERNAMESPACENOTBOUND = "InvalidParameterValue.GroupClusterNamespaceNotBound"

	// 部署组ID不能为空。
	INVALIDPARAMETERVALUE_GROUPIDNULL = "InvalidParameterValue.GroupIdNull"

	// 部署组名称已存在，请更换其他名称。
	INVALIDPARAMETERVALUE_GROUPNAMEEXIST = "InvalidParameterValue.GroupNameExist"

	// 部署组名不能大于50个字符。
	INVALIDPARAMETERVALUE_GROUPNAMELENGTH = "InvalidParameterValue.GroupNameLength"

	// 部署组名不能为空。
	INVALIDPARAMETERVALUE_GROUPNAMENULL = "InvalidParameterValue.GroupNameNull"

	// 部署组名称格式不正确,只能包含小写字母、数字及分隔符("-"),且必须以小写字母开头，数字或小写字母结尾。
	INVALIDPARAMETERVALUE_GROUPNAMEREGXMISMATCH = "InvalidParameterValue.GroupNameRegxMismatch"

	// 无法获取部署组。
	INVALIDPARAMETERVALUE_GROUPNOTEXISTS = "InvalidParameterValue.GroupNotExists"

	// 分组无效的分业参数。
	INVALIDPARAMETERVALUE_GROUPPAGELIMITINVALID = "InvalidParameterValue.GroupPageLimitInvalid"

	// 程序包不能为空。
	INVALIDPARAMETERVALUE_GROUPPKGNULL = "InvalidParameterValue.GroupPkgNull"

	// 无效的部署组状态过滤字段。
	INVALIDPARAMETERVALUE_GROUPSTATUSINVALID = "InvalidParameterValue.GroupStatusInvalid"

	// 分组操作，无有效机器。
	INVALIDPARAMETERVALUE_GROUPVALIDINSTANCENULL = "InvalidParameterValue.GroupValidInstanceNull"

	// 镜像仓库名不合法,示例:tsf-repo/nginx。
	INVALIDPARAMETERVALUE_IMAGEREPOREPONAMEINVALID = "InvalidParameterValue.ImagerepoReponameInvalid"

	// 镜像仓库名不能为空。
	INVALIDPARAMETERVALUE_IMAGEREPOREPONAMENULL = "InvalidParameterValue.ImagerepoReponameNull"

	// imageTags不能为空。
	INVALIDPARAMETERVALUE_IMAGEREPOTAGNAMENULL = "InvalidParameterValue.ImagerepoTagnameNull"

	// 重装系统，无效的镜像id。
	INVALIDPARAMETERVALUE_INSTANCEINVALIDIMAGE = "InvalidParameterValue.InstanceInvalidImage"

	// 参数格式异常。
	INVALIDPARAMETERVALUE_INVALIDPARAMETERFORMAT = "InvalidParameterValue.InvalidParameterFormat"

	// 已经绑定灰度规则，无法删除。
	INVALIDPARAMETERVALUE_LANEINFOALREADYUSED = "InvalidParameterValue.LaneInfoAlreadyUsed"

	// 存在同名的泳道。
	INVALIDPARAMETERVALUE_LANEINFONAMEALREADYUSED = "InvalidParameterValue.LaneInfoNameAlreadyUsed"

	// 泳道名称格式有误。
	INVALIDPARAMETERVALUE_LANEINFONAMEINVALID = "InvalidParameterValue.LaneInfoNameInvalid"

	// 泳道名称不能为空。
	INVALIDPARAMETERVALUE_LANEINFONAMENOTEMPTY = "InvalidParameterValue.LaneInfoNameNotEmpty"

	// 泳道名称不能超过60个字符。
	INVALIDPARAMETERVALUE_LANEINFONAMETOOLONG = "InvalidParameterValue.LaneInfoNameTooLong"

	// 泳道不存在。
	INVALIDPARAMETERVALUE_LANEINFONOTEXIST = "InvalidParameterValue.LaneInfoNotExist"

	// 泳道没有设置任何入口应用。
	INVALIDPARAMETERVALUE_LANEINFONOTEXISTENTRANCE = "InvalidParameterValue.LaneInfoNotExistEntrance"

	// 泳道备注不能超过200个字符。
	INVALIDPARAMETERVALUE_LANEINFOREMARKTOOLONG = "InvalidParameterValue.LaneInfoRemarkTooLong"

	// 全链路灰度规则中的泳道不存在。
	INVALIDPARAMETERVALUE_LANERULEINFONOTEXIST = "InvalidParameterValue.LaneRuleInfoNotExist"

	// 存在同名的全链路灰度规则。
	INVALIDPARAMETERVALUE_LANERULENAMEALREADYUSED = "InvalidParameterValue.LaneRuleNameAlreadyUsed"

	// 全链路灰度规则名称格式有误。
	INVALIDPARAMETERVALUE_LANERULENAMEINVALID = "InvalidParameterValue.LaneRuleNameInvalid"

	// 全链路灰度规则名称不能为空。
	INVALIDPARAMETERVALUE_LANERULENAMENOTEMPTY = "InvalidParameterValue.LaneRuleNameNotEmpty"

	// 全链路灰度规则名称不能超过60个字符。
	INVALIDPARAMETERVALUE_LANERULENAMETOOLONG = "InvalidParameterValue.LaneRuleNameTooLong"

	// 全链路灰度规则不存在。
	INVALIDPARAMETERVALUE_LANERULENOTEXIST = "InvalidParameterValue.LaneRuleNotExist"

	// 全链路灰度规则备注不能超过200个字符。
	INVALIDPARAMETERVALUE_LANERULEREMARKTOOLONG = "InvalidParameterValue.LaneRuleRemarkTooLong"

	// 全链路灰度规则标签名不能为空。
	INVALIDPARAMETERVALUE_LANERULETAGNAMENOTEMPTY = "InvalidParameterValue.LaneRuleTagNameNotEmpty"

	// 全链路灰度规则标签名不能超过32个字符。
	INVALIDPARAMETERVALUE_LANERULETAGNAMETOOLONG = "InvalidParameterValue.LaneRuleTagNameTooLong"

	// 全链路灰度规则必须设置至少一个标签。
	INVALIDPARAMETERVALUE_LANERULETAGNOTEMPTY = "InvalidParameterValue.LaneRuleTagNotEmpty"

	// 全链路灰度规则标签值不能超过128个字符。
	INVALIDPARAMETERVALUE_LANERULETAGVALUETOOLONG = "InvalidParameterValue.LaneRuleTagValueTooLong"

	// 全链路灰度规则总标签值不能超过200个字符。
	INVALIDPARAMETERVALUE_LANERULETAGVALUETOTALTOOLONG = "InvalidParameterValue.LaneRuleTagValueTotalTooLong"

	// 集群已关联该命名空间。
	INVALIDPARAMETERVALUE_NAMESPACEALREADYBINDCLUSTER = "InvalidParameterValue.NamespaceAlreadyBindCluster"

	// 命名空间描格式不正确。
	INVALIDPARAMETERVALUE_NAMESPACEDESCINVALID = "InvalidParameterValue.NamespaceDescInvalid"

	// 命名空间名称已存在，请更换其他名称。
	INVALIDPARAMETERVALUE_NAMESPACENAMEEXIST = "InvalidParameterValue.NamespaceNameExist"

	// 命名空间名称格式不正确。
	INVALIDPARAMETERVALUE_NAMESPACENAMEINVALID = "InvalidParameterValue.NamespaceNameInvalid"

	// 无法获取命名空间。
	INVALIDPARAMETERVALUE_NAMESPACENOTEXISTS = "InvalidParameterValue.NamespaceNotExists"

	// 配置项已经发布，不允许删除。
	INVALIDPARAMETERVALUE_RELEASEDCONFIGCANNOTBEDELETED = "InvalidParameterValue.ReleasedConfigCanNotBeDeleted"

	// 无权限操作资源%s。
	INVALIDPARAMETERVALUE_RESOURCEPERMISSIONDENIED = "InvalidParameterValue.ResourcePermissionDenied"

	// ResourceType 不支持。
	INVALIDPARAMETERVALUE_RESOURCETYPEERROR = "InvalidParameterValue.ResourceTypeError"

	// 服务描述不能大于200字符。
	INVALIDPARAMETERVALUE_SERVICEDESCLENGTH = "InvalidParameterValue.ServiceDescLength"

	// 服务名称重复。
	INVALIDPARAMETERVALUE_SERVICENAMEREPEATED = "InvalidParameterValue.ServiceNameRepeated"

	// 服务不存在或权限不足。
	INVALIDPARAMETERVALUE_SERVICENOTEXISTSORPERMISSIONDENIED = "InvalidParameterValue.ServiceNotExistsOrPermissionDenied"

	// 无效请求参数。
	INVALIDPARAMETERVALUE_TASKPARAMETERINVALID = "InvalidParameterValue.TaskParameterInvalid"

	// 命名空间数达到上限。
	LIMITEXCEEDED_ERRNAMESPACEMAXLIMIT = "LimitExceeded.ErrNamespaceMaxLimit"

	// 仓库达到上限。
	LIMITEXCEEDED_ERRREPOMAXLIMIT = "LimitExceeded.ErrRepoMaxLimit"

	// 最多支持创建五个容器集群，当前已经超过使用上限。
	LIMITEXCEEDED_TKECLUSTERNUMBEREXCEEDLIMIT = "LimitExceeded.TkeClusterNumberExceedLimit"

	// 应用ID不能为空。
	MISSINGPARAMETER_APPLICATIONIDNULL = "MissingParameter.ApplicationIdNull"

	// 应用ID未填写。
	MISSINGPARAMETER_APPLICATIONIDREQUIRED = "MissingParameter.ApplicationIdRequired"

	// 应用类型不能为空。
	MISSINGPARAMETER_APPLICATIONTYPENULL = "MissingParameter.ApplicationTypeNull"

	// 集群所属子网不能为空。
	MISSINGPARAMETER_CLUSTERSUBNETREQUIRED = "MissingParameter.ClusterSubnetRequired"

	// 配置项ID未填写。
	MISSINGPARAMETER_CONFIGIDREQUIRED = "MissingParameter.ConfigIdRequired"

	// 配置项名称未填写。
	MISSINGPARAMETER_CONFIGNAMEREQUIRED = "MissingParameter.ConfigNameRequired"

	// 配置项发布信息ID未填写。
	MISSINGPARAMETER_CONFIGRELEASEIDREQUIRED = "MissingParameter.ConfigReleaseIdRequired"

	// 配置项类型未填写。
	MISSINGPARAMETER_CONFIGTYPEREQUIRED = "MissingParameter.ConfigTypeRequired"

	// 配置项值未填写。
	MISSINGPARAMETER_CONFIGVALUEREQUIRED = "MissingParameter.ConfigValueRequired"

	// 配置项版本未填写。
	MISSINGPARAMETER_CONFIGVERSIONREQUIRED = "MissingParameter.ConfigVersionRequired"

	// 文件配置项文件内容未填写。
	MISSINGPARAMETER_FILECONFIGFILEVALUEREQUIRED = "MissingParameter.FileConfigFileValueRequired"

	// 缺少请求参数:%s。
	MISSINGPARAMETER_GATEWAYPARAMETERREQUIRED = "MissingParameter.GatewayParameterRequired"

	// 分组所属应用不能为空。
	MISSINGPARAMETER_GROUPAPPLICATIONNULL = "MissingParameter.GroupApplicationNull"

	// 分组ID不能为空。
	MISSINGPARAMETER_GROUPIDNULL = "MissingParameter.GroupIdNull"

	// 命名空间ID不能为空。
	MISSINGPARAMETER_NAMESPACEIDREQUIRED = "MissingParameter.NamespaceIdRequired"

	// %s缺失。
	MISSINGPARAMETER_REQUIREDPARAMETERMISSING = "MissingParameter.RequiredParameterMissing"

	// 未填写服务Id。
	MISSINGPARAMETER_SERVICEIDREQUIRED = "MissingParameter.ServiceIdRequired"

	// 缺少必填参数。
	MISSINGPARAMETER_TASKPARAMETERMISSED = "MissingParameter.TaskParameterMissed"

	// 此应用下存在资源，无法执行删除操作。
	RESOURCEINUSE_APPLICATIONCANNOTDELETE = "ResourceInUse.ApplicationCannotDelete"

	// 资源仍在使用中 无法删除。
	RESOURCEINUSE_CVMCAEMASTERCANNOTDELETE = "ResourceInUse.CvmcaeMasterCannotDelete"

	// 默认命名空间不能被删除。
	RESOURCEINUSE_DEFAULTNAMEPSACECANNOTBEDELETED = "ResourceInUse.DefaultNamepsaceCannotBeDeleted"

	// 此分组下存在资源，无法执行删除操作。
	RESOURCEINUSE_GROUPCANNOTDELETE = "ResourceInUse.GroupCannotDelete"

	// 部署组在更新中 请稍后再执行该操作。
	RESOURCEINUSE_GROUPINOPERATION = "ResourceInUse.GroupInOperation"

	// 机器实例已经被使用。
	RESOURCEINUSE_INSTANCEHASBEENUSED = "ResourceInUse.InstanceHasBeenUsed"

	// 此命名空间下存在资源，无法执行删除操作。
	RESOURCEINUSE_NAMESPACECANNOTDELETE = "ResourceInUse.NamespaceCannotDelete"

	// 资源对象已存在。
	RESOURCEINUSE_OBJECTEXIST = "ResourceInUse.ObjectExist"

	// 限流规则已存在，请检查规则名和规则配置。
	RESOURCEINUSE_RATELIMITRULEEXISTERROR = "ResourceInUse.RatelimitRuleExistError"

	// 实例数已达上限。
	RESOURCEINSUFFICIENT_INSTANCEEXCESSLIMIT = "ResourceInsufficient.InstanceExcessLimit"

	// 仓库空间达到上限。
	RESOURCEINSUFFICIENT_PACKAGESPACEFULL = "ResourceInsufficient.PackageSpaceFull"

	// 无法获取应用信息。
	RESOURCENOTFOUND_APPLICATIONNOTEXIST = "ResourceNotFound.ApplicationNotExist"

	// 应用不存在或应用不属于当前项目。
	RESOURCENOTFOUND_APPLICATIONPROJECTNOTMATCH = "ResourceNotFound.ApplicationProjectNotMatch"

	// 无法获取命名空间所属集群。
	RESOURCENOTFOUND_CLUSTERNOTEXIST = "ResourceNotFound.ClusterNotExist"

	// 集群所属私有网络不存在。
	RESOURCENOTFOUND_CLUSTERVPCNOTEXIST = "ResourceNotFound.ClusterVpcNotExist"

	// 找不到集群。
	RESOURCENOTFOUND_CONTAINERGROUPCLUSTERNOTFOUND = "ResourceNotFound.ContainergroupClusterNotfound"

	// 无法找到该部署组所属集群和命名空间。
	RESOURCENOTFOUND_CONTAINERGROUPGROUPNAMESPACECLUSTERNOTFOUND = "ResourceNotFound.ContainergroupGroupNamespaceClusterNotFound"

	// 无法找到该部署组。
	RESOURCENOTFOUND_CONTAINERGROUPGROUPNOTFOUND = "ResourceNotFound.ContainergroupGroupNotFound"

	// TSF MASTER 资源不存在。
	RESOURCENOTFOUND_CVMCAEMASTERRESOURCENOTFOUND = "ResourceNotFound.CvmcaeMasterResourceNotFound"

	// 镜像仓库不存在。
	RESOURCENOTFOUND_ERRNOREPO = "ResourceNotFound.ErrNoRepo"

	// 用户错误。
	RESOURCENOTFOUND_ERRNOUSER = "ResourceNotFound.ErrNoUser"

	// 无法获取分组所属应用。
	RESOURCENOTFOUND_GROUPAPPLICATIONNOTEXIST = "ResourceNotFound.GroupApplicationNotExist"

	// 无法获取分组所属命名空间。
	RESOURCENOTFOUND_GROUPNAMESPACENOTEXIST = "ResourceNotFound.GroupNamespaceNotExist"

	// 此部署组不存在，无法执行该操作。
	RESOURCENOTFOUND_GROUPNOTEXIST = "ResourceNotFound.GroupNotExist"

	// 无法获取机器信息。
	RESOURCENOTFOUND_INSTANCENOTEXIST = "ResourceNotFound.InstanceNotExist"

	// 无法找到License服务器。
	RESOURCENOTFOUND_LICENSESERVERNOTFOUND = "ResourceNotFound.LicenseServerNotFound"

	// 目标微服务已离线[%s]。。
	RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"

	// 无法获取命名空间。
	RESOURCENOTFOUND_NAMESPACENOTEXIST = "ResourceNotFound.NamespaceNotExist"

	// 资源对象不存在。
	RESOURCENOTFOUND_OBJECTNOEXIST = "ResourceNotFound.ObjectNoExist"

	// 无法获取服务，无法执行该操作。
	RESOURCENOTFOUND_SERVICENOTEXIST = "ResourceNotFound.ServiceNotExist"

	// 任务不存在。
	RESOURCENOTFOUND_TASKNOTFOUND = "ResourceNotFound.TaskNotFound"

	// TKE 中不存在该集群。
	RESOURCENOTFOUND_TKECLUSTERNOTEXISTS = "ResourceNotFound.TkeClusterNotExists"

	// 访问 CAM 系统出错，%s。
	UNAUTHORIZEDOPERATION_CAMGENERALERROR = "UnauthorizedOperation.CamGeneralError"

	// 协作者身份未授权，需要主账号授予协作者权限，参考 TSF 官网文档「快速入门/准备工作」。
	UNAUTHORIZEDOPERATION_CAMTSFROLENOPERMISSION = "UnauthorizedOperation.CamTsfRoleNoPermission"

	// 当前主账号未创建TSF_QCSRole或未对子账号授予预设策略QcloudCamSubaccountsAuthorizeRoleFullAccess。请参考产品文档主账号协作者使用说明。。
	UNAUTHORIZEDOPERATION_CAMTSFROLENOTEXIST = "UnauthorizedOperation.CamTsfRoleNotExist"

	// License未激活。。
	UNAUTHORIZEDOPERATION_LICENSEINACTIVE = "UnauthorizedOperation.LicenseInactive"

	// 您所购买的服务不支持该操作。
	UNAUTHORIZEDOPERATION_LICENSEUNAUTHORIZED = "UnauthorizedOperation.LicenseUnauthorized"

	// 缺少License。。
	UNAUTHORIZEDOPERATION_NOLICENSE = "UnauthorizedOperation.NoLicense"

	// 用户无权限访问该接口。。
	UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"

	// 批量操作数量超过限制:%s。
	UNSUPPORTEDOPERATION_GATEWAYTOOMANYREQUESTPARAMETER = "UnsupportedOperation.GatewayTooManyRequestParameter"

	// 操作不支持。
	UNSUPPORTEDOPERATION_TASKNOTSUPPORTED = "UnsupportedOperation.TaskNotSupported"

	// 不支持的ACTION。。
	UNSUPPORTEDOPERATION_UNSUPPORTACTION = "UnsupportedOperation.UnsupportAction"
)
