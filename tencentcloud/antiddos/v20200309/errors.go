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

package v20200309

const (
	// 此产品的特有错误码

	// CAM签名/鉴权错误。
	AUTHFAILURE = "AuthFailure"

	// DryRun 操作，代表请求将会是成功的，只是多传了 DryRun 参数。
	DRYRUNOPERATION = "DryRunOperation"

	// 操作失败。
	FAILEDOPERATION = "FailedOperation"

	// 账户余额不足
	FAILEDOPERATION_INSUFFICIENTBALANCE = "FailedOperation.InsufficientBalance"

	// 内部错误。
	INTERNALERROR = "InternalError"

	// 参数错误。
	INVALIDPARAMETER = "InvalidParameter"

	// 这个参数需要开启白名单才可以使用
	INVALIDPARAMETER_SPECIALPARAMETERFORSPECIALACCOUNT = "InvalidParameter.SpecialParameterForSpecialAccount"

	// 参数取值错误。
	INVALIDPARAMETERVALUE = "InvalidParameterValue"

	// 超过配额限制。
	LIMITEXCEEDED = "LimitExceeded"

	// 缺少参数错误。
	MISSINGPARAMETER = "MissingParameter"

	// 操作被拒绝。
	OPERATIONDENIED = "OperationDenied"

	// 请求的次数超过了频率限制。
	REQUESTLIMITEXCEEDED = "RequestLimitExceeded"

	// 资源被占用。
	RESOURCEINUSE = "ResourceInUse"

	// 资源不足。
	RESOURCEINSUFFICIENT = "ResourceInsufficient"

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

	// 不支持此业务带宽
	UNSUPPORTEDOPERATION_BANDWIDTHNOTSUPPORTED = "UnsupportedOperation.BandwidthNotSupported"

	// 不支持此保底防护带宽的取值
	UNSUPPORTEDOPERATION_BASICPROTECTBANDWIDTHNOTSUPPORTED = "UnsupportedOperation.BasicProtectBandwidthNotSupported"

	// 不支持此弹性带宽的取值
	UNSUPPORTEDOPERATION_ELASTICPROTECTBANDWIDTHNOTSUPPORTED = "UnsupportedOperation.ElasticProtectBandwidthNotSupported"

	// 不支持此种付费类型
	UNSUPPORTEDOPERATION_INSTANCECHARGETYPENOTSUPPORTED = "UnsupportedOperation.InstanceChargeTypeNotSupported"

	// 不支持该计费周期，参考对应高防包支持的计费周期
	UNSUPPORTEDOPERATION_PERIODNOTSUPPORTED = "UnsupportedOperation.PeriodNotSupported"

	// 不支持该防护IP数量取值
	UNSUPPORTEDOPERATION_PROTECTIPCOUNTNOTSUPPORTED = "UnsupportedOperation.ProtectIpCountNotSupported"

	// 不支持此地域
	UNSUPPORTEDOPERATION_REGIONNOTSUPPORTED = "UnsupportedOperation.RegionNotSupported"
)
