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

import (
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2019-08-20"

type Client struct {
    common.Client
}

// Deprecated
func NewClientWithSecretId(secretId, secretKey, region string) (client *Client, err error) {
    cpf := profile.NewClientProfile()
    client = &Client{}
    client.Init(region).WithSecretId(secretId, secretKey).WithProfile(cpf)
    return
}

func NewClient(credential common.CredentialIface, region string, clientProfile *profile.ClientProfile) (client *Client, err error) {
    client = &Client{}
    client.Init(region).
        WithCredential(credential).
        WithProfile(clientProfile)
    return
}


func NewAddContractRequest() (request *AddContractRequest) {
    request = &AddContractRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "AddContract")
    
    
    return
}

func NewAddContractResponse() (response *AddContractResponse) {
    response = &AddContractResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AddContract
// 云支付-添加合同接口
func (c *Client) AddContract(request *AddContractRequest) (response *AddContractResponse, err error) {
    if request == nil {
        request = NewAddContractRequest()
    }
    
    response = NewAddContractResponse()
    err = c.Send(request, response)
    return
}

func NewAddMerchantRequest() (request *AddMerchantRequest) {
    request = &AddMerchantRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "AddMerchant")
    
    
    return
}

func NewAddMerchantResponse() (response *AddMerchantResponse) {
    response = &AddMerchantResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AddMerchant
// 云支付-添加商户接口
func (c *Client) AddMerchant(request *AddMerchantRequest) (response *AddMerchantResponse, err error) {
    if request == nil {
        request = NewAddMerchantRequest()
    }
    
    response = NewAddMerchantResponse()
    err = c.Send(request, response)
    return
}

func NewAddShopRequest() (request *AddShopRequest) {
    request = &AddShopRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "AddShop")
    
    
    return
}

func NewAddShopResponse() (response *AddShopResponse) {
    response = &AddShopResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AddShop
// 云支付-添加门店接口
func (c *Client) AddShop(request *AddShopRequest) (response *AddShopResponse, err error) {
    if request == nil {
        request = NewAddShopRequest()
    }
    
    response = NewAddShopResponse()
    err = c.Send(request, response)
    return
}

func NewApplyApplicationMaterialRequest() (request *ApplyApplicationMaterialRequest) {
    request = &ApplyApplicationMaterialRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "ApplyApplicationMaterial")
    
    
    return
}

func NewApplyApplicationMaterialResponse() (response *ApplyApplicationMaterialResponse) {
    response = &ApplyApplicationMaterialResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ApplyApplicationMaterial
// 跨境-提交申报材料。申报材料的主体是付款人，需要提前调用【跨境-付款人申请】接口提交付款人信息且审核通过后调用。
//
// 可能返回的错误码:
//  INTERNALERROR_PARAMETERERROR = "InternalError.ParameterError"
//  INTERNALERROR_UNKOWNERROR = "InternalError.UnkownError"
//  RESOURCENOTFOUND_MERCHANTINFONOTFOUND = "ResourceNotFound.MerchantInfoNotFound"
func (c *Client) ApplyApplicationMaterial(request *ApplyApplicationMaterialRequest) (response *ApplyApplicationMaterialResponse, err error) {
    if request == nil {
        request = NewApplyApplicationMaterialRequest()
    }
    
    response = NewApplyApplicationMaterialResponse()
    err = c.Send(request, response)
    return
}

func NewApplyOutwardOrderRequest() (request *ApplyOutwardOrderRequest) {
    request = &ApplyOutwardOrderRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "ApplyOutwardOrder")
    
    
    return
}

func NewApplyOutwardOrderResponse() (response *ApplyOutwardOrderResponse) {
    response = &ApplyOutwardOrderResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ApplyOutwardOrder
// 跨境-汇出指令申请。通过该接口可将对接方账户中的人民币余额汇兑成外币，再汇出至指定银行账户。
//
// 可能返回的错误码:
//  INTERNALERROR_PARAMETERERROR = "InternalError.ParameterError"
//  INTERNALERROR_UNKOWNERROR = "InternalError.UnkownError"
//  RESOURCENOTFOUND_MERCHANTINFONOTFOUND = "ResourceNotFound.MerchantInfoNotFound"
func (c *Client) ApplyOutwardOrder(request *ApplyOutwardOrderRequest) (response *ApplyOutwardOrderResponse, err error) {
    if request == nil {
        request = NewApplyOutwardOrderRequest()
    }
    
    response = NewApplyOutwardOrderResponse()
    err = c.Send(request, response)
    return
}

func NewApplyPayerInfoRequest() (request *ApplyPayerInfoRequest) {
    request = &ApplyPayerInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "ApplyPayerInfo")
    
    
    return
}

func NewApplyPayerInfoResponse() (response *ApplyPayerInfoResponse) {
    response = &ApplyPayerInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ApplyPayerInfo
// 跨境-付款人申请。通过该接口提交付款人信息并进行 kyc 审核。
//
// 可能返回的错误码:
//  INTERNALERROR_PARAMETERERROR = "InternalError.ParameterError"
//  INTERNALERROR_UNKOWNERROR = "InternalError.UnkownError"
//  RESOURCENOTFOUND_MERCHANTINFONOTFOUND = "ResourceNotFound.MerchantInfoNotFound"
func (c *Client) ApplyPayerInfo(request *ApplyPayerInfoRequest) (response *ApplyPayerInfoResponse, err error) {
    if request == nil {
        request = NewApplyPayerInfoRequest()
    }
    
    response = NewApplyPayerInfoResponse()
    err = c.Send(request, response)
    return
}

func NewApplyReWithdrawalRequest() (request *ApplyReWithdrawalRequest) {
    request = &ApplyReWithdrawalRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "ApplyReWithdrawal")
    
    
    return
}

func NewApplyReWithdrawalResponse() (response *ApplyReWithdrawalResponse) {
    response = &ApplyReWithdrawalResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ApplyReWithdrawal
// 正常结算提现失败情况下，发起重新提现的请求接口
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_NORECORD = "FailedOperation.NoRecord"
//  FAILEDOPERATION_OCCOMPLETEDORDER = "FailedOperation.OcCompletedOrder"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ApplyReWithdrawal(request *ApplyReWithdrawalRequest) (response *ApplyReWithdrawalResponse, err error) {
    if request == nil {
        request = NewApplyReWithdrawalRequest()
    }
    
    response = NewApplyReWithdrawalResponse()
    err = c.Send(request, response)
    return
}

func NewApplyTradeRequest() (request *ApplyTradeRequest) {
    request = &ApplyTradeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "ApplyTrade")
    
    
    return
}

func NewApplyTradeResponse() (response *ApplyTradeResponse) {
    response = &ApplyTradeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ApplyTrade
// 跨境-提交贸易材料。通过提交贸易材料接口可为对接方累计贸易额度，在额度范围内可发起汇兑汇出交易。
//
// 可能返回的错误码:
//  INTERNALERROR_PARAMETERERROR = "InternalError.ParameterError"
//  INTERNALERROR_UNKOWNERROR = "InternalError.UnkownError"
//  RESOURCENOTFOUND_MERCHANTINFONOTFOUND = "ResourceNotFound.MerchantInfoNotFound"
func (c *Client) ApplyTrade(request *ApplyTradeRequest) (response *ApplyTradeResponse, err error) {
    if request == nil {
        request = NewApplyTradeRequest()
    }
    
    response = NewApplyTradeResponse()
    err = c.Send(request, response)
    return
}

func NewApplyWithdrawalRequest() (request *ApplyWithdrawalRequest) {
    request = &ApplyWithdrawalRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "ApplyWithdrawal")
    
    
    return
}

func NewApplyWithdrawalResponse() (response *ApplyWithdrawalResponse) {
    response = &ApplyWithdrawalResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ApplyWithdrawal
// 商户提现
//
// 可能返回的错误码:
//  FAILEDOPERATION_BANKFAILED = "FailedOperation.BankFailed"
//  FAILEDOPERATION_SERVICEERROR = "FailedOperation.ServiceError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER_ = "MissingParameter."
func (c *Client) ApplyWithdrawal(request *ApplyWithdrawalRequest) (response *ApplyWithdrawalResponse, err error) {
    if request == nil {
        request = NewApplyWithdrawalRequest()
    }
    
    response = NewApplyWithdrawalResponse()
    err = c.Send(request, response)
    return
}

func NewBindAccountRequest() (request *BindAccountRequest) {
    request = &BindAccountRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "BindAccount")
    
    
    return
}

func NewBindAccountResponse() (response *BindAccountResponse) {
    response = &BindAccountResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// BindAccount
// 灵云-绑定账号
//
// 可能返回的错误码:
//  FAILEDOPERATION_BANKFAILED = "FailedOperation.BankFailed"
//  FAILEDOPERATION_SERVICEERROR = "FailedOperation.ServiceError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER_ = "MissingParameter."
func (c *Client) BindAccount(request *BindAccountRequest) (response *BindAccountResponse, err error) {
    if request == nil {
        request = NewBindAccountRequest()
    }
    
    response = NewBindAccountResponse()
    err = c.Send(request, response)
    return
}

func NewBindAcctRequest() (request *BindAcctRequest) {
    request = &BindAcctRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "BindAcct")
    
    
    return
}

func NewBindAcctResponse() (response *BindAcctResponse) {
    response = &BindAcctResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// BindAcct
// 商户绑定提现银行卡，每个商户只能绑定一张提现银行卡
//
// 可能返回的错误码:
//  FAILEDOPERATION_BANKFAILED = "FailedOperation.BankFailed"
//  FAILEDOPERATION_SERVICEERROR = "FailedOperation.ServiceError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER_ = "MissingParameter."
func (c *Client) BindAcct(request *BindAcctRequest) (response *BindAcctResponse, err error) {
    if request == nil {
        request = NewBindAcctRequest()
    }
    
    response = NewBindAcctResponse()
    err = c.Send(request, response)
    return
}

func NewBindRelateAccReUnionPayRequest() (request *BindRelateAccReUnionPayRequest) {
    request = &BindRelateAccReUnionPayRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "BindRelateAccReUnionPay")
    
    
    return
}

func NewBindRelateAccReUnionPayResponse() (response *BindRelateAccReUnionPayResponse) {
    response = &BindRelateAccReUnionPayResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// BindRelateAccReUnionPay
// 会员绑定提现账户-回填银联鉴权短信码。用于会员填写动态验证码后，发往银行进行验证，验证成功则完成绑定。
//
// 可能返回的错误码:
//  FAILEDOPERATION_BACKENDERROR = "FailedOperation.BackendError"
//  INTERNALERROR_BACKENDERROR = "InternalError.BackendError"
//  INTERNALERROR_DBACCESSERROR = "InternalError.DBAccessError"
//  INTERNALERROR_FUNDSUMMARYACCTNOINCONSISTENTERROR = "InternalError.FundSummaryAcctNoInconsistentError"
//  INTERNALERROR_PARAMETERERROR = "InternalError.ParameterError"
//  INTERNALERROR_SAVEDBERROR = "InternalError.SaveDBError"
//  INTERNALERROR_SUBACCOUNTNOTFOUNDERROR = "InternalError.SubAccountNotFoundError"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  INTERNALERROR_UNKOWNERROR = "InternalError.UnkownError"
func (c *Client) BindRelateAccReUnionPay(request *BindRelateAccReUnionPayRequest) (response *BindRelateAccReUnionPayResponse, err error) {
    if request == nil {
        request = NewBindRelateAccReUnionPayRequest()
    }
    
    response = NewBindRelateAccReUnionPayResponse()
    err = c.Send(request, response)
    return
}

func NewBindRelateAcctSmallAmountRequest() (request *BindRelateAcctSmallAmountRequest) {
    request = &BindRelateAcctSmallAmountRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "BindRelateAcctSmallAmount")
    
    
    return
}

func NewBindRelateAcctSmallAmountResponse() (response *BindRelateAcctSmallAmountResponse) {
    response = &BindRelateAcctSmallAmountResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// BindRelateAcctSmallAmount
// 会员绑定提现账户-小额鉴权。会员申请绑定提现账户，绑定后从会员子账户中提现到绑定账户。
//
// 转账鉴权有两种形式：往账鉴权和来账鉴权。
//
// 往账鉴权：该接口发起成功后，银行会向提现账户转入小于等于0.5元的随机金额，并短信通知客户查看，客户查看后，需将收到的金额大小，在电商平台页面上回填，并通知银行。银行验证通过后，完成提现账户绑定。
//
// 来账鉴权：该接口发起成功后，银行以短信通知客户查看，客户查看后，需通过待绑定的账户往市场的监管账户转入短信上指定的金额。银行检索到该笔指定金额的来账是源自待绑定账户，则绑定成功。平安银行的账户，即BankType送1时，大小额行号和超级网银号都不用送。
//
// 可能返回的错误码:
//  FAILEDOPERATION_BACKENDERROR = "FailedOperation.BackendError"
//  INTERNALERROR_BACKENDERROR = "InternalError.BackendError"
//  INTERNALERROR_DBACCESSERROR = "InternalError.DBAccessError"
//  INTERNALERROR_FUNDSUMMARYACCTNOINCONSISTENTERROR = "InternalError.FundSummaryAcctNoInconsistentError"
//  INTERNALERROR_PARAMETERERROR = "InternalError.ParameterError"
//  INTERNALERROR_SAVEDBERROR = "InternalError.SaveDBError"
//  INTERNALERROR_SUBACCOUNTNOTFOUNDERROR = "InternalError.SubAccountNotFoundError"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  INTERNALERROR_UNKOWNERROR = "InternalError.UnkownError"
func (c *Client) BindRelateAcctSmallAmount(request *BindRelateAcctSmallAmountRequest) (response *BindRelateAcctSmallAmountResponse, err error) {
    if request == nil {
        request = NewBindRelateAcctSmallAmountRequest()
    }
    
    response = NewBindRelateAcctSmallAmountResponse()
    err = c.Send(request, response)
    return
}

func NewBindRelateAcctUnionPayRequest() (request *BindRelateAcctUnionPayRequest) {
    request = &BindRelateAcctUnionPayRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "BindRelateAcctUnionPay")
    
    
    return
}

func NewBindRelateAcctUnionPayResponse() (response *BindRelateAcctUnionPayResponse) {
    response = &BindRelateAcctUnionPayResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// BindRelateAcctUnionPay
// 会员绑定提现账户-银联鉴权。用于会员申请绑定提现账户，申请后银行前往银联验证卡信息：姓名、证件、卡号、银行预留手机是否相符，相符则发送给会员手机动态验证码并返回成功，不相符则返回失败。
//
// 平台接收到银行返回成功后，进入输入动态验证码的页面，有效期120秒，若120秒未输入，客户可点击重新发送动态验证码，这个步骤重新调用该接口即可。
//
// 平安银行的账户，大小额行号和超级网银号都不用送。
//
// 超级网银号：单笔转账金额不超过5万，不限制笔数，只用选XX银行，不用具体到支行，可实时知道对方是否收款成功。
//
// 大小额联行号：单笔转账可超过5万，需具体到支行，不能实时知道对方是否收款成功。金额超过5万的，在工作日的8点30-17点间才会成功。
//
// 可能返回的错误码:
//  FAILEDOPERATION_BACKENDERROR = "FailedOperation.BackendError"
//  INTERNALERROR_BACKENDERROR = "InternalError.BackendError"
//  INTERNALERROR_DBACCESSERROR = "InternalError.DBAccessError"
//  INTERNALERROR_FUNDSUMMARYACCTNOINCONSISTENTERROR = "InternalError.FundSummaryAcctNoInconsistentError"
//  INTERNALERROR_PARAMETERERROR = "InternalError.ParameterError"
//  INTERNALERROR_SAVEDBERROR = "InternalError.SaveDBError"
//  INTERNALERROR_SUBACCOUNTNOTFOUNDERROR = "InternalError.SubAccountNotFoundError"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  INTERNALERROR_UNKOWNERROR = "InternalError.UnkownError"
func (c *Client) BindRelateAcctUnionPay(request *BindRelateAcctUnionPayRequest) (response *BindRelateAcctUnionPayResponse, err error) {
    if request == nil {
        request = NewBindRelateAcctUnionPayRequest()
    }
    
    response = NewBindRelateAcctUnionPayResponse()
    err = c.Send(request, response)
    return
}

func NewCheckAcctRequest() (request *CheckAcctRequest) {
    request = &CheckAcctRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "CheckAcct")
    
    
    return
}

func NewCheckAcctResponse() (response *CheckAcctResponse) {
    response = &CheckAcctResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CheckAcct
// 商户绑定提现银行卡的验证接口
//
// 可能返回的错误码:
//  FAILEDOPERATION_BANKFAILED = "FailedOperation.BankFailed"
//  FAILEDOPERATION_SERVICEERROR = "FailedOperation.ServiceError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER_ = "MissingParameter."
func (c *Client) CheckAcct(request *CheckAcctRequest) (response *CheckAcctResponse, err error) {
    if request == nil {
        request = NewCheckAcctRequest()
    }
    
    response = NewCheckAcctResponse()
    err = c.Send(request, response)
    return
}

func NewCheckAmountRequest() (request *CheckAmountRequest) {
    request = &CheckAmountRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "CheckAmount")
    
    
    return
}

func NewCheckAmountResponse() (response *CheckAmountResponse) {
    response = &CheckAmountResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CheckAmount
// 验证鉴权金额。此接口可受理BindRelateAcctSmallAmount接口发起的转账金额（往账鉴权方式）的验证处理。若所回填的验证金额验证通过，则会绑定原申请中的银行账户作为提现账户。通过此接口也可以查得BindRelateAcctSmallAmount接口发起的来账鉴权方式的申请的当前状态。
//
// 可能返回的错误码:
//  FAILEDOPERATION_BACKENDERROR = "FailedOperation.BackendError"
//  INTERNALERROR_BACKENDERROR = "InternalError.BackendError"
//  INTERNALERROR_DBACCESSERROR = "InternalError.DBAccessError"
//  INTERNALERROR_FUNDSUMMARYACCTNOINCONSISTENTERROR = "InternalError.FundSummaryAcctNoInconsistentError"
//  INTERNALERROR_PARAMETERERROR = "InternalError.ParameterError"
//  INTERNALERROR_SAVEDBERROR = "InternalError.SaveDBError"
//  INTERNALERROR_SUBACCOUNTNOTFOUNDERROR = "InternalError.SubAccountNotFoundError"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  INTERNALERROR_UNKOWNERROR = "InternalError.UnkownError"
func (c *Client) CheckAmount(request *CheckAmountRequest) (response *CheckAmountResponse, err error) {
    if request == nil {
        request = NewCheckAmountRequest()
    }
    
    response = NewCheckAmountResponse()
    err = c.Send(request, response)
    return
}

func NewCloseOrderRequest() (request *CloseOrderRequest) {
    request = &CloseOrderRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "CloseOrder")
    
    
    return
}

func NewCloseOrderResponse() (response *CloseOrderResponse) {
    response = &CloseOrderResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CloseOrder
// 通过此接口关闭此前已创建的订单，关闭后，用户将无法继续付款。仅能关闭创建后未支付的订单
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_APPDENY = "FailedOperation.AppDeny"
//  FAILEDOPERATION_NORECORD = "FailedOperation.NoRecord"
//  FAILEDOPERATION_ORDERNOTACTIVATED = "FailedOperation.OrderNotActivated"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CloseOrder(request *CloseOrderRequest) (response *CloseOrderResponse, err error) {
    if request == nil {
        request = NewCloseOrderRequest()
    }
    
    response = NewCloseOrderResponse()
    err = c.Send(request, response)
    return
}

func NewConfirmOrderRequest() (request *ConfirmOrderRequest) {
    request = &ConfirmOrderRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "ConfirmOrder")
    
    
    return
}

func NewConfirmOrderResponse() (response *ConfirmOrderResponse) {
    response = &ConfirmOrderResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ConfirmOrder
// 云鉴-消费订单确认接口
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDPARAMETER = "FailedOperation.InvalidParameter"
//  FAILEDOPERATION_MISSINGPARAMETER = "FailedOperation.MissingParameter"
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
func (c *Client) ConfirmOrder(request *ConfirmOrderRequest) (response *ConfirmOrderResponse, err error) {
    if request == nil {
        request = NewConfirmOrderRequest()
    }
    
    response = NewConfirmOrderResponse()
    err = c.Send(request, response)
    return
}

func NewContractOrderRequest() (request *ContractOrderRequest) {
    request = &ContractOrderRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "ContractOrder")
    
    
    return
}

func NewContractOrderResponse() (response *ContractOrderResponse) {
    response = &ContractOrderResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ContractOrder
// 应用需要先带上签约信息调用本接口生成支付订单号，并将应答的PayInfo透传给聚鑫SDK，拉起客户端（包括微信公众号/微信小程序/客户端App）支付。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_APPDENY = "FailedOperation.AppDeny"
//  FAILEDOPERATION_BACKENDERROR = "FailedOperation.BackendError"
//  FAILEDOPERATION_CHANNELDENY = "FailedOperation.ChannelDeny"
//  FAILEDOPERATION_DBCLIENTCONNECTFAILED = "FailedOperation.DbClientConnectFailed"
//  FAILEDOPERATION_DBCLIENTINSERTTFAILED = "FailedOperation.DbClientInserttFailed"
//  FAILEDOPERATION_EXTERNALMERCHANTCONTRACTINFOCONFIGNOFOUND = "FailedOperation.ExternalMerchantContractInfoConfigNoFound"
//  FAILEDOPERATION_EXTERNALMERCHANTINDEXCONFIGNOFOUND = "FailedOperation.ExternalMerchantIndexConfigNoFound"
//  FAILEDOPERATION_MARSHALERROR = "FailedOperation.MarshalError"
//  FAILEDOPERATION_OCCOMPLETEDORDER = "FailedOperation.OcCompletedOrder"
//  FAILEDOPERATION_OCREPEATORDER = "FailedOperation.OcRepeatOrder"
//  FAILEDOPERATION_PARAMETERERROR = "FailedOperation.ParameterError"
//  FAILEDOPERATION_PARENTAPPIDERROR = "FailedOperation.ParentAppIdError"
//  FAILEDOPERATION_PORTALERROR = "FailedOperation.PortalError"
//  FAILEDOPERATION_UNMARSHALERROR = "FailedOperation.UnmarshalError"
//  FAILEDOPERATION_WECHATERROR = "FailedOperation.WechatError"
//  INTERNALERROR_BACKENDERROR = "InternalError.BackendError"
//  INTERNALERROR_BACKENDINTERNALERROR = "InternalError.BackendInternalError"
//  INTERNALERROR_PARAMETERERROR = "InternalError.ParameterError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ContractOrder(request *ContractOrderRequest) (response *ContractOrderResponse, err error) {
    if request == nil {
        request = NewContractOrderRequest()
    }
    
    response = NewContractOrderResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAcctRequest() (request *CreateAcctRequest) {
    request = &CreateAcctRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "CreateAcct")
    
    
    return
}

func NewCreateAcctResponse() (response *CreateAcctResponse) {
    response = &CreateAcctResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateAcct
// 子商户入驻聚鑫平台
//
// 可能返回的错误码:
//  FAILEDOPERATION_BANKFAILED = "FailedOperation.BankFailed"
//  FAILEDOPERATION_SERVICEERROR = "FailedOperation.ServiceError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER_ = "MissingParameter."
func (c *Client) CreateAcct(request *CreateAcctRequest) (response *CreateAcctResponse, err error) {
    if request == nil {
        request = NewCreateAcctRequest()
    }
    
    response = NewCreateAcctResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAgentTaxPaymentInfosRequest() (request *CreateAgentTaxPaymentInfosRequest) {
    request = &CreateAgentTaxPaymentInfosRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "CreateAgentTaxPaymentInfos")
    
    
    return
}

func NewCreateAgentTaxPaymentInfosResponse() (response *CreateAgentTaxPaymentInfosResponse) {
    response = &CreateAgentTaxPaymentInfosResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateAgentTaxPaymentInfos
// 直播平台-代理商完税信息录入
//
// 可能返回的错误码:
//  INTERNALERROR_DELETEDBERROR = "InternalError.DeleteDBError"
//  INTERNALERROR_SAVEDBERROR = "InternalError.SaveDBError"
//  INTERNALERROR_UNKOWNERROR = "InternalError.UnkownError"
//  INVALIDPARAMETER_LACKPARAMETER = "InvalidParameter.LackParameter"
//  RESOURCENOTFOUND_BATCHINFONOTFOUND = "ResourceNotFound.BatchInfoNotFound"
//  RESOURCENOTFOUND_PLATFORMINFONOTFOUND = "ResourceNotFound.PlatformInfoNotFound"
func (c *Client) CreateAgentTaxPaymentInfos(request *CreateAgentTaxPaymentInfosRequest) (response *CreateAgentTaxPaymentInfosResponse, err error) {
    if request == nil {
        request = NewCreateAgentTaxPaymentInfosRequest()
    }
    
    response = NewCreateAgentTaxPaymentInfosResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAnchorRequest() (request *CreateAnchorRequest) {
    request = &CreateAnchorRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "CreateAnchor")
    
    
    return
}

func NewCreateAnchorResponse() (response *CreateAnchorResponse) {
    response = &CreateAnchorResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateAnchor
// 直播平台-主播入驻
//
// 可能返回的错误码:
//  FAILEDOPERATION_CREATEAGENT = "FailedOperation.CreateAgent"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  MISSINGPARAMETER_ACTION = "MissingParameter.Action"
//  MISSINGPARAMETER_APPID = "MissingParameter.AppId"
func (c *Client) CreateAnchor(request *CreateAnchorRequest) (response *CreateAnchorResponse, err error) {
    if request == nil {
        request = NewCreateAnchorRequest()
    }
    
    response = NewCreateAnchorResponse()
    err = c.Send(request, response)
    return
}

func NewCreateBatchPaymentRequest() (request *CreateBatchPaymentRequest) {
    request = &CreateBatchPaymentRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "CreateBatchPayment")
    
    
    return
}

func NewCreateBatchPaymentResponse() (response *CreateBatchPaymentResponse) {
    response = &CreateBatchPaymentResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateBatchPayment
// 灵云-批量主播转账接口
//
// 可能返回的错误码:
//  FAILEDOPERATION_CREATEAGENT = "FailedOperation.CreateAgent"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  MISSINGPARAMETER_ACTION = "MissingParameter.Action"
//  MISSINGPARAMETER_APPID = "MissingParameter.AppId"
func (c *Client) CreateBatchPayment(request *CreateBatchPaymentRequest) (response *CreateBatchPaymentResponse, err error) {
    if request == nil {
        request = NewCreateBatchPaymentRequest()
    }
    
    response = NewCreateBatchPaymentResponse()
    err = c.Send(request, response)
    return
}

func NewCreateCustAcctIdRequest() (request *CreateCustAcctIdRequest) {
    request = &CreateCustAcctIdRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "CreateCustAcctId")
    
    
    return
}

func NewCreateCustAcctIdResponse() (response *CreateCustAcctIdResponse) {
    response = &CreateCustAcctIdResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateCustAcctId
// 会员子账户开立。会员在银行注册，并开立会员子账户，交易网会员代码即会员在平台端系统的会员编号。
//
// 平台需保存银行返回的子账户账号，后续交易接口都会用到。会员属性字段为预留扩展字段，当前必须送默认值。
//
// 可能返回的错误码:
//  FAILEDOPERATION_BACKENDERROR = "FailedOperation.BackendError"
//  FAILEDOPERATION_SDKERROR = "FailedOperation.SDKError"
//  INTERNALERROR_BACKENDERROR = "InternalError.BackendError"
//  INTERNALERROR_DBACCESSERROR = "InternalError.DBAccessError"
//  INTERNALERROR_FUNDSUMMARYACCTNOINCONSISTENTERROR = "InternalError.FundSummaryAcctNoInconsistentError"
//  INTERNALERROR_PARAMETERERROR = "InternalError.ParameterError"
//  INTERNALERROR_SAVEDBERROR = "InternalError.SaveDBError"
//  INTERNALERROR_SUBACCOUNTNOTFOUNDERROR = "InternalError.SubAccountNotFoundError"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  INTERNALERROR_UNKOWNERROR = "InternalError.UnkownError"
func (c *Client) CreateCustAcctId(request *CreateCustAcctIdRequest) (response *CreateCustAcctIdResponse, err error) {
    if request == nil {
        request = NewCreateCustAcctIdRequest()
    }
    
    response = NewCreateCustAcctIdResponse()
    err = c.Send(request, response)
    return
}

func NewCreateExternalAnchorRequest() (request *CreateExternalAnchorRequest) {
    request = &CreateExternalAnchorRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "CreateExternalAnchor")
    
    
    return
}

func NewCreateExternalAnchorResponse() (response *CreateExternalAnchorResponse) {
    response = &CreateExternalAnchorResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateExternalAnchor
// 灵云-主播入驻
//
// 可能返回的错误码:
//  FAILEDOPERATION_BACKENDERROR = "FailedOperation.BackendError"
//  FAILEDOPERATION_SDKERROR = "FailedOperation.SDKError"
//  INTERNALERROR_BACKENDERROR = "InternalError.BackendError"
//  INTERNALERROR_DBACCESSERROR = "InternalError.DBAccessError"
//  INTERNALERROR_FUNDSUMMARYACCTNOINCONSISTENTERROR = "InternalError.FundSummaryAcctNoInconsistentError"
//  INTERNALERROR_PARAMETERERROR = "InternalError.ParameterError"
//  INTERNALERROR_SAVEDBERROR = "InternalError.SaveDBError"
//  INTERNALERROR_SUBACCOUNTNOTFOUNDERROR = "InternalError.SubAccountNotFoundError"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  INTERNALERROR_UNKOWNERROR = "InternalError.UnkownError"
func (c *Client) CreateExternalAnchor(request *CreateExternalAnchorRequest) (response *CreateExternalAnchorResponse, err error) {
    if request == nil {
        request = NewCreateExternalAnchorRequest()
    }
    
    response = NewCreateExternalAnchorResponse()
    err = c.Send(request, response)
    return
}

func NewCreateInvoiceRequest() (request *CreateInvoiceRequest) {
    request = &CreateInvoiceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "CreateInvoice")
    
    
    return
}

func NewCreateInvoiceResponse() (response *CreateInvoiceResponse) {
    response = &CreateInvoiceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateInvoice
// 智慧零售-发票开具
//
// 可能返回的错误码:
//  FAILEDOPERATION_BACKENDERROR = "FailedOperation.BackendError"
//  FAILEDOPERATION_INVOICEEXIST = "FailedOperation.InvoiceExist"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
//  INTERNALERROR_BACKENDERROR = "InternalError.BackendError"
//  INTERNALERROR_DUPLICATEKEYERROR = "InternalError.DuplicateKeyError"
//  INTERNALERROR_INVOICEEXIST = "InternalError.InvoiceExist"
//  INTERNALERROR_SANDBOXACCESSERROR = "InternalError.SandBoxAccessError"
//  INTERNALERROR_SAVEDBERROR = "InternalError.SaveDBError"
//  INTERNALERROR_SIGGENERROR = "InternalError.SigGenError"
//  INTERNALERROR_UNKOWNERROR = "InternalError.UnkownError"
//  INVALIDPARAMETER_LACKPARAMETER = "InvalidParameter.LackParameter"
//  INVALIDPARAMETER_UNSUPPORTEDPARAMETER = "InvalidParameter.UnsupportedParameter"
//  RESOURCEINSUFFICIENT_THREADPOOLREJECT = "ResourceInsufficient.ThreadPoolReject"
//  RESOURCENOTFOUND_INVOICENOTFOUND = "ResourceNotFound.InvoiceNotFound"
//  RESOURCENOTFOUND_MERCHANTINFONOTFOUND = "ResourceNotFound.MerchantInfoNotFound"
//  RESOURCENOTFOUND_PLATFORMINFONOTFOUND = "ResourceNotFound.PlatformInfoNotFound"
func (c *Client) CreateInvoice(request *CreateInvoiceRequest) (response *CreateInvoiceResponse, err error) {
    if request == nil {
        request = NewCreateInvoiceRequest()
    }
    
    response = NewCreateInvoiceResponse()
    err = c.Send(request, response)
    return
}

func NewCreateInvoiceV2Request() (request *CreateInvoiceV2Request) {
    request = &CreateInvoiceV2Request{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "CreateInvoiceV2")
    
    
    return
}

func NewCreateInvoiceV2Response() (response *CreateInvoiceV2Response) {
    response = &CreateInvoiceV2Response{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateInvoiceV2
// 智慧零售-发票开具V2
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_INVOICENOTFOUND = "ResourceNotFound.InvoiceNotFound"
func (c *Client) CreateInvoiceV2(request *CreateInvoiceV2Request) (response *CreateInvoiceV2Response, err error) {
    if request == nil {
        request = NewCreateInvoiceV2Request()
    }
    
    response = NewCreateInvoiceV2Response()
    err = c.Send(request, response)
    return
}

func NewCreateMerchantRequest() (request *CreateMerchantRequest) {
    request = &CreateMerchantRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "CreateMerchant")
    
    
    return
}

func NewCreateMerchantResponse() (response *CreateMerchantResponse) {
    response = &CreateMerchantResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateMerchant
// 智慧零售-商户注册
//
// 可能返回的错误码:
//  INTERNALERROR_DUPLICATEKEYERROR = "InternalError.DuplicateKeyError"
//  INTERNALERROR_SAVEDBERROR = "InternalError.SaveDBError"
//  INTERNALERROR_SIGGENERROR = "InternalError.SigGenError"
//  INTERNALERROR_UNKOWNERROR = "InternalError.UnkownError"
//  INVALIDPARAMETER_LACKPARAMETER = "InvalidParameter.LackParameter"
//  RESOURCENOTFOUND_INVOICENOTFOUND = "ResourceNotFound.InvoiceNotFound"
//  RESOURCENOTFOUND_MERCHANTINFONOTFOUND = "ResourceNotFound.MerchantInfoNotFound"
//  RESOURCENOTFOUND_PLATFORMINFONOTFOUND = "ResourceNotFound.PlatformInfoNotFound"
func (c *Client) CreateMerchant(request *CreateMerchantRequest) (response *CreateMerchantResponse, err error) {
    if request == nil {
        request = NewCreateMerchantRequest()
    }
    
    response = NewCreateMerchantResponse()
    err = c.Send(request, response)
    return
}

func NewCreateOrderRequest() (request *CreateOrderRequest) {
    request = &CreateOrderRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "CreateOrder")
    
    
    return
}

func NewCreateOrderResponse() (response *CreateOrderResponse) {
    response = &CreateOrderResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateOrder
// 云鉴-消费订单发起的接口
//
// 可能返回的错误码:
//  FAILEDOPERATION_CREATEORDERERROR = "FailedOperation.CreateOrderError"
//  FAILEDOPERATION_CREATEORDERUNKNOWN = "FailedOperation.CreateOrderUnknown"
//  FAILEDOPERATION_INVALIDPARAMETER = "FailedOperation.InvalidParameter"
//  FAILEDOPERATION_MISSINGPARAMETER = "FailedOperation.MissingParameter"
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
func (c *Client) CreateOrder(request *CreateOrderRequest) (response *CreateOrderResponse, err error) {
    if request == nil {
        request = NewCreateOrderRequest()
    }
    
    response = NewCreateOrderResponse()
    err = c.Send(request, response)
    return
}

func NewCreatePayMerchantRequest() (request *CreatePayMerchantRequest) {
    request = &CreatePayMerchantRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "CreatePayMerchant")
    
    
    return
}

func NewCreatePayMerchantResponse() (response *CreatePayMerchantResponse) {
    response = &CreatePayMerchantResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreatePayMerchant
// 商户新增的接口
//
// 可能返回的错误码:
//  FAILEDOPERATION_ADDMERCHANTFAILED = "FailedOperation.AddMerchantFailed"
//  FAILEDOPERATION_BANKFAILED = "FailedOperation.BankFailed"
//  FAILEDOPERATION_INVALIDPARAMETER = "FailedOperation.InvalidParameter"
//  FAILEDOPERATION_MERCHANTCHECKFAILED = "FailedOperation.MerchantCheckFailed"
//  FAILEDOPERATION_MERCHANTCREATEFAILED = "FailedOperation.MerchantCreateFailed"
//  FAILEDOPERATION_MERCHANTEXIST = "FailedOperation.MerchantExist"
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
func (c *Client) CreatePayMerchant(request *CreatePayMerchantRequest) (response *CreatePayMerchantResponse, err error) {
    if request == nil {
        request = NewCreatePayMerchantRequest()
    }
    
    response = NewCreatePayMerchantResponse()
    err = c.Send(request, response)
    return
}

func NewCreateRedInvoiceRequest() (request *CreateRedInvoiceRequest) {
    request = &CreateRedInvoiceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "CreateRedInvoice")
    
    
    return
}

func NewCreateRedInvoiceResponse() (response *CreateRedInvoiceResponse) {
    response = &CreateRedInvoiceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateRedInvoice
// 智慧零售-发票红冲
//
// 可能返回的错误码:
//  FAILEDOPERATION_BACKENDERROR = "FailedOperation.BackendError"
//  FAILEDOPERATION_INVOICEEXIST = "FailedOperation.InvoiceExist"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
//  INTERNALERROR_BACKENDERROR = "InternalError.BackendError"
//  INTERNALERROR_DUPLICATEKEYERROR = "InternalError.DuplicateKeyError"
//  INTERNALERROR_SANDBOXACCESSERROR = "InternalError.SandBoxAccessError"
//  INTERNALERROR_SAVEDBERROR = "InternalError.SaveDBError"
//  INTERNALERROR_SIGGENERROR = "InternalError.SigGenError"
//  INTERNALERROR_UNKOWNERROR = "InternalError.UnkownError"
//  INVALIDPARAMETER_LACKPARAMETER = "InvalidParameter.LackParameter"
//  INVALIDPARAMETER_UNSUPPORTEDPARAMETER = "InvalidParameter.UnsupportedParameter"
//  RESOURCEINSUFFICIENT_THREADPOOLREJECT = "ResourceInsufficient.ThreadPoolReject"
//  RESOURCENOTFOUND_INVOICENOTFOUND = "ResourceNotFound.InvoiceNotFound"
//  RESOURCENOTFOUND_MERCHANTINFONOTFOUND = "ResourceNotFound.MerchantInfoNotFound"
//  RESOURCENOTFOUND_PLATFORMINFONOTFOUND = "ResourceNotFound.PlatformInfoNotFound"
func (c *Client) CreateRedInvoice(request *CreateRedInvoiceRequest) (response *CreateRedInvoiceResponse, err error) {
    if request == nil {
        request = NewCreateRedInvoiceRequest()
    }
    
    response = NewCreateRedInvoiceResponse()
    err = c.Send(request, response)
    return
}

func NewCreateRedInvoiceV2Request() (request *CreateRedInvoiceV2Request) {
    request = &CreateRedInvoiceV2Request{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "CreateRedInvoiceV2")
    
    
    return
}

func NewCreateRedInvoiceV2Response() (response *CreateRedInvoiceV2Response) {
    response = &CreateRedInvoiceV2Response{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateRedInvoiceV2
// 智慧零售-发票红冲V2
//
// 可能返回的错误码:
//  FAILEDOPERATION_BACKENDERROR = "FailedOperation.BackendError"
//  FAILEDOPERATION_INVOICEEXIST = "FailedOperation.InvoiceExist"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
//  INTERNALERROR_BACKENDERROR = "InternalError.BackendError"
//  INTERNALERROR_DUPLICATEKEYERROR = "InternalError.DuplicateKeyError"
//  INTERNALERROR_SANDBOXACCESSERROR = "InternalError.SandBoxAccessError"
//  INTERNALERROR_SAVEDBERROR = "InternalError.SaveDBError"
//  INTERNALERROR_SIGGENERROR = "InternalError.SigGenError"
//  INTERNALERROR_UNKOWNERROR = "InternalError.UnkownError"
//  INVALIDPARAMETER_LACKPARAMETER = "InvalidParameter.LackParameter"
//  INVALIDPARAMETER_UNSUPPORTEDPARAMETER = "InvalidParameter.UnsupportedParameter"
//  RESOURCEINSUFFICIENT_THREADPOOLREJECT = "ResourceInsufficient.ThreadPoolReject"
//  RESOURCENOTFOUND_INVOICENOTFOUND = "ResourceNotFound.InvoiceNotFound"
//  RESOURCENOTFOUND_MERCHANTINFONOTFOUND = "ResourceNotFound.MerchantInfoNotFound"
//  RESOURCENOTFOUND_PLATFORMINFONOTFOUND = "ResourceNotFound.PlatformInfoNotFound"
func (c *Client) CreateRedInvoiceV2(request *CreateRedInvoiceV2Request) (response *CreateRedInvoiceV2Response, err error) {
    if request == nil {
        request = NewCreateRedInvoiceV2Request()
    }
    
    response = NewCreateRedInvoiceV2Response()
    err = c.Send(request, response)
    return
}

func NewCreateSinglePayRequest() (request *CreateSinglePayRequest) {
    request = &CreateSinglePayRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "CreateSinglePay")
    
    
    return
}

func NewCreateSinglePayResponse() (response *CreateSinglePayResponse) {
    response = &CreateSinglePayResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateSinglePay
// 银企直连-单笔支付接口
//
// 可能返回的错误码:
//  INTERNALERROR_BACKENDERROR = "InternalError.BackendError"
//  INTERNALERROR_SANDBOXACCESSERROR = "InternalError.SandBoxAccessError"
//  INTERNALERROR_SIGGENERROR = "InternalError.SigGenError"
//  INTERNALERROR_UNKOWNERROR = "InternalError.UnkownError"
//  INVALIDPARAMETER_LACKPARAMETER = "InvalidParameter.LackParameter"
//  INVALIDPARAMETER_UNSUPPORTEDPARAMETER = "InvalidParameter.UnsupportedParameter"
//  RESOURCENOTFOUND_PLATFORMINFONOTFOUND = "ResourceNotFound.PlatformInfoNotFound"
func (c *Client) CreateSinglePay(request *CreateSinglePayRequest) (response *CreateSinglePayResponse, err error) {
    if request == nil {
        request = NewCreateSinglePayRequest()
    }
    
    response = NewCreateSinglePayResponse()
    err = c.Send(request, response)
    return
}

func NewCreateSinglePaymentRequest() (request *CreateSinglePaymentRequest) {
    request = &CreateSinglePaymentRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "CreateSinglePayment")
    
    
    return
}

func NewCreateSinglePaymentResponse() (response *CreateSinglePaymentResponse) {
    response = &CreateSinglePaymentResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateSinglePayment
// 灵云-单笔主播转账接口
//
// 可能返回的错误码:
//  INTERNALERROR_BACKENDERROR = "InternalError.BackendError"
//  INTERNALERROR_SANDBOXACCESSERROR = "InternalError.SandBoxAccessError"
//  INTERNALERROR_SIGGENERROR = "InternalError.SigGenError"
//  INTERNALERROR_UNKOWNERROR = "InternalError.UnkownError"
//  INVALIDPARAMETER_LACKPARAMETER = "InvalidParameter.LackParameter"
//  INVALIDPARAMETER_UNSUPPORTEDPARAMETER = "InvalidParameter.UnsupportedParameter"
//  RESOURCENOTFOUND_PLATFORMINFONOTFOUND = "ResourceNotFound.PlatformInfoNotFound"
func (c *Client) CreateSinglePayment(request *CreateSinglePaymentRequest) (response *CreateSinglePaymentResponse, err error) {
    if request == nil {
        request = NewCreateSinglePaymentRequest()
    }
    
    response = NewCreateSinglePaymentResponse()
    err = c.Send(request, response)
    return
}

func NewCreateTransferBatchRequest() (request *CreateTransferBatchRequest) {
    request = &CreateTransferBatchRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "CreateTransferBatch")
    
    
    return
}

func NewCreateTransferBatchResponse() (response *CreateTransferBatchResponse) {
    response = &CreateTransferBatchResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateTransferBatch
// 微信商户发起批量转账
//
// 可能返回的错误码:
//  FAILEDOPERATION_ALREADYEXISTS = "FailedOperation.AlreadyExists"
//  FAILEDOPERATION_APPIDMCHIDNOTMATCH = "FailedOperation.AppidMchidNotMatch"
//  FAILEDOPERATION_FREQUENCYLIMITED = "FailedOperation.FrequencyLimited"
//  FAILEDOPERATION_INVALIDREQUEST = "FailedOperation.InvalidRequest"
//  FAILEDOPERATION_NOAUTH = "FailedOperation.NoAuth"
//  FAILEDOPERATION_NOTENOUGH = "FailedOperation.NotEnough"
//  FAILEDOPERATION_NOTFOUND = "FailedOperation.NotFound"
//  FAILEDOPERATION_PARAMERROR = "FailedOperation.ParamError"
//  FAILEDOPERATION_QUOTAEXCEED = "FailedOperation.QuotaExceed"
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR_ = "InternalError."
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER_ = "MissingParameter."
func (c *Client) CreateTransferBatch(request *CreateTransferBatchRequest) (response *CreateTransferBatchResponse, err error) {
    if request == nil {
        request = NewCreateTransferBatchRequest()
    }
    
    response = NewCreateTransferBatchResponse()
    err = c.Send(request, response)
    return
}

func NewDeduceQuotaRequest() (request *DeduceQuotaRequest) {
    request = &DeduceQuotaRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "DeduceQuota")
    
    
    return
}

func NewDeduceQuotaResponse() (response *DeduceQuotaResponse) {
    response = &DeduceQuotaResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeduceQuota
// 直播平台-扣减额度
//
// 可能返回的错误码:
//  FAILEDOPERATION_ALREADYEXISTS = "FailedOperation.AlreadyExists"
//  FAILEDOPERATION_APPIDMCHIDNOTMATCH = "FailedOperation.AppidMchidNotMatch"
//  FAILEDOPERATION_FREQUENCYLIMITED = "FailedOperation.FrequencyLimited"
//  FAILEDOPERATION_INVALIDREQUEST = "FailedOperation.InvalidRequest"
//  FAILEDOPERATION_NOAUTH = "FailedOperation.NoAuth"
//  FAILEDOPERATION_NOTENOUGH = "FailedOperation.NotEnough"
//  FAILEDOPERATION_NOTFOUND = "FailedOperation.NotFound"
//  FAILEDOPERATION_PARAMERROR = "FailedOperation.ParamError"
//  FAILEDOPERATION_QUOTAEXCEED = "FailedOperation.QuotaExceed"
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR_ = "InternalError."
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER_ = "MissingParameter."
func (c *Client) DeduceQuota(request *DeduceQuotaRequest) (response *DeduceQuotaResponse, err error) {
    if request == nil {
        request = NewDeduceQuotaRequest()
    }
    
    response = NewDeduceQuotaResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAgentTaxPaymentInfoRequest() (request *DeleteAgentTaxPaymentInfoRequest) {
    request = &DeleteAgentTaxPaymentInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "DeleteAgentTaxPaymentInfo")
    
    
    return
}

func NewDeleteAgentTaxPaymentInfoResponse() (response *DeleteAgentTaxPaymentInfoResponse) {
    response = &DeleteAgentTaxPaymentInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteAgentTaxPaymentInfo
// 直播平台-删除代理商完税信息
//
// 可能返回的错误码:
//  INTERNALERROR_DELETEDBERROR = "InternalError.DeleteDBError"
//  INTERNALERROR_SAVEDBERROR = "InternalError.SaveDBError"
//  INTERNALERROR_UNKOWNERROR = "InternalError.UnkownError"
//  INVALIDPARAMETER_LACKPARAMETER = "InvalidParameter.LackParameter"
//  RESOURCENOTFOUND_BATCHINFONOTFOUND = "ResourceNotFound.BatchInfoNotFound"
//  RESOURCENOTFOUND_PLATFORMINFONOTFOUND = "ResourceNotFound.PlatformInfoNotFound"
func (c *Client) DeleteAgentTaxPaymentInfo(request *DeleteAgentTaxPaymentInfoRequest) (response *DeleteAgentTaxPaymentInfoResponse, err error) {
    if request == nil {
        request = NewDeleteAgentTaxPaymentInfoRequest()
    }
    
    response = NewDeleteAgentTaxPaymentInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAgentTaxPaymentInfosRequest() (request *DeleteAgentTaxPaymentInfosRequest) {
    request = &DeleteAgentTaxPaymentInfosRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "DeleteAgentTaxPaymentInfos")
    
    
    return
}

func NewDeleteAgentTaxPaymentInfosResponse() (response *DeleteAgentTaxPaymentInfosResponse) {
    response = &DeleteAgentTaxPaymentInfosResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteAgentTaxPaymentInfos
// 直播平台-删除代理商完税信息
//
// 可能返回的错误码:
//  INTERNALERROR_DELETEDBERROR = "InternalError.DeleteDBError"
//  INTERNALERROR_SAVEDBERROR = "InternalError.SaveDBError"
//  INTERNALERROR_UNKOWNERROR = "InternalError.UnkownError"
//  INVALIDPARAMETER_LACKPARAMETER = "InvalidParameter.LackParameter"
//  RESOURCENOTFOUND_BATCHINFONOTFOUND = "ResourceNotFound.BatchInfoNotFound"
//  RESOURCENOTFOUND_PLATFORMINFONOTFOUND = "ResourceNotFound.PlatformInfoNotFound"
func (c *Client) DeleteAgentTaxPaymentInfos(request *DeleteAgentTaxPaymentInfosRequest) (response *DeleteAgentTaxPaymentInfosResponse, err error) {
    if request == nil {
        request = NewDeleteAgentTaxPaymentInfosRequest()
    }
    
    response = NewDeleteAgentTaxPaymentInfosResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeChargeDetailRequest() (request *DescribeChargeDetailRequest) {
    request = &DescribeChargeDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "DescribeChargeDetail")
    
    
    return
}

func NewDescribeChargeDetailResponse() (response *DescribeChargeDetailResponse) {
    response = &DescribeChargeDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeChargeDetail
// 查询充值明细接口
//
// 可能返回的错误码:
//  FAILEDOPERATION_ACTIONINVALID = "FailedOperation.ActionInvalid"
//  FAILEDOPERATION_PABANKERROR = "FailedOperation.PABankError"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  INTERNALERROR_UNKOWNERROR = "InternalError.UnkownError"
//  MISSINGPARAMETER_ACTION = "MissingParameter.Action"
//  MISSINGPARAMETER_APPID = "MissingParameter.AppId"
func (c *Client) DescribeChargeDetail(request *DescribeChargeDetailRequest) (response *DescribeChargeDetailResponse, err error) {
    if request == nil {
        request = NewDescribeChargeDetailRequest()
    }
    
    response = NewDescribeChargeDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeOrderStatusRequest() (request *DescribeOrderStatusRequest) {
    request = &DescribeOrderStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "DescribeOrderStatus")
    
    
    return
}

func NewDescribeOrderStatusResponse() (response *DescribeOrderStatusResponse) {
    response = &DescribeOrderStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeOrderStatus
// 查询单笔订单交易状态
//
// 可能返回的错误码:
//  FAILEDOPERATION_ACTIONINVALID = "FailedOperation.ActionInvalid"
//  FAILEDOPERATION_PABANKERROR = "FailedOperation.PABankError"
func (c *Client) DescribeOrderStatus(request *DescribeOrderStatusRequest) (response *DescribeOrderStatusResponse, err error) {
    if request == nil {
        request = NewDescribeOrderStatusRequest()
    }
    
    response = NewDescribeOrderStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDistributeAccreditQueryRequest() (request *DistributeAccreditQueryRequest) {
    request = &DistributeAccreditQueryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "DistributeAccreditQuery")
    
    
    return
}

func NewDistributeAccreditQueryResponse() (response *DistributeAccreditQueryResponse) {
    response = &DistributeAccreditQueryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DistributeAccreditQuery
// 云支付-分账授权申请查询接口
//
// 可能返回的错误码:
//  FAILEDOPERATION_ACTIONINVALID = "FailedOperation.ActionInvalid"
//  FAILEDOPERATION_PABANKERROR = "FailedOperation.PABankError"
func (c *Client) DistributeAccreditQuery(request *DistributeAccreditQueryRequest) (response *DistributeAccreditQueryResponse, err error) {
    if request == nil {
        request = NewDistributeAccreditQueryRequest()
    }
    
    response = NewDistributeAccreditQueryResponse()
    err = c.Send(request, response)
    return
}

func NewDistributeAccreditTlinxRequest() (request *DistributeAccreditTlinxRequest) {
    request = &DistributeAccreditTlinxRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "DistributeAccreditTlinx")
    
    
    return
}

func NewDistributeAccreditTlinxResponse() (response *DistributeAccreditTlinxResponse) {
    response = &DistributeAccreditTlinxResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DistributeAccreditTlinx
// 云支付-分账授权申请接口
//
// 可能返回的错误码:
//  FAILEDOPERATION_ACTIONINVALID = "FailedOperation.ActionInvalid"
//  FAILEDOPERATION_PABANKERROR = "FailedOperation.PABankError"
func (c *Client) DistributeAccreditTlinx(request *DistributeAccreditTlinxRequest) (response *DistributeAccreditTlinxResponse, err error) {
    if request == nil {
        request = NewDistributeAccreditTlinxRequest()
    }
    
    response = NewDistributeAccreditTlinxResponse()
    err = c.Send(request, response)
    return
}

func NewDistributeAddReceiverRequest() (request *DistributeAddReceiverRequest) {
    request = &DistributeAddReceiverRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "DistributeAddReceiver")
    
    
    return
}

func NewDistributeAddReceiverResponse() (response *DistributeAddReceiverResponse) {
    response = &DistributeAddReceiverResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DistributeAddReceiver
// 云支付-分账添加分账接收方接口
//
// 可能返回的错误码:
//  FAILEDOPERATION_ACTIONINVALID = "FailedOperation.ActionInvalid"
//  FAILEDOPERATION_PABANKERROR = "FailedOperation.PABankError"
func (c *Client) DistributeAddReceiver(request *DistributeAddReceiverRequest) (response *DistributeAddReceiverResponse, err error) {
    if request == nil {
        request = NewDistributeAddReceiverRequest()
    }
    
    response = NewDistributeAddReceiverResponse()
    err = c.Send(request, response)
    return
}

func NewDistributeApplyRequest() (request *DistributeApplyRequest) {
    request = &DistributeApplyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "DistributeApply")
    
    
    return
}

func NewDistributeApplyResponse() (response *DistributeApplyResponse) {
    response = &DistributeApplyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DistributeApply
// 云支付-分账请求接口
//
// 可能返回的错误码:
//  FAILEDOPERATION_ACTIONINVALID = "FailedOperation.ActionInvalid"
//  FAILEDOPERATION_PABANKERROR = "FailedOperation.PABankError"
func (c *Client) DistributeApply(request *DistributeApplyRequest) (response *DistributeApplyResponse, err error) {
    if request == nil {
        request = NewDistributeApplyRequest()
    }
    
    response = NewDistributeApplyResponse()
    err = c.Send(request, response)
    return
}

func NewDistributeCancelRequest() (request *DistributeCancelRequest) {
    request = &DistributeCancelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "DistributeCancel")
    
    
    return
}

func NewDistributeCancelResponse() (response *DistributeCancelResponse) {
    response = &DistributeCancelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DistributeCancel
// 云支付-分账撤销接口
//
// 可能返回的错误码:
//  FAILEDOPERATION_ACTIONINVALID = "FailedOperation.ActionInvalid"
//  FAILEDOPERATION_PABANKERROR = "FailedOperation.PABankError"
func (c *Client) DistributeCancel(request *DistributeCancelRequest) (response *DistributeCancelResponse, err error) {
    if request == nil {
        request = NewDistributeCancelRequest()
    }
    
    response = NewDistributeCancelResponse()
    err = c.Send(request, response)
    return
}

func NewDistributeQueryRequest() (request *DistributeQueryRequest) {
    request = &DistributeQueryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "DistributeQuery")
    
    
    return
}

func NewDistributeQueryResponse() (response *DistributeQueryResponse) {
    response = &DistributeQueryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DistributeQuery
// 云支付-分账结果查询接口
//
// 可能返回的错误码:
//  FAILEDOPERATION_ACTIONINVALID = "FailedOperation.ActionInvalid"
//  FAILEDOPERATION_PABANKERROR = "FailedOperation.PABankError"
func (c *Client) DistributeQuery(request *DistributeQueryRequest) (response *DistributeQueryResponse, err error) {
    if request == nil {
        request = NewDistributeQueryRequest()
    }
    
    response = NewDistributeQueryResponse()
    err = c.Send(request, response)
    return
}

func NewDistributeQueryReceiverRequest() (request *DistributeQueryReceiverRequest) {
    request = &DistributeQueryReceiverRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "DistributeQueryReceiver")
    
    
    return
}

func NewDistributeQueryReceiverResponse() (response *DistributeQueryReceiverResponse) {
    response = &DistributeQueryReceiverResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DistributeQueryReceiver
// 云支付-T查询已添加分账接收方接口
//
// 可能返回的错误码:
//  FAILEDOPERATION_ACTIONINVALID = "FailedOperation.ActionInvalid"
//  FAILEDOPERATION_PABANKERROR = "FailedOperation.PABankError"
func (c *Client) DistributeQueryReceiver(request *DistributeQueryReceiverRequest) (response *DistributeQueryReceiverResponse, err error) {
    if request == nil {
        request = NewDistributeQueryReceiverRequest()
    }
    
    response = NewDistributeQueryReceiverResponse()
    err = c.Send(request, response)
    return
}

func NewDistributeRemoveReceiverRequest() (request *DistributeRemoveReceiverRequest) {
    request = &DistributeRemoveReceiverRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "DistributeRemoveReceiver")
    
    
    return
}

func NewDistributeRemoveReceiverResponse() (response *DistributeRemoveReceiverResponse) {
    response = &DistributeRemoveReceiverResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DistributeRemoveReceiver
// 云支付-分账解除分账接收方接口
//
// 可能返回的错误码:
//  FAILEDOPERATION_ACTIONINVALID = "FailedOperation.ActionInvalid"
//  FAILEDOPERATION_PABANKERROR = "FailedOperation.PABankError"
func (c *Client) DistributeRemoveReceiver(request *DistributeRemoveReceiverRequest) (response *DistributeRemoveReceiverResponse, err error) {
    if request == nil {
        request = NewDistributeRemoveReceiverRequest()
    }
    
    response = NewDistributeRemoveReceiverResponse()
    err = c.Send(request, response)
    return
}

func NewDownloadBillRequest() (request *DownloadBillRequest) {
    request = &DownloadBillRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "DownloadBill")
    
    
    return
}

func NewDownloadBillResponse() (response *DownloadBillResponse) {
    response = &DownloadBillResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DownloadBill
// 账单下载接口，根据本接口返回的URL地址，在D+1日下载对账单。注意，本接口返回的URL地址有时效，请尽快下载。URL超时时效后，请重新调用本接口再次获取。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CONFIGERROR = "FailedOperation.ConfigError"
//  FAILEDOPERATION_DBCONFIGERROR = "FailedOperation.DBConfigError"
//  FAILEDOPERATION_FILENOTEXIST = "FailedOperation.FileNotExist"
//  FAILEDOPERATION_SERVICEERROR = "FailedOperation.ServiceError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DownloadBill(request *DownloadBillRequest) (response *DownloadBillResponse, err error) {
    if request == nil {
        request = NewDownloadBillRequest()
    }
    
    response = NewDownloadBillResponse()
    err = c.Send(request, response)
    return
}

func NewDownloadOrgFileRequest() (request *DownloadOrgFileRequest) {
    request = &DownloadOrgFileRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "DownloadOrgFile")
    
    
    return
}

func NewDownloadOrgFileResponse() (response *DownloadOrgFileResponse) {
    response = &DownloadOrgFileResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DownloadOrgFile
// 云支付-下载机构文件接口
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CONFIGERROR = "FailedOperation.ConfigError"
//  FAILEDOPERATION_DBCONFIGERROR = "FailedOperation.DBConfigError"
//  FAILEDOPERATION_FILENOTEXIST = "FailedOperation.FileNotExist"
//  FAILEDOPERATION_SERVICEERROR = "FailedOperation.ServiceError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DownloadOrgFile(request *DownloadOrgFileRequest) (response *DownloadOrgFileResponse, err error) {
    if request == nil {
        request = NewDownloadOrgFileRequest()
    }
    
    response = NewDownloadOrgFileResponse()
    err = c.Send(request, response)
    return
}

func NewDownloadReconciliationUrlRequest() (request *DownloadReconciliationUrlRequest) {
    request = &DownloadReconciliationUrlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "DownloadReconciliationUrl")
    
    
    return
}

func NewDownloadReconciliationUrlResponse() (response *DownloadReconciliationUrlResponse) {
    response = &DownloadReconciliationUrlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DownloadReconciliationUrl
// 获取对账中心账单下载地址的接口
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDPARAMETER = "FailedOperation.InvalidParameter"
//  FAILEDOPERATION_ISEMPTY = "FailedOperation.IsEmpty"
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
func (c *Client) DownloadReconciliationUrl(request *DownloadReconciliationUrlRequest) (response *DownloadReconciliationUrlResponse, err error) {
    if request == nil {
        request = NewDownloadReconciliationUrlRequest()
    }
    
    response = NewDownloadReconciliationUrlResponse()
    err = c.Send(request, response)
    return
}

func NewExecuteMemberTransactionRequest() (request *ExecuteMemberTransactionRequest) {
    request = &ExecuteMemberTransactionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "ExecuteMemberTransaction")
    
    
    return
}

func NewExecuteMemberTransactionResponse() (response *ExecuteMemberTransactionResponse) {
    response = &ExecuteMemberTransactionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ExecuteMemberTransaction
// 会员间交易接口
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ACTIONINVALID = "FailedOperation.ActionInvalid"
//  FAILEDOPERATION_PABANKERROR = "FailedOperation.PABankError"
func (c *Client) ExecuteMemberTransaction(request *ExecuteMemberTransactionRequest) (response *ExecuteMemberTransactionResponse, err error) {
    if request == nil {
        request = NewExecuteMemberTransactionRequest()
    }
    
    response = NewExecuteMemberTransactionResponse()
    err = c.Send(request, response)
    return
}

func NewMigrateOrderRefundRequest() (request *MigrateOrderRefundRequest) {
    request = &MigrateOrderRefundRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "MigrateOrderRefund")
    
    
    return
}

func NewMigrateOrderRefundResponse() (response *MigrateOrderRefundResponse) {
    response = &MigrateOrderRefundResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// MigrateOrderRefund
// 山姆聚合支付项目-存量订单退款接口。可以通过本接口将支付款全部或部分退还给付款方，在收到用户退款请求并且验证成功之后，按照退款规则将支付款按原路退回到支付帐号。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CHANNELDENY = "FailedOperation.ChannelDeny"
//  FAILEDOPERATION_CHANNELERROR = "FailedOperation.ChannelError"
//  FAILEDOPERATION_NORECORD = "FailedOperation.NoRecord"
//  FAILEDOPERATION_OCREPEATORDER = "FailedOperation.OcRepeatOrder"
//  FAILEDOPERATION_SERVICEERROR = "FailedOperation.ServiceError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
//  INTERNALERROR_BACKENDERROR = "InternalError.BackendError"
//  INTERNALERROR_DBACCESSERROR = "InternalError.DBAccessError"
//  INTERNALERROR_DELETEDBERROR = "InternalError.DeleteDBError"
//  INTERNALERROR_DUPLICATEKEYERROR = "InternalError.DuplicateKeyError"
//  INTERNALERROR_PARAMETERERROR = "InternalError.ParameterError"
//  INTERNALERROR_SANDBOXACCESSERROR = "InternalError.SandBoxAccessError"
//  INTERNALERROR_SAVEDBERROR = "InternalError.SaveDBError"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  INTERNALERROR_UNKOWNERROR = "InternalError.UnkownError"
//  INVALIDPARAMETER_LACKPARAMETER = "InvalidParameter.LackParameter"
//  INVALIDPARAMETER_UNSUPPORTEDPARAMETER = "InvalidParameter.UnsupportedParameter"
//  RESOURCENOTFOUND_MERCHANTINFONOTFOUND = "ResourceNotFound.MerchantInfoNotFound"
func (c *Client) MigrateOrderRefund(request *MigrateOrderRefundRequest) (response *MigrateOrderRefundResponse, err error) {
    if request == nil {
        request = NewMigrateOrderRefundRequest()
    }
    
    response = NewMigrateOrderRefundResponse()
    err = c.Send(request, response)
    return
}

func NewMigrateOrderRefundQueryRequest() (request *MigrateOrderRefundQueryRequest) {
    request = &MigrateOrderRefundQueryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "MigrateOrderRefundQuery")
    
    
    return
}

func NewMigrateOrderRefundQueryResponse() (response *MigrateOrderRefundQueryResponse) {
    response = &MigrateOrderRefundQueryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// MigrateOrderRefundQuery
// 提交退款申请后，通过调用该接口查询退款状态。退款可能有一定延时。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CHANNELDENY = "FailedOperation.ChannelDeny"
//  FAILEDOPERATION_CHANNELERROR = "FailedOperation.ChannelError"
//  FAILEDOPERATION_NORECORD = "FailedOperation.NoRecord"
//  FAILEDOPERATION_OCREPEATORDER = "FailedOperation.OcRepeatOrder"
//  FAILEDOPERATION_SERVICEERROR = "FailedOperation.ServiceError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
//  INTERNALERROR_BACKENDERROR = "InternalError.BackendError"
//  INTERNALERROR_DBACCESSERROR = "InternalError.DBAccessError"
//  INTERNALERROR_PARAMETERERROR = "InternalError.ParameterError"
//  INTERNALERROR_SANDBOXACCESSERROR = "InternalError.SandBoxAccessError"
//  INTERNALERROR_SAVEDBERROR = "InternalError.SaveDBError"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  INTERNALERROR_UNKOWNERROR = "InternalError.UnkownError"
//  INVALIDPARAMETER_LACKPARAMETER = "InvalidParameter.LackParameter"
//  INVALIDPARAMETER_UNSUPPORTEDPARAMETER = "InvalidParameter.UnsupportedParameter"
//  RESOURCENOTFOUND_MERCHANTINFONOTFOUND = "ResourceNotFound.MerchantInfoNotFound"
func (c *Client) MigrateOrderRefundQuery(request *MigrateOrderRefundQueryRequest) (response *MigrateOrderRefundQueryResponse, err error) {
    if request == nil {
        request = NewMigrateOrderRefundQueryRequest()
    }
    
    response = NewMigrateOrderRefundQueryResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAgentTaxPaymentInfoRequest() (request *ModifyAgentTaxPaymentInfoRequest) {
    request = &ModifyAgentTaxPaymentInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "ModifyAgentTaxPaymentInfo")
    
    
    return
}

func NewModifyAgentTaxPaymentInfoResponse() (response *ModifyAgentTaxPaymentInfoResponse) {
    response = &ModifyAgentTaxPaymentInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyAgentTaxPaymentInfo
// 直播平台-修改代理商完税信息
//
// 可能返回的错误码:
//  INTERNALERROR_DELETEDBERROR = "InternalError.DeleteDBError"
//  INTERNALERROR_SAVEDBERROR = "InternalError.SaveDBError"
//  INTERNALERROR_UNKOWNERROR = "InternalError.UnkownError"
//  INVALIDPARAMETER_LACKPARAMETER = "InvalidParameter.LackParameter"
//  RESOURCENOTFOUND_BATCHINFONOTFOUND = "ResourceNotFound.BatchInfoNotFound"
//  RESOURCENOTFOUND_PLATFORMINFONOTFOUND = "ResourceNotFound.PlatformInfoNotFound"
func (c *Client) ModifyAgentTaxPaymentInfo(request *ModifyAgentTaxPaymentInfoRequest) (response *ModifyAgentTaxPaymentInfoResponse, err error) {
    if request == nil {
        request = NewModifyAgentTaxPaymentInfoRequest()
    }
    
    response = NewModifyAgentTaxPaymentInfoResponse()
    err = c.Send(request, response)
    return
}

func NewModifyBindedAccountRequest() (request *ModifyBindedAccountRequest) {
    request = &ModifyBindedAccountRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "ModifyBindedAccount")
    
    
    return
}

func NewModifyBindedAccountResponse() (response *ModifyBindedAccountResponse) {
    response = &ModifyBindedAccountResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyBindedAccount
// 灵云-重新绑定账号
//
// 可能返回的错误码:
//  INTERNALERROR_DELETEDBERROR = "InternalError.DeleteDBError"
//  INTERNALERROR_SAVEDBERROR = "InternalError.SaveDBError"
//  INTERNALERROR_UNKOWNERROR = "InternalError.UnkownError"
//  INVALIDPARAMETER_LACKPARAMETER = "InvalidParameter.LackParameter"
//  RESOURCENOTFOUND_BATCHINFONOTFOUND = "ResourceNotFound.BatchInfoNotFound"
//  RESOURCENOTFOUND_PLATFORMINFONOTFOUND = "ResourceNotFound.PlatformInfoNotFound"
func (c *Client) ModifyBindedAccount(request *ModifyBindedAccountRequest) (response *ModifyBindedAccountResponse, err error) {
    if request == nil {
        request = NewModifyBindedAccountRequest()
    }
    
    response = NewModifyBindedAccountResponse()
    err = c.Send(request, response)
    return
}

func NewModifyMerchantRequest() (request *ModifyMerchantRequest) {
    request = &ModifyMerchantRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "ModifyMerchant")
    
    
    return
}

func NewModifyMerchantResponse() (response *ModifyMerchantResponse) {
    response = &ModifyMerchantResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyMerchant
// 云鉴-商户信息修改的接口
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDPARAMETER = "FailedOperation.InvalidParameter"
//  FAILEDOPERATION_MERCHANTNOTEXIST = "FailedOperation.MerchantNotExist"
//  FAILEDOPERATION_MODIFYMERCHANTFAILED = "FailedOperation.ModifyMerchantFailed"
func (c *Client) ModifyMerchant(request *ModifyMerchantRequest) (response *ModifyMerchantResponse, err error) {
    if request == nil {
        request = NewModifyMerchantRequest()
    }
    
    response = NewModifyMerchantResponse()
    err = c.Send(request, response)
    return
}

func NewModifyMntMbrBindRelateAcctBankCodeRequest() (request *ModifyMntMbrBindRelateAcctBankCodeRequest) {
    request = &ModifyMntMbrBindRelateAcctBankCodeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "ModifyMntMbrBindRelateAcctBankCode")
    
    
    return
}

func NewModifyMntMbrBindRelateAcctBankCodeResponse() (response *ModifyMntMbrBindRelateAcctBankCodeResponse) {
    response = &ModifyMntMbrBindRelateAcctBankCodeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyMntMbrBindRelateAcctBankCode
// 维护会员绑定提现账户联行号。此接口可以支持市场修改会员的提现账户的开户行信息，具体包括开户行行名、开户行的银行联行号（大小额联行号）和超级网银行号。
//
// 可能返回的错误码:
//  FAILEDOPERATION_BACKENDERROR = "FailedOperation.BackendError"
//  INTERNALERROR_BACKENDERROR = "InternalError.BackendError"
//  INTERNALERROR_DBACCESSERROR = "InternalError.DBAccessError"
//  INTERNALERROR_FUNDSUMMARYACCTNOINCONSISTENTERROR = "InternalError.FundSummaryAcctNoInconsistentError"
//  INTERNALERROR_PARAMETERERROR = "InternalError.ParameterError"
//  INTERNALERROR_SAVEDBERROR = "InternalError.SaveDBError"
//  INTERNALERROR_SUBACCOUNTNOTFOUNDERROR = "InternalError.SubAccountNotFoundError"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  INTERNALERROR_UNKOWNERROR = "InternalError.UnkownError"
func (c *Client) ModifyMntMbrBindRelateAcctBankCode(request *ModifyMntMbrBindRelateAcctBankCodeRequest) (response *ModifyMntMbrBindRelateAcctBankCodeResponse, err error) {
    if request == nil {
        request = NewModifyMntMbrBindRelateAcctBankCodeRequest()
    }
    
    response = NewModifyMntMbrBindRelateAcctBankCodeResponse()
    err = c.Send(request, response)
    return
}

func NewQueryAcctBindingRequest() (request *QueryAcctBindingRequest) {
    request = &QueryAcctBindingRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "QueryAcctBinding")
    
    
    return
}

func NewQueryAcctBindingResponse() (response *QueryAcctBindingResponse) {
    response = &QueryAcctBindingResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// QueryAcctBinding
// 聚鑫-查询子账户绑定银行卡
//
// 可能返回的错误码:
//  FAILEDOPERATION_BANKFAILED = "FailedOperation.BankFailed"
//  FAILEDOPERATION_SERVICEERROR = "FailedOperation.ServiceError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER_ = "MissingParameter."
func (c *Client) QueryAcctBinding(request *QueryAcctBindingRequest) (response *QueryAcctBindingResponse, err error) {
    if request == nil {
        request = NewQueryAcctBindingRequest()
    }
    
    response = NewQueryAcctBindingResponse()
    err = c.Send(request, response)
    return
}

func NewQueryAcctInfoRequest() (request *QueryAcctInfoRequest) {
    request = &QueryAcctInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "QueryAcctInfo")
    
    
    return
}

func NewQueryAcctInfoResponse() (response *QueryAcctInfoResponse) {
    response = &QueryAcctInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// QueryAcctInfo
// 聚鑫-开户信息查询
//
// 可能返回的错误码:
//  FAILEDOPERATION_BANKFAILED = "FailedOperation.BankFailed"
//  FAILEDOPERATION_SERVICEERROR = "FailedOperation.ServiceError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER_ = "MissingParameter."
func (c *Client) QueryAcctInfo(request *QueryAcctInfoRequest) (response *QueryAcctInfoResponse, err error) {
    if request == nil {
        request = NewQueryAcctInfoRequest()
    }
    
    response = NewQueryAcctInfoResponse()
    err = c.Send(request, response)
    return
}

func NewQueryAcctInfoListRequest() (request *QueryAcctInfoListRequest) {
    request = &QueryAcctInfoListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "QueryAcctInfoList")
    
    
    return
}

func NewQueryAcctInfoListResponse() (response *QueryAcctInfoListResponse) {
    response = &QueryAcctInfoListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// QueryAcctInfoList
// 聚鑫-开户信息列表查询, 查询某一段时间的开户信息
//
// 可能返回的错误码:
//  FAILEDOPERATION_BANKFAILED = "FailedOperation.BankFailed"
//  FAILEDOPERATION_SERVICEERROR = "FailedOperation.ServiceError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER_ = "MissingParameter."
func (c *Client) QueryAcctInfoList(request *QueryAcctInfoListRequest) (response *QueryAcctInfoListResponse, err error) {
    if request == nil {
        request = NewQueryAcctInfoListRequest()
    }
    
    response = NewQueryAcctInfoListResponse()
    err = c.Send(request, response)
    return
}

func NewQueryAgentStatementsRequest() (request *QueryAgentStatementsRequest) {
    request = &QueryAgentStatementsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "QueryAgentStatements")
    
    
    return
}

func NewQueryAgentStatementsResponse() (response *QueryAgentStatementsResponse) {
    response = &QueryAgentStatementsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// QueryAgentStatements
// 直播平台-查询代理商结算单链接
//
// 可能返回的错误码:
//  AUTHFAILURE_ACCOUNT = "AuthFailure.Account"
//  FAILEDOPERATION_GETLIVEDAILYSUMMARY = "FailedOperation.GetLiveDailySummary"
//  FAILEDOPERATION_QUERYAGENTSTATEMENTS = "FailedOperation.QueryAgentStatements"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  MISSINGPARAMETER_ACTION = "MissingParameter.Action"
//  MISSINGPARAMETER_APPID = "MissingParameter.AppId"
//  RESOURCENOTFOUND_ACCOUNT = "ResourceNotFound.Account"
func (c *Client) QueryAgentStatements(request *QueryAgentStatementsRequest) (response *QueryAgentStatementsResponse, err error) {
    if request == nil {
        request = NewQueryAgentStatementsRequest()
    }
    
    response = NewQueryAgentStatementsResponse()
    err = c.Send(request, response)
    return
}

func NewQueryAgentTaxPaymentBatchRequest() (request *QueryAgentTaxPaymentBatchRequest) {
    request = &QueryAgentTaxPaymentBatchRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "QueryAgentTaxPaymentBatch")
    
    
    return
}

func NewQueryAgentTaxPaymentBatchResponse() (response *QueryAgentTaxPaymentBatchResponse) {
    response = &QueryAgentTaxPaymentBatchResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// QueryAgentTaxPaymentBatch
// 直播平台-查询批次信息
//
// 可能返回的错误码:
//  INTERNALERROR_DELETEDBERROR = "InternalError.DeleteDBError"
//  INTERNALERROR_SAVEDBERROR = "InternalError.SaveDBError"
//  INTERNALERROR_UNKOWNERROR = "InternalError.UnkownError"
//  INVALIDPARAMETER_LACKPARAMETER = "InvalidParameter.LackParameter"
//  RESOURCENOTFOUND_BATCHINFONOTFOUND = "ResourceNotFound.BatchInfoNotFound"
//  RESOURCENOTFOUND_PLATFORMINFONOTFOUND = "ResourceNotFound.PlatformInfoNotFound"
func (c *Client) QueryAgentTaxPaymentBatch(request *QueryAgentTaxPaymentBatchRequest) (response *QueryAgentTaxPaymentBatchResponse, err error) {
    if request == nil {
        request = NewQueryAgentTaxPaymentBatchRequest()
    }
    
    response = NewQueryAgentTaxPaymentBatchResponse()
    err = c.Send(request, response)
    return
}

func NewQueryAnchorContractInfoRequest() (request *QueryAnchorContractInfoRequest) {
    request = &QueryAnchorContractInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "QueryAnchorContractInfo")
    
    
    return
}

func NewQueryAnchorContractInfoResponse() (response *QueryAnchorContractInfoResponse) {
    response = &QueryAnchorContractInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// QueryAnchorContractInfo
// 直播平台-查询主播签约信息
//
// 可能返回的错误码:
//  INTERNALERROR_DELETEDBERROR = "InternalError.DeleteDBError"
//  INTERNALERROR_UNKOWNERROR = "InternalError.UnkownError"
//  INVALIDPARAMETER_LACKPARAMETER = "InvalidParameter.LackParameter"
//  RESOURCENOTFOUND_BATCHINFONOTFOUND = "ResourceNotFound.BatchInfoNotFound"
//  RESOURCENOTFOUND_PLATFORMINFONOTFOUND = "ResourceNotFound.PlatformInfoNotFound"
func (c *Client) QueryAnchorContractInfo(request *QueryAnchorContractInfoRequest) (response *QueryAnchorContractInfoResponse, err error) {
    if request == nil {
        request = NewQueryAnchorContractInfoRequest()
    }
    
    response = NewQueryAnchorContractInfoResponse()
    err = c.Send(request, response)
    return
}

func NewQueryApplicationMaterialRequest() (request *QueryApplicationMaterialRequest) {
    request = &QueryApplicationMaterialRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "QueryApplicationMaterial")
    
    
    return
}

func NewQueryApplicationMaterialResponse() (response *QueryApplicationMaterialResponse) {
    response = &QueryApplicationMaterialResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// QueryApplicationMaterial
// 跨境-成功申报材料查询。查询成功入库的申报材料。
//
// 可能返回的错误码:
//  INTERNALERROR_PARAMETERERROR = "InternalError.ParameterError"
//  INTERNALERROR_UNKOWNERROR = "InternalError.UnkownError"
//  RESOURCENOTFOUND_MERCHANTINFONOTFOUND = "ResourceNotFound.MerchantInfoNotFound"
func (c *Client) QueryApplicationMaterial(request *QueryApplicationMaterialRequest) (response *QueryApplicationMaterialResponse, err error) {
    if request == nil {
        request = NewQueryApplicationMaterialRequest()
    }
    
    response = NewQueryApplicationMaterialResponse()
    err = c.Send(request, response)
    return
}

func NewQueryAssignmentRequest() (request *QueryAssignmentRequest) {
    request = &QueryAssignmentRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "QueryAssignment")
    
    
    return
}

func NewQueryAssignmentResponse() (response *QueryAssignmentResponse) {
    response = &QueryAssignmentResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// QueryAssignment
// 直播平台-查询分配关系
//
// 可能返回的错误码:
//  INTERNALERROR_PARAMETERERROR = "InternalError.ParameterError"
//  INTERNALERROR_UNKOWNERROR = "InternalError.UnkownError"
//  RESOURCENOTFOUND_MERCHANTINFONOTFOUND = "ResourceNotFound.MerchantInfoNotFound"
func (c *Client) QueryAssignment(request *QueryAssignmentRequest) (response *QueryAssignmentResponse, err error) {
    if request == nil {
        request = NewQueryAssignmentRequest()
    }
    
    response = NewQueryAssignmentResponse()
    err = c.Send(request, response)
    return
}

func NewQueryBalanceRequest() (request *QueryBalanceRequest) {
    request = &QueryBalanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "QueryBalance")
    
    
    return
}

func NewQueryBalanceResponse() (response *QueryBalanceResponse) {
    response = &QueryBalanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// QueryBalance
// 子商户余额查询
//
// 可能返回的错误码:
//  FAILEDOPERATION_BANKFAILED = "FailedOperation.BankFailed"
//  FAILEDOPERATION_SERVICEERROR = "FailedOperation.ServiceError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER_ = "MissingParameter."
func (c *Client) QueryBalance(request *QueryBalanceRequest) (response *QueryBalanceResponse, err error) {
    if request == nil {
        request = NewQueryBalanceRequest()
    }
    
    response = NewQueryBalanceResponse()
    err = c.Send(request, response)
    return
}

func NewQueryBankClearRequest() (request *QueryBankClearRequest) {
    request = &QueryBankClearRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "QueryBankClear")
    
    
    return
}

func NewQueryBankClearResponse() (response *QueryBankClearResponse) {
    response = &QueryBankClearResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// QueryBankClear
// 查询银行在途清算结果。查询时间段内交易网的在途清算结果。
//
// 可能返回的错误码:
//  FAILEDOPERATION_BACKENDERROR = "FailedOperation.BackendError"
//  INTERNALERROR_BACKENDERROR = "InternalError.BackendError"
//  INTERNALERROR_DBACCESSERROR = "InternalError.DBAccessError"
//  INTERNALERROR_FUNDSUMMARYACCTNOINCONSISTENTERROR = "InternalError.FundSummaryAcctNoInconsistentError"
//  INTERNALERROR_PARAMETERERROR = "InternalError.ParameterError"
//  INTERNALERROR_SAVEDBERROR = "InternalError.SaveDBError"
//  INTERNALERROR_SUBACCOUNTNOTFOUNDERROR = "InternalError.SubAccountNotFoundError"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  INTERNALERROR_UNKOWNERROR = "InternalError.UnkownError"
func (c *Client) QueryBankClear(request *QueryBankClearRequest) (response *QueryBankClearResponse, err error) {
    if request == nil {
        request = NewQueryBankClearRequest()
    }
    
    response = NewQueryBankClearResponse()
    err = c.Send(request, response)
    return
}

func NewQueryBankTransactionDetailsRequest() (request *QueryBankTransactionDetailsRequest) {
    request = &QueryBankTransactionDetailsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "QueryBankTransactionDetails")
    
    
    return
}

func NewQueryBankTransactionDetailsResponse() (response *QueryBankTransactionDetailsResponse) {
    response = &QueryBankTransactionDetailsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// QueryBankTransactionDetails
// 查询银行时间段内交易明细。查询时间段的会员成功交易。
//
// 可能返回的错误码:
//  FAILEDOPERATION_BACKENDERROR = "FailedOperation.BackendError"
//  INTERNALERROR_BACKENDERROR = "InternalError.BackendError"
//  INTERNALERROR_DBACCESSERROR = "InternalError.DBAccessError"
//  INTERNALERROR_FUNDSUMMARYACCTNOINCONSISTENTERROR = "InternalError.FundSummaryAcctNoInconsistentError"
//  INTERNALERROR_PARAMETERERROR = "InternalError.ParameterError"
//  INTERNALERROR_SAVEDBERROR = "InternalError.SaveDBError"
//  INTERNALERROR_SUBACCOUNTNOTFOUNDERROR = "InternalError.SubAccountNotFoundError"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  INTERNALERROR_UNKOWNERROR = "InternalError.UnkownError"
func (c *Client) QueryBankTransactionDetails(request *QueryBankTransactionDetailsRequest) (response *QueryBankTransactionDetailsResponse, err error) {
    if request == nil {
        request = NewQueryBankTransactionDetailsRequest()
    }
    
    response = NewQueryBankTransactionDetailsResponse()
    err = c.Send(request, response)
    return
}

func NewQueryBankWithdrawCashDetailsRequest() (request *QueryBankWithdrawCashDetailsRequest) {
    request = &QueryBankWithdrawCashDetailsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "QueryBankWithdrawCashDetails")
    
    
    return
}

func NewQueryBankWithdrawCashDetailsResponse() (response *QueryBankWithdrawCashDetailsResponse) {
    response = &QueryBankWithdrawCashDetailsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// QueryBankWithdrawCashDetails
// 查询银行时间段内清分提现明细。查询银行时间段内清分提现明细接口：若为“见证+收单退款”“见证+收单充值”记录时备注Note为“见证+收单充值,订单号”“见证+收单退款,订单号”，此接口可以查到T0/T1的充值明细和退款记录。查询标志：充值记录仍用3清分选项查询，退款记录同提现用2选项查询。
//
// 可能返回的错误码:
//  FAILEDOPERATION_APPDENY = "FailedOperation.AppDeny"
//  INTERNALERROR_BACKENDERROR = "InternalError.BackendError"
//  INTERNALERROR_DBACCESSERROR = "InternalError.DBAccessError"
//  INTERNALERROR_FUNDSUMMARYACCTNOINCONSISTENTERROR = "InternalError.FundSummaryAcctNoInconsistentError"
//  INTERNALERROR_PARAMETERERROR = "InternalError.ParameterError"
//  INTERNALERROR_SAVEDBERROR = "InternalError.SaveDBError"
//  INTERNALERROR_SUBACCOUNTNOTFOUNDERROR = "InternalError.SubAccountNotFoundError"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  INTERNALERROR_UNKOWNERROR = "InternalError.UnkownError"
func (c *Client) QueryBankWithdrawCashDetails(request *QueryBankWithdrawCashDetailsRequest) (response *QueryBankWithdrawCashDetailsResponse, err error) {
    if request == nil {
        request = NewQueryBankWithdrawCashDetailsRequest()
    }
    
    response = NewQueryBankWithdrawCashDetailsResponse()
    err = c.Send(request, response)
    return
}

func NewQueryBatchPaymentResultRequest() (request *QueryBatchPaymentResultRequest) {
    request = &QueryBatchPaymentResultRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "QueryBatchPaymentResult")
    
    
    return
}

func NewQueryBatchPaymentResultResponse() (response *QueryBatchPaymentResultResponse) {
    response = &QueryBatchPaymentResultResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// QueryBatchPaymentResult
// 灵云-批量转账结果查询
//
// 可能返回的错误码:
//  FAILEDOPERATION_APPDENY = "FailedOperation.AppDeny"
//  INTERNALERROR_BACKENDERROR = "InternalError.BackendError"
//  INTERNALERROR_DBACCESSERROR = "InternalError.DBAccessError"
//  INTERNALERROR_FUNDSUMMARYACCTNOINCONSISTENTERROR = "InternalError.FundSummaryAcctNoInconsistentError"
//  INTERNALERROR_PARAMETERERROR = "InternalError.ParameterError"
//  INTERNALERROR_SAVEDBERROR = "InternalError.SaveDBError"
//  INTERNALERROR_SUBACCOUNTNOTFOUNDERROR = "InternalError.SubAccountNotFoundError"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  INTERNALERROR_UNKOWNERROR = "InternalError.UnkownError"
func (c *Client) QueryBatchPaymentResult(request *QueryBatchPaymentResultRequest) (response *QueryBatchPaymentResultResponse, err error) {
    if request == nil {
        request = NewQueryBatchPaymentResultRequest()
    }
    
    response = NewQueryBatchPaymentResultResponse()
    err = c.Send(request, response)
    return
}

func NewQueryBillDownloadURLRequest() (request *QueryBillDownloadURLRequest) {
    request = &QueryBillDownloadURLRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "QueryBillDownloadURL")
    
    
    return
}

func NewQueryBillDownloadURLResponse() (response *QueryBillDownloadURLResponse) {
    response = &QueryBillDownloadURLResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// QueryBillDownloadURL
// 获取单笔代发转账对账单下载URL
//
// 可能返回的错误码:
//  FAILEDOPERATION_APPDENY = "FailedOperation.AppDeny"
//  INTERNALERROR_BACKENDERROR = "InternalError.BackendError"
//  INTERNALERROR_DBACCESSERROR = "InternalError.DBAccessError"
//  INTERNALERROR_FUNDSUMMARYACCTNOINCONSISTENTERROR = "InternalError.FundSummaryAcctNoInconsistentError"
//  INTERNALERROR_PARAMETERERROR = "InternalError.ParameterError"
//  INTERNALERROR_SAVEDBERROR = "InternalError.SaveDBError"
//  INTERNALERROR_SUBACCOUNTNOTFOUNDERROR = "InternalError.SubAccountNotFoundError"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  INTERNALERROR_UNKOWNERROR = "InternalError.UnkownError"
func (c *Client) QueryBillDownloadURL(request *QueryBillDownloadURLRequest) (response *QueryBillDownloadURLResponse, err error) {
    if request == nil {
        request = NewQueryBillDownloadURLRequest()
    }
    
    response = NewQueryBillDownloadURLResponse()
    err = c.Send(request, response)
    return
}

func NewQueryCityCodeRequest() (request *QueryCityCodeRequest) {
    request = &QueryCityCodeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "QueryCityCode")
    
    
    return
}

func NewQueryCityCodeResponse() (response *QueryCityCodeResponse) {
    response = &QueryCityCodeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// QueryCityCode
// 云支付-查询城市编码接口
//
// 可能返回的错误码:
//  FAILEDOPERATION_APPDENY = "FailedOperation.AppDeny"
//  INTERNALERROR_BACKENDERROR = "InternalError.BackendError"
//  INTERNALERROR_DBACCESSERROR = "InternalError.DBAccessError"
//  INTERNALERROR_FUNDSUMMARYACCTNOINCONSISTENTERROR = "InternalError.FundSummaryAcctNoInconsistentError"
//  INTERNALERROR_PARAMETERERROR = "InternalError.ParameterError"
//  INTERNALERROR_SAVEDBERROR = "InternalError.SaveDBError"
//  INTERNALERROR_SUBACCOUNTNOTFOUNDERROR = "InternalError.SubAccountNotFoundError"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  INTERNALERROR_UNKOWNERROR = "InternalError.UnkownError"
func (c *Client) QueryCityCode(request *QueryCityCodeRequest) (response *QueryCityCodeResponse, err error) {
    if request == nil {
        request = NewQueryCityCodeRequest()
    }
    
    response = NewQueryCityCodeResponse()
    err = c.Send(request, response)
    return
}

func NewQueryCommonTransferRechargeRequest() (request *QueryCommonTransferRechargeRequest) {
    request = &QueryCommonTransferRechargeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "QueryCommonTransferRecharge")
    
    
    return
}

func NewQueryCommonTransferRechargeResponse() (response *QueryCommonTransferRechargeResponse) {
    response = &QueryCommonTransferRechargeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// QueryCommonTransferRecharge
// 查询普通转账充值明细。接口用于查询会员主动转账进资金汇总账户的明细情况。若会员使用绑定账号转入，则直接入账到会员子账户。若未使用绑定账号转入，则系统无法自动清分到对应子账户，则转入挂账子账户由平台自行清分。若是 “见证+收单充值”T0充值记录时备注Note为“见证+收单充值,订单号” 此接口可以查到T0到账的“见证+收单充值”充值记录。
//
// 可能返回的错误码:
//  FAILEDOPERATION_BACKENDERROR = "FailedOperation.BackendError"
//  INTERNALERROR_BACKENDERROR = "InternalError.BackendError"
//  INTERNALERROR_DBACCESSERROR = "InternalError.DBAccessError"
//  INTERNALERROR_FUNDSUMMARYACCTNOINCONSISTENTERROR = "InternalError.FundSummaryAcctNoInconsistentError"
//  INTERNALERROR_PARAMETERERROR = "InternalError.ParameterError"
//  INTERNALERROR_SAVEDBERROR = "InternalError.SaveDBError"
//  INTERNALERROR_SUBACCOUNTNOTFOUNDERROR = "InternalError.SubAccountNotFoundError"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  INTERNALERROR_UNKOWNERROR = "InternalError.UnkownError"
func (c *Client) QueryCommonTransferRecharge(request *QueryCommonTransferRechargeRequest) (response *QueryCommonTransferRechargeResponse, err error) {
    if request == nil {
        request = NewQueryCommonTransferRechargeRequest()
    }
    
    response = NewQueryCommonTransferRechargeResponse()
    err = c.Send(request, response)
    return
}

func NewQueryContractRequest() (request *QueryContractRequest) {
    request = &QueryContractRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "QueryContract")
    
    
    return
}

func NewQueryContractResponse() (response *QueryContractResponse) {
    response = &QueryContractResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// QueryContract
// 通过此接口查询签约数据
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_BACKCALLERROR = "FailedOperation.BackCallError"
//  FAILEDOPERATION_CALLCHANNELGATEWAYERROR = "FailedOperation.CallChannelGatewayError"
//  FAILEDOPERATION_DBCLIENTCONNECTFAILED = "FailedOperation.DbClientConnectFailed"
//  FAILEDOPERATION_DBCLIENTQUERYFAILED = "FailedOperation.DbClientQueryFailed"
//  FAILEDOPERATION_DBCLIENTUPDATEFAILED = "FailedOperation.DbClientUpdateFailed"
//  FAILEDOPERATION_EXTERNALCONTRACTINDEXNOTFOUND = "FailedOperation.ExternalContractIndexNotFound"
//  FAILEDOPERATION_EXTERNALCONTRACTNOTFOUND = "FailedOperation.ExternalContractNotFound"
//  FAILEDOPERATION_EXTERNALMERCHANTCONTRACTINFOCONFIGNOFOUND = "FailedOperation.ExternalMerchantContractInfoConfigNoFound"
//  FAILEDOPERATION_EXTERNALMERCHANTINDEXCONFIGNOFOUND = "FailedOperation.ExternalMerchantIndexConfigNoFound"
//  FAILEDOPERATION_EXTERNALMERCHANTINFOCONFIGNOFOUND = "FailedOperation.ExternalMerchantInfoConfigNoFound"
//  FAILEDOPERATION_HTTPDOREQUESTERROR = "FailedOperation.HttpDoRequestError"
//  FAILEDOPERATION_MARSHALERROR = "FailedOperation.MarshalError"
//  FAILEDOPERATION_MERCHANTPERMISSIONERROR = "FailedOperation.MerchantPermissionError"
//  FAILEDOPERATION_NOTIFYURLPARSEERROR = "FailedOperation.NotifyUrlParseError"
//  FAILEDOPERATION_PARAMETERERROR = "FailedOperation.ParameterError"
//  FAILEDOPERATION_QUERYCONTRACTNULL = "FailedOperation.QueryContractNull"
//  FAILEDOPERATION_QUERYMCHANNELERROR = "FailedOperation.QueryMchannelError"
//  FAILEDOPERATION_QUERYMODEERROR = "FailedOperation.QueryModeError"
//  FAILEDOPERATION_QUERYRESULTNULL = "FailedOperation.QueryResultNull"
//  FAILEDOPERATION_SIGNERROR = "FailedOperation.SignError"
//  FAILEDOPERATION_UNMARSHALERROR = "FailedOperation.UnmarshalError"
//  FAILEDOPERATION_UPDATECONTRACTSTATUSFAILED = "FailedOperation.UpdateContractStatusFailed"
//  FAILEDOPERATION_XMLFAIL = "FailedOperation.XmlFail"
//  INTERNALERROR_BACKENDCONNECTIONERROR = "InternalError.BackendConnectionError"
//  INTERNALERROR_BACKENDINTERNALERROR = "InternalError.BackendInternalError"
//  INTERNALERROR_BACKENDNETWORKERROR = "InternalError.BackendNetworkError"
//  INTERNALERROR_BACKENDROUTERERROR = "InternalError.BackendRouterError"
//  INTERNALERROR_BACKENDTIMEOUT = "InternalError.BackendTimeOut"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMMARSHALFAILED = "InvalidParameter.ParamMarshalFailed"
//  INVALIDPARAMETER_PARAMUNMARSHALFAILED = "InvalidParameter.ParamUnmarshalFailed"
//  MISSINGPARAMETER_ = "MissingParameter."
func (c *Client) QueryContract(request *QueryContractRequest) (response *QueryContractResponse, err error) {
    if request == nil {
        request = NewQueryContractRequest()
    }
    
    response = NewQueryContractResponse()
    err = c.Send(request, response)
    return
}

func NewQueryContractPayFeeRequest() (request *QueryContractPayFeeRequest) {
    request = &QueryContractPayFeeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "QueryContractPayFee")
    
    
    return
}

func NewQueryContractPayFeeResponse() (response *QueryContractPayFeeResponse) {
    response = &QueryContractPayFeeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// QueryContractPayFee
// 云支付-查询支付方式费率及自定义表单项接口
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_BACKCALLERROR = "FailedOperation.BackCallError"
//  FAILEDOPERATION_CALLCHANNELGATEWAYERROR = "FailedOperation.CallChannelGatewayError"
//  FAILEDOPERATION_DBCLIENTCONNECTFAILED = "FailedOperation.DbClientConnectFailed"
//  FAILEDOPERATION_DBCLIENTQUERYFAILED = "FailedOperation.DbClientQueryFailed"
//  FAILEDOPERATION_DBCLIENTUPDATEFAILED = "FailedOperation.DbClientUpdateFailed"
//  FAILEDOPERATION_EXTERNALCONTRACTINDEXNOTFOUND = "FailedOperation.ExternalContractIndexNotFound"
//  FAILEDOPERATION_EXTERNALCONTRACTNOTFOUND = "FailedOperation.ExternalContractNotFound"
//  FAILEDOPERATION_EXTERNALMERCHANTCONTRACTINFOCONFIGNOFOUND = "FailedOperation.ExternalMerchantContractInfoConfigNoFound"
//  FAILEDOPERATION_EXTERNALMERCHANTINDEXCONFIGNOFOUND = "FailedOperation.ExternalMerchantIndexConfigNoFound"
//  FAILEDOPERATION_EXTERNALMERCHANTINFOCONFIGNOFOUND = "FailedOperation.ExternalMerchantInfoConfigNoFound"
//  FAILEDOPERATION_HTTPDOREQUESTERROR = "FailedOperation.HttpDoRequestError"
//  FAILEDOPERATION_MARSHALERROR = "FailedOperation.MarshalError"
//  FAILEDOPERATION_MERCHANTPERMISSIONERROR = "FailedOperation.MerchantPermissionError"
//  FAILEDOPERATION_NOTIFYURLPARSEERROR = "FailedOperation.NotifyUrlParseError"
//  FAILEDOPERATION_PARAMETERERROR = "FailedOperation.ParameterError"
//  FAILEDOPERATION_QUERYCONTRACTNULL = "FailedOperation.QueryContractNull"
//  FAILEDOPERATION_QUERYMCHANNELERROR = "FailedOperation.QueryMchannelError"
//  FAILEDOPERATION_QUERYMODEERROR = "FailedOperation.QueryModeError"
//  FAILEDOPERATION_QUERYRESULTNULL = "FailedOperation.QueryResultNull"
//  FAILEDOPERATION_SIGNERROR = "FailedOperation.SignError"
//  FAILEDOPERATION_UNMARSHALERROR = "FailedOperation.UnmarshalError"
//  FAILEDOPERATION_UPDATECONTRACTSTATUSFAILED = "FailedOperation.UpdateContractStatusFailed"
//  FAILEDOPERATION_XMLFAIL = "FailedOperation.XmlFail"
//  INTERNALERROR_BACKENDCONNECTIONERROR = "InternalError.BackendConnectionError"
//  INTERNALERROR_BACKENDINTERNALERROR = "InternalError.BackendInternalError"
//  INTERNALERROR_BACKENDNETWORKERROR = "InternalError.BackendNetworkError"
//  INTERNALERROR_BACKENDROUTERERROR = "InternalError.BackendRouterError"
//  INTERNALERROR_BACKENDTIMEOUT = "InternalError.BackendTimeOut"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMMARSHALFAILED = "InvalidParameter.ParamMarshalFailed"
//  INVALIDPARAMETER_PARAMUNMARSHALFAILED = "InvalidParameter.ParamUnmarshalFailed"
//  MISSINGPARAMETER_ = "MissingParameter."
func (c *Client) QueryContractPayFee(request *QueryContractPayFeeRequest) (response *QueryContractPayFeeResponse, err error) {
    if request == nil {
        request = NewQueryContractPayFeeRequest()
    }
    
    response = NewQueryContractPayFeeResponse()
    err = c.Send(request, response)
    return
}

func NewQueryContractPayWayListRequest() (request *QueryContractPayWayListRequest) {
    request = &QueryContractPayWayListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "QueryContractPayWayList")
    
    
    return
}

func NewQueryContractPayWayListResponse() (response *QueryContractPayWayListResponse) {
    response = &QueryContractPayWayListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// QueryContractPayWayList
// 云支付-查询合同支付方式列表接口
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_BACKCALLERROR = "FailedOperation.BackCallError"
//  FAILEDOPERATION_CALLCHANNELGATEWAYERROR = "FailedOperation.CallChannelGatewayError"
//  FAILEDOPERATION_DBCLIENTCONNECTFAILED = "FailedOperation.DbClientConnectFailed"
//  FAILEDOPERATION_DBCLIENTQUERYFAILED = "FailedOperation.DbClientQueryFailed"
//  FAILEDOPERATION_DBCLIENTUPDATEFAILED = "FailedOperation.DbClientUpdateFailed"
//  FAILEDOPERATION_EXTERNALCONTRACTINDEXNOTFOUND = "FailedOperation.ExternalContractIndexNotFound"
//  FAILEDOPERATION_EXTERNALCONTRACTNOTFOUND = "FailedOperation.ExternalContractNotFound"
//  FAILEDOPERATION_EXTERNALMERCHANTCONTRACTINFOCONFIGNOFOUND = "FailedOperation.ExternalMerchantContractInfoConfigNoFound"
//  FAILEDOPERATION_EXTERNALMERCHANTINDEXCONFIGNOFOUND = "FailedOperation.ExternalMerchantIndexConfigNoFound"
//  FAILEDOPERATION_EXTERNALMERCHANTINFOCONFIGNOFOUND = "FailedOperation.ExternalMerchantInfoConfigNoFound"
//  FAILEDOPERATION_HTTPDOREQUESTERROR = "FailedOperation.HttpDoRequestError"
//  FAILEDOPERATION_MARSHALERROR = "FailedOperation.MarshalError"
//  FAILEDOPERATION_MERCHANTPERMISSIONERROR = "FailedOperation.MerchantPermissionError"
//  FAILEDOPERATION_NOTIFYURLPARSEERROR = "FailedOperation.NotifyUrlParseError"
//  FAILEDOPERATION_PARAMETERERROR = "FailedOperation.ParameterError"
//  FAILEDOPERATION_QUERYCONTRACTNULL = "FailedOperation.QueryContractNull"
//  FAILEDOPERATION_QUERYMCHANNELERROR = "FailedOperation.QueryMchannelError"
//  FAILEDOPERATION_QUERYMODEERROR = "FailedOperation.QueryModeError"
//  FAILEDOPERATION_QUERYRESULTNULL = "FailedOperation.QueryResultNull"
//  FAILEDOPERATION_SIGNERROR = "FailedOperation.SignError"
//  FAILEDOPERATION_UNMARSHALERROR = "FailedOperation.UnmarshalError"
//  FAILEDOPERATION_UPDATECONTRACTSTATUSFAILED = "FailedOperation.UpdateContractStatusFailed"
//  FAILEDOPERATION_XMLFAIL = "FailedOperation.XmlFail"
//  INTERNALERROR_BACKENDCONNECTIONERROR = "InternalError.BackendConnectionError"
//  INTERNALERROR_BACKENDINTERNALERROR = "InternalError.BackendInternalError"
//  INTERNALERROR_BACKENDNETWORKERROR = "InternalError.BackendNetworkError"
//  INTERNALERROR_BACKENDROUTERERROR = "InternalError.BackendRouterError"
//  INTERNALERROR_BACKENDTIMEOUT = "InternalError.BackendTimeOut"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMMARSHALFAILED = "InvalidParameter.ParamMarshalFailed"
//  INVALIDPARAMETER_PARAMUNMARSHALFAILED = "InvalidParameter.ParamUnmarshalFailed"
//  MISSINGPARAMETER_ = "MissingParameter."
func (c *Client) QueryContractPayWayList(request *QueryContractPayWayListRequest) (response *QueryContractPayWayListResponse, err error) {
    if request == nil {
        request = NewQueryContractPayWayListRequest()
    }
    
    response = NewQueryContractPayWayListResponse()
    err = c.Send(request, response)
    return
}

func NewQueryContractRelateShopRequest() (request *QueryContractRelateShopRequest) {
    request = &QueryContractRelateShopRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "QueryContractRelateShop")
    
    
    return
}

func NewQueryContractRelateShopResponse() (response *QueryContractRelateShopResponse) {
    response = &QueryContractRelateShopResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// QueryContractRelateShop
// 云支付-查询合同可关联门店接口
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_BACKCALLERROR = "FailedOperation.BackCallError"
//  FAILEDOPERATION_CALLCHANNELGATEWAYERROR = "FailedOperation.CallChannelGatewayError"
//  FAILEDOPERATION_DBCLIENTCONNECTFAILED = "FailedOperation.DbClientConnectFailed"
//  FAILEDOPERATION_DBCLIENTQUERYFAILED = "FailedOperation.DbClientQueryFailed"
//  FAILEDOPERATION_DBCLIENTUPDATEFAILED = "FailedOperation.DbClientUpdateFailed"
//  FAILEDOPERATION_EXTERNALCONTRACTINDEXNOTFOUND = "FailedOperation.ExternalContractIndexNotFound"
//  FAILEDOPERATION_EXTERNALCONTRACTNOTFOUND = "FailedOperation.ExternalContractNotFound"
//  FAILEDOPERATION_EXTERNALMERCHANTCONTRACTINFOCONFIGNOFOUND = "FailedOperation.ExternalMerchantContractInfoConfigNoFound"
//  FAILEDOPERATION_EXTERNALMERCHANTINDEXCONFIGNOFOUND = "FailedOperation.ExternalMerchantIndexConfigNoFound"
//  FAILEDOPERATION_EXTERNALMERCHANTINFOCONFIGNOFOUND = "FailedOperation.ExternalMerchantInfoConfigNoFound"
//  FAILEDOPERATION_HTTPDOREQUESTERROR = "FailedOperation.HttpDoRequestError"
//  FAILEDOPERATION_MARSHALERROR = "FailedOperation.MarshalError"
//  FAILEDOPERATION_MERCHANTPERMISSIONERROR = "FailedOperation.MerchantPermissionError"
//  FAILEDOPERATION_NOTIFYURLPARSEERROR = "FailedOperation.NotifyUrlParseError"
//  FAILEDOPERATION_PARAMETERERROR = "FailedOperation.ParameterError"
//  FAILEDOPERATION_QUERYCONTRACTNULL = "FailedOperation.QueryContractNull"
//  FAILEDOPERATION_QUERYMCHANNELERROR = "FailedOperation.QueryMchannelError"
//  FAILEDOPERATION_QUERYMODEERROR = "FailedOperation.QueryModeError"
//  FAILEDOPERATION_QUERYRESULTNULL = "FailedOperation.QueryResultNull"
//  FAILEDOPERATION_SIGNERROR = "FailedOperation.SignError"
//  FAILEDOPERATION_UNMARSHALERROR = "FailedOperation.UnmarshalError"
//  FAILEDOPERATION_UPDATECONTRACTSTATUSFAILED = "FailedOperation.UpdateContractStatusFailed"
//  FAILEDOPERATION_XMLFAIL = "FailedOperation.XmlFail"
//  INTERNALERROR_BACKENDCONNECTIONERROR = "InternalError.BackendConnectionError"
//  INTERNALERROR_BACKENDINTERNALERROR = "InternalError.BackendInternalError"
//  INTERNALERROR_BACKENDNETWORKERROR = "InternalError.BackendNetworkError"
//  INTERNALERROR_BACKENDROUTERERROR = "InternalError.BackendRouterError"
//  INTERNALERROR_BACKENDTIMEOUT = "InternalError.BackendTimeOut"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMMARSHALFAILED = "InvalidParameter.ParamMarshalFailed"
//  INVALIDPARAMETER_PARAMUNMARSHALFAILED = "InvalidParameter.ParamUnmarshalFailed"
//  MISSINGPARAMETER_ = "MissingParameter."
func (c *Client) QueryContractRelateShop(request *QueryContractRelateShopRequest) (response *QueryContractRelateShopResponse, err error) {
    if request == nil {
        request = NewQueryContractRelateShopRequest()
    }
    
    response = NewQueryContractRelateShopResponse()
    err = c.Send(request, response)
    return
}

func NewQueryCustAcctIdBalanceRequest() (request *QueryCustAcctIdBalanceRequest) {
    request = &QueryCustAcctIdBalanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "QueryCustAcctIdBalance")
    
    
    return
}

func NewQueryCustAcctIdBalanceResponse() (response *QueryCustAcctIdBalanceResponse) {
    response = &QueryCustAcctIdBalanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// QueryCustAcctIdBalance
// 查询银行子账户余额。查询会员子账户以及平台的功能子账户的余额。
//
// 可能返回的错误码:
//  FAILEDOPERATION_BACKENDERROR = "FailedOperation.BackendError"
//  INTERNALERROR_BACKENDERROR = "InternalError.BackendError"
//  INTERNALERROR_DBACCESSERROR = "InternalError.DBAccessError"
//  INTERNALERROR_FUNDSUMMARYACCTNOINCONSISTENTERROR = "InternalError.FundSummaryAcctNoInconsistentError"
//  INTERNALERROR_PARAMETERERROR = "InternalError.ParameterError"
//  INTERNALERROR_SAVEDBERROR = "InternalError.SaveDBError"
//  INTERNALERROR_SUBACCOUNTNOTFOUNDERROR = "InternalError.SubAccountNotFoundError"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  INTERNALERROR_UNKOWNERROR = "InternalError.UnkownError"
func (c *Client) QueryCustAcctIdBalance(request *QueryCustAcctIdBalanceRequest) (response *QueryCustAcctIdBalanceResponse, err error) {
    if request == nil {
        request = NewQueryCustAcctIdBalanceRequest()
    }
    
    response = NewQueryCustAcctIdBalanceResponse()
    err = c.Send(request, response)
    return
}

func NewQueryDownloadBillURLRequest() (request *QueryDownloadBillURLRequest) {
    request = &QueryDownloadBillURLRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "QueryDownloadBillURL")
    
    
    return
}

func NewQueryDownloadBillURLResponse() (response *QueryDownloadBillURLResponse) {
    response = &QueryDownloadBillURLResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// QueryDownloadBillURL
// 云鉴-查询对账单下载地址的接口
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADBILLERROR = "FailedOperation.DownloadBillError"
//  FAILEDOPERATION_INVALIDPARAMETER = "FailedOperation.InvalidParameter"
//  FAILEDOPERATION_MISSINGPARAMETER = "FailedOperation.MissingParameter"
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
func (c *Client) QueryDownloadBillURL(request *QueryDownloadBillURLRequest) (response *QueryDownloadBillURLResponse, err error) {
    if request == nil {
        request = NewQueryDownloadBillURLRequest()
    }
    
    response = NewQueryDownloadBillURLResponse()
    err = c.Send(request, response)
    return
}

func NewQueryExchangeRateRequest() (request *QueryExchangeRateRequest) {
    request = &QueryExchangeRateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "QueryExchangeRate")
    
    
    return
}

func NewQueryExchangeRateResponse() (response *QueryExchangeRateResponse) {
    response = &QueryExchangeRateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// QueryExchangeRate
// 跨境-查询汇率
//
// 可能返回的错误码:
//  INTERNALERROR_PARAMETERERROR = "InternalError.ParameterError"
//  INTERNALERROR_UNKOWNERROR = "InternalError.UnkownError"
//  RESOURCENOTFOUND_MERCHANTINFONOTFOUND = "ResourceNotFound.MerchantInfoNotFound"
func (c *Client) QueryExchangeRate(request *QueryExchangeRateRequest) (response *QueryExchangeRateResponse, err error) {
    if request == nil {
        request = NewQueryExchangeRateRequest()
    }
    
    response = NewQueryExchangeRateResponse()
    err = c.Send(request, response)
    return
}

func NewQueryInvoiceRequest() (request *QueryInvoiceRequest) {
    request = &QueryInvoiceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "QueryInvoice")
    
    
    return
}

func NewQueryInvoiceResponse() (response *QueryInvoiceResponse) {
    response = &QueryInvoiceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// QueryInvoice
// 智慧零售-发票查询
//
// 可能返回的错误码:
//  FAILEDOPERATION_BACKENDERROR = "FailedOperation.BackendError"
//  FAILEDOPERATION_INVOICEEXIST = "FailedOperation.InvoiceExist"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
//  INTERNALERROR_BACKENDERROR = "InternalError.BackendError"
//  INTERNALERROR_DUPLICATEKEYERROR = "InternalError.DuplicateKeyError"
//  INTERNALERROR_INVOICEEXIST = "InternalError.InvoiceExist"
//  INTERNALERROR_SANDBOXACCESSERROR = "InternalError.SandBoxAccessError"
//  INTERNALERROR_SAVEDBERROR = "InternalError.SaveDBError"
//  INTERNALERROR_SIGGENERROR = "InternalError.SigGenError"
//  INTERNALERROR_UNKOWNERROR = "InternalError.UnkownError"
//  INVALIDPARAMETER_LACKPARAMETER = "InvalidParameter.LackParameter"
//  INVALIDPARAMETER_UNSUPPORTEDPARAMETER = "InvalidParameter.UnsupportedParameter"
//  RESOURCEINSUFFICIENT_THREADPOOLREJECT = "ResourceInsufficient.ThreadPoolReject"
//  RESOURCENOTFOUND_INVOICENOTFOUND = "ResourceNotFound.InvoiceNotFound"
//  RESOURCENOTFOUND_MERCHANTINFONOTFOUND = "ResourceNotFound.MerchantInfoNotFound"
//  RESOURCENOTFOUND_PLATFORMINFONOTFOUND = "ResourceNotFound.PlatformInfoNotFound"
func (c *Client) QueryInvoice(request *QueryInvoiceRequest) (response *QueryInvoiceResponse, err error) {
    if request == nil {
        request = NewQueryInvoiceRequest()
    }
    
    response = NewQueryInvoiceResponse()
    err = c.Send(request, response)
    return
}

func NewQueryInvoiceV2Request() (request *QueryInvoiceV2Request) {
    request = &QueryInvoiceV2Request{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "QueryInvoiceV2")
    
    
    return
}

func NewQueryInvoiceV2Response() (response *QueryInvoiceV2Response) {
    response = &QueryInvoiceV2Response{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// QueryInvoiceV2
// 智慧零售-发票查询V2
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_INVOICENOTFOUND = "ResourceNotFound.InvoiceNotFound"
func (c *Client) QueryInvoiceV2(request *QueryInvoiceV2Request) (response *QueryInvoiceV2Response, err error) {
    if request == nil {
        request = NewQueryInvoiceV2Request()
    }
    
    response = NewQueryInvoiceV2Response()
    err = c.Send(request, response)
    return
}

func NewQueryMaliciousRegistrationRequest() (request *QueryMaliciousRegistrationRequest) {
    request = &QueryMaliciousRegistrationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "QueryMaliciousRegistration")
    
    
    return
}

func NewQueryMaliciousRegistrationResponse() (response *QueryMaliciousRegistrationResponse) {
    response = &QueryMaliciousRegistrationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// QueryMaliciousRegistration
// 商户恶意注册接口
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_INVOICENOTFOUND = "ResourceNotFound.InvoiceNotFound"
func (c *Client) QueryMaliciousRegistration(request *QueryMaliciousRegistrationRequest) (response *QueryMaliciousRegistrationResponse, err error) {
    if request == nil {
        request = NewQueryMaliciousRegistrationRequest()
    }
    
    response = NewQueryMaliciousRegistrationResponse()
    err = c.Send(request, response)
    return
}

func NewQueryMemberBindRequest() (request *QueryMemberBindRequest) {
    request = &QueryMemberBindRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "QueryMemberBind")
    
    
    return
}

func NewQueryMemberBindResponse() (response *QueryMemberBindResponse) {
    response = &QueryMemberBindResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// QueryMemberBind
// 会员绑定信息查询。查询标志为“单个会员”的情况下，返回该会员的有效的绑定账户信息。
//
// 查询标志为“全部会员”的情况下，返回市场下的全部的有效的绑定账户信息。查询标志为“单个会员的证件信息”的情况下，返回市场下的指定的会员的留存在电商见证宝系统的证件信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION_BACKENDERROR = "FailedOperation.BackendError"
//  INTERNALERROR_BACKENDERROR = "InternalError.BackendError"
//  INTERNALERROR_DBACCESSERROR = "InternalError.DBAccessError"
//  INTERNALERROR_FUNDSUMMARYACCTNOINCONSISTENTERROR = "InternalError.FundSummaryAcctNoInconsistentError"
//  INTERNALERROR_PARAMETERERROR = "InternalError.ParameterError"
//  INTERNALERROR_SAVEDBERROR = "InternalError.SaveDBError"
//  INTERNALERROR_SUBACCOUNTNOTFOUNDERROR = "InternalError.SubAccountNotFoundError"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  INTERNALERROR_UNKOWNERROR = "InternalError.UnkownError"
func (c *Client) QueryMemberBind(request *QueryMemberBindRequest) (response *QueryMemberBindResponse, err error) {
    if request == nil {
        request = NewQueryMemberBindRequest()
    }
    
    response = NewQueryMemberBindResponse()
    err = c.Send(request, response)
    return
}

func NewQueryMemberTransactionRequest() (request *QueryMemberTransactionRequest) {
    request = &QueryMemberTransactionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "QueryMemberTransaction")
    
    
    return
}

func NewQueryMemberTransactionResponse() (response *QueryMemberTransactionResponse) {
    response = &QueryMemberTransactionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// QueryMemberTransaction
// 会员间交易-不验证。此接口可以实现会员间的余额的交易，实现资金在会员之间流动。
//
// 可能返回的错误码:
//  FAILEDOPERATION_BACKENDERROR = "FailedOperation.BackendError"
//  INTERNALERROR_BACKENDERROR = "InternalError.BackendError"
//  INTERNALERROR_DBACCESSERROR = "InternalError.DBAccessError"
//  INTERNALERROR_FUNDSUMMARYACCTNOINCONSISTENTERROR = "InternalError.FundSummaryAcctNoInconsistentError"
//  INTERNALERROR_PARAMETERERROR = "InternalError.ParameterError"
//  INTERNALERROR_SAVEDBERROR = "InternalError.SaveDBError"
//  INTERNALERROR_SUBACCOUNTNOTFOUNDERROR = "InternalError.SubAccountNotFoundError"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  INTERNALERROR_UNKOWNERROR = "InternalError.UnkownError"
func (c *Client) QueryMemberTransaction(request *QueryMemberTransactionRequest) (response *QueryMemberTransactionResponse, err error) {
    if request == nil {
        request = NewQueryMemberTransactionRequest()
    }
    
    response = NewQueryMemberTransactionResponse()
    err = c.Send(request, response)
    return
}

func NewQueryMerchantRequest() (request *QueryMerchantRequest) {
    request = &QueryMerchantRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "QueryMerchant")
    
    
    return
}

func NewQueryMerchantResponse() (response *QueryMerchantResponse) {
    response = &QueryMerchantResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// QueryMerchant
// 云鉴-商户信息查询接口
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDPARAMETER = "FailedOperation.InvalidParameter"
//  FAILEDOPERATION_MERCHANTNOTEXIST = "FailedOperation.MerchantNotExist"
//  FAILEDOPERATION_MISSINGPARAMETER = "FailedOperation.MissingParameter"
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
func (c *Client) QueryMerchant(request *QueryMerchantRequest) (response *QueryMerchantResponse, err error) {
    if request == nil {
        request = NewQueryMerchantRequest()
    }
    
    response = NewQueryMerchantResponse()
    err = c.Send(request, response)
    return
}

func NewQueryMerchantBalanceRequest() (request *QueryMerchantBalanceRequest) {
    request = &QueryMerchantBalanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "QueryMerchantBalance")
    
    
    return
}

func NewQueryMerchantBalanceResponse() (response *QueryMerchantBalanceResponse) {
    response = &QueryMerchantBalanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// QueryMerchantBalance
// 跨境-对接方账户余额查询
//
// 可能返回的错误码:
//  INTERNALERROR_PARAMETERERROR = "InternalError.ParameterError"
//  INTERNALERROR_UNKOWNERROR = "InternalError.UnkownError"
//  RESOURCENOTFOUND_MERCHANTINFONOTFOUND = "ResourceNotFound.MerchantInfoNotFound"
func (c *Client) QueryMerchantBalance(request *QueryMerchantBalanceRequest) (response *QueryMerchantBalanceResponse, err error) {
    if request == nil {
        request = NewQueryMerchantBalanceRequest()
    }
    
    response = NewQueryMerchantBalanceResponse()
    err = c.Send(request, response)
    return
}

func NewQueryMerchantClassificationRequest() (request *QueryMerchantClassificationRequest) {
    request = &QueryMerchantClassificationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "QueryMerchantClassification")
    
    
    return
}

func NewQueryMerchantClassificationResponse() (response *QueryMerchantClassificationResponse) {
    response = &QueryMerchantClassificationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// QueryMerchantClassification
// 云支付-查询商户分类接口
//
// 可能返回的错误码:
//  INTERNALERROR_PARAMETERERROR = "InternalError.ParameterError"
//  INTERNALERROR_UNKOWNERROR = "InternalError.UnkownError"
//  RESOURCENOTFOUND_MERCHANTINFONOTFOUND = "ResourceNotFound.MerchantInfoNotFound"
func (c *Client) QueryMerchantClassification(request *QueryMerchantClassificationRequest) (response *QueryMerchantClassificationResponse, err error) {
    if request == nil {
        request = NewQueryMerchantClassificationRequest()
    }
    
    response = NewQueryMerchantClassificationResponse()
    err = c.Send(request, response)
    return
}

func NewQueryMerchantInfoForManagementRequest() (request *QueryMerchantInfoForManagementRequest) {
    request = &QueryMerchantInfoForManagementRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "QueryMerchantInfoForManagement")
    
    
    return
}

func NewQueryMerchantInfoForManagementResponse() (response *QueryMerchantInfoForManagementResponse) {
    response = &QueryMerchantInfoForManagementResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// QueryMerchantInfoForManagement
// 智慧零售-查询管理端商户
//
// 可能返回的错误码:
//  INTERNALERROR_BACKENDERROR = "InternalError.BackendError"
//  INTERNALERROR_DUPLICATEKEYERROR = "InternalError.DuplicateKeyError"
//  INTERNALERROR_SAVEDBERROR = "InternalError.SaveDBError"
//  INTERNALERROR_SIGGENERROR = "InternalError.SigGenError"
//  INTERNALERROR_UNKOWNERROR = "InternalError.UnkownError"
//  INVALIDPARAMETER_LACKPARAMETER = "InvalidParameter.LackParameter"
//  RESOURCENOTFOUND_INVOICENOTFOUND = "ResourceNotFound.InvoiceNotFound"
//  RESOURCENOTFOUND_MERCHANTINFONOTFOUND = "ResourceNotFound.MerchantInfoNotFound"
//  RESOURCENOTFOUND_PLATFORMINFONOTFOUND = "ResourceNotFound.PlatformInfoNotFound"
func (c *Client) QueryMerchantInfoForManagement(request *QueryMerchantInfoForManagementRequest) (response *QueryMerchantInfoForManagementResponse, err error) {
    if request == nil {
        request = NewQueryMerchantInfoForManagementRequest()
    }
    
    response = NewQueryMerchantInfoForManagementResponse()
    err = c.Send(request, response)
    return
}

func NewQueryMerchantOrderRequest() (request *QueryMerchantOrderRequest) {
    request = &QueryMerchantOrderRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "QueryMerchantOrder")
    
    
    return
}

func NewQueryMerchantOrderResponse() (response *QueryMerchantOrderResponse) {
    response = &QueryMerchantOrderResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// QueryMerchantOrder
// 云鉴-消费订单查询接口
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDPARAMETER = "FailedOperation.InvalidParameter"
//  FAILEDOPERATION_MISSINGPARAMETER = "FailedOperation.MissingParameter"
//  FAILEDOPERATION_QUERYORDERERROR = "FailedOperation.QueryOrderError"
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
func (c *Client) QueryMerchantOrder(request *QueryMerchantOrderRequest) (response *QueryMerchantOrderResponse, err error) {
    if request == nil {
        request = NewQueryMerchantOrderRequest()
    }
    
    response = NewQueryMerchantOrderResponse()
    err = c.Send(request, response)
    return
}

func NewQueryMerchantPayWayListRequest() (request *QueryMerchantPayWayListRequest) {
    request = &QueryMerchantPayWayListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "QueryMerchantPayWayList")
    
    
    return
}

func NewQueryMerchantPayWayListResponse() (response *QueryMerchantPayWayListResponse) {
    response = &QueryMerchantPayWayListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// QueryMerchantPayWayList
// 商户查询已开通的支付方式列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDPARAMETER = "FailedOperation.InvalidParameter"
//  FAILEDOPERATION_MISSINGPARAMETER = "FailedOperation.MissingParameter"
//  FAILEDOPERATION_QUERYORDERERROR = "FailedOperation.QueryOrderError"
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
func (c *Client) QueryMerchantPayWayList(request *QueryMerchantPayWayListRequest) (response *QueryMerchantPayWayListResponse, err error) {
    if request == nil {
        request = NewQueryMerchantPayWayListRequest()
    }
    
    response = NewQueryMerchantPayWayListResponse()
    err = c.Send(request, response)
    return
}

func NewQueryOrderRequest() (request *QueryOrderRequest) {
    request = &QueryOrderRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "QueryOrder")
    
    
    return
}

func NewQueryOrderResponse() (response *QueryOrderResponse) {
    response = &QueryOrderResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// QueryOrder
// 根据订单号，或者用户Id，查询支付订单状态 
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_APPDENY = "FailedOperation.AppDeny"
//  FAILEDOPERATION_NORECORD = "FailedOperation.NoRecord"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) QueryOrder(request *QueryOrderRequest) (response *QueryOrderResponse, err error) {
    if request == nil {
        request = NewQueryOrderRequest()
    }
    
    response = NewQueryOrderResponse()
    err = c.Send(request, response)
    return
}

func NewQueryOrderStatusRequest() (request *QueryOrderStatusRequest) {
    request = &QueryOrderStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "QueryOrderStatus")
    
    
    return
}

func NewQueryOrderStatusResponse() (response *QueryOrderStatusResponse) {
    response = &QueryOrderStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// QueryOrderStatus
// 云支付-查询订单付款状态
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_APPDENY = "FailedOperation.AppDeny"
//  FAILEDOPERATION_NORECORD = "FailedOperation.NoRecord"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) QueryOrderStatus(request *QueryOrderStatusRequest) (response *QueryOrderStatusResponse, err error) {
    if request == nil {
        request = NewQueryOrderStatusRequest()
    }
    
    response = NewQueryOrderStatusResponse()
    err = c.Send(request, response)
    return
}

func NewQueryOutwardOrderRequest() (request *QueryOutwardOrderRequest) {
    request = &QueryOutwardOrderRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "QueryOutwardOrder")
    
    
    return
}

func NewQueryOutwardOrderResponse() (response *QueryOutwardOrderResponse) {
    response = &QueryOutwardOrderResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// QueryOutwardOrder
// 跨境-查询汇出结果
//
// 可能返回的错误码:
//  INTERNALERROR_PARAMETERERROR = "InternalError.ParameterError"
//  INTERNALERROR_UNKOWNERROR = "InternalError.UnkownError"
//  RESOURCENOTFOUND_MERCHANTINFONOTFOUND = "ResourceNotFound.MerchantInfoNotFound"
func (c *Client) QueryOutwardOrder(request *QueryOutwardOrderRequest) (response *QueryOutwardOrderResponse, err error) {
    if request == nil {
        request = NewQueryOutwardOrderRequest()
    }
    
    response = NewQueryOutwardOrderResponse()
    err = c.Send(request, response)
    return
}

func NewQueryPayerInfoRequest() (request *QueryPayerInfoRequest) {
    request = &QueryPayerInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "QueryPayerInfo")
    
    
    return
}

func NewQueryPayerInfoResponse() (response *QueryPayerInfoResponse) {
    response = &QueryPayerInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// QueryPayerInfo
// 跨境-付款人查询
//
// 可能返回的错误码:
//  INTERNALERROR_PARAMETERERROR = "InternalError.ParameterError"
//  INTERNALERROR_UNKOWNERROR = "InternalError.UnkownError"
//  RESOURCENOTFOUND_MERCHANTINFONOTFOUND = "ResourceNotFound.MerchantInfoNotFound"
func (c *Client) QueryPayerInfo(request *QueryPayerInfoRequest) (response *QueryPayerInfoResponse, err error) {
    if request == nil {
        request = NewQueryPayerInfoRequest()
    }
    
    response = NewQueryPayerInfoResponse()
    err = c.Send(request, response)
    return
}

func NewQueryReconciliationDocumentRequest() (request *QueryReconciliationDocumentRequest) {
    request = &QueryReconciliationDocumentRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "QueryReconciliationDocument")
    
    
    return
}

func NewQueryReconciliationDocumentResponse() (response *QueryReconciliationDocumentResponse) {
    response = &QueryReconciliationDocumentResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// QueryReconciliationDocument
// 查询对账文件信息。平台调用该接口获取需下载对账文件的文件名称以及密钥。 平台获取到信息后， 可以再调用OPENAPI的文件下载功能。
//
// 可能返回的错误码:
//  FAILEDOPERATION_BACKENDERROR = "FailedOperation.BackendError"
//  INTERNALERROR_BACKENDERROR = "InternalError.BackendError"
//  INTERNALERROR_DBACCESSERROR = "InternalError.DBAccessError"
//  INTERNALERROR_FUNDSUMMARYACCTNOINCONSISTENTERROR = "InternalError.FundSummaryAcctNoInconsistentError"
//  INTERNALERROR_PARAMETERERROR = "InternalError.ParameterError"
//  INTERNALERROR_SAVEDBERROR = "InternalError.SaveDBError"
//  INTERNALERROR_SUBACCOUNTNOTFOUNDERROR = "InternalError.SubAccountNotFoundError"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  INTERNALERROR_UNKOWNERROR = "InternalError.UnkownError"
func (c *Client) QueryReconciliationDocument(request *QueryReconciliationDocumentRequest) (response *QueryReconciliationDocumentResponse, err error) {
    if request == nil {
        request = NewQueryReconciliationDocumentRequest()
    }
    
    response = NewQueryReconciliationDocumentResponse()
    err = c.Send(request, response)
    return
}

func NewQueryRefundRequest() (request *QueryRefundRequest) {
    request = &QueryRefundRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "QueryRefund")
    
    
    return
}

func NewQueryRefundResponse() (response *QueryRefundResponse) {
    response = &QueryRefundResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// QueryRefund
// 提交退款申请后，通过调用该接口查询退款状态。退款可能有一定延时，用微信零钱支付的退款约20分钟内到账，银行卡支付的退款约3个工作日后到账。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_APPDENY = "FailedOperation.AppDeny"
//  FAILEDOPERATION_NORECORD = "FailedOperation.NoRecord"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) QueryRefund(request *QueryRefundRequest) (response *QueryRefundResponse, err error) {
    if request == nil {
        request = NewQueryRefundRequest()
    }
    
    response = NewQueryRefundResponse()
    err = c.Send(request, response)
    return
}

func NewQueryShopOpenIdRequest() (request *QueryShopOpenIdRequest) {
    request = &QueryShopOpenIdRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "QueryShopOpenId")
    
    
    return
}

func NewQueryShopOpenIdResponse() (response *QueryShopOpenIdResponse) {
    response = &QueryShopOpenIdResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// QueryShopOpenId
// 云支付-获取门店OpenId接口
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_APPDENY = "FailedOperation.AppDeny"
//  FAILEDOPERATION_NORECORD = "FailedOperation.NoRecord"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) QueryShopOpenId(request *QueryShopOpenIdRequest) (response *QueryShopOpenIdResponse, err error) {
    if request == nil {
        request = NewQueryShopOpenIdRequest()
    }
    
    response = NewQueryShopOpenIdResponse()
    err = c.Send(request, response)
    return
}

func NewQuerySinglePayRequest() (request *QuerySinglePayRequest) {
    request = &QuerySinglePayRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "QuerySinglePay")
    
    
    return
}

func NewQuerySinglePayResponse() (response *QuerySinglePayResponse) {
    response = &QuerySinglePayResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// QuerySinglePay
// 银企直连-单笔支付状态查询接口
//
// 可能返回的错误码:
//  INTERNALERROR_BACKENDERROR = "InternalError.BackendError"
//  INTERNALERROR_SANDBOXACCESSERROR = "InternalError.SandBoxAccessError"
//  INTERNALERROR_SIGGENERROR = "InternalError.SigGenError"
//  INTERNALERROR_UNKOWNERROR = "InternalError.UnkownError"
//  INVALIDPARAMETER_LACKPARAMETER = "InvalidParameter.LackParameter"
//  RESOURCENOTFOUND_PLATFORMINFONOTFOUND = "ResourceNotFound.PlatformInfoNotFound"
func (c *Client) QuerySinglePay(request *QuerySinglePayRequest) (response *QuerySinglePayResponse, err error) {
    if request == nil {
        request = NewQuerySinglePayRequest()
    }
    
    response = NewQuerySinglePayResponse()
    err = c.Send(request, response)
    return
}

func NewQuerySinglePaymentResultRequest() (request *QuerySinglePaymentResultRequest) {
    request = &QuerySinglePaymentResultRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "QuerySinglePaymentResult")
    
    
    return
}

func NewQuerySinglePaymentResultResponse() (response *QuerySinglePaymentResultResponse) {
    response = &QuerySinglePaymentResultResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// QuerySinglePaymentResult
// 灵云-单笔转账结果查询
//
// 可能返回的错误码:
//  INTERNALERROR_BACKENDERROR = "InternalError.BackendError"
//  INTERNALERROR_SANDBOXACCESSERROR = "InternalError.SandBoxAccessError"
//  INTERNALERROR_SIGGENERROR = "InternalError.SigGenError"
//  INTERNALERROR_UNKOWNERROR = "InternalError.UnkownError"
//  INVALIDPARAMETER_LACKPARAMETER = "InvalidParameter.LackParameter"
//  RESOURCENOTFOUND_PLATFORMINFONOTFOUND = "ResourceNotFound.PlatformInfoNotFound"
func (c *Client) QuerySinglePaymentResult(request *QuerySinglePaymentResultRequest) (response *QuerySinglePaymentResultResponse, err error) {
    if request == nil {
        request = NewQuerySinglePaymentResultRequest()
    }
    
    response = NewQuerySinglePaymentResultResponse()
    err = c.Send(request, response)
    return
}

func NewQuerySingleTransactionStatusRequest() (request *QuerySingleTransactionStatusRequest) {
    request = &QuerySingleTransactionStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "QuerySingleTransactionStatus")
    
    
    return
}

func NewQuerySingleTransactionStatusResponse() (response *QuerySingleTransactionStatusResponse) {
    response = &QuerySingleTransactionStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// QuerySingleTransactionStatus
// 查询银行单笔交易状态。查询单笔交易的状态。
//
// 可能返回的错误码:
//  FAILEDOPERATION_BACKENDERROR = "FailedOperation.BackendError"
//  INTERNALERROR_BACKENDERROR = "InternalError.BackendError"
//  INTERNALERROR_DBACCESSERROR = "InternalError.DBAccessError"
//  INTERNALERROR_FUNDSUMMARYACCTNOINCONSISTENTERROR = "InternalError.FundSummaryAcctNoInconsistentError"
//  INTERNALERROR_PARAMETERERROR = "InternalError.ParameterError"
//  INTERNALERROR_SAVEDBERROR = "InternalError.SaveDBError"
//  INTERNALERROR_SUBACCOUNTNOTFOUNDERROR = "InternalError.SubAccountNotFoundError"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  INTERNALERROR_UNKOWNERROR = "InternalError.UnkownError"
func (c *Client) QuerySingleTransactionStatus(request *QuerySingleTransactionStatusRequest) (response *QuerySingleTransactionStatusResponse, err error) {
    if request == nil {
        request = NewQuerySingleTransactionStatusRequest()
    }
    
    response = NewQuerySingleTransactionStatusResponse()
    err = c.Send(request, response)
    return
}

func NewQuerySmallAmountTransferRequest() (request *QuerySmallAmountTransferRequest) {
    request = &QuerySmallAmountTransferRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "QuerySmallAmountTransfer")
    
    
    return
}

func NewQuerySmallAmountTransferResponse() (response *QuerySmallAmountTransferResponse) {
    response = &QuerySmallAmountTransferResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// QuerySmallAmountTransfer
// 查询小额鉴权转账结果。查询小额往账鉴权的转账状态。
//
// 可能返回的错误码:
//  FAILEDOPERATION_BACKENDERROR = "FailedOperation.BackendError"
//  INTERNALERROR_BACKENDERROR = "InternalError.BackendError"
//  INTERNALERROR_DBACCESSERROR = "InternalError.DBAccessError"
//  INTERNALERROR_FUNDSUMMARYACCTNOINCONSISTENTERROR = "InternalError.FundSummaryAcctNoInconsistentError"
//  INTERNALERROR_PARAMETERERROR = "InternalError.ParameterError"
//  INTERNALERROR_SAVEDBERROR = "InternalError.SaveDBError"
//  INTERNALERROR_SUBACCOUNTNOTFOUNDERROR = "InternalError.SubAccountNotFoundError"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  INTERNALERROR_UNKOWNERROR = "InternalError.UnkownError"
func (c *Client) QuerySmallAmountTransfer(request *QuerySmallAmountTransferRequest) (response *QuerySmallAmountTransferResponse, err error) {
    if request == nil {
        request = NewQuerySmallAmountTransferRequest()
    }
    
    response = NewQuerySmallAmountTransferResponse()
    err = c.Send(request, response)
    return
}

func NewQueryTradeRequest() (request *QueryTradeRequest) {
    request = &QueryTradeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "QueryTrade")
    
    
    return
}

func NewQueryTradeResponse() (response *QueryTradeResponse) {
    response = &QueryTradeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// QueryTrade
// 跨境-贸易材料明细查询。
//
// 可能返回的错误码:
//  INTERNALERROR_PARAMETERERROR = "InternalError.ParameterError"
//  INTERNALERROR_UNKOWNERROR = "InternalError.UnkownError"
//  RESOURCENOTFOUND_MERCHANTINFONOTFOUND = "ResourceNotFound.MerchantInfoNotFound"
func (c *Client) QueryTrade(request *QueryTradeRequest) (response *QueryTradeResponse, err error) {
    if request == nil {
        request = NewQueryTradeRequest()
    }
    
    response = NewQueryTradeResponse()
    err = c.Send(request, response)
    return
}

func NewQueryTransferBatchRequest() (request *QueryTransferBatchRequest) {
    request = &QueryTransferBatchRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "QueryTransferBatch")
    
    
    return
}

func NewQueryTransferBatchResponse() (response *QueryTransferBatchResponse) {
    response = &QueryTransferBatchResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// QueryTransferBatch
// 通过商家批次单号或者微信批次号查询批次单
//
// 可能返回的错误码:
//  FAILEDOPERATION_ALREADYEXISTS = "FailedOperation.AlreadyExists"
//  FAILEDOPERATION_APPIDMCHIDNOTMATCH = "FailedOperation.AppidMchidNotMatch"
//  FAILEDOPERATION_FREQUENCYLIMITED = "FailedOperation.FrequencyLimited"
//  FAILEDOPERATION_INVALIDREQUEST = "FailedOperation.InvalidRequest"
//  FAILEDOPERATION_NOAUTH = "FailedOperation.NoAuth"
//  FAILEDOPERATION_NOTENOUGH = "FailedOperation.NotEnough"
//  FAILEDOPERATION_NOTFOUND = "FailedOperation.NotFound"
//  FAILEDOPERATION_PARAMERROR = "FailedOperation.ParamError"
//  FAILEDOPERATION_QUOTAEXCEED = "FailedOperation.QuotaExceed"
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR_ = "InternalError."
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER_ = "MissingParameter."
func (c *Client) QueryTransferBatch(request *QueryTransferBatchRequest) (response *QueryTransferBatchResponse, err error) {
    if request == nil {
        request = NewQueryTransferBatchRequest()
    }
    
    response = NewQueryTransferBatchResponse()
    err = c.Send(request, response)
    return
}

func NewQueryTransferDetailRequest() (request *QueryTransferDetailRequest) {
    request = &QueryTransferDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "QueryTransferDetail")
    
    
    return
}

func NewQueryTransferDetailResponse() (response *QueryTransferDetailResponse) {
    response = &QueryTransferDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// QueryTransferDetail
// 通过商家或者微信批次明细单号查询明细单
//
// 可能返回的错误码:
//  FAILEDOPERATION_ALREADYEXISTS = "FailedOperation.AlreadyExists"
//  FAILEDOPERATION_APPIDMCHIDNOTMATCH = "FailedOperation.AppidMchidNotMatch"
//  FAILEDOPERATION_FREQUENCYLIMITED = "FailedOperation.FrequencyLimited"
//  FAILEDOPERATION_INVALIDREQUEST = "FailedOperation.InvalidRequest"
//  FAILEDOPERATION_NOAUTH = "FailedOperation.NoAuth"
//  FAILEDOPERATION_NOTENOUGH = "FailedOperation.NotEnough"
//  FAILEDOPERATION_NOTFOUND = "FailedOperation.NotFound"
//  FAILEDOPERATION_PARAMERROR = "FailedOperation.ParamError"
//  FAILEDOPERATION_QUOTAEXCEED = "FailedOperation.QuotaExceed"
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR_ = "InternalError."
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER_ = "MissingParameter."
func (c *Client) QueryTransferDetail(request *QueryTransferDetailRequest) (response *QueryTransferDetailResponse, err error) {
    if request == nil {
        request = NewQueryTransferDetailRequest()
    }
    
    response = NewQueryTransferDetailResponse()
    err = c.Send(request, response)
    return
}

func NewQueryTransferResultRequest() (request *QueryTransferResultRequest) {
    request = &QueryTransferResultRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "QueryTransferResult")
    
    
    return
}

func NewQueryTransferResultResponse() (response *QueryTransferResultResponse) {
    response = &QueryTransferResultResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// QueryTransferResult
// 智能代发-单笔代发转账查询接口
//
// 可能返回的错误码:
//  FAILEDOPERATION_ALREADYEXISTS = "FailedOperation.AlreadyExists"
//  FAILEDOPERATION_APPIDMCHIDNOTMATCH = "FailedOperation.AppidMchidNotMatch"
//  FAILEDOPERATION_FREQUENCYLIMITED = "FailedOperation.FrequencyLimited"
//  FAILEDOPERATION_INVALIDREQUEST = "FailedOperation.InvalidRequest"
//  FAILEDOPERATION_NOAUTH = "FailedOperation.NoAuth"
//  FAILEDOPERATION_NOTENOUGH = "FailedOperation.NotEnough"
//  FAILEDOPERATION_NOTFOUND = "FailedOperation.NotFound"
//  FAILEDOPERATION_PARAMERROR = "FailedOperation.ParamError"
//  FAILEDOPERATION_QUOTAEXCEED = "FailedOperation.QuotaExceed"
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR_ = "InternalError."
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER_ = "MissingParameter."
func (c *Client) QueryTransferResult(request *QueryTransferResultRequest) (response *QueryTransferResultResponse, err error) {
    if request == nil {
        request = NewQueryTransferResultRequest()
    }
    
    response = NewQueryTransferResultResponse()
    err = c.Send(request, response)
    return
}

func NewRechargeByThirdPayRequest() (request *RechargeByThirdPayRequest) {
    request = &RechargeByThirdPayRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "RechargeByThirdPay")
    
    
    return
}

func NewRechargeByThirdPayResponse() (response *RechargeByThirdPayResponse) {
    response = &RechargeByThirdPayResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// RechargeByThirdPay
// 会员在途充值(经第三方支付渠道)接口
//
// 可能返回的错误码:
//  FAILEDOPERATION_ACTIONINVALID = "FailedOperation.ActionInvalid"
//  FAILEDOPERATION_PABANKERROR = "FailedOperation.PABankError"
func (c *Client) RechargeByThirdPay(request *RechargeByThirdPayRequest) (response *RechargeByThirdPayResponse, err error) {
    if request == nil {
        request = NewRechargeByThirdPayRequest()
    }
    
    response = NewRechargeByThirdPayResponse()
    err = c.Send(request, response)
    return
}

func NewRechargeMemberThirdPayRequest() (request *RechargeMemberThirdPayRequest) {
    request = &RechargeMemberThirdPayRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "RechargeMemberThirdPay")
    
    
    return
}

func NewRechargeMemberThirdPayResponse() (response *RechargeMemberThirdPayResponse) {
    response = &RechargeMemberThirdPayResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// RechargeMemberThirdPay
// 见证宝-会员在途充值(经第三方支付渠道)
//
// 可能返回的错误码:
//  FAILEDOPERATION_BACKENDERROR = "FailedOperation.BackendError"
//  INTERNALERROR_BACKENDERROR = "InternalError.BackendError"
//  INTERNALERROR_DBACCESSERROR = "InternalError.DBAccessError"
//  INTERNALERROR_FUNDSUMMARYACCTNOINCONSISTENTERROR = "InternalError.FundSummaryAcctNoInconsistentError"
//  INTERNALERROR_PARAMETERERROR = "InternalError.ParameterError"
//  INTERNALERROR_SAVEDBERROR = "InternalError.SaveDBError"
//  INTERNALERROR_SUBACCOUNTNOTFOUNDERROR = "InternalError.SubAccountNotFoundError"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  INTERNALERROR_UNKOWNERROR = "InternalError.UnkownError"
func (c *Client) RechargeMemberThirdPay(request *RechargeMemberThirdPayRequest) (response *RechargeMemberThirdPayResponse, err error) {
    if request == nil {
        request = NewRechargeMemberThirdPayRequest()
    }
    
    response = NewRechargeMemberThirdPayResponse()
    err = c.Send(request, response)
    return
}

func NewRefundRequest() (request *RefundRequest) {
    request = &RefundRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "Refund")
    
    
    return
}

func NewRefundResponse() (response *RefundResponse) {
    response = &RefundResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// Refund
// 如交易订单需退款，可以通过本接口将支付款全部或部分退还给付款方，聚鑫将在收到退款请求并且验证成功之后，按照退款规则将支付款按原路退回到支付帐号。最长支持1年的订单退款。在订单包含多个子订单的情况下，如果使用本接口传入OutTradeNo或TransactionId退款，则只支持全单退款；如果需要部分退款，请通过传入子订单的方式来指定部分金额退款。 
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_APPDENY = "FailedOperation.AppDeny"
//  FAILEDOPERATION_NORECORD = "FailedOperation.NoRecord"
//  FAILEDOPERATION_PORTALERROR = "FailedOperation.PortalError"
//  FAILEDOPERATION_WXCRTNOTSET = "FailedOperation.WxCrtNotSet"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) Refund(request *RefundRequest) (response *RefundResponse, err error) {
    if request == nil {
        request = NewRefundRequest()
    }
    
    response = NewRefundResponse()
    err = c.Send(request, response)
    return
}

func NewRefundMemberTransactionRequest() (request *RefundMemberTransactionRequest) {
    request = &RefundMemberTransactionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "RefundMemberTransaction")
    
    
    return
}

func NewRefundMemberTransactionResponse() (response *RefundMemberTransactionResponse) {
    response = &RefundMemberTransactionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// RefundMemberTransaction
// 会员间交易退款
//
// 可能返回的错误码:
//  FAILEDOPERATION_ACTIONINVALID = "FailedOperation.ActionInvalid"
//  FAILEDOPERATION_PABANKERROR = "FailedOperation.PABankError"
func (c *Client) RefundMemberTransaction(request *RefundMemberTransactionRequest) (response *RefundMemberTransactionResponse, err error) {
    if request == nil {
        request = NewRefundMemberTransactionRequest()
    }
    
    response = NewRefundMemberTransactionResponse()
    err = c.Send(request, response)
    return
}

func NewRefundOrderRequest() (request *RefundOrderRequest) {
    request = &RefundOrderRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "RefundOrder")
    
    
    return
}

func NewRefundOrderResponse() (response *RefundOrderResponse) {
    response = &RefundOrderResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// RefundOrder
// 云鉴-消费订单退款的接口
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDPARAMETER = "FailedOperation.InvalidParameter"
//  FAILEDOPERATION_MISSINGPARAMETER = "FailedOperation.MissingParameter"
//  FAILEDOPERATION_ORDERREFUNDERROR = "FailedOperation.OrderRefundError"
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
func (c *Client) RefundOrder(request *RefundOrderRequest) (response *RefundOrderResponse, err error) {
    if request == nil {
        request = NewRefundOrderRequest()
    }
    
    response = NewRefundOrderResponse()
    err = c.Send(request, response)
    return
}

func NewRefundTlinxOrderRequest() (request *RefundTlinxOrderRequest) {
    request = &RefundTlinxOrderRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "RefundTlinxOrder")
    
    
    return
}

func NewRefundTlinxOrderResponse() (response *RefundTlinxOrderResponse) {
    response = &RefundTlinxOrderResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// RefundTlinxOrder
// 云支付订单退款接口
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDPARAMETER = "FailedOperation.InvalidParameter"
//  FAILEDOPERATION_MISSINGPARAMETER = "FailedOperation.MissingParameter"
//  FAILEDOPERATION_ORDERREFUNDERROR = "FailedOperation.OrderRefundError"
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
func (c *Client) RefundTlinxOrder(request *RefundTlinxOrderRequest) (response *RefundTlinxOrderResponse, err error) {
    if request == nil {
        request = NewRefundTlinxOrderRequest()
    }
    
    response = NewRefundTlinxOrderResponse()
    err = c.Send(request, response)
    return
}

func NewRegisterBehaviorRequest() (request *RegisterBehaviorRequest) {
    request = &RegisterBehaviorRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "RegisterBehavior")
    
    
    return
}

func NewRegisterBehaviorResponse() (response *RegisterBehaviorResponse) {
    response = &RegisterBehaviorResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// RegisterBehavior
// 商户查询是否签约和签约行为上报
//
// 可能返回的错误码:
//  FAILEDOPERATION_BANKFAILED = "FailedOperation.BankFailed"
//  FAILEDOPERATION_SERVICEERROR = "FailedOperation.ServiceError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER_ = "MissingParameter."
func (c *Client) RegisterBehavior(request *RegisterBehaviorRequest) (response *RegisterBehaviorResponse, err error) {
    if request == nil {
        request = NewRegisterBehaviorRequest()
    }
    
    response = NewRegisterBehaviorResponse()
    err = c.Send(request, response)
    return
}

func NewRegisterBillRequest() (request *RegisterBillRequest) {
    request = &RegisterBillRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "RegisterBill")
    
    
    return
}

func NewRegisterBillResponse() (response *RegisterBillResponse) {
    response = &RegisterBillResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// RegisterBill
// 登记挂账(支持撤销)
//
// 可能返回的错误码:
//  FAILEDOPERATION_ACCTNOTBIND = "FailedOperation.AcctNotBind"
//  FAILEDOPERATION_ACCTNOTEXIST = "FailedOperation.AcctNotExist"
//  FAILEDOPERATION_ACTIONINVALID = "FailedOperation.ActionInvalid"
//  FAILEDOPERATION_BALANCEINSUFFICIENT = "FailedOperation.BalanceInsufficient"
//  FAILEDOPERATION_PABANKERROR = "FailedOperation.PABankError"
func (c *Client) RegisterBill(request *RegisterBillRequest) (response *RegisterBillResponse, err error) {
    if request == nil {
        request = NewRegisterBillRequest()
    }
    
    response = NewRegisterBillResponse()
    err = c.Send(request, response)
    return
}

func NewRegisterBillSupportWithdrawRequest() (request *RegisterBillSupportWithdrawRequest) {
    request = &RegisterBillSupportWithdrawRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "RegisterBillSupportWithdraw")
    
    
    return
}

func NewRegisterBillSupportWithdrawResponse() (response *RegisterBillSupportWithdrawResponse) {
    response = &RegisterBillSupportWithdrawResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// RegisterBillSupportWithdraw
// 登记挂账(支持撤销)。此接口可实现把不明来账或自有资金等已登记在挂账子账户下的资金调整到普通会员子账户。即通过申请调用此接口，将会减少挂账子账户的资金，调增指定的普通会员子账户的可提现余额及可用余额。此接口不支持把挂账子账户资金清分到功能子账户。
//
// 可能返回的错误码:
//  FAILEDOPERATION_BACKENDERROR = "FailedOperation.BackendError"
//  INTERNALERROR_BACKENDERROR = "InternalError.BackendError"
//  INTERNALERROR_DBACCESSERROR = "InternalError.DBAccessError"
//  INTERNALERROR_FUNDSUMMARYACCTNOINCONSISTENTERROR = "InternalError.FundSummaryAcctNoInconsistentError"
//  INTERNALERROR_PARAMETERERROR = "InternalError.ParameterError"
//  INTERNALERROR_SAVEDBERROR = "InternalError.SaveDBError"
//  INTERNALERROR_SUBACCOUNTNOTFOUNDERROR = "InternalError.SubAccountNotFoundError"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  INTERNALERROR_UNKOWNERROR = "InternalError.UnkownError"
func (c *Client) RegisterBillSupportWithdraw(request *RegisterBillSupportWithdrawRequest) (response *RegisterBillSupportWithdrawResponse, err error) {
    if request == nil {
        request = NewRegisterBillSupportWithdrawRequest()
    }
    
    response = NewRegisterBillSupportWithdrawResponse()
    err = c.Send(request, response)
    return
}

func NewRevResigterBillSupportWithdrawRequest() (request *RevResigterBillSupportWithdrawRequest) {
    request = &RevResigterBillSupportWithdrawRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "RevResigterBillSupportWithdraw")
    
    
    return
}

func NewRevResigterBillSupportWithdrawResponse() (response *RevResigterBillSupportWithdrawResponse) {
    response = &RevResigterBillSupportWithdrawResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// RevResigterBillSupportWithdraw
// 登记挂账撤销。此接口可以实现把RegisterBillSupportWithdraw接口完成的登记挂账进行撤销，即调减普通会员子账户的可提现和可用余额，调增挂账子账户的可用余额。
//
// 可能返回的错误码:
//  FAILEDOPERATION_BACKENDERROR = "FailedOperation.BackendError"
//  INTERNALERROR_BACKENDERROR = "InternalError.BackendError"
//  INTERNALERROR_DBACCESSERROR = "InternalError.DBAccessError"
//  INTERNALERROR_FUNDSUMMARYACCTNOINCONSISTENTERROR = "InternalError.FundSummaryAcctNoInconsistentError"
//  INTERNALERROR_PARAMETERERROR = "InternalError.ParameterError"
//  INTERNALERROR_SAVEDBERROR = "InternalError.SaveDBError"
//  INTERNALERROR_SUBACCOUNTNOTFOUNDERROR = "InternalError.SubAccountNotFoundError"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  INTERNALERROR_UNKOWNERROR = "InternalError.UnkownError"
func (c *Client) RevResigterBillSupportWithdraw(request *RevResigterBillSupportWithdrawRequest) (response *RevResigterBillSupportWithdrawResponse, err error) {
    if request == nil {
        request = NewRevResigterBillSupportWithdrawRequest()
    }
    
    response = NewRevResigterBillSupportWithdrawResponse()
    err = c.Send(request, response)
    return
}

func NewReviseMbrPropertyRequest() (request *ReviseMbrPropertyRequest) {
    request = &ReviseMbrPropertyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "ReviseMbrProperty")
    
    
    return
}

func NewReviseMbrPropertyResponse() (response *ReviseMbrPropertyResponse) {
    response = &ReviseMbrPropertyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ReviseMbrProperty
// 修改会员属性-普通商户子账户。修改会员的会员属性。
//
// 可能返回的错误码:
//  FAILEDOPERATION_BACKENDERROR = "FailedOperation.BackendError"
//  INTERNALERROR_BACKENDERROR = "InternalError.BackendError"
//  INTERNALERROR_DBACCESSERROR = "InternalError.DBAccessError"
//  INTERNALERROR_FUNDSUMMARYACCTNOINCONSISTENTERROR = "InternalError.FundSummaryAcctNoInconsistentError"
//  INTERNALERROR_PARAMETERERROR = "InternalError.ParameterError"
//  INTERNALERROR_SAVEDBERROR = "InternalError.SaveDBError"
//  INTERNALERROR_SUBACCOUNTNOTFOUNDERROR = "InternalError.SubAccountNotFoundError"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  INTERNALERROR_UNKOWNERROR = "InternalError.UnkownError"
func (c *Client) ReviseMbrProperty(request *ReviseMbrPropertyRequest) (response *ReviseMbrPropertyResponse, err error) {
    if request == nil {
        request = NewReviseMbrPropertyRequest()
    }
    
    response = NewReviseMbrPropertyResponse()
    err = c.Send(request, response)
    return
}

func NewRevokeMemberRechargeThirdPayRequest() (request *RevokeMemberRechargeThirdPayRequest) {
    request = &RevokeMemberRechargeThirdPayRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "RevokeMemberRechargeThirdPay")
    
    
    return
}

func NewRevokeMemberRechargeThirdPayResponse() (response *RevokeMemberRechargeThirdPayResponse) {
    response = &RevokeMemberRechargeThirdPayResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// RevokeMemberRechargeThirdPay
// 撤销会员在途充值(经第三方支付渠道)
//
// 可能返回的错误码:
//  FAILEDOPERATION_BACKENDERROR = "FailedOperation.BackendError"
//  INTERNALERROR_BACKENDERROR = "InternalError.BackendError"
//  INTERNALERROR_DBACCESSERROR = "InternalError.DBAccessError"
//  INTERNALERROR_FUNDSUMMARYACCTNOINCONSISTENTERROR = "InternalError.FundSummaryAcctNoInconsistentError"
//  INTERNALERROR_PARAMETERERROR = "InternalError.ParameterError"
//  INTERNALERROR_SAVEDBERROR = "InternalError.SaveDBError"
//  INTERNALERROR_SUBACCOUNTNOTFOUNDERROR = "InternalError.SubAccountNotFoundError"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  INTERNALERROR_UNKOWNERROR = "InternalError.UnkownError"
func (c *Client) RevokeMemberRechargeThirdPay(request *RevokeMemberRechargeThirdPayRequest) (response *RevokeMemberRechargeThirdPayResponse, err error) {
    if request == nil {
        request = NewRevokeMemberRechargeThirdPayRequest()
    }
    
    response = NewRevokeMemberRechargeThirdPayResponse()
    err = c.Send(request, response)
    return
}

func NewRevokeRechargeByThirdPayRequest() (request *RevokeRechargeByThirdPayRequest) {
    request = &RevokeRechargeByThirdPayRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "RevokeRechargeByThirdPay")
    
    
    return
}

func NewRevokeRechargeByThirdPayResponse() (response *RevokeRechargeByThirdPayResponse) {
    response = &RevokeRechargeByThirdPayResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// RevokeRechargeByThirdPay
// 撤销会员在途充值(经第三方支付渠道)接口
//
// 可能返回的错误码:
//  FAILEDOPERATION_ACTIONINVALID = "FailedOperation.ActionInvalid"
//  FAILEDOPERATION_PABANKERROR = "FailedOperation.PABankError"
func (c *Client) RevokeRechargeByThirdPay(request *RevokeRechargeByThirdPayRequest) (response *RevokeRechargeByThirdPayResponse, err error) {
    if request == nil {
        request = NewRevokeRechargeByThirdPayRequest()
    }
    
    response = NewRevokeRechargeByThirdPayResponse()
    err = c.Send(request, response)
    return
}

func NewSyncContractDataRequest() (request *SyncContractDataRequest) {
    request = &SyncContractDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "SyncContractData")
    
    
    return
}

func NewSyncContractDataResponse() (response *SyncContractDataResponse) {
    response = &SyncContractDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// SyncContractData
// 对于存量的签约关系导入或者部分场景下米大师无法收到签约通知的场景，需要由调用方主动将签约状态同步至米大师
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_BACKCALLERROR = "FailedOperation.BackCallError"
//  FAILEDOPERATION_CALLCHANNELGATEWAYERROR = "FailedOperation.CallChannelGatewayError"
//  FAILEDOPERATION_DBCLIENTCONNECTFAILED = "FailedOperation.DbClientConnectFailed"
//  FAILEDOPERATION_DBCLIENTQUERYFAILED = "FailedOperation.DbClientQueryFailed"
//  FAILEDOPERATION_DBCLIENTUPDATEFAILED = "FailedOperation.DbClientUpdateFailed"
//  FAILEDOPERATION_EXTERNALCONTRACTNOTFOUND = "FailedOperation.ExternalContractNotFound"
//  FAILEDOPERATION_EXTERNALMERCHANTCONTRACTINFOCONFIGNOFOUND = "FailedOperation.ExternalMerchantContractInfoConfigNoFound"
//  FAILEDOPERATION_EXTERNALMERCHANTINDEXCONFIGNOFOUND = "FailedOperation.ExternalMerchantIndexConfigNoFound"
//  FAILEDOPERATION_EXTERNALMERCHANTINFOCONFIGNOFOUND = "FailedOperation.ExternalMerchantInfoConfigNoFound"
//  FAILEDOPERATION_HTTPDOREQUESTERROR = "FailedOperation.HttpDoRequestError"
//  FAILEDOPERATION_MARSHALERROR = "FailedOperation.MarshalError"
//  FAILEDOPERATION_MERCHANTPERMISSIONERROR = "FailedOperation.MerchantPermissionError"
//  FAILEDOPERATION_NOTIFYURLPARSEERROR = "FailedOperation.NotifyUrlParseError"
//  FAILEDOPERATION_PARAMETERERROR = "FailedOperation.ParameterError"
//  FAILEDOPERATION_SIGNERROR = "FailedOperation.SignError"
//  FAILEDOPERATION_SYNCMCHANNELERROR = "FailedOperation.SyncMchannelError"
//  FAILEDOPERATION_UNMARSHALERROR = "FailedOperation.UnmarshalError"
//  FAILEDOPERATION_XMLFAIL = "FailedOperation.XmlFail"
//  INTERNALERROR_BACKENDCONNECTIONERROR = "InternalError.BackendConnectionError"
//  INTERNALERROR_BACKENDINTERNALERROR = "InternalError.BackendInternalError"
//  INTERNALERROR_BACKENDNETWORKERROR = "InternalError.BackendNetworkError"
//  INTERNALERROR_BACKENDROUTERERROR = "InternalError.BackendRouterError"
//  INTERNALERROR_BACKENDTIMEOUT = "InternalError.BackendTimeOut"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMMARSHALFAILED = "InvalidParameter.ParamMarshalFailed"
//  INVALIDPARAMETER_PARAMUNMARSHALFAILED = "InvalidParameter.ParamUnmarshalFailed"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_ = "MissingParameter."
func (c *Client) SyncContractData(request *SyncContractDataRequest) (response *SyncContractDataResponse, err error) {
    if request == nil {
        request = NewSyncContractDataRequest()
    }
    
    response = NewSyncContractDataResponse()
    err = c.Send(request, response)
    return
}

func NewTerminateContractRequest() (request *TerminateContractRequest) {
    request = &TerminateContractRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "TerminateContract")
    
    
    return
}

func NewTerminateContractResponse() (response *TerminateContractResponse) {
    response = &TerminateContractResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// TerminateContract
// 通过此接口进行解约
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_BACKCALLERROR = "FailedOperation.BackCallError"
//  FAILEDOPERATION_CALLCHANNELGATEWAYERROR = "FailedOperation.CallChannelGatewayError"
//  FAILEDOPERATION_CLOSECONTRACTDBFAILED = "FailedOperation.CloseContractDbFailed"
//  FAILEDOPERATION_CLOSECONTRACTMODEINVALID = "FailedOperation.CloseContractModeInvalid"
//  FAILEDOPERATION_CONTRACTSTATUSERROR = "FailedOperation.ContractStatusError"
//  FAILEDOPERATION_DBCLIENTCONNECTFAILED = "FailedOperation.DbClientConnectFailed"
//  FAILEDOPERATION_DBCLIENTQUERYFAILED = "FailedOperation.DbClientQueryFailed"
//  FAILEDOPERATION_DBCLIENTUPDATEFAILED = "FailedOperation.DbClientUpdateFailed"
//  FAILEDOPERATION_EXTERNALCONTRACTNOTFOUND = "FailedOperation.ExternalContractNotFound"
//  FAILEDOPERATION_EXTERNALMERCHANTCONTRACTINFOCONFIGNOFOUND = "FailedOperation.ExternalMerchantContractInfoConfigNoFound"
//  FAILEDOPERATION_EXTERNALMERCHANTINDEXCONFIGNOFOUND = "FailedOperation.ExternalMerchantIndexConfigNoFound"
//  FAILEDOPERATION_EXTERNALMERCHANTINFOCONFIGNOFOUND = "FailedOperation.ExternalMerchantInfoConfigNoFound"
//  FAILEDOPERATION_HTTPDOREQUESTERROR = "FailedOperation.HttpDoRequestError"
//  FAILEDOPERATION_MARSHALERROR = "FailedOperation.MarshalError"
//  FAILEDOPERATION_MERCHANTPERMISSIONERROR = "FailedOperation.MerchantPermissionError"
//  FAILEDOPERATION_NOTIFYURLPARSEERROR = "FailedOperation.NotifyUrlParseError"
//  FAILEDOPERATION_PARAMETERERROR = "FailedOperation.ParameterError"
//  FAILEDOPERATION_SIGNERROR = "FailedOperation.SignError"
//  FAILEDOPERATION_TERMINATEMCHANNELERROR = "FailedOperation.TerminateMchannelError"
//  FAILEDOPERATION_UNMARSHALERROR = "FailedOperation.UnmarshalError"
//  FAILEDOPERATION_XMLFAIL = "FailedOperation.XmlFail"
//  INTERNALERROR_BACKENDCONNECTIONERROR = "InternalError.BackendConnectionError"
//  INTERNALERROR_BACKENDINTERNALERROR = "InternalError.BackendInternalError"
//  INTERNALERROR_BACKENDNETWORKERROR = "InternalError.BackendNetworkError"
//  INTERNALERROR_BACKENDROUTERERROR = "InternalError.BackendRouterError"
//  INTERNALERROR_BACKENDTIMEOUT = "InternalError.BackendTimeOut"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMMARSHALFAILED = "InvalidParameter.ParamMarshalFailed"
//  INVALIDPARAMETER_PARAMUNMARSHALFAILED = "InvalidParameter.ParamUnmarshalFailed"
//  MISSINGPARAMETER_ = "MissingParameter."
func (c *Client) TerminateContract(request *TerminateContractRequest) (response *TerminateContractResponse, err error) {
    if request == nil {
        request = NewTerminateContractRequest()
    }
    
    response = NewTerminateContractResponse()
    err = c.Send(request, response)
    return
}

func NewTransferSinglePayRequest() (request *TransferSinglePayRequest) {
    request = &TransferSinglePayRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "TransferSinglePay")
    
    
    return
}

func NewTransferSinglePayResponse() (response *TransferSinglePayResponse) {
    response = &TransferSinglePayResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// TransferSinglePay
// 智能代发-单笔代发转账接口
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_BACKCALLERROR = "FailedOperation.BackCallError"
//  FAILEDOPERATION_CALLCHANNELGATEWAYERROR = "FailedOperation.CallChannelGatewayError"
//  FAILEDOPERATION_CLOSECONTRACTDBFAILED = "FailedOperation.CloseContractDbFailed"
//  FAILEDOPERATION_CLOSECONTRACTMODEINVALID = "FailedOperation.CloseContractModeInvalid"
//  FAILEDOPERATION_CONTRACTSTATUSERROR = "FailedOperation.ContractStatusError"
//  FAILEDOPERATION_DBCLIENTCONNECTFAILED = "FailedOperation.DbClientConnectFailed"
//  FAILEDOPERATION_DBCLIENTQUERYFAILED = "FailedOperation.DbClientQueryFailed"
//  FAILEDOPERATION_DBCLIENTUPDATEFAILED = "FailedOperation.DbClientUpdateFailed"
//  FAILEDOPERATION_EXTERNALCONTRACTNOTFOUND = "FailedOperation.ExternalContractNotFound"
//  FAILEDOPERATION_EXTERNALMERCHANTCONTRACTINFOCONFIGNOFOUND = "FailedOperation.ExternalMerchantContractInfoConfigNoFound"
//  FAILEDOPERATION_EXTERNALMERCHANTINDEXCONFIGNOFOUND = "FailedOperation.ExternalMerchantIndexConfigNoFound"
//  FAILEDOPERATION_EXTERNALMERCHANTINFOCONFIGNOFOUND = "FailedOperation.ExternalMerchantInfoConfigNoFound"
//  FAILEDOPERATION_HTTPDOREQUESTERROR = "FailedOperation.HttpDoRequestError"
//  FAILEDOPERATION_MARSHALERROR = "FailedOperation.MarshalError"
//  FAILEDOPERATION_MERCHANTPERMISSIONERROR = "FailedOperation.MerchantPermissionError"
//  FAILEDOPERATION_NOTIFYURLPARSEERROR = "FailedOperation.NotifyUrlParseError"
//  FAILEDOPERATION_PARAMETERERROR = "FailedOperation.ParameterError"
//  FAILEDOPERATION_SIGNERROR = "FailedOperation.SignError"
//  FAILEDOPERATION_TERMINATEMCHANNELERROR = "FailedOperation.TerminateMchannelError"
//  FAILEDOPERATION_UNMARSHALERROR = "FailedOperation.UnmarshalError"
//  FAILEDOPERATION_XMLFAIL = "FailedOperation.XmlFail"
//  INTERNALERROR_BACKENDCONNECTIONERROR = "InternalError.BackendConnectionError"
//  INTERNALERROR_BACKENDINTERNALERROR = "InternalError.BackendInternalError"
//  INTERNALERROR_BACKENDNETWORKERROR = "InternalError.BackendNetworkError"
//  INTERNALERROR_BACKENDROUTERERROR = "InternalError.BackendRouterError"
//  INTERNALERROR_BACKENDTIMEOUT = "InternalError.BackendTimeOut"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMMARSHALFAILED = "InvalidParameter.ParamMarshalFailed"
//  INVALIDPARAMETER_PARAMUNMARSHALFAILED = "InvalidParameter.ParamUnmarshalFailed"
//  MISSINGPARAMETER_ = "MissingParameter."
func (c *Client) TransferSinglePay(request *TransferSinglePayRequest) (response *TransferSinglePayResponse, err error) {
    if request == nil {
        request = NewTransferSinglePayRequest()
    }
    
    response = NewTransferSinglePayResponse()
    err = c.Send(request, response)
    return
}

func NewUnBindAcctRequest() (request *UnBindAcctRequest) {
    request = &UnBindAcctRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "UnBindAcct")
    
    
    return
}

func NewUnBindAcctResponse() (response *UnBindAcctResponse) {
    response = &UnBindAcctResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UnBindAcct
// 商户解除绑定的提现银行卡
//
// 可能返回的错误码:
//  FAILEDOPERATION_BANKFAILED = "FailedOperation.BankFailed"
//  FAILEDOPERATION_SERVICEERROR = "FailedOperation.ServiceError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER_ = "MissingParameter."
func (c *Client) UnBindAcct(request *UnBindAcctRequest) (response *UnBindAcctResponse, err error) {
    if request == nil {
        request = NewUnBindAcctRequest()
    }
    
    response = NewUnBindAcctResponse()
    err = c.Send(request, response)
    return
}

func NewUnbindRelateAcctRequest() (request *UnbindRelateAcctRequest) {
    request = &UnbindRelateAcctRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "UnbindRelateAcct")
    
    
    return
}

func NewUnbindRelateAcctResponse() (response *UnbindRelateAcctResponse) {
    response = &UnbindRelateAcctResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UnbindRelateAcct
// 会员解绑提现账户。此接口可以支持会员解除名下的绑定账户关系。
//
// 可能返回的错误码:
//  FAILEDOPERATION_BACKENDERROR = "FailedOperation.BackendError"
//  INTERNALERROR_BACKENDERROR = "InternalError.BackendError"
//  INTERNALERROR_DBACCESSERROR = "InternalError.DBAccessError"
//  INTERNALERROR_FUNDSUMMARYACCTNOINCONSISTENTERROR = "InternalError.FundSummaryAcctNoInconsistentError"
//  INTERNALERROR_PARAMETERERROR = "InternalError.ParameterError"
//  INTERNALERROR_SAVEDBERROR = "InternalError.SaveDBError"
//  INTERNALERROR_SUBACCOUNTNOTFOUNDERROR = "InternalError.SubAccountNotFoundError"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  INTERNALERROR_UNKOWNERROR = "InternalError.UnkownError"
func (c *Client) UnbindRelateAcct(request *UnbindRelateAcctRequest) (response *UnbindRelateAcctResponse, err error) {
    if request == nil {
        request = NewUnbindRelateAcctRequest()
    }
    
    response = NewUnbindRelateAcctResponse()
    err = c.Send(request, response)
    return
}

func NewUnifiedOrderRequest() (request *UnifiedOrderRequest) {
    request = &UnifiedOrderRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "UnifiedOrder")
    
    
    return
}

func NewUnifiedOrderResponse() (response *UnifiedOrderResponse) {
    response = &UnifiedOrderResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UnifiedOrder
// 应用需要先调用本接口生成支付订单号，并将应答的PayInfo透传给聚鑫SDK，拉起客户端（包括微信公众号/微信小程序/客户端App）支付。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_APPDENY = "FailedOperation.AppDeny"
//  FAILEDOPERATION_CHANNELDENY = "FailedOperation.ChannelDeny"
//  FAILEDOPERATION_OCCOMPLETEDORDER = "FailedOperation.OcCompletedOrder"
//  FAILEDOPERATION_OCREPEATORDER = "FailedOperation.OcRepeatOrder"
//  FAILEDOPERATION_PARENTAPPIDERROR = "FailedOperation.ParentAppIdError"
//  FAILEDOPERATION_PORTALERROR = "FailedOperation.PortalError"
//  FAILEDOPERATION_WECHATERROR = "FailedOperation.WechatError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) UnifiedOrder(request *UnifiedOrderRequest) (response *UnifiedOrderResponse, err error) {
    if request == nil {
        request = NewUnifiedOrderRequest()
    }
    
    response = NewUnifiedOrderResponse()
    err = c.Send(request, response)
    return
}

func NewUnifiedTlinxOrderRequest() (request *UnifiedTlinxOrderRequest) {
    request = &UnifiedTlinxOrderRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "UnifiedTlinxOrder")
    
    
    return
}

func NewUnifiedTlinxOrderResponse() (response *UnifiedTlinxOrderResponse) {
    response = &UnifiedTlinxOrderResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UnifiedTlinxOrder
// 云支付-统一下单接口
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_APPDENY = "FailedOperation.AppDeny"
//  FAILEDOPERATION_CHANNELDENY = "FailedOperation.ChannelDeny"
//  FAILEDOPERATION_OCCOMPLETEDORDER = "FailedOperation.OcCompletedOrder"
//  FAILEDOPERATION_OCREPEATORDER = "FailedOperation.OcRepeatOrder"
//  FAILEDOPERATION_PARENTAPPIDERROR = "FailedOperation.ParentAppIdError"
//  FAILEDOPERATION_PORTALERROR = "FailedOperation.PortalError"
//  FAILEDOPERATION_WECHATERROR = "FailedOperation.WechatError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) UnifiedTlinxOrder(request *UnifiedTlinxOrderRequest) (response *UnifiedTlinxOrderResponse, err error) {
    if request == nil {
        request = NewUnifiedTlinxOrderRequest()
    }
    
    response = NewUnifiedTlinxOrderResponse()
    err = c.Send(request, response)
    return
}

func NewUploadExternalAnchorInfoRequest() (request *UploadExternalAnchorInfoRequest) {
    request = &UploadExternalAnchorInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "UploadExternalAnchorInfo")
    
    
    return
}

func NewUploadExternalAnchorInfoResponse() (response *UploadExternalAnchorInfoResponse) {
    response = &UploadExternalAnchorInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UploadExternalAnchorInfo
// 灵云-上传主播信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_APPDENY = "FailedOperation.AppDeny"
//  FAILEDOPERATION_CHANNELDENY = "FailedOperation.ChannelDeny"
//  FAILEDOPERATION_OCCOMPLETEDORDER = "FailedOperation.OcCompletedOrder"
//  FAILEDOPERATION_OCREPEATORDER = "FailedOperation.OcRepeatOrder"
//  FAILEDOPERATION_PARENTAPPIDERROR = "FailedOperation.ParentAppIdError"
//  FAILEDOPERATION_PORTALERROR = "FailedOperation.PortalError"
//  FAILEDOPERATION_WECHATERROR = "FailedOperation.WechatError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) UploadExternalAnchorInfo(request *UploadExternalAnchorInfoRequest) (response *UploadExternalAnchorInfoResponse, err error) {
    if request == nil {
        request = NewUploadExternalAnchorInfoRequest()
    }
    
    response = NewUploadExternalAnchorInfoResponse()
    err = c.Send(request, response)
    return
}

func NewUploadFileRequest() (request *UploadFileRequest) {
    request = &UploadFileRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "UploadFile")
    
    
    return
}

func NewUploadFileResponse() (response *UploadFileResponse) {
    response = &UploadFileResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UploadFile
// 直播平台-文件上传
//
// 可能返回的错误码:
//  FAILEDOPERATION_ACTION = "FailedOperation.Action"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  MISSINGPARAMETER_ACTION = "MissingParameter.Action"
//  MISSINGPARAMETER_APPID = "MissingParameter.AppId"
func (c *Client) UploadFile(request *UploadFileRequest) (response *UploadFileResponse, err error) {
    if request == nil {
        request = NewUploadFileRequest()
    }
    
    response = NewUploadFileResponse()
    err = c.Send(request, response)
    return
}

func NewUploadOrgFileRequest() (request *UploadOrgFileRequest) {
    request = &UploadOrgFileRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "UploadOrgFile")
    
    
    return
}

func NewUploadOrgFileResponse() (response *UploadOrgFileResponse) {
    response = &UploadOrgFileResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UploadOrgFile
// 云支付-上传机构文件接口
//
// 可能返回的错误码:
//  FAILEDOPERATION_ACTION = "FailedOperation.Action"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  MISSINGPARAMETER_ACTION = "MissingParameter.Action"
//  MISSINGPARAMETER_APPID = "MissingParameter.AppId"
func (c *Client) UploadOrgFile(request *UploadOrgFileRequest) (response *UploadOrgFileResponse, err error) {
    if request == nil {
        request = NewUploadOrgFileRequest()
    }
    
    response = NewUploadOrgFileResponse()
    err = c.Send(request, response)
    return
}

func NewUploadTaxListRequest() (request *UploadTaxListRequest) {
    request = &UploadTaxListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "UploadTaxList")
    
    
    return
}

func NewUploadTaxListResponse() (response *UploadTaxListResponse) {
    response = &UploadTaxListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UploadTaxList
// 直播平台-上传代理商完税列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_UPLOADTAXLIST = "FailedOperation.UploadTaxList"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  MISSINGPARAMETER_ACTION = "MissingParameter.Action"
//  MISSINGPARAMETER_APPID = "MissingParameter.AppId"
//  RESOURCENOTFOUND_ACCOUNT = "ResourceNotFound.Account"
func (c *Client) UploadTaxList(request *UploadTaxListRequest) (response *UploadTaxListResponse, err error) {
    if request == nil {
        request = NewUploadTaxListRequest()
    }
    
    response = NewUploadTaxListResponse()
    err = c.Send(request, response)
    return
}

func NewUploadTaxPaymentRequest() (request *UploadTaxPaymentRequest) {
    request = &UploadTaxPaymentRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "UploadTaxPayment")
    
    
    return
}

func NewUploadTaxPaymentResponse() (response *UploadTaxPaymentResponse) {
    response = &UploadTaxPaymentResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UploadTaxPayment
// 直播平台-上传代理商完税证明
//
// 可能返回的错误码:
//  FAILEDOPERATION_UPLOADTAXLIST = "FailedOperation.UploadTaxList"
//  FAILEDOPERATION_UPLOADTAXPAYMENT = "FailedOperation.UploadTaxPayment"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  MISSINGPARAMETER_ACTION = "MissingParameter.Action"
//  MISSINGPARAMETER_APPID = "MissingParameter.AppId"
//  RESOURCENOTFOUND_ACCOUNT = "ResourceNotFound.Account"
func (c *Client) UploadTaxPayment(request *UploadTaxPaymentRequest) (response *UploadTaxPaymentResponse, err error) {
    if request == nil {
        request = NewUploadTaxPaymentRequest()
    }
    
    response = NewUploadTaxPaymentResponse()
    err = c.Send(request, response)
    return
}

func NewViewContractRequest() (request *ViewContractRequest) {
    request = &ViewContractRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "ViewContract")
    
    
    return
}

func NewViewContractResponse() (response *ViewContractResponse) {
    response = &ViewContractResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ViewContract
// 云支付-查询合同明细接口
//
// 可能返回的错误码:
//  FAILEDOPERATION_UPLOADTAXLIST = "FailedOperation.UploadTaxList"
//  FAILEDOPERATION_UPLOADTAXPAYMENT = "FailedOperation.UploadTaxPayment"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  MISSINGPARAMETER_ACTION = "MissingParameter.Action"
//  MISSINGPARAMETER_APPID = "MissingParameter.AppId"
//  RESOURCENOTFOUND_ACCOUNT = "ResourceNotFound.Account"
func (c *Client) ViewContract(request *ViewContractRequest) (response *ViewContractResponse, err error) {
    if request == nil {
        request = NewViewContractRequest()
    }
    
    response = NewViewContractResponse()
    err = c.Send(request, response)
    return
}

func NewViewMerchantRequest() (request *ViewMerchantRequest) {
    request = &ViewMerchantRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "ViewMerchant")
    
    
    return
}

func NewViewMerchantResponse() (response *ViewMerchantResponse) {
    response = &ViewMerchantResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ViewMerchant
// 云支付-查询商户明细接口
//
// 可能返回的错误码:
//  FAILEDOPERATION_UPLOADTAXLIST = "FailedOperation.UploadTaxList"
//  FAILEDOPERATION_UPLOADTAXPAYMENT = "FailedOperation.UploadTaxPayment"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  MISSINGPARAMETER_ACTION = "MissingParameter.Action"
//  MISSINGPARAMETER_APPID = "MissingParameter.AppId"
//  RESOURCENOTFOUND_ACCOUNT = "ResourceNotFound.Account"
func (c *Client) ViewMerchant(request *ViewMerchantRequest) (response *ViewMerchantResponse, err error) {
    if request == nil {
        request = NewViewMerchantRequest()
    }
    
    response = NewViewMerchantResponse()
    err = c.Send(request, response)
    return
}

func NewViewShopRequest() (request *ViewShopRequest) {
    request = &ViewShopRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "ViewShop")
    
    
    return
}

func NewViewShopResponse() (response *ViewShopResponse) {
    response = &ViewShopResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ViewShop
// 云支付-查询门店明细接口
//
// 可能返回的错误码:
//  FAILEDOPERATION_UPLOADTAXLIST = "FailedOperation.UploadTaxList"
//  FAILEDOPERATION_UPLOADTAXPAYMENT = "FailedOperation.UploadTaxPayment"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  MISSINGPARAMETER_ACTION = "MissingParameter.Action"
//  MISSINGPARAMETER_APPID = "MissingParameter.AppId"
//  RESOURCENOTFOUND_ACCOUNT = "ResourceNotFound.Account"
func (c *Client) ViewShop(request *ViewShopRequest) (response *ViewShopResponse, err error) {
    if request == nil {
        request = NewViewShopRequest()
    }
    
    response = NewViewShopResponse()
    err = c.Send(request, response)
    return
}

func NewWithdrawCashMembershipRequest() (request *WithdrawCashMembershipRequest) {
    request = &WithdrawCashMembershipRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "WithdrawCashMembership")
    
    
    return
}

func NewWithdrawCashMembershipResponse() (response *WithdrawCashMembershipResponse) {
    response = &WithdrawCashMembershipResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// WithdrawCashMembership
// 会员提现-不验证。此接口受理会员发起的提现申请。会员子账户的可提现余额、可用余额会减少，市场的资金汇总账户(监管账户)会减少相应的发生金额，提现到会员申请的收款账户。		
//
// 可能返回的错误码:
//  FAILEDOPERATION_BACKENDERROR = "FailedOperation.BackendError"
//  INTERNALERROR_BACKENDERROR = "InternalError.BackendError"
//  INTERNALERROR_DBACCESSERROR = "InternalError.DBAccessError"
//  INTERNALERROR_FUNDSUMMARYACCTNOINCONSISTENTERROR = "InternalError.FundSummaryAcctNoInconsistentError"
//  INTERNALERROR_PARAMETERERROR = "InternalError.ParameterError"
//  INTERNALERROR_SAVEDBERROR = "InternalError.SaveDBError"
//  INTERNALERROR_SUBACCOUNTNOTFOUNDERROR = "InternalError.SubAccountNotFoundError"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  INTERNALERROR_UNKOWNERROR = "InternalError.UnkownError"
func (c *Client) WithdrawCashMembership(request *WithdrawCashMembershipRequest) (response *WithdrawCashMembershipResponse, err error) {
    if request == nil {
        request = NewWithdrawCashMembershipRequest()
    }
    
    response = NewWithdrawCashMembershipResponse()
    err = c.Send(request, response)
    return
}
