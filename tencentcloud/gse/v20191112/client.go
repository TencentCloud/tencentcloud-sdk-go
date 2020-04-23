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

package v20191112

import (
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2019-11-12"

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


func NewCreateGameServerSessionRequest() (request *CreateGameServerSessionRequest) {
    request = &CreateGameServerSessionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gse", APIVersion, "CreateGameServerSession")
    return
}

func NewCreateGameServerSessionResponse() (response *CreateGameServerSessionResponse) {
    response = &CreateGameServerSessionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（CreateGameServerSession）用于创建游戏服务会话
func (c *Client) CreateGameServerSession(request *CreateGameServerSessionRequest) (response *CreateGameServerSessionResponse, err error) {
    if request == nil {
        request = NewCreateGameServerSessionRequest()
    }
    response = NewCreateGameServerSessionResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeGameServerSessionDetailsRequest() (request *DescribeGameServerSessionDetailsRequest) {
    request = &DescribeGameServerSessionDetailsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gse", APIVersion, "DescribeGameServerSessionDetails")
    return
}

func NewDescribeGameServerSessionDetailsResponse() (response *DescribeGameServerSessionDetailsResponse) {
    response = &DescribeGameServerSessionDetailsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DescribeGameServerSessionDetails）用于查询游戏服务器会话详情列表
func (c *Client) DescribeGameServerSessionDetails(request *DescribeGameServerSessionDetailsRequest) (response *DescribeGameServerSessionDetailsResponse, err error) {
    if request == nil {
        request = NewDescribeGameServerSessionDetailsRequest()
    }
    response = NewDescribeGameServerSessionDetailsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeGameServerSessionPlacementRequest() (request *DescribeGameServerSessionPlacementRequest) {
    request = &DescribeGameServerSessionPlacementRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gse", APIVersion, "DescribeGameServerSessionPlacement")
    return
}

func NewDescribeGameServerSessionPlacementResponse() (response *DescribeGameServerSessionPlacementResponse) {
    response = &DescribeGameServerSessionPlacementResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DescribeGameServerSessionPlacement）用于查询游戏服务器会话的放置
func (c *Client) DescribeGameServerSessionPlacement(request *DescribeGameServerSessionPlacementRequest) (response *DescribeGameServerSessionPlacementResponse, err error) {
    if request == nil {
        request = NewDescribeGameServerSessionPlacementRequest()
    }
    response = NewDescribeGameServerSessionPlacementResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeGameServerSessionsRequest() (request *DescribeGameServerSessionsRequest) {
    request = &DescribeGameServerSessionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gse", APIVersion, "DescribeGameServerSessions")
    return
}

func NewDescribeGameServerSessionsResponse() (response *DescribeGameServerSessionsResponse) {
    response = &DescribeGameServerSessionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DescribeGameServerSessions）用于查询游戏服务器会话列表
func (c *Client) DescribeGameServerSessions(request *DescribeGameServerSessionsRequest) (response *DescribeGameServerSessionsResponse, err error) {
    if request == nil {
        request = NewDescribeGameServerSessionsRequest()
    }
    response = NewDescribeGameServerSessionsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstancesRequest() (request *DescribeInstancesRequest) {
    request = &DescribeInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gse", APIVersion, "DescribeInstances")
    return
}

func NewDescribeInstancesResponse() (response *DescribeInstancesResponse) {
    response = &DescribeInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 用于查询服务器实例列表
func (c *Client) DescribeInstances(request *DescribeInstancesRequest) (response *DescribeInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeInstancesRequest()
    }
    response = NewDescribeInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePlayerSessionsRequest() (request *DescribePlayerSessionsRequest) {
    request = &DescribePlayerSessionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gse", APIVersion, "DescribePlayerSessions")
    return
}

func NewDescribePlayerSessionsResponse() (response *DescribePlayerSessionsResponse) {
    response = &DescribePlayerSessionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DescribePlayerSessions）用于获取玩家会话列表
func (c *Client) DescribePlayerSessions(request *DescribePlayerSessionsRequest) (response *DescribePlayerSessionsResponse, err error) {
    if request == nil {
        request = NewDescribePlayerSessionsRequest()
    }
    response = NewDescribePlayerSessionsResponse()
    err = c.Send(request, response)
    return
}

func NewGetGameServerSessionLogUrlRequest() (request *GetGameServerSessionLogUrlRequest) {
    request = &GetGameServerSessionLogUrlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gse", APIVersion, "GetGameServerSessionLogUrl")
    return
}

func NewGetGameServerSessionLogUrlResponse() (response *GetGameServerSessionLogUrlResponse) {
    response = &GetGameServerSessionLogUrlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（GetGameServerSessionLogUrl）用于获取游戏服务器会话的日志URL
func (c *Client) GetGameServerSessionLogUrl(request *GetGameServerSessionLogUrlRequest) (response *GetGameServerSessionLogUrlResponse, err error) {
    if request == nil {
        request = NewGetGameServerSessionLogUrlRequest()
    }
    response = NewGetGameServerSessionLogUrlResponse()
    err = c.Send(request, response)
    return
}

func NewGetInstanceAccessRequest() (request *GetInstanceAccessRequest) {
    request = &GetInstanceAccessRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gse", APIVersion, "GetInstanceAccess")
    return
}

func NewGetInstanceAccessResponse() (response *GetInstanceAccessResponse) {
    response = &GetInstanceAccessResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取实例登录所需要的凭据
func (c *Client) GetInstanceAccess(request *GetInstanceAccessRequest) (response *GetInstanceAccessResponse, err error) {
    if request == nil {
        request = NewGetInstanceAccessRequest()
    }
    response = NewGetInstanceAccessResponse()
    err = c.Send(request, response)
    return
}

func NewJoinGameServerSessionRequest() (request *JoinGameServerSessionRequest) {
    request = &JoinGameServerSessionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gse", APIVersion, "JoinGameServerSession")
    return
}

func NewJoinGameServerSessionResponse() (response *JoinGameServerSessionResponse) {
    response = &JoinGameServerSessionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（JoinGameServerSession）用于加入游戏服务器会话
func (c *Client) JoinGameServerSession(request *JoinGameServerSessionRequest) (response *JoinGameServerSessionResponse, err error) {
    if request == nil {
        request = NewJoinGameServerSessionRequest()
    }
    response = NewJoinGameServerSessionResponse()
    err = c.Send(request, response)
    return
}

func NewSearchGameServerSessionsRequest() (request *SearchGameServerSessionsRequest) {
    request = &SearchGameServerSessionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gse", APIVersion, "SearchGameServerSessions")
    return
}

func NewSearchGameServerSessionsResponse() (response *SearchGameServerSessionsResponse) {
    response = &SearchGameServerSessionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（SearchGameServerSessions）用于搜索游戏服务器会话列表
func (c *Client) SearchGameServerSessions(request *SearchGameServerSessionsRequest) (response *SearchGameServerSessionsResponse, err error) {
    if request == nil {
        request = NewSearchGameServerSessionsRequest()
    }
    response = NewSearchGameServerSessionsResponse()
    err = c.Send(request, response)
    return
}

func NewStartGameServerSessionPlacementRequest() (request *StartGameServerSessionPlacementRequest) {
    request = &StartGameServerSessionPlacementRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gse", APIVersion, "StartGameServerSessionPlacement")
    return
}

func NewStartGameServerSessionPlacementResponse() (response *StartGameServerSessionPlacementResponse) {
    response = &StartGameServerSessionPlacementResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（StartGameServerSessionPlacement）用于开始放置游戏服务器会话
func (c *Client) StartGameServerSessionPlacement(request *StartGameServerSessionPlacementRequest) (response *StartGameServerSessionPlacementResponse, err error) {
    if request == nil {
        request = NewStartGameServerSessionPlacementRequest()
    }
    response = NewStartGameServerSessionPlacementResponse()
    err = c.Send(request, response)
    return
}

func NewStopGameServerSessionPlacementRequest() (request *StopGameServerSessionPlacementRequest) {
    request = &StopGameServerSessionPlacementRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gse", APIVersion, "StopGameServerSessionPlacement")
    return
}

func NewStopGameServerSessionPlacementResponse() (response *StopGameServerSessionPlacementResponse) {
    response = &StopGameServerSessionPlacementResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（StopGameServerSessionPlacement）用于停止放置游戏服务器会话
func (c *Client) StopGameServerSessionPlacement(request *StopGameServerSessionPlacementRequest) (response *StopGameServerSessionPlacementResponse, err error) {
    if request == nil {
        request = NewStopGameServerSessionPlacementRequest()
    }
    response = NewStopGameServerSessionPlacementResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateGameServerSessionRequest() (request *UpdateGameServerSessionRequest) {
    request = &UpdateGameServerSessionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gse", APIVersion, "UpdateGameServerSession")
    return
}

func NewUpdateGameServerSessionResponse() (response *UpdateGameServerSessionResponse) {
    response = &UpdateGameServerSessionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（UpdateGameServerSession）用于更新游戏服务器会话
func (c *Client) UpdateGameServerSession(request *UpdateGameServerSessionRequest) (response *UpdateGameServerSessionResponse, err error) {
    if request == nil {
        request = NewUpdateGameServerSessionRequest()
    }
    response = NewUpdateGameServerSessionResponse()
    err = c.Send(request, response)
    return
}
