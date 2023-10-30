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

	// 账号已欠费。
	FAILEDOPERATION_ARREARSERROR = "FailedOperation.ArrearsError"

	// 身份证CardSide类型错误
	FAILEDOPERATION_CARDSIDEERROR = "FailedOperation.CardSideError"

	// 今日次数达到限制。
	FAILEDOPERATION_COUNTLIMITERROR = "FailedOperation.CountLimitError"

	// 数据库异常。
	FAILEDOPERATION_DBERROR = "FailedOperation.DbError"

	// 检测失败。
	FAILEDOPERATION_DETECTFAILED = "FailedOperation.DetectFailed"

	// 文件下载失败。
	FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"

	// 图片内容为空。
	FAILEDOPERATION_EMPTYIMAGEERROR = "FailedOperation.EmptyImageError"

	// 引擎识别超时。
	FAILEDOPERATION_ENGINERECOGNIZETIMEOUT = "FailedOperation.EngineRecognizeTimeout"

	// 第二代身份证信息不合法或缺失（身份证号、姓名字段校验非法等）
	FAILEDOPERATION_IDCARDINFOILLEGAL = "FailedOperation.IdCardInfoIllegal"

	// 图片分辨率过小或身份证在原图中的占比过小
	FAILEDOPERATION_IDCARDTOOSMALL = "FailedOperation.IdCardTooSmall"

	// 银行卡信息非法。
	FAILEDOPERATION_ILLEGALBANKCARDERROR = "FailedOperation.IllegalBankCardError"

	// 图片模糊。
	FAILEDOPERATION_IMAGEBLUR = "FailedOperation.ImageBlur"

	// 图片解码失败。
	FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"

	// 照片未检测到名片。
	FAILEDOPERATION_IMAGENOBUSINESSCARD = "FailedOperation.ImageNoBusinessCard"

	// 图片中未检测到第二代身份证或临时身份证。
	FAILEDOPERATION_IMAGENOIDCARD = "FailedOperation.ImageNoIdCard"

	// 非指定卡类别图片
	FAILEDOPERATION_IMAGENOSPECIFIEDCARD = "FailedOperation.ImageNoSpecifiedCard"

	// 图片中未检测到文本。
	FAILEDOPERATION_IMAGENOTEXT = "FailedOperation.ImageNoText"

	// 图片尺寸过大，请参考输入参数中关于图片大小限制的说明。
	FAILEDOPERATION_IMAGESIZETOOLARGE = "FailedOperation.ImageSizeTooLarge"

	// 发票数据不一致。温馨提示：新版发票核验接口功能更完整，请尽快切换，如已切换请忽略。
	FAILEDOPERATION_INVOICEMISMATCH = "FailedOperation.InvoiceMismatch"

	// 输入的Language不支持。
	FAILEDOPERATION_LANGUAGENOTSUPPORT = "FailedOperation.LanguageNotSupport"

	// 入参CardSide错误，与实际上传的卡证正副页类型不符
	FAILEDOPERATION_LICENSECARDSIDEERROR = "FailedOperation.LicenseCardSideError"

	// 图片中存在两张及以上同面卡证，请上传卡证单面图片或一正一反双面图片
	FAILEDOPERATION_LICENSEMULTICARDERROR = "FailedOperation.LicenseMultiCardError"

	// 图片中存在两张及以上同面卡证，请上传卡证单面图片或一正一反双面图片。
	FAILEDOPERATION_MULTICARDERROR = "FailedOperation.MultiCardError"

	// 非银行卡。
	FAILEDOPERATION_NOBANKCARDERROR = "FailedOperation.NoBankCardError"

	// 非营业执照。
	FAILEDOPERATION_NOBIZLICENSE = "FailedOperation.NoBizLicense"

	// 上传的图片非驾驶证
	FAILEDOPERATION_NODRIVERLICENSEERROR = "FailedOperation.NoDriverLicenseError"

	// 非香港身份证。
	FAILEDOPERATION_NOHKIDCARD = "FailedOperation.NoHKIDCard"

	// 非马来身份证。
	FAILEDOPERATION_NOMASIDCARD = "FailedOperation.NoMASIDCard"

	// 非护照。
	FAILEDOPERATION_NOPASSPORT = "FailedOperation.NoPassport"

	// 上传的图片非行驶证
	FAILEDOPERATION_NOVEHICLELICENSEERROR = "FailedOperation.NoVehicleLicenseError"

	// OCR识别失败。
	FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"

	// 未知错误。
	FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"

	// 服务未开通。
	FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"

	// 剩余识别次数不足，请检查资源包状态。
	FAILEDOPERATION_USERQUOTAERROR = "FailedOperation.UserQuotaError"

	// 内部错误。
	INTERNALERROR = "InternalError"

	// Config不是有效的JSON格式。
	INVALIDPARAMETER_CONFIGFORMATERROR = "InvalidParameter.ConfigFormatError"

	// 图片解码失败。
	INVALIDPARAMETER_ENGINEIMAGEDECODEFAILED = "InvalidParameter.EngineImageDecodeFailed"

	// 任务创建失败，文件URL非法。
	INVALIDPARAMETERVALUE_FILEURLILLEGALERROR = "InvalidParameterValue.FileUrlIllegalError"

	// 参数值错误。
	INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"

	// 开票金额或校验码参数值错误。
	INVALIDPARAMETERVALUE_PRICEORVERIFICATIONPARAMETERVALUELIMIT = "InvalidParameterValue.PriceOrVerificationParameterValueLimit"

	// 发票代码参数值错误。
	INVALIDPARAMETERVALUE_TICKETCODEPARAMETERVALUELIMIT = "InvalidParameterValue.TicketCodeParameterValueLimit"

	// 开票日期参数值错误。
	INVALIDPARAMETERVALUE_TICKETDATEPARAMETERVALUELIMIT = "InvalidParameterValue.TicketDateParameterValueLimit"

	// 发票号码参数值错误。
	INVALIDPARAMETERVALUE_TICKETSNPARAMETERVALUELIMIT = "InvalidParameterValue.TicketSnParameterValueLimit"

	// 文件内容太大。
	LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"

	// 请求的次数超过了频率限制。
	REQUESTLIMITEXCEEDED = "RequestLimitExceeded"

	// 地区编码不存在。
	RESOURCENOTFOUND_NOAREACODE = "ResourceNotFound.NoAreaCode"

	// 发票不存在。温馨提示：新版发票核验接口功能更完整，请尽快切换，如已切换请忽略。
	RESOURCENOTFOUND_NOINVOICE = "ResourceNotFound.NoInvoice"

	// 不支持当天发票查询。
	RESOURCENOTFOUND_NOTSUPPORTCURRENTINVOICEQUERY = "ResourceNotFound.NotSupportCurrentInvoiceQuery"

	// 帐号已欠费。
	RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"

	// 账号资源包耗尽。
	RESOURCEUNAVAILABLE_RESOURCEPACKAGERUNOUT = "ResourceUnavailable.ResourcePackageRunOut"

	// 税务局网络异常，请稍后访问。
	RESOURCEUNAVAILABLE_TAXNETWORKERROR = "ResourceUnavailable.TaxNetworkError"

	// 计费状态异常。
	RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
)
