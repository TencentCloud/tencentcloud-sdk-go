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

package v20191209

const (
	// 此产品的特有错误码

	// 验证码签名错误。
	AUTHFAILURE_CAPSIGERROR = "AuthFailure.CapSigError"

	// 请求过期。
	AUTHFAILURE_EXPIRED = "AuthFailure.Expired"

	// 内部错误。
	INTERNALERROR = "InternalError"

	// 业务后端逻辑错误。
	INTERNALERROR_BACKENDLOGICERROR = "InternalError.BackendLogicError"

	// 后端错误。
	INTERNALERROR_SERVERERROR = "InternalError.ServerError"

	// 关键词后端错误。
	INTERNALERROR_SIGNBACKENDERROR = "InternalError.SignBackendError"

	// 参数错误。
	INVALIDPARAMETER = "InvalidParameter"

	// DevToken不可用、过期或与当前场景不匹配。
	INVALIDPARAMETER_DEVTOKENINVALID = "InvalidParameter.DevTokenInvalid"

	// 参数错误。
	INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"

	// 参数校验错误。
	INVALIDPARAMETER_PARAMETERERROR = "InvalidParameter.ParameterError"

	// token不可用或过期。
	INVALIDPARAMETER_TOKENINVALID = "InvalidParameter.TokenInvalid"

	// URL错误。
	INVALIDPARAMETER_URLERROR = "InvalidParameter.UrlError"

	// 版本错误。
	INVALIDPARAMETER_VERSIONERROR = "InvalidParameter.VersionError"

	// 参数取值错误。
	INVALIDPARAMETERVALUE = "InvalidParameterValue"

	// 请求body错误。
	INVALIDPARAMETERVALUE_BADBODY = "InvalidParameterValue.BadBody"

	// 请求包过大。
	INVALIDPARAMETERVALUE_BODYTOOLARGE = "InvalidParameterValue.BodyTooLarge"

	// 验证码不匹配。
	INVALIDPARAMETERVALUE_CAPMISMATCH = "InvalidParameterValue.CapMisMatch"

	// Http方法错误。
	INVALIDPARAMETERVALUE_HTTPMETHODERROR = "InvalidParameterValue.HttpMethodError"

	// 超过配额限制。
	LIMITEXCEEDED = "LimitExceeded"

	// 超过配额。
	LIMITEXCEEDED_FREQCNT = "LimitExceeded.FreqCnt"

	// 超过配额（IP）。
	LIMITEXCEEDED_IPFREQCNT = "LimitExceeded.IpFreqCnt"

	// 关键词频控。
	LIMITEXCEEDED_KEYFREQCNT = "LimitExceeded.KeyFreqCnt"

	// 重放攻击。
	LIMITEXCEEDED_REPLAYATTACK = "LimitExceeded.ReplayAttack"

	// 接口不存在。
	RESOURCENOTFOUND_INTERFACENOTFOUND = "ResourceNotFound.InterfaceNotFound"

	// 资源不可用。
	RESOURCEUNAVAILABLE = "ResourceUnavailable"

	// devid 暂未生成。
	RESOURCEUNAVAILABLE_NOTALLREADY = "ResourceUnavailable.NotAllready"

	// 未开通服务权限。
	RESOURCEUNAVAILABLE_PERMISSIONDENIED = "ResourceUnavailable.PermissionDenied"

	// 鉴权失败。
	UNAUTHORIZEDOPERATION_AUTHFAILED = "UnauthorizedOperation.AuthFailed"

	// 密钥不存在。
	UNKNOWNPARAMETER_SECRETIDNOTEXISTS = "UnknownParameter.SecretIdNotExists"
)
