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

package v20180529

import (
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2018-05-29"

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


func NewAddRealServersRequest() (request *AddRealServersRequest) {
    request = &AddRealServersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gaap", APIVersion, "AddRealServers")
    return
}

func NewAddRealServersResponse() (response *AddRealServersResponse) {
    response = &AddRealServersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 添加源站(服务器)信息，支持IP或域名
func (c *Client) AddRealServers(request *AddRealServersRequest) (response *AddRealServersResponse, err error) {
    if request == nil {
        request = NewAddRealServersRequest()
    }
    response = NewAddRealServersResponse()
    err = c.Send(request, response)
    return
}

func NewBindListenerRealServersRequest() (request *BindListenerRealServersRequest) {
    request = &BindListenerRealServersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gaap", APIVersion, "BindListenerRealServers")
    return
}

func NewBindListenerRealServersResponse() (response *BindListenerRealServersResponse) {
    response = &BindListenerRealServersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（BindListenerRealServers）用于TCP/UDP监听器绑定解绑源站。
// 注意：本接口会解绑之前绑定的源站，绑定本次调用所选择的源站。例如：原来绑定的源站为A，B，C，本次调用的选择绑定的源站为C，D，E，那么调用后所绑定的源站为C，D，E。
func (c *Client) BindListenerRealServers(request *BindListenerRealServersRequest) (response *BindListenerRealServersResponse, err error) {
    if request == nil {
        request = NewBindListenerRealServersRequest()
    }
    response = NewBindListenerRealServersResponse()
    err = c.Send(request, response)
    return
}

func NewBindRuleRealServersRequest() (request *BindRuleRealServersRequest) {
    request = &BindRuleRealServersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gaap", APIVersion, "BindRuleRealServers")
    return
}

func NewBindRuleRealServersResponse() (response *BindRuleRealServersResponse) {
    response = &BindRuleRealServersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 该接口用于7层监听器的转发规则绑定源站。注意：本接口会解绑之前绑定的源站，绑定本次调用所选择的源站。
func (c *Client) BindRuleRealServers(request *BindRuleRealServersRequest) (response *BindRuleRealServersResponse, err error) {
    if request == nil {
        request = NewBindRuleRealServersRequest()
    }
    response = NewBindRuleRealServersResponse()
    err = c.Send(request, response)
    return
}

func NewCheckProxyCreateRequest() (request *CheckProxyCreateRequest) {
    request = &CheckProxyCreateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gaap", APIVersion, "CheckProxyCreate")
    return
}

func NewCheckProxyCreateResponse() (response *CheckProxyCreateResponse) {
    response = &CheckProxyCreateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口(CheckProxyCreate)用于查询能否创建指定配置的加速通道。
func (c *Client) CheckProxyCreate(request *CheckProxyCreateRequest) (response *CheckProxyCreateResponse, err error) {
    if request == nil {
        request = NewCheckProxyCreateRequest()
    }
    response = NewCheckProxyCreateResponse()
    err = c.Send(request, response)
    return
}

func NewCloseProxiesRequest() (request *CloseProxiesRequest) {
    request = &CloseProxiesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gaap", APIVersion, "CloseProxies")
    return
}

func NewCloseProxiesResponse() (response *CloseProxiesResponse) {
    response = &CloseProxiesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（CloseProxies）用于关闭通道。通道关闭后，不再产生流量，但每天仍然收取通道基础配置费用。
func (c *Client) CloseProxies(request *CloseProxiesRequest) (response *CloseProxiesResponse, err error) {
    if request == nil {
        request = NewCloseProxiesRequest()
    }
    response = NewCloseProxiesResponse()
    err = c.Send(request, response)
    return
}

func NewCloseSecurityPolicyRequest() (request *CloseSecurityPolicyRequest) {
    request = &CloseSecurityPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gaap", APIVersion, "CloseSecurityPolicy")
    return
}

func NewCloseSecurityPolicyResponse() (response *CloseSecurityPolicyResponse) {
    response = &CloseSecurityPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 关闭安全策略
func (c *Client) CloseSecurityPolicy(request *CloseSecurityPolicyRequest) (response *CloseSecurityPolicyResponse, err error) {
    if request == nil {
        request = NewCloseSecurityPolicyRequest()
    }
    response = NewCloseSecurityPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewCreateCertificateRequest() (request *CreateCertificateRequest) {
    request = &CreateCertificateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gaap", APIVersion, "CreateCertificate")
    return
}

func NewCreateCertificateResponse() (response *CreateCertificateResponse) {
    response = &CreateCertificateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（CreateCertificate）用于创建Gaap相关证书和配置文件，包括基础认证配置文件，客户端CA证书，服务器SSL证书，Gaap SSL证书以及源站CA证书。
func (c *Client) CreateCertificate(request *CreateCertificateRequest) (response *CreateCertificateResponse, err error) {
    if request == nil {
        request = NewCreateCertificateRequest()
    }
    response = NewCreateCertificateResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDomainRequest() (request *CreateDomainRequest) {
    request = &CreateDomainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gaap", APIVersion, "CreateDomain")
    return
}

func NewCreateDomainResponse() (response *CreateDomainResponse) {
    response = &CreateDomainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（CreateDomain）用于创建HTTP/HTTPS监听器的访问域名，客户端请求通过访问该域名来请求后端业务。
// 该接口仅支持version3.0的通道。
func (c *Client) CreateDomain(request *CreateDomainRequest) (response *CreateDomainResponse, err error) {
    if request == nil {
        request = NewCreateDomainRequest()
    }
    response = NewCreateDomainResponse()
    err = c.Send(request, response)
    return
}

func NewCreateHTTPListenerRequest() (request *CreateHTTPListenerRequest) {
    request = &CreateHTTPListenerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gaap", APIVersion, "CreateHTTPListener")
    return
}

func NewCreateHTTPListenerResponse() (response *CreateHTTPListenerResponse) {
    response = &CreateHTTPListenerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 该接口（CreateHTTPListener）用于在通道实例下创建HTTP协议类型的监听器。
func (c *Client) CreateHTTPListener(request *CreateHTTPListenerRequest) (response *CreateHTTPListenerResponse, err error) {
    if request == nil {
        request = NewCreateHTTPListenerRequest()
    }
    response = NewCreateHTTPListenerResponse()
    err = c.Send(request, response)
    return
}

func NewCreateHTTPSListenerRequest() (request *CreateHTTPSListenerRequest) {
    request = &CreateHTTPSListenerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gaap", APIVersion, "CreateHTTPSListener")
    return
}

func NewCreateHTTPSListenerResponse() (response *CreateHTTPSListenerResponse) {
    response = &CreateHTTPSListenerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 该接口（CreateHTTPSListener）用于在通道实例下创建HTTPS协议类型的监听器。
func (c *Client) CreateHTTPSListener(request *CreateHTTPSListenerRequest) (response *CreateHTTPSListenerResponse, err error) {
    if request == nil {
        request = NewCreateHTTPSListenerRequest()
    }
    response = NewCreateHTTPSListenerResponse()
    err = c.Send(request, response)
    return
}

func NewCreateProxyRequest() (request *CreateProxyRequest) {
    request = &CreateProxyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gaap", APIVersion, "CreateProxy")
    return
}

func NewCreateProxyResponse() (response *CreateProxyResponse) {
    response = &CreateProxyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（CreateProxy）用于创建/复制一个指定配置的加速通道。当复制通道时，需要设置新通道的基本配置参数，并设置ClonedProxyId来指定被复制的通道。
func (c *Client) CreateProxy(request *CreateProxyRequest) (response *CreateProxyResponse, err error) {
    if request == nil {
        request = NewCreateProxyRequest()
    }
    response = NewCreateProxyResponse()
    err = c.Send(request, response)
    return
}

func NewCreateProxyGroupRequest() (request *CreateProxyGroupRequest) {
    request = &CreateProxyGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gaap", APIVersion, "CreateProxyGroup")
    return
}

func NewCreateProxyGroupResponse() (response *CreateProxyGroupResponse) {
    response = &CreateProxyGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（CreateProxyGroup）用于创建通道组。
func (c *Client) CreateProxyGroup(request *CreateProxyGroupRequest) (response *CreateProxyGroupResponse, err error) {
    if request == nil {
        request = NewCreateProxyGroupRequest()
    }
    response = NewCreateProxyGroupResponse()
    err = c.Send(request, response)
    return
}

func NewCreateProxyGroupDomainRequest() (request *CreateProxyGroupDomainRequest) {
    request = &CreateProxyGroupDomainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gaap", APIVersion, "CreateProxyGroupDomain")
    return
}

func NewCreateProxyGroupDomainResponse() (response *CreateProxyGroupDomainResponse) {
    response = &CreateProxyGroupDomainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（CreateProxyGroupDomain）用于创建通道组域名，并开启域名解析。
func (c *Client) CreateProxyGroupDomain(request *CreateProxyGroupDomainRequest) (response *CreateProxyGroupDomainResponse, err error) {
    if request == nil {
        request = NewCreateProxyGroupDomainRequest()
    }
    response = NewCreateProxyGroupDomainResponse()
    err = c.Send(request, response)
    return
}

func NewCreateRuleRequest() (request *CreateRuleRequest) {
    request = &CreateRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gaap", APIVersion, "CreateRule")
    return
}

func NewCreateRuleResponse() (response *CreateRuleResponse) {
    response = &CreateRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 该接口（CreateRule）用于创建HTTP/HTTPS监听器转发规则。
func (c *Client) CreateRule(request *CreateRuleRequest) (response *CreateRuleResponse, err error) {
    if request == nil {
        request = NewCreateRuleRequest()
    }
    response = NewCreateRuleResponse()
    err = c.Send(request, response)
    return
}

func NewCreateSecurityPolicyRequest() (request *CreateSecurityPolicyRequest) {
    request = &CreateSecurityPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gaap", APIVersion, "CreateSecurityPolicy")
    return
}

func NewCreateSecurityPolicyResponse() (response *CreateSecurityPolicyResponse) {
    response = &CreateSecurityPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建安全策略
func (c *Client) CreateSecurityPolicy(request *CreateSecurityPolicyRequest) (response *CreateSecurityPolicyResponse, err error) {
    if request == nil {
        request = NewCreateSecurityPolicyRequest()
    }
    response = NewCreateSecurityPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewCreateSecurityRulesRequest() (request *CreateSecurityRulesRequest) {
    request = &CreateSecurityRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gaap", APIVersion, "CreateSecurityRules")
    return
}

func NewCreateSecurityRulesResponse() (response *CreateSecurityRulesResponse) {
    response = &CreateSecurityRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 添加安全策略规则
func (c *Client) CreateSecurityRules(request *CreateSecurityRulesRequest) (response *CreateSecurityRulesResponse, err error) {
    if request == nil {
        request = NewCreateSecurityRulesRequest()
    }
    response = NewCreateSecurityRulesResponse()
    err = c.Send(request, response)
    return
}

func NewCreateTCPListenersRequest() (request *CreateTCPListenersRequest) {
    request = &CreateTCPListenersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gaap", APIVersion, "CreateTCPListeners")
    return
}

func NewCreateTCPListenersResponse() (response *CreateTCPListenersResponse) {
    response = &CreateTCPListenersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 该接口（CreateTCPListeners）用于批量创建单通道或者通道组的TCP协议类型的监听器。
func (c *Client) CreateTCPListeners(request *CreateTCPListenersRequest) (response *CreateTCPListenersResponse, err error) {
    if request == nil {
        request = NewCreateTCPListenersRequest()
    }
    response = NewCreateTCPListenersResponse()
    err = c.Send(request, response)
    return
}

func NewCreateUDPListenersRequest() (request *CreateUDPListenersRequest) {
    request = &CreateUDPListenersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gaap", APIVersion, "CreateUDPListeners")
    return
}

func NewCreateUDPListenersResponse() (response *CreateUDPListenersResponse) {
    response = &CreateUDPListenersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 该接口（CreateUDPListeners）用于批量创建单通道或者通道组的UDP协议类型的监听器。
func (c *Client) CreateUDPListeners(request *CreateUDPListenersRequest) (response *CreateUDPListenersResponse, err error) {
    if request == nil {
        request = NewCreateUDPListenersRequest()
    }
    response = NewCreateUDPListenersResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteCertificateRequest() (request *DeleteCertificateRequest) {
    request = &DeleteCertificateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gaap", APIVersion, "DeleteCertificate")
    return
}

func NewDeleteCertificateResponse() (response *DeleteCertificateResponse) {
    response = &DeleteCertificateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DeleteCertificate）用于删除证书。
func (c *Client) DeleteCertificate(request *DeleteCertificateRequest) (response *DeleteCertificateResponse, err error) {
    if request == nil {
        request = NewDeleteCertificateRequest()
    }
    response = NewDeleteCertificateResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteDomainRequest() (request *DeleteDomainRequest) {
    request = &DeleteDomainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gaap", APIVersion, "DeleteDomain")
    return
}

func NewDeleteDomainResponse() (response *DeleteDomainResponse) {
    response = &DeleteDomainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DeleteDomain）仅适用于7层监听器，用于删除该监听器下对应域名及域名下的所有规则，所有已绑定源站的规则将自动解绑。
func (c *Client) DeleteDomain(request *DeleteDomainRequest) (response *DeleteDomainResponse, err error) {
    if request == nil {
        request = NewDeleteDomainRequest()
    }
    response = NewDeleteDomainResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteListenersRequest() (request *DeleteListenersRequest) {
    request = &DeleteListenersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gaap", APIVersion, "DeleteListeners")
    return
}

func NewDeleteListenersResponse() (response *DeleteListenersResponse) {
    response = &DeleteListenersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 该接口（DeleteListeners）用于批量删除通道或通道组的监听器，包括4/7层监听器。
func (c *Client) DeleteListeners(request *DeleteListenersRequest) (response *DeleteListenersResponse, err error) {
    if request == nil {
        request = NewDeleteListenersRequest()
    }
    response = NewDeleteListenersResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteProxyGroupRequest() (request *DeleteProxyGroupRequest) {
    request = &DeleteProxyGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gaap", APIVersion, "DeleteProxyGroup")
    return
}

func NewDeleteProxyGroupResponse() (response *DeleteProxyGroupResponse) {
    response = &DeleteProxyGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DeleteProxyGroup）用于删除通道组。
func (c *Client) DeleteProxyGroup(request *DeleteProxyGroupRequest) (response *DeleteProxyGroupResponse, err error) {
    if request == nil {
        request = NewDeleteProxyGroupRequest()
    }
    response = NewDeleteProxyGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteRuleRequest() (request *DeleteRuleRequest) {
    request = &DeleteRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gaap", APIVersion, "DeleteRule")
    return
}

func NewDeleteRuleResponse() (response *DeleteRuleResponse) {
    response = &DeleteRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 该接口（DeleteRule）用于删除HTTP/HTTPS监听器的转发规则。
func (c *Client) DeleteRule(request *DeleteRuleRequest) (response *DeleteRuleResponse, err error) {
    if request == nil {
        request = NewDeleteRuleRequest()
    }
    response = NewDeleteRuleResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteSecurityPolicyRequest() (request *DeleteSecurityPolicyRequest) {
    request = &DeleteSecurityPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gaap", APIVersion, "DeleteSecurityPolicy")
    return
}

func NewDeleteSecurityPolicyResponse() (response *DeleteSecurityPolicyResponse) {
    response = &DeleteSecurityPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除安全策略
func (c *Client) DeleteSecurityPolicy(request *DeleteSecurityPolicyRequest) (response *DeleteSecurityPolicyResponse, err error) {
    if request == nil {
        request = NewDeleteSecurityPolicyRequest()
    }
    response = NewDeleteSecurityPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteSecurityRulesRequest() (request *DeleteSecurityRulesRequest) {
    request = &DeleteSecurityRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gaap", APIVersion, "DeleteSecurityRules")
    return
}

func NewDeleteSecurityRulesResponse() (response *DeleteSecurityRulesResponse) {
    response = &DeleteSecurityRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除安全策略规则
func (c *Client) DeleteSecurityRules(request *DeleteSecurityRulesRequest) (response *DeleteSecurityRulesResponse, err error) {
    if request == nil {
        request = NewDeleteSecurityRulesRequest()
    }
    response = NewDeleteSecurityRulesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAccessRegionsRequest() (request *DescribeAccessRegionsRequest) {
    request = &DescribeAccessRegionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gaap", APIVersion, "DescribeAccessRegions")
    return
}

func NewDescribeAccessRegionsResponse() (response *DescribeAccessRegionsResponse) {
    response = &DescribeAccessRegionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DescribeAccessRegions）用于查询加速区域，即客户端接入区域。
func (c *Client) DescribeAccessRegions(request *DescribeAccessRegionsRequest) (response *DescribeAccessRegionsResponse, err error) {
    if request == nil {
        request = NewDescribeAccessRegionsRequest()
    }
    response = NewDescribeAccessRegionsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAccessRegionsByDestRegionRequest() (request *DescribeAccessRegionsByDestRegionRequest) {
    request = &DescribeAccessRegionsByDestRegionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gaap", APIVersion, "DescribeAccessRegionsByDestRegion")
    return
}

func NewDescribeAccessRegionsByDestRegionResponse() (response *DescribeAccessRegionsByDestRegionResponse) {
    response = &DescribeAccessRegionsByDestRegionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DescribeAccessRegionsByDestRegion）根据源站区域查询可用的加速区域列表
func (c *Client) DescribeAccessRegionsByDestRegion(request *DescribeAccessRegionsByDestRegionRequest) (response *DescribeAccessRegionsByDestRegionResponse, err error) {
    if request == nil {
        request = NewDescribeAccessRegionsByDestRegionRequest()
    }
    response = NewDescribeAccessRegionsByDestRegionResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCertificateDetailRequest() (request *DescribeCertificateDetailRequest) {
    request = &DescribeCertificateDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gaap", APIVersion, "DescribeCertificateDetail")
    return
}

func NewDescribeCertificateDetailResponse() (response *DescribeCertificateDetailResponse) {
    response = &DescribeCertificateDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DescribeCertificateDetail）用于查询证书详情，包括证书ID，证书名字，证书类型，证书内容以及密钥等信息。
func (c *Client) DescribeCertificateDetail(request *DescribeCertificateDetailRequest) (response *DescribeCertificateDetailResponse, err error) {
    if request == nil {
        request = NewDescribeCertificateDetailRequest()
    }
    response = NewDescribeCertificateDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCertificatesRequest() (request *DescribeCertificatesRequest) {
    request = &DescribeCertificatesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gaap", APIVersion, "DescribeCertificates")
    return
}

func NewDescribeCertificatesResponse() (response *DescribeCertificatesResponse) {
    response = &DescribeCertificatesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DescribeCertificates）用来查询可以使用的证书列表。
func (c *Client) DescribeCertificates(request *DescribeCertificatesRequest) (response *DescribeCertificatesResponse, err error) {
    if request == nil {
        request = NewDescribeCertificatesRequest()
    }
    response = NewDescribeCertificatesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCountryAreaMappingRequest() (request *DescribeCountryAreaMappingRequest) {
    request = &DescribeCountryAreaMappingRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gaap", APIVersion, "DescribeCountryAreaMapping")
    return
}

func NewDescribeCountryAreaMappingResponse() (response *DescribeCountryAreaMappingResponse) {
    response = &DescribeCountryAreaMappingResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DescribeCountryAreaMapping）用于获取国家地区编码映射表。
func (c *Client) DescribeCountryAreaMapping(request *DescribeCountryAreaMappingRequest) (response *DescribeCountryAreaMappingResponse, err error) {
    if request == nil {
        request = NewDescribeCountryAreaMappingRequest()
    }
    response = NewDescribeCountryAreaMappingResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDestRegionsRequest() (request *DescribeDestRegionsRequest) {
    request = &DescribeDestRegionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gaap", APIVersion, "DescribeDestRegions")
    return
}

func NewDescribeDestRegionsResponse() (response *DescribeDestRegionsResponse) {
    response = &DescribeDestRegionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DescribeDestRegions）用于查询源站区域，即源站服务器所在区域。
func (c *Client) DescribeDestRegions(request *DescribeDestRegionsRequest) (response *DescribeDestRegionsResponse, err error) {
    if request == nil {
        request = NewDescribeDestRegionsRequest()
    }
    response = NewDescribeDestRegionsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeGroupAndStatisticsProxyRequest() (request *DescribeGroupAndStatisticsProxyRequest) {
    request = &DescribeGroupAndStatisticsProxyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gaap", APIVersion, "DescribeGroupAndStatisticsProxy")
    return
}

func NewDescribeGroupAndStatisticsProxyResponse() (response *DescribeGroupAndStatisticsProxyResponse) {
    response = &DescribeGroupAndStatisticsProxyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 该接口为内部接口，用于查询可以获取统计数据的通道组和通道信息
func (c *Client) DescribeGroupAndStatisticsProxy(request *DescribeGroupAndStatisticsProxyRequest) (response *DescribeGroupAndStatisticsProxyResponse, err error) {
    if request == nil {
        request = NewDescribeGroupAndStatisticsProxyRequest()
    }
    response = NewDescribeGroupAndStatisticsProxyResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeGroupDomainConfigRequest() (request *DescribeGroupDomainConfigRequest) {
    request = &DescribeGroupDomainConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gaap", APIVersion, "DescribeGroupDomainConfig")
    return
}

func NewDescribeGroupDomainConfigResponse() (response *DescribeGroupDomainConfigResponse) {
    response = &DescribeGroupDomainConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DescribeGroupDomainConfig）用于获取通道组域名解析配置详情。
func (c *Client) DescribeGroupDomainConfig(request *DescribeGroupDomainConfigRequest) (response *DescribeGroupDomainConfigResponse, err error) {
    if request == nil {
        request = NewDescribeGroupDomainConfigRequest()
    }
    response = NewDescribeGroupDomainConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeHTTPListenersRequest() (request *DescribeHTTPListenersRequest) {
    request = &DescribeHTTPListenersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gaap", APIVersion, "DescribeHTTPListeners")
    return
}

func NewDescribeHTTPListenersResponse() (response *DescribeHTTPListenersResponse) {
    response = &DescribeHTTPListenersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 该接口（DescribeHTTPListeners）用来查询HTTP监听器信息。
func (c *Client) DescribeHTTPListeners(request *DescribeHTTPListenersRequest) (response *DescribeHTTPListenersResponse, err error) {
    if request == nil {
        request = NewDescribeHTTPListenersRequest()
    }
    response = NewDescribeHTTPListenersResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeHTTPSListenersRequest() (request *DescribeHTTPSListenersRequest) {
    request = &DescribeHTTPSListenersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gaap", APIVersion, "DescribeHTTPSListeners")
    return
}

func NewDescribeHTTPSListenersResponse() (response *DescribeHTTPSListenersResponse) {
    response = &DescribeHTTPSListenersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DescribeHTTPSListeners）用来查询HTTPS监听器信息。
func (c *Client) DescribeHTTPSListeners(request *DescribeHTTPSListenersRequest) (response *DescribeHTTPSListenersResponse, err error) {
    if request == nil {
        request = NewDescribeHTTPSListenersRequest()
    }
    response = NewDescribeHTTPSListenersResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeListenerRealServersRequest() (request *DescribeListenerRealServersRequest) {
    request = &DescribeListenerRealServersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gaap", APIVersion, "DescribeListenerRealServers")
    return
}

func NewDescribeListenerRealServersResponse() (response *DescribeListenerRealServersResponse) {
    response = &DescribeListenerRealServersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 该接口（DescribeListenerRealServers）用于查询TCP/UDP监听器源站列表，包括该监听器已经绑定的源站列表以及可以绑定的源站列表。
func (c *Client) DescribeListenerRealServers(request *DescribeListenerRealServersRequest) (response *DescribeListenerRealServersResponse, err error) {
    if request == nil {
        request = NewDescribeListenerRealServersRequest()
    }
    response = NewDescribeListenerRealServersResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeListenerStatisticsRequest() (request *DescribeListenerStatisticsRequest) {
    request = &DescribeListenerStatisticsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gaap", APIVersion, "DescribeListenerStatistics")
    return
}

func NewDescribeListenerStatisticsResponse() (response *DescribeListenerStatisticsResponse) {
    response = &DescribeListenerStatisticsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 该接口用于查询监听器统计数据，包括出入带宽，出入包量，并发数据。支持300秒, 3600秒和86400秒的细粒度，取值为细粒度范围内最大值。
func (c *Client) DescribeListenerStatistics(request *DescribeListenerStatisticsRequest) (response *DescribeListenerStatisticsResponse, err error) {
    if request == nil {
        request = NewDescribeListenerStatisticsRequest()
    }
    response = NewDescribeListenerStatisticsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeProxiesRequest() (request *DescribeProxiesRequest) {
    request = &DescribeProxiesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gaap", APIVersion, "DescribeProxies")
    return
}

func NewDescribeProxiesResponse() (response *DescribeProxiesResponse) {
    response = &DescribeProxiesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DescribeProxies）用于查询通道实例列表。
func (c *Client) DescribeProxies(request *DescribeProxiesRequest) (response *DescribeProxiesResponse, err error) {
    if request == nil {
        request = NewDescribeProxiesRequest()
    }
    response = NewDescribeProxiesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeProxiesStatusRequest() (request *DescribeProxiesStatusRequest) {
    request = &DescribeProxiesStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gaap", APIVersion, "DescribeProxiesStatus")
    return
}

func NewDescribeProxiesStatusResponse() (response *DescribeProxiesStatusResponse) {
    response = &DescribeProxiesStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DescribeProxiesStatus）用于查询通道状态列表。
func (c *Client) DescribeProxiesStatus(request *DescribeProxiesStatusRequest) (response *DescribeProxiesStatusResponse, err error) {
    if request == nil {
        request = NewDescribeProxiesStatusRequest()
    }
    response = NewDescribeProxiesStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeProxyAndStatisticsListenersRequest() (request *DescribeProxyAndStatisticsListenersRequest) {
    request = &DescribeProxyAndStatisticsListenersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gaap", APIVersion, "DescribeProxyAndStatisticsListeners")
    return
}

func NewDescribeProxyAndStatisticsListenersResponse() (response *DescribeProxyAndStatisticsListenersResponse) {
    response = &DescribeProxyAndStatisticsListenersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 该接口为内部接口，用于查询可以获取统计数据的通道和监听器信息
func (c *Client) DescribeProxyAndStatisticsListeners(request *DescribeProxyAndStatisticsListenersRequest) (response *DescribeProxyAndStatisticsListenersResponse, err error) {
    if request == nil {
        request = NewDescribeProxyAndStatisticsListenersRequest()
    }
    response = NewDescribeProxyAndStatisticsListenersResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeProxyDetailRequest() (request *DescribeProxyDetailRequest) {
    request = &DescribeProxyDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gaap", APIVersion, "DescribeProxyDetail")
    return
}

func NewDescribeProxyDetailResponse() (response *DescribeProxyDetailResponse) {
    response = &DescribeProxyDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DescribeProxyDetail）用于查询通道详情。
func (c *Client) DescribeProxyDetail(request *DescribeProxyDetailRequest) (response *DescribeProxyDetailResponse, err error) {
    if request == nil {
        request = NewDescribeProxyDetailRequest()
    }
    response = NewDescribeProxyDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeProxyGroupDetailsRequest() (request *DescribeProxyGroupDetailsRequest) {
    request = &DescribeProxyGroupDetailsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gaap", APIVersion, "DescribeProxyGroupDetails")
    return
}

func NewDescribeProxyGroupDetailsResponse() (response *DescribeProxyGroupDetailsResponse) {
    response = &DescribeProxyGroupDetailsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DescribeProxyGroupDetails）用于查询通道组详情。
func (c *Client) DescribeProxyGroupDetails(request *DescribeProxyGroupDetailsRequest) (response *DescribeProxyGroupDetailsResponse, err error) {
    if request == nil {
        request = NewDescribeProxyGroupDetailsRequest()
    }
    response = NewDescribeProxyGroupDetailsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeProxyGroupListRequest() (request *DescribeProxyGroupListRequest) {
    request = &DescribeProxyGroupListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gaap", APIVersion, "DescribeProxyGroupList")
    return
}

func NewDescribeProxyGroupListResponse() (response *DescribeProxyGroupListResponse) {
    response = &DescribeProxyGroupListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DescribeProxyGroupList）用于拉取通道组列表及各通道组基本信息。
func (c *Client) DescribeProxyGroupList(request *DescribeProxyGroupListRequest) (response *DescribeProxyGroupListResponse, err error) {
    if request == nil {
        request = NewDescribeProxyGroupListRequest()
    }
    response = NewDescribeProxyGroupListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeProxyGroupStatisticsRequest() (request *DescribeProxyGroupStatisticsRequest) {
    request = &DescribeProxyGroupStatisticsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gaap", APIVersion, "DescribeProxyGroupStatistics")
    return
}

func NewDescribeProxyGroupStatisticsResponse() (response *DescribeProxyGroupStatisticsResponse) {
    response = &DescribeProxyGroupStatisticsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 该接口用于查询监听器统计数据，包括出入带宽，出入包量，并发数据。支持300, 3600和86400的细粒度，取值为细粒度范围内最大值。
func (c *Client) DescribeProxyGroupStatistics(request *DescribeProxyGroupStatisticsRequest) (response *DescribeProxyGroupStatisticsResponse, err error) {
    if request == nil {
        request = NewDescribeProxyGroupStatisticsRequest()
    }
    response = NewDescribeProxyGroupStatisticsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeProxyStatisticsRequest() (request *DescribeProxyStatisticsRequest) {
    request = &DescribeProxyStatisticsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gaap", APIVersion, "DescribeProxyStatistics")
    return
}

func NewDescribeProxyStatisticsResponse() (response *DescribeProxyStatisticsResponse) {
    response = &DescribeProxyStatisticsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 该接口用于查询监听器统计数据，包括出入带宽，出入包量，并发，丢包和时延数据。支持300, 3600和86400的细粒度，取值为细粒度范围内最大值。
func (c *Client) DescribeProxyStatistics(request *DescribeProxyStatisticsRequest) (response *DescribeProxyStatisticsResponse, err error) {
    if request == nil {
        request = NewDescribeProxyStatisticsRequest()
    }
    response = NewDescribeProxyStatisticsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRealServerStatisticsRequest() (request *DescribeRealServerStatisticsRequest) {
    request = &DescribeRealServerStatisticsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gaap", APIVersion, "DescribeRealServerStatistics")
    return
}

func NewDescribeRealServerStatisticsResponse() (response *DescribeRealServerStatisticsResponse) {
    response = &DescribeRealServerStatisticsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 该接口（DescribeRealServerStatistics）用于查询源站健康检查结果的统计数据。源站状态展示位为1：正常或者0：异常。查询的源站需要在监听器或者规则上进行了绑定，查询时需指定绑定的监听器或者规则ID。该接口支持最近1，3，6，12，24小时内1分钟细粒度的源站状态统计数据展示。
func (c *Client) DescribeRealServerStatistics(request *DescribeRealServerStatisticsRequest) (response *DescribeRealServerStatisticsResponse, err error) {
    if request == nil {
        request = NewDescribeRealServerStatisticsRequest()
    }
    response = NewDescribeRealServerStatisticsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRealServersRequest() (request *DescribeRealServersRequest) {
    request = &DescribeRealServersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gaap", APIVersion, "DescribeRealServers")
    return
}

func NewDescribeRealServersResponse() (response *DescribeRealServersResponse) {
    response = &DescribeRealServersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DescribeRealServers）用于查询源站信息，可以根据项目名查询所有的源站信息，此外支持指定IP机或者域名的源站模糊查询。
func (c *Client) DescribeRealServers(request *DescribeRealServersRequest) (response *DescribeRealServersResponse, err error) {
    if request == nil {
        request = NewDescribeRealServersRequest()
    }
    response = NewDescribeRealServersResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRealServersStatusRequest() (request *DescribeRealServersStatusRequest) {
    request = &DescribeRealServersStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gaap", APIVersion, "DescribeRealServersStatus")
    return
}

func NewDescribeRealServersStatusResponse() (response *DescribeRealServersStatusResponse) {
    response = &DescribeRealServersStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DescribeRealServersStatus）用于查询源站是否已被规则或者监听器绑定
func (c *Client) DescribeRealServersStatus(request *DescribeRealServersStatusRequest) (response *DescribeRealServersStatusResponse, err error) {
    if request == nil {
        request = NewDescribeRealServersStatusRequest()
    }
    response = NewDescribeRealServersStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRegionAndPriceRequest() (request *DescribeRegionAndPriceRequest) {
    request = &DescribeRegionAndPriceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gaap", APIVersion, "DescribeRegionAndPrice")
    return
}

func NewDescribeRegionAndPriceResponse() (response *DescribeRegionAndPriceResponse) {
    response = &DescribeRegionAndPriceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 该接口（DescribeRegionAndPrice）用于获取源站区域和带宽梯度价格
func (c *Client) DescribeRegionAndPrice(request *DescribeRegionAndPriceRequest) (response *DescribeRegionAndPriceResponse, err error) {
    if request == nil {
        request = NewDescribeRegionAndPriceRequest()
    }
    response = NewDescribeRegionAndPriceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeResourcesByTagRequest() (request *DescribeResourcesByTagRequest) {
    request = &DescribeResourcesByTagRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gaap", APIVersion, "DescribeResourcesByTag")
    return
}

func NewDescribeResourcesByTagResponse() (response *DescribeResourcesByTagResponse) {
    response = &DescribeResourcesByTagResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DescribeResourcesByTag）用于根据标签来查询对应的资源信息，包括通道，通道组和源站。
func (c *Client) DescribeResourcesByTag(request *DescribeResourcesByTagRequest) (response *DescribeResourcesByTagResponse, err error) {
    if request == nil {
        request = NewDescribeResourcesByTagRequest()
    }
    response = NewDescribeResourcesByTagResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRuleRealServersRequest() (request *DescribeRuleRealServersRequest) {
    request = &DescribeRuleRealServersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gaap", APIVersion, "DescribeRuleRealServers")
    return
}

func NewDescribeRuleRealServersResponse() (response *DescribeRuleRealServersResponse) {
    response = &DescribeRuleRealServersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DescribeRuleRealServers）用于查询转发规则相关的源站信息， 包括该规则可绑定的源站信息和已绑定的源站信息。
func (c *Client) DescribeRuleRealServers(request *DescribeRuleRealServersRequest) (response *DescribeRuleRealServersResponse, err error) {
    if request == nil {
        request = NewDescribeRuleRealServersRequest()
    }
    response = NewDescribeRuleRealServersResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRulesRequest() (request *DescribeRulesRequest) {
    request = &DescribeRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gaap", APIVersion, "DescribeRules")
    return
}

func NewDescribeRulesResponse() (response *DescribeRulesResponse) {
    response = &DescribeRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DescribeRules）用于查询监听器下的所有规则信息，包括规则域名，路径以及该规则下所绑定的源站列表。当通道版本为3.0时，该接口会返回该域名对应的高级认证配置信息。
func (c *Client) DescribeRules(request *DescribeRulesRequest) (response *DescribeRulesResponse, err error) {
    if request == nil {
        request = NewDescribeRulesRequest()
    }
    response = NewDescribeRulesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSecurityPolicyDetailRequest() (request *DescribeSecurityPolicyDetailRequest) {
    request = &DescribeSecurityPolicyDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gaap", APIVersion, "DescribeSecurityPolicyDetail")
    return
}

func NewDescribeSecurityPolicyDetailResponse() (response *DescribeSecurityPolicyDetailResponse) {
    response = &DescribeSecurityPolicyDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取安全策略详情
func (c *Client) DescribeSecurityPolicyDetail(request *DescribeSecurityPolicyDetailRequest) (response *DescribeSecurityPolicyDetailResponse, err error) {
    if request == nil {
        request = NewDescribeSecurityPolicyDetailRequest()
    }
    response = NewDescribeSecurityPolicyDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTCPListenersRequest() (request *DescribeTCPListenersRequest) {
    request = &DescribeTCPListenersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gaap", APIVersion, "DescribeTCPListeners")
    return
}

func NewDescribeTCPListenersResponse() (response *DescribeTCPListenersResponse) {
    response = &DescribeTCPListenersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 该接口（DescribeTCPListeners）用于查询单通道或者通道组下的TCP监听器信息。
func (c *Client) DescribeTCPListeners(request *DescribeTCPListenersRequest) (response *DescribeTCPListenersResponse, err error) {
    if request == nil {
        request = NewDescribeTCPListenersRequest()
    }
    response = NewDescribeTCPListenersResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUDPListenersRequest() (request *DescribeUDPListenersRequest) {
    request = &DescribeUDPListenersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gaap", APIVersion, "DescribeUDPListeners")
    return
}

func NewDescribeUDPListenersResponse() (response *DescribeUDPListenersResponse) {
    response = &DescribeUDPListenersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 该接口（DescribeUDPListeners）用于查询单通道或者通道组下的UDP监听器信息
func (c *Client) DescribeUDPListeners(request *DescribeUDPListenersRequest) (response *DescribeUDPListenersResponse, err error) {
    if request == nil {
        request = NewDescribeUDPListenersRequest()
    }
    response = NewDescribeUDPListenersResponse()
    err = c.Send(request, response)
    return
}

func NewDestroyProxiesRequest() (request *DestroyProxiesRequest) {
    request = &DestroyProxiesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gaap", APIVersion, "DestroyProxies")
    return
}

func NewDestroyProxiesResponse() (response *DestroyProxiesResponse) {
    response = &DestroyProxiesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DestroyProxies）用于销毁。通道销毁后，不再产生任何费用。
func (c *Client) DestroyProxies(request *DestroyProxiesRequest) (response *DestroyProxiesResponse, err error) {
    if request == nil {
        request = NewDestroyProxiesRequest()
    }
    response = NewDestroyProxiesResponse()
    err = c.Send(request, response)
    return
}

func NewInquiryPriceCreateProxyRequest() (request *InquiryPriceCreateProxyRequest) {
    request = &InquiryPriceCreateProxyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gaap", APIVersion, "InquiryPriceCreateProxy")
    return
}

func NewInquiryPriceCreateProxyResponse() (response *InquiryPriceCreateProxyResponse) {
    response = &InquiryPriceCreateProxyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（InquiryPriceCreateProxy）用于创建加速通道询价。
func (c *Client) InquiryPriceCreateProxy(request *InquiryPriceCreateProxyRequest) (response *InquiryPriceCreateProxyResponse, err error) {
    if request == nil {
        request = NewInquiryPriceCreateProxyRequest()
    }
    response = NewInquiryPriceCreateProxyResponse()
    err = c.Send(request, response)
    return
}

func NewModifyCertificateRequest() (request *ModifyCertificateRequest) {
    request = &ModifyCertificateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gaap", APIVersion, "ModifyCertificate")
    return
}

func NewModifyCertificateResponse() (response *ModifyCertificateResponse) {
    response = &ModifyCertificateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（ModifyCertificate）用于修改监听器下的域名对应的证书。该接口仅适用于version3.0的通道。
func (c *Client) ModifyCertificate(request *ModifyCertificateRequest) (response *ModifyCertificateResponse, err error) {
    if request == nil {
        request = NewModifyCertificateRequest()
    }
    response = NewModifyCertificateResponse()
    err = c.Send(request, response)
    return
}

func NewModifyCertificateAttributesRequest() (request *ModifyCertificateAttributesRequest) {
    request = &ModifyCertificateAttributesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gaap", APIVersion, "ModifyCertificateAttributes")
    return
}

func NewModifyCertificateAttributesResponse() (response *ModifyCertificateAttributesResponse) {
    response = &ModifyCertificateAttributesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（ModifyCertificateAttributes）用于修改证书，包括证明名字以及证书内容。
func (c *Client) ModifyCertificateAttributes(request *ModifyCertificateAttributesRequest) (response *ModifyCertificateAttributesResponse, err error) {
    if request == nil {
        request = NewModifyCertificateAttributesRequest()
    }
    response = NewModifyCertificateAttributesResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDomainRequest() (request *ModifyDomainRequest) {
    request = &ModifyDomainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gaap", APIVersion, "ModifyDomain")
    return
}

func NewModifyDomainResponse() (response *ModifyDomainResponse) {
    response = &ModifyDomainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（ModifyDomain）用于监听器下的域名。当通道版本为3.0时，支持对该域名所对应的证书修改。
func (c *Client) ModifyDomain(request *ModifyDomainRequest) (response *ModifyDomainResponse, err error) {
    if request == nil {
        request = NewModifyDomainRequest()
    }
    response = NewModifyDomainResponse()
    err = c.Send(request, response)
    return
}

func NewModifyGroupDomainConfigRequest() (request *ModifyGroupDomainConfigRequest) {
    request = &ModifyGroupDomainConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gaap", APIVersion, "ModifyGroupDomainConfig")
    return
}

func NewModifyGroupDomainConfigResponse() (response *ModifyGroupDomainConfigResponse) {
    response = &ModifyGroupDomainConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（ModifyGroupDomainConfig）用于配置通道组就近接入域名。
func (c *Client) ModifyGroupDomainConfig(request *ModifyGroupDomainConfigRequest) (response *ModifyGroupDomainConfigResponse, err error) {
    if request == nil {
        request = NewModifyGroupDomainConfigRequest()
    }
    response = NewModifyGroupDomainConfigResponse()
    err = c.Send(request, response)
    return
}

func NewModifyHTTPListenerAttributeRequest() (request *ModifyHTTPListenerAttributeRequest) {
    request = &ModifyHTTPListenerAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gaap", APIVersion, "ModifyHTTPListenerAttribute")
    return
}

func NewModifyHTTPListenerAttributeResponse() (response *ModifyHTTPListenerAttributeResponse) {
    response = &ModifyHTTPListenerAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 该接口（ModifyHTTPListenerAttribute）用于修改通道的HTTP监听器配置信息，目前仅支持修改监听器的名称。
// 注意：通道组通道暂时不支持HTTP/HTTPS监听器。
func (c *Client) ModifyHTTPListenerAttribute(request *ModifyHTTPListenerAttributeRequest) (response *ModifyHTTPListenerAttributeResponse, err error) {
    if request == nil {
        request = NewModifyHTTPListenerAttributeRequest()
    }
    response = NewModifyHTTPListenerAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyHTTPSListenerAttributeRequest() (request *ModifyHTTPSListenerAttributeRequest) {
    request = &ModifyHTTPSListenerAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gaap", APIVersion, "ModifyHTTPSListenerAttribute")
    return
}

func NewModifyHTTPSListenerAttributeResponse() (response *ModifyHTTPSListenerAttributeResponse) {
    response = &ModifyHTTPSListenerAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 该接口（ModifyHTTPSListenerAttribute）用于修改HTTPS监听器配置，当前不支持通道组和v1版本通道。
func (c *Client) ModifyHTTPSListenerAttribute(request *ModifyHTTPSListenerAttributeRequest) (response *ModifyHTTPSListenerAttributeResponse, err error) {
    if request == nil {
        request = NewModifyHTTPSListenerAttributeRequest()
    }
    response = NewModifyHTTPSListenerAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyProxiesAttributeRequest() (request *ModifyProxiesAttributeRequest) {
    request = &ModifyProxiesAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gaap", APIVersion, "ModifyProxiesAttribute")
    return
}

func NewModifyProxiesAttributeResponse() (response *ModifyProxiesAttributeResponse) {
    response = &ModifyProxiesAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（ModifyProxiesAttribute）用于修改实例的属性（目前只支持修改通道的名称）。
func (c *Client) ModifyProxiesAttribute(request *ModifyProxiesAttributeRequest) (response *ModifyProxiesAttributeResponse, err error) {
    if request == nil {
        request = NewModifyProxiesAttributeRequest()
    }
    response = NewModifyProxiesAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyProxiesProjectRequest() (request *ModifyProxiesProjectRequest) {
    request = &ModifyProxiesProjectRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gaap", APIVersion, "ModifyProxiesProject")
    return
}

func NewModifyProxiesProjectResponse() (response *ModifyProxiesProjectResponse) {
    response = &ModifyProxiesProjectResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（ModifyProxiesProject）用于修改通道所属项目。
func (c *Client) ModifyProxiesProject(request *ModifyProxiesProjectRequest) (response *ModifyProxiesProjectResponse, err error) {
    if request == nil {
        request = NewModifyProxiesProjectRequest()
    }
    response = NewModifyProxiesProjectResponse()
    err = c.Send(request, response)
    return
}

func NewModifyProxyConfigurationRequest() (request *ModifyProxyConfigurationRequest) {
    request = &ModifyProxyConfigurationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gaap", APIVersion, "ModifyProxyConfiguration")
    return
}

func NewModifyProxyConfigurationResponse() (response *ModifyProxyConfigurationResponse) {
    response = &ModifyProxyConfigurationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（ModifyProxyConfiguration）用于修改通道的配置。根据当前业务的容量需求，扩容或缩容相关通道的配置。仅支持Scalarable为1的通道,Scalarable可通过接口DescribeProxies获取。
func (c *Client) ModifyProxyConfiguration(request *ModifyProxyConfigurationRequest) (response *ModifyProxyConfigurationResponse, err error) {
    if request == nil {
        request = NewModifyProxyConfigurationRequest()
    }
    response = NewModifyProxyConfigurationResponse()
    err = c.Send(request, response)
    return
}

func NewModifyProxyGroupAttributeRequest() (request *ModifyProxyGroupAttributeRequest) {
    request = &ModifyProxyGroupAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gaap", APIVersion, "ModifyProxyGroupAttribute")
    return
}

func NewModifyProxyGroupAttributeResponse() (response *ModifyProxyGroupAttributeResponse) {
    response = &ModifyProxyGroupAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（ModifyProxyGroupAttribute）用于修改通道组属性，目前仅支持修改通道组名称。
func (c *Client) ModifyProxyGroupAttribute(request *ModifyProxyGroupAttributeRequest) (response *ModifyProxyGroupAttributeResponse, err error) {
    if request == nil {
        request = NewModifyProxyGroupAttributeRequest()
    }
    response = NewModifyProxyGroupAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyRealServerNameRequest() (request *ModifyRealServerNameRequest) {
    request = &ModifyRealServerNameRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gaap", APIVersion, "ModifyRealServerName")
    return
}

func NewModifyRealServerNameResponse() (response *ModifyRealServerNameResponse) {
    response = &ModifyRealServerNameResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（ModifyRealServerName）用于修改源站的名称
func (c *Client) ModifyRealServerName(request *ModifyRealServerNameRequest) (response *ModifyRealServerNameResponse, err error) {
    if request == nil {
        request = NewModifyRealServerNameRequest()
    }
    response = NewModifyRealServerNameResponse()
    err = c.Send(request, response)
    return
}

func NewModifyRuleAttributeRequest() (request *ModifyRuleAttributeRequest) {
    request = &ModifyRuleAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gaap", APIVersion, "ModifyRuleAttribute")
    return
}

func NewModifyRuleAttributeResponse() (response *ModifyRuleAttributeResponse) {
    response = &ModifyRuleAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（ModifyRuleAttribute）用于修改转发规则的信息，包括健康检查的配置以及转发策略。
func (c *Client) ModifyRuleAttribute(request *ModifyRuleAttributeRequest) (response *ModifyRuleAttributeResponse, err error) {
    if request == nil {
        request = NewModifyRuleAttributeRequest()
    }
    response = NewModifyRuleAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewModifySecurityRuleRequest() (request *ModifySecurityRuleRequest) {
    request = &ModifySecurityRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gaap", APIVersion, "ModifySecurityRule")
    return
}

func NewModifySecurityRuleResponse() (response *ModifySecurityRuleResponse) {
    response = &ModifySecurityRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改安全策略规则名
func (c *Client) ModifySecurityRule(request *ModifySecurityRuleRequest) (response *ModifySecurityRuleResponse, err error) {
    if request == nil {
        request = NewModifySecurityRuleRequest()
    }
    response = NewModifySecurityRuleResponse()
    err = c.Send(request, response)
    return
}

func NewModifyTCPListenerAttributeRequest() (request *ModifyTCPListenerAttributeRequest) {
    request = &ModifyTCPListenerAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gaap", APIVersion, "ModifyTCPListenerAttribute")
    return
}

func NewModifyTCPListenerAttributeResponse() (response *ModifyTCPListenerAttributeResponse) {
    response = &ModifyTCPListenerAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（ModifyTCPListenerAttribute）用于修改通道实例下TCP监听器配置，包括健康检查的配置，调度策略。
func (c *Client) ModifyTCPListenerAttribute(request *ModifyTCPListenerAttributeRequest) (response *ModifyTCPListenerAttributeResponse, err error) {
    if request == nil {
        request = NewModifyTCPListenerAttributeRequest()
    }
    response = NewModifyTCPListenerAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyUDPListenerAttributeRequest() (request *ModifyUDPListenerAttributeRequest) {
    request = &ModifyUDPListenerAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gaap", APIVersion, "ModifyUDPListenerAttribute")
    return
}

func NewModifyUDPListenerAttributeResponse() (response *ModifyUDPListenerAttributeResponse) {
    response = &ModifyUDPListenerAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（ModifyUDPListenerAttribute）用于修改通道实例下UDP监听器配置，包括监听器名称和调度策略的修改。
func (c *Client) ModifyUDPListenerAttribute(request *ModifyUDPListenerAttributeRequest) (response *ModifyUDPListenerAttributeResponse, err error) {
    if request == nil {
        request = NewModifyUDPListenerAttributeRequest()
    }
    response = NewModifyUDPListenerAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewOpenProxiesRequest() (request *OpenProxiesRequest) {
    request = &OpenProxiesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gaap", APIVersion, "OpenProxies")
    return
}

func NewOpenProxiesResponse() (response *OpenProxiesResponse) {
    response = &OpenProxiesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 该接口（OpenProxies）用于开启一条或者多条通道。
func (c *Client) OpenProxies(request *OpenProxiesRequest) (response *OpenProxiesResponse, err error) {
    if request == nil {
        request = NewOpenProxiesRequest()
    }
    response = NewOpenProxiesResponse()
    err = c.Send(request, response)
    return
}

func NewOpenSecurityPolicyRequest() (request *OpenSecurityPolicyRequest) {
    request = &OpenSecurityPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gaap", APIVersion, "OpenSecurityPolicy")
    return
}

func NewOpenSecurityPolicyResponse() (response *OpenSecurityPolicyResponse) {
    response = &OpenSecurityPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 开启安全策略
func (c *Client) OpenSecurityPolicy(request *OpenSecurityPolicyRequest) (response *OpenSecurityPolicyResponse, err error) {
    if request == nil {
        request = NewOpenSecurityPolicyRequest()
    }
    response = NewOpenSecurityPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewRemoveRealServersRequest() (request *RemoveRealServersRequest) {
    request = &RemoveRealServersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gaap", APIVersion, "RemoveRealServers")
    return
}

func NewRemoveRealServersResponse() (response *RemoveRealServersResponse) {
    response = &RemoveRealServersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除已添加的源站(服务器)IP或域名
func (c *Client) RemoveRealServers(request *RemoveRealServersRequest) (response *RemoveRealServersResponse, err error) {
    if request == nil {
        request = NewRemoveRealServersRequest()
    }
    response = NewRemoveRealServersResponse()
    err = c.Send(request, response)
    return
}

func NewSetAuthenticationRequest() (request *SetAuthenticationRequest) {
    request = &SetAuthenticationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gaap", APIVersion, "SetAuthentication")
    return
}

func NewSetAuthenticationResponse() (response *SetAuthenticationResponse) {
    response = &SetAuthenticationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（SetAuthentication）用于通道的高级认证配置，包括认证方式选择，以及各种认证方式对应的证书选择。仅支持Version3.0的通道。
func (c *Client) SetAuthentication(request *SetAuthenticationRequest) (response *SetAuthenticationResponse, err error) {
    if request == nil {
        request = NewSetAuthenticationRequest()
    }
    response = NewSetAuthenticationResponse()
    err = c.Send(request, response)
    return
}
