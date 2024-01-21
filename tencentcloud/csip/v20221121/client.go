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

func NewCreateDomainAndIpRequest() (request *CreateDomainAndIpRequest) {
    request = &CreateDomainAndIpRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "CreateDomainAndIp")
    
    
    return
}

func NewCreateDomainAndIpResponse() (response *CreateDomainAndIpResponse) {
    response = &CreateDomainAndIpResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateDomainAndIp
// 创建域名、ip相关信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CreateDomainAndIp(request *CreateDomainAndIpRequest) (response *CreateDomainAndIpResponse, err error) {
    return c.CreateDomainAndIpWithContext(context.Background(), request)
}

// CreateDomainAndIp
// 创建域名、ip相关信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CreateDomainAndIpWithContext(ctx context.Context, request *CreateDomainAndIpRequest) (response *CreateDomainAndIpResponse, err error) {
    if request == nil {
        request = NewCreateDomainAndIpRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateDomainAndIp require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateDomainAndIpResponse()
    err = c.Send(request, response)
    return
}

func NewCreateRiskCenterScanTaskRequest() (request *CreateRiskCenterScanTaskRequest) {
    request = &CreateRiskCenterScanTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "CreateRiskCenterScanTask")
    
    
    return
}

func NewCreateRiskCenterScanTaskResponse() (response *CreateRiskCenterScanTaskResponse) {
    response = &CreateRiskCenterScanTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateRiskCenterScanTask
// 创建风险中心扫描任务
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
func (c *Client) CreateRiskCenterScanTask(request *CreateRiskCenterScanTaskRequest) (response *CreateRiskCenterScanTaskResponse, err error) {
    return c.CreateRiskCenterScanTaskWithContext(context.Background(), request)
}

// CreateRiskCenterScanTask
// 创建风险中心扫描任务
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
func (c *Client) CreateRiskCenterScanTaskWithContext(ctx context.Context, request *CreateRiskCenterScanTaskRequest) (response *CreateRiskCenterScanTaskResponse, err error) {
    if request == nil {
        request = NewCreateRiskCenterScanTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateRiskCenterScanTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateRiskCenterScanTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteDomainAndIpRequest() (request *DeleteDomainAndIpRequest) {
    request = &DeleteDomainAndIpRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DeleteDomainAndIp")
    
    
    return
}

func NewDeleteDomainAndIpResponse() (response *DeleteDomainAndIpResponse) {
    response = &DeleteDomainAndIpResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteDomainAndIp
// 删除域名和ip请求
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
func (c *Client) DeleteDomainAndIp(request *DeleteDomainAndIpRequest) (response *DeleteDomainAndIpResponse, err error) {
    return c.DeleteDomainAndIpWithContext(context.Background(), request)
}

// DeleteDomainAndIp
// 删除域名和ip请求
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
func (c *Client) DeleteDomainAndIpWithContext(ctx context.Context, request *DeleteDomainAndIpRequest) (response *DeleteDomainAndIpResponse, err error) {
    if request == nil {
        request = NewDeleteDomainAndIpRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteDomainAndIp require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteDomainAndIpResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteRiskScanTaskRequest() (request *DeleteRiskScanTaskRequest) {
    request = &DeleteRiskScanTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DeleteRiskScanTask")
    
    
    return
}

func NewDeleteRiskScanTaskResponse() (response *DeleteRiskScanTaskResponse) {
    response = &DeleteRiskScanTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteRiskScanTask
// 删除风险中心扫描任务
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
func (c *Client) DeleteRiskScanTask(request *DeleteRiskScanTaskRequest) (response *DeleteRiskScanTaskResponse, err error) {
    return c.DeleteRiskScanTaskWithContext(context.Background(), request)
}

// DeleteRiskScanTask
// 删除风险中心扫描任务
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
func (c *Client) DeleteRiskScanTaskWithContext(ctx context.Context, request *DeleteRiskScanTaskRequest) (response *DeleteRiskScanTaskResponse, err error) {
    if request == nil {
        request = NewDeleteRiskScanTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteRiskScanTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteRiskScanTaskResponse()
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
func (c *Client) DescribeCVMAssetInfo(request *DescribeCVMAssetInfoRequest) (response *DescribeCVMAssetInfoResponse, err error) {
    return c.DescribeCVMAssetInfoWithContext(context.Background(), request)
}

// DescribeCVMAssetInfo
// cvm详情
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
func (c *Client) DescribeCVMAssets(request *DescribeCVMAssetsRequest) (response *DescribeCVMAssetsResponse, err error) {
    return c.DescribeCVMAssetsWithContext(context.Background(), request)
}

// DescribeCVMAssets
// cvm列表
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

func NewDescribeClusterPodAssetsRequest() (request *DescribeClusterPodAssetsRequest) {
    request = &DescribeClusterPodAssetsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeClusterPodAssets")
    
    
    return
}

func NewDescribeClusterPodAssetsResponse() (response *DescribeClusterPodAssetsResponse) {
    response = &DescribeClusterPodAssetsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeClusterPodAssets
// 集群pod列表
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
func (c *Client) DescribeClusterPodAssets(request *DescribeClusterPodAssetsRequest) (response *DescribeClusterPodAssetsResponse, err error) {
    return c.DescribeClusterPodAssetsWithContext(context.Background(), request)
}

// DescribeClusterPodAssets
// 集群pod列表
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
func (c *Client) DescribeClusterPodAssetsWithContext(ctx context.Context, request *DescribeClusterPodAssetsRequest) (response *DescribeClusterPodAssetsResponse, err error) {
    if request == nil {
        request = NewDescribeClusterPodAssetsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeClusterPodAssets require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeClusterPodAssetsResponse()
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
func (c *Client) DescribeDbAssetInfo(request *DescribeDbAssetInfoRequest) (response *DescribeDbAssetInfoResponse, err error) {
    return c.DescribeDbAssetInfoWithContext(context.Background(), request)
}

// DescribeDbAssetInfo
// db资产详情
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
// 数据库资产列表
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
func (c *Client) DescribeDbAssets(request *DescribeDbAssetsRequest) (response *DescribeDbAssetsResponse, err error) {
    return c.DescribeDbAssetsWithContext(context.Background(), request)
}

// DescribeDbAssets
// 数据库资产列表
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

func NewDescribeDomainAssetsRequest() (request *DescribeDomainAssetsRequest) {
    request = &DescribeDomainAssetsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeDomainAssets")
    
    
    return
}

func NewDescribeDomainAssetsResponse() (response *DescribeDomainAssetsResponse) {
    response = &DescribeDomainAssetsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDomainAssets
// 域名列表
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
func (c *Client) DescribeDomainAssets(request *DescribeDomainAssetsRequest) (response *DescribeDomainAssetsResponse, err error) {
    return c.DescribeDomainAssetsWithContext(context.Background(), request)
}

// DescribeDomainAssets
// 域名列表
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
func (c *Client) DescribeDomainAssetsWithContext(ctx context.Context, request *DescribeDomainAssetsRequest) (response *DescribeDomainAssetsResponse, err error) {
    if request == nil {
        request = NewDescribeDomainAssetsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDomainAssets require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDomainAssetsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeGatewayAssetsRequest() (request *DescribeGatewayAssetsRequest) {
    request = &DescribeGatewayAssetsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeGatewayAssets")
    
    
    return
}

func NewDescribeGatewayAssetsResponse() (response *DescribeGatewayAssetsResponse) {
    response = &DescribeGatewayAssetsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeGatewayAssets
// 获取网关列表
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
func (c *Client) DescribeGatewayAssets(request *DescribeGatewayAssetsRequest) (response *DescribeGatewayAssetsResponse, err error) {
    return c.DescribeGatewayAssetsWithContext(context.Background(), request)
}

// DescribeGatewayAssets
// 获取网关列表
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
func (c *Client) DescribeGatewayAssetsWithContext(ctx context.Context, request *DescribeGatewayAssetsRequest) (response *DescribeGatewayAssetsResponse, err error) {
    if request == nil {
        request = NewDescribeGatewayAssetsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeGatewayAssets require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeGatewayAssetsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeListenerListRequest() (request *DescribeListenerListRequest) {
    request = &DescribeListenerListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeListenerList")
    
    
    return
}

func NewDescribeListenerListResponse() (response *DescribeListenerListResponse) {
    response = &DescribeListenerListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeListenerList
// 查询clb监听器列表
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
func (c *Client) DescribeListenerList(request *DescribeListenerListRequest) (response *DescribeListenerListResponse, err error) {
    return c.DescribeListenerListWithContext(context.Background(), request)
}

// DescribeListenerList
// 查询clb监听器列表
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
func (c *Client) DescribeListenerListWithContext(ctx context.Context, request *DescribeListenerListRequest) (response *DescribeListenerListResponse, err error) {
    if request == nil {
        request = NewDescribeListenerListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeListenerList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeListenerListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeNICAssetsRequest() (request *DescribeNICAssetsRequest) {
    request = &DescribeNICAssetsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeNICAssets")
    
    
    return
}

func NewDescribeNICAssetsResponse() (response *DescribeNICAssetsResponse) {
    response = &DescribeNICAssetsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeNICAssets
// 获取网卡列表
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
func (c *Client) DescribeNICAssets(request *DescribeNICAssetsRequest) (response *DescribeNICAssetsResponse, err error) {
    return c.DescribeNICAssetsWithContext(context.Background(), request)
}

// DescribeNICAssets
// 获取网卡列表
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
func (c *Client) DescribeNICAssetsWithContext(ctx context.Context, request *DescribeNICAssetsRequest) (response *DescribeNICAssetsResponse, err error) {
    if request == nil {
        request = NewDescribeNICAssetsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeNICAssets require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeNICAssetsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeOrganizationUserInfoRequest() (request *DescribeOrganizationUserInfoRequest) {
    request = &DescribeOrganizationUserInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeOrganizationUserInfo")
    
    
    return
}

func NewDescribeOrganizationUserInfoResponse() (response *DescribeOrganizationUserInfoResponse) {
    response = &DescribeOrganizationUserInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeOrganizationUserInfo
// 查询集团账号用户列表
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
func (c *Client) DescribeOrganizationUserInfo(request *DescribeOrganizationUserInfoRequest) (response *DescribeOrganizationUserInfoResponse, err error) {
    return c.DescribeOrganizationUserInfoWithContext(context.Background(), request)
}

// DescribeOrganizationUserInfo
// 查询集团账号用户列表
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
func (c *Client) DescribeOrganizationUserInfoWithContext(ctx context.Context, request *DescribeOrganizationUserInfoRequest) (response *DescribeOrganizationUserInfoResponse, err error) {
    if request == nil {
        request = NewDescribeOrganizationUserInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeOrganizationUserInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeOrganizationUserInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePublicIpAssetsRequest() (request *DescribePublicIpAssetsRequest) {
    request = &DescribePublicIpAssetsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribePublicIpAssets")
    
    
    return
}

func NewDescribePublicIpAssetsResponse() (response *DescribePublicIpAssetsResponse) {
    response = &DescribePublicIpAssetsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribePublicIpAssets
// ip公网列表
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
func (c *Client) DescribePublicIpAssets(request *DescribePublicIpAssetsRequest) (response *DescribePublicIpAssetsResponse, err error) {
    return c.DescribePublicIpAssetsWithContext(context.Background(), request)
}

// DescribePublicIpAssets
// ip公网列表
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
func (c *Client) DescribePublicIpAssetsWithContext(ctx context.Context, request *DescribePublicIpAssetsRequest) (response *DescribePublicIpAssetsResponse, err error) {
    if request == nil {
        request = NewDescribePublicIpAssetsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePublicIpAssets require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePublicIpAssetsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRiskCenterAssetViewCFGRiskListRequest() (request *DescribeRiskCenterAssetViewCFGRiskListRequest) {
    request = &DescribeRiskCenterAssetViewCFGRiskListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeRiskCenterAssetViewCFGRiskList")
    
    
    return
}

func NewDescribeRiskCenterAssetViewCFGRiskListResponse() (response *DescribeRiskCenterAssetViewCFGRiskListResponse) {
    response = &DescribeRiskCenterAssetViewCFGRiskListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRiskCenterAssetViewCFGRiskList
// 获取资产视角的配置风险列表
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
func (c *Client) DescribeRiskCenterAssetViewCFGRiskList(request *DescribeRiskCenterAssetViewCFGRiskListRequest) (response *DescribeRiskCenterAssetViewCFGRiskListResponse, err error) {
    return c.DescribeRiskCenterAssetViewCFGRiskListWithContext(context.Background(), request)
}

// DescribeRiskCenterAssetViewCFGRiskList
// 获取资产视角的配置风险列表
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
func (c *Client) DescribeRiskCenterAssetViewCFGRiskListWithContext(ctx context.Context, request *DescribeRiskCenterAssetViewCFGRiskListRequest) (response *DescribeRiskCenterAssetViewCFGRiskListResponse, err error) {
    if request == nil {
        request = NewDescribeRiskCenterAssetViewCFGRiskListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRiskCenterAssetViewCFGRiskList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRiskCenterAssetViewCFGRiskListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRiskCenterAssetViewPortRiskListRequest() (request *DescribeRiskCenterAssetViewPortRiskListRequest) {
    request = &DescribeRiskCenterAssetViewPortRiskListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeRiskCenterAssetViewPortRiskList")
    
    
    return
}

func NewDescribeRiskCenterAssetViewPortRiskListResponse() (response *DescribeRiskCenterAssetViewPortRiskListResponse) {
    response = &DescribeRiskCenterAssetViewPortRiskListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRiskCenterAssetViewPortRiskList
// 获取资产视角的端口风险列表
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
func (c *Client) DescribeRiskCenterAssetViewPortRiskList(request *DescribeRiskCenterAssetViewPortRiskListRequest) (response *DescribeRiskCenterAssetViewPortRiskListResponse, err error) {
    return c.DescribeRiskCenterAssetViewPortRiskListWithContext(context.Background(), request)
}

// DescribeRiskCenterAssetViewPortRiskList
// 获取资产视角的端口风险列表
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
func (c *Client) DescribeRiskCenterAssetViewPortRiskListWithContext(ctx context.Context, request *DescribeRiskCenterAssetViewPortRiskListRequest) (response *DescribeRiskCenterAssetViewPortRiskListResponse, err error) {
    if request == nil {
        request = NewDescribeRiskCenterAssetViewPortRiskListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRiskCenterAssetViewPortRiskList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRiskCenterAssetViewPortRiskListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRiskCenterAssetViewVULRiskListRequest() (request *DescribeRiskCenterAssetViewVULRiskListRequest) {
    request = &DescribeRiskCenterAssetViewVULRiskListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeRiskCenterAssetViewVULRiskList")
    
    
    return
}

func NewDescribeRiskCenterAssetViewVULRiskListResponse() (response *DescribeRiskCenterAssetViewVULRiskListResponse) {
    response = &DescribeRiskCenterAssetViewVULRiskListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRiskCenterAssetViewVULRiskList
// 获取资产视角的漏洞风险列表
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
func (c *Client) DescribeRiskCenterAssetViewVULRiskList(request *DescribeRiskCenterAssetViewVULRiskListRequest) (response *DescribeRiskCenterAssetViewVULRiskListResponse, err error) {
    return c.DescribeRiskCenterAssetViewVULRiskListWithContext(context.Background(), request)
}

// DescribeRiskCenterAssetViewVULRiskList
// 获取资产视角的漏洞风险列表
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
func (c *Client) DescribeRiskCenterAssetViewVULRiskListWithContext(ctx context.Context, request *DescribeRiskCenterAssetViewVULRiskListRequest) (response *DescribeRiskCenterAssetViewVULRiskListResponse, err error) {
    if request == nil {
        request = NewDescribeRiskCenterAssetViewVULRiskListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRiskCenterAssetViewVULRiskList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRiskCenterAssetViewVULRiskListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRiskCenterAssetViewWeakPasswordRiskListRequest() (request *DescribeRiskCenterAssetViewWeakPasswordRiskListRequest) {
    request = &DescribeRiskCenterAssetViewWeakPasswordRiskListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeRiskCenterAssetViewWeakPasswordRiskList")
    
    
    return
}

func NewDescribeRiskCenterAssetViewWeakPasswordRiskListResponse() (response *DescribeRiskCenterAssetViewWeakPasswordRiskListResponse) {
    response = &DescribeRiskCenterAssetViewWeakPasswordRiskListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRiskCenterAssetViewWeakPasswordRiskList
// 获取资产视角的弱口令风险列表
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
func (c *Client) DescribeRiskCenterAssetViewWeakPasswordRiskList(request *DescribeRiskCenterAssetViewWeakPasswordRiskListRequest) (response *DescribeRiskCenterAssetViewWeakPasswordRiskListResponse, err error) {
    return c.DescribeRiskCenterAssetViewWeakPasswordRiskListWithContext(context.Background(), request)
}

// DescribeRiskCenterAssetViewWeakPasswordRiskList
// 获取资产视角的弱口令风险列表
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
func (c *Client) DescribeRiskCenterAssetViewWeakPasswordRiskListWithContext(ctx context.Context, request *DescribeRiskCenterAssetViewWeakPasswordRiskListRequest) (response *DescribeRiskCenterAssetViewWeakPasswordRiskListResponse, err error) {
    if request == nil {
        request = NewDescribeRiskCenterAssetViewWeakPasswordRiskListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRiskCenterAssetViewWeakPasswordRiskList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRiskCenterAssetViewWeakPasswordRiskListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRiskCenterPortViewPortRiskListRequest() (request *DescribeRiskCenterPortViewPortRiskListRequest) {
    request = &DescribeRiskCenterPortViewPortRiskListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeRiskCenterPortViewPortRiskList")
    
    
    return
}

func NewDescribeRiskCenterPortViewPortRiskListResponse() (response *DescribeRiskCenterPortViewPortRiskListResponse) {
    response = &DescribeRiskCenterPortViewPortRiskListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRiskCenterPortViewPortRiskList
// 获取端口视角的端口风险列表
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
func (c *Client) DescribeRiskCenterPortViewPortRiskList(request *DescribeRiskCenterPortViewPortRiskListRequest) (response *DescribeRiskCenterPortViewPortRiskListResponse, err error) {
    return c.DescribeRiskCenterPortViewPortRiskListWithContext(context.Background(), request)
}

// DescribeRiskCenterPortViewPortRiskList
// 获取端口视角的端口风险列表
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
func (c *Client) DescribeRiskCenterPortViewPortRiskListWithContext(ctx context.Context, request *DescribeRiskCenterPortViewPortRiskListRequest) (response *DescribeRiskCenterPortViewPortRiskListResponse, err error) {
    if request == nil {
        request = NewDescribeRiskCenterPortViewPortRiskListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRiskCenterPortViewPortRiskList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRiskCenterPortViewPortRiskListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRiskCenterServerRiskListRequest() (request *DescribeRiskCenterServerRiskListRequest) {
    request = &DescribeRiskCenterServerRiskListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeRiskCenterServerRiskList")
    
    
    return
}

func NewDescribeRiskCenterServerRiskListResponse() (response *DescribeRiskCenterServerRiskListResponse) {
    response = &DescribeRiskCenterServerRiskListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRiskCenterServerRiskList
// 获取风险服务列表
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
func (c *Client) DescribeRiskCenterServerRiskList(request *DescribeRiskCenterServerRiskListRequest) (response *DescribeRiskCenterServerRiskListResponse, err error) {
    return c.DescribeRiskCenterServerRiskListWithContext(context.Background(), request)
}

// DescribeRiskCenterServerRiskList
// 获取风险服务列表
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
func (c *Client) DescribeRiskCenterServerRiskListWithContext(ctx context.Context, request *DescribeRiskCenterServerRiskListRequest) (response *DescribeRiskCenterServerRiskListResponse, err error) {
    if request == nil {
        request = NewDescribeRiskCenterServerRiskListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRiskCenterServerRiskList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRiskCenterServerRiskListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRiskCenterVULViewVULRiskListRequest() (request *DescribeRiskCenterVULViewVULRiskListRequest) {
    request = &DescribeRiskCenterVULViewVULRiskListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeRiskCenterVULViewVULRiskList")
    
    
    return
}

func NewDescribeRiskCenterVULViewVULRiskListResponse() (response *DescribeRiskCenterVULViewVULRiskListResponse) {
    response = &DescribeRiskCenterVULViewVULRiskListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRiskCenterVULViewVULRiskList
// 获取漏洞视角的漏洞风险列表
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
func (c *Client) DescribeRiskCenterVULViewVULRiskList(request *DescribeRiskCenterVULViewVULRiskListRequest) (response *DescribeRiskCenterVULViewVULRiskListResponse, err error) {
    return c.DescribeRiskCenterVULViewVULRiskListWithContext(context.Background(), request)
}

// DescribeRiskCenterVULViewVULRiskList
// 获取漏洞视角的漏洞风险列表
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
func (c *Client) DescribeRiskCenterVULViewVULRiskListWithContext(ctx context.Context, request *DescribeRiskCenterVULViewVULRiskListRequest) (response *DescribeRiskCenterVULViewVULRiskListResponse, err error) {
    if request == nil {
        request = NewDescribeRiskCenterVULViewVULRiskListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRiskCenterVULViewVULRiskList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRiskCenterVULViewVULRiskListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRiskCenterWebsiteRiskListRequest() (request *DescribeRiskCenterWebsiteRiskListRequest) {
    request = &DescribeRiskCenterWebsiteRiskListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeRiskCenterWebsiteRiskList")
    
    
    return
}

func NewDescribeRiskCenterWebsiteRiskListResponse() (response *DescribeRiskCenterWebsiteRiskListResponse) {
    response = &DescribeRiskCenterWebsiteRiskListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRiskCenterWebsiteRiskList
// 获取内容风险列表
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
func (c *Client) DescribeRiskCenterWebsiteRiskList(request *DescribeRiskCenterWebsiteRiskListRequest) (response *DescribeRiskCenterWebsiteRiskListResponse, err error) {
    return c.DescribeRiskCenterWebsiteRiskListWithContext(context.Background(), request)
}

// DescribeRiskCenterWebsiteRiskList
// 获取内容风险列表
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
func (c *Client) DescribeRiskCenterWebsiteRiskListWithContext(ctx context.Context, request *DescribeRiskCenterWebsiteRiskListRequest) (response *DescribeRiskCenterWebsiteRiskListResponse, err error) {
    if request == nil {
        request = NewDescribeRiskCenterWebsiteRiskListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRiskCenterWebsiteRiskList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRiskCenterWebsiteRiskListResponse()
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

func NewDescribeScanTaskListRequest() (request *DescribeScanTaskListRequest) {
    request = &DescribeScanTaskListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeScanTaskList")
    
    
    return
}

func NewDescribeScanTaskListResponse() (response *DescribeScanTaskListResponse) {
    response = &DescribeScanTaskListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeScanTaskList
// 获取扫描任务列表
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
func (c *Client) DescribeScanTaskList(request *DescribeScanTaskListRequest) (response *DescribeScanTaskListResponse, err error) {
    return c.DescribeScanTaskListWithContext(context.Background(), request)
}

// DescribeScanTaskList
// 获取扫描任务列表
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
func (c *Client) DescribeScanTaskListWithContext(ctx context.Context, request *DescribeScanTaskListRequest) (response *DescribeScanTaskListResponse, err error) {
    if request == nil {
        request = NewDescribeScanTaskListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeScanTaskList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeScanTaskListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSearchBugInfoRequest() (request *DescribeSearchBugInfoRequest) {
    request = &DescribeSearchBugInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeSearchBugInfo")
    
    
    return
}

func NewDescribeSearchBugInfoResponse() (response *DescribeSearchBugInfoResponse) {
    response = &DescribeSearchBugInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSearchBugInfo
// 立体防护中心查询漏洞信息
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
func (c *Client) DescribeSearchBugInfo(request *DescribeSearchBugInfoRequest) (response *DescribeSearchBugInfoResponse, err error) {
    return c.DescribeSearchBugInfoWithContext(context.Background(), request)
}

// DescribeSearchBugInfo
// 立体防护中心查询漏洞信息
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
func (c *Client) DescribeSearchBugInfoWithContext(ctx context.Context, request *DescribeSearchBugInfoRequest) (response *DescribeSearchBugInfoResponse, err error) {
    if request == nil {
        request = NewDescribeSearchBugInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSearchBugInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSearchBugInfoResponse()
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

func NewDescribeTaskLogListRequest() (request *DescribeTaskLogListRequest) {
    request = &DescribeTaskLogListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeTaskLogList")
    
    
    return
}

func NewDescribeTaskLogListResponse() (response *DescribeTaskLogListResponse) {
    response = &DescribeTaskLogListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTaskLogList
// 获取任务扫描报告列表
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
func (c *Client) DescribeTaskLogList(request *DescribeTaskLogListRequest) (response *DescribeTaskLogListResponse, err error) {
    return c.DescribeTaskLogListWithContext(context.Background(), request)
}

// DescribeTaskLogList
// 获取任务扫描报告列表
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
func (c *Client) DescribeTaskLogListWithContext(ctx context.Context, request *DescribeTaskLogListRequest) (response *DescribeTaskLogListResponse, err error) {
    if request == nil {
        request = NewDescribeTaskLogListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTaskLogList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTaskLogListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTaskLogURLRequest() (request *DescribeTaskLogURLRequest) {
    request = &DescribeTaskLogURLRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeTaskLogURL")
    
    
    return
}

func NewDescribeTaskLogURLResponse() (response *DescribeTaskLogURLResponse) {
    response = &DescribeTaskLogURLResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTaskLogURL
// 获取报告下载的临时链接
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
func (c *Client) DescribeTaskLogURL(request *DescribeTaskLogURLRequest) (response *DescribeTaskLogURLResponse, err error) {
    return c.DescribeTaskLogURLWithContext(context.Background(), request)
}

// DescribeTaskLogURL
// 获取报告下载的临时链接
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
func (c *Client) DescribeTaskLogURLWithContext(ctx context.Context, request *DescribeTaskLogURLRequest) (response *DescribeTaskLogURLResponse, err error) {
    if request == nil {
        request = NewDescribeTaskLogURLRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTaskLogURL require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTaskLogURLResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVULRiskAdvanceCFGListRequest() (request *DescribeVULRiskAdvanceCFGListRequest) {
    request = &DescribeVULRiskAdvanceCFGListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeVULRiskAdvanceCFGList")
    
    
    return
}

func NewDescribeVULRiskAdvanceCFGListResponse() (response *DescribeVULRiskAdvanceCFGListResponse) {
    response = &DescribeVULRiskAdvanceCFGListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeVULRiskAdvanceCFGList
// 查询漏洞风险高级配置
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
func (c *Client) DescribeVULRiskAdvanceCFGList(request *DescribeVULRiskAdvanceCFGListRequest) (response *DescribeVULRiskAdvanceCFGListResponse, err error) {
    return c.DescribeVULRiskAdvanceCFGListWithContext(context.Background(), request)
}

// DescribeVULRiskAdvanceCFGList
// 查询漏洞风险高级配置
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
func (c *Client) DescribeVULRiskAdvanceCFGListWithContext(ctx context.Context, request *DescribeVULRiskAdvanceCFGListRequest) (response *DescribeVULRiskAdvanceCFGListResponse, err error) {
    if request == nil {
        request = NewDescribeVULRiskAdvanceCFGListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVULRiskAdvanceCFGList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVULRiskAdvanceCFGListResponse()
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

func NewModifyOrganizationAccountStatusRequest() (request *ModifyOrganizationAccountStatusRequest) {
    request = &ModifyOrganizationAccountStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "ModifyOrganizationAccountStatus")
    
    
    return
}

func NewModifyOrganizationAccountStatusResponse() (response *ModifyOrganizationAccountStatusResponse) {
    response = &ModifyOrganizationAccountStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyOrganizationAccountStatus
// 修改集团账号状态
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
func (c *Client) ModifyOrganizationAccountStatus(request *ModifyOrganizationAccountStatusRequest) (response *ModifyOrganizationAccountStatusResponse, err error) {
    return c.ModifyOrganizationAccountStatusWithContext(context.Background(), request)
}

// ModifyOrganizationAccountStatus
// 修改集团账号状态
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
func (c *Client) ModifyOrganizationAccountStatusWithContext(ctx context.Context, request *ModifyOrganizationAccountStatusRequest) (response *ModifyOrganizationAccountStatusResponse, err error) {
    if request == nil {
        request = NewModifyOrganizationAccountStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyOrganizationAccountStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyOrganizationAccountStatusResponse()
    err = c.Send(request, response)
    return
}

func NewModifyRiskCenterRiskStatusRequest() (request *ModifyRiskCenterRiskStatusRequest) {
    request = &ModifyRiskCenterRiskStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "ModifyRiskCenterRiskStatus")
    
    
    return
}

func NewModifyRiskCenterRiskStatusResponse() (response *ModifyRiskCenterRiskStatusResponse) {
    response = &ModifyRiskCenterRiskStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyRiskCenterRiskStatus
// 修改风险中心风险状态
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
func (c *Client) ModifyRiskCenterRiskStatus(request *ModifyRiskCenterRiskStatusRequest) (response *ModifyRiskCenterRiskStatusResponse, err error) {
    return c.ModifyRiskCenterRiskStatusWithContext(context.Background(), request)
}

// ModifyRiskCenterRiskStatus
// 修改风险中心风险状态
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
func (c *Client) ModifyRiskCenterRiskStatusWithContext(ctx context.Context, request *ModifyRiskCenterRiskStatusRequest) (response *ModifyRiskCenterRiskStatusResponse, err error) {
    if request == nil {
        request = NewModifyRiskCenterRiskStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyRiskCenterRiskStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyRiskCenterRiskStatusResponse()
    err = c.Send(request, response)
    return
}

func NewModifyRiskCenterScanTaskRequest() (request *ModifyRiskCenterScanTaskRequest) {
    request = &ModifyRiskCenterScanTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "ModifyRiskCenterScanTask")
    
    
    return
}

func NewModifyRiskCenterScanTaskResponse() (response *ModifyRiskCenterScanTaskResponse) {
    response = &ModifyRiskCenterScanTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyRiskCenterScanTask
// 修改风险中心扫描任务
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
func (c *Client) ModifyRiskCenterScanTask(request *ModifyRiskCenterScanTaskRequest) (response *ModifyRiskCenterScanTaskResponse, err error) {
    return c.ModifyRiskCenterScanTaskWithContext(context.Background(), request)
}

// ModifyRiskCenterScanTask
// 修改风险中心扫描任务
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
func (c *Client) ModifyRiskCenterScanTaskWithContext(ctx context.Context, request *ModifyRiskCenterScanTaskRequest) (response *ModifyRiskCenterScanTaskResponse, err error) {
    if request == nil {
        request = NewModifyRiskCenterScanTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyRiskCenterScanTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyRiskCenterScanTaskResponse()
    err = c.Send(request, response)
    return
}

func NewStopRiskCenterTaskRequest() (request *StopRiskCenterTaskRequest) {
    request = &StopRiskCenterTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "StopRiskCenterTask")
    
    
    return
}

func NewStopRiskCenterTaskResponse() (response *StopRiskCenterTaskResponse) {
    response = &StopRiskCenterTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// StopRiskCenterTask
// 停止扫风险中心扫描任务
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
func (c *Client) StopRiskCenterTask(request *StopRiskCenterTaskRequest) (response *StopRiskCenterTaskResponse, err error) {
    return c.StopRiskCenterTaskWithContext(context.Background(), request)
}

// StopRiskCenterTask
// 停止扫风险中心扫描任务
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
func (c *Client) StopRiskCenterTaskWithContext(ctx context.Context, request *StopRiskCenterTaskRequest) (response *StopRiskCenterTaskResponse, err error) {
    if request == nil {
        request = NewStopRiskCenterTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("StopRiskCenterTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewStopRiskCenterTaskResponse()
    err = c.Send(request, response)
    return
}
