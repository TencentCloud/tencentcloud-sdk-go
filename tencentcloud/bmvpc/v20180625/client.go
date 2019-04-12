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

func NewClient(credential *common.Credential, region string, clientProfile *profile.ClientProfile) (client *Client, err error) {
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

// 接受黑石对等连接
func (c *Client) AcceptVpcPeerConnection(request *AcceptVpcPeerConnectionRequest) (response *AcceptVpcPeerConnectionResponse, err error) {
    if request == nil {
        request = NewAcceptVpcPeerConnectionRequest()
    }
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

// 批量注册虚拟IP，异步接口。通过接口来查询任务进度。每次请求最多注册256个IP
func (c *Client) AsyncRegisterIps(request *AsyncRegisterIpsRequest) (response *AsyncRegisterIpsResponse, err error) {
    if request == nil {
        request = NewAsyncRegisterIpsRequest()
    }
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

// NAT网关绑定EIP接口，可将EIP绑定到NAT网关，该EIP作为访问外网的源IP地址，将流量发送到Internet
func (c *Client) BindEipsToNatGateway(request *BindEipsToNatGatewayRequest) (response *BindEipsToNatGatewayResponse, err error) {
    if request == nil {
        request = NewBindEipsToNatGatewayRequest()
    }
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

// 可用于将子网的部分IP绑定到NAT网关
func (c *Client) BindIpsToNatGateway(request *BindIpsToNatGatewayRequest) (response *BindIpsToNatGatewayResponse, err error) {
    if request == nil {
        request = NewBindIpsToNatGatewayRequest()
    }
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

// NAT网关绑定子网后，该子网内全部IP可出公网
func (c *Client) BindSubnetsToNatGateway(request *BindSubnetsToNatGatewayRequest) (response *BindSubnetsToNatGatewayResponse, err error) {
    if request == nil {
        request = NewBindSubnetsToNatGatewayRequest()
    }
    response = NewBindSubnetsToNatGatewayResponse()
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

// 创建黑石Docker子网， 如果不指定VlanId，将会分配2000--2999范围的VlanId; 子网会关闭分布式网关
func (c *Client) CreateDockerSubnetWithVlan(request *CreateDockerSubnetWithVlanRequest) (response *CreateDockerSubnetWithVlanResponse, err error) {
    if request == nil {
        request = NewCreateDockerSubnetWithVlanRequest()
    }
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

// 本接口（CreateHostedInterface）用于黑石托管机器加入带VLANID不为5的子网。
// 
// 1) 不能加入vlanId 为5的子网，只能加入VLANID范围为2000-2999的子网。
// 2) 每台托管机器最多可以加入20个子网。
// 3) 每次调用最多能支持传入10台托管机器。
func (c *Client) CreateHostedInterface(request *CreateHostedInterfaceRequest) (response *CreateHostedInterfaceResponse, err error) {
    if request == nil {
        request = NewCreateHostedInterfaceRequest()
    }
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

// 物理机加入子网
func (c *Client) CreateInterfaces(request *CreateInterfacesRequest) (response *CreateInterfacesResponse, err error) {
    if request == nil {
        request = NewCreateInterfacesRequest()
    }
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

// 创建NAT网关接口，可针对网段方式、子网全部IP、子网部分IP这三种方式创建NAT网关
func (c *Client) CreateNatGateway(request *CreateNatGatewayRequest) (response *CreateNatGatewayResponse, err error) {
    if request == nil {
        request = NewCreateNatGatewayRequest()
    }
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

// 创建黑石路由表的路由规则
func (c *Client) CreateRoutePolicies(request *CreateRoutePoliciesRequest) (response *CreateRoutePoliciesResponse, err error) {
    if request == nil {
        request = NewCreateRoutePoliciesRequest()
    }
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

// 创建黑石私有网络的子网
// 访问管理: 用户可以对VpcId进行授权操作。比如设置资源为["qcs::bmvpc:::unVpc/vpc-xxxxx"]
func (c *Client) CreateSubnet(request *CreateSubnetRequest) (response *CreateSubnetResponse, err error) {
    if request == nil {
        request = NewCreateSubnetRequest()
    }
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

// 创建黑石虚拟子网， 虚拟子网用于在黑石上创建虚拟网络，与黑石子网要做好规划。虚拟子网会分配2000-2999的VlanId。
func (c *Client) CreateVirtualSubnetWithVlan(request *CreateVirtualSubnetWithVlanRequest) (response *CreateVirtualSubnetWithVlanResponse, err error) {
    if request == nil {
        request = NewCreateVirtualSubnetWithVlanRequest()
    }
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

// 创建黑石私有网络
func (c *Client) CreateVpc(request *CreateVpcRequest) (response *CreateVpcResponse, err error) {
    if request == nil {
        request = NewCreateVpcRequest()
    }
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

// 创建对等连接
func (c *Client) CreateVpcPeerConnection(request *CreateVpcPeerConnectionRequest) (response *CreateVpcPeerConnectionResponse, err error) {
    if request == nil {
        request = NewCreateVpcPeerConnectionRequest()
    }
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

// 本接口（DeleteCustomerGateway）用于删除对端网关。
func (c *Client) DeleteCustomerGateway(request *DeleteCustomerGatewayRequest) (response *DeleteCustomerGatewayResponse, err error) {
    if request == nil {
        request = NewDeleteCustomerGatewayRequest()
    }
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

// 本接口用于托管机器从VLANID不为5的子网中移除。
// 1) 不能从vlanId 为5的子网中移除。
// 2) 每次调用最多能支持传入10台物理机。
func (c *Client) DeleteHostedInterface(request *DeleteHostedInterfaceRequest) (response *DeleteHostedInterfaceResponse, err error) {
    if request == nil {
        request = NewDeleteHostedInterfaceRequest()
    }
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

// 托管机器移除子网批量接口，传入一台托管机器和多个子网，批量移除这些子网。异步接口，接口返回TaskId。
func (c *Client) DeleteHostedInterfaces(request *DeleteHostedInterfacesRequest) (response *DeleteHostedInterfacesResponse, err error) {
    if request == nil {
        request = NewDeleteHostedInterfacesRequest()
    }
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

// 物理机移除子网批量接口，传入一台物理机和多个子网，批量移除这些子网。异步接口，接口返回TaskId。
func (c *Client) DeleteInterfaces(request *DeleteInterfacesRequest) (response *DeleteInterfacesResponse, err error) {
    if request == nil {
        request = NewDeleteInterfacesRequest()
    }
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

// 删除NAT网关
func (c *Client) DeleteNatGateway(request *DeleteNatGatewayRequest) (response *DeleteNatGatewayResponse, err error) {
    if request == nil {
        request = NewDeleteNatGatewayRequest()
    }
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

// 删除黑石路由表路由规则
func (c *Client) DeleteRoutePolicy(request *DeleteRoutePolicyRequest) (response *DeleteRoutePolicyResponse, err error) {
    if request == nil {
        request = NewDeleteRoutePolicyRequest()
    }
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

// 本接口（DeleteSubnet）用于删除黑石私有网络子网。
// 删除子网前，请清理该子网下所有资源，包括物理机、负载均衡、黑石数据库、弹性IP、NAT网关等资源
func (c *Client) DeleteSubnet(request *DeleteSubnetRequest) (response *DeleteSubnetResponse, err error) {
    if request == nil {
        request = NewDeleteSubnetRequest()
    }
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

// 退还虚拟IP。此接口只能退还虚拟IP，物理机IP不能退还。
func (c *Client) DeleteVirtualIp(request *DeleteVirtualIpRequest) (response *DeleteVirtualIpResponse, err error) {
    if request == nil {
        request = NewDeleteVirtualIpRequest()
    }
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

// 本接口(DeleteVpc)用于删除黑石私有网络(VPC)。
// 
// 删除私有网络前，请清理该私有网络下所有资源，包括子网、负载均衡、弹性 IP、对等连接、NAT 网关、专线通道、SSLVPN 等资源。
func (c *Client) DeleteVpc(request *DeleteVpcRequest) (response *DeleteVpcResponse, err error) {
    if request == nil {
        request = NewDeleteVpcRequest()
    }
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

// 删除黑石对等连接
func (c *Client) DeleteVpcPeerConnection(request *DeleteVpcPeerConnectionRequest) (response *DeleteVpcPeerConnectionResponse, err error) {
    if request == nil {
        request = NewDeleteVpcPeerConnectionRequest()
    }
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

// 本接口(DeleteVpnConnection)用于删除VPN通道。
func (c *Client) DeleteVpnConnection(request *DeleteVpnConnectionRequest) (response *DeleteVpnConnectionResponse, err error) {
    if request == nil {
        request = NewDeleteVpnConnectionRequest()
    }
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

// 本接口（DeleteVpnGateway）用于删除VPN网关。
func (c *Client) DeleteVpnGateway(request *DeleteVpnGatewayRequest) (response *DeleteVpnGatewayResponse, err error) {
    if request == nil {
        request = NewDeleteVpnGatewayRequest()
    }
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

// 注销私有网络IP为空闲
func (c *Client) DeregisterIps(request *DeregisterIpsRequest) (response *DeregisterIpsResponse, err error) {
    if request == nil {
        request = NewDeregisterIpsRequest()
    }
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

// 本接口（DescribeCustomerGateways）用于查询对端网关列表。
func (c *Client) DescribeCustomerGateways(request *DescribeCustomerGatewaysRequest) (response *DescribeCustomerGatewaysResponse, err error) {
    if request == nil {
        request = NewDescribeCustomerGatewaysRequest()
    }
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

// 获取NAT网关信息，包括NAT网关 ID、网关名称、私有网络、网关并发连接上限、绑定EIP列表等
func (c *Client) DescribeNatGateways(request *DescribeNatGatewaysRequest) (response *DescribeNatGatewaysResponse, err error) {
    if request == nil {
        request = NewDescribeNatGatewaysRequest()
    }
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

// 可获取NAT网关绑定的子网信息
func (c *Client) DescribeNatSubnets(request *DescribeNatSubnetsRequest) (response *DescribeNatSubnetsResponse, err error) {
    if request == nil {
        request = NewDescribeNatSubnetsRequest()
    }
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

// 本接口（DescribeRoutePolicies）用于查询路由表条目。
func (c *Client) DescribeRoutePolicies(request *DescribeRoutePoliciesRequest) (response *DescribeRoutePoliciesResponse, err error) {
    if request == nil {
        request = NewDescribeRoutePoliciesRequest()
    }
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

// 本接口（DescribeRouteTables）用于查询路由表。
func (c *Client) DescribeRouteTables(request *DescribeRouteTablesRequest) (response *DescribeRouteTablesResponse, err error) {
    if request == nil {
        request = NewDescribeRouteTablesRequest()
    }
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

// 获取子网内可用IP列表
func (c *Client) DescribeSubnetAvailableIps(request *DescribeSubnetAvailableIpsRequest) (response *DescribeSubnetAvailableIpsResponse, err error) {
    if request == nil {
        request = NewDescribeSubnetAvailableIpsRequest()
    }
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

// 物理机可以加入物理机子网，虚拟子网，DOCKER子网，通过此接口可以查询物理机加入的子网。
func (c *Client) DescribeSubnetByDevice(request *DescribeSubnetByDeviceRequest) (response *DescribeSubnetByDeviceResponse, err error) {
    if request == nil {
        request = NewDescribeSubnetByDeviceRequest()
    }
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

// 托管可以加入物理机子网，虚拟子网，DOCKER子网，通过此接口可以查询托管加入的子网。
func (c *Client) DescribeSubnetByHostedDevice(request *DescribeSubnetByHostedDeviceRequest) (response *DescribeSubnetByHostedDeviceResponse, err error) {
    if request == nil {
        request = NewDescribeSubnetByHostedDeviceRequest()
    }
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

// 本接口（DescribeSubnets）用于查询黑石子网列表。
func (c *Client) DescribeSubnets(request *DescribeSubnetsRequest) (response *DescribeSubnetsResponse, err error) {
    if request == nil {
        request = NewDescribeSubnetsRequest()
    }
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

// 根据任务ID，获取任务的执行状态
func (c *Client) DescribeTaskStatus(request *DescribeTaskStatusRequest) (response *DescribeTaskStatusResponse, err error) {
    if request == nil {
        request = NewDescribeTaskStatusRequest()
    }
    response = NewDescribeTaskStatusResponse()
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

// 本接口（DescribeVpcQuota）用于查询用户VPC相关配额限制。
func (c *Client) DescribeVpcQuota(request *DescribeVpcQuotaRequest) (response *DescribeVpcQuotaResponse, err error) {
    if request == nil {
        request = NewDescribeVpcQuotaRequest()
    }
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

// 查询黑石私有网络关联资源
func (c *Client) DescribeVpcResource(request *DescribeVpcResourceRequest) (response *DescribeVpcResourceResponse, err error) {
    if request == nil {
        request = NewDescribeVpcResourceRequest()
    }
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

// 本接口（DescribeVpcView）用于查询VPC网络拓扑视图。
func (c *Client) DescribeVpcView(request *DescribeVpcViewRequest) (response *DescribeVpcViewResponse, err error) {
    if request == nil {
        request = NewDescribeVpcViewRequest()
    }
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

// 本接口（DescribeVpcs）用于查询私有网络列表。
// 本接口不传参数时，返回默认排序下的前20条VPC信息。
func (c *Client) DescribeVpcs(request *DescribeVpcsRequest) (response *DescribeVpcsResponse, err error) {
    if request == nil {
        request = NewDescribeVpcsRequest()
    }
    response = NewDescribeVpcsResponse()
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

// 本接口(DownloadCustomerGatewayConfiguration)用于下载VPN通道配置。
func (c *Client) DownloadCustomerGatewayConfiguration(request *DownloadCustomerGatewayConfigurationRequest) (response *DownloadCustomerGatewayConfigurationResponse, err error) {
    if request == nil {
        request = NewDownloadCustomerGatewayConfigurationRequest()
    }
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

// 本接口（ModifyCustomerGatewayAttribute）用于修改对端网关信息。
func (c *Client) ModifyCustomerGatewayAttribute(request *ModifyCustomerGatewayAttributeRequest) (response *ModifyCustomerGatewayAttributeResponse, err error) {
    if request == nil {
        request = NewModifyCustomerGatewayAttributeRequest()
    }
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

// 修改自定义路由
func (c *Client) ModifyRoutePolicy(request *ModifyRoutePolicyRequest) (response *ModifyRoutePolicyResponse, err error) {
    if request == nil {
        request = NewModifyRoutePolicyRequest()
    }
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

// 修改路由表
func (c *Client) ModifyRouteTable(request *ModifyRouteTableRequest) (response *ModifyRouteTableResponse, err error) {
    if request == nil {
        request = NewModifyRouteTableRequest()
    }
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

// 修改子网属性
func (c *Client) ModifySubnetAttribute(request *ModifySubnetAttributeRequest) (response *ModifySubnetAttributeResponse, err error) {
    if request == nil {
        request = NewModifySubnetAttributeRequest()
    }
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

// 修改子网DHCP Relay属性
func (c *Client) ModifySubnetDHCPRelay(request *ModifySubnetDHCPRelayRequest) (response *ModifySubnetDHCPRelayResponse, err error) {
    if request == nil {
        request = NewModifySubnetDHCPRelayRequest()
    }
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

// 本接口（ModifyVpcAttribute）用于修改VPC的标识名称和控制VPC的监控起停。
func (c *Client) ModifyVpcAttribute(request *ModifyVpcAttributeRequest) (response *ModifyVpcAttributeResponse, err error) {
    if request == nil {
        request = NewModifyVpcAttributeRequest()
    }
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

// 修改黑石对等连接
func (c *Client) ModifyVpcPeerConnection(request *ModifyVpcPeerConnectionRequest) (response *ModifyVpcPeerConnectionResponse, err error) {
    if request == nil {
        request = NewModifyVpcPeerConnectionRequest()
    }
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

// 本接口（ModifyVpnConnectionAttribute）用于修改VPN通道。
func (c *Client) ModifyVpnConnectionAttribute(request *ModifyVpnConnectionAttributeRequest) (response *ModifyVpnConnectionAttributeResponse, err error) {
    if request == nil {
        request = NewModifyVpnConnectionAttributeRequest()
    }
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

// 本接口（ModifyVpnGatewayAttribute）用于修改VPN网关属性。
func (c *Client) ModifyVpnGatewayAttribute(request *ModifyVpnGatewayAttributeRequest) (response *ModifyVpnGatewayAttributeResponse, err error) {
    if request == nil {
        request = NewModifyVpnGatewayAttributeRequest()
    }
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

// 拒绝黑石对等连接申请
func (c *Client) RejectVpcPeerConnection(request *RejectVpcPeerConnectionRequest) (response *RejectVpcPeerConnectionResponse, err error) {
    if request == nil {
        request = NewRejectVpcPeerConnectionRequest()
    }
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

// 本接口(ResetVpnConnection)用于重置VPN通道。
func (c *Client) ResetVpnConnection(request *ResetVpnConnectionRequest) (response *ResetVpnConnectionResponse, err error) {
    if request == nil {
        request = NewResetVpnConnectionRequest()
    }
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

// NAT网关解绑该EIP后，NAT网关将不会使用该EIP作为访问外网的源IP地址
func (c *Client) UnbindEipsFromNatGateway(request *UnbindEipsFromNatGatewayRequest) (response *UnbindEipsFromNatGatewayResponse, err error) {
    if request == nil {
        request = NewUnbindEipsFromNatGatewayRequest()
    }
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

// NAT网关解绑IP接口，可将子网的部分IP从NAT网关中解绑
func (c *Client) UnbindIpsFromNatGateway(request *UnbindIpsFromNatGatewayRequest) (response *UnbindIpsFromNatGatewayResponse, err error) {
    if request == nil {
        request = NewUnbindIpsFromNatGatewayRequest()
    }
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

// NAT网关解绑子网接口，可将子网解绑NAT网关
func (c *Client) UnbindSubnetsFromNatGateway(request *UnbindSubnetsFromNatGatewayRequest) (response *UnbindSubnetsFromNatGatewayResponse, err error) {
    if request == nil {
        request = NewUnbindSubnetsFromNatGatewayRequest()
    }
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

// 升级NAT网关接口，可NAT网关修改为小型NAT网关、中型NAT网关、以及大型NAT网关
func (c *Client) UpgradeNatGateway(request *UpgradeNatGatewayRequest) (response *UpgradeNatGatewayResponse, err error) {
    if request == nil {
        request = NewUpgradeNatGatewayRequest()
    }
    response = NewUpgradeNatGatewayResponse()
    err = c.Send(request, response)
    return
}
