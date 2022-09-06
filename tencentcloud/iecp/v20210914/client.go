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

package v20210914

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2021-09-14"

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


func NewApplyMarketComponentRequest() (request *ApplyMarketComponentRequest) {
    request = &ApplyMarketComponentRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iecp", APIVersion, "ApplyMarketComponent")
    
    
    return
}

func NewApplyMarketComponentResponse() (response *ApplyMarketComponentResponse) {
    response = &ApplyMarketComponentResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ApplyMarketComponent
// 从组件市场选中组件并添加到应用模板列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ApplyMarketComponent(request *ApplyMarketComponentRequest) (response *ApplyMarketComponentResponse, err error) {
    return c.ApplyMarketComponentWithContext(context.Background(), request)
}

// ApplyMarketComponent
// 从组件市场选中组件并添加到应用模板列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ApplyMarketComponentWithContext(ctx context.Context, request *ApplyMarketComponentRequest) (response *ApplyMarketComponentResponse, err error) {
    if request == nil {
        request = NewApplyMarketComponentRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ApplyMarketComponent require credential")
    }

    request.SetContext(ctx)
    
    response = NewApplyMarketComponentResponse()
    err = c.Send(request, response)
    return
}

func NewBuildMessageRouteRequest() (request *BuildMessageRouteRequest) {
    request = &BuildMessageRouteRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iecp", APIVersion, "BuildMessageRoute")
    
    
    return
}

func NewBuildMessageRouteResponse() (response *BuildMessageRouteResponse) {
    response = &BuildMessageRouteResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// BuildMessageRoute
// 建立消息路由
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) BuildMessageRoute(request *BuildMessageRouteRequest) (response *BuildMessageRouteResponse, err error) {
    return c.BuildMessageRouteWithContext(context.Background(), request)
}

// BuildMessageRoute
// 建立消息路由
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) BuildMessageRouteWithContext(ctx context.Context, request *BuildMessageRouteRequest) (response *BuildMessageRouteResponse, err error) {
    if request == nil {
        request = NewBuildMessageRouteRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("BuildMessageRoute require credential")
    }

    request.SetContext(ctx)
    
    response = NewBuildMessageRouteResponse()
    err = c.Send(request, response)
    return
}

func NewCreateApplicationVisualizationRequest() (request *CreateApplicationVisualizationRequest) {
    request = &CreateApplicationVisualizationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iecp", APIVersion, "CreateApplicationVisualization")
    
    
    return
}

func NewCreateApplicationVisualizationResponse() (response *CreateApplicationVisualizationResponse) {
    response = &CreateApplicationVisualizationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateApplicationVisualization
// 创建可视化创建应用模板
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateApplicationVisualization(request *CreateApplicationVisualizationRequest) (response *CreateApplicationVisualizationResponse, err error) {
    return c.CreateApplicationVisualizationWithContext(context.Background(), request)
}

// CreateApplicationVisualization
// 创建可视化创建应用模板
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateApplicationVisualizationWithContext(ctx context.Context, request *CreateApplicationVisualizationRequest) (response *CreateApplicationVisualizationResponse, err error) {
    if request == nil {
        request = NewCreateApplicationVisualizationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateApplicationVisualization require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateApplicationVisualizationResponse()
    err = c.Send(request, response)
    return
}

func NewCreateConfigMapRequest() (request *CreateConfigMapRequest) {
    request = &CreateConfigMapRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iecp", APIVersion, "CreateConfigMap")
    
    
    return
}

func NewCreateConfigMapResponse() (response *CreateConfigMapResponse) {
    response = &CreateConfigMapResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateConfigMap
// 创建ConfigMap
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateConfigMap(request *CreateConfigMapRequest) (response *CreateConfigMapResponse, err error) {
    return c.CreateConfigMapWithContext(context.Background(), request)
}

// CreateConfigMap
// 创建ConfigMap
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateConfigMapWithContext(ctx context.Context, request *CreateConfigMapRequest) (response *CreateConfigMapResponse, err error) {
    if request == nil {
        request = NewCreateConfigMapRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateConfigMap require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateConfigMapResponse()
    err = c.Send(request, response)
    return
}

func NewCreateEdgeNodeRequest() (request *CreateEdgeNodeRequest) {
    request = &CreateEdgeNodeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iecp", APIVersion, "CreateEdgeNode")
    
    
    return
}

func NewCreateEdgeNodeResponse() (response *CreateEdgeNodeResponse) {
    response = &CreateEdgeNodeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateEdgeNode
// 创建边缘节点
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateEdgeNode(request *CreateEdgeNodeRequest) (response *CreateEdgeNodeResponse, err error) {
    return c.CreateEdgeNodeWithContext(context.Background(), request)
}

// CreateEdgeNode
// 创建边缘节点
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateEdgeNodeWithContext(ctx context.Context, request *CreateEdgeNodeRequest) (response *CreateEdgeNodeResponse, err error) {
    if request == nil {
        request = NewCreateEdgeNodeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateEdgeNode require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateEdgeNodeResponse()
    err = c.Send(request, response)
    return
}

func NewCreateEdgeNodeBatchRequest() (request *CreateEdgeNodeBatchRequest) {
    request = &CreateEdgeNodeBatchRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iecp", APIVersion, "CreateEdgeNodeBatch")
    
    
    return
}

func NewCreateEdgeNodeBatchResponse() (response *CreateEdgeNodeBatchResponse) {
    response = &CreateEdgeNodeBatchResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateEdgeNodeBatch
// 批量预注册节点
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateEdgeNodeBatch(request *CreateEdgeNodeBatchRequest) (response *CreateEdgeNodeBatchResponse, err error) {
    return c.CreateEdgeNodeBatchWithContext(context.Background(), request)
}

// CreateEdgeNodeBatch
// 批量预注册节点
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateEdgeNodeBatchWithContext(ctx context.Context, request *CreateEdgeNodeBatchRequest) (response *CreateEdgeNodeBatchResponse, err error) {
    if request == nil {
        request = NewCreateEdgeNodeBatchRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateEdgeNodeBatch require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateEdgeNodeBatchResponse()
    err = c.Send(request, response)
    return
}

func NewCreateEdgeNodeGroupRequest() (request *CreateEdgeNodeGroupRequest) {
    request = &CreateEdgeNodeGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iecp", APIVersion, "CreateEdgeNodeGroup")
    
    
    return
}

func NewCreateEdgeNodeGroupResponse() (response *CreateEdgeNodeGroupResponse) {
    response = &CreateEdgeNodeGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateEdgeNodeGroup
// 创建边缘单元NodeGroup
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateEdgeNodeGroup(request *CreateEdgeNodeGroupRequest) (response *CreateEdgeNodeGroupResponse, err error) {
    return c.CreateEdgeNodeGroupWithContext(context.Background(), request)
}

// CreateEdgeNodeGroup
// 创建边缘单元NodeGroup
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateEdgeNodeGroupWithContext(ctx context.Context, request *CreateEdgeNodeGroupRequest) (response *CreateEdgeNodeGroupResponse, err error) {
    if request == nil {
        request = NewCreateEdgeNodeGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateEdgeNodeGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateEdgeNodeGroupResponse()
    err = c.Send(request, response)
    return
}

func NewCreateEdgeNodeUnitTemplateRequest() (request *CreateEdgeNodeUnitTemplateRequest) {
    request = &CreateEdgeNodeUnitTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iecp", APIVersion, "CreateEdgeNodeUnitTemplate")
    
    
    return
}

func NewCreateEdgeNodeUnitTemplateResponse() (response *CreateEdgeNodeUnitTemplateResponse) {
    response = &CreateEdgeNodeUnitTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateEdgeNodeUnitTemplate
// 创建边缘单元NodeUnit模板
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateEdgeNodeUnitTemplate(request *CreateEdgeNodeUnitTemplateRequest) (response *CreateEdgeNodeUnitTemplateResponse, err error) {
    return c.CreateEdgeNodeUnitTemplateWithContext(context.Background(), request)
}

// CreateEdgeNodeUnitTemplate
// 创建边缘单元NodeUnit模板
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateEdgeNodeUnitTemplateWithContext(ctx context.Context, request *CreateEdgeNodeUnitTemplateRequest) (response *CreateEdgeNodeUnitTemplateResponse, err error) {
    if request == nil {
        request = NewCreateEdgeNodeUnitTemplateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateEdgeNodeUnitTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateEdgeNodeUnitTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewCreateEdgeUnitApplicationVisualizationRequest() (request *CreateEdgeUnitApplicationVisualizationRequest) {
    request = &CreateEdgeUnitApplicationVisualizationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iecp", APIVersion, "CreateEdgeUnitApplicationVisualization")
    
    
    return
}

func NewCreateEdgeUnitApplicationVisualizationResponse() (response *CreateEdgeUnitApplicationVisualizationResponse) {
    response = &CreateEdgeUnitApplicationVisualizationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateEdgeUnitApplicationVisualization
// 可视化创建应用
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateEdgeUnitApplicationVisualization(request *CreateEdgeUnitApplicationVisualizationRequest) (response *CreateEdgeUnitApplicationVisualizationResponse, err error) {
    return c.CreateEdgeUnitApplicationVisualizationWithContext(context.Background(), request)
}

// CreateEdgeUnitApplicationVisualization
// 可视化创建应用
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateEdgeUnitApplicationVisualizationWithContext(ctx context.Context, request *CreateEdgeUnitApplicationVisualizationRequest) (response *CreateEdgeUnitApplicationVisualizationResponse, err error) {
    if request == nil {
        request = NewCreateEdgeUnitApplicationVisualizationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateEdgeUnitApplicationVisualization require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateEdgeUnitApplicationVisualizationResponse()
    err = c.Send(request, response)
    return
}

func NewCreateEdgeUnitApplicationYamlRequest() (request *CreateEdgeUnitApplicationYamlRequest) {
    request = &CreateEdgeUnitApplicationYamlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iecp", APIVersion, "CreateEdgeUnitApplicationYaml")
    
    
    return
}

func NewCreateEdgeUnitApplicationYamlResponse() (response *CreateEdgeUnitApplicationYamlResponse) {
    response = &CreateEdgeUnitApplicationYamlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateEdgeUnitApplicationYaml
// yaml方式创建应用
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateEdgeUnitApplicationYaml(request *CreateEdgeUnitApplicationYamlRequest) (response *CreateEdgeUnitApplicationYamlResponse, err error) {
    return c.CreateEdgeUnitApplicationYamlWithContext(context.Background(), request)
}

// CreateEdgeUnitApplicationYaml
// yaml方式创建应用
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateEdgeUnitApplicationYamlWithContext(ctx context.Context, request *CreateEdgeUnitApplicationYamlRequest) (response *CreateEdgeUnitApplicationYamlResponse, err error) {
    if request == nil {
        request = NewCreateEdgeUnitApplicationYamlRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateEdgeUnitApplicationYaml require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateEdgeUnitApplicationYamlResponse()
    err = c.Send(request, response)
    return
}

func NewCreateEdgeUnitCloudRequest() (request *CreateEdgeUnitCloudRequest) {
    request = &CreateEdgeUnitCloudRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iecp", APIVersion, "CreateEdgeUnitCloud")
    
    
    return
}

func NewCreateEdgeUnitCloudResponse() (response *CreateEdgeUnitCloudResponse) {
    response = &CreateEdgeUnitCloudResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateEdgeUnitCloud
// 创建边缘单元
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateEdgeUnitCloud(request *CreateEdgeUnitCloudRequest) (response *CreateEdgeUnitCloudResponse, err error) {
    return c.CreateEdgeUnitCloudWithContext(context.Background(), request)
}

// CreateEdgeUnitCloud
// 创建边缘单元
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateEdgeUnitCloudWithContext(ctx context.Context, request *CreateEdgeUnitCloudRequest) (response *CreateEdgeUnitCloudResponse, err error) {
    if request == nil {
        request = NewCreateEdgeUnitCloudRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateEdgeUnitCloud require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateEdgeUnitCloudResponse()
    err = c.Send(request, response)
    return
}

func NewCreateEdgeUnitDevicesRequest() (request *CreateEdgeUnitDevicesRequest) {
    request = &CreateEdgeUnitDevicesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iecp", APIVersion, "CreateEdgeUnitDevices")
    
    
    return
}

func NewCreateEdgeUnitDevicesResponse() (response *CreateEdgeUnitDevicesResponse) {
    response = &CreateEdgeUnitDevicesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateEdgeUnitDevices
// 批量绑定设备到单元
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateEdgeUnitDevices(request *CreateEdgeUnitDevicesRequest) (response *CreateEdgeUnitDevicesResponse, err error) {
    return c.CreateEdgeUnitDevicesWithContext(context.Background(), request)
}

// CreateEdgeUnitDevices
// 批量绑定设备到单元
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateEdgeUnitDevicesWithContext(ctx context.Context, request *CreateEdgeUnitDevicesRequest) (response *CreateEdgeUnitDevicesResponse, err error) {
    if request == nil {
        request = NewCreateEdgeUnitDevicesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateEdgeUnitDevices require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateEdgeUnitDevicesResponse()
    err = c.Send(request, response)
    return
}

func NewCreateIotDeviceRequest() (request *CreateIotDeviceRequest) {
    request = &CreateIotDeviceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iecp", APIVersion, "CreateIotDevice")
    
    
    return
}

func NewCreateIotDeviceResponse() (response *CreateIotDeviceResponse) {
    response = &CreateIotDeviceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateIotDevice
// 创建子设备
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateIotDevice(request *CreateIotDeviceRequest) (response *CreateIotDeviceResponse, err error) {
    return c.CreateIotDeviceWithContext(context.Background(), request)
}

// CreateIotDevice
// 创建子设备
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateIotDeviceWithContext(ctx context.Context, request *CreateIotDeviceRequest) (response *CreateIotDeviceResponse, err error) {
    if request == nil {
        request = NewCreateIotDeviceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateIotDevice require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateIotDeviceResponse()
    err = c.Send(request, response)
    return
}

func NewCreateMessageRouteRequest() (request *CreateMessageRouteRequest) {
    request = &CreateMessageRouteRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iecp", APIVersion, "CreateMessageRoute")
    
    
    return
}

func NewCreateMessageRouteResponse() (response *CreateMessageRouteResponse) {
    response = &CreateMessageRouteResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateMessageRoute
// 创建消息路由
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateMessageRoute(request *CreateMessageRouteRequest) (response *CreateMessageRouteResponse, err error) {
    return c.CreateMessageRouteWithContext(context.Background(), request)
}

// CreateMessageRoute
// 创建消息路由
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateMessageRouteWithContext(ctx context.Context, request *CreateMessageRouteRequest) (response *CreateMessageRouteResponse, err error) {
    if request == nil {
        request = NewCreateMessageRouteRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateMessageRoute require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateMessageRouteResponse()
    err = c.Send(request, response)
    return
}

func NewCreateNamespaceRequest() (request *CreateNamespaceRequest) {
    request = &CreateNamespaceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iecp", APIVersion, "CreateNamespace")
    
    
    return
}

func NewCreateNamespaceResponse() (response *CreateNamespaceResponse) {
    response = &CreateNamespaceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateNamespace
// 创建命名空间
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateNamespace(request *CreateNamespaceRequest) (response *CreateNamespaceResponse, err error) {
    return c.CreateNamespaceWithContext(context.Background(), request)
}

// CreateNamespace
// 创建命名空间
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateNamespaceWithContext(ctx context.Context, request *CreateNamespaceRequest) (response *CreateNamespaceResponse, err error) {
    if request == nil {
        request = NewCreateNamespaceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateNamespace require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateNamespaceResponse()
    err = c.Send(request, response)
    return
}

func NewCreateSecretRequest() (request *CreateSecretRequest) {
    request = &CreateSecretRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iecp", APIVersion, "CreateSecret")
    
    
    return
}

func NewCreateSecretResponse() (response *CreateSecretResponse) {
    response = &CreateSecretResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateSecret
// 创建Secret
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateSecret(request *CreateSecretRequest) (response *CreateSecretResponse, err error) {
    return c.CreateSecretWithContext(context.Background(), request)
}

// CreateSecret
// 创建Secret
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateSecretWithContext(ctx context.Context, request *CreateSecretRequest) (response *CreateSecretResponse, err error) {
    if request == nil {
        request = NewCreateSecretRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateSecret require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateSecretResponse()
    err = c.Send(request, response)
    return
}

func NewCreateUpdateNodeUnitRequest() (request *CreateUpdateNodeUnitRequest) {
    request = &CreateUpdateNodeUnitRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iecp", APIVersion, "CreateUpdateNodeUnit")
    
    
    return
}

func NewCreateUpdateNodeUnitResponse() (response *CreateUpdateNodeUnitResponse) {
    response = &CreateUpdateNodeUnitResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateUpdateNodeUnit
// 创建或更新边缘单元NodeUnit
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateUpdateNodeUnit(request *CreateUpdateNodeUnitRequest) (response *CreateUpdateNodeUnitResponse, err error) {
    return c.CreateUpdateNodeUnitWithContext(context.Background(), request)
}

// CreateUpdateNodeUnit
// 创建或更新边缘单元NodeUnit
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateUpdateNodeUnitWithContext(ctx context.Context, request *CreateUpdateNodeUnitRequest) (response *CreateUpdateNodeUnitResponse, err error) {
    if request == nil {
        request = NewCreateUpdateNodeUnitRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateUpdateNodeUnit require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateUpdateNodeUnitResponse()
    err = c.Send(request, response)
    return
}

func NewCreateUserTokenRequest() (request *CreateUserTokenRequest) {
    request = &CreateUserTokenRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iecp", APIVersion, "CreateUserToken")
    
    
    return
}

func NewCreateUserTokenResponse() (response *CreateUserTokenResponse) {
    response = &CreateUserTokenResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateUserToken
// 创建token
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateUserToken(request *CreateUserTokenRequest) (response *CreateUserTokenResponse, err error) {
    return c.CreateUserTokenWithContext(context.Background(), request)
}

// CreateUserToken
// 创建token
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateUserTokenWithContext(ctx context.Context, request *CreateUserTokenRequest) (response *CreateUserTokenResponse, err error) {
    if request == nil {
        request = NewCreateUserTokenRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateUserToken require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateUserTokenResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteApplicationsRequest() (request *DeleteApplicationsRequest) {
    request = &DeleteApplicationsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iecp", APIVersion, "DeleteApplications")
    
    
    return
}

func NewDeleteApplicationsResponse() (response *DeleteApplicationsResponse) {
    response = &DeleteApplicationsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteApplications
// 删除应用模板
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteApplications(request *DeleteApplicationsRequest) (response *DeleteApplicationsResponse, err error) {
    return c.DeleteApplicationsWithContext(context.Background(), request)
}

// DeleteApplications
// 删除应用模板
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteApplicationsWithContext(ctx context.Context, request *DeleteApplicationsRequest) (response *DeleteApplicationsResponse, err error) {
    if request == nil {
        request = NewDeleteApplicationsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteApplications require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteApplicationsResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteConfigMapRequest() (request *DeleteConfigMapRequest) {
    request = &DeleteConfigMapRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iecp", APIVersion, "DeleteConfigMap")
    
    
    return
}

func NewDeleteConfigMapResponse() (response *DeleteConfigMapResponse) {
    response = &DeleteConfigMapResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteConfigMap
// 删除ConfigMap
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteConfigMap(request *DeleteConfigMapRequest) (response *DeleteConfigMapResponse, err error) {
    return c.DeleteConfigMapWithContext(context.Background(), request)
}

// DeleteConfigMap
// 删除ConfigMap
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteConfigMapWithContext(ctx context.Context, request *DeleteConfigMapRequest) (response *DeleteConfigMapResponse, err error) {
    if request == nil {
        request = NewDeleteConfigMapRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteConfigMap require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteConfigMapResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteEdgeNodeGroupRequest() (request *DeleteEdgeNodeGroupRequest) {
    request = &DeleteEdgeNodeGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iecp", APIVersion, "DeleteEdgeNodeGroup")
    
    
    return
}

func NewDeleteEdgeNodeGroupResponse() (response *DeleteEdgeNodeGroupResponse) {
    response = &DeleteEdgeNodeGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteEdgeNodeGroup
// 删除边缘单元NodeGroup
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteEdgeNodeGroup(request *DeleteEdgeNodeGroupRequest) (response *DeleteEdgeNodeGroupResponse, err error) {
    return c.DeleteEdgeNodeGroupWithContext(context.Background(), request)
}

// DeleteEdgeNodeGroup
// 删除边缘单元NodeGroup
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteEdgeNodeGroupWithContext(ctx context.Context, request *DeleteEdgeNodeGroupRequest) (response *DeleteEdgeNodeGroupResponse, err error) {
    if request == nil {
        request = NewDeleteEdgeNodeGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteEdgeNodeGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteEdgeNodeGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteEdgeNodeUnitTemplatesRequest() (request *DeleteEdgeNodeUnitTemplatesRequest) {
    request = &DeleteEdgeNodeUnitTemplatesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iecp", APIVersion, "DeleteEdgeNodeUnitTemplates")
    
    
    return
}

func NewDeleteEdgeNodeUnitTemplatesResponse() (response *DeleteEdgeNodeUnitTemplatesResponse) {
    response = &DeleteEdgeNodeUnitTemplatesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteEdgeNodeUnitTemplates
// 删除边缘单元NodeUnit模板
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteEdgeNodeUnitTemplates(request *DeleteEdgeNodeUnitTemplatesRequest) (response *DeleteEdgeNodeUnitTemplatesResponse, err error) {
    return c.DeleteEdgeNodeUnitTemplatesWithContext(context.Background(), request)
}

// DeleteEdgeNodeUnitTemplates
// 删除边缘单元NodeUnit模板
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteEdgeNodeUnitTemplatesWithContext(ctx context.Context, request *DeleteEdgeNodeUnitTemplatesRequest) (response *DeleteEdgeNodeUnitTemplatesResponse, err error) {
    if request == nil {
        request = NewDeleteEdgeNodeUnitTemplatesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteEdgeNodeUnitTemplates require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteEdgeNodeUnitTemplatesResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteEdgeNodesRequest() (request *DeleteEdgeNodesRequest) {
    request = &DeleteEdgeNodesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iecp", APIVersion, "DeleteEdgeNodes")
    
    
    return
}

func NewDeleteEdgeNodesResponse() (response *DeleteEdgeNodesResponse) {
    response = &DeleteEdgeNodesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteEdgeNodes
// 批量删除边缘节点
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteEdgeNodes(request *DeleteEdgeNodesRequest) (response *DeleteEdgeNodesResponse, err error) {
    return c.DeleteEdgeNodesWithContext(context.Background(), request)
}

// DeleteEdgeNodes
// 批量删除边缘节点
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteEdgeNodesWithContext(ctx context.Context, request *DeleteEdgeNodesRequest) (response *DeleteEdgeNodesResponse, err error) {
    if request == nil {
        request = NewDeleteEdgeNodesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteEdgeNodes require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteEdgeNodesResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteEdgeUnitApplicationsRequest() (request *DeleteEdgeUnitApplicationsRequest) {
    request = &DeleteEdgeUnitApplicationsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iecp", APIVersion, "DeleteEdgeUnitApplications")
    
    
    return
}

func NewDeleteEdgeUnitApplicationsResponse() (response *DeleteEdgeUnitApplicationsResponse) {
    response = &DeleteEdgeUnitApplicationsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteEdgeUnitApplications
// 删除应用列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteEdgeUnitApplications(request *DeleteEdgeUnitApplicationsRequest) (response *DeleteEdgeUnitApplicationsResponse, err error) {
    return c.DeleteEdgeUnitApplicationsWithContext(context.Background(), request)
}

// DeleteEdgeUnitApplications
// 删除应用列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteEdgeUnitApplicationsWithContext(ctx context.Context, request *DeleteEdgeUnitApplicationsRequest) (response *DeleteEdgeUnitApplicationsResponse, err error) {
    if request == nil {
        request = NewDeleteEdgeUnitApplicationsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteEdgeUnitApplications require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteEdgeUnitApplicationsResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteEdgeUnitCloudRequest() (request *DeleteEdgeUnitCloudRequest) {
    request = &DeleteEdgeUnitCloudRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iecp", APIVersion, "DeleteEdgeUnitCloud")
    
    
    return
}

func NewDeleteEdgeUnitCloudResponse() (response *DeleteEdgeUnitCloudResponse) {
    response = &DeleteEdgeUnitCloudResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteEdgeUnitCloud
// 删除边缘单元
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteEdgeUnitCloud(request *DeleteEdgeUnitCloudRequest) (response *DeleteEdgeUnitCloudResponse, err error) {
    return c.DeleteEdgeUnitCloudWithContext(context.Background(), request)
}

// DeleteEdgeUnitCloud
// 删除边缘单元
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteEdgeUnitCloudWithContext(ctx context.Context, request *DeleteEdgeUnitCloudRequest) (response *DeleteEdgeUnitCloudResponse, err error) {
    if request == nil {
        request = NewDeleteEdgeUnitCloudRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteEdgeUnitCloud require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteEdgeUnitCloudResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteEdgeUnitDeployGridItemRequest() (request *DeleteEdgeUnitDeployGridItemRequest) {
    request = &DeleteEdgeUnitDeployGridItemRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iecp", APIVersion, "DeleteEdgeUnitDeployGridItem")
    
    
    return
}

func NewDeleteEdgeUnitDeployGridItemResponse() (response *DeleteEdgeUnitDeployGridItemResponse) {
    response = &DeleteEdgeUnitDeployGridItemResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteEdgeUnitDeployGridItem
// 重新部署边缘单元指定Grid下应用
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteEdgeUnitDeployGridItem(request *DeleteEdgeUnitDeployGridItemRequest) (response *DeleteEdgeUnitDeployGridItemResponse, err error) {
    return c.DeleteEdgeUnitDeployGridItemWithContext(context.Background(), request)
}

// DeleteEdgeUnitDeployGridItem
// 重新部署边缘单元指定Grid下应用
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteEdgeUnitDeployGridItemWithContext(ctx context.Context, request *DeleteEdgeUnitDeployGridItemRequest) (response *DeleteEdgeUnitDeployGridItemResponse, err error) {
    if request == nil {
        request = NewDeleteEdgeUnitDeployGridItemRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteEdgeUnitDeployGridItem require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteEdgeUnitDeployGridItemResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteEdgeUnitDevicesRequest() (request *DeleteEdgeUnitDevicesRequest) {
    request = &DeleteEdgeUnitDevicesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iecp", APIVersion, "DeleteEdgeUnitDevices")
    
    
    return
}

func NewDeleteEdgeUnitDevicesResponse() (response *DeleteEdgeUnitDevicesResponse) {
    response = &DeleteEdgeUnitDevicesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteEdgeUnitDevices
// 批量解绑单元设备
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteEdgeUnitDevices(request *DeleteEdgeUnitDevicesRequest) (response *DeleteEdgeUnitDevicesResponse, err error) {
    return c.DeleteEdgeUnitDevicesWithContext(context.Background(), request)
}

// DeleteEdgeUnitDevices
// 批量解绑单元设备
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteEdgeUnitDevicesWithContext(ctx context.Context, request *DeleteEdgeUnitDevicesRequest) (response *DeleteEdgeUnitDevicesResponse, err error) {
    if request == nil {
        request = NewDeleteEdgeUnitDevicesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteEdgeUnitDevices require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteEdgeUnitDevicesResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteEdgeUnitPodRequest() (request *DeleteEdgeUnitPodRequest) {
    request = &DeleteEdgeUnitPodRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iecp", APIVersion, "DeleteEdgeUnitPod")
    
    
    return
}

func NewDeleteEdgeUnitPodResponse() (response *DeleteEdgeUnitPodResponse) {
    response = &DeleteEdgeUnitPodResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteEdgeUnitPod
// 删除指定pod
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteEdgeUnitPod(request *DeleteEdgeUnitPodRequest) (response *DeleteEdgeUnitPodResponse, err error) {
    return c.DeleteEdgeUnitPodWithContext(context.Background(), request)
}

// DeleteEdgeUnitPod
// 删除指定pod
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteEdgeUnitPodWithContext(ctx context.Context, request *DeleteEdgeUnitPodRequest) (response *DeleteEdgeUnitPodResponse, err error) {
    if request == nil {
        request = NewDeleteEdgeUnitPodRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteEdgeUnitPod require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteEdgeUnitPodResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteIotDeviceRequest() (request *DeleteIotDeviceRequest) {
    request = &DeleteIotDeviceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iecp", APIVersion, "DeleteIotDevice")
    
    
    return
}

func NewDeleteIotDeviceResponse() (response *DeleteIotDeviceResponse) {
    response = &DeleteIotDeviceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteIotDevice
// 删除设备
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteIotDevice(request *DeleteIotDeviceRequest) (response *DeleteIotDeviceResponse, err error) {
    return c.DeleteIotDeviceWithContext(context.Background(), request)
}

// DeleteIotDevice
// 删除设备
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteIotDeviceWithContext(ctx context.Context, request *DeleteIotDeviceRequest) (response *DeleteIotDeviceResponse, err error) {
    if request == nil {
        request = NewDeleteIotDeviceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteIotDevice require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteIotDeviceResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteIotDeviceBatchRequest() (request *DeleteIotDeviceBatchRequest) {
    request = &DeleteIotDeviceBatchRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iecp", APIVersion, "DeleteIotDeviceBatch")
    
    
    return
}

func NewDeleteIotDeviceBatchResponse() (response *DeleteIotDeviceBatchResponse) {
    response = &DeleteIotDeviceBatchResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteIotDeviceBatch
// 批量删除设备
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteIotDeviceBatch(request *DeleteIotDeviceBatchRequest) (response *DeleteIotDeviceBatchResponse, err error) {
    return c.DeleteIotDeviceBatchWithContext(context.Background(), request)
}

// DeleteIotDeviceBatch
// 批量删除设备
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteIotDeviceBatchWithContext(ctx context.Context, request *DeleteIotDeviceBatchRequest) (response *DeleteIotDeviceBatchResponse, err error) {
    if request == nil {
        request = NewDeleteIotDeviceBatchRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteIotDeviceBatch require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteIotDeviceBatchResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteMessageRouteRequest() (request *DeleteMessageRouteRequest) {
    request = &DeleteMessageRouteRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iecp", APIVersion, "DeleteMessageRoute")
    
    
    return
}

func NewDeleteMessageRouteResponse() (response *DeleteMessageRouteResponse) {
    response = &DeleteMessageRouteResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteMessageRoute
// 删除消息路由
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteMessageRoute(request *DeleteMessageRouteRequest) (response *DeleteMessageRouteResponse, err error) {
    return c.DeleteMessageRouteWithContext(context.Background(), request)
}

// DeleteMessageRoute
// 删除消息路由
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteMessageRouteWithContext(ctx context.Context, request *DeleteMessageRouteRequest) (response *DeleteMessageRouteResponse, err error) {
    if request == nil {
        request = NewDeleteMessageRouteRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteMessageRoute require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteMessageRouteResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteNamespaceRequest() (request *DeleteNamespaceRequest) {
    request = &DeleteNamespaceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iecp", APIVersion, "DeleteNamespace")
    
    
    return
}

func NewDeleteNamespaceResponse() (response *DeleteNamespaceResponse) {
    response = &DeleteNamespaceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteNamespace
// 删除命名空间
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteNamespace(request *DeleteNamespaceRequest) (response *DeleteNamespaceResponse, err error) {
    return c.DeleteNamespaceWithContext(context.Background(), request)
}

// DeleteNamespace
// 删除命名空间
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteNamespaceWithContext(ctx context.Context, request *DeleteNamespaceRequest) (response *DeleteNamespaceResponse, err error) {
    if request == nil {
        request = NewDeleteNamespaceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteNamespace require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteNamespaceResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteNodeUnitRequest() (request *DeleteNodeUnitRequest) {
    request = &DeleteNodeUnitRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iecp", APIVersion, "DeleteNodeUnit")
    
    
    return
}

func NewDeleteNodeUnitResponse() (response *DeleteNodeUnitResponse) {
    response = &DeleteNodeUnitResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteNodeUnit
// 删除边缘单元NodeUnit
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteNodeUnit(request *DeleteNodeUnitRequest) (response *DeleteNodeUnitResponse, err error) {
    return c.DeleteNodeUnitWithContext(context.Background(), request)
}

// DeleteNodeUnit
// 删除边缘单元NodeUnit
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteNodeUnitWithContext(ctx context.Context, request *DeleteNodeUnitRequest) (response *DeleteNodeUnitResponse, err error) {
    if request == nil {
        request = NewDeleteNodeUnitRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteNodeUnit require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteNodeUnitResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteSecretRequest() (request *DeleteSecretRequest) {
    request = &DeleteSecretRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iecp", APIVersion, "DeleteSecret")
    
    
    return
}

func NewDeleteSecretResponse() (response *DeleteSecretResponse) {
    response = &DeleteSecretResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteSecret
// 删除Secret
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteSecret(request *DeleteSecretRequest) (response *DeleteSecretResponse, err error) {
    return c.DeleteSecretWithContext(context.Background(), request)
}

// DeleteSecret
// 删除Secret
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteSecretWithContext(ctx context.Context, request *DeleteSecretRequest) (response *DeleteSecretResponse, err error) {
    if request == nil {
        request = NewDeleteSecretRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteSecret require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteSecretResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeApplicationVisualizationRequest() (request *DescribeApplicationVisualizationRequest) {
    request = &DescribeApplicationVisualizationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iecp", APIVersion, "DescribeApplicationVisualization")
    
    
    return
}

func NewDescribeApplicationVisualizationResponse() (response *DescribeApplicationVisualizationResponse) {
    response = &DescribeApplicationVisualizationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeApplicationVisualization
// 获取应用模板可视化配置信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeApplicationVisualization(request *DescribeApplicationVisualizationRequest) (response *DescribeApplicationVisualizationResponse, err error) {
    return c.DescribeApplicationVisualizationWithContext(context.Background(), request)
}

// DescribeApplicationVisualization
// 获取应用模板可视化配置信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeApplicationVisualizationWithContext(ctx context.Context, request *DescribeApplicationVisualizationRequest) (response *DescribeApplicationVisualizationResponse, err error) {
    if request == nil {
        request = NewDescribeApplicationVisualizationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeApplicationVisualization require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeApplicationVisualizationResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeApplicationYamlRequest() (request *DescribeApplicationYamlRequest) {
    request = &DescribeApplicationYamlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iecp", APIVersion, "DescribeApplicationYaml")
    
    
    return
}

func NewDescribeApplicationYamlResponse() (response *DescribeApplicationYamlResponse) {
    response = &DescribeApplicationYamlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeApplicationYaml
// 查询应用模板Yaml
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeApplicationYaml(request *DescribeApplicationYamlRequest) (response *DescribeApplicationYamlResponse, err error) {
    return c.DescribeApplicationYamlWithContext(context.Background(), request)
}

// DescribeApplicationYaml
// 查询应用模板Yaml
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeApplicationYamlWithContext(ctx context.Context, request *DescribeApplicationYamlRequest) (response *DescribeApplicationYamlResponse, err error) {
    if request == nil {
        request = NewDescribeApplicationYamlRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeApplicationYaml require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeApplicationYamlResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeApplicationYamlErrorRequest() (request *DescribeApplicationYamlErrorRequest) {
    request = &DescribeApplicationYamlErrorRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iecp", APIVersion, "DescribeApplicationYamlError")
    
    
    return
}

func NewDescribeApplicationYamlErrorResponse() (response *DescribeApplicationYamlErrorResponse) {
    response = &DescribeApplicationYamlErrorResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeApplicationYamlError
// 检查应用模板的Yaml配置
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeApplicationYamlError(request *DescribeApplicationYamlErrorRequest) (response *DescribeApplicationYamlErrorResponse, err error) {
    return c.DescribeApplicationYamlErrorWithContext(context.Background(), request)
}

// DescribeApplicationYamlError
// 检查应用模板的Yaml配置
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeApplicationYamlErrorWithContext(ctx context.Context, request *DescribeApplicationYamlErrorRequest) (response *DescribeApplicationYamlErrorResponse, err error) {
    if request == nil {
        request = NewDescribeApplicationYamlErrorRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeApplicationYamlError require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeApplicationYamlErrorResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeApplicationsRequest() (request *DescribeApplicationsRequest) {
    request = &DescribeApplicationsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iecp", APIVersion, "DescribeApplications")
    
    
    return
}

func NewDescribeApplicationsResponse() (response *DescribeApplicationsResponse) {
    response = &DescribeApplicationsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeApplications
// 获取应用模板列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeApplications(request *DescribeApplicationsRequest) (response *DescribeApplicationsResponse, err error) {
    return c.DescribeApplicationsWithContext(context.Background(), request)
}

// DescribeApplications
// 获取应用模板列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeApplicationsWithContext(ctx context.Context, request *DescribeApplicationsRequest) (response *DescribeApplicationsResponse, err error) {
    if request == nil {
        request = NewDescribeApplicationsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeApplications require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeApplicationsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeConfigMapRequest() (request *DescribeConfigMapRequest) {
    request = &DescribeConfigMapRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iecp", APIVersion, "DescribeConfigMap")
    
    
    return
}

func NewDescribeConfigMapResponse() (response *DescribeConfigMapResponse) {
    response = &DescribeConfigMapResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeConfigMap
// 获取ConfigMap详情
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeConfigMap(request *DescribeConfigMapRequest) (response *DescribeConfigMapResponse, err error) {
    return c.DescribeConfigMapWithContext(context.Background(), request)
}

// DescribeConfigMap
// 获取ConfigMap详情
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeConfigMapWithContext(ctx context.Context, request *DescribeConfigMapRequest) (response *DescribeConfigMapResponse, err error) {
    if request == nil {
        request = NewDescribeConfigMapRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeConfigMap require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeConfigMapResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeConfigMapYamlErrorRequest() (request *DescribeConfigMapYamlErrorRequest) {
    request = &DescribeConfigMapYamlErrorRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iecp", APIVersion, "DescribeConfigMapYamlError")
    
    
    return
}

func NewDescribeConfigMapYamlErrorResponse() (response *DescribeConfigMapYamlErrorResponse) {
    response = &DescribeConfigMapYamlErrorResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeConfigMapYamlError
// 校验ConfigMap的Yaml语法
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeConfigMapYamlError(request *DescribeConfigMapYamlErrorRequest) (response *DescribeConfigMapYamlErrorResponse, err error) {
    return c.DescribeConfigMapYamlErrorWithContext(context.Background(), request)
}

// DescribeConfigMapYamlError
// 校验ConfigMap的Yaml语法
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeConfigMapYamlErrorWithContext(ctx context.Context, request *DescribeConfigMapYamlErrorRequest) (response *DescribeConfigMapYamlErrorResponse, err error) {
    if request == nil {
        request = NewDescribeConfigMapYamlErrorRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeConfigMapYamlError require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeConfigMapYamlErrorResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeConfigMapsRequest() (request *DescribeConfigMapsRequest) {
    request = &DescribeConfigMapsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iecp", APIVersion, "DescribeConfigMaps")
    
    
    return
}

func NewDescribeConfigMapsResponse() (response *DescribeConfigMapsResponse) {
    response = &DescribeConfigMapsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeConfigMaps
// 获取ConfigMap列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeConfigMaps(request *DescribeConfigMapsRequest) (response *DescribeConfigMapsResponse, err error) {
    return c.DescribeConfigMapsWithContext(context.Background(), request)
}

// DescribeConfigMaps
// 获取ConfigMap列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeConfigMapsWithContext(ctx context.Context, request *DescribeConfigMapsRequest) (response *DescribeConfigMapsResponse, err error) {
    if request == nil {
        request = NewDescribeConfigMapsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeConfigMaps require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeConfigMapsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDracoEdgeNodeInstallerRequest() (request *DescribeDracoEdgeNodeInstallerRequest) {
    request = &DescribeDracoEdgeNodeInstallerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iecp", APIVersion, "DescribeDracoEdgeNodeInstaller")
    
    
    return
}

func NewDescribeDracoEdgeNodeInstallerResponse() (response *DescribeDracoEdgeNodeInstallerResponse) {
    response = &DescribeDracoEdgeNodeInstallerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDracoEdgeNodeInstaller
// 自动获取Draco设备的安装包
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDracoEdgeNodeInstaller(request *DescribeDracoEdgeNodeInstallerRequest) (response *DescribeDracoEdgeNodeInstallerResponse, err error) {
    return c.DescribeDracoEdgeNodeInstallerWithContext(context.Background(), request)
}

// DescribeDracoEdgeNodeInstaller
// 自动获取Draco设备的安装包
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDracoEdgeNodeInstallerWithContext(ctx context.Context, request *DescribeDracoEdgeNodeInstallerRequest) (response *DescribeDracoEdgeNodeInstallerResponse, err error) {
    if request == nil {
        request = NewDescribeDracoEdgeNodeInstallerRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDracoEdgeNodeInstaller require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDracoEdgeNodeInstallerResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEdgeAgentNodeInstallerRequest() (request *DescribeEdgeAgentNodeInstallerRequest) {
    request = &DescribeEdgeAgentNodeInstallerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iecp", APIVersion, "DescribeEdgeAgentNodeInstaller")
    
    
    return
}

func NewDescribeEdgeAgentNodeInstallerResponse() (response *DescribeEdgeAgentNodeInstallerResponse) {
    response = &DescribeEdgeAgentNodeInstallerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeEdgeAgentNodeInstaller
// 获取节点安装信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeEdgeAgentNodeInstaller(request *DescribeEdgeAgentNodeInstallerRequest) (response *DescribeEdgeAgentNodeInstallerResponse, err error) {
    return c.DescribeEdgeAgentNodeInstallerWithContext(context.Background(), request)
}

// DescribeEdgeAgentNodeInstaller
// 获取节点安装信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeEdgeAgentNodeInstallerWithContext(ctx context.Context, request *DescribeEdgeAgentNodeInstallerRequest) (response *DescribeEdgeAgentNodeInstallerResponse, err error) {
    if request == nil {
        request = NewDescribeEdgeAgentNodeInstallerRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeEdgeAgentNodeInstaller require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeEdgeAgentNodeInstallerResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEdgeDefaultVpcRequest() (request *DescribeEdgeDefaultVpcRequest) {
    request = &DescribeEdgeDefaultVpcRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iecp", APIVersion, "DescribeEdgeDefaultVpc")
    
    
    return
}

func NewDescribeEdgeDefaultVpcResponse() (response *DescribeEdgeDefaultVpcResponse) {
    response = &DescribeEdgeDefaultVpcResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeEdgeDefaultVpc
// 获取边缘集群默认VPC信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeEdgeDefaultVpc(request *DescribeEdgeDefaultVpcRequest) (response *DescribeEdgeDefaultVpcResponse, err error) {
    return c.DescribeEdgeDefaultVpcWithContext(context.Background(), request)
}

// DescribeEdgeDefaultVpc
// 获取边缘集群默认VPC信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeEdgeDefaultVpcWithContext(ctx context.Context, request *DescribeEdgeDefaultVpcRequest) (response *DescribeEdgeDefaultVpcResponse, err error) {
    if request == nil {
        request = NewDescribeEdgeDefaultVpcRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeEdgeDefaultVpc require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeEdgeDefaultVpcResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEdgeNodeRequest() (request *DescribeEdgeNodeRequest) {
    request = &DescribeEdgeNodeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iecp", APIVersion, "DescribeEdgeNode")
    
    
    return
}

func NewDescribeEdgeNodeResponse() (response *DescribeEdgeNodeResponse) {
    response = &DescribeEdgeNodeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeEdgeNode
// 获取边缘节点信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeEdgeNode(request *DescribeEdgeNodeRequest) (response *DescribeEdgeNodeResponse, err error) {
    return c.DescribeEdgeNodeWithContext(context.Background(), request)
}

// DescribeEdgeNode
// 获取边缘节点信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeEdgeNodeWithContext(ctx context.Context, request *DescribeEdgeNodeRequest) (response *DescribeEdgeNodeResponse, err error) {
    if request == nil {
        request = NewDescribeEdgeNodeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeEdgeNode require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeEdgeNodeResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEdgeNodePodContainersRequest() (request *DescribeEdgeNodePodContainersRequest) {
    request = &DescribeEdgeNodePodContainersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iecp", APIVersion, "DescribeEdgeNodePodContainers")
    
    
    return
}

func NewDescribeEdgeNodePodContainersResponse() (response *DescribeEdgeNodePodContainersResponse) {
    response = &DescribeEdgeNodePodContainersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeEdgeNodePodContainers
// 查询节点Pod内的容器列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeEdgeNodePodContainers(request *DescribeEdgeNodePodContainersRequest) (response *DescribeEdgeNodePodContainersResponse, err error) {
    return c.DescribeEdgeNodePodContainersWithContext(context.Background(), request)
}

// DescribeEdgeNodePodContainers
// 查询节点Pod内的容器列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeEdgeNodePodContainersWithContext(ctx context.Context, request *DescribeEdgeNodePodContainersRequest) (response *DescribeEdgeNodePodContainersResponse, err error) {
    if request == nil {
        request = NewDescribeEdgeNodePodContainersRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeEdgeNodePodContainers require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeEdgeNodePodContainersResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEdgeNodePodsRequest() (request *DescribeEdgeNodePodsRequest) {
    request = &DescribeEdgeNodePodsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iecp", APIVersion, "DescribeEdgeNodePods")
    
    
    return
}

func NewDescribeEdgeNodePodsResponse() (response *DescribeEdgeNodePodsResponse) {
    response = &DescribeEdgeNodePodsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeEdgeNodePods
// 查询节点Pod列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeEdgeNodePods(request *DescribeEdgeNodePodsRequest) (response *DescribeEdgeNodePodsResponse, err error) {
    return c.DescribeEdgeNodePodsWithContext(context.Background(), request)
}

// DescribeEdgeNodePods
// 查询节点Pod列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeEdgeNodePodsWithContext(ctx context.Context, request *DescribeEdgeNodePodsRequest) (response *DescribeEdgeNodePodsResponse, err error) {
    if request == nil {
        request = NewDescribeEdgeNodePodsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeEdgeNodePods require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeEdgeNodePodsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEdgeNodeRemarkListRequest() (request *DescribeEdgeNodeRemarkListRequest) {
    request = &DescribeEdgeNodeRemarkListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iecp", APIVersion, "DescribeEdgeNodeRemarkList")
    
    
    return
}

func NewDescribeEdgeNodeRemarkListResponse() (response *DescribeEdgeNodeRemarkListResponse) {
    response = &DescribeEdgeNodeRemarkListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeEdgeNodeRemarkList
// 获取节点备注信息列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeEdgeNodeRemarkList(request *DescribeEdgeNodeRemarkListRequest) (response *DescribeEdgeNodeRemarkListResponse, err error) {
    return c.DescribeEdgeNodeRemarkListWithContext(context.Background(), request)
}

// DescribeEdgeNodeRemarkList
// 获取节点备注信息列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeEdgeNodeRemarkListWithContext(ctx context.Context, request *DescribeEdgeNodeRemarkListRequest) (response *DescribeEdgeNodeRemarkListResponse, err error) {
    if request == nil {
        request = NewDescribeEdgeNodeRemarkListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeEdgeNodeRemarkList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeEdgeNodeRemarkListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEdgeNodesRequest() (request *DescribeEdgeNodesRequest) {
    request = &DescribeEdgeNodesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iecp", APIVersion, "DescribeEdgeNodes")
    
    
    return
}

func NewDescribeEdgeNodesResponse() (response *DescribeEdgeNodesResponse) {
    response = &DescribeEdgeNodesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeEdgeNodes
// 查询边缘节点列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeEdgeNodes(request *DescribeEdgeNodesRequest) (response *DescribeEdgeNodesResponse, err error) {
    return c.DescribeEdgeNodesWithContext(context.Background(), request)
}

// DescribeEdgeNodes
// 查询边缘节点列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeEdgeNodesWithContext(ctx context.Context, request *DescribeEdgeNodesRequest) (response *DescribeEdgeNodesResponse, err error) {
    if request == nil {
        request = NewDescribeEdgeNodesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeEdgeNodes require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeEdgeNodesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEdgeOperationLogsRequest() (request *DescribeEdgeOperationLogsRequest) {
    request = &DescribeEdgeOperationLogsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iecp", APIVersion, "DescribeEdgeOperationLogs")
    
    
    return
}

func NewDescribeEdgeOperationLogsResponse() (response *DescribeEdgeOperationLogsResponse) {
    response = &DescribeEdgeOperationLogsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeEdgeOperationLogs
// 查询边缘操作日志
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeEdgeOperationLogs(request *DescribeEdgeOperationLogsRequest) (response *DescribeEdgeOperationLogsResponse, err error) {
    return c.DescribeEdgeOperationLogsWithContext(context.Background(), request)
}

// DescribeEdgeOperationLogs
// 查询边缘操作日志
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeEdgeOperationLogsWithContext(ctx context.Context, request *DescribeEdgeOperationLogsRequest) (response *DescribeEdgeOperationLogsResponse, err error) {
    if request == nil {
        request = NewDescribeEdgeOperationLogsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeEdgeOperationLogs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeEdgeOperationLogsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEdgePodRequest() (request *DescribeEdgePodRequest) {
    request = &DescribeEdgePodRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iecp", APIVersion, "DescribeEdgePod")
    
    
    return
}

func NewDescribeEdgePodResponse() (response *DescribeEdgePodResponse) {
    response = &DescribeEdgePodResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeEdgePod
// 查询边缘单元Pod
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeEdgePod(request *DescribeEdgePodRequest) (response *DescribeEdgePodResponse, err error) {
    return c.DescribeEdgePodWithContext(context.Background(), request)
}

// DescribeEdgePod
// 查询边缘单元Pod
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeEdgePodWithContext(ctx context.Context, request *DescribeEdgePodRequest) (response *DescribeEdgePodResponse, err error) {
    if request == nil {
        request = NewDescribeEdgePodRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeEdgePod require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeEdgePodResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEdgeSnNodesRequest() (request *DescribeEdgeSnNodesRequest) {
    request = &DescribeEdgeSnNodesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iecp", APIVersion, "DescribeEdgeSnNodes")
    
    
    return
}

func NewDescribeEdgeSnNodesResponse() (response *DescribeEdgeSnNodesResponse) {
    response = &DescribeEdgeSnNodesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeEdgeSnNodes
// 查询预注册节点列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeEdgeSnNodes(request *DescribeEdgeSnNodesRequest) (response *DescribeEdgeSnNodesResponse, err error) {
    return c.DescribeEdgeSnNodesWithContext(context.Background(), request)
}

// DescribeEdgeSnNodes
// 查询预注册节点列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeEdgeSnNodesWithContext(ctx context.Context, request *DescribeEdgeSnNodesRequest) (response *DescribeEdgeSnNodesResponse, err error) {
    if request == nil {
        request = NewDescribeEdgeSnNodesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeEdgeSnNodes require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeEdgeSnNodesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEdgeUnitApplicationEventsRequest() (request *DescribeEdgeUnitApplicationEventsRequest) {
    request = &DescribeEdgeUnitApplicationEventsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iecp", APIVersion, "DescribeEdgeUnitApplicationEvents")
    
    
    return
}

func NewDescribeEdgeUnitApplicationEventsResponse() (response *DescribeEdgeUnitApplicationEventsResponse) {
    response = &DescribeEdgeUnitApplicationEventsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeEdgeUnitApplicationEvents
// 获取应用事件列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeEdgeUnitApplicationEvents(request *DescribeEdgeUnitApplicationEventsRequest) (response *DescribeEdgeUnitApplicationEventsResponse, err error) {
    return c.DescribeEdgeUnitApplicationEventsWithContext(context.Background(), request)
}

// DescribeEdgeUnitApplicationEvents
// 获取应用事件列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeEdgeUnitApplicationEventsWithContext(ctx context.Context, request *DescribeEdgeUnitApplicationEventsRequest) (response *DescribeEdgeUnitApplicationEventsResponse, err error) {
    if request == nil {
        request = NewDescribeEdgeUnitApplicationEventsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeEdgeUnitApplicationEvents require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeEdgeUnitApplicationEventsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEdgeUnitApplicationLogsRequest() (request *DescribeEdgeUnitApplicationLogsRequest) {
    request = &DescribeEdgeUnitApplicationLogsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iecp", APIVersion, "DescribeEdgeUnitApplicationLogs")
    
    
    return
}

func NewDescribeEdgeUnitApplicationLogsResponse() (response *DescribeEdgeUnitApplicationLogsResponse) {
    response = &DescribeEdgeUnitApplicationLogsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeEdgeUnitApplicationLogs
// 获取应用日志
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeEdgeUnitApplicationLogs(request *DescribeEdgeUnitApplicationLogsRequest) (response *DescribeEdgeUnitApplicationLogsResponse, err error) {
    return c.DescribeEdgeUnitApplicationLogsWithContext(context.Background(), request)
}

// DescribeEdgeUnitApplicationLogs
// 获取应用日志
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeEdgeUnitApplicationLogsWithContext(ctx context.Context, request *DescribeEdgeUnitApplicationLogsRequest) (response *DescribeEdgeUnitApplicationLogsResponse, err error) {
    if request == nil {
        request = NewDescribeEdgeUnitApplicationLogsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeEdgeUnitApplicationLogs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeEdgeUnitApplicationLogsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEdgeUnitApplicationPodContainersRequest() (request *DescribeEdgeUnitApplicationPodContainersRequest) {
    request = &DescribeEdgeUnitApplicationPodContainersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iecp", APIVersion, "DescribeEdgeUnitApplicationPodContainers")
    
    
    return
}

func NewDescribeEdgeUnitApplicationPodContainersResponse() (response *DescribeEdgeUnitApplicationPodContainersResponse) {
    response = &DescribeEdgeUnitApplicationPodContainersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeEdgeUnitApplicationPodContainers
// 获取应用容器状态
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeEdgeUnitApplicationPodContainers(request *DescribeEdgeUnitApplicationPodContainersRequest) (response *DescribeEdgeUnitApplicationPodContainersResponse, err error) {
    return c.DescribeEdgeUnitApplicationPodContainersWithContext(context.Background(), request)
}

// DescribeEdgeUnitApplicationPodContainers
// 获取应用容器状态
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeEdgeUnitApplicationPodContainersWithContext(ctx context.Context, request *DescribeEdgeUnitApplicationPodContainersRequest) (response *DescribeEdgeUnitApplicationPodContainersResponse, err error) {
    if request == nil {
        request = NewDescribeEdgeUnitApplicationPodContainersRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeEdgeUnitApplicationPodContainers require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeEdgeUnitApplicationPodContainersResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEdgeUnitApplicationPodsRequest() (request *DescribeEdgeUnitApplicationPodsRequest) {
    request = &DescribeEdgeUnitApplicationPodsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iecp", APIVersion, "DescribeEdgeUnitApplicationPods")
    
    
    return
}

func NewDescribeEdgeUnitApplicationPodsResponse() (response *DescribeEdgeUnitApplicationPodsResponse) {
    response = &DescribeEdgeUnitApplicationPodsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeEdgeUnitApplicationPods
// 获取应用下Pod状态
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeEdgeUnitApplicationPods(request *DescribeEdgeUnitApplicationPodsRequest) (response *DescribeEdgeUnitApplicationPodsResponse, err error) {
    return c.DescribeEdgeUnitApplicationPodsWithContext(context.Background(), request)
}

// DescribeEdgeUnitApplicationPods
// 获取应用下Pod状态
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeEdgeUnitApplicationPodsWithContext(ctx context.Context, request *DescribeEdgeUnitApplicationPodsRequest) (response *DescribeEdgeUnitApplicationPodsResponse, err error) {
    if request == nil {
        request = NewDescribeEdgeUnitApplicationPodsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeEdgeUnitApplicationPods require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeEdgeUnitApplicationPodsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEdgeUnitApplicationVisualizationRequest() (request *DescribeEdgeUnitApplicationVisualizationRequest) {
    request = &DescribeEdgeUnitApplicationVisualizationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iecp", APIVersion, "DescribeEdgeUnitApplicationVisualization")
    
    
    return
}

func NewDescribeEdgeUnitApplicationVisualizationResponse() (response *DescribeEdgeUnitApplicationVisualizationResponse) {
    response = &DescribeEdgeUnitApplicationVisualizationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeEdgeUnitApplicationVisualization
// 获取单元可视化配置信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeEdgeUnitApplicationVisualization(request *DescribeEdgeUnitApplicationVisualizationRequest) (response *DescribeEdgeUnitApplicationVisualizationResponse, err error) {
    return c.DescribeEdgeUnitApplicationVisualizationWithContext(context.Background(), request)
}

// DescribeEdgeUnitApplicationVisualization
// 获取单元可视化配置信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeEdgeUnitApplicationVisualizationWithContext(ctx context.Context, request *DescribeEdgeUnitApplicationVisualizationRequest) (response *DescribeEdgeUnitApplicationVisualizationResponse, err error) {
    if request == nil {
        request = NewDescribeEdgeUnitApplicationVisualizationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeEdgeUnitApplicationVisualization require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeEdgeUnitApplicationVisualizationResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEdgeUnitApplicationYamlRequest() (request *DescribeEdgeUnitApplicationYamlRequest) {
    request = &DescribeEdgeUnitApplicationYamlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iecp", APIVersion, "DescribeEdgeUnitApplicationYaml")
    
    
    return
}

func NewDescribeEdgeUnitApplicationYamlResponse() (response *DescribeEdgeUnitApplicationYamlResponse) {
    response = &DescribeEdgeUnitApplicationYamlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeEdgeUnitApplicationYaml
// 获取应用的Yaml配置
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeEdgeUnitApplicationYaml(request *DescribeEdgeUnitApplicationYamlRequest) (response *DescribeEdgeUnitApplicationYamlResponse, err error) {
    return c.DescribeEdgeUnitApplicationYamlWithContext(context.Background(), request)
}

// DescribeEdgeUnitApplicationYaml
// 获取应用的Yaml配置
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeEdgeUnitApplicationYamlWithContext(ctx context.Context, request *DescribeEdgeUnitApplicationYamlRequest) (response *DescribeEdgeUnitApplicationYamlResponse, err error) {
    if request == nil {
        request = NewDescribeEdgeUnitApplicationYamlRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeEdgeUnitApplicationYaml require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeEdgeUnitApplicationYamlResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEdgeUnitApplicationYamlErrorRequest() (request *DescribeEdgeUnitApplicationYamlErrorRequest) {
    request = &DescribeEdgeUnitApplicationYamlErrorRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iecp", APIVersion, "DescribeEdgeUnitApplicationYamlError")
    
    
    return
}

func NewDescribeEdgeUnitApplicationYamlErrorResponse() (response *DescribeEdgeUnitApplicationYamlErrorResponse) {
    response = &DescribeEdgeUnitApplicationYamlErrorResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeEdgeUnitApplicationYamlError
// 检查单元应用的Yaml配置
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeEdgeUnitApplicationYamlError(request *DescribeEdgeUnitApplicationYamlErrorRequest) (response *DescribeEdgeUnitApplicationYamlErrorResponse, err error) {
    return c.DescribeEdgeUnitApplicationYamlErrorWithContext(context.Background(), request)
}

// DescribeEdgeUnitApplicationYamlError
// 检查单元应用的Yaml配置
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeEdgeUnitApplicationYamlErrorWithContext(ctx context.Context, request *DescribeEdgeUnitApplicationYamlErrorRequest) (response *DescribeEdgeUnitApplicationYamlErrorResponse, err error) {
    if request == nil {
        request = NewDescribeEdgeUnitApplicationYamlErrorRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeEdgeUnitApplicationYamlError require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeEdgeUnitApplicationYamlErrorResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEdgeUnitApplicationsRequest() (request *DescribeEdgeUnitApplicationsRequest) {
    request = &DescribeEdgeUnitApplicationsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iecp", APIVersion, "DescribeEdgeUnitApplications")
    
    
    return
}

func NewDescribeEdgeUnitApplicationsResponse() (response *DescribeEdgeUnitApplicationsResponse) {
    response = &DescribeEdgeUnitApplicationsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeEdgeUnitApplications
// 获取单元下应用列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeEdgeUnitApplications(request *DescribeEdgeUnitApplicationsRequest) (response *DescribeEdgeUnitApplicationsResponse, err error) {
    return c.DescribeEdgeUnitApplicationsWithContext(context.Background(), request)
}

// DescribeEdgeUnitApplications
// 获取单元下应用列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeEdgeUnitApplicationsWithContext(ctx context.Context, request *DescribeEdgeUnitApplicationsRequest) (response *DescribeEdgeUnitApplicationsResponse, err error) {
    if request == nil {
        request = NewDescribeEdgeUnitApplicationsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeEdgeUnitApplications require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeEdgeUnitApplicationsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEdgeUnitCloudRequest() (request *DescribeEdgeUnitCloudRequest) {
    request = &DescribeEdgeUnitCloudRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iecp", APIVersion, "DescribeEdgeUnitCloud")
    
    
    return
}

func NewDescribeEdgeUnitCloudResponse() (response *DescribeEdgeUnitCloudResponse) {
    response = &DescribeEdgeUnitCloudResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeEdgeUnitCloud
// 查询边缘集群详情
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeEdgeUnitCloud(request *DescribeEdgeUnitCloudRequest) (response *DescribeEdgeUnitCloudResponse, err error) {
    return c.DescribeEdgeUnitCloudWithContext(context.Background(), request)
}

// DescribeEdgeUnitCloud
// 查询边缘集群详情
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeEdgeUnitCloudWithContext(ctx context.Context, request *DescribeEdgeUnitCloudRequest) (response *DescribeEdgeUnitCloudResponse, err error) {
    if request == nil {
        request = NewDescribeEdgeUnitCloudRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeEdgeUnitCloud require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeEdgeUnitCloudResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEdgeUnitDeployGridRequest() (request *DescribeEdgeUnitDeployGridRequest) {
    request = &DescribeEdgeUnitDeployGridRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iecp", APIVersion, "DescribeEdgeUnitDeployGrid")
    
    
    return
}

func NewDescribeEdgeUnitDeployGridResponse() (response *DescribeEdgeUnitDeployGridResponse) {
    response = &DescribeEdgeUnitDeployGridResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeEdgeUnitDeployGrid
// 查询边缘单元Grid列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeEdgeUnitDeployGrid(request *DescribeEdgeUnitDeployGridRequest) (response *DescribeEdgeUnitDeployGridResponse, err error) {
    return c.DescribeEdgeUnitDeployGridWithContext(context.Background(), request)
}

// DescribeEdgeUnitDeployGrid
// 查询边缘单元Grid列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeEdgeUnitDeployGridWithContext(ctx context.Context, request *DescribeEdgeUnitDeployGridRequest) (response *DescribeEdgeUnitDeployGridResponse, err error) {
    if request == nil {
        request = NewDescribeEdgeUnitDeployGridRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeEdgeUnitDeployGrid require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeEdgeUnitDeployGridResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEdgeUnitDeployGridItemRequest() (request *DescribeEdgeUnitDeployGridItemRequest) {
    request = &DescribeEdgeUnitDeployGridItemRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iecp", APIVersion, "DescribeEdgeUnitDeployGridItem")
    
    
    return
}

func NewDescribeEdgeUnitDeployGridItemResponse() (response *DescribeEdgeUnitDeployGridItemResponse) {
    response = &DescribeEdgeUnitDeployGridItemResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeEdgeUnitDeployGridItem
// 查询边缘单元指定Grid下的部署应用列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeEdgeUnitDeployGridItem(request *DescribeEdgeUnitDeployGridItemRequest) (response *DescribeEdgeUnitDeployGridItemResponse, err error) {
    return c.DescribeEdgeUnitDeployGridItemWithContext(context.Background(), request)
}

// DescribeEdgeUnitDeployGridItem
// 查询边缘单元指定Grid下的部署应用列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeEdgeUnitDeployGridItemWithContext(ctx context.Context, request *DescribeEdgeUnitDeployGridItemRequest) (response *DescribeEdgeUnitDeployGridItemResponse, err error) {
    if request == nil {
        request = NewDescribeEdgeUnitDeployGridItemRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeEdgeUnitDeployGridItem require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeEdgeUnitDeployGridItemResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEdgeUnitDeployGridItemYamlRequest() (request *DescribeEdgeUnitDeployGridItemYamlRequest) {
    request = &DescribeEdgeUnitDeployGridItemYamlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iecp", APIVersion, "DescribeEdgeUnitDeployGridItemYaml")
    
    
    return
}

func NewDescribeEdgeUnitDeployGridItemYamlResponse() (response *DescribeEdgeUnitDeployGridItemYamlResponse) {
    response = &DescribeEdgeUnitDeployGridItemYamlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeEdgeUnitDeployGridItemYaml
// 查询指定Grid下应用的Yaml
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeEdgeUnitDeployGridItemYaml(request *DescribeEdgeUnitDeployGridItemYamlRequest) (response *DescribeEdgeUnitDeployGridItemYamlResponse, err error) {
    return c.DescribeEdgeUnitDeployGridItemYamlWithContext(context.Background(), request)
}

// DescribeEdgeUnitDeployGridItemYaml
// 查询指定Grid下应用的Yaml
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeEdgeUnitDeployGridItemYamlWithContext(ctx context.Context, request *DescribeEdgeUnitDeployGridItemYamlRequest) (response *DescribeEdgeUnitDeployGridItemYamlResponse, err error) {
    if request == nil {
        request = NewDescribeEdgeUnitDeployGridItemYamlRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeEdgeUnitDeployGridItemYaml require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeEdgeUnitDeployGridItemYamlResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEdgeUnitExtraRequest() (request *DescribeEdgeUnitExtraRequest) {
    request = &DescribeEdgeUnitExtraRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iecp", APIVersion, "DescribeEdgeUnitExtra")
    
    
    return
}

func NewDescribeEdgeUnitExtraResponse() (response *DescribeEdgeUnitExtraResponse) {
    response = &DescribeEdgeUnitExtraResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeEdgeUnitExtra
// 查询边缘单元额外信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeEdgeUnitExtra(request *DescribeEdgeUnitExtraRequest) (response *DescribeEdgeUnitExtraResponse, err error) {
    return c.DescribeEdgeUnitExtraWithContext(context.Background(), request)
}

// DescribeEdgeUnitExtra
// 查询边缘单元额外信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeEdgeUnitExtraWithContext(ctx context.Context, request *DescribeEdgeUnitExtraRequest) (response *DescribeEdgeUnitExtraResponse, err error) {
    if request == nil {
        request = NewDescribeEdgeUnitExtraRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeEdgeUnitExtra require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeEdgeUnitExtraResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEdgeUnitGridEventsRequest() (request *DescribeEdgeUnitGridEventsRequest) {
    request = &DescribeEdgeUnitGridEventsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iecp", APIVersion, "DescribeEdgeUnitGridEvents")
    
    
    return
}

func NewDescribeEdgeUnitGridEventsResponse() (response *DescribeEdgeUnitGridEventsResponse) {
    response = &DescribeEdgeUnitGridEventsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeEdgeUnitGridEvents
// 查询边缘单元Grid事件列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeEdgeUnitGridEvents(request *DescribeEdgeUnitGridEventsRequest) (response *DescribeEdgeUnitGridEventsResponse, err error) {
    return c.DescribeEdgeUnitGridEventsWithContext(context.Background(), request)
}

// DescribeEdgeUnitGridEvents
// 查询边缘单元Grid事件列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeEdgeUnitGridEventsWithContext(ctx context.Context, request *DescribeEdgeUnitGridEventsRequest) (response *DescribeEdgeUnitGridEventsResponse, err error) {
    if request == nil {
        request = NewDescribeEdgeUnitGridEventsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeEdgeUnitGridEvents require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeEdgeUnitGridEventsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEdgeUnitGridPodsRequest() (request *DescribeEdgeUnitGridPodsRequest) {
    request = &DescribeEdgeUnitGridPodsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iecp", APIVersion, "DescribeEdgeUnitGridPods")
    
    
    return
}

func NewDescribeEdgeUnitGridPodsResponse() (response *DescribeEdgeUnitGridPodsResponse) {
    response = &DescribeEdgeUnitGridPodsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeEdgeUnitGridPods
// 查询边缘单元Grid的Pod列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeEdgeUnitGridPods(request *DescribeEdgeUnitGridPodsRequest) (response *DescribeEdgeUnitGridPodsResponse, err error) {
    return c.DescribeEdgeUnitGridPodsWithContext(context.Background(), request)
}

// DescribeEdgeUnitGridPods
// 查询边缘单元Grid的Pod列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeEdgeUnitGridPodsWithContext(ctx context.Context, request *DescribeEdgeUnitGridPodsRequest) (response *DescribeEdgeUnitGridPodsResponse, err error) {
    if request == nil {
        request = NewDescribeEdgeUnitGridPodsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeEdgeUnitGridPods require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeEdgeUnitGridPodsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEdgeUnitMonitorStatusRequest() (request *DescribeEdgeUnitMonitorStatusRequest) {
    request = &DescribeEdgeUnitMonitorStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iecp", APIVersion, "DescribeEdgeUnitMonitorStatus")
    
    
    return
}

func NewDescribeEdgeUnitMonitorStatusResponse() (response *DescribeEdgeUnitMonitorStatusResponse) {
    response = &DescribeEdgeUnitMonitorStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeEdgeUnitMonitorStatus
// 查询边缘集群监控状态
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeEdgeUnitMonitorStatus(request *DescribeEdgeUnitMonitorStatusRequest) (response *DescribeEdgeUnitMonitorStatusResponse, err error) {
    return c.DescribeEdgeUnitMonitorStatusWithContext(context.Background(), request)
}

// DescribeEdgeUnitMonitorStatus
// 查询边缘集群监控状态
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeEdgeUnitMonitorStatusWithContext(ctx context.Context, request *DescribeEdgeUnitMonitorStatusRequest) (response *DescribeEdgeUnitMonitorStatusResponse, err error) {
    if request == nil {
        request = NewDescribeEdgeUnitMonitorStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeEdgeUnitMonitorStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeEdgeUnitMonitorStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEdgeUnitNodeGroupRequest() (request *DescribeEdgeUnitNodeGroupRequest) {
    request = &DescribeEdgeUnitNodeGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iecp", APIVersion, "DescribeEdgeUnitNodeGroup")
    
    
    return
}

func NewDescribeEdgeUnitNodeGroupResponse() (response *DescribeEdgeUnitNodeGroupResponse) {
    response = &DescribeEdgeUnitNodeGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeEdgeUnitNodeGroup
// 查询边缘集群NodeGroup
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeEdgeUnitNodeGroup(request *DescribeEdgeUnitNodeGroupRequest) (response *DescribeEdgeUnitNodeGroupResponse, err error) {
    return c.DescribeEdgeUnitNodeGroupWithContext(context.Background(), request)
}

// DescribeEdgeUnitNodeGroup
// 查询边缘集群NodeGroup
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeEdgeUnitNodeGroupWithContext(ctx context.Context, request *DescribeEdgeUnitNodeGroupRequest) (response *DescribeEdgeUnitNodeGroupResponse, err error) {
    if request == nil {
        request = NewDescribeEdgeUnitNodeGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeEdgeUnitNodeGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeEdgeUnitNodeGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEdgeUnitNodeUnitTemplatesRequest() (request *DescribeEdgeUnitNodeUnitTemplatesRequest) {
    request = &DescribeEdgeUnitNodeUnitTemplatesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iecp", APIVersion, "DescribeEdgeUnitNodeUnitTemplates")
    
    
    return
}

func NewDescribeEdgeUnitNodeUnitTemplatesResponse() (response *DescribeEdgeUnitNodeUnitTemplatesResponse) {
    response = &DescribeEdgeUnitNodeUnitTemplatesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeEdgeUnitNodeUnitTemplates
// 查询边缘单元EdgeUnit模板列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeEdgeUnitNodeUnitTemplates(request *DescribeEdgeUnitNodeUnitTemplatesRequest) (response *DescribeEdgeUnitNodeUnitTemplatesResponse, err error) {
    return c.DescribeEdgeUnitNodeUnitTemplatesWithContext(context.Background(), request)
}

// DescribeEdgeUnitNodeUnitTemplates
// 查询边缘单元EdgeUnit模板列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeEdgeUnitNodeUnitTemplatesWithContext(ctx context.Context, request *DescribeEdgeUnitNodeUnitTemplatesRequest) (response *DescribeEdgeUnitNodeUnitTemplatesResponse, err error) {
    if request == nil {
        request = NewDescribeEdgeUnitNodeUnitTemplatesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeEdgeUnitNodeUnitTemplates require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeEdgeUnitNodeUnitTemplatesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEdgeUnitsCloudRequest() (request *DescribeEdgeUnitsCloudRequest) {
    request = &DescribeEdgeUnitsCloudRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iecp", APIVersion, "DescribeEdgeUnitsCloud")
    
    
    return
}

func NewDescribeEdgeUnitsCloudResponse() (response *DescribeEdgeUnitsCloudResponse) {
    response = &DescribeEdgeUnitsCloudResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeEdgeUnitsCloud
// 查询边缘单元列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeEdgeUnitsCloud(request *DescribeEdgeUnitsCloudRequest) (response *DescribeEdgeUnitsCloudResponse, err error) {
    return c.DescribeEdgeUnitsCloudWithContext(context.Background(), request)
}

// DescribeEdgeUnitsCloud
// 查询边缘单元列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeEdgeUnitsCloudWithContext(ctx context.Context, request *DescribeEdgeUnitsCloudRequest) (response *DescribeEdgeUnitsCloudResponse, err error) {
    if request == nil {
        request = NewDescribeEdgeUnitsCloudRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeEdgeUnitsCloud require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeEdgeUnitsCloudResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeIotDeviceRequest() (request *DescribeIotDeviceRequest) {
    request = &DescribeIotDeviceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iecp", APIVersion, "DescribeIotDevice")
    
    
    return
}

func NewDescribeIotDeviceResponse() (response *DescribeIotDeviceResponse) {
    response = &DescribeIotDeviceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeIotDevice
// 获取设备信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeIotDevice(request *DescribeIotDeviceRequest) (response *DescribeIotDeviceResponse, err error) {
    return c.DescribeIotDeviceWithContext(context.Background(), request)
}

// DescribeIotDevice
// 获取设备信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeIotDeviceWithContext(ctx context.Context, request *DescribeIotDeviceRequest) (response *DescribeIotDeviceResponse, err error) {
    if request == nil {
        request = NewDescribeIotDeviceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeIotDevice require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeIotDeviceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeIotDevicesRequest() (request *DescribeIotDevicesRequest) {
    request = &DescribeIotDevicesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iecp", APIVersion, "DescribeIotDevices")
    
    
    return
}

func NewDescribeIotDevicesResponse() (response *DescribeIotDevicesResponse) {
    response = &DescribeIotDevicesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeIotDevices
// 获取设备列表信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeIotDevices(request *DescribeIotDevicesRequest) (response *DescribeIotDevicesResponse, err error) {
    return c.DescribeIotDevicesWithContext(context.Background(), request)
}

// DescribeIotDevices
// 获取设备列表信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeIotDevicesWithContext(ctx context.Context, request *DescribeIotDevicesRequest) (response *DescribeIotDevicesResponse, err error) {
    if request == nil {
        request = NewDescribeIotDevicesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeIotDevices require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeIotDevicesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMessageRouteListRequest() (request *DescribeMessageRouteListRequest) {
    request = &DescribeMessageRouteListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iecp", APIVersion, "DescribeMessageRouteList")
    
    
    return
}

func NewDescribeMessageRouteListResponse() (response *DescribeMessageRouteListResponse) {
    response = &DescribeMessageRouteListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeMessageRouteList
// 获取消息路由列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeMessageRouteList(request *DescribeMessageRouteListRequest) (response *DescribeMessageRouteListResponse, err error) {
    return c.DescribeMessageRouteListWithContext(context.Background(), request)
}

// DescribeMessageRouteList
// 获取消息路由列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeMessageRouteListWithContext(ctx context.Context, request *DescribeMessageRouteListRequest) (response *DescribeMessageRouteListResponse, err error) {
    if request == nil {
        request = NewDescribeMessageRouteListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMessageRouteList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMessageRouteListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMonitorMetricsRequest() (request *DescribeMonitorMetricsRequest) {
    request = &DescribeMonitorMetricsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iecp", APIVersion, "DescribeMonitorMetrics")
    
    
    return
}

func NewDescribeMonitorMetricsResponse() (response *DescribeMonitorMetricsResponse) {
    response = &DescribeMonitorMetricsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeMonitorMetrics
// 查询边缘单元监控数据
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeMonitorMetrics(request *DescribeMonitorMetricsRequest) (response *DescribeMonitorMetricsResponse, err error) {
    return c.DescribeMonitorMetricsWithContext(context.Background(), request)
}

// DescribeMonitorMetrics
// 查询边缘单元监控数据
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeMonitorMetricsWithContext(ctx context.Context, request *DescribeMonitorMetricsRequest) (response *DescribeMonitorMetricsResponse, err error) {
    if request == nil {
        request = NewDescribeMonitorMetricsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMonitorMetrics require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMonitorMetricsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeNamespaceRequest() (request *DescribeNamespaceRequest) {
    request = &DescribeNamespaceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iecp", APIVersion, "DescribeNamespace")
    
    
    return
}

func NewDescribeNamespaceResponse() (response *DescribeNamespaceResponse) {
    response = &DescribeNamespaceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeNamespace
// 获取命名空间
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeNamespace(request *DescribeNamespaceRequest) (response *DescribeNamespaceResponse, err error) {
    return c.DescribeNamespaceWithContext(context.Background(), request)
}

// DescribeNamespace
// 获取命名空间
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeNamespaceWithContext(ctx context.Context, request *DescribeNamespaceRequest) (response *DescribeNamespaceResponse, err error) {
    if request == nil {
        request = NewDescribeNamespaceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeNamespace require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeNamespaceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeNamespaceResourcesRequest() (request *DescribeNamespaceResourcesRequest) {
    request = &DescribeNamespaceResourcesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iecp", APIVersion, "DescribeNamespaceResources")
    
    
    return
}

func NewDescribeNamespaceResourcesResponse() (response *DescribeNamespaceResourcesResponse) {
    response = &DescribeNamespaceResourcesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeNamespaceResources
// 获取命名空间下的资源信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeNamespaceResources(request *DescribeNamespaceResourcesRequest) (response *DescribeNamespaceResourcesResponse, err error) {
    return c.DescribeNamespaceResourcesWithContext(context.Background(), request)
}

// DescribeNamespaceResources
// 获取命名空间下的资源信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeNamespaceResourcesWithContext(ctx context.Context, request *DescribeNamespaceResourcesRequest) (response *DescribeNamespaceResourcesResponse, err error) {
    if request == nil {
        request = NewDescribeNamespaceResourcesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeNamespaceResources require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeNamespaceResourcesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeNamespacesRequest() (request *DescribeNamespacesRequest) {
    request = &DescribeNamespacesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iecp", APIVersion, "DescribeNamespaces")
    
    
    return
}

func NewDescribeNamespacesResponse() (response *DescribeNamespacesResponse) {
    response = &DescribeNamespacesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeNamespaces
// 获取命名空间列表信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeNamespaces(request *DescribeNamespacesRequest) (response *DescribeNamespacesResponse, err error) {
    return c.DescribeNamespacesWithContext(context.Background(), request)
}

// DescribeNamespaces
// 获取命名空间列表信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeNamespacesWithContext(ctx context.Context, request *DescribeNamespacesRequest) (response *DescribeNamespacesResponse, err error) {
    if request == nil {
        request = NewDescribeNamespacesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeNamespaces require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeNamespacesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeNodeUnitRequest() (request *DescribeNodeUnitRequest) {
    request = &DescribeNodeUnitRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iecp", APIVersion, "DescribeNodeUnit")
    
    
    return
}

func NewDescribeNodeUnitResponse() (response *DescribeNodeUnitResponse) {
    response = &DescribeNodeUnitResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeNodeUnit
// 查询边缘单元NodeUnit列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeNodeUnit(request *DescribeNodeUnitRequest) (response *DescribeNodeUnitResponse, err error) {
    return c.DescribeNodeUnitWithContext(context.Background(), request)
}

// DescribeNodeUnit
// 查询边缘单元NodeUnit列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeNodeUnitWithContext(ctx context.Context, request *DescribeNodeUnitRequest) (response *DescribeNodeUnitResponse, err error) {
    if request == nil {
        request = NewDescribeNodeUnitRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeNodeUnit require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeNodeUnitResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeNodeUnitTemplateOnNodeGroupRequest() (request *DescribeNodeUnitTemplateOnNodeGroupRequest) {
    request = &DescribeNodeUnitTemplateOnNodeGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iecp", APIVersion, "DescribeNodeUnitTemplateOnNodeGroup")
    
    
    return
}

func NewDescribeNodeUnitTemplateOnNodeGroupResponse() (response *DescribeNodeUnitTemplateOnNodeGroupResponse) {
    response = &DescribeNodeUnitTemplateOnNodeGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeNodeUnitTemplateOnNodeGroup
// 查询指定NodeGroup下NodeUnit模板列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeNodeUnitTemplateOnNodeGroup(request *DescribeNodeUnitTemplateOnNodeGroupRequest) (response *DescribeNodeUnitTemplateOnNodeGroupResponse, err error) {
    return c.DescribeNodeUnitTemplateOnNodeGroupWithContext(context.Background(), request)
}

// DescribeNodeUnitTemplateOnNodeGroup
// 查询指定NodeGroup下NodeUnit模板列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeNodeUnitTemplateOnNodeGroupWithContext(ctx context.Context, request *DescribeNodeUnitTemplateOnNodeGroupRequest) (response *DescribeNodeUnitTemplateOnNodeGroupResponse, err error) {
    if request == nil {
        request = NewDescribeNodeUnitTemplateOnNodeGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeNodeUnitTemplateOnNodeGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeNodeUnitTemplateOnNodeGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSecretRequest() (request *DescribeSecretRequest) {
    request = &DescribeSecretRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iecp", APIVersion, "DescribeSecret")
    
    
    return
}

func NewDescribeSecretResponse() (response *DescribeSecretResponse) {
    response = &DescribeSecretResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeSecret
// 获取Secret详情
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeSecret(request *DescribeSecretRequest) (response *DescribeSecretResponse, err error) {
    return c.DescribeSecretWithContext(context.Background(), request)
}

// DescribeSecret
// 获取Secret详情
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeSecretWithContext(ctx context.Context, request *DescribeSecretRequest) (response *DescribeSecretResponse, err error) {
    if request == nil {
        request = NewDescribeSecretRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSecret require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSecretResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSecretYamlErrorRequest() (request *DescribeSecretYamlErrorRequest) {
    request = &DescribeSecretYamlErrorRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iecp", APIVersion, "DescribeSecretYamlError")
    
    
    return
}

func NewDescribeSecretYamlErrorResponse() (response *DescribeSecretYamlErrorResponse) {
    response = &DescribeSecretYamlErrorResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeSecretYamlError
// 校验Secret的Yaml语法
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeSecretYamlError(request *DescribeSecretYamlErrorRequest) (response *DescribeSecretYamlErrorResponse, err error) {
    return c.DescribeSecretYamlErrorWithContext(context.Background(), request)
}

// DescribeSecretYamlError
// 校验Secret的Yaml语法
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeSecretYamlErrorWithContext(ctx context.Context, request *DescribeSecretYamlErrorRequest) (response *DescribeSecretYamlErrorResponse, err error) {
    if request == nil {
        request = NewDescribeSecretYamlErrorRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSecretYamlError require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSecretYamlErrorResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSecretsRequest() (request *DescribeSecretsRequest) {
    request = &DescribeSecretsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iecp", APIVersion, "DescribeSecrets")
    
    
    return
}

func NewDescribeSecretsResponse() (response *DescribeSecretsResponse) {
    response = &DescribeSecretsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeSecrets
// 获取Secrets列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeSecrets(request *DescribeSecretsRequest) (response *DescribeSecretsResponse, err error) {
    return c.DescribeSecretsWithContext(context.Background(), request)
}

// DescribeSecrets
// 获取Secrets列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeSecretsWithContext(ctx context.Context, request *DescribeSecretsRequest) (response *DescribeSecretsResponse, err error) {
    if request == nil {
        request = NewDescribeSecretsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSecrets require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSecretsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeYeheResourceLimitRequest() (request *DescribeYeheResourceLimitRequest) {
    request = &DescribeYeheResourceLimitRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iecp", APIVersion, "DescribeYeheResourceLimit")
    
    
    return
}

func NewDescribeYeheResourceLimitResponse() (response *DescribeYeheResourceLimitResponse) {
    response = &DescribeYeheResourceLimitResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeYeheResourceLimit
// 查询用户的资源限制
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeYeheResourceLimit(request *DescribeYeheResourceLimitRequest) (response *DescribeYeheResourceLimitResponse, err error) {
    return c.DescribeYeheResourceLimitWithContext(context.Background(), request)
}

// DescribeYeheResourceLimit
// 查询用户的资源限制
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeYeheResourceLimitWithContext(ctx context.Context, request *DescribeYeheResourceLimitRequest) (response *DescribeYeheResourceLimitResponse, err error) {
    if request == nil {
        request = NewDescribeYeheResourceLimitRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeYeheResourceLimit require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeYeheResourceLimitResponse()
    err = c.Send(request, response)
    return
}

func NewGetMarketComponentRequest() (request *GetMarketComponentRequest) {
    request = &GetMarketComponentRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iecp", APIVersion, "GetMarketComponent")
    
    
    return
}

func NewGetMarketComponentResponse() (response *GetMarketComponentResponse) {
    response = &GetMarketComponentResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetMarketComponent
// 获取组件市场的组件信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) GetMarketComponent(request *GetMarketComponentRequest) (response *GetMarketComponentResponse, err error) {
    return c.GetMarketComponentWithContext(context.Background(), request)
}

// GetMarketComponent
// 获取组件市场的组件信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) GetMarketComponentWithContext(ctx context.Context, request *GetMarketComponentRequest) (response *GetMarketComponentResponse, err error) {
    if request == nil {
        request = NewGetMarketComponentRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetMarketComponent require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetMarketComponentResponse()
    err = c.Send(request, response)
    return
}

func NewGetMarketComponentListRequest() (request *GetMarketComponentListRequest) {
    request = &GetMarketComponentListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iecp", APIVersion, "GetMarketComponentList")
    
    
    return
}

func NewGetMarketComponentListResponse() (response *GetMarketComponentListResponse) {
    response = &GetMarketComponentListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetMarketComponentList
// 获取组件市场组件列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) GetMarketComponentList(request *GetMarketComponentListRequest) (response *GetMarketComponentListResponse, err error) {
    return c.GetMarketComponentListWithContext(context.Background(), request)
}

// GetMarketComponentList
// 获取组件市场组件列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) GetMarketComponentListWithContext(ctx context.Context, request *GetMarketComponentListRequest) (response *GetMarketComponentListResponse, err error) {
    if request == nil {
        request = NewGetMarketComponentListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetMarketComponentList require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetMarketComponentListResponse()
    err = c.Send(request, response)
    return
}

func NewModifyApplicationBasicInfoRequest() (request *ModifyApplicationBasicInfoRequest) {
    request = &ModifyApplicationBasicInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iecp", APIVersion, "ModifyApplicationBasicInfo")
    
    
    return
}

func NewModifyApplicationBasicInfoResponse() (response *ModifyApplicationBasicInfoResponse) {
    response = &ModifyApplicationBasicInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyApplicationBasicInfo
// 修改应用模板基本信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyApplicationBasicInfo(request *ModifyApplicationBasicInfoRequest) (response *ModifyApplicationBasicInfoResponse, err error) {
    return c.ModifyApplicationBasicInfoWithContext(context.Background(), request)
}

// ModifyApplicationBasicInfo
// 修改应用模板基本信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyApplicationBasicInfoWithContext(ctx context.Context, request *ModifyApplicationBasicInfoRequest) (response *ModifyApplicationBasicInfoResponse, err error) {
    if request == nil {
        request = NewModifyApplicationBasicInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyApplicationBasicInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyApplicationBasicInfoResponse()
    err = c.Send(request, response)
    return
}

func NewModifyApplicationVisualizationRequest() (request *ModifyApplicationVisualizationRequest) {
    request = &ModifyApplicationVisualizationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iecp", APIVersion, "ModifyApplicationVisualization")
    
    
    return
}

func NewModifyApplicationVisualizationResponse() (response *ModifyApplicationVisualizationResponse) {
    response = &ModifyApplicationVisualizationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyApplicationVisualization
// 修改应用模板配置
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyApplicationVisualization(request *ModifyApplicationVisualizationRequest) (response *ModifyApplicationVisualizationResponse, err error) {
    return c.ModifyApplicationVisualizationWithContext(context.Background(), request)
}

// ModifyApplicationVisualization
// 修改应用模板配置
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyApplicationVisualizationWithContext(ctx context.Context, request *ModifyApplicationVisualizationRequest) (response *ModifyApplicationVisualizationResponse, err error) {
    if request == nil {
        request = NewModifyApplicationVisualizationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyApplicationVisualization require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyApplicationVisualizationResponse()
    err = c.Send(request, response)
    return
}

func NewModifyConfigMapRequest() (request *ModifyConfigMapRequest) {
    request = &ModifyConfigMapRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iecp", APIVersion, "ModifyConfigMap")
    
    
    return
}

func NewModifyConfigMapResponse() (response *ModifyConfigMapResponse) {
    response = &ModifyConfigMapResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyConfigMap
// 修改ConfigMap
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyConfigMap(request *ModifyConfigMapRequest) (response *ModifyConfigMapResponse, err error) {
    return c.ModifyConfigMapWithContext(context.Background(), request)
}

// ModifyConfigMap
// 修改ConfigMap
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyConfigMapWithContext(ctx context.Context, request *ModifyConfigMapRequest) (response *ModifyConfigMapResponse, err error) {
    if request == nil {
        request = NewModifyConfigMapRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyConfigMap require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyConfigMapResponse()
    err = c.Send(request, response)
    return
}

func NewModifyEdgeDracoNodeRequest() (request *ModifyEdgeDracoNodeRequest) {
    request = &ModifyEdgeDracoNodeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iecp", APIVersion, "ModifyEdgeDracoNode")
    
    
    return
}

func NewModifyEdgeDracoNodeResponse() (response *ModifyEdgeDracoNodeResponse) {
    response = &ModifyEdgeDracoNodeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyEdgeDracoNode
// 编辑draco设备信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyEdgeDracoNode(request *ModifyEdgeDracoNodeRequest) (response *ModifyEdgeDracoNodeResponse, err error) {
    return c.ModifyEdgeDracoNodeWithContext(context.Background(), request)
}

// ModifyEdgeDracoNode
// 编辑draco设备信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyEdgeDracoNodeWithContext(ctx context.Context, request *ModifyEdgeDracoNodeRequest) (response *ModifyEdgeDracoNodeResponse, err error) {
    if request == nil {
        request = NewModifyEdgeDracoNodeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyEdgeDracoNode require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyEdgeDracoNodeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyEdgeNodeLabelsRequest() (request *ModifyEdgeNodeLabelsRequest) {
    request = &ModifyEdgeNodeLabelsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iecp", APIVersion, "ModifyEdgeNodeLabels")
    
    
    return
}

func NewModifyEdgeNodeLabelsResponse() (response *ModifyEdgeNodeLabelsResponse) {
    response = &ModifyEdgeNodeLabelsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyEdgeNodeLabels
// 编辑边缘节点标签
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyEdgeNodeLabels(request *ModifyEdgeNodeLabelsRequest) (response *ModifyEdgeNodeLabelsResponse, err error) {
    return c.ModifyEdgeNodeLabelsWithContext(context.Background(), request)
}

// ModifyEdgeNodeLabels
// 编辑边缘节点标签
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyEdgeNodeLabelsWithContext(ctx context.Context, request *ModifyEdgeNodeLabelsRequest) (response *ModifyEdgeNodeLabelsResponse, err error) {
    if request == nil {
        request = NewModifyEdgeNodeLabelsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyEdgeNodeLabels require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyEdgeNodeLabelsResponse()
    err = c.Send(request, response)
    return
}

func NewModifyEdgeUnitRequest() (request *ModifyEdgeUnitRequest) {
    request = &ModifyEdgeUnitRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iecp", APIVersion, "ModifyEdgeUnit")
    
    
    return
}

func NewModifyEdgeUnitResponse() (response *ModifyEdgeUnitResponse) {
    response = &ModifyEdgeUnitResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyEdgeUnit
// 修改边缘集群
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyEdgeUnit(request *ModifyEdgeUnitRequest) (response *ModifyEdgeUnitResponse, err error) {
    return c.ModifyEdgeUnitWithContext(context.Background(), request)
}

// ModifyEdgeUnit
// 修改边缘集群
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyEdgeUnitWithContext(ctx context.Context, request *ModifyEdgeUnitRequest) (response *ModifyEdgeUnitResponse, err error) {
    if request == nil {
        request = NewModifyEdgeUnitRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyEdgeUnit require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyEdgeUnitResponse()
    err = c.Send(request, response)
    return
}

func NewModifyEdgeUnitApplicationBasicInfoRequest() (request *ModifyEdgeUnitApplicationBasicInfoRequest) {
    request = &ModifyEdgeUnitApplicationBasicInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iecp", APIVersion, "ModifyEdgeUnitApplicationBasicInfo")
    
    
    return
}

func NewModifyEdgeUnitApplicationBasicInfoResponse() (response *ModifyEdgeUnitApplicationBasicInfoResponse) {
    response = &ModifyEdgeUnitApplicationBasicInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyEdgeUnitApplicationBasicInfo
// 修改单元应用基本信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyEdgeUnitApplicationBasicInfo(request *ModifyEdgeUnitApplicationBasicInfoRequest) (response *ModifyEdgeUnitApplicationBasicInfoResponse, err error) {
    return c.ModifyEdgeUnitApplicationBasicInfoWithContext(context.Background(), request)
}

// ModifyEdgeUnitApplicationBasicInfo
// 修改单元应用基本信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyEdgeUnitApplicationBasicInfoWithContext(ctx context.Context, request *ModifyEdgeUnitApplicationBasicInfoRequest) (response *ModifyEdgeUnitApplicationBasicInfoResponse, err error) {
    if request == nil {
        request = NewModifyEdgeUnitApplicationBasicInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyEdgeUnitApplicationBasicInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyEdgeUnitApplicationBasicInfoResponse()
    err = c.Send(request, response)
    return
}

func NewModifyEdgeUnitApplicationVisualizationRequest() (request *ModifyEdgeUnitApplicationVisualizationRequest) {
    request = &ModifyEdgeUnitApplicationVisualizationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iecp", APIVersion, "ModifyEdgeUnitApplicationVisualization")
    
    
    return
}

func NewModifyEdgeUnitApplicationVisualizationResponse() (response *ModifyEdgeUnitApplicationVisualizationResponse) {
    response = &ModifyEdgeUnitApplicationVisualizationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyEdgeUnitApplicationVisualization
// 可视化修改应用配置
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyEdgeUnitApplicationVisualization(request *ModifyEdgeUnitApplicationVisualizationRequest) (response *ModifyEdgeUnitApplicationVisualizationResponse, err error) {
    return c.ModifyEdgeUnitApplicationVisualizationWithContext(context.Background(), request)
}

// ModifyEdgeUnitApplicationVisualization
// 可视化修改应用配置
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyEdgeUnitApplicationVisualizationWithContext(ctx context.Context, request *ModifyEdgeUnitApplicationVisualizationRequest) (response *ModifyEdgeUnitApplicationVisualizationResponse, err error) {
    if request == nil {
        request = NewModifyEdgeUnitApplicationVisualizationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyEdgeUnitApplicationVisualization require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyEdgeUnitApplicationVisualizationResponse()
    err = c.Send(request, response)
    return
}

func NewModifyEdgeUnitApplicationYamlRequest() (request *ModifyEdgeUnitApplicationYamlRequest) {
    request = &ModifyEdgeUnitApplicationYamlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iecp", APIVersion, "ModifyEdgeUnitApplicationYaml")
    
    
    return
}

func NewModifyEdgeUnitApplicationYamlResponse() (response *ModifyEdgeUnitApplicationYamlResponse) {
    response = &ModifyEdgeUnitApplicationYamlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyEdgeUnitApplicationYaml
// Yaml方式修改应用配置
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyEdgeUnitApplicationYaml(request *ModifyEdgeUnitApplicationYamlRequest) (response *ModifyEdgeUnitApplicationYamlResponse, err error) {
    return c.ModifyEdgeUnitApplicationYamlWithContext(context.Background(), request)
}

// ModifyEdgeUnitApplicationYaml
// Yaml方式修改应用配置
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyEdgeUnitApplicationYamlWithContext(ctx context.Context, request *ModifyEdgeUnitApplicationYamlRequest) (response *ModifyEdgeUnitApplicationYamlResponse, err error) {
    if request == nil {
        request = NewModifyEdgeUnitApplicationYamlRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyEdgeUnitApplicationYaml require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyEdgeUnitApplicationYamlResponse()
    err = c.Send(request, response)
    return
}

func NewModifyEdgeUnitCloudApiRequest() (request *ModifyEdgeUnitCloudApiRequest) {
    request = &ModifyEdgeUnitCloudApiRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iecp", APIVersion, "ModifyEdgeUnitCloudApi")
    
    
    return
}

func NewModifyEdgeUnitCloudApiResponse() (response *ModifyEdgeUnitCloudApiResponse) {
    response = &ModifyEdgeUnitCloudApiResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyEdgeUnitCloudApi
// 更新边缘单元信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ModifyEdgeUnitCloudApi(request *ModifyEdgeUnitCloudApiRequest) (response *ModifyEdgeUnitCloudApiResponse, err error) {
    return c.ModifyEdgeUnitCloudApiWithContext(context.Background(), request)
}

// ModifyEdgeUnitCloudApi
// 更新边缘单元信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ModifyEdgeUnitCloudApiWithContext(ctx context.Context, request *ModifyEdgeUnitCloudApiRequest) (response *ModifyEdgeUnitCloudApiResponse, err error) {
    if request == nil {
        request = NewModifyEdgeUnitCloudApiRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyEdgeUnitCloudApi require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyEdgeUnitCloudApiResponse()
    err = c.Send(request, response)
    return
}

func NewModifyEdgeUnitDeployGridItemRequest() (request *ModifyEdgeUnitDeployGridItemRequest) {
    request = &ModifyEdgeUnitDeployGridItemRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iecp", APIVersion, "ModifyEdgeUnitDeployGridItem")
    
    
    return
}

func NewModifyEdgeUnitDeployGridItemResponse() (response *ModifyEdgeUnitDeployGridItemResponse) {
    response = &ModifyEdgeUnitDeployGridItemResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyEdgeUnitDeployGridItem
// 修改边缘单元Grid部署应用副本数
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyEdgeUnitDeployGridItem(request *ModifyEdgeUnitDeployGridItemRequest) (response *ModifyEdgeUnitDeployGridItemResponse, err error) {
    return c.ModifyEdgeUnitDeployGridItemWithContext(context.Background(), request)
}

// ModifyEdgeUnitDeployGridItem
// 修改边缘单元Grid部署应用副本数
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyEdgeUnitDeployGridItemWithContext(ctx context.Context, request *ModifyEdgeUnitDeployGridItemRequest) (response *ModifyEdgeUnitDeployGridItemResponse, err error) {
    if request == nil {
        request = NewModifyEdgeUnitDeployGridItemRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyEdgeUnitDeployGridItem require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyEdgeUnitDeployGridItemResponse()
    err = c.Send(request, response)
    return
}

func NewModifyIotDeviceRequest() (request *ModifyIotDeviceRequest) {
    request = &ModifyIotDeviceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iecp", APIVersion, "ModifyIotDevice")
    
    
    return
}

func NewModifyIotDeviceResponse() (response *ModifyIotDeviceResponse) {
    response = &ModifyIotDeviceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyIotDevice
// 修改设备信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyIotDevice(request *ModifyIotDeviceRequest) (response *ModifyIotDeviceResponse, err error) {
    return c.ModifyIotDeviceWithContext(context.Background(), request)
}

// ModifyIotDevice
// 修改设备信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyIotDeviceWithContext(ctx context.Context, request *ModifyIotDeviceRequest) (response *ModifyIotDeviceResponse, err error) {
    if request == nil {
        request = NewModifyIotDeviceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyIotDevice require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyIotDeviceResponse()
    err = c.Send(request, response)
    return
}

func NewModifyNodeUnitTemplateRequest() (request *ModifyNodeUnitTemplateRequest) {
    request = &ModifyNodeUnitTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iecp", APIVersion, "ModifyNodeUnitTemplate")
    
    
    return
}

func NewModifyNodeUnitTemplateResponse() (response *ModifyNodeUnitTemplateResponse) {
    response = &ModifyNodeUnitTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyNodeUnitTemplate
// 修改边缘单元NodeUnit模板
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyNodeUnitTemplate(request *ModifyNodeUnitTemplateRequest) (response *ModifyNodeUnitTemplateResponse, err error) {
    return c.ModifyNodeUnitTemplateWithContext(context.Background(), request)
}

// ModifyNodeUnitTemplate
// 修改边缘单元NodeUnit模板
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyNodeUnitTemplateWithContext(ctx context.Context, request *ModifyNodeUnitTemplateRequest) (response *ModifyNodeUnitTemplateResponse, err error) {
    if request == nil {
        request = NewModifyNodeUnitTemplateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyNodeUnitTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyNodeUnitTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewModifySecretRequest() (request *ModifySecretRequest) {
    request = &ModifySecretRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iecp", APIVersion, "ModifySecret")
    
    
    return
}

func NewModifySecretResponse() (response *ModifySecretResponse) {
    response = &ModifySecretResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifySecret
// 修改Secret
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifySecret(request *ModifySecretRequest) (response *ModifySecretResponse, err error) {
    return c.ModifySecretWithContext(context.Background(), request)
}

// ModifySecret
// 修改Secret
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifySecretWithContext(ctx context.Context, request *ModifySecretRequest) (response *ModifySecretResponse, err error) {
    if request == nil {
        request = NewModifySecretRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifySecret require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifySecretResponse()
    err = c.Send(request, response)
    return
}

func NewRedeployEdgeUnitApplicationRequest() (request *RedeployEdgeUnitApplicationRequest) {
    request = &RedeployEdgeUnitApplicationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iecp", APIVersion, "RedeployEdgeUnitApplication")
    
    
    return
}

func NewRedeployEdgeUnitApplicationResponse() (response *RedeployEdgeUnitApplicationResponse) {
    response = &RedeployEdgeUnitApplicationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// RedeployEdgeUnitApplication
// 单元应用重部署
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) RedeployEdgeUnitApplication(request *RedeployEdgeUnitApplicationRequest) (response *RedeployEdgeUnitApplicationResponse, err error) {
    return c.RedeployEdgeUnitApplicationWithContext(context.Background(), request)
}

// RedeployEdgeUnitApplication
// 单元应用重部署
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) RedeployEdgeUnitApplicationWithContext(ctx context.Context, request *RedeployEdgeUnitApplicationRequest) (response *RedeployEdgeUnitApplicationResponse, err error) {
    if request == nil {
        request = NewRedeployEdgeUnitApplicationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RedeployEdgeUnitApplication require credential")
    }

    request.SetContext(ctx)
    
    response = NewRedeployEdgeUnitApplicationResponse()
    err = c.Send(request, response)
    return
}

func NewSetRouteOnOffRequest() (request *SetRouteOnOffRequest) {
    request = &SetRouteOnOffRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iecp", APIVersion, "SetRouteOnOff")
    
    
    return
}

func NewSetRouteOnOffResponse() (response *SetRouteOnOffResponse) {
    response = &SetRouteOnOffResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// SetRouteOnOff
// 开关消息路由
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) SetRouteOnOff(request *SetRouteOnOffRequest) (response *SetRouteOnOffResponse, err error) {
    return c.SetRouteOnOffWithContext(context.Background(), request)
}

// SetRouteOnOff
// 开关消息路由
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) SetRouteOnOffWithContext(ctx context.Context, request *SetRouteOnOffRequest) (response *SetRouteOnOffResponse, err error) {
    if request == nil {
        request = NewSetRouteOnOffRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SetRouteOnOff require credential")
    }

    request.SetContext(ctx)
    
    response = NewSetRouteOnOffResponse()
    err = c.Send(request, response)
    return
}
