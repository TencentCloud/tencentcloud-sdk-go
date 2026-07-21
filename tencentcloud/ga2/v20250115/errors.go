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

package v20250115

const (
	// 此产品的特有错误码

	// 创建全球加速实例失败，请重试你的请求，如果问题仍然存在，请联系腾讯云客服。
	FAILEDOPERATION_CREATEGLOBALACCELERATORFAILED = "FailedOperation.CreateGlobalAcceleratorFailed"

	// 创建该端口段所需资源不足。
	FAILEDOPERATION_INSUFFICIENTRESOURCES = "FailedOperation.InsufficientResources"

	// 七层监听器不支持携带参数`%(value)s` 。
	INVALIDPARAMETER_APPLICATIONLAYERLISTENERCANNOTCARRYPARAMETERS = "InvalidParameter.ApplicationLayerListenerCannotCarryParameters"

	// 会话保持关闭，不支持携带此参数。
	INVALIDPARAMETER_CLIENTAFFINITYCLOSE = "InvalidParameter.ClientAffinityCLose"

	// 终端节点组配置为关闭健康检查，禁止携带参数 `%(parameter)s`。
	INVALIDPARAMETER_DISABLEHEALTHCHECKNOTCARRYPARAMETERS = "InvalidParameter.DisableHealthCheckNotCarryParameters"

	// 终端节点组配置为关闭健康检查，禁止携带参数 `%(parameter)s`。
	INVALIDPARAMETER_ENABLEHEALTHCHECKNOTCARRYPARAMETERS = "InvalidParameter.EnableHealthCheckNotCarryParameters"

	// 当配置HTTP检查时需要携带参数`%(value)s` 。
	INVALIDPARAMETER_ENDPOINTGROUPCHECKTYPEHTTPCARRYPARAMETERS = "InvalidParameter.EndpointGroupCheckTypeHttpCarryParameters"

	// Http监听器不支持携带参数`%(value)s` 。
	INVALIDPARAMETER_HTTPLISTENERCANNOTCARRYPARAMETERS = "InvalidParameter.HttpListenerCannotCarryParameters"

	// Https监听器不支持携带参数`%(value)s` 。
	INVALIDPARAMETER_HTTPSLISTENERCANNOTCARRYPARAMETERS = "InvalidParameter.HttpsListenerCannotCarryParameters"

	// 仅HTTPS回源协议支持携带参数`%(value)s`。
	INVALIDPARAMETER_HTTPSORIGINSERVERPROTOCOLSUPPORTSPARAMETERS = "InvalidParameter.HttpsOriginServerProtocolSupportsParameters"

	// 请求中传入参数 `%(key)s` 与存量数据存在重复。
	INVALIDPARAMETER_INPUTDUPLICATEWITHEXISTINGDATA = "InvalidParameter.InputDuplicateWithExistingData"

	// 请求中传入参数必选有 `%(value)s` 其中一个。
	INVALIDPARAMETER_INPUTMUSTEXISTONE = "InvalidParameter.InputMustExistOne"

	// 请求中传入参数 `%(key)s` 必选再 `%(value)s` 范围内。
	INVALIDPARAMETER_INPUTOUTOFRANGE = "InvalidParameter.InputOutOfRange"

	// 请求中传入参数 `%(key)s` 存在重复值。
	INVALIDPARAMETER_INPUTREPEAT = "InvalidParameter.InputRepeat"

	// 仅当配置HTTP检查时需要携带参数`%(value)s` 。
	INVALIDPARAMETER_ONLYENDPOINTGROUPCHECKTYPEHTTPCARRYPARAMETERS = "InvalidParameter.OnlyEndpointGroupCheckTypeHttpCarryParameters"

	// 仅支持的单个端口。
	INVALIDPARAMETER_SINGLEPORT = "InvalidParameter.SinglePort"

	// Tcp终端节点组不支持携带参数`%(value)s` 。
	INVALIDPARAMETER_TCPENDPOINTGROUPCANNOTCARRYPARAMETERS = "InvalidParameter.TcpEndpointGroupCannotCarryParameters"

	// 仅Tcp终端节点组支持携带参数`%(value)s` 。
	INVALIDPARAMETER_TCPENDPOINTGROUPCARRYPARAMETERS = "InvalidParameter.TcpEndpointGroupCarryParameters"

	// 四层终端节点组不支持携带参数`%(value)s` 。
	INVALIDPARAMETER_TRANSPORTLAYERENDPOINTGROUPCANNOTCARRYPARAMETERS = "InvalidParameter.TransportLayerEndpointGroupCannotCarryParameters"

	// 四层监听器不支持携带参数`%(value)s` 。
	INVALIDPARAMETER_TRANSPORTLAYERLISTENERCANNOTCARRYPARAMETERS = "InvalidParameter.TransportLayerListenerCannotCarryParameters"

	// Udp终端节点组不支持携带参数`%(value)s` 。
	INVALIDPARAMETER_UDPENDPOINTGROUPCANNOTCARRYPARAMETERS = "InvalidParameter.UdpEndpointGroupCannotCarryParameters"

	// 参数取值错误。
	INVALIDPARAMETERVALUE = "InvalidParameterValue"

	// 只能是可打印ASCII字符。
	INVALIDPARAMETERVALUE_ASCIICHARACTERS = "InvalidParameterValue.AsciiCharacters"

	// 证书 `%(value)s` 和存量证书重复。
	INVALIDPARAMETERVALUE_CERTIFICATESCONFLICT = "InvalidParameterValue.CertificatesConflict"

	// 参数 `%(parameter)s` 不能为空。
	INVALIDPARAMETERVALUE_EMPTY = "InvalidParameterValue.Empty"

	// 七层转发策略域名和存量重复。
	INVALIDPARAMETERVALUE_FORWARDINGPOLICYHOSTCONFLICT = "InvalidParameterValue.ForwardingPolicyHostConflict"

	// 七层转发规则Path和存量重复。
	INVALIDPARAMETERVALUE_FORWARDINGRULEPATHCONFLICT = "InvalidParameterValue.ForwardingRulePathConflict"

	// 实例对象 `%(value)s` 关系不匹配。
	INVALIDPARAMETERVALUE_INSTANCEMISMATCH = "InvalidParameterValue.InstanceMismatch"

	// 实例名称仅支持以大小写字符或中文开头，支持数字、英文句号、或段划线、下划线。
	INVALIDPARAMETERVALUE_INSTANCENAME = "InvalidParameterValue.InstanceName"

	// 参数 `%(parameter)s` 值 `%(value)s` 数量超过限制。不能大于 `%(limit)s` 个。
	INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"

	// 终端节点组监听端口要和终端节点组所属监听器端口保持一致。
	INVALIDPARAMETERVALUE_LISTENERPORTNOTEQUAL = "InvalidParameterValue.ListenerPortNotEqual"

	// 参数`%(value)s`日志主题所在地域与终端节点地域`%(region)s`未正确对应。
	INVALIDPARAMETERVALUE_LOGTASKLOCATEDERROR = "InvalidParameterValue.LogTaskLocatedError"

	// 日志集主题`%(value)s`不存在
	INVALIDPARAMETERVALUE_LOGSETNOTEXIST = "InvalidParameterValue.LogsetNotExist"

	// 参数 `%(parameter)s` 值 `%(value)s` 是无效的。正确且完整的值形如 `%(template)s`。
	INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"

	// 当前不支持一个监听器端口映射到多个回源端口。
	INVALIDPARAMETERVALUE_NOTMAPPINGMANYPORT = "InvalidParameterValue.NotMappingManyPort"

	// 参数 `%(value)s` 不在公网网段内。
	INVALIDPARAMETERVALUE_NOTWITHINPUBLICNETWORK = "InvalidParameterValue.NotWithinPublicNetwork"

	// 参数 `%(parameter)s` 值 `%(value)s` 必须在 `%(value_range)s` 范围内。
	INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"

	// FromPort `%(fromPort)s` 应小于等于ToPort `%(toPort)s`。
	INVALIDPARAMETERVALUE_SEGMENTPORTRANGE = "InvalidParameterValue.SegmentPortRange"

	// 参数值 `%(value)s` 不能含有特殊字符 `%(key)s`。
	INVALIDPARAMETERVALUE_SPECIALCHARACTERS = "InvalidParameterValue.SpecialCharacters"

	// 参数值 `%(value)s` 仅能含有字符 `%(key)s`。
	INVALIDPARAMETERVALUE_SPECIFICCHARACTERS = "InvalidParameterValue.SpecificCharacters"

	// 异步任务ID不存在。
	INVALIDPARAMETERVALUE_TASKNOTFOUND = "InvalidParameterValue.TaskNotFound"

	// 当前TCP协议系列监听器端口已和存量监听器端口重复。
	INVALIDPARAMETERVALUE_TCPSERIESLISTENERPORTEQUAL = "InvalidParameterValue.TcpSeriesListenerPortEqual"

	// 当前全球加速实例已配置第三方节点加速地域，监听器端口不允许一致。
	INVALIDPARAMETERVALUE_THIRDAREALISTENERPORTEQUAL = "InvalidParameterValue.ThirdAreaListenerPortEqual"

	// 当前全球加速实例第三方节点不支持配置监听器端口段。
	INVALIDPARAMETERVALUE_THIRDAREALISTENERPORTSEGMENT = "InvalidParameterValue.ThirdAreaListenerPortSegment"

	// 参数值 `%(value)s` 长度不能大于 `%(max_size)s`。
	INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"

	// 当前UDP协议系列监听器端口已和存量监听器端口重复。
	INVALIDPARAMETERVALUE_UDPSERIESLISTENERPORTEQUAL = "InvalidParameterValue.UdpSeriesListenerPortEqual"

	// 超过配额限制。
	LIMITEXCEEDED = "LimitExceeded"

	// 缺少参数错误。
	MISSINGPARAMETER = "MissingParameter"

	// 当前为七层终端节点组，请求缺少必填参数 `%(parameter)s`。
	MISSINGPARAMETER_APPLICATIONLAYERENDPOINTGROUPPARAMETER = "MissingParameter.ApplicationLayerEndpointGroupParameter"

	// 终端节点组配置为自定义检查时，请求缺少必填参数 `%(parameter)s`。
	MISSINGPARAMETER_CUSTOMCHECKTYPEPARAMETER = "MissingParameter.CustomCheckTypeParameter"

	// 终端节点组配置为开启健康检查，请求缺少必填参数 `%(parameter)s`。
	MISSINGPARAMETER_ENABLEHEALTHCHECKPARAMETER = "MissingParameter.EnableHealthCheckParameter"

	// 资源不存在。
	RESOURCENOTFOUND = "ResourceNotFound"

	// 加速地域`%(value)s` 已经存在。
	UNSUPPORTEDOPERATION_ACCELERATEREGIONREPEAT = "UnsupportedOperation.AccelerateRegionRepeat"

	// 当前账号不支持配置访问日志。
	UNSUPPORTEDOPERATION_ACCESSLOG = "UnsupportedOperation.AccessLog"

	// 账户被冻结，不支持当前操作。
	UNSUPPORTEDOPERATION_ACCOUNTFROZEN = "UnsupportedOperation.AccountFrozen"

	// 已经开启跨境，不支持重复操作。
	UNSUPPORTEDOPERATION_ALREADYENABLECROSSBORDER = "UnsupportedOperation.AlreadyEnableCrossBorder"

	// 七层默认转发规则不允许删除。
	UNSUPPORTEDOPERATION_APPLICATIONLAYERENDPOINTGROUPNOTDELETE = "UnsupportedOperation.ApplicationLayerEndpointGroupNotDelete"

	// 抱歉，您的操作暂时无法完成，请稍后重试或联系客服。
	UNSUPPORTEDOPERATION_BILLINGFAILED = "UnsupportedOperation.BillingFailed"

	// CA证书不允许操作。
	UNSUPPORTEDOPERATION_CACERTIFICATESOPERATE = "UnsupportedOperation.CaCertificatesOperate"

	// 证书状态是 `%(value)s` ，不支持当前操作。
	UNSUPPORTEDOPERATION_CERTIFICATEINVALIDSTATUS = "UnsupportedOperation.CertificateInvalidStatus"

	// 跨境承诺书没签约，不支持当前操作。
	UNSUPPORTEDOPERATION_CROSSBORDERPROMISENOTSIGNED = "UnsupportedOperation.CrossBorderPromiseNotSigned"

	// 当前账号不支持此操作。
	UNSUPPORTEDOPERATION_CURRENTACCOUNTNOTALLOWED = "UnsupportedOperation.CurrentAccountNotAllowed"

	// 默认证书不允许操作。
	UNSUPPORTEDOPERATION_DEFAULTCERTIFICATESOPERATE = "UnsupportedOperation.DefaultCertificatesOperate"

	// 默认终端节点组不支持修改。
	UNSUPPORTEDOPERATION_DEFAULTENDPOINTGROUPMODIFY = "UnsupportedOperation.DefaultEndpointGroupModify"

	// 默认七层转发策略规则不允许操作。
	UNSUPPORTEDOPERATION_DEFAULTFORWARDINGPOLICYOPERATE = "UnsupportedOperation.DefaultForwardingPolicyOperate"

	// 替换证书域名和旧证书域名不匹配。
	UNSUPPORTEDOPERATION_DOMAINMISMATCHED = "UnsupportedOperation.DomainMismatched"

	// 请勿重复操作当前实例状态。
	UNSUPPORTEDOPERATION_DUPLICATEINSTANCESTATUS = "UnsupportedOperation.DuplicateInstanceStatus"

	// 当前用户，不支持设置跨境类型。
	UNSUPPORTEDOPERATION_ENABLECROSSBORDER = "UnsupportedOperation.EnableCrossBorder"

	// 存在加速地域，不支持当前操作。
	UNSUPPORTEDOPERATION_EXISTACCELERATORAREA = "UnsupportedOperation.ExistAcceleratorArea"

	// 存在默认准许所有流量访问通道访问控制策略，不支持当前操作。
	UNSUPPORTEDOPERATION_EXISTACCEPTACLPOLICY = "UnsupportedOperation.ExistAcceptAclPolicy"

	// 存在访问日志，不支持当前操作。
	UNSUPPORTEDOPERATION_EXISTACCESSLOG = "UnsupportedOperation.ExistAccessLog"

	// 配置了策略为拒绝的访问规则，不支持当前操作。
	UNSUPPORTEDOPERATION_EXISTDROPACCEPTACLRULE = "UnsupportedOperation.ExistDropAcceptAclRule"

	// 当前实例存在终端节点组，不支持当前操作。
	UNSUPPORTEDOPERATION_EXISTENDPOINTGROUP = "UnsupportedOperation.ExistEndpointGroup"

	// 存在七层转发规则配置了终端节点组`%(value)s`，不支持当前操作。
	UNSUPPORTEDOPERATION_EXISTFORWARDINGRULE = "UnsupportedOperation.ExistForwardingRule"

	// 存在访问控制策略，不支持当前操作。
	UNSUPPORTEDOPERATION_EXISTGLOBALACCELERATORACLPOLICY = "UnsupportedOperation.ExistGlobalAcceleratorAclPolicy"

	// 存在监听器，不支持当前操作。
	UNSUPPORTEDOPERATION_EXISTLISTENER = "UnsupportedOperation.ExistListener"

	// 存在第三方节点，不允许配置默认准许所有流量访问通道策略。
	UNSUPPORTEDOPERATION_EXISTTHIRDPARTYNODES = "UnsupportedOperation.ExistThirdPartyNodes"

	// 转发规则不支持配置默认终端节点组。
	UNSUPPORTEDOPERATION_FORWARDGROUPFORWARDINGRULEUNSUPPORTEDDEFAULTENDPOINTGROUP = "UnsupportedOperation.ForwardGroupForwardingRuleUnsupportedDefaultEndpointGroup"

	// 监听器和终端节点组的HTTP最大协议版本需要保持一致。
	UNSUPPORTEDOPERATION_HTTPVERSIONINCONSISTENT = "UnsupportedOperation.HttpVersionInconsistent"

	// 实例不是运行状态，不支持当前操作。
	UNSUPPORTEDOPERATION_INSTANCENOTRUNNING = "UnsupportedOperation.InstanceNotRunning"

	// 当前实例状态不允许此操作。
	UNSUPPORTEDOPERATION_INSTANCESTATENOTALLOWEDOPERATE = "UnsupportedOperation.InstanceStateNotAllowedOperate"

	// 账户余额不足。
	UNSUPPORTEDOPERATION_INSUFFICIENTFUNDS = "UnsupportedOperation.InsufficientFunds"

	// 当前全球加速实例未开通内网回源特性。
	UNSUPPORTEDOPERATION_INTERNALENDPOINTFEATURENOTENABLED = "UnsupportedOperation.InternalEndpointFeatureNotEnabled"

	// 不支持IPV6。
	UNSUPPORTEDOPERATION_IPV6NOTSUPPORT = "UnsupportedOperation.Ipv6NotSupport"

	// 当前账号不支持配置端口段。
	UNSUPPORTEDOPERATION_LISTENERPORTSEGMENT = "UnsupportedOperation.ListenerPortSegment"

	// 暂不支持，请先授权开启日志任务。
	UNSUPPORTEDOPERATION_NOTAUTHORIZATION = "UnsupportedOperation.NotAuthorization"

	// 当前全球加速实例没有设置跨境类型。无法创建跨境加速地域或终端节点组。
	UNSUPPORTEDOPERATION_NOTSETCROSSBORDERTYPE = "UnsupportedOperation.NotSetCrossBorderType"

	// 仅TCP监听器，支持通过TOA获取客户端IP。
	UNSUPPORTEDOPERATION_ONLYTCPLISTENERSUPPORTTOA = "UnsupportedOperation.OnlyTcpListenerSupportToa"

	// 请求参数错误，不支持当前操作。具体出错信息为：`%(info)s` 。
	UNSUPPORTEDOPERATION_REQUESTPARAMETERSERROR = "UnsupportedOperation.RequestParametersError"

	// 当前操作使终端节点组域名和全球加速实例域名一致，不支持当前操作。
	UNSUPPORTEDOPERATION_SAMEDOMAIN = "UnsupportedOperation.SameDomain"

	// 当前操作使终端节点组域名和全球加速实例加速地域公网IP一致，不支持当前操作。
	UNSUPPORTEDOPERATION_SAMEPUBLICIP = "UnsupportedOperation.SamePublicIp"

	// 当前全球加速实例未开通TOA特性，请联系腾讯云客服申请。
	UNSUPPORTEDOPERATION_TOAFEATURENOTENABLED = "UnsupportedOperation.TOAFeatureNotEnabled"

	// 不支持创建三网类型的加速地域。
	UNSUPPORTEDOPERATION_THREENETWORKSACCELERATEAREAS = "UnsupportedOperation.ThreeNetworksAccelerateAreas"

	// 不支持创建三网类型的终端节点组。
	UNSUPPORTEDOPERATION_THREENETWORKSENDPOINTGROUP = "UnsupportedOperation.ThreeNetworksEndpointGroup"

	// 四层监听器不支持创建转发策略。
	UNSUPPORTEDOPERATION_TRANSPORTLAYERUNSUPPORTEDOPERATEFORWARDINGPOLICY = "UnsupportedOperation.TransportLayerUnsupportedOperateForwardingPolicy"

	// 不支持创建第三方节点。
	UNSUPPORTEDOPERATION_UNABLECREATETHIRDPARTYNODES = "UnsupportedOperation.UnableCreateThirdPartyNodes"

	// 当前全球加速实例无法创建跨境加速地域或终端节点组。
	UNSUPPORTEDOPERATION_UNABLECROSSBORDER = "UnsupportedOperation.UnableCrossBorder"

	// 虚拟终端节点组不支持Tcp和Udp类型终端节点组。
	UNSUPPORTEDOPERATION_VIRTUALENDPOINTGROUPUNSUPPORTEDTCPANDUDP = "UnsupportedOperation.VirtualEndpointGroupUnsupportedTcpAndUdp"
)
