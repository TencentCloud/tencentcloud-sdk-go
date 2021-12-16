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

package v20190328

const (
	// 此产品的特有错误码

	// 内部错误。
	INTERNALERROR = "InternalError"

	// 查询超时
	INTERNALERROR_TIMEOUT = "InternalError.Timeout"

	// 参数错误。
	INVALIDPARAMETER = "InvalidParameter"

	// CertMd5参数错误
	INVALIDPARAMETER_CERTMD5 = "InvalidParameter.CertMd5"

	// FileMd5参数错误
	INVALIDPARAMETER_FILEMD5 = "InvalidParameter.FileMd5"

	// FileSize参数错误
	INVALIDPARAMETER_FILESIZE = "InvalidParameter.FileSize"

	// Imei参数错误
	INVALIDPARAMETER_IMEI = "InvalidParameter.Imei"

	// 接口不存在
	INVALIDPARAMETER_INVALIDACTION = "InvalidParameter.InvalidAction"

	// IP参数错误
	INVALIDPARAMETER_IP = "InvalidParameter.Ip"

	// 包名填写错误
	INVALIDPARAMETER_PACKAGENAME = "InvalidParameter.PackageName"

	// PhoneNumber参数错误
	INVALIDPARAMETER_PHONENUMBER = "InvalidParameter.PhoneNumber"

	// QQ参数错误
	INVALIDPARAMETER_QQ = "InvalidParameter.QQ"

	// Service参数错误
	INVALIDPARAMETER_SERVICE = "InvalidParameter.Service"

	// Url参数错误
	INVALIDPARAMETER_URL = "InvalidParameter.Url"

	// Wechat参数错误
	INVALIDPARAMETER_WECHAT = "InvalidParameter.Wechat"

	// 参数取值错误。
	INVALIDPARAMETERVALUE = "InvalidParameterValue"

	// 超过配额限制。
	LIMITEXCEEDED = "LimitExceeded"

	// 缺少参数错误。
	MISSINGPARAMETER = "MissingParameter"

	// 资源不足。
	RESOURCEINSUFFICIENT = "ResourceInsufficient"

	// 资源不存在。
	RESOURCENOTFOUND = "ResourceNotFound"

	// 资源不可用。
	RESOURCEUNAVAILABLE = "ResourceUnavailable"

	// 未知参数错误。
	UNKNOWNPARAMETER = "UnknownParameter"
)
