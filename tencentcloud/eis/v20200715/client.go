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

package v20200715

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2020-07-15"

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


func NewDescribeEisConnectorConfigRequest() (request *DescribeEisConnectorConfigRequest) {
    request = &DescribeEisConnectorConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("eis", APIVersion, "DescribeEisConnectorConfig")
    
    
    return
}

func NewDescribeEisConnectorConfigResponse() (response *DescribeEisConnectorConfigResponse) {
    response = &DescribeEisConnectorConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeEisConnectorConfig
// 获取连接器配置参数
//
// 可能返回的错误码:
//  FAILEDOPERATION_AUTHENTICATIONFAILED = "FailedOperation.AuthenticationFailed"
//  FAILEDOPERATION_METACOMPILERERROR = "FailedOperation.MetaCompilerError"
//  INTERNALERROR_DATABASEERROR = "InternalError.DatabaseError"
//  INTERNALERROR_RPCPILOTSERVERERROR = "InternalError.RpcPilotServerError"
//  INVALIDPARAMETERVALUE_CONNECTORNOTEXIST = "InvalidParameterValue.ConnectorNotExist"
func (c *Client) DescribeEisConnectorConfig(request *DescribeEisConnectorConfigRequest) (response *DescribeEisConnectorConfigResponse, err error) {
    return c.DescribeEisConnectorConfigWithContext(context.Background(), request)
}

// DescribeEisConnectorConfig
// 获取连接器配置参数
//
// 可能返回的错误码:
//  FAILEDOPERATION_AUTHENTICATIONFAILED = "FailedOperation.AuthenticationFailed"
//  FAILEDOPERATION_METACOMPILERERROR = "FailedOperation.MetaCompilerError"
//  INTERNALERROR_DATABASEERROR = "InternalError.DatabaseError"
//  INTERNALERROR_RPCPILOTSERVERERROR = "InternalError.RpcPilotServerError"
//  INVALIDPARAMETERVALUE_CONNECTORNOTEXIST = "InvalidParameterValue.ConnectorNotExist"
func (c *Client) DescribeEisConnectorConfigWithContext(ctx context.Context, request *DescribeEisConnectorConfigRequest) (response *DescribeEisConnectorConfigResponse, err error) {
    if request == nil {
        request = NewDescribeEisConnectorConfigRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeEisConnectorConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeEisConnectorConfigResponse()
    err = c.Send(request, response)
    return
}

func NewListEisConnectorOperationsRequest() (request *ListEisConnectorOperationsRequest) {
    request = &ListEisConnectorOperationsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("eis", APIVersion, "ListEisConnectorOperations")
    
    
    return
}

func NewListEisConnectorOperationsResponse() (response *ListEisConnectorOperationsResponse) {
    response = &ListEisConnectorOperationsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ListEisConnectorOperations
// 获取连接器操作列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_AUTHENTICATIONFAILED = "FailedOperation.AuthenticationFailed"
//  FAILEDOPERATION_METACOMPILERERROR = "FailedOperation.MetaCompilerError"
//  INTERNALERROR_DATABASEERROR = "InternalError.DatabaseError"
//  INTERNALERROR_RPCPILOTSERVERERROR = "InternalError.RpcPilotServerError"
//  INVALIDPARAMETERVALUE_CONNECTORNOTEXIST = "InvalidParameterValue.ConnectorNotExist"
func (c *Client) ListEisConnectorOperations(request *ListEisConnectorOperationsRequest) (response *ListEisConnectorOperationsResponse, err error) {
    return c.ListEisConnectorOperationsWithContext(context.Background(), request)
}

// ListEisConnectorOperations
// 获取连接器操作列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_AUTHENTICATIONFAILED = "FailedOperation.AuthenticationFailed"
//  FAILEDOPERATION_METACOMPILERERROR = "FailedOperation.MetaCompilerError"
//  INTERNALERROR_DATABASEERROR = "InternalError.DatabaseError"
//  INTERNALERROR_RPCPILOTSERVERERROR = "InternalError.RpcPilotServerError"
//  INVALIDPARAMETERVALUE_CONNECTORNOTEXIST = "InvalidParameterValue.ConnectorNotExist"
func (c *Client) ListEisConnectorOperationsWithContext(ctx context.Context, request *ListEisConnectorOperationsRequest) (response *ListEisConnectorOperationsResponse, err error) {
    if request == nil {
        request = NewListEisConnectorOperationsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListEisConnectorOperations require credential")
    }

    request.SetContext(ctx)
    
    response = NewListEisConnectorOperationsResponse()
    err = c.Send(request, response)
    return
}

func NewListEisConnectorsRequest() (request *ListEisConnectorsRequest) {
    request = &ListEisConnectorsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("eis", APIVersion, "ListEisConnectors")
    
    
    return
}

func NewListEisConnectorsResponse() (response *ListEisConnectorsResponse) {
    response = &ListEisConnectorsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ListEisConnectors
// 连接器列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_AUTHENTICATIONFAILED = "FailedOperation.AuthenticationFailed"
//  INTERNALERROR_DATABASEERROR = "InternalError.DatabaseError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ListEisConnectors(request *ListEisConnectorsRequest) (response *ListEisConnectorsResponse, err error) {
    return c.ListEisConnectorsWithContext(context.Background(), request)
}

// ListEisConnectors
// 连接器列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_AUTHENTICATIONFAILED = "FailedOperation.AuthenticationFailed"
//  INTERNALERROR_DATABASEERROR = "InternalError.DatabaseError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ListEisConnectorsWithContext(ctx context.Context, request *ListEisConnectorsRequest) (response *ListEisConnectorsResponse, err error) {
    if request == nil {
        request = NewListEisConnectorsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListEisConnectors require credential")
    }

    request.SetContext(ctx)
    
    response = NewListEisConnectorsResponse()
    err = c.Send(request, response)
    return
}
