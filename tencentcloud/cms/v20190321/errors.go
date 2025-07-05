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

package v20190321

const (
	// 此产品的特有错误码

	// CAM签名/鉴权错误。
	AUTHFAILURE = "AuthFailure"

	// 操作失败。
	FAILEDOPERATION = "FailedOperation"

	// 内部错误
	INTERNALERROR = "InternalError"

	// 服务内部错误。
	INTERNALERROR_INTERNALERROR = "InternalError.InternalError"

	// 参数错误。
	INVALIDPARAMETER = "InvalidParameter"

	// 图片长宽比太大
	INVALIDPARAMETER_IMAGEASPECTRATIOTOOLARGE = "InvalidParameter.ImageAspectRatioTooLarge"

	// 图片体积太小
	INVALIDPARAMETER_IMAGEDATATOOSMALL = "InvalidParameter.ImageDataTooSmall"

	// 图片尺寸过小。
	INVALIDPARAMETER_IMAGESIZETOOSMALL = "InvalidParameter.ImageSizeTooSmall"

	// 图片内容错误。
	INVALIDPARAMETER_INVALIDIMAGECONTENT = "InvalidParameter.InvalidImageContent"

	// 参数不可用
	INVALIDPARAMETER_PARAMETERERROR = "InvalidParameter.ParameterError"

	// FileContent不可用：需Base64编码
	INVALIDPARAMETERVALUE_ERRFILECONTENT = "InvalidParameterValue.ErrFileContent"

	// 图片尺寸错误。
	INVALIDPARAMETERVALUE_ERRIMAGESIZE = "InvalidParameterValue.ErrImageSize"

	// 文本内容过长。
	INVALIDPARAMETERVALUE_ERRTEXTCONTENTLEN = "InvalidParameterValue.ErrTextContentLen"

	// 文本内容类型错误：需base64编码
	INVALIDPARAMETERVALUE_ERRTEXTCONTENTTYPE = "InvalidParameterValue.ErrTextContentType"

	// Content参数错误
	INVALIDPARAMETERVALUE_INVALIDCONTENT = "InvalidParameterValue.InvalidContent"

	// 图片文件内容大小异常。
	INVALIDPARAMETERVALUE_INVALIDFILECONTENTSIZE = "InvalidParameterValue.InvalidFileContentSize"

	// 图片内容错误
	INVALIDPARAMETERVALUE_INVALIDIMAGECONTENT = "InvalidParameterValue.InvalidImageContent"

	// 参数取值错误
	INVALIDPARAMETERVALUE_INVALIDPARAMETER = "InvalidParameterValue.InvalidParameter"

	// 超过配额限制。
	LIMITEXCEEDED = "LimitExceeded"

	// 缺少参数错误。
	MISSINGPARAMETER = "MissingParameter"

	// FileUrl或FileContent都为空
	MISSINGPARAMETER_ERRFILEURL = "MissingParameter.ErrFileUrl"

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

	// 文件链接下载内部错误
	RESOURCENOTFOUND_ERRDOWDOWNINTERNALERROR = "ResourceNotFound.ErrDowdownInternalError"

	// 文件链接下载服务参数异常
	RESOURCENOTFOUND_ERRDOWDOWNPARAMSERROR = "ResourceNotFound.ErrDowdownParamsError"

	// 文件链接下载源错误
	RESOURCENOTFOUND_ERRDOWDOWNSOURCEERROR = "ResourceNotFound.ErrDowdownSourceError"

	// 文件链接下载超时
	RESOURCENOTFOUND_ERRDOWDOWNTIMEOUT = "ResourceNotFound.ErrDowdownTimeOut"

	// 资源不可用。
	RESOURCEUNAVAILABLE = "ResourceUnavailable"

	// 图片识别服务超时
	RESOURCEUNAVAILABLE_ERRIMAGETIMEOUT = "ResourceUnavailable.ErrImageTimeOut"

	// 文本识别服务超时
	RESOURCEUNAVAILABLE_ERRTEXTTIMEOUT = "ResourceUnavailable.ErrTextTimeOut"

	// 图片文件下载失败。
	RESOURCEUNAVAILABLE_IMAGEDOWNLOADERROR = "ResourceUnavailable.ImageDownloadError"

	// 未授权操作。
	UNAUTHORIZEDOPERATION = "UnauthorizedOperation"

	// 鉴权失败
	UNAUTHORIZEDOPERATION_ERRAUTH = "UnauthorizedOperation.ErrAuth"

	// 未开通权限/无有效套餐包/账号已欠费。
	UNAUTHORIZEDOPERATION_UNAUTHORIZED = "UnauthorizedOperation.Unauthorized"

	// 未知参数错误。
	UNKNOWNPARAMETER = "UnknownParameter"

	// 操作不支持。
	UNSUPPORTEDOPERATION = "UnsupportedOperation"
)
