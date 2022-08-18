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

package v20220110

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2022-01-10"

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


func NewApplyConcurrentRequest() (request *ApplyConcurrentRequest) {
    request = &ApplyConcurrentRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("car", APIVersion, "ApplyConcurrent")
    
    
    return
}

func NewApplyConcurrentResponse() (response *ApplyConcurrentResponse) {
    response = &ApplyConcurrentResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ApplyConcurrent
// 申请并发
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
func (c *Client) ApplyConcurrent(request *ApplyConcurrentRequest) (response *ApplyConcurrentResponse, err error) {
    return c.ApplyConcurrentWithContext(context.Background(), request)
}

// ApplyConcurrent
// 申请并发
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
func (c *Client) ApplyConcurrentWithContext(ctx context.Context, request *ApplyConcurrentRequest) (response *ApplyConcurrentResponse, err error) {
    if request == nil {
        request = NewApplyConcurrentRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ApplyConcurrent require credential")
    }

    request.SetContext(ctx)
    
    response = NewApplyConcurrentResponse()
    err = c.Send(request, response)
    return
}

func NewCreateSessionRequest() (request *CreateSessionRequest) {
    request = &CreateSessionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("car", APIVersion, "CreateSession")
    
    
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
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_JSONPARSEERROR = "InvalidParameter.JsonParseError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
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
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_JSONPARSEERROR = "InvalidParameter.JsonParseError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
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

func NewDestroySessionRequest() (request *DestroySessionRequest) {
    request = &DestroySessionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("car", APIVersion, "DestroySession")
    
    
    return
}

func NewDestroySessionResponse() (response *DestroySessionResponse) {
    response = &DestroySessionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DestroySession
// 销毁会话
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PROCESSTIMEOUT = "FailedOperation.ProcessTimeout"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_JSONPARSEERROR = "InvalidParameter.JsonParseError"
//  RESOURCEUNAVAILABLE_ACCESSFAILED = "ResourceUnavailable.AccessFailed"
func (c *Client) DestroySession(request *DestroySessionRequest) (response *DestroySessionResponse, err error) {
    return c.DestroySessionWithContext(context.Background(), request)
}

// DestroySession
// 销毁会话
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PROCESSTIMEOUT = "FailedOperation.ProcessTimeout"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_JSONPARSEERROR = "InvalidParameter.JsonParseError"
//  RESOURCEUNAVAILABLE_ACCESSFAILED = "ResourceUnavailable.AccessFailed"
func (c *Client) DestroySessionWithContext(ctx context.Context, request *DestroySessionRequest) (response *DestroySessionResponse, err error) {
    if request == nil {
        request = NewDestroySessionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DestroySession require credential")
    }

    request.SetContext(ctx)
    
    response = NewDestroySessionResponse()
    err = c.Send(request, response)
    return
}
