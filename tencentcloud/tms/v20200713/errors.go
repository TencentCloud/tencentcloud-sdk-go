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

package v20200713

const (
	// 此产品的特有错误码

	// CAM签名/鉴权错误。
	AUTHFAILURE = "AuthFailure"

	// 内部错误。
	INTERNALERROR = "InternalError"

	// 请求超时。
	INTERNALERROR_ERRTEXTTIMEOUT = "InternalError.ErrTextTimeOut"

	// 参数错误。
	INVALIDPARAMETER = "InvalidParameter"

	// 错误的action。
	INVALIDPARAMETER_ERRACTION = "InvalidParameter.ErrAction"

	// 请求的文本长度过长。
	INVALIDPARAMETER_ERRTEXTCONTENTLEN = "InvalidParameter.ErrTextContentLen"

	// 文本类型错误，需要base64的文本。
	INVALIDPARAMETER_ERRTEXTCONTENTTYPE = "InvalidParameter.ErrTextContentType"

	// 参数内容格式错误。
	INVALIDPARAMETER_INVALIDPARAMETERCONTENT = "InvalidParameter.InvalidParameterContent"

	// 参数取值错误。
	INVALIDPARAMETERVALUE = "InvalidParameterValue"

	// 请求的文本格式错误（需要base64编码格式的文本）。
	INVALIDPARAMETERVALUE_ERRTEXTCONTENTTYPE = "InvalidParameterValue.ErrTextContentType"

	// EvilContent参数太长。
	INVALIDPARAMETERVALUE_INVALIDEVILCONTENT = "InvalidParameterValue.InvalidEvilContent"

	// EvilType取值错误。
	INVALIDPARAMETERVALUE_INVALIDEVILTYPE = "InvalidParameterValue.InvalidEvilType"

	// ReportedAccount参数错误。
	INVALIDPARAMETERVALUE_INVALIDREPORTEDACCOUNT = "InvalidParameterValue.InvalidReportedAccount"

	// ReportedAccount参数太长。
	INVALIDPARAMETERVALUE_INVALIDREPORTEDACCOUNTLENGTH = "InvalidParameterValue.InvalidReportedAccountLength"

	// ReportedAccountType参数错误。
	INVALIDPARAMETERVALUE_INVALIDREPORTEDACCOUNTTYPE = "InvalidParameterValue.InvalidReportedAccountType"

	// SenderAccount参数错误。
	INVALIDPARAMETERVALUE_INVALIDSENDERACCOUNT = "InvalidParameterValue.InvalidSenderAccount"

	// SenderAccount参数太长。
	INVALIDPARAMETERVALUE_INVALIDSENDERACCOUNTLENGTH = "InvalidParameterValue.InvalidSenderAccountLength"

	// SenderAccountType参数错误。
	INVALIDPARAMETERVALUE_INVALIDSENDERACCOUNTTYPE = "InvalidParameterValue.InvalidSenderAccountType"

	// SenderIP参数错误。
	INVALIDPARAMETERVALUE_INVALIDSENDERIP = "InvalidParameterValue.InvalidSenderIP"

	// 超过配额限制。
	LIMITEXCEEDED = "LimitExceeded"

	// 缺少参数错误。
	MISSINGPARAMETER = "MissingParameter"

	// 资源被占用。
	RESOURCEINUSE = "ResourceInUse"

	// 资源不存在。
	RESOURCENOTFOUND = "ResourceNotFound"

	// 未获取到接口授权。
	UNAUTHORIZEDOPERATION_UNAUTHORIZED = "UnauthorizedOperation.Unauthorized"
)
