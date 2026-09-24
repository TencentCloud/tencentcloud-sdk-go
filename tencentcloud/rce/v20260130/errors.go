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

	// 请求字段缺失
	INTERNALERROR_TURINGFIELDMISSED = "InternalError.TuringFieldMissed"

	// 内部错误
	INTERNALERROR_TURINGINTERNALERROR = "InternalError.TuringInternalError"

	// 时间戳字段不合法
	INTERNALERROR_TURINGINVALIDTIMESTAMP = "InternalError.TuringInvalidTimestamp"

	// 已购配额已用完
	INTERNALERROR_TURINGLIMITEXCEEDED = "InternalError.TuringLimitExceeded"

	// QPS超过频率限制
	INTERNALERROR_TURINGREQUESTLIMITEXCEEDED = "InternalError.TuringRequestLimitExceeded"

	// 服务调用失败
	INTERNALERROR_TURINGSERVICEFAILED = "InternalError.TuringServiceFailed"

	// 子服务超时
	INTERNALERROR_TURINGSUBSERVICETIMEOUT = "InternalError.TuringSubServiceTimeout"

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

	// 服务参数错误
	INVALIDPARAMETERVALUE_INVALIDTURINGTOKEN = "InvalidParameterValue.InvalidTuringToken"

	// 租户不存在。
	INVALIDPARAMETERVALUE_TENANTNOTEXIST = "InvalidParameterValue.TenantNotExist"

	// 缺少参数错误。
	MISSINGPARAMETER = "MissingParameter"

	// 缺少参数错误，必传参数未填。
	MISSINGPARAMETER_FIELDMISSED = "MissingParameter.FieldMissed"

	// 未授权操作。
	UNAUTHORIZEDOPERATION = "UnauthorizedOperation"

	// 渠道号不匹配。
	UNAUTHORIZEDOPERATION_CHANNELIDMISMATCH = "UnauthorizedOperation.ChannelIdMismatch"

	// 签名检验失败
	UNAUTHORIZEDOPERATION_INVALIDTURINGSIGNATURE = "UnauthorizedOperation.InvalidTuringSignature"

	// 服务未找到
	UNAUTHORIZEDOPERATION_RESOURCENOTFOUND = "UnauthorizedOperation.ResourceNotFound"

	// 重放请求
	UNAUTHORIZEDOPERATION_TURINGREPLAYREQUEST = "UnauthorizedOperation.TuringReplayRequest"

	// 服务访问未授权
	UNAUTHORIZEDOPERATION_TURINGSERVICEUNAUTHORIZED = "UnauthorizedOperation.TuringServiceUnauthorized"

	// appid未配置
	UNAUTHORIZEDOPERATION_UNKNOWNTURINGAPPID = "UnauthorizedOperation.UnknownTuringAppid"
)
