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

package v20191213

const (
	// 此产品的特有错误码

	// 操作失败。
	FAILEDOPERATION = "FailedOperation"

	// 图片美颜失败，请更换图片。
	FAILEDOPERATION_BEAUTIFYFAILED = "FailedOperation.BeautifyFailed"

	// 撤销任务无法被成功执行, 请重试。
	FAILEDOPERATION_CANCELJOBFAILURE = "FailedOperation.CancelJobFailure"

	// 未检测到人脸。
	FAILEDOPERATION_DETECTNOFACE = "FailedOperation.DetectNoFace"

	// 操作太频繁，触发频控。
	FAILEDOPERATION_EFFECTFREQCTRL = "FailedOperation.EffectFreqCtrl"

	// 特效服务内部错误。
	FAILEDOPERATION_EFFECTINNERERROR = "FailedOperation.EffectInnerError"

	// 人脸因太小被过滤，建议人脸分辨率不小于34*34。
	FAILEDOPERATION_FACESIZETOOSMALL = "FailedOperation.FaceSizeTooSmall"

	// 操作太频繁，触发频控，请稍后重试。
	FAILEDOPERATION_FREQCTRL = "FailedOperation.FreqCtrl"

	// 图片解码失败。
	FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"

	// 图片下载失败。
	FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"

	// 不支持灰色图。
	FAILEDOPERATION_IMAGEGRAYNOTSUPPORT = "FailedOperation.ImageGrayNotSupport"

	// 不支持的图片文件。
	FAILEDOPERATION_IMAGENOTSUPPORTED = "FailedOperation.ImageNotSupported"

	// 图片分辨率过大，超过2000*2000。
	FAILEDOPERATION_IMAGERESOLUTIONEXCEED = "FailedOperation.ImageResolutionExceed"

	// 图片短边分辨率小于64。
	FAILEDOPERATION_IMAGERESOLUTIONTOOSMALL = "FailedOperation.ImageResolutionTooSmall"

	// 图片数据太大。
	FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"

	// 图片上传失败。
	FAILEDOPERATION_IMAGEUPLOADFAILED = "FailedOperation.ImageUploadFailed"

	// 服务内部错误，请重试。
	FAILEDOPERATION_INNERERROR = "FailedOperation.InnerError"

	// 内部错误。
	FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"

	// 任务冲突。
	FAILEDOPERATION_JOBCONFLICT = "FailedOperation.JobConflict"

	// 任务已撤销，请重新提交任务。
	FAILEDOPERATION_JOBHASBEENCANCELED = "FailedOperation.JobHasBeenCanceled"

	// 任务已停止处理，请重新提交任务。
	FAILEDOPERATION_JOBSTOPPROCESSING = "FailedOperation.JobStopProcessing"

	// 素材超过数量限制。
	FAILEDOPERATION_MODELVALUEEXCEED = "FailedOperation.ModelValueExceed"

	// 参数错误。
	FAILEDOPERATION_PARAMETERVALUEERROR = "FailedOperation.ParameterValueError"

	// 请求包体太大，建议在6M内（通常是参数中包括大的图片）。
	FAILEDOPERATION_REQUESTENTITYTOOLARGE = "FailedOperation.RequestEntityTooLarge"

	// 后端服务超时，请重试。
	FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"

	// RPC请求失败，一般为算法微服务故障。
	FAILEDOPERATION_RPCFAIL = "FailedOperation.RpcFail"

	// 系统内部错误。
	FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"

	// 未知错误。
	FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"

	// 内部错误。
	INTERNALERROR = "InternalError"

	// 参数错误。
	INVALIDPARAMETER = "InvalidParameter"

	// 参数不合法。
	INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"

	// 图片中没有人脸。
	INVALIDPARAMETER_NOFACEINPHOTO = "InvalidParameter.NoFaceInPhoto"

	// 参数取值错误。
	INVALIDPARAMETERVALUE = "InvalidParameterValue"

	// 大眼参数不合法。
	INVALIDPARAMETERVALUE_EYEENLARGINGILLEGAL = "InvalidParameterValue.EyeEnlargingIllegal"

	// 瘦脸参数不合法。
	INVALIDPARAMETERVALUE_FACELIFTINGILLEGAL = "InvalidParameterValue.FaceLiftingIllegal"

	// 人脸框参数不合法。
	INVALIDPARAMETERVALUE_FACERECTINVALID = "InvalidParameterValue.FaceRectInvalid"

	// 第1个人脸框参数不合法。
	INVALIDPARAMETERVALUE_FACERECTINVALIDFIRST = "InvalidParameterValue.FaceRectInvalidFirst"

	// 第2个人脸框参数不合法。
	INVALIDPARAMETERVALUE_FACERECTINVALIDSECOND = "InvalidParameterValue.FaceRectInvalidSecond"

	// 第3个人脸框参数不合法。
	INVALIDPARAMETERVALUE_FACERECTINVALIDTHRID = "InvalidParameterValue.FaceRectInvalidThrid"

	// 缺少参数，请检查图片参数是否为空。
	INVALIDPARAMETERVALUE_IMAGEEMPTY = "InvalidParameterValue.ImageEmpty"

	// 图片不合法。
	INVALIDPARAMETERVALUE_IMAGEINVALID = "InvalidParameterValue.ImageInvalid"

	// 图片数据太大。
	INVALIDPARAMETERVALUE_IMAGESIZEEXCEED = "InvalidParameterValue.ImageSizeExceed"

	// 素材图片不合法, 必须是512*512的PNG图片。
	INVALIDPARAMETERVALUE_LUTIMAGEINVALID = "InvalidParameterValue.LutImageInvalid"

	// 素材图片尺寸不合法，必须是512*512的PNG图片。
	INVALIDPARAMETERVALUE_LUTIMAGESIZEINVALID = "InvalidParameterValue.LutImageSizeInvalid"

	// 素材不存在。
	INVALIDPARAMETERVALUE_MODELIDNOTFOUND = "InvalidParameterValue.ModelIdNotFound"

	// 图片中没有人脸。
	INVALIDPARAMETERVALUE_NOFACEINPHOTO = "InvalidParameterValue.NoFaceInPhoto"

	// 参数错误。
	INVALIDPARAMETERVALUE_PARAMETERVALUEERROR = "InvalidParameterValue.ParameterValueError"

	// 磨皮参数不合法。
	INVALIDPARAMETERVALUE_SMOOTHINGILLEGAL = "InvalidParameterValue.SmoothingIllegal"

	// URL格式不合法。
	INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"

	// 美白参数不合法。
	INVALIDPARAMETERVALUE_WHITENINGILLEGAL = "InvalidParameterValue.WhiteningIllegal"

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
