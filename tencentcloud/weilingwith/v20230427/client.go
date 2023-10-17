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

package v20230427

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2023-04-27"

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


func NewDescribeApplicationListRequest() (request *DescribeApplicationListRequest) {
    request = &DescribeApplicationListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("weilingwith", APIVersion, "DescribeApplicationList")
    
    
    return
}

func NewDescribeApplicationListResponse() (response *DescribeApplicationListResponse) {
    response = &DescribeApplicationListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeApplicationList
// 查询指定空间关联的应用列表
//
// 可能返回的错误码:
//  AUTHFAILURE_APIAUTHORIZATIONNOTFOUND = "AuthFailure.ApiAuthorizationNotFound"
//  AUTHFAILURE_TOKENNOTFOUND = "AuthFailure.TokenNotFound"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_APIGATEWAYINTERNALERROR = "InternalError.ApiGatewayInternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDAPPLICATIONID = "InvalidParameterValue.InvalidApplicationId"
func (c *Client) DescribeApplicationList(request *DescribeApplicationListRequest) (response *DescribeApplicationListResponse, err error) {
    return c.DescribeApplicationListWithContext(context.Background(), request)
}

// DescribeApplicationList
// 查询指定空间关联的应用列表
//
// 可能返回的错误码:
//  AUTHFAILURE_APIAUTHORIZATIONNOTFOUND = "AuthFailure.ApiAuthorizationNotFound"
//  AUTHFAILURE_TOKENNOTFOUND = "AuthFailure.TokenNotFound"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_APIGATEWAYINTERNALERROR = "InternalError.ApiGatewayInternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDAPPLICATIONID = "InvalidParameterValue.InvalidApplicationId"
func (c *Client) DescribeApplicationListWithContext(ctx context.Context, request *DescribeApplicationListRequest) (response *DescribeApplicationListResponse, err error) {
    if request == nil {
        request = NewDescribeApplicationListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeApplicationList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeApplicationListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEdgeApplicationTokenRequest() (request *DescribeEdgeApplicationTokenRequest) {
    request = &DescribeEdgeApplicationTokenRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("weilingwith", APIVersion, "DescribeEdgeApplicationToken")
    
    
    return
}

func NewDescribeEdgeApplicationTokenResponse() (response *DescribeEdgeApplicationTokenResponse) {
    response = &DescribeEdgeApplicationTokenResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeEdgeApplicationToken
// 查询边缘应用凭证
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeEdgeApplicationToken(request *DescribeEdgeApplicationTokenRequest) (response *DescribeEdgeApplicationTokenResponse, err error) {
    return c.DescribeEdgeApplicationTokenWithContext(context.Background(), request)
}

// DescribeEdgeApplicationToken
// 查询边缘应用凭证
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeEdgeApplicationTokenWithContext(ctx context.Context, request *DescribeEdgeApplicationTokenRequest) (response *DescribeEdgeApplicationTokenResponse, err error) {
    if request == nil {
        request = NewDescribeEdgeApplicationTokenRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeEdgeApplicationToken require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeEdgeApplicationTokenResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInterfaceListRequest() (request *DescribeInterfaceListRequest) {
    request = &DescribeInterfaceListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("weilingwith", APIVersion, "DescribeInterfaceList")
    
    
    return
}

func NewDescribeInterfaceListResponse() (response *DescribeInterfaceListResponse) {
    response = &DescribeInterfaceListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeInterfaceList
// 查询接口列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeInterfaceList(request *DescribeInterfaceListRequest) (response *DescribeInterfaceListResponse, err error) {
    return c.DescribeInterfaceListWithContext(context.Background(), request)
}

// DescribeInterfaceList
// 查询接口列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeInterfaceListWithContext(ctx context.Context, request *DescribeInterfaceListRequest) (response *DescribeInterfaceListResponse, err error) {
    if request == nil {
        request = NewDescribeInterfaceListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInterfaceList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInterfaceListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeWorkspaceListRequest() (request *DescribeWorkspaceListRequest) {
    request = &DescribeWorkspaceListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("weilingwith", APIVersion, "DescribeWorkspaceList")
    
    
    return
}

func NewDescribeWorkspaceListResponse() (response *DescribeWorkspaceListResponse) {
    response = &DescribeWorkspaceListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeWorkspaceList
// 获取租户下的空间列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDWORKSPACEID = "InvalidParameterValue.InvalidWorkspaceId"
func (c *Client) DescribeWorkspaceList(request *DescribeWorkspaceListRequest) (response *DescribeWorkspaceListResponse, err error) {
    return c.DescribeWorkspaceListWithContext(context.Background(), request)
}

// DescribeWorkspaceList
// 获取租户下的空间列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDWORKSPACEID = "InvalidParameterValue.InvalidWorkspaceId"
func (c *Client) DescribeWorkspaceListWithContext(ctx context.Context, request *DescribeWorkspaceListRequest) (response *DescribeWorkspaceListResponse, err error) {
    if request == nil {
        request = NewDescribeWorkspaceListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeWorkspaceList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeWorkspaceListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeWorkspaceUserListRequest() (request *DescribeWorkspaceUserListRequest) {
    request = &DescribeWorkspaceUserListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("weilingwith", APIVersion, "DescribeWorkspaceUserList")
    
    
    return
}

func NewDescribeWorkspaceUserListResponse() (response *DescribeWorkspaceUserListResponse) {
    response = &DescribeWorkspaceUserListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeWorkspaceUserList
// 查询项目空间人员列表
//
// 可能返回的错误码:
//  INTERNALERROR_APPAPINOSPACEPERMISSION = "InternalError.AppApiNoSpacePermission"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeWorkspaceUserList(request *DescribeWorkspaceUserListRequest) (response *DescribeWorkspaceUserListResponse, err error) {
    return c.DescribeWorkspaceUserListWithContext(context.Background(), request)
}

// DescribeWorkspaceUserList
// 查询项目空间人员列表
//
// 可能返回的错误码:
//  INTERNALERROR_APPAPINOSPACEPERMISSION = "InternalError.AppApiNoSpacePermission"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeWorkspaceUserListWithContext(ctx context.Context, request *DescribeWorkspaceUserListRequest) (response *DescribeWorkspaceUserListResponse, err error) {
    if request == nil {
        request = NewDescribeWorkspaceUserListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeWorkspaceUserList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeWorkspaceUserListResponse()
    err = c.Send(request, response)
    return
}
