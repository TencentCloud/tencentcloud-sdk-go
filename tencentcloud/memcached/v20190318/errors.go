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

package v20190318

const (
	// 此产品的特有错误码

	// 数据库操作异常
	FAILEDOPERATION_DBOPERATIONFAILED = "FailedOperation.DbOperationFailed"

	// 系统内部错误
	FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"

	// 内部错误
	INTERNALERROR_INTERNALERROR = "InternalError.InternalError"

	// 业务参数错误
	INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"

	// 参数为空
	MISSINGPARAMETER_EMPTYPARAM = "MissingParameter.EmptyParam"

	// 未提供有效的账号
	RESOURCENOTFOUND_ACCOUNTDOESNOTEXISTS = "ResourceNotFound.AccountDoesNotExists"

	// Cam鉴权失败
	UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
)
