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

package v20180125

const (
	// 此产品的特有错误码

	// CAM签名/鉴权错误。
	AUTHFAILURE = "AuthFailure"

	// 当前实例未购买api安全套餐
	AUTHFAILURE_ERRCODENOPURCHASED = "AuthFailure.ErrCodeNoPurchased"

	// DryRun 操作，代表请求将会是成功的，只是多传了 DryRun 参数。
	DRYRUNOPERATION = "DryRunOperation"

	// 操作失败。
	FAILEDOPERATION = "FailedOperation"

	// 调用CLS日志服务API失败
	FAILEDOPERATION_CLSDBOPERATIONFAILED = "FailedOperation.CLSDBOperationFailed"

	// CLS内部错误。
	FAILEDOPERATION_CLSINTERNALERROR = "FailedOperation.CLSInternalError"

	// 操作CH数据库失败
	FAILEDOPERATION_CLICKHOUSEOPERATIONFAILED = "FailedOperation.ClickHouseOperationFailed"

	// 操作Mysql数据库失败
	FAILEDOPERATION_MYSQLDBOPERATIONFAILED = "FailedOperation.MysqlDBOperationFailed"

	// 操作Redis数据库失败
	FAILEDOPERATION_REDISOPERATIONFAILED = "FailedOperation.RedisOperationFailed"

	// 删除的Session正在被启用
	FAILEDOPERATION_SESSIONINUSED = "FailedOperation.SessionInUsed"

	// 黑白名单添加数超过上限
	FAILEDOPERATION_THENUMBEROFADDEDBLACKANDWHITELISTEXCEEDSTHEUPPERLIMIT = "FailedOperation.TheNumberOfAddedBlackAndWhiteListExceedsTheUpperLimit"

	// 一次性删除数量达到上限
	FAILEDOPERATION_THENUMBEROFONETIMEDELETIONSREACHEDTHEUPPERLIMIT = "FailedOperation.TheNumberOfOneTimeDeletionsReachedTheUpperLimit"

	// 内部错误。
	INTERNALERROR = "InternalError"

	// 异步调用失败
	INTERNALERROR_ASYNCHRONOUSCALLFAILED = "InternalError.AsynchronousCallFailed"

	// DBErr
	INTERNALERROR_DBERR = "InternalError.DBErr"

	// 存在内部错误，请联系我们
	INTERNALERROR_UNKNOWNERR = "InternalError.UnknownErr"

	// 参数错误。
	INVALIDPARAMETER = "InvalidParameter"

	// 证书信息参数错误
	INVALIDPARAMETER_CERTIFICATIONPARAMETERERR = "InvalidParameter.CertificationParameterErr"

	// 域名数量超出限制错误
	INVALIDPARAMETER_DOMAINEXCEEDSLIMITERR = "InvalidParameter.DomainExceedsLimitErr"

	// 域名未备案
	INVALIDPARAMETER_DOMAINNOTRECORD = "InvalidParameter.DomainNotRecord"

	// 证书内容非法。
	INVALIDPARAMETER_INVALIDCERTIFICATE = "InvalidParameter.InvalidCertificate"

	// 逻辑错误：SQL检索语句中的逻辑错误也可能导致错误。例如，使用错误的运算符、使用错误的条件等
	INVALIDPARAMETER_LOGICERR = "InvalidParameter.LogicErr"

	// 端口信息参数错误
	INVALIDPARAMETER_PORTPARAMETERERR = "InvalidParameter.PortParameterErr"

	// 防护域名参数错误
	INVALIDPARAMETER_PROTECTIONDOMAINPARAMETERERR = "InvalidParameter.ProtectionDomainParameterErr"

	// 根据ID查询证书失败。
	INVALIDPARAMETER_QUERYCERTBYSSLIDFAILED = "InvalidParameter.QueryCertBySSLIDFailed"

	// 语法错误：逻辑表达式语法解析出错
	INVALIDPARAMETER_QUERYSTRINGSYNTAXERR = "InvalidParameter.QueryStringSyntaxErr"

	// 语法错误：SQL检索语句必须遵循特定的语法规则，如果语法错误，就会导致SQL语句无法执行。例如，缺少关键字、拼写错误、缺少分号等。
	INVALIDPARAMETER_SQLSYNTAXERR = "InvalidParameter.SQLSyntaxErr"

	// 当前实例版本不支持开启TLS自定义，请升级到高级版及以上
	INVALIDPARAMETER_SUPPORTTLSCONFFAILED = "InvalidParameter.SupportTLSConfFailed"

	// TLS或加密套件参数错误
	INVALIDPARAMETER_TLSPARAMETERERR = "InvalidParameter.TLSParameterErr"

	// 数据类型错误：SQL检索语句中的数据类型必须与数据库中的数据类型匹配，否则会导致错误。例如，将字符串与整数进行比较、将日期格式不正确等。
	INVALIDPARAMETER_TYPEMISMATCH = "InvalidParameter.TypeMismatch"

	// 越权参数错误
	INVALIDPARAMETER_UNAUTHORIZEDOPERATIONPARAMETERERR = "InvalidParameter.UnauthorizedOperationParameterErr"

	// UnknownAction
	INVALIDPARAMETER_UNKNOWNACTION = "InvalidParameter.UnknownAction"

	// 回源信息参数错误
	INVALIDPARAMETER_UPSTREAMPARAMETERERR = "InvalidParameter.UpstreamParameterErr"

	// XFF重置参数错误
	INVALIDPARAMETER_XFFRESETPARAMETERERR = "InvalidParameter.XFFResetParameterErr"

	// 参数取值错误。
	INVALIDPARAMETERVALUE = "InvalidParameterValue"

	// 超过配额限制。
	LIMITEXCEEDED = "LimitExceeded"

	// SpecificationErr
	LIMITEXCEEDED_SPECIFICATIONERR = "LimitExceeded.SpecificationErr"

	// 缺少参数错误。
	MISSINGPARAMETER = "MissingParameter"

	// 操作被拒绝。
	OPERATIONDENIED = "OperationDenied"

	// 请求的次数超过了频率限制。
	REQUESTLIMITEXCEEDED = "RequestLimitExceeded"

	// 资源被占用。
	RESOURCEINUSE = "ResourceInUse"

	// EmptyErr
	RESOURCEINUSE_EMPTYERR = "ResourceInUse.EmptyErr"

	// 资源不足。
	RESOURCEINSUFFICIENT = "ResourceInsufficient"

	// 资源不存在。
	RESOURCENOTFOUND = "ResourceNotFound"

	// 资源不可用。
	RESOURCEUNAVAILABLE = "ResourceUnavailable"

	// IPV6网络正在分配，请耐心等待
	RESOURCEUNAVAILABLE_DOMAINIPV6INCONFIGERR = "ResourceUnavailable.DomainIpv6InConfigErr"

	// 资源售罄。
	RESOURCESSOLDOUT = "ResourcesSoldOut"

	// 未授权操作。
	UNAUTHORIZEDOPERATION = "UnauthorizedOperation"

	// 未知参数错误。
	UNKNOWNPARAMETER = "UnknownParameter"

	// 操作不支持。
	UNSUPPORTEDOPERATION = "UnsupportedOperation"

	// InvalidRequest
	UNSUPPORTEDOPERATION_INVALIDREQUEST = "UnsupportedOperation.InvalidRequest"
)
