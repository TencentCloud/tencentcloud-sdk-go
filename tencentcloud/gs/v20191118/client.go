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
//  UNSUPPORTEDOPERATION_STOPPING = "UnsupportedOperation.Stopping"
func (c *Client) CreateSession(request *CreateSessionRequest) (response *CreateSessionResponse, err error) {
    if request == nil {
        request = NewCreateSessionRequest()
    }
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
// 获取实例总数和运行数
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_JSONPARSEERROR = "InvalidParameter.JsonParseError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeInstancesCount(request *DescribeInstancesCountRequest) (response *DescribeInstancesCountResponse, err error) {
    if request == nil {
        request = NewDescribeInstancesCountRequest()
    }
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
//  FAILEDOPERATION_SLOWDOWN = "FailedOperation.SlowDown"
//  FAILEDOPERATION_TOOFREQUENTLY = "FailedOperation.TooFrequently"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_JSONPARSEERROR = "InvalidParameter.JsonParseError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND_SESSIONNOTFOUND = "ResourceNotFound.SessionNotFound"
func (c *Client) SaveGameArchive(request *SaveGameArchiveRequest) (response *SaveGameArchiveResponse, err error) {
    if request == nil {
        request = NewSaveGameArchiveRequest()
    }
    response = NewSaveGameArchiveResponse()
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
//  FAILEDOPERATION_PROCESSTIMEOUT = "FailedOperation.ProcessTimeout"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_JSONPARSEERROR = "InvalidParameter.JsonParseError"
func (c *Client) StopGame(request *StopGameRequest) (response *StopGameResponse, err error) {
    if request == nil {
        request = NewStopGameRequest()
    }
    response = NewStopGameResponse()
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
//  FAILEDOPERATION_SLOWDOWN = "FailedOperation.SlowDown"
//  FAILEDOPERATION_TOOFREQUENTLY = "FailedOperation.TooFrequently"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_JSONPARSEERROR = "InvalidParameter.JsonParseError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND_SESSIONNOTFOUND = "ResourceNotFound.SessionNotFound"
func (c *Client) SwitchGameArchive(request *SwitchGameArchiveRequest) (response *SwitchGameArchiveResponse, err error) {
    if request == nil {
        request = NewSwitchGameArchiveRequest()
    }
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
//  FAILEDOPERATION_PROCESSTIMEOUT = "FailedOperation.ProcessTimeout"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_JSONPARSEERROR = "InvalidParameter.JsonParseError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_NOIDLE = "ResourceNotFound.NoIdle"
//  RESOURCEUNAVAILABLE_INITIALIZATION = "ResourceUnavailable.Initialization"
//  UNSUPPORTEDOPERATION_STOPPING = "UnsupportedOperation.Stopping"
func (c *Client) TrylockWorker(request *TrylockWorkerRequest) (response *TrylockWorkerResponse, err error) {
    if request == nil {
        request = NewTrylockWorkerRequest()
    }
    response = NewTrylockWorkerResponse()
    err = c.Send(request, response)
    return
}
