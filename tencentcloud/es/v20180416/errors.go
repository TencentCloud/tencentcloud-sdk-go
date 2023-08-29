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

package v20180416

const (
	// 此产品的特有错误码

	// CAM签名/鉴权错误。
	AUTHFAILURE = "AuthFailure"

	// 操作未授权。
	AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnAuthorizedOperation"

	// 操作失败。
	FAILEDOPERATION = "FailedOperation"

	// 集群资源配额限制错误。
	FAILEDOPERATION_CLUSTERRESOURCELIMITERROR = "FailedOperation.ClusterResourceLimitError"

	// 节点磁盘块数参数检查失败。
	FAILEDOPERATION_DISKCOUNTPARAMERROR = "FailedOperation.DiskCountParamError"

	// 集群状态错误。
	FAILEDOPERATION_ERRORCLUSTERSTATE = "FailedOperation.ErrorClusterState"

	// 集群索引没有副本存在。
	FAILEDOPERATION_ERRORCLUSTERSTATENOREPLICATION = "FailedOperation.ErrorClusterStateNoReplication"

	// 集群状态不健康。
	FAILEDOPERATION_ERRORCLUSTERSTATEUNHEALTH = "FailedOperation.ErrorClusterStateUnhealth"

	// 获取实例的标签列表错误。
	FAILEDOPERATION_GETTAGINFOERROR = "FailedOperation.GetTagInfoError"

	// 账户未绑定信用卡或paypal，无法支付。
	FAILEDOPERATION_NOPAYMENT = "FailedOperation.NoPayment"

	// 用户未实名认证。
	FAILEDOPERATION_NOTAUTHENTICATED = "FailedOperation.NotAuthenticated"

	// 不支持在滚动重启扩容计算资源同时扩容磁盘数量。
	FAILEDOPERATION_UNSUPPORTEDRESETNODETYPEANDSCALEOUTDISK = "FailedOperation.UnsupportedResetNodeTypeAndScaleOutDisk"

	// 不支持在滚动重启缩容计算资源同时修改磁盘大小
	FAILEDOPERATION_UNSUPPORTEDRESTSCALEDOWNANDMODIFYDISK = "FailedOperation.UnsupportedRestScaleDownAndModifyDisk"

	// 不支持反向调节节点配置和磁盘容量。
	FAILEDOPERATION_UNSUPPORTEDREVERSEREGULATIONNODETYPEANDDISK = "FailedOperation.UnsupportedReverseRegulationNodeTypeAndDisk"

	// 内部错误。
	INTERNALERROR = "InternalError"

	// 参数错误。
	INVALIDPARAMETER = "InvalidParameter"

	// AppId的取值和预期不符。
	INVALIDPARAMETER_INVALIDAPPID = "InvalidParameter.InvalidAppId"

	// 是否自动使用代金券AutoVoucher的取值和预期不符。
	INVALIDPARAMETER_INVALIDAUTOVOUCHER = "InvalidParameter.InvalidAutoVoucher"

	// 集群部署方式InvalidDeployMode的取值和预期不符。
	INVALIDPARAMETER_INVALIDDEPLOYMODE = "InvalidParameter.InvalidDeployMode"

	// 节点磁盘块数DiskCount取值和预期不符。
	INVALIDPARAMETER_INVALIDDISKCOUNT = "InvalidParameter.InvalidDiskCount"

	// 节点磁盘加密信息DiskEncrypt的取值与预期不符。
	INVALIDPARAMETER_INVALIDDISKENCRYPT = "InvalidParameter.InvalidDiskEncrypt"

	// 是否启用增强型ssd云盘DiskEnhance的取值和预期不符。
	INVALIDPARAMETER_INVALIDDISKENHANCE = "InvalidParameter.InvalidDiskEnhance"

	// 实例版本EsVersion的取值和预期不符。
	INVALIDPARAMETER_INVALIDESVERSION = "InvalidParameter.InvalidEsVersion"

	// 无效的InstanceId，没有找到对应资源。
	INVALIDPARAMETER_INVALIDINSTANCEID = "InvalidParameter.InvalidInstanceId"

	// Ip地址取值和预期不符。
	INVALIDPARAMETER_INVALIDIP = "InvalidParameter.InvalidIp"

	// 私有网络vip列表IpList的取值和预期不符。
	INVALIDPARAMETER_INVALIDIPLIST = "InvalidParameter.InvalidIpList"

	// 分页大小Limit的取值和预期不符。
	INVALIDPARAMETER_INVALIDLIMIT = "InvalidParameter.InvalidLimit"

	// 多可用区部署ZoneDetail的信息与预期不符。
	INVALIDPARAMETER_INVALIDMULTIZONEINFO = "InvalidParameter.InvalidMultiZoneInfo"

	// 节点数量NodeNum的取值和预期不符。
	INVALIDPARAMETER_INVALIDNODENUM = "InvalidParameter.InvalidNodeNum"

	// 节点规格NodeType的取值和预期不符。
	INVALIDPARAMETER_INVALIDNODETYPE = "InvalidParameter.InvalidNodeType"

	// 分页起始值Offset的取值和预期不符。
	INVALIDPARAMETER_INVALIDOFFSET = "InvalidParameter.InvalidOffset"

	// 操作类型OpType的取值和预期不符。
	INVALIDPARAMETER_INVALIDOPTYPE = "InvalidParameter.InvalidOpType"

	// 可维护时间段OperationDuration的取值和预期不符。
	INVALIDPARAMETER_INVALIDOPERATIONDURATION = "InvalidParameter.InvalidOperationDuration"

	// 排序字段OrderByKey的取值和预期不符。
	INVALIDPARAMETER_INVALIDORDERBYKEY = "InvalidParameter.InvalidOrderByKey"

	// 排序方式OrderByType的取值和预期不符。
	INVALIDPARAMETER_INVALIDORDERBYTYPE = "InvalidParameter.InvalidOrderByType"

	// 地域Region的取值与预期不符。
	INVALIDPARAMETER_INVALIDREGION = "InvalidParameter.InvalidRegion"

	// 安全组id列表SecurityGroupIds的取值和预期不符。
	INVALIDPARAMETER_INVALIDSECURITYGROUPIDS = "InvalidParameter.InvalidSecurityGroupIds"

	// 节点标签信息TagInfo的取值和预期不符。
	INVALIDPARAMETER_INVALIDTAGINFO = "InvalidParameter.InvalidTagInfo"

	// 节点标签信息列表TagList的取值和预期不符。
	INVALIDPARAMETER_INVALIDTAGLIST = "InvalidParameter.InvalidTagList"

	// 节点类型Type的取值和预期不符。
	INVALIDPARAMETER_INVALIDTYPE = "InvalidParameter.InvalidType"

	// 代金券ID列表VoucherIds的取值和预期不符。
	INVALIDPARAMETER_INVALIDVOUCHERIDS = "InvalidParameter.InvalidVoucherIds"

	// 可用区Zone的信息与预期不符。
	INVALIDPARAMETER_INVALIDZONE = "InvalidParameter.InvalidZone"

	// 参数取值错误。
	INVALIDPARAMETERVALUE = "InvalidParameterValue"

	// 计费类型ChargeType的参数取值有误。
	INVALIDPARAMETERVALUE_CHARGETYPE = "InvalidParameterValue.ChargeType"

	// 自动续费标识RenewFlag的参数取值有误。
	INVALIDPARAMETERVALUE_RENEWFLAG = "InvalidParameterValue.RenewFlag"

	// 该账号下的集群数超过限额。
	LIMITEXCEEDED_CLUSTERNUM = "LimitExceeded.ClusterNum"

	// 缺少参数错误。
	MISSINGPARAMETER = "MissingParameter"

	// 请求的次数超过了频率限制。
	REQUESTLIMITEXCEEDED = "RequestLimitExceeded"

	// 资源被占用。
	RESOURCEINUSE = "ResourceInUse"

	// 订单被锁定。
	RESOURCEINUSE_ORDER = "ResourceInUse.Order"

	// 资源不足。
	RESOURCEINSUFFICIENT = "ResourceInsufficient"

	// 账户余额不足。
	RESOURCEINSUFFICIENT_BALANCE = "ResourceInsufficient.Balance"

	// 隐藏可用区专用主节点资源不足。
	RESOURCEINSUFFICIENT_HIDDENZONE = "ResourceInsufficient.HiddenZone"

	// 子网剩余ip数量不足。
	RESOURCEINSUFFICIENT_SUBNET = "ResourceInsufficient.Subnet"

	// 可用区资源不足。
	RESOURCEINSUFFICIENT_ZONE = "ResourceInsufficient.Zone"

	// 资源不存在。
	RESOURCENOTFOUND = "ResourceNotFound"

	// CAM资源未找到。
	RESOURCENOTFOUND_CAMINFONOTFOUND = "ResourceNotFound.CAMInfoNotFound"

	// 数据库资源获取失败。
	RESOURCENOTFOUND_DBINFONOTFOUND = "ResourceNotFound.DBInfoNotFound"

	// 磁盘相关资源获取失败。
	RESOURCENOTFOUND_DISKINFONOTFOUND = "ResourceNotFound.DiskInfoNotFound"

	// 安全组信息获取失败。
	RESOURCENOTFOUND_SECURITYGROUPNOTFOUND = "ResourceNotFound.SecurityGroupNotFound"

	// 获取计费资源失败。
	RESOURCENOTFOUND_TRADECGWNOTFOUND = "ResourceNotFound.TradeCgwNotFound"

	// 白名单资源获取失败。
	RESOURCENOTFOUND_WHITELISTNOTFOUND = "ResourceNotFound.WhiteListNotFound"

	// 资源不可用。
	RESOURCEUNAVAILABLE = "ResourceUnavailable"

	// 资源售罄。
	RESOURCESSOLDOUT = "ResourcesSoldOut"

	// 未授权操作。
	UNAUTHORIZEDOPERATION = "UnauthorizedOperation"

	// Uin不在白名单中。
	UNAUTHORIZEDOPERATION_UINNOTINWHITELIST = "UnauthorizedOperation.UinNotInWhiteList"

	// 未知参数错误。
	UNKNOWNPARAMETER = "UnknownParameter"

	// 操作不支持。
	UNSUPPORTEDOPERATION = "UnsupportedOperation"

	// 不支持该认证方式。
	UNSUPPORTEDOPERATION_BASICSECURITYTYPE = "UnsupportedOperation.BasicSecurityType"

	// 不支持该操作，license的类型有误。
	UNSUPPORTEDOPERATION_LICENSEERROR = "UnsupportedOperation.LicenseError"

	// 不支持该操作，实例状态有误。
	UNSUPPORTEDOPERATION_STATUSNOTSUPPORT = "UnsupportedOperation.StatusNotSupport"

	// VPC资源获取失败。
	UNSUPPORTEDOPERATION_VPCINFONOTFOUND = "UnsupportedOperation.VPCInfoNotFound"
)
