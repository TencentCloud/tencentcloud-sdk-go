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
    "errors"
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
    return c.AssociateDDoSEipAddressWithContext(context.Background(), request)
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("AssociateDDoSEipAddress require credential")
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
    return c.AssociateDDoSEipLoadBalancerWithContext(context.Background(), request)
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("AssociateDDoSEipLoadBalancer require credential")
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
    return c.CreateBlackWhiteIpListWithContext(context.Background(), request)
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateBlackWhiteIpList require credential")
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
    return c.CreateBoundIPWithContext(context.Background(), request)
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateBoundIP require credential")
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
    return c.CreateCCPrecisionPolicyWithContext(context.Background(), request)
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCCPrecisionPolicy require credential")
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
    return c.CreateCCReqLimitPolicyWithContext(context.Background(), request)
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCCReqLimitPolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateCCReqLimitPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewCreateCcBlackWhiteIpListRequest() (request *CreateCcBlackWhiteIpListRequest) {
    request = &CreateCcBlackWhiteIpListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("antiddos", APIVersion, "CreateCcBlackWhiteIpList")
    
    
    return
}

func NewCreateCcBlackWhiteIpListResponse() (response *CreateCcBlackWhiteIpListResponse) {
    response = &CreateCcBlackWhiteIpListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateCcBlackWhiteIpList
// 新建CC四层黑白名单
//
// 可能返回的错误码:
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateCcBlackWhiteIpList(request *CreateCcBlackWhiteIpListRequest) (response *CreateCcBlackWhiteIpListResponse, err error) {
    return c.CreateCcBlackWhiteIpListWithContext(context.Background(), request)
}

// CreateCcBlackWhiteIpList
// 新建CC四层黑白名单
//
// 可能返回的错误码:
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateCcBlackWhiteIpListWithContext(ctx context.Context, request *CreateCcBlackWhiteIpListRequest) (response *CreateCcBlackWhiteIpListResponse, err error) {
    if request == nil {
        request = NewCreateCcBlackWhiteIpListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCcBlackWhiteIpList require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateCcBlackWhiteIpListResponse()
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
    return c.CreateCcGeoIPBlockConfigWithContext(context.Background(), request)
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCcGeoIPBlockConfig require credential")
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
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateDDoSAI(request *CreateDDoSAIRequest) (response *CreateDDoSAIResponse, err error) {
    return c.CreateDDoSAIWithContext(context.Background(), request)
}

// CreateDDoSAI
// 设置DDoS防护的AI防护开关
//
// 可能返回的错误码:
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateDDoSAIWithContext(ctx context.Context, request *CreateDDoSAIRequest) (response *CreateDDoSAIResponse, err error) {
    if request == nil {
        request = NewCreateDDoSAIRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateDDoSAI require credential")
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
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateDDoSBlackWhiteIpList(request *CreateDDoSBlackWhiteIpListRequest) (response *CreateDDoSBlackWhiteIpListResponse, err error) {
    return c.CreateDDoSBlackWhiteIpListWithContext(context.Background(), request)
}

// CreateDDoSBlackWhiteIpList
// 添加DDoS防护的IP网段黑白名单
//
// 可能返回的错误码:
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateDDoSBlackWhiteIpListWithContext(ctx context.Context, request *CreateDDoSBlackWhiteIpListRequest) (response *CreateDDoSBlackWhiteIpListResponse, err error) {
    if request == nil {
        request = NewCreateDDoSBlackWhiteIpListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateDDoSBlackWhiteIpList require credential")
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
    return c.CreateDDoSConnectLimitWithContext(context.Background(), request)
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateDDoSConnectLimit require credential")
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
    return c.CreateDDoSGeoIPBlockConfigWithContext(context.Background(), request)
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateDDoSGeoIPBlockConfig require credential")
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
    return c.CreateDDoSSpeedLimitConfigWithContext(context.Background(), request)
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateDDoSSpeedLimitConfig require credential")
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
    return c.CreateDefaultAlarmThresholdWithContext(context.Background(), request)
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateDefaultAlarmThreshold require credential")
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
    return c.CreateIPAlarmThresholdConfigWithContext(context.Background(), request)
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateIPAlarmThresholdConfig require credential")
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
    return c.CreateL7RuleCertsWithContext(context.Background(), request)
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateL7RuleCerts require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateL7RuleCertsResponse()
    err = c.Send(request, response)
    return
}

func NewCreateNewL7RulesRequest() (request *CreateNewL7RulesRequest) {
    request = &CreateNewL7RulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("antiddos", APIVersion, "CreateNewL7Rules")
    
    
    return
}

func NewCreateNewL7RulesResponse() (response *CreateNewL7RulesResponse) {
    response = &CreateNewL7RulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateNewL7Rules
// 添加7层转发规则
//
// 可能返回的错误码:
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateNewL7Rules(request *CreateNewL7RulesRequest) (response *CreateNewL7RulesResponse, err error) {
    return c.CreateNewL7RulesWithContext(context.Background(), request)
}

// CreateNewL7Rules
// 添加7层转发规则
//
// 可能返回的错误码:
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateNewL7RulesWithContext(ctx context.Context, request *CreateNewL7RulesRequest) (response *CreateNewL7RulesResponse, err error) {
    if request == nil {
        request = NewCreateNewL7RulesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateNewL7Rules require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateNewL7RulesResponse()
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
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreatePacketFilterConfig(request *CreatePacketFilterConfigRequest) (response *CreatePacketFilterConfigResponse, err error) {
    return c.CreatePacketFilterConfigWithContext(context.Background(), request)
}

// CreatePacketFilterConfig
// 添加DDoS防护的特征过滤规则
//
// 可能返回的错误码:
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreatePacketFilterConfigWithContext(ctx context.Context, request *CreatePacketFilterConfigRequest) (response *CreatePacketFilterConfigResponse, err error) {
    if request == nil {
        request = NewCreatePacketFilterConfigRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreatePacketFilterConfig require credential")
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
    return c.CreatePortAclConfigWithContext(context.Background(), request)
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreatePortAclConfig require credential")
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
    return c.CreatePortAclConfigListWithContext(context.Background(), request)
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreatePortAclConfigList require credential")
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
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateProtocolBlockConfig(request *CreateProtocolBlockConfigRequest) (response *CreateProtocolBlockConfigResponse, err error) {
    return c.CreateProtocolBlockConfigWithContext(context.Background(), request)
}

// CreateProtocolBlockConfig
// 设置DDoS防护的协议封禁配置
//
// 可能返回的错误码:
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateProtocolBlockConfigWithContext(ctx context.Context, request *CreateProtocolBlockConfigRequest) (response *CreateProtocolBlockConfigResponse, err error) {
    if request == nil {
        request = NewCreateProtocolBlockConfigRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateProtocolBlockConfig require credential")
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
    return c.CreateSchedulingDomainWithContext(context.Background(), request)
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateSchedulingDomain require credential")
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
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateWaterPrintConfig(request *CreateWaterPrintConfigRequest) (response *CreateWaterPrintConfigResponse, err error) {
    return c.CreateWaterPrintConfigWithContext(context.Background(), request)
}

// CreateWaterPrintConfig
// 添加DDoS防护的水印防护配置
//
// 可能返回的错误码:
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateWaterPrintConfigWithContext(ctx context.Context, request *CreateWaterPrintConfigRequest) (response *CreateWaterPrintConfigResponse, err error) {
    if request == nil {
        request = NewCreateWaterPrintConfigRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateWaterPrintConfig require credential")
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
    return c.CreateWaterPrintKeyWithContext(context.Background(), request)
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateWaterPrintKey require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateWaterPrintKeyResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteCCLevelPolicyRequest() (request *DeleteCCLevelPolicyRequest) {
    request = &DeleteCCLevelPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("antiddos", APIVersion, "DeleteCCLevelPolicy")
    
    
    return
}

func NewDeleteCCLevelPolicyResponse() (response *DeleteCCLevelPolicyResponse) {
    response = &DeleteCCLevelPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteCCLevelPolicy
// 删除CC分级策略
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
func (c *Client) DeleteCCLevelPolicy(request *DeleteCCLevelPolicyRequest) (response *DeleteCCLevelPolicyResponse, err error) {
    return c.DeleteCCLevelPolicyWithContext(context.Background(), request)
}

// DeleteCCLevelPolicy
// 删除CC分级策略
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
func (c *Client) DeleteCCLevelPolicyWithContext(ctx context.Context, request *DeleteCCLevelPolicyRequest) (response *DeleteCCLevelPolicyResponse, err error) {
    if request == nil {
        request = NewDeleteCCLevelPolicyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteCCLevelPolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteCCLevelPolicyResponse()
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
    return c.DeleteCCPrecisionPolicyWithContext(context.Background(), request)
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteCCPrecisionPolicy require credential")
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
    return c.DeleteCCRequestLimitPolicyWithContext(context.Background(), request)
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteCCRequestLimitPolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteCCRequestLimitPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteCCThresholdPolicyRequest() (request *DeleteCCThresholdPolicyRequest) {
    request = &DeleteCCThresholdPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("antiddos", APIVersion, "DeleteCCThresholdPolicy")
    
    
    return
}

func NewDeleteCCThresholdPolicyResponse() (response *DeleteCCThresholdPolicyResponse) {
    response = &DeleteCCThresholdPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteCCThresholdPolicy
// 删除CC清洗阈值策略
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
func (c *Client) DeleteCCThresholdPolicy(request *DeleteCCThresholdPolicyRequest) (response *DeleteCCThresholdPolicyResponse, err error) {
    return c.DeleteCCThresholdPolicyWithContext(context.Background(), request)
}

// DeleteCCThresholdPolicy
// 删除CC清洗阈值策略
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
func (c *Client) DeleteCCThresholdPolicyWithContext(ctx context.Context, request *DeleteCCThresholdPolicyRequest) (response *DeleteCCThresholdPolicyResponse, err error) {
    if request == nil {
        request = NewDeleteCCThresholdPolicyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteCCThresholdPolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteCCThresholdPolicyResponse()
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
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteCcBlackWhiteIpList(request *DeleteCcBlackWhiteIpListRequest) (response *DeleteCcBlackWhiteIpListResponse, err error) {
    return c.DeleteCcBlackWhiteIpListWithContext(context.Background(), request)
}

// DeleteCcBlackWhiteIpList
// 删除CC四层黑白名单
//
// 可能返回的错误码:
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteCcBlackWhiteIpListWithContext(ctx context.Context, request *DeleteCcBlackWhiteIpListRequest) (response *DeleteCcBlackWhiteIpListResponse, err error) {
    if request == nil {
        request = NewDeleteCcBlackWhiteIpListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteCcBlackWhiteIpList require credential")
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
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteCcGeoIPBlockConfig(request *DeleteCcGeoIPBlockConfigRequest) (response *DeleteCcGeoIPBlockConfigResponse, err error) {
    return c.DeleteCcGeoIPBlockConfigWithContext(context.Background(), request)
}

// DeleteCcGeoIPBlockConfig
// 删除CC防护的区域封禁配置
//
// 可能返回的错误码:
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteCcGeoIPBlockConfigWithContext(ctx context.Context, request *DeleteCcGeoIPBlockConfigRequest) (response *DeleteCcGeoIPBlockConfigResponse, err error) {
    if request == nil {
        request = NewDeleteCcGeoIPBlockConfigRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteCcGeoIPBlockConfig require credential")
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
    return c.DeleteDDoSBlackWhiteIpListWithContext(context.Background(), request)
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteDDoSBlackWhiteIpList require credential")
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
    return c.DeleteDDoSGeoIPBlockConfigWithContext(context.Background(), request)
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteDDoSGeoIPBlockConfig require credential")
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
    return c.DeleteDDoSSpeedLimitConfigWithContext(context.Background(), request)
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteDDoSSpeedLimitConfig require credential")
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
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeletePacketFilterConfig(request *DeletePacketFilterConfigRequest) (response *DeletePacketFilterConfigResponse, err error) {
    return c.DeletePacketFilterConfigWithContext(context.Background(), request)
}

// DeletePacketFilterConfig
// 删除DDoS防护的特征过滤规则
//
// 可能返回的错误码:
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeletePacketFilterConfigWithContext(ctx context.Context, request *DeletePacketFilterConfigRequest) (response *DeletePacketFilterConfigResponse, err error) {
    if request == nil {
        request = NewDeletePacketFilterConfigRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeletePacketFilterConfig require credential")
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
    return c.DeletePortAclConfigWithContext(context.Background(), request)
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeletePortAclConfig require credential")
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
    return c.DeleteWaterPrintConfigWithContext(context.Background(), request)
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteWaterPrintConfig require credential")
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
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteWaterPrintKey(request *DeleteWaterPrintKeyRequest) (response *DeleteWaterPrintKeyResponse, err error) {
    return c.DeleteWaterPrintKeyWithContext(context.Background(), request)
}

// DeleteWaterPrintKey
// 删除DDoS防护的水印防护密钥
//
// 可能返回的错误码:
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteWaterPrintKeyWithContext(ctx context.Context, request *DeleteWaterPrintKeyRequest) (response *DeleteWaterPrintKeyResponse, err error) {
    if request == nil {
        request = NewDeleteWaterPrintKeyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteWaterPrintKey require credential")
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
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeBasicDeviceStatus(request *DescribeBasicDeviceStatusRequest) (response *DescribeBasicDeviceStatusResponse, err error) {
    return c.DescribeBasicDeviceStatusWithContext(context.Background(), request)
}

// DescribeBasicDeviceStatus
// 获取基础防护攻击状态
//
// 可能返回的错误码:
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeBasicDeviceStatusWithContext(ctx context.Context, request *DescribeBasicDeviceStatusRequest) (response *DescribeBasicDeviceStatusResponse, err error) {
    if request == nil {
        request = NewDescribeBasicDeviceStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBasicDeviceStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBasicDeviceStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBgpBizTrendRequest() (request *DescribeBgpBizTrendRequest) {
    request = &DescribeBgpBizTrendRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("antiddos", APIVersion, "DescribeBgpBizTrend")
    
    
    return
}

func NewDescribeBgpBizTrendResponse() (response *DescribeBgpBizTrendResponse) {
    response = &DescribeBgpBizTrendResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeBgpBizTrend
// 获取高防包流量折线图
//
// 可能返回的错误码:
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeBgpBizTrend(request *DescribeBgpBizTrendRequest) (response *DescribeBgpBizTrendResponse, err error) {
    return c.DescribeBgpBizTrendWithContext(context.Background(), request)
}

// DescribeBgpBizTrend
// 获取高防包流量折线图
//
// 可能返回的错误码:
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeBgpBizTrendWithContext(ctx context.Context, request *DescribeBgpBizTrendRequest) (response *DescribeBgpBizTrendResponse, err error) {
    if request == nil {
        request = NewDescribeBgpBizTrendRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBgpBizTrend require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBgpBizTrendResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBizHttpStatusRequest() (request *DescribeBizHttpStatusRequest) {
    request = &DescribeBizHttpStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("antiddos", APIVersion, "DescribeBizHttpStatus")
    
    
    return
}

func NewDescribeBizHttpStatusResponse() (response *DescribeBizHttpStatusResponse) {
    response = &DescribeBizHttpStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeBizHttpStatus
// 获取业务流量状态码统计列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeBizHttpStatus(request *DescribeBizHttpStatusRequest) (response *DescribeBizHttpStatusResponse, err error) {
    return c.DescribeBizHttpStatusWithContext(context.Background(), request)
}

// DescribeBizHttpStatus
// 获取业务流量状态码统计列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeBizHttpStatusWithContext(ctx context.Context, request *DescribeBizHttpStatusRequest) (response *DescribeBizHttpStatusResponse, err error) {
    if request == nil {
        request = NewDescribeBizHttpStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBizHttpStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBizHttpStatusResponse()
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
    return c.DescribeBizTrendWithContext(context.Background(), request)
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBizTrend require credential")
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
    return c.DescribeBlackWhiteIpListWithContext(context.Background(), request)
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBlackWhiteIpList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBlackWhiteIpListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCCLevelListRequest() (request *DescribeCCLevelListRequest) {
    request = &DescribeCCLevelListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("antiddos", APIVersion, "DescribeCCLevelList")
    
    
    return
}

func NewDescribeCCLevelListResponse() (response *DescribeCCLevelListResponse) {
    response = &DescribeCCLevelListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCCLevelList
// 获取边界防护CC防护等级列表
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
func (c *Client) DescribeCCLevelList(request *DescribeCCLevelListRequest) (response *DescribeCCLevelListResponse, err error) {
    return c.DescribeCCLevelListWithContext(context.Background(), request)
}

// DescribeCCLevelList
// 获取边界防护CC防护等级列表
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
func (c *Client) DescribeCCLevelListWithContext(ctx context.Context, request *DescribeCCLevelListRequest) (response *DescribeCCLevelListResponse, err error) {
    if request == nil {
        request = NewDescribeCCLevelListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCCLevelList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCCLevelListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCCLevelPolicyRequest() (request *DescribeCCLevelPolicyRequest) {
    request = &DescribeCCLevelPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("antiddos", APIVersion, "DescribeCCLevelPolicy")
    
    
    return
}

func NewDescribeCCLevelPolicyResponse() (response *DescribeCCLevelPolicyResponse) {
    response = &DescribeCCLevelPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCCLevelPolicy
// 获取CC分级策略
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
func (c *Client) DescribeCCLevelPolicy(request *DescribeCCLevelPolicyRequest) (response *DescribeCCLevelPolicyResponse, err error) {
    return c.DescribeCCLevelPolicyWithContext(context.Background(), request)
}

// DescribeCCLevelPolicy
// 获取CC分级策略
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
func (c *Client) DescribeCCLevelPolicyWithContext(ctx context.Context, request *DescribeCCLevelPolicyRequest) (response *DescribeCCLevelPolicyResponse, err error) {
    if request == nil {
        request = NewDescribeCCLevelPolicyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCCLevelPolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCCLevelPolicyResponse()
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
    return c.DescribeCCPrecisionPlyListWithContext(context.Background(), request)
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCCPrecisionPlyList require credential")
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
    return c.DescribeCCReqLimitPolicyListWithContext(context.Background(), request)
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCCReqLimitPolicyList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCCReqLimitPolicyListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCCThresholdListRequest() (request *DescribeCCThresholdListRequest) {
    request = &DescribeCCThresholdListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("antiddos", APIVersion, "DescribeCCThresholdList")
    
    
    return
}

func NewDescribeCCThresholdListResponse() (response *DescribeCCThresholdListResponse) {
    response = &DescribeCCThresholdListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCCThresholdList
// 获取CC清洗阈值列表
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
func (c *Client) DescribeCCThresholdList(request *DescribeCCThresholdListRequest) (response *DescribeCCThresholdListResponse, err error) {
    return c.DescribeCCThresholdListWithContext(context.Background(), request)
}

// DescribeCCThresholdList
// 获取CC清洗阈值列表
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
func (c *Client) DescribeCCThresholdListWithContext(ctx context.Context, request *DescribeCCThresholdListRequest) (response *DescribeCCThresholdListResponse, err error) {
    if request == nil {
        request = NewDescribeCCThresholdListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCCThresholdList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCCThresholdListResponse()
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
    return c.DescribeCCTrendWithContext(context.Background(), request)
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCCTrend require credential")
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
    return c.DescribeCcBlackWhiteIpListWithContext(context.Background(), request)
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCcBlackWhiteIpList require credential")
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
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeCcGeoIPBlockConfigList(request *DescribeCcGeoIPBlockConfigListRequest) (response *DescribeCcGeoIPBlockConfigListResponse, err error) {
    return c.DescribeCcGeoIPBlockConfigListWithContext(context.Background(), request)
}

// DescribeCcGeoIPBlockConfigList
// 获取CC防护的区域封禁配置列表
//
// 可能返回的错误码:
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeCcGeoIPBlockConfigListWithContext(ctx context.Context, request *DescribeCcGeoIPBlockConfigListRequest) (response *DescribeCcGeoIPBlockConfigListResponse, err error) {
    if request == nil {
        request = NewDescribeCcGeoIPBlockConfigListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCcGeoIPBlockConfigList require credential")
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
    return c.DescribeDDoSBlackWhiteIpListWithContext(context.Background(), request)
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDDoSBlackWhiteIpList require credential")
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
    return c.DescribeDDoSConnectLimitListWithContext(context.Background(), request)
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDDoSConnectLimitList require credential")
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
    return c.DescribeDDoSTrendWithContext(context.Background(), request)
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDDoSTrend require credential")
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
    return c.DescribeDefaultAlarmThresholdWithContext(context.Background(), request)
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDefaultAlarmThreshold require credential")
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
    return c.DescribeL7RulesBySSLCertIdWithContext(context.Background(), request)
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeL7RulesBySSLCertId require credential")
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
//  AUTHFAILURE = "AuthFailure"
func (c *Client) DescribeListBGPIPInstances(request *DescribeListBGPIPInstancesRequest) (response *DescribeListBGPIPInstancesResponse, err error) {
    return c.DescribeListBGPIPInstancesWithContext(context.Background(), request)
}

// DescribeListBGPIPInstances
// 获取高防IP资产实例列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
func (c *Client) DescribeListBGPIPInstancesWithContext(ctx context.Context, request *DescribeListBGPIPInstancesRequest) (response *DescribeListBGPIPInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeListBGPIPInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeListBGPIPInstances require credential")
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
//  AUTHFAILURE = "AuthFailure"
func (c *Client) DescribeListBGPInstances(request *DescribeListBGPInstancesRequest) (response *DescribeListBGPInstancesResponse, err error) {
    return c.DescribeListBGPInstancesWithContext(context.Background(), request)
}

// DescribeListBGPInstances
// 获取高防包资产实例列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
func (c *Client) DescribeListBGPInstancesWithContext(ctx context.Context, request *DescribeListBGPInstancesRequest) (response *DescribeListBGPInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeListBGPInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeListBGPInstances require credential")
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
//  AUTHFAILURE = "AuthFailure"
func (c *Client) DescribeListBlackWhiteIpList(request *DescribeListBlackWhiteIpListRequest) (response *DescribeListBlackWhiteIpListResponse, err error) {
    return c.DescribeListBlackWhiteIpListWithContext(context.Background(), request)
}

// DescribeListBlackWhiteIpList
// 获取DDoS防护的IP黑白名单列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
func (c *Client) DescribeListBlackWhiteIpListWithContext(ctx context.Context, request *DescribeListBlackWhiteIpListRequest) (response *DescribeListBlackWhiteIpListResponse, err error) {
    if request == nil {
        request = NewDescribeListBlackWhiteIpListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeListBlackWhiteIpList require credential")
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
//  AUTHFAILURE = "AuthFailure"
func (c *Client) DescribeListDDoSAI(request *DescribeListDDoSAIRequest) (response *DescribeListDDoSAIResponse, err error) {
    return c.DescribeListDDoSAIWithContext(context.Background(), request)
}

// DescribeListDDoSAI
// 获取DDoS防护的AI防护开关列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
func (c *Client) DescribeListDDoSAIWithContext(ctx context.Context, request *DescribeListDDoSAIRequest) (response *DescribeListDDoSAIResponse, err error) {
    if request == nil {
        request = NewDescribeListDDoSAIRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeListDDoSAI require credential")
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
//  AUTHFAILURE = "AuthFailure"
func (c *Client) DescribeListDDoSGeoIPBlockConfig(request *DescribeListDDoSGeoIPBlockConfigRequest) (response *DescribeListDDoSGeoIPBlockConfigResponse, err error) {
    return c.DescribeListDDoSGeoIPBlockConfigWithContext(context.Background(), request)
}

// DescribeListDDoSGeoIPBlockConfig
// 获取DDoS防护的区域封禁配置列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
func (c *Client) DescribeListDDoSGeoIPBlockConfigWithContext(ctx context.Context, request *DescribeListDDoSGeoIPBlockConfigRequest) (response *DescribeListDDoSGeoIPBlockConfigResponse, err error) {
    if request == nil {
        request = NewDescribeListDDoSGeoIPBlockConfigRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeListDDoSGeoIPBlockConfig require credential")
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
//  AUTHFAILURE = "AuthFailure"
func (c *Client) DescribeListDDoSSpeedLimitConfig(request *DescribeListDDoSSpeedLimitConfigRequest) (response *DescribeListDDoSSpeedLimitConfigResponse, err error) {
    return c.DescribeListDDoSSpeedLimitConfigWithContext(context.Background(), request)
}

// DescribeListDDoSSpeedLimitConfig
// 获取DDoS防护的访问限速配置列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
func (c *Client) DescribeListDDoSSpeedLimitConfigWithContext(ctx context.Context, request *DescribeListDDoSSpeedLimitConfigRequest) (response *DescribeListDDoSSpeedLimitConfigResponse, err error) {
    if request == nil {
        request = NewDescribeListDDoSSpeedLimitConfigRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeListDDoSSpeedLimitConfig require credential")
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
//  AUTHFAILURE = "AuthFailure"
func (c *Client) DescribeListIPAlarmConfig(request *DescribeListIPAlarmConfigRequest) (response *DescribeListIPAlarmConfigResponse, err error) {
    return c.DescribeListIPAlarmConfigWithContext(context.Background(), request)
}

// DescribeListIPAlarmConfig
// 获取单IP告警阈值配置列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
func (c *Client) DescribeListIPAlarmConfigWithContext(ctx context.Context, request *DescribeListIPAlarmConfigRequest) (response *DescribeListIPAlarmConfigResponse, err error) {
    if request == nil {
        request = NewDescribeListIPAlarmConfigRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeListIPAlarmConfig require credential")
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
//  AUTHFAILURE = "AuthFailure"
func (c *Client) DescribeListListener(request *DescribeListListenerRequest) (response *DescribeListListenerResponse, err error) {
    return c.DescribeListListenerWithContext(context.Background(), request)
}

// DescribeListListener
// 获取转发监听器列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
func (c *Client) DescribeListListenerWithContext(ctx context.Context, request *DescribeListListenerRequest) (response *DescribeListListenerResponse, err error) {
    if request == nil {
        request = NewDescribeListListenerRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeListListener require credential")
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
//  AUTHFAILURE = "AuthFailure"
func (c *Client) DescribeListPacketFilterConfig(request *DescribeListPacketFilterConfigRequest) (response *DescribeListPacketFilterConfigResponse, err error) {
    return c.DescribeListPacketFilterConfigWithContext(context.Background(), request)
}

// DescribeListPacketFilterConfig
// 获取DDoS防护的特征过滤规则列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
func (c *Client) DescribeListPacketFilterConfigWithContext(ctx context.Context, request *DescribeListPacketFilterConfigRequest) (response *DescribeListPacketFilterConfigResponse, err error) {
    if request == nil {
        request = NewDescribeListPacketFilterConfigRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeListPacketFilterConfig require credential")
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
    return c.DescribeListPortAclListWithContext(context.Background(), request)
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeListPortAclList require credential")
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
    return c.DescribeListProtectThresholdConfigWithContext(context.Background(), request)
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeListProtectThresholdConfig require credential")
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
    return c.DescribeListProtocolBlockConfigWithContext(context.Background(), request)
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeListProtocolBlockConfig require credential")
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
    return c.DescribeListSchedulingDomainWithContext(context.Background(), request)
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeListSchedulingDomain require credential")
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
    return c.DescribeListWaterPrintConfigWithContext(context.Background(), request)
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeListWaterPrintConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeListWaterPrintConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeNewL7RulesRequest() (request *DescribeNewL7RulesRequest) {
    request = &DescribeNewL7RulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("antiddos", APIVersion, "DescribeNewL7Rules")
    
    
    return
}

func NewDescribeNewL7RulesResponse() (response *DescribeNewL7RulesResponse) {
    response = &DescribeNewL7RulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeNewL7Rules
// 高防IP获取7层规则
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
func (c *Client) DescribeNewL7Rules(request *DescribeNewL7RulesRequest) (response *DescribeNewL7RulesResponse, err error) {
    return c.DescribeNewL7RulesWithContext(context.Background(), request)
}

// DescribeNewL7Rules
// 高防IP获取7层规则
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
func (c *Client) DescribeNewL7RulesWithContext(ctx context.Context, request *DescribeNewL7RulesRequest) (response *DescribeNewL7RulesResponse, err error) {
    if request == nil {
        request = NewDescribeNewL7RulesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeNewL7Rules require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeNewL7RulesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeNewL7RulesErrHealthRequest() (request *DescribeNewL7RulesErrHealthRequest) {
    request = &DescribeNewL7RulesErrHealthRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("antiddos", APIVersion, "DescribeNewL7RulesErrHealth")
    
    
    return
}

func NewDescribeNewL7RulesErrHealthResponse() (response *DescribeNewL7RulesErrHealthResponse) {
    response = &DescribeNewL7RulesErrHealthResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeNewL7RulesErrHealth
// 获取L7转发规则健康检查异常结果列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
func (c *Client) DescribeNewL7RulesErrHealth(request *DescribeNewL7RulesErrHealthRequest) (response *DescribeNewL7RulesErrHealthResponse, err error) {
    return c.DescribeNewL7RulesErrHealthWithContext(context.Background(), request)
}

// DescribeNewL7RulesErrHealth
// 获取L7转发规则健康检查异常结果列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
func (c *Client) DescribeNewL7RulesErrHealthWithContext(ctx context.Context, request *DescribeNewL7RulesErrHealthRequest) (response *DescribeNewL7RulesErrHealthResponse, err error) {
    if request == nil {
        request = NewDescribeNewL7RulesErrHealthRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeNewL7RulesErrHealth require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeNewL7RulesErrHealthResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeOverviewAttackTrendRequest() (request *DescribeOverviewAttackTrendRequest) {
    request = &DescribeOverviewAttackTrendRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("antiddos", APIVersion, "DescribeOverviewAttackTrend")
    
    
    return
}

func NewDescribeOverviewAttackTrendResponse() (response *DescribeOverviewAttackTrendResponse) {
    response = &DescribeOverviewAttackTrendResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeOverviewAttackTrend
// 拉取防护概览攻击趋势
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
func (c *Client) DescribeOverviewAttackTrend(request *DescribeOverviewAttackTrendRequest) (response *DescribeOverviewAttackTrendResponse, err error) {
    return c.DescribeOverviewAttackTrendWithContext(context.Background(), request)
}

// DescribeOverviewAttackTrend
// 拉取防护概览攻击趋势
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
func (c *Client) DescribeOverviewAttackTrendWithContext(ctx context.Context, request *DescribeOverviewAttackTrendRequest) (response *DescribeOverviewAttackTrendResponse, err error) {
    if request == nil {
        request = NewDescribeOverviewAttackTrendRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeOverviewAttackTrend require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeOverviewAttackTrendResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeOverviewCCTrendRequest() (request *DescribeOverviewCCTrendRequest) {
    request = &DescribeOverviewCCTrendRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("antiddos", APIVersion, "DescribeOverviewCCTrend")
    
    
    return
}

func NewDescribeOverviewCCTrendResponse() (response *DescribeOverviewCCTrendResponse) {
    response = &DescribeOverviewCCTrendResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeOverviewCCTrend
// 获取防护概览总请求峰值(QPS)和攻击请求(QPS)以及总请求次数和攻击请求次数
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
func (c *Client) DescribeOverviewCCTrend(request *DescribeOverviewCCTrendRequest) (response *DescribeOverviewCCTrendResponse, err error) {
    return c.DescribeOverviewCCTrendWithContext(context.Background(), request)
}

// DescribeOverviewCCTrend
// 获取防护概览总请求峰值(QPS)和攻击请求(QPS)以及总请求次数和攻击请求次数
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
func (c *Client) DescribeOverviewCCTrendWithContext(ctx context.Context, request *DescribeOverviewCCTrendRequest) (response *DescribeOverviewCCTrendResponse, err error) {
    if request == nil {
        request = NewDescribeOverviewCCTrendRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeOverviewCCTrend require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeOverviewCCTrendResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeOverviewDDoSEventListRequest() (request *DescribeOverviewDDoSEventListRequest) {
    request = &DescribeOverviewDDoSEventListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("antiddos", APIVersion, "DescribeOverviewDDoSEventList")
    
    
    return
}

func NewDescribeOverviewDDoSEventListResponse() (response *DescribeOverviewDDoSEventListResponse) {
    response = &DescribeOverviewDDoSEventListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeOverviewDDoSEventList
// 获取防护概览的ddos攻击事件
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
func (c *Client) DescribeOverviewDDoSEventList(request *DescribeOverviewDDoSEventListRequest) (response *DescribeOverviewDDoSEventListResponse, err error) {
    return c.DescribeOverviewDDoSEventListWithContext(context.Background(), request)
}

// DescribeOverviewDDoSEventList
// 获取防护概览的ddos攻击事件
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
func (c *Client) DescribeOverviewDDoSEventListWithContext(ctx context.Context, request *DescribeOverviewDDoSEventListRequest) (response *DescribeOverviewDDoSEventListResponse, err error) {
    if request == nil {
        request = NewDescribeOverviewDDoSEventListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeOverviewDDoSEventList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeOverviewDDoSEventListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeOverviewDDoSTrendRequest() (request *DescribeOverviewDDoSTrendRequest) {
    request = &DescribeOverviewDDoSTrendRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("antiddos", APIVersion, "DescribeOverviewDDoSTrend")
    
    
    return
}

func NewDescribeOverviewDDoSTrendResponse() (response *DescribeOverviewDDoSTrendResponse) {
    response = &DescribeOverviewDDoSTrendResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeOverviewDDoSTrend
// 获取防护概览DDoS攻击流量带宽和攻击包速率数据
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
func (c *Client) DescribeOverviewDDoSTrend(request *DescribeOverviewDDoSTrendRequest) (response *DescribeOverviewDDoSTrendResponse, err error) {
    return c.DescribeOverviewDDoSTrendWithContext(context.Background(), request)
}

// DescribeOverviewDDoSTrend
// 获取防护概览DDoS攻击流量带宽和攻击包速率数据
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
func (c *Client) DescribeOverviewDDoSTrendWithContext(ctx context.Context, request *DescribeOverviewDDoSTrendRequest) (response *DescribeOverviewDDoSTrendResponse, err error) {
    if request == nil {
        request = NewDescribeOverviewDDoSTrendRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeOverviewDDoSTrend require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeOverviewDDoSTrendResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeOverviewIndexRequest() (request *DescribeOverviewIndexRequest) {
    request = &DescribeOverviewIndexRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("antiddos", APIVersion, "DescribeOverviewIndex")
    
    
    return
}

func NewDescribeOverviewIndexResponse() (response *DescribeOverviewIndexResponse) {
    response = &DescribeOverviewIndexResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeOverviewIndex
// 拉取防护概览指标
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
func (c *Client) DescribeOverviewIndex(request *DescribeOverviewIndexRequest) (response *DescribeOverviewIndexResponse, err error) {
    return c.DescribeOverviewIndexWithContext(context.Background(), request)
}

// DescribeOverviewIndex
// 拉取防护概览指标
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
func (c *Client) DescribeOverviewIndexWithContext(ctx context.Context, request *DescribeOverviewIndexRequest) (response *DescribeOverviewIndexResponse, err error) {
    if request == nil {
        request = NewDescribeOverviewIndexRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeOverviewIndex require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeOverviewIndexResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePendingRiskInfoRequest() (request *DescribePendingRiskInfoRequest) {
    request = &DescribePendingRiskInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("antiddos", APIVersion, "DescribePendingRiskInfo")
    
    
    return
}

func NewDescribePendingRiskInfoResponse() (response *DescribePendingRiskInfoResponse) {
    response = &DescribePendingRiskInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribePendingRiskInfo
// 查询账号维度待处理风险信息，包括是否为付费用户，查询攻击中、封堵中、过期资源数量等
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
func (c *Client) DescribePendingRiskInfo(request *DescribePendingRiskInfoRequest) (response *DescribePendingRiskInfoResponse, err error) {
    return c.DescribePendingRiskInfoWithContext(context.Background(), request)
}

// DescribePendingRiskInfo
// 查询账号维度待处理风险信息，包括是否为付费用户，查询攻击中、封堵中、过期资源数量等
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
func (c *Client) DescribePendingRiskInfoWithContext(ctx context.Context, request *DescribePendingRiskInfoRequest) (response *DescribePendingRiskInfoResponse, err error) {
    if request == nil {
        request = NewDescribePendingRiskInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePendingRiskInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePendingRiskInfoResponse()
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
    return c.DisassociateDDoSEipAddressWithContext(context.Background(), request)
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DisassociateDDoSEipAddress require credential")
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
    return c.ModifyCCLevelPolicyWithContext(context.Background(), request)
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyCCLevelPolicy require credential")
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
    return c.ModifyCCPrecisionPolicyWithContext(context.Background(), request)
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyCCPrecisionPolicy require credential")
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
    return c.ModifyCCReqLimitPolicyWithContext(context.Background(), request)
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyCCReqLimitPolicy require credential")
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
    return c.ModifyCCThresholdPolicyWithContext(context.Background(), request)
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyCCThresholdPolicy require credential")
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
    return c.ModifyCcBlackWhiteIpListWithContext(context.Background(), request)
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyCcBlackWhiteIpList require credential")
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
    return c.ModifyDDoSBlackWhiteIpListWithContext(context.Background(), request)
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDDoSBlackWhiteIpList require credential")
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
    return c.ModifyDDoSGeoIPBlockConfigWithContext(context.Background(), request)
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDDoSGeoIPBlockConfig require credential")
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
    return c.ModifyDDoSLevelWithContext(context.Background(), request)
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDDoSLevel require credential")
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
    return c.ModifyDDoSSpeedLimitConfigWithContext(context.Background(), request)
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDDoSSpeedLimitConfig require credential")
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
    return c.ModifyDDoSThresholdWithContext(context.Background(), request)
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDDoSThreshold require credential")
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
    return c.ModifyDomainUsrNameWithContext(context.Background(), request)
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDomainUsrName require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDomainUsrNameResponse()
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
    return c.ModifyNewDomainRulesWithContext(context.Background(), request)
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyNewDomainRules require credential")
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
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyPacketFilterConfig(request *ModifyPacketFilterConfigRequest) (response *ModifyPacketFilterConfigResponse, err error) {
    return c.ModifyPacketFilterConfigWithContext(context.Background(), request)
}

// ModifyPacketFilterConfig
// 修改DDoS防护的特征过滤规则
//
// 可能返回的错误码:
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyPacketFilterConfigWithContext(ctx context.Context, request *ModifyPacketFilterConfigRequest) (response *ModifyPacketFilterConfigResponse, err error) {
    if request == nil {
        request = NewModifyPacketFilterConfigRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyPacketFilterConfig require credential")
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
    return c.ModifyPortAclConfigWithContext(context.Background(), request)
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyPortAclConfig require credential")
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
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) SwitchWaterPrintConfig(request *SwitchWaterPrintConfigRequest) (response *SwitchWaterPrintConfigResponse, err error) {
    return c.SwitchWaterPrintConfigWithContext(context.Background(), request)
}

// SwitchWaterPrintConfig
// 开启或关闭DDoS防护的水印防护配置
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) SwitchWaterPrintConfigWithContext(ctx context.Context, request *SwitchWaterPrintConfigRequest) (response *SwitchWaterPrintConfigResponse, err error) {
    if request == nil {
        request = NewSwitchWaterPrintConfigRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SwitchWaterPrintConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewSwitchWaterPrintConfigResponse()
    err = c.Send(request, response)
    return
}
