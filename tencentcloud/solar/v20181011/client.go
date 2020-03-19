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

func NewClient(credential *common.Credential, region string, clientProfile *profile.ClientProfile) (client *Client, err error) {
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

// 员工渠道更改员工状态
func (c *Client) CheckStaffChUser(request *CheckStaffChUserRequest) (response *CheckStaffChUserResponse, err error) {
    if request == nil {
        request = NewCheckStaffChUserRequest()
    }
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

// 复制活动渠道的策略
func (c *Client) CopyActivityChannel(request *CopyActivityChannelRequest) (response *CopyActivityChannelResponse, err error) {
    if request == nil {
        request = NewCopyActivityChannelRequest()
    }
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

// 创建项目
func (c *Client) CreateProject(request *CreateProjectRequest) (response *CreateProjectResponse, err error) {
    if request == nil {
        request = NewCreateProjectRequest()
    }
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

// 创建子项目
func (c *Client) CreateSubProject(request *CreateSubProjectRequest) (response *CreateSubProjectResponse, err error) {
    if request == nil {
        request = NewCreateSubProjectRequest()
    }
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

// 删除项目
func (c *Client) DeleteProject(request *DeleteProjectRequest) (response *DeleteProjectResponse, err error) {
    if request == nil {
        request = NewDeleteProjectRequest()
    }
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

// 客户档案查询客户详情
func (c *Client) DescribeCustomer(request *DescribeCustomerRequest) (response *DescribeCustomerResponse, err error) {
    if request == nil {
        request = NewDescribeCustomerRequest()
    }
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

// 查询客户档案列表
func (c *Client) DescribeCustomers(request *DescribeCustomersRequest) (response *DescribeCustomersResponse, err error) {
    if request == nil {
        request = NewDescribeCustomersRequest()
    }
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

// 项目详情展示
func (c *Client) DescribeProject(request *DescribeProjectRequest) (response *DescribeProjectResponse, err error) {
    if request == nil {
        request = NewDescribeProjectRequest()
    }
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

// 项目库存详情
func (c *Client) DescribeProjectStock(request *DescribeProjectStockRequest) (response *DescribeProjectStockResponse, err error) {
    if request == nil {
        request = NewDescribeProjectStockRequest()
    }
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

// 项目列表展示
func (c *Client) DescribeProjects(request *DescribeProjectsRequest) (response *DescribeProjectsResponse, err error) {
    if request == nil {
        request = NewDescribeProjectsRequest()
    }
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

// 素材查询服务号模板的列表（样例）
func (c *Client) DescribeResourceTemplateHeaders(request *DescribeResourceTemplateHeadersRequest) (response *DescribeResourceTemplateHeadersResponse, err error) {
    if request == nil {
        request = NewDescribeResourceTemplateHeadersRequest()
    }
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

// 子项目详情
func (c *Client) DescribeSubProject(request *DescribeSubProjectRequest) (response *DescribeSubProjectResponse, err error) {
    if request == nil {
        request = NewDescribeSubProjectRequest()
    }
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

// 把审批中的工单置为已失效
func (c *Client) ExpireFlow(request *ExpireFlowRequest) (response *ExpireFlowResponse, err error) {
    if request == nil {
        request = NewExpireFlowRequest()
    }
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

// 修改项目
func (c *Client) ModifyProject(request *ModifyProjectRequest) (response *ModifyProjectResponse, err error) {
    if request == nil {
        request = NewModifyProjectRequest()
    }
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

// 下线项目
func (c *Client) OffLineProject(request *OffLineProjectRequest) (response *OffLineProjectResponse, err error) {
    if request == nil {
        request = NewOffLineProjectRequest()
    }
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

// 补充子项目库存
func (c *Client) ReplenishProjectStock(request *ReplenishProjectStockRequest) (response *ReplenishProjectStockResponse, err error) {
    if request == nil {
        request = NewReplenishProjectStockRequest()
    }
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

// 发送企业微信触达任务
func (c *Client) SendWxTouchTask(request *SendWxTouchTaskRequest) (response *SendWxTouchTaskResponse, err error) {
    if request == nil {
        request = NewSendWxTouchTaskRequest()
    }
    response = NewSendWxTouchTaskResponse()
    err = c.Send(request, response)
    return
}
