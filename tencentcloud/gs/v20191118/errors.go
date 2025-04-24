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

package v20191118

const (
	// 此产品的特有错误码

	// 操作失败。
	FAILEDOPERATION = "FailedOperation"

	// 应用创建失败。
	FAILEDOPERATION_APPLICATIONCREATEFAIL = "FailedOperation.ApplicationCreateFail"

	// 应用锁定失败。
	FAILEDOPERATION_APPLICATIONLOCKFAIL = "FailedOperation.ApplicationLockFail"

	// 应用不存在。
	FAILEDOPERATION_APPLICATIONNOTFIND = "FailedOperation.ApplicationNotFind"

	// 游戏锁定失败。
	FAILEDOPERATION_GAMELOCKFAIL = "FailedOperation.GameLockFail"

	// 游戏不存在。
	FAILEDOPERATION_GAMENOTFIND = "FailedOperation.GameNotFind"

	// 锁定机器超时或未调用TrylockWorker。
	FAILEDOPERATION_LOCKTIMEOUT = "FailedOperation.LockTimeout"

	// 处理超时。
	FAILEDOPERATION_PROCESSTIMEOUT = "FailedOperation.ProcessTimeout"

	// 请降低访问频率。
	FAILEDOPERATION_SLOWDOWN = "FailedOperation.SlowDown"

	// 操作超时。
	FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"

	// 请求太频繁。
	FAILEDOPERATION_TOOFREQUENTLY = "FailedOperation.TooFrequently"

	// 内部错误。
	INTERNALERROR = "InternalError"

	// 参数错误。
	INVALIDPARAMETER = "InvalidParameter"

	// Json解析失败。
	INVALIDPARAMETER_JSONPARSEERROR = "InvalidParameter.JsonParseError"

	// 参数取值错误。
	INVALIDPARAMETERVALUE = "InvalidParameterValue"

	// 【多人游戏】对应的角色人数超过限制。
	LIMITEXCEEDED_ROLE = "LimitExceeded.Role"

	// 缺少参数错误。
	MISSINGPARAMETER = "MissingParameter"

	// 操作被拒绝。
	OPERATIONDENIED = "OperationDenied"

	// 超出应用数量限制。
	OPERATIONDENIED_APPLICATIONLIMITEXCEEDED = "OperationDenied.ApplicationLimitExceeded"

	// 版本正在创建。
	OPERATIONDENIED_VERSIONCREATING = "OperationDenied.VersionCreating"

	// 超过版本数量限制。
	OPERATIONDENIED_VERSIONLIMITEXCEEDED = "OperationDenied.VersionLimitExceeded"

	// 资源不存在。
	RESOURCENOTFOUND = "ResourceNotFound"

	// 没有空闲机器。
	RESOURCENOTFOUND_NOIDLE = "ResourceNotFound.NoIdle"

	// 会话未找到。
	RESOURCENOTFOUND_SESSIONNOTFOUND = "ResourceNotFound.SessionNotFound"

	// 访问并发实例失败。
	RESOURCEUNAVAILABLE_ACCESSFAILED = "ResourceUnavailable.AccessFailed"

	// 机器还在初始化中。
	RESOURCEUNAVAILABLE_INITIALIZATION = "ResourceUnavailable.Initialization"

	// 操作不支持。
	UNSUPPORTEDOPERATION = "UnsupportedOperation"

	// 机器未运行。
	UNSUPPORTEDOPERATION_NOTRUNNING = "UnsupportedOperation.NotRunning"

	// 退出游戏中。
	UNSUPPORTEDOPERATION_STOPPING = "UnsupportedOperation.Stopping"
)
