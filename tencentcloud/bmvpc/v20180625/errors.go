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

package v20180625

const (
	// 此产品的特有错误码

	// 操作失败。
	FAILEDOPERATION = "FailedOperation"

	// 内部错误。
	INTERNALERROR = "InternalError"

	// 参数错误。
	INVALIDPARAMETER = "InvalidParameter"

	// 超过配额限制。
	LIMITEXCEEDED = "LimitExceeded"

	// 可创建的VPN通道SPD ACL数达到上限。
	LIMITEXCEEDED_SPDACLLIMIT = "LimitExceeded.SpdAclLimit"

	// 可创建的VPN通道SPD ACL目的地址数达到上限。
	LIMITEXCEEDED_SPDDNETLIMIT = "LimitExceeded.SpdDnetLimit"

	// 可创建的VPN通道SPD ACL源地址数达到上限。
	LIMITEXCEEDED_SPDSNETLIMIT = "LimitExceeded.SpdSnetLimit"

	// 资源被占用。
	RESOURCEINUSE = "ResourceInUse"

	// 资源不存在。
	RESOURCENOTFOUND = "ResourceNotFound"

	// 对端网关不存在。
	RESOURCENOTFOUND_CUSTOMERGATEWAYNOTEXIST = "ResourceNotFound.CustomerGatewayNotExist"

	// 联通状态的通道才可以重置。
	RESOURCENOTFOUND_NOTAVAIBLE = "ResourceNotFound.NotAvaible"

	// VPC不存在。
	RESOURCENOTFOUND_VPCNOTEXIST = "ResourceNotFound.VpcNotExist"

	// 对等连接不存在。
	RESOURCENOTFOUND_VPCPEERNOTEXIST = "ResourceNotFound.VpcPeerNotExist"

	// VPN通道不存在。
	RESOURCENOTFOUND_VPNCONNNOTEXIST = "ResourceNotFound.VpnConnNotExist"

	// VPN网关不存在。
	RESOURCENOTFOUND_VPNGWNOTEXIST = "ResourceNotFound.VpnGwNotExist"

	// 资源不可用。
	RESOURCEUNAVAILABLE = "ResourceUnavailable"

	// 未授权操作。
	UNAUTHORIZEDOPERATION = "UnauthorizedOperation"

	// 操作不支持。
	UNSUPPORTEDOPERATION = "UnsupportedOperation"

	// 对端网关地址已存在。
	UNSUPPORTEDOPERATION_CUSTOMERGATEWAYADDREXIST = "UnsupportedOperation.CustomerGatewayAddrExist"

	// 对端网关地址不合法。
	UNSUPPORTEDOPERATION_CUSTOMERGATEWAYADDRINVALID = "UnsupportedOperation.CustomerGatewayAddrInvalid"

	// 对等连接状态不正确。
	UNSUPPORTEDOPERATION_INVALIDVPCPEERSTATE = "UnsupportedOperation.InvalidVpcPeerState"

	// VPN通道状态不正确。
	UNSUPPORTEDOPERATION_INVALIDVPNCONNSTATE = "UnsupportedOperation.InvalidVpnConnState"

	// SPD策略源目的地址冲突。
	UNSUPPORTEDOPERATION_SPDACLCIDRINVALID = "UnsupportedOperation.SpdAclCidrInvalid"

	// SPD策略源地址不在VPC网段内。
	UNSUPPORTEDOPERATION_SPDSNETNOTINCIDR = "UnsupportedOperation.SpdSnetNotInCidr"

	// 对等连接CIDR冲突。
	UNSUPPORTEDOPERATION_VPCCIDRCONFICT = "UnsupportedOperation.VpcCidrConfict"

	// 对等连接已存在。
	UNSUPPORTEDOPERATION_VPCPEEREXIST = "UnsupportedOperation.VpcPeerExist"

	// VPN通道已存在。
	UNSUPPORTEDOPERATION_VPNCONNEXIST = "UnsupportedOperation.VpnConnExist"

	// 存在VPN通道引用。
	UNSUPPORTEDOPERATION_VPNCONNINUSE = "UnsupportedOperation.VpnConnInUse"
)
