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

package v20230308

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2023-03-08"

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


func NewCreateConsumerGroupRequest() (request *CreateConsumerGroupRequest) {
    request = &CreateConsumerGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trocket", APIVersion, "CreateConsumerGroup")
    
    
    return
}

func NewCreateConsumerGroupResponse() (response *CreateConsumerGroupResponse) {
    response = &CreateConsumerGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateConsumerGroup
// 创建消费组
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INSTANCENOTREADY = "FailedOperation.InstanceNotReady"
//  UNSUPPORTEDOPERATION_RESOURCEALREADYEXISTS = "UnsupportedOperation.ResourceAlreadyExists"
func (c *Client) CreateConsumerGroup(request *CreateConsumerGroupRequest) (response *CreateConsumerGroupResponse, err error) {
    return c.CreateConsumerGroupWithContext(context.Background(), request)
}

// CreateConsumerGroup
// 创建消费组
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INSTANCENOTREADY = "FailedOperation.InstanceNotReady"
//  UNSUPPORTEDOPERATION_RESOURCEALREADYEXISTS = "UnsupportedOperation.ResourceAlreadyExists"
func (c *Client) CreateConsumerGroupWithContext(ctx context.Context, request *CreateConsumerGroupRequest) (response *CreateConsumerGroupResponse, err error) {
    if request == nil {
        request = NewCreateConsumerGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateConsumerGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateConsumerGroupResponse()
    err = c.Send(request, response)
    return
}

func NewCreateInstanceRequest() (request *CreateInstanceRequest) {
    request = &CreateInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trocket", APIVersion, "CreateInstance")
    
    
    return
}

func NewCreateInstanceResponse() (response *CreateInstanceResponse) {
    response = &CreateInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateInstance
// 购买新实例
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CreateInstance(request *CreateInstanceRequest) (response *CreateInstanceResponse, err error) {
    return c.CreateInstanceWithContext(context.Background(), request)
}

// CreateInstance
// 购买新实例
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CreateInstanceWithContext(ctx context.Context, request *CreateInstanceRequest) (response *CreateInstanceResponse, err error) {
    if request == nil {
        request = NewCreateInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewCreateTopicRequest() (request *CreateTopicRequest) {
    request = &CreateTopicRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trocket", APIVersion, "CreateTopic")
    
    
    return
}

func NewCreateTopicResponse() (response *CreateTopicResponse) {
    response = &CreateTopicResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateTopic
// 创建主题
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INSTANCENOTREADY = "FailedOperation.InstanceNotReady"
//  LIMITEXCEEDED_TOPICNUM = "LimitExceeded.TopicNum"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
//  UNSUPPORTEDOPERATION_RESOURCEALREADYEXISTS = "UnsupportedOperation.ResourceAlreadyExists"
func (c *Client) CreateTopic(request *CreateTopicRequest) (response *CreateTopicResponse, err error) {
    return c.CreateTopicWithContext(context.Background(), request)
}

// CreateTopic
// 创建主题
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INSTANCENOTREADY = "FailedOperation.InstanceNotReady"
//  LIMITEXCEEDED_TOPICNUM = "LimitExceeded.TopicNum"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
//  UNSUPPORTEDOPERATION_RESOURCEALREADYEXISTS = "UnsupportedOperation.ResourceAlreadyExists"
func (c *Client) CreateTopicWithContext(ctx context.Context, request *CreateTopicRequest) (response *CreateTopicResponse, err error) {
    if request == nil {
        request = NewCreateTopicRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateTopic require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateTopicResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteConsumerGroupRequest() (request *DeleteConsumerGroupRequest) {
    request = &DeleteConsumerGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trocket", APIVersion, "DeleteConsumerGroup")
    
    
    return
}

func NewDeleteConsumerGroupResponse() (response *DeleteConsumerGroupResponse) {
    response = &DeleteConsumerGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteConsumerGroup
// 删除消费组
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCENOTFOUND_GROUP = "ResourceNotFound.Group"
func (c *Client) DeleteConsumerGroup(request *DeleteConsumerGroupRequest) (response *DeleteConsumerGroupResponse, err error) {
    return c.DeleteConsumerGroupWithContext(context.Background(), request)
}

// DeleteConsumerGroup
// 删除消费组
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCENOTFOUND_GROUP = "ResourceNotFound.Group"
func (c *Client) DeleteConsumerGroupWithContext(ctx context.Context, request *DeleteConsumerGroupRequest) (response *DeleteConsumerGroupResponse, err error) {
    if request == nil {
        request = NewDeleteConsumerGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteConsumerGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteConsumerGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteInstanceRequest() (request *DeleteInstanceRequest) {
    request = &DeleteInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trocket", APIVersion, "DeleteInstance")
    
    
    return
}

func NewDeleteInstanceResponse() (response *DeleteInstanceResponse) {
    response = &DeleteInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteInstance
// 删除实例
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
func (c *Client) DeleteInstance(request *DeleteInstanceRequest) (response *DeleteInstanceResponse, err error) {
    return c.DeleteInstanceWithContext(context.Background(), request)
}

// DeleteInstance
// 删除实例
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
func (c *Client) DeleteInstanceWithContext(ctx context.Context, request *DeleteInstanceRequest) (response *DeleteInstanceResponse, err error) {
    if request == nil {
        request = NewDeleteInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteTopicRequest() (request *DeleteTopicRequest) {
    request = &DeleteTopicRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trocket", APIVersion, "DeleteTopic")
    
    
    return
}

func NewDeleteTopicResponse() (response *DeleteTopicResponse) {
    response = &DeleteTopicResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteTopic
// 删除主题
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DeleteTopic(request *DeleteTopicRequest) (response *DeleteTopicResponse, err error) {
    return c.DeleteTopicWithContext(context.Background(), request)
}

// DeleteTopic
// 删除主题
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DeleteTopicWithContext(ctx context.Context, request *DeleteTopicRequest) (response *DeleteTopicResponse, err error) {
    if request == nil {
        request = NewDeleteTopicRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteTopic require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteTopicResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeConsumerGroupRequest() (request *DescribeConsumerGroupRequest) {
    request = &DescribeConsumerGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trocket", APIVersion, "DescribeConsumerGroup")
    
    
    return
}

func NewDescribeConsumerGroupResponse() (response *DescribeConsumerGroupResponse) {
    response = &DescribeConsumerGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeConsumerGroup
// 查询消费组详情
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_GROUP = "ResourceNotFound.Group"
func (c *Client) DescribeConsumerGroup(request *DescribeConsumerGroupRequest) (response *DescribeConsumerGroupResponse, err error) {
    return c.DescribeConsumerGroupWithContext(context.Background(), request)
}

// DescribeConsumerGroup
// 查询消费组详情
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_GROUP = "ResourceNotFound.Group"
func (c *Client) DescribeConsumerGroupWithContext(ctx context.Context, request *DescribeConsumerGroupRequest) (response *DescribeConsumerGroupResponse, err error) {
    if request == nil {
        request = NewDescribeConsumerGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeConsumerGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeConsumerGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceRequest() (request *DescribeInstanceRequest) {
    request = &DescribeInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trocket", APIVersion, "DescribeInstance")
    
    
    return
}

func NewDescribeInstanceResponse() (response *DescribeInstanceResponse) {
    response = &DescribeInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeInstance
// 查询实例信息
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
func (c *Client) DescribeInstance(request *DescribeInstanceRequest) (response *DescribeInstanceResponse, err error) {
    return c.DescribeInstanceWithContext(context.Background(), request)
}

// DescribeInstance
// 查询实例信息
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
func (c *Client) DescribeInstanceWithContext(ctx context.Context, request *DescribeInstanceRequest) (response *DescribeInstanceResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceListRequest() (request *DescribeInstanceListRequest) {
    request = &DescribeInstanceListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trocket", APIVersion, "DescribeInstanceList")
    
    
    return
}

func NewDescribeInstanceListResponse() (response *DescribeInstanceListResponse) {
    response = &DescribeInstanceListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeInstanceList
// 获取实例列表，Filters参数使用说明如下：
//
// 1. InstanceName, 名称模糊查询
//
// 2. InstanceId，实例ID查询
//
// 3. InstanceType, 实例类型查询，支持多选
//
// 3. InstanceStatus，实例状态查询，支持多选
//
// 
//
// 当使用TagFilters查询时，Filters参数失效。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeInstanceList(request *DescribeInstanceListRequest) (response *DescribeInstanceListResponse, err error) {
    return c.DescribeInstanceListWithContext(context.Background(), request)
}

// DescribeInstanceList
// 获取实例列表，Filters参数使用说明如下：
//
// 1. InstanceName, 名称模糊查询
//
// 2. InstanceId，实例ID查询
//
// 3. InstanceType, 实例类型查询，支持多选
//
// 3. InstanceStatus，实例状态查询，支持多选
//
// 
//
// 当使用TagFilters查询时，Filters参数失效。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeInstanceListWithContext(ctx context.Context, request *DescribeInstanceListRequest) (response *DescribeInstanceListResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstanceList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstanceListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTopicRequest() (request *DescribeTopicRequest) {
    request = &DescribeTopicRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trocket", APIVersion, "DescribeTopic")
    
    
    return
}

func NewDescribeTopicResponse() (response *DescribeTopicResponse) {
    response = &DescribeTopicResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeTopic
// 查询主题详情，Offset和Limit参数是指订阅该主题的消费组查询分页参数，Filter参数使用说明如下：
//
// 
//
// ConsumerGroup，消费组名称过滤
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
//  RESOURCENOTFOUND_TOPIC = "ResourceNotFound.Topic"
func (c *Client) DescribeTopic(request *DescribeTopicRequest) (response *DescribeTopicResponse, err error) {
    return c.DescribeTopicWithContext(context.Background(), request)
}

// DescribeTopic
// 查询主题详情，Offset和Limit参数是指订阅该主题的消费组查询分页参数，Filter参数使用说明如下：
//
// 
//
// ConsumerGroup，消费组名称过滤
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
//  RESOURCENOTFOUND_TOPIC = "ResourceNotFound.Topic"
func (c *Client) DescribeTopicWithContext(ctx context.Context, request *DescribeTopicRequest) (response *DescribeTopicResponse, err error) {
    if request == nil {
        request = NewDescribeTopicRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTopic require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTopicResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTopicListRequest() (request *DescribeTopicListRequest) {
    request = &DescribeTopicListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trocket", APIVersion, "DescribeTopicList")
    
    
    return
}

func NewDescribeTopicListResponse() (response *DescribeTopicListResponse) {
    response = &DescribeTopicListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeTopicList
// 获取主题列表，Filter参数使用说明如下：
//
// 
//
// 1. TopicName，主题名称模糊搜索
//
// 2. TopicType，主题类型查询，支持多选，可选值：Normal,Order,Transaction,DelayScheduled
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
func (c *Client) DescribeTopicList(request *DescribeTopicListRequest) (response *DescribeTopicListResponse, err error) {
    return c.DescribeTopicListWithContext(context.Background(), request)
}

// DescribeTopicList
// 获取主题列表，Filter参数使用说明如下：
//
// 
//
// 1. TopicName，主题名称模糊搜索
//
// 2. TopicType，主题类型查询，支持多选，可选值：Normal,Order,Transaction,DelayScheduled
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
func (c *Client) DescribeTopicListWithContext(ctx context.Context, request *DescribeTopicListRequest) (response *DescribeTopicListResponse, err error) {
    if request == nil {
        request = NewDescribeTopicListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTopicList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTopicListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTopicStatsOpRequest() (request *DescribeTopicStatsOpRequest) {
    request = &DescribeTopicStatsOpRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trocket", APIVersion, "DescribeTopicStatsOp")
    
    
    return
}

func NewDescribeTopicStatsOpResponse() (response *DescribeTopicStatsOpResponse) {
    response = &DescribeTopicStatsOpResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeTopicStatsOp
// 运营端查询topicStata
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
func (c *Client) DescribeTopicStatsOp(request *DescribeTopicStatsOpRequest) (response *DescribeTopicStatsOpResponse, err error) {
    return c.DescribeTopicStatsOpWithContext(context.Background(), request)
}

// DescribeTopicStatsOp
// 运营端查询topicStata
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
func (c *Client) DescribeTopicStatsOpWithContext(ctx context.Context, request *DescribeTopicStatsOpRequest) (response *DescribeTopicStatsOpResponse, err error) {
    if request == nil {
        request = NewDescribeTopicStatsOpRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTopicStatsOp require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTopicStatsOpResponse()
    err = c.Send(request, response)
    return
}

func NewModifyConsumerGroupRequest() (request *ModifyConsumerGroupRequest) {
    request = &ModifyConsumerGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trocket", APIVersion, "ModifyConsumerGroup")
    
    
    return
}

func NewModifyConsumerGroupResponse() (response *ModifyConsumerGroupResponse) {
    response = &ModifyConsumerGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyConsumerGroup
// 修改消费组属性
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
func (c *Client) ModifyConsumerGroup(request *ModifyConsumerGroupRequest) (response *ModifyConsumerGroupResponse, err error) {
    return c.ModifyConsumerGroupWithContext(context.Background(), request)
}

// ModifyConsumerGroup
// 修改消费组属性
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
func (c *Client) ModifyConsumerGroupWithContext(ctx context.Context, request *ModifyConsumerGroupRequest) (response *ModifyConsumerGroupResponse, err error) {
    if request == nil {
        request = NewModifyConsumerGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyConsumerGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyConsumerGroupResponse()
    err = c.Send(request, response)
    return
}

func NewModifyInstanceRequest() (request *ModifyInstanceRequest) {
    request = &ModifyInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trocket", APIVersion, "ModifyInstance")
    
    
    return
}

func NewModifyInstanceResponse() (response *ModifyInstanceResponse) {
    response = &ModifyInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyInstance
// 修改实例属性
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INSTANCENOTREADY = "FailedOperation.InstanceNotReady"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) ModifyInstance(request *ModifyInstanceRequest) (response *ModifyInstanceResponse, err error) {
    return c.ModifyInstanceWithContext(context.Background(), request)
}

// ModifyInstance
// 修改实例属性
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INSTANCENOTREADY = "FailedOperation.InstanceNotReady"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) ModifyInstanceWithContext(ctx context.Context, request *ModifyInstanceRequest) (response *ModifyInstanceResponse, err error) {
    if request == nil {
        request = NewModifyInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewModifyTopicRequest() (request *ModifyTopicRequest) {
    request = &ModifyTopicRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trocket", APIVersion, "ModifyTopic")
    
    
    return
}

func NewModifyTopicResponse() (response *ModifyTopicResponse) {
    response = &ModifyTopicResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyTopic
// 修改主题属性
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ModifyTopic(request *ModifyTopicRequest) (response *ModifyTopicResponse, err error) {
    return c.ModifyTopicWithContext(context.Background(), request)
}

// ModifyTopic
// 修改主题属性
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ModifyTopicWithContext(ctx context.Context, request *ModifyTopicRequest) (response *ModifyTopicResponse, err error) {
    if request == nil {
        request = NewModifyTopicRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyTopic require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyTopicResponse()
    err = c.Send(request, response)
    return
}
