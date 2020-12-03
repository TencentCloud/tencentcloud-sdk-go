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

package v20180326

import (
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2018-03-26"

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


func NewAddClusterInstancesRequest() (request *AddClusterInstancesRequest) {
    request = &AddClusterInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "AddClusterInstances")
    return
}

func NewAddClusterInstancesResponse() (response *AddClusterInstancesResponse) {
    response = &AddClusterInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 添加云主机节点至TSF集群
func (c *Client) AddClusterInstances(request *AddClusterInstancesRequest) (response *AddClusterInstancesResponse, err error) {
    if request == nil {
        request = NewAddClusterInstancesRequest()
    }
    response = NewAddClusterInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewAddInstancesRequest() (request *AddInstancesRequest) {
    request = &AddInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "AddInstances")
    return
}

func NewAddInstancesResponse() (response *AddInstancesResponse) {
    response = &AddInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 添加云主机节点至TSF集群
func (c *Client) AddInstances(request *AddInstancesRequest) (response *AddInstancesResponse, err error) {
    if request == nil {
        request = NewAddInstancesRequest()
    }
    response = NewAddInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewBindApiGroupRequest() (request *BindApiGroupRequest) {
    request = &BindApiGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "BindApiGroup")
    return
}

func NewBindApiGroupResponse() (response *BindApiGroupResponse) {
    response = &BindApiGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 网关与API分组批量绑定
func (c *Client) BindApiGroup(request *BindApiGroupRequest) (response *BindApiGroupResponse, err error) {
    if request == nil {
        request = NewBindApiGroupRequest()
    }
    response = NewBindApiGroupResponse()
    err = c.Send(request, response)
    return
}

func NewChangeApiUsableStatusRequest() (request *ChangeApiUsableStatusRequest) {
    request = &ChangeApiUsableStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "ChangeApiUsableStatus")
    return
}

func NewChangeApiUsableStatusResponse() (response *ChangeApiUsableStatusResponse) {
    response = &ChangeApiUsableStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 启用或禁用API
func (c *Client) ChangeApiUsableStatus(request *ChangeApiUsableStatusRequest) (response *ChangeApiUsableStatusResponse, err error) {
    if request == nil {
        request = NewChangeApiUsableStatusRequest()
    }
    response = NewChangeApiUsableStatusResponse()
    err = c.Send(request, response)
    return
}

func NewContinueRunFailedTaskBatchRequest() (request *ContinueRunFailedTaskBatchRequest) {
    request = &ContinueRunFailedTaskBatchRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "ContinueRunFailedTaskBatch")
    return
}

func NewContinueRunFailedTaskBatchResponse() (response *ContinueRunFailedTaskBatchResponse) {
    response = &ContinueRunFailedTaskBatchResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 对执行失败的任务批次执行续跑
func (c *Client) ContinueRunFailedTaskBatch(request *ContinueRunFailedTaskBatchRequest) (response *ContinueRunFailedTaskBatchResponse, err error) {
    if request == nil {
        request = NewContinueRunFailedTaskBatchRequest()
    }
    response = NewContinueRunFailedTaskBatchResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAllGatewayApiAsyncRequest() (request *CreateAllGatewayApiAsyncRequest) {
    request = &CreateAllGatewayApiAsyncRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "CreateAllGatewayApiAsync")
    return
}

func NewCreateAllGatewayApiAsyncResponse() (response *CreateAllGatewayApiAsyncResponse) {
    response = &CreateAllGatewayApiAsyncResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 一键导入API分组
func (c *Client) CreateAllGatewayApiAsync(request *CreateAllGatewayApiAsyncRequest) (response *CreateAllGatewayApiAsyncResponse, err error) {
    if request == nil {
        request = NewCreateAllGatewayApiAsyncRequest()
    }
    response = NewCreateAllGatewayApiAsyncResponse()
    err = c.Send(request, response)
    return
}

func NewCreateApiGroupRequest() (request *CreateApiGroupRequest) {
    request = &CreateApiGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "CreateApiGroup")
    return
}

func NewCreateApiGroupResponse() (response *CreateApiGroupResponse) {
    response = &CreateApiGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建API分组
func (c *Client) CreateApiGroup(request *CreateApiGroupRequest) (response *CreateApiGroupResponse, err error) {
    if request == nil {
        request = NewCreateApiGroupRequest()
    }
    response = NewCreateApiGroupResponse()
    err = c.Send(request, response)
    return
}

func NewCreateApiRateLimitRuleRequest() (request *CreateApiRateLimitRuleRequest) {
    request = &CreateApiRateLimitRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "CreateApiRateLimitRule")
    return
}

func NewCreateApiRateLimitRuleResponse() (response *CreateApiRateLimitRuleResponse) {
    response = &CreateApiRateLimitRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建API限流规则
func (c *Client) CreateApiRateLimitRule(request *CreateApiRateLimitRuleRequest) (response *CreateApiRateLimitRuleResponse, err error) {
    if request == nil {
        request = NewCreateApiRateLimitRuleRequest()
    }
    response = NewCreateApiRateLimitRuleResponse()
    err = c.Send(request, response)
    return
}

func NewCreateApplicationRequest() (request *CreateApplicationRequest) {
    request = &CreateApplicationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "CreateApplication")
    return
}

func NewCreateApplicationResponse() (response *CreateApplicationResponse) {
    response = &CreateApplicationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建应用
func (c *Client) CreateApplication(request *CreateApplicationRequest) (response *CreateApplicationResponse, err error) {
    if request == nil {
        request = NewCreateApplicationRequest()
    }
    response = NewCreateApplicationResponse()
    err = c.Send(request, response)
    return
}

func NewCreateClusterRequest() (request *CreateClusterRequest) {
    request = &CreateClusterRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "CreateCluster")
    return
}

func NewCreateClusterResponse() (response *CreateClusterResponse) {
    response = &CreateClusterResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建集群
func (c *Client) CreateCluster(request *CreateClusterRequest) (response *CreateClusterResponse, err error) {
    if request == nil {
        request = NewCreateClusterRequest()
    }
    response = NewCreateClusterResponse()
    err = c.Send(request, response)
    return
}

func NewCreateConfigRequest() (request *CreateConfigRequest) {
    request = &CreateConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "CreateConfig")
    return
}

func NewCreateConfigResponse() (response *CreateConfigResponse) {
    response = &CreateConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建配置项
func (c *Client) CreateConfig(request *CreateConfigRequest) (response *CreateConfigResponse, err error) {
    if request == nil {
        request = NewCreateConfigRequest()
    }
    response = NewCreateConfigResponse()
    err = c.Send(request, response)
    return
}

func NewCreateContainGroupRequest() (request *CreateContainGroupRequest) {
    request = &CreateContainGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "CreateContainGroup")
    return
}

func NewCreateContainGroupResponse() (response *CreateContainGroupResponse) {
    response = &CreateContainGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建容器部署组
func (c *Client) CreateContainGroup(request *CreateContainGroupRequest) (response *CreateContainGroupResponse, err error) {
    if request == nil {
        request = NewCreateContainGroupRequest()
    }
    response = NewCreateContainGroupResponse()
    err = c.Send(request, response)
    return
}

func NewCreateGatewayApiRequest() (request *CreateGatewayApiRequest) {
    request = &CreateGatewayApiRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "CreateGatewayApi")
    return
}

func NewCreateGatewayApiResponse() (response *CreateGatewayApiResponse) {
    response = &CreateGatewayApiResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 批量导入API至api分组(也支持新建API到分组)
func (c *Client) CreateGatewayApi(request *CreateGatewayApiRequest) (response *CreateGatewayApiResponse, err error) {
    if request == nil {
        request = NewCreateGatewayApiRequest()
    }
    response = NewCreateGatewayApiResponse()
    err = c.Send(request, response)
    return
}

func NewCreateGroupRequest() (request *CreateGroupRequest) {
    request = &CreateGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "CreateGroup")
    return
}

func NewCreateGroupResponse() (response *CreateGroupResponse) {
    response = &CreateGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建虚拟机部署组
func (c *Client) CreateGroup(request *CreateGroupRequest) (response *CreateGroupResponse, err error) {
    if request == nil {
        request = NewCreateGroupRequest()
    }
    response = NewCreateGroupResponse()
    err = c.Send(request, response)
    return
}

func NewCreateLaneRequest() (request *CreateLaneRequest) {
    request = &CreateLaneRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "CreateLane")
    return
}

func NewCreateLaneResponse() (response *CreateLaneResponse) {
    response = &CreateLaneResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建泳道
func (c *Client) CreateLane(request *CreateLaneRequest) (response *CreateLaneResponse, err error) {
    if request == nil {
        request = NewCreateLaneRequest()
    }
    response = NewCreateLaneResponse()
    err = c.Send(request, response)
    return
}

func NewCreateLaneRuleRequest() (request *CreateLaneRuleRequest) {
    request = &CreateLaneRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "CreateLaneRule")
    return
}

func NewCreateLaneRuleResponse() (response *CreateLaneRuleResponse) {
    response = &CreateLaneRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建泳道规则
func (c *Client) CreateLaneRule(request *CreateLaneRuleRequest) (response *CreateLaneRuleResponse, err error) {
    if request == nil {
        request = NewCreateLaneRuleRequest()
    }
    response = NewCreateLaneRuleResponse()
    err = c.Send(request, response)
    return
}

func NewCreateMicroserviceRequest() (request *CreateMicroserviceRequest) {
    request = &CreateMicroserviceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "CreateMicroservice")
    return
}

func NewCreateMicroserviceResponse() (response *CreateMicroserviceResponse) {
    response = &CreateMicroserviceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 新增微服务
func (c *Client) CreateMicroservice(request *CreateMicroserviceRequest) (response *CreateMicroserviceResponse, err error) {
    if request == nil {
        request = NewCreateMicroserviceRequest()
    }
    response = NewCreateMicroserviceResponse()
    err = c.Send(request, response)
    return
}

func NewCreateNamespaceRequest() (request *CreateNamespaceRequest) {
    request = &CreateNamespaceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "CreateNamespace")
    return
}

func NewCreateNamespaceResponse() (response *CreateNamespaceResponse) {
    response = &CreateNamespaceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建命名空间
func (c *Client) CreateNamespace(request *CreateNamespaceRequest) (response *CreateNamespaceResponse, err error) {
    if request == nil {
        request = NewCreateNamespaceRequest()
    }
    response = NewCreateNamespaceResponse()
    err = c.Send(request, response)
    return
}

func NewCreatePublicConfigRequest() (request *CreatePublicConfigRequest) {
    request = &CreatePublicConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "CreatePublicConfig")
    return
}

func NewCreatePublicConfigResponse() (response *CreatePublicConfigResponse) {
    response = &CreatePublicConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建公共配置项
func (c *Client) CreatePublicConfig(request *CreatePublicConfigRequest) (response *CreatePublicConfigResponse, err error) {
    if request == nil {
        request = NewCreatePublicConfigRequest()
    }
    response = NewCreatePublicConfigResponse()
    err = c.Send(request, response)
    return
}

func NewCreateRepositoryRequest() (request *CreateRepositoryRequest) {
    request = &CreateRepositoryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "CreateRepository")
    return
}

func NewCreateRepositoryResponse() (response *CreateRepositoryResponse) {
    response = &CreateRepositoryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建仓库
func (c *Client) CreateRepository(request *CreateRepositoryRequest) (response *CreateRepositoryResponse, err error) {
    if request == nil {
        request = NewCreateRepositoryRequest()
    }
    response = NewCreateRepositoryResponse()
    err = c.Send(request, response)
    return
}

func NewCreateServerlessGroupRequest() (request *CreateServerlessGroupRequest) {
    request = &CreateServerlessGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "CreateServerlessGroup")
    return
}

func NewCreateServerlessGroupResponse() (response *CreateServerlessGroupResponse) {
    response = &CreateServerlessGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建Serverless部署组
func (c *Client) CreateServerlessGroup(request *CreateServerlessGroupRequest) (response *CreateServerlessGroupResponse, err error) {
    if request == nil {
        request = NewCreateServerlessGroupRequest()
    }
    response = NewCreateServerlessGroupResponse()
    err = c.Send(request, response)
    return
}

func NewCreateTaskRequest() (request *CreateTaskRequest) {
    request = &CreateTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "CreateTask")
    return
}

func NewCreateTaskResponse() (response *CreateTaskResponse) {
    response = &CreateTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建任务
func (c *Client) CreateTask(request *CreateTaskRequest) (response *CreateTaskResponse, err error) {
    if request == nil {
        request = NewCreateTaskRequest()
    }
    response = NewCreateTaskResponse()
    err = c.Send(request, response)
    return
}

func NewCreateTaskFlowRequest() (request *CreateTaskFlowRequest) {
    request = &CreateTaskFlowRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "CreateTaskFlow")
    return
}

func NewCreateTaskFlowResponse() (response *CreateTaskFlowResponse) {
    response = &CreateTaskFlowResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建工作流
func (c *Client) CreateTaskFlow(request *CreateTaskFlowRequest) (response *CreateTaskFlowResponse, err error) {
    if request == nil {
        request = NewCreateTaskFlowRequest()
    }
    response = NewCreateTaskFlowResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteApiGroupRequest() (request *DeleteApiGroupRequest) {
    request = &DeleteApiGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DeleteApiGroup")
    return
}

func NewDeleteApiGroupResponse() (response *DeleteApiGroupResponse) {
    response = &DeleteApiGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除Api分组
func (c *Client) DeleteApiGroup(request *DeleteApiGroupRequest) (response *DeleteApiGroupResponse, err error) {
    if request == nil {
        request = NewDeleteApiGroupRequest()
    }
    response = NewDeleteApiGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteApplicationRequest() (request *DeleteApplicationRequest) {
    request = &DeleteApplicationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DeleteApplication")
    return
}

func NewDeleteApplicationResponse() (response *DeleteApplicationResponse) {
    response = &DeleteApplicationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除应用
func (c *Client) DeleteApplication(request *DeleteApplicationRequest) (response *DeleteApplicationResponse, err error) {
    if request == nil {
        request = NewDeleteApplicationRequest()
    }
    response = NewDeleteApplicationResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteConfigRequest() (request *DeleteConfigRequest) {
    request = &DeleteConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DeleteConfig")
    return
}

func NewDeleteConfigResponse() (response *DeleteConfigResponse) {
    response = &DeleteConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除配置项
func (c *Client) DeleteConfig(request *DeleteConfigRequest) (response *DeleteConfigResponse, err error) {
    if request == nil {
        request = NewDeleteConfigRequest()
    }
    response = NewDeleteConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteContainerGroupRequest() (request *DeleteContainerGroupRequest) {
    request = &DeleteContainerGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DeleteContainerGroup")
    return
}

func NewDeleteContainerGroupResponse() (response *DeleteContainerGroupResponse) {
    response = &DeleteContainerGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除容器部署组
func (c *Client) DeleteContainerGroup(request *DeleteContainerGroupRequest) (response *DeleteContainerGroupResponse, err error) {
    if request == nil {
        request = NewDeleteContainerGroupRequest()
    }
    response = NewDeleteContainerGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteGroupRequest() (request *DeleteGroupRequest) {
    request = &DeleteGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DeleteGroup")
    return
}

func NewDeleteGroupResponse() (response *DeleteGroupResponse) {
    response = &DeleteGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除容器部署组
func (c *Client) DeleteGroup(request *DeleteGroupRequest) (response *DeleteGroupResponse, err error) {
    if request == nil {
        request = NewDeleteGroupRequest()
    }
    response = NewDeleteGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteImageTagsRequest() (request *DeleteImageTagsRequest) {
    request = &DeleteImageTagsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DeleteImageTags")
    return
}

func NewDeleteImageTagsResponse() (response *DeleteImageTagsResponse) {
    response = &DeleteImageTagsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 批量删除镜像版本
func (c *Client) DeleteImageTags(request *DeleteImageTagsRequest) (response *DeleteImageTagsResponse, err error) {
    if request == nil {
        request = NewDeleteImageTagsRequest()
    }
    response = NewDeleteImageTagsResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteLaneRequest() (request *DeleteLaneRequest) {
    request = &DeleteLaneRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DeleteLane")
    return
}

func NewDeleteLaneResponse() (response *DeleteLaneResponse) {
    response = &DeleteLaneResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除泳道
func (c *Client) DeleteLane(request *DeleteLaneRequest) (response *DeleteLaneResponse, err error) {
    if request == nil {
        request = NewDeleteLaneRequest()
    }
    response = NewDeleteLaneResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteMicroserviceRequest() (request *DeleteMicroserviceRequest) {
    request = &DeleteMicroserviceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DeleteMicroservice")
    return
}

func NewDeleteMicroserviceResponse() (response *DeleteMicroserviceResponse) {
    response = &DeleteMicroserviceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除微服务
func (c *Client) DeleteMicroservice(request *DeleteMicroserviceRequest) (response *DeleteMicroserviceResponse, err error) {
    if request == nil {
        request = NewDeleteMicroserviceRequest()
    }
    response = NewDeleteMicroserviceResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteNamespaceRequest() (request *DeleteNamespaceRequest) {
    request = &DeleteNamespaceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DeleteNamespace")
    return
}

func NewDeleteNamespaceResponse() (response *DeleteNamespaceResponse) {
    response = &DeleteNamespaceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除命名空间
func (c *Client) DeleteNamespace(request *DeleteNamespaceRequest) (response *DeleteNamespaceResponse, err error) {
    if request == nil {
        request = NewDeleteNamespaceRequest()
    }
    response = NewDeleteNamespaceResponse()
    err = c.Send(request, response)
    return
}

func NewDeletePkgsRequest() (request *DeletePkgsRequest) {
    request = &DeletePkgsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DeletePkgs")
    return
}

func NewDeletePkgsResponse() (response *DeletePkgsResponse) {
    response = &DeletePkgsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 从软件仓库批量删除程序包。
// 一次最多支持删除1000个包，数量超过1000，返回UpperDeleteLimit错误。
func (c *Client) DeletePkgs(request *DeletePkgsRequest) (response *DeletePkgsResponse, err error) {
    if request == nil {
        request = NewDeletePkgsRequest()
    }
    response = NewDeletePkgsResponse()
    err = c.Send(request, response)
    return
}

func NewDeletePublicConfigRequest() (request *DeletePublicConfigRequest) {
    request = &DeletePublicConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DeletePublicConfig")
    return
}

func NewDeletePublicConfigResponse() (response *DeletePublicConfigResponse) {
    response = &DeletePublicConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除公共配置项
func (c *Client) DeletePublicConfig(request *DeletePublicConfigRequest) (response *DeletePublicConfigResponse, err error) {
    if request == nil {
        request = NewDeletePublicConfigRequest()
    }
    response = NewDeletePublicConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteRepositoryRequest() (request *DeleteRepositoryRequest) {
    request = &DeleteRepositoryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DeleteRepository")
    return
}

func NewDeleteRepositoryResponse() (response *DeleteRepositoryResponse) {
    response = &DeleteRepositoryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除仓库
func (c *Client) DeleteRepository(request *DeleteRepositoryRequest) (response *DeleteRepositoryResponse, err error) {
    if request == nil {
        request = NewDeleteRepositoryRequest()
    }
    response = NewDeleteRepositoryResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteServerlessGroupRequest() (request *DeleteServerlessGroupRequest) {
    request = &DeleteServerlessGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DeleteServerlessGroup")
    return
}

func NewDeleteServerlessGroupResponse() (response *DeleteServerlessGroupResponse) {
    response = &DeleteServerlessGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除Serverless部署组
func (c *Client) DeleteServerlessGroup(request *DeleteServerlessGroupRequest) (response *DeleteServerlessGroupResponse, err error) {
    if request == nil {
        request = NewDeleteServerlessGroupRequest()
    }
    response = NewDeleteServerlessGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteTaskRequest() (request *DeleteTaskRequest) {
    request = &DeleteTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DeleteTask")
    return
}

func NewDeleteTaskResponse() (response *DeleteTaskResponse) {
    response = &DeleteTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除任务
func (c *Client) DeleteTask(request *DeleteTaskRequest) (response *DeleteTaskResponse, err error) {
    if request == nil {
        request = NewDeleteTaskRequest()
    }
    response = NewDeleteTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDeployContainerGroupRequest() (request *DeployContainerGroupRequest) {
    request = &DeployContainerGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DeployContainerGroup")
    return
}

func NewDeployContainerGroupResponse() (response *DeployContainerGroupResponse) {
    response = &DeployContainerGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 部署容器应用
func (c *Client) DeployContainerGroup(request *DeployContainerGroupRequest) (response *DeployContainerGroupResponse, err error) {
    if request == nil {
        request = NewDeployContainerGroupRequest()
    }
    response = NewDeployContainerGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDeployGroupRequest() (request *DeployGroupRequest) {
    request = &DeployGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DeployGroup")
    return
}

func NewDeployGroupResponse() (response *DeployGroupResponse) {
    response = &DeployGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 部署虚拟机部署组应用
func (c *Client) DeployGroup(request *DeployGroupRequest) (response *DeployGroupResponse, err error) {
    if request == nil {
        request = NewDeployGroupRequest()
    }
    response = NewDeployGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDeployServerlessGroupRequest() (request *DeployServerlessGroupRequest) {
    request = &DeployServerlessGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DeployServerlessGroup")
    return
}

func NewDeployServerlessGroupResponse() (response *DeployServerlessGroupResponse) {
    response = &DeployServerlessGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 部署Serverless应用
func (c *Client) DeployServerlessGroup(request *DeployServerlessGroupRequest) (response *DeployServerlessGroupResponse, err error) {
    if request == nil {
        request = NewDeployServerlessGroupRequest()
    }
    response = NewDeployServerlessGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeApiDetailRequest() (request *DescribeApiDetailRequest) {
    request = &DescribeApiDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeApiDetail")
    return
}

func NewDescribeApiDetailResponse() (response *DescribeApiDetailResponse) {
    response = &DescribeApiDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询API详情
func (c *Client) DescribeApiDetail(request *DescribeApiDetailRequest) (response *DescribeApiDetailResponse, err error) {
    if request == nil {
        request = NewDescribeApiDetailRequest()
    }
    response = NewDescribeApiDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeApiGroupRequest() (request *DescribeApiGroupRequest) {
    request = &DescribeApiGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeApiGroup")
    return
}

func NewDescribeApiGroupResponse() (response *DescribeApiGroupResponse) {
    response = &DescribeApiGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询API分组
func (c *Client) DescribeApiGroup(request *DescribeApiGroupRequest) (response *DescribeApiGroupResponse, err error) {
    if request == nil {
        request = NewDescribeApiGroupRequest()
    }
    response = NewDescribeApiGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeApiGroupsRequest() (request *DescribeApiGroupsRequest) {
    request = &DescribeApiGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeApiGroups")
    return
}

func NewDescribeApiGroupsResponse() (response *DescribeApiGroupsResponse) {
    response = &DescribeApiGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询API 分组信息列表
func (c *Client) DescribeApiGroups(request *DescribeApiGroupsRequest) (response *DescribeApiGroupsResponse, err error) {
    if request == nil {
        request = NewDescribeApiGroupsRequest()
    }
    response = NewDescribeApiGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeApiRateLimitRulesRequest() (request *DescribeApiRateLimitRulesRequest) {
    request = &DescribeApiRateLimitRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeApiRateLimitRules")
    return
}

func NewDescribeApiRateLimitRulesResponse() (response *DescribeApiRateLimitRulesResponse) {
    response = &DescribeApiRateLimitRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询API限流规则
func (c *Client) DescribeApiRateLimitRules(request *DescribeApiRateLimitRulesRequest) (response *DescribeApiRateLimitRulesResponse, err error) {
    if request == nil {
        request = NewDescribeApiRateLimitRulesRequest()
    }
    response = NewDescribeApiRateLimitRulesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeApiUseDetailRequest() (request *DescribeApiUseDetailRequest) {
    request = &DescribeApiUseDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeApiUseDetail")
    return
}

func NewDescribeApiUseDetailResponse() (response *DescribeApiUseDetailResponse) {
    response = &DescribeApiUseDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询网关API监控明细数据
func (c *Client) DescribeApiUseDetail(request *DescribeApiUseDetailRequest) (response *DescribeApiUseDetailResponse, err error) {
    if request == nil {
        request = NewDescribeApiUseDetailRequest()
    }
    response = NewDescribeApiUseDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeApiVersionsRequest() (request *DescribeApiVersionsRequest) {
    request = &DescribeApiVersionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeApiVersions")
    return
}

func NewDescribeApiVersionsResponse() (response *DescribeApiVersionsResponse) {
    response = &DescribeApiVersionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询API 版本
func (c *Client) DescribeApiVersions(request *DescribeApiVersionsRequest) (response *DescribeApiVersionsResponse, err error) {
    if request == nil {
        request = NewDescribeApiVersionsRequest()
    }
    response = NewDescribeApiVersionsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeApplicationRequest() (request *DescribeApplicationRequest) {
    request = &DescribeApplicationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeApplication")
    return
}

func NewDescribeApplicationResponse() (response *DescribeApplicationResponse) {
    response = &DescribeApplicationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取应用详情
func (c *Client) DescribeApplication(request *DescribeApplicationRequest) (response *DescribeApplicationResponse, err error) {
    if request == nil {
        request = NewDescribeApplicationRequest()
    }
    response = NewDescribeApplicationResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeApplicationAttributeRequest() (request *DescribeApplicationAttributeRequest) {
    request = &DescribeApplicationAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeApplicationAttribute")
    return
}

func NewDescribeApplicationAttributeResponse() (response *DescribeApplicationAttributeResponse) {
    response = &DescribeApplicationAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取应用列表其它字段，如实例数量信息等
func (c *Client) DescribeApplicationAttribute(request *DescribeApplicationAttributeRequest) (response *DescribeApplicationAttributeResponse, err error) {
    if request == nil {
        request = NewDescribeApplicationAttributeRequest()
    }
    response = NewDescribeApplicationAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeApplicationsRequest() (request *DescribeApplicationsRequest) {
    request = &DescribeApplicationsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeApplications")
    return
}

func NewDescribeApplicationsResponse() (response *DescribeApplicationsResponse) {
    response = &DescribeApplicationsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取应用列表
func (c *Client) DescribeApplications(request *DescribeApplicationsRequest) (response *DescribeApplicationsResponse, err error) {
    if request == nil {
        request = NewDescribeApplicationsRequest()
    }
    response = NewDescribeApplicationsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBasicResourceUsageRequest() (request *DescribeBasicResourceUsageRequest) {
    request = &DescribeBasicResourceUsageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeBasicResourceUsage")
    return
}

func NewDescribeBasicResourceUsageResponse() (response *DescribeBasicResourceUsageResponse) {
    response = &DescribeBasicResourceUsageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// TSF基本资源信息概览接口
func (c *Client) DescribeBasicResourceUsage(request *DescribeBasicResourceUsageRequest) (response *DescribeBasicResourceUsageResponse, err error) {
    if request == nil {
        request = NewDescribeBasicResourceUsageRequest()
    }
    response = NewDescribeBasicResourceUsageResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeClusterInstancesRequest() (request *DescribeClusterInstancesRequest) {
    request = &DescribeClusterInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeClusterInstances")
    return
}

func NewDescribeClusterInstancesResponse() (response *DescribeClusterInstancesResponse) {
    response = &DescribeClusterInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询集群实例
func (c *Client) DescribeClusterInstances(request *DescribeClusterInstancesRequest) (response *DescribeClusterInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeClusterInstancesRequest()
    }
    response = NewDescribeClusterInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeConfigRequest() (request *DescribeConfigRequest) {
    request = &DescribeConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeConfig")
    return
}

func NewDescribeConfigResponse() (response *DescribeConfigResponse) {
    response = &DescribeConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询配置
func (c *Client) DescribeConfig(request *DescribeConfigRequest) (response *DescribeConfigResponse, err error) {
    if request == nil {
        request = NewDescribeConfigRequest()
    }
    response = NewDescribeConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeConfigReleaseLogsRequest() (request *DescribeConfigReleaseLogsRequest) {
    request = &DescribeConfigReleaseLogsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeConfigReleaseLogs")
    return
}

func NewDescribeConfigReleaseLogsResponse() (response *DescribeConfigReleaseLogsResponse) {
    response = &DescribeConfigReleaseLogsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询配置发布历史
func (c *Client) DescribeConfigReleaseLogs(request *DescribeConfigReleaseLogsRequest) (response *DescribeConfigReleaseLogsResponse, err error) {
    if request == nil {
        request = NewDescribeConfigReleaseLogsRequest()
    }
    response = NewDescribeConfigReleaseLogsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeConfigReleasesRequest() (request *DescribeConfigReleasesRequest) {
    request = &DescribeConfigReleasesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeConfigReleases")
    return
}

func NewDescribeConfigReleasesResponse() (response *DescribeConfigReleasesResponse) {
    response = &DescribeConfigReleasesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询配置发布信息
func (c *Client) DescribeConfigReleases(request *DescribeConfigReleasesRequest) (response *DescribeConfigReleasesResponse, err error) {
    if request == nil {
        request = NewDescribeConfigReleasesRequest()
    }
    response = NewDescribeConfigReleasesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeConfigSummaryRequest() (request *DescribeConfigSummaryRequest) {
    request = &DescribeConfigSummaryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeConfigSummary")
    return
}

func NewDescribeConfigSummaryResponse() (response *DescribeConfigSummaryResponse) {
    response = &DescribeConfigSummaryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询配置汇总列表
func (c *Client) DescribeConfigSummary(request *DescribeConfigSummaryRequest) (response *DescribeConfigSummaryResponse, err error) {
    if request == nil {
        request = NewDescribeConfigSummaryRequest()
    }
    response = NewDescribeConfigSummaryResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeConfigsRequest() (request *DescribeConfigsRequest) {
    request = &DescribeConfigsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeConfigs")
    return
}

func NewDescribeConfigsResponse() (response *DescribeConfigsResponse) {
    response = &DescribeConfigsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询配置项列表
func (c *Client) DescribeConfigs(request *DescribeConfigsRequest) (response *DescribeConfigsResponse, err error) {
    if request == nil {
        request = NewDescribeConfigsRequest()
    }
    response = NewDescribeConfigsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeContainerGroupDetailRequest() (request *DescribeContainerGroupDetailRequest) {
    request = &DescribeContainerGroupDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeContainerGroupDetail")
    return
}

func NewDescribeContainerGroupDetailResponse() (response *DescribeContainerGroupDetailResponse) {
    response = &DescribeContainerGroupDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

//  容器部署组详情
func (c *Client) DescribeContainerGroupDetail(request *DescribeContainerGroupDetailRequest) (response *DescribeContainerGroupDetailResponse, err error) {
    if request == nil {
        request = NewDescribeContainerGroupDetailRequest()
    }
    response = NewDescribeContainerGroupDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeContainerGroupsRequest() (request *DescribeContainerGroupsRequest) {
    request = &DescribeContainerGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeContainerGroups")
    return
}

func NewDescribeContainerGroupsResponse() (response *DescribeContainerGroupsResponse) {
    response = &DescribeContainerGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 容器部署组列表
func (c *Client) DescribeContainerGroups(request *DescribeContainerGroupsRequest) (response *DescribeContainerGroupsResponse, err error) {
    if request == nil {
        request = NewDescribeContainerGroupsRequest()
    }
    response = NewDescribeContainerGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCreateGatewayApiStatusRequest() (request *DescribeCreateGatewayApiStatusRequest) {
    request = &DescribeCreateGatewayApiStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeCreateGatewayApiStatus")
    return
}

func NewDescribeCreateGatewayApiStatusResponse() (response *DescribeCreateGatewayApiStatusResponse) {
    response = &DescribeCreateGatewayApiStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询一键导入API分组任务的状态
func (c *Client) DescribeCreateGatewayApiStatus(request *DescribeCreateGatewayApiStatusRequest) (response *DescribeCreateGatewayApiStatusResponse, err error) {
    if request == nil {
        request = NewDescribeCreateGatewayApiStatusRequest()
    }
    response = NewDescribeCreateGatewayApiStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDownloadInfoRequest() (request *DescribeDownloadInfoRequest) {
    request = &DescribeDownloadInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeDownloadInfo")
    return
}

func NewDescribeDownloadInfoResponse() (response *DescribeDownloadInfoResponse) {
    response = &DescribeDownloadInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// TSF上传的程序包存放在腾讯云对象存储（COS）中，通过该API可以获取从COS下载程序包需要的信息，包括包所在的桶、存储路径、鉴权信息等，之后使用COS API（或SDK）进行下载。
// COS相关文档请查阅：https://cloud.tencent.com/document/product/436
func (c *Client) DescribeDownloadInfo(request *DescribeDownloadInfoRequest) (response *DescribeDownloadInfoResponse, err error) {
    if request == nil {
        request = NewDescribeDownloadInfoRequest()
    }
    response = NewDescribeDownloadInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeFlowLastBatchStateRequest() (request *DescribeFlowLastBatchStateRequest) {
    request = &DescribeFlowLastBatchStateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeFlowLastBatchState")
    return
}

func NewDescribeFlowLastBatchStateResponse() (response *DescribeFlowLastBatchStateResponse) {
    response = &DescribeFlowLastBatchStateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询工作流最新一个批次的状态信息
func (c *Client) DescribeFlowLastBatchState(request *DescribeFlowLastBatchStateRequest) (response *DescribeFlowLastBatchStateResponse, err error) {
    if request == nil {
        request = NewDescribeFlowLastBatchStateRequest()
    }
    response = NewDescribeFlowLastBatchStateResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeGatewayAllGroupApisRequest() (request *DescribeGatewayAllGroupApisRequest) {
    request = &DescribeGatewayAllGroupApisRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeGatewayAllGroupApis")
    return
}

func NewDescribeGatewayAllGroupApisResponse() (response *DescribeGatewayAllGroupApisResponse) {
    response = &DescribeGatewayAllGroupApisResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询网关所有分组下Api列表
func (c *Client) DescribeGatewayAllGroupApis(request *DescribeGatewayAllGroupApisRequest) (response *DescribeGatewayAllGroupApisResponse, err error) {
    if request == nil {
        request = NewDescribeGatewayAllGroupApisRequest()
    }
    response = NewDescribeGatewayAllGroupApisResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeGatewayMonitorOverviewRequest() (request *DescribeGatewayMonitorOverviewRequest) {
    request = &DescribeGatewayMonitorOverviewRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeGatewayMonitorOverview")
    return
}

func NewDescribeGatewayMonitorOverviewResponse() (response *DescribeGatewayMonitorOverviewResponse) {
    response = &DescribeGatewayMonitorOverviewResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询网关监控概览
func (c *Client) DescribeGatewayMonitorOverview(request *DescribeGatewayMonitorOverviewRequest) (response *DescribeGatewayMonitorOverviewResponse, err error) {
    if request == nil {
        request = NewDescribeGatewayMonitorOverviewRequest()
    }
    response = NewDescribeGatewayMonitorOverviewResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeGroupRequest() (request *DescribeGroupRequest) {
    request = &DescribeGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeGroup")
    return
}

func NewDescribeGroupResponse() (response *DescribeGroupResponse) {
    response = &DescribeGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询虚拟机部署组详情
func (c *Client) DescribeGroup(request *DescribeGroupRequest) (response *DescribeGroupResponse, err error) {
    if request == nil {
        request = NewDescribeGroupRequest()
    }
    response = NewDescribeGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeGroupBindedGatewaysRequest() (request *DescribeGroupBindedGatewaysRequest) {
    request = &DescribeGroupBindedGatewaysRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeGroupBindedGateways")
    return
}

func NewDescribeGroupBindedGatewaysResponse() (response *DescribeGroupBindedGatewaysResponse) {
    response = &DescribeGroupBindedGatewaysResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询某个API分组已绑定的网关部署组信息列表
func (c *Client) DescribeGroupBindedGateways(request *DescribeGroupBindedGatewaysRequest) (response *DescribeGroupBindedGatewaysResponse, err error) {
    if request == nil {
        request = NewDescribeGroupBindedGatewaysRequest()
    }
    response = NewDescribeGroupBindedGatewaysResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeGroupGatewaysRequest() (request *DescribeGroupGatewaysRequest) {
    request = &DescribeGroupGatewaysRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeGroupGateways")
    return
}

func NewDescribeGroupGatewaysResponse() (response *DescribeGroupGatewaysResponse) {
    response = &DescribeGroupGatewaysResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询某个网关绑定的API 分组信息列表
func (c *Client) DescribeGroupGateways(request *DescribeGroupGatewaysRequest) (response *DescribeGroupGatewaysResponse, err error) {
    if request == nil {
        request = NewDescribeGroupGatewaysRequest()
    }
    response = NewDescribeGroupGatewaysResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeGroupInstancesRequest() (request *DescribeGroupInstancesRequest) {
    request = &DescribeGroupInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeGroupInstances")
    return
}

func NewDescribeGroupInstancesResponse() (response *DescribeGroupInstancesResponse) {
    response = &DescribeGroupInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询虚拟机部署组云主机列表
func (c *Client) DescribeGroupInstances(request *DescribeGroupInstancesRequest) (response *DescribeGroupInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeGroupInstancesRequest()
    }
    response = NewDescribeGroupInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeGroupUseDetailRequest() (request *DescribeGroupUseDetailRequest) {
    request = &DescribeGroupUseDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeGroupUseDetail")
    return
}

func NewDescribeGroupUseDetailResponse() (response *DescribeGroupUseDetailResponse) {
    response = &DescribeGroupUseDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询网关分组监控明细数据
func (c *Client) DescribeGroupUseDetail(request *DescribeGroupUseDetailRequest) (response *DescribeGroupUseDetailResponse, err error) {
    if request == nil {
        request = NewDescribeGroupUseDetailRequest()
    }
    response = NewDescribeGroupUseDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeGroupsRequest() (request *DescribeGroupsRequest) {
    request = &DescribeGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeGroups")
    return
}

func NewDescribeGroupsResponse() (response *DescribeGroupsResponse) {
    response = &DescribeGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取虚拟机部署组列表
func (c *Client) DescribeGroups(request *DescribeGroupsRequest) (response *DescribeGroupsResponse, err error) {
    if request == nil {
        request = NewDescribeGroupsRequest()
    }
    response = NewDescribeGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeImageRepositoryRequest() (request *DescribeImageRepositoryRequest) {
    request = &DescribeImageRepositoryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeImageRepository")
    return
}

func NewDescribeImageRepositoryResponse() (response *DescribeImageRepositoryResponse) {
    response = &DescribeImageRepositoryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 镜像仓库列表 
func (c *Client) DescribeImageRepository(request *DescribeImageRepositoryRequest) (response *DescribeImageRepositoryResponse, err error) {
    if request == nil {
        request = NewDescribeImageRepositoryRequest()
    }
    response = NewDescribeImageRepositoryResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeImageTagsRequest() (request *DescribeImageTagsRequest) {
    request = &DescribeImageTagsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeImageTags")
    return
}

func NewDescribeImageTagsResponse() (response *DescribeImageTagsResponse) {
    response = &DescribeImageTagsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 镜像版本列表
func (c *Client) DescribeImageTags(request *DescribeImageTagsRequest) (response *DescribeImageTagsResponse, err error) {
    if request == nil {
        request = NewDescribeImageTagsRequest()
    }
    response = NewDescribeImageTagsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLaneRulesRequest() (request *DescribeLaneRulesRequest) {
    request = &DescribeLaneRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeLaneRules")
    return
}

func NewDescribeLaneRulesResponse() (response *DescribeLaneRulesResponse) {
    response = &DescribeLaneRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询泳道规则列表
func (c *Client) DescribeLaneRules(request *DescribeLaneRulesRequest) (response *DescribeLaneRulesResponse, err error) {
    if request == nil {
        request = NewDescribeLaneRulesRequest()
    }
    response = NewDescribeLaneRulesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLanesRequest() (request *DescribeLanesRequest) {
    request = &DescribeLanesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeLanes")
    return
}

func NewDescribeLanesResponse() (response *DescribeLanesResponse) {
    response = &DescribeLanesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询泳道列表
func (c *Client) DescribeLanes(request *DescribeLanesRequest) (response *DescribeLanesResponse, err error) {
    if request == nil {
        request = NewDescribeLanesRequest()
    }
    response = NewDescribeLanesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMicroserviceRequest() (request *DescribeMicroserviceRequest) {
    request = &DescribeMicroserviceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeMicroservice")
    return
}

func NewDescribeMicroserviceResponse() (response *DescribeMicroserviceResponse) {
    response = &DescribeMicroserviceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询微服务详情
func (c *Client) DescribeMicroservice(request *DescribeMicroserviceRequest) (response *DescribeMicroserviceResponse, err error) {
    if request == nil {
        request = NewDescribeMicroserviceRequest()
    }
    response = NewDescribeMicroserviceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMicroservicesRequest() (request *DescribeMicroservicesRequest) {
    request = &DescribeMicroservicesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeMicroservices")
    return
}

func NewDescribeMicroservicesResponse() (response *DescribeMicroservicesResponse) {
    response = &DescribeMicroservicesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取微服务列表
func (c *Client) DescribeMicroservices(request *DescribeMicroservicesRequest) (response *DescribeMicroservicesResponse, err error) {
    if request == nil {
        request = NewDescribeMicroservicesRequest()
    }
    response = NewDescribeMicroservicesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMsApiListRequest() (request *DescribeMsApiListRequest) {
    request = &DescribeMsApiListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeMsApiList")
    return
}

func NewDescribeMsApiListResponse() (response *DescribeMsApiListResponse) {
    response = &DescribeMsApiListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询服务API列表
func (c *Client) DescribeMsApiList(request *DescribeMsApiListRequest) (response *DescribeMsApiListResponse, err error) {
    if request == nil {
        request = NewDescribeMsApiListRequest()
    }
    response = NewDescribeMsApiListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePkgsRequest() (request *DescribePkgsRequest) {
    request = &DescribePkgsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribePkgs")
    return
}

func NewDescribePkgsResponse() (response *DescribePkgsResponse) {
    response = &DescribePkgsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 无
func (c *Client) DescribePkgs(request *DescribePkgsRequest) (response *DescribePkgsResponse, err error) {
    if request == nil {
        request = NewDescribePkgsRequest()
    }
    response = NewDescribePkgsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePodInstancesRequest() (request *DescribePodInstancesRequest) {
    request = &DescribePodInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribePodInstances")
    return
}

func NewDescribePodInstancesResponse() (response *DescribePodInstancesResponse) {
    response = &DescribePodInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取部署组实例列表
func (c *Client) DescribePodInstances(request *DescribePodInstancesRequest) (response *DescribePodInstancesResponse, err error) {
    if request == nil {
        request = NewDescribePodInstancesRequest()
    }
    response = NewDescribePodInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePublicConfigRequest() (request *DescribePublicConfigRequest) {
    request = &DescribePublicConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribePublicConfig")
    return
}

func NewDescribePublicConfigResponse() (response *DescribePublicConfigResponse) {
    response = &DescribePublicConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询公共配置（单条）
func (c *Client) DescribePublicConfig(request *DescribePublicConfigRequest) (response *DescribePublicConfigResponse, err error) {
    if request == nil {
        request = NewDescribePublicConfigRequest()
    }
    response = NewDescribePublicConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePublicConfigReleaseLogsRequest() (request *DescribePublicConfigReleaseLogsRequest) {
    request = &DescribePublicConfigReleaseLogsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribePublicConfigReleaseLogs")
    return
}

func NewDescribePublicConfigReleaseLogsResponse() (response *DescribePublicConfigReleaseLogsResponse) {
    response = &DescribePublicConfigReleaseLogsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询公共配置发布历史
func (c *Client) DescribePublicConfigReleaseLogs(request *DescribePublicConfigReleaseLogsRequest) (response *DescribePublicConfigReleaseLogsResponse, err error) {
    if request == nil {
        request = NewDescribePublicConfigReleaseLogsRequest()
    }
    response = NewDescribePublicConfigReleaseLogsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePublicConfigReleasesRequest() (request *DescribePublicConfigReleasesRequest) {
    request = &DescribePublicConfigReleasesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribePublicConfigReleases")
    return
}

func NewDescribePublicConfigReleasesResponse() (response *DescribePublicConfigReleasesResponse) {
    response = &DescribePublicConfigReleasesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询公共配置发布信息
func (c *Client) DescribePublicConfigReleases(request *DescribePublicConfigReleasesRequest) (response *DescribePublicConfigReleasesResponse, err error) {
    if request == nil {
        request = NewDescribePublicConfigReleasesRequest()
    }
    response = NewDescribePublicConfigReleasesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePublicConfigSummaryRequest() (request *DescribePublicConfigSummaryRequest) {
    request = &DescribePublicConfigSummaryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribePublicConfigSummary")
    return
}

func NewDescribePublicConfigSummaryResponse() (response *DescribePublicConfigSummaryResponse) {
    response = &DescribePublicConfigSummaryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询公共配置汇总列表
func (c *Client) DescribePublicConfigSummary(request *DescribePublicConfigSummaryRequest) (response *DescribePublicConfigSummaryResponse, err error) {
    if request == nil {
        request = NewDescribePublicConfigSummaryRequest()
    }
    response = NewDescribePublicConfigSummaryResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePublicConfigsRequest() (request *DescribePublicConfigsRequest) {
    request = &DescribePublicConfigsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribePublicConfigs")
    return
}

func NewDescribePublicConfigsResponse() (response *DescribePublicConfigsResponse) {
    response = &DescribePublicConfigsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询公共配置项列表
func (c *Client) DescribePublicConfigs(request *DescribePublicConfigsRequest) (response *DescribePublicConfigsResponse, err error) {
    if request == nil {
        request = NewDescribePublicConfigsRequest()
    }
    response = NewDescribePublicConfigsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeReleasedConfigRequest() (request *DescribeReleasedConfigRequest) {
    request = &DescribeReleasedConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeReleasedConfig")
    return
}

func NewDescribeReleasedConfigResponse() (response *DescribeReleasedConfigResponse) {
    response = &DescribeReleasedConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询group发布的配置
func (c *Client) DescribeReleasedConfig(request *DescribeReleasedConfigRequest) (response *DescribeReleasedConfigResponse, err error) {
    if request == nil {
        request = NewDescribeReleasedConfigRequest()
    }
    response = NewDescribeReleasedConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRepositoriesRequest() (request *DescribeRepositoriesRequest) {
    request = &DescribeRepositoriesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeRepositories")
    return
}

func NewDescribeRepositoriesResponse() (response *DescribeRepositoriesResponse) {
    response = &DescribeRepositoriesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询仓库列表
func (c *Client) DescribeRepositories(request *DescribeRepositoriesRequest) (response *DescribeRepositoriesResponse, err error) {
    if request == nil {
        request = NewDescribeRepositoriesRequest()
    }
    response = NewDescribeRepositoriesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRepositoryRequest() (request *DescribeRepositoryRequest) {
    request = &DescribeRepositoryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeRepository")
    return
}

func NewDescribeRepositoryResponse() (response *DescribeRepositoryResponse) {
    response = &DescribeRepositoryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询仓库信息
func (c *Client) DescribeRepository(request *DescribeRepositoryRequest) (response *DescribeRepositoryResponse, err error) {
    if request == nil {
        request = NewDescribeRepositoryRequest()
    }
    response = NewDescribeRepositoryResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeServerlessGroupRequest() (request *DescribeServerlessGroupRequest) {
    request = &DescribeServerlessGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeServerlessGroup")
    return
}

func NewDescribeServerlessGroupResponse() (response *DescribeServerlessGroupResponse) {
    response = &DescribeServerlessGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询Serverless部署组明细
func (c *Client) DescribeServerlessGroup(request *DescribeServerlessGroupRequest) (response *DescribeServerlessGroupResponse, err error) {
    if request == nil {
        request = NewDescribeServerlessGroupRequest()
    }
    response = NewDescribeServerlessGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeServerlessGroupsRequest() (request *DescribeServerlessGroupsRequest) {
    request = &DescribeServerlessGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeServerlessGroups")
    return
}

func NewDescribeServerlessGroupsResponse() (response *DescribeServerlessGroupsResponse) {
    response = &DescribeServerlessGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询Serverless部署组列表
func (c *Client) DescribeServerlessGroups(request *DescribeServerlessGroupsRequest) (response *DescribeServerlessGroupsResponse, err error) {
    if request == nil {
        request = NewDescribeServerlessGroupsRequest()
    }
    response = NewDescribeServerlessGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSimpleApplicationsRequest() (request *DescribeSimpleApplicationsRequest) {
    request = &DescribeSimpleApplicationsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeSimpleApplications")
    return
}

func NewDescribeSimpleApplicationsResponse() (response *DescribeSimpleApplicationsResponse) {
    response = &DescribeSimpleApplicationsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询简单应用列表
func (c *Client) DescribeSimpleApplications(request *DescribeSimpleApplicationsRequest) (response *DescribeSimpleApplicationsResponse, err error) {
    if request == nil {
        request = NewDescribeSimpleApplicationsRequest()
    }
    response = NewDescribeSimpleApplicationsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSimpleClustersRequest() (request *DescribeSimpleClustersRequest) {
    request = &DescribeSimpleClustersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeSimpleClusters")
    return
}

func NewDescribeSimpleClustersResponse() (response *DescribeSimpleClustersResponse) {
    response = &DescribeSimpleClustersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询简单集群列表
func (c *Client) DescribeSimpleClusters(request *DescribeSimpleClustersRequest) (response *DescribeSimpleClustersResponse, err error) {
    if request == nil {
        request = NewDescribeSimpleClustersRequest()
    }
    response = NewDescribeSimpleClustersResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSimpleGroupsRequest() (request *DescribeSimpleGroupsRequest) {
    request = &DescribeSimpleGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeSimpleGroups")
    return
}

func NewDescribeSimpleGroupsResponse() (response *DescribeSimpleGroupsResponse) {
    response = &DescribeSimpleGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询简单部署组列表
func (c *Client) DescribeSimpleGroups(request *DescribeSimpleGroupsRequest) (response *DescribeSimpleGroupsResponse, err error) {
    if request == nil {
        request = NewDescribeSimpleGroupsRequest()
    }
    response = NewDescribeSimpleGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSimpleNamespacesRequest() (request *DescribeSimpleNamespacesRequest) {
    request = &DescribeSimpleNamespacesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeSimpleNamespaces")
    return
}

func NewDescribeSimpleNamespacesResponse() (response *DescribeSimpleNamespacesResponse) {
    response = &DescribeSimpleNamespacesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询简单命名空间列表 
func (c *Client) DescribeSimpleNamespaces(request *DescribeSimpleNamespacesRequest) (response *DescribeSimpleNamespacesResponse, err error) {
    if request == nil {
        request = NewDescribeSimpleNamespacesRequest()
    }
    response = NewDescribeSimpleNamespacesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTaskLastStatusRequest() (request *DescribeTaskLastStatusRequest) {
    request = &DescribeTaskLastStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeTaskLastStatus")
    return
}

func NewDescribeTaskLastStatusResponse() (response *DescribeTaskLastStatusResponse) {
    response = &DescribeTaskLastStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询任务最近一次执行状态
func (c *Client) DescribeTaskLastStatus(request *DescribeTaskLastStatusRequest) (response *DescribeTaskLastStatusResponse, err error) {
    if request == nil {
        request = NewDescribeTaskLastStatusRequest()
    }
    response = NewDescribeTaskLastStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUploadInfoRequest() (request *DescribeUploadInfoRequest) {
    request = &DescribeUploadInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeUploadInfo")
    return
}

func NewDescribeUploadInfoResponse() (response *DescribeUploadInfoResponse) {
    response = &DescribeUploadInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// TSF会将软件包上传到腾讯云对象存储（COS）。调用此接口获取上传信息，如目标地域，桶，包Id，存储路径，鉴权信息等，之后请使用COS API（或SDK）进行上传。
// COS相关文档请查阅：https://cloud.tencent.com/document/product/436
func (c *Client) DescribeUploadInfo(request *DescribeUploadInfoRequest) (response *DescribeUploadInfoResponse, err error) {
    if request == nil {
        request = NewDescribeUploadInfoRequest()
    }
    response = NewDescribeUploadInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDisableTaskRequest() (request *DisableTaskRequest) {
    request = &DisableTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DisableTask")
    return
}

func NewDisableTaskResponse() (response *DisableTaskResponse) {
    response = &DisableTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 停用任务
func (c *Client) DisableTask(request *DisableTaskRequest) (response *DisableTaskResponse, err error) {
    if request == nil {
        request = NewDisableTaskRequest()
    }
    response = NewDisableTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDisableTaskFlowRequest() (request *DisableTaskFlowRequest) {
    request = &DisableTaskFlowRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DisableTaskFlow")
    return
}

func NewDisableTaskFlowResponse() (response *DisableTaskFlowResponse) {
    response = &DisableTaskFlowResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 停用工作流
func (c *Client) DisableTaskFlow(request *DisableTaskFlowRequest) (response *DisableTaskFlowResponse, err error) {
    if request == nil {
        request = NewDisableTaskFlowRequest()
    }
    response = NewDisableTaskFlowResponse()
    err = c.Send(request, response)
    return
}

func NewDraftApiGroupRequest() (request *DraftApiGroupRequest) {
    request = &DraftApiGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DraftApiGroup")
    return
}

func NewDraftApiGroupResponse() (response *DraftApiGroupResponse) {
    response = &DraftApiGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 下线Api分组
func (c *Client) DraftApiGroup(request *DraftApiGroupRequest) (response *DraftApiGroupResponse, err error) {
    if request == nil {
        request = NewDraftApiGroupRequest()
    }
    response = NewDraftApiGroupResponse()
    err = c.Send(request, response)
    return
}

func NewEnableTaskRequest() (request *EnableTaskRequest) {
    request = &EnableTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "EnableTask")
    return
}

func NewEnableTaskResponse() (response *EnableTaskResponse) {
    response = &EnableTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 启用任务
func (c *Client) EnableTask(request *EnableTaskRequest) (response *EnableTaskResponse, err error) {
    if request == nil {
        request = NewEnableTaskRequest()
    }
    response = NewEnableTaskResponse()
    err = c.Send(request, response)
    return
}

func NewEnableTaskFlowRequest() (request *EnableTaskFlowRequest) {
    request = &EnableTaskFlowRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "EnableTaskFlow")
    return
}

func NewEnableTaskFlowResponse() (response *EnableTaskFlowResponse) {
    response = &EnableTaskFlowResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 启用工作流
func (c *Client) EnableTaskFlow(request *EnableTaskFlowRequest) (response *EnableTaskFlowResponse, err error) {
    if request == nil {
        request = NewEnableTaskFlowRequest()
    }
    response = NewEnableTaskFlowResponse()
    err = c.Send(request, response)
    return
}

func NewExecuteTaskRequest() (request *ExecuteTaskRequest) {
    request = &ExecuteTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "ExecuteTask")
    return
}

func NewExecuteTaskResponse() (response *ExecuteTaskResponse) {
    response = &ExecuteTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 手动执行一次任务。
func (c *Client) ExecuteTask(request *ExecuteTaskRequest) (response *ExecuteTaskResponse, err error) {
    if request == nil {
        request = NewExecuteTaskRequest()
    }
    response = NewExecuteTaskResponse()
    err = c.Send(request, response)
    return
}

func NewExecuteTaskFlowRequest() (request *ExecuteTaskFlowRequest) {
    request = &ExecuteTaskFlowRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "ExecuteTaskFlow")
    return
}

func NewExecuteTaskFlowResponse() (response *ExecuteTaskFlowResponse) {
    response = &ExecuteTaskFlowResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 执行一次工作流
func (c *Client) ExecuteTaskFlow(request *ExecuteTaskFlowRequest) (response *ExecuteTaskFlowResponse, err error) {
    if request == nil {
        request = NewExecuteTaskFlowRequest()
    }
    response = NewExecuteTaskFlowResponse()
    err = c.Send(request, response)
    return
}

func NewExpandGroupRequest() (request *ExpandGroupRequest) {
    request = &ExpandGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "ExpandGroup")
    return
}

func NewExpandGroupResponse() (response *ExpandGroupResponse) {
    response = &ExpandGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 虚拟机部署组添加实例
func (c *Client) ExpandGroup(request *ExpandGroupRequest) (response *ExpandGroupResponse, err error) {
    if request == nil {
        request = NewExpandGroupRequest()
    }
    response = NewExpandGroupResponse()
    err = c.Send(request, response)
    return
}

func NewModifyContainerGroupRequest() (request *ModifyContainerGroupRequest) {
    request = &ModifyContainerGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "ModifyContainerGroup")
    return
}

func NewModifyContainerGroupResponse() (response *ModifyContainerGroupResponse) {
    response = &ModifyContainerGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改容器部署组
func (c *Client) ModifyContainerGroup(request *ModifyContainerGroupRequest) (response *ModifyContainerGroupResponse, err error) {
    if request == nil {
        request = NewModifyContainerGroupRequest()
    }
    response = NewModifyContainerGroupResponse()
    err = c.Send(request, response)
    return
}

func NewModifyContainerReplicasRequest() (request *ModifyContainerReplicasRequest) {
    request = &ModifyContainerReplicasRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "ModifyContainerReplicas")
    return
}

func NewModifyContainerReplicasResponse() (response *ModifyContainerReplicasResponse) {
    response = &ModifyContainerReplicasResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改容器部署组实例数
func (c *Client) ModifyContainerReplicas(request *ModifyContainerReplicasRequest) (response *ModifyContainerReplicasResponse, err error) {
    if request == nil {
        request = NewModifyContainerReplicasRequest()
    }
    response = NewModifyContainerReplicasResponse()
    err = c.Send(request, response)
    return
}

func NewModifyLaneRequest() (request *ModifyLaneRequest) {
    request = &ModifyLaneRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "ModifyLane")
    return
}

func NewModifyLaneResponse() (response *ModifyLaneResponse) {
    response = &ModifyLaneResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 更新泳道信息
func (c *Client) ModifyLane(request *ModifyLaneRequest) (response *ModifyLaneResponse, err error) {
    if request == nil {
        request = NewModifyLaneRequest()
    }
    response = NewModifyLaneResponse()
    err = c.Send(request, response)
    return
}

func NewModifyLaneRuleRequest() (request *ModifyLaneRuleRequest) {
    request = &ModifyLaneRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "ModifyLaneRule")
    return
}

func NewModifyLaneRuleResponse() (response *ModifyLaneRuleResponse) {
    response = &ModifyLaneRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 更新泳道规则
func (c *Client) ModifyLaneRule(request *ModifyLaneRuleRequest) (response *ModifyLaneRuleResponse, err error) {
    if request == nil {
        request = NewModifyLaneRuleRequest()
    }
    response = NewModifyLaneRuleResponse()
    err = c.Send(request, response)
    return
}

func NewModifyMicroserviceRequest() (request *ModifyMicroserviceRequest) {
    request = &ModifyMicroserviceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "ModifyMicroservice")
    return
}

func NewModifyMicroserviceResponse() (response *ModifyMicroserviceResponse) {
    response = &ModifyMicroserviceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改微服务详情
func (c *Client) ModifyMicroservice(request *ModifyMicroserviceRequest) (response *ModifyMicroserviceResponse, err error) {
    if request == nil {
        request = NewModifyMicroserviceRequest()
    }
    response = NewModifyMicroserviceResponse()
    err = c.Send(request, response)
    return
}

func NewModifyTaskRequest() (request *ModifyTaskRequest) {
    request = &ModifyTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "ModifyTask")
    return
}

func NewModifyTaskResponse() (response *ModifyTaskResponse) {
    response = &ModifyTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改任务
func (c *Client) ModifyTask(request *ModifyTaskRequest) (response *ModifyTaskResponse, err error) {
    if request == nil {
        request = NewModifyTaskRequest()
    }
    response = NewModifyTaskResponse()
    err = c.Send(request, response)
    return
}

func NewModifyUploadInfoRequest() (request *ModifyUploadInfoRequest) {
    request = &ModifyUploadInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "ModifyUploadInfo")
    return
}

func NewModifyUploadInfoResponse() (response *ModifyUploadInfoResponse) {
    response = &ModifyUploadInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 调用该接口和COS的上传接口后，需要调用此接口更新TSF中保存的程序包状态。
// 调用此接口完成后，才标志上传包流程结束。
func (c *Client) ModifyUploadInfo(request *ModifyUploadInfoRequest) (response *ModifyUploadInfoResponse, err error) {
    if request == nil {
        request = NewModifyUploadInfoRequest()
    }
    response = NewModifyUploadInfoResponse()
    err = c.Send(request, response)
    return
}

func NewRedoTaskRequest() (request *RedoTaskRequest) {
    request = &RedoTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "RedoTask")
    return
}

func NewRedoTaskResponse() (response *RedoTaskResponse) {
    response = &RedoTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 重新执行任务
func (c *Client) RedoTask(request *RedoTaskRequest) (response *RedoTaskResponse, err error) {
    if request == nil {
        request = NewRedoTaskRequest()
    }
    response = NewRedoTaskResponse()
    err = c.Send(request, response)
    return
}

func NewRedoTaskBatchRequest() (request *RedoTaskBatchRequest) {
    request = &RedoTaskBatchRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "RedoTaskBatch")
    return
}

func NewRedoTaskBatchResponse() (response *RedoTaskBatchResponse) {
    response = &RedoTaskBatchResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 重新执行任务批次
func (c *Client) RedoTaskBatch(request *RedoTaskBatchRequest) (response *RedoTaskBatchResponse, err error) {
    if request == nil {
        request = NewRedoTaskBatchRequest()
    }
    response = NewRedoTaskBatchResponse()
    err = c.Send(request, response)
    return
}

func NewRedoTaskExecuteRequest() (request *RedoTaskExecuteRequest) {
    request = &RedoTaskExecuteRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "RedoTaskExecute")
    return
}

func NewRedoTaskExecuteResponse() (response *RedoTaskExecuteResponse) {
    response = &RedoTaskExecuteResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 重新执行在某个节点上执行任务。
func (c *Client) RedoTaskExecute(request *RedoTaskExecuteRequest) (response *RedoTaskExecuteResponse, err error) {
    if request == nil {
        request = NewRedoTaskExecuteRequest()
    }
    response = NewRedoTaskExecuteResponse()
    err = c.Send(request, response)
    return
}

func NewRedoTaskFlowBatchRequest() (request *RedoTaskFlowBatchRequest) {
    request = &RedoTaskFlowBatchRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "RedoTaskFlowBatch")
    return
}

func NewRedoTaskFlowBatchResponse() (response *RedoTaskFlowBatchResponse) {
    response = &RedoTaskFlowBatchResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 重新执行工作流批次
func (c *Client) RedoTaskFlowBatch(request *RedoTaskFlowBatchRequest) (response *RedoTaskFlowBatchResponse, err error) {
    if request == nil {
        request = NewRedoTaskFlowBatchRequest()
    }
    response = NewRedoTaskFlowBatchResponse()
    err = c.Send(request, response)
    return
}

func NewReleaseApiGroupRequest() (request *ReleaseApiGroupRequest) {
    request = &ReleaseApiGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "ReleaseApiGroup")
    return
}

func NewReleaseApiGroupResponse() (response *ReleaseApiGroupResponse) {
    response = &ReleaseApiGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 发布Api分组
func (c *Client) ReleaseApiGroup(request *ReleaseApiGroupRequest) (response *ReleaseApiGroupResponse, err error) {
    if request == nil {
        request = NewReleaseApiGroupRequest()
    }
    response = NewReleaseApiGroupResponse()
    err = c.Send(request, response)
    return
}

func NewReleaseConfigRequest() (request *ReleaseConfigRequest) {
    request = &ReleaseConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "ReleaseConfig")
    return
}

func NewReleaseConfigResponse() (response *ReleaseConfigResponse) {
    response = &ReleaseConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 发布配置
func (c *Client) ReleaseConfig(request *ReleaseConfigRequest) (response *ReleaseConfigResponse, err error) {
    if request == nil {
        request = NewReleaseConfigRequest()
    }
    response = NewReleaseConfigResponse()
    err = c.Send(request, response)
    return
}

func NewReleasePublicConfigRequest() (request *ReleasePublicConfigRequest) {
    request = &ReleasePublicConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "ReleasePublicConfig")
    return
}

func NewReleasePublicConfigResponse() (response *ReleasePublicConfigResponse) {
    response = &ReleasePublicConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 发布公共配置
func (c *Client) ReleasePublicConfig(request *ReleasePublicConfigRequest) (response *ReleasePublicConfigResponse, err error) {
    if request == nil {
        request = NewReleasePublicConfigRequest()
    }
    response = NewReleasePublicConfigResponse()
    err = c.Send(request, response)
    return
}

func NewRemoveInstancesRequest() (request *RemoveInstancesRequest) {
    request = &RemoveInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "RemoveInstances")
    return
}

func NewRemoveInstancesResponse() (response *RemoveInstancesResponse) {
    response = &RemoveInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 从 TSF 集群中批量移除云主机节点
func (c *Client) RemoveInstances(request *RemoveInstancesRequest) (response *RemoveInstancesResponse, err error) {
    if request == nil {
        request = NewRemoveInstancesRequest()
    }
    response = NewRemoveInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewRevocationConfigRequest() (request *RevocationConfigRequest) {
    request = &RevocationConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "RevocationConfig")
    return
}

func NewRevocationConfigResponse() (response *RevocationConfigResponse) {
    response = &RevocationConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 撤回已发布的配置
func (c *Client) RevocationConfig(request *RevocationConfigRequest) (response *RevocationConfigResponse, err error) {
    if request == nil {
        request = NewRevocationConfigRequest()
    }
    response = NewRevocationConfigResponse()
    err = c.Send(request, response)
    return
}

func NewRevocationPublicConfigRequest() (request *RevocationPublicConfigRequest) {
    request = &RevocationPublicConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "RevocationPublicConfig")
    return
}

func NewRevocationPublicConfigResponse() (response *RevocationPublicConfigResponse) {
    response = &RevocationPublicConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 撤回已发布的公共配置
func (c *Client) RevocationPublicConfig(request *RevocationPublicConfigRequest) (response *RevocationPublicConfigResponse, err error) {
    if request == nil {
        request = NewRevocationPublicConfigRequest()
    }
    response = NewRevocationPublicConfigResponse()
    err = c.Send(request, response)
    return
}

func NewRollbackConfigRequest() (request *RollbackConfigRequest) {
    request = &RollbackConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "RollbackConfig")
    return
}

func NewRollbackConfigResponse() (response *RollbackConfigResponse) {
    response = &RollbackConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 回滚配置
func (c *Client) RollbackConfig(request *RollbackConfigRequest) (response *RollbackConfigResponse, err error) {
    if request == nil {
        request = NewRollbackConfigRequest()
    }
    response = NewRollbackConfigResponse()
    err = c.Send(request, response)
    return
}

func NewShrinkGroupRequest() (request *ShrinkGroupRequest) {
    request = &ShrinkGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "ShrinkGroup")
    return
}

func NewShrinkGroupResponse() (response *ShrinkGroupResponse) {
    response = &ShrinkGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 下线部署组所有机器实例
func (c *Client) ShrinkGroup(request *ShrinkGroupRequest) (response *ShrinkGroupResponse, err error) {
    if request == nil {
        request = NewShrinkGroupRequest()
    }
    response = NewShrinkGroupResponse()
    err = c.Send(request, response)
    return
}

func NewShrinkInstancesRequest() (request *ShrinkInstancesRequest) {
    request = &ShrinkInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "ShrinkInstances")
    return
}

func NewShrinkInstancesResponse() (response *ShrinkInstancesResponse) {
    response = &ShrinkInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 虚拟机部署组下线实例
func (c *Client) ShrinkInstances(request *ShrinkInstancesRequest) (response *ShrinkInstancesResponse, err error) {
    if request == nil {
        request = NewShrinkInstancesRequest()
    }
    response = NewShrinkInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewStartContainerGroupRequest() (request *StartContainerGroupRequest) {
    request = &StartContainerGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "StartContainerGroup")
    return
}

func NewStartContainerGroupResponse() (response *StartContainerGroupResponse) {
    response = &StartContainerGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 启动容器部署组
func (c *Client) StartContainerGroup(request *StartContainerGroupRequest) (response *StartContainerGroupResponse, err error) {
    if request == nil {
        request = NewStartContainerGroupRequest()
    }
    response = NewStartContainerGroupResponse()
    err = c.Send(request, response)
    return
}

func NewStartGroupRequest() (request *StartGroupRequest) {
    request = &StartGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "StartGroup")
    return
}

func NewStartGroupResponse() (response *StartGroupResponse) {
    response = &StartGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 启动分组
func (c *Client) StartGroup(request *StartGroupRequest) (response *StartGroupResponse, err error) {
    if request == nil {
        request = NewStartGroupRequest()
    }
    response = NewStartGroupResponse()
    err = c.Send(request, response)
    return
}

func NewStopContainerGroupRequest() (request *StopContainerGroupRequest) {
    request = &StopContainerGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "StopContainerGroup")
    return
}

func NewStopContainerGroupResponse() (response *StopContainerGroupResponse) {
    response = &StopContainerGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 停止容器部署组
func (c *Client) StopContainerGroup(request *StopContainerGroupRequest) (response *StopContainerGroupResponse, err error) {
    if request == nil {
        request = NewStopContainerGroupRequest()
    }
    response = NewStopContainerGroupResponse()
    err = c.Send(request, response)
    return
}

func NewStopGroupRequest() (request *StopGroupRequest) {
    request = &StopGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "StopGroup")
    return
}

func NewStopGroupResponse() (response *StopGroupResponse) {
    response = &StopGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 停止虚拟机部署组
func (c *Client) StopGroup(request *StopGroupRequest) (response *StopGroupResponse, err error) {
    if request == nil {
        request = NewStopGroupRequest()
    }
    response = NewStopGroupResponse()
    err = c.Send(request, response)
    return
}

func NewStopTaskBatchRequest() (request *StopTaskBatchRequest) {
    request = &StopTaskBatchRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "StopTaskBatch")
    return
}

func NewStopTaskBatchResponse() (response *StopTaskBatchResponse) {
    response = &StopTaskBatchResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 停止执行中的任务批次， 非运行中的任务不可调用。
func (c *Client) StopTaskBatch(request *StopTaskBatchRequest) (response *StopTaskBatchResponse, err error) {
    if request == nil {
        request = NewStopTaskBatchRequest()
    }
    response = NewStopTaskBatchResponse()
    err = c.Send(request, response)
    return
}

func NewStopTaskExecuteRequest() (request *StopTaskExecuteRequest) {
    request = &StopTaskExecuteRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "StopTaskExecute")
    return
}

func NewStopTaskExecuteResponse() (response *StopTaskExecuteResponse) {
    response = &StopTaskExecuteResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 停止正在某个节点上执行的任务
func (c *Client) StopTaskExecute(request *StopTaskExecuteRequest) (response *StopTaskExecuteResponse, err error) {
    if request == nil {
        request = NewStopTaskExecuteRequest()
    }
    response = NewStopTaskExecuteResponse()
    err = c.Send(request, response)
    return
}

func NewTerminateTaskFlowBatchRequest() (request *TerminateTaskFlowBatchRequest) {
    request = &TerminateTaskFlowBatchRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "TerminateTaskFlowBatch")
    return
}

func NewTerminateTaskFlowBatchResponse() (response *TerminateTaskFlowBatchResponse) {
    response = &TerminateTaskFlowBatchResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 停止一个工作流批次
func (c *Client) TerminateTaskFlowBatch(request *TerminateTaskFlowBatchRequest) (response *TerminateTaskFlowBatchResponse, err error) {
    if request == nil {
        request = NewTerminateTaskFlowBatchRequest()
    }
    response = NewTerminateTaskFlowBatchResponse()
    err = c.Send(request, response)
    return
}

func NewUnbindApiGroupRequest() (request *UnbindApiGroupRequest) {
    request = &UnbindApiGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "UnbindApiGroup")
    return
}

func NewUnbindApiGroupResponse() (response *UnbindApiGroupResponse) {
    response = &UnbindApiGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// API分组批量与网关解绑
func (c *Client) UnbindApiGroup(request *UnbindApiGroupRequest) (response *UnbindApiGroupResponse, err error) {
    if request == nil {
        request = NewUnbindApiGroupRequest()
    }
    response = NewUnbindApiGroupResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateApiGroupRequest() (request *UpdateApiGroupRequest) {
    request = &UpdateApiGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "UpdateApiGroup")
    return
}

func NewUpdateApiGroupResponse() (response *UpdateApiGroupResponse) {
    response = &UpdateApiGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 更新Api分组
func (c *Client) UpdateApiGroup(request *UpdateApiGroupRequest) (response *UpdateApiGroupResponse, err error) {
    if request == nil {
        request = NewUpdateApiGroupRequest()
    }
    response = NewUpdateApiGroupResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateApiRateLimitRuleRequest() (request *UpdateApiRateLimitRuleRequest) {
    request = &UpdateApiRateLimitRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "UpdateApiRateLimitRule")
    return
}

func NewUpdateApiRateLimitRuleResponse() (response *UpdateApiRateLimitRuleResponse) {
    response = &UpdateApiRateLimitRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 更新API限流规则
func (c *Client) UpdateApiRateLimitRule(request *UpdateApiRateLimitRuleRequest) (response *UpdateApiRateLimitRuleResponse, err error) {
    if request == nil {
        request = NewUpdateApiRateLimitRuleRequest()
    }
    response = NewUpdateApiRateLimitRuleResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateApiRateLimitRulesRequest() (request *UpdateApiRateLimitRulesRequest) {
    request = &UpdateApiRateLimitRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "UpdateApiRateLimitRules")
    return
}

func NewUpdateApiRateLimitRulesResponse() (response *UpdateApiRateLimitRulesResponse) {
    response = &UpdateApiRateLimitRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 批量更新API限流规则
func (c *Client) UpdateApiRateLimitRules(request *UpdateApiRateLimitRulesRequest) (response *UpdateApiRateLimitRulesResponse, err error) {
    if request == nil {
        request = NewUpdateApiRateLimitRulesRequest()
    }
    response = NewUpdateApiRateLimitRulesResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateGatewayApiRequest() (request *UpdateGatewayApiRequest) {
    request = &UpdateGatewayApiRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "UpdateGatewayApi")
    return
}

func NewUpdateGatewayApiResponse() (response *UpdateGatewayApiResponse) {
    response = &UpdateGatewayApiResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 更新API
func (c *Client) UpdateGatewayApi(request *UpdateGatewayApiRequest) (response *UpdateGatewayApiResponse, err error) {
    if request == nil {
        request = NewUpdateGatewayApiRequest()
    }
    response = NewUpdateGatewayApiResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateHealthCheckSettingsRequest() (request *UpdateHealthCheckSettingsRequest) {
    request = &UpdateHealthCheckSettingsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "UpdateHealthCheckSettings")
    return
}

func NewUpdateHealthCheckSettingsResponse() (response *UpdateHealthCheckSettingsResponse) {
    response = &UpdateHealthCheckSettingsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 更新健康检查配置
func (c *Client) UpdateHealthCheckSettings(request *UpdateHealthCheckSettingsRequest) (response *UpdateHealthCheckSettingsResponse, err error) {
    if request == nil {
        request = NewUpdateHealthCheckSettingsRequest()
    }
    response = NewUpdateHealthCheckSettingsResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateRepositoryRequest() (request *UpdateRepositoryRequest) {
    request = &UpdateRepositoryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "UpdateRepository")
    return
}

func NewUpdateRepositoryResponse() (response *UpdateRepositoryResponse) {
    response = &UpdateRepositoryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 更新仓库信息
func (c *Client) UpdateRepository(request *UpdateRepositoryRequest) (response *UpdateRepositoryResponse, err error) {
    if request == nil {
        request = NewUpdateRepositoryRequest()
    }
    response = NewUpdateRepositoryResponse()
    err = c.Send(request, response)
    return
}
