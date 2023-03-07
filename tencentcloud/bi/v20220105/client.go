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

package v20220105

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2022-01-05"

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


func NewApplyEmbedIntervalRequest() (request *ApplyEmbedIntervalRequest) {
    request = &ApplyEmbedIntervalRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bi", APIVersion, "ApplyEmbedInterval")
    
    
    return
}

func NewApplyEmbedIntervalResponse() (response *ApplyEmbedIntervalResponse) {
    response = &ApplyEmbedIntervalResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ApplyEmbedInterval
// 申请延长Token可用时间接口-强鉴权
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNAL = "InternalError.Internal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_EMBED = "InvalidParameter.Embed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_MISSINGPARAM = "MissingParameter.MissingParam"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_AUTHORIZE = "UnauthorizedOperation.Authorize"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_BIERROR = "UnsupportedOperation.BIError"
func (c *Client) ApplyEmbedInterval(request *ApplyEmbedIntervalRequest) (response *ApplyEmbedIntervalResponse, err error) {
    return c.ApplyEmbedIntervalWithContext(context.Background(), request)
}

// ApplyEmbedInterval
// 申请延长Token可用时间接口-强鉴权
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNAL = "InternalError.Internal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_EMBED = "InvalidParameter.Embed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_MISSINGPARAM = "MissingParameter.MissingParam"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_AUTHORIZE = "UnauthorizedOperation.Authorize"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_BIERROR = "UnsupportedOperation.BIError"
func (c *Client) ApplyEmbedIntervalWithContext(ctx context.Context, request *ApplyEmbedIntervalRequest) (response *ApplyEmbedIntervalResponse, err error) {
    if request == nil {
        request = NewApplyEmbedIntervalRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ApplyEmbedInterval require credential")
    }

    request.SetContext(ctx)
    
    response = NewApplyEmbedIntervalResponse()
    err = c.Send(request, response)
    return
}

func NewCreateEmbedTokenRequest() (request *CreateEmbedTokenRequest) {
    request = &CreateEmbedTokenRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bi", APIVersion, "CreateEmbedToken")
    
    
    return
}

func NewCreateEmbedTokenResponse() (response *CreateEmbedTokenResponse) {
    response = &CreateEmbedTokenResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateEmbedToken
// 创建嵌出报表-强鉴权
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNAL = "InternalError.Internal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_EMBED = "InvalidParameter.Embed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_MISSINGPARAM = "MissingParameter.MissingParam"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_AUTHORIZE = "UnauthorizedOperation.Authorize"
//  UNAUTHORIZEDOPERATION_USERNOTEXIST = "UnauthorizedOperation.UserNotExist"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_BIERROR = "UnsupportedOperation.BIError"
func (c *Client) CreateEmbedToken(request *CreateEmbedTokenRequest) (response *CreateEmbedTokenResponse, err error) {
    return c.CreateEmbedTokenWithContext(context.Background(), request)
}

// CreateEmbedToken
// 创建嵌出报表-强鉴权
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNAL = "InternalError.Internal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_EMBED = "InvalidParameter.Embed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_MISSINGPARAM = "MissingParameter.MissingParam"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_AUTHORIZE = "UnauthorizedOperation.Authorize"
//  UNAUTHORIZEDOPERATION_USERNOTEXIST = "UnauthorizedOperation.UserNotExist"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_BIERROR = "UnsupportedOperation.BIError"
func (c *Client) CreateEmbedTokenWithContext(ctx context.Context, request *CreateEmbedTokenRequest) (response *CreateEmbedTokenResponse, err error) {
    if request == nil {
        request = NewCreateEmbedTokenRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateEmbedToken require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateEmbedTokenResponse()
    err = c.Send(request, response)
    return
}
