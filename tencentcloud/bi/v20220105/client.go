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

package v20220105

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2022-01-05"

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


func NewApplyEmbedIntervalRequest() (request *ApplyEmbedIntervalRequest) {
    request = &ApplyEmbedIntervalRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bi", APIVersion, "ApplyEmbedInterval")
    
    
    return
}

func NewApplyEmbedIntervalResponse() (response *ApplyEmbedIntervalResponse) {
    response = &ApplyEmbedIntervalResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ApplyEmbedInterval
// 申请延长Token可用时间接口-强鉴权
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNAL = "InternalError.Internal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_EMBED = "InvalidParameter.Embed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_MISSINGPARAM = "MissingParameter.MissingParam"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_AUTHORIZE = "UnauthorizedOperation.Authorize"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_BIERROR = "UnsupportedOperation.BIError"
func (c *Client) ApplyEmbedInterval(request *ApplyEmbedIntervalRequest) (response *ApplyEmbedIntervalResponse, err error) {
    return c.ApplyEmbedIntervalWithContext(context.Background(), request)
}

// ApplyEmbedInterval
// 申请延长Token可用时间接口-强鉴权
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNAL = "InternalError.Internal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_EMBED = "InvalidParameter.Embed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_MISSINGPARAM = "MissingParameter.MissingParam"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_AUTHORIZE = "UnauthorizedOperation.Authorize"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_BIERROR = "UnsupportedOperation.BIError"
func (c *Client) ApplyEmbedIntervalWithContext(ctx context.Context, request *ApplyEmbedIntervalRequest) (response *ApplyEmbedIntervalResponse, err error) {
    if request == nil {
        request = NewApplyEmbedIntervalRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bi", APIVersion, "ApplyEmbedInterval")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ApplyEmbedInterval require credential")
    }

    request.SetContext(ctx)
    
    response = NewApplyEmbedIntervalResponse()
    err = c.Send(request, response)
    return
}

func NewClearEmbedTokenRequest() (request *ClearEmbedTokenRequest) {
    request = &ClearEmbedTokenRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bi", APIVersion, "ClearEmbedToken")
    
    
    return
}

func NewClearEmbedTokenResponse() (response *ClearEmbedTokenResponse) {
    response = &ClearEmbedTokenResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ClearEmbedToken
// 强鉴权token 清理，只有企业管理员才能调用该接口
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNAL = "InternalError.Internal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_EMBED = "InvalidParameter.Embed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_MISSINGPARAM = "MissingParameter.MissingParam"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_AUTHORIZE = "UnauthorizedOperation.Authorize"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_BIERROR = "UnsupportedOperation.BIError"
func (c *Client) ClearEmbedToken(request *ClearEmbedTokenRequest) (response *ClearEmbedTokenResponse, err error) {
    return c.ClearEmbedTokenWithContext(context.Background(), request)
}

// ClearEmbedToken
// 强鉴权token 清理，只有企业管理员才能调用该接口
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNAL = "InternalError.Internal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_EMBED = "InvalidParameter.Embed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_MISSINGPARAM = "MissingParameter.MissingParam"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_AUTHORIZE = "UnauthorizedOperation.Authorize"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_BIERROR = "UnsupportedOperation.BIError"
func (c *Client) ClearEmbedTokenWithContext(ctx context.Context, request *ClearEmbedTokenRequest) (response *ClearEmbedTokenResponse, err error) {
    if request == nil {
        request = NewClearEmbedTokenRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bi", APIVersion, "ClearEmbedToken")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ClearEmbedToken require credential")
    }

    request.SetContext(ctx)
    
    response = NewClearEmbedTokenResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAuthApiKeyRequest() (request *CreateAuthApiKeyRequest) {
    request = &CreateAuthApiKeyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bi", APIVersion, "CreateAuthApiKey")
    
    
    return
}

func NewCreateAuthApiKeyResponse() (response *CreateAuthApiKeyResponse) {
    response = &CreateAuthApiKeyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateAuthApiKey
// 创建ApiKey
//
// 可能返回的错误码:
//  INTERNALERROR_INTERNAL = "InternalError.Internal"
//  MISSINGPARAMETER_MISSINGPARAM = "MissingParameter.MissingParam"
//  UNAUTHORIZEDOPERATION_AUTHORIZE = "UnauthorizedOperation.Authorize"
//  UNAUTHORIZEDOPERATION_INACTIVE = "UnauthorizedOperation.Inactive"
//  UNSUPPORTEDOPERATION_BIERROR = "UnsupportedOperation.BIError"
func (c *Client) CreateAuthApiKey(request *CreateAuthApiKeyRequest) (response *CreateAuthApiKeyResponse, err error) {
    return c.CreateAuthApiKeyWithContext(context.Background(), request)
}

// CreateAuthApiKey
// 创建ApiKey
//
// 可能返回的错误码:
//  INTERNALERROR_INTERNAL = "InternalError.Internal"
//  MISSINGPARAMETER_MISSINGPARAM = "MissingParameter.MissingParam"
//  UNAUTHORIZEDOPERATION_AUTHORIZE = "UnauthorizedOperation.Authorize"
//  UNAUTHORIZEDOPERATION_INACTIVE = "UnauthorizedOperation.Inactive"
//  UNSUPPORTEDOPERATION_BIERROR = "UnsupportedOperation.BIError"
func (c *Client) CreateAuthApiKeyWithContext(ctx context.Context, request *CreateAuthApiKeyRequest) (response *CreateAuthApiKeyResponse, err error) {
    if request == nil {
        request = NewCreateAuthApiKeyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bi", APIVersion, "CreateAuthApiKey")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAuthApiKey require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAuthApiKeyResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDataTableRequest() (request *CreateDataTableRequest) {
    request = &CreateDataTableRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bi", APIVersion, "CreateDataTable")
    
    
    return
}

func NewCreateDataTableResponse() (response *CreateDataTableResponse) {
    response = &CreateDataTableResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateDataTable
// 添加数据表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNAL = "InternalError.Internal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_MISSINGPARAM = "MissingParameter.MissingParam"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_AUTHORIZE = "UnauthorizedOperation.Authorize"
//  UNAUTHORIZEDOPERATION_USERNOTEXIST = "UnauthorizedOperation.UserNotExist"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_BIERROR = "UnsupportedOperation.BIError"
func (c *Client) CreateDataTable(request *CreateDataTableRequest) (response *CreateDataTableResponse, err error) {
    return c.CreateDataTableWithContext(context.Background(), request)
}

// CreateDataTable
// 添加数据表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNAL = "InternalError.Internal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_MISSINGPARAM = "MissingParameter.MissingParam"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_AUTHORIZE = "UnauthorizedOperation.Authorize"
//  UNAUTHORIZEDOPERATION_USERNOTEXIST = "UnauthorizedOperation.UserNotExist"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_BIERROR = "UnsupportedOperation.BIError"
func (c *Client) CreateDataTableWithContext(ctx context.Context, request *CreateDataTableRequest) (response *CreateDataTableResponse, err error) {
    if request == nil {
        request = NewCreateDataTableRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bi", APIVersion, "CreateDataTable")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateDataTable require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateDataTableResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDatasourceRequest() (request *CreateDatasourceRequest) {
    request = &CreateDatasourceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bi", APIVersion, "CreateDatasource")
    
    
    return
}

func NewCreateDatasourceResponse() (response *CreateDatasourceResponse) {
    response = &CreateDatasourceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateDatasource
// 创建数据源
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNAL = "InternalError.Internal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_MISSINGPARAM = "MissingParameter.MissingParam"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_AUTHORIZE = "UnauthorizedOperation.Authorize"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_BIERROR = "UnsupportedOperation.BIError"
func (c *Client) CreateDatasource(request *CreateDatasourceRequest) (response *CreateDatasourceResponse, err error) {
    return c.CreateDatasourceWithContext(context.Background(), request)
}

// CreateDatasource
// 创建数据源
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNAL = "InternalError.Internal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_MISSINGPARAM = "MissingParameter.MissingParam"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_AUTHORIZE = "UnauthorizedOperation.Authorize"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_BIERROR = "UnsupportedOperation.BIError"
func (c *Client) CreateDatasourceWithContext(ctx context.Context, request *CreateDatasourceRequest) (response *CreateDatasourceResponse, err error) {
    if request == nil {
        request = NewCreateDatasourceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bi", APIVersion, "CreateDatasource")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateDatasource require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateDatasourceResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDatasourceCloudRequest() (request *CreateDatasourceCloudRequest) {
    request = &CreateDatasourceCloudRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bi", APIVersion, "CreateDatasourceCloud")
    
    
    return
}

func NewCreateDatasourceCloudResponse() (response *CreateDatasourceCloudResponse) {
    response = &CreateDatasourceCloudResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateDatasourceCloud
// 创建云数据库
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNAL = "InternalError.Internal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_MISSINGPARAM = "MissingParameter.MissingParam"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_AUTHORIZE = "UnauthorizedOperation.Authorize"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_BIERROR = "UnsupportedOperation.BIError"
func (c *Client) CreateDatasourceCloud(request *CreateDatasourceCloudRequest) (response *CreateDatasourceCloudResponse, err error) {
    return c.CreateDatasourceCloudWithContext(context.Background(), request)
}

// CreateDatasourceCloud
// 创建云数据库
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNAL = "InternalError.Internal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_MISSINGPARAM = "MissingParameter.MissingParam"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_AUTHORIZE = "UnauthorizedOperation.Authorize"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_BIERROR = "UnsupportedOperation.BIError"
func (c *Client) CreateDatasourceCloudWithContext(ctx context.Context, request *CreateDatasourceCloudRequest) (response *CreateDatasourceCloudResponse, err error) {
    if request == nil {
        request = NewCreateDatasourceCloudRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bi", APIVersion, "CreateDatasourceCloud")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateDatasourceCloud require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateDatasourceCloudResponse()
    err = c.Send(request, response)
    return
}

func NewCreateEmbedTokenRequest() (request *CreateEmbedTokenRequest) {
    request = &CreateEmbedTokenRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bi", APIVersion, "CreateEmbedToken")
    
    
    return
}

func NewCreateEmbedTokenResponse() (response *CreateEmbedTokenResponse) {
    response = &CreateEmbedTokenResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateEmbedToken
// 创建嵌出报表-强鉴权
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNAL = "InternalError.Internal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_EMBED = "InvalidParameter.Embed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_MISSINGPARAM = "MissingParameter.MissingParam"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_AUTHORIZE = "UnauthorizedOperation.Authorize"
//  UNAUTHORIZEDOPERATION_USERNOTEXIST = "UnauthorizedOperation.UserNotExist"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_BIERROR = "UnsupportedOperation.BIError"
func (c *Client) CreateEmbedToken(request *CreateEmbedTokenRequest) (response *CreateEmbedTokenResponse, err error) {
    return c.CreateEmbedTokenWithContext(context.Background(), request)
}

// CreateEmbedToken
// 创建嵌出报表-强鉴权
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNAL = "InternalError.Internal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_EMBED = "InvalidParameter.Embed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_MISSINGPARAM = "MissingParameter.MissingParam"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_AUTHORIZE = "UnauthorizedOperation.Authorize"
//  UNAUTHORIZEDOPERATION_USERNOTEXIST = "UnauthorizedOperation.UserNotExist"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_BIERROR = "UnsupportedOperation.BIError"
func (c *Client) CreateEmbedTokenWithContext(ctx context.Context, request *CreateEmbedTokenRequest) (response *CreateEmbedTokenResponse, err error) {
    if request == nil {
        request = NewCreateEmbedTokenRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bi", APIVersion, "CreateEmbedToken")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateEmbedToken require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateEmbedTokenResponse()
    err = c.Send(request, response)
    return
}

func NewCreatePermissionRanksRequest() (request *CreatePermissionRanksRequest) {
    request = &CreatePermissionRanksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bi", APIVersion, "CreatePermissionRanks")
    
    
    return
}

func NewCreatePermissionRanksResponse() (response *CreatePermissionRanksResponse) {
    response = &CreatePermissionRanksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreatePermissionRanks
// 创建行列权限
//
// 可能返回的错误码:
//  INTERNALERROR_INTERNAL = "InternalError.Internal"
//  MISSINGPARAMETER_MISSINGPARAM = "MissingParameter.MissingParam"
//  UNAUTHORIZEDOPERATION_AUTHORIZE = "UnauthorizedOperation.Authorize"
//  UNSUPPORTEDOPERATION_BIERROR = "UnsupportedOperation.BIError"
func (c *Client) CreatePermissionRanks(request *CreatePermissionRanksRequest) (response *CreatePermissionRanksResponse, err error) {
    return c.CreatePermissionRanksWithContext(context.Background(), request)
}

// CreatePermissionRanks
// 创建行列权限
//
// 可能返回的错误码:
//  INTERNALERROR_INTERNAL = "InternalError.Internal"
//  MISSINGPARAMETER_MISSINGPARAM = "MissingParameter.MissingParam"
//  UNAUTHORIZEDOPERATION_AUTHORIZE = "UnauthorizedOperation.Authorize"
//  UNSUPPORTEDOPERATION_BIERROR = "UnsupportedOperation.BIError"
func (c *Client) CreatePermissionRanksWithContext(ctx context.Context, request *CreatePermissionRanksRequest) (response *CreatePermissionRanksResponse, err error) {
    if request == nil {
        request = NewCreatePermissionRanksRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bi", APIVersion, "CreatePermissionRanks")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreatePermissionRanks require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreatePermissionRanksResponse()
    err = c.Send(request, response)
    return
}

func NewCreateProjectRequest() (request *CreateProjectRequest) {
    request = &CreateProjectRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bi", APIVersion, "CreateProject")
    
    
    return
}

func NewCreateProjectResponse() (response *CreateProjectResponse) {
    response = &CreateProjectResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateProject
// 创建项目
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNAL = "InternalError.Internal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_MISSINGPARAM = "MissingParameter.MissingParam"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_AUTHORIZE = "UnauthorizedOperation.Authorize"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateProject(request *CreateProjectRequest) (response *CreateProjectResponse, err error) {
    return c.CreateProjectWithContext(context.Background(), request)
}

// CreateProject
// 创建项目
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNAL = "InternalError.Internal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_MISSINGPARAM = "MissingParameter.MissingParam"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_AUTHORIZE = "UnauthorizedOperation.Authorize"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateProjectWithContext(ctx context.Context, request *CreateProjectRequest) (response *CreateProjectResponse, err error) {
    if request == nil {
        request = NewCreateProjectRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bi", APIVersion, "CreateProject")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateProject require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateProjectResponse()
    err = c.Send(request, response)
    return
}

func NewCreateUserGroupRequest() (request *CreateUserGroupRequest) {
    request = &CreateUserGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bi", APIVersion, "CreateUserGroup")
    
    
    return
}

func NewCreateUserGroupResponse() (response *CreateUserGroupResponse) {
    response = &CreateUserGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateUserGroup
// CreateUserGroup
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNAL = "InternalError.Internal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_EMBED = "InvalidParameter.Embed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_MISSINGPARAM = "MissingParameter.MissingParam"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_AUTHORIZE = "UnauthorizedOperation.Authorize"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_BIERROR = "UnsupportedOperation.BIError"
func (c *Client) CreateUserGroup(request *CreateUserGroupRequest) (response *CreateUserGroupResponse, err error) {
    return c.CreateUserGroupWithContext(context.Background(), request)
}

// CreateUserGroup
// CreateUserGroup
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNAL = "InternalError.Internal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_EMBED = "InvalidParameter.Embed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_MISSINGPARAM = "MissingParameter.MissingParam"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_AUTHORIZE = "UnauthorizedOperation.Authorize"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_BIERROR = "UnsupportedOperation.BIError"
func (c *Client) CreateUserGroupWithContext(ctx context.Context, request *CreateUserGroupRequest) (response *CreateUserGroupResponse, err error) {
    if request == nil {
        request = NewCreateUserGroupRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bi", APIVersion, "CreateUserGroup")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateUserGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateUserGroupResponse()
    err = c.Send(request, response)
    return
}

func NewCreateUserGroupMemberRequest() (request *CreateUserGroupMemberRequest) {
    request = &CreateUserGroupMemberRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bi", APIVersion, "CreateUserGroupMember")
    
    
    return
}

func NewCreateUserGroupMemberResponse() (response *CreateUserGroupMemberResponse) {
    response = &CreateUserGroupMemberResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateUserGroupMember
// CreateUserGroupMember
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNAL = "InternalError.Internal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_EMBED = "InvalidParameter.Embed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_MISSINGPARAM = "MissingParameter.MissingParam"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_AUTHORIZE = "UnauthorizedOperation.Authorize"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_BIERROR = "UnsupportedOperation.BIError"
func (c *Client) CreateUserGroupMember(request *CreateUserGroupMemberRequest) (response *CreateUserGroupMemberResponse, err error) {
    return c.CreateUserGroupMemberWithContext(context.Background(), request)
}

// CreateUserGroupMember
// CreateUserGroupMember
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNAL = "InternalError.Internal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_EMBED = "InvalidParameter.Embed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_MISSINGPARAM = "MissingParameter.MissingParam"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_AUTHORIZE = "UnauthorizedOperation.Authorize"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_BIERROR = "UnsupportedOperation.BIError"
func (c *Client) CreateUserGroupMemberWithContext(ctx context.Context, request *CreateUserGroupMemberRequest) (response *CreateUserGroupMemberResponse, err error) {
    if request == nil {
        request = NewCreateUserGroupMemberRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bi", APIVersion, "CreateUserGroupMember")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateUserGroupMember require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateUserGroupMemberResponse()
    err = c.Send(request, response)
    return
}

func NewCreateUserRoleRequest() (request *CreateUserRoleRequest) {
    request = &CreateUserRoleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bi", APIVersion, "CreateUserRole")
    
    
    return
}

func NewCreateUserRoleResponse() (response *CreateUserRoleResponse) {
    response = &CreateUserRoleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateUserRole
// 创建用户角色
//
// 可能返回的错误码:
//  INTERNALERROR_INTERNAL = "InternalError.Internal"
//  MISSINGPARAMETER_MISSINGPARAM = "MissingParameter.MissingParam"
//  UNAUTHORIZEDOPERATION_AUTHORIZE = "UnauthorizedOperation.Authorize"
//  UNAUTHORIZEDOPERATION_USERNOTEXIST = "UnauthorizedOperation.UserNotExist"
func (c *Client) CreateUserRole(request *CreateUserRoleRequest) (response *CreateUserRoleResponse, err error) {
    return c.CreateUserRoleWithContext(context.Background(), request)
}

// CreateUserRole
// 创建用户角色
//
// 可能返回的错误码:
//  INTERNALERROR_INTERNAL = "InternalError.Internal"
//  MISSINGPARAMETER_MISSINGPARAM = "MissingParameter.MissingParam"
//  UNAUTHORIZEDOPERATION_AUTHORIZE = "UnauthorizedOperation.Authorize"
//  UNAUTHORIZEDOPERATION_USERNOTEXIST = "UnauthorizedOperation.UserNotExist"
func (c *Client) CreateUserRoleWithContext(ctx context.Context, request *CreateUserRoleRequest) (response *CreateUserRoleResponse, err error) {
    if request == nil {
        request = NewCreateUserRoleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bi", APIVersion, "CreateUserRole")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateUserRole require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateUserRoleResponse()
    err = c.Send(request, response)
    return
}

func NewCreateUserRoleProjectRequest() (request *CreateUserRoleProjectRequest) {
    request = &CreateUserRoleProjectRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bi", APIVersion, "CreateUserRoleProject")
    
    
    return
}

func NewCreateUserRoleProjectResponse() (response *CreateUserRoleProjectResponse) {
    response = &CreateUserRoleProjectResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateUserRoleProject
// 项目内-创建用户角色
//
// 可能返回的错误码:
//  INTERNALERROR_INTERNAL = "InternalError.Internal"
//  MISSINGPARAMETER_MISSINGPARAM = "MissingParameter.MissingParam"
//  UNAUTHORIZEDOPERATION_AUTHORIZE = "UnauthorizedOperation.Authorize"
//  UNAUTHORIZEDOPERATION_INACTIVE = "UnauthorizedOperation.Inactive"
func (c *Client) CreateUserRoleProject(request *CreateUserRoleProjectRequest) (response *CreateUserRoleProjectResponse, err error) {
    return c.CreateUserRoleProjectWithContext(context.Background(), request)
}

// CreateUserRoleProject
// 项目内-创建用户角色
//
// 可能返回的错误码:
//  INTERNALERROR_INTERNAL = "InternalError.Internal"
//  MISSINGPARAMETER_MISSINGPARAM = "MissingParameter.MissingParam"
//  UNAUTHORIZEDOPERATION_AUTHORIZE = "UnauthorizedOperation.Authorize"
//  UNAUTHORIZEDOPERATION_INACTIVE = "UnauthorizedOperation.Inactive"
func (c *Client) CreateUserRoleProjectWithContext(ctx context.Context, request *CreateUserRoleProjectRequest) (response *CreateUserRoleProjectResponse, err error) {
    if request == nil {
        request = NewCreateUserRoleProjectRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bi", APIVersion, "CreateUserRoleProject")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateUserRoleProject require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateUserRoleProjectResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAuthApiKeyRequest() (request *DeleteAuthApiKeyRequest) {
    request = &DeleteAuthApiKeyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bi", APIVersion, "DeleteAuthApiKey")
    
    
    return
}

func NewDeleteAuthApiKeyResponse() (response *DeleteAuthApiKeyResponse) {
    response = &DeleteAuthApiKeyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteAuthApiKey
// 删除ApiKey
//
// 可能返回的错误码:
//  INTERNALERROR_INTERNAL = "InternalError.Internal"
//  MISSINGPARAMETER_MISSINGPARAM = "MissingParameter.MissingParam"
//  UNAUTHORIZEDOPERATION_AUTHORIZE = "UnauthorizedOperation.Authorize"
//  UNAUTHORIZEDOPERATION_INACTIVE = "UnauthorizedOperation.Inactive"
//  UNSUPPORTEDOPERATION_BIERROR = "UnsupportedOperation.BIError"
func (c *Client) DeleteAuthApiKey(request *DeleteAuthApiKeyRequest) (response *DeleteAuthApiKeyResponse, err error) {
    return c.DeleteAuthApiKeyWithContext(context.Background(), request)
}

// DeleteAuthApiKey
// 删除ApiKey
//
// 可能返回的错误码:
//  INTERNALERROR_INTERNAL = "InternalError.Internal"
//  MISSINGPARAMETER_MISSINGPARAM = "MissingParameter.MissingParam"
//  UNAUTHORIZEDOPERATION_AUTHORIZE = "UnauthorizedOperation.Authorize"
//  UNAUTHORIZEDOPERATION_INACTIVE = "UnauthorizedOperation.Inactive"
//  UNSUPPORTEDOPERATION_BIERROR = "UnsupportedOperation.BIError"
func (c *Client) DeleteAuthApiKeyWithContext(ctx context.Context, request *DeleteAuthApiKeyRequest) (response *DeleteAuthApiKeyResponse, err error) {
    if request == nil {
        request = NewDeleteAuthApiKeyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bi", APIVersion, "DeleteAuthApiKey")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteAuthApiKey require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteAuthApiKeyResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteDatasourceRequest() (request *DeleteDatasourceRequest) {
    request = &DeleteDatasourceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bi", APIVersion, "DeleteDatasource")
    
    
    return
}

func NewDeleteDatasourceResponse() (response *DeleteDatasourceResponse) {
    response = &DeleteDatasourceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteDatasource
// 删除数据源
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNAL = "InternalError.Internal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_MISSINGPARAM = "MissingParameter.MissingParam"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_AUTHORIZE = "UnauthorizedOperation.Authorize"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_BIERROR = "UnsupportedOperation.BIError"
func (c *Client) DeleteDatasource(request *DeleteDatasourceRequest) (response *DeleteDatasourceResponse, err error) {
    return c.DeleteDatasourceWithContext(context.Background(), request)
}

// DeleteDatasource
// 删除数据源
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNAL = "InternalError.Internal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_MISSINGPARAM = "MissingParameter.MissingParam"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_AUTHORIZE = "UnauthorizedOperation.Authorize"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_BIERROR = "UnsupportedOperation.BIError"
func (c *Client) DeleteDatasourceWithContext(ctx context.Context, request *DeleteDatasourceRequest) (response *DeleteDatasourceResponse, err error) {
    if request == nil {
        request = NewDeleteDatasourceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bi", APIVersion, "DeleteDatasource")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteDatasource require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteDatasourceResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteProjectRequest() (request *DeleteProjectRequest) {
    request = &DeleteProjectRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bi", APIVersion, "DeleteProject")
    
    
    return
}

func NewDeleteProjectResponse() (response *DeleteProjectResponse) {
    response = &DeleteProjectResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteProject
// 删除项目
//
// 可能返回的错误码:
//  UNAUTHORIZEDOPERATION_AUTHORIZE = "UnauthorizedOperation.Authorize"
func (c *Client) DeleteProject(request *DeleteProjectRequest) (response *DeleteProjectResponse, err error) {
    return c.DeleteProjectWithContext(context.Background(), request)
}

// DeleteProject
// 删除项目
//
// 可能返回的错误码:
//  UNAUTHORIZEDOPERATION_AUTHORIZE = "UnauthorizedOperation.Authorize"
func (c *Client) DeleteProjectWithContext(ctx context.Context, request *DeleteProjectRequest) (response *DeleteProjectResponse, err error) {
    if request == nil {
        request = NewDeleteProjectRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bi", APIVersion, "DeleteProject")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteProject require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteProjectResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteUserGroupRequest() (request *DeleteUserGroupRequest) {
    request = &DeleteUserGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bi", APIVersion, "DeleteUserGroup")
    
    
    return
}

func NewDeleteUserGroupResponse() (response *DeleteUserGroupResponse) {
    response = &DeleteUserGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteUserGroup
// DeleteUserGroup
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNAL = "InternalError.Internal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_EMBED = "InvalidParameter.Embed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_MISSINGPARAM = "MissingParameter.MissingParam"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_AUTHORIZE = "UnauthorizedOperation.Authorize"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_BIERROR = "UnsupportedOperation.BIError"
func (c *Client) DeleteUserGroup(request *DeleteUserGroupRequest) (response *DeleteUserGroupResponse, err error) {
    return c.DeleteUserGroupWithContext(context.Background(), request)
}

// DeleteUserGroup
// DeleteUserGroup
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNAL = "InternalError.Internal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_EMBED = "InvalidParameter.Embed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_MISSINGPARAM = "MissingParameter.MissingParam"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_AUTHORIZE = "UnauthorizedOperation.Authorize"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_BIERROR = "UnsupportedOperation.BIError"
func (c *Client) DeleteUserGroupWithContext(ctx context.Context, request *DeleteUserGroupRequest) (response *DeleteUserGroupResponse, err error) {
    if request == nil {
        request = NewDeleteUserGroupRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bi", APIVersion, "DeleteUserGroup")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteUserGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteUserGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteUserGroupMemberRequest() (request *DeleteUserGroupMemberRequest) {
    request = &DeleteUserGroupMemberRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bi", APIVersion, "DeleteUserGroupMember")
    
    
    return
}

func NewDeleteUserGroupMemberResponse() (response *DeleteUserGroupMemberResponse) {
    response = &DeleteUserGroupMemberResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteUserGroupMember
// DeleteUserGroupMember
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNAL = "InternalError.Internal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_EMBED = "InvalidParameter.Embed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_MISSINGPARAM = "MissingParameter.MissingParam"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_AUTHORIZE = "UnauthorizedOperation.Authorize"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_BIERROR = "UnsupportedOperation.BIError"
func (c *Client) DeleteUserGroupMember(request *DeleteUserGroupMemberRequest) (response *DeleteUserGroupMemberResponse, err error) {
    return c.DeleteUserGroupMemberWithContext(context.Background(), request)
}

// DeleteUserGroupMember
// DeleteUserGroupMember
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNAL = "InternalError.Internal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_EMBED = "InvalidParameter.Embed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_MISSINGPARAM = "MissingParameter.MissingParam"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_AUTHORIZE = "UnauthorizedOperation.Authorize"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_BIERROR = "UnsupportedOperation.BIError"
func (c *Client) DeleteUserGroupMemberWithContext(ctx context.Context, request *DeleteUserGroupMemberRequest) (response *DeleteUserGroupMemberResponse, err error) {
    if request == nil {
        request = NewDeleteUserGroupMemberRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bi", APIVersion, "DeleteUserGroupMember")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteUserGroupMember require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteUserGroupMemberResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteUserRoleRequest() (request *DeleteUserRoleRequest) {
    request = &DeleteUserRoleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bi", APIVersion, "DeleteUserRole")
    
    
    return
}

func NewDeleteUserRoleResponse() (response *DeleteUserRoleResponse) {
    response = &DeleteUserRoleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteUserRole
// 删除用户角色，会删除用户
//
// 可能返回的错误码:
//  INTERNALERROR_INTERNAL = "InternalError.Internal"
//  UNAUTHORIZEDOPERATION_AUTHORIZE = "UnauthorizedOperation.Authorize"
//  UNAUTHORIZEDOPERATION_INACTIVE = "UnauthorizedOperation.Inactive"
//  UNAUTHORIZEDOPERATION_USERNOTEXIST = "UnauthorizedOperation.UserNotExist"
func (c *Client) DeleteUserRole(request *DeleteUserRoleRequest) (response *DeleteUserRoleResponse, err error) {
    return c.DeleteUserRoleWithContext(context.Background(), request)
}

// DeleteUserRole
// 删除用户角色，会删除用户
//
// 可能返回的错误码:
//  INTERNALERROR_INTERNAL = "InternalError.Internal"
//  UNAUTHORIZEDOPERATION_AUTHORIZE = "UnauthorizedOperation.Authorize"
//  UNAUTHORIZEDOPERATION_INACTIVE = "UnauthorizedOperation.Inactive"
//  UNAUTHORIZEDOPERATION_USERNOTEXIST = "UnauthorizedOperation.UserNotExist"
func (c *Client) DeleteUserRoleWithContext(ctx context.Context, request *DeleteUserRoleRequest) (response *DeleteUserRoleResponse, err error) {
    if request == nil {
        request = NewDeleteUserRoleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bi", APIVersion, "DeleteUserRole")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteUserRole require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteUserRoleResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteUserRoleProjectRequest() (request *DeleteUserRoleProjectRequest) {
    request = &DeleteUserRoleProjectRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bi", APIVersion, "DeleteUserRoleProject")
    
    
    return
}

func NewDeleteUserRoleProjectResponse() (response *DeleteUserRoleProjectResponse) {
    response = &DeleteUserRoleProjectResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteUserRoleProject
// 项目内-删除用户角色
//
// 可能返回的错误码:
//  MISSINGPARAMETER_MISSINGPARAM = "MissingParameter.MissingParam"
//  UNAUTHORIZEDOPERATION_AUTHORIZE = "UnauthorizedOperation.Authorize"
func (c *Client) DeleteUserRoleProject(request *DeleteUserRoleProjectRequest) (response *DeleteUserRoleProjectResponse, err error) {
    return c.DeleteUserRoleProjectWithContext(context.Background(), request)
}

// DeleteUserRoleProject
// 项目内-删除用户角色
//
// 可能返回的错误码:
//  MISSINGPARAMETER_MISSINGPARAM = "MissingParameter.MissingParam"
//  UNAUTHORIZEDOPERATION_AUTHORIZE = "UnauthorizedOperation.Authorize"
func (c *Client) DeleteUserRoleProjectWithContext(ctx context.Context, request *DeleteUserRoleProjectRequest) (response *DeleteUserRoleProjectResponse, err error) {
    if request == nil {
        request = NewDeleteUserRoleProjectRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bi", APIVersion, "DeleteUserRoleProject")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteUserRoleProject require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteUserRoleProjectResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAuthApiKeyInfoRequest() (request *DescribeAuthApiKeyInfoRequest) {
    request = &DescribeAuthApiKeyInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bi", APIVersion, "DescribeAuthApiKeyInfo")
    
    
    return
}

func NewDescribeAuthApiKeyInfoResponse() (response *DescribeAuthApiKeyInfoResponse) {
    response = &DescribeAuthApiKeyInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAuthApiKeyInfo
// ApiKey信息
//
// 可能返回的错误码:
//  INTERNALERROR_INTERNAL = "InternalError.Internal"
//  MISSINGPARAMETER_MISSINGPARAM = "MissingParameter.MissingParam"
//  UNAUTHORIZEDOPERATION_AUTHORIZE = "UnauthorizedOperation.Authorize"
//  UNAUTHORIZEDOPERATION_INACTIVE = "UnauthorizedOperation.Inactive"
//  UNSUPPORTEDOPERATION_BIERROR = "UnsupportedOperation.BIError"
func (c *Client) DescribeAuthApiKeyInfo(request *DescribeAuthApiKeyInfoRequest) (response *DescribeAuthApiKeyInfoResponse, err error) {
    return c.DescribeAuthApiKeyInfoWithContext(context.Background(), request)
}

// DescribeAuthApiKeyInfo
// ApiKey信息
//
// 可能返回的错误码:
//  INTERNALERROR_INTERNAL = "InternalError.Internal"
//  MISSINGPARAMETER_MISSINGPARAM = "MissingParameter.MissingParam"
//  UNAUTHORIZEDOPERATION_AUTHORIZE = "UnauthorizedOperation.Authorize"
//  UNAUTHORIZEDOPERATION_INACTIVE = "UnauthorizedOperation.Inactive"
//  UNSUPPORTEDOPERATION_BIERROR = "UnsupportedOperation.BIError"
func (c *Client) DescribeAuthApiKeyInfoWithContext(ctx context.Context, request *DescribeAuthApiKeyInfoRequest) (response *DescribeAuthApiKeyInfoResponse, err error) {
    if request == nil {
        request = NewDescribeAuthApiKeyInfoRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bi", APIVersion, "DescribeAuthApiKeyInfo")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAuthApiKeyInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAuthApiKeyInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAuthApiKeyListRequest() (request *DescribeAuthApiKeyListRequest) {
    request = &DescribeAuthApiKeyListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bi", APIVersion, "DescribeAuthApiKeyList")
    
    
    return
}

func NewDescribeAuthApiKeyListResponse() (response *DescribeAuthApiKeyListResponse) {
    response = &DescribeAuthApiKeyListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAuthApiKeyList
// ApiKey列表
//
// 可能返回的错误码:
//  MISSINGPARAMETER_MISSINGPARAM = "MissingParameter.MissingParam"
//  UNAUTHORIZEDOPERATION_AUTHORIZE = "UnauthorizedOperation.Authorize"
//  UNAUTHORIZEDOPERATION_INACTIVE = "UnauthorizedOperation.Inactive"
//  UNAUTHORIZEDOPERATION_USERNOTEXIST = "UnauthorizedOperation.UserNotExist"
func (c *Client) DescribeAuthApiKeyList(request *DescribeAuthApiKeyListRequest) (response *DescribeAuthApiKeyListResponse, err error) {
    return c.DescribeAuthApiKeyListWithContext(context.Background(), request)
}

// DescribeAuthApiKeyList
// ApiKey列表
//
// 可能返回的错误码:
//  MISSINGPARAMETER_MISSINGPARAM = "MissingParameter.MissingParam"
//  UNAUTHORIZEDOPERATION_AUTHORIZE = "UnauthorizedOperation.Authorize"
//  UNAUTHORIZEDOPERATION_INACTIVE = "UnauthorizedOperation.Inactive"
//  UNAUTHORIZEDOPERATION_USERNOTEXIST = "UnauthorizedOperation.UserNotExist"
func (c *Client) DescribeAuthApiKeyListWithContext(ctx context.Context, request *DescribeAuthApiKeyListRequest) (response *DescribeAuthApiKeyListResponse, err error) {
    if request == nil {
        request = NewDescribeAuthApiKeyListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bi", APIVersion, "DescribeAuthApiKeyList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAuthApiKeyList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAuthApiKeyListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDatasourceListRequest() (request *DescribeDatasourceListRequest) {
    request = &DescribeDatasourceListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bi", APIVersion, "DescribeDatasourceList")
    
    
    return
}

func NewDescribeDatasourceListResponse() (response *DescribeDatasourceListResponse) {
    response = &DescribeDatasourceListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDatasourceList
// 查询数据源列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNAL = "InternalError.Internal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_MISSINGPARAM = "MissingParameter.MissingParam"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_AUTHORIZE = "UnauthorizedOperation.Authorize"
//  UNAUTHORIZEDOPERATION_INACTIVE = "UnauthorizedOperation.Inactive"
//  UNAUTHORIZEDOPERATION_USERNOTEXIST = "UnauthorizedOperation.UserNotExist"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_BIERROR = "UnsupportedOperation.BIError"
func (c *Client) DescribeDatasourceList(request *DescribeDatasourceListRequest) (response *DescribeDatasourceListResponse, err error) {
    return c.DescribeDatasourceListWithContext(context.Background(), request)
}

// DescribeDatasourceList
// 查询数据源列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNAL = "InternalError.Internal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_MISSINGPARAM = "MissingParameter.MissingParam"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_AUTHORIZE = "UnauthorizedOperation.Authorize"
//  UNAUTHORIZEDOPERATION_INACTIVE = "UnauthorizedOperation.Inactive"
//  UNAUTHORIZEDOPERATION_USERNOTEXIST = "UnauthorizedOperation.UserNotExist"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_BIERROR = "UnsupportedOperation.BIError"
func (c *Client) DescribeDatasourceListWithContext(ctx context.Context, request *DescribeDatasourceListRequest) (response *DescribeDatasourceListResponse, err error) {
    if request == nil {
        request = NewDescribeDatasourceListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bi", APIVersion, "DescribeDatasourceList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDatasourceList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDatasourceListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePageWidgetListRequest() (request *DescribePageWidgetListRequest) {
    request = &DescribePageWidgetListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bi", APIVersion, "DescribePageWidgetList")
    
    
    return
}

func NewDescribePageWidgetListResponse() (response *DescribePageWidgetListResponse) {
    response = &DescribePageWidgetListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribePageWidgetList
// 查询页面组件信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNAL = "InternalError.Internal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_MISSINGPARAM = "MissingParameter.MissingParam"
//  UNSUPPORTEDOPERATION_BIERROR = "UnsupportedOperation.BIError"
func (c *Client) DescribePageWidgetList(request *DescribePageWidgetListRequest) (response *DescribePageWidgetListResponse, err error) {
    return c.DescribePageWidgetListWithContext(context.Background(), request)
}

// DescribePageWidgetList
// 查询页面组件信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNAL = "InternalError.Internal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_MISSINGPARAM = "MissingParameter.MissingParam"
//  UNSUPPORTEDOPERATION_BIERROR = "UnsupportedOperation.BIError"
func (c *Client) DescribePageWidgetListWithContext(ctx context.Context, request *DescribePageWidgetListRequest) (response *DescribePageWidgetListResponse, err error) {
    if request == nil {
        request = NewDescribePageWidgetListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bi", APIVersion, "DescribePageWidgetList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePageWidgetList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePageWidgetListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePermissionRanksInfoRequest() (request *DescribePermissionRanksInfoRequest) {
    request = &DescribePermissionRanksInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bi", APIVersion, "DescribePermissionRanksInfo")
    
    
    return
}

func NewDescribePermissionRanksInfoResponse() (response *DescribePermissionRanksInfoResponse) {
    response = &DescribePermissionRanksInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribePermissionRanksInfo
// 根据角色或标签查询行列权限配置
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNAL = "InternalError.Internal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_MISSINGPARAM = "MissingParameter.MissingParam"
//  UNSUPPORTEDOPERATION_BIERROR = "UnsupportedOperation.BIError"
func (c *Client) DescribePermissionRanksInfo(request *DescribePermissionRanksInfoRequest) (response *DescribePermissionRanksInfoResponse, err error) {
    return c.DescribePermissionRanksInfoWithContext(context.Background(), request)
}

// DescribePermissionRanksInfo
// 根据角色或标签查询行列权限配置
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNAL = "InternalError.Internal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_MISSINGPARAM = "MissingParameter.MissingParam"
//  UNSUPPORTEDOPERATION_BIERROR = "UnsupportedOperation.BIError"
func (c *Client) DescribePermissionRanksInfoWithContext(ctx context.Context, request *DescribePermissionRanksInfoRequest) (response *DescribePermissionRanksInfoResponse, err error) {
    if request == nil {
        request = NewDescribePermissionRanksInfoRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bi", APIVersion, "DescribePermissionRanksInfo")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePermissionRanksInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePermissionRanksInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePermissionRoleInfoRequest() (request *DescribePermissionRoleInfoRequest) {
    request = &DescribePermissionRoleInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bi", APIVersion, "DescribePermissionRoleInfo")
    
    
    return
}

func NewDescribePermissionRoleInfoResponse() (response *DescribePermissionRoleInfoResponse) {
    response = &DescribePermissionRoleInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribePermissionRoleInfo
// 行列权限项目内角色列表接口1
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNAL = "InternalError.Internal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_MISSINGPARAM = "MissingParameter.MissingParam"
//  UNSUPPORTEDOPERATION_BIERROR = "UnsupportedOperation.BIError"
func (c *Client) DescribePermissionRoleInfo(request *DescribePermissionRoleInfoRequest) (response *DescribePermissionRoleInfoResponse, err error) {
    return c.DescribePermissionRoleInfoWithContext(context.Background(), request)
}

// DescribePermissionRoleInfo
// 行列权限项目内角色列表接口1
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNAL = "InternalError.Internal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_MISSINGPARAM = "MissingParameter.MissingParam"
//  UNSUPPORTEDOPERATION_BIERROR = "UnsupportedOperation.BIError"
func (c *Client) DescribePermissionRoleInfoWithContext(ctx context.Context, request *DescribePermissionRoleInfoRequest) (response *DescribePermissionRoleInfoResponse, err error) {
    if request == nil {
        request = NewDescribePermissionRoleInfoRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bi", APIVersion, "DescribePermissionRoleInfo")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePermissionRoleInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePermissionRoleInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePermissionStatusInfoRequest() (request *DescribePermissionStatusInfoRequest) {
    request = &DescribePermissionStatusInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bi", APIVersion, "DescribePermissionStatusInfo")
    
    
    return
}

func NewDescribePermissionStatusInfoResponse() (response *DescribePermissionStatusInfoResponse) {
    response = &DescribePermissionStatusInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribePermissionStatusInfo
// 查询行列权限初始状态1
//
// 可能返回的错误码:
//  MISSINGPARAMETER_MISSINGPARAM = "MissingParameter.MissingParam"
//  UNSUPPORTEDOPERATION_BIERROR = "UnsupportedOperation.BIError"
func (c *Client) DescribePermissionStatusInfo(request *DescribePermissionStatusInfoRequest) (response *DescribePermissionStatusInfoResponse, err error) {
    return c.DescribePermissionStatusInfoWithContext(context.Background(), request)
}

// DescribePermissionStatusInfo
// 查询行列权限初始状态1
//
// 可能返回的错误码:
//  MISSINGPARAMETER_MISSINGPARAM = "MissingParameter.MissingParam"
//  UNSUPPORTEDOPERATION_BIERROR = "UnsupportedOperation.BIError"
func (c *Client) DescribePermissionStatusInfoWithContext(ctx context.Context, request *DescribePermissionStatusInfoRequest) (response *DescribePermissionStatusInfoResponse, err error) {
    if request == nil {
        request = NewDescribePermissionStatusInfoRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bi", APIVersion, "DescribePermissionStatusInfo")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePermissionStatusInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePermissionStatusInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeProjectInfoRequest() (request *DescribeProjectInfoRequest) {
    request = &DescribeProjectInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bi", APIVersion, "DescribeProjectInfo")
    
    
    return
}

func NewDescribeProjectInfoResponse() (response *DescribeProjectInfoResponse) {
    response = &DescribeProjectInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeProjectInfo
// 项目详情接口
//
// 可能返回的错误码:
//  MISSINGPARAMETER_MISSINGPARAM = "MissingParameter.MissingParam"
//  UNAUTHORIZEDOPERATION_AUTHORIZE = "UnauthorizedOperation.Authorize"
//  UNAUTHORIZEDOPERATION_USERNOTEXIST = "UnauthorizedOperation.UserNotExist"
//  UNSUPPORTEDOPERATION_BIERROR = "UnsupportedOperation.BIError"
func (c *Client) DescribeProjectInfo(request *DescribeProjectInfoRequest) (response *DescribeProjectInfoResponse, err error) {
    return c.DescribeProjectInfoWithContext(context.Background(), request)
}

// DescribeProjectInfo
// 项目详情接口
//
// 可能返回的错误码:
//  MISSINGPARAMETER_MISSINGPARAM = "MissingParameter.MissingParam"
//  UNAUTHORIZEDOPERATION_AUTHORIZE = "UnauthorizedOperation.Authorize"
//  UNAUTHORIZEDOPERATION_USERNOTEXIST = "UnauthorizedOperation.UserNotExist"
//  UNSUPPORTEDOPERATION_BIERROR = "UnsupportedOperation.BIError"
func (c *Client) DescribeProjectInfoWithContext(ctx context.Context, request *DescribeProjectInfoRequest) (response *DescribeProjectInfoResponse, err error) {
    if request == nil {
        request = NewDescribeProjectInfoRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bi", APIVersion, "DescribeProjectInfo")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeProjectInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeProjectInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeProjectListRequest() (request *DescribeProjectListRequest) {
    request = &DescribeProjectListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bi", APIVersion, "DescribeProjectList")
    
    
    return
}

func NewDescribeProjectListResponse() (response *DescribeProjectListResponse) {
    response = &DescribeProjectListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeProjectList
// 项目信息
//
// 可能返回的错误码:
//  INTERNALERROR_INTERNAL = "InternalError.Internal"
//  MISSINGPARAMETER_MISSINGPARAM = "MissingParameter.MissingParam"
//  UNAUTHORIZEDOPERATION_AUTHORIZE = "UnauthorizedOperation.Authorize"
//  UNAUTHORIZEDOPERATION_INACTIVE = "UnauthorizedOperation.Inactive"
//  UNAUTHORIZEDOPERATION_USERNOTEXIST = "UnauthorizedOperation.UserNotExist"
func (c *Client) DescribeProjectList(request *DescribeProjectListRequest) (response *DescribeProjectListResponse, err error) {
    return c.DescribeProjectListWithContext(context.Background(), request)
}

// DescribeProjectList
// 项目信息
//
// 可能返回的错误码:
//  INTERNALERROR_INTERNAL = "InternalError.Internal"
//  MISSINGPARAMETER_MISSINGPARAM = "MissingParameter.MissingParam"
//  UNAUTHORIZEDOPERATION_AUTHORIZE = "UnauthorizedOperation.Authorize"
//  UNAUTHORIZEDOPERATION_INACTIVE = "UnauthorizedOperation.Inactive"
//  UNAUTHORIZEDOPERATION_USERNOTEXIST = "UnauthorizedOperation.UserNotExist"
func (c *Client) DescribeProjectListWithContext(ctx context.Context, request *DescribeProjectListRequest) (response *DescribeProjectListResponse, err error) {
    if request == nil {
        request = NewDescribeProjectListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bi", APIVersion, "DescribeProjectList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeProjectList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeProjectListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeResourceUserGroupPageListRequest() (request *DescribeResourceUserGroupPageListRequest) {
    request = &DescribeResourceUserGroupPageListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bi", APIVersion, "DescribeResourceUserGroupPageList")
    
    
    return
}

func NewDescribeResourceUserGroupPageListResponse() (response *DescribeResourceUserGroupPageListResponse) {
    response = &DescribeResourceUserGroupPageListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeResourceUserGroupPageList
// 用户组资源权限查询接口
//
// 可能返回的错误码:
//  INTERNALERROR_INTERNAL = "InternalError.Internal"
//  MISSINGPARAMETER_MISSINGPARAM = "MissingParameter.MissingParam"
//  UNAUTHORIZEDOPERATION_AUTHORIZE = "UnauthorizedOperation.Authorize"
func (c *Client) DescribeResourceUserGroupPageList(request *DescribeResourceUserGroupPageListRequest) (response *DescribeResourceUserGroupPageListResponse, err error) {
    return c.DescribeResourceUserGroupPageListWithContext(context.Background(), request)
}

// DescribeResourceUserGroupPageList
// 用户组资源权限查询接口
//
// 可能返回的错误码:
//  INTERNALERROR_INTERNAL = "InternalError.Internal"
//  MISSINGPARAMETER_MISSINGPARAM = "MissingParameter.MissingParam"
//  UNAUTHORIZEDOPERATION_AUTHORIZE = "UnauthorizedOperation.Authorize"
func (c *Client) DescribeResourceUserGroupPageListWithContext(ctx context.Context, request *DescribeResourceUserGroupPageListRequest) (response *DescribeResourceUserGroupPageListResponse, err error) {
    if request == nil {
        request = NewDescribeResourceUserGroupPageListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bi", APIVersion, "DescribeResourceUserGroupPageList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeResourceUserGroupPageList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeResourceUserGroupPageListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSourceFieldListRequest() (request *DescribeSourceFieldListRequest) {
    request = &DescribeSourceFieldListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bi", APIVersion, "DescribeSourceFieldList")
    
    
    return
}

func NewDescribeSourceFieldListResponse() (response *DescribeSourceFieldListResponse) {
    response = &DescribeSourceFieldListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSourceFieldList
// 原始数据表字段接口信息
//
// 可能返回的错误码:
//  INTERNALERROR_INTERNAL = "InternalError.Internal"
//  MISSINGPARAMETER_MISSINGPARAM = "MissingParameter.MissingParam"
//  UNAUTHORIZEDOPERATION_AUTHORIZE = "UnauthorizedOperation.Authorize"
//  UNSUPPORTEDOPERATION_BIERROR = "UnsupportedOperation.BIError"
func (c *Client) DescribeSourceFieldList(request *DescribeSourceFieldListRequest) (response *DescribeSourceFieldListResponse, err error) {
    return c.DescribeSourceFieldListWithContext(context.Background(), request)
}

// DescribeSourceFieldList
// 原始数据表字段接口信息
//
// 可能返回的错误码:
//  INTERNALERROR_INTERNAL = "InternalError.Internal"
//  MISSINGPARAMETER_MISSINGPARAM = "MissingParameter.MissingParam"
//  UNAUTHORIZEDOPERATION_AUTHORIZE = "UnauthorizedOperation.Authorize"
//  UNSUPPORTEDOPERATION_BIERROR = "UnsupportedOperation.BIError"
func (c *Client) DescribeSourceFieldListWithContext(ctx context.Context, request *DescribeSourceFieldListRequest) (response *DescribeSourceFieldListResponse, err error) {
    if request == nil {
        request = NewDescribeSourceFieldListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bi", APIVersion, "DescribeSourceFieldList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSourceFieldList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSourceFieldListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUserGroupInfoRequest() (request *DescribeUserGroupInfoRequest) {
    request = &DescribeUserGroupInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bi", APIVersion, "DescribeUserGroupInfo")
    
    
    return
}

func NewDescribeUserGroupInfoResponse() (response *DescribeUserGroupInfoResponse) {
    response = &DescribeUserGroupInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeUserGroupInfo
// DescribeUserGroupInfo
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNAL = "InternalError.Internal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_EMBED = "InvalidParameter.Embed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_MISSINGPARAM = "MissingParameter.MissingParam"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_AUTHORIZE = "UnauthorizedOperation.Authorize"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_BIERROR = "UnsupportedOperation.BIError"
func (c *Client) DescribeUserGroupInfo(request *DescribeUserGroupInfoRequest) (response *DescribeUserGroupInfoResponse, err error) {
    return c.DescribeUserGroupInfoWithContext(context.Background(), request)
}

// DescribeUserGroupInfo
// DescribeUserGroupInfo
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNAL = "InternalError.Internal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_EMBED = "InvalidParameter.Embed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_MISSINGPARAM = "MissingParameter.MissingParam"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_AUTHORIZE = "UnauthorizedOperation.Authorize"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_BIERROR = "UnsupportedOperation.BIError"
func (c *Client) DescribeUserGroupInfoWithContext(ctx context.Context, request *DescribeUserGroupInfoRequest) (response *DescribeUserGroupInfoResponse, err error) {
    if request == nil {
        request = NewDescribeUserGroupInfoRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bi", APIVersion, "DescribeUserGroupInfo")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUserGroupInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeUserGroupInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUserGroupMemberListRequest() (request *DescribeUserGroupMemberListRequest) {
    request = &DescribeUserGroupMemberListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bi", APIVersion, "DescribeUserGroupMemberList")
    
    
    return
}

func NewDescribeUserGroupMemberListResponse() (response *DescribeUserGroupMemberListResponse) {
    response = &DescribeUserGroupMemberListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeUserGroupMemberList
// DescribeUserGroupMemberList
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNAL = "InternalError.Internal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_EMBED = "InvalidParameter.Embed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_MISSINGPARAM = "MissingParameter.MissingParam"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_AUTHORIZE = "UnauthorizedOperation.Authorize"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_BIERROR = "UnsupportedOperation.BIError"
func (c *Client) DescribeUserGroupMemberList(request *DescribeUserGroupMemberListRequest) (response *DescribeUserGroupMemberListResponse, err error) {
    return c.DescribeUserGroupMemberListWithContext(context.Background(), request)
}

// DescribeUserGroupMemberList
// DescribeUserGroupMemberList
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNAL = "InternalError.Internal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_EMBED = "InvalidParameter.Embed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_MISSINGPARAM = "MissingParameter.MissingParam"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_AUTHORIZE = "UnauthorizedOperation.Authorize"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_BIERROR = "UnsupportedOperation.BIError"
func (c *Client) DescribeUserGroupMemberListWithContext(ctx context.Context, request *DescribeUserGroupMemberListRequest) (response *DescribeUserGroupMemberListResponse, err error) {
    if request == nil {
        request = NewDescribeUserGroupMemberListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bi", APIVersion, "DescribeUserGroupMemberList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUserGroupMemberList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeUserGroupMemberListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUserGroupTreeListRequest() (request *DescribeUserGroupTreeListRequest) {
    request = &DescribeUserGroupTreeListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bi", APIVersion, "DescribeUserGroupTreeList")
    
    
    return
}

func NewDescribeUserGroupTreeListResponse() (response *DescribeUserGroupTreeListResponse) {
    response = &DescribeUserGroupTreeListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeUserGroupTreeList
// 用户组数查询接口
//
// 可能返回的错误码:
//  INTERNALERROR_INTERNAL = "InternalError.Internal"
//  MISSINGPARAMETER_MISSINGPARAM = "MissingParameter.MissingParam"
//  UNAUTHORIZEDOPERATION_AUTHORIZE = "UnauthorizedOperation.Authorize"
func (c *Client) DescribeUserGroupTreeList(request *DescribeUserGroupTreeListRequest) (response *DescribeUserGroupTreeListResponse, err error) {
    return c.DescribeUserGroupTreeListWithContext(context.Background(), request)
}

// DescribeUserGroupTreeList
// 用户组数查询接口
//
// 可能返回的错误码:
//  INTERNALERROR_INTERNAL = "InternalError.Internal"
//  MISSINGPARAMETER_MISSINGPARAM = "MissingParameter.MissingParam"
//  UNAUTHORIZEDOPERATION_AUTHORIZE = "UnauthorizedOperation.Authorize"
func (c *Client) DescribeUserGroupTreeListWithContext(ctx context.Context, request *DescribeUserGroupTreeListRequest) (response *DescribeUserGroupTreeListResponse, err error) {
    if request == nil {
        request = NewDescribeUserGroupTreeListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bi", APIVersion, "DescribeUserGroupTreeList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUserGroupTreeList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeUserGroupTreeListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUserProjectListRequest() (request *DescribeUserProjectListRequest) {
    request = &DescribeUserProjectListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bi", APIVersion, "DescribeUserProjectList")
    
    
    return
}

func NewDescribeUserProjectListResponse() (response *DescribeUserProjectListResponse) {
    response = &DescribeUserProjectListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeUserProjectList
// 项目内-用户接口
//
// 可能返回的错误码:
//  MISSINGPARAMETER_MISSINGPARAM = "MissingParameter.MissingParam"
//  UNSUPPORTEDOPERATION_BIERROR = "UnsupportedOperation.BIError"
func (c *Client) DescribeUserProjectList(request *DescribeUserProjectListRequest) (response *DescribeUserProjectListResponse, err error) {
    return c.DescribeUserProjectListWithContext(context.Background(), request)
}

// DescribeUserProjectList
// 项目内-用户接口
//
// 可能返回的错误码:
//  MISSINGPARAMETER_MISSINGPARAM = "MissingParameter.MissingParam"
//  UNSUPPORTEDOPERATION_BIERROR = "UnsupportedOperation.BIError"
func (c *Client) DescribeUserProjectListWithContext(ctx context.Context, request *DescribeUserProjectListRequest) (response *DescribeUserProjectListResponse, err error) {
    if request == nil {
        request = NewDescribeUserProjectListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bi", APIVersion, "DescribeUserProjectList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUserProjectList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeUserProjectListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUserRoleListRequest() (request *DescribeUserRoleListRequest) {
    request = &DescribeUserRoleListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bi", APIVersion, "DescribeUserRoleList")
    
    
    return
}

func NewDescribeUserRoleListResponse() (response *DescribeUserRoleListResponse) {
    response = &DescribeUserRoleListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeUserRoleList
// 用户角色列表
//
// 可能返回的错误码:
//  INTERNALERROR_INTERNAL = "InternalError.Internal"
//  MISSINGPARAMETER_MISSINGPARAM = "MissingParameter.MissingParam"
//  UNAUTHORIZEDOPERATION_AUTHORIZE = "UnauthorizedOperation.Authorize"
//  UNAUTHORIZEDOPERATION_INACTIVE = "UnauthorizedOperation.Inactive"
//  UNAUTHORIZEDOPERATION_USERNOTEXIST = "UnauthorizedOperation.UserNotExist"
//  UNSUPPORTEDOPERATION_BIERROR = "UnsupportedOperation.BIError"
func (c *Client) DescribeUserRoleList(request *DescribeUserRoleListRequest) (response *DescribeUserRoleListResponse, err error) {
    return c.DescribeUserRoleListWithContext(context.Background(), request)
}

// DescribeUserRoleList
// 用户角色列表
//
// 可能返回的错误码:
//  INTERNALERROR_INTERNAL = "InternalError.Internal"
//  MISSINGPARAMETER_MISSINGPARAM = "MissingParameter.MissingParam"
//  UNAUTHORIZEDOPERATION_AUTHORIZE = "UnauthorizedOperation.Authorize"
//  UNAUTHORIZEDOPERATION_INACTIVE = "UnauthorizedOperation.Inactive"
//  UNAUTHORIZEDOPERATION_USERNOTEXIST = "UnauthorizedOperation.UserNotExist"
//  UNSUPPORTEDOPERATION_BIERROR = "UnsupportedOperation.BIError"
func (c *Client) DescribeUserRoleListWithContext(ctx context.Context, request *DescribeUserRoleListRequest) (response *DescribeUserRoleListResponse, err error) {
    if request == nil {
        request = NewDescribeUserRoleListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bi", APIVersion, "DescribeUserRoleList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUserRoleList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeUserRoleListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUserRoleProjectListRequest() (request *DescribeUserRoleProjectListRequest) {
    request = &DescribeUserRoleProjectListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bi", APIVersion, "DescribeUserRoleProjectList")
    
    
    return
}

func NewDescribeUserRoleProjectListResponse() (response *DescribeUserRoleProjectListResponse) {
    response = &DescribeUserRoleProjectListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeUserRoleProjectList
// 项目内-用户角色列表
//
// 可能返回的错误码:
//  INTERNALERROR_INTERNAL = "InternalError.Internal"
//  MISSINGPARAMETER_MISSINGPARAM = "MissingParameter.MissingParam"
//  UNAUTHORIZEDOPERATION_AUTHORIZE = "UnauthorizedOperation.Authorize"
//  UNAUTHORIZEDOPERATION_INACTIVE = "UnauthorizedOperation.Inactive"
//  UNAUTHORIZEDOPERATION_USERNOTEXIST = "UnauthorizedOperation.UserNotExist"
//  UNSUPPORTEDOPERATION_BIERROR = "UnsupportedOperation.BIError"
func (c *Client) DescribeUserRoleProjectList(request *DescribeUserRoleProjectListRequest) (response *DescribeUserRoleProjectListResponse, err error) {
    return c.DescribeUserRoleProjectListWithContext(context.Background(), request)
}

// DescribeUserRoleProjectList
// 项目内-用户角色列表
//
// 可能返回的错误码:
//  INTERNALERROR_INTERNAL = "InternalError.Internal"
//  MISSINGPARAMETER_MISSINGPARAM = "MissingParameter.MissingParam"
//  UNAUTHORIZEDOPERATION_AUTHORIZE = "UnauthorizedOperation.Authorize"
//  UNAUTHORIZEDOPERATION_INACTIVE = "UnauthorizedOperation.Inactive"
//  UNAUTHORIZEDOPERATION_USERNOTEXIST = "UnauthorizedOperation.UserNotExist"
//  UNSUPPORTEDOPERATION_BIERROR = "UnsupportedOperation.BIError"
func (c *Client) DescribeUserRoleProjectListWithContext(ctx context.Context, request *DescribeUserRoleProjectListRequest) (response *DescribeUserRoleProjectListResponse, err error) {
    if request == nil {
        request = NewDescribeUserRoleProjectListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bi", APIVersion, "DescribeUserRoleProjectList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUserRoleProjectList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeUserRoleProjectListResponse()
    err = c.Send(request, response)
    return
}

func NewExportScreenPageRequest() (request *ExportScreenPageRequest) {
    request = &ExportScreenPageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bi", APIVersion, "ExportScreenPage")
    
    
    return
}

func NewExportScreenPageResponse() (response *ExportScreenPageResponse) {
    response = &ExportScreenPageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ExportScreenPage
// 页面截图导出
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNAL = "InternalError.Internal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_MISSINGPARAM = "MissingParameter.MissingParam"
//  UNSUPPORTEDOPERATION_BIERROR = "UnsupportedOperation.BIError"
func (c *Client) ExportScreenPage(request *ExportScreenPageRequest) (response *ExportScreenPageResponse, err error) {
    return c.ExportScreenPageWithContext(context.Background(), request)
}

// ExportScreenPage
// 页面截图导出
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNAL = "InternalError.Internal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_MISSINGPARAM = "MissingParameter.MissingParam"
//  UNSUPPORTEDOPERATION_BIERROR = "UnsupportedOperation.BIError"
func (c *Client) ExportScreenPageWithContext(ctx context.Context, request *ExportScreenPageRequest) (response *ExportScreenPageResponse, err error) {
    if request == nil {
        request = NewExportScreenPageRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bi", APIVersion, "ExportScreenPage")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ExportScreenPage require credential")
    }

    request.SetContext(ctx)
    
    response = NewExportScreenPageResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAuthApiKeyRequest() (request *ModifyAuthApiKeyRequest) {
    request = &ModifyAuthApiKeyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bi", APIVersion, "ModifyAuthApiKey")
    
    
    return
}

func NewModifyAuthApiKeyResponse() (response *ModifyAuthApiKeyResponse) {
    response = &ModifyAuthApiKeyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyAuthApiKey
// 更新ApiKey
//
// 可能返回的错误码:
//  INTERNALERROR_INTERNAL = "InternalError.Internal"
//  MISSINGPARAMETER_MISSINGPARAM = "MissingParameter.MissingParam"
//  UNAUTHORIZEDOPERATION_AUTHORIZE = "UnauthorizedOperation.Authorize"
//  UNAUTHORIZEDOPERATION_INACTIVE = "UnauthorizedOperation.Inactive"
//  UNSUPPORTEDOPERATION_BIERROR = "UnsupportedOperation.BIError"
func (c *Client) ModifyAuthApiKey(request *ModifyAuthApiKeyRequest) (response *ModifyAuthApiKeyResponse, err error) {
    return c.ModifyAuthApiKeyWithContext(context.Background(), request)
}

// ModifyAuthApiKey
// 更新ApiKey
//
// 可能返回的错误码:
//  INTERNALERROR_INTERNAL = "InternalError.Internal"
//  MISSINGPARAMETER_MISSINGPARAM = "MissingParameter.MissingParam"
//  UNAUTHORIZEDOPERATION_AUTHORIZE = "UnauthorizedOperation.Authorize"
//  UNAUTHORIZEDOPERATION_INACTIVE = "UnauthorizedOperation.Inactive"
//  UNSUPPORTEDOPERATION_BIERROR = "UnsupportedOperation.BIError"
func (c *Client) ModifyAuthApiKeyWithContext(ctx context.Context, request *ModifyAuthApiKeyRequest) (response *ModifyAuthApiKeyResponse, err error) {
    if request == nil {
        request = NewModifyAuthApiKeyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bi", APIVersion, "ModifyAuthApiKey")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAuthApiKey require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAuthApiKeyResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDatasourceRequest() (request *ModifyDatasourceRequest) {
    request = &ModifyDatasourceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bi", APIVersion, "ModifyDatasource")
    
    
    return
}

func NewModifyDatasourceResponse() (response *ModifyDatasourceResponse) {
    response = &ModifyDatasourceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyDatasource
// 更新数据源
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNAL = "InternalError.Internal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_MISSINGPARAM = "MissingParameter.MissingParam"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_AUTHORIZE = "UnauthorizedOperation.Authorize"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_BIERROR = "UnsupportedOperation.BIError"
func (c *Client) ModifyDatasource(request *ModifyDatasourceRequest) (response *ModifyDatasourceResponse, err error) {
    return c.ModifyDatasourceWithContext(context.Background(), request)
}

// ModifyDatasource
// 更新数据源
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNAL = "InternalError.Internal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_MISSINGPARAM = "MissingParameter.MissingParam"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_AUTHORIZE = "UnauthorizedOperation.Authorize"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_BIERROR = "UnsupportedOperation.BIError"
func (c *Client) ModifyDatasourceWithContext(ctx context.Context, request *ModifyDatasourceRequest) (response *ModifyDatasourceResponse, err error) {
    if request == nil {
        request = NewModifyDatasourceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bi", APIVersion, "ModifyDatasource")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDatasource require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDatasourceResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDatasourceCloudRequest() (request *ModifyDatasourceCloudRequest) {
    request = &ModifyDatasourceCloudRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bi", APIVersion, "ModifyDatasourceCloud")
    
    
    return
}

func NewModifyDatasourceCloudResponse() (response *ModifyDatasourceCloudResponse) {
    response = &ModifyDatasourceCloudResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyDatasourceCloud
// 更新云数据库
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNAL = "InternalError.Internal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_MISSINGPARAM = "MissingParameter.MissingParam"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_AUTHORIZE = "UnauthorizedOperation.Authorize"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_BIERROR = "UnsupportedOperation.BIError"
func (c *Client) ModifyDatasourceCloud(request *ModifyDatasourceCloudRequest) (response *ModifyDatasourceCloudResponse, err error) {
    return c.ModifyDatasourceCloudWithContext(context.Background(), request)
}

// ModifyDatasourceCloud
// 更新云数据库
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNAL = "InternalError.Internal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_MISSINGPARAM = "MissingParameter.MissingParam"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_AUTHORIZE = "UnauthorizedOperation.Authorize"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_BIERROR = "UnsupportedOperation.BIError"
func (c *Client) ModifyDatasourceCloudWithContext(ctx context.Context, request *ModifyDatasourceCloudRequest) (response *ModifyDatasourceCloudResponse, err error) {
    if request == nil {
        request = NewModifyDatasourceCloudRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bi", APIVersion, "ModifyDatasourceCloud")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDatasourceCloud require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDatasourceCloudResponse()
    err = c.Send(request, response)
    return
}

func NewModifyProjectRequest() (request *ModifyProjectRequest) {
    request = &ModifyProjectRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bi", APIVersion, "ModifyProject")
    
    
    return
}

func NewModifyProjectResponse() (response *ModifyProjectResponse) {
    response = &ModifyProjectResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyProject
// 修改项目信息
//
// 可能返回的错误码:
//  MISSINGPARAMETER_MISSINGPARAM = "MissingParameter.MissingParam"
//  UNAUTHORIZEDOPERATION_AUTHORIZE = "UnauthorizedOperation.Authorize"
func (c *Client) ModifyProject(request *ModifyProjectRequest) (response *ModifyProjectResponse, err error) {
    return c.ModifyProjectWithContext(context.Background(), request)
}

// ModifyProject
// 修改项目信息
//
// 可能返回的错误码:
//  MISSINGPARAMETER_MISSINGPARAM = "MissingParameter.MissingParam"
//  UNAUTHORIZEDOPERATION_AUTHORIZE = "UnauthorizedOperation.Authorize"
func (c *Client) ModifyProjectWithContext(ctx context.Context, request *ModifyProjectRequest) (response *ModifyProjectResponse, err error) {
    if request == nil {
        request = NewModifyProjectRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bi", APIVersion, "ModifyProject")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyProject require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyProjectResponse()
    err = c.Send(request, response)
    return
}

func NewModifyResourceUserGroupRequest() (request *ModifyResourceUserGroupRequest) {
    request = &ModifyResourceUserGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bi", APIVersion, "ModifyResourceUserGroup")
    
    
    return
}

func NewModifyResourceUserGroupResponse() (response *ModifyResourceUserGroupResponse) {
    response = &ModifyResourceUserGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyResourceUserGroup
// 更新用户组权限
//
// 可能返回的错误码:
//  INTERNALERROR_INTERNAL = "InternalError.Internal"
//  MISSINGPARAMETER_MISSINGPARAM = "MissingParameter.MissingParam"
//  UNAUTHORIZEDOPERATION_AUTHORIZE = "UnauthorizedOperation.Authorize"
func (c *Client) ModifyResourceUserGroup(request *ModifyResourceUserGroupRequest) (response *ModifyResourceUserGroupResponse, err error) {
    return c.ModifyResourceUserGroupWithContext(context.Background(), request)
}

// ModifyResourceUserGroup
// 更新用户组权限
//
// 可能返回的错误码:
//  INTERNALERROR_INTERNAL = "InternalError.Internal"
//  MISSINGPARAMETER_MISSINGPARAM = "MissingParameter.MissingParam"
//  UNAUTHORIZEDOPERATION_AUTHORIZE = "UnauthorizedOperation.Authorize"
func (c *Client) ModifyResourceUserGroupWithContext(ctx context.Context, request *ModifyResourceUserGroupRequest) (response *ModifyResourceUserGroupResponse, err error) {
    if request == nil {
        request = NewModifyResourceUserGroupRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bi", APIVersion, "ModifyResourceUserGroup")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyResourceUserGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyResourceUserGroupResponse()
    err = c.Send(request, response)
    return
}

func NewModifyResourceUserGroupResourceRequest() (request *ModifyResourceUserGroupResourceRequest) {
    request = &ModifyResourceUserGroupResourceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bi", APIVersion, "ModifyResourceUserGroupResource")
    
    
    return
}

func NewModifyResourceUserGroupResourceResponse() (response *ModifyResourceUserGroupResourceResponse) {
    response = &ModifyResourceUserGroupResourceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyResourceUserGroupResource
// 按资源 - 更新用户组权限
//
// 可能返回的错误码:
//  INTERNALERROR_INTERNAL = "InternalError.Internal"
//  MISSINGPARAMETER_MISSINGPARAM = "MissingParameter.MissingParam"
//  UNAUTHORIZEDOPERATION_AUTHORIZE = "UnauthorizedOperation.Authorize"
func (c *Client) ModifyResourceUserGroupResource(request *ModifyResourceUserGroupResourceRequest) (response *ModifyResourceUserGroupResourceResponse, err error) {
    return c.ModifyResourceUserGroupResourceWithContext(context.Background(), request)
}

// ModifyResourceUserGroupResource
// 按资源 - 更新用户组权限
//
// 可能返回的错误码:
//  INTERNALERROR_INTERNAL = "InternalError.Internal"
//  MISSINGPARAMETER_MISSINGPARAM = "MissingParameter.MissingParam"
//  UNAUTHORIZEDOPERATION_AUTHORIZE = "UnauthorizedOperation.Authorize"
func (c *Client) ModifyResourceUserGroupResourceWithContext(ctx context.Context, request *ModifyResourceUserGroupResourceRequest) (response *ModifyResourceUserGroupResourceResponse, err error) {
    if request == nil {
        request = NewModifyResourceUserGroupResourceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bi", APIVersion, "ModifyResourceUserGroupResource")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyResourceUserGroupResource require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyResourceUserGroupResourceResponse()
    err = c.Send(request, response)
    return
}

func NewModifyUserDetailInfoRequest() (request *ModifyUserDetailInfoRequest) {
    request = &ModifyUserDetailInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bi", APIVersion, "ModifyUserDetailInfo")
    
    
    return
}

func NewModifyUserDetailInfoResponse() (response *ModifyUserDetailInfoResponse) {
    response = &ModifyUserDetailInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyUserDetailInfo
// 修改用户角色信息
//
// 可能返回的错误码:
//  INTERNALERROR_INTERNAL = "InternalError.Internal"
//  MISSINGPARAMETER_MISSINGPARAM = "MissingParameter.MissingParam"
//  UNAUTHORIZEDOPERATION_AUTHORIZE = "UnauthorizedOperation.Authorize"
func (c *Client) ModifyUserDetailInfo(request *ModifyUserDetailInfoRequest) (response *ModifyUserDetailInfoResponse, err error) {
    return c.ModifyUserDetailInfoWithContext(context.Background(), request)
}

// ModifyUserDetailInfo
// 修改用户角色信息
//
// 可能返回的错误码:
//  INTERNALERROR_INTERNAL = "InternalError.Internal"
//  MISSINGPARAMETER_MISSINGPARAM = "MissingParameter.MissingParam"
//  UNAUTHORIZEDOPERATION_AUTHORIZE = "UnauthorizedOperation.Authorize"
func (c *Client) ModifyUserDetailInfoWithContext(ctx context.Context, request *ModifyUserDetailInfoRequest) (response *ModifyUserDetailInfoResponse, err error) {
    if request == nil {
        request = NewModifyUserDetailInfoRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bi", APIVersion, "ModifyUserDetailInfo")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyUserDetailInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyUserDetailInfoResponse()
    err = c.Send(request, response)
    return
}

func NewModifyUserGroupRequest() (request *ModifyUserGroupRequest) {
    request = &ModifyUserGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bi", APIVersion, "ModifyUserGroup")
    
    
    return
}

func NewModifyUserGroupResponse() (response *ModifyUserGroupResponse) {
    response = &ModifyUserGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyUserGroup
// ModifyUserGroup
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNAL = "InternalError.Internal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_EMBED = "InvalidParameter.Embed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_MISSINGPARAM = "MissingParameter.MissingParam"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_AUTHORIZE = "UnauthorizedOperation.Authorize"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_BIERROR = "UnsupportedOperation.BIError"
func (c *Client) ModifyUserGroup(request *ModifyUserGroupRequest) (response *ModifyUserGroupResponse, err error) {
    return c.ModifyUserGroupWithContext(context.Background(), request)
}

// ModifyUserGroup
// ModifyUserGroup
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNAL = "InternalError.Internal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_EMBED = "InvalidParameter.Embed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_MISSINGPARAM = "MissingParameter.MissingParam"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_AUTHORIZE = "UnauthorizedOperation.Authorize"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_BIERROR = "UnsupportedOperation.BIError"
func (c *Client) ModifyUserGroupWithContext(ctx context.Context, request *ModifyUserGroupRequest) (response *ModifyUserGroupResponse, err error) {
    if request == nil {
        request = NewModifyUserGroupRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bi", APIVersion, "ModifyUserGroup")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyUserGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyUserGroupResponse()
    err = c.Send(request, response)
    return
}

func NewModifyUserRoleRequest() (request *ModifyUserRoleRequest) {
    request = &ModifyUserRoleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bi", APIVersion, "ModifyUserRole")
    
    
    return
}

func NewModifyUserRoleResponse() (response *ModifyUserRoleResponse) {
    response = &ModifyUserRoleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyUserRole
// 修改用户角色信息
//
// 可能返回的错误码:
//  INTERNALERROR_INTERNAL = "InternalError.Internal"
//  MISSINGPARAMETER_MISSINGPARAM = "MissingParameter.MissingParam"
//  UNAUTHORIZEDOPERATION_AUTHORIZE = "UnauthorizedOperation.Authorize"
func (c *Client) ModifyUserRole(request *ModifyUserRoleRequest) (response *ModifyUserRoleResponse, err error) {
    return c.ModifyUserRoleWithContext(context.Background(), request)
}

// ModifyUserRole
// 修改用户角色信息
//
// 可能返回的错误码:
//  INTERNALERROR_INTERNAL = "InternalError.Internal"
//  MISSINGPARAMETER_MISSINGPARAM = "MissingParameter.MissingParam"
//  UNAUTHORIZEDOPERATION_AUTHORIZE = "UnauthorizedOperation.Authorize"
func (c *Client) ModifyUserRoleWithContext(ctx context.Context, request *ModifyUserRoleRequest) (response *ModifyUserRoleResponse, err error) {
    if request == nil {
        request = NewModifyUserRoleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bi", APIVersion, "ModifyUserRole")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyUserRole require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyUserRoleResponse()
    err = c.Send(request, response)
    return
}

func NewModifyUserRoleProjectRequest() (request *ModifyUserRoleProjectRequest) {
    request = &ModifyUserRoleProjectRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bi", APIVersion, "ModifyUserRoleProject")
    
    
    return
}

func NewModifyUserRoleProjectResponse() (response *ModifyUserRoleProjectResponse) {
    response = &ModifyUserRoleProjectResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyUserRoleProject
// 项目-修改用户角色信息
//
// 可能返回的错误码:
//  INTERNALERROR_INTERNAL = "InternalError.Internal"
//  MISSINGPARAMETER_MISSINGPARAM = "MissingParameter.MissingParam"
//  UNAUTHORIZEDOPERATION_AUTHORIZE = "UnauthorizedOperation.Authorize"
func (c *Client) ModifyUserRoleProject(request *ModifyUserRoleProjectRequest) (response *ModifyUserRoleProjectResponse, err error) {
    return c.ModifyUserRoleProjectWithContext(context.Background(), request)
}

// ModifyUserRoleProject
// 项目-修改用户角色信息
//
// 可能返回的错误码:
//  INTERNALERROR_INTERNAL = "InternalError.Internal"
//  MISSINGPARAMETER_MISSINGPARAM = "MissingParameter.MissingParam"
//  UNAUTHORIZEDOPERATION_AUTHORIZE = "UnauthorizedOperation.Authorize"
func (c *Client) ModifyUserRoleProjectWithContext(ctx context.Context, request *ModifyUserRoleProjectRequest) (response *ModifyUserRoleProjectResponse, err error) {
    if request == nil {
        request = NewModifyUserRoleProjectRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bi", APIVersion, "ModifyUserRoleProject")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyUserRoleProject require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyUserRoleProjectResponse()
    err = c.Send(request, response)
    return
}

func NewQueryUserGroupMemberRequest() (request *QueryUserGroupMemberRequest) {
    request = &QueryUserGroupMemberRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bi", APIVersion, "QueryUserGroupMember")
    
    
    return
}

func NewQueryUserGroupMemberResponse() (response *QueryUserGroupMemberResponse) {
    response = &QueryUserGroupMemberResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// QueryUserGroupMember
// QueryUserGroupMember
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNAL = "InternalError.Internal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_EMBED = "InvalidParameter.Embed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_MISSINGPARAM = "MissingParameter.MissingParam"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_AUTHORIZE = "UnauthorizedOperation.Authorize"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_BIERROR = "UnsupportedOperation.BIError"
func (c *Client) QueryUserGroupMember(request *QueryUserGroupMemberRequest) (response *QueryUserGroupMemberResponse, err error) {
    return c.QueryUserGroupMemberWithContext(context.Background(), request)
}

// QueryUserGroupMember
// QueryUserGroupMember
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNAL = "InternalError.Internal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_EMBED = "InvalidParameter.Embed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_MISSINGPARAM = "MissingParameter.MissingParam"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_AUTHORIZE = "UnauthorizedOperation.Authorize"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_BIERROR = "UnsupportedOperation.BIError"
func (c *Client) QueryUserGroupMemberWithContext(ctx context.Context, request *QueryUserGroupMemberRequest) (response *QueryUserGroupMemberResponse, err error) {
    if request == nil {
        request = NewQueryUserGroupMemberRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bi", APIVersion, "QueryUserGroupMember")
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryUserGroupMember require credential")
    }

    request.SetContext(ctx)
    
    response = NewQueryUserGroupMemberResponse()
    err = c.Send(request, response)
    return
}
