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

package v20190313

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2019-03-13"

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


func NewCreateSessionRequest() (request *CreateSessionRequest) {
    request = &CreateSessionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bizlive", APIVersion, "CreateSession")
    
    
    return
}

func NewCreateSessionResponse() (response *CreateSessionResponse) {
    response = &CreateSessionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateSession
// 创建会话
//
// 可能返回的错误码:
//  FAILEDOPERATION_LACKBANDWIDTH = "FailedOperation.LackBandwidth"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_JSONPARSEERROR = "InvalidParameter.JsonParseError"
//  RESOURCENOTFOUND_NOIDLE = "ResourceNotFound.NoIdle"
func (c *Client) CreateSession(request *CreateSessionRequest) (response *CreateSessionResponse, err error) {
    return c.CreateSessionWithContext(context.Background(), request)
}

// CreateSession
// 创建会话
//
// 可能返回的错误码:
//  FAILEDOPERATION_LACKBANDWIDTH = "FailedOperation.LackBandwidth"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_JSONPARSEERROR = "InvalidParameter.JsonParseError"
//  RESOURCENOTFOUND_NOIDLE = "ResourceNotFound.NoIdle"
func (c *Client) CreateSessionWithContext(ctx context.Context, request *CreateSessionRequest) (response *CreateSessionResponse, err error) {
    if request == nil {
        request = NewCreateSessionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateSession require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateSessionResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeStreamPlayInfoListRequest() (request *DescribeStreamPlayInfoListRequest) {
    request = &DescribeStreamPlayInfoListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bizlive", APIVersion, "DescribeStreamPlayInfoList")
    
    
    return
}

func NewDescribeStreamPlayInfoListResponse() (response *DescribeStreamPlayInfoListResponse) {
    response = &DescribeStreamPlayInfoListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeStreamPlayInfoList
// 查询播放数据，支持按流名称查询详细播放数据，也可按播放域名查询详细总数据。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeStreamPlayInfoList(request *DescribeStreamPlayInfoListRequest) (response *DescribeStreamPlayInfoListResponse, err error) {
    return c.DescribeStreamPlayInfoListWithContext(context.Background(), request)
}

// DescribeStreamPlayInfoList
// 查询播放数据，支持按流名称查询详细播放数据，也可按播放域名查询详细总数据。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeStreamPlayInfoListWithContext(ctx context.Context, request *DescribeStreamPlayInfoListRequest) (response *DescribeStreamPlayInfoListResponse, err error) {
    if request == nil {
        request = NewDescribeStreamPlayInfoListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeStreamPlayInfoList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeStreamPlayInfoListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeWorkersRequest() (request *DescribeWorkersRequest) {
    request = &DescribeWorkersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bizlive", APIVersion, "DescribeWorkers")
    
    
    return
}

func NewDescribeWorkersResponse() (response *DescribeWorkersResponse) {
    response = &DescribeWorkersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeWorkers
// 查询空闲机器数量
//
// 可能返回的错误码:
//  INVALIDPARAMETER_JSONPARSEERROR = "InvalidParameter.JsonParseError"
func (c *Client) DescribeWorkers(request *DescribeWorkersRequest) (response *DescribeWorkersResponse, err error) {
    return c.DescribeWorkersWithContext(context.Background(), request)
}

// DescribeWorkers
// 查询空闲机器数量
//
// 可能返回的错误码:
//  INVALIDPARAMETER_JSONPARSEERROR = "InvalidParameter.JsonParseError"
func (c *Client) DescribeWorkersWithContext(ctx context.Context, request *DescribeWorkersRequest) (response *DescribeWorkersResponse, err error) {
    if request == nil {
        request = NewDescribeWorkersRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeWorkers require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeWorkersResponse()
    err = c.Send(request, response)
    return
}

func NewForbidLiveStreamRequest() (request *ForbidLiveStreamRequest) {
    request = &ForbidLiveStreamRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bizlive", APIVersion, "ForbidLiveStream")
    
    
    return
}

func NewForbidLiveStreamResponse() (response *ForbidLiveStreamResponse) {
    response = &ForbidLiveStreamResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ForbidLiveStream
// 禁止某条流的推送，可以预设某个时刻将流恢复。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CALLOTHERSVRERROR = "InternalError.CallOtherSvrError"
//  INTERNALERROR_CONFIGNOTEXIST = "InternalError.ConfigNotExist"
//  INTERNALERROR_GETBIZIDERROR = "InternalError.GetBizidError"
//  INTERNALERROR_GETSTREAMINFOERROR = "InternalError.GetStreamInfoError"
//  INTERNALERROR_GETUPSTREAMINFOERROR = "InternalError.GetUpstreamInfoError"
//  INTERNALERROR_NOTPERMMITOPERAT = "InternalError.NotPermmitOperat"
//  INTERNALERROR_STREAMSTATUSERROR = "InternalError.StreamStatusError"
//  INTERNALERROR_UPDATEDATAERROR = "InternalError.UpdateDataError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ForbidLiveStream(request *ForbidLiveStreamRequest) (response *ForbidLiveStreamResponse, err error) {
    return c.ForbidLiveStreamWithContext(context.Background(), request)
}

// ForbidLiveStream
// 禁止某条流的推送，可以预设某个时刻将流恢复。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CALLOTHERSVRERROR = "InternalError.CallOtherSvrError"
//  INTERNALERROR_CONFIGNOTEXIST = "InternalError.ConfigNotExist"
//  INTERNALERROR_GETBIZIDERROR = "InternalError.GetBizidError"
//  INTERNALERROR_GETSTREAMINFOERROR = "InternalError.GetStreamInfoError"
//  INTERNALERROR_GETUPSTREAMINFOERROR = "InternalError.GetUpstreamInfoError"
//  INTERNALERROR_NOTPERMMITOPERAT = "InternalError.NotPermmitOperat"
//  INTERNALERROR_STREAMSTATUSERROR = "InternalError.StreamStatusError"
//  INTERNALERROR_UPDATEDATAERROR = "InternalError.UpdateDataError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ForbidLiveStreamWithContext(ctx context.Context, request *ForbidLiveStreamRequest) (response *ForbidLiveStreamResponse, err error) {
    if request == nil {
        request = NewForbidLiveStreamRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ForbidLiveStream require credential")
    }

    request.SetContext(ctx)
    
    response = NewForbidLiveStreamResponse()
    err = c.Send(request, response)
    return
}

func NewRegisterIMRequest() (request *RegisterIMRequest) {
    request = &RegisterIMRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bizlive", APIVersion, "RegisterIM")
    
    
    return
}

func NewRegisterIMResponse() (response *RegisterIMResponse) {
    response = &RegisterIMResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// RegisterIM
// 注册聊天室
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_JSONPARSEERROR = "InvalidParameter.JsonParseError"
//  LIMITEXCEEDED_NOIMABILITY = "LimitExceeded.NoIMAbility"
func (c *Client) RegisterIM(request *RegisterIMRequest) (response *RegisterIMResponse, err error) {
    return c.RegisterIMWithContext(context.Background(), request)
}

// RegisterIM
// 注册聊天室
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_JSONPARSEERROR = "InvalidParameter.JsonParseError"
//  LIMITEXCEEDED_NOIMABILITY = "LimitExceeded.NoIMAbility"
func (c *Client) RegisterIMWithContext(ctx context.Context, request *RegisterIMRequest) (response *RegisterIMResponse, err error) {
    if request == nil {
        request = NewRegisterIMRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RegisterIM require credential")
    }

    request.SetContext(ctx)
    
    response = NewRegisterIMResponse()
    err = c.Send(request, response)
    return
}

func NewStopGameRequest() (request *StopGameRequest) {
    request = &StopGameRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bizlive", APIVersion, "StopGame")
    
    
    return
}

func NewStopGameResponse() (response *StopGameResponse) {
    response = &StopGameResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// StopGame
// 强制退出游戏
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_JSONPARSEERROR = "InvalidParameter.JsonParseError"
func (c *Client) StopGame(request *StopGameRequest) (response *StopGameResponse, err error) {
    return c.StopGameWithContext(context.Background(), request)
}

// StopGame
// 强制退出游戏
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_JSONPARSEERROR = "InvalidParameter.JsonParseError"
func (c *Client) StopGameWithContext(ctx context.Context, request *StopGameRequest) (response *StopGameResponse, err error) {
    if request == nil {
        request = NewStopGameRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("StopGame require credential")
    }

    request.SetContext(ctx)
    
    response = NewStopGameResponse()
    err = c.Send(request, response)
    return
}
