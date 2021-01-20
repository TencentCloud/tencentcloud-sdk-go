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

package v20200217

import (
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2020-02-17"

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


func NewCreateClusterRequest() (request *CreateClusterRequest) {
    request = &CreateClusterRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tdmq", APIVersion, "CreateCluster")
    return
}

func NewCreateClusterResponse() (response *CreateClusterResponse) {
    response = &CreateClusterResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建用户的集群
func (c *Client) CreateCluster(request *CreateClusterRequest) (response *CreateClusterResponse, err error) {
    if request == nil {
        request = NewCreateClusterRequest()
    }
    response = NewCreateClusterResponse()
    err = c.Send(request, response)
    return
}

func NewCreateEnvironmentRequest() (request *CreateEnvironmentRequest) {
    request = &CreateEnvironmentRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tdmq", APIVersion, "CreateEnvironment")
    return
}

func NewCreateEnvironmentResponse() (response *CreateEnvironmentResponse) {
    response = &CreateEnvironmentResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 用于在用户账户下创建消息队列 Tdmq环境（命名空间）
func (c *Client) CreateEnvironment(request *CreateEnvironmentRequest) (response *CreateEnvironmentResponse, err error) {
    if request == nil {
        request = NewCreateEnvironmentRequest()
    }
    response = NewCreateEnvironmentResponse()
    err = c.Send(request, response)
    return
}

func NewCreateSubscriptionRequest() (request *CreateSubscriptionRequest) {
    request = &CreateSubscriptionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tdmq", APIVersion, "CreateSubscription")
    return
}

func NewCreateSubscriptionResponse() (response *CreateSubscriptionResponse) {
    response = &CreateSubscriptionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建一个主题的订阅关系
func (c *Client) CreateSubscription(request *CreateSubscriptionRequest) (response *CreateSubscriptionResponse, err error) {
    if request == nil {
        request = NewCreateSubscriptionRequest()
    }
    response = NewCreateSubscriptionResponse()
    err = c.Send(request, response)
    return
}

func NewCreateTopicRequest() (request *CreateTopicRequest) {
    request = &CreateTopicRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tdmq", APIVersion, "CreateTopic")
    return
}

func NewCreateTopicResponse() (response *CreateTopicResponse) {
    response = &CreateTopicResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 新增指定分区、类型的消息主题
func (c *Client) CreateTopic(request *CreateTopicRequest) (response *CreateTopicResponse, err error) {
    if request == nil {
        request = NewCreateTopicRequest()
    }
    response = NewCreateTopicResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteClusterRequest() (request *DeleteClusterRequest) {
    request = &DeleteClusterRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tdmq", APIVersion, "DeleteCluster")
    return
}

func NewDeleteClusterResponse() (response *DeleteClusterResponse) {
    response = &DeleteClusterResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除集群
func (c *Client) DeleteCluster(request *DeleteClusterRequest) (response *DeleteClusterResponse, err error) {
    if request == nil {
        request = NewDeleteClusterRequest()
    }
    response = NewDeleteClusterResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteEnvironmentsRequest() (request *DeleteEnvironmentsRequest) {
    request = &DeleteEnvironmentsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tdmq", APIVersion, "DeleteEnvironments")
    return
}

func NewDeleteEnvironmentsResponse() (response *DeleteEnvironmentsResponse) {
    response = &DeleteEnvironmentsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 批量删除租户下的环境
func (c *Client) DeleteEnvironments(request *DeleteEnvironmentsRequest) (response *DeleteEnvironmentsResponse, err error) {
    if request == nil {
        request = NewDeleteEnvironmentsRequest()
    }
    response = NewDeleteEnvironmentsResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteSubscriptionsRequest() (request *DeleteSubscriptionsRequest) {
    request = &DeleteSubscriptionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tdmq", APIVersion, "DeleteSubscriptions")
    return
}

func NewDeleteSubscriptionsResponse() (response *DeleteSubscriptionsResponse) {
    response = &DeleteSubscriptionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除订阅关系
func (c *Client) DeleteSubscriptions(request *DeleteSubscriptionsRequest) (response *DeleteSubscriptionsResponse, err error) {
    if request == nil {
        request = NewDeleteSubscriptionsRequest()
    }
    response = NewDeleteSubscriptionsResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteTopicsRequest() (request *DeleteTopicsRequest) {
    request = &DeleteTopicsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tdmq", APIVersion, "DeleteTopics")
    return
}

func NewDeleteTopicsResponse() (response *DeleteTopicsResponse) {
    response = &DeleteTopicsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 批量删除topics
func (c *Client) DeleteTopics(request *DeleteTopicsRequest) (response *DeleteTopicsResponse, err error) {
    if request == nil {
        request = NewDeleteTopicsRequest()
    }
    response = NewDeleteTopicsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBindClustersRequest() (request *DescribeBindClustersRequest) {
    request = &DescribeBindClustersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tdmq", APIVersion, "DescribeBindClusters")
    return
}

func NewDescribeBindClustersResponse() (response *DescribeBindClustersResponse) {
    response = &DescribeBindClustersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取用户绑定的专享集群列表
func (c *Client) DescribeBindClusters(request *DescribeBindClustersRequest) (response *DescribeBindClustersResponse, err error) {
    if request == nil {
        request = NewDescribeBindClustersRequest()
    }
    response = NewDescribeBindClustersResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBindVpcsRequest() (request *DescribeBindVpcsRequest) {
    request = &DescribeBindVpcsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tdmq", APIVersion, "DescribeBindVpcs")
    return
}

func NewDescribeBindVpcsResponse() (response *DescribeBindVpcsResponse) {
    response = &DescribeBindVpcsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取租户VPC绑定关系
func (c *Client) DescribeBindVpcs(request *DescribeBindVpcsRequest) (response *DescribeBindVpcsResponse, err error) {
    if request == nil {
        request = NewDescribeBindVpcsRequest()
    }
    response = NewDescribeBindVpcsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeClusterDetailRequest() (request *DescribeClusterDetailRequest) {
    request = &DescribeClusterDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tdmq", APIVersion, "DescribeClusterDetail")
    return
}

func NewDescribeClusterDetailResponse() (response *DescribeClusterDetailResponse) {
    response = &DescribeClusterDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取集群的详细信息
func (c *Client) DescribeClusterDetail(request *DescribeClusterDetailRequest) (response *DescribeClusterDetailResponse, err error) {
    if request == nil {
        request = NewDescribeClusterDetailRequest()
    }
    response = NewDescribeClusterDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeClustersRequest() (request *DescribeClustersRequest) {
    request = &DescribeClustersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tdmq", APIVersion, "DescribeClusters")
    return
}

func NewDescribeClustersResponse() (response *DescribeClustersResponse) {
    response = &DescribeClustersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取集群列表
func (c *Client) DescribeClusters(request *DescribeClustersRequest) (response *DescribeClustersResponse, err error) {
    if request == nil {
        request = NewDescribeClustersRequest()
    }
    response = NewDescribeClustersResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEnvironmentAttributesRequest() (request *DescribeEnvironmentAttributesRequest) {
    request = &DescribeEnvironmentAttributesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tdmq", APIVersion, "DescribeEnvironmentAttributes")
    return
}

func NewDescribeEnvironmentAttributesResponse() (response *DescribeEnvironmentAttributesResponse) {
    response = &DescribeEnvironmentAttributesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取指定环境的属性
func (c *Client) DescribeEnvironmentAttributes(request *DescribeEnvironmentAttributesRequest) (response *DescribeEnvironmentAttributesResponse, err error) {
    if request == nil {
        request = NewDescribeEnvironmentAttributesRequest()
    }
    response = NewDescribeEnvironmentAttributesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEnvironmentRolesRequest() (request *DescribeEnvironmentRolesRequest) {
    request = &DescribeEnvironmentRolesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tdmq", APIVersion, "DescribeEnvironmentRoles")
    return
}

func NewDescribeEnvironmentRolesResponse() (response *DescribeEnvironmentRolesResponse) {
    response = &DescribeEnvironmentRolesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取环境角色列表
func (c *Client) DescribeEnvironmentRoles(request *DescribeEnvironmentRolesRequest) (response *DescribeEnvironmentRolesResponse, err error) {
    if request == nil {
        request = NewDescribeEnvironmentRolesRequest()
    }
    response = NewDescribeEnvironmentRolesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEnvironmentsRequest() (request *DescribeEnvironmentsRequest) {
    request = &DescribeEnvironmentsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tdmq", APIVersion, "DescribeEnvironments")
    return
}

func NewDescribeEnvironmentsResponse() (response *DescribeEnvironmentsResponse) {
    response = &DescribeEnvironmentsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取租户下环境列表
func (c *Client) DescribeEnvironments(request *DescribeEnvironmentsRequest) (response *DescribeEnvironmentsResponse, err error) {
    if request == nil {
        request = NewDescribeEnvironmentsRequest()
    }
    response = NewDescribeEnvironmentsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeProducersRequest() (request *DescribeProducersRequest) {
    request = &DescribeProducersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tdmq", APIVersion, "DescribeProducers")
    return
}

func NewDescribeProducersResponse() (response *DescribeProducersResponse) {
    response = &DescribeProducersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取生产者列表，仅显示在线的生产者
func (c *Client) DescribeProducers(request *DescribeProducersRequest) (response *DescribeProducersResponse, err error) {
    if request == nil {
        request = NewDescribeProducersRequest()
    }
    response = NewDescribeProducersResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSubscriptionsRequest() (request *DescribeSubscriptionsRequest) {
    request = &DescribeSubscriptionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tdmq", APIVersion, "DescribeSubscriptions")
    return
}

func NewDescribeSubscriptionsResponse() (response *DescribeSubscriptionsResponse) {
    response = &DescribeSubscriptionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询指定环境和主题下的订阅者列表
func (c *Client) DescribeSubscriptions(request *DescribeSubscriptionsRequest) (response *DescribeSubscriptionsResponse, err error) {
    if request == nil {
        request = NewDescribeSubscriptionsRequest()
    }
    response = NewDescribeSubscriptionsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTopicsRequest() (request *DescribeTopicsRequest) {
    request = &DescribeTopicsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tdmq", APIVersion, "DescribeTopics")
    return
}

func NewDescribeTopicsResponse() (response *DescribeTopicsResponse) {
    response = &DescribeTopicsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取环境下主题列表
func (c *Client) DescribeTopics(request *DescribeTopicsRequest) (response *DescribeTopicsResponse, err error) {
    if request == nil {
        request = NewDescribeTopicsRequest()
    }
    response = NewDescribeTopicsResponse()
    err = c.Send(request, response)
    return
}

func NewModifyClusterRequest() (request *ModifyClusterRequest) {
    request = &ModifyClusterRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tdmq", APIVersion, "ModifyCluster")
    return
}

func NewModifyClusterResponse() (response *ModifyClusterResponse) {
    response = &ModifyClusterResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 更新集群信息
func (c *Client) ModifyCluster(request *ModifyClusterRequest) (response *ModifyClusterResponse, err error) {
    if request == nil {
        request = NewModifyClusterRequest()
    }
    response = NewModifyClusterResponse()
    err = c.Send(request, response)
    return
}

func NewModifyEnvironmentAttributesRequest() (request *ModifyEnvironmentAttributesRequest) {
    request = &ModifyEnvironmentAttributesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tdmq", APIVersion, "ModifyEnvironmentAttributes")
    return
}

func NewModifyEnvironmentAttributesResponse() (response *ModifyEnvironmentAttributesResponse) {
    response = &ModifyEnvironmentAttributesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改指定环境的属性值
func (c *Client) ModifyEnvironmentAttributes(request *ModifyEnvironmentAttributesRequest) (response *ModifyEnvironmentAttributesResponse, err error) {
    if request == nil {
        request = NewModifyEnvironmentAttributesRequest()
    }
    response = NewModifyEnvironmentAttributesResponse()
    err = c.Send(request, response)
    return
}

func NewModifyTopicRequest() (request *ModifyTopicRequest) {
    request = &ModifyTopicRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tdmq", APIVersion, "ModifyTopic")
    return
}

func NewModifyTopicResponse() (response *ModifyTopicResponse) {
    response = &ModifyTopicResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改主题备注和分区数
func (c *Client) ModifyTopic(request *ModifyTopicRequest) (response *ModifyTopicResponse, err error) {
    if request == nil {
        request = NewModifyTopicRequest()
    }
    response = NewModifyTopicResponse()
    err = c.Send(request, response)
    return
}

func NewResetMsgSubOffsetByTimestampRequest() (request *ResetMsgSubOffsetByTimestampRequest) {
    request = &ResetMsgSubOffsetByTimestampRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tdmq", APIVersion, "ResetMsgSubOffsetByTimestamp")
    return
}

func NewResetMsgSubOffsetByTimestampResponse() (response *ResetMsgSubOffsetByTimestampResponse) {
    response = &ResetMsgSubOffsetByTimestampResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 根据时间戳进行消息回溯，精确到毫秒
func (c *Client) ResetMsgSubOffsetByTimestamp(request *ResetMsgSubOffsetByTimestampRequest) (response *ResetMsgSubOffsetByTimestampResponse, err error) {
    if request == nil {
        request = NewResetMsgSubOffsetByTimestampRequest()
    }
    response = NewResetMsgSubOffsetByTimestampResponse()
    err = c.Send(request, response)
    return
}
