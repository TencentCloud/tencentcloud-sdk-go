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

package v20180319

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2018-03-19"

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


func NewDeregisterMigrationTaskRequest() (request *DeregisterMigrationTaskRequest) {
    request = &DeregisterMigrationTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("msp", APIVersion, "DeregisterMigrationTask")
    
    
    return
}

func NewDeregisterMigrationTaskResponse() (response *DeregisterMigrationTaskResponse) {
    response = &DeregisterMigrationTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeregisterMigrationTask
// 取消注册迁移任务
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DeregisterMigrationTask(request *DeregisterMigrationTaskRequest) (response *DeregisterMigrationTaskResponse, err error) {
    return c.DeregisterMigrationTaskWithContext(context.Background(), request)
}

// DeregisterMigrationTask
// 取消注册迁移任务
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DeregisterMigrationTaskWithContext(ctx context.Context, request *DeregisterMigrationTaskRequest) (response *DeregisterMigrationTaskResponse, err error) {
    if request == nil {
        request = NewDeregisterMigrationTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeregisterMigrationTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeregisterMigrationTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMigrationTaskRequest() (request *DescribeMigrationTaskRequest) {
    request = &DescribeMigrationTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("msp", APIVersion, "DescribeMigrationTask")
    
    
    return
}

func NewDescribeMigrationTaskResponse() (response *DescribeMigrationTaskResponse) {
    response = &DescribeMigrationTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeMigrationTask
// 获取指定迁移任务详情
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeMigrationTask(request *DescribeMigrationTaskRequest) (response *DescribeMigrationTaskResponse, err error) {
    return c.DescribeMigrationTaskWithContext(context.Background(), request)
}

// DescribeMigrationTask
// 获取指定迁移任务详情
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeMigrationTaskWithContext(ctx context.Context, request *DescribeMigrationTaskRequest) (response *DescribeMigrationTaskResponse, err error) {
    if request == nil {
        request = NewDescribeMigrationTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMigrationTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMigrationTaskResponse()
    err = c.Send(request, response)
    return
}

func NewListMigrationProjectRequest() (request *ListMigrationProjectRequest) {
    request = &ListMigrationProjectRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("msp", APIVersion, "ListMigrationProject")
    
    
    return
}

func NewListMigrationProjectResponse() (response *ListMigrationProjectResponse) {
    response = &ListMigrationProjectResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ListMigrationProject
// 获取迁移项目名称列表
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) ListMigrationProject(request *ListMigrationProjectRequest) (response *ListMigrationProjectResponse, err error) {
    return c.ListMigrationProjectWithContext(context.Background(), request)
}

// ListMigrationProject
// 获取迁移项目名称列表
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) ListMigrationProjectWithContext(ctx context.Context, request *ListMigrationProjectRequest) (response *ListMigrationProjectResponse, err error) {
    if request == nil {
        request = NewListMigrationProjectRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListMigrationProject require credential")
    }

    request.SetContext(ctx)
    
    response = NewListMigrationProjectResponse()
    err = c.Send(request, response)
    return
}

func NewListMigrationTaskRequest() (request *ListMigrationTaskRequest) {
    request = &ListMigrationTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("msp", APIVersion, "ListMigrationTask")
    
    
    return
}

func NewListMigrationTaskResponse() (response *ListMigrationTaskResponse) {
    response = &ListMigrationTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ListMigrationTask
// 获取迁移任务列表
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) ListMigrationTask(request *ListMigrationTaskRequest) (response *ListMigrationTaskResponse, err error) {
    return c.ListMigrationTaskWithContext(context.Background(), request)
}

// ListMigrationTask
// 获取迁移任务列表
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) ListMigrationTaskWithContext(ctx context.Context, request *ListMigrationTaskRequest) (response *ListMigrationTaskResponse, err error) {
    if request == nil {
        request = NewListMigrationTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListMigrationTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewListMigrationTaskResponse()
    err = c.Send(request, response)
    return
}

func NewModifyMigrationTaskBelongToProjectRequest() (request *ModifyMigrationTaskBelongToProjectRequest) {
    request = &ModifyMigrationTaskBelongToProjectRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("msp", APIVersion, "ModifyMigrationTaskBelongToProject")
    
    
    return
}

func NewModifyMigrationTaskBelongToProjectResponse() (response *ModifyMigrationTaskBelongToProjectResponse) {
    response = &ModifyMigrationTaskBelongToProjectResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyMigrationTaskBelongToProject
// 更改迁移任务所属项目
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) ModifyMigrationTaskBelongToProject(request *ModifyMigrationTaskBelongToProjectRequest) (response *ModifyMigrationTaskBelongToProjectResponse, err error) {
    return c.ModifyMigrationTaskBelongToProjectWithContext(context.Background(), request)
}

// ModifyMigrationTaskBelongToProject
// 更改迁移任务所属项目
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) ModifyMigrationTaskBelongToProjectWithContext(ctx context.Context, request *ModifyMigrationTaskBelongToProjectRequest) (response *ModifyMigrationTaskBelongToProjectResponse, err error) {
    if request == nil {
        request = NewModifyMigrationTaskBelongToProjectRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyMigrationTaskBelongToProject require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyMigrationTaskBelongToProjectResponse()
    err = c.Send(request, response)
    return
}

func NewModifyMigrationTaskStatusRequest() (request *ModifyMigrationTaskStatusRequest) {
    request = &ModifyMigrationTaskStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("msp", APIVersion, "ModifyMigrationTaskStatus")
    
    
    return
}

func NewModifyMigrationTaskStatusResponse() (response *ModifyMigrationTaskStatusResponse) {
    response = &ModifyMigrationTaskStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyMigrationTaskStatus
// 更新迁移任务状态
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) ModifyMigrationTaskStatus(request *ModifyMigrationTaskStatusRequest) (response *ModifyMigrationTaskStatusResponse, err error) {
    return c.ModifyMigrationTaskStatusWithContext(context.Background(), request)
}

// ModifyMigrationTaskStatus
// 更新迁移任务状态
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) ModifyMigrationTaskStatusWithContext(ctx context.Context, request *ModifyMigrationTaskStatusRequest) (response *ModifyMigrationTaskStatusResponse, err error) {
    if request == nil {
        request = NewModifyMigrationTaskStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyMigrationTaskStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyMigrationTaskStatusResponse()
    err = c.Send(request, response)
    return
}

func NewRegisterMigrationTaskRequest() (request *RegisterMigrationTaskRequest) {
    request = &RegisterMigrationTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("msp", APIVersion, "RegisterMigrationTask")
    
    
    return
}

func NewRegisterMigrationTaskResponse() (response *RegisterMigrationTaskResponse) {
    response = &RegisterMigrationTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// RegisterMigrationTask
// 注册迁移任务
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) RegisterMigrationTask(request *RegisterMigrationTaskRequest) (response *RegisterMigrationTaskResponse, err error) {
    return c.RegisterMigrationTaskWithContext(context.Background(), request)
}

// RegisterMigrationTask
// 注册迁移任务
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) RegisterMigrationTaskWithContext(ctx context.Context, request *RegisterMigrationTaskRequest) (response *RegisterMigrationTaskResponse, err error) {
    if request == nil {
        request = NewRegisterMigrationTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RegisterMigrationTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewRegisterMigrationTaskResponse()
    err = c.Send(request, response)
    return
}
