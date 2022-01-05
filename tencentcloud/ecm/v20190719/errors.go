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

package v20190719

const (
	// 此产品的特有错误码

	// 鉴权失败。
	AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"

	// 操作失败。
	FAILEDOPERATION = "FailedOperation"

	// 余额不足。
	FAILEDOPERATION_BLOCKBALANCE = "FailedOperation.BlockBalance"

	// 数据操作失败。
	FAILEDOPERATION_DATAOPERATIONFAILED = "FailedOperation.DataOperationFailed"

	// FailedOperation.DiskAttached
	FAILEDOPERATION_DISKATTACHED = "FailedOperation.DiskAttached"

	// 镜像正在使用中。
	FAILEDOPERATION_IMAGEINUSE = "FailedOperation.ImageInUse"

	// 模块下存在实例。
	FAILEDOPERATION_INSTANCEINMODULE = "FailedOperation.InstanceInModule"

	// 实例没有全部关机。
	FAILEDOPERATION_INSTANCENOTALLSTOPPED = "FailedOperation.InstanceNotAllStopped"

	// 实例不属于当前模块。
	FAILEDOPERATION_INSTANCEOWNERCHECKFAILED = "FailedOperation.InstanceOwnerCheckFailed"

	// 内部操作错误。
	FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"

	// 当前状态无法执行该操作。
	FAILEDOPERATION_INVALIDSTATUS = "FailedOperation.InvalidStatus"

	// 其他操作正在运行，无法进行当前操作。
	FAILEDOPERATION_OPERATIONCONFLICT = "FailedOperation.OperationConflict"

	// 不允许执行当前操作。
	FAILEDOPERATION_OPERATIONNOTALLOW = "FailedOperation.OperationNotAllow"

	// 该内网IP已经绑定了弹性IP。
	FAILEDOPERATION_PRIVATEIPADDRESSBINDED = "FailedOperation.PrivateIpAddressBinded"

	// 内网IP状态非可用状态。
	FAILEDOPERATION_PRIVATEIPADDRESSUNAVAILABLE = "FailedOperation.PrivateIpAddressUnavailable"

	// 内部错误。
	INTERNALERROR = "InternalError"

	// InternalError.ComponentError
	INTERNALERROR_COMPONENTERROR = "InternalError.ComponentError"

	// InternalError.FailQueryResource
	INTERNALERROR_FAILQUERYRESOURCE = "InternalError.FailQueryResource"

	// InternalError.ResourceOpFailed
	INTERNALERROR_RESOURCEOPFAILED = "InternalError.ResourceOpFailed"

	// 参数错误。
	INVALIDPARAMETER = "InvalidParameter"

	// 参数不支持同时指定。
	INVALIDPARAMETER_COEXIST = "InvalidParameter.Coexist"

	// InvalidParameter.DiskConfigNotSupported
	INVALIDPARAMETER_DISKCONFIGNOTSUPPORTED = "InvalidParameter.DiskConfigNotSupported"

	// 参数格式错误。
	INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"

	// 数据格式不正确。
	INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"

	// 指定的两个参数冲突，不能同时存在。 EIP只能绑定在实例上或指定网卡的指定内网 IP 上。
	INVALIDPARAMETER_INVALIDPARAMETERCONFLICT = "InvalidParameter.InvalidParameterConflict"

	// 负载均衡实例ID错误。
	INVALIDPARAMETER_LBIDNOTFOUND = "InvalidParameter.LBIdNotFound"

	// 监听器ID错误。
	INVALIDPARAMETER_LISTENERIDNOTFOUND = "InvalidParameter.ListenerIdNotFound"

	// 查找不到符合条件的转发规则。
	INVALIDPARAMETER_LOCATIONNOTFOUND = "InvalidParameter.LocationNotFound"

	// KeepImageLogin, Password, KeyIds 同时只能使用1个。
	INVALIDPARAMETER_PARAMETERCONFLICT = "InvalidParameter.ParameterConflict"

	// 监听器端口检查失败，比如端口冲突。
	INVALIDPARAMETER_PORTCHECKFAILED = "InvalidParameter.PortCheckFailed"

	// InvalidParameter.ProjectIdNotExist
	INVALIDPARAMETER_PROJECTIDNOTEXIST = "InvalidParameter.ProjectIdNotExist"

	// 监听器协议检查失败，比如相关协议不支持对应操作。
	INVALIDPARAMETER_PROTOCOLCHECKFAILED = "InvalidParameter.ProtocolCheckFailed"

	// 地域无效。
	INVALIDPARAMETER_REGIONNOTFOUND = "InvalidParameter.RegionNotFound"

	// 参数取值错误。
	INVALIDPARAMETERVALUE = "InvalidParameterValue"

	// 申请数量超限。
	INVALIDPARAMETERVALUE_ADDRESSQUOTALIMITEXCEEDED = "InvalidParameterValue.AddressQuotaLimitExceeded"

	// 超多带宽限制。
	INVALIDPARAMETERVALUE_BANDWIDTHOUTOFRANGE = "InvalidParameterValue.BandwidthOutOfRange"

	// 目的网段不在对端VPC的CIDR范围内。
	INVALIDPARAMETERVALUE_CIDRNOTINPEERVPC = "InvalidParameterValue.CidrNotInPeerVpc"

	// 重复数据。
	INVALIDPARAMETERVALUE_DUPLICATE = "InvalidParameterValue.Duplicate"

	// Filter参数数值超过限制。
	INVALIDPARAMETERVALUE_FILTERLIMITEXCEEDED = "InvalidParameterValue.FilterLimitExceeded"

	// 镜像已导入。
	INVALIDPARAMETERVALUE_IMAGEDUPLICATE = "InvalidParameterValue.ImageDuplicate"

	// 镜像名称重复。
	INVALIDPARAMETERVALUE_IMAGENAMEDUPLICATE = "InvalidParameterValue.ImageNameDuplicate"

	// 镜像大小大于系统盘大小。
	INVALIDPARAMETERVALUE_IMAGESIZELARGETHANSYSDISKSIZE = "InvalidParameterValue.ImageSizeLargeThanSysDiskSize"

	// 实例配置不匹配。
	INVALIDPARAMETERVALUE_INSTANCECONFIGNOTMATCH = "InvalidParameterValue.InstanceConfigNotMatch"

	// 实例不支持当前操作。
	INVALIDPARAMETERVALUE_INSTANCEIDNOTSUPPORTED = "InvalidParameterValue.InstanceIdNotSupported"

	// 实例名字长度超出限制。
	INVALIDPARAMETERVALUE_INSTANCENAMETOOLONG = "InvalidParameterValue.InstanceNameTooLong"

	// 实例机型和pid不匹配。
	INVALIDPARAMETERVALUE_INSTANCETYPENOTMATCHPID = "InvalidParameterValue.InstanceTypeNotMatchPid"

	// InvalidParameterValue.InsufficientRefundQuota
	INVALIDPARAMETERVALUE_INSUFFICIENTREFUNDQUOTA = "InvalidParameterValue.InsufficientRefundQuota"

	// 无效的EIP。
	INVALIDPARAMETERVALUE_INVAILDADDRESSID = "InvalidParameterValue.InvaildAddressId"

	// 您的输入有误或无相应的操作权限。
	INVALIDPARAMETERVALUE_INVAILDDESCRIBEINSTANCE = "InvalidParameterValue.InvaildDescribeInstance"

	// 无效的弹性网卡ID。
	INVALIDPARAMETERVALUE_INVAILDENIID = "InvalidParameterValue.InvaildEniID"

	// HostName不合法。
	INVALIDPARAMETERVALUE_INVAILDHOSTNAME = "InvalidParameterValue.InvaildHostName"

	// 参数错误。
	INVALIDPARAMETERVALUE_INVAILDMODIFYPARAM = "InvalidParameterValue.InvaildModifyParam"

	// 模块数量不合法。
	INVALIDPARAMETERVALUE_INVAILDMODULENUM = "InvalidParameterValue.InvaildModuleNum"

	// 分页参数不合法。
	INVALIDPARAMETERVALUE_INVAILDPAGEPARAM = "InvalidParameterValue.InvaildPageParam"

	// 密码不合法。
	INVALIDPARAMETERVALUE_INVAILDPASSWORD = "InvalidParameterValue.InvaildPassword"

	// 当同步绑定弹性网卡时，一次只能申请单个弹性IP。
	INVALIDPARAMETERVALUE_INVALIDADDRESSCOUNT = "InvalidParameterValue.InvalidAddressCount"

	// 带宽大小不合法。
	INVALIDPARAMETERVALUE_INVALIDBANDWIDTH = "InvalidParameterValue.InvalidBandwidth"

	// 入带宽大小超过限额。
	INVALIDPARAMETERVALUE_INVALIDBANDWIDTHIN = "InvalidParameterValue.InvalidBandwidthIn"

	// 出入带宽必须一致。
	INVALIDPARAMETERVALUE_INVALIDBANDWIDTHINANDOUT = "InvalidParameterValue.InvalidBandwidthInAndOut"

	// 该用户不支持当前所选择的计费类型。
	INVALIDPARAMETERVALUE_INVALIDBILLINGTYPE = "InvalidParameterValue.InvalidBillingType"

	// 数据盘个数不合法。
	INVALIDPARAMETERVALUE_INVALIDDATADISKNUM = "InvalidParameterValue.InvalidDataDiskNum"

	// 数据盘大小不合法。
	INVALIDPARAMETERVALUE_INVALIDDATADISKSIZE = "InvalidParameterValue.InvalidDataDiskSize"

	// 数据盘类型不合法。
	INVALIDPARAMETERVALUE_INVALIDDATADISKTYPE = "InvalidParameterValue.InvalidDataDiskType"

	// InvalidParameterValue.InvalidDisk
	INVALIDPARAMETERVALUE_INVALIDDISK = "InvalidParameterValue.InvalidDisk"

	// InvalidParameterValue.InvalidDiskId
	INVALIDPARAMETERVALUE_INVALIDDISKID = "InvalidParameterValue.InvalidDiskId"

	// InvalidParameterValue.InvalidDiskType
	INVALIDPARAMETERVALUE_INVALIDDISKTYPE = "InvalidParameterValue.InvalidDiskType"

	// windows镜像不支持IP直通。
	INVALIDPARAMETERVALUE_INVALIDEIPDIRECTSERVICE = "InvalidParameterValue.InvalidEIPDirectService"

	// 地域参数不合法。
	INVALIDPARAMETERVALUE_INVALIDECMREGION = "InvalidParameterValue.InvalidEcmRegion"

	// 无效过滤器。
	INVALIDPARAMETERVALUE_INVALIDFILTER = "InvalidParameterValue.InvalidFilter"

	// 运营商不合法。
	INVALIDPARAMETERVALUE_INVALIDISPINNODE = "InvalidParameterValue.InvalidISPInNode"

	// 内存4G以上机型不支持32位操作系统镜像。
	INVALIDPARAMETERVALUE_INVALIDIMAGEARCHITECTURE = "InvalidParameterValue.InvalidImageArchitecture"

	// 镜像ID不合法。
	INVALIDPARAMETERVALUE_INVALIDIMAGEID = "InvalidParameterValue.InvalidImageID"

	// 实例计费类型不合法。
	INVALIDPARAMETERVALUE_INVALIDINSTANCECHARGETYPE = "InvalidParameterValue.InvalidInstanceChargeType"

	// 实例ID不合法。
	INVALIDPARAMETERVALUE_INVALIDINSTANCEID = "InvalidParameterValue.InvalidInstanceID"

	// 不支持调整高IO机型配置。
	INVALIDPARAMETERVALUE_INVALIDINSTANCETYPE = "InvalidParameterValue.InvalidInstanceType"

	// 机型ID不合法。
	INVALIDPARAMETERVALUE_INVALIDINSTANCETYPECONFIGID = "InvalidParameterValue.InvalidInstanceTypeConfigID"

	// 计费类型不合法。
	INVALIDPARAMETERVALUE_INVALIDINTERNETCHARGETYPE = "InvalidParameterValue.InvalidInternetChargeType"

	// 秘钥ID无效。
	INVALIDPARAMETERVALUE_INVALIDKEYPAIRID = "InvalidParameterValue.InvalidKeyPairId"

	// 秘钥名称异常。
	INVALIDPARAMETERVALUE_INVALIDKEYPAIRNAME = "InvalidParameterValue.InvalidKeyPairName"

	// 负载均衡数量不合法。
	INVALIDPARAMETERVALUE_INVALIDLOADBALANCERNUM = "InvalidParameterValue.InvalidLoadBalancerNum"

	// 负载均衡类型不合法。
	INVALIDPARAMETERVALUE_INVALIDLOADBALANCERTYPE = "InvalidParameterValue.InvalidLoadBalancerType"

	// 模块ID不合法。
	INVALIDPARAMETERVALUE_INVALIDMODULEID = "InvalidParameterValue.InvalidModuleID"

	// ModuleId和InstanceType参数不可同时为空。
	INVALIDPARAMETERVALUE_INVALIDMODULEIDANDINSTANCETYPECONFIGID = "InvalidParameterValue.InvalidModuleIDAndInstanceTypeConfigID"

	// ModuleId和InstanceType参数不可同时为空。
	INVALIDPARAMETERVALUE_INVALIDMODULEIDANDINSTANCETYPEID = "InvalidParameterValue.InvalidModuleIDAndInstanceTypeID"

	// 模块名称冲突。
	INVALIDPARAMETERVALUE_INVALIDMODULENAME = "InvalidParameterValue.InvalidModuleName"

	// 排序字段不合法。
	INVALIDPARAMETERVALUE_INVALIDORDERBYFIELD = "InvalidParameterValue.InvalidOrderByField"

	// 公共参数不合法。
	INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"

	// ecm region不合法。
	INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"

	// 全组ID不合法。
	INVALIDPARAMETERVALUE_INVALIDSECURITYGROUPID = "InvalidParameterValue.InvalidSecurityGroupID"

	// InvalidParameterValue.InvalidSnapshot
	INVALIDPARAMETERVALUE_INVALIDSNAPSHOT = "InvalidParameterValue.InvalidSnapshot"

	// InvalidParameterValue.InvalidSnapshotId
	INVALIDPARAMETERVALUE_INVALIDSNAPSHOTID = "InvalidParameterValue.InvalidSnapshotId"

	// 系统盘大小不合法。
	INVALIDPARAMETERVALUE_INVALIDSYSTEMDISKSIZE = "InvalidParameterValue.InvalidSystemDiskSize"

	// 系统盘类型不合法。
	INVALIDPARAMETERVALUE_INVALIDSYSTEMDISKTYPE = "InvalidParameterValue.InvalidSystemDiskType"

	// 时间不合法。
	INVALIDPARAMETERVALUE_INVALIDTIME = "InvalidParameterValue.InvalidTime"

	// zone不合法。
	INVALIDPARAMETERVALUE_INVALIDZONE = "InvalidParameterValue.InvalidZone"

	// 城市或者数量不合法。
	INVALIDPARAMETERVALUE_INVALIDZONEINSTANCECOUNT = "InvalidParameterValue.InvalidZoneInstanceCount"

	// 用区不支持此机型。
	INVALIDPARAMETERVALUE_INVALIDZONEINSTANCETYPE = "InvalidParameterValue.InvalidZoneInstanceType"

	// 参数长度错误。
	INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"

	// 参数值超出限制。
	INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"

	// 入参格式不合法。
	INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"

	// 已有实例的边缘模块不支持调整默认配置。
	INVALIDPARAMETERVALUE_MODULENOTALLOWCHANGE = "InvalidParameterValue.ModuleNotAllowChange"

	// 节点不支持ipv6。
	INVALIDPARAMETERVALUE_NODENOTSUPPORTIPV6 = "InvalidParameterValue.NodeNotSupportIPv6"

	// 对象不在当前子网内。
	INVALIDPARAMETERVALUE_OBJECTNOTCURRENTSUBNET = "InvalidParameterValue.ObjectNotCurrentSubnet"

	// 指定对象不是当前的私有网络。
	INVALIDPARAMETERVALUE_OBJECTVPCNOTCURRENTVPC = "InvalidParameterValue.ObjectVpcNotCurrentVpc"

	// 参数值超出限制。
	INVALIDPARAMETERVALUE_PARAMETERVALUETOOLARGE = "InvalidParameterValue.ParameterValueTooLarge"

	// 参数值不在指定范围。
	INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"

	// 参数值是一个系统保留对象。
	INVALIDPARAMETERVALUE_RESERVED = "InvalidParameterValue.Reserved"

	// 指定CIDR与同一个私有网络内的另一个子网发生冲突。
	INVALIDPARAMETERVALUE_SUBNETCONFLICT = "InvalidParameterValue.SubnetConflict"

	// 无效的子网网段。
	INVALIDPARAMETERVALUE_SUBNETRANGE = "InvalidParameterValue.SubnetRange"

	// 标签数量超过限制。
	INVALIDPARAMETERVALUE_TAGNUMOUTOFRANGE = "InvalidParameterValue.TagNumOutOfRange"

	// 无法找到任务。
	INVALIDPARAMETERVALUE_TASKNOTFOUND = "InvalidParameterValue.TaskNotFound"

	// 定时销毁时间早于当前时间。
	INVALIDPARAMETERVALUE_TERMINATETIMESMALLER = "InvalidParameterValue.TerminateTimeSmaller"

	// 参数值太长。
	INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"

	// 该机型不支持当前所选择的计费类型。
	INVALIDPARAMETERVALUE_UNMATCHEDBILLINGTYPE = "InvalidParameterValue.UnmatchedBillingType"

	// 不支持该区域。
	INVALIDPARAMETERVALUE_UNSUPPORTEDREGION = "InvalidParameterValue.UnsupportedRegion"

	// 用户不支持ipv6。
	INVALIDPARAMETERVALUE_USERNOTSUPPORTIPV6 = "InvalidParameterValue.UserNotSupportIPv6"

	// 目的网段和当前VPC的CIDR冲突。
	INVALIDPARAMETERVALUE_VPCCIDRCONFLICT = "InvalidParameterValue.VpcCidrConflict"

	// 超过配额限制。
	LIMITEXCEEDED = "LimitExceeded"

	// 申请数量超限。
	LIMITEXCEEDED_ADDRESSQUOTALIMITEXCEEDED = "LimitExceeded.AddressQuotaLimitExceeded"

	// 申购次数不足，每个腾讯云账户每个地域每天申购次数为配额数*2 次。
	LIMITEXCEEDED_ADDRESSQUOTALIMITEXCEEDEDDAILYALLOCATE = "LimitExceeded.AddressQuotaLimitExceededDailyAllocate"

	// LimitExceeded.AttachedDiskLimitExceeded
	LIMITEXCEEDED_ATTACHEDDISKLIMITEXCEEDED = "LimitExceeded.AttachedDiskLimitExceeded"

	// 网卡配额不足。
	LIMITEXCEEDED_ENIQUOTALIMITEXCEEDED = "LimitExceeded.EniQuotaLimitExceeded"

	// 实例绑定的安全组超出限制。
	LIMITEXCEEDED_INSTANCESECURITYGROUPLIMITEXCEEDED = "LimitExceeded.InstanceSecurityGroupLimitExceeded"

	// 申请的clb超过限制。
	LIMITEXCEEDED_LBLIMITEXCEEDED = "LimitExceeded.LBLimitExceeded"

	// 模块默认安全组数目限制。
	LIMITEXCEEDED_MODULESECURITYGROUPLIMITEXCEEDED = "LimitExceeded.ModuleSecurityGroupLimitExceeded"

	// 弹性网卡或公网IP数超过CPU规定的限额。
	LIMITEXCEEDED_NICORIPLIMITEXCEEDED = "LimitExceeded.NicOrIPLimitExceeded"

	// 内网IP超出限制。
	LIMITEXCEEDED_PRIVATEIPQUOTALIMITEXCEEDED = "LimitExceeded.PrivateIPQuotaLimitExceeded"

	// 安全组绑定的实例数超限。
	LIMITEXCEEDED_SECURITYGROUPINSTANCELIMITEXCEEDED = "LimitExceeded.SecurityGroupInstanceLimitExceeded"

	// 安全组可关联模块数目限制。
	LIMITEXCEEDED_SECURITYGROUPMODULELIMITEXCEEDED = "LimitExceeded.SecurityGroupModuleLimitExceeded"

	// 安全组规则数量超过上限。
	LIMITEXCEEDED_SECURITYGROUPPOLICYSET = "LimitExceeded.SecurityGroupPolicySet"

	// 申请的cpu核数超限。
	LIMITEXCEEDED_VCPULIMITEXCEEDED = "LimitExceeded.VcpuLimitExceeded"

	// 缺少参数错误。
	MISSINGPARAMETER = "MissingParameter"

	// 缺少需关联的实体参数。
	MISSINGPARAMETER_MISSINGASSOCIATEENTITY = "MissingParameter.MissingAssociateEntity"

	// 获取基础配置请求参数不全。
	MISSINGPARAMETER_MISSINGBASECONFIGPARAMETER = "MissingParameter.MissingBaseConfigParameter"

	// 镜像操作参数不全。
	MISSINGPARAMETER_MISSINGIMAGEPARAMETER = "MissingParameter.MissingImageParameter"

	// 机型操作参数不全。
	MISSINGPARAMETER_MISSINGINSTANCETYPECONFIGPARAMETER = "MissingParameter.MissingInstanceTypeConfigParameter"

	// 实例操作参数不全。
	MISSINGPARAMETER_MISSINGINSTANCESPARAMETER = "MissingParameter.MissingInstancesParameter"

	// 模块操作参数不全。
	MISSINGPARAMETER_MISSINGMODULEPARAMETER = "MissingParameter.MissingModuleParameter"

	// 缺少网卡操作配置的请求参数。
	MISSINGPARAMETER_MISSINGNETWORKINTERFACEPARAMETER = "MissingParameter.MissingNetworkInterfaceParameter"

	// 节点操作参数不全。
	MISSINGPARAMETER_MISSINGNODEPARAMETER = "MissingParameter.MissingNodeParameter"

	// 获取概览页配置请求参数不全。
	MISSINGPARAMETER_MISSINGOVERVIEWPARAMETER = "MissingParameter.MissingOverViewParameter"

	// 缺少私有IP地址。
	MISSINGPARAMETER_MISSINGPRIVATEIPADDRESS = "MissingParameter.MissingPrivateIpAddress"

	// 不允许执行当前操作。
	OPERATIONDENIED_INSTANCEOPERATIONINPROGRESS = "OperationDenied.InstanceOperationInProgress"

	// 资源被占用。
	RESOURCEINUSE = "ResourceInUse"

	// 资源不足。
	RESOURCEINSUFFICIENT = "ResourceInsufficient"

	// IP资源不足。
	RESOURCEINSUFFICIENT_IPQUOTANOTENOUGH = "ResourceInsufficient.IPQuotaNotEnough"

	// 实例资源不足。
	RESOURCEINSUFFICIENT_INSTANCEQUOTANOTENOUGH = "ResourceInsufficient.InstanceQuotaNotEnough"

	// 私有镜像数量超出限制。
	RESOURCEINSUFFICIENT_INVAILDPRIVATEIMAGENUM = "ResourceInsufficient.InvaildPrivateImageNum"

	// 资源不存在。
	RESOURCENOTFOUND = "ResourceNotFound"

	// 实例不合法。
	RESOURCENOTFOUND_INSTANCENOTEXIST = "ResourceNotFound.InstanceNotExist"

	// ResourceNotFound.NotFound
	RESOURCENOTFOUND_NOTFOUND = "ResourceNotFound.NotFound"

	// ResourceUnavailable.Attached
	RESOURCEUNAVAILABLE_ATTACHED = "ResourceUnavailable.Attached"

	// ResourceUnavailable.Expire
	RESOURCEUNAVAILABLE_EXPIRE = "ResourceUnavailable.Expire"

	// 实例未运行。
	RESOURCEUNAVAILABLE_INSTANCENOTRUNNING = "ResourceUnavailable.InstanceNotRunning"

	// ResourceUnavailable.NotPortable
	RESOURCEUNAVAILABLE_NOTPORTABLE = "ResourceUnavailable.NotPortable"

	// ResourceUnavailable.NotSupported
	RESOURCEUNAVAILABLE_NOTSUPPORTED = "ResourceUnavailable.NotSupported"

	// ResourceUnavailable.RepeatRefund
	RESOURCEUNAVAILABLE_REPEATREFUND = "ResourceUnavailable.RepeatRefund"

	// ResourceUnavailable.TypeError
	RESOURCEUNAVAILABLE_TYPEERROR = "ResourceUnavailable.TypeError"

	// 负载均衡资源已经售罄。
	RESOURCESSOLDOUT_LOADBALANCERSOLDOUT = "ResourcesSoldOut.LoadBalancerSoldOut"

	// 未授权操作。
	UNAUTHORIZEDOPERATION = "UnauthorizedOperation"

	// 没有权限进行该操作。
	UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"

	// UnauthorizedOperation.MFAExpired
	UNAUTHORIZEDOPERATION_MFAEXPIRED = "UnauthorizedOperation.MFAExpired"

	// UnauthorizedOperation.NotCertification
	UNAUTHORIZEDOPERATION_NOTCERTIFICATION = "UnauthorizedOperation.NotCertification"

	// UnauthorizedOperation.NotHavePaymentRight
	UNAUTHORIZEDOPERATION_NOTHAVEPAYMENTRIGHT = "UnauthorizedOperation.NotHavePaymentRight"

	// 无windows镜像权限。
	UNAUTHORIZEDOPERATION_WINDOWSIMAGE = "UnauthorizedOperation.WindowsImage"

	// 操作不支持。
	UNSUPPORTEDOPERATION = "UnsupportedOperation"

	// 未找到相关IP。
	UNSUPPORTEDOPERATION_ADDRESSIDNOTFOUND = "UnsupportedOperation.AddressIdNotFound"

	// 指定实例已经绑定了EIP。需先解绑当前的EIP才能再次进行绑定操作。
	UNSUPPORTEDOPERATION_ALREADYBINDEIP = "UnsupportedOperation.AlreadyBindEip"

	// 弹性网卡与实例是已关联的。
	UNSUPPORTEDOPERATION_ATTACHMENTALREADYEXISTS = "UnsupportedOperation.AttachmentAlreadyExists"

	// 实例是未关联的。
	UNSUPPORTEDOPERATION_ATTACHMENTNOTFOUND = "UnsupportedOperation.AttachmentNotFound"

	// 禁止删除默认路由表。
	UNSUPPORTEDOPERATION_DELDEFAULTROUTE = "UnsupportedOperation.DelDefaultRoute"

	// 禁止删除已关联子网的路由表。
	UNSUPPORTEDOPERATION_DELROUTEWITHSUBNET = "UnsupportedOperation.DelRouteWithSubnet"

	// 安全组规则重复。
	UNSUPPORTEDOPERATION_DUPLICATEPOLICY = "UnsupportedOperation.DuplicatePolicy"

	// 不支持ECMP。
	UNSUPPORTEDOPERATION_ECMP = "UnsupportedOperation.Ecmp"

	// 和云联网的路由形成ECMP。
	UNSUPPORTEDOPERATION_ECMPWITHCCNROUTE = "UnsupportedOperation.EcmpWithCcnRoute"

	// 和用户自定义的路由形成ECMP。
	UNSUPPORTEDOPERATION_ECMPWITHUSERROUTE = "UnsupportedOperation.EcmpWithUserRoute"

	// 无效实例ID。指定的实例ID不存在。
	UNSUPPORTEDOPERATION_INSTANCEIDNOTFOUND = "UnsupportedOperation.InstanceIdNotFound"

	// 不被支持的实例。
	UNSUPPORTEDOPERATION_INSTANCEIDNOTSUPPORTED = "UnsupportedOperation.InstanceIdNotSupported"

	// 当前机型不支持所选镜像。
	UNSUPPORTEDOPERATION_INSTANCETYPENOTSUPPORTIMAGE = "UnsupportedOperation.InstanceTypeNotSupportImage"

	// 当前状态不能进行该操作。
	UNSUPPORTEDOPERATION_INVALIDINSTANCESTATE = "UnsupportedOperation.InvalidInstanceState"

	// 指定 NetworkInterfaceId 不存在或指定的PrivateIpAddress不在NetworkInterfaceId上。
	UNSUPPORTEDOPERATION_INVALIDNETWORKINTERFACEIDNOTFOUND = "UnsupportedOperation.InvalidNetworkInterfaceIdNotFound"

	// 指定弹性网卡的指定内网IP已经绑定了EIP，不能重复绑定。
	UNSUPPORTEDOPERATION_INVALIDPRIVATEIPADDRESSALREADYBINDEIP = "UnsupportedOperation.InvalidPrivateIpAddressAlreadyBindEip"

	// 资源状态不合法。
	UNSUPPORTEDOPERATION_INVALIDSTATE = "UnsupportedOperation.InvalidState"

	// 请确认提供的IP地址是否完整。
	UNSUPPORTEDOPERATION_MALFORMED = "UnsupportedOperation.Malformed"

	// 资源互斥操作任务正在执行。
	UNSUPPORTEDOPERATION_MUTEXOPERATIONTASKRUNNING = "UnsupportedOperation.MutexOperationTaskRunning"

	// 实例规格仅支持绑定三个EIP。
	UNSUPPORTEDOPERATION_QUOTALIMITEXCEEDED = "UnsupportedOperation.QuotaLimitExceeded"

	// UnsupportedOperation.SnapHasShared
	UNSUPPORTEDOPERATION_SNAPHASSHARED = "UnsupportedOperation.SnapHasShared"

	// UnsupportedOperation.SnapshotHasBindedImage
	UNSUPPORTEDOPERATION_SNAPSHOTHASBINDEDIMAGE = "UnsupportedOperation.SnapshotHasBindedImage"

	// 当前状态不能进行此操作。
	UNSUPPORTEDOPERATION_STATUSNOTPERMIT = "UnsupportedOperation.StatusNotPermit"

	// 系统路由，禁止操作。
	UNSUPPORTEDOPERATION_SYSTEMROUTE = "UnsupportedOperation.SystemRoute"

	// 指定安全组规则版本号和当前最新版本不一致。
	UNSUPPORTEDOPERATION_VERSIONMISMATCH = "UnsupportedOperation.VersionMismatch"

	// 资源不属于同一个VPC。
	UNSUPPORTEDOPERATION_VPCMISMATCH = "UnsupportedOperation.VpcMismatch"
)
