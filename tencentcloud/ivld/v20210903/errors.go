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

package v20210903

const (
	// 此产品的特有错误码

	// 用户不存在。
	AUTHFAILURE_USERNOTFOUND = "AuthFailure.UserNotFound"

	// 匹配的模板不存在。
	FAILEDOPERATION_AITEMPLATENOTEXIST = "FailedOperation.AiTemplateNotExist"

	// 内部DB连接失败。
	FAILEDOPERATION_DBCONNECTIONERROR = "FailedOperation.DBConnectionError"

	// 内部DB操作错误。
	FAILEDOPERATION_DBOPERATIONERROR = "FailedOperation.DBOperationError"

	// 媒资文件下载失败。
	FAILEDOPERATION_DOWNLOADFAILED = "FailedOperation.DownloadFailed"

	// 获取任务列表失败。
	FAILEDOPERATION_GETTASKLISTFAILED = "FailedOperation.GetTaskListFailed"

	// 获取媒资信息失败。
	FAILEDOPERATION_GETVIDEOMETADATAFAILED = "FailedOperation.GetVideoMetadataFailed"

	// MD5不匹配。
	FAILEDOPERATION_MD5MISMATCH = "FailedOperation.MD5Mismatch"

	// 媒资文件已经存在。
	FAILEDOPERATION_MEDIAALREADYEXIST = "FailedOperation.MediaAlreadyExist"

	// 媒资正在使用。
	FAILEDOPERATION_MEDIAINUSE = "FailedOperation.MediaInUse"

	// 媒体文件未就绪。
	FAILEDOPERATION_MEDIANOTREADY = "FailedOperation.MediaNotReady"

	// 存在相同的任务。
	FAILEDOPERATION_TASKALREADYEXIST = "FailedOperation.TaskAlreadyExist"

	// 视频分析未完成。
	FAILEDOPERATION_TASKNOTFINISHED = "FailedOperation.TaskNotFinished"

	// 转码失败。
	FAILEDOPERATION_TRANSCODEFAILED = "FailedOperation.TranscodeFailed"

	// 内部错误。
	INTERNALERROR_INNERERROR = "InternalError.InnerError"

	// 参数错误。
	INVALIDPARAMETER = "InvalidParameter"

	// 文件路径不合法。
	INVALIDPARAMETER_INVALIDFILEPATH = "InvalidParameter.InvalidFilePath"

	// MD5非法。
	INVALIDPARAMETER_INVALIDMD5 = "InvalidParameter.InvalidMD5"

	// 媒体ID不合法。
	INVALIDPARAMETER_INVALIDMEDIAID = "InvalidParameter.InvalidMediaId"

	// 媒体名称非法。
	INVALIDPARAMETER_INVALIDMEDIANAME = "InvalidParameter.InvalidMediaName"

	// 媒资状态不合法。
	INVALIDPARAMETER_INVALIDMEDIASTATUS = "InvalidParameter.InvalidMediaStatus"

	// 名称不合法。
	INVALIDPARAMETER_INVALIDNAME = "InvalidParameter.InvalidName"

	// 分页序号不合法。
	INVALIDPARAMETER_INVALIDPAGENUMBER = "InvalidParameter.InvalidPageNumber"

	// 分页大小不合法。
	INVALIDPARAMETER_INVALIDPAGESIZE = "InvalidParameter.InvalidPageSize"

	// 排序字段不合法。
	INVALIDPARAMETER_INVALIDSORTBY = "InvalidParameter.InvalidSortBy"

	// 排序方式不合法。
	INVALIDPARAMETER_INVALIDSORTORDER = "InvalidParameter.InvalidSortOrder"

	// 任务ID不合法。
	INVALIDPARAMETER_INVALIDTASKID = "InvalidParameter.InvalidTaskId"

	// 任务名称不合法。
	INVALIDPARAMETER_INVALIDTASKNAME = "InvalidParameter.InvalidTaskName"

	// 任务状态不合法。
	INVALIDPARAMETER_INVALIDTASKSTATUS = "InvalidParameter.InvalidTaskStatus"

	// URL不合法。
	INVALIDPARAMETER_INVALIDURL = "InvalidParameter.InvalidURL"

	// 用户Uin无效。
	INVALIDPARAMETER_INVALIDUIN = "InvalidParameter.InvalidUin"

	// 名称超过长度限制。
	INVALIDPARAMETER_NAMETOOLONG = "InvalidParameter.NameTooLong"

	// 不支持的URL类型。
	INVALIDPARAMETER_UNSUPPORTURL = "InvalidParameter.UnsupportURL"

	// 使用量超过限制。
	LIMITEXCEEDED_USAGELIMITEXCEEDED = "LimitExceeded.UsageLimitExceeded"

	// 同时发起过多任务。
	REQUESTLIMITEXCEEDED_CONCURRENCYOVERFLOW = "RequestLimitExceeded.ConcurrencyOverflow"

	// 媒资文件不存在。
	RESOURCENOTFOUND_MEDIANOTFOUND = "ResourceNotFound.MediaNotFound"

	// 任务不存在。
	RESOURCENOTFOUND_TASKNOTFOUND = "ResourceNotFound.TaskNotFound"

	// 未授权操作。
	UNAUTHORIZEDOPERATION = "UnauthorizedOperation"

	// 用户未激活该产品。
	UNAUTHORIZEDOPERATION_UNAUTHORIZEDPRODUCT = "UnauthorizedOperation.UnauthorizedProduct"

	// 媒资文件不可访问。
	UNSUPPORTEDOPERATION_MEDIANOTACCESSIBLE = "UnsupportedOperation.MediaNotAccessible"
)
