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

package v20231130

const (
	// 此产品的特有错误码

	// 操作失败。
	FAILEDOPERATION = "FailedOperation"

	// 文件下载失败。
	FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"

	// 文件解码失败
	FAILEDOPERATION_FILEDECODEFAILED = "FailedOperation.FileDecodeFailed"

	// 图片解码失败。
	FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"

	// 图片尺寸过大，请参考输入参数中关于图片大小限制的说明。
	FAILEDOPERATION_IMAGESIZETOOLARGE = "FailedOperation.ImageSizeTooLarge"

	// OCR识别失败。
	FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"

	// 未知错误。
	FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"

	// 服务未开通。
	FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"

	// 参数错误。
	INVALIDPARAMETER = "InvalidParameter"

	// 参数值错误。
	INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"

	// 缺少参数错误。
	MISSINGPARAMETER = "MissingParameter"

	// 账号已欠费。
	RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"

	// 账号资源包耗尽。
	RESOURCEUNAVAILABLE_RESOURCEPACKAGERUNOUT = "ResourceUnavailable.ResourcePackageRunOut"

	// 计费状态异常。
	RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
)
