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

	// SecredId失效。
	AUTHFAILURE_INVALIDSECRETID = "AuthFailure.InvalidSecretId"

	// MFA失败。
	AUTHFAILURE_MFAFAILURE = "AuthFailure.MFAFailure"

	// SecretId不存在。
	AUTHFAILURE_SECRETIDNOTFOUND = "AuthFailure.SecretIdNotFound"

	// 签名已过期。
	AUTHFAILURE_SIGNATUREEXPIRE = "AuthFailure.SignatureExpire"

	// 签名校验失败。
	AUTHFAILURE_SIGNATUREFAILURE = "AuthFailure.SignatureFailure"

	// 任务已完成。
	AUTHFAILURE_TASKFINISHED = "AuthFailure.TaskFinished"

	// 令牌失败。
	AUTHFAILURE_TOKENFAILURE = "AuthFailure.TokenFailure"

	// 用户已激活。
	AUTHFAILURE_USERACTIVATED = "AuthFailure.UserActivated"

	// 用户状态异常。
	AUTHFAILURE_USERINVALIDSTATUS = "AuthFailure.UserInvalidStatus"

	// 用户无权限。
	AUTHFAILURE_USERNOTFOUND = "AuthFailure.UserNotFound"

	// 用户已欠费停服。
	AUTHFAILURE_USERSTOPARREAR = "AuthFailure.UserStopArrear"

	// 操作失败。
	FAILEDOPERATION = "FailedOperation"

	// 匹配的模板不存在。
	FAILEDOPERATION_AITEMPLATENOTEXIST = "FailedOperation.AiTemplateNotExist"

	// 自定义人物分类已存在。
	FAILEDOPERATION_CATEGORYEXIST = "FailedOperation.CategoryExist"

	// 自定义类型层级变化。
	FAILEDOPERATION_CATEGORYLEVELCHANGED = "FailedOperation.CategoryLevelChanged"

	// 自定义人物分类被引用，不能删除。
	FAILEDOPERATION_CATEGORYREFERRED = "FailedOperation.CategoryReferred"

	// 自定义人物库已存在。
	FAILEDOPERATION_CUSTOMGROUPALREADYEXIST = "FailedOperation.CustomGroupAlreadyExist"

	// 内部DB连接失败。
	FAILEDOPERATION_DBCONNECTIONERROR = "FailedOperation.DBConnectionError"

	// 媒资文件下载失败。
	FAILEDOPERATION_DOWNLOADFAILED = "FailedOperation.DownloadFailed"

	// 图片特征提取失败。
	FAILEDOPERATION_FEATUREALGOFAILED = "FailedOperation.FeatureAlgoFailed"

	// 获取CAM临时鉴权失败。
	FAILEDOPERATION_GETCAMTOKENFAILED = "FailedOperation.GetCAMTokenFailed"

	// 获取任务列表失败。
	FAILEDOPERATION_GETTASKLISTFAILED = "FailedOperation.GetTaskListFailed"

	// 获取媒资信息失败。
	FAILEDOPERATION_GETVIDEOMETADATAFAILED = "FailedOperation.GetVideoMetadataFailed"

	// 图片数量过多。
	FAILEDOPERATION_IMAGENUMEXCEEDED = "FailedOperation.ImageNumExceeded"

	// MD5不匹配。
	FAILEDOPERATION_MD5MISMATCH = "FailedOperation.MD5Mismatch"

	// 媒资文件已经存在。
	FAILEDOPERATION_MEDIAALREADYEXIST = "FailedOperation.MediaAlreadyExist"

	// 媒资文件已经过期。
	FAILEDOPERATION_MEDIAEXPIRED = "FailedOperation.MediaExpired"

	// 媒资正在使用。
	FAILEDOPERATION_MEDIAINUSE = "FailedOperation.MediaInUse"

	// 媒体文件未就绪。
	FAILEDOPERATION_MEDIANOTREADY = "FailedOperation.MediaNotReady"

	// 图片中包含多张人脸。
	FAILEDOPERATION_MULTIPLEFACESINIMAGE = "FailedOperation.MultipleFacesInImage"

	// 图片中不包含人脸。
	FAILEDOPERATION_NOFACEINIMAGE = "FailedOperation.NoFaceInImage"

	// 计费开通失败。
	FAILEDOPERATION_OPENCHARGEFAILED = "FailedOperation.OpenChargeFailed"

	// 人脸库中存在相似的人脸: %!s(MISSING)。
	FAILEDOPERATION_PERSONDUPLICATED = "FailedOperation.PersonDuplicated"

	// 人脸图片不属于已知人物。
	FAILEDOPERATION_PERSONNOTMATCHED = "FailedOperation.PersonNotMatched"

	// 自定义人物数量过多。
	FAILEDOPERATION_PERSONNUMEXCEEDED = "FailedOperation.PersonNumExceeded"

	// 图片质量分检测失败。
	FAILEDOPERATION_QUALITYALGOFAILED = "FailedOperation.QualityAlgoFailed"

	// 图片质量分过低。
	FAILEDOPERATION_QUALITYTOOLOW = "FailedOperation.QualityTooLow"

	// 结果快照反序列化失败。
	FAILEDOPERATION_SNAPSHOTDESERIALIZEFAILED = "FailedOperation.SnapshotDeserializeFailed"

	// 停止AI工作室任务失败。
	FAILEDOPERATION_STOPFLOWFAILED = "FailedOperation.StopFlowFailed"

	// 存在相同的任务。
	FAILEDOPERATION_TASKALREADYEXIST = "FailedOperation.TaskAlreadyExist"

	// 视频分析未完成。
	FAILEDOPERATION_TASKNOTFINISHED = "FailedOperation.TaskNotFinished"

	// 转码失败。
	FAILEDOPERATION_TRANSCODEFAILED = "FailedOperation.TranscodeFailed"

	// 上传文件失败。
	FAILEDOPERATION_UPLOADFAILED = "FailedOperation.UploadFailed"

	// 内部错误。
	INTERNALERROR = "InternalError"

	// 内部DB连接失败。
	INTERNALERROR_DBCONNECTIONERROR = "InternalError.DBConnectionError"

	// 内部DB操作错误。
	INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"

	// 内部错误。
	INTERNALERROR_INNERERROR = "InternalError.InnerError"

	// 自定义人物请求超过限制。
	INTERNALERROR_INTERNALOVERFLOW = "InternalError.InternalOverflow"

	// 参数错误。
	INVALIDPARAMETER = "InvalidParameter"

	// 自定义人物类型ID不合法。
	INVALIDPARAMETER_INVALIDCATEGORYID = "InvalidParameter.InvalidCategoryId"

	// 文件路径不合法。
	INVALIDPARAMETER_INVALIDFILEPATH = "InvalidParameter.InvalidFilePath"

	// 图片不合法。
	INVALIDPARAMETER_INVALIDIMAGE = "InvalidParameter.InvalidImage"

	// 图片ID不合法。
	INVALIDPARAMETER_INVALIDIMAGEID = "InvalidParameter.InvalidImageId"

	// 一级自定义类型不合法。
	INVALIDPARAMETER_INVALIDL1CATEGORY = "InvalidParameter.InvalidL1Category"

	// 二级自定义类型不合法。
	INVALIDPARAMETER_INVALIDL2CATEGORY = "InvalidParameter.InvalidL2Category"

	// MD5不合法。
	INVALIDPARAMETER_INVALIDMD5 = "InvalidParameter.InvalidMD5"

	// 媒体ID不合法。
	INVALIDPARAMETER_INVALIDMEDIAID = "InvalidParameter.InvalidMediaId"

	// MediaLabel无效。
	INVALIDPARAMETER_INVALIDMEDIALABEL = "InvalidParameter.InvalidMediaLabel"

	// MediaLang无效。
	INVALIDPARAMETER_INVALIDMEDIALANG = "InvalidParameter.InvalidMediaLang"

	// 媒体名称非法。
	INVALIDPARAMETER_INVALIDMEDIANAME = "InvalidParameter.InvalidMediaName"

	// MediaPreknownInfo无效。
	INVALIDPARAMETER_INVALIDMEDIAPREKNOWNINFO = "InvalidParameter.InvalidMediaPreknownInfo"

	// 媒资状态不合法。
	INVALIDPARAMETER_INVALIDMEDIASTATUS = "InvalidParameter.InvalidMediaStatus"

	// MediaType无效。
	INVALIDPARAMETER_INVALIDMEDIATYPE = "InvalidParameter.InvalidMediaType"

	// 名称不合法。
	INVALIDPARAMETER_INVALIDNAME = "InvalidParameter.InvalidName"

	// 分页序号不合法。
	INVALIDPARAMETER_INVALIDPAGENUMBER = "InvalidParameter.InvalidPageNumber"

	// 分页大小不合法。
	INVALIDPARAMETER_INVALIDPAGESIZE = "InvalidParameter.InvalidPageSize"

	// 输入字段 %!s(MISSING) 不合法。
	INVALIDPARAMETER_INVALIDPARAM = "InvalidParameter.InvalidParam"

	// 人物ID不合法。
	INVALIDPARAMETER_INVALIDPERSONID = "InvalidParameter.InvalidPersonId"

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

	// 参数超过长度限制。
	INVALIDPARAMETER_PARAMTOOLONG = "InvalidParameter.ParamTooLong"

	// 输入URL域名无法解析。
	INVALIDPARAMETER_URLNOTRESOLVED = "InvalidParameter.URLNotResolved"

	// 不支持的URL类型。
	INVALIDPARAMETER_UNSUPPORTURL = "InvalidParameter.UnsupportURL"

	// 使用量超过限制。
	LIMITEXCEEDED_USAGELIMITEXCEEDED = "LimitExceeded.UsageLimitExceeded"

	// 批量导入超过限制。
	REQUESTLIMITEXCEEDED_BATCHIMPORTOVERFLOW = "RequestLimitExceeded.BatchImportOverflow"

	// 同时发起过多任务。
	REQUESTLIMITEXCEEDED_CONCURRENCYOVERFLOW = "RequestLimitExceeded.ConcurrencyOverflow"

	// 自定义人物类型不存在。
	RESOURCENOTFOUND_CUSTOMCATEGORYNOTFOUND = "ResourceNotFound.CustomCategoryNotFound"

	// 自定义人物库不存在。
	RESOURCENOTFOUND_CUSTOMGROUPNOTFOUND = "ResourceNotFound.CustomGroupNotFound"

	// 媒资文件不存在。
	RESOURCENOTFOUND_MEDIANOTFOUND = "ResourceNotFound.MediaNotFound"

	// 记录不存在。
	RESOURCENOTFOUND_RECORDNOTFOUND = "ResourceNotFound.RecordNotFound"

	// 任务不存在。
	RESOURCENOTFOUND_TASKNOTFOUND = "ResourceNotFound.TaskNotFound"

	// 未授权操作。
	UNAUTHORIZEDOPERATION = "UnauthorizedOperation"

	// 用户未激活该产品。
	UNAUTHORIZEDOPERATION_UNAUTHORIZEDPRODUCT = "UnauthorizedOperation.UnauthorizedProduct"

	// 媒资文件不可访问。
	UNSUPPORTEDOPERATION_MEDIANOTACCESSIBLE = "UnsupportedOperation.MediaNotAccessible"

	// 任务不可访问。
	UNSUPPORTEDOPERATION_TASKNOTACCESSIBLE = "UnsupportedOperation.TaskNotAccessible"
)
