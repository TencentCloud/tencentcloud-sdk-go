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

package v20221121

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2022-11-21"

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


func NewAddNewBindRoleUserRequest() (request *AddNewBindRoleUserRequest) {
    request = &AddNewBindRoleUserRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "AddNewBindRoleUser")
    
    
    return
}

func NewAddNewBindRoleUserResponse() (response *AddNewBindRoleUserResponse) {
    response = &AddNewBindRoleUserResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AddNewBindRoleUser
// csip角色授权绑定接口
func (c *Client) AddNewBindRoleUser(request *AddNewBindRoleUserRequest) (response *AddNewBindRoleUserResponse, err error) {
    return c.AddNewBindRoleUserWithContext(context.Background(), request)
}

// AddNewBindRoleUser
// csip角色授权绑定接口
func (c *Client) AddNewBindRoleUserWithContext(ctx context.Context, request *AddNewBindRoleUserRequest) (response *AddNewBindRoleUserResponse, err error) {
    if request == nil {
        request = NewAddNewBindRoleUserRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddNewBindRoleUser require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddNewBindRoleUserResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCVMAssetInfoRequest() (request *DescribeCVMAssetInfoRequest) {
    request = &DescribeCVMAssetInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeCVMAssetInfo")
    
    
    return
}

func NewDescribeCVMAssetInfoResponse() (response *DescribeCVMAssetInfoResponse) {
    response = &DescribeCVMAssetInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCVMAssetInfo
// cvm详情
func (c *Client) DescribeCVMAssetInfo(request *DescribeCVMAssetInfoRequest) (response *DescribeCVMAssetInfoResponse, err error) {
    return c.DescribeCVMAssetInfoWithContext(context.Background(), request)
}

// DescribeCVMAssetInfo
// cvm详情
func (c *Client) DescribeCVMAssetInfoWithContext(ctx context.Context, request *DescribeCVMAssetInfoRequest) (response *DescribeCVMAssetInfoResponse, err error) {
    if request == nil {
        request = NewDescribeCVMAssetInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCVMAssetInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCVMAssetInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCVMAssetsRequest() (request *DescribeCVMAssetsRequest) {
    request = &DescribeCVMAssetsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeCVMAssets")
    
    
    return
}

func NewDescribeCVMAssetsResponse() (response *DescribeCVMAssetsResponse) {
    response = &DescribeCVMAssetsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCVMAssets
// cvm列表
func (c *Client) DescribeCVMAssets(request *DescribeCVMAssetsRequest) (response *DescribeCVMAssetsResponse, err error) {
    return c.DescribeCVMAssetsWithContext(context.Background(), request)
}

// DescribeCVMAssets
// cvm列表
func (c *Client) DescribeCVMAssetsWithContext(ctx context.Context, request *DescribeCVMAssetsRequest) (response *DescribeCVMAssetsResponse, err error) {
    if request == nil {
        request = NewDescribeCVMAssetsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCVMAssets require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCVMAssetsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDbAssetInfoRequest() (request *DescribeDbAssetInfoRequest) {
    request = &DescribeDbAssetInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeDbAssetInfo")
    
    
    return
}

func NewDescribeDbAssetInfoResponse() (response *DescribeDbAssetInfoResponse) {
    response = &DescribeDbAssetInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDbAssetInfo
// db资产详情
func (c *Client) DescribeDbAssetInfo(request *DescribeDbAssetInfoRequest) (response *DescribeDbAssetInfoResponse, err error) {
    return c.DescribeDbAssetInfoWithContext(context.Background(), request)
}

// DescribeDbAssetInfo
// db资产详情
func (c *Client) DescribeDbAssetInfoWithContext(ctx context.Context, request *DescribeDbAssetInfoRequest) (response *DescribeDbAssetInfoResponse, err error) {
    if request == nil {
        request = NewDescribeDbAssetInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDbAssetInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDbAssetInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDbAssetsRequest() (request *DescribeDbAssetsRequest) {
    request = &DescribeDbAssetsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeDbAssets")
    
    
    return
}

func NewDescribeDbAssetsResponse() (response *DescribeDbAssetsResponse) {
    response = &DescribeDbAssetsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDbAssets
// 资产列表
func (c *Client) DescribeDbAssets(request *DescribeDbAssetsRequest) (response *DescribeDbAssetsResponse, err error) {
    return c.DescribeDbAssetsWithContext(context.Background(), request)
}

// DescribeDbAssets
// 资产列表
func (c *Client) DescribeDbAssetsWithContext(ctx context.Context, request *DescribeDbAssetsRequest) (response *DescribeDbAssetsResponse, err error) {
    if request == nil {
        request = NewDescribeDbAssetsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDbAssets require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDbAssetsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeScanReportListRequest() (request *DescribeScanReportListRequest) {
    request = &DescribeScanReportListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeScanReportList")
    
    
    return
}

func NewDescribeScanReportListResponse() (response *DescribeScanReportListResponse) {
    response = &DescribeScanReportListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeScanReportList
// 获取扫描报告列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
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
func (c *Client) DescribeScanReportList(request *DescribeScanReportListRequest) (response *DescribeScanReportListResponse, err error) {
    return c.DescribeScanReportListWithContext(context.Background(), request)
}

// DescribeScanReportList
// 获取扫描报告列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
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
func (c *Client) DescribeScanReportListWithContext(ctx context.Context, request *DescribeScanReportListRequest) (response *DescribeScanReportListResponse, err error) {
    if request == nil {
        request = NewDescribeScanReportListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeScanReportList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeScanReportListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSubnetAssetsRequest() (request *DescribeSubnetAssetsRequest) {
    request = &DescribeSubnetAssetsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeSubnetAssets")
    
    
    return
}

func NewDescribeSubnetAssetsResponse() (response *DescribeSubnetAssetsResponse) {
    response = &DescribeSubnetAssetsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeSubnetAssets
// 获取子网列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
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
func (c *Client) DescribeSubnetAssets(request *DescribeSubnetAssetsRequest) (response *DescribeSubnetAssetsResponse, err error) {
    return c.DescribeSubnetAssetsWithContext(context.Background(), request)
}

// DescribeSubnetAssets
// 获取子网列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
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
func (c *Client) DescribeSubnetAssetsWithContext(ctx context.Context, request *DescribeSubnetAssetsRequest) (response *DescribeSubnetAssetsResponse, err error) {
    if request == nil {
        request = NewDescribeSubnetAssetsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSubnetAssets require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSubnetAssetsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVpcAssetsRequest() (request *DescribeVpcAssetsRequest) {
    request = &DescribeVpcAssetsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeVpcAssets")
    
    
    return
}

func NewDescribeVpcAssetsResponse() (response *DescribeVpcAssetsResponse) {
    response = &DescribeVpcAssetsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeVpcAssets
// 获取vpc列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
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
func (c *Client) DescribeVpcAssets(request *DescribeVpcAssetsRequest) (response *DescribeVpcAssetsResponse, err error) {
    return c.DescribeVpcAssetsWithContext(context.Background(), request)
}

// DescribeVpcAssets
// 获取vpc列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
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
func (c *Client) DescribeVpcAssetsWithContext(ctx context.Context, request *DescribeVpcAssetsRequest) (response *DescribeVpcAssetsResponse, err error) {
    if request == nil {
        request = NewDescribeVpcAssetsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVpcAssets require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVpcAssetsResponse()
    err = c.Send(request, response)
    return
}
