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


func NewActivateDeviceCertificateRequest() (request *ActivateDeviceCertificateRequest) {
    request = &ActivateDeviceCertificateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mqtt", APIVersion, "ActivateDeviceCertificate")
    
    
    return
}

func NewActivateDeviceCertificateResponse() (response *ActivateDeviceCertificateResponse) {
    response = &ActivateDeviceCertificateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ActivateDeviceCertificate
// 生效设备证书
//
// 可能返回的错误码:
//  FAILEDOPERATION_INSTANCENOTREADY = "FailedOperation.InstanceNotReady"
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
func (c *Client) ActivateDeviceCertificate(request *ActivateDeviceCertificateRequest) (response *ActivateDeviceCertificateResponse, err error) {
    return c.ActivateDeviceCertificateWithContext(context.Background(), request)
}

// ActivateDeviceCertificate
// 生效设备证书
//
// 可能返回的错误码:
//  FAILEDOPERATION_INSTANCENOTREADY = "FailedOperation.InstanceNotReady"
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
func (c *Client) ActivateDeviceCertificateWithContext(ctx context.Context, request *ActivateDeviceCertificateRequest) (response *ActivateDeviceCertificateResponse, err error) {
    if request == nil {
        request = NewActivateDeviceCertificateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ActivateDeviceCertificate require credential")
    }

    request.SetContext(ctx)
    
    response = NewActivateDeviceCertificateResponse()
    err = c.Send(request, response)
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
//
// 可能返回的错误码:
//  FAILEDOPERATION_INSTANCENOTREADY = "FailedOperation.InstanceNotReady"
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
func (c *Client) CreateAuthorizationPolicy(request *CreateAuthorizationPolicyRequest) (response *CreateAuthorizationPolicyResponse, err error) {
    return c.CreateAuthorizationPolicyWithContext(context.Background(), request)
}

// CreateAuthorizationPolicy
// 创建MQTT实例的性能测试任务
//
// 可能返回的错误码:
//  FAILEDOPERATION_INSTANCENOTREADY = "FailedOperation.InstanceNotReady"
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
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

func NewCreateInsPublicEndpointRequest() (request *CreateInsPublicEndpointRequest) {
    request = &CreateInsPublicEndpointRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mqtt", APIVersion, "CreateInsPublicEndpoint")
    
    
    return
}

func NewCreateInsPublicEndpointResponse() (response *CreateInsPublicEndpointResponse) {
    response = &CreateInsPublicEndpointResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateInsPublicEndpoint
// 为MQTT实例创建公网接入点
//
// 可能返回的错误码:
//  FAILEDOPERATION_INSTANCENOTREADY = "FailedOperation.InstanceNotReady"
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
func (c *Client) CreateInsPublicEndpoint(request *CreateInsPublicEndpointRequest) (response *CreateInsPublicEndpointResponse, err error) {
    return c.CreateInsPublicEndpointWithContext(context.Background(), request)
}

// CreateInsPublicEndpoint
// 为MQTT实例创建公网接入点
//
// 可能返回的错误码:
//  FAILEDOPERATION_INSTANCENOTREADY = "FailedOperation.InstanceNotReady"
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
func (c *Client) CreateInsPublicEndpointWithContext(ctx context.Context, request *CreateInsPublicEndpointRequest) (response *CreateInsPublicEndpointResponse, err error) {
    if request == nil {
        request = NewCreateInsPublicEndpointRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateInsPublicEndpoint require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateInsPublicEndpointResponse()
    err = c.Send(request, response)
    return
}

func NewCreateInstanceRequest() (request *CreateInstanceRequest) {
    request = &CreateInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mqtt", APIVersion, "CreateInstance")
    
    
    return
}

func NewCreateInstanceResponse() (response *CreateInstanceResponse) {
    response = &CreateInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateInstance
// 购买新的MQTT实例
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CreateInstance(request *CreateInstanceRequest) (response *CreateInstanceResponse, err error) {
    return c.CreateInstanceWithContext(context.Background(), request)
}

// CreateInstance
// 购买新的MQTT实例
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

func NewCreateUserRequest() (request *CreateUserRequest) {
    request = &CreateUserRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mqtt", APIVersion, "CreateUser")
    
    
    return
}

func NewCreateUserResponse() (response *CreateUserResponse) {
    response = &CreateUserResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateUser
// 添加mqtt角色
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INSTANCENOTREADY = "FailedOperation.InstanceNotReady"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
//  UNSUPPORTEDOPERATION_RESOURCEALREADYEXISTS = "UnsupportedOperation.ResourceAlreadyExists"
func (c *Client) CreateUser(request *CreateUserRequest) (response *CreateUserResponse, err error) {
    return c.CreateUserWithContext(context.Background(), request)
}

// CreateUser
// 添加mqtt角色
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INSTANCENOTREADY = "FailedOperation.InstanceNotReady"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
//  UNSUPPORTEDOPERATION_RESOURCEALREADYEXISTS = "UnsupportedOperation.ResourceAlreadyExists"
func (c *Client) CreateUserWithContext(ctx context.Context, request *CreateUserRequest) (response *CreateUserResponse, err error) {
    if request == nil {
        request = NewCreateUserRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateUser require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateUserResponse()
    err = c.Send(request, response)
    return
}

func NewDeactivateDeviceCertificateRequest() (request *DeactivateDeviceCertificateRequest) {
    request = &DeactivateDeviceCertificateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mqtt", APIVersion, "DeactivateDeviceCertificate")
    
    
    return
}

func NewDeactivateDeviceCertificateResponse() (response *DeactivateDeviceCertificateResponse) {
    response = &DeactivateDeviceCertificateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeactivateDeviceCertificate
// 失效Ca证书
//
// 可能返回的错误码:
//  FAILEDOPERATION_INSTANCENOTREADY = "FailedOperation.InstanceNotReady"
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
func (c *Client) DeactivateDeviceCertificate(request *DeactivateDeviceCertificateRequest) (response *DeactivateDeviceCertificateResponse, err error) {
    return c.DeactivateDeviceCertificateWithContext(context.Background(), request)
}

// DeactivateDeviceCertificate
// 失效Ca证书
//
// 可能返回的错误码:
//  FAILEDOPERATION_INSTANCENOTREADY = "FailedOperation.InstanceNotReady"
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
func (c *Client) DeactivateDeviceCertificateWithContext(ctx context.Context, request *DeactivateDeviceCertificateRequest) (response *DeactivateDeviceCertificateResponse, err error) {
    if request == nil {
        request = NewDeactivateDeviceCertificateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeactivateDeviceCertificate require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeactivateDeviceCertificateResponse()
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

func NewDeleteDeviceCertificateRequest() (request *DeleteDeviceCertificateRequest) {
    request = &DeleteDeviceCertificateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mqtt", APIVersion, "DeleteDeviceCertificate")
    
    
    return
}

func NewDeleteDeviceCertificateResponse() (response *DeleteDeviceCertificateResponse) {
    response = &DeleteDeviceCertificateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteDeviceCertificate
// 删除设备证书
//
// 可能返回的错误码:
//  FAILEDOPERATION_INSTANCENOTREADY = "FailedOperation.InstanceNotReady"
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
func (c *Client) DeleteDeviceCertificate(request *DeleteDeviceCertificateRequest) (response *DeleteDeviceCertificateResponse, err error) {
    return c.DeleteDeviceCertificateWithContext(context.Background(), request)
}

// DeleteDeviceCertificate
// 删除设备证书
//
// 可能返回的错误码:
//  FAILEDOPERATION_INSTANCENOTREADY = "FailedOperation.InstanceNotReady"
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
func (c *Client) DeleteDeviceCertificateWithContext(ctx context.Context, request *DeleteDeviceCertificateRequest) (response *DeleteDeviceCertificateResponse, err error) {
    if request == nil {
        request = NewDeleteDeviceCertificateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteDeviceCertificate require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteDeviceCertificateResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteInsPublicEndpointRequest() (request *DeleteInsPublicEndpointRequest) {
    request = &DeleteInsPublicEndpointRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mqtt", APIVersion, "DeleteInsPublicEndpoint")
    
    
    return
}

func NewDeleteInsPublicEndpointResponse() (response *DeleteInsPublicEndpointResponse) {
    response = &DeleteInsPublicEndpointResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteInsPublicEndpoint
// 删除MQTT实例的公网接入点
//
// 可能返回的错误码:
//  FAILEDOPERATION_INSTANCENOTREADY = "FailedOperation.InstanceNotReady"
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
func (c *Client) DeleteInsPublicEndpoint(request *DeleteInsPublicEndpointRequest) (response *DeleteInsPublicEndpointResponse, err error) {
    return c.DeleteInsPublicEndpointWithContext(context.Background(), request)
}

// DeleteInsPublicEndpoint
// 删除MQTT实例的公网接入点
//
// 可能返回的错误码:
//  FAILEDOPERATION_INSTANCENOTREADY = "FailedOperation.InstanceNotReady"
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
func (c *Client) DeleteInsPublicEndpointWithContext(ctx context.Context, request *DeleteInsPublicEndpointRequest) (response *DeleteInsPublicEndpointResponse, err error) {
    if request == nil {
        request = NewDeleteInsPublicEndpointRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteInsPublicEndpoint require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteInsPublicEndpointResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteInstanceRequest() (request *DeleteInstanceRequest) {
    request = &DeleteInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mqtt", APIVersion, "DeleteInstance")
    
    
    return
}

func NewDeleteInstanceResponse() (response *DeleteInstanceResponse) {
    response = &DeleteInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteInstance
// 删除MQTT实例
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
func (c *Client) DeleteInstance(request *DeleteInstanceRequest) (response *DeleteInstanceResponse, err error) {
    return c.DeleteInstanceWithContext(context.Background(), request)
}

// DeleteInstance
// 删除MQTT实例
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

func NewDeleteUserRequest() (request *DeleteUserRequest) {
    request = &DeleteUserRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mqtt", APIVersion, "DeleteUser")
    
    
    return
}

func NewDeleteUserResponse() (response *DeleteUserResponse) {
    response = &DeleteUserResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteUser
// 删除MQTT访问用户
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
//  RESOURCENOTFOUND_ROLE = "ResourceNotFound.Role"
func (c *Client) DeleteUser(request *DeleteUserRequest) (response *DeleteUserResponse, err error) {
    return c.DeleteUserWithContext(context.Background(), request)
}

// DeleteUser
// 删除MQTT访问用户
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
//  RESOURCENOTFOUND_ROLE = "ResourceNotFound.Role"
func (c *Client) DeleteUserWithContext(ctx context.Context, request *DeleteUserRequest) (response *DeleteUserResponse, err error) {
    if request == nil {
        request = NewDeleteUserRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteUser require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteUserResponse()
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

func NewDescribeDeviceCertificateRequest() (request *DescribeDeviceCertificateRequest) {
    request = &DescribeDeviceCertificateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mqtt", APIVersion, "DescribeDeviceCertificate")
    
    
    return
}

func NewDescribeDeviceCertificateResponse() (response *DescribeDeviceCertificateResponse) {
    response = &DescribeDeviceCertificateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDeviceCertificate
// 查询设备证书详情接口
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
func (c *Client) DescribeDeviceCertificate(request *DescribeDeviceCertificateRequest) (response *DescribeDeviceCertificateResponse, err error) {
    return c.DescribeDeviceCertificateWithContext(context.Background(), request)
}

// DescribeDeviceCertificate
// 查询设备证书详情接口
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
func (c *Client) DescribeDeviceCertificateWithContext(ctx context.Context, request *DescribeDeviceCertificateRequest) (response *DescribeDeviceCertificateResponse, err error) {
    if request == nil {
        request = NewDescribeDeviceCertificateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDeviceCertificate require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDeviceCertificateResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDeviceCertificatesRequest() (request *DescribeDeviceCertificatesRequest) {
    request = &DescribeDeviceCertificatesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mqtt", APIVersion, "DescribeDeviceCertificates")
    
    
    return
}

func NewDescribeDeviceCertificatesResponse() (response *DescribeDeviceCertificatesResponse) {
    response = &DescribeDeviceCertificatesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDeviceCertificates
// 分页查询设备证书
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
func (c *Client) DescribeDeviceCertificates(request *DescribeDeviceCertificatesRequest) (response *DescribeDeviceCertificatesResponse, err error) {
    return c.DescribeDeviceCertificatesWithContext(context.Background(), request)
}

// DescribeDeviceCertificates
// 分页查询设备证书
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
func (c *Client) DescribeDeviceCertificatesWithContext(ctx context.Context, request *DescribeDeviceCertificatesRequest) (response *DescribeDeviceCertificatesResponse, err error) {
    if request == nil {
        request = NewDescribeDeviceCertificatesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDeviceCertificates require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDeviceCertificatesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInsPublicEndpointsRequest() (request *DescribeInsPublicEndpointsRequest) {
    request = &DescribeInsPublicEndpointsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mqtt", APIVersion, "DescribeInsPublicEndpoints")
    
    
    return
}

func NewDescribeInsPublicEndpointsResponse() (response *DescribeInsPublicEndpointsResponse) {
    response = &DescribeInsPublicEndpointsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeInsPublicEndpoints
// 查询MQTT实例公网接入点
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
func (c *Client) DescribeInsPublicEndpoints(request *DescribeInsPublicEndpointsRequest) (response *DescribeInsPublicEndpointsResponse, err error) {
    return c.DescribeInsPublicEndpointsWithContext(context.Background(), request)
}

// DescribeInsPublicEndpoints
// 查询MQTT实例公网接入点
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
func (c *Client) DescribeInsPublicEndpointsWithContext(ctx context.Context, request *DescribeInsPublicEndpointsRequest) (response *DescribeInsPublicEndpointsResponse, err error) {
    if request == nil {
        request = NewDescribeInsPublicEndpointsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInsPublicEndpoints require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInsPublicEndpointsResponse()
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

func NewDescribeUserListRequest() (request *DescribeUserListRequest) {
    request = &DescribeUserListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mqtt", APIVersion, "DescribeUserList")
    
    
    return
}

func NewDescribeUserListResponse() (response *DescribeUserListResponse) {
    response = &DescribeUserListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeUserList
// 查询用户列表，Filter参数使用说明如下：
//
// 
//
// 1. Username，用户名称模糊搜索
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
func (c *Client) DescribeUserList(request *DescribeUserListRequest) (response *DescribeUserListResponse, err error) {
    return c.DescribeUserListWithContext(context.Background(), request)
}

// DescribeUserList
// 查询用户列表，Filter参数使用说明如下：
//
// 
//
// 1. Username，用户名称模糊搜索
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
func (c *Client) DescribeUserListWithContext(ctx context.Context, request *DescribeUserListRequest) (response *DescribeUserListResponse, err error) {
    if request == nil {
        request = NewDescribeUserListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUserList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeUserListResponse()
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

func NewModifyInsPublicEndpointRequest() (request *ModifyInsPublicEndpointRequest) {
    request = &ModifyInsPublicEndpointRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mqtt", APIVersion, "ModifyInsPublicEndpoint")
    
    
    return
}

func NewModifyInsPublicEndpointResponse() (response *ModifyInsPublicEndpointResponse) {
    response = &ModifyInsPublicEndpointResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyInsPublicEndpoint
// 更新MQTT实例公网接入点
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
func (c *Client) ModifyInsPublicEndpoint(request *ModifyInsPublicEndpointRequest) (response *ModifyInsPublicEndpointResponse, err error) {
    return c.ModifyInsPublicEndpointWithContext(context.Background(), request)
}

// ModifyInsPublicEndpoint
// 更新MQTT实例公网接入点
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
func (c *Client) ModifyInsPublicEndpointWithContext(ctx context.Context, request *ModifyInsPublicEndpointRequest) (response *ModifyInsPublicEndpointResponse, err error) {
    if request == nil {
        request = NewModifyInsPublicEndpointRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyInsPublicEndpoint require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyInsPublicEndpointResponse()
    err = c.Send(request, response)
    return
}

func NewModifyInstanceRequest() (request *ModifyInstanceRequest) {
    request = &ModifyInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mqtt", APIVersion, "ModifyInstance")
    
    
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

func NewModifyUserRequest() (request *ModifyUserRequest) {
    request = &ModifyUserRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mqtt", APIVersion, "ModifyUser")
    
    
    return
}

func NewModifyUserResponse() (response *ModifyUserResponse) {
    response = &ModifyUserResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyUser
// 修改MQTT角色
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ModifyUser(request *ModifyUserRequest) (response *ModifyUserResponse, err error) {
    return c.ModifyUserWithContext(context.Background(), request)
}

// ModifyUser
// 修改MQTT角色
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ModifyUserWithContext(ctx context.Context, request *ModifyUserRequest) (response *ModifyUserResponse, err error) {
    if request == nil {
        request = NewModifyUserRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyUser require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyUserResponse()
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

func NewRevokedDeviceCertificateRequest() (request *RevokedDeviceCertificateRequest) {
    request = &RevokedDeviceCertificateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mqtt", APIVersion, "RevokedDeviceCertificate")
    
    
    return
}

func NewRevokedDeviceCertificateResponse() (response *RevokedDeviceCertificateResponse) {
    response = &RevokedDeviceCertificateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RevokedDeviceCertificate
// 吊销设备证书
//
// 可能返回的错误码:
//  FAILEDOPERATION_INSTANCENOTREADY = "FailedOperation.InstanceNotReady"
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
func (c *Client) RevokedDeviceCertificate(request *RevokedDeviceCertificateRequest) (response *RevokedDeviceCertificateResponse, err error) {
    return c.RevokedDeviceCertificateWithContext(context.Background(), request)
}

// RevokedDeviceCertificate
// 吊销设备证书
//
// 可能返回的错误码:
//  FAILEDOPERATION_INSTANCENOTREADY = "FailedOperation.InstanceNotReady"
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
func (c *Client) RevokedDeviceCertificateWithContext(ctx context.Context, request *RevokedDeviceCertificateRequest) (response *RevokedDeviceCertificateResponse, err error) {
    if request == nil {
        request = NewRevokedDeviceCertificateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RevokedDeviceCertificate require credential")
    }

    request.SetContext(ctx)
    
    response = NewRevokedDeviceCertificateResponse()
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
