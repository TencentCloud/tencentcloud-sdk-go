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

package v20200910

const (
	// 此产品的特有错误码

	// CAM签名/鉴权错误。
	AUTHFAILURE = "AuthFailure"

	// 用量消息推送失败。
	FAILEDOPERATION_PUSHUSAGEMESSAGEERROR = "FailedOperation.PushUsageMessageError"

	// 服务未开通。
	FAILEDOPERATION_SERVICENOTOPEN = "FailedOperation.ServiceNotOpen"

	// 未知错误。
	FAILEDOPERATION_UNKNOWNERROR = "FailedOperation.UnknownError"

	// 内部错误。
	INTERNALERROR = "InternalError"

	// 图片ocr识别异常。
	INTERNALERROR_IMAGEOCRERROR = "InternalError.ImageOcrError"

	// 图片处理异常。
	INTERNALERROR_IMAGEPROCESSERROR = "InternalError.ImageProcessError"

	// 服务调用超时。
	INTERNALERROR_SERVERTIMEOUTERROR = "InternalError.ServerTimeOutError"

	// 报告结构化异常。
	INTERNALERROR_STRUCTIONERROR = "InternalError.StructionError"

	// 报告文本分类异常。
	INTERNALERROR_TEXTCLASSIFYERROR = "InternalError.TextClassifyError"

	// 参数错误。
	INVALIDPARAMETER = "InvalidParameter"

	// 参数AutoFitDirection无效。
	INVALIDPARAMETER_AUTOFITDIRECTION = "InvalidParameter.AutoFitDirection"

	// 参数ImageInfoList无效。
	INVALIDPARAMETER_IMAGEINFOLIST = "InvalidParameter.ImageInfoList"

	// 参数ImageOriginalSize无效。
	INVALIDPARAMETER_IMAGEORIGINALSIZE = "InvalidParameter.ImageOriginalSize"

	// 请求Action无效。
	INVALIDPARAMETER_INVALIDACTION = "InvalidParameter.InvalidAction"

	// 参数OcrEngineType无效。
	INVALIDPARAMETER_OCRENGINETYPE = "InvalidParameter.OcrEngineType"

	// 参数RotateTheAngle无效。
	INVALIDPARAMETER_ROTATETHEANGLE = "InvalidParameter.RotateTheAngle"

	// 参数Text无效。
	INVALIDPARAMETER_TEXT = "InvalidParameter.Text"

	// 图片编码无效。
	INVALIDPARAMETERVALUE_IMAGECODEINVALID = "InvalidParameterValue.ImageCodeInvalid"

	// 图片没有文字。
	INVALIDPARAMETERVALUE_IMAGEISNOTEXT = "InvalidParameterValue.ImageIsNoText"

	// 图片URL无效。
	INVALIDPARAMETERVALUE_IMAGEURLINVALID = "InvalidParameterValue.ImageURLInvalid"

	// 不支持的报告类型。
	OPERATIONDENIED_UNSUPPORTTHISTYPE = "OperationDenied.UnSupportThisType"

	// 当前无权限，请检查BisinsessId。
	UNAUTHORIZEDOPERATION_PERMISSIONDENIEDERROR = "UnauthorizedOperation.PermissionDeniedError"

	// 当前报告类型不支持。
	UNSUPPORTEDOPERATION_UNSUPPORTTHISTYPE = "UnsupportedOperation.UnSupportThisType"
)
