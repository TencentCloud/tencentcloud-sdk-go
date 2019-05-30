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
