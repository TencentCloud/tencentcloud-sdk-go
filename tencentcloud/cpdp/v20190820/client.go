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
    "context"
    "errors"
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
    return c.AddContractWithContext(context.Background(), request)
}

// AddContract
// 云支付-添加合同接口
func (c *Client) AddContractWithContext(ctx context.Context, request *AddContractRequest) (response *AddContractResponse, err error) {
    if request == nil {
        request = NewAddContractRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddContract require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddContractResponse()
    err = c.Send(request, response)
    return
}

func NewAddFlexFundingAccountRequest() (request *AddFlexFundingAccountRequest) {
    request = &AddFlexFundingAccountRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cpdp", APIVersion, "AddFlexFundingAccount")
    
    
    return
}

func NewAddFlexFundingAccountResponse() (response *AddFlexFundingAccountResponse) {
    response = &AddFlexFundingAccountResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AddFlexFundingAccount
// 灵云V2-绑定收款用户资金账号信息
//
// 可能返回的错误码:
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
func (c *Client) AddFlexFundingAccount(request *AddFlexFundingAccountRequest) (response *AddFlexFundingAccountResponse, err error) {
    return c.AddFlexFundingAccountWithContext(context.Background(), request)
}

// AddFlexFundingAccount
// 灵云V2-绑定收款用户资金账号信息
//
// 可能返回的错误码:
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
func (c *Client) AddFlexFundingAccountWithContext(ctx context.Context, request *AddFlexFundingAccountRequest) (response *AddFlexFundingAccountResponse, err error) {
    if request == nil {
        request = NewAddFlexFundingAccountRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddFlexFundingAccount require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddFlexFundingAccountResponse()
    err = c.Send(request, response)
    return
}

func NewAddFlexIdInfoRequest() (request *AddFlexIdInfoRequest) {
    request = &AddFlexIdInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cpdp", APIVersion, "AddFlexIdInfo")
    
    
    return
}

func NewAddFlexIdInfoResponse() (response *AddFlexIdInfoResponse) {
    response = &AddFlexIdInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AddFlexIdInfo
// 灵云V2-补充证件信息
//
// 可能返回的错误码:
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
func (c *Client) AddFlexIdInfo(request *AddFlexIdInfoRequest) (response *AddFlexIdInfoResponse, err error) {
    return c.AddFlexIdInfoWithContext(context.Background(), request)
}

// AddFlexIdInfo
// 灵云V2-补充证件信息
//
// 可能返回的错误码:
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
func (c *Client) AddFlexIdInfoWithContext(ctx context.Context, request *AddFlexIdInfoRequest) (response *AddFlexIdInfoResponse, err error) {
    if request == nil {
        request = NewAddFlexIdInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddFlexIdInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddFlexIdInfoResponse()
    err = c.Send(request, response)
    return
}

func NewAddFlexPhoneNoRequest() (request *AddFlexPhoneNoRequest) {
    request = &AddFlexPhoneNoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cpdp", APIVersion, "AddFlexPhoneNo")
    
    
    return
}

func NewAddFlexPhoneNoResponse() (response *AddFlexPhoneNoResponse) {
    response = &AddFlexPhoneNoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AddFlexPhoneNo
// 灵云V2-补充手机号信息
//
// 可能返回的错误码:
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
func (c *Client) AddFlexPhoneNo(request *AddFlexPhoneNoRequest) (response *AddFlexPhoneNoResponse, err error) {
    return c.AddFlexPhoneNoWithContext(context.Background(), request)
}

// AddFlexPhoneNo
// 灵云V2-补充手机号信息
//
// 可能返回的错误码:
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
func (c *Client) AddFlexPhoneNoWithContext(ctx context.Context, request *AddFlexPhoneNoRequest) (response *AddFlexPhoneNoResponse, err error) {
    if request == nil {
        request = NewAddFlexPhoneNoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddFlexPhoneNo require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddFlexPhoneNoResponse()
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
//
// 可能返回的错误码:
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
func (c *Client) AddMerchant(request *AddMerchantRequest) (response *AddMerchantResponse, err error) {
    return c.AddMerchantWithContext(context.Background(), request)
}

// AddMerchant
// 云支付-添加商户接口
//
// 可能返回的错误码:
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
func (c *Client) AddMerchantWithContext(ctx context.Context, request *AddMerchantRequest) (response *AddMerchantResponse, err error) {
    if request == nil {
        request = NewAddMerchantRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddMerchant require credential")
    }

    request.SetContext(ctx)
    
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
//
// 可能返回的错误码:
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
func (c *Client) AddShop(request *AddShopRequest) (response *AddShopResponse, err error) {
    return c.AddShopWithContext(context.Background(), request)
}

// AddShop
// 云支付-添加门店接口
//
// 可能返回的错误码:
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
func (c *Client) AddShopWithContext(ctx context.Context, request *AddShopRequest) (response *AddShopResponse, err error) {
    if request == nil {
        request = NewAddShopRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddShop require credential")
    }

    request.SetContext(ctx)
    
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
    return c.ApplyApplicationMaterialWithContext(context.Background(), request)
}

// ApplyApplicationMaterial
// 跨境-提交申报材料。申报材料的主体是付款人，需要提前调用【跨境-付款人申请】接口提交付款人信息且审核通过后调用。
//
// 可能返回的错误码:
//  INTERNALERROR_PARAMETERERROR = "InternalError.ParameterError"
//  INTERNALERROR_UNKOWNERROR = "InternalError.UnkownError"
//  RESOURCENOTFOUND_MERCHANTINFONOTFOUND = "ResourceNotFound.MerchantInfoNotFound"
func (c *Client) ApplyApplicationMaterialWithContext(ctx context.Context, request *ApplyApplicationMaterialRequest) (response *ApplyApplicationMaterialResponse, err error) {
    if request == nil {
        request = NewApplyApplicationMaterialRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ApplyApplicationMaterial require credential")
    }

    request.SetContext(ctx)
    
    response = NewApplyApplicationMaterialResponse()
    err = c.Send(request, response)
    return
}

func NewApplyFlexPaymentRequest() (request *ApplyFlexPaymentRequest) {
    request = &ApplyFlexPaymentRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cpdp", APIVersion, "ApplyFlexPayment")
    
    
    return
}

func NewApplyFlexPaymentResponse() (response *ApplyFlexPaymentResponse) {
    response = &ApplyFlexPaymentResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ApplyFlexPayment
// 灵云V2-付款
//
// 可能返回的错误码:
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
func (c *Client) ApplyFlexPayment(request *ApplyFlexPaymentRequest) (response *ApplyFlexPaymentResponse, err error) {
    return c.ApplyFlexPaymentWithContext(context.Background(), request)
}

// ApplyFlexPayment
// 灵云V2-付款
//
// 可能返回的错误码:
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
func (c *Client) ApplyFlexPaymentWithContext(ctx context.Context, request *ApplyFlexPaymentRequest) (response *ApplyFlexPaymentResponse, err error) {
    if request == nil {
        request = NewApplyFlexPaymentRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ApplyFlexPayment require credential")
    }

    request.SetContext(ctx)
    
    response = NewApplyFlexPaymentResponse()
    err = c.Send(request, response)
    return
}

func NewApplyFlexSettlementRequest() (request *ApplyFlexSettlementRequest) {
    request = &ApplyFlexSettlementRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cpdp", APIVersion, "ApplyFlexSettlement")
    
    
    return
}

func NewApplyFlexSettlementResponse() (response *ApplyFlexSettlementResponse) {
    response = &ApplyFlexSettlementResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ApplyFlexSettlement
// 灵云V2-结算
//
// 可能返回的错误码:
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
func (c *Client) ApplyFlexSettlement(request *ApplyFlexSettlementRequest) (response *ApplyFlexSettlementResponse, err error) {
    return c.ApplyFlexSettlementWithContext(context.Background(), request)
}

// ApplyFlexSettlement
// 灵云V2-结算
//
// 可能返回的错误码:
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
func (c *Client) ApplyFlexSettlementWithContext(ctx context.Context, request *ApplyFlexSettlementRequest) (response *ApplyFlexSettlementResponse, err error) {
    if request == nil {
        request = NewApplyFlexSettlementRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ApplyFlexSettlement require credential")
    }

    request.SetContext(ctx)
    
    response = NewApplyFlexSettlementResponse()
    err = c.Send(request, response)
    return
}

func NewApplyOpenBankOrderDetailReceiptRequest() (request *ApplyOpenBankOrderDetailReceiptRequest) {
    request = &ApplyOpenBankOrderDetailReceiptRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cpdp", APIVersion, "ApplyOpenBankOrderDetailReceipt")
    
    
    return
}

func NewApplyOpenBankOrderDetailReceiptResponse() (response *ApplyOpenBankOrderDetailReceiptResponse) {
    response = &ApplyOpenBankOrderDetailReceiptResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ApplyOpenBankOrderDetailReceipt
// 云企付-申请单笔交易回单
//
// 可能返回的错误码:
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
func (c *Client) ApplyOpenBankOrderDetailReceipt(request *ApplyOpenBankOrderDetailReceiptRequest) (response *ApplyOpenBankOrderDetailReceiptResponse, err error) {
    return c.ApplyOpenBankOrderDetailReceiptWithContext(context.Background(), request)
}

// ApplyOpenBankOrderDetailReceipt
// 云企付-申请单笔交易回单
//
// 可能返回的错误码:
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
func (c *Client) ApplyOpenBankOrderDetailReceiptWithContext(ctx context.Context, request *ApplyOpenBankOrderDetailReceiptRequest) (response *ApplyOpenBankOrderDetailReceiptResponse, err error) {
    if request == nil {
        request = NewApplyOpenBankOrderDetailReceiptRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ApplyOpenBankOrderDetailReceipt require credential")
    }

    request.SetContext(ctx)
    
    response = NewApplyOpenBankOrderDetailReceiptResponse()
    err = c.Send(request, response)
    return
}

func NewApplyOpenBankSettleOrderRequest() (request *ApplyOpenBankSettleOrderRequest) {
    request = &ApplyOpenBankSettleOrderRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cpdp", APIVersion, "ApplyOpenBankSettleOrder")
    
    
    return
}

func NewApplyOpenBankSettleOrderResponse() (response *ApplyOpenBankSettleOrderResponse) {
    response = &ApplyOpenBankSettleOrderResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ApplyOpenBankSettleOrder
// 云企付-结算申请接口
//
// 可能返回的错误码:
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
func (c *Client) ApplyOpenBankSettleOrder(request *ApplyOpenBankSettleOrderRequest) (response *ApplyOpenBankSettleOrderResponse, err error) {
    return c.ApplyOpenBankSettleOrderWithContext(context.Background(), request)
}

// ApplyOpenBankSettleOrder
// 云企付-结算申请接口
//
// 可能返回的错误码:
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
func (c *Client) ApplyOpenBankSettleOrderWithContext(ctx context.Context, request *ApplyOpenBankSettleOrderRequest) (response *ApplyOpenBankSettleOrderResponse, err error) {
    if request == nil {
        request = NewApplyOpenBankSettleOrderRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ApplyOpenBankSettleOrder require credential")
    }

    request.SetContext(ctx)
    
    response = NewApplyOpenBankSettleOrderResponse()
    err = c.Send(request, response)
    return
}

func NewApplyOpenBankSubMerchantSignOnlineRequest() (request *ApplyOpenBankSubMerchantSignOnlineRequest) {
    request = &ApplyOpenBankSubMerchantSignOnlineRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cpdp", APIVersion, "ApplyOpenBankSubMerchantSignOnline")
    
    
    return
}

func NewApplyOpenBankSubMerchantSignOnlineResponse() (response *ApplyOpenBankSubMerchantSignOnlineResponse) {
    response = &ApplyOpenBankSubMerchantSignOnlineResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ApplyOpenBankSubMerchantSignOnline
// 子商户在线签约
//
// 可能返回的错误码:
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
func (c *Client) ApplyOpenBankSubMerchantSignOnline(request *ApplyOpenBankSubMerchantSignOnlineRequest) (response *ApplyOpenBankSubMerchantSignOnlineResponse, err error) {
    return c.ApplyOpenBankSubMerchantSignOnlineWithContext(context.Background(), request)
}

// ApplyOpenBankSubMerchantSignOnline
// 子商户在线签约
//
// 可能返回的错误码:
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
func (c *Client) ApplyOpenBankSubMerchantSignOnlineWithContext(ctx context.Context, request *ApplyOpenBankSubMerchantSignOnlineRequest) (response *ApplyOpenBankSubMerchantSignOnlineResponse, err error) {
    if request == nil {
        request = NewApplyOpenBankSubMerchantSignOnlineRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ApplyOpenBankSubMerchantSignOnline require credential")
    }

    request.SetContext(ctx)
    
    response = NewApplyOpenBankSubMerchantSignOnlineResponse()
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
    return c.ApplyOutwardOrderWithContext(context.Background(), request)
}

// ApplyOutwardOrder
// 跨境-汇出指令申请。通过该接口可将对接方账户中的人民币余额汇兑成外币，再汇出至指定银行账户。
//
// 可能返回的错误码:
//  INTERNALERROR_PARAMETERERROR = "InternalError.ParameterError"
//  INTERNALERROR_UNKOWNERROR = "InternalError.UnkownError"
//  RESOURCENOTFOUND_MERCHANTINFONOTFOUND = "ResourceNotFound.MerchantInfoNotFound"
func (c *Client) ApplyOutwardOrderWithContext(ctx context.Context, request *ApplyOutwardOrderRequest) (response *ApplyOutwardOrderResponse, err error) {
    if request == nil {
        request = NewApplyOutwardOrderRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ApplyOutwardOrder require credential")
    }

    request.SetContext(ctx)
    
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
    return c.ApplyPayerInfoWithContext(context.Background(), request)
}

// ApplyPayerInfo
// 跨境-付款人申请。通过该接口提交付款人信息并进行 kyc 审核。
//
// 可能返回的错误码:
//  INTERNALERROR_PARAMETERERROR = "InternalError.ParameterError"
//  INTERNALERROR_UNKOWNERROR = "InternalError.UnkownError"
//  RESOURCENOTFOUND_MERCHANTINFONOTFOUND = "ResourceNotFound.MerchantInfoNotFound"
func (c *Client) ApplyPayerInfoWithContext(ctx context.Context, request *ApplyPayerInfoRequest) (response *ApplyPayerInfoResponse, err error) {
    if request == nil {
        request = NewApplyPayerInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ApplyPayerInfo require credential")
    }

    request.SetContext(ctx)
    
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
//  AUTHFAILURE_VERIFYERROR = "AuthFailure.VerifyError"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_NORECORD = "FailedOperation.NoRecord"
//  FAILEDOPERATION_OCCOMPLETEDORDER = "FailedOperation.OcCompletedOrder"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ApplyReWithdrawal(request *ApplyReWithdrawalRequest) (response *ApplyReWithdrawalResponse, err error) {
    return c.ApplyReWithdrawalWithContext(context.Background(), request)
}

// ApplyReWithdrawal
// 正常结算提现失败情况下，发起重新提现的请求接口
//
// 可能返回的错误码:
//  AUTHFAILURE_VERIFYERROR = "AuthFailure.VerifyError"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_NORECORD = "FailedOperation.NoRecord"
//  FAILEDOPERATION_OCCOMPLETEDORDER = "FailedOperation.OcCompletedOrder"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ApplyReWithdrawalWithContext(ctx context.Context, request *ApplyReWithdrawalRequest) (response *ApplyReWithdrawalResponse, err error) {
    if request == nil {
        request = NewApplyReWithdrawalRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ApplyReWithdrawal require credential")
    }

    request.SetContext(ctx)
    
    response = NewApplyReWithdrawalResponse()
    err = c.Send(request, response)
    return
}

func NewApplyReconciliationFileRequest() (request *ApplyReconciliationFileRequest) {
    request = &ApplyReconciliationFileRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cpdp", APIVersion, "ApplyReconciliationFile")
    
    
    return
}

func NewApplyReconciliationFileResponse() (response *ApplyReconciliationFileResponse) {
    response = &ApplyReconciliationFileResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ApplyReconciliationFile
// 聚鑫-申请对账文件
//
// 可能返回的错误码:
//  AUTHFAILURE_VERIFYERROR = "AuthFailure.VerifyError"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_NORECORD = "FailedOperation.NoRecord"
//  FAILEDOPERATION_OCCOMPLETEDORDER = "FailedOperation.OcCompletedOrder"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ApplyReconciliationFile(request *ApplyReconciliationFileRequest) (response *ApplyReconciliationFileResponse, err error) {
    return c.ApplyReconciliationFileWithContext(context.Background(), request)
}

// ApplyReconciliationFile
// 聚鑫-申请对账文件
//
// 可能返回的错误码:
//  AUTHFAILURE_VERIFYERROR = "AuthFailure.VerifyError"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_NORECORD = "FailedOperation.NoRecord"
//  FAILEDOPERATION_OCCOMPLETEDORDER = "FailedOperation.OcCompletedOrder"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ApplyReconciliationFileWithContext(ctx context.Context, request *ApplyReconciliationFileRequest) (response *ApplyReconciliationFileResponse, err error) {
    if request == nil {
        request = NewApplyReconciliationFileRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ApplyReconciliationFile require credential")
    }

    request.SetContext(ctx)
    
    response = NewApplyReconciliationFileResponse()
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
    return c.ApplyTradeWithContext(context.Background(), request)
}

// ApplyTrade
// 跨境-提交贸易材料。通过提交贸易材料接口可为对接方累计贸易额度，在额度范围内可发起汇兑汇出交易。
//
// 可能返回的错误码:
//  INTERNALERROR_PARAMETERERROR = "InternalError.ParameterError"
//  INTERNALERROR_UNKOWNERROR = "InternalError.UnkownError"
//  RESOURCENOTFOUND_MERCHANTINFONOTFOUND = "ResourceNotFound.MerchantInfoNotFound"
func (c *Client) ApplyTradeWithContext(ctx context.Context, request *ApplyTradeRequest) (response *ApplyTradeResponse, err error) {
    if request == nil {
        request = NewApplyTradeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ApplyTrade require credential")
    }

    request.SetContext(ctx)
    
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
    return c.ApplyWithdrawalWithContext(context.Background(), request)
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
func (c *Client) ApplyWithdrawalWithContext(ctx context.Context, request *ApplyWithdrawalRequest) (response *ApplyWithdrawalResponse, err error) {
    if request == nil {
        request = NewApplyWithdrawalRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ApplyWithdrawal require credential")
    }

    request.SetContext(ctx)
    
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
    return c.BindAccountWithContext(context.Background(), request)
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
func (c *Client) BindAccountWithContext(ctx context.Context, request *BindAccountRequest) (response *BindAccountResponse, err error) {
    if request == nil {
        request = NewBindAccountRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("BindAccount require credential")
    }

    request.SetContext(ctx)
    
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
    return c.BindAcctWithContext(context.Background(), request)
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
func (c *Client) BindAcctWithContext(ctx context.Context, request *BindAcctRequest) (response *BindAcctResponse, err error) {
    if request == nil {
        request = NewBindAcctRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("BindAcct require credential")
    }

    request.SetContext(ctx)
    
    response = NewBindAcctResponse()
    err = c.Send(request, response)
    return
}

func NewBindOpenBankExternalSubMerchantBankAccountRequest() (request *BindOpenBankExternalSubMerchantBankAccountRequest) {
    request = &BindOpenBankExternalSubMerchantBankAccountRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cpdp", APIVersion, "BindOpenBankExternalSubMerchantBankAccount")
    
    
    return
}

func NewBindOpenBankExternalSubMerchantBankAccountResponse() (response *BindOpenBankExternalSubMerchantBankAccountResponse) {
    response = &BindOpenBankExternalSubMerchantBankAccountResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// BindOpenBankExternalSubMerchantBankAccount
// 云企付-子商户银行卡绑定
//
// 可能返回的错误码:
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
func (c *Client) BindOpenBankExternalSubMerchantBankAccount(request *BindOpenBankExternalSubMerchantBankAccountRequest) (response *BindOpenBankExternalSubMerchantBankAccountResponse, err error) {
    return c.BindOpenBankExternalSubMerchantBankAccountWithContext(context.Background(), request)
}

// BindOpenBankExternalSubMerchantBankAccount
// 云企付-子商户银行卡绑定
//
// 可能返回的错误码:
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
func (c *Client) BindOpenBankExternalSubMerchantBankAccountWithContext(ctx context.Context, request *BindOpenBankExternalSubMerchantBankAccountRequest) (response *BindOpenBankExternalSubMerchantBankAccountResponse, err error) {
    if request == nil {
        request = NewBindOpenBankExternalSubMerchantBankAccountRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("BindOpenBankExternalSubMerchantBankAccount require credential")
    }

    request.SetContext(ctx)
    
    response = NewBindOpenBankExternalSubMerchantBankAccountResponse()
    err = c.Send(request, response)
    return
}

func NewBindOpenBankProfitSharePayeeRequest() (request *BindOpenBankProfitSharePayeeRequest) {
    request = &BindOpenBankProfitSharePayeeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cpdp", APIVersion, "BindOpenBankProfitSharePayee")
    
    
    return
}

func NewBindOpenBankProfitSharePayeeResponse() (response *BindOpenBankProfitSharePayeeResponse) {
    response = &BindOpenBankProfitSharePayeeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// BindOpenBankProfitSharePayee
// 云企付-绑定分账收款方
//
// 可能返回的错误码:
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
func (c *Client) BindOpenBankProfitSharePayee(request *BindOpenBankProfitSharePayeeRequest) (response *BindOpenBankProfitSharePayeeResponse, err error) {
    return c.BindOpenBankProfitSharePayeeWithContext(context.Background(), request)
}

// BindOpenBankProfitSharePayee
// 云企付-绑定分账收款方
//
// 可能返回的错误码:
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
func (c *Client) BindOpenBankProfitSharePayeeWithContext(ctx context.Context, request *BindOpenBankProfitSharePayeeRequest) (response *BindOpenBankProfitSharePayeeResponse, err error) {
    if request == nil {
        request = NewBindOpenBankProfitSharePayeeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("BindOpenBankProfitSharePayee require credential")
    }

    request.SetContext(ctx)
    
    response = NewBindOpenBankProfitSharePayeeResponse()
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
    return c.BindRelateAccReUnionPayWithContext(context.Background(), request)
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
func (c *Client) BindRelateAccReUnionPayWithContext(ctx context.Context, request *BindRelateAccReUnionPayRequest) (response *BindRelateAccReUnionPayResponse, err error) {
    if request == nil {
        request = NewBindRelateAccReUnionPayRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("BindRelateAccReUnionPay require credential")
    }

    request.SetContext(ctx)
    
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
    return c.BindRelateAcctSmallAmountWithContext(context.Background(), request)
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
func (c *Client) BindRelateAcctSmallAmountWithContext(ctx context.Context, request *BindRelateAcctSmallAmountRequest) (response *BindRelateAcctSmallAmountResponse, err error) {
    if request == nil {
        request = NewBindRelateAcctSmallAmountRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("BindRelateAcctSmallAmount require credential")
    }

    request.SetContext(ctx)
    
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
    return c.BindRelateAcctUnionPayWithContext(context.Background(), request)
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
func (c *Client) BindRelateAcctUnionPayWithContext(ctx context.Context, request *BindRelateAcctUnionPayRequest) (response *BindRelateAcctUnionPayResponse, err error) {
    if request == nil {
        request = NewBindRelateAcctUnionPayRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("BindRelateAcctUnionPay require credential")
    }

    request.SetContext(ctx)
    
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
    return c.CheckAcctWithContext(context.Background(), request)
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
func (c *Client) CheckAcctWithContext(ctx context.Context, request *CheckAcctRequest) (response *CheckAcctResponse, err error) {
    if request == nil {
        request = NewCheckAcctRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CheckAcct require credential")
    }

    request.SetContext(ctx)
    
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
    return c.CheckAmountWithContext(context.Background(), request)
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
func (c *Client) CheckAmountWithContext(ctx context.Context, request *CheckAmountRequest) (response *CheckAmountResponse, err error) {
    if request == nil {
        request = NewCheckAmountRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CheckAmount require credential")
    }

    request.SetContext(ctx)
    
    response = NewCheckAmountResponse()
    err = c.Send(request, response)
    return
}

func NewCloseCloudOrderRequest() (request *CloseCloudOrderRequest) {
    request = &CloseCloudOrderRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cpdp", APIVersion, "CloseCloudOrder")
    
    
    return
}

func NewCloseCloudOrderResponse() (response *CloseCloudOrderResponse) {
    response = &CloseCloudOrderResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CloseCloudOrder
// 通过此接口关闭此前已创建的订单。关闭后，用户将无法继续付款，仅能关闭创建后未支付的订单。
//
// 可能返回的错误码:
//  AUTHFAILURE_MIDAS = "AuthFailure.Midas"
//  FAILEDOPERATION_ACTION = "FailedOperation.Action"
//  FAILEDOPERATION_MIDASINTERNALERROR = "FailedOperation.MidasInternalError"
//  FAILEDOPERATION_MIDASINVALIDREQUEST = "FailedOperation.MidasInvalidRequest"
//  FAILEDOPERATION_MIDASNEEDRETRY = "FailedOperation.MidasNeedRetry"
//  FAILEDOPERATION_MIDASREGISTERUNFINISHED = "FailedOperation.MidasRegisterUnfinished"
//  FAILEDOPERATION_MIDASREPEATORDER = "FailedOperation.MidasRepeatOrder"
//  FAILEDOPERATION_MIDASRISK = "FailedOperation.MidasRisk"
//  FAILEDOPERATION_MIDASSTATUSNOTMATCH = "FailedOperation.MidasStatusNotMatch"
//  FAILEDOPERATION_MIDASUNSUPPORTEDACTION = "FailedOperation.MidasUnsupportedAction"
//  INTERNALERROR_MIDAS = "InternalError.Midas"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  INVALIDPARAMETER_MIDAS = "InvalidParameter.Midas"
//  INVALIDPARAMETER_MIDASENVIRONMENT = "InvalidParameter.MidasEnvironment"
//  INVALIDPARAMETER_MIDASEXTERNALAPP = "InvalidParameter.MidasExternalApp"
//  INVALIDPARAMETER_MIDASFILETYPE = "InvalidParameter.MidasFileType"
//  INVALIDPARAMETER_MIDASHASH = "InvalidParameter.MidasHash"
//  INVALIDPARAMETER_MIDASSIGNID = "InvalidParameter.MidasSignId"
//  LIMITEXCEEDED_MIDASLARGEFILE = "LimitExceeded.MidasLargeFile"
//  LIMITEXCEEDED_MIDASORDER = "LimitExceeded.MidasOrder"
//  LIMITEXCEEDED_MIDASORDERCANCELED = "LimitExceeded.MidasOrderCanceled"
//  LIMITEXCEEDED_MIDASORDERCLOSED = "LimitExceeded.MidasOrderClosed"
//  LIMITEXCEEDED_MIDASORDEREXPIRED = "LimitExceeded.MidasOrderExpired"
//  LIMITEXCEEDED_MIDASORDERFAILED = "LimitExceeded.MidasOrderFailed"
//  LIMITEXCEEDED_MIDASORDERPARTIALSUCCESS = "LimitExceeded.MidasOrderPartialSuccess"
//  LIMITEXCEEDED_MIDASORDERPROCESSING = "LimitExceeded.MidasOrderProcessing"
//  LIMITEXCEEDED_MIDASORDERSUCCESS = "LimitExceeded.MidasOrderSuccess"
//  REQUESTLIMITEXCEEDED_MIDAS = "RequestLimitExceeded.Midas"
//  REQUESTLIMITEXCEEDED_MIDASINVALIDREQUEST = "RequestLimitExceeded.MidasInvalidRequest"
//  RESOURCEINUSE_MIDAS = "ResourceInUse.Midas"
//  RESOURCENOTFOUND_ACCOUNT = "ResourceNotFound.Account"
//  RESOURCENOTFOUND_MIDASEXTERNALAPP = "ResourceNotFound.MidasExternalApp"
//  RESOURCENOTFOUND_MIDASEXTERNALORDER = "ResourceNotFound.MidasExternalOrder"
//  RESOURCENOTFOUND_MIDASORDER = "ResourceNotFound.MidasOrder"
//  RESOURCENOTFOUND_MIDASSIGN = "ResourceNotFound.MidasSign"
//  RESOURCEUNAVAILABLE_MIDASBALANCE = "ResourceUnavailable.MidasBalance"
//  RESOURCEUNAVAILABLE_MIDASDAY = "ResourceUnavailable.MidasDay"
//  RESOURCEUNAVAILABLE_MIDASFROZENAMOUNT = "ResourceUnavailable.MidasFrozenAmount"
//  RESOURCEUNAVAILABLE_MIDASMERCHANTBALANCE = "ResourceUnavailable.MidasMerchantBalance"
//  RESOURCEUNAVAILABLE_MIDASORDER = "ResourceUnavailable.MidasOrder"
//  RESOURCEUNAVAILABLE_MIDASUSERBALANCE = "ResourceUnavailable.MidasUserBalance"
//  RESOURCEUNAVAILABLE_MIDASWALLET = "ResourceUnavailable.MidasWallet"
//  UNAUTHORIZEDOPERATION_MIDAS = "UnauthorizedOperation.Midas"
func (c *Client) CloseCloudOrder(request *CloseCloudOrderRequest) (response *CloseCloudOrderResponse, err error) {
    return c.CloseCloudOrderWithContext(context.Background(), request)
}

// CloseCloudOrder
// 通过此接口关闭此前已创建的订单。关闭后，用户将无法继续付款，仅能关闭创建后未支付的订单。
//
// 可能返回的错误码:
//  AUTHFAILURE_MIDAS = "AuthFailure.Midas"
//  FAILEDOPERATION_ACTION = "FailedOperation.Action"
//  FAILEDOPERATION_MIDASINTERNALERROR = "FailedOperation.MidasInternalError"
//  FAILEDOPERATION_MIDASINVALIDREQUEST = "FailedOperation.MidasInvalidRequest"
//  FAILEDOPERATION_MIDASNEEDRETRY = "FailedOperation.MidasNeedRetry"
//  FAILEDOPERATION_MIDASREGISTERUNFINISHED = "FailedOperation.MidasRegisterUnfinished"
//  FAILEDOPERATION_MIDASREPEATORDER = "FailedOperation.MidasRepeatOrder"
//  FAILEDOPERATION_MIDASRISK = "FailedOperation.MidasRisk"
//  FAILEDOPERATION_MIDASSTATUSNOTMATCH = "FailedOperation.MidasStatusNotMatch"
//  FAILEDOPERATION_MIDASUNSUPPORTEDACTION = "FailedOperation.MidasUnsupportedAction"
//  INTERNALERROR_MIDAS = "InternalError.Midas"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  INVALIDPARAMETER_MIDAS = "InvalidParameter.Midas"
//  INVALIDPARAMETER_MIDASENVIRONMENT = "InvalidParameter.MidasEnvironment"
//  INVALIDPARAMETER_MIDASEXTERNALAPP = "InvalidParameter.MidasExternalApp"
//  INVALIDPARAMETER_MIDASFILETYPE = "InvalidParameter.MidasFileType"
//  INVALIDPARAMETER_MIDASHASH = "InvalidParameter.MidasHash"
//  INVALIDPARAMETER_MIDASSIGNID = "InvalidParameter.MidasSignId"
//  LIMITEXCEEDED_MIDASLARGEFILE = "LimitExceeded.MidasLargeFile"
//  LIMITEXCEEDED_MIDASORDER = "LimitExceeded.MidasOrder"
//  LIMITEXCEEDED_MIDASORDERCANCELED = "LimitExceeded.MidasOrderCanceled"
//  LIMITEXCEEDED_MIDASORDERCLOSED = "LimitExceeded.MidasOrderClosed"
//  LIMITEXCEEDED_MIDASORDEREXPIRED = "LimitExceeded.MidasOrderExpired"
//  LIMITEXCEEDED_MIDASORDERFAILED = "LimitExceeded.MidasOrderFailed"
//  LIMITEXCEEDED_MIDASORDERPARTIALSUCCESS = "LimitExceeded.MidasOrderPartialSuccess"
//  LIMITEXCEEDED_MIDASORDERPROCESSING = "LimitExceeded.MidasOrderProcessing"
//  LIMITEXCEEDED_MIDASORDERSUCCESS = "LimitExceeded.MidasOrderSuccess"
//  REQUESTLIMITEXCEEDED_MIDAS = "RequestLimitExceeded.Midas"
//  REQUESTLIMITEXCEEDED_MIDASINVALIDREQUEST = "RequestLimitExceeded.MidasInvalidRequest"
//  RESOURCEINUSE_MIDAS = "ResourceInUse.Midas"
//  RESOURCENOTFOUND_ACCOUNT = "ResourceNotFound.Account"
//  RESOURCENOTFOUND_MIDASEXTERNALAPP = "ResourceNotFound.MidasExternalApp"
//  RESOURCENOTFOUND_MIDASEXTERNALORDER = "ResourceNotFound.MidasExternalOrder"
//  RESOURCENOTFOUND_MIDASORDER = "ResourceNotFound.MidasOrder"
//  RESOURCENOTFOUND_MIDASSIGN = "ResourceNotFound.MidasSign"
//  RESOURCEUNAVAILABLE_MIDASBALANCE = "ResourceUnavailable.MidasBalance"
//  RESOURCEUNAVAILABLE_MIDASDAY = "ResourceUnavailable.MidasDay"
//  RESOURCEUNAVAILABLE_MIDASFROZENAMOUNT = "ResourceUnavailable.MidasFrozenAmount"
//  RESOURCEUNAVAILABLE_MIDASMERCHANTBALANCE = "ResourceUnavailable.MidasMerchantBalance"
//  RESOURCEUNAVAILABLE_MIDASORDER = "ResourceUnavailable.MidasOrder"
//  RESOURCEUNAVAILABLE_MIDASUSERBALANCE = "ResourceUnavailable.MidasUserBalance"
//  RESOURCEUNAVAILABLE_MIDASWALLET = "ResourceUnavailable.MidasWallet"
//  UNAUTHORIZEDOPERATION_MIDAS = "UnauthorizedOperation.Midas"
func (c *Client) CloseCloudOrderWithContext(ctx context.Context, request *CloseCloudOrderRequest) (response *CloseCloudOrderResponse, err error) {
    if request == nil {
        request = NewCloseCloudOrderRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CloseCloudOrder require credential")
    }

    request.SetContext(ctx)
    
    response = NewCloseCloudOrderResponse()
    err = c.Send(request, response)
    return
}

func NewCloseOpenBankPaymentOrderRequest() (request *CloseOpenBankPaymentOrderRequest) {
    request = &CloseOpenBankPaymentOrderRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cpdp", APIVersion, "CloseOpenBankPaymentOrder")
    
    
    return
}

func NewCloseOpenBankPaymentOrderResponse() (response *CloseOpenBankPaymentOrderResponse) {
    response = &CloseOpenBankPaymentOrderResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CloseOpenBankPaymentOrder
// 云企付-关闭订单
//
// 可能返回的错误码:
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
func (c *Client) CloseOpenBankPaymentOrder(request *CloseOpenBankPaymentOrderRequest) (response *CloseOpenBankPaymentOrderResponse, err error) {
    return c.CloseOpenBankPaymentOrderWithContext(context.Background(), request)
}

// CloseOpenBankPaymentOrder
// 云企付-关闭订单
//
// 可能返回的错误码:
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
func (c *Client) CloseOpenBankPaymentOrderWithContext(ctx context.Context, request *CloseOpenBankPaymentOrderRequest) (response *CloseOpenBankPaymentOrderResponse, err error) {
    if request == nil {
        request = NewCloseOpenBankPaymentOrderRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CloseOpenBankPaymentOrder require credential")
    }

    request.SetContext(ctx)
    
    response = NewCloseOpenBankPaymentOrderResponse()
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
//  AUTHFAILURE_SECRETKEYNOTFOUND = "AuthFailure.SecretKeyNotFound"
//  AUTHFAILURE_VERIFYERROR = "AuthFailure.VerifyError"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_APPDENY = "FailedOperation.AppDeny"
//  FAILEDOPERATION_NORECORD = "FailedOperation.NoRecord"
//  FAILEDOPERATION_ORDERNOTACTIVATED = "FailedOperation.OrderNotActivated"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CloseOrder(request *CloseOrderRequest) (response *CloseOrderResponse, err error) {
    return c.CloseOrderWithContext(context.Background(), request)
}

// CloseOrder
// 通过此接口关闭此前已创建的订单，关闭后，用户将无法继续付款。仅能关闭创建后未支付的订单
//
// 可能返回的错误码:
//  AUTHFAILURE_SECRETKEYNOTFOUND = "AuthFailure.SecretKeyNotFound"
//  AUTHFAILURE_VERIFYERROR = "AuthFailure.VerifyError"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_APPDENY = "FailedOperation.AppDeny"
//  FAILEDOPERATION_NORECORD = "FailedOperation.NoRecord"
//  FAILEDOPERATION_ORDERNOTACTIVATED = "FailedOperation.OrderNotActivated"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CloseOrderWithContext(ctx context.Context, request *CloseOrderRequest) (response *CloseOrderResponse, err error) {
    if request == nil {
        request = NewCloseOrderRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CloseOrder require credential")
    }

    request.SetContext(ctx)
    
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
    return c.ConfirmOrderWithContext(context.Background(), request)
}

// ConfirmOrder
// 云鉴-消费订单确认接口
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDPARAMETER = "FailedOperation.InvalidParameter"
//  FAILEDOPERATION_MISSINGPARAMETER = "FailedOperation.MissingParameter"
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
func (c *Client) ConfirmOrderWithContext(ctx context.Context, request *ConfirmOrderRequest) (response *ConfirmOrderResponse, err error) {
    if request == nil {
        request = NewConfirmOrderRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ConfirmOrder require credential")
    }

    request.SetContext(ctx)
    
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
//  AUTHFAILURE_VERIFYERROR = "AuthFailure.VerifyError"
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
    return c.ContractOrderWithContext(context.Background(), request)
}

// ContractOrder
// 应用需要先带上签约信息调用本接口生成支付订单号，并将应答的PayInfo透传给聚鑫SDK，拉起客户端（包括微信公众号/微信小程序/客户端App）支付。
//
// 可能返回的错误码:
//  AUTHFAILURE_VERIFYERROR = "AuthFailure.VerifyError"
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
func (c *Client) ContractOrderWithContext(ctx context.Context, request *ContractOrderRequest) (response *ContractOrderResponse, err error) {
    if request == nil {
        request = NewContractOrderRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ContractOrder require credential")
    }

    request.SetContext(ctx)
    
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
//  AUTHFAILURE_SECRETKEYNOTFOUND = "AuthFailure.SecretKeyNotFound"
//  FAILEDOPERATION_BANKFAILED = "FailedOperation.BankFailed"
//  FAILEDOPERATION_SERVICEERROR = "FailedOperation.ServiceError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER_ = "MissingParameter."
func (c *Client) CreateAcct(request *CreateAcctRequest) (response *CreateAcctResponse, err error) {
    return c.CreateAcctWithContext(context.Background(), request)
}

// CreateAcct
// 子商户入驻聚鑫平台
//
// 可能返回的错误码:
//  AUTHFAILURE_SECRETKEYNOTFOUND = "AuthFailure.SecretKeyNotFound"
//  FAILEDOPERATION_BANKFAILED = "FailedOperation.BankFailed"
//  FAILEDOPERATION_SERVICEERROR = "FailedOperation.ServiceError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER_ = "MissingParameter."
func (c *Client) CreateAcctWithContext(ctx context.Context, request *CreateAcctRequest) (response *CreateAcctResponse, err error) {
    if request == nil {
        request = NewCreateAcctRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAcct require credential")
    }

    request.SetContext(ctx)
    
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
    return c.CreateAgentTaxPaymentInfosWithContext(context.Background(), request)
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
func (c *Client) CreateAgentTaxPaymentInfosWithContext(ctx context.Context, request *CreateAgentTaxPaymentInfosRequest) (response *CreateAgentTaxPaymentInfosResponse, err error) {
    if request == nil {
        request = NewCreateAgentTaxPaymentInfosRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAgentTaxPaymentInfos require credential")
    }

    request.SetContext(ctx)
    
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
//  FAILEDOPERATION_ACTION = "FailedOperation.Action"
//  FAILEDOPERATION_CREATEAGENT = "FailedOperation.CreateAgent"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  MISSINGPARAMETER_ACTION = "MissingParameter.Action"
//  MISSINGPARAMETER_APPID = "MissingParameter.AppId"
func (c *Client) CreateAnchor(request *CreateAnchorRequest) (response *CreateAnchorResponse, err error) {
    return c.CreateAnchorWithContext(context.Background(), request)
}

// CreateAnchor
// 直播平台-主播入驻
//
// 可能返回的错误码:
//  FAILEDOPERATION_ACTION = "FailedOperation.Action"
//  FAILEDOPERATION_CREATEAGENT = "FailedOperation.CreateAgent"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  MISSINGPARAMETER_ACTION = "MissingParameter.Action"
//  MISSINGPARAMETER_APPID = "MissingParameter.AppId"
func (c *Client) CreateAnchorWithContext(ctx context.Context, request *CreateAnchorRequest) (response *CreateAnchorResponse, err error) {
    if request == nil {
        request = NewCreateAnchorRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAnchor require credential")
    }

    request.SetContext(ctx)
    
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
//  FAILEDOPERATION_ACTION = "FailedOperation.Action"
//  FAILEDOPERATION_CREATEAGENT = "FailedOperation.CreateAgent"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  MISSINGPARAMETER_ACTION = "MissingParameter.Action"
//  MISSINGPARAMETER_APPID = "MissingParameter.AppId"
func (c *Client) CreateBatchPayment(request *CreateBatchPaymentRequest) (response *CreateBatchPaymentResponse, err error) {
    return c.CreateBatchPaymentWithContext(context.Background(), request)
}

// CreateBatchPayment
// 灵云-批量主播转账接口
//
// 可能返回的错误码:
//  FAILEDOPERATION_ACTION = "FailedOperation.Action"
//  FAILEDOPERATION_CREATEAGENT = "FailedOperation.CreateAgent"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  MISSINGPARAMETER_ACTION = "MissingParameter.Action"
//  MISSINGPARAMETER_APPID = "MissingParameter.AppId"
func (c *Client) CreateBatchPaymentWithContext(ctx context.Context, request *CreateBatchPaymentRequest) (response *CreateBatchPaymentResponse, err error) {
    if request == nil {
        request = NewCreateBatchPaymentRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateBatchPayment require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateBatchPaymentResponse()
    err = c.Send(request, response)
    return
}

func NewCreateCloudSubMerchantRequest() (request *CreateCloudSubMerchantRequest) {
    request = &CreateCloudSubMerchantRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cpdp", APIVersion, "CreateCloudSubMerchant")
    
    
    return
}

func NewCreateCloudSubMerchantResponse() (response *CreateCloudSubMerchantResponse) {
    response = &CreateCloudSubMerchantResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateCloudSubMerchant
// 创建子商户
//
// 可能返回的错误码:
//  AUTHFAILURE_MIDAS = "AuthFailure.Midas"
//  FAILEDOPERATION_ACTION = "FailedOperation.Action"
//  FAILEDOPERATION_MIDASINTERNALERROR = "FailedOperation.MidasInternalError"
//  FAILEDOPERATION_MIDASINVALIDREQUEST = "FailedOperation.MidasInvalidRequest"
//  FAILEDOPERATION_MIDASNEEDRETRY = "FailedOperation.MidasNeedRetry"
//  FAILEDOPERATION_MIDASREGISTERUNFINISHED = "FailedOperation.MidasRegisterUnfinished"
//  FAILEDOPERATION_MIDASREPEATORDER = "FailedOperation.MidasRepeatOrder"
//  FAILEDOPERATION_MIDASRISK = "FailedOperation.MidasRisk"
//  FAILEDOPERATION_MIDASSTATUSNOTMATCH = "FailedOperation.MidasStatusNotMatch"
//  FAILEDOPERATION_MIDASUNSUPPORTEDACTION = "FailedOperation.MidasUnsupportedAction"
//  INTERNALERROR_MIDAS = "InternalError.Midas"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  INVALIDPARAMETER_MIDAS = "InvalidParameter.Midas"
//  INVALIDPARAMETER_MIDASENVIRONMENT = "InvalidParameter.MidasEnvironment"
//  INVALIDPARAMETER_MIDASEXTERNALAPP = "InvalidParameter.MidasExternalApp"
//  INVALIDPARAMETER_MIDASFILETYPE = "InvalidParameter.MidasFileType"
//  INVALIDPARAMETER_MIDASHASH = "InvalidParameter.MidasHash"
//  INVALIDPARAMETER_MIDASSIGNID = "InvalidParameter.MidasSignId"
//  LIMITEXCEEDED_MIDASLARGEFILE = "LimitExceeded.MidasLargeFile"
//  LIMITEXCEEDED_MIDASORDER = "LimitExceeded.MidasOrder"
//  LIMITEXCEEDED_MIDASORDERCANCELED = "LimitExceeded.MidasOrderCanceled"
//  LIMITEXCEEDED_MIDASORDERCLOSED = "LimitExceeded.MidasOrderClosed"
//  LIMITEXCEEDED_MIDASORDEREXPIRED = "LimitExceeded.MidasOrderExpired"
//  LIMITEXCEEDED_MIDASORDERFAILED = "LimitExceeded.MidasOrderFailed"
//  LIMITEXCEEDED_MIDASORDERPARTIALSUCCESS = "LimitExceeded.MidasOrderPartialSuccess"
//  LIMITEXCEEDED_MIDASORDERPROCESSING = "LimitExceeded.MidasOrderProcessing"
//  LIMITEXCEEDED_MIDASORDERSUCCESS = "LimitExceeded.MidasOrderSuccess"
//  REQUESTLIMITEXCEEDED_MIDAS = "RequestLimitExceeded.Midas"
//  REQUESTLIMITEXCEEDED_MIDASINVALIDREQUEST = "RequestLimitExceeded.MidasInvalidRequest"
//  RESOURCEINUSE_MIDAS = "ResourceInUse.Midas"
//  RESOURCENOTFOUND_ACCOUNT = "ResourceNotFound.Account"
//  RESOURCENOTFOUND_MIDASEXTERNALAPP = "ResourceNotFound.MidasExternalApp"
//  RESOURCENOTFOUND_MIDASEXTERNALORDER = "ResourceNotFound.MidasExternalOrder"
//  RESOURCENOTFOUND_MIDASORDER = "ResourceNotFound.MidasOrder"
//  RESOURCENOTFOUND_MIDASSIGN = "ResourceNotFound.MidasSign"
//  RESOURCEUNAVAILABLE_MIDASBALANCE = "ResourceUnavailable.MidasBalance"
//  RESOURCEUNAVAILABLE_MIDASDAY = "ResourceUnavailable.MidasDay"
//  RESOURCEUNAVAILABLE_MIDASFROZENAMOUNT = "ResourceUnavailable.MidasFrozenAmount"
//  RESOURCEUNAVAILABLE_MIDASMERCHANTBALANCE = "ResourceUnavailable.MidasMerchantBalance"
//  RESOURCEUNAVAILABLE_MIDASORDER = "ResourceUnavailable.MidasOrder"
//  RESOURCEUNAVAILABLE_MIDASUSERBALANCE = "ResourceUnavailable.MidasUserBalance"
//  RESOURCEUNAVAILABLE_MIDASWALLET = "ResourceUnavailable.MidasWallet"
//  UNAUTHORIZEDOPERATION_MIDAS = "UnauthorizedOperation.Midas"
func (c *Client) CreateCloudSubMerchant(request *CreateCloudSubMerchantRequest) (response *CreateCloudSubMerchantResponse, err error) {
    return c.CreateCloudSubMerchantWithContext(context.Background(), request)
}

// CreateCloudSubMerchant
// 创建子商户
//
// 可能返回的错误码:
//  AUTHFAILURE_MIDAS = "AuthFailure.Midas"
//  FAILEDOPERATION_ACTION = "FailedOperation.Action"
//  FAILEDOPERATION_MIDASINTERNALERROR = "FailedOperation.MidasInternalError"
//  FAILEDOPERATION_MIDASINVALIDREQUEST = "FailedOperation.MidasInvalidRequest"
//  FAILEDOPERATION_MIDASNEEDRETRY = "FailedOperation.MidasNeedRetry"
//  FAILEDOPERATION_MIDASREGISTERUNFINISHED = "FailedOperation.MidasRegisterUnfinished"
//  FAILEDOPERATION_MIDASREPEATORDER = "FailedOperation.MidasRepeatOrder"
//  FAILEDOPERATION_MIDASRISK = "FailedOperation.MidasRisk"
//  FAILEDOPERATION_MIDASSTATUSNOTMATCH = "FailedOperation.MidasStatusNotMatch"
//  FAILEDOPERATION_MIDASUNSUPPORTEDACTION = "FailedOperation.MidasUnsupportedAction"
//  INTERNALERROR_MIDAS = "InternalError.Midas"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  INVALIDPARAMETER_MIDAS = "InvalidParameter.Midas"
//  INVALIDPARAMETER_MIDASENVIRONMENT = "InvalidParameter.MidasEnvironment"
//  INVALIDPARAMETER_MIDASEXTERNALAPP = "InvalidParameter.MidasExternalApp"
//  INVALIDPARAMETER_MIDASFILETYPE = "InvalidParameter.MidasFileType"
//  INVALIDPARAMETER_MIDASHASH = "InvalidParameter.MidasHash"
//  INVALIDPARAMETER_MIDASSIGNID = "InvalidParameter.MidasSignId"
//  LIMITEXCEEDED_MIDASLARGEFILE = "LimitExceeded.MidasLargeFile"
//  LIMITEXCEEDED_MIDASORDER = "LimitExceeded.MidasOrder"
//  LIMITEXCEEDED_MIDASORDERCANCELED = "LimitExceeded.MidasOrderCanceled"
//  LIMITEXCEEDED_MIDASORDERCLOSED = "LimitExceeded.MidasOrderClosed"
//  LIMITEXCEEDED_MIDASORDEREXPIRED = "LimitExceeded.MidasOrderExpired"
//  LIMITEXCEEDED_MIDASORDERFAILED = "LimitExceeded.MidasOrderFailed"
//  LIMITEXCEEDED_MIDASORDERPARTIALSUCCESS = "LimitExceeded.MidasOrderPartialSuccess"
//  LIMITEXCEEDED_MIDASORDERPROCESSING = "LimitExceeded.MidasOrderProcessing"
//  LIMITEXCEEDED_MIDASORDERSUCCESS = "LimitExceeded.MidasOrderSuccess"
//  REQUESTLIMITEXCEEDED_MIDAS = "RequestLimitExceeded.Midas"
//  REQUESTLIMITEXCEEDED_MIDASINVALIDREQUEST = "RequestLimitExceeded.MidasInvalidRequest"
//  RESOURCEINUSE_MIDAS = "ResourceInUse.Midas"
//  RESOURCENOTFOUND_ACCOUNT = "ResourceNotFound.Account"
//  RESOURCENOTFOUND_MIDASEXTERNALAPP = "ResourceNotFound.MidasExternalApp"
//  RESOURCENOTFOUND_MIDASEXTERNALORDER = "ResourceNotFound.MidasExternalOrder"
//  RESOURCENOTFOUND_MIDASORDER = "ResourceNotFound.MidasOrder"
//  RESOURCENOTFOUND_MIDASSIGN = "ResourceNotFound.MidasSign"
//  RESOURCEUNAVAILABLE_MIDASBALANCE = "ResourceUnavailable.MidasBalance"
//  RESOURCEUNAVAILABLE_MIDASDAY = "ResourceUnavailable.MidasDay"
//  RESOURCEUNAVAILABLE_MIDASFROZENAMOUNT = "ResourceUnavailable.MidasFrozenAmount"
//  RESOURCEUNAVAILABLE_MIDASMERCHANTBALANCE = "ResourceUnavailable.MidasMerchantBalance"
//  RESOURCEUNAVAILABLE_MIDASORDER = "ResourceUnavailable.MidasOrder"
//  RESOURCEUNAVAILABLE_MIDASUSERBALANCE = "ResourceUnavailable.MidasUserBalance"
//  RESOURCEUNAVAILABLE_MIDASWALLET = "ResourceUnavailable.MidasWallet"
//  UNAUTHORIZEDOPERATION_MIDAS = "UnauthorizedOperation.Midas"
func (c *Client) CreateCloudSubMerchantWithContext(ctx context.Context, request *CreateCloudSubMerchantRequest) (response *CreateCloudSubMerchantResponse, err error) {
    if request == nil {
        request = NewCreateCloudSubMerchantRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCloudSubMerchant require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateCloudSubMerchantResponse()
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
    return c.CreateCustAcctIdWithContext(context.Background(), request)
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
func (c *Client) CreateCustAcctIdWithContext(ctx context.Context, request *CreateCustAcctIdRequest) (response *CreateCustAcctIdResponse, err error) {
    if request == nil {
        request = NewCreateCustAcctIdRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCustAcctId require credential")
    }

    request.SetContext(ctx)
    
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
    return c.CreateExternalAnchorWithContext(context.Background(), request)
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
func (c *Client) CreateExternalAnchorWithContext(ctx context.Context, request *CreateExternalAnchorRequest) (response *CreateExternalAnchorResponse, err error) {
    if request == nil {
        request = NewCreateExternalAnchorRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateExternalAnchor require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateExternalAnchorResponse()
    err = c.Send(request, response)
    return
}

func NewCreateFlexPayeeRequest() (request *CreateFlexPayeeRequest) {
    request = &CreateFlexPayeeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cpdp", APIVersion, "CreateFlexPayee")
    
    
    return
}

func NewCreateFlexPayeeResponse() (response *CreateFlexPayeeResponse) {
    response = &CreateFlexPayeeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateFlexPayee
// 灵云V2-收款用户开立
//
// 可能返回的错误码:
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
func (c *Client) CreateFlexPayee(request *CreateFlexPayeeRequest) (response *CreateFlexPayeeResponse, err error) {
    return c.CreateFlexPayeeWithContext(context.Background(), request)
}

// CreateFlexPayee
// 灵云V2-收款用户开立
//
// 可能返回的错误码:
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
func (c *Client) CreateFlexPayeeWithContext(ctx context.Context, request *CreateFlexPayeeRequest) (response *CreateFlexPayeeResponse, err error) {
    if request == nil {
        request = NewCreateFlexPayeeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateFlexPayee require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateFlexPayeeResponse()
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
    return c.CreateInvoiceWithContext(context.Background(), request)
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
func (c *Client) CreateInvoiceWithContext(ctx context.Context, request *CreateInvoiceRequest) (response *CreateInvoiceResponse, err error) {
    if request == nil {
        request = NewCreateInvoiceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateInvoice require credential")
    }

    request.SetContext(ctx)
    
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
    return c.CreateInvoiceV2WithContext(context.Background(), request)
}

// CreateInvoiceV2
// 智慧零售-发票开具V2
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_INVOICENOTFOUND = "ResourceNotFound.InvoiceNotFound"
func (c *Client) CreateInvoiceV2WithContext(ctx context.Context, request *CreateInvoiceV2Request) (response *CreateInvoiceV2Response, err error) {
    if request == nil {
        request = NewCreateInvoiceV2Request()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateInvoiceV2 require credential")
    }

    request.SetContext(ctx)
    
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
    return c.CreateMerchantWithContext(context.Background(), request)
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
func (c *Client) CreateMerchantWithContext(ctx context.Context, request *CreateMerchantRequest) (response *CreateMerchantResponse, err error) {
    if request == nil {
        request = NewCreateMerchantRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateMerchant require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateMerchantResponse()
    err = c.Send(request, response)
    return
}

func NewCreateOpenBankAggregatedSubMerchantRegistrationRequest() (request *CreateOpenBankAggregatedSubMerchantRegistrationRequest) {
    request = &CreateOpenBankAggregatedSubMerchantRegistrationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cpdp", APIVersion, "CreateOpenBankAggregatedSubMerchantRegistration")
    
    
    return
}

func NewCreateOpenBankAggregatedSubMerchantRegistrationResponse() (response *CreateOpenBankAggregatedSubMerchantRegistrationResponse) {
    response = &CreateOpenBankAggregatedSubMerchantRegistrationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateOpenBankAggregatedSubMerchantRegistration
// 云企付-子商户进件V2
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
func (c *Client) CreateOpenBankAggregatedSubMerchantRegistration(request *CreateOpenBankAggregatedSubMerchantRegistrationRequest) (response *CreateOpenBankAggregatedSubMerchantRegistrationResponse, err error) {
    return c.CreateOpenBankAggregatedSubMerchantRegistrationWithContext(context.Background(), request)
}

// CreateOpenBankAggregatedSubMerchantRegistration
// 云企付-子商户进件V2
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
func (c *Client) CreateOpenBankAggregatedSubMerchantRegistrationWithContext(ctx context.Context, request *CreateOpenBankAggregatedSubMerchantRegistrationRequest) (response *CreateOpenBankAggregatedSubMerchantRegistrationResponse, err error) {
    if request == nil {
        request = NewCreateOpenBankAggregatedSubMerchantRegistrationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateOpenBankAggregatedSubMerchantRegistration require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateOpenBankAggregatedSubMerchantRegistrationResponse()
    err = c.Send(request, response)
    return
}

func NewCreateOpenBankExternalSubMerchantAccountBookRequest() (request *CreateOpenBankExternalSubMerchantAccountBookRequest) {
    request = &CreateOpenBankExternalSubMerchantAccountBookRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cpdp", APIVersion, "CreateOpenBankExternalSubMerchantAccountBook")
    
    
    return
}

func NewCreateOpenBankExternalSubMerchantAccountBookResponse() (response *CreateOpenBankExternalSubMerchantAccountBookResponse) {
    response = &CreateOpenBankExternalSubMerchantAccountBookResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateOpenBankExternalSubMerchantAccountBook
// 第三方子商户电子记账本创建接口
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
func (c *Client) CreateOpenBankExternalSubMerchantAccountBook(request *CreateOpenBankExternalSubMerchantAccountBookRequest) (response *CreateOpenBankExternalSubMerchantAccountBookResponse, err error) {
    return c.CreateOpenBankExternalSubMerchantAccountBookWithContext(context.Background(), request)
}

// CreateOpenBankExternalSubMerchantAccountBook
// 第三方子商户电子记账本创建接口
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
func (c *Client) CreateOpenBankExternalSubMerchantAccountBookWithContext(ctx context.Context, request *CreateOpenBankExternalSubMerchantAccountBookRequest) (response *CreateOpenBankExternalSubMerchantAccountBookResponse, err error) {
    if request == nil {
        request = NewCreateOpenBankExternalSubMerchantAccountBookRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateOpenBankExternalSubMerchantAccountBook require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateOpenBankExternalSubMerchantAccountBookResponse()
    err = c.Send(request, response)
    return
}

func NewCreateOpenBankExternalSubMerchantRegistrationRequest() (request *CreateOpenBankExternalSubMerchantRegistrationRequest) {
    request = &CreateOpenBankExternalSubMerchantRegistrationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cpdp", APIVersion, "CreateOpenBankExternalSubMerchantRegistration")
    
    
    return
}

func NewCreateOpenBankExternalSubMerchantRegistrationResponse() (response *CreateOpenBankExternalSubMerchantRegistrationResponse) {
    response = &CreateOpenBankExternalSubMerchantRegistrationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateOpenBankExternalSubMerchantRegistration
// 云企付-子商户进件
//
// 可能返回的错误码:
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
func (c *Client) CreateOpenBankExternalSubMerchantRegistration(request *CreateOpenBankExternalSubMerchantRegistrationRequest) (response *CreateOpenBankExternalSubMerchantRegistrationResponse, err error) {
    return c.CreateOpenBankExternalSubMerchantRegistrationWithContext(context.Background(), request)
}

// CreateOpenBankExternalSubMerchantRegistration
// 云企付-子商户进件
//
// 可能返回的错误码:
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
func (c *Client) CreateOpenBankExternalSubMerchantRegistrationWithContext(ctx context.Context, request *CreateOpenBankExternalSubMerchantRegistrationRequest) (response *CreateOpenBankExternalSubMerchantRegistrationResponse, err error) {
    if request == nil {
        request = NewCreateOpenBankExternalSubMerchantRegistrationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateOpenBankExternalSubMerchantRegistration require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateOpenBankExternalSubMerchantRegistrationResponse()
    err = c.Send(request, response)
    return
}

func NewCreateOpenBankGlobalPaymentOrderRequest() (request *CreateOpenBankGlobalPaymentOrderRequest) {
    request = &CreateOpenBankGlobalPaymentOrderRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cpdp", APIVersion, "CreateOpenBankGlobalPaymentOrder")
    
    
    return
}

func NewCreateOpenBankGlobalPaymentOrderResponse() (response *CreateOpenBankGlobalPaymentOrderResponse) {
    response = &CreateOpenBankGlobalPaymentOrderResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateOpenBankGlobalPaymentOrder
// 云企付-跨境支付下单
//
// 可能返回的错误码:
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
func (c *Client) CreateOpenBankGlobalPaymentOrder(request *CreateOpenBankGlobalPaymentOrderRequest) (response *CreateOpenBankGlobalPaymentOrderResponse, err error) {
    return c.CreateOpenBankGlobalPaymentOrderWithContext(context.Background(), request)
}

// CreateOpenBankGlobalPaymentOrder
// 云企付-跨境支付下单
//
// 可能返回的错误码:
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
func (c *Client) CreateOpenBankGlobalPaymentOrderWithContext(ctx context.Context, request *CreateOpenBankGlobalPaymentOrderRequest) (response *CreateOpenBankGlobalPaymentOrderResponse, err error) {
    if request == nil {
        request = NewCreateOpenBankGlobalPaymentOrderRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateOpenBankGlobalPaymentOrder require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateOpenBankGlobalPaymentOrderResponse()
    err = c.Send(request, response)
    return
}

func NewCreateOpenBankMerchantRequest() (request *CreateOpenBankMerchantRequest) {
    request = &CreateOpenBankMerchantRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cpdp", APIVersion, "CreateOpenBankMerchant")
    
    
    return
}

func NewCreateOpenBankMerchantResponse() (response *CreateOpenBankMerchantResponse) {
    response = &CreateOpenBankMerchantResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateOpenBankMerchant
// 云企付-创建商户
//
// 可能返回的错误码:
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
func (c *Client) CreateOpenBankMerchant(request *CreateOpenBankMerchantRequest) (response *CreateOpenBankMerchantResponse, err error) {
    return c.CreateOpenBankMerchantWithContext(context.Background(), request)
}

// CreateOpenBankMerchant
// 云企付-创建商户
//
// 可能返回的错误码:
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
func (c *Client) CreateOpenBankMerchantWithContext(ctx context.Context, request *CreateOpenBankMerchantRequest) (response *CreateOpenBankMerchantResponse, err error) {
    if request == nil {
        request = NewCreateOpenBankMerchantRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateOpenBankMerchant require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateOpenBankMerchantResponse()
    err = c.Send(request, response)
    return
}

func NewCreateOpenBankPaymentOrderRequest() (request *CreateOpenBankPaymentOrderRequest) {
    request = &CreateOpenBankPaymentOrderRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cpdp", APIVersion, "CreateOpenBankPaymentOrder")
    
    
    return
}

func NewCreateOpenBankPaymentOrderResponse() (response *CreateOpenBankPaymentOrderResponse) {
    response = &CreateOpenBankPaymentOrderResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateOpenBankPaymentOrder
// 云企付-创建支付订单。支持B2B网关支付，B2C转账下单。
//
// 可能返回的错误码:
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
func (c *Client) CreateOpenBankPaymentOrder(request *CreateOpenBankPaymentOrderRequest) (response *CreateOpenBankPaymentOrderResponse, err error) {
    return c.CreateOpenBankPaymentOrderWithContext(context.Background(), request)
}

// CreateOpenBankPaymentOrder
// 云企付-创建支付订单。支持B2B网关支付，B2C转账下单。
//
// 可能返回的错误码:
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
func (c *Client) CreateOpenBankPaymentOrderWithContext(ctx context.Context, request *CreateOpenBankPaymentOrderRequest) (response *CreateOpenBankPaymentOrderResponse, err error) {
    if request == nil {
        request = NewCreateOpenBankPaymentOrderRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateOpenBankPaymentOrder require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateOpenBankPaymentOrderResponse()
    err = c.Send(request, response)
    return
}

func NewCreateOpenBankRechargeOrderRequest() (request *CreateOpenBankRechargeOrderRequest) {
    request = &CreateOpenBankRechargeOrderRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cpdp", APIVersion, "CreateOpenBankRechargeOrder")
    
    
    return
}

func NewCreateOpenBankRechargeOrderResponse() (response *CreateOpenBankRechargeOrderResponse) {
    response = &CreateOpenBankRechargeOrderResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateOpenBankRechargeOrder
// 云企付-创建充值订单
//
// 可能返回的错误码:
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
func (c *Client) CreateOpenBankRechargeOrder(request *CreateOpenBankRechargeOrderRequest) (response *CreateOpenBankRechargeOrderResponse, err error) {
    return c.CreateOpenBankRechargeOrderWithContext(context.Background(), request)
}

// CreateOpenBankRechargeOrder
// 云企付-创建充值订单
//
// 可能返回的错误码:
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
func (c *Client) CreateOpenBankRechargeOrderWithContext(ctx context.Context, request *CreateOpenBankRechargeOrderRequest) (response *CreateOpenBankRechargeOrderResponse, err error) {
    if request == nil {
        request = NewCreateOpenBankRechargeOrderRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateOpenBankRechargeOrder require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateOpenBankRechargeOrderResponse()
    err = c.Send(request, response)
    return
}

func NewCreateOpenBankSubMerchantRateConfigureRequest() (request *CreateOpenBankSubMerchantRateConfigureRequest) {
    request = &CreateOpenBankSubMerchantRateConfigureRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cpdp", APIVersion, "CreateOpenBankSubMerchantRateConfigure")
    
    
    return
}

func NewCreateOpenBankSubMerchantRateConfigureResponse() (response *CreateOpenBankSubMerchantRateConfigureResponse) {
    response = &CreateOpenBankSubMerchantRateConfigureResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateOpenBankSubMerchantRateConfigure
// 云企付-子商户费率配置
//
// 可能返回的错误码:
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
func (c *Client) CreateOpenBankSubMerchantRateConfigure(request *CreateOpenBankSubMerchantRateConfigureRequest) (response *CreateOpenBankSubMerchantRateConfigureResponse, err error) {
    return c.CreateOpenBankSubMerchantRateConfigureWithContext(context.Background(), request)
}

// CreateOpenBankSubMerchantRateConfigure
// 云企付-子商户费率配置
//
// 可能返回的错误码:
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
func (c *Client) CreateOpenBankSubMerchantRateConfigureWithContext(ctx context.Context, request *CreateOpenBankSubMerchantRateConfigureRequest) (response *CreateOpenBankSubMerchantRateConfigureResponse, err error) {
    if request == nil {
        request = NewCreateOpenBankSubMerchantRateConfigureRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateOpenBankSubMerchantRateConfigure require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateOpenBankSubMerchantRateConfigureResponse()
    err = c.Send(request, response)
    return
}

func NewCreateOpenBankUnifiedOrderRequest() (request *CreateOpenBankUnifiedOrderRequest) {
    request = &CreateOpenBankUnifiedOrderRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cpdp", APIVersion, "CreateOpenBankUnifiedOrder")
    
    
    return
}

func NewCreateOpenBankUnifiedOrderResponse() (response *CreateOpenBankUnifiedOrderResponse) {
    response = &CreateOpenBankUnifiedOrderResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateOpenBankUnifiedOrder
// 云企付-聚合下单
//
// 可能返回的错误码:
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
func (c *Client) CreateOpenBankUnifiedOrder(request *CreateOpenBankUnifiedOrderRequest) (response *CreateOpenBankUnifiedOrderResponse, err error) {
    return c.CreateOpenBankUnifiedOrderWithContext(context.Background(), request)
}

// CreateOpenBankUnifiedOrder
// 云企付-聚合下单
//
// 可能返回的错误码:
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
func (c *Client) CreateOpenBankUnifiedOrderWithContext(ctx context.Context, request *CreateOpenBankUnifiedOrderRequest) (response *CreateOpenBankUnifiedOrderResponse, err error) {
    if request == nil {
        request = NewCreateOpenBankUnifiedOrderRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateOpenBankUnifiedOrder require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateOpenBankUnifiedOrderResponse()
    err = c.Send(request, response)
    return
}

func NewCreateOpenBankVerificationOrderRequest() (request *CreateOpenBankVerificationOrderRequest) {
    request = &CreateOpenBankVerificationOrderRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cpdp", APIVersion, "CreateOpenBankVerificationOrder")
    
    
    return
}

func NewCreateOpenBankVerificationOrderResponse() (response *CreateOpenBankVerificationOrderResponse) {
    response = &CreateOpenBankVerificationOrderResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateOpenBankVerificationOrder
// 云企付-创建核销申请，适用于针对支付订单维度的确认收货，解冻等业务场景。目前支持的渠道有TENPAY下的EBANK_PAYMENT付款方式创建支付订单时，选择担保支付下单的订单进行解冻。
//
// 可能返回的错误码:
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
func (c *Client) CreateOpenBankVerificationOrder(request *CreateOpenBankVerificationOrderRequest) (response *CreateOpenBankVerificationOrderResponse, err error) {
    return c.CreateOpenBankVerificationOrderWithContext(context.Background(), request)
}

// CreateOpenBankVerificationOrder
// 云企付-创建核销申请，适用于针对支付订单维度的确认收货，解冻等业务场景。目前支持的渠道有TENPAY下的EBANK_PAYMENT付款方式创建支付订单时，选择担保支付下单的订单进行解冻。
//
// 可能返回的错误码:
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
func (c *Client) CreateOpenBankVerificationOrderWithContext(ctx context.Context, request *CreateOpenBankVerificationOrderRequest) (response *CreateOpenBankVerificationOrderResponse, err error) {
    if request == nil {
        request = NewCreateOpenBankVerificationOrderRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateOpenBankVerificationOrder require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateOpenBankVerificationOrderResponse()
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
    return c.CreateOrderWithContext(context.Background(), request)
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
func (c *Client) CreateOrderWithContext(ctx context.Context, request *CreateOrderRequest) (response *CreateOrderResponse, err error) {
    if request == nil {
        request = NewCreateOrderRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateOrder require credential")
    }

    request.SetContext(ctx)
    
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
    return c.CreatePayMerchantWithContext(context.Background(), request)
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
func (c *Client) CreatePayMerchantWithContext(ctx context.Context, request *CreatePayMerchantRequest) (response *CreatePayMerchantResponse, err error) {
    if request == nil {
        request = NewCreatePayMerchantRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreatePayMerchant require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreatePayMerchantResponse()
    err = c.Send(request, response)
    return
}

func NewCreatePayRollPreOrderRequest() (request *CreatePayRollPreOrderRequest) {
    request = &CreatePayRollPreOrderRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cpdp", APIVersion, "CreatePayRollPreOrder")
    
    
    return
}

func NewCreatePayRollPreOrderResponse() (response *CreatePayRollPreOrderResponse) {
    response = &CreatePayRollPreOrderResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreatePayRollPreOrder
// 务工卡-核身预下单
//
// 可能返回的错误码:
//  AUTHFAILURE_SIGN = "AuthFailure.Sign"
//  FAILEDOPERATION_ACTION = "FailedOperation.Action"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  MISSINGPARAMETER_ACTION = "MissingParameter.Action"
//  MISSINGPARAMETER_APPID = "MissingParameter.AppId"
//  RESOURCENOTFOUND_ACCOUNT = "ResourceNotFound.Account"
//  RESOURCENOTFOUND_KEY = "ResourceNotFound.Key"
func (c *Client) CreatePayRollPreOrder(request *CreatePayRollPreOrderRequest) (response *CreatePayRollPreOrderResponse, err error) {
    return c.CreatePayRollPreOrderWithContext(context.Background(), request)
}

// CreatePayRollPreOrder
// 务工卡-核身预下单
//
// 可能返回的错误码:
//  AUTHFAILURE_SIGN = "AuthFailure.Sign"
//  FAILEDOPERATION_ACTION = "FailedOperation.Action"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  MISSINGPARAMETER_ACTION = "MissingParameter.Action"
//  MISSINGPARAMETER_APPID = "MissingParameter.AppId"
//  RESOURCENOTFOUND_ACCOUNT = "ResourceNotFound.Account"
//  RESOURCENOTFOUND_KEY = "ResourceNotFound.Key"
func (c *Client) CreatePayRollPreOrderWithContext(ctx context.Context, request *CreatePayRollPreOrderRequest) (response *CreatePayRollPreOrderResponse, err error) {
    if request == nil {
        request = NewCreatePayRollPreOrderRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreatePayRollPreOrder require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreatePayRollPreOrderResponse()
    err = c.Send(request, response)
    return
}

func NewCreatePayRollPreOrderWithAuthRequest() (request *CreatePayRollPreOrderWithAuthRequest) {
    request = &CreatePayRollPreOrderWithAuthRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cpdp", APIVersion, "CreatePayRollPreOrderWithAuth")
    
    
    return
}

func NewCreatePayRollPreOrderWithAuthResponse() (response *CreatePayRollPreOrderWithAuthResponse) {
    response = &CreatePayRollPreOrderWithAuthResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreatePayRollPreOrderWithAuth
// 务工卡-核身预下单带授权
//
// 可能返回的错误码:
//  AUTHFAILURE_SIGN = "AuthFailure.Sign"
//  FAILEDOPERATION_ACTION = "FailedOperation.Action"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  MISSINGPARAMETER_ACTION = "MissingParameter.Action"
//  MISSINGPARAMETER_APPID = "MissingParameter.AppId"
//  RESOURCENOTFOUND_ACCOUNT = "ResourceNotFound.Account"
//  RESOURCENOTFOUND_KEY = "ResourceNotFound.Key"
func (c *Client) CreatePayRollPreOrderWithAuth(request *CreatePayRollPreOrderWithAuthRequest) (response *CreatePayRollPreOrderWithAuthResponse, err error) {
    return c.CreatePayRollPreOrderWithAuthWithContext(context.Background(), request)
}

// CreatePayRollPreOrderWithAuth
// 务工卡-核身预下单带授权
//
// 可能返回的错误码:
//  AUTHFAILURE_SIGN = "AuthFailure.Sign"
//  FAILEDOPERATION_ACTION = "FailedOperation.Action"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  MISSINGPARAMETER_ACTION = "MissingParameter.Action"
//  MISSINGPARAMETER_APPID = "MissingParameter.AppId"
//  RESOURCENOTFOUND_ACCOUNT = "ResourceNotFound.Account"
//  RESOURCENOTFOUND_KEY = "ResourceNotFound.Key"
func (c *Client) CreatePayRollPreOrderWithAuthWithContext(ctx context.Context, request *CreatePayRollPreOrderWithAuthRequest) (response *CreatePayRollPreOrderWithAuthResponse, err error) {
    if request == nil {
        request = NewCreatePayRollPreOrderWithAuthRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreatePayRollPreOrderWithAuth require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreatePayRollPreOrderWithAuthResponse()
    err = c.Send(request, response)
    return
}

func NewCreatePayRollTokenRequest() (request *CreatePayRollTokenRequest) {
    request = &CreatePayRollTokenRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cpdp", APIVersion, "CreatePayRollToken")
    
    
    return
}

func NewCreatePayRollTokenResponse() (response *CreatePayRollTokenResponse) {
    response = &CreatePayRollTokenResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreatePayRollToken
// 务工卡-生成授权令牌
//
// 可能返回的错误码:
//  AUTHFAILURE_SIGN = "AuthFailure.Sign"
//  FAILEDOPERATION_ACTION = "FailedOperation.Action"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  MISSINGPARAMETER_ACTION = "MissingParameter.Action"
//  MISSINGPARAMETER_APPID = "MissingParameter.AppId"
//  RESOURCENOTFOUND_ACCOUNT = "ResourceNotFound.Account"
//  RESOURCENOTFOUND_KEY = "ResourceNotFound.Key"
func (c *Client) CreatePayRollToken(request *CreatePayRollTokenRequest) (response *CreatePayRollTokenResponse, err error) {
    return c.CreatePayRollTokenWithContext(context.Background(), request)
}

// CreatePayRollToken
// 务工卡-生成授权令牌
//
// 可能返回的错误码:
//  AUTHFAILURE_SIGN = "AuthFailure.Sign"
//  FAILEDOPERATION_ACTION = "FailedOperation.Action"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  MISSINGPARAMETER_ACTION = "MissingParameter.Action"
//  MISSINGPARAMETER_APPID = "MissingParameter.AppId"
//  RESOURCENOTFOUND_ACCOUNT = "ResourceNotFound.Account"
//  RESOURCENOTFOUND_KEY = "ResourceNotFound.Key"
func (c *Client) CreatePayRollTokenWithContext(ctx context.Context, request *CreatePayRollTokenRequest) (response *CreatePayRollTokenResponse, err error) {
    if request == nil {
        request = NewCreatePayRollTokenRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreatePayRollToken require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreatePayRollTokenResponse()
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
    return c.CreateRedInvoiceWithContext(context.Background(), request)
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
func (c *Client) CreateRedInvoiceWithContext(ctx context.Context, request *CreateRedInvoiceRequest) (response *CreateRedInvoiceResponse, err error) {
    if request == nil {
        request = NewCreateRedInvoiceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateRedInvoice require credential")
    }

    request.SetContext(ctx)
    
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
//  RESOURCENOTFOUND_INVOICENOTFOUND = "ResourceNotFound.InvoiceNotFound"
func (c *Client) CreateRedInvoiceV2(request *CreateRedInvoiceV2Request) (response *CreateRedInvoiceV2Response, err error) {
    return c.CreateRedInvoiceV2WithContext(context.Background(), request)
}

// CreateRedInvoiceV2
// 智慧零售-发票红冲V2
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_INVOICENOTFOUND = "ResourceNotFound.InvoiceNotFound"
func (c *Client) CreateRedInvoiceV2WithContext(ctx context.Context, request *CreateRedInvoiceV2Request) (response *CreateRedInvoiceV2Response, err error) {
    if request == nil {
        request = NewCreateRedInvoiceV2Request()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateRedInvoiceV2 require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateRedInvoiceV2Response()
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
//  RESOURCENOTFOUND_INVOICENOTFOUND = "ResourceNotFound.InvoiceNotFound"
func (c *Client) CreateSinglePayment(request *CreateSinglePaymentRequest) (response *CreateSinglePaymentResponse, err error) {
    return c.CreateSinglePaymentWithContext(context.Background(), request)
}

// CreateSinglePayment
// 灵云-单笔主播转账接口
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_INVOICENOTFOUND = "ResourceNotFound.InvoiceNotFound"
func (c *Client) CreateSinglePaymentWithContext(ctx context.Context, request *CreateSinglePaymentRequest) (response *CreateSinglePaymentResponse, err error) {
    if request == nil {
        request = NewCreateSinglePaymentRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateSinglePayment require credential")
    }

    request.SetContext(ctx)
    
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
    return c.CreateTransferBatchWithContext(context.Background(), request)
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
func (c *Client) CreateTransferBatchWithContext(ctx context.Context, request *CreateTransferBatchRequest) (response *CreateTransferBatchResponse, err error) {
    if request == nil {
        request = NewCreateTransferBatchRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateTransferBatch require credential")
    }

    request.SetContext(ctx)
    
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
    return c.DeduceQuotaWithContext(context.Background(), request)
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
func (c *Client) DeduceQuotaWithContext(ctx context.Context, request *DeduceQuotaRequest) (response *DeduceQuotaResponse, err error) {
    if request == nil {
        request = NewDeduceQuotaRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeduceQuota require credential")
    }

    request.SetContext(ctx)
    
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
    return c.DeleteAgentTaxPaymentInfoWithContext(context.Background(), request)
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
func (c *Client) DeleteAgentTaxPaymentInfoWithContext(ctx context.Context, request *DeleteAgentTaxPaymentInfoRequest) (response *DeleteAgentTaxPaymentInfoResponse, err error) {
    if request == nil {
        request = NewDeleteAgentTaxPaymentInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteAgentTaxPaymentInfo require credential")
    }

    request.SetContext(ctx)
    
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
    return c.DeleteAgentTaxPaymentInfosWithContext(context.Background(), request)
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
func (c *Client) DeleteAgentTaxPaymentInfosWithContext(ctx context.Context, request *DeleteAgentTaxPaymentInfosRequest) (response *DeleteAgentTaxPaymentInfosResponse, err error) {
    if request == nil {
        request = NewDeleteAgentTaxPaymentInfosRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteAgentTaxPaymentInfos require credential")
    }

    request.SetContext(ctx)
    
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
    return c.DescribeChargeDetailWithContext(context.Background(), request)
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
func (c *Client) DescribeChargeDetailWithContext(ctx context.Context, request *DescribeChargeDetailRequest) (response *DescribeChargeDetailResponse, err error) {
    if request == nil {
        request = NewDescribeChargeDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeChargeDetail require credential")
    }

    request.SetContext(ctx)
    
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
    return c.DescribeOrderStatusWithContext(context.Background(), request)
}

// DescribeOrderStatus
// 查询单笔订单交易状态
//
// 可能返回的错误码:
//  FAILEDOPERATION_ACTIONINVALID = "FailedOperation.ActionInvalid"
//  FAILEDOPERATION_PABANKERROR = "FailedOperation.PABankError"
func (c *Client) DescribeOrderStatusWithContext(ctx context.Context, request *DescribeOrderStatusRequest) (response *DescribeOrderStatusResponse, err error) {
    if request == nil {
        request = NewDescribeOrderStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeOrderStatus require credential")
    }

    request.SetContext(ctx)
    
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
    return c.DistributeAccreditQueryWithContext(context.Background(), request)
}

// DistributeAccreditQuery
// 云支付-分账授权申请查询接口
//
// 可能返回的错误码:
//  FAILEDOPERATION_ACTIONINVALID = "FailedOperation.ActionInvalid"
//  FAILEDOPERATION_PABANKERROR = "FailedOperation.PABankError"
func (c *Client) DistributeAccreditQueryWithContext(ctx context.Context, request *DistributeAccreditQueryRequest) (response *DistributeAccreditQueryResponse, err error) {
    if request == nil {
        request = NewDistributeAccreditQueryRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DistributeAccreditQuery require credential")
    }

    request.SetContext(ctx)
    
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
    return c.DistributeAccreditTlinxWithContext(context.Background(), request)
}

// DistributeAccreditTlinx
// 云支付-分账授权申请接口
//
// 可能返回的错误码:
//  FAILEDOPERATION_ACTIONINVALID = "FailedOperation.ActionInvalid"
//  FAILEDOPERATION_PABANKERROR = "FailedOperation.PABankError"
func (c *Client) DistributeAccreditTlinxWithContext(ctx context.Context, request *DistributeAccreditTlinxRequest) (response *DistributeAccreditTlinxResponse, err error) {
    if request == nil {
        request = NewDistributeAccreditTlinxRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DistributeAccreditTlinx require credential")
    }

    request.SetContext(ctx)
    
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
    return c.DistributeAddReceiverWithContext(context.Background(), request)
}

// DistributeAddReceiver
// 云支付-分账添加分账接收方接口
//
// 可能返回的错误码:
//  FAILEDOPERATION_ACTIONINVALID = "FailedOperation.ActionInvalid"
//  FAILEDOPERATION_PABANKERROR = "FailedOperation.PABankError"
func (c *Client) DistributeAddReceiverWithContext(ctx context.Context, request *DistributeAddReceiverRequest) (response *DistributeAddReceiverResponse, err error) {
    if request == nil {
        request = NewDistributeAddReceiverRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DistributeAddReceiver require credential")
    }

    request.SetContext(ctx)
    
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
    return c.DistributeApplyWithContext(context.Background(), request)
}

// DistributeApply
// 云支付-分账请求接口
//
// 可能返回的错误码:
//  FAILEDOPERATION_ACTIONINVALID = "FailedOperation.ActionInvalid"
//  FAILEDOPERATION_PABANKERROR = "FailedOperation.PABankError"
func (c *Client) DistributeApplyWithContext(ctx context.Context, request *DistributeApplyRequest) (response *DistributeApplyResponse, err error) {
    if request == nil {
        request = NewDistributeApplyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DistributeApply require credential")
    }

    request.SetContext(ctx)
    
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
    return c.DistributeCancelWithContext(context.Background(), request)
}

// DistributeCancel
// 云支付-分账撤销接口
//
// 可能返回的错误码:
//  FAILEDOPERATION_ACTIONINVALID = "FailedOperation.ActionInvalid"
//  FAILEDOPERATION_PABANKERROR = "FailedOperation.PABankError"
func (c *Client) DistributeCancelWithContext(ctx context.Context, request *DistributeCancelRequest) (response *DistributeCancelResponse, err error) {
    if request == nil {
        request = NewDistributeCancelRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DistributeCancel require credential")
    }

    request.SetContext(ctx)
    
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
    return c.DistributeQueryWithContext(context.Background(), request)
}

// DistributeQuery
// 云支付-分账结果查询接口
//
// 可能返回的错误码:
//  FAILEDOPERATION_ACTIONINVALID = "FailedOperation.ActionInvalid"
//  FAILEDOPERATION_PABANKERROR = "FailedOperation.PABankError"
func (c *Client) DistributeQueryWithContext(ctx context.Context, request *DistributeQueryRequest) (response *DistributeQueryResponse, err error) {
    if request == nil {
        request = NewDistributeQueryRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DistributeQuery require credential")
    }

    request.SetContext(ctx)
    
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
// 云支付-查询已添加分账接收方接口
//
// 可能返回的错误码:
//  FAILEDOPERATION_ACTIONINVALID = "FailedOperation.ActionInvalid"
//  FAILEDOPERATION_PABANKERROR = "FailedOperation.PABankError"
func (c *Client) DistributeQueryReceiver(request *DistributeQueryReceiverRequest) (response *DistributeQueryReceiverResponse, err error) {
    return c.DistributeQueryReceiverWithContext(context.Background(), request)
}

// DistributeQueryReceiver
// 云支付-查询已添加分账接收方接口
//
// 可能返回的错误码:
//  FAILEDOPERATION_ACTIONINVALID = "FailedOperation.ActionInvalid"
//  FAILEDOPERATION_PABANKERROR = "FailedOperation.PABankError"
func (c *Client) DistributeQueryReceiverWithContext(ctx context.Context, request *DistributeQueryReceiverRequest) (response *DistributeQueryReceiverResponse, err error) {
    if request == nil {
        request = NewDistributeQueryReceiverRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DistributeQueryReceiver require credential")
    }

    request.SetContext(ctx)
    
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
    return c.DistributeRemoveReceiverWithContext(context.Background(), request)
}

// DistributeRemoveReceiver
// 云支付-分账解除分账接收方接口
//
// 可能返回的错误码:
//  FAILEDOPERATION_ACTIONINVALID = "FailedOperation.ActionInvalid"
//  FAILEDOPERATION_PABANKERROR = "FailedOperation.PABankError"
func (c *Client) DistributeRemoveReceiverWithContext(ctx context.Context, request *DistributeRemoveReceiverRequest) (response *DistributeRemoveReceiverResponse, err error) {
    if request == nil {
        request = NewDistributeRemoveReceiverRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DistributeRemoveReceiver require credential")
    }

    request.SetContext(ctx)
    
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
    return c.DownloadBillWithContext(context.Background(), request)
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
func (c *Client) DownloadBillWithContext(ctx context.Context, request *DownloadBillRequest) (response *DownloadBillResponse, err error) {
    if request == nil {
        request = NewDownloadBillRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DownloadBill require credential")
    }

    request.SetContext(ctx)
    
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
    return c.DownloadOrgFileWithContext(context.Background(), request)
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
func (c *Client) DownloadOrgFileWithContext(ctx context.Context, request *DownloadOrgFileRequest) (response *DownloadOrgFileResponse, err error) {
    if request == nil {
        request = NewDownloadOrgFileRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DownloadOrgFile require credential")
    }

    request.SetContext(ctx)
    
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
    return c.DownloadReconciliationUrlWithContext(context.Background(), request)
}

// DownloadReconciliationUrl
// 获取对账中心账单下载地址的接口
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDPARAMETER = "FailedOperation.InvalidParameter"
//  FAILEDOPERATION_ISEMPTY = "FailedOperation.IsEmpty"
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
func (c *Client) DownloadReconciliationUrlWithContext(ctx context.Context, request *DownloadReconciliationUrlRequest) (response *DownloadReconciliationUrlResponse, err error) {
    if request == nil {
        request = NewDownloadReconciliationUrlRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DownloadReconciliationUrl require credential")
    }

    request.SetContext(ctx)
    
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
    return c.ExecuteMemberTransactionWithContext(context.Background(), request)
}

// ExecuteMemberTransaction
// 会员间交易接口
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ACTIONINVALID = "FailedOperation.ActionInvalid"
//  FAILEDOPERATION_PABANKERROR = "FailedOperation.PABankError"
func (c *Client) ExecuteMemberTransactionWithContext(ctx context.Context, request *ExecuteMemberTransactionRequest) (response *ExecuteMemberTransactionResponse, err error) {
    if request == nil {
        request = NewExecuteMemberTransactionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ExecuteMemberTransaction require credential")
    }

    request.SetContext(ctx)
    
    response = NewExecuteMemberTransactionResponse()
    err = c.Send(request, response)
    return
}

func NewFreezeFlexBalanceRequest() (request *FreezeFlexBalanceRequest) {
    request = &FreezeFlexBalanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cpdp", APIVersion, "FreezeFlexBalance")
    
    
    return
}

func NewFreezeFlexBalanceResponse() (response *FreezeFlexBalanceResponse) {
    response = &FreezeFlexBalanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// FreezeFlexBalance
// 灵云V2-冻结余额
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ACTIONINVALID = "FailedOperation.ActionInvalid"
//  FAILEDOPERATION_PABANKERROR = "FailedOperation.PABankError"
func (c *Client) FreezeFlexBalance(request *FreezeFlexBalanceRequest) (response *FreezeFlexBalanceResponse, err error) {
    return c.FreezeFlexBalanceWithContext(context.Background(), request)
}

// FreezeFlexBalance
// 灵云V2-冻结余额
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ACTIONINVALID = "FailedOperation.ActionInvalid"
//  FAILEDOPERATION_PABANKERROR = "FailedOperation.PABankError"
func (c *Client) FreezeFlexBalanceWithContext(ctx context.Context, request *FreezeFlexBalanceRequest) (response *FreezeFlexBalanceResponse, err error) {
    if request == nil {
        request = NewFreezeFlexBalanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("FreezeFlexBalance require credential")
    }

    request.SetContext(ctx)
    
    response = NewFreezeFlexBalanceResponse()
    err = c.Send(request, response)
    return
}

func NewGetBillDownloadUrlRequest() (request *GetBillDownloadUrlRequest) {
    request = &GetBillDownloadUrlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cpdp", APIVersion, "GetBillDownloadUrl")
    
    
    return
}

func NewGetBillDownloadUrlResponse() (response *GetBillDownloadUrlResponse) {
    response = &GetBillDownloadUrlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetBillDownloadUrl
// 调用该接口返回对账单下载地址，对账单下载URL通过GET方式访问，返回zip包，解压后为csv格式文件。文件首行如下：
//
// 订单号,订单归属日期,机构编号,订单描述,交易类型,订单状态,支付场景,原始金额,折扣金额,实际交易金额,支付渠道优惠金额,抹零金额,币种,下单时间,付款成功时间,商户编号,门店编号,付款方式编号,付款方式名称,商户手续费T1,商户扣率,是否信用卡交易,原始订单号,用户账号,外部订单号,订单备注
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ACTIONINVALID = "FailedOperation.ActionInvalid"
//  FAILEDOPERATION_PABANKERROR = "FailedOperation.PABankError"
func (c *Client) GetBillDownloadUrl(request *GetBillDownloadUrlRequest) (response *GetBillDownloadUrlResponse, err error) {
    return c.GetBillDownloadUrlWithContext(context.Background(), request)
}

// GetBillDownloadUrl
// 调用该接口返回对账单下载地址，对账单下载URL通过GET方式访问，返回zip包，解压后为csv格式文件。文件首行如下：
//
// 订单号,订单归属日期,机构编号,订单描述,交易类型,订单状态,支付场景,原始金额,折扣金额,实际交易金额,支付渠道优惠金额,抹零金额,币种,下单时间,付款成功时间,商户编号,门店编号,付款方式编号,付款方式名称,商户手续费T1,商户扣率,是否信用卡交易,原始订单号,用户账号,外部订单号,订单备注
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ACTIONINVALID = "FailedOperation.ActionInvalid"
//  FAILEDOPERATION_PABANKERROR = "FailedOperation.PABankError"
func (c *Client) GetBillDownloadUrlWithContext(ctx context.Context, request *GetBillDownloadUrlRequest) (response *GetBillDownloadUrlResponse, err error) {
    if request == nil {
        request = NewGetBillDownloadUrlRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetBillDownloadUrl require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetBillDownloadUrlResponse()
    err = c.Send(request, response)
    return
}

func NewGetDistributeBillDownloadUrlRequest() (request *GetDistributeBillDownloadUrlRequest) {
    request = &GetDistributeBillDownloadUrlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cpdp", APIVersion, "GetDistributeBillDownloadUrl")
    
    
    return
}

func NewGetDistributeBillDownloadUrlResponse() (response *GetDistributeBillDownloadUrlResponse) {
    response = &GetDistributeBillDownloadUrlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetDistributeBillDownloadUrl
// 调用该接口返回对账单下载地址，对账单下载URL通过GET方式访问，返回zip包，解压后为csv格式文件。文件首行如下：
//
// 商户号,订单号,支付订单号,分账订单总金额,分账详情（通过|分割每笔明细：商户号1#分账金额1|商户号2#分账金额2）,交易手续费承担方商户号,交易手续费,发起时间,分账状态,结算日期,非交易主体分账金额,商户退款订单号,商户分账单号
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ACTIONINVALID = "FailedOperation.ActionInvalid"
//  FAILEDOPERATION_PABANKERROR = "FailedOperation.PABankError"
func (c *Client) GetDistributeBillDownloadUrl(request *GetDistributeBillDownloadUrlRequest) (response *GetDistributeBillDownloadUrlResponse, err error) {
    return c.GetDistributeBillDownloadUrlWithContext(context.Background(), request)
}

// GetDistributeBillDownloadUrl
// 调用该接口返回对账单下载地址，对账单下载URL通过GET方式访问，返回zip包，解压后为csv格式文件。文件首行如下：
//
// 商户号,订单号,支付订单号,分账订单总金额,分账详情（通过|分割每笔明细：商户号1#分账金额1|商户号2#分账金额2）,交易手续费承担方商户号,交易手续费,发起时间,分账状态,结算日期,非交易主体分账金额,商户退款订单号,商户分账单号
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ACTIONINVALID = "FailedOperation.ActionInvalid"
//  FAILEDOPERATION_PABANKERROR = "FailedOperation.PABankError"
func (c *Client) GetDistributeBillDownloadUrlWithContext(ctx context.Context, request *GetDistributeBillDownloadUrlRequest) (response *GetDistributeBillDownloadUrlResponse, err error) {
    if request == nil {
        request = NewGetDistributeBillDownloadUrlRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetDistributeBillDownloadUrl require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetDistributeBillDownloadUrlResponse()
    err = c.Send(request, response)
    return
}

func NewGetPayRollAuthRequest() (request *GetPayRollAuthRequest) {
    request = &GetPayRollAuthRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cpdp", APIVersion, "GetPayRollAuth")
    
    
    return
}

func NewGetPayRollAuthResponse() (response *GetPayRollAuthResponse) {
    response = &GetPayRollAuthResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetPayRollAuth
// 务工卡-查询授权关系
//
// 可能返回的错误码:
//  AUTHFAILURE_SIGN = "AuthFailure.Sign"
//  FAILEDOPERATION_ACTION = "FailedOperation.Action"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  MISSINGPARAMETER_ACTION = "MissingParameter.Action"
//  MISSINGPARAMETER_APPID = "MissingParameter.AppId"
//  RESOURCENOTFOUND_ACCOUNT = "ResourceNotFound.Account"
//  RESOURCENOTFOUND_KEY = "ResourceNotFound.Key"
func (c *Client) GetPayRollAuth(request *GetPayRollAuthRequest) (response *GetPayRollAuthResponse, err error) {
    return c.GetPayRollAuthWithContext(context.Background(), request)
}

// GetPayRollAuth
// 务工卡-查询授权关系
//
// 可能返回的错误码:
//  AUTHFAILURE_SIGN = "AuthFailure.Sign"
//  FAILEDOPERATION_ACTION = "FailedOperation.Action"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  MISSINGPARAMETER_ACTION = "MissingParameter.Action"
//  MISSINGPARAMETER_APPID = "MissingParameter.AppId"
//  RESOURCENOTFOUND_ACCOUNT = "ResourceNotFound.Account"
//  RESOURCENOTFOUND_KEY = "ResourceNotFound.Key"
func (c *Client) GetPayRollAuthWithContext(ctx context.Context, request *GetPayRollAuthRequest) (response *GetPayRollAuthResponse, err error) {
    if request == nil {
        request = NewGetPayRollAuthRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetPayRollAuth require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetPayRollAuthResponse()
    err = c.Send(request, response)
    return
}

func NewGetPayRollAuthListRequest() (request *GetPayRollAuthListRequest) {
    request = &GetPayRollAuthListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cpdp", APIVersion, "GetPayRollAuthList")
    
    
    return
}

func NewGetPayRollAuthListResponse() (response *GetPayRollAuthListResponse) {
    response = &GetPayRollAuthListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetPayRollAuthList
// 务工卡-查询核身记录
//
// 可能返回的错误码:
//  AUTHFAILURE_SIGN = "AuthFailure.Sign"
//  FAILEDOPERATION_ACTION = "FailedOperation.Action"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  MISSINGPARAMETER_ACTION = "MissingParameter.Action"
//  MISSINGPARAMETER_APPID = "MissingParameter.AppId"
//  RESOURCENOTFOUND_ACCOUNT = "ResourceNotFound.Account"
//  RESOURCENOTFOUND_KEY = "ResourceNotFound.Key"
func (c *Client) GetPayRollAuthList(request *GetPayRollAuthListRequest) (response *GetPayRollAuthListResponse, err error) {
    return c.GetPayRollAuthListWithContext(context.Background(), request)
}

// GetPayRollAuthList
// 务工卡-查询核身记录
//
// 可能返回的错误码:
//  AUTHFAILURE_SIGN = "AuthFailure.Sign"
//  FAILEDOPERATION_ACTION = "FailedOperation.Action"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  MISSINGPARAMETER_ACTION = "MissingParameter.Action"
//  MISSINGPARAMETER_APPID = "MissingParameter.AppId"
//  RESOURCENOTFOUND_ACCOUNT = "ResourceNotFound.Account"
//  RESOURCENOTFOUND_KEY = "ResourceNotFound.Key"
func (c *Client) GetPayRollAuthListWithContext(ctx context.Context, request *GetPayRollAuthListRequest) (response *GetPayRollAuthListResponse, err error) {
    if request == nil {
        request = NewGetPayRollAuthListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetPayRollAuthList require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetPayRollAuthListResponse()
    err = c.Send(request, response)
    return
}

func NewGetPayRollAuthResultRequest() (request *GetPayRollAuthResultRequest) {
    request = &GetPayRollAuthResultRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cpdp", APIVersion, "GetPayRollAuthResult")
    
    
    return
}

func NewGetPayRollAuthResultResponse() (response *GetPayRollAuthResultResponse) {
    response = &GetPayRollAuthResultResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetPayRollAuthResult
// 务工卡-获取核身结果
//
// 可能返回的错误码:
//  AUTHFAILURE_SIGN = "AuthFailure.Sign"
//  FAILEDOPERATION_ACTION = "FailedOperation.Action"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  MISSINGPARAMETER_ACTION = "MissingParameter.Action"
//  MISSINGPARAMETER_APPID = "MissingParameter.AppId"
//  RESOURCENOTFOUND_ACCOUNT = "ResourceNotFound.Account"
//  RESOURCENOTFOUND_KEY = "ResourceNotFound.Key"
func (c *Client) GetPayRollAuthResult(request *GetPayRollAuthResultRequest) (response *GetPayRollAuthResultResponse, err error) {
    return c.GetPayRollAuthResultWithContext(context.Background(), request)
}

// GetPayRollAuthResult
// 务工卡-获取核身结果
//
// 可能返回的错误码:
//  AUTHFAILURE_SIGN = "AuthFailure.Sign"
//  FAILEDOPERATION_ACTION = "FailedOperation.Action"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  MISSINGPARAMETER_ACTION = "MissingParameter.Action"
//  MISSINGPARAMETER_APPID = "MissingParameter.AppId"
//  RESOURCENOTFOUND_ACCOUNT = "ResourceNotFound.Account"
//  RESOURCENOTFOUND_KEY = "ResourceNotFound.Key"
func (c *Client) GetPayRollAuthResultWithContext(ctx context.Context, request *GetPayRollAuthResultRequest) (response *GetPayRollAuthResultResponse, err error) {
    if request == nil {
        request = NewGetPayRollAuthResultRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetPayRollAuthResult require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetPayRollAuthResultResponse()
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
    return c.MigrateOrderRefundWithContext(context.Background(), request)
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
func (c *Client) MigrateOrderRefundWithContext(ctx context.Context, request *MigrateOrderRefundRequest) (response *MigrateOrderRefundResponse, err error) {
    if request == nil {
        request = NewMigrateOrderRefundRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("MigrateOrderRefund require credential")
    }

    request.SetContext(ctx)
    
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
    return c.MigrateOrderRefundQueryWithContext(context.Background(), request)
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
func (c *Client) MigrateOrderRefundQueryWithContext(ctx context.Context, request *MigrateOrderRefundQueryRequest) (response *MigrateOrderRefundQueryResponse, err error) {
    if request == nil {
        request = NewMigrateOrderRefundQueryRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("MigrateOrderRefundQuery require credential")
    }

    request.SetContext(ctx)
    
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
    return c.ModifyAgentTaxPaymentInfoWithContext(context.Background(), request)
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
func (c *Client) ModifyAgentTaxPaymentInfoWithContext(ctx context.Context, request *ModifyAgentTaxPaymentInfoRequest) (response *ModifyAgentTaxPaymentInfoResponse, err error) {
    if request == nil {
        request = NewModifyAgentTaxPaymentInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAgentTaxPaymentInfo require credential")
    }

    request.SetContext(ctx)
    
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
    return c.ModifyBindedAccountWithContext(context.Background(), request)
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
func (c *Client) ModifyBindedAccountWithContext(ctx context.Context, request *ModifyBindedAccountRequest) (response *ModifyBindedAccountResponse, err error) {
    if request == nil {
        request = NewModifyBindedAccountRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyBindedAccount require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyBindedAccountResponse()
    err = c.Send(request, response)
    return
}

func NewModifyFlexFundingAccountRequest() (request *ModifyFlexFundingAccountRequest) {
    request = &ModifyFlexFundingAccountRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cpdp", APIVersion, "ModifyFlexFundingAccount")
    
    
    return
}

func NewModifyFlexFundingAccountResponse() (response *ModifyFlexFundingAccountResponse) {
    response = &ModifyFlexFundingAccountResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyFlexFundingAccount
// 灵云V2-修改收款用户资金账号信息
//
// 可能返回的错误码:
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
func (c *Client) ModifyFlexFundingAccount(request *ModifyFlexFundingAccountRequest) (response *ModifyFlexFundingAccountResponse, err error) {
    return c.ModifyFlexFundingAccountWithContext(context.Background(), request)
}

// ModifyFlexFundingAccount
// 灵云V2-修改收款用户资金账号信息
//
// 可能返回的错误码:
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
func (c *Client) ModifyFlexFundingAccountWithContext(ctx context.Context, request *ModifyFlexFundingAccountRequest) (response *ModifyFlexFundingAccountResponse, err error) {
    if request == nil {
        request = NewModifyFlexFundingAccountRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyFlexFundingAccount require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyFlexFundingAccountResponse()
    err = c.Send(request, response)
    return
}

func NewModifyFlexPayeeAccountRightStatusRequest() (request *ModifyFlexPayeeAccountRightStatusRequest) {
    request = &ModifyFlexPayeeAccountRightStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cpdp", APIVersion, "ModifyFlexPayeeAccountRightStatus")
    
    
    return
}

func NewModifyFlexPayeeAccountRightStatusResponse() (response *ModifyFlexPayeeAccountRightStatusResponse) {
    response = &ModifyFlexPayeeAccountRightStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyFlexPayeeAccountRightStatus
// 灵云V2-收款用户账户权益状态修改
//
// 可能返回的错误码:
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
func (c *Client) ModifyFlexPayeeAccountRightStatus(request *ModifyFlexPayeeAccountRightStatusRequest) (response *ModifyFlexPayeeAccountRightStatusResponse, err error) {
    return c.ModifyFlexPayeeAccountRightStatusWithContext(context.Background(), request)
}

// ModifyFlexPayeeAccountRightStatus
// 灵云V2-收款用户账户权益状态修改
//
// 可能返回的错误码:
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
func (c *Client) ModifyFlexPayeeAccountRightStatusWithContext(ctx context.Context, request *ModifyFlexPayeeAccountRightStatusRequest) (response *ModifyFlexPayeeAccountRightStatusResponse, err error) {
    if request == nil {
        request = NewModifyFlexPayeeAccountRightStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyFlexPayeeAccountRightStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyFlexPayeeAccountRightStatusResponse()
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
    return c.ModifyMerchantWithContext(context.Background(), request)
}

// ModifyMerchant
// 云鉴-商户信息修改的接口
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDPARAMETER = "FailedOperation.InvalidParameter"
//  FAILEDOPERATION_MERCHANTNOTEXIST = "FailedOperation.MerchantNotExist"
//  FAILEDOPERATION_MODIFYMERCHANTFAILED = "FailedOperation.ModifyMerchantFailed"
func (c *Client) ModifyMerchantWithContext(ctx context.Context, request *ModifyMerchantRequest) (response *ModifyMerchantResponse, err error) {
    if request == nil {
        request = NewModifyMerchantRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyMerchant require credential")
    }

    request.SetContext(ctx)
    
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
    return c.ModifyMntMbrBindRelateAcctBankCodeWithContext(context.Background(), request)
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
func (c *Client) ModifyMntMbrBindRelateAcctBankCodeWithContext(ctx context.Context, request *ModifyMntMbrBindRelateAcctBankCodeRequest) (response *ModifyMntMbrBindRelateAcctBankCodeResponse, err error) {
    if request == nil {
        request = NewModifyMntMbrBindRelateAcctBankCodeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyMntMbrBindRelateAcctBankCode require credential")
    }

    request.SetContext(ctx)
    
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
//  AUTHFAILURE_SECRETKEYNOTFOUND = "AuthFailure.SecretKeyNotFound"
//  AUTHFAILURE_VERIFYERROR = "AuthFailure.VerifyError"
//  FAILEDOPERATION_BANKFAILED = "FailedOperation.BankFailed"
//  FAILEDOPERATION_SERVICEERROR = "FailedOperation.ServiceError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER_ = "MissingParameter."
func (c *Client) QueryAcctBinding(request *QueryAcctBindingRequest) (response *QueryAcctBindingResponse, err error) {
    return c.QueryAcctBindingWithContext(context.Background(), request)
}

// QueryAcctBinding
// 聚鑫-查询子账户绑定银行卡
//
// 可能返回的错误码:
//  AUTHFAILURE_SECRETKEYNOTFOUND = "AuthFailure.SecretKeyNotFound"
//  AUTHFAILURE_VERIFYERROR = "AuthFailure.VerifyError"
//  FAILEDOPERATION_BANKFAILED = "FailedOperation.BankFailed"
//  FAILEDOPERATION_SERVICEERROR = "FailedOperation.ServiceError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER_ = "MissingParameter."
func (c *Client) QueryAcctBindingWithContext(ctx context.Context, request *QueryAcctBindingRequest) (response *QueryAcctBindingResponse, err error) {
    if request == nil {
        request = NewQueryAcctBindingRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryAcctBinding require credential")
    }

    request.SetContext(ctx)
    
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
//  AUTHFAILURE_SECRETKEYNOTFOUND = "AuthFailure.SecretKeyNotFound"
//  AUTHFAILURE_VERIFYERROR = "AuthFailure.VerifyError"
//  FAILEDOPERATION_BANKFAILED = "FailedOperation.BankFailed"
//  FAILEDOPERATION_SERVICEERROR = "FailedOperation.ServiceError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER_ = "MissingParameter."
func (c *Client) QueryAcctInfo(request *QueryAcctInfoRequest) (response *QueryAcctInfoResponse, err error) {
    return c.QueryAcctInfoWithContext(context.Background(), request)
}

// QueryAcctInfo
// 聚鑫-开户信息查询
//
// 可能返回的错误码:
//  AUTHFAILURE_SECRETKEYNOTFOUND = "AuthFailure.SecretKeyNotFound"
//  AUTHFAILURE_VERIFYERROR = "AuthFailure.VerifyError"
//  FAILEDOPERATION_BANKFAILED = "FailedOperation.BankFailed"
//  FAILEDOPERATION_SERVICEERROR = "FailedOperation.ServiceError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER_ = "MissingParameter."
func (c *Client) QueryAcctInfoWithContext(ctx context.Context, request *QueryAcctInfoRequest) (response *QueryAcctInfoResponse, err error) {
    if request == nil {
        request = NewQueryAcctInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryAcctInfo require credential")
    }

    request.SetContext(ctx)
    
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
//  AUTHFAILURE_SECRETKEYNOTFOUND = "AuthFailure.SecretKeyNotFound"
//  FAILEDOPERATION_BANKFAILED = "FailedOperation.BankFailed"
//  FAILEDOPERATION_SERVICEERROR = "FailedOperation.ServiceError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER_ = "MissingParameter."
func (c *Client) QueryAcctInfoList(request *QueryAcctInfoListRequest) (response *QueryAcctInfoListResponse, err error) {
    return c.QueryAcctInfoListWithContext(context.Background(), request)
}

// QueryAcctInfoList
// 聚鑫-开户信息列表查询, 查询某一段时间的开户信息
//
// 可能返回的错误码:
//  AUTHFAILURE_SECRETKEYNOTFOUND = "AuthFailure.SecretKeyNotFound"
//  FAILEDOPERATION_BANKFAILED = "FailedOperation.BankFailed"
//  FAILEDOPERATION_SERVICEERROR = "FailedOperation.ServiceError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER_ = "MissingParameter."
func (c *Client) QueryAcctInfoListWithContext(ctx context.Context, request *QueryAcctInfoListRequest) (response *QueryAcctInfoListResponse, err error) {
    if request == nil {
        request = NewQueryAcctInfoListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryAcctInfoList require credential")
    }

    request.SetContext(ctx)
    
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
    return c.QueryAgentStatementsWithContext(context.Background(), request)
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
func (c *Client) QueryAgentStatementsWithContext(ctx context.Context, request *QueryAgentStatementsRequest) (response *QueryAgentStatementsResponse, err error) {
    if request == nil {
        request = NewQueryAgentStatementsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryAgentStatements require credential")
    }

    request.SetContext(ctx)
    
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
    return c.QueryAgentTaxPaymentBatchWithContext(context.Background(), request)
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
func (c *Client) QueryAgentTaxPaymentBatchWithContext(ctx context.Context, request *QueryAgentTaxPaymentBatchRequest) (response *QueryAgentTaxPaymentBatchResponse, err error) {
    if request == nil {
        request = NewQueryAgentTaxPaymentBatchRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryAgentTaxPaymentBatch require credential")
    }

    request.SetContext(ctx)
    
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
    return c.QueryAnchorContractInfoWithContext(context.Background(), request)
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
func (c *Client) QueryAnchorContractInfoWithContext(ctx context.Context, request *QueryAnchorContractInfoRequest) (response *QueryAnchorContractInfoResponse, err error) {
    if request == nil {
        request = NewQueryAnchorContractInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryAnchorContractInfo require credential")
    }

    request.SetContext(ctx)
    
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
    return c.QueryApplicationMaterialWithContext(context.Background(), request)
}

// QueryApplicationMaterial
// 跨境-成功申报材料查询。查询成功入库的申报材料。
//
// 可能返回的错误码:
//  INTERNALERROR_PARAMETERERROR = "InternalError.ParameterError"
//  INTERNALERROR_UNKOWNERROR = "InternalError.UnkownError"
//  RESOURCENOTFOUND_MERCHANTINFONOTFOUND = "ResourceNotFound.MerchantInfoNotFound"
func (c *Client) QueryApplicationMaterialWithContext(ctx context.Context, request *QueryApplicationMaterialRequest) (response *QueryApplicationMaterialResponse, err error) {
    if request == nil {
        request = NewQueryApplicationMaterialRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryApplicationMaterial require credential")
    }

    request.SetContext(ctx)
    
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
    return c.QueryAssignmentWithContext(context.Background(), request)
}

// QueryAssignment
// 直播平台-查询分配关系
//
// 可能返回的错误码:
//  INTERNALERROR_PARAMETERERROR = "InternalError.ParameterError"
//  INTERNALERROR_UNKOWNERROR = "InternalError.UnkownError"
//  RESOURCENOTFOUND_MERCHANTINFONOTFOUND = "ResourceNotFound.MerchantInfoNotFound"
func (c *Client) QueryAssignmentWithContext(ctx context.Context, request *QueryAssignmentRequest) (response *QueryAssignmentResponse, err error) {
    if request == nil {
        request = NewQueryAssignmentRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryAssignment require credential")
    }

    request.SetContext(ctx)
    
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
//  AUTHFAILURE_SECRETKEYNOTFOUND = "AuthFailure.SecretKeyNotFound"
//  FAILEDOPERATION_BANKFAILED = "FailedOperation.BankFailed"
//  FAILEDOPERATION_SERVICEERROR = "FailedOperation.ServiceError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER_ = "MissingParameter."
func (c *Client) QueryBalance(request *QueryBalanceRequest) (response *QueryBalanceResponse, err error) {
    return c.QueryBalanceWithContext(context.Background(), request)
}

// QueryBalance
// 子商户余额查询
//
// 可能返回的错误码:
//  AUTHFAILURE_SECRETKEYNOTFOUND = "AuthFailure.SecretKeyNotFound"
//  FAILEDOPERATION_BANKFAILED = "FailedOperation.BankFailed"
//  FAILEDOPERATION_SERVICEERROR = "FailedOperation.ServiceError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER_ = "MissingParameter."
func (c *Client) QueryBalanceWithContext(ctx context.Context, request *QueryBalanceRequest) (response *QueryBalanceResponse, err error) {
    if request == nil {
        request = NewQueryBalanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryBalance require credential")
    }

    request.SetContext(ctx)
    
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
    return c.QueryBankClearWithContext(context.Background(), request)
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
func (c *Client) QueryBankClearWithContext(ctx context.Context, request *QueryBankClearRequest) (response *QueryBankClearResponse, err error) {
    if request == nil {
        request = NewQueryBankClearRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryBankClear require credential")
    }

    request.SetContext(ctx)
    
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
    return c.QueryBankTransactionDetailsWithContext(context.Background(), request)
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
func (c *Client) QueryBankTransactionDetailsWithContext(ctx context.Context, request *QueryBankTransactionDetailsRequest) (response *QueryBankTransactionDetailsResponse, err error) {
    if request == nil {
        request = NewQueryBankTransactionDetailsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryBankTransactionDetails require credential")
    }

    request.SetContext(ctx)
    
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
    return c.QueryBankWithdrawCashDetailsWithContext(context.Background(), request)
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
func (c *Client) QueryBankWithdrawCashDetailsWithContext(ctx context.Context, request *QueryBankWithdrawCashDetailsRequest) (response *QueryBankWithdrawCashDetailsResponse, err error) {
    if request == nil {
        request = NewQueryBankWithdrawCashDetailsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryBankWithdrawCashDetails require credential")
    }

    request.SetContext(ctx)
    
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
    return c.QueryBatchPaymentResultWithContext(context.Background(), request)
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
func (c *Client) QueryBatchPaymentResultWithContext(ctx context.Context, request *QueryBatchPaymentResultRequest) (response *QueryBatchPaymentResultResponse, err error) {
    if request == nil {
        request = NewQueryBatchPaymentResultRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryBatchPaymentResult require credential")
    }

    request.SetContext(ctx)
    
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
    return c.QueryBillDownloadURLWithContext(context.Background(), request)
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
func (c *Client) QueryBillDownloadURLWithContext(ctx context.Context, request *QueryBillDownloadURLRequest) (response *QueryBillDownloadURLResponse, err error) {
    if request == nil {
        request = NewQueryBillDownloadURLRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryBillDownloadURL require credential")
    }

    request.SetContext(ctx)
    
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
    return c.QueryCityCodeWithContext(context.Background(), request)
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
func (c *Client) QueryCityCodeWithContext(ctx context.Context, request *QueryCityCodeRequest) (response *QueryCityCodeResponse, err error) {
    if request == nil {
        request = NewQueryCityCodeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryCityCode require credential")
    }

    request.SetContext(ctx)
    
    response = NewQueryCityCodeResponse()
    err = c.Send(request, response)
    return
}

func NewQueryCloudChannelDataRequest() (request *QueryCloudChannelDataRequest) {
    request = &QueryCloudChannelDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cpdp", APIVersion, "QueryCloudChannelData")
    
    
    return
}

func NewQueryCloudChannelDataResponse() (response *QueryCloudChannelDataResponse) {
    response = &QueryCloudChannelDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// QueryCloudChannelData
// 发起支付等渠道操作后，可以调用该接口查询渠道的数据。
//
// 可能返回的错误码:
//  AUTHFAILURE_MIDAS = "AuthFailure.Midas"
//  FAILEDOPERATION_ACTION = "FailedOperation.Action"
//  FAILEDOPERATION_MIDASINTERNALERROR = "FailedOperation.MidasInternalError"
//  FAILEDOPERATION_MIDASINVALIDREQUEST = "FailedOperation.MidasInvalidRequest"
//  FAILEDOPERATION_MIDASNEEDRETRY = "FailedOperation.MidasNeedRetry"
//  FAILEDOPERATION_MIDASREGISTERUNFINISHED = "FailedOperation.MidasRegisterUnfinished"
//  FAILEDOPERATION_MIDASREPEATORDER = "FailedOperation.MidasRepeatOrder"
//  FAILEDOPERATION_MIDASRISK = "FailedOperation.MidasRisk"
//  FAILEDOPERATION_MIDASSTATUSNOTMATCH = "FailedOperation.MidasStatusNotMatch"
//  FAILEDOPERATION_MIDASUNSUPPORTEDACTION = "FailedOperation.MidasUnsupportedAction"
//  INTERNALERROR_MIDAS = "InternalError.Midas"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  INVALIDPARAMETER_MIDAS = "InvalidParameter.Midas"
//  INVALIDPARAMETER_MIDASENVIRONMENT = "InvalidParameter.MidasEnvironment"
//  INVALIDPARAMETER_MIDASEXTERNALAPP = "InvalidParameter.MidasExternalApp"
//  INVALIDPARAMETER_MIDASFILETYPE = "InvalidParameter.MidasFileType"
//  INVALIDPARAMETER_MIDASHASH = "InvalidParameter.MidasHash"
//  INVALIDPARAMETER_MIDASSIGNID = "InvalidParameter.MidasSignId"
//  LIMITEXCEEDED_MIDASLARGEFILE = "LimitExceeded.MidasLargeFile"
//  LIMITEXCEEDED_MIDASORDER = "LimitExceeded.MidasOrder"
//  LIMITEXCEEDED_MIDASORDERCANCELED = "LimitExceeded.MidasOrderCanceled"
//  LIMITEXCEEDED_MIDASORDERCLOSED = "LimitExceeded.MidasOrderClosed"
//  LIMITEXCEEDED_MIDASORDEREXPIRED = "LimitExceeded.MidasOrderExpired"
//  LIMITEXCEEDED_MIDASORDERFAILED = "LimitExceeded.MidasOrderFailed"
//  LIMITEXCEEDED_MIDASORDERPARTIALSUCCESS = "LimitExceeded.MidasOrderPartialSuccess"
//  LIMITEXCEEDED_MIDASORDERPROCESSING = "LimitExceeded.MidasOrderProcessing"
//  LIMITEXCEEDED_MIDASORDERSUCCESS = "LimitExceeded.MidasOrderSuccess"
//  REQUESTLIMITEXCEEDED_MIDAS = "RequestLimitExceeded.Midas"
//  REQUESTLIMITEXCEEDED_MIDASINVALIDREQUEST = "RequestLimitExceeded.MidasInvalidRequest"
//  RESOURCEINUSE_MIDAS = "ResourceInUse.Midas"
//  RESOURCENOTFOUND_ACCOUNT = "ResourceNotFound.Account"
//  RESOURCENOTFOUND_MIDASEXTERNALAPP = "ResourceNotFound.MidasExternalApp"
//  RESOURCENOTFOUND_MIDASEXTERNALORDER = "ResourceNotFound.MidasExternalOrder"
//  RESOURCENOTFOUND_MIDASORDER = "ResourceNotFound.MidasOrder"
//  RESOURCENOTFOUND_MIDASSIGN = "ResourceNotFound.MidasSign"
//  RESOURCEUNAVAILABLE_MIDASBALANCE = "ResourceUnavailable.MidasBalance"
//  RESOURCEUNAVAILABLE_MIDASDAY = "ResourceUnavailable.MidasDay"
//  RESOURCEUNAVAILABLE_MIDASFROZENAMOUNT = "ResourceUnavailable.MidasFrozenAmount"
//  RESOURCEUNAVAILABLE_MIDASMERCHANTBALANCE = "ResourceUnavailable.MidasMerchantBalance"
//  RESOURCEUNAVAILABLE_MIDASORDER = "ResourceUnavailable.MidasOrder"
//  RESOURCEUNAVAILABLE_MIDASUSERBALANCE = "ResourceUnavailable.MidasUserBalance"
//  RESOURCEUNAVAILABLE_MIDASWALLET = "ResourceUnavailable.MidasWallet"
//  UNAUTHORIZEDOPERATION_MIDAS = "UnauthorizedOperation.Midas"
func (c *Client) QueryCloudChannelData(request *QueryCloudChannelDataRequest) (response *QueryCloudChannelDataResponse, err error) {
    return c.QueryCloudChannelDataWithContext(context.Background(), request)
}

// QueryCloudChannelData
// 发起支付等渠道操作后，可以调用该接口查询渠道的数据。
//
// 可能返回的错误码:
//  AUTHFAILURE_MIDAS = "AuthFailure.Midas"
//  FAILEDOPERATION_ACTION = "FailedOperation.Action"
//  FAILEDOPERATION_MIDASINTERNALERROR = "FailedOperation.MidasInternalError"
//  FAILEDOPERATION_MIDASINVALIDREQUEST = "FailedOperation.MidasInvalidRequest"
//  FAILEDOPERATION_MIDASNEEDRETRY = "FailedOperation.MidasNeedRetry"
//  FAILEDOPERATION_MIDASREGISTERUNFINISHED = "FailedOperation.MidasRegisterUnfinished"
//  FAILEDOPERATION_MIDASREPEATORDER = "FailedOperation.MidasRepeatOrder"
//  FAILEDOPERATION_MIDASRISK = "FailedOperation.MidasRisk"
//  FAILEDOPERATION_MIDASSTATUSNOTMATCH = "FailedOperation.MidasStatusNotMatch"
//  FAILEDOPERATION_MIDASUNSUPPORTEDACTION = "FailedOperation.MidasUnsupportedAction"
//  INTERNALERROR_MIDAS = "InternalError.Midas"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  INVALIDPARAMETER_MIDAS = "InvalidParameter.Midas"
//  INVALIDPARAMETER_MIDASENVIRONMENT = "InvalidParameter.MidasEnvironment"
//  INVALIDPARAMETER_MIDASEXTERNALAPP = "InvalidParameter.MidasExternalApp"
//  INVALIDPARAMETER_MIDASFILETYPE = "InvalidParameter.MidasFileType"
//  INVALIDPARAMETER_MIDASHASH = "InvalidParameter.MidasHash"
//  INVALIDPARAMETER_MIDASSIGNID = "InvalidParameter.MidasSignId"
//  LIMITEXCEEDED_MIDASLARGEFILE = "LimitExceeded.MidasLargeFile"
//  LIMITEXCEEDED_MIDASORDER = "LimitExceeded.MidasOrder"
//  LIMITEXCEEDED_MIDASORDERCANCELED = "LimitExceeded.MidasOrderCanceled"
//  LIMITEXCEEDED_MIDASORDERCLOSED = "LimitExceeded.MidasOrderClosed"
//  LIMITEXCEEDED_MIDASORDEREXPIRED = "LimitExceeded.MidasOrderExpired"
//  LIMITEXCEEDED_MIDASORDERFAILED = "LimitExceeded.MidasOrderFailed"
//  LIMITEXCEEDED_MIDASORDERPARTIALSUCCESS = "LimitExceeded.MidasOrderPartialSuccess"
//  LIMITEXCEEDED_MIDASORDERPROCESSING = "LimitExceeded.MidasOrderProcessing"
//  LIMITEXCEEDED_MIDASORDERSUCCESS = "LimitExceeded.MidasOrderSuccess"
//  REQUESTLIMITEXCEEDED_MIDAS = "RequestLimitExceeded.Midas"
//  REQUESTLIMITEXCEEDED_MIDASINVALIDREQUEST = "RequestLimitExceeded.MidasInvalidRequest"
//  RESOURCEINUSE_MIDAS = "ResourceInUse.Midas"
//  RESOURCENOTFOUND_ACCOUNT = "ResourceNotFound.Account"
//  RESOURCENOTFOUND_MIDASEXTERNALAPP = "ResourceNotFound.MidasExternalApp"
//  RESOURCENOTFOUND_MIDASEXTERNALORDER = "ResourceNotFound.MidasExternalOrder"
//  RESOURCENOTFOUND_MIDASORDER = "ResourceNotFound.MidasOrder"
//  RESOURCENOTFOUND_MIDASSIGN = "ResourceNotFound.MidasSign"
//  RESOURCEUNAVAILABLE_MIDASBALANCE = "ResourceUnavailable.MidasBalance"
//  RESOURCEUNAVAILABLE_MIDASDAY = "ResourceUnavailable.MidasDay"
//  RESOURCEUNAVAILABLE_MIDASFROZENAMOUNT = "ResourceUnavailable.MidasFrozenAmount"
//  RESOURCEUNAVAILABLE_MIDASMERCHANTBALANCE = "ResourceUnavailable.MidasMerchantBalance"
//  RESOURCEUNAVAILABLE_MIDASORDER = "ResourceUnavailable.MidasOrder"
//  RESOURCEUNAVAILABLE_MIDASUSERBALANCE = "ResourceUnavailable.MidasUserBalance"
//  RESOURCEUNAVAILABLE_MIDASWALLET = "ResourceUnavailable.MidasWallet"
//  UNAUTHORIZEDOPERATION_MIDAS = "UnauthorizedOperation.Midas"
func (c *Client) QueryCloudChannelDataWithContext(ctx context.Context, request *QueryCloudChannelDataRequest) (response *QueryCloudChannelDataResponse, err error) {
    if request == nil {
        request = NewQueryCloudChannelDataRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryCloudChannelData require credential")
    }

    request.SetContext(ctx)
    
    response = NewQueryCloudChannelDataResponse()
    err = c.Send(request, response)
    return
}

func NewQueryCloudOrderRequest() (request *QueryCloudOrderRequest) {
    request = &QueryCloudOrderRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cpdp", APIVersion, "QueryCloudOrder")
    
    
    return
}

func NewQueryCloudOrderResponse() (response *QueryCloudOrderResponse) {
    response = &QueryCloudOrderResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// QueryCloudOrder
// 根据订单号或用户ID，查询支付订单状态。
//
// 可能返回的错误码:
//  AUTHFAILURE_MIDAS = "AuthFailure.Midas"
//  FAILEDOPERATION_ACTION = "FailedOperation.Action"
//  FAILEDOPERATION_MIDASINTERNALERROR = "FailedOperation.MidasInternalError"
//  FAILEDOPERATION_MIDASINVALIDREQUEST = "FailedOperation.MidasInvalidRequest"
//  FAILEDOPERATION_MIDASNEEDRETRY = "FailedOperation.MidasNeedRetry"
//  FAILEDOPERATION_MIDASREGISTERUNFINISHED = "FailedOperation.MidasRegisterUnfinished"
//  FAILEDOPERATION_MIDASREPEATORDER = "FailedOperation.MidasRepeatOrder"
//  FAILEDOPERATION_MIDASRISK = "FailedOperation.MidasRisk"
//  FAILEDOPERATION_MIDASSTATUSNOTMATCH = "FailedOperation.MidasStatusNotMatch"
//  FAILEDOPERATION_MIDASUNSUPPORTEDACTION = "FailedOperation.MidasUnsupportedAction"
//  INTERNALERROR_MIDAS = "InternalError.Midas"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  INVALIDPARAMETER_MIDAS = "InvalidParameter.Midas"
//  INVALIDPARAMETER_MIDASENVIRONMENT = "InvalidParameter.MidasEnvironment"
//  INVALIDPARAMETER_MIDASEXTERNALAPP = "InvalidParameter.MidasExternalApp"
//  INVALIDPARAMETER_MIDASFILETYPE = "InvalidParameter.MidasFileType"
//  INVALIDPARAMETER_MIDASHASH = "InvalidParameter.MidasHash"
//  INVALIDPARAMETER_MIDASSIGNID = "InvalidParameter.MidasSignId"
//  LIMITEXCEEDED_MIDASLARGEFILE = "LimitExceeded.MidasLargeFile"
//  LIMITEXCEEDED_MIDASORDER = "LimitExceeded.MidasOrder"
//  LIMITEXCEEDED_MIDASORDERCANCELED = "LimitExceeded.MidasOrderCanceled"
//  LIMITEXCEEDED_MIDASORDERCLOSED = "LimitExceeded.MidasOrderClosed"
//  LIMITEXCEEDED_MIDASORDEREXPIRED = "LimitExceeded.MidasOrderExpired"
//  LIMITEXCEEDED_MIDASORDERFAILED = "LimitExceeded.MidasOrderFailed"
//  LIMITEXCEEDED_MIDASORDERPARTIALSUCCESS = "LimitExceeded.MidasOrderPartialSuccess"
//  LIMITEXCEEDED_MIDASORDERPROCESSING = "LimitExceeded.MidasOrderProcessing"
//  LIMITEXCEEDED_MIDASORDERSUCCESS = "LimitExceeded.MidasOrderSuccess"
//  REQUESTLIMITEXCEEDED_MIDAS = "RequestLimitExceeded.Midas"
//  REQUESTLIMITEXCEEDED_MIDASINVALIDREQUEST = "RequestLimitExceeded.MidasInvalidRequest"
//  RESOURCEINUSE_MIDAS = "ResourceInUse.Midas"
//  RESOURCENOTFOUND_ACCOUNT = "ResourceNotFound.Account"
//  RESOURCENOTFOUND_MIDASEXTERNALAPP = "ResourceNotFound.MidasExternalApp"
//  RESOURCENOTFOUND_MIDASEXTERNALORDER = "ResourceNotFound.MidasExternalOrder"
//  RESOURCENOTFOUND_MIDASORDER = "ResourceNotFound.MidasOrder"
//  RESOURCENOTFOUND_MIDASSIGN = "ResourceNotFound.MidasSign"
//  RESOURCEUNAVAILABLE_MIDASBALANCE = "ResourceUnavailable.MidasBalance"
//  RESOURCEUNAVAILABLE_MIDASDAY = "ResourceUnavailable.MidasDay"
//  RESOURCEUNAVAILABLE_MIDASFROZENAMOUNT = "ResourceUnavailable.MidasFrozenAmount"
//  RESOURCEUNAVAILABLE_MIDASMERCHANTBALANCE = "ResourceUnavailable.MidasMerchantBalance"
//  RESOURCEUNAVAILABLE_MIDASORDER = "ResourceUnavailable.MidasOrder"
//  RESOURCEUNAVAILABLE_MIDASUSERBALANCE = "ResourceUnavailable.MidasUserBalance"
//  RESOURCEUNAVAILABLE_MIDASWALLET = "ResourceUnavailable.MidasWallet"
//  UNAUTHORIZEDOPERATION_MIDAS = "UnauthorizedOperation.Midas"
func (c *Client) QueryCloudOrder(request *QueryCloudOrderRequest) (response *QueryCloudOrderResponse, err error) {
    return c.QueryCloudOrderWithContext(context.Background(), request)
}

// QueryCloudOrder
// 根据订单号或用户ID，查询支付订单状态。
//
// 可能返回的错误码:
//  AUTHFAILURE_MIDAS = "AuthFailure.Midas"
//  FAILEDOPERATION_ACTION = "FailedOperation.Action"
//  FAILEDOPERATION_MIDASINTERNALERROR = "FailedOperation.MidasInternalError"
//  FAILEDOPERATION_MIDASINVALIDREQUEST = "FailedOperation.MidasInvalidRequest"
//  FAILEDOPERATION_MIDASNEEDRETRY = "FailedOperation.MidasNeedRetry"
//  FAILEDOPERATION_MIDASREGISTERUNFINISHED = "FailedOperation.MidasRegisterUnfinished"
//  FAILEDOPERATION_MIDASREPEATORDER = "FailedOperation.MidasRepeatOrder"
//  FAILEDOPERATION_MIDASRISK = "FailedOperation.MidasRisk"
//  FAILEDOPERATION_MIDASSTATUSNOTMATCH = "FailedOperation.MidasStatusNotMatch"
//  FAILEDOPERATION_MIDASUNSUPPORTEDACTION = "FailedOperation.MidasUnsupportedAction"
//  INTERNALERROR_MIDAS = "InternalError.Midas"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  INVALIDPARAMETER_MIDAS = "InvalidParameter.Midas"
//  INVALIDPARAMETER_MIDASENVIRONMENT = "InvalidParameter.MidasEnvironment"
//  INVALIDPARAMETER_MIDASEXTERNALAPP = "InvalidParameter.MidasExternalApp"
//  INVALIDPARAMETER_MIDASFILETYPE = "InvalidParameter.MidasFileType"
//  INVALIDPARAMETER_MIDASHASH = "InvalidParameter.MidasHash"
//  INVALIDPARAMETER_MIDASSIGNID = "InvalidParameter.MidasSignId"
//  LIMITEXCEEDED_MIDASLARGEFILE = "LimitExceeded.MidasLargeFile"
//  LIMITEXCEEDED_MIDASORDER = "LimitExceeded.MidasOrder"
//  LIMITEXCEEDED_MIDASORDERCANCELED = "LimitExceeded.MidasOrderCanceled"
//  LIMITEXCEEDED_MIDASORDERCLOSED = "LimitExceeded.MidasOrderClosed"
//  LIMITEXCEEDED_MIDASORDEREXPIRED = "LimitExceeded.MidasOrderExpired"
//  LIMITEXCEEDED_MIDASORDERFAILED = "LimitExceeded.MidasOrderFailed"
//  LIMITEXCEEDED_MIDASORDERPARTIALSUCCESS = "LimitExceeded.MidasOrderPartialSuccess"
//  LIMITEXCEEDED_MIDASORDERPROCESSING = "LimitExceeded.MidasOrderProcessing"
//  LIMITEXCEEDED_MIDASORDERSUCCESS = "LimitExceeded.MidasOrderSuccess"
//  REQUESTLIMITEXCEEDED_MIDAS = "RequestLimitExceeded.Midas"
//  REQUESTLIMITEXCEEDED_MIDASINVALIDREQUEST = "RequestLimitExceeded.MidasInvalidRequest"
//  RESOURCEINUSE_MIDAS = "ResourceInUse.Midas"
//  RESOURCENOTFOUND_ACCOUNT = "ResourceNotFound.Account"
//  RESOURCENOTFOUND_MIDASEXTERNALAPP = "ResourceNotFound.MidasExternalApp"
//  RESOURCENOTFOUND_MIDASEXTERNALORDER = "ResourceNotFound.MidasExternalOrder"
//  RESOURCENOTFOUND_MIDASORDER = "ResourceNotFound.MidasOrder"
//  RESOURCENOTFOUND_MIDASSIGN = "ResourceNotFound.MidasSign"
//  RESOURCEUNAVAILABLE_MIDASBALANCE = "ResourceUnavailable.MidasBalance"
//  RESOURCEUNAVAILABLE_MIDASDAY = "ResourceUnavailable.MidasDay"
//  RESOURCEUNAVAILABLE_MIDASFROZENAMOUNT = "ResourceUnavailable.MidasFrozenAmount"
//  RESOURCEUNAVAILABLE_MIDASMERCHANTBALANCE = "ResourceUnavailable.MidasMerchantBalance"
//  RESOURCEUNAVAILABLE_MIDASORDER = "ResourceUnavailable.MidasOrder"
//  RESOURCEUNAVAILABLE_MIDASUSERBALANCE = "ResourceUnavailable.MidasUserBalance"
//  RESOURCEUNAVAILABLE_MIDASWALLET = "ResourceUnavailable.MidasWallet"
//  UNAUTHORIZEDOPERATION_MIDAS = "UnauthorizedOperation.Midas"
func (c *Client) QueryCloudOrderWithContext(ctx context.Context, request *QueryCloudOrderRequest) (response *QueryCloudOrderResponse, err error) {
    if request == nil {
        request = NewQueryCloudOrderRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryCloudOrder require credential")
    }

    request.SetContext(ctx)
    
    response = NewQueryCloudOrderResponse()
    err = c.Send(request, response)
    return
}

func NewQueryCloudRefundOrderRequest() (request *QueryCloudRefundOrderRequest) {
    request = &QueryCloudRefundOrderRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cpdp", APIVersion, "QueryCloudRefundOrder")
    
    
    return
}

func NewQueryCloudRefundOrderResponse() (response *QueryCloudRefundOrderResponse) {
    response = &QueryCloudRefundOrderResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// QueryCloudRefundOrder
// 提交退款申请后，通过调用该接口查询退款状态。退款可能有一定延时，用微信零钱支付的退款约20分钟内到账，银行卡支付的退款约3个工作日后到账。
//
// 可能返回的错误码:
//  AUTHFAILURE_MIDAS = "AuthFailure.Midas"
//  FAILEDOPERATION_ACTION = "FailedOperation.Action"
//  FAILEDOPERATION_MIDASINTERNALERROR = "FailedOperation.MidasInternalError"
//  FAILEDOPERATION_MIDASINVALIDREQUEST = "FailedOperation.MidasInvalidRequest"
//  FAILEDOPERATION_MIDASNEEDRETRY = "FailedOperation.MidasNeedRetry"
//  FAILEDOPERATION_MIDASREGISTERUNFINISHED = "FailedOperation.MidasRegisterUnfinished"
//  FAILEDOPERATION_MIDASREPEATORDER = "FailedOperation.MidasRepeatOrder"
//  FAILEDOPERATION_MIDASRISK = "FailedOperation.MidasRisk"
//  FAILEDOPERATION_MIDASSTATUSNOTMATCH = "FailedOperation.MidasStatusNotMatch"
//  FAILEDOPERATION_MIDASUNSUPPORTEDACTION = "FailedOperation.MidasUnsupportedAction"
//  INTERNALERROR_MIDAS = "InternalError.Midas"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  INVALIDPARAMETER_MIDAS = "InvalidParameter.Midas"
//  INVALIDPARAMETER_MIDASENVIRONMENT = "InvalidParameter.MidasEnvironment"
//  INVALIDPARAMETER_MIDASEXTERNALAPP = "InvalidParameter.MidasExternalApp"
//  INVALIDPARAMETER_MIDASFILETYPE = "InvalidParameter.MidasFileType"
//  INVALIDPARAMETER_MIDASHASH = "InvalidParameter.MidasHash"
//  INVALIDPARAMETER_MIDASSIGNID = "InvalidParameter.MidasSignId"
//  LIMITEXCEEDED_MIDASLARGEFILE = "LimitExceeded.MidasLargeFile"
//  LIMITEXCEEDED_MIDASORDER = "LimitExceeded.MidasOrder"
//  LIMITEXCEEDED_MIDASORDERCANCELED = "LimitExceeded.MidasOrderCanceled"
//  LIMITEXCEEDED_MIDASORDERCLOSED = "LimitExceeded.MidasOrderClosed"
//  LIMITEXCEEDED_MIDASORDEREXPIRED = "LimitExceeded.MidasOrderExpired"
//  LIMITEXCEEDED_MIDASORDERFAILED = "LimitExceeded.MidasOrderFailed"
//  LIMITEXCEEDED_MIDASORDERPARTIALSUCCESS = "LimitExceeded.MidasOrderPartialSuccess"
//  LIMITEXCEEDED_MIDASORDERPROCESSING = "LimitExceeded.MidasOrderProcessing"
//  LIMITEXCEEDED_MIDASORDERSUCCESS = "LimitExceeded.MidasOrderSuccess"
//  REQUESTLIMITEXCEEDED_MIDAS = "RequestLimitExceeded.Midas"
//  REQUESTLIMITEXCEEDED_MIDASINVALIDREQUEST = "RequestLimitExceeded.MidasInvalidRequest"
//  RESOURCEINUSE_MIDAS = "ResourceInUse.Midas"
//  RESOURCENOTFOUND_ACCOUNT = "ResourceNotFound.Account"
//  RESOURCENOTFOUND_MIDASEXTERNALAPP = "ResourceNotFound.MidasExternalApp"
//  RESOURCENOTFOUND_MIDASEXTERNALORDER = "ResourceNotFound.MidasExternalOrder"
//  RESOURCENOTFOUND_MIDASORDER = "ResourceNotFound.MidasOrder"
//  RESOURCENOTFOUND_MIDASSIGN = "ResourceNotFound.MidasSign"
//  RESOURCEUNAVAILABLE_MIDASBALANCE = "ResourceUnavailable.MidasBalance"
//  RESOURCEUNAVAILABLE_MIDASDAY = "ResourceUnavailable.MidasDay"
//  RESOURCEUNAVAILABLE_MIDASFROZENAMOUNT = "ResourceUnavailable.MidasFrozenAmount"
//  RESOURCEUNAVAILABLE_MIDASMERCHANTBALANCE = "ResourceUnavailable.MidasMerchantBalance"
//  RESOURCEUNAVAILABLE_MIDASORDER = "ResourceUnavailable.MidasOrder"
//  RESOURCEUNAVAILABLE_MIDASUSERBALANCE = "ResourceUnavailable.MidasUserBalance"
//  RESOURCEUNAVAILABLE_MIDASWALLET = "ResourceUnavailable.MidasWallet"
//  UNAUTHORIZEDOPERATION_MIDAS = "UnauthorizedOperation.Midas"
func (c *Client) QueryCloudRefundOrder(request *QueryCloudRefundOrderRequest) (response *QueryCloudRefundOrderResponse, err error) {
    return c.QueryCloudRefundOrderWithContext(context.Background(), request)
}

// QueryCloudRefundOrder
// 提交退款申请后，通过调用该接口查询退款状态。退款可能有一定延时，用微信零钱支付的退款约20分钟内到账，银行卡支付的退款约3个工作日后到账。
//
// 可能返回的错误码:
//  AUTHFAILURE_MIDAS = "AuthFailure.Midas"
//  FAILEDOPERATION_ACTION = "FailedOperation.Action"
//  FAILEDOPERATION_MIDASINTERNALERROR = "FailedOperation.MidasInternalError"
//  FAILEDOPERATION_MIDASINVALIDREQUEST = "FailedOperation.MidasInvalidRequest"
//  FAILEDOPERATION_MIDASNEEDRETRY = "FailedOperation.MidasNeedRetry"
//  FAILEDOPERATION_MIDASREGISTERUNFINISHED = "FailedOperation.MidasRegisterUnfinished"
//  FAILEDOPERATION_MIDASREPEATORDER = "FailedOperation.MidasRepeatOrder"
//  FAILEDOPERATION_MIDASRISK = "FailedOperation.MidasRisk"
//  FAILEDOPERATION_MIDASSTATUSNOTMATCH = "FailedOperation.MidasStatusNotMatch"
//  FAILEDOPERATION_MIDASUNSUPPORTEDACTION = "FailedOperation.MidasUnsupportedAction"
//  INTERNALERROR_MIDAS = "InternalError.Midas"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  INVALIDPARAMETER_MIDAS = "InvalidParameter.Midas"
//  INVALIDPARAMETER_MIDASENVIRONMENT = "InvalidParameter.MidasEnvironment"
//  INVALIDPARAMETER_MIDASEXTERNALAPP = "InvalidParameter.MidasExternalApp"
//  INVALIDPARAMETER_MIDASFILETYPE = "InvalidParameter.MidasFileType"
//  INVALIDPARAMETER_MIDASHASH = "InvalidParameter.MidasHash"
//  INVALIDPARAMETER_MIDASSIGNID = "InvalidParameter.MidasSignId"
//  LIMITEXCEEDED_MIDASLARGEFILE = "LimitExceeded.MidasLargeFile"
//  LIMITEXCEEDED_MIDASORDER = "LimitExceeded.MidasOrder"
//  LIMITEXCEEDED_MIDASORDERCANCELED = "LimitExceeded.MidasOrderCanceled"
//  LIMITEXCEEDED_MIDASORDERCLOSED = "LimitExceeded.MidasOrderClosed"
//  LIMITEXCEEDED_MIDASORDEREXPIRED = "LimitExceeded.MidasOrderExpired"
//  LIMITEXCEEDED_MIDASORDERFAILED = "LimitExceeded.MidasOrderFailed"
//  LIMITEXCEEDED_MIDASORDERPARTIALSUCCESS = "LimitExceeded.MidasOrderPartialSuccess"
//  LIMITEXCEEDED_MIDASORDERPROCESSING = "LimitExceeded.MidasOrderProcessing"
//  LIMITEXCEEDED_MIDASORDERSUCCESS = "LimitExceeded.MidasOrderSuccess"
//  REQUESTLIMITEXCEEDED_MIDAS = "RequestLimitExceeded.Midas"
//  REQUESTLIMITEXCEEDED_MIDASINVALIDREQUEST = "RequestLimitExceeded.MidasInvalidRequest"
//  RESOURCEINUSE_MIDAS = "ResourceInUse.Midas"
//  RESOURCENOTFOUND_ACCOUNT = "ResourceNotFound.Account"
//  RESOURCENOTFOUND_MIDASEXTERNALAPP = "ResourceNotFound.MidasExternalApp"
//  RESOURCENOTFOUND_MIDASEXTERNALORDER = "ResourceNotFound.MidasExternalOrder"
//  RESOURCENOTFOUND_MIDASORDER = "ResourceNotFound.MidasOrder"
//  RESOURCENOTFOUND_MIDASSIGN = "ResourceNotFound.MidasSign"
//  RESOURCEUNAVAILABLE_MIDASBALANCE = "ResourceUnavailable.MidasBalance"
//  RESOURCEUNAVAILABLE_MIDASDAY = "ResourceUnavailable.MidasDay"
//  RESOURCEUNAVAILABLE_MIDASFROZENAMOUNT = "ResourceUnavailable.MidasFrozenAmount"
//  RESOURCEUNAVAILABLE_MIDASMERCHANTBALANCE = "ResourceUnavailable.MidasMerchantBalance"
//  RESOURCEUNAVAILABLE_MIDASORDER = "ResourceUnavailable.MidasOrder"
//  RESOURCEUNAVAILABLE_MIDASUSERBALANCE = "ResourceUnavailable.MidasUserBalance"
//  RESOURCEUNAVAILABLE_MIDASWALLET = "ResourceUnavailable.MidasWallet"
//  UNAUTHORIZEDOPERATION_MIDAS = "UnauthorizedOperation.Midas"
func (c *Client) QueryCloudRefundOrderWithContext(ctx context.Context, request *QueryCloudRefundOrderRequest) (response *QueryCloudRefundOrderResponse, err error) {
    if request == nil {
        request = NewQueryCloudRefundOrderRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryCloudRefundOrder require credential")
    }

    request.SetContext(ctx)
    
    response = NewQueryCloudRefundOrderResponse()
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
    return c.QueryCommonTransferRechargeWithContext(context.Background(), request)
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
func (c *Client) QueryCommonTransferRechargeWithContext(ctx context.Context, request *QueryCommonTransferRechargeRequest) (response *QueryCommonTransferRechargeResponse, err error) {
    if request == nil {
        request = NewQueryCommonTransferRechargeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryCommonTransferRecharge require credential")
    }

    request.SetContext(ctx)
    
    response = NewQueryCommonTransferRechargeResponse()
    err = c.Send(request, response)
    return
}

func NewQueryCompanyTitleRequest() (request *QueryCompanyTitleRequest) {
    request = &QueryCompanyTitleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cpdp", APIVersion, "QueryCompanyTitle")
    
    
    return
}

func NewQueryCompanyTitleResponse() (response *QueryCompanyTitleResponse) {
    response = &QueryCompanyTitleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// QueryCompanyTitle
// 智慧零售-查询公司抬头
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_INVOICENOTFOUND = "ResourceNotFound.InvoiceNotFound"
func (c *Client) QueryCompanyTitle(request *QueryCompanyTitleRequest) (response *QueryCompanyTitleResponse, err error) {
    return c.QueryCompanyTitleWithContext(context.Background(), request)
}

// QueryCompanyTitle
// 智慧零售-查询公司抬头
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_INVOICENOTFOUND = "ResourceNotFound.InvoiceNotFound"
func (c *Client) QueryCompanyTitleWithContext(ctx context.Context, request *QueryCompanyTitleRequest) (response *QueryCompanyTitleResponse, err error) {
    if request == nil {
        request = NewQueryCompanyTitleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryCompanyTitle require credential")
    }

    request.SetContext(ctx)
    
    response = NewQueryCompanyTitleResponse()
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
//  AUTHFAILURE_SECRETKEYNOTFOUND = "AuthFailure.SecretKeyNotFound"
//  AUTHFAILURE_VERIFYERROR = "AuthFailure.VerifyError"
//  AUTHFAILURE_VERIFYTOKENFAILURE = "AuthFailure.VerifyTokenFailure"
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
    return c.QueryContractWithContext(context.Background(), request)
}

// QueryContract
// 通过此接口查询签约数据
//
// 可能返回的错误码:
//  AUTHFAILURE_SECRETKEYNOTFOUND = "AuthFailure.SecretKeyNotFound"
//  AUTHFAILURE_VERIFYERROR = "AuthFailure.VerifyError"
//  AUTHFAILURE_VERIFYTOKENFAILURE = "AuthFailure.VerifyTokenFailure"
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
func (c *Client) QueryContractWithContext(ctx context.Context, request *QueryContractRequest) (response *QueryContractResponse, err error) {
    if request == nil {
        request = NewQueryContractRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryContract require credential")
    }

    request.SetContext(ctx)
    
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
//  AUTHFAILURE_SECRETKEYNOTFOUND = "AuthFailure.SecretKeyNotFound"
//  AUTHFAILURE_VERIFYERROR = "AuthFailure.VerifyError"
//  AUTHFAILURE_VERIFYTOKENFAILURE = "AuthFailure.VerifyTokenFailure"
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
    return c.QueryContractPayFeeWithContext(context.Background(), request)
}

// QueryContractPayFee
// 云支付-查询支付方式费率及自定义表单项接口
//
// 可能返回的错误码:
//  AUTHFAILURE_SECRETKEYNOTFOUND = "AuthFailure.SecretKeyNotFound"
//  AUTHFAILURE_VERIFYERROR = "AuthFailure.VerifyError"
//  AUTHFAILURE_VERIFYTOKENFAILURE = "AuthFailure.VerifyTokenFailure"
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
func (c *Client) QueryContractPayFeeWithContext(ctx context.Context, request *QueryContractPayFeeRequest) (response *QueryContractPayFeeResponse, err error) {
    if request == nil {
        request = NewQueryContractPayFeeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryContractPayFee require credential")
    }

    request.SetContext(ctx)
    
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
//  AUTHFAILURE_SECRETKEYNOTFOUND = "AuthFailure.SecretKeyNotFound"
//  AUTHFAILURE_VERIFYERROR = "AuthFailure.VerifyError"
//  AUTHFAILURE_VERIFYTOKENFAILURE = "AuthFailure.VerifyTokenFailure"
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
    return c.QueryContractPayWayListWithContext(context.Background(), request)
}

// QueryContractPayWayList
// 云支付-查询合同支付方式列表接口
//
// 可能返回的错误码:
//  AUTHFAILURE_SECRETKEYNOTFOUND = "AuthFailure.SecretKeyNotFound"
//  AUTHFAILURE_VERIFYERROR = "AuthFailure.VerifyError"
//  AUTHFAILURE_VERIFYTOKENFAILURE = "AuthFailure.VerifyTokenFailure"
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
func (c *Client) QueryContractPayWayListWithContext(ctx context.Context, request *QueryContractPayWayListRequest) (response *QueryContractPayWayListResponse, err error) {
    if request == nil {
        request = NewQueryContractPayWayListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryContractPayWayList require credential")
    }

    request.SetContext(ctx)
    
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
//  AUTHFAILURE_SECRETKEYNOTFOUND = "AuthFailure.SecretKeyNotFound"
//  AUTHFAILURE_VERIFYERROR = "AuthFailure.VerifyError"
//  AUTHFAILURE_VERIFYTOKENFAILURE = "AuthFailure.VerifyTokenFailure"
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
    return c.QueryContractRelateShopWithContext(context.Background(), request)
}

// QueryContractRelateShop
// 云支付-查询合同可关联门店接口
//
// 可能返回的错误码:
//  AUTHFAILURE_SECRETKEYNOTFOUND = "AuthFailure.SecretKeyNotFound"
//  AUTHFAILURE_VERIFYERROR = "AuthFailure.VerifyError"
//  AUTHFAILURE_VERIFYTOKENFAILURE = "AuthFailure.VerifyTokenFailure"
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
func (c *Client) QueryContractRelateShopWithContext(ctx context.Context, request *QueryContractRelateShopRequest) (response *QueryContractRelateShopResponse, err error) {
    if request == nil {
        request = NewQueryContractRelateShopRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryContractRelateShop require credential")
    }

    request.SetContext(ctx)
    
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
    return c.QueryCustAcctIdBalanceWithContext(context.Background(), request)
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
func (c *Client) QueryCustAcctIdBalanceWithContext(ctx context.Context, request *QueryCustAcctIdBalanceRequest) (response *QueryCustAcctIdBalanceResponse, err error) {
    if request == nil {
        request = NewQueryCustAcctIdBalanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryCustAcctIdBalance require credential")
    }

    request.SetContext(ctx)
    
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
    return c.QueryDownloadBillURLWithContext(context.Background(), request)
}

// QueryDownloadBillURL
// 云鉴-查询对账单下载地址的接口
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADBILLERROR = "FailedOperation.DownloadBillError"
//  FAILEDOPERATION_INVALIDPARAMETER = "FailedOperation.InvalidParameter"
//  FAILEDOPERATION_MISSINGPARAMETER = "FailedOperation.MissingParameter"
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
func (c *Client) QueryDownloadBillURLWithContext(ctx context.Context, request *QueryDownloadBillURLRequest) (response *QueryDownloadBillURLResponse, err error) {
    if request == nil {
        request = NewQueryDownloadBillURLRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryDownloadBillURL require credential")
    }

    request.SetContext(ctx)
    
    response = NewQueryDownloadBillURLResponse()
    err = c.Send(request, response)
    return
}

func NewQueryExceedingInfoRequest() (request *QueryExceedingInfoRequest) {
    request = &QueryExceedingInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cpdp", APIVersion, "QueryExceedingInfo")
    
    
    return
}

func NewQueryExceedingInfoResponse() (response *QueryExceedingInfoResponse) {
    response = &QueryExceedingInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// QueryExceedingInfo
// 灵云-查询超额信息
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADBILLERROR = "FailedOperation.DownloadBillError"
//  FAILEDOPERATION_INVALIDPARAMETER = "FailedOperation.InvalidParameter"
//  FAILEDOPERATION_MISSINGPARAMETER = "FailedOperation.MissingParameter"
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
func (c *Client) QueryExceedingInfo(request *QueryExceedingInfoRequest) (response *QueryExceedingInfoResponse, err error) {
    return c.QueryExceedingInfoWithContext(context.Background(), request)
}

// QueryExceedingInfo
// 灵云-查询超额信息
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADBILLERROR = "FailedOperation.DownloadBillError"
//  FAILEDOPERATION_INVALIDPARAMETER = "FailedOperation.InvalidParameter"
//  FAILEDOPERATION_MISSINGPARAMETER = "FailedOperation.MissingParameter"
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
func (c *Client) QueryExceedingInfoWithContext(ctx context.Context, request *QueryExceedingInfoRequest) (response *QueryExceedingInfoResponse, err error) {
    if request == nil {
        request = NewQueryExceedingInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryExceedingInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewQueryExceedingInfoResponse()
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
    return c.QueryExchangeRateWithContext(context.Background(), request)
}

// QueryExchangeRate
// 跨境-查询汇率
//
// 可能返回的错误码:
//  INTERNALERROR_PARAMETERERROR = "InternalError.ParameterError"
//  INTERNALERROR_UNKOWNERROR = "InternalError.UnkownError"
//  RESOURCENOTFOUND_MERCHANTINFONOTFOUND = "ResourceNotFound.MerchantInfoNotFound"
func (c *Client) QueryExchangeRateWithContext(ctx context.Context, request *QueryExchangeRateRequest) (response *QueryExchangeRateResponse, err error) {
    if request == nil {
        request = NewQueryExchangeRateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryExchangeRate require credential")
    }

    request.SetContext(ctx)
    
    response = NewQueryExchangeRateResponse()
    err = c.Send(request, response)
    return
}

func NewQueryFinancialDataUrlRequest() (request *QueryFinancialDataUrlRequest) {
    request = &QueryFinancialDataUrlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cpdp", APIVersion, "QueryFinancialDataUrl")
    
    
    return
}

func NewQueryFinancialDataUrlResponse() (response *QueryFinancialDataUrlResponse) {
    response = &QueryFinancialDataUrlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// QueryFinancialDataUrl
// 财税-查询金融数据文件下载链接
//
// 可能返回的错误码:
//  FAILEDOPERATION_MOUNTNOTFOUND = "FailedOperation.MountNotFound"
func (c *Client) QueryFinancialDataUrl(request *QueryFinancialDataUrlRequest) (response *QueryFinancialDataUrlResponse, err error) {
    return c.QueryFinancialDataUrlWithContext(context.Background(), request)
}

// QueryFinancialDataUrl
// 财税-查询金融数据文件下载链接
//
// 可能返回的错误码:
//  FAILEDOPERATION_MOUNTNOTFOUND = "FailedOperation.MountNotFound"
func (c *Client) QueryFinancialDataUrlWithContext(ctx context.Context, request *QueryFinancialDataUrlRequest) (response *QueryFinancialDataUrlResponse, err error) {
    if request == nil {
        request = NewQueryFinancialDataUrlRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryFinancialDataUrl require credential")
    }

    request.SetContext(ctx)
    
    response = NewQueryFinancialDataUrlResponse()
    err = c.Send(request, response)
    return
}

func NewQueryFlexAmountBeforeTaxRequest() (request *QueryFlexAmountBeforeTaxRequest) {
    request = &QueryFlexAmountBeforeTaxRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cpdp", APIVersion, "QueryFlexAmountBeforeTax")
    
    
    return
}

func NewQueryFlexAmountBeforeTaxResponse() (response *QueryFlexAmountBeforeTaxResponse) {
    response = &QueryFlexAmountBeforeTaxResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// QueryFlexAmountBeforeTax
// 灵云V2-查询税前金额
//
// 可能返回的错误码:
//  FAILEDOPERATION_MOUNTNOTFOUND = "FailedOperation.MountNotFound"
func (c *Client) QueryFlexAmountBeforeTax(request *QueryFlexAmountBeforeTaxRequest) (response *QueryFlexAmountBeforeTaxResponse, err error) {
    return c.QueryFlexAmountBeforeTaxWithContext(context.Background(), request)
}

// QueryFlexAmountBeforeTax
// 灵云V2-查询税前金额
//
// 可能返回的错误码:
//  FAILEDOPERATION_MOUNTNOTFOUND = "FailedOperation.MountNotFound"
func (c *Client) QueryFlexAmountBeforeTaxWithContext(ctx context.Context, request *QueryFlexAmountBeforeTaxRequest) (response *QueryFlexAmountBeforeTaxResponse, err error) {
    if request == nil {
        request = NewQueryFlexAmountBeforeTaxRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryFlexAmountBeforeTax require credential")
    }

    request.SetContext(ctx)
    
    response = NewQueryFlexAmountBeforeTaxResponse()
    err = c.Send(request, response)
    return
}

func NewQueryFlexBillDownloadUrlRequest() (request *QueryFlexBillDownloadUrlRequest) {
    request = &QueryFlexBillDownloadUrlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cpdp", APIVersion, "QueryFlexBillDownloadUrl")
    
    
    return
}

func NewQueryFlexBillDownloadUrlResponse() (response *QueryFlexBillDownloadUrlResponse) {
    response = &QueryFlexBillDownloadUrlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// QueryFlexBillDownloadUrl
// 灵云V2-查询对账单文件下载链接
//
// 可能返回的错误码:
//  FAILEDOPERATION_MOUNTNOTFOUND = "FailedOperation.MountNotFound"
func (c *Client) QueryFlexBillDownloadUrl(request *QueryFlexBillDownloadUrlRequest) (response *QueryFlexBillDownloadUrlResponse, err error) {
    return c.QueryFlexBillDownloadUrlWithContext(context.Background(), request)
}

// QueryFlexBillDownloadUrl
// 灵云V2-查询对账单文件下载链接
//
// 可能返回的错误码:
//  FAILEDOPERATION_MOUNTNOTFOUND = "FailedOperation.MountNotFound"
func (c *Client) QueryFlexBillDownloadUrlWithContext(ctx context.Context, request *QueryFlexBillDownloadUrlRequest) (response *QueryFlexBillDownloadUrlResponse, err error) {
    if request == nil {
        request = NewQueryFlexBillDownloadUrlRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryFlexBillDownloadUrl require credential")
    }

    request.SetContext(ctx)
    
    response = NewQueryFlexBillDownloadUrlResponse()
    err = c.Send(request, response)
    return
}

func NewQueryFlexFreezeOrderListRequest() (request *QueryFlexFreezeOrderListRequest) {
    request = &QueryFlexFreezeOrderListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cpdp", APIVersion, "QueryFlexFreezeOrderList")
    
    
    return
}

func NewQueryFlexFreezeOrderListResponse() (response *QueryFlexFreezeOrderListResponse) {
    response = &QueryFlexFreezeOrderListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// QueryFlexFreezeOrderList
// 灵云V2-查询冻结订单列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_MOUNTNOTFOUND = "FailedOperation.MountNotFound"
func (c *Client) QueryFlexFreezeOrderList(request *QueryFlexFreezeOrderListRequest) (response *QueryFlexFreezeOrderListResponse, err error) {
    return c.QueryFlexFreezeOrderListWithContext(context.Background(), request)
}

// QueryFlexFreezeOrderList
// 灵云V2-查询冻结订单列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_MOUNTNOTFOUND = "FailedOperation.MountNotFound"
func (c *Client) QueryFlexFreezeOrderListWithContext(ctx context.Context, request *QueryFlexFreezeOrderListRequest) (response *QueryFlexFreezeOrderListResponse, err error) {
    if request == nil {
        request = NewQueryFlexFreezeOrderListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryFlexFreezeOrderList require credential")
    }

    request.SetContext(ctx)
    
    response = NewQueryFlexFreezeOrderListResponse()
    err = c.Send(request, response)
    return
}

func NewQueryFlexOrderSummaryListRequest() (request *QueryFlexOrderSummaryListRequest) {
    request = &QueryFlexOrderSummaryListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cpdp", APIVersion, "QueryFlexOrderSummaryList")
    
    
    return
}

func NewQueryFlexOrderSummaryListResponse() (response *QueryFlexOrderSummaryListResponse) {
    response = &QueryFlexOrderSummaryListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// QueryFlexOrderSummaryList
// 灵云V2-订单汇总列表查询
//
// 可能返回的错误码:
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
func (c *Client) QueryFlexOrderSummaryList(request *QueryFlexOrderSummaryListRequest) (response *QueryFlexOrderSummaryListResponse, err error) {
    return c.QueryFlexOrderSummaryListWithContext(context.Background(), request)
}

// QueryFlexOrderSummaryList
// 灵云V2-订单汇总列表查询
//
// 可能返回的错误码:
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
func (c *Client) QueryFlexOrderSummaryListWithContext(ctx context.Context, request *QueryFlexOrderSummaryListRequest) (response *QueryFlexOrderSummaryListResponse, err error) {
    if request == nil {
        request = NewQueryFlexOrderSummaryListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryFlexOrderSummaryList require credential")
    }

    request.SetContext(ctx)
    
    response = NewQueryFlexOrderSummaryListResponse()
    err = c.Send(request, response)
    return
}

func NewQueryFlexPayeeAccountBalanceRequest() (request *QueryFlexPayeeAccountBalanceRequest) {
    request = &QueryFlexPayeeAccountBalanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cpdp", APIVersion, "QueryFlexPayeeAccountBalance")
    
    
    return
}

func NewQueryFlexPayeeAccountBalanceResponse() (response *QueryFlexPayeeAccountBalanceResponse) {
    response = &QueryFlexPayeeAccountBalanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// QueryFlexPayeeAccountBalance
// 灵云V2-收款用户账户余额查询
//
// 可能返回的错误码:
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
func (c *Client) QueryFlexPayeeAccountBalance(request *QueryFlexPayeeAccountBalanceRequest) (response *QueryFlexPayeeAccountBalanceResponse, err error) {
    return c.QueryFlexPayeeAccountBalanceWithContext(context.Background(), request)
}

// QueryFlexPayeeAccountBalance
// 灵云V2-收款用户账户余额查询
//
// 可能返回的错误码:
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
func (c *Client) QueryFlexPayeeAccountBalanceWithContext(ctx context.Context, request *QueryFlexPayeeAccountBalanceRequest) (response *QueryFlexPayeeAccountBalanceResponse, err error) {
    if request == nil {
        request = NewQueryFlexPayeeAccountBalanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryFlexPayeeAccountBalance require credential")
    }

    request.SetContext(ctx)
    
    response = NewQueryFlexPayeeAccountBalanceResponse()
    err = c.Send(request, response)
    return
}

func NewQueryFlexPayeeAccountInfoRequest() (request *QueryFlexPayeeAccountInfoRequest) {
    request = &QueryFlexPayeeAccountInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cpdp", APIVersion, "QueryFlexPayeeAccountInfo")
    
    
    return
}

func NewQueryFlexPayeeAccountInfoResponse() (response *QueryFlexPayeeAccountInfoResponse) {
    response = &QueryFlexPayeeAccountInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// QueryFlexPayeeAccountInfo
// 灵云V2-收款用户账户信息查询
//
// 可能返回的错误码:
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
func (c *Client) QueryFlexPayeeAccountInfo(request *QueryFlexPayeeAccountInfoRequest) (response *QueryFlexPayeeAccountInfoResponse, err error) {
    return c.QueryFlexPayeeAccountInfoWithContext(context.Background(), request)
}

// QueryFlexPayeeAccountInfo
// 灵云V2-收款用户账户信息查询
//
// 可能返回的错误码:
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
func (c *Client) QueryFlexPayeeAccountInfoWithContext(ctx context.Context, request *QueryFlexPayeeAccountInfoRequest) (response *QueryFlexPayeeAccountInfoResponse, err error) {
    if request == nil {
        request = NewQueryFlexPayeeAccountInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryFlexPayeeAccountInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewQueryFlexPayeeAccountInfoResponse()
    err = c.Send(request, response)
    return
}

func NewQueryFlexPayeeAccountListRequest() (request *QueryFlexPayeeAccountListRequest) {
    request = &QueryFlexPayeeAccountListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cpdp", APIVersion, "QueryFlexPayeeAccountList")
    
    
    return
}

func NewQueryFlexPayeeAccountListResponse() (response *QueryFlexPayeeAccountListResponse) {
    response = &QueryFlexPayeeAccountListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// QueryFlexPayeeAccountList
// 灵云V2-收款用户账户列表查询
//
// 可能返回的错误码:
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
func (c *Client) QueryFlexPayeeAccountList(request *QueryFlexPayeeAccountListRequest) (response *QueryFlexPayeeAccountListResponse, err error) {
    return c.QueryFlexPayeeAccountListWithContext(context.Background(), request)
}

// QueryFlexPayeeAccountList
// 灵云V2-收款用户账户列表查询
//
// 可能返回的错误码:
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
func (c *Client) QueryFlexPayeeAccountListWithContext(ctx context.Context, request *QueryFlexPayeeAccountListRequest) (response *QueryFlexPayeeAccountListResponse, err error) {
    if request == nil {
        request = NewQueryFlexPayeeAccountListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryFlexPayeeAccountList require credential")
    }

    request.SetContext(ctx)
    
    response = NewQueryFlexPayeeAccountListResponse()
    err = c.Send(request, response)
    return
}

func NewQueryFlexPayeeInfoRequest() (request *QueryFlexPayeeInfoRequest) {
    request = &QueryFlexPayeeInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cpdp", APIVersion, "QueryFlexPayeeInfo")
    
    
    return
}

func NewQueryFlexPayeeInfoResponse() (response *QueryFlexPayeeInfoResponse) {
    response = &QueryFlexPayeeInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// QueryFlexPayeeInfo
// 灵云V2-收款用户信息查询
//
// 可能返回的错误码:
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
func (c *Client) QueryFlexPayeeInfo(request *QueryFlexPayeeInfoRequest) (response *QueryFlexPayeeInfoResponse, err error) {
    return c.QueryFlexPayeeInfoWithContext(context.Background(), request)
}

// QueryFlexPayeeInfo
// 灵云V2-收款用户信息查询
//
// 可能返回的错误码:
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
func (c *Client) QueryFlexPayeeInfoWithContext(ctx context.Context, request *QueryFlexPayeeInfoRequest) (response *QueryFlexPayeeInfoResponse, err error) {
    if request == nil {
        request = NewQueryFlexPayeeInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryFlexPayeeInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewQueryFlexPayeeInfoResponse()
    err = c.Send(request, response)
    return
}

func NewQueryFlexPaymentOrderListRequest() (request *QueryFlexPaymentOrderListRequest) {
    request = &QueryFlexPaymentOrderListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cpdp", APIVersion, "QueryFlexPaymentOrderList")
    
    
    return
}

func NewQueryFlexPaymentOrderListResponse() (response *QueryFlexPaymentOrderListResponse) {
    response = &QueryFlexPaymentOrderListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// QueryFlexPaymentOrderList
// 灵云V2-查询付款订单列表
//
// 可能返回的错误码:
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
func (c *Client) QueryFlexPaymentOrderList(request *QueryFlexPaymentOrderListRequest) (response *QueryFlexPaymentOrderListResponse, err error) {
    return c.QueryFlexPaymentOrderListWithContext(context.Background(), request)
}

// QueryFlexPaymentOrderList
// 灵云V2-查询付款订单列表
//
// 可能返回的错误码:
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
func (c *Client) QueryFlexPaymentOrderListWithContext(ctx context.Context, request *QueryFlexPaymentOrderListRequest) (response *QueryFlexPaymentOrderListResponse, err error) {
    if request == nil {
        request = NewQueryFlexPaymentOrderListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryFlexPaymentOrderList require credential")
    }

    request.SetContext(ctx)
    
    response = NewQueryFlexPaymentOrderListResponse()
    err = c.Send(request, response)
    return
}

func NewQueryFlexPaymentOrderStatusRequest() (request *QueryFlexPaymentOrderStatusRequest) {
    request = &QueryFlexPaymentOrderStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cpdp", APIVersion, "QueryFlexPaymentOrderStatus")
    
    
    return
}

func NewQueryFlexPaymentOrderStatusResponse() (response *QueryFlexPaymentOrderStatusResponse) {
    response = &QueryFlexPaymentOrderStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// QueryFlexPaymentOrderStatus
// 灵云V2-查询付款订单状态
//
// 可能返回的错误码:
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
func (c *Client) QueryFlexPaymentOrderStatus(request *QueryFlexPaymentOrderStatusRequest) (response *QueryFlexPaymentOrderStatusResponse, err error) {
    return c.QueryFlexPaymentOrderStatusWithContext(context.Background(), request)
}

// QueryFlexPaymentOrderStatus
// 灵云V2-查询付款订单状态
//
// 可能返回的错误码:
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
func (c *Client) QueryFlexPaymentOrderStatusWithContext(ctx context.Context, request *QueryFlexPaymentOrderStatusRequest) (response *QueryFlexPaymentOrderStatusResponse, err error) {
    if request == nil {
        request = NewQueryFlexPaymentOrderStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryFlexPaymentOrderStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewQueryFlexPaymentOrderStatusResponse()
    err = c.Send(request, response)
    return
}

func NewQueryFlexPlatformAccountBalanceRequest() (request *QueryFlexPlatformAccountBalanceRequest) {
    request = &QueryFlexPlatformAccountBalanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cpdp", APIVersion, "QueryFlexPlatformAccountBalance")
    
    
    return
}

func NewQueryFlexPlatformAccountBalanceResponse() (response *QueryFlexPlatformAccountBalanceResponse) {
    response = &QueryFlexPlatformAccountBalanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// QueryFlexPlatformAccountBalance
// 灵云V2-平台账户余额查询
//
// 可能返回的错误码:
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
func (c *Client) QueryFlexPlatformAccountBalance(request *QueryFlexPlatformAccountBalanceRequest) (response *QueryFlexPlatformAccountBalanceResponse, err error) {
    return c.QueryFlexPlatformAccountBalanceWithContext(context.Background(), request)
}

// QueryFlexPlatformAccountBalance
// 灵云V2-平台账户余额查询
//
// 可能返回的错误码:
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
func (c *Client) QueryFlexPlatformAccountBalanceWithContext(ctx context.Context, request *QueryFlexPlatformAccountBalanceRequest) (response *QueryFlexPlatformAccountBalanceResponse, err error) {
    if request == nil {
        request = NewQueryFlexPlatformAccountBalanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryFlexPlatformAccountBalance require credential")
    }

    request.SetContext(ctx)
    
    response = NewQueryFlexPlatformAccountBalanceResponse()
    err = c.Send(request, response)
    return
}

func NewQueryFlexServiceProviderAccountBalanceRequest() (request *QueryFlexServiceProviderAccountBalanceRequest) {
    request = &QueryFlexServiceProviderAccountBalanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cpdp", APIVersion, "QueryFlexServiceProviderAccountBalance")
    
    
    return
}

func NewQueryFlexServiceProviderAccountBalanceResponse() (response *QueryFlexServiceProviderAccountBalanceResponse) {
    response = &QueryFlexServiceProviderAccountBalanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// QueryFlexServiceProviderAccountBalance
// 灵云V2-查询服务商账户余额
//
// 可能返回的错误码:
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
func (c *Client) QueryFlexServiceProviderAccountBalance(request *QueryFlexServiceProviderAccountBalanceRequest) (response *QueryFlexServiceProviderAccountBalanceResponse, err error) {
    return c.QueryFlexServiceProviderAccountBalanceWithContext(context.Background(), request)
}

// QueryFlexServiceProviderAccountBalance
// 灵云V2-查询服务商账户余额
//
// 可能返回的错误码:
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
func (c *Client) QueryFlexServiceProviderAccountBalanceWithContext(ctx context.Context, request *QueryFlexServiceProviderAccountBalanceRequest) (response *QueryFlexServiceProviderAccountBalanceResponse, err error) {
    if request == nil {
        request = NewQueryFlexServiceProviderAccountBalanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryFlexServiceProviderAccountBalance require credential")
    }

    request.SetContext(ctx)
    
    response = NewQueryFlexServiceProviderAccountBalanceResponse()
    err = c.Send(request, response)
    return
}

func NewQueryFlexSettlementOrderListRequest() (request *QueryFlexSettlementOrderListRequest) {
    request = &QueryFlexSettlementOrderListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cpdp", APIVersion, "QueryFlexSettlementOrderList")
    
    
    return
}

func NewQueryFlexSettlementOrderListResponse() (response *QueryFlexSettlementOrderListResponse) {
    response = &QueryFlexSettlementOrderListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// QueryFlexSettlementOrderList
// 灵云V2-查询结算订单列表
//
// 可能返回的错误码:
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
func (c *Client) QueryFlexSettlementOrderList(request *QueryFlexSettlementOrderListRequest) (response *QueryFlexSettlementOrderListResponse, err error) {
    return c.QueryFlexSettlementOrderListWithContext(context.Background(), request)
}

// QueryFlexSettlementOrderList
// 灵云V2-查询结算订单列表
//
// 可能返回的错误码:
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
func (c *Client) QueryFlexSettlementOrderListWithContext(ctx context.Context, request *QueryFlexSettlementOrderListRequest) (response *QueryFlexSettlementOrderListResponse, err error) {
    if request == nil {
        request = NewQueryFlexSettlementOrderListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryFlexSettlementOrderList require credential")
    }

    request.SetContext(ctx)
    
    response = NewQueryFlexSettlementOrderListResponse()
    err = c.Send(request, response)
    return
}

func NewQueryFundsTransactionDetailsRequest() (request *QueryFundsTransactionDetailsRequest) {
    request = &QueryFundsTransactionDetailsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cpdp", APIVersion, "QueryFundsTransactionDetails")
    
    
    return
}

func NewQueryFundsTransactionDetailsResponse() (response *QueryFundsTransactionDetailsResponse) {
    response = &QueryFundsTransactionDetailsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// QueryFundsTransactionDetails
// 聚鑫-查询会员资金交易信息列表
//
// 可能返回的错误码:
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
func (c *Client) QueryFundsTransactionDetails(request *QueryFundsTransactionDetailsRequest) (response *QueryFundsTransactionDetailsResponse, err error) {
    return c.QueryFundsTransactionDetailsWithContext(context.Background(), request)
}

// QueryFundsTransactionDetails
// 聚鑫-查询会员资金交易信息列表
//
// 可能返回的错误码:
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
func (c *Client) QueryFundsTransactionDetailsWithContext(ctx context.Context, request *QueryFundsTransactionDetailsRequest) (response *QueryFundsTransactionDetailsResponse, err error) {
    if request == nil {
        request = NewQueryFundsTransactionDetailsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryFundsTransactionDetails require credential")
    }

    request.SetContext(ctx)
    
    response = NewQueryFundsTransactionDetailsResponse()
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
    return c.QueryInvoiceWithContext(context.Background(), request)
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
func (c *Client) QueryInvoiceWithContext(ctx context.Context, request *QueryInvoiceRequest) (response *QueryInvoiceResponse, err error) {
    if request == nil {
        request = NewQueryInvoiceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryInvoice require credential")
    }

    request.SetContext(ctx)
    
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
    return c.QueryInvoiceV2WithContext(context.Background(), request)
}

// QueryInvoiceV2
// 智慧零售-发票查询V2
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_INVOICENOTFOUND = "ResourceNotFound.InvoiceNotFound"
func (c *Client) QueryInvoiceV2WithContext(ctx context.Context, request *QueryInvoiceV2Request) (response *QueryInvoiceV2Response, err error) {
    if request == nil {
        request = NewQueryInvoiceV2Request()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryInvoiceV2 require credential")
    }

    request.SetContext(ctx)
    
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
    return c.QueryMaliciousRegistrationWithContext(context.Background(), request)
}

// QueryMaliciousRegistration
// 商户恶意注册接口
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_INVOICENOTFOUND = "ResourceNotFound.InvoiceNotFound"
func (c *Client) QueryMaliciousRegistrationWithContext(ctx context.Context, request *QueryMaliciousRegistrationRequest) (response *QueryMaliciousRegistrationResponse, err error) {
    if request == nil {
        request = NewQueryMaliciousRegistrationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryMaliciousRegistration require credential")
    }

    request.SetContext(ctx)
    
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
    return c.QueryMemberBindWithContext(context.Background(), request)
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
func (c *Client) QueryMemberBindWithContext(ctx context.Context, request *QueryMemberBindRequest) (response *QueryMemberBindResponse, err error) {
    if request == nil {
        request = NewQueryMemberBindRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryMemberBind require credential")
    }

    request.SetContext(ctx)
    
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
    return c.QueryMemberTransactionWithContext(context.Background(), request)
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
func (c *Client) QueryMemberTransactionWithContext(ctx context.Context, request *QueryMemberTransactionRequest) (response *QueryMemberTransactionResponse, err error) {
    if request == nil {
        request = NewQueryMemberTransactionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryMemberTransaction require credential")
    }

    request.SetContext(ctx)
    
    response = NewQueryMemberTransactionResponse()
    err = c.Send(request, response)
    return
}

func NewQueryMemberTransactionDetailsRequest() (request *QueryMemberTransactionDetailsRequest) {
    request = &QueryMemberTransactionDetailsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cpdp", APIVersion, "QueryMemberTransactionDetails")
    
    
    return
}

func NewQueryMemberTransactionDetailsResponse() (response *QueryMemberTransactionDetailsResponse) {
    response = &QueryMemberTransactionDetailsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// QueryMemberTransactionDetails
// 聚鑫-查询会员间交易信息列表
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
func (c *Client) QueryMemberTransactionDetails(request *QueryMemberTransactionDetailsRequest) (response *QueryMemberTransactionDetailsResponse, err error) {
    return c.QueryMemberTransactionDetailsWithContext(context.Background(), request)
}

// QueryMemberTransactionDetails
// 聚鑫-查询会员间交易信息列表
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
func (c *Client) QueryMemberTransactionDetailsWithContext(ctx context.Context, request *QueryMemberTransactionDetailsRequest) (response *QueryMemberTransactionDetailsResponse, err error) {
    if request == nil {
        request = NewQueryMemberTransactionDetailsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryMemberTransactionDetails require credential")
    }

    request.SetContext(ctx)
    
    response = NewQueryMemberTransactionDetailsResponse()
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
    return c.QueryMerchantWithContext(context.Background(), request)
}

// QueryMerchant
// 云鉴-商户信息查询接口
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDPARAMETER = "FailedOperation.InvalidParameter"
//  FAILEDOPERATION_MERCHANTNOTEXIST = "FailedOperation.MerchantNotExist"
//  FAILEDOPERATION_MISSINGPARAMETER = "FailedOperation.MissingParameter"
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
func (c *Client) QueryMerchantWithContext(ctx context.Context, request *QueryMerchantRequest) (response *QueryMerchantResponse, err error) {
    if request == nil {
        request = NewQueryMerchantRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryMerchant require credential")
    }

    request.SetContext(ctx)
    
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
    return c.QueryMerchantBalanceWithContext(context.Background(), request)
}

// QueryMerchantBalance
// 跨境-对接方账户余额查询
//
// 可能返回的错误码:
//  INTERNALERROR_PARAMETERERROR = "InternalError.ParameterError"
//  INTERNALERROR_UNKOWNERROR = "InternalError.UnkownError"
//  RESOURCENOTFOUND_MERCHANTINFONOTFOUND = "ResourceNotFound.MerchantInfoNotFound"
func (c *Client) QueryMerchantBalanceWithContext(ctx context.Context, request *QueryMerchantBalanceRequest) (response *QueryMerchantBalanceResponse, err error) {
    if request == nil {
        request = NewQueryMerchantBalanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryMerchantBalance require credential")
    }

    request.SetContext(ctx)
    
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
    return c.QueryMerchantClassificationWithContext(context.Background(), request)
}

// QueryMerchantClassification
// 云支付-查询商户分类接口
//
// 可能返回的错误码:
//  INTERNALERROR_PARAMETERERROR = "InternalError.ParameterError"
//  INTERNALERROR_UNKOWNERROR = "InternalError.UnkownError"
//  RESOURCENOTFOUND_MERCHANTINFONOTFOUND = "ResourceNotFound.MerchantInfoNotFound"
func (c *Client) QueryMerchantClassificationWithContext(ctx context.Context, request *QueryMerchantClassificationRequest) (response *QueryMerchantClassificationResponse, err error) {
    if request == nil {
        request = NewQueryMerchantClassificationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryMerchantClassification require credential")
    }

    request.SetContext(ctx)
    
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
    return c.QueryMerchantInfoForManagementWithContext(context.Background(), request)
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
func (c *Client) QueryMerchantInfoForManagementWithContext(ctx context.Context, request *QueryMerchantInfoForManagementRequest) (response *QueryMerchantInfoForManagementResponse, err error) {
    if request == nil {
        request = NewQueryMerchantInfoForManagementRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryMerchantInfoForManagement require credential")
    }

    request.SetContext(ctx)
    
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
    return c.QueryMerchantOrderWithContext(context.Background(), request)
}

// QueryMerchantOrder
// 云鉴-消费订单查询接口
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDPARAMETER = "FailedOperation.InvalidParameter"
//  FAILEDOPERATION_MISSINGPARAMETER = "FailedOperation.MissingParameter"
//  FAILEDOPERATION_QUERYORDERERROR = "FailedOperation.QueryOrderError"
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
func (c *Client) QueryMerchantOrderWithContext(ctx context.Context, request *QueryMerchantOrderRequest) (response *QueryMerchantOrderResponse, err error) {
    if request == nil {
        request = NewQueryMerchantOrderRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryMerchantOrder require credential")
    }

    request.SetContext(ctx)
    
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
    return c.QueryMerchantPayWayListWithContext(context.Background(), request)
}

// QueryMerchantPayWayList
// 商户查询已开通的支付方式列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDPARAMETER = "FailedOperation.InvalidParameter"
//  FAILEDOPERATION_MISSINGPARAMETER = "FailedOperation.MissingParameter"
//  FAILEDOPERATION_QUERYORDERERROR = "FailedOperation.QueryOrderError"
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
func (c *Client) QueryMerchantPayWayListWithContext(ctx context.Context, request *QueryMerchantPayWayListRequest) (response *QueryMerchantPayWayListResponse, err error) {
    if request == nil {
        request = NewQueryMerchantPayWayListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryMerchantPayWayList require credential")
    }

    request.SetContext(ctx)
    
    response = NewQueryMerchantPayWayListResponse()
    err = c.Send(request, response)
    return
}

func NewQueryOpenBankBankAccountBalanceRequest() (request *QueryOpenBankBankAccountBalanceRequest) {
    request = &QueryOpenBankBankAccountBalanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cpdp", APIVersion, "QueryOpenBankBankAccountBalance")
    
    
    return
}

func NewQueryOpenBankBankAccountBalanceResponse() (response *QueryOpenBankBankAccountBalanceResponse) {
    response = &QueryOpenBankBankAccountBalanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// QueryOpenBankBankAccountBalance
// 云企付-子商户银行卡余额查询
//
// 可能返回的错误码:
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
func (c *Client) QueryOpenBankBankAccountBalance(request *QueryOpenBankBankAccountBalanceRequest) (response *QueryOpenBankBankAccountBalanceResponse, err error) {
    return c.QueryOpenBankBankAccountBalanceWithContext(context.Background(), request)
}

// QueryOpenBankBankAccountBalance
// 云企付-子商户银行卡余额查询
//
// 可能返回的错误码:
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
func (c *Client) QueryOpenBankBankAccountBalanceWithContext(ctx context.Context, request *QueryOpenBankBankAccountBalanceRequest) (response *QueryOpenBankBankAccountBalanceResponse, err error) {
    if request == nil {
        request = NewQueryOpenBankBankAccountBalanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryOpenBankBankAccountBalance require credential")
    }

    request.SetContext(ctx)
    
    response = NewQueryOpenBankBankAccountBalanceResponse()
    err = c.Send(request, response)
    return
}

func NewQueryOpenBankBankBranchListRequest() (request *QueryOpenBankBankBranchListRequest) {
    request = &QueryOpenBankBankBranchListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cpdp", APIVersion, "QueryOpenBankBankBranchList")
    
    
    return
}

func NewQueryOpenBankBankBranchListResponse() (response *QueryOpenBankBankBranchListResponse) {
    response = &QueryOpenBankBankBranchListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// QueryOpenBankBankBranchList
// 云企付-查询联行号
//
// 可能返回的错误码:
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
func (c *Client) QueryOpenBankBankBranchList(request *QueryOpenBankBankBranchListRequest) (response *QueryOpenBankBankBranchListResponse, err error) {
    return c.QueryOpenBankBankBranchListWithContext(context.Background(), request)
}

// QueryOpenBankBankBranchList
// 云企付-查询联行号
//
// 可能返回的错误码:
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
func (c *Client) QueryOpenBankBankBranchListWithContext(ctx context.Context, request *QueryOpenBankBankBranchListRequest) (response *QueryOpenBankBankBranchListResponse, err error) {
    if request == nil {
        request = NewQueryOpenBankBankBranchListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryOpenBankBankBranchList require credential")
    }

    request.SetContext(ctx)
    
    response = NewQueryOpenBankBankBranchListResponse()
    err = c.Send(request, response)
    return
}

func NewQueryOpenBankBillDataPageRequest() (request *QueryOpenBankBillDataPageRequest) {
    request = &QueryOpenBankBillDataPageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cpdp", APIVersion, "QueryOpenBankBillDataPage")
    
    
    return
}

func NewQueryOpenBankBillDataPageResponse() (response *QueryOpenBankBillDataPageResponse) {
    response = &QueryOpenBankBillDataPageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// QueryOpenBankBillDataPage
// 云企付-分页查询对账单数据
//
// 可能返回的错误码:
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
func (c *Client) QueryOpenBankBillDataPage(request *QueryOpenBankBillDataPageRequest) (response *QueryOpenBankBillDataPageResponse, err error) {
    return c.QueryOpenBankBillDataPageWithContext(context.Background(), request)
}

// QueryOpenBankBillDataPage
// 云企付-分页查询对账单数据
//
// 可能返回的错误码:
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
func (c *Client) QueryOpenBankBillDataPageWithContext(ctx context.Context, request *QueryOpenBankBillDataPageRequest) (response *QueryOpenBankBillDataPageResponse, err error) {
    if request == nil {
        request = NewQueryOpenBankBillDataPageRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryOpenBankBillDataPage require credential")
    }

    request.SetContext(ctx)
    
    response = NewQueryOpenBankBillDataPageResponse()
    err = c.Send(request, response)
    return
}

func NewQueryOpenBankBindExternalSubMerchantBankAccountRequest() (request *QueryOpenBankBindExternalSubMerchantBankAccountRequest) {
    request = &QueryOpenBankBindExternalSubMerchantBankAccountRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cpdp", APIVersion, "QueryOpenBankBindExternalSubMerchantBankAccount")
    
    
    return
}

func NewQueryOpenBankBindExternalSubMerchantBankAccountResponse() (response *QueryOpenBankBindExternalSubMerchantBankAccountResponse) {
    response = &QueryOpenBankBindExternalSubMerchantBankAccountResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// QueryOpenBankBindExternalSubMerchantBankAccount
// 云企付-子商户银行卡绑定结果查询
//
// 可能返回的错误码:
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
func (c *Client) QueryOpenBankBindExternalSubMerchantBankAccount(request *QueryOpenBankBindExternalSubMerchantBankAccountRequest) (response *QueryOpenBankBindExternalSubMerchantBankAccountResponse, err error) {
    return c.QueryOpenBankBindExternalSubMerchantBankAccountWithContext(context.Background(), request)
}

// QueryOpenBankBindExternalSubMerchantBankAccount
// 云企付-子商户银行卡绑定结果查询
//
// 可能返回的错误码:
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
func (c *Client) QueryOpenBankBindExternalSubMerchantBankAccountWithContext(ctx context.Context, request *QueryOpenBankBindExternalSubMerchantBankAccountRequest) (response *QueryOpenBankBindExternalSubMerchantBankAccountResponse, err error) {
    if request == nil {
        request = NewQueryOpenBankBindExternalSubMerchantBankAccountRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryOpenBankBindExternalSubMerchantBankAccount require credential")
    }

    request.SetContext(ctx)
    
    response = NewQueryOpenBankBindExternalSubMerchantBankAccountResponse()
    err = c.Send(request, response)
    return
}

func NewQueryOpenBankDailyReceiptDownloadUrlRequest() (request *QueryOpenBankDailyReceiptDownloadUrlRequest) {
    request = &QueryOpenBankDailyReceiptDownloadUrlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cpdp", APIVersion, "QueryOpenBankDailyReceiptDownloadUrl")
    
    
    return
}

func NewQueryOpenBankDailyReceiptDownloadUrlResponse() (response *QueryOpenBankDailyReceiptDownloadUrlResponse) {
    response = &QueryOpenBankDailyReceiptDownloadUrlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// QueryOpenBankDailyReceiptDownloadUrl
// 云企付-按日期批量查询回单下载地址
//
// 可能返回的错误码:
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
func (c *Client) QueryOpenBankDailyReceiptDownloadUrl(request *QueryOpenBankDailyReceiptDownloadUrlRequest) (response *QueryOpenBankDailyReceiptDownloadUrlResponse, err error) {
    return c.QueryOpenBankDailyReceiptDownloadUrlWithContext(context.Background(), request)
}

// QueryOpenBankDailyReceiptDownloadUrl
// 云企付-按日期批量查询回单下载地址
//
// 可能返回的错误码:
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
func (c *Client) QueryOpenBankDailyReceiptDownloadUrlWithContext(ctx context.Context, request *QueryOpenBankDailyReceiptDownloadUrlRequest) (response *QueryOpenBankDailyReceiptDownloadUrlResponse, err error) {
    if request == nil {
        request = NewQueryOpenBankDailyReceiptDownloadUrlRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryOpenBankDailyReceiptDownloadUrl require credential")
    }

    request.SetContext(ctx)
    
    response = NewQueryOpenBankDailyReceiptDownloadUrlResponse()
    err = c.Send(request, response)
    return
}

func NewQueryOpenBankDownLoadUrlRequest() (request *QueryOpenBankDownLoadUrlRequest) {
    request = &QueryOpenBankDownLoadUrlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cpdp", APIVersion, "QueryOpenBankDownLoadUrl")
    
    
    return
}

func NewQueryOpenBankDownLoadUrlResponse() (response *QueryOpenBankDownLoadUrlResponse) {
    response = &QueryOpenBankDownLoadUrlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// QueryOpenBankDownLoadUrl
// 云企付-查询对账单下载地址
//
// 可能返回的错误码:
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
func (c *Client) QueryOpenBankDownLoadUrl(request *QueryOpenBankDownLoadUrlRequest) (response *QueryOpenBankDownLoadUrlResponse, err error) {
    return c.QueryOpenBankDownLoadUrlWithContext(context.Background(), request)
}

// QueryOpenBankDownLoadUrl
// 云企付-查询对账单下载地址
//
// 可能返回的错误码:
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
func (c *Client) QueryOpenBankDownLoadUrlWithContext(ctx context.Context, request *QueryOpenBankDownLoadUrlRequest) (response *QueryOpenBankDownLoadUrlResponse, err error) {
    if request == nil {
        request = NewQueryOpenBankDownLoadUrlRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryOpenBankDownLoadUrl require credential")
    }

    request.SetContext(ctx)
    
    response = NewQueryOpenBankDownLoadUrlResponse()
    err = c.Send(request, response)
    return
}

func NewQueryOpenBankExternalSubAccountBookBalanceRequest() (request *QueryOpenBankExternalSubAccountBookBalanceRequest) {
    request = &QueryOpenBankExternalSubAccountBookBalanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cpdp", APIVersion, "QueryOpenBankExternalSubAccountBookBalance")
    
    
    return
}

func NewQueryOpenBankExternalSubAccountBookBalanceResponse() (response *QueryOpenBankExternalSubAccountBookBalanceResponse) {
    response = &QueryOpenBankExternalSubAccountBookBalanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// QueryOpenBankExternalSubAccountBookBalance
// 第三方子商户电子记账本余额查询接口
//
// 可能返回的错误码:
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
func (c *Client) QueryOpenBankExternalSubAccountBookBalance(request *QueryOpenBankExternalSubAccountBookBalanceRequest) (response *QueryOpenBankExternalSubAccountBookBalanceResponse, err error) {
    return c.QueryOpenBankExternalSubAccountBookBalanceWithContext(context.Background(), request)
}

// QueryOpenBankExternalSubAccountBookBalance
// 第三方子商户电子记账本余额查询接口
//
// 可能返回的错误码:
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
func (c *Client) QueryOpenBankExternalSubAccountBookBalanceWithContext(ctx context.Context, request *QueryOpenBankExternalSubAccountBookBalanceRequest) (response *QueryOpenBankExternalSubAccountBookBalanceResponse, err error) {
    if request == nil {
        request = NewQueryOpenBankExternalSubAccountBookBalanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryOpenBankExternalSubAccountBookBalance require credential")
    }

    request.SetContext(ctx)
    
    response = NewQueryOpenBankExternalSubAccountBookBalanceResponse()
    err = c.Send(request, response)
    return
}

func NewQueryOpenBankExternalSubMerchantBankAccountRequest() (request *QueryOpenBankExternalSubMerchantBankAccountRequest) {
    request = &QueryOpenBankExternalSubMerchantBankAccountRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cpdp", APIVersion, "QueryOpenBankExternalSubMerchantBankAccount")
    
    
    return
}

func NewQueryOpenBankExternalSubMerchantBankAccountResponse() (response *QueryOpenBankExternalSubMerchantBankAccountResponse) {
    response = &QueryOpenBankExternalSubMerchantBankAccountResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// QueryOpenBankExternalSubMerchantBankAccount
// 云企付-子商户银行卡列表查询
//
// 可能返回的错误码:
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
func (c *Client) QueryOpenBankExternalSubMerchantBankAccount(request *QueryOpenBankExternalSubMerchantBankAccountRequest) (response *QueryOpenBankExternalSubMerchantBankAccountResponse, err error) {
    return c.QueryOpenBankExternalSubMerchantBankAccountWithContext(context.Background(), request)
}

// QueryOpenBankExternalSubMerchantBankAccount
// 云企付-子商户银行卡列表查询
//
// 可能返回的错误码:
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
func (c *Client) QueryOpenBankExternalSubMerchantBankAccountWithContext(ctx context.Context, request *QueryOpenBankExternalSubMerchantBankAccountRequest) (response *QueryOpenBankExternalSubMerchantBankAccountResponse, err error) {
    if request == nil {
        request = NewQueryOpenBankExternalSubMerchantBankAccountRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryOpenBankExternalSubMerchantBankAccount require credential")
    }

    request.SetContext(ctx)
    
    response = NewQueryOpenBankExternalSubMerchantBankAccountResponse()
    err = c.Send(request, response)
    return
}

func NewQueryOpenBankExternalSubMerchantRegistrationRequest() (request *QueryOpenBankExternalSubMerchantRegistrationRequest) {
    request = &QueryOpenBankExternalSubMerchantRegistrationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cpdp", APIVersion, "QueryOpenBankExternalSubMerchantRegistration")
    
    
    return
}

func NewQueryOpenBankExternalSubMerchantRegistrationResponse() (response *QueryOpenBankExternalSubMerchantRegistrationResponse) {
    response = &QueryOpenBankExternalSubMerchantRegistrationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// QueryOpenBankExternalSubMerchantRegistration
// 云企付-子商户进件结果查询
//
// 可能返回的错误码:
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
func (c *Client) QueryOpenBankExternalSubMerchantRegistration(request *QueryOpenBankExternalSubMerchantRegistrationRequest) (response *QueryOpenBankExternalSubMerchantRegistrationResponse, err error) {
    return c.QueryOpenBankExternalSubMerchantRegistrationWithContext(context.Background(), request)
}

// QueryOpenBankExternalSubMerchantRegistration
// 云企付-子商户进件结果查询
//
// 可能返回的错误码:
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
func (c *Client) QueryOpenBankExternalSubMerchantRegistrationWithContext(ctx context.Context, request *QueryOpenBankExternalSubMerchantRegistrationRequest) (response *QueryOpenBankExternalSubMerchantRegistrationResponse, err error) {
    if request == nil {
        request = NewQueryOpenBankExternalSubMerchantRegistrationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryOpenBankExternalSubMerchantRegistration require credential")
    }

    request.SetContext(ctx)
    
    response = NewQueryOpenBankExternalSubMerchantRegistrationResponse()
    err = c.Send(request, response)
    return
}

func NewQueryOpenBankOrderDetailReceiptInfoRequest() (request *QueryOpenBankOrderDetailReceiptInfoRequest) {
    request = &QueryOpenBankOrderDetailReceiptInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cpdp", APIVersion, "QueryOpenBankOrderDetailReceiptInfo")
    
    
    return
}

func NewQueryOpenBankOrderDetailReceiptInfoResponse() (response *QueryOpenBankOrderDetailReceiptInfoResponse) {
    response = &QueryOpenBankOrderDetailReceiptInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// QueryOpenBankOrderDetailReceiptInfo
// 云企付-单笔交易回单申请结果查询
//
// 可能返回的错误码:
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
func (c *Client) QueryOpenBankOrderDetailReceiptInfo(request *QueryOpenBankOrderDetailReceiptInfoRequest) (response *QueryOpenBankOrderDetailReceiptInfoResponse, err error) {
    return c.QueryOpenBankOrderDetailReceiptInfoWithContext(context.Background(), request)
}

// QueryOpenBankOrderDetailReceiptInfo
// 云企付-单笔交易回单申请结果查询
//
// 可能返回的错误码:
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
func (c *Client) QueryOpenBankOrderDetailReceiptInfoWithContext(ctx context.Context, request *QueryOpenBankOrderDetailReceiptInfoRequest) (response *QueryOpenBankOrderDetailReceiptInfoResponse, err error) {
    if request == nil {
        request = NewQueryOpenBankOrderDetailReceiptInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryOpenBankOrderDetailReceiptInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewQueryOpenBankOrderDetailReceiptInfoResponse()
    err = c.Send(request, response)
    return
}

func NewQueryOpenBankPaymentOrderRequest() (request *QueryOpenBankPaymentOrderRequest) {
    request = &QueryOpenBankPaymentOrderRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cpdp", APIVersion, "QueryOpenBankPaymentOrder")
    
    
    return
}

func NewQueryOpenBankPaymentOrderResponse() (response *QueryOpenBankPaymentOrderResponse) {
    response = &QueryOpenBankPaymentOrderResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// QueryOpenBankPaymentOrder
// 云企付-查询订单支付结果
//
// 可能返回的错误码:
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
func (c *Client) QueryOpenBankPaymentOrder(request *QueryOpenBankPaymentOrderRequest) (response *QueryOpenBankPaymentOrderResponse, err error) {
    return c.QueryOpenBankPaymentOrderWithContext(context.Background(), request)
}

// QueryOpenBankPaymentOrder
// 云企付-查询订单支付结果
//
// 可能返回的错误码:
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
func (c *Client) QueryOpenBankPaymentOrderWithContext(ctx context.Context, request *QueryOpenBankPaymentOrderRequest) (response *QueryOpenBankPaymentOrderResponse, err error) {
    if request == nil {
        request = NewQueryOpenBankPaymentOrderRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryOpenBankPaymentOrder require credential")
    }

    request.SetContext(ctx)
    
    response = NewQueryOpenBankPaymentOrderResponse()
    err = c.Send(request, response)
    return
}

func NewQueryOpenBankProfitSharePayeeRequest() (request *QueryOpenBankProfitSharePayeeRequest) {
    request = &QueryOpenBankProfitSharePayeeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cpdp", APIVersion, "QueryOpenBankProfitSharePayee")
    
    
    return
}

func NewQueryOpenBankProfitSharePayeeResponse() (response *QueryOpenBankProfitSharePayeeResponse) {
    response = &QueryOpenBankProfitSharePayeeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// QueryOpenBankProfitSharePayee
// 云企付-绑定分账收款方查询
//
// 可能返回的错误码:
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
func (c *Client) QueryOpenBankProfitSharePayee(request *QueryOpenBankProfitSharePayeeRequest) (response *QueryOpenBankProfitSharePayeeResponse, err error) {
    return c.QueryOpenBankProfitSharePayeeWithContext(context.Background(), request)
}

// QueryOpenBankProfitSharePayee
// 云企付-绑定分账收款方查询
//
// 可能返回的错误码:
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
func (c *Client) QueryOpenBankProfitSharePayeeWithContext(ctx context.Context, request *QueryOpenBankProfitSharePayeeRequest) (response *QueryOpenBankProfitSharePayeeResponse, err error) {
    if request == nil {
        request = NewQueryOpenBankProfitSharePayeeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryOpenBankProfitSharePayee require credential")
    }

    request.SetContext(ctx)
    
    response = NewQueryOpenBankProfitSharePayeeResponse()
    err = c.Send(request, response)
    return
}

func NewQueryOpenBankRefundOrderRequest() (request *QueryOpenBankRefundOrderRequest) {
    request = &QueryOpenBankRefundOrderRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cpdp", APIVersion, "QueryOpenBankRefundOrder")
    
    
    return
}

func NewQueryOpenBankRefundOrderResponse() (response *QueryOpenBankRefundOrderResponse) {
    response = &QueryOpenBankRefundOrderResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// QueryOpenBankRefundOrder
// 云企付-退款结果查询
//
// 可能返回的错误码:
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
func (c *Client) QueryOpenBankRefundOrder(request *QueryOpenBankRefundOrderRequest) (response *QueryOpenBankRefundOrderResponse, err error) {
    return c.QueryOpenBankRefundOrderWithContext(context.Background(), request)
}

// QueryOpenBankRefundOrder
// 云企付-退款结果查询
//
// 可能返回的错误码:
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
func (c *Client) QueryOpenBankRefundOrderWithContext(ctx context.Context, request *QueryOpenBankRefundOrderRequest) (response *QueryOpenBankRefundOrderResponse, err error) {
    if request == nil {
        request = NewQueryOpenBankRefundOrderRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryOpenBankRefundOrder require credential")
    }

    request.SetContext(ctx)
    
    response = NewQueryOpenBankRefundOrderResponse()
    err = c.Send(request, response)
    return
}

func NewQueryOpenBankSettleOrderRequest() (request *QueryOpenBankSettleOrderRequest) {
    request = &QueryOpenBankSettleOrderRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cpdp", APIVersion, "QueryOpenBankSettleOrder")
    
    
    return
}

func NewQueryOpenBankSettleOrderResponse() (response *QueryOpenBankSettleOrderResponse) {
    response = &QueryOpenBankSettleOrderResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// QueryOpenBankSettleOrder
// 云企付-结算单查询结果
//
// 可能返回的错误码:
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
func (c *Client) QueryOpenBankSettleOrder(request *QueryOpenBankSettleOrderRequest) (response *QueryOpenBankSettleOrderResponse, err error) {
    return c.QueryOpenBankSettleOrderWithContext(context.Background(), request)
}

// QueryOpenBankSettleOrder
// 云企付-结算单查询结果
//
// 可能返回的错误码:
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
func (c *Client) QueryOpenBankSettleOrderWithContext(ctx context.Context, request *QueryOpenBankSettleOrderRequest) (response *QueryOpenBankSettleOrderResponse, err error) {
    if request == nil {
        request = NewQueryOpenBankSettleOrderRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryOpenBankSettleOrder require credential")
    }

    request.SetContext(ctx)
    
    response = NewQueryOpenBankSettleOrderResponse()
    err = c.Send(request, response)
    return
}

func NewQueryOpenBankSubMerchantCredentialRequest() (request *QueryOpenBankSubMerchantCredentialRequest) {
    request = &QueryOpenBankSubMerchantCredentialRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cpdp", APIVersion, "QueryOpenBankSubMerchantCredential")
    
    
    return
}

func NewQueryOpenBankSubMerchantCredentialResponse() (response *QueryOpenBankSubMerchantCredentialResponse) {
    response = &QueryOpenBankSubMerchantCredentialResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// QueryOpenBankSubMerchantCredential
// 云企付-子商户资质文件查询
//
// 可能返回的错误码:
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
func (c *Client) QueryOpenBankSubMerchantCredential(request *QueryOpenBankSubMerchantCredentialRequest) (response *QueryOpenBankSubMerchantCredentialResponse, err error) {
    return c.QueryOpenBankSubMerchantCredentialWithContext(context.Background(), request)
}

// QueryOpenBankSubMerchantCredential
// 云企付-子商户资质文件查询
//
// 可能返回的错误码:
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
func (c *Client) QueryOpenBankSubMerchantCredentialWithContext(ctx context.Context, request *QueryOpenBankSubMerchantCredentialRequest) (response *QueryOpenBankSubMerchantCredentialResponse, err error) {
    if request == nil {
        request = NewQueryOpenBankSubMerchantCredentialRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryOpenBankSubMerchantCredential require credential")
    }

    request.SetContext(ctx)
    
    response = NewQueryOpenBankSubMerchantCredentialResponse()
    err = c.Send(request, response)
    return
}

func NewQueryOpenBankSubMerchantRateConfigureRequest() (request *QueryOpenBankSubMerchantRateConfigureRequest) {
    request = &QueryOpenBankSubMerchantRateConfigureRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cpdp", APIVersion, "QueryOpenBankSubMerchantRateConfigure")
    
    
    return
}

func NewQueryOpenBankSubMerchantRateConfigureResponse() (response *QueryOpenBankSubMerchantRateConfigureResponse) {
    response = &QueryOpenBankSubMerchantRateConfigureResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// QueryOpenBankSubMerchantRateConfigure
// 云企付-子商户费率配置结果查询
//
// 可能返回的错误码:
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
func (c *Client) QueryOpenBankSubMerchantRateConfigure(request *QueryOpenBankSubMerchantRateConfigureRequest) (response *QueryOpenBankSubMerchantRateConfigureResponse, err error) {
    return c.QueryOpenBankSubMerchantRateConfigureWithContext(context.Background(), request)
}

// QueryOpenBankSubMerchantRateConfigure
// 云企付-子商户费率配置结果查询
//
// 可能返回的错误码:
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
func (c *Client) QueryOpenBankSubMerchantRateConfigureWithContext(ctx context.Context, request *QueryOpenBankSubMerchantRateConfigureRequest) (response *QueryOpenBankSubMerchantRateConfigureResponse, err error) {
    if request == nil {
        request = NewQueryOpenBankSubMerchantRateConfigureRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryOpenBankSubMerchantRateConfigure require credential")
    }

    request.SetContext(ctx)
    
    response = NewQueryOpenBankSubMerchantRateConfigureResponse()
    err = c.Send(request, response)
    return
}

func NewQueryOpenBankSubMerchantSignOnlineRequest() (request *QueryOpenBankSubMerchantSignOnlineRequest) {
    request = &QueryOpenBankSubMerchantSignOnlineRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cpdp", APIVersion, "QueryOpenBankSubMerchantSignOnline")
    
    
    return
}

func NewQueryOpenBankSubMerchantSignOnlineResponse() (response *QueryOpenBankSubMerchantSignOnlineResponse) {
    response = &QueryOpenBankSubMerchantSignOnlineResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// QueryOpenBankSubMerchantSignOnline
// 子商户在线签约查询
//
// 可能返回的错误码:
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
func (c *Client) QueryOpenBankSubMerchantSignOnline(request *QueryOpenBankSubMerchantSignOnlineRequest) (response *QueryOpenBankSubMerchantSignOnlineResponse, err error) {
    return c.QueryOpenBankSubMerchantSignOnlineWithContext(context.Background(), request)
}

// QueryOpenBankSubMerchantSignOnline
// 子商户在线签约查询
//
// 可能返回的错误码:
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
func (c *Client) QueryOpenBankSubMerchantSignOnlineWithContext(ctx context.Context, request *QueryOpenBankSubMerchantSignOnlineRequest) (response *QueryOpenBankSubMerchantSignOnlineResponse, err error) {
    if request == nil {
        request = NewQueryOpenBankSubMerchantSignOnlineRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryOpenBankSubMerchantSignOnline require credential")
    }

    request.SetContext(ctx)
    
    response = NewQueryOpenBankSubMerchantSignOnlineResponse()
    err = c.Send(request, response)
    return
}

func NewQueryOpenBankSupportBankListRequest() (request *QueryOpenBankSupportBankListRequest) {
    request = &QueryOpenBankSupportBankListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cpdp", APIVersion, "QueryOpenBankSupportBankList")
    
    
    return
}

func NewQueryOpenBankSupportBankListResponse() (response *QueryOpenBankSupportBankListResponse) {
    response = &QueryOpenBankSupportBankListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// QueryOpenBankSupportBankList
// 云企付-查询支持银行列表
//
// 可能返回的错误码:
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
func (c *Client) QueryOpenBankSupportBankList(request *QueryOpenBankSupportBankListRequest) (response *QueryOpenBankSupportBankListResponse, err error) {
    return c.QueryOpenBankSupportBankListWithContext(context.Background(), request)
}

// QueryOpenBankSupportBankList
// 云企付-查询支持银行列表
//
// 可能返回的错误码:
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
func (c *Client) QueryOpenBankSupportBankListWithContext(ctx context.Context, request *QueryOpenBankSupportBankListRequest) (response *QueryOpenBankSupportBankListResponse, err error) {
    if request == nil {
        request = NewQueryOpenBankSupportBankListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryOpenBankSupportBankList require credential")
    }

    request.SetContext(ctx)
    
    response = NewQueryOpenBankSupportBankListResponse()
    err = c.Send(request, response)
    return
}

func NewQueryOpenBankUnbindExternalSubMerchantBankAccountRequest() (request *QueryOpenBankUnbindExternalSubMerchantBankAccountRequest) {
    request = &QueryOpenBankUnbindExternalSubMerchantBankAccountRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cpdp", APIVersion, "QueryOpenBankUnbindExternalSubMerchantBankAccount")
    
    
    return
}

func NewQueryOpenBankUnbindExternalSubMerchantBankAccountResponse() (response *QueryOpenBankUnbindExternalSubMerchantBankAccountResponse) {
    response = &QueryOpenBankUnbindExternalSubMerchantBankAccountResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// QueryOpenBankUnbindExternalSubMerchantBankAccount
// 云企付-子商户银行卡解绑结果查询
//
// 可能返回的错误码:
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
func (c *Client) QueryOpenBankUnbindExternalSubMerchantBankAccount(request *QueryOpenBankUnbindExternalSubMerchantBankAccountRequest) (response *QueryOpenBankUnbindExternalSubMerchantBankAccountResponse, err error) {
    return c.QueryOpenBankUnbindExternalSubMerchantBankAccountWithContext(context.Background(), request)
}

// QueryOpenBankUnbindExternalSubMerchantBankAccount
// 云企付-子商户银行卡解绑结果查询
//
// 可能返回的错误码:
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
func (c *Client) QueryOpenBankUnbindExternalSubMerchantBankAccountWithContext(ctx context.Context, request *QueryOpenBankUnbindExternalSubMerchantBankAccountRequest) (response *QueryOpenBankUnbindExternalSubMerchantBankAccountResponse, err error) {
    if request == nil {
        request = NewQueryOpenBankUnbindExternalSubMerchantBankAccountRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryOpenBankUnbindExternalSubMerchantBankAccount require credential")
    }

    request.SetContext(ctx)
    
    response = NewQueryOpenBankUnbindExternalSubMerchantBankAccountResponse()
    err = c.Send(request, response)
    return
}

func NewQueryOpenBankVerificationOrderRequest() (request *QueryOpenBankVerificationOrderRequest) {
    request = &QueryOpenBankVerificationOrderRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cpdp", APIVersion, "QueryOpenBankVerificationOrder")
    
    
    return
}

func NewQueryOpenBankVerificationOrderResponse() (response *QueryOpenBankVerificationOrderResponse) {
    response = &QueryOpenBankVerificationOrderResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// QueryOpenBankVerificationOrder
// 云企付-查询核销订单状态，客户可以使用该接口来查询核销申请的订单状态。目前仅支持TENPAY渠道EBANK_PAYMENT付款方式的担保支付订单查询。
//
// 可能返回的错误码:
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
func (c *Client) QueryOpenBankVerificationOrder(request *QueryOpenBankVerificationOrderRequest) (response *QueryOpenBankVerificationOrderResponse, err error) {
    return c.QueryOpenBankVerificationOrderWithContext(context.Background(), request)
}

// QueryOpenBankVerificationOrder
// 云企付-查询核销订单状态，客户可以使用该接口来查询核销申请的订单状态。目前仅支持TENPAY渠道EBANK_PAYMENT付款方式的担保支付订单查询。
//
// 可能返回的错误码:
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
func (c *Client) QueryOpenBankVerificationOrderWithContext(ctx context.Context, request *QueryOpenBankVerificationOrderRequest) (response *QueryOpenBankVerificationOrderResponse, err error) {
    if request == nil {
        request = NewQueryOpenBankVerificationOrderRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryOpenBankVerificationOrder require credential")
    }

    request.SetContext(ctx)
    
    response = NewQueryOpenBankVerificationOrderResponse()
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
//  AUTHFAILURE_SECRETKEYNOTFOUND = "AuthFailure.SecretKeyNotFound"
//  AUTHFAILURE_VERIFYERROR = "AuthFailure.VerifyError"
//  AUTHFAILURE_VERIFYTOKENFAILURE = "AuthFailure.VerifyTokenFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_APPDENY = "FailedOperation.AppDeny"
//  FAILEDOPERATION_NORECORD = "FailedOperation.NoRecord"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) QueryOrder(request *QueryOrderRequest) (response *QueryOrderResponse, err error) {
    return c.QueryOrderWithContext(context.Background(), request)
}

// QueryOrder
// 根据订单号，或者用户Id，查询支付订单状态 
//
// 可能返回的错误码:
//  AUTHFAILURE_SECRETKEYNOTFOUND = "AuthFailure.SecretKeyNotFound"
//  AUTHFAILURE_VERIFYERROR = "AuthFailure.VerifyError"
//  AUTHFAILURE_VERIFYTOKENFAILURE = "AuthFailure.VerifyTokenFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_APPDENY = "FailedOperation.AppDeny"
//  FAILEDOPERATION_NORECORD = "FailedOperation.NoRecord"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) QueryOrderWithContext(ctx context.Context, request *QueryOrderRequest) (response *QueryOrderResponse, err error) {
    if request == nil {
        request = NewQueryOrderRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryOrder require credential")
    }

    request.SetContext(ctx)
    
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
//  AUTHFAILURE_SECRETKEYNOTFOUND = "AuthFailure.SecretKeyNotFound"
//  AUTHFAILURE_VERIFYERROR = "AuthFailure.VerifyError"
//  AUTHFAILURE_VERIFYTOKENFAILURE = "AuthFailure.VerifyTokenFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_APPDENY = "FailedOperation.AppDeny"
//  FAILEDOPERATION_NORECORD = "FailedOperation.NoRecord"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) QueryOrderStatus(request *QueryOrderStatusRequest) (response *QueryOrderStatusResponse, err error) {
    return c.QueryOrderStatusWithContext(context.Background(), request)
}

// QueryOrderStatus
// 云支付-查询订单付款状态
//
// 可能返回的错误码:
//  AUTHFAILURE_SECRETKEYNOTFOUND = "AuthFailure.SecretKeyNotFound"
//  AUTHFAILURE_VERIFYERROR = "AuthFailure.VerifyError"
//  AUTHFAILURE_VERIFYTOKENFAILURE = "AuthFailure.VerifyTokenFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_APPDENY = "FailedOperation.AppDeny"
//  FAILEDOPERATION_NORECORD = "FailedOperation.NoRecord"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) QueryOrderStatusWithContext(ctx context.Context, request *QueryOrderStatusRequest) (response *QueryOrderStatusResponse, err error) {
    if request == nil {
        request = NewQueryOrderStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryOrderStatus require credential")
    }

    request.SetContext(ctx)
    
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
    return c.QueryOutwardOrderWithContext(context.Background(), request)
}

// QueryOutwardOrder
// 跨境-查询汇出结果
//
// 可能返回的错误码:
//  INTERNALERROR_PARAMETERERROR = "InternalError.ParameterError"
//  INTERNALERROR_UNKOWNERROR = "InternalError.UnkownError"
//  RESOURCENOTFOUND_MERCHANTINFONOTFOUND = "ResourceNotFound.MerchantInfoNotFound"
func (c *Client) QueryOutwardOrderWithContext(ctx context.Context, request *QueryOutwardOrderRequest) (response *QueryOutwardOrderResponse, err error) {
    if request == nil {
        request = NewQueryOutwardOrderRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryOutwardOrder require credential")
    }

    request.SetContext(ctx)
    
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
    return c.QueryPayerInfoWithContext(context.Background(), request)
}

// QueryPayerInfo
// 跨境-付款人查询
//
// 可能返回的错误码:
//  INTERNALERROR_PARAMETERERROR = "InternalError.ParameterError"
//  INTERNALERROR_UNKOWNERROR = "InternalError.UnkownError"
//  RESOURCENOTFOUND_MERCHANTINFONOTFOUND = "ResourceNotFound.MerchantInfoNotFound"
func (c *Client) QueryPayerInfoWithContext(ctx context.Context, request *QueryPayerInfoRequest) (response *QueryPayerInfoResponse, err error) {
    if request == nil {
        request = NewQueryPayerInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryPayerInfo require credential")
    }

    request.SetContext(ctx)
    
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
    return c.QueryReconciliationDocumentWithContext(context.Background(), request)
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
func (c *Client) QueryReconciliationDocumentWithContext(ctx context.Context, request *QueryReconciliationDocumentRequest) (response *QueryReconciliationDocumentResponse, err error) {
    if request == nil {
        request = NewQueryReconciliationDocumentRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryReconciliationDocument require credential")
    }

    request.SetContext(ctx)
    
    response = NewQueryReconciliationDocumentResponse()
    err = c.Send(request, response)
    return
}

func NewQueryReconciliationFileApplyInfoRequest() (request *QueryReconciliationFileApplyInfoRequest) {
    request = &QueryReconciliationFileApplyInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cpdp", APIVersion, "QueryReconciliationFileApplyInfo")
    
    
    return
}

func NewQueryReconciliationFileApplyInfoResponse() (response *QueryReconciliationFileApplyInfoResponse) {
    response = &QueryReconciliationFileApplyInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// QueryReconciliationFileApplyInfo
// 聚鑫-查询对账文件申请结果
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
func (c *Client) QueryReconciliationFileApplyInfo(request *QueryReconciliationFileApplyInfoRequest) (response *QueryReconciliationFileApplyInfoResponse, err error) {
    return c.QueryReconciliationFileApplyInfoWithContext(context.Background(), request)
}

// QueryReconciliationFileApplyInfo
// 聚鑫-查询对账文件申请结果
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
func (c *Client) QueryReconciliationFileApplyInfoWithContext(ctx context.Context, request *QueryReconciliationFileApplyInfoRequest) (response *QueryReconciliationFileApplyInfoResponse, err error) {
    if request == nil {
        request = NewQueryReconciliationFileApplyInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryReconciliationFileApplyInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewQueryReconciliationFileApplyInfoResponse()
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
//  AUTHFAILURE_SECRETKEYNOTFOUND = "AuthFailure.SecretKeyNotFound"
//  AUTHFAILURE_VERIFYERROR = "AuthFailure.VerifyError"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_APPDENY = "FailedOperation.AppDeny"
//  FAILEDOPERATION_NORECORD = "FailedOperation.NoRecord"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) QueryRefund(request *QueryRefundRequest) (response *QueryRefundResponse, err error) {
    return c.QueryRefundWithContext(context.Background(), request)
}

// QueryRefund
// 提交退款申请后，通过调用该接口查询退款状态。退款可能有一定延时，用微信零钱支付的退款约20分钟内到账，银行卡支付的退款约3个工作日后到账。
//
// 可能返回的错误码:
//  AUTHFAILURE_SECRETKEYNOTFOUND = "AuthFailure.SecretKeyNotFound"
//  AUTHFAILURE_VERIFYERROR = "AuthFailure.VerifyError"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_APPDENY = "FailedOperation.AppDeny"
//  FAILEDOPERATION_NORECORD = "FailedOperation.NoRecord"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) QueryRefundWithContext(ctx context.Context, request *QueryRefundRequest) (response *QueryRefundResponse, err error) {
    if request == nil {
        request = NewQueryRefundRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryRefund require credential")
    }

    request.SetContext(ctx)
    
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
//  AUTHFAILURE_SECRETKEYNOTFOUND = "AuthFailure.SecretKeyNotFound"
//  AUTHFAILURE_VERIFYERROR = "AuthFailure.VerifyError"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_APPDENY = "FailedOperation.AppDeny"
//  FAILEDOPERATION_NORECORD = "FailedOperation.NoRecord"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) QueryShopOpenId(request *QueryShopOpenIdRequest) (response *QueryShopOpenIdResponse, err error) {
    return c.QueryShopOpenIdWithContext(context.Background(), request)
}

// QueryShopOpenId
// 云支付-获取门店OpenId接口
//
// 可能返回的错误码:
//  AUTHFAILURE_SECRETKEYNOTFOUND = "AuthFailure.SecretKeyNotFound"
//  AUTHFAILURE_VERIFYERROR = "AuthFailure.VerifyError"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_APPDENY = "FailedOperation.AppDeny"
//  FAILEDOPERATION_NORECORD = "FailedOperation.NoRecord"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) QueryShopOpenIdWithContext(ctx context.Context, request *QueryShopOpenIdRequest) (response *QueryShopOpenIdResponse, err error) {
    if request == nil {
        request = NewQueryShopOpenIdRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryShopOpenId require credential")
    }

    request.SetContext(ctx)
    
    response = NewQueryShopOpenIdResponse()
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
//  AUTHFAILURE_SECRETKEYNOTFOUND = "AuthFailure.SecretKeyNotFound"
//  AUTHFAILURE_VERIFYERROR = "AuthFailure.VerifyError"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_APPDENY = "FailedOperation.AppDeny"
//  FAILEDOPERATION_NORECORD = "FailedOperation.NoRecord"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) QuerySinglePaymentResult(request *QuerySinglePaymentResultRequest) (response *QuerySinglePaymentResultResponse, err error) {
    return c.QuerySinglePaymentResultWithContext(context.Background(), request)
}

// QuerySinglePaymentResult
// 灵云-单笔转账结果查询
//
// 可能返回的错误码:
//  AUTHFAILURE_SECRETKEYNOTFOUND = "AuthFailure.SecretKeyNotFound"
//  AUTHFAILURE_VERIFYERROR = "AuthFailure.VerifyError"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_APPDENY = "FailedOperation.AppDeny"
//  FAILEDOPERATION_NORECORD = "FailedOperation.NoRecord"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) QuerySinglePaymentResultWithContext(ctx context.Context, request *QuerySinglePaymentResultRequest) (response *QuerySinglePaymentResultResponse, err error) {
    if request == nil {
        request = NewQuerySinglePaymentResultRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QuerySinglePaymentResult require credential")
    }

    request.SetContext(ctx)
    
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
    return c.QuerySingleTransactionStatusWithContext(context.Background(), request)
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
func (c *Client) QuerySingleTransactionStatusWithContext(ctx context.Context, request *QuerySingleTransactionStatusRequest) (response *QuerySingleTransactionStatusResponse, err error) {
    if request == nil {
        request = NewQuerySingleTransactionStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QuerySingleTransactionStatus require credential")
    }

    request.SetContext(ctx)
    
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
    return c.QuerySmallAmountTransferWithContext(context.Background(), request)
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
func (c *Client) QuerySmallAmountTransferWithContext(ctx context.Context, request *QuerySmallAmountTransferRequest) (response *QuerySmallAmountTransferResponse, err error) {
    if request == nil {
        request = NewQuerySmallAmountTransferRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QuerySmallAmountTransfer require credential")
    }

    request.SetContext(ctx)
    
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
    return c.QueryTradeWithContext(context.Background(), request)
}

// QueryTrade
// 跨境-贸易材料明细查询。
//
// 可能返回的错误码:
//  INTERNALERROR_PARAMETERERROR = "InternalError.ParameterError"
//  INTERNALERROR_UNKOWNERROR = "InternalError.UnkownError"
//  RESOURCENOTFOUND_MERCHANTINFONOTFOUND = "ResourceNotFound.MerchantInfoNotFound"
func (c *Client) QueryTradeWithContext(ctx context.Context, request *QueryTradeRequest) (response *QueryTradeResponse, err error) {
    if request == nil {
        request = NewQueryTradeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryTrade require credential")
    }

    request.SetContext(ctx)
    
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
    return c.QueryTransferBatchWithContext(context.Background(), request)
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
func (c *Client) QueryTransferBatchWithContext(ctx context.Context, request *QueryTransferBatchRequest) (response *QueryTransferBatchResponse, err error) {
    if request == nil {
        request = NewQueryTransferBatchRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryTransferBatch require credential")
    }

    request.SetContext(ctx)
    
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
    return c.QueryTransferDetailWithContext(context.Background(), request)
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
func (c *Client) QueryTransferDetailWithContext(ctx context.Context, request *QueryTransferDetailRequest) (response *QueryTransferDetailResponse, err error) {
    if request == nil {
        request = NewQueryTransferDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryTransferDetail require credential")
    }

    request.SetContext(ctx)
    
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
    return c.QueryTransferResultWithContext(context.Background(), request)
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
func (c *Client) QueryTransferResultWithContext(ctx context.Context, request *QueryTransferResultRequest) (response *QueryTransferResultResponse, err error) {
    if request == nil {
        request = NewQueryTransferResultRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryTransferResult require credential")
    }

    request.SetContext(ctx)
    
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
    return c.RechargeByThirdPayWithContext(context.Background(), request)
}

// RechargeByThirdPay
// 会员在途充值(经第三方支付渠道)接口
//
// 可能返回的错误码:
//  FAILEDOPERATION_ACTIONINVALID = "FailedOperation.ActionInvalid"
//  FAILEDOPERATION_PABANKERROR = "FailedOperation.PABankError"
func (c *Client) RechargeByThirdPayWithContext(ctx context.Context, request *RechargeByThirdPayRequest) (response *RechargeByThirdPayResponse, err error) {
    if request == nil {
        request = NewRechargeByThirdPayRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RechargeByThirdPay require credential")
    }

    request.SetContext(ctx)
    
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
    return c.RechargeMemberThirdPayWithContext(context.Background(), request)
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
func (c *Client) RechargeMemberThirdPayWithContext(ctx context.Context, request *RechargeMemberThirdPayRequest) (response *RechargeMemberThirdPayResponse, err error) {
    if request == nil {
        request = NewRechargeMemberThirdPayRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RechargeMemberThirdPay require credential")
    }

    request.SetContext(ctx)
    
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
//  AUTHFAILURE_VERIFYERROR = "AuthFailure.VerifyError"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ABNORMALMERCHANTSTATE = "FailedOperation.AbnormalMerchantState"
//  FAILEDOPERATION_ABNORMALORDERSTATE = "FailedOperation.AbnormalOrderState"
//  FAILEDOPERATION_APPDENY = "FailedOperation.AppDeny"
//  FAILEDOPERATION_BANLANCENOTENOUGHERROR = "FailedOperation.BanlanceNotEnoughError"
//  FAILEDOPERATION_CHANNELREFUNDFAILED = "FailedOperation.ChannelRefundFailed"
//  FAILEDOPERATION_CHANNELREFUNDFREQUENCYLIMITED = "FailedOperation.ChannelRefundFrequencyLimited"
//  FAILEDOPERATION_INTERNALSERVICETIMEOUT = "FailedOperation.InternalServiceTimeout"
//  FAILEDOPERATION_INVALIDREFUNDAMT = "FailedOperation.InvalidRefundAmt"
//  FAILEDOPERATION_MERCHANTBALANCENOTENOUGH = "FailedOperation.MerchantBalanceNotEnough"
//  FAILEDOPERATION_MERCHANTNOTEXISTS = "FailedOperation.MerchantNotExists"
//  FAILEDOPERATION_NORECORD = "FailedOperation.NoRecord"
//  FAILEDOPERATION_ORDERLOCKED = "FailedOperation.OrderLocked"
//  FAILEDOPERATION_PORTALERROR = "FailedOperation.PortalError"
//  FAILEDOPERATION_REFUNDINFODUPLICATE = "FailedOperation.RefundInfoDuplicate"
//  FAILEDOPERATION_REFUNDNOTRETRIEABLE = "FailedOperation.RefundNotRetrieable"
//  FAILEDOPERATION_REFUNDPROCESSING = "FailedOperation.RefundProcessIng"
//  FAILEDOPERATION_REFUNDTRANSACTIONCLOSED = "FailedOperation.RefundTransactionClosed"
//  FAILEDOPERATION_REFUNDTRANSACTIONFINISHED = "FailedOperation.RefundTransactionFinished"
//  FAILEDOPERATION_WXCRTNOTSET = "FailedOperation.WxCrtNotSet"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) Refund(request *RefundRequest) (response *RefundResponse, err error) {
    return c.RefundWithContext(context.Background(), request)
}

// Refund
// 如交易订单需退款，可以通过本接口将支付款全部或部分退还给付款方，聚鑫将在收到退款请求并且验证成功之后，按照退款规则将支付款按原路退回到支付帐号。最长支持1年的订单退款。在订单包含多个子订单的情况下，如果使用本接口传入OutTradeNo或TransactionId退款，则只支持全单退款；如果需要部分退款，请通过传入子订单的方式来指定部分金额退款。 
//
// 可能返回的错误码:
//  AUTHFAILURE_VERIFYERROR = "AuthFailure.VerifyError"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ABNORMALMERCHANTSTATE = "FailedOperation.AbnormalMerchantState"
//  FAILEDOPERATION_ABNORMALORDERSTATE = "FailedOperation.AbnormalOrderState"
//  FAILEDOPERATION_APPDENY = "FailedOperation.AppDeny"
//  FAILEDOPERATION_BANLANCENOTENOUGHERROR = "FailedOperation.BanlanceNotEnoughError"
//  FAILEDOPERATION_CHANNELREFUNDFAILED = "FailedOperation.ChannelRefundFailed"
//  FAILEDOPERATION_CHANNELREFUNDFREQUENCYLIMITED = "FailedOperation.ChannelRefundFrequencyLimited"
//  FAILEDOPERATION_INTERNALSERVICETIMEOUT = "FailedOperation.InternalServiceTimeout"
//  FAILEDOPERATION_INVALIDREFUNDAMT = "FailedOperation.InvalidRefundAmt"
//  FAILEDOPERATION_MERCHANTBALANCENOTENOUGH = "FailedOperation.MerchantBalanceNotEnough"
//  FAILEDOPERATION_MERCHANTNOTEXISTS = "FailedOperation.MerchantNotExists"
//  FAILEDOPERATION_NORECORD = "FailedOperation.NoRecord"
//  FAILEDOPERATION_ORDERLOCKED = "FailedOperation.OrderLocked"
//  FAILEDOPERATION_PORTALERROR = "FailedOperation.PortalError"
//  FAILEDOPERATION_REFUNDINFODUPLICATE = "FailedOperation.RefundInfoDuplicate"
//  FAILEDOPERATION_REFUNDNOTRETRIEABLE = "FailedOperation.RefundNotRetrieable"
//  FAILEDOPERATION_REFUNDPROCESSING = "FailedOperation.RefundProcessIng"
//  FAILEDOPERATION_REFUNDTRANSACTIONCLOSED = "FailedOperation.RefundTransactionClosed"
//  FAILEDOPERATION_REFUNDTRANSACTIONFINISHED = "FailedOperation.RefundTransactionFinished"
//  FAILEDOPERATION_WXCRTNOTSET = "FailedOperation.WxCrtNotSet"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) RefundWithContext(ctx context.Context, request *RefundRequest) (response *RefundResponse, err error) {
    if request == nil {
        request = NewRefundRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("Refund require credential")
    }

    request.SetContext(ctx)
    
    response = NewRefundResponse()
    err = c.Send(request, response)
    return
}

func NewRefundCloudOrderRequest() (request *RefundCloudOrderRequest) {
    request = &RefundCloudOrderRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cpdp", APIVersion, "RefundCloudOrder")
    
    
    return
}

func NewRefundCloudOrderResponse() (response *RefundCloudOrderResponse) {
    response = &RefundCloudOrderResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// RefundCloudOrder
// 如交易订单需退款，可以通过本接口将支付款全部或部分退还给付款方，聚鑫将在收到退款请求并且验证成功之后，按照退款规则将支付款按原路退回到支付帐号。最长支持1年的订单退款。在订单包含多个子订单的情况下，如果使用本接口传入OutTradeNo或TransactionId退款，则只支持全单退款；如果需要部分退款，请通过传入子订单的方式来指定部分金额退款。 
//
// 可能返回的错误码:
//  AUTHFAILURE_MIDAS = "AuthFailure.Midas"
//  FAILEDOPERATION_ACTION = "FailedOperation.Action"
//  FAILEDOPERATION_MIDASINTERNALERROR = "FailedOperation.MidasInternalError"
//  FAILEDOPERATION_MIDASINVALIDREQUEST = "FailedOperation.MidasInvalidRequest"
//  FAILEDOPERATION_MIDASNEEDRETRY = "FailedOperation.MidasNeedRetry"
//  FAILEDOPERATION_MIDASREGISTERUNFINISHED = "FailedOperation.MidasRegisterUnfinished"
//  FAILEDOPERATION_MIDASREPEATORDER = "FailedOperation.MidasRepeatOrder"
//  FAILEDOPERATION_MIDASRISK = "FailedOperation.MidasRisk"
//  FAILEDOPERATION_MIDASSTATUSNOTMATCH = "FailedOperation.MidasStatusNotMatch"
//  FAILEDOPERATION_MIDASUNSUPPORTEDACTION = "FailedOperation.MidasUnsupportedAction"
//  INTERNALERROR_MIDAS = "InternalError.Midas"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  INVALIDPARAMETER_MIDAS = "InvalidParameter.Midas"
//  INVALIDPARAMETER_MIDASENVIRONMENT = "InvalidParameter.MidasEnvironment"
//  INVALIDPARAMETER_MIDASEXTERNALAPP = "InvalidParameter.MidasExternalApp"
//  INVALIDPARAMETER_MIDASFILETYPE = "InvalidParameter.MidasFileType"
//  INVALIDPARAMETER_MIDASHASH = "InvalidParameter.MidasHash"
//  INVALIDPARAMETER_MIDASSIGNID = "InvalidParameter.MidasSignId"
//  LIMITEXCEEDED_MIDASLARGEFILE = "LimitExceeded.MidasLargeFile"
//  LIMITEXCEEDED_MIDASORDER = "LimitExceeded.MidasOrder"
//  LIMITEXCEEDED_MIDASORDERCANCELED = "LimitExceeded.MidasOrderCanceled"
//  LIMITEXCEEDED_MIDASORDERCLOSED = "LimitExceeded.MidasOrderClosed"
//  LIMITEXCEEDED_MIDASORDEREXPIRED = "LimitExceeded.MidasOrderExpired"
//  LIMITEXCEEDED_MIDASORDERFAILED = "LimitExceeded.MidasOrderFailed"
//  LIMITEXCEEDED_MIDASORDERPARTIALSUCCESS = "LimitExceeded.MidasOrderPartialSuccess"
//  LIMITEXCEEDED_MIDASORDERPROCESSING = "LimitExceeded.MidasOrderProcessing"
//  LIMITEXCEEDED_MIDASORDERSUCCESS = "LimitExceeded.MidasOrderSuccess"
//  REQUESTLIMITEXCEEDED_MIDAS = "RequestLimitExceeded.Midas"
//  REQUESTLIMITEXCEEDED_MIDASINVALIDREQUEST = "RequestLimitExceeded.MidasInvalidRequest"
//  RESOURCEINUSE_MIDAS = "ResourceInUse.Midas"
//  RESOURCENOTFOUND_ACCOUNT = "ResourceNotFound.Account"
//  RESOURCENOTFOUND_MIDASEXTERNALAPP = "ResourceNotFound.MidasExternalApp"
//  RESOURCENOTFOUND_MIDASEXTERNALORDER = "ResourceNotFound.MidasExternalOrder"
//  RESOURCENOTFOUND_MIDASORDER = "ResourceNotFound.MidasOrder"
//  RESOURCENOTFOUND_MIDASSIGN = "ResourceNotFound.MidasSign"
//  RESOURCEUNAVAILABLE_MIDASBALANCE = "ResourceUnavailable.MidasBalance"
//  RESOURCEUNAVAILABLE_MIDASDAY = "ResourceUnavailable.MidasDay"
//  RESOURCEUNAVAILABLE_MIDASFROZENAMOUNT = "ResourceUnavailable.MidasFrozenAmount"
//  RESOURCEUNAVAILABLE_MIDASMERCHANTBALANCE = "ResourceUnavailable.MidasMerchantBalance"
//  RESOURCEUNAVAILABLE_MIDASORDER = "ResourceUnavailable.MidasOrder"
//  RESOURCEUNAVAILABLE_MIDASUSERBALANCE = "ResourceUnavailable.MidasUserBalance"
//  RESOURCEUNAVAILABLE_MIDASWALLET = "ResourceUnavailable.MidasWallet"
//  UNAUTHORIZEDOPERATION_MIDAS = "UnauthorizedOperation.Midas"
func (c *Client) RefundCloudOrder(request *RefundCloudOrderRequest) (response *RefundCloudOrderResponse, err error) {
    return c.RefundCloudOrderWithContext(context.Background(), request)
}

// RefundCloudOrder
// 如交易订单需退款，可以通过本接口将支付款全部或部分退还给付款方，聚鑫将在收到退款请求并且验证成功之后，按照退款规则将支付款按原路退回到支付帐号。最长支持1年的订单退款。在订单包含多个子订单的情况下，如果使用本接口传入OutTradeNo或TransactionId退款，则只支持全单退款；如果需要部分退款，请通过传入子订单的方式来指定部分金额退款。 
//
// 可能返回的错误码:
//  AUTHFAILURE_MIDAS = "AuthFailure.Midas"
//  FAILEDOPERATION_ACTION = "FailedOperation.Action"
//  FAILEDOPERATION_MIDASINTERNALERROR = "FailedOperation.MidasInternalError"
//  FAILEDOPERATION_MIDASINVALIDREQUEST = "FailedOperation.MidasInvalidRequest"
//  FAILEDOPERATION_MIDASNEEDRETRY = "FailedOperation.MidasNeedRetry"
//  FAILEDOPERATION_MIDASREGISTERUNFINISHED = "FailedOperation.MidasRegisterUnfinished"
//  FAILEDOPERATION_MIDASREPEATORDER = "FailedOperation.MidasRepeatOrder"
//  FAILEDOPERATION_MIDASRISK = "FailedOperation.MidasRisk"
//  FAILEDOPERATION_MIDASSTATUSNOTMATCH = "FailedOperation.MidasStatusNotMatch"
//  FAILEDOPERATION_MIDASUNSUPPORTEDACTION = "FailedOperation.MidasUnsupportedAction"
//  INTERNALERROR_MIDAS = "InternalError.Midas"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  INVALIDPARAMETER_MIDAS = "InvalidParameter.Midas"
//  INVALIDPARAMETER_MIDASENVIRONMENT = "InvalidParameter.MidasEnvironment"
//  INVALIDPARAMETER_MIDASEXTERNALAPP = "InvalidParameter.MidasExternalApp"
//  INVALIDPARAMETER_MIDASFILETYPE = "InvalidParameter.MidasFileType"
//  INVALIDPARAMETER_MIDASHASH = "InvalidParameter.MidasHash"
//  INVALIDPARAMETER_MIDASSIGNID = "InvalidParameter.MidasSignId"
//  LIMITEXCEEDED_MIDASLARGEFILE = "LimitExceeded.MidasLargeFile"
//  LIMITEXCEEDED_MIDASORDER = "LimitExceeded.MidasOrder"
//  LIMITEXCEEDED_MIDASORDERCANCELED = "LimitExceeded.MidasOrderCanceled"
//  LIMITEXCEEDED_MIDASORDERCLOSED = "LimitExceeded.MidasOrderClosed"
//  LIMITEXCEEDED_MIDASORDEREXPIRED = "LimitExceeded.MidasOrderExpired"
//  LIMITEXCEEDED_MIDASORDERFAILED = "LimitExceeded.MidasOrderFailed"
//  LIMITEXCEEDED_MIDASORDERPARTIALSUCCESS = "LimitExceeded.MidasOrderPartialSuccess"
//  LIMITEXCEEDED_MIDASORDERPROCESSING = "LimitExceeded.MidasOrderProcessing"
//  LIMITEXCEEDED_MIDASORDERSUCCESS = "LimitExceeded.MidasOrderSuccess"
//  REQUESTLIMITEXCEEDED_MIDAS = "RequestLimitExceeded.Midas"
//  REQUESTLIMITEXCEEDED_MIDASINVALIDREQUEST = "RequestLimitExceeded.MidasInvalidRequest"
//  RESOURCEINUSE_MIDAS = "ResourceInUse.Midas"
//  RESOURCENOTFOUND_ACCOUNT = "ResourceNotFound.Account"
//  RESOURCENOTFOUND_MIDASEXTERNALAPP = "ResourceNotFound.MidasExternalApp"
//  RESOURCENOTFOUND_MIDASEXTERNALORDER = "ResourceNotFound.MidasExternalOrder"
//  RESOURCENOTFOUND_MIDASORDER = "ResourceNotFound.MidasOrder"
//  RESOURCENOTFOUND_MIDASSIGN = "ResourceNotFound.MidasSign"
//  RESOURCEUNAVAILABLE_MIDASBALANCE = "ResourceUnavailable.MidasBalance"
//  RESOURCEUNAVAILABLE_MIDASDAY = "ResourceUnavailable.MidasDay"
//  RESOURCEUNAVAILABLE_MIDASFROZENAMOUNT = "ResourceUnavailable.MidasFrozenAmount"
//  RESOURCEUNAVAILABLE_MIDASMERCHANTBALANCE = "ResourceUnavailable.MidasMerchantBalance"
//  RESOURCEUNAVAILABLE_MIDASORDER = "ResourceUnavailable.MidasOrder"
//  RESOURCEUNAVAILABLE_MIDASUSERBALANCE = "ResourceUnavailable.MidasUserBalance"
//  RESOURCEUNAVAILABLE_MIDASWALLET = "ResourceUnavailable.MidasWallet"
//  UNAUTHORIZEDOPERATION_MIDAS = "UnauthorizedOperation.Midas"
func (c *Client) RefundCloudOrderWithContext(ctx context.Context, request *RefundCloudOrderRequest) (response *RefundCloudOrderResponse, err error) {
    if request == nil {
        request = NewRefundCloudOrderRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RefundCloudOrder require credential")
    }

    request.SetContext(ctx)
    
    response = NewRefundCloudOrderResponse()
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
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ACTIONINVALID = "FailedOperation.ActionInvalid"
//  FAILEDOPERATION_PABANKERROR = "FailedOperation.PABankError"
func (c *Client) RefundMemberTransaction(request *RefundMemberTransactionRequest) (response *RefundMemberTransactionResponse, err error) {
    return c.RefundMemberTransactionWithContext(context.Background(), request)
}

// RefundMemberTransaction
// 会员间交易退款
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ACTIONINVALID = "FailedOperation.ActionInvalid"
//  FAILEDOPERATION_PABANKERROR = "FailedOperation.PABankError"
func (c *Client) RefundMemberTransactionWithContext(ctx context.Context, request *RefundMemberTransactionRequest) (response *RefundMemberTransactionResponse, err error) {
    if request == nil {
        request = NewRefundMemberTransactionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RefundMemberTransaction require credential")
    }

    request.SetContext(ctx)
    
    response = NewRefundMemberTransactionResponse()
    err = c.Send(request, response)
    return
}

func NewRefundOpenBankOrderRequest() (request *RefundOpenBankOrderRequest) {
    request = &RefundOpenBankOrderRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cpdp", APIVersion, "RefundOpenBankOrder")
    
    
    return
}

func NewRefundOpenBankOrderResponse() (response *RefundOpenBankOrderResponse) {
    response = &RefundOpenBankOrderResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// RefundOpenBankOrder
// 云企付-退款申请
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ACTIONINVALID = "FailedOperation.ActionInvalid"
//  FAILEDOPERATION_PABANKERROR = "FailedOperation.PABankError"
func (c *Client) RefundOpenBankOrder(request *RefundOpenBankOrderRequest) (response *RefundOpenBankOrderResponse, err error) {
    return c.RefundOpenBankOrderWithContext(context.Background(), request)
}

// RefundOpenBankOrder
// 云企付-退款申请
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ACTIONINVALID = "FailedOperation.ActionInvalid"
//  FAILEDOPERATION_PABANKERROR = "FailedOperation.PABankError"
func (c *Client) RefundOpenBankOrderWithContext(ctx context.Context, request *RefundOpenBankOrderRequest) (response *RefundOpenBankOrderResponse, err error) {
    if request == nil {
        request = NewRefundOpenBankOrderRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RefundOpenBankOrder require credential")
    }

    request.SetContext(ctx)
    
    response = NewRefundOpenBankOrderResponse()
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
    return c.RefundOrderWithContext(context.Background(), request)
}

// RefundOrder
// 云鉴-消费订单退款的接口
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDPARAMETER = "FailedOperation.InvalidParameter"
//  FAILEDOPERATION_MISSINGPARAMETER = "FailedOperation.MissingParameter"
//  FAILEDOPERATION_ORDERREFUNDERROR = "FailedOperation.OrderRefundError"
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
func (c *Client) RefundOrderWithContext(ctx context.Context, request *RefundOrderRequest) (response *RefundOrderResponse, err error) {
    if request == nil {
        request = NewRefundOrderRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RefundOrder require credential")
    }

    request.SetContext(ctx)
    
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
    return c.RefundTlinxOrderWithContext(context.Background(), request)
}

// RefundTlinxOrder
// 云支付订单退款接口
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDPARAMETER = "FailedOperation.InvalidParameter"
//  FAILEDOPERATION_MISSINGPARAMETER = "FailedOperation.MissingParameter"
//  FAILEDOPERATION_ORDERREFUNDERROR = "FailedOperation.OrderRefundError"
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
func (c *Client) RefundTlinxOrderWithContext(ctx context.Context, request *RefundTlinxOrderRequest) (response *RefundTlinxOrderResponse, err error) {
    if request == nil {
        request = NewRefundTlinxOrderRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RefundTlinxOrder require credential")
    }

    request.SetContext(ctx)
    
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
//  AUTHFAILURE_VERIFYERROR = "AuthFailure.VerifyError"
//  FAILEDOPERATION_BANKFAILED = "FailedOperation.BankFailed"
//  FAILEDOPERATION_SERVICEERROR = "FailedOperation.ServiceError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER_ = "MissingParameter."
func (c *Client) RegisterBehavior(request *RegisterBehaviorRequest) (response *RegisterBehaviorResponse, err error) {
    return c.RegisterBehaviorWithContext(context.Background(), request)
}

// RegisterBehavior
// 商户查询是否签约和签约行为上报
//
// 可能返回的错误码:
//  AUTHFAILURE_VERIFYERROR = "AuthFailure.VerifyError"
//  FAILEDOPERATION_BANKFAILED = "FailedOperation.BankFailed"
//  FAILEDOPERATION_SERVICEERROR = "FailedOperation.ServiceError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER_ = "MissingParameter."
func (c *Client) RegisterBehaviorWithContext(ctx context.Context, request *RegisterBehaviorRequest) (response *RegisterBehaviorResponse, err error) {
    if request == nil {
        request = NewRegisterBehaviorRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RegisterBehavior require credential")
    }

    request.SetContext(ctx)
    
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
    return c.RegisterBillWithContext(context.Background(), request)
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
func (c *Client) RegisterBillWithContext(ctx context.Context, request *RegisterBillRequest) (response *RegisterBillResponse, err error) {
    if request == nil {
        request = NewRegisterBillRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RegisterBill require credential")
    }

    request.SetContext(ctx)
    
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
    return c.RegisterBillSupportWithdrawWithContext(context.Background(), request)
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
func (c *Client) RegisterBillSupportWithdrawWithContext(ctx context.Context, request *RegisterBillSupportWithdrawRequest) (response *RegisterBillSupportWithdrawResponse, err error) {
    if request == nil {
        request = NewRegisterBillSupportWithdrawRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RegisterBillSupportWithdraw require credential")
    }

    request.SetContext(ctx)
    
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
    return c.RevResigterBillSupportWithdrawWithContext(context.Background(), request)
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
func (c *Client) RevResigterBillSupportWithdrawWithContext(ctx context.Context, request *RevResigterBillSupportWithdrawRequest) (response *RevResigterBillSupportWithdrawResponse, err error) {
    if request == nil {
        request = NewRevResigterBillSupportWithdrawRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RevResigterBillSupportWithdraw require credential")
    }

    request.SetContext(ctx)
    
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
    return c.ReviseMbrPropertyWithContext(context.Background(), request)
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
func (c *Client) ReviseMbrPropertyWithContext(ctx context.Context, request *ReviseMbrPropertyRequest) (response *ReviseMbrPropertyResponse, err error) {
    if request == nil {
        request = NewReviseMbrPropertyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ReviseMbrProperty require credential")
    }

    request.SetContext(ctx)
    
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
    return c.RevokeMemberRechargeThirdPayWithContext(context.Background(), request)
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
func (c *Client) RevokeMemberRechargeThirdPayWithContext(ctx context.Context, request *RevokeMemberRechargeThirdPayRequest) (response *RevokeMemberRechargeThirdPayResponse, err error) {
    if request == nil {
        request = NewRevokeMemberRechargeThirdPayRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RevokeMemberRechargeThirdPay require credential")
    }

    request.SetContext(ctx)
    
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
    return c.RevokeRechargeByThirdPayWithContext(context.Background(), request)
}

// RevokeRechargeByThirdPay
// 撤销会员在途充值(经第三方支付渠道)接口
//
// 可能返回的错误码:
//  FAILEDOPERATION_ACTIONINVALID = "FailedOperation.ActionInvalid"
//  FAILEDOPERATION_PABANKERROR = "FailedOperation.PABankError"
func (c *Client) RevokeRechargeByThirdPayWithContext(ctx context.Context, request *RevokeRechargeByThirdPayRequest) (response *RevokeRechargeByThirdPayResponse, err error) {
    if request == nil {
        request = NewRevokeRechargeByThirdPayRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RevokeRechargeByThirdPay require credential")
    }

    request.SetContext(ctx)
    
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
//  AUTHFAILURE_VERIFYERROR = "AuthFailure.VerifyError"
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
    return c.SyncContractDataWithContext(context.Background(), request)
}

// SyncContractData
// 对于存量的签约关系导入或者部分场景下米大师无法收到签约通知的场景，需要由调用方主动将签约状态同步至米大师
//
// 可能返回的错误码:
//  AUTHFAILURE_VERIFYERROR = "AuthFailure.VerifyError"
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
func (c *Client) SyncContractDataWithContext(ctx context.Context, request *SyncContractDataRequest) (response *SyncContractDataResponse, err error) {
    if request == nil {
        request = NewSyncContractDataRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SyncContractData require credential")
    }

    request.SetContext(ctx)
    
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
//  AUTHFAILURE_SECRETKEYNOTFOUND = "AuthFailure.SecretKeyNotFound"
//  AUTHFAILURE_VERIFYERROR = "AuthFailure.VerifyError"
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
//  FAILEDOPERATION_EXTERNALCONTRACTSTATUSINVALID = "FailedOperation.ExternalContractStatusInvalid"
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
    return c.TerminateContractWithContext(context.Background(), request)
}

// TerminateContract
// 通过此接口进行解约
//
// 可能返回的错误码:
//  AUTHFAILURE_SECRETKEYNOTFOUND = "AuthFailure.SecretKeyNotFound"
//  AUTHFAILURE_VERIFYERROR = "AuthFailure.VerifyError"
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
//  FAILEDOPERATION_EXTERNALCONTRACTSTATUSINVALID = "FailedOperation.ExternalContractStatusInvalid"
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
func (c *Client) TerminateContractWithContext(ctx context.Context, request *TerminateContractRequest) (response *TerminateContractResponse, err error) {
    if request == nil {
        request = NewTerminateContractRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("TerminateContract require credential")
    }

    request.SetContext(ctx)
    
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
//  AUTHFAILURE_SECRETKEYNOTFOUND = "AuthFailure.SecretKeyNotFound"
//  AUTHFAILURE_VERIFYERROR = "AuthFailure.VerifyError"
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
//  FAILEDOPERATION_EXTERNALCONTRACTSTATUSINVALID = "FailedOperation.ExternalContractStatusInvalid"
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
    return c.TransferSinglePayWithContext(context.Background(), request)
}

// TransferSinglePay
// 智能代发-单笔代发转账接口
//
// 可能返回的错误码:
//  AUTHFAILURE_SECRETKEYNOTFOUND = "AuthFailure.SecretKeyNotFound"
//  AUTHFAILURE_VERIFYERROR = "AuthFailure.VerifyError"
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
//  FAILEDOPERATION_EXTERNALCONTRACTSTATUSINVALID = "FailedOperation.ExternalContractStatusInvalid"
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
func (c *Client) TransferSinglePayWithContext(ctx context.Context, request *TransferSinglePayRequest) (response *TransferSinglePayResponse, err error) {
    if request == nil {
        request = NewTransferSinglePayRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("TransferSinglePay require credential")
    }

    request.SetContext(ctx)
    
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
//  AUTHFAILURE_SECRETKEYNOTFOUND = "AuthFailure.SecretKeyNotFound"
//  FAILEDOPERATION_BANKFAILED = "FailedOperation.BankFailed"
//  FAILEDOPERATION_SERVICEERROR = "FailedOperation.ServiceError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER_ = "MissingParameter."
func (c *Client) UnBindAcct(request *UnBindAcctRequest) (response *UnBindAcctResponse, err error) {
    return c.UnBindAcctWithContext(context.Background(), request)
}

// UnBindAcct
// 商户解除绑定的提现银行卡
//
// 可能返回的错误码:
//  AUTHFAILURE_SECRETKEYNOTFOUND = "AuthFailure.SecretKeyNotFound"
//  FAILEDOPERATION_BANKFAILED = "FailedOperation.BankFailed"
//  FAILEDOPERATION_SERVICEERROR = "FailedOperation.ServiceError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER_ = "MissingParameter."
func (c *Client) UnBindAcctWithContext(ctx context.Context, request *UnBindAcctRequest) (response *UnBindAcctResponse, err error) {
    if request == nil {
        request = NewUnBindAcctRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UnBindAcct require credential")
    }

    request.SetContext(ctx)
    
    response = NewUnBindAcctResponse()
    err = c.Send(request, response)
    return
}

func NewUnbindOpenBankExternalSubMerchantBankAccountRequest() (request *UnbindOpenBankExternalSubMerchantBankAccountRequest) {
    request = &UnbindOpenBankExternalSubMerchantBankAccountRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cpdp", APIVersion, "UnbindOpenBankExternalSubMerchantBankAccount")
    
    
    return
}

func NewUnbindOpenBankExternalSubMerchantBankAccountResponse() (response *UnbindOpenBankExternalSubMerchantBankAccountResponse) {
    response = &UnbindOpenBankExternalSubMerchantBankAccountResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UnbindOpenBankExternalSubMerchantBankAccount
// 云企付-子商户银行卡解绑
//
// 可能返回的错误码:
//  AUTHFAILURE_SECRETKEYNOTFOUND = "AuthFailure.SecretKeyNotFound"
//  FAILEDOPERATION_BANKFAILED = "FailedOperation.BankFailed"
//  FAILEDOPERATION_SERVICEERROR = "FailedOperation.ServiceError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER_ = "MissingParameter."
func (c *Client) UnbindOpenBankExternalSubMerchantBankAccount(request *UnbindOpenBankExternalSubMerchantBankAccountRequest) (response *UnbindOpenBankExternalSubMerchantBankAccountResponse, err error) {
    return c.UnbindOpenBankExternalSubMerchantBankAccountWithContext(context.Background(), request)
}

// UnbindOpenBankExternalSubMerchantBankAccount
// 云企付-子商户银行卡解绑
//
// 可能返回的错误码:
//  AUTHFAILURE_SECRETKEYNOTFOUND = "AuthFailure.SecretKeyNotFound"
//  FAILEDOPERATION_BANKFAILED = "FailedOperation.BankFailed"
//  FAILEDOPERATION_SERVICEERROR = "FailedOperation.ServiceError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER_ = "MissingParameter."
func (c *Client) UnbindOpenBankExternalSubMerchantBankAccountWithContext(ctx context.Context, request *UnbindOpenBankExternalSubMerchantBankAccountRequest) (response *UnbindOpenBankExternalSubMerchantBankAccountResponse, err error) {
    if request == nil {
        request = NewUnbindOpenBankExternalSubMerchantBankAccountRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UnbindOpenBankExternalSubMerchantBankAccount require credential")
    }

    request.SetContext(ctx)
    
    response = NewUnbindOpenBankExternalSubMerchantBankAccountResponse()
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
    return c.UnbindRelateAcctWithContext(context.Background(), request)
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
func (c *Client) UnbindRelateAcctWithContext(ctx context.Context, request *UnbindRelateAcctRequest) (response *UnbindRelateAcctResponse, err error) {
    if request == nil {
        request = NewUnbindRelateAcctRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UnbindRelateAcct require credential")
    }

    request.SetContext(ctx)
    
    response = NewUnbindRelateAcctResponse()
    err = c.Send(request, response)
    return
}

func NewUnifiedCloudOrderRequest() (request *UnifiedCloudOrderRequest) {
    request = &UnifiedCloudOrderRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cpdp", APIVersion, "UnifiedCloudOrder")
    
    
    return
}

func NewUnifiedCloudOrderResponse() (response *UnifiedCloudOrderResponse) {
    response = &UnifiedCloudOrderResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UnifiedCloudOrder
// 应用需要先调用本接口生成支付订单号，并将应答的PayInfo透传给聚鑫SDK，拉起客户端（包括微信公众号/微信小程序/客户端App）支付。
//
// 可能返回的错误码:
//  AUTHFAILURE_MIDAS = "AuthFailure.Midas"
//  FAILEDOPERATION_ACTION = "FailedOperation.Action"
//  FAILEDOPERATION_MIDASINTERNALERROR = "FailedOperation.MidasInternalError"
//  FAILEDOPERATION_MIDASINVALIDREQUEST = "FailedOperation.MidasInvalidRequest"
//  FAILEDOPERATION_MIDASNEEDRETRY = "FailedOperation.MidasNeedRetry"
//  FAILEDOPERATION_MIDASREGISTERUNFINISHED = "FailedOperation.MidasRegisterUnfinished"
//  FAILEDOPERATION_MIDASREPEATORDER = "FailedOperation.MidasRepeatOrder"
//  FAILEDOPERATION_MIDASRISK = "FailedOperation.MidasRisk"
//  FAILEDOPERATION_MIDASSTATUSNOTMATCH = "FailedOperation.MidasStatusNotMatch"
//  FAILEDOPERATION_MIDASUNSUPPORTEDACTION = "FailedOperation.MidasUnsupportedAction"
//  FAILEDOPERATION_MISSINGPARAMETER = "FailedOperation.MissingParameter"
//  INTERNALERROR_MIDAS = "InternalError.Midas"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  INVALIDPARAMETER_MIDAS = "InvalidParameter.Midas"
//  INVALIDPARAMETER_MIDASENVIRONMENT = "InvalidParameter.MidasEnvironment"
//  INVALIDPARAMETER_MIDASEXTERNALAPP = "InvalidParameter.MidasExternalApp"
//  INVALIDPARAMETER_MIDASFILETYPE = "InvalidParameter.MidasFileType"
//  INVALIDPARAMETER_MIDASHASH = "InvalidParameter.MidasHash"
//  INVALIDPARAMETER_MIDASSIGNID = "InvalidParameter.MidasSignId"
//  LIMITEXCEEDED_MIDASLARGEFILE = "LimitExceeded.MidasLargeFile"
//  LIMITEXCEEDED_MIDASORDER = "LimitExceeded.MidasOrder"
//  LIMITEXCEEDED_MIDASORDERCANCELED = "LimitExceeded.MidasOrderCanceled"
//  LIMITEXCEEDED_MIDASORDERCLOSED = "LimitExceeded.MidasOrderClosed"
//  LIMITEXCEEDED_MIDASORDEREXPIRED = "LimitExceeded.MidasOrderExpired"
//  LIMITEXCEEDED_MIDASORDERFAILED = "LimitExceeded.MidasOrderFailed"
//  LIMITEXCEEDED_MIDASORDERPARTIALSUCCESS = "LimitExceeded.MidasOrderPartialSuccess"
//  LIMITEXCEEDED_MIDASORDERPROCESSING = "LimitExceeded.MidasOrderProcessing"
//  LIMITEXCEEDED_MIDASORDERSUCCESS = "LimitExceeded.MidasOrderSuccess"
//  REQUESTLIMITEXCEEDED_MIDAS = "RequestLimitExceeded.Midas"
//  REQUESTLIMITEXCEEDED_MIDASINVALIDREQUEST = "RequestLimitExceeded.MidasInvalidRequest"
//  RESOURCEINUSE_MIDAS = "ResourceInUse.Midas"
//  RESOURCENOTFOUND_ACCOUNT = "ResourceNotFound.Account"
//  RESOURCENOTFOUND_MIDASEXTERNALAPP = "ResourceNotFound.MidasExternalApp"
//  RESOURCENOTFOUND_MIDASEXTERNALORDER = "ResourceNotFound.MidasExternalOrder"
//  RESOURCENOTFOUND_MIDASORDER = "ResourceNotFound.MidasOrder"
//  RESOURCENOTFOUND_MIDASSIGN = "ResourceNotFound.MidasSign"
//  RESOURCEUNAVAILABLE_MIDASBALANCE = "ResourceUnavailable.MidasBalance"
//  RESOURCEUNAVAILABLE_MIDASDAY = "ResourceUnavailable.MidasDay"
//  RESOURCEUNAVAILABLE_MIDASFROZENAMOUNT = "ResourceUnavailable.MidasFrozenAmount"
//  RESOURCEUNAVAILABLE_MIDASMERCHANTBALANCE = "ResourceUnavailable.MidasMerchantBalance"
//  RESOURCEUNAVAILABLE_MIDASORDER = "ResourceUnavailable.MidasOrder"
//  RESOURCEUNAVAILABLE_MIDASUSERBALANCE = "ResourceUnavailable.MidasUserBalance"
//  RESOURCEUNAVAILABLE_MIDASWALLET = "ResourceUnavailable.MidasWallet"
//  UNAUTHORIZEDOPERATION_MIDAS = "UnauthorizedOperation.Midas"
func (c *Client) UnifiedCloudOrder(request *UnifiedCloudOrderRequest) (response *UnifiedCloudOrderResponse, err error) {
    return c.UnifiedCloudOrderWithContext(context.Background(), request)
}

// UnifiedCloudOrder
// 应用需要先调用本接口生成支付订单号，并将应答的PayInfo透传给聚鑫SDK，拉起客户端（包括微信公众号/微信小程序/客户端App）支付。
//
// 可能返回的错误码:
//  AUTHFAILURE_MIDAS = "AuthFailure.Midas"
//  FAILEDOPERATION_ACTION = "FailedOperation.Action"
//  FAILEDOPERATION_MIDASINTERNALERROR = "FailedOperation.MidasInternalError"
//  FAILEDOPERATION_MIDASINVALIDREQUEST = "FailedOperation.MidasInvalidRequest"
//  FAILEDOPERATION_MIDASNEEDRETRY = "FailedOperation.MidasNeedRetry"
//  FAILEDOPERATION_MIDASREGISTERUNFINISHED = "FailedOperation.MidasRegisterUnfinished"
//  FAILEDOPERATION_MIDASREPEATORDER = "FailedOperation.MidasRepeatOrder"
//  FAILEDOPERATION_MIDASRISK = "FailedOperation.MidasRisk"
//  FAILEDOPERATION_MIDASSTATUSNOTMATCH = "FailedOperation.MidasStatusNotMatch"
//  FAILEDOPERATION_MIDASUNSUPPORTEDACTION = "FailedOperation.MidasUnsupportedAction"
//  FAILEDOPERATION_MISSINGPARAMETER = "FailedOperation.MissingParameter"
//  INTERNALERROR_MIDAS = "InternalError.Midas"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  INVALIDPARAMETER_MIDAS = "InvalidParameter.Midas"
//  INVALIDPARAMETER_MIDASENVIRONMENT = "InvalidParameter.MidasEnvironment"
//  INVALIDPARAMETER_MIDASEXTERNALAPP = "InvalidParameter.MidasExternalApp"
//  INVALIDPARAMETER_MIDASFILETYPE = "InvalidParameter.MidasFileType"
//  INVALIDPARAMETER_MIDASHASH = "InvalidParameter.MidasHash"
//  INVALIDPARAMETER_MIDASSIGNID = "InvalidParameter.MidasSignId"
//  LIMITEXCEEDED_MIDASLARGEFILE = "LimitExceeded.MidasLargeFile"
//  LIMITEXCEEDED_MIDASORDER = "LimitExceeded.MidasOrder"
//  LIMITEXCEEDED_MIDASORDERCANCELED = "LimitExceeded.MidasOrderCanceled"
//  LIMITEXCEEDED_MIDASORDERCLOSED = "LimitExceeded.MidasOrderClosed"
//  LIMITEXCEEDED_MIDASORDEREXPIRED = "LimitExceeded.MidasOrderExpired"
//  LIMITEXCEEDED_MIDASORDERFAILED = "LimitExceeded.MidasOrderFailed"
//  LIMITEXCEEDED_MIDASORDERPARTIALSUCCESS = "LimitExceeded.MidasOrderPartialSuccess"
//  LIMITEXCEEDED_MIDASORDERPROCESSING = "LimitExceeded.MidasOrderProcessing"
//  LIMITEXCEEDED_MIDASORDERSUCCESS = "LimitExceeded.MidasOrderSuccess"
//  REQUESTLIMITEXCEEDED_MIDAS = "RequestLimitExceeded.Midas"
//  REQUESTLIMITEXCEEDED_MIDASINVALIDREQUEST = "RequestLimitExceeded.MidasInvalidRequest"
//  RESOURCEINUSE_MIDAS = "ResourceInUse.Midas"
//  RESOURCENOTFOUND_ACCOUNT = "ResourceNotFound.Account"
//  RESOURCENOTFOUND_MIDASEXTERNALAPP = "ResourceNotFound.MidasExternalApp"
//  RESOURCENOTFOUND_MIDASEXTERNALORDER = "ResourceNotFound.MidasExternalOrder"
//  RESOURCENOTFOUND_MIDASORDER = "ResourceNotFound.MidasOrder"
//  RESOURCENOTFOUND_MIDASSIGN = "ResourceNotFound.MidasSign"
//  RESOURCEUNAVAILABLE_MIDASBALANCE = "ResourceUnavailable.MidasBalance"
//  RESOURCEUNAVAILABLE_MIDASDAY = "ResourceUnavailable.MidasDay"
//  RESOURCEUNAVAILABLE_MIDASFROZENAMOUNT = "ResourceUnavailable.MidasFrozenAmount"
//  RESOURCEUNAVAILABLE_MIDASMERCHANTBALANCE = "ResourceUnavailable.MidasMerchantBalance"
//  RESOURCEUNAVAILABLE_MIDASORDER = "ResourceUnavailable.MidasOrder"
//  RESOURCEUNAVAILABLE_MIDASUSERBALANCE = "ResourceUnavailable.MidasUserBalance"
//  RESOURCEUNAVAILABLE_MIDASWALLET = "ResourceUnavailable.MidasWallet"
//  UNAUTHORIZEDOPERATION_MIDAS = "UnauthorizedOperation.Midas"
func (c *Client) UnifiedCloudOrderWithContext(ctx context.Context, request *UnifiedCloudOrderRequest) (response *UnifiedCloudOrderResponse, err error) {
    if request == nil {
        request = NewUnifiedCloudOrderRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UnifiedCloudOrder require credential")
    }

    request.SetContext(ctx)
    
    response = NewUnifiedCloudOrderResponse()
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
//  AUTHFAILURE_FORBIDDEN = "AuthFailure.Forbidden"
//  AUTHFAILURE_SECRETKEYNOTFOUND = "AuthFailure.SecretKeyNotFound"
//  AUTHFAILURE_VERIFYERROR = "AuthFailure.VerifyError"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_APPDENY = "FailedOperation.AppDeny"
//  FAILEDOPERATION_CHANNELDENY = "FailedOperation.ChannelDeny"
//  FAILEDOPERATION_OCCOMPLETEDORDER = "FailedOperation.OcCompletedOrder"
//  FAILEDOPERATION_OCREPEATORDER = "FailedOperation.OcRepeatOrder"
//  FAILEDOPERATION_PARENTAPPIDERROR = "FailedOperation.ParentAppIdError"
//  FAILEDOPERATION_PORTALERROR = "FailedOperation.PortalError"
//  FAILEDOPERATION_WECHATERROR = "FailedOperation.WechatError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_BACKENDCGIERROR = "InvalidParameter.BackendCgiError"
func (c *Client) UnifiedOrder(request *UnifiedOrderRequest) (response *UnifiedOrderResponse, err error) {
    return c.UnifiedOrderWithContext(context.Background(), request)
}

// UnifiedOrder
// 应用需要先调用本接口生成支付订单号，并将应答的PayInfo透传给聚鑫SDK，拉起客户端（包括微信公众号/微信小程序/客户端App）支付。
//
// 可能返回的错误码:
//  AUTHFAILURE_FORBIDDEN = "AuthFailure.Forbidden"
//  AUTHFAILURE_SECRETKEYNOTFOUND = "AuthFailure.SecretKeyNotFound"
//  AUTHFAILURE_VERIFYERROR = "AuthFailure.VerifyError"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_APPDENY = "FailedOperation.AppDeny"
//  FAILEDOPERATION_CHANNELDENY = "FailedOperation.ChannelDeny"
//  FAILEDOPERATION_OCCOMPLETEDORDER = "FailedOperation.OcCompletedOrder"
//  FAILEDOPERATION_OCREPEATORDER = "FailedOperation.OcRepeatOrder"
//  FAILEDOPERATION_PARENTAPPIDERROR = "FailedOperation.ParentAppIdError"
//  FAILEDOPERATION_PORTALERROR = "FailedOperation.PortalError"
//  FAILEDOPERATION_WECHATERROR = "FailedOperation.WechatError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_BACKENDCGIERROR = "InvalidParameter.BackendCgiError"
func (c *Client) UnifiedOrderWithContext(ctx context.Context, request *UnifiedOrderRequest) (response *UnifiedOrderResponse, err error) {
    if request == nil {
        request = NewUnifiedOrderRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UnifiedOrder require credential")
    }

    request.SetContext(ctx)
    
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
//  AUTHFAILURE_FORBIDDEN = "AuthFailure.Forbidden"
//  AUTHFAILURE_SECRETKEYNOTFOUND = "AuthFailure.SecretKeyNotFound"
//  AUTHFAILURE_VERIFYERROR = "AuthFailure.VerifyError"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_APPDENY = "FailedOperation.AppDeny"
//  FAILEDOPERATION_CHANNELDENY = "FailedOperation.ChannelDeny"
//  FAILEDOPERATION_OCCOMPLETEDORDER = "FailedOperation.OcCompletedOrder"
//  FAILEDOPERATION_OCREPEATORDER = "FailedOperation.OcRepeatOrder"
//  FAILEDOPERATION_PARENTAPPIDERROR = "FailedOperation.ParentAppIdError"
//  FAILEDOPERATION_PORTALERROR = "FailedOperation.PortalError"
//  FAILEDOPERATION_WECHATERROR = "FailedOperation.WechatError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_BACKENDCGIERROR = "InvalidParameter.BackendCgiError"
func (c *Client) UnifiedTlinxOrder(request *UnifiedTlinxOrderRequest) (response *UnifiedTlinxOrderResponse, err error) {
    return c.UnifiedTlinxOrderWithContext(context.Background(), request)
}

// UnifiedTlinxOrder
// 云支付-统一下单接口
//
// 可能返回的错误码:
//  AUTHFAILURE_FORBIDDEN = "AuthFailure.Forbidden"
//  AUTHFAILURE_SECRETKEYNOTFOUND = "AuthFailure.SecretKeyNotFound"
//  AUTHFAILURE_VERIFYERROR = "AuthFailure.VerifyError"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_APPDENY = "FailedOperation.AppDeny"
//  FAILEDOPERATION_CHANNELDENY = "FailedOperation.ChannelDeny"
//  FAILEDOPERATION_OCCOMPLETEDORDER = "FailedOperation.OcCompletedOrder"
//  FAILEDOPERATION_OCREPEATORDER = "FailedOperation.OcRepeatOrder"
//  FAILEDOPERATION_PARENTAPPIDERROR = "FailedOperation.ParentAppIdError"
//  FAILEDOPERATION_PORTALERROR = "FailedOperation.PortalError"
//  FAILEDOPERATION_WECHATERROR = "FailedOperation.WechatError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_BACKENDCGIERROR = "InvalidParameter.BackendCgiError"
func (c *Client) UnifiedTlinxOrderWithContext(ctx context.Context, request *UnifiedTlinxOrderRequest) (response *UnifiedTlinxOrderResponse, err error) {
    if request == nil {
        request = NewUnifiedTlinxOrderRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UnifiedTlinxOrder require credential")
    }

    request.SetContext(ctx)
    
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
//  AUTHFAILURE_FORBIDDEN = "AuthFailure.Forbidden"
//  AUTHFAILURE_SECRETKEYNOTFOUND = "AuthFailure.SecretKeyNotFound"
//  AUTHFAILURE_VERIFYERROR = "AuthFailure.VerifyError"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_APPDENY = "FailedOperation.AppDeny"
//  FAILEDOPERATION_CHANNELDENY = "FailedOperation.ChannelDeny"
//  FAILEDOPERATION_OCCOMPLETEDORDER = "FailedOperation.OcCompletedOrder"
//  FAILEDOPERATION_OCREPEATORDER = "FailedOperation.OcRepeatOrder"
//  FAILEDOPERATION_PARENTAPPIDERROR = "FailedOperation.ParentAppIdError"
//  FAILEDOPERATION_PORTALERROR = "FailedOperation.PortalError"
//  FAILEDOPERATION_WECHATERROR = "FailedOperation.WechatError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_BACKENDCGIERROR = "InvalidParameter.BackendCgiError"
func (c *Client) UploadExternalAnchorInfo(request *UploadExternalAnchorInfoRequest) (response *UploadExternalAnchorInfoResponse, err error) {
    return c.UploadExternalAnchorInfoWithContext(context.Background(), request)
}

// UploadExternalAnchorInfo
// 灵云-上传主播信息
//
// 可能返回的错误码:
//  AUTHFAILURE_FORBIDDEN = "AuthFailure.Forbidden"
//  AUTHFAILURE_SECRETKEYNOTFOUND = "AuthFailure.SecretKeyNotFound"
//  AUTHFAILURE_VERIFYERROR = "AuthFailure.VerifyError"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_APPDENY = "FailedOperation.AppDeny"
//  FAILEDOPERATION_CHANNELDENY = "FailedOperation.ChannelDeny"
//  FAILEDOPERATION_OCCOMPLETEDORDER = "FailedOperation.OcCompletedOrder"
//  FAILEDOPERATION_OCREPEATORDER = "FailedOperation.OcRepeatOrder"
//  FAILEDOPERATION_PARENTAPPIDERROR = "FailedOperation.ParentAppIdError"
//  FAILEDOPERATION_PORTALERROR = "FailedOperation.PortalError"
//  FAILEDOPERATION_WECHATERROR = "FailedOperation.WechatError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_BACKENDCGIERROR = "InvalidParameter.BackendCgiError"
func (c *Client) UploadExternalAnchorInfoWithContext(ctx context.Context, request *UploadExternalAnchorInfoRequest) (response *UploadExternalAnchorInfoResponse, err error) {
    if request == nil {
        request = NewUploadExternalAnchorInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UploadExternalAnchorInfo require credential")
    }

    request.SetContext(ctx)
    
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
    return c.UploadFileWithContext(context.Background(), request)
}

// UploadFile
// 直播平台-文件上传
//
// 可能返回的错误码:
//  FAILEDOPERATION_ACTION = "FailedOperation.Action"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  MISSINGPARAMETER_ACTION = "MissingParameter.Action"
//  MISSINGPARAMETER_APPID = "MissingParameter.AppId"
func (c *Client) UploadFileWithContext(ctx context.Context, request *UploadFileRequest) (response *UploadFileResponse, err error) {
    if request == nil {
        request = NewUploadFileRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UploadFile require credential")
    }

    request.SetContext(ctx)
    
    response = NewUploadFileResponse()
    err = c.Send(request, response)
    return
}

func NewUploadOpenBankSubMerchantCredentialRequest() (request *UploadOpenBankSubMerchantCredentialRequest) {
    request = &UploadOpenBankSubMerchantCredentialRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cpdp", APIVersion, "UploadOpenBankSubMerchantCredential")
    
    
    return
}

func NewUploadOpenBankSubMerchantCredentialResponse() (response *UploadOpenBankSubMerchantCredentialResponse) {
    response = &UploadOpenBankSubMerchantCredentialResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UploadOpenBankSubMerchantCredential
// 云企付-子商户资质文件上传
//
// 可能返回的错误码:
//  FAILEDOPERATION_ACTION = "FailedOperation.Action"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  MISSINGPARAMETER_ACTION = "MissingParameter.Action"
//  MISSINGPARAMETER_APPID = "MissingParameter.AppId"
func (c *Client) UploadOpenBankSubMerchantCredential(request *UploadOpenBankSubMerchantCredentialRequest) (response *UploadOpenBankSubMerchantCredentialResponse, err error) {
    return c.UploadOpenBankSubMerchantCredentialWithContext(context.Background(), request)
}

// UploadOpenBankSubMerchantCredential
// 云企付-子商户资质文件上传
//
// 可能返回的错误码:
//  FAILEDOPERATION_ACTION = "FailedOperation.Action"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  MISSINGPARAMETER_ACTION = "MissingParameter.Action"
//  MISSINGPARAMETER_APPID = "MissingParameter.AppId"
func (c *Client) UploadOpenBankSubMerchantCredentialWithContext(ctx context.Context, request *UploadOpenBankSubMerchantCredentialRequest) (response *UploadOpenBankSubMerchantCredentialResponse, err error) {
    if request == nil {
        request = NewUploadOpenBankSubMerchantCredentialRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UploadOpenBankSubMerchantCredential require credential")
    }

    request.SetContext(ctx)
    
    response = NewUploadOpenBankSubMerchantCredentialResponse()
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
    return c.UploadOrgFileWithContext(context.Background(), request)
}

// UploadOrgFile
// 云支付-上传机构文件接口
//
// 可能返回的错误码:
//  FAILEDOPERATION_ACTION = "FailedOperation.Action"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  MISSINGPARAMETER_ACTION = "MissingParameter.Action"
//  MISSINGPARAMETER_APPID = "MissingParameter.AppId"
func (c *Client) UploadOrgFileWithContext(ctx context.Context, request *UploadOrgFileRequest) (response *UploadOrgFileResponse, err error) {
    if request == nil {
        request = NewUploadOrgFileRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UploadOrgFile require credential")
    }

    request.SetContext(ctx)
    
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
    return c.UploadTaxListWithContext(context.Background(), request)
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
func (c *Client) UploadTaxListWithContext(ctx context.Context, request *UploadTaxListRequest) (response *UploadTaxListResponse, err error) {
    if request == nil {
        request = NewUploadTaxListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UploadTaxList require credential")
    }

    request.SetContext(ctx)
    
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
    return c.UploadTaxPaymentWithContext(context.Background(), request)
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
func (c *Client) UploadTaxPaymentWithContext(ctx context.Context, request *UploadTaxPaymentRequest) (response *UploadTaxPaymentResponse, err error) {
    if request == nil {
        request = NewUploadTaxPaymentRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UploadTaxPayment require credential")
    }

    request.SetContext(ctx)
    
    response = NewUploadTaxPaymentResponse()
    err = c.Send(request, response)
    return
}

func NewVerifyOpenBankAccountRequest() (request *VerifyOpenBankAccountRequest) {
    request = &VerifyOpenBankAccountRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cpdp", APIVersion, "VerifyOpenBankAccount")
    
    
    return
}

func NewVerifyOpenBankAccountResponse() (response *VerifyOpenBankAccountResponse) {
    response = &VerifyOpenBankAccountResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// VerifyOpenBankAccount
// 云企付-子商户银行卡打款验证，在接入TENPAY渠道EBANK_PAYMENT付款时，若客户期望接入担保支付，需在接入前先完成，收款商户绑定的银行卡进行打款验证。验证成功后，才可以调用CreateOpenBankPaymentOrder接口进行担保支付下单。
//
// 可能返回的错误码:
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
func (c *Client) VerifyOpenBankAccount(request *VerifyOpenBankAccountRequest) (response *VerifyOpenBankAccountResponse, err error) {
    return c.VerifyOpenBankAccountWithContext(context.Background(), request)
}

// VerifyOpenBankAccount
// 云企付-子商户银行卡打款验证，在接入TENPAY渠道EBANK_PAYMENT付款时，若客户期望接入担保支付，需在接入前先完成，收款商户绑定的银行卡进行打款验证。验证成功后，才可以调用CreateOpenBankPaymentOrder接口进行担保支付下单。
//
// 可能返回的错误码:
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
func (c *Client) VerifyOpenBankAccountWithContext(ctx context.Context, request *VerifyOpenBankAccountRequest) (response *VerifyOpenBankAccountResponse, err error) {
    if request == nil {
        request = NewVerifyOpenBankAccountRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("VerifyOpenBankAccount require credential")
    }

    request.SetContext(ctx)
    
    response = NewVerifyOpenBankAccountResponse()
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
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
func (c *Client) ViewContract(request *ViewContractRequest) (response *ViewContractResponse, err error) {
    return c.ViewContractWithContext(context.Background(), request)
}

// ViewContract
// 云支付-查询合同明细接口
//
// 可能返回的错误码:
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
func (c *Client) ViewContractWithContext(ctx context.Context, request *ViewContractRequest) (response *ViewContractResponse, err error) {
    if request == nil {
        request = NewViewContractRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ViewContract require credential")
    }

    request.SetContext(ctx)
    
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
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
func (c *Client) ViewMerchant(request *ViewMerchantRequest) (response *ViewMerchantResponse, err error) {
    return c.ViewMerchantWithContext(context.Background(), request)
}

// ViewMerchant
// 云支付-查询商户明细接口
//
// 可能返回的错误码:
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
func (c *Client) ViewMerchantWithContext(ctx context.Context, request *ViewMerchantRequest) (response *ViewMerchantResponse, err error) {
    if request == nil {
        request = NewViewMerchantRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ViewMerchant require credential")
    }

    request.SetContext(ctx)
    
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
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
func (c *Client) ViewShop(request *ViewShopRequest) (response *ViewShopResponse, err error) {
    return c.ViewShopWithContext(context.Background(), request)
}

// ViewShop
// 云支付-查询门店明细接口
//
// 可能返回的错误码:
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
func (c *Client) ViewShopWithContext(ctx context.Context, request *ViewShopRequest) (response *ViewShopResponse, err error) {
    if request == nil {
        request = NewViewShopRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ViewShop require credential")
    }

    request.SetContext(ctx)
    
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
    return c.WithdrawCashMembershipWithContext(context.Background(), request)
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
func (c *Client) WithdrawCashMembershipWithContext(ctx context.Context, request *WithdrawCashMembershipRequest) (response *WithdrawCashMembershipResponse, err error) {
    if request == nil {
        request = NewWithdrawCashMembershipRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("WithdrawCashMembership require credential")
    }

    request.SetContext(ctx)
    
    response = NewWithdrawCashMembershipResponse()
    err = c.Send(request, response)
    return
}
