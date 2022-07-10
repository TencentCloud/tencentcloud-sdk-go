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

package v20201229

const (
	// 此产品的特有错误码

	// 请求超时。
	INTERNALERROR_ERRTEXTTIMEOUT = "InternalError.ErrTextTimeOut"

	// 错误的action。
	INVALIDPARAMETER_ERRACTION = "InvalidParameter.ErrAction"

	// 请求的文本长度过长。
	INVALIDPARAMETER_ERRTEXTCONTENTLEN = "InvalidParameter.ErrTextContentLen"

	// 文本类型错误，需要base64的文本。
	INVALIDPARAMETER_ERRTEXTCONTENTTYPE = "InvalidParameter.ErrTextContentType"

	// InvalidParameter.ParameterError
	INVALIDPARAMETER_PARAMETERERROR = "InvalidParameter.ParameterError"

	// FileContent不可用，传入的Base64编码无法转换成标准utf8内容。
	INVALIDPARAMETERVALUE_ERRFILECONTENT = "InvalidParameterValue.ErrFileContent"

	// 请求的文本长度超过限制。
	INVALIDPARAMETERVALUE_ERRTEXTCONTENTLEN = "InvalidParameterValue.ErrTextContentLen"

	// 请求的文本格式错误（需要base64编码格式的文本）。
	INVALIDPARAMETERVALUE_ERRTEXTCONTENTTYPE = "InvalidParameterValue.ErrTextContentType"

	// 未开通权限/无有效套餐包/账号已欠费。
	UNAUTHORIZEDOPERATION_UNAUTHORIZED = "UnauthorizedOperation.Unauthorized"
)
