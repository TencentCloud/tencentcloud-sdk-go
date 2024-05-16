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

package v20240223

const (
	// 此产品的特有错误码

	// 音频解码失败。
	FAILEDOPERATION_AUDIODECODEFAILED = "FailedOperation.AudioDecodeFailed"

	// 音频处理失败。
	FAILEDOPERATION_AUDIOPROCESSFAILED = "FailedOperation.AudioProcessFailed"

	// 音频处理任务未完成，不能执行翻译结果确认
	FAILEDOPERATION_AUDIOPROCESSNOTFINISHED = "FailedOperation.AudioProcessNotFinished"

	// 文件下载失败。
	FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownloadError"

	// 任务不存在。
	FAILEDOPERATION_JOBNOTEXIST = "FailedOperation.JobNotExist"

	// 后端服务超时。
	FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"

	// 系统内部错误。
	FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"

	// 音频翻译结果已确认
	FAILEDOPERATION_TRANSLATIONCONFIRMHASFINISHED = "FailedOperation.TranslationConfirmHasFinished"

	// 用户未选择确认音频翻译结果
	FAILEDOPERATION_TRANSLATIONNOTNEEDCONFIRM = "FailedOperation.TranslationNotNeedConfirm"

	// 内部错误。
	FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"

	// 视频时长超限制。
	FAILEDOPERATION_VIDEODURATIONEXCEED = "FailedOperation.VideoDurationExceed"

	// 视频Fps超限制。
	FAILEDOPERATION_VIDEOFPSEXCEED = "FailedOperation.VideoFpsExceed"

	// 视频分辨率超限制。
	FAILEDOPERATION_VIDEORESOLUTIONEXCEED = "FailedOperation.VideoResolutionExceed"

	// 视频分辨率超限制。
	FAILEDOPERATION_VIDEOSIZEEXCEED = "FailedOperation.VideoSizeExceed"

	// 参数不合法。
	INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"

	// URL格式不合法。
	INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"

	// 当前并发数已超上限
	REQUESTLIMITEXCEEDED_JOBNUMEXCEED = "RequestLimitExceeded.JobNumExceed"

	// 用户账号超出了限制。
	REQUESTLIMITEXCEEDED_UINLIMITEXCEEDED = "RequestLimitExceeded.UinLimitExceeded"

	// 资源正在发货中。
	RESOURCEUNAVAILABLE_DELIVERING = "ResourceUnavailable.Delivering"

	// 账号已被冻结。
	RESOURCEUNAVAILABLE_FREEZE = "ResourceUnavailable.Freeze"

	// 账号已欠费。
	RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"

	// 服务正在开通中，请稍等。
	RESOURCEUNAVAILABLE_ISOPENING = "ResourceUnavailable.IsOpening"

	// 余额不足。
	RESOURCEUNAVAILABLE_LOWBALANCE = "ResourceUnavailable.LowBalance"

	// 计费状态未知，请确认是否已在控制台开通服务。
	RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"

	// 服务未开通。
	RESOURCEUNAVAILABLE_NOTREADY = "ResourceUnavailable.NotReady"

	// 资源已被回收。
	RESOURCEUNAVAILABLE_RECOVER = "ResourceUnavailable.Recover"

	// 计费状态未知。
	RESOURCEUNAVAILABLE_UNKNOWNSTATUS = "ResourceUnavailable.UnknownStatus"

	// 账号已欠费。
	RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
)
