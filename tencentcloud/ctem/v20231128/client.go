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

package v20231128

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2023-11-28"

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


func NewCreateAppRequest() (request *CreateAppRequest) {
    request = &CreateAppRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ctem", APIVersion, "CreateApp")
    
    
    return
}

func NewCreateAppResponse() (response *CreateAppResponse) {
    response = &CreateAppResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateApp
// 添加APP资产
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
func (c *Client) CreateApp(request *CreateAppRequest) (response *CreateAppResponse, err error) {
    return c.CreateAppWithContext(context.Background(), request)
}

// CreateApp
// 添加APP资产
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
func (c *Client) CreateAppWithContext(ctx context.Context, request *CreateAppRequest) (response *CreateAppResponse, err error) {
    if request == nil {
        request = NewCreateAppRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ctem", APIVersion, "CreateApp")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateApp require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAppResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAssetRequest() (request *CreateAssetRequest) {
    request = &CreateAssetRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ctem", APIVersion, "CreateAsset")
    
    
    return
}

func NewCreateAssetResponse() (response *CreateAssetResponse) {
    response = &CreateAssetResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateAsset
// 添加主机资产
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
func (c *Client) CreateAsset(request *CreateAssetRequest) (response *CreateAssetResponse, err error) {
    return c.CreateAssetWithContext(context.Background(), request)
}

// CreateAsset
// 添加主机资产
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
func (c *Client) CreateAssetWithContext(ctx context.Context, request *CreateAssetRequest) (response *CreateAssetResponse, err error) {
    if request == nil {
        request = NewCreateAssetRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ctem", APIVersion, "CreateAsset")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAsset require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAssetResponse()
    err = c.Send(request, response)
    return
}

func NewCreateCustomerRequest() (request *CreateCustomerRequest) {
    request = &CreateCustomerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ctem", APIVersion, "CreateCustomer")
    
    
    return
}

func NewCreateCustomerResponse() (response *CreateCustomerResponse) {
    response = &CreateCustomerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateCustomer
// 创建企业
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
func (c *Client) CreateCustomer(request *CreateCustomerRequest) (response *CreateCustomerResponse, err error) {
    return c.CreateCustomerWithContext(context.Background(), request)
}

// CreateCustomer
// 创建企业
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
func (c *Client) CreateCustomerWithContext(ctx context.Context, request *CreateCustomerRequest) (response *CreateCustomerResponse, err error) {
    if request == nil {
        request = NewCreateCustomerRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ctem", APIVersion, "CreateCustomer")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCustomer require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateCustomerResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDomainRequest() (request *CreateDomainRequest) {
    request = &CreateDomainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ctem", APIVersion, "CreateDomain")
    
    
    return
}

func NewCreateDomainResponse() (response *CreateDomainResponse) {
    response = &CreateDomainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateDomain
// 添加主域名数据
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
func (c *Client) CreateDomain(request *CreateDomainRequest) (response *CreateDomainResponse, err error) {
    return c.CreateDomainWithContext(context.Background(), request)
}

// CreateDomain
// 添加主域名数据
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
func (c *Client) CreateDomainWithContext(ctx context.Context, request *CreateDomainRequest) (response *CreateDomainResponse, err error) {
    if request == nil {
        request = NewCreateDomainRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ctem", APIVersion, "CreateDomain")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateDomain require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateDomainResponse()
    err = c.Send(request, response)
    return
}

func NewCreateEnterpriseRequest() (request *CreateEnterpriseRequest) {
    request = &CreateEnterpriseRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ctem", APIVersion, "CreateEnterprise")
    
    
    return
}

func NewCreateEnterpriseResponse() (response *CreateEnterpriseResponse) {
    response = &CreateEnterpriseResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateEnterprise
// 添加企业架构资产
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
func (c *Client) CreateEnterprise(request *CreateEnterpriseRequest) (response *CreateEnterpriseResponse, err error) {
    return c.CreateEnterpriseWithContext(context.Background(), request)
}

// CreateEnterprise
// 添加企业架构资产
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
func (c *Client) CreateEnterpriseWithContext(ctx context.Context, request *CreateEnterpriseRequest) (response *CreateEnterpriseResponse, err error) {
    if request == nil {
        request = NewCreateEnterpriseRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ctem", APIVersion, "CreateEnterprise")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateEnterprise require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateEnterpriseResponse()
    err = c.Send(request, response)
    return
}

func NewCreateHttpRequest() (request *CreateHttpRequest) {
    request = &CreateHttpRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ctem", APIVersion, "CreateHttp")
    
    
    return
}

func NewCreateHttpResponse() (response *CreateHttpResponse) {
    response = &CreateHttpResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateHttp
// 添加网站资产
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
func (c *Client) CreateHttp(request *CreateHttpRequest) (response *CreateHttpResponse, err error) {
    return c.CreateHttpWithContext(context.Background(), request)
}

// CreateHttp
// 添加网站资产
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
func (c *Client) CreateHttpWithContext(ctx context.Context, request *CreateHttpRequest) (response *CreateHttpResponse, err error) {
    if request == nil {
        request = NewCreateHttpRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ctem", APIVersion, "CreateHttp")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateHttp require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateHttpResponse()
    err = c.Send(request, response)
    return
}

func NewCreateJobRecordRequest() (request *CreateJobRecordRequest) {
    request = &CreateJobRecordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ctem", APIVersion, "CreateJobRecord")
    
    
    return
}

func NewCreateJobRecordResponse() (response *CreateJobRecordResponse) {
    response = &CreateJobRecordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateJobRecord
// 启动测绘
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
func (c *Client) CreateJobRecord(request *CreateJobRecordRequest) (response *CreateJobRecordResponse, err error) {
    return c.CreateJobRecordWithContext(context.Background(), request)
}

// CreateJobRecord
// 启动测绘
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
func (c *Client) CreateJobRecordWithContext(ctx context.Context, request *CreateJobRecordRequest) (response *CreateJobRecordResponse, err error) {
    if request == nil {
        request = NewCreateJobRecordRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ctem", APIVersion, "CreateJobRecord")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateJobRecord require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateJobRecordResponse()
    err = c.Send(request, response)
    return
}

func NewCreateManageRequest() (request *CreateManageRequest) {
    request = &CreateManageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ctem", APIVersion, "CreateManage")
    
    
    return
}

func NewCreateManageResponse() (response *CreateManageResponse) {
    response = &CreateManageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateManage
// 添加后台数据
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
func (c *Client) CreateManage(request *CreateManageRequest) (response *CreateManageResponse, err error) {
    return c.CreateManageWithContext(context.Background(), request)
}

// CreateManage
// 添加后台数据
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
func (c *Client) CreateManageWithContext(ctx context.Context, request *CreateManageRequest) (response *CreateManageResponse, err error) {
    if request == nil {
        request = NewCreateManageRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ctem", APIVersion, "CreateManage")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateManage require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateManageResponse()
    err = c.Send(request, response)
    return
}

func NewCreatePortRequest() (request *CreatePortRequest) {
    request = &CreatePortRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ctem", APIVersion, "CreatePort")
    
    
    return
}

func NewCreatePortResponse() (response *CreatePortResponse) {
    response = &CreatePortResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreatePort
// 添加端口服务资产
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
func (c *Client) CreatePort(request *CreatePortRequest) (response *CreatePortResponse, err error) {
    return c.CreatePortWithContext(context.Background(), request)
}

// CreatePort
// 添加端口服务资产
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
func (c *Client) CreatePortWithContext(ctx context.Context, request *CreatePortRequest) (response *CreatePortResponse, err error) {
    if request == nil {
        request = NewCreatePortRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ctem", APIVersion, "CreatePort")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreatePort require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreatePortResponse()
    err = c.Send(request, response)
    return
}

func NewCreateSeedsRequest() (request *CreateSeedsRequest) {
    request = &CreateSeedsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ctem", APIVersion, "CreateSeeds")
    
    
    return
}

func NewCreateSeedsResponse() (response *CreateSeedsResponse) {
    response = &CreateSeedsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateSeeds
// 创建种子
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
func (c *Client) CreateSeeds(request *CreateSeedsRequest) (response *CreateSeedsResponse, err error) {
    return c.CreateSeedsWithContext(context.Background(), request)
}

// CreateSeeds
// 创建种子
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
func (c *Client) CreateSeedsWithContext(ctx context.Context, request *CreateSeedsRequest) (response *CreateSeedsResponse, err error) {
    if request == nil {
        request = NewCreateSeedsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ctem", APIVersion, "CreateSeeds")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateSeeds require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateSeedsResponse()
    err = c.Send(request, response)
    return
}

func NewCreateSubDomainRequest() (request *CreateSubDomainRequest) {
    request = &CreateSubDomainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ctem", APIVersion, "CreateSubDomain")
    
    
    return
}

func NewCreateSubDomainResponse() (response *CreateSubDomainResponse) {
    response = &CreateSubDomainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateSubDomain
// 添加子域名数据
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
func (c *Client) CreateSubDomain(request *CreateSubDomainRequest) (response *CreateSubDomainResponse, err error) {
    return c.CreateSubDomainWithContext(context.Background(), request)
}

// CreateSubDomain
// 添加子域名数据
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
func (c *Client) CreateSubDomainWithContext(ctx context.Context, request *CreateSubDomainRequest) (response *CreateSubDomainResponse, err error) {
    if request == nil {
        request = NewCreateSubDomainRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ctem", APIVersion, "CreateSubDomain")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateSubDomain require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateSubDomainResponse()
    err = c.Send(request, response)
    return
}

func NewCreateSuspiciousAssetRequest() (request *CreateSuspiciousAssetRequest) {
    request = &CreateSuspiciousAssetRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ctem", APIVersion, "CreateSuspiciousAsset")
    
    
    return
}

func NewCreateSuspiciousAssetResponse() (response *CreateSuspiciousAssetResponse) {
    response = &CreateSuspiciousAssetResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateSuspiciousAsset
// 添加影子资产
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
func (c *Client) CreateSuspiciousAsset(request *CreateSuspiciousAssetRequest) (response *CreateSuspiciousAssetResponse, err error) {
    return c.CreateSuspiciousAssetWithContext(context.Background(), request)
}

// CreateSuspiciousAsset
// 添加影子资产
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
func (c *Client) CreateSuspiciousAssetWithContext(ctx context.Context, request *CreateSuspiciousAssetRequest) (response *CreateSuspiciousAssetResponse, err error) {
    if request == nil {
        request = NewCreateSuspiciousAssetRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ctem", APIVersion, "CreateSuspiciousAsset")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateSuspiciousAsset require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateSuspiciousAssetResponse()
    err = c.Send(request, response)
    return
}

func NewCreateWechatAppletRequest() (request *CreateWechatAppletRequest) {
    request = &CreateWechatAppletRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ctem", APIVersion, "CreateWechatApplet")
    
    
    return
}

func NewCreateWechatAppletResponse() (response *CreateWechatAppletResponse) {
    response = &CreateWechatAppletResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateWechatApplet
// 添加微信小程序资产
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
func (c *Client) CreateWechatApplet(request *CreateWechatAppletRequest) (response *CreateWechatAppletResponse, err error) {
    return c.CreateWechatAppletWithContext(context.Background(), request)
}

// CreateWechatApplet
// 添加微信小程序资产
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
func (c *Client) CreateWechatAppletWithContext(ctx context.Context, request *CreateWechatAppletRequest) (response *CreateWechatAppletResponse, err error) {
    if request == nil {
        request = NewCreateWechatAppletRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ctem", APIVersion, "CreateWechatApplet")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateWechatApplet require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateWechatAppletResponse()
    err = c.Send(request, response)
    return
}

func NewCreateWechatOfficialAccountRequest() (request *CreateWechatOfficialAccountRequest) {
    request = &CreateWechatOfficialAccountRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ctem", APIVersion, "CreateWechatOfficialAccount")
    
    
    return
}

func NewCreateWechatOfficialAccountResponse() (response *CreateWechatOfficialAccountResponse) {
    response = &CreateWechatOfficialAccountResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateWechatOfficialAccount
// 添加微信公众号资产
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
func (c *Client) CreateWechatOfficialAccount(request *CreateWechatOfficialAccountRequest) (response *CreateWechatOfficialAccountResponse, err error) {
    return c.CreateWechatOfficialAccountWithContext(context.Background(), request)
}

// CreateWechatOfficialAccount
// 添加微信公众号资产
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
func (c *Client) CreateWechatOfficialAccountWithContext(ctx context.Context, request *CreateWechatOfficialAccountRequest) (response *CreateWechatOfficialAccountResponse, err error) {
    if request == nil {
        request = NewCreateWechatOfficialAccountRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ctem", APIVersion, "CreateWechatOfficialAccount")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateWechatOfficialAccount require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateWechatOfficialAccountResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAppsRequest() (request *DeleteAppsRequest) {
    request = &DeleteAppsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ctem", APIVersion, "DeleteApps")
    
    
    return
}

func NewDeleteAppsResponse() (response *DeleteAppsResponse) {
    response = &DeleteAppsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteApps
// 删除APP数据
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
func (c *Client) DeleteApps(request *DeleteAppsRequest) (response *DeleteAppsResponse, err error) {
    return c.DeleteAppsWithContext(context.Background(), request)
}

// DeleteApps
// 删除APP数据
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
func (c *Client) DeleteAppsWithContext(ctx context.Context, request *DeleteAppsRequest) (response *DeleteAppsResponse, err error) {
    if request == nil {
        request = NewDeleteAppsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ctem", APIVersion, "DeleteApps")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteApps require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteAppsResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAssetsRequest() (request *DeleteAssetsRequest) {
    request = &DeleteAssetsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ctem", APIVersion, "DeleteAssets")
    
    
    return
}

func NewDeleteAssetsResponse() (response *DeleteAssetsResponse) {
    response = &DeleteAssetsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteAssets
// 删除主机资产数据
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
func (c *Client) DeleteAssets(request *DeleteAssetsRequest) (response *DeleteAssetsResponse, err error) {
    return c.DeleteAssetsWithContext(context.Background(), request)
}

// DeleteAssets
// 删除主机资产数据
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
func (c *Client) DeleteAssetsWithContext(ctx context.Context, request *DeleteAssetsRequest) (response *DeleteAssetsResponse, err error) {
    if request == nil {
        request = NewDeleteAssetsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ctem", APIVersion, "DeleteAssets")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteAssets require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteAssetsResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteDomainsRequest() (request *DeleteDomainsRequest) {
    request = &DeleteDomainsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ctem", APIVersion, "DeleteDomains")
    
    
    return
}

func NewDeleteDomainsResponse() (response *DeleteDomainsResponse) {
    response = &DeleteDomainsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteDomains
// 删除主域名数据
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
func (c *Client) DeleteDomains(request *DeleteDomainsRequest) (response *DeleteDomainsResponse, err error) {
    return c.DeleteDomainsWithContext(context.Background(), request)
}

// DeleteDomains
// 删除主域名数据
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
func (c *Client) DeleteDomainsWithContext(ctx context.Context, request *DeleteDomainsRequest) (response *DeleteDomainsResponse, err error) {
    if request == nil {
        request = NewDeleteDomainsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ctem", APIVersion, "DeleteDomains")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteDomains require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteDomainsResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteEnterprisesRequest() (request *DeleteEnterprisesRequest) {
    request = &DeleteEnterprisesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ctem", APIVersion, "DeleteEnterprises")
    
    
    return
}

func NewDeleteEnterprisesResponse() (response *DeleteEnterprisesResponse) {
    response = &DeleteEnterprisesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteEnterprises
// 删除企业架构数据
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
func (c *Client) DeleteEnterprises(request *DeleteEnterprisesRequest) (response *DeleteEnterprisesResponse, err error) {
    return c.DeleteEnterprisesWithContext(context.Background(), request)
}

// DeleteEnterprises
// 删除企业架构数据
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
func (c *Client) DeleteEnterprisesWithContext(ctx context.Context, request *DeleteEnterprisesRequest) (response *DeleteEnterprisesResponse, err error) {
    if request == nil {
        request = NewDeleteEnterprisesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ctem", APIVersion, "DeleteEnterprises")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteEnterprises require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteEnterprisesResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteHttpsRequest() (request *DeleteHttpsRequest) {
    request = &DeleteHttpsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ctem", APIVersion, "DeleteHttps")
    
    
    return
}

func NewDeleteHttpsResponse() (response *DeleteHttpsResponse) {
    response = &DeleteHttpsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteHttps
// 删除网站资产数据
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
func (c *Client) DeleteHttps(request *DeleteHttpsRequest) (response *DeleteHttpsResponse, err error) {
    return c.DeleteHttpsWithContext(context.Background(), request)
}

// DeleteHttps
// 删除网站资产数据
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
func (c *Client) DeleteHttpsWithContext(ctx context.Context, request *DeleteHttpsRequest) (response *DeleteHttpsResponse, err error) {
    if request == nil {
        request = NewDeleteHttpsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ctem", APIVersion, "DeleteHttps")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteHttps require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteHttpsResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteManagesRequest() (request *DeleteManagesRequest) {
    request = &DeleteManagesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ctem", APIVersion, "DeleteManages")
    
    
    return
}

func NewDeleteManagesResponse() (response *DeleteManagesResponse) {
    response = &DeleteManagesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteManages
// 删除后台数据
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
func (c *Client) DeleteManages(request *DeleteManagesRequest) (response *DeleteManagesResponse, err error) {
    return c.DeleteManagesWithContext(context.Background(), request)
}

// DeleteManages
// 删除后台数据
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
func (c *Client) DeleteManagesWithContext(ctx context.Context, request *DeleteManagesRequest) (response *DeleteManagesResponse, err error) {
    if request == nil {
        request = NewDeleteManagesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ctem", APIVersion, "DeleteManages")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteManages require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteManagesResponse()
    err = c.Send(request, response)
    return
}

func NewDeletePortsRequest() (request *DeletePortsRequest) {
    request = &DeletePortsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ctem", APIVersion, "DeletePorts")
    
    
    return
}

func NewDeletePortsResponse() (response *DeletePortsResponse) {
    response = &DeletePortsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeletePorts
// 删除端口数据
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
func (c *Client) DeletePorts(request *DeletePortsRequest) (response *DeletePortsResponse, err error) {
    return c.DeletePortsWithContext(context.Background(), request)
}

// DeletePorts
// 删除端口数据
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
func (c *Client) DeletePortsWithContext(ctx context.Context, request *DeletePortsRequest) (response *DeletePortsResponse, err error) {
    if request == nil {
        request = NewDeletePortsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ctem", APIVersion, "DeletePorts")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeletePorts require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeletePortsResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteSeedsRequest() (request *DeleteSeedsRequest) {
    request = &DeleteSeedsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ctem", APIVersion, "DeleteSeeds")
    
    
    return
}

func NewDeleteSeedsResponse() (response *DeleteSeedsResponse) {
    response = &DeleteSeedsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteSeeds
// 删除种子
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
func (c *Client) DeleteSeeds(request *DeleteSeedsRequest) (response *DeleteSeedsResponse, err error) {
    return c.DeleteSeedsWithContext(context.Background(), request)
}

// DeleteSeeds
// 删除种子
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
func (c *Client) DeleteSeedsWithContext(ctx context.Context, request *DeleteSeedsRequest) (response *DeleteSeedsResponse, err error) {
    if request == nil {
        request = NewDeleteSeedsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ctem", APIVersion, "DeleteSeeds")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteSeeds require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteSeedsResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteSubDomainsRequest() (request *DeleteSubDomainsRequest) {
    request = &DeleteSubDomainsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ctem", APIVersion, "DeleteSubDomains")
    
    
    return
}

func NewDeleteSubDomainsResponse() (response *DeleteSubDomainsResponse) {
    response = &DeleteSubDomainsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteSubDomains
// 删除子域名数据
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
func (c *Client) DeleteSubDomains(request *DeleteSubDomainsRequest) (response *DeleteSubDomainsResponse, err error) {
    return c.DeleteSubDomainsWithContext(context.Background(), request)
}

// DeleteSubDomains
// 删除子域名数据
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
func (c *Client) DeleteSubDomainsWithContext(ctx context.Context, request *DeleteSubDomainsRequest) (response *DeleteSubDomainsResponse, err error) {
    if request == nil {
        request = NewDeleteSubDomainsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ctem", APIVersion, "DeleteSubDomains")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteSubDomains require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteSubDomainsResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteSuspiciousAssetsRequest() (request *DeleteSuspiciousAssetsRequest) {
    request = &DeleteSuspiciousAssetsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ctem", APIVersion, "DeleteSuspiciousAssets")
    
    
    return
}

func NewDeleteSuspiciousAssetsResponse() (response *DeleteSuspiciousAssetsResponse) {
    response = &DeleteSuspiciousAssetsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteSuspiciousAssets
// 删除影子资产数据
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
func (c *Client) DeleteSuspiciousAssets(request *DeleteSuspiciousAssetsRequest) (response *DeleteSuspiciousAssetsResponse, err error) {
    return c.DeleteSuspiciousAssetsWithContext(context.Background(), request)
}

// DeleteSuspiciousAssets
// 删除影子资产数据
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
func (c *Client) DeleteSuspiciousAssetsWithContext(ctx context.Context, request *DeleteSuspiciousAssetsRequest) (response *DeleteSuspiciousAssetsResponse, err error) {
    if request == nil {
        request = NewDeleteSuspiciousAssetsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ctem", APIVersion, "DeleteSuspiciousAssets")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteSuspiciousAssets require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteSuspiciousAssetsResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteWechatAppletsRequest() (request *DeleteWechatAppletsRequest) {
    request = &DeleteWechatAppletsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ctem", APIVersion, "DeleteWechatApplets")
    
    
    return
}

func NewDeleteWechatAppletsResponse() (response *DeleteWechatAppletsResponse) {
    response = &DeleteWechatAppletsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteWechatApplets
// 删除微信小程序数据
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
func (c *Client) DeleteWechatApplets(request *DeleteWechatAppletsRequest) (response *DeleteWechatAppletsResponse, err error) {
    return c.DeleteWechatAppletsWithContext(context.Background(), request)
}

// DeleteWechatApplets
// 删除微信小程序数据
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
func (c *Client) DeleteWechatAppletsWithContext(ctx context.Context, request *DeleteWechatAppletsRequest) (response *DeleteWechatAppletsResponse, err error) {
    if request == nil {
        request = NewDeleteWechatAppletsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ctem", APIVersion, "DeleteWechatApplets")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteWechatApplets require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteWechatAppletsResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteWechatOfficialAccountsRequest() (request *DeleteWechatOfficialAccountsRequest) {
    request = &DeleteWechatOfficialAccountsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ctem", APIVersion, "DeleteWechatOfficialAccounts")
    
    
    return
}

func NewDeleteWechatOfficialAccountsResponse() (response *DeleteWechatOfficialAccountsResponse) {
    response = &DeleteWechatOfficialAccountsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteWechatOfficialAccounts
// 删除微信公众号数据
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
func (c *Client) DeleteWechatOfficialAccounts(request *DeleteWechatOfficialAccountsRequest) (response *DeleteWechatOfficialAccountsResponse, err error) {
    return c.DeleteWechatOfficialAccountsWithContext(context.Background(), request)
}

// DeleteWechatOfficialAccounts
// 删除微信公众号数据
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
func (c *Client) DeleteWechatOfficialAccountsWithContext(ctx context.Context, request *DeleteWechatOfficialAccountsRequest) (response *DeleteWechatOfficialAccountsResponse, err error) {
    if request == nil {
        request = NewDeleteWechatOfficialAccountsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ctem", APIVersion, "DeleteWechatOfficialAccounts")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteWechatOfficialAccounts require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteWechatOfficialAccountsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeApiSecsRequest() (request *DescribeApiSecsRequest) {
    request = &DescribeApiSecsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ctem", APIVersion, "DescribeApiSecs")
    
    
    return
}

func NewDescribeApiSecsResponse() (response *DescribeApiSecsResponse) {
    response = &DescribeApiSecsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeApiSecs
// 查看API安全数据
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
func (c *Client) DescribeApiSecs(request *DescribeApiSecsRequest) (response *DescribeApiSecsResponse, err error) {
    return c.DescribeApiSecsWithContext(context.Background(), request)
}

// DescribeApiSecs
// 查看API安全数据
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
func (c *Client) DescribeApiSecsWithContext(ctx context.Context, request *DescribeApiSecsRequest) (response *DescribeApiSecsResponse, err error) {
    if request == nil {
        request = NewDescribeApiSecsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ctem", APIVersion, "DescribeApiSecs")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeApiSecs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeApiSecsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAppsRequest() (request *DescribeAppsRequest) {
    request = &DescribeAppsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ctem", APIVersion, "DescribeApps")
    
    
    return
}

func NewDescribeAppsResponse() (response *DescribeAppsResponse) {
    response = &DescribeAppsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeApps
// 查看移动端资产
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
func (c *Client) DescribeApps(request *DescribeAppsRequest) (response *DescribeAppsResponse, err error) {
    return c.DescribeAppsWithContext(context.Background(), request)
}

// DescribeApps
// 查看移动端资产
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
func (c *Client) DescribeAppsWithContext(ctx context.Context, request *DescribeAppsRequest) (response *DescribeAppsResponse, err error) {
    if request == nil {
        request = NewDescribeAppsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ctem", APIVersion, "DescribeApps")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeApps require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAppsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAssetsRequest() (request *DescribeAssetsRequest) {
    request = &DescribeAssetsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ctem", APIVersion, "DescribeAssets")
    
    
    return
}

func NewDescribeAssetsResponse() (response *DescribeAssetsResponse) {
    response = &DescribeAssetsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAssets
// 查看主机资产
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
func (c *Client) DescribeAssets(request *DescribeAssetsRequest) (response *DescribeAssetsResponse, err error) {
    return c.DescribeAssetsWithContext(context.Background(), request)
}

// DescribeAssets
// 查看主机资产
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
func (c *Client) DescribeAssetsWithContext(ctx context.Context, request *DescribeAssetsRequest) (response *DescribeAssetsResponse, err error) {
    if request == nil {
        request = NewDescribeAssetsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ctem", APIVersion, "DescribeAssets")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAssets require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAssetsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeConfigsRequest() (request *DescribeConfigsRequest) {
    request = &DescribeConfigsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ctem", APIVersion, "DescribeConfigs")
    
    
    return
}

func NewDescribeConfigsResponse() (response *DescribeConfigsResponse) {
    response = &DescribeConfigsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeConfigs
// 查看目录爆破数据
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
func (c *Client) DescribeConfigs(request *DescribeConfigsRequest) (response *DescribeConfigsResponse, err error) {
    return c.DescribeConfigsWithContext(context.Background(), request)
}

// DescribeConfigs
// 查看目录爆破数据
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
func (c *Client) DescribeConfigsWithContext(ctx context.Context, request *DescribeConfigsRequest) (response *DescribeConfigsResponse, err error) {
    if request == nil {
        request = NewDescribeConfigsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ctem", APIVersion, "DescribeConfigs")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeConfigs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeConfigsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCustomersRequest() (request *DescribeCustomersRequest) {
    request = &DescribeCustomersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ctem", APIVersion, "DescribeCustomers")
    
    
    return
}

func NewDescribeCustomersResponse() (response *DescribeCustomersResponse) {
    response = &DescribeCustomersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCustomers
// 查看企业列表
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
func (c *Client) DescribeCustomers(request *DescribeCustomersRequest) (response *DescribeCustomersResponse, err error) {
    return c.DescribeCustomersWithContext(context.Background(), request)
}

// DescribeCustomers
// 查看企业列表
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
func (c *Client) DescribeCustomersWithContext(ctx context.Context, request *DescribeCustomersRequest) (response *DescribeCustomersResponse, err error) {
    if request == nil {
        request = NewDescribeCustomersRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ctem", APIVersion, "DescribeCustomers")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCustomers require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCustomersResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDarkWebsRequest() (request *DescribeDarkWebsRequest) {
    request = &DescribeDarkWebsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ctem", APIVersion, "DescribeDarkWebs")
    
    
    return
}

func NewDescribeDarkWebsResponse() (response *DescribeDarkWebsResponse) {
    response = &DescribeDarkWebsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDarkWebs
// 查看暗网数据
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
func (c *Client) DescribeDarkWebs(request *DescribeDarkWebsRequest) (response *DescribeDarkWebsResponse, err error) {
    return c.DescribeDarkWebsWithContext(context.Background(), request)
}

// DescribeDarkWebs
// 查看暗网数据
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
func (c *Client) DescribeDarkWebsWithContext(ctx context.Context, request *DescribeDarkWebsRequest) (response *DescribeDarkWebsResponse, err error) {
    if request == nil {
        request = NewDescribeDarkWebsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ctem", APIVersion, "DescribeDarkWebs")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDarkWebs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDarkWebsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDomainsRequest() (request *DescribeDomainsRequest) {
    request = &DescribeDomainsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ctem", APIVersion, "DescribeDomains")
    
    
    return
}

func NewDescribeDomainsResponse() (response *DescribeDomainsResponse) {
    response = &DescribeDomainsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDomains
// 查看主域名数据
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
func (c *Client) DescribeDomains(request *DescribeDomainsRequest) (response *DescribeDomainsResponse, err error) {
    return c.DescribeDomainsWithContext(context.Background(), request)
}

// DescribeDomains
// 查看主域名数据
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
func (c *Client) DescribeDomainsWithContext(ctx context.Context, request *DescribeDomainsRequest) (response *DescribeDomainsResponse, err error) {
    if request == nil {
        request = NewDescribeDomainsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ctem", APIVersion, "DescribeDomains")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDomains require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDomainsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEnterprisesRequest() (request *DescribeEnterprisesRequest) {
    request = &DescribeEnterprisesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ctem", APIVersion, "DescribeEnterprises")
    
    
    return
}

func NewDescribeEnterprisesResponse() (response *DescribeEnterprisesResponse) {
    response = &DescribeEnterprisesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeEnterprises
// 查看企业架构数据
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
func (c *Client) DescribeEnterprises(request *DescribeEnterprisesRequest) (response *DescribeEnterprisesResponse, err error) {
    return c.DescribeEnterprisesWithContext(context.Background(), request)
}

// DescribeEnterprises
// 查看企业架构数据
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
func (c *Client) DescribeEnterprisesWithContext(ctx context.Context, request *DescribeEnterprisesRequest) (response *DescribeEnterprisesResponse, err error) {
    if request == nil {
        request = NewDescribeEnterprisesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ctem", APIVersion, "DescribeEnterprises")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeEnterprises require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeEnterprisesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeFakeAppsRequest() (request *DescribeFakeAppsRequest) {
    request = &DescribeFakeAppsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ctem", APIVersion, "DescribeFakeApps")
    
    
    return
}

func NewDescribeFakeAppsResponse() (response *DescribeFakeAppsResponse) {
    response = &DescribeFakeAppsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeFakeApps
// 查询仿冒应用
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
func (c *Client) DescribeFakeApps(request *DescribeFakeAppsRequest) (response *DescribeFakeAppsResponse, err error) {
    return c.DescribeFakeAppsWithContext(context.Background(), request)
}

// DescribeFakeApps
// 查询仿冒应用
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
func (c *Client) DescribeFakeAppsWithContext(ctx context.Context, request *DescribeFakeAppsRequest) (response *DescribeFakeAppsResponse, err error) {
    if request == nil {
        request = NewDescribeFakeAppsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ctem", APIVersion, "DescribeFakeApps")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeFakeApps require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeFakeAppsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeFakeMiniProgramsRequest() (request *DescribeFakeMiniProgramsRequest) {
    request = &DescribeFakeMiniProgramsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ctem", APIVersion, "DescribeFakeMiniPrograms")
    
    
    return
}

func NewDescribeFakeMiniProgramsResponse() (response *DescribeFakeMiniProgramsResponse) {
    response = &DescribeFakeMiniProgramsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeFakeMiniPrograms
// 查询仿冒小程序
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
func (c *Client) DescribeFakeMiniPrograms(request *DescribeFakeMiniProgramsRequest) (response *DescribeFakeMiniProgramsResponse, err error) {
    return c.DescribeFakeMiniProgramsWithContext(context.Background(), request)
}

// DescribeFakeMiniPrograms
// 查询仿冒小程序
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
func (c *Client) DescribeFakeMiniProgramsWithContext(ctx context.Context, request *DescribeFakeMiniProgramsRequest) (response *DescribeFakeMiniProgramsResponse, err error) {
    if request == nil {
        request = NewDescribeFakeMiniProgramsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ctem", APIVersion, "DescribeFakeMiniPrograms")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeFakeMiniPrograms require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeFakeMiniProgramsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeFakeWebsitesRequest() (request *DescribeFakeWebsitesRequest) {
    request = &DescribeFakeWebsitesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ctem", APIVersion, "DescribeFakeWebsites")
    
    
    return
}

func NewDescribeFakeWebsitesResponse() (response *DescribeFakeWebsitesResponse) {
    response = &DescribeFakeWebsitesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeFakeWebsites
// 查询仿冒网站
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
func (c *Client) DescribeFakeWebsites(request *DescribeFakeWebsitesRequest) (response *DescribeFakeWebsitesResponse, err error) {
    return c.DescribeFakeWebsitesWithContext(context.Background(), request)
}

// DescribeFakeWebsites
// 查询仿冒网站
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
func (c *Client) DescribeFakeWebsitesWithContext(ctx context.Context, request *DescribeFakeWebsitesRequest) (response *DescribeFakeWebsitesResponse, err error) {
    if request == nil {
        request = NewDescribeFakeWebsitesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ctem", APIVersion, "DescribeFakeWebsites")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeFakeWebsites require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeFakeWebsitesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeFakeWechatOfficialsRequest() (request *DescribeFakeWechatOfficialsRequest) {
    request = &DescribeFakeWechatOfficialsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ctem", APIVersion, "DescribeFakeWechatOfficials")
    
    
    return
}

func NewDescribeFakeWechatOfficialsResponse() (response *DescribeFakeWechatOfficialsResponse) {
    response = &DescribeFakeWechatOfficialsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeFakeWechatOfficials
// 查询仿冒公众号
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
func (c *Client) DescribeFakeWechatOfficials(request *DescribeFakeWechatOfficialsRequest) (response *DescribeFakeWechatOfficialsResponse, err error) {
    return c.DescribeFakeWechatOfficialsWithContext(context.Background(), request)
}

// DescribeFakeWechatOfficials
// 查询仿冒公众号
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
func (c *Client) DescribeFakeWechatOfficialsWithContext(ctx context.Context, request *DescribeFakeWechatOfficialsRequest) (response *DescribeFakeWechatOfficialsResponse, err error) {
    if request == nil {
        request = NewDescribeFakeWechatOfficialsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ctem", APIVersion, "DescribeFakeWechatOfficials")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeFakeWechatOfficials require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeFakeWechatOfficialsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeGithubsRequest() (request *DescribeGithubsRequest) {
    request = &DescribeGithubsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ctem", APIVersion, "DescribeGithubs")
    
    
    return
}

func NewDescribeGithubsResponse() (response *DescribeGithubsResponse) {
    response = &DescribeGithubsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeGithubs
// 查看Github泄露数据
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
func (c *Client) DescribeGithubs(request *DescribeGithubsRequest) (response *DescribeGithubsResponse, err error) {
    return c.DescribeGithubsWithContext(context.Background(), request)
}

// DescribeGithubs
// 查看Github泄露数据
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
func (c *Client) DescribeGithubsWithContext(ctx context.Context, request *DescribeGithubsRequest) (response *DescribeGithubsResponse, err error) {
    if request == nil {
        request = NewDescribeGithubsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ctem", APIVersion, "DescribeGithubs")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeGithubs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeGithubsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeHttpsRequest() (request *DescribeHttpsRequest) {
    request = &DescribeHttpsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ctem", APIVersion, "DescribeHttps")
    
    
    return
}

func NewDescribeHttpsResponse() (response *DescribeHttpsResponse) {
    response = &DescribeHttpsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeHttps
// 查看http数据
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
func (c *Client) DescribeHttps(request *DescribeHttpsRequest) (response *DescribeHttpsResponse, err error) {
    return c.DescribeHttpsWithContext(context.Background(), request)
}

// DescribeHttps
// 查看http数据
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
func (c *Client) DescribeHttpsWithContext(ctx context.Context, request *DescribeHttpsRequest) (response *DescribeHttpsResponse, err error) {
    if request == nil {
        request = NewDescribeHttpsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ctem", APIVersion, "DescribeHttps")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeHttps require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeHttpsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeJobRecordDetailsRequest() (request *DescribeJobRecordDetailsRequest) {
    request = &DescribeJobRecordDetailsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ctem", APIVersion, "DescribeJobRecordDetails")
    
    
    return
}

func NewDescribeJobRecordDetailsResponse() (response *DescribeJobRecordDetailsResponse) {
    response = &DescribeJobRecordDetailsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeJobRecordDetails
// 查看链路详情
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
func (c *Client) DescribeJobRecordDetails(request *DescribeJobRecordDetailsRequest) (response *DescribeJobRecordDetailsResponse, err error) {
    return c.DescribeJobRecordDetailsWithContext(context.Background(), request)
}

// DescribeJobRecordDetails
// 查看链路详情
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
func (c *Client) DescribeJobRecordDetailsWithContext(ctx context.Context, request *DescribeJobRecordDetailsRequest) (response *DescribeJobRecordDetailsResponse, err error) {
    if request == nil {
        request = NewDescribeJobRecordDetailsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ctem", APIVersion, "DescribeJobRecordDetails")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeJobRecordDetails require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeJobRecordDetailsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeJobRecordsRequest() (request *DescribeJobRecordsRequest) {
    request = &DescribeJobRecordsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ctem", APIVersion, "DescribeJobRecords")
    
    
    return
}

func NewDescribeJobRecordsResponse() (response *DescribeJobRecordsResponse) {
    response = &DescribeJobRecordsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeJobRecords
// 查看任务运行记录列表
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
func (c *Client) DescribeJobRecords(request *DescribeJobRecordsRequest) (response *DescribeJobRecordsResponse, err error) {
    return c.DescribeJobRecordsWithContext(context.Background(), request)
}

// DescribeJobRecords
// 查看任务运行记录列表
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
func (c *Client) DescribeJobRecordsWithContext(ctx context.Context, request *DescribeJobRecordsRequest) (response *DescribeJobRecordsResponse, err error) {
    if request == nil {
        request = NewDescribeJobRecordsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ctem", APIVersion, "DescribeJobRecords")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeJobRecords require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeJobRecordsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLeakageCodesRequest() (request *DescribeLeakageCodesRequest) {
    request = &DescribeLeakageCodesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ctem", APIVersion, "DescribeLeakageCodes")
    
    
    return
}

func NewDescribeLeakageCodesResponse() (response *DescribeLeakageCodesResponse) {
    response = &DescribeLeakageCodesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeLeakageCodes
// 获取代码泄露数据
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
func (c *Client) DescribeLeakageCodes(request *DescribeLeakageCodesRequest) (response *DescribeLeakageCodesResponse, err error) {
    return c.DescribeLeakageCodesWithContext(context.Background(), request)
}

// DescribeLeakageCodes
// 获取代码泄露数据
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
func (c *Client) DescribeLeakageCodesWithContext(ctx context.Context, request *DescribeLeakageCodesRequest) (response *DescribeLeakageCodesResponse, err error) {
    if request == nil {
        request = NewDescribeLeakageCodesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ctem", APIVersion, "DescribeLeakageCodes")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLeakageCodes require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLeakageCodesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLeakageDatasRequest() (request *DescribeLeakageDatasRequest) {
    request = &DescribeLeakageDatasRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ctem", APIVersion, "DescribeLeakageDatas")
    
    
    return
}

func NewDescribeLeakageDatasResponse() (response *DescribeLeakageDatasResponse) {
    response = &DescribeLeakageDatasResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeLeakageDatas
// 获取数据泄露事件
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
func (c *Client) DescribeLeakageDatas(request *DescribeLeakageDatasRequest) (response *DescribeLeakageDatasResponse, err error) {
    return c.DescribeLeakageDatasWithContext(context.Background(), request)
}

// DescribeLeakageDatas
// 获取数据泄露事件
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
func (c *Client) DescribeLeakageDatasWithContext(ctx context.Context, request *DescribeLeakageDatasRequest) (response *DescribeLeakageDatasResponse, err error) {
    if request == nil {
        request = NewDescribeLeakageDatasRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ctem", APIVersion, "DescribeLeakageDatas")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLeakageDatas require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLeakageDatasResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLeakageEmailsRequest() (request *DescribeLeakageEmailsRequest) {
    request = &DescribeLeakageEmailsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ctem", APIVersion, "DescribeLeakageEmails")
    
    
    return
}

func NewDescribeLeakageEmailsResponse() (response *DescribeLeakageEmailsResponse) {
    response = &DescribeLeakageEmailsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeLeakageEmails
// 获取邮箱泄露数据
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
func (c *Client) DescribeLeakageEmails(request *DescribeLeakageEmailsRequest) (response *DescribeLeakageEmailsResponse, err error) {
    return c.DescribeLeakageEmailsWithContext(context.Background(), request)
}

// DescribeLeakageEmails
// 获取邮箱泄露数据
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
func (c *Client) DescribeLeakageEmailsWithContext(ctx context.Context, request *DescribeLeakageEmailsRequest) (response *DescribeLeakageEmailsResponse, err error) {
    if request == nil {
        request = NewDescribeLeakageEmailsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ctem", APIVersion, "DescribeLeakageEmails")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLeakageEmails require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLeakageEmailsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeManagesRequest() (request *DescribeManagesRequest) {
    request = &DescribeManagesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ctem", APIVersion, "DescribeManages")
    
    
    return
}

func NewDescribeManagesResponse() (response *DescribeManagesResponse) {
    response = &DescribeManagesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeManages
// 查看后台管理数据
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
func (c *Client) DescribeManages(request *DescribeManagesRequest) (response *DescribeManagesResponse, err error) {
    return c.DescribeManagesWithContext(context.Background(), request)
}

// DescribeManages
// 查看后台管理数据
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
func (c *Client) DescribeManagesWithContext(ctx context.Context, request *DescribeManagesRequest) (response *DescribeManagesResponse, err error) {
    if request == nil {
        request = NewDescribeManagesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ctem", APIVersion, "DescribeManages")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeManages require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeManagesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeNetDisksRequest() (request *DescribeNetDisksRequest) {
    request = &DescribeNetDisksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ctem", APIVersion, "DescribeNetDisks")
    
    
    return
}

func NewDescribeNetDisksResponse() (response *DescribeNetDisksResponse) {
    response = &DescribeNetDisksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeNetDisks
// 查看网盘泄露数据
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
func (c *Client) DescribeNetDisks(request *DescribeNetDisksRequest) (response *DescribeNetDisksResponse, err error) {
    return c.DescribeNetDisksWithContext(context.Background(), request)
}

// DescribeNetDisks
// 查看网盘泄露数据
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
func (c *Client) DescribeNetDisksWithContext(ctx context.Context, request *DescribeNetDisksRequest) (response *DescribeNetDisksResponse, err error) {
    if request == nil {
        request = NewDescribeNetDisksRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ctem", APIVersion, "DescribeNetDisks")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeNetDisks require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeNetDisksResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePortsRequest() (request *DescribePortsRequest) {
    request = &DescribePortsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ctem", APIVersion, "DescribePorts")
    
    
    return
}

func NewDescribePortsResponse() (response *DescribePortsResponse) {
    response = &DescribePortsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribePorts
// 查看端口数据
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
func (c *Client) DescribePorts(request *DescribePortsRequest) (response *DescribePortsResponse, err error) {
    return c.DescribePortsWithContext(context.Background(), request)
}

// DescribePorts
// 查看端口数据
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
func (c *Client) DescribePortsWithContext(ctx context.Context, request *DescribePortsRequest) (response *DescribePortsResponse, err error) {
    if request == nil {
        request = NewDescribePortsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ctem", APIVersion, "DescribePorts")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePorts require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePortsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSeedsRequest() (request *DescribeSeedsRequest) {
    request = &DescribeSeedsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ctem", APIVersion, "DescribeSeeds")
    
    
    return
}

func NewDescribeSeedsResponse() (response *DescribeSeedsResponse) {
    response = &DescribeSeedsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSeeds
// 查看种子列表
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
func (c *Client) DescribeSeeds(request *DescribeSeedsRequest) (response *DescribeSeedsResponse, err error) {
    return c.DescribeSeedsWithContext(context.Background(), request)
}

// DescribeSeeds
// 查看种子列表
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
func (c *Client) DescribeSeedsWithContext(ctx context.Context, request *DescribeSeedsRequest) (response *DescribeSeedsResponse, err error) {
    if request == nil {
        request = NewDescribeSeedsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ctem", APIVersion, "DescribeSeeds")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSeeds require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSeedsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSensitiveInfosRequest() (request *DescribeSensitiveInfosRequest) {
    request = &DescribeSensitiveInfosRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ctem", APIVersion, "DescribeSensitiveInfos")
    
    
    return
}

func NewDescribeSensitiveInfosResponse() (response *DescribeSensitiveInfosResponse) {
    response = &DescribeSensitiveInfosResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSensitiveInfos
// 查看敏感信息泄露数据
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
func (c *Client) DescribeSensitiveInfos(request *DescribeSensitiveInfosRequest) (response *DescribeSensitiveInfosResponse, err error) {
    return c.DescribeSensitiveInfosWithContext(context.Background(), request)
}

// DescribeSensitiveInfos
// 查看敏感信息泄露数据
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
func (c *Client) DescribeSensitiveInfosWithContext(ctx context.Context, request *DescribeSensitiveInfosRequest) (response *DescribeSensitiveInfosResponse, err error) {
    if request == nil {
        request = NewDescribeSensitiveInfosRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ctem", APIVersion, "DescribeSensitiveInfos")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSensitiveInfos require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSensitiveInfosResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSubDomainsRequest() (request *DescribeSubDomainsRequest) {
    request = &DescribeSubDomainsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ctem", APIVersion, "DescribeSubDomains")
    
    
    return
}

func NewDescribeSubDomainsResponse() (response *DescribeSubDomainsResponse) {
    response = &DescribeSubDomainsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSubDomains
// 查看子域名数据
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
func (c *Client) DescribeSubDomains(request *DescribeSubDomainsRequest) (response *DescribeSubDomainsResponse, err error) {
    return c.DescribeSubDomainsWithContext(context.Background(), request)
}

// DescribeSubDomains
// 查看子域名数据
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
func (c *Client) DescribeSubDomainsWithContext(ctx context.Context, request *DescribeSubDomainsRequest) (response *DescribeSubDomainsResponse, err error) {
    if request == nil {
        request = NewDescribeSubDomainsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ctem", APIVersion, "DescribeSubDomains")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSubDomains require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSubDomainsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSuspiciousAssetsRequest() (request *DescribeSuspiciousAssetsRequest) {
    request = &DescribeSuspiciousAssetsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ctem", APIVersion, "DescribeSuspiciousAssets")
    
    
    return
}

func NewDescribeSuspiciousAssetsResponse() (response *DescribeSuspiciousAssetsResponse) {
    response = &DescribeSuspiciousAssetsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSuspiciousAssets
// 查看影子资产
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
func (c *Client) DescribeSuspiciousAssets(request *DescribeSuspiciousAssetsRequest) (response *DescribeSuspiciousAssetsResponse, err error) {
    return c.DescribeSuspiciousAssetsWithContext(context.Background(), request)
}

// DescribeSuspiciousAssets
// 查看影子资产
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
func (c *Client) DescribeSuspiciousAssetsWithContext(ctx context.Context, request *DescribeSuspiciousAssetsRequest) (response *DescribeSuspiciousAssetsResponse, err error) {
    if request == nil {
        request = NewDescribeSuspiciousAssetsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ctem", APIVersion, "DescribeSuspiciousAssets")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSuspiciousAssets require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSuspiciousAssetsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVulsRequest() (request *DescribeVulsRequest) {
    request = &DescribeVulsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ctem", APIVersion, "DescribeVuls")
    
    
    return
}

func NewDescribeVulsResponse() (response *DescribeVulsResponse) {
    response = &DescribeVulsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeVuls
// 查看漏洞数据
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
func (c *Client) DescribeVuls(request *DescribeVulsRequest) (response *DescribeVulsResponse, err error) {
    return c.DescribeVulsWithContext(context.Background(), request)
}

// DescribeVuls
// 查看漏洞数据
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
func (c *Client) DescribeVulsWithContext(ctx context.Context, request *DescribeVulsRequest) (response *DescribeVulsResponse, err error) {
    if request == nil {
        request = NewDescribeVulsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ctem", APIVersion, "DescribeVuls")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVuls require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVulsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeWeakPasswordsRequest() (request *DescribeWeakPasswordsRequest) {
    request = &DescribeWeakPasswordsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ctem", APIVersion, "DescribeWeakPasswords")
    
    
    return
}

func NewDescribeWeakPasswordsResponse() (response *DescribeWeakPasswordsResponse) {
    response = &DescribeWeakPasswordsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeWeakPasswords
// 查看弱口令数据
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
func (c *Client) DescribeWeakPasswords(request *DescribeWeakPasswordsRequest) (response *DescribeWeakPasswordsResponse, err error) {
    return c.DescribeWeakPasswordsWithContext(context.Background(), request)
}

// DescribeWeakPasswords
// 查看弱口令数据
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
func (c *Client) DescribeWeakPasswordsWithContext(ctx context.Context, request *DescribeWeakPasswordsRequest) (response *DescribeWeakPasswordsResponse, err error) {
    if request == nil {
        request = NewDescribeWeakPasswordsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ctem", APIVersion, "DescribeWeakPasswords")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeWeakPasswords require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeWeakPasswordsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeWechatAppletsRequest() (request *DescribeWechatAppletsRequest) {
    request = &DescribeWechatAppletsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ctem", APIVersion, "DescribeWechatApplets")
    
    
    return
}

func NewDescribeWechatAppletsResponse() (response *DescribeWechatAppletsResponse) {
    response = &DescribeWechatAppletsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeWechatApplets
// 查看微信小程序
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
func (c *Client) DescribeWechatApplets(request *DescribeWechatAppletsRequest) (response *DescribeWechatAppletsResponse, err error) {
    return c.DescribeWechatAppletsWithContext(context.Background(), request)
}

// DescribeWechatApplets
// 查看微信小程序
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
func (c *Client) DescribeWechatAppletsWithContext(ctx context.Context, request *DescribeWechatAppletsRequest) (response *DescribeWechatAppletsResponse, err error) {
    if request == nil {
        request = NewDescribeWechatAppletsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ctem", APIVersion, "DescribeWechatApplets")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeWechatApplets require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeWechatAppletsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeWechatOfficialAccountsRequest() (request *DescribeWechatOfficialAccountsRequest) {
    request = &DescribeWechatOfficialAccountsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ctem", APIVersion, "DescribeWechatOfficialAccounts")
    
    
    return
}

func NewDescribeWechatOfficialAccountsResponse() (response *DescribeWechatOfficialAccountsResponse) {
    response = &DescribeWechatOfficialAccountsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeWechatOfficialAccounts
// 查看公众号数据
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
func (c *Client) DescribeWechatOfficialAccounts(request *DescribeWechatOfficialAccountsRequest) (response *DescribeWechatOfficialAccountsResponse, err error) {
    return c.DescribeWechatOfficialAccountsWithContext(context.Background(), request)
}

// DescribeWechatOfficialAccounts
// 查看公众号数据
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
func (c *Client) DescribeWechatOfficialAccountsWithContext(ctx context.Context, request *DescribeWechatOfficialAccountsRequest) (response *DescribeWechatOfficialAccountsResponse, err error) {
    if request == nil {
        request = NewDescribeWechatOfficialAccountsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ctem", APIVersion, "DescribeWechatOfficialAccounts")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeWechatOfficialAccounts require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeWechatOfficialAccountsResponse()
    err = c.Send(request, response)
    return
}

func NewIgnoreDataRequest() (request *IgnoreDataRequest) {
    request = &IgnoreDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ctem", APIVersion, "IgnoreData")
    
    
    return
}

func NewIgnoreDataResponse() (response *IgnoreDataResponse) {
    response = &IgnoreDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// IgnoreData
// 忽略数据
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
func (c *Client) IgnoreData(request *IgnoreDataRequest) (response *IgnoreDataResponse, err error) {
    return c.IgnoreDataWithContext(context.Background(), request)
}

// IgnoreData
// 忽略数据
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
func (c *Client) IgnoreDataWithContext(ctx context.Context, request *IgnoreDataRequest) (response *IgnoreDataResponse, err error) {
    if request == nil {
        request = NewIgnoreDataRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ctem", APIVersion, "IgnoreData")
    
    if c.GetCredential() == nil {
        return nil, errors.New("IgnoreData require credential")
    }

    request.SetContext(ctx)
    
    response = NewIgnoreDataResponse()
    err = c.Send(request, response)
    return
}

func NewModifyCustomerRequest() (request *ModifyCustomerRequest) {
    request = &ModifyCustomerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ctem", APIVersion, "ModifyCustomer")
    
    
    return
}

func NewModifyCustomerResponse() (response *ModifyCustomerResponse) {
    response = &ModifyCustomerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyCustomer
// 编辑企业
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
func (c *Client) ModifyCustomer(request *ModifyCustomerRequest) (response *ModifyCustomerResponse, err error) {
    return c.ModifyCustomerWithContext(context.Background(), request)
}

// ModifyCustomer
// 编辑企业
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
func (c *Client) ModifyCustomerWithContext(ctx context.Context, request *ModifyCustomerRequest) (response *ModifyCustomerResponse, err error) {
    if request == nil {
        request = NewModifyCustomerRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ctem", APIVersion, "ModifyCustomer")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyCustomer require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyCustomerResponse()
    err = c.Send(request, response)
    return
}

func NewModifyLabelRequest() (request *ModifyLabelRequest) {
    request = &ModifyLabelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ctem", APIVersion, "ModifyLabel")
    
    
    return
}

func NewModifyLabelResponse() (response *ModifyLabelResponse) {
    response = &ModifyLabelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyLabel
// 修改标签
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
func (c *Client) ModifyLabel(request *ModifyLabelRequest) (response *ModifyLabelResponse, err error) {
    return c.ModifyLabelWithContext(context.Background(), request)
}

// ModifyLabel
// 修改标签
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
func (c *Client) ModifyLabelWithContext(ctx context.Context, request *ModifyLabelRequest) (response *ModifyLabelResponse, err error) {
    if request == nil {
        request = NewModifyLabelRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ctem", APIVersion, "ModifyLabel")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyLabel require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyLabelResponse()
    err = c.Send(request, response)
    return
}

func NewModifySeedStatusRequest() (request *ModifySeedStatusRequest) {
    request = &ModifySeedStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ctem", APIVersion, "ModifySeedStatus")
    
    
    return
}

func NewModifySeedStatusResponse() (response *ModifySeedStatusResponse) {
    response = &ModifySeedStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifySeedStatus
// 修改种子状态
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
func (c *Client) ModifySeedStatus(request *ModifySeedStatusRequest) (response *ModifySeedStatusResponse, err error) {
    return c.ModifySeedStatusWithContext(context.Background(), request)
}

// ModifySeedStatus
// 修改种子状态
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
func (c *Client) ModifySeedStatusWithContext(ctx context.Context, request *ModifySeedStatusRequest) (response *ModifySeedStatusResponse, err error) {
    if request == nil {
        request = NewModifySeedStatusRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ctem", APIVersion, "ModifySeedStatus")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifySeedStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifySeedStatusResponse()
    err = c.Send(request, response)
    return
}

func NewStopJobRecordRequest() (request *StopJobRecordRequest) {
    request = &StopJobRecordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ctem", APIVersion, "StopJobRecord")
    
    
    return
}

func NewStopJobRecordResponse() (response *StopJobRecordResponse) {
    response = &StopJobRecordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// StopJobRecord
// 停止扫描
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
func (c *Client) StopJobRecord(request *StopJobRecordRequest) (response *StopJobRecordResponse, err error) {
    return c.StopJobRecordWithContext(context.Background(), request)
}

// StopJobRecord
// 停止扫描
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
func (c *Client) StopJobRecordWithContext(ctx context.Context, request *StopJobRecordRequest) (response *StopJobRecordResponse, err error) {
    if request == nil {
        request = NewStopJobRecordRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ctem", APIVersion, "StopJobRecord")
    
    if c.GetCredential() == nil {
        return nil, errors.New("StopJobRecord require credential")
    }

    request.SetContext(ctx)
    
    response = NewStopJobRecordResponse()
    err = c.Send(request, response)
    return
}
