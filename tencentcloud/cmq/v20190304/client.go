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

package v20190304

import (
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2019-03-04"

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


func NewClearQueueRequest() (request *ClearQueueRequest) {
    request = &ClearQueueRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cmq", APIVersion, "ClearQueue")
    return
}

func NewClearQueueResponse() (response *ClearQueueResponse) {
    response = &ClearQueueResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 清除queue中的所有消息
func (c *Client) ClearQueue(request *ClearQueueRequest) (response *ClearQueueResponse, err error) {
    if request == nil {
        request = NewClearQueueRequest()
    }
    response = NewClearQueueResponse()
    err = c.Send(request, response)
    return
}

func NewClearSubscriptionFilterTagsRequest() (request *ClearSubscriptionFilterTagsRequest) {
    request = &ClearSubscriptionFilterTagsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cmq", APIVersion, "ClearSubscriptionFilterTags")
    return
}

func NewClearSubscriptionFilterTagsResponse() (response *ClearSubscriptionFilterTagsResponse) {
    response = &ClearSubscriptionFilterTagsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 清空订阅者消息标签
func (c *Client) ClearSubscriptionFilterTags(request *ClearSubscriptionFilterTagsRequest) (response *ClearSubscriptionFilterTagsResponse, err error) {
    if request == nil {
        request = NewClearSubscriptionFilterTagsRequest()
    }
    response = NewClearSubscriptionFilterTagsResponse()
    err = c.Send(request, response)
    return
}

func NewCreateQueueRequest() (request *CreateQueueRequest) {
    request = &CreateQueueRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cmq", APIVersion, "CreateQueue")
    return
}

func NewCreateQueueResponse() (response *CreateQueueResponse) {
    response = &CreateQueueResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建队列接口
func (c *Client) CreateQueue(request *CreateQueueRequest) (response *CreateQueueResponse, err error) {
    if request == nil {
        request = NewCreateQueueRequest()
    }
    response = NewCreateQueueResponse()
    err = c.Send(request, response)
    return
}

func NewCreateSubscribeRequest() (request *CreateSubscribeRequest) {
    request = &CreateSubscribeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cmq", APIVersion, "CreateSubscribe")
    return
}

func NewCreateSubscribeResponse() (response *CreateSubscribeResponse) {
    response = &CreateSubscribeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建订阅接口
func (c *Client) CreateSubscribe(request *CreateSubscribeRequest) (response *CreateSubscribeResponse, err error) {
    if request == nil {
        request = NewCreateSubscribeRequest()
    }
    response = NewCreateSubscribeResponse()
    err = c.Send(request, response)
    return
}

func NewCreateTopicRequest() (request *CreateTopicRequest) {
    request = &CreateTopicRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cmq", APIVersion, "CreateTopic")
    return
}

func NewCreateTopicResponse() (response *CreateTopicResponse) {
    response = &CreateTopicResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建主题
func (c *Client) CreateTopic(request *CreateTopicRequest) (response *CreateTopicResponse, err error) {
    if request == nil {
        request = NewCreateTopicRequest()
    }
    response = NewCreateTopicResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteQueueRequest() (request *DeleteQueueRequest) {
    request = &DeleteQueueRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cmq", APIVersion, "DeleteQueue")
    return
}

func NewDeleteQueueResponse() (response *DeleteQueueResponse) {
    response = &DeleteQueueResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteQueue
func (c *Client) DeleteQueue(request *DeleteQueueRequest) (response *DeleteQueueResponse, err error) {
    if request == nil {
        request = NewDeleteQueueRequest()
    }
    response = NewDeleteQueueResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteSubscribeRequest() (request *DeleteSubscribeRequest) {
    request = &DeleteSubscribeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cmq", APIVersion, "DeleteSubscribe")
    return
}

func NewDeleteSubscribeResponse() (response *DeleteSubscribeResponse) {
    response = &DeleteSubscribeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除订阅
func (c *Client) DeleteSubscribe(request *DeleteSubscribeRequest) (response *DeleteSubscribeResponse, err error) {
    if request == nil {
        request = NewDeleteSubscribeRequest()
    }
    response = NewDeleteSubscribeResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteTopicRequest() (request *DeleteTopicRequest) {
    request = &DeleteTopicRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cmq", APIVersion, "DeleteTopic")
    return
}

func NewDeleteTopicResponse() (response *DeleteTopicResponse) {
    response = &DeleteTopicResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除主题
func (c *Client) DeleteTopic(request *DeleteTopicRequest) (response *DeleteTopicResponse, err error) {
    if request == nil {
        request = NewDeleteTopicRequest()
    }
    response = NewDeleteTopicResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDeadLetterSourceQueuesRequest() (request *DescribeDeadLetterSourceQueuesRequest) {
    request = &DescribeDeadLetterSourceQueuesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cmq", APIVersion, "DescribeDeadLetterSourceQueues")
    return
}

func NewDescribeDeadLetterSourceQueuesResponse() (response *DescribeDeadLetterSourceQueuesResponse) {
    response = &DescribeDeadLetterSourceQueuesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 枚举死信队列源队列
func (c *Client) DescribeDeadLetterSourceQueues(request *DescribeDeadLetterSourceQueuesRequest) (response *DescribeDeadLetterSourceQueuesResponse, err error) {
    if request == nil {
        request = NewDescribeDeadLetterSourceQueuesRequest()
    }
    response = NewDescribeDeadLetterSourceQueuesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeQueueDetailRequest() (request *DescribeQueueDetailRequest) {
    request = &DescribeQueueDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cmq", APIVersion, "DescribeQueueDetail")
    return
}

func NewDescribeQueueDetailResponse() (response *DescribeQueueDetailResponse) {
    response = &DescribeQueueDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 枚举队列
func (c *Client) DescribeQueueDetail(request *DescribeQueueDetailRequest) (response *DescribeQueueDetailResponse, err error) {
    if request == nil {
        request = NewDescribeQueueDetailRequest()
    }
    response = NewDescribeQueueDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSubscriptionDetailRequest() (request *DescribeSubscriptionDetailRequest) {
    request = &DescribeSubscriptionDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cmq", APIVersion, "DescribeSubscriptionDetail")
    return
}

func NewDescribeSubscriptionDetailResponse() (response *DescribeSubscriptionDetailResponse) {
    response = &DescribeSubscriptionDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询订阅详情
func (c *Client) DescribeSubscriptionDetail(request *DescribeSubscriptionDetailRequest) (response *DescribeSubscriptionDetailResponse, err error) {
    if request == nil {
        request = NewDescribeSubscriptionDetailRequest()
    }
    response = NewDescribeSubscriptionDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTopicDetailRequest() (request *DescribeTopicDetailRequest) {
    request = &DescribeTopicDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cmq", APIVersion, "DescribeTopicDetail")
    return
}

func NewDescribeTopicDetailResponse() (response *DescribeTopicDetailResponse) {
    response = &DescribeTopicDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询主题详情
func (c *Client) DescribeTopicDetail(request *DescribeTopicDetailRequest) (response *DescribeTopicDetailResponse, err error) {
    if request == nil {
        request = NewDescribeTopicDetailRequest()
    }
    response = NewDescribeTopicDetailResponse()
    err = c.Send(request, response)
    return
}

func NewModifyQueueAttributeRequest() (request *ModifyQueueAttributeRequest) {
    request = &ModifyQueueAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cmq", APIVersion, "ModifyQueueAttribute")
    return
}

func NewModifyQueueAttributeResponse() (response *ModifyQueueAttributeResponse) {
    response = &ModifyQueueAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改队列属性
func (c *Client) ModifyQueueAttribute(request *ModifyQueueAttributeRequest) (response *ModifyQueueAttributeResponse, err error) {
    if request == nil {
        request = NewModifyQueueAttributeRequest()
    }
    response = NewModifyQueueAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewModifySubscriptionAttributeRequest() (request *ModifySubscriptionAttributeRequest) {
    request = &ModifySubscriptionAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cmq", APIVersion, "ModifySubscriptionAttribute")
    return
}

func NewModifySubscriptionAttributeResponse() (response *ModifySubscriptionAttributeResponse) {
    response = &ModifySubscriptionAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改订阅属性
func (c *Client) ModifySubscriptionAttribute(request *ModifySubscriptionAttributeRequest) (response *ModifySubscriptionAttributeResponse, err error) {
    if request == nil {
        request = NewModifySubscriptionAttributeRequest()
    }
    response = NewModifySubscriptionAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyTopicAttributeRequest() (request *ModifyTopicAttributeRequest) {
    request = &ModifyTopicAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cmq", APIVersion, "ModifyTopicAttribute")
    return
}

func NewModifyTopicAttributeResponse() (response *ModifyTopicAttributeResponse) {
    response = &ModifyTopicAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改主题属性
func (c *Client) ModifyTopicAttribute(request *ModifyTopicAttributeRequest) (response *ModifyTopicAttributeResponse, err error) {
    if request == nil {
        request = NewModifyTopicAttributeRequest()
    }
    response = NewModifyTopicAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewRewindQueueRequest() (request *RewindQueueRequest) {
    request = &RewindQueueRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cmq", APIVersion, "RewindQueue")
    return
}

func NewRewindQueueResponse() (response *RewindQueueResponse) {
    response = &RewindQueueResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 回溯队列
func (c *Client) RewindQueue(request *RewindQueueRequest) (response *RewindQueueResponse, err error) {
    if request == nil {
        request = NewRewindQueueRequest()
    }
    response = NewRewindQueueResponse()
    err = c.Send(request, response)
    return
}

func NewUnbindDeadLetterRequest() (request *UnbindDeadLetterRequest) {
    request = &UnbindDeadLetterRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cmq", APIVersion, "UnbindDeadLetter")
    return
}

func NewUnbindDeadLetterResponse() (response *UnbindDeadLetterResponse) {
    response = &UnbindDeadLetterResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 解绑死信队列
func (c *Client) UnbindDeadLetter(request *UnbindDeadLetterRequest) (response *UnbindDeadLetterResponse, err error) {
    if request == nil {
        request = NewUnbindDeadLetterRequest()
    }
    response = NewUnbindDeadLetterResponse()
    err = c.Send(request, response)
    return
}
