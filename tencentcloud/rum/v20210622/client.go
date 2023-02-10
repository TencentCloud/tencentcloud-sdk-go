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
    "errors"
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


func NewCreateLogExportRequest() (request *CreateLogExportRequest) {
    request = &CreateLogExportRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "CreateLogExport")
    
    
    return
}

func NewCreateLogExportResponse() (response *CreateLogExportResponse) {
    response = &CreateLogExportResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateLogExport
// 接口请求域名： rum.tencentcloudapi.com 。
//
// 
//
// 本接口用于创建日志下载任务
//
// 
//
// 默认接口请求频率限制：20次/秒。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateLogExport(request *CreateLogExportRequest) (response *CreateLogExportResponse, err error) {
    return c.CreateLogExportWithContext(context.Background(), request)
}

// CreateLogExport
// 接口请求域名： rum.tencentcloudapi.com 。
//
// 
//
// 本接口用于创建日志下载任务
//
// 
//
// 默认接口请求频率限制：20次/秒。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateLogExportWithContext(ctx context.Context, request *CreateLogExportRequest) (response *CreateLogExportResponse, err error) {
    if request == nil {
        request = NewCreateLogExportRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateLogExport require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateLogExportResponse()
    err = c.Send(request, response)
    return
}

func NewCreateOfflineLogConfigRequest() (request *CreateOfflineLogConfigRequest) {
    request = &CreateOfflineLogConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "CreateOfflineLogConfig")
    
    
    return
}

func NewCreateOfflineLogConfigResponse() (response *CreateOfflineLogConfigResponse) {
    response = &CreateOfflineLogConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateOfflineLogConfig
// 创建离线日志监听，对应用户的离线日志将上报
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateOfflineLogConfig(request *CreateOfflineLogConfigRequest) (response *CreateOfflineLogConfigResponse, err error) {
    return c.CreateOfflineLogConfigWithContext(context.Background(), request)
}

// CreateOfflineLogConfig
// 创建离线日志监听，对应用户的离线日志将上报
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateOfflineLogConfigWithContext(ctx context.Context, request *CreateOfflineLogConfigRequest) (response *CreateOfflineLogConfigResponse, err error) {
    if request == nil {
        request = NewCreateOfflineLogConfigRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateOfflineLogConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateOfflineLogConfigResponse()
    err = c.Send(request, response)
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
// 创建 RUM 应用（归属于某个团队）
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
    return c.CreateProjectWithContext(context.Background(), request)
}

// CreateProject
// 创建 RUM 应用（归属于某个团队）
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateProject require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateProjectResponse()
    err = c.Send(request, response)
    return
}

func NewCreateReleaseFileRequest() (request *CreateReleaseFileRequest) {
    request = &CreateReleaseFileRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "CreateReleaseFile")
    
    
    return
}

func NewCreateReleaseFileResponse() (response *CreateReleaseFileResponse) {
    response = &CreateReleaseFileResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateReleaseFile
// 创建对应项目的文件记录
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
//  RESOURCENOTFOUND_NOINSTANCE = "ResourceNotFound.NoInstance"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateReleaseFile(request *CreateReleaseFileRequest) (response *CreateReleaseFileResponse, err error) {
    return c.CreateReleaseFileWithContext(context.Background(), request)
}

// CreateReleaseFile
// 创建对应项目的文件记录
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
//  RESOURCENOTFOUND_NOINSTANCE = "ResourceNotFound.NoInstance"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateReleaseFileWithContext(ctx context.Context, request *CreateReleaseFileRequest) (response *CreateReleaseFileResponse, err error) {
    if request == nil {
        request = NewCreateReleaseFileRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateReleaseFile require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateReleaseFileResponse()
    err = c.Send(request, response)
    return
}

func NewCreateStarProjectRequest() (request *CreateStarProjectRequest) {
    request = &CreateStarProjectRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "CreateStarProject")
    
    
    return
}

func NewCreateStarProjectResponse() (response *CreateStarProjectResponse) {
    response = &CreateStarProjectResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateStarProject
// 个人用户添加星标项目
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
func (c *Client) CreateStarProject(request *CreateStarProjectRequest) (response *CreateStarProjectResponse, err error) {
    return c.CreateStarProjectWithContext(context.Background(), request)
}

// CreateStarProject
// 个人用户添加星标项目
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
func (c *Client) CreateStarProjectWithContext(ctx context.Context, request *CreateStarProjectRequest) (response *CreateStarProjectResponse, err error) {
    if request == nil {
        request = NewCreateStarProjectRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateStarProject require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateStarProjectResponse()
    err = c.Send(request, response)
    return
}

func NewCreateTawInstanceRequest() (request *CreateTawInstanceRequest) {
    request = &CreateTawInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "CreateTawInstance")
    
    
    return
}

func NewCreateTawInstanceResponse() (response *CreateTawInstanceResponse) {
    response = &CreateTawInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateTawInstance
// 创建 RUM 业务系统
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CHARGENOBALANCE = "FailedOperation.ChargeNoBalance"
//  FAILEDOPERATION_CHARGENOPAYRIGHT = "FailedOperation.ChargeNoPayRight"
//  FAILEDOPERATION_CHARGEPARAMINVALID = "FailedOperation.ChargeParamInvalid"
//  FAILEDOPERATION_INFRASTRUCTUREERROR = "FailedOperation.InfrastructureError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateTawInstance(request *CreateTawInstanceRequest) (response *CreateTawInstanceResponse, err error) {
    return c.CreateTawInstanceWithContext(context.Background(), request)
}

// CreateTawInstance
// 创建 RUM 业务系统
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CHARGENOBALANCE = "FailedOperation.ChargeNoBalance"
//  FAILEDOPERATION_CHARGENOPAYRIGHT = "FailedOperation.ChargeNoPayRight"
//  FAILEDOPERATION_CHARGEPARAMINVALID = "FailedOperation.ChargeParamInvalid"
//  FAILEDOPERATION_INFRASTRUCTUREERROR = "FailedOperation.InfrastructureError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateTawInstanceWithContext(ctx context.Context, request *CreateTawInstanceRequest) (response *CreateTawInstanceResponse, err error) {
    if request == nil {
        request = NewCreateTawInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateTawInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateTawInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewCreateWhitelistRequest() (request *CreateWhitelistRequest) {
    request = &CreateWhitelistRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "CreateWhitelist")
    
    
    return
}

func NewCreateWhitelistResponse() (response *CreateWhitelistResponse) {
    response = &CreateWhitelistResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateWhitelist
// 创建白名单
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
func (c *Client) CreateWhitelist(request *CreateWhitelistRequest) (response *CreateWhitelistResponse, err error) {
    return c.CreateWhitelistWithContext(context.Background(), request)
}

// CreateWhitelist
// 创建白名单
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
func (c *Client) CreateWhitelistWithContext(ctx context.Context, request *CreateWhitelistRequest) (response *CreateWhitelistResponse, err error) {
    if request == nil {
        request = NewCreateWhitelistRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateWhitelist require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateWhitelistResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteInstanceRequest() (request *DeleteInstanceRequest) {
    request = &DeleteInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "DeleteInstance")
    
    
    return
}

func NewDeleteInstanceResponse() (response *DeleteInstanceResponse) {
    response = &DeleteInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteInstance
// 删除实例，谨慎操作，不可恢复
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteInstance(request *DeleteInstanceRequest) (response *DeleteInstanceResponse, err error) {
    return c.DeleteInstanceWithContext(context.Background(), request)
}

// DeleteInstance
// 删除实例，谨慎操作，不可恢复
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteInstanceWithContext(ctx context.Context, request *DeleteInstanceRequest) (response *DeleteInstanceResponse, err error) {
    if request == nil {
        request = NewDeleteInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteLogExportRequest() (request *DeleteLogExportRequest) {
    request = &DeleteLogExportRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "DeleteLogExport")
    
    
    return
}

func NewDeleteLogExportResponse() (response *DeleteLogExportResponse) {
    response = &DeleteLogExportResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteLogExport
// 接口请求域名： rum.tencentcloudapi.com 。
//
// 
//
// 本接口用于删除日志下载任务
//
// 
//
// 默认接口请求频率限制：20次/秒。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteLogExport(request *DeleteLogExportRequest) (response *DeleteLogExportResponse, err error) {
    return c.DeleteLogExportWithContext(context.Background(), request)
}

// DeleteLogExport
// 接口请求域名： rum.tencentcloudapi.com 。
//
// 
//
// 本接口用于删除日志下载任务
//
// 
//
// 默认接口请求频率限制：20次/秒。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteLogExportWithContext(ctx context.Context, request *DeleteLogExportRequest) (response *DeleteLogExportResponse, err error) {
    if request == nil {
        request = NewDeleteLogExportRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteLogExport require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteLogExportResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteOfflineLogConfigRequest() (request *DeleteOfflineLogConfigRequest) {
    request = &DeleteOfflineLogConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "DeleteOfflineLogConfig")
    
    
    return
}

func NewDeleteOfflineLogConfigResponse() (response *DeleteOfflineLogConfigResponse) {
    response = &DeleteOfflineLogConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteOfflineLogConfig
// 删除 rum 离线日志监听 - 对应用户的离线日志将不会上报
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
//  RESOURCENOTFOUND_NOINSTANCE = "ResourceNotFound.NoInstance"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteOfflineLogConfig(request *DeleteOfflineLogConfigRequest) (response *DeleteOfflineLogConfigResponse, err error) {
    return c.DeleteOfflineLogConfigWithContext(context.Background(), request)
}

// DeleteOfflineLogConfig
// 删除 rum 离线日志监听 - 对应用户的离线日志将不会上报
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
//  RESOURCENOTFOUND_NOINSTANCE = "ResourceNotFound.NoInstance"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteOfflineLogConfigWithContext(ctx context.Context, request *DeleteOfflineLogConfigRequest) (response *DeleteOfflineLogConfigResponse, err error) {
    if request == nil {
        request = NewDeleteOfflineLogConfigRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteOfflineLogConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteOfflineLogConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteOfflineLogRecordRequest() (request *DeleteOfflineLogRecordRequest) {
    request = &DeleteOfflineLogRecordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "DeleteOfflineLogRecord")
    
    
    return
}

func NewDeleteOfflineLogRecordResponse() (response *DeleteOfflineLogRecordResponse) {
    response = &DeleteOfflineLogRecordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteOfflineLogRecord
// 删除对应的离线日志记录
//
// 可能返回的错误码:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteOfflineLogRecord(request *DeleteOfflineLogRecordRequest) (response *DeleteOfflineLogRecordResponse, err error) {
    return c.DeleteOfflineLogRecordWithContext(context.Background(), request)
}

// DeleteOfflineLogRecord
// 删除对应的离线日志记录
//
// 可能返回的错误码:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteOfflineLogRecordWithContext(ctx context.Context, request *DeleteOfflineLogRecordRequest) (response *DeleteOfflineLogRecordResponse, err error) {
    if request == nil {
        request = NewDeleteOfflineLogRecordRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteOfflineLogRecord require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteOfflineLogRecordResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteProjectRequest() (request *DeleteProjectRequest) {
    request = &DeleteProjectRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "DeleteProject")
    
    
    return
}

func NewDeleteProjectResponse() (response *DeleteProjectResponse) {
    response = &DeleteProjectResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteProject
// 删除给定的 rum 的项目
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteProject(request *DeleteProjectRequest) (response *DeleteProjectResponse, err error) {
    return c.DeleteProjectWithContext(context.Background(), request)
}

// DeleteProject
// 删除给定的 rum 的项目
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteProjectWithContext(ctx context.Context, request *DeleteProjectRequest) (response *DeleteProjectResponse, err error) {
    if request == nil {
        request = NewDeleteProjectRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteProject require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteProjectResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteReleaseFileRequest() (request *DeleteReleaseFileRequest) {
    request = &DeleteReleaseFileRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "DeleteReleaseFile")
    
    
    return
}

func NewDeleteReleaseFileResponse() (response *DeleteReleaseFileResponse) {
    response = &DeleteReleaseFileResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteReleaseFile
// 将对应 sourcemap 文件删除
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
//  RESOURCENOTFOUND_NOINSTANCE = "ResourceNotFound.NoInstance"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteReleaseFile(request *DeleteReleaseFileRequest) (response *DeleteReleaseFileResponse, err error) {
    return c.DeleteReleaseFileWithContext(context.Background(), request)
}

// DeleteReleaseFile
// 将对应 sourcemap 文件删除
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
//  RESOURCENOTFOUND_NOINSTANCE = "ResourceNotFound.NoInstance"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteReleaseFileWithContext(ctx context.Context, request *DeleteReleaseFileRequest) (response *DeleteReleaseFileResponse, err error) {
    if request == nil {
        request = NewDeleteReleaseFileRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteReleaseFile require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteReleaseFileResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteStarProjectRequest() (request *DeleteStarProjectRequest) {
    request = &DeleteStarProjectRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "DeleteStarProject")
    
    
    return
}

func NewDeleteStarProjectResponse() (response *DeleteStarProjectResponse) {
    response = &DeleteStarProjectResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteStarProject
// 删除用户名下的星标项目
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
func (c *Client) DeleteStarProject(request *DeleteStarProjectRequest) (response *DeleteStarProjectResponse, err error) {
    return c.DeleteStarProjectWithContext(context.Background(), request)
}

// DeleteStarProject
// 删除用户名下的星标项目
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
func (c *Client) DeleteStarProjectWithContext(ctx context.Context, request *DeleteStarProjectRequest) (response *DeleteStarProjectResponse, err error) {
    if request == nil {
        request = NewDeleteStarProjectRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteStarProject require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteStarProjectResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteWhitelistRequest() (request *DeleteWhitelistRequest) {
    request = &DeleteWhitelistRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "DeleteWhitelist")
    
    
    return
}

func NewDeleteWhitelistResponse() (response *DeleteWhitelistResponse) {
    response = &DeleteWhitelistResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteWhitelist
// 删除白名单
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
func (c *Client) DeleteWhitelist(request *DeleteWhitelistRequest) (response *DeleteWhitelistResponse, err error) {
    return c.DeleteWhitelistWithContext(context.Background(), request)
}

// DeleteWhitelist
// 删除白名单
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
func (c *Client) DeleteWhitelistWithContext(ctx context.Context, request *DeleteWhitelistRequest) (response *DeleteWhitelistResponse, err error) {
    if request == nil {
        request = NewDeleteWhitelistRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteWhitelist require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteWhitelistResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDataRequest() (request *DescribeDataRequest) {
    request = &DescribeDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "DescribeData")
    
    
    return
}

func NewDescribeDataResponse() (response *DescribeDataResponse) {
    response = &DescribeDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeData
// 转发monitor查询
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
func (c *Client) DescribeData(request *DescribeDataRequest) (response *DescribeDataResponse, err error) {
    return c.DescribeDataWithContext(context.Background(), request)
}

// DescribeData
// 转发monitor查询
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
func (c *Client) DescribeDataWithContext(ctx context.Context, request *DescribeDataRequest) (response *DescribeDataResponse, err error) {
    if request == nil {
        request = NewDescribeDataRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeData require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDataResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDataCustomUrlRequest() (request *DescribeDataCustomUrlRequest) {
    request = &DescribeDataCustomUrlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "DescribeDataCustomUrl")
    
    
    return
}

func NewDescribeDataCustomUrlResponse() (response *DescribeDataCustomUrlResponse) {
    response = &DescribeDataCustomUrlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDataCustomUrl
// 获取DescribeDataCustomUrl信息
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
func (c *Client) DescribeDataCustomUrl(request *DescribeDataCustomUrlRequest) (response *DescribeDataCustomUrlResponse, err error) {
    return c.DescribeDataCustomUrlWithContext(context.Background(), request)
}

// DescribeDataCustomUrl
// 获取DescribeDataCustomUrl信息
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
func (c *Client) DescribeDataCustomUrlWithContext(ctx context.Context, request *DescribeDataCustomUrlRequest) (response *DescribeDataCustomUrlResponse, err error) {
    if request == nil {
        request = NewDescribeDataCustomUrlRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDataCustomUrl require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDataCustomUrlResponse()
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
    return c.DescribeDataEventUrlWithContext(context.Background(), request)
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDataEventUrl require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDataEventUrlResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDataFetchProjectRequest() (request *DescribeDataFetchProjectRequest) {
    request = &DescribeDataFetchProjectRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "DescribeDataFetchProject")
    
    
    return
}

func NewDescribeDataFetchProjectResponse() (response *DescribeDataFetchProjectResponse) {
    response = &DescribeDataFetchProjectResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDataFetchProject
// 获取DescribeDataFetchProject信息。已下线，请使用DescribeDataFetchUrl
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
func (c *Client) DescribeDataFetchProject(request *DescribeDataFetchProjectRequest) (response *DescribeDataFetchProjectResponse, err error) {
    return c.DescribeDataFetchProjectWithContext(context.Background(), request)
}

// DescribeDataFetchProject
// 获取DescribeDataFetchProject信息。已下线，请使用DescribeDataFetchUrl
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
func (c *Client) DescribeDataFetchProjectWithContext(ctx context.Context, request *DescribeDataFetchProjectRequest) (response *DescribeDataFetchProjectResponse, err error) {
    if request == nil {
        request = NewDescribeDataFetchProjectRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDataFetchProject require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDataFetchProjectResponse()
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
    return c.DescribeDataFetchUrlWithContext(context.Background(), request)
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDataFetchUrl require credential")
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
    return c.DescribeDataFetchUrlInfoWithContext(context.Background(), request)
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDataFetchUrlInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDataFetchUrlInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDataLogUrlInfoRequest() (request *DescribeDataLogUrlInfoRequest) {
    request = &DescribeDataLogUrlInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "DescribeDataLogUrlInfo")
    
    
    return
}

func NewDescribeDataLogUrlInfoResponse() (response *DescribeDataLogUrlInfoResponse) {
    response = &DescribeDataLogUrlInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDataLogUrlInfo
// 获取loginfo信息
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
func (c *Client) DescribeDataLogUrlInfo(request *DescribeDataLogUrlInfoRequest) (response *DescribeDataLogUrlInfoResponse, err error) {
    return c.DescribeDataLogUrlInfoWithContext(context.Background(), request)
}

// DescribeDataLogUrlInfo
// 获取loginfo信息
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
func (c *Client) DescribeDataLogUrlInfoWithContext(ctx context.Context, request *DescribeDataLogUrlInfoRequest) (response *DescribeDataLogUrlInfoResponse, err error) {
    if request == nil {
        request = NewDescribeDataLogUrlInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDataLogUrlInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDataLogUrlInfoResponse()
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
    return c.DescribeDataLogUrlStatisticsWithContext(context.Background(), request)
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDataLogUrlStatistics require credential")
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
    return c.DescribeDataPerformancePageWithContext(context.Background(), request)
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDataPerformancePage require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDataPerformancePageResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDataPerformanceProjectRequest() (request *DescribeDataPerformanceProjectRequest) {
    request = &DescribeDataPerformanceProjectRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "DescribeDataPerformanceProject")
    
    
    return
}

func NewDescribeDataPerformanceProjectResponse() (response *DescribeDataPerformanceProjectResponse) {
    response = &DescribeDataPerformanceProjectResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDataPerformanceProject
// 获取PerformanceProject信息
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
func (c *Client) DescribeDataPerformanceProject(request *DescribeDataPerformanceProjectRequest) (response *DescribeDataPerformanceProjectResponse, err error) {
    return c.DescribeDataPerformanceProjectWithContext(context.Background(), request)
}

// DescribeDataPerformanceProject
// 获取PerformanceProject信息
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
func (c *Client) DescribeDataPerformanceProjectWithContext(ctx context.Context, request *DescribeDataPerformanceProjectRequest) (response *DescribeDataPerformanceProjectResponse, err error) {
    if request == nil {
        request = NewDescribeDataPerformanceProjectRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDataPerformanceProject require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDataPerformanceProjectResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDataPvUrlInfoRequest() (request *DescribeDataPvUrlInfoRequest) {
    request = &DescribeDataPvUrlInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "DescribeDataPvUrlInfo")
    
    
    return
}

func NewDescribeDataPvUrlInfoResponse() (response *DescribeDataPvUrlInfoResponse) {
    response = &DescribeDataPvUrlInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDataPvUrlInfo
// 获取PvUrlInfo信息
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
func (c *Client) DescribeDataPvUrlInfo(request *DescribeDataPvUrlInfoRequest) (response *DescribeDataPvUrlInfoResponse, err error) {
    return c.DescribeDataPvUrlInfoWithContext(context.Background(), request)
}

// DescribeDataPvUrlInfo
// 获取PvUrlInfo信息
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
func (c *Client) DescribeDataPvUrlInfoWithContext(ctx context.Context, request *DescribeDataPvUrlInfoRequest) (response *DescribeDataPvUrlInfoResponse, err error) {
    if request == nil {
        request = NewDescribeDataPvUrlInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDataPvUrlInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDataPvUrlInfoResponse()
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
    return c.DescribeDataPvUrlStatisticsWithContext(context.Background(), request)
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDataPvUrlStatistics require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDataPvUrlStatisticsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDataReportCountRequest() (request *DescribeDataReportCountRequest) {
    request = &DescribeDataReportCountRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "DescribeDataReportCount")
    
    
    return
}

func NewDescribeDataReportCountResponse() (response *DescribeDataReportCountResponse) {
    response = &DescribeDataReportCountResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDataReportCount
// 获取项目上报量
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
func (c *Client) DescribeDataReportCount(request *DescribeDataReportCountRequest) (response *DescribeDataReportCountResponse, err error) {
    return c.DescribeDataReportCountWithContext(context.Background(), request)
}

// DescribeDataReportCount
// 获取项目上报量
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
func (c *Client) DescribeDataReportCountWithContext(ctx context.Context, request *DescribeDataReportCountRequest) (response *DescribeDataReportCountResponse, err error) {
    if request == nil {
        request = NewDescribeDataReportCountRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDataReportCount require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDataReportCountResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDataSetUrlStatisticsRequest() (request *DescribeDataSetUrlStatisticsRequest) {
    request = &DescribeDataSetUrlStatisticsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "DescribeDataSetUrlStatistics")
    
    
    return
}

func NewDescribeDataSetUrlStatisticsResponse() (response *DescribeDataSetUrlStatisticsResponse) {
    response = &DescribeDataSetUrlStatisticsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDataSetUrlStatistics
// 获取DescribeDataSetUrlStatistics信息
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
func (c *Client) DescribeDataSetUrlStatistics(request *DescribeDataSetUrlStatisticsRequest) (response *DescribeDataSetUrlStatisticsResponse, err error) {
    return c.DescribeDataSetUrlStatisticsWithContext(context.Background(), request)
}

// DescribeDataSetUrlStatistics
// 获取DescribeDataSetUrlStatistics信息
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
func (c *Client) DescribeDataSetUrlStatisticsWithContext(ctx context.Context, request *DescribeDataSetUrlStatisticsRequest) (response *DescribeDataSetUrlStatisticsResponse, err error) {
    if request == nil {
        request = NewDescribeDataSetUrlStatisticsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDataSetUrlStatistics require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDataSetUrlStatisticsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDataStaticProjectRequest() (request *DescribeDataStaticProjectRequest) {
    request = &DescribeDataStaticProjectRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "DescribeDataStaticProject")
    
    
    return
}

func NewDescribeDataStaticProjectResponse() (response *DescribeDataStaticProjectResponse) {
    response = &DescribeDataStaticProjectResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDataStaticProject
// 获取DescribeDataStaticProject信息
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
func (c *Client) DescribeDataStaticProject(request *DescribeDataStaticProjectRequest) (response *DescribeDataStaticProjectResponse, err error) {
    return c.DescribeDataStaticProjectWithContext(context.Background(), request)
}

// DescribeDataStaticProject
// 获取DescribeDataStaticProject信息
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
func (c *Client) DescribeDataStaticProjectWithContext(ctx context.Context, request *DescribeDataStaticProjectRequest) (response *DescribeDataStaticProjectResponse, err error) {
    if request == nil {
        request = NewDescribeDataStaticProjectRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDataStaticProject require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDataStaticProjectResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDataStaticResourceRequest() (request *DescribeDataStaticResourceRequest) {
    request = &DescribeDataStaticResourceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "DescribeDataStaticResource")
    
    
    return
}

func NewDescribeDataStaticResourceResponse() (response *DescribeDataStaticResourceResponse) {
    response = &DescribeDataStaticResourceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDataStaticResource
// 获取DescribeDataStaticResource信息
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
func (c *Client) DescribeDataStaticResource(request *DescribeDataStaticResourceRequest) (response *DescribeDataStaticResourceResponse, err error) {
    return c.DescribeDataStaticResourceWithContext(context.Background(), request)
}

// DescribeDataStaticResource
// 获取DescribeDataStaticResource信息
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
func (c *Client) DescribeDataStaticResourceWithContext(ctx context.Context, request *DescribeDataStaticResourceRequest) (response *DescribeDataStaticResourceResponse, err error) {
    if request == nil {
        request = NewDescribeDataStaticResourceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDataStaticResource require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDataStaticResourceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDataStaticUrlRequest() (request *DescribeDataStaticUrlRequest) {
    request = &DescribeDataStaticUrlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "DescribeDataStaticUrl")
    
    
    return
}

func NewDescribeDataStaticUrlResponse() (response *DescribeDataStaticUrlResponse) {
    response = &DescribeDataStaticUrlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDataStaticUrl
// 获取DescribeDataStaticUrl信息
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
func (c *Client) DescribeDataStaticUrl(request *DescribeDataStaticUrlRequest) (response *DescribeDataStaticUrlResponse, err error) {
    return c.DescribeDataStaticUrlWithContext(context.Background(), request)
}

// DescribeDataStaticUrl
// 获取DescribeDataStaticUrl信息
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
func (c *Client) DescribeDataStaticUrlWithContext(ctx context.Context, request *DescribeDataStaticUrlRequest) (response *DescribeDataStaticUrlResponse, err error) {
    if request == nil {
        request = NewDescribeDataStaticUrlRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDataStaticUrl require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDataStaticUrlResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDataWebVitalsPageRequest() (request *DescribeDataWebVitalsPageRequest) {
    request = &DescribeDataWebVitalsPageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "DescribeDataWebVitalsPage")
    
    
    return
}

func NewDescribeDataWebVitalsPageResponse() (response *DescribeDataWebVitalsPageResponse) {
    response = &DescribeDataWebVitalsPageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDataWebVitalsPage
// 获取DescribeDataWebVitalsPage信息，用户核心活动信息
//
// 页面加载性能之Web Vitals。性能关键点
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
func (c *Client) DescribeDataWebVitalsPage(request *DescribeDataWebVitalsPageRequest) (response *DescribeDataWebVitalsPageResponse, err error) {
    return c.DescribeDataWebVitalsPageWithContext(context.Background(), request)
}

// DescribeDataWebVitalsPage
// 获取DescribeDataWebVitalsPage信息，用户核心活动信息
//
// 页面加载性能之Web Vitals。性能关键点
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
func (c *Client) DescribeDataWebVitalsPageWithContext(ctx context.Context, request *DescribeDataWebVitalsPageRequest) (response *DescribeDataWebVitalsPageResponse, err error) {
    if request == nil {
        request = NewDescribeDataWebVitalsPageRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDataWebVitalsPage require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDataWebVitalsPageResponse()
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
    return c.DescribeErrorWithContext(context.Background(), request)
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeError require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeErrorResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLogExportsRequest() (request *DescribeLogExportsRequest) {
    request = &DescribeLogExportsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "DescribeLogExports")
    
    
    return
}

func NewDescribeLogExportsResponse() (response *DescribeLogExportsResponse) {
    response = &DescribeLogExportsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeLogExports
// 接口请求域名： rum.tencentcloudapi.com 。
//
// 
//
// 本接口用于获取日志下载任务列表
//
// 
//
// 默认接口请求频率限制：20次/秒
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeLogExports(request *DescribeLogExportsRequest) (response *DescribeLogExportsResponse, err error) {
    return c.DescribeLogExportsWithContext(context.Background(), request)
}

// DescribeLogExports
// 接口请求域名： rum.tencentcloudapi.com 。
//
// 
//
// 本接口用于获取日志下载任务列表
//
// 
//
// 默认接口请求频率限制：20次/秒
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeLogExportsWithContext(ctx context.Context, request *DescribeLogExportsRequest) (response *DescribeLogExportsResponse, err error) {
    if request == nil {
        request = NewDescribeLogExportsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLogExports require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLogExportsResponse()
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
// (已下线，请用DescribeRumLogList)
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
    return c.DescribeLogListWithContext(context.Background(), request)
}

// DescribeLogList
// (已下线，请用DescribeRumLogList)
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLogList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLogListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeOfflineLogConfigsRequest() (request *DescribeOfflineLogConfigsRequest) {
    request = &DescribeOfflineLogConfigsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "DescribeOfflineLogConfigs")
    
    
    return
}

func NewDescribeOfflineLogConfigsResponse() (response *DescribeOfflineLogConfigsResponse) {
    response = &DescribeOfflineLogConfigsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeOfflineLogConfigs
// 获取设置的离线日志监听配置 - 返回设置的用户唯一标识
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeOfflineLogConfigs(request *DescribeOfflineLogConfigsRequest) (response *DescribeOfflineLogConfigsResponse, err error) {
    return c.DescribeOfflineLogConfigsWithContext(context.Background(), request)
}

// DescribeOfflineLogConfigs
// 获取设置的离线日志监听配置 - 返回设置的用户唯一标识
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeOfflineLogConfigsWithContext(ctx context.Context, request *DescribeOfflineLogConfigsRequest) (response *DescribeOfflineLogConfigsResponse, err error) {
    if request == nil {
        request = NewDescribeOfflineLogConfigsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeOfflineLogConfigs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeOfflineLogConfigsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeOfflineLogRecordsRequest() (request *DescribeOfflineLogRecordsRequest) {
    request = &DescribeOfflineLogRecordsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "DescribeOfflineLogRecords")
    
    
    return
}

func NewDescribeOfflineLogRecordsResponse() (response *DescribeOfflineLogRecordsResponse) {
    response = &DescribeOfflineLogRecordsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeOfflineLogRecords
// 获取所有离线日志记录(最多100条)
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeOfflineLogRecords(request *DescribeOfflineLogRecordsRequest) (response *DescribeOfflineLogRecordsResponse, err error) {
    return c.DescribeOfflineLogRecordsWithContext(context.Background(), request)
}

// DescribeOfflineLogRecords
// 获取所有离线日志记录(最多100条)
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeOfflineLogRecordsWithContext(ctx context.Context, request *DescribeOfflineLogRecordsRequest) (response *DescribeOfflineLogRecordsResponse, err error) {
    if request == nil {
        request = NewDescribeOfflineLogRecordsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeOfflineLogRecords require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeOfflineLogRecordsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeOfflineLogsRequest() (request *DescribeOfflineLogsRequest) {
    request = &DescribeOfflineLogsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "DescribeOfflineLogs")
    
    
    return
}

func NewDescribeOfflineLogsResponse() (response *DescribeOfflineLogsResponse) {
    response = &DescribeOfflineLogsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeOfflineLogs
// 获取对应离线日志
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeOfflineLogs(request *DescribeOfflineLogsRequest) (response *DescribeOfflineLogsResponse, err error) {
    return c.DescribeOfflineLogsWithContext(context.Background(), request)
}

// DescribeOfflineLogs
// 获取对应离线日志
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeOfflineLogsWithContext(ctx context.Context, request *DescribeOfflineLogsRequest) (response *DescribeOfflineLogsResponse, err error) {
    if request == nil {
        request = NewDescribeOfflineLogsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeOfflineLogs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeOfflineLogsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeProjectLimitsRequest() (request *DescribeProjectLimitsRequest) {
    request = &DescribeProjectLimitsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "DescribeProjectLimits")
    
    
    return
}

func NewDescribeProjectLimitsResponse() (response *DescribeProjectLimitsResponse) {
    response = &DescribeProjectLimitsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeProjectLimits
// 获取应用上报抽样信息
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
func (c *Client) DescribeProjectLimits(request *DescribeProjectLimitsRequest) (response *DescribeProjectLimitsResponse, err error) {
    return c.DescribeProjectLimitsWithContext(context.Background(), request)
}

// DescribeProjectLimits
// 获取应用上报抽样信息
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
func (c *Client) DescribeProjectLimitsWithContext(ctx context.Context, request *DescribeProjectLimitsRequest) (response *DescribeProjectLimitsResponse, err error) {
    if request == nil {
        request = NewDescribeProjectLimitsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeProjectLimits require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeProjectLimitsResponse()
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
    return c.DescribeProjectsWithContext(context.Background(), request)
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeProjects require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeProjectsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePvListRequest() (request *DescribePvListRequest) {
    request = &DescribePvListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "DescribePvList")
    
    
    return
}

func NewDescribePvListResponse() (response *DescribePvListResponse) {
    response = &DescribePvListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribePvList
// 获取项目下的PV列表
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
func (c *Client) DescribePvList(request *DescribePvListRequest) (response *DescribePvListResponse, err error) {
    return c.DescribePvListWithContext(context.Background(), request)
}

// DescribePvList
// 获取项目下的PV列表
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
func (c *Client) DescribePvListWithContext(ctx context.Context, request *DescribePvListRequest) (response *DescribePvListResponse, err error) {
    if request == nil {
        request = NewDescribePvListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePvList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePvListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeReleaseFileSignRequest() (request *DescribeReleaseFileSignRequest) {
    request = &DescribeReleaseFileSignRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "DescribeReleaseFileSign")
    
    
    return
}

func NewDescribeReleaseFileSignResponse() (response *DescribeReleaseFileSignResponse) {
    response = &DescribeReleaseFileSignResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeReleaseFileSign
// 获取上传文件存储的临时密钥
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
//  RESOURCENOTFOUND_NOINSTANCE = "ResourceNotFound.NoInstance"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeReleaseFileSign(request *DescribeReleaseFileSignRequest) (response *DescribeReleaseFileSignResponse, err error) {
    return c.DescribeReleaseFileSignWithContext(context.Background(), request)
}

// DescribeReleaseFileSign
// 获取上传文件存储的临时密钥
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
//  RESOURCENOTFOUND_NOINSTANCE = "ResourceNotFound.NoInstance"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeReleaseFileSignWithContext(ctx context.Context, request *DescribeReleaseFileSignRequest) (response *DescribeReleaseFileSignResponse, err error) {
    if request == nil {
        request = NewDescribeReleaseFileSignRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeReleaseFileSign require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeReleaseFileSignResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeReleaseFilesRequest() (request *DescribeReleaseFilesRequest) {
    request = &DescribeReleaseFilesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "DescribeReleaseFiles")
    
    
    return
}

func NewDescribeReleaseFilesResponse() (response *DescribeReleaseFilesResponse) {
    response = &DescribeReleaseFilesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeReleaseFiles
// 获取应用对应sourcemap文件列表
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
//  RESOURCENOTFOUND_NOINSTANCE = "ResourceNotFound.NoInstance"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeReleaseFiles(request *DescribeReleaseFilesRequest) (response *DescribeReleaseFilesResponse, err error) {
    return c.DescribeReleaseFilesWithContext(context.Background(), request)
}

// DescribeReleaseFiles
// 获取应用对应sourcemap文件列表
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
//  RESOURCENOTFOUND_NOINSTANCE = "ResourceNotFound.NoInstance"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeReleaseFilesWithContext(ctx context.Context, request *DescribeReleaseFilesRequest) (response *DescribeReleaseFilesResponse, err error) {
    if request == nil {
        request = NewDescribeReleaseFilesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeReleaseFiles require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeReleaseFilesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRumGroupLogRequest() (request *DescribeRumGroupLogRequest) {
    request = &DescribeRumGroupLogRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "DescribeRumGroupLog")
    
    
    return
}

func NewDescribeRumGroupLogResponse() (response *DescribeRumGroupLogResponse) {
    response = &DescribeRumGroupLogResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeRumGroupLog
// 获取项目下的日志聚合信息
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
func (c *Client) DescribeRumGroupLog(request *DescribeRumGroupLogRequest) (response *DescribeRumGroupLogResponse, err error) {
    return c.DescribeRumGroupLogWithContext(context.Background(), request)
}

// DescribeRumGroupLog
// 获取项目下的日志聚合信息
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
func (c *Client) DescribeRumGroupLogWithContext(ctx context.Context, request *DescribeRumGroupLogRequest) (response *DescribeRumGroupLogResponse, err error) {
    if request == nil {
        request = NewDescribeRumGroupLogRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRumGroupLog require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRumGroupLogResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRumLogExportRequest() (request *DescribeRumLogExportRequest) {
    request = &DescribeRumLogExportRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "DescribeRumLogExport")
    
    
    return
}

func NewDescribeRumLogExportResponse() (response *DescribeRumLogExportResponse) {
    response = &DescribeRumLogExportResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeRumLogExport
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
func (c *Client) DescribeRumLogExport(request *DescribeRumLogExportRequest) (response *DescribeRumLogExportResponse, err error) {
    return c.DescribeRumLogExportWithContext(context.Background(), request)
}

// DescribeRumLogExport
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
func (c *Client) DescribeRumLogExportWithContext(ctx context.Context, request *DescribeRumLogExportRequest) (response *DescribeRumLogExportResponse, err error) {
    if request == nil {
        request = NewDescribeRumLogExportRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRumLogExport require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRumLogExportResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRumLogExportsRequest() (request *DescribeRumLogExportsRequest) {
    request = &DescribeRumLogExportsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "DescribeRumLogExports")
    
    
    return
}

func NewDescribeRumLogExportsResponse() (response *DescribeRumLogExportsResponse) {
    response = &DescribeRumLogExportsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeRumLogExports
// 获取项目下的日志导出列表
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
func (c *Client) DescribeRumLogExports(request *DescribeRumLogExportsRequest) (response *DescribeRumLogExportsResponse, err error) {
    return c.DescribeRumLogExportsWithContext(context.Background(), request)
}

// DescribeRumLogExports
// 获取项目下的日志导出列表
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
func (c *Client) DescribeRumLogExportsWithContext(ctx context.Context, request *DescribeRumLogExportsRequest) (response *DescribeRumLogExportsResponse, err error) {
    if request == nil {
        request = NewDescribeRumLogExportsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRumLogExports require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRumLogExportsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRumLogListRequest() (request *DescribeRumLogListRequest) {
    request = &DescribeRumLogListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "DescribeRumLogList")
    
    
    return
}

func NewDescribeRumLogListResponse() (response *DescribeRumLogListResponse) {
    response = &DescribeRumLogListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeRumLogList
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
func (c *Client) DescribeRumLogList(request *DescribeRumLogListRequest) (response *DescribeRumLogListResponse, err error) {
    return c.DescribeRumLogListWithContext(context.Background(), request)
}

// DescribeRumLogList
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
func (c *Client) DescribeRumLogListWithContext(ctx context.Context, request *DescribeRumLogListRequest) (response *DescribeRumLogListResponse, err error) {
    if request == nil {
        request = NewDescribeRumLogListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRumLogList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRumLogListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRumStatsLogListRequest() (request *DescribeRumStatsLogListRequest) {
    request = &DescribeRumStatsLogListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "DescribeRumStatsLogList")
    
    
    return
}

func NewDescribeRumStatsLogListResponse() (response *DescribeRumStatsLogListResponse) {
    response = &DescribeRumStatsLogListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeRumStatsLogList
// 获取项目下的日志列表，分钟级
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
func (c *Client) DescribeRumStatsLogList(request *DescribeRumStatsLogListRequest) (response *DescribeRumStatsLogListResponse, err error) {
    return c.DescribeRumStatsLogListWithContext(context.Background(), request)
}

// DescribeRumStatsLogList
// 获取项目下的日志列表，分钟级
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
func (c *Client) DescribeRumStatsLogListWithContext(ctx context.Context, request *DescribeRumStatsLogListRequest) (response *DescribeRumStatsLogListResponse, err error) {
    if request == nil {
        request = NewDescribeRumStatsLogListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRumStatsLogList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRumStatsLogListResponse()
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
    return c.DescribeScoresWithContext(context.Background(), request)
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeScores require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeScoresResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTawAreasRequest() (request *DescribeTawAreasRequest) {
    request = &DescribeTawAreasRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "DescribeTawAreas")
    
    
    return
}

func NewDescribeTawAreasResponse() (response *DescribeTawAreasResponse) {
    response = &DescribeTawAreasResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeTawAreas
// 查询片区信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeTawAreas(request *DescribeTawAreasRequest) (response *DescribeTawAreasResponse, err error) {
    return c.DescribeTawAreasWithContext(context.Background(), request)
}

// DescribeTawAreas
// 查询片区信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeTawAreasWithContext(ctx context.Context, request *DescribeTawAreasRequest) (response *DescribeTawAreasResponse, err error) {
    if request == nil {
        request = NewDescribeTawAreasRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTawAreas require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTawAreasResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTawInstancesRequest() (request *DescribeTawInstancesRequest) {
    request = &DescribeTawInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "DescribeTawInstances")
    
    
    return
}

func NewDescribeTawInstancesResponse() (response *DescribeTawInstancesResponse) {
    response = &DescribeTawInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeTawInstances
// 查询实例信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeTawInstances(request *DescribeTawInstancesRequest) (response *DescribeTawInstancesResponse, err error) {
    return c.DescribeTawInstancesWithContext(context.Background(), request)
}

// DescribeTawInstances
// 查询实例信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeTawInstancesWithContext(ctx context.Context, request *DescribeTawInstancesRequest) (response *DescribeTawInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeTawInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTawInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTawInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUvListRequest() (request *DescribeUvListRequest) {
    request = &DescribeUvListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "DescribeUvList")
    
    
    return
}

func NewDescribeUvListResponse() (response *DescribeUvListResponse) {
    response = &DescribeUvListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeUvList
// 获取项目下的UV列表
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
func (c *Client) DescribeUvList(request *DescribeUvListRequest) (response *DescribeUvListResponse, err error) {
    return c.DescribeUvListWithContext(context.Background(), request)
}

// DescribeUvList
// 获取项目下的UV列表
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
func (c *Client) DescribeUvListWithContext(ctx context.Context, request *DescribeUvListRequest) (response *DescribeUvListResponse, err error) {
    if request == nil {
        request = NewDescribeUvListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUvList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeUvListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeWhitelistsRequest() (request *DescribeWhitelistsRequest) {
    request = &DescribeWhitelistsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "DescribeWhitelists")
    
    
    return
}

func NewDescribeWhitelistsResponse() (response *DescribeWhitelistsResponse) {
    response = &DescribeWhitelistsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeWhitelists
// 获取白名单列表
//
// 可能返回的错误码:
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
func (c *Client) DescribeWhitelists(request *DescribeWhitelistsRequest) (response *DescribeWhitelistsResponse, err error) {
    return c.DescribeWhitelistsWithContext(context.Background(), request)
}

// DescribeWhitelists
// 获取白名单列表
//
// 可能返回的错误码:
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
func (c *Client) DescribeWhitelistsWithContext(ctx context.Context, request *DescribeWhitelistsRequest) (response *DescribeWhitelistsResponse, err error) {
    if request == nil {
        request = NewDescribeWhitelistsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeWhitelists require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeWhitelistsResponse()
    err = c.Send(request, response)
    return
}

func NewModifyInstanceRequest() (request *ModifyInstanceRequest) {
    request = &ModifyInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "ModifyInstance")
    
    
    return
}

func NewModifyInstanceResponse() (response *ModifyInstanceResponse) {
    response = &ModifyInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyInstance
// 修改 RUM 业务系统
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyInstance(request *ModifyInstanceRequest) (response *ModifyInstanceResponse, err error) {
    return c.ModifyInstanceWithContext(context.Background(), request)
}

// ModifyInstance
// 修改 RUM 业务系统
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyInstanceWithContext(ctx context.Context, request *ModifyInstanceRequest) (response *ModifyInstanceResponse, err error) {
    if request == nil {
        request = NewModifyInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewModifyProjectRequest() (request *ModifyProjectRequest) {
    request = &ModifyProjectRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "ModifyProject")
    
    
    return
}

func NewModifyProjectResponse() (response *ModifyProjectResponse) {
    response = &ModifyProjectResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyProject
// 修改 RUM 应用信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyProject(request *ModifyProjectRequest) (response *ModifyProjectResponse, err error) {
    return c.ModifyProjectWithContext(context.Background(), request)
}

// ModifyProject
// 修改 RUM 应用信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyProjectWithContext(ctx context.Context, request *ModifyProjectRequest) (response *ModifyProjectResponse, err error) {
    if request == nil {
        request = NewModifyProjectRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyProject require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyProjectResponse()
    err = c.Send(request, response)
    return
}

func NewModifyProjectLimitRequest() (request *ModifyProjectLimitRequest) {
    request = &ModifyProjectLimitRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "ModifyProjectLimit")
    
    
    return
}

func NewModifyProjectLimitResponse() (response *ModifyProjectLimitResponse) {
    response = &ModifyProjectLimitResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyProjectLimit
// 新增修改限流
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
func (c *Client) ModifyProjectLimit(request *ModifyProjectLimitRequest) (response *ModifyProjectLimitResponse, err error) {
    return c.ModifyProjectLimitWithContext(context.Background(), request)
}

// ModifyProjectLimit
// 新增修改限流
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
func (c *Client) ModifyProjectLimitWithContext(ctx context.Context, request *ModifyProjectLimitRequest) (response *ModifyProjectLimitResponse, err error) {
    if request == nil {
        request = NewModifyProjectLimitRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyProjectLimit require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyProjectLimitResponse()
    err = c.Send(request, response)
    return
}

func NewResumeInstanceRequest() (request *ResumeInstanceRequest) {
    request = &ResumeInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "ResumeInstance")
    
    
    return
}

func NewResumeInstanceResponse() (response *ResumeInstanceResponse) {
    response = &ResumeInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ResumeInstance
// 恢复 RUM 业务系统，恢复后，用户可以正常使用和上报数据
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ResumeInstance(request *ResumeInstanceRequest) (response *ResumeInstanceResponse, err error) {
    return c.ResumeInstanceWithContext(context.Background(), request)
}

// ResumeInstance
// 恢复 RUM 业务系统，恢复后，用户可以正常使用和上报数据
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ResumeInstanceWithContext(ctx context.Context, request *ResumeInstanceRequest) (response *ResumeInstanceResponse, err error) {
    if request == nil {
        request = NewResumeInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ResumeInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewResumeInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewResumeProjectRequest() (request *ResumeProjectRequest) {
    request = &ResumeProjectRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "ResumeProject")
    
    
    return
}

func NewResumeProjectResponse() (response *ResumeProjectResponse) {
    response = &ResumeProjectResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ResumeProject
// 恢复应用使用与上报数据
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ResumeProject(request *ResumeProjectRequest) (response *ResumeProjectResponse, err error) {
    return c.ResumeProjectWithContext(context.Background(), request)
}

// ResumeProject
// 恢复应用使用与上报数据
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ResumeProjectWithContext(ctx context.Context, request *ResumeProjectRequest) (response *ResumeProjectResponse, err error) {
    if request == nil {
        request = NewResumeProjectRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ResumeProject require credential")
    }

    request.SetContext(ctx)
    
    response = NewResumeProjectResponse()
    err = c.Send(request, response)
    return
}

func NewStopInstanceRequest() (request *StopInstanceRequest) {
    request = &StopInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "StopInstance")
    
    
    return
}

func NewStopInstanceResponse() (response *StopInstanceResponse) {
    response = &StopInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// StopInstance
// 停止实例
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) StopInstance(request *StopInstanceRequest) (response *StopInstanceResponse, err error) {
    return c.StopInstanceWithContext(context.Background(), request)
}

// StopInstance
// 停止实例
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) StopInstanceWithContext(ctx context.Context, request *StopInstanceRequest) (response *StopInstanceResponse, err error) {
    if request == nil {
        request = NewStopInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("StopInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewStopInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewStopProjectRequest() (request *StopProjectRequest) {
    request = &StopProjectRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "StopProject")
    
    
    return
}

func NewStopProjectResponse() (response *StopProjectResponse) {
    response = &StopProjectResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// StopProject
// 停止项目使用与上报数据
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) StopProject(request *StopProjectRequest) (response *StopProjectResponse, err error) {
    return c.StopProjectWithContext(context.Background(), request)
}

// StopProject
// 停止项目使用与上报数据
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) StopProjectWithContext(ctx context.Context, request *StopProjectRequest) (response *StopProjectResponse, err error) {
    if request == nil {
        request = NewStopProjectRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("StopProject require credential")
    }

    request.SetContext(ctx)
    
    response = NewStopProjectResponse()
    err = c.Send(request, response)
    return
}
