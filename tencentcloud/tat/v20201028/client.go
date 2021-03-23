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

package v20201028

import (
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2020-10-28"

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


func NewCreateCommandRequest() (request *CreateCommandRequest) {
    request = &CreateCommandRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tat", APIVersion, "CreateCommand")
    return
}

func NewCreateCommandResponse() (response *CreateCommandResponse) {
    response = &CreateCommandResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 此接口用于创建命令。
func (c *Client) CreateCommand(request *CreateCommandRequest) (response *CreateCommandResponse, err error) {
    if request == nil {
        request = NewCreateCommandRequest()
    }
    response = NewCreateCommandResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteCommandRequest() (request *DeleteCommandRequest) {
    request = &DeleteCommandRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tat", APIVersion, "DeleteCommand")
    return
}

func NewDeleteCommandResponse() (response *DeleteCommandResponse) {
    response = &DeleteCommandResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 此接口用于删除命令。
func (c *Client) DeleteCommand(request *DeleteCommandRequest) (response *DeleteCommandResponse, err error) {
    if request == nil {
        request = NewDeleteCommandRequest()
    }
    response = NewDeleteCommandResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAutomationAgentStatusRequest() (request *DescribeAutomationAgentStatusRequest) {
    request = &DescribeAutomationAgentStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tat", APIVersion, "DescribeAutomationAgentStatus")
    return
}

func NewDescribeAutomationAgentStatusResponse() (response *DescribeAutomationAgentStatusResponse) {
    response = &DescribeAutomationAgentStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 此接口用于查询自动化助手客户端的状态。
func (c *Client) DescribeAutomationAgentStatus(request *DescribeAutomationAgentStatusRequest) (response *DescribeAutomationAgentStatusResponse, err error) {
    if request == nil {
        request = NewDescribeAutomationAgentStatusRequest()
    }
    response = NewDescribeAutomationAgentStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCommandsRequest() (request *DescribeCommandsRequest) {
    request = &DescribeCommandsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tat", APIVersion, "DescribeCommands")
    return
}

func NewDescribeCommandsResponse() (response *DescribeCommandsResponse) {
    response = &DescribeCommandsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 此接口用于查询命令详情。
func (c *Client) DescribeCommands(request *DescribeCommandsRequest) (response *DescribeCommandsResponse, err error) {
    if request == nil {
        request = NewDescribeCommandsRequest()
    }
    response = NewDescribeCommandsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInvocationTasksRequest() (request *DescribeInvocationTasksRequest) {
    request = &DescribeInvocationTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tat", APIVersion, "DescribeInvocationTasks")
    return
}

func NewDescribeInvocationTasksResponse() (response *DescribeInvocationTasksResponse) {
    response = &DescribeInvocationTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 此接口用于查询执行任务详情。
func (c *Client) DescribeInvocationTasks(request *DescribeInvocationTasksRequest) (response *DescribeInvocationTasksResponse, err error) {
    if request == nil {
        request = NewDescribeInvocationTasksRequest()
    }
    response = NewDescribeInvocationTasksResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInvocationsRequest() (request *DescribeInvocationsRequest) {
    request = &DescribeInvocationsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tat", APIVersion, "DescribeInvocations")
    return
}

func NewDescribeInvocationsResponse() (response *DescribeInvocationsResponse) {
    response = &DescribeInvocationsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 此接口用于查询执行活动详情。
func (c *Client) DescribeInvocations(request *DescribeInvocationsRequest) (response *DescribeInvocationsResponse, err error) {
    if request == nil {
        request = NewDescribeInvocationsRequest()
    }
    response = NewDescribeInvocationsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRegionsRequest() (request *DescribeRegionsRequest) {
    request = &DescribeRegionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tat", APIVersion, "DescribeRegions")
    return
}

func NewDescribeRegionsResponse() (response *DescribeRegionsResponse) {
    response = &DescribeRegionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 此接口用于查询地域列表
func (c *Client) DescribeRegions(request *DescribeRegionsRequest) (response *DescribeRegionsResponse, err error) {
    if request == nil {
        request = NewDescribeRegionsRequest()
    }
    response = NewDescribeRegionsResponse()
    err = c.Send(request, response)
    return
}

func NewInvokeCommandRequest() (request *InvokeCommandRequest) {
    request = &InvokeCommandRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tat", APIVersion, "InvokeCommand")
    return
}

func NewInvokeCommandResponse() (response *InvokeCommandResponse) {
    response = &InvokeCommandResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 在指定的实例上触发命令，调用成功返回执行活动ID（inv-xxxxxxxx），每个执行活动内部有一个或多个执行任务（invt-xxxxxxxx），每个执行任务代表命令在一台 CVM 或一台 Lighthouse 上的执行记录。
// 
// * 如果指定实例未安装 agent，或 agent 不在线，返回失败
// * 如果命令类型与 agent 运行环境不符，返回失败
// * 指定的实例需要处于 VPC 网络
// * 指定的实例需要处于 RUNNING 状态
// * 不可同时指定 CVM 和 Lighthouse
func (c *Client) InvokeCommand(request *InvokeCommandRequest) (response *InvokeCommandResponse, err error) {
    if request == nil {
        request = NewInvokeCommandRequest()
    }
    response = NewInvokeCommandResponse()
    err = c.Send(request, response)
    return
}

func NewModifyCommandRequest() (request *ModifyCommandRequest) {
    request = &ModifyCommandRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tat", APIVersion, "ModifyCommand")
    return
}

func NewModifyCommandResponse() (response *ModifyCommandResponse) {
    response = &ModifyCommandResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 此接口用于修改命令。
func (c *Client) ModifyCommand(request *ModifyCommandRequest) (response *ModifyCommandResponse, err error) {
    if request == nil {
        request = NewModifyCommandRequest()
    }
    response = NewModifyCommandResponse()
    err = c.Send(request, response)
    return
}

func NewPreviewReplacedCommandContentRequest() (request *PreviewReplacedCommandContentRequest) {
    request = &PreviewReplacedCommandContentRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tat", APIVersion, "PreviewReplacedCommandContent")
    return
}

func NewPreviewReplacedCommandContentResponse() (response *PreviewReplacedCommandContentResponse) {
    response = &PreviewReplacedCommandContentResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 此接口用于预览自定义参数替换后的命令内容。不会触发真实执行。
func (c *Client) PreviewReplacedCommandContent(request *PreviewReplacedCommandContentRequest) (response *PreviewReplacedCommandContentResponse, err error) {
    if request == nil {
        request = NewPreviewReplacedCommandContentRequest()
    }
    response = NewPreviewReplacedCommandContentResponse()
    err = c.Send(request, response)
    return
}

func NewRunCommandRequest() (request *RunCommandRequest) {
    request = &RunCommandRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tat", APIVersion, "RunCommand")
    return
}

func NewRunCommandResponse() (response *RunCommandResponse) {
    response = &RunCommandResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 执行命令，调用成功返回执行活动ID（inv-xxxxxxxx），每个执行活动内部有一个或多个执行任务（invt-xxxxxxxx），每个执行任务代表命令在一台 CVM 或一台 Lighthouse 上的执行记录。
// 
// * 如果指定实例未安装 agent，或 agent 不在线，返回失败
// * 如果命令类型与 agent 运行环境不符，返回失败
// * 指定的实例需要处于 VPC 网络
// * 指定的实例需要处于 `RUNNING` 状态
// * 不可同时指定 CVM 和 Lighthouse
func (c *Client) RunCommand(request *RunCommandRequest) (response *RunCommandResponse, err error) {
    if request == nil {
        request = NewRunCommandRequest()
    }
    response = NewRunCommandResponse()
    err = c.Send(request, response)
    return
}
