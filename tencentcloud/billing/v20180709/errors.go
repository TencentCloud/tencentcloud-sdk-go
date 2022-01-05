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

package v20180709

const (
	// 此产品的特有错误码

	// CAM签名/鉴权错误。
	AUTHFAILURE = "AuthFailure"

	// 没有权限。
	AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"

	// 操作失败。
	FAILEDOPERATION = "FailedOperation"

	// 代理支付设备不能降配。
	FAILEDOPERATION_AGENTPAYDEALCANNOTDOWN = "FailedOperation.AgentPayDealCannotDown"

	// 账户余额不足。
	FAILEDOPERATION_BALANCEINSUFFICIENT = "FailedOperation.BalanceInsufficient"

	// 订单状态错误，只有未支付订单才能支付。
	FAILEDOPERATION_INVALIDDEAL = "FailedOperation.InvalidDeal"

	// 代金券不可用。
	FAILEDOPERATION_INVALIDVOUCHER = "FailedOperation.InvalidVoucher"

	// 一起购买的订单必须同时支付。
	FAILEDOPERATION_NEEDPAYTOGETER = "FailedOperation.NeedPayTogeter"

	// 套餐订单需一起购买。
	FAILEDOPERATION_NEEDPAYTOGETHER = "FailedOperation.NeedPayTogether"

	// 支付失败，请联系腾讯云工作人员处理。
	FAILEDOPERATION_PAYPRICEERROR = "FailedOperation.PayPriceError"

	// 支付成功但发货失败，请联系腾讯云工作人员处理。
	FAILEDOPERATION_PAYSUCCDELIVERFAILED = "FailedOperation.PaySuccDeliverFailed"

	// 获取数据条数失败。
	FAILEDOPERATION_QUERYCOUNTFAILED = "FailedOperation.QueryCountFailed"

	// 汇总数据正在构建中，请稍后再试。
	FAILEDOPERATION_SUMMARYDATANOTREADY = "FailedOperation.SummaryDataNotReady"

	// 不存在该分账标签键。
	FAILEDOPERATION_TAGKEYNOTEXIST = "FailedOperation.TagKeyNotExist"

	// 内部错误。
	INTERNALERROR = "InternalError"

	// 网关错误。
	INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"

	// 参数错误。
	INVALIDPARAMETER = "InvalidParameter"

	// 参数取值错误。
	INVALIDPARAMETERVALUE = "InvalidParameterValue"

	// 账号没有cam授权。
	UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"

	// 因账号安全升级，购买云资源需完善您的实名信息。
	UNAUTHORIZEDOPERATION_CERTIFICATIONNEEDUPGRADE = "UnauthorizedOperation.CertificationNeedUpgrade"

	// 账号没有实名认证，支付失败。
	UNAUTHORIZEDOPERATION_NOTCERTIFICATION = "UnauthorizedOperation.NotCertification"

	// 操作不支持。
	UNSUPPORTEDOPERATION = "UnsupportedOperation"
)
