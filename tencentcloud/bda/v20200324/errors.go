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

package v20200324

const (
	// 此产品的特有错误码

	// 音频解码失败。
	FAILEDOPERATION_AUDIODECODEFAILED = "FailedOperation.AudioDecodeFailed"

	// 音频编码失败。
	FAILEDOPERATION_AUDIOENCODEFAILED = "FailedOperation.AudioEncodeFailed"

	// 余额不足，开通失败，请充值后再开通。
	FAILEDOPERATION_BALANCEINSUFFICIENT = "FailedOperation.BalanceInsufficient"

	// 文件下载失败。
	FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownloadError"

	// 图片解码失败。
	FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"

	// 图片下载错误。
	FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"

	// 人脸检测失败。
	FAILEDOPERATION_IMAGEFACEDETECTFAILED = "FailedOperation.ImageFacedetectFailed"

	// 图片不存在前景。
	FAILEDOPERATION_IMAGENOTFOREGROUND = "FailedOperation.ImageNotForeground"

	// 不支持的图片文件。
	FAILEDOPERATION_IMAGENOTSUPPORTED = "FailedOperation.ImageNotSupported"

	// 图片分辨率过大。
	FAILEDOPERATION_IMAGERESOLUTIONEXCEED = "FailedOperation.ImageResolutionExceed"

	// 图片分辨率过小。
	FAILEDOPERATION_IMAGERESOLUTIONINSUFFICIENT = "FailedOperation.ImageResolutionInsufficient"

	// base64编码后的图片数据过大。
	FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"

	// 服务内部错误，请重试。
	FAILEDOPERATION_INNERERROR = "FailedOperation.InnerError"

	// 任务冲突。
	FAILEDOPERATION_JOBCONFLICT = "FailedOperation.JobConflict"

	// 任务队列已满。
	FAILEDOPERATION_JOBQUEUEFULL = "FailedOperation.JobQueueFull"

	// 人像数过多。
	FAILEDOPERATION_PROFILENUMEXCEED = "FailedOperation.ProfileNumExceed"

	// 整个请求体太大（通常主要是图片）。
	FAILEDOPERATION_REQUESTENTITYTOOLARGE = "FailedOperation.RequestEntityTooLarge"

	// 后端服务超时。
	FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"

	// RPC请求失败，一般为算法微服务故障。
	FAILEDOPERATION_RPCFAIL = "FailedOperation.RpcFail"

	// 人像分割失败。
	FAILEDOPERATION_SEGMENTFAILED = "FailedOperation.SegmentFailed"

	// 算法服务异常，请重试。
	FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"

	// 任务超过上限。
	FAILEDOPERATION_TASKLIMITEXCEEDED = "FailedOperation.TaskLimitExceeded"

	// 任务不存在。
	FAILEDOPERATION_TASKNOTEXIST = "FailedOperation.TaskNotExist"

	// 任务取消失败。
	FAILEDOPERATION_TERMINATETASKFAILED = "FailedOperation.TerminateTaskFailed"

	// 文件太大。
	FAILEDOPERATION_TOOLARGEFILEERROR = "FailedOperation.TooLargeFileError"

	// 内部错误。
	FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"

	// 视频解码失败。
	FAILEDOPERATION_VIDEODECODEFAILED = "FailedOperation.VideoDecodeFailed"

	// 参数不合法。
	INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"

	// 图片为空。
	INVALIDPARAMETERVALUE_IMAGEEMPTY = "InvalidParameterValue.ImageEmpty"

	// 图片中没有人脸。
	INVALIDPARAMETERVALUE_NOFACEINPHOTO = "InvalidParameterValue.NoFaceInPhoto"

	// URL格式不合法。
	INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"

	// 文件太大。
	LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"

	// 账号已欠费。
	RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"

	// 计费状态未知，请确认是否已在控制台开通服务。
	RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"

	// 未知方法名。
	UNSUPPORTEDOPERATION_UNKNOWMETHOD = "UnsupportedOperation.UnknowMethod"
)
