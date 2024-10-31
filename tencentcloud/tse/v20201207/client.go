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


func NewBindAutoScalerResourceStrategyToGroupsRequest() (request *BindAutoScalerResourceStrategyToGroupsRequest) {
    request = &BindAutoScalerResourceStrategyToGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "BindAutoScalerResourceStrategyToGroups")
    
    
    return
}

func NewBindAutoScalerResourceStrategyToGroupsResponse() (response *BindAutoScalerResourceStrategyToGroupsResponse) {
    response = &BindAutoScalerResourceStrategyToGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// BindAutoScalerResourceStrategyToGroups
// 弹性伸缩策略批量绑定网关分组
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMNOAUTH = "InternalError.CamNoAuth"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INTERNALERROR_OPERATIONFAILED = "InternalError.OperationFailed"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_GATEWAYID = "InvalidParameterValue.GatewayId"
//  INVALIDPARAMETERVALUE_QUERYERROR = "InvalidParameterValue.QueryError"
//  LIMITEXCEEDED_LIMITEXCEEDED = "LimitExceeded.LimitExceeded"
//  OPERATIONDENIED_OPERATIONDENIED = "OperationDenied.OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_FORBIDDEN = "ResourceNotFound.Forbidden"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) BindAutoScalerResourceStrategyToGroups(request *BindAutoScalerResourceStrategyToGroupsRequest) (response *BindAutoScalerResourceStrategyToGroupsResponse, err error) {
    return c.BindAutoScalerResourceStrategyToGroupsWithContext(context.Background(), request)
}

// BindAutoScalerResourceStrategyToGroups
// 弹性伸缩策略批量绑定网关分组
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMNOAUTH = "InternalError.CamNoAuth"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INTERNALERROR_OPERATIONFAILED = "InternalError.OperationFailed"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_GATEWAYID = "InvalidParameterValue.GatewayId"
//  INVALIDPARAMETERVALUE_QUERYERROR = "InvalidParameterValue.QueryError"
//  LIMITEXCEEDED_LIMITEXCEEDED = "LimitExceeded.LimitExceeded"
//  OPERATIONDENIED_OPERATIONDENIED = "OperationDenied.OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_FORBIDDEN = "ResourceNotFound.Forbidden"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) BindAutoScalerResourceStrategyToGroupsWithContext(ctx context.Context, request *BindAutoScalerResourceStrategyToGroupsRequest) (response *BindAutoScalerResourceStrategyToGroupsResponse, err error) {
    if request == nil {
        request = NewBindAutoScalerResourceStrategyToGroupsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("BindAutoScalerResourceStrategyToGroups require credential")
    }

    request.SetContext(ctx)
    
    response = NewBindAutoScalerResourceStrategyToGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewCloseWafProtectionRequest() (request *CloseWafProtectionRequest) {
    request = &CloseWafProtectionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "CloseWafProtection")
    
    
    return
}

func NewCloseWafProtectionResponse() (response *CloseWafProtectionResponse) {
    response = &CloseWafProtectionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CloseWafProtection
// 关闭 WAF 防护
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
func (c *Client) CloseWafProtection(request *CloseWafProtectionRequest) (response *CloseWafProtectionResponse, err error) {
    return c.CloseWafProtectionWithContext(context.Background(), request)
}

// CloseWafProtection
// 关闭 WAF 防护
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
func (c *Client) CloseWafProtectionWithContext(ctx context.Context, request *CloseWafProtectionRequest) (response *CloseWafProtectionResponse, err error) {
    if request == nil {
        request = NewCloseWafProtectionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CloseWafProtection require credential")
    }

    request.SetContext(ctx)
    
    response = NewCloseWafProtectionResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAutoScalerResourceStrategyRequest() (request *CreateAutoScalerResourceStrategyRequest) {
    request = &CreateAutoScalerResourceStrategyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "CreateAutoScalerResourceStrategy")
    
    
    return
}

func NewCreateAutoScalerResourceStrategyResponse() (response *CreateAutoScalerResourceStrategyResponse) {
    response = &CreateAutoScalerResourceStrategyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateAutoScalerResourceStrategy
// 创建弹性伸缩策略
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMNOAUTH = "InternalError.CamNoAuth"
//  INTERNALERROR_CREATEERROR = "InternalError.CreateError"
//  INTERNALERROR_DECODEERROR = "InternalError.DecodeError"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_UPDATEERROR = "InternalError.UpdateError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_LIMITEXCEEDED = "LimitExceeded.LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
func (c *Client) CreateAutoScalerResourceStrategy(request *CreateAutoScalerResourceStrategyRequest) (response *CreateAutoScalerResourceStrategyResponse, err error) {
    return c.CreateAutoScalerResourceStrategyWithContext(context.Background(), request)
}

// CreateAutoScalerResourceStrategy
// 创建弹性伸缩策略
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMNOAUTH = "InternalError.CamNoAuth"
//  INTERNALERROR_CREATEERROR = "InternalError.CreateError"
//  INTERNALERROR_DECODEERROR = "InternalError.DecodeError"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_UPDATEERROR = "InternalError.UpdateError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_LIMITEXCEEDED = "LimitExceeded.LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
func (c *Client) CreateAutoScalerResourceStrategyWithContext(ctx context.Context, request *CreateAutoScalerResourceStrategyRequest) (response *CreateAutoScalerResourceStrategyResponse, err error) {
    if request == nil {
        request = NewCreateAutoScalerResourceStrategyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAutoScalerResourceStrategy require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAutoScalerResourceStrategyResponse()
    err = c.Send(request, response)
    return
}

func NewCreateCloudNativeAPIGatewayRequest() (request *CreateCloudNativeAPIGatewayRequest) {
    request = &CreateCloudNativeAPIGatewayRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "CreateCloudNativeAPIGateway")
    
    
    return
}

func NewCreateCloudNativeAPIGatewayResponse() (response *CreateCloudNativeAPIGatewayResponse) {
    response = &CreateCloudNativeAPIGatewayResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateCloudNativeAPIGateway
// 创建云原生API网关实例
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLS = "FailedOperation.Cls"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_RESOURCE = "FailedOperation.Resource"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  FAILEDOPERATION_VPC = "FailedOperation.Vpc"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CREATEERROR = "InternalError.CreateError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_CREATEERROR = "InvalidParameterValue.CreateError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  INVALIDPARAMETERVALUE_SPECIFICATION = "InvalidParameterValue.Specification"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  LIMITEXCEEDED_LIMITEXCEEDED = "LimitExceeded.LimitExceeded"
//  MISSINGPARAMETER_MISSPARAMETER = "MissingParameter.MissParameter"
//  OPERATIONDENIED_OPERATIONDENIED = "OperationDenied.OperationDenied"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_CLSNOTACTIVATED = "UnauthorizedOperation.ClsNotActivated"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) CreateCloudNativeAPIGateway(request *CreateCloudNativeAPIGatewayRequest) (response *CreateCloudNativeAPIGatewayResponse, err error) {
    return c.CreateCloudNativeAPIGatewayWithContext(context.Background(), request)
}

// CreateCloudNativeAPIGateway
// 创建云原生API网关实例
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLS = "FailedOperation.Cls"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_RESOURCE = "FailedOperation.Resource"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  FAILEDOPERATION_VPC = "FailedOperation.Vpc"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CREATEERROR = "InternalError.CreateError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_CREATEERROR = "InvalidParameterValue.CreateError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  INVALIDPARAMETERVALUE_SPECIFICATION = "InvalidParameterValue.Specification"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  LIMITEXCEEDED_LIMITEXCEEDED = "LimitExceeded.LimitExceeded"
//  MISSINGPARAMETER_MISSPARAMETER = "MissingParameter.MissParameter"
//  OPERATIONDENIED_OPERATIONDENIED = "OperationDenied.OperationDenied"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_CLSNOTACTIVATED = "UnauthorizedOperation.ClsNotActivated"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) CreateCloudNativeAPIGatewayWithContext(ctx context.Context, request *CreateCloudNativeAPIGatewayRequest) (response *CreateCloudNativeAPIGatewayResponse, err error) {
    if request == nil {
        request = NewCreateCloudNativeAPIGatewayRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCloudNativeAPIGateway require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateCloudNativeAPIGatewayResponse()
    err = c.Send(request, response)
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

func NewCreateCloudNativeAPIGatewayCertificateRequest() (request *CreateCloudNativeAPIGatewayCertificateRequest) {
    request = &CreateCloudNativeAPIGatewayCertificateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "CreateCloudNativeAPIGatewayCertificate")
    
    
    return
}

func NewCreateCloudNativeAPIGatewayCertificateResponse() (response *CreateCloudNativeAPIGatewayCertificateResponse) {
    response = &CreateCloudNativeAPIGatewayCertificateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateCloudNativeAPIGatewayCertificate
// 创建云原生网关证书
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  FAILEDOPERATION_VPC = "FailedOperation.Vpc"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_OPERATIONFAILED = "InternalError.OperationFailed"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  INVALIDPARAMETERVALUE_SPECIFICATION = "InvalidParameterValue.Specification"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  LIMITEXCEEDED_LBDOMAINS = "LimitExceeded.LBDomains"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_CAMPASSROLENOTEXIST = "UnauthorizedOperation.CamPassRoleNotExist"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) CreateCloudNativeAPIGatewayCertificate(request *CreateCloudNativeAPIGatewayCertificateRequest) (response *CreateCloudNativeAPIGatewayCertificateResponse, err error) {
    return c.CreateCloudNativeAPIGatewayCertificateWithContext(context.Background(), request)
}

// CreateCloudNativeAPIGatewayCertificate
// 创建云原生网关证书
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  FAILEDOPERATION_VPC = "FailedOperation.Vpc"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_OPERATIONFAILED = "InternalError.OperationFailed"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  INVALIDPARAMETERVALUE_SPECIFICATION = "InvalidParameterValue.Specification"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  LIMITEXCEEDED_LBDOMAINS = "LimitExceeded.LBDomains"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_CAMPASSROLENOTEXIST = "UnauthorizedOperation.CamPassRoleNotExist"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) CreateCloudNativeAPIGatewayCertificateWithContext(ctx context.Context, request *CreateCloudNativeAPIGatewayCertificateRequest) (response *CreateCloudNativeAPIGatewayCertificateResponse, err error) {
    if request == nil {
        request = NewCreateCloudNativeAPIGatewayCertificateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCloudNativeAPIGatewayCertificate require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateCloudNativeAPIGatewayCertificateResponse()
    err = c.Send(request, response)
    return
}

func NewCreateCloudNativeAPIGatewayPublicNetworkRequest() (request *CreateCloudNativeAPIGatewayPublicNetworkRequest) {
    request = &CreateCloudNativeAPIGatewayPublicNetworkRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "CreateCloudNativeAPIGatewayPublicNetwork")
    
    
    return
}

func NewCreateCloudNativeAPIGatewayPublicNetworkResponse() (response *CreateCloudNativeAPIGatewayPublicNetworkResponse) {
    response = &CreateCloudNativeAPIGatewayPublicNetworkResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateCloudNativeAPIGatewayPublicNetwork
// 创建公网网络配置
//
// 可能返回的错误码:
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_CREATEERROR = "InvalidParameterValue.CreateError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
func (c *Client) CreateCloudNativeAPIGatewayPublicNetwork(request *CreateCloudNativeAPIGatewayPublicNetworkRequest) (response *CreateCloudNativeAPIGatewayPublicNetworkResponse, err error) {
    return c.CreateCloudNativeAPIGatewayPublicNetworkWithContext(context.Background(), request)
}

// CreateCloudNativeAPIGatewayPublicNetwork
// 创建公网网络配置
//
// 可能返回的错误码:
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_CREATEERROR = "InvalidParameterValue.CreateError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
func (c *Client) CreateCloudNativeAPIGatewayPublicNetworkWithContext(ctx context.Context, request *CreateCloudNativeAPIGatewayPublicNetworkRequest) (response *CreateCloudNativeAPIGatewayPublicNetworkResponse, err error) {
    if request == nil {
        request = NewCreateCloudNativeAPIGatewayPublicNetworkRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCloudNativeAPIGatewayPublicNetwork require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateCloudNativeAPIGatewayPublicNetworkResponse()
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

func NewCreateConfigFileRequest() (request *CreateConfigFileRequest) {
    request = &CreateConfigFileRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "CreateConfigFile")
    
    
    return
}

func NewCreateConfigFileResponse() (response *CreateConfigFileResponse) {
    response = &CreateConfigFileResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateConfigFile
// 创建配置文件
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GOVERNANCEFAILURE = "InternalError.GovernanceFailure"
//  INTERNALERROR_IOERROR = "InternalError.IOError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_CREATEERROR = "InvalidParameterValue.CreateError"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
func (c *Client) CreateConfigFile(request *CreateConfigFileRequest) (response *CreateConfigFileResponse, err error) {
    return c.CreateConfigFileWithContext(context.Background(), request)
}

// CreateConfigFile
// 创建配置文件
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GOVERNANCEFAILURE = "InternalError.GovernanceFailure"
//  INTERNALERROR_IOERROR = "InternalError.IOError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_CREATEERROR = "InvalidParameterValue.CreateError"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
func (c *Client) CreateConfigFileWithContext(ctx context.Context, request *CreateConfigFileRequest) (response *CreateConfigFileResponse, err error) {
    if request == nil {
        request = NewCreateConfigFileRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateConfigFile require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateConfigFileResponse()
    err = c.Send(request, response)
    return
}

func NewCreateConfigFileGroupRequest() (request *CreateConfigFileGroupRequest) {
    request = &CreateConfigFileGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "CreateConfigFileGroup")
    
    
    return
}

func NewCreateConfigFileGroupResponse() (response *CreateConfigFileGroupResponse) {
    response = &CreateConfigFileGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateConfigFileGroup
// 创建服务治理中心配置文件组
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CREATEERROR = "InternalError.CreateError"
//  INTERNALERROR_GOVERNANCEFAILURE = "InternalError.GovernanceFailure"
//  INTERNALERROR_IOERROR = "InternalError.IOError"
//  INTERNALERROR_OPERATIONFAILED = "InternalError.OperationFailed"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_CREATEERROR = "MissingParameter.CreateError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateConfigFileGroup(request *CreateConfigFileGroupRequest) (response *CreateConfigFileGroupResponse, err error) {
    return c.CreateConfigFileGroupWithContext(context.Background(), request)
}

// CreateConfigFileGroup
// 创建服务治理中心配置文件组
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CREATEERROR = "InternalError.CreateError"
//  INTERNALERROR_GOVERNANCEFAILURE = "InternalError.GovernanceFailure"
//  INTERNALERROR_IOERROR = "InternalError.IOError"
//  INTERNALERROR_OPERATIONFAILED = "InternalError.OperationFailed"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_CREATEERROR = "MissingParameter.CreateError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateConfigFileGroupWithContext(ctx context.Context, request *CreateConfigFileGroupRequest) (response *CreateConfigFileGroupResponse, err error) {
    if request == nil {
        request = NewCreateConfigFileGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateConfigFileGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateConfigFileGroupResponse()
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

func NewCreateGovernanceAliasRequest() (request *CreateGovernanceAliasRequest) {
    request = &CreateGovernanceAliasRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "CreateGovernanceAlias")
    
    
    return
}

func NewCreateGovernanceAliasResponse() (response *CreateGovernanceAliasResponse) {
    response = &CreateGovernanceAliasResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateGovernanceAlias
// 创建治理中心服务别名
//
// 可能返回的错误码:
//  INTERNALERROR_GOVERNANCEFAILURE = "InternalError.GovernanceFailure"
//  INTERNALERROR_IOERROR = "InternalError.IOError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_CREATEERROR = "InvalidParameterValue.CreateError"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateGovernanceAlias(request *CreateGovernanceAliasRequest) (response *CreateGovernanceAliasResponse, err error) {
    return c.CreateGovernanceAliasWithContext(context.Background(), request)
}

// CreateGovernanceAlias
// 创建治理中心服务别名
//
// 可能返回的错误码:
//  INTERNALERROR_GOVERNANCEFAILURE = "InternalError.GovernanceFailure"
//  INTERNALERROR_IOERROR = "InternalError.IOError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_CREATEERROR = "InvalidParameterValue.CreateError"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateGovernanceAliasWithContext(ctx context.Context, request *CreateGovernanceAliasRequest) (response *CreateGovernanceAliasResponse, err error) {
    if request == nil {
        request = NewCreateGovernanceAliasRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateGovernanceAlias require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateGovernanceAliasResponse()
    err = c.Send(request, response)
    return
}

func NewCreateGovernanceInstancesRequest() (request *CreateGovernanceInstancesRequest) {
    request = &CreateGovernanceInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "CreateGovernanceInstances")
    
    
    return
}

func NewCreateGovernanceInstancesResponse() (response *CreateGovernanceInstancesResponse) {
    response = &CreateGovernanceInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateGovernanceInstances
// 创建服务实例
//
// 可能返回的错误码:
//  INTERNALERROR_GOVERNANCEFAILURE = "InternalError.GovernanceFailure"
//  INTERNALERROR_IOERROR = "InternalError.IOError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_CREATEERROR = "InvalidParameterValue.CreateError"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateGovernanceInstances(request *CreateGovernanceInstancesRequest) (response *CreateGovernanceInstancesResponse, err error) {
    return c.CreateGovernanceInstancesWithContext(context.Background(), request)
}

// CreateGovernanceInstances
// 创建服务实例
//
// 可能返回的错误码:
//  INTERNALERROR_GOVERNANCEFAILURE = "InternalError.GovernanceFailure"
//  INTERNALERROR_IOERROR = "InternalError.IOError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_CREATEERROR = "InvalidParameterValue.CreateError"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateGovernanceInstancesWithContext(ctx context.Context, request *CreateGovernanceInstancesRequest) (response *CreateGovernanceInstancesResponse, err error) {
    if request == nil {
        request = NewCreateGovernanceInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateGovernanceInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateGovernanceInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewCreateGovernanceNamespacesRequest() (request *CreateGovernanceNamespacesRequest) {
    request = &CreateGovernanceNamespacesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "CreateGovernanceNamespaces")
    
    
    return
}

func NewCreateGovernanceNamespacesResponse() (response *CreateGovernanceNamespacesResponse) {
    response = &CreateGovernanceNamespacesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateGovernanceNamespaces
// 创建治理中心命名空间
//
// 可能返回的错误码:
//  INTERNALERROR_GOVERNANCEFAILURE = "InternalError.GovernanceFailure"
//  INTERNALERROR_IOERROR = "InternalError.IOError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_CREATEERROR = "InvalidParameterValue.CreateError"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateGovernanceNamespaces(request *CreateGovernanceNamespacesRequest) (response *CreateGovernanceNamespacesResponse, err error) {
    return c.CreateGovernanceNamespacesWithContext(context.Background(), request)
}

// CreateGovernanceNamespaces
// 创建治理中心命名空间
//
// 可能返回的错误码:
//  INTERNALERROR_GOVERNANCEFAILURE = "InternalError.GovernanceFailure"
//  INTERNALERROR_IOERROR = "InternalError.IOError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_CREATEERROR = "InvalidParameterValue.CreateError"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateGovernanceNamespacesWithContext(ctx context.Context, request *CreateGovernanceNamespacesRequest) (response *CreateGovernanceNamespacesResponse, err error) {
    if request == nil {
        request = NewCreateGovernanceNamespacesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateGovernanceNamespaces require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateGovernanceNamespacesResponse()
    err = c.Send(request, response)
    return
}

func NewCreateGovernanceServicesRequest() (request *CreateGovernanceServicesRequest) {
    request = &CreateGovernanceServicesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "CreateGovernanceServices")
    
    
    return
}

func NewCreateGovernanceServicesResponse() (response *CreateGovernanceServicesResponse) {
    response = &CreateGovernanceServicesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateGovernanceServices
// 创建治理中心服务
//
// 可能返回的错误码:
//  INTERNALERROR_GOVERNANCEFAILURE = "InternalError.GovernanceFailure"
//  INTERNALERROR_IOERROR = "InternalError.IOError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_CREATEERROR = "InvalidParameterValue.CreateError"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateGovernanceServices(request *CreateGovernanceServicesRequest) (response *CreateGovernanceServicesResponse, err error) {
    return c.CreateGovernanceServicesWithContext(context.Background(), request)
}

// CreateGovernanceServices
// 创建治理中心服务
//
// 可能返回的错误码:
//  INTERNALERROR_GOVERNANCEFAILURE = "InternalError.GovernanceFailure"
//  INTERNALERROR_IOERROR = "InternalError.IOError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_CREATEERROR = "InvalidParameterValue.CreateError"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateGovernanceServicesWithContext(ctx context.Context, request *CreateGovernanceServicesRequest) (response *CreateGovernanceServicesResponse, err error) {
    if request == nil {
        request = NewCreateGovernanceServicesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateGovernanceServices require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateGovernanceServicesResponse()
    err = c.Send(request, response)
    return
}

func NewCreateNativeGatewayServerGroupRequest() (request *CreateNativeGatewayServerGroupRequest) {
    request = &CreateNativeGatewayServerGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "CreateNativeGatewayServerGroup")
    
    
    return
}

func NewCreateNativeGatewayServerGroupResponse() (response *CreateNativeGatewayServerGroupResponse) {
    response = &CreateNativeGatewayServerGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateNativeGatewayServerGroup
// 创建云原生网关引擎分组
//
// 可能返回的错误码:
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"
//  OPERATIONDENIED_OPERATIONDENIED = "OperationDenied.OperationDenied"
func (c *Client) CreateNativeGatewayServerGroup(request *CreateNativeGatewayServerGroupRequest) (response *CreateNativeGatewayServerGroupResponse, err error) {
    return c.CreateNativeGatewayServerGroupWithContext(context.Background(), request)
}

// CreateNativeGatewayServerGroup
// 创建云原生网关引擎分组
//
// 可能返回的错误码:
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"
//  OPERATIONDENIED_OPERATIONDENIED = "OperationDenied.OperationDenied"
func (c *Client) CreateNativeGatewayServerGroupWithContext(ctx context.Context, request *CreateNativeGatewayServerGroupRequest) (response *CreateNativeGatewayServerGroupResponse, err error) {
    if request == nil {
        request = NewCreateNativeGatewayServerGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateNativeGatewayServerGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateNativeGatewayServerGroupResponse()
    err = c.Send(request, response)
    return
}

func NewCreateOrUpdateConfigFileAndReleaseRequest() (request *CreateOrUpdateConfigFileAndReleaseRequest) {
    request = &CreateOrUpdateConfigFileAndReleaseRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "CreateOrUpdateConfigFileAndRelease")
    
    
    return
}

func NewCreateOrUpdateConfigFileAndReleaseResponse() (response *CreateOrUpdateConfigFileAndReleaseResponse) {
    response = &CreateOrUpdateConfigFileAndReleaseResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateOrUpdateConfigFileAndRelease
// 创建或更新配置文件并发布配置
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GOVERNANCEFAILURE = "InternalError.GovernanceFailure"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_QUERYERROR = "InvalidParameterValue.QueryError"
func (c *Client) CreateOrUpdateConfigFileAndRelease(request *CreateOrUpdateConfigFileAndReleaseRequest) (response *CreateOrUpdateConfigFileAndReleaseResponse, err error) {
    return c.CreateOrUpdateConfigFileAndReleaseWithContext(context.Background(), request)
}

// CreateOrUpdateConfigFileAndRelease
// 创建或更新配置文件并发布配置
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GOVERNANCEFAILURE = "InternalError.GovernanceFailure"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_QUERYERROR = "InvalidParameterValue.QueryError"
func (c *Client) CreateOrUpdateConfigFileAndReleaseWithContext(ctx context.Context, request *CreateOrUpdateConfigFileAndReleaseRequest) (response *CreateOrUpdateConfigFileAndReleaseResponse, err error) {
    if request == nil {
        request = NewCreateOrUpdateConfigFileAndReleaseRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateOrUpdateConfigFileAndRelease require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateOrUpdateConfigFileAndReleaseResponse()
    err = c.Send(request, response)
    return
}

func NewCreateWafDomainsRequest() (request *CreateWafDomainsRequest) {
    request = &CreateWafDomainsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "CreateWafDomains")
    
    
    return
}

func NewCreateWafDomainsResponse() (response *CreateWafDomainsResponse) {
    response = &CreateWafDomainsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateWafDomains
// 新建 WAF 防护域名
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  FAILEDOPERATION_VPC = "FailedOperation.Vpc"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INTERNALERROR_OPERATIONFAILED = "InternalError.OperationFailed"
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
func (c *Client) CreateWafDomains(request *CreateWafDomainsRequest) (response *CreateWafDomainsResponse, err error) {
    return c.CreateWafDomainsWithContext(context.Background(), request)
}

// CreateWafDomains
// 新建 WAF 防护域名
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  FAILEDOPERATION_VPC = "FailedOperation.Vpc"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INTERNALERROR_OPERATIONFAILED = "InternalError.OperationFailed"
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
func (c *Client) CreateWafDomainsWithContext(ctx context.Context, request *CreateWafDomainsRequest) (response *CreateWafDomainsResponse, err error) {
    if request == nil {
        request = NewCreateWafDomainsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateWafDomains require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateWafDomainsResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAutoScalerResourceStrategyRequest() (request *DeleteAutoScalerResourceStrategyRequest) {
    request = &DeleteAutoScalerResourceStrategyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "DeleteAutoScalerResourceStrategy")
    
    
    return
}

func NewDeleteAutoScalerResourceStrategyResponse() (response *DeleteAutoScalerResourceStrategyResponse) {
    response = &DeleteAutoScalerResourceStrategyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteAutoScalerResourceStrategy
// 删除弹性伸缩策略
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  INTERNALERROR_CAMNOAUTH = "InternalError.CamNoAuth"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_GATEWAYID = "InvalidParameterValue.GatewayId"
//  OPERATIONDENIED_OPERATIONDENIED = "OperationDenied.OperationDenied"
//  RESOURCENOTFOUND_FORBIDDEN = "ResourceNotFound.Forbidden"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DeleteAutoScalerResourceStrategy(request *DeleteAutoScalerResourceStrategyRequest) (response *DeleteAutoScalerResourceStrategyResponse, err error) {
    return c.DeleteAutoScalerResourceStrategyWithContext(context.Background(), request)
}

// DeleteAutoScalerResourceStrategy
// 删除弹性伸缩策略
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  INTERNALERROR_CAMNOAUTH = "InternalError.CamNoAuth"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_GATEWAYID = "InvalidParameterValue.GatewayId"
//  OPERATIONDENIED_OPERATIONDENIED = "OperationDenied.OperationDenied"
//  RESOURCENOTFOUND_FORBIDDEN = "ResourceNotFound.Forbidden"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DeleteAutoScalerResourceStrategyWithContext(ctx context.Context, request *DeleteAutoScalerResourceStrategyRequest) (response *DeleteAutoScalerResourceStrategyResponse, err error) {
    if request == nil {
        request = NewDeleteAutoScalerResourceStrategyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteAutoScalerResourceStrategy require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteAutoScalerResourceStrategyResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteCloudNativeAPIGatewayRequest() (request *DeleteCloudNativeAPIGatewayRequest) {
    request = &DeleteCloudNativeAPIGatewayRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "DeleteCloudNativeAPIGateway")
    
    
    return
}

func NewDeleteCloudNativeAPIGatewayResponse() (response *DeleteCloudNativeAPIGatewayResponse) {
    response = &DeleteCloudNativeAPIGatewayResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteCloudNativeAPIGateway
// 删除云原生API网关实例
//
// 可能返回的错误码:
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_GATEWAYID = "InvalidParameterValue.GatewayId"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"
//  LIMITEXCEEDED_LIMITEXCEEDED = "LimitExceeded.LimitExceeded"
//  OPERATIONDENIED_OPERATIONDENIED = "OperationDenied.OperationDenied"
//  RESOURCENOTFOUND_FORBIDDEN = "ResourceNotFound.Forbidden"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMPASSROLENOTEXIST = "UnauthorizedOperation.CamPassRoleNotExist"
//  UNAUTHORIZEDOPERATION_CLSNOTACTIVATED = "UnauthorizedOperation.ClsNotActivated"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DeleteCloudNativeAPIGateway(request *DeleteCloudNativeAPIGatewayRequest) (response *DeleteCloudNativeAPIGatewayResponse, err error) {
    return c.DeleteCloudNativeAPIGatewayWithContext(context.Background(), request)
}

// DeleteCloudNativeAPIGateway
// 删除云原生API网关实例
//
// 可能返回的错误码:
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_GATEWAYID = "InvalidParameterValue.GatewayId"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"
//  LIMITEXCEEDED_LIMITEXCEEDED = "LimitExceeded.LimitExceeded"
//  OPERATIONDENIED_OPERATIONDENIED = "OperationDenied.OperationDenied"
//  RESOURCENOTFOUND_FORBIDDEN = "ResourceNotFound.Forbidden"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMPASSROLENOTEXIST = "UnauthorizedOperation.CamPassRoleNotExist"
//  UNAUTHORIZEDOPERATION_CLSNOTACTIVATED = "UnauthorizedOperation.ClsNotActivated"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DeleteCloudNativeAPIGatewayWithContext(ctx context.Context, request *DeleteCloudNativeAPIGatewayRequest) (response *DeleteCloudNativeAPIGatewayResponse, err error) {
    if request == nil {
        request = NewDeleteCloudNativeAPIGatewayRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteCloudNativeAPIGateway require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteCloudNativeAPIGatewayResponse()
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

func NewDeleteCloudNativeAPIGatewayCertificateRequest() (request *DeleteCloudNativeAPIGatewayCertificateRequest) {
    request = &DeleteCloudNativeAPIGatewayCertificateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "DeleteCloudNativeAPIGatewayCertificate")
    
    
    return
}

func NewDeleteCloudNativeAPIGatewayCertificateResponse() (response *DeleteCloudNativeAPIGatewayCertificateResponse) {
    response = &DeleteCloudNativeAPIGatewayCertificateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteCloudNativeAPIGatewayCertificate
// 删除云原生网关证书
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
func (c *Client) DeleteCloudNativeAPIGatewayCertificate(request *DeleteCloudNativeAPIGatewayCertificateRequest) (response *DeleteCloudNativeAPIGatewayCertificateResponse, err error) {
    return c.DeleteCloudNativeAPIGatewayCertificateWithContext(context.Background(), request)
}

// DeleteCloudNativeAPIGatewayCertificate
// 删除云原生网关证书
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
func (c *Client) DeleteCloudNativeAPIGatewayCertificateWithContext(ctx context.Context, request *DeleteCloudNativeAPIGatewayCertificateRequest) (response *DeleteCloudNativeAPIGatewayCertificateResponse, err error) {
    if request == nil {
        request = NewDeleteCloudNativeAPIGatewayCertificateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteCloudNativeAPIGatewayCertificate require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteCloudNativeAPIGatewayCertificateResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteCloudNativeAPIGatewayPublicNetworkRequest() (request *DeleteCloudNativeAPIGatewayPublicNetworkRequest) {
    request = &DeleteCloudNativeAPIGatewayPublicNetworkRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "DeleteCloudNativeAPIGatewayPublicNetwork")
    
    
    return
}

func NewDeleteCloudNativeAPIGatewayPublicNetworkResponse() (response *DeleteCloudNativeAPIGatewayPublicNetworkResponse) {
    response = &DeleteCloudNativeAPIGatewayPublicNetworkResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteCloudNativeAPIGatewayPublicNetwork
// 删除公网网络配置
//
// 可能返回的错误码:
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  INVALIDPARAMETERVALUE_UPDATEERROR = "InvalidParameterValue.UpdateError"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
func (c *Client) DeleteCloudNativeAPIGatewayPublicNetwork(request *DeleteCloudNativeAPIGatewayPublicNetworkRequest) (response *DeleteCloudNativeAPIGatewayPublicNetworkResponse, err error) {
    return c.DeleteCloudNativeAPIGatewayPublicNetworkWithContext(context.Background(), request)
}

// DeleteCloudNativeAPIGatewayPublicNetwork
// 删除公网网络配置
//
// 可能返回的错误码:
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  INVALIDPARAMETERVALUE_UPDATEERROR = "InvalidParameterValue.UpdateError"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
func (c *Client) DeleteCloudNativeAPIGatewayPublicNetworkWithContext(ctx context.Context, request *DeleteCloudNativeAPIGatewayPublicNetworkRequest) (response *DeleteCloudNativeAPIGatewayPublicNetworkResponse, err error) {
    if request == nil {
        request = NewDeleteCloudNativeAPIGatewayPublicNetworkRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteCloudNativeAPIGatewayPublicNetwork require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteCloudNativeAPIGatewayPublicNetworkResponse()
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

func NewDeleteConfigFileGroupRequest() (request *DeleteConfigFileGroupRequest) {
    request = &DeleteConfigFileGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "DeleteConfigFileGroup")
    
    
    return
}

func NewDeleteConfigFileGroupResponse() (response *DeleteConfigFileGroupResponse) {
    response = &DeleteConfigFileGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteConfigFileGroup
// 删除配置文件分组
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GOVERNANCEFAILURE = "InternalError.GovernanceFailure"
//  INTERNALERROR_IOERROR = "InternalError.IOError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_UPDATEERROR = "InvalidParameterValue.UpdateError"
//  RESOURCEINUSE_SERVICESEXISTEDINNAMESPACE = "ResourceInUse.ServicesExistedInNamespace"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteConfigFileGroup(request *DeleteConfigFileGroupRequest) (response *DeleteConfigFileGroupResponse, err error) {
    return c.DeleteConfigFileGroupWithContext(context.Background(), request)
}

// DeleteConfigFileGroup
// 删除配置文件分组
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GOVERNANCEFAILURE = "InternalError.GovernanceFailure"
//  INTERNALERROR_IOERROR = "InternalError.IOError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_UPDATEERROR = "InvalidParameterValue.UpdateError"
//  RESOURCEINUSE_SERVICESEXISTEDINNAMESPACE = "ResourceInUse.ServicesExistedInNamespace"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteConfigFileGroupWithContext(ctx context.Context, request *DeleteConfigFileGroupRequest) (response *DeleteConfigFileGroupResponse, err error) {
    if request == nil {
        request = NewDeleteConfigFileGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteConfigFileGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteConfigFileGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteConfigFileReleasesRequest() (request *DeleteConfigFileReleasesRequest) {
    request = &DeleteConfigFileReleasesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "DeleteConfigFileReleases")
    
    
    return
}

func NewDeleteConfigFileReleasesResponse() (response *DeleteConfigFileReleasesResponse) {
    response = &DeleteConfigFileReleasesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteConfigFileReleases
// 删除配置发布
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GOVERNANCEFAILURE = "InternalError.GovernanceFailure"
//  INTERNALERROR_IOERROR = "InternalError.IOError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_QUERYERROR = "InvalidParameterValue.QueryError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteConfigFileReleases(request *DeleteConfigFileReleasesRequest) (response *DeleteConfigFileReleasesResponse, err error) {
    return c.DeleteConfigFileReleasesWithContext(context.Background(), request)
}

// DeleteConfigFileReleases
// 删除配置发布
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GOVERNANCEFAILURE = "InternalError.GovernanceFailure"
//  INTERNALERROR_IOERROR = "InternalError.IOError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_QUERYERROR = "InvalidParameterValue.QueryError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteConfigFileReleasesWithContext(ctx context.Context, request *DeleteConfigFileReleasesRequest) (response *DeleteConfigFileReleasesResponse, err error) {
    if request == nil {
        request = NewDeleteConfigFileReleasesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteConfigFileReleases require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteConfigFileReleasesResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteConfigFilesRequest() (request *DeleteConfigFilesRequest) {
    request = &DeleteConfigFilesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "DeleteConfigFiles")
    
    
    return
}

func NewDeleteConfigFilesResponse() (response *DeleteConfigFilesResponse) {
    response = &DeleteConfigFilesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteConfigFiles
// 删除配置文件
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GOVERNANCEFAILURE = "InternalError.GovernanceFailure"
//  INTERNALERROR_IOERROR = "InternalError.IOError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_UPDATEERROR = "InvalidParameterValue.UpdateError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteConfigFiles(request *DeleteConfigFilesRequest) (response *DeleteConfigFilesResponse, err error) {
    return c.DeleteConfigFilesWithContext(context.Background(), request)
}

// DeleteConfigFiles
// 删除配置文件
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GOVERNANCEFAILURE = "InternalError.GovernanceFailure"
//  INTERNALERROR_IOERROR = "InternalError.IOError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_UPDATEERROR = "InvalidParameterValue.UpdateError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteConfigFilesWithContext(ctx context.Context, request *DeleteConfigFilesRequest) (response *DeleteConfigFilesResponse, err error) {
    if request == nil {
        request = NewDeleteConfigFilesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteConfigFiles require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteConfigFilesResponse()
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

func NewDeleteGovernanceAliasesRequest() (request *DeleteGovernanceAliasesRequest) {
    request = &DeleteGovernanceAliasesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "DeleteGovernanceAliases")
    
    
    return
}

func NewDeleteGovernanceAliasesResponse() (response *DeleteGovernanceAliasesResponse) {
    response = &DeleteGovernanceAliasesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteGovernanceAliases
// 删除治理中心服务别名
//
// 可能返回的错误码:
//  INTERNALERROR_GOVERNANCEFAILURE = "InternalError.GovernanceFailure"
//  INTERNALERROR_IOERROR = "InternalError.IOError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_CREATEERROR = "InvalidParameterValue.CreateError"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteGovernanceAliases(request *DeleteGovernanceAliasesRequest) (response *DeleteGovernanceAliasesResponse, err error) {
    return c.DeleteGovernanceAliasesWithContext(context.Background(), request)
}

// DeleteGovernanceAliases
// 删除治理中心服务别名
//
// 可能返回的错误码:
//  INTERNALERROR_GOVERNANCEFAILURE = "InternalError.GovernanceFailure"
//  INTERNALERROR_IOERROR = "InternalError.IOError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_CREATEERROR = "InvalidParameterValue.CreateError"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteGovernanceAliasesWithContext(ctx context.Context, request *DeleteGovernanceAliasesRequest) (response *DeleteGovernanceAliasesResponse, err error) {
    if request == nil {
        request = NewDeleteGovernanceAliasesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteGovernanceAliases require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteGovernanceAliasesResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteGovernanceInstancesRequest() (request *DeleteGovernanceInstancesRequest) {
    request = &DeleteGovernanceInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "DeleteGovernanceInstances")
    
    
    return
}

func NewDeleteGovernanceInstancesResponse() (response *DeleteGovernanceInstancesResponse) {
    response = &DeleteGovernanceInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteGovernanceInstances
// 删除服务实例
//
// 可能返回的错误码:
//  INTERNALERROR_GOVERNANCEFAILURE = "InternalError.GovernanceFailure"
//  INTERNALERROR_IOERROR = "InternalError.IOError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_UPDATEERROR = "InvalidParameterValue.UpdateError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteGovernanceInstances(request *DeleteGovernanceInstancesRequest) (response *DeleteGovernanceInstancesResponse, err error) {
    return c.DeleteGovernanceInstancesWithContext(context.Background(), request)
}

// DeleteGovernanceInstances
// 删除服务实例
//
// 可能返回的错误码:
//  INTERNALERROR_GOVERNANCEFAILURE = "InternalError.GovernanceFailure"
//  INTERNALERROR_IOERROR = "InternalError.IOError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_UPDATEERROR = "InvalidParameterValue.UpdateError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteGovernanceInstancesWithContext(ctx context.Context, request *DeleteGovernanceInstancesRequest) (response *DeleteGovernanceInstancesResponse, err error) {
    if request == nil {
        request = NewDeleteGovernanceInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteGovernanceInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteGovernanceInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteGovernanceInstancesByHostRequest() (request *DeleteGovernanceInstancesByHostRequest) {
    request = &DeleteGovernanceInstancesByHostRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "DeleteGovernanceInstancesByHost")
    
    
    return
}

func NewDeleteGovernanceInstancesByHostResponse() (response *DeleteGovernanceInstancesByHostResponse) {
    response = &DeleteGovernanceInstancesByHostResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteGovernanceInstancesByHost
// 删除治理中心服务实例
//
// 可能返回的错误码:
//  INTERNALERROR_GOVERNANCEFAILURE = "InternalError.GovernanceFailure"
//  INTERNALERROR_IOERROR = "InternalError.IOError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_UPDATEERROR = "InvalidParameterValue.UpdateError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteGovernanceInstancesByHost(request *DeleteGovernanceInstancesByHostRequest) (response *DeleteGovernanceInstancesByHostResponse, err error) {
    return c.DeleteGovernanceInstancesByHostWithContext(context.Background(), request)
}

// DeleteGovernanceInstancesByHost
// 删除治理中心服务实例
//
// 可能返回的错误码:
//  INTERNALERROR_GOVERNANCEFAILURE = "InternalError.GovernanceFailure"
//  INTERNALERROR_IOERROR = "InternalError.IOError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_UPDATEERROR = "InvalidParameterValue.UpdateError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteGovernanceInstancesByHostWithContext(ctx context.Context, request *DeleteGovernanceInstancesByHostRequest) (response *DeleteGovernanceInstancesByHostResponse, err error) {
    if request == nil {
        request = NewDeleteGovernanceInstancesByHostRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteGovernanceInstancesByHost require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteGovernanceInstancesByHostResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteGovernanceNamespacesRequest() (request *DeleteGovernanceNamespacesRequest) {
    request = &DeleteGovernanceNamespacesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "DeleteGovernanceNamespaces")
    
    
    return
}

func NewDeleteGovernanceNamespacesResponse() (response *DeleteGovernanceNamespacesResponse) {
    response = &DeleteGovernanceNamespacesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteGovernanceNamespaces
// 删除治理中心命名空间
//
// 可能返回的错误码:
//  INTERNALERROR_GOVERNANCEFAILURE = "InternalError.GovernanceFailure"
//  INTERNALERROR_IOERROR = "InternalError.IOError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_UPDATEERROR = "InvalidParameterValue.UpdateError"
//  RESOURCEINUSE_SERVICESEXISTEDINNAMESPACE = "ResourceInUse.ServicesExistedInNamespace"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteGovernanceNamespaces(request *DeleteGovernanceNamespacesRequest) (response *DeleteGovernanceNamespacesResponse, err error) {
    return c.DeleteGovernanceNamespacesWithContext(context.Background(), request)
}

// DeleteGovernanceNamespaces
// 删除治理中心命名空间
//
// 可能返回的错误码:
//  INTERNALERROR_GOVERNANCEFAILURE = "InternalError.GovernanceFailure"
//  INTERNALERROR_IOERROR = "InternalError.IOError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_UPDATEERROR = "InvalidParameterValue.UpdateError"
//  RESOURCEINUSE_SERVICESEXISTEDINNAMESPACE = "ResourceInUse.ServicesExistedInNamespace"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteGovernanceNamespacesWithContext(ctx context.Context, request *DeleteGovernanceNamespacesRequest) (response *DeleteGovernanceNamespacesResponse, err error) {
    if request == nil {
        request = NewDeleteGovernanceNamespacesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteGovernanceNamespaces require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteGovernanceNamespacesResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteGovernanceServicesRequest() (request *DeleteGovernanceServicesRequest) {
    request = &DeleteGovernanceServicesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "DeleteGovernanceServices")
    
    
    return
}

func NewDeleteGovernanceServicesResponse() (response *DeleteGovernanceServicesResponse) {
    response = &DeleteGovernanceServicesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteGovernanceServices
// 删除治理中心服务
//
// 可能返回的错误码:
//  INTERNALERROR_GOVERNANCEFAILURE = "InternalError.GovernanceFailure"
//  INTERNALERROR_IOERROR = "InternalError.IOError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"
//  INVALIDPARAMETERVALUE_UPDATEERROR = "InvalidParameterValue.UpdateError"
//  RESOURCEINUSE_INSTANCESEXISTEDINSERVICE = "ResourceInUse.InstancesExistedInService"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteGovernanceServices(request *DeleteGovernanceServicesRequest) (response *DeleteGovernanceServicesResponse, err error) {
    return c.DeleteGovernanceServicesWithContext(context.Background(), request)
}

// DeleteGovernanceServices
// 删除治理中心服务
//
// 可能返回的错误码:
//  INTERNALERROR_GOVERNANCEFAILURE = "InternalError.GovernanceFailure"
//  INTERNALERROR_IOERROR = "InternalError.IOError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"
//  INVALIDPARAMETERVALUE_UPDATEERROR = "InvalidParameterValue.UpdateError"
//  RESOURCEINUSE_INSTANCESEXISTEDINSERVICE = "ResourceInUse.InstancesExistedInService"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteGovernanceServicesWithContext(ctx context.Context, request *DeleteGovernanceServicesRequest) (response *DeleteGovernanceServicesResponse, err error) {
    if request == nil {
        request = NewDeleteGovernanceServicesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteGovernanceServices require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteGovernanceServicesResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteNativeGatewayServerGroupRequest() (request *DeleteNativeGatewayServerGroupRequest) {
    request = &DeleteNativeGatewayServerGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "DeleteNativeGatewayServerGroup")
    
    
    return
}

func NewDeleteNativeGatewayServerGroupResponse() (response *DeleteNativeGatewayServerGroupResponse) {
    response = &DeleteNativeGatewayServerGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteNativeGatewayServerGroup
// 删除网关实例分组
//
// 可能返回的错误码:
//  INTERNALERROR_GOVERNANCEFAILURE = "InternalError.GovernanceFailure"
//  INTERNALERROR_IOERROR = "InternalError.IOError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"
//  INVALIDPARAMETERVALUE_UPDATEERROR = "InvalidParameterValue.UpdateError"
//  RESOURCEINUSE_INSTANCESEXISTEDINSERVICE = "ResourceInUse.InstancesExistedInService"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteNativeGatewayServerGroup(request *DeleteNativeGatewayServerGroupRequest) (response *DeleteNativeGatewayServerGroupResponse, err error) {
    return c.DeleteNativeGatewayServerGroupWithContext(context.Background(), request)
}

// DeleteNativeGatewayServerGroup
// 删除网关实例分组
//
// 可能返回的错误码:
//  INTERNALERROR_GOVERNANCEFAILURE = "InternalError.GovernanceFailure"
//  INTERNALERROR_IOERROR = "InternalError.IOError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"
//  INVALIDPARAMETERVALUE_UPDATEERROR = "InvalidParameterValue.UpdateError"
//  RESOURCEINUSE_INSTANCESEXISTEDINSERVICE = "ResourceInUse.InstancesExistedInService"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteNativeGatewayServerGroupWithContext(ctx context.Context, request *DeleteNativeGatewayServerGroupRequest) (response *DeleteNativeGatewayServerGroupResponse, err error) {
    if request == nil {
        request = NewDeleteNativeGatewayServerGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteNativeGatewayServerGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteNativeGatewayServerGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteWafDomainsRequest() (request *DeleteWafDomainsRequest) {
    request = &DeleteWafDomainsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "DeleteWafDomains")
    
    
    return
}

func NewDeleteWafDomainsResponse() (response *DeleteWafDomainsResponse) {
    response = &DeleteWafDomainsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteWafDomains
// 删除 WAF 防护域名
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
func (c *Client) DeleteWafDomains(request *DeleteWafDomainsRequest) (response *DeleteWafDomainsResponse, err error) {
    return c.DeleteWafDomainsWithContext(context.Background(), request)
}

// DeleteWafDomains
// 删除 WAF 防护域名
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
func (c *Client) DeleteWafDomainsWithContext(ctx context.Context, request *DeleteWafDomainsRequest) (response *DeleteWafDomainsResponse, err error) {
    if request == nil {
        request = NewDeleteWafDomainsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteWafDomains require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteWafDomainsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAllConfigFileTemplatesRequest() (request *DescribeAllConfigFileTemplatesRequest) {
    request = &DescribeAllConfigFileTemplatesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "DescribeAllConfigFileTemplates")
    
    
    return
}

func NewDescribeAllConfigFileTemplatesResponse() (response *DescribeAllConfigFileTemplatesResponse) {
    response = &DescribeAllConfigFileTemplatesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAllConfigFileTemplates
// 获取全量配置文件模板列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAllConfigFileTemplates(request *DescribeAllConfigFileTemplatesRequest) (response *DescribeAllConfigFileTemplatesResponse, err error) {
    return c.DescribeAllConfigFileTemplatesWithContext(context.Background(), request)
}

// DescribeAllConfigFileTemplates
// 获取全量配置文件模板列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAllConfigFileTemplatesWithContext(ctx context.Context, request *DescribeAllConfigFileTemplatesRequest) (response *DescribeAllConfigFileTemplatesResponse, err error) {
    if request == nil {
        request = NewDescribeAllConfigFileTemplatesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAllConfigFileTemplates require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAllConfigFileTemplatesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAutoScalerResourceStrategiesRequest() (request *DescribeAutoScalerResourceStrategiesRequest) {
    request = &DescribeAutoScalerResourceStrategiesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "DescribeAutoScalerResourceStrategies")
    
    
    return
}

func NewDescribeAutoScalerResourceStrategiesResponse() (response *DescribeAutoScalerResourceStrategiesResponse) {
    response = &DescribeAutoScalerResourceStrategiesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAutoScalerResourceStrategies
// 查看弹性伸缩策略列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMNOAUTH = "InternalError.CamNoAuth"
//  INTERNALERROR_CREATEERROR = "InternalError.CreateError"
//  INTERNALERROR_DECODEERROR = "InternalError.DecodeError"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_UPDATEERROR = "InternalError.UpdateError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
func (c *Client) DescribeAutoScalerResourceStrategies(request *DescribeAutoScalerResourceStrategiesRequest) (response *DescribeAutoScalerResourceStrategiesResponse, err error) {
    return c.DescribeAutoScalerResourceStrategiesWithContext(context.Background(), request)
}

// DescribeAutoScalerResourceStrategies
// 查看弹性伸缩策略列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMNOAUTH = "InternalError.CamNoAuth"
//  INTERNALERROR_CREATEERROR = "InternalError.CreateError"
//  INTERNALERROR_DECODEERROR = "InternalError.DecodeError"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_UPDATEERROR = "InternalError.UpdateError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
func (c *Client) DescribeAutoScalerResourceStrategiesWithContext(ctx context.Context, request *DescribeAutoScalerResourceStrategiesRequest) (response *DescribeAutoScalerResourceStrategiesResponse, err error) {
    if request == nil {
        request = NewDescribeAutoScalerResourceStrategiesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAutoScalerResourceStrategies require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAutoScalerResourceStrategiesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAutoScalerResourceStrategyBindingGroupsRequest() (request *DescribeAutoScalerResourceStrategyBindingGroupsRequest) {
    request = &DescribeAutoScalerResourceStrategyBindingGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "DescribeAutoScalerResourceStrategyBindingGroups")
    
    
    return
}

func NewDescribeAutoScalerResourceStrategyBindingGroupsResponse() (response *DescribeAutoScalerResourceStrategyBindingGroupsResponse) {
    response = &DescribeAutoScalerResourceStrategyBindingGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAutoScalerResourceStrategyBindingGroups
// 查看弹性伸缩策略绑定的网关分组
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMNOAUTH = "InternalError.CamNoAuth"
//  INTERNALERROR_DECODEERROR = "InternalError.DecodeError"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INTERNALERROR_OPERATIONFAILED = "InternalError.OperationFailed"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_GATEWAYID = "InvalidParameterValue.GatewayId"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_MISSPARAMETER = "MissingParameter.MissParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_OPERATIONDENIED = "OperationDenied.OperationDenied"
//  RESOURCENOTFOUND_FORBIDDEN = "ResourceNotFound.Forbidden"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DescribeAutoScalerResourceStrategyBindingGroups(request *DescribeAutoScalerResourceStrategyBindingGroupsRequest) (response *DescribeAutoScalerResourceStrategyBindingGroupsResponse, err error) {
    return c.DescribeAutoScalerResourceStrategyBindingGroupsWithContext(context.Background(), request)
}

// DescribeAutoScalerResourceStrategyBindingGroups
// 查看弹性伸缩策略绑定的网关分组
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMNOAUTH = "InternalError.CamNoAuth"
//  INTERNALERROR_DECODEERROR = "InternalError.DecodeError"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INTERNALERROR_OPERATIONFAILED = "InternalError.OperationFailed"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_GATEWAYID = "InvalidParameterValue.GatewayId"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_MISSPARAMETER = "MissingParameter.MissParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_OPERATIONDENIED = "OperationDenied.OperationDenied"
//  RESOURCENOTFOUND_FORBIDDEN = "ResourceNotFound.Forbidden"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DescribeAutoScalerResourceStrategyBindingGroupsWithContext(ctx context.Context, request *DescribeAutoScalerResourceStrategyBindingGroupsRequest) (response *DescribeAutoScalerResourceStrategyBindingGroupsResponse, err error) {
    if request == nil {
        request = NewDescribeAutoScalerResourceStrategyBindingGroupsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAutoScalerResourceStrategyBindingGroups require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAutoScalerResourceStrategyBindingGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCloudNativeAPIGatewayRequest() (request *DescribeCloudNativeAPIGatewayRequest) {
    request = &DescribeCloudNativeAPIGatewayRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "DescribeCloudNativeAPIGateway")
    
    
    return
}

func NewDescribeCloudNativeAPIGatewayResponse() (response *DescribeCloudNativeAPIGatewayResponse) {
    response = &DescribeCloudNativeAPIGatewayResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCloudNativeAPIGateway
// 获取云原生API网关实例信息
//
// 可能返回的错误码:
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  INTERNALERROR_IOERROR = "InternalError.IOError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_GATEWAYID = "InvalidParameterValue.GatewayId"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  LIMITEXCEEDED_LIMITEXCEEDED = "LimitExceeded.LimitExceeded"
//  MISSINGPARAMETER_MISSPARAMETER = "MissingParameter.MissParameter"
//  OPERATIONDENIED_OPERATIONDENIED = "OperationDenied.OperationDenied"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DescribeCloudNativeAPIGateway(request *DescribeCloudNativeAPIGatewayRequest) (response *DescribeCloudNativeAPIGatewayResponse, err error) {
    return c.DescribeCloudNativeAPIGatewayWithContext(context.Background(), request)
}

// DescribeCloudNativeAPIGateway
// 获取云原生API网关实例信息
//
// 可能返回的错误码:
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  INTERNALERROR_IOERROR = "InternalError.IOError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_GATEWAYID = "InvalidParameterValue.GatewayId"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  LIMITEXCEEDED_LIMITEXCEEDED = "LimitExceeded.LimitExceeded"
//  MISSINGPARAMETER_MISSPARAMETER = "MissingParameter.MissParameter"
//  OPERATIONDENIED_OPERATIONDENIED = "OperationDenied.OperationDenied"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DescribeCloudNativeAPIGatewayWithContext(ctx context.Context, request *DescribeCloudNativeAPIGatewayRequest) (response *DescribeCloudNativeAPIGatewayResponse, err error) {
    if request == nil {
        request = NewDescribeCloudNativeAPIGatewayRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCloudNativeAPIGateway require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCloudNativeAPIGatewayResponse()
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

func NewDescribeCloudNativeAPIGatewayCertificateDetailsRequest() (request *DescribeCloudNativeAPIGatewayCertificateDetailsRequest) {
    request = &DescribeCloudNativeAPIGatewayCertificateDetailsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "DescribeCloudNativeAPIGatewayCertificateDetails")
    
    
    return
}

func NewDescribeCloudNativeAPIGatewayCertificateDetailsResponse() (response *DescribeCloudNativeAPIGatewayCertificateDetailsResponse) {
    response = &DescribeCloudNativeAPIGatewayCertificateDetailsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCloudNativeAPIGatewayCertificateDetails
// 查询云原生网关单个证书详情
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
func (c *Client) DescribeCloudNativeAPIGatewayCertificateDetails(request *DescribeCloudNativeAPIGatewayCertificateDetailsRequest) (response *DescribeCloudNativeAPIGatewayCertificateDetailsResponse, err error) {
    return c.DescribeCloudNativeAPIGatewayCertificateDetailsWithContext(context.Background(), request)
}

// DescribeCloudNativeAPIGatewayCertificateDetails
// 查询云原生网关单个证书详情
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
func (c *Client) DescribeCloudNativeAPIGatewayCertificateDetailsWithContext(ctx context.Context, request *DescribeCloudNativeAPIGatewayCertificateDetailsRequest) (response *DescribeCloudNativeAPIGatewayCertificateDetailsResponse, err error) {
    if request == nil {
        request = NewDescribeCloudNativeAPIGatewayCertificateDetailsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCloudNativeAPIGatewayCertificateDetails require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCloudNativeAPIGatewayCertificateDetailsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCloudNativeAPIGatewayCertificatesRequest() (request *DescribeCloudNativeAPIGatewayCertificatesRequest) {
    request = &DescribeCloudNativeAPIGatewayCertificatesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "DescribeCloudNativeAPIGatewayCertificates")
    
    
    return
}

func NewDescribeCloudNativeAPIGatewayCertificatesResponse() (response *DescribeCloudNativeAPIGatewayCertificatesResponse) {
    response = &DescribeCloudNativeAPIGatewayCertificatesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCloudNativeAPIGatewayCertificates
// 查询云原生网关证书列表
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
func (c *Client) DescribeCloudNativeAPIGatewayCertificates(request *DescribeCloudNativeAPIGatewayCertificatesRequest) (response *DescribeCloudNativeAPIGatewayCertificatesResponse, err error) {
    return c.DescribeCloudNativeAPIGatewayCertificatesWithContext(context.Background(), request)
}

// DescribeCloudNativeAPIGatewayCertificates
// 查询云原生网关证书列表
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
func (c *Client) DescribeCloudNativeAPIGatewayCertificatesWithContext(ctx context.Context, request *DescribeCloudNativeAPIGatewayCertificatesRequest) (response *DescribeCloudNativeAPIGatewayCertificatesResponse, err error) {
    if request == nil {
        request = NewDescribeCloudNativeAPIGatewayCertificatesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCloudNativeAPIGatewayCertificates require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCloudNativeAPIGatewayCertificatesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCloudNativeAPIGatewayConfigRequest() (request *DescribeCloudNativeAPIGatewayConfigRequest) {
    request = &DescribeCloudNativeAPIGatewayConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "DescribeCloudNativeAPIGatewayConfig")
    
    
    return
}

func NewDescribeCloudNativeAPIGatewayConfigResponse() (response *DescribeCloudNativeAPIGatewayConfigResponse) {
    response = &DescribeCloudNativeAPIGatewayConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCloudNativeAPIGatewayConfig
// 获取云原生API网关实例网络配置信息
//
// 可能返回的错误码:
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  INTERNALERROR_CREATEERROR = "InternalError.CreateError"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_GATEWAYID = "InvalidParameterValue.GatewayId"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  LIMITEXCEEDED_LIMITEXCEEDED = "LimitExceeded.LimitExceeded"
//  MISSINGPARAMETER_MISSPARAMETER = "MissingParameter.MissParameter"
//  OPERATIONDENIED_OPERATIONDENIED = "OperationDenied.OperationDenied"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DescribeCloudNativeAPIGatewayConfig(request *DescribeCloudNativeAPIGatewayConfigRequest) (response *DescribeCloudNativeAPIGatewayConfigResponse, err error) {
    return c.DescribeCloudNativeAPIGatewayConfigWithContext(context.Background(), request)
}

// DescribeCloudNativeAPIGatewayConfig
// 获取云原生API网关实例网络配置信息
//
// 可能返回的错误码:
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  INTERNALERROR_CREATEERROR = "InternalError.CreateError"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_GATEWAYID = "InvalidParameterValue.GatewayId"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  LIMITEXCEEDED_LIMITEXCEEDED = "LimitExceeded.LimitExceeded"
//  MISSINGPARAMETER_MISSPARAMETER = "MissingParameter.MissParameter"
//  OPERATIONDENIED_OPERATIONDENIED = "OperationDenied.OperationDenied"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DescribeCloudNativeAPIGatewayConfigWithContext(ctx context.Context, request *DescribeCloudNativeAPIGatewayConfigRequest) (response *DescribeCloudNativeAPIGatewayConfigResponse, err error) {
    if request == nil {
        request = NewDescribeCloudNativeAPIGatewayConfigRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCloudNativeAPIGatewayConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCloudNativeAPIGatewayConfigResponse()
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

func NewDescribeCloudNativeAPIGatewayPortsRequest() (request *DescribeCloudNativeAPIGatewayPortsRequest) {
    request = &DescribeCloudNativeAPIGatewayPortsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "DescribeCloudNativeAPIGatewayPorts")
    
    
    return
}

func NewDescribeCloudNativeAPIGatewayPortsResponse() (response *DescribeCloudNativeAPIGatewayPortsResponse) {
    response = &DescribeCloudNativeAPIGatewayPortsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCloudNativeAPIGatewayPorts
// 获取云原生API网关实例端口信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
func (c *Client) DescribeCloudNativeAPIGatewayPorts(request *DescribeCloudNativeAPIGatewayPortsRequest) (response *DescribeCloudNativeAPIGatewayPortsResponse, err error) {
    return c.DescribeCloudNativeAPIGatewayPortsWithContext(context.Background(), request)
}

// DescribeCloudNativeAPIGatewayPorts
// 获取云原生API网关实例端口信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
func (c *Client) DescribeCloudNativeAPIGatewayPortsWithContext(ctx context.Context, request *DescribeCloudNativeAPIGatewayPortsRequest) (response *DescribeCloudNativeAPIGatewayPortsResponse, err error) {
    if request == nil {
        request = NewDescribeCloudNativeAPIGatewayPortsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCloudNativeAPIGatewayPorts require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCloudNativeAPIGatewayPortsResponse()
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

func NewDescribeCloudNativeAPIGatewayUpstreamRequest() (request *DescribeCloudNativeAPIGatewayUpstreamRequest) {
    request = &DescribeCloudNativeAPIGatewayUpstreamRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "DescribeCloudNativeAPIGatewayUpstream")
    
    
    return
}

func NewDescribeCloudNativeAPIGatewayUpstreamResponse() (response *DescribeCloudNativeAPIGatewayUpstreamResponse) {
    response = &DescribeCloudNativeAPIGatewayUpstreamResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCloudNativeAPIGatewayUpstream
// 获取云原生网关服务详情下的Upstream列表
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
func (c *Client) DescribeCloudNativeAPIGatewayUpstream(request *DescribeCloudNativeAPIGatewayUpstreamRequest) (response *DescribeCloudNativeAPIGatewayUpstreamResponse, err error) {
    return c.DescribeCloudNativeAPIGatewayUpstreamWithContext(context.Background(), request)
}

// DescribeCloudNativeAPIGatewayUpstream
// 获取云原生网关服务详情下的Upstream列表
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
func (c *Client) DescribeCloudNativeAPIGatewayUpstreamWithContext(ctx context.Context, request *DescribeCloudNativeAPIGatewayUpstreamRequest) (response *DescribeCloudNativeAPIGatewayUpstreamResponse, err error) {
    if request == nil {
        request = NewDescribeCloudNativeAPIGatewayUpstreamRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCloudNativeAPIGatewayUpstream require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCloudNativeAPIGatewayUpstreamResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCloudNativeAPIGatewaysRequest() (request *DescribeCloudNativeAPIGatewaysRequest) {
    request = &DescribeCloudNativeAPIGatewaysRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "DescribeCloudNativeAPIGateways")
    
    
    return
}

func NewDescribeCloudNativeAPIGatewaysResponse() (response *DescribeCloudNativeAPIGatewaysResponse) {
    response = &DescribeCloudNativeAPIGatewaysResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCloudNativeAPIGateways
// 获取云原生API网关实例列表
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  INTERNALERROR_HTTPSTATUSCODEERROR = "InternalError.HttpStatusCodeError"
//  INTERNALERROR_IOERROR = "InternalError.IOError"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_GATEWAYID = "InvalidParameterValue.GatewayId"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_QUERYERROR = "InvalidParameterValue.QueryError"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  LIMITEXCEEDED_LIMITEXCEEDED = "LimitExceeded.LimitExceeded"
//  MISSINGPARAMETER_MISSPARAMETER = "MissingParameter.MissParameter"
//  OPERATIONDENIED_OPERATIONDENIED = "OperationDenied.OperationDenied"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeCloudNativeAPIGateways(request *DescribeCloudNativeAPIGatewaysRequest) (response *DescribeCloudNativeAPIGatewaysResponse, err error) {
    return c.DescribeCloudNativeAPIGatewaysWithContext(context.Background(), request)
}

// DescribeCloudNativeAPIGateways
// 获取云原生API网关实例列表
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  INTERNALERROR_HTTPSTATUSCODEERROR = "InternalError.HttpStatusCodeError"
//  INTERNALERROR_IOERROR = "InternalError.IOError"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_GATEWAYID = "InvalidParameterValue.GatewayId"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_QUERYERROR = "InvalidParameterValue.QueryError"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  LIMITEXCEEDED_LIMITEXCEEDED = "LimitExceeded.LimitExceeded"
//  MISSINGPARAMETER_MISSPARAMETER = "MissingParameter.MissParameter"
//  OPERATIONDENIED_OPERATIONDENIED = "OperationDenied.OperationDenied"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeCloudNativeAPIGatewaysWithContext(ctx context.Context, request *DescribeCloudNativeAPIGatewaysRequest) (response *DescribeCloudNativeAPIGatewaysResponse, err error) {
    if request == nil {
        request = NewDescribeCloudNativeAPIGatewaysRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCloudNativeAPIGateways require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCloudNativeAPIGatewaysResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeConfigFileRequest() (request *DescribeConfigFileRequest) {
    request = &DescribeConfigFileRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "DescribeConfigFile")
    
    
    return
}

func NewDescribeConfigFileResponse() (response *DescribeConfigFileResponse) {
    response = &DescribeConfigFileResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeConfigFile
// 根据命名空间、组、名字查找配置文件
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GOVERNANCEFAILURE = "InternalError.GovernanceFailure"
//  INTERNALERROR_IOERROR = "InternalError.IOError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_QUERYERROR = "InvalidParameterValue.QueryError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeConfigFile(request *DescribeConfigFileRequest) (response *DescribeConfigFileResponse, err error) {
    return c.DescribeConfigFileWithContext(context.Background(), request)
}

// DescribeConfigFile
// 根据命名空间、组、名字查找配置文件
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GOVERNANCEFAILURE = "InternalError.GovernanceFailure"
//  INTERNALERROR_IOERROR = "InternalError.IOError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_QUERYERROR = "InvalidParameterValue.QueryError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeConfigFileWithContext(ctx context.Context, request *DescribeConfigFileRequest) (response *DescribeConfigFileResponse, err error) {
    if request == nil {
        request = NewDescribeConfigFileRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeConfigFile require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeConfigFileResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeConfigFileGroupsRequest() (request *DescribeConfigFileGroupsRequest) {
    request = &DescribeConfigFileGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "DescribeConfigFileGroups")
    
    
    return
}

func NewDescribeConfigFileGroupsResponse() (response *DescribeConfigFileGroupsResponse) {
    response = &DescribeConfigFileGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeConfigFileGroups
// 根据条件分页查询配置文件组
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GOVERNANCEFAILURE = "InternalError.GovernanceFailure"
//  INTERNALERROR_IOERROR = "InternalError.IOError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_QUERYERROR = "InvalidParameterValue.QueryError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeConfigFileGroups(request *DescribeConfigFileGroupsRequest) (response *DescribeConfigFileGroupsResponse, err error) {
    return c.DescribeConfigFileGroupsWithContext(context.Background(), request)
}

// DescribeConfigFileGroups
// 根据条件分页查询配置文件组
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GOVERNANCEFAILURE = "InternalError.GovernanceFailure"
//  INTERNALERROR_IOERROR = "InternalError.IOError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_QUERYERROR = "InvalidParameterValue.QueryError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeConfigFileGroupsWithContext(ctx context.Context, request *DescribeConfigFileGroupsRequest) (response *DescribeConfigFileGroupsResponse, err error) {
    if request == nil {
        request = NewDescribeConfigFileGroupsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeConfigFileGroups require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeConfigFileGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeConfigFileReleaseRequest() (request *DescribeConfigFileReleaseRequest) {
    request = &DescribeConfigFileReleaseRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "DescribeConfigFileRelease")
    
    
    return
}

func NewDescribeConfigFileReleaseResponse() (response *DescribeConfigFileReleaseResponse) {
    response = &DescribeConfigFileReleaseResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeConfigFileRelease
// 获取配置文件发布
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_IOERROR = "InternalError.IOError"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeConfigFileRelease(request *DescribeConfigFileReleaseRequest) (response *DescribeConfigFileReleaseResponse, err error) {
    return c.DescribeConfigFileReleaseWithContext(context.Background(), request)
}

// DescribeConfigFileRelease
// 获取配置文件发布
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_IOERROR = "InternalError.IOError"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeConfigFileReleaseWithContext(ctx context.Context, request *DescribeConfigFileReleaseRequest) (response *DescribeConfigFileReleaseResponse, err error) {
    if request == nil {
        request = NewDescribeConfigFileReleaseRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeConfigFileRelease require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeConfigFileReleaseResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeConfigFileReleaseHistoriesRequest() (request *DescribeConfigFileReleaseHistoriesRequest) {
    request = &DescribeConfigFileReleaseHistoriesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "DescribeConfigFileReleaseHistories")
    
    
    return
}

func NewDescribeConfigFileReleaseHistoriesResponse() (response *DescribeConfigFileReleaseHistoriesResponse) {
    response = &DescribeConfigFileReleaseHistoriesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeConfigFileReleaseHistories
// 获取配置文件发布历史列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GOVERNANCEFAILURE = "InternalError.GovernanceFailure"
//  INTERNALERROR_IOERROR = "InternalError.IOError"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeConfigFileReleaseHistories(request *DescribeConfigFileReleaseHistoriesRequest) (response *DescribeConfigFileReleaseHistoriesResponse, err error) {
    return c.DescribeConfigFileReleaseHistoriesWithContext(context.Background(), request)
}

// DescribeConfigFileReleaseHistories
// 获取配置文件发布历史列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GOVERNANCEFAILURE = "InternalError.GovernanceFailure"
//  INTERNALERROR_IOERROR = "InternalError.IOError"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeConfigFileReleaseHistoriesWithContext(ctx context.Context, request *DescribeConfigFileReleaseHistoriesRequest) (response *DescribeConfigFileReleaseHistoriesResponse, err error) {
    if request == nil {
        request = NewDescribeConfigFileReleaseHistoriesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeConfigFileReleaseHistories require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeConfigFileReleaseHistoriesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeConfigFileReleaseVersionsRequest() (request *DescribeConfigFileReleaseVersionsRequest) {
    request = &DescribeConfigFileReleaseVersionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "DescribeConfigFileReleaseVersions")
    
    
    return
}

func NewDescribeConfigFileReleaseVersionsResponse() (response *DescribeConfigFileReleaseVersionsResponse) {
    response = &DescribeConfigFileReleaseVersionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeConfigFileReleaseVersions
// 查询某个配置所有版本信息
//
// 可能返回的错误码:
//  INTERNALERROR_GOVERNANCEFAILURE = "InternalError.GovernanceFailure"
//  INTERNALERROR_IOERROR = "InternalError.IOError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_QUERYERROR = "InvalidParameterValue.QueryError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeConfigFileReleaseVersions(request *DescribeConfigFileReleaseVersionsRequest) (response *DescribeConfigFileReleaseVersionsResponse, err error) {
    return c.DescribeConfigFileReleaseVersionsWithContext(context.Background(), request)
}

// DescribeConfigFileReleaseVersions
// 查询某个配置所有版本信息
//
// 可能返回的错误码:
//  INTERNALERROR_GOVERNANCEFAILURE = "InternalError.GovernanceFailure"
//  INTERNALERROR_IOERROR = "InternalError.IOError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_QUERYERROR = "InvalidParameterValue.QueryError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeConfigFileReleaseVersionsWithContext(ctx context.Context, request *DescribeConfigFileReleaseVersionsRequest) (response *DescribeConfigFileReleaseVersionsResponse, err error) {
    if request == nil {
        request = NewDescribeConfigFileReleaseVersionsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeConfigFileReleaseVersions require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeConfigFileReleaseVersionsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeConfigFileReleasesRequest() (request *DescribeConfigFileReleasesRequest) {
    request = &DescribeConfigFileReleasesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "DescribeConfigFileReleases")
    
    
    return
}

func NewDescribeConfigFileReleasesResponse() (response *DescribeConfigFileReleasesResponse) {
    response = &DescribeConfigFileReleasesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeConfigFileReleases
// 查询配置版本列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GOVERNANCEFAILURE = "InternalError.GovernanceFailure"
//  INTERNALERROR_IOERROR = "InternalError.IOError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_QUERYERROR = "InvalidParameterValue.QueryError"
func (c *Client) DescribeConfigFileReleases(request *DescribeConfigFileReleasesRequest) (response *DescribeConfigFileReleasesResponse, err error) {
    return c.DescribeConfigFileReleasesWithContext(context.Background(), request)
}

// DescribeConfigFileReleases
// 查询配置版本列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GOVERNANCEFAILURE = "InternalError.GovernanceFailure"
//  INTERNALERROR_IOERROR = "InternalError.IOError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_QUERYERROR = "InvalidParameterValue.QueryError"
func (c *Client) DescribeConfigFileReleasesWithContext(ctx context.Context, request *DescribeConfigFileReleasesRequest) (response *DescribeConfigFileReleasesResponse, err error) {
    if request == nil {
        request = NewDescribeConfigFileReleasesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeConfigFileReleases require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeConfigFileReleasesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeConfigFilesRequest() (request *DescribeConfigFilesRequest) {
    request = &DescribeConfigFilesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "DescribeConfigFiles")
    
    
    return
}

func NewDescribeConfigFilesResponse() (response *DescribeConfigFilesResponse) {
    response = &DescribeConfigFilesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeConfigFiles
// 根据命名空间、组名、名称、标签查询配置文件列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GOVERNANCEFAILURE = "InternalError.GovernanceFailure"
//  INTERNALERROR_IOERROR = "InternalError.IOError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_QUERYERROR = "InvalidParameterValue.QueryError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeConfigFiles(request *DescribeConfigFilesRequest) (response *DescribeConfigFilesResponse, err error) {
    return c.DescribeConfigFilesWithContext(context.Background(), request)
}

// DescribeConfigFiles
// 根据命名空间、组名、名称、标签查询配置文件列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GOVERNANCEFAILURE = "InternalError.GovernanceFailure"
//  INTERNALERROR_IOERROR = "InternalError.IOError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_QUERYERROR = "InvalidParameterValue.QueryError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeConfigFilesWithContext(ctx context.Context, request *DescribeConfigFilesRequest) (response *DescribeConfigFilesResponse, err error) {
    if request == nil {
        request = NewDescribeConfigFilesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeConfigFiles require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeConfigFilesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeConfigFilesByGroupRequest() (request *DescribeConfigFilesByGroupRequest) {
    request = &DescribeConfigFilesByGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "DescribeConfigFilesByGroup")
    
    
    return
}

func NewDescribeConfigFilesByGroupResponse() (response *DescribeConfigFilesByGroupResponse) {
    response = &DescribeConfigFilesByGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeConfigFilesByGroup
// 根据group查询配置文件列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GOVERNANCEFAILURE = "InternalError.GovernanceFailure"
//  INTERNALERROR_IOERROR = "InternalError.IOError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_QUERYERROR = "InvalidParameterValue.QueryError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeConfigFilesByGroup(request *DescribeConfigFilesByGroupRequest) (response *DescribeConfigFilesByGroupResponse, err error) {
    return c.DescribeConfigFilesByGroupWithContext(context.Background(), request)
}

// DescribeConfigFilesByGroup
// 根据group查询配置文件列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GOVERNANCEFAILURE = "InternalError.GovernanceFailure"
//  INTERNALERROR_IOERROR = "InternalError.IOError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_QUERYERROR = "InvalidParameterValue.QueryError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeConfigFilesByGroupWithContext(ctx context.Context, request *DescribeConfigFilesByGroupRequest) (response *DescribeConfigFilesByGroupResponse, err error) {
    if request == nil {
        request = NewDescribeConfigFilesByGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeConfigFilesByGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeConfigFilesByGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeGovernanceAliasesRequest() (request *DescribeGovernanceAliasesRequest) {
    request = &DescribeGovernanceAliasesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "DescribeGovernanceAliases")
    
    
    return
}

func NewDescribeGovernanceAliasesResponse() (response *DescribeGovernanceAliasesResponse) {
    response = &DescribeGovernanceAliasesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeGovernanceAliases
// 查询治理中心服务别名列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GOVERNANCEFAILURE = "InternalError.GovernanceFailure"
//  INTERNALERROR_IOERROR = "InternalError.IOError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_QUERYERROR = "InvalidParameterValue.QueryError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeGovernanceAliases(request *DescribeGovernanceAliasesRequest) (response *DescribeGovernanceAliasesResponse, err error) {
    return c.DescribeGovernanceAliasesWithContext(context.Background(), request)
}

// DescribeGovernanceAliases
// 查询治理中心服务别名列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GOVERNANCEFAILURE = "InternalError.GovernanceFailure"
//  INTERNALERROR_IOERROR = "InternalError.IOError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_QUERYERROR = "InvalidParameterValue.QueryError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeGovernanceAliasesWithContext(ctx context.Context, request *DescribeGovernanceAliasesRequest) (response *DescribeGovernanceAliasesResponse, err error) {
    if request == nil {
        request = NewDescribeGovernanceAliasesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeGovernanceAliases require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeGovernanceAliasesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeGovernanceInstancesRequest() (request *DescribeGovernanceInstancesRequest) {
    request = &DescribeGovernanceInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "DescribeGovernanceInstances")
    
    
    return
}

func NewDescribeGovernanceInstancesResponse() (response *DescribeGovernanceInstancesResponse) {
    response = &DescribeGovernanceInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeGovernanceInstances
// 查询服务实例
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GOVERNANCEFAILURE = "InternalError.GovernanceFailure"
//  INTERNALERROR_IOERROR = "InternalError.IOError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_QUERYERROR = "InvalidParameterValue.QueryError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeGovernanceInstances(request *DescribeGovernanceInstancesRequest) (response *DescribeGovernanceInstancesResponse, err error) {
    return c.DescribeGovernanceInstancesWithContext(context.Background(), request)
}

// DescribeGovernanceInstances
// 查询服务实例
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GOVERNANCEFAILURE = "InternalError.GovernanceFailure"
//  INTERNALERROR_IOERROR = "InternalError.IOError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_QUERYERROR = "InvalidParameterValue.QueryError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeGovernanceInstancesWithContext(ctx context.Context, request *DescribeGovernanceInstancesRequest) (response *DescribeGovernanceInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeGovernanceInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeGovernanceInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeGovernanceInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeGovernanceNamespacesRequest() (request *DescribeGovernanceNamespacesRequest) {
    request = &DescribeGovernanceNamespacesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "DescribeGovernanceNamespaces")
    
    
    return
}

func NewDescribeGovernanceNamespacesResponse() (response *DescribeGovernanceNamespacesResponse) {
    response = &DescribeGovernanceNamespacesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeGovernanceNamespaces
// 查询服务治理中心命名空间列表
//
// 可能返回的错误码:
//  INTERNALERROR_GOVERNANCEFAILURE = "InternalError.GovernanceFailure"
//  INTERNALERROR_IOERROR = "InternalError.IOError"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_QUERYERROR = "InvalidParameterValue.QueryError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeGovernanceNamespaces(request *DescribeGovernanceNamespacesRequest) (response *DescribeGovernanceNamespacesResponse, err error) {
    return c.DescribeGovernanceNamespacesWithContext(context.Background(), request)
}

// DescribeGovernanceNamespaces
// 查询服务治理中心命名空间列表
//
// 可能返回的错误码:
//  INTERNALERROR_GOVERNANCEFAILURE = "InternalError.GovernanceFailure"
//  INTERNALERROR_IOERROR = "InternalError.IOError"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_QUERYERROR = "InvalidParameterValue.QueryError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeGovernanceNamespacesWithContext(ctx context.Context, request *DescribeGovernanceNamespacesRequest) (response *DescribeGovernanceNamespacesResponse, err error) {
    if request == nil {
        request = NewDescribeGovernanceNamespacesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeGovernanceNamespaces require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeGovernanceNamespacesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeGovernanceServiceContractVersionsRequest() (request *DescribeGovernanceServiceContractVersionsRequest) {
    request = &DescribeGovernanceServiceContractVersionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "DescribeGovernanceServiceContractVersions")
    
    
    return
}

func NewDescribeGovernanceServiceContractVersionsResponse() (response *DescribeGovernanceServiceContractVersionsResponse) {
    response = &DescribeGovernanceServiceContractVersionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeGovernanceServiceContractVersions
// 查询服务下契约版本列表
//
// 可能返回的错误码:
//  INTERNALERROR_GOVERNANCEFAILURE = "InternalError.GovernanceFailure"
//  INTERNALERROR_IOERROR = "InternalError.IOError"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_QUERYERROR = "InvalidParameterValue.QueryError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeGovernanceServiceContractVersions(request *DescribeGovernanceServiceContractVersionsRequest) (response *DescribeGovernanceServiceContractVersionsResponse, err error) {
    return c.DescribeGovernanceServiceContractVersionsWithContext(context.Background(), request)
}

// DescribeGovernanceServiceContractVersions
// 查询服务下契约版本列表
//
// 可能返回的错误码:
//  INTERNALERROR_GOVERNANCEFAILURE = "InternalError.GovernanceFailure"
//  INTERNALERROR_IOERROR = "InternalError.IOError"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_QUERYERROR = "InvalidParameterValue.QueryError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeGovernanceServiceContractVersionsWithContext(ctx context.Context, request *DescribeGovernanceServiceContractVersionsRequest) (response *DescribeGovernanceServiceContractVersionsResponse, err error) {
    if request == nil {
        request = NewDescribeGovernanceServiceContractVersionsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeGovernanceServiceContractVersions require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeGovernanceServiceContractVersionsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeGovernanceServiceContractsRequest() (request *DescribeGovernanceServiceContractsRequest) {
    request = &DescribeGovernanceServiceContractsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "DescribeGovernanceServiceContracts")
    
    
    return
}

func NewDescribeGovernanceServiceContractsResponse() (response *DescribeGovernanceServiceContractsResponse) {
    response = &DescribeGovernanceServiceContractsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeGovernanceServiceContracts
// 查询服务契约定义列表
//
// 可能返回的错误码:
//  INTERNALERROR_GOVERNANCEFAILURE = "InternalError.GovernanceFailure"
//  INTERNALERROR_IOERROR = "InternalError.IOError"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_QUERYERROR = "InvalidParameterValue.QueryError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeGovernanceServiceContracts(request *DescribeGovernanceServiceContractsRequest) (response *DescribeGovernanceServiceContractsResponse, err error) {
    return c.DescribeGovernanceServiceContractsWithContext(context.Background(), request)
}

// DescribeGovernanceServiceContracts
// 查询服务契约定义列表
//
// 可能返回的错误码:
//  INTERNALERROR_GOVERNANCEFAILURE = "InternalError.GovernanceFailure"
//  INTERNALERROR_IOERROR = "InternalError.IOError"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_QUERYERROR = "InvalidParameterValue.QueryError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeGovernanceServiceContractsWithContext(ctx context.Context, request *DescribeGovernanceServiceContractsRequest) (response *DescribeGovernanceServiceContractsResponse, err error) {
    if request == nil {
        request = NewDescribeGovernanceServiceContractsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeGovernanceServiceContracts require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeGovernanceServiceContractsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeGovernanceServicesRequest() (request *DescribeGovernanceServicesRequest) {
    request = &DescribeGovernanceServicesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "DescribeGovernanceServices")
    
    
    return
}

func NewDescribeGovernanceServicesResponse() (response *DescribeGovernanceServicesResponse) {
    response = &DescribeGovernanceServicesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeGovernanceServices
// 查询治理中心服务列表
//
// 可能返回的错误码:
//  INTERNALERROR_GOVERNANCEFAILURE = "InternalError.GovernanceFailure"
//  INTERNALERROR_IOERROR = "InternalError.IOError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_QUERYERROR = "InvalidParameterValue.QueryError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeGovernanceServices(request *DescribeGovernanceServicesRequest) (response *DescribeGovernanceServicesResponse, err error) {
    return c.DescribeGovernanceServicesWithContext(context.Background(), request)
}

// DescribeGovernanceServices
// 查询治理中心服务列表
//
// 可能返回的错误码:
//  INTERNALERROR_GOVERNANCEFAILURE = "InternalError.GovernanceFailure"
//  INTERNALERROR_IOERROR = "InternalError.IOError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_QUERYERROR = "InvalidParameterValue.QueryError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeGovernanceServicesWithContext(ctx context.Context, request *DescribeGovernanceServicesRequest) (response *DescribeGovernanceServicesResponse, err error) {
    if request == nil {
        request = NewDescribeGovernanceServicesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeGovernanceServices require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeGovernanceServicesResponse()
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

func NewDescribeNativeGatewayServerGroupsRequest() (request *DescribeNativeGatewayServerGroupsRequest) {
    request = &DescribeNativeGatewayServerGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "DescribeNativeGatewayServerGroups")
    
    
    return
}

func NewDescribeNativeGatewayServerGroupsResponse() (response *DescribeNativeGatewayServerGroupsResponse) {
    response = &DescribeNativeGatewayServerGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeNativeGatewayServerGroups
// 查询云原生网关分组信息
//
// 可能返回的错误码:
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
func (c *Client) DescribeNativeGatewayServerGroups(request *DescribeNativeGatewayServerGroupsRequest) (response *DescribeNativeGatewayServerGroupsResponse, err error) {
    return c.DescribeNativeGatewayServerGroupsWithContext(context.Background(), request)
}

// DescribeNativeGatewayServerGroups
// 查询云原生网关分组信息
//
// 可能返回的错误码:
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
func (c *Client) DescribeNativeGatewayServerGroupsWithContext(ctx context.Context, request *DescribeNativeGatewayServerGroupsRequest) (response *DescribeNativeGatewayServerGroupsResponse, err error) {
    if request == nil {
        request = NewDescribeNativeGatewayServerGroupsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeNativeGatewayServerGroups require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeNativeGatewayServerGroupsResponse()
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

func NewDescribePublicAddressConfigRequest() (request *DescribePublicAddressConfigRequest) {
    request = &DescribePublicAddressConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "DescribePublicAddressConfig")
    
    
    return
}

func NewDescribePublicAddressConfigResponse() (response *DescribePublicAddressConfigResponse) {
    response = &DescribePublicAddressConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribePublicAddressConfig
// 查询公网地址信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
func (c *Client) DescribePublicAddressConfig(request *DescribePublicAddressConfigRequest) (response *DescribePublicAddressConfigResponse, err error) {
    return c.DescribePublicAddressConfigWithContext(context.Background(), request)
}

// DescribePublicAddressConfig
// 查询公网地址信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
func (c *Client) DescribePublicAddressConfigWithContext(ctx context.Context, request *DescribePublicAddressConfigRequest) (response *DescribePublicAddressConfigResponse, err error) {
    if request == nil {
        request = NewDescribePublicAddressConfigRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePublicAddressConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePublicAddressConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePublicNetworkRequest() (request *DescribePublicNetworkRequest) {
    request = &DescribePublicNetworkRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "DescribePublicNetwork")
    
    
    return
}

func NewDescribePublicNetworkResponse() (response *DescribePublicNetworkResponse) {
    response = &DescribePublicNetworkResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribePublicNetwork
// 查询云原生API网关实例公网详情
//
// 可能返回的错误码:
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  INTERNALERROR_CREATEERROR = "InternalError.CreateError"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_GATEWAYID = "InvalidParameterValue.GatewayId"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  LIMITEXCEEDED_LIMITEXCEEDED = "LimitExceeded.LimitExceeded"
//  MISSINGPARAMETER_MISSPARAMETER = "MissingParameter.MissParameter"
//  OPERATIONDENIED_OPERATIONDENIED = "OperationDenied.OperationDenied"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DescribePublicNetwork(request *DescribePublicNetworkRequest) (response *DescribePublicNetworkResponse, err error) {
    return c.DescribePublicNetworkWithContext(context.Background(), request)
}

// DescribePublicNetwork
// 查询云原生API网关实例公网详情
//
// 可能返回的错误码:
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  INTERNALERROR_CREATEERROR = "InternalError.CreateError"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_GATEWAYID = "InvalidParameterValue.GatewayId"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  LIMITEXCEEDED_LIMITEXCEEDED = "LimitExceeded.LimitExceeded"
//  MISSINGPARAMETER_MISSPARAMETER = "MissingParameter.MissParameter"
//  OPERATIONDENIED_OPERATIONDENIED = "OperationDenied.OperationDenied"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DescribePublicNetworkWithContext(ctx context.Context, request *DescribePublicNetworkRequest) (response *DescribePublicNetworkResponse, err error) {
    if request == nil {
        request = NewDescribePublicNetworkRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePublicNetwork require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePublicNetworkResponse()
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
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
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
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
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

func NewDescribeUpstreamHealthCheckConfigRequest() (request *DescribeUpstreamHealthCheckConfigRequest) {
    request = &DescribeUpstreamHealthCheckConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "DescribeUpstreamHealthCheckConfig")
    
    
    return
}

func NewDescribeUpstreamHealthCheckConfigResponse() (response *DescribeUpstreamHealthCheckConfigResponse) {
    response = &DescribeUpstreamHealthCheckConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeUpstreamHealthCheckConfig
// 获取云原生网关服务健康检查配置
//
// 可能返回的错误码:
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
func (c *Client) DescribeUpstreamHealthCheckConfig(request *DescribeUpstreamHealthCheckConfigRequest) (response *DescribeUpstreamHealthCheckConfigResponse, err error) {
    return c.DescribeUpstreamHealthCheckConfigWithContext(context.Background(), request)
}

// DescribeUpstreamHealthCheckConfig
// 获取云原生网关服务健康检查配置
//
// 可能返回的错误码:
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
func (c *Client) DescribeUpstreamHealthCheckConfigWithContext(ctx context.Context, request *DescribeUpstreamHealthCheckConfigRequest) (response *DescribeUpstreamHealthCheckConfigResponse, err error) {
    if request == nil {
        request = NewDescribeUpstreamHealthCheckConfigRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUpstreamHealthCheckConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeUpstreamHealthCheckConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeWafDomainsRequest() (request *DescribeWafDomainsRequest) {
    request = &DescribeWafDomainsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "DescribeWafDomains")
    
    
    return
}

func NewDescribeWafDomainsResponse() (response *DescribeWafDomainsResponse) {
    response = &DescribeWafDomainsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeWafDomains
// 获取 WAF 防护域名
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
func (c *Client) DescribeWafDomains(request *DescribeWafDomainsRequest) (response *DescribeWafDomainsResponse, err error) {
    return c.DescribeWafDomainsWithContext(context.Background(), request)
}

// DescribeWafDomains
// 获取 WAF 防护域名
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
func (c *Client) DescribeWafDomainsWithContext(ctx context.Context, request *DescribeWafDomainsRequest) (response *DescribeWafDomainsResponse, err error) {
    if request == nil {
        request = NewDescribeWafDomainsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeWafDomains require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeWafDomainsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeWafProtectionRequest() (request *DescribeWafProtectionRequest) {
    request = &DescribeWafProtectionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "DescribeWafProtection")
    
    
    return
}

func NewDescribeWafProtectionResponse() (response *DescribeWafProtectionResponse) {
    response = &DescribeWafProtectionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeWafProtection
// 获取 WAF 防护状态
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
func (c *Client) DescribeWafProtection(request *DescribeWafProtectionRequest) (response *DescribeWafProtectionResponse, err error) {
    return c.DescribeWafProtectionWithContext(context.Background(), request)
}

// DescribeWafProtection
// 获取 WAF 防护状态
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
func (c *Client) DescribeWafProtectionWithContext(ctx context.Context, request *DescribeWafProtectionRequest) (response *DescribeWafProtectionResponse, err error) {
    if request == nil {
        request = NewDescribeWafProtectionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeWafProtection require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeWafProtectionResponse()
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
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
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
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
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

func NewModifyAutoScalerResourceStrategyRequest() (request *ModifyAutoScalerResourceStrategyRequest) {
    request = &ModifyAutoScalerResourceStrategyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "ModifyAutoScalerResourceStrategy")
    
    
    return
}

func NewModifyAutoScalerResourceStrategyResponse() (response *ModifyAutoScalerResourceStrategyResponse) {
    response = &ModifyAutoScalerResourceStrategyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyAutoScalerResourceStrategy
// 更新弹性伸缩策略
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_DESCRIPTION = "InvalidParameterValue.Description"
//  INVALIDPARAMETERVALUE_GATEWAYID = "InvalidParameterValue.GatewayId"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  MISSINGPARAMETER_MISSPARAMETER = "MissingParameter.MissParameter"
//  OPERATIONDENIED_OPERATIONDENIED = "OperationDenied.OperationDenied"
//  RESOURCENOTFOUND_FORBIDDEN = "ResourceNotFound.Forbidden"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) ModifyAutoScalerResourceStrategy(request *ModifyAutoScalerResourceStrategyRequest) (response *ModifyAutoScalerResourceStrategyResponse, err error) {
    return c.ModifyAutoScalerResourceStrategyWithContext(context.Background(), request)
}

// ModifyAutoScalerResourceStrategy
// 更新弹性伸缩策略
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_DESCRIPTION = "InvalidParameterValue.Description"
//  INVALIDPARAMETERVALUE_GATEWAYID = "InvalidParameterValue.GatewayId"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  MISSINGPARAMETER_MISSPARAMETER = "MissingParameter.MissParameter"
//  OPERATIONDENIED_OPERATIONDENIED = "OperationDenied.OperationDenied"
//  RESOURCENOTFOUND_FORBIDDEN = "ResourceNotFound.Forbidden"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) ModifyAutoScalerResourceStrategyWithContext(ctx context.Context, request *ModifyAutoScalerResourceStrategyRequest) (response *ModifyAutoScalerResourceStrategyResponse, err error) {
    if request == nil {
        request = NewModifyAutoScalerResourceStrategyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAutoScalerResourceStrategy require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAutoScalerResourceStrategyResponse()
    err = c.Send(request, response)
    return
}

func NewModifyCloudNativeAPIGatewayRequest() (request *ModifyCloudNativeAPIGatewayRequest) {
    request = &ModifyCloudNativeAPIGatewayRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "ModifyCloudNativeAPIGateway")
    
    
    return
}

func NewModifyCloudNativeAPIGatewayResponse() (response *ModifyCloudNativeAPIGatewayResponse) {
    response = &ModifyCloudNativeAPIGatewayResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyCloudNativeAPIGateway
// 修改云原生API网关实例基础信息
//
// 可能返回的错误码:
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_DESCRIPTION = "InvalidParameterValue.Description"
//  INVALIDPARAMETERVALUE_GATEWAYID = "InvalidParameterValue.GatewayId"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  MISSINGPARAMETER_MISSPARAMETER = "MissingParameter.MissParameter"
//  OPERATIONDENIED_OPERATIONDENIED = "OperationDenied.OperationDenied"
//  RESOURCENOTFOUND_FORBIDDEN = "ResourceNotFound.Forbidden"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) ModifyCloudNativeAPIGateway(request *ModifyCloudNativeAPIGatewayRequest) (response *ModifyCloudNativeAPIGatewayResponse, err error) {
    return c.ModifyCloudNativeAPIGatewayWithContext(context.Background(), request)
}

// ModifyCloudNativeAPIGateway
// 修改云原生API网关实例基础信息
//
// 可能返回的错误码:
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_DESCRIPTION = "InvalidParameterValue.Description"
//  INVALIDPARAMETERVALUE_GATEWAYID = "InvalidParameterValue.GatewayId"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  MISSINGPARAMETER_MISSPARAMETER = "MissingParameter.MissParameter"
//  OPERATIONDENIED_OPERATIONDENIED = "OperationDenied.OperationDenied"
//  RESOURCENOTFOUND_FORBIDDEN = "ResourceNotFound.Forbidden"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) ModifyCloudNativeAPIGatewayWithContext(ctx context.Context, request *ModifyCloudNativeAPIGatewayRequest) (response *ModifyCloudNativeAPIGatewayResponse, err error) {
    if request == nil {
        request = NewModifyCloudNativeAPIGatewayRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyCloudNativeAPIGateway require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyCloudNativeAPIGatewayResponse()
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

func NewModifyCloudNativeAPIGatewayCertificateRequest() (request *ModifyCloudNativeAPIGatewayCertificateRequest) {
    request = &ModifyCloudNativeAPIGatewayCertificateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "ModifyCloudNativeAPIGatewayCertificate")
    
    
    return
}

func NewModifyCloudNativeAPIGatewayCertificateResponse() (response *ModifyCloudNativeAPIGatewayCertificateResponse) {
    response = &ModifyCloudNativeAPIGatewayCertificateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyCloudNativeAPIGatewayCertificate
// 更新云原生网关证书
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  FAILEDOPERATION_VPC = "FailedOperation.Vpc"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_OPERATIONFAILED = "InternalError.OperationFailed"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  INVALIDPARAMETERVALUE_SPECIFICATION = "InvalidParameterValue.Specification"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  LIMITEXCEEDED_LBDOMAINS = "LimitExceeded.LBDomains"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_CAMPASSROLENOTEXIST = "UnauthorizedOperation.CamPassRoleNotExist"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) ModifyCloudNativeAPIGatewayCertificate(request *ModifyCloudNativeAPIGatewayCertificateRequest) (response *ModifyCloudNativeAPIGatewayCertificateResponse, err error) {
    return c.ModifyCloudNativeAPIGatewayCertificateWithContext(context.Background(), request)
}

// ModifyCloudNativeAPIGatewayCertificate
// 更新云原生网关证书
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  FAILEDOPERATION_VPC = "FailedOperation.Vpc"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_OPERATIONFAILED = "InternalError.OperationFailed"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  INVALIDPARAMETERVALUE_SPECIFICATION = "InvalidParameterValue.Specification"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  LIMITEXCEEDED_LBDOMAINS = "LimitExceeded.LBDomains"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_CAMPASSROLENOTEXIST = "UnauthorizedOperation.CamPassRoleNotExist"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) ModifyCloudNativeAPIGatewayCertificateWithContext(ctx context.Context, request *ModifyCloudNativeAPIGatewayCertificateRequest) (response *ModifyCloudNativeAPIGatewayCertificateResponse, err error) {
    if request == nil {
        request = NewModifyCloudNativeAPIGatewayCertificateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyCloudNativeAPIGatewayCertificate require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyCloudNativeAPIGatewayCertificateResponse()
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
//  MISSINGPARAMETER_CREATEERROR = "MissingParameter.CreateError"
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
//  MISSINGPARAMETER_CREATEERROR = "MissingParameter.CreateError"
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

func NewModifyConfigFileGroupRequest() (request *ModifyConfigFileGroupRequest) {
    request = &ModifyConfigFileGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "ModifyConfigFileGroup")
    
    
    return
}

func NewModifyConfigFileGroupResponse() (response *ModifyConfigFileGroupResponse) {
    response = &ModifyConfigFileGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyConfigFileGroup
// 批量修改配置文件组
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GOVERNANCEFAILURE = "InternalError.GovernanceFailure"
//  INTERNALERROR_IOERROR = "InternalError.IOError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_UPDATEERROR = "InvalidParameterValue.UpdateError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyConfigFileGroup(request *ModifyConfigFileGroupRequest) (response *ModifyConfigFileGroupResponse, err error) {
    return c.ModifyConfigFileGroupWithContext(context.Background(), request)
}

// ModifyConfigFileGroup
// 批量修改配置文件组
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GOVERNANCEFAILURE = "InternalError.GovernanceFailure"
//  INTERNALERROR_IOERROR = "InternalError.IOError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_UPDATEERROR = "InvalidParameterValue.UpdateError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyConfigFileGroupWithContext(ctx context.Context, request *ModifyConfigFileGroupRequest) (response *ModifyConfigFileGroupResponse, err error) {
    if request == nil {
        request = NewModifyConfigFileGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyConfigFileGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyConfigFileGroupResponse()
    err = c.Send(request, response)
    return
}

func NewModifyConfigFilesRequest() (request *ModifyConfigFilesRequest) {
    request = &ModifyConfigFilesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "ModifyConfigFiles")
    
    
    return
}

func NewModifyConfigFilesResponse() (response *ModifyConfigFilesResponse) {
    response = &ModifyConfigFilesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyConfigFiles
// 修改配置文件
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GOVERNANCEFAILURE = "InternalError.GovernanceFailure"
//  INTERNALERROR_IOERROR = "InternalError.IOError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_UPDATEERROR = "InvalidParameterValue.UpdateError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyConfigFiles(request *ModifyConfigFilesRequest) (response *ModifyConfigFilesResponse, err error) {
    return c.ModifyConfigFilesWithContext(context.Background(), request)
}

// ModifyConfigFiles
// 修改配置文件
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GOVERNANCEFAILURE = "InternalError.GovernanceFailure"
//  INTERNALERROR_IOERROR = "InternalError.IOError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_UPDATEERROR = "InvalidParameterValue.UpdateError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyConfigFilesWithContext(ctx context.Context, request *ModifyConfigFilesRequest) (response *ModifyConfigFilesResponse, err error) {
    if request == nil {
        request = NewModifyConfigFilesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyConfigFiles require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyConfigFilesResponse()
    err = c.Send(request, response)
    return
}

func NewModifyConsoleNetworkRequest() (request *ModifyConsoleNetworkRequest) {
    request = &ModifyConsoleNetworkRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "ModifyConsoleNetwork")
    
    
    return
}

func NewModifyConsoleNetworkResponse() (response *ModifyConsoleNetworkResponse) {
    response = &ModifyConsoleNetworkResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyConsoleNetwork
// 修改网关实例Konga网络配置
//
// 可能返回的错误码:
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_UPDATEERROR = "InvalidParameterValue.UpdateError"
//  MISSINGPARAMETER_MISSPARAMETER = "MissingParameter.MissParameter"
//  OPERATIONDENIED_OPERATIONDENIED = "OperationDenied.OperationDenied"
//  RESOURCENOTFOUND_FORBIDDEN = "ResourceNotFound.Forbidden"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) ModifyConsoleNetwork(request *ModifyConsoleNetworkRequest) (response *ModifyConsoleNetworkResponse, err error) {
    return c.ModifyConsoleNetworkWithContext(context.Background(), request)
}

// ModifyConsoleNetwork
// 修改网关实例Konga网络配置
//
// 可能返回的错误码:
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_UPDATEERROR = "InvalidParameterValue.UpdateError"
//  MISSINGPARAMETER_MISSPARAMETER = "MissingParameter.MissParameter"
//  OPERATIONDENIED_OPERATIONDENIED = "OperationDenied.OperationDenied"
//  RESOURCENOTFOUND_FORBIDDEN = "ResourceNotFound.Forbidden"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) ModifyConsoleNetworkWithContext(ctx context.Context, request *ModifyConsoleNetworkRequest) (response *ModifyConsoleNetworkResponse, err error) {
    if request == nil {
        request = NewModifyConsoleNetworkRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyConsoleNetwork require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyConsoleNetworkResponse()
    err = c.Send(request, response)
    return
}

func NewModifyGovernanceAliasRequest() (request *ModifyGovernanceAliasRequest) {
    request = &ModifyGovernanceAliasRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "ModifyGovernanceAlias")
    
    
    return
}

func NewModifyGovernanceAliasResponse() (response *ModifyGovernanceAliasResponse) {
    response = &ModifyGovernanceAliasResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyGovernanceAlias
// 修改治理中心服务别名
//
// 可能返回的错误码:
//  INTERNALERROR_GOVERNANCEFAILURE = "InternalError.GovernanceFailure"
//  INTERNALERROR_IOERROR = "InternalError.IOError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_CREATEERROR = "InvalidParameterValue.CreateError"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyGovernanceAlias(request *ModifyGovernanceAliasRequest) (response *ModifyGovernanceAliasResponse, err error) {
    return c.ModifyGovernanceAliasWithContext(context.Background(), request)
}

// ModifyGovernanceAlias
// 修改治理中心服务别名
//
// 可能返回的错误码:
//  INTERNALERROR_GOVERNANCEFAILURE = "InternalError.GovernanceFailure"
//  INTERNALERROR_IOERROR = "InternalError.IOError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_CREATEERROR = "InvalidParameterValue.CreateError"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyGovernanceAliasWithContext(ctx context.Context, request *ModifyGovernanceAliasRequest) (response *ModifyGovernanceAliasResponse, err error) {
    if request == nil {
        request = NewModifyGovernanceAliasRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyGovernanceAlias require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyGovernanceAliasResponse()
    err = c.Send(request, response)
    return
}

func NewModifyGovernanceInstancesRequest() (request *ModifyGovernanceInstancesRequest) {
    request = &ModifyGovernanceInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "ModifyGovernanceInstances")
    
    
    return
}

func NewModifyGovernanceInstancesResponse() (response *ModifyGovernanceInstancesResponse) {
    response = &ModifyGovernanceInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyGovernanceInstances
// 修改治理中心服务实例
//
// 可能返回的错误码:
//  INTERNALERROR_GOVERNANCEFAILURE = "InternalError.GovernanceFailure"
//  INTERNALERROR_IOERROR = "InternalError.IOError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_UPDATEERROR = "InvalidParameterValue.UpdateError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyGovernanceInstances(request *ModifyGovernanceInstancesRequest) (response *ModifyGovernanceInstancesResponse, err error) {
    return c.ModifyGovernanceInstancesWithContext(context.Background(), request)
}

// ModifyGovernanceInstances
// 修改治理中心服务实例
//
// 可能返回的错误码:
//  INTERNALERROR_GOVERNANCEFAILURE = "InternalError.GovernanceFailure"
//  INTERNALERROR_IOERROR = "InternalError.IOError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_UPDATEERROR = "InvalidParameterValue.UpdateError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyGovernanceInstancesWithContext(ctx context.Context, request *ModifyGovernanceInstancesRequest) (response *ModifyGovernanceInstancesResponse, err error) {
    if request == nil {
        request = NewModifyGovernanceInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyGovernanceInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyGovernanceInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewModifyGovernanceNamespacesRequest() (request *ModifyGovernanceNamespacesRequest) {
    request = &ModifyGovernanceNamespacesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "ModifyGovernanceNamespaces")
    
    
    return
}

func NewModifyGovernanceNamespacesResponse() (response *ModifyGovernanceNamespacesResponse) {
    response = &ModifyGovernanceNamespacesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyGovernanceNamespaces
// 修改治理中心命名空间
//
// 可能返回的错误码:
//  INTERNALERROR_GOVERNANCEFAILURE = "InternalError.GovernanceFailure"
//  INTERNALERROR_IOERROR = "InternalError.IOError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_UPDATEERROR = "InvalidParameterValue.UpdateError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyGovernanceNamespaces(request *ModifyGovernanceNamespacesRequest) (response *ModifyGovernanceNamespacesResponse, err error) {
    return c.ModifyGovernanceNamespacesWithContext(context.Background(), request)
}

// ModifyGovernanceNamespaces
// 修改治理中心命名空间
//
// 可能返回的错误码:
//  INTERNALERROR_GOVERNANCEFAILURE = "InternalError.GovernanceFailure"
//  INTERNALERROR_IOERROR = "InternalError.IOError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_UPDATEERROR = "InvalidParameterValue.UpdateError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyGovernanceNamespacesWithContext(ctx context.Context, request *ModifyGovernanceNamespacesRequest) (response *ModifyGovernanceNamespacesResponse, err error) {
    if request == nil {
        request = NewModifyGovernanceNamespacesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyGovernanceNamespaces require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyGovernanceNamespacesResponse()
    err = c.Send(request, response)
    return
}

func NewModifyGovernanceServicesRequest() (request *ModifyGovernanceServicesRequest) {
    request = &ModifyGovernanceServicesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "ModifyGovernanceServices")
    
    
    return
}

func NewModifyGovernanceServicesResponse() (response *ModifyGovernanceServicesResponse) {
    response = &ModifyGovernanceServicesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyGovernanceServices
// 修改治理中心服务
//
// 可能返回的错误码:
//  INTERNALERROR_GOVERNANCEFAILURE = "InternalError.GovernanceFailure"
//  INTERNALERROR_IOERROR = "InternalError.IOError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_UPDATEERROR = "InvalidParameterValue.UpdateError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyGovernanceServices(request *ModifyGovernanceServicesRequest) (response *ModifyGovernanceServicesResponse, err error) {
    return c.ModifyGovernanceServicesWithContext(context.Background(), request)
}

// ModifyGovernanceServices
// 修改治理中心服务
//
// 可能返回的错误码:
//  INTERNALERROR_GOVERNANCEFAILURE = "InternalError.GovernanceFailure"
//  INTERNALERROR_IOERROR = "InternalError.IOError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_UPDATEERROR = "InvalidParameterValue.UpdateError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyGovernanceServicesWithContext(ctx context.Context, request *ModifyGovernanceServicesRequest) (response *ModifyGovernanceServicesResponse, err error) {
    if request == nil {
        request = NewModifyGovernanceServicesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyGovernanceServices require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyGovernanceServicesResponse()
    err = c.Send(request, response)
    return
}

func NewModifyNativeGatewayServerGroupRequest() (request *ModifyNativeGatewayServerGroupRequest) {
    request = &ModifyNativeGatewayServerGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "ModifyNativeGatewayServerGroup")
    
    
    return
}

func NewModifyNativeGatewayServerGroupResponse() (response *ModifyNativeGatewayServerGroupResponse) {
    response = &ModifyNativeGatewayServerGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyNativeGatewayServerGroup
// 修改云原生API网关实例分组基础信息
//
// 可能返回的错误码:
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_DESCRIPTION = "InvalidParameterValue.Description"
//  INVALIDPARAMETERVALUE_GATEWAYID = "InvalidParameterValue.GatewayId"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  MISSINGPARAMETER_MISSPARAMETER = "MissingParameter.MissParameter"
//  OPERATIONDENIED_OPERATIONDENIED = "OperationDenied.OperationDenied"
//  RESOURCENOTFOUND_FORBIDDEN = "ResourceNotFound.Forbidden"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) ModifyNativeGatewayServerGroup(request *ModifyNativeGatewayServerGroupRequest) (response *ModifyNativeGatewayServerGroupResponse, err error) {
    return c.ModifyNativeGatewayServerGroupWithContext(context.Background(), request)
}

// ModifyNativeGatewayServerGroup
// 修改云原生API网关实例分组基础信息
//
// 可能返回的错误码:
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_DESCRIPTION = "InvalidParameterValue.Description"
//  INVALIDPARAMETERVALUE_GATEWAYID = "InvalidParameterValue.GatewayId"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  MISSINGPARAMETER_MISSPARAMETER = "MissingParameter.MissParameter"
//  OPERATIONDENIED_OPERATIONDENIED = "OperationDenied.OperationDenied"
//  RESOURCENOTFOUND_FORBIDDEN = "ResourceNotFound.Forbidden"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) ModifyNativeGatewayServerGroupWithContext(ctx context.Context, request *ModifyNativeGatewayServerGroupRequest) (response *ModifyNativeGatewayServerGroupResponse, err error) {
    if request == nil {
        request = NewModifyNativeGatewayServerGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyNativeGatewayServerGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyNativeGatewayServerGroupResponse()
    err = c.Send(request, response)
    return
}

func NewModifyNetworkAccessStrategyRequest() (request *ModifyNetworkAccessStrategyRequest) {
    request = &ModifyNetworkAccessStrategyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "ModifyNetworkAccessStrategy")
    
    
    return
}

func NewModifyNetworkAccessStrategyResponse() (response *ModifyNetworkAccessStrategyResponse) {
    response = &ModifyNetworkAccessStrategyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyNetworkAccessStrategy
// 修改云原生API网关实例Kong访问策略，支持白名单或者黑名单。
//
// 可能返回的错误码:
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_UPDATEERROR = "InvalidParameterValue.UpdateError"
//  MISSINGPARAMETER_MISSPARAMETER = "MissingParameter.MissParameter"
//  OPERATIONDENIED_OPERATIONDENIED = "OperationDenied.OperationDenied"
//  RESOURCENOTFOUND_FORBIDDEN = "ResourceNotFound.Forbidden"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) ModifyNetworkAccessStrategy(request *ModifyNetworkAccessStrategyRequest) (response *ModifyNetworkAccessStrategyResponse, err error) {
    return c.ModifyNetworkAccessStrategyWithContext(context.Background(), request)
}

// ModifyNetworkAccessStrategy
// 修改云原生API网关实例Kong访问策略，支持白名单或者黑名单。
//
// 可能返回的错误码:
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_UPDATEERROR = "InvalidParameterValue.UpdateError"
//  MISSINGPARAMETER_MISSPARAMETER = "MissingParameter.MissParameter"
//  OPERATIONDENIED_OPERATIONDENIED = "OperationDenied.OperationDenied"
//  RESOURCENOTFOUND_FORBIDDEN = "ResourceNotFound.Forbidden"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) ModifyNetworkAccessStrategyWithContext(ctx context.Context, request *ModifyNetworkAccessStrategyRequest) (response *ModifyNetworkAccessStrategyResponse, err error) {
    if request == nil {
        request = NewModifyNetworkAccessStrategyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyNetworkAccessStrategy require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyNetworkAccessStrategyResponse()
    err = c.Send(request, response)
    return
}

func NewModifyNetworkBasicInfoRequest() (request *ModifyNetworkBasicInfoRequest) {
    request = &ModifyNetworkBasicInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "ModifyNetworkBasicInfo")
    
    
    return
}

func NewModifyNetworkBasicInfoResponse() (response *ModifyNetworkBasicInfoResponse) {
    response = &ModifyNetworkBasicInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyNetworkBasicInfo
// 修改云原生API网关实例网络基本信息，例如带宽以及描述，只支持修改客户端公网/内网的信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_UPDATEERROR = "InvalidParameterValue.UpdateError"
//  MISSINGPARAMETER_MISSPARAMETER = "MissingParameter.MissParameter"
//  OPERATIONDENIED_OPERATIONDENIED = "OperationDenied.OperationDenied"
//  RESOURCENOTFOUND_FORBIDDEN = "ResourceNotFound.Forbidden"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) ModifyNetworkBasicInfo(request *ModifyNetworkBasicInfoRequest) (response *ModifyNetworkBasicInfoResponse, err error) {
    return c.ModifyNetworkBasicInfoWithContext(context.Background(), request)
}

// ModifyNetworkBasicInfo
// 修改云原生API网关实例网络基本信息，例如带宽以及描述，只支持修改客户端公网/内网的信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_UPDATEERROR = "InvalidParameterValue.UpdateError"
//  MISSINGPARAMETER_MISSPARAMETER = "MissingParameter.MissParameter"
//  OPERATIONDENIED_OPERATIONDENIED = "OperationDenied.OperationDenied"
//  RESOURCENOTFOUND_FORBIDDEN = "ResourceNotFound.Forbidden"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) ModifyNetworkBasicInfoWithContext(ctx context.Context, request *ModifyNetworkBasicInfoRequest) (response *ModifyNetworkBasicInfoResponse, err error) {
    if request == nil {
        request = NewModifyNetworkBasicInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyNetworkBasicInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyNetworkBasicInfoResponse()
    err = c.Send(request, response)
    return
}

func NewModifyUpstreamNodeStatusRequest() (request *ModifyUpstreamNodeStatusRequest) {
    request = &ModifyUpstreamNodeStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "ModifyUpstreamNodeStatus")
    
    
    return
}

func NewModifyUpstreamNodeStatusResponse() (response *ModifyUpstreamNodeStatusResponse) {
    response = &ModifyUpstreamNodeStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyUpstreamNodeStatus
// 修改云原生网关上游实例节点健康状态
//
// 可能返回的错误码:
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_UPDATEERROR = "InvalidParameterValue.UpdateError"
//  MISSINGPARAMETER_MISSPARAMETER = "MissingParameter.MissParameter"
//  OPERATIONDENIED_OPERATIONDENIED = "OperationDenied.OperationDenied"
//  RESOURCENOTFOUND_FORBIDDEN = "ResourceNotFound.Forbidden"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) ModifyUpstreamNodeStatus(request *ModifyUpstreamNodeStatusRequest) (response *ModifyUpstreamNodeStatusResponse, err error) {
    return c.ModifyUpstreamNodeStatusWithContext(context.Background(), request)
}

// ModifyUpstreamNodeStatus
// 修改云原生网关上游实例节点健康状态
//
// 可能返回的错误码:
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_UPDATEERROR = "InvalidParameterValue.UpdateError"
//  MISSINGPARAMETER_MISSPARAMETER = "MissingParameter.MissParameter"
//  OPERATIONDENIED_OPERATIONDENIED = "OperationDenied.OperationDenied"
//  RESOURCENOTFOUND_FORBIDDEN = "ResourceNotFound.Forbidden"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) ModifyUpstreamNodeStatusWithContext(ctx context.Context, request *ModifyUpstreamNodeStatusRequest) (response *ModifyUpstreamNodeStatusResponse, err error) {
    if request == nil {
        request = NewModifyUpstreamNodeStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyUpstreamNodeStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyUpstreamNodeStatusResponse()
    err = c.Send(request, response)
    return
}

func NewOpenWafProtectionRequest() (request *OpenWafProtectionRequest) {
    request = &OpenWafProtectionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "OpenWafProtection")
    
    
    return
}

func NewOpenWafProtectionResponse() (response *OpenWafProtectionResponse) {
    response = &OpenWafProtectionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// OpenWafProtection
// 开启 WAF 防护
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  FAILEDOPERATION_VPC = "FailedOperation.Vpc"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_OPERATIONFAILED = "InternalError.OperationFailed"
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
func (c *Client) OpenWafProtection(request *OpenWafProtectionRequest) (response *OpenWafProtectionResponse, err error) {
    return c.OpenWafProtectionWithContext(context.Background(), request)
}

// OpenWafProtection
// 开启 WAF 防护
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  FAILEDOPERATION_VPC = "FailedOperation.Vpc"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_OPERATIONFAILED = "InternalError.OperationFailed"
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
func (c *Client) OpenWafProtectionWithContext(ctx context.Context, request *OpenWafProtectionRequest) (response *OpenWafProtectionResponse, err error) {
    if request == nil {
        request = NewOpenWafProtectionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("OpenWafProtection require credential")
    }

    request.SetContext(ctx)
    
    response = NewOpenWafProtectionResponse()
    err = c.Send(request, response)
    return
}

func NewPublishConfigFilesRequest() (request *PublishConfigFilesRequest) {
    request = &PublishConfigFilesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "PublishConfigFiles")
    
    
    return
}

func NewPublishConfigFilesResponse() (response *PublishConfigFilesResponse) {
    response = &PublishConfigFilesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// PublishConfigFiles
// 发布配置文件
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GOVERNANCEFAILURE = "InternalError.GovernanceFailure"
//  INTERNALERROR_IOERROR = "InternalError.IOError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
func (c *Client) PublishConfigFiles(request *PublishConfigFilesRequest) (response *PublishConfigFilesResponse, err error) {
    return c.PublishConfigFilesWithContext(context.Background(), request)
}

// PublishConfigFiles
// 发布配置文件
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GOVERNANCEFAILURE = "InternalError.GovernanceFailure"
//  INTERNALERROR_IOERROR = "InternalError.IOError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
func (c *Client) PublishConfigFilesWithContext(ctx context.Context, request *PublishConfigFilesRequest) (response *PublishConfigFilesResponse, err error) {
    if request == nil {
        request = NewPublishConfigFilesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("PublishConfigFiles require credential")
    }

    request.SetContext(ctx)
    
    response = NewPublishConfigFilesResponse()
    err = c.Send(request, response)
    return
}

func NewRestartSREInstanceRequest() (request *RestartSREInstanceRequest) {
    request = &RestartSREInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "RestartSREInstance")
    
    
    return
}

func NewRestartSREInstanceResponse() (response *RestartSREInstanceResponse) {
    response = &RestartSREInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RestartSREInstance
// 重启微服务引擎实例
//
// 可能返回的错误码:
//  INTERNALERROR_OPERATIONFAILED = "InternalError.OperationFailed"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) RestartSREInstance(request *RestartSREInstanceRequest) (response *RestartSREInstanceResponse, err error) {
    return c.RestartSREInstanceWithContext(context.Background(), request)
}

// RestartSREInstance
// 重启微服务引擎实例
//
// 可能返回的错误码:
//  INTERNALERROR_OPERATIONFAILED = "InternalError.OperationFailed"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) RestartSREInstanceWithContext(ctx context.Context, request *RestartSREInstanceRequest) (response *RestartSREInstanceResponse, err error) {
    if request == nil {
        request = NewRestartSREInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RestartSREInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewRestartSREInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewRollbackConfigFileReleasesRequest() (request *RollbackConfigFileReleasesRequest) {
    request = &RollbackConfigFileReleasesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "RollbackConfigFileReleases")
    
    
    return
}

func NewRollbackConfigFileReleasesResponse() (response *RollbackConfigFileReleasesResponse) {
    response = &RollbackConfigFileReleasesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RollbackConfigFileReleases
// 回滚配置发布
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GOVERNANCEFAILURE = "InternalError.GovernanceFailure"
//  INTERNALERROR_IOERROR = "InternalError.IOError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_QUERYERROR = "InvalidParameterValue.QueryError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) RollbackConfigFileReleases(request *RollbackConfigFileReleasesRequest) (response *RollbackConfigFileReleasesResponse, err error) {
    return c.RollbackConfigFileReleasesWithContext(context.Background(), request)
}

// RollbackConfigFileReleases
// 回滚配置发布
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GOVERNANCEFAILURE = "InternalError.GovernanceFailure"
//  INTERNALERROR_IOERROR = "InternalError.IOError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_QUERYERROR = "InvalidParameterValue.QueryError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) RollbackConfigFileReleasesWithContext(ctx context.Context, request *RollbackConfigFileReleasesRequest) (response *RollbackConfigFileReleasesResponse, err error) {
    if request == nil {
        request = NewRollbackConfigFileReleasesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RollbackConfigFileReleases require credential")
    }

    request.SetContext(ctx)
    
    response = NewRollbackConfigFileReleasesResponse()
    err = c.Send(request, response)
    return
}

func NewUnbindAutoScalerResourceStrategyFromGroupsRequest() (request *UnbindAutoScalerResourceStrategyFromGroupsRequest) {
    request = &UnbindAutoScalerResourceStrategyFromGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "UnbindAutoScalerResourceStrategyFromGroups")
    
    
    return
}

func NewUnbindAutoScalerResourceStrategyFromGroupsResponse() (response *UnbindAutoScalerResourceStrategyFromGroupsResponse) {
    response = &UnbindAutoScalerResourceStrategyFromGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UnbindAutoScalerResourceStrategyFromGroups
// 弹性伸缩策略批量解绑网关分组
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMNOAUTH = "InternalError.CamNoAuth"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INTERNALERROR_OPERATIONFAILED = "InternalError.OperationFailed"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_GATEWAYID = "InvalidParameterValue.GatewayId"
//  INVALIDPARAMETERVALUE_QUERYERROR = "InvalidParameterValue.QueryError"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_MISSPARAMETER = "MissingParameter.MissParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_OPERATIONDENIED = "OperationDenied.OperationDenied"
//  RESOURCENOTFOUND_FORBIDDEN = "ResourceNotFound.Forbidden"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) UnbindAutoScalerResourceStrategyFromGroups(request *UnbindAutoScalerResourceStrategyFromGroupsRequest) (response *UnbindAutoScalerResourceStrategyFromGroupsResponse, err error) {
    return c.UnbindAutoScalerResourceStrategyFromGroupsWithContext(context.Background(), request)
}

// UnbindAutoScalerResourceStrategyFromGroups
// 弹性伸缩策略批量解绑网关分组
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMNOAUTH = "InternalError.CamNoAuth"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INTERNALERROR_OPERATIONFAILED = "InternalError.OperationFailed"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_GATEWAYID = "InvalidParameterValue.GatewayId"
//  INVALIDPARAMETERVALUE_QUERYERROR = "InvalidParameterValue.QueryError"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_MISSPARAMETER = "MissingParameter.MissParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_OPERATIONDENIED = "OperationDenied.OperationDenied"
//  RESOURCENOTFOUND_FORBIDDEN = "ResourceNotFound.Forbidden"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) UnbindAutoScalerResourceStrategyFromGroupsWithContext(ctx context.Context, request *UnbindAutoScalerResourceStrategyFromGroupsRequest) (response *UnbindAutoScalerResourceStrategyFromGroupsResponse, err error) {
    if request == nil {
        request = NewUnbindAutoScalerResourceStrategyFromGroupsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UnbindAutoScalerResourceStrategyFromGroups require credential")
    }

    request.SetContext(ctx)
    
    response = NewUnbindAutoScalerResourceStrategyFromGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateCloudNativeAPIGatewayCertificateInfoRequest() (request *UpdateCloudNativeAPIGatewayCertificateInfoRequest) {
    request = &UpdateCloudNativeAPIGatewayCertificateInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "UpdateCloudNativeAPIGatewayCertificateInfo")
    
    
    return
}

func NewUpdateCloudNativeAPIGatewayCertificateInfoResponse() (response *UpdateCloudNativeAPIGatewayCertificateInfoResponse) {
    response = &UpdateCloudNativeAPIGatewayCertificateInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateCloudNativeAPIGatewayCertificateInfo
// 修改云原生网关证书信息
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
//  LIMITEXCEEDED_LBDOMAINS = "LimitExceeded.LBDomains"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_CAMPASSROLENOTEXIST = "UnauthorizedOperation.CamPassRoleNotExist"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) UpdateCloudNativeAPIGatewayCertificateInfo(request *UpdateCloudNativeAPIGatewayCertificateInfoRequest) (response *UpdateCloudNativeAPIGatewayCertificateInfoResponse, err error) {
    return c.UpdateCloudNativeAPIGatewayCertificateInfoWithContext(context.Background(), request)
}

// UpdateCloudNativeAPIGatewayCertificateInfo
// 修改云原生网关证书信息
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
//  LIMITEXCEEDED_LBDOMAINS = "LimitExceeded.LBDomains"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_CAMPASSROLENOTEXIST = "UnauthorizedOperation.CamPassRoleNotExist"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) UpdateCloudNativeAPIGatewayCertificateInfoWithContext(ctx context.Context, request *UpdateCloudNativeAPIGatewayCertificateInfoRequest) (response *UpdateCloudNativeAPIGatewayCertificateInfoResponse, err error) {
    if request == nil {
        request = NewUpdateCloudNativeAPIGatewayCertificateInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateCloudNativeAPIGatewayCertificateInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateCloudNativeAPIGatewayCertificateInfoResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateCloudNativeAPIGatewaySpecRequest() (request *UpdateCloudNativeAPIGatewaySpecRequest) {
    request = &UpdateCloudNativeAPIGatewaySpecRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "UpdateCloudNativeAPIGatewaySpec")
    
    
    return
}

func NewUpdateCloudNativeAPIGatewaySpecResponse() (response *UpdateCloudNativeAPIGatewaySpecResponse) {
    response = &UpdateCloudNativeAPIGatewaySpecResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateCloudNativeAPIGatewaySpec
// 修改云原生API网关实例的节点规格信息，例如节点扩缩容或者升降配
//
// 可能返回的错误码:
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  MISSINGPARAMETER_MISSPARAMETER = "MissingParameter.MissParameter"
//  OPERATIONDENIED_OPERATIONDENIED = "OperationDenied.OperationDenied"
//  RESOURCENOTFOUND_FORBIDDEN = "ResourceNotFound.Forbidden"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) UpdateCloudNativeAPIGatewaySpec(request *UpdateCloudNativeAPIGatewaySpecRequest) (response *UpdateCloudNativeAPIGatewaySpecResponse, err error) {
    return c.UpdateCloudNativeAPIGatewaySpecWithContext(context.Background(), request)
}

// UpdateCloudNativeAPIGatewaySpec
// 修改云原生API网关实例的节点规格信息，例如节点扩缩容或者升降配
//
// 可能返回的错误码:
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  MISSINGPARAMETER_MISSPARAMETER = "MissingParameter.MissParameter"
//  OPERATIONDENIED_OPERATIONDENIED = "OperationDenied.OperationDenied"
//  RESOURCENOTFOUND_FORBIDDEN = "ResourceNotFound.Forbidden"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) UpdateCloudNativeAPIGatewaySpecWithContext(ctx context.Context, request *UpdateCloudNativeAPIGatewaySpecRequest) (response *UpdateCloudNativeAPIGatewaySpecResponse, err error) {
    if request == nil {
        request = NewUpdateCloudNativeAPIGatewaySpecRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateCloudNativeAPIGatewaySpec require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateCloudNativeAPIGatewaySpecResponse()
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

func NewUpdateUpstreamHealthCheckConfigRequest() (request *UpdateUpstreamHealthCheckConfigRequest) {
    request = &UpdateUpstreamHealthCheckConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "UpdateUpstreamHealthCheckConfig")
    
    
    return
}

func NewUpdateUpstreamHealthCheckConfigResponse() (response *UpdateUpstreamHealthCheckConfigResponse) {
    response = &UpdateUpstreamHealthCheckConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateUpstreamHealthCheckConfig
// 更新云原生网关健康检查配置
//
// 可能返回的错误码:
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
func (c *Client) UpdateUpstreamHealthCheckConfig(request *UpdateUpstreamHealthCheckConfigRequest) (response *UpdateUpstreamHealthCheckConfigResponse, err error) {
    return c.UpdateUpstreamHealthCheckConfigWithContext(context.Background(), request)
}

// UpdateUpstreamHealthCheckConfig
// 更新云原生网关健康检查配置
//
// 可能返回的错误码:
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
func (c *Client) UpdateUpstreamHealthCheckConfigWithContext(ctx context.Context, request *UpdateUpstreamHealthCheckConfigRequest) (response *UpdateUpstreamHealthCheckConfigResponse, err error) {
    if request == nil {
        request = NewUpdateUpstreamHealthCheckConfigRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateUpstreamHealthCheckConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateUpstreamHealthCheckConfigResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateUpstreamTargetsRequest() (request *UpdateUpstreamTargetsRequest) {
    request = &UpdateUpstreamTargetsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "UpdateUpstreamTargets")
    
    
    return
}

func NewUpdateUpstreamTargetsResponse() (response *UpdateUpstreamTargetsResponse) {
    response = &UpdateUpstreamTargetsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateUpstreamTargets
// 更新网关上游实例列表，仅支持IPList服务类型
//
// 可能返回的错误码:
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
func (c *Client) UpdateUpstreamTargets(request *UpdateUpstreamTargetsRequest) (response *UpdateUpstreamTargetsResponse, err error) {
    return c.UpdateUpstreamTargetsWithContext(context.Background(), request)
}

// UpdateUpstreamTargets
// 更新网关上游实例列表，仅支持IPList服务类型
//
// 可能返回的错误码:
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
func (c *Client) UpdateUpstreamTargetsWithContext(ctx context.Context, request *UpdateUpstreamTargetsRequest) (response *UpdateUpstreamTargetsResponse, err error) {
    if request == nil {
        request = NewUpdateUpstreamTargetsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateUpstreamTargets require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateUpstreamTargetsResponse()
    err = c.Send(request, response)
    return
}
