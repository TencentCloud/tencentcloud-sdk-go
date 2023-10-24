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
    
    if c.GetCredential() == nil {
        return nil, errors.New("ApplyEmbedInterval require credential")
    }

    request.SetContext(ctx)
    
    response = NewApplyEmbedIntervalResponse()
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateEmbedToken require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateEmbedTokenResponse()
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateProject require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateProjectResponse()
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateUserRoleProject require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateUserRoleProjectResponse()
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteProject require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteProjectResponse()
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteUserRoleProject require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteUserRoleProjectResponse()
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDatasourceList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDatasourceListResponse()
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeProjectList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeProjectListResponse()
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUserRoleProjectList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeUserRoleProjectListResponse()
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyProject require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyProjectResponse()
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyUserRoleProject require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyUserRoleProjectResponse()
    err = c.Send(request, response)
    return
}
