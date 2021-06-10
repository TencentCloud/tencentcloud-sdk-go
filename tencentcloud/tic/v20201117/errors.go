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

package v20201117

const (
	// 此产品的特有错误码

	// 操作失败。
	FAILEDOPERATION = "FailedOperation"

	// HTTP请求错误。
	FAILEDOPERATION_HTTPREQUESTERROR = "FailedOperation.HttpRequestError"

	// 资源不存在。
	FAILEDOPERATION_NOTEXIST = "FailedOperation.NotExist"

	// 内部错误。
	INTERNALERROR = "InternalError"

	// 使用的云存储错误。
	INTERNALERROR_COSERROR = "InternalError.CosError"

	// 数据库操作错误。
	INTERNALERROR_DBERROR = "InternalError.DbError"

	// 使用的消息队列错误。
	INTERNALERROR_MQERROR = "InternalError.MqError"

	// 系统错误。
	INTERNALERROR_SYSTEM = "InternalError.System"

	// 参数错误。
	INVALIDPARAMETER = "InvalidParameter"

	// 参数错误。
	INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"

	// 禁止操作。
	UNSUPPORTEDOPERATION_FORBIDOP = "UnsupportedOperation.ForbidOp"
)
