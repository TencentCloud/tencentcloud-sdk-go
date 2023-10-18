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

package v20230508

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2023-05-08"

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


func NewCreateWorkspaceRequest() (request *CreateWorkspaceRequest) {
    request = &CreateWorkspaceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cloudstudio", APIVersion, "CreateWorkspace")
    
    
    return
}

func NewCreateWorkspaceResponse() (response *CreateWorkspaceResponse) {
    response = &CreateWorkspaceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateWorkspace
// 创建工作空间
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_WORKSPACENAMEDUPLICATE = "FailedOperation.WorkspaceNameDuplicate"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) CreateWorkspace(request *CreateWorkspaceRequest) (response *CreateWorkspaceResponse, err error) {
    return c.CreateWorkspaceWithContext(context.Background(), request)
}

// CreateWorkspace
// 创建工作空间
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_WORKSPACENAMEDUPLICATE = "FailedOperation.WorkspaceNameDuplicate"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) CreateWorkspaceWithContext(ctx context.Context, request *CreateWorkspaceRequest) (response *CreateWorkspaceResponse, err error) {
    if request == nil {
        request = NewCreateWorkspaceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateWorkspace require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateWorkspaceResponse()
    err = c.Send(request, response)
    return
}

func NewCreateWorkspaceTokenRequest() (request *CreateWorkspaceTokenRequest) {
    request = &CreateWorkspaceTokenRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cloudstudio", APIVersion, "CreateWorkspaceToken")
    
    
    return
}

func NewCreateWorkspaceTokenResponse() (response *CreateWorkspaceTokenResponse) {
    response = &CreateWorkspaceTokenResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateWorkspaceToken
// 创建工作空间临时访问凭证，重复调用会创建新的 Token，旧的 Token 将会自动失效
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) CreateWorkspaceToken(request *CreateWorkspaceTokenRequest) (response *CreateWorkspaceTokenResponse, err error) {
    return c.CreateWorkspaceTokenWithContext(context.Background(), request)
}

// CreateWorkspaceToken
// 创建工作空间临时访问凭证，重复调用会创建新的 Token，旧的 Token 将会自动失效
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) CreateWorkspaceTokenWithContext(ctx context.Context, request *CreateWorkspaceTokenRequest) (response *CreateWorkspaceTokenResponse, err error) {
    if request == nil {
        request = NewCreateWorkspaceTokenRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateWorkspaceToken require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateWorkspaceTokenResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeConfigRequest() (request *DescribeConfigRequest) {
    request = &DescribeConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cloudstudio", APIVersion, "DescribeConfig")
    
    
    return
}

func NewDescribeConfigResponse() (response *DescribeConfigResponse) {
    response = &DescribeConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeConfig
// 获取用户配置
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeConfig(request *DescribeConfigRequest) (response *DescribeConfigResponse, err error) {
    return c.DescribeConfigWithContext(context.Background(), request)
}

// DescribeConfig
// 获取用户配置
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeConfigWithContext(ctx context.Context, request *DescribeConfigRequest) (response *DescribeConfigResponse, err error) {
    if request == nil {
        request = NewDescribeConfigRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeImagesRequest() (request *DescribeImagesRequest) {
    request = &DescribeImagesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cloudstudio", APIVersion, "DescribeImages")
    
    
    return
}

func NewDescribeImagesResponse() (response *DescribeImagesResponse) {
    response = &DescribeImagesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeImages
// 获取基础镜像列表
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeImages(request *DescribeImagesRequest) (response *DescribeImagesResponse, err error) {
    return c.DescribeImagesWithContext(context.Background(), request)
}

// DescribeImages
// 获取基础镜像列表
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeImagesWithContext(ctx context.Context, request *DescribeImagesRequest) (response *DescribeImagesResponse, err error) {
    if request == nil {
        request = NewDescribeImagesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeImages require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeImagesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeWorkspacesRequest() (request *DescribeWorkspacesRequest) {
    request = &DescribeWorkspacesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cloudstudio", APIVersion, "DescribeWorkspaces")
    
    
    return
}

func NewDescribeWorkspacesResponse() (response *DescribeWorkspacesResponse) {
    response = &DescribeWorkspacesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeWorkspaces
// 获取用户工作空间列表
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeWorkspaces(request *DescribeWorkspacesRequest) (response *DescribeWorkspacesResponse, err error) {
    return c.DescribeWorkspacesWithContext(context.Background(), request)
}

// DescribeWorkspaces
// 获取用户工作空间列表
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeWorkspacesWithContext(ctx context.Context, request *DescribeWorkspacesRequest) (response *DescribeWorkspacesResponse, err error) {
    if request == nil {
        request = NewDescribeWorkspacesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeWorkspaces require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeWorkspacesResponse()
    err = c.Send(request, response)
    return
}

func NewModifyWorkspaceRequest() (request *ModifyWorkspaceRequest) {
    request = &ModifyWorkspaceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cloudstudio", APIVersion, "ModifyWorkspace")
    
    
    return
}

func NewModifyWorkspaceResponse() (response *ModifyWorkspaceResponse) {
    response = &ModifyWorkspaceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyWorkspace
// 修改工作空间
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_WORKSPACENAMEDUPLICATE = "FailedOperation.WorkspaceNameDuplicate"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) ModifyWorkspace(request *ModifyWorkspaceRequest) (response *ModifyWorkspaceResponse, err error) {
    return c.ModifyWorkspaceWithContext(context.Background(), request)
}

// ModifyWorkspace
// 修改工作空间
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_WORKSPACENAMEDUPLICATE = "FailedOperation.WorkspaceNameDuplicate"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) ModifyWorkspaceWithContext(ctx context.Context, request *ModifyWorkspaceRequest) (response *ModifyWorkspaceResponse, err error) {
    if request == nil {
        request = NewModifyWorkspaceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyWorkspace require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyWorkspaceResponse()
    err = c.Send(request, response)
    return
}

func NewRemoveWorkspaceRequest() (request *RemoveWorkspaceRequest) {
    request = &RemoveWorkspaceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cloudstudio", APIVersion, "RemoveWorkspace")
    
    
    return
}

func NewRemoveWorkspaceResponse() (response *RemoveWorkspaceResponse) {
    response = &RemoveWorkspaceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RemoveWorkspace
// 删除工作空间
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_WORKSPACENAMEDUPLICATE = "FailedOperation.WorkspaceNameDuplicate"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) RemoveWorkspace(request *RemoveWorkspaceRequest) (response *RemoveWorkspaceResponse, err error) {
    return c.RemoveWorkspaceWithContext(context.Background(), request)
}

// RemoveWorkspace
// 删除工作空间
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_WORKSPACENAMEDUPLICATE = "FailedOperation.WorkspaceNameDuplicate"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) RemoveWorkspaceWithContext(ctx context.Context, request *RemoveWorkspaceRequest) (response *RemoveWorkspaceResponse, err error) {
    if request == nil {
        request = NewRemoveWorkspaceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RemoveWorkspace require credential")
    }

    request.SetContext(ctx)
    
    response = NewRemoveWorkspaceResponse()
    err = c.Send(request, response)
    return
}

func NewRunWorkspaceRequest() (request *RunWorkspaceRequest) {
    request = &RunWorkspaceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cloudstudio", APIVersion, "RunWorkspace")
    
    
    return
}

func NewRunWorkspaceResponse() (response *RunWorkspaceResponse) {
    response = &RunWorkspaceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RunWorkspace
// 运行空间
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) RunWorkspace(request *RunWorkspaceRequest) (response *RunWorkspaceResponse, err error) {
    return c.RunWorkspaceWithContext(context.Background(), request)
}

// RunWorkspace
// 运行空间
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) RunWorkspaceWithContext(ctx context.Context, request *RunWorkspaceRequest) (response *RunWorkspaceResponse, err error) {
    if request == nil {
        request = NewRunWorkspaceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RunWorkspace require credential")
    }

    request.SetContext(ctx)
    
    response = NewRunWorkspaceResponse()
    err = c.Send(request, response)
    return
}

func NewStopWorkspaceRequest() (request *StopWorkspaceRequest) {
    request = &StopWorkspaceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cloudstudio", APIVersion, "StopWorkspace")
    
    
    return
}

func NewStopWorkspaceResponse() (response *StopWorkspaceResponse) {
    response = &StopWorkspaceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// StopWorkspace
// 停止运行空间
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) StopWorkspace(request *StopWorkspaceRequest) (response *StopWorkspaceResponse, err error) {
    return c.StopWorkspaceWithContext(context.Background(), request)
}

// StopWorkspace
// 停止运行空间
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) StopWorkspaceWithContext(ctx context.Context, request *StopWorkspaceRequest) (response *StopWorkspaceResponse, err error) {
    if request == nil {
        request = NewStopWorkspaceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("StopWorkspace require credential")
    }

    request.SetContext(ctx)
    
    response = NewStopWorkspaceResponse()
    err = c.Send(request, response)
    return
}
