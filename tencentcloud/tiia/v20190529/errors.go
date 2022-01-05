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

package v20190529

const (
	// 此产品的特有错误码

	// 认证失败。
	AUTHFAILURE_INVALIDAUTHORIZATION = "AuthFailure.InvalidAuthorization"

	// 文件下载失败。
	FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"

	// 图片内容为空。
	FAILEDOPERATION_EMPTYIMAGEERROR = "FailedOperation.EmptyImageError"

	// 图片解码失败。
	FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"

	// 图片删除失败。
	FAILEDOPERATION_IMAGEDELETEFAILED = "FailedOperation.ImageDeleteFailed"

	// 图片下载错误。
	FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"

	// 超出Entity数量限制。
	FAILEDOPERATION_IMAGEENTITYCOUNTEXCEED = "FailedOperation.ImageEntityCountExceed"

	// 图库为空。
	FAILEDOPERATION_IMAGEGROUPEMPTY = "FailedOperation.ImageGroupEmpty"

	// 未找到图片信息。
	FAILEDOPERATION_IMAGENOTFOUNDINFO = "FailedOperation.ImageNotFoundInfo"

	// 不支持的图片文件。
	FAILEDOPERATION_IMAGENOTSUPPORTED = "FailedOperation.ImageNotSupported"

	// 超出图库限制。
	FAILEDOPERATION_IMAGENUMEXCEED = "FailedOperation.ImageNumExceed"

	// 未查询到结果。
	FAILEDOPERATION_IMAGESEARCHINVALID = "FailedOperation.ImageSearchInvalid"

	// base64编码后的图片数据过大。
	FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"

	// 图片不满足检测要求。
	FAILEDOPERATION_IMAGEUNQUALIFIED = "FailedOperation.ImageUnQualified"

	// url地址不合法，无法下载。
	FAILEDOPERATION_IMAGEURLINVALID = "FailedOperation.ImageUrlInvalid"

	// 内部错误。
	FAILEDOPERATION_INNERERROR = "FailedOperation.InnerError"

	// 调用计费返回失败。
	FAILEDOPERATION_INVOKECHARGEERROR = "FailedOperation.InvokeChargeError"

	// 参数为空。
	FAILEDOPERATION_PARAMETEREMPTY = "FailedOperation.ParameterEmpty"

	// 车辆识别失败。
	FAILEDOPERATION_RECOGNIZEFAILDED = "FailedOperation.RecognizeFailded"

	// 后端服务请求失败。
	FAILEDOPERATION_REQUESTERROR = "FailedOperation.RequestError"

	// 后端服务超时。
	FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"

	// RPC请求失败，一般为算法微服务故障。
	FAILEDOPERATION_RPCFAIL = "FailedOperation.RpcFail"

	// 算法服务异常，请重试。
	FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"

	// 文件太大。
	FAILEDOPERATION_TOOLARGEFILEERROR = "FailedOperation.TooLargeFileError"

	// 内部错误。
	FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"

	// 服务未开通。
	FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"

	// 未知错误。
	FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"

	// 参数取值错误。
	INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"

	// 图库简介过长。
	INVALIDPARAMETERVALUE_BRIEFTOOLONG = "InvalidParameterValue.BriefTooLong"

	// 自定义内容过长。
	INVALIDPARAMETERVALUE_CUSTOMCONTENTTOOLONG = "InvalidParameterValue.CustomContentTooLong"

	// 物品ID为空。
	INVALIDPARAMETERVALUE_ENTITYIDEMPTY = "InvalidParameterValue.EntityIdEmpty"

	// 物品ID超出长度限制。
	INVALIDPARAMETERVALUE_ENTITYIDTOOLONG = "InvalidParameterValue.EntityIdTooLong"

	// Filter参数不合法。
	INVALIDPARAMETERVALUE_FILTERINVALID = "InvalidParameterValue.FilterInvalid"

	// Filter参数过长。
	INVALIDPARAMETERVALUE_FILTERSIZEEXCEED = "InvalidParameterValue.FilterSizeExceed"

	// 图库ID已存在。
	INVALIDPARAMETERVALUE_IMAGEGROUPIDALREADYEXIST = "InvalidParameterValue.ImageGroupIdAlreadyExist"

	// 图库ID不合法。
	INVALIDPARAMETERVALUE_IMAGEGROUPIDILLEGAL = "InvalidParameterValue.ImageGroupIdIllegal"

	// 图库ID不存在。
	INVALIDPARAMETERVALUE_IMAGEGROUPIDNOTEXIST = "InvalidParameterValue.ImageGroupIdNotExist"

	// 图库ID超出长度限制。
	INVALIDPARAMETERVALUE_IMAGEGROUPIDTOOLONG = "InvalidParameterValue.ImageGroupIdTooLong"

	// 图库名称为空。
	INVALIDPARAMETERVALUE_IMAGEGROUPNAMEEMPTY = "InvalidParameterValue.ImageGroupNameEmpty"

	// 图库名称超出长度限制。
	INVALIDPARAMETERVALUE_IMAGEGROUPNAMETOOLONG = "InvalidParameterValue.ImageGroupNameTooLong"

	// 参数值错误。
	INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"

	// 返回数量不在合法范围内。
	INVALIDPARAMETERVALUE_LIMITEXCEED = "InvalidParameterValue.LimitExceed"

	// 图片中未检测到人脸。
	INVALIDPARAMETERVALUE_NOFACEINPHOTO = "InvalidParameterValue.NoFaceInPhoto"

	// 图片已经存在。
	INVALIDPARAMETERVALUE_PICNAMEALREADYEXIST = "InvalidParameterValue.PicNameAlreadyExist"

	// 图片名称为空。
	INVALIDPARAMETERVALUE_PICNAMEEMPTY = "InvalidParameterValue.PicNameEmpty"

	// 图片名称超出长度限制。
	INVALIDPARAMETERVALUE_PICNAMETOOLONG = "InvalidParameterValue.PicNameTooLong"

	// 标签数量过多。
	INVALIDPARAMETERVALUE_TAGSKEYSEXCEED = "InvalidParameterValue.TagsKeysExceed"

	// 标签值类型不合法。
	INVALIDPARAMETERVALUE_TAGSVALUEILLEGAL = "InvalidParameterValue.TagsValueIllegal"

	// 标签值长度过长。
	INVALIDPARAMETERVALUE_TAGSVALUESIZEEXCEED = "InvalidParameterValue.TagsValueSizeExceed"

	// URL格式不合法。
	INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"

	// 文件太大。
	LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"

	// 必选参数为空。
	MISSINGPARAMETER_ERRORPARAMETEREMPTY = "MissingParameter.ErrorParameterEmpty"

	// 账号已欠费。
	RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"

	// 计费状态未知，请确认是否已在控制台开通服务。
	RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"

	// 计费状态异常。
	RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
)
