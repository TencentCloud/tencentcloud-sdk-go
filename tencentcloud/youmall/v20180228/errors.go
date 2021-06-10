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

package v20180228

const (
	// 此产品的特有错误码

	// 登录名称已存在。
	FAILEDOPERATION_ACCOUNTEXIST = "FailedOperation.AccountExist"

	// 人脸质量不达标。
	FAILEDOPERATION_BADFACEQUALITY = "FailedOperation.BadFaceQuality"

	// 图片特征提取失败。
	FAILEDOPERATION_EXTRACTFEATUREERROR = "FailedOperation.ExtractFeatureError"

	// 未检测到人脸。
	FAILEDOPERATION_FACENOTFOUND = "FailedOperation.FaceNotFound"

	// 人脸大小过小。
	FAILEDOPERATION_FACESIZEERROR = "FailedOperation.FaceSizeError"

	// 人脸已存在。
	FAILEDOPERATION_HASEXISTPERSON = "FailedOperation.HasExistPerson"

	// 内部错误。
	FAILEDOPERATION_INNERERROR = "FailedOperation.InnerError"

	// 检测到多个人脸。
	FAILEDOPERATION_MULTIFACEDETECTED = "FailedOperation.MultiFaceDetected"

	// 需要预授权。
	FAILEDOPERATION_NEEDGRANTROLEFIRST = "FailedOperation.NeedGrantRoleFirst"

	// 无数据。
	FAILEDOPERATION_NODATA = "FailedOperation.NoData"

	// 未检索到顾客。
	FAILEDOPERATION_NOPERSON = "FailedOperation.NoPerson"

	// 商户无权限。
	FAILEDOPERATION_NORIGHT = "FailedOperation.NoRight"

	// 客户门店编码未匹配。
	FAILEDOPERATION_NOTMATCHSHOPCODE = "FailedOperation.NotMatchShopCode"

	// 图片检索失败。
	FAILEDOPERATION_OTHERS = "FailedOperation.Others"

	// 请求处理参数错误。
	FAILEDOPERATION_PARAMETERERROR = "FailedOperation.ParameterError"

	// 处理失败。
	FAILEDOPERATION_PROCESSFAIL = "FailedOperation.ProcessFail"

	// 内部错误
	INTERNALERROR = "InternalError"

	// 输入数据错误。
	INTERNALERROR_DATAERROR = "InternalError.DataError"

	// 数据已存在。
	INTERNALERROR_DATAHASEXISTS = "InternalError.DataHasExists"

	// 内部系统处理失败。
	INTERNALERROR_INNERSERVERFAILED = "InternalError.InnerServerFailed"

	// 元数据处理失败。
	INTERNALERROR_METADATAOPFAILED = "InternalError.MetaDataOpFailed"

	// 用户不存在。
	INTERNALERROR_USERNOTEXIST = "InternalError.UserNotExist"

	// 参数错误
	INVALIDPARAMETER = "InvalidParameter"

	// 参数取值错误
	INVALIDPARAMETERVALUE = "InvalidParameterValue"

	// 请求JSON解析错误。
	INVALIDPARAMETERVALUE_JSONPARSEERR = "InvalidParameterValue.JsonParseErr"

	// 未授权操作
	UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
)
