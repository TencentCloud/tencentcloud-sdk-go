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

package v20190318

const (
	// 此产品的特有错误码

	// 未支持的功能。
	FAILEDOPERATION_NOTSUPPORTEDFUNCTIONERROR = "FailedOperation.NotSupportedFunctionError"

	// 内部错误
	INTERNALERROR = "InternalError"

	// 任务已被强制取消
	INTERNALERROR_JOBCANCELED = "InternalError.JobCanceled"

	// 任务失败重试次数超过上限
	INTERNALERROR_JOBREACHMAXRETRYTIMES = "InternalError.JobReachMaxRetryTimes"

	// 查询人脸库失败
	INTERNALERROR_STATUSQUERYFACELIBFAILED = "InternalError.StatusQueryFaceLibFailed"

	// 参数错误
	INVALIDPARAMETER = "InvalidParameter"

	// 人体检索失败
	INVALIDPARAMETER_BODYRETRIEVALFAILED = "InvalidParameter.BodyRetrievalFailed"

	// 无法检测到人脸，请确认输入源无误。
	INVALIDPARAMETER_CANNOTFINDFACEERROR = "InvalidParameter.CannotFindFaceError"

	// 图片中没有人脸
	INVALIDPARAMETER_CANNOTFINDFACEINPICTURE = "InvalidParameter.CannotFindFaceInPicture"

	// 从指定URL读取视频失败。
	INVALIDPARAMETER_CANNOTREADVIDEOFROMURLERROR = "InvalidParameter.CannotReadVideoFromUrlError"

	// 裁剪帧失败
	INVALIDPARAMETER_CLIPPINGFRAMEFAILED = "InvalidParameter.ClippingFrameFailed"

	// 计算相似度失败
	INVALIDPARAMETER_COMPUTATIONALSIMILARITYFAILED = "InvalidParameter.ComputationalSimilarityFailed"

	// 添加 Face 人脸失败
	INVALIDPARAMETER_CREATEFACEFAILED = "InvalidParameter.CreateFaceFailed"

	// 添加 Person 个体失败
	INVALIDPARAMETER_CREATEPERSONFAILED = "InvalidParameter.CreatePersonFailed"

	// 人脸检测失败
	INVALIDPARAMETER_FACEDETECTEDFAILED = "InvalidParameter.FaceDetectedFailed"

	// 删除face人脸失败，不存在该faceid
	INVALIDPARAMETER_FACENOTEXIST = "InvalidParameter.FaceNotExist"

	// 人脸不存在
	INVALIDPARAMETER_FACENOTEXISTED = "InvalidParameter.FaceNotExisted"

	// 超出单次检索的人脸数上线
	INVALIDPARAMETER_FACENUMBEREXCEEDLIMITOFSINGLEDETECTION = "InvalidParameter.FaceNumberExceedLimitOfSingleDetection"

	// 该个体人脸已经到最大值
	INVALIDPARAMETER_FACENUMBERLIMIT = "InvalidParameter.FaceNumberLimit"

	// 人脸配准失败
	INVALIDPARAMETER_FACEREGISTRATIONFAILED = "InvalidParameter.FaceRegistrationFailed"

	// 人脸检索失败
	INVALIDPARAMETER_FACERETRIEVALFAILED = "InvalidParameter.FaceRetrievalFailed"

	// 人脸太小
	INVALIDPARAMETER_FACETOOSMALL = "InvalidParameter.FaceTooSmall"

	// 解码图片失败
	INVALIDPARAMETER_FAILEDTODECODEIMAGE = "InvalidParameter.FailedToDecodeImage"

	// 特征提取失败
	INVALIDPARAMETER_FEATUREEXTRACTIONFAILED = "InvalidParameter.FeatureExtractionFailed"

	// 输入内容为空
	INVALIDPARAMETER_FILECONTENTEMPTY = "InvalidParameter.FileContentEmpty"

	// 输入图片无法解码
	INVALIDPARAMETER_IMAGEDECODEFAILED = "InvalidParameter.ImageDecodeFailed"

	// 图片太大
	INVALIDPARAMETER_IMAGETOOLARGE = "InvalidParameter.ImageTooLarge"

	// 图片太小
	INVALIDPARAMETER_IMAGETOOSMALL = "InvalidParameter.ImageTooSmall"

	// 不支持的 FileType
	INVALIDPARAMETER_INVALIDFILETYPE = "InvalidParameter.InvalidFileType"

	// 无效的语言类型，只支持0 英文，1 中文
	INVALIDPARAMETER_INVALIDLANG = "InvalidParameter.InvalidLang"

	// 资源标识符应为字母、数字或者下划线组成
	INVALIDPARAMETER_INVALIDRESOURCEIDENTIFY = "InvalidParameter.InvalidResourceIdentify"

	// 名称应为60个字符内的汉字、字母、数字或者下划线组成
	INVALIDPARAMETER_INVALIDRESOURCENAME = "InvalidParameter.InvalidResourceName"

	// 分片序号错误
	INVALIDPARAMETER_INVALIDSEQID = "InvalidParameter.InvalidSeqId"

	// 输入音频文件无效
	INVALIDPARAMETER_INVALIDURL = "InvalidParameter.InvalidUrl"

	// 任务查询失败，请检查任务标识符后重试
	INVALIDPARAMETER_JOBNOTVALID = "InvalidParameter.JobNotValid"

	// 员库中尚未有任何人员，请确认输入人员库ID无误。
	INVALIDPARAMETER_LIBHAVENOPERSON = "InvalidParameter.LibHaveNoPerson"

	// 请确认输入人员库ID。
	INVALIDPARAMETER_LIBISEMPTY = "InvalidParameter.LibIsEmpty"

	// 人员库不存在
	INVALIDPARAMETER_LIBRARYNOTEXISTED = "InvalidParameter.LibraryNotExisted"

	// 人脸质量低，包括不清晰
	INVALIDPARAMETER_LOWQUALITYPICTURE = "InvalidParameter.LowQualityPicture"

	// 服务不存在
	INVALIDPARAMETER_NOSUCHSERVICE = "InvalidParameter.NoSuchService"

	// 非人员考勤任务
	INVALIDPARAMETER_NOTATTENDANCEJOB = "InvalidParameter.NotAttendanceJob"

	// 人脸角度pitch不合法
	INVALIDPARAMETER_NOTVALIDPITCH = "InvalidParameter.NotValidPitch"

	// 人脸角度row不合法
	INVALIDPARAMETER_NOTVALIDROW = "InvalidParameter.NotValidRow"

	// 人脸角度yaw不合法
	INVALIDPARAMETER_NOTVALIDYAW = "InvalidParameter.NotValidYaw"

	// 指定人员不存在
	INVALIDPARAMETER_PERSONEXISTED = "InvalidParameter.PersonExisted"

	// 人员标识不合法
	INVALIDPARAMETER_PERSONIDNOTVALID = "InvalidParameter.PersonIdNotValid"

	// 删除person个体失败，不存在该个体
	INVALIDPARAMETER_PERSONNOTEXIST = "InvalidParameter.PersonNotExist"

	// 人员不存在
	INVALIDPARAMETER_PERSONNOTEXISTED = "InvalidParameter.PersonNotExisted"

	// 人员未注册
	INVALIDPARAMETER_PERSONNOTREGISTERED = "InvalidParameter.PersonNotRegistered"

	// 任务处理中，请稍后查询。
	INVALIDPARAMETER_PROCESSUNREADY = "InvalidParameter.ProcessUnReady"

	// 资源名称需要保持唯一
	INVALIDPARAMETER_RESOURCENAMEDUPLICATE = "InvalidParameter.ResourceNameDuplicate"

	// 任务执行失败，请确保拉流地址稳定
	INVALIDPARAMETER_STATUSJOBFAILED = "InvalidParameter.StatusJobFailed"

	// 任务不存在
	INVALIDPARAMETER_STATUSJOBNOTFOUND = "InvalidParameter.StatusJobNotFound"

	// 任务未完成
	INVALIDPARAMETER_STATUSJOBUNFINISHED = "InvalidParameter.StatusJobUnfinished"

	// 没有找到摄像头
	INVALIDPARAMETER_STATUSNOCAMERAFOUND = "InvalidParameter.StatusNoCameraFound"

	// 图片中含有多张人脸
	INVALIDPARAMETER_TOOMANYFACE = "InvalidParameter.TooManyFace"

	// 轨迹追踪失败
	INVALIDPARAMETER_TRACKINGFAILED = "InvalidParameter.TrackingFailed"

	// 不支持MOV类型文件
	INVALIDPARAMETER_UNSUPPORTEDFILETYPEMOV = "InvalidParameter.UnsupportedFileTypeMov"

	// 视频尺寸不符合要求(1080p及以下)
	INVALIDPARAMETER_UNSUPPORTEDVIDEOSIZE = "InvalidParameter.UnsupportedVideoSize"

	// 任务已经启动分析，请勿重复提交。
	INVALIDPARAMETER_VIDEOALREDYPROCESSEDERROR = "InvalidParameter.VideoAlredyProcessedError"

	// 获取请求url http包体失败。
	INVALIDPARAMETERVALUE_GETHTTPBODYERROR = "InvalidParameterValue.GetHttpBodyError"

	// 人员库自定义描述字段数组长度超过限制。最多可以创建5个。
	INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSEXCEED = "InvalidParameterValue.GroupExDescriptionsExceed"

	// 人员库自定义描述字段名称不可重复。
	INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMEIDENTICAL = "InvalidParameterValue.GroupExDescriptionsNameIdentical"

	// 人员库自定义描述字段名称包含非法字符。人员库自定义描述字段名称只支持中英文、-、_、数字。
	INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMEILLEGAL = "InvalidParameterValue.GroupExDescriptionsNameIllegal"

	// 人员库自定义描述字段名称长度超出限制。
	INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMETOOLONG = "InvalidParameterValue.GroupExDescriptionsNameTooLong"

	// 图片错误
	INVALIDPARAMETERVALUE_IMAGEILLEGAL = "InvalidParameterValue.ImageIllegal"

	// 分片序号错误。
	INVALIDPARAMETERVALUE_INVALIDSEQID = "InvalidParameterValue.InvalidSeqId"

	// 未支持的功能。
	INVALIDPARAMETERVALUE_NOTSUPPORTEDFUNCTIONERROR = "InvalidParameterValue.NotSupportedFunctionError"

	// 人员自定义描述字段数组长度超过限制。最多5个。
	INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONINFOSEXCEED = "InvalidParameterValue.PersonExDescriptionInfosExceed"

	// 该教室的摄像头数量超出最大限制
	LIMITEXCEEDED_STATUSCAMERACOUNTOVERMAX = "LimitExceeded.StatusCameraCountOverMax"

	// 资源不存在
	RESOURCENOTFOUND = "ResourceNotFound"

	// 无法找到精彩集锦任务ID，请确认输入JobId无误。
	RESOURCENOTFOUND_HLJOBIDNOTFOUND = "ResourceNotFound.HLJobIdNotFound"

	// 评估之前没有初始化或已过期。
	RESOURCEUNAVAILABLE_CANNOTFINDSESSION = "ResourceUnavailable.CannotFindSession"

	// 提交任务时获取拉流地址失败，请确认该教室的摄像头都配置好
	RESOURCEUNAVAILABLE_STATUSQUERYSTREAMFAILED = "ResourceUnavailable.StatusQueryStreamFailed"
)
