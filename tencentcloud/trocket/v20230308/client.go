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
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
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
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
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
// 创建 RocketMQ 5.x 集群
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CreateInstance(request *CreateInstanceRequest) (response *CreateInstanceResponse, err error) {
    return c.CreateInstanceWithContext(context.Background(), request)
}

// CreateInstance
// 创建 RocketMQ 5.x 集群
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

func NewCreateMQTTInsPublicEndpointRequest() (request *CreateMQTTInsPublicEndpointRequest) {
    request = &CreateMQTTInsPublicEndpointRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trocket", APIVersion, "CreateMQTTInsPublicEndpoint")
    
    
    return
}

func NewCreateMQTTInsPublicEndpointResponse() (response *CreateMQTTInsPublicEndpointResponse) {
    response = &CreateMQTTInsPublicEndpointResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateMQTTInsPublicEndpoint
// 为MQTT实例创建公网接入点
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CreateMQTTInsPublicEndpoint(request *CreateMQTTInsPublicEndpointRequest) (response *CreateMQTTInsPublicEndpointResponse, err error) {
    return c.CreateMQTTInsPublicEndpointWithContext(context.Background(), request)
}

// CreateMQTTInsPublicEndpoint
// 为MQTT实例创建公网接入点
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CreateMQTTInsPublicEndpointWithContext(ctx context.Context, request *CreateMQTTInsPublicEndpointRequest) (response *CreateMQTTInsPublicEndpointResponse, err error) {
    if request == nil {
        request = NewCreateMQTTInsPublicEndpointRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateMQTTInsPublicEndpoint require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateMQTTInsPublicEndpointResponse()
    err = c.Send(request, response)
    return
}

func NewCreateMQTTInstanceRequest() (request *CreateMQTTInstanceRequest) {
    request = &CreateMQTTInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trocket", APIVersion, "CreateMQTTInstance")
    
    
    return
}

func NewCreateMQTTInstanceResponse() (response *CreateMQTTInstanceResponse) {
    response = &CreateMQTTInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateMQTTInstance
// 购买新的MQTT实例
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CreateMQTTInstance(request *CreateMQTTInstanceRequest) (response *CreateMQTTInstanceResponse, err error) {
    return c.CreateMQTTInstanceWithContext(context.Background(), request)
}

// CreateMQTTInstance
// 购买新的MQTT实例
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CreateMQTTInstanceWithContext(ctx context.Context, request *CreateMQTTInstanceRequest) (response *CreateMQTTInstanceResponse, err error) {
    if request == nil {
        request = NewCreateMQTTInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateMQTTInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateMQTTInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewCreateMQTTTopicRequest() (request *CreateMQTTTopicRequest) {
    request = &CreateMQTTTopicRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trocket", APIVersion, "CreateMQTTTopic")
    
    
    return
}

func NewCreateMQTTTopicResponse() (response *CreateMQTTTopicResponse) {
    response = &CreateMQTTTopicResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateMQTTTopic
// 创建主题
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INSTANCENOTREADY = "FailedOperation.InstanceNotReady"
//  LIMITEXCEEDED_TOPICNUM = "LimitExceeded.TopicNum"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
//  UNSUPPORTEDOPERATION_RESOURCEALREADYEXISTS = "UnsupportedOperation.ResourceAlreadyExists"
func (c *Client) CreateMQTTTopic(request *CreateMQTTTopicRequest) (response *CreateMQTTTopicResponse, err error) {
    return c.CreateMQTTTopicWithContext(context.Background(), request)
}

// CreateMQTTTopic
// 创建主题
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INSTANCENOTREADY = "FailedOperation.InstanceNotReady"
//  LIMITEXCEEDED_TOPICNUM = "LimitExceeded.TopicNum"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
//  UNSUPPORTEDOPERATION_RESOURCEALREADYEXISTS = "UnsupportedOperation.ResourceAlreadyExists"
func (c *Client) CreateMQTTTopicWithContext(ctx context.Context, request *CreateMQTTTopicRequest) (response *CreateMQTTTopicResponse, err error) {
    if request == nil {
        request = NewCreateMQTTTopicRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateMQTTTopic require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateMQTTTopicResponse()
    err = c.Send(request, response)
    return
}

func NewCreateMQTTUserRequest() (request *CreateMQTTUserRequest) {
    request = &CreateMQTTUserRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trocket", APIVersion, "CreateMQTTUser")
    
    
    return
}

func NewCreateMQTTUserResponse() (response *CreateMQTTUserResponse) {
    response = &CreateMQTTUserResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateMQTTUser
// 添加mqtt角色
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INSTANCENOTREADY = "FailedOperation.InstanceNotReady"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
//  UNSUPPORTEDOPERATION_RESOURCEALREADYEXISTS = "UnsupportedOperation.ResourceAlreadyExists"
func (c *Client) CreateMQTTUser(request *CreateMQTTUserRequest) (response *CreateMQTTUserResponse, err error) {
    return c.CreateMQTTUserWithContext(context.Background(), request)
}

// CreateMQTTUser
// 添加mqtt角色
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INSTANCENOTREADY = "FailedOperation.InstanceNotReady"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
//  UNSUPPORTEDOPERATION_RESOURCEALREADYEXISTS = "UnsupportedOperation.ResourceAlreadyExists"
func (c *Client) CreateMQTTUserWithContext(ctx context.Context, request *CreateMQTTUserRequest) (response *CreateMQTTUserResponse, err error) {
    if request == nil {
        request = NewCreateMQTTUserRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateMQTTUser require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateMQTTUserResponse()
    err = c.Send(request, response)
    return
}

func NewCreateRoleRequest() (request *CreateRoleRequest) {
    request = &CreateRoleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trocket", APIVersion, "CreateRole")
    
    
    return
}

func NewCreateRoleResponse() (response *CreateRoleResponse) {
    response = &CreateRoleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateRole
// 添加角色
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INSTANCENOTREADY = "FailedOperation.InstanceNotReady"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
//  UNSUPPORTEDOPERATION_RESOURCEALREADYEXISTS = "UnsupportedOperation.ResourceAlreadyExists"
func (c *Client) CreateRole(request *CreateRoleRequest) (response *CreateRoleResponse, err error) {
    return c.CreateRoleWithContext(context.Background(), request)
}

// CreateRole
// 添加角色
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INSTANCENOTREADY = "FailedOperation.InstanceNotReady"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
//  UNSUPPORTEDOPERATION_RESOURCEALREADYEXISTS = "UnsupportedOperation.ResourceAlreadyExists"
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

func NewDeleteMQTTInsPublicEndpointRequest() (request *DeleteMQTTInsPublicEndpointRequest) {
    request = &DeleteMQTTInsPublicEndpointRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trocket", APIVersion, "DeleteMQTTInsPublicEndpoint")
    
    
    return
}

func NewDeleteMQTTInsPublicEndpointResponse() (response *DeleteMQTTInsPublicEndpointResponse) {
    response = &DeleteMQTTInsPublicEndpointResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteMQTTInsPublicEndpoint
// 删除MQTT实例的公网接入点
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
func (c *Client) DeleteMQTTInsPublicEndpoint(request *DeleteMQTTInsPublicEndpointRequest) (response *DeleteMQTTInsPublicEndpointResponse, err error) {
    return c.DeleteMQTTInsPublicEndpointWithContext(context.Background(), request)
}

// DeleteMQTTInsPublicEndpoint
// 删除MQTT实例的公网接入点
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
func (c *Client) DeleteMQTTInsPublicEndpointWithContext(ctx context.Context, request *DeleteMQTTInsPublicEndpointRequest) (response *DeleteMQTTInsPublicEndpointResponse, err error) {
    if request == nil {
        request = NewDeleteMQTTInsPublicEndpointRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteMQTTInsPublicEndpoint require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteMQTTInsPublicEndpointResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteMQTTInstanceRequest() (request *DeleteMQTTInstanceRequest) {
    request = &DeleteMQTTInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trocket", APIVersion, "DeleteMQTTInstance")
    
    
    return
}

func NewDeleteMQTTInstanceResponse() (response *DeleteMQTTInstanceResponse) {
    response = &DeleteMQTTInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteMQTTInstance
// 删除MQTT实例
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
func (c *Client) DeleteMQTTInstance(request *DeleteMQTTInstanceRequest) (response *DeleteMQTTInstanceResponse, err error) {
    return c.DeleteMQTTInstanceWithContext(context.Background(), request)
}

// DeleteMQTTInstance
// 删除MQTT实例
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
func (c *Client) DeleteMQTTInstanceWithContext(ctx context.Context, request *DeleteMQTTInstanceRequest) (response *DeleteMQTTInstanceResponse, err error) {
    if request == nil {
        request = NewDeleteMQTTInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteMQTTInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteMQTTInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteMQTTTopicRequest() (request *DeleteMQTTTopicRequest) {
    request = &DeleteMQTTTopicRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trocket", APIVersion, "DeleteMQTTTopic")
    
    
    return
}

func NewDeleteMQTTTopicResponse() (response *DeleteMQTTTopicResponse) {
    response = &DeleteMQTTTopicResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteMQTTTopic
// 删除MQTT主题
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCENOTFOUND_TOPIC = "ResourceNotFound.Topic"
func (c *Client) DeleteMQTTTopic(request *DeleteMQTTTopicRequest) (response *DeleteMQTTTopicResponse, err error) {
    return c.DeleteMQTTTopicWithContext(context.Background(), request)
}

// DeleteMQTTTopic
// 删除MQTT主题
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCENOTFOUND_TOPIC = "ResourceNotFound.Topic"
func (c *Client) DeleteMQTTTopicWithContext(ctx context.Context, request *DeleteMQTTTopicRequest) (response *DeleteMQTTTopicResponse, err error) {
    if request == nil {
        request = NewDeleteMQTTTopicRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteMQTTTopic require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteMQTTTopicResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteMQTTUserRequest() (request *DeleteMQTTUserRequest) {
    request = &DeleteMQTTUserRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trocket", APIVersion, "DeleteMQTTUser")
    
    
    return
}

func NewDeleteMQTTUserResponse() (response *DeleteMQTTUserResponse) {
    response = &DeleteMQTTUserResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteMQTTUser
// 删除MQTT访问用户
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
//  RESOURCENOTFOUND_ROLE = "ResourceNotFound.Role"
func (c *Client) DeleteMQTTUser(request *DeleteMQTTUserRequest) (response *DeleteMQTTUserResponse, err error) {
    return c.DeleteMQTTUserWithContext(context.Background(), request)
}

// DeleteMQTTUser
// 删除MQTT访问用户
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
//  RESOURCENOTFOUND_ROLE = "ResourceNotFound.Role"
func (c *Client) DeleteMQTTUserWithContext(ctx context.Context, request *DeleteMQTTUserRequest) (response *DeleteMQTTUserResponse, err error) {
    if request == nil {
        request = NewDeleteMQTTUserRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteMQTTUser require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteMQTTUserResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteRoleRequest() (request *DeleteRoleRequest) {
    request = &DeleteRoleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trocket", APIVersion, "DeleteRole")
    
    
    return
}

func NewDeleteRoleResponse() (response *DeleteRoleResponse) {
    response = &DeleteRoleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteRole
// 删除角色
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
//  RESOURCENOTFOUND_ROLE = "ResourceNotFound.Role"
func (c *Client) DeleteRole(request *DeleteRoleRequest) (response *DeleteRoleResponse, err error) {
    return c.DeleteRoleWithContext(context.Background(), request)
}

// DeleteRole
// 删除角色
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
//  RESOURCENOTFOUND_ROLE = "ResourceNotFound.Role"
func (c *Client) DeleteRoleWithContext(ctx context.Context, request *DeleteRoleRequest) (response *DeleteRoleResponse, err error) {
    if request == nil {
        request = NewDeleteRoleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteRole require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteRoleResponse()
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
//  RESOURCENOTFOUND_TOPIC = "ResourceNotFound.Topic"
func (c *Client) DeleteTopic(request *DeleteTopicRequest) (response *DeleteTopicResponse, err error) {
    return c.DeleteTopicWithContext(context.Background(), request)
}

// DeleteTopic
// 删除主题
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCENOTFOUND_TOPIC = "ResourceNotFound.Topic"
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
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
func (c *Client) DescribeConsumerGroup(request *DescribeConsumerGroupRequest) (response *DescribeConsumerGroupResponse, err error) {
    return c.DescribeConsumerGroupWithContext(context.Background(), request)
}

// DescribeConsumerGroup
// 查询消费组详情
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_GROUP = "ResourceNotFound.Group"
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
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

func NewDescribeConsumerGroupListRequest() (request *DescribeConsumerGroupListRequest) {
    request = &DescribeConsumerGroupListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trocket", APIVersion, "DescribeConsumerGroupList")
    
    
    return
}

func NewDescribeConsumerGroupListResponse() (response *DescribeConsumerGroupListResponse) {
    response = &DescribeConsumerGroupListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeConsumerGroupList
// 获取消费组列表，Filter参数使用说明如下：
//
// 
//
// 1. ConsumerGroupName，名称模糊查询
//
// 2. ConsumeMessageOrderly，投递顺序性。"true":顺序投递；"false":并发投递
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
func (c *Client) DescribeConsumerGroupList(request *DescribeConsumerGroupListRequest) (response *DescribeConsumerGroupListResponse, err error) {
    return c.DescribeConsumerGroupListWithContext(context.Background(), request)
}

// DescribeConsumerGroupList
// 获取消费组列表，Filter参数使用说明如下：
//
// 
//
// 1. ConsumerGroupName，名称模糊查询
//
// 2. ConsumeMessageOrderly，投递顺序性。"true":顺序投递；"false":并发投递
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
func (c *Client) DescribeConsumerGroupListWithContext(ctx context.Context, request *DescribeConsumerGroupListRequest) (response *DescribeConsumerGroupListResponse, err error) {
    if request == nil {
        request = NewDescribeConsumerGroupListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeConsumerGroupList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeConsumerGroupListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeFusionInstanceListRequest() (request *DescribeFusionInstanceListRequest) {
    request = &DescribeFusionInstanceListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trocket", APIVersion, "DescribeFusionInstanceList")
    
    
    return
}

func NewDescribeFusionInstanceListResponse() (response *DescribeFusionInstanceListResponse) {
    response = &DescribeFusionInstanceListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeFusionInstanceList
// 获取实例列表，Filters参数使用说明如下：
//
// 1. InstanceName, 名称模糊查询
//
// 2. InstanceId，实例ID查询
//
// 3. InstanceType, 实例类型查询，支持多选
//
// 4. Version，实例版本查询
//
// 当使用TagFilters查询时，Filters参数失效。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeFusionInstanceList(request *DescribeFusionInstanceListRequest) (response *DescribeFusionInstanceListResponse, err error) {
    return c.DescribeFusionInstanceListWithContext(context.Background(), request)
}

// DescribeFusionInstanceList
// 获取实例列表，Filters参数使用说明如下：
//
// 1. InstanceName, 名称模糊查询
//
// 2. InstanceId，实例ID查询
//
// 3. InstanceType, 实例类型查询，支持多选
//
// 4. Version，实例版本查询
//
// 当使用TagFilters查询时，Filters参数失效。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeFusionInstanceListWithContext(ctx context.Context, request *DescribeFusionInstanceListRequest) (response *DescribeFusionInstanceListResponse, err error) {
    if request == nil {
        request = NewDescribeFusionInstanceListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeFusionInstanceList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeFusionInstanceListResponse()
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

func NewDescribeMQTTClientRequest() (request *DescribeMQTTClientRequest) {
    request = &DescribeMQTTClientRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trocket", APIVersion, "DescribeMQTTClient")
    
    
    return
}

func NewDescribeMQTTClientResponse() (response *DescribeMQTTClientResponse) {
    response = &DescribeMQTTClientResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeMQTTClient
// 查询 MQTT 客户端详情
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_CLIENT = "ResourceNotFound.Client"
func (c *Client) DescribeMQTTClient(request *DescribeMQTTClientRequest) (response *DescribeMQTTClientResponse, err error) {
    return c.DescribeMQTTClientWithContext(context.Background(), request)
}

// DescribeMQTTClient
// 查询 MQTT 客户端详情
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_CLIENT = "ResourceNotFound.Client"
func (c *Client) DescribeMQTTClientWithContext(ctx context.Context, request *DescribeMQTTClientRequest) (response *DescribeMQTTClientResponse, err error) {
    if request == nil {
        request = NewDescribeMQTTClientRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMQTTClient require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMQTTClientResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMQTTInsPublicEndpointsRequest() (request *DescribeMQTTInsPublicEndpointsRequest) {
    request = &DescribeMQTTInsPublicEndpointsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trocket", APIVersion, "DescribeMQTTInsPublicEndpoints")
    
    
    return
}

func NewDescribeMQTTInsPublicEndpointsResponse() (response *DescribeMQTTInsPublicEndpointsResponse) {
    response = &DescribeMQTTInsPublicEndpointsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeMQTTInsPublicEndpoints
// 查询MQTT实例公网接入点
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_CLIENT = "ResourceNotFound.Client"
func (c *Client) DescribeMQTTInsPublicEndpoints(request *DescribeMQTTInsPublicEndpointsRequest) (response *DescribeMQTTInsPublicEndpointsResponse, err error) {
    return c.DescribeMQTTInsPublicEndpointsWithContext(context.Background(), request)
}

// DescribeMQTTInsPublicEndpoints
// 查询MQTT实例公网接入点
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_CLIENT = "ResourceNotFound.Client"
func (c *Client) DescribeMQTTInsPublicEndpointsWithContext(ctx context.Context, request *DescribeMQTTInsPublicEndpointsRequest) (response *DescribeMQTTInsPublicEndpointsResponse, err error) {
    if request == nil {
        request = NewDescribeMQTTInsPublicEndpointsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMQTTInsPublicEndpoints require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMQTTInsPublicEndpointsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMQTTInsVPCEndpointsRequest() (request *DescribeMQTTInsVPCEndpointsRequest) {
    request = &DescribeMQTTInsVPCEndpointsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trocket", APIVersion, "DescribeMQTTInsVPCEndpoints")
    
    
    return
}

func NewDescribeMQTTInsVPCEndpointsResponse() (response *DescribeMQTTInsVPCEndpointsResponse) {
    response = &DescribeMQTTInsVPCEndpointsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeMQTTInsVPCEndpoints
// 查询MQTT实例公网接入点
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_CLIENT = "ResourceNotFound.Client"
func (c *Client) DescribeMQTTInsVPCEndpoints(request *DescribeMQTTInsVPCEndpointsRequest) (response *DescribeMQTTInsVPCEndpointsResponse, err error) {
    return c.DescribeMQTTInsVPCEndpointsWithContext(context.Background(), request)
}

// DescribeMQTTInsVPCEndpoints
// 查询MQTT实例公网接入点
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_CLIENT = "ResourceNotFound.Client"
func (c *Client) DescribeMQTTInsVPCEndpointsWithContext(ctx context.Context, request *DescribeMQTTInsVPCEndpointsRequest) (response *DescribeMQTTInsVPCEndpointsResponse, err error) {
    if request == nil {
        request = NewDescribeMQTTInsVPCEndpointsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMQTTInsVPCEndpoints require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMQTTInsVPCEndpointsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMQTTInstanceRequest() (request *DescribeMQTTInstanceRequest) {
    request = &DescribeMQTTInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trocket", APIVersion, "DescribeMQTTInstance")
    
    
    return
}

func NewDescribeMQTTInstanceResponse() (response *DescribeMQTTInstanceResponse) {
    response = &DescribeMQTTInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeMQTTInstance
// 查询实例信息
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
func (c *Client) DescribeMQTTInstance(request *DescribeMQTTInstanceRequest) (response *DescribeMQTTInstanceResponse, err error) {
    return c.DescribeMQTTInstanceWithContext(context.Background(), request)
}

// DescribeMQTTInstance
// 查询实例信息
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
func (c *Client) DescribeMQTTInstanceWithContext(ctx context.Context, request *DescribeMQTTInstanceRequest) (response *DescribeMQTTInstanceResponse, err error) {
    if request == nil {
        request = NewDescribeMQTTInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMQTTInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMQTTInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMQTTInstanceCertRequest() (request *DescribeMQTTInstanceCertRequest) {
    request = &DescribeMQTTInstanceCertRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trocket", APIVersion, "DescribeMQTTInstanceCert")
    
    
    return
}

func NewDescribeMQTTInstanceCertResponse() (response *DescribeMQTTInstanceCertResponse) {
    response = &DescribeMQTTInstanceCertResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeMQTTInstanceCert
// 查询MQTT集群证书列表
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
func (c *Client) DescribeMQTTInstanceCert(request *DescribeMQTTInstanceCertRequest) (response *DescribeMQTTInstanceCertResponse, err error) {
    return c.DescribeMQTTInstanceCertWithContext(context.Background(), request)
}

// DescribeMQTTInstanceCert
// 查询MQTT集群证书列表
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
func (c *Client) DescribeMQTTInstanceCertWithContext(ctx context.Context, request *DescribeMQTTInstanceCertRequest) (response *DescribeMQTTInstanceCertResponse, err error) {
    if request == nil {
        request = NewDescribeMQTTInstanceCertRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMQTTInstanceCert require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMQTTInstanceCertResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMQTTInstanceListRequest() (request *DescribeMQTTInstanceListRequest) {
    request = &DescribeMQTTInstanceListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trocket", APIVersion, "DescribeMQTTInstanceList")
    
    
    return
}

func NewDescribeMQTTInstanceListResponse() (response *DescribeMQTTInstanceListResponse) {
    response = &DescribeMQTTInstanceListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeMQTTInstanceList
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
func (c *Client) DescribeMQTTInstanceList(request *DescribeMQTTInstanceListRequest) (response *DescribeMQTTInstanceListResponse, err error) {
    return c.DescribeMQTTInstanceListWithContext(context.Background(), request)
}

// DescribeMQTTInstanceList
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
func (c *Client) DescribeMQTTInstanceListWithContext(ctx context.Context, request *DescribeMQTTInstanceListRequest) (response *DescribeMQTTInstanceListResponse, err error) {
    if request == nil {
        request = NewDescribeMQTTInstanceListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMQTTInstanceList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMQTTInstanceListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMQTTMessageRequest() (request *DescribeMQTTMessageRequest) {
    request = &DescribeMQTTMessageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trocket", APIVersion, "DescribeMQTTMessage")
    
    
    return
}

func NewDescribeMQTTMessageResponse() (response *DescribeMQTTMessageResponse) {
    response = &DescribeMQTTMessageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeMQTTMessage
// 查询MQTT消息详情
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
//  RESOURCENOTFOUND_MESSAGE = "ResourceNotFound.Message"
func (c *Client) DescribeMQTTMessage(request *DescribeMQTTMessageRequest) (response *DescribeMQTTMessageResponse, err error) {
    return c.DescribeMQTTMessageWithContext(context.Background(), request)
}

// DescribeMQTTMessage
// 查询MQTT消息详情
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
//  RESOURCENOTFOUND_MESSAGE = "ResourceNotFound.Message"
func (c *Client) DescribeMQTTMessageWithContext(ctx context.Context, request *DescribeMQTTMessageRequest) (response *DescribeMQTTMessageResponse, err error) {
    if request == nil {
        request = NewDescribeMQTTMessageRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMQTTMessage require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMQTTMessageResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMQTTMessageListRequest() (request *DescribeMQTTMessageListRequest) {
    request = &DescribeMQTTMessageListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trocket", APIVersion, "DescribeMQTTMessageList")
    
    
    return
}

func NewDescribeMQTTMessageListResponse() (response *DescribeMQTTMessageListResponse) {
    response = &DescribeMQTTMessageListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeMQTTMessageList
// 查询消息列表，如查询死信，请设置ConsumerGroup参数
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
func (c *Client) DescribeMQTTMessageList(request *DescribeMQTTMessageListRequest) (response *DescribeMQTTMessageListResponse, err error) {
    return c.DescribeMQTTMessageListWithContext(context.Background(), request)
}

// DescribeMQTTMessageList
// 查询消息列表，如查询死信，请设置ConsumerGroup参数
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
func (c *Client) DescribeMQTTMessageListWithContext(ctx context.Context, request *DescribeMQTTMessageListRequest) (response *DescribeMQTTMessageListResponse, err error) {
    if request == nil {
        request = NewDescribeMQTTMessageListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMQTTMessageList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMQTTMessageListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMQTTProductSKUListRequest() (request *DescribeMQTTProductSKUListRequest) {
    request = &DescribeMQTTProductSKUListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trocket", APIVersion, "DescribeMQTTProductSKUList")
    
    
    return
}

func NewDescribeMQTTProductSKUListResponse() (response *DescribeMQTTProductSKUListResponse) {
    response = &DescribeMQTTProductSKUListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeMQTTProductSKUList
// 获取产品售卖规格
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
func (c *Client) DescribeMQTTProductSKUList(request *DescribeMQTTProductSKUListRequest) (response *DescribeMQTTProductSKUListResponse, err error) {
    return c.DescribeMQTTProductSKUListWithContext(context.Background(), request)
}

// DescribeMQTTProductSKUList
// 获取产品售卖规格
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
func (c *Client) DescribeMQTTProductSKUListWithContext(ctx context.Context, request *DescribeMQTTProductSKUListRequest) (response *DescribeMQTTProductSKUListResponse, err error) {
    if request == nil {
        request = NewDescribeMQTTProductSKUListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMQTTProductSKUList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMQTTProductSKUListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMQTTTopicRequest() (request *DescribeMQTTTopicRequest) {
    request = &DescribeMQTTTopicRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trocket", APIVersion, "DescribeMQTTTopic")
    
    
    return
}

func NewDescribeMQTTTopicResponse() (response *DescribeMQTTTopicResponse) {
    response = &DescribeMQTTTopicResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeMQTTTopic
// 查询mqtt主题详情
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
//  RESOURCENOTFOUND_TOPIC = "ResourceNotFound.Topic"
func (c *Client) DescribeMQTTTopic(request *DescribeMQTTTopicRequest) (response *DescribeMQTTTopicResponse, err error) {
    return c.DescribeMQTTTopicWithContext(context.Background(), request)
}

// DescribeMQTTTopic
// 查询mqtt主题详情
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
//  RESOURCENOTFOUND_TOPIC = "ResourceNotFound.Topic"
func (c *Client) DescribeMQTTTopicWithContext(ctx context.Context, request *DescribeMQTTTopicRequest) (response *DescribeMQTTTopicResponse, err error) {
    if request == nil {
        request = NewDescribeMQTTTopicRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMQTTTopic require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMQTTTopicResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMQTTTopicListRequest() (request *DescribeMQTTTopicListRequest) {
    request = &DescribeMQTTTopicListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trocket", APIVersion, "DescribeMQTTTopicList")
    
    
    return
}

func NewDescribeMQTTTopicListResponse() (response *DescribeMQTTTopicListResponse) {
    response = &DescribeMQTTTopicListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeMQTTTopicList
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
func (c *Client) DescribeMQTTTopicList(request *DescribeMQTTTopicListRequest) (response *DescribeMQTTTopicListResponse, err error) {
    return c.DescribeMQTTTopicListWithContext(context.Background(), request)
}

// DescribeMQTTTopicList
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
func (c *Client) DescribeMQTTTopicListWithContext(ctx context.Context, request *DescribeMQTTTopicListRequest) (response *DescribeMQTTTopicListResponse, err error) {
    if request == nil {
        request = NewDescribeMQTTTopicListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMQTTTopicList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMQTTTopicListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMQTTUserListRequest() (request *DescribeMQTTUserListRequest) {
    request = &DescribeMQTTUserListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trocket", APIVersion, "DescribeMQTTUserList")
    
    
    return
}

func NewDescribeMQTTUserListResponse() (response *DescribeMQTTUserListResponse) {
    response = &DescribeMQTTUserListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeMQTTUserList
// 查询用户列表，Filter参数使用说明如下：
//
// 
//
// 1. Username，用户名称模糊搜索
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
func (c *Client) DescribeMQTTUserList(request *DescribeMQTTUserListRequest) (response *DescribeMQTTUserListResponse, err error) {
    return c.DescribeMQTTUserListWithContext(context.Background(), request)
}

// DescribeMQTTUserList
// 查询用户列表，Filter参数使用说明如下：
//
// 
//
// 1. Username，用户名称模糊搜索
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
func (c *Client) DescribeMQTTUserListWithContext(ctx context.Context, request *DescribeMQTTUserListRequest) (response *DescribeMQTTUserListResponse, err error) {
    if request == nil {
        request = NewDescribeMQTTUserListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMQTTUserList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMQTTUserListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRoleListRequest() (request *DescribeRoleListRequest) {
    request = &DescribeRoleListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trocket", APIVersion, "DescribeRoleList")
    
    
    return
}

func NewDescribeRoleListResponse() (response *DescribeRoleListResponse) {
    response = &DescribeRoleListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRoleList
// 查询角色列表，Filter参数使用说明如下：
//
// 
//
// 1. RoleName，角色名称模糊搜索
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
func (c *Client) DescribeRoleList(request *DescribeRoleListRequest) (response *DescribeRoleListResponse, err error) {
    return c.DescribeRoleListWithContext(context.Background(), request)
}

// DescribeRoleList
// 查询角色列表，Filter参数使用说明如下：
//
// 
//
// 1. RoleName，角色名称模糊搜索
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
func (c *Client) DescribeRoleListWithContext(ctx context.Context, request *DescribeRoleListRequest) (response *DescribeRoleListResponse, err error) {
    if request == nil {
        request = NewDescribeRoleListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRoleList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRoleListResponse()
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

func NewImportSourceClusterConsumerGroupsRequest() (request *ImportSourceClusterConsumerGroupsRequest) {
    request = &ImportSourceClusterConsumerGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trocket", APIVersion, "ImportSourceClusterConsumerGroups")
    
    
    return
}

func NewImportSourceClusterConsumerGroupsResponse() (response *ImportSourceClusterConsumerGroupsResponse) {
    response = &ImportSourceClusterConsumerGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ImportSourceClusterConsumerGroups
// 导入消费者组列表
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
func (c *Client) ImportSourceClusterConsumerGroups(request *ImportSourceClusterConsumerGroupsRequest) (response *ImportSourceClusterConsumerGroupsResponse, err error) {
    return c.ImportSourceClusterConsumerGroupsWithContext(context.Background(), request)
}

// ImportSourceClusterConsumerGroups
// 导入消费者组列表
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
func (c *Client) ImportSourceClusterConsumerGroupsWithContext(ctx context.Context, request *ImportSourceClusterConsumerGroupsRequest) (response *ImportSourceClusterConsumerGroupsResponse, err error) {
    if request == nil {
        request = NewImportSourceClusterConsumerGroupsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ImportSourceClusterConsumerGroups require credential")
    }

    request.SetContext(ctx)
    
    response = NewImportSourceClusterConsumerGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewImportSourceClusterTopicsRequest() (request *ImportSourceClusterTopicsRequest) {
    request = &ImportSourceClusterTopicsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trocket", APIVersion, "ImportSourceClusterTopics")
    
    
    return
}

func NewImportSourceClusterTopicsResponse() (response *ImportSourceClusterTopicsResponse) {
    response = &ImportSourceClusterTopicsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ImportSourceClusterTopics
// 导入topic列表
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
func (c *Client) ImportSourceClusterTopics(request *ImportSourceClusterTopicsRequest) (response *ImportSourceClusterTopicsResponse, err error) {
    return c.ImportSourceClusterTopicsWithContext(context.Background(), request)
}

// ImportSourceClusterTopics
// 导入topic列表
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
func (c *Client) ImportSourceClusterTopicsWithContext(ctx context.Context, request *ImportSourceClusterTopicsRequest) (response *ImportSourceClusterTopicsResponse, err error) {
    if request == nil {
        request = NewImportSourceClusterTopicsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ImportSourceClusterTopics require credential")
    }

    request.SetContext(ctx)
    
    response = NewImportSourceClusterTopicsResponse()
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
//  UNSUPPORTEDOPERATION_INSTANCETOPICNUMDOWNGRADE = "UnsupportedOperation.InstanceTopicNumDowngrade"
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
//  UNSUPPORTEDOPERATION_INSTANCETOPICNUMDOWNGRADE = "UnsupportedOperation.InstanceTopicNumDowngrade"
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

func NewModifyMQTTInsPublicEndpointRequest() (request *ModifyMQTTInsPublicEndpointRequest) {
    request = &ModifyMQTTInsPublicEndpointRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trocket", APIVersion, "ModifyMQTTInsPublicEndpoint")
    
    
    return
}

func NewModifyMQTTInsPublicEndpointResponse() (response *ModifyMQTTInsPublicEndpointResponse) {
    response = &ModifyMQTTInsPublicEndpointResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyMQTTInsPublicEndpoint
// 更新MQTT实例公网接入点
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INSTANCENOTREADY = "FailedOperation.InstanceNotReady"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNSUPPORTEDOPERATION_INSTANCETOPICNUMDOWNGRADE = "UnsupportedOperation.InstanceTopicNumDowngrade"
func (c *Client) ModifyMQTTInsPublicEndpoint(request *ModifyMQTTInsPublicEndpointRequest) (response *ModifyMQTTInsPublicEndpointResponse, err error) {
    return c.ModifyMQTTInsPublicEndpointWithContext(context.Background(), request)
}

// ModifyMQTTInsPublicEndpoint
// 更新MQTT实例公网接入点
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INSTANCENOTREADY = "FailedOperation.InstanceNotReady"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNSUPPORTEDOPERATION_INSTANCETOPICNUMDOWNGRADE = "UnsupportedOperation.InstanceTopicNumDowngrade"
func (c *Client) ModifyMQTTInsPublicEndpointWithContext(ctx context.Context, request *ModifyMQTTInsPublicEndpointRequest) (response *ModifyMQTTInsPublicEndpointResponse, err error) {
    if request == nil {
        request = NewModifyMQTTInsPublicEndpointRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyMQTTInsPublicEndpoint require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyMQTTInsPublicEndpointResponse()
    err = c.Send(request, response)
    return
}

func NewModifyMQTTInstanceRequest() (request *ModifyMQTTInstanceRequest) {
    request = &ModifyMQTTInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trocket", APIVersion, "ModifyMQTTInstance")
    
    
    return
}

func NewModifyMQTTInstanceResponse() (response *ModifyMQTTInstanceResponse) {
    response = &ModifyMQTTInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyMQTTInstance
// 修改实例属性
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INSTANCENOTREADY = "FailedOperation.InstanceNotReady"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) ModifyMQTTInstance(request *ModifyMQTTInstanceRequest) (response *ModifyMQTTInstanceResponse, err error) {
    return c.ModifyMQTTInstanceWithContext(context.Background(), request)
}

// ModifyMQTTInstance
// 修改实例属性
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INSTANCENOTREADY = "FailedOperation.InstanceNotReady"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) ModifyMQTTInstanceWithContext(ctx context.Context, request *ModifyMQTTInstanceRequest) (response *ModifyMQTTInstanceResponse, err error) {
    if request == nil {
        request = NewModifyMQTTInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyMQTTInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyMQTTInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewModifyMQTTInstanceCertBindingRequest() (request *ModifyMQTTInstanceCertBindingRequest) {
    request = &ModifyMQTTInstanceCertBindingRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trocket", APIVersion, "ModifyMQTTInstanceCertBinding")
    
    
    return
}

func NewModifyMQTTInstanceCertBindingResponse() (response *ModifyMQTTInstanceCertBindingResponse) {
    response = &ModifyMQTTInstanceCertBindingResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyMQTTInstanceCertBinding
// 更新MQTT集群绑定证书
//
// 参数传空，则为删除证书
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INSTANCENOTREADY = "FailedOperation.InstanceNotReady"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) ModifyMQTTInstanceCertBinding(request *ModifyMQTTInstanceCertBindingRequest) (response *ModifyMQTTInstanceCertBindingResponse, err error) {
    return c.ModifyMQTTInstanceCertBindingWithContext(context.Background(), request)
}

// ModifyMQTTInstanceCertBinding
// 更新MQTT集群绑定证书
//
// 参数传空，则为删除证书
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INSTANCENOTREADY = "FailedOperation.InstanceNotReady"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) ModifyMQTTInstanceCertBindingWithContext(ctx context.Context, request *ModifyMQTTInstanceCertBindingRequest) (response *ModifyMQTTInstanceCertBindingResponse, err error) {
    if request == nil {
        request = NewModifyMQTTInstanceCertBindingRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyMQTTInstanceCertBinding require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyMQTTInstanceCertBindingResponse()
    err = c.Send(request, response)
    return
}

func NewModifyMQTTTopicRequest() (request *ModifyMQTTTopicRequest) {
    request = &ModifyMQTTTopicRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trocket", APIVersion, "ModifyMQTTTopic")
    
    
    return
}

func NewModifyMQTTTopicResponse() (response *ModifyMQTTTopicResponse) {
    response = &ModifyMQTTTopicResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyMQTTTopic
// 修改主题属性
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ModifyMQTTTopic(request *ModifyMQTTTopicRequest) (response *ModifyMQTTTopicResponse, err error) {
    return c.ModifyMQTTTopicWithContext(context.Background(), request)
}

// ModifyMQTTTopic
// 修改主题属性
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ModifyMQTTTopicWithContext(ctx context.Context, request *ModifyMQTTTopicRequest) (response *ModifyMQTTTopicResponse, err error) {
    if request == nil {
        request = NewModifyMQTTTopicRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyMQTTTopic require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyMQTTTopicResponse()
    err = c.Send(request, response)
    return
}

func NewModifyMQTTUserRequest() (request *ModifyMQTTUserRequest) {
    request = &ModifyMQTTUserRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trocket", APIVersion, "ModifyMQTTUser")
    
    
    return
}

func NewModifyMQTTUserResponse() (response *ModifyMQTTUserResponse) {
    response = &ModifyMQTTUserResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyMQTTUser
// 修改MQTT角色
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ModifyMQTTUser(request *ModifyMQTTUserRequest) (response *ModifyMQTTUserResponse, err error) {
    return c.ModifyMQTTUserWithContext(context.Background(), request)
}

// ModifyMQTTUser
// 修改MQTT角色
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ModifyMQTTUserWithContext(ctx context.Context, request *ModifyMQTTUserRequest) (response *ModifyMQTTUserResponse, err error) {
    if request == nil {
        request = NewModifyMQTTUserRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyMQTTUser require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyMQTTUserResponse()
    err = c.Send(request, response)
    return
}

func NewModifyRoleRequest() (request *ModifyRoleRequest) {
    request = &ModifyRoleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trocket", APIVersion, "ModifyRole")
    
    
    return
}

func NewModifyRoleResponse() (response *ModifyRoleResponse) {
    response = &ModifyRoleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyRole
// 修改角色
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ModifyRole(request *ModifyRoleRequest) (response *ModifyRoleResponse, err error) {
    return c.ModifyRoleWithContext(context.Background(), request)
}

// ModifyRole
// 修改角色
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
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
