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

package v20210519

import (
    "context"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2021-05-19"

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


func NewCreateCredentialRequest() (request *CreateCredentialRequest) {
    request = &CreateCredentialRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tdid", APIVersion, "CreateCredential")
    
    
    return
}

func NewCreateCredentialResponse() (response *CreateCredentialResponse) {
    response = &CreateCredentialResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateCredential
// 创建凭证
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_OPERATIONEXCEPTION = "FailedOperation.OperationException"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateCredential(request *CreateCredentialRequest) (response *CreateCredentialResponse, err error) {
    if request == nil {
        request = NewCreateCredentialRequest()
    }
    
    response = NewCreateCredentialResponse()
    err = c.Send(request, response)
    return
}

// CreateCredential
// 创建凭证
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_OPERATIONEXCEPTION = "FailedOperation.OperationException"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateCredentialWithContext(ctx context.Context, request *CreateCredentialRequest) (response *CreateCredentialResponse, err error) {
    if request == nil {
        request = NewCreateCredentialRequest()
    }
    request.SetContext(ctx)
    
    response = NewCreateCredentialResponse()
    err = c.Send(request, response)
    return
}

func NewCreateSelectiveCredentialRequest() (request *CreateSelectiveCredentialRequest) {
    request = &CreateSelectiveCredentialRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tdid", APIVersion, "CreateSelectiveCredential")
    
    
    return
}

func NewCreateSelectiveCredentialResponse() (response *CreateSelectiveCredentialResponse) {
    response = &CreateSelectiveCredentialResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateSelectiveCredential
// 创建选择性批露凭证
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_INVALIDAUTH = "FailedOperation.InvalidAuth"
//  FAILEDOPERATION_OPERATIONEXCEPTION = "FailedOperation.OperationException"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVEREXCEPTION = "InternalError.ServerException"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INTERNALERROR_SERVICEPANIC = "InternalError.ServicePanic"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateSelectiveCredential(request *CreateSelectiveCredentialRequest) (response *CreateSelectiveCredentialResponse, err error) {
    if request == nil {
        request = NewCreateSelectiveCredentialRequest()
    }
    
    response = NewCreateSelectiveCredentialResponse()
    err = c.Send(request, response)
    return
}

// CreateSelectiveCredential
// 创建选择性批露凭证
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_INVALIDAUTH = "FailedOperation.InvalidAuth"
//  FAILEDOPERATION_OPERATIONEXCEPTION = "FailedOperation.OperationException"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVEREXCEPTION = "InternalError.ServerException"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INTERNALERROR_SERVICEPANIC = "InternalError.ServicePanic"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateSelectiveCredentialWithContext(ctx context.Context, request *CreateSelectiveCredentialRequest) (response *CreateSelectiveCredentialResponse, err error) {
    if request == nil {
        request = NewCreateSelectiveCredentialRequest()
    }
    request.SetContext(ctx)
    
    response = NewCreateSelectiveCredentialResponse()
    err = c.Send(request, response)
    return
}

func NewCreateTDidRequest() (request *CreateTDidRequest) {
    request = &CreateTDidRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tdid", APIVersion, "CreateTDid")
    
    
    return
}

func NewCreateTDidResponse() (response *CreateTDidResponse) {
    response = &CreateTDidResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateTDid
// 创建机构DID
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_INVALIDAUTH = "FailedOperation.InvalidAuth"
//  FAILEDOPERATION_OPERATIONEXCEPTION = "FailedOperation.OperationException"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVEREXCEPTION = "InternalError.ServerException"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INTERNALERROR_SERVICEPANIC = "InternalError.ServicePanic"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
func (c *Client) CreateTDid(request *CreateTDidRequest) (response *CreateTDidResponse, err error) {
    if request == nil {
        request = NewCreateTDidRequest()
    }
    
    response = NewCreateTDidResponse()
    err = c.Send(request, response)
    return
}

// CreateTDid
// 创建机构DID
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_INVALIDAUTH = "FailedOperation.InvalidAuth"
//  FAILEDOPERATION_OPERATIONEXCEPTION = "FailedOperation.OperationException"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVEREXCEPTION = "InternalError.ServerException"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INTERNALERROR_SERVICEPANIC = "InternalError.ServicePanic"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
func (c *Client) CreateTDidWithContext(ctx context.Context, request *CreateTDidRequest) (response *CreateTDidResponse, err error) {
    if request == nil {
        request = NewCreateTDidRequest()
    }
    request.SetContext(ctx)
    
    response = NewCreateTDidResponse()
    err = c.Send(request, response)
    return
}

func NewVerifyCredentialRequest() (request *VerifyCredentialRequest) {
    request = &VerifyCredentialRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tdid", APIVersion, "VerifyCredential")
    
    
    return
}

func NewVerifyCredentialResponse() (response *VerifyCredentialResponse) {
    response = &VerifyCredentialResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// VerifyCredential
// 验证凭证
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_INVALIDAUTH = "FailedOperation.InvalidAuth"
//  FAILEDOPERATION_OPERATIONEXCEPTION = "FailedOperation.OperationException"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVEREXCEPTION = "InternalError.ServerException"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INTERNALERROR_SERVICEPANIC = "InternalError.ServicePanic"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) VerifyCredential(request *VerifyCredentialRequest) (response *VerifyCredentialResponse, err error) {
    if request == nil {
        request = NewVerifyCredentialRequest()
    }
    
    response = NewVerifyCredentialResponse()
    err = c.Send(request, response)
    return
}

// VerifyCredential
// 验证凭证
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_INVALIDAUTH = "FailedOperation.InvalidAuth"
//  FAILEDOPERATION_OPERATIONEXCEPTION = "FailedOperation.OperationException"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVEREXCEPTION = "InternalError.ServerException"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INTERNALERROR_SERVICEPANIC = "InternalError.ServicePanic"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) VerifyCredentialWithContext(ctx context.Context, request *VerifyCredentialRequest) (response *VerifyCredentialResponse, err error) {
    if request == nil {
        request = NewVerifyCredentialRequest()
    }
    request.SetContext(ctx)
    
    response = NewVerifyCredentialResponse()
    err = c.Send(request, response)
    return
}
