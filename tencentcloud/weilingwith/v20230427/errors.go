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

package v20230427

const (
	// 此产品的特有错误码

	// 未找到此应用该api授权信息
	AUTHFAILURE_APIAUTHORIZATIONNOTFOUND = "AuthFailure.ApiAuthorizationNotFound"

	// 未找到该token信息
	AUTHFAILURE_TOKENNOTFOUND = "AuthFailure.TokenNotFound"

	// 内部错误。
	INTERNALERROR = "InternalError"

	// API网关内部错误
	INTERNALERROR_APIGATEWAYINTERNALERROR = "InternalError.ApiGatewayInternalError"

	// 该应用未关联该项目空间数据，无法获取该项目空间数据
	INTERNALERROR_APPAPINOSPACEPERMISSION = "InternalError.AppApiNoSpacePermission"

	// 未知错误
	INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"

	// 参数错误。
	INVALIDPARAMETER = "InvalidParameter"

	// 参数取值错误。
	INVALIDPARAMETERVALUE = "InvalidParameterValue"

	// 应用id非法
	INVALIDPARAMETERVALUE_INVALIDAPPLICATIONID = "InvalidParameterValue.InvalidApplicationId"

	// 错误的工作空间Id
	INVALIDPARAMETERVALUE_INVALIDWORKSPACEID = "InvalidParameterValue.InvalidWorkspaceId"
)
