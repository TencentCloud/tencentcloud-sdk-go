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

package v20230901

const (
	// 此产品的特有错误码

	// 操作失败。
	FAILEDOPERATION = "FailedOperation"

	// 引擎层请求超时；请稍后重试。
	FAILEDOPERATION_ENGINEREQUESTTIMEOUT = "FailedOperation.EngineRequestTimeout"

	// 引擎层内部错误；请稍后重试。
	FAILEDOPERATION_ENGINESERVERERROR = "FailedOperation.EngineServerError"

	// 引擎层请求超过限额；请稍后重试。
	FAILEDOPERATION_ENGINESERVERLIMITEXCEEDED = "FailedOperation.EngineServerLimitExceeded"

	// 免费资源包余量已用尽，请购买资源包或开通后付费。
	FAILEDOPERATION_FREERESOURCEPACKEXHAUSTED = "FailedOperation.FreeResourcePackExhausted"

	// 资源包余量已用尽，请购买资源包或开通后付费。
	FAILEDOPERATION_RESOURCEPACKEXHAUSTED = "FailedOperation.ResourcePackExhausted"

	// 服务未开通，请前往控制台申请试用。
	FAILEDOPERATION_SERVICENOTACTIVATED = "FailedOperation.ServiceNotActivated"

	// 用户主动停服。
	FAILEDOPERATION_SERVICESTOP = "FailedOperation.ServiceStop"

	// 欠费停服。
	FAILEDOPERATION_SERVICESTOPARREARS = "FailedOperation.ServiceStopArrears"

	// 内部错误。
	INTERNALERROR = "InternalError"

	// 参数错误。
	INVALIDPARAMETER = "InvalidParameter"

	// 参数取值错误。
	INVALIDPARAMETERVALUE = "InvalidParameterValue"

	// 模型不存在。
	INVALIDPARAMETERVALUE_MODEL = "InvalidParameterValue.Model"

	// 超过配额限制。
	LIMITEXCEEDED = "LimitExceeded"
)
