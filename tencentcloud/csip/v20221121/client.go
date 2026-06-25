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


func NewAddDspmAssetManagerRequest() (request *AddDspmAssetManagerRequest) {
    request = &AddDspmAssetManagerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "AddDspmAssetManager")
    
    
    return
}

func NewAddDspmAssetManagerResponse() (response *AddDspmAssetManagerResponse) {
    response = &AddDspmAssetManagerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AddDspmAssetManager
// 添加资产管理员
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) AddDspmAssetManager(request *AddDspmAssetManagerRequest) (response *AddDspmAssetManagerResponse, err error) {
    return c.AddDspmAssetManagerWithContext(context.Background(), request)
}

// AddDspmAssetManager
// 添加资产管理员
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) AddDspmAssetManagerWithContext(ctx context.Context, request *AddDspmAssetManagerRequest) (response *AddDspmAssetManagerResponse, err error) {
    if request == nil {
        request = NewAddDspmAssetManagerRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "AddDspmAssetManager")
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddDspmAssetManager require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddDspmAssetManagerResponse()
    err = c.Send(request, response)
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
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) AddNewBindRoleUser(request *AddNewBindRoleUserRequest) (response *AddNewBindRoleUserResponse, err error) {
    return c.AddNewBindRoleUserWithContext(context.Background(), request)
}

// AddNewBindRoleUser
// csip角色授权绑定接口
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) AddNewBindRoleUserWithContext(ctx context.Context, request *AddNewBindRoleUserRequest) (response *AddNewBindRoleUserResponse, err error) {
    if request == nil {
        request = NewAddNewBindRoleUserRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "AddNewBindRoleUser")
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddNewBindRoleUser require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddNewBindRoleUserResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAccessKeyCheckTaskRequest() (request *CreateAccessKeyCheckTaskRequest) {
    request = &CreateAccessKeyCheckTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "CreateAccessKeyCheckTask")
    
    
    return
}

func NewCreateAccessKeyCheckTaskResponse() (response *CreateAccessKeyCheckTaskResponse) {
    response = &CreateAccessKeyCheckTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateAccessKeyCheckTask
// 检测AK 异步任务
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateAccessKeyCheckTask(request *CreateAccessKeyCheckTaskRequest) (response *CreateAccessKeyCheckTaskResponse, err error) {
    return c.CreateAccessKeyCheckTaskWithContext(context.Background(), request)
}

// CreateAccessKeyCheckTask
// 检测AK 异步任务
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateAccessKeyCheckTaskWithContext(ctx context.Context, request *CreateAccessKeyCheckTaskRequest) (response *CreateAccessKeyCheckTaskResponse, err error) {
    if request == nil {
        request = NewCreateAccessKeyCheckTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "CreateAccessKeyCheckTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAccessKeyCheckTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAccessKeyCheckTaskResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAccessKeySyncTaskRequest() (request *CreateAccessKeySyncTaskRequest) {
    request = &CreateAccessKeySyncTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "CreateAccessKeySyncTask")
    
    
    return
}

func NewCreateAccessKeySyncTaskResponse() (response *CreateAccessKeySyncTaskResponse) {
    response = &CreateAccessKeySyncTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateAccessKeySyncTask
// 发起AK资产同步任务
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateAccessKeySyncTask(request *CreateAccessKeySyncTaskRequest) (response *CreateAccessKeySyncTaskResponse, err error) {
    return c.CreateAccessKeySyncTaskWithContext(context.Background(), request)
}

// CreateAccessKeySyncTask
// 发起AK资产同步任务
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateAccessKeySyncTaskWithContext(ctx context.Context, request *CreateAccessKeySyncTaskRequest) (response *CreateAccessKeySyncTaskResponse, err error) {
    if request == nil {
        request = NewCreateAccessKeySyncTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "CreateAccessKeySyncTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAccessKeySyncTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAccessKeySyncTaskResponse()
    err = c.Send(request, response)
    return
}

func NewCreateCosAssetSyncTaskRequest() (request *CreateCosAssetSyncTaskRequest) {
    request = &CreateCosAssetSyncTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "CreateCosAssetSyncTask")
    
    
    return
}

func NewCreateCosAssetSyncTaskResponse() (response *CreateCosAssetSyncTaskResponse) {
    response = &CreateCosAssetSyncTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateCosAssetSyncTask
// 创建资产同步任务
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateCosAssetSyncTask(request *CreateCosAssetSyncTaskRequest) (response *CreateCosAssetSyncTaskResponse, err error) {
    return c.CreateCosAssetSyncTaskWithContext(context.Background(), request)
}

// CreateCosAssetSyncTask
// 创建资产同步任务
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateCosAssetSyncTaskWithContext(ctx context.Context, request *CreateCosAssetSyncTaskRequest) (response *CreateCosAssetSyncTaskResponse, err error) {
    if request == nil {
        request = NewCreateCosAssetSyncTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "CreateCosAssetSyncTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCosAssetSyncTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateCosAssetSyncTaskResponse()
    err = c.Send(request, response)
    return
}

func NewCreateCosObjectScanTaskRequest() (request *CreateCosObjectScanTaskRequest) {
    request = &CreateCosObjectScanTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "CreateCosObjectScanTask")
    
    
    return
}

func NewCreateCosObjectScanTaskResponse() (response *CreateCosObjectScanTaskResponse) {
    response = &CreateCosObjectScanTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateCosObjectScanTask
// 创建cos病毒扫描、敏感数据识别任务
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateCosObjectScanTask(request *CreateCosObjectScanTaskRequest) (response *CreateCosObjectScanTaskResponse, err error) {
    return c.CreateCosObjectScanTaskWithContext(context.Background(), request)
}

// CreateCosObjectScanTask
// 创建cos病毒扫描、敏感数据识别任务
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateCosObjectScanTaskWithContext(ctx context.Context, request *CreateCosObjectScanTaskRequest) (response *CreateCosObjectScanTaskResponse, err error) {
    if request == nil {
        request = NewCreateCosObjectScanTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "CreateCosObjectScanTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCosObjectScanTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateCosObjectScanTaskResponse()
    err = c.Send(request, response)
    return
}

func NewCreateCosPolicyRequest() (request *CreateCosPolicyRequest) {
    request = &CreateCosPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "CreateCosPolicy")
    
    
    return
}

func NewCreateCosPolicyResponse() (response *CreateCosPolicyResponse) {
    response = &CreateCosPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateCosPolicy
// 添加cos告警策略
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateCosPolicy(request *CreateCosPolicyRequest) (response *CreateCosPolicyResponse, err error) {
    return c.CreateCosPolicyWithContext(context.Background(), request)
}

// CreateCosPolicy
// 添加cos告警策略
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateCosPolicyWithContext(ctx context.Context, request *CreateCosPolicyRequest) (response *CreateCosPolicyResponse, err error) {
    if request == nil {
        request = NewCreateCosPolicyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "CreateCosPolicy")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCosPolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateCosPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewCreateCosRiskScanTaskRequest() (request *CreateCosRiskScanTaskRequest) {
    request = &CreateCosRiskScanTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "CreateCosRiskScanTask")
    
    
    return
}

func NewCreateCosRiskScanTaskResponse() (response *CreateCosRiskScanTaskResponse) {
    response = &CreateCosRiskScanTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateCosRiskScanTask
// 创建风险监测任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) CreateCosRiskScanTask(request *CreateCosRiskScanTaskRequest) (response *CreateCosRiskScanTaskResponse, err error) {
    return c.CreateCosRiskScanTaskWithContext(context.Background(), request)
}

// CreateCosRiskScanTask
// 创建风险监测任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) CreateCosRiskScanTaskWithContext(ctx context.Context, request *CreateCosRiskScanTaskRequest) (response *CreateCosRiskScanTaskResponse, err error) {
    if request == nil {
        request = NewCreateCosRiskScanTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "CreateCosRiskScanTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCosRiskScanTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateCosRiskScanTaskResponse()
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
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "CreateDomainAndIp")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateDomainAndIp require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateDomainAndIpResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDspmAccessExportJobRequest() (request *CreateDspmAccessExportJobRequest) {
    request = &CreateDspmAccessExportJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "CreateDspmAccessExportJob")
    
    
    return
}

func NewCreateDspmAccessExportJobResponse() (response *CreateDspmAccessExportJobResponse) {
    response = &CreateDspmAccessExportJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateDspmAccessExportJob
// 创建Dspm访问记录导出任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CreateDspmAccessExportJob(request *CreateDspmAccessExportJobRequest) (response *CreateDspmAccessExportJobResponse, err error) {
    return c.CreateDspmAccessExportJobWithContext(context.Background(), request)
}

// CreateDspmAccessExportJob
// 创建Dspm访问记录导出任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CreateDspmAccessExportJobWithContext(ctx context.Context, request *CreateDspmAccessExportJobRequest) (response *CreateDspmAccessExportJobResponse, err error) {
    if request == nil {
        request = NewCreateDspmAccessExportJobRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "CreateDspmAccessExportJob")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateDspmAccessExportJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateDspmAccessExportJobResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDspmApplyOrderRequest() (request *CreateDspmApplyOrderRequest) {
    request = &CreateDspmApplyOrderRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "CreateDspmApplyOrder")
    
    
    return
}

func NewCreateDspmApplyOrderResponse() (response *CreateDspmApplyOrderResponse) {
    response = &CreateDspmApplyOrderResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateDspmApplyOrder
// 创建Dspm申请单
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateDspmApplyOrder(request *CreateDspmApplyOrderRequest) (response *CreateDspmApplyOrderResponse, err error) {
    return c.CreateDspmApplyOrderWithContext(context.Background(), request)
}

// CreateDspmApplyOrder
// 创建Dspm申请单
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateDspmApplyOrderWithContext(ctx context.Context, request *CreateDspmApplyOrderRequest) (response *CreateDspmApplyOrderResponse, err error) {
    if request == nil {
        request = NewCreateDspmApplyOrderRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "CreateDspmApplyOrder")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateDspmApplyOrder require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateDspmApplyOrderResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDspmApproveHistoryExportJobRequest() (request *CreateDspmApproveHistoryExportJobRequest) {
    request = &CreateDspmApproveHistoryExportJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "CreateDspmApproveHistoryExportJob")
    
    
    return
}

func NewCreateDspmApproveHistoryExportJobResponse() (response *CreateDspmApproveHistoryExportJobResponse) {
    response = &CreateDspmApproveHistoryExportJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateDspmApproveHistoryExportJob
// 创建Dspm审批历史导出任务
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateDspmApproveHistoryExportJob(request *CreateDspmApproveHistoryExportJobRequest) (response *CreateDspmApproveHistoryExportJobResponse, err error) {
    return c.CreateDspmApproveHistoryExportJobWithContext(context.Background(), request)
}

// CreateDspmApproveHistoryExportJob
// 创建Dspm审批历史导出任务
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateDspmApproveHistoryExportJobWithContext(ctx context.Context, request *CreateDspmApproveHistoryExportJobRequest) (response *CreateDspmApproveHistoryExportJobResponse, err error) {
    if request == nil {
        request = NewCreateDspmApproveHistoryExportJobRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "CreateDspmApproveHistoryExportJob")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateDspmApproveHistoryExportJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateDspmApproveHistoryExportJobResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDspmAssetAccessTopologyExportJobRequest() (request *CreateDspmAssetAccessTopologyExportJobRequest) {
    request = &CreateDspmAssetAccessTopologyExportJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "CreateDspmAssetAccessTopologyExportJob")
    
    
    return
}

func NewCreateDspmAssetAccessTopologyExportJobResponse() (response *CreateDspmAssetAccessTopologyExportJobResponse) {
    response = &CreateDspmAssetAccessTopologyExportJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateDspmAssetAccessTopologyExportJob
// 创建Dspm资产访问拓扑导出任务
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateDspmAssetAccessTopologyExportJob(request *CreateDspmAssetAccessTopologyExportJobRequest) (response *CreateDspmAssetAccessTopologyExportJobResponse, err error) {
    return c.CreateDspmAssetAccessTopologyExportJobWithContext(context.Background(), request)
}

// CreateDspmAssetAccessTopologyExportJob
// 创建Dspm资产访问拓扑导出任务
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateDspmAssetAccessTopologyExportJobWithContext(ctx context.Context, request *CreateDspmAssetAccessTopologyExportJobRequest) (response *CreateDspmAssetAccessTopologyExportJobResponse, err error) {
    if request == nil {
        request = NewCreateDspmAssetAccessTopologyExportJobRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "CreateDspmAssetAccessTopologyExportJob")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateDspmAssetAccessTopologyExportJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateDspmAssetAccessTopologyExportJobResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDspmAssetsExportJobRequest() (request *CreateDspmAssetsExportJobRequest) {
    request = &CreateDspmAssetsExportJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "CreateDspmAssetsExportJob")
    
    
    return
}

func NewCreateDspmAssetsExportJobResponse() (response *CreateDspmAssetsExportJobResponse) {
    response = &CreateDspmAssetsExportJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateDspmAssetsExportJob
// 创建Dspm资产列表导出任务
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateDspmAssetsExportJob(request *CreateDspmAssetsExportJobRequest) (response *CreateDspmAssetsExportJobResponse, err error) {
    return c.CreateDspmAssetsExportJobWithContext(context.Background(), request)
}

// CreateDspmAssetsExportJob
// 创建Dspm资产列表导出任务
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateDspmAssetsExportJobWithContext(ctx context.Context, request *CreateDspmAssetsExportJobRequest) (response *CreateDspmAssetsExportJobResponse, err error) {
    if request == nil {
        request = NewCreateDspmAssetsExportJobRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "CreateDspmAssetsExportJob")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateDspmAssetsExportJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateDspmAssetsExportJobResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDspmExportTaskRequest() (request *CreateDspmExportTaskRequest) {
    request = &CreateDspmExportTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "CreateDspmExportTask")
    
    
    return
}

func NewCreateDspmExportTaskResponse() (response *CreateDspmExportTaskResponse) {
    response = &CreateDspmExportTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateDspmExportTask
// 创建日志导出任务
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateDspmExportTask(request *CreateDspmExportTaskRequest) (response *CreateDspmExportTaskResponse, err error) {
    return c.CreateDspmExportTaskWithContext(context.Background(), request)
}

// CreateDspmExportTask
// 创建日志导出任务
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateDspmExportTaskWithContext(ctx context.Context, request *CreateDspmExportTaskRequest) (response *CreateDspmExportTaskResponse, err error) {
    if request == nil {
        request = NewCreateDspmExportTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "CreateDspmExportTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateDspmExportTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateDspmExportTaskResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDspmIdentifyInfoListExportJobRequest() (request *CreateDspmIdentifyInfoListExportJobRequest) {
    request = &CreateDspmIdentifyInfoListExportJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "CreateDspmIdentifyInfoListExportJob")
    
    
    return
}

func NewCreateDspmIdentifyInfoListExportJobResponse() (response *CreateDspmIdentifyInfoListExportJobResponse) {
    response = &CreateDspmIdentifyInfoListExportJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateDspmIdentifyInfoListExportJob
// 创建Dspm身份列表导出任务
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateDspmIdentifyInfoListExportJob(request *CreateDspmIdentifyInfoListExportJobRequest) (response *CreateDspmIdentifyInfoListExportJobResponse, err error) {
    return c.CreateDspmIdentifyInfoListExportJobWithContext(context.Background(), request)
}

// CreateDspmIdentifyInfoListExportJob
// 创建Dspm身份列表导出任务
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateDspmIdentifyInfoListExportJobWithContext(ctx context.Context, request *CreateDspmIdentifyInfoListExportJobRequest) (response *CreateDspmIdentifyInfoListExportJobResponse, err error) {
    if request == nil {
        request = NewCreateDspmIdentifyInfoListExportJobRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "CreateDspmIdentifyInfoListExportJob")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateDspmIdentifyInfoListExportJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateDspmIdentifyInfoListExportJobResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDspmPersonalIdentifyRequest() (request *CreateDspmPersonalIdentifyRequest) {
    request = &CreateDspmPersonalIdentifyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "CreateDspmPersonalIdentify")
    
    
    return
}

func NewCreateDspmPersonalIdentifyResponse() (response *CreateDspmPersonalIdentifyResponse) {
    response = &CreateDspmPersonalIdentifyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateDspmPersonalIdentify
// 创建Dspm个人身份id
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) CreateDspmPersonalIdentify(request *CreateDspmPersonalIdentifyRequest) (response *CreateDspmPersonalIdentifyResponse, err error) {
    return c.CreateDspmPersonalIdentifyWithContext(context.Background(), request)
}

// CreateDspmPersonalIdentify
// 创建Dspm个人身份id
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) CreateDspmPersonalIdentifyWithContext(ctx context.Context, request *CreateDspmPersonalIdentifyRequest) (response *CreateDspmPersonalIdentifyResponse, err error) {
    if request == nil {
        request = NewCreateDspmPersonalIdentifyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "CreateDspmPersonalIdentify")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateDspmPersonalIdentify require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateDspmPersonalIdentifyResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDspmRiskExportJobRequest() (request *CreateDspmRiskExportJobRequest) {
    request = &CreateDspmRiskExportJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "CreateDspmRiskExportJob")
    
    
    return
}

func NewCreateDspmRiskExportJobResponse() (response *CreateDspmRiskExportJobResponse) {
    response = &CreateDspmRiskExportJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateDspmRiskExportJob
// 创建Dspm风险导出任务
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) CreateDspmRiskExportJob(request *CreateDspmRiskExportJobRequest) (response *CreateDspmRiskExportJobResponse, err error) {
    return c.CreateDspmRiskExportJobWithContext(context.Background(), request)
}

// CreateDspmRiskExportJob
// 创建Dspm风险导出任务
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) CreateDspmRiskExportJobWithContext(ctx context.Context, request *CreateDspmRiskExportJobRequest) (response *CreateDspmRiskExportJobResponse, err error) {
    if request == nil {
        request = NewCreateDspmRiskExportJobRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "CreateDspmRiskExportJob")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateDspmRiskExportJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateDspmRiskExportJobResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDspmWhitelistStrategyRequest() (request *CreateDspmWhitelistStrategyRequest) {
    request = &CreateDspmWhitelistStrategyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "CreateDspmWhitelistStrategy")
    
    
    return
}

func NewCreateDspmWhitelistStrategyResponse() (response *CreateDspmWhitelistStrategyResponse) {
    response = &CreateDspmWhitelistStrategyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateDspmWhitelistStrategy
// 创建Dspm白名单策略
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) CreateDspmWhitelistStrategy(request *CreateDspmWhitelistStrategyRequest) (response *CreateDspmWhitelistStrategyResponse, err error) {
    return c.CreateDspmWhitelistStrategyWithContext(context.Background(), request)
}

// CreateDspmWhitelistStrategy
// 创建Dspm白名单策略
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) CreateDspmWhitelistStrategyWithContext(ctx context.Context, request *CreateDspmWhitelistStrategyRequest) (response *CreateDspmWhitelistStrategyResponse, err error) {
    if request == nil {
        request = NewCreateDspmWhitelistStrategyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "CreateDspmWhitelistStrategy")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateDspmWhitelistStrategy require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateDspmWhitelistStrategyResponse()
    err = c.Send(request, response)
    return
}

func NewCreateIaCAccessTokenRequest() (request *CreateIaCAccessTokenRequest) {
    request = &CreateIaCAccessTokenRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "CreateIaCAccessToken")
    
    
    return
}

func NewCreateIaCAccessTokenResponse() (response *CreateIaCAccessTokenResponse) {
    response = &CreateIaCAccessTokenResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateIaCAccessToken
// 创建IaC检测接入Token
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CreateIaCAccessToken(request *CreateIaCAccessTokenRequest) (response *CreateIaCAccessTokenResponse, err error) {
    return c.CreateIaCAccessTokenWithContext(context.Background(), request)
}

// CreateIaCAccessToken
// 创建IaC检测接入Token
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CreateIaCAccessTokenWithContext(ctx context.Context, request *CreateIaCAccessTokenRequest) (response *CreateIaCAccessTokenResponse, err error) {
    if request == nil {
        request = NewCreateIaCAccessTokenRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "CreateIaCAccessToken")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateIaCAccessToken require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateIaCAccessTokenResponse()
    err = c.Send(request, response)
    return
}

func NewCreateIaCFileExportJobRequest() (request *CreateIaCFileExportJobRequest) {
    request = &CreateIaCFileExportJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "CreateIaCFileExportJob")
    
    
    return
}

func NewCreateIaCFileExportJobResponse() (response *CreateIaCFileExportJobResponse) {
    response = &CreateIaCFileExportJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateIaCFileExportJob
// 创建IaC检测文件导出任务
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CreateIaCFileExportJob(request *CreateIaCFileExportJobRequest) (response *CreateIaCFileExportJobResponse, err error) {
    return c.CreateIaCFileExportJobWithContext(context.Background(), request)
}

// CreateIaCFileExportJob
// 创建IaC检测文件导出任务
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CreateIaCFileExportJobWithContext(ctx context.Context, request *CreateIaCFileExportJobRequest) (response *CreateIaCFileExportJobResponse, err error) {
    if request == nil {
        request = NewCreateIaCFileExportJobRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "CreateIaCFileExportJob")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateIaCFileExportJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateIaCFileExportJobResponse()
    err = c.Send(request, response)
    return
}

func NewCreateIaCFileReScanTaskRequest() (request *CreateIaCFileReScanTaskRequest) {
    request = &CreateIaCFileReScanTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "CreateIaCFileReScanTask")
    
    
    return
}

func NewCreateIaCFileReScanTaskResponse() (response *CreateIaCFileReScanTaskResponse) {
    response = &CreateIaCFileReScanTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateIaCFileReScanTask
// 创建IaC检测文件重新扫描任务
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateIaCFileReScanTask(request *CreateIaCFileReScanTaskRequest) (response *CreateIaCFileReScanTaskResponse, err error) {
    return c.CreateIaCFileReScanTaskWithContext(context.Background(), request)
}

// CreateIaCFileReScanTask
// 创建IaC检测文件重新扫描任务
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateIaCFileReScanTaskWithContext(ctx context.Context, request *CreateIaCFileReScanTaskRequest) (response *CreateIaCFileReScanTaskResponse, err error) {
    if request == nil {
        request = NewCreateIaCFileReScanTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "CreateIaCFileReScanTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateIaCFileReScanTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateIaCFileReScanTaskResponse()
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
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "CreateRiskCenterScanTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateRiskCenterScanTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateRiskCenterScanTaskResponse()
    err = c.Send(request, response)
    return
}

func NewCreateSkillScanRequest() (request *CreateSkillScanRequest) {
    request = &CreateSkillScanRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "CreateSkillScan")
    
    
    return
}

func NewCreateSkillScanResponse() (response *CreateSkillScanResponse) {
    response = &CreateSkillScanResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateSkillScan
// 上传 Skill ZIP 文件，触发异步安全检测。上传成功后应使用返回的 ContentHash + EngineVersion 轮询 DescribeSkillScanResult 接口获取结果。上传接口具备幂等性，同一 Hash 的文件重复上传不会创建重复任务。检测结果保留90天，超期后需重新上传检测。
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
func (c *Client) CreateSkillScan(request *CreateSkillScanRequest) (response *CreateSkillScanResponse, err error) {
    return c.CreateSkillScanWithContext(context.Background(), request)
}

// CreateSkillScan
// 上传 Skill ZIP 文件，触发异步安全检测。上传成功后应使用返回的 ContentHash + EngineVersion 轮询 DescribeSkillScanResult 接口获取结果。上传接口具备幂等性，同一 Hash 的文件重复上传不会创建重复任务。检测结果保留90天，超期后需重新上传检测。
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
func (c *Client) CreateSkillScanWithContext(ctx context.Context, request *CreateSkillScanRequest) (response *CreateSkillScanResponse, err error) {
    if request == nil {
        request = NewCreateSkillScanRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "CreateSkillScan")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateSkillScan require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateSkillScanResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteCosAkAssetRequest() (request *DeleteCosAkAssetRequest) {
    request = &DeleteCosAkAssetRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DeleteCosAkAsset")
    
    
    return
}

func NewDeleteCosAkAssetResponse() (response *DeleteCosAkAssetResponse) {
    response = &DeleteCosAkAssetResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteCosAkAsset
// 删除已删除的cos ak资产
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteCosAkAsset(request *DeleteCosAkAssetRequest) (response *DeleteCosAkAssetResponse, err error) {
    return c.DeleteCosAkAssetWithContext(context.Background(), request)
}

// DeleteCosAkAsset
// 删除已删除的cos ak资产
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteCosAkAssetWithContext(ctx context.Context, request *DeleteCosAkAssetRequest) (response *DeleteCosAkAssetResponse, err error) {
    if request == nil {
        request = NewDeleteCosAkAssetRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DeleteCosAkAsset")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteCosAkAsset require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteCosAkAssetResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteCosPolicyRequest() (request *DeleteCosPolicyRequest) {
    request = &DeleteCosPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DeleteCosPolicy")
    
    
    return
}

func NewDeleteCosPolicyResponse() (response *DeleteCosPolicyResponse) {
    response = &DeleteCosPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteCosPolicy
// 删除策略
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DeleteCosPolicy(request *DeleteCosPolicyRequest) (response *DeleteCosPolicyResponse, err error) {
    return c.DeleteCosPolicyWithContext(context.Background(), request)
}

// DeleteCosPolicy
// 删除策略
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DeleteCosPolicyWithContext(ctx context.Context, request *DeleteCosPolicyRequest) (response *DeleteCosPolicyResponse, err error) {
    if request == nil {
        request = NewDeleteCosPolicyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DeleteCosPolicy")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteCosPolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteCosPolicyResponse()
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
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DeleteDomainAndIp")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteDomainAndIp require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteDomainAndIpResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteDspmApplyOrderRequest() (request *DeleteDspmApplyOrderRequest) {
    request = &DeleteDspmApplyOrderRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DeleteDspmApplyOrder")
    
    
    return
}

func NewDeleteDspmApplyOrderResponse() (response *DeleteDspmApplyOrderResponse) {
    response = &DeleteDspmApplyOrderResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteDspmApplyOrder
// 删除Dspm申请单
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteDspmApplyOrder(request *DeleteDspmApplyOrderRequest) (response *DeleteDspmApplyOrderResponse, err error) {
    return c.DeleteDspmApplyOrderWithContext(context.Background(), request)
}

// DeleteDspmApplyOrder
// 删除Dspm申请单
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteDspmApplyOrderWithContext(ctx context.Context, request *DeleteDspmApplyOrderRequest) (response *DeleteDspmApplyOrderResponse, err error) {
    if request == nil {
        request = NewDeleteDspmApplyOrderRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DeleteDspmApplyOrder")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteDspmApplyOrder require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteDspmApplyOrderResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteDspmAssetAccountRequest() (request *DeleteDspmAssetAccountRequest) {
    request = &DeleteDspmAssetAccountRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DeleteDspmAssetAccount")
    
    
    return
}

func NewDeleteDspmAssetAccountResponse() (response *DeleteDspmAssetAccountResponse) {
    response = &DeleteDspmAssetAccountResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteDspmAssetAccount
// 删除Dspm资产账号
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteDspmAssetAccount(request *DeleteDspmAssetAccountRequest) (response *DeleteDspmAssetAccountResponse, err error) {
    return c.DeleteDspmAssetAccountWithContext(context.Background(), request)
}

// DeleteDspmAssetAccount
// 删除Dspm资产账号
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteDspmAssetAccountWithContext(ctx context.Context, request *DeleteDspmAssetAccountRequest) (response *DeleteDspmAssetAccountResponse, err error) {
    if request == nil {
        request = NewDeleteDspmAssetAccountRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DeleteDspmAssetAccount")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteDspmAssetAccount require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteDspmAssetAccountResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteDspmBackupLogListRequest() (request *DeleteDspmBackupLogListRequest) {
    request = &DeleteDspmBackupLogListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DeleteDspmBackupLogList")
    
    
    return
}

func NewDeleteDspmBackupLogListResponse() (response *DeleteDspmBackupLogListResponse) {
    response = &DeleteDspmBackupLogListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteDspmBackupLogList
// 删除备份日志
//
// 可能返回的错误码:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteDspmBackupLogList(request *DeleteDspmBackupLogListRequest) (response *DeleteDspmBackupLogListResponse, err error) {
    return c.DeleteDspmBackupLogListWithContext(context.Background(), request)
}

// DeleteDspmBackupLogList
// 删除备份日志
//
// 可能返回的错误码:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteDspmBackupLogListWithContext(ctx context.Context, request *DeleteDspmBackupLogListRequest) (response *DeleteDspmBackupLogListResponse, err error) {
    if request == nil {
        request = NewDeleteDspmBackupLogListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DeleteDspmBackupLogList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteDspmBackupLogList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteDspmBackupLogListResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteDspmExportTaskRequest() (request *DeleteDspmExportTaskRequest) {
    request = &DeleteDspmExportTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DeleteDspmExportTask")
    
    
    return
}

func NewDeleteDspmExportTaskResponse() (response *DeleteDspmExportTaskResponse) {
    response = &DeleteDspmExportTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteDspmExportTask
// 删除导出任务
//
// 可能返回的错误码:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteDspmExportTask(request *DeleteDspmExportTaskRequest) (response *DeleteDspmExportTaskResponse, err error) {
    return c.DeleteDspmExportTaskWithContext(context.Background(), request)
}

// DeleteDspmExportTask
// 删除导出任务
//
// 可能返回的错误码:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteDspmExportTaskWithContext(ctx context.Context, request *DeleteDspmExportTaskRequest) (response *DeleteDspmExportTaskResponse, err error) {
    if request == nil {
        request = NewDeleteDspmExportTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DeleteDspmExportTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteDspmExportTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteDspmExportTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteDspmPersonalIdentifyRequest() (request *DeleteDspmPersonalIdentifyRequest) {
    request = &DeleteDspmPersonalIdentifyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DeleteDspmPersonalIdentify")
    
    
    return
}

func NewDeleteDspmPersonalIdentifyResponse() (response *DeleteDspmPersonalIdentifyResponse) {
    response = &DeleteDspmPersonalIdentifyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteDspmPersonalIdentify
// 删除Dspm个人身份id
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DeleteDspmPersonalIdentify(request *DeleteDspmPersonalIdentifyRequest) (response *DeleteDspmPersonalIdentifyResponse, err error) {
    return c.DeleteDspmPersonalIdentifyWithContext(context.Background(), request)
}

// DeleteDspmPersonalIdentify
// 删除Dspm个人身份id
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DeleteDspmPersonalIdentifyWithContext(ctx context.Context, request *DeleteDspmPersonalIdentifyRequest) (response *DeleteDspmPersonalIdentifyResponse, err error) {
    if request == nil {
        request = NewDeleteDspmPersonalIdentifyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DeleteDspmPersonalIdentify")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteDspmPersonalIdentify require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteDspmPersonalIdentifyResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteDspmRestoreLogListRequest() (request *DeleteDspmRestoreLogListRequest) {
    request = &DeleteDspmRestoreLogListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DeleteDspmRestoreLogList")
    
    
    return
}

func NewDeleteDspmRestoreLogListResponse() (response *DeleteDspmRestoreLogListResponse) {
    response = &DeleteDspmRestoreLogListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteDspmRestoreLogList
// 删除恢复日志
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DeleteDspmRestoreLogList(request *DeleteDspmRestoreLogListRequest) (response *DeleteDspmRestoreLogListResponse, err error) {
    return c.DeleteDspmRestoreLogListWithContext(context.Background(), request)
}

// DeleteDspmRestoreLogList
// 删除恢复日志
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DeleteDspmRestoreLogListWithContext(ctx context.Context, request *DeleteDspmRestoreLogListRequest) (response *DeleteDspmRestoreLogListResponse, err error) {
    if request == nil {
        request = NewDeleteDspmRestoreLogListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DeleteDspmRestoreLogList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteDspmRestoreLogList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteDspmRestoreLogListResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteDspmWhitelistStrategyRequest() (request *DeleteDspmWhitelistStrategyRequest) {
    request = &DeleteDspmWhitelistStrategyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DeleteDspmWhitelistStrategy")
    
    
    return
}

func NewDeleteDspmWhitelistStrategyResponse() (response *DeleteDspmWhitelistStrategyResponse) {
    response = &DeleteDspmWhitelistStrategyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteDspmWhitelistStrategy
// 删除Dspm白名单策略
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DeleteDspmWhitelistStrategy(request *DeleteDspmWhitelistStrategyRequest) (response *DeleteDspmWhitelistStrategyResponse, err error) {
    return c.DeleteDspmWhitelistStrategyWithContext(context.Background(), request)
}

// DeleteDspmWhitelistStrategy
// 删除Dspm白名单策略
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DeleteDspmWhitelistStrategyWithContext(ctx context.Context, request *DeleteDspmWhitelistStrategyRequest) (response *DeleteDspmWhitelistStrategyResponse, err error) {
    if request == nil {
        request = NewDeleteDspmWhitelistStrategyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DeleteDspmWhitelistStrategy")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteDspmWhitelistStrategy require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteDspmWhitelistStrategyResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteIaCAccessTokenRequest() (request *DeleteIaCAccessTokenRequest) {
    request = &DeleteIaCAccessTokenRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DeleteIaCAccessToken")
    
    
    return
}

func NewDeleteIaCAccessTokenResponse() (response *DeleteIaCAccessTokenResponse) {
    response = &DeleteIaCAccessTokenResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteIaCAccessToken
// 删除IaC检测接入Token
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DeleteIaCAccessToken(request *DeleteIaCAccessTokenRequest) (response *DeleteIaCAccessTokenResponse, err error) {
    return c.DeleteIaCAccessTokenWithContext(context.Background(), request)
}

// DeleteIaCAccessToken
// 删除IaC检测接入Token
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DeleteIaCAccessTokenWithContext(ctx context.Context, request *DeleteIaCAccessTokenRequest) (response *DeleteIaCAccessTokenResponse, err error) {
    if request == nil {
        request = NewDeleteIaCAccessTokenRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DeleteIaCAccessToken")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteIaCAccessToken require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteIaCAccessTokenResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteIaCFileRequest() (request *DeleteIaCFileRequest) {
    request = &DeleteIaCFileRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DeleteIaCFile")
    
    
    return
}

func NewDeleteIaCFileResponse() (response *DeleteIaCFileResponse) {
    response = &DeleteIaCFileResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteIaCFile
// 删除IaC检测文件
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DeleteIaCFile(request *DeleteIaCFileRequest) (response *DeleteIaCFileResponse, err error) {
    return c.DeleteIaCFileWithContext(context.Background(), request)
}

// DeleteIaCFile
// 删除IaC检测文件
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DeleteIaCFileWithContext(ctx context.Context, request *DeleteIaCFileRequest) (response *DeleteIaCFileResponse, err error) {
    if request == nil {
        request = NewDeleteIaCFileRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DeleteIaCFile")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteIaCFile require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteIaCFileResponse()
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
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DeleteRiskScanTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteRiskScanTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteRiskScanTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAIAgentAssetListRequest() (request *DescribeAIAgentAssetListRequest) {
    request = &DescribeAIAgentAssetListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeAIAgentAssetList")
    
    
    return
}

func NewDescribeAIAgentAssetListResponse() (response *DescribeAIAgentAssetListResponse) {
    response = &DescribeAIAgentAssetListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAIAgentAssetList
// 获取 AI agent 资产列表
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
func (c *Client) DescribeAIAgentAssetList(request *DescribeAIAgentAssetListRequest) (response *DescribeAIAgentAssetListResponse, err error) {
    return c.DescribeAIAgentAssetListWithContext(context.Background(), request)
}

// DescribeAIAgentAssetList
// 获取 AI agent 资产列表
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
func (c *Client) DescribeAIAgentAssetListWithContext(ctx context.Context, request *DescribeAIAgentAssetListRequest) (response *DescribeAIAgentAssetListResponse, err error) {
    if request == nil {
        request = NewDescribeAIAgentAssetListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeAIAgentAssetList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAIAgentAssetList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAIAgentAssetListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAKAnalysisDetailRequest() (request *DescribeAKAnalysisDetailRequest) {
    request = &DescribeAKAnalysisDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeAKAnalysisDetail")
    
    
    return
}

func NewDescribeAKAnalysisDetailResponse() (response *DescribeAKAnalysisDetailResponse) {
    response = &DescribeAKAnalysisDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAKAnalysisDetail
// 访问密钥告警记录AI分析结果详情
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
func (c *Client) DescribeAKAnalysisDetail(request *DescribeAKAnalysisDetailRequest) (response *DescribeAKAnalysisDetailResponse, err error) {
    return c.DescribeAKAnalysisDetailWithContext(context.Background(), request)
}

// DescribeAKAnalysisDetail
// 访问密钥告警记录AI分析结果详情
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
func (c *Client) DescribeAKAnalysisDetailWithContext(ctx context.Context, request *DescribeAKAnalysisDetailRequest) (response *DescribeAKAnalysisDetailResponse, err error) {
    if request == nil {
        request = NewDescribeAKAnalysisDetailRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeAKAnalysisDetail")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAKAnalysisDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAKAnalysisDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAbnormalCallRecordRequest() (request *DescribeAbnormalCallRecordRequest) {
    request = &DescribeAbnormalCallRecordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeAbnormalCallRecord")
    
    
    return
}

func NewDescribeAbnormalCallRecordResponse() (response *DescribeAbnormalCallRecordResponse) {
    response = &DescribeAbnormalCallRecordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAbnormalCallRecord
// 获取调用记录列表
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
func (c *Client) DescribeAbnormalCallRecord(request *DescribeAbnormalCallRecordRequest) (response *DescribeAbnormalCallRecordResponse, err error) {
    return c.DescribeAbnormalCallRecordWithContext(context.Background(), request)
}

// DescribeAbnormalCallRecord
// 获取调用记录列表
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
func (c *Client) DescribeAbnormalCallRecordWithContext(ctx context.Context, request *DescribeAbnormalCallRecordRequest) (response *DescribeAbnormalCallRecordResponse, err error) {
    if request == nil {
        request = NewDescribeAbnormalCallRecordRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeAbnormalCallRecord")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAbnormalCallRecord require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAbnormalCallRecordResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAccessKeyAlarmRequest() (request *DescribeAccessKeyAlarmRequest) {
    request = &DescribeAccessKeyAlarmRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeAccessKeyAlarm")
    
    
    return
}

func NewDescribeAccessKeyAlarmResponse() (response *DescribeAccessKeyAlarmResponse) {
    response = &DescribeAccessKeyAlarmResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAccessKeyAlarm
// 访问密钥告警记录列表
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
func (c *Client) DescribeAccessKeyAlarm(request *DescribeAccessKeyAlarmRequest) (response *DescribeAccessKeyAlarmResponse, err error) {
    return c.DescribeAccessKeyAlarmWithContext(context.Background(), request)
}

// DescribeAccessKeyAlarm
// 访问密钥告警记录列表
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
func (c *Client) DescribeAccessKeyAlarmWithContext(ctx context.Context, request *DescribeAccessKeyAlarmRequest) (response *DescribeAccessKeyAlarmResponse, err error) {
    if request == nil {
        request = NewDescribeAccessKeyAlarmRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeAccessKeyAlarm")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAccessKeyAlarm require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAccessKeyAlarmResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAccessKeyAlarmDetailRequest() (request *DescribeAccessKeyAlarmDetailRequest) {
    request = &DescribeAccessKeyAlarmDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeAccessKeyAlarmDetail")
    
    
    return
}

func NewDescribeAccessKeyAlarmDetailResponse() (response *DescribeAccessKeyAlarmDetailResponse) {
    response = &DescribeAccessKeyAlarmDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAccessKeyAlarmDetail
// 访问密钥告警记录详情
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
func (c *Client) DescribeAccessKeyAlarmDetail(request *DescribeAccessKeyAlarmDetailRequest) (response *DescribeAccessKeyAlarmDetailResponse, err error) {
    return c.DescribeAccessKeyAlarmDetailWithContext(context.Background(), request)
}

// DescribeAccessKeyAlarmDetail
// 访问密钥告警记录详情
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
func (c *Client) DescribeAccessKeyAlarmDetailWithContext(ctx context.Context, request *DescribeAccessKeyAlarmDetailRequest) (response *DescribeAccessKeyAlarmDetailResponse, err error) {
    if request == nil {
        request = NewDescribeAccessKeyAlarmDetailRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeAccessKeyAlarmDetail")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAccessKeyAlarmDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAccessKeyAlarmDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAccessKeyAssetRequest() (request *DescribeAccessKeyAssetRequest) {
    request = &DescribeAccessKeyAssetRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeAccessKeyAsset")
    
    
    return
}

func NewDescribeAccessKeyAssetResponse() (response *DescribeAccessKeyAssetResponse) {
    response = &DescribeAccessKeyAssetResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAccessKeyAsset
// 获取用户访问密钥资产列表
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
func (c *Client) DescribeAccessKeyAsset(request *DescribeAccessKeyAssetRequest) (response *DescribeAccessKeyAssetResponse, err error) {
    return c.DescribeAccessKeyAssetWithContext(context.Background(), request)
}

// DescribeAccessKeyAsset
// 获取用户访问密钥资产列表
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
func (c *Client) DescribeAccessKeyAssetWithContext(ctx context.Context, request *DescribeAccessKeyAssetRequest) (response *DescribeAccessKeyAssetResponse, err error) {
    if request == nil {
        request = NewDescribeAccessKeyAssetRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeAccessKeyAsset")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAccessKeyAsset require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAccessKeyAssetResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAccessKeyRiskRequest() (request *DescribeAccessKeyRiskRequest) {
    request = &DescribeAccessKeyRiskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeAccessKeyRisk")
    
    
    return
}

func NewDescribeAccessKeyRiskResponse() (response *DescribeAccessKeyRiskResponse) {
    response = &DescribeAccessKeyRiskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAccessKeyRisk
// 访问密钥风险记录列表
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
func (c *Client) DescribeAccessKeyRisk(request *DescribeAccessKeyRiskRequest) (response *DescribeAccessKeyRiskResponse, err error) {
    return c.DescribeAccessKeyRiskWithContext(context.Background(), request)
}

// DescribeAccessKeyRisk
// 访问密钥风险记录列表
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
func (c *Client) DescribeAccessKeyRiskWithContext(ctx context.Context, request *DescribeAccessKeyRiskRequest) (response *DescribeAccessKeyRiskResponse, err error) {
    if request == nil {
        request = NewDescribeAccessKeyRiskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeAccessKeyRisk")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAccessKeyRisk require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAccessKeyRiskResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAccessKeyRiskDetailRequest() (request *DescribeAccessKeyRiskDetailRequest) {
    request = &DescribeAccessKeyRiskDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeAccessKeyRiskDetail")
    
    
    return
}

func NewDescribeAccessKeyRiskDetailResponse() (response *DescribeAccessKeyRiskDetailResponse) {
    response = &DescribeAccessKeyRiskDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAccessKeyRiskDetail
// 访问密钥风险记录详情
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
func (c *Client) DescribeAccessKeyRiskDetail(request *DescribeAccessKeyRiskDetailRequest) (response *DescribeAccessKeyRiskDetailResponse, err error) {
    return c.DescribeAccessKeyRiskDetailWithContext(context.Background(), request)
}

// DescribeAccessKeyRiskDetail
// 访问密钥风险记录详情
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
func (c *Client) DescribeAccessKeyRiskDetailWithContext(ctx context.Context, request *DescribeAccessKeyRiskDetailRequest) (response *DescribeAccessKeyRiskDetailResponse, err error) {
    if request == nil {
        request = NewDescribeAccessKeyRiskDetailRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeAccessKeyRiskDetail")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAccessKeyRiskDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAccessKeyRiskDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAccessKeyUserDetailRequest() (request *DescribeAccessKeyUserDetailRequest) {
    request = &DescribeAccessKeyUserDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeAccessKeyUserDetail")
    
    
    return
}

func NewDescribeAccessKeyUserDetailResponse() (response *DescribeAccessKeyUserDetailResponse) {
    response = &DescribeAccessKeyUserDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAccessKeyUserDetail
// 查询用户的账号详情
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
func (c *Client) DescribeAccessKeyUserDetail(request *DescribeAccessKeyUserDetailRequest) (response *DescribeAccessKeyUserDetailResponse, err error) {
    return c.DescribeAccessKeyUserDetailWithContext(context.Background(), request)
}

// DescribeAccessKeyUserDetail
// 查询用户的账号详情
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
func (c *Client) DescribeAccessKeyUserDetailWithContext(ctx context.Context, request *DescribeAccessKeyUserDetailRequest) (response *DescribeAccessKeyUserDetailResponse, err error) {
    if request == nil {
        request = NewDescribeAccessKeyUserDetailRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeAccessKeyUserDetail")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAccessKeyUserDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAccessKeyUserDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAccessKeyUserListRequest() (request *DescribeAccessKeyUserListRequest) {
    request = &DescribeAccessKeyUserListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeAccessKeyUserList")
    
    
    return
}

func NewDescribeAccessKeyUserListResponse() (response *DescribeAccessKeyUserListResponse) {
    response = &DescribeAccessKeyUserListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAccessKeyUserList
// 查询用户的账号列表
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
func (c *Client) DescribeAccessKeyUserList(request *DescribeAccessKeyUserListRequest) (response *DescribeAccessKeyUserListResponse, err error) {
    return c.DescribeAccessKeyUserListWithContext(context.Background(), request)
}

// DescribeAccessKeyUserList
// 查询用户的账号列表
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
func (c *Client) DescribeAccessKeyUserListWithContext(ctx context.Context, request *DescribeAccessKeyUserListRequest) (response *DescribeAccessKeyUserListResponse, err error) {
    if request == nil {
        request = NewDescribeAccessKeyUserListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeAccessKeyUserList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAccessKeyUserList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAccessKeyUserListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAlertListRequest() (request *DescribeAlertListRequest) {
    request = &DescribeAlertListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeAlertList")
    
    
    return
}

func NewDescribeAlertListResponse() (response *DescribeAlertListResponse) {
    response = &DescribeAlertListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAlertList
// 告警中心全量告警列表接口
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
func (c *Client) DescribeAlertList(request *DescribeAlertListRequest) (response *DescribeAlertListResponse, err error) {
    return c.DescribeAlertListWithContext(context.Background(), request)
}

// DescribeAlertList
// 告警中心全量告警列表接口
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
func (c *Client) DescribeAlertListWithContext(ctx context.Context, request *DescribeAlertListRequest) (response *DescribeAlertListResponse, err error) {
    if request == nil {
        request = NewDescribeAlertListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeAlertList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAlertList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAlertListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAssetProcessListRequest() (request *DescribeAssetProcessListRequest) {
    request = &DescribeAssetProcessListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeAssetProcessList")
    
    
    return
}

func NewDescribeAssetProcessListResponse() (response *DescribeAssetProcessListResponse) {
    response = &DescribeAssetProcessListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAssetProcessList
// 查询云边界分析-暴露路径下主机节点的进程列表
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
func (c *Client) DescribeAssetProcessList(request *DescribeAssetProcessListRequest) (response *DescribeAssetProcessListResponse, err error) {
    return c.DescribeAssetProcessListWithContext(context.Background(), request)
}

// DescribeAssetProcessList
// 查询云边界分析-暴露路径下主机节点的进程列表
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
func (c *Client) DescribeAssetProcessListWithContext(ctx context.Context, request *DescribeAssetProcessListRequest) (response *DescribeAssetProcessListResponse, err error) {
    if request == nil {
        request = NewDescribeAssetProcessListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeAssetProcessList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAssetProcessList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAssetProcessListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAssetRiskListRequest() (request *DescribeAssetRiskListRequest) {
    request = &DescribeAssetRiskListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeAssetRiskList")
    
    
    return
}

func NewDescribeAssetRiskListResponse() (response *DescribeAssetRiskListResponse) {
    response = &DescribeAssetRiskListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAssetRiskList
// 资产视角下云资源配置风险列表
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
func (c *Client) DescribeAssetRiskList(request *DescribeAssetRiskListRequest) (response *DescribeAssetRiskListResponse, err error) {
    return c.DescribeAssetRiskListWithContext(context.Background(), request)
}

// DescribeAssetRiskList
// 资产视角下云资源配置风险列表
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
func (c *Client) DescribeAssetRiskListWithContext(ctx context.Context, request *DescribeAssetRiskListRequest) (response *DescribeAssetRiskListResponse, err error) {
    if request == nil {
        request = NewDescribeAssetRiskListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeAssetRiskList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAssetRiskList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAssetRiskListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAssetViewVulRiskListRequest() (request *DescribeAssetViewVulRiskListRequest) {
    request = &DescribeAssetViewVulRiskListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeAssetViewVulRiskList")
    
    
    return
}

func NewDescribeAssetViewVulRiskListResponse() (response *DescribeAssetViewVulRiskListResponse) {
    response = &DescribeAssetViewVulRiskListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAssetViewVulRiskList
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
func (c *Client) DescribeAssetViewVulRiskList(request *DescribeAssetViewVulRiskListRequest) (response *DescribeAssetViewVulRiskListResponse, err error) {
    return c.DescribeAssetViewVulRiskListWithContext(context.Background(), request)
}

// DescribeAssetViewVulRiskList
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
func (c *Client) DescribeAssetViewVulRiskListWithContext(ctx context.Context, request *DescribeAssetViewVulRiskListRequest) (response *DescribeAssetViewVulRiskListResponse, err error) {
    if request == nil {
        request = NewDescribeAssetViewVulRiskListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeAssetViewVulRiskList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAssetViewVulRiskList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAssetViewVulRiskListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAssumeRoleRequest() (request *DescribeAssumeRoleRequest) {
    request = &DescribeAssumeRoleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeAssumeRole")
    
    
    return
}

func NewDescribeAssumeRoleResponse() (response *DescribeAssumeRoleResponse) {
    response = &DescribeAssumeRoleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAssumeRole
// 查询是否绑定角色
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAssumeRole(request *DescribeAssumeRoleRequest) (response *DescribeAssumeRoleResponse, err error) {
    return c.DescribeAssumeRoleWithContext(context.Background(), request)
}

// DescribeAssumeRole
// 查询是否绑定角色
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAssumeRoleWithContext(ctx context.Context, request *DescribeAssumeRoleRequest) (response *DescribeAssumeRoleResponse, err error) {
    if request == nil {
        request = NewDescribeAssumeRoleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeAssumeRole")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAssumeRole require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAssumeRoleResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBucketInvokeIpListRequest() (request *DescribeBucketInvokeIpListRequest) {
    request = &DescribeBucketInvokeIpListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeBucketInvokeIpList")
    
    
    return
}

func NewDescribeBucketInvokeIpListResponse() (response *DescribeBucketInvokeIpListResponse) {
    response = &DescribeBucketInvokeIpListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeBucketInvokeIpList
// 查看存储桶调用源ip列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
func (c *Client) DescribeBucketInvokeIpList(request *DescribeBucketInvokeIpListRequest) (response *DescribeBucketInvokeIpListResponse, err error) {
    return c.DescribeBucketInvokeIpListWithContext(context.Background(), request)
}

// DescribeBucketInvokeIpList
// 查看存储桶调用源ip列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
func (c *Client) DescribeBucketInvokeIpListWithContext(ctx context.Context, request *DescribeBucketInvokeIpListRequest) (response *DescribeBucketInvokeIpListResponse, err error) {
    if request == nil {
        request = NewDescribeBucketInvokeIpListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeBucketInvokeIpList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBucketInvokeIpList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBucketInvokeIpListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCFWAssetStatisticsRequest() (request *DescribeCFWAssetStatisticsRequest) {
    request = &DescribeCFWAssetStatisticsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeCFWAssetStatistics")
    
    
    return
}

func NewDescribeCFWAssetStatisticsResponse() (response *DescribeCFWAssetStatisticsResponse) {
    response = &DescribeCFWAssetStatisticsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCFWAssetStatistics
// 云防资产中心统计数据
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
func (c *Client) DescribeCFWAssetStatistics(request *DescribeCFWAssetStatisticsRequest) (response *DescribeCFWAssetStatisticsResponse, err error) {
    return c.DescribeCFWAssetStatisticsWithContext(context.Background(), request)
}

// DescribeCFWAssetStatistics
// 云防资产中心统计数据
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
func (c *Client) DescribeCFWAssetStatisticsWithContext(ctx context.Context, request *DescribeCFWAssetStatisticsRequest) (response *DescribeCFWAssetStatisticsResponse, err error) {
    if request == nil {
        request = NewDescribeCFWAssetStatisticsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeCFWAssetStatistics")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCFWAssetStatistics require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCFWAssetStatisticsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCSIPRiskStatisticsRequest() (request *DescribeCSIPRiskStatisticsRequest) {
    request = &DescribeCSIPRiskStatisticsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeCSIPRiskStatistics")
    
    
    return
}

func NewDescribeCSIPRiskStatisticsResponse() (response *DescribeCSIPRiskStatisticsResponse) {
    response = &DescribeCSIPRiskStatisticsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCSIPRiskStatistics
// 获取风险中心风险概况示例
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
func (c *Client) DescribeCSIPRiskStatistics(request *DescribeCSIPRiskStatisticsRequest) (response *DescribeCSIPRiskStatisticsResponse, err error) {
    return c.DescribeCSIPRiskStatisticsWithContext(context.Background(), request)
}

// DescribeCSIPRiskStatistics
// 获取风险中心风险概况示例
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
func (c *Client) DescribeCSIPRiskStatisticsWithContext(ctx context.Context, request *DescribeCSIPRiskStatisticsRequest) (response *DescribeCSIPRiskStatisticsResponse, err error) {
    if request == nil {
        request = NewDescribeCSIPRiskStatisticsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeCSIPRiskStatistics")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCSIPRiskStatistics require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCSIPRiskStatisticsResponse()
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
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
func (c *Client) DescribeCVMAssetInfo(request *DescribeCVMAssetInfoRequest) (response *DescribeCVMAssetInfoResponse, err error) {
    return c.DescribeCVMAssetInfoWithContext(context.Background(), request)
}

// DescribeCVMAssetInfo
// cvm详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
func (c *Client) DescribeCVMAssetInfoWithContext(ctx context.Context, request *DescribeCVMAssetInfoRequest) (response *DescribeCVMAssetInfoResponse, err error) {
    if request == nil {
        request = NewDescribeCVMAssetInfoRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeCVMAssetInfo")
    
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
// 获取cvm列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
func (c *Client) DescribeCVMAssets(request *DescribeCVMAssetsRequest) (response *DescribeCVMAssetsResponse, err error) {
    return c.DescribeCVMAssetsWithContext(context.Background(), request)
}

// DescribeCVMAssets
// 获取cvm列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
func (c *Client) DescribeCVMAssetsWithContext(ctx context.Context, request *DescribeCVMAssetsRequest) (response *DescribeCVMAssetsResponse, err error) {
    if request == nil {
        request = NewDescribeCVMAssetsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeCVMAssets")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCVMAssets require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCVMAssetsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCWPMachineDetailRequest() (request *DescribeCWPMachineDetailRequest) {
    request = &DescribeCWPMachineDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeCWPMachineDetail")
    
    
    return
}

func NewDescribeCWPMachineDetailResponse() (response *DescribeCWPMachineDetailResponse) {
    response = &DescribeCWPMachineDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCWPMachineDetail
// 主机详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
func (c *Client) DescribeCWPMachineDetail(request *DescribeCWPMachineDetailRequest) (response *DescribeCWPMachineDetailResponse, err error) {
    return c.DescribeCWPMachineDetailWithContext(context.Background(), request)
}

// DescribeCWPMachineDetail
// 主机详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
func (c *Client) DescribeCWPMachineDetailWithContext(ctx context.Context, request *DescribeCWPMachineDetailRequest) (response *DescribeCWPMachineDetailResponse, err error) {
    if request == nil {
        request = NewDescribeCWPMachineDetailRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeCWPMachineDetail")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCWPMachineDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCWPMachineDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCWPMachinesRequest() (request *DescribeCWPMachinesRequest) {
    request = &DescribeCWPMachinesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeCWPMachines")
    
    
    return
}

func NewDescribeCWPMachinesResponse() (response *DescribeCWPMachinesResponse) {
    response = &DescribeCWPMachinesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCWPMachines
// 主机列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
func (c *Client) DescribeCWPMachines(request *DescribeCWPMachinesRequest) (response *DescribeCWPMachinesResponse, err error) {
    return c.DescribeCWPMachinesWithContext(context.Background(), request)
}

// DescribeCWPMachines
// 主机列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
func (c *Client) DescribeCWPMachinesWithContext(ctx context.Context, request *DescribeCWPMachinesRequest) (response *DescribeCWPMachinesResponse, err error) {
    if request == nil {
        request = NewDescribeCWPMachinesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeCWPMachines")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCWPMachines require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCWPMachinesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCallRecordRequest() (request *DescribeCallRecordRequest) {
    request = &DescribeCallRecordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeCallRecord")
    
    
    return
}

func NewDescribeCallRecordResponse() (response *DescribeCallRecordResponse) {
    response = &DescribeCallRecordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCallRecord
// 获取调用记录列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
func (c *Client) DescribeCallRecord(request *DescribeCallRecordRequest) (response *DescribeCallRecordResponse, err error) {
    return c.DescribeCallRecordWithContext(context.Background(), request)
}

// DescribeCallRecord
// 获取调用记录列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
func (c *Client) DescribeCallRecordWithContext(ctx context.Context, request *DescribeCallRecordRequest) (response *DescribeCallRecordResponse, err error) {
    if request == nil {
        request = NewDescribeCallRecordRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeCallRecord")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCallRecord require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCallRecordResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCheckViewRisksRequest() (request *DescribeCheckViewRisksRequest) {
    request = &DescribeCheckViewRisksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeCheckViewRisks")
    
    
    return
}

func NewDescribeCheckViewRisksResponse() (response *DescribeCheckViewRisksResponse) {
    response = &DescribeCheckViewRisksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCheckViewRisks
// 检查视角下云资源配置风险列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
func (c *Client) DescribeCheckViewRisks(request *DescribeCheckViewRisksRequest) (response *DescribeCheckViewRisksResponse, err error) {
    return c.DescribeCheckViewRisksWithContext(context.Background(), request)
}

// DescribeCheckViewRisks
// 检查视角下云资源配置风险列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
func (c *Client) DescribeCheckViewRisksWithContext(ctx context.Context, request *DescribeCheckViewRisksRequest) (response *DescribeCheckViewRisksResponse, err error) {
    if request == nil {
        request = NewDescribeCheckViewRisksRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeCheckViewRisks")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCheckViewRisks require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCheckViewRisksResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeClusterAssetsRequest() (request *DescribeClusterAssetsRequest) {
    request = &DescribeClusterAssetsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeClusterAssets")
    
    
    return
}

func NewDescribeClusterAssetsResponse() (response *DescribeClusterAssetsResponse) {
    response = &DescribeClusterAssetsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeClusterAssets
// 集群列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
func (c *Client) DescribeClusterAssets(request *DescribeClusterAssetsRequest) (response *DescribeClusterAssetsResponse, err error) {
    return c.DescribeClusterAssetsWithContext(context.Background(), request)
}

// DescribeClusterAssets
// 集群列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
func (c *Client) DescribeClusterAssetsWithContext(ctx context.Context, request *DescribeClusterAssetsRequest) (response *DescribeClusterAssetsResponse, err error) {
    if request == nil {
        request = NewDescribeClusterAssetsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeClusterAssets")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeClusterAssets require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeClusterAssetsResponse()
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
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
func (c *Client) DescribeClusterPodAssets(request *DescribeClusterPodAssetsRequest) (response *DescribeClusterPodAssetsResponse, err error) {
    return c.DescribeClusterPodAssetsWithContext(context.Background(), request)
}

// DescribeClusterPodAssets
// 集群pod列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
func (c *Client) DescribeClusterPodAssetsWithContext(ctx context.Context, request *DescribeClusterPodAssetsRequest) (response *DescribeClusterPodAssetsResponse, err error) {
    if request == nil {
        request = NewDescribeClusterPodAssetsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeClusterPodAssets")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeClusterPodAssets require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeClusterPodAssetsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeConfigCheckRulesRequest() (request *DescribeConfigCheckRulesRequest) {
    request = &DescribeConfigCheckRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeConfigCheckRules")
    
    
    return
}

func NewDescribeConfigCheckRulesResponse() (response *DescribeConfigCheckRulesResponse) {
    response = &DescribeConfigCheckRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeConfigCheckRules
// 云资源配置风险规则列表示例
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
func (c *Client) DescribeConfigCheckRules(request *DescribeConfigCheckRulesRequest) (response *DescribeConfigCheckRulesResponse, err error) {
    return c.DescribeConfigCheckRulesWithContext(context.Background(), request)
}

// DescribeConfigCheckRules
// 云资源配置风险规则列表示例
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
func (c *Client) DescribeConfigCheckRulesWithContext(ctx context.Context, request *DescribeConfigCheckRulesRequest) (response *DescribeConfigCheckRulesResponse, err error) {
    if request == nil {
        request = NewDescribeConfigCheckRulesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeConfigCheckRules")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeConfigCheckRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeConfigCheckRulesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCosAccessPermissionRequest() (request *DescribeCosAccessPermissionRequest) {
    request = &DescribeCosAccessPermissionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeCosAccessPermission")
    
    
    return
}

func NewDescribeCosAccessPermissionResponse() (response *DescribeCosAccessPermissionResponse) {
    response = &DescribeCosAccessPermissionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCosAccessPermission
// 查看cos桶访问权限信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DECRYPTERROR = "InvalidParameter.DecryptError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeCosAccessPermission(request *DescribeCosAccessPermissionRequest) (response *DescribeCosAccessPermissionResponse, err error) {
    return c.DescribeCosAccessPermissionWithContext(context.Background(), request)
}

// DescribeCosAccessPermission
// 查看cos桶访问权限信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DECRYPTERROR = "InvalidParameter.DecryptError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeCosAccessPermissionWithContext(ctx context.Context, request *DescribeCosAccessPermissionRequest) (response *DescribeCosAccessPermissionResponse, err error) {
    if request == nil {
        request = NewDescribeCosAccessPermissionRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeCosAccessPermission")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCosAccessPermission require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCosAccessPermissionResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCosAccessPermissionsRequest() (request *DescribeCosAccessPermissionsRequest) {
    request = &DescribeCosAccessPermissionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeCosAccessPermissions")
    
    
    return
}

func NewDescribeCosAccessPermissionsResponse() (response *DescribeCosAccessPermissionsResponse) {
    response = &DescribeCosAccessPermissionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCosAccessPermissions
// 查看对象存储访问权限列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeCosAccessPermissions(request *DescribeCosAccessPermissionsRequest) (response *DescribeCosAccessPermissionsResponse, err error) {
    return c.DescribeCosAccessPermissionsWithContext(context.Background(), request)
}

// DescribeCosAccessPermissions
// 查看对象存储访问权限列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeCosAccessPermissionsWithContext(ctx context.Context, request *DescribeCosAccessPermissionsRequest) (response *DescribeCosAccessPermissionsResponse, err error) {
    if request == nil {
        request = NewDescribeCosAccessPermissionsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeCosAccessPermissions")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCosAccessPermissions require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCosAccessPermissionsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCosActionListRequest() (request *DescribeCosActionListRequest) {
    request = &DescribeCosActionListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeCosActionList")
    
    
    return
}

func NewDescribeCosActionListResponse() (response *DescribeCosActionListResponse) {
    response = &DescribeCosActionListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCosActionList
// 查看COS接口列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeCosActionList(request *DescribeCosActionListRequest) (response *DescribeCosActionListResponse, err error) {
    return c.DescribeCosActionListWithContext(context.Background(), request)
}

// DescribeCosActionList
// 查看COS接口列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeCosActionListWithContext(ctx context.Context, request *DescribeCosActionListRequest) (response *DescribeCosActionListResponse, err error) {
    if request == nil {
        request = NewDescribeCosActionListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeCosActionList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCosActionList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCosActionListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCosAkAssetRequest() (request *DescribeCosAkAssetRequest) {
    request = &DescribeCosAkAssetRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeCosAkAsset")
    
    
    return
}

func NewDescribeCosAkAssetResponse() (response *DescribeCosAkAssetResponse) {
    response = &DescribeCosAkAssetResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCosAkAsset
// 查看ak资产列表信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeCosAkAsset(request *DescribeCosAkAssetRequest) (response *DescribeCosAkAssetResponse, err error) {
    return c.DescribeCosAkAssetWithContext(context.Background(), request)
}

// DescribeCosAkAsset
// 查看ak资产列表信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeCosAkAssetWithContext(ctx context.Context, request *DescribeCosAkAssetRequest) (response *DescribeCosAkAssetResponse, err error) {
    if request == nil {
        request = NewDescribeCosAkAssetRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeCosAkAsset")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCosAkAsset require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCosAkAssetResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCosAkInvokeIpListRequest() (request *DescribeCosAkInvokeIpListRequest) {
    request = &DescribeCosAkInvokeIpListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeCosAkInvokeIpList")
    
    
    return
}

func NewDescribeCosAkInvokeIpListResponse() (response *DescribeCosAkInvokeIpListResponse) {
    response = &DescribeCosAkInvokeIpListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCosAkInvokeIpList
// 查看存储桶调用源ip列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
func (c *Client) DescribeCosAkInvokeIpList(request *DescribeCosAkInvokeIpListRequest) (response *DescribeCosAkInvokeIpListResponse, err error) {
    return c.DescribeCosAkInvokeIpListWithContext(context.Background(), request)
}

// DescribeCosAkInvokeIpList
// 查看存储桶调用源ip列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
func (c *Client) DescribeCosAkInvokeIpListWithContext(ctx context.Context, request *DescribeCosAkInvokeIpListRequest) (response *DescribeCosAkInvokeIpListResponse, err error) {
    if request == nil {
        request = NewDescribeCosAkInvokeIpListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeCosAkInvokeIpList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCosAkInvokeIpList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCosAkInvokeIpListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCosAlarmListRequest() (request *DescribeCosAlarmListRequest) {
    request = &DescribeCosAlarmListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeCosAlarmList")
    
    
    return
}

func NewDescribeCosAlarmListResponse() (response *DescribeCosAlarmListResponse) {
    response = &DescribeCosAlarmListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCosAlarmList
// 查看告警列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeCosAlarmList(request *DescribeCosAlarmListRequest) (response *DescribeCosAlarmListResponse, err error) {
    return c.DescribeCosAlarmListWithContext(context.Background(), request)
}

// DescribeCosAlarmList
// 查看告警列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeCosAlarmListWithContext(ctx context.Context, request *DescribeCosAlarmListRequest) (response *DescribeCosAlarmListResponse, err error) {
    if request == nil {
        request = NewDescribeCosAlarmListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeCosAlarmList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCosAlarmList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCosAlarmListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCosAlarmTrendDataRequest() (request *DescribeCosAlarmTrendDataRequest) {
    request = &DescribeCosAlarmTrendDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeCosAlarmTrendData")
    
    
    return
}

func NewDescribeCosAlarmTrendDataResponse() (response *DescribeCosAlarmTrendDataResponse) {
    response = &DescribeCosAlarmTrendDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCosAlarmTrendData
// 每日告警新增数据
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCEINUSE = "ResourceInUse"
func (c *Client) DescribeCosAlarmTrendData(request *DescribeCosAlarmTrendDataRequest) (response *DescribeCosAlarmTrendDataResponse, err error) {
    return c.DescribeCosAlarmTrendDataWithContext(context.Background(), request)
}

// DescribeCosAlarmTrendData
// 每日告警新增数据
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCEINUSE = "ResourceInUse"
func (c *Client) DescribeCosAlarmTrendDataWithContext(ctx context.Context, request *DescribeCosAlarmTrendDataRequest) (response *DescribeCosAlarmTrendDataResponse, err error) {
    if request == nil {
        request = NewDescribeCosAlarmTrendDataRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeCosAlarmTrendData")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCosAlarmTrendData require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCosAlarmTrendDataResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCosAssetRequest() (request *DescribeCosAssetRequest) {
    request = &DescribeCosAssetRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeCosAsset")
    
    
    return
}

func NewDescribeCosAssetResponse() (response *DescribeCosAssetResponse) {
    response = &DescribeCosAssetResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCosAsset
// 查看cos资产列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeCosAsset(request *DescribeCosAssetRequest) (response *DescribeCosAssetResponse, err error) {
    return c.DescribeCosAssetWithContext(context.Background(), request)
}

// DescribeCosAsset
// 查看cos资产列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeCosAssetWithContext(ctx context.Context, request *DescribeCosAssetRequest) (response *DescribeCosAssetResponse, err error) {
    if request == nil {
        request = NewDescribeCosAssetRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeCosAsset")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCosAsset require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCosAssetResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCosAssetSyncTaskRequest() (request *DescribeCosAssetSyncTaskRequest) {
    request = &DescribeCosAssetSyncTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeCosAssetSyncTask")
    
    
    return
}

func NewDescribeCosAssetSyncTaskResponse() (response *DescribeCosAssetSyncTaskResponse) {
    response = &DescribeCosAssetSyncTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCosAssetSyncTask
// 获取对应appid对应的当前正在扫描的taskid
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeCosAssetSyncTask(request *DescribeCosAssetSyncTaskRequest) (response *DescribeCosAssetSyncTaskResponse, err error) {
    return c.DescribeCosAssetSyncTaskWithContext(context.Background(), request)
}

// DescribeCosAssetSyncTask
// 获取对应appid对应的当前正在扫描的taskid
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeCosAssetSyncTaskWithContext(ctx context.Context, request *DescribeCosAssetSyncTaskRequest) (response *DescribeCosAssetSyncTaskResponse, err error) {
    if request == nil {
        request = NewDescribeCosAssetSyncTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeCosAssetSyncTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCosAssetSyncTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCosAssetSyncTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCosAuditAppIdListRequest() (request *DescribeCosAuditAppIdListRequest) {
    request = &DescribeCosAuditAppIdListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeCosAuditAppIdList")
    
    
    return
}

func NewDescribeCosAuditAppIdListResponse() (response *DescribeCosAuditAppIdListResponse) {
    response = &DescribeCosAuditAppIdListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCosAuditAppIdList
// 查看该appid下已购买的appid集合
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeCosAuditAppIdList(request *DescribeCosAuditAppIdListRequest) (response *DescribeCosAuditAppIdListResponse, err error) {
    return c.DescribeCosAuditAppIdListWithContext(context.Background(), request)
}

// DescribeCosAuditAppIdList
// 查看该appid下已购买的appid集合
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeCosAuditAppIdListWithContext(ctx context.Context, request *DescribeCosAuditAppIdListRequest) (response *DescribeCosAuditAppIdListResponse, err error) {
    if request == nil {
        request = NewDescribeCosAuditAppIdListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeCosAuditAppIdList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCosAuditAppIdList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCosAuditAppIdListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCosAuditDictionaryListRequest() (request *DescribeCosAuditDictionaryListRequest) {
    request = &DescribeCosAuditDictionaryListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeCosAuditDictionaryList")
    
    
    return
}

func NewDescribeCosAuditDictionaryListResponse() (response *DescribeCosAuditDictionaryListResponse) {
    response = &DescribeCosAuditDictionaryListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCosAuditDictionaryList
// 查询cos审计字典信息列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_RESOURCEID = "InvalidParameter.ResourceId"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeCosAuditDictionaryList(request *DescribeCosAuditDictionaryListRequest) (response *DescribeCosAuditDictionaryListResponse, err error) {
    return c.DescribeCosAuditDictionaryListWithContext(context.Background(), request)
}

// DescribeCosAuditDictionaryList
// 查询cos审计字典信息列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_RESOURCEID = "InvalidParameter.ResourceId"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeCosAuditDictionaryListWithContext(ctx context.Context, request *DescribeCosAuditDictionaryListRequest) (response *DescribeCosAuditDictionaryListResponse, err error) {
    if request == nil {
        request = NewDescribeCosAuditDictionaryListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeCosAuditDictionaryList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCosAuditDictionaryList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCosAuditDictionaryListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCosAuditPayInfoRequest() (request *DescribeCosAuditPayInfoRequest) {
    request = &DescribeCosAuditPayInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeCosAuditPayInfo")
    
    
    return
}

func NewDescribeCosAuditPayInfoResponse() (response *DescribeCosAuditPayInfoResponse) {
    response = &DescribeCosAuditPayInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCosAuditPayInfo
// 获取审计支付信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER_RESOURCEID = "InvalidParameter.ResourceId"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeCosAuditPayInfo(request *DescribeCosAuditPayInfoRequest) (response *DescribeCosAuditPayInfoResponse, err error) {
    return c.DescribeCosAuditPayInfoWithContext(context.Background(), request)
}

// DescribeCosAuditPayInfo
// 获取审计支付信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER_RESOURCEID = "InvalidParameter.ResourceId"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeCosAuditPayInfoWithContext(ctx context.Context, request *DescribeCosAuditPayInfoRequest) (response *DescribeCosAuditPayInfoResponse, err error) {
    if request == nil {
        request = NewDescribeCosAuditPayInfoRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeCosAuditPayInfo")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCosAuditPayInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCosAuditPayInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCosBucketBillingInfoRequest() (request *DescribeCosBucketBillingInfoRequest) {
    request = &DescribeCosBucketBillingInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeCosBucketBillingInfo")
    
    
    return
}

func NewDescribeCosBucketBillingInfoResponse() (response *DescribeCosBucketBillingInfoResponse) {
    response = &DescribeCosBucketBillingInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCosBucketBillingInfo
// 获取存储桶计费信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeCosBucketBillingInfo(request *DescribeCosBucketBillingInfoRequest) (response *DescribeCosBucketBillingInfoResponse, err error) {
    return c.DescribeCosBucketBillingInfoWithContext(context.Background(), request)
}

// DescribeCosBucketBillingInfo
// 获取存储桶计费信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeCosBucketBillingInfoWithContext(ctx context.Context, request *DescribeCosBucketBillingInfoRequest) (response *DescribeCosBucketBillingInfoResponse, err error) {
    if request == nil {
        request = NewDescribeCosBucketBillingInfoRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeCosBucketBillingInfo")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCosBucketBillingInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCosBucketBillingInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCosBucketListRequest() (request *DescribeCosBucketListRequest) {
    request = &DescribeCosBucketListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeCosBucketList")
    
    
    return
}

func NewDescribeCosBucketListResponse() (response *DescribeCosBucketListResponse) {
    response = &DescribeCosBucketListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCosBucketList
// 获取存储桶信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeCosBucketList(request *DescribeCosBucketListRequest) (response *DescribeCosBucketListResponse, err error) {
    return c.DescribeCosBucketListWithContext(context.Background(), request)
}

// DescribeCosBucketList
// 获取存储桶信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeCosBucketListWithContext(ctx context.Context, request *DescribeCosBucketListRequest) (response *DescribeCosBucketListResponse, err error) {
    if request == nil {
        request = NewDescribeCosBucketListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeCosBucketList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCosBucketList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCosBucketListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCosBucketRiskRequest() (request *DescribeCosBucketRiskRequest) {
    request = &DescribeCosBucketRiskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeCosBucketRisk")
    
    
    return
}

func NewDescribeCosBucketRiskResponse() (response *DescribeCosBucketRiskResponse) {
    response = &DescribeCosBucketRiskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCosBucketRisk
// 查看风险资产视角
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeCosBucketRisk(request *DescribeCosBucketRiskRequest) (response *DescribeCosBucketRiskResponse, err error) {
    return c.DescribeCosBucketRiskWithContext(context.Background(), request)
}

// DescribeCosBucketRisk
// 查看风险资产视角
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeCosBucketRiskWithContext(ctx context.Context, request *DescribeCosBucketRiskRequest) (response *DescribeCosBucketRiskResponse, err error) {
    if request == nil {
        request = NewDescribeCosBucketRiskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeCosBucketRisk")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCosBucketRisk require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCosBucketRiskResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCosIdentifyFileListRequest() (request *DescribeCosIdentifyFileListRequest) {
    request = &DescribeCosIdentifyFileListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeCosIdentifyFileList")
    
    
    return
}

func NewDescribeCosIdentifyFileListResponse() (response *DescribeCosIdentifyFileListResponse) {
    response = &DescribeCosIdentifyFileListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCosIdentifyFileList
// 查询cos文件数据识别结果列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_RESOURCEID = "InvalidParameter.ResourceId"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeCosIdentifyFileList(request *DescribeCosIdentifyFileListRequest) (response *DescribeCosIdentifyFileListResponse, err error) {
    return c.DescribeCosIdentifyFileListWithContext(context.Background(), request)
}

// DescribeCosIdentifyFileList
// 查询cos文件数据识别结果列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_RESOURCEID = "InvalidParameter.ResourceId"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeCosIdentifyFileListWithContext(ctx context.Context, request *DescribeCosIdentifyFileListRequest) (response *DescribeCosIdentifyFileListResponse, err error) {
    if request == nil {
        request = NewDescribeCosIdentifyFileListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeCosIdentifyFileList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCosIdentifyFileList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCosIdentifyFileListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCosInvokeUaRequest() (request *DescribeCosInvokeUaRequest) {
    request = &DescribeCosInvokeUaRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeCosInvokeUa")
    
    
    return
}

func NewDescribeCosInvokeUaResponse() (response *DescribeCosInvokeUaResponse) {
    response = &DescribeCosInvokeUaResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCosInvokeUa
// 查看调用记录关联的文件信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_SQLQUERYFAILED = "InvalidParameterValue.SQLQueryFailed"
//  LIMITEXCEEDED = "LimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeCosInvokeUa(request *DescribeCosInvokeUaRequest) (response *DescribeCosInvokeUaResponse, err error) {
    return c.DescribeCosInvokeUaWithContext(context.Background(), request)
}

// DescribeCosInvokeUa
// 查看调用记录关联的文件信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_SQLQUERYFAILED = "InvalidParameterValue.SQLQueryFailed"
//  LIMITEXCEEDED = "LimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeCosInvokeUaWithContext(ctx context.Context, request *DescribeCosInvokeUaRequest) (response *DescribeCosInvokeUaResponse, err error) {
    if request == nil {
        request = NewDescribeCosInvokeUaRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeCosInvokeUa")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCosInvokeUa require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCosInvokeUaResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCosIpInvokeLogRequest() (request *DescribeCosIpInvokeLogRequest) {
    request = &DescribeCosIpInvokeLogRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeCosIpInvokeLog")
    
    
    return
}

func NewDescribeCosIpInvokeLogResponse() (response *DescribeCosIpInvokeLogResponse) {
    response = &DescribeCosIpInvokeLogResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCosIpInvokeLog
// 查看cos调用日志
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeCosIpInvokeLog(request *DescribeCosIpInvokeLogRequest) (response *DescribeCosIpInvokeLogResponse, err error) {
    return c.DescribeCosIpInvokeLogWithContext(context.Background(), request)
}

// DescribeCosIpInvokeLog
// 查看cos调用日志
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeCosIpInvokeLogWithContext(ctx context.Context, request *DescribeCosIpInvokeLogRequest) (response *DescribeCosIpInvokeLogResponse, err error) {
    if request == nil {
        request = NewDescribeCosIpInvokeLogRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeCosIpInvokeLog")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCosIpInvokeLog require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCosIpInvokeLogResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCosIpInvokeRecordFileRequest() (request *DescribeCosIpInvokeRecordFileRequest) {
    request = &DescribeCosIpInvokeRecordFileRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeCosIpInvokeRecordFile")
    
    
    return
}

func NewDescribeCosIpInvokeRecordFileResponse() (response *DescribeCosIpInvokeRecordFileResponse) {
    response = &DescribeCosIpInvokeRecordFileResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCosIpInvokeRecordFile
// 查看调用记录关联的文件信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_SQLQUERYFAILED = "InvalidParameterValue.SQLQueryFailed"
//  LIMITEXCEEDED = "LimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeCosIpInvokeRecordFile(request *DescribeCosIpInvokeRecordFileRequest) (response *DescribeCosIpInvokeRecordFileResponse, err error) {
    return c.DescribeCosIpInvokeRecordFileWithContext(context.Background(), request)
}

// DescribeCosIpInvokeRecordFile
// 查看调用记录关联的文件信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_SQLQUERYFAILED = "InvalidParameterValue.SQLQueryFailed"
//  LIMITEXCEEDED = "LimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeCosIpInvokeRecordFileWithContext(ctx context.Context, request *DescribeCosIpInvokeRecordFileRequest) (response *DescribeCosIpInvokeRecordFileResponse, err error) {
    if request == nil {
        request = NewDescribeCosIpInvokeRecordFileRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeCosIpInvokeRecordFile")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCosIpInvokeRecordFile require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCosIpInvokeRecordFileResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCosOverviewRequest() (request *DescribeCosOverviewRequest) {
    request = &DescribeCosOverviewRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeCosOverview")
    
    
    return
}

func NewDescribeCosOverviewResponse() (response *DescribeCosOverviewResponse) {
    response = &DescribeCosOverviewResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCosOverview
// cos概览信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeCosOverview(request *DescribeCosOverviewRequest) (response *DescribeCosOverviewResponse, err error) {
    return c.DescribeCosOverviewWithContext(context.Background(), request)
}

// DescribeCosOverview
// cos概览信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeCosOverviewWithContext(ctx context.Context, request *DescribeCosOverviewRequest) (response *DescribeCosOverviewResponse, err error) {
    if request == nil {
        request = NewDescribeCosOverviewRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeCosOverview")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCosOverview require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCosOverviewResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCosPolicyRequest() (request *DescribeCosPolicyRequest) {
    request = &DescribeCosPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeCosPolicy")
    
    
    return
}

func NewDescribeCosPolicyResponse() (response *DescribeCosPolicyResponse) {
    response = &DescribeCosPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCosPolicy
// 获取策略列表信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeCosPolicy(request *DescribeCosPolicyRequest) (response *DescribeCosPolicyResponse, err error) {
    return c.DescribeCosPolicyWithContext(context.Background(), request)
}

// DescribeCosPolicy
// 获取策略列表信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeCosPolicyWithContext(ctx context.Context, request *DescribeCosPolicyRequest) (response *DescribeCosPolicyResponse, err error) {
    if request == nil {
        request = NewDescribeCosPolicyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeCosPolicy")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCosPolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCosPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCosRiskActionListRequest() (request *DescribeCosRiskActionListRequest) {
    request = &DescribeCosRiskActionListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeCosRiskActionList")
    
    
    return
}

func NewDescribeCosRiskActionListResponse() (response *DescribeCosRiskActionListResponse) {
    response = &DescribeCosRiskActionListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCosRiskActionList
// 风险接口列表信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeCosRiskActionList(request *DescribeCosRiskActionListRequest) (response *DescribeCosRiskActionListResponse, err error) {
    return c.DescribeCosRiskActionListWithContext(context.Background(), request)
}

// DescribeCosRiskActionList
// 风险接口列表信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeCosRiskActionListWithContext(ctx context.Context, request *DescribeCosRiskActionListRequest) (response *DescribeCosRiskActionListResponse, err error) {
    if request == nil {
        request = NewDescribeCosRiskActionListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeCosRiskActionList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCosRiskActionList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCosRiskActionListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCosRiskEvidenceRequest() (request *DescribeCosRiskEvidenceRequest) {
    request = &DescribeCosRiskEvidenceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeCosRiskEvidence")
    
    
    return
}

func NewDescribeCosRiskEvidenceResponse() (response *DescribeCosRiskEvidenceResponse) {
    response = &DescribeCosRiskEvidenceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCosRiskEvidence
// 查看风险证据以及描述
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeCosRiskEvidence(request *DescribeCosRiskEvidenceRequest) (response *DescribeCosRiskEvidenceResponse, err error) {
    return c.DescribeCosRiskEvidenceWithContext(context.Background(), request)
}

// DescribeCosRiskEvidence
// 查看风险证据以及描述
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeCosRiskEvidenceWithContext(ctx context.Context, request *DescribeCosRiskEvidenceRequest) (response *DescribeCosRiskEvidenceResponse, err error) {
    if request == nil {
        request = NewDescribeCosRiskEvidenceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeCosRiskEvidence")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCosRiskEvidence require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCosRiskEvidenceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCosRiskScanTaskRequest() (request *DescribeCosRiskScanTaskRequest) {
    request = &DescribeCosRiskScanTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeCosRiskScanTask")
    
    
    return
}

func NewDescribeCosRiskScanTaskResponse() (response *DescribeCosRiskScanTaskResponse) {
    response = &DescribeCosRiskScanTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCosRiskScanTask
// 查看存储桶扫描任务详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeCosRiskScanTask(request *DescribeCosRiskScanTaskRequest) (response *DescribeCosRiskScanTaskResponse, err error) {
    return c.DescribeCosRiskScanTaskWithContext(context.Background(), request)
}

// DescribeCosRiskScanTask
// 查看存储桶扫描任务详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeCosRiskScanTaskWithContext(ctx context.Context, request *DescribeCosRiskScanTaskRequest) (response *DescribeCosRiskScanTaskResponse, err error) {
    if request == nil {
        request = NewDescribeCosRiskScanTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeCosRiskScanTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCosRiskScanTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCosRiskScanTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCosRoleAccessPermissionRequest() (request *DescribeCosRoleAccessPermissionRequest) {
    request = &DescribeCosRoleAccessPermissionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeCosRoleAccessPermission")
    
    
    return
}

func NewDescribeCosRoleAccessPermissionResponse() (response *DescribeCosRoleAccessPermissionResponse) {
    response = &DescribeCosRoleAccessPermissionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCosRoleAccessPermission
// 查看cos桶访问权限信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DECRYPTERROR = "InvalidParameter.DecryptError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeCosRoleAccessPermission(request *DescribeCosRoleAccessPermissionRequest) (response *DescribeCosRoleAccessPermissionResponse, err error) {
    return c.DescribeCosRoleAccessPermissionWithContext(context.Background(), request)
}

// DescribeCosRoleAccessPermission
// 查看cos桶访问权限信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DECRYPTERROR = "InvalidParameter.DecryptError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeCosRoleAccessPermissionWithContext(ctx context.Context, request *DescribeCosRoleAccessPermissionRequest) (response *DescribeCosRoleAccessPermissionResponse, err error) {
    if request == nil {
        request = NewDescribeCosRoleAccessPermissionRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeCosRoleAccessPermission")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCosRoleAccessPermission require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCosRoleAccessPermissionResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCosRoleAccessPermissionsRequest() (request *DescribeCosRoleAccessPermissionsRequest) {
    request = &DescribeCosRoleAccessPermissionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeCosRoleAccessPermissions")
    
    
    return
}

func NewDescribeCosRoleAccessPermissionsResponse() (response *DescribeCosRoleAccessPermissionsResponse) {
    response = &DescribeCosRoleAccessPermissionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCosRoleAccessPermissions
// 获取存储桶角色权限列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeCosRoleAccessPermissions(request *DescribeCosRoleAccessPermissionsRequest) (response *DescribeCosRoleAccessPermissionsResponse, err error) {
    return c.DescribeCosRoleAccessPermissionsWithContext(context.Background(), request)
}

// DescribeCosRoleAccessPermissions
// 获取存储桶角色权限列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeCosRoleAccessPermissionsWithContext(ctx context.Context, request *DescribeCosRoleAccessPermissionsRequest) (response *DescribeCosRoleAccessPermissionsResponse, err error) {
    if request == nil {
        request = NewDescribeCosRoleAccessPermissionsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeCosRoleAccessPermissions")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCosRoleAccessPermissions require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCosRoleAccessPermissionsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCosSourceIpRequest() (request *DescribeCosSourceIpRequest) {
    request = &DescribeCosSourceIpRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeCosSourceIp")
    
    
    return
}

func NewDescribeCosSourceIpResponse() (response *DescribeCosSourceIpResponse) {
    response = &DescribeCosSourceIpResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCosSourceIp
// 调用源ip列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeCosSourceIp(request *DescribeCosSourceIpRequest) (response *DescribeCosSourceIpResponse, err error) {
    return c.DescribeCosSourceIpWithContext(context.Background(), request)
}

// DescribeCosSourceIp
// 调用源ip列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeCosSourceIpWithContext(ctx context.Context, request *DescribeCosSourceIpRequest) (response *DescribeCosSourceIpResponse, err error) {
    if request == nil {
        request = NewDescribeCosSourceIpRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeCosSourceIp")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCosSourceIp require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCosSourceIpResponse()
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
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDbAssetInfo(request *DescribeDbAssetInfoRequest) (response *DescribeDbAssetInfoResponse, err error) {
    return c.DescribeDbAssetInfoWithContext(context.Background(), request)
}

// DescribeDbAssetInfo
// db资产详情
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDbAssetInfoWithContext(ctx context.Context, request *DescribeDbAssetInfoRequest) (response *DescribeDbAssetInfoResponse, err error) {
    if request == nil {
        request = NewDescribeDbAssetInfoRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeDbAssetInfo")
    
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
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDbAssets(request *DescribeDbAssetsRequest) (response *DescribeDbAssetsResponse, err error) {
    return c.DescribeDbAssetsWithContext(context.Background(), request)
}

// DescribeDbAssets
// 数据库资产列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDbAssetsWithContext(ctx context.Context, request *DescribeDbAssetsRequest) (response *DescribeDbAssetsResponse, err error) {
    if request == nil {
        request = NewDescribeDbAssetsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeDbAssets")
    
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
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDomainAssets(request *DescribeDomainAssetsRequest) (response *DescribeDomainAssetsResponse, err error) {
    return c.DescribeDomainAssetsWithContext(context.Background(), request)
}

// DescribeDomainAssets
// 域名列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDomainAssetsWithContext(ctx context.Context, request *DescribeDomainAssetsRequest) (response *DescribeDomainAssetsResponse, err error) {
    if request == nil {
        request = NewDescribeDomainAssetsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeDomainAssets")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDomainAssets require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDomainAssetsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDspmAccessRecordRequest() (request *DescribeDspmAccessRecordRequest) {
    request = &DescribeDspmAccessRecordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeDspmAccessRecord")
    
    
    return
}

func NewDescribeDspmAccessRecordResponse() (response *DescribeDspmAccessRecordResponse) {
    response = &DescribeDspmAccessRecordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDspmAccessRecord
// 查询Dspm访问记录
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeDspmAccessRecord(request *DescribeDspmAccessRecordRequest) (response *DescribeDspmAccessRecordResponse, err error) {
    return c.DescribeDspmAccessRecordWithContext(context.Background(), request)
}

// DescribeDspmAccessRecord
// 查询Dspm访问记录
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeDspmAccessRecordWithContext(ctx context.Context, request *DescribeDspmAccessRecordRequest) (response *DescribeDspmAccessRecordResponse, err error) {
    if request == nil {
        request = NewDescribeDspmAccessRecordRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeDspmAccessRecord")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDspmAccessRecord require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDspmAccessRecordResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDspmAccessTopologyAccountsRequest() (request *DescribeDspmAccessTopologyAccountsRequest) {
    request = &DescribeDspmAccessTopologyAccountsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeDspmAccessTopologyAccounts")
    
    
    return
}

func NewDescribeDspmAccessTopologyAccountsResponse() (response *DescribeDspmAccessTopologyAccountsResponse) {
    response = &DescribeDspmAccessTopologyAccountsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDspmAccessTopologyAccounts
// 查询Dspm访问拓扑账号列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeDspmAccessTopologyAccounts(request *DescribeDspmAccessTopologyAccountsRequest) (response *DescribeDspmAccessTopologyAccountsResponse, err error) {
    return c.DescribeDspmAccessTopologyAccountsWithContext(context.Background(), request)
}

// DescribeDspmAccessTopologyAccounts
// 查询Dspm访问拓扑账号列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeDspmAccessTopologyAccountsWithContext(ctx context.Context, request *DescribeDspmAccessTopologyAccountsRequest) (response *DescribeDspmAccessTopologyAccountsResponse, err error) {
    if request == nil {
        request = NewDescribeDspmAccessTopologyAccountsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeDspmAccessTopologyAccounts")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDspmAccessTopologyAccounts require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDspmAccessTopologyAccountsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDspmAccessTopologyAssetsRequest() (request *DescribeDspmAccessTopologyAssetsRequest) {
    request = &DescribeDspmAccessTopologyAssetsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeDspmAccessTopologyAssets")
    
    
    return
}

func NewDescribeDspmAccessTopologyAssetsResponse() (response *DescribeDspmAccessTopologyAssetsResponse) {
    response = &DescribeDspmAccessTopologyAssetsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDspmAccessTopologyAssets
// 查询Dspm访问拓扑资产列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeDspmAccessTopologyAssets(request *DescribeDspmAccessTopologyAssetsRequest) (response *DescribeDspmAccessTopologyAssetsResponse, err error) {
    return c.DescribeDspmAccessTopologyAssetsWithContext(context.Background(), request)
}

// DescribeDspmAccessTopologyAssets
// 查询Dspm访问拓扑资产列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeDspmAccessTopologyAssetsWithContext(ctx context.Context, request *DescribeDspmAccessTopologyAssetsRequest) (response *DescribeDspmAccessTopologyAssetsResponse, err error) {
    if request == nil {
        request = NewDescribeDspmAccessTopologyAssetsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeDspmAccessTopologyAssets")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDspmAccessTopologyAssets require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDspmAccessTopologyAssetsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDspmAccessTopologyIpsRequest() (request *DescribeDspmAccessTopologyIpsRequest) {
    request = &DescribeDspmAccessTopologyIpsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeDspmAccessTopologyIps")
    
    
    return
}

func NewDescribeDspmAccessTopologyIpsResponse() (response *DescribeDspmAccessTopologyIpsResponse) {
    response = &DescribeDspmAccessTopologyIpsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDspmAccessTopologyIps
// 查询Dspm访问拓扑ip列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeDspmAccessTopologyIps(request *DescribeDspmAccessTopologyIpsRequest) (response *DescribeDspmAccessTopologyIpsResponse, err error) {
    return c.DescribeDspmAccessTopologyIpsWithContext(context.Background(), request)
}

// DescribeDspmAccessTopologyIps
// 查询Dspm访问拓扑ip列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeDspmAccessTopologyIpsWithContext(ctx context.Context, request *DescribeDspmAccessTopologyIpsRequest) (response *DescribeDspmAccessTopologyIpsResponse, err error) {
    if request == nil {
        request = NewDescribeDspmAccessTopologyIpsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeDspmAccessTopologyIps")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDspmAccessTopologyIps require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDspmAccessTopologyIpsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDspmApplyHistoryRequest() (request *DescribeDspmApplyHistoryRequest) {
    request = &DescribeDspmApplyHistoryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeDspmApplyHistory")
    
    
    return
}

func NewDescribeDspmApplyHistoryResponse() (response *DescribeDspmApplyHistoryResponse) {
    response = &DescribeDspmApplyHistoryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDspmApplyHistory
// 查询Dspm申请历史
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDspmApplyHistory(request *DescribeDspmApplyHistoryRequest) (response *DescribeDspmApplyHistoryResponse, err error) {
    return c.DescribeDspmApplyHistoryWithContext(context.Background(), request)
}

// DescribeDspmApplyHistory
// 查询Dspm申请历史
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDspmApplyHistoryWithContext(ctx context.Context, request *DescribeDspmApplyHistoryRequest) (response *DescribeDspmApplyHistoryResponse, err error) {
    if request == nil {
        request = NewDescribeDspmApplyHistoryRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeDspmApplyHistory")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDspmApplyHistory require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDspmApplyHistoryResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDspmApplyOrderListRequest() (request *DescribeDspmApplyOrderListRequest) {
    request = &DescribeDspmApplyOrderListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeDspmApplyOrderList")
    
    
    return
}

func NewDescribeDspmApplyOrderListResponse() (response *DescribeDspmApplyOrderListResponse) {
    response = &DescribeDspmApplyOrderListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDspmApplyOrderList
// 查询Dspm申请单列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDspmApplyOrderList(request *DescribeDspmApplyOrderListRequest) (response *DescribeDspmApplyOrderListResponse, err error) {
    return c.DescribeDspmApplyOrderListWithContext(context.Background(), request)
}

// DescribeDspmApplyOrderList
// 查询Dspm申请单列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDspmApplyOrderListWithContext(ctx context.Context, request *DescribeDspmApplyOrderListRequest) (response *DescribeDspmApplyOrderListResponse, err error) {
    if request == nil {
        request = NewDescribeDspmApplyOrderListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeDspmApplyOrderList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDspmApplyOrderList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDspmApplyOrderListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDspmApproveHistoryRequest() (request *DescribeDspmApproveHistoryRequest) {
    request = &DescribeDspmApproveHistoryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeDspmApproveHistory")
    
    
    return
}

func NewDescribeDspmApproveHistoryResponse() (response *DescribeDspmApproveHistoryResponse) {
    response = &DescribeDspmApproveHistoryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDspmApproveHistory
// 查询Dspm审批历史
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDspmApproveHistory(request *DescribeDspmApproveHistoryRequest) (response *DescribeDspmApproveHistoryResponse, err error) {
    return c.DescribeDspmApproveHistoryWithContext(context.Background(), request)
}

// DescribeDspmApproveHistory
// 查询Dspm审批历史
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDspmApproveHistoryWithContext(ctx context.Context, request *DescribeDspmApproveHistoryRequest) (response *DescribeDspmApproveHistoryResponse, err error) {
    if request == nil {
        request = NewDescribeDspmApproveHistoryRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeDspmApproveHistory")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDspmApproveHistory require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDspmApproveHistoryResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDspmApproveOrderListRequest() (request *DescribeDspmApproveOrderListRequest) {
    request = &DescribeDspmApproveOrderListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeDspmApproveOrderList")
    
    
    return
}

func NewDescribeDspmApproveOrderListResponse() (response *DescribeDspmApproveOrderListResponse) {
    response = &DescribeDspmApproveOrderListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDspmApproveOrderList
// 查询Dspm审批单列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDspmApproveOrderList(request *DescribeDspmApproveOrderListRequest) (response *DescribeDspmApproveOrderListResponse, err error) {
    return c.DescribeDspmApproveOrderListWithContext(context.Background(), request)
}

// DescribeDspmApproveOrderList
// 查询Dspm审批单列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDspmApproveOrderListWithContext(ctx context.Context, request *DescribeDspmApproveOrderListRequest) (response *DescribeDspmApproveOrderListResponse, err error) {
    if request == nil {
        request = NewDescribeDspmApproveOrderListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeDspmApproveOrderList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDspmApproveOrderList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDspmApproveOrderListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDspmAssetAccessTopologyRequest() (request *DescribeDspmAssetAccessTopologyRequest) {
    request = &DescribeDspmAssetAccessTopologyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeDspmAssetAccessTopology")
    
    
    return
}

func NewDescribeDspmAssetAccessTopologyResponse() (response *DescribeDspmAssetAccessTopologyResponse) {
    response = &DescribeDspmAssetAccessTopologyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDspmAssetAccessTopology
// 查询Dspm资产访问拓扑
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeDspmAssetAccessTopology(request *DescribeDspmAssetAccessTopologyRequest) (response *DescribeDspmAssetAccessTopologyResponse, err error) {
    return c.DescribeDspmAssetAccessTopologyWithContext(context.Background(), request)
}

// DescribeDspmAssetAccessTopology
// 查询Dspm资产访问拓扑
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeDspmAssetAccessTopologyWithContext(ctx context.Context, request *DescribeDspmAssetAccessTopologyRequest) (response *DescribeDspmAssetAccessTopologyResponse, err error) {
    if request == nil {
        request = NewDescribeDspmAssetAccessTopologyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeDspmAssetAccessTopology")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDspmAssetAccessTopology require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDspmAssetAccessTopologyResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDspmAssetAccountIdentifyRequest() (request *DescribeDspmAssetAccountIdentifyRequest) {
    request = &DescribeDspmAssetAccountIdentifyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeDspmAssetAccountIdentify")
    
    
    return
}

func NewDescribeDspmAssetAccountIdentifyResponse() (response *DescribeDspmAssetAccountIdentifyResponse) {
    response = &DescribeDspmAssetAccountIdentifyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDspmAssetAccountIdentify
// 查询Dspm资产账号身份信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDspmAssetAccountIdentify(request *DescribeDspmAssetAccountIdentifyRequest) (response *DescribeDspmAssetAccountIdentifyResponse, err error) {
    return c.DescribeDspmAssetAccountIdentifyWithContext(context.Background(), request)
}

// DescribeDspmAssetAccountIdentify
// 查询Dspm资产账号身份信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDspmAssetAccountIdentifyWithContext(ctx context.Context, request *DescribeDspmAssetAccountIdentifyRequest) (response *DescribeDspmAssetAccountIdentifyResponse, err error) {
    if request == nil {
        request = NewDescribeDspmAssetAccountIdentifyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeDspmAssetAccountIdentify")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDspmAssetAccountIdentify require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDspmAssetAccountIdentifyResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDspmAssetAccountPresetPrivilegesRequest() (request *DescribeDspmAssetAccountPresetPrivilegesRequest) {
    request = &DescribeDspmAssetAccountPresetPrivilegesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeDspmAssetAccountPresetPrivileges")
    
    
    return
}

func NewDescribeDspmAssetAccountPresetPrivilegesResponse() (response *DescribeDspmAssetAccountPresetPrivilegesResponse) {
    response = &DescribeDspmAssetAccountPresetPrivilegesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDspmAssetAccountPresetPrivileges
// 查询Dspm资产账号预设特权信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDspmAssetAccountPresetPrivileges(request *DescribeDspmAssetAccountPresetPrivilegesRequest) (response *DescribeDspmAssetAccountPresetPrivilegesResponse, err error) {
    return c.DescribeDspmAssetAccountPresetPrivilegesWithContext(context.Background(), request)
}

// DescribeDspmAssetAccountPresetPrivileges
// 查询Dspm资产账号预设特权信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDspmAssetAccountPresetPrivilegesWithContext(ctx context.Context, request *DescribeDspmAssetAccountPresetPrivilegesRequest) (response *DescribeDspmAssetAccountPresetPrivilegesResponse, err error) {
    if request == nil {
        request = NewDescribeDspmAssetAccountPresetPrivilegesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeDspmAssetAccountPresetPrivileges")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDspmAssetAccountPresetPrivileges require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDspmAssetAccountPresetPrivilegesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDspmAssetAccountRecycledPrivilegesRequest() (request *DescribeDspmAssetAccountRecycledPrivilegesRequest) {
    request = &DescribeDspmAssetAccountRecycledPrivilegesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeDspmAssetAccountRecycledPrivileges")
    
    
    return
}

func NewDescribeDspmAssetAccountRecycledPrivilegesResponse() (response *DescribeDspmAssetAccountRecycledPrivilegesResponse) {
    response = &DescribeDspmAssetAccountRecycledPrivilegesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDspmAssetAccountRecycledPrivileges
// 查询Dspm资产账号回收后特权信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDspmAssetAccountRecycledPrivileges(request *DescribeDspmAssetAccountRecycledPrivilegesRequest) (response *DescribeDspmAssetAccountRecycledPrivilegesResponse, err error) {
    return c.DescribeDspmAssetAccountRecycledPrivilegesWithContext(context.Background(), request)
}

// DescribeDspmAssetAccountRecycledPrivileges
// 查询Dspm资产账号回收后特权信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDspmAssetAccountRecycledPrivilegesWithContext(ctx context.Context, request *DescribeDspmAssetAccountRecycledPrivilegesRequest) (response *DescribeDspmAssetAccountRecycledPrivilegesResponse, err error) {
    if request == nil {
        request = NewDescribeDspmAssetAccountRecycledPrivilegesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeDspmAssetAccountRecycledPrivileges")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDspmAssetAccountRecycledPrivileges require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDspmAssetAccountRecycledPrivilegesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDspmAssetAccountsRequest() (request *DescribeDspmAssetAccountsRequest) {
    request = &DescribeDspmAssetAccountsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeDspmAssetAccounts")
    
    
    return
}

func NewDescribeDspmAssetAccountsResponse() (response *DescribeDspmAssetAccountsResponse) {
    response = &DescribeDspmAssetAccountsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDspmAssetAccounts
// 查询Dspm资产账号列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDspmAssetAccounts(request *DescribeDspmAssetAccountsRequest) (response *DescribeDspmAssetAccountsResponse, err error) {
    return c.DescribeDspmAssetAccountsWithContext(context.Background(), request)
}

// DescribeDspmAssetAccounts
// 查询Dspm资产账号列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDspmAssetAccountsWithContext(ctx context.Context, request *DescribeDspmAssetAccountsRequest) (response *DescribeDspmAssetAccountsResponse, err error) {
    if request == nil {
        request = NewDescribeDspmAssetAccountsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeDspmAssetAccounts")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDspmAssetAccounts require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDspmAssetAccountsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDspmAssetDatabaseListRequest() (request *DescribeDspmAssetDatabaseListRequest) {
    request = &DescribeDspmAssetDatabaseListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeDspmAssetDatabaseList")
    
    
    return
}

func NewDescribeDspmAssetDatabaseListResponse() (response *DescribeDspmAssetDatabaseListResponse) {
    response = &DescribeDspmAssetDatabaseListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDspmAssetDatabaseList
// 查询资产数据库信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_RESOURCEID = "InvalidParameter.ResourceId"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeDspmAssetDatabaseList(request *DescribeDspmAssetDatabaseListRequest) (response *DescribeDspmAssetDatabaseListResponse, err error) {
    return c.DescribeDspmAssetDatabaseListWithContext(context.Background(), request)
}

// DescribeDspmAssetDatabaseList
// 查询资产数据库信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_RESOURCEID = "InvalidParameter.ResourceId"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeDspmAssetDatabaseListWithContext(ctx context.Context, request *DescribeDspmAssetDatabaseListRequest) (response *DescribeDspmAssetDatabaseListResponse, err error) {
    if request == nil {
        request = NewDescribeDspmAssetDatabaseListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeDspmAssetDatabaseList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDspmAssetDatabaseList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDspmAssetDatabaseListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDspmAssetDatabasesRequest() (request *DescribeDspmAssetDatabasesRequest) {
    request = &DescribeDspmAssetDatabasesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeDspmAssetDatabases")
    
    
    return
}

func NewDescribeDspmAssetDatabasesResponse() (response *DescribeDspmAssetDatabasesResponse) {
    response = &DescribeDspmAssetDatabasesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDspmAssetDatabases
// 查询Dspm资产数据库列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDspmAssetDatabases(request *DescribeDspmAssetDatabasesRequest) (response *DescribeDspmAssetDatabasesResponse, err error) {
    return c.DescribeDspmAssetDatabasesWithContext(context.Background(), request)
}

// DescribeDspmAssetDatabases
// 查询Dspm资产数据库列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDspmAssetDatabasesWithContext(ctx context.Context, request *DescribeDspmAssetDatabasesRequest) (response *DescribeDspmAssetDatabasesResponse, err error) {
    if request == nil {
        request = NewDescribeDspmAssetDatabasesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeDspmAssetDatabases")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDspmAssetDatabases require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDspmAssetDatabasesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDspmAssetFieldListRequest() (request *DescribeDspmAssetFieldListRequest) {
    request = &DescribeDspmAssetFieldListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeDspmAssetFieldList")
    
    
    return
}

func NewDescribeDspmAssetFieldListResponse() (response *DescribeDspmAssetFieldListResponse) {
    response = &DescribeDspmAssetFieldListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDspmAssetFieldList
// 查询dspm资产字段信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_RESOURCEID = "InvalidParameter.ResourceId"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeDspmAssetFieldList(request *DescribeDspmAssetFieldListRequest) (response *DescribeDspmAssetFieldListResponse, err error) {
    return c.DescribeDspmAssetFieldListWithContext(context.Background(), request)
}

// DescribeDspmAssetFieldList
// 查询dspm资产字段信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_RESOURCEID = "InvalidParameter.ResourceId"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeDspmAssetFieldListWithContext(ctx context.Context, request *DescribeDspmAssetFieldListRequest) (response *DescribeDspmAssetFieldListResponse, err error) {
    if request == nil {
        request = NewDescribeDspmAssetFieldListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeDspmAssetFieldList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDspmAssetFieldList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDspmAssetFieldListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDspmAssetIdsRequest() (request *DescribeDspmAssetIdsRequest) {
    request = &DescribeDspmAssetIdsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeDspmAssetIds")
    
    
    return
}

func NewDescribeDspmAssetIdsResponse() (response *DescribeDspmAssetIdsResponse) {
    response = &DescribeDspmAssetIdsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDspmAssetIds
// 查询Dspm资产id列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDspmAssetIds(request *DescribeDspmAssetIdsRequest) (response *DescribeDspmAssetIdsResponse, err error) {
    return c.DescribeDspmAssetIdsWithContext(context.Background(), request)
}

// DescribeDspmAssetIds
// 查询Dspm资产id列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDspmAssetIdsWithContext(ctx context.Context, request *DescribeDspmAssetIdsRequest) (response *DescribeDspmAssetIdsResponse, err error) {
    if request == nil {
        request = NewDescribeDspmAssetIdsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeDspmAssetIds")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDspmAssetIds require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDspmAssetIdsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDspmAssetLoginCredentialRequest() (request *DescribeDspmAssetLoginCredentialRequest) {
    request = &DescribeDspmAssetLoginCredentialRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeDspmAssetLoginCredential")
    
    
    return
}

func NewDescribeDspmAssetLoginCredentialResponse() (response *DescribeDspmAssetLoginCredentialResponse) {
    response = &DescribeDspmAssetLoginCredentialResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDspmAssetLoginCredential
// 查询Dspm资产登录凭据
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDspmAssetLoginCredential(request *DescribeDspmAssetLoginCredentialRequest) (response *DescribeDspmAssetLoginCredentialResponse, err error) {
    return c.DescribeDspmAssetLoginCredentialWithContext(context.Background(), request)
}

// DescribeDspmAssetLoginCredential
// 查询Dspm资产登录凭据
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDspmAssetLoginCredentialWithContext(ctx context.Context, request *DescribeDspmAssetLoginCredentialRequest) (response *DescribeDspmAssetLoginCredentialResponse, err error) {
    if request == nil {
        request = NewDescribeDspmAssetLoginCredentialRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeDspmAssetLoginCredential")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDspmAssetLoginCredential require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDspmAssetLoginCredentialResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDspmAssetSecurityAnalyseStatusRequest() (request *DescribeDspmAssetSecurityAnalyseStatusRequest) {
    request = &DescribeDspmAssetSecurityAnalyseStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeDspmAssetSecurityAnalyseStatus")
    
    
    return
}

func NewDescribeDspmAssetSecurityAnalyseStatusResponse() (response *DescribeDspmAssetSecurityAnalyseStatusResponse) {
    response = &DescribeDspmAssetSecurityAnalyseStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDspmAssetSecurityAnalyseStatus
// 查询Dspm资产安全分析状态
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDspmAssetSecurityAnalyseStatus(request *DescribeDspmAssetSecurityAnalyseStatusRequest) (response *DescribeDspmAssetSecurityAnalyseStatusResponse, err error) {
    return c.DescribeDspmAssetSecurityAnalyseStatusWithContext(context.Background(), request)
}

// DescribeDspmAssetSecurityAnalyseStatus
// 查询Dspm资产安全分析状态
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDspmAssetSecurityAnalyseStatusWithContext(ctx context.Context, request *DescribeDspmAssetSecurityAnalyseStatusRequest) (response *DescribeDspmAssetSecurityAnalyseStatusResponse, err error) {
    if request == nil {
        request = NewDescribeDspmAssetSecurityAnalyseStatusRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeDspmAssetSecurityAnalyseStatus")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDspmAssetSecurityAnalyseStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDspmAssetSecurityAnalyseStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDspmAssetSupportedPrivilegesRequest() (request *DescribeDspmAssetSupportedPrivilegesRequest) {
    request = &DescribeDspmAssetSupportedPrivilegesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeDspmAssetSupportedPrivileges")
    
    
    return
}

func NewDescribeDspmAssetSupportedPrivilegesResponse() (response *DescribeDspmAssetSupportedPrivilegesResponse) {
    response = &DescribeDspmAssetSupportedPrivilegesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDspmAssetSupportedPrivileges
// 查询Dspm资产支持的权限
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDspmAssetSupportedPrivileges(request *DescribeDspmAssetSupportedPrivilegesRequest) (response *DescribeDspmAssetSupportedPrivilegesResponse, err error) {
    return c.DescribeDspmAssetSupportedPrivilegesWithContext(context.Background(), request)
}

// DescribeDspmAssetSupportedPrivileges
// 查询Dspm资产支持的权限
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDspmAssetSupportedPrivilegesWithContext(ctx context.Context, request *DescribeDspmAssetSupportedPrivilegesRequest) (response *DescribeDspmAssetSupportedPrivilegesResponse, err error) {
    if request == nil {
        request = NewDescribeDspmAssetSupportedPrivilegesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeDspmAssetSupportedPrivileges")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDspmAssetSupportedPrivileges require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDspmAssetSupportedPrivilegesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDspmAssetTableListRequest() (request *DescribeDspmAssetTableListRequest) {
    request = &DescribeDspmAssetTableListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeDspmAssetTableList")
    
    
    return
}

func NewDescribeDspmAssetTableListResponse() (response *DescribeDspmAssetTableListResponse) {
    response = &DescribeDspmAssetTableListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDspmAssetTableList
// 查询资产表信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_RESOURCEID = "InvalidParameter.ResourceId"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeDspmAssetTableList(request *DescribeDspmAssetTableListRequest) (response *DescribeDspmAssetTableListResponse, err error) {
    return c.DescribeDspmAssetTableListWithContext(context.Background(), request)
}

// DescribeDspmAssetTableList
// 查询资产表信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_RESOURCEID = "InvalidParameter.ResourceId"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeDspmAssetTableListWithContext(ctx context.Context, request *DescribeDspmAssetTableListRequest) (response *DescribeDspmAssetTableListResponse, err error) {
    if request == nil {
        request = NewDescribeDspmAssetTableListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeDspmAssetTableList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDspmAssetTableList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDspmAssetTableListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDspmAssetsRequest() (request *DescribeDspmAssetsRequest) {
    request = &DescribeDspmAssetsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeDspmAssets")
    
    
    return
}

func NewDescribeDspmAssetsResponse() (response *DescribeDspmAssetsResponse) {
    response = &DescribeDspmAssetsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDspmAssets
// 查询Dspm资产列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDspmAssets(request *DescribeDspmAssetsRequest) (response *DescribeDspmAssetsResponse, err error) {
    return c.DescribeDspmAssetsWithContext(context.Background(), request)
}

// DescribeDspmAssets
// 查询Dspm资产列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDspmAssetsWithContext(ctx context.Context, request *DescribeDspmAssetsRequest) (response *DescribeDspmAssetsResponse, err error) {
    if request == nil {
        request = NewDescribeDspmAssetsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeDspmAssets")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDspmAssets require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDspmAssetsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDspmBackupLogListRequest() (request *DescribeDspmBackupLogListRequest) {
    request = &DescribeDspmBackupLogListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeDspmBackupLogList")
    
    
    return
}

func NewDescribeDspmBackupLogListResponse() (response *DescribeDspmBackupLogListResponse) {
    response = &DescribeDspmBackupLogListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDspmBackupLogList
// 查询备份日志列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDspmBackupLogList(request *DescribeDspmBackupLogListRequest) (response *DescribeDspmBackupLogListResponse, err error) {
    return c.DescribeDspmBackupLogListWithContext(context.Background(), request)
}

// DescribeDspmBackupLogList
// 查询备份日志列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDspmBackupLogListWithContext(ctx context.Context, request *DescribeDspmBackupLogListRequest) (response *DescribeDspmBackupLogListResponse, err error) {
    if request == nil {
        request = NewDescribeDspmBackupLogListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeDspmBackupLogList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDspmBackupLogList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDspmBackupLogListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDspmBackupSettingRequest() (request *DescribeDspmBackupSettingRequest) {
    request = &DescribeDspmBackupSettingRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeDspmBackupSetting")
    
    
    return
}

func NewDescribeDspmBackupSettingResponse() (response *DescribeDspmBackupSettingResponse) {
    response = &DescribeDspmBackupSettingResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDspmBackupSetting
// 查询日志备份配置
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDspmBackupSetting(request *DescribeDspmBackupSettingRequest) (response *DescribeDspmBackupSettingResponse, err error) {
    return c.DescribeDspmBackupSettingWithContext(context.Background(), request)
}

// DescribeDspmBackupSetting
// 查询日志备份配置
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDspmBackupSettingWithContext(ctx context.Context, request *DescribeDspmBackupSettingRequest) (response *DescribeDspmBackupSettingResponse, err error) {
    if request == nil {
        request = NewDescribeDspmBackupSettingRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeDspmBackupSetting")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDspmBackupSetting require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDspmBackupSettingResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDspmDictionaryListRequest() (request *DescribeDspmDictionaryListRequest) {
    request = &DescribeDspmDictionaryListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeDspmDictionaryList")
    
    
    return
}

func NewDescribeDspmDictionaryListResponse() (response *DescribeDspmDictionaryListResponse) {
    response = &DescribeDspmDictionaryListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDspmDictionaryList
// 查询dspm字典信息列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_RESOURCEID = "InvalidParameter.ResourceId"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeDspmDictionaryList(request *DescribeDspmDictionaryListRequest) (response *DescribeDspmDictionaryListResponse, err error) {
    return c.DescribeDspmDictionaryListWithContext(context.Background(), request)
}

// DescribeDspmDictionaryList
// 查询dspm字典信息列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_RESOURCEID = "InvalidParameter.ResourceId"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeDspmDictionaryListWithContext(ctx context.Context, request *DescribeDspmDictionaryListRequest) (response *DescribeDspmDictionaryListResponse, err error) {
    if request == nil {
        request = NewDescribeDspmDictionaryListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeDspmDictionaryList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDspmDictionaryList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDspmDictionaryListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDspmExportTaskRequest() (request *DescribeDspmExportTaskRequest) {
    request = &DescribeDspmExportTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeDspmExportTask")
    
    
    return
}

func NewDescribeDspmExportTaskResponse() (response *DescribeDspmExportTaskResponse) {
    response = &DescribeDspmExportTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDspmExportTask
// 查询导出任务
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_RESOURCEID = "InvalidParameter.ResourceId"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeDspmExportTask(request *DescribeDspmExportTaskRequest) (response *DescribeDspmExportTaskResponse, err error) {
    return c.DescribeDspmExportTaskWithContext(context.Background(), request)
}

// DescribeDspmExportTask
// 查询导出任务
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_RESOURCEID = "InvalidParameter.ResourceId"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeDspmExportTaskWithContext(ctx context.Context, request *DescribeDspmExportTaskRequest) (response *DescribeDspmExportTaskResponse, err error) {
    if request == nil {
        request = NewDescribeDspmExportTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeDspmExportTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDspmExportTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDspmExportTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDspmIdentifyIdListRequest() (request *DescribeDspmIdentifyIdListRequest) {
    request = &DescribeDspmIdentifyIdListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeDspmIdentifyIdList")
    
    
    return
}

func NewDescribeDspmIdentifyIdListResponse() (response *DescribeDspmIdentifyIdListResponse) {
    response = &DescribeDspmIdentifyIdListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDspmIdentifyIdList
// 查询Dspm身份id列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDspmIdentifyIdList(request *DescribeDspmIdentifyIdListRequest) (response *DescribeDspmIdentifyIdListResponse, err error) {
    return c.DescribeDspmIdentifyIdListWithContext(context.Background(), request)
}

// DescribeDspmIdentifyIdList
// 查询Dspm身份id列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDspmIdentifyIdListWithContext(ctx context.Context, request *DescribeDspmIdentifyIdListRequest) (response *DescribeDspmIdentifyIdListResponse, err error) {
    if request == nil {
        request = NewDescribeDspmIdentifyIdListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeDspmIdentifyIdList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDspmIdentifyIdList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDspmIdentifyIdListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDspmIdentifyInfoRequest() (request *DescribeDspmIdentifyInfoRequest) {
    request = &DescribeDspmIdentifyInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeDspmIdentifyInfo")
    
    
    return
}

func NewDescribeDspmIdentifyInfoResponse() (response *DescribeDspmIdentifyInfoResponse) {
    response = &DescribeDspmIdentifyInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDspmIdentifyInfo
// 查询Dspm身份信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeDspmIdentifyInfo(request *DescribeDspmIdentifyInfoRequest) (response *DescribeDspmIdentifyInfoResponse, err error) {
    return c.DescribeDspmIdentifyInfoWithContext(context.Background(), request)
}

// DescribeDspmIdentifyInfo
// 查询Dspm身份信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeDspmIdentifyInfoWithContext(ctx context.Context, request *DescribeDspmIdentifyInfoRequest) (response *DescribeDspmIdentifyInfoResponse, err error) {
    if request == nil {
        request = NewDescribeDspmIdentifyInfoRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeDspmIdentifyInfo")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDspmIdentifyInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDspmIdentifyInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDspmIdentifyInfoListRequest() (request *DescribeDspmIdentifyInfoListRequest) {
    request = &DescribeDspmIdentifyInfoListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeDspmIdentifyInfoList")
    
    
    return
}

func NewDescribeDspmIdentifyInfoListResponse() (response *DescribeDspmIdentifyInfoListResponse) {
    response = &DescribeDspmIdentifyInfoListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDspmIdentifyInfoList
// 查询Dspm身份信息列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDspmIdentifyInfoList(request *DescribeDspmIdentifyInfoListRequest) (response *DescribeDspmIdentifyInfoListResponse, err error) {
    return c.DescribeDspmIdentifyInfoListWithContext(context.Background(), request)
}

// DescribeDspmIdentifyInfoList
// 查询Dspm身份信息列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDspmIdentifyInfoListWithContext(ctx context.Context, request *DescribeDspmIdentifyInfoListRequest) (response *DescribeDspmIdentifyInfoListResponse, err error) {
    if request == nil {
        request = NewDescribeDspmIdentifyInfoListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeDspmIdentifyInfoList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDspmIdentifyInfoList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDspmIdentifyInfoListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDspmLogListRequest() (request *DescribeDspmLogListRequest) {
    request = &DescribeDspmLogListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeDspmLogList")
    
    
    return
}

func NewDescribeDspmLogListResponse() (response *DescribeDspmLogListResponse) {
    response = &DescribeDspmLogListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDspmLogList
// 查询日志列表信息
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
func (c *Client) DescribeDspmLogList(request *DescribeDspmLogListRequest) (response *DescribeDspmLogListResponse, err error) {
    return c.DescribeDspmLogListWithContext(context.Background(), request)
}

// DescribeDspmLogList
// 查询日志列表信息
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
func (c *Client) DescribeDspmLogListWithContext(ctx context.Context, request *DescribeDspmLogListRequest) (response *DescribeDspmLogListResponse, err error) {
    if request == nil {
        request = NewDescribeDspmLogListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeDspmLogList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDspmLogList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDspmLogListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDspmPayInfoRequest() (request *DescribeDspmPayInfoRequest) {
    request = &DescribeDspmPayInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeDspmPayInfo")
    
    
    return
}

func NewDescribeDspmPayInfoResponse() (response *DescribeDspmPayInfoResponse) {
    response = &DescribeDspmPayInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDspmPayInfo
// 获取已购Dspm订单信息
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
func (c *Client) DescribeDspmPayInfo(request *DescribeDspmPayInfoRequest) (response *DescribeDspmPayInfoResponse, err error) {
    return c.DescribeDspmPayInfoWithContext(context.Background(), request)
}

// DescribeDspmPayInfo
// 获取已购Dspm订单信息
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
func (c *Client) DescribeDspmPayInfoWithContext(ctx context.Context, request *DescribeDspmPayInfoRequest) (response *DescribeDspmPayInfoResponse, err error) {
    if request == nil {
        request = NewDescribeDspmPayInfoRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeDspmPayInfo")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDspmPayInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDspmPayInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDspmPersonApplyHistoryRequest() (request *DescribeDspmPersonApplyHistoryRequest) {
    request = &DescribeDspmPersonApplyHistoryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeDspmPersonApplyHistory")
    
    
    return
}

func NewDescribeDspmPersonApplyHistoryResponse() (response *DescribeDspmPersonApplyHistoryResponse) {
    response = &DescribeDspmPersonApplyHistoryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDspmPersonApplyHistory
// 查询Dspm访客申请记录
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDspmPersonApplyHistory(request *DescribeDspmPersonApplyHistoryRequest) (response *DescribeDspmPersonApplyHistoryResponse, err error) {
    return c.DescribeDspmPersonApplyHistoryWithContext(context.Background(), request)
}

// DescribeDspmPersonApplyHistory
// 查询Dspm访客申请记录
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDspmPersonApplyHistoryWithContext(ctx context.Context, request *DescribeDspmPersonApplyHistoryRequest) (response *DescribeDspmPersonApplyHistoryResponse, err error) {
    if request == nil {
        request = NewDescribeDspmPersonApplyHistoryRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeDspmPersonApplyHistory")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDspmPersonApplyHistory require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDspmPersonApplyHistoryResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDspmPersonalIdentifyListRequest() (request *DescribeDspmPersonalIdentifyListRequest) {
    request = &DescribeDspmPersonalIdentifyListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeDspmPersonalIdentifyList")
    
    
    return
}

func NewDescribeDspmPersonalIdentifyListResponse() (response *DescribeDspmPersonalIdentifyListResponse) {
    response = &DescribeDspmPersonalIdentifyListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDspmPersonalIdentifyList
// 查询Dspm个人身份信息列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDspmPersonalIdentifyList(request *DescribeDspmPersonalIdentifyListRequest) (response *DescribeDspmPersonalIdentifyListResponse, err error) {
    return c.DescribeDspmPersonalIdentifyListWithContext(context.Background(), request)
}

// DescribeDspmPersonalIdentifyList
// 查询Dspm个人身份信息列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDspmPersonalIdentifyListWithContext(ctx context.Context, request *DescribeDspmPersonalIdentifyListRequest) (response *DescribeDspmPersonalIdentifyListResponse, err error) {
    if request == nil {
        request = NewDescribeDspmPersonalIdentifyListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeDspmPersonalIdentifyList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDspmPersonalIdentifyList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDspmPersonalIdentifyListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDspmRiskRequest() (request *DescribeDspmRiskRequest) {
    request = &DescribeDspmRiskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeDspmRisk")
    
    
    return
}

func NewDescribeDspmRiskResponse() (response *DescribeDspmRiskResponse) {
    response = &DescribeDspmRiskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDspmRisk
// 查询Dspm风险记录
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeDspmRisk(request *DescribeDspmRiskRequest) (response *DescribeDspmRiskResponse, err error) {
    return c.DescribeDspmRiskWithContext(context.Background(), request)
}

// DescribeDspmRisk
// 查询Dspm风险记录
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeDspmRiskWithContext(ctx context.Context, request *DescribeDspmRiskRequest) (response *DescribeDspmRiskResponse, err error) {
    if request == nil {
        request = NewDescribeDspmRiskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeDspmRisk")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDspmRisk require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDspmRiskResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDspmRiskDetailRequest() (request *DescribeDspmRiskDetailRequest) {
    request = &DescribeDspmRiskDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeDspmRiskDetail")
    
    
    return
}

func NewDescribeDspmRiskDetailResponse() (response *DescribeDspmRiskDetailResponse) {
    response = &DescribeDspmRiskDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDspmRiskDetail
// 查询Dspm风险详情
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeDspmRiskDetail(request *DescribeDspmRiskDetailRequest) (response *DescribeDspmRiskDetailResponse, err error) {
    return c.DescribeDspmRiskDetailWithContext(context.Background(), request)
}

// DescribeDspmRiskDetail
// 查询Dspm风险详情
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeDspmRiskDetailWithContext(ctx context.Context, request *DescribeDspmRiskDetailRequest) (response *DescribeDspmRiskDetailResponse, err error) {
    if request == nil {
        request = NewDescribeDspmRiskDetailRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeDspmRiskDetail")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDspmRiskDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDspmRiskDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDspmRiskStrategyRequest() (request *DescribeDspmRiskStrategyRequest) {
    request = &DescribeDspmRiskStrategyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeDspmRiskStrategy")
    
    
    return
}

func NewDescribeDspmRiskStrategyResponse() (response *DescribeDspmRiskStrategyResponse) {
    response = &DescribeDspmRiskStrategyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDspmRiskStrategy
// 查询Dspm风险策略
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeDspmRiskStrategy(request *DescribeDspmRiskStrategyRequest) (response *DescribeDspmRiskStrategyResponse, err error) {
    return c.DescribeDspmRiskStrategyWithContext(context.Background(), request)
}

// DescribeDspmRiskStrategy
// 查询Dspm风险策略
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeDspmRiskStrategyWithContext(ctx context.Context, request *DescribeDspmRiskStrategyRequest) (response *DescribeDspmRiskStrategyResponse, err error) {
    if request == nil {
        request = NewDescribeDspmRiskStrategyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeDspmRiskStrategy")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDspmRiskStrategy require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDspmRiskStrategyResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDspmRiskStrategyGroupRequest() (request *DescribeDspmRiskStrategyGroupRequest) {
    request = &DescribeDspmRiskStrategyGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeDspmRiskStrategyGroup")
    
    
    return
}

func NewDescribeDspmRiskStrategyGroupResponse() (response *DescribeDspmRiskStrategyGroupResponse) {
    response = &DescribeDspmRiskStrategyGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDspmRiskStrategyGroup
// 查询Dspm风险分组策略
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeDspmRiskStrategyGroup(request *DescribeDspmRiskStrategyGroupRequest) (response *DescribeDspmRiskStrategyGroupResponse, err error) {
    return c.DescribeDspmRiskStrategyGroupWithContext(context.Background(), request)
}

// DescribeDspmRiskStrategyGroup
// 查询Dspm风险分组策略
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeDspmRiskStrategyGroupWithContext(ctx context.Context, request *DescribeDspmRiskStrategyGroupRequest) (response *DescribeDspmRiskStrategyGroupResponse, err error) {
    if request == nil {
        request = NewDescribeDspmRiskStrategyGroupRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeDspmRiskStrategyGroup")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDspmRiskStrategyGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDspmRiskStrategyGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDspmRiskTendencyRequest() (request *DescribeDspmRiskTendencyRequest) {
    request = &DescribeDspmRiskTendencyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeDspmRiskTendency")
    
    
    return
}

func NewDescribeDspmRiskTendencyResponse() (response *DescribeDspmRiskTendencyResponse) {
    response = &DescribeDspmRiskTendencyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDspmRiskTendency
// 查询Dspm风险趋势
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeDspmRiskTendency(request *DescribeDspmRiskTendencyRequest) (response *DescribeDspmRiskTendencyResponse, err error) {
    return c.DescribeDspmRiskTendencyWithContext(context.Background(), request)
}

// DescribeDspmRiskTendency
// 查询Dspm风险趋势
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeDspmRiskTendencyWithContext(ctx context.Context, request *DescribeDspmRiskTendencyRequest) (response *DescribeDspmRiskTendencyResponse, err error) {
    if request == nil {
        request = NewDescribeDspmRiskTendencyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeDspmRiskTendency")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDspmRiskTendency require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDspmRiskTendencyResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDspmStatisticsRequest() (request *DescribeDspmStatisticsRequest) {
    request = &DescribeDspmStatisticsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeDspmStatistics")
    
    
    return
}

func NewDescribeDspmStatisticsResponse() (response *DescribeDspmStatisticsResponse) {
    response = &DescribeDspmStatisticsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDspmStatistics
// 查询Dspm统计信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeDspmStatistics(request *DescribeDspmStatisticsRequest) (response *DescribeDspmStatisticsResponse, err error) {
    return c.DescribeDspmStatisticsWithContext(context.Background(), request)
}

// DescribeDspmStatistics
// 查询Dspm统计信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeDspmStatisticsWithContext(ctx context.Context, request *DescribeDspmStatisticsRequest) (response *DescribeDspmStatisticsResponse, err error) {
    if request == nil {
        request = NewDescribeDspmStatisticsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeDspmStatistics")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDspmStatistics require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDspmStatisticsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDspmSupportedAssetTypeRequest() (request *DescribeDspmSupportedAssetTypeRequest) {
    request = &DescribeDspmSupportedAssetTypeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeDspmSupportedAssetType")
    
    
    return
}

func NewDescribeDspmSupportedAssetTypeResponse() (response *DescribeDspmSupportedAssetTypeResponse) {
    response = &DescribeDspmSupportedAssetTypeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDspmSupportedAssetType
// 查询Dspm支持的资产类型信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDspmSupportedAssetType(request *DescribeDspmSupportedAssetTypeRequest) (response *DescribeDspmSupportedAssetTypeResponse, err error) {
    return c.DescribeDspmSupportedAssetTypeWithContext(context.Background(), request)
}

// DescribeDspmSupportedAssetType
// 查询Dspm支持的资产类型信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDspmSupportedAssetTypeWithContext(ctx context.Context, request *DescribeDspmSupportedAssetTypeRequest) (response *DescribeDspmSupportedAssetTypeResponse, err error) {
    if request == nil {
        request = NewDescribeDspmSupportedAssetTypeRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeDspmSupportedAssetType")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDspmSupportedAssetType require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDspmSupportedAssetTypeResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDspmSyncAssetsStatusRequest() (request *DescribeDspmSyncAssetsStatusRequest) {
    request = &DescribeDspmSyncAssetsStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeDspmSyncAssetsStatus")
    
    
    return
}

func NewDescribeDspmSyncAssetsStatusResponse() (response *DescribeDspmSyncAssetsStatusResponse) {
    response = &DescribeDspmSyncAssetsStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDspmSyncAssetsStatus
// 查询Dspm同步资产状态
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeDspmSyncAssetsStatus(request *DescribeDspmSyncAssetsStatusRequest) (response *DescribeDspmSyncAssetsStatusResponse, err error) {
    return c.DescribeDspmSyncAssetsStatusWithContext(context.Background(), request)
}

// DescribeDspmSyncAssetsStatus
// 查询Dspm同步资产状态
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeDspmSyncAssetsStatusWithContext(ctx context.Context, request *DescribeDspmSyncAssetsStatusRequest) (response *DescribeDspmSyncAssetsStatusResponse, err error) {
    if request == nil {
        request = NewDescribeDspmSyncAssetsStatusRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeDspmSyncAssetsStatus")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDspmSyncAssetsStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDspmSyncAssetsStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDspmSyncUsersStatusRequest() (request *DescribeDspmSyncUsersStatusRequest) {
    request = &DescribeDspmSyncUsersStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeDspmSyncUsersStatus")
    
    
    return
}

func NewDescribeDspmSyncUsersStatusResponse() (response *DescribeDspmSyncUsersStatusResponse) {
    response = &DescribeDspmSyncUsersStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDspmSyncUsersStatus
// 查询Dspm同步用户状态
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeDspmSyncUsersStatus(request *DescribeDspmSyncUsersStatusRequest) (response *DescribeDspmSyncUsersStatusResponse, err error) {
    return c.DescribeDspmSyncUsersStatusWithContext(context.Background(), request)
}

// DescribeDspmSyncUsersStatus
// 查询Dspm同步用户状态
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeDspmSyncUsersStatusWithContext(ctx context.Context, request *DescribeDspmSyncUsersStatusRequest) (response *DescribeDspmSyncUsersStatusResponse, err error) {
    if request == nil {
        request = NewDescribeDspmSyncUsersStatusRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeDspmSyncUsersStatus")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDspmSyncUsersStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDspmSyncUsersStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDspmWhitelistStrategyRequest() (request *DescribeDspmWhitelistStrategyRequest) {
    request = &DescribeDspmWhitelistStrategyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeDspmWhitelistStrategy")
    
    
    return
}

func NewDescribeDspmWhitelistStrategyResponse() (response *DescribeDspmWhitelistStrategyResponse) {
    response = &DescribeDspmWhitelistStrategyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDspmWhitelistStrategy
// 查询Dspm白名单策略
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeDspmWhitelistStrategy(request *DescribeDspmWhitelistStrategyRequest) (response *DescribeDspmWhitelistStrategyResponse, err error) {
    return c.DescribeDspmWhitelistStrategyWithContext(context.Background(), request)
}

// DescribeDspmWhitelistStrategy
// 查询Dspm白名单策略
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeDspmWhitelistStrategyWithContext(ctx context.Context, request *DescribeDspmWhitelistStrategyRequest) (response *DescribeDspmWhitelistStrategyResponse, err error) {
    if request == nil {
        request = NewDescribeDspmWhitelistStrategyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeDspmWhitelistStrategy")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDspmWhitelistStrategy require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDspmWhitelistStrategyResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeExposeAssetCategoryRequest() (request *DescribeExposeAssetCategoryRequest) {
    request = &DescribeExposeAssetCategoryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeExposeAssetCategory")
    
    
    return
}

func NewDescribeExposeAssetCategoryResponse() (response *DescribeExposeAssetCategoryResponse) {
    response = &DescribeExposeAssetCategoryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeExposeAssetCategory
// 云边界分析资产分类
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeExposeAssetCategory(request *DescribeExposeAssetCategoryRequest) (response *DescribeExposeAssetCategoryResponse, err error) {
    return c.DescribeExposeAssetCategoryWithContext(context.Background(), request)
}

// DescribeExposeAssetCategory
// 云边界分析资产分类
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeExposeAssetCategoryWithContext(ctx context.Context, request *DescribeExposeAssetCategoryRequest) (response *DescribeExposeAssetCategoryResponse, err error) {
    if request == nil {
        request = NewDescribeExposeAssetCategoryRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeExposeAssetCategory")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeExposeAssetCategory require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeExposeAssetCategoryResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeExposePathRequest() (request *DescribeExposePathRequest) {
    request = &DescribeExposePathRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeExposePath")
    
    
    return
}

func NewDescribeExposePathResponse() (response *DescribeExposePathResponse) {
    response = &DescribeExposePathResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeExposePath
// 查询云边界分析路径节点
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeExposePath(request *DescribeExposePathRequest) (response *DescribeExposePathResponse, err error) {
    return c.DescribeExposePathWithContext(context.Background(), request)
}

// DescribeExposePath
// 查询云边界分析路径节点
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeExposePathWithContext(ctx context.Context, request *DescribeExposePathRequest) (response *DescribeExposePathResponse, err error) {
    if request == nil {
        request = NewDescribeExposePathRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeExposePath")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeExposePath require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeExposePathResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeExposuresRequest() (request *DescribeExposuresRequest) {
    request = &DescribeExposuresRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeExposures")
    
    
    return
}

func NewDescribeExposuresResponse() (response *DescribeExposuresResponse) {
    response = &DescribeExposuresResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeExposures
// 云边界分析资产列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeExposures(request *DescribeExposuresRequest) (response *DescribeExposuresResponse, err error) {
    return c.DescribeExposuresWithContext(context.Background(), request)
}

// DescribeExposures
// 云边界分析资产列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeExposuresWithContext(ctx context.Context, request *DescribeExposuresRequest) (response *DescribeExposuresResponse, err error) {
    if request == nil {
        request = NewDescribeExposuresRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeExposures")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeExposures require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeExposuresResponse()
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
//  FAILEDOPERATION = "FailedOperation"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeGatewayAssets(request *DescribeGatewayAssetsRequest) (response *DescribeGatewayAssetsResponse, err error) {
    return c.DescribeGatewayAssetsWithContext(context.Background(), request)
}

// DescribeGatewayAssets
// 获取网关列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeGatewayAssetsWithContext(ctx context.Context, request *DescribeGatewayAssetsRequest) (response *DescribeGatewayAssetsResponse, err error) {
    if request == nil {
        request = NewDescribeGatewayAssetsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeGatewayAssets")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeGatewayAssets require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeGatewayAssetsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeHighBaseLineRiskListRequest() (request *DescribeHighBaseLineRiskListRequest) {
    request = &DescribeHighBaseLineRiskListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeHighBaseLineRiskList")
    
    
    return
}

func NewDescribeHighBaseLineRiskListResponse() (response *DescribeHighBaseLineRiskListResponse) {
    response = &DescribeHighBaseLineRiskListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeHighBaseLineRiskList
// 查询云边界分析-暴露路径下主机节点的高危基线风险列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeHighBaseLineRiskList(request *DescribeHighBaseLineRiskListRequest) (response *DescribeHighBaseLineRiskListResponse, err error) {
    return c.DescribeHighBaseLineRiskListWithContext(context.Background(), request)
}

// DescribeHighBaseLineRiskList
// 查询云边界分析-暴露路径下主机节点的高危基线风险列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeHighBaseLineRiskListWithContext(ctx context.Context, request *DescribeHighBaseLineRiskListRequest) (response *DescribeHighBaseLineRiskListResponse, err error) {
    if request == nil {
        request = NewDescribeHighBaseLineRiskListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeHighBaseLineRiskList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeHighBaseLineRiskList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeHighBaseLineRiskListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeIaCFileListRequest() (request *DescribeIaCFileListRequest) {
    request = &DescribeIaCFileListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeIaCFileList")
    
    
    return
}

func NewDescribeIaCFileListResponse() (response *DescribeIaCFileListResponse) {
    response = &DescribeIaCFileListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeIaCFileList
// 获取IaC检测文件列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeIaCFileList(request *DescribeIaCFileListRequest) (response *DescribeIaCFileListResponse, err error) {
    return c.DescribeIaCFileListWithContext(context.Background(), request)
}

// DescribeIaCFileList
// 获取IaC检测文件列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeIaCFileListWithContext(ctx context.Context, request *DescribeIaCFileListRequest) (response *DescribeIaCFileListResponse, err error) {
    if request == nil {
        request = NewDescribeIaCFileListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeIaCFileList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeIaCFileList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeIaCFileListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeIaCFileOverviewRequest() (request *DescribeIaCFileOverviewRequest) {
    request = &DescribeIaCFileOverviewRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeIaCFileOverview")
    
    
    return
}

func NewDescribeIaCFileOverviewResponse() (response *DescribeIaCFileOverviewResponse) {
    response = &DescribeIaCFileOverviewResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeIaCFileOverview
// 获取IaC检测文件概览
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeIaCFileOverview(request *DescribeIaCFileOverviewRequest) (response *DescribeIaCFileOverviewResponse, err error) {
    return c.DescribeIaCFileOverviewWithContext(context.Background(), request)
}

// DescribeIaCFileOverview
// 获取IaC检测文件概览
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeIaCFileOverviewWithContext(ctx context.Context, request *DescribeIaCFileOverviewRequest) (response *DescribeIaCFileOverviewResponse, err error) {
    if request == nil {
        request = NewDescribeIaCFileOverviewRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeIaCFileOverview")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeIaCFileOverview require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeIaCFileOverviewResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeIaCFileReportRequest() (request *DescribeIaCFileReportRequest) {
    request = &DescribeIaCFileReportRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeIaCFileReport")
    
    
    return
}

func NewDescribeIaCFileReportResponse() (response *DescribeIaCFileReportResponse) {
    response = &DescribeIaCFileReportResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeIaCFileReport
// 获取IaC检测文件报告
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeIaCFileReport(request *DescribeIaCFileReportRequest) (response *DescribeIaCFileReportResponse, err error) {
    return c.DescribeIaCFileReportWithContext(context.Background(), request)
}

// DescribeIaCFileReport
// 获取IaC检测文件报告
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeIaCFileReportWithContext(ctx context.Context, request *DescribeIaCFileReportRequest) (response *DescribeIaCFileReportResponse, err error) {
    if request == nil {
        request = NewDescribeIaCFileReportRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeIaCFileReport")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeIaCFileReport require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeIaCFileReportResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeIaCTokenListRequest() (request *DescribeIaCTokenListRequest) {
    request = &DescribeIaCTokenListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeIaCTokenList")
    
    
    return
}

func NewDescribeIaCTokenListResponse() (response *DescribeIaCTokenListResponse) {
    response = &DescribeIaCTokenListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeIaCTokenList
// 获取IaC检测接入Token列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeIaCTokenList(request *DescribeIaCTokenListRequest) (response *DescribeIaCTokenListResponse, err error) {
    return c.DescribeIaCTokenListWithContext(context.Background(), request)
}

// DescribeIaCTokenList
// 获取IaC检测接入Token列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeIaCTokenListWithContext(ctx context.Context, request *DescribeIaCTokenListRequest) (response *DescribeIaCTokenListResponse, err error) {
    if request == nil {
        request = NewDescribeIaCTokenListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeIaCTokenList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeIaCTokenList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeIaCTokenListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeIpInvokeRecordRequest() (request *DescribeIpInvokeRecordRequest) {
    request = &DescribeIpInvokeRecordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeIpInvokeRecord")
    
    
    return
}

func NewDescribeIpInvokeRecordResponse() (response *DescribeIpInvokeRecordResponse) {
    response = &DescribeIpInvokeRecordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeIpInvokeRecord
// 对象存储异常检测调用记录信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeIpInvokeRecord(request *DescribeIpInvokeRecordRequest) (response *DescribeIpInvokeRecordResponse, err error) {
    return c.DescribeIpInvokeRecordWithContext(context.Background(), request)
}

// DescribeIpInvokeRecord
// 对象存储异常检测调用记录信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeIpInvokeRecordWithContext(ctx context.Context, request *DescribeIpInvokeRecordRequest) (response *DescribeIpInvokeRecordResponse, err error) {
    if request == nil {
        request = NewDescribeIpInvokeRecordRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeIpInvokeRecord")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeIpInvokeRecord require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeIpInvokeRecordResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeIpInvokeRecordDetailRequest() (request *DescribeIpInvokeRecordDetailRequest) {
    request = &DescribeIpInvokeRecordDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeIpInvokeRecordDetail")
    
    
    return
}

func NewDescribeIpInvokeRecordDetailResponse() (response *DescribeIpInvokeRecordDetailResponse) {
    response = &DescribeIpInvokeRecordDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeIpInvokeRecordDetail
// ip访问列表详情信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeIpInvokeRecordDetail(request *DescribeIpInvokeRecordDetailRequest) (response *DescribeIpInvokeRecordDetailResponse, err error) {
    return c.DescribeIpInvokeRecordDetailWithContext(context.Background(), request)
}

// DescribeIpInvokeRecordDetail
// ip访问列表详情信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeIpInvokeRecordDetailWithContext(ctx context.Context, request *DescribeIpInvokeRecordDetailRequest) (response *DescribeIpInvokeRecordDetailResponse, err error) {
    if request == nil {
        request = NewDescribeIpInvokeRecordDetailRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeIpInvokeRecordDetail")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeIpInvokeRecordDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeIpInvokeRecordDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeKeySandboxCredentialRequest() (request *DescribeKeySandboxCredentialRequest) {
    request = &DescribeKeySandboxCredentialRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeKeySandboxCredential")
    
    
    return
}

func NewDescribeKeySandboxCredentialResponse() (response *DescribeKeySandboxCredentialResponse) {
    response = &DescribeKeySandboxCredentialResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeKeySandboxCredential
// 查询凭证详情，返回凭证元数据和打码后的凭据数据。access类型返回Access数组（Key原文、Value打码），sts类型返回STS对象（System原文、SecretID和SecretKey打码）
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeKeySandboxCredential(request *DescribeKeySandboxCredentialRequest) (response *DescribeKeySandboxCredentialResponse, err error) {
    return c.DescribeKeySandboxCredentialWithContext(context.Background(), request)
}

// DescribeKeySandboxCredential
// 查询凭证详情，返回凭证元数据和打码后的凭据数据。access类型返回Access数组（Key原文、Value打码），sts类型返回STS对象（System原文、SecretID和SecretKey打码）
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeKeySandboxCredentialWithContext(ctx context.Context, request *DescribeKeySandboxCredentialRequest) (response *DescribeKeySandboxCredentialResponse, err error) {
    if request == nil {
        request = NewDescribeKeySandboxCredentialRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeKeySandboxCredential")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeKeySandboxCredential require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeKeySandboxCredentialResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeKeySandboxCredentialListRequest() (request *DescribeKeySandboxCredentialListRequest) {
    request = &DescribeKeySandboxCredentialListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeKeySandboxCredentialList")
    
    
    return
}

func NewDescribeKeySandboxCredentialListResponse() (response *DescribeKeySandboxCredentialListResponse) {
    response = &DescribeKeySandboxCredentialListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeKeySandboxCredentialList
// 查询凭证列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeKeySandboxCredentialList(request *DescribeKeySandboxCredentialListRequest) (response *DescribeKeySandboxCredentialListResponse, err error) {
    return c.DescribeKeySandboxCredentialListWithContext(context.Background(), request)
}

// DescribeKeySandboxCredentialList
// 查询凭证列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeKeySandboxCredentialListWithContext(ctx context.Context, request *DescribeKeySandboxCredentialListRequest) (response *DescribeKeySandboxCredentialListResponse, err error) {
    if request == nil {
        request = NewDescribeKeySandboxCredentialListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeKeySandboxCredentialList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeKeySandboxCredentialList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeKeySandboxCredentialListResponse()
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
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeListenerList(request *DescribeListenerListRequest) (response *DescribeListenerListResponse, err error) {
    return c.DescribeListenerListWithContext(context.Background(), request)
}

// DescribeListenerList
// 查询clb监听器列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeListenerListWithContext(ctx context.Context, request *DescribeListenerListRequest) (response *DescribeListenerListResponse, err error) {
    if request == nil {
        request = NewDescribeListenerListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeListenerList")
    
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
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeNICAssets(request *DescribeNICAssetsRequest) (response *DescribeNICAssetsResponse, err error) {
    return c.DescribeNICAssetsWithContext(context.Background(), request)
}

// DescribeNICAssets
// 获取网卡列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeNICAssetsWithContext(ctx context.Context, request *DescribeNICAssetsRequest) (response *DescribeNICAssetsResponse, err error) {
    if request == nil {
        request = NewDescribeNICAssetsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeNICAssets")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeNICAssets require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeNICAssetsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeOrganizationInfoRequest() (request *DescribeOrganizationInfoRequest) {
    request = &DescribeOrganizationInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeOrganizationInfo")
    
    
    return
}

func NewDescribeOrganizationInfoResponse() (response *DescribeOrganizationInfoResponse) {
    response = &DescribeOrganizationInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeOrganizationInfo
// 查询集团账号详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeOrganizationInfo(request *DescribeOrganizationInfoRequest) (response *DescribeOrganizationInfoResponse, err error) {
    return c.DescribeOrganizationInfoWithContext(context.Background(), request)
}

// DescribeOrganizationInfo
// 查询集团账号详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeOrganizationInfoWithContext(ctx context.Context, request *DescribeOrganizationInfoRequest) (response *DescribeOrganizationInfoResponse, err error) {
    if request == nil {
        request = NewDescribeOrganizationInfoRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeOrganizationInfo")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeOrganizationInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeOrganizationInfoResponse()
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
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeOrganizationUserInfo(request *DescribeOrganizationUserInfoRequest) (response *DescribeOrganizationUserInfoResponse, err error) {
    return c.DescribeOrganizationUserInfoWithContext(context.Background(), request)
}

// DescribeOrganizationUserInfo
// 查询集团账号用户列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeOrganizationUserInfoWithContext(ctx context.Context, request *DescribeOrganizationUserInfoRequest) (response *DescribeOrganizationUserInfoResponse, err error) {
    if request == nil {
        request = NewDescribeOrganizationUserInfoRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeOrganizationUserInfo")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeOrganizationUserInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeOrganizationUserInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeOtherCloudAssetsRequest() (request *DescribeOtherCloudAssetsRequest) {
    request = &DescribeOtherCloudAssetsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeOtherCloudAssets")
    
    
    return
}

func NewDescribeOtherCloudAssetsResponse() (response *DescribeOtherCloudAssetsResponse) {
    response = &DescribeOtherCloudAssetsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeOtherCloudAssets
// 资产列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeOtherCloudAssets(request *DescribeOtherCloudAssetsRequest) (response *DescribeOtherCloudAssetsResponse, err error) {
    return c.DescribeOtherCloudAssetsWithContext(context.Background(), request)
}

// DescribeOtherCloudAssets
// 资产列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeOtherCloudAssetsWithContext(ctx context.Context, request *DescribeOtherCloudAssetsRequest) (response *DescribeOtherCloudAssetsResponse, err error) {
    if request == nil {
        request = NewDescribeOtherCloudAssetsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeOtherCloudAssets")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeOtherCloudAssets require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeOtherCloudAssetsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePolicyHitDataRequest() (request *DescribePolicyHitDataRequest) {
    request = &DescribePolicyHitDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribePolicyHitData")
    
    
    return
}

func NewDescribePolicyHitDataResponse() (response *DescribePolicyHitDataResponse) {
    response = &DescribePolicyHitDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribePolicyHitData
// 按日期查看策略命中详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribePolicyHitData(request *DescribePolicyHitDataRequest) (response *DescribePolicyHitDataResponse, err error) {
    return c.DescribePolicyHitDataWithContext(context.Background(), request)
}

// DescribePolicyHitData
// 按日期查看策略命中详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribePolicyHitDataWithContext(ctx context.Context, request *DescribePolicyHitDataRequest) (response *DescribePolicyHitDataResponse, err error) {
    if request == nil {
        request = NewDescribePolicyHitDataRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribePolicyHitData")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePolicyHitData require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePolicyHitDataResponse()
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
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribePublicIpAssets(request *DescribePublicIpAssetsRequest) (response *DescribePublicIpAssetsResponse, err error) {
    return c.DescribePublicIpAssetsWithContext(context.Background(), request)
}

// DescribePublicIpAssets
// ip公网列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribePublicIpAssetsWithContext(ctx context.Context, request *DescribePublicIpAssetsRequest) (response *DescribePublicIpAssetsResponse, err error) {
    if request == nil {
        request = NewDescribePublicIpAssetsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribePublicIpAssets")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePublicIpAssets require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePublicIpAssetsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRepositoryImageAssetsRequest() (request *DescribeRepositoryImageAssetsRequest) {
    request = &DescribeRepositoryImageAssetsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeRepositoryImageAssets")
    
    
    return
}

func NewDescribeRepositoryImageAssetsResponse() (response *DescribeRepositoryImageAssetsResponse) {
    response = &DescribeRepositoryImageAssetsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRepositoryImageAssets
// 仓库镜像列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeRepositoryImageAssets(request *DescribeRepositoryImageAssetsRequest) (response *DescribeRepositoryImageAssetsResponse, err error) {
    return c.DescribeRepositoryImageAssetsWithContext(context.Background(), request)
}

// DescribeRepositoryImageAssets
// 仓库镜像列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeRepositoryImageAssetsWithContext(ctx context.Context, request *DescribeRepositoryImageAssetsRequest) (response *DescribeRepositoryImageAssetsResponse, err error) {
    if request == nil {
        request = NewDescribeRepositoryImageAssetsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeRepositoryImageAssets")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRepositoryImageAssets require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRepositoryImageAssetsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRiskBucketListRequest() (request *DescribeRiskBucketListRequest) {
    request = &DescribeRiskBucketListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeRiskBucketList")
    
    
    return
}

func NewDescribeRiskBucketListResponse() (response *DescribeRiskBucketListResponse) {
    response = &DescribeRiskBucketListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRiskBucketList
// 查看风险关联的存储桶信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeRiskBucketList(request *DescribeRiskBucketListRequest) (response *DescribeRiskBucketListResponse, err error) {
    return c.DescribeRiskBucketListWithContext(context.Background(), request)
}

// DescribeRiskBucketList
// 查看风险关联的存储桶信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeRiskBucketListWithContext(ctx context.Context, request *DescribeRiskBucketListRequest) (response *DescribeRiskBucketListResponse, err error) {
    if request == nil {
        request = NewDescribeRiskBucketListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeRiskBucketList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRiskBucketList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRiskBucketListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRiskCallRecordRequest() (request *DescribeRiskCallRecordRequest) {
    request = &DescribeRiskCallRecordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeRiskCallRecord")
    
    
    return
}

func NewDescribeRiskCallRecordResponse() (response *DescribeRiskCallRecordResponse) {
    response = &DescribeRiskCallRecordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRiskCallRecord
// 获取风险调用记录列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeRiskCallRecord(request *DescribeRiskCallRecordRequest) (response *DescribeRiskCallRecordResponse, err error) {
    return c.DescribeRiskCallRecordWithContext(context.Background(), request)
}

// DescribeRiskCallRecord
// 获取风险调用记录列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeRiskCallRecordWithContext(ctx context.Context, request *DescribeRiskCallRecordRequest) (response *DescribeRiskCallRecordResponse, err error) {
    if request == nil {
        request = NewDescribeRiskCallRecordRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeRiskCallRecord")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRiskCallRecord require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRiskCallRecordResponse()
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
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeRiskCenterAssetViewCFGRiskList(request *DescribeRiskCenterAssetViewCFGRiskListRequest) (response *DescribeRiskCenterAssetViewCFGRiskListResponse, err error) {
    return c.DescribeRiskCenterAssetViewCFGRiskListWithContext(context.Background(), request)
}

// DescribeRiskCenterAssetViewCFGRiskList
// 获取资产视角的配置风险列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeRiskCenterAssetViewCFGRiskListWithContext(ctx context.Context, request *DescribeRiskCenterAssetViewCFGRiskListRequest) (response *DescribeRiskCenterAssetViewCFGRiskListResponse, err error) {
    if request == nil {
        request = NewDescribeRiskCenterAssetViewCFGRiskListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeRiskCenterAssetViewCFGRiskList")
    
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
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DECRYPTERROR = "InvalidParameter.DecryptError"
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
//  UNAUTHORIZEDOPERATION_CLOUDAUDIT = "UnauthorizedOperation.CloudAudit"
//  UNAUTHORIZEDOPERATION_COS = "UnauthorizedOperation.Cos"
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
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DECRYPTERROR = "InvalidParameter.DecryptError"
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
//  UNAUTHORIZEDOPERATION_CLOUDAUDIT = "UnauthorizedOperation.CloudAudit"
//  UNAUTHORIZEDOPERATION_COS = "UnauthorizedOperation.Cos"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeRiskCenterAssetViewPortRiskListWithContext(ctx context.Context, request *DescribeRiskCenterAssetViewPortRiskListRequest) (response *DescribeRiskCenterAssetViewPortRiskListResponse, err error) {
    if request == nil {
        request = NewDescribeRiskCenterAssetViewPortRiskListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeRiskCenterAssetViewPortRiskList")
    
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
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DECRYPTERROR = "InvalidParameter.DecryptError"
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
//  UNAUTHORIZEDOPERATION_CLOUDAUDIT = "UnauthorizedOperation.CloudAudit"
//  UNAUTHORIZEDOPERATION_COS = "UnauthorizedOperation.Cos"
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
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DECRYPTERROR = "InvalidParameter.DecryptError"
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
//  UNAUTHORIZEDOPERATION_CLOUDAUDIT = "UnauthorizedOperation.CloudAudit"
//  UNAUTHORIZEDOPERATION_COS = "UnauthorizedOperation.Cos"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeRiskCenterAssetViewVULRiskListWithContext(ctx context.Context, request *DescribeRiskCenterAssetViewVULRiskListRequest) (response *DescribeRiskCenterAssetViewVULRiskListResponse, err error) {
    if request == nil {
        request = NewDescribeRiskCenterAssetViewVULRiskListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeRiskCenterAssetViewVULRiskList")
    
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
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DECRYPTERROR = "InvalidParameter.DecryptError"
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
//  UNAUTHORIZEDOPERATION_CLOUDAUDIT = "UnauthorizedOperation.CloudAudit"
//  UNAUTHORIZEDOPERATION_COS = "UnauthorizedOperation.Cos"
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
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DECRYPTERROR = "InvalidParameter.DecryptError"
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
//  UNAUTHORIZEDOPERATION_CLOUDAUDIT = "UnauthorizedOperation.CloudAudit"
//  UNAUTHORIZEDOPERATION_COS = "UnauthorizedOperation.Cos"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeRiskCenterAssetViewWeakPasswordRiskListWithContext(ctx context.Context, request *DescribeRiskCenterAssetViewWeakPasswordRiskListRequest) (response *DescribeRiskCenterAssetViewWeakPasswordRiskListResponse, err error) {
    if request == nil {
        request = NewDescribeRiskCenterAssetViewWeakPasswordRiskListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeRiskCenterAssetViewWeakPasswordRiskList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRiskCenterAssetViewWeakPasswordRiskList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRiskCenterAssetViewWeakPasswordRiskListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRiskCenterCFGViewCFGRiskListRequest() (request *DescribeRiskCenterCFGViewCFGRiskListRequest) {
    request = &DescribeRiskCenterCFGViewCFGRiskListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeRiskCenterCFGViewCFGRiskList")
    
    
    return
}

func NewDescribeRiskCenterCFGViewCFGRiskListResponse() (response *DescribeRiskCenterCFGViewCFGRiskListResponse) {
    response = &DescribeRiskCenterCFGViewCFGRiskListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRiskCenterCFGViewCFGRiskList
// 获取配置视角的配置风险列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DECRYPTERROR = "InvalidParameter.DecryptError"
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
//  UNAUTHORIZEDOPERATION_CLOUDAUDIT = "UnauthorizedOperation.CloudAudit"
//  UNAUTHORIZEDOPERATION_COS = "UnauthorizedOperation.Cos"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeRiskCenterCFGViewCFGRiskList(request *DescribeRiskCenterCFGViewCFGRiskListRequest) (response *DescribeRiskCenterCFGViewCFGRiskListResponse, err error) {
    return c.DescribeRiskCenterCFGViewCFGRiskListWithContext(context.Background(), request)
}

// DescribeRiskCenterCFGViewCFGRiskList
// 获取配置视角的配置风险列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DECRYPTERROR = "InvalidParameter.DecryptError"
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
//  UNAUTHORIZEDOPERATION_CLOUDAUDIT = "UnauthorizedOperation.CloudAudit"
//  UNAUTHORIZEDOPERATION_COS = "UnauthorizedOperation.Cos"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeRiskCenterCFGViewCFGRiskListWithContext(ctx context.Context, request *DescribeRiskCenterCFGViewCFGRiskListRequest) (response *DescribeRiskCenterCFGViewCFGRiskListResponse, err error) {
    if request == nil {
        request = NewDescribeRiskCenterCFGViewCFGRiskListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeRiskCenterCFGViewCFGRiskList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRiskCenterCFGViewCFGRiskList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRiskCenterCFGViewCFGRiskListResponse()
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
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DECRYPTERROR = "InvalidParameter.DecryptError"
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
//  UNAUTHORIZEDOPERATION_CLOUDAUDIT = "UnauthorizedOperation.CloudAudit"
//  UNAUTHORIZEDOPERATION_COS = "UnauthorizedOperation.Cos"
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
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DECRYPTERROR = "InvalidParameter.DecryptError"
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
//  UNAUTHORIZEDOPERATION_CLOUDAUDIT = "UnauthorizedOperation.CloudAudit"
//  UNAUTHORIZEDOPERATION_COS = "UnauthorizedOperation.Cos"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeRiskCenterPortViewPortRiskListWithContext(ctx context.Context, request *DescribeRiskCenterPortViewPortRiskListRequest) (response *DescribeRiskCenterPortViewPortRiskListResponse, err error) {
    if request == nil {
        request = NewDescribeRiskCenterPortViewPortRiskListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeRiskCenterPortViewPortRiskList")
    
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
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DECRYPTERROR = "InvalidParameter.DecryptError"
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
//  UNAUTHORIZEDOPERATION_CLOUDAUDIT = "UnauthorizedOperation.CloudAudit"
//  UNAUTHORIZEDOPERATION_COS = "UnauthorizedOperation.Cos"
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
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DECRYPTERROR = "InvalidParameter.DecryptError"
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
//  UNAUTHORIZEDOPERATION_CLOUDAUDIT = "UnauthorizedOperation.CloudAudit"
//  UNAUTHORIZEDOPERATION_COS = "UnauthorizedOperation.Cos"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeRiskCenterServerRiskListWithContext(ctx context.Context, request *DescribeRiskCenterServerRiskListRequest) (response *DescribeRiskCenterServerRiskListResponse, err error) {
    if request == nil {
        request = NewDescribeRiskCenterServerRiskListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeRiskCenterServerRiskList")
    
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
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeRiskCenterVULViewVULRiskList")
    
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
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeRiskCenterWebsiteRiskList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRiskCenterWebsiteRiskList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRiskCenterWebsiteRiskListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRiskDetailListRequest() (request *DescribeRiskDetailListRequest) {
    request = &DescribeRiskDetailListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeRiskDetailList")
    
    
    return
}

func NewDescribeRiskDetailListResponse() (response *DescribeRiskDetailListResponse) {
    response = &DescribeRiskDetailListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRiskDetailList
// 风险详情列表示例
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
func (c *Client) DescribeRiskDetailList(request *DescribeRiskDetailListRequest) (response *DescribeRiskDetailListResponse, err error) {
    return c.DescribeRiskDetailListWithContext(context.Background(), request)
}

// DescribeRiskDetailList
// 风险详情列表示例
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
func (c *Client) DescribeRiskDetailListWithContext(ctx context.Context, request *DescribeRiskDetailListRequest) (response *DescribeRiskDetailListResponse, err error) {
    if request == nil {
        request = NewDescribeRiskDetailListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeRiskDetailList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRiskDetailList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRiskDetailListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRiskItemListRequest() (request *DescribeRiskItemListRequest) {
    request = &DescribeRiskItemListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeRiskItemList")
    
    
    return
}

func NewDescribeRiskItemListResponse() (response *DescribeRiskItemListResponse) {
    response = &DescribeRiskItemListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRiskItemList
// 获取风险项视角列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeRiskItemList(request *DescribeRiskItemListRequest) (response *DescribeRiskItemListResponse, err error) {
    return c.DescribeRiskItemListWithContext(context.Background(), request)
}

// DescribeRiskItemList
// 获取风险项视角列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeRiskItemListWithContext(ctx context.Context, request *DescribeRiskItemListRequest) (response *DescribeRiskItemListResponse, err error) {
    if request == nil {
        request = NewDescribeRiskItemListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeRiskItemList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRiskItemList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRiskItemListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRiskRuleDetailRequest() (request *DescribeRiskRuleDetailRequest) {
    request = &DescribeRiskRuleDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeRiskRuleDetail")
    
    
    return
}

func NewDescribeRiskRuleDetailResponse() (response *DescribeRiskRuleDetailResponse) {
    response = &DescribeRiskRuleDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRiskRuleDetail
// 查询风险规则详情示例
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeRiskRuleDetail(request *DescribeRiskRuleDetailRequest) (response *DescribeRiskRuleDetailResponse, err error) {
    return c.DescribeRiskRuleDetailWithContext(context.Background(), request)
}

// DescribeRiskRuleDetail
// 查询风险规则详情示例
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeRiskRuleDetailWithContext(ctx context.Context, request *DescribeRiskRuleDetailRequest) (response *DescribeRiskRuleDetailResponse, err error) {
    if request == nil {
        request = NewDescribeRiskRuleDetailRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeRiskRuleDetail")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRiskRuleDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRiskRuleDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRiskRulesRequest() (request *DescribeRiskRulesRequest) {
    request = &DescribeRiskRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeRiskRules")
    
    
    return
}

func NewDescribeRiskRulesResponse() (response *DescribeRiskRulesResponse) {
    response = &DescribeRiskRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRiskRules
// 高级配置风险规则列表示例
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeRiskRules(request *DescribeRiskRulesRequest) (response *DescribeRiskRulesResponse, err error) {
    return c.DescribeRiskRulesWithContext(context.Background(), request)
}

// DescribeRiskRules
// 高级配置风险规则列表示例
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeRiskRulesWithContext(ctx context.Context, request *DescribeRiskRulesRequest) (response *DescribeRiskRulesResponse, err error) {
    if request == nil {
        request = NewDescribeRiskRulesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeRiskRules")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRiskRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRiskRulesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRiskTrendDataRequest() (request *DescribeRiskTrendDataRequest) {
    request = &DescribeRiskTrendDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeRiskTrendData")
    
    
    return
}

func NewDescribeRiskTrendDataResponse() (response *DescribeRiskTrendDataResponse) {
    response = &DescribeRiskTrendDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRiskTrendData
// 查看风险趋势图
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeRiskTrendData(request *DescribeRiskTrendDataRequest) (response *DescribeRiskTrendDataResponse, err error) {
    return c.DescribeRiskTrendDataWithContext(context.Background(), request)
}

// DescribeRiskTrendData
// 查看风险趋势图
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeRiskTrendDataWithContext(ctx context.Context, request *DescribeRiskTrendDataRequest) (response *DescribeRiskTrendDataResponse, err error) {
    if request == nil {
        request = NewDescribeRiskTrendDataRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeRiskTrendData")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRiskTrendData require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRiskTrendDataResponse()
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
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeScanReportList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeScanReportList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeScanReportListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeScanStatisticRequest() (request *DescribeScanStatisticRequest) {
    request = &DescribeScanStatisticRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeScanStatistic")
    
    
    return
}

func NewDescribeScanStatisticResponse() (response *DescribeScanStatisticResponse) {
    response = &DescribeScanStatisticResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeScanStatistic
// 查询云边界分析扫描结果统计信息
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
func (c *Client) DescribeScanStatistic(request *DescribeScanStatisticRequest) (response *DescribeScanStatisticResponse, err error) {
    return c.DescribeScanStatisticWithContext(context.Background(), request)
}

// DescribeScanStatistic
// 查询云边界分析扫描结果统计信息
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
func (c *Client) DescribeScanStatisticWithContext(ctx context.Context, request *DescribeScanStatisticRequest) (response *DescribeScanStatisticResponse, err error) {
    if request == nil {
        request = NewDescribeScanStatisticRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeScanStatistic")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeScanStatistic require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeScanStatisticResponse()
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
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeScanTaskList")
    
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
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeSearchBugInfo")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSearchBugInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSearchBugInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSkillScanPayInfoRequest() (request *DescribeSkillScanPayInfoRequest) {
    request = &DescribeSkillScanPayInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeSkillScanPayInfo")
    
    
    return
}

func NewDescribeSkillScanPayInfoResponse() (response *DescribeSkillScanPayInfoResponse) {
    response = &DescribeSkillScanPayInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSkillScanPayInfo
// 查询 Skill 安全检测计费信息，包括订单状态、总配额、已消耗配额、到期时间、支付模式等。无订单时返回零值（仅含 TimeNow 和 BetaEndTime）。试用订单通过 ModifyTrialStatus(Module=9) 领取，正式订单通过计费系统创建。
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
func (c *Client) DescribeSkillScanPayInfo(request *DescribeSkillScanPayInfoRequest) (response *DescribeSkillScanPayInfoResponse, err error) {
    return c.DescribeSkillScanPayInfoWithContext(context.Background(), request)
}

// DescribeSkillScanPayInfo
// 查询 Skill 安全检测计费信息，包括订单状态、总配额、已消耗配额、到期时间、支付模式等。无订单时返回零值（仅含 TimeNow 和 BetaEndTime）。试用订单通过 ModifyTrialStatus(Module=9) 领取，正式订单通过计费系统创建。
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
func (c *Client) DescribeSkillScanPayInfoWithContext(ctx context.Context, request *DescribeSkillScanPayInfoRequest) (response *DescribeSkillScanPayInfoResponse, err error) {
    if request == nil {
        request = NewDescribeSkillScanPayInfoRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeSkillScanPayInfo")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSkillScanPayInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSkillScanPayInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSkillScanResultRequest() (request *DescribeSkillScanResultRequest) {
    request = &DescribeSkillScanResultRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeSkillScanResult")
    
    
    return
}

func NewDescribeSkillScanResultResponse() (response *DescribeSkillScanResultResponse) {
    response = &DescribeSkillScanResultResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSkillScanResult
// 查询 Skill 安全检测结果。调用 CreateSkillScan 成功后使用返回的 ContentHash + EngineVersion 轮询本接口获取结果。上传成功后建议5分钟后首次轮询，如未检测完成之后每隔1分钟轮询一次。响应通过 Status 字段区分四种状态：检测完成（SUCCESS）、检测中（SCANNING）、无记录（NOT_FOUND）、检测失败（FAILED）。注意：检测结果保留90天，超期后将返回 NOT_FOUND。
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
func (c *Client) DescribeSkillScanResult(request *DescribeSkillScanResultRequest) (response *DescribeSkillScanResultResponse, err error) {
    return c.DescribeSkillScanResultWithContext(context.Background(), request)
}

// DescribeSkillScanResult
// 查询 Skill 安全检测结果。调用 CreateSkillScan 成功后使用返回的 ContentHash + EngineVersion 轮询本接口获取结果。上传成功后建议5分钟后首次轮询，如未检测完成之后每隔1分钟轮询一次。响应通过 Status 字段区分四种状态：检测完成（SUCCESS）、检测中（SCANNING）、无记录（NOT_FOUND）、检测失败（FAILED）。注意：检测结果保留90天，超期后将返回 NOT_FOUND。
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
func (c *Client) DescribeSkillScanResultWithContext(ctx context.Context, request *DescribeSkillScanResultRequest) (response *DescribeSkillScanResultResponse, err error) {
    if request == nil {
        request = NewDescribeSkillScanResultRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeSkillScanResult")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSkillScanResult require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSkillScanResultResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSourceIPAssetRequest() (request *DescribeSourceIPAssetRequest) {
    request = &DescribeSourceIPAssetRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeSourceIPAsset")
    
    
    return
}

func NewDescribeSourceIPAssetResponse() (response *DescribeSourceIPAssetResponse) {
    response = &DescribeSourceIPAssetResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSourceIPAsset
// 获取用户访问密钥资产列表（源IP视角）
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
func (c *Client) DescribeSourceIPAsset(request *DescribeSourceIPAssetRequest) (response *DescribeSourceIPAssetResponse, err error) {
    return c.DescribeSourceIPAssetWithContext(context.Background(), request)
}

// DescribeSourceIPAsset
// 获取用户访问密钥资产列表（源IP视角）
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
func (c *Client) DescribeSourceIPAssetWithContext(ctx context.Context, request *DescribeSourceIPAssetRequest) (response *DescribeSourceIPAssetResponse, err error) {
    if request == nil {
        request = NewDescribeSourceIPAssetRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeSourceIPAsset")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSourceIPAsset require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSourceIPAssetResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSubUserInfoRequest() (request *DescribeSubUserInfoRequest) {
    request = &DescribeSubUserInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeSubUserInfo")
    
    
    return
}

func NewDescribeSubUserInfoResponse() (response *DescribeSubUserInfoResponse) {
    response = &DescribeSubUserInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSubUserInfo
// 查询集团的子账号列表
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
func (c *Client) DescribeSubUserInfo(request *DescribeSubUserInfoRequest) (response *DescribeSubUserInfoResponse, err error) {
    return c.DescribeSubUserInfoWithContext(context.Background(), request)
}

// DescribeSubUserInfo
// 查询集团的子账号列表
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
func (c *Client) DescribeSubUserInfoWithContext(ctx context.Context, request *DescribeSubUserInfoRequest) (response *DescribeSubUserInfoResponse, err error) {
    if request == nil {
        request = NewDescribeSubUserInfoRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeSubUserInfo")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSubUserInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSubUserInfoResponse()
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
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeSubnetAssets")
    
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
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeTaskLogList")
    
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
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeTaskLogURL")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTaskLogURL require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTaskLogURLResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTopAttackInfoRequest() (request *DescribeTopAttackInfoRequest) {
    request = &DescribeTopAttackInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeTopAttackInfo")
    
    
    return
}

func NewDescribeTopAttackInfoResponse() (response *DescribeTopAttackInfoResponse) {
    response = &DescribeTopAttackInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTopAttackInfo
// 查询TOP攻击信息
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
func (c *Client) DescribeTopAttackInfo(request *DescribeTopAttackInfoRequest) (response *DescribeTopAttackInfoResponse, err error) {
    return c.DescribeTopAttackInfoWithContext(context.Background(), request)
}

// DescribeTopAttackInfo
// 查询TOP攻击信息
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
func (c *Client) DescribeTopAttackInfoWithContext(ctx context.Context, request *DescribeTopAttackInfoRequest) (response *DescribeTopAttackInfoResponse, err error) {
    if request == nil {
        request = NewDescribeTopAttackInfoRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeTopAttackInfo")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTopAttackInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTopAttackInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUebaRuleRequest() (request *DescribeUebaRuleRequest) {
    request = &DescribeUebaRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeUebaRule")
    
    
    return
}

func NewDescribeUebaRuleResponse() (response *DescribeUebaRuleResponse) {
    response = &DescribeUebaRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeUebaRule
// 查询用户行为分析策略列表
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
func (c *Client) DescribeUebaRule(request *DescribeUebaRuleRequest) (response *DescribeUebaRuleResponse, err error) {
    return c.DescribeUebaRuleWithContext(context.Background(), request)
}

// DescribeUebaRule
// 查询用户行为分析策略列表
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
func (c *Client) DescribeUebaRuleWithContext(ctx context.Context, request *DescribeUebaRuleRequest) (response *DescribeUebaRuleResponse, err error) {
    if request == nil {
        request = NewDescribeUebaRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeUebaRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUebaRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeUebaRuleResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUserCallRecordRequest() (request *DescribeUserCallRecordRequest) {
    request = &DescribeUserCallRecordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeUserCallRecord")
    
    
    return
}

func NewDescribeUserCallRecordResponse() (response *DescribeUserCallRecordResponse) {
    response = &DescribeUserCallRecordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeUserCallRecord
// 获取账号调用记录列表
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
func (c *Client) DescribeUserCallRecord(request *DescribeUserCallRecordRequest) (response *DescribeUserCallRecordResponse, err error) {
    return c.DescribeUserCallRecordWithContext(context.Background(), request)
}

// DescribeUserCallRecord
// 获取账号调用记录列表
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
func (c *Client) DescribeUserCallRecordWithContext(ctx context.Context, request *DescribeUserCallRecordRequest) (response *DescribeUserCallRecordResponse, err error) {
    if request == nil {
        request = NewDescribeUserCallRecordRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeUserCallRecord")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUserCallRecord require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeUserCallRecordResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUserDspmInfoListRequest() (request *DescribeUserDspmInfoListRequest) {
    request = &DescribeUserDspmInfoListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeUserDspmInfoList")
    
    
    return
}

func NewDescribeUserDspmInfoListResponse() (response *DescribeUserDspmInfoListResponse) {
    response = &DescribeUserDspmInfoListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeUserDspmInfoList
// 获取账号dspm信息列表
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
func (c *Client) DescribeUserDspmInfoList(request *DescribeUserDspmInfoListRequest) (response *DescribeUserDspmInfoListResponse, err error) {
    return c.DescribeUserDspmInfoListWithContext(context.Background(), request)
}

// DescribeUserDspmInfoList
// 获取账号dspm信息列表
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
func (c *Client) DescribeUserDspmInfoListWithContext(ctx context.Context, request *DescribeUserDspmInfoListRequest) (response *DescribeUserDspmInfoListResponse, err error) {
    if request == nil {
        request = NewDescribeUserDspmInfoListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeUserDspmInfoList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUserDspmInfoList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeUserDspmInfoListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVULListRequest() (request *DescribeVULListRequest) {
    request = &DescribeVULListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeVULList")
    
    
    return
}

func NewDescribeVULListResponse() (response *DescribeVULListResponse) {
    response = &DescribeVULListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeVULList
// 新安全中心风险中心-漏洞列表
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
func (c *Client) DescribeVULList(request *DescribeVULListRequest) (response *DescribeVULListResponse, err error) {
    return c.DescribeVULListWithContext(context.Background(), request)
}

// DescribeVULList
// 新安全中心风险中心-漏洞列表
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
func (c *Client) DescribeVULListWithContext(ctx context.Context, request *DescribeVULListRequest) (response *DescribeVULListResponse, err error) {
    if request == nil {
        request = NewDescribeVULListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeVULList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVULList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVULListResponse()
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
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeVULRiskAdvanceCFGList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVULRiskAdvanceCFGList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVULRiskAdvanceCFGListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVULRiskDetailRequest() (request *DescribeVULRiskDetailRequest) {
    request = &DescribeVULRiskDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeVULRiskDetail")
    
    
    return
}

func NewDescribeVULRiskDetailResponse() (response *DescribeVULRiskDetailResponse) {
    response = &DescribeVULRiskDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeVULRiskDetail
// 获取漏洞展开详情
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
func (c *Client) DescribeVULRiskDetail(request *DescribeVULRiskDetailRequest) (response *DescribeVULRiskDetailResponse, err error) {
    return c.DescribeVULRiskDetailWithContext(context.Background(), request)
}

// DescribeVULRiskDetail
// 获取漏洞展开详情
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
func (c *Client) DescribeVULRiskDetailWithContext(ctx context.Context, request *DescribeVULRiskDetailRequest) (response *DescribeVULRiskDetailResponse, err error) {
    if request == nil {
        request = NewDescribeVULRiskDetailRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeVULRiskDetail")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVULRiskDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVULRiskDetailResponse()
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
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeVpcAssets")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVpcAssets require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVpcAssetsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVulRiskListRequest() (request *DescribeVulRiskListRequest) {
    request = &DescribeVulRiskListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeVulRiskList")
    
    
    return
}

func NewDescribeVulRiskListResponse() (response *DescribeVulRiskListResponse) {
    response = &DescribeVulRiskListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeVulRiskList
// 查询云边界分析-暴露路径下主机节点的漏洞列表
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
func (c *Client) DescribeVulRiskList(request *DescribeVulRiskListRequest) (response *DescribeVulRiskListResponse, err error) {
    return c.DescribeVulRiskListWithContext(context.Background(), request)
}

// DescribeVulRiskList
// 查询云边界分析-暴露路径下主机节点的漏洞列表
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
func (c *Client) DescribeVulRiskListWithContext(ctx context.Context, request *DescribeVulRiskListRequest) (response *DescribeVulRiskListResponse, err error) {
    if request == nil {
        request = NewDescribeVulRiskListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeVulRiskList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVulRiskList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVulRiskListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVulViewVulRiskListRequest() (request *DescribeVulViewVulRiskListRequest) {
    request = &DescribeVulViewVulRiskListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeVulViewVulRiskList")
    
    
    return
}

func NewDescribeVulViewVulRiskListResponse() (response *DescribeVulViewVulRiskListResponse) {
    response = &DescribeVulViewVulRiskListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeVulViewVulRiskList
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
func (c *Client) DescribeVulViewVulRiskList(request *DescribeVulViewVulRiskListRequest) (response *DescribeVulViewVulRiskListResponse, err error) {
    return c.DescribeVulViewVulRiskListWithContext(context.Background(), request)
}

// DescribeVulViewVulRiskList
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
func (c *Client) DescribeVulViewVulRiskListWithContext(ctx context.Context, request *DescribeVulViewVulRiskListRequest) (response *DescribeVulViewVulRiskListResponse, err error) {
    if request == nil {
        request = NewDescribeVulViewVulRiskListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeVulViewVulRiskList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVulViewVulRiskList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVulViewVulRiskListResponse()
    err = c.Send(request, response)
    return
}

func NewDownloadDspmExportLogRequest() (request *DownloadDspmExportLogRequest) {
    request = &DownloadDspmExportLogRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DownloadDspmExportLog")
    
    
    return
}

func NewDownloadDspmExportLogResponse() (response *DownloadDspmExportLogResponse) {
    response = &DownloadDspmExportLogResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DownloadDspmExportLog
// 下载导出日志
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
func (c *Client) DownloadDspmExportLog(request *DownloadDspmExportLogRequest) (response *DownloadDspmExportLogResponse, err error) {
    return c.DownloadDspmExportLogWithContext(context.Background(), request)
}

// DownloadDspmExportLog
// 下载导出日志
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
func (c *Client) DownloadDspmExportLogWithContext(ctx context.Context, request *DownloadDspmExportLogRequest) (response *DownloadDspmExportLogResponse, err error) {
    if request == nil {
        request = NewDownloadDspmExportLogRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DownloadDspmExportLog")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DownloadDspmExportLog require credential")
    }

    request.SetContext(ctx)
    
    response = NewDownloadDspmExportLogResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAlarmRiskStatusRequest() (request *ModifyAlarmRiskStatusRequest) {
    request = &ModifyAlarmRiskStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "ModifyAlarmRiskStatus")
    
    
    return
}

func NewModifyAlarmRiskStatusResponse() (response *ModifyAlarmRiskStatusResponse) {
    response = &ModifyAlarmRiskStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyAlarmRiskStatus
// 修改或者更改处置状态
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyAlarmRiskStatus(request *ModifyAlarmRiskStatusRequest) (response *ModifyAlarmRiskStatusResponse, err error) {
    return c.ModifyAlarmRiskStatusWithContext(context.Background(), request)
}

// ModifyAlarmRiskStatus
// 修改或者更改处置状态
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyAlarmRiskStatusWithContext(ctx context.Context, request *ModifyAlarmRiskStatusRequest) (response *ModifyAlarmRiskStatusResponse, err error) {
    if request == nil {
        request = NewModifyAlarmRiskStatusRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "ModifyAlarmRiskStatus")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAlarmRiskStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAlarmRiskStatusResponse()
    err = c.Send(request, response)
    return
}

func NewModifyCosAuditMonitorAccountRequest() (request *ModifyCosAuditMonitorAccountRequest) {
    request = &ModifyCosAuditMonitorAccountRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "ModifyCosAuditMonitorAccount")
    
    
    return
}

func NewModifyCosAuditMonitorAccountResponse() (response *ModifyCosAuditMonitorAccountResponse) {
    response = &ModifyCosAuditMonitorAccountResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyCosAuditMonitorAccount
// 修改cos审计监测账号
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) ModifyCosAuditMonitorAccount(request *ModifyCosAuditMonitorAccountRequest) (response *ModifyCosAuditMonitorAccountResponse, err error) {
    return c.ModifyCosAuditMonitorAccountWithContext(context.Background(), request)
}

// ModifyCosAuditMonitorAccount
// 修改cos审计监测账号
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) ModifyCosAuditMonitorAccountWithContext(ctx context.Context, request *ModifyCosAuditMonitorAccountRequest) (response *ModifyCosAuditMonitorAccountResponse, err error) {
    if request == nil {
        request = NewModifyCosAuditMonitorAccountRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "ModifyCosAuditMonitorAccount")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyCosAuditMonitorAccount require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyCosAuditMonitorAccountResponse()
    err = c.Send(request, response)
    return
}

func NewModifyCosMarkInfoRequest() (request *ModifyCosMarkInfoRequest) {
    request = &ModifyCosMarkInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "ModifyCosMarkInfo")
    
    
    return
}

func NewModifyCosMarkInfoResponse() (response *ModifyCosMarkInfoResponse) {
    response = &ModifyCosMarkInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyCosMarkInfo
// 修改对象存储备注
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyCosMarkInfo(request *ModifyCosMarkInfoRequest) (response *ModifyCosMarkInfoResponse, err error) {
    return c.ModifyCosMarkInfoWithContext(context.Background(), request)
}

// ModifyCosMarkInfo
// 修改对象存储备注
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyCosMarkInfoWithContext(ctx context.Context, request *ModifyCosMarkInfoRequest) (response *ModifyCosMarkInfoResponse, err error) {
    if request == nil {
        request = NewModifyCosMarkInfoRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "ModifyCosMarkInfo")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyCosMarkInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyCosMarkInfoResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDspmAccessRecordRequest() (request *ModifyDspmAccessRecordRequest) {
    request = &ModifyDspmAccessRecordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "ModifyDspmAccessRecord")
    
    
    return
}

func NewModifyDspmAccessRecordResponse() (response *ModifyDspmAccessRecordResponse) {
    response = &ModifyDspmAccessRecordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyDspmAccessRecord
// 修改Dspm访问管理信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) ModifyDspmAccessRecord(request *ModifyDspmAccessRecordRequest) (response *ModifyDspmAccessRecordResponse, err error) {
    return c.ModifyDspmAccessRecordWithContext(context.Background(), request)
}

// ModifyDspmAccessRecord
// 修改Dspm访问管理信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) ModifyDspmAccessRecordWithContext(ctx context.Context, request *ModifyDspmAccessRecordRequest) (response *ModifyDspmAccessRecordResponse, err error) {
    if request == nil {
        request = NewModifyDspmAccessRecordRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "ModifyDspmAccessRecord")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDspmAccessRecord require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDspmAccessRecordResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDspmApproveStatusRequest() (request *ModifyDspmApproveStatusRequest) {
    request = &ModifyDspmApproveStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "ModifyDspmApproveStatus")
    
    
    return
}

func NewModifyDspmApproveStatusResponse() (response *ModifyDspmApproveStatusResponse) {
    response = &ModifyDspmApproveStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyDspmApproveStatus
// 修改Dspm审批单状态
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyDspmApproveStatus(request *ModifyDspmApproveStatusRequest) (response *ModifyDspmApproveStatusResponse, err error) {
    return c.ModifyDspmApproveStatusWithContext(context.Background(), request)
}

// ModifyDspmApproveStatus
// 修改Dspm审批单状态
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyDspmApproveStatusWithContext(ctx context.Context, request *ModifyDspmApproveStatusRequest) (response *ModifyDspmApproveStatusResponse, err error) {
    if request == nil {
        request = NewModifyDspmApproveStatusRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "ModifyDspmApproveStatus")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDspmApproveStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDspmApproveStatusResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDspmAssetAccountRequest() (request *ModifyDspmAssetAccountRequest) {
    request = &ModifyDspmAssetAccountRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "ModifyDspmAssetAccount")
    
    
    return
}

func NewModifyDspmAssetAccountResponse() (response *ModifyDspmAssetAccountResponse) {
    response = &ModifyDspmAssetAccountResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyDspmAssetAccount
// 修改Dspm资产账号信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyDspmAssetAccount(request *ModifyDspmAssetAccountRequest) (response *ModifyDspmAssetAccountResponse, err error) {
    return c.ModifyDspmAssetAccountWithContext(context.Background(), request)
}

// ModifyDspmAssetAccount
// 修改Dspm资产账号信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyDspmAssetAccountWithContext(ctx context.Context, request *ModifyDspmAssetAccountRequest) (response *ModifyDspmAssetAccountResponse, err error) {
    if request == nil {
        request = NewModifyDspmAssetAccountRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "ModifyDspmAssetAccount")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDspmAssetAccount require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDspmAssetAccountResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDspmAssetAccountPrivilegesRequest() (request *ModifyDspmAssetAccountPrivilegesRequest) {
    request = &ModifyDspmAssetAccountPrivilegesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "ModifyDspmAssetAccountPrivileges")
    
    
    return
}

func NewModifyDspmAssetAccountPrivilegesResponse() (response *ModifyDspmAssetAccountPrivilegesResponse) {
    response = &ModifyDspmAssetAccountPrivilegesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyDspmAssetAccountPrivileges
// 修改Dspm资产账号权限
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyDspmAssetAccountPrivileges(request *ModifyDspmAssetAccountPrivilegesRequest) (response *ModifyDspmAssetAccountPrivilegesResponse, err error) {
    return c.ModifyDspmAssetAccountPrivilegesWithContext(context.Background(), request)
}

// ModifyDspmAssetAccountPrivileges
// 修改Dspm资产账号权限
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyDspmAssetAccountPrivilegesWithContext(ctx context.Context, request *ModifyDspmAssetAccountPrivilegesRequest) (response *ModifyDspmAssetAccountPrivilegesResponse, err error) {
    if request == nil {
        request = NewModifyDspmAssetAccountPrivilegesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "ModifyDspmAssetAccountPrivileges")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDspmAssetAccountPrivileges require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDspmAssetAccountPrivilegesResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDspmAssetDataScanTaskRequest() (request *ModifyDspmAssetDataScanTaskRequest) {
    request = &ModifyDspmAssetDataScanTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "ModifyDspmAssetDataScanTask")
    
    
    return
}

func NewModifyDspmAssetDataScanTaskResponse() (response *ModifyDspmAssetDataScanTaskResponse) {
    response = &ModifyDspmAssetDataScanTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyDspmAssetDataScanTask
// 修改Dspm资产数据扫描任务
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyDspmAssetDataScanTask(request *ModifyDspmAssetDataScanTaskRequest) (response *ModifyDspmAssetDataScanTaskResponse, err error) {
    return c.ModifyDspmAssetDataScanTaskWithContext(context.Background(), request)
}

// ModifyDspmAssetDataScanTask
// 修改Dspm资产数据扫描任务
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyDspmAssetDataScanTaskWithContext(ctx context.Context, request *ModifyDspmAssetDataScanTaskRequest) (response *ModifyDspmAssetDataScanTaskResponse, err error) {
    if request == nil {
        request = NewModifyDspmAssetDataScanTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "ModifyDspmAssetDataScanTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDspmAssetDataScanTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDspmAssetDataScanTaskResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDspmAssetLogDeliverySwitchRequest() (request *ModifyDspmAssetLogDeliverySwitchRequest) {
    request = &ModifyDspmAssetLogDeliverySwitchRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "ModifyDspmAssetLogDeliverySwitch")
    
    
    return
}

func NewModifyDspmAssetLogDeliverySwitchResponse() (response *ModifyDspmAssetLogDeliverySwitchResponse) {
    response = &ModifyDspmAssetLogDeliverySwitchResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyDspmAssetLogDeliverySwitch
// 修改Dspm资产日志投递开关
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyDspmAssetLogDeliverySwitch(request *ModifyDspmAssetLogDeliverySwitchRequest) (response *ModifyDspmAssetLogDeliverySwitchResponse, err error) {
    return c.ModifyDspmAssetLogDeliverySwitchWithContext(context.Background(), request)
}

// ModifyDspmAssetLogDeliverySwitch
// 修改Dspm资产日志投递开关
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyDspmAssetLogDeliverySwitchWithContext(ctx context.Context, request *ModifyDspmAssetLogDeliverySwitchRequest) (response *ModifyDspmAssetLogDeliverySwitchResponse, err error) {
    if request == nil {
        request = NewModifyDspmAssetLogDeliverySwitchRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "ModifyDspmAssetLogDeliverySwitch")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDspmAssetLogDeliverySwitch require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDspmAssetLogDeliverySwitchResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDspmAssetSecurityAnalysisSwitchRequest() (request *ModifyDspmAssetSecurityAnalysisSwitchRequest) {
    request = &ModifyDspmAssetSecurityAnalysisSwitchRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "ModifyDspmAssetSecurityAnalysisSwitch")
    
    
    return
}

func NewModifyDspmAssetSecurityAnalysisSwitchResponse() (response *ModifyDspmAssetSecurityAnalysisSwitchResponse) {
    response = &ModifyDspmAssetSecurityAnalysisSwitchResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyDspmAssetSecurityAnalysisSwitch
// 修改Dspm资产日志投递开关
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyDspmAssetSecurityAnalysisSwitch(request *ModifyDspmAssetSecurityAnalysisSwitchRequest) (response *ModifyDspmAssetSecurityAnalysisSwitchResponse, err error) {
    return c.ModifyDspmAssetSecurityAnalysisSwitchWithContext(context.Background(), request)
}

// ModifyDspmAssetSecurityAnalysisSwitch
// 修改Dspm资产日志投递开关
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyDspmAssetSecurityAnalysisSwitchWithContext(ctx context.Context, request *ModifyDspmAssetSecurityAnalysisSwitchRequest) (response *ModifyDspmAssetSecurityAnalysisSwitchResponse, err error) {
    if request == nil {
        request = NewModifyDspmAssetSecurityAnalysisSwitchRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "ModifyDspmAssetSecurityAnalysisSwitch")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDspmAssetSecurityAnalysisSwitch require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDspmAssetSecurityAnalysisSwitchResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDspmBackupSettingRequest() (request *ModifyDspmBackupSettingRequest) {
    request = &ModifyDspmBackupSettingRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "ModifyDspmBackupSetting")
    
    
    return
}

func NewModifyDspmBackupSettingResponse() (response *ModifyDspmBackupSettingResponse) {
    response = &ModifyDspmBackupSettingResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyDspmBackupSetting
// 修改日志备份设置
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyDspmBackupSetting(request *ModifyDspmBackupSettingRequest) (response *ModifyDspmBackupSettingResponse, err error) {
    return c.ModifyDspmBackupSettingWithContext(context.Background(), request)
}

// ModifyDspmBackupSetting
// 修改日志备份设置
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyDspmBackupSettingWithContext(ctx context.Context, request *ModifyDspmBackupSettingRequest) (response *ModifyDspmBackupSettingResponse, err error) {
    if request == nil {
        request = NewModifyDspmBackupSettingRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "ModifyDspmBackupSetting")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDspmBackupSetting require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDspmBackupSettingResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDspmIdentifyInfoRequest() (request *ModifyDspmIdentifyInfoRequest) {
    request = &ModifyDspmIdentifyInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "ModifyDspmIdentifyInfo")
    
    
    return
}

func NewModifyDspmIdentifyInfoResponse() (response *ModifyDspmIdentifyInfoResponse) {
    response = &ModifyDspmIdentifyInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyDspmIdentifyInfo
// 修改Dspm身份信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyDspmIdentifyInfo(request *ModifyDspmIdentifyInfoRequest) (response *ModifyDspmIdentifyInfoResponse, err error) {
    return c.ModifyDspmIdentifyInfoWithContext(context.Background(), request)
}

// ModifyDspmIdentifyInfo
// 修改Dspm身份信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyDspmIdentifyInfoWithContext(ctx context.Context, request *ModifyDspmIdentifyInfoRequest) (response *ModifyDspmIdentifyInfoResponse, err error) {
    if request == nil {
        request = NewModifyDspmIdentifyInfoRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "ModifyDspmIdentifyInfo")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDspmIdentifyInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDspmIdentifyInfoResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDspmIpInfoRequest() (request *ModifyDspmIpInfoRequest) {
    request = &ModifyDspmIpInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "ModifyDspmIpInfo")
    
    
    return
}

func NewModifyDspmIpInfoResponse() (response *ModifyDspmIpInfoResponse) {
    response = &ModifyDspmIpInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyDspmIpInfo
// 修改DspmIp信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyDspmIpInfo(request *ModifyDspmIpInfoRequest) (response *ModifyDspmIpInfoResponse, err error) {
    return c.ModifyDspmIpInfoWithContext(context.Background(), request)
}

// ModifyDspmIpInfo
// 修改DspmIp信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyDspmIpInfoWithContext(ctx context.Context, request *ModifyDspmIpInfoRequest) (response *ModifyDspmIpInfoResponse, err error) {
    if request == nil {
        request = NewModifyDspmIpInfoRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "ModifyDspmIpInfo")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDspmIpInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDspmIpInfoResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDspmPersonalIdentifyRequest() (request *ModifyDspmPersonalIdentifyRequest) {
    request = &ModifyDspmPersonalIdentifyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "ModifyDspmPersonalIdentify")
    
    
    return
}

func NewModifyDspmPersonalIdentifyResponse() (response *ModifyDspmPersonalIdentifyResponse) {
    response = &ModifyDspmPersonalIdentifyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyDspmPersonalIdentify
// 修改Dspm个人身份id
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) ModifyDspmPersonalIdentify(request *ModifyDspmPersonalIdentifyRequest) (response *ModifyDspmPersonalIdentifyResponse, err error) {
    return c.ModifyDspmPersonalIdentifyWithContext(context.Background(), request)
}

// ModifyDspmPersonalIdentify
// 修改Dspm个人身份id
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) ModifyDspmPersonalIdentifyWithContext(ctx context.Context, request *ModifyDspmPersonalIdentifyRequest) (response *ModifyDspmPersonalIdentifyResponse, err error) {
    if request == nil {
        request = NewModifyDspmPersonalIdentifyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "ModifyDspmPersonalIdentify")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDspmPersonalIdentify require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDspmPersonalIdentifyResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDspmRestoreLogTaskRequest() (request *ModifyDspmRestoreLogTaskRequest) {
    request = &ModifyDspmRestoreLogTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "ModifyDspmRestoreLogTask")
    
    
    return
}

func NewModifyDspmRestoreLogTaskResponse() (response *ModifyDspmRestoreLogTaskResponse) {
    response = &ModifyDspmRestoreLogTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyDspmRestoreLogTask
// 恢复备份日志
//
// 可能返回的错误码:
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyDspmRestoreLogTask(request *ModifyDspmRestoreLogTaskRequest) (response *ModifyDspmRestoreLogTaskResponse, err error) {
    return c.ModifyDspmRestoreLogTaskWithContext(context.Background(), request)
}

// ModifyDspmRestoreLogTask
// 恢复备份日志
//
// 可能返回的错误码:
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyDspmRestoreLogTaskWithContext(ctx context.Context, request *ModifyDspmRestoreLogTaskRequest) (response *ModifyDspmRestoreLogTaskResponse, err error) {
    if request == nil {
        request = NewModifyDspmRestoreLogTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "ModifyDspmRestoreLogTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDspmRestoreLogTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDspmRestoreLogTaskResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDspmRiskInfoRequest() (request *ModifyDspmRiskInfoRequest) {
    request = &ModifyDspmRiskInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "ModifyDspmRiskInfo")
    
    
    return
}

func NewModifyDspmRiskInfoResponse() (response *ModifyDspmRiskInfoResponse) {
    response = &ModifyDspmRiskInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyDspmRiskInfo
// 修改Dspm风险信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyDspmRiskInfo(request *ModifyDspmRiskInfoRequest) (response *ModifyDspmRiskInfoResponse, err error) {
    return c.ModifyDspmRiskInfoWithContext(context.Background(), request)
}

// ModifyDspmRiskInfo
// 修改Dspm风险信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyDspmRiskInfoWithContext(ctx context.Context, request *ModifyDspmRiskInfoRequest) (response *ModifyDspmRiskInfoResponse, err error) {
    if request == nil {
        request = NewModifyDspmRiskInfoRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "ModifyDspmRiskInfo")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDspmRiskInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDspmRiskInfoResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDspmRiskStrategyRequest() (request *ModifyDspmRiskStrategyRequest) {
    request = &ModifyDspmRiskStrategyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "ModifyDspmRiskStrategy")
    
    
    return
}

func NewModifyDspmRiskStrategyResponse() (response *ModifyDspmRiskStrategyResponse) {
    response = &ModifyDspmRiskStrategyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyDspmRiskStrategy
// 修改Dspm风险策略
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) ModifyDspmRiskStrategy(request *ModifyDspmRiskStrategyRequest) (response *ModifyDspmRiskStrategyResponse, err error) {
    return c.ModifyDspmRiskStrategyWithContext(context.Background(), request)
}

// ModifyDspmRiskStrategy
// 修改Dspm风险策略
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) ModifyDspmRiskStrategyWithContext(ctx context.Context, request *ModifyDspmRiskStrategyRequest) (response *ModifyDspmRiskStrategyResponse, err error) {
    if request == nil {
        request = NewModifyDspmRiskStrategyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "ModifyDspmRiskStrategy")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDspmRiskStrategy require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDspmRiskStrategyResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDspmWhitelistStrategyRequest() (request *ModifyDspmWhitelistStrategyRequest) {
    request = &ModifyDspmWhitelistStrategyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "ModifyDspmWhitelistStrategy")
    
    
    return
}

func NewModifyDspmWhitelistStrategyResponse() (response *ModifyDspmWhitelistStrategyResponse) {
    response = &ModifyDspmWhitelistStrategyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyDspmWhitelistStrategy
// 修改Dspm白名单策略
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) ModifyDspmWhitelistStrategy(request *ModifyDspmWhitelistStrategyRequest) (response *ModifyDspmWhitelistStrategyResponse, err error) {
    return c.ModifyDspmWhitelistStrategyWithContext(context.Background(), request)
}

// ModifyDspmWhitelistStrategy
// 修改Dspm白名单策略
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) ModifyDspmWhitelistStrategyWithContext(ctx context.Context, request *ModifyDspmWhitelistStrategyRequest) (response *ModifyDspmWhitelistStrategyResponse, err error) {
    if request == nil {
        request = NewModifyDspmWhitelistStrategyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "ModifyDspmWhitelistStrategy")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDspmWhitelistStrategy require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDspmWhitelistStrategyResponse()
    err = c.Send(request, response)
    return
}

func NewModifyIaCTokenPeriodRequest() (request *ModifyIaCTokenPeriodRequest) {
    request = &ModifyIaCTokenPeriodRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "ModifyIaCTokenPeriod")
    
    
    return
}

func NewModifyIaCTokenPeriodResponse() (response *ModifyIaCTokenPeriodResponse) {
    response = &ModifyIaCTokenPeriodResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyIaCTokenPeriod
// 修改IaC检测接入Token存储周期
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyIaCTokenPeriod(request *ModifyIaCTokenPeriodRequest) (response *ModifyIaCTokenPeriodResponse, err error) {
    return c.ModifyIaCTokenPeriodWithContext(context.Background(), request)
}

// ModifyIaCTokenPeriod
// 修改IaC检测接入Token存储周期
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyIaCTokenPeriodWithContext(ctx context.Context, request *ModifyIaCTokenPeriodRequest) (response *ModifyIaCTokenPeriodResponse, err error) {
    if request == nil {
        request = NewModifyIaCTokenPeriodRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "ModifyIaCTokenPeriod")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyIaCTokenPeriod require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyIaCTokenPeriodResponse()
    err = c.Send(request, response)
    return
}

func NewModifyMachineRemarkRequest() (request *ModifyMachineRemarkRequest) {
    request = &ModifyMachineRemarkRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "ModifyMachineRemark")
    
    
    return
}

func NewModifyMachineRemarkResponse() (response *ModifyMachineRemarkResponse) {
    response = &ModifyMachineRemarkResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyMachineRemark
// 修改主机资产备注信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyMachineRemark(request *ModifyMachineRemarkRequest) (response *ModifyMachineRemarkResponse, err error) {
    return c.ModifyMachineRemarkWithContext(context.Background(), request)
}

// ModifyMachineRemark
// 修改主机资产备注信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyMachineRemarkWithContext(ctx context.Context, request *ModifyMachineRemarkRequest) (response *ModifyMachineRemarkResponse, err error) {
    if request == nil {
        request = NewModifyMachineRemarkRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "ModifyMachineRemark")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyMachineRemark require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyMachineRemarkResponse()
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
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyOrganizationAccountStatus(request *ModifyOrganizationAccountStatusRequest) (response *ModifyOrganizationAccountStatusResponse, err error) {
    return c.ModifyOrganizationAccountStatusWithContext(context.Background(), request)
}

// ModifyOrganizationAccountStatus
// 修改集团账号状态
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyOrganizationAccountStatusWithContext(ctx context.Context, request *ModifyOrganizationAccountStatusRequest) (response *ModifyOrganizationAccountStatusResponse, err error) {
    if request == nil {
        request = NewModifyOrganizationAccountStatusRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "ModifyOrganizationAccountStatus")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyOrganizationAccountStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyOrganizationAccountStatusResponse()
    err = c.Send(request, response)
    return
}

func NewModifyPolicyStatusRequest() (request *ModifyPolicyStatusRequest) {
    request = &ModifyPolicyStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "ModifyPolicyStatus")
    
    
    return
}

func NewModifyPolicyStatusResponse() (response *ModifyPolicyStatusResponse) {
    response = &ModifyPolicyStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyPolicyStatus
// 修改策略状态
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) ModifyPolicyStatus(request *ModifyPolicyStatusRequest) (response *ModifyPolicyStatusResponse, err error) {
    return c.ModifyPolicyStatusWithContext(context.Background(), request)
}

// ModifyPolicyStatus
// 修改策略状态
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) ModifyPolicyStatusWithContext(ctx context.Context, request *ModifyPolicyStatusRequest) (response *ModifyPolicyStatusResponse, err error) {
    if request == nil {
        request = NewModifyPolicyStatusRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "ModifyPolicyStatus")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyPolicyStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyPolicyStatusResponse()
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
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "ModifyRiskCenterRiskStatus")
    
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
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "ModifyRiskCenterScanTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyRiskCenterScanTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyRiskCenterScanTaskResponse()
    err = c.Send(request, response)
    return
}

func NewModifyUebaRuleSwitchRequest() (request *ModifyUebaRuleSwitchRequest) {
    request = &ModifyUebaRuleSwitchRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "ModifyUebaRuleSwitch")
    
    
    return
}

func NewModifyUebaRuleSwitchResponse() (response *ModifyUebaRuleSwitchResponse) {
    response = &ModifyUebaRuleSwitchResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyUebaRuleSwitch
// 更新自定义策略的开关
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
func (c *Client) ModifyUebaRuleSwitch(request *ModifyUebaRuleSwitchRequest) (response *ModifyUebaRuleSwitchResponse, err error) {
    return c.ModifyUebaRuleSwitchWithContext(context.Background(), request)
}

// ModifyUebaRuleSwitch
// 更新自定义策略的开关
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
func (c *Client) ModifyUebaRuleSwitchWithContext(ctx context.Context, request *ModifyUebaRuleSwitchRequest) (response *ModifyUebaRuleSwitchResponse, err error) {
    if request == nil {
        request = NewModifyUebaRuleSwitchRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "ModifyUebaRuleSwitch")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyUebaRuleSwitch require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyUebaRuleSwitchResponse()
    err = c.Send(request, response)
    return
}

func NewResetDspmAssetAccountPasswordRequest() (request *ResetDspmAssetAccountPasswordRequest) {
    request = &ResetDspmAssetAccountPasswordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "ResetDspmAssetAccountPassword")
    
    
    return
}

func NewResetDspmAssetAccountPasswordResponse() (response *ResetDspmAssetAccountPasswordResponse) {
    response = &ResetDspmAssetAccountPasswordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ResetDspmAssetAccountPassword
// 重置Dspm资产账号密码
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ResetDspmAssetAccountPassword(request *ResetDspmAssetAccountPasswordRequest) (response *ResetDspmAssetAccountPasswordResponse, err error) {
    return c.ResetDspmAssetAccountPasswordWithContext(context.Background(), request)
}

// ResetDspmAssetAccountPassword
// 重置Dspm资产账号密码
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ResetDspmAssetAccountPasswordWithContext(ctx context.Context, request *ResetDspmAssetAccountPasswordRequest) (response *ResetDspmAssetAccountPasswordResponse, err error) {
    if request == nil {
        request = NewResetDspmAssetAccountPasswordRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "ResetDspmAssetAccountPassword")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ResetDspmAssetAccountPassword require credential")
    }

    request.SetContext(ctx)
    
    response = NewResetDspmAssetAccountPasswordResponse()
    err = c.Send(request, response)
    return
}

func NewRetryDspmExportLogRequest() (request *RetryDspmExportLogRequest) {
    request = &RetryDspmExportLogRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "RetryDspmExportLog")
    
    
    return
}

func NewRetryDspmExportLogResponse() (response *RetryDspmExportLogResponse) {
    response = &RetryDspmExportLogResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RetryDspmExportLog
// RetryExportLog
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) RetryDspmExportLog(request *RetryDspmExportLogRequest) (response *RetryDspmExportLogResponse, err error) {
    return c.RetryDspmExportLogWithContext(context.Background(), request)
}

// RetryDspmExportLog
// RetryExportLog
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) RetryDspmExportLogWithContext(ctx context.Context, request *RetryDspmExportLogRequest) (response *RetryDspmExportLogResponse, err error) {
    if request == nil {
        request = NewRetryDspmExportLogRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "RetryDspmExportLog")
    
    if c.GetCredential() == nil {
        return nil, errors.New("RetryDspmExportLog require credential")
    }

    request.SetContext(ctx)
    
    response = NewRetryDspmExportLogResponse()
    err = c.Send(request, response)
    return
}

func NewRevertDspmAssetAccountRequest() (request *RevertDspmAssetAccountRequest) {
    request = &RevertDspmAssetAccountRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "RevertDspmAssetAccount")
    
    
    return
}

func NewRevertDspmAssetAccountResponse() (response *RevertDspmAssetAccountResponse) {
    response = &RevertDspmAssetAccountResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RevertDspmAssetAccount
// 恢复Dspm资产账号
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) RevertDspmAssetAccount(request *RevertDspmAssetAccountRequest) (response *RevertDspmAssetAccountResponse, err error) {
    return c.RevertDspmAssetAccountWithContext(context.Background(), request)
}

// RevertDspmAssetAccount
// 恢复Dspm资产账号
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) RevertDspmAssetAccountWithContext(ctx context.Context, request *RevertDspmAssetAccountRequest) (response *RevertDspmAssetAccountResponse, err error) {
    if request == nil {
        request = NewRevertDspmAssetAccountRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "RevertDspmAssetAccount")
    
    if c.GetCredential() == nil {
        return nil, errors.New("RevertDspmAssetAccount require credential")
    }

    request.SetContext(ctx)
    
    response = NewRevertDspmAssetAccountResponse()
    err = c.Send(request, response)
    return
}

func NewSendDspmAssetLoginSmsCodeRequest() (request *SendDspmAssetLoginSmsCodeRequest) {
    request = &SendDspmAssetLoginSmsCodeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "SendDspmAssetLoginSmsCode")
    
    
    return
}

func NewSendDspmAssetLoginSmsCodeResponse() (response *SendDspmAssetLoginSmsCodeResponse) {
    response = &SendDspmAssetLoginSmsCodeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SendDspmAssetLoginSmsCode
// 发送Dspm资产访问验证码
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) SendDspmAssetLoginSmsCode(request *SendDspmAssetLoginSmsCodeRequest) (response *SendDspmAssetLoginSmsCodeResponse, err error) {
    return c.SendDspmAssetLoginSmsCodeWithContext(context.Background(), request)
}

// SendDspmAssetLoginSmsCode
// 发送Dspm资产访问验证码
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) SendDspmAssetLoginSmsCodeWithContext(ctx context.Context, request *SendDspmAssetLoginSmsCodeRequest) (response *SendDspmAssetLoginSmsCodeResponse, err error) {
    if request == nil {
        request = NewSendDspmAssetLoginSmsCodeRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "SendDspmAssetLoginSmsCode")
    
    if c.GetCredential() == nil {
        return nil, errors.New("SendDspmAssetLoginSmsCode require credential")
    }

    request.SetContext(ctx)
    
    response = NewSendDspmAssetLoginSmsCodeResponse()
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
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "StopRiskCenterTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("StopRiskCenterTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewStopRiskCenterTaskResponse()
    err = c.Send(request, response)
    return
}

func NewSyncDspmAssetsRequest() (request *SyncDspmAssetsRequest) {
    request = &SyncDspmAssetsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "SyncDspmAssets")
    
    
    return
}

func NewSyncDspmAssetsResponse() (response *SyncDspmAssetsResponse) {
    response = &SyncDspmAssetsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SyncDspmAssets
// 同步dspm支持的资产
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) SyncDspmAssets(request *SyncDspmAssetsRequest) (response *SyncDspmAssetsResponse, err error) {
    return c.SyncDspmAssetsWithContext(context.Background(), request)
}

// SyncDspmAssets
// 同步dspm支持的资产
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) SyncDspmAssetsWithContext(ctx context.Context, request *SyncDspmAssetsRequest) (response *SyncDspmAssetsResponse, err error) {
    if request == nil {
        request = NewSyncDspmAssetsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "SyncDspmAssets")
    
    if c.GetCredential() == nil {
        return nil, errors.New("SyncDspmAssets require credential")
    }

    request.SetContext(ctx)
    
    response = NewSyncDspmAssetsResponse()
    err = c.Send(request, response)
    return
}

func NewSyncDspmUsersRequest() (request *SyncDspmUsersRequest) {
    request = &SyncDspmUsersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "SyncDspmUsers")
    
    
    return
}

func NewSyncDspmUsersResponse() (response *SyncDspmUsersResponse) {
    response = &SyncDspmUsersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SyncDspmUsers
// 同步dspm用户列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) SyncDspmUsers(request *SyncDspmUsersRequest) (response *SyncDspmUsersResponse, err error) {
    return c.SyncDspmUsersWithContext(context.Background(), request)
}

// SyncDspmUsers
// 同步dspm用户列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) SyncDspmUsersWithContext(ctx context.Context, request *SyncDspmUsersRequest) (response *SyncDspmUsersResponse, err error) {
    if request == nil {
        request = NewSyncDspmUsersRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "SyncDspmUsers")
    
    if c.GetCredential() == nil {
        return nil, errors.New("SyncDspmUsers require credential")
    }

    request.SetContext(ctx)
    
    response = NewSyncDspmUsersResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateAccessKeyAlarmStatusRequest() (request *UpdateAccessKeyAlarmStatusRequest) {
    request = &UpdateAccessKeyAlarmStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "UpdateAccessKeyAlarmStatus")
    
    
    return
}

func NewUpdateAccessKeyAlarmStatusResponse() (response *UpdateAccessKeyAlarmStatusResponse) {
    response = &UpdateAccessKeyAlarmStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateAccessKeyAlarmStatus
// 标记风险或者告警为 已处置/已忽略
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) UpdateAccessKeyAlarmStatus(request *UpdateAccessKeyAlarmStatusRequest) (response *UpdateAccessKeyAlarmStatusResponse, err error) {
    return c.UpdateAccessKeyAlarmStatusWithContext(context.Background(), request)
}

// UpdateAccessKeyAlarmStatus
// 标记风险或者告警为 已处置/已忽略
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) UpdateAccessKeyAlarmStatusWithContext(ctx context.Context, request *UpdateAccessKeyAlarmStatusRequest) (response *UpdateAccessKeyAlarmStatusResponse, err error) {
    if request == nil {
        request = NewUpdateAccessKeyAlarmStatusRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "UpdateAccessKeyAlarmStatus")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateAccessKeyAlarmStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateAccessKeyAlarmStatusResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateAccessKeyRemarkRequest() (request *UpdateAccessKeyRemarkRequest) {
    request = &UpdateAccessKeyRemarkRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "UpdateAccessKeyRemark")
    
    
    return
}

func NewUpdateAccessKeyRemarkResponse() (response *UpdateAccessKeyRemarkResponse) {
    response = &UpdateAccessKeyRemarkResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateAccessKeyRemark
// 编辑访问密钥/源IP备注
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) UpdateAccessKeyRemark(request *UpdateAccessKeyRemarkRequest) (response *UpdateAccessKeyRemarkResponse, err error) {
    return c.UpdateAccessKeyRemarkWithContext(context.Background(), request)
}

// UpdateAccessKeyRemark
// 编辑访问密钥/源IP备注
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) UpdateAccessKeyRemarkWithContext(ctx context.Context, request *UpdateAccessKeyRemarkRequest) (response *UpdateAccessKeyRemarkResponse, err error) {
    if request == nil {
        request = NewUpdateAccessKeyRemarkRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "UpdateAccessKeyRemark")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateAccessKeyRemark require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateAccessKeyRemarkResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateAlertStatusListRequest() (request *UpdateAlertStatusListRequest) {
    request = &UpdateAlertStatusListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "UpdateAlertStatusList")
    
    
    return
}

func NewUpdateAlertStatusListResponse() (response *UpdateAlertStatusListResponse) {
    response = &UpdateAlertStatusListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateAlertStatusList
// 批量告警状态处理接口
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) UpdateAlertStatusList(request *UpdateAlertStatusListRequest) (response *UpdateAlertStatusListResponse, err error) {
    return c.UpdateAlertStatusListWithContext(context.Background(), request)
}

// UpdateAlertStatusList
// 批量告警状态处理接口
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) UpdateAlertStatusListWithContext(ctx context.Context, request *UpdateAlertStatusListRequest) (response *UpdateAlertStatusListResponse, err error) {
    if request == nil {
        request = NewUpdateAlertStatusListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "UpdateAlertStatusList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateAlertStatusList require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateAlertStatusListResponse()
    err = c.Send(request, response)
    return
}

func NewVerifyDspmAssetLoginCodeRequest() (request *VerifyDspmAssetLoginCodeRequest) {
    request = &VerifyDspmAssetLoginCodeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "VerifyDspmAssetLoginCode")
    
    
    return
}

func NewVerifyDspmAssetLoginCodeResponse() (response *VerifyDspmAssetLoginCodeResponse) {
    response = &VerifyDspmAssetLoginCodeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// VerifyDspmAssetLoginCode
// 验证Dspm资产登录验证码
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) VerifyDspmAssetLoginCode(request *VerifyDspmAssetLoginCodeRequest) (response *VerifyDspmAssetLoginCodeResponse, err error) {
    return c.VerifyDspmAssetLoginCodeWithContext(context.Background(), request)
}

// VerifyDspmAssetLoginCode
// 验证Dspm资产登录验证码
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) VerifyDspmAssetLoginCodeWithContext(ctx context.Context, request *VerifyDspmAssetLoginCodeRequest) (response *VerifyDspmAssetLoginCodeResponse, err error) {
    if request == nil {
        request = NewVerifyDspmAssetLoginCodeRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "VerifyDspmAssetLoginCode")
    
    if c.GetCredential() == nil {
        return nil, errors.New("VerifyDspmAssetLoginCode require credential")
    }

    request.SetContext(ctx)
    
    response = NewVerifyDspmAssetLoginCodeResponse()
    err = c.Send(request, response)
    return
}
