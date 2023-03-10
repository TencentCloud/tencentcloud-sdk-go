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

package v20221229

const (
	// 此产品的特有错误码

	// 控制台服务异常。
	FAILEDOPERATION_CONSOLESERVERERROR = "FailedOperation.ConsoleServerError"

	// 生成图片审核不通过，请重试。
	FAILEDOPERATION_GENERATEIMAGEFAILED = "FailedOperation.GenerateImageFailed"

	// 图片解码失败。
	FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"

	// 图片下载错误。
	FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"

	// 图片分辨率过大，超过2000*2000。
	FAILEDOPERATION_IMAGERESOLUTIONEXCEED = "FailedOperation.ImageResolutionExceed"

	// base64编码后的图片数据大小不超过10M。
	FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"

	// 整个请求体太大（通常主要是图片）。
	FAILEDOPERATION_REQUESTENTITYTOOLARGE = "FailedOperation.RequestEntityTooLarge"

	// 后端服务超时。
	FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"

	// RPC请求失败，一般为算法微服务故障。
	FAILEDOPERATION_RPCFAIL = "FailedOperation.RpcFail"

	// 服务内部错误。
	FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"

	// 未知错误。
	FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"

	// 参数不合法。
	INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"

	// 图片为空。
	INVALIDPARAMETERVALUE_IMAGEEMPTY = "InvalidParameterValue.ImageEmpty"

	// 参数字段或者值有误
	INVALIDPARAMETERVALUE_PARAMETERVALUEERROR = "InvalidParameterValue.ParameterValueError"

	// 1xx和其他风格不可混合使用。
	INVALIDPARAMETERVALUE_STYLECONFLICT = "InvalidParameterValue.StyleConflict"

	// 输入文本过长，请更换短一点的文本后重试。
	INVALIDPARAMETERVALUE_TEXTLENGTHEXCEED = "InvalidParameterValue.TextLengthExceed"

	// URL格式不合法。
	INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"

	// 图片包含非法信息，审核不通过。
	OPERATIONDENIED_IMAGEILLEGALDETECTED = "OperationDenied.ImageIllegalDetected"

	// 文本包含非法信息，审核不通过。
	OPERATIONDENIED_TEXTILLEGALDETECTED = "OperationDenied.TextIllegalDetected"

	// 请求的次数超过了频率限制。
	REQUESTLIMITEXCEEDED = "RequestLimitExceeded"

	// 资源正在发货中。
	RESOURCEUNAVAILABLE_DELIVERING = "ResourceUnavailable.Delivering"

	// 帐号已欠费。
	RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"

	// 余额不足。
	RESOURCEUNAVAILABLE_LOWBALANCE = "ResourceUnavailable.LowBalance"

	// 计费状态未知，请确认是否已在控制台开通服务。
	RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"

	// 帐号已停服。
	RESOURCEUNAVAILABLE_STOPUSING = "ResourceUnavailable.StopUsing"

	// 计费状态异常。
	RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
)
