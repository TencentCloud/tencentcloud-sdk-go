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
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCustomer require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateCustomerResponse()
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateJobRecord require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateJobRecordResponse()
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeEnterprises require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeEnterprisesResponse()
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePorts require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePortsResponse()
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeWechatOfficialAccounts require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeWechatOfficialAccountsResponse()
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyCustomer require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyCustomerResponse()
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("StopJobRecord require credential")
    }

    request.SetContext(ctx)
    
    response = NewStopJobRecordResponse()
    err = c.Send(request, response)
    return
}
