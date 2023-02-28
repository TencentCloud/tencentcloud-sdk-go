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

package v20210624

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2021-06-24"

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


func NewCreateBPFakeURLRequest() (request *CreateBPFakeURLRequest) {
    request = &CreateBPFakeURLRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bma", APIVersion, "CreateBPFakeURL")
    
    
    return
}

func NewCreateBPFakeURLResponse() (response *CreateBPFakeURLResponse) {
    response = &CreateBPFakeURLResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateBPFakeURL
// 添加仿冒链接（举报）
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDACTION = "InvalidAction"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  NOSUCHVERSION = "NoSuchVersion"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateBPFakeURL(request *CreateBPFakeURLRequest) (response *CreateBPFakeURLResponse, err error) {
    return c.CreateBPFakeURLWithContext(context.Background(), request)
}

// CreateBPFakeURL
// 添加仿冒链接（举报）
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDACTION = "InvalidAction"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  NOSUCHVERSION = "NoSuchVersion"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateBPFakeURLWithContext(ctx context.Context, request *CreateBPFakeURLRequest) (response *CreateBPFakeURLResponse, err error) {
    if request == nil {
        request = NewCreateBPFakeURLRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateBPFakeURL require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateBPFakeURLResponse()
    err = c.Send(request, response)
    return
}

func NewCreateBPFalseTicketRequest() (request *CreateBPFalseTicketRequest) {
    request = &CreateBPFalseTicketRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bma", APIVersion, "CreateBPFalseTicket")
    
    
    return
}

func NewCreateBPFalseTicketResponse() (response *CreateBPFalseTicketResponse) {
    response = &CreateBPFalseTicketResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateBPFalseTicket
// 添加误报工单
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDACTION = "InvalidAction"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  NOSUCHVERSION = "NoSuchVersion"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateBPFalseTicket(request *CreateBPFalseTicketRequest) (response *CreateBPFalseTicketResponse, err error) {
    return c.CreateBPFalseTicketWithContext(context.Background(), request)
}

// CreateBPFalseTicket
// 添加误报工单
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDACTION = "InvalidAction"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  NOSUCHVERSION = "NoSuchVersion"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateBPFalseTicketWithContext(ctx context.Context, request *CreateBPFalseTicketRequest) (response *CreateBPFalseTicketResponse, err error) {
    if request == nil {
        request = NewCreateBPFalseTicketRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateBPFalseTicket require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateBPFalseTicketResponse()
    err = c.Send(request, response)
    return
}

func NewCreateBPOfflineAttachmentRequest() (request *CreateBPOfflineAttachmentRequest) {
    request = &CreateBPOfflineAttachmentRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bma", APIVersion, "CreateBPOfflineAttachment")
    
    
    return
}

func NewCreateBPOfflineAttachmentResponse() (response *CreateBPOfflineAttachmentResponse) {
    response = &CreateBPOfflineAttachmentResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateBPOfflineAttachment
// 添加下线材料
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDACTION = "InvalidAction"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  NOSUCHVERSION = "NoSuchVersion"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateBPOfflineAttachment(request *CreateBPOfflineAttachmentRequest) (response *CreateBPOfflineAttachmentResponse, err error) {
    return c.CreateBPOfflineAttachmentWithContext(context.Background(), request)
}

// CreateBPOfflineAttachment
// 添加下线材料
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDACTION = "InvalidAction"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  NOSUCHVERSION = "NoSuchVersion"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateBPOfflineAttachmentWithContext(ctx context.Context, request *CreateBPOfflineAttachmentRequest) (response *CreateBPOfflineAttachmentResponse, err error) {
    if request == nil {
        request = NewCreateBPOfflineAttachmentRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateBPOfflineAttachment require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateBPOfflineAttachmentResponse()
    err = c.Send(request, response)
    return
}

func NewCreateBPOfflineTicketRequest() (request *CreateBPOfflineTicketRequest) {
    request = &CreateBPOfflineTicketRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bma", APIVersion, "CreateBPOfflineTicket")
    
    
    return
}

func NewCreateBPOfflineTicketResponse() (response *CreateBPOfflineTicketResponse) {
    response = &CreateBPOfflineTicketResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateBPOfflineTicket
// 添加下线工单
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDACTION = "InvalidAction"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  NOSUCHVERSION = "NoSuchVersion"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateBPOfflineTicket(request *CreateBPOfflineTicketRequest) (response *CreateBPOfflineTicketResponse, err error) {
    return c.CreateBPOfflineTicketWithContext(context.Background(), request)
}

// CreateBPOfflineTicket
// 添加下线工单
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDACTION = "InvalidAction"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  NOSUCHVERSION = "NoSuchVersion"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateBPOfflineTicketWithContext(ctx context.Context, request *CreateBPOfflineTicketRequest) (response *CreateBPOfflineTicketResponse, err error) {
    if request == nil {
        request = NewCreateBPOfflineTicketRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateBPOfflineTicket require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateBPOfflineTicketResponse()
    err = c.Send(request, response)
    return
}

func NewCreateBPProtectURLsRequest() (request *CreateBPProtectURLsRequest) {
    request = &CreateBPProtectURLsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bma", APIVersion, "CreateBPProtectURLs")
    
    
    return
}

func NewCreateBPProtectURLsResponse() (response *CreateBPProtectURLsResponse) {
    response = &CreateBPProtectURLsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateBPProtectURLs
// 添加保护网站
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDACTION = "InvalidAction"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  NOSUCHVERSION = "NoSuchVersion"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateBPProtectURLs(request *CreateBPProtectURLsRequest) (response *CreateBPProtectURLsResponse, err error) {
    return c.CreateBPProtectURLsWithContext(context.Background(), request)
}

// CreateBPProtectURLs
// 添加保护网站
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDACTION = "InvalidAction"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  NOSUCHVERSION = "NoSuchVersion"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateBPProtectURLsWithContext(ctx context.Context, request *CreateBPProtectURLsRequest) (response *CreateBPProtectURLsResponse, err error) {
    if request == nil {
        request = NewCreateBPProtectURLsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateBPProtectURLs require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateBPProtectURLsResponse()
    err = c.Send(request, response)
    return
}

func NewCreateCRBlockRequest() (request *CreateCRBlockRequest) {
    request = &CreateCRBlockRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bma", APIVersion, "CreateCRBlock")
    
    
    return
}

func NewCreateCRBlockResponse() (response *CreateCRBlockResponse) {
    response = &CreateCRBlockResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateCRBlock
// 新建协查处置
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDACTION = "InvalidAction"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  NOSUCHVERSION = "NoSuchVersion"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateCRBlock(request *CreateCRBlockRequest) (response *CreateCRBlockResponse, err error) {
    return c.CreateCRBlockWithContext(context.Background(), request)
}

// CreateCRBlock
// 新建协查处置
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDACTION = "InvalidAction"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  NOSUCHVERSION = "NoSuchVersion"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateCRBlockWithContext(ctx context.Context, request *CreateCRBlockRequest) (response *CreateCRBlockResponse, err error) {
    if request == nil {
        request = NewCreateCRBlockRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCRBlock require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateCRBlockResponse()
    err = c.Send(request, response)
    return
}

func NewCreateCRCompanyVerifyRequest() (request *CreateCRCompanyVerifyRequest) {
    request = &CreateCRCompanyVerifyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bma", APIVersion, "CreateCRCompanyVerify")
    
    
    return
}

func NewCreateCRCompanyVerifyResponse() (response *CreateCRCompanyVerifyResponse) {
    response = &CreateCRCompanyVerifyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateCRCompanyVerify
// 本接口用于企业认证，新接入用户必须认证后才可以进行后续操作（个人认证和企业认证二选一），只需认证一次即可
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDACTION = "InvalidAction"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  NOSUCHVERSION = "NoSuchVersion"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateCRCompanyVerify(request *CreateCRCompanyVerifyRequest) (response *CreateCRCompanyVerifyResponse, err error) {
    return c.CreateCRCompanyVerifyWithContext(context.Background(), request)
}

// CreateCRCompanyVerify
// 本接口用于企业认证，新接入用户必须认证后才可以进行后续操作（个人认证和企业认证二选一），只需认证一次即可
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDACTION = "InvalidAction"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  NOSUCHVERSION = "NoSuchVersion"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateCRCompanyVerifyWithContext(ctx context.Context, request *CreateCRCompanyVerifyRequest) (response *CreateCRCompanyVerifyResponse, err error) {
    if request == nil {
        request = NewCreateCRCompanyVerifyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCRCompanyVerify require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateCRCompanyVerifyResponse()
    err = c.Send(request, response)
    return
}

func NewCreateCRDesktopCodeRequest() (request *CreateCRDesktopCodeRequest) {
    request = &CreateCRDesktopCodeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bma", APIVersion, "CreateCRDesktopCode")
    
    
    return
}

func NewCreateCRDesktopCodeResponse() (response *CreateCRDesktopCodeResponse) {
    response = &CreateCRDesktopCodeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateCRDesktopCode
// 新建过程取证码
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDACTION = "InvalidAction"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  NOSUCHVERSION = "NoSuchVersion"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateCRDesktopCode(request *CreateCRDesktopCodeRequest) (response *CreateCRDesktopCodeResponse, err error) {
    return c.CreateCRDesktopCodeWithContext(context.Background(), request)
}

// CreateCRDesktopCode
// 新建过程取证码
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDACTION = "InvalidAction"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  NOSUCHVERSION = "NoSuchVersion"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateCRDesktopCodeWithContext(ctx context.Context, request *CreateCRDesktopCodeRequest) (response *CreateCRDesktopCodeResponse, err error) {
    if request == nil {
        request = NewCreateCRDesktopCodeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCRDesktopCode require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateCRDesktopCodeResponse()
    err = c.Send(request, response)
    return
}

func NewCreateCRObtainRequest() (request *CreateCRObtainRequest) {
    request = &CreateCRObtainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bma", APIVersion, "CreateCRObtain")
    
    
    return
}

func NewCreateCRObtainResponse() (response *CreateCRObtainResponse) {
    response = &CreateCRObtainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateCRObtain
// 版权保护-新建取证接口
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDACTION = "InvalidAction"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  NOSUCHVERSION = "NoSuchVersion"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateCRObtain(request *CreateCRObtainRequest) (response *CreateCRObtainResponse, err error) {
    return c.CreateCRObtainWithContext(context.Background(), request)
}

// CreateCRObtain
// 版权保护-新建取证接口
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDACTION = "InvalidAction"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  NOSUCHVERSION = "NoSuchVersion"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateCRObtainWithContext(ctx context.Context, request *CreateCRObtainRequest) (response *CreateCRObtainResponse, err error) {
    if request == nil {
        request = NewCreateCRObtainRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCRObtain require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateCRObtainResponse()
    err = c.Send(request, response)
    return
}

func NewCreateCRRightRequest() (request *CreateCRRightRequest) {
    request = &CreateCRRightRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bma", APIVersion, "CreateCRRight")
    
    
    return
}

func NewCreateCRRightResponse() (response *CreateCRRightResponse) {
    response = &CreateCRRightResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateCRRight
// 版权保护-新建发函接口
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDACTION = "InvalidAction"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  NOSUCHVERSION = "NoSuchVersion"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateCRRight(request *CreateCRRightRequest) (response *CreateCRRightResponse, err error) {
    return c.CreateCRRightWithContext(context.Background(), request)
}

// CreateCRRight
// 版权保护-新建发函接口
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDACTION = "InvalidAction"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  NOSUCHVERSION = "NoSuchVersion"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateCRRightWithContext(ctx context.Context, request *CreateCRRightRequest) (response *CreateCRRightResponse, err error) {
    if request == nil {
        request = NewCreateCRRightRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCRRight require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateCRRightResponse()
    err = c.Send(request, response)
    return
}

func NewCreateCRRightFileRequest() (request *CreateCRRightFileRequest) {
    request = &CreateCRRightFileRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bma", APIVersion, "CreateCRRightFile")
    
    
    return
}

func NewCreateCRRightFileResponse() (response *CreateCRRightFileResponse) {
    response = &CreateCRRightFileResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateCRRightFile
// 权属文件添加
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDACTION = "InvalidAction"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  NOSUCHVERSION = "NoSuchVersion"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateCRRightFile(request *CreateCRRightFileRequest) (response *CreateCRRightFileResponse, err error) {
    return c.CreateCRRightFileWithContext(context.Background(), request)
}

// CreateCRRightFile
// 权属文件添加
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDACTION = "InvalidAction"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  NOSUCHVERSION = "NoSuchVersion"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateCRRightFileWithContext(ctx context.Context, request *CreateCRRightFileRequest) (response *CreateCRRightFileResponse, err error) {
    if request == nil {
        request = NewCreateCRRightFileRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCRRightFile require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateCRRightFileResponse()
    err = c.Send(request, response)
    return
}

func NewCreateCRTortRequest() (request *CreateCRTortRequest) {
    request = &CreateCRTortRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bma", APIVersion, "CreateCRTort")
    
    
    return
}

func NewCreateCRTortResponse() (response *CreateCRTortResponse) {
    response = &CreateCRTortResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateCRTort
// 举报侵权链接
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDACTION = "InvalidAction"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  NOSUCHVERSION = "NoSuchVersion"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateCRTort(request *CreateCRTortRequest) (response *CreateCRTortResponse, err error) {
    return c.CreateCRTortWithContext(context.Background(), request)
}

// CreateCRTort
// 举报侵权链接
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDACTION = "InvalidAction"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  NOSUCHVERSION = "NoSuchVersion"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateCRTortWithContext(ctx context.Context, request *CreateCRTortRequest) (response *CreateCRTortResponse, err error) {
    if request == nil {
        request = NewCreateCRTortRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCRTort require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateCRTortResponse()
    err = c.Send(request, response)
    return
}

func NewCreateCRUserVerifyRequest() (request *CreateCRUserVerifyRequest) {
    request = &CreateCRUserVerifyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bma", APIVersion, "CreateCRUserVerify")
    
    
    return
}

func NewCreateCRUserVerifyResponse() (response *CreateCRUserVerifyResponse) {
    response = &CreateCRUserVerifyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateCRUserVerify
// 本接口用于个人认证，新接入用户必须认证后才可以进行后续操作（个人认证和企业认证二选一），只需认证一次即可
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDACTION = "InvalidAction"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  NOSUCHVERSION = "NoSuchVersion"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateCRUserVerify(request *CreateCRUserVerifyRequest) (response *CreateCRUserVerifyResponse, err error) {
    return c.CreateCRUserVerifyWithContext(context.Background(), request)
}

// CreateCRUserVerify
// 本接口用于个人认证，新接入用户必须认证后才可以进行后续操作（个人认证和企业认证二选一），只需认证一次即可
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDACTION = "InvalidAction"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  NOSUCHVERSION = "NoSuchVersion"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateCRUserVerifyWithContext(ctx context.Context, request *CreateCRUserVerifyRequest) (response *CreateCRUserVerifyResponse, err error) {
    if request == nil {
        request = NewCreateCRUserVerifyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCRUserVerify require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateCRUserVerifyResponse()
    err = c.Send(request, response)
    return
}

func NewCreateCRWorkRequest() (request *CreateCRWorkRequest) {
    request = &CreateCRWorkRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bma", APIVersion, "CreateCRWork")
    
    
    return
}

func NewCreateCRWorkResponse() (response *CreateCRWorkResponse) {
    response = &CreateCRWorkResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateCRWork
// 新建作品
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDACTION = "InvalidAction"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  NOSUCHVERSION = "NoSuchVersion"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateCRWork(request *CreateCRWorkRequest) (response *CreateCRWorkResponse, err error) {
    return c.CreateCRWorkWithContext(context.Background(), request)
}

// CreateCRWork
// 新建作品
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDACTION = "InvalidAction"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  NOSUCHVERSION = "NoSuchVersion"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateCRWorkWithContext(ctx context.Context, request *CreateCRWorkRequest) (response *CreateCRWorkResponse, err error) {
    if request == nil {
        request = NewCreateCRWorkRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCRWork require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateCRWorkResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBPCompanyInfoRequest() (request *DescribeBPCompanyInfoRequest) {
    request = &DescribeBPCompanyInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bma", APIVersion, "DescribeBPCompanyInfo")
    
    
    return
}

func NewDescribeBPCompanyInfoResponse() (response *DescribeBPCompanyInfoResponse) {
    response = &DescribeBPCompanyInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeBPCompanyInfo
// 查询企业信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDACTION = "InvalidAction"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  NOSUCHVERSION = "NoSuchVersion"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeBPCompanyInfo(request *DescribeBPCompanyInfoRequest) (response *DescribeBPCompanyInfoResponse, err error) {
    return c.DescribeBPCompanyInfoWithContext(context.Background(), request)
}

// DescribeBPCompanyInfo
// 查询企业信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDACTION = "InvalidAction"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  NOSUCHVERSION = "NoSuchVersion"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeBPCompanyInfoWithContext(ctx context.Context, request *DescribeBPCompanyInfoRequest) (response *DescribeBPCompanyInfoResponse, err error) {
    if request == nil {
        request = NewDescribeBPCompanyInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBPCompanyInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBPCompanyInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBPFakeURLsRequest() (request *DescribeBPFakeURLsRequest) {
    request = &DescribeBPFakeURLsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bma", APIVersion, "DescribeBPFakeURLs")
    
    
    return
}

func NewDescribeBPFakeURLsResponse() (response *DescribeBPFakeURLsResponse) {
    response = &DescribeBPFakeURLsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeBPFakeURLs
// 查询仿冒链接
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeBPFakeURLs(request *DescribeBPFakeURLsRequest) (response *DescribeBPFakeURLsResponse, err error) {
    return c.DescribeBPFakeURLsWithContext(context.Background(), request)
}

// DescribeBPFakeURLs
// 查询仿冒链接
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeBPFakeURLsWithContext(ctx context.Context, request *DescribeBPFakeURLsRequest) (response *DescribeBPFakeURLsResponse, err error) {
    if request == nil {
        request = NewDescribeBPFakeURLsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBPFakeURLs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBPFakeURLsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBPProtectURLsRequest() (request *DescribeBPProtectURLsRequest) {
    request = &DescribeBPProtectURLsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bma", APIVersion, "DescribeBPProtectURLs")
    
    
    return
}

func NewDescribeBPProtectURLsResponse() (response *DescribeBPProtectURLsResponse) {
    response = &DescribeBPProtectURLsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeBPProtectURLs
// 查询保护网站
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDACTION = "InvalidAction"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  NOSUCHVERSION = "NoSuchVersion"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeBPProtectURLs(request *DescribeBPProtectURLsRequest) (response *DescribeBPProtectURLsResponse, err error) {
    return c.DescribeBPProtectURLsWithContext(context.Background(), request)
}

// DescribeBPProtectURLs
// 查询保护网站
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDACTION = "InvalidAction"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  NOSUCHVERSION = "NoSuchVersion"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeBPProtectURLsWithContext(ctx context.Context, request *DescribeBPProtectURLsRequest) (response *DescribeBPProtectURLsResponse, err error) {
    if request == nil {
        request = NewDescribeBPProtectURLsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBPProtectURLs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBPProtectURLsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBPReportFakeURLsRequest() (request *DescribeBPReportFakeURLsRequest) {
    request = &DescribeBPReportFakeURLsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bma", APIVersion, "DescribeBPReportFakeURLs")
    
    
    return
}

func NewDescribeBPReportFakeURLsResponse() (response *DescribeBPReportFakeURLsResponse) {
    response = &DescribeBPReportFakeURLsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeBPReportFakeURLs
// 查询举报列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDACTION = "InvalidAction"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  NOSUCHVERSION = "NoSuchVersion"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeBPReportFakeURLs(request *DescribeBPReportFakeURLsRequest) (response *DescribeBPReportFakeURLsResponse, err error) {
    return c.DescribeBPReportFakeURLsWithContext(context.Background(), request)
}

// DescribeBPReportFakeURLs
// 查询举报列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDACTION = "InvalidAction"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  NOSUCHVERSION = "NoSuchVersion"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeBPReportFakeURLsWithContext(ctx context.Context, request *DescribeBPReportFakeURLsRequest) (response *DescribeBPReportFakeURLsResponse, err error) {
    if request == nil {
        request = NewDescribeBPReportFakeURLsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBPReportFakeURLs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBPReportFakeURLsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCRMonitorDetailRequest() (request *DescribeCRMonitorDetailRequest) {
    request = &DescribeCRMonitorDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bma", APIVersion, "DescribeCRMonitorDetail")
    
    
    return
}

func NewDescribeCRMonitorDetailResponse() (response *DescribeCRMonitorDetailResponse) {
    response = &DescribeCRMonitorDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCRMonitorDetail
// 版权保护-查询作品监测详情接口
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDACTION = "InvalidAction"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  NOSUCHVERSION = "NoSuchVersion"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeCRMonitorDetail(request *DescribeCRMonitorDetailRequest) (response *DescribeCRMonitorDetailResponse, err error) {
    return c.DescribeCRMonitorDetailWithContext(context.Background(), request)
}

// DescribeCRMonitorDetail
// 版权保护-查询作品监测详情接口
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDACTION = "InvalidAction"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  NOSUCHVERSION = "NoSuchVersion"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeCRMonitorDetailWithContext(ctx context.Context, request *DescribeCRMonitorDetailRequest) (response *DescribeCRMonitorDetailResponse, err error) {
    if request == nil {
        request = NewDescribeCRMonitorDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCRMonitorDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCRMonitorDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCRMonitorsRequest() (request *DescribeCRMonitorsRequest) {
    request = &DescribeCRMonitorsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bma", APIVersion, "DescribeCRMonitors")
    
    
    return
}

func NewDescribeCRMonitorsResponse() (response *DescribeCRMonitorsResponse) {
    response = &DescribeCRMonitorsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCRMonitors
// 版权保护-查询监测列表接口
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDACTION = "InvalidAction"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  NOSUCHVERSION = "NoSuchVersion"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeCRMonitors(request *DescribeCRMonitorsRequest) (response *DescribeCRMonitorsResponse, err error) {
    return c.DescribeCRMonitorsWithContext(context.Background(), request)
}

// DescribeCRMonitors
// 版权保护-查询监测列表接口
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDACTION = "InvalidAction"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  NOSUCHVERSION = "NoSuchVersion"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeCRMonitorsWithContext(ctx context.Context, request *DescribeCRMonitorsRequest) (response *DescribeCRMonitorsResponse, err error) {
    if request == nil {
        request = NewDescribeCRMonitorsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCRMonitors require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCRMonitorsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCRObtainDetailRequest() (request *DescribeCRObtainDetailRequest) {
    request = &DescribeCRObtainDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bma", APIVersion, "DescribeCRObtainDetail")
    
    
    return
}

func NewDescribeCRObtainDetailResponse() (response *DescribeCRObtainDetailResponse) {
    response = &DescribeCRObtainDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCRObtainDetail
// 查询取证详情
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDACTION = "InvalidAction"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  NOSUCHVERSION = "NoSuchVersion"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeCRObtainDetail(request *DescribeCRObtainDetailRequest) (response *DescribeCRObtainDetailResponse, err error) {
    return c.DescribeCRObtainDetailWithContext(context.Background(), request)
}

// DescribeCRObtainDetail
// 查询取证详情
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDACTION = "InvalidAction"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  NOSUCHVERSION = "NoSuchVersion"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeCRObtainDetailWithContext(ctx context.Context, request *DescribeCRObtainDetailRequest) (response *DescribeCRObtainDetailResponse, err error) {
    if request == nil {
        request = NewDescribeCRObtainDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCRObtainDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCRObtainDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCRWorkInfoRequest() (request *DescribeCRWorkInfoRequest) {
    request = &DescribeCRWorkInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bma", APIVersion, "DescribeCRWorkInfo")
    
    
    return
}

func NewDescribeCRWorkInfoResponse() (response *DescribeCRWorkInfoResponse) {
    response = &DescribeCRWorkInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCRWorkInfo
// 查询作品基本信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDACTION = "InvalidAction"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  NOSUCHVERSION = "NoSuchVersion"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeCRWorkInfo(request *DescribeCRWorkInfoRequest) (response *DescribeCRWorkInfoResponse, err error) {
    return c.DescribeCRWorkInfoWithContext(context.Background(), request)
}

// DescribeCRWorkInfo
// 查询作品基本信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDACTION = "InvalidAction"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  NOSUCHVERSION = "NoSuchVersion"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeCRWorkInfoWithContext(ctx context.Context, request *DescribeCRWorkInfoRequest) (response *DescribeCRWorkInfoResponse, err error) {
    if request == nil {
        request = NewDescribeCRWorkInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCRWorkInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCRWorkInfoResponse()
    err = c.Send(request, response)
    return
}

func NewModifyBPOfflineAttachmentRequest() (request *ModifyBPOfflineAttachmentRequest) {
    request = &ModifyBPOfflineAttachmentRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bma", APIVersion, "ModifyBPOfflineAttachment")
    
    
    return
}

func NewModifyBPOfflineAttachmentResponse() (response *ModifyBPOfflineAttachmentResponse) {
    response = &ModifyBPOfflineAttachmentResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyBPOfflineAttachment
// 修改下线材料
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDACTION = "InvalidAction"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  NOSUCHVERSION = "NoSuchVersion"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyBPOfflineAttachment(request *ModifyBPOfflineAttachmentRequest) (response *ModifyBPOfflineAttachmentResponse, err error) {
    return c.ModifyBPOfflineAttachmentWithContext(context.Background(), request)
}

// ModifyBPOfflineAttachment
// 修改下线材料
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDACTION = "InvalidAction"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  NOSUCHVERSION = "NoSuchVersion"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyBPOfflineAttachmentWithContext(ctx context.Context, request *ModifyBPOfflineAttachmentRequest) (response *ModifyBPOfflineAttachmentResponse, err error) {
    if request == nil {
        request = NewModifyBPOfflineAttachmentRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyBPOfflineAttachment require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyBPOfflineAttachmentResponse()
    err = c.Send(request, response)
    return
}

func NewModifyCRBlockStatusRequest() (request *ModifyCRBlockStatusRequest) {
    request = &ModifyCRBlockStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bma", APIVersion, "ModifyCRBlockStatus")
    
    
    return
}

func NewModifyCRBlockStatusResponse() (response *ModifyCRBlockStatusResponse) {
    response = &ModifyCRBlockStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyCRBlockStatus
// 协查处置申请
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDACTION = "InvalidAction"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  NOSUCHVERSION = "NoSuchVersion"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyCRBlockStatus(request *ModifyCRBlockStatusRequest) (response *ModifyCRBlockStatusResponse, err error) {
    return c.ModifyCRBlockStatusWithContext(context.Background(), request)
}

// ModifyCRBlockStatus
// 协查处置申请
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDACTION = "InvalidAction"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  NOSUCHVERSION = "NoSuchVersion"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyCRBlockStatusWithContext(ctx context.Context, request *ModifyCRBlockStatusRequest) (response *ModifyCRBlockStatusResponse, err error) {
    if request == nil {
        request = NewModifyCRBlockStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyCRBlockStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyCRBlockStatusResponse()
    err = c.Send(request, response)
    return
}

func NewModifyCRMonitorRequest() (request *ModifyCRMonitorRequest) {
    request = &ModifyCRMonitorRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bma", APIVersion, "ModifyCRMonitor")
    
    
    return
}

func NewModifyCRMonitorResponse() (response *ModifyCRMonitorResponse) {
    response = &ModifyCRMonitorResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyCRMonitor
// 开启/关闭监测
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDACTION = "InvalidAction"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  NOSUCHVERSION = "NoSuchVersion"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyCRMonitor(request *ModifyCRMonitorRequest) (response *ModifyCRMonitorResponse, err error) {
    return c.ModifyCRMonitorWithContext(context.Background(), request)
}

// ModifyCRMonitor
// 开启/关闭监测
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDACTION = "InvalidAction"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  NOSUCHVERSION = "NoSuchVersion"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyCRMonitorWithContext(ctx context.Context, request *ModifyCRMonitorRequest) (response *ModifyCRMonitorResponse, err error) {
    if request == nil {
        request = NewModifyCRMonitorRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyCRMonitor require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyCRMonitorResponse()
    err = c.Send(request, response)
    return
}

func NewModifyCRObtainStatusRequest() (request *ModifyCRObtainStatusRequest) {
    request = &ModifyCRObtainStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bma", APIVersion, "ModifyCRObtainStatus")
    
    
    return
}

func NewModifyCRObtainStatusResponse() (response *ModifyCRObtainStatusResponse) {
    response = &ModifyCRObtainStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyCRObtainStatus
// 取证申请
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDACTION = "InvalidAction"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  NOSUCHVERSION = "NoSuchVersion"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyCRObtainStatus(request *ModifyCRObtainStatusRequest) (response *ModifyCRObtainStatusResponse, err error) {
    return c.ModifyCRObtainStatusWithContext(context.Background(), request)
}

// ModifyCRObtainStatus
// 取证申请
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDACTION = "InvalidAction"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  NOSUCHVERSION = "NoSuchVersion"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyCRObtainStatusWithContext(ctx context.Context, request *ModifyCRObtainStatusRequest) (response *ModifyCRObtainStatusResponse, err error) {
    if request == nil {
        request = NewModifyCRObtainStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyCRObtainStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyCRObtainStatusResponse()
    err = c.Send(request, response)
    return
}

func NewModifyCRRightStatusRequest() (request *ModifyCRRightStatusRequest) {
    request = &ModifyCRRightStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bma", APIVersion, "ModifyCRRightStatus")
    
    
    return
}

func NewModifyCRRightStatusResponse() (response *ModifyCRRightStatusResponse) {
    response = &ModifyCRRightStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyCRRightStatus
// 发函申请
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDACTION = "InvalidAction"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  NOSUCHVERSION = "NoSuchVersion"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyCRRightStatus(request *ModifyCRRightStatusRequest) (response *ModifyCRRightStatusResponse, err error) {
    return c.ModifyCRRightStatusWithContext(context.Background(), request)
}

// ModifyCRRightStatus
// 发函申请
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDACTION = "InvalidAction"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  NOSUCHVERSION = "NoSuchVersion"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyCRRightStatusWithContext(ctx context.Context, request *ModifyCRRightStatusRequest) (response *ModifyCRRightStatusResponse, err error) {
    if request == nil {
        request = NewModifyCRRightStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyCRRightStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyCRRightStatusResponse()
    err = c.Send(request, response)
    return
}

func NewModifyCRWhiteListRequest() (request *ModifyCRWhiteListRequest) {
    request = &ModifyCRWhiteListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bma", APIVersion, "ModifyCRWhiteList")
    
    
    return
}

func NewModifyCRWhiteListResponse() (response *ModifyCRWhiteListResponse) {
    response = &ModifyCRWhiteListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyCRWhiteList
// 修改白名单列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDACTION = "InvalidAction"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  NOSUCHVERSION = "NoSuchVersion"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyCRWhiteList(request *ModifyCRWhiteListRequest) (response *ModifyCRWhiteListResponse, err error) {
    return c.ModifyCRWhiteListWithContext(context.Background(), request)
}

// ModifyCRWhiteList
// 修改白名单列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDACTION = "InvalidAction"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  NOSUCHVERSION = "NoSuchVersion"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyCRWhiteListWithContext(ctx context.Context, request *ModifyCRWhiteListRequest) (response *ModifyCRWhiteListResponse, err error) {
    if request == nil {
        request = NewModifyCRWhiteListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyCRWhiteList require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyCRWhiteListResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateCRWorkRequest() (request *UpdateCRWorkRequest) {
    request = &UpdateCRWorkRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bma", APIVersion, "UpdateCRWork")
    
    
    return
}

func NewUpdateCRWorkResponse() (response *UpdateCRWorkResponse) {
    response = &UpdateCRWorkResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpdateCRWork
// 更新作品
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDACTION = "InvalidAction"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  NOSUCHVERSION = "NoSuchVersion"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) UpdateCRWork(request *UpdateCRWorkRequest) (response *UpdateCRWorkResponse, err error) {
    return c.UpdateCRWorkWithContext(context.Background(), request)
}

// UpdateCRWork
// 更新作品
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDACTION = "InvalidAction"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  NOSUCHVERSION = "NoSuchVersion"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) UpdateCRWorkWithContext(ctx context.Context, request *UpdateCRWorkRequest) (response *UpdateCRWorkResponse, err error) {
    if request == nil {
        request = NewUpdateCRWorkRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateCRWork require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateCRWorkResponse()
    err = c.Send(request, response)
    return
}
