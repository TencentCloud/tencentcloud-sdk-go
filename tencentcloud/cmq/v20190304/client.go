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
    "context"
    "errors"
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

func NewClient(credential common.CredentialIface, region string, clientProfile *profile.ClientProfile) (client *Client, err error) {
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

// ClearQueue
// 清除queue中的所有消息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TRYLATER = "FailedOperation.TryLater"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_NOTASKID = "InvalidParameterValue.NoTaskId"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ClearQueue(request *ClearQueueRequest) (response *ClearQueueResponse, err error) {
    return c.ClearQueueWithContext(context.Background(), request)
}

// ClearQueue
// 清除queue中的所有消息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TRYLATER = "FailedOperation.TryLater"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_NOTASKID = "InvalidParameterValue.NoTaskId"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ClearQueueWithContext(ctx context.Context, request *ClearQueueRequest) (response *ClearQueueResponse, err error) {
    if request == nil {
        request = NewClearQueueRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ClearQueue require credential")
    }

    request.SetContext(ctx)
    
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

// ClearSubscriptionFilterTags
// 清空订阅者消息标签
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TRYLATER = "FailedOperation.TryLater"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_NOTASKID = "InvalidParameterValue.NoTaskId"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ClearSubscriptionFilterTags(request *ClearSubscriptionFilterTagsRequest) (response *ClearSubscriptionFilterTagsResponse, err error) {
    return c.ClearSubscriptionFilterTagsWithContext(context.Background(), request)
}

// ClearSubscriptionFilterTags
// 清空订阅者消息标签
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TRYLATER = "FailedOperation.TryLater"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_NOTASKID = "InvalidParameterValue.NoTaskId"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ClearSubscriptionFilterTagsWithContext(ctx context.Context, request *ClearSubscriptionFilterTagsRequest) (response *ClearSubscriptionFilterTagsResponse, err error) {
    if request == nil {
        request = NewClearSubscriptionFilterTagsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ClearSubscriptionFilterTags require credential")
    }

    request.SetContext(ctx)
    
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

// CreateQueue
// 创建队列接口
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TRYLATER = "FailedOperation.TryLater"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateQueue(request *CreateQueueRequest) (response *CreateQueueResponse, err error) {
    return c.CreateQueueWithContext(context.Background(), request)
}

// CreateQueue
// 创建队列接口
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TRYLATER = "FailedOperation.TryLater"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateQueueWithContext(ctx context.Context, request *CreateQueueRequest) (response *CreateQueueResponse, err error) {
    if request == nil {
        request = NewCreateQueueRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateQueue require credential")
    }

    request.SetContext(ctx)
    
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

// CreateSubscribe
// 创建订阅接口
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TRYLATER = "FailedOperation.TryLater"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_NOTASKID = "InvalidParameterValue.NoTaskId"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateSubscribe(request *CreateSubscribeRequest) (response *CreateSubscribeResponse, err error) {
    return c.CreateSubscribeWithContext(context.Background(), request)
}

// CreateSubscribe
// 创建订阅接口
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TRYLATER = "FailedOperation.TryLater"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_NOTASKID = "InvalidParameterValue.NoTaskId"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateSubscribeWithContext(ctx context.Context, request *CreateSubscribeRequest) (response *CreateSubscribeResponse, err error) {
    if request == nil {
        request = NewCreateSubscribeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateSubscribe require credential")
    }

    request.SetContext(ctx)
    
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

// CreateTopic
// 创建主题
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TRYLATER = "FailedOperation.TryLater"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateTopic(request *CreateTopicRequest) (response *CreateTopicResponse, err error) {
    return c.CreateTopicWithContext(context.Background(), request)
}

// CreateTopic
// 创建主题
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TRYLATER = "FailedOperation.TryLater"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
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
// DeleteQueue
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeleteQueue(request *DeleteQueueRequest) (response *DeleteQueueResponse, err error) {
    return c.DeleteQueueWithContext(context.Background(), request)
}

// DeleteQueue
// DeleteQueue
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeleteQueueWithContext(ctx context.Context, request *DeleteQueueRequest) (response *DeleteQueueResponse, err error) {
    if request == nil {
        request = NewDeleteQueueRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteQueue require credential")
    }

    request.SetContext(ctx)
    
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

// DeleteSubscribe
// 删除订阅
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TRYLATER = "FailedOperation.TryLater"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeleteSubscribe(request *DeleteSubscribeRequest) (response *DeleteSubscribeResponse, err error) {
    return c.DeleteSubscribeWithContext(context.Background(), request)
}

// DeleteSubscribe
// 删除订阅
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TRYLATER = "FailedOperation.TryLater"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeleteSubscribeWithContext(ctx context.Context, request *DeleteSubscribeRequest) (response *DeleteSubscribeResponse, err error) {
    if request == nil {
        request = NewDeleteSubscribeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteSubscribe require credential")
    }

    request.SetContext(ctx)
    
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

// DeleteTopic
// 删除主题
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TRYLATER = "FailedOperation.TryLater"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeleteTopic(request *DeleteTopicRequest) (response *DeleteTopicResponse, err error) {
    return c.DeleteTopicWithContext(context.Background(), request)
}

// DeleteTopic
// 删除主题
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TRYLATER = "FailedOperation.TryLater"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
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

// DescribeDeadLetterSourceQueues
// 枚举死信队列源队列
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TRYLATER = "FailedOperation.TryLater"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeDeadLetterSourceQueues(request *DescribeDeadLetterSourceQueuesRequest) (response *DescribeDeadLetterSourceQueuesResponse, err error) {
    return c.DescribeDeadLetterSourceQueuesWithContext(context.Background(), request)
}

// DescribeDeadLetterSourceQueues
// 枚举死信队列源队列
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TRYLATER = "FailedOperation.TryLater"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeDeadLetterSourceQueuesWithContext(ctx context.Context, request *DescribeDeadLetterSourceQueuesRequest) (response *DescribeDeadLetterSourceQueuesResponse, err error) {
    if request == nil {
        request = NewDescribeDeadLetterSourceQueuesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDeadLetterSourceQueues require credential")
    }

    request.SetContext(ctx)
    
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

// DescribeQueueDetail
// 枚举队列
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TRYLATER = "FailedOperation.TryLater"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeQueueDetail(request *DescribeQueueDetailRequest) (response *DescribeQueueDetailResponse, err error) {
    return c.DescribeQueueDetailWithContext(context.Background(), request)
}

// DescribeQueueDetail
// 枚举队列
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TRYLATER = "FailedOperation.TryLater"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeQueueDetailWithContext(ctx context.Context, request *DescribeQueueDetailRequest) (response *DescribeQueueDetailResponse, err error) {
    if request == nil {
        request = NewDescribeQueueDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeQueueDetail require credential")
    }

    request.SetContext(ctx)
    
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

// DescribeSubscriptionDetail
// 查询订阅详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TRYLATER = "FailedOperation.TryLater"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_NOTASKID = "InvalidParameterValue.NoTaskId"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeSubscriptionDetail(request *DescribeSubscriptionDetailRequest) (response *DescribeSubscriptionDetailResponse, err error) {
    return c.DescribeSubscriptionDetailWithContext(context.Background(), request)
}

// DescribeSubscriptionDetail
// 查询订阅详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TRYLATER = "FailedOperation.TryLater"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_NOTASKID = "InvalidParameterValue.NoTaskId"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeSubscriptionDetailWithContext(ctx context.Context, request *DescribeSubscriptionDetailRequest) (response *DescribeSubscriptionDetailResponse, err error) {
    if request == nil {
        request = NewDescribeSubscriptionDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSubscriptionDetail require credential")
    }

    request.SetContext(ctx)
    
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

// DescribeTopicDetail
// 查询主题详情 
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TRYLATER = "FailedOperation.TryLater"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_NOTASKID = "InvalidParameterValue.NoTaskId"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeTopicDetail(request *DescribeTopicDetailRequest) (response *DescribeTopicDetailResponse, err error) {
    return c.DescribeTopicDetailWithContext(context.Background(), request)
}

// DescribeTopicDetail
// 查询主题详情 
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TRYLATER = "FailedOperation.TryLater"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_NOTASKID = "InvalidParameterValue.NoTaskId"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeTopicDetailWithContext(ctx context.Context, request *DescribeTopicDetailRequest) (response *DescribeTopicDetailResponse, err error) {
    if request == nil {
        request = NewDescribeTopicDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTopicDetail require credential")
    }

    request.SetContext(ctx)
    
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

// ModifyQueueAttribute
// 修改队列属性
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TRYLATER = "FailedOperation.TryLater"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_NOTASKID = "InvalidParameterValue.NoTaskId"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyQueueAttribute(request *ModifyQueueAttributeRequest) (response *ModifyQueueAttributeResponse, err error) {
    return c.ModifyQueueAttributeWithContext(context.Background(), request)
}

// ModifyQueueAttribute
// 修改队列属性
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TRYLATER = "FailedOperation.TryLater"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_NOTASKID = "InvalidParameterValue.NoTaskId"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyQueueAttributeWithContext(ctx context.Context, request *ModifyQueueAttributeRequest) (response *ModifyQueueAttributeResponse, err error) {
    if request == nil {
        request = NewModifyQueueAttributeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyQueueAttribute require credential")
    }

    request.SetContext(ctx)
    
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

// ModifySubscriptionAttribute
// 修改订阅属性
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TRYLATER = "FailedOperation.TryLater"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_NOTASKID = "InvalidParameterValue.NoTaskId"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifySubscriptionAttribute(request *ModifySubscriptionAttributeRequest) (response *ModifySubscriptionAttributeResponse, err error) {
    return c.ModifySubscriptionAttributeWithContext(context.Background(), request)
}

// ModifySubscriptionAttribute
// 修改订阅属性
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TRYLATER = "FailedOperation.TryLater"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_NOTASKID = "InvalidParameterValue.NoTaskId"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifySubscriptionAttributeWithContext(ctx context.Context, request *ModifySubscriptionAttributeRequest) (response *ModifySubscriptionAttributeResponse, err error) {
    if request == nil {
        request = NewModifySubscriptionAttributeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifySubscriptionAttribute require credential")
    }

    request.SetContext(ctx)
    
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

// ModifyTopicAttribute
// 修改主题属性
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TRYLATER = "FailedOperation.TryLater"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyTopicAttribute(request *ModifyTopicAttributeRequest) (response *ModifyTopicAttributeResponse, err error) {
    return c.ModifyTopicAttributeWithContext(context.Background(), request)
}

// ModifyTopicAttribute
// 修改主题属性
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TRYLATER = "FailedOperation.TryLater"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyTopicAttributeWithContext(ctx context.Context, request *ModifyTopicAttributeRequest) (response *ModifyTopicAttributeResponse, err error) {
    if request == nil {
        request = NewModifyTopicAttributeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyTopicAttribute require credential")
    }

    request.SetContext(ctx)
    
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

// RewindQueue
// 回溯队列
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TRYLATER = "FailedOperation.TryLater"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) RewindQueue(request *RewindQueueRequest) (response *RewindQueueResponse, err error) {
    return c.RewindQueueWithContext(context.Background(), request)
}

// RewindQueue
// 回溯队列
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TRYLATER = "FailedOperation.TryLater"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) RewindQueueWithContext(ctx context.Context, request *RewindQueueRequest) (response *RewindQueueResponse, err error) {
    if request == nil {
        request = NewRewindQueueRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RewindQueue require credential")
    }

    request.SetContext(ctx)
    
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

// UnbindDeadLetter
// 解绑死信队列
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TRYLATER = "FailedOperation.TryLater"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_NOTASKID = "InvalidParameterValue.NoTaskId"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) UnbindDeadLetter(request *UnbindDeadLetterRequest) (response *UnbindDeadLetterResponse, err error) {
    return c.UnbindDeadLetterWithContext(context.Background(), request)
}

// UnbindDeadLetter
// 解绑死信队列
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TRYLATER = "FailedOperation.TryLater"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_NOTASKID = "InvalidParameterValue.NoTaskId"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) UnbindDeadLetterWithContext(ctx context.Context, request *UnbindDeadLetterRequest) (response *UnbindDeadLetterResponse, err error) {
    if request == nil {
        request = NewUnbindDeadLetterRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UnbindDeadLetter require credential")
    }

    request.SetContext(ctx)
    
    response = NewUnbindDeadLetterResponse()
    err = c.Send(request, response)
    return
}
