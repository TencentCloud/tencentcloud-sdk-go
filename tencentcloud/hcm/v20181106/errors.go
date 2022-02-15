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

package v20181106

const (
	// 此产品的特有错误码

	// 计费次数统计失败。
	FAILEDOPERATION_CHARGECOUNTERROR = "FailedOperation.ChargeCountError"

	// 引擎请求失败。
	INTERNALERROR_ENGINEREQUESTFAILED = "InternalError.EngineRequestFailed"

	// 引擎识别失败。
	INTERNALERROR_ENGINERESULTERROR = "InternalError.EngineResultError"

	// 初始化参数错误。
	INTERNALERROR_INITIALPARAMETERERROR = "InternalError.InitialParameterError"

	// 服务器内部错误。
	INTERNALERROR_SERVERINTERNALERROR = "InternalError.ServerInternalError"

	// 无法找到图像，请确认Image参数与Url参数无误。
	INVALIDPARAMETERVALUE_CANNOTFINDIMAGEERROR = "InvalidParameterValue.CannotFindImageError"

	// 评估之前没有初始化或已过期。
	INVALIDPARAMETERVALUE_CANNOTFINDSESSION = "InvalidParameterValue.CannotFindSession"

	// Image参数为空，请重新填写。
	INVALIDPARAMETERVALUE_EMPTYIMAGEERROR = "InvalidParameterValue.EmptyImageError"

	// 必填参数为空，请核实传入参数。
	INVALIDPARAMETERVALUE_EMPTYINPUTERROR = "InvalidParameterValue.EmptyInputError"

	// 超过图片大小限制，请裁剪后再评估。
	INVALIDPARAMETERVALUE_EXCEEDDOWNLOADIMAGESIZEERROR = "InvalidParameterValue.ExceedDownloadImageSizeError"

	// 图像解码错误，请重新核实图像信息。
	INVALIDPARAMETERVALUE_FAILDECODEERROR = "InvalidParameterValue.FailDecodeError"

	// 图片下载失败，请核实图像下载地址。
	INVALIDPARAMETERVALUE_FAILDOWNLOADIMAGEERROR = "InvalidParameterValue.FailDownloadImageError"

	// 算式信息获取失败，请核实图像内容。
	INVALIDPARAMETERVALUE_FAILRECOGNIZEERROR = "InvalidParameterValue.FailRecognizeError"

	// 输入的图片为非速算图片，请核实图片中存在手写数学公式，且背景没有太多干扰。
	INVALIDPARAMETERVALUE_INVALIDIMAGEERROR = "InvalidParameterValue.InvalidImageError"

	// 无法找到用户，请确认已在控制台开通服务并使用了正确的HCMAPPID。
	RESOURCENOTFOUND_CANNOTFINDUSER = "ResourceNotFound.CannotFindUser"
)
