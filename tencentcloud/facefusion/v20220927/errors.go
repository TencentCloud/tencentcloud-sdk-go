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

package v20220927

const (
	// 此产品的特有错误码

	// 操作失败。
	FAILEDOPERATION = "FailedOperation"

	// 余额不足，开通失败，请充值后再开通。
	FAILEDOPERATION_BALANCEINSUFFICIENT = "FailedOperation.BalanceInsufficient"

	// 人脸因太小被过滤，建议人脸分辨率不小于34*34。
	FAILEDOPERATION_FACESIZETOOSMALL = "FailedOperation.FaceSizeTooSmall"

	// 素材未经过审核。
	FAILEDOPERATION_FUSEMATERIALNOTAUTH = "FailedOperation.FuseMaterialNotAuth"

	// 素材不存在。
	FAILEDOPERATION_FUSEMATERIALNOTEXIST = "FailedOperation.FuseMaterialNotExist"

	// 人脸检测-图片解码失败。
	FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"

	// 图片下载失败。
	FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"

	// 素材尺寸超过1080*1080像素。
	FAILEDOPERATION_IMAGEPIXELEXCEED = "FailedOperation.ImagePixelExceed"

	// 图片短边分辨率小于64。
	FAILEDOPERATION_IMAGERESOLUTIONTOOSMALL = "FailedOperation.ImageResolutionTooSmall"

	// 输入图片base64数据大小超过5M。
	FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"

	// 图片尺寸过大或者过小；不满足算法要求。
	FAILEDOPERATION_IMAGESIZEINVALID = "FailedOperation.ImageSizeInvalid"

	// 服务内部错误。
	FAILEDOPERATION_INNERERROR = "FailedOperation.InnerError"

	// 无法检测出人脸, 人脸框配准分低于阈值。
	FAILEDOPERATION_NOFACEDETECTED = "FailedOperation.NoFaceDetected"

	// 参数字段或者值有误。
	FAILEDOPERATION_PARAMETERVALUEERROR = "FailedOperation.ParameterValueError"

	// 后端服务超时。
	FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"

	// RPC请求失败，一般为算法微服务故障。
	FAILEDOPERATION_RPCFAIL = "FailedOperation.RpcFail"

	// 系统内部错误。
	FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"

	// 素材人脸ID不存在。
	FAILEDOPERATION_TEMPLATEFACEIDNOTEXIST = "FailedOperation.TemplateFaceIDNotExist"

	// 内部错误。
	FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"

	// 未知错误。
	FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"

	// 未查找到活动id。
	INVALIDPARAMETERVALUE_ACTIVITYIDNOTFOUND = "InvalidParameterValue.ActivityIdNotFound"

	// 人脸框参数有误或者人脸框太小。
	INVALIDPARAMETERVALUE_FACERECTPARAMETERVALUEERROR = "InvalidParameterValue.FaceRectParameterValueError"

	// 未查找到素材Id。
	INVALIDPARAMETERVALUE_MATERIALIDNOTFOUND = "InvalidParameterValue.MaterialIdNotFound"

	// 参数字段或者值有误。
	INVALIDPARAMETERVALUE_PARAMETERVALUEERROR = "InvalidParameterValue.ParameterValueError"

	// URL格式不合法。
	INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"

	// 请求的次数超过了频率限制。
	REQUESTLIMITEXCEEDED = "RequestLimitExceeded"

	// 资源不足。
	RESOURCEINSUFFICIENT = "ResourceInsufficient"

	// 资源不存在。
	RESOURCENOTFOUND = "ResourceNotFound"

	// 账号已被冻结。
	RESOURCEUNAVAILABLE_FREEZE = "ResourceUnavailable.Freeze"

	// 账号已欠费。
	RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"

	// 服务正在开通中，请稍等。
	RESOURCEUNAVAILABLE_ISOPENING = "ResourceUnavailable.IsOpening"

	// 计费状态未知，请确认是否已在控制台开通服务。
	RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"

	// 资源已被回收。
	RESOURCEUNAVAILABLE_RECOVER = "ResourceUnavailable.Recover"

	// 账号已停服。
	RESOURCEUNAVAILABLE_STOPUSING = "ResourceUnavailable.StopUsing"
)
