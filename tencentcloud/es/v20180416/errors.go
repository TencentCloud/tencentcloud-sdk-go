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

	// 账户未绑定信用卡或paypal，无法支付。
	FAILEDOPERATION_NOPAYMENT = "FailedOperation.NoPayment"

	// 用户未实名认证。
	FAILEDOPERATION_NOTAUTHENTICATED = "FailedOperation.NotAuthenticated"

	// 不支持反向调节节点配置和磁盘容量。
	FAILEDOPERATION_UNSUPPORTREVERSEREGULATIONNODETYPEANDDISK = "FailedOperation.UnsupportReverseRegulationNodeTypeAndDisk"

	// 内部错误。
	INTERNALERROR = "InternalError"

	// 参数错误。
	INVALIDPARAMETER = "InvalidParameter"

	// 资源被占用。
	RESOURCEINUSE = "ResourceInUse"

	// 资源不足。
	RESOURCEINSUFFICIENT = "ResourceInsufficient"

	// 账户余额不足。
	RESOURCEINSUFFICIENT_BALANCE = "ResourceInsufficient.Balance"

	// 隐藏可用区专用主节点资源不足。
	RESOURCEINSUFFICIENT_HIDDENZONE = "ResourceInsufficient.HiddenZone"

	// 子网剩余ip数量不足。
	RESOURCEINSUFFICIENT_SUBNET = "ResourceInsufficient.Subnet"

	// 资源不存在。
	RESOURCENOTFOUND = "ResourceNotFound"

	// 资源不可用。
	RESOURCEUNAVAILABLE = "ResourceUnavailable"

	// 资源售罄。
	RESOURCESSOLDOUT = "ResourcesSoldOut"

	// 未授权操作。
	UNAUTHORIZEDOPERATION = "UnauthorizedOperation"

	// 未知参数错误。
	UNKNOWNPARAMETER = "UnknownParameter"

	// 操作不支持。
	UNSUPPORTEDOPERATION = "UnsupportedOperation"
)
