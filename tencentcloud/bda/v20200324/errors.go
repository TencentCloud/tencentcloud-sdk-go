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

package v20200324

const (
	// 此产品的特有错误码

	// 音频解码失败。
	FAILEDOPERATION_AUDIODECODEFAILED = "FailedOperation.AudioDecodeFailed"

	// 音频编码失败。
	FAILEDOPERATION_AUDIOENCODEFAILED = "FailedOperation.AudioEncodeFailed"

	// 人体特征检测失败。
	FAILEDOPERATION_BODYFEATUREFAIL = "FailedOperation.BodyFeatureFail"

	// 人体关键点检测失败。
	FAILEDOPERATION_BODYJOINTSFAIL = "FailedOperation.BodyJointsFail"

	// 人体质量分过低。
	FAILEDOPERATION_BODYQUALITYNOTQUALIFIED = "FailedOperation.BodyQualityNotQualified"

	// 输入的人体框不合法。
	FAILEDOPERATION_BODYRECTILLEGAL = "FailedOperation.BodyRectIllegal"

	// 输入的人体框数量不合法。
	FAILEDOPERATION_BODYRECTNUMILLEGAL = "FailedOperation.BodyRectNumIllegal"

	// 创建轨迹失败，请选择符合要求的人体图片。
	FAILEDOPERATION_CREATETRACEFAILED = "FailedOperation.CreateTraceFailed"

	// 文件下载失败。
	FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownloadError"

	// 搜索的人体库为空。
	FAILEDOPERATION_GROUPEMPTY = "FailedOperation.GroupEmpty"

	// 人体检测失败。
	FAILEDOPERATION_IMAGEBODYDETECTFAILED = "FailedOperation.ImageBodyDetectFailed"

	// 图片解码失败。
	FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"

	// 图片下载错误。
	FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"

	// 人脸检测失败。
	FAILEDOPERATION_IMAGEFACEDETECTFAILED = "FailedOperation.ImageFacedetectFailed"

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

	// 图片中没有人体。
	FAILEDOPERATION_NOBODYINPHOTO = "FailedOperation.NoBodyInPhoto"

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

	// 未知错误。
	FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"

	// 视频解码失败。
	FAILEDOPERATION_VIDEODECODEFAILED = "FailedOperation.VideoDecodeFailed"

	// 参数不合法。
	INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"

	// 账号人体轨迹数量超出限制。
	INVALIDPARAMETERVALUE_ACCOUNTTRACENUMEXCEED = "InvalidParameterValue.AccountTraceNumExceed"

	// 算法模型版本不合法。
	INVALIDPARAMETERVALUE_BODYMODELVERSIONILLEGAL = "InvalidParameterValue.BodyModelVersionIllegal"

	// 传入的人体框过多。
	INVALIDPARAMETERVALUE_BODYRECTSEXCEED = "InvalidParameterValue.BodyRectsExceed"

	// 人体库ID已经存在。人体库ID不可重复。
	INVALIDPARAMETERVALUE_GROUPIDALREADYEXIST = "InvalidParameterValue.GroupIdAlreadyExist"

	// 人体库ID包含非法字符。人体库ID只支持英文、数字、-%@#&_。
	INVALIDPARAMETERVALUE_GROUPIDILLEGAL = "InvalidParameterValue.GroupIdIllegal"

	// 人体库ID不存在。
	INVALIDPARAMETERVALUE_GROUPIDNOTEXIST = "InvalidParameterValue.GroupIdNotExist"

	// 人体库ID超出长度限制。
	INVALIDPARAMETERVALUE_GROUPIDTOOLONG = "InvalidParameterValue.GroupIdTooLong"

	// 人体库名称已经存在。人体库名称不可重复。
	INVALIDPARAMETERVALUE_GROUPNAMEALREADYEXIST = "InvalidParameterValue.GroupNameAlreadyExist"

	// 人体库名称超出长度限制。
	INVALIDPARAMETERVALUE_GROUPNAMETOOLONG = "InvalidParameterValue.GroupNameTooLong"

	// 人体库数量超出限制。
	INVALIDPARAMETERVALUE_GROUPNUMEXCEED = "InvalidParameterValue.GroupNumExceed"

	// 人体库备注超出长度限制。
	INVALIDPARAMETERVALUE_GROUPTAGTOOLONG = "InvalidParameterValue.GroupTagTooLong"

	// 人体库人体轨迹数量超出限制。
	INVALIDPARAMETERVALUE_GROUPTRACENUMEXCEED = "InvalidParameterValue.GroupTraceNumExceed"

	// 图片为空。
	INVALIDPARAMETERVALUE_IMAGEEMPTY = "InvalidParameterValue.ImageEmpty"

	// 返回数量不在合法范围内。
	INVALIDPARAMETERVALUE_LIMITEXCEED = "InvalidParameterValue.LimitExceed"

	// 图片中没有人脸。
	INVALIDPARAMETERVALUE_NOFACEINPHOTO = "InvalidParameterValue.NoFaceInPhoto"

	// 起始序号过大。请检查需要请求的数组长度。
	INVALIDPARAMETERVALUE_OFFSETEXCEED = "InvalidParameterValue.OffsetExceed"

	// 人员ID已经存在。人员ID不可重复。
	INVALIDPARAMETERVALUE_PERSONIDALREADYEXIST = "InvalidParameterValue.PersonIdAlreadyExist"

	// 人员ID包含非法字符。人员ID只支持英文、数字、-%@#&_。
	INVALIDPARAMETERVALUE_PERSONIDILLEGAL = "InvalidParameterValue.PersonIdIllegal"

	// 人员ID不存在。
	INVALIDPARAMETERVALUE_PERSONIDNOTEXIST = "InvalidParameterValue.PersonIdNotExist"

	// 人员ID超出长度限制。
	INVALIDPARAMETERVALUE_PERSONIDTOOLONG = "InvalidParameterValue.PersonIdTooLong"

	// 人员名称包含非法字符。
	INVALIDPARAMETERVALUE_PERSONNAMEILLEGAL = "InvalidParameterValue.PersonNameIllegal"

	// 人员名称超出长度限制。
	INVALIDPARAMETERVALUE_PERSONNAMETOOLONG = "InvalidParameterValue.PersonNameTooLong"

	// 人员人体轨迹数量超出限制。
	INVALIDPARAMETERVALUE_PERSONTRACENUMEXCEED = "InvalidParameterValue.PersonTraceNumExceed"

	// 搜索的人员数目超过限制。
	INVALIDPARAMETERVALUE_SEARCHPERSONSEXCEED = "InvalidParameterValue.SearchPersonsExceed"

	// 创建人体轨迹的人体图片数量超出限制。
	INVALIDPARAMETERVALUE_TRACEBODYNUMEXCEED = "InvalidParameterValue.TraceBodyNumExceed"

	// TraceMatchThreshold参数不合法。
	INVALIDPARAMETERVALUE_TRACEMATCHTHRESHOLDILLEGAL = "InvalidParameterValue.TraceMatchThresholdIllegal"

	// URL格式不合法。
	INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"

	// 文件太大。
	LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"

	// 必选参数为空。
	MISSINGPARAMETER_ERRORPARAMETEREMPTY = "MissingParameter.ErrorParameterEmpty"

	// 账号已欠费。
	RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"

	// 计费状态未知，请确认是否已在控制台开通服务。
	RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"

	// 服务未开通。
	RESOURCEUNAVAILABLE_NOTREADY = "ResourceUnavailable.NotReady"

	// 未知方法名。
	UNSUPPORTEDOPERATION_UNKNOWMETHOD = "UnsupportedOperation.UnknowMethod"
)
