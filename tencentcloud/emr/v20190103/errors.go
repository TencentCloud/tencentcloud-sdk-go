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

package v20190103

const (
	// 此产品的特有错误码

	// 操作失败。
	FAILEDOPERATION = "FailedOperation"

	// 重复的订单，请检查emr控制台。
	FAILEDOPERATION_DUPLICATEORDERNOTALLOWED = "FailedOperation.DuplicateOrderNotAllowed"

	// 内部错误。
	INTERNALERROR = "InternalError"

	// 内部服务调用异常。
	INTERNALERROR_ACCOUNTCGWERROR = "InternalError.AccountCgwError"

	// 内部服务调用异常。
	INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"

	// 内部服务调用异常。
	INTERNALERROR_CAMERROR = "InternalError.CamError"

	// 内部服务调用异常。
	INTERNALERROR_CBSCGWERROR = "InternalError.CbsCgwError"

	// 内部服务调用异常。
	INTERNALERROR_CBSERROR = "InternalError.CbsError"

	// 内部服务调用异常。
	INTERNALERROR_CDBCGWERROR = "InternalError.CdbCgwError"

	// 内部服务调用异常。
	INTERNALERROR_CDBERROR = "InternalError.CdbError"

	// cvm或cbs资源不够或软件不合法。
	INTERNALERROR_CHECKQUOTAERR = "InternalError.CheckQuotaErr"

	// 内部服务调用异常。
	INTERNALERROR_CONFIGCGWERROR = "InternalError.ConfigCgwError"

	// 内部服务调用异常。
	INTERNALERROR_CVMERROR = "InternalError.CvmError"

	// 内部服务调用异常。
	INTERNALERROR_KMSERROR = "InternalError.KmsError"

	// 内部服务调用异常。
	INTERNALERROR_PROJECTCGWERROR = "InternalError.ProjectCgwError"

	// 安全组接口调用异常。
	INTERNALERROR_SGERROR = "InternalError.SgError"

	// TKE调用出错。
	INTERNALERROR_TKEERROR = "InternalError.TKEError"

	// 内部服务调用异常。
	INTERNALERROR_TAGERROR = "InternalError.TagError"

	// 内部服务调用异常。
	INTERNALERROR_TRADECGWERROR = "InternalError.TradeCgwError"

	// 内部服务调用异常。
	INTERNALERROR_VPCCGWERROR = "InternalError.VpcCgwError"

	// 内部服务调用异常。
	INTERNALERROR_VPCERROR = "InternalError.VpcError"

	// 内部服务调用异常。
	INTERNALERROR_WOODSERVERERROR = "InternalError.WoodServerError"

	// 参数错误。
	INVALIDPARAMETER = "InvalidParameter"

	// 展示策略错误。
	INVALIDPARAMETER_DISPLAYSTRATEGYNOTMATCH = "InvalidParameter.DisplayStrategyNotMatch"

	// 参数错误。
	INVALIDPARAMETER_HALESSMASTERCOUNT = "InvalidParameter.HALessMasterCount"

	// Common节点数量无效。
	INVALIDPARAMETER_INCORRECTCOMMONCOUNT = "InvalidParameter.IncorrectCommonCount"

	// Master节点数量无效。
	INVALIDPARAMETER_INCORRECTMASTERCOUNT = "InvalidParameter.IncorrectMasterCount"

	// core节点的数量不能超过20。
	INVALIDPARAMETER_INVAILDCORECOUNT = "InvalidParameter.InvaildCoreCount"

	// 无效参数，AppId。
	INVALIDPARAMETER_INVALIDAPPID = "InvalidParameter.InvalidAppId"

	// 无效的自动续费标识。
	INVALIDPARAMETER_INVALIDAUTORENEW = "InvalidParameter.InvalidAutoRenew"

	// 无效的引导脚本。
	INVALIDPARAMETER_INVALIDBOOTSTRAPACTION = "InvalidParameter.InvalidBootstrapAction"

	// 无效的ClientToken。
	INVALIDPARAMETER_INVALIDCLIENTTOKEN = "InvalidParameter.InvalidClientToken"

	// 无效参数，ClusterId。
	INVALIDPARAMETER_INVALIDCLUSTERID = "InvalidParameter.InvalidClusterId"

	// 无效的组件。
	INVALIDPARAMETER_INVALIDCOMPONENT = "InvalidParameter.InvalidComponent"

	// Core节点数量无效。
	INVALIDPARAMETER_INVALIDCORECOUNT = "InvalidParameter.InvalidCoreCount"

	// 扩容数量必须大于0。
	INVALIDPARAMETER_INVALIDCOUNT = "InvalidParameter.InvalidCount"

	// 同一请求只能扩容Task或者Core节点。
	INVALIDPARAMETER_INVALIDCOUNTNUM = "InvalidParameter.InvalidCountNum"

	// 错误信息：Invalid PodParameter。
	INVALIDPARAMETER_INVALIDCUSTOMIZEDPODPARAM = "InvalidParameter.InvalidCustomizedPodParam"

	// 无效的磁盘大小。
	INVALIDPARAMETER_INVALIDDISKSIZE = "InvalidParameter.InvalidDiskSize"

	// 无效的EKS实例。
	INVALIDPARAMETER_INVALIDEKSINSTANCE = "InvalidParameter.InvalidEksInstance"

	// CustomConfig参数值无效。
	INVALIDPARAMETER_INVALIDEXTENDFIELD = "InvalidParameter.InvalidExtendField"

	// 无效的任务失败处理策略。
	INVALIDPARAMETER_INVALIDFAILUREPOLICY = "InvalidParameter.InvalidFailurePolicy"

	// 无效参数，EMR实例不符合要求。
	INVALIDPARAMETER_INVALIDINSTANCE = "InvalidParameter.InvalidInstance"

	// 无效的集群名称。
	INVALIDPARAMETER_INVALIDINSTANCENAME = "InvalidParameter.InvalidInstanceName"

	// 无效的集群保留策略。
	INVALIDPARAMETER_INVALIDINSTANCEPOLICY = "InvalidParameter.InvalidInstancePolicy"

	// 无效的机型。
	INVALIDPARAMETER_INVALIDINSTANCETYPE = "InvalidParameter.InvalidInstanceType"

	// 无效的流程任务。
	INVALIDPARAMETER_INVALIDJOBFLOW = "InvalidParameter.InvalidJobFlow"

	// 无效的任务步骤类型。
	INVALIDPARAMETER_INVALIDJOBTYPE = "InvalidParameter.InvalidJobType"

	// 无效的登录设置。
	INVALIDPARAMETER_INVALIDLOGINSETTING = "InvalidParameter.InvalidLoginSetting"

	// 无效的元数据表类型。
	INVALIDPARAMETER_INVALIDMETATYPE = "InvalidParameter.InvalidMetaType"

	// 变配规格无效。
	INVALIDPARAMETER_INVALIDMODIFYSPEC = "InvalidParameter.InvalidModifySpec"

	// 无效的NodeType。
	INVALIDPARAMETER_INVALIDNODETYPE = "InvalidParameter.InvalidNodeType"

	// 无效密码。
	INVALIDPARAMETER_INVALIDPASSWORD = "InvalidParameter.InvalidPassword"

	// 无效的付费类型。
	INVALIDPARAMETER_INVALIDPAYMODE = "InvalidParameter.InvalidPaymode"

	// 无效的引导操作脚本。
	INVALIDPARAMETER_INVALIDPREEXECUTEDFILE = "InvalidParameter.InvalidPreExecutedFile"

	// 无效参数，不符合EMR版本。
	INVALIDPARAMETER_INVALIDPRODUCT = "InvalidParameter.InvalidProduct"

	// 无效的产品ID。
	INVALIDPARAMETER_INVALIDPRODUCTID = "InvalidParameter.InvalidProductId"

	// 无效的项目ID。
	INVALIDPARAMETER_INVALIDPROJECTID = "InvalidParameter.InvalidProjectId"

	// 资源ID无效。
	INVALIDPARAMETER_INVALIDRESOURCEIDS = "InvalidParameter.InvalidResourceIds"

	// 无效的资源规格。
	INVALIDPARAMETER_INVALIDRESOURCESPEC = "InvalidParameter.InvalidResourceSpec"

	// 无效的安全组ID。
	INVALIDPARAMETER_INVALIDSERCURITYGRPUPID = "InvalidParameter.InvalidSercurityGrpupId"

	// 服务名无效。
	INVALIDPARAMETER_INVALIDSERVICENAME = "InvalidParameter.InvalidServiceName"

	// 参数ServiceNodeInfo无效或错误。
	INVALIDPARAMETER_INVALIDSERVICENODEINFO = "InvalidParameter.InvalidServiceNodeInfo"

	// 参数InvalidSoftDeployInfo无效或错误。
	INVALIDPARAMETER_INVALIDSOFTDEPLOYINFO = "InvalidParameter.InvalidSoftDeployInfo"

	// 无效的SoftInfo。
	INVALIDPARAMETER_INVALIDSOFTINFO = "InvalidParameter.InvalidSoftInfo"

	// 参数错误。
	INVALIDPARAMETER_INVALIDSOFTWARE = "InvalidParameter.InvalidSoftWare"

	// 软件名无效。
	INVALIDPARAMETER_INVALIDSOFTWARENAME = "InvalidParameter.InvalidSoftWareName"

	// 软件版本无效。
	INVALIDPARAMETER_INVALIDSOFTWAREVERSION = "InvalidParameter.InvalidSoftWareVersion"

	// 无效的子网ID。
	INVALIDPARAMETER_INVALIDSUBNETID = "InvalidParameter.InvalidSubnetId"

	// 无效的高可用参数。
	INVALIDPARAMETER_INVALIDSUPPORTHA = "InvalidParameter.InvalidSupportHA"

	// task的数量不能超过20。
	INVALIDPARAMETER_INVALIDTASKCOUNT = "InvalidParameter.InvalidTaskCount"

	// 无效的timespan。
	INVALIDPARAMETER_INVALIDTIMESPAN = "InvalidParameter.InvalidTimeSpan"

	// 无效的TimeUnit。
	INVALIDPARAMETER_INVALIDTIMEUNIT = "InvalidParameter.InvalidTimeUnit"

	// 无效的Tke集群ID，或Tke集群不符合条件。
	INVALIDPARAMETER_INVALIDTKEINSTANCE = "InvalidParameter.InvalidTkeInstance"

	// 无效的统一元数据库。
	INVALIDPARAMETER_INVALIDUNIFYMETA = "InvalidParameter.InvalidUnifyMeta"

	// 无效的私有网络ID。
	INVALIDPARAMETER_INVALIDVPCID = "InvalidParameter.InvalidVpcId"

	// 无效的可用区。
	INVALIDPARAMETER_INVALIDZONE = "InvalidParameter.InvalidZone"

	// 无效参数，不满足必须组件。
	INVALIDPARAMETER_NOTCONTAINMUSTSELECTSOFTWARE = "InvalidParameter.NotContainMustSelectSoftware"

	// 排序字段错误。
	INVALIDPARAMETER_ORDERFIELDNOTMATCH = "InvalidParameter.OrderFieldNotMatch"

	// 付费模式与资源不匹配。
	INVALIDPARAMETER_PAYMODERESOURCENOTMATCH = "InvalidParameter.PayModeResourceNotMatch"

	// 项目与资源不匹配。
	INVALIDPARAMETER_PROJECTRESOURCENOTMATCH = "InvalidParameter.ProjectResourceNotMatch"

	// 存在无效的产品组件。
	INVALIDPARAMETER_SOFTWARENOTINPRODUCT = "InvalidParameter.SoftwareNotInProduct"

	// 策略为授权。
	INVALIDPARAMETER_UNGRANTEDPOLICY = "InvalidParameter.UngrantedPolicy"

	// 角色未授权。
	INVALIDPARAMETER_UNGRANTEDROLE = "InvalidParameter.UngrantedRole"

	// 可用区与资源不匹配。
	INVALIDPARAMETER_ZONERESOURCENOTMATCH = "InvalidParameter.ZoneResourceNotMatch"

	// 参数取值错误。
	INVALIDPARAMETERVALUE = "InvalidParameterValue"

	// 无效的Tke集群ID，或Tke集群不符合条件。
	INVALIDPARAMETERVALUE_INVALIDTKEINSTANCE = "InvalidParameterValue.InvalidTkeInstance"

	// 引导脚本数量超过限制。
	LIMITEXCEEDED_BOOTSTRAPACTIONSNUMLIMITEXCEEDED = "LimitExceeded.BootstrapActionsNumLimitExceeded"

	// 缺少参数错误。
	MISSINGPARAMETER = "MissingParameter"

	// 实例在流程中。
	RESOURCEINUSE_INSTANCEINPROCESS = "ResourceInUse.InstanceInProcess"

	// 硬盘规格不满足。
	RESOURCEINSUFFICIENT_DISKINSUFFICIENT = "ResourceInsufficient.DiskInsufficient"

	// 不支持或售罄的节点规格。
	RESOURCEINSUFFICIENT_INSTANCEINSUFFICIENT = "ResourceInsufficient.InstanceInsufficient"

	// 无法找到该实例。
	RESOURCENOTFOUND_CLUSTERNOTFOUND = "ResourceNotFound.ClusterNotFound"

	// 无法找到硬件信息。
	RESOURCENOTFOUND_HARDWAREINFONOTFOUND = "ResourceNotFound.HardwareInfoNotFound"

	// 无法找到该实例。
	RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"

	// 无法找到监控元数据。
	RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"

	// 找不到对应的子网。
	RESOURCENOTFOUND_SUBNETNOTFOUND = "ResourceNotFound.SubnetNotFound"

	// tke集群前置组件未部署。
	RESOURCENOTFOUND_TKEPRECONDITIONNOTFOUND = "ResourceNotFound.TKEPreconditionNotFound"

	// 没有查找到指定标签。
	RESOURCENOTFOUND_TAGSNOTFOUND = "ResourceNotFound.TagsNotFound"

	// 当前资源规格不存在默认规格。
	RESOURCEUNAVAILABLE_RESOURCESPECNOTDEFAULTSPEC = "ResourceUnavailable.ResourceSpecNotDefaultSpec"

	// 当前资源规格不存在默认规格。
	RESOURCEUNAVAILABLE_RESOURCESPEC_NOTDEFAULTSPEC = "ResourceUnavailable.ResourceSpec_NotDefaultSpec"

	// 资源售罄。
	RESOURCESSOLDOUT_ = "ResourcesSoldOut."

	// CBS资源售罄。
	RESOURCESSOLDOUT_CBSSOLDOUT = "ResourcesSoldOut.CbsSoldOut"

	// 云服务器已售罄。
	RESOURCESSOLDOUT_CVMSOLDOUT = "ResourcesSoldOut.CvmSoldOut"

	// 未知参数错误。
	UNKNOWNPARAMETER = "UnknownParameter"

	// 操作不支持。
	UNSUPPORTEDOPERATION = "UnsupportedOperation"

	// 该服务不支持此操作。
	UNSUPPORTEDOPERATION_SERVICENOTSUPPORT = "UnsupportedOperation.ServiceNotSupport"
)
