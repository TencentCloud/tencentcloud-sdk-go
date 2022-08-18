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

package v20180625

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2018-06-25"

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


func NewAcceptVpcPeerConnectionRequest() (request *AcceptVpcPeerConnectionRequest) {
    request = &AcceptVpcPeerConnectionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bmvpc", APIVersion, "AcceptVpcPeerConnection")
    
    
    return
}

func NewAcceptVpcPeerConnectionResponse() (response *AcceptVpcPeerConnectionResponse) {
    response = &AcceptVpcPeerConnectionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AcceptVpcPeerConnection
// 接受黑石对等连接
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_VPCNOTEXIST = "ResourceNotFound.VpcNotExist"
//  RESOURCENOTFOUND_VPCPEERNOTEXIST = "ResourceNotFound.VpcPeerNotExist"
//  UNSUPPORTEDOPERATION_INVALIDVPCPEERSTATE = "UnsupportedOperation.InvalidVpcPeerState"
//  UNSUPPORTEDOPERATION_VPCCIDRCONFICT = "UnsupportedOperation.VpcCidrConfict"
//  UNSUPPORTEDOPERATION_VPCPEEREXIST = "UnsupportedOperation.VpcPeerExist"
func (c *Client) AcceptVpcPeerConnection(request *AcceptVpcPeerConnectionRequest) (response *AcceptVpcPeerConnectionResponse, err error) {
    return c.AcceptVpcPeerConnectionWithContext(context.Background(), request)
}

// AcceptVpcPeerConnection
// 接受黑石对等连接
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_VPCNOTEXIST = "ResourceNotFound.VpcNotExist"
//  RESOURCENOTFOUND_VPCPEERNOTEXIST = "ResourceNotFound.VpcPeerNotExist"
//  UNSUPPORTEDOPERATION_INVALIDVPCPEERSTATE = "UnsupportedOperation.InvalidVpcPeerState"
//  UNSUPPORTEDOPERATION_VPCCIDRCONFICT = "UnsupportedOperation.VpcCidrConfict"
//  UNSUPPORTEDOPERATION_VPCPEEREXIST = "UnsupportedOperation.VpcPeerExist"
func (c *Client) AcceptVpcPeerConnectionWithContext(ctx context.Context, request *AcceptVpcPeerConnectionRequest) (response *AcceptVpcPeerConnectionResponse, err error) {
    if request == nil {
        request = NewAcceptVpcPeerConnectionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AcceptVpcPeerConnection require credential")
    }

    request.SetContext(ctx)
    
    response = NewAcceptVpcPeerConnectionResponse()
    err = c.Send(request, response)
    return
}

func NewAsyncRegisterIpsRequest() (request *AsyncRegisterIpsRequest) {
    request = &AsyncRegisterIpsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bmvpc", APIVersion, "AsyncRegisterIps")
    
    
    return
}

func NewAsyncRegisterIpsResponse() (response *AsyncRegisterIpsResponse) {
    response = &AsyncRegisterIpsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AsyncRegisterIps
// 批量注册虚拟IP，异步接口。通过接口来查询任务进度。每次请求最多注册256个IP
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) AsyncRegisterIps(request *AsyncRegisterIpsRequest) (response *AsyncRegisterIpsResponse, err error) {
    return c.AsyncRegisterIpsWithContext(context.Background(), request)
}

// AsyncRegisterIps
// 批量注册虚拟IP，异步接口。通过接口来查询任务进度。每次请求最多注册256个IP
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) AsyncRegisterIpsWithContext(ctx context.Context, request *AsyncRegisterIpsRequest) (response *AsyncRegisterIpsResponse, err error) {
    if request == nil {
        request = NewAsyncRegisterIpsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AsyncRegisterIps require credential")
    }

    request.SetContext(ctx)
    
    response = NewAsyncRegisterIpsResponse()
    err = c.Send(request, response)
    return
}

func NewBindEipsToNatGatewayRequest() (request *BindEipsToNatGatewayRequest) {
    request = &BindEipsToNatGatewayRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bmvpc", APIVersion, "BindEipsToNatGateway")
    
    
    return
}

func NewBindEipsToNatGatewayResponse() (response *BindEipsToNatGatewayResponse) {
    response = &BindEipsToNatGatewayResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// BindEipsToNatGateway
// NAT网关绑定EIP接口，可将EIP绑定到NAT网关，该EIP作为访问外网的源IP地址，将流量发送到Internet
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) BindEipsToNatGateway(request *BindEipsToNatGatewayRequest) (response *BindEipsToNatGatewayResponse, err error) {
    return c.BindEipsToNatGatewayWithContext(context.Background(), request)
}

// BindEipsToNatGateway
// NAT网关绑定EIP接口，可将EIP绑定到NAT网关，该EIP作为访问外网的源IP地址，将流量发送到Internet
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) BindEipsToNatGatewayWithContext(ctx context.Context, request *BindEipsToNatGatewayRequest) (response *BindEipsToNatGatewayResponse, err error) {
    if request == nil {
        request = NewBindEipsToNatGatewayRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("BindEipsToNatGateway require credential")
    }

    request.SetContext(ctx)
    
    response = NewBindEipsToNatGatewayResponse()
    err = c.Send(request, response)
    return
}

func NewBindIpsToNatGatewayRequest() (request *BindIpsToNatGatewayRequest) {
    request = &BindIpsToNatGatewayRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bmvpc", APIVersion, "BindIpsToNatGateway")
    
    
    return
}

func NewBindIpsToNatGatewayResponse() (response *BindIpsToNatGatewayResponse) {
    response = &BindIpsToNatGatewayResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// BindIpsToNatGateway
// 可用于将子网的部分IP绑定到NAT网关
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) BindIpsToNatGateway(request *BindIpsToNatGatewayRequest) (response *BindIpsToNatGatewayResponse, err error) {
    return c.BindIpsToNatGatewayWithContext(context.Background(), request)
}

// BindIpsToNatGateway
// 可用于将子网的部分IP绑定到NAT网关
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) BindIpsToNatGatewayWithContext(ctx context.Context, request *BindIpsToNatGatewayRequest) (response *BindIpsToNatGatewayResponse, err error) {
    if request == nil {
        request = NewBindIpsToNatGatewayRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("BindIpsToNatGateway require credential")
    }

    request.SetContext(ctx)
    
    response = NewBindIpsToNatGatewayResponse()
    err = c.Send(request, response)
    return
}

func NewBindSubnetsToNatGatewayRequest() (request *BindSubnetsToNatGatewayRequest) {
    request = &BindSubnetsToNatGatewayRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bmvpc", APIVersion, "BindSubnetsToNatGateway")
    
    
    return
}

func NewBindSubnetsToNatGatewayResponse() (response *BindSubnetsToNatGatewayResponse) {
    response = &BindSubnetsToNatGatewayResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// BindSubnetsToNatGateway
// NAT网关绑定子网后，该子网内全部IP可出公网
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) BindSubnetsToNatGateway(request *BindSubnetsToNatGatewayRequest) (response *BindSubnetsToNatGatewayResponse, err error) {
    return c.BindSubnetsToNatGatewayWithContext(context.Background(), request)
}

// BindSubnetsToNatGateway
// NAT网关绑定子网后，该子网内全部IP可出公网
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) BindSubnetsToNatGatewayWithContext(ctx context.Context, request *BindSubnetsToNatGatewayRequest) (response *BindSubnetsToNatGatewayResponse, err error) {
    if request == nil {
        request = NewBindSubnetsToNatGatewayRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("BindSubnetsToNatGateway require credential")
    }

    request.SetContext(ctx)
    
    response = NewBindSubnetsToNatGatewayResponse()
    err = c.Send(request, response)
    return
}

func NewCreateCustomerGatewayRequest() (request *CreateCustomerGatewayRequest) {
    request = &CreateCustomerGatewayRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bmvpc", APIVersion, "CreateCustomerGateway")
    
    
    return
}

func NewCreateCustomerGatewayResponse() (response *CreateCustomerGatewayResponse) {
    response = &CreateCustomerGatewayResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateCustomerGateway
// 本接口（CreateCustomerGateway）用于创建对端网关。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_CUSTOMERGATEWAYNOTEXIST = "ResourceNotFound.CustomerGatewayNotExist"
//  UNSUPPORTEDOPERATION_CUSTOMERGATEWAYADDREXIST = "UnsupportedOperation.CustomerGatewayAddrExist"
//  UNSUPPORTEDOPERATION_CUSTOMERGATEWAYADDRINVALID = "UnsupportedOperation.CustomerGatewayAddrInvalid"
//  UNSUPPORTEDOPERATION_VPNCONNINUSE = "UnsupportedOperation.VpnConnInUse"
func (c *Client) CreateCustomerGateway(request *CreateCustomerGatewayRequest) (response *CreateCustomerGatewayResponse, err error) {
    return c.CreateCustomerGatewayWithContext(context.Background(), request)
}

// CreateCustomerGateway
// 本接口（CreateCustomerGateway）用于创建对端网关。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_CUSTOMERGATEWAYNOTEXIST = "ResourceNotFound.CustomerGatewayNotExist"
//  UNSUPPORTEDOPERATION_CUSTOMERGATEWAYADDREXIST = "UnsupportedOperation.CustomerGatewayAddrExist"
//  UNSUPPORTEDOPERATION_CUSTOMERGATEWAYADDRINVALID = "UnsupportedOperation.CustomerGatewayAddrInvalid"
//  UNSUPPORTEDOPERATION_VPNCONNINUSE = "UnsupportedOperation.VpnConnInUse"
func (c *Client) CreateCustomerGatewayWithContext(ctx context.Context, request *CreateCustomerGatewayRequest) (response *CreateCustomerGatewayResponse, err error) {
    if request == nil {
        request = NewCreateCustomerGatewayRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCustomerGateway require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateCustomerGatewayResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDockerSubnetWithVlanRequest() (request *CreateDockerSubnetWithVlanRequest) {
    request = &CreateDockerSubnetWithVlanRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bmvpc", APIVersion, "CreateDockerSubnetWithVlan")
    
    
    return
}

func NewCreateDockerSubnetWithVlanResponse() (response *CreateDockerSubnetWithVlanResponse) {
    response = &CreateDockerSubnetWithVlanResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateDockerSubnetWithVlan
// 创建黑石Docker子网， 如果不指定VlanId，将会分配2000--2999范围的VlanId; 子网会关闭分布式网关
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateDockerSubnetWithVlan(request *CreateDockerSubnetWithVlanRequest) (response *CreateDockerSubnetWithVlanResponse, err error) {
    return c.CreateDockerSubnetWithVlanWithContext(context.Background(), request)
}

// CreateDockerSubnetWithVlan
// 创建黑石Docker子网， 如果不指定VlanId，将会分配2000--2999范围的VlanId; 子网会关闭分布式网关
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateDockerSubnetWithVlanWithContext(ctx context.Context, request *CreateDockerSubnetWithVlanRequest) (response *CreateDockerSubnetWithVlanResponse, err error) {
    if request == nil {
        request = NewCreateDockerSubnetWithVlanRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateDockerSubnetWithVlan require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateDockerSubnetWithVlanResponse()
    err = c.Send(request, response)
    return
}

func NewCreateHostedInterfaceRequest() (request *CreateHostedInterfaceRequest) {
    request = &CreateHostedInterfaceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bmvpc", APIVersion, "CreateHostedInterface")
    
    
    return
}

func NewCreateHostedInterfaceResponse() (response *CreateHostedInterfaceResponse) {
    response = &CreateHostedInterfaceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateHostedInterface
// 本接口（CreateHostedInterface）用于黑石托管机器加入带VLANID不为5的子网。
//
// 
//
// 1) 不能加入vlanId 为5的子网，只能加入VLANID范围为2000-2999的子网。
//
// 2) 每台托管机器最多可以加入20个子网。
//
// 3) 每次调用最多能支持传入10台托管机器。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateHostedInterface(request *CreateHostedInterfaceRequest) (response *CreateHostedInterfaceResponse, err error) {
    return c.CreateHostedInterfaceWithContext(context.Background(), request)
}

// CreateHostedInterface
// 本接口（CreateHostedInterface）用于黑石托管机器加入带VLANID不为5的子网。
//
// 
//
// 1) 不能加入vlanId 为5的子网，只能加入VLANID范围为2000-2999的子网。
//
// 2) 每台托管机器最多可以加入20个子网。
//
// 3) 每次调用最多能支持传入10台托管机器。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateHostedInterfaceWithContext(ctx context.Context, request *CreateHostedInterfaceRequest) (response *CreateHostedInterfaceResponse, err error) {
    if request == nil {
        request = NewCreateHostedInterfaceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateHostedInterface require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateHostedInterfaceResponse()
    err = c.Send(request, response)
    return
}

func NewCreateInterfacesRequest() (request *CreateInterfacesRequest) {
    request = &CreateInterfacesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bmvpc", APIVersion, "CreateInterfaces")
    
    
    return
}

func NewCreateInterfacesResponse() (response *CreateInterfacesResponse) {
    response = &CreateInterfacesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateInterfaces
// 物理机加入子网
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateInterfaces(request *CreateInterfacesRequest) (response *CreateInterfacesResponse, err error) {
    return c.CreateInterfacesWithContext(context.Background(), request)
}

// CreateInterfaces
// 物理机加入子网
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateInterfacesWithContext(ctx context.Context, request *CreateInterfacesRequest) (response *CreateInterfacesResponse, err error) {
    if request == nil {
        request = NewCreateInterfacesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateInterfaces require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateInterfacesResponse()
    err = c.Send(request, response)
    return
}

func NewCreateNatGatewayRequest() (request *CreateNatGatewayRequest) {
    request = &CreateNatGatewayRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bmvpc", APIVersion, "CreateNatGateway")
    
    
    return
}

func NewCreateNatGatewayResponse() (response *CreateNatGatewayResponse) {
    response = &CreateNatGatewayResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateNatGateway
// 创建NAT网关接口，可针对网段方式、子网全部IP、子网部分IP这三种方式创建NAT网关
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateNatGateway(request *CreateNatGatewayRequest) (response *CreateNatGatewayResponse, err error) {
    return c.CreateNatGatewayWithContext(context.Background(), request)
}

// CreateNatGateway
// 创建NAT网关接口，可针对网段方式、子网全部IP、子网部分IP这三种方式创建NAT网关
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateNatGatewayWithContext(ctx context.Context, request *CreateNatGatewayRequest) (response *CreateNatGatewayResponse, err error) {
    if request == nil {
        request = NewCreateNatGatewayRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateNatGateway require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateNatGatewayResponse()
    err = c.Send(request, response)
    return
}

func NewCreateRoutePoliciesRequest() (request *CreateRoutePoliciesRequest) {
    request = &CreateRoutePoliciesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bmvpc", APIVersion, "CreateRoutePolicies")
    
    
    return
}

func NewCreateRoutePoliciesResponse() (response *CreateRoutePoliciesResponse) {
    response = &CreateRoutePoliciesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateRoutePolicies
// 创建黑石路由表的路由规则
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateRoutePolicies(request *CreateRoutePoliciesRequest) (response *CreateRoutePoliciesResponse, err error) {
    return c.CreateRoutePoliciesWithContext(context.Background(), request)
}

// CreateRoutePolicies
// 创建黑石路由表的路由规则
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateRoutePoliciesWithContext(ctx context.Context, request *CreateRoutePoliciesRequest) (response *CreateRoutePoliciesResponse, err error) {
    if request == nil {
        request = NewCreateRoutePoliciesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateRoutePolicies require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateRoutePoliciesResponse()
    err = c.Send(request, response)
    return
}

func NewCreateSubnetRequest() (request *CreateSubnetRequest) {
    request = &CreateSubnetRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bmvpc", APIVersion, "CreateSubnet")
    
    
    return
}

func NewCreateSubnetResponse() (response *CreateSubnetResponse) {
    response = &CreateSubnetResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateSubnet
// 创建黑石私有网络的子网
//
// 访问管理: 用户可以对VpcId进行授权操作。例如设置资源为["qcs::bmvpc:::unVpc/vpc-xxxxx"]
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateSubnet(request *CreateSubnetRequest) (response *CreateSubnetResponse, err error) {
    return c.CreateSubnetWithContext(context.Background(), request)
}

// CreateSubnet
// 创建黑石私有网络的子网
//
// 访问管理: 用户可以对VpcId进行授权操作。例如设置资源为["qcs::bmvpc:::unVpc/vpc-xxxxx"]
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateSubnetWithContext(ctx context.Context, request *CreateSubnetRequest) (response *CreateSubnetResponse, err error) {
    if request == nil {
        request = NewCreateSubnetRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateSubnet require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateSubnetResponse()
    err = c.Send(request, response)
    return
}

func NewCreateVirtualSubnetWithVlanRequest() (request *CreateVirtualSubnetWithVlanRequest) {
    request = &CreateVirtualSubnetWithVlanRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bmvpc", APIVersion, "CreateVirtualSubnetWithVlan")
    
    
    return
}

func NewCreateVirtualSubnetWithVlanResponse() (response *CreateVirtualSubnetWithVlanResponse) {
    response = &CreateVirtualSubnetWithVlanResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateVirtualSubnetWithVlan
// 创建黑石虚拟子网， 虚拟子网用于在黑石上创建虚拟网络，与黑石子网要做好规划。虚拟子网会分配2000-2999的VlanId。
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateVirtualSubnetWithVlan(request *CreateVirtualSubnetWithVlanRequest) (response *CreateVirtualSubnetWithVlanResponse, err error) {
    return c.CreateVirtualSubnetWithVlanWithContext(context.Background(), request)
}

// CreateVirtualSubnetWithVlan
// 创建黑石虚拟子网， 虚拟子网用于在黑石上创建虚拟网络，与黑石子网要做好规划。虚拟子网会分配2000-2999的VlanId。
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateVirtualSubnetWithVlanWithContext(ctx context.Context, request *CreateVirtualSubnetWithVlanRequest) (response *CreateVirtualSubnetWithVlanResponse, err error) {
    if request == nil {
        request = NewCreateVirtualSubnetWithVlanRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateVirtualSubnetWithVlan require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateVirtualSubnetWithVlanResponse()
    err = c.Send(request, response)
    return
}

func NewCreateVpcRequest() (request *CreateVpcRequest) {
    request = &CreateVpcRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bmvpc", APIVersion, "CreateVpc")
    
    
    return
}

func NewCreateVpcResponse() (response *CreateVpcResponse) {
    response = &CreateVpcResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateVpc
// 创建黑石私有网络
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
func (c *Client) CreateVpc(request *CreateVpcRequest) (response *CreateVpcResponse, err error) {
    return c.CreateVpcWithContext(context.Background(), request)
}

// CreateVpc
// 创建黑石私有网络
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
func (c *Client) CreateVpcWithContext(ctx context.Context, request *CreateVpcRequest) (response *CreateVpcResponse, err error) {
    if request == nil {
        request = NewCreateVpcRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateVpc require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateVpcResponse()
    err = c.Send(request, response)
    return
}

func NewCreateVpcPeerConnectionRequest() (request *CreateVpcPeerConnectionRequest) {
    request = &CreateVpcPeerConnectionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bmvpc", APIVersion, "CreateVpcPeerConnection")
    
    
    return
}

func NewCreateVpcPeerConnectionResponse() (response *CreateVpcPeerConnectionResponse) {
    response = &CreateVpcPeerConnectionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateVpcPeerConnection
// 创建对等连接
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND_VPCNOTEXIST = "ResourceNotFound.VpcNotExist"
//  RESOURCENOTFOUND_VPCPEERNOTEXIST = "ResourceNotFound.VpcPeerNotExist"
//  UNSUPPORTEDOPERATION_INVALIDVPCPEERSTATE = "UnsupportedOperation.InvalidVpcPeerState"
//  UNSUPPORTEDOPERATION_VPCCIDRCONFICT = "UnsupportedOperation.VpcCidrConfict"
//  UNSUPPORTEDOPERATION_VPCPEEREXIST = "UnsupportedOperation.VpcPeerExist"
func (c *Client) CreateVpcPeerConnection(request *CreateVpcPeerConnectionRequest) (response *CreateVpcPeerConnectionResponse, err error) {
    return c.CreateVpcPeerConnectionWithContext(context.Background(), request)
}

// CreateVpcPeerConnection
// 创建对等连接
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND_VPCNOTEXIST = "ResourceNotFound.VpcNotExist"
//  RESOURCENOTFOUND_VPCPEERNOTEXIST = "ResourceNotFound.VpcPeerNotExist"
//  UNSUPPORTEDOPERATION_INVALIDVPCPEERSTATE = "UnsupportedOperation.InvalidVpcPeerState"
//  UNSUPPORTEDOPERATION_VPCCIDRCONFICT = "UnsupportedOperation.VpcCidrConfict"
//  UNSUPPORTEDOPERATION_VPCPEEREXIST = "UnsupportedOperation.VpcPeerExist"
func (c *Client) CreateVpcPeerConnectionWithContext(ctx context.Context, request *CreateVpcPeerConnectionRequest) (response *CreateVpcPeerConnectionResponse, err error) {
    if request == nil {
        request = NewCreateVpcPeerConnectionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateVpcPeerConnection require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateVpcPeerConnectionResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteCustomerGatewayRequest() (request *DeleteCustomerGatewayRequest) {
    request = &DeleteCustomerGatewayRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bmvpc", APIVersion, "DeleteCustomerGateway")
    
    
    return
}

func NewDeleteCustomerGatewayResponse() (response *DeleteCustomerGatewayResponse) {
    response = &DeleteCustomerGatewayResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteCustomerGateway
// 本接口（DeleteCustomerGateway）用于删除对端网关。
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_CUSTOMERGATEWAYNOTEXIST = "ResourceNotFound.CustomerGatewayNotExist"
//  UNSUPPORTEDOPERATION_CUSTOMERGATEWAYADDREXIST = "UnsupportedOperation.CustomerGatewayAddrExist"
//  UNSUPPORTEDOPERATION_CUSTOMERGATEWAYADDRINVALID = "UnsupportedOperation.CustomerGatewayAddrInvalid"
//  UNSUPPORTEDOPERATION_VPNCONNINUSE = "UnsupportedOperation.VpnConnInUse"
func (c *Client) DeleteCustomerGateway(request *DeleteCustomerGatewayRequest) (response *DeleteCustomerGatewayResponse, err error) {
    return c.DeleteCustomerGatewayWithContext(context.Background(), request)
}

// DeleteCustomerGateway
// 本接口（DeleteCustomerGateway）用于删除对端网关。
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_CUSTOMERGATEWAYNOTEXIST = "ResourceNotFound.CustomerGatewayNotExist"
//  UNSUPPORTEDOPERATION_CUSTOMERGATEWAYADDREXIST = "UnsupportedOperation.CustomerGatewayAddrExist"
//  UNSUPPORTEDOPERATION_CUSTOMERGATEWAYADDRINVALID = "UnsupportedOperation.CustomerGatewayAddrInvalid"
//  UNSUPPORTEDOPERATION_VPNCONNINUSE = "UnsupportedOperation.VpnConnInUse"
func (c *Client) DeleteCustomerGatewayWithContext(ctx context.Context, request *DeleteCustomerGatewayRequest) (response *DeleteCustomerGatewayResponse, err error) {
    if request == nil {
        request = NewDeleteCustomerGatewayRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteCustomerGateway require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteCustomerGatewayResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteHostedInterfaceRequest() (request *DeleteHostedInterfaceRequest) {
    request = &DeleteHostedInterfaceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bmvpc", APIVersion, "DeleteHostedInterface")
    
    
    return
}

func NewDeleteHostedInterfaceResponse() (response *DeleteHostedInterfaceResponse) {
    response = &DeleteHostedInterfaceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteHostedInterface
// 本接口用于托管机器从VLANID不为5的子网中移除。
//
// 1) 不能从vlanId 为5的子网中移除。
//
// 2) 每次调用最多能支持传入10台物理机。
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteHostedInterface(request *DeleteHostedInterfaceRequest) (response *DeleteHostedInterfaceResponse, err error) {
    return c.DeleteHostedInterfaceWithContext(context.Background(), request)
}

// DeleteHostedInterface
// 本接口用于托管机器从VLANID不为5的子网中移除。
//
// 1) 不能从vlanId 为5的子网中移除。
//
// 2) 每次调用最多能支持传入10台物理机。
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteHostedInterfaceWithContext(ctx context.Context, request *DeleteHostedInterfaceRequest) (response *DeleteHostedInterfaceResponse, err error) {
    if request == nil {
        request = NewDeleteHostedInterfaceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteHostedInterface require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteHostedInterfaceResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteHostedInterfacesRequest() (request *DeleteHostedInterfacesRequest) {
    request = &DeleteHostedInterfacesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bmvpc", APIVersion, "DeleteHostedInterfaces")
    
    
    return
}

func NewDeleteHostedInterfacesResponse() (response *DeleteHostedInterfacesResponse) {
    response = &DeleteHostedInterfacesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteHostedInterfaces
// 托管机器移除子网批量接口，传入一台托管机器和多个子网，批量移除这些子网。异步接口，接口返回TaskId。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteHostedInterfaces(request *DeleteHostedInterfacesRequest) (response *DeleteHostedInterfacesResponse, err error) {
    return c.DeleteHostedInterfacesWithContext(context.Background(), request)
}

// DeleteHostedInterfaces
// 托管机器移除子网批量接口，传入一台托管机器和多个子网，批量移除这些子网。异步接口，接口返回TaskId。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteHostedInterfacesWithContext(ctx context.Context, request *DeleteHostedInterfacesRequest) (response *DeleteHostedInterfacesResponse, err error) {
    if request == nil {
        request = NewDeleteHostedInterfacesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteHostedInterfaces require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteHostedInterfacesResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteInterfacesRequest() (request *DeleteInterfacesRequest) {
    request = &DeleteInterfacesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bmvpc", APIVersion, "DeleteInterfaces")
    
    
    return
}

func NewDeleteInterfacesResponse() (response *DeleteInterfacesResponse) {
    response = &DeleteInterfacesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteInterfaces
// 物理机移除子网批量接口，传入一台物理机和多个子网，批量移除这些子网。异步接口，接口返回TaskId。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteInterfaces(request *DeleteInterfacesRequest) (response *DeleteInterfacesResponse, err error) {
    return c.DeleteInterfacesWithContext(context.Background(), request)
}

// DeleteInterfaces
// 物理机移除子网批量接口，传入一台物理机和多个子网，批量移除这些子网。异步接口，接口返回TaskId。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteInterfacesWithContext(ctx context.Context, request *DeleteInterfacesRequest) (response *DeleteInterfacesResponse, err error) {
    if request == nil {
        request = NewDeleteInterfacesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteInterfaces require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteInterfacesResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteNatGatewayRequest() (request *DeleteNatGatewayRequest) {
    request = &DeleteNatGatewayRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bmvpc", APIVersion, "DeleteNatGateway")
    
    
    return
}

func NewDeleteNatGatewayResponse() (response *DeleteNatGatewayResponse) {
    response = &DeleteNatGatewayResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteNatGateway
// 删除NAT网关
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteNatGateway(request *DeleteNatGatewayRequest) (response *DeleteNatGatewayResponse, err error) {
    return c.DeleteNatGatewayWithContext(context.Background(), request)
}

// DeleteNatGateway
// 删除NAT网关
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteNatGatewayWithContext(ctx context.Context, request *DeleteNatGatewayRequest) (response *DeleteNatGatewayResponse, err error) {
    if request == nil {
        request = NewDeleteNatGatewayRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteNatGateway require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteNatGatewayResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteRoutePolicyRequest() (request *DeleteRoutePolicyRequest) {
    request = &DeleteRoutePolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bmvpc", APIVersion, "DeleteRoutePolicy")
    
    
    return
}

func NewDeleteRoutePolicyResponse() (response *DeleteRoutePolicyResponse) {
    response = &DeleteRoutePolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteRoutePolicy
// 删除黑石路由表路由规则
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteRoutePolicy(request *DeleteRoutePolicyRequest) (response *DeleteRoutePolicyResponse, err error) {
    return c.DeleteRoutePolicyWithContext(context.Background(), request)
}

// DeleteRoutePolicy
// 删除黑石路由表路由规则
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteRoutePolicyWithContext(ctx context.Context, request *DeleteRoutePolicyRequest) (response *DeleteRoutePolicyResponse, err error) {
    if request == nil {
        request = NewDeleteRoutePolicyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteRoutePolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteRoutePolicyResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteSubnetRequest() (request *DeleteSubnetRequest) {
    request = &DeleteSubnetRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bmvpc", APIVersion, "DeleteSubnet")
    
    
    return
}

func NewDeleteSubnetResponse() (response *DeleteSubnetResponse) {
    response = &DeleteSubnetResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteSubnet
// 本接口（DeleteSubnet）用于删除黑石私有网络子网。
//
// 删除子网前，请清理该子网下所有资源，包括物理机、负载均衡、黑石数据库、弹性IP、NAT网关等资源
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeleteSubnet(request *DeleteSubnetRequest) (response *DeleteSubnetResponse, err error) {
    return c.DeleteSubnetWithContext(context.Background(), request)
}

// DeleteSubnet
// 本接口（DeleteSubnet）用于删除黑石私有网络子网。
//
// 删除子网前，请清理该子网下所有资源，包括物理机、负载均衡、黑石数据库、弹性IP、NAT网关等资源
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeleteSubnetWithContext(ctx context.Context, request *DeleteSubnetRequest) (response *DeleteSubnetResponse, err error) {
    if request == nil {
        request = NewDeleteSubnetRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteSubnet require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteSubnetResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteVirtualIpRequest() (request *DeleteVirtualIpRequest) {
    request = &DeleteVirtualIpRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bmvpc", APIVersion, "DeleteVirtualIp")
    
    
    return
}

func NewDeleteVirtualIpResponse() (response *DeleteVirtualIpResponse) {
    response = &DeleteVirtualIpResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteVirtualIp
// 退还虚拟IP。此接口只能退还虚拟IP，物理机IP不能退还。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteVirtualIp(request *DeleteVirtualIpRequest) (response *DeleteVirtualIpResponse, err error) {
    return c.DeleteVirtualIpWithContext(context.Background(), request)
}

// DeleteVirtualIp
// 退还虚拟IP。此接口只能退还虚拟IP，物理机IP不能退还。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteVirtualIpWithContext(ctx context.Context, request *DeleteVirtualIpRequest) (response *DeleteVirtualIpResponse, err error) {
    if request == nil {
        request = NewDeleteVirtualIpRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteVirtualIp require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteVirtualIpResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteVpcRequest() (request *DeleteVpcRequest) {
    request = &DeleteVpcRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bmvpc", APIVersion, "DeleteVpc")
    
    
    return
}

func NewDeleteVpcResponse() (response *DeleteVpcResponse) {
    response = &DeleteVpcResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteVpc
// 本接口(DeleteVpc)用于删除黑石私有网络(VPC)。
//
// 
//
// 删除私有网络前，请清理该私有网络下所有资源，包括子网、负载均衡、弹性 IP、对等连接、NAT 网关、专线通道、SSLVPN 等资源。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteVpc(request *DeleteVpcRequest) (response *DeleteVpcResponse, err error) {
    return c.DeleteVpcWithContext(context.Background(), request)
}

// DeleteVpc
// 本接口(DeleteVpc)用于删除黑石私有网络(VPC)。
//
// 
//
// 删除私有网络前，请清理该私有网络下所有资源，包括子网、负载均衡、弹性 IP、对等连接、NAT 网关、专线通道、SSLVPN 等资源。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteVpcWithContext(ctx context.Context, request *DeleteVpcRequest) (response *DeleteVpcResponse, err error) {
    if request == nil {
        request = NewDeleteVpcRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteVpc require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteVpcResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteVpcPeerConnectionRequest() (request *DeleteVpcPeerConnectionRequest) {
    request = &DeleteVpcPeerConnectionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bmvpc", APIVersion, "DeleteVpcPeerConnection")
    
    
    return
}

func NewDeleteVpcPeerConnectionResponse() (response *DeleteVpcPeerConnectionResponse) {
    response = &DeleteVpcPeerConnectionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteVpcPeerConnection
// 删除黑石对等连接
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_VPCNOTEXIST = "ResourceNotFound.VpcNotExist"
//  RESOURCENOTFOUND_VPCPEERNOTEXIST = "ResourceNotFound.VpcPeerNotExist"
//  UNSUPPORTEDOPERATION_INVALIDVPCPEERSTATE = "UnsupportedOperation.InvalidVpcPeerState"
//  UNSUPPORTEDOPERATION_VPCCIDRCONFICT = "UnsupportedOperation.VpcCidrConfict"
//  UNSUPPORTEDOPERATION_VPCPEEREXIST = "UnsupportedOperation.VpcPeerExist"
func (c *Client) DeleteVpcPeerConnection(request *DeleteVpcPeerConnectionRequest) (response *DeleteVpcPeerConnectionResponse, err error) {
    return c.DeleteVpcPeerConnectionWithContext(context.Background(), request)
}

// DeleteVpcPeerConnection
// 删除黑石对等连接
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_VPCNOTEXIST = "ResourceNotFound.VpcNotExist"
//  RESOURCENOTFOUND_VPCPEERNOTEXIST = "ResourceNotFound.VpcPeerNotExist"
//  UNSUPPORTEDOPERATION_INVALIDVPCPEERSTATE = "UnsupportedOperation.InvalidVpcPeerState"
//  UNSUPPORTEDOPERATION_VPCCIDRCONFICT = "UnsupportedOperation.VpcCidrConfict"
//  UNSUPPORTEDOPERATION_VPCPEEREXIST = "UnsupportedOperation.VpcPeerExist"
func (c *Client) DeleteVpcPeerConnectionWithContext(ctx context.Context, request *DeleteVpcPeerConnectionRequest) (response *DeleteVpcPeerConnectionResponse, err error) {
    if request == nil {
        request = NewDeleteVpcPeerConnectionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteVpcPeerConnection require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteVpcPeerConnectionResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteVpnConnectionRequest() (request *DeleteVpnConnectionRequest) {
    request = &DeleteVpnConnectionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bmvpc", APIVersion, "DeleteVpnConnection")
    
    
    return
}

func NewDeleteVpnConnectionResponse() (response *DeleteVpnConnectionResponse) {
    response = &DeleteVpnConnectionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteVpnConnection
// 本接口(DeleteVpnConnection)用于删除VPN通道。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED_SPDACLLIMIT = "LimitExceeded.SpdAclLimit"
//  LIMITEXCEEDED_SPDDNETLIMIT = "LimitExceeded.SpdDnetLimit"
//  LIMITEXCEEDED_SPDSNETLIMIT = "LimitExceeded.SpdSnetLimit"
//  RESOURCENOTFOUND_NOTAVAIBLE = "ResourceNotFound.NotAvaible"
//  RESOURCENOTFOUND_VPNCONNNOTEXIST = "ResourceNotFound.VpnConnNotExist"
//  UNSUPPORTEDOPERATION_INVALIDVPNCONNSTATE = "UnsupportedOperation.InvalidVpnConnState"
//  UNSUPPORTEDOPERATION_SPDACLCIDRINVALID = "UnsupportedOperation.SpdAclCidrInvalid"
//  UNSUPPORTEDOPERATION_SPDSNETNOTINCIDR = "UnsupportedOperation.SpdSnetNotInCidr"
//  UNSUPPORTEDOPERATION_VPNCONNEXIST = "UnsupportedOperation.VpnConnExist"
func (c *Client) DeleteVpnConnection(request *DeleteVpnConnectionRequest) (response *DeleteVpnConnectionResponse, err error) {
    return c.DeleteVpnConnectionWithContext(context.Background(), request)
}

// DeleteVpnConnection
// 本接口(DeleteVpnConnection)用于删除VPN通道。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED_SPDACLLIMIT = "LimitExceeded.SpdAclLimit"
//  LIMITEXCEEDED_SPDDNETLIMIT = "LimitExceeded.SpdDnetLimit"
//  LIMITEXCEEDED_SPDSNETLIMIT = "LimitExceeded.SpdSnetLimit"
//  RESOURCENOTFOUND_NOTAVAIBLE = "ResourceNotFound.NotAvaible"
//  RESOURCENOTFOUND_VPNCONNNOTEXIST = "ResourceNotFound.VpnConnNotExist"
//  UNSUPPORTEDOPERATION_INVALIDVPNCONNSTATE = "UnsupportedOperation.InvalidVpnConnState"
//  UNSUPPORTEDOPERATION_SPDACLCIDRINVALID = "UnsupportedOperation.SpdAclCidrInvalid"
//  UNSUPPORTEDOPERATION_SPDSNETNOTINCIDR = "UnsupportedOperation.SpdSnetNotInCidr"
//  UNSUPPORTEDOPERATION_VPNCONNEXIST = "UnsupportedOperation.VpnConnExist"
func (c *Client) DeleteVpnConnectionWithContext(ctx context.Context, request *DeleteVpnConnectionRequest) (response *DeleteVpnConnectionResponse, err error) {
    if request == nil {
        request = NewDeleteVpnConnectionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteVpnConnection require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteVpnConnectionResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteVpnGatewayRequest() (request *DeleteVpnGatewayRequest) {
    request = &DeleteVpnGatewayRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bmvpc", APIVersion, "DeleteVpnGateway")
    
    
    return
}

func NewDeleteVpnGatewayResponse() (response *DeleteVpnGatewayResponse) {
    response = &DeleteVpnGatewayResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteVpnGateway
// 本接口（DeleteVpnGateway）用于删除VPN网关。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_VPNGWNOTEXIST = "ResourceNotFound.VpnGwNotExist"
//  UNSUPPORTEDOPERATION_VPNCONNINUSE = "UnsupportedOperation.VpnConnInUse"
func (c *Client) DeleteVpnGateway(request *DeleteVpnGatewayRequest) (response *DeleteVpnGatewayResponse, err error) {
    return c.DeleteVpnGatewayWithContext(context.Background(), request)
}

// DeleteVpnGateway
// 本接口（DeleteVpnGateway）用于删除VPN网关。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_VPNGWNOTEXIST = "ResourceNotFound.VpnGwNotExist"
//  UNSUPPORTEDOPERATION_VPNCONNINUSE = "UnsupportedOperation.VpnConnInUse"
func (c *Client) DeleteVpnGatewayWithContext(ctx context.Context, request *DeleteVpnGatewayRequest) (response *DeleteVpnGatewayResponse, err error) {
    if request == nil {
        request = NewDeleteVpnGatewayRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteVpnGateway require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteVpnGatewayResponse()
    err = c.Send(request, response)
    return
}

func NewDeregisterIpsRequest() (request *DeregisterIpsRequest) {
    request = &DeregisterIpsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bmvpc", APIVersion, "DeregisterIps")
    
    
    return
}

func NewDeregisterIpsResponse() (response *DeregisterIpsResponse) {
    response = &DeregisterIpsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeregisterIps
// 注销私有网络IP为空闲
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeregisterIps(request *DeregisterIpsRequest) (response *DeregisterIpsResponse, err error) {
    return c.DeregisterIpsWithContext(context.Background(), request)
}

// DeregisterIps
// 注销私有网络IP为空闲
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeregisterIpsWithContext(ctx context.Context, request *DeregisterIpsRequest) (response *DeregisterIpsResponse, err error) {
    if request == nil {
        request = NewDeregisterIpsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeregisterIps require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeregisterIpsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCustomerGatewaysRequest() (request *DescribeCustomerGatewaysRequest) {
    request = &DescribeCustomerGatewaysRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bmvpc", APIVersion, "DescribeCustomerGateways")
    
    
    return
}

func NewDescribeCustomerGatewaysResponse() (response *DescribeCustomerGatewaysResponse) {
    response = &DescribeCustomerGatewaysResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCustomerGateways
// 本接口（DescribeCustomerGateways）用于查询对端网关列表。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeCustomerGateways(request *DescribeCustomerGatewaysRequest) (response *DescribeCustomerGatewaysResponse, err error) {
    return c.DescribeCustomerGatewaysWithContext(context.Background(), request)
}

// DescribeCustomerGateways
// 本接口（DescribeCustomerGateways）用于查询对端网关列表。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeCustomerGatewaysWithContext(ctx context.Context, request *DescribeCustomerGatewaysRequest) (response *DescribeCustomerGatewaysResponse, err error) {
    if request == nil {
        request = NewDescribeCustomerGatewaysRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCustomerGateways require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCustomerGatewaysResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeNatGatewaysRequest() (request *DescribeNatGatewaysRequest) {
    request = &DescribeNatGatewaysRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bmvpc", APIVersion, "DescribeNatGateways")
    
    
    return
}

func NewDescribeNatGatewaysResponse() (response *DescribeNatGatewaysResponse) {
    response = &DescribeNatGatewaysResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeNatGateways
// 获取NAT网关信息，包括NAT网关 ID、网关名称、私有网络、网关并发连接上限、绑定EIP列表等
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeNatGateways(request *DescribeNatGatewaysRequest) (response *DescribeNatGatewaysResponse, err error) {
    return c.DescribeNatGatewaysWithContext(context.Background(), request)
}

// DescribeNatGateways
// 获取NAT网关信息，包括NAT网关 ID、网关名称、私有网络、网关并发连接上限、绑定EIP列表等
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeNatGatewaysWithContext(ctx context.Context, request *DescribeNatGatewaysRequest) (response *DescribeNatGatewaysResponse, err error) {
    if request == nil {
        request = NewDescribeNatGatewaysRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeNatGateways require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeNatGatewaysResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeNatSubnetsRequest() (request *DescribeNatSubnetsRequest) {
    request = &DescribeNatSubnetsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bmvpc", APIVersion, "DescribeNatSubnets")
    
    
    return
}

func NewDescribeNatSubnetsResponse() (response *DescribeNatSubnetsResponse) {
    response = &DescribeNatSubnetsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeNatSubnets
// 可获取NAT网关绑定的子网信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeNatSubnets(request *DescribeNatSubnetsRequest) (response *DescribeNatSubnetsResponse, err error) {
    return c.DescribeNatSubnetsWithContext(context.Background(), request)
}

// DescribeNatSubnets
// 可获取NAT网关绑定的子网信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeNatSubnetsWithContext(ctx context.Context, request *DescribeNatSubnetsRequest) (response *DescribeNatSubnetsResponse, err error) {
    if request == nil {
        request = NewDescribeNatSubnetsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeNatSubnets require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeNatSubnetsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRoutePoliciesRequest() (request *DescribeRoutePoliciesRequest) {
    request = &DescribeRoutePoliciesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bmvpc", APIVersion, "DescribeRoutePolicies")
    
    
    return
}

func NewDescribeRoutePoliciesResponse() (response *DescribeRoutePoliciesResponse) {
    response = &DescribeRoutePoliciesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeRoutePolicies
// 本接口（DescribeRoutePolicies）用于查询路由表条目。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeRoutePolicies(request *DescribeRoutePoliciesRequest) (response *DescribeRoutePoliciesResponse, err error) {
    return c.DescribeRoutePoliciesWithContext(context.Background(), request)
}

// DescribeRoutePolicies
// 本接口（DescribeRoutePolicies）用于查询路由表条目。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeRoutePoliciesWithContext(ctx context.Context, request *DescribeRoutePoliciesRequest) (response *DescribeRoutePoliciesResponse, err error) {
    if request == nil {
        request = NewDescribeRoutePoliciesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRoutePolicies require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRoutePoliciesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRouteTablesRequest() (request *DescribeRouteTablesRequest) {
    request = &DescribeRouteTablesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bmvpc", APIVersion, "DescribeRouteTables")
    
    
    return
}

func NewDescribeRouteTablesResponse() (response *DescribeRouteTablesResponse) {
    response = &DescribeRouteTablesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeRouteTables
// 本接口（DescribeRouteTables）用于查询路由表。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeRouteTables(request *DescribeRouteTablesRequest) (response *DescribeRouteTablesResponse, err error) {
    return c.DescribeRouteTablesWithContext(context.Background(), request)
}

// DescribeRouteTables
// 本接口（DescribeRouteTables）用于查询路由表。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeRouteTablesWithContext(ctx context.Context, request *DescribeRouteTablesRequest) (response *DescribeRouteTablesResponse, err error) {
    if request == nil {
        request = NewDescribeRouteTablesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRouteTables require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRouteTablesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSubnetAvailableIpsRequest() (request *DescribeSubnetAvailableIpsRequest) {
    request = &DescribeSubnetAvailableIpsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bmvpc", APIVersion, "DescribeSubnetAvailableIps")
    
    
    return
}

func NewDescribeSubnetAvailableIpsResponse() (response *DescribeSubnetAvailableIpsResponse) {
    response = &DescribeSubnetAvailableIpsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeSubnetAvailableIps
// 获取子网内可用IP列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeSubnetAvailableIps(request *DescribeSubnetAvailableIpsRequest) (response *DescribeSubnetAvailableIpsResponse, err error) {
    return c.DescribeSubnetAvailableIpsWithContext(context.Background(), request)
}

// DescribeSubnetAvailableIps
// 获取子网内可用IP列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeSubnetAvailableIpsWithContext(ctx context.Context, request *DescribeSubnetAvailableIpsRequest) (response *DescribeSubnetAvailableIpsResponse, err error) {
    if request == nil {
        request = NewDescribeSubnetAvailableIpsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSubnetAvailableIps require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSubnetAvailableIpsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSubnetByDeviceRequest() (request *DescribeSubnetByDeviceRequest) {
    request = &DescribeSubnetByDeviceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bmvpc", APIVersion, "DescribeSubnetByDevice")
    
    
    return
}

func NewDescribeSubnetByDeviceResponse() (response *DescribeSubnetByDeviceResponse) {
    response = &DescribeSubnetByDeviceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeSubnetByDevice
// 物理机可以加入物理机子网，虚拟子网，DOCKER子网，通过此接口可以查询物理机加入的子网。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeSubnetByDevice(request *DescribeSubnetByDeviceRequest) (response *DescribeSubnetByDeviceResponse, err error) {
    return c.DescribeSubnetByDeviceWithContext(context.Background(), request)
}

// DescribeSubnetByDevice
// 物理机可以加入物理机子网，虚拟子网，DOCKER子网，通过此接口可以查询物理机加入的子网。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeSubnetByDeviceWithContext(ctx context.Context, request *DescribeSubnetByDeviceRequest) (response *DescribeSubnetByDeviceResponse, err error) {
    if request == nil {
        request = NewDescribeSubnetByDeviceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSubnetByDevice require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSubnetByDeviceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSubnetByHostedDeviceRequest() (request *DescribeSubnetByHostedDeviceRequest) {
    request = &DescribeSubnetByHostedDeviceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bmvpc", APIVersion, "DescribeSubnetByHostedDevice")
    
    
    return
}

func NewDescribeSubnetByHostedDeviceResponse() (response *DescribeSubnetByHostedDeviceResponse) {
    response = &DescribeSubnetByHostedDeviceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeSubnetByHostedDevice
// 托管可以加入物理机子网，虚拟子网，DOCKER子网，通过此接口可以查询托管加入的子网。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeSubnetByHostedDevice(request *DescribeSubnetByHostedDeviceRequest) (response *DescribeSubnetByHostedDeviceResponse, err error) {
    return c.DescribeSubnetByHostedDeviceWithContext(context.Background(), request)
}

// DescribeSubnetByHostedDevice
// 托管可以加入物理机子网，虚拟子网，DOCKER子网，通过此接口可以查询托管加入的子网。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeSubnetByHostedDeviceWithContext(ctx context.Context, request *DescribeSubnetByHostedDeviceRequest) (response *DescribeSubnetByHostedDeviceResponse, err error) {
    if request == nil {
        request = NewDescribeSubnetByHostedDeviceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSubnetByHostedDevice require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSubnetByHostedDeviceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSubnetsRequest() (request *DescribeSubnetsRequest) {
    request = &DescribeSubnetsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bmvpc", APIVersion, "DescribeSubnets")
    
    
    return
}

func NewDescribeSubnetsResponse() (response *DescribeSubnetsResponse) {
    response = &DescribeSubnetsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeSubnets
// 本接口（DescribeSubnets）用于查询黑石子网列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeSubnets(request *DescribeSubnetsRequest) (response *DescribeSubnetsResponse, err error) {
    return c.DescribeSubnetsWithContext(context.Background(), request)
}

// DescribeSubnets
// 本接口（DescribeSubnets）用于查询黑石子网列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeSubnetsWithContext(ctx context.Context, request *DescribeSubnetsRequest) (response *DescribeSubnetsResponse, err error) {
    if request == nil {
        request = NewDescribeSubnetsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSubnets require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSubnetsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTaskStatusRequest() (request *DescribeTaskStatusRequest) {
    request = &DescribeTaskStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bmvpc", APIVersion, "DescribeTaskStatus")
    
    
    return
}

func NewDescribeTaskStatusResponse() (response *DescribeTaskStatusResponse) {
    response = &DescribeTaskStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeTaskStatus
// 根据任务ID，获取任务的执行状态
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeTaskStatus(request *DescribeTaskStatusRequest) (response *DescribeTaskStatusResponse, err error) {
    return c.DescribeTaskStatusWithContext(context.Background(), request)
}

// DescribeTaskStatus
// 根据任务ID，获取任务的执行状态
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeTaskStatusWithContext(ctx context.Context, request *DescribeTaskStatusRequest) (response *DescribeTaskStatusResponse, err error) {
    if request == nil {
        request = NewDescribeTaskStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTaskStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTaskStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVpcPeerConnectionsRequest() (request *DescribeVpcPeerConnectionsRequest) {
    request = &DescribeVpcPeerConnectionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bmvpc", APIVersion, "DescribeVpcPeerConnections")
    
    
    return
}

func NewDescribeVpcPeerConnectionsResponse() (response *DescribeVpcPeerConnectionsResponse) {
    response = &DescribeVpcPeerConnectionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeVpcPeerConnections
// 获取对等连接列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeVpcPeerConnections(request *DescribeVpcPeerConnectionsRequest) (response *DescribeVpcPeerConnectionsResponse, err error) {
    return c.DescribeVpcPeerConnectionsWithContext(context.Background(), request)
}

// DescribeVpcPeerConnections
// 获取对等连接列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeVpcPeerConnectionsWithContext(ctx context.Context, request *DescribeVpcPeerConnectionsRequest) (response *DescribeVpcPeerConnectionsResponse, err error) {
    if request == nil {
        request = NewDescribeVpcPeerConnectionsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVpcPeerConnections require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVpcPeerConnectionsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVpcQuotaRequest() (request *DescribeVpcQuotaRequest) {
    request = &DescribeVpcQuotaRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bmvpc", APIVersion, "DescribeVpcQuota")
    
    
    return
}

func NewDescribeVpcQuotaResponse() (response *DescribeVpcQuotaResponse) {
    response = &DescribeVpcQuotaResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeVpcQuota
// 本接口（DescribeVpcQuota）用于查询用户VPC相关配额限制。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeVpcQuota(request *DescribeVpcQuotaRequest) (response *DescribeVpcQuotaResponse, err error) {
    return c.DescribeVpcQuotaWithContext(context.Background(), request)
}

// DescribeVpcQuota
// 本接口（DescribeVpcQuota）用于查询用户VPC相关配额限制。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeVpcQuotaWithContext(ctx context.Context, request *DescribeVpcQuotaRequest) (response *DescribeVpcQuotaResponse, err error) {
    if request == nil {
        request = NewDescribeVpcQuotaRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVpcQuota require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVpcQuotaResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVpcResourceRequest() (request *DescribeVpcResourceRequest) {
    request = &DescribeVpcResourceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bmvpc", APIVersion, "DescribeVpcResource")
    
    
    return
}

func NewDescribeVpcResourceResponse() (response *DescribeVpcResourceResponse) {
    response = &DescribeVpcResourceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeVpcResource
// 查询黑石私有网络关联资源
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeVpcResource(request *DescribeVpcResourceRequest) (response *DescribeVpcResourceResponse, err error) {
    return c.DescribeVpcResourceWithContext(context.Background(), request)
}

// DescribeVpcResource
// 查询黑石私有网络关联资源
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeVpcResourceWithContext(ctx context.Context, request *DescribeVpcResourceRequest) (response *DescribeVpcResourceResponse, err error) {
    if request == nil {
        request = NewDescribeVpcResourceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVpcResource require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVpcResourceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVpcViewRequest() (request *DescribeVpcViewRequest) {
    request = &DescribeVpcViewRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bmvpc", APIVersion, "DescribeVpcView")
    
    
    return
}

func NewDescribeVpcViewResponse() (response *DescribeVpcViewResponse) {
    response = &DescribeVpcViewResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeVpcView
// 本接口（DescribeVpcView）用于查询VPC网络拓扑视图。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeVpcView(request *DescribeVpcViewRequest) (response *DescribeVpcViewResponse, err error) {
    return c.DescribeVpcViewWithContext(context.Background(), request)
}

// DescribeVpcView
// 本接口（DescribeVpcView）用于查询VPC网络拓扑视图。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeVpcViewWithContext(ctx context.Context, request *DescribeVpcViewRequest) (response *DescribeVpcViewResponse, err error) {
    if request == nil {
        request = NewDescribeVpcViewRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVpcView require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVpcViewResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVpcsRequest() (request *DescribeVpcsRequest) {
    request = &DescribeVpcsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bmvpc", APIVersion, "DescribeVpcs")
    
    
    return
}

func NewDescribeVpcsResponse() (response *DescribeVpcsResponse) {
    response = &DescribeVpcsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeVpcs
// 本接口（DescribeVpcs）用于查询私有网络列表。
//
// 本接口不传参数时，返回默认排序下的前20条VPC信息。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeVpcs(request *DescribeVpcsRequest) (response *DescribeVpcsResponse, err error) {
    return c.DescribeVpcsWithContext(context.Background(), request)
}

// DescribeVpcs
// 本接口（DescribeVpcs）用于查询私有网络列表。
//
// 本接口不传参数时，返回默认排序下的前20条VPC信息。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeVpcsWithContext(ctx context.Context, request *DescribeVpcsRequest) (response *DescribeVpcsResponse, err error) {
    if request == nil {
        request = NewDescribeVpcsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVpcs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVpcsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVpnConnectionsRequest() (request *DescribeVpnConnectionsRequest) {
    request = &DescribeVpnConnectionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bmvpc", APIVersion, "DescribeVpnConnections")
    
    
    return
}

func NewDescribeVpnConnectionsResponse() (response *DescribeVpnConnectionsResponse) {
    response = &DescribeVpnConnectionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeVpnConnections
//  本接口（DescribeVpnConnections）查询VPN通道列表。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeVpnConnections(request *DescribeVpnConnectionsRequest) (response *DescribeVpnConnectionsResponse, err error) {
    return c.DescribeVpnConnectionsWithContext(context.Background(), request)
}

// DescribeVpnConnections
//  本接口（DescribeVpnConnections）查询VPN通道列表。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeVpnConnectionsWithContext(ctx context.Context, request *DescribeVpnConnectionsRequest) (response *DescribeVpnConnectionsResponse, err error) {
    if request == nil {
        request = NewDescribeVpnConnectionsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVpnConnections require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVpnConnectionsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVpnGatewaysRequest() (request *DescribeVpnGatewaysRequest) {
    request = &DescribeVpnGatewaysRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bmvpc", APIVersion, "DescribeVpnGateways")
    
    
    return
}

func NewDescribeVpnGatewaysResponse() (response *DescribeVpnGatewaysResponse) {
    response = &DescribeVpnGatewaysResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeVpnGateways
// 本接口（DescribeVpnGateways）用于查询VPN网关列表。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeVpnGateways(request *DescribeVpnGatewaysRequest) (response *DescribeVpnGatewaysResponse, err error) {
    return c.DescribeVpnGatewaysWithContext(context.Background(), request)
}

// DescribeVpnGateways
// 本接口（DescribeVpnGateways）用于查询VPN网关列表。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeVpnGatewaysWithContext(ctx context.Context, request *DescribeVpnGatewaysRequest) (response *DescribeVpnGatewaysResponse, err error) {
    if request == nil {
        request = NewDescribeVpnGatewaysRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVpnGateways require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVpnGatewaysResponse()
    err = c.Send(request, response)
    return
}

func NewDownloadCustomerGatewayConfigurationRequest() (request *DownloadCustomerGatewayConfigurationRequest) {
    request = &DownloadCustomerGatewayConfigurationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bmvpc", APIVersion, "DownloadCustomerGatewayConfiguration")
    
    
    return
}

func NewDownloadCustomerGatewayConfigurationResponse() (response *DownloadCustomerGatewayConfigurationResponse) {
    response = &DownloadCustomerGatewayConfigurationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DownloadCustomerGatewayConfiguration
// 本接口(DownloadCustomerGatewayConfiguration)用于下载VPN通道配置。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_VPNCONNNOTEXIST = "ResourceNotFound.VpnConnNotExist"
func (c *Client) DownloadCustomerGatewayConfiguration(request *DownloadCustomerGatewayConfigurationRequest) (response *DownloadCustomerGatewayConfigurationResponse, err error) {
    return c.DownloadCustomerGatewayConfigurationWithContext(context.Background(), request)
}

// DownloadCustomerGatewayConfiguration
// 本接口(DownloadCustomerGatewayConfiguration)用于下载VPN通道配置。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_VPNCONNNOTEXIST = "ResourceNotFound.VpnConnNotExist"
func (c *Client) DownloadCustomerGatewayConfigurationWithContext(ctx context.Context, request *DownloadCustomerGatewayConfigurationRequest) (response *DownloadCustomerGatewayConfigurationResponse, err error) {
    if request == nil {
        request = NewDownloadCustomerGatewayConfigurationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DownloadCustomerGatewayConfiguration require credential")
    }

    request.SetContext(ctx)
    
    response = NewDownloadCustomerGatewayConfigurationResponse()
    err = c.Send(request, response)
    return
}

func NewModifyCustomerGatewayAttributeRequest() (request *ModifyCustomerGatewayAttributeRequest) {
    request = &ModifyCustomerGatewayAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bmvpc", APIVersion, "ModifyCustomerGatewayAttribute")
    
    
    return
}

func NewModifyCustomerGatewayAttributeResponse() (response *ModifyCustomerGatewayAttributeResponse) {
    response = &ModifyCustomerGatewayAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyCustomerGatewayAttribute
// 本接口（ModifyCustomerGatewayAttribute）用于修改对端网关信息。
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_CUSTOMERGATEWAYNOTEXIST = "ResourceNotFound.CustomerGatewayNotExist"
//  UNSUPPORTEDOPERATION_CUSTOMERGATEWAYADDREXIST = "UnsupportedOperation.CustomerGatewayAddrExist"
//  UNSUPPORTEDOPERATION_CUSTOMERGATEWAYADDRINVALID = "UnsupportedOperation.CustomerGatewayAddrInvalid"
//  UNSUPPORTEDOPERATION_VPNCONNINUSE = "UnsupportedOperation.VpnConnInUse"
func (c *Client) ModifyCustomerGatewayAttribute(request *ModifyCustomerGatewayAttributeRequest) (response *ModifyCustomerGatewayAttributeResponse, err error) {
    return c.ModifyCustomerGatewayAttributeWithContext(context.Background(), request)
}

// ModifyCustomerGatewayAttribute
// 本接口（ModifyCustomerGatewayAttribute）用于修改对端网关信息。
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_CUSTOMERGATEWAYNOTEXIST = "ResourceNotFound.CustomerGatewayNotExist"
//  UNSUPPORTEDOPERATION_CUSTOMERGATEWAYADDREXIST = "UnsupportedOperation.CustomerGatewayAddrExist"
//  UNSUPPORTEDOPERATION_CUSTOMERGATEWAYADDRINVALID = "UnsupportedOperation.CustomerGatewayAddrInvalid"
//  UNSUPPORTEDOPERATION_VPNCONNINUSE = "UnsupportedOperation.VpnConnInUse"
func (c *Client) ModifyCustomerGatewayAttributeWithContext(ctx context.Context, request *ModifyCustomerGatewayAttributeRequest) (response *ModifyCustomerGatewayAttributeResponse, err error) {
    if request == nil {
        request = NewModifyCustomerGatewayAttributeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyCustomerGatewayAttribute require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyCustomerGatewayAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyRoutePolicyRequest() (request *ModifyRoutePolicyRequest) {
    request = &ModifyRoutePolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bmvpc", APIVersion, "ModifyRoutePolicy")
    
    
    return
}

func NewModifyRoutePolicyResponse() (response *ModifyRoutePolicyResponse) {
    response = &ModifyRoutePolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyRoutePolicy
// 修改自定义路由
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyRoutePolicy(request *ModifyRoutePolicyRequest) (response *ModifyRoutePolicyResponse, err error) {
    return c.ModifyRoutePolicyWithContext(context.Background(), request)
}

// ModifyRoutePolicy
// 修改自定义路由
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyRoutePolicyWithContext(ctx context.Context, request *ModifyRoutePolicyRequest) (response *ModifyRoutePolicyResponse, err error) {
    if request == nil {
        request = NewModifyRoutePolicyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyRoutePolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyRoutePolicyResponse()
    err = c.Send(request, response)
    return
}

func NewModifyRouteTableRequest() (request *ModifyRouteTableRequest) {
    request = &ModifyRouteTableRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bmvpc", APIVersion, "ModifyRouteTable")
    
    
    return
}

func NewModifyRouteTableResponse() (response *ModifyRouteTableResponse) {
    response = &ModifyRouteTableResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyRouteTable
// 修改路由表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyRouteTable(request *ModifyRouteTableRequest) (response *ModifyRouteTableResponse, err error) {
    return c.ModifyRouteTableWithContext(context.Background(), request)
}

// ModifyRouteTable
// 修改路由表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyRouteTableWithContext(ctx context.Context, request *ModifyRouteTableRequest) (response *ModifyRouteTableResponse, err error) {
    if request == nil {
        request = NewModifyRouteTableRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyRouteTable require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyRouteTableResponse()
    err = c.Send(request, response)
    return
}

func NewModifySubnetAttributeRequest() (request *ModifySubnetAttributeRequest) {
    request = &ModifySubnetAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bmvpc", APIVersion, "ModifySubnetAttribute")
    
    
    return
}

func NewModifySubnetAttributeResponse() (response *ModifySubnetAttributeResponse) {
    response = &ModifySubnetAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifySubnetAttribute
// 修改子网属性
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifySubnetAttribute(request *ModifySubnetAttributeRequest) (response *ModifySubnetAttributeResponse, err error) {
    return c.ModifySubnetAttributeWithContext(context.Background(), request)
}

// ModifySubnetAttribute
// 修改子网属性
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifySubnetAttributeWithContext(ctx context.Context, request *ModifySubnetAttributeRequest) (response *ModifySubnetAttributeResponse, err error) {
    if request == nil {
        request = NewModifySubnetAttributeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifySubnetAttribute require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifySubnetAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewModifySubnetDHCPRelayRequest() (request *ModifySubnetDHCPRelayRequest) {
    request = &ModifySubnetDHCPRelayRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bmvpc", APIVersion, "ModifySubnetDHCPRelay")
    
    
    return
}

func NewModifySubnetDHCPRelayResponse() (response *ModifySubnetDHCPRelayResponse) {
    response = &ModifySubnetDHCPRelayResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifySubnetDHCPRelay
// 修改子网DHCP Relay属性
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifySubnetDHCPRelay(request *ModifySubnetDHCPRelayRequest) (response *ModifySubnetDHCPRelayResponse, err error) {
    return c.ModifySubnetDHCPRelayWithContext(context.Background(), request)
}

// ModifySubnetDHCPRelay
// 修改子网DHCP Relay属性
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifySubnetDHCPRelayWithContext(ctx context.Context, request *ModifySubnetDHCPRelayRequest) (response *ModifySubnetDHCPRelayResponse, err error) {
    if request == nil {
        request = NewModifySubnetDHCPRelayRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifySubnetDHCPRelay require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifySubnetDHCPRelayResponse()
    err = c.Send(request, response)
    return
}

func NewModifyVpcAttributeRequest() (request *ModifyVpcAttributeRequest) {
    request = &ModifyVpcAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bmvpc", APIVersion, "ModifyVpcAttribute")
    
    
    return
}

func NewModifyVpcAttributeResponse() (response *ModifyVpcAttributeResponse) {
    response = &ModifyVpcAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyVpcAttribute
// 本接口（ModifyVpcAttribute）用于修改VPC的标识名称和控制VPC的监控起停。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyVpcAttribute(request *ModifyVpcAttributeRequest) (response *ModifyVpcAttributeResponse, err error) {
    return c.ModifyVpcAttributeWithContext(context.Background(), request)
}

// ModifyVpcAttribute
// 本接口（ModifyVpcAttribute）用于修改VPC的标识名称和控制VPC的监控起停。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyVpcAttributeWithContext(ctx context.Context, request *ModifyVpcAttributeRequest) (response *ModifyVpcAttributeResponse, err error) {
    if request == nil {
        request = NewModifyVpcAttributeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyVpcAttribute require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyVpcAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyVpcPeerConnectionRequest() (request *ModifyVpcPeerConnectionRequest) {
    request = &ModifyVpcPeerConnectionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bmvpc", APIVersion, "ModifyVpcPeerConnection")
    
    
    return
}

func NewModifyVpcPeerConnectionResponse() (response *ModifyVpcPeerConnectionResponse) {
    response = &ModifyVpcPeerConnectionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyVpcPeerConnection
// 修改黑石对等连接
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_VPCNOTEXIST = "ResourceNotFound.VpcNotExist"
//  RESOURCENOTFOUND_VPCPEERNOTEXIST = "ResourceNotFound.VpcPeerNotExist"
//  UNSUPPORTEDOPERATION_INVALIDVPCPEERSTATE = "UnsupportedOperation.InvalidVpcPeerState"
//  UNSUPPORTEDOPERATION_VPCCIDRCONFICT = "UnsupportedOperation.VpcCidrConfict"
//  UNSUPPORTEDOPERATION_VPCPEEREXIST = "UnsupportedOperation.VpcPeerExist"
func (c *Client) ModifyVpcPeerConnection(request *ModifyVpcPeerConnectionRequest) (response *ModifyVpcPeerConnectionResponse, err error) {
    return c.ModifyVpcPeerConnectionWithContext(context.Background(), request)
}

// ModifyVpcPeerConnection
// 修改黑石对等连接
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_VPCNOTEXIST = "ResourceNotFound.VpcNotExist"
//  RESOURCENOTFOUND_VPCPEERNOTEXIST = "ResourceNotFound.VpcPeerNotExist"
//  UNSUPPORTEDOPERATION_INVALIDVPCPEERSTATE = "UnsupportedOperation.InvalidVpcPeerState"
//  UNSUPPORTEDOPERATION_VPCCIDRCONFICT = "UnsupportedOperation.VpcCidrConfict"
//  UNSUPPORTEDOPERATION_VPCPEEREXIST = "UnsupportedOperation.VpcPeerExist"
func (c *Client) ModifyVpcPeerConnectionWithContext(ctx context.Context, request *ModifyVpcPeerConnectionRequest) (response *ModifyVpcPeerConnectionResponse, err error) {
    if request == nil {
        request = NewModifyVpcPeerConnectionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyVpcPeerConnection require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyVpcPeerConnectionResponse()
    err = c.Send(request, response)
    return
}

func NewModifyVpnConnectionAttributeRequest() (request *ModifyVpnConnectionAttributeRequest) {
    request = &ModifyVpnConnectionAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bmvpc", APIVersion, "ModifyVpnConnectionAttribute")
    
    
    return
}

func NewModifyVpnConnectionAttributeResponse() (response *ModifyVpnConnectionAttributeResponse) {
    response = &ModifyVpnConnectionAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyVpnConnectionAttribute
// 本接口（ModifyVpnConnectionAttribute）用于修改VPN通道。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_SPDACLLIMIT = "LimitExceeded.SpdAclLimit"
//  LIMITEXCEEDED_SPDDNETLIMIT = "LimitExceeded.SpdDnetLimit"
//  LIMITEXCEEDED_SPDSNETLIMIT = "LimitExceeded.SpdSnetLimit"
//  RESOURCENOTFOUND_NOTAVAIBLE = "ResourceNotFound.NotAvaible"
//  RESOURCENOTFOUND_VPCNOTEXIST = "ResourceNotFound.VpcNotExist"
//  RESOURCENOTFOUND_VPNCONNNOTEXIST = "ResourceNotFound.VpnConnNotExist"
//  UNSUPPORTEDOPERATION_INVALIDVPNCONNSTATE = "UnsupportedOperation.InvalidVpnConnState"
//  UNSUPPORTEDOPERATION_SPDACLCIDRINVALID = "UnsupportedOperation.SpdAclCidrInvalid"
//  UNSUPPORTEDOPERATION_SPDSNETNOTINCIDR = "UnsupportedOperation.SpdSnetNotInCidr"
func (c *Client) ModifyVpnConnectionAttribute(request *ModifyVpnConnectionAttributeRequest) (response *ModifyVpnConnectionAttributeResponse, err error) {
    return c.ModifyVpnConnectionAttributeWithContext(context.Background(), request)
}

// ModifyVpnConnectionAttribute
// 本接口（ModifyVpnConnectionAttribute）用于修改VPN通道。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_SPDACLLIMIT = "LimitExceeded.SpdAclLimit"
//  LIMITEXCEEDED_SPDDNETLIMIT = "LimitExceeded.SpdDnetLimit"
//  LIMITEXCEEDED_SPDSNETLIMIT = "LimitExceeded.SpdSnetLimit"
//  RESOURCENOTFOUND_NOTAVAIBLE = "ResourceNotFound.NotAvaible"
//  RESOURCENOTFOUND_VPCNOTEXIST = "ResourceNotFound.VpcNotExist"
//  RESOURCENOTFOUND_VPNCONNNOTEXIST = "ResourceNotFound.VpnConnNotExist"
//  UNSUPPORTEDOPERATION_INVALIDVPNCONNSTATE = "UnsupportedOperation.InvalidVpnConnState"
//  UNSUPPORTEDOPERATION_SPDACLCIDRINVALID = "UnsupportedOperation.SpdAclCidrInvalid"
//  UNSUPPORTEDOPERATION_SPDSNETNOTINCIDR = "UnsupportedOperation.SpdSnetNotInCidr"
func (c *Client) ModifyVpnConnectionAttributeWithContext(ctx context.Context, request *ModifyVpnConnectionAttributeRequest) (response *ModifyVpnConnectionAttributeResponse, err error) {
    if request == nil {
        request = NewModifyVpnConnectionAttributeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyVpnConnectionAttribute require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyVpnConnectionAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyVpnGatewayAttributeRequest() (request *ModifyVpnGatewayAttributeRequest) {
    request = &ModifyVpnGatewayAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bmvpc", APIVersion, "ModifyVpnGatewayAttribute")
    
    
    return
}

func NewModifyVpnGatewayAttributeResponse() (response *ModifyVpnGatewayAttributeResponse) {
    response = &ModifyVpnGatewayAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyVpnGatewayAttribute
// 本接口（ModifyVpnGatewayAttribute）用于修改VPN网关属性。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_VPNGWNOTEXIST = "ResourceNotFound.VpnGwNotExist"
//  UNSUPPORTEDOPERATION_VPNCONNINUSE = "UnsupportedOperation.VpnConnInUse"
func (c *Client) ModifyVpnGatewayAttribute(request *ModifyVpnGatewayAttributeRequest) (response *ModifyVpnGatewayAttributeResponse, err error) {
    return c.ModifyVpnGatewayAttributeWithContext(context.Background(), request)
}

// ModifyVpnGatewayAttribute
// 本接口（ModifyVpnGatewayAttribute）用于修改VPN网关属性。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_VPNGWNOTEXIST = "ResourceNotFound.VpnGwNotExist"
//  UNSUPPORTEDOPERATION_VPNCONNINUSE = "UnsupportedOperation.VpnConnInUse"
func (c *Client) ModifyVpnGatewayAttributeWithContext(ctx context.Context, request *ModifyVpnGatewayAttributeRequest) (response *ModifyVpnGatewayAttributeResponse, err error) {
    if request == nil {
        request = NewModifyVpnGatewayAttributeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyVpnGatewayAttribute require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyVpnGatewayAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewRejectVpcPeerConnectionRequest() (request *RejectVpcPeerConnectionRequest) {
    request = &RejectVpcPeerConnectionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bmvpc", APIVersion, "RejectVpcPeerConnection")
    
    
    return
}

func NewRejectVpcPeerConnectionResponse() (response *RejectVpcPeerConnectionResponse) {
    response = &RejectVpcPeerConnectionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// RejectVpcPeerConnection
// 拒绝黑石对等连接申请
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_VPCNOTEXIST = "ResourceNotFound.VpcNotExist"
//  RESOURCENOTFOUND_VPCPEERNOTEXIST = "ResourceNotFound.VpcPeerNotExist"
//  UNSUPPORTEDOPERATION_INVALIDVPCPEERSTATE = "UnsupportedOperation.InvalidVpcPeerState"
//  UNSUPPORTEDOPERATION_VPCCIDRCONFICT = "UnsupportedOperation.VpcCidrConfict"
//  UNSUPPORTEDOPERATION_VPCPEEREXIST = "UnsupportedOperation.VpcPeerExist"
func (c *Client) RejectVpcPeerConnection(request *RejectVpcPeerConnectionRequest) (response *RejectVpcPeerConnectionResponse, err error) {
    return c.RejectVpcPeerConnectionWithContext(context.Background(), request)
}

// RejectVpcPeerConnection
// 拒绝黑石对等连接申请
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_VPCNOTEXIST = "ResourceNotFound.VpcNotExist"
//  RESOURCENOTFOUND_VPCPEERNOTEXIST = "ResourceNotFound.VpcPeerNotExist"
//  UNSUPPORTEDOPERATION_INVALIDVPCPEERSTATE = "UnsupportedOperation.InvalidVpcPeerState"
//  UNSUPPORTEDOPERATION_VPCCIDRCONFICT = "UnsupportedOperation.VpcCidrConfict"
//  UNSUPPORTEDOPERATION_VPCPEEREXIST = "UnsupportedOperation.VpcPeerExist"
func (c *Client) RejectVpcPeerConnectionWithContext(ctx context.Context, request *RejectVpcPeerConnectionRequest) (response *RejectVpcPeerConnectionResponse, err error) {
    if request == nil {
        request = NewRejectVpcPeerConnectionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RejectVpcPeerConnection require credential")
    }

    request.SetContext(ctx)
    
    response = NewRejectVpcPeerConnectionResponse()
    err = c.Send(request, response)
    return
}

func NewResetVpnConnectionRequest() (request *ResetVpnConnectionRequest) {
    request = &ResetVpnConnectionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bmvpc", APIVersion, "ResetVpnConnection")
    
    
    return
}

func NewResetVpnConnectionResponse() (response *ResetVpnConnectionResponse) {
    response = &ResetVpnConnectionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ResetVpnConnection
// 本接口(ResetVpnConnection)用于重置VPN通道。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_NOTAVAIBLE = "ResourceNotFound.NotAvaible"
//  RESOURCENOTFOUND_VPCNOTEXIST = "ResourceNotFound.VpcNotExist"
//  RESOURCENOTFOUND_VPNCONNNOTEXIST = "ResourceNotFound.VpnConnNotExist"
func (c *Client) ResetVpnConnection(request *ResetVpnConnectionRequest) (response *ResetVpnConnectionResponse, err error) {
    return c.ResetVpnConnectionWithContext(context.Background(), request)
}

// ResetVpnConnection
// 本接口(ResetVpnConnection)用于重置VPN通道。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_NOTAVAIBLE = "ResourceNotFound.NotAvaible"
//  RESOURCENOTFOUND_VPCNOTEXIST = "ResourceNotFound.VpcNotExist"
//  RESOURCENOTFOUND_VPNCONNNOTEXIST = "ResourceNotFound.VpnConnNotExist"
func (c *Client) ResetVpnConnectionWithContext(ctx context.Context, request *ResetVpnConnectionRequest) (response *ResetVpnConnectionResponse, err error) {
    if request == nil {
        request = NewResetVpnConnectionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ResetVpnConnection require credential")
    }

    request.SetContext(ctx)
    
    response = NewResetVpnConnectionResponse()
    err = c.Send(request, response)
    return
}

func NewUnbindEipsFromNatGatewayRequest() (request *UnbindEipsFromNatGatewayRequest) {
    request = &UnbindEipsFromNatGatewayRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bmvpc", APIVersion, "UnbindEipsFromNatGateway")
    
    
    return
}

func NewUnbindEipsFromNatGatewayResponse() (response *UnbindEipsFromNatGatewayResponse) {
    response = &UnbindEipsFromNatGatewayResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UnbindEipsFromNatGateway
// NAT网关解绑该EIP后，NAT网关将不会使用该EIP作为访问外网的源IP地址
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) UnbindEipsFromNatGateway(request *UnbindEipsFromNatGatewayRequest) (response *UnbindEipsFromNatGatewayResponse, err error) {
    return c.UnbindEipsFromNatGatewayWithContext(context.Background(), request)
}

// UnbindEipsFromNatGateway
// NAT网关解绑该EIP后，NAT网关将不会使用该EIP作为访问外网的源IP地址
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) UnbindEipsFromNatGatewayWithContext(ctx context.Context, request *UnbindEipsFromNatGatewayRequest) (response *UnbindEipsFromNatGatewayResponse, err error) {
    if request == nil {
        request = NewUnbindEipsFromNatGatewayRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UnbindEipsFromNatGateway require credential")
    }

    request.SetContext(ctx)
    
    response = NewUnbindEipsFromNatGatewayResponse()
    err = c.Send(request, response)
    return
}

func NewUnbindIpsFromNatGatewayRequest() (request *UnbindIpsFromNatGatewayRequest) {
    request = &UnbindIpsFromNatGatewayRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bmvpc", APIVersion, "UnbindIpsFromNatGateway")
    
    
    return
}

func NewUnbindIpsFromNatGatewayResponse() (response *UnbindIpsFromNatGatewayResponse) {
    response = &UnbindIpsFromNatGatewayResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UnbindIpsFromNatGateway
// NAT网关解绑IP接口，可将子网的部分IP从NAT网关中解绑
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) UnbindIpsFromNatGateway(request *UnbindIpsFromNatGatewayRequest) (response *UnbindIpsFromNatGatewayResponse, err error) {
    return c.UnbindIpsFromNatGatewayWithContext(context.Background(), request)
}

// UnbindIpsFromNatGateway
// NAT网关解绑IP接口，可将子网的部分IP从NAT网关中解绑
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) UnbindIpsFromNatGatewayWithContext(ctx context.Context, request *UnbindIpsFromNatGatewayRequest) (response *UnbindIpsFromNatGatewayResponse, err error) {
    if request == nil {
        request = NewUnbindIpsFromNatGatewayRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UnbindIpsFromNatGateway require credential")
    }

    request.SetContext(ctx)
    
    response = NewUnbindIpsFromNatGatewayResponse()
    err = c.Send(request, response)
    return
}

func NewUnbindSubnetsFromNatGatewayRequest() (request *UnbindSubnetsFromNatGatewayRequest) {
    request = &UnbindSubnetsFromNatGatewayRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bmvpc", APIVersion, "UnbindSubnetsFromNatGateway")
    
    
    return
}

func NewUnbindSubnetsFromNatGatewayResponse() (response *UnbindSubnetsFromNatGatewayResponse) {
    response = &UnbindSubnetsFromNatGatewayResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UnbindSubnetsFromNatGateway
// NAT网关解绑子网接口，可将子网解绑NAT网关
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) UnbindSubnetsFromNatGateway(request *UnbindSubnetsFromNatGatewayRequest) (response *UnbindSubnetsFromNatGatewayResponse, err error) {
    return c.UnbindSubnetsFromNatGatewayWithContext(context.Background(), request)
}

// UnbindSubnetsFromNatGateway
// NAT网关解绑子网接口，可将子网解绑NAT网关
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) UnbindSubnetsFromNatGatewayWithContext(ctx context.Context, request *UnbindSubnetsFromNatGatewayRequest) (response *UnbindSubnetsFromNatGatewayResponse, err error) {
    if request == nil {
        request = NewUnbindSubnetsFromNatGatewayRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UnbindSubnetsFromNatGateway require credential")
    }

    request.SetContext(ctx)
    
    response = NewUnbindSubnetsFromNatGatewayResponse()
    err = c.Send(request, response)
    return
}

func NewUpgradeNatGatewayRequest() (request *UpgradeNatGatewayRequest) {
    request = &UpgradeNatGatewayRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bmvpc", APIVersion, "UpgradeNatGateway")
    
    
    return
}

func NewUpgradeNatGatewayResponse() (response *UpgradeNatGatewayResponse) {
    response = &UpgradeNatGatewayResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpgradeNatGateway
// 升级NAT网关接口，可NAT网关修改为小型NAT网关、中型NAT网关、以及大型NAT网关
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) UpgradeNatGateway(request *UpgradeNatGatewayRequest) (response *UpgradeNatGatewayResponse, err error) {
    return c.UpgradeNatGatewayWithContext(context.Background(), request)
}

// UpgradeNatGateway
// 升级NAT网关接口，可NAT网关修改为小型NAT网关、中型NAT网关、以及大型NAT网关
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) UpgradeNatGatewayWithContext(ctx context.Context, request *UpgradeNatGatewayRequest) (response *UpgradeNatGatewayResponse, err error) {
    if request == nil {
        request = NewUpgradeNatGatewayRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpgradeNatGateway require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpgradeNatGatewayResponse()
    err = c.Send(request, response)
    return
}
