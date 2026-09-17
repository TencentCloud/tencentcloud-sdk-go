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

package v20260401

const (
	// 此产品的特有错误码

	// 操作失败。
	FAILEDOPERATION = "FailedOperation"

	// IP 仍绑定物理服务器，需先销毁物理机再释放IP
	FAILEDOPERATION_IPSTILLBOUNDTOSERVER = "FailedOperation.IpStillBoundToServer"

	// 实例为 BGP/OSPF 模式，已自动分配或释放IP，不需要申请或释放IP
	FAILEDOPERATION_NOTSUPPORTEDFORDYNAMICINSTANCE = "FailedOperation.NotSupportedForDynamicInstance"

	// 该 AppId 在指定可用区已存在私网实例
	FAILEDOPERATION_PRIVATEINSTANCEDUPLICATE = "FailedOperation.PrivateInstanceDuplicate"

	// 该 AppId 在指定可用区已存在公网实例
	FAILEDOPERATION_PUBLICINSTANCEDUPLICATE = "FailedOperation.PublicInstanceDuplicate"

	// 内部错误。
	INTERNALERROR = "InternalError"

	// 参数错误。
	INVALIDPARAMETER = "InvalidParameter"

	// **BGP/OSPF 模式**：配额值不在有效范围内
	INVALIDPARAMETER_INVALIDQUOTA = "InvalidParameter.InvalidQuota"

	// 参数错误
	INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"

	// 参数取值错误。
	INVALIDPARAMETERVALUE = "InvalidParameterValue"

	// 子网网络地址格式不合法
	INVALIDPARAMETERVALUE_INVALIDNETWORK = "InvalidParameterValue.InvalidNetwork"

	// 指定 IPv4 不属于该实例或未曾申请（Type=ipv4）
	INVALIDPARAMETERVALUE_PUBLICIPNOTAVAILABLE = "InvalidParameterValue.PublicIpNotAvailable"

	// 指定 IPv6 不属于该实例或未曾申请（Type=ipv6）
	INVALIDPARAMETERVALUE_PUBLICIPV6NOTAVAILABLE = "InvalidParameterValue.PublicIpv6NotAvailable"

	// LimitExceeded.InstanceQuota
	LIMITEXCEEDED_INSTANCEQUOTA = "LimitExceeded.InstanceQuota"

	// 申请Ipv6数量超出配额限制
	LIMITEXCEEDED_IPV6QUOTAEXCEEDED = "LimitExceeded.Ipv6QuotaExceeded"

	// 用户在该可用区未配置 IPv6 配额（Type=ipv6）
	LIMITEXCEEDED_IPV6QUOTANOTCONFIGURED = "LimitExceeded.Ipv6QuotaNotConfigured"

	// 申请Ipv4数量超出配额限制
	LIMITEXCEEDED_QUOTAEXCEEDED = "LimitExceeded.QuotaExceeded"

	// **BGP/OSPF 模式**：用户在该可用区未配置 IPv4 配额
	LIMITEXCEEDED_QUOTANOTCONFIGURED = "LimitExceeded.QuotaNotConfigured"

	// 缺少参数错误。
	MISSINGPARAMETER = "MissingParameter"

	// 实例下仍有服务器绑定，需先释放所有物理机
	RESOURCEINUSE_PRIVATEINSTANCEINUSE = "ResourceInUse.PrivateInstanceInUse"

	// 实例下仍有服务器或Ip绑定，需先释放所有物理机或Ip
	RESOURCEINUSE_PUBLICINSTANCEINUSE = "ResourceInUse.PublicInstanceInUse"

	// 资源不足。
	RESOURCEINSUFFICIENT = "ResourceInsufficient"

	// **BGP/OSPF 模式**：可用区无空闲 CIDR 段（/26 或 /25）可分配
	RESOURCEINSUFFICIENT_NOAVAILABLECIDR = "ResourceInsufficient.NoAvailableCidr"

	// IPv4 剩余地址不足
	RESOURCEINSUFFICIENT_PUBLICIPINSUFFICIENT = "ResourceInsufficient.PublicIpInsufficient"

	// IPv6 地址剩余不足
	RESOURCEINSUFFICIENT_PUBLICIPV6INSUFFICIENT = "ResourceInsufficient.PublicIpv6Insufficient"

	// 物理机实例不存在
	RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"

	// 指定的私网实例不存在
	RESOURCENOTFOUND_PRIVATEINSTANCENOTFOUND = "ResourceNotFound.PrivateInstanceNotFound"

	// 指定的公网实例不存在。
	RESOURCENOTFOUND_PUBLICINSTANCENOTFOUND = "ResourceNotFound.PublicInstanceNotFound"

	// ResourceNotFound.Zone
	RESOURCENOTFOUND_ZONE = "ResourceNotFound.Zone"

	// 可用区不存在。
	RESOURCENOTFOUND_ZONENOTFOUND = "ResourceNotFound.ZoneNotFound"

	// AppId 与实例归属不匹配
	UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"

	// UnsupportedOperation.InvalidInstanceState
	UNSUPPORTEDOPERATION_INVALIDINSTANCESTATE = "UnsupportedOperation.InvalidInstanceState"

	// UnsupportedOperation.LegacyCompatBatchMixed
	UNSUPPORTEDOPERATION_LEGACYCOMPATBATCHMIXED = "UnsupportedOperation.LegacyCompatBatchMixed"
)
