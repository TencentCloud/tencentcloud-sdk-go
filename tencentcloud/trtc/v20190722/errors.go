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

package v20190722

const (
	// 此产品的特有错误码

	// CAM签名/鉴权错误。
	AUTHFAILURE = "AuthFailure"

	// 未完成实名认证，不允许此操作。
	AUTHFAILURE_UNREALNAMEAUTHENTICATED = "AuthFailure.UnRealNameAuthenticated"

	// CAM鉴权失败。
	AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"

	// 不支持该操作。
	AUTHFAILURE_UNSUPPORTEDOPERATION = "AuthFailure.UnsupportedOperation"

	// 操作失败。
	FAILEDOPERATION = "FailedOperation"

	// 云端录制方法不支持。
	FAILEDOPERATION_CRUNSUPPORTMETHOD = "FailedOperation.CRUnsupportMethod"

	// 房间中没有设置混流模板。
	FAILEDOPERATION_MIXSESSIONNOTEXIST = "FailedOperation.MixSessionNotExist"

	// 需要解锁所需能力位
	FAILEDOPERATION_NOTABILITY = "FailedOperation.NotAbility"

	// 不允许此操作，请提交工单联系我们
	FAILEDOPERATION_NOTALLOWED = "FailedOperation.NotAllowed"

	// 查询任务失败
	FAILEDOPERATION_QUERYTASKINFOFAILED = "FailedOperation.QueryTaskInfoFailed"

	// 请求过期。
	FAILEDOPERATION_REQUESTOUTDATED = "FailedOperation.RequestOutdated"

	// 云API混流模板和SDK混流冲突。
	FAILEDOPERATION_REQUESTREJECTION = "FailedOperation.RequestRejection"

	// 单用户并发过载，请联系我们增大并发路数到合理值。
	FAILEDOPERATION_RESTRICTEDCONCURRENCY = "FailedOperation.RestrictedConcurrency"

	// 房间不存在。
	FAILEDOPERATION_ROOMNOTEXIST = "FailedOperation.RoomNotExist"

	// 应用ID不存在。
	FAILEDOPERATION_SDKAPPIDNOTEXIST = "FailedOperation.SdkAppIdNotExist"

	// AppId下没有此SdkAppId的资源
	FAILEDOPERATION_SDKAPPIDNOTUNDERAPPID = "FailedOperation.SdkAppIdNotUnderAppId"

	// 任务已存在
	FAILEDOPERATION_TASKEXIST = "FailedOperation.TaskExist"

	// 调用接口时任务已结束。
	FAILEDOPERATION_TASKFINISHED = "FailedOperation.TaskFinished"

	// 任务不存在或已结束
	FAILEDOPERATION_TASKNOTEXIST = "FailedOperation.TaskNotExist"

	// 内部错误。
	INTERNALERROR = "InternalError"

	// 云端录制内部服务错误。
	INTERNALERROR_CRINTERNALERROR = "InternalError.CRInternalError"

	// 数据库查询异常。
	INTERNALERROR_DBERROR = "InternalError.DBError"

	// ES查询异常。
	INTERNALERROR_ESQUERYERROR = "InternalError.EsQueryError"

	// 查询房间失败。
	INTERNALERROR_GETROOMCACHEIPERROR = "InternalError.GetRoomCacheIpError"

	// 获取房间信息失败。
	INTERNALERROR_GETROOMFROMCACHEERROR = "InternalError.GetRoomFromCacheError"

	// http请求解析失败。
	INTERNALERROR_HTTPPARASEFALIED = "InternalError.HttpParaseFalied"

	// http请求解析失败。
	INTERNALERROR_HTTPPARSEFAILED = "InternalError.HttpParseFailed"

	// 接口错误。
	INTERNALERROR_INTERFACEERR = "InternalError.InterfaceErr"

	// 内部错误，请重试。
	INTERNALERROR_INTERNALERROR = "InternalError.InternalError"

	// 不支持的方法。
	INTERNALERROR_METHODERR = "InternalError.MethodErr"

	// 参数错误。
	INVALIDPARAMETER = "InvalidParameter"

	// AppId校验失败。
	INVALIDPARAMETER_APPID = "InvalidParameter.AppId"

	// 音频编码参数错误。
	INVALIDPARAMETER_AUDIOENCODEPARAMS = "InvalidParameter.AudioEncodeParams"

	// BackgroundImageUrl参数错误。
	INVALIDPARAMETER_BACKGROUNDIMAGEURL = "InvalidParameter.BackgroundImageUrl"

	// body 解析参数失败。
	INVALIDPARAMETER_BODYPARAMSERROR = "InvalidParameter.BodyParamsError"

	// 图片过大。
	INVALIDPARAMETER_CHECKCONTENTFAILED = "InvalidParameter.CheckContentFailed"

	// 后缀名校验失败。
	INVALIDPARAMETER_CHECKSUFFIXFAILED = "InvalidParameter.CheckSuffixFailed"

	// EncodeParams参数错误。
	INVALIDPARAMETER_ENCODEPARAMS = "InvalidParameter.EncodeParams"

	// EndTs参数错误。
	INVALIDPARAMETER_ENDTS = "InvalidParameter.EndTs"

	// invalid sdkAppId
	INVALIDPARAMETER_INVALIDSDKAPPID = "InvalidParameter.InvalidSdkAppId"

	// 大画面居右显示参数错误。
	INVALIDPARAMETER_MAINVIDEORIGHTALIGN = "InvalidParameter.MainVideoRightAlign"

	// 大画面流类型错误。
	INVALIDPARAMETER_MAINVIDEOSTREAMTYPE = "InvalidParameter.MainVideoStreamType"

	// 参数超出范围。
	INVALIDPARAMETER_OUTOFRANGE = "InvalidParameter.OutOfRange"

	// OutputParams参数错误。
	INVALIDPARAMETER_OUTPUTPARAMS = "InvalidParameter.OutputParams"

	// PageNumber 参数错误。
	INVALIDPARAMETER_PAGENUMBER = "InvalidParameter.PageNumber"

	// PageSize参数错误。
	INVALIDPARAMETER_PAGESIZE = "InvalidParameter.PageSize"

	// PageSize超过100。
	INVALIDPARAMETER_PAGESIZEOVERSIZE = "InvalidParameter.PageSizeOversize"

	// 待操作的图片未找到。
	INVALIDPARAMETER_PICTURENOTFOUND = "InvalidParameter.PictureNotFound"

	// 自定义布局参数错误。
	INVALIDPARAMETER_PRESETLAYOUTCONFIG = "InvalidParameter.PresetLayoutConfig"

	// PublishCdnUrls参数校验失败。
	INVALIDPARAMETER_PUBLISHCDNURLS = "InvalidParameter.PublishCdnUrls"

	// 纯音频推流参数错误。
	INVALIDPARAMETER_PUREAUDIOSTREAM = "InvalidParameter.PureAudioStream"

	// 查询范围超过文档限制。
	INVALIDPARAMETER_QUERYSCALEOVERSIZE = "InvalidParameter.QueryScaleOversize"

	// 纯音频录制参数错误。
	INVALIDPARAMETER_RECORDAUDIOONLY = "InvalidParameter.RecordAudioOnly"

	// RecordId参数错误。
	INVALIDPARAMETER_RECORDID = "InvalidParameter.RecordId"

	// RoomId参数错误。
	INVALIDPARAMETER_ROOMID = "InvalidParameter.RoomId"

	// SdkAppId参数错误。
	INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"

	// 小画面布局参数错误。
	INVALIDPARAMETER_SMALLVIDEOLAYOUTPARAMS = "InvalidParameter.SmallVideoLayoutParams"

	// 小画面布局中流类型参数错误。
	INVALIDPARAMETER_SMALLVIDEOSTREAMTYPE = "InvalidParameter.SmallVideoStreamType"

	// 查询开始时间超过文档限制。
	INVALIDPARAMETER_STARTTIMEEXPIRE = "InvalidParameter.StartTimeExpire"

	// 查询起始时间超过当前监控仪表盘功能版本限制，详情请见:https://cloud.tencent.com/document/product/647/81331。
	INVALIDPARAMETER_STARTTIMEOVERSIZE = "InvalidParameter.StartTimeOversize"

	// StartTs参数错误。
	INVALIDPARAMETER_STARTTS = "InvalidParameter.StartTs"

	// 查询开始时间超过文档限制。
	INVALIDPARAMETER_STARTTSOVERSIZE = "InvalidParameter.StartTsOversize"

	// StrRoomId参数错误。
	INVALIDPARAMETER_STRROOMID = "InvalidParameter.StrRoomId"

	// StreamId参数错误。
	INVALIDPARAMETER_STREAMID = "InvalidParameter.StreamId"

	// StreamUrl参数校验失败
	INVALIDPARAMETER_STREAMURL = "InvalidParameter.StreamUrl"

	// TaskId 参数错误。
	INVALIDPARAMETER_TASKID = "InvalidParameter.TaskId"

	// Url解析参数失败。
	INVALIDPARAMETER_URLPARAMSERROR = "InvalidParameter.UrlParamsError"

	// UserId参数错误。
	INVALIDPARAMETER_USERID = "InvalidParameter.UserId"

	// UserIds参数错误。
	INVALIDPARAMETER_USERIDS = "InvalidParameter.UserIds"

	// 用户数超过6个。
	INVALIDPARAMETER_USERIDSMORETHANSIX = "InvalidParameter.UserIdsMorethanSix"

	// UserSig过期或错误。
	INVALIDPARAMETER_USERSIG = "InvalidParameter.UserSig"

	// UserSig 不是超级管理员。
	INVALIDPARAMETER_USERSIGNOTADMIN = "InvalidParameter.UserSigNotAdmin"

	// 视频分辨率参数错误。
	INVALIDPARAMETER_VIDEORESOLUTION = "InvalidParameter.VideoResolution"

	// RoomId值错误。
	INVALIDPARAMETERVALUE_ROOMID = "InvalidParameterValue.RoomId"

	// 缺少参数错误。
	MISSINGPARAMETER = "MissingParameter"

	// 缺少AccessKey参数。
	MISSINGPARAMETER_ACCESSKEY = "MissingParameter.AccessKey"

	// 缺少AppId参数。
	MISSINGPARAMETER_APPID = "MissingParameter.AppId"

	// EncodeParams中缺少音频输出参数。
	MISSINGPARAMETER_AUDIOENCODEPARAMS = "MissingParameter.AudioEncodeParams"

	// 转推参数中缺少BizId。
	MISSINGPARAMETER_BIZID = "MissingParameter.BizId"

	// 缺少Bucket参数。
	MISSINGPARAMETER_BUCKET = "MissingParameter.Bucket"

	// 缺少CloudStorage参数。
	MISSINGPARAMETER_CLOUDSTORAGE = "MissingParameter.CloudStorage"

	// 缺少CommId参数。
	MISSINGPARAMETER_COMMID = "MissingParameter.CommId"

	// 缺少SdkAppId参数或CommID参数。
	MISSINGPARAMETER_COMMIDORSDKAPPID = "MissingParameter.CommIdOrSdkAppId"

	// 缺少EncodeParams参数。
	MISSINGPARAMETER_ENCODEPARAMS = "MissingParameter.EncodeParams"

	// 缺少endTS_s参数。
	MISSINGPARAMETER_ENDTS = "MissingParameter.EndTs"

	// 缺少OutputParams参数。
	MISSINGPARAMETER_OUTPUTPARAMS = "MissingParameter.OutputParams"

	// 缺少自定义布局参数。
	MISSINGPARAMETER_PRESETLAYOUTCONFIG = "MissingParameter.PresetLayoutConfig"

	// 缺少转推参数。
	MISSINGPARAMETER_PUBLISHCDNPARAMS = "MissingParameter.PublishCdnParams"

	// 转推参数中缺少转推目的地址。
	MISSINGPARAMETER_PUBLISHCDNURLS = "MissingParameter.PublishCdnUrls"

	// 缺少RecordMode参数。
	MISSINGPARAMETER_RECORDMODE = "MissingParameter.RecordMode"

	// 缺少RecordParams参数。
	MISSINGPARAMETER_RECORDPARAMS = "MissingParameter.RecordParams"

	// 缺少Region参数。
	MISSINGPARAMETER_REGION = "MissingParameter.Region"

	// 缺少RoomId参数。
	MISSINGPARAMETER_ROOMID = "MissingParameter.RoomId"

	// 缺少RoomNum参数。
	MISSINGPARAMETER_ROOMNUM = "MissingParameter.RoomNum"

	// 缺少SdkAppId参数。
	MISSINGPARAMETER_SDKAPPID = "MissingParameter.SdkAppId"

	// 缺少SecretKey参数。
	MISSINGPARAMETER_SECRETKEY = "MissingParameter.SecretKey"

	// 缺少startTS_s参数。
	MISSINGPARAMETER_STARTTS = "MissingParameter.StartTs"

	// 缺少StorageParams参数。
	MISSINGPARAMETER_STORAGEPARAMS = "MissingParameter.StorageParams"

	// OutputParams中缺少StreamId参数。
	MISSINGPARAMETER_STREAMID = "MissingParameter.StreamId"

	// 缺少StreamType参数。
	MISSINGPARAMETER_STREAMTYPE = "MissingParameter.StreamType"

	// 缺少TaskId参数。
	MISSINGPARAMETER_TASKID = "MissingParameter.TaskId"

	// 缺少UserId参数。
	MISSINGPARAMETER_USERID = "MissingParameter.UserId"

	// 缺少UserIds参数。
	MISSINGPARAMETER_USERIDS = "MissingParameter.UserIds"

	// 缺少UserSig参数。
	MISSINGPARAMETER_USERSIG = "MissingParameter.UserSig"

	// 缺少Vendor参数。
	MISSINGPARAMETER_VENDOR = "MissingParameter.Vendor"

	// EncodeParams中缺少视频输出参数。
	MISSINGPARAMETER_VIDEOENCODEPARAMS = "MissingParameter.VideoEncodeParams"

	// 资源不足。
	RESOURCEINSUFFICIENT_REQUESTREJECTION = "ResourceInsufficient.RequestRejection"

	// 资源不存在。
	RESOURCENOTFOUND = "ResourceNotFound"

	// 没有操作SdkAppId的权限。
	UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"

	// 未知参数错误。
	UNKNOWNPARAMETER = "UnknownParameter"

	// 操作不支持。
	UNSUPPORTEDOPERATION = "UnsupportedOperation"
)
