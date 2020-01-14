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

package v20180416

import (
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2018-04-16"

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

func NewClient(credential *common.Credential, region string, clientProfile *profile.ClientProfile) (client *Client, err error) {
    client = &Client{}
    client.Init(region).
        WithCredential(credential).
        WithProfile(clientProfile)
    return
}


func NewApplyUserCertRequest() (request *ApplyUserCertRequest) {
    request = &ApplyUserCertRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbaas", APIVersion, "ApplyUserCert")
    return
}

func NewApplyUserCertResponse() (response *ApplyUserCertResponse) {
    response = &ApplyUserCertResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 申请用户证书
func (c *Client) ApplyUserCert(request *ApplyUserCertRequest) (response *ApplyUserCertResponse, err error) {
    if request == nil {
        request = NewApplyUserCertRequest()
    }
    response = NewApplyUserCertResponse()
    err = c.Send(request, response)
    return
}

func NewBlockByNumberHandlerRequest() (request *BlockByNumberHandlerRequest) {
    request = &BlockByNumberHandlerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbaas", APIVersion, "BlockByNumberHandler")
    return
}

func NewBlockByNumberHandlerResponse() (response *BlockByNumberHandlerResponse) {
    response = &BlockByNumberHandlerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// Bcos根据块高查询区块信息
func (c *Client) BlockByNumberHandler(request *BlockByNumberHandlerRequest) (response *BlockByNumberHandlerResponse, err error) {
    if request == nil {
        request = NewBlockByNumberHandlerRequest()
    }
    response = NewBlockByNumberHandlerResponse()
    err = c.Send(request, response)
    return
}

func NewDeployDynamicContractHandlerRequest() (request *DeployDynamicContractHandlerRequest) {
    request = &DeployDynamicContractHandlerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbaas", APIVersion, "DeployDynamicContractHandler")
    return
}

func NewDeployDynamicContractHandlerResponse() (response *DeployDynamicContractHandlerResponse) {
    response = &DeployDynamicContractHandlerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 动态部署合约
func (c *Client) DeployDynamicContractHandler(request *DeployDynamicContractHandlerRequest) (response *DeployDynamicContractHandlerResponse, err error) {
    if request == nil {
        request = NewDeployDynamicContractHandlerRequest()
    }
    response = NewDeployDynamicContractHandlerResponse()
    err = c.Send(request, response)
    return
}

func NewDownloadUserCertRequest() (request *DownloadUserCertRequest) {
    request = &DownloadUserCertRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbaas", APIVersion, "DownloadUserCert")
    return
}

func NewDownloadUserCertResponse() (response *DownloadUserCertResponse) {
    response = &DownloadUserCertResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 下载用户证书
func (c *Client) DownloadUserCert(request *DownloadUserCertRequest) (response *DownloadUserCertResponse, err error) {
    if request == nil {
        request = NewDownloadUserCertRequest()
    }
    response = NewDownloadUserCertResponse()
    err = c.Send(request, response)
    return
}

func NewGetBlockListRequest() (request *GetBlockListRequest) {
    request = &GetBlockListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbaas", APIVersion, "GetBlockList")
    return
}

func NewGetBlockListResponse() (response *GetBlockListResponse) {
    response = &GetBlockListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查看当前网络下的所有区块列表，分页展示
func (c *Client) GetBlockList(request *GetBlockListRequest) (response *GetBlockListResponse, err error) {
    if request == nil {
        request = NewGetBlockListRequest()
    }
    response = NewGetBlockListResponse()
    err = c.Send(request, response)
    return
}

func NewGetBlockListHandlerRequest() (request *GetBlockListHandlerRequest) {
    request = &GetBlockListHandlerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbaas", APIVersion, "GetBlockListHandler")
    return
}

func NewGetBlockListHandlerResponse() (response *GetBlockListHandlerResponse) {
    response = &GetBlockListHandlerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// Bcos分页查询当前群组下的区块列表
func (c *Client) GetBlockListHandler(request *GetBlockListHandlerRequest) (response *GetBlockListHandlerResponse, err error) {
    if request == nil {
        request = NewGetBlockListHandlerRequest()
    }
    response = NewGetBlockListHandlerResponse()
    err = c.Send(request, response)
    return
}

func NewGetBlockTransactionListForUserRequest() (request *GetBlockTransactionListForUserRequest) {
    request = &GetBlockTransactionListForUserRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbaas", APIVersion, "GetBlockTransactionListForUser")
    return
}

func NewGetBlockTransactionListForUserResponse() (response *GetBlockTransactionListForUserResponse) {
    response = &GetBlockTransactionListForUserResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取区块内的交易列表
func (c *Client) GetBlockTransactionListForUser(request *GetBlockTransactionListForUserRequest) (response *GetBlockTransactionListForUserResponse, err error) {
    if request == nil {
        request = NewGetBlockTransactionListForUserRequest()
    }
    response = NewGetBlockTransactionListForUserResponse()
    err = c.Send(request, response)
    return
}

func NewGetClusterSummaryRequest() (request *GetClusterSummaryRequest) {
    request = &GetClusterSummaryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbaas", APIVersion, "GetClusterSummary")
    return
}

func NewGetClusterSummaryResponse() (response *GetClusterSummaryResponse) {
    response = &GetClusterSummaryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取区块链网络概要
func (c *Client) GetClusterSummary(request *GetClusterSummaryRequest) (response *GetClusterSummaryResponse, err error) {
    if request == nil {
        request = NewGetClusterSummaryRequest()
    }
    response = NewGetClusterSummaryResponse()
    err = c.Send(request, response)
    return
}

func NewGetInvokeTxRequest() (request *GetInvokeTxRequest) {
    request = &GetInvokeTxRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbaas", APIVersion, "GetInvokeTx")
    return
}

func NewGetInvokeTxResponse() (response *GetInvokeTxResponse) {
    response = &GetInvokeTxResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// Invoke异步调用结果查询
func (c *Client) GetInvokeTx(request *GetInvokeTxRequest) (response *GetInvokeTxResponse, err error) {
    if request == nil {
        request = NewGetInvokeTxRequest()
    }
    response = NewGetInvokeTxResponse()
    err = c.Send(request, response)
    return
}

func NewGetLatesdTransactionListRequest() (request *GetLatesdTransactionListRequest) {
    request = &GetLatesdTransactionListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbaas", APIVersion, "GetLatesdTransactionList")
    return
}

func NewGetLatesdTransactionListResponse() (response *GetLatesdTransactionListResponse) {
    response = &GetLatesdTransactionListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取最新交易列表
func (c *Client) GetLatesdTransactionList(request *GetLatesdTransactionListRequest) (response *GetLatesdTransactionListResponse, err error) {
    if request == nil {
        request = NewGetLatesdTransactionListRequest()
    }
    response = NewGetLatesdTransactionListResponse()
    err = c.Send(request, response)
    return
}

func NewGetTransByHashHandlerRequest() (request *GetTransByHashHandlerRequest) {
    request = &GetTransByHashHandlerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbaas", APIVersion, "GetTransByHashHandler")
    return
}

func NewGetTransByHashHandlerResponse() (response *GetTransByHashHandlerResponse) {
    response = &GetTransByHashHandlerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// Bcos根据交易哈希查看交易详细信息
func (c *Client) GetTransByHashHandler(request *GetTransByHashHandlerRequest) (response *GetTransByHashHandlerResponse, err error) {
    if request == nil {
        request = NewGetTransByHashHandlerRequest()
    }
    response = NewGetTransByHashHandlerResponse()
    err = c.Send(request, response)
    return
}

func NewGetTransListHandlerRequest() (request *GetTransListHandlerRequest) {
    request = &GetTransListHandlerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbaas", APIVersion, "GetTransListHandler")
    return
}

func NewGetTransListHandlerResponse() (response *GetTransListHandlerResponse) {
    response = &GetTransListHandlerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// Bcos分页查询当前群组的交易信息列表
func (c *Client) GetTransListHandler(request *GetTransListHandlerRequest) (response *GetTransListHandlerResponse, err error) {
    if request == nil {
        request = NewGetTransListHandlerRequest()
    }
    response = NewGetTransListHandlerResponse()
    err = c.Send(request, response)
    return
}

func NewGetTransactionDetailForUserRequest() (request *GetTransactionDetailForUserRequest) {
    request = &GetTransactionDetailForUserRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbaas", APIVersion, "GetTransactionDetailForUser")
    return
}

func NewGetTransactionDetailForUserResponse() (response *GetTransactionDetailForUserResponse) {
    response = &GetTransactionDetailForUserResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取交易详情
func (c *Client) GetTransactionDetailForUser(request *GetTransactionDetailForUserRequest) (response *GetTransactionDetailForUserResponse, err error) {
    if request == nil {
        request = NewGetTransactionDetailForUserRequest()
    }
    response = NewGetTransactionDetailForUserResponse()
    err = c.Send(request, response)
    return
}

func NewInvokeRequest() (request *InvokeRequest) {
    request = &InvokeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbaas", APIVersion, "Invoke")
    return
}

func NewInvokeResponse() (response *InvokeResponse) {
    response = &InvokeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 新增交易
func (c *Client) Invoke(request *InvokeRequest) (response *InvokeResponse, err error) {
    if request == nil {
        request = NewInvokeRequest()
    }
    response = NewInvokeResponse()
    err = c.Send(request, response)
    return
}

func NewQueryRequest() (request *QueryRequest) {
    request = &QueryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbaas", APIVersion, "Query")
    return
}

func NewQueryResponse() (response *QueryResponse) {
    response = &QueryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询交易
func (c *Client) Query(request *QueryRequest) (response *QueryResponse, err error) {
    if request == nil {
        request = NewQueryRequest()
    }
    response = NewQueryResponse()
    err = c.Send(request, response)
    return
}

func NewSendTransactionHandlerRequest() (request *SendTransactionHandlerRequest) {
    request = &SendTransactionHandlerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbaas", APIVersion, "SendTransactionHandler")
    return
}

func NewSendTransactionHandlerResponse() (response *SendTransactionHandlerResponse) {
    response = &SendTransactionHandlerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// Bcos发送交易
func (c *Client) SendTransactionHandler(request *SendTransactionHandlerRequest) (response *SendTransactionHandlerResponse, err error) {
    if request == nil {
        request = NewSendTransactionHandlerRequest()
    }
    response = NewSendTransactionHandlerResponse()
    err = c.Send(request, response)
    return
}

func NewSrvInvokeRequest() (request *SrvInvokeRequest) {
    request = &SrvInvokeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbaas", APIVersion, "SrvInvoke")
    return
}

func NewSrvInvokeResponse() (response *SrvInvokeResponse) {
    response = &SrvInvokeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// trustsql服务统一接口
func (c *Client) SrvInvoke(request *SrvInvokeRequest) (response *SrvInvokeResponse, err error) {
    if request == nil {
        request = NewSrvInvokeRequest()
    }
    response = NewSrvInvokeResponse()
    err = c.Send(request, response)
    return
}

func NewTransByDynamicContractHandlerRequest() (request *TransByDynamicContractHandlerRequest) {
    request = &TransByDynamicContractHandlerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbaas", APIVersion, "TransByDynamicContractHandler")
    return
}

func NewTransByDynamicContractHandlerResponse() (response *TransByDynamicContractHandlerResponse) {
    response = &TransByDynamicContractHandlerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 根据动态部署的合约发送交易
func (c *Client) TransByDynamicContractHandler(request *TransByDynamicContractHandlerRequest) (response *TransByDynamicContractHandlerResponse, err error) {
    if request == nil {
        request = NewTransByDynamicContractHandlerRequest()
    }
    response = NewTransByDynamicContractHandlerResponse()
    err = c.Send(request, response)
    return
}
