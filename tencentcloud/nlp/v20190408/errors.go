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

	// 操作失败。
	FAILEDOPERATION = "FailedOperation"

	// 非法文本输入导致返回异常
	FAILEDOPERATION_ILLEGALTEXTERROR = "FailedOperation.IllegalTextError"

	// 文本向量化失败
	FAILEDOPERATION_TEXTEMBEDDINGFAILED = "FailedOperation.TextEmbeddingFailed"

	// 查找不到词语
	FAILEDOPERATION_WORDNOTFOUND = "FailedOperation.WordNotFound"

	// 内部错误。
	INTERNALERROR = "InternalError"

	// 外部服务调用失败
	INTERNALERROR_EXTERNALSERVICEERROR = "InternalError.ExternalServiceError"

	// 图数据库引擎内部错误	
	INTERNALERROR_GRAPHENGINEERROR = "InternalError.GraphEngineError"

	// 图查询引擎计算超时错误	
	INTERNALERROR_GREMLINTIMEOUTERROR = "InternalError.GremlinTimeoutError"

	// 请求数据解析失败
	INTERNALERROR_REQUESTENCODEERROR = "InternalError.RequestEncodeError"

	// 读取请求数据内容失败
	INTERNALERROR_REQUESTPARSEERROR = "InternalError.RequestParseError"

	// 资源请求错误
	INTERNALERROR_RESOURCEREQUESTERROR = "InternalError.ResourceRequestError"

	// 系统响应解析失败
	INTERNALERROR_RESPONSEDECODEERROR = "InternalError.ResponseDecodeError"

	// 服务调用错误
	INTERNALERROR_SERVICECALLERROR = "InternalError.ServiceCallError"

	// 参数错误。
	INVALIDPARAMETER = "InvalidParameter"

	// 接口不存在
	INVALIDPARAMETER_INVALIDACTION = "InvalidParameter.InvalidAction"

	// 不可用的接口操作错误
	INVALIDPARAMETER_INVALIDACTIONERROR = "InvalidParameter.InvalidActionError"

	// 参数取值错误。
	INVALIDPARAMETERVALUE = "InvalidParameterValue"

	// 参数空值错误
	INVALIDPARAMETERVALUE_EMPTYVALUEERROR = "InvalidParameterValue.EmptyValueError"

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

	// 资源用量达到上限
	LIMITEXCEEDED_RESOURCEREACHEDLIMIT = "LimitExceeded.ResourceReachedLimit"

	// 缺少参数错误。
	MISSINGPARAMETER = "MissingParameter"

	// 资源被占用。
	RESOURCEINUSE = "ResourceInUse"

	// 名称已存在
	RESOURCEINUSE_NAMEEXISTS = "ResourceInUse.NameExists"

	// 资源正在操作中
	RESOURCEINUSE_RESOURCEOPERATING = "ResourceInUse.ResourceOperating"

	// 资源不足。
	RESOURCEINSUFFICIENT = "ResourceInsufficient"

	// 额度用尽，请充值后重试
	RESOURCEINSUFFICIENT_QUOTARUNOUT = "ResourceInsufficient.QuotaRunOut"

	// 数据资源不存在
	RESOURCENOTFOUND_DATANOTFOUND = "ResourceNotFound.DataNotFound"

	// 文件资源不存在
	RESOURCENOTFOUND_FILENOTFOUND = "ResourceNotFound.FileNotFound"

	// 资源不可用。
	RESOURCEUNAVAILABLE = "ResourceUnavailable"

	// 文件资源不可用
	RESOURCEUNAVAILABLE_FILEUNAVAILABLE = "ResourceUnavailable.FileUnavailable"

	// 您的账号尚未开通NLP服务，请登录腾讯云NLP控制台进行服务开通后再使用
	RESOURCEUNAVAILABLE_SERVICENOTOPENEDERROR = "ResourceUnavailable.ServiceNotOpenedError"

	// 未授权操作。
	UNAUTHORIZEDOPERATION = "UnauthorizedOperation"

	// 实名认证失败
	UNAUTHORIZEDOPERATION_AUTHENTICATEFAILED = "UnauthorizedOperation.AuthenticateFailed"
)
