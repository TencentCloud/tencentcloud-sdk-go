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

package v20180301

const (
	// 此产品的特有错误码

	// CAM签名/鉴权错误。
	AUTHFAILURE_INVALIDAUTHORIZATION = "AuthFailure.InvalidAuthorization"

	// 未检测到闭眼动作。
	FAILEDOPERATION_ACTIONCLOSEEYE = "FailedOperation.ActionCloseEye"

	// 脸离屏幕太近。
	FAILEDOPERATION_ACTIONFACECLOSE = "FailedOperation.ActionFaceClose"

	// 脸离屏幕太远。
	FAILEDOPERATION_ACTIONFACEFAR = "FailedOperation.ActionFaceFar"

	// 脸离屏幕太左。
	FAILEDOPERATION_ACTIONFACELEFT = "FailedOperation.ActionFaceLeft"

	// 脸离屏幕太右。
	FAILEDOPERATION_ACTIONFACERIGHT = "FailedOperation.ActionFaceRight"

	// 未检测到动作配合。
	FAILEDOPERATION_ACTIONFIRSTACTION = "FailedOperation.ActionFirstAction"

	// 光线太暗。
	FAILEDOPERATION_ACTIONLIGHTDARK = "FailedOperation.ActionLightDark"

	// 光线太强。
	FAILEDOPERATION_ACTIONLIGHTSTRONG = "FailedOperation.ActionLightStrong"

	// 未能检测到完整人脸。
	FAILEDOPERATION_ACTIONNODETECTFACE = "FailedOperation.ActionNodetectFace"

	// 未检测到张嘴动作。
	FAILEDOPERATION_ACTIONOPENMOUTH = "FailedOperation.ActionOpenMouth"

	// 比对失败。
	FAILEDOPERATION_COMPAREFAIL = "FailedOperation.CompareFail"

	// 比对相似度未达到通过标准。
	FAILEDOPERATION_COMPARELOWSIMILARITY = "FailedOperation.CompareLowSimilarity"

	// 调用比对引擎接口出错。
	FAILEDOPERATION_COMPARESYSTEMERROR = "FailedOperation.CompareSystemError"

	// 数据库异常。
	FAILEDOPERATION_DBERROR = "FailedOperation.DbError"

	// 解密失败。
	FAILEDOPERATION_DECRYPTSYSTEMERROR = "FailedOperation.DecryptSystemError"

	// 文件下载失败。
	FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"

	// 文件下载超时。
	FAILEDOPERATION_DOWNLOADTIMEOUTERROR = "FailedOperation.DownLoadTimeoutError"

	// 图片内容为空。
	FAILEDOPERATION_EMPTYIMAGEERROR = "FailedOperation.EmptyImageError"

	// 加密失败。
	FAILEDOPERATION_ENCRYPTSYSTEMERROR = "FailedOperation.EncryptSystemError"

	// 文件存储失败，请稍后重试。
	FAILEDOPERATION_FILESAVEERROR = "FailedOperation.FileSaveError"

	// 输入的身份证号有误。
	FAILEDOPERATION_IDFORMATERROR = "FailedOperation.IdFormatError"

	// 姓名和身份证号不一致，请核实后重试。
	FAILEDOPERATION_IDNAMEMISMATCH = "FailedOperation.IdNameMisMatch"

	// 库中无此号，请到户籍所在地进行核实。
	FAILEDOPERATION_IDNOEXISTSYSTEM = "FailedOperation.IdNoExistSystem"

	// 库中无此号照片，请到户籍所在地进行核实。
	FAILEDOPERATION_IDPHOTONOEXIST = "FailedOperation.IdPhotoNoExist"

	// 证件图片分辨率太低，请重新上传。
	FAILEDOPERATION_IDPHOTOPOORQUALITY = "FailedOperation.IdPhotoPoorQuality"

	// 客户库自建库或认证中心返照失败，请稍后再试。
	FAILEDOPERATION_IDPHOTOSYSTEMNOANSWER = "FailedOperation.IdPhotoSystemNoanswer"

	// 图片模糊。
	FAILEDOPERATION_IMAGEBLUR = "FailedOperation.ImageBlur"

	// 图片解码失败。
	FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"

	// 图片中未检测到身份证。
	FAILEDOPERATION_IMAGENOIDCARD = "FailedOperation.ImageNoIdCard"

	// 图片尺寸过大。
	FAILEDOPERATION_IMAGESIZETOOLARGE = "FailedOperation.ImageSizeTooLarge"

	// 检测到多张人脸。
	FAILEDOPERATION_LIFEPHOTODETECTFACES = "FailedOperation.LifePhotoDetectFaces"

	// 实人比对没通过。
	FAILEDOPERATION_LIFEPHOTODETECTFAKE = "FailedOperation.LifePhotoDetectFake"

	// 未能检测到完整人脸。
	FAILEDOPERATION_LIFEPHOTODETECTNOFACES = "FailedOperation.LifePhotoDetectNoFaces"

	// 传入图片分辨率太低，请重新上传。
	FAILEDOPERATION_LIFEPHOTOPOORQUALITY = "FailedOperation.LifePhotoPoorQuality"

	// 传入图片过大或过小。
	FAILEDOPERATION_LIFEPHOTOSIZEERROR = "FailedOperation.LifePhotoSizeError"

	// 脸部未完整露出。
	FAILEDOPERATION_LIPFACEINCOMPLETE = "FailedOperation.LipFaceIncomplete"

	// 嘴唇动作幅度过小。
	FAILEDOPERATION_LIPMOVESMALL = "FailedOperation.LipMoveSmall"

	// 视频拉取失败，请重试。
	FAILEDOPERATION_LIPNETFAILED = "FailedOperation.LipNetFailed"

	// 视频为空，或大小不合适，请控制录制时长在6s左右。
	FAILEDOPERATION_LIPSIZEERROR = "FailedOperation.LipSizeError"

	// 视频格式有误。
	FAILEDOPERATION_LIPVIDEOINVALID = "FailedOperation.LipVideoInvalid"

	// 视频像素太低。
	FAILEDOPERATION_LIPVIDEOQUAILITY = "FailedOperation.LipVideoQuaility"

	// 未检测到声音。
	FAILEDOPERATION_LIPVOICEDETECT = "FailedOperation.LipVoiceDetect"

	// 视频声音太小。
	FAILEDOPERATION_LIPVOICELOW = "FailedOperation.LipVoiceLow"

	// 声音识别失败。
	FAILEDOPERATION_LIPVOICERECOGNIZE = "FailedOperation.LipVoiceRecognize"

	// 人脸检测失败，无法提取比对照。
	FAILEDOPERATION_LIVESSBESTFRAMEERROR = "FailedOperation.LivessBestFrameError"

	// 活体检测没通过。
	FAILEDOPERATION_LIVESSDETECTFAIL = "FailedOperation.LivessDetectFail"

	// 疑似非真人录制。
	FAILEDOPERATION_LIVESSDETECTFAKE = "FailedOperation.LivessDetectFake"

	// 调用活体引擎接口出错。
	FAILEDOPERATION_LIVESSSYSTEMERROR = "FailedOperation.LivessSystemError"

	// 视频实人检测没通过。
	FAILEDOPERATION_LIVESSUNKNOWNERROR = "FailedOperation.LivessUnknownError"

	// 输入的姓名有误。
	FAILEDOPERATION_NAMEFORMATERROR = "FailedOperation.NameFormatError"

	// Ocr识别失败。
	FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"

	// 调用次数超出限制。
	FAILEDOPERATION_REQUESTLIMITEXCEEDED = "FailedOperation.RequestLimitExceeded"

	// 实人检测失败。
	FAILEDOPERATION_SILENTDETECTFAIL = "FailedOperation.SilentDetectFail"

	// 眼睛检测失败。
	FAILEDOPERATION_SILENTEYELIVEFAIL = "FailedOperation.SilentEyeLiveFail"

	// 视频未检测到人脸。
	FAILEDOPERATION_SILENTFACEDETECTFAIL = "FailedOperation.SilentFaceDetectFail"

	// 视频中人脸质量低。
	FAILEDOPERATION_SILENTFACEQUALITYFAIL = "FailedOperation.SilentFaceQualityFail"

	// 检测到带口罩。
	FAILEDOPERATION_SILENTFACEWITHMASKFAIL = "FailedOperation.SilentFaceWithMaskFail"

	// 嘴巴检测失败。
	FAILEDOPERATION_SILENTMOUTHLIVEFAIL = "FailedOperation.SilentMouthLiveFail"

	// 视频检测中有多个人脸。
	FAILEDOPERATION_SILENTMULTIFACEFAIL = "FailedOperation.SilentMultiFaceFail"

	// 疑似翻拍。
	FAILEDOPERATION_SILENTPICTURELIVEFAIL = "FailedOperation.SilentPictureLiveFail"

	// 实人检测未达到通过标准。
	FAILEDOPERATION_SILENTTHRESHOLD = "FailedOperation.SilentThreshold"

	// 视频录制时间过短，请录制2秒以上的视频。
	FAILEDOPERATION_SILENTTOOSHORT = "FailedOperation.SilentTooShort"

	// STS未授权。
	FAILEDOPERATION_STSUNAUTHERRERROR = "FailedOperation.StsUnAuthErrError"

	// 内部未知错误。
	FAILEDOPERATION_UNKNOWN = "FailedOperation.UnKnown"

	// 该用户未注册E证通，请先注册并跟权威库核验。
	FAILEDOPERATION_UNREGISTEREDEID = "FailedOperation.UnregisteredEid"

	// 认证不通过。
	FAILEDOPERATION_VERIFICATIONFAIL = "FailedOperation.VerificationFail"

	// 内部错误。
	INTERNALERROR = "InternalError"

	// 加密失败。
	INTERNALERROR_ENCRYPTSYSTEMERROR = "InternalError.EncryptSystemError"

	// 内部未知错误。
	INTERNALERROR_UNKNOWN = "InternalError.UnKnown"

	// 参数错误。
	INVALIDPARAMETER = "InvalidParameter"

	// RuleId不存在。
	INVALIDPARAMETER_RULEID = "InvalidParameter.RuleId"

	// 存在不加密的字段，请参考文档修改。
	INVALIDPARAMETER_UNSUPPORTENCRYPTFIELD = "InvalidParameter.UnsupportEncryptField"

	// 参数取值错误。
	INVALIDPARAMETERVALUE = "InvalidParameterValue"

	// BizToken过期。
	INVALIDPARAMETERVALUE_BIZTOKENEXPIRED = "InvalidParameterValue.BizTokenExpired"

	// BizToken不合法。
	INVALIDPARAMETERVALUE_BIZTOKENILLEGAL = "InvalidParameterValue.BizTokenIllegal"

	// 该ruleid已被您停用，请确认后重试。
	INVALIDPARAMETERVALUE_RULEIDDISABLED = "InvalidParameterValue.RuleIdDisabled"

	// RuleId不存在，请到人脸核身控制台申请。
	INVALIDPARAMETERVALUE_RULEIDNOTEXIST = "InvalidParameterValue.RuleIdNotExist"

	// 请求的次数超过了频率限制。
	REQUESTLIMITEXCEEDED = "RequestLimitExceeded"

	// 未授权操作。
	UNAUTHORIZEDOPERATION = "UnauthorizedOperation"

	// 服务开通异常。
	UNAUTHORIZEDOPERATION_ACTIVATEERROR = "UnauthorizedOperation.ActivateError"

	// 服务开通中。
	UNAUTHORIZEDOPERATION_ACTIVATING = "UnauthorizedOperation.Activating"

	// 帐号已欠费。
	UNAUTHORIZEDOPERATION_ARREARS = "UnauthorizedOperation.Arrears"

	// 计费状态异常。
	UNAUTHORIZEDOPERATION_CHARGESTATUSEXCEPTION = "UnauthorizedOperation.ChargeStatusException"

	// 账号未实名。
	UNAUTHORIZEDOPERATION_NONAUTHORIZE = "UnauthorizedOperation.NonAuthorize"

	// 未开通服务。
	UNAUTHORIZEDOPERATION_NONACTIVATED = "UnauthorizedOperation.Nonactivated"

	// 操作不支持。
	UNSUPPORTEDOPERATION = "UnsupportedOperation"
)
