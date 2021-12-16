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

package v20200727

const (
	// 此产品的特有错误码

	// 操作失败。
	FAILEDOPERATION = "FailedOperation"

	// 数据库查询错误。
	FAILEDOPERATION_DATABASEFINDERROR = "FailedOperation.DatabaseFindError"

	// 文件不存在。
	FAILEDOPERATION_FILENOTEXISTS = "FailedOperation.FileNotExists"

	// 权限不足。
	FAILEDOPERATION_INSUFFICIENTPRIVILEGE = "FailedOperation.InsufficientPrivilege"

	// 指定的服务不存在。
	FAILEDOPERATION_SERVICENOTFOUND = "FailedOperation.ServiceNotFound"

	// 内部错误。
	INTERNALERROR = "InternalError"

	// 参数错误。
	INVALIDPARAMETER = "InvalidParameter"

	// 参数值不能为空。
	INVALIDPARAMETER_EMPTYPARAMETERS = "InvalidParameter.EmptyParameters"

	// 文件格式解析失败。
	INVALIDPARAMETER_INVALIDFILEFORMAT = "InvalidParameter.InvalidFileFormat"

	// 文件大小不符合要求。
	INVALIDPARAMETER_INVALIDFILESIZE = "InvalidParameter.InvalidFileSize"

	// 参数取值错误。
	INVALIDPARAMETER_INVALIDPARAMETERVALUE = "InvalidParameter.InvalidParameterValue"

	// 参数解析失败。
	INVALIDPARAMETER_REQUESTPARSEERROR = "InvalidParameter.RequestParseError"
)
