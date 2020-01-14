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

package v20190823

import (
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2019-08-23"

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


func NewCreateCallBackRequest() (request *CreateCallBackRequest) {
    request = &CreateCallBackRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("npp", APIVersion, "CreateCallBack")
    return
}

func NewCreateCallBackResponse() (response *CreateCallBackResponse) {
    response = &CreateCallBackResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 回拨呼叫请求
func (c *Client) CreateCallBack(request *CreateCallBackRequest) (response *CreateCallBackResponse, err error) {
    if request == nil {
        request = NewCreateCallBackRequest()
    }
    response = NewCreateCallBackResponse()
    err = c.Send(request, response)
    return
}

func NewDelVirtualNumRequest() (request *DelVirtualNumRequest) {
    request = &DelVirtualNumRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("npp", APIVersion, "DelVirtualNum")
    return
}

func NewDelVirtualNumResponse() (response *DelVirtualNumResponse) {
    response = &DelVirtualNumResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 直拨解绑中间号
func (c *Client) DelVirtualNum(request *DelVirtualNumRequest) (response *DelVirtualNumResponse, err error) {
    if request == nil {
        request = NewDelVirtualNumRequest()
    }
    response = NewDelVirtualNumResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteCallBackRequest() (request *DeleteCallBackRequest) {
    request = &DeleteCallBackRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("npp", APIVersion, "DeleteCallBack")
    return
}

func NewDeleteCallBackResponse() (response *DeleteCallBackResponse) {
    response = &DeleteCallBackResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 回拨呼叫取消
func (c *Client) DeleteCallBack(request *DeleteCallBackRequest) (response *DeleteCallBackResponse, err error) {
    if request == nil {
        request = NewDeleteCallBackRequest()
    }
    response = NewDeleteCallBackResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCallBackCdrRequest() (request *DescribeCallBackCdrRequest) {
    request = &DescribeCallBackCdrRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("npp", APIVersion, "DescribeCallBackCdr")
    return
}

func NewDescribeCallBackCdrResponse() (response *DescribeCallBackCdrResponse) {
    response = &DescribeCallBackCdrResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 回拨话单获取接口
func (c *Client) DescribeCallBackCdr(request *DescribeCallBackCdrRequest) (response *DescribeCallBackCdrResponse, err error) {
    if request == nil {
        request = NewDescribeCallBackCdrRequest()
    }
    response = NewDescribeCallBackCdrResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCallBackStatusRequest() (request *DescribeCallBackStatusRequest) {
    request = &DescribeCallBackStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("npp", APIVersion, "DescribeCallBackStatus")
    return
}

func NewDescribeCallBackStatusResponse() (response *DescribeCallBackStatusResponse) {
    response = &DescribeCallBackStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 回拨通话状态获取
func (c *Client) DescribeCallBackStatus(request *DescribeCallBackStatusRequest) (response *DescribeCallBackStatusResponse, err error) {
    if request == nil {
        request = NewDescribeCallBackStatusRequest()
    }
    response = NewDescribeCallBackStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCallerDisplayListRequest() (request *DescribeCallerDisplayListRequest) {
    request = &DescribeCallerDisplayListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("npp", APIVersion, "DescribeCallerDisplayList")
    return
}

func NewDescribeCallerDisplayListResponse() (response *DescribeCallerDisplayListResponse) {
    response = &DescribeCallerDisplayListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 回拨拉取主叫显号号码集合
func (c *Client) DescribeCallerDisplayList(request *DescribeCallerDisplayListRequest) (response *DescribeCallerDisplayListResponse, err error) {
    if request == nil {
        request = NewDescribeCallerDisplayListRequest()
    }
    response = NewDescribeCallerDisplayListResponse()
    err = c.Send(request, response)
    return
}

func NewGet400CdrRequest() (request *Get400CdrRequest) {
    request = &Get400CdrRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("npp", APIVersion, "Get400Cdr")
    return
}

func NewGet400CdrResponse() (response *Get400CdrResponse) {
    response = &Get400CdrResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 直拨话单获取接口
func (c *Client) Get400Cdr(request *Get400CdrRequest) (response *Get400CdrResponse, err error) {
    if request == nil {
        request = NewGet400CdrRequest()
    }
    response = NewGet400CdrResponse()
    err = c.Send(request, response)
    return
}

func NewGetVirtualNumRequest() (request *GetVirtualNumRequest) {
    request = &GetVirtualNumRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("npp", APIVersion, "GetVirtualNum")
    return
}

func NewGetVirtualNumResponse() (response *GetVirtualNumResponse) {
    response = &GetVirtualNumResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 直拨获取中间号（App 使用方发起）
func (c *Client) GetVirtualNum(request *GetVirtualNumRequest) (response *GetVirtualNumResponse, err error) {
    if request == nil {
        request = NewGetVirtualNumRequest()
    }
    response = NewGetVirtualNumResponse()
    err = c.Send(request, response)
    return
}
