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

package v20220928

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2022-09-28"

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


func NewAllocateCustomerCreditRequest() (request *AllocateCustomerCreditRequest) {
    request = &AllocateCustomerCreditRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("intlpartnersmgt", APIVersion, "AllocateCustomerCredit")
    
    
    return
}

func NewAllocateCustomerCreditResponse() (response *AllocateCustomerCreditResponse) {
    response = &AllocateCustomerCreditResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AllocateCustomerCredit
// 合作伙伴可以为关联客户设置信用额度，包括调高额度、降低额度、设置额度为0
//
// 1、信用额度长期有效，不会定期清0；
//
// 2、可用信用额度为0，会触发客户停服，请谨慎操作；
//
// 3、如需限制客户新购，但不影响已购产品使用，可与渠道经理申请客户欠费不停服特权后，设置可用信用额度为0；
//
// 4、设置的额度，为当前可用信用额度的增量，最大不能超过合作伙伴可分配的剩余额度，设置负数代表回收额度，可用信用额度最低设置为0。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) AllocateCustomerCredit(request *AllocateCustomerCreditRequest) (response *AllocateCustomerCreditResponse, err error) {
    return c.AllocateCustomerCreditWithContext(context.Background(), request)
}

// AllocateCustomerCredit
// 合作伙伴可以为关联客户设置信用额度，包括调高额度、降低额度、设置额度为0
//
// 1、信用额度长期有效，不会定期清0；
//
// 2、可用信用额度为0，会触发客户停服，请谨慎操作；
//
// 3、如需限制客户新购，但不影响已购产品使用，可与渠道经理申请客户欠费不停服特权后，设置可用信用额度为0；
//
// 4、设置的额度，为当前可用信用额度的增量，最大不能超过合作伙伴可分配的剩余额度，设置负数代表回收额度，可用信用额度最低设置为0。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) AllocateCustomerCreditWithContext(ctx context.Context, request *AllocateCustomerCreditRequest) (response *AllocateCustomerCreditResponse, err error) {
    if request == nil {
        request = NewAllocateCustomerCreditRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AllocateCustomerCredit require credential")
    }

    request.SetContext(ctx)
    
    response = NewAllocateCustomerCreditResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAccountRequest() (request *CreateAccountRequest) {
    request = &CreateAccountRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("intlpartnersmgt", APIVersion, "CreateAccount")
    
    
    return
}

func NewCreateAccountResponse() (response *CreateAccountResponse) {
    response = &CreateAccountResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateAccount
// 在合作伙伴平台，创建腾讯云账号，子客户注册完成后，自动与合作伙伴账号绑定。
//
// 
//
// 注意事项：<br>
//
// 1、创建腾讯云账号，输入的邮箱、手机号，需要合作伙伴做有效性验证。<br>
//
// 2、客户首次登录需要补充个人信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_MAILISREGISTERED = "FailedOperation.MailIsRegistered"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ACCOUNTTYPECONTENTINCORRECT = "InvalidParameter.AccountTypeContentIncorrect"
//  INVALIDPARAMETER_AREACONTENTINCORRECT = "InvalidParameter.AreaContentIncorrect"
//  INVALIDPARAMETER_AREAFORMATINCORRECT = "InvalidParameter.AreaFormatIncorrect"
//  INVALIDPARAMETER_CONFIRMPASSWORDCONTENTINCORRECT = "InvalidParameter.ConfirmPasswordContentIncorrect"
//  INVALIDPARAMETER_COUNTRYCODECONTENTINCORRECT = "InvalidParameter.CountryCodeContentIncorrect"
//  INVALIDPARAMETER_COUNTRYCODEFORMATINCORRECT = "InvalidParameter.CountryCodeFormatIncorrect"
//  INVALIDPARAMETER_MAILFORMATINCORRECT = "InvalidParameter.MailFormatIncorrect"
//  INVALIDPARAMETER_PASSWORDCONTENTINCORRECT = "InvalidParameter.PasswordContentIncorrect"
//  INVALIDPARAMETER_PASSWORDFORMATINCORRECT = "InvalidParameter.PasswordFormatIncorrect"
//  INVALIDPARAMETER_PHONENUMFORMATINCORRECT = "InvalidParameter.PhoneNumFormatIncorrect"
//  INVALIDPARAMETERVALUE_ACCOUNTTYPEEMPTY = "InvalidParameterValue.AccountTypeEmpty"
//  INVALIDPARAMETERVALUE_AREAEMPTY = "InvalidParameterValue.AreaEmpty"
//  INVALIDPARAMETERVALUE_COUNTRYCODEEMPTY = "InvalidParameterValue.CountryCodeEmpty"
//  INVALIDPARAMETERVALUE_MAILEMPTY = "InvalidParameterValue.MailEmpty"
//  INVALIDPARAMETERVALUE_PASSWORDEMPTY = "InvalidParameterValue.PasswordEmpty"
//  INVALIDPARAMETERVALUE_PHONENUMEMPTY = "InvalidParameterValue.PhoneNumEmpty"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateAccount(request *CreateAccountRequest) (response *CreateAccountResponse, err error) {
    return c.CreateAccountWithContext(context.Background(), request)
}

// CreateAccount
// 在合作伙伴平台，创建腾讯云账号，子客户注册完成后，自动与合作伙伴账号绑定。
//
// 
//
// 注意事项：<br>
//
// 1、创建腾讯云账号，输入的邮箱、手机号，需要合作伙伴做有效性验证。<br>
//
// 2、客户首次登录需要补充个人信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_MAILISREGISTERED = "FailedOperation.MailIsRegistered"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ACCOUNTTYPECONTENTINCORRECT = "InvalidParameter.AccountTypeContentIncorrect"
//  INVALIDPARAMETER_AREACONTENTINCORRECT = "InvalidParameter.AreaContentIncorrect"
//  INVALIDPARAMETER_AREAFORMATINCORRECT = "InvalidParameter.AreaFormatIncorrect"
//  INVALIDPARAMETER_CONFIRMPASSWORDCONTENTINCORRECT = "InvalidParameter.ConfirmPasswordContentIncorrect"
//  INVALIDPARAMETER_COUNTRYCODECONTENTINCORRECT = "InvalidParameter.CountryCodeContentIncorrect"
//  INVALIDPARAMETER_COUNTRYCODEFORMATINCORRECT = "InvalidParameter.CountryCodeFormatIncorrect"
//  INVALIDPARAMETER_MAILFORMATINCORRECT = "InvalidParameter.MailFormatIncorrect"
//  INVALIDPARAMETER_PASSWORDCONTENTINCORRECT = "InvalidParameter.PasswordContentIncorrect"
//  INVALIDPARAMETER_PASSWORDFORMATINCORRECT = "InvalidParameter.PasswordFormatIncorrect"
//  INVALIDPARAMETER_PHONENUMFORMATINCORRECT = "InvalidParameter.PhoneNumFormatIncorrect"
//  INVALIDPARAMETERVALUE_ACCOUNTTYPEEMPTY = "InvalidParameterValue.AccountTypeEmpty"
//  INVALIDPARAMETERVALUE_AREAEMPTY = "InvalidParameterValue.AreaEmpty"
//  INVALIDPARAMETERVALUE_COUNTRYCODEEMPTY = "InvalidParameterValue.CountryCodeEmpty"
//  INVALIDPARAMETERVALUE_MAILEMPTY = "InvalidParameterValue.MailEmpty"
//  INVALIDPARAMETERVALUE_PASSWORDEMPTY = "InvalidParameterValue.PasswordEmpty"
//  INVALIDPARAMETERVALUE_PHONENUMEMPTY = "InvalidParameterValue.PhoneNumEmpty"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateAccountWithContext(ctx context.Context, request *CreateAccountRequest) (response *CreateAccountResponse, err error) {
    if request == nil {
        request = NewCreateAccountRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAccount require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAccountResponse()
    err = c.Send(request, response)
    return
}

func NewGetCountryCodesRequest() (request *GetCountryCodesRequest) {
    request = &GetCountryCodesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("intlpartnersmgt", APIVersion, "GetCountryCodes")
    
    
    return
}

func NewGetCountryCodesResponse() (response *GetCountryCodesResponse) {
    response = &GetCountryCodesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetCountryCodes
// 获取国家和地区的代码
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) GetCountryCodes(request *GetCountryCodesRequest) (response *GetCountryCodesResponse, err error) {
    return c.GetCountryCodesWithContext(context.Background(), request)
}

// GetCountryCodes
// 获取国家和地区的代码
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) GetCountryCodesWithContext(ctx context.Context, request *GetCountryCodesRequest) (response *GetCountryCodesResponse, err error) {
    if request == nil {
        request = NewGetCountryCodesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetCountryCodes require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetCountryCodesResponse()
    err = c.Send(request, response)
    return
}

func NewQueryCreditAllocationHistoryRequest() (request *QueryCreditAllocationHistoryRequest) {
    request = &QueryCreditAllocationHistoryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("intlpartnersmgt", APIVersion, "QueryCreditAllocationHistory")
    
    
    return
}

func NewQueryCreditAllocationHistoryResponse() (response *QueryCreditAllocationHistoryResponse) {
    response = &QueryCreditAllocationHistoryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// QueryCreditAllocationHistory
// 查询单个客户的全部历史分配记录
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) QueryCreditAllocationHistory(request *QueryCreditAllocationHistoryRequest) (response *QueryCreditAllocationHistoryResponse, err error) {
    return c.QueryCreditAllocationHistoryWithContext(context.Background(), request)
}

// QueryCreditAllocationHistory
// 查询单个客户的全部历史分配记录
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) QueryCreditAllocationHistoryWithContext(ctx context.Context, request *QueryCreditAllocationHistoryRequest) (response *QueryCreditAllocationHistoryResponse, err error) {
    if request == nil {
        request = NewQueryCreditAllocationHistoryRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryCreditAllocationHistory require credential")
    }

    request.SetContext(ctx)
    
    response = NewQueryCreditAllocationHistoryResponse()
    err = c.Send(request, response)
    return
}

func NewQueryCreditByUinListRequest() (request *QueryCreditByUinListRequest) {
    request = &QueryCreditByUinListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("intlpartnersmgt", APIVersion, "QueryCreditByUinList")
    
    
    return
}

func NewQueryCreditByUinListResponse() (response *QueryCreditByUinListResponse) {
    response = &QueryCreditByUinListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// QueryCreditByUinList
// 查询用户列表信用
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) QueryCreditByUinList(request *QueryCreditByUinListRequest) (response *QueryCreditByUinListResponse, err error) {
    return c.QueryCreditByUinListWithContext(context.Background(), request)
}

// QueryCreditByUinList
// 查询用户列表信用
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) QueryCreditByUinListWithContext(ctx context.Context, request *QueryCreditByUinListRequest) (response *QueryCreditByUinListResponse, err error) {
    if request == nil {
        request = NewQueryCreditByUinListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryCreditByUinList require credential")
    }

    request.SetContext(ctx)
    
    response = NewQueryCreditByUinListResponse()
    err = c.Send(request, response)
    return
}

func NewQueryCustomersCreditRequest() (request *QueryCustomersCreditRequest) {
    request = &QueryCustomersCreditRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("intlpartnersmgt", APIVersion, "QueryCustomersCredit")
    
    
    return
}

func NewQueryCustomersCreditResponse() (response *QueryCustomersCreditResponse) {
    response = &QueryCustomersCreditResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// QueryCustomersCredit
// 合作伙伴可以查询关联客户的信用额度，以及客户的基础信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) QueryCustomersCredit(request *QueryCustomersCreditRequest) (response *QueryCustomersCreditResponse, err error) {
    return c.QueryCustomersCreditWithContext(context.Background(), request)
}

// QueryCustomersCredit
// 合作伙伴可以查询关联客户的信用额度，以及客户的基础信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) QueryCustomersCreditWithContext(ctx context.Context, request *QueryCustomersCreditRequest) (response *QueryCustomersCreditResponse, err error) {
    if request == nil {
        request = NewQueryCustomersCreditRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryCustomersCredit require credential")
    }

    request.SetContext(ctx)
    
    response = NewQueryCustomersCreditResponse()
    err = c.Send(request, response)
    return
}

func NewQueryDirectCustomersCreditRequest() (request *QueryDirectCustomersCreditRequest) {
    request = &QueryDirectCustomersCreditRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("intlpartnersmgt", APIVersion, "QueryDirectCustomersCredit")
    
    
    return
}

func NewQueryDirectCustomersCreditResponse() (response *QueryDirectCustomersCreditResponse) {
    response = &QueryDirectCustomersCreditResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// QueryDirectCustomersCredit
// 查询直接子客信用
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) QueryDirectCustomersCredit(request *QueryDirectCustomersCreditRequest) (response *QueryDirectCustomersCreditResponse, err error) {
    return c.QueryDirectCustomersCreditWithContext(context.Background(), request)
}

// QueryDirectCustomersCredit
// 查询直接子客信用
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) QueryDirectCustomersCreditWithContext(ctx context.Context, request *QueryDirectCustomersCreditRequest) (response *QueryDirectCustomersCreditResponse, err error) {
    if request == nil {
        request = NewQueryDirectCustomersCreditRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryDirectCustomersCredit require credential")
    }

    request.SetContext(ctx)
    
    response = NewQueryDirectCustomersCreditResponse()
    err = c.Send(request, response)
    return
}

func NewQueryPartnerCreditRequest() (request *QueryPartnerCreditRequest) {
    request = &QueryPartnerCreditRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("intlpartnersmgt", APIVersion, "QueryPartnerCredit")
    
    
    return
}

func NewQueryPartnerCreditResponse() (response *QueryPartnerCreditResponse) {
    response = &QueryPartnerCreditResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// QueryPartnerCredit
// 查询合作伙伴自己的总信用额度、可用信用额度、已使用信用额度，单位为美元
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) QueryPartnerCredit(request *QueryPartnerCreditRequest) (response *QueryPartnerCreditResponse, err error) {
    return c.QueryPartnerCreditWithContext(context.Background(), request)
}

// QueryPartnerCredit
// 查询合作伙伴自己的总信用额度、可用信用额度、已使用信用额度，单位为美元
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) QueryPartnerCreditWithContext(ctx context.Context, request *QueryPartnerCreditRequest) (response *QueryPartnerCreditResponse, err error) {
    if request == nil {
        request = NewQueryPartnerCreditRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryPartnerCredit require credential")
    }

    request.SetContext(ctx)
    
    response = NewQueryPartnerCreditResponse()
    err = c.Send(request, response)
    return
}

func NewQueryVoucherAmountByUinRequest() (request *QueryVoucherAmountByUinRequest) {
    request = &QueryVoucherAmountByUinRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("intlpartnersmgt", APIVersion, "QueryVoucherAmountByUin")
    
    
    return
}

func NewQueryVoucherAmountByUinResponse() (response *QueryVoucherAmountByUinResponse) {
    response = &QueryVoucherAmountByUinResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// QueryVoucherAmountByUin
// 根据客户uin查询代金券额度
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) QueryVoucherAmountByUin(request *QueryVoucherAmountByUinRequest) (response *QueryVoucherAmountByUinResponse, err error) {
    return c.QueryVoucherAmountByUinWithContext(context.Background(), request)
}

// QueryVoucherAmountByUin
// 根据客户uin查询代金券额度
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) QueryVoucherAmountByUinWithContext(ctx context.Context, request *QueryVoucherAmountByUinRequest) (response *QueryVoucherAmountByUinResponse, err error) {
    if request == nil {
        request = NewQueryVoucherAmountByUinRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryVoucherAmountByUin require credential")
    }

    request.SetContext(ctx)
    
    response = NewQueryVoucherAmountByUinResponse()
    err = c.Send(request, response)
    return
}

func NewQueryVoucherListByUinRequest() (request *QueryVoucherListByUinRequest) {
    request = &QueryVoucherListByUinRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("intlpartnersmgt", APIVersion, "QueryVoucherListByUin")
    
    
    return
}

func NewQueryVoucherListByUinResponse() (response *QueryVoucherListByUinResponse) {
    response = &QueryVoucherListByUinResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// QueryVoucherListByUin
// 根据客户uin查询代金券列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) QueryVoucherListByUin(request *QueryVoucherListByUinRequest) (response *QueryVoucherListByUinResponse, err error) {
    return c.QueryVoucherListByUinWithContext(context.Background(), request)
}

// QueryVoucherListByUin
// 根据客户uin查询代金券列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) QueryVoucherListByUinWithContext(ctx context.Context, request *QueryVoucherListByUinRequest) (response *QueryVoucherListByUinResponse, err error) {
    if request == nil {
        request = NewQueryVoucherListByUinRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryVoucherListByUin require credential")
    }

    request.SetContext(ctx)
    
    response = NewQueryVoucherListByUinResponse()
    err = c.Send(request, response)
    return
}

func NewQueryVoucherPoolRequest() (request *QueryVoucherPoolRequest) {
    request = &QueryVoucherPoolRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("intlpartnersmgt", APIVersion, "QueryVoucherPool")
    
    
    return
}

func NewQueryVoucherPoolResponse() (response *QueryVoucherPoolResponse) {
    response = &QueryVoucherPoolResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// QueryVoucherPool
// 查询代金券额度池
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) QueryVoucherPool(request *QueryVoucherPoolRequest) (response *QueryVoucherPoolResponse, err error) {
    return c.QueryVoucherPoolWithContext(context.Background(), request)
}

// QueryVoucherPool
// 查询代金券额度池
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) QueryVoucherPoolWithContext(ctx context.Context, request *QueryVoucherPoolRequest) (response *QueryVoucherPoolResponse, err error) {
    if request == nil {
        request = NewQueryVoucherPoolRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryVoucherPool require credential")
    }

    request.SetContext(ctx)
    
    response = NewQueryVoucherPoolResponse()
    err = c.Send(request, response)
    return
}
