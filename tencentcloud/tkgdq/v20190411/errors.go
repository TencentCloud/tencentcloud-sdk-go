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

package v20190411

const (
	// 此产品的特有错误码

	// 内部错误
	INTERNALERROR = "InternalError"

	// 图数据库引擎内部错误
	INTERNALERROR_GRAPHENGINEERROR = "InternalError.GraphEngineError"

	// 图查询引擎计算超时错误
	INTERNALERROR_GREMLINTIMEOUTERROR = "InternalError.GremlinTimeoutError"

	// 请求数据解析失败
	INTERNALERROR_REQUESTENCODEERROR = "InternalError.RequestEncodeError"

	// 读取请求数据内容失败
	INTERNALERROR_REQUESTPARSEERROR = "InternalError.RequestParseError"

	// 系统响应解析失败
	INTERNALERROR_RESPONSEDECODEERROR = "InternalError.ResponseDecodeError"

	// 参数错误
	INVALIDPARAMETER = "InvalidParameter"

	// 不可用的接口操作错误
	INVALIDPARAMETER_INVALIDACTIONERROR = "InvalidParameter.InvalidActionError"

	// 参数取值错误
	INVALIDPARAMETERVALUE = "InvalidParameterValue"

	// 账户余量不足错误
	LIMITEXCEEDED_INSUFFICIENTBALANCEERROR = "LimitExceeded.InsufficientBalanceError"

	// 未授权的用户请求错误
	UNAUTHORIZEDOPERATION_UNAUTHORIZEDACCOUNTERROR = "UnauthorizedOperation.UnauthorizedAccountError"
)
