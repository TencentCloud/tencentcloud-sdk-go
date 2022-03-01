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
    "context"
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
//  INTERNALERROR_ACTIONREADTIMEOUT = "InternalError.ActionReadTimeout"
//  INTERNALERROR_CREATESERVICEERROR = "InternalError.CreateServiceError"
//  INVALIDPARAMETERVALUE_SERVICENAMEDUPLICATEERROR = "InvalidParameterValue.ServiceNameDuplicateError"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
func (c *Client) CreateApplication(request *CreateApplicationRequest) (response *CreateApplicationResponse, err error) {
    if request == nil {
        request = NewCreateApplicationRequest()
    }
    
    response = NewCreateApplicationResponse()
    err = c.Send(request, response)
    return
}

// CreateApplication
// 创建应用
//
// 可能返回的错误码:
//  INTERNALERROR_ACTIONREADTIMEOUT = "InternalError.ActionReadTimeout"
//  INTERNALERROR_CREATESERVICEERROR = "InternalError.CreateServiceError"
//  INVALIDPARAMETERVALUE_SERVICENAMEDUPLICATEERROR = "InvalidParameterValue.ServiceNameDuplicateError"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
func (c *Client) CreateApplicationWithContext(ctx context.Context, request *CreateApplicationRequest) (response *CreateApplicationResponse, err error) {
    if request == nil {
        request = NewCreateApplicationRequest()
    }
    request.SetContext(ctx)
    
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
//  INTERNALERROR_ACTIONREADTIMEOUT = "InternalError.ActionReadTimeout"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
func (c *Client) CreateCosToken(request *CreateCosTokenRequest) (response *CreateCosTokenResponse, err error) {
    if request == nil {
        request = NewCreateCosTokenRequest()
    }
    
    response = NewCreateCosTokenResponse()
    err = c.Send(request, response)
    return
}

// CreateCosToken
// 生成Cos临时秘钥
//
// 可能返回的错误码:
//  INTERNALERROR_ACTIONREADTIMEOUT = "InternalError.ActionReadTimeout"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
func (c *Client) CreateCosTokenWithContext(ctx context.Context, request *CreateCosTokenRequest) (response *CreateCosTokenResponse, err error) {
    if request == nil {
        request = NewCreateCosTokenRequest()
    }
    request.SetContext(ctx)
    
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
//  INTERNALERROR_CREATEEKSCLUSTERERROR = "InternalError.CreateEksClusterError"
//  INTERNALERROR_DEFAULTINTERNALERROR = "InternalError.DefaultInternalError"
//  INVALIDPARAMETERVALUE_NAMESPACEDUPLICATEERROR = "InvalidParameterValue.NamespaceDuplicateError"
func (c *Client) CreateEnvironment(request *CreateEnvironmentRequest) (response *CreateEnvironmentResponse, err error) {
    if request == nil {
        request = NewCreateEnvironmentRequest()
    }
    
    response = NewCreateEnvironmentResponse()
    err = c.Send(request, response)
    return
}

// CreateEnvironment
// 创建环境
//
// 可能返回的错误码:
//  INTERNALERROR_CREATEEKSCLUSTERERROR = "InternalError.CreateEksClusterError"
//  INTERNALERROR_DEFAULTINTERNALERROR = "InternalError.DefaultInternalError"
//  INVALIDPARAMETERVALUE_NAMESPACEDUPLICATEERROR = "InvalidParameterValue.NamespaceDuplicateError"
func (c *Client) CreateEnvironmentWithContext(ctx context.Context, request *CreateEnvironmentRequest) (response *CreateEnvironmentResponse, err error) {
    if request == nil {
        request = NewCreateEnvironmentRequest()
    }
    request.SetContext(ctx)
    
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
//  INVALIDPARAMETERVALUE_NAMESPACEREACHMAXIMUM = "InvalidParameterValue.NamespaceReachMaximum"
//  RESOURCENOTFOUND_NAMESPACENOTFOUND = "ResourceNotFound.NamespaceNotFound"
func (c *Client) CreateResource(request *CreateResourceRequest) (response *CreateResourceResponse, err error) {
    if request == nil {
        request = NewCreateResourceRequest()
    }
    
    response = NewCreateResourceResponse()
    err = c.Send(request, response)
    return
}

// CreateResource
// 绑定云资源
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_NAMESPACEREACHMAXIMUM = "InvalidParameterValue.NamespaceReachMaximum"
//  RESOURCENOTFOUND_NAMESPACENOTFOUND = "ResourceNotFound.NamespaceNotFound"
func (c *Client) CreateResourceWithContext(ctx context.Context, request *CreateResourceRequest) (response *CreateResourceResponse, err error) {
    if request == nil {
        request = NewCreateResourceRequest()
    }
    request.SetContext(ctx)
    
    response = NewCreateResourceResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteApplicationRequest() (request *DeleteApplicationRequest) {
    request = &DeleteApplicationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tem", APIVersion, "DeleteApplication")
    
    
    return
}

func NewDeleteApplicationResponse() (response *DeleteApplicationResponse) {
    response = &DeleteApplicationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteApplication
// 服务删除
//
//   - 停止当前运行服务
//
//   - 删除服务相关资源
//
//   - 删除服务
//
// 可能返回的错误码:
//  INTERNALERROR_ACTIONREADTIMEOUT = "InternalError.ActionReadTimeout"
//  INTERNALERROR_DEFAULTINTERNALERROR = "InternalError.DefaultInternalError"
//  INTERNALERROR_DELETESERVICEERROR = "InternalError.DeleteServiceError"
//  INVALIDPARAMETERVALUE_SERVICEFOUNDRUNNINGVERSION = "InvalidParameterValue.ServiceFoundRunningVersion"
//  INVALIDPARAMETERVALUE_VERSIONROUTERATENOTZERO = "InvalidParameterValue.VersionRouteRateNotZero"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  RESOURCENOTFOUND_SERVICENOTFOUND = "ResourceNotFound.ServiceNotFound"
//  RESOURCENOTFOUND_SERVICERUNNINGVERSIONNOTFOUND = "ResourceNotFound.ServiceRunningVersionNotFound"
//  RESOURCENOTFOUND_VERSIONNAMESPACENOTFOUND = "ResourceNotFound.VersionNamespaceNotFound"
//  RESOURCENOTFOUND_VERSIONSERVICENOTFOUND = "ResourceNotFound.VersionServiceNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DeleteApplication(request *DeleteApplicationRequest) (response *DeleteApplicationResponse, err error) {
    if request == nil {
        request = NewDeleteApplicationRequest()
    }
    
    response = NewDeleteApplicationResponse()
    err = c.Send(request, response)
    return
}

// DeleteApplication
// 服务删除
//
//   - 停止当前运行服务
//
//   - 删除服务相关资源
//
//   - 删除服务
//
// 可能返回的错误码:
//  INTERNALERROR_ACTIONREADTIMEOUT = "InternalError.ActionReadTimeout"
//  INTERNALERROR_DEFAULTINTERNALERROR = "InternalError.DefaultInternalError"
//  INTERNALERROR_DELETESERVICEERROR = "InternalError.DeleteServiceError"
//  INVALIDPARAMETERVALUE_SERVICEFOUNDRUNNINGVERSION = "InvalidParameterValue.ServiceFoundRunningVersion"
//  INVALIDPARAMETERVALUE_VERSIONROUTERATENOTZERO = "InvalidParameterValue.VersionRouteRateNotZero"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  RESOURCENOTFOUND_SERVICENOTFOUND = "ResourceNotFound.ServiceNotFound"
//  RESOURCENOTFOUND_SERVICERUNNINGVERSIONNOTFOUND = "ResourceNotFound.ServiceRunningVersionNotFound"
//  RESOURCENOTFOUND_VERSIONNAMESPACENOTFOUND = "ResourceNotFound.VersionNamespaceNotFound"
//  RESOURCENOTFOUND_VERSIONSERVICENOTFOUND = "ResourceNotFound.VersionServiceNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DeleteApplicationWithContext(ctx context.Context, request *DeleteApplicationRequest) (response *DeleteApplicationResponse, err error) {
    if request == nil {
        request = NewDeleteApplicationRequest()
    }
    request.SetContext(ctx)
    
    response = NewDeleteApplicationResponse()
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
//  INTERNALERROR_ACTIONREADTIMEOUT = "InternalError.ActionReadTimeout"
//  INTERNALERROR_DEFAULTINTERNALERROR = "InternalError.DefaultInternalError"
//  INTERNALERROR_DELETESERVICEERROR = "InternalError.DeleteServiceError"
//  INVALIDPARAMETERVALUE_SERVICEFOUNDRUNNINGVERSION = "InvalidParameterValue.ServiceFoundRunningVersion"
//  INVALIDPARAMETERVALUE_VERSIONROUTERATENOTZERO = "InvalidParameterValue.VersionRouteRateNotZero"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  RESOURCENOTFOUND_SERVICENOTFOUND = "ResourceNotFound.ServiceNotFound"
//  RESOURCENOTFOUND_SERVICERUNNINGVERSIONNOTFOUND = "ResourceNotFound.ServiceRunningVersionNotFound"
//  RESOURCENOTFOUND_VERSIONNAMESPACENOTFOUND = "ResourceNotFound.VersionNamespaceNotFound"
//  RESOURCENOTFOUND_VERSIONSERVICENOTFOUND = "ResourceNotFound.VersionServiceNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DeleteIngress(request *DeleteIngressRequest) (response *DeleteIngressResponse, err error) {
    if request == nil {
        request = NewDeleteIngressRequest()
    }
    
    response = NewDeleteIngressResponse()
    err = c.Send(request, response)
    return
}

// DeleteIngress
// 删除 Ingress 规则
//
// 可能返回的错误码:
//  INTERNALERROR_ACTIONREADTIMEOUT = "InternalError.ActionReadTimeout"
//  INTERNALERROR_DEFAULTINTERNALERROR = "InternalError.DefaultInternalError"
//  INTERNALERROR_DELETESERVICEERROR = "InternalError.DeleteServiceError"
//  INVALIDPARAMETERVALUE_SERVICEFOUNDRUNNINGVERSION = "InvalidParameterValue.ServiceFoundRunningVersion"
//  INVALIDPARAMETERVALUE_VERSIONROUTERATENOTZERO = "InvalidParameterValue.VersionRouteRateNotZero"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  RESOURCENOTFOUND_SERVICENOTFOUND = "ResourceNotFound.ServiceNotFound"
//  RESOURCENOTFOUND_SERVICERUNNINGVERSIONNOTFOUND = "ResourceNotFound.ServiceRunningVersionNotFound"
//  RESOURCENOTFOUND_VERSIONNAMESPACENOTFOUND = "ResourceNotFound.VersionNamespaceNotFound"
//  RESOURCENOTFOUND_VERSIONSERVICENOTFOUND = "ResourceNotFound.VersionServiceNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DeleteIngressWithContext(ctx context.Context, request *DeleteIngressRequest) (response *DeleteIngressResponse, err error) {
    if request == nil {
        request = NewDeleteIngressRequest()
    }
    request.SetContext(ctx)
    
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
//  INTERNALERROR_CREATEAPMRESOURCEERROR = "InternalError.CreateApmResourceError"
//  INTERNALERROR_DEFAULTINTERNALERROR = "InternalError.DefaultInternalError"
//  INTERNALERROR_DEPLOYVERSIONERROR = "InternalError.DeployVersionError"
//  INVALIDPARAMETERVALUE_INVALIDDEPLOYVERSION = "InvalidParameterValue.InvalidDeployVersion"
//  INVALIDPARAMETERVALUE_TRAITSTRACINGNOTSUPPORTED = "InvalidParameterValue.TraitsTracingNotSupported"
//  INVALIDPARAMETERVALUE_VERSIONLOWERCASE = "InvalidParameterValue.VersionLowerCase"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  RESOURCENOTFOUND_VERSIONNAMESPACENOTFOUND = "ResourceNotFound.VersionNamespaceNotFound"
//  RESOURCEUNAVAILABLE_WAITFORKRUISE = "ResourceUnavailable.WaitForKruise"
func (c *Client) DeployApplication(request *DeployApplicationRequest) (response *DeployApplicationResponse, err error) {
    if request == nil {
        request = NewDeployApplicationRequest()
    }
    
    response = NewDeployApplicationResponse()
    err = c.Send(request, response)
    return
}

// DeployApplication
// 应用部署
//
// 可能返回的错误码:
//  INTERNALERROR_CREATEAPMRESOURCEERROR = "InternalError.CreateApmResourceError"
//  INTERNALERROR_DEFAULTINTERNALERROR = "InternalError.DefaultInternalError"
//  INTERNALERROR_DEPLOYVERSIONERROR = "InternalError.DeployVersionError"
//  INVALIDPARAMETERVALUE_INVALIDDEPLOYVERSION = "InvalidParameterValue.InvalidDeployVersion"
//  INVALIDPARAMETERVALUE_TRAITSTRACINGNOTSUPPORTED = "InvalidParameterValue.TraitsTracingNotSupported"
//  INVALIDPARAMETERVALUE_VERSIONLOWERCASE = "InvalidParameterValue.VersionLowerCase"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  RESOURCENOTFOUND_VERSIONNAMESPACENOTFOUND = "ResourceNotFound.VersionNamespaceNotFound"
//  RESOURCEUNAVAILABLE_WAITFORKRUISE = "ResourceUnavailable.WaitForKruise"
func (c *Client) DeployApplicationWithContext(ctx context.Context, request *DeployApplicationRequest) (response *DeployApplicationResponse, err error) {
    if request == nil {
        request = NewDeployApplicationRequest()
    }
    request.SetContext(ctx)
    
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
//  INTERNALERROR_DEFAULTINTERNALERROR = "InternalError.DefaultInternalError"
//  INTERNALERROR_DESCRIBERUNPODLISTERROR = "InternalError.DescribeRunPodListError"
//  RESOURCENOTFOUND_SERVICENOTFOUND = "ResourceNotFound.ServiceNotFound"
//  RESOURCENOTFOUND_SERVICERUNNINGVERSIONNOTFOUND = "ResourceNotFound.ServiceRunningVersionNotFound"
func (c *Client) DescribeApplicationPods(request *DescribeApplicationPodsRequest) (response *DescribeApplicationPodsResponse, err error) {
    if request == nil {
        request = NewDescribeApplicationPodsRequest()
    }
    
    response = NewDescribeApplicationPodsResponse()
    err = c.Send(request, response)
    return
}

// DescribeApplicationPods
// 获取应用实例列表
//
// 可能返回的错误码:
//  INTERNALERROR_DEFAULTINTERNALERROR = "InternalError.DefaultInternalError"
//  INTERNALERROR_DESCRIBERUNPODLISTERROR = "InternalError.DescribeRunPodListError"
//  RESOURCENOTFOUND_SERVICENOTFOUND = "ResourceNotFound.ServiceNotFound"
//  RESOURCENOTFOUND_SERVICERUNNINGVERSIONNOTFOUND = "ResourceNotFound.ServiceRunningVersionNotFound"
func (c *Client) DescribeApplicationPodsWithContext(ctx context.Context, request *DescribeApplicationPodsRequest) (response *DescribeApplicationPodsResponse, err error) {
    if request == nil {
        request = NewDescribeApplicationPodsRequest()
    }
    request.SetContext(ctx)
    
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
//  INTERNALERROR_DESCRIBERUNPODLISTERROR = "InternalError.DescribeRunPodListError"
//  MISSINGPARAMETER_NAMESPACEIDNULL = "MissingParameter.NamespaceIdNull"
//  RESOURCENOTFOUND_SERVICENOTFOUND = "ResourceNotFound.ServiceNotFound"
//  RESOURCENOTFOUND_VERSIONNAMESPACENOTFOUND = "ResourceNotFound.VersionNamespaceNotFound"
func (c *Client) DescribeDeployApplicationDetail(request *DescribeDeployApplicationDetailRequest) (response *DescribeDeployApplicationDetailResponse, err error) {
    if request == nil {
        request = NewDescribeDeployApplicationDetailRequest()
    }
    
    response = NewDescribeDeployApplicationDetailResponse()
    err = c.Send(request, response)
    return
}

// DescribeDeployApplicationDetail
// 获取分批发布详情
//
// 可能返回的错误码:
//  INTERNALERROR_DESCRIBERUNPODLISTERROR = "InternalError.DescribeRunPodListError"
//  MISSINGPARAMETER_NAMESPACEIDNULL = "MissingParameter.NamespaceIdNull"
//  RESOURCENOTFOUND_SERVICENOTFOUND = "ResourceNotFound.ServiceNotFound"
//  RESOURCENOTFOUND_VERSIONNAMESPACENOTFOUND = "ResourceNotFound.VersionNamespaceNotFound"
func (c *Client) DescribeDeployApplicationDetailWithContext(ctx context.Context, request *DescribeDeployApplicationDetailRequest) (response *DescribeDeployApplicationDetailResponse, err error) {
    if request == nil {
        request = NewDescribeDeployApplicationDetailRequest()
    }
    request.SetContext(ctx)
    
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
//  INTERNALERROR_DEFAULTINTERNALERROR = "InternalError.DefaultInternalError"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
func (c *Client) DescribeEnvironments(request *DescribeEnvironmentsRequest) (response *DescribeEnvironmentsResponse, err error) {
    if request == nil {
        request = NewDescribeEnvironmentsRequest()
    }
    
    response = NewDescribeEnvironmentsResponse()
    err = c.Send(request, response)
    return
}

// DescribeEnvironments
// 获取租户环境列表
//
// 可能返回的错误码:
//  INTERNALERROR_DEFAULTINTERNALERROR = "InternalError.DefaultInternalError"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
func (c *Client) DescribeEnvironmentsWithContext(ctx context.Context, request *DescribeEnvironmentsRequest) (response *DescribeEnvironmentsResponse, err error) {
    if request == nil {
        request = NewDescribeEnvironmentsRequest()
    }
    request.SetContext(ctx)
    
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
//  INTERNALERROR_DEFAULTINTERNALERROR = "InternalError.DefaultInternalError"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
func (c *Client) DescribeIngress(request *DescribeIngressRequest) (response *DescribeIngressResponse, err error) {
    if request == nil {
        request = NewDescribeIngressRequest()
    }
    
    response = NewDescribeIngressResponse()
    err = c.Send(request, response)
    return
}

// DescribeIngress
// 查询 Ingress 规则
//
// 可能返回的错误码:
//  INTERNALERROR_DEFAULTINTERNALERROR = "InternalError.DefaultInternalError"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
func (c *Client) DescribeIngressWithContext(ctx context.Context, request *DescribeIngressRequest) (response *DescribeIngressResponse, err error) {
    if request == nil {
        request = NewDescribeIngressRequest()
    }
    request.SetContext(ctx)
    
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
//  MISSINGPARAMETER_NAMESPACEIDNULL = "MissingParameter.NamespaceIdNull"
func (c *Client) DescribeIngresses(request *DescribeIngressesRequest) (response *DescribeIngressesResponse, err error) {
    if request == nil {
        request = NewDescribeIngressesRequest()
    }
    
    response = NewDescribeIngressesResponse()
    err = c.Send(request, response)
    return
}

// DescribeIngresses
// 查询 Ingress 规则列表
//
// 可能返回的错误码:
//  MISSINGPARAMETER_NAMESPACEIDNULL = "MissingParameter.NamespaceIdNull"
func (c *Client) DescribeIngressesWithContext(ctx context.Context, request *DescribeIngressesRequest) (response *DescribeIngressesResponse, err error) {
    if request == nil {
        request = NewDescribeIngressesRequest()
    }
    request.SetContext(ctx)
    
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
//  MISSINGPARAMETER_NAMESPACEIDNULL = "MissingParameter.NamespaceIdNull"
func (c *Client) DescribeRelatedIngresses(request *DescribeRelatedIngressesRequest) (response *DescribeRelatedIngressesResponse, err error) {
    if request == nil {
        request = NewDescribeRelatedIngressesRequest()
    }
    
    response = NewDescribeRelatedIngressesResponse()
    err = c.Send(request, response)
    return
}

// DescribeRelatedIngresses
// 查询应用关联的 Ingress 规则列表
//
// 可能返回的错误码:
//  MISSINGPARAMETER_NAMESPACEIDNULL = "MissingParameter.NamespaceIdNull"
func (c *Client) DescribeRelatedIngressesWithContext(ctx context.Context, request *DescribeRelatedIngressesRequest) (response *DescribeRelatedIngressesResponse, err error) {
    if request == nil {
        request = NewDescribeRelatedIngressesRequest()
    }
    request.SetContext(ctx)
    
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
//  MISSINGPARAMETER_NAMESPACEIDNULL = "MissingParameter.NamespaceIdNull"
func (c *Client) GenerateApplicationPackageDownloadUrl(request *GenerateApplicationPackageDownloadUrlRequest) (response *GenerateApplicationPackageDownloadUrlResponse, err error) {
    if request == nil {
        request = NewGenerateApplicationPackageDownloadUrlRequest()
    }
    
    response = NewGenerateApplicationPackageDownloadUrlResponse()
    err = c.Send(request, response)
    return
}

// GenerateApplicationPackageDownloadUrl
// 生成应用程序包预签名下载链接
//
// 可能返回的错误码:
//  MISSINGPARAMETER_NAMESPACEIDNULL = "MissingParameter.NamespaceIdNull"
func (c *Client) GenerateApplicationPackageDownloadUrlWithContext(ctx context.Context, request *GenerateApplicationPackageDownloadUrlRequest) (response *GenerateApplicationPackageDownloadUrlResponse, err error) {
    if request == nil {
        request = NewGenerateApplicationPackageDownloadUrlRequest()
    }
    request.SetContext(ctx)
    
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

// ModifyApplicationInfo
// 修改应用基本信息
//
// 可能返回的错误码:
//  INTERNALERROR_CREATESERVICEERROR = "InternalError.CreateServiceError"
func (c *Client) ModifyApplicationInfoWithContext(ctx context.Context, request *ModifyApplicationInfoRequest) (response *ModifyApplicationInfoResponse, err error) {
    if request == nil {
        request = NewModifyApplicationInfoRequest()
    }
    request.SetContext(ctx)
    
    response = NewModifyApplicationInfoResponse()
    err = c.Send(request, response)
    return
}

func NewModifyApplicationReplicasRequest() (request *ModifyApplicationReplicasRequest) {
    request = &ModifyApplicationReplicasRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tem", APIVersion, "ModifyApplicationReplicas")
    
    
    return
}

func NewModifyApplicationReplicasResponse() (response *ModifyApplicationReplicasResponse) {
    response = &ModifyApplicationReplicasResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyApplicationReplicas
// 修改应用实例数量
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_SERVICERUNNINGVERSIONNOTFOUND = "ResourceNotFound.ServiceRunningVersionNotFound"
func (c *Client) ModifyApplicationReplicas(request *ModifyApplicationReplicasRequest) (response *ModifyApplicationReplicasResponse, err error) {
    if request == nil {
        request = NewModifyApplicationReplicasRequest()
    }
    
    response = NewModifyApplicationReplicasResponse()
    err = c.Send(request, response)
    return
}

// ModifyApplicationReplicas
// 修改应用实例数量
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_SERVICERUNNINGVERSIONNOTFOUND = "ResourceNotFound.ServiceRunningVersionNotFound"
func (c *Client) ModifyApplicationReplicasWithContext(ctx context.Context, request *ModifyApplicationReplicasRequest) (response *ModifyApplicationReplicasResponse, err error) {
    if request == nil {
        request = NewModifyApplicationReplicasRequest()
    }
    request.SetContext(ctx)
    
    response = NewModifyApplicationReplicasResponse()
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
//  INTERNALERROR_DEFAULTINTERNALERROR = "InternalError.DefaultInternalError"
func (c *Client) ModifyEnvironment(request *ModifyEnvironmentRequest) (response *ModifyEnvironmentResponse, err error) {
    if request == nil {
        request = NewModifyEnvironmentRequest()
    }
    
    response = NewModifyEnvironmentResponse()
    err = c.Send(request, response)
    return
}

// ModifyEnvironment
// 编辑环境
//
// 可能返回的错误码:
//  INTERNALERROR_DEFAULTINTERNALERROR = "InternalError.DefaultInternalError"
func (c *Client) ModifyEnvironmentWithContext(ctx context.Context, request *ModifyEnvironmentRequest) (response *ModifyEnvironmentResponse, err error) {
    if request == nil {
        request = NewModifyEnvironmentRequest()
    }
    request.SetContext(ctx)
    
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
//  INTERNALERROR_UPDATEINGRESSERROR = "InternalError.UpdateIngressError"
func (c *Client) ModifyIngress(request *ModifyIngressRequest) (response *ModifyIngressResponse, err error) {
    if request == nil {
        request = NewModifyIngressRequest()
    }
    
    response = NewModifyIngressResponse()
    err = c.Send(request, response)
    return
}

// ModifyIngress
// 创建或者更新 Ingress 规则
//
// 可能返回的错误码:
//  INTERNALERROR_UPDATEINGRESSERROR = "InternalError.UpdateIngressError"
func (c *Client) ModifyIngressWithContext(ctx context.Context, request *ModifyIngressRequest) (response *ModifyIngressResponse, err error) {
    if request == nil {
        request = NewModifyIngressRequest()
    }
    request.SetContext(ctx)
    
    response = NewModifyIngressResponse()
    err = c.Send(request, response)
    return
}

func NewRestartApplicationRequest() (request *RestartApplicationRequest) {
    request = &RestartApplicationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tem", APIVersion, "RestartApplication")
    
    
    return
}

func NewRestartApplicationResponse() (response *RestartApplicationResponse) {
    response = &RestartApplicationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// RestartApplication
// 服务重启
//
// 可能返回的错误码:
//  INTERNALERROR_RESTARTAPPLICATIONERROR = "InternalError.RestartApplicationError"
//  MISSINGPARAMETER_NAMESPACEIDNULL = "MissingParameter.NamespaceIdNull"
func (c *Client) RestartApplication(request *RestartApplicationRequest) (response *RestartApplicationResponse, err error) {
    if request == nil {
        request = NewRestartApplicationRequest()
    }
    
    response = NewRestartApplicationResponse()
    err = c.Send(request, response)
    return
}

// RestartApplication
// 服务重启
//
// 可能返回的错误码:
//  INTERNALERROR_RESTARTAPPLICATIONERROR = "InternalError.RestartApplicationError"
//  MISSINGPARAMETER_NAMESPACEIDNULL = "MissingParameter.NamespaceIdNull"
func (c *Client) RestartApplicationWithContext(ctx context.Context, request *RestartApplicationRequest) (response *RestartApplicationResponse, err error) {
    if request == nil {
        request = NewRestartApplicationRequest()
    }
    request.SetContext(ctx)
    
    response = NewRestartApplicationResponse()
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
//  INTERNALERROR_RESTARTAPPLICATIONERROR = "InternalError.RestartApplicationError"
//  MISSINGPARAMETER_NAMESPACEIDNULL = "MissingParameter.NamespaceIdNull"
func (c *Client) RestartApplicationPod(request *RestartApplicationPodRequest) (response *RestartApplicationPodResponse, err error) {
    if request == nil {
        request = NewRestartApplicationPodRequest()
    }
    
    response = NewRestartApplicationPodResponse()
    err = c.Send(request, response)
    return
}

// RestartApplicationPod
// 重启应用实例
//
// 可能返回的错误码:
//  INTERNALERROR_RESTARTAPPLICATIONERROR = "InternalError.RestartApplicationError"
//  MISSINGPARAMETER_NAMESPACEIDNULL = "MissingParameter.NamespaceIdNull"
func (c *Client) RestartApplicationPodWithContext(ctx context.Context, request *RestartApplicationPodRequest) (response *RestartApplicationPodResponse, err error) {
    if request == nil {
        request = NewRestartApplicationPodRequest()
    }
    request.SetContext(ctx)
    
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
//  INTERNALERROR_RESTARTAPPLICATIONERROR = "InternalError.RestartApplicationError"
//  MISSINGPARAMETER_NAMESPACEIDNULL = "MissingParameter.NamespaceIdNull"
func (c *Client) ResumeDeployApplication(request *ResumeDeployApplicationRequest) (response *ResumeDeployApplicationResponse, err error) {
    if request == nil {
        request = NewResumeDeployApplicationRequest()
    }
    
    response = NewResumeDeployApplicationResponse()
    err = c.Send(request, response)
    return
}

// ResumeDeployApplication
// 开始下一批次发布
//
// 可能返回的错误码:
//  INTERNALERROR_RESTARTAPPLICATIONERROR = "InternalError.RestartApplicationError"
//  MISSINGPARAMETER_NAMESPACEIDNULL = "MissingParameter.NamespaceIdNull"
func (c *Client) ResumeDeployApplicationWithContext(ctx context.Context, request *ResumeDeployApplicationRequest) (response *ResumeDeployApplicationResponse, err error) {
    if request == nil {
        request = NewResumeDeployApplicationRequest()
    }
    request.SetContext(ctx)
    
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
//  INTERNALERROR_RESTARTAPPLICATIONERROR = "InternalError.RestartApplicationError"
//  MISSINGPARAMETER_NAMESPACEIDNULL = "MissingParameter.NamespaceIdNull"
func (c *Client) RevertDeployApplication(request *RevertDeployApplicationRequest) (response *RevertDeployApplicationResponse, err error) {
    if request == nil {
        request = NewRevertDeployApplicationRequest()
    }
    
    response = NewRevertDeployApplicationResponse()
    err = c.Send(request, response)
    return
}

// RevertDeployApplication
// 回滚分批发布
//
// 可能返回的错误码:
//  INTERNALERROR_RESTARTAPPLICATIONERROR = "InternalError.RestartApplicationError"
//  MISSINGPARAMETER_NAMESPACEIDNULL = "MissingParameter.NamespaceIdNull"
func (c *Client) RevertDeployApplicationWithContext(ctx context.Context, request *RevertDeployApplicationRequest) (response *RevertDeployApplicationResponse, err error) {
    if request == nil {
        request = NewRevertDeployApplicationRequest()
    }
    request.SetContext(ctx)
    
    response = NewRevertDeployApplicationResponse()
    err = c.Send(request, response)
    return
}

func NewRollingUpdateApplicationByVersionRequest() (request *RollingUpdateApplicationByVersionRequest) {
    request = &RollingUpdateApplicationByVersionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tem", APIVersion, "RollingUpdateApplicationByVersion")
    
    
    return
}

func NewRollingUpdateApplicationByVersionResponse() (response *RollingUpdateApplicationByVersionResponse) {
    response = &RollingUpdateApplicationByVersionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// RollingUpdateApplicationByVersion
// 更新应用部署版本
//
// 可能返回的错误码:
//  INTERNALERROR_DEFAULTINTERNALERROR = "InternalError.DefaultInternalError"
//  INTERNALERROR_DEPLOYVERSIONERROR = "InternalError.DeployVersionError"
//  INVALIDPARAMETERVALUE_INVALIDDEPLOYVERSION = "InvalidParameterValue.InvalidDeployVersion"
//  INVALIDPARAMETERVALUE_VERSIONLOWERCASE = "InvalidParameterValue.VersionLowerCase"
//  RESOURCENOTFOUND_SERVICENOTFOUND = "ResourceNotFound.ServiceNotFound"
//  RESOURCENOTFOUND_VERSIONNAMESPACENOTFOUND = "ResourceNotFound.VersionNamespaceNotFound"
func (c *Client) RollingUpdateApplicationByVersion(request *RollingUpdateApplicationByVersionRequest) (response *RollingUpdateApplicationByVersionResponse, err error) {
    if request == nil {
        request = NewRollingUpdateApplicationByVersionRequest()
    }
    
    response = NewRollingUpdateApplicationByVersionResponse()
    err = c.Send(request, response)
    return
}

// RollingUpdateApplicationByVersion
// 更新应用部署版本
//
// 可能返回的错误码:
//  INTERNALERROR_DEFAULTINTERNALERROR = "InternalError.DefaultInternalError"
//  INTERNALERROR_DEPLOYVERSIONERROR = "InternalError.DeployVersionError"
//  INVALIDPARAMETERVALUE_INVALIDDEPLOYVERSION = "InvalidParameterValue.InvalidDeployVersion"
//  INVALIDPARAMETERVALUE_VERSIONLOWERCASE = "InvalidParameterValue.VersionLowerCase"
//  RESOURCENOTFOUND_SERVICENOTFOUND = "ResourceNotFound.ServiceNotFound"
//  RESOURCENOTFOUND_VERSIONNAMESPACENOTFOUND = "ResourceNotFound.VersionNamespaceNotFound"
func (c *Client) RollingUpdateApplicationByVersionWithContext(ctx context.Context, request *RollingUpdateApplicationByVersionRequest) (response *RollingUpdateApplicationByVersionResponse, err error) {
    if request == nil {
        request = NewRollingUpdateApplicationByVersionRequest()
    }
    request.SetContext(ctx)
    
    response = NewRollingUpdateApplicationByVersionResponse()
    err = c.Send(request, response)
    return
}

func NewStopApplicationRequest() (request *StopApplicationRequest) {
    request = &StopApplicationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tem", APIVersion, "StopApplication")
    
    
    return
}

func NewStopApplicationResponse() (response *StopApplicationResponse) {
    response = &StopApplicationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// StopApplication
// 服务停止
//
// 可能返回的错误码:
//  MISSINGPARAMETER_NAMESPACEIDNULL = "MissingParameter.NamespaceIdNull"
//  RESOURCENOTFOUND_SERVICERUNNINGVERSIONNOTFOUND = "ResourceNotFound.ServiceRunningVersionNotFound"
func (c *Client) StopApplication(request *StopApplicationRequest) (response *StopApplicationResponse, err error) {
    if request == nil {
        request = NewStopApplicationRequest()
    }
    
    response = NewStopApplicationResponse()
    err = c.Send(request, response)
    return
}

// StopApplication
// 服务停止
//
// 可能返回的错误码:
//  MISSINGPARAMETER_NAMESPACEIDNULL = "MissingParameter.NamespaceIdNull"
//  RESOURCENOTFOUND_SERVICERUNNINGVERSIONNOTFOUND = "ResourceNotFound.ServiceRunningVersionNotFound"
func (c *Client) StopApplicationWithContext(ctx context.Context, request *StopApplicationRequest) (response *StopApplicationResponse, err error) {
    if request == nil {
        request = NewStopApplicationRequest()
    }
    request.SetContext(ctx)
    
    response = NewStopApplicationResponse()
    err = c.Send(request, response)
    return
}
