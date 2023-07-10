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

package v20201207

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2020-12-07"

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


func NewCreateCloudNativeAPIGatewayCanaryRuleRequest() (request *CreateCloudNativeAPIGatewayCanaryRuleRequest) {
    request = &CreateCloudNativeAPIGatewayCanaryRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "CreateCloudNativeAPIGatewayCanaryRule")
    
    
    return
}

func NewCreateCloudNativeAPIGatewayCanaryRuleResponse() (response *CreateCloudNativeAPIGatewayCanaryRuleResponse) {
    response = &CreateCloudNativeAPIGatewayCanaryRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateCloudNativeAPIGatewayCanaryRule
// 创建云原生网关的灰度规则
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
//  INVALIDPARAMETERVALUE_CREATEERROR = "InvalidParameterValue.CreateError"
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
func (c *Client) CreateCloudNativeAPIGatewayCanaryRule(request *CreateCloudNativeAPIGatewayCanaryRuleRequest) (response *CreateCloudNativeAPIGatewayCanaryRuleResponse, err error) {
    return c.CreateCloudNativeAPIGatewayCanaryRuleWithContext(context.Background(), request)
}

// CreateCloudNativeAPIGatewayCanaryRule
// 创建云原生网关的灰度规则
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
//  INVALIDPARAMETERVALUE_CREATEERROR = "InvalidParameterValue.CreateError"
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
func (c *Client) CreateCloudNativeAPIGatewayCanaryRuleWithContext(ctx context.Context, request *CreateCloudNativeAPIGatewayCanaryRuleRequest) (response *CreateCloudNativeAPIGatewayCanaryRuleResponse, err error) {
    if request == nil {
        request = NewCreateCloudNativeAPIGatewayCanaryRuleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCloudNativeAPIGatewayCanaryRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateCloudNativeAPIGatewayCanaryRuleResponse()
    err = c.Send(request, response)
    return
}

func NewCreateCloudNativeAPIGatewayRouteRequest() (request *CreateCloudNativeAPIGatewayRouteRequest) {
    request = &CreateCloudNativeAPIGatewayRouteRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "CreateCloudNativeAPIGatewayRoute")
    
    
    return
}

func NewCreateCloudNativeAPIGatewayRouteResponse() (response *CreateCloudNativeAPIGatewayRouteResponse) {
    response = &CreateCloudNativeAPIGatewayRouteResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateCloudNativeAPIGatewayRoute
// 创建云原生网关路由
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
func (c *Client) CreateCloudNativeAPIGatewayRoute(request *CreateCloudNativeAPIGatewayRouteRequest) (response *CreateCloudNativeAPIGatewayRouteResponse, err error) {
    return c.CreateCloudNativeAPIGatewayRouteWithContext(context.Background(), request)
}

// CreateCloudNativeAPIGatewayRoute
// 创建云原生网关路由
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
func (c *Client) CreateCloudNativeAPIGatewayRouteWithContext(ctx context.Context, request *CreateCloudNativeAPIGatewayRouteRequest) (response *CreateCloudNativeAPIGatewayRouteResponse, err error) {
    if request == nil {
        request = NewCreateCloudNativeAPIGatewayRouteRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCloudNativeAPIGatewayRoute require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateCloudNativeAPIGatewayRouteResponse()
    err = c.Send(request, response)
    return
}

func NewCreateCloudNativeAPIGatewayRouteRateLimitRequest() (request *CreateCloudNativeAPIGatewayRouteRateLimitRequest) {
    request = &CreateCloudNativeAPIGatewayRouteRateLimitRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "CreateCloudNativeAPIGatewayRouteRateLimit")
    
    
    return
}

func NewCreateCloudNativeAPIGatewayRouteRateLimitResponse() (response *CreateCloudNativeAPIGatewayRouteRateLimitResponse) {
    response = &CreateCloudNativeAPIGatewayRouteRateLimitResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateCloudNativeAPIGatewayRouteRateLimit
// 创建云原生网关限流插件(路由)
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
func (c *Client) CreateCloudNativeAPIGatewayRouteRateLimit(request *CreateCloudNativeAPIGatewayRouteRateLimitRequest) (response *CreateCloudNativeAPIGatewayRouteRateLimitResponse, err error) {
    return c.CreateCloudNativeAPIGatewayRouteRateLimitWithContext(context.Background(), request)
}

// CreateCloudNativeAPIGatewayRouteRateLimit
// 创建云原生网关限流插件(路由)
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
func (c *Client) CreateCloudNativeAPIGatewayRouteRateLimitWithContext(ctx context.Context, request *CreateCloudNativeAPIGatewayRouteRateLimitRequest) (response *CreateCloudNativeAPIGatewayRouteRateLimitResponse, err error) {
    if request == nil {
        request = NewCreateCloudNativeAPIGatewayRouteRateLimitRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCloudNativeAPIGatewayRouteRateLimit require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateCloudNativeAPIGatewayRouteRateLimitResponse()
    err = c.Send(request, response)
    return
}

func NewCreateCloudNativeAPIGatewayServiceRequest() (request *CreateCloudNativeAPIGatewayServiceRequest) {
    request = &CreateCloudNativeAPIGatewayServiceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "CreateCloudNativeAPIGatewayService")
    
    
    return
}

func NewCreateCloudNativeAPIGatewayServiceResponse() (response *CreateCloudNativeAPIGatewayServiceResponse) {
    response = &CreateCloudNativeAPIGatewayServiceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateCloudNativeAPIGatewayService
// 创建云原生网关服务
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
func (c *Client) CreateCloudNativeAPIGatewayService(request *CreateCloudNativeAPIGatewayServiceRequest) (response *CreateCloudNativeAPIGatewayServiceResponse, err error) {
    return c.CreateCloudNativeAPIGatewayServiceWithContext(context.Background(), request)
}

// CreateCloudNativeAPIGatewayService
// 创建云原生网关服务
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
func (c *Client) CreateCloudNativeAPIGatewayServiceWithContext(ctx context.Context, request *CreateCloudNativeAPIGatewayServiceRequest) (response *CreateCloudNativeAPIGatewayServiceResponse, err error) {
    if request == nil {
        request = NewCreateCloudNativeAPIGatewayServiceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCloudNativeAPIGatewayService require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateCloudNativeAPIGatewayServiceResponse()
    err = c.Send(request, response)
    return
}

func NewCreateCloudNativeAPIGatewayServiceRateLimitRequest() (request *CreateCloudNativeAPIGatewayServiceRateLimitRequest) {
    request = &CreateCloudNativeAPIGatewayServiceRateLimitRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "CreateCloudNativeAPIGatewayServiceRateLimit")
    
    
    return
}

func NewCreateCloudNativeAPIGatewayServiceRateLimitResponse() (response *CreateCloudNativeAPIGatewayServiceRateLimitResponse) {
    response = &CreateCloudNativeAPIGatewayServiceRateLimitResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateCloudNativeAPIGatewayServiceRateLimit
// 创建云原生网关限流插件(服务)
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
//  MISSINGPARAMETER_MISSPARAMETER = "MissingParameter.MissParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_CAMPASSROLENOTEXIST = "UnauthorizedOperation.CamPassRoleNotExist"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) CreateCloudNativeAPIGatewayServiceRateLimit(request *CreateCloudNativeAPIGatewayServiceRateLimitRequest) (response *CreateCloudNativeAPIGatewayServiceRateLimitResponse, err error) {
    return c.CreateCloudNativeAPIGatewayServiceRateLimitWithContext(context.Background(), request)
}

// CreateCloudNativeAPIGatewayServiceRateLimit
// 创建云原生网关限流插件(服务)
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
//  MISSINGPARAMETER_MISSPARAMETER = "MissingParameter.MissParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_CAMPASSROLENOTEXIST = "UnauthorizedOperation.CamPassRoleNotExist"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) CreateCloudNativeAPIGatewayServiceRateLimitWithContext(ctx context.Context, request *CreateCloudNativeAPIGatewayServiceRateLimitRequest) (response *CreateCloudNativeAPIGatewayServiceRateLimitResponse, err error) {
    if request == nil {
        request = NewCreateCloudNativeAPIGatewayServiceRateLimitRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCloudNativeAPIGatewayServiceRateLimit require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateCloudNativeAPIGatewayServiceRateLimitResponse()
    err = c.Send(request, response)
    return
}

func NewCreateEngineRequest() (request *CreateEngineRequest) {
    request = &CreateEngineRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "CreateEngine")
    
    
    return
}

func NewCreateEngineResponse() (response *CreateEngineResponse) {
    response = &CreateEngineResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateEngine
// 创建引擎实例
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  INTERNALERROR_CREATEERROR = "InternalError.CreateError"
//  INTERNALERROR_GETCREDENTIAL = "InternalError.GetCredential"
//  INTERNALERROR_TKEFAILURE = "InternalError.TKEFailure"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INTERNALERROR_VPCFAILURE = "InternalError.VPCFailure"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_CREATEERROR = "InvalidParameterValue.CreateError"
//  INVALIDPARAMETERVALUE_QUERYERROR = "InvalidParameterValue.QueryError"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER_CREATEERROR = "MissingParameter.CreateError"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_CAMPASSROLENOTEXIST = "UnauthorizedOperation.CamPassRoleNotExist"
func (c *Client) CreateEngine(request *CreateEngineRequest) (response *CreateEngineResponse, err error) {
    return c.CreateEngineWithContext(context.Background(), request)
}

// CreateEngine
// 创建引擎实例
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  INTERNALERROR_CREATEERROR = "InternalError.CreateError"
//  INTERNALERROR_GETCREDENTIAL = "InternalError.GetCredential"
//  INTERNALERROR_TKEFAILURE = "InternalError.TKEFailure"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INTERNALERROR_VPCFAILURE = "InternalError.VPCFailure"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_CREATEERROR = "InvalidParameterValue.CreateError"
//  INVALIDPARAMETERVALUE_QUERYERROR = "InvalidParameterValue.QueryError"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER_CREATEERROR = "MissingParameter.CreateError"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_CAMPASSROLENOTEXIST = "UnauthorizedOperation.CamPassRoleNotExist"
func (c *Client) CreateEngineWithContext(ctx context.Context, request *CreateEngineRequest) (response *CreateEngineResponse, err error) {
    if request == nil {
        request = NewCreateEngineRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateEngine require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateEngineResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteCloudNativeAPIGatewayCanaryRuleRequest() (request *DeleteCloudNativeAPIGatewayCanaryRuleRequest) {
    request = &DeleteCloudNativeAPIGatewayCanaryRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "DeleteCloudNativeAPIGatewayCanaryRule")
    
    
    return
}

func NewDeleteCloudNativeAPIGatewayCanaryRuleResponse() (response *DeleteCloudNativeAPIGatewayCanaryRuleResponse) {
    response = &DeleteCloudNativeAPIGatewayCanaryRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteCloudNativeAPIGatewayCanaryRule
// 删除云原生网关的灰度规则
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
func (c *Client) DeleteCloudNativeAPIGatewayCanaryRule(request *DeleteCloudNativeAPIGatewayCanaryRuleRequest) (response *DeleteCloudNativeAPIGatewayCanaryRuleResponse, err error) {
    return c.DeleteCloudNativeAPIGatewayCanaryRuleWithContext(context.Background(), request)
}

// DeleteCloudNativeAPIGatewayCanaryRule
// 删除云原生网关的灰度规则
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
func (c *Client) DeleteCloudNativeAPIGatewayCanaryRuleWithContext(ctx context.Context, request *DeleteCloudNativeAPIGatewayCanaryRuleRequest) (response *DeleteCloudNativeAPIGatewayCanaryRuleResponse, err error) {
    if request == nil {
        request = NewDeleteCloudNativeAPIGatewayCanaryRuleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteCloudNativeAPIGatewayCanaryRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteCloudNativeAPIGatewayCanaryRuleResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteCloudNativeAPIGatewayRouteRequest() (request *DeleteCloudNativeAPIGatewayRouteRequest) {
    request = &DeleteCloudNativeAPIGatewayRouteRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "DeleteCloudNativeAPIGatewayRoute")
    
    
    return
}

func NewDeleteCloudNativeAPIGatewayRouteResponse() (response *DeleteCloudNativeAPIGatewayRouteResponse) {
    response = &DeleteCloudNativeAPIGatewayRouteResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteCloudNativeAPIGatewayRoute
// 删除云原生网关路由
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
func (c *Client) DeleteCloudNativeAPIGatewayRoute(request *DeleteCloudNativeAPIGatewayRouteRequest) (response *DeleteCloudNativeAPIGatewayRouteResponse, err error) {
    return c.DeleteCloudNativeAPIGatewayRouteWithContext(context.Background(), request)
}

// DeleteCloudNativeAPIGatewayRoute
// 删除云原生网关路由
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
func (c *Client) DeleteCloudNativeAPIGatewayRouteWithContext(ctx context.Context, request *DeleteCloudNativeAPIGatewayRouteRequest) (response *DeleteCloudNativeAPIGatewayRouteResponse, err error) {
    if request == nil {
        request = NewDeleteCloudNativeAPIGatewayRouteRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteCloudNativeAPIGatewayRoute require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteCloudNativeAPIGatewayRouteResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteCloudNativeAPIGatewayRouteRateLimitRequest() (request *DeleteCloudNativeAPIGatewayRouteRateLimitRequest) {
    request = &DeleteCloudNativeAPIGatewayRouteRateLimitRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "DeleteCloudNativeAPIGatewayRouteRateLimit")
    
    
    return
}

func NewDeleteCloudNativeAPIGatewayRouteRateLimitResponse() (response *DeleteCloudNativeAPIGatewayRouteRateLimitResponse) {
    response = &DeleteCloudNativeAPIGatewayRouteRateLimitResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteCloudNativeAPIGatewayRouteRateLimit
// 删除云原生网关的限流插件(路由)
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
func (c *Client) DeleteCloudNativeAPIGatewayRouteRateLimit(request *DeleteCloudNativeAPIGatewayRouteRateLimitRequest) (response *DeleteCloudNativeAPIGatewayRouteRateLimitResponse, err error) {
    return c.DeleteCloudNativeAPIGatewayRouteRateLimitWithContext(context.Background(), request)
}

// DeleteCloudNativeAPIGatewayRouteRateLimit
// 删除云原生网关的限流插件(路由)
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
func (c *Client) DeleteCloudNativeAPIGatewayRouteRateLimitWithContext(ctx context.Context, request *DeleteCloudNativeAPIGatewayRouteRateLimitRequest) (response *DeleteCloudNativeAPIGatewayRouteRateLimitResponse, err error) {
    if request == nil {
        request = NewDeleteCloudNativeAPIGatewayRouteRateLimitRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteCloudNativeAPIGatewayRouteRateLimit require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteCloudNativeAPIGatewayRouteRateLimitResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteCloudNativeAPIGatewayServiceRequest() (request *DeleteCloudNativeAPIGatewayServiceRequest) {
    request = &DeleteCloudNativeAPIGatewayServiceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "DeleteCloudNativeAPIGatewayService")
    
    
    return
}

func NewDeleteCloudNativeAPIGatewayServiceResponse() (response *DeleteCloudNativeAPIGatewayServiceResponse) {
    response = &DeleteCloudNativeAPIGatewayServiceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteCloudNativeAPIGatewayService
// 删除云原生网关服务
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
func (c *Client) DeleteCloudNativeAPIGatewayService(request *DeleteCloudNativeAPIGatewayServiceRequest) (response *DeleteCloudNativeAPIGatewayServiceResponse, err error) {
    return c.DeleteCloudNativeAPIGatewayServiceWithContext(context.Background(), request)
}

// DeleteCloudNativeAPIGatewayService
// 删除云原生网关服务
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
func (c *Client) DeleteCloudNativeAPIGatewayServiceWithContext(ctx context.Context, request *DeleteCloudNativeAPIGatewayServiceRequest) (response *DeleteCloudNativeAPIGatewayServiceResponse, err error) {
    if request == nil {
        request = NewDeleteCloudNativeAPIGatewayServiceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteCloudNativeAPIGatewayService require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteCloudNativeAPIGatewayServiceResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteCloudNativeAPIGatewayServiceRateLimitRequest() (request *DeleteCloudNativeAPIGatewayServiceRateLimitRequest) {
    request = &DeleteCloudNativeAPIGatewayServiceRateLimitRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "DeleteCloudNativeAPIGatewayServiceRateLimit")
    
    
    return
}

func NewDeleteCloudNativeAPIGatewayServiceRateLimitResponse() (response *DeleteCloudNativeAPIGatewayServiceRateLimitResponse) {
    response = &DeleteCloudNativeAPIGatewayServiceRateLimitResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteCloudNativeAPIGatewayServiceRateLimit
// 删除云原生网关的限流插件(服务)
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
func (c *Client) DeleteCloudNativeAPIGatewayServiceRateLimit(request *DeleteCloudNativeAPIGatewayServiceRateLimitRequest) (response *DeleteCloudNativeAPIGatewayServiceRateLimitResponse, err error) {
    return c.DeleteCloudNativeAPIGatewayServiceRateLimitWithContext(context.Background(), request)
}

// DeleteCloudNativeAPIGatewayServiceRateLimit
// 删除云原生网关的限流插件(服务)
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
func (c *Client) DeleteCloudNativeAPIGatewayServiceRateLimitWithContext(ctx context.Context, request *DeleteCloudNativeAPIGatewayServiceRateLimitRequest) (response *DeleteCloudNativeAPIGatewayServiceRateLimitResponse, err error) {
    if request == nil {
        request = NewDeleteCloudNativeAPIGatewayServiceRateLimitRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteCloudNativeAPIGatewayServiceRateLimit require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteCloudNativeAPIGatewayServiceRateLimitResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteEngineRequest() (request *DeleteEngineRequest) {
    request = &DeleteEngineRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "DeleteEngine")
    
    
    return
}

func NewDeleteEngineResponse() (response *DeleteEngineResponse) {
    response = &DeleteEngineResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteEngine
// 删除引擎实例
//
// 可能返回的错误码:
//  INTERNALERROR_OPERATIONFAILED = "InternalError.OperationFailed"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INTERNALERROR_UPDATEERROR = "InternalError.UpdateError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_UPDATEERROR = "InvalidParameterValue.UpdateError"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteEngine(request *DeleteEngineRequest) (response *DeleteEngineResponse, err error) {
    return c.DeleteEngineWithContext(context.Background(), request)
}

// DeleteEngine
// 删除引擎实例
//
// 可能返回的错误码:
//  INTERNALERROR_OPERATIONFAILED = "InternalError.OperationFailed"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INTERNALERROR_UPDATEERROR = "InternalError.UpdateError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_UPDATEERROR = "InvalidParameterValue.UpdateError"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteEngineWithContext(ctx context.Context, request *DeleteEngineRequest) (response *DeleteEngineResponse, err error) {
    if request == nil {
        request = NewDeleteEngineRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteEngine require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteEngineResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCloudNativeAPIGatewayCanaryRulesRequest() (request *DescribeCloudNativeAPIGatewayCanaryRulesRequest) {
    request = &DescribeCloudNativeAPIGatewayCanaryRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "DescribeCloudNativeAPIGatewayCanaryRules")
    
    
    return
}

func NewDescribeCloudNativeAPIGatewayCanaryRulesResponse() (response *DescribeCloudNativeAPIGatewayCanaryRulesResponse) {
    response = &DescribeCloudNativeAPIGatewayCanaryRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCloudNativeAPIGatewayCanaryRules
// 查询云原生网关灰度规则列表
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
func (c *Client) DescribeCloudNativeAPIGatewayCanaryRules(request *DescribeCloudNativeAPIGatewayCanaryRulesRequest) (response *DescribeCloudNativeAPIGatewayCanaryRulesResponse, err error) {
    return c.DescribeCloudNativeAPIGatewayCanaryRulesWithContext(context.Background(), request)
}

// DescribeCloudNativeAPIGatewayCanaryRules
// 查询云原生网关灰度规则列表
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
func (c *Client) DescribeCloudNativeAPIGatewayCanaryRulesWithContext(ctx context.Context, request *DescribeCloudNativeAPIGatewayCanaryRulesRequest) (response *DescribeCloudNativeAPIGatewayCanaryRulesResponse, err error) {
    if request == nil {
        request = NewDescribeCloudNativeAPIGatewayCanaryRulesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCloudNativeAPIGatewayCanaryRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCloudNativeAPIGatewayCanaryRulesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCloudNativeAPIGatewayNodesRequest() (request *DescribeCloudNativeAPIGatewayNodesRequest) {
    request = &DescribeCloudNativeAPIGatewayNodesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "DescribeCloudNativeAPIGatewayNodes")
    
    
    return
}

func NewDescribeCloudNativeAPIGatewayNodesResponse() (response *DescribeCloudNativeAPIGatewayNodesResponse) {
    response = &DescribeCloudNativeAPIGatewayNodesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCloudNativeAPIGatewayNodes
// 获取云原生网关节点列表
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETERVALUE_ACTION = "InvalidParameterValue.Action"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_QUERYERROR = "InvalidParameterValue.QueryError"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
func (c *Client) DescribeCloudNativeAPIGatewayNodes(request *DescribeCloudNativeAPIGatewayNodesRequest) (response *DescribeCloudNativeAPIGatewayNodesResponse, err error) {
    return c.DescribeCloudNativeAPIGatewayNodesWithContext(context.Background(), request)
}

// DescribeCloudNativeAPIGatewayNodes
// 获取云原生网关节点列表
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETERVALUE_ACTION = "InvalidParameterValue.Action"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_QUERYERROR = "InvalidParameterValue.QueryError"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
func (c *Client) DescribeCloudNativeAPIGatewayNodesWithContext(ctx context.Context, request *DescribeCloudNativeAPIGatewayNodesRequest) (response *DescribeCloudNativeAPIGatewayNodesResponse, err error) {
    if request == nil {
        request = NewDescribeCloudNativeAPIGatewayNodesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCloudNativeAPIGatewayNodes require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCloudNativeAPIGatewayNodesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCloudNativeAPIGatewayRouteRateLimitRequest() (request *DescribeCloudNativeAPIGatewayRouteRateLimitRequest) {
    request = &DescribeCloudNativeAPIGatewayRouteRateLimitRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "DescribeCloudNativeAPIGatewayRouteRateLimit")
    
    
    return
}

func NewDescribeCloudNativeAPIGatewayRouteRateLimitResponse() (response *DescribeCloudNativeAPIGatewayRouteRateLimitResponse) {
    response = &DescribeCloudNativeAPIGatewayRouteRateLimitResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCloudNativeAPIGatewayRouteRateLimit
// 查询云原生网关的限流插件(路由)
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
func (c *Client) DescribeCloudNativeAPIGatewayRouteRateLimit(request *DescribeCloudNativeAPIGatewayRouteRateLimitRequest) (response *DescribeCloudNativeAPIGatewayRouteRateLimitResponse, err error) {
    return c.DescribeCloudNativeAPIGatewayRouteRateLimitWithContext(context.Background(), request)
}

// DescribeCloudNativeAPIGatewayRouteRateLimit
// 查询云原生网关的限流插件(路由)
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
func (c *Client) DescribeCloudNativeAPIGatewayRouteRateLimitWithContext(ctx context.Context, request *DescribeCloudNativeAPIGatewayRouteRateLimitRequest) (response *DescribeCloudNativeAPIGatewayRouteRateLimitResponse, err error) {
    if request == nil {
        request = NewDescribeCloudNativeAPIGatewayRouteRateLimitRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCloudNativeAPIGatewayRouteRateLimit require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCloudNativeAPIGatewayRouteRateLimitResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCloudNativeAPIGatewayRoutesRequest() (request *DescribeCloudNativeAPIGatewayRoutesRequest) {
    request = &DescribeCloudNativeAPIGatewayRoutesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "DescribeCloudNativeAPIGatewayRoutes")
    
    
    return
}

func NewDescribeCloudNativeAPIGatewayRoutesResponse() (response *DescribeCloudNativeAPIGatewayRoutesResponse) {
    response = &DescribeCloudNativeAPIGatewayRoutesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCloudNativeAPIGatewayRoutes
// 查询云原生网关路由列表
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
func (c *Client) DescribeCloudNativeAPIGatewayRoutes(request *DescribeCloudNativeAPIGatewayRoutesRequest) (response *DescribeCloudNativeAPIGatewayRoutesResponse, err error) {
    return c.DescribeCloudNativeAPIGatewayRoutesWithContext(context.Background(), request)
}

// DescribeCloudNativeAPIGatewayRoutes
// 查询云原生网关路由列表
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
func (c *Client) DescribeCloudNativeAPIGatewayRoutesWithContext(ctx context.Context, request *DescribeCloudNativeAPIGatewayRoutesRequest) (response *DescribeCloudNativeAPIGatewayRoutesResponse, err error) {
    if request == nil {
        request = NewDescribeCloudNativeAPIGatewayRoutesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCloudNativeAPIGatewayRoutes require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCloudNativeAPIGatewayRoutesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCloudNativeAPIGatewayServiceRateLimitRequest() (request *DescribeCloudNativeAPIGatewayServiceRateLimitRequest) {
    request = &DescribeCloudNativeAPIGatewayServiceRateLimitRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "DescribeCloudNativeAPIGatewayServiceRateLimit")
    
    
    return
}

func NewDescribeCloudNativeAPIGatewayServiceRateLimitResponse() (response *DescribeCloudNativeAPIGatewayServiceRateLimitResponse) {
    response = &DescribeCloudNativeAPIGatewayServiceRateLimitResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCloudNativeAPIGatewayServiceRateLimit
// 查询云原生网关的限流插件(服务)
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
func (c *Client) DescribeCloudNativeAPIGatewayServiceRateLimit(request *DescribeCloudNativeAPIGatewayServiceRateLimitRequest) (response *DescribeCloudNativeAPIGatewayServiceRateLimitResponse, err error) {
    return c.DescribeCloudNativeAPIGatewayServiceRateLimitWithContext(context.Background(), request)
}

// DescribeCloudNativeAPIGatewayServiceRateLimit
// 查询云原生网关的限流插件(服务)
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
func (c *Client) DescribeCloudNativeAPIGatewayServiceRateLimitWithContext(ctx context.Context, request *DescribeCloudNativeAPIGatewayServiceRateLimitRequest) (response *DescribeCloudNativeAPIGatewayServiceRateLimitResponse, err error) {
    if request == nil {
        request = NewDescribeCloudNativeAPIGatewayServiceRateLimitRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCloudNativeAPIGatewayServiceRateLimit require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCloudNativeAPIGatewayServiceRateLimitResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCloudNativeAPIGatewayServicesRequest() (request *DescribeCloudNativeAPIGatewayServicesRequest) {
    request = &DescribeCloudNativeAPIGatewayServicesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "DescribeCloudNativeAPIGatewayServices")
    
    
    return
}

func NewDescribeCloudNativeAPIGatewayServicesResponse() (response *DescribeCloudNativeAPIGatewayServicesResponse) {
    response = &DescribeCloudNativeAPIGatewayServicesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCloudNativeAPIGatewayServices
// 查询云原生网关服务列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  FAILEDOPERATION_VPC = "FailedOperation.Vpc"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_IOERROR = "InternalError.IOError"
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
func (c *Client) DescribeCloudNativeAPIGatewayServices(request *DescribeCloudNativeAPIGatewayServicesRequest) (response *DescribeCloudNativeAPIGatewayServicesResponse, err error) {
    return c.DescribeCloudNativeAPIGatewayServicesWithContext(context.Background(), request)
}

// DescribeCloudNativeAPIGatewayServices
// 查询云原生网关服务列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  FAILEDOPERATION_VPC = "FailedOperation.Vpc"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_IOERROR = "InternalError.IOError"
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
func (c *Client) DescribeCloudNativeAPIGatewayServicesWithContext(ctx context.Context, request *DescribeCloudNativeAPIGatewayServicesRequest) (response *DescribeCloudNativeAPIGatewayServicesResponse, err error) {
    if request == nil {
        request = NewDescribeCloudNativeAPIGatewayServicesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCloudNativeAPIGatewayServices require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCloudNativeAPIGatewayServicesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeNacosReplicasRequest() (request *DescribeNacosReplicasRequest) {
    request = &DescribeNacosReplicasRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "DescribeNacosReplicas")
    
    
    return
}

func NewDescribeNacosReplicasResponse() (response *DescribeNacosReplicasResponse) {
    response = &DescribeNacosReplicasResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeNacosReplicas
// 查询Nacos类型引擎实例副本信息
//
// 可能返回的错误码:
//  INTERNALERROR_GETCREDENTIAL = "InternalError.GetCredential"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_QUERYERROR = "InvalidParameterValue.QueryError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeNacosReplicas(request *DescribeNacosReplicasRequest) (response *DescribeNacosReplicasResponse, err error) {
    return c.DescribeNacosReplicasWithContext(context.Background(), request)
}

// DescribeNacosReplicas
// 查询Nacos类型引擎实例副本信息
//
// 可能返回的错误码:
//  INTERNALERROR_GETCREDENTIAL = "InternalError.GetCredential"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_QUERYERROR = "InvalidParameterValue.QueryError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeNacosReplicasWithContext(ctx context.Context, request *DescribeNacosReplicasRequest) (response *DescribeNacosReplicasResponse, err error) {
    if request == nil {
        request = NewDescribeNacosReplicasRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeNacosReplicas require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeNacosReplicasResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeNacosServerInterfacesRequest() (request *DescribeNacosServerInterfacesRequest) {
    request = &DescribeNacosServerInterfacesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "DescribeNacosServerInterfaces")
    
    
    return
}

func NewDescribeNacosServerInterfacesResponse() (response *DescribeNacosServerInterfacesResponse) {
    response = &DescribeNacosServerInterfacesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeNacosServerInterfaces
// 查询nacos服务接口列表
//
// 可能返回的错误码:
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_QUERYERROR = "InvalidParameterValue.QueryError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeNacosServerInterfaces(request *DescribeNacosServerInterfacesRequest) (response *DescribeNacosServerInterfacesResponse, err error) {
    return c.DescribeNacosServerInterfacesWithContext(context.Background(), request)
}

// DescribeNacosServerInterfaces
// 查询nacos服务接口列表
//
// 可能返回的错误码:
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_QUERYERROR = "InvalidParameterValue.QueryError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeNacosServerInterfacesWithContext(ctx context.Context, request *DescribeNacosServerInterfacesRequest) (response *DescribeNacosServerInterfacesResponse, err error) {
    if request == nil {
        request = NewDescribeNacosServerInterfacesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeNacosServerInterfaces require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeNacosServerInterfacesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeOneCloudNativeAPIGatewayServiceRequest() (request *DescribeOneCloudNativeAPIGatewayServiceRequest) {
    request = &DescribeOneCloudNativeAPIGatewayServiceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "DescribeOneCloudNativeAPIGatewayService")
    
    
    return
}

func NewDescribeOneCloudNativeAPIGatewayServiceResponse() (response *DescribeOneCloudNativeAPIGatewayServiceResponse) {
    response = &DescribeOneCloudNativeAPIGatewayServiceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeOneCloudNativeAPIGatewayService
// 获取云原生网关服务详情
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
func (c *Client) DescribeOneCloudNativeAPIGatewayService(request *DescribeOneCloudNativeAPIGatewayServiceRequest) (response *DescribeOneCloudNativeAPIGatewayServiceResponse, err error) {
    return c.DescribeOneCloudNativeAPIGatewayServiceWithContext(context.Background(), request)
}

// DescribeOneCloudNativeAPIGatewayService
// 获取云原生网关服务详情
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
func (c *Client) DescribeOneCloudNativeAPIGatewayServiceWithContext(ctx context.Context, request *DescribeOneCloudNativeAPIGatewayServiceRequest) (response *DescribeOneCloudNativeAPIGatewayServiceResponse, err error) {
    if request == nil {
        request = NewDescribeOneCloudNativeAPIGatewayServiceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeOneCloudNativeAPIGatewayService require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeOneCloudNativeAPIGatewayServiceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSREInstanceAccessAddressRequest() (request *DescribeSREInstanceAccessAddressRequest) {
    request = &DescribeSREInstanceAccessAddressRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "DescribeSREInstanceAccessAddress")
    
    
    return
}

func NewDescribeSREInstanceAccessAddressResponse() (response *DescribeSREInstanceAccessAddressResponse) {
    response = &DescribeSREInstanceAccessAddressResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeSREInstanceAccessAddress
// 查询引擎实例访问地址
//
// 可能返回的错误码:
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_QUERYERROR = "InvalidParameterValue.QueryError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeSREInstanceAccessAddress(request *DescribeSREInstanceAccessAddressRequest) (response *DescribeSREInstanceAccessAddressResponse, err error) {
    return c.DescribeSREInstanceAccessAddressWithContext(context.Background(), request)
}

// DescribeSREInstanceAccessAddress
// 查询引擎实例访问地址
//
// 可能返回的错误码:
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_QUERYERROR = "InvalidParameterValue.QueryError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeSREInstanceAccessAddressWithContext(ctx context.Context, request *DescribeSREInstanceAccessAddressRequest) (response *DescribeSREInstanceAccessAddressResponse, err error) {
    if request == nil {
        request = NewDescribeSREInstanceAccessAddressRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSREInstanceAccessAddress require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSREInstanceAccessAddressResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSREInstancesRequest() (request *DescribeSREInstancesRequest) {
    request = &DescribeSREInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "DescribeSREInstances")
    
    
    return
}

func NewDescribeSREInstancesResponse() (response *DescribeSREInstancesResponse) {
    response = &DescribeSREInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeSREInstances
// 用于查询引擎实例列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_GETROLEERROR = "InternalError.GetRoleError"
//  INTERNALERROR_HTTPSTATUSCODEERROR = "InternalError.HttpStatusCodeError"
//  INTERNALERROR_IOERROR = "InternalError.IOError"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_TAGFAILURE = "InternalError.TagFailure"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_QUERYERROR = "InvalidParameterValue.QueryError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
func (c *Client) DescribeSREInstances(request *DescribeSREInstancesRequest) (response *DescribeSREInstancesResponse, err error) {
    return c.DescribeSREInstancesWithContext(context.Background(), request)
}

// DescribeSREInstances
// 用于查询引擎实例列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_GETROLEERROR = "InternalError.GetRoleError"
//  INTERNALERROR_HTTPSTATUSCODEERROR = "InternalError.HttpStatusCodeError"
//  INTERNALERROR_IOERROR = "InternalError.IOError"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_TAGFAILURE = "InternalError.TagFailure"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_QUERYERROR = "InvalidParameterValue.QueryError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
func (c *Client) DescribeSREInstancesWithContext(ctx context.Context, request *DescribeSREInstancesRequest) (response *DescribeSREInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeSREInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSREInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSREInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeZookeeperReplicasRequest() (request *DescribeZookeeperReplicasRequest) {
    request = &DescribeZookeeperReplicasRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "DescribeZookeeperReplicas")
    
    
    return
}

func NewDescribeZookeeperReplicasResponse() (response *DescribeZookeeperReplicasResponse) {
    response = &DescribeZookeeperReplicasResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeZookeeperReplicas
// 查询Zookeeper类型注册引擎实例副本信息
//
// 可能返回的错误码:
//  INTERNALERROR_GETCREDENTIAL = "InternalError.GetCredential"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"
//  INVALIDPARAMETERVALUE_QUERYERROR = "InvalidParameterValue.QueryError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeZookeeperReplicas(request *DescribeZookeeperReplicasRequest) (response *DescribeZookeeperReplicasResponse, err error) {
    return c.DescribeZookeeperReplicasWithContext(context.Background(), request)
}

// DescribeZookeeperReplicas
// 查询Zookeeper类型注册引擎实例副本信息
//
// 可能返回的错误码:
//  INTERNALERROR_GETCREDENTIAL = "InternalError.GetCredential"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"
//  INVALIDPARAMETERVALUE_QUERYERROR = "InvalidParameterValue.QueryError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeZookeeperReplicasWithContext(ctx context.Context, request *DescribeZookeeperReplicasRequest) (response *DescribeZookeeperReplicasResponse, err error) {
    if request == nil {
        request = NewDescribeZookeeperReplicasRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeZookeeperReplicas require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeZookeeperReplicasResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeZookeeperServerInterfacesRequest() (request *DescribeZookeeperServerInterfacesRequest) {
    request = &DescribeZookeeperServerInterfacesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "DescribeZookeeperServerInterfaces")
    
    
    return
}

func NewDescribeZookeeperServerInterfacesResponse() (response *DescribeZookeeperServerInterfacesResponse) {
    response = &DescribeZookeeperServerInterfacesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeZookeeperServerInterfaces
// 查询zookeeper服务接口列表
//
// 可能返回的错误码:
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_QUERYERROR = "InvalidParameterValue.QueryError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeZookeeperServerInterfaces(request *DescribeZookeeperServerInterfacesRequest) (response *DescribeZookeeperServerInterfacesResponse, err error) {
    return c.DescribeZookeeperServerInterfacesWithContext(context.Background(), request)
}

// DescribeZookeeperServerInterfaces
// 查询zookeeper服务接口列表
//
// 可能返回的错误码:
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_QUERYERROR = "InvalidParameterValue.QueryError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeZookeeperServerInterfacesWithContext(ctx context.Context, request *DescribeZookeeperServerInterfacesRequest) (response *DescribeZookeeperServerInterfacesResponse, err error) {
    if request == nil {
        request = NewDescribeZookeeperServerInterfacesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeZookeeperServerInterfaces require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeZookeeperServerInterfacesResponse()
    err = c.Send(request, response)
    return
}

func NewModifyCloudNativeAPIGatewayCanaryRuleRequest() (request *ModifyCloudNativeAPIGatewayCanaryRuleRequest) {
    request = &ModifyCloudNativeAPIGatewayCanaryRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "ModifyCloudNativeAPIGatewayCanaryRule")
    
    
    return
}

func NewModifyCloudNativeAPIGatewayCanaryRuleResponse() (response *ModifyCloudNativeAPIGatewayCanaryRuleResponse) {
    response = &ModifyCloudNativeAPIGatewayCanaryRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyCloudNativeAPIGatewayCanaryRule
// 修改云原生网关的灰度规则
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
//  INVALIDPARAMETERVALUE_UPDATEERROR = "InvalidParameterValue.UpdateError"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_CAMPASSROLENOTEXIST = "UnauthorizedOperation.CamPassRoleNotExist"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) ModifyCloudNativeAPIGatewayCanaryRule(request *ModifyCloudNativeAPIGatewayCanaryRuleRequest) (response *ModifyCloudNativeAPIGatewayCanaryRuleResponse, err error) {
    return c.ModifyCloudNativeAPIGatewayCanaryRuleWithContext(context.Background(), request)
}

// ModifyCloudNativeAPIGatewayCanaryRule
// 修改云原生网关的灰度规则
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
//  INVALIDPARAMETERVALUE_UPDATEERROR = "InvalidParameterValue.UpdateError"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_CAMPASSROLENOTEXIST = "UnauthorizedOperation.CamPassRoleNotExist"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) ModifyCloudNativeAPIGatewayCanaryRuleWithContext(ctx context.Context, request *ModifyCloudNativeAPIGatewayCanaryRuleRequest) (response *ModifyCloudNativeAPIGatewayCanaryRuleResponse, err error) {
    if request == nil {
        request = NewModifyCloudNativeAPIGatewayCanaryRuleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyCloudNativeAPIGatewayCanaryRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyCloudNativeAPIGatewayCanaryRuleResponse()
    err = c.Send(request, response)
    return
}

func NewModifyCloudNativeAPIGatewayRouteRequest() (request *ModifyCloudNativeAPIGatewayRouteRequest) {
    request = &ModifyCloudNativeAPIGatewayRouteRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "ModifyCloudNativeAPIGatewayRoute")
    
    
    return
}

func NewModifyCloudNativeAPIGatewayRouteResponse() (response *ModifyCloudNativeAPIGatewayRouteResponse) {
    response = &ModifyCloudNativeAPIGatewayRouteResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyCloudNativeAPIGatewayRoute
// 修改云原生网关路由
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
func (c *Client) ModifyCloudNativeAPIGatewayRoute(request *ModifyCloudNativeAPIGatewayRouteRequest) (response *ModifyCloudNativeAPIGatewayRouteResponse, err error) {
    return c.ModifyCloudNativeAPIGatewayRouteWithContext(context.Background(), request)
}

// ModifyCloudNativeAPIGatewayRoute
// 修改云原生网关路由
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
func (c *Client) ModifyCloudNativeAPIGatewayRouteWithContext(ctx context.Context, request *ModifyCloudNativeAPIGatewayRouteRequest) (response *ModifyCloudNativeAPIGatewayRouteResponse, err error) {
    if request == nil {
        request = NewModifyCloudNativeAPIGatewayRouteRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyCloudNativeAPIGatewayRoute require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyCloudNativeAPIGatewayRouteResponse()
    err = c.Send(request, response)
    return
}

func NewModifyCloudNativeAPIGatewayRouteRateLimitRequest() (request *ModifyCloudNativeAPIGatewayRouteRateLimitRequest) {
    request = &ModifyCloudNativeAPIGatewayRouteRateLimitRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "ModifyCloudNativeAPIGatewayRouteRateLimit")
    
    
    return
}

func NewModifyCloudNativeAPIGatewayRouteRateLimitResponse() (response *ModifyCloudNativeAPIGatewayRouteRateLimitResponse) {
    response = &ModifyCloudNativeAPIGatewayRouteRateLimitResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyCloudNativeAPIGatewayRouteRateLimit
// 修改云原生网关限流插件(路由)
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
func (c *Client) ModifyCloudNativeAPIGatewayRouteRateLimit(request *ModifyCloudNativeAPIGatewayRouteRateLimitRequest) (response *ModifyCloudNativeAPIGatewayRouteRateLimitResponse, err error) {
    return c.ModifyCloudNativeAPIGatewayRouteRateLimitWithContext(context.Background(), request)
}

// ModifyCloudNativeAPIGatewayRouteRateLimit
// 修改云原生网关限流插件(路由)
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
func (c *Client) ModifyCloudNativeAPIGatewayRouteRateLimitWithContext(ctx context.Context, request *ModifyCloudNativeAPIGatewayRouteRateLimitRequest) (response *ModifyCloudNativeAPIGatewayRouteRateLimitResponse, err error) {
    if request == nil {
        request = NewModifyCloudNativeAPIGatewayRouteRateLimitRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyCloudNativeAPIGatewayRouteRateLimit require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyCloudNativeAPIGatewayRouteRateLimitResponse()
    err = c.Send(request, response)
    return
}

func NewModifyCloudNativeAPIGatewayServiceRequest() (request *ModifyCloudNativeAPIGatewayServiceRequest) {
    request = &ModifyCloudNativeAPIGatewayServiceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "ModifyCloudNativeAPIGatewayService")
    
    
    return
}

func NewModifyCloudNativeAPIGatewayServiceResponse() (response *ModifyCloudNativeAPIGatewayServiceResponse) {
    response = &ModifyCloudNativeAPIGatewayServiceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyCloudNativeAPIGatewayService
// 修改云原生网关服务
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
func (c *Client) ModifyCloudNativeAPIGatewayService(request *ModifyCloudNativeAPIGatewayServiceRequest) (response *ModifyCloudNativeAPIGatewayServiceResponse, err error) {
    return c.ModifyCloudNativeAPIGatewayServiceWithContext(context.Background(), request)
}

// ModifyCloudNativeAPIGatewayService
// 修改云原生网关服务
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
func (c *Client) ModifyCloudNativeAPIGatewayServiceWithContext(ctx context.Context, request *ModifyCloudNativeAPIGatewayServiceRequest) (response *ModifyCloudNativeAPIGatewayServiceResponse, err error) {
    if request == nil {
        request = NewModifyCloudNativeAPIGatewayServiceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyCloudNativeAPIGatewayService require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyCloudNativeAPIGatewayServiceResponse()
    err = c.Send(request, response)
    return
}

func NewModifyCloudNativeAPIGatewayServiceRateLimitRequest() (request *ModifyCloudNativeAPIGatewayServiceRateLimitRequest) {
    request = &ModifyCloudNativeAPIGatewayServiceRateLimitRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "ModifyCloudNativeAPIGatewayServiceRateLimit")
    
    
    return
}

func NewModifyCloudNativeAPIGatewayServiceRateLimitResponse() (response *ModifyCloudNativeAPIGatewayServiceRateLimitResponse) {
    response = &ModifyCloudNativeAPIGatewayServiceRateLimitResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyCloudNativeAPIGatewayServiceRateLimit
// 修改云原生网关限流插件(服务)
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
func (c *Client) ModifyCloudNativeAPIGatewayServiceRateLimit(request *ModifyCloudNativeAPIGatewayServiceRateLimitRequest) (response *ModifyCloudNativeAPIGatewayServiceRateLimitResponse, err error) {
    return c.ModifyCloudNativeAPIGatewayServiceRateLimitWithContext(context.Background(), request)
}

// ModifyCloudNativeAPIGatewayServiceRateLimit
// 修改云原生网关限流插件(服务)
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
func (c *Client) ModifyCloudNativeAPIGatewayServiceRateLimitWithContext(ctx context.Context, request *ModifyCloudNativeAPIGatewayServiceRateLimitRequest) (response *ModifyCloudNativeAPIGatewayServiceRateLimitResponse, err error) {
    if request == nil {
        request = NewModifyCloudNativeAPIGatewayServiceRateLimitRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyCloudNativeAPIGatewayServiceRateLimit require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyCloudNativeAPIGatewayServiceRateLimitResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateEngineInternetAccessRequest() (request *UpdateEngineInternetAccessRequest) {
    request = &UpdateEngineInternetAccessRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "UpdateEngineInternetAccess")
    
    
    return
}

func NewUpdateEngineInternetAccessResponse() (response *UpdateEngineInternetAccessResponse) {
    response = &UpdateEngineInternetAccessResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpdateEngineInternetAccess
// 修改引擎公网访问配置
//
// 可能返回的错误码:
//  INTERNALERROR_OPERATIONFAILED = "InternalError.OperationFailed"
//  INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"
//  MISSINGPARAMETER_UPDATEERROR = "MissingParameter.UpdateError"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) UpdateEngineInternetAccess(request *UpdateEngineInternetAccessRequest) (response *UpdateEngineInternetAccessResponse, err error) {
    return c.UpdateEngineInternetAccessWithContext(context.Background(), request)
}

// UpdateEngineInternetAccess
// 修改引擎公网访问配置
//
// 可能返回的错误码:
//  INTERNALERROR_OPERATIONFAILED = "InternalError.OperationFailed"
//  INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"
//  MISSINGPARAMETER_UPDATEERROR = "MissingParameter.UpdateError"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) UpdateEngineInternetAccessWithContext(ctx context.Context, request *UpdateEngineInternetAccessRequest) (response *UpdateEngineInternetAccessResponse, err error) {
    if request == nil {
        request = NewUpdateEngineInternetAccessRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateEngineInternetAccess require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateEngineInternetAccessResponse()
    err = c.Send(request, response)
    return
}
