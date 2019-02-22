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


func NewBindL4BackendsRequest() (request *BindL4BackendsRequest) {
    request = &BindL4BackendsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bmlb", APIVersion, "BindL4Backends")
    return
}

func NewBindL4BackendsResponse() (response *BindL4BackendsResponse) {
    response = &BindL4BackendsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 绑定黑石服务器到四层监听器。服务器包括物理服务器、虚拟机以及半托管机器。
func (c *Client) BindL4Backends(request *BindL4BackendsRequest) (response *BindL4BackendsResponse, err error) {
    if request == nil {
        request = NewBindL4BackendsRequest()
    }
    response = NewBindL4BackendsResponse()
    err = c.Send(request, response)
    return
}

func NewBindL7BackendsRequest() (request *BindL7BackendsRequest) {
    request = &BindL7BackendsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bmlb", APIVersion, "BindL7Backends")
    return
}

func NewBindL7BackendsResponse() (response *BindL7BackendsResponse) {
    response = &BindL7BackendsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 绑定黑石物理服务器或半托管服务器到七层转发路径。
func (c *Client) BindL7Backends(request *BindL7BackendsRequest) (response *BindL7BackendsResponse, err error) {
    if request == nil {
        request = NewBindL7BackendsRequest()
    }
    response = NewBindL7BackendsResponse()
    err = c.Send(request, response)
    return
}

func NewBindTrafficMirrorListenersRequest() (request *BindTrafficMirrorListenersRequest) {
    request = &BindTrafficMirrorListenersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bmlb", APIVersion, "BindTrafficMirrorListeners")
    return
}

func NewBindTrafficMirrorListenersResponse() (response *BindTrafficMirrorListenersResponse) {
    response = &BindTrafficMirrorListenersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 绑定黑石服务器七层监听器到流量镜像实例。
func (c *Client) BindTrafficMirrorListeners(request *BindTrafficMirrorListenersRequest) (response *BindTrafficMirrorListenersResponse, err error) {
    if request == nil {
        request = NewBindTrafficMirrorListenersRequest()
    }
    response = NewBindTrafficMirrorListenersResponse()
    err = c.Send(request, response)
    return
}

func NewBindTrafficMirrorReceiversRequest() (request *BindTrafficMirrorReceiversRequest) {
    request = &BindTrafficMirrorReceiversRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bmlb", APIVersion, "BindTrafficMirrorReceivers")
    return
}

func NewBindTrafficMirrorReceiversResponse() (response *BindTrafficMirrorReceiversResponse) {
    response = &BindTrafficMirrorReceiversResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 绑定黑石物理服务器成为流量镜像接收机。
func (c *Client) BindTrafficMirrorReceivers(request *BindTrafficMirrorReceiversRequest) (response *BindTrafficMirrorReceiversResponse, err error) {
    if request == nil {
        request = NewBindTrafficMirrorReceiversRequest()
    }
    response = NewBindTrafficMirrorReceiversResponse()
    err = c.Send(request, response)
    return
}

func NewCreateL4ListenersRequest() (request *CreateL4ListenersRequest) {
    request = &CreateL4ListenersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bmlb", APIVersion, "CreateL4Listeners")
    return
}

func NewCreateL4ListenersResponse() (response *CreateL4ListenersResponse) {
    response = &CreateL4ListenersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建黑石四层负载均衡监听器功能。黑石负载均衡四层监听器提供了转发用户请求的具体规则，包括端口、协议、会话保持、健康检查等参数。
func (c *Client) CreateL4Listeners(request *CreateL4ListenersRequest) (response *CreateL4ListenersResponse, err error) {
    if request == nil {
        request = NewCreateL4ListenersRequest()
    }
    response = NewCreateL4ListenersResponse()
    err = c.Send(request, response)
    return
}

func NewCreateL7ListenersRequest() (request *CreateL7ListenersRequest) {
    request = &CreateL7ListenersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bmlb", APIVersion, "CreateL7Listeners")
    return
}

func NewCreateL7ListenersResponse() (response *CreateL7ListenersResponse) {
    response = &CreateL7ListenersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建黑石负载均衡七层监听器功能。负载均衡七层监听器提供了转发用户请求的具体规则，包括端口、协议等参数。
func (c *Client) CreateL7Listeners(request *CreateL7ListenersRequest) (response *CreateL7ListenersResponse, err error) {
    if request == nil {
        request = NewCreateL7ListenersRequest()
    }
    response = NewCreateL7ListenersResponse()
    err = c.Send(request, response)
    return
}

func NewCreateL7RulesRequest() (request *CreateL7RulesRequest) {
    request = &CreateL7RulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bmlb", APIVersion, "CreateL7Rules")
    return
}

func NewCreateL7RulesResponse() (response *CreateL7RulesResponse) {
    response = &CreateL7RulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建黑石负载均衡七层转发规则。
func (c *Client) CreateL7Rules(request *CreateL7RulesRequest) (response *CreateL7RulesResponse, err error) {
    if request == nil {
        request = NewCreateL7RulesRequest()
    }
    response = NewCreateL7RulesResponse()
    err = c.Send(request, response)
    return
}

func NewCreateLoadBalancersRequest() (request *CreateLoadBalancersRequest) {
    request = &CreateLoadBalancersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bmlb", APIVersion, "CreateLoadBalancers")
    return
}

func NewCreateLoadBalancersResponse() (response *CreateLoadBalancersResponse) {
    response = &CreateLoadBalancersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 用来创建黑石负载均衡。为了使用黑石负载均衡服务，您必须要创建一个或者多个负载均衡实例。通过成功调用该接口，会返回负载均衡实例的唯一ID。用户可以购买的黑石负载均衡实例类型分为：公网类型、内网类型。公网类型负载均衡对应一个BGP VIP，可用于快速访问公网负载均衡绑定的物理服务器；内网类型负载均衡对应一个腾讯云内部的VIP，不能通过Internet访问，可快速访问内网负载均衡绑定的物理服务器。
func (c *Client) CreateLoadBalancers(request *CreateLoadBalancersRequest) (response *CreateLoadBalancersResponse, err error) {
    if request == nil {
        request = NewCreateLoadBalancersRequest()
    }
    response = NewCreateLoadBalancersResponse()
    err = c.Send(request, response)
    return
}

func NewCreateTrafficMirrorRequest() (request *CreateTrafficMirrorRequest) {
    request = &CreateTrafficMirrorRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bmlb", APIVersion, "CreateTrafficMirror")
    return
}

func NewCreateTrafficMirrorResponse() (response *CreateTrafficMirrorResponse) {
    response = &CreateTrafficMirrorResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建流量镜像实例。
func (c *Client) CreateTrafficMirror(request *CreateTrafficMirrorRequest) (response *CreateTrafficMirrorResponse, err error) {
    if request == nil {
        request = NewCreateTrafficMirrorRequest()
    }
    response = NewCreateTrafficMirrorResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteL7DomainsRequest() (request *DeleteL7DomainsRequest) {
    request = &DeleteL7DomainsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bmlb", APIVersion, "DeleteL7Domains")
    return
}

func NewDeleteL7DomainsResponse() (response *DeleteL7DomainsResponse) {
    response = &DeleteL7DomainsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除黑石负载均衡七层转发域名。
func (c *Client) DeleteL7Domains(request *DeleteL7DomainsRequest) (response *DeleteL7DomainsResponse, err error) {
    if request == nil {
        request = NewDeleteL7DomainsRequest()
    }
    response = NewDeleteL7DomainsResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteL7RulesRequest() (request *DeleteL7RulesRequest) {
    request = &DeleteL7RulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bmlb", APIVersion, "DeleteL7Rules")
    return
}

func NewDeleteL7RulesResponse() (response *DeleteL7RulesResponse) {
    response = &DeleteL7RulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除黑石负载均衡七层转发规则。
func (c *Client) DeleteL7Rules(request *DeleteL7RulesRequest) (response *DeleteL7RulesResponse, err error) {
    if request == nil {
        request = NewDeleteL7RulesRequest()
    }
    response = NewDeleteL7RulesResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteListenersRequest() (request *DeleteListenersRequest) {
    request = &DeleteListenersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bmlb", APIVersion, "DeleteListeners")
    return
}

func NewDeleteListenersResponse() (response *DeleteListenersResponse) {
    response = &DeleteListenersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除黑石负载均衡监听器。
func (c *Client) DeleteListeners(request *DeleteListenersRequest) (response *DeleteListenersResponse, err error) {
    if request == nil {
        request = NewDeleteListenersRequest()
    }
    response = NewDeleteListenersResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteLoadBalancerRequest() (request *DeleteLoadBalancerRequest) {
    request = &DeleteLoadBalancerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bmlb", APIVersion, "DeleteLoadBalancer")
    return
}

func NewDeleteLoadBalancerResponse() (response *DeleteLoadBalancerResponse) {
    response = &DeleteLoadBalancerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除用户指定的黑石负载均衡实例。
func (c *Client) DeleteLoadBalancer(request *DeleteLoadBalancerRequest) (response *DeleteLoadBalancerResponse, err error) {
    if request == nil {
        request = NewDeleteLoadBalancerRequest()
    }
    response = NewDeleteLoadBalancerResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteTrafficMirrorRequest() (request *DeleteTrafficMirrorRequest) {
    request = &DeleteTrafficMirrorRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bmlb", APIVersion, "DeleteTrafficMirror")
    return
}

func NewDeleteTrafficMirrorResponse() (response *DeleteTrafficMirrorResponse) {
    response = &DeleteTrafficMirrorResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除已创建的黑石流量镜像实例，删除过程是异步执行的，因此需要使用查询任务接口获取删除的结果。
func (c *Client) DeleteTrafficMirror(request *DeleteTrafficMirrorRequest) (response *DeleteTrafficMirrorResponse, err error) {
    if request == nil {
        request = NewDeleteTrafficMirrorRequest()
    }
    response = NewDeleteTrafficMirrorResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCertDetailRequest() (request *DescribeCertDetailRequest) {
    request = &DescribeCertDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bmlb", APIVersion, "DescribeCertDetail")
    return
}

func NewDescribeCertDetailResponse() (response *DescribeCertDetailResponse) {
    response = &DescribeCertDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取黑石负载均衡证书详情。
func (c *Client) DescribeCertDetail(request *DescribeCertDetailRequest) (response *DescribeCertDetailResponse, err error) {
    if request == nil {
        request = NewDescribeCertDetailRequest()
    }
    response = NewDescribeCertDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDevicesBindInfoRequest() (request *DescribeDevicesBindInfoRequest) {
    request = &DescribeDevicesBindInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bmlb", APIVersion, "DescribeDevicesBindInfo")
    return
}

func NewDescribeDevicesBindInfoResponse() (response *DescribeDevicesBindInfoResponse) {
    response = &DescribeDevicesBindInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询黑石物理机和虚机以及托管服务器绑定的黑石负载均衡详情。
func (c *Client) DescribeDevicesBindInfo(request *DescribeDevicesBindInfoRequest) (response *DescribeDevicesBindInfoResponse, err error) {
    if request == nil {
        request = NewDescribeDevicesBindInfoRequest()
    }
    response = NewDescribeDevicesBindInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeL4BackendsRequest() (request *DescribeL4BackendsRequest) {
    request = &DescribeL4BackendsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bmlb", APIVersion, "DescribeL4Backends")
    return
}

func NewDescribeL4BackendsResponse() (response *DescribeL4BackendsResponse) {
    response = &DescribeL4BackendsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取黑石负载均衡四层监听器绑定的主机列表。
func (c *Client) DescribeL4Backends(request *DescribeL4BackendsRequest) (response *DescribeL4BackendsResponse, err error) {
    if request == nil {
        request = NewDescribeL4BackendsRequest()
    }
    response = NewDescribeL4BackendsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeL4ListenerInfoRequest() (request *DescribeL4ListenerInfoRequest) {
    request = &DescribeL4ListenerInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bmlb", APIVersion, "DescribeL4ListenerInfo")
    return
}

func NewDescribeL4ListenerInfoResponse() (response *DescribeL4ListenerInfoResponse) {
    response = &DescribeL4ListenerInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查找绑定了某主机或者指定监听器名称的黑石负载均衡四层监听器。
func (c *Client) DescribeL4ListenerInfo(request *DescribeL4ListenerInfoRequest) (response *DescribeL4ListenerInfoResponse, err error) {
    if request == nil {
        request = NewDescribeL4ListenerInfoRequest()
    }
    response = NewDescribeL4ListenerInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeL4ListenersRequest() (request *DescribeL4ListenersRequest) {
    request = &DescribeL4ListenersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bmlb", APIVersion, "DescribeL4Listeners")
    return
}

func NewDescribeL4ListenersResponse() (response *DescribeL4ListenersResponse) {
    response = &DescribeL4ListenersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取黑石负载均衡四层监听器。
func (c *Client) DescribeL4Listeners(request *DescribeL4ListenersRequest) (response *DescribeL4ListenersResponse, err error) {
    if request == nil {
        request = NewDescribeL4ListenersRequest()
    }
    response = NewDescribeL4ListenersResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeL7BackendsRequest() (request *DescribeL7BackendsRequest) {
    request = &DescribeL7BackendsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bmlb", APIVersion, "DescribeL7Backends")
    return
}

func NewDescribeL7BackendsResponse() (response *DescribeL7BackendsResponse) {
    response = &DescribeL7BackendsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取黑石负载均衡七层监听器绑定的主机列表
func (c *Client) DescribeL7Backends(request *DescribeL7BackendsRequest) (response *DescribeL7BackendsResponse, err error) {
    if request == nil {
        request = NewDescribeL7BackendsRequest()
    }
    response = NewDescribeL7BackendsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeL7ListenerInfoRequest() (request *DescribeL7ListenerInfoRequest) {
    request = &DescribeL7ListenerInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bmlb", APIVersion, "DescribeL7ListenerInfo")
    return
}

func NewDescribeL7ListenerInfoResponse() (response *DescribeL7ListenerInfoResponse) {
    response = &DescribeL7ListenerInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查找绑定了某主机或者有某转发域名黑石负载均衡七层监听器。
func (c *Client) DescribeL7ListenerInfo(request *DescribeL7ListenerInfoRequest) (response *DescribeL7ListenerInfoResponse, err error) {
    if request == nil {
        request = NewDescribeL7ListenerInfoRequest()
    }
    response = NewDescribeL7ListenerInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeL7ListenersRequest() (request *DescribeL7ListenersRequest) {
    request = &DescribeL7ListenersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bmlb", APIVersion, "DescribeL7Listeners")
    return
}

func NewDescribeL7ListenersResponse() (response *DescribeL7ListenersResponse) {
    response = &DescribeL7ListenersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取黑石负载均衡七层监听器列表信息。
func (c *Client) DescribeL7Listeners(request *DescribeL7ListenersRequest) (response *DescribeL7ListenersResponse, err error) {
    if request == nil {
        request = NewDescribeL7ListenersRequest()
    }
    response = NewDescribeL7ListenersResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeL7ListenersExRequest() (request *DescribeL7ListenersExRequest) {
    request = &DescribeL7ListenersExRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bmlb", APIVersion, "DescribeL7ListenersEx")
    return
}

func NewDescribeL7ListenersExResponse() (response *DescribeL7ListenersExResponse) {
    response = &DescribeL7ListenersExResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取指定VPC下的7层监听器(支持模糊匹配)。
func (c *Client) DescribeL7ListenersEx(request *DescribeL7ListenersExRequest) (response *DescribeL7ListenersExResponse, err error) {
    if request == nil {
        request = NewDescribeL7ListenersExRequest()
    }
    response = NewDescribeL7ListenersExResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeL7RulesRequest() (request *DescribeL7RulesRequest) {
    request = &DescribeL7RulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bmlb", APIVersion, "DescribeL7Rules")
    return
}

func NewDescribeL7RulesResponse() (response *DescribeL7RulesResponse) {
    response = &DescribeL7RulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取黑石负载均衡七层转发规则。
func (c *Client) DescribeL7Rules(request *DescribeL7RulesRequest) (response *DescribeL7RulesResponse, err error) {
    if request == nil {
        request = NewDescribeL7RulesRequest()
    }
    response = NewDescribeL7RulesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLoadBalancerPortInfoRequest() (request *DescribeLoadBalancerPortInfoRequest) {
    request = &DescribeLoadBalancerPortInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bmlb", APIVersion, "DescribeLoadBalancerPortInfo")
    return
}

func NewDescribeLoadBalancerPortInfoResponse() (response *DescribeLoadBalancerPortInfoResponse) {
    response = &DescribeLoadBalancerPortInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取黑石负载均衡端口相关信息。
func (c *Client) DescribeLoadBalancerPortInfo(request *DescribeLoadBalancerPortInfoRequest) (response *DescribeLoadBalancerPortInfoResponse, err error) {
    if request == nil {
        request = NewDescribeLoadBalancerPortInfoRequest()
    }
    response = NewDescribeLoadBalancerPortInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLoadBalancerTaskResultRequest() (request *DescribeLoadBalancerTaskResultRequest) {
    request = &DescribeLoadBalancerTaskResultRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bmlb", APIVersion, "DescribeLoadBalancerTaskResult")
    return
}

func NewDescribeLoadBalancerTaskResultResponse() (response *DescribeLoadBalancerTaskResultResponse) {
    response = &DescribeLoadBalancerTaskResultResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询负载均衡实例异步任务的执行情况。
func (c *Client) DescribeLoadBalancerTaskResult(request *DescribeLoadBalancerTaskResultRequest) (response *DescribeLoadBalancerTaskResultResponse, err error) {
    if request == nil {
        request = NewDescribeLoadBalancerTaskResultRequest()
    }
    response = NewDescribeLoadBalancerTaskResultResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLoadBalancersRequest() (request *DescribeLoadBalancersRequest) {
    request = &DescribeLoadBalancersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bmlb", APIVersion, "DescribeLoadBalancers")
    return
}

func NewDescribeLoadBalancersResponse() (response *DescribeLoadBalancersResponse) {
    response = &DescribeLoadBalancersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取黑石负载均衡实例列表
func (c *Client) DescribeLoadBalancers(request *DescribeLoadBalancersRequest) (response *DescribeLoadBalancersResponse, err error) {
    if request == nil {
        request = NewDescribeLoadBalancersRequest()
    }
    response = NewDescribeLoadBalancersResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTrafficMirrorListenersRequest() (request *DescribeTrafficMirrorListenersRequest) {
    request = &DescribeTrafficMirrorListenersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bmlb", APIVersion, "DescribeTrafficMirrorListeners")
    return
}

func NewDescribeTrafficMirrorListenersResponse() (response *DescribeTrafficMirrorListenersResponse) {
    response = &DescribeTrafficMirrorListenersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取流量镜像的监听器列表信息。
func (c *Client) DescribeTrafficMirrorListeners(request *DescribeTrafficMirrorListenersRequest) (response *DescribeTrafficMirrorListenersResponse, err error) {
    if request == nil {
        request = NewDescribeTrafficMirrorListenersRequest()
    }
    response = NewDescribeTrafficMirrorListenersResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTrafficMirrorReceiverHealthStatusRequest() (request *DescribeTrafficMirrorReceiverHealthStatusRequest) {
    request = &DescribeTrafficMirrorReceiverHealthStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bmlb", APIVersion, "DescribeTrafficMirrorReceiverHealthStatus")
    return
}

func NewDescribeTrafficMirrorReceiverHealthStatusResponse() (response *DescribeTrafficMirrorReceiverHealthStatusResponse) {
    response = &DescribeTrafficMirrorReceiverHealthStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取流量镜像接收机健康状态。
func (c *Client) DescribeTrafficMirrorReceiverHealthStatus(request *DescribeTrafficMirrorReceiverHealthStatusRequest) (response *DescribeTrafficMirrorReceiverHealthStatusResponse, err error) {
    if request == nil {
        request = NewDescribeTrafficMirrorReceiverHealthStatusRequest()
    }
    response = NewDescribeTrafficMirrorReceiverHealthStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTrafficMirrorReceiversRequest() (request *DescribeTrafficMirrorReceiversRequest) {
    request = &DescribeTrafficMirrorReceiversRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bmlb", APIVersion, "DescribeTrafficMirrorReceivers")
    return
}

func NewDescribeTrafficMirrorReceiversResponse() (response *DescribeTrafficMirrorReceiversResponse) {
    response = &DescribeTrafficMirrorReceiversResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取指定流量镜像实例的接收机信息。
func (c *Client) DescribeTrafficMirrorReceivers(request *DescribeTrafficMirrorReceiversRequest) (response *DescribeTrafficMirrorReceiversResponse, err error) {
    if request == nil {
        request = NewDescribeTrafficMirrorReceiversRequest()
    }
    response = NewDescribeTrafficMirrorReceiversResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTrafficMirrorsRequest() (request *DescribeTrafficMirrorsRequest) {
    request = &DescribeTrafficMirrorsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bmlb", APIVersion, "DescribeTrafficMirrors")
    return
}

func NewDescribeTrafficMirrorsResponse() (response *DescribeTrafficMirrorsResponse) {
    response = &DescribeTrafficMirrorsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取流量镜像实例的列表信息。
func (c *Client) DescribeTrafficMirrors(request *DescribeTrafficMirrorsRequest) (response *DescribeTrafficMirrorsResponse, err error) {
    if request == nil {
        request = NewDescribeTrafficMirrorsRequest()
    }
    response = NewDescribeTrafficMirrorsResponse()
    err = c.Send(request, response)
    return
}

func NewModifyL4BackendPortRequest() (request *ModifyL4BackendPortRequest) {
    request = &ModifyL4BackendPortRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bmlb", APIVersion, "ModifyL4BackendPort")
    return
}

func NewModifyL4BackendPortResponse() (response *ModifyL4BackendPortResponse) {
    response = &ModifyL4BackendPortResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改黑石负载均衡四层监听器后端实例端口。
func (c *Client) ModifyL4BackendPort(request *ModifyL4BackendPortRequest) (response *ModifyL4BackendPortResponse, err error) {
    if request == nil {
        request = NewModifyL4BackendPortRequest()
    }
    response = NewModifyL4BackendPortResponse()
    err = c.Send(request, response)
    return
}

func NewModifyL4BackendProbePortRequest() (request *ModifyL4BackendProbePortRequest) {
    request = &ModifyL4BackendProbePortRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bmlb", APIVersion, "ModifyL4BackendProbePort")
    return
}

func NewModifyL4BackendProbePortResponse() (response *ModifyL4BackendProbePortResponse) {
    response = &ModifyL4BackendProbePortResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改黑石负载均衡四层监听器后端探测端口。
func (c *Client) ModifyL4BackendProbePort(request *ModifyL4BackendProbePortRequest) (response *ModifyL4BackendProbePortResponse, err error) {
    if request == nil {
        request = NewModifyL4BackendProbePortRequest()
    }
    response = NewModifyL4BackendProbePortResponse()
    err = c.Send(request, response)
    return
}

func NewModifyL4BackendWeightRequest() (request *ModifyL4BackendWeightRequest) {
    request = &ModifyL4BackendWeightRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bmlb", APIVersion, "ModifyL4BackendWeight")
    return
}

func NewModifyL4BackendWeightResponse() (response *ModifyL4BackendWeightResponse) {
    response = &ModifyL4BackendWeightResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改黑石负载均衡四层监听器后端实例权重功能。
func (c *Client) ModifyL4BackendWeight(request *ModifyL4BackendWeightRequest) (response *ModifyL4BackendWeightResponse, err error) {
    if request == nil {
        request = NewModifyL4BackendWeightRequest()
    }
    response = NewModifyL4BackendWeightResponse()
    err = c.Send(request, response)
    return
}

func NewModifyL4ListenerRequest() (request *ModifyL4ListenerRequest) {
    request = &ModifyL4ListenerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bmlb", APIVersion, "ModifyL4Listener")
    return
}

func NewModifyL4ListenerResponse() (response *ModifyL4ListenerResponse) {
    response = &ModifyL4ListenerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改黑石负载均衡四层监听器。
func (c *Client) ModifyL4Listener(request *ModifyL4ListenerRequest) (response *ModifyL4ListenerResponse, err error) {
    if request == nil {
        request = NewModifyL4ListenerRequest()
    }
    response = NewModifyL4ListenerResponse()
    err = c.Send(request, response)
    return
}

func NewModifyL7BackendPortRequest() (request *ModifyL7BackendPortRequest) {
    request = &ModifyL7BackendPortRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bmlb", APIVersion, "ModifyL7BackendPort")
    return
}

func NewModifyL7BackendPortResponse() (response *ModifyL7BackendPortResponse) {
    response = &ModifyL7BackendPortResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改黑石负载均衡七层转发路径后端实例端口。
func (c *Client) ModifyL7BackendPort(request *ModifyL7BackendPortRequest) (response *ModifyL7BackendPortResponse, err error) {
    if request == nil {
        request = NewModifyL7BackendPortRequest()
    }
    response = NewModifyL7BackendPortResponse()
    err = c.Send(request, response)
    return
}

func NewModifyL7BackendWeightRequest() (request *ModifyL7BackendWeightRequest) {
    request = &ModifyL7BackendWeightRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bmlb", APIVersion, "ModifyL7BackendWeight")
    return
}

func NewModifyL7BackendWeightResponse() (response *ModifyL7BackendWeightResponse) {
    response = &ModifyL7BackendWeightResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改黑石负载均衡七层转发路径后端实例权重。
func (c *Client) ModifyL7BackendWeight(request *ModifyL7BackendWeightRequest) (response *ModifyL7BackendWeightResponse, err error) {
    if request == nil {
        request = NewModifyL7BackendWeightRequest()
    }
    response = NewModifyL7BackendWeightResponse()
    err = c.Send(request, response)
    return
}

func NewModifyL7ListenerRequest() (request *ModifyL7ListenerRequest) {
    request = &ModifyL7ListenerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bmlb", APIVersion, "ModifyL7Listener")
    return
}

func NewModifyL7ListenerResponse() (response *ModifyL7ListenerResponse) {
    response = &ModifyL7ListenerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改黑石负载均衡七层监听器。
func (c *Client) ModifyL7Listener(request *ModifyL7ListenerRequest) (response *ModifyL7ListenerResponse, err error) {
    if request == nil {
        request = NewModifyL7ListenerRequest()
    }
    response = NewModifyL7ListenerResponse()
    err = c.Send(request, response)
    return
}

func NewModifyL7LocationsRequest() (request *ModifyL7LocationsRequest) {
    request = &ModifyL7LocationsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bmlb", APIVersion, "ModifyL7Locations")
    return
}

func NewModifyL7LocationsResponse() (response *ModifyL7LocationsResponse) {
    response = &ModifyL7LocationsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改黑石负载均衡七层转发路径。
func (c *Client) ModifyL7Locations(request *ModifyL7LocationsRequest) (response *ModifyL7LocationsResponse, err error) {
    if request == nil {
        request = NewModifyL7LocationsRequest()
    }
    response = NewModifyL7LocationsResponse()
    err = c.Send(request, response)
    return
}

func NewModifyLoadBalancerRequest() (request *ModifyLoadBalancerRequest) {
    request = &ModifyLoadBalancerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bmlb", APIVersion, "ModifyLoadBalancer")
    return
}

func NewModifyLoadBalancerResponse() (response *ModifyLoadBalancerResponse) {
    response = &ModifyLoadBalancerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 根据输入参数来修改黑石负载均衡实例的基本配置信息。可能的信息包括负载均衡实例的名称，域名前缀。
func (c *Client) ModifyLoadBalancer(request *ModifyLoadBalancerRequest) (response *ModifyLoadBalancerResponse, err error) {
    if request == nil {
        request = NewModifyLoadBalancerRequest()
    }
    response = NewModifyLoadBalancerResponse()
    err = c.Send(request, response)
    return
}

func NewModifyLoadBalancerChargeModeRequest() (request *ModifyLoadBalancerChargeModeRequest) {
    request = &ModifyLoadBalancerChargeModeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bmlb", APIVersion, "ModifyLoadBalancerChargeMode")
    return
}

func NewModifyLoadBalancerChargeModeResponse() (response *ModifyLoadBalancerChargeModeResponse) {
    response = &ModifyLoadBalancerChargeModeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 更改黑石负载均衡的计费方式
func (c *Client) ModifyLoadBalancerChargeMode(request *ModifyLoadBalancerChargeModeRequest) (response *ModifyLoadBalancerChargeModeResponse, err error) {
    if request == nil {
        request = NewModifyLoadBalancerChargeModeRequest()
    }
    response = NewModifyLoadBalancerChargeModeResponse()
    err = c.Send(request, response)
    return
}

func NewReplaceCertRequest() (request *ReplaceCertRequest) {
    request = &ReplaceCertRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bmlb", APIVersion, "ReplaceCert")
    return
}

func NewReplaceCertResponse() (response *ReplaceCertResponse) {
    response = &ReplaceCertResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 更新黑石负载均衡证书。
func (c *Client) ReplaceCert(request *ReplaceCertRequest) (response *ReplaceCertResponse, err error) {
    if request == nil {
        request = NewReplaceCertRequest()
    }
    response = NewReplaceCertResponse()
    err = c.Send(request, response)
    return
}

func NewSetTrafficMirrorAliasRequest() (request *SetTrafficMirrorAliasRequest) {
    request = &SetTrafficMirrorAliasRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bmlb", APIVersion, "SetTrafficMirrorAlias")
    return
}

func NewSetTrafficMirrorAliasResponse() (response *SetTrafficMirrorAliasResponse) {
    response = &SetTrafficMirrorAliasResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 设置流量镜像的别名。
func (c *Client) SetTrafficMirrorAlias(request *SetTrafficMirrorAliasRequest) (response *SetTrafficMirrorAliasResponse, err error) {
    if request == nil {
        request = NewSetTrafficMirrorAliasRequest()
    }
    response = NewSetTrafficMirrorAliasResponse()
    err = c.Send(request, response)
    return
}

func NewSetTrafficMirrorHealthSwitchRequest() (request *SetTrafficMirrorHealthSwitchRequest) {
    request = &SetTrafficMirrorHealthSwitchRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bmlb", APIVersion, "SetTrafficMirrorHealthSwitch")
    return
}

func NewSetTrafficMirrorHealthSwitchResponse() (response *SetTrafficMirrorHealthSwitchResponse) {
    response = &SetTrafficMirrorHealthSwitchResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 设置流量镜像的健康检查参数。
func (c *Client) SetTrafficMirrorHealthSwitch(request *SetTrafficMirrorHealthSwitchRequest) (response *SetTrafficMirrorHealthSwitchResponse, err error) {
    if request == nil {
        request = NewSetTrafficMirrorHealthSwitchRequest()
    }
    response = NewSetTrafficMirrorHealthSwitchResponse()
    err = c.Send(request, response)
    return
}

func NewUnbindL4BackendsRequest() (request *UnbindL4BackendsRequest) {
    request = &UnbindL4BackendsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bmlb", APIVersion, "UnbindL4Backends")
    return
}

func NewUnbindL4BackendsResponse() (response *UnbindL4BackendsResponse) {
    response = &UnbindL4BackendsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 解绑黑石负载均衡四层监听器物理服务器。
func (c *Client) UnbindL4Backends(request *UnbindL4BackendsRequest) (response *UnbindL4BackendsResponse, err error) {
    if request == nil {
        request = NewUnbindL4BackendsRequest()
    }
    response = NewUnbindL4BackendsResponse()
    err = c.Send(request, response)
    return
}

func NewUnbindL7BackendsRequest() (request *UnbindL7BackendsRequest) {
    request = &UnbindL7BackendsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bmlb", APIVersion, "UnbindL7Backends")
    return
}

func NewUnbindL7BackendsResponse() (response *UnbindL7BackendsResponse) {
    response = &UnbindL7BackendsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 解绑黑石物理服务器或者托管服务器到七层转发路径功能。
func (c *Client) UnbindL7Backends(request *UnbindL7BackendsRequest) (response *UnbindL7BackendsResponse, err error) {
    if request == nil {
        request = NewUnbindL7BackendsRequest()
    }
    response = NewUnbindL7BackendsResponse()
    err = c.Send(request, response)
    return
}

func NewUnbindTrafficMirrorListenersRequest() (request *UnbindTrafficMirrorListenersRequest) {
    request = &UnbindTrafficMirrorListenersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bmlb", APIVersion, "UnbindTrafficMirrorListeners")
    return
}

func NewUnbindTrafficMirrorListenersResponse() (response *UnbindTrafficMirrorListenersResponse) {
    response = &UnbindTrafficMirrorListenersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 解绑流量镜像监听器。
func (c *Client) UnbindTrafficMirrorListeners(request *UnbindTrafficMirrorListenersRequest) (response *UnbindTrafficMirrorListenersResponse, err error) {
    if request == nil {
        request = NewUnbindTrafficMirrorListenersRequest()
    }
    response = NewUnbindTrafficMirrorListenersResponse()
    err = c.Send(request, response)
    return
}

func NewUnbindTrafficMirrorReceiversRequest() (request *UnbindTrafficMirrorReceiversRequest) {
    request = &UnbindTrafficMirrorReceiversRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bmlb", APIVersion, "UnbindTrafficMirrorReceivers")
    return
}

func NewUnbindTrafficMirrorReceiversResponse() (response *UnbindTrafficMirrorReceiversResponse) {
    response = &UnbindTrafficMirrorReceiversResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 从流量镜像实例上解绑流量镜像接收机。
func (c *Client) UnbindTrafficMirrorReceivers(request *UnbindTrafficMirrorReceiversRequest) (response *UnbindTrafficMirrorReceiversResponse, err error) {
    if request == nil {
        request = NewUnbindTrafficMirrorReceiversRequest()
    }
    response = NewUnbindTrafficMirrorReceiversResponse()
    err = c.Send(request, response)
    return
}

func NewUploadCertRequest() (request *UploadCertRequest) {
    request = &UploadCertRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bmlb", APIVersion, "UploadCert")
    return
}

func NewUploadCertResponse() (response *UploadCertResponse) {
    response = &UploadCertResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建黑石负载均衡证书。
func (c *Client) UploadCert(request *UploadCertRequest) (response *UploadCertResponse, err error) {
    if request == nil {
        request = NewUploadCertRequest()
    }
    response = NewUploadCertResponse()
    err = c.Send(request, response)
    return
}
