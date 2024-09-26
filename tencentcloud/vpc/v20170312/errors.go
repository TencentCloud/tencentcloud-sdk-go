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

package v20170312

const (
	// 此产品的特有错误码

	// 账户配额不足，每个腾讯云账户每个地域下最多可创建 20 个 EIP。
	ADDRESSQUOTALIMITEXCEEDED = "AddressQuotaLimitExceeded"

	// 申购次数不足，每个腾讯云账户每个地域每天申购次数为配额数*2 次。
	ADDRESSQUOTALIMITEXCEEDED_DAILYALLOCATE = "AddressQuotaLimitExceeded.DailyAllocate"

	// CAM签名/鉴权错误。
	AUTHFAILURE = "AuthFailure"

	// 地址没有弹性网卡信息。
	FAILEDOPERATION_ADDRESSENIINFONOTFOUND = "FailedOperation.AddressEniInfoNotFound"

	// 账户余额不足。
	FAILEDOPERATION_BALANCEINSUFFICIENT = "FailedOperation.BalanceInsufficient"

	// 不支持的地域。
	FAILEDOPERATION_INVALIDREGION = "FailedOperation.InvalidRegion"

	// 不支持的IP类型。
	FAILEDOPERATION_IPTYPENOTPERMIT = "FailedOperation.IpTypeNotPermit"

	// 未找到实例的主网卡。
	FAILEDOPERATION_MASTERENINOTFOUND = "FailedOperation.MasterEniNotFound"

	// 网络探测超时，请稍后重试。
	FAILEDOPERATION_NETDETECTTIMEOUT = "FailedOperation.NetDetectTimeOut"

	// 任务执行失败。
	FAILEDOPERATION_TASKFAILED = "FailedOperation.TaskFailed"

	// 内部错误。
	INTERNALERROR = "InternalError"

	// 创建Ckafka路由失败，请稍后重试。
	INTERNALERROR_CREATECKAFKAROUTEERROR = "InternalError.CreateCkafkaRouteError"

	// 内部模块错误。
	INTERNALERROR_MODULEERROR = "InternalError.ModuleError"

	// 操作内部错误。
	INTERNALSERVERERROR = "InternalServerError"

	// 不支持此账户。
	INVALIDACCOUNT_NOTSUPPORTED = "InvalidAccount.NotSupported"

	// 指定EIP处于被封堵状态。当EIP处于封堵状态的时候是不能够进行绑定操作的，需要先进行解封。
	INVALIDADDRESSID_BLOCKED = "InvalidAddressId.Blocked"

	//  指定的EIP不存在。
	INVALIDADDRESSID_NOTFOUND = "InvalidAddressId.NotFound"

	// 指定EIP处于欠费状态。
	INVALIDADDRESSIDSTATE_INARREARS = "InvalidAddressIdState.InArrears"

	// 指定 EIP 当前状态不能进行绑定操作。只有 EIP 的状态是 UNBIND 时才能进行绑定操作。
	INVALIDADDRESSIDSTATUS_NOTPERMIT = "InvalidAddressIdStatus.NotPermit"

	// 指定EIP的当前状态不允许进行该操作。
	INVALIDADDRESSSTATE = "InvalidAddressState"

	// 不被支持的实例。
	INVALIDINSTANCE_NOTSUPPORTED = "InvalidInstance.NotSupported"

	// 指定实例已经绑定了EIP。需先解绑当前的EIP才能再次进行绑定操作。
	INVALIDINSTANCEID_ALREADYBINDEIP = "InvalidInstanceId.AlreadyBindEip"

	// 无效实例ID。指定的实例ID不存在。
	INVALIDINSTANCEID_NOTFOUND = "InvalidInstanceId.NotFound"

	// 指定 NetworkInterfaceId 不存在或指定的PrivateIpAddress不在NetworkInterfaceId上。
	INVALIDNETWORKINTERFACEID_NOTFOUND = "InvalidNetworkInterfaceId.NotFound"

	// 参数错误。
	INVALIDPARAMETER = "InvalidParameter"

	// ACL ID与ACL类型不匹配。
	INVALIDPARAMETER_ACLTYPEMISMATCH = "InvalidParameter.AclTypeMismatch"

	// 参数不支持同时指定。
	INVALIDPARAMETER_COEXIST = "InvalidParameter.Coexist"

	// 指定过滤条件不存在。
	INVALIDPARAMETER_FILTERINVALIDKEY = "InvalidParameter.FilterInvalidKey"

	// 指定过滤条件不是键值对。
	INVALIDPARAMETER_FILTERNOTDICT = "InvalidParameter.FilterNotDict"

	// 指定过滤选项值不是列表。
	INVALIDPARAMETER_FILTERVALUESNOTLIST = "InvalidParameter.FilterValuesNotList"

	// 该过滤规则不合法。
	INVALIDPARAMETER_INVALIDFILTER = "InvalidParameter.InvalidFilter"

	// 指定参数值不是预期的字典格式。
	INVALIDPARAMETER_INVALIDKEY = "InvalidParameter.InvalidKey"

	// 下一跳类型与下一跳网关不匹配。
	INVALIDPARAMETER_NEXTHOPMISMATCH = "InvalidParameter.NextHopMismatch"

	// 指定参数值不是键值对。
	INVALIDPARAMETER_NOTDICT = "InvalidParameter.NotDict"

	// 指定参数值不是一个列表。
	INVALIDPARAMETER_NOTLIST = "InvalidParameter.NotList"

	// 指定键值不是一个字符串型。
	INVALIDPARAMETER_NOTSTR = "InvalidParameter.NotStr"

	// 专线网关跨可用区容灾组不存在。
	INVALIDPARAMETER_VPGHAGROUPNOTFOUND = "InvalidParameter.VpgHaGroupNotFound"

	// 指定的两个参数冲突，不能同时存在。 EIP只能绑定在实例上或指定网卡的指定内网 IP 上。
	INVALIDPARAMETERCONFLICT = "InvalidParameterConflict"

	// 参数取值错误。
	INVALIDPARAMETERVALUE = "InvalidParameterValue"

	// 被攻击的IP地址。
	INVALIDPARAMETERVALUE_ADDRESSATTACKED = "InvalidParameterValue.AddressAttacked"

	// 该地址ID不合法。
	INVALIDPARAMETERVALUE_ADDRESSIDMALFORMED = "InvalidParameterValue.AddressIdMalformed"

	// 该地址计费方式与其他地址冲突。
	INVALIDPARAMETERVALUE_ADDRESSINTERNETCHARGETYPECONFLICT = "InvalidParameterValue.AddressInternetChargeTypeConflict"

	// 该IP地址现在不可用。
	INVALIDPARAMETERVALUE_ADDRESSIPNOTAVAILABLE = "InvalidParameterValue.AddressIpNotAvailable"

	// IP地址未找到。
	INVALIDPARAMETERVALUE_ADDRESSIPNOTFOUND = "InvalidParameterValue.AddressIpNotFound"

	// VPC中不存在此IP地址。
	INVALIDPARAMETERVALUE_ADDRESSIPNOTINVPC = "InvalidParameterValue.AddressIpNotInVpc"

	// 此IPv6地址未发布。
	INVALIDPARAMETERVALUE_ADDRESSIPNOTPUBLIC = "InvalidParameterValue.AddressIpNotPublic"

	// 未查询到该地址。
	INVALIDPARAMETERVALUE_ADDRESSIPSNOTFOUND = "InvalidParameterValue.AddressIpsNotFound"

	// 该地址不可与此实例申请。
	INVALIDPARAMETERVALUE_ADDRESSNOTAPPLICABLE = "InvalidParameterValue.AddressNotApplicable"

	// 该地址不是CalcIP。
	INVALIDPARAMETERVALUE_ADDRESSNOTCALCIP = "InvalidParameterValue.AddressNotCalcIP"

	// 未找到该地址。
	INVALIDPARAMETERVALUE_ADDRESSNOTFOUND = "InvalidParameterValue.AddressNotFound"

	// 该IPv6地址已经发布。
	INVALIDPARAMETERVALUE_ADDRESSPUBLISHED = "InvalidParameterValue.AddressPublished"

	// 当前IP地址类型不正确。
	INVALIDPARAMETERVALUE_ADDRESSTYPECONFLICT = "InvalidParameterValue.AddressTypeConflict"

	// 带宽超出限制。
	INVALIDPARAMETERVALUE_BANDWIDTHOUTOFRANGE = "InvalidParameterValue.BandwidthOutOfRange"

	// 带宽包ID不正确。
	INVALIDPARAMETERVALUE_BANDWIDTHPACKAGEIDMALFORMED = "InvalidParameterValue.BandwidthPackageIdMalformed"

	// 该带宽包正在被使用。
	INVALIDPARAMETERVALUE_BANDWIDTHPACKAGEINUSE = "InvalidParameterValue.BandwidthPackageInUse"

	// 未查询到该带宽包。
	INVALIDPARAMETERVALUE_BANDWIDTHPACKAGENOTFOUND = "InvalidParameterValue.BandwidthPackageNotFound"

	// 选择带宽低于可允许的最小范围。
	INVALIDPARAMETERVALUE_BANDWIDTHTOOSMALL = "InvalidParameterValue.BandwidthTooSmall"

	// 指定云联网关联黑石私有网络数量达到上限。
	INVALIDPARAMETERVALUE_CCNATTACHBMVPCLIMITEXCEEDED = "InvalidParameterValue.CcnAttachBmvpcLimitExceeded"

	// 目的网段不在对端VPC的CIDR范围内。
	INVALIDPARAMETERVALUE_CIDRNOTINPEERVPC = "InvalidParameterValue.CidrNotInPeerVpc"

	// 指定CIDR不在SSL-VPN所属私有网络CIDR内。
	INVALIDPARAMETERVALUE_CIDRNOTINSSLVPNVPC = "InvalidParameterValue.CidrNotInSslVpnVpc"

	// 非法入参组合。
	INVALIDPARAMETERVALUE_COMBINATION = "InvalidParameterValue.Combination"

	// 入参重复。
	INVALIDPARAMETERVALUE_DUPLICATE = "InvalidParameterValue.Duplicate"

	// 参数值存在重复。
	INVALIDPARAMETERVALUE_DUPLICATEPARA = "InvalidParameterValue.DuplicatePara"

	// 本端地域和端地域重复。
	INVALIDPARAMETERVALUE_DUPLICATEREGION = "InvalidParameterValue.DuplicateRegion"

	// 值超过上限。
	INVALIDPARAMETERVALUE_EIPBRANDWIDTHOUTINVALID = "InvalidParameterValue.EIPBrandWidthOutInvalid"

	// 缺少参数。
	INVALIDPARAMETERVALUE_EMPTY = "InvalidParameterValue.Empty"

	// IPv6转换实例ID已经存在。
	INVALIDPARAMETERVALUE_IPV6RULEIDEXISTED = "InvalidParameterValue.IPv6RuleIdExisted"

	// IPv6规则没有更改。
	INVALIDPARAMETERVALUE_IPV6RULENOTCHANGE = "InvalidParameterValue.IPv6RuleNotChange"

	// 资源格式错误
	INVALIDPARAMETERVALUE_ILLEGAL = "InvalidParameterValue.Illegal"

	// 不能在VPC CIDR范围内。
	INVALIDPARAMETERVALUE_INVPCCIDR = "InvalidParameterValue.InVpcCidr"

	// 该实例的计费方式与其他实例不同。
	INVALIDPARAMETERVALUE_INCONSISTENTINSTANCEINTERNETCHARGETYPE = "InvalidParameterValue.InconsistentInstanceInternetChargeType"

	// 该实例不支持AnycastEIP。
	INVALIDPARAMETERVALUE_INSTANCEDOESNOTSUPPORTANYCAST = "InvalidParameterValue.InstanceDoesNotSupportAnycast"

	// 实例不存在公网IP。
	INVALIDPARAMETERVALUE_INSTANCEHASNOWANIP = "InvalidParameterValue.InstanceHasNoWanIP"

	// 该实例已有WanIP。
	INVALIDPARAMETERVALUE_INSTANCEHASWANIP = "InvalidParameterValue.InstanceHasWanIP"

	// 实例ID错误。
	INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"

	// 该实例没有CalcIP，无法完成请求。
	INVALIDPARAMETERVALUE_INSTANCENOCALCIP = "InvalidParameterValue.InstanceNoCalcIP"

	// 该实例没有WanIP，无法完成请求。
	INVALIDPARAMETERVALUE_INSTANCENOWANIP = "InvalidParameterValue.InstanceNoWanIP"

	// 由于该IP被禁用，无法绑定该实例。
	INVALIDPARAMETERVALUE_INSTANCENORMALPUBLICIPBLOCKED = "InvalidParameterValue.InstanceNormalPublicIpBlocked"

	// 弹性网卡绑定的实例与地址绑定的实例不一致。
	INVALIDPARAMETERVALUE_INSTANCENOTMATCHASSOCIATEENI = "InvalidParameterValue.InstanceNotMatchAssociateEni"

	// 网络计费模式没有更改。
	INVALIDPARAMETERVALUE_INTERNETCHARGETYPENOTCHANGED = "InvalidParameterValue.InternetChargeTypeNotChanged"

	// 无效的带宽包计费方式。
	INVALIDPARAMETERVALUE_INVALIDBANDWIDTHPACKAGECHARGETYPE = "InvalidParameterValue.InvalidBandwidthPackageChargeType"

	// 参数的值不存在或不支持。
	INVALIDPARAMETERVALUE_INVALIDBUSINESS = "InvalidParameterValue.InvalidBusiness"

	// 传入的DedicatedClusterId有误。
	INVALIDPARAMETERVALUE_INVALIDDEDICATEDCLUSTERID = "InvalidParameterValue.InvalidDedicatedClusterId"

	// 该IP只能绑定小时流量后付费和带宽包实例。
	INVALIDPARAMETERVALUE_INVALIDINSTANCEINTERNETCHARGETYPE = "InvalidParameterValue.InvalidInstanceInternetChargeType"

	// 该实例状态无法完成操作。
	INVALIDPARAMETERVALUE_INVALIDINSTANCESTATE = "InvalidParameterValue.InvalidInstanceState"

	// 无效的IPv6地址。
	INVALIDPARAMETERVALUE_INVALIDIPV6 = "InvalidParameterValue.InvalidIpv6"

	// 该Tag不合法。
	INVALIDPARAMETERVALUE_INVALIDTAG = "InvalidParameterValue.InvalidTag"

	// 未查询到该IPv6规则。
	INVALIDPARAMETERVALUE_IP6RULENOTFOUND = "InvalidParameterValue.Ip6RuleNotFound"

	// 未查询到该IPv6翻译器。
	INVALIDPARAMETERVALUE_IP6TRANSLATORNOTFOUND = "InvalidParameterValue.Ip6TranslatorNotFound"

	// 负载均衡实例已经绑定了EIP。
	INVALIDPARAMETERVALUE_LBALREADYBINDEIP = "InvalidParameterValue.LBAlreadyBindEip"

	// 参数值超出限制。
	INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"

	// 入参格式不合法。
	INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"

	// 指定审批单号和资源不匹配。
	INVALIDPARAMETERVALUE_MEMBERAPPROVALAPPLICATIONIDMISMATCH = "InvalidParameterValue.MemberApprovalApplicationIdMismatch"

	// 流程服务审批单未审批。
	INVALIDPARAMETERVALUE_MEMBERAPPROVALAPPLICATIONNOTAPPROVED = "InvalidParameterValue.MemberApprovalApplicationNotApproved"

	// 流程服务审批单被拒绝。
	INVALIDPARAMETERVALUE_MEMBERAPPROVALAPPLICATIONREJECTED = "InvalidParameterValue.MemberApprovalApplicationRejected"

	// 该请求需要走BPAAS流程服务审批，当前发起审批中。
	INVALIDPARAMETERVALUE_MEMBERAPPROVALAPPLICATIONSTARTED = "InvalidParameterValue.MemberApprovalApplicationStarted"

	// 缺少绑定的实例。
	INVALIDPARAMETERVALUE_MISSINGASSOCIATEENTITY = "InvalidParameterValue.MissingAssociateEntity"

	// 集群类型不同的IP不可在同一请求中。
	INVALIDPARAMETERVALUE_MIXEDADDRESSIPSETTYPE = "InvalidParameterValue.MixedAddressIpSetType"

	// NAT网关的DNAT转换规则已存在。
	INVALIDPARAMETERVALUE_NATGATEWAYDNATRULEEXISTED = "InvalidParameterValue.NatGatewayDnatRuleExisted"

	// NAT网关的DNAT转换规则不存在。
	INVALIDPARAMETERVALUE_NATGATEWAYDNATRULENOTEXISTS = "InvalidParameterValue.NatGatewayDnatRuleNotExists"

	// DNAT转换规则的内网IP需为虚拟机上网卡所用的IP。
	INVALIDPARAMETERVALUE_NATGATEWAYDNATRULEPIPNEEDVM = "InvalidParameterValue.NatGatewayDnatRulePipNeedVm"

	// 新增NAT网关的DNAT转换规则已重复。
	INVALIDPARAMETERVALUE_NATGATEWAYDNATRULEREPEATED = "InvalidParameterValue.NatGatewayDnatRuleRepeated"

	// NAT网关的SNAT转换规则不存在。
	INVALIDPARAMETERVALUE_NATGATEWAYSNATRULENOTEXISTS = "InvalidParameterValue.NatGatewaySnatRuleNotExists"

	// NAT网关的SNAT规则已经存在。
	INVALIDPARAMETERVALUE_NATSNATRULEEXISTS = "InvalidParameterValue.NatSnatRuleExists"

	// 探测目的IP和网络探测在同一个VPC内。
	INVALIDPARAMETERVALUE_NETDETECTINVPC = "InvalidParameterValue.NetDetectInVpc"

	// 探测目的IP在云联网的路由表中找不到匹配的下一跳。
	INVALIDPARAMETERVALUE_NETDETECTNOTFOUNDIP = "InvalidParameterValue.NetDetectNotFoundIp"

	// 探测目的IP与同一个私有网络内的同一个子网下的其他网络探测的探测目的IP相同。
	INVALIDPARAMETERVALUE_NETDETECTSAMEIP = "InvalidParameterValue.NetDetectSameIp"

	// 网络接口ID不正确。
	INVALIDPARAMETERVALUE_NETWORKINTERFACEIDMALFORMED = "InvalidParameterValue.NetworkInterfaceIdMalformed"

	// 未找到网络接口ID，或私有IP地址未在网络接口配置。
	INVALIDPARAMETERVALUE_NETWORKINTERFACENOTFOUND = "InvalidParameterValue.NetworkInterfaceNotFound"

	// 必须在VPC CIDR范围内。
	INVALIDPARAMETERVALUE_NOTINVPCCIDR = "InvalidParameterValue.NotInVpcCidr"

	// 不是UTF8编码。
	INVALIDPARAMETERVALUE_NOTUTF8ENCODINGERROR = "InvalidParameterValue.NotUtf8EncodingError"

	// 该操作仅对主网卡支持。
	INVALIDPARAMETERVALUE_ONLYSUPPORTEDFORMASTERNETWORKCARD = "InvalidParameterValue.OnlySupportedForMasterNetworkCard"

	// 参数值格式不匹配。
	INVALIDPARAMETERVALUE_PARAMETERMISMATCH = "InvalidParameterValue.ParameterMismatch"

	// 私网NAT网关不存在
	INVALIDPARAMETERVALUE_PRIVATENATNOTEXISTS = "InvalidParameterValue.PrivateNatNotExists"

	// 私网NAT网关规则不存在。
	INVALIDPARAMETERVALUE_PRIVATENATRULENOTEXISTS = "InvalidParameterValue.PrivateNatRuleNotExists"

	// 私网NAT网关传入规则类型不支持。
	INVALIDPARAMETERVALUE_PRIVATENATSNATRULENOTSUPPORT = "InvalidParameterValue.PrivateNatSnatRuleNotSupport"

	// 参数值不在指定范围。
	INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"

	// 参数值是一个系统保留对象。
	INVALIDPARAMETERVALUE_RESERVED = "InvalidParameterValue.Reserved"

	// 该资源已加入其他带宽包。
	INVALIDPARAMETERVALUE_RESOURCEALREADYEXISTED = "InvalidParameterValue.ResourceAlreadyExisted"

	// 该资源已过期。
	INVALIDPARAMETERVALUE_RESOURCEEXPIRED = "InvalidParameterValue.ResourceExpired"

	// 资源ID不正确。
	INVALIDPARAMETERVALUE_RESOURCEIDMALFORMED = "InvalidParameterValue.ResourceIdMalformed"

	// 该资源不在此带宽包中。
	INVALIDPARAMETERVALUE_RESOURCENOTEXISTED = "InvalidParameterValue.ResourceNotExisted"

	// 未查询到该资源。
	INVALIDPARAMETERVALUE_RESOURCENOTFOUND = "InvalidParameterValue.ResourceNotFound"

	// 该资源不支持此操作。
	INVALIDPARAMETERVALUE_RESOURCENOTSUPPORT = "InvalidParameterValue.ResourceNotSupport"

	// 指定的优先级之间冲突或与已存在的优先级冲突。
	INVALIDPARAMETERVALUE_ROUTEPOLICYPRIORITYCONFLICT = "InvalidParameterValue.RoutePolicyPriorityConflict"

	// SSL-VPN-SERVER 云端网段和SSL-VPN-SERVER 客户端网段重叠。
	INVALIDPARAMETERVALUE_SSLCCNVPNSERVERCIDRCONFLICT = "InvalidParameterValue.SslCcnVpnServerCidrConflict"

	// 存在关机的主机还在使用当前资源，无法操作。
	INVALIDPARAMETERVALUE_STOPCHARGINGINSTANCEINUSE = "InvalidParameterValue.StopChargingInstanceInUse"

	// 子网CIDR冲突。
	INVALIDPARAMETERVALUE_SUBNETCONFLICT = "InvalidParameterValue.SubnetConflict"

	// CIDR与同一个私有网络内的另一个子网发生重叠。
	INVALIDPARAMETERVALUE_SUBNETOVERLAP = "InvalidParameterValue.SubnetOverlap"

	// 子网与辅助Cidr网段重叠。
	INVALIDPARAMETERVALUE_SUBNETOVERLAPASSISTCIDR = "InvalidParameterValue.SubnetOverlapAssistCidr"

	// 子网CIDR不合法。
	INVALIDPARAMETERVALUE_SUBNETRANGE = "InvalidParameterValue.SubnetRange"

	// 标签键重复。
	INVALIDPARAMETERVALUE_TAGDUPLICATEKEY = "InvalidParameterValue.TagDuplicateKey"

	// 重复的标签资源类型。
	INVALIDPARAMETERVALUE_TAGDUPLICATERESOURCETYPE = "InvalidParameterValue.TagDuplicateResourceType"

	// 标签键无效。
	INVALIDPARAMETERVALUE_TAGINVALIDKEY = "InvalidParameterValue.TagInvalidKey"

	// 标签键长度无效。
	INVALIDPARAMETERVALUE_TAGINVALIDKEYLEN = "InvalidParameterValue.TagInvalidKeyLen"

	// 标签值无效。
	INVALIDPARAMETERVALUE_TAGINVALIDVAL = "InvalidParameterValue.TagInvalidVal"

	// 标签键不存在。
	INVALIDPARAMETERVALUE_TAGKEYNOTEXISTS = "InvalidParameterValue.TagKeyNotExists"

	// 标签没有分配配额。
	INVALIDPARAMETERVALUE_TAGNOTALLOCATEDQUOTA = "InvalidParameterValue.TagNotAllocatedQuota"

	// 该标签不存在。
	INVALIDPARAMETERVALUE_TAGNOTEXISTED = "InvalidParameterValue.TagNotExisted"

	// 不支持的标签。
	INVALIDPARAMETERVALUE_TAGNOTSUPPORTTAG = "InvalidParameterValue.TagNotSupportTag"

	// 标签资源格式错误。
	INVALIDPARAMETERVALUE_TAGRESOURCEFORMATERROR = "InvalidParameterValue.TagResourceFormatError"

	// 标签时间戳超配。
	INVALIDPARAMETERVALUE_TAGTIMESTAMPEXCEEDED = "InvalidParameterValue.TagTimestampExceeded"

	// 标签值不存在。
	INVALIDPARAMETERVALUE_TAGVALNOTEXISTS = "InvalidParameterValue.TagValNotExists"

	// 无效参数值。参数值太长。
	INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"

	// 流量包ID格式错误。
	INVALIDPARAMETERVALUE_TRAFFICPACKAGEID = "InvalidParameterValue.TrafficPackageId"

	// 该流量包ID不合法。
	INVALIDPARAMETERVALUE_TRAFFICPACKAGEIDMALFORMED = "InvalidParameterValue.TrafficPackageIdMalformed"

	// 未查询到此流量包。
	INVALIDPARAMETERVALUE_TRAFFICPACKAGENOTFOUND = "InvalidParameterValue.TrafficPackageNotFound"

	// 指定的流量包不支持此操作
	INVALIDPARAMETERVALUE_TRAFFICPACKAGENOTSUPPORTED = "InvalidParameterValue.TrafficPackageNotSupported"

	// 该可用区不可用。
	INVALIDPARAMETERVALUE_UNAVAILABLEZONE = "InvalidParameterValue.UnavailableZone"

	// 目的网段和当前VPC的CIDR冲突。
	INVALIDPARAMETERVALUE_VPCCIDRCONFLICT = "InvalidParameterValue.VpcCidrConflict"

	// 当前功能不支持此专线网关。
	INVALIDPARAMETERVALUE_VPGTYPENOTMATCH = "InvalidParameterValue.VpgTypeNotMatch"

	// 当前BGP隧道网段：`%(key)s`和已有BGP隧道网段：`%(value)s`重叠。
	INVALIDPARAMETERVALUE_VPNCONNBGPTUNNELCIDRCONFLICT = "InvalidParameterValue.VpnConnBgpTunnelCidrConflict"

	// BGP隧道网段必须为掩码为30的网段。
	INVALIDPARAMETERVALUE_VPNCONNBGPTUNNELCIDRMASK = "InvalidParameterValue.VpnConnBgpTunnelCidrMask"

	// 云端或用户端BGP地址：`%(value)s`必须在BGP隧道网段`%(key)s`内。
	INVALIDPARAMETERVALUE_VPNCONNBGPTUNNELCIDRNOTSUPPORTED = "InvalidParameterValue.VpnConnBgpTunnelCidrNotSupported"

	// 目的网段和当前VPN通道的CIDR冲突。
	INVALIDPARAMETERVALUE_VPNCONNCIDRCONFLICT = "InvalidParameterValue.VpnConnCidrConflict"

	// VPN通道探测ip冲突。
	INVALIDPARAMETERVALUE_VPNCONNHEALTHCHECKIPCONFLICT = "InvalidParameterValue.VpnConnHealthCheckIpConflict"

	// 参数Zone的值与CDC所在Zone冲突。
	INVALIDPARAMETERVALUE_ZONECONFLICT = "InvalidParameterValue.ZoneConflict"

	// 指定弹性网卡的指定内网IP已经绑定了EIP，不能重复绑定。
	INVALIDPRIVATEIPADDRESS_ALREADYBINDEIP = "InvalidPrivateIpAddress.AlreadyBindEip"

	// 无效的路由策略ID（RouteId）。
	INVALIDROUTEID_NOTFOUND = "InvalidRouteId.NotFound"

	// 无效的路由表,路由表实例ID不合法。
	INVALIDROUTETABLEID_MALFORMED = "InvalidRouteTableId.Malformed"

	// 无效的路由表,路由表资源不存在，请再次核实您输入的资源信息是否正确。
	INVALIDROUTETABLEID_NOTFOUND = "InvalidRouteTableId.NotFound"

	// 无效的安全组,安全组实例ID不合法。
	INVALIDSECURITYGROUPID_MALFORMED = "InvalidSecurityGroupID.Malformed"

	// 无效的安全组,安全组实例ID不存在。
	INVALIDSECURITYGROUPID_NOTFOUND = "InvalidSecurityGroupID.NotFound"

	// 无效的VPC,VPC实例ID不合法。
	INVALIDVPCID_MALFORMED = "InvalidVpcId.Malformed"

	// 无效的VPC,VPC资源不存在。
	INVALIDVPCID_NOTFOUND = "InvalidVpcId.NotFound"

	// 无效的VPN网关,VPN实例ID不合法。
	INVALIDVPNGATEWAYID_MALFORMED = "InvalidVpnGatewayId.Malformed"

	// 无效的VPN网关,VPN实例不存在，请再次核实您输入的资源信息是否正确。
	INVALIDVPNGATEWAYID_NOTFOUND = "InvalidVpnGatewayId.NotFound"

	// 超过配额限制。
	LIMITEXCEEDED = "LimitExceeded"

	// 账号退还配额超过限制。
	LIMITEXCEEDED_ACCOUNTRETURNQUOTA = "LimitExceeded.AccountReturnQuota"

	// 接口请求次数超过限频。
	LIMITEXCEEDED_ACTIONLIMITED = "LimitExceeded.ActionLimited"

	// 分配IP地址数量达到上限。
	LIMITEXCEEDED_ADDRESS = "LimitExceeded.Address"

	// 租户申请的弹性IP超过上限。
	LIMITEXCEEDED_ADDRESSQUOTALIMITEXCEEDED = "LimitExceeded.AddressQuotaLimitExceeded"

	// 实例关联快照策略数量达到上限。
	LIMITEXCEEDED_ATTACHEDSNAPSHOTPOLICYEXCEEDED = "LimitExceeded.AttachedSnapshotPolicyExceeded"

	// 带宽包配额超过限制。
	LIMITEXCEEDED_BANDWIDTHPACKAGEQUOTA = "LimitExceeded.BandwidthPackageQuota"

	// 当前带宽包加入资源上限。
	LIMITEXCEEDED_BANDWIDTHPACKAGERESOURCEQUOTA = "LimitExceeded.BandwidthPackageResourceQuota"

	// 云联网路由传播策略数量已达到上限。
	LIMITEXCEEDED_CCNROUTEBROADCASTPOLICY = "LimitExceeded.CcnRouteBroadcastPolicy"

	// 云联网路由传播策略条件数量已达到上限。
	LIMITEXCEEDED_CCNROUTEBROADCASTPOLICYCOND = "LimitExceeded.CcnRouteBroadcastPolicyCond"

	// 超过更换IP配额。
	LIMITEXCEEDED_CHANGEADDRESSQUOTA = "LimitExceeded.ChangeAddressQuota"

	// VPC分配网段数量达到上限。
	LIMITEXCEEDED_CIDRBLOCK = "LimitExceeded.CidrBlock"

	// 当前实例关联的云联网数量达到上限。
	LIMITEXCEEDED_CURRENTINSTANCEATTACHEDCCNINSTANCES = "LimitExceeded.CurrentInstanceAttachedCcnInstances"

	// 租户每天申请的弹性IP超过上限。
	LIMITEXCEEDED_DAILYALLOCATEADDRESSQUOTALIMITEXCEEDED = "LimitExceeded.DailyAllocateAddressQuotaLimitExceeded"

	// 超过每日更换IP配额。
	LIMITEXCEEDED_DAILYCHANGEADDRESSQUOTA = "LimitExceeded.DailyChangeAddressQuota"

	// 默认VPC数量已达到上限。
	LIMITEXCEEDED_DEFAULTVPCLIMITEXCEEDED = "LimitExceeded.DefaultVpcLimitExceeded"

	// 实例绑定的弹性IP超过配额。
	LIMITEXCEEDED_INSTANCEADDRESSQUOTA = "LimitExceeded.InstanceAddressQuota"

	// 修改地址网络计费模式配额超过限制。
	LIMITEXCEEDED_MODIFYADDRESSINTERNETCHARGETYPEQUOTA = "LimitExceeded.ModifyAddressInternetChargeTypeQuota"

	// 每月地址找回配额超过限制。
	LIMITEXCEEDED_MONTHLYADDRESSRECOVERYQUOTA = "LimitExceeded.MonthlyAddressRecoveryQuota"

	// NAT网关的DNAT转换规则数量已达到上限。
	LIMITEXCEEDED_NATGATEWAYDNATLIMITEXCEEDED = "LimitExceeded.NatGatewayDnatLimitExceeded"

	// NAT网关数量已达到上限。
	LIMITEXCEEDED_NATGATEWAYLIMITEXCEEDED = "LimitExceeded.NatGatewayLimitExceeded"

	// 私有网络创建的NAT网关超过上限。
	LIMITEXCEEDED_NATGATEWAYPERVPCLIMITEXCEEDED = "LimitExceeded.NatGatewayPerVpcLimitExceeded"

	// VPC内创建的弹性网卡数量超过上限
	LIMITEXCEEDED_NETWORKINTERFACELIMITEXCEEDED = "LimitExceeded.NetworkInterfaceLimitExceeded"

	// 过滤参数名称超过限制。
	LIMITEXCEEDED_NUMBEROFFILTERS = "LimitExceeded.NumberOfFilters"

	// 私网NAT网关转换规则的访问控制规则数量达到上限。
	LIMITEXCEEDED_PRIVATENATTRANSLATIONACLRULE = "LimitExceeded.PrivateNatTranslationAclRule"

	// 私网NAT网关的转换数量达到上限。
	LIMITEXCEEDED_PRIVATENATTRANSLATIONNATRULE = "LimitExceeded.PrivateNatTranslationNatRule"

	// NAT网关绑定的弹性IP超过上限。
	LIMITEXCEEDED_PUBLICIPADDRESSPERNATGATEWAYLIMITEXCEEDED = "LimitExceeded.PublicIpAddressPerNatGatewayLimitExceeded"

	// 安全组规则数量超过上限。
	LIMITEXCEEDED_SECURITYGROUPPOLICYSET = "LimitExceeded.SecurityGroupPolicySet"

	// 子网分配子网段数量达到上限。
	LIMITEXCEEDED_SUBNETCIDRBLOCK = "LimitExceeded.SubnetCidrBlock"

	// 标签键已达到上限。
	LIMITEXCEEDED_TAGKEYEXCEEDED = "LimitExceeded.TagKeyExceeded"

	// 每个资源的标签键已达到上限。
	LIMITEXCEEDED_TAGKEYPERRESOURCEEXCEEDED = "LimitExceeded.TagKeyPerResourceExceeded"

	// 没有足够的标签配额。
	LIMITEXCEEDED_TAGNOTENOUGHQUOTA = "LimitExceeded.TagNotEnoughQuota"

	// 标签配额已满，无法创建资源。
	LIMITEXCEEDED_TAGQUOTA = "LimitExceeded.TagQuota"

	// 标签配额已达到上限。
	LIMITEXCEEDED_TAGQUOTAEXCEEDED = "LimitExceeded.TagQuotaExceeded"

	// 标签键的数目已达到上限。
	LIMITEXCEEDED_TAGTAGSEXCEEDED = "LimitExceeded.TagTagsExceeded"

	// 流量包配额超过限制。
	LIMITEXCEEDED_TRAFFICPACKAGEQUOTA = "LimitExceeded.TrafficPackageQuota"

	// 有效的对等个数超过配额上限。
	LIMITEXCEEDED_VPCPEERAVALIMITEXCEEDED = "LimitExceeded.VpcPeerAvaLimitExceeded"

	// 可创建的对等连接个数超过总上限。
	LIMITEXCEEDED_VPCPEERTOTALLIMITEXCEEDED = "LimitExceeded.VpcPeerTotalLimitExceeded"

	// 缺少参数错误。
	MISSINGPARAMETER = "MissingParameter"

	// 缺少多个参数错误
	MISSINGPARAMETER_MULTIMISSINGPARAMETER = "MissingParameter.MultiMissingParameter"

	// 指定公网IP处于隔离状态。
	OPERATIONDENIED_ADDRESSINARREARS = "OperationDenied.AddressInArrears"

	// 互斥的任务正在执行。
	OPERATIONDENIED_MUTEXTASKRUNNING = "OperationDenied.MutexTaskRunning"

	// 资源被占用。
	RESOURCEINUSE = "ResourceInUse"

	// 指定IP地址已经在使用中。
	RESOURCEINUSE_ADDRESS = "ResourceInUse.Address"

	// 资源不足。
	RESOURCEINSUFFICIENT = "ResourceInsufficient"

	// 网段资源不足。
	RESOURCEINSUFFICIENT_CIDRBLOCK = "ResourceInsufficient.CidrBlock"

	// 指定实例类型已售罄。
	RESOURCEINSUFFICIENT_INSTANCE = "ResourceInsufficient.Instance"

	// 子网IP资源不足, 无法分配IP。
	RESOURCEINSUFFICIENT_SUBNET = "ResourceInsufficient.Subnet"

	// 资源不存在。
	RESOURCENOTFOUND = "ResourceNotFound"

	// Svc不存在。
	RESOURCENOTFOUND_SVCNOTEXIST = "ResourceNotFound.SvcNotExist"

	// 资源不可用。
	RESOURCEUNAVAILABLE = "ResourceUnavailable"

	// 当前用户不在指定终端节点服务的白名单内。
	RESOURCEUNAVAILABLE_SERVICEWHITELISTNOTADDED = "ResourceUnavailable.ServiceWhiteListNotAdded"

	// 未授权操作。
	UNAUTHORIZEDOPERATION = "UnauthorizedOperation"

	// 无权限申请AnycastEip资源。
	UNAUTHORIZEDOPERATION_ANYCASTEIP = "UnauthorizedOperation.AnycastEip"

	// 绑定关系不存在。
	UNAUTHORIZEDOPERATION_ATTACHMENTNOTFOUND = "UnauthorizedOperation.AttachmentNotFound"

	// 未授权的用户。
	UNAUTHORIZEDOPERATION_INVALIDACCOUNT = "UnauthorizedOperation.InvalidAccount"

	// 账号未实名。
	UNAUTHORIZEDOPERATION_NOREALNAMEAUTHENTICATION = "UnauthorizedOperation.NoRealNameAuthentication"

	// 主IP不支持指定操作。
	UNAUTHORIZEDOPERATION_PRIMARYIP = "UnauthorizedOperation.PrimaryIp"

	// 对等连接本端VPC与对端VPC存在CIDR冲突,或一端与已建立的对等连接某一端冲突。
	UNAUTHORIZEDOPERATION_VPCPEERCIDRCONFLICT = "UnauthorizedOperation.VpcPeerCidrConflict"

	// 未知参数错误。
	UNKNOWNPARAMETER = "UnknownParameter"

	// 参数无法识别，可以尝试相似参数代替。
	UNKNOWNPARAMETER_WITHGUESS = "UnknownParameter.WithGuess"

	// 操作不支持。
	UNSUPPORTEDOPERATION = "UnsupportedOperation"

	// 不支持的账户。
	UNSUPPORTEDOPERATION_ACCOUNTNOTSUPPORTED = "UnsupportedOperation.AccountNotSupported"

	// 接口不存在。
	UNSUPPORTEDOPERATION_ACTIONNOTFOUND = "UnsupportedOperation.ActionNotFound"

	// 欠费状态不支持该操作。
	UNSUPPORTEDOPERATION_ADDRESSIPINARREAR = "UnsupportedOperation.AddressIpInArrear"

	// 此付费模式的IP地址不支持该操作。
	UNSUPPORTEDOPERATION_ADDRESSIPINTERNETCHARGETYPENOTPERMIT = "UnsupportedOperation.AddressIpInternetChargeTypeNotPermit"

	// 绑定此实例的IP地址不支持该操作。
	UNSUPPORTEDOPERATION_ADDRESSIPNOTSUPPORTINSTANCE = "UnsupportedOperation.AddressIpNotSupportInstance"

	// 此IP地址状态不支持该操作。
	UNSUPPORTEDOPERATION_ADDRESSIPSTATUSNOTPERMIT = "UnsupportedOperation.AddressIpStatusNotPermit"

	// 该地址状态不支持此操作。
	UNSUPPORTEDOPERATION_ADDRESSSTATUSNOTPERMIT = "UnsupportedOperation.AddressStatusNotPermit"

	// 资源不在指定的AppId下。
	UNSUPPORTEDOPERATION_APPIDMISMATCH = "UnsupportedOperation.AppIdMismatch"

	// APPId不存在。
	UNSUPPORTEDOPERATION_APPIDNOTFOUND = "UnsupportedOperation.AppIdNotFound"

	// CCN关联的其他vpc已经存在nat的路由
	UNSUPPORTEDOPERATION_ASSOCIATEDVPCOFCCNHADNATROUTE = "UnsupportedOperation.AssociatedVpcOfCcnHadNatRoute"

	// 绑定关系已存在。
	UNSUPPORTEDOPERATION_ATTACHMENTALREADYEXISTS = "UnsupportedOperation.AttachmentAlreadyExists"

	// 绑定关系不存在。
	UNSUPPORTEDOPERATION_ATTACHMENTNOTFOUND = "UnsupportedOperation.AttachmentNotFound"

	// 当前云联网还有预付费带宽未到期，不支持主动删除。
	UNSUPPORTEDOPERATION_BANDWIDTHNOTEXPIRED = "UnsupportedOperation.BandwidthNotExpired"

	// 该带宽包不支持此操作。
	UNSUPPORTEDOPERATION_BANDWIDTHPACKAGEIDNOTSUPPORTED = "UnsupportedOperation.BandwidthPackageIdNotSupported"

	// 已绑定EIP。
	UNSUPPORTEDOPERATION_BINDEIP = "UnsupportedOperation.BindEIP"

	// 指定VPC CIDR范围不支持私有网络和基础网络设备互通。
	UNSUPPORTEDOPERATION_CIDRUNSUPPORTEDCLASSICLINK = "UnsupportedOperation.CIDRUnSupportedClassicLink"

	// 路由表选择策略不支持默认路由表。
	UNSUPPORTEDOPERATION_CANNOTASSIGNDEFAULTROUTETABLE = "UnsupportedOperation.CanNotAssignDefaultRouteTable"

	// 实例已关联CCN。
	UNSUPPORTEDOPERATION_CCNATTACHED = "UnsupportedOperation.CcnAttached"

	// 云联网实例不支持跨账号关联。
	UNSUPPORTEDOPERATION_CCNCROSSACCOUNT = "UnsupportedOperation.CcnCrossAccount"

	// 当前云联网有流日志，不支持删除。
	UNSUPPORTEDOPERATION_CCNHASFLOWLOG = "UnsupportedOperation.CcnHasFlowLog"

	// CCN实例所属账号未通过联通审批。
	UNSUPPORTEDOPERATION_CCNINSTANCEACCOUNTNOTAPPROVEDBYUNICOM = "UnsupportedOperation.CcnInstanceAccountNotApprovedByUnicom"

	// 实例未关联CCN。
	UNSUPPORTEDOPERATION_CCNNOTATTACHED = "UnsupportedOperation.CcnNotAttached"

	// 当前云联网未开启路由传播策略。
	UNSUPPORTEDOPERATION_CCNNOTENABLEBROADCASTPOLICY = "UnsupportedOperation.CcnNotEnableBroadcastPolicy"

	// 跨账号场景下不支持自驾云账号实例 关联普通账号云联网。
	UNSUPPORTEDOPERATION_CCNORDINARYACCOUNTREFUSEATTACH = "UnsupportedOperation.CcnOrdinaryAccountRefuseAttach"

	// 指定的路由表不存在。
	UNSUPPORTEDOPERATION_CCNROUTETABLENOTEXIST = "UnsupportedOperation.CcnRouteTableNotExist"

	// CDC子网不支持创建非本地网关类型的路由。
	UNSUPPORTEDOPERATION_CDCSUBNETNOTSUPPORTUNLOCALGATEWAY = "UnsupportedOperation.CdcSubnetNotSupportUnLocalGateway"

	// 实例已经和VPC绑定。
	UNSUPPORTEDOPERATION_CLASSICINSTANCEIDALREADYEXISTS = "UnsupportedOperation.ClassicInstanceIdAlreadyExists"

	// 负载均衡的安全组规则已达到上限。
	UNSUPPORTEDOPERATION_CLBPOLICYEXCEEDLIMIT = "UnsupportedOperation.ClbPolicyExceedLimit"

	// 公网Clb不支持该规则。
	UNSUPPORTEDOPERATION_CLBPOLICYLIMIT = "UnsupportedOperation.ClbPolicyLimit"

	// 与该VPC下的TKE容器的网段重叠。
	UNSUPPORTEDOPERATION_CONFLICTWITHDOCKERROUTE = "UnsupportedOperation.ConflictWithDockerRoute"

	// 当前账号非联通账号。
	UNSUPPORTEDOPERATION_CURRENTACCOUNTISNOTUNICOMACCOUNT = "UnsupportedOperation.CurrentAccountIsNotUnicomAccount"

	// 当前查询地域非跨境。
	UNSUPPORTEDOPERATION_CURRENTQUERYREGIONISNOTCROSSBORDER = "UnsupportedOperation.CurrentQueryRegionIsNotCrossBorder"

	// 该专线网关存在关联的NAT规则，不允许删除，请先删调所有的NAT规则。
	UNSUPPORTEDOPERATION_DCGATEWAYNATRULEEXISTS = "UnsupportedOperation.DCGatewayNatRuleExists"

	// 指定的VPC未发现专线网关。
	UNSUPPORTEDOPERATION_DCGATEWAYSNOTFOUNDINVPC = "UnsupportedOperation.DcGatewaysNotFoundInVpc"

	// 禁止删除默认路由表。
	UNSUPPORTEDOPERATION_DELDEFAULTROUTE = "UnsupportedOperation.DelDefaultRoute"

	// 禁止删除已关联子网的路由表。
	UNSUPPORTEDOPERATION_DELROUTEWITHSUBNET = "UnsupportedOperation.DelRouteWithSubnet"

	// VPN通道状态为更新中/销毁中/创建中，不支持此操作。
	UNSUPPORTEDOPERATION_DELETEVPNCONNINVALIDSTATE = "UnsupportedOperation.DeleteVpnConnInvalidState"

	// 发货失败。
	UNSUPPORTEDOPERATION_DELIVERYFAILED = "UnsupportedOperation.DeliveryFailed"

	// 专线网关正在更新BGP Community属性。
	UNSUPPORTEDOPERATION_DIRECTCONNECTGATEWAYISUPDATINGCOMMUNITY = "UnsupportedOperation.DirectConnectGatewayIsUpdatingCommunity"

	// 指定的路由策略已发布至云联网，请先撤销。
	UNSUPPORTEDOPERATION_DISABLEDNOTIFYCCN = "UnsupportedOperation.DisabledNotifyCcn"

	// 创建DPDK NAT流日志时，采集类型只支持全部。
	UNSUPPORTEDOPERATION_DPDKNATFLOWLOGONLYSUPPORTALLTRAFFICTYPE = "UnsupportedOperation.DpdkNatFlowLogOnlySupportAllTrafficType"

	// 安全组规则重复。
	UNSUPPORTEDOPERATION_DUPLICATEPOLICY = "UnsupportedOperation.DuplicatePolicy"

	// 不支持ECMP。
	UNSUPPORTEDOPERATION_ECMP = "UnsupportedOperation.Ecmp"

	// 和云联网的路由形成ECMP。
	UNSUPPORTEDOPERATION_ECMPWITHCCNROUTE = "UnsupportedOperation.EcmpWithCcnRoute"

	// 和用户自定义的路由形成ECMP。
	UNSUPPORTEDOPERATION_ECMPWITHUSERROUTE = "UnsupportedOperation.EcmpWithUserRoute"

	// 当前地域不支持启用组播。
	UNSUPPORTEDOPERATION_ENABLEMULTICAST = "UnsupportedOperation.EnableMulticast"

	// 终端节点和终端节点服务的CdcId不一致。
	UNSUPPORTEDOPERATION_ENDPOINTMISMATCHENDPOINTSERVICECDCID = "UnsupportedOperation.EndPointMismatchEndPointServiceCdcId"

	// 终端节点服务本身不能是终端节点。
	UNSUPPORTEDOPERATION_ENDPOINTSERVICE = "UnsupportedOperation.EndPointService"

	// 指定ResourceId对应的流日志已经创建
	UNSUPPORTEDOPERATION_FLOWLOGINSTANCEEXISTED = "UnsupportedOperation.FlowLogInstanceExisted"

	// 不支持创建流日志：当前弹性网卡绑定的是KO机型。
	UNSUPPORTEDOPERATION_FLOWLOGSNOTSUPPORTKOINSTANCEENI = "UnsupportedOperation.FlowLogsNotSupportKoInstanceEni"

	// 不支持创建流日志：当前弹性网卡未绑定实例。
	UNSUPPORTEDOPERATION_FLOWLOGSNOTSUPPORTNULLINSTANCEENI = "UnsupportedOperation.FlowLogsNotSupportNullInstanceEni"

	// TGW还没有投放IPv6网段
	UNSUPPORTEDOPERATION_IPV6CIDRNOTDEPLOYED = "UnsupportedOperation.IPV6CidrNotDeployed"

	// 指定的客户令牌已经被使用。
	UNSUPPORTEDOPERATION_IDEMPOTENTPARAMETERMISMATCH = "UnsupportedOperation.IdempotentParameterMismatch"

	// 先前的幂等请求仍在处理中，请稍后重试。
	UNSUPPORTEDOPERATION_IDEMPOTENTPROCESSING = "UnsupportedOperation.IdempotentProcessing"

	// 该种类型地址不支持此操作。
	UNSUPPORTEDOPERATION_INCORRECTADDRESSRESOURCETYPE = "UnsupportedOperation.IncorrectAddressResourceType"

	// 用户配置的实例和路由表不匹配。
	UNSUPPORTEDOPERATION_INSTANCEANDRTBNOTMATCH = "UnsupportedOperation.InstanceAndRtbNotMatch"

	// 当前云联网`%(value)s`的CdcId与传入实例的CdcId不一致，不支持关联。
	UNSUPPORTEDOPERATION_INSTANCECDCIDNOTMATCHCCNCDCID = "UnsupportedOperation.InstanceCdcIdNotMatchCcnCdcId"

	// 指定实例资源不匹配。
	UNSUPPORTEDOPERATION_INSTANCEMISMATCH = "UnsupportedOperation.InstanceMismatch"

	// 路由表选择策略配置的实例不存在。
	UNSUPPORTEDOPERATION_INSTANCENOTEXIST = "UnsupportedOperation.InstanceNotExist"

	// 跨账号场景下不支持普通账号实例关联自驾云账号云联网。
	UNSUPPORTEDOPERATION_INSTANCEORDINARYACCOUNTREFUSEATTACH = "UnsupportedOperation.InstanceOrdinaryAccountRefuseAttach"

	// 该地址绑定的实例状态不支持此操作。
	UNSUPPORTEDOPERATION_INSTANCESTATENOTSUPPORTED = "UnsupportedOperation.InstanceStateNotSupported"

	// 账户余额不足。
	UNSUPPORTEDOPERATION_INSUFFICIENTFUNDS = "UnsupportedOperation.InsufficientFunds"

	// 至少需要两个运营商
	UNSUPPORTEDOPERATION_INSUFFICIENTINTERNETSERVICEPROVIDERS = "UnsupportedOperation.InsufficientInternetServiceProviders"

	// 不支持该操作。
	UNSUPPORTEDOPERATION_INVALIDACTION = "UnsupportedOperation.InvalidAction"

	// 该地址的网络付费方式不支持此操作。
	UNSUPPORTEDOPERATION_INVALIDADDRESSINTERNETCHARGETYPE = "UnsupportedOperation.InvalidAddressInternetChargeType"

	// 该地址状态不支持此操作。
	UNSUPPORTEDOPERATION_INVALIDADDRESSSTATE = "UnsupportedOperation.InvalidAddressState"

	// 无效的实例状态。
	UNSUPPORTEDOPERATION_INVALIDINSTANCESTATE = "UnsupportedOperation.InvalidInstanceState"

	// 该计费方式不支持此操作。
	UNSUPPORTEDOPERATION_INVALIDRESOURCEINTERNETCHARGETYPE = "UnsupportedOperation.InvalidResourceInternetChargeType"

	// 不支持加入此协议的带宽包。
	UNSUPPORTEDOPERATION_INVALIDRESOURCEPROTOCOL = "UnsupportedOperation.InvalidResourceProtocol"

	// 资源状态不合法。
	UNSUPPORTEDOPERATION_INVALIDSTATE = "UnsupportedOperation.InvalidState"

	// 当前状态不支持发布至云联网，请重试。
	UNSUPPORTEDOPERATION_INVALIDSTATUSNOTIFYCCN = "UnsupportedOperation.InvalidStatusNotifyCcn"

	// 关联当前云联网的实例的账号存在不是金融云账号。
	UNSUPPORTEDOPERATION_ISNOTFINANCEACCOUNT = "UnsupportedOperation.IsNotFinanceAccount"

	// 该ISP不支持此操作。
	UNSUPPORTEDOPERATION_ISPNOTSUPPORTED = "UnsupportedOperation.IspNotSupported"

	// 需要先解绑 IDC通道
	UNSUPPORTEDOPERATION_LDCXDISASSOCIATE = "UnsupportedOperation.LDCXDisassociate"

	// 指定的CDC已存在本地网关。
	UNSUPPORTEDOPERATION_LOCALGATEWAYALREADYEXISTS = "UnsupportedOperation.LocalGatewayAlreadyExists"

	// 资源被锁定。
	UNSUPPORTEDOPERATION_LOCKEDRESOURCES = "UnsupportedOperation.LockedResources"

	// 账户不支持修改公网IP的该属性。
	UNSUPPORTEDOPERATION_MODIFYADDRESSATTRIBUTE = "UnsupportedOperation.ModifyAddressAttribute"

	// VPC实例内部有账号纬度的IPv6白名单，不支持关联多云联网。
	UNSUPPORTEDOPERATION_MULTIPLEVPCNOTSUPPORTATTACHACCOUNTHASIPV6 = "UnsupportedOperation.MultipleVpcNotSupportAttachAccountHasIpv6"

	// 资源互斥操作任务正在执行。
	UNSUPPORTEDOPERATION_MUTEXOPERATIONTASKRUNNING = "UnsupportedOperation.MutexOperationTaskRunning"

	// NAT网关的公网IP不存在。
	UNSUPPORTEDOPERATION_NATGATEWAYEIPNOTEXISTS = "UnsupportedOperation.NatGatewayEipNotExists"

	// NAT网关存在未解绑的IP。
	UNSUPPORTEDOPERATION_NATGATEWAYHADEIPUNASSOCIATE = "UnsupportedOperation.NatGatewayHadEipUnassociate"

	// NAT网关已被封禁，不支持此操作。
	UNSUPPORTEDOPERATION_NATGATEWAYRESTRICTED = "UnsupportedOperation.NatGatewayRestricted"

	// SNAT/DNAT转换规则所指定的内网IP已绑定了其他的规则，无法重复绑定。
	UNSUPPORTEDOPERATION_NATGATEWAYRULEPIPEXISTS = "UnsupportedOperation.NatGatewayRulePipExists"

	// SNAT转换规则的内网IP需为虚拟机上网卡所用的IP。
	UNSUPPORTEDOPERATION_NATGATEWAYSNATPIPNEEDVM = "UnsupportedOperation.NatGatewaySnatPipNeedVm"

	// NAT网关的SNAT转换规则不存在。
	UNSUPPORTEDOPERATION_NATGATEWAYSNATRULENOTEXISTS = "UnsupportedOperation.NatGatewaySnatRuleNotExists"

	// NAT网关类型不支持SNAT规则。
	UNSUPPORTEDOPERATION_NATGATEWAYTYPENOTSUPPORTSNAT = "UnsupportedOperation.NatGatewayTypeNotSupportSNAT"

	// NAT实例不支持该操作。
	UNSUPPORTEDOPERATION_NATNOTSUPPORTED = "UnsupportedOperation.NatNotSupported"

	// 指定的子网不支持创建本地网关类型的路由。
	UNSUPPORTEDOPERATION_NORMALSUBNETNOTSUPPORTLOCALGATEWAY = "UnsupportedOperation.NormalSubnetNotSupportLocalGateway"

	// 当前实例已被封禁，无法进行此操作。
	UNSUPPORTEDOPERATION_NOTLOCKEDINSTANCEOPERATION = "UnsupportedOperation.NotLockedInstanceOperation"

	// 目的端的服务在IP申请中使用的实例ID和这里传入的不匹配。
	UNSUPPORTEDOPERATION_NOTMATCHTARGETSERVICE = "UnsupportedOperation.NotMatchTargetService"

	// 当前云联网实例未处于申请中状态，无法进行操作。
	UNSUPPORTEDOPERATION_NOTPENDINGCCNINSTANCE = "UnsupportedOperation.NotPendingCcnInstance"

	// 当前云联网为非后付费类型，无法进行此操作。
	UNSUPPORTEDOPERATION_NOTPOSTPAIDCCNOPERATION = "UnsupportedOperation.NotPostpaidCcnOperation"

	// 当前云联网不支持同时关联EDGE实例和跨境实例
	UNSUPPORTEDOPERATION_NOTSUPPORTATTACHEDGEANDCROSSBORDERINSTANCE = "UnsupportedOperation.NotSupportAttachEdgeAndCrossBorderInstance"

	// 没有开启多路由表特性，不能创建用户路由表。
	UNSUPPORTEDOPERATION_NOTSUPPORTCREATECCNROUTETABLE = "UnsupportedOperation.NotSupportCreateCcnRouteTable"

	// 默认路由表，不支持删除。
	UNSUPPORTEDOPERATION_NOTSUPPORTDELETEDEFAULTCCNROUTETABLE = "UnsupportedOperation.NotSupportDeleteDefaultCcnRouteTable"

	// 不支持删除默认路由表。
	UNSUPPORTEDOPERATION_NOTSUPPORTDELETEDEFAULTROUTETABLE = "UnsupportedOperation.NotSupportDeleteDefaultRouteTable"

	// 公有云到黑石的对等连接不支持删除。
	UNSUPPORTEDOPERATION_NOTSUPPORTDELETEVPCBMPEER = "UnsupportedOperation.NotSupportDeleteVpcBmPeer"

	// 默认路由表，不支持修改。
	UNSUPPORTEDOPERATION_NOTSUPPORTMODIFYDEFAULTCCNROUTETABLE = "UnsupportedOperation.NotSupportModifyDefaultCcnRouteTable"

	// 不支持编辑相同的实例和源地址。
	UNSUPPORTEDOPERATION_NOTSUPPORTSAMECCNINSTANCEANDSOURCEADDRESS = "UnsupportedOperation.NotSupportSameCcnInstanceAndSourceAddress"

	// 不支持的可用区
	UNSUPPORTEDOPERATION_NOTSUPPORTZONE = "UnsupportedOperation.NotSupportZone"

	// 该地址类型不支持释放操作。
	UNSUPPORTEDOPERATION_NOTSUPPORTEDADDRESSIPSCHARGETYPE = "UnsupportedOperation.NotSupportedAddressIpsChargeType"

	// 此地域没有上线出口二资源，请到北京/广州/南京购买。
	UNSUPPORTEDOPERATION_NOTSUPPORTEDPURCHASECENTEREGRESSRESOURCE = "UnsupportedOperation.NotSupportedPurchaseCenterEgressResource"

	// 当前云联网不支持更新路由发布。
	UNSUPPORTEDOPERATION_NOTSUPPORTEDUPDATECCNROUTEPUBLISH = "UnsupportedOperation.NotSupportedUpdateCcnRoutePublish"

	// 指定的路由策略不支持发布或撤销至云联网。
	UNSUPPORTEDOPERATION_NOTIFYCCN = "UnsupportedOperation.NotifyCcn"

	// 此产品计费方式已下线，请尝试其他计费方式。
	UNSUPPORTEDOPERATION_OFFLINECHARGETYPE = "UnsupportedOperation.OfflineChargeType"

	// 仅支持专业版Ckafka。
	UNSUPPORTEDOPERATION_ONLYSUPPORTPROFESSIONKAFKA = "UnsupportedOperation.OnlySupportProfessionKafka"

	// 预付费云联网只支持地域间限速。
	UNSUPPORTEDOPERATION_PREPAIDCCNONLYSUPPORTINTERREGIONLIMIT = "UnsupportedOperation.PrepaidCcnOnlySupportInterRegionLimit"

	// 指定的值是主IP。
	UNSUPPORTEDOPERATION_PRIMARYIP = "UnsupportedOperation.PrimaryIp"

	// 私网NAT网关存在关联规则。
	UNSUPPORTEDOPERATION_PRIVATENATGATEWAYASSOCIATIONEXISTS = "UnsupportedOperation.PrivateNatGatewayAssociationExists"

	// Nat网关至少存在一个弹性IP，弹性IP不能解绑。
	UNSUPPORTEDOPERATION_PUBLICIPADDRESSDISASSOCIATE = "UnsupportedOperation.PublicIpAddressDisassociate"

	// 绑定NAT网关的弹性IP不是BGP性质的IP。
	UNSUPPORTEDOPERATION_PUBLICIPADDRESSISNOTBGPIP = "UnsupportedOperation.PublicIpAddressIsNotBGPIp"

	// 绑定NAT网关的弹性IP不存在。
	UNSUPPORTEDOPERATION_PUBLICIPADDRESSISNOTEXISTED = "UnsupportedOperation.PublicIpAddressIsNotExisted"

	// 绑定NAT网关的弹性IP不是按流量计费的。
	UNSUPPORTEDOPERATION_PUBLICIPADDRESSNOTBILLEDBYTRAFFIC = "UnsupportedOperation.PublicIpAddressNotBilledByTraffic"

	// 当前账号不能在该地域使用产品。
	UNSUPPORTEDOPERATION_PURCHASELIMIT = "UnsupportedOperation.PurchaseLimit"

	// 记录已存在。
	UNSUPPORTEDOPERATION_RECORDEXISTS = "UnsupportedOperation.RecordExists"

	// 记录不存在。
	UNSUPPORTEDOPERATION_RECORDNOTEXISTS = "UnsupportedOperation.RecordNotExists"

	// 资源处于不可用状态，禁止操作。
	UNSUPPORTEDOPERATION_RESOURCEISINVALIDSTATE = "UnsupportedOperation.ResourceIsInvalidState"

	// 输入的资源ID与IP绑定的资源不匹配，请检查。
	UNSUPPORTEDOPERATION_RESOURCEMISMATCH = "UnsupportedOperation.ResourceMismatch"

	// 未找到相关角色，请确认角色是否授权。
	UNSUPPORTEDOPERATION_ROLENOTFOUND = "UnsupportedOperation.RoleNotFound"

	// 当前路由表删除失败，请先检查是否存在关联的策略。
	UNSUPPORTEDOPERATION_ROUTETABLECANNOTDELETE = "UnsupportedOperation.RouteTableCanNotDelete"

	// 当前云联网的路由表数量已超过限制。
	UNSUPPORTEDOPERATION_ROUTETABLEEXCEEDPERVBCLIMIT = "UnsupportedOperation.RouteTableExceedPerVbcLimit"

	// 路由表绑定了子网。
	UNSUPPORTEDOPERATION_ROUTETABLEHASSUBNETRULE = "UnsupportedOperation.RouteTableHasSubnetRule"

	// 当前云联网的路由表选择策略数量已超过限制。
	UNSUPPORTEDOPERATION_ROUTETABLESELECTPOLICYEXCEED = "UnsupportedOperation.RouteTableSelectPolicyExceed"

	// SslVpnClientIds：`vpnc-20f9b3d7` 证书状态已禁用或Client证书状态不可用，不支持禁用证书。
	UNSUPPORTEDOPERATION_SSLCLIENTCERTALREADYDISABLEORCERTABNORMAL = "UnsupportedOperation.SSLClientCertAlreadyDisableOrCertAbnormal"

	// SslVpnClientIds：`vpnc-20f9b3d7` 证书状态已启用或Client证书状态不可用，不支持启用证书。
	UNSUPPORTEDOPERATION_SSLCLIENTCERTALREADYENABLEORCERTABNORMAL = "UnsupportedOperation.SSLClientCertAlreadyEnableOrCertAbnormal"

	// SSL客户端状态不可用，不支持下载
	UNSUPPORTEDOPERATION_SSLCLIENTCERTDISABLEUNSUPPORTEDDOWNLOADSSLCLIENTCERT = "UnsupportedOperation.SSLClientCertDisableUnsupportedDownloadSSLClientCert"

	// 实例已关联快照策略。
	UNSUPPORTEDOPERATION_SNAPSHOTATTACHED = "UnsupportedOperation.SnapshotAttached"

	// 快照备份策略不支持修改。
	UNSUPPORTEDOPERATION_SNAPSHOTBACKUPTYPEMODIFY = "UnsupportedOperation.SnapshotBackupTypeModify"

	// 快照文件生成失败。
	UNSUPPORTEDOPERATION_SNAPSHOTFILEFAILED = "UnsupportedOperation.SnapshotFileFailed"

	// 快照文件已过期或删除。
	UNSUPPORTEDOPERATION_SNAPSHOTFILENOEXIST = "UnsupportedOperation.SnapshotFileNoExist"

	// 快照文件正在生成中，请稍后查看。
	UNSUPPORTEDOPERATION_SNAPSHOTFILEPROCESSING = "UnsupportedOperation.SnapshotFileProcessing"

	// 一次仅支持关联一个地域的实例。
	UNSUPPORTEDOPERATION_SNAPSHOTINSTANCEREGIONDIFF = "UnsupportedOperation.SnapshotInstanceRegionDiff"

	// 实例未关联快照策略。
	UNSUPPORTEDOPERATION_SNAPSHOTNOTATTACHED = "UnsupportedOperation.SnapshotNotAttached"

	// SNAT子网 不支持分配IP。
	UNSUPPORTEDOPERATION_SNATSUBNET = "UnsupportedOperation.SnatSubnet"

	// 指定的终端节点服务所创建的终端节点不支持绑定安全组。
	UNSUPPORTEDOPERATION_SPECIALENDPOINTSERVICE = "UnsupportedOperation.SpecialEndPointService"

	// SslVpnClientId 不存在。
	UNSUPPORTEDOPERATION_SSLVPNCLIENTIDNOTFOUND = "UnsupportedOperation.SslVpnClientIdNotFound"

	// 中继网卡不支持该操作。
	UNSUPPORTEDOPERATION_SUBENINOTSUPPORTTRUNKING = "UnsupportedOperation.SubEniNotSupportTrunking"

	// 子网不存在。
	UNSUPPORTEDOPERATION_SUBNETNOTEXISTS = "UnsupportedOperation.SubnetNotExists"

	// 系统路由，禁止操作。
	UNSUPPORTEDOPERATION_SYSTEMROUTE = "UnsupportedOperation.SystemRoute"

	// 标签正在分配中。
	UNSUPPORTEDOPERATION_TAGALLOCATE = "UnsupportedOperation.TagAllocate"

	// 标签正在释放中。
	UNSUPPORTEDOPERATION_TAGFREE = "UnsupportedOperation.TagFree"

	// 标签没有权限。
	UNSUPPORTEDOPERATION_TAGNOTPERMIT = "UnsupportedOperation.TagNotPermit"

	// 不支持使用系统预留的标签键。
	UNSUPPORTEDOPERATION_TAGSYSTEMRESERVEDTAGKEY = "UnsupportedOperation.TagSystemReservedTagKey"

	// 账号ID不存在。
	UNSUPPORTEDOPERATION_UINNOTFOUND = "UnsupportedOperation.UinNotFound"

	// 不支持跨境。
	UNSUPPORTEDOPERATION_UNABLECROSSBORDER = "UnsupportedOperation.UnableCrossBorder"

	// 当前云联网无法关联金融云实例。
	UNSUPPORTEDOPERATION_UNABLECROSSFINANCE = "UnsupportedOperation.UnableCrossFinance"

	// 未分配IPv6网段。
	UNSUPPORTEDOPERATION_UNASSIGNCIDRBLOCK = "UnsupportedOperation.UnassignCidrBlock"

	// 未绑定EIP。
	UNSUPPORTEDOPERATION_UNBINDEIP = "UnsupportedOperation.UnbindEIP"

	// 账户还有未支付订单，请先完成付款。
	UNSUPPORTEDOPERATION_UNPAIDORDERALREADYEXISTS = "UnsupportedOperation.UnpaidOrderAlreadyExists"

	// 不支持绑定LocalZone弹性公网IP。
	UNSUPPORTEDOPERATION_UNSUPPORTEDBINDLOCALZONEEIP = "UnsupportedOperation.UnsupportedBindLocalZoneEIP"

	// 指定机型不支持弹性网卡。
	UNSUPPORTEDOPERATION_UNSUPPORTEDINSTANCEFAMILY = "UnsupportedOperation.UnsupportedInstanceFamily"

	// 暂无法在此国家/地区提供该服务。
	UNSUPPORTEDOPERATION_UNSUPPORTEDREGION = "UnsupportedOperation.UnsupportedRegion"

	// 当前用户付费类型不支持创建所选付费类型的云联网。
	UNSUPPORTEDOPERATION_USERANDCCNCHARGETYPENOTMATCH = "UnsupportedOperation.UserAndCcnChargeTypeNotMatch"

	// 指定安全组规则版本号和当前最新版本不一致。
	UNSUPPORTEDOPERATION_VERSIONMISMATCH = "UnsupportedOperation.VersionMismatch"

	// 资源不属于同一个VPC。
	UNSUPPORTEDOPERATION_VPCMISMATCH = "UnsupportedOperation.VpcMismatch"

	// 对等连接已存在。
	UNSUPPORTEDOPERATION_VPCPEERALREADYEXIST = "UnsupportedOperation.VpcPeerAlreadyExist"

	// VPC网段存在CIDR冲突。
	UNSUPPORTEDOPERATION_VPCPEERCIDRCONFLICT = "UnsupportedOperation.VpcPeerCidrConflict"

	// 对等连接状态错误。
	UNSUPPORTEDOPERATION_VPCPEERINVALIDSTATECHANGE = "UnsupportedOperation.VpcPeerInvalidStateChange"

	// 该账不能发起操作。
	UNSUPPORTEDOPERATION_VPCPEERPURVIEWERROR = "UnsupportedOperation.VpcPeerPurviewError"

	// 当前通道为非可用状态，不支持该操作。
	UNSUPPORTEDOPERATION_VPNCONNINVALIDSTATE = "UnsupportedOperation.VpnConnInvalidState"

	// SPD本端网段冲突，请检查后重试。
	UNSUPPORTEDOPERATION_VPNCONNSPDOVERLAP = "UnsupportedOperation.VpnConnSPDOverlap"

	// VPC类型VPN网关必须携带VpcId。
	UNSUPPORTEDOPERATION_VPNGWVPCIDMUSTHAVE = "UnsupportedOperation.VpnGwVpcIdMustHave"

	// VPN不支持BGP
	UNSUPPORTEDOPERATION_VPNUNSUPPORTEDBGP = "UnsupportedOperation.VpnUnsupportedBgp"

	// 对端网关BGP ASN和已有的通道对端或云上VPN的BGP ASN相同。
	UNSUPPORTEDOPERATION_VPNUNSUPPORTEDBGPASNEQUAL = "UnsupportedOperation.VpnUnsupportedBgpAsnEqual"

	// `%(value)s`，不支持降低VPN带宽。
	UNSUPPORTEDOPERATION_VPNUNSUPPORTEDMODIFYBANDWIDTH = "UnsupportedOperation.VpnUnsupportedModifyBandwidth"

	// VPN不支持修改BGP ASN。
	UNSUPPORTEDOPERATION_VPNUNSUPPORTEDMODIFYBGPASN = "UnsupportedOperation.VpnUnsupportedModifyBgpAsn"

	// VPN未配置BGP ASN。
	UNSUPPORTEDOPERATION_VPNUNSUPPORTEDNOTEXISTBGPASN = "UnsupportedOperation.VpnUnsupportedNotExistBgpAsn"

	// 指定资源在不同的可用区。
	UNSUPPORTEDOPERATION_ZONEMISMATCH = "UnsupportedOperation.ZoneMismatch"

	// 已经达到指定区域vpc资源申请数量上限。
	VPCLIMITEXCEEDED = "VpcLimitExceeded"
)
