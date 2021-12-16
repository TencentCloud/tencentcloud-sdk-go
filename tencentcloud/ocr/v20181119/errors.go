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

package v20181119

const (
	// 此产品的特有错误码

	// 帐号已欠费。
	FAILEDOPERATION_ARREARSERROR = "FailedOperation.ArrearsError"

	// 今日次数达到限制。
	FAILEDOPERATION_COUNTLIMITERROR = "FailedOperation.CountLimitError"

	// 检测失败。
	FAILEDOPERATION_DETECTFAILED = "FailedOperation.DetectFailed"

	// 文件下载失败。
	FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"

	// 图片内容为空。
	FAILEDOPERATION_EMPTYIMAGEERROR = "FailedOperation.EmptyImageError"

	// 引擎识别超时。
	FAILEDOPERATION_ENGINERECOGNIZETIMEOUT = "FailedOperation.EngineRecognizeTimeout"

	// 身份证信息不合法（身份证号、姓名字段校验非法等）。
	FAILEDOPERATION_IDCARDINFOILLEGAL = "FailedOperation.IdCardInfoIllegal"

	// 图片模糊。
	FAILEDOPERATION_IMAGEBLUR = "FailedOperation.ImageBlur"

	// 图片解码失败。
	FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"

	// 照片未检测到名片。
	FAILEDOPERATION_IMAGENOBUSINESSCARD = "FailedOperation.ImageNoBusinessCard"

	// 图片中未检测到身份证。
	FAILEDOPERATION_IMAGENOIDCARD = "FailedOperation.ImageNoIdCard"

	// 图片中未检测到文本。
	FAILEDOPERATION_IMAGENOTEXT = "FailedOperation.ImageNoText"

	// 图片尺寸过大，请参考输出参数中关于图片大小限制的说明。
	FAILEDOPERATION_IMAGESIZETOOLARGE = "FailedOperation.ImageSizeTooLarge"

	// 发票数据不一致。
	FAILEDOPERATION_INVOICEMISMATCH = "FailedOperation.InvoiceMismatch"

	// 输入的Language不支持。
	FAILEDOPERATION_LANGUAGENOTSUPPORT = "FailedOperation.LanguageNotSupport"

	// 照片中存在多张卡。
	FAILEDOPERATION_MULTICARDERROR = "FailedOperation.MultiCardError"

	// 非香港身份证。
	FAILEDOPERATION_NOHKIDCARD = "FailedOperation.NoHKIDCard"

	// 非马来身份证。
	FAILEDOPERATION_NOMASIDCARD = "FailedOperation.NoMASIDCard"

	// 非护照。
	FAILEDOPERATION_NOPASSPORT = "FailedOperation.NoPassport"

	// OCR识别失败。
	FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"

	// 查询无记录。
	FAILEDOPERATION_QUERYNORECORD = "FailedOperation.QueryNoRecord"

	// 未知错误。
	FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"

	// 服务未开通。
	FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"

	// 内部错误。
	INTERNALERROR = "InternalError"

	// Config不是有效的JSON格式。
	INVALIDPARAMETER_CONFIGFORMATERROR = "InvalidParameter.ConfigFormatError"

	// 图片解码失败。
	INVALIDPARAMETER_ENGINEIMAGEDECODEFAILED = "InvalidParameter.EngineImageDecodeFailed"

	// 无效的GTIN。
	INVALIDPARAMETER_INVALIDGTINERROR = "InvalidParameter.InvalidGTINError"

	// 参数值错误。
	INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"

	// 文件内容太大。
	LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"

	// 发票不存在。
	RESOURCENOTFOUND_NOINVOICE = "ResourceNotFound.NoInvoice"

	// 不支持当天发票查询。
	RESOURCENOTFOUND_NOTSUPPORTCURRENTINVOICEQUERY = "ResourceNotFound.NotSupportCurrentInvoiceQuery"

	// 计费状态异常。
	RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
)
