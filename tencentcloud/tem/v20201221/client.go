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

package v20201221

import (
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2020-12-21"

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


func NewCreateCosTokenRequest() (request *CreateCosTokenRequest) {
    request = &CreateCosTokenRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tem", APIVersion, "CreateCosToken")
    
    
    return
}

func NewCreateCosTokenResponse() (response *CreateCosTokenResponse) {
    response = &CreateCosTokenResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateCosToken
// 生成Cos临时秘钥
func (c *Client) CreateCosToken(request *CreateCosTokenRequest) (response *CreateCosTokenResponse, err error) {
    if request == nil {
        request = NewCreateCosTokenRequest()
    }
    
    response = NewCreateCosTokenResponse()
    err = c.Send(request, response)
    return
}

func NewCreateCosTokenV2Request() (request *CreateCosTokenV2Request) {
    request = &CreateCosTokenV2Request{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tem", APIVersion, "CreateCosTokenV2")
    
    
    return
}

func NewCreateCosTokenV2Response() (response *CreateCosTokenV2Response) {
    response = &CreateCosTokenV2Response{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateCosTokenV2
// 生成Cos临时秘钥
func (c *Client) CreateCosTokenV2(request *CreateCosTokenV2Request) (response *CreateCosTokenV2Response, err error) {
    if request == nil {
        request = NewCreateCosTokenV2Request()
    }
    
    response = NewCreateCosTokenV2Response()
    err = c.Send(request, response)
    return
}

func NewCreateNamespaceRequest() (request *CreateNamespaceRequest) {
    request = &CreateNamespaceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tem", APIVersion, "CreateNamespace")
    
    
    return
}

func NewCreateNamespaceResponse() (response *CreateNamespaceResponse) {
    response = &CreateNamespaceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateNamespace
// 创建环境
func (c *Client) CreateNamespace(request *CreateNamespaceRequest) (response *CreateNamespaceResponse, err error) {
    if request == nil {
        request = NewCreateNamespaceRequest()
    }
    
    response = NewCreateNamespaceResponse()
    err = c.Send(request, response)
    return
}

func NewCreateResourceRequest() (request *CreateResourceRequest) {
    request = &CreateResourceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tem", APIVersion, "CreateResource")
    
    
    return
}

func NewCreateResourceResponse() (response *CreateResourceResponse) {
    response = &CreateResourceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateResource
// 绑定云资源
func (c *Client) CreateResource(request *CreateResourceRequest) (response *CreateResourceResponse, err error) {
    if request == nil {
        request = NewCreateResourceRequest()
    }
    
    response = NewCreateResourceResponse()
    err = c.Send(request, response)
    return
}

func NewCreateServiceV2Request() (request *CreateServiceV2Request) {
    request = &CreateServiceV2Request{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tem", APIVersion, "CreateServiceV2")
    
    
    return
}

func NewCreateServiceV2Response() (response *CreateServiceV2Response) {
    response = &CreateServiceV2Response{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateServiceV2
// 创建服务
//
// 可能返回的错误码:
//  INTERNALERROR_CREATESERVICEERROR = "InternalError.CreateServiceError"
func (c *Client) CreateServiceV2(request *CreateServiceV2Request) (response *CreateServiceV2Response, err error) {
    if request == nil {
        request = NewCreateServiceV2Request()
    }
    
    response = NewCreateServiceV2Response()
    err = c.Send(request, response)
    return
}

func NewDeleteIngressRequest() (request *DeleteIngressRequest) {
    request = &DeleteIngressRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tem", APIVersion, "DeleteIngress")
    
    
    return
}

func NewDeleteIngressResponse() (response *DeleteIngressResponse) {
    response = &DeleteIngressResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteIngress
// 删除 Ingress 规则
//
// 可能返回的错误码:
//  INTERNALERROR_CREATESERVICEERROR = "InternalError.CreateServiceError"
func (c *Client) DeleteIngress(request *DeleteIngressRequest) (response *DeleteIngressResponse, err error) {
    if request == nil {
        request = NewDeleteIngressRequest()
    }
    
    response = NewDeleteIngressResponse()
    err = c.Send(request, response)
    return
}

func NewDeployServiceV2Request() (request *DeployServiceV2Request) {
    request = &DeployServiceV2Request{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tem", APIVersion, "DeployServiceV2")
    
    
    return
}

func NewDeployServiceV2Response() (response *DeployServiceV2Response) {
    response = &DeployServiceV2Response{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeployServiceV2
// 服务部署
//
// 可能返回的错误码:
//  INTERNALERROR_CREATESERVICEERROR = "InternalError.CreateServiceError"
func (c *Client) DeployServiceV2(request *DeployServiceV2Request) (response *DeployServiceV2Response, err error) {
    if request == nil {
        request = NewDeployServiceV2Request()
    }
    
    response = NewDeployServiceV2Response()
    err = c.Send(request, response)
    return
}

func NewDescribeIngressRequest() (request *DescribeIngressRequest) {
    request = &DescribeIngressRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tem", APIVersion, "DescribeIngress")
    
    
    return
}

func NewDescribeIngressResponse() (response *DescribeIngressResponse) {
    response = &DescribeIngressResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeIngress
// 查询 Ingress 规则
//
// 可能返回的错误码:
//  INTERNALERROR_CREATESERVICEERROR = "InternalError.CreateServiceError"
func (c *Client) DescribeIngress(request *DescribeIngressRequest) (response *DescribeIngressResponse, err error) {
    if request == nil {
        request = NewDescribeIngressRequest()
    }
    
    response = NewDescribeIngressResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeIngressesRequest() (request *DescribeIngressesRequest) {
    request = &DescribeIngressesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tem", APIVersion, "DescribeIngresses")
    
    
    return
}

func NewDescribeIngressesResponse() (response *DescribeIngressesResponse) {
    response = &DescribeIngressesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeIngresses
// 查询 Ingress 规则列表
//
// 可能返回的错误码:
//  INTERNALERROR_CREATESERVICEERROR = "InternalError.CreateServiceError"
func (c *Client) DescribeIngresses(request *DescribeIngressesRequest) (response *DescribeIngressesResponse, err error) {
    if request == nil {
        request = NewDescribeIngressesRequest()
    }
    
    response = NewDescribeIngressesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeNamespacesRequest() (request *DescribeNamespacesRequest) {
    request = &DescribeNamespacesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tem", APIVersion, "DescribeNamespaces")
    
    
    return
}

func NewDescribeNamespacesResponse() (response *DescribeNamespacesResponse) {
    response = &DescribeNamespacesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeNamespaces
// 获取租户环境列表
//
// 可能返回的错误码:
//  INTERNALERROR_CREATESERVICEERROR = "InternalError.CreateServiceError"
func (c *Client) DescribeNamespaces(request *DescribeNamespacesRequest) (response *DescribeNamespacesResponse, err error) {
    if request == nil {
        request = NewDescribeNamespacesRequest()
    }
    
    response = NewDescribeNamespacesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRelatedIngressesRequest() (request *DescribeRelatedIngressesRequest) {
    request = &DescribeRelatedIngressesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tem", APIVersion, "DescribeRelatedIngresses")
    
    
    return
}

func NewDescribeRelatedIngressesResponse() (response *DescribeRelatedIngressesResponse) {
    response = &DescribeRelatedIngressesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeRelatedIngresses
// 查询服务关联的 Ingress 规则列表
//
// 可能返回的错误码:
//  INTERNALERROR_CREATESERVICEERROR = "InternalError.CreateServiceError"
func (c *Client) DescribeRelatedIngresses(request *DescribeRelatedIngressesRequest) (response *DescribeRelatedIngressesResponse, err error) {
    if request == nil {
        request = NewDescribeRelatedIngressesRequest()
    }
    
    response = NewDescribeRelatedIngressesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeServiceRunPodListV2Request() (request *DescribeServiceRunPodListV2Request) {
    request = &DescribeServiceRunPodListV2Request{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tem", APIVersion, "DescribeServiceRunPodListV2")
    
    
    return
}

func NewDescribeServiceRunPodListV2Response() (response *DescribeServiceRunPodListV2Response) {
    response = &DescribeServiceRunPodListV2Response{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeServiceRunPodListV2
// 获取服务下面运行pod列表
//
// 可能返回的错误码:
//  INTERNALERROR_CREATESERVICEERROR = "InternalError.CreateServiceError"
func (c *Client) DescribeServiceRunPodListV2(request *DescribeServiceRunPodListV2Request) (response *DescribeServiceRunPodListV2Response, err error) {
    if request == nil {
        request = NewDescribeServiceRunPodListV2Request()
    }
    
    response = NewDescribeServiceRunPodListV2Response()
    err = c.Send(request, response)
    return
}

func NewGenerateDownloadUrlRequest() (request *GenerateDownloadUrlRequest) {
    request = &GenerateDownloadUrlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tem", APIVersion, "GenerateDownloadUrl")
    
    
    return
}

func NewGenerateDownloadUrlResponse() (response *GenerateDownloadUrlResponse) {
    response = &GenerateDownloadUrlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GenerateDownloadUrl
// 生成包预签名下载链接
//
// 可能返回的错误码:
//  INTERNALERROR_CREATESERVICEERROR = "InternalError.CreateServiceError"
func (c *Client) GenerateDownloadUrl(request *GenerateDownloadUrlRequest) (response *GenerateDownloadUrlResponse, err error) {
    if request == nil {
        request = NewGenerateDownloadUrlRequest()
    }
    
    response = NewGenerateDownloadUrlResponse()
    err = c.Send(request, response)
    return
}

func NewModifyIngressRequest() (request *ModifyIngressRequest) {
    request = &ModifyIngressRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tem", APIVersion, "ModifyIngress")
    
    
    return
}

func NewModifyIngressResponse() (response *ModifyIngressResponse) {
    response = &ModifyIngressResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyIngress
// 创建或者更新 Ingress 规则
//
// 可能返回的错误码:
//  INTERNALERROR_CREATESERVICEERROR = "InternalError.CreateServiceError"
func (c *Client) ModifyIngress(request *ModifyIngressRequest) (response *ModifyIngressResponse, err error) {
    if request == nil {
        request = NewModifyIngressRequest()
    }
    
    response = NewModifyIngressResponse()
    err = c.Send(request, response)
    return
}

func NewModifyNamespaceRequest() (request *ModifyNamespaceRequest) {
    request = &ModifyNamespaceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tem", APIVersion, "ModifyNamespace")
    
    
    return
}

func NewModifyNamespaceResponse() (response *ModifyNamespaceResponse) {
    response = &ModifyNamespaceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyNamespace
// 编辑环境
//
// 可能返回的错误码:
//  INTERNALERROR_CREATESERVICEERROR = "InternalError.CreateServiceError"
func (c *Client) ModifyNamespace(request *ModifyNamespaceRequest) (response *ModifyNamespaceResponse, err error) {
    if request == nil {
        request = NewModifyNamespaceRequest()
    }
    
    response = NewModifyNamespaceResponse()
    err = c.Send(request, response)
    return
}

func NewModifyServiceInfoRequest() (request *ModifyServiceInfoRequest) {
    request = &ModifyServiceInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tem", APIVersion, "ModifyServiceInfo")
    
    
    return
}

func NewModifyServiceInfoResponse() (response *ModifyServiceInfoResponse) {
    response = &ModifyServiceInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyServiceInfo
// 修改服务基本信息
//
// 可能返回的错误码:
//  INTERNALERROR_CREATESERVICEERROR = "InternalError.CreateServiceError"
func (c *Client) ModifyServiceInfo(request *ModifyServiceInfoRequest) (response *ModifyServiceInfoResponse, err error) {
    if request == nil {
        request = NewModifyServiceInfoRequest()
    }
    
    response = NewModifyServiceInfoResponse()
    err = c.Send(request, response)
    return
}

func NewRestartServiceRunPodRequest() (request *RestartServiceRunPodRequest) {
    request = &RestartServiceRunPodRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tem", APIVersion, "RestartServiceRunPod")
    
    
    return
}

func NewRestartServiceRunPodResponse() (response *RestartServiceRunPodResponse) {
    response = &RestartServiceRunPodResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// RestartServiceRunPod
// 重启实例
//
// 可能返回的错误码:
//  INTERNALERROR_CREATESERVICEERROR = "InternalError.CreateServiceError"
func (c *Client) RestartServiceRunPod(request *RestartServiceRunPodRequest) (response *RestartServiceRunPodResponse, err error) {
    if request == nil {
        request = NewRestartServiceRunPodRequest()
    }
    
    response = NewRestartServiceRunPodResponse()
    err = c.Send(request, response)
    return
}
