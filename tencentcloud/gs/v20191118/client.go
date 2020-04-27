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

package v20191118

import (
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2019-11-18"

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


func NewCreateSessionRequest() (request *CreateSessionRequest) {
    request = &CreateSessionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gs", APIVersion, "CreateSession")
    return
}

func NewCreateSessionResponse() (response *CreateSessionResponse) {
    response = &CreateSessionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建会话
func (c *Client) CreateSession(request *CreateSessionRequest) (response *CreateSessionResponse, err error) {
    if request == nil {
        request = NewCreateSessionRequest()
    }
    response = NewCreateSessionResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeWorkersRequest() (request *DescribeWorkersRequest) {
    request = &DescribeWorkersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gs", APIVersion, "DescribeWorkers")
    return
}

func NewDescribeWorkersResponse() (response *DescribeWorkersResponse) {
    response = &DescribeWorkersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询空闲机器数量
func (c *Client) DescribeWorkers(request *DescribeWorkersRequest) (response *DescribeWorkersResponse, err error) {
    if request == nil {
        request = NewDescribeWorkersRequest()
    }
    response = NewDescribeWorkersResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeWorkersInfoRequest() (request *DescribeWorkersInfoRequest) {
    request = &DescribeWorkersInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gs", APIVersion, "DescribeWorkersInfo")
    return
}

func NewDescribeWorkersInfoResponse() (response *DescribeWorkersInfoResponse) {
    response = &DescribeWorkersInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取机器信息
func (c *Client) DescribeWorkersInfo(request *DescribeWorkersInfoRequest) (response *DescribeWorkersInfoResponse, err error) {
    if request == nil {
        request = NewDescribeWorkersInfoRequest()
    }
    response = NewDescribeWorkersInfoResponse()
    err = c.Send(request, response)
    return
}

func NewEnterQueueRequest() (request *EnterQueueRequest) {
    request = &EnterQueueRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gs", APIVersion, "EnterQueue")
    return
}

func NewEnterQueueResponse() (response *EnterQueueResponse) {
    response = &EnterQueueResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 进入排队锁定机器
func (c *Client) EnterQueue(request *EnterQueueRequest) (response *EnterQueueResponse, err error) {
    if request == nil {
        request = NewEnterQueueRequest()
    }
    response = NewEnterQueueResponse()
    err = c.Send(request, response)
    return
}

func NewModifyWorkersRequest() (request *ModifyWorkersRequest) {
    request = &ModifyWorkersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gs", APIVersion, "ModifyWorkers")
    return
}

func NewModifyWorkersResponse() (response *ModifyWorkersResponse) {
    response = &ModifyWorkersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改机器信息
func (c *Client) ModifyWorkers(request *ModifyWorkersRequest) (response *ModifyWorkersResponse, err error) {
    if request == nil {
        request = NewModifyWorkersRequest()
    }
    response = NewModifyWorkersResponse()
    err = c.Send(request, response)
    return
}

func NewQuitQueueRequest() (request *QuitQueueRequest) {
    request = &QuitQueueRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gs", APIVersion, "QuitQueue")
    return
}

func NewQuitQueueResponse() (response *QuitQueueResponse) {
    response = &QuitQueueResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 退出排队
func (c *Client) QuitQueue(request *QuitQueueRequest) (response *QuitQueueResponse, err error) {
    if request == nil {
        request = NewQuitQueueRequest()
    }
    response = NewQuitQueueResponse()
    err = c.Send(request, response)
    return
}

func NewStopGameRequest() (request *StopGameRequest) {
    request = &StopGameRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gs", APIVersion, "StopGame")
    return
}

func NewStopGameResponse() (response *StopGameResponse) {
    response = &StopGameResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 强制退出游戏
func (c *Client) StopGame(request *StopGameRequest) (response *StopGameResponse, err error) {
    if request == nil {
        request = NewStopGameRequest()
    }
    response = NewStopGameResponse()
    err = c.Send(request, response)
    return
}

func NewTrylockWorkerRequest() (request *TrylockWorkerRequest) {
    request = &TrylockWorkerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gs", APIVersion, "TrylockWorker")
    return
}

func NewTrylockWorkerResponse() (response *TrylockWorkerResponse) {
    response = &TrylockWorkerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 尝试锁定机器
func (c *Client) TrylockWorker(request *TrylockWorkerRequest) (response *TrylockWorkerResponse, err error) {
    if request == nil {
        request = NewTrylockWorkerRequest()
    }
    response = NewTrylockWorkerResponse()
    err = c.Send(request, response)
    return
}
