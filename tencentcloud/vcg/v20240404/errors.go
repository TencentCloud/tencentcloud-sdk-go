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

package v20240404

const (
	// 此产品的特有错误码

	// 下载视频出错。
	FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownloadError"

	// 任务不存在。
	FAILEDOPERATION_TASKNOTEXIST = "FailedOperation.TaskNotExist"

	// 任务状态异常。
	FAILEDOPERATION_TASKSTATUSERROR = "FailedOperation.TaskStatusError"

	// 视频解码失败。
	FAILEDOPERATION_VIDEODECODEFAILED = "FailedOperation.VideoDecodeFailed"

	// 参数取值错误。
	INVALIDPARAMETERVALUE = "InvalidParameterValue"

	// 不支持的视频宽高比。
	INVALIDPARAMETERVALUE_INVALIDVIDEOASPECTRATIO = "InvalidParameterValue.InvalidVideoAspectRatio"

	// 视频时长超过限制。
	INVALIDPARAMETERVALUE_INVALIDVIDEODURATION = "InvalidParameterValue.InvalidVideoDuration"

	// 不支持的视频FPS。
	INVALIDPARAMETERVALUE_INVALIDVIDEOFPS = "InvalidParameterValue.InvalidVideoFPS"

	// 不支持的视频格式。
	INVALIDPARAMETERVALUE_INVALIDVIDEOFORMAT = "InvalidParameterValue.InvalidVideoFormat"

	// 不支持的分辨率。
	INVALIDPARAMETERVALUE_INVALIDVIDEORESOLUTION = "InvalidParameterValue.InvalidVideoResolution"

	// 参数字段或者值有误。
	INVALIDPARAMETERVALUE_PARAMETERVALUEERROR = "InvalidParameterValue.ParameterValueError"

	// 风格不存在。
	INVALIDPARAMETERVALUE_STYLENOTEXIST = "InvalidParameterValue.StyleNotExist"

	// 视频URL格式不合法。
	INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"

	// 视频大小超过限制。
	INVALIDPARAMETERVALUE_VIDEOSIZEEXCEED = "InvalidParameterValue.VideoSizeExceed"

	// 请求的次数超过了频率限制。
	REQUESTLIMITEXCEEDED = "RequestLimitExceeded"

	// 提交任务数超过最大并发。
	REQUESTLIMITEXCEEDED_JOBNUMEXCEED = "RequestLimitExceeded.JobNumExceed"

	// 资源不足。
	RESOURCEINSUFFICIENT = "ResourceInsufficient"

	// 计费状态未知。
	RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
)
