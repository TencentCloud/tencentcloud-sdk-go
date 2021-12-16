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

package v20181213

const (
	// 此产品的特有错误码

	// 内部错误。
	INTERNALERROR = "InternalError"

	// 批改错误。
	INTERNALERROR_CORRECTERROR = "InternalError.CorrectError"

	// 识别错误。
	INTERNALERROR_OCRERROR = "InternalError.OcrError"

	// 服务器内部错误，初始化失败。
	INTERNALERROR_OCRSERVERINTERNERROR = "InternalError.OcrServerInternError"

	// 其它错误。
	INTERNALERROR_OTHERERROR = "InternalError.OtherError"

	// 服务器过载，请联系相关客服。
	INTERNALERROR_OVERLOADERROR = "InternalError.OverLoadError"

	// 图片识别错误。
	INTERNALERROR_RECOGNIZEERROR = "InternalError.RecognizeError"

	// 无法连接图像下载服务器。
	INTERNALERROR_SERVERCONNECTDOWNLOADERROR = "InternalError.ServerConnectDownloadError"

	// 图片切割错误。
	INTERNALERROR_SPLITERROR = "InternalError.SplitError"

	// 参数为空。
	INVALIDPARAMETER_EMPTYPARAMETERERROR = "InvalidParameter.EmptyParameterError"

	// 传入的参数有误。
	INVALIDPARAMETER_INPUTERROR = "InvalidParameter.InputError"

	// 任务不存在。
	INVALIDPARAMETER_TASKNOTFOUND = "InvalidParameter.TaskNotFound"

	// Appid无效。
	INVALIDPARAMETERVALUE_APPIDINVALIDERROR = "InvalidParameterValue.AppidInvalidError"

	// 图片解码失败，请核实输入信息。
	INVALIDPARAMETERVALUE_DECODEIMAGEERROR = "InvalidParameterValue.DecodeImageError"

	// url图片下载失败。
	INVALIDPARAMETERVALUE_DOWNLOADIMAGEFAILERROR = "InvalidParameterValue.DownloadImageFailError"

	// 图片数据为空。
	INVALIDPARAMETERVALUE_EMPTYIMAGEERROR = "InvalidParameterValue.EmptyImageError"

	// 图片下载失败。
	INVALIDPARAMETERVALUE_IMAGEDOWNLOADFAILERROR = "InvalidParameterValue.ImageDownloadFailError"

	// 图片超出下载限制。
	INVALIDPARAMETERVALUE_IMAGESIZEEXCEEDERROR = "InvalidParameterValue.ImageSizeExceedError"

	// 图片尺寸太大。
	INVALIDPARAMETERVALUE_IMAGETOOBIGERROR = "InvalidParameterValue.ImageTooBigError"

	// 输入错误，请核实InputType参数。
	INVALIDPARAMETERVALUE_INPUTTYPEVALUEERROR = "InvalidParameterValue.InputTypeValueError"

	// SessionId无效。
	INVALIDPARAMETERVALUE_SESSIONERROR = "InvalidParameterValue.SessionError"

	// 图像请求URL的格式错误。
	INVALIDPARAMETERVALUE_URLFROMATIVADLIDERROR = "InvalidParameterValue.UrlFromatIvadlidError"

	// 频率限制。
	LIMITEXCEEDED_FREQLIMITFORBIDDENACCESSERROR = "LimitExceeded.FreqLimitForbiddenAccessError"

	// 无法找到用户，请确认已在控制台开通服务并使用了正确的 ECCAPPID。
	RESOURCENOTFOUND_CANNOTFINDUSER = "ResourceNotFound.CannotFindUser"

	// 无效的服务名称。
	RESOURCENOTFOUND_SERVERNAMENOTEXISTINLICENSEERROR = "ResourceNotFound.ServerNameNotExistInLicenseError"

	// 服务未开通或已欠费。
	RESOURCEUNAVAILABLE_AUTHORIZEERROR = "ResourceUnavailable.AuthorizeError"

	// license无效。
	UNAUTHORIZEDOPERATION_LICENSEINVALIDFORBIDDENACCESSERROR = "UnauthorizedOperation.LicenseInvalidForbiddenAccessError"

	// license中未授权该服务。
	UNAUTHORIZEDOPERATION_SERVERNAMEUNAUTHORIZEDINERROR = "UnauthorizedOperation.ServerNameUnauthorizedInError"
)
