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

package v20220601

const (
	// 此产品的特有错误码

	// 查询数据失败。
	FAILEDOPERATION_QUERYDATA = "FailedOperation.QueryData"

	// 内部错误，数据库异常。
	INTERNALERROR_DATABASEEXCEPTION = "InternalError.DatabaseException"

	// 内部未知错误。
	INTERNALERROR_UNKNOWN = "InternalError.Unknown"

	// 请求参数错误。
	INVALIDPARAMETER_REQUESTPARAM = "InvalidParameter.RequestParam"

	// 缺少公共参数。
	MISSINGPARAMETER_COMMONPARAM = "MissingParameter.CommonParam"

	// 资源不存在。
	RESOURCENOTFOUND_NOTFOUND = "ResourceNotFound.NotFound"

	// 未授权的操作。
	UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
)
