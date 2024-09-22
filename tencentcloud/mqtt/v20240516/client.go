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

package v20240516

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2024-05-16"

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


func NewCreateAuthorizationPolicyRequest() (request *CreateAuthorizationPolicyRequest) {
    request = &CreateAuthorizationPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mqtt", APIVersion, "CreateAuthorizationPolicy")
    
    
    return
}

func NewCreateAuthorizationPolicyResponse() (response *CreateAuthorizationPolicyResponse) {
    response = &CreateAuthorizationPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateAuthorizationPolicy
// 创建MQTT实例的性能测试任务
func (c *Client) CreateAuthorizationPolicy(request *CreateAuthorizationPolicyRequest) (response *CreateAuthorizationPolicyResponse, err error) {
    return c.CreateAuthorizationPolicyWithContext(context.Background(), request)
}

// CreateAuthorizationPolicy
// 创建MQTT实例的性能测试任务
func (c *Client) CreateAuthorizationPolicyWithContext(ctx context.Context, request *CreateAuthorizationPolicyRequest) (response *CreateAuthorizationPolicyResponse, err error) {
    if request == nil {
        request = NewCreateAuthorizationPolicyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAuthorizationPolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAuthorizationPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewCreateJWKSAuthenticatorRequest() (request *CreateJWKSAuthenticatorRequest) {
    request = &CreateJWKSAuthenticatorRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mqtt", APIVersion, "CreateJWKSAuthenticator")
    
    
    return
}

func NewCreateJWKSAuthenticatorResponse() (response *CreateJWKSAuthenticatorResponse) {
    response = &CreateJWKSAuthenticatorResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateJWKSAuthenticator
// 创建一个jwks的认证
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INSTANCENOTREADY = "FailedOperation.InstanceNotReady"
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
//  UNSUPPORTEDOPERATION_RESOURCEALREADYEXISTS = "UnsupportedOperation.ResourceAlreadyExists"
func (c *Client) CreateJWKSAuthenticator(request *CreateJWKSAuthenticatorRequest) (response *CreateJWKSAuthenticatorResponse, err error) {
    return c.CreateJWKSAuthenticatorWithContext(context.Background(), request)
}

// CreateJWKSAuthenticator
// 创建一个jwks的认证
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INSTANCENOTREADY = "FailedOperation.InstanceNotReady"
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
//  UNSUPPORTEDOPERATION_RESOURCEALREADYEXISTS = "UnsupportedOperation.ResourceAlreadyExists"
func (c *Client) CreateJWKSAuthenticatorWithContext(ctx context.Context, request *CreateJWKSAuthenticatorRequest) (response *CreateJWKSAuthenticatorResponse, err error) {
    if request == nil {
        request = NewCreateJWKSAuthenticatorRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateJWKSAuthenticator require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateJWKSAuthenticatorResponse()
    err = c.Send(request, response)
    return
}

func NewCreateJWTAuthenticatorRequest() (request *CreateJWTAuthenticatorRequest) {
    request = &CreateJWTAuthenticatorRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mqtt", APIVersion, "CreateJWTAuthenticator")
    
    
    return
}

func NewCreateJWTAuthenticatorResponse() (response *CreateJWTAuthenticatorResponse) {
    response = &CreateJWTAuthenticatorResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateJWTAuthenticator
// 创建一个jwks的认证
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INSTANCENOTREADY = "FailedOperation.InstanceNotReady"
//  FAILEDOPERATION_PUBLICKEYVERIFYFAILED = "FailedOperation.PublicKeyVerifyFailed"
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
//  UNSUPPORTEDOPERATION_RESOURCEALREADYEXISTS = "UnsupportedOperation.ResourceAlreadyExists"
func (c *Client) CreateJWTAuthenticator(request *CreateJWTAuthenticatorRequest) (response *CreateJWTAuthenticatorResponse, err error) {
    return c.CreateJWTAuthenticatorWithContext(context.Background(), request)
}

// CreateJWTAuthenticator
// 创建一个jwks的认证
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INSTANCENOTREADY = "FailedOperation.InstanceNotReady"
//  FAILEDOPERATION_PUBLICKEYVERIFYFAILED = "FailedOperation.PublicKeyVerifyFailed"
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
//  UNSUPPORTEDOPERATION_RESOURCEALREADYEXISTS = "UnsupportedOperation.ResourceAlreadyExists"
func (c *Client) CreateJWTAuthenticatorWithContext(ctx context.Context, request *CreateJWTAuthenticatorRequest) (response *CreateJWTAuthenticatorResponse, err error) {
    if request == nil {
        request = NewCreateJWTAuthenticatorRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateJWTAuthenticator require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateJWTAuthenticatorResponse()
    err = c.Send(request, response)
    return
}

func NewCreateTopicRequest() (request *CreateTopicRequest) {
    request = &CreateTopicRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mqtt", APIVersion, "CreateTopic")
    
    
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

func NewDeleteAuthenticatorRequest() (request *DeleteAuthenticatorRequest) {
    request = &DeleteAuthenticatorRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mqtt", APIVersion, "DeleteAuthenticator")
    
    
    return
}

func NewDeleteAuthenticatorResponse() (response *DeleteAuthenticatorResponse) {
    response = &DeleteAuthenticatorResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteAuthenticator
// 根据认证器类型删除一个MQTT认证器
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
//  RESOURCENOTFOUND_ROLE = "ResourceNotFound.Role"
func (c *Client) DeleteAuthenticator(request *DeleteAuthenticatorRequest) (response *DeleteAuthenticatorResponse, err error) {
    return c.DeleteAuthenticatorWithContext(context.Background(), request)
}

// DeleteAuthenticator
// 根据认证器类型删除一个MQTT认证器
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
//  RESOURCENOTFOUND_ROLE = "ResourceNotFound.Role"
func (c *Client) DeleteAuthenticatorWithContext(ctx context.Context, request *DeleteAuthenticatorRequest) (response *DeleteAuthenticatorResponse, err error) {
    if request == nil {
        request = NewDeleteAuthenticatorRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteAuthenticator require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteAuthenticatorResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAuthorizationPolicyRequest() (request *DeleteAuthorizationPolicyRequest) {
    request = &DeleteAuthorizationPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mqtt", APIVersion, "DeleteAuthorizationPolicy")
    
    
    return
}

func NewDeleteAuthorizationPolicyResponse() (response *DeleteAuthorizationPolicyResponse) {
    response = &DeleteAuthorizationPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteAuthorizationPolicy
// 删除策略规则
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
//  RESOURCENOTFOUND_ROLE = "ResourceNotFound.Role"
func (c *Client) DeleteAuthorizationPolicy(request *DeleteAuthorizationPolicyRequest) (response *DeleteAuthorizationPolicyResponse, err error) {
    return c.DeleteAuthorizationPolicyWithContext(context.Background(), request)
}

// DeleteAuthorizationPolicy
// 删除策略规则
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
//  RESOURCENOTFOUND_ROLE = "ResourceNotFound.Role"
func (c *Client) DeleteAuthorizationPolicyWithContext(ctx context.Context, request *DeleteAuthorizationPolicyRequest) (response *DeleteAuthorizationPolicyResponse, err error) {
    if request == nil {
        request = NewDeleteAuthorizationPolicyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteAuthorizationPolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteAuthorizationPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteTopicRequest() (request *DeleteTopicRequest) {
    request = &DeleteTopicRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mqtt", APIVersion, "DeleteTopic")
    
    
    return
}

func NewDeleteTopicResponse() (response *DeleteTopicResponse) {
    response = &DeleteTopicResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteTopic
// 删除MQTT主题
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCENOTFOUND_TOPIC = "ResourceNotFound.Topic"
func (c *Client) DeleteTopic(request *DeleteTopicRequest) (response *DeleteTopicResponse, err error) {
    return c.DeleteTopicWithContext(context.Background(), request)
}

// DeleteTopic
// 删除MQTT主题
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

func NewDescribeAuthenticatorRequest() (request *DescribeAuthenticatorRequest) {
    request = &DescribeAuthenticatorRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mqtt", APIVersion, "DescribeAuthenticator")
    
    
    return
}

func NewDescribeAuthenticatorResponse() (response *DescribeAuthenticatorResponse) {
    response = &DescribeAuthenticatorResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAuthenticator
// 查询MQTT认证器
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
func (c *Client) DescribeAuthenticator(request *DescribeAuthenticatorRequest) (response *DescribeAuthenticatorResponse, err error) {
    return c.DescribeAuthenticatorWithContext(context.Background(), request)
}

// DescribeAuthenticator
// 查询MQTT认证器
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
func (c *Client) DescribeAuthenticatorWithContext(ctx context.Context, request *DescribeAuthenticatorRequest) (response *DescribeAuthenticatorResponse, err error) {
    if request == nil {
        request = NewDescribeAuthenticatorRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAuthenticator require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAuthenticatorResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAuthorizationPoliciesRequest() (request *DescribeAuthorizationPoliciesRequest) {
    request = &DescribeAuthorizationPoliciesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mqtt", APIVersion, "DescribeAuthorizationPolicies")
    
    
    return
}

func NewDescribeAuthorizationPoliciesResponse() (response *DescribeAuthorizationPoliciesResponse) {
    response = &DescribeAuthorizationPoliciesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAuthorizationPolicies
// 查询授权规则
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
func (c *Client) DescribeAuthorizationPolicies(request *DescribeAuthorizationPoliciesRequest) (response *DescribeAuthorizationPoliciesResponse, err error) {
    return c.DescribeAuthorizationPoliciesWithContext(context.Background(), request)
}

// DescribeAuthorizationPolicies
// 查询授权规则
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
func (c *Client) DescribeAuthorizationPoliciesWithContext(ctx context.Context, request *DescribeAuthorizationPoliciesRequest) (response *DescribeAuthorizationPoliciesResponse, err error) {
    if request == nil {
        request = NewDescribeAuthorizationPoliciesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAuthorizationPolicies require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAuthorizationPoliciesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceRequest() (request *DescribeInstanceRequest) {
    request = &DescribeInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mqtt", APIVersion, "DescribeInstance")
    
    
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
    
    request.Init().WithApiInfo("mqtt", APIVersion, "DescribeInstanceList")
    
    
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
    
    request.Init().WithApiInfo("mqtt", APIVersion, "DescribeTopic")
    
    
    return
}

func NewDescribeTopicResponse() (response *DescribeTopicResponse) {
    response = &DescribeTopicResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTopic
// 查询mqtt主题详情
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
//  RESOURCENOTFOUND_TOPIC = "ResourceNotFound.Topic"
func (c *Client) DescribeTopic(request *DescribeTopicRequest) (response *DescribeTopicResponse, err error) {
    return c.DescribeTopicWithContext(context.Background(), request)
}

// DescribeTopic
// 查询mqtt主题详情
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
    
    request.Init().WithApiInfo("mqtt", APIVersion, "DescribeTopicList")
    
    
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

func NewModifyAuthorizationPolicyRequest() (request *ModifyAuthorizationPolicyRequest) {
    request = &ModifyAuthorizationPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mqtt", APIVersion, "ModifyAuthorizationPolicy")
    
    
    return
}

func NewModifyAuthorizationPolicyResponse() (response *ModifyAuthorizationPolicyResponse) {
    response = &ModifyAuthorizationPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyAuthorizationPolicy
// 修改策略规则
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
func (c *Client) ModifyAuthorizationPolicy(request *ModifyAuthorizationPolicyRequest) (response *ModifyAuthorizationPolicyResponse, err error) {
    return c.ModifyAuthorizationPolicyWithContext(context.Background(), request)
}

// ModifyAuthorizationPolicy
// 修改策略规则
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
func (c *Client) ModifyAuthorizationPolicyWithContext(ctx context.Context, request *ModifyAuthorizationPolicyRequest) (response *ModifyAuthorizationPolicyResponse, err error) {
    if request == nil {
        request = NewModifyAuthorizationPolicyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAuthorizationPolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAuthorizationPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewModifyJWKSAuthenticatorRequest() (request *ModifyJWKSAuthenticatorRequest) {
    request = &ModifyJWKSAuthenticatorRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mqtt", APIVersion, "ModifyJWKSAuthenticator")
    
    
    return
}

func NewModifyJWKSAuthenticatorResponse() (response *ModifyJWKSAuthenticatorResponse) {
    response = &ModifyJWKSAuthenticatorResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyJWKSAuthenticator
// 修改MQTT JWKS 认证器
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ModifyJWKSAuthenticator(request *ModifyJWKSAuthenticatorRequest) (response *ModifyJWKSAuthenticatorResponse, err error) {
    return c.ModifyJWKSAuthenticatorWithContext(context.Background(), request)
}

// ModifyJWKSAuthenticator
// 修改MQTT JWKS 认证器
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ModifyJWKSAuthenticatorWithContext(ctx context.Context, request *ModifyJWKSAuthenticatorRequest) (response *ModifyJWKSAuthenticatorResponse, err error) {
    if request == nil {
        request = NewModifyJWKSAuthenticatorRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyJWKSAuthenticator require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyJWKSAuthenticatorResponse()
    err = c.Send(request, response)
    return
}

func NewModifyJWTAuthenticatorRequest() (request *ModifyJWTAuthenticatorRequest) {
    request = &ModifyJWTAuthenticatorRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mqtt", APIVersion, "ModifyJWTAuthenticator")
    
    
    return
}

func NewModifyJWTAuthenticatorResponse() (response *ModifyJWTAuthenticatorResponse) {
    response = &ModifyJWTAuthenticatorResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyJWTAuthenticator
// 修改MQTT JWKS 认证器
//
// 可能返回的错误码:
//  FAILEDOPERATION_PUBLICKEYVERIFYFAILED = "FailedOperation.PublicKeyVerifyFailed"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ModifyJWTAuthenticator(request *ModifyJWTAuthenticatorRequest) (response *ModifyJWTAuthenticatorResponse, err error) {
    return c.ModifyJWTAuthenticatorWithContext(context.Background(), request)
}

// ModifyJWTAuthenticator
// 修改MQTT JWKS 认证器
//
// 可能返回的错误码:
//  FAILEDOPERATION_PUBLICKEYVERIFYFAILED = "FailedOperation.PublicKeyVerifyFailed"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ModifyJWTAuthenticatorWithContext(ctx context.Context, request *ModifyJWTAuthenticatorRequest) (response *ModifyJWTAuthenticatorResponse, err error) {
    if request == nil {
        request = NewModifyJWTAuthenticatorRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyJWTAuthenticator require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyJWTAuthenticatorResponse()
    err = c.Send(request, response)
    return
}

func NewModifyTopicRequest() (request *ModifyTopicRequest) {
    request = &ModifyTopicRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mqtt", APIVersion, "ModifyTopic")
    
    
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

func NewRegisterDeviceCertificateRequest() (request *RegisterDeviceCertificateRequest) {
    request = &RegisterDeviceCertificateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mqtt", APIVersion, "RegisterDeviceCertificate")
    
    
    return
}

func NewRegisterDeviceCertificateResponse() (response *RegisterDeviceCertificateResponse) {
    response = &RegisterDeviceCertificateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RegisterDeviceCertificate
// 注册设备证书
//
// 可能返回的错误码:
//  FAILEDOPERATION_INSTANCENOTREADY = "FailedOperation.InstanceNotReady"
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
func (c *Client) RegisterDeviceCertificate(request *RegisterDeviceCertificateRequest) (response *RegisterDeviceCertificateResponse, err error) {
    return c.RegisterDeviceCertificateWithContext(context.Background(), request)
}

// RegisterDeviceCertificate
// 注册设备证书
//
// 可能返回的错误码:
//  FAILEDOPERATION_INSTANCENOTREADY = "FailedOperation.InstanceNotReady"
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
func (c *Client) RegisterDeviceCertificateWithContext(ctx context.Context, request *RegisterDeviceCertificateRequest) (response *RegisterDeviceCertificateResponse, err error) {
    if request == nil {
        request = NewRegisterDeviceCertificateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RegisterDeviceCertificate require credential")
    }

    request.SetContext(ctx)
    
    response = NewRegisterDeviceCertificateResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateAuthorizationPolicyPriorityRequest() (request *UpdateAuthorizationPolicyPriorityRequest) {
    request = &UpdateAuthorizationPolicyPriorityRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mqtt", APIVersion, "UpdateAuthorizationPolicyPriority")
    
    
    return
}

func NewUpdateAuthorizationPolicyPriorityResponse() (response *UpdateAuthorizationPolicyPriorityResponse) {
    response = &UpdateAuthorizationPolicyPriorityResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateAuthorizationPolicyPriority
// 修改策略规则优先级
//
// 可能返回的错误码:
//  FAILEDOPERATION_INSTANCENOTREADY = "FailedOperation.InstanceNotReady"
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
func (c *Client) UpdateAuthorizationPolicyPriority(request *UpdateAuthorizationPolicyPriorityRequest) (response *UpdateAuthorizationPolicyPriorityResponse, err error) {
    return c.UpdateAuthorizationPolicyPriorityWithContext(context.Background(), request)
}

// UpdateAuthorizationPolicyPriority
// 修改策略规则优先级
//
// 可能返回的错误码:
//  FAILEDOPERATION_INSTANCENOTREADY = "FailedOperation.InstanceNotReady"
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
func (c *Client) UpdateAuthorizationPolicyPriorityWithContext(ctx context.Context, request *UpdateAuthorizationPolicyPriorityRequest) (response *UpdateAuthorizationPolicyPriorityResponse, err error) {
    if request == nil {
        request = NewUpdateAuthorizationPolicyPriorityRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateAuthorizationPolicyPriority require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateAuthorizationPolicyPriorityResponse()
    err = c.Send(request, response)
    return
}
