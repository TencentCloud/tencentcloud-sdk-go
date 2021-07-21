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

package v20190919

const (
	// 此产品的特有错误码

	// CAM签名/鉴权错误。
	AUTHFAILURE = "AuthFailure"

	// 操作失败。
	FAILEDOPERATION = "FailedOperation"

	// 文档下载失败，请检查请求参数中URL是否正确，或者如果您使用其他的文件存储服务，请检查文件存储服务的上传带宽，文档转码服务仅提供1分钟的下载时间，如果下载不成功本次的转码请求将以失败终止。
	FAILEDOPERATION_FILEDOWNLOADFAIL = "FailedOperation.FileDownloadFail"

	// 文档格式错误，不支持转换只读文档或者已加密的文档。
	FAILEDOPERATION_FILEFORMATERROR = "FailedOperation.FileFormatError"

	// 文档打开失败，请检查提交转码的文档是否加密或有其他格式问题。
	FAILEDOPERATION_FILEOPENFAIL = "FailedOperation.FileOpenFail"

	// 转码后上传结果失败，请稍候重试。
	FAILEDOPERATION_FILEUPLOADFAIL = "FailedOperation.FileUploadFail"

	// 录制失败，具体请参考错误描述。
	FAILEDOPERATION_RECORD = "FailedOperation.Record"

	// 转码失败，具体请参考错误描述或联系客服人员。
	FAILEDOPERATION_TRANSCODE = "FailedOperation.Transcode"

	// 转码服务出现内部错误，请稍候重试或联系客户人员。
	FAILEDOPERATION_TRANSCODESERVERERROR = "FailedOperation.TranscodeServerError"

	// 白板推流失败，具体请参考错误描述。
	FAILEDOPERATION_WHITEBOARDPUSH = "FailedOperation.WhiteboardPush"

	// 内部错误。
	INTERNALERROR = "InternalError"

	// 参数错误。
	INVALIDPARAMETER = "InvalidParameter"

	// 参数类型不匹配。
	INVALIDPARAMETER_BODYPARAMETERTYPEUNMATCHED = "InvalidParameter.BodyParameterTypeUnmatched"

	// 回调地址格式错误。
	INVALIDPARAMETER_CALLBACKADDRESSFORMATERROR = "InvalidParameter.CallbackAddressFormatError"

	// 文档后缀名对应的格式不支持。
	INVALIDPARAMETER_FILEFORMATUNSUPPORTED = "InvalidParameter.FileFormatUnsupported"

	// 额外指定的特殊功能不存在。
	INVALIDPARAMETER_INVALIDEXTRA = "InvalidParameter.InvalidExtra"

	// 实时录制参数格式不正确。
	INVALIDPARAMETER_RECORDPARAMETER = "InvalidParameter.RecordParameter"

	// SdkAppId不存在或格式错误。
	INVALIDPARAMETER_SDKAPPIDNOTFOUND = "InvalidParameter.SdkAppIdNotFound"

	// 需要查询的任务不存在。
	INVALIDPARAMETER_TASKNOTFOUND = "InvalidParameter.TaskNotFound"

	// 文档转码参数格式不正确。
	INVALIDPARAMETER_TRANSCODEPARAMETER = "InvalidParameter.TranscodeParameter"

	// 文档下载Url格式错误，请检查请求参数里的Url。
	INVALIDPARAMETER_URLFORMATERROR = "InvalidParameter.UrlFormatError"

	// 转码或录制任务并发数量超过限制，请参考错误描述或稍后重试。
	LIMITEXCEEDED_TASKCONCURRENCY = "LimitExceeded.TaskConcurrency"

	// 超过文档最大页数限制，目前不支持超过500页的文件转码，如有特殊需求，请联系客服人员。
	LIMITEXCEEDED_TRANSCODEPAGESLIMITATION = "LimitExceeded.TranscodePagesLimitation"

	// 请求的次数超过了频率限制。
	REQUESTLIMITEXCEEDED = "RequestLimitExceeded"

	// 实时录制任务录制用户已被其他录制任务使用。
	RESOURCEINUSE_RECORDUSERID = "ResourceInUse.RecordUserId"

	// 资源不可用。
	RESOURCEUNAVAILABLE = "ResourceUnavailable"

	// 未开通互动白板。
	RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"

	// 账户欠费或者互动白板服务已过期。
	RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"

	// 未授权操作。
	UNAUTHORIZEDOPERATION = "UnauthorizedOperation"

	// SdkAppId不存在或者SdkAppId与当前腾讯云账号不对应。
	UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"

	// 未知参数错误。
	UNKNOWNPARAMETER = "UnknownParameter"

	// 操作不支持。
	UNSUPPORTEDOPERATION = "UnsupportedOperation"

	// 当前未完成的任务不能此状态下执行指定操作，例如对正在录制的任务执行恢复录制等。
	UNSUPPORTEDOPERATION_INVALIDTASKSTATUS = "UnsupportedOperation.InvalidTaskStatus"

	// 任务结束，不能执行指定操作。
	UNSUPPORTEDOPERATION_TASKHASALREADYSTOPPED = "UnsupportedOperation.TaskHasAlreadyStopped"
)
