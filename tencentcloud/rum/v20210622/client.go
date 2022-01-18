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

package v20210622

import (
    "context"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2021-06-22"

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


func NewCreateProjectRequest() (request *CreateProjectRequest) {
    request = &CreateProjectRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("rum", APIVersion, "CreateProject")
    
    
    return
}

func NewCreateProjectResponse() (response *CreateProjectResponse) {
    response = &CreateProjectResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateProject
// 创建项目（归属于某个团队）
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateProject(request *CreateProjectRequest) (response *CreateProjectResponse, err error) {
    if request == nil {
        request = NewCreateProjectRequest()
    }
    
    response = NewCreateProjectResponse()
    err = c.Send(request, response)
    return
}

// CreateProject
// 创建项目（归属于某个团队）
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateProjectWithContext(ctx context.Context, request *CreateProjectRequest) (response *CreateProjectResponse, err error) {
    if request == nil {
        request = NewCreateProjectRequest()
    }
    request.SetContext(ctx)
    
    response = NewCreateProjectResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDataEventUrlRequest() (request *DescribeDataEventUrlRequest) {
    request = &DescribeDataEventUrlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("rum", APIVersion, "DescribeDataEventUrl")
    
    
    return
}

func NewDescribeDataEventUrlResponse() (response *DescribeDataEventUrlResponse) {
    response = &DescribeDataEventUrlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDataEventUrl
// 获取DescribeDataEventUrl信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  FAILEDOPERATION_INFRASTRUCTUREERROR = "FailedOperation.InfrastructureError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_NOINSTANCE = "ResourceNotFound.NoInstance"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDataEventUrl(request *DescribeDataEventUrlRequest) (response *DescribeDataEventUrlResponse, err error) {
    if request == nil {
        request = NewDescribeDataEventUrlRequest()
    }
    
    response = NewDescribeDataEventUrlResponse()
    err = c.Send(request, response)
    return
}

// DescribeDataEventUrl
// 获取DescribeDataEventUrl信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  FAILEDOPERATION_INFRASTRUCTUREERROR = "FailedOperation.InfrastructureError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_NOINSTANCE = "ResourceNotFound.NoInstance"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDataEventUrlWithContext(ctx context.Context, request *DescribeDataEventUrlRequest) (response *DescribeDataEventUrlResponse, err error) {
    if request == nil {
        request = NewDescribeDataEventUrlRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeDataEventUrlResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDataFetchUrlRequest() (request *DescribeDataFetchUrlRequest) {
    request = &DescribeDataFetchUrlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("rum", APIVersion, "DescribeDataFetchUrl")
    
    
    return
}

func NewDescribeDataFetchUrlResponse() (response *DescribeDataFetchUrlResponse) {
    response = &DescribeDataFetchUrlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDataFetchUrl
// 获取DescribeDataFetchUrl信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  FAILEDOPERATION_INFRASTRUCTUREERROR = "FailedOperation.InfrastructureError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_NOINSTANCE = "ResourceNotFound.NoInstance"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDataFetchUrl(request *DescribeDataFetchUrlRequest) (response *DescribeDataFetchUrlResponse, err error) {
    if request == nil {
        request = NewDescribeDataFetchUrlRequest()
    }
    
    response = NewDescribeDataFetchUrlResponse()
    err = c.Send(request, response)
    return
}

// DescribeDataFetchUrl
// 获取DescribeDataFetchUrl信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  FAILEDOPERATION_INFRASTRUCTUREERROR = "FailedOperation.InfrastructureError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_NOINSTANCE = "ResourceNotFound.NoInstance"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDataFetchUrlWithContext(ctx context.Context, request *DescribeDataFetchUrlRequest) (response *DescribeDataFetchUrlResponse, err error) {
    if request == nil {
        request = NewDescribeDataFetchUrlRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeDataFetchUrlResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDataFetchUrlInfoRequest() (request *DescribeDataFetchUrlInfoRequest) {
    request = &DescribeDataFetchUrlInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("rum", APIVersion, "DescribeDataFetchUrlInfo")
    
    
    return
}

func NewDescribeDataFetchUrlInfoResponse() (response *DescribeDataFetchUrlInfoResponse) {
    response = &DescribeDataFetchUrlInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDataFetchUrlInfo
// 获取DescribeDataFetchUrlInfo信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  FAILEDOPERATION_INFRASTRUCTUREERROR = "FailedOperation.InfrastructureError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_NOINSTANCE = "ResourceNotFound.NoInstance"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDataFetchUrlInfo(request *DescribeDataFetchUrlInfoRequest) (response *DescribeDataFetchUrlInfoResponse, err error) {
    if request == nil {
        request = NewDescribeDataFetchUrlInfoRequest()
    }
    
    response = NewDescribeDataFetchUrlInfoResponse()
    err = c.Send(request, response)
    return
}

// DescribeDataFetchUrlInfo
// 获取DescribeDataFetchUrlInfo信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  FAILEDOPERATION_INFRASTRUCTUREERROR = "FailedOperation.InfrastructureError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_NOINSTANCE = "ResourceNotFound.NoInstance"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDataFetchUrlInfoWithContext(ctx context.Context, request *DescribeDataFetchUrlInfoRequest) (response *DescribeDataFetchUrlInfoResponse, err error) {
    if request == nil {
        request = NewDescribeDataFetchUrlInfoRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeDataFetchUrlInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDataLogUrlStatisticsRequest() (request *DescribeDataLogUrlStatisticsRequest) {
    request = &DescribeDataLogUrlStatisticsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("rum", APIVersion, "DescribeDataLogUrlStatistics")
    
    
    return
}

func NewDescribeDataLogUrlStatisticsResponse() (response *DescribeDataLogUrlStatisticsResponse) {
    response = &DescribeDataLogUrlStatisticsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDataLogUrlStatistics
// 获取LogUrlStatistics信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  FAILEDOPERATION_INFRASTRUCTUREERROR = "FailedOperation.InfrastructureError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
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
func (c *Client) DescribeDataLogUrlStatistics(request *DescribeDataLogUrlStatisticsRequest) (response *DescribeDataLogUrlStatisticsResponse, err error) {
    if request == nil {
        request = NewDescribeDataLogUrlStatisticsRequest()
    }
    
    response = NewDescribeDataLogUrlStatisticsResponse()
    err = c.Send(request, response)
    return
}

// DescribeDataLogUrlStatistics
// 获取LogUrlStatistics信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  FAILEDOPERATION_INFRASTRUCTUREERROR = "FailedOperation.InfrastructureError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
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
func (c *Client) DescribeDataLogUrlStatisticsWithContext(ctx context.Context, request *DescribeDataLogUrlStatisticsRequest) (response *DescribeDataLogUrlStatisticsResponse, err error) {
    if request == nil {
        request = NewDescribeDataLogUrlStatisticsRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeDataLogUrlStatisticsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDataPerformancePageRequest() (request *DescribeDataPerformancePageRequest) {
    request = &DescribeDataPerformancePageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("rum", APIVersion, "DescribeDataPerformancePage")
    
    
    return
}

func NewDescribeDataPerformancePageResponse() (response *DescribeDataPerformancePageResponse) {
    response = &DescribeDataPerformancePageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDataPerformancePage
// 获取PerformancePage信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  FAILEDOPERATION_INFRASTRUCTUREERROR = "FailedOperation.InfrastructureError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_NOINSTANCE = "ResourceNotFound.NoInstance"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDataPerformancePage(request *DescribeDataPerformancePageRequest) (response *DescribeDataPerformancePageResponse, err error) {
    if request == nil {
        request = NewDescribeDataPerformancePageRequest()
    }
    
    response = NewDescribeDataPerformancePageResponse()
    err = c.Send(request, response)
    return
}

// DescribeDataPerformancePage
// 获取PerformancePage信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  FAILEDOPERATION_INFRASTRUCTUREERROR = "FailedOperation.InfrastructureError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_NOINSTANCE = "ResourceNotFound.NoInstance"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDataPerformancePageWithContext(ctx context.Context, request *DescribeDataPerformancePageRequest) (response *DescribeDataPerformancePageResponse, err error) {
    if request == nil {
        request = NewDescribeDataPerformancePageRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeDataPerformancePageResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDataPvUrlStatisticsRequest() (request *DescribeDataPvUrlStatisticsRequest) {
    request = &DescribeDataPvUrlStatisticsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("rum", APIVersion, "DescribeDataPvUrlStatistics")
    
    
    return
}

func NewDescribeDataPvUrlStatisticsResponse() (response *DescribeDataPvUrlStatisticsResponse) {
    response = &DescribeDataPvUrlStatisticsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDataPvUrlStatistics
// 获取DescribeDataPvUrlStatistics信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  FAILEDOPERATION_INFRASTRUCTUREERROR = "FailedOperation.InfrastructureError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_NOINSTANCE = "ResourceNotFound.NoInstance"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDataPvUrlStatistics(request *DescribeDataPvUrlStatisticsRequest) (response *DescribeDataPvUrlStatisticsResponse, err error) {
    if request == nil {
        request = NewDescribeDataPvUrlStatisticsRequest()
    }
    
    response = NewDescribeDataPvUrlStatisticsResponse()
    err = c.Send(request, response)
    return
}

// DescribeDataPvUrlStatistics
// 获取DescribeDataPvUrlStatistics信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  FAILEDOPERATION_INFRASTRUCTUREERROR = "FailedOperation.InfrastructureError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_NOINSTANCE = "ResourceNotFound.NoInstance"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDataPvUrlStatisticsWithContext(ctx context.Context, request *DescribeDataPvUrlStatisticsRequest) (response *DescribeDataPvUrlStatisticsResponse, err error) {
    if request == nil {
        request = NewDescribeDataPvUrlStatisticsRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeDataPvUrlStatisticsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeErrorRequest() (request *DescribeErrorRequest) {
    request = &DescribeErrorRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("rum", APIVersion, "DescribeError")
    
    
    return
}

func NewDescribeErrorResponse() (response *DescribeErrorResponse) {
    response = &DescribeErrorResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeError
// 获取首页错误信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
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
func (c *Client) DescribeError(request *DescribeErrorRequest) (response *DescribeErrorResponse, err error) {
    if request == nil {
        request = NewDescribeErrorRequest()
    }
    
    response = NewDescribeErrorResponse()
    err = c.Send(request, response)
    return
}

// DescribeError
// 获取首页错误信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
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
func (c *Client) DescribeErrorWithContext(ctx context.Context, request *DescribeErrorRequest) (response *DescribeErrorResponse, err error) {
    if request == nil {
        request = NewDescribeErrorRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeErrorResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLogListRequest() (request *DescribeLogListRequest) {
    request = &DescribeLogListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("rum", APIVersion, "DescribeLogList")
    
    
    return
}

func NewDescribeLogListResponse() (response *DescribeLogListResponse) {
    response = &DescribeLogListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeLogList
// 获取项目下的日志列表（实例创建的项目下的日志列表）
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
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
func (c *Client) DescribeLogList(request *DescribeLogListRequest) (response *DescribeLogListResponse, err error) {
    if request == nil {
        request = NewDescribeLogListRequest()
    }
    
    response = NewDescribeLogListResponse()
    err = c.Send(request, response)
    return
}

// DescribeLogList
// 获取项目下的日志列表（实例创建的项目下的日志列表）
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
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
func (c *Client) DescribeLogListWithContext(ctx context.Context, request *DescribeLogListRequest) (response *DescribeLogListResponse, err error) {
    if request == nil {
        request = NewDescribeLogListRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeLogListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeProjectsRequest() (request *DescribeProjectsRequest) {
    request = &DescribeProjectsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("rum", APIVersion, "DescribeProjects")
    
    
    return
}

func NewDescribeProjectsResponse() (response *DescribeProjectsResponse) {
    response = &DescribeProjectsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeProjects
// 获取项目列表（实例创建的团队下的项目列表）
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_NOINSTANCE = "ResourceNotFound.NoInstance"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeProjects(request *DescribeProjectsRequest) (response *DescribeProjectsResponse, err error) {
    if request == nil {
        request = NewDescribeProjectsRequest()
    }
    
    response = NewDescribeProjectsResponse()
    err = c.Send(request, response)
    return
}

// DescribeProjects
// 获取项目列表（实例创建的团队下的项目列表）
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_NOINSTANCE = "ResourceNotFound.NoInstance"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeProjectsWithContext(ctx context.Context, request *DescribeProjectsRequest) (response *DescribeProjectsResponse, err error) {
    if request == nil {
        request = NewDescribeProjectsRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeProjectsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeScoresRequest() (request *DescribeScoresRequest) {
    request = &DescribeScoresRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("rum", APIVersion, "DescribeScores")
    
    
    return
}

func NewDescribeScoresResponse() (response *DescribeScoresResponse) {
    response = &DescribeScoresResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeScores
// 获取首页分数列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
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
func (c *Client) DescribeScores(request *DescribeScoresRequest) (response *DescribeScoresResponse, err error) {
    if request == nil {
        request = NewDescribeScoresRequest()
    }
    
    response = NewDescribeScoresResponse()
    err = c.Send(request, response)
    return
}

// DescribeScores
// 获取首页分数列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
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
func (c *Client) DescribeScoresWithContext(ctx context.Context, request *DescribeScoresRequest) (response *DescribeScoresResponse, err error) {
    if request == nil {
        request = NewDescribeScoresRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeScoresResponse()
    err = c.Send(request, response)
    return
}
