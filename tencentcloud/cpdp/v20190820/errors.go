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

package v20190820

const (
	// 此产品的特有错误码

	// 未进行实名认证。
	AUTHFAILURE_ACCOUNT = "AuthFailure.Account"

	// 权限限额，禁止操作。
	AUTHFAILURE_FORBIDDEN = "AuthFailure.Forbidden"

	// 聚鑫签名信息不匹配。
	AUTHFAILURE_MIDAS = "AuthFailure.Midas"

	// 未找到密钥。
	AUTHFAILURE_SECRETKEYNOTFOUND = "AuthFailure.SecretKeyNotFound"

	// 账户尚未签约。
	AUTHFAILURE_SIGN = "AuthFailure.Sign"

	// 验证失败。
	AUTHFAILURE_VERIFYERROR = "AuthFailure.VerifyError"

	// Token校验失败。
	AUTHFAILURE_VERIFYTOKENFAILURE = "AuthFailure.VerifyTokenFailure"

	// 操作失败。
	FAILEDOPERATION = "FailedOperation"

	// 商户状态异常。
	FAILEDOPERATION_ABNORMALMERCHANTSTATE = "FailedOperation.AbnormalMerchantState"

	// 订单状态异常。
	FAILEDOPERATION_ABNORMALORDERSTATE = "FailedOperation.AbnormalOrderState"

	// 账户未绑定。
	FAILEDOPERATION_ACCTNOTBIND = "FailedOperation.AcctNotBind"

	// 账户不存在。
	FAILEDOPERATION_ACCTNOTEXIST = "FailedOperation.AcctNotExist"

	// 接口调用出错。
	FAILEDOPERATION_ACTION = "FailedOperation.Action"

	// 配置参数action无效。
	FAILEDOPERATION_ACTIONINVALID = "FailedOperation.ActionInvalid"

	// 商户创建失败。
	FAILEDOPERATION_ADDMERCHANTFAILED = "FailedOperation.AddMerchantFailed"

	// 订单已经存在。
	FAILEDOPERATION_ALREADYEXISTS = "FailedOperation.AlreadyExists"

	// 该业务代码对应的配置没有发布到相应的环境。
	FAILEDOPERATION_APPDENY = "FailedOperation.AppDeny"

	// 商户号和appid没有绑定关系。
	FAILEDOPERATION_APPIDMCHIDNOTMATCH = "FailedOperation.AppidMchidNotMatch"

	// 请求下游服务错误。
	FAILEDOPERATION_BACKCALLERROR = "FailedOperation.BackCallError"

	// 云鉴内部调用失败。
	FAILEDOPERATION_BACKENDERROR = "FailedOperation.BackendError"

	// 账户余额不足。
	FAILEDOPERATION_BALANCEINSUFFICIENT = "FailedOperation.BalanceInsufficient"

	// 调用银行接口失败，可能原因： 银行接口超时或为获取到返回值 银行接口返回非法。
	FAILEDOPERATION_BANKFAILED = "FailedOperation.BankFailed"

	// 余额不足。
	FAILEDOPERATION_BANLANCENOTENOUGHERROR = "FailedOperation.BanlanceNotEnoughError"

	// 内部请求渠道网关错误。
	FAILEDOPERATION_CALLCHANNELGATEWAYERROR = "FailedOperation.CallChannelGatewayError"

	// 支付渠道错误，请确认配置信息是 否提交至聚鑫。
	FAILEDOPERATION_CHANNELDENY = "FailedOperation.ChannelDeny"

	// 支付渠道错误。
	FAILEDOPERATION_CHANNELERROR = "FailedOperation.ChannelError"

	// 渠道方退款失败。
	FAILEDOPERATION_CHANNELREFUNDFAILED = "FailedOperation.ChannelRefundFailed"

	// 渠道退款频率受限。
	FAILEDOPERATION_CHANNELREFUNDFREQUENCYLIMITED = "FailedOperation.ChannelRefundFrequencyLimited"

	// 终止合约失败。
	FAILEDOPERATION_CLOSECONTRACTDBFAILED = "FailedOperation.CloseContractDbFailed"

	// 终止合约模式非法。
	FAILEDOPERATION_CLOSECONTRACTMODEINVALID = "FailedOperation.CloseContractModeInvalid"

	// 下载服务器配置错误。
	FAILEDOPERATION_CONFIGERROR = "FailedOperation.ConfigError"

	// 签约状态不正确。
	FAILEDOPERATION_CONTRACTSTATUSERROR = "FailedOperation.ContractStatusError"

	// 录入代理商出错。
	FAILEDOPERATION_CREATEAGENT = "FailedOperation.CreateAgent"

	// 消费订单发起失败。
	FAILEDOPERATION_CREATEORDERERROR = "FailedOperation.CreateOrderError"

	// 消费订单发起状态未知，请调用查询接口进行查询。
	FAILEDOPERATION_CREATEORDERUNKNOWN = "FailedOperation.CreateOrderUnknown"

	// 下载服务器数据库配置错误。
	FAILEDOPERATION_DBCONFIGERROR = "FailedOperation.DBConfigError"

	// 连接数据库失败。
	FAILEDOPERATION_DBCLIENTCONNECTFAILED = "FailedOperation.DbClientConnectFailed"

	// 数据库插入数据失败。
	FAILEDOPERATION_DBCLIENTINSERTTFAILED = "FailedOperation.DbClientInserttFailed"

	// 查询数据失败。
	FAILEDOPERATION_DBCLIENTQUERYFAILED = "FailedOperation.DbClientQueryFailed"

	// 数据更新失败。
	FAILEDOPERATION_DBCLIENTUPDATEFAILED = "FailedOperation.DbClientUpdateFailed"

	// 对账单下载失败。
	FAILEDOPERATION_DOWNLOADBILLERROR = "FailedOperation.DownloadBillError"

	// 未查到第三方渠道合约数据。
	FAILEDOPERATION_EXTERNALCONTRACTINDEXNOTFOUND = "FailedOperation.ExternalContractIndexNotFound"

	// 未查到第三方渠道合约数据。
	FAILEDOPERATION_EXTERNALCONTRACTNOTFOUND = "FailedOperation.ExternalContractNotFound"

	// 外部签约状态无效。
	FAILEDOPERATION_EXTERNALCONTRACTSTATUSINVALID = "FailedOperation.ExternalContractStatusInvalid"

	// 第三方渠道商户合约配置查询错误。
	FAILEDOPERATION_EXTERNALMERCHANTCONTRACTINFOCONFIGNOFOUND = "FailedOperation.ExternalMerchantContractInfoConfigNoFound"

	// 第三方渠道商户配置信息查询错误。
	FAILEDOPERATION_EXTERNALMERCHANTINDEXCONFIGNOFOUND = "FailedOperation.ExternalMerchantIndexConfigNoFound"

	// 第三方渠道商户配置信息查询错误。
	FAILEDOPERATION_EXTERNALMERCHANTINFOCONFIGNOFOUND = "FailedOperation.ExternalMerchantInfoConfigNoFound"

	// 对账文件不存在，尚未生成。
	FAILEDOPERATION_FILENOTEXIST = "FailedOperation.FileNotExist"

	// 频率超限。
	FAILEDOPERATION_FREQUENCYLIMITED = "FailedOperation.FrequencyLimited"

	// 获取日结报表列表出错。
	FAILEDOPERATION_GETLIVEDAILYSUMMARY = "FailedOperation.GetLiveDailySummary"

	// 内部http方式请求下游服务错误。
	FAILEDOPERATION_HTTPDOREQUESTERROR = "FailedOperation.HttpDoRequestError"

	// 内部服务超时。
	FAILEDOPERATION_INTERNALSERVICETIMEOUT = "FailedOperation.InternalServiceTimeout"

	// 参数错误。
	FAILEDOPERATION_INVALIDPARAMETER = "FailedOperation.InvalidParameter"

	// 退款金额无效。
	FAILEDOPERATION_INVALIDREFUNDAMT = "FailedOperation.InvalidRefundAmt"

	// 请求参数符合参数格式，但不符合业务规则。
	FAILEDOPERATION_INVALIDREQUEST = "FailedOperation.InvalidRequest"

	// 发票已提交。
	FAILEDOPERATION_INVOICEEXIST = "FailedOperation.InvoiceExist"

	// 文件为空，未生成或者当天无数据。
	FAILEDOPERATION_ISEMPTY = "FailedOperation.IsEmpty"

	// 序列化出错。
	FAILEDOPERATION_MARSHALERROR = "FailedOperation.MarshalError"

	// 商户余额不足。
	FAILEDOPERATION_MERCHANTBALANCENOTENOUGH = "FailedOperation.MerchantBalanceNotEnough"

	// 商户信息验证失败。
	FAILEDOPERATION_MERCHANTCHECKFAILED = "FailedOperation.MerchantCheckFailed"

	// 商户创建失败。
	FAILEDOPERATION_MERCHANTCREATEFAILED = "FailedOperation.MerchantCreateFailed"

	// 商户已存在。
	FAILEDOPERATION_MERCHANTEXIST = "FailedOperation.MerchantExist"

	// 查无此商户。
	FAILEDOPERATION_MERCHANTNOTEXIST = "FailedOperation.MerchantNotExist"

	// 商户不存在。
	FAILEDOPERATION_MERCHANTNOTEXISTS = "FailedOperation.MerchantNotExists"

	// 商户权限错误。
	FAILEDOPERATION_MERCHANTPERMISSIONERROR = "FailedOperation.MerchantPermissionError"

	// 聚鑫内部系统错误。
	FAILEDOPERATION_MIDASINTERNALERROR = "FailedOperation.MidasInternalError"

	// 聚鑫请求无效。
	FAILEDOPERATION_MIDASINVALIDREQUEST = "FailedOperation.MidasInvalidRequest"

	// 聚鑫需要重试。
	FAILEDOPERATION_MIDASNEEDRETRY = "FailedOperation.MidasNeedRetry"

	// 聚鑫当前企业对该产品能力有未完成的申请单。
	FAILEDOPERATION_MIDASREGISTERUNFINISHED = "FailedOperation.MidasRegisterUnfinished"

	// 聚鑫订单已经提交。
	FAILEDOPERATION_MIDASREPEATORDER = "FailedOperation.MidasRepeatOrder"

	// 通用风控系统错误，被风控拦截。
	FAILEDOPERATION_MIDASRISK = "FailedOperation.MidasRisk"

	// 聚鑫支付状态不匹配。
	FAILEDOPERATION_MIDASSTATUSNOTMATCH = "FailedOperation.MidasStatusNotMatch"

	// 聚鑫不支持该操作。
	FAILEDOPERATION_MIDASUNSUPPORTEDACTION = "FailedOperation.MidasUnsupportedAction"

	// 参数缺失。
	FAILEDOPERATION_MISSINGPARAMETER = "FailedOperation.MissingParameter"

	// 商户修改失败。
	FAILEDOPERATION_MODIFYMERCHANTFAILED = "FailedOperation.ModifyMerchantFailed"

	// 记录不存在。
	FAILEDOPERATION_MOUNTNOTFOUND = "FailedOperation.MountNotFound"

	// 商户信息不合法。
	FAILEDOPERATION_NOAUTH = "FailedOperation.NoAuth"

	// 订单不存在。
	FAILEDOPERATION_NORECORD = "FailedOperation.NoRecord"

	// 资金不足。
	FAILEDOPERATION_NOTENOUGH = "FailedOperation.NotEnough"

	// 记录不存在。
	FAILEDOPERATION_NOTFOUND = "FailedOperation.NotFound"

	// 业务的签约通知url解析错误。
	FAILEDOPERATION_NOTIFYURLPARSEERROR = "FailedOperation.NotifyUrlParseError"

	// 订单已经支付。
	FAILEDOPERATION_OCCOMPLETEDORDER = "FailedOperation.OcCompletedOrder"

	// 订单号重复，但是两次请求参数不 一致。
	FAILEDOPERATION_OCREPEATORDER = "FailedOperation.OcRepeatOrder"

	// 退款主单被锁，请待前退款单完成后再发起退款。
	FAILEDOPERATION_ORDERLOCKED = "FailedOperation.OrderLocked"

	// 订单状态不可用。
	FAILEDOPERATION_ORDERNOTACTIVATED = "FailedOperation.OrderNotActivated"

	// 消费订单退款失败。
	FAILEDOPERATION_ORDERREFUNDERROR = "FailedOperation.OrderRefundError"

	// 平安银行返回错误。
	FAILEDOPERATION_PABANKERROR = "FailedOperation.PABankError"

	// 微信返回参数错误。
	FAILEDOPERATION_PARAMERROR = "FailedOperation.ParamError"

	// 请求参数错误。
	FAILEDOPERATION_PARAMETERERROR = "FailedOperation.ParameterError"

	// AppId、SubAppId不匹配。
	FAILEDOPERATION_PARENTAPPIDERROR = "FailedOperation.ParentAppIdError"

	// 聚鑫内部系统调用失败。
	FAILEDOPERATION_PORTALERROR = "FailedOperation.PortalError"

	// 查询代理商结算单链接出错。
	FAILEDOPERATION_QUERYAGENTSTATEMENTS = "FailedOperation.QueryAgentStatements"

	// 查询签约关系不存在。
	FAILEDOPERATION_QUERYCONTRACTNULL = "FailedOperation.QueryContractNull"

	// 查询渠道返回错误。
	FAILEDOPERATION_QUERYMCHANNELERROR = "FailedOperation.QueryMchannelError"

	// 查询模式错误。
	FAILEDOPERATION_QUERYMODEERROR = "FailedOperation.QueryModeError"

	// 消费订单查询失败。
	FAILEDOPERATION_QUERYORDERERROR = "FailedOperation.QueryOrderError"

	// 渠道查询结果为空，请传入正确参数。
	FAILEDOPERATION_QUERYRESULTNULL = "FailedOperation.QueryResultNull"

	// 超出商户单日转账额度。
	FAILEDOPERATION_QUOTAEXCEED = "FailedOperation.QuotaExceed"

	// 退款信息重复。
	FAILEDOPERATION_REFUNDINFODUPLICATE = "FailedOperation.RefundInfoDuplicate"

	// 不可重试退款。
	FAILEDOPERATION_REFUNDNOTRETRIEABLE = "FailedOperation.RefundNotRetrieable"

	// 退款处理中。
	FAILEDOPERATION_REFUNDPROCESSING = "FailedOperation.RefundProcessIng"

	// 退款交易已结束。
	FAILEDOPERATION_REFUNDTRANSACTIONCLOSED = "FailedOperation.RefundTransactionClosed"

	// 退款交易已完成，请勿重复操作。
	FAILEDOPERATION_REFUNDTRANSACTIONFINISHED = "FailedOperation.RefundTransactionFinished"

	// 云鉴内部调用失败。
	FAILEDOPERATION_SDKERROR = "FailedOperation.SDKError"

	// 接口执行失败。
	FAILEDOPERATION_SERVICEERROR = "FailedOperation.ServiceError"

	// 签名错误。
	FAILEDOPERATION_SIGNERROR = "FailedOperation.SignError"

	// 调用后端同步服务失败。
	FAILEDOPERATION_SYNCMCHANNELERROR = "FailedOperation.SyncMchannelError"

	// 系统内部错误。
	FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"

	// 调用后端解约服务失败。
	FAILEDOPERATION_TERMINATEMCHANNELERROR = "FailedOperation.TerminateMchannelError"

	// 未知错误。
	FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"

	// 反序列化出错。
	FAILEDOPERATION_UNMARSHALERROR = "FailedOperation.UnmarshalError"

	// 冲正签约状态失败。
	FAILEDOPERATION_UPDATECONTRACTSTATUSFAILED = "FailedOperation.UpdateContractStatusFailed"

	// 上传代理商完税列表出错。
	FAILEDOPERATION_UPLOADTAXLIST = "FailedOperation.UploadTaxList"

	// 上传代理商完税证明出错。
	FAILEDOPERATION_UPLOADTAXPAYMENT = "FailedOperation.UploadTaxPayment"

	// 调用渠道错误，请确认渠道如微信 wxappid是否配置正确。
	FAILEDOPERATION_WECHATERROR = "FailedOperation.WechatError"

	// 微信支付证书未配置。
	FAILEDOPERATION_WXCRTNOTSET = "FailedOperation.WxCrtNotSet"

	// xml格式错误。
	FAILEDOPERATION_XMLFAIL = "FailedOperation.XmlFail"

	// 系统错误。
	INTERNALERROR_ = "InternalError."

	// 内部连接后端错误。
	INTERNALERROR_BACKENDCONNECTIONERROR = "InternalError.BackendConnectionError"

	// 后端服务访问错误。
	INTERNALERROR_BACKENDERROR = "InternalError.BackendError"

	// 服务内部错误。
	INTERNALERROR_BACKENDINTERNALERROR = "InternalError.BackendInternalError"

	// 内部调用后端网络错误。
	INTERNALERROR_BACKENDNETWORKERROR = "InternalError.BackendNetworkError"

	// 内部调用路由错误。
	INTERNALERROR_BACKENDROUTERERROR = "InternalError.BackendRouterError"

	// 内部调用后端超时。
	INTERNALERROR_BACKENDTIMEOUT = "InternalError.BackendTimeOut"

	// 数据库访问错误。
	INTERNALERROR_DBACCESSERROR = "InternalError.DBAccessError"

	// 数据删除失败。
	INTERNALERROR_DELETEDBERROR = "InternalError.DeleteDBError"

	// 数据重复插入错误。
	INTERNALERROR_DUPLICATEKEYERROR = "InternalError.DuplicateKeyError"

	// 资金汇总账号不一致错误。
	INTERNALERROR_FUNDSUMMARYACCTNOINCONSISTENTERROR = "InternalError.FundSummaryAcctNoInconsistentError"

	// 发票已开具。
	INTERNALERROR_INVOICEEXIST = "InternalError.InvoiceExist"

	// 聚鑫内部系统未知错误。
	INTERNALERROR_MIDAS = "InternalError.Midas"

	// 参数错误。
	INTERNALERROR_PARAMETERERROR = "InternalError.ParameterError"

	// 沙箱环境访问错误。
	INTERNALERROR_SANDBOXACCESSERROR = "InternalError.SandBoxAccessError"

	// 数据保存失败。
	INTERNALERROR_SAVEDBERROR = "InternalError.SaveDBError"

	// 生成签名失败。
	INTERNALERROR_SIGGENERROR = "InternalError.SigGenError"

	// 子账户不存在错误。
	INTERNALERROR_SUBACCOUNTNOTFOUNDERROR = "InternalError.SubAccountNotFoundError"

	// 内部错误。
	INTERNALERROR_UNKNOWN = "InternalError.Unknown"

	// 未知错误。
	INTERNALERROR_UNKOWNERROR = "InternalError.UnkownError"

	// 参数错误。
	INVALIDPARAMETER = "InvalidParameter"

	// 传入参数有误，请检查传入参数是否正确。
	INVALIDPARAMETER_BACKENDCGIERROR = "InvalidParameter.BackendCgiError"

	// 缺少必填参数。
	INVALIDPARAMETER_LACKPARAMETER = "InvalidParameter.LackParameter"

	// 聚鑫参数无效。
	INVALIDPARAMETER_MIDAS = "InvalidParameter.Midas"

	// 聚鑫环境参数非法。
	INVALIDPARAMETER_MIDASENVIRONMENT = "InvalidParameter.MidasEnvironment"

	// 聚鑫第三方应用无效。
	INVALIDPARAMETER_MIDASEXTERNALAPP = "InvalidParameter.MidasExternalApp"

	// 聚鑫文件格式错误。
	INVALIDPARAMETER_MIDASFILETYPE = "InvalidParameter.MidasFileType"

	// 聚鑫摘要错误。
	INVALIDPARAMETER_MIDASHASH = "InvalidParameter.MidasHash"

	// 聚鑫签约ID非法。
	INVALIDPARAMETER_MIDASSIGNID = "InvalidParameter.MidasSignId"

	// 参数编码失败。
	INVALIDPARAMETER_PARAMMARSHALFAILED = "InvalidParameter.ParamMarshalFailed"

	// 参数解析失败。
	INVALIDPARAMETER_PARAMUNMARSHALFAILED = "InvalidParameter.ParamUnmarshalFailed"

	// 无效参数。
	INVALIDPARAMETER_UNSUPPORTEDPARAMETER = "InvalidParameter.UnsupportedParameter"

	// 聚鑫文件过大。
	LIMITEXCEEDED_MIDASLARGEFILE = "LimitExceeded.MidasLargeFile"

	// 聚鑫不允许并发下单。
	LIMITEXCEEDED_MIDASORDER = "LimitExceeded.MidasOrder"

	// 聚鑫订单已取消。
	LIMITEXCEEDED_MIDASORDERCANCELED = "LimitExceeded.MidasOrderCanceled"

	// 聚鑫已关单。
	LIMITEXCEEDED_MIDASORDERCLOSED = "LimitExceeded.MidasOrderClosed"

	// 聚鑫订单已过期，包括超时未支付、超过退款期限等情况。
	LIMITEXCEEDED_MIDASORDEREXPIRED = "LimitExceeded.MidasOrderExpired"

	// 聚鑫处理失败。
	LIMITEXCEEDED_MIDASORDERFAILED = "LimitExceeded.MidasOrderFailed"

	// 聚鑫处理部分成功部分失败。
	LIMITEXCEEDED_MIDASORDERPARTIALSUCCESS = "LimitExceeded.MidasOrderPartialSuccess"

	// 聚鑫处理中。
	LIMITEXCEEDED_MIDASORDERPROCESSING = "LimitExceeded.MidasOrderProcessing"

	// 聚鑫处理成功，请勿再提交。
	LIMITEXCEEDED_MIDASORDERSUCCESS = "LimitExceeded.MidasOrderSuccess"

	// 缺少参数错误。
	MISSINGPARAMETER = "MissingParameter"

	// 参数缺失。
	MISSINGPARAMETER_ = "MissingParameter."

	// 参数缺少Action或Action不存在。
	MISSINGPARAMETER_ACTION = "MissingParameter.Action"

	// 参数缺少Appid。
	MISSINGPARAMETER_APPID = "MissingParameter.AppId"

	// 聚鑫接口请求频率限制。
	REQUESTLIMITEXCEEDED_MIDAS = "RequestLimitExceeded.Midas"

	// 聚鑫接口无效请求过多。
	REQUESTLIMITEXCEEDED_MIDASINVALIDREQUEST = "RequestLimitExceeded.MidasInvalidRequest"

	// 聚鑫流程进行中，不能重入。
	RESOURCEINUSE_MIDAS = "ResourceInUse.Midas"

	// 运行资源不足。
	RESOURCEINSUFFICIENT_THREADPOOLREJECT = "ResourceInsufficient.ThreadPoolReject"

	// 账户不匹配或不存在。
	RESOURCENOTFOUND_ACCOUNT = "ResourceNotFound.Account"

	// 批次信息不存在。
	RESOURCENOTFOUND_BATCHINFONOTFOUND = "ResourceNotFound.BatchInfoNotFound"

	// 发票信息不存在。
	RESOURCENOTFOUND_INVOICENOTFOUND = "ResourceNotFound.InvoiceNotFound"

	// 密钥不匹配或不存在。
	RESOURCENOTFOUND_KEY = "ResourceNotFound.Key"

	// 商户信息不存在。
	RESOURCENOTFOUND_MERCHANTINFONOTFOUND = "ResourceNotFound.MerchantInfoNotFound"

	// 聚鑫第三方应用未查找到。
	RESOURCENOTFOUND_MIDASEXTERNALAPP = "ResourceNotFound.MidasExternalApp"

	// 聚鑫第三方渠道找不到订单，ORDER_NOT_FOUND表示支付中台找不到订单。
	RESOURCENOTFOUND_MIDASEXTERNALORDER = "ResourceNotFound.MidasExternalOrder"

	// 聚鑫订单没有查到记录。
	RESOURCENOTFOUND_MIDASORDER = "ResourceNotFound.MidasOrder"

	// 聚鑫签约关系不存在。
	RESOURCENOTFOUND_MIDASSIGN = "ResourceNotFound.MidasSign"

	// 平台信息不存在。
	RESOURCENOTFOUND_PLATFORMINFONOTFOUND = "ResourceNotFound.PlatformInfoNotFound"

	// 聚鑫可用余额不足。
	RESOURCEUNAVAILABLE_MIDASBALANCE = "ResourceUnavailable.MidasBalance"

	// 聚鑫单日限额。
	RESOURCEUNAVAILABLE_MIDASDAY = "ResourceUnavailable.MidasDay"

	// 聚鑫订单剩余冻结金额不足，因此会完结失败。
	RESOURCEUNAVAILABLE_MIDASFROZENAMOUNT = "ResourceUnavailable.MidasFrozenAmount"

	// 聚鑫商户可用余额不足。
	RESOURCEUNAVAILABLE_MIDASMERCHANTBALANCE = "ResourceUnavailable.MidasMerchantBalance"

	// 聚鑫单笔限额。
	RESOURCEUNAVAILABLE_MIDASORDER = "ResourceUnavailable.MidasOrder"

	// 聚鑫用户可用余额不足。
	RESOURCEUNAVAILABLE_MIDASUSERBALANCE = "ResourceUnavailable.MidasUserBalance"

	// 聚鑫钱包支付超限。
	RESOURCEUNAVAILABLE_MIDASWALLET = "ResourceUnavailable.MidasWallet"

	// 聚鑫系统未授权。
	UNAUTHORIZEDOPERATION_MIDAS = "UnauthorizedOperation.Midas"
)
