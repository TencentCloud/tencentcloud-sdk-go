// Copyright (c) 2017-2025 Tencent. All Rights Reserved.
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

package v20230418

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2023-04-18"

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


func NewCreateRabbitMQServerlessBindingRequest() (request *CreateRabbitMQServerlessBindingRequest) {
    request = &CreateRabbitMQServerlessBindingRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trabbit", APIVersion, "CreateRabbitMQServerlessBinding")
    
    
    return
}

func NewCreateRabbitMQServerlessBindingResponse() (response *CreateRabbitMQServerlessBindingResponse) {
    response = &CreateRabbitMQServerlessBindingResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateRabbitMQServerlessBinding
// 创建RabbitMQ路由关系
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateRabbitMQServerlessBinding(request *CreateRabbitMQServerlessBindingRequest) (response *CreateRabbitMQServerlessBindingResponse, err error) {
    return c.CreateRabbitMQServerlessBindingWithContext(context.Background(), request)
}

// CreateRabbitMQServerlessBinding
// 创建RabbitMQ路由关系
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateRabbitMQServerlessBindingWithContext(ctx context.Context, request *CreateRabbitMQServerlessBindingRequest) (response *CreateRabbitMQServerlessBindingResponse, err error) {
    if request == nil {
        request = NewCreateRabbitMQServerlessBindingRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateRabbitMQServerlessBinding require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateRabbitMQServerlessBindingResponse()
    err = c.Send(request, response)
    return
}

func NewCreateRabbitMQServerlessExchangeRequest() (request *CreateRabbitMQServerlessExchangeRequest) {
    request = &CreateRabbitMQServerlessExchangeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trabbit", APIVersion, "CreateRabbitMQServerlessExchange")
    
    
    return
}

func NewCreateRabbitMQServerlessExchangeResponse() (response *CreateRabbitMQServerlessExchangeResponse) {
    response = &CreateRabbitMQServerlessExchangeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateRabbitMQServerlessExchange
// 创建RabbitMQ exchange
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateRabbitMQServerlessExchange(request *CreateRabbitMQServerlessExchangeRequest) (response *CreateRabbitMQServerlessExchangeResponse, err error) {
    return c.CreateRabbitMQServerlessExchangeWithContext(context.Background(), request)
}

// CreateRabbitMQServerlessExchange
// 创建RabbitMQ exchange
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateRabbitMQServerlessExchangeWithContext(ctx context.Context, request *CreateRabbitMQServerlessExchangeRequest) (response *CreateRabbitMQServerlessExchangeResponse, err error) {
    if request == nil {
        request = NewCreateRabbitMQServerlessExchangeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateRabbitMQServerlessExchange require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateRabbitMQServerlessExchangeResponse()
    err = c.Send(request, response)
    return
}

func NewCreateRabbitMQServerlessQueueRequest() (request *CreateRabbitMQServerlessQueueRequest) {
    request = &CreateRabbitMQServerlessQueueRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trabbit", APIVersion, "CreateRabbitMQServerlessQueue")
    
    
    return
}

func NewCreateRabbitMQServerlessQueueResponse() (response *CreateRabbitMQServerlessQueueResponse) {
    response = &CreateRabbitMQServerlessQueueResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateRabbitMQServerlessQueue
// 创建RabbitMQ队列
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateRabbitMQServerlessQueue(request *CreateRabbitMQServerlessQueueRequest) (response *CreateRabbitMQServerlessQueueResponse, err error) {
    return c.CreateRabbitMQServerlessQueueWithContext(context.Background(), request)
}

// CreateRabbitMQServerlessQueue
// 创建RabbitMQ队列
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateRabbitMQServerlessQueueWithContext(ctx context.Context, request *CreateRabbitMQServerlessQueueRequest) (response *CreateRabbitMQServerlessQueueResponse, err error) {
    if request == nil {
        request = NewCreateRabbitMQServerlessQueueRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateRabbitMQServerlessQueue require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateRabbitMQServerlessQueueResponse()
    err = c.Send(request, response)
    return
}

func NewCreateRabbitMQServerlessUserRequest() (request *CreateRabbitMQServerlessUserRequest) {
    request = &CreateRabbitMQServerlessUserRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trabbit", APIVersion, "CreateRabbitMQServerlessUser")
    
    
    return
}

func NewCreateRabbitMQServerlessUserResponse() (response *CreateRabbitMQServerlessUserResponse) {
    response = &CreateRabbitMQServerlessUserResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateRabbitMQServerlessUser
// 创建RabbitMQ的用户
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateRabbitMQServerlessUser(request *CreateRabbitMQServerlessUserRequest) (response *CreateRabbitMQServerlessUserResponse, err error) {
    return c.CreateRabbitMQServerlessUserWithContext(context.Background(), request)
}

// CreateRabbitMQServerlessUser
// 创建RabbitMQ的用户
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateRabbitMQServerlessUserWithContext(ctx context.Context, request *CreateRabbitMQServerlessUserRequest) (response *CreateRabbitMQServerlessUserResponse, err error) {
    if request == nil {
        request = NewCreateRabbitMQServerlessUserRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateRabbitMQServerlessUser require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateRabbitMQServerlessUserResponse()
    err = c.Send(request, response)
    return
}

func NewCreateRabbitMQServerlessVirtualHostRequest() (request *CreateRabbitMQServerlessVirtualHostRequest) {
    request = &CreateRabbitMQServerlessVirtualHostRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trabbit", APIVersion, "CreateRabbitMQServerlessVirtualHost")
    
    
    return
}

func NewCreateRabbitMQServerlessVirtualHostResponse() (response *CreateRabbitMQServerlessVirtualHostResponse) {
    response = &CreateRabbitMQServerlessVirtualHostResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateRabbitMQServerlessVirtualHost
// 创建RabbitMQ的vhost
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateRabbitMQServerlessVirtualHost(request *CreateRabbitMQServerlessVirtualHostRequest) (response *CreateRabbitMQServerlessVirtualHostResponse, err error) {
    return c.CreateRabbitMQServerlessVirtualHostWithContext(context.Background(), request)
}

// CreateRabbitMQServerlessVirtualHost
// 创建RabbitMQ的vhost
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateRabbitMQServerlessVirtualHostWithContext(ctx context.Context, request *CreateRabbitMQServerlessVirtualHostRequest) (response *CreateRabbitMQServerlessVirtualHostResponse, err error) {
    if request == nil {
        request = NewCreateRabbitMQServerlessVirtualHostRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateRabbitMQServerlessVirtualHost require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateRabbitMQServerlessVirtualHostResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteRabbitMQServerlessBindingRequest() (request *DeleteRabbitMQServerlessBindingRequest) {
    request = &DeleteRabbitMQServerlessBindingRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trabbit", APIVersion, "DeleteRabbitMQServerlessBinding")
    
    
    return
}

func NewDeleteRabbitMQServerlessBindingResponse() (response *DeleteRabbitMQServerlessBindingResponse) {
    response = &DeleteRabbitMQServerlessBindingResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteRabbitMQServerlessBinding
// 解绑RabbitMQ路由关系
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteRabbitMQServerlessBinding(request *DeleteRabbitMQServerlessBindingRequest) (response *DeleteRabbitMQServerlessBindingResponse, err error) {
    return c.DeleteRabbitMQServerlessBindingWithContext(context.Background(), request)
}

// DeleteRabbitMQServerlessBinding
// 解绑RabbitMQ路由关系
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteRabbitMQServerlessBindingWithContext(ctx context.Context, request *DeleteRabbitMQServerlessBindingRequest) (response *DeleteRabbitMQServerlessBindingResponse, err error) {
    if request == nil {
        request = NewDeleteRabbitMQServerlessBindingRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteRabbitMQServerlessBinding require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteRabbitMQServerlessBindingResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteRabbitMQServerlessExchangeRequest() (request *DeleteRabbitMQServerlessExchangeRequest) {
    request = &DeleteRabbitMQServerlessExchangeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trabbit", APIVersion, "DeleteRabbitMQServerlessExchange")
    
    
    return
}

func NewDeleteRabbitMQServerlessExchangeResponse() (response *DeleteRabbitMQServerlessExchangeResponse) {
    response = &DeleteRabbitMQServerlessExchangeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteRabbitMQServerlessExchange
// 删除RabbitMQ exchange
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteRabbitMQServerlessExchange(request *DeleteRabbitMQServerlessExchangeRequest) (response *DeleteRabbitMQServerlessExchangeResponse, err error) {
    return c.DeleteRabbitMQServerlessExchangeWithContext(context.Background(), request)
}

// DeleteRabbitMQServerlessExchange
// 删除RabbitMQ exchange
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteRabbitMQServerlessExchangeWithContext(ctx context.Context, request *DeleteRabbitMQServerlessExchangeRequest) (response *DeleteRabbitMQServerlessExchangeResponse, err error) {
    if request == nil {
        request = NewDeleteRabbitMQServerlessExchangeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteRabbitMQServerlessExchange require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteRabbitMQServerlessExchangeResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteRabbitMQServerlessPermissionRequest() (request *DeleteRabbitMQServerlessPermissionRequest) {
    request = &DeleteRabbitMQServerlessPermissionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trabbit", APIVersion, "DeleteRabbitMQServerlessPermission")
    
    
    return
}

func NewDeleteRabbitMQServerlessPermissionResponse() (response *DeleteRabbitMQServerlessPermissionResponse) {
    response = &DeleteRabbitMQServerlessPermissionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteRabbitMQServerlessPermission
// 删除RabbitMQ的权限
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteRabbitMQServerlessPermission(request *DeleteRabbitMQServerlessPermissionRequest) (response *DeleteRabbitMQServerlessPermissionResponse, err error) {
    return c.DeleteRabbitMQServerlessPermissionWithContext(context.Background(), request)
}

// DeleteRabbitMQServerlessPermission
// 删除RabbitMQ的权限
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteRabbitMQServerlessPermissionWithContext(ctx context.Context, request *DeleteRabbitMQServerlessPermissionRequest) (response *DeleteRabbitMQServerlessPermissionResponse, err error) {
    if request == nil {
        request = NewDeleteRabbitMQServerlessPermissionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteRabbitMQServerlessPermission require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteRabbitMQServerlessPermissionResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteRabbitMQServerlessQueueRequest() (request *DeleteRabbitMQServerlessQueueRequest) {
    request = &DeleteRabbitMQServerlessQueueRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trabbit", APIVersion, "DeleteRabbitMQServerlessQueue")
    
    
    return
}

func NewDeleteRabbitMQServerlessQueueResponse() (response *DeleteRabbitMQServerlessQueueResponse) {
    response = &DeleteRabbitMQServerlessQueueResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteRabbitMQServerlessQueue
// 删除RabbitMQ队列
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteRabbitMQServerlessQueue(request *DeleteRabbitMQServerlessQueueRequest) (response *DeleteRabbitMQServerlessQueueResponse, err error) {
    return c.DeleteRabbitMQServerlessQueueWithContext(context.Background(), request)
}

// DeleteRabbitMQServerlessQueue
// 删除RabbitMQ队列
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteRabbitMQServerlessQueueWithContext(ctx context.Context, request *DeleteRabbitMQServerlessQueueRequest) (response *DeleteRabbitMQServerlessQueueResponse, err error) {
    if request == nil {
        request = NewDeleteRabbitMQServerlessQueueRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteRabbitMQServerlessQueue require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteRabbitMQServerlessQueueResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteRabbitMQServerlessUserRequest() (request *DeleteRabbitMQServerlessUserRequest) {
    request = &DeleteRabbitMQServerlessUserRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trabbit", APIVersion, "DeleteRabbitMQServerlessUser")
    
    
    return
}

func NewDeleteRabbitMQServerlessUserResponse() (response *DeleteRabbitMQServerlessUserResponse) {
    response = &DeleteRabbitMQServerlessUserResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteRabbitMQServerlessUser
// 删除RabbitMQ的用户
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteRabbitMQServerlessUser(request *DeleteRabbitMQServerlessUserRequest) (response *DeleteRabbitMQServerlessUserResponse, err error) {
    return c.DeleteRabbitMQServerlessUserWithContext(context.Background(), request)
}

// DeleteRabbitMQServerlessUser
// 删除RabbitMQ的用户
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteRabbitMQServerlessUserWithContext(ctx context.Context, request *DeleteRabbitMQServerlessUserRequest) (response *DeleteRabbitMQServerlessUserResponse, err error) {
    if request == nil {
        request = NewDeleteRabbitMQServerlessUserRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteRabbitMQServerlessUser require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteRabbitMQServerlessUserResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteRabbitMQServerlessVirtualHostRequest() (request *DeleteRabbitMQServerlessVirtualHostRequest) {
    request = &DeleteRabbitMQServerlessVirtualHostRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trabbit", APIVersion, "DeleteRabbitMQServerlessVirtualHost")
    
    
    return
}

func NewDeleteRabbitMQServerlessVirtualHostResponse() (response *DeleteRabbitMQServerlessVirtualHostResponse) {
    response = &DeleteRabbitMQServerlessVirtualHostResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteRabbitMQServerlessVirtualHost
// 删除RabbitMQ的vhost
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteRabbitMQServerlessVirtualHost(request *DeleteRabbitMQServerlessVirtualHostRequest) (response *DeleteRabbitMQServerlessVirtualHostResponse, err error) {
    return c.DeleteRabbitMQServerlessVirtualHostWithContext(context.Background(), request)
}

// DeleteRabbitMQServerlessVirtualHost
// 删除RabbitMQ的vhost
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteRabbitMQServerlessVirtualHostWithContext(ctx context.Context, request *DeleteRabbitMQServerlessVirtualHostRequest) (response *DeleteRabbitMQServerlessVirtualHostResponse, err error) {
    if request == nil {
        request = NewDeleteRabbitMQServerlessVirtualHostRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteRabbitMQServerlessVirtualHost require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteRabbitMQServerlessVirtualHostResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRabbitMQServerlessBindingsRequest() (request *DescribeRabbitMQServerlessBindingsRequest) {
    request = &DescribeRabbitMQServerlessBindingsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trabbit", APIVersion, "DescribeRabbitMQServerlessBindings")
    
    
    return
}

func NewDescribeRabbitMQServerlessBindingsResponse() (response *DescribeRabbitMQServerlessBindingsResponse) {
    response = &DescribeRabbitMQServerlessBindingsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRabbitMQServerlessBindings
// 获取路由关系列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeRabbitMQServerlessBindings(request *DescribeRabbitMQServerlessBindingsRequest) (response *DescribeRabbitMQServerlessBindingsResponse, err error) {
    return c.DescribeRabbitMQServerlessBindingsWithContext(context.Background(), request)
}

// DescribeRabbitMQServerlessBindings
// 获取路由关系列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeRabbitMQServerlessBindingsWithContext(ctx context.Context, request *DescribeRabbitMQServerlessBindingsRequest) (response *DescribeRabbitMQServerlessBindingsResponse, err error) {
    if request == nil {
        request = NewDescribeRabbitMQServerlessBindingsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRabbitMQServerlessBindings require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRabbitMQServerlessBindingsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRabbitMQServerlessConnectionRequest() (request *DescribeRabbitMQServerlessConnectionRequest) {
    request = &DescribeRabbitMQServerlessConnectionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trabbit", APIVersion, "DescribeRabbitMQServerlessConnection")
    
    
    return
}

func NewDescribeRabbitMQServerlessConnectionResponse() (response *DescribeRabbitMQServerlessConnectionResponse) {
    response = &DescribeRabbitMQServerlessConnectionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRabbitMQServerlessConnection
// 查询RabbitMQ连接列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeRabbitMQServerlessConnection(request *DescribeRabbitMQServerlessConnectionRequest) (response *DescribeRabbitMQServerlessConnectionResponse, err error) {
    return c.DescribeRabbitMQServerlessConnectionWithContext(context.Background(), request)
}

// DescribeRabbitMQServerlessConnection
// 查询RabbitMQ连接列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeRabbitMQServerlessConnectionWithContext(ctx context.Context, request *DescribeRabbitMQServerlessConnectionRequest) (response *DescribeRabbitMQServerlessConnectionResponse, err error) {
    if request == nil {
        request = NewDescribeRabbitMQServerlessConnectionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRabbitMQServerlessConnection require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRabbitMQServerlessConnectionResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRabbitMQServerlessConsumersRequest() (request *DescribeRabbitMQServerlessConsumersRequest) {
    request = &DescribeRabbitMQServerlessConsumersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trabbit", APIVersion, "DescribeRabbitMQServerlessConsumers")
    
    
    return
}

func NewDescribeRabbitMQServerlessConsumersResponse() (response *DescribeRabbitMQServerlessConsumersResponse) {
    response = &DescribeRabbitMQServerlessConsumersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRabbitMQServerlessConsumers
// 查询RabbitMQ队列消费者列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeRabbitMQServerlessConsumers(request *DescribeRabbitMQServerlessConsumersRequest) (response *DescribeRabbitMQServerlessConsumersResponse, err error) {
    return c.DescribeRabbitMQServerlessConsumersWithContext(context.Background(), request)
}

// DescribeRabbitMQServerlessConsumers
// 查询RabbitMQ队列消费者列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeRabbitMQServerlessConsumersWithContext(ctx context.Context, request *DescribeRabbitMQServerlessConsumersRequest) (response *DescribeRabbitMQServerlessConsumersResponse, err error) {
    if request == nil {
        request = NewDescribeRabbitMQServerlessConsumersRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRabbitMQServerlessConsumers require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRabbitMQServerlessConsumersResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRabbitMQServerlessExchangeDetailRequest() (request *DescribeRabbitMQServerlessExchangeDetailRequest) {
    request = &DescribeRabbitMQServerlessExchangeDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trabbit", APIVersion, "DescribeRabbitMQServerlessExchangeDetail")
    
    
    return
}

func NewDescribeRabbitMQServerlessExchangeDetailResponse() (response *DescribeRabbitMQServerlessExchangeDetailResponse) {
    response = &DescribeRabbitMQServerlessExchangeDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRabbitMQServerlessExchangeDetail
// 查询RabbitMQ exchange 详情
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeRabbitMQServerlessExchangeDetail(request *DescribeRabbitMQServerlessExchangeDetailRequest) (response *DescribeRabbitMQServerlessExchangeDetailResponse, err error) {
    return c.DescribeRabbitMQServerlessExchangeDetailWithContext(context.Background(), request)
}

// DescribeRabbitMQServerlessExchangeDetail
// 查询RabbitMQ exchange 详情
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeRabbitMQServerlessExchangeDetailWithContext(ctx context.Context, request *DescribeRabbitMQServerlessExchangeDetailRequest) (response *DescribeRabbitMQServerlessExchangeDetailResponse, err error) {
    if request == nil {
        request = NewDescribeRabbitMQServerlessExchangeDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRabbitMQServerlessExchangeDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRabbitMQServerlessExchangeDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRabbitMQServerlessExchangesRequest() (request *DescribeRabbitMQServerlessExchangesRequest) {
    request = &DescribeRabbitMQServerlessExchangesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trabbit", APIVersion, "DescribeRabbitMQServerlessExchanges")
    
    
    return
}

func NewDescribeRabbitMQServerlessExchangesResponse() (response *DescribeRabbitMQServerlessExchangesResponse) {
    response = &DescribeRabbitMQServerlessExchangesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRabbitMQServerlessExchanges
// 查询RabbitMQ exchange 列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeRabbitMQServerlessExchanges(request *DescribeRabbitMQServerlessExchangesRequest) (response *DescribeRabbitMQServerlessExchangesResponse, err error) {
    return c.DescribeRabbitMQServerlessExchangesWithContext(context.Background(), request)
}

// DescribeRabbitMQServerlessExchanges
// 查询RabbitMQ exchange 列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeRabbitMQServerlessExchangesWithContext(ctx context.Context, request *DescribeRabbitMQServerlessExchangesRequest) (response *DescribeRabbitMQServerlessExchangesResponse, err error) {
    if request == nil {
        request = NewDescribeRabbitMQServerlessExchangesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRabbitMQServerlessExchanges require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRabbitMQServerlessExchangesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRabbitMQServerlessInstanceRequest() (request *DescribeRabbitMQServerlessInstanceRequest) {
    request = &DescribeRabbitMQServerlessInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trabbit", APIVersion, "DescribeRabbitMQServerlessInstance")
    
    
    return
}

func NewDescribeRabbitMQServerlessInstanceResponse() (response *DescribeRabbitMQServerlessInstanceResponse) {
    response = &DescribeRabbitMQServerlessInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRabbitMQServerlessInstance
// 获取单个RabbitMQ专享实例信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeRabbitMQServerlessInstance(request *DescribeRabbitMQServerlessInstanceRequest) (response *DescribeRabbitMQServerlessInstanceResponse, err error) {
    return c.DescribeRabbitMQServerlessInstanceWithContext(context.Background(), request)
}

// DescribeRabbitMQServerlessInstance
// 获取单个RabbitMQ专享实例信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeRabbitMQServerlessInstanceWithContext(ctx context.Context, request *DescribeRabbitMQServerlessInstanceRequest) (response *DescribeRabbitMQServerlessInstanceResponse, err error) {
    if request == nil {
        request = NewDescribeRabbitMQServerlessInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRabbitMQServerlessInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRabbitMQServerlessInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRabbitMQServerlessPermissionRequest() (request *DescribeRabbitMQServerlessPermissionRequest) {
    request = &DescribeRabbitMQServerlessPermissionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trabbit", APIVersion, "DescribeRabbitMQServerlessPermission")
    
    
    return
}

func NewDescribeRabbitMQServerlessPermissionResponse() (response *DescribeRabbitMQServerlessPermissionResponse) {
    response = &DescribeRabbitMQServerlessPermissionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRabbitMQServerlessPermission
// 查询RabbitMQ权限列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeRabbitMQServerlessPermission(request *DescribeRabbitMQServerlessPermissionRequest) (response *DescribeRabbitMQServerlessPermissionResponse, err error) {
    return c.DescribeRabbitMQServerlessPermissionWithContext(context.Background(), request)
}

// DescribeRabbitMQServerlessPermission
// 查询RabbitMQ权限列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeRabbitMQServerlessPermissionWithContext(ctx context.Context, request *DescribeRabbitMQServerlessPermissionRequest) (response *DescribeRabbitMQServerlessPermissionResponse, err error) {
    if request == nil {
        request = NewDescribeRabbitMQServerlessPermissionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRabbitMQServerlessPermission require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRabbitMQServerlessPermissionResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRabbitMQServerlessQueueDetailRequest() (request *DescribeRabbitMQServerlessQueueDetailRequest) {
    request = &DescribeRabbitMQServerlessQueueDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trabbit", APIVersion, "DescribeRabbitMQServerlessQueueDetail")
    
    
    return
}

func NewDescribeRabbitMQServerlessQueueDetailResponse() (response *DescribeRabbitMQServerlessQueueDetailResponse) {
    response = &DescribeRabbitMQServerlessQueueDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRabbitMQServerlessQueueDetail
// 查询RabbitMQ队列详情
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeRabbitMQServerlessQueueDetail(request *DescribeRabbitMQServerlessQueueDetailRequest) (response *DescribeRabbitMQServerlessQueueDetailResponse, err error) {
    return c.DescribeRabbitMQServerlessQueueDetailWithContext(context.Background(), request)
}

// DescribeRabbitMQServerlessQueueDetail
// 查询RabbitMQ队列详情
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeRabbitMQServerlessQueueDetailWithContext(ctx context.Context, request *DescribeRabbitMQServerlessQueueDetailRequest) (response *DescribeRabbitMQServerlessQueueDetailResponse, err error) {
    if request == nil {
        request = NewDescribeRabbitMQServerlessQueueDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRabbitMQServerlessQueueDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRabbitMQServerlessQueueDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRabbitMQServerlessQueuesRequest() (request *DescribeRabbitMQServerlessQueuesRequest) {
    request = &DescribeRabbitMQServerlessQueuesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trabbit", APIVersion, "DescribeRabbitMQServerlessQueues")
    
    
    return
}

func NewDescribeRabbitMQServerlessQueuesResponse() (response *DescribeRabbitMQServerlessQueuesResponse) {
    response = &DescribeRabbitMQServerlessQueuesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRabbitMQServerlessQueues
// 查询RabbitMQ队列列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeRabbitMQServerlessQueues(request *DescribeRabbitMQServerlessQueuesRequest) (response *DescribeRabbitMQServerlessQueuesResponse, err error) {
    return c.DescribeRabbitMQServerlessQueuesWithContext(context.Background(), request)
}

// DescribeRabbitMQServerlessQueues
// 查询RabbitMQ队列列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeRabbitMQServerlessQueuesWithContext(ctx context.Context, request *DescribeRabbitMQServerlessQueuesRequest) (response *DescribeRabbitMQServerlessQueuesResponse, err error) {
    if request == nil {
        request = NewDescribeRabbitMQServerlessQueuesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRabbitMQServerlessQueues require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRabbitMQServerlessQueuesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRabbitMQServerlessUserRequest() (request *DescribeRabbitMQServerlessUserRequest) {
    request = &DescribeRabbitMQServerlessUserRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trabbit", APIVersion, "DescribeRabbitMQServerlessUser")
    
    
    return
}

func NewDescribeRabbitMQServerlessUserResponse() (response *DescribeRabbitMQServerlessUserResponse) {
    response = &DescribeRabbitMQServerlessUserResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRabbitMQServerlessUser
// 查询RabbitMQ用户列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeRabbitMQServerlessUser(request *DescribeRabbitMQServerlessUserRequest) (response *DescribeRabbitMQServerlessUserResponse, err error) {
    return c.DescribeRabbitMQServerlessUserWithContext(context.Background(), request)
}

// DescribeRabbitMQServerlessUser
// 查询RabbitMQ用户列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeRabbitMQServerlessUserWithContext(ctx context.Context, request *DescribeRabbitMQServerlessUserRequest) (response *DescribeRabbitMQServerlessUserResponse, err error) {
    if request == nil {
        request = NewDescribeRabbitMQServerlessUserRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRabbitMQServerlessUser require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRabbitMQServerlessUserResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRabbitMQServerlessVirtualHostRequest() (request *DescribeRabbitMQServerlessVirtualHostRequest) {
    request = &DescribeRabbitMQServerlessVirtualHostRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trabbit", APIVersion, "DescribeRabbitMQServerlessVirtualHost")
    
    
    return
}

func NewDescribeRabbitMQServerlessVirtualHostResponse() (response *DescribeRabbitMQServerlessVirtualHostResponse) {
    response = &DescribeRabbitMQServerlessVirtualHostResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRabbitMQServerlessVirtualHost
// 查询RabbitMQ vhost列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeRabbitMQServerlessVirtualHost(request *DescribeRabbitMQServerlessVirtualHostRequest) (response *DescribeRabbitMQServerlessVirtualHostResponse, err error) {
    return c.DescribeRabbitMQServerlessVirtualHostWithContext(context.Background(), request)
}

// DescribeRabbitMQServerlessVirtualHost
// 查询RabbitMQ vhost列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeRabbitMQServerlessVirtualHostWithContext(ctx context.Context, request *DescribeRabbitMQServerlessVirtualHostRequest) (response *DescribeRabbitMQServerlessVirtualHostResponse, err error) {
    if request == nil {
        request = NewDescribeRabbitMQServerlessVirtualHostRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRabbitMQServerlessVirtualHost require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRabbitMQServerlessVirtualHostResponse()
    err = c.Send(request, response)
    return
}

func NewListRabbitMQServerlessInstancesRequest() (request *ListRabbitMQServerlessInstancesRequest) {
    request = &ListRabbitMQServerlessInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trabbit", APIVersion, "ListRabbitMQServerlessInstances")
    
    
    return
}

func NewListRabbitMQServerlessInstancesResponse() (response *ListRabbitMQServerlessInstancesResponse) {
    response = &ListRabbitMQServerlessInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListRabbitMQServerlessInstances
// 获取实例列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) ListRabbitMQServerlessInstances(request *ListRabbitMQServerlessInstancesRequest) (response *ListRabbitMQServerlessInstancesResponse, err error) {
    return c.ListRabbitMQServerlessInstancesWithContext(context.Background(), request)
}

// ListRabbitMQServerlessInstances
// 获取实例列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) ListRabbitMQServerlessInstancesWithContext(ctx context.Context, request *ListRabbitMQServerlessInstancesRequest) (response *ListRabbitMQServerlessInstancesResponse, err error) {
    if request == nil {
        request = NewListRabbitMQServerlessInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListRabbitMQServerlessInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewListRabbitMQServerlessInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewModifyRabbitMQServerlessExchangeRequest() (request *ModifyRabbitMQServerlessExchangeRequest) {
    request = &ModifyRabbitMQServerlessExchangeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trabbit", APIVersion, "ModifyRabbitMQServerlessExchange")
    
    
    return
}

func NewModifyRabbitMQServerlessExchangeResponse() (response *ModifyRabbitMQServerlessExchangeResponse) {
    response = &ModifyRabbitMQServerlessExchangeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyRabbitMQServerlessExchange
// 修改RabbitMQ exchange
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyRabbitMQServerlessExchange(request *ModifyRabbitMQServerlessExchangeRequest) (response *ModifyRabbitMQServerlessExchangeResponse, err error) {
    return c.ModifyRabbitMQServerlessExchangeWithContext(context.Background(), request)
}

// ModifyRabbitMQServerlessExchange
// 修改RabbitMQ exchange
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyRabbitMQServerlessExchangeWithContext(ctx context.Context, request *ModifyRabbitMQServerlessExchangeRequest) (response *ModifyRabbitMQServerlessExchangeResponse, err error) {
    if request == nil {
        request = NewModifyRabbitMQServerlessExchangeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyRabbitMQServerlessExchange require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyRabbitMQServerlessExchangeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyRabbitMQServerlessInstanceRequest() (request *ModifyRabbitMQServerlessInstanceRequest) {
    request = &ModifyRabbitMQServerlessInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trabbit", APIVersion, "ModifyRabbitMQServerlessInstance")
    
    
    return
}

func NewModifyRabbitMQServerlessInstanceResponse() (response *ModifyRabbitMQServerlessInstanceResponse) {
    response = &ModifyRabbitMQServerlessInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyRabbitMQServerlessInstance
// 修改集群信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyRabbitMQServerlessInstance(request *ModifyRabbitMQServerlessInstanceRequest) (response *ModifyRabbitMQServerlessInstanceResponse, err error) {
    return c.ModifyRabbitMQServerlessInstanceWithContext(context.Background(), request)
}

// ModifyRabbitMQServerlessInstance
// 修改集群信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyRabbitMQServerlessInstanceWithContext(ctx context.Context, request *ModifyRabbitMQServerlessInstanceRequest) (response *ModifyRabbitMQServerlessInstanceResponse, err error) {
    if request == nil {
        request = NewModifyRabbitMQServerlessInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyRabbitMQServerlessInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyRabbitMQServerlessInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewModifyRabbitMQServerlessPermissionRequest() (request *ModifyRabbitMQServerlessPermissionRequest) {
    request = &ModifyRabbitMQServerlessPermissionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trabbit", APIVersion, "ModifyRabbitMQServerlessPermission")
    
    
    return
}

func NewModifyRabbitMQServerlessPermissionResponse() (response *ModifyRabbitMQServerlessPermissionResponse) {
    response = &ModifyRabbitMQServerlessPermissionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyRabbitMQServerlessPermission
// 修改RabbitMQ的权限
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyRabbitMQServerlessPermission(request *ModifyRabbitMQServerlessPermissionRequest) (response *ModifyRabbitMQServerlessPermissionResponse, err error) {
    return c.ModifyRabbitMQServerlessPermissionWithContext(context.Background(), request)
}

// ModifyRabbitMQServerlessPermission
// 修改RabbitMQ的权限
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyRabbitMQServerlessPermissionWithContext(ctx context.Context, request *ModifyRabbitMQServerlessPermissionRequest) (response *ModifyRabbitMQServerlessPermissionResponse, err error) {
    if request == nil {
        request = NewModifyRabbitMQServerlessPermissionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyRabbitMQServerlessPermission require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyRabbitMQServerlessPermissionResponse()
    err = c.Send(request, response)
    return
}

func NewModifyRabbitMQServerlessQueueRequest() (request *ModifyRabbitMQServerlessQueueRequest) {
    request = &ModifyRabbitMQServerlessQueueRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trabbit", APIVersion, "ModifyRabbitMQServerlessQueue")
    
    
    return
}

func NewModifyRabbitMQServerlessQueueResponse() (response *ModifyRabbitMQServerlessQueueResponse) {
    response = &ModifyRabbitMQServerlessQueueResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyRabbitMQServerlessQueue
// 修改RabbitMQ队列
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyRabbitMQServerlessQueue(request *ModifyRabbitMQServerlessQueueRequest) (response *ModifyRabbitMQServerlessQueueResponse, err error) {
    return c.ModifyRabbitMQServerlessQueueWithContext(context.Background(), request)
}

// ModifyRabbitMQServerlessQueue
// 修改RabbitMQ队列
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyRabbitMQServerlessQueueWithContext(ctx context.Context, request *ModifyRabbitMQServerlessQueueRequest) (response *ModifyRabbitMQServerlessQueueResponse, err error) {
    if request == nil {
        request = NewModifyRabbitMQServerlessQueueRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyRabbitMQServerlessQueue require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyRabbitMQServerlessQueueResponse()
    err = c.Send(request, response)
    return
}

func NewModifyRabbitMQServerlessUserRequest() (request *ModifyRabbitMQServerlessUserRequest) {
    request = &ModifyRabbitMQServerlessUserRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trabbit", APIVersion, "ModifyRabbitMQServerlessUser")
    
    
    return
}

func NewModifyRabbitMQServerlessUserResponse() (response *ModifyRabbitMQServerlessUserResponse) {
    response = &ModifyRabbitMQServerlessUserResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyRabbitMQServerlessUser
// 修改RabbitMQ的用户
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyRabbitMQServerlessUser(request *ModifyRabbitMQServerlessUserRequest) (response *ModifyRabbitMQServerlessUserResponse, err error) {
    return c.ModifyRabbitMQServerlessUserWithContext(context.Background(), request)
}

// ModifyRabbitMQServerlessUser
// 修改RabbitMQ的用户
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyRabbitMQServerlessUserWithContext(ctx context.Context, request *ModifyRabbitMQServerlessUserRequest) (response *ModifyRabbitMQServerlessUserResponse, err error) {
    if request == nil {
        request = NewModifyRabbitMQServerlessUserRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyRabbitMQServerlessUser require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyRabbitMQServerlessUserResponse()
    err = c.Send(request, response)
    return
}

func NewModifyRabbitMQServerlessVirtualHostRequest() (request *ModifyRabbitMQServerlessVirtualHostRequest) {
    request = &ModifyRabbitMQServerlessVirtualHostRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trabbit", APIVersion, "ModifyRabbitMQServerlessVirtualHost")
    
    
    return
}

func NewModifyRabbitMQServerlessVirtualHostResponse() (response *ModifyRabbitMQServerlessVirtualHostResponse) {
    response = &ModifyRabbitMQServerlessVirtualHostResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyRabbitMQServerlessVirtualHost
// 修改RabbitMQ的vhost
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyRabbitMQServerlessVirtualHost(request *ModifyRabbitMQServerlessVirtualHostRequest) (response *ModifyRabbitMQServerlessVirtualHostResponse, err error) {
    return c.ModifyRabbitMQServerlessVirtualHostWithContext(context.Background(), request)
}

// ModifyRabbitMQServerlessVirtualHost
// 修改RabbitMQ的vhost
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyRabbitMQServerlessVirtualHostWithContext(ctx context.Context, request *ModifyRabbitMQServerlessVirtualHostRequest) (response *ModifyRabbitMQServerlessVirtualHostResponse, err error) {
    if request == nil {
        request = NewModifyRabbitMQServerlessVirtualHostRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyRabbitMQServerlessVirtualHost require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyRabbitMQServerlessVirtualHostResponse()
    err = c.Send(request, response)
    return
}
