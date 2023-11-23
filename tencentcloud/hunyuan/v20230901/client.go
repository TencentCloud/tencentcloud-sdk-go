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

package v20230901

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2023-09-01"

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


func NewChatProRequest() (request *ChatProRequest) {
    request = &ChatProRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("hunyuan", APIVersion, "ChatPro")
    
    
    return
}

func NewChatProResponse() (response *ChatProResponse) {
    response = &ChatProResponse{} 
    return

}

// ChatPro
// 腾讯混元大模型高级版是由腾讯研发的大语言模型，具备强大的中文创作能力，复杂语境下的逻辑推理能力，以及可靠的任务执行能力。本接口为SSE协议。
//
// 可能返回的错误码:
//  FAILEDOPERATION_ENGINEREQUESTTIMEOUT = "FailedOperation.EngineRequestTimeout"
//  FAILEDOPERATION_ENGINESERVERERROR = "FailedOperation.EngineServerError"
//  FAILEDOPERATION_ENGINESERVERLIMITEXCEEDED = "FailedOperation.EngineServerLimitExceeded"
//  FAILEDOPERATION_FREERESOURCEPACKEXHAUSTED = "FailedOperation.FreeResourcePackExhausted"
//  FAILEDOPERATION_RESOURCEPACKEXHAUSTED = "FailedOperation.ResourcePackExhausted"
//  FAILEDOPERATION_SERVICENOTACTIVATED = "FailedOperation.ServiceNotActivated"
//  FAILEDOPERATION_SERVICESTOP = "FailedOperation.ServiceStop"
//  FAILEDOPERATION_SERVICESTOPARREARS = "FailedOperation.ServiceStopArrears"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNSUPPORTEDOPERATION_NONWHITELISTACCOUNT = "UnsupportedOperation.NonWhitelistAccount"
func (c *Client) ChatPro(request *ChatProRequest) (response *ChatProResponse, err error) {
    return c.ChatProWithContext(context.Background(), request)
}

// ChatPro
// 腾讯混元大模型高级版是由腾讯研发的大语言模型，具备强大的中文创作能力，复杂语境下的逻辑推理能力，以及可靠的任务执行能力。本接口为SSE协议。
//
// 可能返回的错误码:
//  FAILEDOPERATION_ENGINEREQUESTTIMEOUT = "FailedOperation.EngineRequestTimeout"
//  FAILEDOPERATION_ENGINESERVERERROR = "FailedOperation.EngineServerError"
//  FAILEDOPERATION_ENGINESERVERLIMITEXCEEDED = "FailedOperation.EngineServerLimitExceeded"
//  FAILEDOPERATION_FREERESOURCEPACKEXHAUSTED = "FailedOperation.FreeResourcePackExhausted"
//  FAILEDOPERATION_RESOURCEPACKEXHAUSTED = "FailedOperation.ResourcePackExhausted"
//  FAILEDOPERATION_SERVICENOTACTIVATED = "FailedOperation.ServiceNotActivated"
//  FAILEDOPERATION_SERVICESTOP = "FailedOperation.ServiceStop"
//  FAILEDOPERATION_SERVICESTOPARREARS = "FailedOperation.ServiceStopArrears"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNSUPPORTEDOPERATION_NONWHITELISTACCOUNT = "UnsupportedOperation.NonWhitelistAccount"
func (c *Client) ChatProWithContext(ctx context.Context, request *ChatProRequest) (response *ChatProResponse, err error) {
    if request == nil {
        request = NewChatProRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ChatPro require credential")
    }

    request.SetContext(ctx)
    
    response = NewChatProResponse()
    err = c.Send(request, response)
    return
}

func NewChatStdRequest() (request *ChatStdRequest) {
    request = &ChatStdRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("hunyuan", APIVersion, "ChatStd")
    
    
    return
}

func NewChatStdResponse() (response *ChatStdResponse) {
    response = &ChatStdResponse{} 
    return

}

// ChatStd
// 腾讯混元大模型标准版是由腾讯研发的大语言模型，具备强大的中文创作能力，复杂语境下的逻辑推理能力，以及可靠的任务执行能力。本接口为SSE协议。
//
// 可能返回的错误码:
//  FAILEDOPERATION_ENGINEREQUESTTIMEOUT = "FailedOperation.EngineRequestTimeout"
//  FAILEDOPERATION_ENGINESERVERERROR = "FailedOperation.EngineServerError"
//  FAILEDOPERATION_ENGINESERVERLIMITEXCEEDED = "FailedOperation.EngineServerLimitExceeded"
//  FAILEDOPERATION_FREERESOURCEPACKEXHAUSTED = "FailedOperation.FreeResourcePackExhausted"
//  FAILEDOPERATION_RESOURCEPACKEXHAUSTED = "FailedOperation.ResourcePackExhausted"
//  FAILEDOPERATION_SERVICENOTACTIVATED = "FailedOperation.ServiceNotActivated"
//  FAILEDOPERATION_SERVICESTOP = "FailedOperation.ServiceStop"
//  FAILEDOPERATION_SERVICESTOPARREARS = "FailedOperation.ServiceStopArrears"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNSUPPORTEDOPERATION_NONWHITELISTACCOUNT = "UnsupportedOperation.NonWhitelistAccount"
func (c *Client) ChatStd(request *ChatStdRequest) (response *ChatStdResponse, err error) {
    return c.ChatStdWithContext(context.Background(), request)
}

// ChatStd
// 腾讯混元大模型标准版是由腾讯研发的大语言模型，具备强大的中文创作能力，复杂语境下的逻辑推理能力，以及可靠的任务执行能力。本接口为SSE协议。
//
// 可能返回的错误码:
//  FAILEDOPERATION_ENGINEREQUESTTIMEOUT = "FailedOperation.EngineRequestTimeout"
//  FAILEDOPERATION_ENGINESERVERERROR = "FailedOperation.EngineServerError"
//  FAILEDOPERATION_ENGINESERVERLIMITEXCEEDED = "FailedOperation.EngineServerLimitExceeded"
//  FAILEDOPERATION_FREERESOURCEPACKEXHAUSTED = "FailedOperation.FreeResourcePackExhausted"
//  FAILEDOPERATION_RESOURCEPACKEXHAUSTED = "FailedOperation.ResourcePackExhausted"
//  FAILEDOPERATION_SERVICENOTACTIVATED = "FailedOperation.ServiceNotActivated"
//  FAILEDOPERATION_SERVICESTOP = "FailedOperation.ServiceStop"
//  FAILEDOPERATION_SERVICESTOPARREARS = "FailedOperation.ServiceStopArrears"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNSUPPORTEDOPERATION_NONWHITELISTACCOUNT = "UnsupportedOperation.NonWhitelistAccount"
func (c *Client) ChatStdWithContext(ctx context.Context, request *ChatStdRequest) (response *ChatStdResponse, err error) {
    if request == nil {
        request = NewChatStdRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ChatStd require credential")
    }

    request.SetContext(ctx)
    
    response = NewChatStdResponse()
    err = c.Send(request, response)
    return
}

func NewGetTokenCountRequest() (request *GetTokenCountRequest) {
    request = &GetTokenCountRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("hunyuan", APIVersion, "GetTokenCount")
    
    
    return
}

func NewGetTokenCountResponse() (response *GetTokenCountResponse) {
    response = &GetTokenCountResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetTokenCount
// 该接口用于计算文本对应Token数、字符数。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) GetTokenCount(request *GetTokenCountRequest) (response *GetTokenCountResponse, err error) {
    return c.GetTokenCountWithContext(context.Background(), request)
}

// GetTokenCount
// 该接口用于计算文本对应Token数、字符数。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) GetTokenCountWithContext(ctx context.Context, request *GetTokenCountRequest) (response *GetTokenCountResponse, err error) {
    if request == nil {
        request = NewGetTokenCountRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetTokenCount require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetTokenCountResponse()
    err = c.Send(request, response)
    return
}
