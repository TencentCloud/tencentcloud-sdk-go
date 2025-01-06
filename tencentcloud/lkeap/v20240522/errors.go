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

package v20240522

const (
	// 此产品的特有错误码

	// 操作失败。
	FAILEDOPERATION = "FailedOperation"

	// FailedOperation.DownLoadError
	FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"

	// FailedOperation.FileDecodeFailed
	FAILEDOPERATION_FILEDECODEFAILED = "FailedOperation.FileDecodeFailed"

	// FailedOperation.ImageDecodeFailed
	FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"

	// 暂不支持解析该文件
	FAILEDOPERATION_NONSUPPORTPARSE = "FailedOperation.NonsupportParse"

	// 内部未知错误。
	FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"

	// FailedOperation.UnKnowFileTypeError
	FAILEDOPERATION_UNKNOWFILETYPEERROR = "FailedOperation.UnKnowFileTypeError"

	// 服务未开通。
	FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"

	// 生成识别结果文件访问url失败，请稍后重试。
	FAILEDOPERATION_UPLOADRESULTFILEFAILED = "FailedOperation.UploadResultFileFailed"

	// 内部错误。
	INTERNALERROR = "InternalError"

	// 参数错误。
	INVALIDPARAMETER = "InvalidParameter"

	// 参数值错误。
	INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"

	// 超过最大文件页数限制
	LIMITEXCEEDED_EXCEEDEDMAXPAGESERROR = "LimitExceeded.ExceededMaxPagesError"

	// 文件太大
	LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"

	// 帐号已欠费。
	RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"

	// 账号资源包耗尽。
	RESOURCEUNAVAILABLE_RESOURCEPACKAGERUNOUT = "ResourceUnavailable.ResourcePackageRunOut"

	// 计费状态异常。
	RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
)
