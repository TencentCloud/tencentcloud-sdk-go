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

package v20210701

import (
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2021-07-01"

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


func NewCreateApplicationRequest() (request *CreateApplicationRequest) {
    request = &CreateApplicationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tem", APIVersion, "CreateApplication")
    return
}

func NewCreateApplicationResponse() (response *CreateApplicationResponse) {
    response = &CreateApplicationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateApplication
// 创建应用
//
// 可能返回的错误码:
//  INTERNALERROR_CREATESERVICEERROR = "InternalError.CreateServiceError"
func (c *Client) CreateApplication(request *CreateApplicationRequest) (response *CreateApplicationResponse, err error) {
    if request == nil {
        request = NewCreateApplicationRequest()
    }
    response = NewCreateApplicationResponse()
    err = c.Send(request, response)
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
//
// 可能返回的错误码:
//  INTERNALERROR_CREATESERVICEERROR = "InternalError.CreateServiceError"
func (c *Client) CreateCosToken(request *CreateCosTokenRequest) (response *CreateCosTokenResponse, err error) {
    if request == nil {
        request = NewCreateCosTokenRequest()
    }
    response = NewCreateCosTokenResponse()
    err = c.Send(request, response)
    return
}

func NewCreateEnvironmentRequest() (request *CreateEnvironmentRequest) {
    request = &CreateEnvironmentRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tem", APIVersion, "CreateEnvironment")
    return
}

func NewCreateEnvironmentResponse() (response *CreateEnvironmentResponse) {
    response = &CreateEnvironmentResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateEnvironment
// 创建环境
//
// 可能返回的错误码:
//  INTERNALERROR_CREATESERVICEERROR = "InternalError.CreateServiceError"
func (c *Client) CreateEnvironment(request *CreateEnvironmentRequest) (response *CreateEnvironmentResponse, err error) {
    if request == nil {
        request = NewCreateEnvironmentRequest()
    }
    response = NewCreateEnvironmentResponse()
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
//
// 可能返回的错误码:
//  INTERNALERROR_CREATESERVICEERROR = "InternalError.CreateServiceError"
func (c *Client) CreateResource(request *CreateResourceRequest) (response *CreateResourceResponse, err error) {
    if request == nil {
        request = NewCreateResourceRequest()
    }
    response = NewCreateResourceResponse()
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

func NewDeployApplicationRequest() (request *DeployApplicationRequest) {
    request = &DeployApplicationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tem", APIVersion, "DeployApplication")
    return
}

func NewDeployApplicationResponse() (response *DeployApplicationResponse) {
    response = &DeployApplicationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeployApplication
// 应用部署
//
// 可能返回的错误码:
//  INTERNALERROR_CREATESERVICEERROR = "InternalError.CreateServiceError"
func (c *Client) DeployApplication(request *DeployApplicationRequest) (response *DeployApplicationResponse, err error) {
    if request == nil {
        request = NewDeployApplicationRequest()
    }
    response = NewDeployApplicationResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeApplicationPodsRequest() (request *DescribeApplicationPodsRequest) {
    request = &DescribeApplicationPodsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tem", APIVersion, "DescribeApplicationPods")
    return
}

func NewDescribeApplicationPodsResponse() (response *DescribeApplicationPodsResponse) {
    response = &DescribeApplicationPodsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeApplicationPods
// 获取应用实例列表
//
// 可能返回的错误码:
//  INTERNALERROR_CREATESERVICEERROR = "InternalError.CreateServiceError"
func (c *Client) DescribeApplicationPods(request *DescribeApplicationPodsRequest) (response *DescribeApplicationPodsResponse, err error) {
    if request == nil {
        request = NewDescribeApplicationPodsRequest()
    }
    response = NewDescribeApplicationPodsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDeployApplicationDetailRequest() (request *DescribeDeployApplicationDetailRequest) {
    request = &DescribeDeployApplicationDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tem", APIVersion, "DescribeDeployApplicationDetail")
    return
}

func NewDescribeDeployApplicationDetailResponse() (response *DescribeDeployApplicationDetailResponse) {
    response = &DescribeDeployApplicationDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDeployApplicationDetail
// 获取分批发布详情
//
// 可能返回的错误码:
//  INTERNALERROR_CREATESERVICEERROR = "InternalError.CreateServiceError"
func (c *Client) DescribeDeployApplicationDetail(request *DescribeDeployApplicationDetailRequest) (response *DescribeDeployApplicationDetailResponse, err error) {
    if request == nil {
        request = NewDescribeDeployApplicationDetailRequest()
    }
    response = NewDescribeDeployApplicationDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEnvironmentsRequest() (request *DescribeEnvironmentsRequest) {
    request = &DescribeEnvironmentsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tem", APIVersion, "DescribeEnvironments")
    return
}

func NewDescribeEnvironmentsResponse() (response *DescribeEnvironmentsResponse) {
    response = &DescribeEnvironmentsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeEnvironments
// 获取租户环境列表
//
// 可能返回的错误码:
//  INTERNALERROR_CREATESERVICEERROR = "InternalError.CreateServiceError"
func (c *Client) DescribeEnvironments(request *DescribeEnvironmentsRequest) (response *DescribeEnvironmentsResponse, err error) {
    if request == nil {
        request = NewDescribeEnvironmentsRequest()
    }
    response = NewDescribeEnvironmentsResponse()
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
// 查询应用关联的 Ingress 规则列表
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

func NewGenerateApplicationPackageDownloadUrlRequest() (request *GenerateApplicationPackageDownloadUrlRequest) {
    request = &GenerateApplicationPackageDownloadUrlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tem", APIVersion, "GenerateApplicationPackageDownloadUrl")
    return
}

func NewGenerateApplicationPackageDownloadUrlResponse() (response *GenerateApplicationPackageDownloadUrlResponse) {
    response = &GenerateApplicationPackageDownloadUrlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GenerateApplicationPackageDownloadUrl
// 生成应用程序包预签名下载链接
//
// 可能返回的错误码:
//  INTERNALERROR_CREATESERVICEERROR = "InternalError.CreateServiceError"
func (c *Client) GenerateApplicationPackageDownloadUrl(request *GenerateApplicationPackageDownloadUrlRequest) (response *GenerateApplicationPackageDownloadUrlResponse, err error) {
    if request == nil {
        request = NewGenerateApplicationPackageDownloadUrlRequest()
    }
    response = NewGenerateApplicationPackageDownloadUrlResponse()
    err = c.Send(request, response)
    return
}

func NewModifyApplicationInfoRequest() (request *ModifyApplicationInfoRequest) {
    request = &ModifyApplicationInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tem", APIVersion, "ModifyApplicationInfo")
    return
}

func NewModifyApplicationInfoResponse() (response *ModifyApplicationInfoResponse) {
    response = &ModifyApplicationInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyApplicationInfo
// 修改应用基本信息
//
// 可能返回的错误码:
//  INTERNALERROR_CREATESERVICEERROR = "InternalError.CreateServiceError"
func (c *Client) ModifyApplicationInfo(request *ModifyApplicationInfoRequest) (response *ModifyApplicationInfoResponse, err error) {
    if request == nil {
        request = NewModifyApplicationInfoRequest()
    }
    response = NewModifyApplicationInfoResponse()
    err = c.Send(request, response)
    return
}

func NewModifyEnvironmentRequest() (request *ModifyEnvironmentRequest) {
    request = &ModifyEnvironmentRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tem", APIVersion, "ModifyEnvironment")
    return
}

func NewModifyEnvironmentResponse() (response *ModifyEnvironmentResponse) {
    response = &ModifyEnvironmentResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyEnvironment
// 编辑环境
//
// 可能返回的错误码:
//  INTERNALERROR_CREATESERVICEERROR = "InternalError.CreateServiceError"
func (c *Client) ModifyEnvironment(request *ModifyEnvironmentRequest) (response *ModifyEnvironmentResponse, err error) {
    if request == nil {
        request = NewModifyEnvironmentRequest()
    }
    response = NewModifyEnvironmentResponse()
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

func NewRestartApplicationPodRequest() (request *RestartApplicationPodRequest) {
    request = &RestartApplicationPodRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tem", APIVersion, "RestartApplicationPod")
    return
}

func NewRestartApplicationPodResponse() (response *RestartApplicationPodResponse) {
    response = &RestartApplicationPodResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// RestartApplicationPod
// 重启应用实例
//
// 可能返回的错误码:
//  INTERNALERROR_CREATESERVICEERROR = "InternalError.CreateServiceError"
func (c *Client) RestartApplicationPod(request *RestartApplicationPodRequest) (response *RestartApplicationPodResponse, err error) {
    if request == nil {
        request = NewRestartApplicationPodRequest()
    }
    response = NewRestartApplicationPodResponse()
    err = c.Send(request, response)
    return
}

func NewResumeDeployApplicationRequest() (request *ResumeDeployApplicationRequest) {
    request = &ResumeDeployApplicationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tem", APIVersion, "ResumeDeployApplication")
    return
}

func NewResumeDeployApplicationResponse() (response *ResumeDeployApplicationResponse) {
    response = &ResumeDeployApplicationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ResumeDeployApplication
// 开始下一批次发布
//
// 可能返回的错误码:
//  INTERNALERROR_CREATESERVICEERROR = "InternalError.CreateServiceError"
func (c *Client) ResumeDeployApplication(request *ResumeDeployApplicationRequest) (response *ResumeDeployApplicationResponse, err error) {
    if request == nil {
        request = NewResumeDeployApplicationRequest()
    }
    response = NewResumeDeployApplicationResponse()
    err = c.Send(request, response)
    return
}

func NewRevertDeployApplicationRequest() (request *RevertDeployApplicationRequest) {
    request = &RevertDeployApplicationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tem", APIVersion, "RevertDeployApplication")
    return
}

func NewRevertDeployApplicationResponse() (response *RevertDeployApplicationResponse) {
    response = &RevertDeployApplicationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// RevertDeployApplication
// 回滚分批发布
//
// 可能返回的错误码:
//  INTERNALERROR_CREATESERVICEERROR = "InternalError.CreateServiceError"
func (c *Client) RevertDeployApplication(request *RevertDeployApplicationRequest) (response *RevertDeployApplicationResponse, err error) {
    if request == nil {
        request = NewRevertDeployApplicationRequest()
    }
    response = NewRevertDeployApplicationResponse()
    err = c.Send(request, response)
    return
}
