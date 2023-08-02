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
    "context"
    "errors"
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

func NewClient(credential common.CredentialIface, region string, clientProfile *profile.ClientProfile) (client *Client, err error) {
    client = &Client{}
    client.Init(region).
        WithCredential(credential).
        WithProfile(clientProfile)
    return
}


func NewAcknowledgeMessageRequest() (request *AcknowledgeMessageRequest) {
    request = &AcknowledgeMessageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "AcknowledgeMessage")
    
    
    return
}

func NewAcknowledgeMessageResponse() (response *AcknowledgeMessageResponse) {
    response = &AcknowledgeMessageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AcknowledgeMessage
// 根据提供的 MessageID 确认指定 topic 中的消息
//
// 可能返回的错误码:
//  FAILEDOPERATION_CREATEPRODUCERERROR = "FailedOperation.CreateProducerError"
//  FAILEDOPERATION_CREATEPULSARCLIENTERROR = "FailedOperation.CreatePulsarClientError"
//  FAILEDOPERATION_MAXMESSAGESIZEERROR = "FailedOperation.MaxMessageSizeError"
//  FAILEDOPERATION_MESSAGEIDERROR = "FailedOperation.MessageIDError"
//  FAILEDOPERATION_RECEIVEERROR = "FailedOperation.ReceiveError"
//  FAILEDOPERATION_RECEIVETIMEOUT = "FailedOperation.ReceiveTimeout"
//  FAILEDOPERATION_TOPICTYPEERROR = "FailedOperation.TopicTypeError"
//  INVALIDPARAMETER_TENANTNOTFOUND = "InvalidParameter.TenantNotFound"
//  INVALIDPARAMETER_TOKENNOTFOUND = "InvalidParameter.TokenNotFound"
//  INVALIDPARAMETERVALUE_TOPICNOTFOUND = "InvalidParameterValue.TopicNotFound"
func (c *Client) AcknowledgeMessage(request *AcknowledgeMessageRequest) (response *AcknowledgeMessageResponse, err error) {
    return c.AcknowledgeMessageWithContext(context.Background(), request)
}

// AcknowledgeMessage
// 根据提供的 MessageID 确认指定 topic 中的消息
//
// 可能返回的错误码:
//  FAILEDOPERATION_CREATEPRODUCERERROR = "FailedOperation.CreateProducerError"
//  FAILEDOPERATION_CREATEPULSARCLIENTERROR = "FailedOperation.CreatePulsarClientError"
//  FAILEDOPERATION_MAXMESSAGESIZEERROR = "FailedOperation.MaxMessageSizeError"
//  FAILEDOPERATION_MESSAGEIDERROR = "FailedOperation.MessageIDError"
//  FAILEDOPERATION_RECEIVEERROR = "FailedOperation.ReceiveError"
//  FAILEDOPERATION_RECEIVETIMEOUT = "FailedOperation.ReceiveTimeout"
//  FAILEDOPERATION_TOPICTYPEERROR = "FailedOperation.TopicTypeError"
//  INVALIDPARAMETER_TENANTNOTFOUND = "InvalidParameter.TenantNotFound"
//  INVALIDPARAMETER_TOKENNOTFOUND = "InvalidParameter.TokenNotFound"
//  INVALIDPARAMETERVALUE_TOPICNOTFOUND = "InvalidParameterValue.TopicNotFound"
func (c *Client) AcknowledgeMessageWithContext(ctx context.Context, request *AcknowledgeMessageRequest) (response *AcknowledgeMessageResponse, err error) {
    if request == nil {
        request = NewAcknowledgeMessageRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AcknowledgeMessage require credential")
    }

    request.SetContext(ctx)
    
    response = NewAcknowledgeMessageResponse()
    err = c.Send(request, response)
    return
}

func NewClearCmqQueueRequest() (request *ClearCmqQueueRequest) {
    request = &ClearCmqQueueRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "ClearCmqQueue")
    
    
    return
}

func NewClearCmqQueueResponse() (response *ClearCmqQueueResponse) {
    response = &ClearCmqQueueResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ClearCmqQueue
// 清空cmq消息队列中的消息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) ClearCmqQueue(request *ClearCmqQueueRequest) (response *ClearCmqQueueResponse, err error) {
    return c.ClearCmqQueueWithContext(context.Background(), request)
}

// ClearCmqQueue
// 清空cmq消息队列中的消息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) ClearCmqQueueWithContext(ctx context.Context, request *ClearCmqQueueRequest) (response *ClearCmqQueueResponse, err error) {
    if request == nil {
        request = NewClearCmqQueueRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ClearCmqQueue require credential")
    }

    request.SetContext(ctx)
    
    response = NewClearCmqQueueResponse()
    err = c.Send(request, response)
    return
}

func NewClearCmqSubscriptionFilterTagsRequest() (request *ClearCmqSubscriptionFilterTagsRequest) {
    request = &ClearCmqSubscriptionFilterTagsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "ClearCmqSubscriptionFilterTags")
    
    
    return
}

func NewClearCmqSubscriptionFilterTagsResponse() (response *ClearCmqSubscriptionFilterTagsResponse) {
    response = &ClearCmqSubscriptionFilterTagsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ClearCmqSubscriptionFilterTags
// 清空订阅者消息标签
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ClearCmqSubscriptionFilterTags(request *ClearCmqSubscriptionFilterTagsRequest) (response *ClearCmqSubscriptionFilterTagsResponse, err error) {
    return c.ClearCmqSubscriptionFilterTagsWithContext(context.Background(), request)
}

// ClearCmqSubscriptionFilterTags
// 清空订阅者消息标签
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ClearCmqSubscriptionFilterTagsWithContext(ctx context.Context, request *ClearCmqSubscriptionFilterTagsRequest) (response *ClearCmqSubscriptionFilterTagsResponse, err error) {
    if request == nil {
        request = NewClearCmqSubscriptionFilterTagsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ClearCmqSubscriptionFilterTags require credential")
    }

    request.SetContext(ctx)
    
    response = NewClearCmqSubscriptionFilterTagsResponse()
    err = c.Send(request, response)
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

// CreateCluster
// 创建用户的集群
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CREATECLUSTER = "FailedOperation.CreateCluster"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  LIMITEXCEEDED_CLUSTERS = "LimitExceeded.Clusters"
//  MISSINGPARAMETER_NEEDMOREPARAMS = "MissingParameter.NeedMoreParams"
//  RESOURCEINUSE_CLUSTER = "ResourceInUse.Cluster"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_CREATEFAILED = "ResourceUnavailable.CreateFailed"
//  RESOURCEUNAVAILABLE_FUNDREQUIRED = "ResourceUnavailable.FundRequired"
func (c *Client) CreateCluster(request *CreateClusterRequest) (response *CreateClusterResponse, err error) {
    return c.CreateClusterWithContext(context.Background(), request)
}

// CreateCluster
// 创建用户的集群
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CREATECLUSTER = "FailedOperation.CreateCluster"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  LIMITEXCEEDED_CLUSTERS = "LimitExceeded.Clusters"
//  MISSINGPARAMETER_NEEDMOREPARAMS = "MissingParameter.NeedMoreParams"
//  RESOURCEINUSE_CLUSTER = "ResourceInUse.Cluster"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_CREATEFAILED = "ResourceUnavailable.CreateFailed"
//  RESOURCEUNAVAILABLE_FUNDREQUIRED = "ResourceUnavailable.FundRequired"
func (c *Client) CreateClusterWithContext(ctx context.Context, request *CreateClusterRequest) (response *CreateClusterResponse, err error) {
    if request == nil {
        request = NewCreateClusterRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCluster require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateClusterResponse()
    err = c.Send(request, response)
    return
}

func NewCreateCmqQueueRequest() (request *CreateCmqQueueRequest) {
    request = &CreateCmqQueueRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "CreateCmqQueue")
    
    
    return
}

func NewCreateCmqQueueResponse() (response *CreateCmqQueueResponse) {
    response = &CreateCmqQueueResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateCmqQueue
// 创建cmq队列接口
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CREATECLUSTER = "FailedOperation.CreateCluster"
//  FAILEDOPERATION_CREATETOPIC = "FailedOperation.CreateTopic"
//  FAILEDOPERATION_SETRETENTIONPOLICY = "FailedOperation.SetRetentionPolicy"
//  FAILEDOPERATION_SETTTL = "FailedOperation.SetTTL"
//  INTERNALERROR_BROKERSERVICE = "InternalError.BrokerService"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINUSE_QUEUE = "ResourceInUse.Queue"
//  RESOURCENOTFOUND_ENVIRONMENT = "ResourceNotFound.Environment"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_CREATEFAILED = "ResourceUnavailable.CreateFailed"
//  RESOURCEUNAVAILABLE_FUNDREQUIRED = "ResourceUnavailable.FundRequired"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateCmqQueue(request *CreateCmqQueueRequest) (response *CreateCmqQueueResponse, err error) {
    return c.CreateCmqQueueWithContext(context.Background(), request)
}

// CreateCmqQueue
// 创建cmq队列接口
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CREATECLUSTER = "FailedOperation.CreateCluster"
//  FAILEDOPERATION_CREATETOPIC = "FailedOperation.CreateTopic"
//  FAILEDOPERATION_SETRETENTIONPOLICY = "FailedOperation.SetRetentionPolicy"
//  FAILEDOPERATION_SETTTL = "FailedOperation.SetTTL"
//  INTERNALERROR_BROKERSERVICE = "InternalError.BrokerService"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINUSE_QUEUE = "ResourceInUse.Queue"
//  RESOURCENOTFOUND_ENVIRONMENT = "ResourceNotFound.Environment"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_CREATEFAILED = "ResourceUnavailable.CreateFailed"
//  RESOURCEUNAVAILABLE_FUNDREQUIRED = "ResourceUnavailable.FundRequired"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateCmqQueueWithContext(ctx context.Context, request *CreateCmqQueueRequest) (response *CreateCmqQueueResponse, err error) {
    if request == nil {
        request = NewCreateCmqQueueRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCmqQueue require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateCmqQueueResponse()
    err = c.Send(request, response)
    return
}

func NewCreateCmqSubscribeRequest() (request *CreateCmqSubscribeRequest) {
    request = &CreateCmqSubscribeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "CreateCmqSubscribe")
    
    
    return
}

func NewCreateCmqSubscribeResponse() (response *CreateCmqSubscribeResponse) {
    response = &CreateCmqSubscribeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateCmqSubscribe
// 创建cmq订阅接口
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CREATESUBSCRIPTION = "FailedOperation.CreateSubscription"
//  FAILEDOPERATION_CREATETOPIC = "FailedOperation.CreateTopic"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINUSE_SUBSCRIPTION = "ResourceInUse.Subscription"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_FUNDREQUIRED = "ResourceUnavailable.FundRequired"
func (c *Client) CreateCmqSubscribe(request *CreateCmqSubscribeRequest) (response *CreateCmqSubscribeResponse, err error) {
    return c.CreateCmqSubscribeWithContext(context.Background(), request)
}

// CreateCmqSubscribe
// 创建cmq订阅接口
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CREATESUBSCRIPTION = "FailedOperation.CreateSubscription"
//  FAILEDOPERATION_CREATETOPIC = "FailedOperation.CreateTopic"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINUSE_SUBSCRIPTION = "ResourceInUse.Subscription"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_FUNDREQUIRED = "ResourceUnavailable.FundRequired"
func (c *Client) CreateCmqSubscribeWithContext(ctx context.Context, request *CreateCmqSubscribeRequest) (response *CreateCmqSubscribeResponse, err error) {
    if request == nil {
        request = NewCreateCmqSubscribeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCmqSubscribe require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateCmqSubscribeResponse()
    err = c.Send(request, response)
    return
}

func NewCreateCmqTopicRequest() (request *CreateCmqTopicRequest) {
    request = &CreateCmqTopicRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "CreateCmqTopic")
    
    
    return
}

func NewCreateCmqTopicResponse() (response *CreateCmqTopicResponse) {
    response = &CreateCmqTopicResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateCmqTopic
// 创建cmq主题
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CREATECLUSTER = "FailedOperation.CreateCluster"
//  FAILEDOPERATION_CREATETOPIC = "FailedOperation.CreateTopic"
//  FAILEDOPERATION_SETRETENTIONPOLICY = "FailedOperation.SetRetentionPolicy"
//  FAILEDOPERATION_SETTTL = "FailedOperation.SetTTL"
//  INTERNALERROR_BROKERSERVICE = "InternalError.BrokerService"
//  LIMITEXCEEDED_TOPICS = "LimitExceeded.Topics"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINUSE_TOPIC = "ResourceInUse.Topic"
//  RESOURCENOTFOUND_ENVIRONMENT = "ResourceNotFound.Environment"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_CREATEFAILED = "ResourceUnavailable.CreateFailed"
//  RESOURCEUNAVAILABLE_FUNDREQUIRED = "ResourceUnavailable.FundRequired"
func (c *Client) CreateCmqTopic(request *CreateCmqTopicRequest) (response *CreateCmqTopicResponse, err error) {
    return c.CreateCmqTopicWithContext(context.Background(), request)
}

// CreateCmqTopic
// 创建cmq主题
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CREATECLUSTER = "FailedOperation.CreateCluster"
//  FAILEDOPERATION_CREATETOPIC = "FailedOperation.CreateTopic"
//  FAILEDOPERATION_SETRETENTIONPOLICY = "FailedOperation.SetRetentionPolicy"
//  FAILEDOPERATION_SETTTL = "FailedOperation.SetTTL"
//  INTERNALERROR_BROKERSERVICE = "InternalError.BrokerService"
//  LIMITEXCEEDED_TOPICS = "LimitExceeded.Topics"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINUSE_TOPIC = "ResourceInUse.Topic"
//  RESOURCENOTFOUND_ENVIRONMENT = "ResourceNotFound.Environment"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_CREATEFAILED = "ResourceUnavailable.CreateFailed"
//  RESOURCEUNAVAILABLE_FUNDREQUIRED = "ResourceUnavailable.FundRequired"
func (c *Client) CreateCmqTopicWithContext(ctx context.Context, request *CreateCmqTopicRequest) (response *CreateCmqTopicResponse, err error) {
    if request == nil {
        request = NewCreateCmqTopicRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCmqTopic require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateCmqTopicResponse()
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

// CreateEnvironment
// 用于在用户账户下创建消息队列 Tdmq 命名空间
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CREATEENVIRONMENT = "FailedOperation.CreateEnvironment"
//  FAILEDOPERATION_CREATENAMESPACE = "FailedOperation.CreateNamespace"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMS = "InvalidParameterValue.InvalidParams"
//  INVALIDPARAMETERVALUE_TTL = "InvalidParameterValue.TTL"
//  LIMITEXCEEDED_ENVIRONMENTS = "LimitExceeded.Environments"
//  LIMITEXCEEDED_NAMESPACES = "LimitExceeded.Namespaces"
//  LIMITEXCEEDED_RETENTIONSIZE = "LimitExceeded.RetentionSize"
//  LIMITEXCEEDED_RETENTIONTIME = "LimitExceeded.RetentionTime"
//  MISSINGPARAMETER_NEEDMOREPARAMS = "MissingParameter.NeedMoreParams"
//  OPERATIONDENIED_DEFAULTENVIRONMENT = "OperationDenied.DefaultEnvironment"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINUSE_NAMESPACE = "ResourceInUse.Namespace"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
func (c *Client) CreateEnvironment(request *CreateEnvironmentRequest) (response *CreateEnvironmentResponse, err error) {
    return c.CreateEnvironmentWithContext(context.Background(), request)
}

// CreateEnvironment
// 用于在用户账户下创建消息队列 Tdmq 命名空间
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CREATEENVIRONMENT = "FailedOperation.CreateEnvironment"
//  FAILEDOPERATION_CREATENAMESPACE = "FailedOperation.CreateNamespace"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMS = "InvalidParameterValue.InvalidParams"
//  INVALIDPARAMETERVALUE_TTL = "InvalidParameterValue.TTL"
//  LIMITEXCEEDED_ENVIRONMENTS = "LimitExceeded.Environments"
//  LIMITEXCEEDED_NAMESPACES = "LimitExceeded.Namespaces"
//  LIMITEXCEEDED_RETENTIONSIZE = "LimitExceeded.RetentionSize"
//  LIMITEXCEEDED_RETENTIONTIME = "LimitExceeded.RetentionTime"
//  MISSINGPARAMETER_NEEDMOREPARAMS = "MissingParameter.NeedMoreParams"
//  OPERATIONDENIED_DEFAULTENVIRONMENT = "OperationDenied.DefaultEnvironment"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINUSE_NAMESPACE = "ResourceInUse.Namespace"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
func (c *Client) CreateEnvironmentWithContext(ctx context.Context, request *CreateEnvironmentRequest) (response *CreateEnvironmentResponse, err error) {
    if request == nil {
        request = NewCreateEnvironmentRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateEnvironment require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateEnvironmentResponse()
    err = c.Send(request, response)
    return
}

func NewCreateEnvironmentRoleRequest() (request *CreateEnvironmentRoleRequest) {
    request = &CreateEnvironmentRoleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "CreateEnvironmentRole")
    
    
    return
}

func NewCreateEnvironmentRoleResponse() (response *CreateEnvironmentRoleResponse) {
    response = &CreateEnvironmentRoleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateEnvironmentRole
// 创建环境角色授权
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CREATEENVIRONMENTROLE = "FailedOperation.CreateEnvironmentRole"
//  FAILEDOPERATION_UPDATEENVIRONMENTROLE = "FailedOperation.UpdateEnvironmentRole"
//  INVALIDPARAMETERVALUE_INVALIDPARAMS = "InvalidParameterValue.InvalidParams"
//  MISSINGPARAMETER_NEEDMOREPARAMS = "MissingParameter.NeedMoreParams"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINUSE_ENVIRONMENTROLE = "ResourceInUse.EnvironmentRole"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
//  RESOURCENOTFOUND_ENVIRONMENT = "ResourceNotFound.Environment"
//  RESOURCENOTFOUND_ROLE = "ResourceNotFound.Role"
func (c *Client) CreateEnvironmentRole(request *CreateEnvironmentRoleRequest) (response *CreateEnvironmentRoleResponse, err error) {
    return c.CreateEnvironmentRoleWithContext(context.Background(), request)
}

// CreateEnvironmentRole
// 创建环境角色授权
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CREATEENVIRONMENTROLE = "FailedOperation.CreateEnvironmentRole"
//  FAILEDOPERATION_UPDATEENVIRONMENTROLE = "FailedOperation.UpdateEnvironmentRole"
//  INVALIDPARAMETERVALUE_INVALIDPARAMS = "InvalidParameterValue.InvalidParams"
//  MISSINGPARAMETER_NEEDMOREPARAMS = "MissingParameter.NeedMoreParams"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINUSE_ENVIRONMENTROLE = "ResourceInUse.EnvironmentRole"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
//  RESOURCENOTFOUND_ENVIRONMENT = "ResourceNotFound.Environment"
//  RESOURCENOTFOUND_ROLE = "ResourceNotFound.Role"
func (c *Client) CreateEnvironmentRoleWithContext(ctx context.Context, request *CreateEnvironmentRoleRequest) (response *CreateEnvironmentRoleResponse, err error) {
    if request == nil {
        request = NewCreateEnvironmentRoleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateEnvironmentRole require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateEnvironmentRoleResponse()
    err = c.Send(request, response)
    return
}

func NewCreateRabbitMQUserRequest() (request *CreateRabbitMQUserRequest) {
    request = &CreateRabbitMQUserRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "CreateRabbitMQUser")
    
    
    return
}

func NewCreateRabbitMQUserResponse() (response *CreateRabbitMQUserResponse) {
    response = &CreateRabbitMQUserResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateRabbitMQUser
// 创建RabbitMQ的用户
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CreateRabbitMQUser(request *CreateRabbitMQUserRequest) (response *CreateRabbitMQUserResponse, err error) {
    return c.CreateRabbitMQUserWithContext(context.Background(), request)
}

// CreateRabbitMQUser
// 创建RabbitMQ的用户
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CreateRabbitMQUserWithContext(ctx context.Context, request *CreateRabbitMQUserRequest) (response *CreateRabbitMQUserResponse, err error) {
    if request == nil {
        request = NewCreateRabbitMQUserRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateRabbitMQUser require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateRabbitMQUserResponse()
    err = c.Send(request, response)
    return
}

func NewCreateRabbitMQVipInstanceRequest() (request *CreateRabbitMQVipInstanceRequest) {
    request = &CreateRabbitMQVipInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "CreateRabbitMQVipInstance")
    
    
    return
}

func NewCreateRabbitMQVipInstanceResponse() (response *CreateRabbitMQVipInstanceResponse) {
    response = &CreateRabbitMQVipInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateRabbitMQVipInstance
// 创建RabbitMQ专享版实例
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CreateRabbitMQVipInstance(request *CreateRabbitMQVipInstanceRequest) (response *CreateRabbitMQVipInstanceResponse, err error) {
    return c.CreateRabbitMQVipInstanceWithContext(context.Background(), request)
}

// CreateRabbitMQVipInstance
// 创建RabbitMQ专享版实例
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CreateRabbitMQVipInstanceWithContext(ctx context.Context, request *CreateRabbitMQVipInstanceRequest) (response *CreateRabbitMQVipInstanceResponse, err error) {
    if request == nil {
        request = NewCreateRabbitMQVipInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateRabbitMQVipInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateRabbitMQVipInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewCreateRabbitMQVirtualHostRequest() (request *CreateRabbitMQVirtualHostRequest) {
    request = &CreateRabbitMQVirtualHostRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "CreateRabbitMQVirtualHost")
    
    
    return
}

func NewCreateRabbitMQVirtualHostResponse() (response *CreateRabbitMQVirtualHostResponse) {
    response = &CreateRabbitMQVirtualHostResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateRabbitMQVirtualHost
// 创建RabbitMQ的vhost
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CreateRabbitMQVirtualHost(request *CreateRabbitMQVirtualHostRequest) (response *CreateRabbitMQVirtualHostResponse, err error) {
    return c.CreateRabbitMQVirtualHostWithContext(context.Background(), request)
}

// CreateRabbitMQVirtualHost
// 创建RabbitMQ的vhost
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CreateRabbitMQVirtualHostWithContext(ctx context.Context, request *CreateRabbitMQVirtualHostRequest) (response *CreateRabbitMQVirtualHostResponse, err error) {
    if request == nil {
        request = NewCreateRabbitMQVirtualHostRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateRabbitMQVirtualHost require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateRabbitMQVirtualHostResponse()
    err = c.Send(request, response)
    return
}

func NewCreateRocketMQClusterRequest() (request *CreateRocketMQClusterRequest) {
    request = &CreateRocketMQClusterRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "CreateRocketMQCluster")
    
    
    return
}

func NewCreateRocketMQClusterResponse() (response *CreateRocketMQClusterResponse) {
    response = &CreateRocketMQClusterResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateRocketMQCluster
// 此接口用于创建一个RocketMQ集群
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CREATECLUSTER = "FailedOperation.CreateCluster"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_CLUSTERS = "LimitExceeded.Clusters"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_CREATEFAILED = "ResourceUnavailable.CreateFailed"
//  RESOURCEUNAVAILABLE_FUNDREQUIRED = "ResourceUnavailable.FundRequired"
func (c *Client) CreateRocketMQCluster(request *CreateRocketMQClusterRequest) (response *CreateRocketMQClusterResponse, err error) {
    return c.CreateRocketMQClusterWithContext(context.Background(), request)
}

// CreateRocketMQCluster
// 此接口用于创建一个RocketMQ集群
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CREATECLUSTER = "FailedOperation.CreateCluster"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_CLUSTERS = "LimitExceeded.Clusters"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_CREATEFAILED = "ResourceUnavailable.CreateFailed"
//  RESOURCEUNAVAILABLE_FUNDREQUIRED = "ResourceUnavailable.FundRequired"
func (c *Client) CreateRocketMQClusterWithContext(ctx context.Context, request *CreateRocketMQClusterRequest) (response *CreateRocketMQClusterResponse, err error) {
    if request == nil {
        request = NewCreateRocketMQClusterRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateRocketMQCluster require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateRocketMQClusterResponse()
    err = c.Send(request, response)
    return
}

func NewCreateRocketMQGroupRequest() (request *CreateRocketMQGroupRequest) {
    request = &CreateRocketMQGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "CreateRocketMQGroup")
    
    
    return
}

func NewCreateRocketMQGroupResponse() (response *CreateRocketMQGroupResponse) {
    response = &CreateRocketMQGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateRocketMQGroup
// 创建RocketMQ消费组
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ILLEGALMESSAGE = "InternalError.IllegalMessage"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
func (c *Client) CreateRocketMQGroup(request *CreateRocketMQGroupRequest) (response *CreateRocketMQGroupResponse, err error) {
    return c.CreateRocketMQGroupWithContext(context.Background(), request)
}

// CreateRocketMQGroup
// 创建RocketMQ消费组
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ILLEGALMESSAGE = "InternalError.IllegalMessage"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
func (c *Client) CreateRocketMQGroupWithContext(ctx context.Context, request *CreateRocketMQGroupRequest) (response *CreateRocketMQGroupResponse, err error) {
    if request == nil {
        request = NewCreateRocketMQGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateRocketMQGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateRocketMQGroupResponse()
    err = c.Send(request, response)
    return
}

func NewCreateRocketMQNamespaceRequest() (request *CreateRocketMQNamespaceRequest) {
    request = &CreateRocketMQNamespaceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "CreateRocketMQNamespace")
    
    
    return
}

func NewCreateRocketMQNamespaceResponse() (response *CreateRocketMQNamespaceResponse) {
    response = &CreateRocketMQNamespaceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateRocketMQNamespace
// 创建RocketMQ命名空间
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_SETTTL = "FailedOperation.SetTTL"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ILLEGALMESSAGE = "InternalError.IllegalMessage"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
//  RESOURCEUNAVAILABLE_FUNDREQUIRED = "ResourceUnavailable.FundRequired"
func (c *Client) CreateRocketMQNamespace(request *CreateRocketMQNamespaceRequest) (response *CreateRocketMQNamespaceResponse, err error) {
    return c.CreateRocketMQNamespaceWithContext(context.Background(), request)
}

// CreateRocketMQNamespace
// 创建RocketMQ命名空间
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_SETTTL = "FailedOperation.SetTTL"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ILLEGALMESSAGE = "InternalError.IllegalMessage"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
//  RESOURCEUNAVAILABLE_FUNDREQUIRED = "ResourceUnavailable.FundRequired"
func (c *Client) CreateRocketMQNamespaceWithContext(ctx context.Context, request *CreateRocketMQNamespaceRequest) (response *CreateRocketMQNamespaceResponse, err error) {
    if request == nil {
        request = NewCreateRocketMQNamespaceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateRocketMQNamespace require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateRocketMQNamespaceResponse()
    err = c.Send(request, response)
    return
}

func NewCreateRocketMQTopicRequest() (request *CreateRocketMQTopicRequest) {
    request = &CreateRocketMQTopicRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "CreateRocketMQTopic")
    
    
    return
}

func NewCreateRocketMQTopicResponse() (response *CreateRocketMQTopicResponse) {
    response = &CreateRocketMQTopicResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateRocketMQTopic
// 创建RocketMQ主题
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CREATETOPIC = "FailedOperation.CreateTopic"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ILLEGALMESSAGE = "InternalError.IllegalMessage"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARTITION = "InvalidParameter.Partition"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
//  RESOURCENOTFOUND_NAMESPACE = "ResourceNotFound.Namespace"
//  RESOURCEUNAVAILABLE_FUNDREQUIRED = "ResourceUnavailable.FundRequired"
func (c *Client) CreateRocketMQTopic(request *CreateRocketMQTopicRequest) (response *CreateRocketMQTopicResponse, err error) {
    return c.CreateRocketMQTopicWithContext(context.Background(), request)
}

// CreateRocketMQTopic
// 创建RocketMQ主题
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CREATETOPIC = "FailedOperation.CreateTopic"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ILLEGALMESSAGE = "InternalError.IllegalMessage"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARTITION = "InvalidParameter.Partition"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
//  RESOURCENOTFOUND_NAMESPACE = "ResourceNotFound.Namespace"
//  RESOURCEUNAVAILABLE_FUNDREQUIRED = "ResourceUnavailable.FundRequired"
func (c *Client) CreateRocketMQTopicWithContext(ctx context.Context, request *CreateRocketMQTopicRequest) (response *CreateRocketMQTopicResponse, err error) {
    if request == nil {
        request = NewCreateRocketMQTopicRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateRocketMQTopic require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateRocketMQTopicResponse()
    err = c.Send(request, response)
    return
}

func NewCreateRocketMQVipInstanceRequest() (request *CreateRocketMQVipInstanceRequest) {
    request = &CreateRocketMQVipInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "CreateRocketMQVipInstance")
    
    
    return
}

func NewCreateRocketMQVipInstanceResponse() (response *CreateRocketMQVipInstanceResponse) {
    response = &CreateRocketMQVipInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateRocketMQVipInstance
// 创建RocketMQ专享实例
//
// 可能返回的错误码:
//  FAILEDOPERATION_CALLTRADE = "FailedOperation.CallTrade"
func (c *Client) CreateRocketMQVipInstance(request *CreateRocketMQVipInstanceRequest) (response *CreateRocketMQVipInstanceResponse, err error) {
    return c.CreateRocketMQVipInstanceWithContext(context.Background(), request)
}

// CreateRocketMQVipInstance
// 创建RocketMQ专享实例
//
// 可能返回的错误码:
//  FAILEDOPERATION_CALLTRADE = "FailedOperation.CallTrade"
func (c *Client) CreateRocketMQVipInstanceWithContext(ctx context.Context, request *CreateRocketMQVipInstanceRequest) (response *CreateRocketMQVipInstanceResponse, err error) {
    if request == nil {
        request = NewCreateRocketMQVipInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateRocketMQVipInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateRocketMQVipInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewCreateRoleRequest() (request *CreateRoleRequest) {
    request = &CreateRoleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "CreateRole")
    
    
    return
}

func NewCreateRoleResponse() (response *CreateRoleResponse) {
    response = &CreateRoleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateRole
// 创建角色
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CREATEROLE = "FailedOperation.CreateRole"
//  FAILEDOPERATION_CREATESECRETKEY = "FailedOperation.CreateSecretKey"
//  FAILEDOPERATION_SAVESECRETKEY = "FailedOperation.SaveSecretKey"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMS = "InvalidParameterValue.InvalidParams"
//  MISSINGPARAMETER_NEEDMOREPARAMS = "MissingParameter.NeedMoreParams"
//  RESOURCEINUSE_ROLE = "ResourceInUse.Role"
//  RESOURCENOTFOUND_BROKERCLUSTER = "ResourceNotFound.BrokerCluster"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
//  RESOURCEUNAVAILABLE_FUNDREQUIRED = "ResourceUnavailable.FundRequired"
func (c *Client) CreateRole(request *CreateRoleRequest) (response *CreateRoleResponse, err error) {
    return c.CreateRoleWithContext(context.Background(), request)
}

// CreateRole
// 创建角色
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CREATEROLE = "FailedOperation.CreateRole"
//  FAILEDOPERATION_CREATESECRETKEY = "FailedOperation.CreateSecretKey"
//  FAILEDOPERATION_SAVESECRETKEY = "FailedOperation.SaveSecretKey"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMS = "InvalidParameterValue.InvalidParams"
//  MISSINGPARAMETER_NEEDMOREPARAMS = "MissingParameter.NeedMoreParams"
//  RESOURCEINUSE_ROLE = "ResourceInUse.Role"
//  RESOURCENOTFOUND_BROKERCLUSTER = "ResourceNotFound.BrokerCluster"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
//  RESOURCEUNAVAILABLE_FUNDREQUIRED = "ResourceUnavailable.FundRequired"
func (c *Client) CreateRoleWithContext(ctx context.Context, request *CreateRoleRequest) (response *CreateRoleResponse, err error) {
    if request == nil {
        request = NewCreateRoleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateRole require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateRoleResponse()
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

// CreateSubscription
// 创建一个主题的订阅关系
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CREATESUBSCRIPTION = "FailedOperation.CreateSubscription"
//  FAILEDOPERATION_GETTOPICPARTITIONSFAILED = "FailedOperation.GetTopicPartitionsFailed"
//  INTERNALERROR_RETRY = "InternalError.Retry"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMS = "InvalidParameterValue.InvalidParams"
//  LIMITEXCEEDED_SUBSCRIPTIONS = "LimitExceeded.Subscriptions"
//  LIMITEXCEEDED_TOPICS = "LimitExceeded.Topics"
//  RESOURCEINUSE_SUBSCRIPTION = "ResourceInUse.Subscription"
//  RESOURCENOTFOUND_BROKERCLUSTER = "ResourceNotFound.BrokerCluster"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
//  RESOURCENOTFOUND_TOPIC = "ResourceNotFound.Topic"
//  RESOURCEUNAVAILABLE_FUNDREQUIRED = "ResourceUnavailable.FundRequired"
func (c *Client) CreateSubscription(request *CreateSubscriptionRequest) (response *CreateSubscriptionResponse, err error) {
    return c.CreateSubscriptionWithContext(context.Background(), request)
}

// CreateSubscription
// 创建一个主题的订阅关系
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CREATESUBSCRIPTION = "FailedOperation.CreateSubscription"
//  FAILEDOPERATION_GETTOPICPARTITIONSFAILED = "FailedOperation.GetTopicPartitionsFailed"
//  INTERNALERROR_RETRY = "InternalError.Retry"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMS = "InvalidParameterValue.InvalidParams"
//  LIMITEXCEEDED_SUBSCRIPTIONS = "LimitExceeded.Subscriptions"
//  LIMITEXCEEDED_TOPICS = "LimitExceeded.Topics"
//  RESOURCEINUSE_SUBSCRIPTION = "ResourceInUse.Subscription"
//  RESOURCENOTFOUND_BROKERCLUSTER = "ResourceNotFound.BrokerCluster"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
//  RESOURCENOTFOUND_TOPIC = "ResourceNotFound.Topic"
//  RESOURCEUNAVAILABLE_FUNDREQUIRED = "ResourceUnavailable.FundRequired"
func (c *Client) CreateSubscriptionWithContext(ctx context.Context, request *CreateSubscriptionRequest) (response *CreateSubscriptionResponse, err error) {
    if request == nil {
        request = NewCreateSubscriptionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateSubscription require credential")
    }

    request.SetContext(ctx)
    
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

// CreateTopic
// 新增指定分区、类型的消息主题
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CREATETOPIC = "FailedOperation.CreateTopic"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPARAMS = "InvalidParameterValue.InvalidParams"
//  LIMITEXCEEDED_TOPICS = "LimitExceeded.Topics"
//  MISSINGPARAMETER_NEEDMOREPARAMS = "MissingParameter.NeedMoreParams"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINUSE_TOPIC = "ResourceInUse.Topic"
//  RESOURCENOTFOUND_BROKERCLUSTER = "ResourceNotFound.BrokerCluster"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
//  RESOURCENOTFOUND_ENVIRONMENT = "ResourceNotFound.Environment"
//  RESOURCEUNAVAILABLE_FUNDREQUIRED = "ResourceUnavailable.FundRequired"
func (c *Client) CreateTopic(request *CreateTopicRequest) (response *CreateTopicResponse, err error) {
    return c.CreateTopicWithContext(context.Background(), request)
}

// CreateTopic
// 新增指定分区、类型的消息主题
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CREATETOPIC = "FailedOperation.CreateTopic"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPARAMS = "InvalidParameterValue.InvalidParams"
//  LIMITEXCEEDED_TOPICS = "LimitExceeded.Topics"
//  MISSINGPARAMETER_NEEDMOREPARAMS = "MissingParameter.NeedMoreParams"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINUSE_TOPIC = "ResourceInUse.Topic"
//  RESOURCENOTFOUND_BROKERCLUSTER = "ResourceNotFound.BrokerCluster"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
//  RESOURCENOTFOUND_ENVIRONMENT = "ResourceNotFound.Environment"
//  RESOURCEUNAVAILABLE_FUNDREQUIRED = "ResourceUnavailable.FundRequired"
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

// DeleteCluster
// 删除集群
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DELETECLUSTER = "FailedOperation.DeleteCluster"
//  FAILEDOPERATION_NAMESPACEINUSE = "FailedOperation.NamespaceInUse"
//  FAILEDOPERATION_ROLEINUSE = "FailedOperation.RoleInUse"
//  FAILEDOPERATION_VPCINUSE = "FailedOperation.VpcInUse"
//  OPERATIONDENIED_DEFAULTENVIRONMENT = "OperationDenied.DefaultEnvironment"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
func (c *Client) DeleteCluster(request *DeleteClusterRequest) (response *DeleteClusterResponse, err error) {
    return c.DeleteClusterWithContext(context.Background(), request)
}

// DeleteCluster
// 删除集群
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DELETECLUSTER = "FailedOperation.DeleteCluster"
//  FAILEDOPERATION_NAMESPACEINUSE = "FailedOperation.NamespaceInUse"
//  FAILEDOPERATION_ROLEINUSE = "FailedOperation.RoleInUse"
//  FAILEDOPERATION_VPCINUSE = "FailedOperation.VpcInUse"
//  OPERATIONDENIED_DEFAULTENVIRONMENT = "OperationDenied.DefaultEnvironment"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
func (c *Client) DeleteClusterWithContext(ctx context.Context, request *DeleteClusterRequest) (response *DeleteClusterResponse, err error) {
    if request == nil {
        request = NewDeleteClusterRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteCluster require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteClusterResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteCmqQueueRequest() (request *DeleteCmqQueueRequest) {
    request = &DeleteCmqQueueRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "DeleteCmqQueue")
    
    
    return
}

func NewDeleteCmqQueueResponse() (response *DeleteCmqQueueResponse) {
    response = &DeleteCmqQueueResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteCmqQueue
// 删除cmq队列
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DELETECLUSTER = "FailedOperation.DeleteCluster"
//  FAILEDOPERATION_DELETETOPICS = "FailedOperation.DeleteTopics"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DeleteCmqQueue(request *DeleteCmqQueueRequest) (response *DeleteCmqQueueResponse, err error) {
    return c.DeleteCmqQueueWithContext(context.Background(), request)
}

// DeleteCmqQueue
// 删除cmq队列
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DELETECLUSTER = "FailedOperation.DeleteCluster"
//  FAILEDOPERATION_DELETETOPICS = "FailedOperation.DeleteTopics"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DeleteCmqQueueWithContext(ctx context.Context, request *DeleteCmqQueueRequest) (response *DeleteCmqQueueResponse, err error) {
    if request == nil {
        request = NewDeleteCmqQueueRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteCmqQueue require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteCmqQueueResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteCmqSubscribeRequest() (request *DeleteCmqSubscribeRequest) {
    request = &DeleteCmqSubscribeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "DeleteCmqSubscribe")
    
    
    return
}

func NewDeleteCmqSubscribeResponse() (response *DeleteCmqSubscribeResponse) {
    response = &DeleteCmqSubscribeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteCmqSubscribe
// 删除cmq订阅
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DELETESUBSCRIPTIONS = "FailedOperation.DeleteSubscriptions"
//  FAILEDOPERATION_DELETETOPICS = "FailedOperation.DeleteTopics"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TOPIC = "ResourceNotFound.Topic"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DeleteCmqSubscribe(request *DeleteCmqSubscribeRequest) (response *DeleteCmqSubscribeResponse, err error) {
    return c.DeleteCmqSubscribeWithContext(context.Background(), request)
}

// DeleteCmqSubscribe
// 删除cmq订阅
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DELETESUBSCRIPTIONS = "FailedOperation.DeleteSubscriptions"
//  FAILEDOPERATION_DELETETOPICS = "FailedOperation.DeleteTopics"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TOPIC = "ResourceNotFound.Topic"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DeleteCmqSubscribeWithContext(ctx context.Context, request *DeleteCmqSubscribeRequest) (response *DeleteCmqSubscribeResponse, err error) {
    if request == nil {
        request = NewDeleteCmqSubscribeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteCmqSubscribe require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteCmqSubscribeResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteCmqTopicRequest() (request *DeleteCmqTopicRequest) {
    request = &DeleteCmqTopicRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "DeleteCmqTopic")
    
    
    return
}

func NewDeleteCmqTopicResponse() (response *DeleteCmqTopicResponse) {
    response = &DeleteCmqTopicResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteCmqTopic
// 删除cmq主题
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DELETECLUSTER = "FailedOperation.DeleteCluster"
//  RESOURCENOTFOUND_TOPIC = "ResourceNotFound.Topic"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DeleteCmqTopic(request *DeleteCmqTopicRequest) (response *DeleteCmqTopicResponse, err error) {
    return c.DeleteCmqTopicWithContext(context.Background(), request)
}

// DeleteCmqTopic
// 删除cmq主题
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DELETECLUSTER = "FailedOperation.DeleteCluster"
//  RESOURCENOTFOUND_TOPIC = "ResourceNotFound.Topic"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DeleteCmqTopicWithContext(ctx context.Context, request *DeleteCmqTopicRequest) (response *DeleteCmqTopicResponse, err error) {
    if request == nil {
        request = NewDeleteCmqTopicRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteCmqTopic require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteCmqTopicResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteEnvironmentRolesRequest() (request *DeleteEnvironmentRolesRequest) {
    request = &DeleteEnvironmentRolesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "DeleteEnvironmentRoles")
    
    
    return
}

func NewDeleteEnvironmentRolesResponse() (response *DeleteEnvironmentRolesResponse) {
    response = &DeleteEnvironmentRolesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteEnvironmentRoles
// 删除环境角色授权。
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DELETEENVIRONMENTROLES = "FailedOperation.DeleteEnvironmentRoles"
//  INVALIDPARAMETERVALUE_INVALIDPARAMS = "InvalidParameterValue.InvalidParams"
//  MISSINGPARAMETER_NEEDMOREPARAMS = "MissingParameter.NeedMoreParams"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
//  RESOURCENOTFOUND_ENVIRONMENTROLE = "ResourceNotFound.EnvironmentRole"
func (c *Client) DeleteEnvironmentRoles(request *DeleteEnvironmentRolesRequest) (response *DeleteEnvironmentRolesResponse, err error) {
    return c.DeleteEnvironmentRolesWithContext(context.Background(), request)
}

// DeleteEnvironmentRoles
// 删除环境角色授权。
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DELETEENVIRONMENTROLES = "FailedOperation.DeleteEnvironmentRoles"
//  INVALIDPARAMETERVALUE_INVALIDPARAMS = "InvalidParameterValue.InvalidParams"
//  MISSINGPARAMETER_NEEDMOREPARAMS = "MissingParameter.NeedMoreParams"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
//  RESOURCENOTFOUND_ENVIRONMENTROLE = "ResourceNotFound.EnvironmentRole"
func (c *Client) DeleteEnvironmentRolesWithContext(ctx context.Context, request *DeleteEnvironmentRolesRequest) (response *DeleteEnvironmentRolesResponse, err error) {
    if request == nil {
        request = NewDeleteEnvironmentRolesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteEnvironmentRoles require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteEnvironmentRolesResponse()
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

// DeleteEnvironments
// 批量删除租户下的命名空间
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_DELETEENVIRONMENTS = "FailedOperation.DeleteEnvironments"
//  FAILEDOPERATION_DELETENAMESPACE = "FailedOperation.DeleteNamespace"
//  FAILEDOPERATION_ROLEINUSE = "FailedOperation.RoleInUse"
//  FAILEDOPERATION_TOPICINUSE = "FailedOperation.TopicInUse"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDPARAMS = "InvalidParameterValue.InvalidParams"
//  MISSINGPARAMETER_NEEDMOREPARAMS = "MissingParameter.NeedMoreParams"
//  OPERATIONDENIED_DEFAULTENVIRONMENT = "OperationDenied.DefaultEnvironment"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
//  RESOURCENOTFOUND_ENVIRONMENT = "ResourceNotFound.Environment"
func (c *Client) DeleteEnvironments(request *DeleteEnvironmentsRequest) (response *DeleteEnvironmentsResponse, err error) {
    return c.DeleteEnvironmentsWithContext(context.Background(), request)
}

// DeleteEnvironments
// 批量删除租户下的命名空间
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_DELETEENVIRONMENTS = "FailedOperation.DeleteEnvironments"
//  FAILEDOPERATION_DELETENAMESPACE = "FailedOperation.DeleteNamespace"
//  FAILEDOPERATION_ROLEINUSE = "FailedOperation.RoleInUse"
//  FAILEDOPERATION_TOPICINUSE = "FailedOperation.TopicInUse"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDPARAMS = "InvalidParameterValue.InvalidParams"
//  MISSINGPARAMETER_NEEDMOREPARAMS = "MissingParameter.NeedMoreParams"
//  OPERATIONDENIED_DEFAULTENVIRONMENT = "OperationDenied.DefaultEnvironment"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
//  RESOURCENOTFOUND_ENVIRONMENT = "ResourceNotFound.Environment"
func (c *Client) DeleteEnvironmentsWithContext(ctx context.Context, request *DeleteEnvironmentsRequest) (response *DeleteEnvironmentsResponse, err error) {
    if request == nil {
        request = NewDeleteEnvironmentsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteEnvironments require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteEnvironmentsResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteRabbitMQUserRequest() (request *DeleteRabbitMQUserRequest) {
    request = &DeleteRabbitMQUserRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "DeleteRabbitMQUser")
    
    
    return
}

func NewDeleteRabbitMQUserResponse() (response *DeleteRabbitMQUserResponse) {
    response = &DeleteRabbitMQUserResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteRabbitMQUser
// 删除RabbitMQ的用户
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteRabbitMQUser(request *DeleteRabbitMQUserRequest) (response *DeleteRabbitMQUserResponse, err error) {
    return c.DeleteRabbitMQUserWithContext(context.Background(), request)
}

// DeleteRabbitMQUser
// 删除RabbitMQ的用户
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteRabbitMQUserWithContext(ctx context.Context, request *DeleteRabbitMQUserRequest) (response *DeleteRabbitMQUserResponse, err error) {
    if request == nil {
        request = NewDeleteRabbitMQUserRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteRabbitMQUser require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteRabbitMQUserResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteRabbitMQVipInstanceRequest() (request *DeleteRabbitMQVipInstanceRequest) {
    request = &DeleteRabbitMQVipInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "DeleteRabbitMQVipInstance")
    
    
    return
}

func NewDeleteRabbitMQVipInstanceResponse() (response *DeleteRabbitMQVipInstanceResponse) {
    response = &DeleteRabbitMQVipInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteRabbitMQVipInstance
// 删除RabbitMQ专享版实例
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DeleteRabbitMQVipInstance(request *DeleteRabbitMQVipInstanceRequest) (response *DeleteRabbitMQVipInstanceResponse, err error) {
    return c.DeleteRabbitMQVipInstanceWithContext(context.Background(), request)
}

// DeleteRabbitMQVipInstance
// 删除RabbitMQ专享版实例
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DeleteRabbitMQVipInstanceWithContext(ctx context.Context, request *DeleteRabbitMQVipInstanceRequest) (response *DeleteRabbitMQVipInstanceResponse, err error) {
    if request == nil {
        request = NewDeleteRabbitMQVipInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteRabbitMQVipInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteRabbitMQVipInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteRabbitMQVirtualHostRequest() (request *DeleteRabbitMQVirtualHostRequest) {
    request = &DeleteRabbitMQVirtualHostRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "DeleteRabbitMQVirtualHost")
    
    
    return
}

func NewDeleteRabbitMQVirtualHostResponse() (response *DeleteRabbitMQVirtualHostResponse) {
    response = &DeleteRabbitMQVirtualHostResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteRabbitMQVirtualHost
// 删除RabbitMQ的vhost
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteRabbitMQVirtualHost(request *DeleteRabbitMQVirtualHostRequest) (response *DeleteRabbitMQVirtualHostResponse, err error) {
    return c.DeleteRabbitMQVirtualHostWithContext(context.Background(), request)
}

// DeleteRabbitMQVirtualHost
// 删除RabbitMQ的vhost
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteRabbitMQVirtualHostWithContext(ctx context.Context, request *DeleteRabbitMQVirtualHostRequest) (response *DeleteRabbitMQVirtualHostResponse, err error) {
    if request == nil {
        request = NewDeleteRabbitMQVirtualHostRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteRabbitMQVirtualHost require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteRabbitMQVirtualHostResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteRocketMQClusterRequest() (request *DeleteRocketMQClusterRequest) {
    request = &DeleteRocketMQClusterRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "DeleteRocketMQCluster")
    
    
    return
}

func NewDeleteRocketMQClusterResponse() (response *DeleteRocketMQClusterResponse) {
    response = &DeleteRocketMQClusterResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteRocketMQCluster
// 删除RocketMQ集群
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DELETECLUSTER = "FailedOperation.DeleteCluster"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
func (c *Client) DeleteRocketMQCluster(request *DeleteRocketMQClusterRequest) (response *DeleteRocketMQClusterResponse, err error) {
    return c.DeleteRocketMQClusterWithContext(context.Background(), request)
}

// DeleteRocketMQCluster
// 删除RocketMQ集群
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DELETECLUSTER = "FailedOperation.DeleteCluster"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
func (c *Client) DeleteRocketMQClusterWithContext(ctx context.Context, request *DeleteRocketMQClusterRequest) (response *DeleteRocketMQClusterResponse, err error) {
    if request == nil {
        request = NewDeleteRocketMQClusterRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteRocketMQCluster require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteRocketMQClusterResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteRocketMQGroupRequest() (request *DeleteRocketMQGroupRequest) {
    request = &DeleteRocketMQGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "DeleteRocketMQGroup")
    
    
    return
}

func NewDeleteRocketMQGroupResponse() (response *DeleteRocketMQGroupResponse) {
    response = &DeleteRocketMQGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteRocketMQGroup
// 删除RocketMQ消费组
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ILLEGALMESSAGE = "InternalError.IllegalMessage"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteRocketMQGroup(request *DeleteRocketMQGroupRequest) (response *DeleteRocketMQGroupResponse, err error) {
    return c.DeleteRocketMQGroupWithContext(context.Background(), request)
}

// DeleteRocketMQGroup
// 删除RocketMQ消费组
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ILLEGALMESSAGE = "InternalError.IllegalMessage"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteRocketMQGroupWithContext(ctx context.Context, request *DeleteRocketMQGroupRequest) (response *DeleteRocketMQGroupResponse, err error) {
    if request == nil {
        request = NewDeleteRocketMQGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteRocketMQGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteRocketMQGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteRocketMQNamespaceRequest() (request *DeleteRocketMQNamespaceRequest) {
    request = &DeleteRocketMQNamespaceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "DeleteRocketMQNamespace")
    
    
    return
}

func NewDeleteRocketMQNamespaceResponse() (response *DeleteRocketMQNamespaceResponse) {
    response = &DeleteRocketMQNamespaceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteRocketMQNamespace
// 删除RocketMQ命名空间
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ROLEINUSE = "FailedOperation.RoleInUse"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteRocketMQNamespace(request *DeleteRocketMQNamespaceRequest) (response *DeleteRocketMQNamespaceResponse, err error) {
    return c.DeleteRocketMQNamespaceWithContext(context.Background(), request)
}

// DeleteRocketMQNamespace
// 删除RocketMQ命名空间
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ROLEINUSE = "FailedOperation.RoleInUse"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteRocketMQNamespaceWithContext(ctx context.Context, request *DeleteRocketMQNamespaceRequest) (response *DeleteRocketMQNamespaceResponse, err error) {
    if request == nil {
        request = NewDeleteRocketMQNamespaceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteRocketMQNamespace require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteRocketMQNamespaceResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteRocketMQTopicRequest() (request *DeleteRocketMQTopicRequest) {
    request = &DeleteRocketMQTopicRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "DeleteRocketMQTopic")
    
    
    return
}

func NewDeleteRocketMQTopicResponse() (response *DeleteRocketMQTopicResponse) {
    response = &DeleteRocketMQTopicResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteRocketMQTopic
// 删除RocketMQ主题
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DELETETOPICS = "FailedOperation.DeleteTopics"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteRocketMQTopic(request *DeleteRocketMQTopicRequest) (response *DeleteRocketMQTopicResponse, err error) {
    return c.DeleteRocketMQTopicWithContext(context.Background(), request)
}

// DeleteRocketMQTopic
// 删除RocketMQ主题
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DELETETOPICS = "FailedOperation.DeleteTopics"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteRocketMQTopicWithContext(ctx context.Context, request *DeleteRocketMQTopicRequest) (response *DeleteRocketMQTopicResponse, err error) {
    if request == nil {
        request = NewDeleteRocketMQTopicRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteRocketMQTopic require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteRocketMQTopicResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteRocketMQVipInstanceRequest() (request *DeleteRocketMQVipInstanceRequest) {
    request = &DeleteRocketMQVipInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "DeleteRocketMQVipInstance")
    
    
    return
}

func NewDeleteRocketMQVipInstanceResponse() (response *DeleteRocketMQVipInstanceResponse) {
    response = &DeleteRocketMQVipInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteRocketMQVipInstance
// 删除RocketMQ专享实例
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
func (c *Client) DeleteRocketMQVipInstance(request *DeleteRocketMQVipInstanceRequest) (response *DeleteRocketMQVipInstanceResponse, err error) {
    return c.DeleteRocketMQVipInstanceWithContext(context.Background(), request)
}

// DeleteRocketMQVipInstance
// 删除RocketMQ专享实例
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
func (c *Client) DeleteRocketMQVipInstanceWithContext(ctx context.Context, request *DeleteRocketMQVipInstanceRequest) (response *DeleteRocketMQVipInstanceResponse, err error) {
    if request == nil {
        request = NewDeleteRocketMQVipInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteRocketMQVipInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteRocketMQVipInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteRolesRequest() (request *DeleteRolesRequest) {
    request = &DeleteRolesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "DeleteRoles")
    
    
    return
}

func NewDeleteRolesResponse() (response *DeleteRolesResponse) {
    response = &DeleteRolesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteRoles
// 删除角色，支持批量。
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DELETEROLES = "FailedOperation.DeleteRoles"
//  INVALIDPARAMETERVALUE_INVALIDPARAMS = "InvalidParameterValue.InvalidParams"
//  MISSINGPARAMETER_NEEDMOREPARAMS = "MissingParameter.NeedMoreParams"
//  RESOURCEINUSE_ENVIRONMENTROLE = "ResourceInUse.EnvironmentRole"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
//  RESOURCENOTFOUND_ROLE = "ResourceNotFound.Role"
func (c *Client) DeleteRoles(request *DeleteRolesRequest) (response *DeleteRolesResponse, err error) {
    return c.DeleteRolesWithContext(context.Background(), request)
}

// DeleteRoles
// 删除角色，支持批量。
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DELETEROLES = "FailedOperation.DeleteRoles"
//  INVALIDPARAMETERVALUE_INVALIDPARAMS = "InvalidParameterValue.InvalidParams"
//  MISSINGPARAMETER_NEEDMOREPARAMS = "MissingParameter.NeedMoreParams"
//  RESOURCEINUSE_ENVIRONMENTROLE = "ResourceInUse.EnvironmentRole"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
//  RESOURCENOTFOUND_ROLE = "ResourceNotFound.Role"
func (c *Client) DeleteRolesWithContext(ctx context.Context, request *DeleteRolesRequest) (response *DeleteRolesResponse, err error) {
    if request == nil {
        request = NewDeleteRolesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteRoles require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteRolesResponse()
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

// DeleteSubscriptions
// 删除订阅关系
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DELETESUBSCRIPTIONS = "FailedOperation.DeleteSubscriptions"
//  FAILEDOPERATION_GETTOPICPARTITIONSFAILED = "FailedOperation.GetTopicPartitionsFailed"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMS = "InvalidParameterValue.InvalidParams"
//  MISSINGPARAMETER_NEEDMOREPARAMS = "MissingParameter.NeedMoreParams"
//  OPERATIONDENIED_CONSUMERRUNNING = "OperationDenied.ConsumerRunning"
//  RESOURCEINUSE_SUBSCRIPTION = "ResourceInUse.Subscription"
//  RESOURCENOTFOUND_BROKERCLUSTER = "ResourceNotFound.BrokerCluster"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
//  RESOURCENOTFOUND_SUBSCRIPTION = "ResourceNotFound.Subscription"
//  RESOURCENOTFOUND_TOPIC = "ResourceNotFound.Topic"
func (c *Client) DeleteSubscriptions(request *DeleteSubscriptionsRequest) (response *DeleteSubscriptionsResponse, err error) {
    return c.DeleteSubscriptionsWithContext(context.Background(), request)
}

// DeleteSubscriptions
// 删除订阅关系
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DELETESUBSCRIPTIONS = "FailedOperation.DeleteSubscriptions"
//  FAILEDOPERATION_GETTOPICPARTITIONSFAILED = "FailedOperation.GetTopicPartitionsFailed"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMS = "InvalidParameterValue.InvalidParams"
//  MISSINGPARAMETER_NEEDMOREPARAMS = "MissingParameter.NeedMoreParams"
//  OPERATIONDENIED_CONSUMERRUNNING = "OperationDenied.ConsumerRunning"
//  RESOURCEINUSE_SUBSCRIPTION = "ResourceInUse.Subscription"
//  RESOURCENOTFOUND_BROKERCLUSTER = "ResourceNotFound.BrokerCluster"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
//  RESOURCENOTFOUND_SUBSCRIPTION = "ResourceNotFound.Subscription"
//  RESOURCENOTFOUND_TOPIC = "ResourceNotFound.Topic"
func (c *Client) DeleteSubscriptionsWithContext(ctx context.Context, request *DeleteSubscriptionsRequest) (response *DeleteSubscriptionsResponse, err error) {
    if request == nil {
        request = NewDeleteSubscriptionsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteSubscriptions require credential")
    }

    request.SetContext(ctx)
    
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

// DeleteTopics
// 批量删除topics
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DELETETOPICS = "FailedOperation.DeleteTopics"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMS = "InvalidParameterValue.InvalidParams"
//  MISSINGPARAMETER_NEEDMOREPARAMS = "MissingParameter.NeedMoreParams"
//  OPERATIONDENIED_DEFAULTENVIRONMENT = "OperationDenied.DefaultEnvironment"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
//  RESOURCENOTFOUND_ENVIRONMENT = "ResourceNotFound.Environment"
//  RESOURCENOTFOUND_TOPIC = "ResourceNotFound.Topic"
func (c *Client) DeleteTopics(request *DeleteTopicsRequest) (response *DeleteTopicsResponse, err error) {
    return c.DeleteTopicsWithContext(context.Background(), request)
}

// DeleteTopics
// 批量删除topics
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DELETETOPICS = "FailedOperation.DeleteTopics"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMS = "InvalidParameterValue.InvalidParams"
//  MISSINGPARAMETER_NEEDMOREPARAMS = "MissingParameter.NeedMoreParams"
//  OPERATIONDENIED_DEFAULTENVIRONMENT = "OperationDenied.DefaultEnvironment"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
//  RESOURCENOTFOUND_ENVIRONMENT = "ResourceNotFound.Environment"
//  RESOURCENOTFOUND_TOPIC = "ResourceNotFound.Topic"
func (c *Client) DeleteTopicsWithContext(ctx context.Context, request *DeleteTopicsRequest) (response *DeleteTopicsResponse, err error) {
    if request == nil {
        request = NewDeleteTopicsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteTopics require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteTopicsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAMQPClustersRequest() (request *DescribeAMQPClustersRequest) {
    request = &DescribeAMQPClustersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "DescribeAMQPClusters")
    
    
    return
}

func NewDescribeAMQPClustersResponse() (response *DescribeAMQPClustersResponse) {
    response = &DescribeAMQPClustersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAMQPClusters
// 获取amqp集群列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAMQPClusters(request *DescribeAMQPClustersRequest) (response *DescribeAMQPClustersResponse, err error) {
    return c.DescribeAMQPClustersWithContext(context.Background(), request)
}

// DescribeAMQPClusters
// 获取amqp集群列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAMQPClustersWithContext(ctx context.Context, request *DescribeAMQPClustersRequest) (response *DescribeAMQPClustersResponse, err error) {
    if request == nil {
        request = NewDescribeAMQPClustersRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAMQPClusters require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAMQPClustersResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAllTenantsRequest() (request *DescribeAllTenantsRequest) {
    request = &DescribeAllTenantsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "DescribeAllTenants")
    
    
    return
}

func NewDescribeAllTenantsResponse() (response *DescribeAllTenantsResponse) {
    response = &DescribeAllTenantsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAllTenants
// 获取某个租户的虚拟集群列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAllTenants(request *DescribeAllTenantsRequest) (response *DescribeAllTenantsResponse, err error) {
    return c.DescribeAllTenantsWithContext(context.Background(), request)
}

// DescribeAllTenants
// 获取某个租户的虚拟集群列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAllTenantsWithContext(ctx context.Context, request *DescribeAllTenantsRequest) (response *DescribeAllTenantsResponse, err error) {
    if request == nil {
        request = NewDescribeAllTenantsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAllTenants require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAllTenantsResponse()
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

// DescribeBindClusters
// 获取用户绑定的专享集群列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_SYSTEMUPGRADE = "ResourceUnavailable.SystemUpgrade"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeBindClusters(request *DescribeBindClustersRequest) (response *DescribeBindClustersResponse, err error) {
    return c.DescribeBindClustersWithContext(context.Background(), request)
}

// DescribeBindClusters
// 获取用户绑定的专享集群列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_SYSTEMUPGRADE = "ResourceUnavailable.SystemUpgrade"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeBindClustersWithContext(ctx context.Context, request *DescribeBindClustersRequest) (response *DescribeBindClustersResponse, err error) {
    if request == nil {
        request = NewDescribeBindClustersRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBindClusters require credential")
    }

    request.SetContext(ctx)
    
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

// DescribeBindVpcs
// 获取租户VPC绑定关系
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMS = "InvalidParameterValue.InvalidParams"
//  INVALIDPARAMETERVALUE_NEEDMOREPARAMS = "InvalidParameterValue.NeedMoreParams"
//  MISSINGPARAMETER_NEEDMOREPARAMS = "MissingParameter.NeedMoreParams"
//  RESOURCENOTFOUND_BROKERCLUSTER = "ResourceNotFound.BrokerCluster"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
func (c *Client) DescribeBindVpcs(request *DescribeBindVpcsRequest) (response *DescribeBindVpcsResponse, err error) {
    return c.DescribeBindVpcsWithContext(context.Background(), request)
}

// DescribeBindVpcs
// 获取租户VPC绑定关系
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMS = "InvalidParameterValue.InvalidParams"
//  INVALIDPARAMETERVALUE_NEEDMOREPARAMS = "InvalidParameterValue.NeedMoreParams"
//  MISSINGPARAMETER_NEEDMOREPARAMS = "MissingParameter.NeedMoreParams"
//  RESOURCENOTFOUND_BROKERCLUSTER = "ResourceNotFound.BrokerCluster"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
func (c *Client) DescribeBindVpcsWithContext(ctx context.Context, request *DescribeBindVpcsRequest) (response *DescribeBindVpcsResponse, err error) {
    if request == nil {
        request = NewDescribeBindVpcsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBindVpcs require credential")
    }

    request.SetContext(ctx)
    
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

// DescribeClusterDetail
// 获取集群的详细信息
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
func (c *Client) DescribeClusterDetail(request *DescribeClusterDetailRequest) (response *DescribeClusterDetailResponse, err error) {
    return c.DescribeClusterDetailWithContext(context.Background(), request)
}

// DescribeClusterDetail
// 获取集群的详细信息
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
func (c *Client) DescribeClusterDetailWithContext(ctx context.Context, request *DescribeClusterDetailRequest) (response *DescribeClusterDetailResponse, err error) {
    if request == nil {
        request = NewDescribeClusterDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeClusterDetail require credential")
    }

    request.SetContext(ctx)
    
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

// DescribeClusters
// 获取集群列表
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
func (c *Client) DescribeClusters(request *DescribeClustersRequest) (response *DescribeClustersResponse, err error) {
    return c.DescribeClustersWithContext(context.Background(), request)
}

// DescribeClusters
// 获取集群列表
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
func (c *Client) DescribeClustersWithContext(ctx context.Context, request *DescribeClustersRequest) (response *DescribeClustersResponse, err error) {
    if request == nil {
        request = NewDescribeClustersRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeClusters require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeClustersResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCmqDeadLetterSourceQueuesRequest() (request *DescribeCmqDeadLetterSourceQueuesRequest) {
    request = &DescribeCmqDeadLetterSourceQueuesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "DescribeCmqDeadLetterSourceQueues")
    
    
    return
}

func NewDescribeCmqDeadLetterSourceQueuesResponse() (response *DescribeCmqDeadLetterSourceQueuesResponse) {
    response = &DescribeCmqDeadLetterSourceQueuesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCmqDeadLetterSourceQueues
// 枚举cmq死信队列源队列
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
func (c *Client) DescribeCmqDeadLetterSourceQueues(request *DescribeCmqDeadLetterSourceQueuesRequest) (response *DescribeCmqDeadLetterSourceQueuesResponse, err error) {
    return c.DescribeCmqDeadLetterSourceQueuesWithContext(context.Background(), request)
}

// DescribeCmqDeadLetterSourceQueues
// 枚举cmq死信队列源队列
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
func (c *Client) DescribeCmqDeadLetterSourceQueuesWithContext(ctx context.Context, request *DescribeCmqDeadLetterSourceQueuesRequest) (response *DescribeCmqDeadLetterSourceQueuesResponse, err error) {
    if request == nil {
        request = NewDescribeCmqDeadLetterSourceQueuesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCmqDeadLetterSourceQueues require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCmqDeadLetterSourceQueuesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCmqQueueDetailRequest() (request *DescribeCmqQueueDetailRequest) {
    request = &DescribeCmqQueueDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "DescribeCmqQueueDetail")
    
    
    return
}

func NewDescribeCmqQueueDetailResponse() (response *DescribeCmqQueueDetailResponse) {
    response = &DescribeCmqQueueDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCmqQueueDetail
// 查询cmq队列详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeCmqQueueDetail(request *DescribeCmqQueueDetailRequest) (response *DescribeCmqQueueDetailResponse, err error) {
    return c.DescribeCmqQueueDetailWithContext(context.Background(), request)
}

// DescribeCmqQueueDetail
// 查询cmq队列详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeCmqQueueDetailWithContext(ctx context.Context, request *DescribeCmqQueueDetailRequest) (response *DescribeCmqQueueDetailResponse, err error) {
    if request == nil {
        request = NewDescribeCmqQueueDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCmqQueueDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCmqQueueDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCmqQueuesRequest() (request *DescribeCmqQueuesRequest) {
    request = &DescribeCmqQueuesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "DescribeCmqQueues")
    
    
    return
}

func NewDescribeCmqQueuesResponse() (response *DescribeCmqQueuesResponse) {
    response = &DescribeCmqQueuesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCmqQueues
// 查询cmq全量队列
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeCmqQueues(request *DescribeCmqQueuesRequest) (response *DescribeCmqQueuesResponse, err error) {
    return c.DescribeCmqQueuesWithContext(context.Background(), request)
}

// DescribeCmqQueues
// 查询cmq全量队列
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeCmqQueuesWithContext(ctx context.Context, request *DescribeCmqQueuesRequest) (response *DescribeCmqQueuesResponse, err error) {
    if request == nil {
        request = NewDescribeCmqQueuesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCmqQueues require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCmqQueuesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCmqSubscriptionDetailRequest() (request *DescribeCmqSubscriptionDetailRequest) {
    request = &DescribeCmqSubscriptionDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "DescribeCmqSubscriptionDetail")
    
    
    return
}

func NewDescribeCmqSubscriptionDetailResponse() (response *DescribeCmqSubscriptionDetailResponse) {
    response = &DescribeCmqSubscriptionDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCmqSubscriptionDetail
// 查询cmq订阅详情
//
// 可能返回的错误码:
//  FAILEDOPERATION_CMQBACKENDERROR = "FailedOperation.CmqBackendError"
//  RESOURCENOTFOUND_TOPIC = "ResourceNotFound.Topic"
func (c *Client) DescribeCmqSubscriptionDetail(request *DescribeCmqSubscriptionDetailRequest) (response *DescribeCmqSubscriptionDetailResponse, err error) {
    return c.DescribeCmqSubscriptionDetailWithContext(context.Background(), request)
}

// DescribeCmqSubscriptionDetail
// 查询cmq订阅详情
//
// 可能返回的错误码:
//  FAILEDOPERATION_CMQBACKENDERROR = "FailedOperation.CmqBackendError"
//  RESOURCENOTFOUND_TOPIC = "ResourceNotFound.Topic"
func (c *Client) DescribeCmqSubscriptionDetailWithContext(ctx context.Context, request *DescribeCmqSubscriptionDetailRequest) (response *DescribeCmqSubscriptionDetailResponse, err error) {
    if request == nil {
        request = NewDescribeCmqSubscriptionDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCmqSubscriptionDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCmqSubscriptionDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCmqTopicDetailRequest() (request *DescribeCmqTopicDetailRequest) {
    request = &DescribeCmqTopicDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "DescribeCmqTopicDetail")
    
    
    return
}

func NewDescribeCmqTopicDetailResponse() (response *DescribeCmqTopicDetailResponse) {
    response = &DescribeCmqTopicDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCmqTopicDetail
// 查询cmq主题详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CMQBACKENDERROR = "FailedOperation.CmqBackendError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TOPIC = "ResourceNotFound.Topic"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeCmqTopicDetail(request *DescribeCmqTopicDetailRequest) (response *DescribeCmqTopicDetailResponse, err error) {
    return c.DescribeCmqTopicDetailWithContext(context.Background(), request)
}

// DescribeCmqTopicDetail
// 查询cmq主题详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CMQBACKENDERROR = "FailedOperation.CmqBackendError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TOPIC = "ResourceNotFound.Topic"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeCmqTopicDetailWithContext(ctx context.Context, request *DescribeCmqTopicDetailRequest) (response *DescribeCmqTopicDetailResponse, err error) {
    if request == nil {
        request = NewDescribeCmqTopicDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCmqTopicDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCmqTopicDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCmqTopicsRequest() (request *DescribeCmqTopicsRequest) {
    request = &DescribeCmqTopicsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "DescribeCmqTopics")
    
    
    return
}

func NewDescribeCmqTopicsResponse() (response *DescribeCmqTopicsResponse) {
    response = &DescribeCmqTopicsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCmqTopics
// 枚举cmq全量主题
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CMQBACKENDERROR = "FailedOperation.CmqBackendError"
func (c *Client) DescribeCmqTopics(request *DescribeCmqTopicsRequest) (response *DescribeCmqTopicsResponse, err error) {
    return c.DescribeCmqTopicsWithContext(context.Background(), request)
}

// DescribeCmqTopics
// 枚举cmq全量主题
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CMQBACKENDERROR = "FailedOperation.CmqBackendError"
func (c *Client) DescribeCmqTopicsWithContext(ctx context.Context, request *DescribeCmqTopicsRequest) (response *DescribeCmqTopicsResponse, err error) {
    if request == nil {
        request = NewDescribeCmqTopicsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCmqTopics require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCmqTopicsResponse()
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

// DescribeEnvironmentAttributes
// 获取指定命名空间的属性
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_GETENVIRONMENTATTRIBUTESFAILED = "FailedOperation.GetEnvironmentAttributesFailed"
//  INTERNALERROR_BROKERSERVICE = "InternalError.BrokerService"
//  INTERNALERROR_GETATTRIBUTESFAILED = "InternalError.GetAttributesFailed"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMS = "InvalidParameterValue.InvalidParams"
//  MISSINGPARAMETER_NEEDMOREPARAMS = "MissingParameter.NeedMoreParams"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
//  RESOURCENOTFOUND_ENVIRONMENT = "ResourceNotFound.Environment"
func (c *Client) DescribeEnvironmentAttributes(request *DescribeEnvironmentAttributesRequest) (response *DescribeEnvironmentAttributesResponse, err error) {
    return c.DescribeEnvironmentAttributesWithContext(context.Background(), request)
}

// DescribeEnvironmentAttributes
// 获取指定命名空间的属性
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_GETENVIRONMENTATTRIBUTESFAILED = "FailedOperation.GetEnvironmentAttributesFailed"
//  INTERNALERROR_BROKERSERVICE = "InternalError.BrokerService"
//  INTERNALERROR_GETATTRIBUTESFAILED = "InternalError.GetAttributesFailed"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMS = "InvalidParameterValue.InvalidParams"
//  MISSINGPARAMETER_NEEDMOREPARAMS = "MissingParameter.NeedMoreParams"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
//  RESOURCENOTFOUND_ENVIRONMENT = "ResourceNotFound.Environment"
func (c *Client) DescribeEnvironmentAttributesWithContext(ctx context.Context, request *DescribeEnvironmentAttributesRequest) (response *DescribeEnvironmentAttributesResponse, err error) {
    if request == nil {
        request = NewDescribeEnvironmentAttributesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeEnvironmentAttributes require credential")
    }

    request.SetContext(ctx)
    
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

// DescribeEnvironmentRoles
// 获取命名空间角色列表
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  INVALIDPARAMETERVALUE_INVALIDPARAMS = "InvalidParameterValue.InvalidParams"
//  MISSINGPARAMETER_NEEDMOREPARAMS = "MissingParameter.NeedMoreParams"
//  RESOURCENOTFOUND_BROKERCLUSTER = "ResourceNotFound.BrokerCluster"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
//  RESOURCENOTFOUND_ENVIRONMENT = "ResourceNotFound.Environment"
//  RESOURCENOTFOUND_ROLE = "ResourceNotFound.Role"
func (c *Client) DescribeEnvironmentRoles(request *DescribeEnvironmentRolesRequest) (response *DescribeEnvironmentRolesResponse, err error) {
    return c.DescribeEnvironmentRolesWithContext(context.Background(), request)
}

// DescribeEnvironmentRoles
// 获取命名空间角色列表
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  INVALIDPARAMETERVALUE_INVALIDPARAMS = "InvalidParameterValue.InvalidParams"
//  MISSINGPARAMETER_NEEDMOREPARAMS = "MissingParameter.NeedMoreParams"
//  RESOURCENOTFOUND_BROKERCLUSTER = "ResourceNotFound.BrokerCluster"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
//  RESOURCENOTFOUND_ENVIRONMENT = "ResourceNotFound.Environment"
//  RESOURCENOTFOUND_ROLE = "ResourceNotFound.Role"
func (c *Client) DescribeEnvironmentRolesWithContext(ctx context.Context, request *DescribeEnvironmentRolesRequest) (response *DescribeEnvironmentRolesResponse, err error) {
    if request == nil {
        request = NewDescribeEnvironmentRolesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeEnvironmentRoles require credential")
    }

    request.SetContext(ctx)
    
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

// DescribeEnvironments
// 获取租户下命名空间列表
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMS = "InvalidParameterValue.InvalidParams"
//  MISSINGPARAMETER_NEEDMOREPARAMS = "MissingParameter.NeedMoreParams"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
func (c *Client) DescribeEnvironments(request *DescribeEnvironmentsRequest) (response *DescribeEnvironmentsResponse, err error) {
    return c.DescribeEnvironmentsWithContext(context.Background(), request)
}

// DescribeEnvironments
// 获取租户下命名空间列表
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMS = "InvalidParameterValue.InvalidParams"
//  MISSINGPARAMETER_NEEDMOREPARAMS = "MissingParameter.NeedMoreParams"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
func (c *Client) DescribeEnvironmentsWithContext(ctx context.Context, request *DescribeEnvironmentsRequest) (response *DescribeEnvironmentsResponse, err error) {
    if request == nil {
        request = NewDescribeEnvironmentsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeEnvironments require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeEnvironmentsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeNamespaceBundlesOptRequest() (request *DescribeNamespaceBundlesOptRequest) {
    request = &DescribeNamespaceBundlesOptRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "DescribeNamespaceBundlesOpt")
    
    
    return
}

func NewDescribeNamespaceBundlesOptResponse() (response *DescribeNamespaceBundlesOptResponse) {
    response = &DescribeNamespaceBundlesOptResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeNamespaceBundlesOpt
// 运营端获取命名空间bundle列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_PULSARADMINERROR = "FailedOperation.PulsarAdminError"
//  OPERATIONDENIED_DEFAULTENVIRONMENT = "OperationDenied.DefaultEnvironment"
func (c *Client) DescribeNamespaceBundlesOpt(request *DescribeNamespaceBundlesOptRequest) (response *DescribeNamespaceBundlesOptResponse, err error) {
    return c.DescribeNamespaceBundlesOptWithContext(context.Background(), request)
}

// DescribeNamespaceBundlesOpt
// 运营端获取命名空间bundle列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_PULSARADMINERROR = "FailedOperation.PulsarAdminError"
//  OPERATIONDENIED_DEFAULTENVIRONMENT = "OperationDenied.DefaultEnvironment"
func (c *Client) DescribeNamespaceBundlesOptWithContext(ctx context.Context, request *DescribeNamespaceBundlesOptRequest) (response *DescribeNamespaceBundlesOptResponse, err error) {
    if request == nil {
        request = NewDescribeNamespaceBundlesOptRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeNamespaceBundlesOpt require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeNamespaceBundlesOptResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeNodeHealthOptRequest() (request *DescribeNodeHealthOptRequest) {
    request = &DescribeNodeHealthOptRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "DescribeNodeHealthOpt")
    
    
    return
}

func NewDescribeNodeHealthOptResponse() (response *DescribeNodeHealthOptResponse) {
    response = &DescribeNodeHealthOptResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeNodeHealthOpt
// 运营端获节点健康状态
//
// 可能返回的错误码:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeNodeHealthOpt(request *DescribeNodeHealthOptRequest) (response *DescribeNodeHealthOptResponse, err error) {
    return c.DescribeNodeHealthOptWithContext(context.Background(), request)
}

// DescribeNodeHealthOpt
// 运营端获节点健康状态
//
// 可能返回的错误码:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeNodeHealthOptWithContext(ctx context.Context, request *DescribeNodeHealthOptRequest) (response *DescribeNodeHealthOptResponse, err error) {
    if request == nil {
        request = NewDescribeNodeHealthOptRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeNodeHealthOpt require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeNodeHealthOptResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePublisherSummaryRequest() (request *DescribePublisherSummaryRequest) {
    request = &DescribePublisherSummaryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "DescribePublisherSummary")
    
    
    return
}

func NewDescribePublisherSummaryResponse() (response *DescribePublisherSummaryResponse) {
    response = &DescribePublisherSummaryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribePublisherSummary
// 获取消息生产概览信息
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_TOPIC = "ResourceNotFound.Topic"
func (c *Client) DescribePublisherSummary(request *DescribePublisherSummaryRequest) (response *DescribePublisherSummaryResponse, err error) {
    return c.DescribePublisherSummaryWithContext(context.Background(), request)
}

// DescribePublisherSummary
// 获取消息生产概览信息
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_TOPIC = "ResourceNotFound.Topic"
func (c *Client) DescribePublisherSummaryWithContext(ctx context.Context, request *DescribePublisherSummaryRequest) (response *DescribePublisherSummaryResponse, err error) {
    if request == nil {
        request = NewDescribePublisherSummaryRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePublisherSummary require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePublisherSummaryResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePublishersRequest() (request *DescribePublishersRequest) {
    request = &DescribePublishersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "DescribePublishers")
    
    
    return
}

func NewDescribePublishersResponse() (response *DescribePublishersResponse) {
    response = &DescribePublishersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribePublishers
// 获取生产者信息列表
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_TOPIC = "ResourceNotFound.Topic"
func (c *Client) DescribePublishers(request *DescribePublishersRequest) (response *DescribePublishersResponse, err error) {
    return c.DescribePublishersWithContext(context.Background(), request)
}

// DescribePublishers
// 获取生产者信息列表
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_TOPIC = "ResourceNotFound.Topic"
func (c *Client) DescribePublishersWithContext(ctx context.Context, request *DescribePublishersRequest) (response *DescribePublishersResponse, err error) {
    if request == nil {
        request = NewDescribePublishersRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePublishers require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePublishersResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePulsarProInstanceDetailRequest() (request *DescribePulsarProInstanceDetailRequest) {
    request = &DescribePulsarProInstanceDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "DescribePulsarProInstanceDetail")
    
    
    return
}

func NewDescribePulsarProInstanceDetailResponse() (response *DescribePulsarProInstanceDetailResponse) {
    response = &DescribePulsarProInstanceDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribePulsarProInstanceDetail
// 获取Pulsar专业版集群实例信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_ILLEGALMESSAGE = "InternalError.IllegalMessage"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
func (c *Client) DescribePulsarProInstanceDetail(request *DescribePulsarProInstanceDetailRequest) (response *DescribePulsarProInstanceDetailResponse, err error) {
    return c.DescribePulsarProInstanceDetailWithContext(context.Background(), request)
}

// DescribePulsarProInstanceDetail
// 获取Pulsar专业版集群实例信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_ILLEGALMESSAGE = "InternalError.IllegalMessage"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
func (c *Client) DescribePulsarProInstanceDetailWithContext(ctx context.Context, request *DescribePulsarProInstanceDetailRequest) (response *DescribePulsarProInstanceDetailResponse, err error) {
    if request == nil {
        request = NewDescribePulsarProInstanceDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePulsarProInstanceDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePulsarProInstanceDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePulsarProInstancesRequest() (request *DescribePulsarProInstancesRequest) {
    request = &DescribePulsarProInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "DescribePulsarProInstances")
    
    
    return
}

func NewDescribePulsarProInstancesResponse() (response *DescribePulsarProInstancesResponse) {
    response = &DescribePulsarProInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribePulsarProInstances
// 查询用户已购的Pulsar专业版实例列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_ILLEGALMESSAGE = "InternalError.IllegalMessage"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
func (c *Client) DescribePulsarProInstances(request *DescribePulsarProInstancesRequest) (response *DescribePulsarProInstancesResponse, err error) {
    return c.DescribePulsarProInstancesWithContext(context.Background(), request)
}

// DescribePulsarProInstances
// 查询用户已购的Pulsar专业版实例列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_ILLEGALMESSAGE = "InternalError.IllegalMessage"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
func (c *Client) DescribePulsarProInstancesWithContext(ctx context.Context, request *DescribePulsarProInstancesRequest) (response *DescribePulsarProInstancesResponse, err error) {
    if request == nil {
        request = NewDescribePulsarProInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePulsarProInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePulsarProInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRabbitMQNodeListRequest() (request *DescribeRabbitMQNodeListRequest) {
    request = &DescribeRabbitMQNodeListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "DescribeRabbitMQNodeList")
    
    
    return
}

func NewDescribeRabbitMQNodeListResponse() (response *DescribeRabbitMQNodeListResponse) {
    response = &DescribeRabbitMQNodeListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeRabbitMQNodeList
// RabbitMQ专享版查询节点列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
func (c *Client) DescribeRabbitMQNodeList(request *DescribeRabbitMQNodeListRequest) (response *DescribeRabbitMQNodeListResponse, err error) {
    return c.DescribeRabbitMQNodeListWithContext(context.Background(), request)
}

// DescribeRabbitMQNodeList
// RabbitMQ专享版查询节点列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
func (c *Client) DescribeRabbitMQNodeListWithContext(ctx context.Context, request *DescribeRabbitMQNodeListRequest) (response *DescribeRabbitMQNodeListResponse, err error) {
    if request == nil {
        request = NewDescribeRabbitMQNodeListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRabbitMQNodeList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRabbitMQNodeListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRabbitMQUserRequest() (request *DescribeRabbitMQUserRequest) {
    request = &DescribeRabbitMQUserRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "DescribeRabbitMQUser")
    
    
    return
}

func NewDescribeRabbitMQUserResponse() (response *DescribeRabbitMQUserResponse) {
    response = &DescribeRabbitMQUserResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeRabbitMQUser
// 查询RabbitMQ用户列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeRabbitMQUser(request *DescribeRabbitMQUserRequest) (response *DescribeRabbitMQUserResponse, err error) {
    return c.DescribeRabbitMQUserWithContext(context.Background(), request)
}

// DescribeRabbitMQUser
// 查询RabbitMQ用户列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeRabbitMQUserWithContext(ctx context.Context, request *DescribeRabbitMQUserRequest) (response *DescribeRabbitMQUserResponse, err error) {
    if request == nil {
        request = NewDescribeRabbitMQUserRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRabbitMQUser require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRabbitMQUserResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRabbitMQVipInstanceRequest() (request *DescribeRabbitMQVipInstanceRequest) {
    request = &DescribeRabbitMQVipInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "DescribeRabbitMQVipInstance")
    
    
    return
}

func NewDescribeRabbitMQVipInstanceResponse() (response *DescribeRabbitMQVipInstanceResponse) {
    response = &DescribeRabbitMQVipInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeRabbitMQVipInstance
// 获取单个RabbitMQ专享实例信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_ILLEGALMESSAGE = "InternalError.IllegalMessage"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
func (c *Client) DescribeRabbitMQVipInstance(request *DescribeRabbitMQVipInstanceRequest) (response *DescribeRabbitMQVipInstanceResponse, err error) {
    return c.DescribeRabbitMQVipInstanceWithContext(context.Background(), request)
}

// DescribeRabbitMQVipInstance
// 获取单个RabbitMQ专享实例信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_ILLEGALMESSAGE = "InternalError.IllegalMessage"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
func (c *Client) DescribeRabbitMQVipInstanceWithContext(ctx context.Context, request *DescribeRabbitMQVipInstanceRequest) (response *DescribeRabbitMQVipInstanceResponse, err error) {
    if request == nil {
        request = NewDescribeRabbitMQVipInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRabbitMQVipInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRabbitMQVipInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRabbitMQVipInstancesRequest() (request *DescribeRabbitMQVipInstancesRequest) {
    request = &DescribeRabbitMQVipInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "DescribeRabbitMQVipInstances")
    
    
    return
}

func NewDescribeRabbitMQVipInstancesResponse() (response *DescribeRabbitMQVipInstancesResponse) {
    response = &DescribeRabbitMQVipInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeRabbitMQVipInstances
// 查询用户已购的RabbitMQ专享实例列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_ILLEGALMESSAGE = "InternalError.IllegalMessage"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
func (c *Client) DescribeRabbitMQVipInstances(request *DescribeRabbitMQVipInstancesRequest) (response *DescribeRabbitMQVipInstancesResponse, err error) {
    return c.DescribeRabbitMQVipInstancesWithContext(context.Background(), request)
}

// DescribeRabbitMQVipInstances
// 查询用户已购的RabbitMQ专享实例列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_ILLEGALMESSAGE = "InternalError.IllegalMessage"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
func (c *Client) DescribeRabbitMQVipInstancesWithContext(ctx context.Context, request *DescribeRabbitMQVipInstancesRequest) (response *DescribeRabbitMQVipInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeRabbitMQVipInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRabbitMQVipInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRabbitMQVipInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRabbitMQVirtualHostRequest() (request *DescribeRabbitMQVirtualHostRequest) {
    request = &DescribeRabbitMQVirtualHostRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "DescribeRabbitMQVirtualHost")
    
    
    return
}

func NewDescribeRabbitMQVirtualHostResponse() (response *DescribeRabbitMQVirtualHostResponse) {
    response = &DescribeRabbitMQVirtualHostResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeRabbitMQVirtualHost
// 查询RabbitMQ vhost列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeRabbitMQVirtualHost(request *DescribeRabbitMQVirtualHostRequest) (response *DescribeRabbitMQVirtualHostResponse, err error) {
    return c.DescribeRabbitMQVirtualHostWithContext(context.Background(), request)
}

// DescribeRabbitMQVirtualHost
// 查询RabbitMQ vhost列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeRabbitMQVirtualHostWithContext(ctx context.Context, request *DescribeRabbitMQVirtualHostRequest) (response *DescribeRabbitMQVirtualHostResponse, err error) {
    if request == nil {
        request = NewDescribeRabbitMQVirtualHostRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRabbitMQVirtualHost require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRabbitMQVirtualHostResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRabbitMQVirtualHostListRequest() (request *DescribeRabbitMQVirtualHostListRequest) {
    request = &DescribeRabbitMQVirtualHostListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "DescribeRabbitMQVirtualHostList")
    
    
    return
}

func NewDescribeRabbitMQVirtualHostListResponse() (response *DescribeRabbitMQVirtualHostListResponse) {
    response = &DescribeRabbitMQVirtualHostListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeRabbitMQVirtualHostList
// RabbitMQ专享版查询虚拟主机列表
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
func (c *Client) DescribeRabbitMQVirtualHostList(request *DescribeRabbitMQVirtualHostListRequest) (response *DescribeRabbitMQVirtualHostListResponse, err error) {
    return c.DescribeRabbitMQVirtualHostListWithContext(context.Background(), request)
}

// DescribeRabbitMQVirtualHostList
// RabbitMQ专享版查询虚拟主机列表
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
func (c *Client) DescribeRabbitMQVirtualHostListWithContext(ctx context.Context, request *DescribeRabbitMQVirtualHostListRequest) (response *DescribeRabbitMQVirtualHostListResponse, err error) {
    if request == nil {
        request = NewDescribeRabbitMQVirtualHostListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRabbitMQVirtualHostList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRabbitMQVirtualHostListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRocketMQClusterRequest() (request *DescribeRocketMQClusterRequest) {
    request = &DescribeRocketMQClusterRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "DescribeRocketMQCluster")
    
    
    return
}

func NewDescribeRocketMQClusterResponse() (response *DescribeRocketMQClusterResponse) {
    response = &DescribeRocketMQClusterResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeRocketMQCluster
// 获取单个RocketMQ集群信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_ILLEGALMESSAGE = "InternalError.IllegalMessage"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
func (c *Client) DescribeRocketMQCluster(request *DescribeRocketMQClusterRequest) (response *DescribeRocketMQClusterResponse, err error) {
    return c.DescribeRocketMQClusterWithContext(context.Background(), request)
}

// DescribeRocketMQCluster
// 获取单个RocketMQ集群信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_ILLEGALMESSAGE = "InternalError.IllegalMessage"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
func (c *Client) DescribeRocketMQClusterWithContext(ctx context.Context, request *DescribeRocketMQClusterRequest) (response *DescribeRocketMQClusterResponse, err error) {
    if request == nil {
        request = NewDescribeRocketMQClusterRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRocketMQCluster require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRocketMQClusterResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRocketMQClustersRequest() (request *DescribeRocketMQClustersRequest) {
    request = &DescribeRocketMQClustersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "DescribeRocketMQClusters")
    
    
    return
}

func NewDescribeRocketMQClustersResponse() (response *DescribeRocketMQClustersResponse) {
    response = &DescribeRocketMQClustersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeRocketMQClusters
// 获取RocketMQ集群列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeRocketMQClusters(request *DescribeRocketMQClustersRequest) (response *DescribeRocketMQClustersResponse, err error) {
    return c.DescribeRocketMQClustersWithContext(context.Background(), request)
}

// DescribeRocketMQClusters
// 获取RocketMQ集群列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeRocketMQClustersWithContext(ctx context.Context, request *DescribeRocketMQClustersRequest) (response *DescribeRocketMQClustersResponse, err error) {
    if request == nil {
        request = NewDescribeRocketMQClustersRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRocketMQClusters require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRocketMQClustersResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRocketMQGroupsRequest() (request *DescribeRocketMQGroupsRequest) {
    request = &DescribeRocketMQGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "DescribeRocketMQGroups")
    
    
    return
}

func NewDescribeRocketMQGroupsResponse() (response *DescribeRocketMQGroupsResponse) {
    response = &DescribeRocketMQGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeRocketMQGroups
// 获取RocketMQ消费组列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeRocketMQGroups(request *DescribeRocketMQGroupsRequest) (response *DescribeRocketMQGroupsResponse, err error) {
    return c.DescribeRocketMQGroupsWithContext(context.Background(), request)
}

// DescribeRocketMQGroups
// 获取RocketMQ消费组列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeRocketMQGroupsWithContext(ctx context.Context, request *DescribeRocketMQGroupsRequest) (response *DescribeRocketMQGroupsResponse, err error) {
    if request == nil {
        request = NewDescribeRocketMQGroupsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRocketMQGroups require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRocketMQGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRocketMQMsgRequest() (request *DescribeRocketMQMsgRequest) {
    request = &DescribeRocketMQMsgRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "DescribeRocketMQMsg")
    
    
    return
}

func NewDescribeRocketMQMsgResponse() (response *DescribeRocketMQMsgResponse) {
    response = &DescribeRocketMQMsgResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeRocketMQMsg
// rocketmq消息详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
func (c *Client) DescribeRocketMQMsg(request *DescribeRocketMQMsgRequest) (response *DescribeRocketMQMsgResponse, err error) {
    return c.DescribeRocketMQMsgWithContext(context.Background(), request)
}

// DescribeRocketMQMsg
// rocketmq消息详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
func (c *Client) DescribeRocketMQMsgWithContext(ctx context.Context, request *DescribeRocketMQMsgRequest) (response *DescribeRocketMQMsgResponse, err error) {
    if request == nil {
        request = NewDescribeRocketMQMsgRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRocketMQMsg require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRocketMQMsgResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRocketMQNamespacesRequest() (request *DescribeRocketMQNamespacesRequest) {
    request = &DescribeRocketMQNamespacesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "DescribeRocketMQNamespaces")
    
    
    return
}

func NewDescribeRocketMQNamespacesResponse() (response *DescribeRocketMQNamespacesResponse) {
    response = &DescribeRocketMQNamespacesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeRocketMQNamespaces
// 获取RocketMQ命名空间列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ILLEGALMESSAGE = "InternalError.IllegalMessage"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeRocketMQNamespaces(request *DescribeRocketMQNamespacesRequest) (response *DescribeRocketMQNamespacesResponse, err error) {
    return c.DescribeRocketMQNamespacesWithContext(context.Background(), request)
}

// DescribeRocketMQNamespaces
// 获取RocketMQ命名空间列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ILLEGALMESSAGE = "InternalError.IllegalMessage"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeRocketMQNamespacesWithContext(ctx context.Context, request *DescribeRocketMQNamespacesRequest) (response *DescribeRocketMQNamespacesResponse, err error) {
    if request == nil {
        request = NewDescribeRocketMQNamespacesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRocketMQNamespaces require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRocketMQNamespacesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRocketMQTopicsRequest() (request *DescribeRocketMQTopicsRequest) {
    request = &DescribeRocketMQTopicsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "DescribeRocketMQTopics")
    
    
    return
}

func NewDescribeRocketMQTopicsResponse() (response *DescribeRocketMQTopicsResponse) {
    response = &DescribeRocketMQTopicsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeRocketMQTopics
// 获取RocketMQ主题列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ILLEGALMESSAGE = "InternalError.IllegalMessage"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_BROKERCLUSTER = "ResourceNotFound.BrokerCluster"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeRocketMQTopics(request *DescribeRocketMQTopicsRequest) (response *DescribeRocketMQTopicsResponse, err error) {
    return c.DescribeRocketMQTopicsWithContext(context.Background(), request)
}

// DescribeRocketMQTopics
// 获取RocketMQ主题列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ILLEGALMESSAGE = "InternalError.IllegalMessage"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_BROKERCLUSTER = "ResourceNotFound.BrokerCluster"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeRocketMQTopicsWithContext(ctx context.Context, request *DescribeRocketMQTopicsRequest) (response *DescribeRocketMQTopicsResponse, err error) {
    if request == nil {
        request = NewDescribeRocketMQTopicsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRocketMQTopics require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRocketMQTopicsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRocketMQVipInstanceDetailRequest() (request *DescribeRocketMQVipInstanceDetailRequest) {
    request = &DescribeRocketMQVipInstanceDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "DescribeRocketMQVipInstanceDetail")
    
    
    return
}

func NewDescribeRocketMQVipInstanceDetailResponse() (response *DescribeRocketMQVipInstanceDetailResponse) {
    response = &DescribeRocketMQVipInstanceDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeRocketMQVipInstanceDetail
// 获取单个RocketMQ专享集群信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_ILLEGALMESSAGE = "InternalError.IllegalMessage"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
func (c *Client) DescribeRocketMQVipInstanceDetail(request *DescribeRocketMQVipInstanceDetailRequest) (response *DescribeRocketMQVipInstanceDetailResponse, err error) {
    return c.DescribeRocketMQVipInstanceDetailWithContext(context.Background(), request)
}

// DescribeRocketMQVipInstanceDetail
// 获取单个RocketMQ专享集群信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_ILLEGALMESSAGE = "InternalError.IllegalMessage"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
func (c *Client) DescribeRocketMQVipInstanceDetailWithContext(ctx context.Context, request *DescribeRocketMQVipInstanceDetailRequest) (response *DescribeRocketMQVipInstanceDetailResponse, err error) {
    if request == nil {
        request = NewDescribeRocketMQVipInstanceDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRocketMQVipInstanceDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRocketMQVipInstanceDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRocketMQVipInstancesRequest() (request *DescribeRocketMQVipInstancesRequest) {
    request = &DescribeRocketMQVipInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "DescribeRocketMQVipInstances")
    
    
    return
}

func NewDescribeRocketMQVipInstancesResponse() (response *DescribeRocketMQVipInstancesResponse) {
    response = &DescribeRocketMQVipInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeRocketMQVipInstances
// 查询用户已购的RocketMQ专享实例列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_ILLEGALMESSAGE = "InternalError.IllegalMessage"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
func (c *Client) DescribeRocketMQVipInstances(request *DescribeRocketMQVipInstancesRequest) (response *DescribeRocketMQVipInstancesResponse, err error) {
    return c.DescribeRocketMQVipInstancesWithContext(context.Background(), request)
}

// DescribeRocketMQVipInstances
// 查询用户已购的RocketMQ专享实例列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_ILLEGALMESSAGE = "InternalError.IllegalMessage"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
func (c *Client) DescribeRocketMQVipInstancesWithContext(ctx context.Context, request *DescribeRocketMQVipInstancesRequest) (response *DescribeRocketMQVipInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeRocketMQVipInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRocketMQVipInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRocketMQVipInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRolesRequest() (request *DescribeRolesRequest) {
    request = &DescribeRolesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "DescribeRoles")
    
    
    return
}

func NewDescribeRolesResponse() (response *DescribeRolesResponse) {
    response = &DescribeRolesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeRoles
// 获取角色列表
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMS = "InvalidParameterValue.InvalidParams"
//  MISSINGPARAMETER_NEEDMOREPARAMS = "MissingParameter.NeedMoreParams"
//  RESOURCENOTFOUND_BROKERCLUSTER = "ResourceNotFound.BrokerCluster"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
func (c *Client) DescribeRoles(request *DescribeRolesRequest) (response *DescribeRolesResponse, err error) {
    return c.DescribeRolesWithContext(context.Background(), request)
}

// DescribeRoles
// 获取角色列表
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMS = "InvalidParameterValue.InvalidParams"
//  MISSINGPARAMETER_NEEDMOREPARAMS = "MissingParameter.NeedMoreParams"
//  RESOURCENOTFOUND_BROKERCLUSTER = "ResourceNotFound.BrokerCluster"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
func (c *Client) DescribeRolesWithContext(ctx context.Context, request *DescribeRolesRequest) (response *DescribeRolesResponse, err error) {
    if request == nil {
        request = NewDescribeRolesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRoles require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRolesResponse()
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

// DescribeSubscriptions
// 查询指定环境和主题下的订阅者列表
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DESCRIBESUBSCRIPTION = "FailedOperation.DescribeSubscription"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMS = "InvalidParameterValue.InvalidParams"
//  MISSINGPARAMETER_NEEDMOREPARAMS = "MissingParameter.NeedMoreParams"
//  RESOURCENOTFOUND_BROKERCLUSTER = "ResourceNotFound.BrokerCluster"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
//  RESOURCENOTFOUND_ENVIRONMENT = "ResourceNotFound.Environment"
//  RESOURCENOTFOUND_TOPIC = "ResourceNotFound.Topic"
func (c *Client) DescribeSubscriptions(request *DescribeSubscriptionsRequest) (response *DescribeSubscriptionsResponse, err error) {
    return c.DescribeSubscriptionsWithContext(context.Background(), request)
}

// DescribeSubscriptions
// 查询指定环境和主题下的订阅者列表
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DESCRIBESUBSCRIPTION = "FailedOperation.DescribeSubscription"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMS = "InvalidParameterValue.InvalidParams"
//  MISSINGPARAMETER_NEEDMOREPARAMS = "MissingParameter.NeedMoreParams"
//  RESOURCENOTFOUND_BROKERCLUSTER = "ResourceNotFound.BrokerCluster"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
//  RESOURCENOTFOUND_ENVIRONMENT = "ResourceNotFound.Environment"
//  RESOURCENOTFOUND_TOPIC = "ResourceNotFound.Topic"
func (c *Client) DescribeSubscriptionsWithContext(ctx context.Context, request *DescribeSubscriptionsRequest) (response *DescribeSubscriptionsResponse, err error) {
    if request == nil {
        request = NewDescribeSubscriptionsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSubscriptions require credential")
    }

    request.SetContext(ctx)
    
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

// DescribeTopics
// 获取环境下主题列表
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INVALIDADMINURL = "InvalidParameter.InvalidAdminUrl"
//  INVALIDPARAMETERVALUE_INVALIDPARAMS = "InvalidParameterValue.InvalidParams"
//  MISSINGPARAMETER_NEEDMOREPARAMS = "MissingParameter.NeedMoreParams"
//  RESOURCENOTFOUND_BROKERCLUSTER = "ResourceNotFound.BrokerCluster"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
//  RESOURCENOTFOUND_ENVIRONMENT = "ResourceNotFound.Environment"
func (c *Client) DescribeTopics(request *DescribeTopicsRequest) (response *DescribeTopicsResponse, err error) {
    return c.DescribeTopicsWithContext(context.Background(), request)
}

// DescribeTopics
// 获取环境下主题列表
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INVALIDADMINURL = "InvalidParameter.InvalidAdminUrl"
//  INVALIDPARAMETERVALUE_INVALIDPARAMS = "InvalidParameterValue.InvalidParams"
//  MISSINGPARAMETER_NEEDMOREPARAMS = "MissingParameter.NeedMoreParams"
//  RESOURCENOTFOUND_BROKERCLUSTER = "ResourceNotFound.BrokerCluster"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
//  RESOURCENOTFOUND_ENVIRONMENT = "ResourceNotFound.Environment"
func (c *Client) DescribeTopicsWithContext(ctx context.Context, request *DescribeTopicsRequest) (response *DescribeTopicsResponse, err error) {
    if request == nil {
        request = NewDescribeTopicsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTopics require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTopicsResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAMQPClusterRequest() (request *ModifyAMQPClusterRequest) {
    request = &ModifyAMQPClusterRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "ModifyAMQPCluster")
    
    
    return
}

func NewModifyAMQPClusterResponse() (response *ModifyAMQPClusterResponse) {
    response = &ModifyAMQPClusterResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyAMQPCluster
// 更新Amqp集群信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyAMQPCluster(request *ModifyAMQPClusterRequest) (response *ModifyAMQPClusterResponse, err error) {
    return c.ModifyAMQPClusterWithContext(context.Background(), request)
}

// ModifyAMQPCluster
// 更新Amqp集群信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyAMQPClusterWithContext(ctx context.Context, request *ModifyAMQPClusterRequest) (response *ModifyAMQPClusterResponse, err error) {
    if request == nil {
        request = NewModifyAMQPClusterRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAMQPCluster require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAMQPClusterResponse()
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

// ModifyCluster
// 更新集群信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE_CLUSTERNAMEDUPLICATION = "InvalidParameterValue.ClusterNameDuplication"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEUNAVAILABLE_FUNDREQUIRED = "ResourceUnavailable.FundRequired"
func (c *Client) ModifyCluster(request *ModifyClusterRequest) (response *ModifyClusterResponse, err error) {
    return c.ModifyClusterWithContext(context.Background(), request)
}

// ModifyCluster
// 更新集群信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE_CLUSTERNAMEDUPLICATION = "InvalidParameterValue.ClusterNameDuplication"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEUNAVAILABLE_FUNDREQUIRED = "ResourceUnavailable.FundRequired"
func (c *Client) ModifyClusterWithContext(ctx context.Context, request *ModifyClusterRequest) (response *ModifyClusterResponse, err error) {
    if request == nil {
        request = NewModifyClusterRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyCluster require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyClusterResponse()
    err = c.Send(request, response)
    return
}

func NewModifyCmqQueueAttributeRequest() (request *ModifyCmqQueueAttributeRequest) {
    request = &ModifyCmqQueueAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "ModifyCmqQueueAttribute")
    
    
    return
}

func NewModifyCmqQueueAttributeResponse() (response *ModifyCmqQueueAttributeResponse) {
    response = &ModifyCmqQueueAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyCmqQueueAttribute
// 修改cmq队列属性
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_SETRETENTIONPOLICY = "FailedOperation.SetRetentionPolicy"
//  FAILEDOPERATION_SETTTL = "FailedOperation.SetTTL"
//  INTERNALERROR_BROKERSERVICE = "InternalError.BrokerService"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_ENVIRONMENT = "ResourceNotFound.Environment"
//  RESOURCEUNAVAILABLE_FUNDREQUIRED = "ResourceUnavailable.FundRequired"
func (c *Client) ModifyCmqQueueAttribute(request *ModifyCmqQueueAttributeRequest) (response *ModifyCmqQueueAttributeResponse, err error) {
    return c.ModifyCmqQueueAttributeWithContext(context.Background(), request)
}

// ModifyCmqQueueAttribute
// 修改cmq队列属性
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_SETRETENTIONPOLICY = "FailedOperation.SetRetentionPolicy"
//  FAILEDOPERATION_SETTTL = "FailedOperation.SetTTL"
//  INTERNALERROR_BROKERSERVICE = "InternalError.BrokerService"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_ENVIRONMENT = "ResourceNotFound.Environment"
//  RESOURCEUNAVAILABLE_FUNDREQUIRED = "ResourceUnavailable.FundRequired"
func (c *Client) ModifyCmqQueueAttributeWithContext(ctx context.Context, request *ModifyCmqQueueAttributeRequest) (response *ModifyCmqQueueAttributeResponse, err error) {
    if request == nil {
        request = NewModifyCmqQueueAttributeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyCmqQueueAttribute require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyCmqQueueAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyCmqSubscriptionAttributeRequest() (request *ModifyCmqSubscriptionAttributeRequest) {
    request = &ModifyCmqSubscriptionAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "ModifyCmqSubscriptionAttribute")
    
    
    return
}

func NewModifyCmqSubscriptionAttributeResponse() (response *ModifyCmqSubscriptionAttributeResponse) {
    response = &ModifyCmqSubscriptionAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyCmqSubscriptionAttribute
// 修改cmq订阅属性
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ModifyCmqSubscriptionAttribute(request *ModifyCmqSubscriptionAttributeRequest) (response *ModifyCmqSubscriptionAttributeResponse, err error) {
    return c.ModifyCmqSubscriptionAttributeWithContext(context.Background(), request)
}

// ModifyCmqSubscriptionAttribute
// 修改cmq订阅属性
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ModifyCmqSubscriptionAttributeWithContext(ctx context.Context, request *ModifyCmqSubscriptionAttributeRequest) (response *ModifyCmqSubscriptionAttributeResponse, err error) {
    if request == nil {
        request = NewModifyCmqSubscriptionAttributeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyCmqSubscriptionAttribute require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyCmqSubscriptionAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyCmqTopicAttributeRequest() (request *ModifyCmqTopicAttributeRequest) {
    request = &ModifyCmqTopicAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "ModifyCmqTopicAttribute")
    
    
    return
}

func NewModifyCmqTopicAttributeResponse() (response *ModifyCmqTopicAttributeResponse) {
    response = &ModifyCmqTopicAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyCmqTopicAttribute
// 修改cmq主题属性
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCENOTFOUND_TOPIC = "ResourceNotFound.Topic"
func (c *Client) ModifyCmqTopicAttribute(request *ModifyCmqTopicAttributeRequest) (response *ModifyCmqTopicAttributeResponse, err error) {
    return c.ModifyCmqTopicAttributeWithContext(context.Background(), request)
}

// ModifyCmqTopicAttribute
// 修改cmq主题属性
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCENOTFOUND_TOPIC = "ResourceNotFound.Topic"
func (c *Client) ModifyCmqTopicAttributeWithContext(ctx context.Context, request *ModifyCmqTopicAttributeRequest) (response *ModifyCmqTopicAttributeResponse, err error) {
    if request == nil {
        request = NewModifyCmqTopicAttributeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyCmqTopicAttribute require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyCmqTopicAttributeResponse()
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

// ModifyEnvironmentAttributes
// 修改指定命名空间的属性值
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_SETRETENTIONPOLICY = "FailedOperation.SetRetentionPolicy"
//  FAILEDOPERATION_SETTTL = "FailedOperation.SetTTL"
//  FAILEDOPERATION_UPDATEENVIRONMENT = "FailedOperation.UpdateEnvironment"
//  INTERNALERROR_BROKERSERVICE = "InternalError.BrokerService"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMS = "InvalidParameterValue.InvalidParams"
//  INVALIDPARAMETERVALUE_TTL = "InvalidParameterValue.TTL"
//  LIMITEXCEEDED_RETENTIONTIME = "LimitExceeded.RetentionTime"
//  MISSINGPARAMETER_NEEDMOREPARAMS = "MissingParameter.NeedMoreParams"
//  OPERATIONDENIED_DEFAULTENVIRONMENT = "OperationDenied.DefaultEnvironment"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
//  RESOURCENOTFOUND_ENVIRONMENT = "ResourceNotFound.Environment"
//  RESOURCENOTFOUND_NAMESPACE = "ResourceNotFound.Namespace"
//  RESOURCEUNAVAILABLE_FUNDREQUIRED = "ResourceUnavailable.FundRequired"
func (c *Client) ModifyEnvironmentAttributes(request *ModifyEnvironmentAttributesRequest) (response *ModifyEnvironmentAttributesResponse, err error) {
    return c.ModifyEnvironmentAttributesWithContext(context.Background(), request)
}

// ModifyEnvironmentAttributes
// 修改指定命名空间的属性值
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_SETRETENTIONPOLICY = "FailedOperation.SetRetentionPolicy"
//  FAILEDOPERATION_SETTTL = "FailedOperation.SetTTL"
//  FAILEDOPERATION_UPDATEENVIRONMENT = "FailedOperation.UpdateEnvironment"
//  INTERNALERROR_BROKERSERVICE = "InternalError.BrokerService"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMS = "InvalidParameterValue.InvalidParams"
//  INVALIDPARAMETERVALUE_TTL = "InvalidParameterValue.TTL"
//  LIMITEXCEEDED_RETENTIONTIME = "LimitExceeded.RetentionTime"
//  MISSINGPARAMETER_NEEDMOREPARAMS = "MissingParameter.NeedMoreParams"
//  OPERATIONDENIED_DEFAULTENVIRONMENT = "OperationDenied.DefaultEnvironment"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
//  RESOURCENOTFOUND_ENVIRONMENT = "ResourceNotFound.Environment"
//  RESOURCENOTFOUND_NAMESPACE = "ResourceNotFound.Namespace"
//  RESOURCEUNAVAILABLE_FUNDREQUIRED = "ResourceUnavailable.FundRequired"
func (c *Client) ModifyEnvironmentAttributesWithContext(ctx context.Context, request *ModifyEnvironmentAttributesRequest) (response *ModifyEnvironmentAttributesResponse, err error) {
    if request == nil {
        request = NewModifyEnvironmentAttributesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyEnvironmentAttributes require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyEnvironmentAttributesResponse()
    err = c.Send(request, response)
    return
}

func NewModifyEnvironmentRoleRequest() (request *ModifyEnvironmentRoleRequest) {
    request = &ModifyEnvironmentRoleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "ModifyEnvironmentRole")
    
    
    return
}

func NewModifyEnvironmentRoleResponse() (response *ModifyEnvironmentRoleResponse) {
    response = &ModifyEnvironmentRoleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyEnvironmentRole
// 修改环境角色授权。
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_UPDATEENVIRONMENTROLE = "FailedOperation.UpdateEnvironmentRole"
//  INVALIDPARAMETERVALUE_INVALIDPARAMS = "InvalidParameterValue.InvalidParams"
//  MISSINGPARAMETER_NEEDMOREPARAMS = "MissingParameter.NeedMoreParams"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
//  RESOURCENOTFOUND_ENVIRONMENTROLE = "ResourceNotFound.EnvironmentRole"
//  RESOURCEUNAVAILABLE_FUNDREQUIRED = "ResourceUnavailable.FundRequired"
func (c *Client) ModifyEnvironmentRole(request *ModifyEnvironmentRoleRequest) (response *ModifyEnvironmentRoleResponse, err error) {
    return c.ModifyEnvironmentRoleWithContext(context.Background(), request)
}

// ModifyEnvironmentRole
// 修改环境角色授权。
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_UPDATEENVIRONMENTROLE = "FailedOperation.UpdateEnvironmentRole"
//  INVALIDPARAMETERVALUE_INVALIDPARAMS = "InvalidParameterValue.InvalidParams"
//  MISSINGPARAMETER_NEEDMOREPARAMS = "MissingParameter.NeedMoreParams"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
//  RESOURCENOTFOUND_ENVIRONMENTROLE = "ResourceNotFound.EnvironmentRole"
//  RESOURCEUNAVAILABLE_FUNDREQUIRED = "ResourceUnavailable.FundRequired"
func (c *Client) ModifyEnvironmentRoleWithContext(ctx context.Context, request *ModifyEnvironmentRoleRequest) (response *ModifyEnvironmentRoleResponse, err error) {
    if request == nil {
        request = NewModifyEnvironmentRoleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyEnvironmentRole require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyEnvironmentRoleResponse()
    err = c.Send(request, response)
    return
}

func NewModifyRabbitMQUserRequest() (request *ModifyRabbitMQUserRequest) {
    request = &ModifyRabbitMQUserRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "ModifyRabbitMQUser")
    
    
    return
}

func NewModifyRabbitMQUserResponse() (response *ModifyRabbitMQUserResponse) {
    response = &ModifyRabbitMQUserResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyRabbitMQUser
// 修改RabbitMQ的用户
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyRabbitMQUser(request *ModifyRabbitMQUserRequest) (response *ModifyRabbitMQUserResponse, err error) {
    return c.ModifyRabbitMQUserWithContext(context.Background(), request)
}

// ModifyRabbitMQUser
// 修改RabbitMQ的用户
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyRabbitMQUserWithContext(ctx context.Context, request *ModifyRabbitMQUserRequest) (response *ModifyRabbitMQUserResponse, err error) {
    if request == nil {
        request = NewModifyRabbitMQUserRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyRabbitMQUser require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyRabbitMQUserResponse()
    err = c.Send(request, response)
    return
}

func NewModifyRabbitMQVipInstanceRequest() (request *ModifyRabbitMQVipInstanceRequest) {
    request = &ModifyRabbitMQVipInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "ModifyRabbitMQVipInstance")
    
    
    return
}

func NewModifyRabbitMQVipInstanceResponse() (response *ModifyRabbitMQVipInstanceResponse) {
    response = &ModifyRabbitMQVipInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyRabbitMQVipInstance
// 修改RabbitMQ专享版实例
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ModifyRabbitMQVipInstance(request *ModifyRabbitMQVipInstanceRequest) (response *ModifyRabbitMQVipInstanceResponse, err error) {
    return c.ModifyRabbitMQVipInstanceWithContext(context.Background(), request)
}

// ModifyRabbitMQVipInstance
// 修改RabbitMQ专享版实例
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ModifyRabbitMQVipInstanceWithContext(ctx context.Context, request *ModifyRabbitMQVipInstanceRequest) (response *ModifyRabbitMQVipInstanceResponse, err error) {
    if request == nil {
        request = NewModifyRabbitMQVipInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyRabbitMQVipInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyRabbitMQVipInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewModifyRabbitMQVirtualHostRequest() (request *ModifyRabbitMQVirtualHostRequest) {
    request = &ModifyRabbitMQVirtualHostRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "ModifyRabbitMQVirtualHost")
    
    
    return
}

func NewModifyRabbitMQVirtualHostResponse() (response *ModifyRabbitMQVirtualHostResponse) {
    response = &ModifyRabbitMQVirtualHostResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyRabbitMQVirtualHost
// 修改RabbitMQ的vhost
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyRabbitMQVirtualHost(request *ModifyRabbitMQVirtualHostRequest) (response *ModifyRabbitMQVirtualHostResponse, err error) {
    return c.ModifyRabbitMQVirtualHostWithContext(context.Background(), request)
}

// ModifyRabbitMQVirtualHost
// 修改RabbitMQ的vhost
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyRabbitMQVirtualHostWithContext(ctx context.Context, request *ModifyRabbitMQVirtualHostRequest) (response *ModifyRabbitMQVirtualHostResponse, err error) {
    if request == nil {
        request = NewModifyRabbitMQVirtualHostRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyRabbitMQVirtualHost require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyRabbitMQVirtualHostResponse()
    err = c.Send(request, response)
    return
}

func NewModifyRocketMQClusterRequest() (request *ModifyRocketMQClusterRequest) {
    request = &ModifyRocketMQClusterRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "ModifyRocketMQCluster")
    
    
    return
}

func NewModifyRocketMQClusterResponse() (response *ModifyRocketMQClusterResponse) {
    response = &ModifyRocketMQClusterResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyRocketMQCluster
// 更新RocketMQ集群信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
//  RESOURCEUNAVAILABLE_FUNDREQUIRED = "ResourceUnavailable.FundRequired"
func (c *Client) ModifyRocketMQCluster(request *ModifyRocketMQClusterRequest) (response *ModifyRocketMQClusterResponse, err error) {
    return c.ModifyRocketMQClusterWithContext(context.Background(), request)
}

// ModifyRocketMQCluster
// 更新RocketMQ集群信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
//  RESOURCEUNAVAILABLE_FUNDREQUIRED = "ResourceUnavailable.FundRequired"
func (c *Client) ModifyRocketMQClusterWithContext(ctx context.Context, request *ModifyRocketMQClusterRequest) (response *ModifyRocketMQClusterResponse, err error) {
    if request == nil {
        request = NewModifyRocketMQClusterRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyRocketMQCluster require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyRocketMQClusterResponse()
    err = c.Send(request, response)
    return
}

func NewModifyRocketMQGroupRequest() (request *ModifyRocketMQGroupRequest) {
    request = &ModifyRocketMQGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "ModifyRocketMQGroup")
    
    
    return
}

func NewModifyRocketMQGroupResponse() (response *ModifyRocketMQGroupResponse) {
    response = &ModifyRocketMQGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyRocketMQGroup
// 更新RocketMQ消费组信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyRocketMQGroup(request *ModifyRocketMQGroupRequest) (response *ModifyRocketMQGroupResponse, err error) {
    return c.ModifyRocketMQGroupWithContext(context.Background(), request)
}

// ModifyRocketMQGroup
// 更新RocketMQ消费组信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyRocketMQGroupWithContext(ctx context.Context, request *ModifyRocketMQGroupRequest) (response *ModifyRocketMQGroupResponse, err error) {
    if request == nil {
        request = NewModifyRocketMQGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyRocketMQGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyRocketMQGroupResponse()
    err = c.Send(request, response)
    return
}

func NewModifyRocketMQInstanceSpecRequest() (request *ModifyRocketMQInstanceSpecRequest) {
    request = &ModifyRocketMQInstanceSpecRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "ModifyRocketMQInstanceSpec")
    
    
    return
}

func NewModifyRocketMQInstanceSpecResponse() (response *ModifyRocketMQInstanceSpecResponse) {
    response = &ModifyRocketMQInstanceSpecResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyRocketMQInstanceSpec
// 本API用于修改RocketMQ专享实例配置，可以支持实例规格、节点数和存储的升配和实例规格的降配。本API发起订单并成功支付后进入实例配置变更的流程，可通过DescribeRocketMQVipInstances查询实例是否已变更完成。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CALLTRADE = "FailedOperation.CallTrade"
//  FAILEDOPERATION_INSTANCENOTREADY = "FailedOperation.InstanceNotReady"
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
//  UNSUPPORTEDOPERATION_INSTANCEDOWNGRADE = "UnsupportedOperation.InstanceDowngrade"
func (c *Client) ModifyRocketMQInstanceSpec(request *ModifyRocketMQInstanceSpecRequest) (response *ModifyRocketMQInstanceSpecResponse, err error) {
    return c.ModifyRocketMQInstanceSpecWithContext(context.Background(), request)
}

// ModifyRocketMQInstanceSpec
// 本API用于修改RocketMQ专享实例配置，可以支持实例规格、节点数和存储的升配和实例规格的降配。本API发起订单并成功支付后进入实例配置变更的流程，可通过DescribeRocketMQVipInstances查询实例是否已变更完成。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CALLTRADE = "FailedOperation.CallTrade"
//  FAILEDOPERATION_INSTANCENOTREADY = "FailedOperation.InstanceNotReady"
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
//  UNSUPPORTEDOPERATION_INSTANCEDOWNGRADE = "UnsupportedOperation.InstanceDowngrade"
func (c *Client) ModifyRocketMQInstanceSpecWithContext(ctx context.Context, request *ModifyRocketMQInstanceSpecRequest) (response *ModifyRocketMQInstanceSpecResponse, err error) {
    if request == nil {
        request = NewModifyRocketMQInstanceSpecRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyRocketMQInstanceSpec require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyRocketMQInstanceSpecResponse()
    err = c.Send(request, response)
    return
}

func NewModifyRocketMQNamespaceRequest() (request *ModifyRocketMQNamespaceRequest) {
    request = &ModifyRocketMQNamespaceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "ModifyRocketMQNamespace")
    
    
    return
}

func NewModifyRocketMQNamespaceResponse() (response *ModifyRocketMQNamespaceResponse) {
    response = &ModifyRocketMQNamespaceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyRocketMQNamespace
// 更新RocketMQ命名空间
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_SETTTL = "FailedOperation.SetTTL"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
//  RESOURCEUNAVAILABLE_FUNDREQUIRED = "ResourceUnavailable.FundRequired"
func (c *Client) ModifyRocketMQNamespace(request *ModifyRocketMQNamespaceRequest) (response *ModifyRocketMQNamespaceResponse, err error) {
    return c.ModifyRocketMQNamespaceWithContext(context.Background(), request)
}

// ModifyRocketMQNamespace
// 更新RocketMQ命名空间
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_SETTTL = "FailedOperation.SetTTL"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
//  RESOURCEUNAVAILABLE_FUNDREQUIRED = "ResourceUnavailable.FundRequired"
func (c *Client) ModifyRocketMQNamespaceWithContext(ctx context.Context, request *ModifyRocketMQNamespaceRequest) (response *ModifyRocketMQNamespaceResponse, err error) {
    if request == nil {
        request = NewModifyRocketMQNamespaceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyRocketMQNamespace require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyRocketMQNamespaceResponse()
    err = c.Send(request, response)
    return
}

func NewModifyRocketMQTopicRequest() (request *ModifyRocketMQTopicRequest) {
    request = &ModifyRocketMQTopicRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "ModifyRocketMQTopic")
    
    
    return
}

func NewModifyRocketMQTopicResponse() (response *ModifyRocketMQTopicResponse) {
    response = &ModifyRocketMQTopicResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyRocketMQTopic
// 更新RocketMQ主题信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARTITION = "InvalidParameter.Partition"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyRocketMQTopic(request *ModifyRocketMQTopicRequest) (response *ModifyRocketMQTopicResponse, err error) {
    return c.ModifyRocketMQTopicWithContext(context.Background(), request)
}

// ModifyRocketMQTopic
// 更新RocketMQ主题信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARTITION = "InvalidParameter.Partition"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyRocketMQTopicWithContext(ctx context.Context, request *ModifyRocketMQTopicRequest) (response *ModifyRocketMQTopicResponse, err error) {
    if request == nil {
        request = NewModifyRocketMQTopicRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyRocketMQTopic require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyRocketMQTopicResponse()
    err = c.Send(request, response)
    return
}

func NewModifyRoleRequest() (request *ModifyRoleRequest) {
    request = &ModifyRoleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "ModifyRole")
    
    
    return
}

func NewModifyRoleResponse() (response *ModifyRoleResponse) {
    response = &ModifyRoleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyRole
// 角色修改
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_UPDATEROLE = "FailedOperation.UpdateRole"
//  INVALIDPARAMETERVALUE_INVALIDPARAMS = "InvalidParameterValue.InvalidParams"
//  MISSINGPARAMETER_NEEDMOREPARAMS = "MissingParameter.NeedMoreParams"
//  RESOURCENOTFOUND_BROKERCLUSTER = "ResourceNotFound.BrokerCluster"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
//  RESOURCENOTFOUND_ROLE = "ResourceNotFound.Role"
//  RESOURCEUNAVAILABLE_FUNDREQUIRED = "ResourceUnavailable.FundRequired"
func (c *Client) ModifyRole(request *ModifyRoleRequest) (response *ModifyRoleResponse, err error) {
    return c.ModifyRoleWithContext(context.Background(), request)
}

// ModifyRole
// 角色修改
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_UPDATEROLE = "FailedOperation.UpdateRole"
//  INVALIDPARAMETERVALUE_INVALIDPARAMS = "InvalidParameterValue.InvalidParams"
//  MISSINGPARAMETER_NEEDMOREPARAMS = "MissingParameter.NeedMoreParams"
//  RESOURCENOTFOUND_BROKERCLUSTER = "ResourceNotFound.BrokerCluster"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
//  RESOURCENOTFOUND_ROLE = "ResourceNotFound.Role"
//  RESOURCEUNAVAILABLE_FUNDREQUIRED = "ResourceUnavailable.FundRequired"
func (c *Client) ModifyRoleWithContext(ctx context.Context, request *ModifyRoleRequest) (response *ModifyRoleResponse, err error) {
    if request == nil {
        request = NewModifyRoleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyRole require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyRoleResponse()
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

// ModifyTopic
// 修改主题备注和分区数
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_GETTOPICPARTITIONSFAILED = "FailedOperation.GetTopicPartitionsFailed"
//  FAILEDOPERATION_UPDATETOPIC = "FailedOperation.UpdateTopic"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMS = "InvalidParameterValue.InvalidParams"
//  MISSINGPARAMETER_NEEDMOREPARAMS = "MissingParameter.NeedMoreParams"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
//  RESOURCENOTFOUND_TOPIC = "ResourceNotFound.Topic"
func (c *Client) ModifyTopic(request *ModifyTopicRequest) (response *ModifyTopicResponse, err error) {
    return c.ModifyTopicWithContext(context.Background(), request)
}

// ModifyTopic
// 修改主题备注和分区数
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_GETTOPICPARTITIONSFAILED = "FailedOperation.GetTopicPartitionsFailed"
//  FAILEDOPERATION_UPDATETOPIC = "FailedOperation.UpdateTopic"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMS = "InvalidParameterValue.InvalidParams"
//  MISSINGPARAMETER_NEEDMOREPARAMS = "MissingParameter.NeedMoreParams"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
//  RESOURCENOTFOUND_TOPIC = "ResourceNotFound.Topic"
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

func NewPublishCmqMsgRequest() (request *PublishCmqMsgRequest) {
    request = &PublishCmqMsgRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "PublishCmqMsg")
    
    
    return
}

func NewPublishCmqMsgResponse() (response *PublishCmqMsgResponse) {
    response = &PublishCmqMsgResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// PublishCmqMsg
// 发送cmq主题消息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) PublishCmqMsg(request *PublishCmqMsgRequest) (response *PublishCmqMsgResponse, err error) {
    return c.PublishCmqMsgWithContext(context.Background(), request)
}

// PublishCmqMsg
// 发送cmq主题消息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) PublishCmqMsgWithContext(ctx context.Context, request *PublishCmqMsgRequest) (response *PublishCmqMsgResponse, err error) {
    if request == nil {
        request = NewPublishCmqMsgRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("PublishCmqMsg require credential")
    }

    request.SetContext(ctx)
    
    response = NewPublishCmqMsgResponse()
    err = c.Send(request, response)
    return
}

func NewReceiveMessageRequest() (request *ReceiveMessageRequest) {
    request = &ReceiveMessageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "ReceiveMessage")
    
    
    return
}

func NewReceiveMessageResponse() (response *ReceiveMessageResponse) {
    response = &ReceiveMessageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ReceiveMessage
// 当前 ReceiveMessage 接口只支持 Partitioned 类型的 Topic。该接口用于接收发送到指定 Partitioned Topic 中的消息，当 Partitioned Topic 中没有消息但还去尝试调用该接口时，会抛出 ReceiveTimeout 的异常。
//
// 
//
// 如何使用 BatchReceivePolicy：
//
// 
//
// BatchReceive 接口提供了如下三个参数：
//
// 
//
// ● MaxNumMessages: 即每次使用 BatchReceive 的时候，最多一次Receive接口返回多少条消息。
//
// ● MaxNumBytes：即每次使用 BatchReceive 的时候，最多一次Receive接口返回多大内容的消息，单位是：bytes。
//
// ● Timeout：即每次使用 BatchReceive 的时候，最多一次 Receive 接口的超时时间是多久，单位是：MS。
//
// 
//
// 默认如果上述三个参数都不指定，即关闭 BatchReceive 的特性。如果三个参数中的任意一个参数指定的数值大于 0，即开启 BatchReceive。BatchReceive 的结束条件为到达上述三个参数中任意一个指定的阈值。
//
// 
//
// 注意：MaxNumMessages 和 MaxNumBytes 每一次接收的最大消息同时受限于 ReceiveQueueSize 的大小，如果 ReceiveQueueSize 的大小设置为 5，MaxNumMessages 设置为10，那么一次 BatchReceive 接收的最多的消息是 5条，而不是10条。
//
// 
//
// 
//
// 
//
// BatchReceivePolicy 的接口会一次性返回多条消息：
//
// 
//
// 1. 多条消息的内容之间使用特殊字符 '###' 来进行分割，业务侧接收到消息之后，可以利用不同语言提供的 Split 工具分割不同的消息。
//
// 2. 多条消息的 MessageID 之间使用特殊字符 '###' 来进行分割，业务侧接收到消息之后，可以利用不同语言提供的 Split 工具分割不同的消息。（用于在调用 AcknowledgeMessage 接口中填入所需要的 MessageID 字段信息）
//
// 可能返回的错误码:
//  FAILEDOPERATION_CREATEPRODUCERERROR = "FailedOperation.CreateProducerError"
//  FAILEDOPERATION_CREATEPULSARCLIENTERROR = "FailedOperation.CreatePulsarClientError"
//  FAILEDOPERATION_MAXMESSAGESIZEERROR = "FailedOperation.MaxMessageSizeError"
//  FAILEDOPERATION_MESSAGEIDERROR = "FailedOperation.MessageIDError"
//  FAILEDOPERATION_RECEIVEERROR = "FailedOperation.ReceiveError"
//  FAILEDOPERATION_RECEIVETIMEOUT = "FailedOperation.ReceiveTimeout"
//  FAILEDOPERATION_SENDMESSAGETIMEOUTERROR = "FailedOperation.SendMessageTimeoutError"
//  FAILEDOPERATION_TOPICTYPEERROR = "FailedOperation.TopicTypeError"
//  INVALIDPARAMETER_TENANTNOTFOUND = "InvalidParameter.TenantNotFound"
//  INVALIDPARAMETER_TOKENNOTFOUND = "InvalidParameter.TokenNotFound"
//  INVALIDPARAMETERVALUE_TOPICNOTFOUND = "InvalidParameterValue.TopicNotFound"
func (c *Client) ReceiveMessage(request *ReceiveMessageRequest) (response *ReceiveMessageResponse, err error) {
    return c.ReceiveMessageWithContext(context.Background(), request)
}

// ReceiveMessage
// 当前 ReceiveMessage 接口只支持 Partitioned 类型的 Topic。该接口用于接收发送到指定 Partitioned Topic 中的消息，当 Partitioned Topic 中没有消息但还去尝试调用该接口时，会抛出 ReceiveTimeout 的异常。
//
// 
//
// 如何使用 BatchReceivePolicy：
//
// 
//
// BatchReceive 接口提供了如下三个参数：
//
// 
//
// ● MaxNumMessages: 即每次使用 BatchReceive 的时候，最多一次Receive接口返回多少条消息。
//
// ● MaxNumBytes：即每次使用 BatchReceive 的时候，最多一次Receive接口返回多大内容的消息，单位是：bytes。
//
// ● Timeout：即每次使用 BatchReceive 的时候，最多一次 Receive 接口的超时时间是多久，单位是：MS。
//
// 
//
// 默认如果上述三个参数都不指定，即关闭 BatchReceive 的特性。如果三个参数中的任意一个参数指定的数值大于 0，即开启 BatchReceive。BatchReceive 的结束条件为到达上述三个参数中任意一个指定的阈值。
//
// 
//
// 注意：MaxNumMessages 和 MaxNumBytes 每一次接收的最大消息同时受限于 ReceiveQueueSize 的大小，如果 ReceiveQueueSize 的大小设置为 5，MaxNumMessages 设置为10，那么一次 BatchReceive 接收的最多的消息是 5条，而不是10条。
//
// 
//
// 
//
// 
//
// BatchReceivePolicy 的接口会一次性返回多条消息：
//
// 
//
// 1. 多条消息的内容之间使用特殊字符 '###' 来进行分割，业务侧接收到消息之后，可以利用不同语言提供的 Split 工具分割不同的消息。
//
// 2. 多条消息的 MessageID 之间使用特殊字符 '###' 来进行分割，业务侧接收到消息之后，可以利用不同语言提供的 Split 工具分割不同的消息。（用于在调用 AcknowledgeMessage 接口中填入所需要的 MessageID 字段信息）
//
// 可能返回的错误码:
//  FAILEDOPERATION_CREATEPRODUCERERROR = "FailedOperation.CreateProducerError"
//  FAILEDOPERATION_CREATEPULSARCLIENTERROR = "FailedOperation.CreatePulsarClientError"
//  FAILEDOPERATION_MAXMESSAGESIZEERROR = "FailedOperation.MaxMessageSizeError"
//  FAILEDOPERATION_MESSAGEIDERROR = "FailedOperation.MessageIDError"
//  FAILEDOPERATION_RECEIVEERROR = "FailedOperation.ReceiveError"
//  FAILEDOPERATION_RECEIVETIMEOUT = "FailedOperation.ReceiveTimeout"
//  FAILEDOPERATION_SENDMESSAGETIMEOUTERROR = "FailedOperation.SendMessageTimeoutError"
//  FAILEDOPERATION_TOPICTYPEERROR = "FailedOperation.TopicTypeError"
//  INVALIDPARAMETER_TENANTNOTFOUND = "InvalidParameter.TenantNotFound"
//  INVALIDPARAMETER_TOKENNOTFOUND = "InvalidParameter.TokenNotFound"
//  INVALIDPARAMETERVALUE_TOPICNOTFOUND = "InvalidParameterValue.TopicNotFound"
func (c *Client) ReceiveMessageWithContext(ctx context.Context, request *ReceiveMessageRequest) (response *ReceiveMessageResponse, err error) {
    if request == nil {
        request = NewReceiveMessageRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ReceiveMessage require credential")
    }

    request.SetContext(ctx)
    
    response = NewReceiveMessageResponse()
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

// ResetMsgSubOffsetByTimestamp
// 根据时间戳进行消息回溯，精确到毫秒
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_RESETMSGSUBOFFSETBYTIMESTAMPFAILED = "FailedOperation.ResetMsgSubOffsetByTimestampFailed"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMS = "InvalidParameterValue.InvalidParams"
//  MISSINGPARAMETER_NEEDMOREPARAMS = "MissingParameter.NeedMoreParams"
//  RESOURCENOTFOUND_BROKERCLUSTER = "ResourceNotFound.BrokerCluster"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
//  RESOURCENOTFOUND_SUBSCRIPTION = "ResourceNotFound.Subscription"
//  RESOURCENOTFOUND_TOPIC = "ResourceNotFound.Topic"
//  RESOURCEUNAVAILABLE_FUNDREQUIRED = "ResourceUnavailable.FundRequired"
func (c *Client) ResetMsgSubOffsetByTimestamp(request *ResetMsgSubOffsetByTimestampRequest) (response *ResetMsgSubOffsetByTimestampResponse, err error) {
    return c.ResetMsgSubOffsetByTimestampWithContext(context.Background(), request)
}

// ResetMsgSubOffsetByTimestamp
// 根据时间戳进行消息回溯，精确到毫秒
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_RESETMSGSUBOFFSETBYTIMESTAMPFAILED = "FailedOperation.ResetMsgSubOffsetByTimestampFailed"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMS = "InvalidParameterValue.InvalidParams"
//  MISSINGPARAMETER_NEEDMOREPARAMS = "MissingParameter.NeedMoreParams"
//  RESOURCENOTFOUND_BROKERCLUSTER = "ResourceNotFound.BrokerCluster"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
//  RESOURCENOTFOUND_SUBSCRIPTION = "ResourceNotFound.Subscription"
//  RESOURCENOTFOUND_TOPIC = "ResourceNotFound.Topic"
//  RESOURCEUNAVAILABLE_FUNDREQUIRED = "ResourceUnavailable.FundRequired"
func (c *Client) ResetMsgSubOffsetByTimestampWithContext(ctx context.Context, request *ResetMsgSubOffsetByTimestampRequest) (response *ResetMsgSubOffsetByTimestampResponse, err error) {
    if request == nil {
        request = NewResetMsgSubOffsetByTimestampRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ResetMsgSubOffsetByTimestamp require credential")
    }

    request.SetContext(ctx)
    
    response = NewResetMsgSubOffsetByTimestampResponse()
    err = c.Send(request, response)
    return
}

func NewResetRocketMQConsumerOffSetRequest() (request *ResetRocketMQConsumerOffSetRequest) {
    request = &ResetRocketMQConsumerOffSetRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "ResetRocketMQConsumerOffSet")
    
    
    return
}

func NewResetRocketMQConsumerOffSetResponse() (response *ResetRocketMQConsumerOffSetResponse) {
    response = &ResetRocketMQConsumerOffSetResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ResetRocketMQConsumerOffSet
// 重置指定Group的消费位点到指定时间戳
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ResetRocketMQConsumerOffSet(request *ResetRocketMQConsumerOffSetRequest) (response *ResetRocketMQConsumerOffSetResponse, err error) {
    return c.ResetRocketMQConsumerOffSetWithContext(context.Background(), request)
}

// ResetRocketMQConsumerOffSet
// 重置指定Group的消费位点到指定时间戳
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ResetRocketMQConsumerOffSetWithContext(ctx context.Context, request *ResetRocketMQConsumerOffSetRequest) (response *ResetRocketMQConsumerOffSetResponse, err error) {
    if request == nil {
        request = NewResetRocketMQConsumerOffSetRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ResetRocketMQConsumerOffSet require credential")
    }

    request.SetContext(ctx)
    
    response = NewResetRocketMQConsumerOffSetResponse()
    err = c.Send(request, response)
    return
}

func NewRewindCmqQueueRequest() (request *RewindCmqQueueRequest) {
    request = &RewindCmqQueueRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "RewindCmqQueue")
    
    
    return
}

func NewRewindCmqQueueResponse() (response *RewindCmqQueueResponse) {
    response = &RewindCmqQueueResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// RewindCmqQueue
// 回溯cmq队列
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) RewindCmqQueue(request *RewindCmqQueueRequest) (response *RewindCmqQueueResponse, err error) {
    return c.RewindCmqQueueWithContext(context.Background(), request)
}

// RewindCmqQueue
// 回溯cmq队列
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) RewindCmqQueueWithContext(ctx context.Context, request *RewindCmqQueueRequest) (response *RewindCmqQueueResponse, err error) {
    if request == nil {
        request = NewRewindCmqQueueRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RewindCmqQueue require credential")
    }

    request.SetContext(ctx)
    
    response = NewRewindCmqQueueResponse()
    err = c.Send(request, response)
    return
}

func NewSendBatchMessagesRequest() (request *SendBatchMessagesRequest) {
    request = &SendBatchMessagesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "SendBatchMessages")
    
    
    return
}

func NewSendBatchMessagesResponse() (response *SendBatchMessagesResponse) {
    response = &SendBatchMessagesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// SendBatchMessages
// 批量发送消息
//
// 
//
// 注意：TDMQ 批量发送消息的接口是在 TDMQ-HTTP 的服务侧将消息打包为一个 Batch，然后将该 Batch 在服务内部当作一次 TCP 请求发送出去。所以在使用过程中，用户还是按照单条消息发送的逻辑，每一条消息是一个独立的 HTTP 的请求，在 TDMQ-HTTP 的服务内部，会将多个 HTTP 的请求聚合为一个 Batch 发送到服务端。即，批量发送消息在使用上与发送单条消息是一致的，batch 的聚合是在 TDMQ-HTTP 的服务内部完成的。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CREATEPRODUCERERROR = "FailedOperation.CreateProducerError"
//  FAILEDOPERATION_CREATEPULSARCLIENTERROR = "FailedOperation.CreatePulsarClientError"
//  FAILEDOPERATION_MAXMESSAGESIZEERROR = "FailedOperation.MaxMessageSizeError"
//  FAILEDOPERATION_MESSAGEIDERROR = "FailedOperation.MessageIDError"
//  FAILEDOPERATION_RECEIVEERROR = "FailedOperation.ReceiveError"
//  FAILEDOPERATION_RECEIVETIMEOUT = "FailedOperation.ReceiveTimeout"
//  FAILEDOPERATION_SENDMESSAGETIMEOUTERROR = "FailedOperation.SendMessageTimeoutError"
//  FAILEDOPERATION_TOPICTYPEERROR = "FailedOperation.TopicTypeError"
//  INVALIDPARAMETER_TENANTNOTFOUND = "InvalidParameter.TenantNotFound"
//  INVALIDPARAMETER_TOKENNOTFOUND = "InvalidParameter.TokenNotFound"
//  INVALIDPARAMETERVALUE_TOPICNOTFOUND = "InvalidParameterValue.TopicNotFound"
func (c *Client) SendBatchMessages(request *SendBatchMessagesRequest) (response *SendBatchMessagesResponse, err error) {
    return c.SendBatchMessagesWithContext(context.Background(), request)
}

// SendBatchMessages
// 批量发送消息
//
// 
//
// 注意：TDMQ 批量发送消息的接口是在 TDMQ-HTTP 的服务侧将消息打包为一个 Batch，然后将该 Batch 在服务内部当作一次 TCP 请求发送出去。所以在使用过程中，用户还是按照单条消息发送的逻辑，每一条消息是一个独立的 HTTP 的请求，在 TDMQ-HTTP 的服务内部，会将多个 HTTP 的请求聚合为一个 Batch 发送到服务端。即，批量发送消息在使用上与发送单条消息是一致的，batch 的聚合是在 TDMQ-HTTP 的服务内部完成的。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CREATEPRODUCERERROR = "FailedOperation.CreateProducerError"
//  FAILEDOPERATION_CREATEPULSARCLIENTERROR = "FailedOperation.CreatePulsarClientError"
//  FAILEDOPERATION_MAXMESSAGESIZEERROR = "FailedOperation.MaxMessageSizeError"
//  FAILEDOPERATION_MESSAGEIDERROR = "FailedOperation.MessageIDError"
//  FAILEDOPERATION_RECEIVEERROR = "FailedOperation.ReceiveError"
//  FAILEDOPERATION_RECEIVETIMEOUT = "FailedOperation.ReceiveTimeout"
//  FAILEDOPERATION_SENDMESSAGETIMEOUTERROR = "FailedOperation.SendMessageTimeoutError"
//  FAILEDOPERATION_TOPICTYPEERROR = "FailedOperation.TopicTypeError"
//  INVALIDPARAMETER_TENANTNOTFOUND = "InvalidParameter.TenantNotFound"
//  INVALIDPARAMETER_TOKENNOTFOUND = "InvalidParameter.TokenNotFound"
//  INVALIDPARAMETERVALUE_TOPICNOTFOUND = "InvalidParameterValue.TopicNotFound"
func (c *Client) SendBatchMessagesWithContext(ctx context.Context, request *SendBatchMessagesRequest) (response *SendBatchMessagesResponse, err error) {
    if request == nil {
        request = NewSendBatchMessagesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SendBatchMessages require credential")
    }

    request.SetContext(ctx)
    
    response = NewSendBatchMessagesResponse()
    err = c.Send(request, response)
    return
}

func NewSendCmqMsgRequest() (request *SendCmqMsgRequest) {
    request = &SendCmqMsgRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "SendCmqMsg")
    
    
    return
}

func NewSendCmqMsgResponse() (response *SendCmqMsgResponse) {
    response = &SendCmqMsgResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// SendCmqMsg
// 发送cmq消息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) SendCmqMsg(request *SendCmqMsgRequest) (response *SendCmqMsgResponse, err error) {
    return c.SendCmqMsgWithContext(context.Background(), request)
}

// SendCmqMsg
// 发送cmq消息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) SendCmqMsgWithContext(ctx context.Context, request *SendCmqMsgRequest) (response *SendCmqMsgResponse, err error) {
    if request == nil {
        request = NewSendCmqMsgRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SendCmqMsg require credential")
    }

    request.SetContext(ctx)
    
    response = NewSendCmqMsgResponse()
    err = c.Send(request, response)
    return
}

func NewSendMessagesRequest() (request *SendMessagesRequest) {
    request = &SendMessagesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "SendMessages")
    
    
    return
}

func NewSendMessagesResponse() (response *SendMessagesResponse) {
    response = &SendMessagesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// SendMessages
// 发送单条消息
//
// 不支持持久topic
//
// 可能返回的错误码:
//  FAILEDOPERATION_CREATEPRODUCERERROR = "FailedOperation.CreateProducerError"
//  FAILEDOPERATION_CREATEPULSARCLIENTERROR = "FailedOperation.CreatePulsarClientError"
//  FAILEDOPERATION_MAXMESSAGESIZEERROR = "FailedOperation.MaxMessageSizeError"
//  FAILEDOPERATION_MESSAGEIDERROR = "FailedOperation.MessageIDError"
//  FAILEDOPERATION_RECEIVEERROR = "FailedOperation.ReceiveError"
//  FAILEDOPERATION_RECEIVETIMEOUT = "FailedOperation.ReceiveTimeout"
//  FAILEDOPERATION_SENDMESSAGETIMEOUTERROR = "FailedOperation.SendMessageTimeoutError"
//  FAILEDOPERATION_TOPICTYPEERROR = "FailedOperation.TopicTypeError"
//  INVALIDPARAMETER_TENANTNOTFOUND = "InvalidParameter.TenantNotFound"
//  INVALIDPARAMETER_TOKENNOTFOUND = "InvalidParameter.TokenNotFound"
//  INVALIDPARAMETERVALUE_TOPICNOTFOUND = "InvalidParameterValue.TopicNotFound"
func (c *Client) SendMessages(request *SendMessagesRequest) (response *SendMessagesResponse, err error) {
    return c.SendMessagesWithContext(context.Background(), request)
}

// SendMessages
// 发送单条消息
//
// 不支持持久topic
//
// 可能返回的错误码:
//  FAILEDOPERATION_CREATEPRODUCERERROR = "FailedOperation.CreateProducerError"
//  FAILEDOPERATION_CREATEPULSARCLIENTERROR = "FailedOperation.CreatePulsarClientError"
//  FAILEDOPERATION_MAXMESSAGESIZEERROR = "FailedOperation.MaxMessageSizeError"
//  FAILEDOPERATION_MESSAGEIDERROR = "FailedOperation.MessageIDError"
//  FAILEDOPERATION_RECEIVEERROR = "FailedOperation.ReceiveError"
//  FAILEDOPERATION_RECEIVETIMEOUT = "FailedOperation.ReceiveTimeout"
//  FAILEDOPERATION_SENDMESSAGETIMEOUTERROR = "FailedOperation.SendMessageTimeoutError"
//  FAILEDOPERATION_TOPICTYPEERROR = "FailedOperation.TopicTypeError"
//  INVALIDPARAMETER_TENANTNOTFOUND = "InvalidParameter.TenantNotFound"
//  INVALIDPARAMETER_TOKENNOTFOUND = "InvalidParameter.TokenNotFound"
//  INVALIDPARAMETERVALUE_TOPICNOTFOUND = "InvalidParameterValue.TopicNotFound"
func (c *Client) SendMessagesWithContext(ctx context.Context, request *SendMessagesRequest) (response *SendMessagesResponse, err error) {
    if request == nil {
        request = NewSendMessagesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SendMessages require credential")
    }

    request.SetContext(ctx)
    
    response = NewSendMessagesResponse()
    err = c.Send(request, response)
    return
}

func NewSendMsgRequest() (request *SendMsgRequest) {
    request = &SendMsgRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "SendMsg")
    
    
    return
}

func NewSendMsgResponse() (response *SendMsgResponse) {
    response = &SendMsgResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// SendMsg
// 此接口仅用于测试发生消息，不能作为现网正式生产使用
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_GETTOPICPARTITIONSFAILED = "FailedOperation.GetTopicPartitionsFailed"
//  FAILEDOPERATION_SENDMSGFAILED = "FailedOperation.SendMsgFailed"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMS = "InvalidParameterValue.InvalidParams"
//  MISSINGPARAMETER_NEEDMOREPARAMS = "MissingParameter.NeedMoreParams"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
//  RESOURCENOTFOUND_TOPIC = "ResourceNotFound.Topic"
//  RESOURCEUNAVAILABLE_FUNDREQUIRED = "ResourceUnavailable.FundRequired"
func (c *Client) SendMsg(request *SendMsgRequest) (response *SendMsgResponse, err error) {
    return c.SendMsgWithContext(context.Background(), request)
}

// SendMsg
// 此接口仅用于测试发生消息，不能作为现网正式生产使用
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_GETTOPICPARTITIONSFAILED = "FailedOperation.GetTopicPartitionsFailed"
//  FAILEDOPERATION_SENDMSGFAILED = "FailedOperation.SendMsgFailed"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMS = "InvalidParameterValue.InvalidParams"
//  MISSINGPARAMETER_NEEDMOREPARAMS = "MissingParameter.NeedMoreParams"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
//  RESOURCENOTFOUND_TOPIC = "ResourceNotFound.Topic"
//  RESOURCEUNAVAILABLE_FUNDREQUIRED = "ResourceUnavailable.FundRequired"
func (c *Client) SendMsgWithContext(ctx context.Context, request *SendMsgRequest) (response *SendMsgResponse, err error) {
    if request == nil {
        request = NewSendMsgRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SendMsg require credential")
    }

    request.SetContext(ctx)
    
    response = NewSendMsgResponse()
    err = c.Send(request, response)
    return
}

func NewSendRocketMQMessageRequest() (request *SendRocketMQMessageRequest) {
    request = &SendRocketMQMessageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "SendRocketMQMessage")
    
    
    return
}

func NewSendRocketMQMessageResponse() (response *SendRocketMQMessageResponse) {
    response = &SendRocketMQMessageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// SendRocketMQMessage
// 发送RocketMQ消息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
//  RESOURCENOTFOUND_TOPIC = "ResourceNotFound.Topic"
func (c *Client) SendRocketMQMessage(request *SendRocketMQMessageRequest) (response *SendRocketMQMessageResponse, err error) {
    return c.SendRocketMQMessageWithContext(context.Background(), request)
}

// SendRocketMQMessage
// 发送RocketMQ消息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
//  RESOURCENOTFOUND_TOPIC = "ResourceNotFound.Topic"
func (c *Client) SendRocketMQMessageWithContext(ctx context.Context, request *SendRocketMQMessageRequest) (response *SendRocketMQMessageResponse, err error) {
    if request == nil {
        request = NewSendRocketMQMessageRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SendRocketMQMessage require credential")
    }

    request.SetContext(ctx)
    
    response = NewSendRocketMQMessageResponse()
    err = c.Send(request, response)
    return
}

func NewUnbindCmqDeadLetterRequest() (request *UnbindCmqDeadLetterRequest) {
    request = &UnbindCmqDeadLetterRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "UnbindCmqDeadLetter")
    
    
    return
}

func NewUnbindCmqDeadLetterResponse() (response *UnbindCmqDeadLetterResponse) {
    response = &UnbindCmqDeadLetterResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UnbindCmqDeadLetter
// 解绑cmq死信队列
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) UnbindCmqDeadLetter(request *UnbindCmqDeadLetterRequest) (response *UnbindCmqDeadLetterResponse, err error) {
    return c.UnbindCmqDeadLetterWithContext(context.Background(), request)
}

// UnbindCmqDeadLetter
// 解绑cmq死信队列
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) UnbindCmqDeadLetterWithContext(ctx context.Context, request *UnbindCmqDeadLetterRequest) (response *UnbindCmqDeadLetterResponse, err error) {
    if request == nil {
        request = NewUnbindCmqDeadLetterRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UnbindCmqDeadLetter require credential")
    }

    request.SetContext(ctx)
    
    response = NewUnbindCmqDeadLetterResponse()
    err = c.Send(request, response)
    return
}
