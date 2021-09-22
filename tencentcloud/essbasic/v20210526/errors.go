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

package v20210526

const (
	// 此产品的特有错误码

	// 操作失败。
	FAILEDOPERATION = "FailedOperation"

	// 内部错误。
	INTERNALERROR = "InternalError"

	// 其他API错误。
	INTERNALERROR_API = "InternalError.Api"

	// 数据库错误。
	INTERNALERROR_DB = "InternalError.Db"

	// 解密错误。
	INTERNALERROR_DECRYPTION = "InternalError.Decryption"

	// 加密错误。
	INTERNALERROR_ENCRYPTION = "InternalError.Encryption"

	// 第三方错误。
	INTERNALERROR_THIRDPARTY = "InternalError.ThirdParty"

	// 参数错误。
	INVALIDPARAMETER = "InvalidParameter"

	// 数据不存在。
	INVALIDPARAMETER_DATANOTFOUND = "InvalidParameter.DataNotFound"

	// 参数错误。
	INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"

	// 参数取值错误。
	INVALIDPARAMETERVALUE = "InvalidParameterValue"

	// 缺少参数错误。
	MISSINGPARAMETER = "MissingParameter"

	// 操作被拒绝。
	OPERATIONDENIED = "OperationDenied"

	// 应用号已被禁止。
	OPERATIONDENIED_BANNEDAPPLICATION = "OperationDenied.BannedApplication"

	// 没有API权限。
	OPERATIONDENIED_NOAPIAUTH = "OperationDenied.NoApiAuth"

	// 用户与企业不对应。
	OPERATIONDENIED_USERNOTINORGANIZATION = "OperationDenied.UserNotInOrganization"

	// 资源不存在。
	RESOURCENOTFOUND = "ResourceNotFound"

	// 应用号不存在。
	RESOURCENOTFOUND_APPLICATION = "ResourceNotFound.Application"

	// 模板不存在。
	RESOURCENOTFOUND_TEMPLATE = "ResourceNotFound.Template"

	// 未授权操作。
	UNAUTHORIZEDOPERATION = "UnauthorizedOperation"

	// 未知参数错误。
	UNKNOWNPARAMETER = "UnknownParameter"

	// 操作不支持。
	UNSUPPORTEDOPERATION = "UnsupportedOperation"
)
