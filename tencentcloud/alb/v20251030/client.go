// Copyright (c) 2017-2025 Tencent. All Rights Reserved.
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

package v20251030

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2025-10-30"

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


func NewAddTargetsToTargetGroupRequest() (request *AddTargetsToTargetGroupRequest) {
    request = &AddTargetsToTargetGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("alb", APIVersion, "AddTargetsToTargetGroup")
    
    
    return
}

func NewAddTargetsToTargetGroupResponse() (response *AddTargetsToTargetGroupResponse) {
    response = &AddTargetsToTargetGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AddTargetsToTargetGroup
// 向目标组内添加后端服务
func (c *Client) AddTargetsToTargetGroup(request *AddTargetsToTargetGroupRequest) (response *AddTargetsToTargetGroupResponse, err error) {
    return c.AddTargetsToTargetGroupWithContext(context.Background(), request)
}

// AddTargetsToTargetGroup
// 向目标组内添加后端服务
func (c *Client) AddTargetsToTargetGroupWithContext(ctx context.Context, request *AddTargetsToTargetGroupRequest) (response *AddTargetsToTargetGroupResponse, err error) {
    if request == nil {
        request = NewAddTargetsToTargetGroupRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "alb", APIVersion, "AddTargetsToTargetGroup")
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddTargetsToTargetGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddTargetsToTargetGroupResponse()
    err = c.Send(request, response)
    return
}

func NewAssociateBandwidthPackageWithLoadBalancerRequest() (request *AssociateBandwidthPackageWithLoadBalancerRequest) {
    request = &AssociateBandwidthPackageWithLoadBalancerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("alb", APIVersion, "AssociateBandwidthPackageWithLoadBalancer")
    
    
    return
}

func NewAssociateBandwidthPackageWithLoadBalancerResponse() (response *AssociateBandwidthPackageWithLoadBalancerResponse) {
    response = &AssociateBandwidthPackageWithLoadBalancerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AssociateBandwidthPackageWithLoadBalancer
// 将共享带宽包绑定到应用型负载均衡实例。
func (c *Client) AssociateBandwidthPackageWithLoadBalancer(request *AssociateBandwidthPackageWithLoadBalancerRequest) (response *AssociateBandwidthPackageWithLoadBalancerResponse, err error) {
    return c.AssociateBandwidthPackageWithLoadBalancerWithContext(context.Background(), request)
}

// AssociateBandwidthPackageWithLoadBalancer
// 将共享带宽包绑定到应用型负载均衡实例。
func (c *Client) AssociateBandwidthPackageWithLoadBalancerWithContext(ctx context.Context, request *AssociateBandwidthPackageWithLoadBalancerRequest) (response *AssociateBandwidthPackageWithLoadBalancerResponse, err error) {
    if request == nil {
        request = NewAssociateBandwidthPackageWithLoadBalancerRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "alb", APIVersion, "AssociateBandwidthPackageWithLoadBalancer")
    
    if c.GetCredential() == nil {
        return nil, errors.New("AssociateBandwidthPackageWithLoadBalancer require credential")
    }

    request.SetContext(ctx)
    
    response = NewAssociateBandwidthPackageWithLoadBalancerResponse()
    err = c.Send(request, response)
    return
}

func NewAssociateListenerAdditionalCertificatesRequest() (request *AssociateListenerAdditionalCertificatesRequest) {
    request = &AssociateListenerAdditionalCertificatesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("alb", APIVersion, "AssociateListenerAdditionalCertificates")
    
    
    return
}

func NewAssociateListenerAdditionalCertificatesResponse() (response *AssociateListenerAdditionalCertificatesResponse) {
    response = &AssociateListenerAdditionalCertificatesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AssociateListenerAdditionalCertificates
// AssociateListenerAdditionalCertificates属于异步接口，即系统返回一个请求 ID，但该扩展证书尚未添加成功，系统后台的添加任务仍在进行。您可以调用DescribeListenerCertificates接口查询扩展证书的添加状态：
//
// 当HTTPS和QUIC监听器处于Associating状态时，表示扩展证书正在添加中。
//
// 当HTTPS和QUIC监听器处于Associated状态时，表示扩展证书添加成功。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_CERTIFICATE = "ResourceNotFound.Certificate"
//  RESOURCENOTFOUND_LISTENER = "ResourceNotFound.Listener"
//  RESOURCENOTFOUND_LOADBALANCER = "ResourceNotFound.LoadBalancer"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_CERTIFICATE = "ResourceUnavailable.Certificate"
//  RESOURCEUNAVAILABLE_LISTENER = "ResourceUnavailable.Listener"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_CERTIFICATEALREADYBOUND = "UnsupportedOperation.CertificateAlreadyBound"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDPROTOCOL = "UnsupportedOperation.UnsupportedProtocol"
func (c *Client) AssociateListenerAdditionalCertificates(request *AssociateListenerAdditionalCertificatesRequest) (response *AssociateListenerAdditionalCertificatesResponse, err error) {
    return c.AssociateListenerAdditionalCertificatesWithContext(context.Background(), request)
}

// AssociateListenerAdditionalCertificates
// AssociateListenerAdditionalCertificates属于异步接口，即系统返回一个请求 ID，但该扩展证书尚未添加成功，系统后台的添加任务仍在进行。您可以调用DescribeListenerCertificates接口查询扩展证书的添加状态：
//
// 当HTTPS和QUIC监听器处于Associating状态时，表示扩展证书正在添加中。
//
// 当HTTPS和QUIC监听器处于Associated状态时，表示扩展证书添加成功。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_CERTIFICATE = "ResourceNotFound.Certificate"
//  RESOURCENOTFOUND_LISTENER = "ResourceNotFound.Listener"
//  RESOURCENOTFOUND_LOADBALANCER = "ResourceNotFound.LoadBalancer"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_CERTIFICATE = "ResourceUnavailable.Certificate"
//  RESOURCEUNAVAILABLE_LISTENER = "ResourceUnavailable.Listener"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_CERTIFICATEALREADYBOUND = "UnsupportedOperation.CertificateAlreadyBound"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDPROTOCOL = "UnsupportedOperation.UnsupportedProtocol"
func (c *Client) AssociateListenerAdditionalCertificatesWithContext(ctx context.Context, request *AssociateListenerAdditionalCertificatesRequest) (response *AssociateListenerAdditionalCertificatesResponse, err error) {
    if request == nil {
        request = NewAssociateListenerAdditionalCertificatesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "alb", APIVersion, "AssociateListenerAdditionalCertificates")
    
    if c.GetCredential() == nil {
        return nil, errors.New("AssociateListenerAdditionalCertificates require credential")
    }

    request.SetContext(ctx)
    
    response = NewAssociateListenerAdditionalCertificatesResponse()
    err = c.Send(request, response)
    return
}

func NewCreateHealthCheckTemplateRequest() (request *CreateHealthCheckTemplateRequest) {
    request = &CreateHealthCheckTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("alb", APIVersion, "CreateHealthCheckTemplate")
    
    
    return
}

func NewCreateHealthCheckTemplateResponse() (response *CreateHealthCheckTemplateResponse) {
    response = &CreateHealthCheckTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateHealthCheckTemplate
// 创建健康检查模板
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_CERTIFICATE = "ResourceNotFound.Certificate"
//  RESOURCENOTFOUND_LISTENER = "ResourceNotFound.Listener"
//  RESOURCENOTFOUND_LOADBALANCER = "ResourceNotFound.LoadBalancer"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_CERTIFICATE = "ResourceUnavailable.Certificate"
//  RESOURCEUNAVAILABLE_LISTENER = "ResourceUnavailable.Listener"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_CERTIFICATEALREADYBOUND = "UnsupportedOperation.CertificateAlreadyBound"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDPROTOCOL = "UnsupportedOperation.UnsupportedProtocol"
func (c *Client) CreateHealthCheckTemplate(request *CreateHealthCheckTemplateRequest) (response *CreateHealthCheckTemplateResponse, err error) {
    return c.CreateHealthCheckTemplateWithContext(context.Background(), request)
}

// CreateHealthCheckTemplate
// 创建健康检查模板
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_CERTIFICATE = "ResourceNotFound.Certificate"
//  RESOURCENOTFOUND_LISTENER = "ResourceNotFound.Listener"
//  RESOURCENOTFOUND_LOADBALANCER = "ResourceNotFound.LoadBalancer"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_CERTIFICATE = "ResourceUnavailable.Certificate"
//  RESOURCEUNAVAILABLE_LISTENER = "ResourceUnavailable.Listener"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_CERTIFICATEALREADYBOUND = "UnsupportedOperation.CertificateAlreadyBound"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDPROTOCOL = "UnsupportedOperation.UnsupportedProtocol"
func (c *Client) CreateHealthCheckTemplateWithContext(ctx context.Context, request *CreateHealthCheckTemplateRequest) (response *CreateHealthCheckTemplateResponse, err error) {
    if request == nil {
        request = NewCreateHealthCheckTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "alb", APIVersion, "CreateHealthCheckTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateHealthCheckTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateHealthCheckTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewCreateListenerRequest() (request *CreateListenerRequest) {
    request = &CreateListenerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("alb", APIVersion, "CreateListener")
    
    
    return
}

func NewCreateListenerResponse() (response *CreateListenerResponse) {
    response = &CreateListenerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateListener
// 创建监听器
//
// 可能返回的错误码:
//  UNSUPPORTEDOPERATION_TARGETGROUPPROTOCOLMISMATCH = "UnsupportedOperation.TargetGroupProtocolMismatch"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDPROTOCOL = "UnsupportedOperation.UnsupportedProtocol"
func (c *Client) CreateListener(request *CreateListenerRequest) (response *CreateListenerResponse, err error) {
    return c.CreateListenerWithContext(context.Background(), request)
}

// CreateListener
// 创建监听器
//
// 可能返回的错误码:
//  UNSUPPORTEDOPERATION_TARGETGROUPPROTOCOLMISMATCH = "UnsupportedOperation.TargetGroupProtocolMismatch"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDPROTOCOL = "UnsupportedOperation.UnsupportedProtocol"
func (c *Client) CreateListenerWithContext(ctx context.Context, request *CreateListenerRequest) (response *CreateListenerResponse, err error) {
    if request == nil {
        request = NewCreateListenerRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "alb", APIVersion, "CreateListener")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateListener require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateListenerResponse()
    err = c.Send(request, response)
    return
}

func NewCreateLoadBalancerRequest() (request *CreateLoadBalancerRequest) {
    request = &CreateLoadBalancerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("alb", APIVersion, "CreateLoadBalancer")
    
    
    return
}

func NewCreateLoadBalancerResponse() (response *CreateLoadBalancerResponse) {
    response = &CreateLoadBalancerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateLoadBalancer
// **CreateLoadBalancer**接口属于异步接口，即系统返回一个实例ID，但该应用型负载均衡实例尚未创建成功，系统后台的创建任务仍在进行。您可以调用[DescribeLoadBalancerDetail](214362)查询应用型负载均衡实例的创建状态：
//
// - 当应用型负载均衡实例处于**Provisioning**状态时，表示应用型负载均衡实例正在创建中。
//
// - 当应用型负载均衡实例处于**Active**状态时，表示应用型负载均衡实例创建成功。
//
// 可能返回的错误码:
//  UNSUPPORTEDOPERATION_TARGETGROUPPROTOCOLMISMATCH = "UnsupportedOperation.TargetGroupProtocolMismatch"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDPROTOCOL = "UnsupportedOperation.UnsupportedProtocol"
func (c *Client) CreateLoadBalancer(request *CreateLoadBalancerRequest) (response *CreateLoadBalancerResponse, err error) {
    return c.CreateLoadBalancerWithContext(context.Background(), request)
}

// CreateLoadBalancer
// **CreateLoadBalancer**接口属于异步接口，即系统返回一个实例ID，但该应用型负载均衡实例尚未创建成功，系统后台的创建任务仍在进行。您可以调用[DescribeLoadBalancerDetail](214362)查询应用型负载均衡实例的创建状态：
//
// - 当应用型负载均衡实例处于**Provisioning**状态时，表示应用型负载均衡实例正在创建中。
//
// - 当应用型负载均衡实例处于**Active**状态时，表示应用型负载均衡实例创建成功。
//
// 可能返回的错误码:
//  UNSUPPORTEDOPERATION_TARGETGROUPPROTOCOLMISMATCH = "UnsupportedOperation.TargetGroupProtocolMismatch"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDPROTOCOL = "UnsupportedOperation.UnsupportedProtocol"
func (c *Client) CreateLoadBalancerWithContext(ctx context.Context, request *CreateLoadBalancerRequest) (response *CreateLoadBalancerResponse, err error) {
    if request == nil {
        request = NewCreateLoadBalancerRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "alb", APIVersion, "CreateLoadBalancer")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateLoadBalancer require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateLoadBalancerResponse()
    err = c.Send(request, response)
    return
}

func NewCreateRulesRequest() (request *CreateRulesRequest) {
    request = &CreateRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("alb", APIVersion, "CreateRules")
    
    
    return
}

func NewCreateRulesResponse() (response *CreateRulesResponse) {
    response = &CreateRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateRules
// CreateRules创建转发规则，本接口为异步接口，返回成功后需以返回的RequestID为入参，调用DescribeAsyncJobs接口查询本次任务是否成功。
//
// 一条规则最多支持10个转发条件（Conditions），5个转发动作（Actions）。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_RESOURCEALREADYEXISTS = "FailedOperation.ResourceAlreadyExists"
//  FAILEDOPERATION_RULEALREADYEXISTS = "FailedOperation.RuleAlreadyExists"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDCLIENTTOKEN = "InvalidParameter.InvalidClientToken"
//  INVALIDPARAMETER_INVALIDFIELDFORMAT = "InvalidParameter.InvalidFieldFormat"
//  INVALIDPARAMETER_INVALIDFIELDTYPE = "InvalidParameter.InvalidFieldType"
//  INVALIDPARAMETER_INVALIDFIELDVALUE = "InvalidParameter.InvalidFieldValue"
//  INVALIDPARAMETER_MISSINGACTION = "InvalidParameter.MissingAction"
//  INVALIDPARAMETER_MISSINGFIELD = "InvalidParameter.MissingField"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_BUSINESSPARAMETERMISMATCH = "InvalidParameterValue.BusinessParameterMismatch"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_QUOTA = "LimitExceeded.Quota"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_LISTENER = "ResourceNotFound.Listener"
//  RESOURCENOTFOUND_LOADBALANCER = "ResourceNotFound.LoadBalancer"
//  RESOURCENOTFOUND_TARGETGROUP = "ResourceNotFound.TargetGroup"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_LISTENER = "ResourceUnavailable.Listener"
//  RESOURCEUNAVAILABLE_LOADBALANCER = "ResourceUnavailable.LoadBalancer"
//  RESOURCEUNAVAILABLE_PREVIOUSTASKNOTCOMPLETED = "ResourceUnavailable.PreviousTaskNotCompleted"
//  RESOURCEUNAVAILABLE_TARGETGROUP = "ResourceUnavailable.TargetGroup"
//  RESOURCEUNAVAILABLE_UPSTREAMSTATUSBLOCKED = "ResourceUnavailable.UpstreamStatusBlocked"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_SERVICEROLEUNAUTHORIZED = "UnauthorizedOperation.ServiceRoleUnauthorized"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_INVALIDSTATETRANSITION = "UnsupportedOperation.InvalidStateTransition"
func (c *Client) CreateRules(request *CreateRulesRequest) (response *CreateRulesResponse, err error) {
    return c.CreateRulesWithContext(context.Background(), request)
}

// CreateRules
// CreateRules创建转发规则，本接口为异步接口，返回成功后需以返回的RequestID为入参，调用DescribeAsyncJobs接口查询本次任务是否成功。
//
// 一条规则最多支持10个转发条件（Conditions），5个转发动作（Actions）。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_RESOURCEALREADYEXISTS = "FailedOperation.ResourceAlreadyExists"
//  FAILEDOPERATION_RULEALREADYEXISTS = "FailedOperation.RuleAlreadyExists"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDCLIENTTOKEN = "InvalidParameter.InvalidClientToken"
//  INVALIDPARAMETER_INVALIDFIELDFORMAT = "InvalidParameter.InvalidFieldFormat"
//  INVALIDPARAMETER_INVALIDFIELDTYPE = "InvalidParameter.InvalidFieldType"
//  INVALIDPARAMETER_INVALIDFIELDVALUE = "InvalidParameter.InvalidFieldValue"
//  INVALIDPARAMETER_MISSINGACTION = "InvalidParameter.MissingAction"
//  INVALIDPARAMETER_MISSINGFIELD = "InvalidParameter.MissingField"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_BUSINESSPARAMETERMISMATCH = "InvalidParameterValue.BusinessParameterMismatch"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_QUOTA = "LimitExceeded.Quota"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_LISTENER = "ResourceNotFound.Listener"
//  RESOURCENOTFOUND_LOADBALANCER = "ResourceNotFound.LoadBalancer"
//  RESOURCENOTFOUND_TARGETGROUP = "ResourceNotFound.TargetGroup"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_LISTENER = "ResourceUnavailable.Listener"
//  RESOURCEUNAVAILABLE_LOADBALANCER = "ResourceUnavailable.LoadBalancer"
//  RESOURCEUNAVAILABLE_PREVIOUSTASKNOTCOMPLETED = "ResourceUnavailable.PreviousTaskNotCompleted"
//  RESOURCEUNAVAILABLE_TARGETGROUP = "ResourceUnavailable.TargetGroup"
//  RESOURCEUNAVAILABLE_UPSTREAMSTATUSBLOCKED = "ResourceUnavailable.UpstreamStatusBlocked"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_SERVICEROLEUNAUTHORIZED = "UnauthorizedOperation.ServiceRoleUnauthorized"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_INVALIDSTATETRANSITION = "UnsupportedOperation.InvalidStateTransition"
func (c *Client) CreateRulesWithContext(ctx context.Context, request *CreateRulesRequest) (response *CreateRulesResponse, err error) {
    if request == nil {
        request = NewCreateRulesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "alb", APIVersion, "CreateRules")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateRulesResponse()
    err = c.Send(request, response)
    return
}

func NewCreateSecurityPolicyRequest() (request *CreateSecurityPolicyRequest) {
    request = &CreateSecurityPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("alb", APIVersion, "CreateSecurityPolicy")
    
    
    return
}

func NewCreateSecurityPolicyResponse() (response *CreateSecurityPolicyResponse) {
    response = &CreateSecurityPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateSecurityPolicy
// 创建自定义安全策略，用于配置 HTTPS 监听器的 TLS 协议版本和加密套件。通过安全策略，您可以灵活控制客户端与负载均衡之间 HTTPS 通信的安全级别。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFIELDVALUE = "InvalidParameter.InvalidFieldValue"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateSecurityPolicy(request *CreateSecurityPolicyRequest) (response *CreateSecurityPolicyResponse, err error) {
    return c.CreateSecurityPolicyWithContext(context.Background(), request)
}

// CreateSecurityPolicy
// 创建自定义安全策略，用于配置 HTTPS 监听器的 TLS 协议版本和加密套件。通过安全策略，您可以灵活控制客户端与负载均衡之间 HTTPS 通信的安全级别。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFIELDVALUE = "InvalidParameter.InvalidFieldValue"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateSecurityPolicyWithContext(ctx context.Context, request *CreateSecurityPolicyRequest) (response *CreateSecurityPolicyResponse, err error) {
    if request == nil {
        request = NewCreateSecurityPolicyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "alb", APIVersion, "CreateSecurityPolicy")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateSecurityPolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateSecurityPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewCreateTargetGroupRequest() (request *CreateTargetGroupRequest) {
    request = &CreateTargetGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("alb", APIVersion, "CreateTargetGroup")
    
    
    return
}

func NewCreateTargetGroupResponse() (response *CreateTargetGroupResponse) {
    response = &CreateTargetGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateTargetGroup
// 目标组相关接口
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFIELDVALUE = "InvalidParameter.InvalidFieldValue"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateTargetGroup(request *CreateTargetGroupRequest) (response *CreateTargetGroupResponse, err error) {
    return c.CreateTargetGroupWithContext(context.Background(), request)
}

// CreateTargetGroup
// 目标组相关接口
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFIELDVALUE = "InvalidParameter.InvalidFieldValue"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateTargetGroupWithContext(ctx context.Context, request *CreateTargetGroupRequest) (response *CreateTargetGroupResponse, err error) {
    if request == nil {
        request = NewCreateTargetGroupRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "alb", APIVersion, "CreateTargetGroup")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateTargetGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateTargetGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteHealthCheckTemplatesRequest() (request *DeleteHealthCheckTemplatesRequest) {
    request = &DeleteHealthCheckTemplatesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("alb", APIVersion, "DeleteHealthCheckTemplates")
    
    
    return
}

func NewDeleteHealthCheckTemplatesResponse() (response *DeleteHealthCheckTemplatesResponse) {
    response = &DeleteHealthCheckTemplatesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteHealthCheckTemplates
// 删除健康检查模板
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFIELDVALUE = "InvalidParameter.InvalidFieldValue"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteHealthCheckTemplates(request *DeleteHealthCheckTemplatesRequest) (response *DeleteHealthCheckTemplatesResponse, err error) {
    return c.DeleteHealthCheckTemplatesWithContext(context.Background(), request)
}

// DeleteHealthCheckTemplates
// 删除健康检查模板
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFIELDVALUE = "InvalidParameter.InvalidFieldValue"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteHealthCheckTemplatesWithContext(ctx context.Context, request *DeleteHealthCheckTemplatesRequest) (response *DeleteHealthCheckTemplatesResponse, err error) {
    if request == nil {
        request = NewDeleteHealthCheckTemplatesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "alb", APIVersion, "DeleteHealthCheckTemplates")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteHealthCheckTemplates require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteHealthCheckTemplatesResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteListenerRequest() (request *DeleteListenerRequest) {
    request = &DeleteListenerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("alb", APIVersion, "DeleteListener")
    
    
    return
}

func NewDeleteListenerResponse() (response *DeleteListenerResponse) {
    response = &DeleteListenerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteListener
// 删除监听器
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFIELDVALUE = "InvalidParameter.InvalidFieldValue"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteListener(request *DeleteListenerRequest) (response *DeleteListenerResponse, err error) {
    return c.DeleteListenerWithContext(context.Background(), request)
}

// DeleteListener
// 删除监听器
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFIELDVALUE = "InvalidParameter.InvalidFieldValue"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteListenerWithContext(ctx context.Context, request *DeleteListenerRequest) (response *DeleteListenerResponse, err error) {
    if request == nil {
        request = NewDeleteListenerRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "alb", APIVersion, "DeleteListener")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteListener require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteListenerResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteLoadBalancersRequest() (request *DeleteLoadBalancersRequest) {
    request = &DeleteLoadBalancersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("alb", APIVersion, "DeleteLoadBalancers")
    
    
    return
}

func NewDeleteLoadBalancersResponse() (response *DeleteLoadBalancersResponse) {
    response = &DeleteLoadBalancersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteLoadBalancers
// **DeleteLoadBalancers**接口属于异步接口，即系统返回一个请求ID，但该应用型负载均衡实例尚未删除成功，系统后台的删除任务仍在进行。您可以调用[DescribeLoadBalancerDetails](214362)查询应用型负载均衡实例的删除状态：
//
// - 当应用型负载均衡实例处于**Deleting**状态时，表示应用型负载均衡实例正在删除中。
//
// - 当查询不到指定的应用型负载均衡实例时，表示应用型负载均衡实例删除成功。
//
// 可能返回的错误码:
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE_LOADBALANCERHASLISTENERS = "ResourceInUse.LoadBalancerHasListeners"
func (c *Client) DeleteLoadBalancers(request *DeleteLoadBalancersRequest) (response *DeleteLoadBalancersResponse, err error) {
    return c.DeleteLoadBalancersWithContext(context.Background(), request)
}

// DeleteLoadBalancers
// **DeleteLoadBalancers**接口属于异步接口，即系统返回一个请求ID，但该应用型负载均衡实例尚未删除成功，系统后台的删除任务仍在进行。您可以调用[DescribeLoadBalancerDetails](214362)查询应用型负载均衡实例的删除状态：
//
// - 当应用型负载均衡实例处于**Deleting**状态时，表示应用型负载均衡实例正在删除中。
//
// - 当查询不到指定的应用型负载均衡实例时，表示应用型负载均衡实例删除成功。
//
// 可能返回的错误码:
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE_LOADBALANCERHASLISTENERS = "ResourceInUse.LoadBalancerHasListeners"
func (c *Client) DeleteLoadBalancersWithContext(ctx context.Context, request *DeleteLoadBalancersRequest) (response *DeleteLoadBalancersResponse, err error) {
    if request == nil {
        request = NewDeleteLoadBalancersRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "alb", APIVersion, "DeleteLoadBalancers")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteLoadBalancers require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteLoadBalancersResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteRulesRequest() (request *DeleteRulesRequest) {
    request = &DeleteRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("alb", APIVersion, "DeleteRules")
    
    
    return
}

func NewDeleteRulesResponse() (response *DeleteRulesResponse) {
    response = &DeleteRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteRules
// DeleteRules删除转发规则，本接口为异步接口，返回成功后需以返回的RequestID为入参，调用DescribeAsyncJobs接口查询本次任务是否成功。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDCLIENTTOKEN = "InvalidParameter.InvalidClientToken"
//  INVALIDPARAMETER_INVALIDFIELDFORMAT = "InvalidParameter.InvalidFieldFormat"
//  INVALIDPARAMETER_INVALIDFIELDTYPE = "InvalidParameter.InvalidFieldType"
//  INVALIDPARAMETER_INVALIDFIELDVALUE = "InvalidParameter.InvalidFieldValue"
//  INVALIDPARAMETER_MISSINGACTION = "InvalidParameter.MissingAction"
//  INVALIDPARAMETER_MISSINGFIELD = "InvalidParameter.MissingField"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_LISTENER = "ResourceNotFound.Listener"
//  RESOURCENOTFOUND_LOADBALANCER = "ResourceNotFound.LoadBalancer"
//  RESOURCENOTFOUND_RULE = "ResourceNotFound.Rule"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_LISTENER = "ResourceUnavailable.Listener"
//  RESOURCEUNAVAILABLE_LOADBALANCER = "ResourceUnavailable.LoadBalancer"
//  RESOURCEUNAVAILABLE_PREVIOUSTASKNOTCOMPLETED = "ResourceUnavailable.PreviousTaskNotCompleted"
//  RESOURCEUNAVAILABLE_RULE = "ResourceUnavailable.Rule"
//  RESOURCEUNAVAILABLE_UPSTREAMSTATUSBLOCKED = "ResourceUnavailable.UpstreamStatusBlocked"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_SERVICEROLEUNAUTHORIZED = "UnauthorizedOperation.ServiceRoleUnauthorized"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteRules(request *DeleteRulesRequest) (response *DeleteRulesResponse, err error) {
    return c.DeleteRulesWithContext(context.Background(), request)
}

// DeleteRules
// DeleteRules删除转发规则，本接口为异步接口，返回成功后需以返回的RequestID为入参，调用DescribeAsyncJobs接口查询本次任务是否成功。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDCLIENTTOKEN = "InvalidParameter.InvalidClientToken"
//  INVALIDPARAMETER_INVALIDFIELDFORMAT = "InvalidParameter.InvalidFieldFormat"
//  INVALIDPARAMETER_INVALIDFIELDTYPE = "InvalidParameter.InvalidFieldType"
//  INVALIDPARAMETER_INVALIDFIELDVALUE = "InvalidParameter.InvalidFieldValue"
//  INVALIDPARAMETER_MISSINGACTION = "InvalidParameter.MissingAction"
//  INVALIDPARAMETER_MISSINGFIELD = "InvalidParameter.MissingField"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_LISTENER = "ResourceNotFound.Listener"
//  RESOURCENOTFOUND_LOADBALANCER = "ResourceNotFound.LoadBalancer"
//  RESOURCENOTFOUND_RULE = "ResourceNotFound.Rule"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_LISTENER = "ResourceUnavailable.Listener"
//  RESOURCEUNAVAILABLE_LOADBALANCER = "ResourceUnavailable.LoadBalancer"
//  RESOURCEUNAVAILABLE_PREVIOUSTASKNOTCOMPLETED = "ResourceUnavailable.PreviousTaskNotCompleted"
//  RESOURCEUNAVAILABLE_RULE = "ResourceUnavailable.Rule"
//  RESOURCEUNAVAILABLE_UPSTREAMSTATUSBLOCKED = "ResourceUnavailable.UpstreamStatusBlocked"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_SERVICEROLEUNAUTHORIZED = "UnauthorizedOperation.ServiceRoleUnauthorized"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteRulesWithContext(ctx context.Context, request *DeleteRulesRequest) (response *DeleteRulesResponse, err error) {
    if request == nil {
        request = NewDeleteRulesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "alb", APIVersion, "DeleteRules")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteRulesResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteSecurityPolicyRequest() (request *DeleteSecurityPolicyRequest) {
    request = &DeleteSecurityPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("alb", APIVersion, "DeleteSecurityPolicy")
    
    
    return
}

func NewDeleteSecurityPolicyResponse() (response *DeleteSecurityPolicyResponse) {
    response = &DeleteSecurityPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteSecurityPolicy
// 删除一个或多个自定义安全策略。删除安全策略前，请确保该策略未被任何 HTTPS 监听器引用，否则删除操作将失败。
//
// 可能返回的错误码:
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteSecurityPolicy(request *DeleteSecurityPolicyRequest) (response *DeleteSecurityPolicyResponse, err error) {
    return c.DeleteSecurityPolicyWithContext(context.Background(), request)
}

// DeleteSecurityPolicy
// 删除一个或多个自定义安全策略。删除安全策略前，请确保该策略未被任何 HTTPS 监听器引用，否则删除操作将失败。
//
// 可能返回的错误码:
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteSecurityPolicyWithContext(ctx context.Context, request *DeleteSecurityPolicyRequest) (response *DeleteSecurityPolicyResponse, err error) {
    if request == nil {
        request = NewDeleteSecurityPolicyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "alb", APIVersion, "DeleteSecurityPolicy")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteSecurityPolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteSecurityPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteTargetGroupsRequest() (request *DeleteTargetGroupsRequest) {
    request = &DeleteTargetGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("alb", APIVersion, "DeleteTargetGroups")
    
    
    return
}

func NewDeleteTargetGroupsResponse() (response *DeleteTargetGroupsResponse) {
    response = &DeleteTargetGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteTargetGroups
// 删除目标组。
//
// 可能返回的错误码:
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteTargetGroups(request *DeleteTargetGroupsRequest) (response *DeleteTargetGroupsResponse, err error) {
    return c.DeleteTargetGroupsWithContext(context.Background(), request)
}

// DeleteTargetGroups
// 删除目标组。
//
// 可能返回的错误码:
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteTargetGroupsWithContext(ctx context.Context, request *DeleteTargetGroupsRequest) (response *DeleteTargetGroupsResponse, err error) {
    if request == nil {
        request = NewDeleteTargetGroupsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "alb", APIVersion, "DeleteTargetGroups")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteTargetGroups require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteTargetGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAsyncJobsRequest() (request *DescribeAsyncJobsRequest) {
    request = &DescribeAsyncJobsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("alb", APIVersion, "DescribeAsyncJobs")
    
    
    return
}

func NewDescribeAsyncJobsResponse() (response *DescribeAsyncJobsResponse) {
    response = &DescribeAsyncJobsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAsyncJobs
// 异步任务查询接口
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAsyncJobs(request *DescribeAsyncJobsRequest) (response *DescribeAsyncJobsResponse, err error) {
    return c.DescribeAsyncJobsWithContext(context.Background(), request)
}

// DescribeAsyncJobs
// 异步任务查询接口
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAsyncJobsWithContext(ctx context.Context, request *DescribeAsyncJobsRequest) (response *DescribeAsyncJobsResponse, err error) {
    if request == nil {
        request = NewDescribeAsyncJobsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "alb", APIVersion, "DescribeAsyncJobs")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAsyncJobs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAsyncJobsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeHealthCheckTemplatesRequest() (request *DescribeHealthCheckTemplatesRequest) {
    request = &DescribeHealthCheckTemplatesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("alb", APIVersion, "DescribeHealthCheckTemplates")
    
    
    return
}

func NewDescribeHealthCheckTemplatesResponse() (response *DescribeHealthCheckTemplatesResponse) {
    response = &DescribeHealthCheckTemplatesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeHealthCheckTemplates
// 查询健康检查模板列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeHealthCheckTemplates(request *DescribeHealthCheckTemplatesRequest) (response *DescribeHealthCheckTemplatesResponse, err error) {
    return c.DescribeHealthCheckTemplatesWithContext(context.Background(), request)
}

// DescribeHealthCheckTemplates
// 查询健康检查模板列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeHealthCheckTemplatesWithContext(ctx context.Context, request *DescribeHealthCheckTemplatesRequest) (response *DescribeHealthCheckTemplatesResponse, err error) {
    if request == nil {
        request = NewDescribeHealthCheckTemplatesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "alb", APIVersion, "DescribeHealthCheckTemplates")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeHealthCheckTemplates require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeHealthCheckTemplatesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeListenerCertificatesRequest() (request *DescribeListenerCertificatesRequest) {
    request = &DescribeListenerCertificatesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("alb", APIVersion, "DescribeListenerCertificates")
    
    
    return
}

func NewDescribeListenerCertificatesResponse() (response *DescribeListenerCertificatesResponse) {
    response = &DescribeListenerCertificatesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeListenerCertificates
// 根据实例id和监听器id，查询指定监听器绑定的证书列表
//
// 若输入CertificateType为SVR，返回扩展服务器证书与默认服务器证书的信息
//
// 若输入CertificateType为CA，返回默认CA证书的信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeListenerCertificates(request *DescribeListenerCertificatesRequest) (response *DescribeListenerCertificatesResponse, err error) {
    return c.DescribeListenerCertificatesWithContext(context.Background(), request)
}

// DescribeListenerCertificates
// 根据实例id和监听器id，查询指定监听器绑定的证书列表
//
// 若输入CertificateType为SVR，返回扩展服务器证书与默认服务器证书的信息
//
// 若输入CertificateType为CA，返回默认CA证书的信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeListenerCertificatesWithContext(ctx context.Context, request *DescribeListenerCertificatesRequest) (response *DescribeListenerCertificatesResponse, err error) {
    if request == nil {
        request = NewDescribeListenerCertificatesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "alb", APIVersion, "DescribeListenerCertificates")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeListenerCertificates require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeListenerCertificatesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeListenerDetailRequest() (request *DescribeListenerDetailRequest) {
    request = &DescribeListenerDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("alb", APIVersion, "DescribeListenerDetail")
    
    
    return
}

func NewDescribeListenerDetailResponse() (response *DescribeListenerDetailResponse) {
    response = &DescribeListenerDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeListenerDetail
// 查询单个监听器详情
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeListenerDetail(request *DescribeListenerDetailRequest) (response *DescribeListenerDetailResponse, err error) {
    return c.DescribeListenerDetailWithContext(context.Background(), request)
}

// DescribeListenerDetail
// 查询单个监听器详情
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeListenerDetailWithContext(ctx context.Context, request *DescribeListenerDetailRequest) (response *DescribeListenerDetailResponse, err error) {
    if request == nil {
        request = NewDescribeListenerDetailRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "alb", APIVersion, "DescribeListenerDetail")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeListenerDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeListenerDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeListenerHealthStatusRequest() (request *DescribeListenerHealthStatusRequest) {
    request = &DescribeListenerHealthStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("alb", APIVersion, "DescribeListenerHealthStatus")
    
    
    return
}

func NewDescribeListenerHealthStatusResponse() (response *DescribeListenerHealthStatusResponse) {
    response = &DescribeListenerHealthStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeListenerHealthStatus
// 查询监听器健康状态。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeListenerHealthStatus(request *DescribeListenerHealthStatusRequest) (response *DescribeListenerHealthStatusResponse, err error) {
    return c.DescribeListenerHealthStatusWithContext(context.Background(), request)
}

// DescribeListenerHealthStatus
// 查询监听器健康状态。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeListenerHealthStatusWithContext(ctx context.Context, request *DescribeListenerHealthStatusRequest) (response *DescribeListenerHealthStatusResponse, err error) {
    if request == nil {
        request = NewDescribeListenerHealthStatusRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "alb", APIVersion, "DescribeListenerHealthStatus")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeListenerHealthStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeListenerHealthStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeListenersRequest() (request *DescribeListenersRequest) {
    request = &DescribeListenersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("alb", APIVersion, "DescribeListeners")
    
    
    return
}

func NewDescribeListenersResponse() (response *DescribeListenersResponse) {
    response = &DescribeListenersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeListeners
// 查询监听器列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeListeners(request *DescribeListenersRequest) (response *DescribeListenersResponse, err error) {
    return c.DescribeListenersWithContext(context.Background(), request)
}

// DescribeListeners
// 查询监听器列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeListenersWithContext(ctx context.Context, request *DescribeListenersRequest) (response *DescribeListenersResponse, err error) {
    if request == nil {
        request = NewDescribeListenersRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "alb", APIVersion, "DescribeListeners")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeListeners require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeListenersResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLoadBalancerDetailRequest() (request *DescribeLoadBalancerDetailRequest) {
    request = &DescribeLoadBalancerDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("alb", APIVersion, "DescribeLoadBalancerDetail")
    
    
    return
}

func NewDescribeLoadBalancerDetailResponse() (response *DescribeLoadBalancerDetailResponse) {
    response = &DescribeLoadBalancerDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeLoadBalancerDetail
// 查询指定负载均衡实例的详细信息。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeLoadBalancerDetail(request *DescribeLoadBalancerDetailRequest) (response *DescribeLoadBalancerDetailResponse, err error) {
    return c.DescribeLoadBalancerDetailWithContext(context.Background(), request)
}

// DescribeLoadBalancerDetail
// 查询指定负载均衡实例的详细信息。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeLoadBalancerDetailWithContext(ctx context.Context, request *DescribeLoadBalancerDetailRequest) (response *DescribeLoadBalancerDetailResponse, err error) {
    if request == nil {
        request = NewDescribeLoadBalancerDetailRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "alb", APIVersion, "DescribeLoadBalancerDetail")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLoadBalancerDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLoadBalancerDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLoadBalancersRequest() (request *DescribeLoadBalancersRequest) {
    request = &DescribeLoadBalancersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("alb", APIVersion, "DescribeLoadBalancers")
    
    
    return
}

func NewDescribeLoadBalancersResponse() (response *DescribeLoadBalancersResponse) {
    response = &DescribeLoadBalancersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeLoadBalancers
// 查询实例配置。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeLoadBalancers(request *DescribeLoadBalancersRequest) (response *DescribeLoadBalancersResponse, err error) {
    return c.DescribeLoadBalancersWithContext(context.Background(), request)
}

// DescribeLoadBalancers
// 查询实例配置。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeLoadBalancersWithContext(ctx context.Context, request *DescribeLoadBalancersRequest) (response *DescribeLoadBalancersResponse, err error) {
    if request == nil {
        request = NewDescribeLoadBalancersRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "alb", APIVersion, "DescribeLoadBalancers")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLoadBalancers require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLoadBalancersResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeQuotaRequest() (request *DescribeQuotaRequest) {
    request = &DescribeQuotaRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("alb", APIVersion, "DescribeQuota")
    
    
    return
}

func NewDescribeQuotaResponse() (response *DescribeQuotaResponse) {
    response = &DescribeQuotaResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeQuota
// 查询当前账号的 ALB 配额配置。支持按配额类型查询，也支持传入资源ID查询资源级配额；可通过 DisplayFields 按需返回已使用量和剩余可用量。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeQuota(request *DescribeQuotaRequest) (response *DescribeQuotaResponse, err error) {
    return c.DescribeQuotaWithContext(context.Background(), request)
}

// DescribeQuota
// 查询当前账号的 ALB 配额配置。支持按配额类型查询，也支持传入资源ID查询资源级配额；可通过 DisplayFields 按需返回已使用量和剩余可用量。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeQuotaWithContext(ctx context.Context, request *DescribeQuotaRequest) (response *DescribeQuotaResponse, err error) {
    if request == nil {
        request = NewDescribeQuotaRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "alb", APIVersion, "DescribeQuota")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeQuota require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeQuotaResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRulesRequest() (request *DescribeRulesRequest) {
    request = &DescribeRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("alb", APIVersion, "DescribeRules")
    
    
    return
}

func NewDescribeRulesResponse() (response *DescribeRulesResponse) {
    response = &DescribeRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRules
// 查询转发规则
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFIELDTYPE = "InvalidParameter.InvalidFieldType"
//  INVALIDPARAMETER_INVALIDFIELDVALUE = "InvalidParameter.InvalidFieldValue"
//  INVALIDPARAMETER_MISSINGACTION = "InvalidParameter.MissingAction"
//  INVALIDPARAMETER_MISSINGFIELD = "InvalidParameter.MissingField"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_LISTENER = "ResourceNotFound.Listener"
//  RESOURCENOTFOUND_LOADBALANCER = "ResourceNotFound.LoadBalancer"
//  RESOURCENOTFOUND_RULE = "ResourceNotFound.Rule"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_SERVICEROLEUNAUTHORIZED = "UnauthorizedOperation.ServiceRoleUnauthorized"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeRules(request *DescribeRulesRequest) (response *DescribeRulesResponse, err error) {
    return c.DescribeRulesWithContext(context.Background(), request)
}

// DescribeRules
// 查询转发规则
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFIELDTYPE = "InvalidParameter.InvalidFieldType"
//  INVALIDPARAMETER_INVALIDFIELDVALUE = "InvalidParameter.InvalidFieldValue"
//  INVALIDPARAMETER_MISSINGACTION = "InvalidParameter.MissingAction"
//  INVALIDPARAMETER_MISSINGFIELD = "InvalidParameter.MissingField"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_LISTENER = "ResourceNotFound.Listener"
//  RESOURCENOTFOUND_LOADBALANCER = "ResourceNotFound.LoadBalancer"
//  RESOURCENOTFOUND_RULE = "ResourceNotFound.Rule"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_SERVICEROLEUNAUTHORIZED = "UnauthorizedOperation.ServiceRoleUnauthorized"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeRulesWithContext(ctx context.Context, request *DescribeRulesRequest) (response *DescribeRulesResponse, err error) {
    if request == nil {
        request = NewDescribeRulesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "alb", APIVersion, "DescribeRules")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRulesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSecurityPoliciesRequest() (request *DescribeSecurityPoliciesRequest) {
    request = &DescribeSecurityPoliciesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("alb", APIVersion, "DescribeSecurityPolicies")
    
    
    return
}

func NewDescribeSecurityPoliciesResponse() (response *DescribeSecurityPoliciesResponse) {
    response = &DescribeSecurityPoliciesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSecurityPolicies
// 查询自定义安全策略列表，支持按安全策略 ID、名称或标签进行筛选，并支持分页查询。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFIELDVALUE = "InvalidParameter.InvalidFieldValue"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeSecurityPolicies(request *DescribeSecurityPoliciesRequest) (response *DescribeSecurityPoliciesResponse, err error) {
    return c.DescribeSecurityPoliciesWithContext(context.Background(), request)
}

// DescribeSecurityPolicies
// 查询自定义安全策略列表，支持按安全策略 ID、名称或标签进行筛选，并支持分页查询。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFIELDVALUE = "InvalidParameter.InvalidFieldValue"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeSecurityPoliciesWithContext(ctx context.Context, request *DescribeSecurityPoliciesRequest) (response *DescribeSecurityPoliciesResponse, err error) {
    if request == nil {
        request = NewDescribeSecurityPoliciesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "alb", APIVersion, "DescribeSecurityPolicies")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSecurityPolicies require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSecurityPoliciesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSecurityPolicyCapabilitiesRequest() (request *DescribeSecurityPolicyCapabilitiesRequest) {
    request = &DescribeSecurityPolicyCapabilitiesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("alb", APIVersion, "DescribeSecurityPolicyCapabilities")
    
    
    return
}

func NewDescribeSecurityPolicyCapabilitiesResponse() (response *DescribeSecurityPolicyCapabilitiesResponse) {
    response = &DescribeSecurityPolicyCapabilitiesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSecurityPolicyCapabilities
// 查询当前地域支持的安全策略配置能力，包括可选的 TLS 协议版本及各版本对应的加密套件列表。在创建或修改自定义安全策略前，建议先调用此接口获取可用的配置选项。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFIELDVALUE = "InvalidParameter.InvalidFieldValue"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeSecurityPolicyCapabilities(request *DescribeSecurityPolicyCapabilitiesRequest) (response *DescribeSecurityPolicyCapabilitiesResponse, err error) {
    return c.DescribeSecurityPolicyCapabilitiesWithContext(context.Background(), request)
}

// DescribeSecurityPolicyCapabilities
// 查询当前地域支持的安全策略配置能力，包括可选的 TLS 协议版本及各版本对应的加密套件列表。在创建或修改自定义安全策略前，建议先调用此接口获取可用的配置选项。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFIELDVALUE = "InvalidParameter.InvalidFieldValue"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeSecurityPolicyCapabilitiesWithContext(ctx context.Context, request *DescribeSecurityPolicyCapabilitiesRequest) (response *DescribeSecurityPolicyCapabilitiesResponse, err error) {
    if request == nil {
        request = NewDescribeSecurityPolicyCapabilitiesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "alb", APIVersion, "DescribeSecurityPolicyCapabilities")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSecurityPolicyCapabilities require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSecurityPolicyCapabilitiesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSecurityPolicyRelationsRequest() (request *DescribeSecurityPolicyRelationsRequest) {
    request = &DescribeSecurityPolicyRelationsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("alb", APIVersion, "DescribeSecurityPolicyRelations")
    
    
    return
}

func NewDescribeSecurityPolicyRelationsResponse() (response *DescribeSecurityPolicyRelationsResponse) {
    response = &DescribeSecurityPolicyRelationsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSecurityPolicyRelations
// 查询安全策略的关联关系，即安全策略被哪些 HTTPS 监听器引用。在删除或修改安全策略前，建议先调用此接口确认影响范围。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeSecurityPolicyRelations(request *DescribeSecurityPolicyRelationsRequest) (response *DescribeSecurityPolicyRelationsResponse, err error) {
    return c.DescribeSecurityPolicyRelationsWithContext(context.Background(), request)
}

// DescribeSecurityPolicyRelations
// 查询安全策略的关联关系，即安全策略被哪些 HTTPS 监听器引用。在删除或修改安全策略前，建议先调用此接口确认影响范围。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeSecurityPolicyRelationsWithContext(ctx context.Context, request *DescribeSecurityPolicyRelationsRequest) (response *DescribeSecurityPolicyRelationsResponse, err error) {
    if request == nil {
        request = NewDescribeSecurityPolicyRelationsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "alb", APIVersion, "DescribeSecurityPolicyRelations")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSecurityPolicyRelations require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSecurityPolicyRelationsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSystemSecurityPoliciesRequest() (request *DescribeSystemSecurityPoliciesRequest) {
    request = &DescribeSystemSecurityPoliciesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("alb", APIVersion, "DescribeSystemSecurityPolicies")
    
    
    return
}

func NewDescribeSystemSecurityPoliciesResponse() (response *DescribeSystemSecurityPoliciesResponse) {
    response = &DescribeSystemSecurityPoliciesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSystemSecurityPolicies
// 查询系统安全策略。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeSystemSecurityPolicies(request *DescribeSystemSecurityPoliciesRequest) (response *DescribeSystemSecurityPoliciesResponse, err error) {
    return c.DescribeSystemSecurityPoliciesWithContext(context.Background(), request)
}

// DescribeSystemSecurityPolicies
// 查询系统安全策略。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeSystemSecurityPoliciesWithContext(ctx context.Context, request *DescribeSystemSecurityPoliciesRequest) (response *DescribeSystemSecurityPoliciesResponse, err error) {
    if request == nil {
        request = NewDescribeSystemSecurityPoliciesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "alb", APIVersion, "DescribeSystemSecurityPolicies")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSystemSecurityPolicies require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSystemSecurityPoliciesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTargetGroupTargetsRequest() (request *DescribeTargetGroupTargetsRequest) {
    request = &DescribeTargetGroupTargetsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("alb", APIVersion, "DescribeTargetGroupTargets")
    
    
    return
}

func NewDescribeTargetGroupTargetsResponse() (response *DescribeTargetGroupTargetsResponse) {
    response = &DescribeTargetGroupTargetsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTargetGroupTargets
// 查询目标组内后端服务
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeTargetGroupTargets(request *DescribeTargetGroupTargetsRequest) (response *DescribeTargetGroupTargetsResponse, err error) {
    return c.DescribeTargetGroupTargetsWithContext(context.Background(), request)
}

// DescribeTargetGroupTargets
// 查询目标组内后端服务
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeTargetGroupTargetsWithContext(ctx context.Context, request *DescribeTargetGroupTargetsRequest) (response *DescribeTargetGroupTargetsResponse, err error) {
    if request == nil {
        request = NewDescribeTargetGroupTargetsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "alb", APIVersion, "DescribeTargetGroupTargets")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTargetGroupTargets require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTargetGroupTargetsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTargetGroupsRequest() (request *DescribeTargetGroupsRequest) {
    request = &DescribeTargetGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("alb", APIVersion, "DescribeTargetGroups")
    
    
    return
}

func NewDescribeTargetGroupsResponse() (response *DescribeTargetGroupsResponse) {
    response = &DescribeTargetGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTargetGroups
// 查询目标组列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeTargetGroups(request *DescribeTargetGroupsRequest) (response *DescribeTargetGroupsResponse, err error) {
    return c.DescribeTargetGroupsWithContext(context.Background(), request)
}

// DescribeTargetGroups
// 查询目标组列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeTargetGroupsWithContext(ctx context.Context, request *DescribeTargetGroupsRequest) (response *DescribeTargetGroupsResponse, err error) {
    if request == nil {
        request = NewDescribeTargetGroupsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "alb", APIVersion, "DescribeTargetGroups")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTargetGroups require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTargetGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTargetGroupsByTargetRequest() (request *DescribeTargetGroupsByTargetRequest) {
    request = &DescribeTargetGroupsByTargetRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("alb", APIVersion, "DescribeTargetGroupsByTarget")
    
    
    return
}

func NewDescribeTargetGroupsByTargetResponse() (response *DescribeTargetGroupsByTargetResponse) {
    response = &DescribeTargetGroupsByTargetResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTargetGroupsByTarget
// 根据子机查询绑定的目标组
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeTargetGroupsByTarget(request *DescribeTargetGroupsByTargetRequest) (response *DescribeTargetGroupsByTargetResponse, err error) {
    return c.DescribeTargetGroupsByTargetWithContext(context.Background(), request)
}

// DescribeTargetGroupsByTarget
// 根据子机查询绑定的目标组
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeTargetGroupsByTargetWithContext(ctx context.Context, request *DescribeTargetGroupsByTargetRequest) (response *DescribeTargetGroupsByTargetResponse, err error) {
    if request == nil {
        request = NewDescribeTargetGroupsByTargetRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "alb", APIVersion, "DescribeTargetGroupsByTarget")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTargetGroupsByTarget require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTargetGroupsByTargetResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeZonesRequest() (request *DescribeZonesRequest) {
    request = &DescribeZonesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("alb", APIVersion, "DescribeZones")
    
    
    return
}

func NewDescribeZonesResponse() (response *DescribeZonesResponse) {
    response = &DescribeZonesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeZones
// 查询可用区
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeZones(request *DescribeZonesRequest) (response *DescribeZonesResponse, err error) {
    return c.DescribeZonesWithContext(context.Background(), request)
}

// DescribeZones
// 查询可用区
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeZonesWithContext(ctx context.Context, request *DescribeZonesRequest) (response *DescribeZonesResponse, err error) {
    if request == nil {
        request = NewDescribeZonesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "alb", APIVersion, "DescribeZones")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeZones require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeZonesResponse()
    err = c.Send(request, response)
    return
}

func NewDisassociateBandwidthPackageFromLoadBalancerRequest() (request *DisassociateBandwidthPackageFromLoadBalancerRequest) {
    request = &DisassociateBandwidthPackageFromLoadBalancerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("alb", APIVersion, "DisassociateBandwidthPackageFromLoadBalancer")
    
    
    return
}

func NewDisassociateBandwidthPackageFromLoadBalancerResponse() (response *DisassociateBandwidthPackageFromLoadBalancerResponse) {
    response = &DisassociateBandwidthPackageFromLoadBalancerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DisassociateBandwidthPackageFromLoadBalancer
// 将共享带宽包从应用型负载均衡实例解绑。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DisassociateBandwidthPackageFromLoadBalancer(request *DisassociateBandwidthPackageFromLoadBalancerRequest) (response *DisassociateBandwidthPackageFromLoadBalancerResponse, err error) {
    return c.DisassociateBandwidthPackageFromLoadBalancerWithContext(context.Background(), request)
}

// DisassociateBandwidthPackageFromLoadBalancer
// 将共享带宽包从应用型负载均衡实例解绑。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DisassociateBandwidthPackageFromLoadBalancerWithContext(ctx context.Context, request *DisassociateBandwidthPackageFromLoadBalancerRequest) (response *DisassociateBandwidthPackageFromLoadBalancerResponse, err error) {
    if request == nil {
        request = NewDisassociateBandwidthPackageFromLoadBalancerRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "alb", APIVersion, "DisassociateBandwidthPackageFromLoadBalancer")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DisassociateBandwidthPackageFromLoadBalancer require credential")
    }

    request.SetContext(ctx)
    
    response = NewDisassociateBandwidthPackageFromLoadBalancerResponse()
    err = c.Send(request, response)
    return
}

func NewDisassociateListenerAdditionalCertificatesRequest() (request *DisassociateListenerAdditionalCertificatesRequest) {
    request = &DisassociateListenerAdditionalCertificatesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("alb", APIVersion, "DisassociateListenerAdditionalCertificates")
    
    
    return
}

func NewDisassociateListenerAdditionalCertificatesResponse() (response *DisassociateListenerAdditionalCertificatesResponse) {
    response = &DisassociateListenerAdditionalCertificatesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DisassociateListenerAdditionalCertificates
// DisassociateListenerAdditionalCertificates属于异步接口，即系统返回一个请求 ID，但该扩展证书尚未解绑成功，系统后台的解绑任务仍在进行。您可以调用DescribeListenerCertificates接口查询证书的解绑状态：若证书处于Disassociating状态，则证书正在解绑中。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_LISTENER = "ResourceNotFound.Listener"
//  RESOURCENOTFOUND_LOADBALANCER = "ResourceNotFound.LoadBalancer"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_CERTIFICATE = "ResourceUnavailable.Certificate"
//  RESOURCEUNAVAILABLE_LISTENER = "ResourceUnavailable.Listener"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_CERTIFICATENOTBOUND = "UnsupportedOperation.CertificateNotBound"
//  UNSUPPORTEDOPERATION_DISASSOCIATEDEFAULTCERTIFICATE = "UnsupportedOperation.DisassociateDefaultCertificate"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDPROTOCOL = "UnsupportedOperation.UnsupportedProtocol"
func (c *Client) DisassociateListenerAdditionalCertificates(request *DisassociateListenerAdditionalCertificatesRequest) (response *DisassociateListenerAdditionalCertificatesResponse, err error) {
    return c.DisassociateListenerAdditionalCertificatesWithContext(context.Background(), request)
}

// DisassociateListenerAdditionalCertificates
// DisassociateListenerAdditionalCertificates属于异步接口，即系统返回一个请求 ID，但该扩展证书尚未解绑成功，系统后台的解绑任务仍在进行。您可以调用DescribeListenerCertificates接口查询证书的解绑状态：若证书处于Disassociating状态，则证书正在解绑中。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_LISTENER = "ResourceNotFound.Listener"
//  RESOURCENOTFOUND_LOADBALANCER = "ResourceNotFound.LoadBalancer"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_CERTIFICATE = "ResourceUnavailable.Certificate"
//  RESOURCEUNAVAILABLE_LISTENER = "ResourceUnavailable.Listener"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_CERTIFICATENOTBOUND = "UnsupportedOperation.CertificateNotBound"
//  UNSUPPORTEDOPERATION_DISASSOCIATEDEFAULTCERTIFICATE = "UnsupportedOperation.DisassociateDefaultCertificate"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDPROTOCOL = "UnsupportedOperation.UnsupportedProtocol"
func (c *Client) DisassociateListenerAdditionalCertificatesWithContext(ctx context.Context, request *DisassociateListenerAdditionalCertificatesRequest) (response *DisassociateListenerAdditionalCertificatesResponse, err error) {
    if request == nil {
        request = NewDisassociateListenerAdditionalCertificatesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "alb", APIVersion, "DisassociateListenerAdditionalCertificates")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DisassociateListenerAdditionalCertificates require credential")
    }

    request.SetContext(ctx)
    
    response = NewDisassociateListenerAdditionalCertificatesResponse()
    err = c.Send(request, response)
    return
}

func NewInquirePriceCreateLoadBalancerRequest() (request *InquirePriceCreateLoadBalancerRequest) {
    request = &InquirePriceCreateLoadBalancerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("alb", APIVersion, "InquirePriceCreateLoadBalancer")
    
    
    return
}

func NewInquirePriceCreateLoadBalancerResponse() (response *InquirePriceCreateLoadBalancerResponse) {
    response = &InquirePriceCreateLoadBalancerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// InquirePriceCreateLoadBalancer
// InquirePriceCreateLoadBalancer接口查询创建负载均衡的价格。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_LISTENER = "ResourceNotFound.Listener"
//  RESOURCENOTFOUND_LOADBALANCER = "ResourceNotFound.LoadBalancer"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_CERTIFICATE = "ResourceUnavailable.Certificate"
//  RESOURCEUNAVAILABLE_LISTENER = "ResourceUnavailable.Listener"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_CERTIFICATENOTBOUND = "UnsupportedOperation.CertificateNotBound"
//  UNSUPPORTEDOPERATION_DISASSOCIATEDEFAULTCERTIFICATE = "UnsupportedOperation.DisassociateDefaultCertificate"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDPROTOCOL = "UnsupportedOperation.UnsupportedProtocol"
func (c *Client) InquirePriceCreateLoadBalancer(request *InquirePriceCreateLoadBalancerRequest) (response *InquirePriceCreateLoadBalancerResponse, err error) {
    return c.InquirePriceCreateLoadBalancerWithContext(context.Background(), request)
}

// InquirePriceCreateLoadBalancer
// InquirePriceCreateLoadBalancer接口查询创建负载均衡的价格。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_LISTENER = "ResourceNotFound.Listener"
//  RESOURCENOTFOUND_LOADBALANCER = "ResourceNotFound.LoadBalancer"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_CERTIFICATE = "ResourceUnavailable.Certificate"
//  RESOURCEUNAVAILABLE_LISTENER = "ResourceUnavailable.Listener"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_CERTIFICATENOTBOUND = "UnsupportedOperation.CertificateNotBound"
//  UNSUPPORTEDOPERATION_DISASSOCIATEDEFAULTCERTIFICATE = "UnsupportedOperation.DisassociateDefaultCertificate"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDPROTOCOL = "UnsupportedOperation.UnsupportedProtocol"
func (c *Client) InquirePriceCreateLoadBalancerWithContext(ctx context.Context, request *InquirePriceCreateLoadBalancerRequest) (response *InquirePriceCreateLoadBalancerResponse, err error) {
    if request == nil {
        request = NewInquirePriceCreateLoadBalancerRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "alb", APIVersion, "InquirePriceCreateLoadBalancer")
    
    if c.GetCredential() == nil {
        return nil, errors.New("InquirePriceCreateLoadBalancer require credential")
    }

    request.SetContext(ctx)
    
    response = NewInquirePriceCreateLoadBalancerResponse()
    err = c.Send(request, response)
    return
}

func NewModifyHealthCheckTemplateRequest() (request *ModifyHealthCheckTemplateRequest) {
    request = &ModifyHealthCheckTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("alb", APIVersion, "ModifyHealthCheckTemplate")
    
    
    return
}

func NewModifyHealthCheckTemplateResponse() (response *ModifyHealthCheckTemplateResponse) {
    response = &ModifyHealthCheckTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyHealthCheckTemplate
// 修改健康检查模板
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyHealthCheckTemplate(request *ModifyHealthCheckTemplateRequest) (response *ModifyHealthCheckTemplateResponse, err error) {
    return c.ModifyHealthCheckTemplateWithContext(context.Background(), request)
}

// ModifyHealthCheckTemplate
// 修改健康检查模板
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyHealthCheckTemplateWithContext(ctx context.Context, request *ModifyHealthCheckTemplateRequest) (response *ModifyHealthCheckTemplateResponse, err error) {
    if request == nil {
        request = NewModifyHealthCheckTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "alb", APIVersion, "ModifyHealthCheckTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyHealthCheckTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyHealthCheckTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewModifyListenerAttributesRequest() (request *ModifyListenerAttributesRequest) {
    request = &ModifyListenerAttributesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("alb", APIVersion, "ModifyListenerAttributes")
    
    
    return
}

func NewModifyListenerAttributesResponse() (response *ModifyListenerAttributesResponse) {
    response = &ModifyListenerAttributesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyListenerAttributes
// 修改监听器属性
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyListenerAttributes(request *ModifyListenerAttributesRequest) (response *ModifyListenerAttributesResponse, err error) {
    return c.ModifyListenerAttributesWithContext(context.Background(), request)
}

// ModifyListenerAttributes
// 修改监听器属性
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyListenerAttributesWithContext(ctx context.Context, request *ModifyListenerAttributesRequest) (response *ModifyListenerAttributesResponse, err error) {
    if request == nil {
        request = NewModifyListenerAttributesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "alb", APIVersion, "ModifyListenerAttributes")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyListenerAttributes require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyListenerAttributesResponse()
    err = c.Send(request, response)
    return
}

func NewModifyLoadBalancerAddressTypeRequest() (request *ModifyLoadBalancerAddressTypeRequest) {
    request = &ModifyLoadBalancerAddressTypeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("alb", APIVersion, "ModifyLoadBalancerAddressType")
    
    
    return
}

func NewModifyLoadBalancerAddressTypeResponse() (response *ModifyLoadBalancerAddressTypeResponse) {
    response = &ModifyLoadBalancerAddressTypeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyLoadBalancerAddressType
// **前提条件：**
//
// 您已经创建应用型负载均衡实例，具体操作，请参见 CreateLoadBalancer 。
//
// 当您需要通过此接口将应用型负载均衡实例的网络类型由私网变为公网时，您需要先创建一个弹性公网 IP。
//
// **使用说明：**
//
// ModifyLoadBalancerAddressType 接口属于异步接口，即系统返回一个请求 ID，但该应用型负载均衡实例的网络类型尚未变更成功，系统后台的变更任务仍在进行。您可以调用 DescribeLoadBalancerDetail 查询应用型负载均衡实例的网络类型的变更状态：
//
// 当应用型负载均衡实例处于 Configuring 状态时，表示应用型负载均衡实例的网络类型正在变更中。
//
// 当应用型负载均衡实例处于 Active 状态时，表示应用型负载均衡实例的网络类型变更成功。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyLoadBalancerAddressType(request *ModifyLoadBalancerAddressTypeRequest) (response *ModifyLoadBalancerAddressTypeResponse, err error) {
    return c.ModifyLoadBalancerAddressTypeWithContext(context.Background(), request)
}

// ModifyLoadBalancerAddressType
// **前提条件：**
//
// 您已经创建应用型负载均衡实例，具体操作，请参见 CreateLoadBalancer 。
//
// 当您需要通过此接口将应用型负载均衡实例的网络类型由私网变为公网时，您需要先创建一个弹性公网 IP。
//
// **使用说明：**
//
// ModifyLoadBalancerAddressType 接口属于异步接口，即系统返回一个请求 ID，但该应用型负载均衡实例的网络类型尚未变更成功，系统后台的变更任务仍在进行。您可以调用 DescribeLoadBalancerDetail 查询应用型负载均衡实例的网络类型的变更状态：
//
// 当应用型负载均衡实例处于 Configuring 状态时，表示应用型负载均衡实例的网络类型正在变更中。
//
// 当应用型负载均衡实例处于 Active 状态时，表示应用型负载均衡实例的网络类型变更成功。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyLoadBalancerAddressTypeWithContext(ctx context.Context, request *ModifyLoadBalancerAddressTypeRequest) (response *ModifyLoadBalancerAddressTypeResponse, err error) {
    if request == nil {
        request = NewModifyLoadBalancerAddressTypeRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "alb", APIVersion, "ModifyLoadBalancerAddressType")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyLoadBalancerAddressType require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyLoadBalancerAddressTypeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyLoadBalancerAttributesRequest() (request *ModifyLoadBalancerAttributesRequest) {
    request = &ModifyLoadBalancerAttributesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("alb", APIVersion, "ModifyLoadBalancerAttributes")
    
    
    return
}

func NewModifyLoadBalancerAttributesResponse() (response *ModifyLoadBalancerAttributesResponse) {
    response = &ModifyLoadBalancerAttributesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyLoadBalancerAttributes
// **ModifyLoadBalancerAttributes**接口属于异步接口，即系统返回一个请求ID，但该应用型负载均衡实例属性尚未修改成功，系统后台的修改任务仍在进行。您可以调用[DescribeLoadBalancerAttribute](214362)查询应用型负载均衡实例属性的修改状态：
//
// - 当应用型负载均衡实例属性处于**Configuring**状态时，表示应用型负载均衡实例属性正在修改中。
//
// - 当应用型负载均衡实例属性处于**Active**状态时，表示应用型负载均衡实例属性修改成功。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyLoadBalancerAttributes(request *ModifyLoadBalancerAttributesRequest) (response *ModifyLoadBalancerAttributesResponse, err error) {
    return c.ModifyLoadBalancerAttributesWithContext(context.Background(), request)
}

// ModifyLoadBalancerAttributes
// **ModifyLoadBalancerAttributes**接口属于异步接口，即系统返回一个请求ID，但该应用型负载均衡实例属性尚未修改成功，系统后台的修改任务仍在进行。您可以调用[DescribeLoadBalancerAttribute](214362)查询应用型负载均衡实例属性的修改状态：
//
// - 当应用型负载均衡实例属性处于**Configuring**状态时，表示应用型负载均衡实例属性正在修改中。
//
// - 当应用型负载均衡实例属性处于**Active**状态时，表示应用型负载均衡实例属性修改成功。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyLoadBalancerAttributesWithContext(ctx context.Context, request *ModifyLoadBalancerAttributesRequest) (response *ModifyLoadBalancerAttributesResponse, err error) {
    if request == nil {
        request = NewModifyLoadBalancerAttributesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "alb", APIVersion, "ModifyLoadBalancerAttributes")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyLoadBalancerAttributes require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyLoadBalancerAttributesResponse()
    err = c.Send(request, response)
    return
}

func NewModifyLoadBalancerModificationProtectionRequest() (request *ModifyLoadBalancerModificationProtectionRequest) {
    request = &ModifyLoadBalancerModificationProtectionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("alb", APIVersion, "ModifyLoadBalancerModificationProtection")
    
    
    return
}

func NewModifyLoadBalancerModificationProtectionResponse() (response *ModifyLoadBalancerModificationProtectionResponse) {
    response = &ModifyLoadBalancerModificationProtectionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyLoadBalancerModificationProtection
// 设置负载均衡实例修改保护。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  REGIONERROR = "RegionError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyLoadBalancerModificationProtection(request *ModifyLoadBalancerModificationProtectionRequest) (response *ModifyLoadBalancerModificationProtectionResponse, err error) {
    return c.ModifyLoadBalancerModificationProtectionWithContext(context.Background(), request)
}

// ModifyLoadBalancerModificationProtection
// 设置负载均衡实例修改保护。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  REGIONERROR = "RegionError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyLoadBalancerModificationProtectionWithContext(ctx context.Context, request *ModifyLoadBalancerModificationProtectionRequest) (response *ModifyLoadBalancerModificationProtectionResponse, err error) {
    if request == nil {
        request = NewModifyLoadBalancerModificationProtectionRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "alb", APIVersion, "ModifyLoadBalancerModificationProtection")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyLoadBalancerModificationProtection require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyLoadBalancerModificationProtectionResponse()
    err = c.Send(request, response)
    return
}

func NewModifyRulesAttributesRequest() (request *ModifyRulesAttributesRequest) {
    request = &ModifyRulesAttributesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("alb", APIVersion, "ModifyRulesAttributes")
    
    
    return
}

func NewModifyRulesAttributesResponse() (response *ModifyRulesAttributesResponse) {
    response = &ModifyRulesAttributesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyRulesAttributes
// ModifyRulesAttributes修改转发规则属性，本接口为异步接口，返回成功后需以返回的RequestID为入参，调用DescribeAsyncJobs接口查询本次任务是否成功。
//
// 一条规则最多支持10个转发条件（Conditions），5个转发动作（Actions）。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDCLIENTTOKEN = "InvalidParameter.InvalidClientToken"
//  INVALIDPARAMETER_INVALIDFIELDTYPE = "InvalidParameter.InvalidFieldType"
//  INVALIDPARAMETER_INVALIDFIELDVALUE = "InvalidParameter.InvalidFieldValue"
//  INVALIDPARAMETER_INVALIDQUOTACHECKREQUEST = "InvalidParameter.InvalidQuotaCheckRequest"
//  INVALIDPARAMETER_MISSINGACTION = "InvalidParameter.MissingAction"
//  INVALIDPARAMETER_MISSINGFIELD = "InvalidParameter.MissingField"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_LISTENER = "ResourceNotFound.Listener"
//  RESOURCENOTFOUND_LOADBALANCER = "ResourceNotFound.LoadBalancer"
//  RESOURCENOTFOUND_RULE = "ResourceNotFound.Rule"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_LISTENER = "ResourceUnavailable.Listener"
//  RESOURCEUNAVAILABLE_LOADBALANCER = "ResourceUnavailable.LoadBalancer"
//  RESOURCEUNAVAILABLE_RULE = "ResourceUnavailable.Rule"
//  RESOURCEUNAVAILABLE_TARGETGROUP = "ResourceUnavailable.TargetGroup"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyRulesAttributes(request *ModifyRulesAttributesRequest) (response *ModifyRulesAttributesResponse, err error) {
    return c.ModifyRulesAttributesWithContext(context.Background(), request)
}

// ModifyRulesAttributes
// ModifyRulesAttributes修改转发规则属性，本接口为异步接口，返回成功后需以返回的RequestID为入参，调用DescribeAsyncJobs接口查询本次任务是否成功。
//
// 一条规则最多支持10个转发条件（Conditions），5个转发动作（Actions）。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDCLIENTTOKEN = "InvalidParameter.InvalidClientToken"
//  INVALIDPARAMETER_INVALIDFIELDTYPE = "InvalidParameter.InvalidFieldType"
//  INVALIDPARAMETER_INVALIDFIELDVALUE = "InvalidParameter.InvalidFieldValue"
//  INVALIDPARAMETER_INVALIDQUOTACHECKREQUEST = "InvalidParameter.InvalidQuotaCheckRequest"
//  INVALIDPARAMETER_MISSINGACTION = "InvalidParameter.MissingAction"
//  INVALIDPARAMETER_MISSINGFIELD = "InvalidParameter.MissingField"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_LISTENER = "ResourceNotFound.Listener"
//  RESOURCENOTFOUND_LOADBALANCER = "ResourceNotFound.LoadBalancer"
//  RESOURCENOTFOUND_RULE = "ResourceNotFound.Rule"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_LISTENER = "ResourceUnavailable.Listener"
//  RESOURCEUNAVAILABLE_LOADBALANCER = "ResourceUnavailable.LoadBalancer"
//  RESOURCEUNAVAILABLE_RULE = "ResourceUnavailable.Rule"
//  RESOURCEUNAVAILABLE_TARGETGROUP = "ResourceUnavailable.TargetGroup"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyRulesAttributesWithContext(ctx context.Context, request *ModifyRulesAttributesRequest) (response *ModifyRulesAttributesResponse, err error) {
    if request == nil {
        request = NewModifyRulesAttributesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "alb", APIVersion, "ModifyRulesAttributes")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyRulesAttributes require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyRulesAttributesResponse()
    err = c.Send(request, response)
    return
}

func NewModifySecurityPolicyAttributesRequest() (request *ModifySecurityPolicyAttributesRequest) {
    request = &ModifySecurityPolicyAttributesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("alb", APIVersion, "ModifySecurityPolicyAttributes")
    
    
    return
}

func NewModifySecurityPolicyAttributesResponse() (response *ModifySecurityPolicyAttributesResponse) {
    response = &ModifySecurityPolicyAttributesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifySecurityPolicyAttributes
// 修改自定义安全策略的属性，包括策略名称、TLS 协议版本和加密套件。修改后的配置将立即应用到所有关联该策略的 HTTPS 监听器。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifySecurityPolicyAttributes(request *ModifySecurityPolicyAttributesRequest) (response *ModifySecurityPolicyAttributesResponse, err error) {
    return c.ModifySecurityPolicyAttributesWithContext(context.Background(), request)
}

// ModifySecurityPolicyAttributes
// 修改自定义安全策略的属性，包括策略名称、TLS 协议版本和加密套件。修改后的配置将立即应用到所有关联该策略的 HTTPS 监听器。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifySecurityPolicyAttributesWithContext(ctx context.Context, request *ModifySecurityPolicyAttributesRequest) (response *ModifySecurityPolicyAttributesResponse, err error) {
    if request == nil {
        request = NewModifySecurityPolicyAttributesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "alb", APIVersion, "ModifySecurityPolicyAttributes")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifySecurityPolicyAttributes require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifySecurityPolicyAttributesResponse()
    err = c.Send(request, response)
    return
}

func NewModifyTargetGroupAttributesRequest() (request *ModifyTargetGroupAttributesRequest) {
    request = &ModifyTargetGroupAttributesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("alb", APIVersion, "ModifyTargetGroupAttributes")
    
    
    return
}

func NewModifyTargetGroupAttributesResponse() (response *ModifyTargetGroupAttributesResponse) {
    response = &ModifyTargetGroupAttributesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyTargetGroupAttributes
// 修改目标组。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyTargetGroupAttributes(request *ModifyTargetGroupAttributesRequest) (response *ModifyTargetGroupAttributesResponse, err error) {
    return c.ModifyTargetGroupAttributesWithContext(context.Background(), request)
}

// ModifyTargetGroupAttributes
// 修改目标组。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyTargetGroupAttributesWithContext(ctx context.Context, request *ModifyTargetGroupAttributesRequest) (response *ModifyTargetGroupAttributesResponse, err error) {
    if request == nil {
        request = NewModifyTargetGroupAttributesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "alb", APIVersion, "ModifyTargetGroupAttributes")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyTargetGroupAttributes require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyTargetGroupAttributesResponse()
    err = c.Send(request, response)
    return
}

func NewModifyTargetsInTargetGroupRequest() (request *ModifyTargetsInTargetGroupRequest) {
    request = &ModifyTargetsInTargetGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("alb", APIVersion, "ModifyTargetsInTargetGroup")
    
    
    return
}

func NewModifyTargetsInTargetGroupResponse() (response *ModifyTargetsInTargetGroupResponse) {
    response = &ModifyTargetsInTargetGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyTargetsInTargetGroup
// 修改目标组内后端服务信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyTargetsInTargetGroup(request *ModifyTargetsInTargetGroupRequest) (response *ModifyTargetsInTargetGroupResponse, err error) {
    return c.ModifyTargetsInTargetGroupWithContext(context.Background(), request)
}

// ModifyTargetsInTargetGroup
// 修改目标组内后端服务信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyTargetsInTargetGroupWithContext(ctx context.Context, request *ModifyTargetsInTargetGroupRequest) (response *ModifyTargetsInTargetGroupResponse, err error) {
    if request == nil {
        request = NewModifyTargetsInTargetGroupRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "alb", APIVersion, "ModifyTargetsInTargetGroup")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyTargetsInTargetGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyTargetsInTargetGroupResponse()
    err = c.Send(request, response)
    return
}

func NewNotifyUnbindTargetRequest() (request *NotifyUnbindTargetRequest) {
    request = &NotifyUnbindTargetRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("alb", APIVersion, "NotifyUnbindTarget")
    
    
    return
}

func NewNotifyUnbindTargetResponse() (response *NotifyUnbindTargetResponse) {
    response = &NotifyUnbindTargetResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// NotifyUnbindTarget
// 通知负载均衡解绑后端服务
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) NotifyUnbindTarget(request *NotifyUnbindTargetRequest) (response *NotifyUnbindTargetResponse, err error) {
    return c.NotifyUnbindTargetWithContext(context.Background(), request)
}

// NotifyUnbindTarget
// 通知负载均衡解绑后端服务
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) NotifyUnbindTargetWithContext(ctx context.Context, request *NotifyUnbindTargetRequest) (response *NotifyUnbindTargetResponse, err error) {
    if request == nil {
        request = NewNotifyUnbindTargetRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "alb", APIVersion, "NotifyUnbindTarget")
    
    if c.GetCredential() == nil {
        return nil, errors.New("NotifyUnbindTarget require credential")
    }

    request.SetContext(ctx)
    
    response = NewNotifyUnbindTargetResponse()
    err = c.Send(request, response)
    return
}

func NewRemoveTargetsFromTargetGroupRequest() (request *RemoveTargetsFromTargetGroupRequest) {
    request = &RemoveTargetsFromTargetGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("alb", APIVersion, "RemoveTargetsFromTargetGroup")
    
    
    return
}

func NewRemoveTargetsFromTargetGroupResponse() (response *RemoveTargetsFromTargetGroupResponse) {
    response = &RemoveTargetsFromTargetGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RemoveTargetsFromTargetGroup
// 从目标组内移除后端服务
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) RemoveTargetsFromTargetGroup(request *RemoveTargetsFromTargetGroupRequest) (response *RemoveTargetsFromTargetGroupResponse, err error) {
    return c.RemoveTargetsFromTargetGroupWithContext(context.Background(), request)
}

// RemoveTargetsFromTargetGroup
// 从目标组内移除后端服务
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) RemoveTargetsFromTargetGroupWithContext(ctx context.Context, request *RemoveTargetsFromTargetGroupRequest) (response *RemoveTargetsFromTargetGroupResponse, err error) {
    if request == nil {
        request = NewRemoveTargetsFromTargetGroupRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "alb", APIVersion, "RemoveTargetsFromTargetGroup")
    
    if c.GetCredential() == nil {
        return nil, errors.New("RemoveTargetsFromTargetGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewRemoveTargetsFromTargetGroupResponse()
    err = c.Send(request, response)
    return
}

func NewSetLoadBalancerSecurityGroupsRequest() (request *SetLoadBalancerSecurityGroupsRequest) {
    request = &SetLoadBalancerSecurityGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("alb", APIVersion, "SetLoadBalancerSecurityGroups")
    
    
    return
}

func NewSetLoadBalancerSecurityGroupsResponse() (response *SetLoadBalancerSecurityGroupsResponse) {
    response = &SetLoadBalancerSecurityGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SetLoadBalancerSecurityGroups
// SetLoadBalancerSecurityGroups 接口支持对一个公网负载均衡实例执行设置（绑定、解绑）安全组操作。查询一个负载均衡实例目前已绑定的安全组，可使用 [DescribeLoadBalancerDetail](xxx) 接口。本接口是set语义，
//
// 绑定操作时，入参需要传入负载均衡实例要绑定的所有安全组（已绑定的+新增绑定的）。
//
// 解绑操作时，入参需要传入负载均衡实例执行解绑后所绑定的所有安全组；如果要解绑所有安全组，可不传此参数，或传入空数组。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) SetLoadBalancerSecurityGroups(request *SetLoadBalancerSecurityGroupsRequest) (response *SetLoadBalancerSecurityGroupsResponse, err error) {
    return c.SetLoadBalancerSecurityGroupsWithContext(context.Background(), request)
}

// SetLoadBalancerSecurityGroups
// SetLoadBalancerSecurityGroups 接口支持对一个公网负载均衡实例执行设置（绑定、解绑）安全组操作。查询一个负载均衡实例目前已绑定的安全组，可使用 [DescribeLoadBalancerDetail](xxx) 接口。本接口是set语义，
//
// 绑定操作时，入参需要传入负载均衡实例要绑定的所有安全组（已绑定的+新增绑定的）。
//
// 解绑操作时，入参需要传入负载均衡实例执行解绑后所绑定的所有安全组；如果要解绑所有安全组，可不传此参数，或传入空数组。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) SetLoadBalancerSecurityGroupsWithContext(ctx context.Context, request *SetLoadBalancerSecurityGroupsRequest) (response *SetLoadBalancerSecurityGroupsResponse, err error) {
    if request == nil {
        request = NewSetLoadBalancerSecurityGroupsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "alb", APIVersion, "SetLoadBalancerSecurityGroups")
    
    if c.GetCredential() == nil {
        return nil, errors.New("SetLoadBalancerSecurityGroups require credential")
    }

    request.SetContext(ctx)
    
    response = NewSetLoadBalancerSecurityGroupsResponse()
    err = c.Send(request, response)
    return
}
