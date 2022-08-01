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
// 版权保护-新建拦截接口
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
// 版权保护-新建拦截接口
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
// 品牌经营管家-版权保护模块企业认证接口
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
// 品牌经营管家-版权保护模块企业认证接口
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
// 版权保护-添加作品接口
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
// 版权保护-添加作品接口
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
// 版权保护-拦截申请接口
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
// 版权保护-拦截申请接口
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
// 版权保护-修改监测状态接口
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
// 版权保护-修改监测状态接口
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
// 申请取证
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
// 申请取证
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
// 版权保护-维权申请接口
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
// 版权保护-维权申请接口
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
