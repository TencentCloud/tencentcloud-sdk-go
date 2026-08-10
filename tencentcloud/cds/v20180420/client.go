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

package v20180420

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2018-04-20"

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


func NewCreateReportPdfRequest() (request *CreateReportPdfRequest) {
    request = &CreateReportPdfRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cds", APIVersion, "CreateReportPdf")
    
    
    return
}

func NewCreateReportPdfResponse() (response *CreateReportPdfResponse) {
    response = &CreateReportPdfResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateReportPdf
// 下载报表PDF
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DASBAMOUNTNOTENOUGH = "FailedOperation.DasbAmountNotEnough"
//  FAILEDOPERATION_DASBERRORCODE = "FailedOperation.DasbErrorCode"
//  FAILEDOPERATION_DASBINVALIDSECRETID = "FailedOperation.DasbInvalidSecretId"
//  FAILEDOPERATION_DASBINVALIDSECRETKEY = "FailedOperation.DasbInvalidSecretKey"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_RESOURCEID = "InvalidParameter.ResourceId"
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
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateReportPdf(request *CreateReportPdfRequest) (response *CreateReportPdfResponse, err error) {
    return c.CreateReportPdfWithContext(context.Background(), request)
}

// CreateReportPdf
// 下载报表PDF
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DASBAMOUNTNOTENOUGH = "FailedOperation.DasbAmountNotEnough"
//  FAILEDOPERATION_DASBERRORCODE = "FailedOperation.DasbErrorCode"
//  FAILEDOPERATION_DASBINVALIDSECRETID = "FailedOperation.DasbInvalidSecretId"
//  FAILEDOPERATION_DASBINVALIDSECRETKEY = "FailedOperation.DasbInvalidSecretKey"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_RESOURCEID = "InvalidParameter.ResourceId"
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
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateReportPdfWithContext(ctx context.Context, request *CreateReportPdfRequest) (response *CreateReportPdfResponse, err error) {
    if request == nil {
        request = NewCreateReportPdfRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cds", APIVersion, "CreateReportPdf")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateReportPdf require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateReportPdfResponse()
    err = c.Send(request, response)
    return
}

func NewCreateTimerReportRequest() (request *CreateTimerReportRequest) {
    request = &CreateTimerReportRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cds", APIVersion, "CreateTimerReport")
    
    
    return
}

func NewCreateTimerReportResponse() (response *CreateTimerReportResponse) {
    response = &CreateTimerReportResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateTimerReport
// 新建报表任务
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DASBAMOUNTNOTENOUGH = "FailedOperation.DasbAmountNotEnough"
//  FAILEDOPERATION_DASBERRORCODE = "FailedOperation.DasbErrorCode"
//  FAILEDOPERATION_DASBINVALIDSECRETID = "FailedOperation.DasbInvalidSecretId"
//  FAILEDOPERATION_DASBINVALIDSECRETKEY = "FailedOperation.DasbInvalidSecretKey"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_RESOURCEID = "InvalidParameter.ResourceId"
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
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateTimerReport(request *CreateTimerReportRequest) (response *CreateTimerReportResponse, err error) {
    return c.CreateTimerReportWithContext(context.Background(), request)
}

// CreateTimerReport
// 新建报表任务
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DASBAMOUNTNOTENOUGH = "FailedOperation.DasbAmountNotEnough"
//  FAILEDOPERATION_DASBERRORCODE = "FailedOperation.DasbErrorCode"
//  FAILEDOPERATION_DASBINVALIDSECRETID = "FailedOperation.DasbInvalidSecretId"
//  FAILEDOPERATION_DASBINVALIDSECRETKEY = "FailedOperation.DasbInvalidSecretKey"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_RESOURCEID = "InvalidParameter.ResourceId"
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
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateTimerReportWithContext(ctx context.Context, request *CreateTimerReportRequest) (response *CreateTimerReportResponse, err error) {
    if request == nil {
        request = NewCreateTimerReportRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cds", APIVersion, "CreateTimerReport")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateTimerReport require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateTimerReportResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAssetsListRequest() (request *DescribeAssetsListRequest) {
    request = &DescribeAssetsListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cds", APIVersion, "DescribeAssetsList")
    
    
    return
}

func NewDescribeAssetsListResponse() (response *DescribeAssetsListResponse) {
    response = &DescribeAssetsListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAssetsList
// 查询资产列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DASBAMOUNTNOTENOUGH = "FailedOperation.DasbAmountNotEnough"
//  FAILEDOPERATION_DASBERRORCODE = "FailedOperation.DasbErrorCode"
//  FAILEDOPERATION_DASBINVALIDSECRETID = "FailedOperation.DasbInvalidSecretId"
//  FAILEDOPERATION_DASBINVALIDSECRETKEY = "FailedOperation.DasbInvalidSecretKey"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_RESOURCEID = "InvalidParameter.ResourceId"
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
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAssetsList(request *DescribeAssetsListRequest) (response *DescribeAssetsListResponse, err error) {
    return c.DescribeAssetsListWithContext(context.Background(), request)
}

// DescribeAssetsList
// 查询资产列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DASBAMOUNTNOTENOUGH = "FailedOperation.DasbAmountNotEnough"
//  FAILEDOPERATION_DASBERRORCODE = "FailedOperation.DasbErrorCode"
//  FAILEDOPERATION_DASBINVALIDSECRETID = "FailedOperation.DasbInvalidSecretId"
//  FAILEDOPERATION_DASBINVALIDSECRETKEY = "FailedOperation.DasbInvalidSecretKey"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_RESOURCEID = "InvalidParameter.ResourceId"
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
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAssetsListWithContext(ctx context.Context, request *DescribeAssetsListRequest) (response *DescribeAssetsListResponse, err error) {
    if request == nil {
        request = NewDescribeAssetsListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cds", APIVersion, "DescribeAssetsList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAssetsList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAssetsListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDbauditInstanceTypeRequest() (request *DescribeDbauditInstanceTypeRequest) {
    request = &DescribeDbauditInstanceTypeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cds", APIVersion, "DescribeDbauditInstanceType")
    
    
    return
}

func NewDescribeDbauditInstanceTypeResponse() (response *DescribeDbauditInstanceTypeResponse) {
    response = &DescribeDbauditInstanceTypeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDbauditInstanceType
// 本接口 (DescribeDbauditInstanceType) 用于查询可售卖的产品规格列表。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DASBAMOUNTNOTENOUGH = "FailedOperation.DasbAmountNotEnough"
//  FAILEDOPERATION_DASBERRORCODE = "FailedOperation.DasbErrorCode"
//  FAILEDOPERATION_DASBINVALIDSECRETID = "FailedOperation.DasbInvalidSecretId"
//  FAILEDOPERATION_DASBINVALIDSECRETKEY = "FailedOperation.DasbInvalidSecretKey"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_RESOURCEID = "InvalidParameter.ResourceId"
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
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDbauditInstanceType(request *DescribeDbauditInstanceTypeRequest) (response *DescribeDbauditInstanceTypeResponse, err error) {
    return c.DescribeDbauditInstanceTypeWithContext(context.Background(), request)
}

// DescribeDbauditInstanceType
// 本接口 (DescribeDbauditInstanceType) 用于查询可售卖的产品规格列表。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DASBAMOUNTNOTENOUGH = "FailedOperation.DasbAmountNotEnough"
//  FAILEDOPERATION_DASBERRORCODE = "FailedOperation.DasbErrorCode"
//  FAILEDOPERATION_DASBINVALIDSECRETID = "FailedOperation.DasbInvalidSecretId"
//  FAILEDOPERATION_DASBINVALIDSECRETKEY = "FailedOperation.DasbInvalidSecretKey"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_RESOURCEID = "InvalidParameter.ResourceId"
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
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDbauditInstanceTypeWithContext(ctx context.Context, request *DescribeDbauditInstanceTypeRequest) (response *DescribeDbauditInstanceTypeResponse, err error) {
    if request == nil {
        request = NewDescribeDbauditInstanceTypeRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cds", APIVersion, "DescribeDbauditInstanceType")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDbauditInstanceType require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDbauditInstanceTypeResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDbauditInstancesRequest() (request *DescribeDbauditInstancesRequest) {
    request = &DescribeDbauditInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cds", APIVersion, "DescribeDbauditInstances")
    
    
    return
}

func NewDescribeDbauditInstancesResponse() (response *DescribeDbauditInstancesResponse) {
    response = &DescribeDbauditInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDbauditInstances
// 本接口 (DescribeDbauditInstances) 用于查询数据安全审计实例列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DASBAMOUNTNOTENOUGH = "FailedOperation.DasbAmountNotEnough"
//  FAILEDOPERATION_DASBERRORCODE = "FailedOperation.DasbErrorCode"
//  FAILEDOPERATION_DASBINVALIDSECRETID = "FailedOperation.DasbInvalidSecretId"
//  FAILEDOPERATION_DASBINVALIDSECRETKEY = "FailedOperation.DasbInvalidSecretKey"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_RESOURCEID = "InvalidParameter.ResourceId"
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
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDbauditInstances(request *DescribeDbauditInstancesRequest) (response *DescribeDbauditInstancesResponse, err error) {
    return c.DescribeDbauditInstancesWithContext(context.Background(), request)
}

// DescribeDbauditInstances
// 本接口 (DescribeDbauditInstances) 用于查询数据安全审计实例列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DASBAMOUNTNOTENOUGH = "FailedOperation.DasbAmountNotEnough"
//  FAILEDOPERATION_DASBERRORCODE = "FailedOperation.DasbErrorCode"
//  FAILEDOPERATION_DASBINVALIDSECRETID = "FailedOperation.DasbInvalidSecretId"
//  FAILEDOPERATION_DASBINVALIDSECRETKEY = "FailedOperation.DasbInvalidSecretKey"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_RESOURCEID = "InvalidParameter.ResourceId"
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
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDbauditInstancesWithContext(ctx context.Context, request *DescribeDbauditInstancesRequest) (response *DescribeDbauditInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeDbauditInstancesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cds", APIVersion, "DescribeDbauditInstances")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDbauditInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDbauditInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDbauditUsedRegionsRequest() (request *DescribeDbauditUsedRegionsRequest) {
    request = &DescribeDbauditUsedRegionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cds", APIVersion, "DescribeDbauditUsedRegions")
    
    
    return
}

func NewDescribeDbauditUsedRegionsResponse() (response *DescribeDbauditUsedRegionsResponse) {
    response = &DescribeDbauditUsedRegionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDbauditUsedRegions
// 本接口 (DescribeDbauditUsedRegions) 用于查询可售卖地域列表。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DASBAMOUNTNOTENOUGH = "FailedOperation.DasbAmountNotEnough"
//  FAILEDOPERATION_DASBERRORCODE = "FailedOperation.DasbErrorCode"
//  FAILEDOPERATION_DASBINVALIDSECRETID = "FailedOperation.DasbInvalidSecretId"
//  FAILEDOPERATION_DASBINVALIDSECRETKEY = "FailedOperation.DasbInvalidSecretKey"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_RESOURCEID = "InvalidParameter.ResourceId"
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
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDbauditUsedRegions(request *DescribeDbauditUsedRegionsRequest) (response *DescribeDbauditUsedRegionsResponse, err error) {
    return c.DescribeDbauditUsedRegionsWithContext(context.Background(), request)
}

// DescribeDbauditUsedRegions
// 本接口 (DescribeDbauditUsedRegions) 用于查询可售卖地域列表。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DASBAMOUNTNOTENOUGH = "FailedOperation.DasbAmountNotEnough"
//  FAILEDOPERATION_DASBERRORCODE = "FailedOperation.DasbErrorCode"
//  FAILEDOPERATION_DASBINVALIDSECRETID = "FailedOperation.DasbInvalidSecretId"
//  FAILEDOPERATION_DASBINVALIDSECRETKEY = "FailedOperation.DasbInvalidSecretKey"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_RESOURCEID = "InvalidParameter.ResourceId"
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
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDbauditUsedRegionsWithContext(ctx context.Context, request *DescribeDbauditUsedRegionsRequest) (response *DescribeDbauditUsedRegionsResponse, err error) {
    if request == nil {
        request = NewDescribeDbauditUsedRegionsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cds", APIVersion, "DescribeDbauditUsedRegions")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDbauditUsedRegions require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDbauditUsedRegionsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeReportListRequest() (request *DescribeReportListRequest) {
    request = &DescribeReportListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cds", APIVersion, "DescribeReportList")
    
    
    return
}

func NewDescribeReportListResponse() (response *DescribeReportListResponse) {
    response = &DescribeReportListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeReportList
// 查询报表列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DASBAMOUNTNOTENOUGH = "FailedOperation.DasbAmountNotEnough"
//  FAILEDOPERATION_DASBERRORCODE = "FailedOperation.DasbErrorCode"
//  FAILEDOPERATION_DASBINVALIDSECRETID = "FailedOperation.DasbInvalidSecretId"
//  FAILEDOPERATION_DASBINVALIDSECRETKEY = "FailedOperation.DasbInvalidSecretKey"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_RESOURCEID = "InvalidParameter.ResourceId"
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
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeReportList(request *DescribeReportListRequest) (response *DescribeReportListResponse, err error) {
    return c.DescribeReportListWithContext(context.Background(), request)
}

// DescribeReportList
// 查询报表列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DASBAMOUNTNOTENOUGH = "FailedOperation.DasbAmountNotEnough"
//  FAILEDOPERATION_DASBERRORCODE = "FailedOperation.DasbErrorCode"
//  FAILEDOPERATION_DASBINVALIDSECRETID = "FailedOperation.DasbInvalidSecretId"
//  FAILEDOPERATION_DASBINVALIDSECRETKEY = "FailedOperation.DasbInvalidSecretKey"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_RESOURCEID = "InvalidParameter.ResourceId"
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
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeReportListWithContext(ctx context.Context, request *DescribeReportListRequest) (response *DescribeReportListResponse, err error) {
    if request == nil {
        request = NewDescribeReportListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cds", APIVersion, "DescribeReportList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeReportList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeReportListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeReportMissionListRequest() (request *DescribeReportMissionListRequest) {
    request = &DescribeReportMissionListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cds", APIVersion, "DescribeReportMissionList")
    
    
    return
}

func NewDescribeReportMissionListResponse() (response *DescribeReportMissionListResponse) {
    response = &DescribeReportMissionListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeReportMissionList
// 查询报表任务列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DASBAMOUNTNOTENOUGH = "FailedOperation.DasbAmountNotEnough"
//  FAILEDOPERATION_DASBERRORCODE = "FailedOperation.DasbErrorCode"
//  FAILEDOPERATION_DASBINVALIDSECRETID = "FailedOperation.DasbInvalidSecretId"
//  FAILEDOPERATION_DASBINVALIDSECRETKEY = "FailedOperation.DasbInvalidSecretKey"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_RESOURCEID = "InvalidParameter.ResourceId"
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
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeReportMissionList(request *DescribeReportMissionListRequest) (response *DescribeReportMissionListResponse, err error) {
    return c.DescribeReportMissionListWithContext(context.Background(), request)
}

// DescribeReportMissionList
// 查询报表任务列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DASBAMOUNTNOTENOUGH = "FailedOperation.DasbAmountNotEnough"
//  FAILEDOPERATION_DASBERRORCODE = "FailedOperation.DasbErrorCode"
//  FAILEDOPERATION_DASBINVALIDSECRETID = "FailedOperation.DasbInvalidSecretId"
//  FAILEDOPERATION_DASBINVALIDSECRETKEY = "FailedOperation.DasbInvalidSecretKey"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_RESOURCEID = "InvalidParameter.ResourceId"
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
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeReportMissionListWithContext(ctx context.Context, request *DescribeReportMissionListRequest) (response *DescribeReportMissionListResponse, err error) {
    if request == nil {
        request = NewDescribeReportMissionListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cds", APIVersion, "DescribeReportMissionList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeReportMissionList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeReportMissionListResponse()
    err = c.Send(request, response)
    return
}

func NewInquiryPriceDbauditInstanceRequest() (request *InquiryPriceDbauditInstanceRequest) {
    request = &InquiryPriceDbauditInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cds", APIVersion, "InquiryPriceDbauditInstance")
    
    
    return
}

func NewInquiryPriceDbauditInstanceResponse() (response *InquiryPriceDbauditInstanceResponse) {
    response = &InquiryPriceDbauditInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// InquiryPriceDbauditInstance
// 用于查询数据安全审计产品实例价格
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DASBAMOUNTNOTENOUGH = "FailedOperation.DasbAmountNotEnough"
//  FAILEDOPERATION_DASBERRORCODE = "FailedOperation.DasbErrorCode"
//  FAILEDOPERATION_DASBINVALIDSECRETID = "FailedOperation.DasbInvalidSecretId"
//  FAILEDOPERATION_DASBINVALIDSECRETKEY = "FailedOperation.DasbInvalidSecretKey"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_RESOURCEID = "InvalidParameter.ResourceId"
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
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) InquiryPriceDbauditInstance(request *InquiryPriceDbauditInstanceRequest) (response *InquiryPriceDbauditInstanceResponse, err error) {
    return c.InquiryPriceDbauditInstanceWithContext(context.Background(), request)
}

// InquiryPriceDbauditInstance
// 用于查询数据安全审计产品实例价格
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DASBAMOUNTNOTENOUGH = "FailedOperation.DasbAmountNotEnough"
//  FAILEDOPERATION_DASBERRORCODE = "FailedOperation.DasbErrorCode"
//  FAILEDOPERATION_DASBINVALIDSECRETID = "FailedOperation.DasbInvalidSecretId"
//  FAILEDOPERATION_DASBINVALIDSECRETKEY = "FailedOperation.DasbInvalidSecretKey"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_RESOURCEID = "InvalidParameter.ResourceId"
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
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) InquiryPriceDbauditInstanceWithContext(ctx context.Context, request *InquiryPriceDbauditInstanceRequest) (response *InquiryPriceDbauditInstanceResponse, err error) {
    if request == nil {
        request = NewInquiryPriceDbauditInstanceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cds", APIVersion, "InquiryPriceDbauditInstance")
    
    if c.GetCredential() == nil {
        return nil, errors.New("InquiryPriceDbauditInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewInquiryPriceDbauditInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDbauditInstancesRenewFlagRequest() (request *ModifyDbauditInstancesRenewFlagRequest) {
    request = &ModifyDbauditInstancesRenewFlagRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cds", APIVersion, "ModifyDbauditInstancesRenewFlag")
    
    
    return
}

func NewModifyDbauditInstancesRenewFlagResponse() (response *ModifyDbauditInstancesRenewFlagResponse) {
    response = &ModifyDbauditInstancesRenewFlagResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyDbauditInstancesRenewFlag
// 本接口 (ModifyDbauditInstancesRenewFlag) 用于修改数据安全审计产品实例续费标识
//
// 可能返回的错误码:
//  INVALIDPARAMETER_RESOURCEID = "InvalidParameter.ResourceId"
//  INVALIDPARAMETER_RESOURCEIDERROR = "InvalidParameter.ResourceIdError"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) ModifyDbauditInstancesRenewFlag(request *ModifyDbauditInstancesRenewFlagRequest) (response *ModifyDbauditInstancesRenewFlagResponse, err error) {
    return c.ModifyDbauditInstancesRenewFlagWithContext(context.Background(), request)
}

// ModifyDbauditInstancesRenewFlag
// 本接口 (ModifyDbauditInstancesRenewFlag) 用于修改数据安全审计产品实例续费标识
//
// 可能返回的错误码:
//  INVALIDPARAMETER_RESOURCEID = "InvalidParameter.ResourceId"
//  INVALIDPARAMETER_RESOURCEIDERROR = "InvalidParameter.ResourceIdError"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) ModifyDbauditInstancesRenewFlagWithContext(ctx context.Context, request *ModifyDbauditInstancesRenewFlagRequest) (response *ModifyDbauditInstancesRenewFlagResponse, err error) {
    if request == nil {
        request = NewModifyDbauditInstancesRenewFlagRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cds", APIVersion, "ModifyDbauditInstancesRenewFlag")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDbauditInstancesRenewFlag require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDbauditInstancesRenewFlagResponse()
    err = c.Send(request, response)
    return
}
