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

package v20190408

const (
	// 此产品的特有错误码

	// 请求头部的 Authorization 不符合腾讯云标准。
	AUTHFAILURE_INVALIDAUTHORIZATION = "AuthFailure.InvalidAuthorization"

	// 操作失败。
	FAILEDOPERATION = "FailedOperation"

	// 余额不足，开通失败，请充值后再开通。
	FAILEDOPERATION_BALANCEINSUFFICIENT = "FailedOperation.BalanceInsufficient"

	// 非法文本输入导致返回异常
	FAILEDOPERATION_ILLEGALTEXTERROR = "FailedOperation.IllegalTextError"

	// 暂无春联生成，请更换关键词重试。
	FAILEDOPERATION_NOCOUPLETS = "FailedOperation.NoCouplets"

	// 暂无诗词生成，请更换关键词重试。
	FAILEDOPERATION_NOPOETRY = "FailedOperation.NoPoetry"

	// 未查询到结果。
	FAILEDOPERATION_NOTFOUNDDATA = "FailedOperation.NotFoundData"

	// 后端服务超时。
	FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"

	// RPC请求失败，一般为算法微服务故障。
	FAILEDOPERATION_RPCFAIL = "FailedOperation.RpcFail"

	// 文本向量化失败
	FAILEDOPERATION_TEXTEMBEDDINGFAILED = "FailedOperation.TextEmbeddingFailed"

	// 内部错误。
	FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"

	// 未知错误。
	FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"

	// 查找不到词语
	FAILEDOPERATION_WORDNOTFOUND = "FailedOperation.WordNotFound"

	// 内部错误。
	INTERNALERROR = "InternalError"

	// 资源请求错误
	INTERNALERROR_RESOURCEREQUESTERROR = "InternalError.ResourceRequestError"

	// 内部服务调用错误。
	INTERNALERROR_SERVICECALLERROR = "InternalError.ServiceCallError"

	// 内部服务调用失败。
	INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"

	// 参数错误。
	INVALIDPARAMETER = "InvalidParameter"

	// 参数空值错误
	INVALIDPARAMETERVALUE_EMPTYVALUEERROR = "InvalidParameterValue.EmptyValueError"

	// Genre非法，请参考Genre参数说明。
	INVALIDPARAMETERVALUE_GENRE = "InvalidParameterValue.Genre"

	// PoetryType非法，请参考PoetryType参数说明。
	INVALIDPARAMETERVALUE_POETRYTYPE = "InvalidParameterValue.PoetryType"

	// Text输入含有敏感信息。
	INVALIDPARAMETERVALUE_SENSITIVETEXT = "InvalidParameterValue.SensitiveText"

	// TargetType非法，请参考TargetType参数说明。
	INVALIDPARAMETERVALUE_TARGETTYPE = "InvalidParameterValue.TargetType"

	// Text非法，请参考Text参数说明。
	INVALIDPARAMETERVALUE_TEXT = "InvalidParameterValue.Text"

	// 文本编码错误，不符合utf-8
	INVALIDPARAMETERVALUE_TEXTENCODEERROR = "InvalidParameterValue.TextEncodeError"

	// 文本输入格式错误
	INVALIDPARAMETERVALUE_TEXTFORMATERROR = "InvalidParameterValue.TextFormatError"

	// 输入文本超出数量限制
	INVALIDPARAMETERVALUE_TEXTNUMTOOMUCH = "InvalidParameterValue.TextNumTooMuch"

	// 输入文本超出长度限制
	INVALIDPARAMETERVALUE_TEXTTOOLONG = "InvalidParameterValue.TextTooLong"

	// 参数取值范围错误
	INVALIDPARAMETERVALUE_VALUERANGEERROR = "InvalidParameterValue.ValueRangeError"

	// 资源用量达到上限。
	LIMITEXCEEDED_RESOURCEREACHEDLIMIT = "LimitExceeded.ResourceReachedLimit"

	// 请求的次数超过了频率限制。
	REQUESTLIMITEXCEEDED_UINLIMITEXCEEDED = "RequestLimitExceeded.UinLimitExceeded"

	// 名称已存在
	RESOURCEINUSE_NAMEEXISTS = "ResourceInUse.NameExists"

	// 资源正在操作中
	RESOURCEINUSE_RESOURCEOPERATING = "ResourceInUse.ResourceOperating"

	// 额度用尽，请充值后重试
	RESOURCEINSUFFICIENT_QUOTARUNOUT = "ResourceInsufficient.QuotaRunOut"

	// 资源不存在。
	RESOURCENOTFOUND = "ResourceNotFound"

	// 数据资源不存在
	RESOURCENOTFOUND_DATANOTFOUND = "ResourceNotFound.DataNotFound"

	// 文件资源不存在
	RESOURCENOTFOUND_FILENOTFOUND = "ResourceNotFound.FileNotFound"

	// 资源不可用。
	RESOURCEUNAVAILABLE = "ResourceUnavailable"

	// 文件资源不可用
	RESOURCEUNAVAILABLE_FILEUNAVAILABLE = "ResourceUnavailable.FileUnavailable"

	// 帐号已被冻结。
	RESOURCEUNAVAILABLE_FREEZE = "ResourceUnavailable.Freeze"

	// 账号已欠费。
	RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"

	// 服务正在开通中，请稍等。
	RESOURCEUNAVAILABLE_ISOPENING = "ResourceUnavailable.IsOpening"

	// 计费状态未知，请确认是否已在控制台开通服务。
	RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"

	// 资源已被回收。
	RESOURCEUNAVAILABLE_RECOVER = "ResourceUnavailable.Recover"

	// 您的账号尚未开通NLP服务，请登录腾讯云NLP控制台进行服务开通后再使用
	RESOURCEUNAVAILABLE_SERVICENOTOPENEDERROR = "ResourceUnavailable.ServiceNotOpenedError"

	// 帐号已停服。
	RESOURCEUNAVAILABLE_STOPUSING = "ResourceUnavailable.StopUsing"

	// 实名认证失败
	UNAUTHORIZEDOPERATION_AUTHENTICATEFAILED = "UnauthorizedOperation.AuthenticateFailed"
)
