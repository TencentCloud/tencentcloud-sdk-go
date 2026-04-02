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

	// 请求中传入参数 `%(key)s` 必选再 `%(value)s` 范围内。
	INVALIDPARAMETER_INPUTOUTOFRANGE = "InvalidParameter.InputOutOfRange"

	// 参数取值错误。
	INVALIDPARAMETERVALUE = "InvalidParameterValue"

	// 实例名称仅支持以大小写字符或中文开头，支持数字、英文句号、或段划线、下划线。
	INVALIDPARAMETERVALUE_INSTANCENAME = "InvalidParameterValue.InstanceName"

	// 参数 `%(parameter)s` 值 `%(value)s` 是无效的。正确且完整的值形如 `%(template)s`。
	INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"

	// 参数值 `%(value)s` 长度不能大于 `%(max_size)s`。
	INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"

	// 缺少参数错误。
	MISSINGPARAMETER = "MissingParameter"

	// 账户被冻结，不支持当前操作。
	UNSUPPORTEDOPERATION_ACCOUNTFROZEN = "UnsupportedOperation.AccountFrozen"

	// 抱歉，您的操作暂时无法完成，请稍后重试或联系客服。
	UNSUPPORTEDOPERATION_BILLINGFAILED = "UnsupportedOperation.BillingFailed"

	// 跨境承诺书没签约，不支持当前操作。
	UNSUPPORTEDOPERATION_CROSSBORDERPROMISENOTSIGNED = "UnsupportedOperation.CrossBorderPromiseNotSigned"

	// 当前账号不支持此操作。
	UNSUPPORTEDOPERATION_CURRENTACCOUNTNOTALLOWED = "UnsupportedOperation.CurrentAccountNotAllowed"

	// 当前用户，不支持设置跨境类型。
	UNSUPPORTEDOPERATION_ENABLECROSSBORDER = "UnsupportedOperation.EnableCrossBorder"

	// 账户余额不足。
	UNSUPPORTEDOPERATION_INSUFFICIENTFUNDS = "UnsupportedOperation.InsufficientFunds"

	// 当前全球加速实例无法创建跨境加速地域或终端节点组。
	UNSUPPORTEDOPERATION_UNABLECROSSBORDER = "UnsupportedOperation.UnableCrossBorder"
)
