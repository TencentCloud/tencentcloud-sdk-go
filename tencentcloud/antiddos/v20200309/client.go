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

package v20200309

import (
    "context"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2020-03-09"

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


func NewAssociateDDoSEipAddressRequest() (request *AssociateDDoSEipAddressRequest) {
    request = &AssociateDDoSEipAddressRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("antiddos", APIVersion, "AssociateDDoSEipAddress")
    
    
    return
}

func NewAssociateDDoSEipAddressResponse() (response *AssociateDDoSEipAddressResponse) {
    response = &AssociateDDoSEipAddressResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AssociateDDoSEipAddress
// 本接口 (AssociateDDoSEipAddress) 用于将高防弹性公网IP绑定到实例或弹性网卡的指定内网 IP 上。
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
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) AssociateDDoSEipAddress(request *AssociateDDoSEipAddressRequest) (response *AssociateDDoSEipAddressResponse, err error) {
    if request == nil {
        request = NewAssociateDDoSEipAddressRequest()
    }
    
    response = NewAssociateDDoSEipAddressResponse()
    err = c.Send(request, response)
    return
}

// AssociateDDoSEipAddress
// 本接口 (AssociateDDoSEipAddress) 用于将高防弹性公网IP绑定到实例或弹性网卡的指定内网 IP 上。
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
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) AssociateDDoSEipAddressWithContext(ctx context.Context, request *AssociateDDoSEipAddressRequest) (response *AssociateDDoSEipAddressResponse, err error) {
    if request == nil {
        request = NewAssociateDDoSEipAddressRequest()
    }
    request.SetContext(ctx)
    
    response = NewAssociateDDoSEipAddressResponse()
    err = c.Send(request, response)
    return
}

func NewAssociateDDoSEipLoadBalancerRequest() (request *AssociateDDoSEipLoadBalancerRequest) {
    request = &AssociateDDoSEipLoadBalancerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("antiddos", APIVersion, "AssociateDDoSEipLoadBalancer")
    
    
    return
}

func NewAssociateDDoSEipLoadBalancerResponse() (response *AssociateDDoSEipLoadBalancerResponse) {
    response = &AssociateDDoSEipLoadBalancerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AssociateDDoSEipLoadBalancer
// 本接口 (AssociateDDoSEipLoadBalancer) 用于将高防弹性公网IP绑定到负载均衡指定内网 IP 上。
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
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) AssociateDDoSEipLoadBalancer(request *AssociateDDoSEipLoadBalancerRequest) (response *AssociateDDoSEipLoadBalancerResponse, err error) {
    if request == nil {
        request = NewAssociateDDoSEipLoadBalancerRequest()
    }
    
    response = NewAssociateDDoSEipLoadBalancerResponse()
    err = c.Send(request, response)
    return
}

// AssociateDDoSEipLoadBalancer
// 本接口 (AssociateDDoSEipLoadBalancer) 用于将高防弹性公网IP绑定到负载均衡指定内网 IP 上。
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
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) AssociateDDoSEipLoadBalancerWithContext(ctx context.Context, request *AssociateDDoSEipLoadBalancerRequest) (response *AssociateDDoSEipLoadBalancerResponse, err error) {
    if request == nil {
        request = NewAssociateDDoSEipLoadBalancerRequest()
    }
    request.SetContext(ctx)
    
    response = NewAssociateDDoSEipLoadBalancerResponse()
    err = c.Send(request, response)
    return
}

func NewCreateBlackWhiteIpListRequest() (request *CreateBlackWhiteIpListRequest) {
    request = &CreateBlackWhiteIpListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("antiddos", APIVersion, "CreateBlackWhiteIpList")
    
    
    return
}

func NewCreateBlackWhiteIpListResponse() (response *CreateBlackWhiteIpListResponse) {
    response = &CreateBlackWhiteIpListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateBlackWhiteIpList
// 添加DDoS防护的IP黑白名单
//
// 可能返回的错误码:
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateBlackWhiteIpList(request *CreateBlackWhiteIpListRequest) (response *CreateBlackWhiteIpListResponse, err error) {
    if request == nil {
        request = NewCreateBlackWhiteIpListRequest()
    }
    
    response = NewCreateBlackWhiteIpListResponse()
    err = c.Send(request, response)
    return
}

// CreateBlackWhiteIpList
// 添加DDoS防护的IP黑白名单
//
// 可能返回的错误码:
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateBlackWhiteIpListWithContext(ctx context.Context, request *CreateBlackWhiteIpListRequest) (response *CreateBlackWhiteIpListResponse, err error) {
    if request == nil {
        request = NewCreateBlackWhiteIpListRequest()
    }
    request.SetContext(ctx)
    
    response = NewCreateBlackWhiteIpListResponse()
    err = c.Send(request, response)
    return
}

func NewCreateBoundIPRequest() (request *CreateBoundIPRequest) {
    request = &CreateBoundIPRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("antiddos", APIVersion, "CreateBoundIP")
    
    
    return
}

func NewCreateBoundIPResponse() (response *CreateBoundIPResponse) {
    response = &CreateBoundIPResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateBoundIP
// 绑定IP到高防包实例，支持独享包、共享包；需要注意的是此接口绑定或解绑IP是异步接口，当处于绑定或解绑中时，则不允许再进行绑定或解绑，需要等待当前绑定或解绑完成。
//
// 可能返回的错误码:
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateBoundIP(request *CreateBoundIPRequest) (response *CreateBoundIPResponse, err error) {
    if request == nil {
        request = NewCreateBoundIPRequest()
    }
    
    response = NewCreateBoundIPResponse()
    err = c.Send(request, response)
    return
}

// CreateBoundIP
// 绑定IP到高防包实例，支持独享包、共享包；需要注意的是此接口绑定或解绑IP是异步接口，当处于绑定或解绑中时，则不允许再进行绑定或解绑，需要等待当前绑定或解绑完成。
//
// 可能返回的错误码:
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateBoundIPWithContext(ctx context.Context, request *CreateBoundIPRequest) (response *CreateBoundIPResponse, err error) {
    if request == nil {
        request = NewCreateBoundIPRequest()
    }
    request.SetContext(ctx)
    
    response = NewCreateBoundIPResponse()
    err = c.Send(request, response)
    return
}

func NewCreateCCPrecisionPolicyRequest() (request *CreateCCPrecisionPolicyRequest) {
    request = &CreateCCPrecisionPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("antiddos", APIVersion, "CreateCCPrecisionPolicy")
    
    
    return
}

func NewCreateCCPrecisionPolicyResponse() (response *CreateCCPrecisionPolicyResponse) {
    response = &CreateCCPrecisionPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateCCPrecisionPolicy
// 新增CC精准防护策略
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
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateCCPrecisionPolicy(request *CreateCCPrecisionPolicyRequest) (response *CreateCCPrecisionPolicyResponse, err error) {
    if request == nil {
        request = NewCreateCCPrecisionPolicyRequest()
    }
    
    response = NewCreateCCPrecisionPolicyResponse()
    err = c.Send(request, response)
    return
}

// CreateCCPrecisionPolicy
// 新增CC精准防护策略
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
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateCCPrecisionPolicyWithContext(ctx context.Context, request *CreateCCPrecisionPolicyRequest) (response *CreateCCPrecisionPolicyResponse, err error) {
    if request == nil {
        request = NewCreateCCPrecisionPolicyRequest()
    }
    request.SetContext(ctx)
    
    response = NewCreateCCPrecisionPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewCreateCCReqLimitPolicyRequest() (request *CreateCCReqLimitPolicyRequest) {
    request = &CreateCCReqLimitPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("antiddos", APIVersion, "CreateCCReqLimitPolicy")
    
    
    return
}

func NewCreateCCReqLimitPolicyResponse() (response *CreateCCReqLimitPolicyResponse) {
    response = &CreateCCReqLimitPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateCCReqLimitPolicy
// 新增CC频率限制策略
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
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateCCReqLimitPolicy(request *CreateCCReqLimitPolicyRequest) (response *CreateCCReqLimitPolicyResponse, err error) {
    if request == nil {
        request = NewCreateCCReqLimitPolicyRequest()
    }
    
    response = NewCreateCCReqLimitPolicyResponse()
    err = c.Send(request, response)
    return
}

// CreateCCReqLimitPolicy
// 新增CC频率限制策略
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
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateCCReqLimitPolicyWithContext(ctx context.Context, request *CreateCCReqLimitPolicyRequest) (response *CreateCCReqLimitPolicyResponse, err error) {
    if request == nil {
        request = NewCreateCCReqLimitPolicyRequest()
    }
    request.SetContext(ctx)
    
    response = NewCreateCCReqLimitPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewCreateCcGeoIPBlockConfigRequest() (request *CreateCcGeoIPBlockConfigRequest) {
    request = &CreateCcGeoIPBlockConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("antiddos", APIVersion, "CreateCcGeoIPBlockConfig")
    
    
    return
}

func NewCreateCcGeoIPBlockConfigResponse() (response *CreateCcGeoIPBlockConfigResponse) {
    response = &CreateCcGeoIPBlockConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateCcGeoIPBlockConfig
// 新建cc防护的地域封禁配置
//
// 可能返回的错误码:
//  LIMITEXCEEDED = "LimitExceeded"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateCcGeoIPBlockConfig(request *CreateCcGeoIPBlockConfigRequest) (response *CreateCcGeoIPBlockConfigResponse, err error) {
    if request == nil {
        request = NewCreateCcGeoIPBlockConfigRequest()
    }
    
    response = NewCreateCcGeoIPBlockConfigResponse()
    err = c.Send(request, response)
    return
}

// CreateCcGeoIPBlockConfig
// 新建cc防护的地域封禁配置
//
// 可能返回的错误码:
//  LIMITEXCEEDED = "LimitExceeded"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateCcGeoIPBlockConfigWithContext(ctx context.Context, request *CreateCcGeoIPBlockConfigRequest) (response *CreateCcGeoIPBlockConfigResponse, err error) {
    if request == nil {
        request = NewCreateCcGeoIPBlockConfigRequest()
    }
    request.SetContext(ctx)
    
    response = NewCreateCcGeoIPBlockConfigResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDDoSAIRequest() (request *CreateDDoSAIRequest) {
    request = &CreateDDoSAIRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("antiddos", APIVersion, "CreateDDoSAI")
    
    
    return
}

func NewCreateDDoSAIResponse() (response *CreateDDoSAIResponse) {
    response = &CreateDDoSAIResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateDDoSAI
// 设置DDoS防护的AI防护开关
//
// 可能返回的错误码:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateDDoSAI(request *CreateDDoSAIRequest) (response *CreateDDoSAIResponse, err error) {
    if request == nil {
        request = NewCreateDDoSAIRequest()
    }
    
    response = NewCreateDDoSAIResponse()
    err = c.Send(request, response)
    return
}

// CreateDDoSAI
// 设置DDoS防护的AI防护开关
//
// 可能返回的错误码:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateDDoSAIWithContext(ctx context.Context, request *CreateDDoSAIRequest) (response *CreateDDoSAIResponse, err error) {
    if request == nil {
        request = NewCreateDDoSAIRequest()
    }
    request.SetContext(ctx)
    
    response = NewCreateDDoSAIResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDDoSBlackWhiteIpListRequest() (request *CreateDDoSBlackWhiteIpListRequest) {
    request = &CreateDDoSBlackWhiteIpListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("antiddos", APIVersion, "CreateDDoSBlackWhiteIpList")
    
    
    return
}

func NewCreateDDoSBlackWhiteIpListResponse() (response *CreateDDoSBlackWhiteIpListResponse) {
    response = &CreateDDoSBlackWhiteIpListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateDDoSBlackWhiteIpList
// 添加DDoS防护的IP网段黑白名单
//
// 可能返回的错误码:
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateDDoSBlackWhiteIpList(request *CreateDDoSBlackWhiteIpListRequest) (response *CreateDDoSBlackWhiteIpListResponse, err error) {
    if request == nil {
        request = NewCreateDDoSBlackWhiteIpListRequest()
    }
    
    response = NewCreateDDoSBlackWhiteIpListResponse()
    err = c.Send(request, response)
    return
}

// CreateDDoSBlackWhiteIpList
// 添加DDoS防护的IP网段黑白名单
//
// 可能返回的错误码:
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateDDoSBlackWhiteIpListWithContext(ctx context.Context, request *CreateDDoSBlackWhiteIpListRequest) (response *CreateDDoSBlackWhiteIpListResponse, err error) {
    if request == nil {
        request = NewCreateDDoSBlackWhiteIpListRequest()
    }
    request.SetContext(ctx)
    
    response = NewCreateDDoSBlackWhiteIpListResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDDoSConnectLimitRequest() (request *CreateDDoSConnectLimitRequest) {
    request = &CreateDDoSConnectLimitRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("antiddos", APIVersion, "CreateDDoSConnectLimit")
    
    
    return
}

func NewCreateDDoSConnectLimitResponse() (response *CreateDDoSConnectLimitResponse) {
    response = &CreateDDoSConnectLimitResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateDDoSConnectLimit
// 配置DDoS连接抑制选项
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
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateDDoSConnectLimit(request *CreateDDoSConnectLimitRequest) (response *CreateDDoSConnectLimitResponse, err error) {
    if request == nil {
        request = NewCreateDDoSConnectLimitRequest()
    }
    
    response = NewCreateDDoSConnectLimitResponse()
    err = c.Send(request, response)
    return
}

// CreateDDoSConnectLimit
// 配置DDoS连接抑制选项
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
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateDDoSConnectLimitWithContext(ctx context.Context, request *CreateDDoSConnectLimitRequest) (response *CreateDDoSConnectLimitResponse, err error) {
    if request == nil {
        request = NewCreateDDoSConnectLimitRequest()
    }
    request.SetContext(ctx)
    
    response = NewCreateDDoSConnectLimitResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDDoSGeoIPBlockConfigRequest() (request *CreateDDoSGeoIPBlockConfigRequest) {
    request = &CreateDDoSGeoIPBlockConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("antiddos", APIVersion, "CreateDDoSGeoIPBlockConfig")
    
    
    return
}

func NewCreateDDoSGeoIPBlockConfigResponse() (response *CreateDDoSGeoIPBlockConfigResponse) {
    response = &CreateDDoSGeoIPBlockConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateDDoSGeoIPBlockConfig
// 添加DDoS防护的区域封禁配置
//
// 可能返回的错误码:
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateDDoSGeoIPBlockConfig(request *CreateDDoSGeoIPBlockConfigRequest) (response *CreateDDoSGeoIPBlockConfigResponse, err error) {
    if request == nil {
        request = NewCreateDDoSGeoIPBlockConfigRequest()
    }
    
    response = NewCreateDDoSGeoIPBlockConfigResponse()
    err = c.Send(request, response)
    return
}

// CreateDDoSGeoIPBlockConfig
// 添加DDoS防护的区域封禁配置
//
// 可能返回的错误码:
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateDDoSGeoIPBlockConfigWithContext(ctx context.Context, request *CreateDDoSGeoIPBlockConfigRequest) (response *CreateDDoSGeoIPBlockConfigResponse, err error) {
    if request == nil {
        request = NewCreateDDoSGeoIPBlockConfigRequest()
    }
    request.SetContext(ctx)
    
    response = NewCreateDDoSGeoIPBlockConfigResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDDoSSpeedLimitConfigRequest() (request *CreateDDoSSpeedLimitConfigRequest) {
    request = &CreateDDoSSpeedLimitConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("antiddos", APIVersion, "CreateDDoSSpeedLimitConfig")
    
    
    return
}

func NewCreateDDoSSpeedLimitConfigResponse() (response *CreateDDoSSpeedLimitConfigResponse) {
    response = &CreateDDoSSpeedLimitConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateDDoSSpeedLimitConfig
// 添加DDoS防护的访问限速配置
//
// 可能返回的错误码:
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateDDoSSpeedLimitConfig(request *CreateDDoSSpeedLimitConfigRequest) (response *CreateDDoSSpeedLimitConfigResponse, err error) {
    if request == nil {
        request = NewCreateDDoSSpeedLimitConfigRequest()
    }
    
    response = NewCreateDDoSSpeedLimitConfigResponse()
    err = c.Send(request, response)
    return
}

// CreateDDoSSpeedLimitConfig
// 添加DDoS防护的访问限速配置
//
// 可能返回的错误码:
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateDDoSSpeedLimitConfigWithContext(ctx context.Context, request *CreateDDoSSpeedLimitConfigRequest) (response *CreateDDoSSpeedLimitConfigResponse, err error) {
    if request == nil {
        request = NewCreateDDoSSpeedLimitConfigRequest()
    }
    request.SetContext(ctx)
    
    response = NewCreateDDoSSpeedLimitConfigResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDefaultAlarmThresholdRequest() (request *CreateDefaultAlarmThresholdRequest) {
    request = &CreateDefaultAlarmThresholdRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("antiddos", APIVersion, "CreateDefaultAlarmThreshold")
    
    
    return
}

func NewCreateDefaultAlarmThresholdResponse() (response *CreateDefaultAlarmThresholdResponse) {
    response = &CreateDefaultAlarmThresholdResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateDefaultAlarmThreshold
// 设置单IP默认告警阈值配置
//
// 可能返回的错误码:
//  RESOURCEINUSE = "ResourceInUse"
func (c *Client) CreateDefaultAlarmThreshold(request *CreateDefaultAlarmThresholdRequest) (response *CreateDefaultAlarmThresholdResponse, err error) {
    if request == nil {
        request = NewCreateDefaultAlarmThresholdRequest()
    }
    
    response = NewCreateDefaultAlarmThresholdResponse()
    err = c.Send(request, response)
    return
}

// CreateDefaultAlarmThreshold
// 设置单IP默认告警阈值配置
//
// 可能返回的错误码:
//  RESOURCEINUSE = "ResourceInUse"
func (c *Client) CreateDefaultAlarmThresholdWithContext(ctx context.Context, request *CreateDefaultAlarmThresholdRequest) (response *CreateDefaultAlarmThresholdResponse, err error) {
    if request == nil {
        request = NewCreateDefaultAlarmThresholdRequest()
    }
    request.SetContext(ctx)
    
    response = NewCreateDefaultAlarmThresholdResponse()
    err = c.Send(request, response)
    return
}

func NewCreateIPAlarmThresholdConfigRequest() (request *CreateIPAlarmThresholdConfigRequest) {
    request = &CreateIPAlarmThresholdConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("antiddos", APIVersion, "CreateIPAlarmThresholdConfig")
    
    
    return
}

func NewCreateIPAlarmThresholdConfigResponse() (response *CreateIPAlarmThresholdConfigResponse) {
    response = &CreateIPAlarmThresholdConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateIPAlarmThresholdConfig
// 设置单IP告警阈值配置
//
// 可能返回的错误码:
//  RESOURCEINUSE = "ResourceInUse"
func (c *Client) CreateIPAlarmThresholdConfig(request *CreateIPAlarmThresholdConfigRequest) (response *CreateIPAlarmThresholdConfigResponse, err error) {
    if request == nil {
        request = NewCreateIPAlarmThresholdConfigRequest()
    }
    
    response = NewCreateIPAlarmThresholdConfigResponse()
    err = c.Send(request, response)
    return
}

// CreateIPAlarmThresholdConfig
// 设置单IP告警阈值配置
//
// 可能返回的错误码:
//  RESOURCEINUSE = "ResourceInUse"
func (c *Client) CreateIPAlarmThresholdConfigWithContext(ctx context.Context, request *CreateIPAlarmThresholdConfigRequest) (response *CreateIPAlarmThresholdConfigResponse, err error) {
    if request == nil {
        request = NewCreateIPAlarmThresholdConfigRequest()
    }
    request.SetContext(ctx)
    
    response = NewCreateIPAlarmThresholdConfigResponse()
    err = c.Send(request, response)
    return
}

func NewCreateL7RuleCertsRequest() (request *CreateL7RuleCertsRequest) {
    request = &CreateL7RuleCertsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("antiddos", APIVersion, "CreateL7RuleCerts")
    
    
    return
}

func NewCreateL7RuleCertsResponse() (response *CreateL7RuleCertsResponse) {
    response = &CreateL7RuleCertsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateL7RuleCerts
// 批量配置L7转发规则的证书供SSL测调用
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) CreateL7RuleCerts(request *CreateL7RuleCertsRequest) (response *CreateL7RuleCertsResponse, err error) {
    if request == nil {
        request = NewCreateL7RuleCertsRequest()
    }
    
    response = NewCreateL7RuleCertsResponse()
    err = c.Send(request, response)
    return
}

// CreateL7RuleCerts
// 批量配置L7转发规则的证书供SSL测调用
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) CreateL7RuleCertsWithContext(ctx context.Context, request *CreateL7RuleCertsRequest) (response *CreateL7RuleCertsResponse, err error) {
    if request == nil {
        request = NewCreateL7RuleCertsRequest()
    }
    request.SetContext(ctx)
    
    response = NewCreateL7RuleCertsResponse()
    err = c.Send(request, response)
    return
}

func NewCreatePacketFilterConfigRequest() (request *CreatePacketFilterConfigRequest) {
    request = &CreatePacketFilterConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("antiddos", APIVersion, "CreatePacketFilterConfig")
    
    
    return
}

func NewCreatePacketFilterConfigResponse() (response *CreatePacketFilterConfigResponse) {
    response = &CreatePacketFilterConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreatePacketFilterConfig
// 添加DDoS防护的特征过滤规则
//
// 可能返回的错误码:
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreatePacketFilterConfig(request *CreatePacketFilterConfigRequest) (response *CreatePacketFilterConfigResponse, err error) {
    if request == nil {
        request = NewCreatePacketFilterConfigRequest()
    }
    
    response = NewCreatePacketFilterConfigResponse()
    err = c.Send(request, response)
    return
}

// CreatePacketFilterConfig
// 添加DDoS防护的特征过滤规则
//
// 可能返回的错误码:
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreatePacketFilterConfigWithContext(ctx context.Context, request *CreatePacketFilterConfigRequest) (response *CreatePacketFilterConfigResponse, err error) {
    if request == nil {
        request = NewCreatePacketFilterConfigRequest()
    }
    request.SetContext(ctx)
    
    response = NewCreatePacketFilterConfigResponse()
    err = c.Send(request, response)
    return
}

func NewCreatePortAclConfigRequest() (request *CreatePortAclConfigRequest) {
    request = &CreatePortAclConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("antiddos", APIVersion, "CreatePortAclConfig")
    
    
    return
}

func NewCreatePortAclConfigResponse() (response *CreatePortAclConfigResponse) {
    response = &CreatePortAclConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreatePortAclConfig
// 添加DDoS防护的端口acl策略
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
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreatePortAclConfig(request *CreatePortAclConfigRequest) (response *CreatePortAclConfigResponse, err error) {
    if request == nil {
        request = NewCreatePortAclConfigRequest()
    }
    
    response = NewCreatePortAclConfigResponse()
    err = c.Send(request, response)
    return
}

// CreatePortAclConfig
// 添加DDoS防护的端口acl策略
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
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreatePortAclConfigWithContext(ctx context.Context, request *CreatePortAclConfigRequest) (response *CreatePortAclConfigResponse, err error) {
    if request == nil {
        request = NewCreatePortAclConfigRequest()
    }
    request.SetContext(ctx)
    
    response = NewCreatePortAclConfigResponse()
    err = c.Send(request, response)
    return
}

func NewCreatePortAclConfigListRequest() (request *CreatePortAclConfigListRequest) {
    request = &CreatePortAclConfigListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("antiddos", APIVersion, "CreatePortAclConfigList")
    
    
    return
}

func NewCreatePortAclConfigListResponse() (response *CreatePortAclConfigListResponse) {
    response = &CreatePortAclConfigListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreatePortAclConfigList
// 批量添加DDoS防护的端口acl策略
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
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreatePortAclConfigList(request *CreatePortAclConfigListRequest) (response *CreatePortAclConfigListResponse, err error) {
    if request == nil {
        request = NewCreatePortAclConfigListRequest()
    }
    
    response = NewCreatePortAclConfigListResponse()
    err = c.Send(request, response)
    return
}

// CreatePortAclConfigList
// 批量添加DDoS防护的端口acl策略
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
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreatePortAclConfigListWithContext(ctx context.Context, request *CreatePortAclConfigListRequest) (response *CreatePortAclConfigListResponse, err error) {
    if request == nil {
        request = NewCreatePortAclConfigListRequest()
    }
    request.SetContext(ctx)
    
    response = NewCreatePortAclConfigListResponse()
    err = c.Send(request, response)
    return
}

func NewCreateProtocolBlockConfigRequest() (request *CreateProtocolBlockConfigRequest) {
    request = &CreateProtocolBlockConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("antiddos", APIVersion, "CreateProtocolBlockConfig")
    
    
    return
}

func NewCreateProtocolBlockConfigResponse() (response *CreateProtocolBlockConfigResponse) {
    response = &CreateProtocolBlockConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateProtocolBlockConfig
// 设置DDoS防护的协议封禁配置
//
// 可能返回的错误码:
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateProtocolBlockConfig(request *CreateProtocolBlockConfigRequest) (response *CreateProtocolBlockConfigResponse, err error) {
    if request == nil {
        request = NewCreateProtocolBlockConfigRequest()
    }
    
    response = NewCreateProtocolBlockConfigResponse()
    err = c.Send(request, response)
    return
}

// CreateProtocolBlockConfig
// 设置DDoS防护的协议封禁配置
//
// 可能返回的错误码:
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateProtocolBlockConfigWithContext(ctx context.Context, request *CreateProtocolBlockConfigRequest) (response *CreateProtocolBlockConfigResponse, err error) {
    if request == nil {
        request = NewCreateProtocolBlockConfigRequest()
    }
    request.SetContext(ctx)
    
    response = NewCreateProtocolBlockConfigResponse()
    err = c.Send(request, response)
    return
}

func NewCreateSchedulingDomainRequest() (request *CreateSchedulingDomainRequest) {
    request = &CreateSchedulingDomainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("antiddos", APIVersion, "CreateSchedulingDomain")
    
    
    return
}

func NewCreateSchedulingDomainResponse() (response *CreateSchedulingDomainResponse) {
    response = &CreateSchedulingDomainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateSchedulingDomain
// 创建一个域名，可用于在封堵时调度切换IP
//
// 可能返回的错误码:
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) CreateSchedulingDomain(request *CreateSchedulingDomainRequest) (response *CreateSchedulingDomainResponse, err error) {
    if request == nil {
        request = NewCreateSchedulingDomainRequest()
    }
    
    response = NewCreateSchedulingDomainResponse()
    err = c.Send(request, response)
    return
}

// CreateSchedulingDomain
// 创建一个域名，可用于在封堵时调度切换IP
//
// 可能返回的错误码:
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) CreateSchedulingDomainWithContext(ctx context.Context, request *CreateSchedulingDomainRequest) (response *CreateSchedulingDomainResponse, err error) {
    if request == nil {
        request = NewCreateSchedulingDomainRequest()
    }
    request.SetContext(ctx)
    
    response = NewCreateSchedulingDomainResponse()
    err = c.Send(request, response)
    return
}

func NewCreateWaterPrintConfigRequest() (request *CreateWaterPrintConfigRequest) {
    request = &CreateWaterPrintConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("antiddos", APIVersion, "CreateWaterPrintConfig")
    
    
    return
}

func NewCreateWaterPrintConfigResponse() (response *CreateWaterPrintConfigResponse) {
    response = &CreateWaterPrintConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateWaterPrintConfig
// 添加DDoS防护的水印防护配置
//
// 可能返回的错误码:
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateWaterPrintConfig(request *CreateWaterPrintConfigRequest) (response *CreateWaterPrintConfigResponse, err error) {
    if request == nil {
        request = NewCreateWaterPrintConfigRequest()
    }
    
    response = NewCreateWaterPrintConfigResponse()
    err = c.Send(request, response)
    return
}

// CreateWaterPrintConfig
// 添加DDoS防护的水印防护配置
//
// 可能返回的错误码:
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateWaterPrintConfigWithContext(ctx context.Context, request *CreateWaterPrintConfigRequest) (response *CreateWaterPrintConfigResponse, err error) {
    if request == nil {
        request = NewCreateWaterPrintConfigRequest()
    }
    request.SetContext(ctx)
    
    response = NewCreateWaterPrintConfigResponse()
    err = c.Send(request, response)
    return
}

func NewCreateWaterPrintKeyRequest() (request *CreateWaterPrintKeyRequest) {
    request = &CreateWaterPrintKeyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("antiddos", APIVersion, "CreateWaterPrintKey")
    
    
    return
}

func NewCreateWaterPrintKeyResponse() (response *CreateWaterPrintKeyResponse) {
    response = &CreateWaterPrintKeyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateWaterPrintKey
// 添加DDoS防护的水印防护密钥
//
// 可能返回的错误码:
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateWaterPrintKey(request *CreateWaterPrintKeyRequest) (response *CreateWaterPrintKeyResponse, err error) {
    if request == nil {
        request = NewCreateWaterPrintKeyRequest()
    }
    
    response = NewCreateWaterPrintKeyResponse()
    err = c.Send(request, response)
    return
}

// CreateWaterPrintKey
// 添加DDoS防护的水印防护密钥
//
// 可能返回的错误码:
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateWaterPrintKeyWithContext(ctx context.Context, request *CreateWaterPrintKeyRequest) (response *CreateWaterPrintKeyResponse, err error) {
    if request == nil {
        request = NewCreateWaterPrintKeyRequest()
    }
    request.SetContext(ctx)
    
    response = NewCreateWaterPrintKeyResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteBlackWhiteIpListRequest() (request *DeleteBlackWhiteIpListRequest) {
    request = &DeleteBlackWhiteIpListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("antiddos", APIVersion, "DeleteBlackWhiteIpList")
    
    
    return
}

func NewDeleteBlackWhiteIpListResponse() (response *DeleteBlackWhiteIpListResponse) {
    response = &DeleteBlackWhiteIpListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteBlackWhiteIpList
// 删除DDoS防护的IP黑白名单
//
// 可能返回的错误码:
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteBlackWhiteIpList(request *DeleteBlackWhiteIpListRequest) (response *DeleteBlackWhiteIpListResponse, err error) {
    if request == nil {
        request = NewDeleteBlackWhiteIpListRequest()
    }
    
    response = NewDeleteBlackWhiteIpListResponse()
    err = c.Send(request, response)
    return
}

// DeleteBlackWhiteIpList
// 删除DDoS防护的IP黑白名单
//
// 可能返回的错误码:
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteBlackWhiteIpListWithContext(ctx context.Context, request *DeleteBlackWhiteIpListRequest) (response *DeleteBlackWhiteIpListResponse, err error) {
    if request == nil {
        request = NewDeleteBlackWhiteIpListRequest()
    }
    request.SetContext(ctx)
    
    response = NewDeleteBlackWhiteIpListResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteCCPrecisionPolicyRequest() (request *DeleteCCPrecisionPolicyRequest) {
    request = &DeleteCCPrecisionPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("antiddos", APIVersion, "DeleteCCPrecisionPolicy")
    
    
    return
}

func NewDeleteCCPrecisionPolicyResponse() (response *DeleteCCPrecisionPolicyResponse) {
    response = &DeleteCCPrecisionPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteCCPrecisionPolicy
// 删除CC精准防护策略
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
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteCCPrecisionPolicy(request *DeleteCCPrecisionPolicyRequest) (response *DeleteCCPrecisionPolicyResponse, err error) {
    if request == nil {
        request = NewDeleteCCPrecisionPolicyRequest()
    }
    
    response = NewDeleteCCPrecisionPolicyResponse()
    err = c.Send(request, response)
    return
}

// DeleteCCPrecisionPolicy
// 删除CC精准防护策略
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
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteCCPrecisionPolicyWithContext(ctx context.Context, request *DeleteCCPrecisionPolicyRequest) (response *DeleteCCPrecisionPolicyResponse, err error) {
    if request == nil {
        request = NewDeleteCCPrecisionPolicyRequest()
    }
    request.SetContext(ctx)
    
    response = NewDeleteCCPrecisionPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteCCRequestLimitPolicyRequest() (request *DeleteCCRequestLimitPolicyRequest) {
    request = &DeleteCCRequestLimitPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("antiddos", APIVersion, "DeleteCCRequestLimitPolicy")
    
    
    return
}

func NewDeleteCCRequestLimitPolicyResponse() (response *DeleteCCRequestLimitPolicyResponse) {
    response = &DeleteCCRequestLimitPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteCCRequestLimitPolicy
// 删除CC频率限制策略
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
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteCCRequestLimitPolicy(request *DeleteCCRequestLimitPolicyRequest) (response *DeleteCCRequestLimitPolicyResponse, err error) {
    if request == nil {
        request = NewDeleteCCRequestLimitPolicyRequest()
    }
    
    response = NewDeleteCCRequestLimitPolicyResponse()
    err = c.Send(request, response)
    return
}

// DeleteCCRequestLimitPolicy
// 删除CC频率限制策略
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
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteCCRequestLimitPolicyWithContext(ctx context.Context, request *DeleteCCRequestLimitPolicyRequest) (response *DeleteCCRequestLimitPolicyResponse, err error) {
    if request == nil {
        request = NewDeleteCCRequestLimitPolicyRequest()
    }
    request.SetContext(ctx)
    
    response = NewDeleteCCRequestLimitPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteCcBlackWhiteIpListRequest() (request *DeleteCcBlackWhiteIpListRequest) {
    request = &DeleteCcBlackWhiteIpListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("antiddos", APIVersion, "DeleteCcBlackWhiteIpList")
    
    
    return
}

func NewDeleteCcBlackWhiteIpListResponse() (response *DeleteCcBlackWhiteIpListResponse) {
    response = &DeleteCcBlackWhiteIpListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteCcBlackWhiteIpList
// 删除CC四层黑白名单
//
// 可能返回的错误码:
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteCcBlackWhiteIpList(request *DeleteCcBlackWhiteIpListRequest) (response *DeleteCcBlackWhiteIpListResponse, err error) {
    if request == nil {
        request = NewDeleteCcBlackWhiteIpListRequest()
    }
    
    response = NewDeleteCcBlackWhiteIpListResponse()
    err = c.Send(request, response)
    return
}

// DeleteCcBlackWhiteIpList
// 删除CC四层黑白名单
//
// 可能返回的错误码:
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteCcBlackWhiteIpListWithContext(ctx context.Context, request *DeleteCcBlackWhiteIpListRequest) (response *DeleteCcBlackWhiteIpListResponse, err error) {
    if request == nil {
        request = NewDeleteCcBlackWhiteIpListRequest()
    }
    request.SetContext(ctx)
    
    response = NewDeleteCcBlackWhiteIpListResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteCcGeoIPBlockConfigRequest() (request *DeleteCcGeoIPBlockConfigRequest) {
    request = &DeleteCcGeoIPBlockConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("antiddos", APIVersion, "DeleteCcGeoIPBlockConfig")
    
    
    return
}

func NewDeleteCcGeoIPBlockConfigResponse() (response *DeleteCcGeoIPBlockConfigResponse) {
    response = &DeleteCcGeoIPBlockConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteCcGeoIPBlockConfig
// 删除CC防护的区域封禁配置
//
// 可能返回的错误码:
//  LIMITEXCEEDED = "LimitExceeded"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteCcGeoIPBlockConfig(request *DeleteCcGeoIPBlockConfigRequest) (response *DeleteCcGeoIPBlockConfigResponse, err error) {
    if request == nil {
        request = NewDeleteCcGeoIPBlockConfigRequest()
    }
    
    response = NewDeleteCcGeoIPBlockConfigResponse()
    err = c.Send(request, response)
    return
}

// DeleteCcGeoIPBlockConfig
// 删除CC防护的区域封禁配置
//
// 可能返回的错误码:
//  LIMITEXCEEDED = "LimitExceeded"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteCcGeoIPBlockConfigWithContext(ctx context.Context, request *DeleteCcGeoIPBlockConfigRequest) (response *DeleteCcGeoIPBlockConfigResponse, err error) {
    if request == nil {
        request = NewDeleteCcGeoIPBlockConfigRequest()
    }
    request.SetContext(ctx)
    
    response = NewDeleteCcGeoIPBlockConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteDDoSBlackWhiteIpListRequest() (request *DeleteDDoSBlackWhiteIpListRequest) {
    request = &DeleteDDoSBlackWhiteIpListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("antiddos", APIVersion, "DeleteDDoSBlackWhiteIpList")
    
    
    return
}

func NewDeleteDDoSBlackWhiteIpListResponse() (response *DeleteDDoSBlackWhiteIpListResponse) {
    response = &DeleteDDoSBlackWhiteIpListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteDDoSBlackWhiteIpList
// 删除DDoS防护的IP网段黑白名单
//
// 可能返回的错误码:
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteDDoSBlackWhiteIpList(request *DeleteDDoSBlackWhiteIpListRequest) (response *DeleteDDoSBlackWhiteIpListResponse, err error) {
    if request == nil {
        request = NewDeleteDDoSBlackWhiteIpListRequest()
    }
    
    response = NewDeleteDDoSBlackWhiteIpListResponse()
    err = c.Send(request, response)
    return
}

// DeleteDDoSBlackWhiteIpList
// 删除DDoS防护的IP网段黑白名单
//
// 可能返回的错误码:
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteDDoSBlackWhiteIpListWithContext(ctx context.Context, request *DeleteDDoSBlackWhiteIpListRequest) (response *DeleteDDoSBlackWhiteIpListResponse, err error) {
    if request == nil {
        request = NewDeleteDDoSBlackWhiteIpListRequest()
    }
    request.SetContext(ctx)
    
    response = NewDeleteDDoSBlackWhiteIpListResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteDDoSGeoIPBlockConfigRequest() (request *DeleteDDoSGeoIPBlockConfigRequest) {
    request = &DeleteDDoSGeoIPBlockConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("antiddos", APIVersion, "DeleteDDoSGeoIPBlockConfig")
    
    
    return
}

func NewDeleteDDoSGeoIPBlockConfigResponse() (response *DeleteDDoSGeoIPBlockConfigResponse) {
    response = &DeleteDDoSGeoIPBlockConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteDDoSGeoIPBlockConfig
// 删除DDoS防护的区域封禁配置
//
// 可能返回的错误码:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteDDoSGeoIPBlockConfig(request *DeleteDDoSGeoIPBlockConfigRequest) (response *DeleteDDoSGeoIPBlockConfigResponse, err error) {
    if request == nil {
        request = NewDeleteDDoSGeoIPBlockConfigRequest()
    }
    
    response = NewDeleteDDoSGeoIPBlockConfigResponse()
    err = c.Send(request, response)
    return
}

// DeleteDDoSGeoIPBlockConfig
// 删除DDoS防护的区域封禁配置
//
// 可能返回的错误码:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteDDoSGeoIPBlockConfigWithContext(ctx context.Context, request *DeleteDDoSGeoIPBlockConfigRequest) (response *DeleteDDoSGeoIPBlockConfigResponse, err error) {
    if request == nil {
        request = NewDeleteDDoSGeoIPBlockConfigRequest()
    }
    request.SetContext(ctx)
    
    response = NewDeleteDDoSGeoIPBlockConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteDDoSSpeedLimitConfigRequest() (request *DeleteDDoSSpeedLimitConfigRequest) {
    request = &DeleteDDoSSpeedLimitConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("antiddos", APIVersion, "DeleteDDoSSpeedLimitConfig")
    
    
    return
}

func NewDeleteDDoSSpeedLimitConfigResponse() (response *DeleteDDoSSpeedLimitConfigResponse) {
    response = &DeleteDDoSSpeedLimitConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteDDoSSpeedLimitConfig
// 删除DDoS防护的访问限速配置
//
// 可能返回的错误码:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteDDoSSpeedLimitConfig(request *DeleteDDoSSpeedLimitConfigRequest) (response *DeleteDDoSSpeedLimitConfigResponse, err error) {
    if request == nil {
        request = NewDeleteDDoSSpeedLimitConfigRequest()
    }
    
    response = NewDeleteDDoSSpeedLimitConfigResponse()
    err = c.Send(request, response)
    return
}

// DeleteDDoSSpeedLimitConfig
// 删除DDoS防护的访问限速配置
//
// 可能返回的错误码:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteDDoSSpeedLimitConfigWithContext(ctx context.Context, request *DeleteDDoSSpeedLimitConfigRequest) (response *DeleteDDoSSpeedLimitConfigResponse, err error) {
    if request == nil {
        request = NewDeleteDDoSSpeedLimitConfigRequest()
    }
    request.SetContext(ctx)
    
    response = NewDeleteDDoSSpeedLimitConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDeletePacketFilterConfigRequest() (request *DeletePacketFilterConfigRequest) {
    request = &DeletePacketFilterConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("antiddos", APIVersion, "DeletePacketFilterConfig")
    
    
    return
}

func NewDeletePacketFilterConfigResponse() (response *DeletePacketFilterConfigResponse) {
    response = &DeletePacketFilterConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeletePacketFilterConfig
// 删除DDoS防护的特征过滤规则
//
// 可能返回的错误码:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeletePacketFilterConfig(request *DeletePacketFilterConfigRequest) (response *DeletePacketFilterConfigResponse, err error) {
    if request == nil {
        request = NewDeletePacketFilterConfigRequest()
    }
    
    response = NewDeletePacketFilterConfigResponse()
    err = c.Send(request, response)
    return
}

// DeletePacketFilterConfig
// 删除DDoS防护的特征过滤规则
//
// 可能返回的错误码:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeletePacketFilterConfigWithContext(ctx context.Context, request *DeletePacketFilterConfigRequest) (response *DeletePacketFilterConfigResponse, err error) {
    if request == nil {
        request = NewDeletePacketFilterConfigRequest()
    }
    request.SetContext(ctx)
    
    response = NewDeletePacketFilterConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDeletePortAclConfigRequest() (request *DeletePortAclConfigRequest) {
    request = &DeletePortAclConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("antiddos", APIVersion, "DeletePortAclConfig")
    
    
    return
}

func NewDeletePortAclConfigResponse() (response *DeletePortAclConfigResponse) {
    response = &DeletePortAclConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeletePortAclConfig
// 删除DDoS防护的端口acl策略
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
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeletePortAclConfig(request *DeletePortAclConfigRequest) (response *DeletePortAclConfigResponse, err error) {
    if request == nil {
        request = NewDeletePortAclConfigRequest()
    }
    
    response = NewDeletePortAclConfigResponse()
    err = c.Send(request, response)
    return
}

// DeletePortAclConfig
// 删除DDoS防护的端口acl策略
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
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeletePortAclConfigWithContext(ctx context.Context, request *DeletePortAclConfigRequest) (response *DeletePortAclConfigResponse, err error) {
    if request == nil {
        request = NewDeletePortAclConfigRequest()
    }
    request.SetContext(ctx)
    
    response = NewDeletePortAclConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteWaterPrintConfigRequest() (request *DeleteWaterPrintConfigRequest) {
    request = &DeleteWaterPrintConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("antiddos", APIVersion, "DeleteWaterPrintConfig")
    
    
    return
}

func NewDeleteWaterPrintConfigResponse() (response *DeleteWaterPrintConfigResponse) {
    response = &DeleteWaterPrintConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteWaterPrintConfig
// 删除DDoS防护的水印防护配置
//
// 可能返回的错误码:
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteWaterPrintConfig(request *DeleteWaterPrintConfigRequest) (response *DeleteWaterPrintConfigResponse, err error) {
    if request == nil {
        request = NewDeleteWaterPrintConfigRequest()
    }
    
    response = NewDeleteWaterPrintConfigResponse()
    err = c.Send(request, response)
    return
}

// DeleteWaterPrintConfig
// 删除DDoS防护的水印防护配置
//
// 可能返回的错误码:
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteWaterPrintConfigWithContext(ctx context.Context, request *DeleteWaterPrintConfigRequest) (response *DeleteWaterPrintConfigResponse, err error) {
    if request == nil {
        request = NewDeleteWaterPrintConfigRequest()
    }
    request.SetContext(ctx)
    
    response = NewDeleteWaterPrintConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteWaterPrintKeyRequest() (request *DeleteWaterPrintKeyRequest) {
    request = &DeleteWaterPrintKeyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("antiddos", APIVersion, "DeleteWaterPrintKey")
    
    
    return
}

func NewDeleteWaterPrintKeyResponse() (response *DeleteWaterPrintKeyResponse) {
    response = &DeleteWaterPrintKeyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteWaterPrintKey
// 删除DDoS防护的水印防护密钥
//
// 可能返回的错误码:
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteWaterPrintKey(request *DeleteWaterPrintKeyRequest) (response *DeleteWaterPrintKeyResponse, err error) {
    if request == nil {
        request = NewDeleteWaterPrintKeyRequest()
    }
    
    response = NewDeleteWaterPrintKeyResponse()
    err = c.Send(request, response)
    return
}

// DeleteWaterPrintKey
// 删除DDoS防护的水印防护密钥
//
// 可能返回的错误码:
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteWaterPrintKeyWithContext(ctx context.Context, request *DeleteWaterPrintKeyRequest) (response *DeleteWaterPrintKeyResponse, err error) {
    if request == nil {
        request = NewDeleteWaterPrintKeyRequest()
    }
    request.SetContext(ctx)
    
    response = NewDeleteWaterPrintKeyResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBasicDeviceStatusRequest() (request *DescribeBasicDeviceStatusRequest) {
    request = &DescribeBasicDeviceStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("antiddos", APIVersion, "DescribeBasicDeviceStatus")
    
    
    return
}

func NewDescribeBasicDeviceStatusResponse() (response *DescribeBasicDeviceStatusResponse) {
    response = &DescribeBasicDeviceStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeBasicDeviceStatus
// 获取基础防护攻击状态
//
// 可能返回的错误码:
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeBasicDeviceStatus(request *DescribeBasicDeviceStatusRequest) (response *DescribeBasicDeviceStatusResponse, err error) {
    if request == nil {
        request = NewDescribeBasicDeviceStatusRequest()
    }
    
    response = NewDescribeBasicDeviceStatusResponse()
    err = c.Send(request, response)
    return
}

// DescribeBasicDeviceStatus
// 获取基础防护攻击状态
//
// 可能返回的错误码:
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeBasicDeviceStatusWithContext(ctx context.Context, request *DescribeBasicDeviceStatusRequest) (response *DescribeBasicDeviceStatusResponse, err error) {
    if request == nil {
        request = NewDescribeBasicDeviceStatusRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeBasicDeviceStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBizTrendRequest() (request *DescribeBizTrendRequest) {
    request = &DescribeBizTrendRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("antiddos", APIVersion, "DescribeBizTrend")
    
    
    return
}

func NewDescribeBizTrendResponse() (response *DescribeBizTrendResponse) {
    response = &DescribeBizTrendResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeBizTrend
// 获取业务流量曲线
//
// 可能返回的错误码:
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeBizTrend(request *DescribeBizTrendRequest) (response *DescribeBizTrendResponse, err error) {
    if request == nil {
        request = NewDescribeBizTrendRequest()
    }
    
    response = NewDescribeBizTrendResponse()
    err = c.Send(request, response)
    return
}

// DescribeBizTrend
// 获取业务流量曲线
//
// 可能返回的错误码:
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeBizTrendWithContext(ctx context.Context, request *DescribeBizTrendRequest) (response *DescribeBizTrendResponse, err error) {
    if request == nil {
        request = NewDescribeBizTrendRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeBizTrendResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBlackWhiteIpListRequest() (request *DescribeBlackWhiteIpListRequest) {
    request = &DescribeBlackWhiteIpListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("antiddos", APIVersion, "DescribeBlackWhiteIpList")
    
    
    return
}

func NewDescribeBlackWhiteIpListResponse() (response *DescribeBlackWhiteIpListResponse) {
    response = &DescribeBlackWhiteIpListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeBlackWhiteIpList
// 获取DDoS防护的IP黑白名单
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
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeBlackWhiteIpList(request *DescribeBlackWhiteIpListRequest) (response *DescribeBlackWhiteIpListResponse, err error) {
    if request == nil {
        request = NewDescribeBlackWhiteIpListRequest()
    }
    
    response = NewDescribeBlackWhiteIpListResponse()
    err = c.Send(request, response)
    return
}

// DescribeBlackWhiteIpList
// 获取DDoS防护的IP黑白名单
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
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeBlackWhiteIpListWithContext(ctx context.Context, request *DescribeBlackWhiteIpListRequest) (response *DescribeBlackWhiteIpListResponse, err error) {
    if request == nil {
        request = NewDescribeBlackWhiteIpListRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeBlackWhiteIpListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCCPrecisionPlyListRequest() (request *DescribeCCPrecisionPlyListRequest) {
    request = &DescribeCCPrecisionPlyListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("antiddos", APIVersion, "DescribeCCPrecisionPlyList")
    
    
    return
}

func NewDescribeCCPrecisionPlyListResponse() (response *DescribeCCPrecisionPlyListResponse) {
    response = &DescribeCCPrecisionPlyListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCCPrecisionPlyList
// 获取CC精准防护列表
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
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeCCPrecisionPlyList(request *DescribeCCPrecisionPlyListRequest) (response *DescribeCCPrecisionPlyListResponse, err error) {
    if request == nil {
        request = NewDescribeCCPrecisionPlyListRequest()
    }
    
    response = NewDescribeCCPrecisionPlyListResponse()
    err = c.Send(request, response)
    return
}

// DescribeCCPrecisionPlyList
// 获取CC精准防护列表
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
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeCCPrecisionPlyListWithContext(ctx context.Context, request *DescribeCCPrecisionPlyListRequest) (response *DescribeCCPrecisionPlyListResponse, err error) {
    if request == nil {
        request = NewDescribeCCPrecisionPlyListRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeCCPrecisionPlyListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCCReqLimitPolicyListRequest() (request *DescribeCCReqLimitPolicyListRequest) {
    request = &DescribeCCReqLimitPolicyListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("antiddos", APIVersion, "DescribeCCReqLimitPolicyList")
    
    
    return
}

func NewDescribeCCReqLimitPolicyListResponse() (response *DescribeCCReqLimitPolicyListResponse) {
    response = &DescribeCCReqLimitPolicyListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCCReqLimitPolicyList
// 获取CC频率限制策略列表
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
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeCCReqLimitPolicyList(request *DescribeCCReqLimitPolicyListRequest) (response *DescribeCCReqLimitPolicyListResponse, err error) {
    if request == nil {
        request = NewDescribeCCReqLimitPolicyListRequest()
    }
    
    response = NewDescribeCCReqLimitPolicyListResponse()
    err = c.Send(request, response)
    return
}

// DescribeCCReqLimitPolicyList
// 获取CC频率限制策略列表
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
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeCCReqLimitPolicyListWithContext(ctx context.Context, request *DescribeCCReqLimitPolicyListRequest) (response *DescribeCCReqLimitPolicyListResponse, err error) {
    if request == nil {
        request = NewDescribeCCReqLimitPolicyListRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeCCReqLimitPolicyListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCCTrendRequest() (request *DescribeCCTrendRequest) {
    request = &DescribeCCTrendRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("antiddos", APIVersion, "DescribeCCTrend")
    
    
    return
}

func NewDescribeCCTrendResponse() (response *DescribeCCTrendResponse) {
    response = &DescribeCCTrendResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCCTrend
// 获取CC攻击指标数据，包括总请求峰值(QPS)和攻击请求(QPS)以及总请求次数和攻击请求次数
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
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeCCTrend(request *DescribeCCTrendRequest) (response *DescribeCCTrendResponse, err error) {
    if request == nil {
        request = NewDescribeCCTrendRequest()
    }
    
    response = NewDescribeCCTrendResponse()
    err = c.Send(request, response)
    return
}

// DescribeCCTrend
// 获取CC攻击指标数据，包括总请求峰值(QPS)和攻击请求(QPS)以及总请求次数和攻击请求次数
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
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeCCTrendWithContext(ctx context.Context, request *DescribeCCTrendRequest) (response *DescribeCCTrendResponse, err error) {
    if request == nil {
        request = NewDescribeCCTrendRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeCCTrendResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCcBlackWhiteIpListRequest() (request *DescribeCcBlackWhiteIpListRequest) {
    request = &DescribeCcBlackWhiteIpListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("antiddos", APIVersion, "DescribeCcBlackWhiteIpList")
    
    
    return
}

func NewDescribeCcBlackWhiteIpListResponse() (response *DescribeCcBlackWhiteIpListResponse) {
    response = &DescribeCcBlackWhiteIpListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCcBlackWhiteIpList
// 获取CC四层黑白名单列表
//
// 可能返回的错误码:
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeCcBlackWhiteIpList(request *DescribeCcBlackWhiteIpListRequest) (response *DescribeCcBlackWhiteIpListResponse, err error) {
    if request == nil {
        request = NewDescribeCcBlackWhiteIpListRequest()
    }
    
    response = NewDescribeCcBlackWhiteIpListResponse()
    err = c.Send(request, response)
    return
}

// DescribeCcBlackWhiteIpList
// 获取CC四层黑白名单列表
//
// 可能返回的错误码:
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeCcBlackWhiteIpListWithContext(ctx context.Context, request *DescribeCcBlackWhiteIpListRequest) (response *DescribeCcBlackWhiteIpListResponse, err error) {
    if request == nil {
        request = NewDescribeCcBlackWhiteIpListRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeCcBlackWhiteIpListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCcGeoIPBlockConfigListRequest() (request *DescribeCcGeoIPBlockConfigListRequest) {
    request = &DescribeCcGeoIPBlockConfigListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("antiddos", APIVersion, "DescribeCcGeoIPBlockConfigList")
    
    
    return
}

func NewDescribeCcGeoIPBlockConfigListResponse() (response *DescribeCcGeoIPBlockConfigListResponse) {
    response = &DescribeCcGeoIPBlockConfigListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCcGeoIPBlockConfigList
// 获取CC防护的区域封禁配置列表
//
// 可能返回的错误码:
//  LIMITEXCEEDED = "LimitExceeded"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeCcGeoIPBlockConfigList(request *DescribeCcGeoIPBlockConfigListRequest) (response *DescribeCcGeoIPBlockConfigListResponse, err error) {
    if request == nil {
        request = NewDescribeCcGeoIPBlockConfigListRequest()
    }
    
    response = NewDescribeCcGeoIPBlockConfigListResponse()
    err = c.Send(request, response)
    return
}

// DescribeCcGeoIPBlockConfigList
// 获取CC防护的区域封禁配置列表
//
// 可能返回的错误码:
//  LIMITEXCEEDED = "LimitExceeded"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeCcGeoIPBlockConfigListWithContext(ctx context.Context, request *DescribeCcGeoIPBlockConfigListRequest) (response *DescribeCcGeoIPBlockConfigListResponse, err error) {
    if request == nil {
        request = NewDescribeCcGeoIPBlockConfigListRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeCcGeoIPBlockConfigListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDDoSBlackWhiteIpListRequest() (request *DescribeDDoSBlackWhiteIpListRequest) {
    request = &DescribeDDoSBlackWhiteIpListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("antiddos", APIVersion, "DescribeDDoSBlackWhiteIpList")
    
    
    return
}

func NewDescribeDDoSBlackWhiteIpListResponse() (response *DescribeDDoSBlackWhiteIpListResponse) {
    response = &DescribeDDoSBlackWhiteIpListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDDoSBlackWhiteIpList
// 获取DDoS防护的IP网段黑白名单
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
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDDoSBlackWhiteIpList(request *DescribeDDoSBlackWhiteIpListRequest) (response *DescribeDDoSBlackWhiteIpListResponse, err error) {
    if request == nil {
        request = NewDescribeDDoSBlackWhiteIpListRequest()
    }
    
    response = NewDescribeDDoSBlackWhiteIpListResponse()
    err = c.Send(request, response)
    return
}

// DescribeDDoSBlackWhiteIpList
// 获取DDoS防护的IP网段黑白名单
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
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDDoSBlackWhiteIpListWithContext(ctx context.Context, request *DescribeDDoSBlackWhiteIpListRequest) (response *DescribeDDoSBlackWhiteIpListResponse, err error) {
    if request == nil {
        request = NewDescribeDDoSBlackWhiteIpListRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeDDoSBlackWhiteIpListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDDoSConnectLimitListRequest() (request *DescribeDDoSConnectLimitListRequest) {
    request = &DescribeDDoSConnectLimitListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("antiddos", APIVersion, "DescribeDDoSConnectLimitList")
    
    
    return
}

func NewDescribeDDoSConnectLimitListResponse() (response *DescribeDDoSConnectLimitListResponse) {
    response = &DescribeDDoSConnectLimitListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDDoSConnectLimitList
// 获取DDoS连接抑制配置列表
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
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDDoSConnectLimitList(request *DescribeDDoSConnectLimitListRequest) (response *DescribeDDoSConnectLimitListResponse, err error) {
    if request == nil {
        request = NewDescribeDDoSConnectLimitListRequest()
    }
    
    response = NewDescribeDDoSConnectLimitListResponse()
    err = c.Send(request, response)
    return
}

// DescribeDDoSConnectLimitList
// 获取DDoS连接抑制配置列表
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
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDDoSConnectLimitListWithContext(ctx context.Context, request *DescribeDDoSConnectLimitListRequest) (response *DescribeDDoSConnectLimitListResponse, err error) {
    if request == nil {
        request = NewDescribeDDoSConnectLimitListRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeDDoSConnectLimitListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDDoSTrendRequest() (request *DescribeDDoSTrendRequest) {
    request = &DescribeDDoSTrendRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("antiddos", APIVersion, "DescribeDDoSTrend")
    
    
    return
}

func NewDescribeDDoSTrendResponse() (response *DescribeDDoSTrendResponse) {
    response = &DescribeDDoSTrendResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDDoSTrend
// 获取DDoS攻击流量带宽和攻击包速率数据
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
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDDoSTrend(request *DescribeDDoSTrendRequest) (response *DescribeDDoSTrendResponse, err error) {
    if request == nil {
        request = NewDescribeDDoSTrendRequest()
    }
    
    response = NewDescribeDDoSTrendResponse()
    err = c.Send(request, response)
    return
}

// DescribeDDoSTrend
// 获取DDoS攻击流量带宽和攻击包速率数据
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
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDDoSTrendWithContext(ctx context.Context, request *DescribeDDoSTrendRequest) (response *DescribeDDoSTrendResponse, err error) {
    if request == nil {
        request = NewDescribeDDoSTrendRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeDDoSTrendResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDefaultAlarmThresholdRequest() (request *DescribeDefaultAlarmThresholdRequest) {
    request = &DescribeDefaultAlarmThresholdRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("antiddos", APIVersion, "DescribeDefaultAlarmThreshold")
    
    
    return
}

func NewDescribeDefaultAlarmThresholdResponse() (response *DescribeDefaultAlarmThresholdResponse) {
    response = &DescribeDefaultAlarmThresholdResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDefaultAlarmThreshold
// 获取单IP默认告警阈值配置
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
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDefaultAlarmThreshold(request *DescribeDefaultAlarmThresholdRequest) (response *DescribeDefaultAlarmThresholdResponse, err error) {
    if request == nil {
        request = NewDescribeDefaultAlarmThresholdRequest()
    }
    
    response = NewDescribeDefaultAlarmThresholdResponse()
    err = c.Send(request, response)
    return
}

// DescribeDefaultAlarmThreshold
// 获取单IP默认告警阈值配置
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
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDefaultAlarmThresholdWithContext(ctx context.Context, request *DescribeDefaultAlarmThresholdRequest) (response *DescribeDefaultAlarmThresholdResponse, err error) {
    if request == nil {
        request = NewDescribeDefaultAlarmThresholdRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeDefaultAlarmThresholdResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeL7RulesBySSLCertIdRequest() (request *DescribeL7RulesBySSLCertIdRequest) {
    request = &DescribeL7RulesBySSLCertIdRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("antiddos", APIVersion, "DescribeL7RulesBySSLCertId")
    
    
    return
}

func NewDescribeL7RulesBySSLCertIdResponse() (response *DescribeL7RulesBySSLCertIdResponse) {
    response = &DescribeL7RulesBySSLCertIdResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeL7RulesBySSLCertId
// 查询与证书ID对于域名匹配的七层规则
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeL7RulesBySSLCertId(request *DescribeL7RulesBySSLCertIdRequest) (response *DescribeL7RulesBySSLCertIdResponse, err error) {
    if request == nil {
        request = NewDescribeL7RulesBySSLCertIdRequest()
    }
    
    response = NewDescribeL7RulesBySSLCertIdResponse()
    err = c.Send(request, response)
    return
}

// DescribeL7RulesBySSLCertId
// 查询与证书ID对于域名匹配的七层规则
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeL7RulesBySSLCertIdWithContext(ctx context.Context, request *DescribeL7RulesBySSLCertIdRequest) (response *DescribeL7RulesBySSLCertIdResponse, err error) {
    if request == nil {
        request = NewDescribeL7RulesBySSLCertIdRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeL7RulesBySSLCertIdResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeListBGPIPInstancesRequest() (request *DescribeListBGPIPInstancesRequest) {
    request = &DescribeListBGPIPInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("antiddos", APIVersion, "DescribeListBGPIPInstances")
    
    
    return
}

func NewDescribeListBGPIPInstancesResponse() (response *DescribeListBGPIPInstancesResponse) {
    response = &DescribeListBGPIPInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeListBGPIPInstances
// 获取高防IP资产实例列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeListBGPIPInstances(request *DescribeListBGPIPInstancesRequest) (response *DescribeListBGPIPInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeListBGPIPInstancesRequest()
    }
    
    response = NewDescribeListBGPIPInstancesResponse()
    err = c.Send(request, response)
    return
}

// DescribeListBGPIPInstances
// 获取高防IP资产实例列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeListBGPIPInstancesWithContext(ctx context.Context, request *DescribeListBGPIPInstancesRequest) (response *DescribeListBGPIPInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeListBGPIPInstancesRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeListBGPIPInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeListBGPInstancesRequest() (request *DescribeListBGPInstancesRequest) {
    request = &DescribeListBGPInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("antiddos", APIVersion, "DescribeListBGPInstances")
    
    
    return
}

func NewDescribeListBGPInstancesResponse() (response *DescribeListBGPInstancesResponse) {
    response = &DescribeListBGPInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeListBGPInstances
// 获取高防包资产实例列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeListBGPInstances(request *DescribeListBGPInstancesRequest) (response *DescribeListBGPInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeListBGPInstancesRequest()
    }
    
    response = NewDescribeListBGPInstancesResponse()
    err = c.Send(request, response)
    return
}

// DescribeListBGPInstances
// 获取高防包资产实例列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeListBGPInstancesWithContext(ctx context.Context, request *DescribeListBGPInstancesRequest) (response *DescribeListBGPInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeListBGPInstancesRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeListBGPInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeListBlackWhiteIpListRequest() (request *DescribeListBlackWhiteIpListRequest) {
    request = &DescribeListBlackWhiteIpListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("antiddos", APIVersion, "DescribeListBlackWhiteIpList")
    
    
    return
}

func NewDescribeListBlackWhiteIpListResponse() (response *DescribeListBlackWhiteIpListResponse) {
    response = &DescribeListBlackWhiteIpListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeListBlackWhiteIpList
// 获取DDoS防护的IP黑白名单列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeListBlackWhiteIpList(request *DescribeListBlackWhiteIpListRequest) (response *DescribeListBlackWhiteIpListResponse, err error) {
    if request == nil {
        request = NewDescribeListBlackWhiteIpListRequest()
    }
    
    response = NewDescribeListBlackWhiteIpListResponse()
    err = c.Send(request, response)
    return
}

// DescribeListBlackWhiteIpList
// 获取DDoS防护的IP黑白名单列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeListBlackWhiteIpListWithContext(ctx context.Context, request *DescribeListBlackWhiteIpListRequest) (response *DescribeListBlackWhiteIpListResponse, err error) {
    if request == nil {
        request = NewDescribeListBlackWhiteIpListRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeListBlackWhiteIpListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeListDDoSAIRequest() (request *DescribeListDDoSAIRequest) {
    request = &DescribeListDDoSAIRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("antiddos", APIVersion, "DescribeListDDoSAI")
    
    
    return
}

func NewDescribeListDDoSAIResponse() (response *DescribeListDDoSAIResponse) {
    response = &DescribeListDDoSAIResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeListDDoSAI
// 获取DDoS防护的AI防护开关列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeListDDoSAI(request *DescribeListDDoSAIRequest) (response *DescribeListDDoSAIResponse, err error) {
    if request == nil {
        request = NewDescribeListDDoSAIRequest()
    }
    
    response = NewDescribeListDDoSAIResponse()
    err = c.Send(request, response)
    return
}

// DescribeListDDoSAI
// 获取DDoS防护的AI防护开关列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeListDDoSAIWithContext(ctx context.Context, request *DescribeListDDoSAIRequest) (response *DescribeListDDoSAIResponse, err error) {
    if request == nil {
        request = NewDescribeListDDoSAIRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeListDDoSAIResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeListDDoSGeoIPBlockConfigRequest() (request *DescribeListDDoSGeoIPBlockConfigRequest) {
    request = &DescribeListDDoSGeoIPBlockConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("antiddos", APIVersion, "DescribeListDDoSGeoIPBlockConfig")
    
    
    return
}

func NewDescribeListDDoSGeoIPBlockConfigResponse() (response *DescribeListDDoSGeoIPBlockConfigResponse) {
    response = &DescribeListDDoSGeoIPBlockConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeListDDoSGeoIPBlockConfig
// 获取DDoS防护的区域封禁配置列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeListDDoSGeoIPBlockConfig(request *DescribeListDDoSGeoIPBlockConfigRequest) (response *DescribeListDDoSGeoIPBlockConfigResponse, err error) {
    if request == nil {
        request = NewDescribeListDDoSGeoIPBlockConfigRequest()
    }
    
    response = NewDescribeListDDoSGeoIPBlockConfigResponse()
    err = c.Send(request, response)
    return
}

// DescribeListDDoSGeoIPBlockConfig
// 获取DDoS防护的区域封禁配置列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeListDDoSGeoIPBlockConfigWithContext(ctx context.Context, request *DescribeListDDoSGeoIPBlockConfigRequest) (response *DescribeListDDoSGeoIPBlockConfigResponse, err error) {
    if request == nil {
        request = NewDescribeListDDoSGeoIPBlockConfigRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeListDDoSGeoIPBlockConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeListDDoSSpeedLimitConfigRequest() (request *DescribeListDDoSSpeedLimitConfigRequest) {
    request = &DescribeListDDoSSpeedLimitConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("antiddos", APIVersion, "DescribeListDDoSSpeedLimitConfig")
    
    
    return
}

func NewDescribeListDDoSSpeedLimitConfigResponse() (response *DescribeListDDoSSpeedLimitConfigResponse) {
    response = &DescribeListDDoSSpeedLimitConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeListDDoSSpeedLimitConfig
// 获取DDoS防护的访问限速配置列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeListDDoSSpeedLimitConfig(request *DescribeListDDoSSpeedLimitConfigRequest) (response *DescribeListDDoSSpeedLimitConfigResponse, err error) {
    if request == nil {
        request = NewDescribeListDDoSSpeedLimitConfigRequest()
    }
    
    response = NewDescribeListDDoSSpeedLimitConfigResponse()
    err = c.Send(request, response)
    return
}

// DescribeListDDoSSpeedLimitConfig
// 获取DDoS防护的访问限速配置列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeListDDoSSpeedLimitConfigWithContext(ctx context.Context, request *DescribeListDDoSSpeedLimitConfigRequest) (response *DescribeListDDoSSpeedLimitConfigResponse, err error) {
    if request == nil {
        request = NewDescribeListDDoSSpeedLimitConfigRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeListDDoSSpeedLimitConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeListIPAlarmConfigRequest() (request *DescribeListIPAlarmConfigRequest) {
    request = &DescribeListIPAlarmConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("antiddos", APIVersion, "DescribeListIPAlarmConfig")
    
    
    return
}

func NewDescribeListIPAlarmConfigResponse() (response *DescribeListIPAlarmConfigResponse) {
    response = &DescribeListIPAlarmConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeListIPAlarmConfig
// 获取单IP告警阈值配置列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeListIPAlarmConfig(request *DescribeListIPAlarmConfigRequest) (response *DescribeListIPAlarmConfigResponse, err error) {
    if request == nil {
        request = NewDescribeListIPAlarmConfigRequest()
    }
    
    response = NewDescribeListIPAlarmConfigResponse()
    err = c.Send(request, response)
    return
}

// DescribeListIPAlarmConfig
// 获取单IP告警阈值配置列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeListIPAlarmConfigWithContext(ctx context.Context, request *DescribeListIPAlarmConfigRequest) (response *DescribeListIPAlarmConfigResponse, err error) {
    if request == nil {
        request = NewDescribeListIPAlarmConfigRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeListIPAlarmConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeListListenerRequest() (request *DescribeListListenerRequest) {
    request = &DescribeListListenerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("antiddos", APIVersion, "DescribeListListener")
    
    
    return
}

func NewDescribeListListenerResponse() (response *DescribeListListenerResponse) {
    response = &DescribeListListenerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeListListener
// 获取转发监听器列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeListListener(request *DescribeListListenerRequest) (response *DescribeListListenerResponse, err error) {
    if request == nil {
        request = NewDescribeListListenerRequest()
    }
    
    response = NewDescribeListListenerResponse()
    err = c.Send(request, response)
    return
}

// DescribeListListener
// 获取转发监听器列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeListListenerWithContext(ctx context.Context, request *DescribeListListenerRequest) (response *DescribeListListenerResponse, err error) {
    if request == nil {
        request = NewDescribeListListenerRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeListListenerResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeListPacketFilterConfigRequest() (request *DescribeListPacketFilterConfigRequest) {
    request = &DescribeListPacketFilterConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("antiddos", APIVersion, "DescribeListPacketFilterConfig")
    
    
    return
}

func NewDescribeListPacketFilterConfigResponse() (response *DescribeListPacketFilterConfigResponse) {
    response = &DescribeListPacketFilterConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeListPacketFilterConfig
// 获取DDoS防护的特征过滤规则列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeListPacketFilterConfig(request *DescribeListPacketFilterConfigRequest) (response *DescribeListPacketFilterConfigResponse, err error) {
    if request == nil {
        request = NewDescribeListPacketFilterConfigRequest()
    }
    
    response = NewDescribeListPacketFilterConfigResponse()
    err = c.Send(request, response)
    return
}

// DescribeListPacketFilterConfig
// 获取DDoS防护的特征过滤规则列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeListPacketFilterConfigWithContext(ctx context.Context, request *DescribeListPacketFilterConfigRequest) (response *DescribeListPacketFilterConfigResponse, err error) {
    if request == nil {
        request = NewDescribeListPacketFilterConfigRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeListPacketFilterConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeListPortAclListRequest() (request *DescribeListPortAclListRequest) {
    request = &DescribeListPortAclListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("antiddos", APIVersion, "DescribeListPortAclList")
    
    
    return
}

func NewDescribeListPortAclListResponse() (response *DescribeListPortAclListResponse) {
    response = &DescribeListPortAclListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeListPortAclList
// 获取DDoS防护的端口acl策略列表
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
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeListPortAclList(request *DescribeListPortAclListRequest) (response *DescribeListPortAclListResponse, err error) {
    if request == nil {
        request = NewDescribeListPortAclListRequest()
    }
    
    response = NewDescribeListPortAclListResponse()
    err = c.Send(request, response)
    return
}

// DescribeListPortAclList
// 获取DDoS防护的端口acl策略列表
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
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeListPortAclListWithContext(ctx context.Context, request *DescribeListPortAclListRequest) (response *DescribeListPortAclListResponse, err error) {
    if request == nil {
        request = NewDescribeListPortAclListRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeListPortAclListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeListProtectThresholdConfigRequest() (request *DescribeListProtectThresholdConfigRequest) {
    request = &DescribeListProtectThresholdConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("antiddos", APIVersion, "DescribeListProtectThresholdConfig")
    
    
    return
}

func NewDescribeListProtectThresholdConfigResponse() (response *DescribeListProtectThresholdConfigResponse) {
    response = &DescribeListProtectThresholdConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeListProtectThresholdConfig
// 获取防护阈值配置列表，包括DDoS的AI、等级、CC阈值开关等
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
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeListProtectThresholdConfig(request *DescribeListProtectThresholdConfigRequest) (response *DescribeListProtectThresholdConfigResponse, err error) {
    if request == nil {
        request = NewDescribeListProtectThresholdConfigRequest()
    }
    
    response = NewDescribeListProtectThresholdConfigResponse()
    err = c.Send(request, response)
    return
}

// DescribeListProtectThresholdConfig
// 获取防护阈值配置列表，包括DDoS的AI、等级、CC阈值开关等
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
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeListProtectThresholdConfigWithContext(ctx context.Context, request *DescribeListProtectThresholdConfigRequest) (response *DescribeListProtectThresholdConfigResponse, err error) {
    if request == nil {
        request = NewDescribeListProtectThresholdConfigRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeListProtectThresholdConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeListProtocolBlockConfigRequest() (request *DescribeListProtocolBlockConfigRequest) {
    request = &DescribeListProtocolBlockConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("antiddos", APIVersion, "DescribeListProtocolBlockConfig")
    
    
    return
}

func NewDescribeListProtocolBlockConfigResponse() (response *DescribeListProtocolBlockConfigResponse) {
    response = &DescribeListProtocolBlockConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeListProtocolBlockConfig
// 获取DDoS防护的协议封禁配置列表
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
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeListProtocolBlockConfig(request *DescribeListProtocolBlockConfigRequest) (response *DescribeListProtocolBlockConfigResponse, err error) {
    if request == nil {
        request = NewDescribeListProtocolBlockConfigRequest()
    }
    
    response = NewDescribeListProtocolBlockConfigResponse()
    err = c.Send(request, response)
    return
}

// DescribeListProtocolBlockConfig
// 获取DDoS防护的协议封禁配置列表
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
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeListProtocolBlockConfigWithContext(ctx context.Context, request *DescribeListProtocolBlockConfigRequest) (response *DescribeListProtocolBlockConfigResponse, err error) {
    if request == nil {
        request = NewDescribeListProtocolBlockConfigRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeListProtocolBlockConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeListSchedulingDomainRequest() (request *DescribeListSchedulingDomainRequest) {
    request = &DescribeListSchedulingDomainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("antiddos", APIVersion, "DescribeListSchedulingDomain")
    
    
    return
}

func NewDescribeListSchedulingDomainResponse() (response *DescribeListSchedulingDomainResponse) {
    response = &DescribeListSchedulingDomainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeListSchedulingDomain
// 获取智能调度域名列表
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
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeListSchedulingDomain(request *DescribeListSchedulingDomainRequest) (response *DescribeListSchedulingDomainResponse, err error) {
    if request == nil {
        request = NewDescribeListSchedulingDomainRequest()
    }
    
    response = NewDescribeListSchedulingDomainResponse()
    err = c.Send(request, response)
    return
}

// DescribeListSchedulingDomain
// 获取智能调度域名列表
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
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeListSchedulingDomainWithContext(ctx context.Context, request *DescribeListSchedulingDomainRequest) (response *DescribeListSchedulingDomainResponse, err error) {
    if request == nil {
        request = NewDescribeListSchedulingDomainRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeListSchedulingDomainResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeListWaterPrintConfigRequest() (request *DescribeListWaterPrintConfigRequest) {
    request = &DescribeListWaterPrintConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("antiddos", APIVersion, "DescribeListWaterPrintConfig")
    
    
    return
}

func NewDescribeListWaterPrintConfigResponse() (response *DescribeListWaterPrintConfigResponse) {
    response = &DescribeListWaterPrintConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeListWaterPrintConfig
// 获取DDoS防护的水印防护配置列表
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
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeListWaterPrintConfig(request *DescribeListWaterPrintConfigRequest) (response *DescribeListWaterPrintConfigResponse, err error) {
    if request == nil {
        request = NewDescribeListWaterPrintConfigRequest()
    }
    
    response = NewDescribeListWaterPrintConfigResponse()
    err = c.Send(request, response)
    return
}

// DescribeListWaterPrintConfig
// 获取DDoS防护的水印防护配置列表
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
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeListWaterPrintConfigWithContext(ctx context.Context, request *DescribeListWaterPrintConfigRequest) (response *DescribeListWaterPrintConfigResponse, err error) {
    if request == nil {
        request = NewDescribeListWaterPrintConfigRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeListWaterPrintConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDisassociateDDoSEipAddressRequest() (request *DisassociateDDoSEipAddressRequest) {
    request = &DisassociateDDoSEipAddressRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("antiddos", APIVersion, "DisassociateDDoSEipAddress")
    
    
    return
}

func NewDisassociateDDoSEipAddressResponse() (response *DisassociateDDoSEipAddressResponse) {
    response = &DisassociateDDoSEipAddressResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DisassociateDDoSEipAddress
// 本接口 (DisassociateDDoSEipAddress) 用于解绑高防弹性公网IP。
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
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DisassociateDDoSEipAddress(request *DisassociateDDoSEipAddressRequest) (response *DisassociateDDoSEipAddressResponse, err error) {
    if request == nil {
        request = NewDisassociateDDoSEipAddressRequest()
    }
    
    response = NewDisassociateDDoSEipAddressResponse()
    err = c.Send(request, response)
    return
}

// DisassociateDDoSEipAddress
// 本接口 (DisassociateDDoSEipAddress) 用于解绑高防弹性公网IP。
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
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DisassociateDDoSEipAddressWithContext(ctx context.Context, request *DisassociateDDoSEipAddressRequest) (response *DisassociateDDoSEipAddressResponse, err error) {
    if request == nil {
        request = NewDisassociateDDoSEipAddressRequest()
    }
    request.SetContext(ctx)
    
    response = NewDisassociateDDoSEipAddressResponse()
    err = c.Send(request, response)
    return
}

func NewModifyCCLevelPolicyRequest() (request *ModifyCCLevelPolicyRequest) {
    request = &ModifyCCLevelPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("antiddos", APIVersion, "ModifyCCLevelPolicy")
    
    
    return
}

func NewModifyCCLevelPolicyResponse() (response *ModifyCCLevelPolicyResponse) {
    response = &ModifyCCLevelPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyCCLevelPolicy
// 修改CC防护等级
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
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyCCLevelPolicy(request *ModifyCCLevelPolicyRequest) (response *ModifyCCLevelPolicyResponse, err error) {
    if request == nil {
        request = NewModifyCCLevelPolicyRequest()
    }
    
    response = NewModifyCCLevelPolicyResponse()
    err = c.Send(request, response)
    return
}

// ModifyCCLevelPolicy
// 修改CC防护等级
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
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyCCLevelPolicyWithContext(ctx context.Context, request *ModifyCCLevelPolicyRequest) (response *ModifyCCLevelPolicyResponse, err error) {
    if request == nil {
        request = NewModifyCCLevelPolicyRequest()
    }
    request.SetContext(ctx)
    
    response = NewModifyCCLevelPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewModifyCCPrecisionPolicyRequest() (request *ModifyCCPrecisionPolicyRequest) {
    request = &ModifyCCPrecisionPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("antiddos", APIVersion, "ModifyCCPrecisionPolicy")
    
    
    return
}

func NewModifyCCPrecisionPolicyResponse() (response *ModifyCCPrecisionPolicyResponse) {
    response = &ModifyCCPrecisionPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyCCPrecisionPolicy
// 修改CC精准防护策略
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
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyCCPrecisionPolicy(request *ModifyCCPrecisionPolicyRequest) (response *ModifyCCPrecisionPolicyResponse, err error) {
    if request == nil {
        request = NewModifyCCPrecisionPolicyRequest()
    }
    
    response = NewModifyCCPrecisionPolicyResponse()
    err = c.Send(request, response)
    return
}

// ModifyCCPrecisionPolicy
// 修改CC精准防护策略
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
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyCCPrecisionPolicyWithContext(ctx context.Context, request *ModifyCCPrecisionPolicyRequest) (response *ModifyCCPrecisionPolicyResponse, err error) {
    if request == nil {
        request = NewModifyCCPrecisionPolicyRequest()
    }
    request.SetContext(ctx)
    
    response = NewModifyCCPrecisionPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewModifyCCReqLimitPolicyRequest() (request *ModifyCCReqLimitPolicyRequest) {
    request = &ModifyCCReqLimitPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("antiddos", APIVersion, "ModifyCCReqLimitPolicy")
    
    
    return
}

func NewModifyCCReqLimitPolicyResponse() (response *ModifyCCReqLimitPolicyResponse) {
    response = &ModifyCCReqLimitPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyCCReqLimitPolicy
// 修改CC频率限制策略
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
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyCCReqLimitPolicy(request *ModifyCCReqLimitPolicyRequest) (response *ModifyCCReqLimitPolicyResponse, err error) {
    if request == nil {
        request = NewModifyCCReqLimitPolicyRequest()
    }
    
    response = NewModifyCCReqLimitPolicyResponse()
    err = c.Send(request, response)
    return
}

// ModifyCCReqLimitPolicy
// 修改CC频率限制策略
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
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyCCReqLimitPolicyWithContext(ctx context.Context, request *ModifyCCReqLimitPolicyRequest) (response *ModifyCCReqLimitPolicyResponse, err error) {
    if request == nil {
        request = NewModifyCCReqLimitPolicyRequest()
    }
    request.SetContext(ctx)
    
    response = NewModifyCCReqLimitPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewModifyCCThresholdPolicyRequest() (request *ModifyCCThresholdPolicyRequest) {
    request = &ModifyCCThresholdPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("antiddos", APIVersion, "ModifyCCThresholdPolicy")
    
    
    return
}

func NewModifyCCThresholdPolicyResponse() (response *ModifyCCThresholdPolicyResponse) {
    response = &ModifyCCThresholdPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyCCThresholdPolicy
// 修改CC清洗阈值
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
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyCCThresholdPolicy(request *ModifyCCThresholdPolicyRequest) (response *ModifyCCThresholdPolicyResponse, err error) {
    if request == nil {
        request = NewModifyCCThresholdPolicyRequest()
    }
    
    response = NewModifyCCThresholdPolicyResponse()
    err = c.Send(request, response)
    return
}

// ModifyCCThresholdPolicy
// 修改CC清洗阈值
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
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyCCThresholdPolicyWithContext(ctx context.Context, request *ModifyCCThresholdPolicyRequest) (response *ModifyCCThresholdPolicyResponse, err error) {
    if request == nil {
        request = NewModifyCCThresholdPolicyRequest()
    }
    request.SetContext(ctx)
    
    response = NewModifyCCThresholdPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewModifyCcBlackWhiteIpListRequest() (request *ModifyCcBlackWhiteIpListRequest) {
    request = &ModifyCcBlackWhiteIpListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("antiddos", APIVersion, "ModifyCcBlackWhiteIpList")
    
    
    return
}

func NewModifyCcBlackWhiteIpListResponse() (response *ModifyCcBlackWhiteIpListResponse) {
    response = &ModifyCcBlackWhiteIpListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyCcBlackWhiteIpList
// 修改CC四层黑白名单
//
// 可能返回的错误码:
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyCcBlackWhiteIpList(request *ModifyCcBlackWhiteIpListRequest) (response *ModifyCcBlackWhiteIpListResponse, err error) {
    if request == nil {
        request = NewModifyCcBlackWhiteIpListRequest()
    }
    
    response = NewModifyCcBlackWhiteIpListResponse()
    err = c.Send(request, response)
    return
}

// ModifyCcBlackWhiteIpList
// 修改CC四层黑白名单
//
// 可能返回的错误码:
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyCcBlackWhiteIpListWithContext(ctx context.Context, request *ModifyCcBlackWhiteIpListRequest) (response *ModifyCcBlackWhiteIpListResponse, err error) {
    if request == nil {
        request = NewModifyCcBlackWhiteIpListRequest()
    }
    request.SetContext(ctx)
    
    response = NewModifyCcBlackWhiteIpListResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDDoSBlackWhiteIpListRequest() (request *ModifyDDoSBlackWhiteIpListRequest) {
    request = &ModifyDDoSBlackWhiteIpListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("antiddos", APIVersion, "ModifyDDoSBlackWhiteIpList")
    
    
    return
}

func NewModifyDDoSBlackWhiteIpListResponse() (response *ModifyDDoSBlackWhiteIpListResponse) {
    response = &ModifyDDoSBlackWhiteIpListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyDDoSBlackWhiteIpList
// 修改DDoS黑白名单列表
//
// 可能返回的错误码:
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyDDoSBlackWhiteIpList(request *ModifyDDoSBlackWhiteIpListRequest) (response *ModifyDDoSBlackWhiteIpListResponse, err error) {
    if request == nil {
        request = NewModifyDDoSBlackWhiteIpListRequest()
    }
    
    response = NewModifyDDoSBlackWhiteIpListResponse()
    err = c.Send(request, response)
    return
}

// ModifyDDoSBlackWhiteIpList
// 修改DDoS黑白名单列表
//
// 可能返回的错误码:
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyDDoSBlackWhiteIpListWithContext(ctx context.Context, request *ModifyDDoSBlackWhiteIpListRequest) (response *ModifyDDoSBlackWhiteIpListResponse, err error) {
    if request == nil {
        request = NewModifyDDoSBlackWhiteIpListRequest()
    }
    request.SetContext(ctx)
    
    response = NewModifyDDoSBlackWhiteIpListResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDDoSGeoIPBlockConfigRequest() (request *ModifyDDoSGeoIPBlockConfigRequest) {
    request = &ModifyDDoSGeoIPBlockConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("antiddos", APIVersion, "ModifyDDoSGeoIPBlockConfig")
    
    
    return
}

func NewModifyDDoSGeoIPBlockConfigResponse() (response *ModifyDDoSGeoIPBlockConfigResponse) {
    response = &ModifyDDoSGeoIPBlockConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyDDoSGeoIPBlockConfig
// 修改DDoS防护的区域封禁配置
//
// 可能返回的错误码:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyDDoSGeoIPBlockConfig(request *ModifyDDoSGeoIPBlockConfigRequest) (response *ModifyDDoSGeoIPBlockConfigResponse, err error) {
    if request == nil {
        request = NewModifyDDoSGeoIPBlockConfigRequest()
    }
    
    response = NewModifyDDoSGeoIPBlockConfigResponse()
    err = c.Send(request, response)
    return
}

// ModifyDDoSGeoIPBlockConfig
// 修改DDoS防护的区域封禁配置
//
// 可能返回的错误码:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyDDoSGeoIPBlockConfigWithContext(ctx context.Context, request *ModifyDDoSGeoIPBlockConfigRequest) (response *ModifyDDoSGeoIPBlockConfigResponse, err error) {
    if request == nil {
        request = NewModifyDDoSGeoIPBlockConfigRequest()
    }
    request.SetContext(ctx)
    
    response = NewModifyDDoSGeoIPBlockConfigResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDDoSLevelRequest() (request *ModifyDDoSLevelRequest) {
    request = &ModifyDDoSLevelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("antiddos", APIVersion, "ModifyDDoSLevel")
    
    
    return
}

func NewModifyDDoSLevelResponse() (response *ModifyDDoSLevelResponse) {
    response = &ModifyDDoSLevelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyDDoSLevel
// 读取或修改DDoS的防护等级
//
// 可能返回的错误码:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyDDoSLevel(request *ModifyDDoSLevelRequest) (response *ModifyDDoSLevelResponse, err error) {
    if request == nil {
        request = NewModifyDDoSLevelRequest()
    }
    
    response = NewModifyDDoSLevelResponse()
    err = c.Send(request, response)
    return
}

// ModifyDDoSLevel
// 读取或修改DDoS的防护等级
//
// 可能返回的错误码:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyDDoSLevelWithContext(ctx context.Context, request *ModifyDDoSLevelRequest) (response *ModifyDDoSLevelResponse, err error) {
    if request == nil {
        request = NewModifyDDoSLevelRequest()
    }
    request.SetContext(ctx)
    
    response = NewModifyDDoSLevelResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDDoSSpeedLimitConfigRequest() (request *ModifyDDoSSpeedLimitConfigRequest) {
    request = &ModifyDDoSSpeedLimitConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("antiddos", APIVersion, "ModifyDDoSSpeedLimitConfig")
    
    
    return
}

func NewModifyDDoSSpeedLimitConfigResponse() (response *ModifyDDoSSpeedLimitConfigResponse) {
    response = &ModifyDDoSSpeedLimitConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyDDoSSpeedLimitConfig
// 修改DDoS防护的访问限速配置
//
// 可能返回的错误码:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyDDoSSpeedLimitConfig(request *ModifyDDoSSpeedLimitConfigRequest) (response *ModifyDDoSSpeedLimitConfigResponse, err error) {
    if request == nil {
        request = NewModifyDDoSSpeedLimitConfigRequest()
    }
    
    response = NewModifyDDoSSpeedLimitConfigResponse()
    err = c.Send(request, response)
    return
}

// ModifyDDoSSpeedLimitConfig
// 修改DDoS防护的访问限速配置
//
// 可能返回的错误码:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyDDoSSpeedLimitConfigWithContext(ctx context.Context, request *ModifyDDoSSpeedLimitConfigRequest) (response *ModifyDDoSSpeedLimitConfigResponse, err error) {
    if request == nil {
        request = NewModifyDDoSSpeedLimitConfigRequest()
    }
    request.SetContext(ctx)
    
    response = NewModifyDDoSSpeedLimitConfigResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDDoSThresholdRequest() (request *ModifyDDoSThresholdRequest) {
    request = &ModifyDDoSThresholdRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("antiddos", APIVersion, "ModifyDDoSThreshold")
    
    
    return
}

func NewModifyDDoSThresholdResponse() (response *ModifyDDoSThresholdResponse) {
    response = &ModifyDDoSThresholdResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyDDoSThreshold
// 修改DDoS清洗阈值
//
// 可能返回的错误码:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyDDoSThreshold(request *ModifyDDoSThresholdRequest) (response *ModifyDDoSThresholdResponse, err error) {
    if request == nil {
        request = NewModifyDDoSThresholdRequest()
    }
    
    response = NewModifyDDoSThresholdResponse()
    err = c.Send(request, response)
    return
}

// ModifyDDoSThreshold
// 修改DDoS清洗阈值
//
// 可能返回的错误码:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyDDoSThresholdWithContext(ctx context.Context, request *ModifyDDoSThresholdRequest) (response *ModifyDDoSThresholdResponse, err error) {
    if request == nil {
        request = NewModifyDDoSThresholdRequest()
    }
    request.SetContext(ctx)
    
    response = NewModifyDDoSThresholdResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDomainUsrNameRequest() (request *ModifyDomainUsrNameRequest) {
    request = &ModifyDomainUsrNameRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("antiddos", APIVersion, "ModifyDomainUsrName")
    
    
    return
}

func NewModifyDomainUsrNameResponse() (response *ModifyDomainUsrNameResponse) {
    response = &ModifyDomainUsrNameResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyDomainUsrName
// 修改智能解析域名名称
//
// 可能返回的错误码:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyDomainUsrName(request *ModifyDomainUsrNameRequest) (response *ModifyDomainUsrNameResponse, err error) {
    if request == nil {
        request = NewModifyDomainUsrNameRequest()
    }
    
    response = NewModifyDomainUsrNameResponse()
    err = c.Send(request, response)
    return
}

// ModifyDomainUsrName
// 修改智能解析域名名称
//
// 可能返回的错误码:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyDomainUsrNameWithContext(ctx context.Context, request *ModifyDomainUsrNameRequest) (response *ModifyDomainUsrNameResponse, err error) {
    if request == nil {
        request = NewModifyDomainUsrNameRequest()
    }
    request.SetContext(ctx)
    
    response = NewModifyDomainUsrNameResponse()
    err = c.Send(request, response)
    return
}

func NewModifyL7RulesEdgeRequest() (request *ModifyL7RulesEdgeRequest) {
    request = &ModifyL7RulesEdgeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("antiddos", APIVersion, "ModifyL7RulesEdge")
    
    
    return
}

func NewModifyL7RulesEdgeResponse() (response *ModifyL7RulesEdgeResponse) {
    response = &ModifyL7RulesEdgeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyL7RulesEdge
// 修改边界防护L7转发规则
//
// 可能返回的错误码:
//  RESOURCEINUSE = "ResourceInUse"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyL7RulesEdge(request *ModifyL7RulesEdgeRequest) (response *ModifyL7RulesEdgeResponse, err error) {
    if request == nil {
        request = NewModifyL7RulesEdgeRequest()
    }
    
    response = NewModifyL7RulesEdgeResponse()
    err = c.Send(request, response)
    return
}

// ModifyL7RulesEdge
// 修改边界防护L7转发规则
//
// 可能返回的错误码:
//  RESOURCEINUSE = "ResourceInUse"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyL7RulesEdgeWithContext(ctx context.Context, request *ModifyL7RulesEdgeRequest) (response *ModifyL7RulesEdgeResponse, err error) {
    if request == nil {
        request = NewModifyL7RulesEdgeRequest()
    }
    request.SetContext(ctx)
    
    response = NewModifyL7RulesEdgeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyNewDomainRulesRequest() (request *ModifyNewDomainRulesRequest) {
    request = &ModifyNewDomainRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("antiddos", APIVersion, "ModifyNewDomainRules")
    
    
    return
}

func NewModifyNewDomainRulesResponse() (response *ModifyNewDomainRulesResponse) {
    response = &ModifyNewDomainRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyNewDomainRules
// 修改7层转发规则
//
// 可能返回的错误码:
//  RESOURCEINUSE = "ResourceInUse"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyNewDomainRules(request *ModifyNewDomainRulesRequest) (response *ModifyNewDomainRulesResponse, err error) {
    if request == nil {
        request = NewModifyNewDomainRulesRequest()
    }
    
    response = NewModifyNewDomainRulesResponse()
    err = c.Send(request, response)
    return
}

// ModifyNewDomainRules
// 修改7层转发规则
//
// 可能返回的错误码:
//  RESOURCEINUSE = "ResourceInUse"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyNewDomainRulesWithContext(ctx context.Context, request *ModifyNewDomainRulesRequest) (response *ModifyNewDomainRulesResponse, err error) {
    if request == nil {
        request = NewModifyNewDomainRulesRequest()
    }
    request.SetContext(ctx)
    
    response = NewModifyNewDomainRulesResponse()
    err = c.Send(request, response)
    return
}

func NewModifyPacketFilterConfigRequest() (request *ModifyPacketFilterConfigRequest) {
    request = &ModifyPacketFilterConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("antiddos", APIVersion, "ModifyPacketFilterConfig")
    
    
    return
}

func NewModifyPacketFilterConfigResponse() (response *ModifyPacketFilterConfigResponse) {
    response = &ModifyPacketFilterConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyPacketFilterConfig
// 修改DDoS防护的特征过滤规则
//
// 可能返回的错误码:
//  RESOURCEINUSE = "ResourceInUse"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyPacketFilterConfig(request *ModifyPacketFilterConfigRequest) (response *ModifyPacketFilterConfigResponse, err error) {
    if request == nil {
        request = NewModifyPacketFilterConfigRequest()
    }
    
    response = NewModifyPacketFilterConfigResponse()
    err = c.Send(request, response)
    return
}

// ModifyPacketFilterConfig
// 修改DDoS防护的特征过滤规则
//
// 可能返回的错误码:
//  RESOURCEINUSE = "ResourceInUse"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyPacketFilterConfigWithContext(ctx context.Context, request *ModifyPacketFilterConfigRequest) (response *ModifyPacketFilterConfigResponse, err error) {
    if request == nil {
        request = NewModifyPacketFilterConfigRequest()
    }
    request.SetContext(ctx)
    
    response = NewModifyPacketFilterConfigResponse()
    err = c.Send(request, response)
    return
}

func NewModifyPortAclConfigRequest() (request *ModifyPortAclConfigRequest) {
    request = &ModifyPortAclConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("antiddos", APIVersion, "ModifyPortAclConfig")
    
    
    return
}

func NewModifyPortAclConfigResponse() (response *ModifyPortAclConfigResponse) {
    response = &ModifyPortAclConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyPortAclConfig
// 修改DDoS防护的端口acl策略
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
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyPortAclConfig(request *ModifyPortAclConfigRequest) (response *ModifyPortAclConfigResponse, err error) {
    if request == nil {
        request = NewModifyPortAclConfigRequest()
    }
    
    response = NewModifyPortAclConfigResponse()
    err = c.Send(request, response)
    return
}

// ModifyPortAclConfig
// 修改DDoS防护的端口acl策略
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
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyPortAclConfigWithContext(ctx context.Context, request *ModifyPortAclConfigRequest) (response *ModifyPortAclConfigResponse, err error) {
    if request == nil {
        request = NewModifyPortAclConfigRequest()
    }
    request.SetContext(ctx)
    
    response = NewModifyPortAclConfigResponse()
    err = c.Send(request, response)
    return
}

func NewSwitchWaterPrintConfigRequest() (request *SwitchWaterPrintConfigRequest) {
    request = &SwitchWaterPrintConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("antiddos", APIVersion, "SwitchWaterPrintConfig")
    
    
    return
}

func NewSwitchWaterPrintConfigResponse() (response *SwitchWaterPrintConfigResponse) {
    response = &SwitchWaterPrintConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// SwitchWaterPrintConfig
// 开启或关闭DDoS防护的水印防护配置
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) SwitchWaterPrintConfig(request *SwitchWaterPrintConfigRequest) (response *SwitchWaterPrintConfigResponse, err error) {
    if request == nil {
        request = NewSwitchWaterPrintConfigRequest()
    }
    
    response = NewSwitchWaterPrintConfigResponse()
    err = c.Send(request, response)
    return
}

// SwitchWaterPrintConfig
// 开启或关闭DDoS防护的水印防护配置
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) SwitchWaterPrintConfigWithContext(ctx context.Context, request *SwitchWaterPrintConfigRequest) (response *SwitchWaterPrintConfigResponse, err error) {
    if request == nil {
        request = NewSwitchWaterPrintConfigRequest()
    }
    request.SetContext(ctx)
    
    response = NewSwitchWaterPrintConfigResponse()
    err = c.Send(request, response)
    return
}
