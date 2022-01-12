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

package v20200513

const (
	// 此产品的特有错误码

	// CAM签名/鉴权错误。
	AUTHFAILURE = "AuthFailure"

	// 操作失败。
	FAILEDOPERATION = "FailedOperation"

	// 内部错误。
	INTERNALERROR = "InternalError"

	// 参数错误。
	INVALIDPARAMETER = "InvalidParameter"

	// 参数取值错误。
	INVALIDPARAMETERVALUE = "InvalidParameterValue"

	// 超过配额限制。
	LIMITEXCEEDED = "LimitExceeded"

	// 图片使用期限已过期。
	LIMITEXCEEDED_ORDEREXPIREDERROR = "LimitExceeded.OrderExpiredError"

	// 下单频率超过限制。
	LIMITEXCEEDED_ORDERLIMITERROR = "LimitExceeded.OrderLimitError"

	// 资源被占用。
	RESOURCEINUSE = "ResourceInUse"

	// 资源不存在。
	RESOURCENOTFOUND = "ResourceNotFound"

	// 图片涉嫌违禁内容
	RESOURCENOTFOUND_SENSITIVEIMAGE = "ResourceNotFound.SensitiveImage"

	// 包含敏感词汇。
	RESOURCENOTFOUND_SENSITIVESEARCH = "ResourceNotFound.SensitiveSearch"

	// 资源不可用。
	RESOURCEUNAVAILABLE = "ResourceUnavailable"
)
