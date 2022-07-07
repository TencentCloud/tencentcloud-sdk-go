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

package v20220110

const (
	// 此产品的特有错误码

	// 操作失败。
	FAILEDOPERATION = "FailedOperation"

	// 未申请并发或申请后超时。
	FAILEDOPERATION_LOCKTIMEOUT = "FailedOperation.LockTimeout"

	// 处理超时。
	FAILEDOPERATION_PROCESSTIMEOUT = "FailedOperation.ProcessTimeout"

	// 请降低访问频率。
	FAILEDOPERATION_SLOWDOWN = "FailedOperation.SlowDown"

	// 内部错误。
	INTERNALERROR = "InternalError"

	// Json 解析失败。
	INVALIDPARAMETER_JSONPARSEERROR = "InvalidParameter.JsonParseError"

	// 参数取值错误。
	INVALIDPARAMETERVALUE = "InvalidParameterValue"

	// 操作被拒绝。
	OPERATIONDENIED = "OperationDenied"

	// 没有空闲并发。
	RESOURCENOTFOUND_NOIDLE = "ResourceNotFound.NoIdle"

	// 未找到会话。
	RESOURCENOTFOUND_SESSIONNOTFOUND = "ResourceNotFound.SessionNotFound"

	// 访问并发实例失败。
	RESOURCEUNAVAILABLE_ACCESSFAILED = "ResourceUnavailable.AccessFailed"

	// 初始化中。
	RESOURCEUNAVAILABLE_INITIALIZATION = "ResourceUnavailable.Initialization"

	// 销毁会话中。
	UNSUPPORTEDOPERATION_STOPPING = "UnsupportedOperation.Stopping"
)
