// Copyright (c) 2017-2025 Tencent. All Rights Reserved.
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

package v20260130

const (
	// 此产品的特有错误码

	// 系统异常。
	INTERNALERROR_SYSTEMEXCEPTION = "InternalError.SystemException"

	// 参数错误。
	INVALIDPARAMETER = "InvalidParameter"

	// 必填字段缺失。
	INVALIDPARAMETER_FIELDMISSED = "InvalidParameter.FieldMissed"

	// 字段格式错误。
	INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"

	// 字段非法取值。
	INVALIDPARAMETER_INVALIDVALUE = "InvalidParameter.InvalidValue"

	// 字段长度超过最大限制。
	INVALIDPARAMETER_LENGTHEXCEED = "InvalidParameter.LengthExceed"

	// 数据未授权，请检查数据授权信息。
	INVALIDPARAMETERVALUE_DATAUNAUTHORIZED = "InvalidParameterValue.DataUnauthorized"

	// 自定义参数Key不允许重复。
	INVALIDPARAMETERVALUE_DUPLICATEDKEY = "InvalidParameterValue.DuplicatedKey"

	// 事件不存在。
	INVALIDPARAMETERVALUE_EVENTNOTEXIST = "InvalidParameterValue.EventNotExist"

	// EventTime与当前系统时间相差过大
	INVALIDPARAMETERVALUE_INVALIDEVENTTIME = "InvalidParameterValue.InvalidEventTime"

	// 租户不存在。
	INVALIDPARAMETERVALUE_TENANTNOTEXIST = "InvalidParameterValue.TenantNotExist"

	// 缺少参数错误。
	MISSINGPARAMETER = "MissingParameter"

	// 缺少参数错误，必传参数未填。
	MISSINGPARAMETER_FIELDMISSED = "MissingParameter.FieldMissed"

	// 未授权操作。
	UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
)
