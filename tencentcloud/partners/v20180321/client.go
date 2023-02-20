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

package v20180321

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2018-03-21"

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


func NewAgentPayDealsRequest() (request *AgentPayDealsRequest) {
    request = &AgentPayDealsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("partners", APIVersion, "AgentPayDeals")
    
    
    return
}

func NewAgentPayDealsResponse() (response *AgentPayDealsResponse) {
    response = &AgentPayDealsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AgentPayDeals
// 代理商支付订单接口，支持自付/代付
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) AgentPayDeals(request *AgentPayDealsRequest) (response *AgentPayDealsResponse, err error) {
    return c.AgentPayDealsWithContext(context.Background(), request)
}

// AgentPayDeals
// 代理商支付订单接口，支持自付/代付
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) AgentPayDealsWithContext(ctx context.Context, request *AgentPayDealsRequest) (response *AgentPayDealsResponse, err error) {
    if request == nil {
        request = NewAgentPayDealsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AgentPayDeals require credential")
    }

    request.SetContext(ctx)
    
    response = NewAgentPayDealsResponse()
    err = c.Send(request, response)
    return
}

func NewAgentTransferMoneyRequest() (request *AgentTransferMoneyRequest) {
    request = &AgentTransferMoneyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("partners", APIVersion, "AgentTransferMoney")
    
    
    return
}

func NewAgentTransferMoneyResponse() (response *AgentTransferMoneyResponse) {
    response = &AgentTransferMoneyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AgentTransferMoney
// 为合作伙伴提供转账给客户能力。仅支持合作伙伴为自己名下客户转账。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) AgentTransferMoney(request *AgentTransferMoneyRequest) (response *AgentTransferMoneyResponse, err error) {
    return c.AgentTransferMoneyWithContext(context.Background(), request)
}

// AgentTransferMoney
// 为合作伙伴提供转账给客户能力。仅支持合作伙伴为自己名下客户转账。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) AgentTransferMoneyWithContext(ctx context.Context, request *AgentTransferMoneyRequest) (response *AgentTransferMoneyResponse, err error) {
    if request == nil {
        request = NewAgentTransferMoneyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AgentTransferMoney require credential")
    }

    request.SetContext(ctx)
    
    response = NewAgentTransferMoneyResponse()
    err = c.Send(request, response)
    return
}

func NewAssignClientsToSalesRequest() (request *AssignClientsToSalesRequest) {
    request = &AssignClientsToSalesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("partners", APIVersion, "AssignClientsToSales")
    
    
    return
}

func NewAssignClientsToSalesResponse() (response *AssignClientsToSalesResponse) {
    response = &AssignClientsToSalesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AssignClientsToSales
// 为代客or申请中代客分派跟进人（业务员），入参可从以下API获取
//
// - 代客列表获取API： [DescribeAgentAuditedClients](https://cloud.tencent.com/document/product/563/19184)
//
// - 申请中代客列表获取API：[DescribeAgentClients](https://cloud.tencent.com/document/product/563/16046)
//
// - 业务员列表获取API：[DescribeSalesmans](https://cloud.tencent.com/document/product/563/35196) <br><br>
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) AssignClientsToSales(request *AssignClientsToSalesRequest) (response *AssignClientsToSalesResponse, err error) {
    return c.AssignClientsToSalesWithContext(context.Background(), request)
}

// AssignClientsToSales
// 为代客or申请中代客分派跟进人（业务员），入参可从以下API获取
//
// - 代客列表获取API： [DescribeAgentAuditedClients](https://cloud.tencent.com/document/product/563/19184)
//
// - 申请中代客列表获取API：[DescribeAgentClients](https://cloud.tencent.com/document/product/563/16046)
//
// - 业务员列表获取API：[DescribeSalesmans](https://cloud.tencent.com/document/product/563/35196) <br><br>
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) AssignClientsToSalesWithContext(ctx context.Context, request *AssignClientsToSalesRequest) (response *AssignClientsToSalesResponse, err error) {
    if request == nil {
        request = NewAssignClientsToSalesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AssignClientsToSales require credential")
    }

    request.SetContext(ctx)
    
    response = NewAssignClientsToSalesResponse()
    err = c.Send(request, response)
    return
}

func NewAuditApplyClientRequest() (request *AuditApplyClientRequest) {
    request = &AuditApplyClientRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("partners", APIVersion, "AuditApplyClient")
    
    
    return
}

func NewAuditApplyClientResponse() (response *AuditApplyClientResponse) {
    response = &AuditApplyClientResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AuditApplyClient
// 代理商可以审核其名下申请中代客
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) AuditApplyClient(request *AuditApplyClientRequest) (response *AuditApplyClientResponse, err error) {
    return c.AuditApplyClientWithContext(context.Background(), request)
}

// AuditApplyClient
// 代理商可以审核其名下申请中代客
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) AuditApplyClientWithContext(ctx context.Context, request *AuditApplyClientRequest) (response *AuditApplyClientResponse, err error) {
    if request == nil {
        request = NewAuditApplyClientRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AuditApplyClient require credential")
    }

    request.SetContext(ctx)
    
    response = NewAuditApplyClientResponse()
    err = c.Send(request, response)
    return
}

func NewCreatePayRelationForClientRequest() (request *CreatePayRelationForClientRequest) {
    request = &CreatePayRelationForClientRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("partners", APIVersion, "CreatePayRelationForClient")
    
    
    return
}

func NewCreatePayRelationForClientResponse() (response *CreatePayRelationForClientResponse) {
    response = &CreatePayRelationForClientResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreatePayRelationForClient
// 合作伙伴为客户创建强代付关系
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CreatePayRelationForClient(request *CreatePayRelationForClientRequest) (response *CreatePayRelationForClientResponse, err error) {
    return c.CreatePayRelationForClientWithContext(context.Background(), request)
}

// CreatePayRelationForClient
// 合作伙伴为客户创建强代付关系
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CreatePayRelationForClientWithContext(ctx context.Context, request *CreatePayRelationForClientRequest) (response *CreatePayRelationForClientResponse, err error) {
    if request == nil {
        request = NewCreatePayRelationForClientRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreatePayRelationForClient require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreatePayRelationForClientResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAgentAuditedClientsRequest() (request *DescribeAgentAuditedClientsRequest) {
    request = &DescribeAgentAuditedClientsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("partners", APIVersion, "DescribeAgentAuditedClients")
    
    
    return
}

func NewDescribeAgentAuditedClientsResponse() (response *DescribeAgentAuditedClientsResponse) {
    response = &DescribeAgentAuditedClientsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAgentAuditedClients
// 查询已审核客户列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeAgentAuditedClients(request *DescribeAgentAuditedClientsRequest) (response *DescribeAgentAuditedClientsResponse, err error) {
    return c.DescribeAgentAuditedClientsWithContext(context.Background(), request)
}

// DescribeAgentAuditedClients
// 查询已审核客户列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeAgentAuditedClientsWithContext(ctx context.Context, request *DescribeAgentAuditedClientsRequest) (response *DescribeAgentAuditedClientsResponse, err error) {
    if request == nil {
        request = NewDescribeAgentAuditedClientsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAgentAuditedClients require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAgentAuditedClientsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAgentBillsRequest() (request *DescribeAgentBillsRequest) {
    request = &DescribeAgentBillsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("partners", APIVersion, "DescribeAgentBills")
    
    
    return
}

func NewDescribeAgentBillsResponse() (response *DescribeAgentBillsResponse) {
    response = &DescribeAgentBillsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAgentBills
// 代理商可查询自己及名下代客所有业务明细
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeAgentBills(request *DescribeAgentBillsRequest) (response *DescribeAgentBillsResponse, err error) {
    return c.DescribeAgentBillsWithContext(context.Background(), request)
}

// DescribeAgentBills
// 代理商可查询自己及名下代客所有业务明细
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeAgentBillsWithContext(ctx context.Context, request *DescribeAgentBillsRequest) (response *DescribeAgentBillsResponse, err error) {
    if request == nil {
        request = NewDescribeAgentBillsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAgentBills require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAgentBillsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAgentClientGradeRequest() (request *DescribeAgentClientGradeRequest) {
    request = &DescribeAgentClientGradeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("partners", APIVersion, "DescribeAgentClientGrade")
    
    
    return
}

func NewDescribeAgentClientGradeResponse() (response *DescribeAgentClientGradeResponse) {
    response = &DescribeAgentClientGradeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAgentClientGrade
// 传入代客uin，查客户级别，客户审核状态，客户实名认证状态
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeAgentClientGrade(request *DescribeAgentClientGradeRequest) (response *DescribeAgentClientGradeResponse, err error) {
    return c.DescribeAgentClientGradeWithContext(context.Background(), request)
}

// DescribeAgentClientGrade
// 传入代客uin，查客户级别，客户审核状态，客户实名认证状态
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeAgentClientGradeWithContext(ctx context.Context, request *DescribeAgentClientGradeRequest) (response *DescribeAgentClientGradeResponse, err error) {
    if request == nil {
        request = NewDescribeAgentClientGradeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAgentClientGrade require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAgentClientGradeResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAgentClientsRequest() (request *DescribeAgentClientsRequest) {
    request = &DescribeAgentClientsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("partners", APIVersion, "DescribeAgentClients")
    
    
    return
}

func NewDescribeAgentClientsResponse() (response *DescribeAgentClientsResponse) {
    response = &DescribeAgentClientsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAgentClients
// 代理商可查询自己名下待审核客户列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeAgentClients(request *DescribeAgentClientsRequest) (response *DescribeAgentClientsResponse, err error) {
    return c.DescribeAgentClientsWithContext(context.Background(), request)
}

// DescribeAgentClients
// 代理商可查询自己名下待审核客户列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeAgentClientsWithContext(ctx context.Context, request *DescribeAgentClientsRequest) (response *DescribeAgentClientsResponse, err error) {
    if request == nil {
        request = NewDescribeAgentClientsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAgentClients require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAgentClientsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAgentDealsByCacheRequest() (request *DescribeAgentDealsByCacheRequest) {
    request = &DescribeAgentDealsByCacheRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("partners", APIVersion, "DescribeAgentDealsByCache")
    
    
    return
}

func NewDescribeAgentDealsByCacheResponse() (response *DescribeAgentDealsByCacheResponse) {
    response = &DescribeAgentDealsByCacheResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAgentDealsByCache
// 供代理商拉取缓存的全量预付费客户订单
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeAgentDealsByCache(request *DescribeAgentDealsByCacheRequest) (response *DescribeAgentDealsByCacheResponse, err error) {
    return c.DescribeAgentDealsByCacheWithContext(context.Background(), request)
}

// DescribeAgentDealsByCache
// 供代理商拉取缓存的全量预付费客户订单
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeAgentDealsByCacheWithContext(ctx context.Context, request *DescribeAgentDealsByCacheRequest) (response *DescribeAgentDealsByCacheResponse, err error) {
    if request == nil {
        request = NewDescribeAgentDealsByCacheRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAgentDealsByCache require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAgentDealsByCacheResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAgentDealsCacheRequest() (request *DescribeAgentDealsCacheRequest) {
    request = &DescribeAgentDealsCacheRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("partners", APIVersion, "DescribeAgentDealsCache")
    
    
    return
}

func NewDescribeAgentDealsCacheResponse() (response *DescribeAgentDealsCacheResponse) {
    response = &DescribeAgentDealsCacheResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAgentDealsCache
// 【该接口已下线，请使用升级版本DescribeAgentDealsByCache】代理商拉取缓存的全量客户订单
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeAgentDealsCache(request *DescribeAgentDealsCacheRequest) (response *DescribeAgentDealsCacheResponse, err error) {
    return c.DescribeAgentDealsCacheWithContext(context.Background(), request)
}

// DescribeAgentDealsCache
// 【该接口已下线，请使用升级版本DescribeAgentDealsByCache】代理商拉取缓存的全量客户订单
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeAgentDealsCacheWithContext(ctx context.Context, request *DescribeAgentDealsCacheRequest) (response *DescribeAgentDealsCacheResponse, err error) {
    if request == nil {
        request = NewDescribeAgentDealsCacheRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAgentDealsCache require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAgentDealsCacheResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAgentPayDealsRequest() (request *DescribeAgentPayDealsRequest) {
    request = &DescribeAgentPayDealsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("partners", APIVersion, "DescribeAgentPayDeals")
    
    
    return
}

func NewDescribeAgentPayDealsResponse() (response *DescribeAgentPayDealsResponse) {
    response = &DescribeAgentPayDealsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAgentPayDeals
// 【该接口已下线，请切换使用升级版本DescribeAgentPayDealsV2】可以查询代理商代付的所有订单
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeAgentPayDeals(request *DescribeAgentPayDealsRequest) (response *DescribeAgentPayDealsResponse, err error) {
    return c.DescribeAgentPayDealsWithContext(context.Background(), request)
}

// DescribeAgentPayDeals
// 【该接口已下线，请切换使用升级版本DescribeAgentPayDealsV2】可以查询代理商代付的所有订单
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeAgentPayDealsWithContext(ctx context.Context, request *DescribeAgentPayDealsRequest) (response *DescribeAgentPayDealsResponse, err error) {
    if request == nil {
        request = NewDescribeAgentPayDealsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAgentPayDeals require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAgentPayDealsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAgentPayDealsV2Request() (request *DescribeAgentPayDealsV2Request) {
    request = &DescribeAgentPayDealsV2Request{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("partners", APIVersion, "DescribeAgentPayDealsV2")
    
    
    return
}

func NewDescribeAgentPayDealsV2Response() (response *DescribeAgentPayDealsV2Response) {
    response = &DescribeAgentPayDealsV2Response{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAgentPayDealsV2
// 可以查询代理商代付的预付费订单
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeAgentPayDealsV2(request *DescribeAgentPayDealsV2Request) (response *DescribeAgentPayDealsV2Response, err error) {
    return c.DescribeAgentPayDealsV2WithContext(context.Background(), request)
}

// DescribeAgentPayDealsV2
// 可以查询代理商代付的预付费订单
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeAgentPayDealsV2WithContext(ctx context.Context, request *DescribeAgentPayDealsV2Request) (response *DescribeAgentPayDealsV2Response, err error) {
    if request == nil {
        request = NewDescribeAgentPayDealsV2Request()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAgentPayDealsV2 require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAgentPayDealsV2Response()
    err = c.Send(request, response)
    return
}

func NewDescribeAgentSelfPayDealsRequest() (request *DescribeAgentSelfPayDealsRequest) {
    request = &DescribeAgentSelfPayDealsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("partners", APIVersion, "DescribeAgentSelfPayDeals")
    
    
    return
}

func NewDescribeAgentSelfPayDealsResponse() (response *DescribeAgentSelfPayDealsResponse) {
    response = &DescribeAgentSelfPayDealsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAgentSelfPayDeals
// 【该接口已下线，请切换使用升级版本DescribeAgentSelfPayDealsV2】可以查询代理商下指定客户的自付订单
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeAgentSelfPayDeals(request *DescribeAgentSelfPayDealsRequest) (response *DescribeAgentSelfPayDealsResponse, err error) {
    return c.DescribeAgentSelfPayDealsWithContext(context.Background(), request)
}

// DescribeAgentSelfPayDeals
// 【该接口已下线，请切换使用升级版本DescribeAgentSelfPayDealsV2】可以查询代理商下指定客户的自付订单
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeAgentSelfPayDealsWithContext(ctx context.Context, request *DescribeAgentSelfPayDealsRequest) (response *DescribeAgentSelfPayDealsResponse, err error) {
    if request == nil {
        request = NewDescribeAgentSelfPayDealsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAgentSelfPayDeals require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAgentSelfPayDealsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAgentSelfPayDealsV2Request() (request *DescribeAgentSelfPayDealsV2Request) {
    request = &DescribeAgentSelfPayDealsV2Request{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("partners", APIVersion, "DescribeAgentSelfPayDealsV2")
    
    
    return
}

func NewDescribeAgentSelfPayDealsV2Response() (response *DescribeAgentSelfPayDealsV2Response) {
    response = &DescribeAgentSelfPayDealsV2Response{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAgentSelfPayDealsV2
// 查询代理商名下指定代客的自付订单（预付费）
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeAgentSelfPayDealsV2(request *DescribeAgentSelfPayDealsV2Request) (response *DescribeAgentSelfPayDealsV2Response, err error) {
    return c.DescribeAgentSelfPayDealsV2WithContext(context.Background(), request)
}

// DescribeAgentSelfPayDealsV2
// 查询代理商名下指定代客的自付订单（预付费）
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeAgentSelfPayDealsV2WithContext(ctx context.Context, request *DescribeAgentSelfPayDealsV2Request) (response *DescribeAgentSelfPayDealsV2Response, err error) {
    if request == nil {
        request = NewDescribeAgentSelfPayDealsV2Request()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAgentSelfPayDealsV2 require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAgentSelfPayDealsV2Response()
    err = c.Send(request, response)
    return
}

func NewDescribeClientBalanceRequest() (request *DescribeClientBalanceRequest) {
    request = &DescribeClientBalanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("partners", APIVersion, "DescribeClientBalance")
    
    
    return
}

func NewDescribeClientBalanceResponse() (response *DescribeClientBalanceResponse) {
    response = &DescribeClientBalanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeClientBalance
// 【该接口将逐步下线，请切换使用升级版本DescribeClientBalanceNew】为合作伙伴提供查询客户余额能力。调用者必须是合作伙伴，只能查询自己名下客户余额.
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeClientBalance(request *DescribeClientBalanceRequest) (response *DescribeClientBalanceResponse, err error) {
    return c.DescribeClientBalanceWithContext(context.Background(), request)
}

// DescribeClientBalance
// 【该接口将逐步下线，请切换使用升级版本DescribeClientBalanceNew】为合作伙伴提供查询客户余额能力。调用者必须是合作伙伴，只能查询自己名下客户余额.
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeClientBalanceWithContext(ctx context.Context, request *DescribeClientBalanceRequest) (response *DescribeClientBalanceResponse, err error) {
    if request == nil {
        request = NewDescribeClientBalanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeClientBalance require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeClientBalanceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeClientBalanceNewRequest() (request *DescribeClientBalanceNewRequest) {
    request = &DescribeClientBalanceNewRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("partners", APIVersion, "DescribeClientBalanceNew")
    
    
    return
}

func NewDescribeClientBalanceNewResponse() (response *DescribeClientBalanceNewResponse) {
    response = &DescribeClientBalanceNewResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeClientBalanceNew
// 为合作伙伴提供查询客户余额能力。调用者必须是合作伙伴，只能查询自己名下客户余额
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeClientBalanceNew(request *DescribeClientBalanceNewRequest) (response *DescribeClientBalanceNewResponse, err error) {
    return c.DescribeClientBalanceNewWithContext(context.Background(), request)
}

// DescribeClientBalanceNew
// 为合作伙伴提供查询客户余额能力。调用者必须是合作伙伴，只能查询自己名下客户余额
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeClientBalanceNewWithContext(ctx context.Context, request *DescribeClientBalanceNewRequest) (response *DescribeClientBalanceNewResponse, err error) {
    if request == nil {
        request = NewDescribeClientBalanceNewRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeClientBalanceNew require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeClientBalanceNewResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRebateInfosRequest() (request *DescribeRebateInfosRequest) {
    request = &DescribeRebateInfosRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("partners", APIVersion, "DescribeRebateInfos")
    
    
    return
}

func NewDescribeRebateInfosResponse() (response *DescribeRebateInfosResponse) {
    response = &DescribeRebateInfosResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeRebateInfos
// 【该接口已下线，请切换使用升级版本DescribeRebateInfosNew】代理商可查询自己名下全部返佣信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeRebateInfos(request *DescribeRebateInfosRequest) (response *DescribeRebateInfosResponse, err error) {
    return c.DescribeRebateInfosWithContext(context.Background(), request)
}

// DescribeRebateInfos
// 【该接口已下线，请切换使用升级版本DescribeRebateInfosNew】代理商可查询自己名下全部返佣信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeRebateInfosWithContext(ctx context.Context, request *DescribeRebateInfosRequest) (response *DescribeRebateInfosResponse, err error) {
    if request == nil {
        request = NewDescribeRebateInfosRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRebateInfos require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRebateInfosResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRebateInfosNewRequest() (request *DescribeRebateInfosNewRequest) {
    request = &DescribeRebateInfosNewRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("partners", APIVersion, "DescribeRebateInfosNew")
    
    
    return
}

func NewDescribeRebateInfosNewResponse() (response *DescribeRebateInfosNewResponse) {
    response = &DescribeRebateInfosNewResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeRebateInfosNew
// 代理商可查询自己名下全部返佣信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeRebateInfosNew(request *DescribeRebateInfosNewRequest) (response *DescribeRebateInfosNewResponse, err error) {
    return c.DescribeRebateInfosNewWithContext(context.Background(), request)
}

// DescribeRebateInfosNew
// 代理商可查询自己名下全部返佣信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeRebateInfosNewWithContext(ctx context.Context, request *DescribeRebateInfosNewRequest) (response *DescribeRebateInfosNewResponse, err error) {
    if request == nil {
        request = NewDescribeRebateInfosNewRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRebateInfosNew require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRebateInfosNewResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSalesmansRequest() (request *DescribeSalesmansRequest) {
    request = &DescribeSalesmansRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("partners", APIVersion, "DescribeSalesmans")
    
    
    return
}

func NewDescribeSalesmansResponse() (response *DescribeSalesmansResponse) {
    response = &DescribeSalesmansResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeSalesmans
// 代理商查询名下业务员列表信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeSalesmans(request *DescribeSalesmansRequest) (response *DescribeSalesmansResponse, err error) {
    return c.DescribeSalesmansWithContext(context.Background(), request)
}

// DescribeSalesmans
// 代理商查询名下业务员列表信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeSalesmansWithContext(ctx context.Context, request *DescribeSalesmansRequest) (response *DescribeSalesmansResponse, err error) {
    if request == nil {
        request = NewDescribeSalesmansRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSalesmans require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSalesmansResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUnbindClientListRequest() (request *DescribeUnbindClientListRequest) {
    request = &DescribeUnbindClientListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("partners", APIVersion, "DescribeUnbindClientList")
    
    
    return
}

func NewDescribeUnbindClientListResponse() (response *DescribeUnbindClientListResponse) {
    response = &DescribeUnbindClientListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeUnbindClientList
// 代理商名下客户解绑记录查询接口
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeUnbindClientList(request *DescribeUnbindClientListRequest) (response *DescribeUnbindClientListResponse, err error) {
    return c.DescribeUnbindClientListWithContext(context.Background(), request)
}

// DescribeUnbindClientList
// 代理商名下客户解绑记录查询接口
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeUnbindClientListWithContext(ctx context.Context, request *DescribeUnbindClientListRequest) (response *DescribeUnbindClientListResponse, err error) {
    if request == nil {
        request = NewDescribeUnbindClientListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUnbindClientList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeUnbindClientListResponse()
    err = c.Send(request, response)
    return
}

func NewModifyClientRemarkRequest() (request *ModifyClientRemarkRequest) {
    request = &ModifyClientRemarkRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("partners", APIVersion, "ModifyClientRemark")
    
    
    return
}

func NewModifyClientRemarkResponse() (response *ModifyClientRemarkResponse) {
    response = &ModifyClientRemarkResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyClientRemark
// 代理商可以对名下客户添加备注、修改备注
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ModifyClientRemark(request *ModifyClientRemarkRequest) (response *ModifyClientRemarkResponse, err error) {
    return c.ModifyClientRemarkWithContext(context.Background(), request)
}

// ModifyClientRemark
// 代理商可以对名下客户添加备注、修改备注
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ModifyClientRemarkWithContext(ctx context.Context, request *ModifyClientRemarkRequest) (response *ModifyClientRemarkResponse, err error) {
    if request == nil {
        request = NewModifyClientRemarkRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyClientRemark require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyClientRemarkResponse()
    err = c.Send(request, response)
    return
}

func NewRemovePayRelationForClientRequest() (request *RemovePayRelationForClientRequest) {
    request = &RemovePayRelationForClientRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("partners", APIVersion, "RemovePayRelationForClient")
    
    
    return
}

func NewRemovePayRelationForClientResponse() (response *RemovePayRelationForClientResponse) {
    response = &RemovePayRelationForClientResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// RemovePayRelationForClient
// 合作伙伴为客户消除强代付关系
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) RemovePayRelationForClient(request *RemovePayRelationForClientRequest) (response *RemovePayRelationForClientResponse, err error) {
    return c.RemovePayRelationForClientWithContext(context.Background(), request)
}

// RemovePayRelationForClient
// 合作伙伴为客户消除强代付关系
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) RemovePayRelationForClientWithContext(ctx context.Context, request *RemovePayRelationForClientRequest) (response *RemovePayRelationForClientResponse, err error) {
    if request == nil {
        request = NewRemovePayRelationForClientRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RemovePayRelationForClient require credential")
    }

    request.SetContext(ctx)
    
    response = NewRemovePayRelationForClientResponse()
    err = c.Send(request, response)
    return
}
