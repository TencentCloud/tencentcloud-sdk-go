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

package v20210728

const (
	// 此产品的特有错误码

	// CAM签名/鉴权错误。
	AUTHFAILURE = "AuthFailure"

	// 操作失败。
	FAILEDOPERATION = "FailedOperation"

	// 获取标签失败。
	FAILEDOPERATION_ACCESSTAGFAIL = "FailedOperation.AccessTagFail"

	// 数据库查询失败。
	FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"

	// 数据库创建记录失败。
	FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"

	// 数据库记录更新失败。
	FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"

	// 任务状态不是运行中。
	FAILEDOPERATION_JOBSTATUSNOTRUNNING = "FailedOperation.JobStatusNotRunning"

	// 任务中没有task。
	FAILEDOPERATION_NOTASKSINJOB = "FailedOperation.NoTasksInJob"

	// 当前环境不支持。
	FAILEDOPERATION_NOTSUPPORTEDINENV = "FailedOperation.NotSupportedInEnv"

	// 资源不存在。
	FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"

	// 请求发送失败。
	FAILEDOPERATION_SENDREQUEST = "FailedOperation.SendRequest"

	// 内部错误。
	INTERNALERROR = "InternalError"

	// 参数错误。
	INVALIDPARAMETER = "InvalidParameter"

	// 参数取值错误。
	INVALIDPARAMETERVALUE = "InvalidParameterValue"

	// 超过配额限制。
	LIMITEXCEEDED = "LimitExceeded"

	// 资源不存在。
	RESOURCENOTFOUND = "ResourceNotFound"
)
