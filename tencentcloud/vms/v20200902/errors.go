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

package v20200902

const (
	// 此产品的特有错误码

	// 访问上下游模块超时。
	FAILEDOPERATION_ACCESSUPSTREAMTIMEOUT = "FailedOperation.AccessUpstreamTimeout"

	// 语音内容中含有敏感词，请[联系我们](https://cloud.tencent.com/document/product/1128/37720)沟通解决。
	FAILEDOPERATION_CONTAINSENSITIVEWORD = "FailedOperation.ContainSensitiveWord"

	// 后端请求包解析失败，通常由于没有遵守 API 接口说明规范导致的，请参见[1004错误详解](https://cloud.tencent.com/document/product/1128/38004#Q7)。
	FAILEDOPERATION_FAILRESOLVEPACKET = "FailedOperation.FailResolvePacket"

	// 套餐包余量不足，请及时[购买语音套餐包](https://cloud.tencent.com/document/product/1128#buyPackage)。
	FAILEDOPERATION_INSUFFICIENTBALANCEINVOICEPACKAGE = "FailedOperation.InsufficientBalanceInVoicePackage"

	// 无效 JSON，请核查发送的请求是否为标准的 JSON 格式。
	FAILEDOPERATION_INVALIDJSONPARAMETERS = "FailedOperation.InvalidJsonParameters"

	// 无效参数，请核查发送的请求参数是否为对应 API 所需参数。
	FAILEDOPERATION_INVALIDPARAMETERS = "FailedOperation.InvalidParameters"

	// 解析请求包体时候失败。
	FAILEDOPERATION_JSONPARSEFAIL = "FailedOperation.JsonParseFail"

	// 语音未知错误，请[联系我们](https://cloud.tencent.com/document/product/1128/37720)沟通解决。
	FAILEDOPERATION_PARAMETERSOTHERERROR = "FailedOperation.ParametersOtherError"

	// 未申请号码或申请的号码资源已过期，请及时支付月功能费用和信息服务费用，具体操作请参见[购买指南](https://cloud.tencent.com/document/product/1128)。
	FAILEDOPERATION_PHONENUMBERUNAPPLIEDOREXPIRED = "FailedOperation.PhonenumberUnappliedOrExpired"

	// 模板未审核或请求的内容与审核通过的模板内容不匹配，请参见[1014错误详解](https://cloud.tencent.com/document/product/1128/38004
	FAILEDOPERATION_TEMPLATEINCORRECTORUNAPPROVED = "FailedOperation.TemplateIncorrectOrUnapproved"

	// 访问上游超时网络，请稍后重试。
	INTERNALERROR_ACCESSUPSTREAMTIMEOUT = "InternalError.AccessUpstreamTimeout"

	// 请求发起时间不正常，通常由您的服务器与腾讯云服务器之间的时间差超过10分钟引起。
	INTERNALERROR_REQUESTTIMEEXCEPTION = "InternalError.RequestTimeException"

	// 后端不存在该 REST API 接口，请核查 REST API 接口说明。
	INTERNALERROR_RESTAPIINTERFACENOTEXIST = "InternalError.RestApiInterfaceNotExist"

	// 后端 Sig 校验失败。
	INTERNALERROR_SIGVERIFICATIONFAIL = "InternalError.SigVerificationFail"

	// 内部sso通道超时。
	INTERNALERROR_SSOSENDRECVFAIL = "InternalError.SsoSendRecvFail"

	// 语音上游错误，请[联系我们](https://cloud.tencent.com/document/product/1128/37720)沟通解决。
	INTERNALERROR_UPSTREAMERROR = "InternalError.UpstreamError"

	// 被叫手机号码格式校验失败。
	INVALIDPARAMETERVALUE_CALLEDNUMBERVERIFYFAIL = "InvalidParameterValue.CalledNumberVerifyFail"

	// 语音模板中单个变量长度超过限制，如需调整限制，请[联系我们](https://cloud.tencent.com/document/product/1128/37720)。
	INVALIDPARAMETERVALUE_CONTENTLENGTHLIMIT = "InvalidParameterValue.ContentLengthLimit"

	// SDK AppID 不存在。
	INVALIDPARAMETERVALUE_SDKAPPIDNOTEXIST = "InvalidParameterValue.SdkAppidNotExist"

	// 下发语音消息时命中频率限制策略，如需申请不受频率限制的测试号码或更改限制策略，请[联系我们](https://cloud.tencent.com/document/product/1128/37720)。
	LIMITEXCEEDED_DELIVERYFREQUENCYLIMIT = "LimitExceeded.DeliveryFrequencyLimit"

	// SDK AppID 禁用发送语音消息，如有需要请[联系我们](https://cloud.tencent.com/document/product/1128/37720)。
	UNAUTHORIZEDOPERATION_SDKAPPIDISDISABLED = "UnauthorizedOperation.SdkAppidIsDisabled"

	// 因腾讯云账号欠费被停止服务，请及时为您的腾讯云账号[充值](https://cloud.tencent.com/document/product/555/7425)缴清欠款。
	UNAUTHORIZEDOPERATION_SERVICESUSPENDDUETOARREARS = "UnauthorizedOperation.ServiceSuspendDueToArrears"

	// VoiceSdkAppid 校验失败。
	UNAUTHORIZEDOPERATION_VOICESDKAPPIDVERIFYFAIL = "UnauthorizedOperation.VoiceSdkAppidVerifyFail"

	// 操作不支持。
	UNSUPPORTEDOPERATION = "UnsupportedOperation"
)
