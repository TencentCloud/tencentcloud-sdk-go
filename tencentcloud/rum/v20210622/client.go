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
    c.InitBaseRequest(&request.BaseRequest, "rum", APIVersion, "CreateReleaseFile")
    
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
    c.InitBaseRequest(&request.BaseRequest, "rum", APIVersion, "CreateStarProject")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateStarProject require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateStarProjectResponse()
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
    c.InitBaseRequest(&request.BaseRequest, "rum", APIVersion, "CreateWhitelist")
    
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
    c.InitBaseRequest(&request.BaseRequest, "rum", APIVersion, "DeleteInstance")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteInstanceResponse()
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
    c.InitBaseRequest(&request.BaseRequest, "rum", APIVersion, "DeleteProject")
    
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
    c.InitBaseRequest(&request.BaseRequest, "rum", APIVersion, "DeleteReleaseFile")
    
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
    c.InitBaseRequest(&request.BaseRequest, "rum", APIVersion, "DeleteStarProject")
    
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
    c.InitBaseRequest(&request.BaseRequest, "rum", APIVersion, "DeleteWhitelist")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteWhitelist require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteWhitelistResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAppDimensionMetricsRequest() (request *DescribeAppDimensionMetricsRequest) {
    request = &DescribeAppDimensionMetricsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "DescribeAppDimensionMetrics")
    
    
    return
}

func NewDescribeAppDimensionMetricsResponse() (response *DescribeAppDimensionMetricsResponse) {
    response = &DescribeAppDimensionMetricsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAppDimensionMetrics
// 用于查询 app 监控多维分析数据
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAppDimensionMetrics(request *DescribeAppDimensionMetricsRequest) (response *DescribeAppDimensionMetricsResponse, err error) {
    return c.DescribeAppDimensionMetricsWithContext(context.Background(), request)
}

// DescribeAppDimensionMetrics
// 用于查询 app 监控多维分析数据
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAppDimensionMetricsWithContext(ctx context.Context, request *DescribeAppDimensionMetricsRequest) (response *DescribeAppDimensionMetricsResponse, err error) {
    if request == nil {
        request = NewDescribeAppDimensionMetricsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "rum", APIVersion, "DescribeAppDimensionMetrics")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAppDimensionMetrics require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAppDimensionMetricsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAppMetricsDataRequest() (request *DescribeAppMetricsDataRequest) {
    request = &DescribeAppMetricsDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "DescribeAppMetricsData")
    
    
    return
}

func NewDescribeAppMetricsDataResponse() (response *DescribeAppMetricsDataResponse) {
    response = &DescribeAppMetricsDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAppMetricsData
// 获取 app 监控指标数据
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAppMetricsData(request *DescribeAppMetricsDataRequest) (response *DescribeAppMetricsDataResponse, err error) {
    return c.DescribeAppMetricsDataWithContext(context.Background(), request)
}

// DescribeAppMetricsData
// 获取 app 监控指标数据
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAppMetricsDataWithContext(ctx context.Context, request *DescribeAppMetricsDataRequest) (response *DescribeAppMetricsDataResponse, err error) {
    if request == nil {
        request = NewDescribeAppMetricsDataRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "rum", APIVersion, "DescribeAppMetricsData")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAppMetricsData require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAppMetricsDataResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAppSingleCaseDetailListRequest() (request *DescribeAppSingleCaseDetailListRequest) {
    request = &DescribeAppSingleCaseDetailListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "DescribeAppSingleCaseDetailList")
    
    
    return
}

func NewDescribeAppSingleCaseDetailListResponse() (response *DescribeAppSingleCaseDetailListResponse) {
    response = &DescribeAppSingleCaseDetailListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAppSingleCaseDetailList
// 查询 app 监控个例样本详情列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeAppSingleCaseDetailList(request *DescribeAppSingleCaseDetailListRequest) (response *DescribeAppSingleCaseDetailListResponse, err error) {
    return c.DescribeAppSingleCaseDetailListWithContext(context.Background(), request)
}

// DescribeAppSingleCaseDetailList
// 查询 app 监控个例样本详情列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeAppSingleCaseDetailListWithContext(ctx context.Context, request *DescribeAppSingleCaseDetailListRequest) (response *DescribeAppSingleCaseDetailListResponse, err error) {
    if request == nil {
        request = NewDescribeAppSingleCaseDetailListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "rum", APIVersion, "DescribeAppSingleCaseDetailList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAppSingleCaseDetailList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAppSingleCaseDetailListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAppSingleCaseListRequest() (request *DescribeAppSingleCaseListRequest) {
    request = &DescribeAppSingleCaseListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "DescribeAppSingleCaseList")
    
    
    return
}

func NewDescribeAppSingleCaseListResponse() (response *DescribeAppSingleCaseListResponse) {
    response = &DescribeAppSingleCaseListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAppSingleCaseList
// 查询 app 监控个例聚合列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAppSingleCaseList(request *DescribeAppSingleCaseListRequest) (response *DescribeAppSingleCaseListResponse, err error) {
    return c.DescribeAppSingleCaseListWithContext(context.Background(), request)
}

// DescribeAppSingleCaseList
// 查询 app 监控个例聚合列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAppSingleCaseListWithContext(ctx context.Context, request *DescribeAppSingleCaseListRequest) (response *DescribeAppSingleCaseListResponse, err error) {
    if request == nil {
        request = NewDescribeAppSingleCaseListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "rum", APIVersion, "DescribeAppSingleCaseList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAppSingleCaseList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAppSingleCaseListResponse()
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
    c.InitBaseRequest(&request.BaseRequest, "rum", APIVersion, "DescribeData")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeData require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDataResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDataBridgeUrlV2Request() (request *DescribeDataBridgeUrlV2Request) {
    request = &DescribeDataBridgeUrlV2Request{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "DescribeDataBridgeUrlV2")
    
    
    return
}

func NewDescribeDataBridgeUrlV2Response() (response *DescribeDataBridgeUrlV2Response) {
    response = &DescribeDataBridgeUrlV2Response{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDataBridgeUrlV2
// 获取DescribeDataBridgeUrl信息
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
func (c *Client) DescribeDataBridgeUrlV2(request *DescribeDataBridgeUrlV2Request) (response *DescribeDataBridgeUrlV2Response, err error) {
    return c.DescribeDataBridgeUrlV2WithContext(context.Background(), request)
}

// DescribeDataBridgeUrlV2
// 获取DescribeDataBridgeUrl信息
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
func (c *Client) DescribeDataBridgeUrlV2WithContext(ctx context.Context, request *DescribeDataBridgeUrlV2Request) (response *DescribeDataBridgeUrlV2Response, err error) {
    if request == nil {
        request = NewDescribeDataBridgeUrlV2Request()
    }
    c.InitBaseRequest(&request.BaseRequest, "rum", APIVersion, "DescribeDataBridgeUrlV2")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDataBridgeUrlV2 require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDataBridgeUrlV2Response()
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
    c.InitBaseRequest(&request.BaseRequest, "rum", APIVersion, "DescribeDataCustomUrl")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDataCustomUrl require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDataCustomUrlResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDataCustomUrlV2Request() (request *DescribeDataCustomUrlV2Request) {
    request = &DescribeDataCustomUrlV2Request{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "DescribeDataCustomUrlV2")
    
    
    return
}

func NewDescribeDataCustomUrlV2Response() (response *DescribeDataCustomUrlV2Response) {
    response = &DescribeDataCustomUrlV2Response{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDataCustomUrlV2
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
func (c *Client) DescribeDataCustomUrlV2(request *DescribeDataCustomUrlV2Request) (response *DescribeDataCustomUrlV2Response, err error) {
    return c.DescribeDataCustomUrlV2WithContext(context.Background(), request)
}

// DescribeDataCustomUrlV2
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
func (c *Client) DescribeDataCustomUrlV2WithContext(ctx context.Context, request *DescribeDataCustomUrlV2Request) (response *DescribeDataCustomUrlV2Response, err error) {
    if request == nil {
        request = NewDescribeDataCustomUrlV2Request()
    }
    c.InitBaseRequest(&request.BaseRequest, "rum", APIVersion, "DescribeDataCustomUrlV2")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDataCustomUrlV2 require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDataCustomUrlV2Response()
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
    c.InitBaseRequest(&request.BaseRequest, "rum", APIVersion, "DescribeDataEventUrl")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDataEventUrl require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDataEventUrlResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDataEventUrlV2Request() (request *DescribeDataEventUrlV2Request) {
    request = &DescribeDataEventUrlV2Request{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "DescribeDataEventUrlV2")
    
    
    return
}

func NewDescribeDataEventUrlV2Response() (response *DescribeDataEventUrlV2Response) {
    response = &DescribeDataEventUrlV2Response{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDataEventUrlV2
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
func (c *Client) DescribeDataEventUrlV2(request *DescribeDataEventUrlV2Request) (response *DescribeDataEventUrlV2Response, err error) {
    return c.DescribeDataEventUrlV2WithContext(context.Background(), request)
}

// DescribeDataEventUrlV2
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
func (c *Client) DescribeDataEventUrlV2WithContext(ctx context.Context, request *DescribeDataEventUrlV2Request) (response *DescribeDataEventUrlV2Response, err error) {
    if request == nil {
        request = NewDescribeDataEventUrlV2Request()
    }
    c.InitBaseRequest(&request.BaseRequest, "rum", APIVersion, "DescribeDataEventUrlV2")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDataEventUrlV2 require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDataEventUrlV2Response()
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
    c.InitBaseRequest(&request.BaseRequest, "rum", APIVersion, "DescribeDataFetchProject")
    
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
    c.InitBaseRequest(&request.BaseRequest, "rum", APIVersion, "DescribeDataFetchUrl")
    
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
    c.InitBaseRequest(&request.BaseRequest, "rum", APIVersion, "DescribeDataFetchUrlInfo")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDataFetchUrlInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDataFetchUrlInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDataFetchUrlV2Request() (request *DescribeDataFetchUrlV2Request) {
    request = &DescribeDataFetchUrlV2Request{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "DescribeDataFetchUrlV2")
    
    
    return
}

func NewDescribeDataFetchUrlV2Response() (response *DescribeDataFetchUrlV2Response) {
    response = &DescribeDataFetchUrlV2Response{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDataFetchUrlV2
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
func (c *Client) DescribeDataFetchUrlV2(request *DescribeDataFetchUrlV2Request) (response *DescribeDataFetchUrlV2Response, err error) {
    return c.DescribeDataFetchUrlV2WithContext(context.Background(), request)
}

// DescribeDataFetchUrlV2
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
func (c *Client) DescribeDataFetchUrlV2WithContext(ctx context.Context, request *DescribeDataFetchUrlV2Request) (response *DescribeDataFetchUrlV2Response, err error) {
    if request == nil {
        request = NewDescribeDataFetchUrlV2Request()
    }
    c.InitBaseRequest(&request.BaseRequest, "rum", APIVersion, "DescribeDataFetchUrlV2")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDataFetchUrlV2 require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDataFetchUrlV2Response()
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
    c.InitBaseRequest(&request.BaseRequest, "rum", APIVersion, "DescribeDataLogUrlInfo")
    
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
    c.InitBaseRequest(&request.BaseRequest, "rum", APIVersion, "DescribeDataLogUrlStatistics")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDataLogUrlStatistics require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDataLogUrlStatisticsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDataLogUrlStatisticsV2Request() (request *DescribeDataLogUrlStatisticsV2Request) {
    request = &DescribeDataLogUrlStatisticsV2Request{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "DescribeDataLogUrlStatisticsV2")
    
    
    return
}

func NewDescribeDataLogUrlStatisticsV2Response() (response *DescribeDataLogUrlStatisticsV2Response) {
    response = &DescribeDataLogUrlStatisticsV2Response{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDataLogUrlStatisticsV2
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
func (c *Client) DescribeDataLogUrlStatisticsV2(request *DescribeDataLogUrlStatisticsV2Request) (response *DescribeDataLogUrlStatisticsV2Response, err error) {
    return c.DescribeDataLogUrlStatisticsV2WithContext(context.Background(), request)
}

// DescribeDataLogUrlStatisticsV2
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
func (c *Client) DescribeDataLogUrlStatisticsV2WithContext(ctx context.Context, request *DescribeDataLogUrlStatisticsV2Request) (response *DescribeDataLogUrlStatisticsV2Response, err error) {
    if request == nil {
        request = NewDescribeDataLogUrlStatisticsV2Request()
    }
    c.InitBaseRequest(&request.BaseRequest, "rum", APIVersion, "DescribeDataLogUrlStatisticsV2")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDataLogUrlStatisticsV2 require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDataLogUrlStatisticsV2Response()
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
    c.InitBaseRequest(&request.BaseRequest, "rum", APIVersion, "DescribeDataPerformancePage")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDataPerformancePage require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDataPerformancePageResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDataPerformancePageV2Request() (request *DescribeDataPerformancePageV2Request) {
    request = &DescribeDataPerformancePageV2Request{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "DescribeDataPerformancePageV2")
    
    
    return
}

func NewDescribeDataPerformancePageV2Response() (response *DescribeDataPerformancePageV2Response) {
    response = &DescribeDataPerformancePageV2Response{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDataPerformancePageV2
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
func (c *Client) DescribeDataPerformancePageV2(request *DescribeDataPerformancePageV2Request) (response *DescribeDataPerformancePageV2Response, err error) {
    return c.DescribeDataPerformancePageV2WithContext(context.Background(), request)
}

// DescribeDataPerformancePageV2
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
func (c *Client) DescribeDataPerformancePageV2WithContext(ctx context.Context, request *DescribeDataPerformancePageV2Request) (response *DescribeDataPerformancePageV2Response, err error) {
    if request == nil {
        request = NewDescribeDataPerformancePageV2Request()
    }
    c.InitBaseRequest(&request.BaseRequest, "rum", APIVersion, "DescribeDataPerformancePageV2")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDataPerformancePageV2 require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDataPerformancePageV2Response()
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
    c.InitBaseRequest(&request.BaseRequest, "rum", APIVersion, "DescribeDataPvUrlInfo")
    
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
    c.InitBaseRequest(&request.BaseRequest, "rum", APIVersion, "DescribeDataPvUrlStatistics")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDataPvUrlStatistics require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDataPvUrlStatisticsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDataPvUrlStatisticsV2Request() (request *DescribeDataPvUrlStatisticsV2Request) {
    request = &DescribeDataPvUrlStatisticsV2Request{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "DescribeDataPvUrlStatisticsV2")
    
    
    return
}

func NewDescribeDataPvUrlStatisticsV2Response() (response *DescribeDataPvUrlStatisticsV2Response) {
    response = &DescribeDataPvUrlStatisticsV2Response{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDataPvUrlStatisticsV2
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
func (c *Client) DescribeDataPvUrlStatisticsV2(request *DescribeDataPvUrlStatisticsV2Request) (response *DescribeDataPvUrlStatisticsV2Response, err error) {
    return c.DescribeDataPvUrlStatisticsV2WithContext(context.Background(), request)
}

// DescribeDataPvUrlStatisticsV2
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
func (c *Client) DescribeDataPvUrlStatisticsV2WithContext(ctx context.Context, request *DescribeDataPvUrlStatisticsV2Request) (response *DescribeDataPvUrlStatisticsV2Response, err error) {
    if request == nil {
        request = NewDescribeDataPvUrlStatisticsV2Request()
    }
    c.InitBaseRequest(&request.BaseRequest, "rum", APIVersion, "DescribeDataPvUrlStatisticsV2")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDataPvUrlStatisticsV2 require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDataPvUrlStatisticsV2Response()
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
    c.InitBaseRequest(&request.BaseRequest, "rum", APIVersion, "DescribeDataReportCount")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDataReportCount require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDataReportCountResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDataReportCountV2Request() (request *DescribeDataReportCountV2Request) {
    request = &DescribeDataReportCountV2Request{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "DescribeDataReportCountV2")
    
    
    return
}

func NewDescribeDataReportCountV2Response() (response *DescribeDataReportCountV2Response) {
    response = &DescribeDataReportCountV2Response{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDataReportCountV2
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
func (c *Client) DescribeDataReportCountV2(request *DescribeDataReportCountV2Request) (response *DescribeDataReportCountV2Response, err error) {
    return c.DescribeDataReportCountV2WithContext(context.Background(), request)
}

// DescribeDataReportCountV2
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
func (c *Client) DescribeDataReportCountV2WithContext(ctx context.Context, request *DescribeDataReportCountV2Request) (response *DescribeDataReportCountV2Response, err error) {
    if request == nil {
        request = NewDescribeDataReportCountV2Request()
    }
    c.InitBaseRequest(&request.BaseRequest, "rum", APIVersion, "DescribeDataReportCountV2")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDataReportCountV2 require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDataReportCountV2Response()
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
    c.InitBaseRequest(&request.BaseRequest, "rum", APIVersion, "DescribeDataSetUrlStatistics")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDataSetUrlStatistics require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDataSetUrlStatisticsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDataSetUrlStatisticsV2Request() (request *DescribeDataSetUrlStatisticsV2Request) {
    request = &DescribeDataSetUrlStatisticsV2Request{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "DescribeDataSetUrlStatisticsV2")
    
    
    return
}

func NewDescribeDataSetUrlStatisticsV2Response() (response *DescribeDataSetUrlStatisticsV2Response) {
    response = &DescribeDataSetUrlStatisticsV2Response{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDataSetUrlStatisticsV2
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
func (c *Client) DescribeDataSetUrlStatisticsV2(request *DescribeDataSetUrlStatisticsV2Request) (response *DescribeDataSetUrlStatisticsV2Response, err error) {
    return c.DescribeDataSetUrlStatisticsV2WithContext(context.Background(), request)
}

// DescribeDataSetUrlStatisticsV2
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
func (c *Client) DescribeDataSetUrlStatisticsV2WithContext(ctx context.Context, request *DescribeDataSetUrlStatisticsV2Request) (response *DescribeDataSetUrlStatisticsV2Response, err error) {
    if request == nil {
        request = NewDescribeDataSetUrlStatisticsV2Request()
    }
    c.InitBaseRequest(&request.BaseRequest, "rum", APIVersion, "DescribeDataSetUrlStatisticsV2")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDataSetUrlStatisticsV2 require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDataSetUrlStatisticsV2Response()
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
    c.InitBaseRequest(&request.BaseRequest, "rum", APIVersion, "DescribeDataStaticProject")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDataStaticProject require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDataStaticProjectResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDataStaticProjectV2Request() (request *DescribeDataStaticProjectV2Request) {
    request = &DescribeDataStaticProjectV2Request{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "DescribeDataStaticProjectV2")
    
    
    return
}

func NewDescribeDataStaticProjectV2Response() (response *DescribeDataStaticProjectV2Response) {
    response = &DescribeDataStaticProjectV2Response{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDataStaticProjectV2
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
func (c *Client) DescribeDataStaticProjectV2(request *DescribeDataStaticProjectV2Request) (response *DescribeDataStaticProjectV2Response, err error) {
    return c.DescribeDataStaticProjectV2WithContext(context.Background(), request)
}

// DescribeDataStaticProjectV2
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
func (c *Client) DescribeDataStaticProjectV2WithContext(ctx context.Context, request *DescribeDataStaticProjectV2Request) (response *DescribeDataStaticProjectV2Response, err error) {
    if request == nil {
        request = NewDescribeDataStaticProjectV2Request()
    }
    c.InitBaseRequest(&request.BaseRequest, "rum", APIVersion, "DescribeDataStaticProjectV2")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDataStaticProjectV2 require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDataStaticProjectV2Response()
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
    c.InitBaseRequest(&request.BaseRequest, "rum", APIVersion, "DescribeDataStaticResource")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDataStaticResource require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDataStaticResourceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDataStaticResourceV2Request() (request *DescribeDataStaticResourceV2Request) {
    request = &DescribeDataStaticResourceV2Request{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "DescribeDataStaticResourceV2")
    
    
    return
}

func NewDescribeDataStaticResourceV2Response() (response *DescribeDataStaticResourceV2Response) {
    response = &DescribeDataStaticResourceV2Response{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDataStaticResourceV2
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
func (c *Client) DescribeDataStaticResourceV2(request *DescribeDataStaticResourceV2Request) (response *DescribeDataStaticResourceV2Response, err error) {
    return c.DescribeDataStaticResourceV2WithContext(context.Background(), request)
}

// DescribeDataStaticResourceV2
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
func (c *Client) DescribeDataStaticResourceV2WithContext(ctx context.Context, request *DescribeDataStaticResourceV2Request) (response *DescribeDataStaticResourceV2Response, err error) {
    if request == nil {
        request = NewDescribeDataStaticResourceV2Request()
    }
    c.InitBaseRequest(&request.BaseRequest, "rum", APIVersion, "DescribeDataStaticResourceV2")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDataStaticResourceV2 require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDataStaticResourceV2Response()
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
    c.InitBaseRequest(&request.BaseRequest, "rum", APIVersion, "DescribeDataStaticUrl")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDataStaticUrl require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDataStaticUrlResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDataStaticUrlV2Request() (request *DescribeDataStaticUrlV2Request) {
    request = &DescribeDataStaticUrlV2Request{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "DescribeDataStaticUrlV2")
    
    
    return
}

func NewDescribeDataStaticUrlV2Response() (response *DescribeDataStaticUrlV2Response) {
    response = &DescribeDataStaticUrlV2Response{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDataStaticUrlV2
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
func (c *Client) DescribeDataStaticUrlV2(request *DescribeDataStaticUrlV2Request) (response *DescribeDataStaticUrlV2Response, err error) {
    return c.DescribeDataStaticUrlV2WithContext(context.Background(), request)
}

// DescribeDataStaticUrlV2
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
func (c *Client) DescribeDataStaticUrlV2WithContext(ctx context.Context, request *DescribeDataStaticUrlV2Request) (response *DescribeDataStaticUrlV2Response, err error) {
    if request == nil {
        request = NewDescribeDataStaticUrlV2Request()
    }
    c.InitBaseRequest(&request.BaseRequest, "rum", APIVersion, "DescribeDataStaticUrlV2")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDataStaticUrlV2 require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDataStaticUrlV2Response()
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
    c.InitBaseRequest(&request.BaseRequest, "rum", APIVersion, "DescribeDataWebVitalsPage")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDataWebVitalsPage require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDataWebVitalsPageResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDataWebVitalsPageV2Request() (request *DescribeDataWebVitalsPageV2Request) {
    request = &DescribeDataWebVitalsPageV2Request{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "DescribeDataWebVitalsPageV2")
    
    
    return
}

func NewDescribeDataWebVitalsPageV2Response() (response *DescribeDataWebVitalsPageV2Response) {
    response = &DescribeDataWebVitalsPageV2Response{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDataWebVitalsPageV2
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
func (c *Client) DescribeDataWebVitalsPageV2(request *DescribeDataWebVitalsPageV2Request) (response *DescribeDataWebVitalsPageV2Response, err error) {
    return c.DescribeDataWebVitalsPageV2WithContext(context.Background(), request)
}

// DescribeDataWebVitalsPageV2
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
func (c *Client) DescribeDataWebVitalsPageV2WithContext(ctx context.Context, request *DescribeDataWebVitalsPageV2Request) (response *DescribeDataWebVitalsPageV2Response, err error) {
    if request == nil {
        request = NewDescribeDataWebVitalsPageV2Request()
    }
    c.InitBaseRequest(&request.BaseRequest, "rum", APIVersion, "DescribeDataWebVitalsPageV2")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDataWebVitalsPageV2 require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDataWebVitalsPageV2Response()
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
    c.InitBaseRequest(&request.BaseRequest, "rum", APIVersion, "DescribeError")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeError require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeErrorResponse()
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
    c.InitBaseRequest(&request.BaseRequest, "rum", APIVersion, "DescribeProjectLimits")
    
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
    c.InitBaseRequest(&request.BaseRequest, "rum", APIVersion, "DescribeProjects")
    
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
    c.InitBaseRequest(&request.BaseRequest, "rum", APIVersion, "DescribePvList")
    
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
    c.InitBaseRequest(&request.BaseRequest, "rum", APIVersion, "DescribeReleaseFileSign")
    
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
    c.InitBaseRequest(&request.BaseRequest, "rum", APIVersion, "DescribeReleaseFiles")
    
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
    c.InitBaseRequest(&request.BaseRequest, "rum", APIVersion, "DescribeRumGroupLog")
    
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
    c.InitBaseRequest(&request.BaseRequest, "rum", APIVersion, "DescribeRumLogExport")
    
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
    c.InitBaseRequest(&request.BaseRequest, "rum", APIVersion, "DescribeRumLogExports")
    
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
    c.InitBaseRequest(&request.BaseRequest, "rum", APIVersion, "DescribeRumLogList")
    
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
    c.InitBaseRequest(&request.BaseRequest, "rum", APIVersion, "DescribeRumStatsLogList")
    
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
    c.InitBaseRequest(&request.BaseRequest, "rum", APIVersion, "DescribeScores")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeScores require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeScoresResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeScoresV2Request() (request *DescribeScoresV2Request) {
    request = &DescribeScoresV2Request{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "DescribeScoresV2")
    
    
    return
}

func NewDescribeScoresV2Response() (response *DescribeScoresV2Response) {
    response = &DescribeScoresV2Response{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeScoresV2
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
func (c *Client) DescribeScoresV2(request *DescribeScoresV2Request) (response *DescribeScoresV2Response, err error) {
    return c.DescribeScoresV2WithContext(context.Background(), request)
}

// DescribeScoresV2
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
func (c *Client) DescribeScoresV2WithContext(ctx context.Context, request *DescribeScoresV2Request) (response *DescribeScoresV2Response, err error) {
    if request == nil {
        request = NewDescribeScoresV2Request()
    }
    c.InitBaseRequest(&request.BaseRequest, "rum", APIVersion, "DescribeScoresV2")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeScoresV2 require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeScoresV2Response()
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
    c.InitBaseRequest(&request.BaseRequest, "rum", APIVersion, "DescribeTawAreas")
    
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
    c.InitBaseRequest(&request.BaseRequest, "rum", APIVersion, "DescribeTawInstances")
    
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
    c.InitBaseRequest(&request.BaseRequest, "rum", APIVersion, "DescribeUvList")
    
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
    c.InitBaseRequest(&request.BaseRequest, "rum", APIVersion, "DescribeWhitelists")
    
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
    c.InitBaseRequest(&request.BaseRequest, "rum", APIVersion, "ModifyInstance")
    
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
    c.InitBaseRequest(&request.BaseRequest, "rum", APIVersion, "ModifyProject")
    
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
    c.InitBaseRequest(&request.BaseRequest, "rum", APIVersion, "ModifyProjectLimit")
    
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
    c.InitBaseRequest(&request.BaseRequest, "rum", APIVersion, "ResumeInstance")
    
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
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ResumeProject(request *ResumeProjectRequest) (response *ResumeProjectResponse, err error) {
    return c.ResumeProjectWithContext(context.Background(), request)
}

// ResumeProject
// 恢复应用使用与上报数据
//
// 可能返回的错误码:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ResumeProjectWithContext(ctx context.Context, request *ResumeProjectRequest) (response *ResumeProjectResponse, err error) {
    if request == nil {
        request = NewResumeProjectRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "rum", APIVersion, "ResumeProject")
    
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
    c.InitBaseRequest(&request.BaseRequest, "rum", APIVersion, "StopInstance")
    
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
    c.InitBaseRequest(&request.BaseRequest, "rum", APIVersion, "StopProject")
    
    if c.GetCredential() == nil {
        return nil, errors.New("StopProject require credential")
    }

    request.SetContext(ctx)
    
    response = NewStopProjectResponse()
    err = c.Send(request, response)
    return
}
