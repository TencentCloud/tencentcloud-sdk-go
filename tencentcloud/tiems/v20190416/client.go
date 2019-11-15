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

package v20190416

import (
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2019-04-16"

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


func NewCreateJobRequest() (request *CreateJobRequest) {
    request = &CreateJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tiems", APIVersion, "CreateJob")
    return
}

func NewCreateJobResponse() (response *CreateJobResponse) {
    response = &CreateJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建任务
func (c *Client) CreateJob(request *CreateJobRequest) (response *CreateJobResponse, err error) {
    if request == nil {
        request = NewCreateJobRequest()
    }
    response = NewCreateJobResponse()
    err = c.Send(request, response)
    return
}

func NewCreateRuntimeRequest() (request *CreateRuntimeRequest) {
    request = &CreateRuntimeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tiems", APIVersion, "CreateRuntime")
    return
}

func NewCreateRuntimeResponse() (response *CreateRuntimeResponse) {
    response = &CreateRuntimeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建运行环境
func (c *Client) CreateRuntime(request *CreateRuntimeRequest) (response *CreateRuntimeResponse, err error) {
    if request == nil {
        request = NewCreateRuntimeRequest()
    }
    response = NewCreateRuntimeResponse()
    err = c.Send(request, response)
    return
}

func NewCreateServiceRequest() (request *CreateServiceRequest) {
    request = &CreateServiceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tiems", APIVersion, "CreateService")
    return
}

func NewCreateServiceResponse() (response *CreateServiceResponse) {
    response = &CreateServiceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建服务
func (c *Client) CreateService(request *CreateServiceRequest) (response *CreateServiceResponse, err error) {
    if request == nil {
        request = NewCreateServiceRequest()
    }
    response = NewCreateServiceResponse()
    err = c.Send(request, response)
    return
}

func NewCreateServiceConfigRequest() (request *CreateServiceConfigRequest) {
    request = &CreateServiceConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tiems", APIVersion, "CreateServiceConfig")
    return
}

func NewCreateServiceConfigResponse() (response *CreateServiceConfigResponse) {
    response = &CreateServiceConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建服务配置
func (c *Client) CreateServiceConfig(request *CreateServiceConfigRequest) (response *CreateServiceConfigResponse, err error) {
    if request == nil {
        request = NewCreateServiceConfigRequest()
    }
    response = NewCreateServiceConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteJobRequest() (request *DeleteJobRequest) {
    request = &DeleteJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tiems", APIVersion, "DeleteJob")
    return
}

func NewDeleteJobResponse() (response *DeleteJobResponse) {
    response = &DeleteJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除任务
func (c *Client) DeleteJob(request *DeleteJobRequest) (response *DeleteJobResponse, err error) {
    if request == nil {
        request = NewDeleteJobRequest()
    }
    response = NewDeleteJobResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteRuntimeRequest() (request *DeleteRuntimeRequest) {
    request = &DeleteRuntimeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tiems", APIVersion, "DeleteRuntime")
    return
}

func NewDeleteRuntimeResponse() (response *DeleteRuntimeResponse) {
    response = &DeleteRuntimeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除运行环境
func (c *Client) DeleteRuntime(request *DeleteRuntimeRequest) (response *DeleteRuntimeResponse, err error) {
    if request == nil {
        request = NewDeleteRuntimeRequest()
    }
    response = NewDeleteRuntimeResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteServiceRequest() (request *DeleteServiceRequest) {
    request = &DeleteServiceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tiems", APIVersion, "DeleteService")
    return
}

func NewDeleteServiceResponse() (response *DeleteServiceResponse) {
    response = &DeleteServiceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除服务
func (c *Client) DeleteService(request *DeleteServiceRequest) (response *DeleteServiceResponse, err error) {
    if request == nil {
        request = NewDeleteServiceRequest()
    }
    response = NewDeleteServiceResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteServiceConfigRequest() (request *DeleteServiceConfigRequest) {
    request = &DeleteServiceConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tiems", APIVersion, "DeleteServiceConfig")
    return
}

func NewDeleteServiceConfigResponse() (response *DeleteServiceConfigResponse) {
    response = &DeleteServiceConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除服务配置
func (c *Client) DeleteServiceConfig(request *DeleteServiceConfigRequest) (response *DeleteServiceConfigResponse, err error) {
    if request == nil {
        request = NewDeleteServiceConfigRequest()
    }
    response = NewDeleteServiceConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstancesRequest() (request *DescribeInstancesRequest) {
    request = &DescribeInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tiems", APIVersion, "DescribeInstances")
    return
}

func NewDescribeInstancesResponse() (response *DescribeInstancesResponse) {
    response = &DescribeInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取节点列表
func (c *Client) DescribeInstances(request *DescribeInstancesRequest) (response *DescribeInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeInstancesRequest()
    }
    response = NewDescribeInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRuntimesRequest() (request *DescribeRuntimesRequest) {
    request = &DescribeRuntimesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tiems", APIVersion, "DescribeRuntimes")
    return
}

func NewDescribeRuntimesResponse() (response *DescribeRuntimesResponse) {
    response = &DescribeRuntimesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 描述服务运行环境
func (c *Client) DescribeRuntimes(request *DescribeRuntimesRequest) (response *DescribeRuntimesResponse, err error) {
    if request == nil {
        request = NewDescribeRuntimesRequest()
    }
    response = NewDescribeRuntimesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeServiceConfigsRequest() (request *DescribeServiceConfigsRequest) {
    request = &DescribeServiceConfigsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tiems", APIVersion, "DescribeServiceConfigs")
    return
}

func NewDescribeServiceConfigsResponse() (response *DescribeServiceConfigsResponse) {
    response = &DescribeServiceConfigsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 描述服务配置
func (c *Client) DescribeServiceConfigs(request *DescribeServiceConfigsRequest) (response *DescribeServiceConfigsResponse, err error) {
    if request == nil {
        request = NewDescribeServiceConfigsRequest()
    }
    response = NewDescribeServiceConfigsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeServicesRequest() (request *DescribeServicesRequest) {
    request = &DescribeServicesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tiems", APIVersion, "DescribeServices")
    return
}

func NewDescribeServicesResponse() (response *DescribeServicesResponse) {
    response = &DescribeServicesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 描述服务
func (c *Client) DescribeServices(request *DescribeServicesRequest) (response *DescribeServicesResponse, err error) {
    if request == nil {
        request = NewDescribeServicesRequest()
    }
    response = NewDescribeServicesResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateJobRequest() (request *UpdateJobRequest) {
    request = &UpdateJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tiems", APIVersion, "UpdateJob")
    return
}

func NewUpdateJobResponse() (response *UpdateJobResponse) {
    response = &UpdateJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 更新任务
func (c *Client) UpdateJob(request *UpdateJobRequest) (response *UpdateJobResponse, err error) {
    if request == nil {
        request = NewUpdateJobRequest()
    }
    response = NewUpdateJobResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateServiceRequest() (request *UpdateServiceRequest) {
    request = &UpdateServiceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tiems", APIVersion, "UpdateService")
    return
}

func NewUpdateServiceResponse() (response *UpdateServiceResponse) {
    response = &UpdateServiceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 更新服务
func (c *Client) UpdateService(request *UpdateServiceRequest) (response *UpdateServiceResponse, err error) {
    if request == nil {
        request = NewUpdateServiceRequest()
    }
    response = NewUpdateServiceResponse()
    err = c.Send(request, response)
    return
}
