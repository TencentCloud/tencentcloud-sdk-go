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

package v20230901

const (
	// 此产品的特有错误码

	// 操作失败。
	FAILEDOPERATION = "FailedOperation"

	// 控制台服务异常。
	FAILEDOPERATION_CONSOLESERVERERROR = "FailedOperation.ConsoleServerError"

	// LogoUrl 或 LogoImage 有误，水印图下载失败。
	FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownloadError"

	// 引擎层请求超时；请稍后重试。
	FAILEDOPERATION_ENGINEREQUESTTIMEOUT = "FailedOperation.EngineRequestTimeout"

	// 引擎层内部错误；请稍后重试。
	FAILEDOPERATION_ENGINESERVERERROR = "FailedOperation.EngineServerError"

	// 引擎层请求超过限额；请稍后重试。
	FAILEDOPERATION_ENGINESERVERLIMITEXCEEDED = "FailedOperation.EngineServerLimitExceeded"

	// 免费资源包余量已用尽，请购买资源包或开通后付费。
	FAILEDOPERATION_FREERESOURCEPACKEXHAUSTED = "FailedOperation.FreeResourcePackExhausted"

	// 图片包含敏感内容
	FAILEDOPERATION_GENERATEIMAGEFAILED = "FailedOperation.GenerateImageFailed"

	// 水印图解码失败
	FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"

	// LogoUrl 或 LogoImage 有误，水印图下载失败。
	FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"

	// 任务不存在。
	FAILEDOPERATION_JOBNOTFOUND = "FailedOperation.JobNotFound"

	// 合作伙伴账号不允许开通，请更换账号。
	FAILEDOPERATION_PARTNERACCOUNTUNSUPPORT = "FailedOperation.PartnerAccountUnSupport"

	// 资源包余量已用尽，请购买资源包或开通后付费。
	FAILEDOPERATION_RESOURCEPACKEXHAUSTED = "FailedOperation.ResourcePackExhausted"

	// 服务未开通，请前往控制台申请试用。
	FAILEDOPERATION_SERVICENOTACTIVATED = "FailedOperation.ServiceNotActivated"

	// 用户主动停服。
	FAILEDOPERATION_SERVICESTOP = "FailedOperation.ServiceStop"

	// 欠费停服。
	FAILEDOPERATION_SERVICESTOPARREARS = "FailedOperation.ServiceStopArrears"

	// 后付费设置次数超过每月限制。
	FAILEDOPERATION_SETPAYMODEEXCEED = "FailedOperation.SetPayModeExceed"

	// 用户未实名，请先进行实名认证。
	FAILEDOPERATION_USERUNAUTHERROR = "FailedOperation.UserUnAuthError"

	// 内部错误。
	INTERNALERROR = "InternalError"

	// 参数错误。
	INVALIDPARAMETER = "InvalidParameter"

	// 参数不合法。
	INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"

	// 参数取值错误。
	INVALIDPARAMETERVALUE = "InvalidParameterValue"

	// 不支持的图片格式。
	INVALIDPARAMETERVALUE_INVALIDIMAGEFORMAT = "InvalidParameterValue.InvalidImageFormat"

	// 不支持的图片分辨率。
	INVALIDPARAMETERVALUE_INVALIDIMAGERESOLUTION = "InvalidParameterValue.InvalidImageResolution"

	// 图片大小超出限制。
	INVALIDPARAMETERVALUE_INVALIDIMAGESIZE = "InvalidParameterValue.InvalidImageSize"

	// 模型不存在。
	INVALIDPARAMETERVALUE_MODEL = "InvalidParameterValue.Model"

	// 参数字段或者值有误
	INVALIDPARAMETERVALUE_PARAMETERVALUEERROR = "InvalidParameterValue.ParameterValueError"

	// URL格式不合法。
	INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"

	// 超过配额限制。
	LIMITEXCEEDED = "LimitExceeded"

	// 图片可能包含敏感信息，请重试
	OPERATIONDENIED_IMAGEILLEGALDETECTED = "OperationDenied.ImageIllegalDetected"

	// 文本包含违法违规信息，审核不通过。
	OPERATIONDENIED_TEXTILLEGALDETECTED = "OperationDenied.TextIllegalDetected"

	// 同时处理的任务数过多，请稍后重试。
	REQUESTLIMITEXCEEDED_JOBNUMEXCEED = "RequestLimitExceeded.JobNumExceed"

	// 资源不足。
	RESOURCEINSUFFICIENT = "ResourceInsufficient"

	// 计费资源已耗尽。
	RESOURCEINSUFFICIENT_CHARGERESOURCEEXHAUST = "ResourceInsufficient.ChargeResourceExhaust"

	// 资源不存在。
	RESOURCENOTFOUND = "ResourceNotFound"

	// 资源不可用。
	RESOURCEUNAVAILABLE = "ResourceUnavailable"

	// 账号已欠费。
	RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"

	// 余额不足。
	RESOURCEUNAVAILABLE_LOWBALANCE = "ResourceUnavailable.LowBalance"

	// 计费状态未知，请确认是否已在控制台开通服务。
	RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"

	// 账号已停服。
	RESOURCEUNAVAILABLE_STOPUSING = "ResourceUnavailable.StopUsing"
)
