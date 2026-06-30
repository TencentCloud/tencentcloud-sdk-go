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

package v20230418

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2023-04-18"

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


func NewAddCloudNativeAPIGatewayConsumerGroupAuthRequest() (request *AddCloudNativeAPIGatewayConsumerGroupAuthRequest) {
    request = &AddCloudNativeAPIGatewayConsumerGroupAuthRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cngw", APIVersion, "AddCloudNativeAPIGatewayConsumerGroupAuth")
    
    
    return
}

func NewAddCloudNativeAPIGatewayConsumerGroupAuthResponse() (response *AddCloudNativeAPIGatewayConsumerGroupAuthResponse) {
    response = &AddCloudNativeAPIGatewayConsumerGroupAuthResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AddCloudNativeAPIGatewayConsumerGroupAuth
// 为资源（模型 API / MCP Server）添加消费者组授权。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  FAILEDOPERATION_VPC = "FailedOperation.Vpc"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  INVALIDPARAMETERVALUE_SPECIFICATION = "InvalidParameterValue.Specification"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_CAMPASSROLENOTEXIST = "UnauthorizedOperation.CamPassRoleNotExist"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) AddCloudNativeAPIGatewayConsumerGroupAuth(request *AddCloudNativeAPIGatewayConsumerGroupAuthRequest) (response *AddCloudNativeAPIGatewayConsumerGroupAuthResponse, err error) {
    return c.AddCloudNativeAPIGatewayConsumerGroupAuthWithContext(context.Background(), request)
}

// AddCloudNativeAPIGatewayConsumerGroupAuth
// 为资源（模型 API / MCP Server）添加消费者组授权。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  FAILEDOPERATION_VPC = "FailedOperation.Vpc"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  INVALIDPARAMETERVALUE_SPECIFICATION = "InvalidParameterValue.Specification"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_CAMPASSROLENOTEXIST = "UnauthorizedOperation.CamPassRoleNotExist"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) AddCloudNativeAPIGatewayConsumerGroupAuthWithContext(ctx context.Context, request *AddCloudNativeAPIGatewayConsumerGroupAuthRequest) (response *AddCloudNativeAPIGatewayConsumerGroupAuthResponse, err error) {
    if request == nil {
        request = NewAddCloudNativeAPIGatewayConsumerGroupAuthRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cngw", APIVersion, "AddCloudNativeAPIGatewayConsumerGroupAuth")
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddCloudNativeAPIGatewayConsumerGroupAuth require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddCloudNativeAPIGatewayConsumerGroupAuthResponse()
    err = c.Send(request, response)
    return
}

func NewAddCloudNativeAPIGatewayConsumerInGroupRequest() (request *AddCloudNativeAPIGatewayConsumerInGroupRequest) {
    request = &AddCloudNativeAPIGatewayConsumerInGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cngw", APIVersion, "AddCloudNativeAPIGatewayConsumerInGroup")
    
    
    return
}

func NewAddCloudNativeAPIGatewayConsumerInGroupResponse() (response *AddCloudNativeAPIGatewayConsumerInGroupResponse) {
    response = &AddCloudNativeAPIGatewayConsumerInGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AddCloudNativeAPIGatewayConsumerInGroup
// 将消费者添加到消费者组。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  FAILEDOPERATION_VPC = "FailedOperation.Vpc"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  INVALIDPARAMETERVALUE_SPECIFICATION = "InvalidParameterValue.Specification"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_CAMPASSROLENOTEXIST = "UnauthorizedOperation.CamPassRoleNotExist"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) AddCloudNativeAPIGatewayConsumerInGroup(request *AddCloudNativeAPIGatewayConsumerInGroupRequest) (response *AddCloudNativeAPIGatewayConsumerInGroupResponse, err error) {
    return c.AddCloudNativeAPIGatewayConsumerInGroupWithContext(context.Background(), request)
}

// AddCloudNativeAPIGatewayConsumerInGroup
// 将消费者添加到消费者组。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  FAILEDOPERATION_VPC = "FailedOperation.Vpc"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  INVALIDPARAMETERVALUE_SPECIFICATION = "InvalidParameterValue.Specification"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_CAMPASSROLENOTEXIST = "UnauthorizedOperation.CamPassRoleNotExist"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) AddCloudNativeAPIGatewayConsumerInGroupWithContext(ctx context.Context, request *AddCloudNativeAPIGatewayConsumerInGroupRequest) (response *AddCloudNativeAPIGatewayConsumerInGroupResponse, err error) {
    if request == nil {
        request = NewAddCloudNativeAPIGatewayConsumerInGroupRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cngw", APIVersion, "AddCloudNativeAPIGatewayConsumerInGroup")
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddCloudNativeAPIGatewayConsumerInGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddCloudNativeAPIGatewayConsumerInGroupResponse()
    err = c.Send(request, response)
    return
}

func NewBindCloudNativeAPIGatewaySecretKeyRequest() (request *BindCloudNativeAPIGatewaySecretKeyRequest) {
    request = &BindCloudNativeAPIGatewaySecretKeyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cngw", APIVersion, "BindCloudNativeAPIGatewaySecretKey")
    
    
    return
}

func NewBindCloudNativeAPIGatewaySecretKeyResponse() (response *BindCloudNativeAPIGatewaySecretKeyResponse) {
    response = &BindCloudNativeAPIGatewaySecretKeyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// BindCloudNativeAPIGatewaySecretKey
// 添加密钥与资源的引用关系接口
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  FAILEDOPERATION_VPC = "FailedOperation.Vpc"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  INVALIDPARAMETERVALUE_SPECIFICATION = "InvalidParameterValue.Specification"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_CAMPASSROLENOTEXIST = "UnauthorizedOperation.CamPassRoleNotExist"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) BindCloudNativeAPIGatewaySecretKey(request *BindCloudNativeAPIGatewaySecretKeyRequest) (response *BindCloudNativeAPIGatewaySecretKeyResponse, err error) {
    return c.BindCloudNativeAPIGatewaySecretKeyWithContext(context.Background(), request)
}

// BindCloudNativeAPIGatewaySecretKey
// 添加密钥与资源的引用关系接口
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  FAILEDOPERATION_VPC = "FailedOperation.Vpc"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  INVALIDPARAMETERVALUE_SPECIFICATION = "InvalidParameterValue.Specification"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_CAMPASSROLENOTEXIST = "UnauthorizedOperation.CamPassRoleNotExist"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) BindCloudNativeAPIGatewaySecretKeyWithContext(ctx context.Context, request *BindCloudNativeAPIGatewaySecretKeyRequest) (response *BindCloudNativeAPIGatewaySecretKeyResponse, err error) {
    if request == nil {
        request = NewBindCloudNativeAPIGatewaySecretKeyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cngw", APIVersion, "BindCloudNativeAPIGatewaySecretKey")
    
    if c.GetCredential() == nil {
        return nil, errors.New("BindCloudNativeAPIGatewaySecretKey require credential")
    }

    request.SetContext(ctx)
    
    response = NewBindCloudNativeAPIGatewaySecretKeyResponse()
    err = c.Send(request, response)
    return
}

func NewCreateCloudNativeAPIGatewayConsumerRequest() (request *CreateCloudNativeAPIGatewayConsumerRequest) {
    request = &CreateCloudNativeAPIGatewayConsumerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cngw", APIVersion, "CreateCloudNativeAPIGatewayConsumer")
    
    
    return
}

func NewCreateCloudNativeAPIGatewayConsumerResponse() (response *CreateCloudNativeAPIGatewayConsumerResponse) {
    response = &CreateCloudNativeAPIGatewayConsumerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateCloudNativeAPIGatewayConsumer
// 创建AI网关消费者。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  FAILEDOPERATION_VPC = "FailedOperation.Vpc"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  INVALIDPARAMETERVALUE_SPECIFICATION = "InvalidParameterValue.Specification"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_CAMPASSROLENOTEXIST = "UnauthorizedOperation.CamPassRoleNotExist"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) CreateCloudNativeAPIGatewayConsumer(request *CreateCloudNativeAPIGatewayConsumerRequest) (response *CreateCloudNativeAPIGatewayConsumerResponse, err error) {
    return c.CreateCloudNativeAPIGatewayConsumerWithContext(context.Background(), request)
}

// CreateCloudNativeAPIGatewayConsumer
// 创建AI网关消费者。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  FAILEDOPERATION_VPC = "FailedOperation.Vpc"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  INVALIDPARAMETERVALUE_SPECIFICATION = "InvalidParameterValue.Specification"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_CAMPASSROLENOTEXIST = "UnauthorizedOperation.CamPassRoleNotExist"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) CreateCloudNativeAPIGatewayConsumerWithContext(ctx context.Context, request *CreateCloudNativeAPIGatewayConsumerRequest) (response *CreateCloudNativeAPIGatewayConsumerResponse, err error) {
    if request == nil {
        request = NewCreateCloudNativeAPIGatewayConsumerRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cngw", APIVersion, "CreateCloudNativeAPIGatewayConsumer")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCloudNativeAPIGatewayConsumer require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateCloudNativeAPIGatewayConsumerResponse()
    err = c.Send(request, response)
    return
}

func NewCreateCloudNativeAPIGatewayConsumerGroupRequest() (request *CreateCloudNativeAPIGatewayConsumerGroupRequest) {
    request = &CreateCloudNativeAPIGatewayConsumerGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cngw", APIVersion, "CreateCloudNativeAPIGatewayConsumerGroup")
    
    
    return
}

func NewCreateCloudNativeAPIGatewayConsumerGroupResponse() (response *CreateCloudNativeAPIGatewayConsumerGroupResponse) {
    response = &CreateCloudNativeAPIGatewayConsumerGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateCloudNativeAPIGatewayConsumerGroup
// 创建AI 网关消费者组
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  FAILEDOPERATION_VPC = "FailedOperation.Vpc"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  INVALIDPARAMETERVALUE_SPECIFICATION = "InvalidParameterValue.Specification"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_CAMPASSROLENOTEXIST = "UnauthorizedOperation.CamPassRoleNotExist"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) CreateCloudNativeAPIGatewayConsumerGroup(request *CreateCloudNativeAPIGatewayConsumerGroupRequest) (response *CreateCloudNativeAPIGatewayConsumerGroupResponse, err error) {
    return c.CreateCloudNativeAPIGatewayConsumerGroupWithContext(context.Background(), request)
}

// CreateCloudNativeAPIGatewayConsumerGroup
// 创建AI 网关消费者组
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  FAILEDOPERATION_VPC = "FailedOperation.Vpc"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  INVALIDPARAMETERVALUE_SPECIFICATION = "InvalidParameterValue.Specification"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_CAMPASSROLENOTEXIST = "UnauthorizedOperation.CamPassRoleNotExist"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) CreateCloudNativeAPIGatewayConsumerGroupWithContext(ctx context.Context, request *CreateCloudNativeAPIGatewayConsumerGroupRequest) (response *CreateCloudNativeAPIGatewayConsumerGroupResponse, err error) {
    if request == nil {
        request = NewCreateCloudNativeAPIGatewayConsumerGroupRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cngw", APIVersion, "CreateCloudNativeAPIGatewayConsumerGroup")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCloudNativeAPIGatewayConsumerGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateCloudNativeAPIGatewayConsumerGroupResponse()
    err = c.Send(request, response)
    return
}

func NewCreateCloudNativeAPIGatewayLLMModelAPIRequest() (request *CreateCloudNativeAPIGatewayLLMModelAPIRequest) {
    request = &CreateCloudNativeAPIGatewayLLMModelAPIRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cngw", APIVersion, "CreateCloudNativeAPIGatewayLLMModelAPI")
    
    
    return
}

func NewCreateCloudNativeAPIGatewayLLMModelAPIResponse() (response *CreateCloudNativeAPIGatewayLLMModelAPIResponse) {
    response = &CreateCloudNativeAPIGatewayLLMModelAPIResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateCloudNativeAPIGatewayLLMModelAPI
// 创建 LLM 模型 API。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  LIMITEXCEEDED_LIMITEXCEEDED = "LimitExceeded.LimitExceeded"
//  MISSINGPARAMETER_MISSPARAMETER = "MissingParameter.MissParameter"
//  OPERATIONDENIED_OPERATIONDENIED = "OperationDenied.OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) CreateCloudNativeAPIGatewayLLMModelAPI(request *CreateCloudNativeAPIGatewayLLMModelAPIRequest) (response *CreateCloudNativeAPIGatewayLLMModelAPIResponse, err error) {
    return c.CreateCloudNativeAPIGatewayLLMModelAPIWithContext(context.Background(), request)
}

// CreateCloudNativeAPIGatewayLLMModelAPI
// 创建 LLM 模型 API。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  LIMITEXCEEDED_LIMITEXCEEDED = "LimitExceeded.LimitExceeded"
//  MISSINGPARAMETER_MISSPARAMETER = "MissingParameter.MissParameter"
//  OPERATIONDENIED_OPERATIONDENIED = "OperationDenied.OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) CreateCloudNativeAPIGatewayLLMModelAPIWithContext(ctx context.Context, request *CreateCloudNativeAPIGatewayLLMModelAPIRequest) (response *CreateCloudNativeAPIGatewayLLMModelAPIResponse, err error) {
    if request == nil {
        request = NewCreateCloudNativeAPIGatewayLLMModelAPIRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cngw", APIVersion, "CreateCloudNativeAPIGatewayLLMModelAPI")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCloudNativeAPIGatewayLLMModelAPI require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateCloudNativeAPIGatewayLLMModelAPIResponse()
    err = c.Send(request, response)
    return
}

func NewCreateCloudNativeAPIGatewayLLMModelServiceRequest() (request *CreateCloudNativeAPIGatewayLLMModelServiceRequest) {
    request = &CreateCloudNativeAPIGatewayLLMModelServiceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cngw", APIVersion, "CreateCloudNativeAPIGatewayLLMModelService")
    
    
    return
}

func NewCreateCloudNativeAPIGatewayLLMModelServiceResponse() (response *CreateCloudNativeAPIGatewayLLMModelServiceResponse) {
    response = &CreateCloudNativeAPIGatewayLLMModelServiceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateCloudNativeAPIGatewayLLMModelService
// 创建 LLM 模型服务。同一网关下 Name 唯一。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  LIMITEXCEEDED_LIMITEXCEEDED = "LimitExceeded.LimitExceeded"
//  MISSINGPARAMETER_MISSPARAMETER = "MissingParameter.MissParameter"
//  OPERATIONDENIED_OPERATIONDENIED = "OperationDenied.OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) CreateCloudNativeAPIGatewayLLMModelService(request *CreateCloudNativeAPIGatewayLLMModelServiceRequest) (response *CreateCloudNativeAPIGatewayLLMModelServiceResponse, err error) {
    return c.CreateCloudNativeAPIGatewayLLMModelServiceWithContext(context.Background(), request)
}

// CreateCloudNativeAPIGatewayLLMModelService
// 创建 LLM 模型服务。同一网关下 Name 唯一。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  LIMITEXCEEDED_LIMITEXCEEDED = "LimitExceeded.LimitExceeded"
//  MISSINGPARAMETER_MISSPARAMETER = "MissingParameter.MissParameter"
//  OPERATIONDENIED_OPERATIONDENIED = "OperationDenied.OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) CreateCloudNativeAPIGatewayLLMModelServiceWithContext(ctx context.Context, request *CreateCloudNativeAPIGatewayLLMModelServiceRequest) (response *CreateCloudNativeAPIGatewayLLMModelServiceResponse, err error) {
    if request == nil {
        request = NewCreateCloudNativeAPIGatewayLLMModelServiceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cngw", APIVersion, "CreateCloudNativeAPIGatewayLLMModelService")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCloudNativeAPIGatewayLLMModelService require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateCloudNativeAPIGatewayLLMModelServiceResponse()
    err = c.Send(request, response)
    return
}

func NewCreateCloudNativeAPIGatewayMCPServerRequest() (request *CreateCloudNativeAPIGatewayMCPServerRequest) {
    request = &CreateCloudNativeAPIGatewayMCPServerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cngw", APIVersion, "CreateCloudNativeAPIGatewayMCPServer")
    
    
    return
}

func NewCreateCloudNativeAPIGatewayMCPServerResponse() (response *CreateCloudNativeAPIGatewayMCPServerResponse) {
    response = &CreateCloudNativeAPIGatewayMCPServerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateCloudNativeAPIGatewayMCPServer
// 创建AI网关MCP Server
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  FAILEDOPERATION_VPC = "FailedOperation.Vpc"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  INVALIDPARAMETERVALUE_SPECIFICATION = "InvalidParameterValue.Specification"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_CAMPASSROLENOTEXIST = "UnauthorizedOperation.CamPassRoleNotExist"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) CreateCloudNativeAPIGatewayMCPServer(request *CreateCloudNativeAPIGatewayMCPServerRequest) (response *CreateCloudNativeAPIGatewayMCPServerResponse, err error) {
    return c.CreateCloudNativeAPIGatewayMCPServerWithContext(context.Background(), request)
}

// CreateCloudNativeAPIGatewayMCPServer
// 创建AI网关MCP Server
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  FAILEDOPERATION_VPC = "FailedOperation.Vpc"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  INVALIDPARAMETERVALUE_SPECIFICATION = "InvalidParameterValue.Specification"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_CAMPASSROLENOTEXIST = "UnauthorizedOperation.CamPassRoleNotExist"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) CreateCloudNativeAPIGatewayMCPServerWithContext(ctx context.Context, request *CreateCloudNativeAPIGatewayMCPServerRequest) (response *CreateCloudNativeAPIGatewayMCPServerResponse, err error) {
    if request == nil {
        request = NewCreateCloudNativeAPIGatewayMCPServerRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cngw", APIVersion, "CreateCloudNativeAPIGatewayMCPServer")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCloudNativeAPIGatewayMCPServer require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateCloudNativeAPIGatewayMCPServerResponse()
    err = c.Send(request, response)
    return
}

func NewCreateCloudNativeAPIGatewayMCPToolRequest() (request *CreateCloudNativeAPIGatewayMCPToolRequest) {
    request = &CreateCloudNativeAPIGatewayMCPToolRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cngw", APIVersion, "CreateCloudNativeAPIGatewayMCPTool")
    
    
    return
}

func NewCreateCloudNativeAPIGatewayMCPToolResponse() (response *CreateCloudNativeAPIGatewayMCPToolResponse) {
    response = &CreateCloudNativeAPIGatewayMCPToolResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateCloudNativeAPIGatewayMCPTool
// 创建AI网关MCP Tool
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  FAILEDOPERATION_VPC = "FailedOperation.Vpc"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  INVALIDPARAMETERVALUE_SPECIFICATION = "InvalidParameterValue.Specification"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_CAMPASSROLENOTEXIST = "UnauthorizedOperation.CamPassRoleNotExist"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) CreateCloudNativeAPIGatewayMCPTool(request *CreateCloudNativeAPIGatewayMCPToolRequest) (response *CreateCloudNativeAPIGatewayMCPToolResponse, err error) {
    return c.CreateCloudNativeAPIGatewayMCPToolWithContext(context.Background(), request)
}

// CreateCloudNativeAPIGatewayMCPTool
// 创建AI网关MCP Tool
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  FAILEDOPERATION_VPC = "FailedOperation.Vpc"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  INVALIDPARAMETERVALUE_SPECIFICATION = "InvalidParameterValue.Specification"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_CAMPASSROLENOTEXIST = "UnauthorizedOperation.CamPassRoleNotExist"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) CreateCloudNativeAPIGatewayMCPToolWithContext(ctx context.Context, request *CreateCloudNativeAPIGatewayMCPToolRequest) (response *CreateCloudNativeAPIGatewayMCPToolResponse, err error) {
    if request == nil {
        request = NewCreateCloudNativeAPIGatewayMCPToolRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cngw", APIVersion, "CreateCloudNativeAPIGatewayMCPTool")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCloudNativeAPIGatewayMCPTool require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateCloudNativeAPIGatewayMCPToolResponse()
    err = c.Send(request, response)
    return
}

func NewCreateCloudNativeAPIGatewaySecretKeyRequest() (request *CreateCloudNativeAPIGatewaySecretKeyRequest) {
    request = &CreateCloudNativeAPIGatewaySecretKeyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cngw", APIVersion, "CreateCloudNativeAPIGatewaySecretKey")
    
    
    return
}

func NewCreateCloudNativeAPIGatewaySecretKeyResponse() (response *CreateCloudNativeAPIGatewaySecretKeyResponse) {
    response = &CreateCloudNativeAPIGatewaySecretKeyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateCloudNativeAPIGatewaySecretKey
// 创建消费者密钥。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  FAILEDOPERATION_VPC = "FailedOperation.Vpc"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  INVALIDPARAMETERVALUE_SPECIFICATION = "InvalidParameterValue.Specification"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_CAMPASSROLENOTEXIST = "UnauthorizedOperation.CamPassRoleNotExist"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) CreateCloudNativeAPIGatewaySecretKey(request *CreateCloudNativeAPIGatewaySecretKeyRequest) (response *CreateCloudNativeAPIGatewaySecretKeyResponse, err error) {
    return c.CreateCloudNativeAPIGatewaySecretKeyWithContext(context.Background(), request)
}

// CreateCloudNativeAPIGatewaySecretKey
// 创建消费者密钥。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  FAILEDOPERATION_VPC = "FailedOperation.Vpc"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  INVALIDPARAMETERVALUE_SPECIFICATION = "InvalidParameterValue.Specification"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_CAMPASSROLENOTEXIST = "UnauthorizedOperation.CamPassRoleNotExist"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) CreateCloudNativeAPIGatewaySecretKeyWithContext(ctx context.Context, request *CreateCloudNativeAPIGatewaySecretKeyRequest) (response *CreateCloudNativeAPIGatewaySecretKeyResponse, err error) {
    if request == nil {
        request = NewCreateCloudNativeAPIGatewaySecretKeyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cngw", APIVersion, "CreateCloudNativeAPIGatewaySecretKey")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCloudNativeAPIGatewaySecretKey require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateCloudNativeAPIGatewaySecretKeyResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteCloudNativeAPIGatewayConsumerRequest() (request *DeleteCloudNativeAPIGatewayConsumerRequest) {
    request = &DeleteCloudNativeAPIGatewayConsumerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cngw", APIVersion, "DeleteCloudNativeAPIGatewayConsumer")
    
    
    return
}

func NewDeleteCloudNativeAPIGatewayConsumerResponse() (response *DeleteCloudNativeAPIGatewayConsumerResponse) {
    response = &DeleteCloudNativeAPIGatewayConsumerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteCloudNativeAPIGatewayConsumer
// 删除AI 网关消费者（被绑定到消费者组/密钥时需先解绑）。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  FAILEDOPERATION_VPC = "FailedOperation.Vpc"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  INVALIDPARAMETERVALUE_SPECIFICATION = "InvalidParameterValue.Specification"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_CAMPASSROLENOTEXIST = "UnauthorizedOperation.CamPassRoleNotExist"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DeleteCloudNativeAPIGatewayConsumer(request *DeleteCloudNativeAPIGatewayConsumerRequest) (response *DeleteCloudNativeAPIGatewayConsumerResponse, err error) {
    return c.DeleteCloudNativeAPIGatewayConsumerWithContext(context.Background(), request)
}

// DeleteCloudNativeAPIGatewayConsumer
// 删除AI 网关消费者（被绑定到消费者组/密钥时需先解绑）。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  FAILEDOPERATION_VPC = "FailedOperation.Vpc"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  INVALIDPARAMETERVALUE_SPECIFICATION = "InvalidParameterValue.Specification"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_CAMPASSROLENOTEXIST = "UnauthorizedOperation.CamPassRoleNotExist"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DeleteCloudNativeAPIGatewayConsumerWithContext(ctx context.Context, request *DeleteCloudNativeAPIGatewayConsumerRequest) (response *DeleteCloudNativeAPIGatewayConsumerResponse, err error) {
    if request == nil {
        request = NewDeleteCloudNativeAPIGatewayConsumerRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cngw", APIVersion, "DeleteCloudNativeAPIGatewayConsumer")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteCloudNativeAPIGatewayConsumer require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteCloudNativeAPIGatewayConsumerResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteCloudNativeAPIGatewayConsumerGroupRequest() (request *DeleteCloudNativeAPIGatewayConsumerGroupRequest) {
    request = &DeleteCloudNativeAPIGatewayConsumerGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cngw", APIVersion, "DeleteCloudNativeAPIGatewayConsumerGroup")
    
    
    return
}

func NewDeleteCloudNativeAPIGatewayConsumerGroupResponse() (response *DeleteCloudNativeAPIGatewayConsumerGroupResponse) {
    response = &DeleteCloudNativeAPIGatewayConsumerGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteCloudNativeAPIGatewayConsumerGroup
// 删除AI网关消费者组
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  FAILEDOPERATION_VPC = "FailedOperation.Vpc"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  INVALIDPARAMETERVALUE_SPECIFICATION = "InvalidParameterValue.Specification"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_CAMPASSROLENOTEXIST = "UnauthorizedOperation.CamPassRoleNotExist"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DeleteCloudNativeAPIGatewayConsumerGroup(request *DeleteCloudNativeAPIGatewayConsumerGroupRequest) (response *DeleteCloudNativeAPIGatewayConsumerGroupResponse, err error) {
    return c.DeleteCloudNativeAPIGatewayConsumerGroupWithContext(context.Background(), request)
}

// DeleteCloudNativeAPIGatewayConsumerGroup
// 删除AI网关消费者组
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  FAILEDOPERATION_VPC = "FailedOperation.Vpc"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  INVALIDPARAMETERVALUE_SPECIFICATION = "InvalidParameterValue.Specification"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_CAMPASSROLENOTEXIST = "UnauthorizedOperation.CamPassRoleNotExist"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DeleteCloudNativeAPIGatewayConsumerGroupWithContext(ctx context.Context, request *DeleteCloudNativeAPIGatewayConsumerGroupRequest) (response *DeleteCloudNativeAPIGatewayConsumerGroupResponse, err error) {
    if request == nil {
        request = NewDeleteCloudNativeAPIGatewayConsumerGroupRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cngw", APIVersion, "DeleteCloudNativeAPIGatewayConsumerGroup")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteCloudNativeAPIGatewayConsumerGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteCloudNativeAPIGatewayConsumerGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteCloudNativeAPIGatewayLLMModelAPIRequest() (request *DeleteCloudNativeAPIGatewayLLMModelAPIRequest) {
    request = &DeleteCloudNativeAPIGatewayLLMModelAPIRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cngw", APIVersion, "DeleteCloudNativeAPIGatewayLLMModelAPI")
    
    
    return
}

func NewDeleteCloudNativeAPIGatewayLLMModelAPIResponse() (response *DeleteCloudNativeAPIGatewayLLMModelAPIResponse) {
    response = &DeleteCloudNativeAPIGatewayLLMModelAPIResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteCloudNativeAPIGatewayLLMModelAPI
// 删除 LLM 模型 API。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  LIMITEXCEEDED_LIMITEXCEEDED = "LimitExceeded.LimitExceeded"
//  MISSINGPARAMETER_MISSPARAMETER = "MissingParameter.MissParameter"
//  OPERATIONDENIED_OPERATIONDENIED = "OperationDenied.OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DeleteCloudNativeAPIGatewayLLMModelAPI(request *DeleteCloudNativeAPIGatewayLLMModelAPIRequest) (response *DeleteCloudNativeAPIGatewayLLMModelAPIResponse, err error) {
    return c.DeleteCloudNativeAPIGatewayLLMModelAPIWithContext(context.Background(), request)
}

// DeleteCloudNativeAPIGatewayLLMModelAPI
// 删除 LLM 模型 API。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  LIMITEXCEEDED_LIMITEXCEEDED = "LimitExceeded.LimitExceeded"
//  MISSINGPARAMETER_MISSPARAMETER = "MissingParameter.MissParameter"
//  OPERATIONDENIED_OPERATIONDENIED = "OperationDenied.OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DeleteCloudNativeAPIGatewayLLMModelAPIWithContext(ctx context.Context, request *DeleteCloudNativeAPIGatewayLLMModelAPIRequest) (response *DeleteCloudNativeAPIGatewayLLMModelAPIResponse, err error) {
    if request == nil {
        request = NewDeleteCloudNativeAPIGatewayLLMModelAPIRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cngw", APIVersion, "DeleteCloudNativeAPIGatewayLLMModelAPI")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteCloudNativeAPIGatewayLLMModelAPI require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteCloudNativeAPIGatewayLLMModelAPIResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteCloudNativeAPIGatewayLLMModelServiceRequest() (request *DeleteCloudNativeAPIGatewayLLMModelServiceRequest) {
    request = &DeleteCloudNativeAPIGatewayLLMModelServiceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cngw", APIVersion, "DeleteCloudNativeAPIGatewayLLMModelService")
    
    
    return
}

func NewDeleteCloudNativeAPIGatewayLLMModelServiceResponse() (response *DeleteCloudNativeAPIGatewayLLMModelServiceResponse) {
    response = &DeleteCloudNativeAPIGatewayLLMModelServiceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteCloudNativeAPIGatewayLLMModelService
// 删除 LLM 模型服务（被模型 API 绑定时需先解绑）。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  LIMITEXCEEDED_LIMITEXCEEDED = "LimitExceeded.LimitExceeded"
//  MISSINGPARAMETER_MISSPARAMETER = "MissingParameter.MissParameter"
//  OPERATIONDENIED_OPERATIONDENIED = "OperationDenied.OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DeleteCloudNativeAPIGatewayLLMModelService(request *DeleteCloudNativeAPIGatewayLLMModelServiceRequest) (response *DeleteCloudNativeAPIGatewayLLMModelServiceResponse, err error) {
    return c.DeleteCloudNativeAPIGatewayLLMModelServiceWithContext(context.Background(), request)
}

// DeleteCloudNativeAPIGatewayLLMModelService
// 删除 LLM 模型服务（被模型 API 绑定时需先解绑）。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  LIMITEXCEEDED_LIMITEXCEEDED = "LimitExceeded.LimitExceeded"
//  MISSINGPARAMETER_MISSPARAMETER = "MissingParameter.MissParameter"
//  OPERATIONDENIED_OPERATIONDENIED = "OperationDenied.OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DeleteCloudNativeAPIGatewayLLMModelServiceWithContext(ctx context.Context, request *DeleteCloudNativeAPIGatewayLLMModelServiceRequest) (response *DeleteCloudNativeAPIGatewayLLMModelServiceResponse, err error) {
    if request == nil {
        request = NewDeleteCloudNativeAPIGatewayLLMModelServiceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cngw", APIVersion, "DeleteCloudNativeAPIGatewayLLMModelService")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteCloudNativeAPIGatewayLLMModelService require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteCloudNativeAPIGatewayLLMModelServiceResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteCloudNativeAPIGatewayMCPServerRequest() (request *DeleteCloudNativeAPIGatewayMCPServerRequest) {
    request = &DeleteCloudNativeAPIGatewayMCPServerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cngw", APIVersion, "DeleteCloudNativeAPIGatewayMCPServer")
    
    
    return
}

func NewDeleteCloudNativeAPIGatewayMCPServerResponse() (response *DeleteCloudNativeAPIGatewayMCPServerResponse) {
    response = &DeleteCloudNativeAPIGatewayMCPServerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteCloudNativeAPIGatewayMCPServer
// 删除AI网关MCP服务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  FAILEDOPERATION_VPC = "FailedOperation.Vpc"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  INVALIDPARAMETERVALUE_SPECIFICATION = "InvalidParameterValue.Specification"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_CAMPASSROLENOTEXIST = "UnauthorizedOperation.CamPassRoleNotExist"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DeleteCloudNativeAPIGatewayMCPServer(request *DeleteCloudNativeAPIGatewayMCPServerRequest) (response *DeleteCloudNativeAPIGatewayMCPServerResponse, err error) {
    return c.DeleteCloudNativeAPIGatewayMCPServerWithContext(context.Background(), request)
}

// DeleteCloudNativeAPIGatewayMCPServer
// 删除AI网关MCP服务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  FAILEDOPERATION_VPC = "FailedOperation.Vpc"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  INVALIDPARAMETERVALUE_SPECIFICATION = "InvalidParameterValue.Specification"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_CAMPASSROLENOTEXIST = "UnauthorizedOperation.CamPassRoleNotExist"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DeleteCloudNativeAPIGatewayMCPServerWithContext(ctx context.Context, request *DeleteCloudNativeAPIGatewayMCPServerRequest) (response *DeleteCloudNativeAPIGatewayMCPServerResponse, err error) {
    if request == nil {
        request = NewDeleteCloudNativeAPIGatewayMCPServerRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cngw", APIVersion, "DeleteCloudNativeAPIGatewayMCPServer")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteCloudNativeAPIGatewayMCPServer require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteCloudNativeAPIGatewayMCPServerResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteCloudNativeAPIGatewayMCPToolRequest() (request *DeleteCloudNativeAPIGatewayMCPToolRequest) {
    request = &DeleteCloudNativeAPIGatewayMCPToolRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cngw", APIVersion, "DeleteCloudNativeAPIGatewayMCPTool")
    
    
    return
}

func NewDeleteCloudNativeAPIGatewayMCPToolResponse() (response *DeleteCloudNativeAPIGatewayMCPToolResponse) {
    response = &DeleteCloudNativeAPIGatewayMCPToolResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteCloudNativeAPIGatewayMCPTool
// 删除AI网关MCP Tool
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  FAILEDOPERATION_VPC = "FailedOperation.Vpc"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  INVALIDPARAMETERVALUE_SPECIFICATION = "InvalidParameterValue.Specification"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_CAMPASSROLENOTEXIST = "UnauthorizedOperation.CamPassRoleNotExist"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DeleteCloudNativeAPIGatewayMCPTool(request *DeleteCloudNativeAPIGatewayMCPToolRequest) (response *DeleteCloudNativeAPIGatewayMCPToolResponse, err error) {
    return c.DeleteCloudNativeAPIGatewayMCPToolWithContext(context.Background(), request)
}

// DeleteCloudNativeAPIGatewayMCPTool
// 删除AI网关MCP Tool
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  FAILEDOPERATION_VPC = "FailedOperation.Vpc"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  INVALIDPARAMETERVALUE_SPECIFICATION = "InvalidParameterValue.Specification"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_CAMPASSROLENOTEXIST = "UnauthorizedOperation.CamPassRoleNotExist"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DeleteCloudNativeAPIGatewayMCPToolWithContext(ctx context.Context, request *DeleteCloudNativeAPIGatewayMCPToolRequest) (response *DeleteCloudNativeAPIGatewayMCPToolResponse, err error) {
    if request == nil {
        request = NewDeleteCloudNativeAPIGatewayMCPToolRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cngw", APIVersion, "DeleteCloudNativeAPIGatewayMCPTool")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteCloudNativeAPIGatewayMCPTool require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteCloudNativeAPIGatewayMCPToolResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteCloudNativeAPIGatewaySecretKeyRequest() (request *DeleteCloudNativeAPIGatewaySecretKeyRequest) {
    request = &DeleteCloudNativeAPIGatewaySecretKeyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cngw", APIVersion, "DeleteCloudNativeAPIGatewaySecretKey")
    
    
    return
}

func NewDeleteCloudNativeAPIGatewaySecretKeyResponse() (response *DeleteCloudNativeAPIGatewaySecretKeyResponse) {
    response = &DeleteCloudNativeAPIGatewaySecretKeyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteCloudNativeAPIGatewaySecretKey
// 删除消费者密钥（被绑定时需先解绑）。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  FAILEDOPERATION_VPC = "FailedOperation.Vpc"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  INVALIDPARAMETERVALUE_SPECIFICATION = "InvalidParameterValue.Specification"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_CAMPASSROLENOTEXIST = "UnauthorizedOperation.CamPassRoleNotExist"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DeleteCloudNativeAPIGatewaySecretKey(request *DeleteCloudNativeAPIGatewaySecretKeyRequest) (response *DeleteCloudNativeAPIGatewaySecretKeyResponse, err error) {
    return c.DeleteCloudNativeAPIGatewaySecretKeyWithContext(context.Background(), request)
}

// DeleteCloudNativeAPIGatewaySecretKey
// 删除消费者密钥（被绑定时需先解绑）。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  FAILEDOPERATION_VPC = "FailedOperation.Vpc"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  INVALIDPARAMETERVALUE_SPECIFICATION = "InvalidParameterValue.Specification"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_CAMPASSROLENOTEXIST = "UnauthorizedOperation.CamPassRoleNotExist"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DeleteCloudNativeAPIGatewaySecretKeyWithContext(ctx context.Context, request *DeleteCloudNativeAPIGatewaySecretKeyRequest) (response *DeleteCloudNativeAPIGatewaySecretKeyResponse, err error) {
    if request == nil {
        request = NewDeleteCloudNativeAPIGatewaySecretKeyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cngw", APIVersion, "DeleteCloudNativeAPIGatewaySecretKey")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteCloudNativeAPIGatewaySecretKey require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteCloudNativeAPIGatewaySecretKeyResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCloudNativeAPIGatewayConsumerRequest() (request *DescribeCloudNativeAPIGatewayConsumerRequest) {
    request = &DescribeCloudNativeAPIGatewayConsumerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cngw", APIVersion, "DescribeCloudNativeAPIGatewayConsumer")
    
    
    return
}

func NewDescribeCloudNativeAPIGatewayConsumerResponse() (response *DescribeCloudNativeAPIGatewayConsumerResponse) {
    response = &DescribeCloudNativeAPIGatewayConsumerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCloudNativeAPIGatewayConsumer
// 查询云原生消费者详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  FAILEDOPERATION_VPC = "FailedOperation.Vpc"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  INVALIDPARAMETERVALUE_SPECIFICATION = "InvalidParameterValue.Specification"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_CAMPASSROLENOTEXIST = "UnauthorizedOperation.CamPassRoleNotExist"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DescribeCloudNativeAPIGatewayConsumer(request *DescribeCloudNativeAPIGatewayConsumerRequest) (response *DescribeCloudNativeAPIGatewayConsumerResponse, err error) {
    return c.DescribeCloudNativeAPIGatewayConsumerWithContext(context.Background(), request)
}

// DescribeCloudNativeAPIGatewayConsumer
// 查询云原生消费者详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  FAILEDOPERATION_VPC = "FailedOperation.Vpc"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  INVALIDPARAMETERVALUE_SPECIFICATION = "InvalidParameterValue.Specification"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_CAMPASSROLENOTEXIST = "UnauthorizedOperation.CamPassRoleNotExist"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DescribeCloudNativeAPIGatewayConsumerWithContext(ctx context.Context, request *DescribeCloudNativeAPIGatewayConsumerRequest) (response *DescribeCloudNativeAPIGatewayConsumerResponse, err error) {
    if request == nil {
        request = NewDescribeCloudNativeAPIGatewayConsumerRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cngw", APIVersion, "DescribeCloudNativeAPIGatewayConsumer")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCloudNativeAPIGatewayConsumer require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCloudNativeAPIGatewayConsumerResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCloudNativeAPIGatewayConsumerGroupRequest() (request *DescribeCloudNativeAPIGatewayConsumerGroupRequest) {
    request = &DescribeCloudNativeAPIGatewayConsumerGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cngw", APIVersion, "DescribeCloudNativeAPIGatewayConsumerGroup")
    
    
    return
}

func NewDescribeCloudNativeAPIGatewayConsumerGroupResponse() (response *DescribeCloudNativeAPIGatewayConsumerGroupResponse) {
    response = &DescribeCloudNativeAPIGatewayConsumerGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCloudNativeAPIGatewayConsumerGroup
// 查询消费者组详情。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  FAILEDOPERATION_VPC = "FailedOperation.Vpc"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  INVALIDPARAMETERVALUE_SPECIFICATION = "InvalidParameterValue.Specification"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_CAMPASSROLENOTEXIST = "UnauthorizedOperation.CamPassRoleNotExist"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DescribeCloudNativeAPIGatewayConsumerGroup(request *DescribeCloudNativeAPIGatewayConsumerGroupRequest) (response *DescribeCloudNativeAPIGatewayConsumerGroupResponse, err error) {
    return c.DescribeCloudNativeAPIGatewayConsumerGroupWithContext(context.Background(), request)
}

// DescribeCloudNativeAPIGatewayConsumerGroup
// 查询消费者组详情。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  FAILEDOPERATION_VPC = "FailedOperation.Vpc"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  INVALIDPARAMETERVALUE_SPECIFICATION = "InvalidParameterValue.Specification"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_CAMPASSROLENOTEXIST = "UnauthorizedOperation.CamPassRoleNotExist"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DescribeCloudNativeAPIGatewayConsumerGroupWithContext(ctx context.Context, request *DescribeCloudNativeAPIGatewayConsumerGroupRequest) (response *DescribeCloudNativeAPIGatewayConsumerGroupResponse, err error) {
    if request == nil {
        request = NewDescribeCloudNativeAPIGatewayConsumerGroupRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cngw", APIVersion, "DescribeCloudNativeAPIGatewayConsumerGroup")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCloudNativeAPIGatewayConsumerGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCloudNativeAPIGatewayConsumerGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCloudNativeAPIGatewayLLMModelAPIRequest() (request *DescribeCloudNativeAPIGatewayLLMModelAPIRequest) {
    request = &DescribeCloudNativeAPIGatewayLLMModelAPIRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cngw", APIVersion, "DescribeCloudNativeAPIGatewayLLMModelAPI")
    
    
    return
}

func NewDescribeCloudNativeAPIGatewayLLMModelAPIResponse() (response *DescribeCloudNativeAPIGatewayLLMModelAPIResponse) {
    response = &DescribeCloudNativeAPIGatewayLLMModelAPIResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCloudNativeAPIGatewayLLMModelAPI
// 查询单个 LLM 模型 API 详情。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  LIMITEXCEEDED_LIMITEXCEEDED = "LimitExceeded.LimitExceeded"
//  MISSINGPARAMETER_MISSPARAMETER = "MissingParameter.MissParameter"
//  OPERATIONDENIED_OPERATIONDENIED = "OperationDenied.OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DescribeCloudNativeAPIGatewayLLMModelAPI(request *DescribeCloudNativeAPIGatewayLLMModelAPIRequest) (response *DescribeCloudNativeAPIGatewayLLMModelAPIResponse, err error) {
    return c.DescribeCloudNativeAPIGatewayLLMModelAPIWithContext(context.Background(), request)
}

// DescribeCloudNativeAPIGatewayLLMModelAPI
// 查询单个 LLM 模型 API 详情。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  LIMITEXCEEDED_LIMITEXCEEDED = "LimitExceeded.LimitExceeded"
//  MISSINGPARAMETER_MISSPARAMETER = "MissingParameter.MissParameter"
//  OPERATIONDENIED_OPERATIONDENIED = "OperationDenied.OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DescribeCloudNativeAPIGatewayLLMModelAPIWithContext(ctx context.Context, request *DescribeCloudNativeAPIGatewayLLMModelAPIRequest) (response *DescribeCloudNativeAPIGatewayLLMModelAPIResponse, err error) {
    if request == nil {
        request = NewDescribeCloudNativeAPIGatewayLLMModelAPIRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cngw", APIVersion, "DescribeCloudNativeAPIGatewayLLMModelAPI")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCloudNativeAPIGatewayLLMModelAPI require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCloudNativeAPIGatewayLLMModelAPIResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCloudNativeAPIGatewayLLMModelAPIsRequest() (request *DescribeCloudNativeAPIGatewayLLMModelAPIsRequest) {
    request = &DescribeCloudNativeAPIGatewayLLMModelAPIsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cngw", APIVersion, "DescribeCloudNativeAPIGatewayLLMModelAPIs")
    
    
    return
}

func NewDescribeCloudNativeAPIGatewayLLMModelAPIsResponse() (response *DescribeCloudNativeAPIGatewayLLMModelAPIsResponse) {
    response = &DescribeCloudNativeAPIGatewayLLMModelAPIsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCloudNativeAPIGatewayLLMModelAPIs
// 查询 LLM 模型 API 列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  LIMITEXCEEDED_LIMITEXCEEDED = "LimitExceeded.LimitExceeded"
//  MISSINGPARAMETER_MISSPARAMETER = "MissingParameter.MissParameter"
//  OPERATIONDENIED_OPERATIONDENIED = "OperationDenied.OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DescribeCloudNativeAPIGatewayLLMModelAPIs(request *DescribeCloudNativeAPIGatewayLLMModelAPIsRequest) (response *DescribeCloudNativeAPIGatewayLLMModelAPIsResponse, err error) {
    return c.DescribeCloudNativeAPIGatewayLLMModelAPIsWithContext(context.Background(), request)
}

// DescribeCloudNativeAPIGatewayLLMModelAPIs
// 查询 LLM 模型 API 列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  LIMITEXCEEDED_LIMITEXCEEDED = "LimitExceeded.LimitExceeded"
//  MISSINGPARAMETER_MISSPARAMETER = "MissingParameter.MissParameter"
//  OPERATIONDENIED_OPERATIONDENIED = "OperationDenied.OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DescribeCloudNativeAPIGatewayLLMModelAPIsWithContext(ctx context.Context, request *DescribeCloudNativeAPIGatewayLLMModelAPIsRequest) (response *DescribeCloudNativeAPIGatewayLLMModelAPIsResponse, err error) {
    if request == nil {
        request = NewDescribeCloudNativeAPIGatewayLLMModelAPIsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cngw", APIVersion, "DescribeCloudNativeAPIGatewayLLMModelAPIs")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCloudNativeAPIGatewayLLMModelAPIs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCloudNativeAPIGatewayLLMModelAPIsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCloudNativeAPIGatewayLLMModelServiceRequest() (request *DescribeCloudNativeAPIGatewayLLMModelServiceRequest) {
    request = &DescribeCloudNativeAPIGatewayLLMModelServiceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cngw", APIVersion, "DescribeCloudNativeAPIGatewayLLMModelService")
    
    
    return
}

func NewDescribeCloudNativeAPIGatewayLLMModelServiceResponse() (response *DescribeCloudNativeAPIGatewayLLMModelServiceResponse) {
    response = &DescribeCloudNativeAPIGatewayLLMModelServiceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCloudNativeAPIGatewayLLMModelService
// 查询单个 LLM 模型服务详情。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  LIMITEXCEEDED_LIMITEXCEEDED = "LimitExceeded.LimitExceeded"
//  MISSINGPARAMETER_MISSPARAMETER = "MissingParameter.MissParameter"
//  OPERATIONDENIED_OPERATIONDENIED = "OperationDenied.OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DescribeCloudNativeAPIGatewayLLMModelService(request *DescribeCloudNativeAPIGatewayLLMModelServiceRequest) (response *DescribeCloudNativeAPIGatewayLLMModelServiceResponse, err error) {
    return c.DescribeCloudNativeAPIGatewayLLMModelServiceWithContext(context.Background(), request)
}

// DescribeCloudNativeAPIGatewayLLMModelService
// 查询单个 LLM 模型服务详情。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  LIMITEXCEEDED_LIMITEXCEEDED = "LimitExceeded.LimitExceeded"
//  MISSINGPARAMETER_MISSPARAMETER = "MissingParameter.MissParameter"
//  OPERATIONDENIED_OPERATIONDENIED = "OperationDenied.OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DescribeCloudNativeAPIGatewayLLMModelServiceWithContext(ctx context.Context, request *DescribeCloudNativeAPIGatewayLLMModelServiceRequest) (response *DescribeCloudNativeAPIGatewayLLMModelServiceResponse, err error) {
    if request == nil {
        request = NewDescribeCloudNativeAPIGatewayLLMModelServiceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cngw", APIVersion, "DescribeCloudNativeAPIGatewayLLMModelService")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCloudNativeAPIGatewayLLMModelService require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCloudNativeAPIGatewayLLMModelServiceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCloudNativeAPIGatewayLLMModelServicesRequest() (request *DescribeCloudNativeAPIGatewayLLMModelServicesRequest) {
    request = &DescribeCloudNativeAPIGatewayLLMModelServicesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cngw", APIVersion, "DescribeCloudNativeAPIGatewayLLMModelServices")
    
    
    return
}

func NewDescribeCloudNativeAPIGatewayLLMModelServicesResponse() (response *DescribeCloudNativeAPIGatewayLLMModelServicesResponse) {
    response = &DescribeCloudNativeAPIGatewayLLMModelServicesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCloudNativeAPIGatewayLLMModelServices
// 查询 LLM 模型服务列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  LIMITEXCEEDED_LIMITEXCEEDED = "LimitExceeded.LimitExceeded"
//  MISSINGPARAMETER_MISSPARAMETER = "MissingParameter.MissParameter"
//  OPERATIONDENIED_OPERATIONDENIED = "OperationDenied.OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DescribeCloudNativeAPIGatewayLLMModelServices(request *DescribeCloudNativeAPIGatewayLLMModelServicesRequest) (response *DescribeCloudNativeAPIGatewayLLMModelServicesResponse, err error) {
    return c.DescribeCloudNativeAPIGatewayLLMModelServicesWithContext(context.Background(), request)
}

// DescribeCloudNativeAPIGatewayLLMModelServices
// 查询 LLM 模型服务列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  LIMITEXCEEDED_LIMITEXCEEDED = "LimitExceeded.LimitExceeded"
//  MISSINGPARAMETER_MISSPARAMETER = "MissingParameter.MissParameter"
//  OPERATIONDENIED_OPERATIONDENIED = "OperationDenied.OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DescribeCloudNativeAPIGatewayLLMModelServicesWithContext(ctx context.Context, request *DescribeCloudNativeAPIGatewayLLMModelServicesRequest) (response *DescribeCloudNativeAPIGatewayLLMModelServicesResponse, err error) {
    if request == nil {
        request = NewDescribeCloudNativeAPIGatewayLLMModelServicesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cngw", APIVersion, "DescribeCloudNativeAPIGatewayLLMModelServices")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCloudNativeAPIGatewayLLMModelServices require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCloudNativeAPIGatewayLLMModelServicesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCloudNativeAPIGatewayLLMTokenUsageListRequest() (request *DescribeCloudNativeAPIGatewayLLMTokenUsageListRequest) {
    request = &DescribeCloudNativeAPIGatewayLLMTokenUsageListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cngw", APIVersion, "DescribeCloudNativeAPIGatewayLLMTokenUsageList")
    
    
    return
}

func NewDescribeCloudNativeAPIGatewayLLMTokenUsageListResponse() (response *DescribeCloudNativeAPIGatewayLLMTokenUsageListResponse) {
    response = &DescribeCloudNativeAPIGatewayLLMTokenUsageListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCloudNativeAPIGatewayLLMTokenUsageList
// 查询 AI 网关Token 消耗统计
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  LIMITEXCEEDED_LIMITEXCEEDED = "LimitExceeded.LimitExceeded"
//  MISSINGPARAMETER_MISSPARAMETER = "MissingParameter.MissParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DescribeCloudNativeAPIGatewayLLMTokenUsageList(request *DescribeCloudNativeAPIGatewayLLMTokenUsageListRequest) (response *DescribeCloudNativeAPIGatewayLLMTokenUsageListResponse, err error) {
    return c.DescribeCloudNativeAPIGatewayLLMTokenUsageListWithContext(context.Background(), request)
}

// DescribeCloudNativeAPIGatewayLLMTokenUsageList
// 查询 AI 网关Token 消耗统计
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  LIMITEXCEEDED_LIMITEXCEEDED = "LimitExceeded.LimitExceeded"
//  MISSINGPARAMETER_MISSPARAMETER = "MissingParameter.MissParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DescribeCloudNativeAPIGatewayLLMTokenUsageListWithContext(ctx context.Context, request *DescribeCloudNativeAPIGatewayLLMTokenUsageListRequest) (response *DescribeCloudNativeAPIGatewayLLMTokenUsageListResponse, err error) {
    if request == nil {
        request = NewDescribeCloudNativeAPIGatewayLLMTokenUsageListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cngw", APIVersion, "DescribeCloudNativeAPIGatewayLLMTokenUsageList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCloudNativeAPIGatewayLLMTokenUsageList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCloudNativeAPIGatewayLLMTokenUsageListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCloudNativeAPIGatewayLLMTokenUsageStatisticsRequest() (request *DescribeCloudNativeAPIGatewayLLMTokenUsageStatisticsRequest) {
    request = &DescribeCloudNativeAPIGatewayLLMTokenUsageStatisticsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cngw", APIVersion, "DescribeCloudNativeAPIGatewayLLMTokenUsageStatistics")
    
    
    return
}

func NewDescribeCloudNativeAPIGatewayLLMTokenUsageStatisticsResponse() (response *DescribeCloudNativeAPIGatewayLLMTokenUsageStatisticsResponse) {
    response = &DescribeCloudNativeAPIGatewayLLMTokenUsageStatisticsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCloudNativeAPIGatewayLLMTokenUsageStatistics
// 查询 AI 网关Token 消耗统计汇总
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  LIMITEXCEEDED_LIMITEXCEEDED = "LimitExceeded.LimitExceeded"
//  MISSINGPARAMETER_MISSPARAMETER = "MissingParameter.MissParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DescribeCloudNativeAPIGatewayLLMTokenUsageStatistics(request *DescribeCloudNativeAPIGatewayLLMTokenUsageStatisticsRequest) (response *DescribeCloudNativeAPIGatewayLLMTokenUsageStatisticsResponse, err error) {
    return c.DescribeCloudNativeAPIGatewayLLMTokenUsageStatisticsWithContext(context.Background(), request)
}

// DescribeCloudNativeAPIGatewayLLMTokenUsageStatistics
// 查询 AI 网关Token 消耗统计汇总
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  LIMITEXCEEDED_LIMITEXCEEDED = "LimitExceeded.LimitExceeded"
//  MISSINGPARAMETER_MISSPARAMETER = "MissingParameter.MissParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DescribeCloudNativeAPIGatewayLLMTokenUsageStatisticsWithContext(ctx context.Context, request *DescribeCloudNativeAPIGatewayLLMTokenUsageStatisticsRequest) (response *DescribeCloudNativeAPIGatewayLLMTokenUsageStatisticsResponse, err error) {
    if request == nil {
        request = NewDescribeCloudNativeAPIGatewayLLMTokenUsageStatisticsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cngw", APIVersion, "DescribeCloudNativeAPIGatewayLLMTokenUsageStatistics")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCloudNativeAPIGatewayLLMTokenUsageStatistics require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCloudNativeAPIGatewayLLMTokenUsageStatisticsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCloudNativeAPIGatewayMCPServerRequest() (request *DescribeCloudNativeAPIGatewayMCPServerRequest) {
    request = &DescribeCloudNativeAPIGatewayMCPServerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cngw", APIVersion, "DescribeCloudNativeAPIGatewayMCPServer")
    
    
    return
}

func NewDescribeCloudNativeAPIGatewayMCPServerResponse() (response *DescribeCloudNativeAPIGatewayMCPServerResponse) {
    response = &DescribeCloudNativeAPIGatewayMCPServerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCloudNativeAPIGatewayMCPServer
// 查询AI 网关MCP服务信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  FAILEDOPERATION_VPC = "FailedOperation.Vpc"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  INVALIDPARAMETERVALUE_SPECIFICATION = "InvalidParameterValue.Specification"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_CAMPASSROLENOTEXIST = "UnauthorizedOperation.CamPassRoleNotExist"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DescribeCloudNativeAPIGatewayMCPServer(request *DescribeCloudNativeAPIGatewayMCPServerRequest) (response *DescribeCloudNativeAPIGatewayMCPServerResponse, err error) {
    return c.DescribeCloudNativeAPIGatewayMCPServerWithContext(context.Background(), request)
}

// DescribeCloudNativeAPIGatewayMCPServer
// 查询AI 网关MCP服务信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  FAILEDOPERATION_VPC = "FailedOperation.Vpc"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  INVALIDPARAMETERVALUE_SPECIFICATION = "InvalidParameterValue.Specification"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_CAMPASSROLENOTEXIST = "UnauthorizedOperation.CamPassRoleNotExist"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DescribeCloudNativeAPIGatewayMCPServerWithContext(ctx context.Context, request *DescribeCloudNativeAPIGatewayMCPServerRequest) (response *DescribeCloudNativeAPIGatewayMCPServerResponse, err error) {
    if request == nil {
        request = NewDescribeCloudNativeAPIGatewayMCPServerRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cngw", APIVersion, "DescribeCloudNativeAPIGatewayMCPServer")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCloudNativeAPIGatewayMCPServer require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCloudNativeAPIGatewayMCPServerResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCloudNativeAPIGatewayMCPServerACLRequest() (request *DescribeCloudNativeAPIGatewayMCPServerACLRequest) {
    request = &DescribeCloudNativeAPIGatewayMCPServerACLRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cngw", APIVersion, "DescribeCloudNativeAPIGatewayMCPServerACL")
    
    
    return
}

func NewDescribeCloudNativeAPIGatewayMCPServerACLResponse() (response *DescribeCloudNativeAPIGatewayMCPServerACLResponse) {
    response = &DescribeCloudNativeAPIGatewayMCPServerACLResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCloudNativeAPIGatewayMCPServerACL
// 查看 MCP Server ACL
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SECRETIDNOTFOUND = "AuthFailure.SecretIdNotFound"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLS = "FailedOperation.Cls"
//  FAILEDOPERATION_CLSCONFIG = "FailedOperation.ClsConfig"
//  FAILEDOPERATION_COS = "FailedOperation.Cos"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_PLUGIN = "FailedOperation.Plugin"
//  FAILEDOPERATION_RESOURCE = "FailedOperation.Resource"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  FAILEDOPERATION_VPC = "FailedOperation.Vpc"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CLBFAILURE = "InternalError.CLBFailure"
//  INTERNALERROR_CAMNOAUTH = "InternalError.CamNoAuth"
//  INTERNALERROR_CREATEERROR = "InternalError.CreateError"
//  INTERNALERROR_DECODEERROR = "InternalError.DecodeError"
//  INTERNALERROR_EKSFAILURE = "InternalError.EKSFailure"
//  INTERNALERROR_GETCREDENTIAL = "InternalError.GetCredential"
//  INTERNALERROR_GETROLEERROR = "InternalError.GetRoleError"
//  INTERNALERROR_GOVERNANCEFAILURE = "InternalError.GovernanceFailure"
//  INTERNALERROR_HTTPSTATUSCODEERROR = "InternalError.HttpStatusCodeError"
//  INTERNALERROR_IOERROR = "InternalError.IOError"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INTERNALERROR_OPERATIONFAILED = "InternalError.OperationFailed"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_TKECLUSTERNOTREADY = "InternalError.TKEClusterNotReady"
//  INTERNALERROR_TKEFAILURE = "InternalError.TKEFailure"
//  INTERNALERROR_TAGFAILURE = "InternalError.TagFailure"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INTERNALERROR_UPDATEERROR = "InternalError.UpdateError"
//  INTERNALERROR_VPCFAILURE = "InternalError.VPCFailure"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARSEJSONERROR = "InvalidParameter.ParseJsonError"
//  INVALIDPARAMETERCOMBINATION = "InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ACTION = "InvalidParameterValue.Action"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_BUCKETNAME = "InvalidParameterValue.BucketName"
//  INVALIDPARAMETERVALUE_COS = "InvalidParameterValue.Cos"
//  INVALIDPARAMETERVALUE_CREATEERROR = "InvalidParameterValue.CreateError"
//  INVALIDPARAMETERVALUE_DESCRIPTION = "InvalidParameterValue.Description"
//  INVALIDPARAMETERVALUE_GATEWAYID = "InvalidParameterValue.GatewayId"
//  INVALIDPARAMETERVALUE_GATEWAYVERSION = "InvalidParameterValue.GatewayVersion"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_KEYNOTEXIST = "InvalidParameterValue.KeyNotExist"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  INVALIDPARAMETERVALUE_OBJECTNAME = "InvalidParameterValue.ObjectName"
//  INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"
//  INVALIDPARAMETERVALUE_PLUGIN = "InvalidParameterValue.Plugin"
//  INVALIDPARAMETERVALUE_QUERYERROR = "InvalidParameterValue.QueryError"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  INVALIDPARAMETERVALUE_SPECIFICATION = "InvalidParameterValue.Specification"
//  INVALIDPARAMETERVALUE_SYSTEMPARAMETER = "InvalidParameterValue.SystemParameter"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  INVALIDPARAMETERVALUE_UPDATEERROR = "InvalidParameterValue.UpdateError"
//  INVALIDPARAMETERVALUE_ZIPFILE = "InvalidParameterValue.ZipFile"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_LBDOMAINS = "LimitExceeded.LBDomains"
//  LIMITEXCEEDED_LIMITEXCEEDED = "LimitExceeded.LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_CREATEERROR = "MissingParameter.CreateError"
//  MISSINGPARAMETER_MISSPARAMETER = "MissingParameter.MissParameter"
//  MISSINGPARAMETER_UPDATEERROR = "MissingParameter.UpdateError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_LIMITEXCEEDED = "OperationDenied.LimitExceeded"
//  OPERATIONDENIED_OPERATIONDENIED = "OperationDenied.OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINUSE_GATEWAYSERVICEEXISTGROUPS = "ResourceInUse.GatewayServiceExistGroups"
//  RESOURCEINUSE_GATEWAYSERVICEEXISTROUTER = "ResourceInUse.GatewayServiceExistRouter"
//  RESOURCEINUSE_GATEWAYSERVICEGROUPEXISTUPSTREAM = "ResourceInUse.GatewayServiceGroupExistUpstream"
//  RESOURCEINUSE_GATEWAYSERVICESOURCEEXISTSERVICE = "ResourceInUse.GatewayServiceSourceExistService"
//  RESOURCEINUSE_INSTANCESEXISTEDINSERVICE = "ResourceInUse.InstancesExistedInService"
//  RESOURCEINUSE_SERVICESEXISTEDINNAMESPACE = "ResourceInUse.ServicesExistedInNamespace"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_CONFIGNOTEXIST = "ResourceNotFound.ConfigNotExist"
//  RESOURCENOTFOUND_FORBIDDEN = "ResourceNotFound.Forbidden"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_CAMPASSROLENOTEXIST = "UnauthorizedOperation.CamPassRoleNotExist"
//  UNAUTHORIZEDOPERATION_CLSNOTACTIVATED = "UnauthorizedOperation.ClsNotActivated"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_GOVERNANCENOTSUPPORTVISIBLE = "UnsupportedOperation.GovernanceNotSupportVisible"
func (c *Client) DescribeCloudNativeAPIGatewayMCPServerACL(request *DescribeCloudNativeAPIGatewayMCPServerACLRequest) (response *DescribeCloudNativeAPIGatewayMCPServerACLResponse, err error) {
    return c.DescribeCloudNativeAPIGatewayMCPServerACLWithContext(context.Background(), request)
}

// DescribeCloudNativeAPIGatewayMCPServerACL
// 查看 MCP Server ACL
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SECRETIDNOTFOUND = "AuthFailure.SecretIdNotFound"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLS = "FailedOperation.Cls"
//  FAILEDOPERATION_CLSCONFIG = "FailedOperation.ClsConfig"
//  FAILEDOPERATION_COS = "FailedOperation.Cos"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_PLUGIN = "FailedOperation.Plugin"
//  FAILEDOPERATION_RESOURCE = "FailedOperation.Resource"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  FAILEDOPERATION_VPC = "FailedOperation.Vpc"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CLBFAILURE = "InternalError.CLBFailure"
//  INTERNALERROR_CAMNOAUTH = "InternalError.CamNoAuth"
//  INTERNALERROR_CREATEERROR = "InternalError.CreateError"
//  INTERNALERROR_DECODEERROR = "InternalError.DecodeError"
//  INTERNALERROR_EKSFAILURE = "InternalError.EKSFailure"
//  INTERNALERROR_GETCREDENTIAL = "InternalError.GetCredential"
//  INTERNALERROR_GETROLEERROR = "InternalError.GetRoleError"
//  INTERNALERROR_GOVERNANCEFAILURE = "InternalError.GovernanceFailure"
//  INTERNALERROR_HTTPSTATUSCODEERROR = "InternalError.HttpStatusCodeError"
//  INTERNALERROR_IOERROR = "InternalError.IOError"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INTERNALERROR_OPERATIONFAILED = "InternalError.OperationFailed"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_TKECLUSTERNOTREADY = "InternalError.TKEClusterNotReady"
//  INTERNALERROR_TKEFAILURE = "InternalError.TKEFailure"
//  INTERNALERROR_TAGFAILURE = "InternalError.TagFailure"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INTERNALERROR_UPDATEERROR = "InternalError.UpdateError"
//  INTERNALERROR_VPCFAILURE = "InternalError.VPCFailure"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARSEJSONERROR = "InvalidParameter.ParseJsonError"
//  INVALIDPARAMETERCOMBINATION = "InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ACTION = "InvalidParameterValue.Action"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_BUCKETNAME = "InvalidParameterValue.BucketName"
//  INVALIDPARAMETERVALUE_COS = "InvalidParameterValue.Cos"
//  INVALIDPARAMETERVALUE_CREATEERROR = "InvalidParameterValue.CreateError"
//  INVALIDPARAMETERVALUE_DESCRIPTION = "InvalidParameterValue.Description"
//  INVALIDPARAMETERVALUE_GATEWAYID = "InvalidParameterValue.GatewayId"
//  INVALIDPARAMETERVALUE_GATEWAYVERSION = "InvalidParameterValue.GatewayVersion"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_KEYNOTEXIST = "InvalidParameterValue.KeyNotExist"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  INVALIDPARAMETERVALUE_OBJECTNAME = "InvalidParameterValue.ObjectName"
//  INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"
//  INVALIDPARAMETERVALUE_PLUGIN = "InvalidParameterValue.Plugin"
//  INVALIDPARAMETERVALUE_QUERYERROR = "InvalidParameterValue.QueryError"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  INVALIDPARAMETERVALUE_SPECIFICATION = "InvalidParameterValue.Specification"
//  INVALIDPARAMETERVALUE_SYSTEMPARAMETER = "InvalidParameterValue.SystemParameter"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  INVALIDPARAMETERVALUE_UPDATEERROR = "InvalidParameterValue.UpdateError"
//  INVALIDPARAMETERVALUE_ZIPFILE = "InvalidParameterValue.ZipFile"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_LBDOMAINS = "LimitExceeded.LBDomains"
//  LIMITEXCEEDED_LIMITEXCEEDED = "LimitExceeded.LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_CREATEERROR = "MissingParameter.CreateError"
//  MISSINGPARAMETER_MISSPARAMETER = "MissingParameter.MissParameter"
//  MISSINGPARAMETER_UPDATEERROR = "MissingParameter.UpdateError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_LIMITEXCEEDED = "OperationDenied.LimitExceeded"
//  OPERATIONDENIED_OPERATIONDENIED = "OperationDenied.OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINUSE_GATEWAYSERVICEEXISTGROUPS = "ResourceInUse.GatewayServiceExistGroups"
//  RESOURCEINUSE_GATEWAYSERVICEEXISTROUTER = "ResourceInUse.GatewayServiceExistRouter"
//  RESOURCEINUSE_GATEWAYSERVICEGROUPEXISTUPSTREAM = "ResourceInUse.GatewayServiceGroupExistUpstream"
//  RESOURCEINUSE_GATEWAYSERVICESOURCEEXISTSERVICE = "ResourceInUse.GatewayServiceSourceExistService"
//  RESOURCEINUSE_INSTANCESEXISTEDINSERVICE = "ResourceInUse.InstancesExistedInService"
//  RESOURCEINUSE_SERVICESEXISTEDINNAMESPACE = "ResourceInUse.ServicesExistedInNamespace"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_CONFIGNOTEXIST = "ResourceNotFound.ConfigNotExist"
//  RESOURCENOTFOUND_FORBIDDEN = "ResourceNotFound.Forbidden"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_CAMPASSROLENOTEXIST = "UnauthorizedOperation.CamPassRoleNotExist"
//  UNAUTHORIZEDOPERATION_CLSNOTACTIVATED = "UnauthorizedOperation.ClsNotActivated"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_GOVERNANCENOTSUPPORTVISIBLE = "UnsupportedOperation.GovernanceNotSupportVisible"
func (c *Client) DescribeCloudNativeAPIGatewayMCPServerACLWithContext(ctx context.Context, request *DescribeCloudNativeAPIGatewayMCPServerACLRequest) (response *DescribeCloudNativeAPIGatewayMCPServerACLResponse, err error) {
    if request == nil {
        request = NewDescribeCloudNativeAPIGatewayMCPServerACLRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cngw", APIVersion, "DescribeCloudNativeAPIGatewayMCPServerACL")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCloudNativeAPIGatewayMCPServerACL require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCloudNativeAPIGatewayMCPServerACLResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCloudNativeAPIGatewayMCPServerAuthRequest() (request *DescribeCloudNativeAPIGatewayMCPServerAuthRequest) {
    request = &DescribeCloudNativeAPIGatewayMCPServerAuthRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cngw", APIVersion, "DescribeCloudNativeAPIGatewayMCPServerAuth")
    
    
    return
}

func NewDescribeCloudNativeAPIGatewayMCPServerAuthResponse() (response *DescribeCloudNativeAPIGatewayMCPServerAuthResponse) {
    response = &DescribeCloudNativeAPIGatewayMCPServerAuthResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCloudNativeAPIGatewayMCPServerAuth
// 查询 MCP Server 的认证配置
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SECRETIDNOTFOUND = "AuthFailure.SecretIdNotFound"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLS = "FailedOperation.Cls"
//  FAILEDOPERATION_CLSCONFIG = "FailedOperation.ClsConfig"
//  FAILEDOPERATION_COS = "FailedOperation.Cos"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_PLUGIN = "FailedOperation.Plugin"
//  FAILEDOPERATION_RESOURCE = "FailedOperation.Resource"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  FAILEDOPERATION_VPC = "FailedOperation.Vpc"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CLBFAILURE = "InternalError.CLBFailure"
//  INTERNALERROR_CAMNOAUTH = "InternalError.CamNoAuth"
//  INTERNALERROR_CREATEERROR = "InternalError.CreateError"
//  INTERNALERROR_DECODEERROR = "InternalError.DecodeError"
//  INTERNALERROR_EKSFAILURE = "InternalError.EKSFailure"
//  INTERNALERROR_GETCREDENTIAL = "InternalError.GetCredential"
//  INTERNALERROR_GETROLEERROR = "InternalError.GetRoleError"
//  INTERNALERROR_GOVERNANCEFAILURE = "InternalError.GovernanceFailure"
//  INTERNALERROR_HTTPSTATUSCODEERROR = "InternalError.HttpStatusCodeError"
//  INTERNALERROR_IOERROR = "InternalError.IOError"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INTERNALERROR_OPERATIONFAILED = "InternalError.OperationFailed"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_TKECLUSTERNOTREADY = "InternalError.TKEClusterNotReady"
//  INTERNALERROR_TKEFAILURE = "InternalError.TKEFailure"
//  INTERNALERROR_TAGFAILURE = "InternalError.TagFailure"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INTERNALERROR_UPDATEERROR = "InternalError.UpdateError"
//  INTERNALERROR_VPCFAILURE = "InternalError.VPCFailure"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARSEJSONERROR = "InvalidParameter.ParseJsonError"
//  INVALIDPARAMETERCOMBINATION = "InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ACTION = "InvalidParameterValue.Action"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_BUCKETNAME = "InvalidParameterValue.BucketName"
//  INVALIDPARAMETERVALUE_COS = "InvalidParameterValue.Cos"
//  INVALIDPARAMETERVALUE_CREATEERROR = "InvalidParameterValue.CreateError"
//  INVALIDPARAMETERVALUE_DESCRIPTION = "InvalidParameterValue.Description"
//  INVALIDPARAMETERVALUE_GATEWAYID = "InvalidParameterValue.GatewayId"
//  INVALIDPARAMETERVALUE_GATEWAYVERSION = "InvalidParameterValue.GatewayVersion"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_KEYNOTEXIST = "InvalidParameterValue.KeyNotExist"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  INVALIDPARAMETERVALUE_OBJECTNAME = "InvalidParameterValue.ObjectName"
//  INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"
//  INVALIDPARAMETERVALUE_PLUGIN = "InvalidParameterValue.Plugin"
//  INVALIDPARAMETERVALUE_QUERYERROR = "InvalidParameterValue.QueryError"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  INVALIDPARAMETERVALUE_SPECIFICATION = "InvalidParameterValue.Specification"
//  INVALIDPARAMETERVALUE_SYSTEMPARAMETER = "InvalidParameterValue.SystemParameter"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  INVALIDPARAMETERVALUE_UPDATEERROR = "InvalidParameterValue.UpdateError"
//  INVALIDPARAMETERVALUE_ZIPFILE = "InvalidParameterValue.ZipFile"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_LBDOMAINS = "LimitExceeded.LBDomains"
//  LIMITEXCEEDED_LIMITEXCEEDED = "LimitExceeded.LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_CREATEERROR = "MissingParameter.CreateError"
//  MISSINGPARAMETER_MISSPARAMETER = "MissingParameter.MissParameter"
//  MISSINGPARAMETER_UPDATEERROR = "MissingParameter.UpdateError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_LIMITEXCEEDED = "OperationDenied.LimitExceeded"
//  OPERATIONDENIED_OPERATIONDENIED = "OperationDenied.OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINUSE_GATEWAYSERVICEEXISTGROUPS = "ResourceInUse.GatewayServiceExistGroups"
//  RESOURCEINUSE_GATEWAYSERVICEEXISTROUTER = "ResourceInUse.GatewayServiceExistRouter"
//  RESOURCEINUSE_GATEWAYSERVICEGROUPEXISTUPSTREAM = "ResourceInUse.GatewayServiceGroupExistUpstream"
//  RESOURCEINUSE_GATEWAYSERVICESOURCEEXISTSERVICE = "ResourceInUse.GatewayServiceSourceExistService"
//  RESOURCEINUSE_INSTANCESEXISTEDINSERVICE = "ResourceInUse.InstancesExistedInService"
//  RESOURCEINUSE_SERVICESEXISTEDINNAMESPACE = "ResourceInUse.ServicesExistedInNamespace"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_CONFIGNOTEXIST = "ResourceNotFound.ConfigNotExist"
//  RESOURCENOTFOUND_FORBIDDEN = "ResourceNotFound.Forbidden"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_CAMPASSROLENOTEXIST = "UnauthorizedOperation.CamPassRoleNotExist"
//  UNAUTHORIZEDOPERATION_CLSNOTACTIVATED = "UnauthorizedOperation.ClsNotActivated"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_GOVERNANCENOTSUPPORTVISIBLE = "UnsupportedOperation.GovernanceNotSupportVisible"
func (c *Client) DescribeCloudNativeAPIGatewayMCPServerAuth(request *DescribeCloudNativeAPIGatewayMCPServerAuthRequest) (response *DescribeCloudNativeAPIGatewayMCPServerAuthResponse, err error) {
    return c.DescribeCloudNativeAPIGatewayMCPServerAuthWithContext(context.Background(), request)
}

// DescribeCloudNativeAPIGatewayMCPServerAuth
// 查询 MCP Server 的认证配置
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SECRETIDNOTFOUND = "AuthFailure.SecretIdNotFound"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLS = "FailedOperation.Cls"
//  FAILEDOPERATION_CLSCONFIG = "FailedOperation.ClsConfig"
//  FAILEDOPERATION_COS = "FailedOperation.Cos"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_PLUGIN = "FailedOperation.Plugin"
//  FAILEDOPERATION_RESOURCE = "FailedOperation.Resource"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  FAILEDOPERATION_VPC = "FailedOperation.Vpc"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CLBFAILURE = "InternalError.CLBFailure"
//  INTERNALERROR_CAMNOAUTH = "InternalError.CamNoAuth"
//  INTERNALERROR_CREATEERROR = "InternalError.CreateError"
//  INTERNALERROR_DECODEERROR = "InternalError.DecodeError"
//  INTERNALERROR_EKSFAILURE = "InternalError.EKSFailure"
//  INTERNALERROR_GETCREDENTIAL = "InternalError.GetCredential"
//  INTERNALERROR_GETROLEERROR = "InternalError.GetRoleError"
//  INTERNALERROR_GOVERNANCEFAILURE = "InternalError.GovernanceFailure"
//  INTERNALERROR_HTTPSTATUSCODEERROR = "InternalError.HttpStatusCodeError"
//  INTERNALERROR_IOERROR = "InternalError.IOError"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INTERNALERROR_OPERATIONFAILED = "InternalError.OperationFailed"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_TKECLUSTERNOTREADY = "InternalError.TKEClusterNotReady"
//  INTERNALERROR_TKEFAILURE = "InternalError.TKEFailure"
//  INTERNALERROR_TAGFAILURE = "InternalError.TagFailure"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INTERNALERROR_UPDATEERROR = "InternalError.UpdateError"
//  INTERNALERROR_VPCFAILURE = "InternalError.VPCFailure"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARSEJSONERROR = "InvalidParameter.ParseJsonError"
//  INVALIDPARAMETERCOMBINATION = "InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ACTION = "InvalidParameterValue.Action"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_BUCKETNAME = "InvalidParameterValue.BucketName"
//  INVALIDPARAMETERVALUE_COS = "InvalidParameterValue.Cos"
//  INVALIDPARAMETERVALUE_CREATEERROR = "InvalidParameterValue.CreateError"
//  INVALIDPARAMETERVALUE_DESCRIPTION = "InvalidParameterValue.Description"
//  INVALIDPARAMETERVALUE_GATEWAYID = "InvalidParameterValue.GatewayId"
//  INVALIDPARAMETERVALUE_GATEWAYVERSION = "InvalidParameterValue.GatewayVersion"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_KEYNOTEXIST = "InvalidParameterValue.KeyNotExist"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  INVALIDPARAMETERVALUE_OBJECTNAME = "InvalidParameterValue.ObjectName"
//  INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"
//  INVALIDPARAMETERVALUE_PLUGIN = "InvalidParameterValue.Plugin"
//  INVALIDPARAMETERVALUE_QUERYERROR = "InvalidParameterValue.QueryError"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  INVALIDPARAMETERVALUE_SPECIFICATION = "InvalidParameterValue.Specification"
//  INVALIDPARAMETERVALUE_SYSTEMPARAMETER = "InvalidParameterValue.SystemParameter"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  INVALIDPARAMETERVALUE_UPDATEERROR = "InvalidParameterValue.UpdateError"
//  INVALIDPARAMETERVALUE_ZIPFILE = "InvalidParameterValue.ZipFile"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_LBDOMAINS = "LimitExceeded.LBDomains"
//  LIMITEXCEEDED_LIMITEXCEEDED = "LimitExceeded.LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_CREATEERROR = "MissingParameter.CreateError"
//  MISSINGPARAMETER_MISSPARAMETER = "MissingParameter.MissParameter"
//  MISSINGPARAMETER_UPDATEERROR = "MissingParameter.UpdateError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_LIMITEXCEEDED = "OperationDenied.LimitExceeded"
//  OPERATIONDENIED_OPERATIONDENIED = "OperationDenied.OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINUSE_GATEWAYSERVICEEXISTGROUPS = "ResourceInUse.GatewayServiceExistGroups"
//  RESOURCEINUSE_GATEWAYSERVICEEXISTROUTER = "ResourceInUse.GatewayServiceExistRouter"
//  RESOURCEINUSE_GATEWAYSERVICEGROUPEXISTUPSTREAM = "ResourceInUse.GatewayServiceGroupExistUpstream"
//  RESOURCEINUSE_GATEWAYSERVICESOURCEEXISTSERVICE = "ResourceInUse.GatewayServiceSourceExistService"
//  RESOURCEINUSE_INSTANCESEXISTEDINSERVICE = "ResourceInUse.InstancesExistedInService"
//  RESOURCEINUSE_SERVICESEXISTEDINNAMESPACE = "ResourceInUse.ServicesExistedInNamespace"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_CONFIGNOTEXIST = "ResourceNotFound.ConfigNotExist"
//  RESOURCENOTFOUND_FORBIDDEN = "ResourceNotFound.Forbidden"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_CAMPASSROLENOTEXIST = "UnauthorizedOperation.CamPassRoleNotExist"
//  UNAUTHORIZEDOPERATION_CLSNOTACTIVATED = "UnauthorizedOperation.ClsNotActivated"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_GOVERNANCENOTSUPPORTVISIBLE = "UnsupportedOperation.GovernanceNotSupportVisible"
func (c *Client) DescribeCloudNativeAPIGatewayMCPServerAuthWithContext(ctx context.Context, request *DescribeCloudNativeAPIGatewayMCPServerAuthRequest) (response *DescribeCloudNativeAPIGatewayMCPServerAuthResponse, err error) {
    if request == nil {
        request = NewDescribeCloudNativeAPIGatewayMCPServerAuthRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cngw", APIVersion, "DescribeCloudNativeAPIGatewayMCPServerAuth")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCloudNativeAPIGatewayMCPServerAuth require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCloudNativeAPIGatewayMCPServerAuthResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCloudNativeAPIGatewayMCPServerListRequest() (request *DescribeCloudNativeAPIGatewayMCPServerListRequest) {
    request = &DescribeCloudNativeAPIGatewayMCPServerListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cngw", APIVersion, "DescribeCloudNativeAPIGatewayMCPServerList")
    
    
    return
}

func NewDescribeCloudNativeAPIGatewayMCPServerListResponse() (response *DescribeCloudNativeAPIGatewayMCPServerListResponse) {
    response = &DescribeCloudNativeAPIGatewayMCPServerListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCloudNativeAPIGatewayMCPServerList
// AI网关查询MCP服务列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  FAILEDOPERATION_VPC = "FailedOperation.Vpc"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  INVALIDPARAMETERVALUE_SPECIFICATION = "InvalidParameterValue.Specification"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_CAMPASSROLENOTEXIST = "UnauthorizedOperation.CamPassRoleNotExist"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DescribeCloudNativeAPIGatewayMCPServerList(request *DescribeCloudNativeAPIGatewayMCPServerListRequest) (response *DescribeCloudNativeAPIGatewayMCPServerListResponse, err error) {
    return c.DescribeCloudNativeAPIGatewayMCPServerListWithContext(context.Background(), request)
}

// DescribeCloudNativeAPIGatewayMCPServerList
// AI网关查询MCP服务列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  FAILEDOPERATION_VPC = "FailedOperation.Vpc"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  INVALIDPARAMETERVALUE_SPECIFICATION = "InvalidParameterValue.Specification"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_CAMPASSROLENOTEXIST = "UnauthorizedOperation.CamPassRoleNotExist"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DescribeCloudNativeAPIGatewayMCPServerListWithContext(ctx context.Context, request *DescribeCloudNativeAPIGatewayMCPServerListRequest) (response *DescribeCloudNativeAPIGatewayMCPServerListResponse, err error) {
    if request == nil {
        request = NewDescribeCloudNativeAPIGatewayMCPServerListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cngw", APIVersion, "DescribeCloudNativeAPIGatewayMCPServerList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCloudNativeAPIGatewayMCPServerList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCloudNativeAPIGatewayMCPServerListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCloudNativeAPIGatewayMCPToolRequest() (request *DescribeCloudNativeAPIGatewayMCPToolRequest) {
    request = &DescribeCloudNativeAPIGatewayMCPToolRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cngw", APIVersion, "DescribeCloudNativeAPIGatewayMCPTool")
    
    
    return
}

func NewDescribeCloudNativeAPIGatewayMCPToolResponse() (response *DescribeCloudNativeAPIGatewayMCPToolResponse) {
    response = &DescribeCloudNativeAPIGatewayMCPToolResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCloudNativeAPIGatewayMCPTool
// 查看AI网关MCP Tool
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  FAILEDOPERATION_VPC = "FailedOperation.Vpc"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  INVALIDPARAMETERVALUE_SPECIFICATION = "InvalidParameterValue.Specification"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_CAMPASSROLENOTEXIST = "UnauthorizedOperation.CamPassRoleNotExist"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DescribeCloudNativeAPIGatewayMCPTool(request *DescribeCloudNativeAPIGatewayMCPToolRequest) (response *DescribeCloudNativeAPIGatewayMCPToolResponse, err error) {
    return c.DescribeCloudNativeAPIGatewayMCPToolWithContext(context.Background(), request)
}

// DescribeCloudNativeAPIGatewayMCPTool
// 查看AI网关MCP Tool
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  FAILEDOPERATION_VPC = "FailedOperation.Vpc"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  INVALIDPARAMETERVALUE_SPECIFICATION = "InvalidParameterValue.Specification"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_CAMPASSROLENOTEXIST = "UnauthorizedOperation.CamPassRoleNotExist"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DescribeCloudNativeAPIGatewayMCPToolWithContext(ctx context.Context, request *DescribeCloudNativeAPIGatewayMCPToolRequest) (response *DescribeCloudNativeAPIGatewayMCPToolResponse, err error) {
    if request == nil {
        request = NewDescribeCloudNativeAPIGatewayMCPToolRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cngw", APIVersion, "DescribeCloudNativeAPIGatewayMCPTool")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCloudNativeAPIGatewayMCPTool require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCloudNativeAPIGatewayMCPToolResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCloudNativeAPIGatewayMCPToolACLListRequest() (request *DescribeCloudNativeAPIGatewayMCPToolACLListRequest) {
    request = &DescribeCloudNativeAPIGatewayMCPToolACLListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cngw", APIVersion, "DescribeCloudNativeAPIGatewayMCPToolACLList")
    
    
    return
}

func NewDescribeCloudNativeAPIGatewayMCPToolACLListResponse() (response *DescribeCloudNativeAPIGatewayMCPToolACLListResponse) {
    response = &DescribeCloudNativeAPIGatewayMCPToolACLListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCloudNativeAPIGatewayMCPToolACLList
// 查询云原生网关 MCP Server 下所有 Tool 的 ACL 状态一览（含 Server ACLType 回显）。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SECRETIDNOTFOUND = "AuthFailure.SecretIdNotFound"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLS = "FailedOperation.Cls"
//  FAILEDOPERATION_CLSCONFIG = "FailedOperation.ClsConfig"
//  FAILEDOPERATION_COS = "FailedOperation.Cos"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_PLUGIN = "FailedOperation.Plugin"
//  FAILEDOPERATION_RESOURCE = "FailedOperation.Resource"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  FAILEDOPERATION_VPC = "FailedOperation.Vpc"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CLBFAILURE = "InternalError.CLBFailure"
//  INTERNALERROR_CAMNOAUTH = "InternalError.CamNoAuth"
//  INTERNALERROR_CREATEERROR = "InternalError.CreateError"
//  INTERNALERROR_DECODEERROR = "InternalError.DecodeError"
//  INTERNALERROR_EKSFAILURE = "InternalError.EKSFailure"
//  INTERNALERROR_GETCREDENTIAL = "InternalError.GetCredential"
//  INTERNALERROR_GETROLEERROR = "InternalError.GetRoleError"
//  INTERNALERROR_GOVERNANCEFAILURE = "InternalError.GovernanceFailure"
//  INTERNALERROR_HTTPSTATUSCODEERROR = "InternalError.HttpStatusCodeError"
//  INTERNALERROR_IOERROR = "InternalError.IOError"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INTERNALERROR_OPERATIONFAILED = "InternalError.OperationFailed"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_TKECLUSTERNOTREADY = "InternalError.TKEClusterNotReady"
//  INTERNALERROR_TKEFAILURE = "InternalError.TKEFailure"
//  INTERNALERROR_TAGFAILURE = "InternalError.TagFailure"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INTERNALERROR_UPDATEERROR = "InternalError.UpdateError"
//  INTERNALERROR_VPCFAILURE = "InternalError.VPCFailure"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARSEJSONERROR = "InvalidParameter.ParseJsonError"
//  INVALIDPARAMETERCOMBINATION = "InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ACTION = "InvalidParameterValue.Action"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_BUCKETNAME = "InvalidParameterValue.BucketName"
//  INVALIDPARAMETERVALUE_COS = "InvalidParameterValue.Cos"
//  INVALIDPARAMETERVALUE_CREATEERROR = "InvalidParameterValue.CreateError"
//  INVALIDPARAMETERVALUE_DESCRIPTION = "InvalidParameterValue.Description"
//  INVALIDPARAMETERVALUE_GATEWAYID = "InvalidParameterValue.GatewayId"
//  INVALIDPARAMETERVALUE_GATEWAYVERSION = "InvalidParameterValue.GatewayVersion"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_KEYNOTEXIST = "InvalidParameterValue.KeyNotExist"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  INVALIDPARAMETERVALUE_OBJECTNAME = "InvalidParameterValue.ObjectName"
//  INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"
//  INVALIDPARAMETERVALUE_PLUGIN = "InvalidParameterValue.Plugin"
//  INVALIDPARAMETERVALUE_QUERYERROR = "InvalidParameterValue.QueryError"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  INVALIDPARAMETERVALUE_SPECIFICATION = "InvalidParameterValue.Specification"
//  INVALIDPARAMETERVALUE_SYSTEMPARAMETER = "InvalidParameterValue.SystemParameter"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  INVALIDPARAMETERVALUE_UPDATEERROR = "InvalidParameterValue.UpdateError"
//  INVALIDPARAMETERVALUE_ZIPFILE = "InvalidParameterValue.ZipFile"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_LBDOMAINS = "LimitExceeded.LBDomains"
//  LIMITEXCEEDED_LIMITEXCEEDED = "LimitExceeded.LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_CREATEERROR = "MissingParameter.CreateError"
//  MISSINGPARAMETER_MISSPARAMETER = "MissingParameter.MissParameter"
//  MISSINGPARAMETER_UPDATEERROR = "MissingParameter.UpdateError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_LIMITEXCEEDED = "OperationDenied.LimitExceeded"
//  OPERATIONDENIED_OPERATIONDENIED = "OperationDenied.OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINUSE_GATEWAYSERVICEEXISTGROUPS = "ResourceInUse.GatewayServiceExistGroups"
//  RESOURCEINUSE_GATEWAYSERVICEEXISTROUTER = "ResourceInUse.GatewayServiceExistRouter"
//  RESOURCEINUSE_GATEWAYSERVICEGROUPEXISTUPSTREAM = "ResourceInUse.GatewayServiceGroupExistUpstream"
//  RESOURCEINUSE_GATEWAYSERVICESOURCEEXISTSERVICE = "ResourceInUse.GatewayServiceSourceExistService"
//  RESOURCEINUSE_INSTANCESEXISTEDINSERVICE = "ResourceInUse.InstancesExistedInService"
//  RESOURCEINUSE_SERVICESEXISTEDINNAMESPACE = "ResourceInUse.ServicesExistedInNamespace"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_CONFIGNOTEXIST = "ResourceNotFound.ConfigNotExist"
//  RESOURCENOTFOUND_FORBIDDEN = "ResourceNotFound.Forbidden"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_CAMPASSROLENOTEXIST = "UnauthorizedOperation.CamPassRoleNotExist"
//  UNAUTHORIZEDOPERATION_CLSNOTACTIVATED = "UnauthorizedOperation.ClsNotActivated"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_GOVERNANCENOTSUPPORTVISIBLE = "UnsupportedOperation.GovernanceNotSupportVisible"
func (c *Client) DescribeCloudNativeAPIGatewayMCPToolACLList(request *DescribeCloudNativeAPIGatewayMCPToolACLListRequest) (response *DescribeCloudNativeAPIGatewayMCPToolACLListResponse, err error) {
    return c.DescribeCloudNativeAPIGatewayMCPToolACLListWithContext(context.Background(), request)
}

// DescribeCloudNativeAPIGatewayMCPToolACLList
// 查询云原生网关 MCP Server 下所有 Tool 的 ACL 状态一览（含 Server ACLType 回显）。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SECRETIDNOTFOUND = "AuthFailure.SecretIdNotFound"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLS = "FailedOperation.Cls"
//  FAILEDOPERATION_CLSCONFIG = "FailedOperation.ClsConfig"
//  FAILEDOPERATION_COS = "FailedOperation.Cos"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_PLUGIN = "FailedOperation.Plugin"
//  FAILEDOPERATION_RESOURCE = "FailedOperation.Resource"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  FAILEDOPERATION_VPC = "FailedOperation.Vpc"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CLBFAILURE = "InternalError.CLBFailure"
//  INTERNALERROR_CAMNOAUTH = "InternalError.CamNoAuth"
//  INTERNALERROR_CREATEERROR = "InternalError.CreateError"
//  INTERNALERROR_DECODEERROR = "InternalError.DecodeError"
//  INTERNALERROR_EKSFAILURE = "InternalError.EKSFailure"
//  INTERNALERROR_GETCREDENTIAL = "InternalError.GetCredential"
//  INTERNALERROR_GETROLEERROR = "InternalError.GetRoleError"
//  INTERNALERROR_GOVERNANCEFAILURE = "InternalError.GovernanceFailure"
//  INTERNALERROR_HTTPSTATUSCODEERROR = "InternalError.HttpStatusCodeError"
//  INTERNALERROR_IOERROR = "InternalError.IOError"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INTERNALERROR_OPERATIONFAILED = "InternalError.OperationFailed"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_TKECLUSTERNOTREADY = "InternalError.TKEClusterNotReady"
//  INTERNALERROR_TKEFAILURE = "InternalError.TKEFailure"
//  INTERNALERROR_TAGFAILURE = "InternalError.TagFailure"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INTERNALERROR_UPDATEERROR = "InternalError.UpdateError"
//  INTERNALERROR_VPCFAILURE = "InternalError.VPCFailure"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARSEJSONERROR = "InvalidParameter.ParseJsonError"
//  INVALIDPARAMETERCOMBINATION = "InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ACTION = "InvalidParameterValue.Action"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_BUCKETNAME = "InvalidParameterValue.BucketName"
//  INVALIDPARAMETERVALUE_COS = "InvalidParameterValue.Cos"
//  INVALIDPARAMETERVALUE_CREATEERROR = "InvalidParameterValue.CreateError"
//  INVALIDPARAMETERVALUE_DESCRIPTION = "InvalidParameterValue.Description"
//  INVALIDPARAMETERVALUE_GATEWAYID = "InvalidParameterValue.GatewayId"
//  INVALIDPARAMETERVALUE_GATEWAYVERSION = "InvalidParameterValue.GatewayVersion"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_KEYNOTEXIST = "InvalidParameterValue.KeyNotExist"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  INVALIDPARAMETERVALUE_OBJECTNAME = "InvalidParameterValue.ObjectName"
//  INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"
//  INVALIDPARAMETERVALUE_PLUGIN = "InvalidParameterValue.Plugin"
//  INVALIDPARAMETERVALUE_QUERYERROR = "InvalidParameterValue.QueryError"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  INVALIDPARAMETERVALUE_SPECIFICATION = "InvalidParameterValue.Specification"
//  INVALIDPARAMETERVALUE_SYSTEMPARAMETER = "InvalidParameterValue.SystemParameter"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  INVALIDPARAMETERVALUE_UPDATEERROR = "InvalidParameterValue.UpdateError"
//  INVALIDPARAMETERVALUE_ZIPFILE = "InvalidParameterValue.ZipFile"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_LBDOMAINS = "LimitExceeded.LBDomains"
//  LIMITEXCEEDED_LIMITEXCEEDED = "LimitExceeded.LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_CREATEERROR = "MissingParameter.CreateError"
//  MISSINGPARAMETER_MISSPARAMETER = "MissingParameter.MissParameter"
//  MISSINGPARAMETER_UPDATEERROR = "MissingParameter.UpdateError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_LIMITEXCEEDED = "OperationDenied.LimitExceeded"
//  OPERATIONDENIED_OPERATIONDENIED = "OperationDenied.OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINUSE_GATEWAYSERVICEEXISTGROUPS = "ResourceInUse.GatewayServiceExistGroups"
//  RESOURCEINUSE_GATEWAYSERVICEEXISTROUTER = "ResourceInUse.GatewayServiceExistRouter"
//  RESOURCEINUSE_GATEWAYSERVICEGROUPEXISTUPSTREAM = "ResourceInUse.GatewayServiceGroupExistUpstream"
//  RESOURCEINUSE_GATEWAYSERVICESOURCEEXISTSERVICE = "ResourceInUse.GatewayServiceSourceExistService"
//  RESOURCEINUSE_INSTANCESEXISTEDINSERVICE = "ResourceInUse.InstancesExistedInService"
//  RESOURCEINUSE_SERVICESEXISTEDINNAMESPACE = "ResourceInUse.ServicesExistedInNamespace"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_CONFIGNOTEXIST = "ResourceNotFound.ConfigNotExist"
//  RESOURCENOTFOUND_FORBIDDEN = "ResourceNotFound.Forbidden"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_CAMPASSROLENOTEXIST = "UnauthorizedOperation.CamPassRoleNotExist"
//  UNAUTHORIZEDOPERATION_CLSNOTACTIVATED = "UnauthorizedOperation.ClsNotActivated"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_GOVERNANCENOTSUPPORTVISIBLE = "UnsupportedOperation.GovernanceNotSupportVisible"
func (c *Client) DescribeCloudNativeAPIGatewayMCPToolACLListWithContext(ctx context.Context, request *DescribeCloudNativeAPIGatewayMCPToolACLListRequest) (response *DescribeCloudNativeAPIGatewayMCPToolACLListResponse, err error) {
    if request == nil {
        request = NewDescribeCloudNativeAPIGatewayMCPToolACLListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cngw", APIVersion, "DescribeCloudNativeAPIGatewayMCPToolACLList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCloudNativeAPIGatewayMCPToolACLList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCloudNativeAPIGatewayMCPToolACLListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCloudNativeAPIGatewayMCPToolListRequest() (request *DescribeCloudNativeAPIGatewayMCPToolListRequest) {
    request = &DescribeCloudNativeAPIGatewayMCPToolListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cngw", APIVersion, "DescribeCloudNativeAPIGatewayMCPToolList")
    
    
    return
}

func NewDescribeCloudNativeAPIGatewayMCPToolListResponse() (response *DescribeCloudNativeAPIGatewayMCPToolListResponse) {
    response = &DescribeCloudNativeAPIGatewayMCPToolListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCloudNativeAPIGatewayMCPToolList
// 查询 AI 网关MCP Tool 列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  FAILEDOPERATION_VPC = "FailedOperation.Vpc"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  INVALIDPARAMETERVALUE_SPECIFICATION = "InvalidParameterValue.Specification"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_CAMPASSROLENOTEXIST = "UnauthorizedOperation.CamPassRoleNotExist"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DescribeCloudNativeAPIGatewayMCPToolList(request *DescribeCloudNativeAPIGatewayMCPToolListRequest) (response *DescribeCloudNativeAPIGatewayMCPToolListResponse, err error) {
    return c.DescribeCloudNativeAPIGatewayMCPToolListWithContext(context.Background(), request)
}

// DescribeCloudNativeAPIGatewayMCPToolList
// 查询 AI 网关MCP Tool 列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  FAILEDOPERATION_VPC = "FailedOperation.Vpc"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  INVALIDPARAMETERVALUE_SPECIFICATION = "InvalidParameterValue.Specification"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_CAMPASSROLENOTEXIST = "UnauthorizedOperation.CamPassRoleNotExist"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DescribeCloudNativeAPIGatewayMCPToolListWithContext(ctx context.Context, request *DescribeCloudNativeAPIGatewayMCPToolListRequest) (response *DescribeCloudNativeAPIGatewayMCPToolListResponse, err error) {
    if request == nil {
        request = NewDescribeCloudNativeAPIGatewayMCPToolListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cngw", APIVersion, "DescribeCloudNativeAPIGatewayMCPToolList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCloudNativeAPIGatewayMCPToolList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCloudNativeAPIGatewayMCPToolListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCloudNativeAPIGatewaySecretKeyRequest() (request *DescribeCloudNativeAPIGatewaySecretKeyRequest) {
    request = &DescribeCloudNativeAPIGatewaySecretKeyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cngw", APIVersion, "DescribeCloudNativeAPIGatewaySecretKey")
    
    
    return
}

func NewDescribeCloudNativeAPIGatewaySecretKeyResponse() (response *DescribeCloudNativeAPIGatewaySecretKeyResponse) {
    response = &DescribeCloudNativeAPIGatewaySecretKeyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCloudNativeAPIGatewaySecretKey
// 查询密钥详情（SecretValue 字段会被掩码）。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  FAILEDOPERATION_VPC = "FailedOperation.Vpc"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  INVALIDPARAMETERVALUE_SPECIFICATION = "InvalidParameterValue.Specification"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_CAMPASSROLENOTEXIST = "UnauthorizedOperation.CamPassRoleNotExist"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DescribeCloudNativeAPIGatewaySecretKey(request *DescribeCloudNativeAPIGatewaySecretKeyRequest) (response *DescribeCloudNativeAPIGatewaySecretKeyResponse, err error) {
    return c.DescribeCloudNativeAPIGatewaySecretKeyWithContext(context.Background(), request)
}

// DescribeCloudNativeAPIGatewaySecretKey
// 查询密钥详情（SecretValue 字段会被掩码）。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  FAILEDOPERATION_VPC = "FailedOperation.Vpc"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  INVALIDPARAMETERVALUE_SPECIFICATION = "InvalidParameterValue.Specification"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_CAMPASSROLENOTEXIST = "UnauthorizedOperation.CamPassRoleNotExist"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DescribeCloudNativeAPIGatewaySecretKeyWithContext(ctx context.Context, request *DescribeCloudNativeAPIGatewaySecretKeyRequest) (response *DescribeCloudNativeAPIGatewaySecretKeyResponse, err error) {
    if request == nil {
        request = NewDescribeCloudNativeAPIGatewaySecretKeyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cngw", APIVersion, "DescribeCloudNativeAPIGatewaySecretKey")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCloudNativeAPIGatewaySecretKey require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCloudNativeAPIGatewaySecretKeyResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCloudNativeAPIGatewaySecretKeyValueRequest() (request *DescribeCloudNativeAPIGatewaySecretKeyValueRequest) {
    request = &DescribeCloudNativeAPIGatewaySecretKeyValueRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cngw", APIVersion, "DescribeCloudNativeAPIGatewaySecretKeyValue")
    
    
    return
}

func NewDescribeCloudNativeAPIGatewaySecretKeyValueResponse() (response *DescribeCloudNativeAPIGatewaySecretKeyValueResponse) {
    response = &DescribeCloudNativeAPIGatewaySecretKeyValueResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCloudNativeAPIGatewaySecretKeyValue
// 查询密钥明文值（KMS 类型密钥不可获取）。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  FAILEDOPERATION_VPC = "FailedOperation.Vpc"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  INVALIDPARAMETERVALUE_SPECIFICATION = "InvalidParameterValue.Specification"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_CAMPASSROLENOTEXIST = "UnauthorizedOperation.CamPassRoleNotExist"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DescribeCloudNativeAPIGatewaySecretKeyValue(request *DescribeCloudNativeAPIGatewaySecretKeyValueRequest) (response *DescribeCloudNativeAPIGatewaySecretKeyValueResponse, err error) {
    return c.DescribeCloudNativeAPIGatewaySecretKeyValueWithContext(context.Background(), request)
}

// DescribeCloudNativeAPIGatewaySecretKeyValue
// 查询密钥明文值（KMS 类型密钥不可获取）。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  FAILEDOPERATION_VPC = "FailedOperation.Vpc"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  INVALIDPARAMETERVALUE_SPECIFICATION = "InvalidParameterValue.Specification"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_CAMPASSROLENOTEXIST = "UnauthorizedOperation.CamPassRoleNotExist"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DescribeCloudNativeAPIGatewaySecretKeyValueWithContext(ctx context.Context, request *DescribeCloudNativeAPIGatewaySecretKeyValueRequest) (response *DescribeCloudNativeAPIGatewaySecretKeyValueResponse, err error) {
    if request == nil {
        request = NewDescribeCloudNativeAPIGatewaySecretKeyValueRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cngw", APIVersion, "DescribeCloudNativeAPIGatewaySecretKeyValue")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCloudNativeAPIGatewaySecretKeyValue require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCloudNativeAPIGatewaySecretKeyValueResponse()
    err = c.Send(request, response)
    return
}

func NewModifyCloudNativeAPIGatewayConsumerRequest() (request *ModifyCloudNativeAPIGatewayConsumerRequest) {
    request = &ModifyCloudNativeAPIGatewayConsumerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cngw", APIVersion, "ModifyCloudNativeAPIGatewayConsumer")
    
    
    return
}

func NewModifyCloudNativeAPIGatewayConsumerResponse() (response *ModifyCloudNativeAPIGatewayConsumerResponse) {
    response = &ModifyCloudNativeAPIGatewayConsumerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyCloudNativeAPIGatewayConsumer
// 修改AI网关消费者
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  FAILEDOPERATION_VPC = "FailedOperation.Vpc"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  INVALIDPARAMETERVALUE_SPECIFICATION = "InvalidParameterValue.Specification"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_CAMPASSROLENOTEXIST = "UnauthorizedOperation.CamPassRoleNotExist"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) ModifyCloudNativeAPIGatewayConsumer(request *ModifyCloudNativeAPIGatewayConsumerRequest) (response *ModifyCloudNativeAPIGatewayConsumerResponse, err error) {
    return c.ModifyCloudNativeAPIGatewayConsumerWithContext(context.Background(), request)
}

// ModifyCloudNativeAPIGatewayConsumer
// 修改AI网关消费者
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  FAILEDOPERATION_VPC = "FailedOperation.Vpc"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  INVALIDPARAMETERVALUE_SPECIFICATION = "InvalidParameterValue.Specification"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_CAMPASSROLENOTEXIST = "UnauthorizedOperation.CamPassRoleNotExist"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) ModifyCloudNativeAPIGatewayConsumerWithContext(ctx context.Context, request *ModifyCloudNativeAPIGatewayConsumerRequest) (response *ModifyCloudNativeAPIGatewayConsumerResponse, err error) {
    if request == nil {
        request = NewModifyCloudNativeAPIGatewayConsumerRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cngw", APIVersion, "ModifyCloudNativeAPIGatewayConsumer")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyCloudNativeAPIGatewayConsumer require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyCloudNativeAPIGatewayConsumerResponse()
    err = c.Send(request, response)
    return
}

func NewModifyCloudNativeAPIGatewayConsumerGroupRequest() (request *ModifyCloudNativeAPIGatewayConsumerGroupRequest) {
    request = &ModifyCloudNativeAPIGatewayConsumerGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cngw", APIVersion, "ModifyCloudNativeAPIGatewayConsumerGroup")
    
    
    return
}

func NewModifyCloudNativeAPIGatewayConsumerGroupResponse() (response *ModifyCloudNativeAPIGatewayConsumerGroupResponse) {
    response = &ModifyCloudNativeAPIGatewayConsumerGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyCloudNativeAPIGatewayConsumerGroup
// 修改消费者组。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  FAILEDOPERATION_VPC = "FailedOperation.Vpc"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  INVALIDPARAMETERVALUE_SPECIFICATION = "InvalidParameterValue.Specification"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_CAMPASSROLENOTEXIST = "UnauthorizedOperation.CamPassRoleNotExist"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) ModifyCloudNativeAPIGatewayConsumerGroup(request *ModifyCloudNativeAPIGatewayConsumerGroupRequest) (response *ModifyCloudNativeAPIGatewayConsumerGroupResponse, err error) {
    return c.ModifyCloudNativeAPIGatewayConsumerGroupWithContext(context.Background(), request)
}

// ModifyCloudNativeAPIGatewayConsumerGroup
// 修改消费者组。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  FAILEDOPERATION_VPC = "FailedOperation.Vpc"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  INVALIDPARAMETERVALUE_SPECIFICATION = "InvalidParameterValue.Specification"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_CAMPASSROLENOTEXIST = "UnauthorizedOperation.CamPassRoleNotExist"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) ModifyCloudNativeAPIGatewayConsumerGroupWithContext(ctx context.Context, request *ModifyCloudNativeAPIGatewayConsumerGroupRequest) (response *ModifyCloudNativeAPIGatewayConsumerGroupResponse, err error) {
    if request == nil {
        request = NewModifyCloudNativeAPIGatewayConsumerGroupRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cngw", APIVersion, "ModifyCloudNativeAPIGatewayConsumerGroup")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyCloudNativeAPIGatewayConsumerGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyCloudNativeAPIGatewayConsumerGroupResponse()
    err = c.Send(request, response)
    return
}

func NewModifyCloudNativeAPIGatewayLLMModelAPIRequest() (request *ModifyCloudNativeAPIGatewayLLMModelAPIRequest) {
    request = &ModifyCloudNativeAPIGatewayLLMModelAPIRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cngw", APIVersion, "ModifyCloudNativeAPIGatewayLLMModelAPI")
    
    
    return
}

func NewModifyCloudNativeAPIGatewayLLMModelAPIResponse() (response *ModifyCloudNativeAPIGatewayLLMModelAPIResponse) {
    response = &ModifyCloudNativeAPIGatewayLLMModelAPIResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyCloudNativeAPIGatewayLLMModelAPI
// 修改 LLM 模型 API。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  LIMITEXCEEDED_LIMITEXCEEDED = "LimitExceeded.LimitExceeded"
//  MISSINGPARAMETER_MISSPARAMETER = "MissingParameter.MissParameter"
//  OPERATIONDENIED_OPERATIONDENIED = "OperationDenied.OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) ModifyCloudNativeAPIGatewayLLMModelAPI(request *ModifyCloudNativeAPIGatewayLLMModelAPIRequest) (response *ModifyCloudNativeAPIGatewayLLMModelAPIResponse, err error) {
    return c.ModifyCloudNativeAPIGatewayLLMModelAPIWithContext(context.Background(), request)
}

// ModifyCloudNativeAPIGatewayLLMModelAPI
// 修改 LLM 模型 API。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  LIMITEXCEEDED_LIMITEXCEEDED = "LimitExceeded.LimitExceeded"
//  MISSINGPARAMETER_MISSPARAMETER = "MissingParameter.MissParameter"
//  OPERATIONDENIED_OPERATIONDENIED = "OperationDenied.OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) ModifyCloudNativeAPIGatewayLLMModelAPIWithContext(ctx context.Context, request *ModifyCloudNativeAPIGatewayLLMModelAPIRequest) (response *ModifyCloudNativeAPIGatewayLLMModelAPIResponse, err error) {
    if request == nil {
        request = NewModifyCloudNativeAPIGatewayLLMModelAPIRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cngw", APIVersion, "ModifyCloudNativeAPIGatewayLLMModelAPI")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyCloudNativeAPIGatewayLLMModelAPI require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyCloudNativeAPIGatewayLLMModelAPIResponse()
    err = c.Send(request, response)
    return
}

func NewModifyCloudNativeAPIGatewayLLMModelServiceRequest() (request *ModifyCloudNativeAPIGatewayLLMModelServiceRequest) {
    request = &ModifyCloudNativeAPIGatewayLLMModelServiceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cngw", APIVersion, "ModifyCloudNativeAPIGatewayLLMModelService")
    
    
    return
}

func NewModifyCloudNativeAPIGatewayLLMModelServiceResponse() (response *ModifyCloudNativeAPIGatewayLLMModelServiceResponse) {
    response = &ModifyCloudNativeAPIGatewayLLMModelServiceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyCloudNativeAPIGatewayLLMModelService
// 修改 LLM 模型服务。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  LIMITEXCEEDED_LIMITEXCEEDED = "LimitExceeded.LimitExceeded"
//  MISSINGPARAMETER_MISSPARAMETER = "MissingParameter.MissParameter"
//  OPERATIONDENIED_OPERATIONDENIED = "OperationDenied.OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) ModifyCloudNativeAPIGatewayLLMModelService(request *ModifyCloudNativeAPIGatewayLLMModelServiceRequest) (response *ModifyCloudNativeAPIGatewayLLMModelServiceResponse, err error) {
    return c.ModifyCloudNativeAPIGatewayLLMModelServiceWithContext(context.Background(), request)
}

// ModifyCloudNativeAPIGatewayLLMModelService
// 修改 LLM 模型服务。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  LIMITEXCEEDED_LIMITEXCEEDED = "LimitExceeded.LimitExceeded"
//  MISSINGPARAMETER_MISSPARAMETER = "MissingParameter.MissParameter"
//  OPERATIONDENIED_OPERATIONDENIED = "OperationDenied.OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) ModifyCloudNativeAPIGatewayLLMModelServiceWithContext(ctx context.Context, request *ModifyCloudNativeAPIGatewayLLMModelServiceRequest) (response *ModifyCloudNativeAPIGatewayLLMModelServiceResponse, err error) {
    if request == nil {
        request = NewModifyCloudNativeAPIGatewayLLMModelServiceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cngw", APIVersion, "ModifyCloudNativeAPIGatewayLLMModelService")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyCloudNativeAPIGatewayLLMModelService require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyCloudNativeAPIGatewayLLMModelServiceResponse()
    err = c.Send(request, response)
    return
}

func NewModifyCloudNativeAPIGatewayMCPServerRequest() (request *ModifyCloudNativeAPIGatewayMCPServerRequest) {
    request = &ModifyCloudNativeAPIGatewayMCPServerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cngw", APIVersion, "ModifyCloudNativeAPIGatewayMCPServer")
    
    
    return
}

func NewModifyCloudNativeAPIGatewayMCPServerResponse() (response *ModifyCloudNativeAPIGatewayMCPServerResponse) {
    response = &ModifyCloudNativeAPIGatewayMCPServerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyCloudNativeAPIGatewayMCPServer
// 修改MCP服务配置
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  FAILEDOPERATION_VPC = "FailedOperation.Vpc"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  INVALIDPARAMETERVALUE_SPECIFICATION = "InvalidParameterValue.Specification"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_CAMPASSROLENOTEXIST = "UnauthorizedOperation.CamPassRoleNotExist"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) ModifyCloudNativeAPIGatewayMCPServer(request *ModifyCloudNativeAPIGatewayMCPServerRequest) (response *ModifyCloudNativeAPIGatewayMCPServerResponse, err error) {
    return c.ModifyCloudNativeAPIGatewayMCPServerWithContext(context.Background(), request)
}

// ModifyCloudNativeAPIGatewayMCPServer
// 修改MCP服务配置
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  FAILEDOPERATION_VPC = "FailedOperation.Vpc"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  INVALIDPARAMETERVALUE_SPECIFICATION = "InvalidParameterValue.Specification"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_CAMPASSROLENOTEXIST = "UnauthorizedOperation.CamPassRoleNotExist"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) ModifyCloudNativeAPIGatewayMCPServerWithContext(ctx context.Context, request *ModifyCloudNativeAPIGatewayMCPServerRequest) (response *ModifyCloudNativeAPIGatewayMCPServerResponse, err error) {
    if request == nil {
        request = NewModifyCloudNativeAPIGatewayMCPServerRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cngw", APIVersion, "ModifyCloudNativeAPIGatewayMCPServer")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyCloudNativeAPIGatewayMCPServer require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyCloudNativeAPIGatewayMCPServerResponse()
    err = c.Send(request, response)
    return
}

func NewModifyCloudNativeAPIGatewayMCPServerACLRequest() (request *ModifyCloudNativeAPIGatewayMCPServerACLRequest) {
    request = &ModifyCloudNativeAPIGatewayMCPServerACLRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cngw", APIVersion, "ModifyCloudNativeAPIGatewayMCPServerACL")
    
    
    return
}

func NewModifyCloudNativeAPIGatewayMCPServerACLResponse() (response *ModifyCloudNativeAPIGatewayMCPServerACLResponse) {
    response = &ModifyCloudNativeAPIGatewayMCPServerACLResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyCloudNativeAPIGatewayMCPServerACL
// 修改 MCP Server ACL
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SECRETIDNOTFOUND = "AuthFailure.SecretIdNotFound"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLS = "FailedOperation.Cls"
//  FAILEDOPERATION_CLSCONFIG = "FailedOperation.ClsConfig"
//  FAILEDOPERATION_COS = "FailedOperation.Cos"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_PLUGIN = "FailedOperation.Plugin"
//  FAILEDOPERATION_RESOURCE = "FailedOperation.Resource"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  FAILEDOPERATION_VPC = "FailedOperation.Vpc"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CLBFAILURE = "InternalError.CLBFailure"
//  INTERNALERROR_CAMNOAUTH = "InternalError.CamNoAuth"
//  INTERNALERROR_CREATEERROR = "InternalError.CreateError"
//  INTERNALERROR_DECODEERROR = "InternalError.DecodeError"
//  INTERNALERROR_EKSFAILURE = "InternalError.EKSFailure"
//  INTERNALERROR_GETCREDENTIAL = "InternalError.GetCredential"
//  INTERNALERROR_GETROLEERROR = "InternalError.GetRoleError"
//  INTERNALERROR_GOVERNANCEFAILURE = "InternalError.GovernanceFailure"
//  INTERNALERROR_HTTPSTATUSCODEERROR = "InternalError.HttpStatusCodeError"
//  INTERNALERROR_IOERROR = "InternalError.IOError"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INTERNALERROR_OPERATIONFAILED = "InternalError.OperationFailed"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_TKECLUSTERNOTREADY = "InternalError.TKEClusterNotReady"
//  INTERNALERROR_TKEFAILURE = "InternalError.TKEFailure"
//  INTERNALERROR_TAGFAILURE = "InternalError.TagFailure"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INTERNALERROR_UPDATEERROR = "InternalError.UpdateError"
//  INTERNALERROR_VPCFAILURE = "InternalError.VPCFailure"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARSEJSONERROR = "InvalidParameter.ParseJsonError"
//  INVALIDPARAMETERCOMBINATION = "InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ACTION = "InvalidParameterValue.Action"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_BUCKETNAME = "InvalidParameterValue.BucketName"
//  INVALIDPARAMETERVALUE_COS = "InvalidParameterValue.Cos"
//  INVALIDPARAMETERVALUE_CREATEERROR = "InvalidParameterValue.CreateError"
//  INVALIDPARAMETERVALUE_DESCRIPTION = "InvalidParameterValue.Description"
//  INVALIDPARAMETERVALUE_GATEWAYID = "InvalidParameterValue.GatewayId"
//  INVALIDPARAMETERVALUE_GATEWAYVERSION = "InvalidParameterValue.GatewayVersion"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_KEYNOTEXIST = "InvalidParameterValue.KeyNotExist"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  INVALIDPARAMETERVALUE_OBJECTNAME = "InvalidParameterValue.ObjectName"
//  INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"
//  INVALIDPARAMETERVALUE_PLUGIN = "InvalidParameterValue.Plugin"
//  INVALIDPARAMETERVALUE_QUERYERROR = "InvalidParameterValue.QueryError"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  INVALIDPARAMETERVALUE_SPECIFICATION = "InvalidParameterValue.Specification"
//  INVALIDPARAMETERVALUE_SYSTEMPARAMETER = "InvalidParameterValue.SystemParameter"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  INVALIDPARAMETERVALUE_UPDATEERROR = "InvalidParameterValue.UpdateError"
//  INVALIDPARAMETERVALUE_ZIPFILE = "InvalidParameterValue.ZipFile"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_LBDOMAINS = "LimitExceeded.LBDomains"
//  LIMITEXCEEDED_LIMITEXCEEDED = "LimitExceeded.LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_CREATEERROR = "MissingParameter.CreateError"
//  MISSINGPARAMETER_MISSPARAMETER = "MissingParameter.MissParameter"
//  MISSINGPARAMETER_UPDATEERROR = "MissingParameter.UpdateError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_LIMITEXCEEDED = "OperationDenied.LimitExceeded"
//  OPERATIONDENIED_OPERATIONDENIED = "OperationDenied.OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINUSE_GATEWAYSERVICEEXISTGROUPS = "ResourceInUse.GatewayServiceExistGroups"
//  RESOURCEINUSE_GATEWAYSERVICEEXISTROUTER = "ResourceInUse.GatewayServiceExistRouter"
//  RESOURCEINUSE_GATEWAYSERVICEGROUPEXISTUPSTREAM = "ResourceInUse.GatewayServiceGroupExistUpstream"
//  RESOURCEINUSE_GATEWAYSERVICESOURCEEXISTSERVICE = "ResourceInUse.GatewayServiceSourceExistService"
//  RESOURCEINUSE_INSTANCESEXISTEDINSERVICE = "ResourceInUse.InstancesExistedInService"
//  RESOURCEINUSE_SERVICESEXISTEDINNAMESPACE = "ResourceInUse.ServicesExistedInNamespace"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_CONFIGNOTEXIST = "ResourceNotFound.ConfigNotExist"
//  RESOURCENOTFOUND_FORBIDDEN = "ResourceNotFound.Forbidden"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_CAMPASSROLENOTEXIST = "UnauthorizedOperation.CamPassRoleNotExist"
//  UNAUTHORIZEDOPERATION_CLSNOTACTIVATED = "UnauthorizedOperation.ClsNotActivated"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_GOVERNANCENOTSUPPORTVISIBLE = "UnsupportedOperation.GovernanceNotSupportVisible"
func (c *Client) ModifyCloudNativeAPIGatewayMCPServerACL(request *ModifyCloudNativeAPIGatewayMCPServerACLRequest) (response *ModifyCloudNativeAPIGatewayMCPServerACLResponse, err error) {
    return c.ModifyCloudNativeAPIGatewayMCPServerACLWithContext(context.Background(), request)
}

// ModifyCloudNativeAPIGatewayMCPServerACL
// 修改 MCP Server ACL
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SECRETIDNOTFOUND = "AuthFailure.SecretIdNotFound"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLS = "FailedOperation.Cls"
//  FAILEDOPERATION_CLSCONFIG = "FailedOperation.ClsConfig"
//  FAILEDOPERATION_COS = "FailedOperation.Cos"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_PLUGIN = "FailedOperation.Plugin"
//  FAILEDOPERATION_RESOURCE = "FailedOperation.Resource"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  FAILEDOPERATION_VPC = "FailedOperation.Vpc"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CLBFAILURE = "InternalError.CLBFailure"
//  INTERNALERROR_CAMNOAUTH = "InternalError.CamNoAuth"
//  INTERNALERROR_CREATEERROR = "InternalError.CreateError"
//  INTERNALERROR_DECODEERROR = "InternalError.DecodeError"
//  INTERNALERROR_EKSFAILURE = "InternalError.EKSFailure"
//  INTERNALERROR_GETCREDENTIAL = "InternalError.GetCredential"
//  INTERNALERROR_GETROLEERROR = "InternalError.GetRoleError"
//  INTERNALERROR_GOVERNANCEFAILURE = "InternalError.GovernanceFailure"
//  INTERNALERROR_HTTPSTATUSCODEERROR = "InternalError.HttpStatusCodeError"
//  INTERNALERROR_IOERROR = "InternalError.IOError"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INTERNALERROR_OPERATIONFAILED = "InternalError.OperationFailed"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_TKECLUSTERNOTREADY = "InternalError.TKEClusterNotReady"
//  INTERNALERROR_TKEFAILURE = "InternalError.TKEFailure"
//  INTERNALERROR_TAGFAILURE = "InternalError.TagFailure"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INTERNALERROR_UPDATEERROR = "InternalError.UpdateError"
//  INTERNALERROR_VPCFAILURE = "InternalError.VPCFailure"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARSEJSONERROR = "InvalidParameter.ParseJsonError"
//  INVALIDPARAMETERCOMBINATION = "InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ACTION = "InvalidParameterValue.Action"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_BUCKETNAME = "InvalidParameterValue.BucketName"
//  INVALIDPARAMETERVALUE_COS = "InvalidParameterValue.Cos"
//  INVALIDPARAMETERVALUE_CREATEERROR = "InvalidParameterValue.CreateError"
//  INVALIDPARAMETERVALUE_DESCRIPTION = "InvalidParameterValue.Description"
//  INVALIDPARAMETERVALUE_GATEWAYID = "InvalidParameterValue.GatewayId"
//  INVALIDPARAMETERVALUE_GATEWAYVERSION = "InvalidParameterValue.GatewayVersion"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_KEYNOTEXIST = "InvalidParameterValue.KeyNotExist"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  INVALIDPARAMETERVALUE_OBJECTNAME = "InvalidParameterValue.ObjectName"
//  INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"
//  INVALIDPARAMETERVALUE_PLUGIN = "InvalidParameterValue.Plugin"
//  INVALIDPARAMETERVALUE_QUERYERROR = "InvalidParameterValue.QueryError"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  INVALIDPARAMETERVALUE_SPECIFICATION = "InvalidParameterValue.Specification"
//  INVALIDPARAMETERVALUE_SYSTEMPARAMETER = "InvalidParameterValue.SystemParameter"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  INVALIDPARAMETERVALUE_UPDATEERROR = "InvalidParameterValue.UpdateError"
//  INVALIDPARAMETERVALUE_ZIPFILE = "InvalidParameterValue.ZipFile"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_LBDOMAINS = "LimitExceeded.LBDomains"
//  LIMITEXCEEDED_LIMITEXCEEDED = "LimitExceeded.LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_CREATEERROR = "MissingParameter.CreateError"
//  MISSINGPARAMETER_MISSPARAMETER = "MissingParameter.MissParameter"
//  MISSINGPARAMETER_UPDATEERROR = "MissingParameter.UpdateError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_LIMITEXCEEDED = "OperationDenied.LimitExceeded"
//  OPERATIONDENIED_OPERATIONDENIED = "OperationDenied.OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINUSE_GATEWAYSERVICEEXISTGROUPS = "ResourceInUse.GatewayServiceExistGroups"
//  RESOURCEINUSE_GATEWAYSERVICEEXISTROUTER = "ResourceInUse.GatewayServiceExistRouter"
//  RESOURCEINUSE_GATEWAYSERVICEGROUPEXISTUPSTREAM = "ResourceInUse.GatewayServiceGroupExistUpstream"
//  RESOURCEINUSE_GATEWAYSERVICESOURCEEXISTSERVICE = "ResourceInUse.GatewayServiceSourceExistService"
//  RESOURCEINUSE_INSTANCESEXISTEDINSERVICE = "ResourceInUse.InstancesExistedInService"
//  RESOURCEINUSE_SERVICESEXISTEDINNAMESPACE = "ResourceInUse.ServicesExistedInNamespace"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_CONFIGNOTEXIST = "ResourceNotFound.ConfigNotExist"
//  RESOURCENOTFOUND_FORBIDDEN = "ResourceNotFound.Forbidden"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_CAMPASSROLENOTEXIST = "UnauthorizedOperation.CamPassRoleNotExist"
//  UNAUTHORIZEDOPERATION_CLSNOTACTIVATED = "UnauthorizedOperation.ClsNotActivated"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_GOVERNANCENOTSUPPORTVISIBLE = "UnsupportedOperation.GovernanceNotSupportVisible"
func (c *Client) ModifyCloudNativeAPIGatewayMCPServerACLWithContext(ctx context.Context, request *ModifyCloudNativeAPIGatewayMCPServerACLRequest) (response *ModifyCloudNativeAPIGatewayMCPServerACLResponse, err error) {
    if request == nil {
        request = NewModifyCloudNativeAPIGatewayMCPServerACLRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cngw", APIVersion, "ModifyCloudNativeAPIGatewayMCPServerACL")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyCloudNativeAPIGatewayMCPServerACL require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyCloudNativeAPIGatewayMCPServerACLResponse()
    err = c.Send(request, response)
    return
}

func NewModifyCloudNativeAPIGatewayMCPServerAuthRequest() (request *ModifyCloudNativeAPIGatewayMCPServerAuthRequest) {
    request = &ModifyCloudNativeAPIGatewayMCPServerAuthRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cngw", APIVersion, "ModifyCloudNativeAPIGatewayMCPServerAuth")
    
    
    return
}

func NewModifyCloudNativeAPIGatewayMCPServerAuthResponse() (response *ModifyCloudNativeAPIGatewayMCPServerAuthResponse) {
    response = &ModifyCloudNativeAPIGatewayMCPServerAuthResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyCloudNativeAPIGatewayMCPServerAuth
// 修改 MCP Server 的认证配置
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SECRETIDNOTFOUND = "AuthFailure.SecretIdNotFound"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLS = "FailedOperation.Cls"
//  FAILEDOPERATION_CLSCONFIG = "FailedOperation.ClsConfig"
//  FAILEDOPERATION_COS = "FailedOperation.Cos"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_PLUGIN = "FailedOperation.Plugin"
//  FAILEDOPERATION_RESOURCE = "FailedOperation.Resource"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  FAILEDOPERATION_VPC = "FailedOperation.Vpc"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CLBFAILURE = "InternalError.CLBFailure"
//  INTERNALERROR_CAMNOAUTH = "InternalError.CamNoAuth"
//  INTERNALERROR_CREATEERROR = "InternalError.CreateError"
//  INTERNALERROR_DECODEERROR = "InternalError.DecodeError"
//  INTERNALERROR_EKSFAILURE = "InternalError.EKSFailure"
//  INTERNALERROR_GETCREDENTIAL = "InternalError.GetCredential"
//  INTERNALERROR_GETROLEERROR = "InternalError.GetRoleError"
//  INTERNALERROR_GOVERNANCEFAILURE = "InternalError.GovernanceFailure"
//  INTERNALERROR_HTTPSTATUSCODEERROR = "InternalError.HttpStatusCodeError"
//  INTERNALERROR_IOERROR = "InternalError.IOError"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INTERNALERROR_OPERATIONFAILED = "InternalError.OperationFailed"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_TKECLUSTERNOTREADY = "InternalError.TKEClusterNotReady"
//  INTERNALERROR_TKEFAILURE = "InternalError.TKEFailure"
//  INTERNALERROR_TAGFAILURE = "InternalError.TagFailure"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INTERNALERROR_UPDATEERROR = "InternalError.UpdateError"
//  INTERNALERROR_VPCFAILURE = "InternalError.VPCFailure"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARSEJSONERROR = "InvalidParameter.ParseJsonError"
//  INVALIDPARAMETERCOMBINATION = "InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ACTION = "InvalidParameterValue.Action"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_BUCKETNAME = "InvalidParameterValue.BucketName"
//  INVALIDPARAMETERVALUE_COS = "InvalidParameterValue.Cos"
//  INVALIDPARAMETERVALUE_CREATEERROR = "InvalidParameterValue.CreateError"
//  INVALIDPARAMETERVALUE_DESCRIPTION = "InvalidParameterValue.Description"
//  INVALIDPARAMETERVALUE_GATEWAYID = "InvalidParameterValue.GatewayId"
//  INVALIDPARAMETERVALUE_GATEWAYVERSION = "InvalidParameterValue.GatewayVersion"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_KEYNOTEXIST = "InvalidParameterValue.KeyNotExist"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  INVALIDPARAMETERVALUE_OBJECTNAME = "InvalidParameterValue.ObjectName"
//  INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"
//  INVALIDPARAMETERVALUE_PLUGIN = "InvalidParameterValue.Plugin"
//  INVALIDPARAMETERVALUE_QUERYERROR = "InvalidParameterValue.QueryError"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  INVALIDPARAMETERVALUE_SPECIFICATION = "InvalidParameterValue.Specification"
//  INVALIDPARAMETERVALUE_SYSTEMPARAMETER = "InvalidParameterValue.SystemParameter"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  INVALIDPARAMETERVALUE_UPDATEERROR = "InvalidParameterValue.UpdateError"
//  INVALIDPARAMETERVALUE_ZIPFILE = "InvalidParameterValue.ZipFile"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_LBDOMAINS = "LimitExceeded.LBDomains"
//  LIMITEXCEEDED_LIMITEXCEEDED = "LimitExceeded.LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_CREATEERROR = "MissingParameter.CreateError"
//  MISSINGPARAMETER_MISSPARAMETER = "MissingParameter.MissParameter"
//  MISSINGPARAMETER_UPDATEERROR = "MissingParameter.UpdateError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_LIMITEXCEEDED = "OperationDenied.LimitExceeded"
//  OPERATIONDENIED_OPERATIONDENIED = "OperationDenied.OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINUSE_GATEWAYSERVICEEXISTGROUPS = "ResourceInUse.GatewayServiceExistGroups"
//  RESOURCEINUSE_GATEWAYSERVICEEXISTROUTER = "ResourceInUse.GatewayServiceExistRouter"
//  RESOURCEINUSE_GATEWAYSERVICEGROUPEXISTUPSTREAM = "ResourceInUse.GatewayServiceGroupExistUpstream"
//  RESOURCEINUSE_GATEWAYSERVICESOURCEEXISTSERVICE = "ResourceInUse.GatewayServiceSourceExistService"
//  RESOURCEINUSE_INSTANCESEXISTEDINSERVICE = "ResourceInUse.InstancesExistedInService"
//  RESOURCEINUSE_SERVICESEXISTEDINNAMESPACE = "ResourceInUse.ServicesExistedInNamespace"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_CONFIGNOTEXIST = "ResourceNotFound.ConfigNotExist"
//  RESOURCENOTFOUND_FORBIDDEN = "ResourceNotFound.Forbidden"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_CAMPASSROLENOTEXIST = "UnauthorizedOperation.CamPassRoleNotExist"
//  UNAUTHORIZEDOPERATION_CLSNOTACTIVATED = "UnauthorizedOperation.ClsNotActivated"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_GOVERNANCENOTSUPPORTVISIBLE = "UnsupportedOperation.GovernanceNotSupportVisible"
func (c *Client) ModifyCloudNativeAPIGatewayMCPServerAuth(request *ModifyCloudNativeAPIGatewayMCPServerAuthRequest) (response *ModifyCloudNativeAPIGatewayMCPServerAuthResponse, err error) {
    return c.ModifyCloudNativeAPIGatewayMCPServerAuthWithContext(context.Background(), request)
}

// ModifyCloudNativeAPIGatewayMCPServerAuth
// 修改 MCP Server 的认证配置
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SECRETIDNOTFOUND = "AuthFailure.SecretIdNotFound"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLS = "FailedOperation.Cls"
//  FAILEDOPERATION_CLSCONFIG = "FailedOperation.ClsConfig"
//  FAILEDOPERATION_COS = "FailedOperation.Cos"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_PLUGIN = "FailedOperation.Plugin"
//  FAILEDOPERATION_RESOURCE = "FailedOperation.Resource"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  FAILEDOPERATION_VPC = "FailedOperation.Vpc"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CLBFAILURE = "InternalError.CLBFailure"
//  INTERNALERROR_CAMNOAUTH = "InternalError.CamNoAuth"
//  INTERNALERROR_CREATEERROR = "InternalError.CreateError"
//  INTERNALERROR_DECODEERROR = "InternalError.DecodeError"
//  INTERNALERROR_EKSFAILURE = "InternalError.EKSFailure"
//  INTERNALERROR_GETCREDENTIAL = "InternalError.GetCredential"
//  INTERNALERROR_GETROLEERROR = "InternalError.GetRoleError"
//  INTERNALERROR_GOVERNANCEFAILURE = "InternalError.GovernanceFailure"
//  INTERNALERROR_HTTPSTATUSCODEERROR = "InternalError.HttpStatusCodeError"
//  INTERNALERROR_IOERROR = "InternalError.IOError"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INTERNALERROR_OPERATIONFAILED = "InternalError.OperationFailed"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_TKECLUSTERNOTREADY = "InternalError.TKEClusterNotReady"
//  INTERNALERROR_TKEFAILURE = "InternalError.TKEFailure"
//  INTERNALERROR_TAGFAILURE = "InternalError.TagFailure"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INTERNALERROR_UPDATEERROR = "InternalError.UpdateError"
//  INTERNALERROR_VPCFAILURE = "InternalError.VPCFailure"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARSEJSONERROR = "InvalidParameter.ParseJsonError"
//  INVALIDPARAMETERCOMBINATION = "InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ACTION = "InvalidParameterValue.Action"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_BUCKETNAME = "InvalidParameterValue.BucketName"
//  INVALIDPARAMETERVALUE_COS = "InvalidParameterValue.Cos"
//  INVALIDPARAMETERVALUE_CREATEERROR = "InvalidParameterValue.CreateError"
//  INVALIDPARAMETERVALUE_DESCRIPTION = "InvalidParameterValue.Description"
//  INVALIDPARAMETERVALUE_GATEWAYID = "InvalidParameterValue.GatewayId"
//  INVALIDPARAMETERVALUE_GATEWAYVERSION = "InvalidParameterValue.GatewayVersion"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_KEYNOTEXIST = "InvalidParameterValue.KeyNotExist"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  INVALIDPARAMETERVALUE_OBJECTNAME = "InvalidParameterValue.ObjectName"
//  INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"
//  INVALIDPARAMETERVALUE_PLUGIN = "InvalidParameterValue.Plugin"
//  INVALIDPARAMETERVALUE_QUERYERROR = "InvalidParameterValue.QueryError"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  INVALIDPARAMETERVALUE_SPECIFICATION = "InvalidParameterValue.Specification"
//  INVALIDPARAMETERVALUE_SYSTEMPARAMETER = "InvalidParameterValue.SystemParameter"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  INVALIDPARAMETERVALUE_UPDATEERROR = "InvalidParameterValue.UpdateError"
//  INVALIDPARAMETERVALUE_ZIPFILE = "InvalidParameterValue.ZipFile"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_LBDOMAINS = "LimitExceeded.LBDomains"
//  LIMITEXCEEDED_LIMITEXCEEDED = "LimitExceeded.LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_CREATEERROR = "MissingParameter.CreateError"
//  MISSINGPARAMETER_MISSPARAMETER = "MissingParameter.MissParameter"
//  MISSINGPARAMETER_UPDATEERROR = "MissingParameter.UpdateError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_LIMITEXCEEDED = "OperationDenied.LimitExceeded"
//  OPERATIONDENIED_OPERATIONDENIED = "OperationDenied.OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINUSE_GATEWAYSERVICEEXISTGROUPS = "ResourceInUse.GatewayServiceExistGroups"
//  RESOURCEINUSE_GATEWAYSERVICEEXISTROUTER = "ResourceInUse.GatewayServiceExistRouter"
//  RESOURCEINUSE_GATEWAYSERVICEGROUPEXISTUPSTREAM = "ResourceInUse.GatewayServiceGroupExistUpstream"
//  RESOURCEINUSE_GATEWAYSERVICESOURCEEXISTSERVICE = "ResourceInUse.GatewayServiceSourceExistService"
//  RESOURCEINUSE_INSTANCESEXISTEDINSERVICE = "ResourceInUse.InstancesExistedInService"
//  RESOURCEINUSE_SERVICESEXISTEDINNAMESPACE = "ResourceInUse.ServicesExistedInNamespace"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_CONFIGNOTEXIST = "ResourceNotFound.ConfigNotExist"
//  RESOURCENOTFOUND_FORBIDDEN = "ResourceNotFound.Forbidden"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_CAMPASSROLENOTEXIST = "UnauthorizedOperation.CamPassRoleNotExist"
//  UNAUTHORIZEDOPERATION_CLSNOTACTIVATED = "UnauthorizedOperation.ClsNotActivated"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_GOVERNANCENOTSUPPORTVISIBLE = "UnsupportedOperation.GovernanceNotSupportVisible"
func (c *Client) ModifyCloudNativeAPIGatewayMCPServerAuthWithContext(ctx context.Context, request *ModifyCloudNativeAPIGatewayMCPServerAuthRequest) (response *ModifyCloudNativeAPIGatewayMCPServerAuthResponse, err error) {
    if request == nil {
        request = NewModifyCloudNativeAPIGatewayMCPServerAuthRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cngw", APIVersion, "ModifyCloudNativeAPIGatewayMCPServerAuth")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyCloudNativeAPIGatewayMCPServerAuth require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyCloudNativeAPIGatewayMCPServerAuthResponse()
    err = c.Send(request, response)
    return
}

func NewModifyCloudNativeAPIGatewayMCPServerStatusRequest() (request *ModifyCloudNativeAPIGatewayMCPServerStatusRequest) {
    request = &ModifyCloudNativeAPIGatewayMCPServerStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cngw", APIVersion, "ModifyCloudNativeAPIGatewayMCPServerStatus")
    
    
    return
}

func NewModifyCloudNativeAPIGatewayMCPServerStatusResponse() (response *ModifyCloudNativeAPIGatewayMCPServerStatusResponse) {
    response = &ModifyCloudNativeAPIGatewayMCPServerStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyCloudNativeAPIGatewayMCPServerStatus
// 创建AI 网关MCP Server
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  FAILEDOPERATION_VPC = "FailedOperation.Vpc"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  INVALIDPARAMETERVALUE_SPECIFICATION = "InvalidParameterValue.Specification"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_CAMPASSROLENOTEXIST = "UnauthorizedOperation.CamPassRoleNotExist"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) ModifyCloudNativeAPIGatewayMCPServerStatus(request *ModifyCloudNativeAPIGatewayMCPServerStatusRequest) (response *ModifyCloudNativeAPIGatewayMCPServerStatusResponse, err error) {
    return c.ModifyCloudNativeAPIGatewayMCPServerStatusWithContext(context.Background(), request)
}

// ModifyCloudNativeAPIGatewayMCPServerStatus
// 创建AI 网关MCP Server
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  FAILEDOPERATION_VPC = "FailedOperation.Vpc"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  INVALIDPARAMETERVALUE_SPECIFICATION = "InvalidParameterValue.Specification"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_CAMPASSROLENOTEXIST = "UnauthorizedOperation.CamPassRoleNotExist"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) ModifyCloudNativeAPIGatewayMCPServerStatusWithContext(ctx context.Context, request *ModifyCloudNativeAPIGatewayMCPServerStatusRequest) (response *ModifyCloudNativeAPIGatewayMCPServerStatusResponse, err error) {
    if request == nil {
        request = NewModifyCloudNativeAPIGatewayMCPServerStatusRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cngw", APIVersion, "ModifyCloudNativeAPIGatewayMCPServerStatus")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyCloudNativeAPIGatewayMCPServerStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyCloudNativeAPIGatewayMCPServerStatusResponse()
    err = c.Send(request, response)
    return
}

func NewModifyCloudNativeAPIGatewayMCPToolRequest() (request *ModifyCloudNativeAPIGatewayMCPToolRequest) {
    request = &ModifyCloudNativeAPIGatewayMCPToolRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cngw", APIVersion, "ModifyCloudNativeAPIGatewayMCPTool")
    
    
    return
}

func NewModifyCloudNativeAPIGatewayMCPToolResponse() (response *ModifyCloudNativeAPIGatewayMCPToolResponse) {
    response = &ModifyCloudNativeAPIGatewayMCPToolResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyCloudNativeAPIGatewayMCPTool
// 修改AI网关MCP Tool
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  FAILEDOPERATION_VPC = "FailedOperation.Vpc"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  INVALIDPARAMETERVALUE_SPECIFICATION = "InvalidParameterValue.Specification"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_CAMPASSROLENOTEXIST = "UnauthorizedOperation.CamPassRoleNotExist"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) ModifyCloudNativeAPIGatewayMCPTool(request *ModifyCloudNativeAPIGatewayMCPToolRequest) (response *ModifyCloudNativeAPIGatewayMCPToolResponse, err error) {
    return c.ModifyCloudNativeAPIGatewayMCPToolWithContext(context.Background(), request)
}

// ModifyCloudNativeAPIGatewayMCPTool
// 修改AI网关MCP Tool
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  FAILEDOPERATION_VPC = "FailedOperation.Vpc"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  INVALIDPARAMETERVALUE_SPECIFICATION = "InvalidParameterValue.Specification"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_CAMPASSROLENOTEXIST = "UnauthorizedOperation.CamPassRoleNotExist"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) ModifyCloudNativeAPIGatewayMCPToolWithContext(ctx context.Context, request *ModifyCloudNativeAPIGatewayMCPToolRequest) (response *ModifyCloudNativeAPIGatewayMCPToolResponse, err error) {
    if request == nil {
        request = NewModifyCloudNativeAPIGatewayMCPToolRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cngw", APIVersion, "ModifyCloudNativeAPIGatewayMCPTool")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyCloudNativeAPIGatewayMCPTool require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyCloudNativeAPIGatewayMCPToolResponse()
    err = c.Send(request, response)
    return
}

func NewModifyCloudNativeAPIGatewayMCPToolACLRequest() (request *ModifyCloudNativeAPIGatewayMCPToolACLRequest) {
    request = &ModifyCloudNativeAPIGatewayMCPToolACLRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cngw", APIVersion, "ModifyCloudNativeAPIGatewayMCPToolACL")
    
    
    return
}

func NewModifyCloudNativeAPIGatewayMCPToolACLResponse() (response *ModifyCloudNativeAPIGatewayMCPToolACLResponse) {
    response = &ModifyCloudNativeAPIGatewayMCPToolACLResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyCloudNativeAPIGatewayMCPToolACL
// 修改 MCP Server Tool ACL
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SECRETIDNOTFOUND = "AuthFailure.SecretIdNotFound"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLS = "FailedOperation.Cls"
//  FAILEDOPERATION_CLSCONFIG = "FailedOperation.ClsConfig"
//  FAILEDOPERATION_COS = "FailedOperation.Cos"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_PLUGIN = "FailedOperation.Plugin"
//  FAILEDOPERATION_RESOURCE = "FailedOperation.Resource"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  FAILEDOPERATION_VPC = "FailedOperation.Vpc"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CLBFAILURE = "InternalError.CLBFailure"
//  INTERNALERROR_CAMNOAUTH = "InternalError.CamNoAuth"
//  INTERNALERROR_CREATEERROR = "InternalError.CreateError"
//  INTERNALERROR_DECODEERROR = "InternalError.DecodeError"
//  INTERNALERROR_EKSFAILURE = "InternalError.EKSFailure"
//  INTERNALERROR_GETCREDENTIAL = "InternalError.GetCredential"
//  INTERNALERROR_GETROLEERROR = "InternalError.GetRoleError"
//  INTERNALERROR_GOVERNANCEFAILURE = "InternalError.GovernanceFailure"
//  INTERNALERROR_HTTPSTATUSCODEERROR = "InternalError.HttpStatusCodeError"
//  INTERNALERROR_IOERROR = "InternalError.IOError"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INTERNALERROR_OPERATIONFAILED = "InternalError.OperationFailed"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_TKECLUSTERNOTREADY = "InternalError.TKEClusterNotReady"
//  INTERNALERROR_TKEFAILURE = "InternalError.TKEFailure"
//  INTERNALERROR_TAGFAILURE = "InternalError.TagFailure"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INTERNALERROR_UPDATEERROR = "InternalError.UpdateError"
//  INTERNALERROR_VPCFAILURE = "InternalError.VPCFailure"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARSEJSONERROR = "InvalidParameter.ParseJsonError"
//  INVALIDPARAMETERCOMBINATION = "InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ACTION = "InvalidParameterValue.Action"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_BUCKETNAME = "InvalidParameterValue.BucketName"
//  INVALIDPARAMETERVALUE_COS = "InvalidParameterValue.Cos"
//  INVALIDPARAMETERVALUE_CREATEERROR = "InvalidParameterValue.CreateError"
//  INVALIDPARAMETERVALUE_DESCRIPTION = "InvalidParameterValue.Description"
//  INVALIDPARAMETERVALUE_GATEWAYID = "InvalidParameterValue.GatewayId"
//  INVALIDPARAMETERVALUE_GATEWAYVERSION = "InvalidParameterValue.GatewayVersion"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_KEYNOTEXIST = "InvalidParameterValue.KeyNotExist"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  INVALIDPARAMETERVALUE_OBJECTNAME = "InvalidParameterValue.ObjectName"
//  INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"
//  INVALIDPARAMETERVALUE_PLUGIN = "InvalidParameterValue.Plugin"
//  INVALIDPARAMETERVALUE_QUERYERROR = "InvalidParameterValue.QueryError"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  INVALIDPARAMETERVALUE_SPECIFICATION = "InvalidParameterValue.Specification"
//  INVALIDPARAMETERVALUE_SYSTEMPARAMETER = "InvalidParameterValue.SystemParameter"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  INVALIDPARAMETERVALUE_UPDATEERROR = "InvalidParameterValue.UpdateError"
//  INVALIDPARAMETERVALUE_ZIPFILE = "InvalidParameterValue.ZipFile"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_LBDOMAINS = "LimitExceeded.LBDomains"
//  LIMITEXCEEDED_LIMITEXCEEDED = "LimitExceeded.LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_CREATEERROR = "MissingParameter.CreateError"
//  MISSINGPARAMETER_MISSPARAMETER = "MissingParameter.MissParameter"
//  MISSINGPARAMETER_UPDATEERROR = "MissingParameter.UpdateError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_LIMITEXCEEDED = "OperationDenied.LimitExceeded"
//  OPERATIONDENIED_OPERATIONDENIED = "OperationDenied.OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINUSE_GATEWAYSERVICEEXISTGROUPS = "ResourceInUse.GatewayServiceExistGroups"
//  RESOURCEINUSE_GATEWAYSERVICEEXISTROUTER = "ResourceInUse.GatewayServiceExistRouter"
//  RESOURCEINUSE_GATEWAYSERVICEGROUPEXISTUPSTREAM = "ResourceInUse.GatewayServiceGroupExistUpstream"
//  RESOURCEINUSE_GATEWAYSERVICESOURCEEXISTSERVICE = "ResourceInUse.GatewayServiceSourceExistService"
//  RESOURCEINUSE_INSTANCESEXISTEDINSERVICE = "ResourceInUse.InstancesExistedInService"
//  RESOURCEINUSE_SERVICESEXISTEDINNAMESPACE = "ResourceInUse.ServicesExistedInNamespace"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_CONFIGNOTEXIST = "ResourceNotFound.ConfigNotExist"
//  RESOURCENOTFOUND_FORBIDDEN = "ResourceNotFound.Forbidden"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_CAMPASSROLENOTEXIST = "UnauthorizedOperation.CamPassRoleNotExist"
//  UNAUTHORIZEDOPERATION_CLSNOTACTIVATED = "UnauthorizedOperation.ClsNotActivated"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_GOVERNANCENOTSUPPORTVISIBLE = "UnsupportedOperation.GovernanceNotSupportVisible"
func (c *Client) ModifyCloudNativeAPIGatewayMCPToolACL(request *ModifyCloudNativeAPIGatewayMCPToolACLRequest) (response *ModifyCloudNativeAPIGatewayMCPToolACLResponse, err error) {
    return c.ModifyCloudNativeAPIGatewayMCPToolACLWithContext(context.Background(), request)
}

// ModifyCloudNativeAPIGatewayMCPToolACL
// 修改 MCP Server Tool ACL
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SECRETIDNOTFOUND = "AuthFailure.SecretIdNotFound"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLS = "FailedOperation.Cls"
//  FAILEDOPERATION_CLSCONFIG = "FailedOperation.ClsConfig"
//  FAILEDOPERATION_COS = "FailedOperation.Cos"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_PLUGIN = "FailedOperation.Plugin"
//  FAILEDOPERATION_RESOURCE = "FailedOperation.Resource"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  FAILEDOPERATION_VPC = "FailedOperation.Vpc"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CLBFAILURE = "InternalError.CLBFailure"
//  INTERNALERROR_CAMNOAUTH = "InternalError.CamNoAuth"
//  INTERNALERROR_CREATEERROR = "InternalError.CreateError"
//  INTERNALERROR_DECODEERROR = "InternalError.DecodeError"
//  INTERNALERROR_EKSFAILURE = "InternalError.EKSFailure"
//  INTERNALERROR_GETCREDENTIAL = "InternalError.GetCredential"
//  INTERNALERROR_GETROLEERROR = "InternalError.GetRoleError"
//  INTERNALERROR_GOVERNANCEFAILURE = "InternalError.GovernanceFailure"
//  INTERNALERROR_HTTPSTATUSCODEERROR = "InternalError.HttpStatusCodeError"
//  INTERNALERROR_IOERROR = "InternalError.IOError"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INTERNALERROR_OPERATIONFAILED = "InternalError.OperationFailed"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_TKECLUSTERNOTREADY = "InternalError.TKEClusterNotReady"
//  INTERNALERROR_TKEFAILURE = "InternalError.TKEFailure"
//  INTERNALERROR_TAGFAILURE = "InternalError.TagFailure"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INTERNALERROR_UPDATEERROR = "InternalError.UpdateError"
//  INTERNALERROR_VPCFAILURE = "InternalError.VPCFailure"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARSEJSONERROR = "InvalidParameter.ParseJsonError"
//  INVALIDPARAMETERCOMBINATION = "InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ACTION = "InvalidParameterValue.Action"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_BUCKETNAME = "InvalidParameterValue.BucketName"
//  INVALIDPARAMETERVALUE_COS = "InvalidParameterValue.Cos"
//  INVALIDPARAMETERVALUE_CREATEERROR = "InvalidParameterValue.CreateError"
//  INVALIDPARAMETERVALUE_DESCRIPTION = "InvalidParameterValue.Description"
//  INVALIDPARAMETERVALUE_GATEWAYID = "InvalidParameterValue.GatewayId"
//  INVALIDPARAMETERVALUE_GATEWAYVERSION = "InvalidParameterValue.GatewayVersion"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_KEYNOTEXIST = "InvalidParameterValue.KeyNotExist"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  INVALIDPARAMETERVALUE_OBJECTNAME = "InvalidParameterValue.ObjectName"
//  INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"
//  INVALIDPARAMETERVALUE_PLUGIN = "InvalidParameterValue.Plugin"
//  INVALIDPARAMETERVALUE_QUERYERROR = "InvalidParameterValue.QueryError"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  INVALIDPARAMETERVALUE_SPECIFICATION = "InvalidParameterValue.Specification"
//  INVALIDPARAMETERVALUE_SYSTEMPARAMETER = "InvalidParameterValue.SystemParameter"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  INVALIDPARAMETERVALUE_UPDATEERROR = "InvalidParameterValue.UpdateError"
//  INVALIDPARAMETERVALUE_ZIPFILE = "InvalidParameterValue.ZipFile"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_LBDOMAINS = "LimitExceeded.LBDomains"
//  LIMITEXCEEDED_LIMITEXCEEDED = "LimitExceeded.LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_CREATEERROR = "MissingParameter.CreateError"
//  MISSINGPARAMETER_MISSPARAMETER = "MissingParameter.MissParameter"
//  MISSINGPARAMETER_UPDATEERROR = "MissingParameter.UpdateError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_LIMITEXCEEDED = "OperationDenied.LimitExceeded"
//  OPERATIONDENIED_OPERATIONDENIED = "OperationDenied.OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINUSE_GATEWAYSERVICEEXISTGROUPS = "ResourceInUse.GatewayServiceExistGroups"
//  RESOURCEINUSE_GATEWAYSERVICEEXISTROUTER = "ResourceInUse.GatewayServiceExistRouter"
//  RESOURCEINUSE_GATEWAYSERVICEGROUPEXISTUPSTREAM = "ResourceInUse.GatewayServiceGroupExistUpstream"
//  RESOURCEINUSE_GATEWAYSERVICESOURCEEXISTSERVICE = "ResourceInUse.GatewayServiceSourceExistService"
//  RESOURCEINUSE_INSTANCESEXISTEDINSERVICE = "ResourceInUse.InstancesExistedInService"
//  RESOURCEINUSE_SERVICESEXISTEDINNAMESPACE = "ResourceInUse.ServicesExistedInNamespace"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_CONFIGNOTEXIST = "ResourceNotFound.ConfigNotExist"
//  RESOURCENOTFOUND_FORBIDDEN = "ResourceNotFound.Forbidden"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_CAMPASSROLENOTEXIST = "UnauthorizedOperation.CamPassRoleNotExist"
//  UNAUTHORIZEDOPERATION_CLSNOTACTIVATED = "UnauthorizedOperation.ClsNotActivated"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_GOVERNANCENOTSUPPORTVISIBLE = "UnsupportedOperation.GovernanceNotSupportVisible"
func (c *Client) ModifyCloudNativeAPIGatewayMCPToolACLWithContext(ctx context.Context, request *ModifyCloudNativeAPIGatewayMCPToolACLRequest) (response *ModifyCloudNativeAPIGatewayMCPToolACLResponse, err error) {
    if request == nil {
        request = NewModifyCloudNativeAPIGatewayMCPToolACLRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cngw", APIVersion, "ModifyCloudNativeAPIGatewayMCPToolACL")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyCloudNativeAPIGatewayMCPToolACL require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyCloudNativeAPIGatewayMCPToolACLResponse()
    err = c.Send(request, response)
    return
}

func NewModifyCloudNativeAPIGatewayMCPToolStatusRequest() (request *ModifyCloudNativeAPIGatewayMCPToolStatusRequest) {
    request = &ModifyCloudNativeAPIGatewayMCPToolStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cngw", APIVersion, "ModifyCloudNativeAPIGatewayMCPToolStatus")
    
    
    return
}

func NewModifyCloudNativeAPIGatewayMCPToolStatusResponse() (response *ModifyCloudNativeAPIGatewayMCPToolStatusResponse) {
    response = &ModifyCloudNativeAPIGatewayMCPToolStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyCloudNativeAPIGatewayMCPToolStatus
// AI网关修改MCP Tool上下线状态
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  FAILEDOPERATION_VPC = "FailedOperation.Vpc"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  INVALIDPARAMETERVALUE_SPECIFICATION = "InvalidParameterValue.Specification"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_CAMPASSROLENOTEXIST = "UnauthorizedOperation.CamPassRoleNotExist"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) ModifyCloudNativeAPIGatewayMCPToolStatus(request *ModifyCloudNativeAPIGatewayMCPToolStatusRequest) (response *ModifyCloudNativeAPIGatewayMCPToolStatusResponse, err error) {
    return c.ModifyCloudNativeAPIGatewayMCPToolStatusWithContext(context.Background(), request)
}

// ModifyCloudNativeAPIGatewayMCPToolStatus
// AI网关修改MCP Tool上下线状态
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  FAILEDOPERATION_VPC = "FailedOperation.Vpc"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  INVALIDPARAMETERVALUE_SPECIFICATION = "InvalidParameterValue.Specification"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_CAMPASSROLENOTEXIST = "UnauthorizedOperation.CamPassRoleNotExist"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) ModifyCloudNativeAPIGatewayMCPToolStatusWithContext(ctx context.Context, request *ModifyCloudNativeAPIGatewayMCPToolStatusRequest) (response *ModifyCloudNativeAPIGatewayMCPToolStatusResponse, err error) {
    if request == nil {
        request = NewModifyCloudNativeAPIGatewayMCPToolStatusRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cngw", APIVersion, "ModifyCloudNativeAPIGatewayMCPToolStatus")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyCloudNativeAPIGatewayMCPToolStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyCloudNativeAPIGatewayMCPToolStatusResponse()
    err = c.Send(request, response)
    return
}

func NewModifyCloudNativeAPIGatewaySecretKeyRequest() (request *ModifyCloudNativeAPIGatewaySecretKeyRequest) {
    request = &ModifyCloudNativeAPIGatewaySecretKeyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cngw", APIVersion, "ModifyCloudNativeAPIGatewaySecretKey")
    
    
    return
}

func NewModifyCloudNativeAPIGatewaySecretKeyResponse() (response *ModifyCloudNativeAPIGatewaySecretKeyResponse) {
    response = &ModifyCloudNativeAPIGatewaySecretKeyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyCloudNativeAPIGatewaySecretKey
// 修改AI网关密钥
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  FAILEDOPERATION_VPC = "FailedOperation.Vpc"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  INVALIDPARAMETERVALUE_SPECIFICATION = "InvalidParameterValue.Specification"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_CAMPASSROLENOTEXIST = "UnauthorizedOperation.CamPassRoleNotExist"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) ModifyCloudNativeAPIGatewaySecretKey(request *ModifyCloudNativeAPIGatewaySecretKeyRequest) (response *ModifyCloudNativeAPIGatewaySecretKeyResponse, err error) {
    return c.ModifyCloudNativeAPIGatewaySecretKeyWithContext(context.Background(), request)
}

// ModifyCloudNativeAPIGatewaySecretKey
// 修改AI网关密钥
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  FAILEDOPERATION_VPC = "FailedOperation.Vpc"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  INVALIDPARAMETERVALUE_SPECIFICATION = "InvalidParameterValue.Specification"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_CAMPASSROLENOTEXIST = "UnauthorizedOperation.CamPassRoleNotExist"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) ModifyCloudNativeAPIGatewaySecretKeyWithContext(ctx context.Context, request *ModifyCloudNativeAPIGatewaySecretKeyRequest) (response *ModifyCloudNativeAPIGatewaySecretKeyResponse, err error) {
    if request == nil {
        request = NewModifyCloudNativeAPIGatewaySecretKeyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cngw", APIVersion, "ModifyCloudNativeAPIGatewaySecretKey")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyCloudNativeAPIGatewaySecretKey require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyCloudNativeAPIGatewaySecretKeyResponse()
    err = c.Send(request, response)
    return
}

func NewRemoveCloudNativeAPIGatewayConsumerGroupAuthRequest() (request *RemoveCloudNativeAPIGatewayConsumerGroupAuthRequest) {
    request = &RemoveCloudNativeAPIGatewayConsumerGroupAuthRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cngw", APIVersion, "RemoveCloudNativeAPIGatewayConsumerGroupAuth")
    
    
    return
}

func NewRemoveCloudNativeAPIGatewayConsumerGroupAuthResponse() (response *RemoveCloudNativeAPIGatewayConsumerGroupAuthResponse) {
    response = &RemoveCloudNativeAPIGatewayConsumerGroupAuthResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RemoveCloudNativeAPIGatewayConsumerGroupAuth
// 从资源（模型 API / MCP Server）移除消费者组授权。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  FAILEDOPERATION_VPC = "FailedOperation.Vpc"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  INVALIDPARAMETERVALUE_SPECIFICATION = "InvalidParameterValue.Specification"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_CAMPASSROLENOTEXIST = "UnauthorizedOperation.CamPassRoleNotExist"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) RemoveCloudNativeAPIGatewayConsumerGroupAuth(request *RemoveCloudNativeAPIGatewayConsumerGroupAuthRequest) (response *RemoveCloudNativeAPIGatewayConsumerGroupAuthResponse, err error) {
    return c.RemoveCloudNativeAPIGatewayConsumerGroupAuthWithContext(context.Background(), request)
}

// RemoveCloudNativeAPIGatewayConsumerGroupAuth
// 从资源（模型 API / MCP Server）移除消费者组授权。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  FAILEDOPERATION_VPC = "FailedOperation.Vpc"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  INVALIDPARAMETERVALUE_SPECIFICATION = "InvalidParameterValue.Specification"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_CAMPASSROLENOTEXIST = "UnauthorizedOperation.CamPassRoleNotExist"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) RemoveCloudNativeAPIGatewayConsumerGroupAuthWithContext(ctx context.Context, request *RemoveCloudNativeAPIGatewayConsumerGroupAuthRequest) (response *RemoveCloudNativeAPIGatewayConsumerGroupAuthResponse, err error) {
    if request == nil {
        request = NewRemoveCloudNativeAPIGatewayConsumerGroupAuthRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cngw", APIVersion, "RemoveCloudNativeAPIGatewayConsumerGroupAuth")
    
    if c.GetCredential() == nil {
        return nil, errors.New("RemoveCloudNativeAPIGatewayConsumerGroupAuth require credential")
    }

    request.SetContext(ctx)
    
    response = NewRemoveCloudNativeAPIGatewayConsumerGroupAuthResponse()
    err = c.Send(request, response)
    return
}

func NewRemoveCloudNativeAPIGatewayConsumerInGroupRequest() (request *RemoveCloudNativeAPIGatewayConsumerInGroupRequest) {
    request = &RemoveCloudNativeAPIGatewayConsumerInGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cngw", APIVersion, "RemoveCloudNativeAPIGatewayConsumerInGroup")
    
    
    return
}

func NewRemoveCloudNativeAPIGatewayConsumerInGroupResponse() (response *RemoveCloudNativeAPIGatewayConsumerInGroupResponse) {
    response = &RemoveCloudNativeAPIGatewayConsumerInGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RemoveCloudNativeAPIGatewayConsumerInGroup
// 将消费者从消费者组移除。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  FAILEDOPERATION_VPC = "FailedOperation.Vpc"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  INVALIDPARAMETERVALUE_SPECIFICATION = "InvalidParameterValue.Specification"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_CAMPASSROLENOTEXIST = "UnauthorizedOperation.CamPassRoleNotExist"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) RemoveCloudNativeAPIGatewayConsumerInGroup(request *RemoveCloudNativeAPIGatewayConsumerInGroupRequest) (response *RemoveCloudNativeAPIGatewayConsumerInGroupResponse, err error) {
    return c.RemoveCloudNativeAPIGatewayConsumerInGroupWithContext(context.Background(), request)
}

// RemoveCloudNativeAPIGatewayConsumerInGroup
// 将消费者从消费者组移除。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  FAILEDOPERATION_VPC = "FailedOperation.Vpc"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  INVALIDPARAMETERVALUE_SPECIFICATION = "InvalidParameterValue.Specification"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_CAMPASSROLENOTEXIST = "UnauthorizedOperation.CamPassRoleNotExist"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) RemoveCloudNativeAPIGatewayConsumerInGroupWithContext(ctx context.Context, request *RemoveCloudNativeAPIGatewayConsumerInGroupRequest) (response *RemoveCloudNativeAPIGatewayConsumerInGroupResponse, err error) {
    if request == nil {
        request = NewRemoveCloudNativeAPIGatewayConsumerInGroupRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cngw", APIVersion, "RemoveCloudNativeAPIGatewayConsumerInGroup")
    
    if c.GetCredential() == nil {
        return nil, errors.New("RemoveCloudNativeAPIGatewayConsumerInGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewRemoveCloudNativeAPIGatewayConsumerInGroupResponse()
    err = c.Send(request, response)
    return
}

func NewUnbindCloudNativeAPIGatewaySecretKeyRequest() (request *UnbindCloudNativeAPIGatewaySecretKeyRequest) {
    request = &UnbindCloudNativeAPIGatewaySecretKeyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cngw", APIVersion, "UnbindCloudNativeAPIGatewaySecretKey")
    
    
    return
}

func NewUnbindCloudNativeAPIGatewaySecretKeyResponse() (response *UnbindCloudNativeAPIGatewaySecretKeyResponse) {
    response = &UnbindCloudNativeAPIGatewaySecretKeyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UnbindCloudNativeAPIGatewaySecretKey
// 解绑密钥
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  FAILEDOPERATION_VPC = "FailedOperation.Vpc"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  INVALIDPARAMETERVALUE_SPECIFICATION = "InvalidParameterValue.Specification"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_CAMPASSROLENOTEXIST = "UnauthorizedOperation.CamPassRoleNotExist"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) UnbindCloudNativeAPIGatewaySecretKey(request *UnbindCloudNativeAPIGatewaySecretKeyRequest) (response *UnbindCloudNativeAPIGatewaySecretKeyResponse, err error) {
    return c.UnbindCloudNativeAPIGatewaySecretKeyWithContext(context.Background(), request)
}

// UnbindCloudNativeAPIGatewaySecretKey
// 解绑密钥
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  FAILEDOPERATION_VPC = "FailedOperation.Vpc"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  INVALIDPARAMETERVALUE_SPECIFICATION = "InvalidParameterValue.Specification"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_CAMPASSROLENOTEXIST = "UnauthorizedOperation.CamPassRoleNotExist"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) UnbindCloudNativeAPIGatewaySecretKeyWithContext(ctx context.Context, request *UnbindCloudNativeAPIGatewaySecretKeyRequest) (response *UnbindCloudNativeAPIGatewaySecretKeyResponse, err error) {
    if request == nil {
        request = NewUnbindCloudNativeAPIGatewaySecretKeyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cngw", APIVersion, "UnbindCloudNativeAPIGatewaySecretKey")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UnbindCloudNativeAPIGatewaySecretKey require credential")
    }

    request.SetContext(ctx)
    
    response = NewUnbindCloudNativeAPIGatewaySecretKeyResponse()
    err = c.Send(request, response)
    return
}
