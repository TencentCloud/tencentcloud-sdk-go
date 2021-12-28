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
    if request == nil {
        request = NewApplyMarketComponentRequest()
    }
    
    response = NewApplyMarketComponentResponse()
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
    if request == nil {
        request = NewCreateApplicationVisualizationRequest()
    }
    
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
    if request == nil {
        request = NewCreateConfigMapRequest()
    }
    
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
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateEdgeNode(request *CreateEdgeNodeRequest) (response *CreateEdgeNodeResponse, err error) {
    if request == nil {
        request = NewCreateEdgeNodeRequest()
    }
    
    response = NewCreateEdgeNodeResponse()
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
    if request == nil {
        request = NewCreateEdgeNodeGroupRequest()
    }
    
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
// 创建边缘单元NodeUnit模版
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
    if request == nil {
        request = NewCreateEdgeNodeUnitTemplateRequest()
    }
    
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
    if request == nil {
        request = NewCreateEdgeUnitApplicationVisualizationRequest()
    }
    
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
    if request == nil {
        request = NewCreateEdgeUnitApplicationYamlRequest()
    }
    
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
    if request == nil {
        request = NewCreateEdgeUnitCloudRequest()
    }
    
    response = NewCreateEdgeUnitCloudResponse()
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
    if request == nil {
        request = NewCreateNamespaceRequest()
    }
    
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
    if request == nil {
        request = NewCreateSecretRequest()
    }
    
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
    if request == nil {
        request = NewCreateUpdateNodeUnitRequest()
    }
    
    response = NewCreateUpdateNodeUnitResponse()
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
    if request == nil {
        request = NewDeleteApplicationsRequest()
    }
    
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
    if request == nil {
        request = NewDeleteConfigMapRequest()
    }
    
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
    if request == nil {
        request = NewDeleteEdgeNodeGroupRequest()
    }
    
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
// 删除边缘单元NodeUnit模版
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
    if request == nil {
        request = NewDeleteEdgeNodeUnitTemplatesRequest()
    }
    
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
    if request == nil {
        request = NewDeleteEdgeNodesRequest()
    }
    
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
    if request == nil {
        request = NewDeleteEdgeUnitApplicationsRequest()
    }
    
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
    if request == nil {
        request = NewDeleteEdgeUnitCloudRequest()
    }
    
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
    if request == nil {
        request = NewDeleteEdgeUnitDeployGridItemRequest()
    }
    
    response = NewDeleteEdgeUnitDeployGridItemResponse()
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
    if request == nil {
        request = NewDeleteEdgeUnitPodRequest()
    }
    
    response = NewDeleteEdgeUnitPodResponse()
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
    if request == nil {
        request = NewDeleteNamespaceRequest()
    }
    
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
    if request == nil {
        request = NewDeleteNodeUnitRequest()
    }
    
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
    if request == nil {
        request = NewDeleteSecretRequest()
    }
    
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
    if request == nil {
        request = NewDescribeApplicationVisualizationRequest()
    }
    
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
    if request == nil {
        request = NewDescribeApplicationYamlRequest()
    }
    
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
    if request == nil {
        request = NewDescribeApplicationYamlErrorRequest()
    }
    
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
    if request == nil {
        request = NewDescribeApplicationsRequest()
    }
    
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
    if request == nil {
        request = NewDescribeConfigMapRequest()
    }
    
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
    if request == nil {
        request = NewDescribeConfigMapYamlErrorRequest()
    }
    
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
    if request == nil {
        request = NewDescribeConfigMapsRequest()
    }
    
    response = NewDescribeConfigMapsResponse()
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
    if request == nil {
        request = NewDescribeEdgeAgentNodeInstallerRequest()
    }
    
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
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeEdgeDefaultVpc(request *DescribeEdgeDefaultVpcRequest) (response *DescribeEdgeDefaultVpcResponse, err error) {
    if request == nil {
        request = NewDescribeEdgeDefaultVpcRequest()
    }
    
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
    if request == nil {
        request = NewDescribeEdgeNodeRequest()
    }
    
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
    if request == nil {
        request = NewDescribeEdgeNodePodContainersRequest()
    }
    
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
    if request == nil {
        request = NewDescribeEdgeNodePodsRequest()
    }
    
    response = NewDescribeEdgeNodePodsResponse()
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
    if request == nil {
        request = NewDescribeEdgeNodesRequest()
    }
    
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
    if request == nil {
        request = NewDescribeEdgeOperationLogsRequest()
    }
    
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
    if request == nil {
        request = NewDescribeEdgePodRequest()
    }
    
    response = NewDescribeEdgePodResponse()
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
    if request == nil {
        request = NewDescribeEdgeUnitApplicationEventsRequest()
    }
    
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
    if request == nil {
        request = NewDescribeEdgeUnitApplicationLogsRequest()
    }
    
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
    if request == nil {
        request = NewDescribeEdgeUnitApplicationPodContainersRequest()
    }
    
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
    if request == nil {
        request = NewDescribeEdgeUnitApplicationPodsRequest()
    }
    
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
    if request == nil {
        request = NewDescribeEdgeUnitApplicationVisualizationRequest()
    }
    
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
    if request == nil {
        request = NewDescribeEdgeUnitApplicationYamlRequest()
    }
    
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
    if request == nil {
        request = NewDescribeEdgeUnitApplicationYamlErrorRequest()
    }
    
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
    if request == nil {
        request = NewDescribeEdgeUnitApplicationsRequest()
    }
    
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
    if request == nil {
        request = NewDescribeEdgeUnitCloudRequest()
    }
    
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
    if request == nil {
        request = NewDescribeEdgeUnitDeployGridRequest()
    }
    
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
    if request == nil {
        request = NewDescribeEdgeUnitDeployGridItemRequest()
    }
    
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
    if request == nil {
        request = NewDescribeEdgeUnitDeployGridItemYamlRequest()
    }
    
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
    if request == nil {
        request = NewDescribeEdgeUnitExtraRequest()
    }
    
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
    if request == nil {
        request = NewDescribeEdgeUnitGridEventsRequest()
    }
    
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
    if request == nil {
        request = NewDescribeEdgeUnitGridPodsRequest()
    }
    
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
    if request == nil {
        request = NewDescribeEdgeUnitMonitorStatusRequest()
    }
    
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
    if request == nil {
        request = NewDescribeEdgeUnitNodeGroupRequest()
    }
    
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
// 查询边缘单元EdgeUnit模版列表
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
    if request == nil {
        request = NewDescribeEdgeUnitNodeUnitTemplatesRequest()
    }
    
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
    if request == nil {
        request = NewDescribeEdgeUnitsCloudRequest()
    }
    
    response = NewDescribeEdgeUnitsCloudResponse()
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
    if request == nil {
        request = NewDescribeMonitorMetricsRequest()
    }
    
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
    if request == nil {
        request = NewDescribeNamespaceRequest()
    }
    
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
    if request == nil {
        request = NewDescribeNamespaceResourcesRequest()
    }
    
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
    if request == nil {
        request = NewDescribeNamespacesRequest()
    }
    
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
    if request == nil {
        request = NewDescribeNodeUnitRequest()
    }
    
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
// 查询指定NodeGroup下NodeUnit模版列表
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
    if request == nil {
        request = NewDescribeNodeUnitTemplateOnNodeGroupRequest()
    }
    
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
    if request == nil {
        request = NewDescribeSecretRequest()
    }
    
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
    if request == nil {
        request = NewDescribeSecretYamlErrorRequest()
    }
    
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
    if request == nil {
        request = NewDescribeSecretsRequest()
    }
    
    response = NewDescribeSecretsResponse()
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
    if request == nil {
        request = NewGetMarketComponentRequest()
    }
    
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
    if request == nil {
        request = NewGetMarketComponentListRequest()
    }
    
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
    if request == nil {
        request = NewModifyApplicationBasicInfoRequest()
    }
    
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
    if request == nil {
        request = NewModifyApplicationVisualizationRequest()
    }
    
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
    if request == nil {
        request = NewModifyConfigMapRequest()
    }
    
    response = NewModifyConfigMapResponse()
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
    if request == nil {
        request = NewModifyEdgeNodeLabelsRequest()
    }
    
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
    if request == nil {
        request = NewModifyEdgeUnitRequest()
    }
    
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
    if request == nil {
        request = NewModifyEdgeUnitApplicationBasicInfoRequest()
    }
    
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
    if request == nil {
        request = NewModifyEdgeUnitApplicationVisualizationRequest()
    }
    
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
    if request == nil {
        request = NewModifyEdgeUnitApplicationYamlRequest()
    }
    
    response = NewModifyEdgeUnitApplicationYamlResponse()
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
    if request == nil {
        request = NewModifyEdgeUnitDeployGridItemRequest()
    }
    
    response = NewModifyEdgeUnitDeployGridItemResponse()
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
// 修改边缘单元NodeUnit模版
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
    if request == nil {
        request = NewModifyNodeUnitTemplateRequest()
    }
    
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
    if request == nil {
        request = NewModifySecretRequest()
    }
    
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
    if request == nil {
        request = NewRedeployEdgeUnitApplicationRequest()
    }
    
    response = NewRedeployEdgeUnitApplicationResponse()
    err = c.Send(request, response)
    return
}
