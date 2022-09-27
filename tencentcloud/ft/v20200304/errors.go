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

package v20200304

const (
	// 此产品的特有错误码

	// 操作失败。
	FAILEDOPERATION = "FailedOperation"

	// 撤销任务无法被成功执行, 请重试。
	FAILEDOPERATION_CANCELJOBFAILURE = "FailedOperation.CancelJobFailure"

	// 未检测到人脸。
	FAILEDOPERATION_DETECTNOFACE = "FailedOperation.DetectNoFace"

	// 人脸检测失败。
	FAILEDOPERATION_FACEDETECTFAILED = "FailedOperation.FaceDetectFailed"

	// 人脸出框，无法使用。
	FAILEDOPERATION_FACEEXCEEDBORDER = "FailedOperation.FaceExceedBorder"

	// 人脸配准失败。
	FAILEDOPERATION_FACESHAPEFAILED = "FailedOperation.FaceShapeFailed"

	// 人脸因太小被过滤，建议人脸分辨率不小于34*34。
	FAILEDOPERATION_FACESIZETOOSMALL = "FailedOperation.FaceSizeTooSmall"

	// 操作太频繁，触发频控，请稍后重试。
	FAILEDOPERATION_FREQCTRL = "FailedOperation.FreqCtrl"

	// 图片解码失败。
	FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"

	// 图片下载错误。
	FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"

	// 不支持的图片文件。
	FAILEDOPERATION_IMAGENOTSUPPORTED = "FailedOperation.ImageNotSupported"

	// 素材尺寸超过2000*2000像素。
	FAILEDOPERATION_IMAGEPIXELEXCEED = "FailedOperation.ImagePixelExceed"

	// 图片分辨率过大，超过2000*2000。
	FAILEDOPERATION_IMAGERESOLUTIONEXCEED = "FailedOperation.ImageResolutionExceed"

	// 图片短边分辨率太小，小于64。
	FAILEDOPERATION_IMAGERESOLUTIONTOOSMALL = "FailedOperation.ImageResolutionTooSmall"

	// 服务内部错误，请重试。
	FAILEDOPERATION_INNERERROR = "FailedOperation.InnerError"

	// 任务冲突。
	FAILEDOPERATION_JOBCONFLICT = "FailedOperation.JobConflict"

	// 任务已撤销，请重新提交任务。
	FAILEDOPERATION_JOBHASBEENCANCELED = "FailedOperation.JobHasBeenCanceled"

	// 任务不存在。
	FAILEDOPERATION_JOBNOTEXIST = "FailedOperation.JobNotExist"

	// 任务已停止处理，请重新提交任务。
	FAILEDOPERATION_JOBSTOPPROCESSING = "FailedOperation.JobStopProcessing"

	// 整个请求体太大（通常主要是图片）。
	FAILEDOPERATION_REQUESTENTITYTOOLARGE = "FailedOperation.RequestEntityTooLarge"

	// 后端服务超时。
	FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"

	// RPC请求失败，一般为算法微服务故障。
	FAILEDOPERATION_RPCFAIL = "FailedOperation.RpcFail"

	// 任务不存在。
	FAILEDOPERATION_TASKNOTEXIST = "FailedOperation.TaskNotExist"

	// 内部错误。
	INTERNALERROR = "InternalError"

	// 参数错误。
	INVALIDPARAMETER = "InvalidParameter"

	// 参数取值错误。
	INVALIDPARAMETERVALUE = "InvalidParameterValue"

	// 人脸框不合法。
	INVALIDPARAMETERVALUE_FACERECTINVALID = "InvalidParameterValue.FaceRectInvalid"

	// 第1个人脸框参数不合法。
	INVALIDPARAMETERVALUE_FACERECTINVALIDFIRST = "InvalidParameterValue.FaceRectInvalidFirst"

	// 第2个人脸框参数不合法。
	INVALIDPARAMETERVALUE_FACERECTINVALIDSECOND = "InvalidParameterValue.FaceRectInvalidSecond"

	// 第3个人脸框参数不合法。
	INVALIDPARAMETERVALUE_FACERECTINVALIDTHRID = "InvalidParameterValue.FaceRectInvalidThrid"

	// 图片为空。
	INVALIDPARAMETERVALUE_IMAGEEMPTY = "InvalidParameterValue.ImageEmpty"

	// 图片数据太大。
	INVALIDPARAMETERVALUE_IMAGESIZEEXCEED = "InvalidParameterValue.ImageSizeExceed"

	// 图片中没有人脸。
	INVALIDPARAMETERVALUE_NOFACEINPHOTO = "InvalidParameterValue.NoFaceInPhoto"

	// 参数不合法。
	INVALIDPARAMETERVALUE_PARAMETERVALUEERROR = "InvalidParameterValue.ParameterValueError"

	// URL格式不合法。
	INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"

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

	// 未知方法名。
	UNSUPPORTEDOPERATION_UNKNOWMETHOD = "UnsupportedOperation.UnknowMethod"
)
