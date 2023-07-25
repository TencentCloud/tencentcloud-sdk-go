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
    "errors"
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


func NewCheckChainRequest() (request *CheckChainRequest) {
    request = &CheckChainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdid", APIVersion, "CheckChain")
    
    
    return
}

func NewCheckChainResponse() (response *CheckChainResponse) {
    response = &CheckChainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CheckChain
// 该接口不再使用
//
// 
//
// 检查区块链信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
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
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CheckChain(request *CheckChainRequest) (response *CheckChainResponse, err error) {
    return c.CheckChainWithContext(context.Background(), request)
}

// CheckChain
// 该接口不再使用
//
// 
//
// 检查区块链信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
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
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CheckChainWithContext(ctx context.Context, request *CheckChainRequest) (response *CheckChainResponse, err error) {
    if request == nil {
        request = NewCheckChainRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CheckChain require credential")
    }

    request.SetContext(ctx)
    
    response = NewCheckChainResponse()
    err = c.Send(request, response)
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
// 该接口不再使用
//
// 
//
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
    return c.CreateCredentialWithContext(context.Background(), request)
}

// CreateCredential
// 该接口不再使用
//
// 
//
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCredential require credential")
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
// 该接口不再使用
//
// 
//
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
    return c.CreateSelectiveCredentialWithContext(context.Background(), request)
}

// CreateSelectiveCredential
// 该接口不再使用
//
// 
//
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateSelectiveCredential require credential")
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
// 该接口不再使用
//
// 
//
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
    return c.CreateTDidWithContext(context.Background(), request)
}

// CreateTDid
// 该接口不再使用
//
// 
//
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateTDid require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateTDidResponse()
    err = c.Send(request, response)
    return
}

func NewCreateTDidByPrivateKeyRequest() (request *CreateTDidByPrivateKeyRequest) {
    request = &CreateTDidByPrivateKeyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdid", APIVersion, "CreateTDidByPrivateKey")
    
    
    return
}

func NewCreateTDidByPrivateKeyResponse() (response *CreateTDidByPrivateKeyResponse) {
    response = &CreateTDidByPrivateKeyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateTDidByPrivateKey
// 该接口不再使用
//
// 
//
// 新建DID根据私钥生成Tdid
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
func (c *Client) CreateTDidByPrivateKey(request *CreateTDidByPrivateKeyRequest) (response *CreateTDidByPrivateKeyResponse, err error) {
    return c.CreateTDidByPrivateKeyWithContext(context.Background(), request)
}

// CreateTDidByPrivateKey
// 该接口不再使用
//
// 
//
// 新建DID根据私钥生成Tdid
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
func (c *Client) CreateTDidByPrivateKeyWithContext(ctx context.Context, request *CreateTDidByPrivateKeyRequest) (response *CreateTDidByPrivateKeyResponse, err error) {
    if request == nil {
        request = NewCreateTDidByPrivateKeyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateTDidByPrivateKey require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateTDidByPrivateKeyResponse()
    err = c.Send(request, response)
    return
}

func NewCreateTDidByPublicKeyRequest() (request *CreateTDidByPublicKeyRequest) {
    request = &CreateTDidByPublicKeyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdid", APIVersion, "CreateTDidByPublicKey")
    
    
    return
}

func NewCreateTDidByPublicKeyResponse() (response *CreateTDidByPublicKeyResponse) {
    response = &CreateTDidByPublicKeyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateTDidByPublicKey
// 该接口不再使用
//
// 
//
//  新建DID根据公钥生成Tdid
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
func (c *Client) CreateTDidByPublicKey(request *CreateTDidByPublicKeyRequest) (response *CreateTDidByPublicKeyResponse, err error) {
    return c.CreateTDidByPublicKeyWithContext(context.Background(), request)
}

// CreateTDidByPublicKey
// 该接口不再使用
//
// 
//
//  新建DID根据公钥生成Tdid
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
func (c *Client) CreateTDidByPublicKeyWithContext(ctx context.Context, request *CreateTDidByPublicKeyRequest) (response *CreateTDidByPublicKeyResponse, err error) {
    if request == nil {
        request = NewCreateTDidByPublicKeyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateTDidByPublicKey require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateTDidByPublicKeyResponse()
    err = c.Send(request, response)
    return
}

func NewGetAuthorityIssuerRequest() (request *GetAuthorityIssuerRequest) {
    request = &GetAuthorityIssuerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdid", APIVersion, "GetAuthorityIssuer")
    
    
    return
}

func NewGetAuthorityIssuerResponse() (response *GetAuthorityIssuerResponse) {
    response = &GetAuthorityIssuerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetAuthorityIssuer
// 该接口不再使用
//
// 
//
// 获取权威机构信息
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
func (c *Client) GetAuthorityIssuer(request *GetAuthorityIssuerRequest) (response *GetAuthorityIssuerResponse, err error) {
    return c.GetAuthorityIssuerWithContext(context.Background(), request)
}

// GetAuthorityIssuer
// 该接口不再使用
//
// 
//
// 获取权威机构信息
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
func (c *Client) GetAuthorityIssuerWithContext(ctx context.Context, request *GetAuthorityIssuerRequest) (response *GetAuthorityIssuerResponse, err error) {
    if request == nil {
        request = NewGetAuthorityIssuerRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetAuthorityIssuer require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetAuthorityIssuerResponse()
    err = c.Send(request, response)
    return
}

func NewGetCptInfoRequest() (request *GetCptInfoRequest) {
    request = &GetCptInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdid", APIVersion, "GetCptInfo")
    
    
    return
}

func NewGetCptInfoResponse() (response *GetCptInfoResponse) {
    response = &GetCptInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetCptInfo
// 该接口不再使用
//
// 
//
// 凭证模版详情
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
func (c *Client) GetCptInfo(request *GetCptInfoRequest) (response *GetCptInfoResponse, err error) {
    return c.GetCptInfoWithContext(context.Background(), request)
}

// GetCptInfo
// 该接口不再使用
//
// 
//
// 凭证模版详情
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
func (c *Client) GetCptInfoWithContext(ctx context.Context, request *GetCptInfoRequest) (response *GetCptInfoResponse, err error) {
    if request == nil {
        request = NewGetCptInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetCptInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetCptInfoResponse()
    err = c.Send(request, response)
    return
}

func NewGetCredentialStatusRequest() (request *GetCredentialStatusRequest) {
    request = &GetCredentialStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdid", APIVersion, "GetCredentialStatus")
    
    
    return
}

func NewGetCredentialStatusResponse() (response *GetCredentialStatusResponse) {
    response = &GetCredentialStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetCredentialStatus
// 该接口不再使用
//
// 
//
// 获取凭证链上状态信息
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
func (c *Client) GetCredentialStatus(request *GetCredentialStatusRequest) (response *GetCredentialStatusResponse, err error) {
    return c.GetCredentialStatusWithContext(context.Background(), request)
}

// GetCredentialStatus
// 该接口不再使用
//
// 
//
// 获取凭证链上状态信息
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
func (c *Client) GetCredentialStatusWithContext(ctx context.Context, request *GetCredentialStatusRequest) (response *GetCredentialStatusResponse, err error) {
    if request == nil {
        request = NewGetCredentialStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetCredentialStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetCredentialStatusResponse()
    err = c.Send(request, response)
    return
}

func NewGetDidDocumentRequest() (request *GetDidDocumentRequest) {
    request = &GetDidDocumentRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdid", APIVersion, "GetDidDocument")
    
    
    return
}

func NewGetDidDocumentResponse() (response *GetDidDocumentResponse) {
    response = &GetDidDocumentResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetDidDocument
// 该接口不再使用
//
// 
//
// 查看DID文档
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
func (c *Client) GetDidDocument(request *GetDidDocumentRequest) (response *GetDidDocumentResponse, err error) {
    return c.GetDidDocumentWithContext(context.Background(), request)
}

// GetDidDocument
// 该接口不再使用
//
// 
//
// 查看DID文档
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
func (c *Client) GetDidDocumentWithContext(ctx context.Context, request *GetDidDocumentRequest) (response *GetDidDocumentResponse, err error) {
    if request == nil {
        request = NewGetDidDocumentRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetDidDocument require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetDidDocumentResponse()
    err = c.Send(request, response)
    return
}

func NewRegisterCptRequest() (request *RegisterCptRequest) {
    request = &RegisterCptRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdid", APIVersion, "RegisterCpt")
    
    
    return
}

func NewRegisterCptResponse() (response *RegisterCptResponse) {
    response = &RegisterCptResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// RegisterCpt
// 该接口不再使用
//
// 
//
// 凭证模版新建
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
func (c *Client) RegisterCpt(request *RegisterCptRequest) (response *RegisterCptResponse, err error) {
    return c.RegisterCptWithContext(context.Background(), request)
}

// RegisterCpt
// 该接口不再使用
//
// 
//
// 凭证模版新建
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
func (c *Client) RegisterCptWithContext(ctx context.Context, request *RegisterCptRequest) (response *RegisterCptResponse, err error) {
    if request == nil {
        request = NewRegisterCptRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RegisterCpt require credential")
    }

    request.SetContext(ctx)
    
    response = NewRegisterCptResponse()
    err = c.Send(request, response)
    return
}

func NewSetCredentialStatusRequest() (request *SetCredentialStatusRequest) {
    request = &SetCredentialStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdid", APIVersion, "SetCredentialStatus")
    
    
    return
}

func NewSetCredentialStatusResponse() (response *SetCredentialStatusResponse) {
    response = &SetCredentialStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// SetCredentialStatus
// 该接口不再使用
//
// 
//
// 设置凭证链上状态
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
func (c *Client) SetCredentialStatus(request *SetCredentialStatusRequest) (response *SetCredentialStatusResponse, err error) {
    return c.SetCredentialStatusWithContext(context.Background(), request)
}

// SetCredentialStatus
// 该接口不再使用
//
// 
//
// 设置凭证链上状态
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
func (c *Client) SetCredentialStatusWithContext(ctx context.Context, request *SetCredentialStatusRequest) (response *SetCredentialStatusResponse, err error) {
    if request == nil {
        request = NewSetCredentialStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SetCredentialStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewSetCredentialStatusResponse()
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
// 该接口不再使用
//
// 
//
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
    return c.VerifyCredentialWithContext(context.Background(), request)
}

// VerifyCredential
// 该接口不再使用
//
// 
//
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("VerifyCredential require credential")
    }

    request.SetContext(ctx)
    
    response = NewVerifyCredentialResponse()
    err = c.Send(request, response)
    return
}
