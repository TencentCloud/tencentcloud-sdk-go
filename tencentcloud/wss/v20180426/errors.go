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

package v20180426

const (
	// 此产品的特有错误码

	// 证书类型错误。
	FAILEDOPERATION_CERTINVALIDPARAM = "FailedOperation.CertInvalidParam"

	// 证书与密钥不对应。
	FAILEDOPERATION_CERTMISMATCH = "FailedOperation.CertMismatch"

	// 证书格式错误。
	FAILEDOPERATION_INVALIDCERT = "FailedOperation.InvalidCert"

	// 内部错误。
	INTERNALERROR = "InternalError"

	// 网络错误，请稍后重试。
	INTERNALERROR_DOMAININTERNALERROR = "InternalError.DomainInternalError"

	// 参数错误。
	INVALIDPARAMETER = "InvalidParameter"

	// 请求类型错误。
	INVALIDPARAMETER_REPTYPEISINVALID = "InvalidParameter.RepTypeIsInvalid"
)
