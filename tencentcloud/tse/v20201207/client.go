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
//  INVALIDPARAMETERVALUE_QUERYERROR = "InvalidParameterValue.QueryError"
func (c *Client) DescribeZookeeperServerInterfaces(request *DescribeZookeeperServerInterfacesRequest) (response *DescribeZookeeperServerInterfacesResponse, err error) {
    return c.DescribeZookeeperServerInterfacesWithContext(context.Background(), request)
}

// DescribeZookeeperServerInterfaces
// 查询zookeeper服务接口列表
//
// 可能返回的错误码:
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_QUERYERROR = "InvalidParameterValue.QueryError"
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
