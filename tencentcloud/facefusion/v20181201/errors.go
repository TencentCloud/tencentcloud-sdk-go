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

package v20181201

const (
	// 此产品的特有错误码

	// 认证失败。
	AUTHFAILURE_INVALIDAUTHORIZATION = "AuthFailure.InvalidAuthorization"

	// 操作失败。
	FAILEDOPERATION = "FailedOperation"

	// 活动状态错误。
	FAILEDOPERATION_ACTIVITYSTATUSINVALID = "FailedOperation.ActivityStatusInvalid"

	// 人脸配准点出框错误码。
	FAILEDOPERATION_FACEBORDERCHECKFAILED = "FailedOperation.FaceBorderCheckFailed"

	// 人脸检测失败。
	FAILEDOPERATION_FACEDETECTFAILED = "FailedOperation.FaceDetectFailed"

	// 人脸出框，无法使用。
	FAILEDOPERATION_FACEEXCEEDBORDER = "FailedOperation.FaceExceedBorder"

	// 人脸提特征失败。
	FAILEDOPERATION_FACEFEATUREFAILED = "FailedOperation.FaceFeatureFailed"

	// 人脸融合失败，请更换图片后重试。
	FAILEDOPERATION_FACEFUSIONERROR = "FailedOperation.FaceFusionError"

	// 人脸姿态检测失败。
	FAILEDOPERATION_FACEPOSEFAILED = "FailedOperation.FacePoseFailed"

	// 人脸框不合法。
	FAILEDOPERATION_FACERECTINVALID = "FailedOperation.FaceRectInvalid"

	// 人脸配准点不合法。
	FAILEDOPERATION_FACESHAPEINVALID = "FailedOperation.FaceShapeInvalid"

	// 人脸因太小被过滤，建议人脸分辨率不小于34*34。
	FAILEDOPERATION_FACESIZETOOSMALL = "FailedOperation.FaceSizeTooSmall"

	// 人脸融合后端服务异常。
	FAILEDOPERATION_FUSEBACKENDSERVERFAULT = "FailedOperation.FuseBackendServerFault"

	// 未检测到人脸。
	FAILEDOPERATION_FUSEDETECTNOFACE = "FailedOperation.FuseDetectNoFace"

	// 操作太频繁，触发频控。
	FAILEDOPERATION_FUSEFREQCTRL = "FailedOperation.FuseFreqCtrl"

	// 图像处理出错。
	FAILEDOPERATION_FUSEIMAGEERROR = "FailedOperation.FuseImageError"

	// 服务内部错误，请重试。
	FAILEDOPERATION_FUSEINNERERROR = "FailedOperation.FuseInnerError"

	// 素材未经过审核。
	FAILEDOPERATION_FUSEMATERIALNOTAUTH = "FailedOperation.FuseMaterialNotAuth"

	// 素材不存在。
	FAILEDOPERATION_FUSEMATERIALNOTEXIST = "FailedOperation.FuseMaterialNotExist"

	// 保存结果图片出错。
	FAILEDOPERATION_FUSESAVEPHOTOFAIL = "FailedOperation.FuseSavePhotoFail"

	// 人脸检测-图片解码失败。
	FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"

	// 图片下载失败。
	FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"

	// 素材尺寸超过1080*1080像素。
	FAILEDOPERATION_IMAGEPIXELEXCEED = "FailedOperation.ImagePixelExceed"

	// 图片分辨率过大。建议您resize压缩到3k*3k以内。
	FAILEDOPERATION_IMAGERESOLUTIONEXCEED = "FailedOperation.ImageResolutionExceed"

	// 图片短边分辨率小于64。
	FAILEDOPERATION_IMAGERESOLUTIONTOOSMALL = "FailedOperation.ImageResolutionTooSmall"

	// 输入图片base64数据大小超过5M。
	FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"

	// base64编码后的图片数据大小不超500k。
	FAILEDOPERATION_IMAGESIZEEXCEEDFIVEHUNDREDKB = "FailedOperation.ImageSizeExceedFiveHundredKB"

	// 图片尺寸过大或者过小；不满足算法要求。
	FAILEDOPERATION_IMAGESIZEINVALID = "FailedOperation.ImageSizeInvalid"

	// 图片上传失败。
	FAILEDOPERATION_IMAGEUPLOADFAILED = "FailedOperation.ImageUploadFailed"

	// 服务内部错误。
	FAILEDOPERATION_INNERERROR = "FailedOperation.InnerError"

	// 素材条数超过上限。
	FAILEDOPERATION_MATERIALVALUEEXCEED = "FailedOperation.MaterialValueExceed"

	// 无法检测出人脸, 人脸框配准分低于阈值。
	FAILEDOPERATION_NOFACEDETECTED = "FailedOperation.NoFaceDetected"

	// 参数字段或者值有误。
	FAILEDOPERATION_PARAMETERVALUEERROR = "FailedOperation.ParameterValueError"

	// 活动未支付授权费或已停用。
	FAILEDOPERATION_PROJECTNOTAUTH = "FailedOperation.ProjectNotAuth"

	// 请求实体太大。
	FAILEDOPERATION_REQUESTENTITYTOOLARGE = "FailedOperation.RequestEntityTooLarge"

	// 后端服务超时。
	FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"

	// RPC请求失败，一般为算法微服务故障。
	FAILEDOPERATION_RPCFAIL = "FailedOperation.RpcFail"

	// 系统内部错误。
	FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"

	// 素材人脸ID不存在。
	FAILEDOPERATION_TEMPLATEFACEIDNOTEXIST = "FailedOperation.TemplateFaceIDNotExist"

	// 未知错误。
	FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"

	// 未查找到活动id。
	INVALIDPARAMETERVALUE_ACTIVITYIDNOTFOUND = "InvalidParameterValue.ActivityIdNotFound"

	// 活动算法版本值错误。
	INVALIDPARAMETERVALUE_ENGINEVALUEERROR = "InvalidParameterValue.EngineValueError"

	// 人脸框参数有误或者人脸框太小。
	INVALIDPARAMETERVALUE_FACERECTPARAMETERVALUEERROR = "InvalidParameterValue.FaceRectParameterValueError"

	// 人脸检测-图片为空。
	INVALIDPARAMETERVALUE_IMAGEEMPTY = "InvalidParameterValue.ImageEmpty"

	// 未查找到素材Id。
	INVALIDPARAMETERVALUE_MATERIALIDNOTFOUND = "InvalidParameterValue.MaterialIdNotFound"

	// 人脸检测-图片没有人脸。
	INVALIDPARAMETERVALUE_NOFACEINPHOTO = "InvalidParameterValue.NoFaceInPhoto"

	// 参数字段或者值有误。
	INVALIDPARAMETERVALUE_PARAMETERVALUEERROR = "InvalidParameterValue.ParameterValueError"

	// 资源正在发货中。
	RESOURCEUNAVAILABLE_DELIVERING = "ResourceUnavailable.Delivering"

	// 帐号已被冻结。
	RESOURCEUNAVAILABLE_FREEZE = "ResourceUnavailable.Freeze"

	// 获取认证信息失败。
	RESOURCEUNAVAILABLE_GETAUTHINFOERROR = "ResourceUnavailable.GetAuthInfoError"

	// 帐号已欠费。
	RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"

	// 余额不足。
	RESOURCEUNAVAILABLE_LOWBALANCE = "ResourceUnavailable.LowBalance"

	// 计费状态未知，请确认是否已在控制台开通服务。
	RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"

	// 服务未开通。
	RESOURCEUNAVAILABLE_NOTREADY = "ResourceUnavailable.NotReady"

	// 资源已被回收。
	RESOURCEUNAVAILABLE_RECOVER = "ResourceUnavailable.Recover"

	// 帐号已停服。
	RESOURCEUNAVAILABLE_STOPUSING = "ResourceUnavailable.StopUsing"

	// 计费状态未知。
	RESOURCEUNAVAILABLE_UNKNOWNSTATUS = "ResourceUnavailable.UnknownStatus"

	// 帐号已欠费。
	RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
)
