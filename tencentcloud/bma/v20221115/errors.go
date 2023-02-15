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

package v20221115

const (
	// 此产品的特有错误码

	// CAM签名/鉴权错误。
	AUTHFAILURE = "AuthFailure"

	// DryRun 操作，代表请求将会是成功的，只是多传了 DryRun 参数。
	DRYRUNOPERATION = "DryRunOperation"

	// 操作失败。
	FAILEDOPERATION = "FailedOperation"

	// 操作失败，企业名称有误。
	FAILEDOPERATION_COMPANYNAMEWRONG = "FailedOperation.CompanyNameWrong"

	// 操作失败，已存在企业信息。
	FAILEDOPERATION_EXISTCOMPANY = "FailedOperation.ExistCompany"

	// 操作失败，已存在检测对象。
	FAILEDOPERATION_EXISTDETECTTARGET = "FailedOperation.ExistDetectTarget"

	// 操作失败，已存在软件信息。
	FAILEDOPERATION_EXISTSOFTWARE = "FailedOperation.ExistSoftware"

	// 该资源已使用，请勿重复使用。
	FAILEDOPERATION_REUSERESOURCE = "FailedOperation.ReuseResource"

	// 操作失败，验证码错误。
	FAILEDOPERATION_SMSVERIFY = "FailedOperation.SmsVerify"

	// 操作失败，资源不可用。
	FAILEDOPERATION_UNUSABLERESOURCE = "FailedOperation.UnusableResource"

	// 内部错误。
	INTERNALERROR = "InternalError"

	// 服务暂时不可用，请稍后重试。若无法解决，请联系智能客服。
	INTERNALERROR_SMSVERIFY = "InternalError.SmsVerify"

	// 接口不存在。
	INVALIDACTION = "InvalidAction"

	// 参数错误。
	INVALIDPARAMETER = "InvalidParameter"

	// 无效的隐私文本，请使用规定方法上传UTF-8编码格式文件。
	INVALIDPARAMETER_PRIVACYTEXTISNOTUTF8 = "InvalidParameter.PrivacyTextIsNotUtf8"

	// 无效的隐私文本下载链接，请使用规定方法上传文件。
	INVALIDPARAMETER_PRIVACYTEXTURL = "InvalidParameter.PrivacyTextURL"

	// 无效的软件下载链接，请使用规定方法上传文件。
	INVALIDPARAMETER_SOFTWAREURL = "InvalidParameter.SoftwareURL"

	// 参数取值错误。
	INVALIDPARAMETERVALUE = "InvalidParameterValue"

	// 超过配额限制。
	LIMITEXCEEDED = "LimitExceeded"

	// 缺少参数错误。
	MISSINGPARAMETER = "MissingParameter"

	// 接口版本不存在。
	NOSUCHVERSION = "NoSuchVersion"

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

	// 接口不支持所传地域。
	UNSUPPORTEDREGION = "UnsupportedRegion"
)
