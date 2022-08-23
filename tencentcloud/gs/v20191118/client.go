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
    "context"
    "errors"
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
    
    request.Init().WithApiInfo("gs", APIVersion, "CreateSession")
    
    
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
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_LOCKTIMEOUT = "FailedOperation.LockTimeout"
//  FAILEDOPERATION_PROCESSTIMEOUT = "FailedOperation.ProcessTimeout"
//  FAILEDOPERATION_SLOWDOWN = "FailedOperation.SlowDown"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  FAILEDOPERATION_TOOFREQUENTLY = "FailedOperation.TooFrequently"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_JSONPARSEERROR = "InvalidParameter.JsonParseError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_ROLE = "LimitExceeded.Role"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND_NOIDLE = "ResourceNotFound.NoIdle"
//  RESOURCENOTFOUND_SESSIONNOTFOUND = "ResourceNotFound.SessionNotFound"
//  RESOURCEUNAVAILABLE_INITIALIZATION = "ResourceUnavailable.Initialization"
//  UNSUPPORTEDOPERATION_STOPPING = "UnsupportedOperation.Stopping"
func (c *Client) CreateSession(request *CreateSessionRequest) (response *CreateSessionResponse, err error) {
    return c.CreateSessionWithContext(context.Background(), request)
}

// CreateSession
// 创建会话
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_LOCKTIMEOUT = "FailedOperation.LockTimeout"
//  FAILEDOPERATION_PROCESSTIMEOUT = "FailedOperation.ProcessTimeout"
//  FAILEDOPERATION_SLOWDOWN = "FailedOperation.SlowDown"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  FAILEDOPERATION_TOOFREQUENTLY = "FailedOperation.TooFrequently"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_JSONPARSEERROR = "InvalidParameter.JsonParseError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_ROLE = "LimitExceeded.Role"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND_NOIDLE = "ResourceNotFound.NoIdle"
//  RESOURCENOTFOUND_SESSIONNOTFOUND = "ResourceNotFound.SessionNotFound"
//  RESOURCEUNAVAILABLE_INITIALIZATION = "ResourceUnavailable.Initialization"
//  UNSUPPORTEDOPERATION_STOPPING = "UnsupportedOperation.Stopping"
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

func NewDescribeInstancesCountRequest() (request *DescribeInstancesCountRequest) {
    request = &DescribeInstancesCountRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("gs", APIVersion, "DescribeInstancesCount")
    
    
    return
}

func NewDescribeInstancesCountResponse() (response *DescribeInstancesCountResponse) {
    response = &DescribeInstancesCountResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeInstancesCount
// 获取并发总数和运行数
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_JSONPARSEERROR = "InvalidParameter.JsonParseError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeInstancesCount(request *DescribeInstancesCountRequest) (response *DescribeInstancesCountResponse, err error) {
    return c.DescribeInstancesCountWithContext(context.Background(), request)
}

// DescribeInstancesCount
// 获取并发总数和运行数
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_JSONPARSEERROR = "InvalidParameter.JsonParseError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeInstancesCountWithContext(ctx context.Context, request *DescribeInstancesCountRequest) (response *DescribeInstancesCountResponse, err error) {
    if request == nil {
        request = NewDescribeInstancesCountRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstancesCount require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstancesCountResponse()
    err = c.Send(request, response)
    return
}

func NewSaveGameArchiveRequest() (request *SaveGameArchiveRequest) {
    request = &SaveGameArchiveRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("gs", APIVersion, "SaveGameArchive")
    
    
    return
}

func NewSaveGameArchiveResponse() (response *SaveGameArchiveResponse) {
    response = &SaveGameArchiveResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// SaveGameArchive
// 保存游戏存档
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_SLOWDOWN = "FailedOperation.SlowDown"
//  FAILEDOPERATION_TOOFREQUENTLY = "FailedOperation.TooFrequently"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_JSONPARSEERROR = "InvalidParameter.JsonParseError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND_SESSIONNOTFOUND = "ResourceNotFound.SessionNotFound"
func (c *Client) SaveGameArchive(request *SaveGameArchiveRequest) (response *SaveGameArchiveResponse, err error) {
    return c.SaveGameArchiveWithContext(context.Background(), request)
}

// SaveGameArchive
// 保存游戏存档
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_SLOWDOWN = "FailedOperation.SlowDown"
//  FAILEDOPERATION_TOOFREQUENTLY = "FailedOperation.TooFrequently"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_JSONPARSEERROR = "InvalidParameter.JsonParseError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND_SESSIONNOTFOUND = "ResourceNotFound.SessionNotFound"
func (c *Client) SaveGameArchiveWithContext(ctx context.Context, request *SaveGameArchiveRequest) (response *SaveGameArchiveResponse, err error) {
    if request == nil {
        request = NewSaveGameArchiveRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SaveGameArchive require credential")
    }

    request.SetContext(ctx)
    
    response = NewSaveGameArchiveResponse()
    err = c.Send(request, response)
    return
}

func NewStartPublishStreamRequest() (request *StartPublishStreamRequest) {
    request = &StartPublishStreamRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("gs", APIVersion, "StartPublishStream")
    
    
    return
}

func NewStartPublishStreamResponse() (response *StartPublishStreamResponse) {
    response = &StartPublishStreamResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// StartPublishStream
// 开始云端推流
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_JSONPARSEERROR = "InvalidParameter.JsonParseError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND_SESSIONNOTFOUND = "ResourceNotFound.SessionNotFound"
//  UNSUPPORTEDOPERATION_NOTRUNNING = "UnsupportedOperation.NotRunning"
func (c *Client) StartPublishStream(request *StartPublishStreamRequest) (response *StartPublishStreamResponse, err error) {
    return c.StartPublishStreamWithContext(context.Background(), request)
}

// StartPublishStream
// 开始云端推流
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_JSONPARSEERROR = "InvalidParameter.JsonParseError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND_SESSIONNOTFOUND = "ResourceNotFound.SessionNotFound"
//  UNSUPPORTEDOPERATION_NOTRUNNING = "UnsupportedOperation.NotRunning"
func (c *Client) StartPublishStreamWithContext(ctx context.Context, request *StartPublishStreamRequest) (response *StartPublishStreamResponse, err error) {
    if request == nil {
        request = NewStartPublishStreamRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("StartPublishStream require credential")
    }

    request.SetContext(ctx)
    
    response = NewStartPublishStreamResponse()
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

// StopGame
// 强制退出游戏
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PROCESSTIMEOUT = "FailedOperation.ProcessTimeout"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_JSONPARSEERROR = "InvalidParameter.JsonParseError"
//  RESOURCEUNAVAILABLE_ACCESSFAILED = "ResourceUnavailable.AccessFailed"
func (c *Client) StopGame(request *StopGameRequest) (response *StopGameResponse, err error) {
    return c.StopGameWithContext(context.Background(), request)
}

// StopGame
// 强制退出游戏
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PROCESSTIMEOUT = "FailedOperation.ProcessTimeout"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_JSONPARSEERROR = "InvalidParameter.JsonParseError"
//  RESOURCEUNAVAILABLE_ACCESSFAILED = "ResourceUnavailable.AccessFailed"
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

func NewStopPublishStreamRequest() (request *StopPublishStreamRequest) {
    request = &StopPublishStreamRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("gs", APIVersion, "StopPublishStream")
    
    
    return
}

func NewStopPublishStreamResponse() (response *StopPublishStreamResponse) {
    response = &StopPublishStreamResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// StopPublishStream
// 停止云端推流
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_JSONPARSEERROR = "InvalidParameter.JsonParseError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND_SESSIONNOTFOUND = "ResourceNotFound.SessionNotFound"
//  UNSUPPORTEDOPERATION_NOTRUNNING = "UnsupportedOperation.NotRunning"
func (c *Client) StopPublishStream(request *StopPublishStreamRequest) (response *StopPublishStreamResponse, err error) {
    return c.StopPublishStreamWithContext(context.Background(), request)
}

// StopPublishStream
// 停止云端推流
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_JSONPARSEERROR = "InvalidParameter.JsonParseError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND_SESSIONNOTFOUND = "ResourceNotFound.SessionNotFound"
//  UNSUPPORTEDOPERATION_NOTRUNNING = "UnsupportedOperation.NotRunning"
func (c *Client) StopPublishStreamWithContext(ctx context.Context, request *StopPublishStreamRequest) (response *StopPublishStreamResponse, err error) {
    if request == nil {
        request = NewStopPublishStreamRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("StopPublishStream require credential")
    }

    request.SetContext(ctx)
    
    response = NewStopPublishStreamResponse()
    err = c.Send(request, response)
    return
}

func NewSwitchGameArchiveRequest() (request *SwitchGameArchiveRequest) {
    request = &SwitchGameArchiveRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("gs", APIVersion, "SwitchGameArchive")
    
    
    return
}

func NewSwitchGameArchiveResponse() (response *SwitchGameArchiveResponse) {
    response = &SwitchGameArchiveResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// SwitchGameArchive
// 切换游戏存档
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_SLOWDOWN = "FailedOperation.SlowDown"
//  FAILEDOPERATION_TOOFREQUENTLY = "FailedOperation.TooFrequently"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_JSONPARSEERROR = "InvalidParameter.JsonParseError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND_SESSIONNOTFOUND = "ResourceNotFound.SessionNotFound"
func (c *Client) SwitchGameArchive(request *SwitchGameArchiveRequest) (response *SwitchGameArchiveResponse, err error) {
    return c.SwitchGameArchiveWithContext(context.Background(), request)
}

// SwitchGameArchive
// 切换游戏存档
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_SLOWDOWN = "FailedOperation.SlowDown"
//  FAILEDOPERATION_TOOFREQUENTLY = "FailedOperation.TooFrequently"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_JSONPARSEERROR = "InvalidParameter.JsonParseError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND_SESSIONNOTFOUND = "ResourceNotFound.SessionNotFound"
func (c *Client) SwitchGameArchiveWithContext(ctx context.Context, request *SwitchGameArchiveRequest) (response *SwitchGameArchiveResponse, err error) {
    if request == nil {
        request = NewSwitchGameArchiveRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SwitchGameArchive require credential")
    }

    request.SetContext(ctx)
    
    response = NewSwitchGameArchiveResponse()
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

// TrylockWorker
// 尝试锁定机器
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PROCESSTIMEOUT = "FailedOperation.ProcessTimeout"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_JSONPARSEERROR = "InvalidParameter.JsonParseError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_NOIDLE = "ResourceNotFound.NoIdle"
//  RESOURCEUNAVAILABLE_INITIALIZATION = "ResourceUnavailable.Initialization"
//  UNSUPPORTEDOPERATION_STOPPING = "UnsupportedOperation.Stopping"
func (c *Client) TrylockWorker(request *TrylockWorkerRequest) (response *TrylockWorkerResponse, err error) {
    return c.TrylockWorkerWithContext(context.Background(), request)
}

// TrylockWorker
// 尝试锁定机器
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PROCESSTIMEOUT = "FailedOperation.ProcessTimeout"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_JSONPARSEERROR = "InvalidParameter.JsonParseError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_NOIDLE = "ResourceNotFound.NoIdle"
//  RESOURCEUNAVAILABLE_INITIALIZATION = "ResourceUnavailable.Initialization"
//  UNSUPPORTEDOPERATION_STOPPING = "UnsupportedOperation.Stopping"
func (c *Client) TrylockWorkerWithContext(ctx context.Context, request *TrylockWorkerRequest) (response *TrylockWorkerResponse, err error) {
    if request == nil {
        request = NewTrylockWorkerRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("TrylockWorker require credential")
    }

    request.SetContext(ctx)
    
    response = NewTrylockWorkerResponse()
    err = c.Send(request, response)
    return
}
