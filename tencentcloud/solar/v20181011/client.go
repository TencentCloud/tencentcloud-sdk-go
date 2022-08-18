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

package v20181011

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2018-10-11"

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


func NewCheckStaffChUserRequest() (request *CheckStaffChUserRequest) {
    request = &CheckStaffChUserRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("solar", APIVersion, "CheckStaffChUser")
    
    
    return
}

func NewCheckStaffChUserResponse() (response *CheckStaffChUserResponse) {
    response = &CheckStaffChUserResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CheckStaffChUser
// 员工渠道更改员工状态
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_THIRDSERVERERROR = "InternalError.ThirdServerError"
func (c *Client) CheckStaffChUser(request *CheckStaffChUserRequest) (response *CheckStaffChUserResponse, err error) {
    return c.CheckStaffChUserWithContext(context.Background(), request)
}

// CheckStaffChUser
// 员工渠道更改员工状态
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_THIRDSERVERERROR = "InternalError.ThirdServerError"
func (c *Client) CheckStaffChUserWithContext(ctx context.Context, request *CheckStaffChUserRequest) (response *CheckStaffChUserResponse, err error) {
    if request == nil {
        request = NewCheckStaffChUserRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CheckStaffChUser require credential")
    }

    request.SetContext(ctx)
    
    response = NewCheckStaffChUserResponse()
    err = c.Send(request, response)
    return
}

func NewCopyActivityChannelRequest() (request *CopyActivityChannelRequest) {
    request = &CopyActivityChannelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("solar", APIVersion, "CopyActivityChannel")
    
    
    return
}

func NewCopyActivityChannelResponse() (response *CopyActivityChannelResponse) {
    response = &CopyActivityChannelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CopyActivityChannel
// 复制活动渠道的策略
//
// 可能返回的错误码:
//  FAILEDOPERATION_GENERATETOKENERROR = "FailedOperation.GenerateTokenError"
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND_CHANNEL = "ResourceNotFound.Channel"
func (c *Client) CopyActivityChannel(request *CopyActivityChannelRequest) (response *CopyActivityChannelResponse, err error) {
    return c.CopyActivityChannelWithContext(context.Background(), request)
}

// CopyActivityChannel
// 复制活动渠道的策略
//
// 可能返回的错误码:
//  FAILEDOPERATION_GENERATETOKENERROR = "FailedOperation.GenerateTokenError"
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND_CHANNEL = "ResourceNotFound.Channel"
func (c *Client) CopyActivityChannelWithContext(ctx context.Context, request *CopyActivityChannelRequest) (response *CopyActivityChannelResponse, err error) {
    if request == nil {
        request = NewCopyActivityChannelRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CopyActivityChannel require credential")
    }

    request.SetContext(ctx)
    
    response = NewCopyActivityChannelResponse()
    err = c.Send(request, response)
    return
}

func NewCreateProjectRequest() (request *CreateProjectRequest) {
    request = &CreateProjectRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("solar", APIVersion, "CreateProject")
    
    
    return
}

func NewCreateProjectResponse() (response *CreateProjectResponse) {
    response = &CreateProjectResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateProject
// 创建项目
//
// 可能返回的错误码:
//  FAILEDOPERATION_GENERATETOKENERROR = "FailedOperation.GenerateTokenError"
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND_CHANNEL = "ResourceNotFound.Channel"
func (c *Client) CreateProject(request *CreateProjectRequest) (response *CreateProjectResponse, err error) {
    return c.CreateProjectWithContext(context.Background(), request)
}

// CreateProject
// 创建项目
//
// 可能返回的错误码:
//  FAILEDOPERATION_GENERATETOKENERROR = "FailedOperation.GenerateTokenError"
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND_CHANNEL = "ResourceNotFound.Channel"
func (c *Client) CreateProjectWithContext(ctx context.Context, request *CreateProjectRequest) (response *CreateProjectResponse, err error) {
    if request == nil {
        request = NewCreateProjectRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateProject require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateProjectResponse()
    err = c.Send(request, response)
    return
}

func NewCreateSubProjectRequest() (request *CreateSubProjectRequest) {
    request = &CreateSubProjectRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("solar", APIVersion, "CreateSubProject")
    
    
    return
}

func NewCreateSubProjectResponse() (response *CreateSubProjectResponse) {
    response = &CreateSubProjectResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateSubProject
// 创建子项目
//
// 可能返回的错误码:
//  FAILEDOPERATION_GENERATETOKENERROR = "FailedOperation.GenerateTokenError"
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND_CHANNEL = "ResourceNotFound.Channel"
func (c *Client) CreateSubProject(request *CreateSubProjectRequest) (response *CreateSubProjectResponse, err error) {
    return c.CreateSubProjectWithContext(context.Background(), request)
}

// CreateSubProject
// 创建子项目
//
// 可能返回的错误码:
//  FAILEDOPERATION_GENERATETOKENERROR = "FailedOperation.GenerateTokenError"
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND_CHANNEL = "ResourceNotFound.Channel"
func (c *Client) CreateSubProjectWithContext(ctx context.Context, request *CreateSubProjectRequest) (response *CreateSubProjectResponse, err error) {
    if request == nil {
        request = NewCreateSubProjectRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateSubProject require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateSubProjectResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteProjectRequest() (request *DeleteProjectRequest) {
    request = &DeleteProjectRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("solar", APIVersion, "DeleteProject")
    
    
    return
}

func NewDeleteProjectResponse() (response *DeleteProjectResponse) {
    response = &DeleteProjectResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteProject
// 删除项目
//
// 可能返回的错误码:
//  UNSUPPORTEDOPERATION_PROJECTNOTALLOWEDTODELETE = "UnsupportedOperation.ProjectNotAllowedToDelete"
func (c *Client) DeleteProject(request *DeleteProjectRequest) (response *DeleteProjectResponse, err error) {
    return c.DeleteProjectWithContext(context.Background(), request)
}

// DeleteProject
// 删除项目
//
// 可能返回的错误码:
//  UNSUPPORTEDOPERATION_PROJECTNOTALLOWEDTODELETE = "UnsupportedOperation.ProjectNotAllowedToDelete"
func (c *Client) DeleteProjectWithContext(ctx context.Context, request *DeleteProjectRequest) (response *DeleteProjectResponse, err error) {
    if request == nil {
        request = NewDeleteProjectRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteProject require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteProjectResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCustomerRequest() (request *DescribeCustomerRequest) {
    request = &DescribeCustomerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("solar", APIVersion, "DescribeCustomer")
    
    
    return
}

func NewDescribeCustomerResponse() (response *DescribeCustomerResponse) {
    response = &DescribeCustomerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCustomer
// 客户档案查询客户详情
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeCustomer(request *DescribeCustomerRequest) (response *DescribeCustomerResponse, err error) {
    return c.DescribeCustomerWithContext(context.Background(), request)
}

// DescribeCustomer
// 客户档案查询客户详情
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeCustomerWithContext(ctx context.Context, request *DescribeCustomerRequest) (response *DescribeCustomerResponse, err error) {
    if request == nil {
        request = NewDescribeCustomerRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCustomer require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCustomerResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCustomersRequest() (request *DescribeCustomersRequest) {
    request = &DescribeCustomersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("solar", APIVersion, "DescribeCustomers")
    
    
    return
}

func NewDescribeCustomersResponse() (response *DescribeCustomersResponse) {
    response = &DescribeCustomersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCustomers
// 查询客户档案列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeCustomers(request *DescribeCustomersRequest) (response *DescribeCustomersResponse, err error) {
    return c.DescribeCustomersWithContext(context.Background(), request)
}

// DescribeCustomers
// 查询客户档案列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeCustomersWithContext(ctx context.Context, request *DescribeCustomersRequest) (response *DescribeCustomersResponse, err error) {
    if request == nil {
        request = NewDescribeCustomersRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCustomers require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCustomersResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeProjectRequest() (request *DescribeProjectRequest) {
    request = &DescribeProjectRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("solar", APIVersion, "DescribeProject")
    
    
    return
}

func NewDescribeProjectResponse() (response *DescribeProjectResponse) {
    response = &DescribeProjectResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeProject
// 项目详情展示
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_PROJECT = "ResourceNotFound.Project"
func (c *Client) DescribeProject(request *DescribeProjectRequest) (response *DescribeProjectResponse, err error) {
    return c.DescribeProjectWithContext(context.Background(), request)
}

// DescribeProject
// 项目详情展示
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_PROJECT = "ResourceNotFound.Project"
func (c *Client) DescribeProjectWithContext(ctx context.Context, request *DescribeProjectRequest) (response *DescribeProjectResponse, err error) {
    if request == nil {
        request = NewDescribeProjectRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeProject require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeProjectResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeProjectStockRequest() (request *DescribeProjectStockRequest) {
    request = &DescribeProjectStockRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("solar", APIVersion, "DescribeProjectStock")
    
    
    return
}

func NewDescribeProjectStockResponse() (response *DescribeProjectStockResponse) {
    response = &DescribeProjectStockResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeProjectStock
// 项目库存详情
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_PROJECT = "ResourceNotFound.Project"
func (c *Client) DescribeProjectStock(request *DescribeProjectStockRequest) (response *DescribeProjectStockResponse, err error) {
    return c.DescribeProjectStockWithContext(context.Background(), request)
}

// DescribeProjectStock
// 项目库存详情
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_PROJECT = "ResourceNotFound.Project"
func (c *Client) DescribeProjectStockWithContext(ctx context.Context, request *DescribeProjectStockRequest) (response *DescribeProjectStockResponse, err error) {
    if request == nil {
        request = NewDescribeProjectStockRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeProjectStock require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeProjectStockResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeProjectsRequest() (request *DescribeProjectsRequest) {
    request = &DescribeProjectsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("solar", APIVersion, "DescribeProjects")
    
    
    return
}

func NewDescribeProjectsResponse() (response *DescribeProjectsResponse) {
    response = &DescribeProjectsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeProjects
// 项目列表展示
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeProjects(request *DescribeProjectsRequest) (response *DescribeProjectsResponse, err error) {
    return c.DescribeProjectsWithContext(context.Background(), request)
}

// DescribeProjects
// 项目列表展示
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeProjectsWithContext(ctx context.Context, request *DescribeProjectsRequest) (response *DescribeProjectsResponse, err error) {
    if request == nil {
        request = NewDescribeProjectsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeProjects require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeProjectsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeResourceTemplateHeadersRequest() (request *DescribeResourceTemplateHeadersRequest) {
    request = &DescribeResourceTemplateHeadersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("solar", APIVersion, "DescribeResourceTemplateHeaders")
    
    
    return
}

func NewDescribeResourceTemplateHeadersResponse() (response *DescribeResourceTemplateHeadersResponse) {
    response = &DescribeResourceTemplateHeadersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeResourceTemplateHeaders
// 素材查询服务号模板的列表（样例）
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_THIRDSERVERERROR = "InternalError.ThirdServerError"
func (c *Client) DescribeResourceTemplateHeaders(request *DescribeResourceTemplateHeadersRequest) (response *DescribeResourceTemplateHeadersResponse, err error) {
    return c.DescribeResourceTemplateHeadersWithContext(context.Background(), request)
}

// DescribeResourceTemplateHeaders
// 素材查询服务号模板的列表（样例）
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_THIRDSERVERERROR = "InternalError.ThirdServerError"
func (c *Client) DescribeResourceTemplateHeadersWithContext(ctx context.Context, request *DescribeResourceTemplateHeadersRequest) (response *DescribeResourceTemplateHeadersResponse, err error) {
    if request == nil {
        request = NewDescribeResourceTemplateHeadersRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeResourceTemplateHeaders require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeResourceTemplateHeadersResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSubProjectRequest() (request *DescribeSubProjectRequest) {
    request = &DescribeSubProjectRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("solar", APIVersion, "DescribeSubProject")
    
    
    return
}

func NewDescribeSubProjectResponse() (response *DescribeSubProjectResponse) {
    response = &DescribeSubProjectResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeSubProject
// 子项目详情
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_THIRDSERVERERROR = "InternalError.ThirdServerError"
func (c *Client) DescribeSubProject(request *DescribeSubProjectRequest) (response *DescribeSubProjectResponse, err error) {
    return c.DescribeSubProjectWithContext(context.Background(), request)
}

// DescribeSubProject
// 子项目详情
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_THIRDSERVERERROR = "InternalError.ThirdServerError"
func (c *Client) DescribeSubProjectWithContext(ctx context.Context, request *DescribeSubProjectRequest) (response *DescribeSubProjectResponse, err error) {
    if request == nil {
        request = NewDescribeSubProjectRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSubProject require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSubProjectResponse()
    err = c.Send(request, response)
    return
}

func NewExpireFlowRequest() (request *ExpireFlowRequest) {
    request = &ExpireFlowRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("solar", APIVersion, "ExpireFlow")
    
    
    return
}

func NewExpireFlowResponse() (response *ExpireFlowResponse) {
    response = &ExpireFlowResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ExpireFlow
// 把审批中的工单置为已失效
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND_FLOW = "ResourceNotFound.Flow"
func (c *Client) ExpireFlow(request *ExpireFlowRequest) (response *ExpireFlowResponse, err error) {
    return c.ExpireFlowWithContext(context.Background(), request)
}

// ExpireFlow
// 把审批中的工单置为已失效
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND_FLOW = "ResourceNotFound.Flow"
func (c *Client) ExpireFlowWithContext(ctx context.Context, request *ExpireFlowRequest) (response *ExpireFlowResponse, err error) {
    if request == nil {
        request = NewExpireFlowRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ExpireFlow require credential")
    }

    request.SetContext(ctx)
    
    response = NewExpireFlowResponse()
    err = c.Send(request, response)
    return
}

func NewModifyProjectRequest() (request *ModifyProjectRequest) {
    request = &ModifyProjectRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("solar", APIVersion, "ModifyProject")
    
    
    return
}

func NewModifyProjectResponse() (response *ModifyProjectResponse) {
    response = &ModifyProjectResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyProject
// 修改项目
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND_FLOW = "ResourceNotFound.Flow"
func (c *Client) ModifyProject(request *ModifyProjectRequest) (response *ModifyProjectResponse, err error) {
    return c.ModifyProjectWithContext(context.Background(), request)
}

// ModifyProject
// 修改项目
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND_FLOW = "ResourceNotFound.Flow"
func (c *Client) ModifyProjectWithContext(ctx context.Context, request *ModifyProjectRequest) (response *ModifyProjectResponse, err error) {
    if request == nil {
        request = NewModifyProjectRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyProject require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyProjectResponse()
    err = c.Send(request, response)
    return
}

func NewOffLineProjectRequest() (request *OffLineProjectRequest) {
    request = &OffLineProjectRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("solar", APIVersion, "OffLineProject")
    
    
    return
}

func NewOffLineProjectResponse() (response *OffLineProjectResponse) {
    response = &OffLineProjectResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// OffLineProject
// 下线项目
//
// 可能返回的错误码:
//  UNSUPPORTEDOPERATION_STATUSOFFLINEPROJECT = "UnsupportedOperation.StatusOffLineProject"
func (c *Client) OffLineProject(request *OffLineProjectRequest) (response *OffLineProjectResponse, err error) {
    return c.OffLineProjectWithContext(context.Background(), request)
}

// OffLineProject
// 下线项目
//
// 可能返回的错误码:
//  UNSUPPORTEDOPERATION_STATUSOFFLINEPROJECT = "UnsupportedOperation.StatusOffLineProject"
func (c *Client) OffLineProjectWithContext(ctx context.Context, request *OffLineProjectRequest) (response *OffLineProjectResponse, err error) {
    if request == nil {
        request = NewOffLineProjectRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("OffLineProject require credential")
    }

    request.SetContext(ctx)
    
    response = NewOffLineProjectResponse()
    err = c.Send(request, response)
    return
}

func NewReplenishProjectStockRequest() (request *ReplenishProjectStockRequest) {
    request = &ReplenishProjectStockRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("solar", APIVersion, "ReplenishProjectStock")
    
    
    return
}

func NewReplenishProjectStockResponse() (response *ReplenishProjectStockResponse) {
    response = &ReplenishProjectStockResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ReplenishProjectStock
// 补充子项目库存
//
// 可能返回的错误码:
//  FAILEDOPERATION_INSUFFICIENTPRIZESTOCK = "FailedOperation.InsufficientPrizeStock"
func (c *Client) ReplenishProjectStock(request *ReplenishProjectStockRequest) (response *ReplenishProjectStockResponse, err error) {
    return c.ReplenishProjectStockWithContext(context.Background(), request)
}

// ReplenishProjectStock
// 补充子项目库存
//
// 可能返回的错误码:
//  FAILEDOPERATION_INSUFFICIENTPRIZESTOCK = "FailedOperation.InsufficientPrizeStock"
func (c *Client) ReplenishProjectStockWithContext(ctx context.Context, request *ReplenishProjectStockRequest) (response *ReplenishProjectStockResponse, err error) {
    if request == nil {
        request = NewReplenishProjectStockRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ReplenishProjectStock require credential")
    }

    request.SetContext(ctx)
    
    response = NewReplenishProjectStockResponse()
    err = c.Send(request, response)
    return
}

func NewSendWxTouchTaskRequest() (request *SendWxTouchTaskRequest) {
    request = &SendWxTouchTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("solar", APIVersion, "SendWxTouchTask")
    
    
    return
}

func NewSendWxTouchTaskResponse() (response *SendWxTouchTaskResponse) {
    response = &SendWxTouchTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// SendWxTouchTask
// 发送企业微信触达任务
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_THIRDSERVERERROR = "InternalError.ThirdServerError"
func (c *Client) SendWxTouchTask(request *SendWxTouchTaskRequest) (response *SendWxTouchTaskResponse, err error) {
    return c.SendWxTouchTaskWithContext(context.Background(), request)
}

// SendWxTouchTask
// 发送企业微信触达任务
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_THIRDSERVERERROR = "InternalError.ThirdServerError"
func (c *Client) SendWxTouchTaskWithContext(ctx context.Context, request *SendWxTouchTaskRequest) (response *SendWxTouchTaskResponse, err error) {
    if request == nil {
        request = NewSendWxTouchTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SendWxTouchTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewSendWxTouchTaskResponse()
    err = c.Send(request, response)
    return
}
