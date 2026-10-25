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


func NewCreateFieldTransformRuleRequest() (request *CreateFieldTransformRuleRequest) {
    request = &CreateFieldTransformRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "CreateFieldTransformRule")
    
    
    return
}

func NewCreateFieldTransformRuleResponse() (response *CreateFieldTransformRuleResponse) {
    response = &CreateFieldTransformRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateFieldTransformRule
// 创建字段配置规则
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CHARGENOBALANCE = "FailedOperation.ChargeNoBalance"
//  FAILEDOPERATION_CHARGENOPAYRIGHT = "FailedOperation.ChargeNoPayRight"
//  FAILEDOPERATION_CHARGEPARAMINVALID = "FailedOperation.ChargeParamInvalid"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  FAILEDOPERATION_INFRASTRUCTUREERROR = "FailedOperation.InfrastructureError"
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
//  RESOURCENOTFOUND_NOINSTANCE = "ResourceNotFound.NoInstance"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateFieldTransformRule(request *CreateFieldTransformRuleRequest) (response *CreateFieldTransformRuleResponse, err error) {
    return c.CreateFieldTransformRuleWithContext(context.Background(), request)
}

// CreateFieldTransformRule
// 创建字段配置规则
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CHARGENOBALANCE = "FailedOperation.ChargeNoBalance"
//  FAILEDOPERATION_CHARGENOPAYRIGHT = "FailedOperation.ChargeNoPayRight"
//  FAILEDOPERATION_CHARGEPARAMINVALID = "FailedOperation.ChargeParamInvalid"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  FAILEDOPERATION_INFRASTRUCTUREERROR = "FailedOperation.InfrastructureError"
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
//  RESOURCENOTFOUND_NOINSTANCE = "ResourceNotFound.NoInstance"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateFieldTransformRuleWithContext(ctx context.Context, request *CreateFieldTransformRuleRequest) (response *CreateFieldTransformRuleResponse, err error) {
    if request == nil {
        request = NewCreateFieldTransformRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "rum", APIVersion, "CreateFieldTransformRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateFieldTransformRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateFieldTransformRuleResponse()
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

func NewCreateTDEMApplicationRequest() (request *CreateTDEMApplicationRequest) {
    request = &CreateTDEMApplicationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "CreateTDEMApplication")
    
    
    return
}

func NewCreateTDEMApplicationResponse() (response *CreateTDEMApplicationResponse) {
    response = &CreateTDEMApplicationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateTDEMApplication
// 在指定业务系统下创建应用。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERCOMBINATION = "InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) CreateTDEMApplication(request *CreateTDEMApplicationRequest) (response *CreateTDEMApplicationResponse, err error) {
    return c.CreateTDEMApplicationWithContext(context.Background(), request)
}

// CreateTDEMApplication
// 在指定业务系统下创建应用。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERCOMBINATION = "InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) CreateTDEMApplicationWithContext(ctx context.Context, request *CreateTDEMApplicationRequest) (response *CreateTDEMApplicationResponse, err error) {
    if request == nil {
        request = NewCreateTDEMApplicationRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "rum", APIVersion, "CreateTDEMApplication")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateTDEMApplication require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateTDEMApplicationResponse()
    err = c.Send(request, response)
    return
}

func NewCreateTDEMBusinessSystemRequest() (request *CreateTDEMBusinessSystemRequest) {
    request = &CreateTDEMBusinessSystemRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "CreateTDEMBusinessSystem")
    
    
    return
}

func NewCreateTDEMBusinessSystemResponse() (response *CreateTDEMBusinessSystemResponse) {
    response = &CreateTDEMBusinessSystemResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateTDEMBusinessSystem
// 创建业务系统。InstanceId 由后端生成，格式 rum-<12 位随机字母数字>。重名返回 InvalidParameterValue。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERCOMBINATION = "InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) CreateTDEMBusinessSystem(request *CreateTDEMBusinessSystemRequest) (response *CreateTDEMBusinessSystemResponse, err error) {
    return c.CreateTDEMBusinessSystemWithContext(context.Background(), request)
}

// CreateTDEMBusinessSystem
// 创建业务系统。InstanceId 由后端生成，格式 rum-<12 位随机字母数字>。重名返回 InvalidParameterValue。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERCOMBINATION = "InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) CreateTDEMBusinessSystemWithContext(ctx context.Context, request *CreateTDEMBusinessSystemRequest) (response *CreateTDEMBusinessSystemResponse, err error) {
    if request == nil {
        request = NewCreateTDEMBusinessSystemRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "rum", APIVersion, "CreateTDEMBusinessSystem")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateTDEMBusinessSystem require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateTDEMBusinessSystemResponse()
    err = c.Send(request, response)
    return
}

func NewCreateTDEMDialTaskRequest() (request *CreateTDEMDialTaskRequest) {
    request = &CreateTDEMDialTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "CreateTDEMDialTask")
    
    
    return
}

func NewCreateTDEMDialTaskResponse() (response *CreateTDEMDialTaskResponse) {
    response = &CreateTDEMDialTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateTDEMDialTask
// 创建一条移动拨测任务。同一应用处于 active 的任务最多 50 条。四项启动范围都为空时不限制启动；填写多项时必须同时满足，同一项多个值命中任意一个即可。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateTDEMDialTask(request *CreateTDEMDialTaskRequest) (response *CreateTDEMDialTaskResponse, err error) {
    return c.CreateTDEMDialTaskWithContext(context.Background(), request)
}

// CreateTDEMDialTask
// 创建一条移动拨测任务。同一应用处于 active 的任务最多 50 条。四项启动范围都为空时不限制启动；填写多项时必须同时满足，同一项多个值命中任意一个即可。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateTDEMDialTaskWithContext(ctx context.Context, request *CreateTDEMDialTaskRequest) (response *CreateTDEMDialTaskResponse, err error) {
    if request == nil {
        request = NewCreateTDEMDialTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "rum", APIVersion, "CreateTDEMDialTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateTDEMDialTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateTDEMDialTaskResponse()
    err = c.Send(request, response)
    return
}

func NewCreateTDEMFieldTransformRuleRequest() (request *CreateTDEMFieldTransformRuleRequest) {
    request = &CreateTDEMFieldTransformRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "CreateTDEMFieldTransformRule")
    
    
    return
}

func NewCreateTDEMFieldTransformRuleResponse() (response *CreateTDEMFieldTransformRuleResponse) {
    response = &CreateTDEMFieldTransformRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateTDEMFieldTransformRule
// 为应用创建字段值替换规则，单应用上限 50 条。替代已废弃的 CreateFieldTransformRule；AppId 为 string。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERCOMBINATION = "InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) CreateTDEMFieldTransformRule(request *CreateTDEMFieldTransformRuleRequest) (response *CreateTDEMFieldTransformRuleResponse, err error) {
    return c.CreateTDEMFieldTransformRuleWithContext(context.Background(), request)
}

// CreateTDEMFieldTransformRule
// 为应用创建字段值替换规则，单应用上限 50 条。替代已废弃的 CreateFieldTransformRule；AppId 为 string。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERCOMBINATION = "InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) CreateTDEMFieldTransformRuleWithContext(ctx context.Context, request *CreateTDEMFieldTransformRuleRequest) (response *CreateTDEMFieldTransformRuleResponse, err error) {
    if request == nil {
        request = NewCreateTDEMFieldTransformRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "rum", APIVersion, "CreateTDEMFieldTransformRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateTDEMFieldTransformRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateTDEMFieldTransformRuleResponse()
    err = c.Send(request, response)
    return
}

func NewCreateTDEMFunnelRequest() (request *CreateTDEMFunnelRequest) {
    request = &CreateTDEMFunnelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "CreateTDEMFunnel")
    
    
    return
}

func NewCreateTDEMFunnelResponse() (response *CreateTDEMFunnelResponse) {
    response = &CreateTDEMFunnelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateTDEMFunnel
// 创建漏斗
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CHARGENOBALANCE = "FailedOperation.ChargeNoBalance"
//  FAILEDOPERATION_CHARGENOPAYRIGHT = "FailedOperation.ChargeNoPayRight"
//  FAILEDOPERATION_CHARGEPARAMINVALID = "FailedOperation.ChargeParamInvalid"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  FAILEDOPERATION_INFRASTRUCTUREERROR = "FailedOperation.InfrastructureError"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERCOMBINATION = "InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
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
func (c *Client) CreateTDEMFunnel(request *CreateTDEMFunnelRequest) (response *CreateTDEMFunnelResponse, err error) {
    return c.CreateTDEMFunnelWithContext(context.Background(), request)
}

// CreateTDEMFunnel
// 创建漏斗
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CHARGENOBALANCE = "FailedOperation.ChargeNoBalance"
//  FAILEDOPERATION_CHARGENOPAYRIGHT = "FailedOperation.ChargeNoPayRight"
//  FAILEDOPERATION_CHARGEPARAMINVALID = "FailedOperation.ChargeParamInvalid"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  FAILEDOPERATION_INFRASTRUCTUREERROR = "FailedOperation.InfrastructureError"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERCOMBINATION = "InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
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
func (c *Client) CreateTDEMFunnelWithContext(ctx context.Context, request *CreateTDEMFunnelRequest) (response *CreateTDEMFunnelResponse, err error) {
    if request == nil {
        request = NewCreateTDEMFunnelRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "rum", APIVersion, "CreateTDEMFunnel")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateTDEMFunnel require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateTDEMFunnelResponse()
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

func NewDeleteFieldTransformRuleRequest() (request *DeleteFieldTransformRuleRequest) {
    request = &DeleteFieldTransformRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "DeleteFieldTransformRule")
    
    
    return
}

func NewDeleteFieldTransformRuleResponse() (response *DeleteFieldTransformRuleResponse) {
    response = &DeleteFieldTransformRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteFieldTransformRule
// 删除字段配置规则
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CHARGENOBALANCE = "FailedOperation.ChargeNoBalance"
//  FAILEDOPERATION_CHARGENOPAYRIGHT = "FailedOperation.ChargeNoPayRight"
//  FAILEDOPERATION_CHARGEPARAMINVALID = "FailedOperation.ChargeParamInvalid"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  FAILEDOPERATION_INFRASTRUCTUREERROR = "FailedOperation.InfrastructureError"
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
//  RESOURCENOTFOUND_NOINSTANCE = "ResourceNotFound.NoInstance"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteFieldTransformRule(request *DeleteFieldTransformRuleRequest) (response *DeleteFieldTransformRuleResponse, err error) {
    return c.DeleteFieldTransformRuleWithContext(context.Background(), request)
}

// DeleteFieldTransformRule
// 删除字段配置规则
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CHARGENOBALANCE = "FailedOperation.ChargeNoBalance"
//  FAILEDOPERATION_CHARGENOPAYRIGHT = "FailedOperation.ChargeNoPayRight"
//  FAILEDOPERATION_CHARGEPARAMINVALID = "FailedOperation.ChargeParamInvalid"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  FAILEDOPERATION_INFRASTRUCTUREERROR = "FailedOperation.InfrastructureError"
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
//  RESOURCENOTFOUND_NOINSTANCE = "ResourceNotFound.NoInstance"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteFieldTransformRuleWithContext(ctx context.Context, request *DeleteFieldTransformRuleRequest) (response *DeleteFieldTransformRuleResponse, err error) {
    if request == nil {
        request = NewDeleteFieldTransformRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "rum", APIVersion, "DeleteFieldTransformRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteFieldTransformRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteFieldTransformRuleResponse()
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

func NewDeleteMiniProgramReplayPackRequest() (request *DeleteMiniProgramReplayPackRequest) {
    request = &DeleteMiniProgramReplayPackRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "DeleteMiniProgramReplayPack")
    
    
    return
}

func NewDeleteMiniProgramReplayPackResponse() (response *DeleteMiniProgramReplayPackResponse) {
    response = &DeleteMiniProgramReplayPackResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteMiniProgramReplayPack
// 删除回放包元数据，并尽力删除COS上回放包与源码zip。不存在返回ResourceNotFound。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CHARGENOBALANCE = "FailedOperation.ChargeNoBalance"
//  FAILEDOPERATION_CHARGENOPAYRIGHT = "FailedOperation.ChargeNoPayRight"
//  FAILEDOPERATION_CHARGEPARAMINVALID = "FailedOperation.ChargeParamInvalid"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  FAILEDOPERATION_INFRASTRUCTUREERROR = "FailedOperation.InfrastructureError"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERCOMBINATION = "InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
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
func (c *Client) DeleteMiniProgramReplayPack(request *DeleteMiniProgramReplayPackRequest) (response *DeleteMiniProgramReplayPackResponse, err error) {
    return c.DeleteMiniProgramReplayPackWithContext(context.Background(), request)
}

// DeleteMiniProgramReplayPack
// 删除回放包元数据，并尽力删除COS上回放包与源码zip。不存在返回ResourceNotFound。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CHARGENOBALANCE = "FailedOperation.ChargeNoBalance"
//  FAILEDOPERATION_CHARGENOPAYRIGHT = "FailedOperation.ChargeNoPayRight"
//  FAILEDOPERATION_CHARGEPARAMINVALID = "FailedOperation.ChargeParamInvalid"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  FAILEDOPERATION_INFRASTRUCTUREERROR = "FailedOperation.InfrastructureError"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERCOMBINATION = "InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
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
func (c *Client) DeleteMiniProgramReplayPackWithContext(ctx context.Context, request *DeleteMiniProgramReplayPackRequest) (response *DeleteMiniProgramReplayPackResponse, err error) {
    if request == nil {
        request = NewDeleteMiniProgramReplayPackRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "rum", APIVersion, "DeleteMiniProgramReplayPack")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteMiniProgramReplayPack require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteMiniProgramReplayPackResponse()
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

func NewDeleteSymbolArtifactRequest() (request *DeleteSymbolArtifactRequest) {
    request = &DeleteSymbolArtifactRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "DeleteSymbolArtifact")
    
    
    return
}

func NewDeleteSymbolArtifactResponse() (response *DeleteSymbolArtifactResponse) {
    response = &DeleteSymbolArtifactResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteSymbolArtifact
// 删除指定的符号制品。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CHARGENOBALANCE = "FailedOperation.ChargeNoBalance"
//  FAILEDOPERATION_CHARGENOPAYRIGHT = "FailedOperation.ChargeNoPayRight"
//  FAILEDOPERATION_CHARGEPARAMINVALID = "FailedOperation.ChargeParamInvalid"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  FAILEDOPERATION_INFRASTRUCTUREERROR = "FailedOperation.InfrastructureError"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERCOMBINATION = "InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
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
func (c *Client) DeleteSymbolArtifact(request *DeleteSymbolArtifactRequest) (response *DeleteSymbolArtifactResponse, err error) {
    return c.DeleteSymbolArtifactWithContext(context.Background(), request)
}

// DeleteSymbolArtifact
// 删除指定的符号制品。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CHARGENOBALANCE = "FailedOperation.ChargeNoBalance"
//  FAILEDOPERATION_CHARGENOPAYRIGHT = "FailedOperation.ChargeNoPayRight"
//  FAILEDOPERATION_CHARGEPARAMINVALID = "FailedOperation.ChargeParamInvalid"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  FAILEDOPERATION_INFRASTRUCTUREERROR = "FailedOperation.InfrastructureError"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERCOMBINATION = "InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
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
func (c *Client) DeleteSymbolArtifactWithContext(ctx context.Context, request *DeleteSymbolArtifactRequest) (response *DeleteSymbolArtifactResponse, err error) {
    if request == nil {
        request = NewDeleteSymbolArtifactRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "rum", APIVersion, "DeleteSymbolArtifact")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteSymbolArtifact require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteSymbolArtifactResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteTDEMApplicationRequest() (request *DeleteTDEMApplicationRequest) {
    request = &DeleteTDEMApplicationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "DeleteTDEMApplication")
    
    
    return
}

func NewDeleteTDEMApplicationResponse() (response *DeleteTDEMApplicationResponse) {
    response = &DeleteTDEMApplicationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteTDEMApplication
// 逻辑删除应用，ProjectKey 即刻失效。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERCOMBINATION = "InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DeleteTDEMApplication(request *DeleteTDEMApplicationRequest) (response *DeleteTDEMApplicationResponse, err error) {
    return c.DeleteTDEMApplicationWithContext(context.Background(), request)
}

// DeleteTDEMApplication
// 逻辑删除应用，ProjectKey 即刻失效。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERCOMBINATION = "InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DeleteTDEMApplicationWithContext(ctx context.Context, request *DeleteTDEMApplicationRequest) (response *DeleteTDEMApplicationResponse, err error) {
    if request == nil {
        request = NewDeleteTDEMApplicationRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "rum", APIVersion, "DeleteTDEMApplication")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteTDEMApplication require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteTDEMApplicationResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteTDEMBusinessSystemRequest() (request *DeleteTDEMBusinessSystemRequest) {
    request = &DeleteTDEMBusinessSystemRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "DeleteTDEMBusinessSystem")
    
    
    return
}

func NewDeleteTDEMBusinessSystemResponse() (response *DeleteTDEMBusinessSystemResponse) {
    response = &DeleteTDEMBusinessSystemResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteTDEMBusinessSystem
// 删除业务系统。名下存在未删除应用时返回 ResourceInUse。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERCOMBINATION = "InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DeleteTDEMBusinessSystem(request *DeleteTDEMBusinessSystemRequest) (response *DeleteTDEMBusinessSystemResponse, err error) {
    return c.DeleteTDEMBusinessSystemWithContext(context.Background(), request)
}

// DeleteTDEMBusinessSystem
// 删除业务系统。名下存在未删除应用时返回 ResourceInUse。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERCOMBINATION = "InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DeleteTDEMBusinessSystemWithContext(ctx context.Context, request *DeleteTDEMBusinessSystemRequest) (response *DeleteTDEMBusinessSystemResponse, err error) {
    if request == nil {
        request = NewDeleteTDEMBusinessSystemRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "rum", APIVersion, "DeleteTDEMBusinessSystem")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteTDEMBusinessSystem require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteTDEMBusinessSystemResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteTDEMDialTaskRequest() (request *DeleteTDEMDialTaskRequest) {
    request = &DeleteTDEMDialTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "DeleteTDEMDialTask")
    
    
    return
}

func NewDeleteTDEMDialTaskResponse() (response *DeleteTDEMDialTaskResponse) {
    response = &DeleteTDEMDialTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteTDEMDialTask
// 将拨测任务标记为已删除。已删除的任务再次删除视为成功。不物理删除任务行，也不删除历史拨测结果。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteTDEMDialTask(request *DeleteTDEMDialTaskRequest) (response *DeleteTDEMDialTaskResponse, err error) {
    return c.DeleteTDEMDialTaskWithContext(context.Background(), request)
}

// DeleteTDEMDialTask
// 将拨测任务标记为已删除。已删除的任务再次删除视为成功。不物理删除任务行，也不删除历史拨测结果。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteTDEMDialTaskWithContext(ctx context.Context, request *DeleteTDEMDialTaskRequest) (response *DeleteTDEMDialTaskResponse, err error) {
    if request == nil {
        request = NewDeleteTDEMDialTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "rum", APIVersion, "DeleteTDEMDialTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteTDEMDialTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteTDEMDialTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteTDEMFieldTransformRuleRequest() (request *DeleteTDEMFieldTransformRuleRequest) {
    request = &DeleteTDEMFieldTransformRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "DeleteTDEMFieldTransformRule")
    
    
    return
}

func NewDeleteTDEMFieldTransformRuleResponse() (response *DeleteTDEMFieldTransformRuleResponse) {
    response = &DeleteTDEMFieldTransformRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteTDEMFieldTransformRule
// 按规则 Id 删除字段值替换规则。替代已废弃的 DeleteFieldTransformRule；AppId 为 string。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERCOMBINATION = "InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DeleteTDEMFieldTransformRule(request *DeleteTDEMFieldTransformRuleRequest) (response *DeleteTDEMFieldTransformRuleResponse, err error) {
    return c.DeleteTDEMFieldTransformRuleWithContext(context.Background(), request)
}

// DeleteTDEMFieldTransformRule
// 按规则 Id 删除字段值替换规则。替代已废弃的 DeleteFieldTransformRule；AppId 为 string。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERCOMBINATION = "InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DeleteTDEMFieldTransformRuleWithContext(ctx context.Context, request *DeleteTDEMFieldTransformRuleRequest) (response *DeleteTDEMFieldTransformRuleResponse, err error) {
    if request == nil {
        request = NewDeleteTDEMFieldTransformRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "rum", APIVersion, "DeleteTDEMFieldTransformRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteTDEMFieldTransformRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteTDEMFieldTransformRuleResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteTDEMFunnelRequest() (request *DeleteTDEMFunnelRequest) {
    request = &DeleteTDEMFunnelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "DeleteTDEMFunnel")
    
    
    return
}

func NewDeleteTDEMFunnelResponse() (response *DeleteTDEMFunnelResponse) {
    response = &DeleteTDEMFunnelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteTDEMFunnel
// 删除漏斗
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CHARGENOBALANCE = "FailedOperation.ChargeNoBalance"
//  FAILEDOPERATION_CHARGENOPAYRIGHT = "FailedOperation.ChargeNoPayRight"
//  FAILEDOPERATION_CHARGEPARAMINVALID = "FailedOperation.ChargeParamInvalid"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  FAILEDOPERATION_INFRASTRUCTUREERROR = "FailedOperation.InfrastructureError"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERCOMBINATION = "InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
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
func (c *Client) DeleteTDEMFunnel(request *DeleteTDEMFunnelRequest) (response *DeleteTDEMFunnelResponse, err error) {
    return c.DeleteTDEMFunnelWithContext(context.Background(), request)
}

// DeleteTDEMFunnel
// 删除漏斗
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CHARGENOBALANCE = "FailedOperation.ChargeNoBalance"
//  FAILEDOPERATION_CHARGENOPAYRIGHT = "FailedOperation.ChargeNoPayRight"
//  FAILEDOPERATION_CHARGEPARAMINVALID = "FailedOperation.ChargeParamInvalid"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  FAILEDOPERATION_INFRASTRUCTUREERROR = "FailedOperation.InfrastructureError"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERCOMBINATION = "InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
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
func (c *Client) DeleteTDEMFunnelWithContext(ctx context.Context, request *DeleteTDEMFunnelRequest) (response *DeleteTDEMFunnelResponse, err error) {
    if request == nil {
        request = NewDeleteTDEMFunnelRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "rum", APIVersion, "DeleteTDEMFunnel")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteTDEMFunnel require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteTDEMFunnelResponse()
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

func NewDescribeApplicationExitReportDetailRequest() (request *DescribeApplicationExitReportDetailRequest) {
    request = &DescribeApplicationExitReportDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "DescribeApplicationExitReportDetail")
    
    
    return
}

func NewDescribeApplicationExitReportDetailResponse() (response *DescribeApplicationExitReportDetailResponse) {
    response = &DescribeApplicationExitReportDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeApplicationExitReportDetail
// 获取应用退出个例上报详情
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeApplicationExitReportDetail(request *DescribeApplicationExitReportDetailRequest) (response *DescribeApplicationExitReportDetailResponse, err error) {
    return c.DescribeApplicationExitReportDetailWithContext(context.Background(), request)
}

// DescribeApplicationExitReportDetail
// 获取应用退出个例上报详情
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeApplicationExitReportDetailWithContext(ctx context.Context, request *DescribeApplicationExitReportDetailRequest) (response *DescribeApplicationExitReportDetailResponse, err error) {
    if request == nil {
        request = NewDescribeApplicationExitReportDetailRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "rum", APIVersion, "DescribeApplicationExitReportDetail")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeApplicationExitReportDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeApplicationExitReportDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeApplicationExitReportListRequest() (request *DescribeApplicationExitReportListRequest) {
    request = &DescribeApplicationExitReportListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "DescribeApplicationExitReportList")
    
    
    return
}

func NewDescribeApplicationExitReportListResponse() (response *DescribeApplicationExitReportListResponse) {
    response = &DescribeApplicationExitReportListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeApplicationExitReportList
// 获取应用退出个例上报列表
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeApplicationExitReportList(request *DescribeApplicationExitReportListRequest) (response *DescribeApplicationExitReportListResponse, err error) {
    return c.DescribeApplicationExitReportListWithContext(context.Background(), request)
}

// DescribeApplicationExitReportList
// 获取应用退出个例上报列表
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeApplicationExitReportListWithContext(ctx context.Context, request *DescribeApplicationExitReportListRequest) (response *DescribeApplicationExitReportListResponse, err error) {
    if request == nil {
        request = NewDescribeApplicationExitReportListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "rum", APIVersion, "DescribeApplicationExitReportList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeApplicationExitReportList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeApplicationExitReportListResponse()
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

func NewDescribeExceptionDetailRequest() (request *DescribeExceptionDetailRequest) {
    request = &DescribeExceptionDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "DescribeExceptionDetail")
    
    
    return
}

func NewDescribeExceptionDetailResponse() (response *DescribeExceptionDetailResponse) {
    response = &DescribeExceptionDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeExceptionDetail
// 获取异常问题详情数据
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeExceptionDetail(request *DescribeExceptionDetailRequest) (response *DescribeExceptionDetailResponse, err error) {
    return c.DescribeExceptionDetailWithContext(context.Background(), request)
}

// DescribeExceptionDetail
// 获取异常问题详情数据
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeExceptionDetailWithContext(ctx context.Context, request *DescribeExceptionDetailRequest) (response *DescribeExceptionDetailResponse, err error) {
    if request == nil {
        request = NewDescribeExceptionDetailRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "rum", APIVersion, "DescribeExceptionDetail")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeExceptionDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeExceptionDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeExceptionReportListRequest() (request *DescribeExceptionReportListRequest) {
    request = &DescribeExceptionReportListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "DescribeExceptionReportList")
    
    
    return
}

func NewDescribeExceptionReportListResponse() (response *DescribeExceptionReportListResponse) {
    response = &DescribeExceptionReportListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeExceptionReportList
// 获取异常个例列表
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeExceptionReportList(request *DescribeExceptionReportListRequest) (response *DescribeExceptionReportListResponse, err error) {
    return c.DescribeExceptionReportListWithContext(context.Background(), request)
}

// DescribeExceptionReportList
// 获取异常个例列表
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeExceptionReportListWithContext(ctx context.Context, request *DescribeExceptionReportListRequest) (response *DescribeExceptionReportListResponse, err error) {
    if request == nil {
        request = NewDescribeExceptionReportListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "rum", APIVersion, "DescribeExceptionReportList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeExceptionReportList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeExceptionReportListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeFOOMMallocProblemDetailRequest() (request *DescribeFOOMMallocProblemDetailRequest) {
    request = &DescribeFOOMMallocProblemDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "DescribeFOOMMallocProblemDetail")
    
    
    return
}

func NewDescribeFOOMMallocProblemDetailResponse() (response *DescribeFOOMMallocProblemDetailResponse) {
    response = &DescribeFOOMMallocProblemDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeFOOMMallocProblemDetail
// 获取FOOM内存分配问题详情数据
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeFOOMMallocProblemDetail(request *DescribeFOOMMallocProblemDetailRequest) (response *DescribeFOOMMallocProblemDetailResponse, err error) {
    return c.DescribeFOOMMallocProblemDetailWithContext(context.Background(), request)
}

// DescribeFOOMMallocProblemDetail
// 获取FOOM内存分配问题详情数据
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeFOOMMallocProblemDetailWithContext(ctx context.Context, request *DescribeFOOMMallocProblemDetailRequest) (response *DescribeFOOMMallocProblemDetailResponse, err error) {
    if request == nil {
        request = NewDescribeFOOMMallocProblemDetailRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "rum", APIVersion, "DescribeFOOMMallocProblemDetail")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeFOOMMallocProblemDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeFOOMMallocProblemDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeFOOMMallocProblemListRequest() (request *DescribeFOOMMallocProblemListRequest) {
    request = &DescribeFOOMMallocProblemListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "DescribeFOOMMallocProblemList")
    
    
    return
}

func NewDescribeFOOMMallocProblemListResponse() (response *DescribeFOOMMallocProblemListResponse) {
    response = &DescribeFOOMMallocProblemListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeFOOMMallocProblemList
// 获取FOOM内存分配问题列表数据
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeFOOMMallocProblemList(request *DescribeFOOMMallocProblemListRequest) (response *DescribeFOOMMallocProblemListResponse, err error) {
    return c.DescribeFOOMMallocProblemListWithContext(context.Background(), request)
}

// DescribeFOOMMallocProblemList
// 获取FOOM内存分配问题列表数据
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeFOOMMallocProblemListWithContext(ctx context.Context, request *DescribeFOOMMallocProblemListRequest) (response *DescribeFOOMMallocProblemListResponse, err error) {
    if request == nil {
        request = NewDescribeFOOMMallocProblemListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "rum", APIVersion, "DescribeFOOMMallocProblemList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeFOOMMallocProblemList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeFOOMMallocProblemListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeFOOMMallocReportListRequest() (request *DescribeFOOMMallocReportListRequest) {
    request = &DescribeFOOMMallocReportListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "DescribeFOOMMallocReportList")
    
    
    return
}

func NewDescribeFOOMMallocReportListResponse() (response *DescribeFOOMMallocReportListResponse) {
    response = &DescribeFOOMMallocReportListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeFOOMMallocReportList
// 获取FOOM内存分配问题上报列表
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeFOOMMallocReportList(request *DescribeFOOMMallocReportListRequest) (response *DescribeFOOMMallocReportListResponse, err error) {
    return c.DescribeFOOMMallocReportListWithContext(context.Background(), request)
}

// DescribeFOOMMallocReportList
// 获取FOOM内存分配问题上报列表
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeFOOMMallocReportListWithContext(ctx context.Context, request *DescribeFOOMMallocReportListRequest) (response *DescribeFOOMMallocReportListResponse, err error) {
    if request == nil {
        request = NewDescribeFOOMMallocReportListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "rum", APIVersion, "DescribeFOOMMallocReportList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeFOOMMallocReportList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeFOOMMallocReportListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeFOOMProblemDetailRequest() (request *DescribeFOOMProblemDetailRequest) {
    request = &DescribeFOOMProblemDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "DescribeFOOMProblemDetail")
    
    
    return
}

func NewDescribeFOOMProblemDetailResponse() (response *DescribeFOOMProblemDetailResponse) {
    response = &DescribeFOOMProblemDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeFOOMProblemDetail
// 获取FOOM问题详情数据
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeFOOMProblemDetail(request *DescribeFOOMProblemDetailRequest) (response *DescribeFOOMProblemDetailResponse, err error) {
    return c.DescribeFOOMProblemDetailWithContext(context.Background(), request)
}

// DescribeFOOMProblemDetail
// 获取FOOM问题详情数据
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeFOOMProblemDetailWithContext(ctx context.Context, request *DescribeFOOMProblemDetailRequest) (response *DescribeFOOMProblemDetailResponse, err error) {
    if request == nil {
        request = NewDescribeFOOMProblemDetailRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "rum", APIVersion, "DescribeFOOMProblemDetail")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeFOOMProblemDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeFOOMProblemDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeFOOMProblemListRequest() (request *DescribeFOOMProblemListRequest) {
    request = &DescribeFOOMProblemListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "DescribeFOOMProblemList")
    
    
    return
}

func NewDescribeFOOMProblemListResponse() (response *DescribeFOOMProblemListResponse) {
    response = &DescribeFOOMProblemListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeFOOMProblemList
// 获取FOOM问题列表数据
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeFOOMProblemList(request *DescribeFOOMProblemListRequest) (response *DescribeFOOMProblemListResponse, err error) {
    return c.DescribeFOOMProblemListWithContext(context.Background(), request)
}

// DescribeFOOMProblemList
// 获取FOOM问题列表数据
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeFOOMProblemListWithContext(ctx context.Context, request *DescribeFOOMProblemListRequest) (response *DescribeFOOMProblemListResponse, err error) {
    if request == nil {
        request = NewDescribeFOOMProblemListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "rum", APIVersion, "DescribeFOOMProblemList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeFOOMProblemList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeFOOMProblemListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeFOOMReportListRequest() (request *DescribeFOOMReportListRequest) {
    request = &DescribeFOOMReportListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "DescribeFOOMReportList")
    
    
    return
}

func NewDescribeFOOMReportListResponse() (response *DescribeFOOMReportListResponse) {
    response = &DescribeFOOMReportListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeFOOMReportList
// 获取FOOM问题上报列表
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeFOOMReportList(request *DescribeFOOMReportListRequest) (response *DescribeFOOMReportListResponse, err error) {
    return c.DescribeFOOMReportListWithContext(context.Background(), request)
}

// DescribeFOOMReportList
// 获取FOOM问题上报列表
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeFOOMReportListWithContext(ctx context.Context, request *DescribeFOOMReportListRequest) (response *DescribeFOOMReportListResponse, err error) {
    if request == nil {
        request = NewDescribeFOOMReportListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "rum", APIVersion, "DescribeFOOMReportList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeFOOMReportList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeFOOMReportListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeFieldTransformRulesRequest() (request *DescribeFieldTransformRulesRequest) {
    request = &DescribeFieldTransformRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "DescribeFieldTransformRules")
    
    
    return
}

func NewDescribeFieldTransformRulesResponse() (response *DescribeFieldTransformRulesResponse) {
    response = &DescribeFieldTransformRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeFieldTransformRules
// 获取字段配置规则
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CHARGENOBALANCE = "FailedOperation.ChargeNoBalance"
//  FAILEDOPERATION_CHARGENOPAYRIGHT = "FailedOperation.ChargeNoPayRight"
//  FAILEDOPERATION_CHARGEPARAMINVALID = "FailedOperation.ChargeParamInvalid"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  FAILEDOPERATION_INFRASTRUCTUREERROR = "FailedOperation.InfrastructureError"
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
//  RESOURCENOTFOUND_NOINSTANCE = "ResourceNotFound.NoInstance"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeFieldTransformRules(request *DescribeFieldTransformRulesRequest) (response *DescribeFieldTransformRulesResponse, err error) {
    return c.DescribeFieldTransformRulesWithContext(context.Background(), request)
}

// DescribeFieldTransformRules
// 获取字段配置规则
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CHARGENOBALANCE = "FailedOperation.ChargeNoBalance"
//  FAILEDOPERATION_CHARGENOPAYRIGHT = "FailedOperation.ChargeNoPayRight"
//  FAILEDOPERATION_CHARGEPARAMINVALID = "FailedOperation.ChargeParamInvalid"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  FAILEDOPERATION_INFRASTRUCTUREERROR = "FailedOperation.InfrastructureError"
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
//  RESOURCENOTFOUND_NOINSTANCE = "ResourceNotFound.NoInstance"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeFieldTransformRulesWithContext(ctx context.Context, request *DescribeFieldTransformRulesRequest) (response *DescribeFieldTransformRulesResponse, err error) {
    if request == nil {
        request = NewDescribeFieldTransformRulesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "rum", APIVersion, "DescribeFieldTransformRules")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeFieldTransformRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeFieldTransformRulesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeFieldValuesRequest() (request *DescribeFieldValuesRequest) {
    request = &DescribeFieldValuesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "DescribeFieldValues")
    
    
    return
}

func NewDescribeFieldValuesResponse() (response *DescribeFieldValuesResponse) {
    response = &DescribeFieldValuesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeFieldValues
// 获取TDEM字段值列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CHARGENOBALANCE = "FailedOperation.ChargeNoBalance"
//  FAILEDOPERATION_CHARGENOPAYRIGHT = "FailedOperation.ChargeNoPayRight"
//  FAILEDOPERATION_CHARGEPARAMINVALID = "FailedOperation.ChargeParamInvalid"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  FAILEDOPERATION_INFRASTRUCTUREERROR = "FailedOperation.InfrastructureError"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERCOMBINATION = "InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
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
func (c *Client) DescribeFieldValues(request *DescribeFieldValuesRequest) (response *DescribeFieldValuesResponse, err error) {
    return c.DescribeFieldValuesWithContext(context.Background(), request)
}

// DescribeFieldValues
// 获取TDEM字段值列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CHARGENOBALANCE = "FailedOperation.ChargeNoBalance"
//  FAILEDOPERATION_CHARGENOPAYRIGHT = "FailedOperation.ChargeNoPayRight"
//  FAILEDOPERATION_CHARGEPARAMINVALID = "FailedOperation.ChargeParamInvalid"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  FAILEDOPERATION_INFRASTRUCTUREERROR = "FailedOperation.InfrastructureError"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERCOMBINATION = "InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
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
func (c *Client) DescribeFieldValuesWithContext(ctx context.Context, request *DescribeFieldValuesRequest) (response *DescribeFieldValuesResponse, err error) {
    if request == nil {
        request = NewDescribeFieldValuesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "rum", APIVersion, "DescribeFieldValues")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeFieldValues require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeFieldValuesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeFileCertificateRequest() (request *DescribeFileCertificateRequest) {
    request = &DescribeFileCertificateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "DescribeFileCertificate")
    
    
    return
}

func NewDescribeFileCertificateResponse() (response *DescribeFileCertificateResponse) {
    response = &DescribeFileCertificateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeFileCertificate
// 获取COS存储的临时密钥，具体到ProjectId
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CHARGENOBALANCE = "FailedOperation.ChargeNoBalance"
//  FAILEDOPERATION_CHARGENOPAYRIGHT = "FailedOperation.ChargeNoPayRight"
//  FAILEDOPERATION_CHARGEPARAMINVALID = "FailedOperation.ChargeParamInvalid"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  FAILEDOPERATION_INFRASTRUCTUREERROR = "FailedOperation.InfrastructureError"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERCOMBINATION = "InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
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
func (c *Client) DescribeFileCertificate(request *DescribeFileCertificateRequest) (response *DescribeFileCertificateResponse, err error) {
    return c.DescribeFileCertificateWithContext(context.Background(), request)
}

// DescribeFileCertificate
// 获取COS存储的临时密钥，具体到ProjectId
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CHARGENOBALANCE = "FailedOperation.ChargeNoBalance"
//  FAILEDOPERATION_CHARGENOPAYRIGHT = "FailedOperation.ChargeNoPayRight"
//  FAILEDOPERATION_CHARGEPARAMINVALID = "FailedOperation.ChargeParamInvalid"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  FAILEDOPERATION_INFRASTRUCTUREERROR = "FailedOperation.InfrastructureError"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERCOMBINATION = "InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
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
func (c *Client) DescribeFileCertificateWithContext(ctx context.Context, request *DescribeFileCertificateRequest) (response *DescribeFileCertificateResponse, err error) {
    if request == nil {
        request = NewDescribeFileCertificateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "rum", APIVersion, "DescribeFileCertificate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeFileCertificate require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeFileCertificateResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeIssueAffectedUsersRequest() (request *DescribeIssueAffectedUsersRequest) {
    request = &DescribeIssueAffectedUsersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "DescribeIssueAffectedUsers")
    
    
    return
}

func NewDescribeIssueAffectedUsersResponse() (response *DescribeIssueAffectedUsersResponse) {
    response = &DescribeIssueAffectedUsersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeIssueAffectedUsers
// 按用户维度聚合某问题的受影响用户清单，服务端分页。AppId 与 IssueId 必填；AppVersion / StartTime / EndTime / Keyword 用于收窄范围，同时作用于清单与 Total，因此 Total 为当前条件下去重后的用户数， 条件全空时等于问题详情的影响用户数。OrderKey / OrderBy 只改变行序、不改变结果集合；排序一律由服务端执行， 并保留 uid 作为次级键以保证翻页不重不漏。接口只读。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeIssueAffectedUsers(request *DescribeIssueAffectedUsersRequest) (response *DescribeIssueAffectedUsersResponse, err error) {
    return c.DescribeIssueAffectedUsersWithContext(context.Background(), request)
}

// DescribeIssueAffectedUsers
// 按用户维度聚合某问题的受影响用户清单，服务端分页。AppId 与 IssueId 必填；AppVersion / StartTime / EndTime / Keyword 用于收窄范围，同时作用于清单与 Total，因此 Total 为当前条件下去重后的用户数， 条件全空时等于问题详情的影响用户数。OrderKey / OrderBy 只改变行序、不改变结果集合；排序一律由服务端执行， 并保留 uid 作为次级键以保证翻页不重不漏。接口只读。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeIssueAffectedUsersWithContext(ctx context.Context, request *DescribeIssueAffectedUsersRequest) (response *DescribeIssueAffectedUsersResponse, err error) {
    if request == nil {
        request = NewDescribeIssueAffectedUsersRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "rum", APIVersion, "DescribeIssueAffectedUsers")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeIssueAffectedUsers require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeIssueAffectedUsersResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeIssueBreadcrumbsRequest() (request *DescribeIssueBreadcrumbsRequest) {
    request = &DescribeIssueBreadcrumbsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "DescribeIssueBreadcrumbs")
    
    
    return
}

func NewDescribeIssueBreadcrumbsResponse() (response *DescribeIssueBreadcrumbsResponse) {
    response = &DescribeIssueBreadcrumbsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeIssueBreadcrumbs
// 获取指定 Issue 当前事件或最新事件之前的 Breadcrumbs 面包屑序列
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CHARGENOBALANCE = "FailedOperation.ChargeNoBalance"
//  FAILEDOPERATION_CHARGENOPAYRIGHT = "FailedOperation.ChargeNoPayRight"
//  FAILEDOPERATION_CHARGEPARAMINVALID = "FailedOperation.ChargeParamInvalid"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  FAILEDOPERATION_INFRASTRUCTUREERROR = "FailedOperation.InfrastructureError"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERCOMBINATION = "InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
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
func (c *Client) DescribeIssueBreadcrumbs(request *DescribeIssueBreadcrumbsRequest) (response *DescribeIssueBreadcrumbsResponse, err error) {
    return c.DescribeIssueBreadcrumbsWithContext(context.Background(), request)
}

// DescribeIssueBreadcrumbs
// 获取指定 Issue 当前事件或最新事件之前的 Breadcrumbs 面包屑序列
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CHARGENOBALANCE = "FailedOperation.ChargeNoBalance"
//  FAILEDOPERATION_CHARGENOPAYRIGHT = "FailedOperation.ChargeNoPayRight"
//  FAILEDOPERATION_CHARGEPARAMINVALID = "FailedOperation.ChargeParamInvalid"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  FAILEDOPERATION_INFRASTRUCTUREERROR = "FailedOperation.InfrastructureError"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERCOMBINATION = "InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
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
func (c *Client) DescribeIssueBreadcrumbsWithContext(ctx context.Context, request *DescribeIssueBreadcrumbsRequest) (response *DescribeIssueBreadcrumbsResponse, err error) {
    if request == nil {
        request = NewDescribeIssueBreadcrumbsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "rum", APIVersion, "DescribeIssueBreadcrumbs")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeIssueBreadcrumbs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeIssueBreadcrumbsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeIssueCrashLogDownloadURLRequest() (request *DescribeIssueCrashLogDownloadURLRequest) {
    request = &DescribeIssueCrashLogDownloadURLRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "DescribeIssueCrashLogDownloadURL")
    
    
    return
}

func NewDescribeIssueCrashLogDownloadURLResponse() (response *DescribeIssueCrashLogDownloadURLResponse) {
    response = &DescribeIssueCrashLogDownloadURLResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeIssueCrashLogDownloadURL
// 获取日志下载链接
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CHARGENOBALANCE = "FailedOperation.ChargeNoBalance"
//  FAILEDOPERATION_CHARGENOPAYRIGHT = "FailedOperation.ChargeNoPayRight"
//  FAILEDOPERATION_CHARGEPARAMINVALID = "FailedOperation.ChargeParamInvalid"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  FAILEDOPERATION_INFRASTRUCTUREERROR = "FailedOperation.InfrastructureError"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERCOMBINATION = "InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
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
func (c *Client) DescribeIssueCrashLogDownloadURL(request *DescribeIssueCrashLogDownloadURLRequest) (response *DescribeIssueCrashLogDownloadURLResponse, err error) {
    return c.DescribeIssueCrashLogDownloadURLWithContext(context.Background(), request)
}

// DescribeIssueCrashLogDownloadURL
// 获取日志下载链接
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CHARGENOBALANCE = "FailedOperation.ChargeNoBalance"
//  FAILEDOPERATION_CHARGENOPAYRIGHT = "FailedOperation.ChargeNoPayRight"
//  FAILEDOPERATION_CHARGEPARAMINVALID = "FailedOperation.ChargeParamInvalid"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  FAILEDOPERATION_INFRASTRUCTUREERROR = "FailedOperation.InfrastructureError"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERCOMBINATION = "InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
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
func (c *Client) DescribeIssueCrashLogDownloadURLWithContext(ctx context.Context, request *DescribeIssueCrashLogDownloadURLRequest) (response *DescribeIssueCrashLogDownloadURLResponse, err error) {
    if request == nil {
        request = NewDescribeIssueCrashLogDownloadURLRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "rum", APIVersion, "DescribeIssueCrashLogDownloadURL")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeIssueCrashLogDownloadURL require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeIssueCrashLogDownloadURLResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeIssueDetailRequest() (request *DescribeIssueDetailRequest) {
    request = &DescribeIssueDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "DescribeIssueDetail")
    
    
    return
}

func NewDescribeIssueDetailResponse() (response *DescribeIssueDetailResponse) {
    response = &DescribeIssueDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeIssueDetail
// 获取指定 Issue 的摘要信息以及默认选中的最新事件详情
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CHARGENOBALANCE = "FailedOperation.ChargeNoBalance"
//  FAILEDOPERATION_CHARGENOPAYRIGHT = "FailedOperation.ChargeNoPayRight"
//  FAILEDOPERATION_CHARGEPARAMINVALID = "FailedOperation.ChargeParamInvalid"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  FAILEDOPERATION_INFRASTRUCTUREERROR = "FailedOperation.InfrastructureError"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERCOMBINATION = "InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
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
func (c *Client) DescribeIssueDetail(request *DescribeIssueDetailRequest) (response *DescribeIssueDetailResponse, err error) {
    return c.DescribeIssueDetailWithContext(context.Background(), request)
}

// DescribeIssueDetail
// 获取指定 Issue 的摘要信息以及默认选中的最新事件详情
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CHARGENOBALANCE = "FailedOperation.ChargeNoBalance"
//  FAILEDOPERATION_CHARGENOPAYRIGHT = "FailedOperation.ChargeNoPayRight"
//  FAILEDOPERATION_CHARGEPARAMINVALID = "FailedOperation.ChargeParamInvalid"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  FAILEDOPERATION_INFRASTRUCTUREERROR = "FailedOperation.InfrastructureError"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERCOMBINATION = "InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
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
func (c *Client) DescribeIssueDetailWithContext(ctx context.Context, request *DescribeIssueDetailRequest) (response *DescribeIssueDetailResponse, err error) {
    if request == nil {
        request = NewDescribeIssueDetailRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "rum", APIVersion, "DescribeIssueDetail")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeIssueDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeIssueDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeIssueEventDetailRequest() (request *DescribeIssueEventDetailRequest) {
    request = &DescribeIssueEventDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "DescribeIssueEventDetail")
    
    
    return
}

func NewDescribeIssueEventDetailResponse() (response *DescribeIssueEventDetailResponse) {
    response = &DescribeIssueEventDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeIssueEventDetail
// 获取指定 Issue 下单个事件的完整详情，包括栈信息、上下文和符号化状态等字段
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CHARGENOBALANCE = "FailedOperation.ChargeNoBalance"
//  FAILEDOPERATION_CHARGENOPAYRIGHT = "FailedOperation.ChargeNoPayRight"
//  FAILEDOPERATION_CHARGEPARAMINVALID = "FailedOperation.ChargeParamInvalid"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  FAILEDOPERATION_INFRASTRUCTUREERROR = "FailedOperation.InfrastructureError"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERCOMBINATION = "InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
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
func (c *Client) DescribeIssueEventDetail(request *DescribeIssueEventDetailRequest) (response *DescribeIssueEventDetailResponse, err error) {
    return c.DescribeIssueEventDetailWithContext(context.Background(), request)
}

// DescribeIssueEventDetail
// 获取指定 Issue 下单个事件的完整详情，包括栈信息、上下文和符号化状态等字段
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CHARGENOBALANCE = "FailedOperation.ChargeNoBalance"
//  FAILEDOPERATION_CHARGENOPAYRIGHT = "FailedOperation.ChargeNoPayRight"
//  FAILEDOPERATION_CHARGEPARAMINVALID = "FailedOperation.ChargeParamInvalid"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  FAILEDOPERATION_INFRASTRUCTUREERROR = "FailedOperation.InfrastructureError"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERCOMBINATION = "InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
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
func (c *Client) DescribeIssueEventDetailWithContext(ctx context.Context, request *DescribeIssueEventDetailRequest) (response *DescribeIssueEventDetailResponse, err error) {
    if request == nil {
        request = NewDescribeIssueEventDetailRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "rum", APIVersion, "DescribeIssueEventDetail")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeIssueEventDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeIssueEventDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeIssueEventsRequest() (request *DescribeIssueEventsRequest) {
    request = &DescribeIssueEventsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "DescribeIssueEvents")
    
    
    return
}

func NewDescribeIssueEventsResponse() (response *DescribeIssueEventsResponse) {
    response = &DescribeIssueEventsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeIssueEvents
// 根据条件获取指定 Issue 的错误事件列表，支持按时间范围、关键字和扩展过滤器查询，返回分页事件数据
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CHARGENOBALANCE = "FailedOperation.ChargeNoBalance"
//  FAILEDOPERATION_CHARGENOPAYRIGHT = "FailedOperation.ChargeNoPayRight"
//  FAILEDOPERATION_CHARGEPARAMINVALID = "FailedOperation.ChargeParamInvalid"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  FAILEDOPERATION_INFRASTRUCTUREERROR = "FailedOperation.InfrastructureError"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERCOMBINATION = "InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
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
func (c *Client) DescribeIssueEvents(request *DescribeIssueEventsRequest) (response *DescribeIssueEventsResponse, err error) {
    return c.DescribeIssueEventsWithContext(context.Background(), request)
}

// DescribeIssueEvents
// 根据条件获取指定 Issue 的错误事件列表，支持按时间范围、关键字和扩展过滤器查询，返回分页事件数据
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CHARGENOBALANCE = "FailedOperation.ChargeNoBalance"
//  FAILEDOPERATION_CHARGENOPAYRIGHT = "FailedOperation.ChargeNoPayRight"
//  FAILEDOPERATION_CHARGEPARAMINVALID = "FailedOperation.ChargeParamInvalid"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  FAILEDOPERATION_INFRASTRUCTUREERROR = "FailedOperation.InfrastructureError"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERCOMBINATION = "InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
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
func (c *Client) DescribeIssueEventsWithContext(ctx context.Context, request *DescribeIssueEventsRequest) (response *DescribeIssueEventsResponse, err error) {
    if request == nil {
        request = NewDescribeIssueEventsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "rum", APIVersion, "DescribeIssueEvents")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeIssueEvents require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeIssueEventsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeIssueTrendRequest() (request *DescribeIssueTrendRequest) {
    request = &DescribeIssueTrendRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "DescribeIssueTrend")
    
    
    return
}

func NewDescribeIssueTrendResponse() (response *DescribeIssueTrendResponse) {
    response = &DescribeIssueTrendResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeIssueTrend
// 获取指定 Issue 在给定时间范围内的趋势分桶统计数据
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CHARGENOBALANCE = "FailedOperation.ChargeNoBalance"
//  FAILEDOPERATION_CHARGENOPAYRIGHT = "FailedOperation.ChargeNoPayRight"
//  FAILEDOPERATION_CHARGEPARAMINVALID = "FailedOperation.ChargeParamInvalid"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  FAILEDOPERATION_INFRASTRUCTUREERROR = "FailedOperation.InfrastructureError"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERCOMBINATION = "InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
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
func (c *Client) DescribeIssueTrend(request *DescribeIssueTrendRequest) (response *DescribeIssueTrendResponse, err error) {
    return c.DescribeIssueTrendWithContext(context.Background(), request)
}

// DescribeIssueTrend
// 获取指定 Issue 在给定时间范围内的趋势分桶统计数据
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CHARGENOBALANCE = "FailedOperation.ChargeNoBalance"
//  FAILEDOPERATION_CHARGENOPAYRIGHT = "FailedOperation.ChargeNoPayRight"
//  FAILEDOPERATION_CHARGEPARAMINVALID = "FailedOperation.ChargeParamInvalid"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  FAILEDOPERATION_INFRASTRUCTUREERROR = "FailedOperation.InfrastructureError"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERCOMBINATION = "InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
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
func (c *Client) DescribeIssueTrendWithContext(ctx context.Context, request *DescribeIssueTrendRequest) (response *DescribeIssueTrendResponse, err error) {
    if request == nil {
        request = NewDescribeIssueTrendRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "rum", APIVersion, "DescribeIssueTrend")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeIssueTrend require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeIssueTrendResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeIssuesRequest() (request *DescribeIssuesRequest) {
    request = &DescribeIssuesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "DescribeIssues")
    
    
    return
}

func NewDescribeIssuesResponse() (response *DescribeIssuesResponse) {
    response = &DescribeIssuesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeIssues
// 根据筛选条件分页查询 Issue 列表，支持环境、Issue 类型、状态、关键字、扩展过滤器和排序条件
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CHARGENOBALANCE = "FailedOperation.ChargeNoBalance"
//  FAILEDOPERATION_CHARGENOPAYRIGHT = "FailedOperation.ChargeNoPayRight"
//  FAILEDOPERATION_CHARGEPARAMINVALID = "FailedOperation.ChargeParamInvalid"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  FAILEDOPERATION_INFRASTRUCTUREERROR = "FailedOperation.InfrastructureError"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERCOMBINATION = "InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
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
func (c *Client) DescribeIssues(request *DescribeIssuesRequest) (response *DescribeIssuesResponse, err error) {
    return c.DescribeIssuesWithContext(context.Background(), request)
}

// DescribeIssues
// 根据筛选条件分页查询 Issue 列表，支持环境、Issue 类型、状态、关键字、扩展过滤器和排序条件
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CHARGENOBALANCE = "FailedOperation.ChargeNoBalance"
//  FAILEDOPERATION_CHARGENOPAYRIGHT = "FailedOperation.ChargeNoPayRight"
//  FAILEDOPERATION_CHARGEPARAMINVALID = "FailedOperation.ChargeParamInvalid"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  FAILEDOPERATION_INFRASTRUCTUREERROR = "FailedOperation.InfrastructureError"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERCOMBINATION = "InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
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
func (c *Client) DescribeIssuesWithContext(ctx context.Context, request *DescribeIssuesRequest) (response *DescribeIssuesResponse, err error) {
    if request == nil {
        request = NewDescribeIssuesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "rum", APIVersion, "DescribeIssues")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeIssues require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeIssuesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeIssuesDistributionRequest() (request *DescribeIssuesDistributionRequest) {
    request = &DescribeIssuesDistributionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "DescribeIssuesDistribution")
    
    
    return
}

func NewDescribeIssuesDistributionResponse() (response *DescribeIssuesDistributionResponse) {
    response = &DescribeIssuesDistributionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeIssuesDistribution
// 获取问题分布数据
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeIssuesDistribution(request *DescribeIssuesDistributionRequest) (response *DescribeIssuesDistributionResponse, err error) {
    return c.DescribeIssuesDistributionWithContext(context.Background(), request)
}

// DescribeIssuesDistribution
// 获取问题分布数据
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeIssuesDistributionWithContext(ctx context.Context, request *DescribeIssuesDistributionRequest) (response *DescribeIssuesDistributionResponse, err error) {
    if request == nil {
        request = NewDescribeIssuesDistributionRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "rum", APIVersion, "DescribeIssuesDistribution")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeIssuesDistribution require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeIssuesDistributionResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeIssuesListRequest() (request *DescribeIssuesListRequest) {
    request = &DescribeIssuesListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "DescribeIssuesList")
    
    
    return
}

func NewDescribeIssuesListResponse() (response *DescribeIssuesListResponse) {
    response = &DescribeIssuesListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeIssuesList
// 获取问题列表数据
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeIssuesList(request *DescribeIssuesListRequest) (response *DescribeIssuesListResponse, err error) {
    return c.DescribeIssuesListWithContext(context.Background(), request)
}

// DescribeIssuesList
// 获取问题列表数据
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeIssuesListWithContext(ctx context.Context, request *DescribeIssuesListRequest) (response *DescribeIssuesListResponse, err error) {
    if request == nil {
        request = NewDescribeIssuesListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "rum", APIVersion, "DescribeIssuesList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeIssuesList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeIssuesListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeIssuesStatisticsTrendRequest() (request *DescribeIssuesStatisticsTrendRequest) {
    request = &DescribeIssuesStatisticsTrendRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "DescribeIssuesStatisticsTrend")
    
    
    return
}

func NewDescribeIssuesStatisticsTrendResponse() (response *DescribeIssuesStatisticsTrendResponse) {
    response = &DescribeIssuesStatisticsTrendResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeIssuesStatisticsTrend
// 获取问题趋势数据
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeIssuesStatisticsTrend(request *DescribeIssuesStatisticsTrendRequest) (response *DescribeIssuesStatisticsTrendResponse, err error) {
    return c.DescribeIssuesStatisticsTrendWithContext(context.Background(), request)
}

// DescribeIssuesStatisticsTrend
// 获取问题趋势数据
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeIssuesStatisticsTrendWithContext(ctx context.Context, request *DescribeIssuesStatisticsTrendRequest) (response *DescribeIssuesStatisticsTrendResponse, err error) {
    if request == nil {
        request = NewDescribeIssuesStatisticsTrendRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "rum", APIVersion, "DescribeIssuesStatisticsTrend")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeIssuesStatisticsTrend require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeIssuesStatisticsTrendResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLagANRProblemAccountDetailRequest() (request *DescribeLagANRProblemAccountDetailRequest) {
    request = &DescribeLagANRProblemAccountDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "DescribeLagANRProblemAccountDetail")
    
    
    return
}

func NewDescribeLagANRProblemAccountDetailResponse() (response *DescribeLagANRProblemAccountDetailResponse) {
    response = &DescribeLagANRProblemAccountDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeLagANRProblemAccountDetail
// 获取卡顿ANR问题用户详情
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeLagANRProblemAccountDetail(request *DescribeLagANRProblemAccountDetailRequest) (response *DescribeLagANRProblemAccountDetailResponse, err error) {
    return c.DescribeLagANRProblemAccountDetailWithContext(context.Background(), request)
}

// DescribeLagANRProblemAccountDetail
// 获取卡顿ANR问题用户详情
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeLagANRProblemAccountDetailWithContext(ctx context.Context, request *DescribeLagANRProblemAccountDetailRequest) (response *DescribeLagANRProblemAccountDetailResponse, err error) {
    if request == nil {
        request = NewDescribeLagANRProblemAccountDetailRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "rum", APIVersion, "DescribeLagANRProblemAccountDetail")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLagANRProblemAccountDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLagANRProblemAccountDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLagANRProblemFeatureAccountsRequest() (request *DescribeLagANRProblemFeatureAccountsRequest) {
    request = &DescribeLagANRProblemFeatureAccountsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "DescribeLagANRProblemFeatureAccounts")
    
    
    return
}

func NewDescribeLagANRProblemFeatureAccountsResponse() (response *DescribeLagANRProblemFeatureAccountsResponse) {
    response = &DescribeLagANRProblemFeatureAccountsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeLagANRProblemFeatureAccounts
// 获取卡顿ANR问题用户列表
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeLagANRProblemFeatureAccounts(request *DescribeLagANRProblemFeatureAccountsRequest) (response *DescribeLagANRProblemFeatureAccountsResponse, err error) {
    return c.DescribeLagANRProblemFeatureAccountsWithContext(context.Background(), request)
}

// DescribeLagANRProblemFeatureAccounts
// 获取卡顿ANR问题用户列表
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeLagANRProblemFeatureAccountsWithContext(ctx context.Context, request *DescribeLagANRProblemFeatureAccountsRequest) (response *DescribeLagANRProblemFeatureAccountsResponse, err error) {
    if request == nil {
        request = NewDescribeLagANRProblemFeatureAccountsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "rum", APIVersion, "DescribeLagANRProblemFeatureAccounts")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLagANRProblemFeatureAccounts require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLagANRProblemFeatureAccountsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLagANRProblemListRequest() (request *DescribeLagANRProblemListRequest) {
    request = &DescribeLagANRProblemListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "DescribeLagANRProblemList")
    
    
    return
}

func NewDescribeLagANRProblemListResponse() (response *DescribeLagANRProblemListResponse) {
    response = &DescribeLagANRProblemListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeLagANRProblemList
// 获取卡顿ANR问题列表数据
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeLagANRProblemList(request *DescribeLagANRProblemListRequest) (response *DescribeLagANRProblemListResponse, err error) {
    return c.DescribeLagANRProblemListWithContext(context.Background(), request)
}

// DescribeLagANRProblemList
// 获取卡顿ANR问题列表数据
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeLagANRProblemListWithContext(ctx context.Context, request *DescribeLagANRProblemListRequest) (response *DescribeLagANRProblemListResponse, err error) {
    if request == nil {
        request = NewDescribeLagANRProblemListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "rum", APIVersion, "DescribeLagANRProblemList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLagANRProblemList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLagANRProblemListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMiniProgramReplayPackRequest() (request *DescribeMiniProgramReplayPackRequest) {
    request = &DescribeMiniProgramReplayPackRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "DescribeMiniProgramReplayPack")
    
    
    return
}

func NewDescribeMiniProgramReplayPackResponse() (response *DescribeMiniProgramReplayPackResponse) {
    response = &DescribeMiniProgramReplayPackResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeMiniProgramReplayPack
// 按AppId+BundleHash取已编译回放包正文。成功含Pack（tmpl/css/js/manifest/hash），不返回WXML/WXSS源码。不存在返回ResourceNotFound。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CHARGENOBALANCE = "FailedOperation.ChargeNoBalance"
//  FAILEDOPERATION_CHARGENOPAYRIGHT = "FailedOperation.ChargeNoPayRight"
//  FAILEDOPERATION_CHARGEPARAMINVALID = "FailedOperation.ChargeParamInvalid"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  FAILEDOPERATION_INFRASTRUCTUREERROR = "FailedOperation.InfrastructureError"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERCOMBINATION = "InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
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
func (c *Client) DescribeMiniProgramReplayPack(request *DescribeMiniProgramReplayPackRequest) (response *DescribeMiniProgramReplayPackResponse, err error) {
    return c.DescribeMiniProgramReplayPackWithContext(context.Background(), request)
}

// DescribeMiniProgramReplayPack
// 按AppId+BundleHash取已编译回放包正文。成功含Pack（tmpl/css/js/manifest/hash），不返回WXML/WXSS源码。不存在返回ResourceNotFound。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CHARGENOBALANCE = "FailedOperation.ChargeNoBalance"
//  FAILEDOPERATION_CHARGENOPAYRIGHT = "FailedOperation.ChargeNoPayRight"
//  FAILEDOPERATION_CHARGEPARAMINVALID = "FailedOperation.ChargeParamInvalid"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  FAILEDOPERATION_INFRASTRUCTUREERROR = "FailedOperation.InfrastructureError"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERCOMBINATION = "InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
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
func (c *Client) DescribeMiniProgramReplayPackWithContext(ctx context.Context, request *DescribeMiniProgramReplayPackRequest) (response *DescribeMiniProgramReplayPackResponse, err error) {
    if request == nil {
        request = NewDescribeMiniProgramReplayPackRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "rum", APIVersion, "DescribeMiniProgramReplayPack")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMiniProgramReplayPack require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMiniProgramReplayPackResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMiniProgramReplayPacksRequest() (request *DescribeMiniProgramReplayPacksRequest) {
    request = &DescribeMiniProgramReplayPacksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "DescribeMiniProgramReplayPacks")
    
    
    return
}

func NewDescribeMiniProgramReplayPacksResponse() (response *DescribeMiniProgramReplayPacksResponse) {
    response = &DescribeMiniProgramReplayPacksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeMiniProgramReplayPacks
// 按AppId列出回放包元数据，不含Pack正文。可按Status过滤。Limit默认20、最大100。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CHARGENOBALANCE = "FailedOperation.ChargeNoBalance"
//  FAILEDOPERATION_CHARGENOPAYRIGHT = "FailedOperation.ChargeNoPayRight"
//  FAILEDOPERATION_CHARGEPARAMINVALID = "FailedOperation.ChargeParamInvalid"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  FAILEDOPERATION_INFRASTRUCTUREERROR = "FailedOperation.InfrastructureError"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERCOMBINATION = "InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
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
func (c *Client) DescribeMiniProgramReplayPacks(request *DescribeMiniProgramReplayPacksRequest) (response *DescribeMiniProgramReplayPacksResponse, err error) {
    return c.DescribeMiniProgramReplayPacksWithContext(context.Background(), request)
}

// DescribeMiniProgramReplayPacks
// 按AppId列出回放包元数据，不含Pack正文。可按Status过滤。Limit默认20、最大100。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CHARGENOBALANCE = "FailedOperation.ChargeNoBalance"
//  FAILEDOPERATION_CHARGENOPAYRIGHT = "FailedOperation.ChargeNoPayRight"
//  FAILEDOPERATION_CHARGEPARAMINVALID = "FailedOperation.ChargeParamInvalid"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  FAILEDOPERATION_INFRASTRUCTUREERROR = "FailedOperation.InfrastructureError"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERCOMBINATION = "InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
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
func (c *Client) DescribeMiniProgramReplayPacksWithContext(ctx context.Context, request *DescribeMiniProgramReplayPacksRequest) (response *DescribeMiniProgramReplayPacksResponse, err error) {
    if request == nil {
        request = NewDescribeMiniProgramReplayPacksRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "rum", APIVersion, "DescribeMiniProgramReplayPacks")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMiniProgramReplayPacks require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMiniProgramReplayPacksResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePerformanceMetricsRequest() (request *DescribePerformanceMetricsRequest) {
    request = &DescribePerformanceMetricsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "DescribePerformanceMetrics")
    
    
    return
}

func NewDescribePerformanceMetricsResponse() (response *DescribePerformanceMetricsResponse) {
    response = &DescribePerformanceMetricsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribePerformanceMetrics
// 获取TDEM性能指标数据
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CHARGENOBALANCE = "FailedOperation.ChargeNoBalance"
//  FAILEDOPERATION_CHARGENOPAYRIGHT = "FailedOperation.ChargeNoPayRight"
//  FAILEDOPERATION_CHARGEPARAMINVALID = "FailedOperation.ChargeParamInvalid"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  FAILEDOPERATION_INFRASTRUCTUREERROR = "FailedOperation.InfrastructureError"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERCOMBINATION = "InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
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
func (c *Client) DescribePerformanceMetrics(request *DescribePerformanceMetricsRequest) (response *DescribePerformanceMetricsResponse, err error) {
    return c.DescribePerformanceMetricsWithContext(context.Background(), request)
}

// DescribePerformanceMetrics
// 获取TDEM性能指标数据
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CHARGENOBALANCE = "FailedOperation.ChargeNoBalance"
//  FAILEDOPERATION_CHARGENOPAYRIGHT = "FailedOperation.ChargeNoPayRight"
//  FAILEDOPERATION_CHARGEPARAMINVALID = "FailedOperation.ChargeParamInvalid"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  FAILEDOPERATION_INFRASTRUCTUREERROR = "FailedOperation.InfrastructureError"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERCOMBINATION = "InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
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
func (c *Client) DescribePerformanceMetricsWithContext(ctx context.Context, request *DescribePerformanceMetricsRequest) (response *DescribePerformanceMetricsResponse, err error) {
    if request == nil {
        request = NewDescribePerformanceMetricsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "rum", APIVersion, "DescribePerformanceMetrics")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePerformanceMetrics require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePerformanceMetricsResponse()
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

func NewDescribeReplayDataRequest() (request *DescribeReplayDataRequest) {
    request = &DescribeReplayDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "DescribeReplayData")
    
    
    return
}

func NewDescribeReplayDataResponse() (response *DescribeReplayDataResponse) {
    response = &DescribeReplayDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeReplayData
// 从对象存储读取录屏帧数据
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CHARGENOBALANCE = "FailedOperation.ChargeNoBalance"
//  FAILEDOPERATION_CHARGENOPAYRIGHT = "FailedOperation.ChargeNoPayRight"
//  FAILEDOPERATION_CHARGEPARAMINVALID = "FailedOperation.ChargeParamInvalid"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  FAILEDOPERATION_INFRASTRUCTUREERROR = "FailedOperation.InfrastructureError"
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
//  RESOURCENOTFOUND_NOINSTANCE = "ResourceNotFound.NoInstance"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeReplayData(request *DescribeReplayDataRequest) (response *DescribeReplayDataResponse, err error) {
    return c.DescribeReplayDataWithContext(context.Background(), request)
}

// DescribeReplayData
// 从对象存储读取录屏帧数据
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CHARGENOBALANCE = "FailedOperation.ChargeNoBalance"
//  FAILEDOPERATION_CHARGENOPAYRIGHT = "FailedOperation.ChargeNoPayRight"
//  FAILEDOPERATION_CHARGEPARAMINVALID = "FailedOperation.ChargeParamInvalid"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  FAILEDOPERATION_INFRASTRUCTUREERROR = "FailedOperation.InfrastructureError"
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
//  RESOURCENOTFOUND_NOINSTANCE = "ResourceNotFound.NoInstance"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeReplayDataWithContext(ctx context.Context, request *DescribeReplayDataRequest) (response *DescribeReplayDataResponse, err error) {
    if request == nil {
        request = NewDescribeReplayDataRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "rum", APIVersion, "DescribeReplayData")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeReplayData require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeReplayDataResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeReplayDetailsRequest() (request *DescribeReplayDetailsRequest) {
    request = &DescribeReplayDetailsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "DescribeReplayDetails")
    
    
    return
}

func NewDescribeReplayDetailsResponse() (response *DescribeReplayDetailsResponse) {
    response = &DescribeReplayDetailsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeReplayDetails
// 获取回放日志
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CHARGENOBALANCE = "FailedOperation.ChargeNoBalance"
//  FAILEDOPERATION_CHARGENOPAYRIGHT = "FailedOperation.ChargeNoPayRight"
//  FAILEDOPERATION_CHARGEPARAMINVALID = "FailedOperation.ChargeParamInvalid"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  FAILEDOPERATION_INFRASTRUCTUREERROR = "FailedOperation.InfrastructureError"
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
//  RESOURCENOTFOUND_NOINSTANCE = "ResourceNotFound.NoInstance"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeReplayDetails(request *DescribeReplayDetailsRequest) (response *DescribeReplayDetailsResponse, err error) {
    return c.DescribeReplayDetailsWithContext(context.Background(), request)
}

// DescribeReplayDetails
// 获取回放日志
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CHARGENOBALANCE = "FailedOperation.ChargeNoBalance"
//  FAILEDOPERATION_CHARGENOPAYRIGHT = "FailedOperation.ChargeNoPayRight"
//  FAILEDOPERATION_CHARGEPARAMINVALID = "FailedOperation.ChargeParamInvalid"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  FAILEDOPERATION_INFRASTRUCTUREERROR = "FailedOperation.InfrastructureError"
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
//  RESOURCENOTFOUND_NOINSTANCE = "ResourceNotFound.NoInstance"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeReplayDetailsWithContext(ctx context.Context, request *DescribeReplayDetailsRequest) (response *DescribeReplayDetailsResponse, err error) {
    if request == nil {
        request = NewDescribeReplayDetailsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "rum", APIVersion, "DescribeReplayDetails")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeReplayDetails require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeReplayDetailsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeReplayMetaRequest() (request *DescribeReplayMetaRequest) {
    request = &DescribeReplayMetaRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "DescribeReplayMeta")
    
    
    return
}

func NewDescribeReplayMetaResponse() (response *DescribeReplayMetaResponse) {
    response = &DescribeReplayMetaResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeReplayMeta
// 获取回放元数据
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CHARGENOBALANCE = "FailedOperation.ChargeNoBalance"
//  FAILEDOPERATION_CHARGENOPAYRIGHT = "FailedOperation.ChargeNoPayRight"
//  FAILEDOPERATION_CHARGEPARAMINVALID = "FailedOperation.ChargeParamInvalid"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  FAILEDOPERATION_INFRASTRUCTUREERROR = "FailedOperation.InfrastructureError"
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
//  RESOURCENOTFOUND_NOINSTANCE = "ResourceNotFound.NoInstance"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeReplayMeta(request *DescribeReplayMetaRequest) (response *DescribeReplayMetaResponse, err error) {
    return c.DescribeReplayMetaWithContext(context.Background(), request)
}

// DescribeReplayMeta
// 获取回放元数据
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CHARGENOBALANCE = "FailedOperation.ChargeNoBalance"
//  FAILEDOPERATION_CHARGENOPAYRIGHT = "FailedOperation.ChargeNoPayRight"
//  FAILEDOPERATION_CHARGEPARAMINVALID = "FailedOperation.ChargeParamInvalid"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  FAILEDOPERATION_INFRASTRUCTUREERROR = "FailedOperation.InfrastructureError"
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
//  RESOURCENOTFOUND_NOINSTANCE = "ResourceNotFound.NoInstance"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeReplayMetaWithContext(ctx context.Context, request *DescribeReplayMetaRequest) (response *DescribeReplayMetaResponse, err error) {
    if request == nil {
        request = NewDescribeReplayMetaRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "rum", APIVersion, "DescribeReplayMeta")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeReplayMeta require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeReplayMetaResponse()
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

func NewDescribeRumGroupLogV2Request() (request *DescribeRumGroupLogV2Request) {
    request = &DescribeRumGroupLogV2Request{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "DescribeRumGroupLogV2")
    
    
    return
}

func NewDescribeRumGroupLogV2Response() (response *DescribeRumGroupLogV2Response) {
    response = &DescribeRumGroupLogV2Response{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRumGroupLogV2
// 根据label 聚合分析日志
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CHARGENOBALANCE = "FailedOperation.ChargeNoBalance"
//  FAILEDOPERATION_CHARGENOPAYRIGHT = "FailedOperation.ChargeNoPayRight"
//  FAILEDOPERATION_CHARGEPARAMINVALID = "FailedOperation.ChargeParamInvalid"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  FAILEDOPERATION_INFRASTRUCTUREERROR = "FailedOperation.InfrastructureError"
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
//  RESOURCENOTFOUND_NOINSTANCE = "ResourceNotFound.NoInstance"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeRumGroupLogV2(request *DescribeRumGroupLogV2Request) (response *DescribeRumGroupLogV2Response, err error) {
    return c.DescribeRumGroupLogV2WithContext(context.Background(), request)
}

// DescribeRumGroupLogV2
// 根据label 聚合分析日志
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CHARGENOBALANCE = "FailedOperation.ChargeNoBalance"
//  FAILEDOPERATION_CHARGENOPAYRIGHT = "FailedOperation.ChargeNoPayRight"
//  FAILEDOPERATION_CHARGEPARAMINVALID = "FailedOperation.ChargeParamInvalid"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  FAILEDOPERATION_INFRASTRUCTUREERROR = "FailedOperation.InfrastructureError"
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
//  RESOURCENOTFOUND_NOINSTANCE = "ResourceNotFound.NoInstance"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeRumGroupLogV2WithContext(ctx context.Context, request *DescribeRumGroupLogV2Request) (response *DescribeRumGroupLogV2Response, err error) {
    if request == nil {
        request = NewDescribeRumGroupLogV2Request()
    }
    c.InitBaseRequest(&request.BaseRequest, "rum", APIVersion, "DescribeRumGroupLogV2")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRumGroupLogV2 require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRumGroupLogV2Response()
    err = c.Send(request, response)
    return
}

func NewDescribeRumLogDetailsV2Request() (request *DescribeRumLogDetailsV2Request) {
    request = &DescribeRumLogDetailsV2Request{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "DescribeRumLogDetailsV2")
    
    
    return
}

func NewDescribeRumLogDetailsV2Response() (response *DescribeRumLogDetailsV2Response) {
    response = &DescribeRumLogDetailsV2Response{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRumLogDetailsV2
// 查询日志明细
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CHARGENOBALANCE = "FailedOperation.ChargeNoBalance"
//  FAILEDOPERATION_CHARGENOPAYRIGHT = "FailedOperation.ChargeNoPayRight"
//  FAILEDOPERATION_CHARGEPARAMINVALID = "FailedOperation.ChargeParamInvalid"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  FAILEDOPERATION_INFRASTRUCTUREERROR = "FailedOperation.InfrastructureError"
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
//  RESOURCENOTFOUND_NOINSTANCE = "ResourceNotFound.NoInstance"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeRumLogDetailsV2(request *DescribeRumLogDetailsV2Request) (response *DescribeRumLogDetailsV2Response, err error) {
    return c.DescribeRumLogDetailsV2WithContext(context.Background(), request)
}

// DescribeRumLogDetailsV2
// 查询日志明细
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CHARGENOBALANCE = "FailedOperation.ChargeNoBalance"
//  FAILEDOPERATION_CHARGENOPAYRIGHT = "FailedOperation.ChargeNoPayRight"
//  FAILEDOPERATION_CHARGEPARAMINVALID = "FailedOperation.ChargeParamInvalid"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  FAILEDOPERATION_INFRASTRUCTUREERROR = "FailedOperation.InfrastructureError"
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
//  RESOURCENOTFOUND_NOINSTANCE = "ResourceNotFound.NoInstance"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeRumLogDetailsV2WithContext(ctx context.Context, request *DescribeRumLogDetailsV2Request) (response *DescribeRumLogDetailsV2Response, err error) {
    if request == nil {
        request = NewDescribeRumLogDetailsV2Request()
    }
    c.InitBaseRequest(&request.BaseRequest, "rum", APIVersion, "DescribeRumLogDetailsV2")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRumLogDetailsV2 require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRumLogDetailsV2Response()
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

func NewDescribeRumLogExportV2Request() (request *DescribeRumLogExportV2Request) {
    request = &DescribeRumLogExportV2Request{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "DescribeRumLogExportV2")
    
    
    return
}

func NewDescribeRumLogExportV2Response() (response *DescribeRumLogExportV2Response) {
    response = &DescribeRumLogExportV2Response{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRumLogExportV2
// 创建日志导出
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CHARGENOBALANCE = "FailedOperation.ChargeNoBalance"
//  FAILEDOPERATION_CHARGENOPAYRIGHT = "FailedOperation.ChargeNoPayRight"
//  FAILEDOPERATION_CHARGEPARAMINVALID = "FailedOperation.ChargeParamInvalid"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  FAILEDOPERATION_INFRASTRUCTUREERROR = "FailedOperation.InfrastructureError"
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
//  RESOURCENOTFOUND_NOINSTANCE = "ResourceNotFound.NoInstance"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeRumLogExportV2(request *DescribeRumLogExportV2Request) (response *DescribeRumLogExportV2Response, err error) {
    return c.DescribeRumLogExportV2WithContext(context.Background(), request)
}

// DescribeRumLogExportV2
// 创建日志导出
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CHARGENOBALANCE = "FailedOperation.ChargeNoBalance"
//  FAILEDOPERATION_CHARGENOPAYRIGHT = "FailedOperation.ChargeNoPayRight"
//  FAILEDOPERATION_CHARGEPARAMINVALID = "FailedOperation.ChargeParamInvalid"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  FAILEDOPERATION_INFRASTRUCTUREERROR = "FailedOperation.InfrastructureError"
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
//  RESOURCENOTFOUND_NOINSTANCE = "ResourceNotFound.NoInstance"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeRumLogExportV2WithContext(ctx context.Context, request *DescribeRumLogExportV2Request) (response *DescribeRumLogExportV2Response, err error) {
    if request == nil {
        request = NewDescribeRumLogExportV2Request()
    }
    c.InitBaseRequest(&request.BaseRequest, "rum", APIVersion, "DescribeRumLogExportV2")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRumLogExportV2 require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRumLogExportV2Response()
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

func NewDescribeRumLogExportsV2Request() (request *DescribeRumLogExportsV2Request) {
    request = &DescribeRumLogExportsV2Request{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "DescribeRumLogExportsV2")
    
    
    return
}

func NewDescribeRumLogExportsV2Response() (response *DescribeRumLogExportsV2Response) {
    response = &DescribeRumLogExportsV2Response{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRumLogExportsV2
// 获取日志导出列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CHARGENOBALANCE = "FailedOperation.ChargeNoBalance"
//  FAILEDOPERATION_CHARGENOPAYRIGHT = "FailedOperation.ChargeNoPayRight"
//  FAILEDOPERATION_CHARGEPARAMINVALID = "FailedOperation.ChargeParamInvalid"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  FAILEDOPERATION_INFRASTRUCTUREERROR = "FailedOperation.InfrastructureError"
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
//  RESOURCENOTFOUND_NOINSTANCE = "ResourceNotFound.NoInstance"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeRumLogExportsV2(request *DescribeRumLogExportsV2Request) (response *DescribeRumLogExportsV2Response, err error) {
    return c.DescribeRumLogExportsV2WithContext(context.Background(), request)
}

// DescribeRumLogExportsV2
// 获取日志导出列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CHARGENOBALANCE = "FailedOperation.ChargeNoBalance"
//  FAILEDOPERATION_CHARGENOPAYRIGHT = "FailedOperation.ChargeNoPayRight"
//  FAILEDOPERATION_CHARGEPARAMINVALID = "FailedOperation.ChargeParamInvalid"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  FAILEDOPERATION_INFRASTRUCTUREERROR = "FailedOperation.InfrastructureError"
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
//  RESOURCENOTFOUND_NOINSTANCE = "ResourceNotFound.NoInstance"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeRumLogExportsV2WithContext(ctx context.Context, request *DescribeRumLogExportsV2Request) (response *DescribeRumLogExportsV2Response, err error) {
    if request == nil {
        request = NewDescribeRumLogExportsV2Request()
    }
    c.InitBaseRequest(&request.BaseRequest, "rum", APIVersion, "DescribeRumLogExportsV2")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRumLogExportsV2 require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRumLogExportsV2Response()
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

func NewDescribeRumLogTotalV2Request() (request *DescribeRumLogTotalV2Request) {
    request = &DescribeRumLogTotalV2Request{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "DescribeRumLogTotalV2")
    
    
    return
}

func NewDescribeRumLogTotalV2Response() (response *DescribeRumLogTotalV2Response) {
    response = &DescribeRumLogTotalV2Response{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRumLogTotalV2
// 查询原始日志总量
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CHARGENOBALANCE = "FailedOperation.ChargeNoBalance"
//  FAILEDOPERATION_CHARGENOPAYRIGHT = "FailedOperation.ChargeNoPayRight"
//  FAILEDOPERATION_CHARGEPARAMINVALID = "FailedOperation.ChargeParamInvalid"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  FAILEDOPERATION_INFRASTRUCTUREERROR = "FailedOperation.InfrastructureError"
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
//  RESOURCENOTFOUND_NOINSTANCE = "ResourceNotFound.NoInstance"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeRumLogTotalV2(request *DescribeRumLogTotalV2Request) (response *DescribeRumLogTotalV2Response, err error) {
    return c.DescribeRumLogTotalV2WithContext(context.Background(), request)
}

// DescribeRumLogTotalV2
// 查询原始日志总量
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CHARGENOBALANCE = "FailedOperation.ChargeNoBalance"
//  FAILEDOPERATION_CHARGENOPAYRIGHT = "FailedOperation.ChargeNoPayRight"
//  FAILEDOPERATION_CHARGEPARAMINVALID = "FailedOperation.ChargeParamInvalid"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  FAILEDOPERATION_INFRASTRUCTUREERROR = "FailedOperation.InfrastructureError"
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
//  RESOURCENOTFOUND_NOINSTANCE = "ResourceNotFound.NoInstance"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeRumLogTotalV2WithContext(ctx context.Context, request *DescribeRumLogTotalV2Request) (response *DescribeRumLogTotalV2Response, err error) {
    if request == nil {
        request = NewDescribeRumLogTotalV2Request()
    }
    c.InitBaseRequest(&request.BaseRequest, "rum", APIVersion, "DescribeRumLogTotalV2")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRumLogTotalV2 require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRumLogTotalV2Response()
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

func NewDescribeRumStatsLogListV2Request() (request *DescribeRumStatsLogListV2Request) {
    request = &DescribeRumStatsLogListV2Request{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "DescribeRumStatsLogListV2")
    
    
    return
}

func NewDescribeRumStatsLogListV2Response() (response *DescribeRumStatsLogListV2Response) {
    response = &DescribeRumStatsLogListV2Response{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRumStatsLogListV2
// 查询日志时间分布
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CHARGENOBALANCE = "FailedOperation.ChargeNoBalance"
//  FAILEDOPERATION_CHARGENOPAYRIGHT = "FailedOperation.ChargeNoPayRight"
//  FAILEDOPERATION_CHARGEPARAMINVALID = "FailedOperation.ChargeParamInvalid"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  FAILEDOPERATION_INFRASTRUCTUREERROR = "FailedOperation.InfrastructureError"
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
//  RESOURCENOTFOUND_NOINSTANCE = "ResourceNotFound.NoInstance"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeRumStatsLogListV2(request *DescribeRumStatsLogListV2Request) (response *DescribeRumStatsLogListV2Response, err error) {
    return c.DescribeRumStatsLogListV2WithContext(context.Background(), request)
}

// DescribeRumStatsLogListV2
// 查询日志时间分布
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CHARGENOBALANCE = "FailedOperation.ChargeNoBalance"
//  FAILEDOPERATION_CHARGENOPAYRIGHT = "FailedOperation.ChargeNoPayRight"
//  FAILEDOPERATION_CHARGEPARAMINVALID = "FailedOperation.ChargeParamInvalid"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  FAILEDOPERATION_INFRASTRUCTUREERROR = "FailedOperation.InfrastructureError"
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
//  RESOURCENOTFOUND_NOINSTANCE = "ResourceNotFound.NoInstance"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeRumStatsLogListV2WithContext(ctx context.Context, request *DescribeRumStatsLogListV2Request) (response *DescribeRumStatsLogListV2Response, err error) {
    if request == nil {
        request = NewDescribeRumStatsLogListV2Request()
    }
    c.InitBaseRequest(&request.BaseRequest, "rum", APIVersion, "DescribeRumStatsLogListV2")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRumStatsLogListV2 require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRumStatsLogListV2Response()
    err = c.Send(request, response)
    return
}

func NewDescribeSampleRateRequest() (request *DescribeSampleRateRequest) {
    request = &DescribeSampleRateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "DescribeSampleRate")
    
    
    return
}

func NewDescribeSampleRateResponse() (response *DescribeSampleRateResponse) {
    response = &DescribeSampleRateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSampleRate
// 获取应用的采样率
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CHARGENOBALANCE = "FailedOperation.ChargeNoBalance"
//  FAILEDOPERATION_CHARGENOPAYRIGHT = "FailedOperation.ChargeNoPayRight"
//  FAILEDOPERATION_CHARGEPARAMINVALID = "FailedOperation.ChargeParamInvalid"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  FAILEDOPERATION_INFRASTRUCTUREERROR = "FailedOperation.InfrastructureError"
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
//  RESOURCENOTFOUND_NOINSTANCE = "ResourceNotFound.NoInstance"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeSampleRate(request *DescribeSampleRateRequest) (response *DescribeSampleRateResponse, err error) {
    return c.DescribeSampleRateWithContext(context.Background(), request)
}

// DescribeSampleRate
// 获取应用的采样率
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CHARGENOBALANCE = "FailedOperation.ChargeNoBalance"
//  FAILEDOPERATION_CHARGENOPAYRIGHT = "FailedOperation.ChargeNoPayRight"
//  FAILEDOPERATION_CHARGEPARAMINVALID = "FailedOperation.ChargeParamInvalid"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  FAILEDOPERATION_INFRASTRUCTUREERROR = "FailedOperation.InfrastructureError"
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
//  RESOURCENOTFOUND_NOINSTANCE = "ResourceNotFound.NoInstance"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeSampleRateWithContext(ctx context.Context, request *DescribeSampleRateRequest) (response *DescribeSampleRateResponse, err error) {
    if request == nil {
        request = NewDescribeSampleRateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "rum", APIVersion, "DescribeSampleRate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSampleRate require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSampleRateResponse()
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

func NewDescribeSessionListRequest() (request *DescribeSessionListRequest) {
    request = &DescribeSessionListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "DescribeSessionList")
    
    
    return
}

func NewDescribeSessionListResponse() (response *DescribeSessionListResponse) {
    response = &DescribeSessionListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSessionList
// 查询会话列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CHARGENOBALANCE = "FailedOperation.ChargeNoBalance"
//  FAILEDOPERATION_CHARGENOPAYRIGHT = "FailedOperation.ChargeNoPayRight"
//  FAILEDOPERATION_CHARGEPARAMINVALID = "FailedOperation.ChargeParamInvalid"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  FAILEDOPERATION_INFRASTRUCTUREERROR = "FailedOperation.InfrastructureError"
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
//  RESOURCENOTFOUND_NOINSTANCE = "ResourceNotFound.NoInstance"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeSessionList(request *DescribeSessionListRequest) (response *DescribeSessionListResponse, err error) {
    return c.DescribeSessionListWithContext(context.Background(), request)
}

// DescribeSessionList
// 查询会话列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CHARGENOBALANCE = "FailedOperation.ChargeNoBalance"
//  FAILEDOPERATION_CHARGENOPAYRIGHT = "FailedOperation.ChargeNoPayRight"
//  FAILEDOPERATION_CHARGEPARAMINVALID = "FailedOperation.ChargeParamInvalid"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  FAILEDOPERATION_INFRASTRUCTUREERROR = "FailedOperation.InfrastructureError"
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
//  RESOURCENOTFOUND_NOINSTANCE = "ResourceNotFound.NoInstance"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeSessionListWithContext(ctx context.Context, request *DescribeSessionListRequest) (response *DescribeSessionListResponse, err error) {
    if request == nil {
        request = NewDescribeSessionListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "rum", APIVersion, "DescribeSessionList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSessionList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSessionListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeStruggleDistributionRequest() (request *DescribeStruggleDistributionRequest) {
    request = &DescribeStruggleDistributionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "DescribeStruggleDistribution")
    
    
    return
}

func NewDescribeStruggleDistributionResponse() (response *DescribeStruggleDistributionResponse) {
    response = &DescribeStruggleDistributionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeStruggleDistribution
// 获取挣扎事件分布数据
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CHARGENOBALANCE = "FailedOperation.ChargeNoBalance"
//  FAILEDOPERATION_CHARGENOPAYRIGHT = "FailedOperation.ChargeNoPayRight"
//  FAILEDOPERATION_CHARGEPARAMINVALID = "FailedOperation.ChargeParamInvalid"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  FAILEDOPERATION_INFRASTRUCTUREERROR = "FailedOperation.InfrastructureError"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERCOMBINATION = "InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
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
func (c *Client) DescribeStruggleDistribution(request *DescribeStruggleDistributionRequest) (response *DescribeStruggleDistributionResponse, err error) {
    return c.DescribeStruggleDistributionWithContext(context.Background(), request)
}

// DescribeStruggleDistribution
// 获取挣扎事件分布数据
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CHARGENOBALANCE = "FailedOperation.ChargeNoBalance"
//  FAILEDOPERATION_CHARGENOPAYRIGHT = "FailedOperation.ChargeNoPayRight"
//  FAILEDOPERATION_CHARGEPARAMINVALID = "FailedOperation.ChargeParamInvalid"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  FAILEDOPERATION_INFRASTRUCTUREERROR = "FailedOperation.InfrastructureError"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERCOMBINATION = "InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
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
func (c *Client) DescribeStruggleDistributionWithContext(ctx context.Context, request *DescribeStruggleDistributionRequest) (response *DescribeStruggleDistributionResponse, err error) {
    if request == nil {
        request = NewDescribeStruggleDistributionRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "rum", APIVersion, "DescribeStruggleDistribution")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeStruggleDistribution require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeStruggleDistributionResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSymbolArtifactUploadCredentialRequest() (request *DescribeSymbolArtifactUploadCredentialRequest) {
    request = &DescribeSymbolArtifactUploadCredentialRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "DescribeSymbolArtifactUploadCredential")
    
    
    return
}

func NewDescribeSymbolArtifactUploadCredentialResponse() (response *DescribeSymbolArtifactUploadCredentialResponse) {
    response = &DescribeSymbolArtifactUploadCredentialResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSymbolArtifactUploadCredential
// 获取符号制品上传临时凭证。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CHARGENOBALANCE = "FailedOperation.ChargeNoBalance"
//  FAILEDOPERATION_CHARGENOPAYRIGHT = "FailedOperation.ChargeNoPayRight"
//  FAILEDOPERATION_CHARGEPARAMINVALID = "FailedOperation.ChargeParamInvalid"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  FAILEDOPERATION_INFRASTRUCTUREERROR = "FailedOperation.InfrastructureError"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERCOMBINATION = "InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
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
func (c *Client) DescribeSymbolArtifactUploadCredential(request *DescribeSymbolArtifactUploadCredentialRequest) (response *DescribeSymbolArtifactUploadCredentialResponse, err error) {
    return c.DescribeSymbolArtifactUploadCredentialWithContext(context.Background(), request)
}

// DescribeSymbolArtifactUploadCredential
// 获取符号制品上传临时凭证。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CHARGENOBALANCE = "FailedOperation.ChargeNoBalance"
//  FAILEDOPERATION_CHARGENOPAYRIGHT = "FailedOperation.ChargeNoPayRight"
//  FAILEDOPERATION_CHARGEPARAMINVALID = "FailedOperation.ChargeParamInvalid"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  FAILEDOPERATION_INFRASTRUCTUREERROR = "FailedOperation.InfrastructureError"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERCOMBINATION = "InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
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
func (c *Client) DescribeSymbolArtifactUploadCredentialWithContext(ctx context.Context, request *DescribeSymbolArtifactUploadCredentialRequest) (response *DescribeSymbolArtifactUploadCredentialResponse, err error) {
    if request == nil {
        request = NewDescribeSymbolArtifactUploadCredentialRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "rum", APIVersion, "DescribeSymbolArtifactUploadCredential")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSymbolArtifactUploadCredential require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSymbolArtifactUploadCredentialResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSymbolArtifactsRequest() (request *DescribeSymbolArtifactsRequest) {
    request = &DescribeSymbolArtifactsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "DescribeSymbolArtifacts")
    
    
    return
}

func NewDescribeSymbolArtifactsResponse() (response *DescribeSymbolArtifactsResponse) {
    response = &DescribeSymbolArtifactsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSymbolArtifacts
// 分页查询指定类型的符号制品。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CHARGENOBALANCE = "FailedOperation.ChargeNoBalance"
//  FAILEDOPERATION_CHARGENOPAYRIGHT = "FailedOperation.ChargeNoPayRight"
//  FAILEDOPERATION_CHARGEPARAMINVALID = "FailedOperation.ChargeParamInvalid"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  FAILEDOPERATION_INFRASTRUCTUREERROR = "FailedOperation.InfrastructureError"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERCOMBINATION = "InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
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
func (c *Client) DescribeSymbolArtifacts(request *DescribeSymbolArtifactsRequest) (response *DescribeSymbolArtifactsResponse, err error) {
    return c.DescribeSymbolArtifactsWithContext(context.Background(), request)
}

// DescribeSymbolArtifacts
// 分页查询指定类型的符号制品。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CHARGENOBALANCE = "FailedOperation.ChargeNoBalance"
//  FAILEDOPERATION_CHARGENOPAYRIGHT = "FailedOperation.ChargeNoPayRight"
//  FAILEDOPERATION_CHARGEPARAMINVALID = "FailedOperation.ChargeParamInvalid"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  FAILEDOPERATION_INFRASTRUCTUREERROR = "FailedOperation.InfrastructureError"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERCOMBINATION = "InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
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
func (c *Client) DescribeSymbolArtifactsWithContext(ctx context.Context, request *DescribeSymbolArtifactsRequest) (response *DescribeSymbolArtifactsResponse, err error) {
    if request == nil {
        request = NewDescribeSymbolArtifactsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "rum", APIVersion, "DescribeSymbolArtifacts")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSymbolArtifacts require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSymbolArtifactsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTDEMApplicationsRequest() (request *DescribeTDEMApplicationsRequest) {
    request = &DescribeTDEMApplicationsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "DescribeTDEMApplications")
    
    
    return
}

func NewDescribeTDEMApplicationsResponse() (response *DescribeTDEMApplicationsResponse) {
    response = &DescribeTDEMApplicationsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTDEMApplications
// 查询应用列表，默认仅当前用户名下有权限的应用；可按业务系统 InstanceId 筛选。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERCOMBINATION = "InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeTDEMApplications(request *DescribeTDEMApplicationsRequest) (response *DescribeTDEMApplicationsResponse, err error) {
    return c.DescribeTDEMApplicationsWithContext(context.Background(), request)
}

// DescribeTDEMApplications
// 查询应用列表，默认仅当前用户名下有权限的应用；可按业务系统 InstanceId 筛选。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERCOMBINATION = "InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeTDEMApplicationsWithContext(ctx context.Context, request *DescribeTDEMApplicationsRequest) (response *DescribeTDEMApplicationsResponse, err error) {
    if request == nil {
        request = NewDescribeTDEMApplicationsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "rum", APIVersion, "DescribeTDEMApplications")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTDEMApplications require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTDEMApplicationsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTDEMBusinessSystemsRequest() (request *DescribeTDEMBusinessSystemsRequest) {
    request = &DescribeTDEMBusinessSystemsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "DescribeTDEMBusinessSystems")
    
    
    return
}

func NewDescribeTDEMBusinessSystemsResponse() (response *DescribeTDEMBusinessSystemsResponse) {
    response = &DescribeTDEMBusinessSystemsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTDEMBusinessSystems
// 查询当前用户名下业务系统列表（超管返回全部）。分页 Offset/Limit，默认 Limit=20 上限 100。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERCOMBINATION = "InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeTDEMBusinessSystems(request *DescribeTDEMBusinessSystemsRequest) (response *DescribeTDEMBusinessSystemsResponse, err error) {
    return c.DescribeTDEMBusinessSystemsWithContext(context.Background(), request)
}

// DescribeTDEMBusinessSystems
// 查询当前用户名下业务系统列表（超管返回全部）。分页 Offset/Limit，默认 Limit=20 上限 100。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERCOMBINATION = "InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeTDEMBusinessSystemsWithContext(ctx context.Context, request *DescribeTDEMBusinessSystemsRequest) (response *DescribeTDEMBusinessSystemsResponse, err error) {
    if request == nil {
        request = NewDescribeTDEMBusinessSystemsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "rum", APIVersion, "DescribeTDEMBusinessSystems")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTDEMBusinessSystems require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTDEMBusinessSystemsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTDEMDialDimensionsRequest() (request *DescribeTDEMDialDimensionsRequest) {
    request = &DescribeTDEMDialDimensionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "DescribeTDEMDialDimensions")
    
    
    return
}

func NewDescribeTDEMDialDimensionsResponse() (response *DescribeTDEMDialDimensionsResponse) {
    response = &DescribeTDEMDialDimensionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTDEMDialDimensions
// 按一个维度拆分拨测指标。Dimension 取值说明：platform、app_version、device、net_type、operator、region、country、target_ip。target_region 不是合法值。region 对应客户端省份，country 对应国家。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeTDEMDialDimensions(request *DescribeTDEMDialDimensionsRequest) (response *DescribeTDEMDialDimensionsResponse, err error) {
    return c.DescribeTDEMDialDimensionsWithContext(context.Background(), request)
}

// DescribeTDEMDialDimensions
// 按一个维度拆分拨测指标。Dimension 取值说明：platform、app_version、device、net_type、operator、region、country、target_ip。target_region 不是合法值。region 对应客户端省份，country 对应国家。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeTDEMDialDimensionsWithContext(ctx context.Context, request *DescribeTDEMDialDimensionsRequest) (response *DescribeTDEMDialDimensionsResponse, err error) {
    if request == nil {
        request = NewDescribeTDEMDialDimensionsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "rum", APIVersion, "DescribeTDEMDialDimensions")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTDEMDialDimensions require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTDEMDialDimensionsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTDEMDialExecutionDetailRequest() (request *DescribeTDEMDialExecutionDetailRequest) {
    request = &DescribeTDEMDialExecutionDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "DescribeTDEMDialExecutionDetail")
    
    
    return
}

func NewDescribeTDEMDialExecutionDetailResponse() (response *DescribeTDEMDialExecutionDetailResponse) {
    response = &DescribeTDEMDialExecutionDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTDEMDialExecutionDetail
// 返回一个拨测批次的探测结果和路由追踪。Probes 的 Region、Operator 是目标 IP 归属。逐跳归属同时写在 Probes.MTR.Hops 和顶层 MTR.Hops。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeTDEMDialExecutionDetail(request *DescribeTDEMDialExecutionDetailRequest) (response *DescribeTDEMDialExecutionDetailResponse, err error) {
    return c.DescribeTDEMDialExecutionDetailWithContext(context.Background(), request)
}

// DescribeTDEMDialExecutionDetail
// 返回一个拨测批次的探测结果和路由追踪。Probes 的 Region、Operator 是目标 IP 归属。逐跳归属同时写在 Probes.MTR.Hops 和顶层 MTR.Hops。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeTDEMDialExecutionDetailWithContext(ctx context.Context, request *DescribeTDEMDialExecutionDetailRequest) (response *DescribeTDEMDialExecutionDetailResponse, err error) {
    if request == nil {
        request = NewDescribeTDEMDialExecutionDetailRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "rum", APIVersion, "DescribeTDEMDialExecutionDetail")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTDEMDialExecutionDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTDEMDialExecutionDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTDEMDialExecutionListRequest() (request *DescribeTDEMDialExecutionListRequest) {
    request = &DescribeTDEMDialExecutionListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "DescribeTDEMDialExecutionList")
    
    
    return
}

func NewDescribeTDEMDialExecutionListResponse() (response *DescribeTDEMDialExecutionListResponse) {
    response = &DescribeTDEMDialExecutionListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTDEMDialExecutionList
// 按批次返回拨测执行记录，一批次一行。TargetIps 是目标地址，TargetIpGeo 是这些地址的地区和运营商。行上的 Region 和 Operator 仍是客户端地区和运营商。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeTDEMDialExecutionList(request *DescribeTDEMDialExecutionListRequest) (response *DescribeTDEMDialExecutionListResponse, err error) {
    return c.DescribeTDEMDialExecutionListWithContext(context.Background(), request)
}

// DescribeTDEMDialExecutionList
// 按批次返回拨测执行记录，一批次一行。TargetIps 是目标地址，TargetIpGeo 是这些地址的地区和运营商。行上的 Region 和 Operator 仍是客户端地区和运营商。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeTDEMDialExecutionListWithContext(ctx context.Context, request *DescribeTDEMDialExecutionListRequest) (response *DescribeTDEMDialExecutionListResponse, err error) {
    if request == nil {
        request = NewDescribeTDEMDialExecutionListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "rum", APIVersion, "DescribeTDEMDialExecutionList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTDEMDialExecutionList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTDEMDialExecutionListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTDEMDialOverviewRequest() (request *DescribeTDEMDialOverviewRequest) {
    request = &DescribeTDEMDialOverviewRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "DescribeTDEMDialOverview")
    
    
    return
}

func NewDescribeTDEMDialOverviewResponse() (response *DescribeTDEMDialOverviewResponse) {
    response = &DescribeTDEMDialOverviewResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTDEMDialOverview
// 按时间窗汇总拨测批次。一次多 IP 探测计为 1 个批次。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeTDEMDialOverview(request *DescribeTDEMDialOverviewRequest) (response *DescribeTDEMDialOverviewResponse, err error) {
    return c.DescribeTDEMDialOverviewWithContext(context.Background(), request)
}

// DescribeTDEMDialOverview
// 按时间窗汇总拨测批次。一次多 IP 探测计为 1 个批次。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeTDEMDialOverviewWithContext(ctx context.Context, request *DescribeTDEMDialOverviewRequest) (response *DescribeTDEMDialOverviewResponse, err error) {
    if request == nil {
        request = NewDescribeTDEMDialOverviewRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "rum", APIVersion, "DescribeTDEMDialOverview")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTDEMDialOverview require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTDEMDialOverviewResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTDEMDialTasksRequest() (request *DescribeTDEMDialTasksRequest) {
    request = &DescribeTDEMDialTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "DescribeTDEMDialTasks")
    
    
    return
}

func NewDescribeTDEMDialTasksResponse() (response *DescribeTDEMDialTasksResponse) {
    response = &DescribeTDEMDialTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTDEMDialTasks
// 查询应用下未删除的拨测任务，包含启动范围。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeTDEMDialTasks(request *DescribeTDEMDialTasksRequest) (response *DescribeTDEMDialTasksResponse, err error) {
    return c.DescribeTDEMDialTasksWithContext(context.Background(), request)
}

// DescribeTDEMDialTasks
// 查询应用下未删除的拨测任务，包含启动范围。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeTDEMDialTasksWithContext(ctx context.Context, request *DescribeTDEMDialTasksRequest) (response *DescribeTDEMDialTasksResponse, err error) {
    if request == nil {
        request = NewDescribeTDEMDialTasksRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "rum", APIVersion, "DescribeTDEMDialTasks")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTDEMDialTasks require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTDEMDialTasksResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTDEMDialTrendRequest() (request *DescribeTDEMDialTrendRequest) {
    request = &DescribeTDEMDialTrendRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "DescribeTDEMDialTrend")
    
    
    return
}

func NewDescribeTDEMDialTrendResponse() (response *DescribeTDEMDialTrendResponse) {
    response = &DescribeTDEMDialTrendResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTDEMDialTrend
// 按小时或天返回拨测趋势。Granularity 不传时按小时。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeTDEMDialTrend(request *DescribeTDEMDialTrendRequest) (response *DescribeTDEMDialTrendResponse, err error) {
    return c.DescribeTDEMDialTrendWithContext(context.Background(), request)
}

// DescribeTDEMDialTrend
// 按小时或天返回拨测趋势。Granularity 不传时按小时。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeTDEMDialTrendWithContext(ctx context.Context, request *DescribeTDEMDialTrendRequest) (response *DescribeTDEMDialTrendResponse, err error) {
    if request == nil {
        request = NewDescribeTDEMDialTrendRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "rum", APIVersion, "DescribeTDEMDialTrend")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTDEMDialTrend require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTDEMDialTrendResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTDEMEventDetailRequest() (request *DescribeTDEMEventDetailRequest) {
    request = &DescribeTDEMEventDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "DescribeTDEMEventDetail")
    
    
    return
}

func NewDescribeTDEMEventDetailResponse() (response *DescribeTDEMEventDetailResponse) {
    response = &DescribeTDEMEventDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTDEMEventDetail
// 按 AppId+EventId 查询事件详情，含主表字段与可选符号化结果。Date 推荐传以命中分区；或用 StartDate+EndDate（<=2 天）。越权与不存在统一 ResourceNotFound。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERCOMBINATION = "InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeTDEMEventDetail(request *DescribeTDEMEventDetailRequest) (response *DescribeTDEMEventDetailResponse, err error) {
    return c.DescribeTDEMEventDetailWithContext(context.Background(), request)
}

// DescribeTDEMEventDetail
// 按 AppId+EventId 查询事件详情，含主表字段与可选符号化结果。Date 推荐传以命中分区；或用 StartDate+EndDate（<=2 天）。越权与不存在统一 ResourceNotFound。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERCOMBINATION = "InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeTDEMEventDetailWithContext(ctx context.Context, request *DescribeTDEMEventDetailRequest) (response *DescribeTDEMEventDetailResponse, err error) {
    if request == nil {
        request = NewDescribeTDEMEventDetailRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "rum", APIVersion, "DescribeTDEMEventDetail")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTDEMEventDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTDEMEventDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTDEMEventRankRequest() (request *DescribeTDEMEventRankRequest) {
    request = &DescribeTDEMEventRankRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "DescribeTDEMEventRank")
    
    
    return
}

func NewDescribeTDEMEventRankResponse() (response *DescribeTDEMEventRankResponse) {
    response = &DescribeTDEMEventRankResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTDEMEventRank
// 错误指纹或 API/资源 URL（含按访问页、状态、标签）排行。读 tdem_events，不走主分析立方 GroupBy。Kind：error/api/resource/api_page/resource_page/api_status/api_status_top/resource_status/tag。 Kind=exception_dimension返回当前普通维度或标签值的同批完整分类计数；展开不另行取数。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERCOMBINATION = "InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeTDEMEventRank(request *DescribeTDEMEventRankRequest) (response *DescribeTDEMEventRankResponse, err error) {
    return c.DescribeTDEMEventRankWithContext(context.Background(), request)
}

// DescribeTDEMEventRank
// 错误指纹或 API/资源 URL（含按访问页、状态、标签）排行。读 tdem_events，不走主分析立方 GroupBy。Kind：error/api/resource/api_page/resource_page/api_status/api_status_top/resource_status/tag。 Kind=exception_dimension返回当前普通维度或标签值的同批完整分类计数；展开不另行取数。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERCOMBINATION = "InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeTDEMEventRankWithContext(ctx context.Context, request *DescribeTDEMEventRankRequest) (response *DescribeTDEMEventRankResponse, err error) {
    if request == nil {
        request = NewDescribeTDEMEventRankRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "rum", APIVersion, "DescribeTDEMEventRank")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTDEMEventRank require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTDEMEventRankResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTDEMEventSearchRequest() (request *DescribeTDEMEventSearchRequest) {
    request = &DescribeTDEMEventSearchRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "DescribeTDEMEventSearch")
    
    
    return
}

func NewDescribeTDEMEventSearchResponse() (response *DescribeTDEMEventSearchResponse) {
    response = &DescribeTDEMEventSearchResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTDEMEventSearch
// 数据探索事件明细分页。会话分类不得调用本接口（Category 禁止 session）。时间单位 Unix 毫秒。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERCOMBINATION = "InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeTDEMEventSearch(request *DescribeTDEMEventSearchRequest) (response *DescribeTDEMEventSearchResponse, err error) {
    return c.DescribeTDEMEventSearchWithContext(context.Background(), request)
}

// DescribeTDEMEventSearch
// 数据探索事件明细分页。会话分类不得调用本接口（Category 禁止 session）。时间单位 Unix 毫秒。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERCOMBINATION = "InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeTDEMEventSearchWithContext(ctx context.Context, request *DescribeTDEMEventSearchRequest) (response *DescribeTDEMEventSearchResponse, err error) {
    if request == nil {
        request = NewDescribeTDEMEventSearchRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "rum", APIVersion, "DescribeTDEMEventSearch")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTDEMEventSearch require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTDEMEventSearchResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTDEMEventSearchSummaryRequest() (request *DescribeTDEMEventSearchSummaryRequest) {
    request = &DescribeTDEMEventSearchSummaryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "DescribeTDEMEventSearchSummary")
    
    
    return
}

func NewDescribeTDEMEventSearchSummaryResponse() (response *DescribeTDEMEventSearchSummaryResponse) {
    response = &DescribeTDEMEventSearchSummaryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTDEMEventSearchSummary
// 同一筛选下的分类计数、访问用户、业务用户、窗口合计与分桶趋势。Category 可含 session。空窗返回 0 与空 Series。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERCOMBINATION = "InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeTDEMEventSearchSummary(request *DescribeTDEMEventSearchSummaryRequest) (response *DescribeTDEMEventSearchSummaryResponse, err error) {
    return c.DescribeTDEMEventSearchSummaryWithContext(context.Background(), request)
}

// DescribeTDEMEventSearchSummary
// 同一筛选下的分类计数、访问用户、业务用户、窗口合计与分桶趋势。Category 可含 session。空窗返回 0 与空 Series。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERCOMBINATION = "InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeTDEMEventSearchSummaryWithContext(ctx context.Context, request *DescribeTDEMEventSearchSummaryRequest) (response *DescribeTDEMEventSearchSummaryResponse, err error) {
    if request == nil {
        request = NewDescribeTDEMEventSearchSummaryRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "rum", APIVersion, "DescribeTDEMEventSearchSummary")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTDEMEventSearchSummary require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTDEMEventSearchSummaryResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTDEMFieldTransformRulesRequest() (request *DescribeTDEMFieldTransformRulesRequest) {
    request = &DescribeTDEMFieldTransformRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "DescribeTDEMFieldTransformRules")
    
    
    return
}

func NewDescribeTDEMFieldTransformRulesResponse() (response *DescribeTDEMFieldTransformRulesResponse) {
    response = &DescribeTDEMFieldTransformRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTDEMFieldTransformRules
// 查询应用下字段值替换规则列表。替代已废弃的 DescribeFieldTransformRules；AppId 为 string。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERCOMBINATION = "InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeTDEMFieldTransformRules(request *DescribeTDEMFieldTransformRulesRequest) (response *DescribeTDEMFieldTransformRulesResponse, err error) {
    return c.DescribeTDEMFieldTransformRulesWithContext(context.Background(), request)
}

// DescribeTDEMFieldTransformRules
// 查询应用下字段值替换规则列表。替代已废弃的 DescribeFieldTransformRules；AppId 为 string。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERCOMBINATION = "InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeTDEMFieldTransformRulesWithContext(ctx context.Context, request *DescribeTDEMFieldTransformRulesRequest) (response *DescribeTDEMFieldTransformRulesResponse, err error) {
    if request == nil {
        request = NewDescribeTDEMFieldTransformRulesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "rum", APIVersion, "DescribeTDEMFieldTransformRules")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTDEMFieldTransformRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTDEMFieldTransformRulesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTDEMFunnelDetailRequest() (request *DescribeTDEMFunnelDetailRequest) {
    request = &DescribeTDEMFunnelDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "DescribeTDEMFunnelDetail")
    
    
    return
}

func NewDescribeTDEMFunnelDetailResponse() (response *DescribeTDEMFunnelDetailResponse) {
    response = &DescribeTDEMFunnelDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTDEMFunnelDetail
// 获取漏斗详情
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CHARGENOBALANCE = "FailedOperation.ChargeNoBalance"
//  FAILEDOPERATION_CHARGENOPAYRIGHT = "FailedOperation.ChargeNoPayRight"
//  FAILEDOPERATION_CHARGEPARAMINVALID = "FailedOperation.ChargeParamInvalid"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  FAILEDOPERATION_INFRASTRUCTUREERROR = "FailedOperation.InfrastructureError"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERCOMBINATION = "InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
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
func (c *Client) DescribeTDEMFunnelDetail(request *DescribeTDEMFunnelDetailRequest) (response *DescribeTDEMFunnelDetailResponse, err error) {
    return c.DescribeTDEMFunnelDetailWithContext(context.Background(), request)
}

// DescribeTDEMFunnelDetail
// 获取漏斗详情
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CHARGENOBALANCE = "FailedOperation.ChargeNoBalance"
//  FAILEDOPERATION_CHARGENOPAYRIGHT = "FailedOperation.ChargeNoPayRight"
//  FAILEDOPERATION_CHARGEPARAMINVALID = "FailedOperation.ChargeParamInvalid"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  FAILEDOPERATION_INFRASTRUCTUREERROR = "FailedOperation.InfrastructureError"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERCOMBINATION = "InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
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
func (c *Client) DescribeTDEMFunnelDetailWithContext(ctx context.Context, request *DescribeTDEMFunnelDetailRequest) (response *DescribeTDEMFunnelDetailResponse, err error) {
    if request == nil {
        request = NewDescribeTDEMFunnelDetailRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "rum", APIVersion, "DescribeTDEMFunnelDetail")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTDEMFunnelDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTDEMFunnelDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTDEMFunnelDimensionsRequest() (request *DescribeTDEMFunnelDimensionsRequest) {
    request = &DescribeTDEMFunnelDimensionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "DescribeTDEMFunnelDimensions")
    
    
    return
}

func NewDescribeTDEMFunnelDimensionsResponse() (response *DescribeTDEMFunnelDimensionsResponse) {
    response = &DescribeTDEMFunnelDimensionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTDEMFunnelDimensions
// 获取漏斗分析报告维度数据
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CHARGENOBALANCE = "FailedOperation.ChargeNoBalance"
//  FAILEDOPERATION_CHARGENOPAYRIGHT = "FailedOperation.ChargeNoPayRight"
//  FAILEDOPERATION_CHARGEPARAMINVALID = "FailedOperation.ChargeParamInvalid"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  FAILEDOPERATION_INFRASTRUCTUREERROR = "FailedOperation.InfrastructureError"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERCOMBINATION = "InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
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
func (c *Client) DescribeTDEMFunnelDimensions(request *DescribeTDEMFunnelDimensionsRequest) (response *DescribeTDEMFunnelDimensionsResponse, err error) {
    return c.DescribeTDEMFunnelDimensionsWithContext(context.Background(), request)
}

// DescribeTDEMFunnelDimensions
// 获取漏斗分析报告维度数据
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CHARGENOBALANCE = "FailedOperation.ChargeNoBalance"
//  FAILEDOPERATION_CHARGENOPAYRIGHT = "FailedOperation.ChargeNoPayRight"
//  FAILEDOPERATION_CHARGEPARAMINVALID = "FailedOperation.ChargeParamInvalid"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  FAILEDOPERATION_INFRASTRUCTUREERROR = "FailedOperation.InfrastructureError"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERCOMBINATION = "InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
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
func (c *Client) DescribeTDEMFunnelDimensionsWithContext(ctx context.Context, request *DescribeTDEMFunnelDimensionsRequest) (response *DescribeTDEMFunnelDimensionsResponse, err error) {
    if request == nil {
        request = NewDescribeTDEMFunnelDimensionsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "rum", APIVersion, "DescribeTDEMFunnelDimensions")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTDEMFunnelDimensions require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTDEMFunnelDimensionsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTDEMFunnelListRequest() (request *DescribeTDEMFunnelListRequest) {
    request = &DescribeTDEMFunnelListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "DescribeTDEMFunnelList")
    
    
    return
}

func NewDescribeTDEMFunnelListResponse() (response *DescribeTDEMFunnelListResponse) {
    response = &DescribeTDEMFunnelListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTDEMFunnelList
// 获取漏斗列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CHARGENOBALANCE = "FailedOperation.ChargeNoBalance"
//  FAILEDOPERATION_CHARGENOPAYRIGHT = "FailedOperation.ChargeNoPayRight"
//  FAILEDOPERATION_CHARGEPARAMINVALID = "FailedOperation.ChargeParamInvalid"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  FAILEDOPERATION_INFRASTRUCTUREERROR = "FailedOperation.InfrastructureError"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERCOMBINATION = "InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
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
func (c *Client) DescribeTDEMFunnelList(request *DescribeTDEMFunnelListRequest) (response *DescribeTDEMFunnelListResponse, err error) {
    return c.DescribeTDEMFunnelListWithContext(context.Background(), request)
}

// DescribeTDEMFunnelList
// 获取漏斗列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CHARGENOBALANCE = "FailedOperation.ChargeNoBalance"
//  FAILEDOPERATION_CHARGENOPAYRIGHT = "FailedOperation.ChargeNoPayRight"
//  FAILEDOPERATION_CHARGEPARAMINVALID = "FailedOperation.ChargeParamInvalid"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  FAILEDOPERATION_INFRASTRUCTUREERROR = "FailedOperation.InfrastructureError"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERCOMBINATION = "InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
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
func (c *Client) DescribeTDEMFunnelListWithContext(ctx context.Context, request *DescribeTDEMFunnelListRequest) (response *DescribeTDEMFunnelListResponse, err error) {
    if request == nil {
        request = NewDescribeTDEMFunnelListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "rum", APIVersion, "DescribeTDEMFunnelList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTDEMFunnelList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTDEMFunnelListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTDEMFunnelReportRequest() (request *DescribeTDEMFunnelReportRequest) {
    request = &DescribeTDEMFunnelReportRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "DescribeTDEMFunnelReport")
    
    
    return
}

func NewDescribeTDEMFunnelReportResponse() (response *DescribeTDEMFunnelReportResponse) {
    response = &DescribeTDEMFunnelReportResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTDEMFunnelReport
// 获取漏斗报告
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CHARGENOBALANCE = "FailedOperation.ChargeNoBalance"
//  FAILEDOPERATION_CHARGENOPAYRIGHT = "FailedOperation.ChargeNoPayRight"
//  FAILEDOPERATION_CHARGEPARAMINVALID = "FailedOperation.ChargeParamInvalid"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  FAILEDOPERATION_INFRASTRUCTUREERROR = "FailedOperation.InfrastructureError"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERCOMBINATION = "InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
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
func (c *Client) DescribeTDEMFunnelReport(request *DescribeTDEMFunnelReportRequest) (response *DescribeTDEMFunnelReportResponse, err error) {
    return c.DescribeTDEMFunnelReportWithContext(context.Background(), request)
}

// DescribeTDEMFunnelReport
// 获取漏斗报告
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CHARGENOBALANCE = "FailedOperation.ChargeNoBalance"
//  FAILEDOPERATION_CHARGENOPAYRIGHT = "FailedOperation.ChargeNoPayRight"
//  FAILEDOPERATION_CHARGEPARAMINVALID = "FailedOperation.ChargeParamInvalid"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  FAILEDOPERATION_INFRASTRUCTUREERROR = "FailedOperation.InfrastructureError"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERCOMBINATION = "InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
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
func (c *Client) DescribeTDEMFunnelReportWithContext(ctx context.Context, request *DescribeTDEMFunnelReportRequest) (response *DescribeTDEMFunnelReportResponse, err error) {
    if request == nil {
        request = NewDescribeTDEMFunnelReportRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "rum", APIVersion, "DescribeTDEMFunnelReport")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTDEMFunnelReport require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTDEMFunnelReportResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTDEMFunnelStepSessionListRequest() (request *DescribeTDEMFunnelStepSessionListRequest) {
    request = &DescribeTDEMFunnelStepSessionListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "DescribeTDEMFunnelStepSessionList")
    
    
    return
}

func NewDescribeTDEMFunnelStepSessionListResponse() (response *DescribeTDEMFunnelStepSessionListResponse) {
    response = &DescribeTDEMFunnelStepSessionListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTDEMFunnelStepSessionList
// 获取漏斗步骤会话
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CHARGENOBALANCE = "FailedOperation.ChargeNoBalance"
//  FAILEDOPERATION_CHARGENOPAYRIGHT = "FailedOperation.ChargeNoPayRight"
//  FAILEDOPERATION_CHARGEPARAMINVALID = "FailedOperation.ChargeParamInvalid"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  FAILEDOPERATION_INFRASTRUCTUREERROR = "FailedOperation.InfrastructureError"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERCOMBINATION = "InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
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
func (c *Client) DescribeTDEMFunnelStepSessionList(request *DescribeTDEMFunnelStepSessionListRequest) (response *DescribeTDEMFunnelStepSessionListResponse, err error) {
    return c.DescribeTDEMFunnelStepSessionListWithContext(context.Background(), request)
}

// DescribeTDEMFunnelStepSessionList
// 获取漏斗步骤会话
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CHARGENOBALANCE = "FailedOperation.ChargeNoBalance"
//  FAILEDOPERATION_CHARGENOPAYRIGHT = "FailedOperation.ChargeNoPayRight"
//  FAILEDOPERATION_CHARGEPARAMINVALID = "FailedOperation.ChargeParamInvalid"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  FAILEDOPERATION_INFRASTRUCTUREERROR = "FailedOperation.InfrastructureError"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERCOMBINATION = "InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
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
func (c *Client) DescribeTDEMFunnelStepSessionListWithContext(ctx context.Context, request *DescribeTDEMFunnelStepSessionListRequest) (response *DescribeTDEMFunnelStepSessionListResponse, err error) {
    if request == nil {
        request = NewDescribeTDEMFunnelStepSessionListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "rum", APIVersion, "DescribeTDEMFunnelStepSessionList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTDEMFunnelStepSessionList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTDEMFunnelStepSessionListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTDEMFunnelTrendRequest() (request *DescribeTDEMFunnelTrendRequest) {
    request = &DescribeTDEMFunnelTrendRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "DescribeTDEMFunnelTrend")
    
    
    return
}

func NewDescribeTDEMFunnelTrendResponse() (response *DescribeTDEMFunnelTrendResponse) {
    response = &DescribeTDEMFunnelTrendResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTDEMFunnelTrend
// 获取漏斗趋势
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CHARGENOBALANCE = "FailedOperation.ChargeNoBalance"
//  FAILEDOPERATION_CHARGENOPAYRIGHT = "FailedOperation.ChargeNoPayRight"
//  FAILEDOPERATION_CHARGEPARAMINVALID = "FailedOperation.ChargeParamInvalid"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  FAILEDOPERATION_INFRASTRUCTUREERROR = "FailedOperation.InfrastructureError"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERCOMBINATION = "InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
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
func (c *Client) DescribeTDEMFunnelTrend(request *DescribeTDEMFunnelTrendRequest) (response *DescribeTDEMFunnelTrendResponse, err error) {
    return c.DescribeTDEMFunnelTrendWithContext(context.Background(), request)
}

// DescribeTDEMFunnelTrend
// 获取漏斗趋势
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CHARGENOBALANCE = "FailedOperation.ChargeNoBalance"
//  FAILEDOPERATION_CHARGENOPAYRIGHT = "FailedOperation.ChargeNoPayRight"
//  FAILEDOPERATION_CHARGEPARAMINVALID = "FailedOperation.ChargeParamInvalid"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  FAILEDOPERATION_INFRASTRUCTUREERROR = "FailedOperation.InfrastructureError"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERCOMBINATION = "InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
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
func (c *Client) DescribeTDEMFunnelTrendWithContext(ctx context.Context, request *DescribeTDEMFunnelTrendRequest) (response *DescribeTDEMFunnelTrendResponse, err error) {
    if request == nil {
        request = NewDescribeTDEMFunnelTrendRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "rum", APIVersion, "DescribeTDEMFunnelTrend")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTDEMFunnelTrend require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTDEMFunnelTrendResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTDEMMetricDistributionRequest() (request *DescribeTDEMMetricDistributionRequest) {
    request = &DescribeTDEMMetricDistributionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "DescribeTDEMMetricDistribution")
    
    
    return
}

func NewDescribeTDEMMetricDistributionResponse() (response *DescribeTDEMMetricDistributionResponse) {
    response = &DescribeTDEMMetricDistributionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTDEMMetricDistribution
// 按单维度 GroupBy 查询分析指标分布。读 tdem_metric_agg。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERCOMBINATION = "InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeTDEMMetricDistribution(request *DescribeTDEMMetricDistributionRequest) (response *DescribeTDEMMetricDistributionResponse, err error) {
    return c.DescribeTDEMMetricDistributionWithContext(context.Background(), request)
}

// DescribeTDEMMetricDistribution
// 按单维度 GroupBy 查询分析指标分布。读 tdem_metric_agg。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERCOMBINATION = "InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeTDEMMetricDistributionWithContext(ctx context.Context, request *DescribeTDEMMetricDistributionRequest) (response *DescribeTDEMMetricDistributionResponse, err error) {
    if request == nil {
        request = NewDescribeTDEMMetricDistributionRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "rum", APIVersion, "DescribeTDEMMetricDistribution")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTDEMMetricDistribution require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTDEMMetricDistributionResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTDEMMetricSummaryRequest() (request *DescribeTDEMMetricSummaryRequest) {
    request = &DescribeTDEMMetricSummaryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "DescribeTDEMMetricSummary")
    
    
    return
}

func NewDescribeTDEMMetricSummaryResponse() (response *DescribeTDEMMetricSummaryResponse) {
    response = &DescribeTDEMMetricSummaryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTDEMMetricSummary
// 查询指定应用与时间窗内的分析指标区间合计。单查返回 Metrics；传 Queries 时返回 Results。读 tdem_metric_agg。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERCOMBINATION = "InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeTDEMMetricSummary(request *DescribeTDEMMetricSummaryRequest) (response *DescribeTDEMMetricSummaryResponse, err error) {
    return c.DescribeTDEMMetricSummaryWithContext(context.Background(), request)
}

// DescribeTDEMMetricSummary
// 查询指定应用与时间窗内的分析指标区间合计。单查返回 Metrics；传 Queries 时返回 Results。读 tdem_metric_agg。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERCOMBINATION = "InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeTDEMMetricSummaryWithContext(ctx context.Context, request *DescribeTDEMMetricSummaryRequest) (response *DescribeTDEMMetricSummaryResponse, err error) {
    if request == nil {
        request = NewDescribeTDEMMetricSummaryRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "rum", APIVersion, "DescribeTDEMMetricSummary")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTDEMMetricSummary require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTDEMMetricSummaryResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTDEMMetricTrendRequest() (request *DescribeTDEMMetricTrendRequest) {
    request = &DescribeTDEMMetricTrendRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "DescribeTDEMMetricTrend")
    
    
    return
}

func NewDescribeTDEMMetricTrendResponse() (response *DescribeTDEMMetricTrendResponse) {
    response = &DescribeTDEMMetricTrendResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTDEMMetricTrend
// 查询分析指标时序。可传 CompareStartTime/CompareEndTime 做对比窗（仅批量）。读 tdem_metric_agg。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERCOMBINATION = "InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeTDEMMetricTrend(request *DescribeTDEMMetricTrendRequest) (response *DescribeTDEMMetricTrendResponse, err error) {
    return c.DescribeTDEMMetricTrendWithContext(context.Background(), request)
}

// DescribeTDEMMetricTrend
// 查询分析指标时序。可传 CompareStartTime/CompareEndTime 做对比窗（仅批量）。读 tdem_metric_agg。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERCOMBINATION = "InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeTDEMMetricTrendWithContext(ctx context.Context, request *DescribeTDEMMetricTrendRequest) (response *DescribeTDEMMetricTrendResponse, err error) {
    if request == nil {
        request = NewDescribeTDEMMetricTrendRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "rum", APIVersion, "DescribeTDEMMetricTrend")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTDEMMetricTrend require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTDEMMetricTrendResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTDEMOverviewScoresRequest() (request *DescribeTDEMOverviewScoresRequest) {
    request = &DescribeTDEMOverviewScoresRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "DescribeTDEMOverviewScores")
    
    
    return
}

func NewDescribeTDEMOverviewScoresResponse() (response *DescribeTDEMOverviewScoresResponse) {
    response = &DescribeTDEMOverviewScoresResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTDEMOverviewScores
// 客户端总览列表批量评分。按应用列表一次返回各计分项单项分与加权总分。AppIds 1-200；不归属当前用户的应用不出现在响应中。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERCOMBINATION = "InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeTDEMOverviewScores(request *DescribeTDEMOverviewScoresRequest) (response *DescribeTDEMOverviewScoresResponse, err error) {
    return c.DescribeTDEMOverviewScoresWithContext(context.Background(), request)
}

// DescribeTDEMOverviewScores
// 客户端总览列表批量评分。按应用列表一次返回各计分项单项分与加权总分。AppIds 1-200；不归属当前用户的应用不出现在响应中。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERCOMBINATION = "InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeTDEMOverviewScoresWithContext(ctx context.Context, request *DescribeTDEMOverviewScoresRequest) (response *DescribeTDEMOverviewScoresResponse, err error) {
    if request == nil {
        request = NewDescribeTDEMOverviewScoresRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "rum", APIVersion, "DescribeTDEMOverviewScores")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTDEMOverviewScores require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTDEMOverviewScoresResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTDEMRelatedEventsRequest() (request *DescribeTDEMRelatedEventsRequest) {
    request = &DescribeTDEMRelatedEventsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "DescribeTDEMRelatedEvents")
    
    
    return
}

func NewDescribeTDEMRelatedEventsResponse() (response *DescribeTDEMRelatedEventsResponse) {
    response = &DescribeTDEMRelatedEventsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTDEMRelatedEvents
// 会话时间线 / 页面子事件 / 错误前轨迹。Mode：timeline / page / breadcrumb。Limit 默认 200，上限 200。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERCOMBINATION = "InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeTDEMRelatedEvents(request *DescribeTDEMRelatedEventsRequest) (response *DescribeTDEMRelatedEventsResponse, err error) {
    return c.DescribeTDEMRelatedEventsWithContext(context.Background(), request)
}

// DescribeTDEMRelatedEvents
// 会话时间线 / 页面子事件 / 错误前轨迹。Mode：timeline / page / breadcrumb。Limit 默认 200，上限 200。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERCOMBINATION = "InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeTDEMRelatedEventsWithContext(ctx context.Context, request *DescribeTDEMRelatedEventsRequest) (response *DescribeTDEMRelatedEventsResponse, err error) {
    if request == nil {
        request = NewDescribeTDEMRelatedEventsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "rum", APIVersion, "DescribeTDEMRelatedEvents")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTDEMRelatedEvents require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTDEMRelatedEventsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTDEMSSEDimensionsRequest() (request *DescribeTDEMSSEDimensionsRequest) {
    request = &DescribeTDEMSSEDimensionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "DescribeTDEMSSEDimensions")
    
    
    return
}

func NewDescribeTDEMSSEDimensionsResponse() (response *DescribeTDEMSSEDimensionsResponse) {
    response = &DescribeTDEMSSEDimensionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTDEMSSEDimensions
// 按一个维度返回大模型对话次数和当前指标值，按指标值降序，最多 200 行。region 由 country 与 province 拼接。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERCOMBINATION = "InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeTDEMSSEDimensions(request *DescribeTDEMSSEDimensionsRequest) (response *DescribeTDEMSSEDimensionsResponse, err error) {
    return c.DescribeTDEMSSEDimensionsWithContext(context.Background(), request)
}

// DescribeTDEMSSEDimensions
// 按一个维度返回大模型对话次数和当前指标值，按指标值降序，最多 200 行。region 由 country 与 province 拼接。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERCOMBINATION = "InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeTDEMSSEDimensionsWithContext(ctx context.Context, request *DescribeTDEMSSEDimensionsRequest) (response *DescribeTDEMSSEDimensionsResponse, err error) {
    if request == nil {
        request = NewDescribeTDEMSSEDimensionsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "rum", APIVersion, "DescribeTDEMSSEDimensions")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTDEMSSEDimensions require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTDEMSSEDimensionsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTDEMSSEDistributionRequest() (request *DescribeTDEMSSEDistributionRequest) {
    request = &DescribeTDEMSSEDistributionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "DescribeTDEMSSEDistribution")
    
    
    return
}

func NewDescribeTDEMSSEDistributionResponse() (response *DescribeTDEMSSEDistributionResponse) {
    response = &DescribeTDEMSSEDistributionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTDEMSSEDistribution
// 返回吞吐率、首次消息时延或流畅度的分档分布。档位左闭右开，最后一档有限区间右端闭合，超出右端的样本进入溢出档。错误率与卡死率返回 InvalidParameter。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERCOMBINATION = "InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeTDEMSSEDistribution(request *DescribeTDEMSSEDistributionRequest) (response *DescribeTDEMSSEDistributionResponse, err error) {
    return c.DescribeTDEMSSEDistributionWithContext(context.Background(), request)
}

// DescribeTDEMSSEDistribution
// 返回吞吐率、首次消息时延或流畅度的分档分布。档位左闭右开，最后一档有限区间右端闭合，超出右端的样本进入溢出档。错误率与卡死率返回 InvalidParameter。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERCOMBINATION = "InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeTDEMSSEDistributionWithContext(ctx context.Context, request *DescribeTDEMSSEDistributionRequest) (response *DescribeTDEMSSEDistributionResponse, err error) {
    if request == nil {
        request = NewDescribeTDEMSSEDistributionRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "rum", APIVersion, "DescribeTDEMSSEDistribution")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTDEMSSEDistribution require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTDEMSSEDistributionResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTDEMSSEPercentileRequest() (request *DescribeTDEMSSEPercentileRequest) {
    request = &DescribeTDEMSSEPercentileRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "DescribeTDEMSSEPercentile")
    
    
    return
}

func NewDescribeTDEMSSEPercentileResponse() (response *DescribeTDEMSSEPercentileResponse) {
    response = &DescribeTDEMSSEPercentileResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTDEMSSEPercentile
// 返回吞吐率、首次消息时延或流畅度在整个时间窗及各时间桶上的平均值和 P50、P75、P95、P99。错误率与卡死率返回 InvalidParameter。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERCOMBINATION = "InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeTDEMSSEPercentile(request *DescribeTDEMSSEPercentileRequest) (response *DescribeTDEMSSEPercentileResponse, err error) {
    return c.DescribeTDEMSSEPercentileWithContext(context.Background(), request)
}

// DescribeTDEMSSEPercentile
// 返回吞吐率、首次消息时延或流畅度在整个时间窗及各时间桶上的平均值和 P50、P75、P95、P99。错误率与卡死率返回 InvalidParameter。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERCOMBINATION = "InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeTDEMSSEPercentileWithContext(ctx context.Context, request *DescribeTDEMSSEPercentileRequest) (response *DescribeTDEMSSEPercentileResponse, err error) {
    if request == nil {
        request = NewDescribeTDEMSSEPercentileRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "rum", APIVersion, "DescribeTDEMSSEPercentile")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTDEMSSEPercentile require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTDEMSSEPercentileResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTDEMSSESummaryRequest() (request *DescribeTDEMSSESummaryRequest) {
    request = &DescribeTDEMSSESummaryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "DescribeTDEMSSESummary")
    
    
    return
}

func NewDescribeTDEMSSESummaryResponse() (response *DescribeTDEMSSESummaryResponse) {
    response = &DescribeTDEMSSESummaryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTDEMSSESummary
// 查询指定应用与时间窗内的大模型对话汇总，包含对话次数、token 总数、吞吐率、首次消息时延、错误率、卡死率和流畅度。只读 tdem.tdem_events 中 event_type 为 sse_call 的事件。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERCOMBINATION = "InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeTDEMSSESummary(request *DescribeTDEMSSESummaryRequest) (response *DescribeTDEMSSESummaryResponse, err error) {
    return c.DescribeTDEMSSESummaryWithContext(context.Background(), request)
}

// DescribeTDEMSSESummary
// 查询指定应用与时间窗内的大模型对话汇总，包含对话次数、token 总数、吞吐率、首次消息时延、错误率、卡死率和流畅度。只读 tdem.tdem_events 中 event_type 为 sse_call 的事件。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERCOMBINATION = "InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeTDEMSSESummaryWithContext(ctx context.Context, request *DescribeTDEMSSESummaryRequest) (response *DescribeTDEMSSESummaryResponse, err error) {
    if request == nil {
        request = NewDescribeTDEMSSESummaryRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "rum", APIVersion, "DescribeTDEMSSESummary")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTDEMSSESummary require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTDEMSSESummaryResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTDEMSSETrendRequest() (request *DescribeTDEMSSETrendRequest) {
    request = &DescribeTDEMSSETrendRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "DescribeTDEMSSETrend")
    
    
    return
}

func NewDescribeTDEMSSETrendResponse() (response *DescribeTDEMSSETrendResponse) {
    response = &DescribeTDEMSSETrendResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTDEMSSETrend
// 按时间桶返回大模型对话趋势。overall 填充 Points，其他维度按整个窗口的指标值从高到低取 TopN 并填充 Series。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERCOMBINATION = "InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeTDEMSSETrend(request *DescribeTDEMSSETrendRequest) (response *DescribeTDEMSSETrendResponse, err error) {
    return c.DescribeTDEMSSETrendWithContext(context.Background(), request)
}

// DescribeTDEMSSETrend
// 按时间桶返回大模型对话趋势。overall 填充 Points，其他维度按整个窗口的指标值从高到低取 TopN 并填充 Series。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERCOMBINATION = "InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeTDEMSSETrendWithContext(ctx context.Context, request *DescribeTDEMSSETrendRequest) (response *DescribeTDEMSSETrendResponse, err error) {
    if request == nil {
        request = NewDescribeTDEMSSETrendRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "rum", APIVersion, "DescribeTDEMSSETrend")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTDEMSSETrend require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTDEMSSETrendResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTDEMSampleConfigRequest() (request *DescribeTDEMSampleConfigRequest) {
    request = &DescribeTDEMSampleConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "DescribeTDEMSampleConfig")
    
    
    return
}

func NewDescribeTDEMSampleConfigResponse() (response *DescribeTDEMSampleConfigResponse) {
    response = &DescribeTDEMSampleConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTDEMSampleConfig
// 读取当前生效的 Web 规则化采样文档。ProjectId 必填且 > 0。无行则返回内置缺省（Event.Enabled=true、Types 空、Replay 关闭且体验率 0），不自动插库。SnapshotId 无快照为 0。query handler 做 PascalCase ↔ 存储 snake_case 映射。非 Web 应用不得走本接口，继续 DescribeSampleRate。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERCOMBINATION = "InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeTDEMSampleConfig(request *DescribeTDEMSampleConfigRequest) (response *DescribeTDEMSampleConfigResponse, err error) {
    return c.DescribeTDEMSampleConfigWithContext(context.Background(), request)
}

// DescribeTDEMSampleConfig
// 读取当前生效的 Web 规则化采样文档。ProjectId 必填且 > 0。无行则返回内置缺省（Event.Enabled=true、Types 空、Replay 关闭且体验率 0），不自动插库。SnapshotId 无快照为 0。query handler 做 PascalCase ↔ 存储 snake_case 映射。非 Web 应用不得走本接口，继续 DescribeSampleRate。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERCOMBINATION = "InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeTDEMSampleConfigWithContext(ctx context.Context, request *DescribeTDEMSampleConfigRequest) (response *DescribeTDEMSampleConfigResponse, err error) {
    if request == nil {
        request = NewDescribeTDEMSampleConfigRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "rum", APIVersion, "DescribeTDEMSampleConfig")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTDEMSampleConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTDEMSampleConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTDEMSampleRateRequest() (request *DescribeTDEMSampleRateRequest) {
    request = &DescribeTDEMSampleRateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "DescribeTDEMSampleRate")
    
    
    return
}

func NewDescribeTDEMSampleRateResponse() (response *DescribeTDEMSampleRateResponse) {
    response = &DescribeTDEMSampleRateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTDEMSampleRate
// 查询单个应用的采集采样率与会话回放采样率。替代已废弃的 DescribeSampleRate；AppId 为 string。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERCOMBINATION = "InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeTDEMSampleRate(request *DescribeTDEMSampleRateRequest) (response *DescribeTDEMSampleRateResponse, err error) {
    return c.DescribeTDEMSampleRateWithContext(context.Background(), request)
}

// DescribeTDEMSampleRate
// 查询单个应用的采集采样率与会话回放采样率。替代已废弃的 DescribeSampleRate；AppId 为 string。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERCOMBINATION = "InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeTDEMSampleRateWithContext(ctx context.Context, request *DescribeTDEMSampleRateRequest) (response *DescribeTDEMSampleRateResponse, err error) {
    if request == nil {
        request = NewDescribeTDEMSampleRateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "rum", APIVersion, "DescribeTDEMSampleRate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTDEMSampleRate require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTDEMSampleRateResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTDEMSessionReplayTimelineRequest() (request *DescribeTDEMSessionReplayTimelineRequest) {
    request = &DescribeTDEMSessionReplayTimelineRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "DescribeTDEMSessionReplayTimeline")
    
    
    return
}

func NewDescribeTDEMSessionReplayTimelineResponse() (response *DescribeTDEMSessionReplayTimelineResponse) {
    response = &DescribeTDEMSessionReplayTimelineResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTDEMSessionReplayTimeline
// 获取回放时间线数据
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CHARGENOBALANCE = "FailedOperation.ChargeNoBalance"
//  FAILEDOPERATION_CHARGENOPAYRIGHT = "FailedOperation.ChargeNoPayRight"
//  FAILEDOPERATION_CHARGEPARAMINVALID = "FailedOperation.ChargeParamInvalid"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  FAILEDOPERATION_INFRASTRUCTUREERROR = "FailedOperation.InfrastructureError"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERCOMBINATION = "InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
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
func (c *Client) DescribeTDEMSessionReplayTimeline(request *DescribeTDEMSessionReplayTimelineRequest) (response *DescribeTDEMSessionReplayTimelineResponse, err error) {
    return c.DescribeTDEMSessionReplayTimelineWithContext(context.Background(), request)
}

// DescribeTDEMSessionReplayTimeline
// 获取回放时间线数据
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CHARGENOBALANCE = "FailedOperation.ChargeNoBalance"
//  FAILEDOPERATION_CHARGENOPAYRIGHT = "FailedOperation.ChargeNoPayRight"
//  FAILEDOPERATION_CHARGEPARAMINVALID = "FailedOperation.ChargeParamInvalid"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  FAILEDOPERATION_INFRASTRUCTUREERROR = "FailedOperation.InfrastructureError"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERCOMBINATION = "InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
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
func (c *Client) DescribeTDEMSessionReplayTimelineWithContext(ctx context.Context, request *DescribeTDEMSessionReplayTimelineRequest) (response *DescribeTDEMSessionReplayTimelineResponse, err error) {
    if request == nil {
        request = NewDescribeTDEMSessionReplayTimelineRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "rum", APIVersion, "DescribeTDEMSessionReplayTimeline")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTDEMSessionReplayTimeline require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTDEMSessionReplayTimelineResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTDEMUsageDetailRequest() (request *DescribeTDEMUsageDetailRequest) {
    request = &DescribeTDEMUsageDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "DescribeTDEMUsageDetail")
    
    
    return
}

func NewDescribeTDEMUsageDetailResponse() (response *DescribeTDEMUsageDetailResponse) {
    response = &DescribeTDEMUsageDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTDEMUsageDetail
// 按业务系统/应用查询用量明细（含行趋势、端类型、删除/停用标记）。Keyword 匹配业务系统名、应用名或应用 ID。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERCOMBINATION = "InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeTDEMUsageDetail(request *DescribeTDEMUsageDetailRequest) (response *DescribeTDEMUsageDetailResponse, err error) {
    return c.DescribeTDEMUsageDetailWithContext(context.Background(), request)
}

// DescribeTDEMUsageDetail
// 按业务系统/应用查询用量明细（含行趋势、端类型、删除/停用标记）。Keyword 匹配业务系统名、应用名或应用 ID。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERCOMBINATION = "InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeTDEMUsageDetailWithContext(ctx context.Context, request *DescribeTDEMUsageDetailRequest) (response *DescribeTDEMUsageDetailResponse, err error) {
    if request == nil {
        request = NewDescribeTDEMUsageDetailRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "rum", APIVersion, "DescribeTDEMUsageDetail")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTDEMUsageDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTDEMUsageDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTDEMUsageSummaryRequest() (request *DescribeTDEMUsageSummaryRequest) {
    request = &DescribeTDEMUsageSummaryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "DescribeTDEMUsageSummary")
    
    
    return
}

func NewDescribeTDEMUsageSummaryResponse() (response *DescribeTDEMUsageSummaryResponse) {
    response = &DescribeTDEMUsageSummaryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTDEMUsageSummary
// 查询当前使用者名下（超管为全部）上报总量与趋势。时间 StartTime/EndTime 为 YYYY-MM-DD HH:00:00 小时对齐。不要求请求体带 AppId。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERCOMBINATION = "InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeTDEMUsageSummary(request *DescribeTDEMUsageSummaryRequest) (response *DescribeTDEMUsageSummaryResponse, err error) {
    return c.DescribeTDEMUsageSummaryWithContext(context.Background(), request)
}

// DescribeTDEMUsageSummary
// 查询当前使用者名下（超管为全部）上报总量与趋势。时间 StartTime/EndTime 为 YYYY-MM-DD HH:00:00 小时对齐。不要求请求体带 AppId。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERCOMBINATION = "InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeTDEMUsageSummaryWithContext(ctx context.Context, request *DescribeTDEMUsageSummaryRequest) (response *DescribeTDEMUsageSummaryResponse, err error) {
    if request == nil {
        request = NewDescribeTDEMUsageSummaryRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "rum", APIVersion, "DescribeTDEMUsageSummary")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTDEMUsageSummary require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTDEMUsageSummaryResponse()
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

func NewDescribeTokenRequest() (request *DescribeTokenRequest) {
    request = &DescribeTokenRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "DescribeToken")
    
    
    return
}

func NewDescribeTokenResponse() (response *DescribeTokenResponse) {
    response = &DescribeTokenResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeToken
// 查询令牌
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeToken(request *DescribeTokenRequest) (response *DescribeTokenResponse, err error) {
    return c.DescribeTokenWithContext(context.Background(), request)
}

// DescribeToken
// 查询令牌
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeTokenWithContext(ctx context.Context, request *DescribeTokenRequest) (response *DescribeTokenResponse, err error) {
    if request == nil {
        request = NewDescribeTokenRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "rum", APIVersion, "DescribeToken")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeToken require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTokenResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTopIssuesRequest() (request *DescribeTopIssuesRequest) {
    request = &DescribeTopIssuesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "DescribeTopIssues")
    
    
    return
}

func NewDescribeTopIssuesResponse() (response *DescribeTopIssuesResponse) {
    response = &DescribeTopIssuesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTopIssues
// 查询top问题列表
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeTopIssues(request *DescribeTopIssuesRequest) (response *DescribeTopIssuesResponse, err error) {
    return c.DescribeTopIssuesWithContext(context.Background(), request)
}

// DescribeTopIssues
// 查询top问题列表
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeTopIssuesWithContext(ctx context.Context, request *DescribeTopIssuesRequest) (response *DescribeTopIssuesResponse, err error) {
    if request == nil {
        request = NewDescribeTopIssuesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "rum", APIVersion, "DescribeTopIssues")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTopIssues require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTopIssuesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUserJourneySanKeyRequest() (request *DescribeUserJourneySanKeyRequest) {
    request = &DescribeUserJourneySanKeyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "DescribeUserJourneySanKey")
    
    
    return
}

func NewDescribeUserJourneySanKeyResponse() (response *DescribeUserJourneySanKeyResponse) {
    response = &DescribeUserJourneySanKeyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeUserJourneySanKey
// 获取用户旅程数据
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CHARGENOBALANCE = "FailedOperation.ChargeNoBalance"
//  FAILEDOPERATION_CHARGENOPAYRIGHT = "FailedOperation.ChargeNoPayRight"
//  FAILEDOPERATION_CHARGEPARAMINVALID = "FailedOperation.ChargeParamInvalid"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  FAILEDOPERATION_INFRASTRUCTUREERROR = "FailedOperation.InfrastructureError"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERCOMBINATION = "InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
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
func (c *Client) DescribeUserJourneySanKey(request *DescribeUserJourneySanKeyRequest) (response *DescribeUserJourneySanKeyResponse, err error) {
    return c.DescribeUserJourneySanKeyWithContext(context.Background(), request)
}

// DescribeUserJourneySanKey
// 获取用户旅程数据
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CHARGENOBALANCE = "FailedOperation.ChargeNoBalance"
//  FAILEDOPERATION_CHARGENOPAYRIGHT = "FailedOperation.ChargeNoPayRight"
//  FAILEDOPERATION_CHARGEPARAMINVALID = "FailedOperation.ChargeParamInvalid"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  FAILEDOPERATION_INFRASTRUCTUREERROR = "FailedOperation.InfrastructureError"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERCOMBINATION = "InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
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
func (c *Client) DescribeUserJourneySanKeyWithContext(ctx context.Context, request *DescribeUserJourneySanKeyRequest) (response *DescribeUserJourneySanKeyResponse, err error) {
    if request == nil {
        request = NewDescribeUserJourneySanKeyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "rum", APIVersion, "DescribeUserJourneySanKey")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUserJourneySanKey require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeUserJourneySanKeyResponse()
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

func NewModifyFieldTransformRuleRequest() (request *ModifyFieldTransformRuleRequest) {
    request = &ModifyFieldTransformRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "ModifyFieldTransformRule")
    
    
    return
}

func NewModifyFieldTransformRuleResponse() (response *ModifyFieldTransformRuleResponse) {
    response = &ModifyFieldTransformRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyFieldTransformRule
// 修改字段配置规则
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CHARGENOBALANCE = "FailedOperation.ChargeNoBalance"
//  FAILEDOPERATION_CHARGENOPAYRIGHT = "FailedOperation.ChargeNoPayRight"
//  FAILEDOPERATION_CHARGEPARAMINVALID = "FailedOperation.ChargeParamInvalid"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  FAILEDOPERATION_INFRASTRUCTUREERROR = "FailedOperation.InfrastructureError"
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
//  RESOURCENOTFOUND_NOINSTANCE = "ResourceNotFound.NoInstance"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyFieldTransformRule(request *ModifyFieldTransformRuleRequest) (response *ModifyFieldTransformRuleResponse, err error) {
    return c.ModifyFieldTransformRuleWithContext(context.Background(), request)
}

// ModifyFieldTransformRule
// 修改字段配置规则
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CHARGENOBALANCE = "FailedOperation.ChargeNoBalance"
//  FAILEDOPERATION_CHARGENOPAYRIGHT = "FailedOperation.ChargeNoPayRight"
//  FAILEDOPERATION_CHARGEPARAMINVALID = "FailedOperation.ChargeParamInvalid"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  FAILEDOPERATION_INFRASTRUCTUREERROR = "FailedOperation.InfrastructureError"
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
//  RESOURCENOTFOUND_NOINSTANCE = "ResourceNotFound.NoInstance"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyFieldTransformRuleWithContext(ctx context.Context, request *ModifyFieldTransformRuleRequest) (response *ModifyFieldTransformRuleResponse, err error) {
    if request == nil {
        request = NewModifyFieldTransformRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "rum", APIVersion, "ModifyFieldTransformRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyFieldTransformRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyFieldTransformRuleResponse()
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

func NewModifyIssueStatusRequest() (request *ModifyIssueStatusRequest) {
    request = &ModifyIssueStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "ModifyIssueStatus")
    
    
    return
}

func NewModifyIssueStatusResponse() (response *ModifyIssueStatusResponse) {
    response = &ModifyIssueStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyIssueStatus
// 更新 Issue 状态，支持 resolve、ignore、reopen 等状态流转动作
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CHARGENOBALANCE = "FailedOperation.ChargeNoBalance"
//  FAILEDOPERATION_CHARGENOPAYRIGHT = "FailedOperation.ChargeNoPayRight"
//  FAILEDOPERATION_CHARGEPARAMINVALID = "FailedOperation.ChargeParamInvalid"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  FAILEDOPERATION_INFRASTRUCTUREERROR = "FailedOperation.InfrastructureError"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERCOMBINATION = "InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
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
func (c *Client) ModifyIssueStatus(request *ModifyIssueStatusRequest) (response *ModifyIssueStatusResponse, err error) {
    return c.ModifyIssueStatusWithContext(context.Background(), request)
}

// ModifyIssueStatus
// 更新 Issue 状态，支持 resolve、ignore、reopen 等状态流转动作
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CHARGENOBALANCE = "FailedOperation.ChargeNoBalance"
//  FAILEDOPERATION_CHARGENOPAYRIGHT = "FailedOperation.ChargeNoPayRight"
//  FAILEDOPERATION_CHARGEPARAMINVALID = "FailedOperation.ChargeParamInvalid"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  FAILEDOPERATION_INFRASTRUCTUREERROR = "FailedOperation.InfrastructureError"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERCOMBINATION = "InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
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
func (c *Client) ModifyIssueStatusWithContext(ctx context.Context, request *ModifyIssueStatusRequest) (response *ModifyIssueStatusResponse, err error) {
    if request == nil {
        request = NewModifyIssueStatusRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "rum", APIVersion, "ModifyIssueStatus")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyIssueStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyIssueStatusResponse()
    err = c.Send(request, response)
    return
}

func NewModifyMiniProgramReplayPackRequest() (request *ModifyMiniProgramReplayPackRequest) {
    request = &ModifyMiniProgramReplayPackRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "ModifyMiniProgramReplayPack")
    
    
    return
}

func NewModifyMiniProgramReplayPackResponse() (response *ModifyMiniProgramReplayPackResponse) {
    response = &ModifyMiniProgramReplayPackResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyMiniProgramReplayPack
// 更新回放包元数据，仅支持改Release，不重编译。换源码请再调UploadMiniProgramSource。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CHARGENOBALANCE = "FailedOperation.ChargeNoBalance"
//  FAILEDOPERATION_CHARGENOPAYRIGHT = "FailedOperation.ChargeNoPayRight"
//  FAILEDOPERATION_CHARGEPARAMINVALID = "FailedOperation.ChargeParamInvalid"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  FAILEDOPERATION_INFRASTRUCTUREERROR = "FailedOperation.InfrastructureError"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERCOMBINATION = "InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
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
func (c *Client) ModifyMiniProgramReplayPack(request *ModifyMiniProgramReplayPackRequest) (response *ModifyMiniProgramReplayPackResponse, err error) {
    return c.ModifyMiniProgramReplayPackWithContext(context.Background(), request)
}

// ModifyMiniProgramReplayPack
// 更新回放包元数据，仅支持改Release，不重编译。换源码请再调UploadMiniProgramSource。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CHARGENOBALANCE = "FailedOperation.ChargeNoBalance"
//  FAILEDOPERATION_CHARGENOPAYRIGHT = "FailedOperation.ChargeNoPayRight"
//  FAILEDOPERATION_CHARGEPARAMINVALID = "FailedOperation.ChargeParamInvalid"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  FAILEDOPERATION_INFRASTRUCTUREERROR = "FailedOperation.InfrastructureError"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERCOMBINATION = "InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
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
func (c *Client) ModifyMiniProgramReplayPackWithContext(ctx context.Context, request *ModifyMiniProgramReplayPackRequest) (response *ModifyMiniProgramReplayPackResponse, err error) {
    if request == nil {
        request = NewModifyMiniProgramReplayPackRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "rum", APIVersion, "ModifyMiniProgramReplayPack")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyMiniProgramReplayPack require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyMiniProgramReplayPackResponse()
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

func NewModifySampleRateRequest() (request *ModifySampleRateRequest) {
    request = &ModifySampleRateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "ModifySampleRate")
    
    
    return
}

func NewModifySampleRateResponse() (response *ModifySampleRateResponse) {
    response = &ModifySampleRateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifySampleRate
// 配置tdem采样率
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CHARGENOBALANCE = "FailedOperation.ChargeNoBalance"
//  FAILEDOPERATION_CHARGENOPAYRIGHT = "FailedOperation.ChargeNoPayRight"
//  FAILEDOPERATION_CHARGEPARAMINVALID = "FailedOperation.ChargeParamInvalid"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  FAILEDOPERATION_INFRASTRUCTUREERROR = "FailedOperation.InfrastructureError"
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
//  RESOURCENOTFOUND_NOINSTANCE = "ResourceNotFound.NoInstance"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifySampleRate(request *ModifySampleRateRequest) (response *ModifySampleRateResponse, err error) {
    return c.ModifySampleRateWithContext(context.Background(), request)
}

// ModifySampleRate
// 配置tdem采样率
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CHARGENOBALANCE = "FailedOperation.ChargeNoBalance"
//  FAILEDOPERATION_CHARGENOPAYRIGHT = "FailedOperation.ChargeNoPayRight"
//  FAILEDOPERATION_CHARGEPARAMINVALID = "FailedOperation.ChargeParamInvalid"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  FAILEDOPERATION_INFRASTRUCTUREERROR = "FailedOperation.InfrastructureError"
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
//  RESOURCENOTFOUND_NOINSTANCE = "ResourceNotFound.NoInstance"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifySampleRateWithContext(ctx context.Context, request *ModifySampleRateRequest) (response *ModifySampleRateResponse, err error) {
    if request == nil {
        request = NewModifySampleRateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "rum", APIVersion, "ModifySampleRate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifySampleRate require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifySampleRateResponse()
    err = c.Send(request, response)
    return
}

func NewModifyTDEMApplicationRequest() (request *ModifyTDEMApplicationRequest) {
    request = &ModifyTDEMApplicationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "ModifyTDEMApplication")
    
    
    return
}

func NewModifyTDEMApplicationResponse() (response *ModifyTDEMApplicationResponse) {
    response = &ModifyTDEMApplicationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyTDEMApplication
// 修改应用，仅传需修改字段。AppId 必填。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERCOMBINATION = "InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) ModifyTDEMApplication(request *ModifyTDEMApplicationRequest) (response *ModifyTDEMApplicationResponse, err error) {
    return c.ModifyTDEMApplicationWithContext(context.Background(), request)
}

// ModifyTDEMApplication
// 修改应用，仅传需修改字段。AppId 必填。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERCOMBINATION = "InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) ModifyTDEMApplicationWithContext(ctx context.Context, request *ModifyTDEMApplicationRequest) (response *ModifyTDEMApplicationResponse, err error) {
    if request == nil {
        request = NewModifyTDEMApplicationRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "rum", APIVersion, "ModifyTDEMApplication")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyTDEMApplication require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyTDEMApplicationResponse()
    err = c.Send(request, response)
    return
}

func NewModifyTDEMBusinessSystemRequest() (request *ModifyTDEMBusinessSystemRequest) {
    request = &ModifyTDEMBusinessSystemRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "ModifyTDEMBusinessSystem")
    
    
    return
}

func NewModifyTDEMBusinessSystemResponse() (response *ModifyTDEMBusinessSystemResponse) {
    response = &ModifyTDEMBusinessSystemResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyTDEMBusinessSystem
// 修改业务系统名称或描述。越权 UnauthorizedOperation；不存在或已删除 ResourceNotFound。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERCOMBINATION = "InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) ModifyTDEMBusinessSystem(request *ModifyTDEMBusinessSystemRequest) (response *ModifyTDEMBusinessSystemResponse, err error) {
    return c.ModifyTDEMBusinessSystemWithContext(context.Background(), request)
}

// ModifyTDEMBusinessSystem
// 修改业务系统名称或描述。越权 UnauthorizedOperation；不存在或已删除 ResourceNotFound。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERCOMBINATION = "InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) ModifyTDEMBusinessSystemWithContext(ctx context.Context, request *ModifyTDEMBusinessSystemRequest) (response *ModifyTDEMBusinessSystemResponse, err error) {
    if request == nil {
        request = NewModifyTDEMBusinessSystemRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "rum", APIVersion, "ModifyTDEMBusinessSystem")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyTDEMBusinessSystem require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyTDEMBusinessSystemResponse()
    err = c.Send(request, response)
    return
}

func NewModifyTDEMDialTaskRequest() (request *ModifyTDEMDialTaskRequest) {
    request = &ModifyTDEMDialTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "ModifyTDEMDialTask")
    
    
    return
}

func NewModifyTDEMDialTaskResponse() (response *ModifyTDEMDialTaskResponse) {
    response = &ModifyTDEMDialTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyTDEMDialTask
// 修改拨测任务。未出现在请求中的字段保持原值。启动范围传空字符串表示清空该项。Type 和 ExecutionTime 创建后不可修改。暂停和恢复通过 Status 完成，不提供单独的状态接口。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyTDEMDialTask(request *ModifyTDEMDialTaskRequest) (response *ModifyTDEMDialTaskResponse, err error) {
    return c.ModifyTDEMDialTaskWithContext(context.Background(), request)
}

// ModifyTDEMDialTask
// 修改拨测任务。未出现在请求中的字段保持原值。启动范围传空字符串表示清空该项。Type 和 ExecutionTime 创建后不可修改。暂停和恢复通过 Status 完成，不提供单独的状态接口。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyTDEMDialTaskWithContext(ctx context.Context, request *ModifyTDEMDialTaskRequest) (response *ModifyTDEMDialTaskResponse, err error) {
    if request == nil {
        request = NewModifyTDEMDialTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "rum", APIVersion, "ModifyTDEMDialTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyTDEMDialTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyTDEMDialTaskResponse()
    err = c.Send(request, response)
    return
}

func NewModifyTDEMFieldTransformRuleRequest() (request *ModifyTDEMFieldTransformRuleRequest) {
    request = &ModifyTDEMFieldTransformRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "ModifyTDEMFieldTransformRule")
    
    
    return
}

func NewModifyTDEMFieldTransformRuleResponse() (response *ModifyTDEMFieldTransformRuleResponse) {
    response = &ModifyTDEMFieldTransformRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyTDEMFieldTransformRule
// 按规则 Id 修改字段值替换规则。替代已废弃的 ModifyFieldTransformRule；AppId 为 string。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERCOMBINATION = "InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) ModifyTDEMFieldTransformRule(request *ModifyTDEMFieldTransformRuleRequest) (response *ModifyTDEMFieldTransformRuleResponse, err error) {
    return c.ModifyTDEMFieldTransformRuleWithContext(context.Background(), request)
}

// ModifyTDEMFieldTransformRule
// 按规则 Id 修改字段值替换规则。替代已废弃的 ModifyFieldTransformRule；AppId 为 string。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERCOMBINATION = "InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) ModifyTDEMFieldTransformRuleWithContext(ctx context.Context, request *ModifyTDEMFieldTransformRuleRequest) (response *ModifyTDEMFieldTransformRuleResponse, err error) {
    if request == nil {
        request = NewModifyTDEMFieldTransformRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "rum", APIVersion, "ModifyTDEMFieldTransformRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyTDEMFieldTransformRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyTDEMFieldTransformRuleResponse()
    err = c.Send(request, response)
    return
}

func NewModifyTDEMFunnelRequest() (request *ModifyTDEMFunnelRequest) {
    request = &ModifyTDEMFunnelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "ModifyTDEMFunnel")
    
    
    return
}

func NewModifyTDEMFunnelResponse() (response *ModifyTDEMFunnelResponse) {
    response = &ModifyTDEMFunnelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyTDEMFunnel
// 修改漏斗
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CHARGENOBALANCE = "FailedOperation.ChargeNoBalance"
//  FAILEDOPERATION_CHARGENOPAYRIGHT = "FailedOperation.ChargeNoPayRight"
//  FAILEDOPERATION_CHARGEPARAMINVALID = "FailedOperation.ChargeParamInvalid"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  FAILEDOPERATION_INFRASTRUCTUREERROR = "FailedOperation.InfrastructureError"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERCOMBINATION = "InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
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
func (c *Client) ModifyTDEMFunnel(request *ModifyTDEMFunnelRequest) (response *ModifyTDEMFunnelResponse, err error) {
    return c.ModifyTDEMFunnelWithContext(context.Background(), request)
}

// ModifyTDEMFunnel
// 修改漏斗
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CHARGENOBALANCE = "FailedOperation.ChargeNoBalance"
//  FAILEDOPERATION_CHARGENOPAYRIGHT = "FailedOperation.ChargeNoPayRight"
//  FAILEDOPERATION_CHARGEPARAMINVALID = "FailedOperation.ChargeParamInvalid"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  FAILEDOPERATION_INFRASTRUCTUREERROR = "FailedOperation.InfrastructureError"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERCOMBINATION = "InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
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
func (c *Client) ModifyTDEMFunnelWithContext(ctx context.Context, request *ModifyTDEMFunnelRequest) (response *ModifyTDEMFunnelResponse, err error) {
    if request == nil {
        request = NewModifyTDEMFunnelRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "rum", APIVersion, "ModifyTDEMFunnel")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyTDEMFunnel require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyTDEMFunnelResponse()
    err = c.Send(request, response)
    return
}

func NewModifyTDEMSampleConfigRequest() (request *ModifyTDEMSampleConfigRequest) {
    request = &ModifyTDEMSampleConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "ModifyTDEMSampleConfig")
    
    
    return
}

func NewModifyTDEMSampleConfigResponse() (response *ModifyTDEMSampleConfigResponse) {
    response = &ModifyTDEMSampleConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyTDEMSampleConfig
// 整页保存 Web 规则化采样文档：校验 → INSERT 快照 → UPSERT 当前行 → 刷新 Redis 下发缓存。不写 tdem_sample_rate。采样率 0–100、最多两位小数；规则 Conditions 不可为空；Types 含非闭集 key 返回 InvalidParameter。校验失败不写库。无 Restore/Import/Export/Diff。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERCOMBINATION = "InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) ModifyTDEMSampleConfig(request *ModifyTDEMSampleConfigRequest) (response *ModifyTDEMSampleConfigResponse, err error) {
    return c.ModifyTDEMSampleConfigWithContext(context.Background(), request)
}

// ModifyTDEMSampleConfig
// 整页保存 Web 规则化采样文档：校验 → INSERT 快照 → UPSERT 当前行 → 刷新 Redis 下发缓存。不写 tdem_sample_rate。采样率 0–100、最多两位小数；规则 Conditions 不可为空；Types 含非闭集 key 返回 InvalidParameter。校验失败不写库。无 Restore/Import/Export/Diff。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERCOMBINATION = "InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) ModifyTDEMSampleConfigWithContext(ctx context.Context, request *ModifyTDEMSampleConfigRequest) (response *ModifyTDEMSampleConfigResponse, err error) {
    if request == nil {
        request = NewModifyTDEMSampleConfigRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "rum", APIVersion, "ModifyTDEMSampleConfig")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyTDEMSampleConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyTDEMSampleConfigResponse()
    err = c.Send(request, response)
    return
}

func NewModifyTDEMSampleRateRequest() (request *ModifyTDEMSampleRateRequest) {
    request = &ModifyTDEMSampleRateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "ModifyTDEMSampleRate")
    
    
    return
}

func NewModifyTDEMSampleRateResponse() (response *ModifyTDEMSampleRateResponse) {
    response = &ModifyTDEMSampleRateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyTDEMSampleRate
// 设置或更新应用采集采样率；ReplaySampleRate 可选。替代已废弃的 ModifySampleRate；AppId 为 string。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERCOMBINATION = "InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) ModifyTDEMSampleRate(request *ModifyTDEMSampleRateRequest) (response *ModifyTDEMSampleRateResponse, err error) {
    return c.ModifyTDEMSampleRateWithContext(context.Background(), request)
}

// ModifyTDEMSampleRate
// 设置或更新应用采集采样率；ReplaySampleRate 可选。替代已废弃的 ModifySampleRate；AppId 为 string。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERCOMBINATION = "InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) ModifyTDEMSampleRateWithContext(ctx context.Context, request *ModifyTDEMSampleRateRequest) (response *ModifyTDEMSampleRateResponse, err error) {
    if request == nil {
        request = NewModifyTDEMSampleRateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "rum", APIVersion, "ModifyTDEMSampleRate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyTDEMSampleRate require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyTDEMSampleRateResponse()
    err = c.Send(request, response)
    return
}

func NewPreviewTDEMFunnelReportRequest() (request *PreviewTDEMFunnelReportRequest) {
    request = &PreviewTDEMFunnelReportRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "PreviewTDEMFunnelReport")
    
    
    return
}

func NewPreviewTDEMFunnelReportResponse() (response *PreviewTDEMFunnelReportResponse) {
    response = &PreviewTDEMFunnelReportResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// PreviewTDEMFunnelReport
// 根据步骤配置直接生成漏斗预览报告（不保存漏斗，不需要FunnelId）
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CHARGENOBALANCE = "FailedOperation.ChargeNoBalance"
//  FAILEDOPERATION_CHARGENOPAYRIGHT = "FailedOperation.ChargeNoPayRight"
//  FAILEDOPERATION_CHARGEPARAMINVALID = "FailedOperation.ChargeParamInvalid"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  FAILEDOPERATION_INFRASTRUCTUREERROR = "FailedOperation.InfrastructureError"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERCOMBINATION = "InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
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
func (c *Client) PreviewTDEMFunnelReport(request *PreviewTDEMFunnelReportRequest) (response *PreviewTDEMFunnelReportResponse, err error) {
    return c.PreviewTDEMFunnelReportWithContext(context.Background(), request)
}

// PreviewTDEMFunnelReport
// 根据步骤配置直接生成漏斗预览报告（不保存漏斗，不需要FunnelId）
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CHARGENOBALANCE = "FailedOperation.ChargeNoBalance"
//  FAILEDOPERATION_CHARGENOPAYRIGHT = "FailedOperation.ChargeNoPayRight"
//  FAILEDOPERATION_CHARGEPARAMINVALID = "FailedOperation.ChargeParamInvalid"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  FAILEDOPERATION_INFRASTRUCTUREERROR = "FailedOperation.InfrastructureError"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERCOMBINATION = "InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
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
func (c *Client) PreviewTDEMFunnelReportWithContext(ctx context.Context, request *PreviewTDEMFunnelReportRequest) (response *PreviewTDEMFunnelReportResponse, err error) {
    if request == nil {
        request = NewPreviewTDEMFunnelReportRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "rum", APIVersion, "PreviewTDEMFunnelReport")
    
    if c.GetCredential() == nil {
        return nil, errors.New("PreviewTDEMFunnelReport require credential")
    }

    request.SetContext(ctx)
    
    response = NewPreviewTDEMFunnelReportResponse()
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

func NewResumeTDEMApplicationRequest() (request *ResumeTDEMApplicationRequest) {
    request = &ResumeTDEMApplicationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "ResumeTDEMApplication")
    
    
    return
}

func NewResumeTDEMApplicationResponse() (response *ResumeTDEMApplicationResponse) {
    response = &ResumeTDEMApplicationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ResumeTDEMApplication
// 恢复已停止的应用，Status 变为 2。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERCOMBINATION = "InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) ResumeTDEMApplication(request *ResumeTDEMApplicationRequest) (response *ResumeTDEMApplicationResponse, err error) {
    return c.ResumeTDEMApplicationWithContext(context.Background(), request)
}

// ResumeTDEMApplication
// 恢复已停止的应用，Status 变为 2。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERCOMBINATION = "InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) ResumeTDEMApplicationWithContext(ctx context.Context, request *ResumeTDEMApplicationRequest) (response *ResumeTDEMApplicationResponse, err error) {
    if request == nil {
        request = NewResumeTDEMApplicationRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "rum", APIVersion, "ResumeTDEMApplication")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ResumeTDEMApplication require credential")
    }

    request.SetContext(ctx)
    
    response = NewResumeTDEMApplicationResponse()
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

func NewStopTDEMApplicationRequest() (request *StopTDEMApplicationRequest) {
    request = &StopTDEMApplicationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "StopTDEMApplication")
    
    
    return
}

func NewStopTDEMApplicationResponse() (response *StopTDEMApplicationResponse) {
    response = &StopTDEMApplicationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// StopTDEMApplication
// 停止应用采集，Status 变为 6。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERCOMBINATION = "InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) StopTDEMApplication(request *StopTDEMApplicationRequest) (response *StopTDEMApplicationResponse, err error) {
    return c.StopTDEMApplicationWithContext(context.Background(), request)
}

// StopTDEMApplication
// 停止应用采集，Status 变为 6。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERCOMBINATION = "InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) StopTDEMApplicationWithContext(ctx context.Context, request *StopTDEMApplicationRequest) (response *StopTDEMApplicationResponse, err error) {
    if request == nil {
        request = NewStopTDEMApplicationRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "rum", APIVersion, "StopTDEMApplication")
    
    if c.GetCredential() == nil {
        return nil, errors.New("StopTDEMApplication require credential")
    }

    request.SetContext(ctx)
    
    response = NewStopTDEMApplicationResponse()
    err = c.Send(request, response)
    return
}

func NewUploadMiniProgramSourceRequest() (request *UploadMiniProgramSourceRequest) {
    request = &UploadMiniProgramSourceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "UploadMiniProgramSource")
    
    
    return
}

func NewUploadMiniProgramSourceResponse() (response *UploadMiniProgramSourceResponse) {
    response = &UploadMiniProgramSourceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UploadMiniProgramSource
// 上传微信小程序源码zip并同步编译回放包。控制台传FileBase64；zip需含app.json，跳过node_modules。同AppId+BundleHash且已ready则幂等。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CHARGENOBALANCE = "FailedOperation.ChargeNoBalance"
//  FAILEDOPERATION_CHARGENOPAYRIGHT = "FailedOperation.ChargeNoPayRight"
//  FAILEDOPERATION_CHARGEPARAMINVALID = "FailedOperation.ChargeParamInvalid"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  FAILEDOPERATION_INFRASTRUCTUREERROR = "FailedOperation.InfrastructureError"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERCOMBINATION = "InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
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
func (c *Client) UploadMiniProgramSource(request *UploadMiniProgramSourceRequest) (response *UploadMiniProgramSourceResponse, err error) {
    return c.UploadMiniProgramSourceWithContext(context.Background(), request)
}

// UploadMiniProgramSource
// 上传微信小程序源码zip并同步编译回放包。控制台传FileBase64；zip需含app.json，跳过node_modules。同AppId+BundleHash且已ready则幂等。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CHARGENOBALANCE = "FailedOperation.ChargeNoBalance"
//  FAILEDOPERATION_CHARGENOPAYRIGHT = "FailedOperation.ChargeNoPayRight"
//  FAILEDOPERATION_CHARGEPARAMINVALID = "FailedOperation.ChargeParamInvalid"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  FAILEDOPERATION_INFRASTRUCTUREERROR = "FailedOperation.InfrastructureError"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERCOMBINATION = "InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
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
func (c *Client) UploadMiniProgramSourceWithContext(ctx context.Context, request *UploadMiniProgramSourceRequest) (response *UploadMiniProgramSourceResponse, err error) {
    if request == nil {
        request = NewUploadMiniProgramSourceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "rum", APIVersion, "UploadMiniProgramSource")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UploadMiniProgramSource require credential")
    }

    request.SetContext(ctx)
    
    response = NewUploadMiniProgramSourceResponse()
    err = c.Send(request, response)
    return
}

func NewUploadSymbolArtifactRequest() (request *UploadSymbolArtifactRequest) {
    request = &UploadSymbolArtifactRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rum", APIVersion, "UploadSymbolArtifact")
    
    
    return
}

func NewUploadSymbolArtifactResponse() (response *UploadSymbolArtifactResponse) {
    response = &UploadSymbolArtifactResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UploadSymbolArtifact
// 登记已上传的符号制品。请先获取上传凭证并将文件上传到 COS。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CHARGENOBALANCE = "FailedOperation.ChargeNoBalance"
//  FAILEDOPERATION_CHARGENOPAYRIGHT = "FailedOperation.ChargeNoPayRight"
//  FAILEDOPERATION_CHARGEPARAMINVALID = "FailedOperation.ChargeParamInvalid"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  FAILEDOPERATION_INFRASTRUCTUREERROR = "FailedOperation.InfrastructureError"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERCOMBINATION = "InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
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
func (c *Client) UploadSymbolArtifact(request *UploadSymbolArtifactRequest) (response *UploadSymbolArtifactResponse, err error) {
    return c.UploadSymbolArtifactWithContext(context.Background(), request)
}

// UploadSymbolArtifact
// 登记已上传的符号制品。请先获取上传凭证并将文件上传到 COS。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CHARGENOBALANCE = "FailedOperation.ChargeNoBalance"
//  FAILEDOPERATION_CHARGENOPAYRIGHT = "FailedOperation.ChargeNoPayRight"
//  FAILEDOPERATION_CHARGEPARAMINVALID = "FailedOperation.ChargeParamInvalid"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  FAILEDOPERATION_INFRASTRUCTUREERROR = "FailedOperation.InfrastructureError"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERCOMBINATION = "InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
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
func (c *Client) UploadSymbolArtifactWithContext(ctx context.Context, request *UploadSymbolArtifactRequest) (response *UploadSymbolArtifactResponse, err error) {
    if request == nil {
        request = NewUploadSymbolArtifactRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "rum", APIVersion, "UploadSymbolArtifact")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UploadSymbolArtifact require credential")
    }

    request.SetContext(ctx)
    
    response = NewUploadSymbolArtifactResponse()
    err = c.Send(request, response)
    return
}
