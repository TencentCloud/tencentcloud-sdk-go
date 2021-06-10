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

package v20200304

const (
	// 此产品的特有错误码

	// CAM签名/鉴权错误。
	AUTHFAILURE = "AuthFailure"

	// 操作失败。
	FAILEDOPERATION = "FailedOperation"

	// Cos存储结果错误。
	FAILEDOPERATION_COSSTORAGEERROR = "FailedOperation.CosStorageError"

	// 转码时截取失败。
	FAILEDOPERATION_EDITERROR = "FailedOperation.EditError"

	// 编码格式或参数不支持。
	FAILEDOPERATION_ENCODEFORMATERROR = "FailedOperation.EncodeFormatError"

	// 运行中的任务达到上限，如需要增加任务上限，请提交工单。
	FAILEDOPERATION_RUNNINGTASKEXCEED = "FailedOperation.RunningTaskExceed"

	// 转码后切片失败。
	FAILEDOPERATION_SEGMENTERROR = "FailedOperation.SegmentError"

	// 系统繁忙，请稍后重试。
	FAILEDOPERATION_SERVERBUSY = "FailedOperation.ServerBusy"

	// 内部服务错误。
	FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"

	// 转码服务异常。
	FAILEDOPERATION_TRANSCODEERROR = "FailedOperation.TranscodeError"

	// 转码服务未知错误。
	FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnknowError"

	// 视频源下载失败或超时。
	FAILEDOPERATION_VIDEODOWNLOADERROR = "FailedOperation.VideoDownloadError"

	// 视频源解析出错。
	FAILEDOPERATION_VIDEOPARSEERROR = "FailedOperation.VideoParseError"

	// 视频大小超过限制。
	FAILEDOPERATION_VIDEOSIZEEXCEED = "FailedOperation.VideoSizeExceed"

	// 内部错误。
	INTERNALERROR = "InternalError"

	// 参数错误。
	INVALIDPARAMETER = "InvalidParameter"

	// 参数取值错误。
	INVALIDPARAMETERVALUE = "InvalidParameterValue"

	// 任务不支持该操作。
	INVALIDPARAMETERVALUE_ACTIONNOTSUPPORT = "InvalidParameterValue.ActionNotSupport"

	// CallbackUrl 不安全。
	INVALIDPARAMETERVALUE_CALLBACKURLERROR = "InvalidParameterValue.CallbackUrlError"

	// 直播场景回调地址必选。
	INVALIDPARAMETERVALUE_CALLBACKURLNOTEXIST = "InvalidParameterValue.CallbackUrlNotExist"

	// Cos授权信息错误。
	INVALIDPARAMETERVALUE_COSAUTHMODEERROR = "InvalidParameterValue.CosAuthModeError"

	// Cos托管ID不存在。
	INVALIDPARAMETERVALUE_COSHOSTEDIDNOTEXIST = "InvalidParameterValue.CosHostedIdNotExist"

	// 视频源地址格式错误。
	INVALIDPARAMETERVALUE_DOWNINFOFORMATWRONG = "InvalidParameterValue.DownInfoFormatWrong"

	// 视频源下载类型错误。
	INVALIDPARAMETERVALUE_DOWNINFOTYPEWRONG = "InvalidParameterValue.DownInfoTypeWrong"

	// 该任务不支持直播流。
	INVALIDPARAMETERVALUE_LIVESOURCENOTSUPPORT = "InvalidParameterValue.LiveSourceNotSupport"

	// 任务已经结束。
	INVALIDPARAMETERVALUE_TASKALREADYDONE = "InvalidParameterValue.TaskAlreadyDone"

	// 任务已经删除。
	INVALIDPARAMETERVALUE_TASKDELETED = "InvalidParameterValue.TaskDeleted"

	// 任务ID不存在。
	INVALIDPARAMETERVALUE_TASKIDNOTEXIST = "InvalidParameterValue.TaskIdNotExist"

	// 请求uri错误。
	INVALIDPARAMETERVALUE_URIERROR = "InvalidParameterValue.UriError"

	// URL 不安全。
	INVALIDPARAMETERVALUE_URLINFOURLERROR = "InvalidParameterValue.UrlInfoUrlError"

	// 视频格式不支持。
	INVALIDPARAMETERVALUE_VIDEOFORMATERROR = "InvalidParameterValue.VideoFormatError"

	// 请求的次数超过了频率限制。
	REQUESTLIMITEXCEEDED = "RequestLimitExceeded"

	// 帐号已欠费。
	RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"

	// 未授权操作。
	UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
)
