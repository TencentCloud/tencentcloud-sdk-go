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
    "errors"
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
//  INTERNALERROR_DEFAULTINTERNALERROR = "InternalError.DefaultInternalError"
//  INVALIDPARAMETERVALUE_INVALIDSERVICENAME = "InvalidParameterValue.InvalidServiceName"
//  INVALIDPARAMETERVALUE_PUBLICREPOTYPEPARAMETERERROR = "InvalidParameterValue.PublicRepoTypeParameterError"
//  INVALIDPARAMETERVALUE_SERVICELOWERCASE = "InvalidParameterValue.ServiceLowerCase"
//  INVALIDPARAMETERVALUE_SERVICENAMEDUPLICATEERROR = "InvalidParameterValue.ServiceNameDuplicateError"
//  INVALIDPARAMETERVALUE_SERVICEREACHMAXIMUM = "InvalidParameterValue.ServiceReachMaximum"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNSUPPORTEDOPERATION_UNSUPPORTACTION = "UnsupportedOperation.UnsupportAction"
func (c *Client) CreateApplication(request *CreateApplicationRequest) (response *CreateApplicationResponse, err error) {
    return c.CreateApplicationWithContext(context.Background(), request)
}

// CreateApplication
// 创建应用
//
// 可能返回的错误码:
//  INTERNALERROR_ACTIONREADTIMEOUT = "InternalError.ActionReadTimeout"
//  INTERNALERROR_CREATESERVICEERROR = "InternalError.CreateServiceError"
//  INTERNALERROR_DEFAULTINTERNALERROR = "InternalError.DefaultInternalError"
//  INVALIDPARAMETERVALUE_INVALIDSERVICENAME = "InvalidParameterValue.InvalidServiceName"
//  INVALIDPARAMETERVALUE_PUBLICREPOTYPEPARAMETERERROR = "InvalidParameterValue.PublicRepoTypeParameterError"
//  INVALIDPARAMETERVALUE_SERVICELOWERCASE = "InvalidParameterValue.ServiceLowerCase"
//  INVALIDPARAMETERVALUE_SERVICENAMEDUPLICATEERROR = "InvalidParameterValue.ServiceNameDuplicateError"
//  INVALIDPARAMETERVALUE_SERVICEREACHMAXIMUM = "InvalidParameterValue.ServiceReachMaximum"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNSUPPORTEDOPERATION_UNSUPPORTACTION = "UnsupportedOperation.UnsupportAction"
func (c *Client) CreateApplicationWithContext(ctx context.Context, request *CreateApplicationRequest) (response *CreateApplicationResponse, err error) {
    if request == nil {
        request = NewCreateApplicationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateApplication require credential")
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
    return c.CreateCosTokenWithContext(context.Background(), request)
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCosToken require credential")
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
    return c.CreateEnvironmentWithContext(context.Background(), request)
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateEnvironment require credential")
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
//  INVALIDPARAMETERVALUE_NAMESPACERESOURCEREACHMAXIMUM = "InvalidParameterValue.NamespaceResourceReachMaximum"
//  RESOURCEINUSE_RESOURCEALREADYUSED = "ResourceInUse.ResourceAlreadyUsed"
//  RESOURCENOTFOUND_NAMESPACENOTFOUND = "ResourceNotFound.NamespaceNotFound"
func (c *Client) CreateResource(request *CreateResourceRequest) (response *CreateResourceResponse, err error) {
    return c.CreateResourceWithContext(context.Background(), request)
}

// CreateResource
// 绑定云资源
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_NAMESPACEREACHMAXIMUM = "InvalidParameterValue.NamespaceReachMaximum"
//  INVALIDPARAMETERVALUE_NAMESPACERESOURCEREACHMAXIMUM = "InvalidParameterValue.NamespaceResourceReachMaximum"
//  RESOURCEINUSE_RESOURCEALREADYUSED = "ResourceInUse.ResourceAlreadyUsed"
//  RESOURCENOTFOUND_NAMESPACENOTFOUND = "ResourceNotFound.NamespaceNotFound"
func (c *Client) CreateResourceWithContext(ctx context.Context, request *CreateResourceRequest) (response *CreateResourceResponse, err error) {
    if request == nil {
        request = NewCreateResourceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateResource require credential")
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
//  INVALIDPARAMETERVALUE_NAMESPACENOTBELONGTOAPPID = "InvalidParameterValue.NamespaceNotBelongToAppid"
//  INVALIDPARAMETERVALUE_SERVICEFOUNDRUNNINGVERSION = "InvalidParameterValue.ServiceFoundRunningVersion"
//  INVALIDPARAMETERVALUE_VERSIONROUTERATENOTZERO = "InvalidParameterValue.VersionRouteRateNotZero"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  RESOURCENOTFOUND_SERVICENOTFOUND = "ResourceNotFound.ServiceNotFound"
//  RESOURCENOTFOUND_SERVICERUNNINGVERSIONNOTFOUND = "ResourceNotFound.ServiceRunningVersionNotFound"
//  RESOURCENOTFOUND_VERSIONNAMESPACENOTFOUND = "ResourceNotFound.VersionNamespaceNotFound"
//  RESOURCENOTFOUND_VERSIONSERVICENOTFOUND = "ResourceNotFound.VersionServiceNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DeleteApplication(request *DeleteApplicationRequest) (response *DeleteApplicationResponse, err error) {
    return c.DeleteApplicationWithContext(context.Background(), request)
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
//  INVALIDPARAMETERVALUE_NAMESPACENOTBELONGTOAPPID = "InvalidParameterValue.NamespaceNotBelongToAppid"
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteApplication require credential")
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
//  INVALIDPARAMETERVALUE_NAMESPACENOTBELONGTOAPPID = "InvalidParameterValue.NamespaceNotBelongToAppid"
//  INVALIDPARAMETERVALUE_SERVICEFOUNDRUNNINGVERSION = "InvalidParameterValue.ServiceFoundRunningVersion"
//  INVALIDPARAMETERVALUE_VERSIONROUTERATENOTZERO = "InvalidParameterValue.VersionRouteRateNotZero"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  RESOURCENOTFOUND_SERVICENOTFOUND = "ResourceNotFound.ServiceNotFound"
//  RESOURCENOTFOUND_SERVICERUNNINGVERSIONNOTFOUND = "ResourceNotFound.ServiceRunningVersionNotFound"
//  RESOURCENOTFOUND_VERSIONNAMESPACENOTFOUND = "ResourceNotFound.VersionNamespaceNotFound"
//  RESOURCENOTFOUND_VERSIONSERVICENOTFOUND = "ResourceNotFound.VersionServiceNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DeleteIngress(request *DeleteIngressRequest) (response *DeleteIngressResponse, err error) {
    return c.DeleteIngressWithContext(context.Background(), request)
}

// DeleteIngress
// 删除 Ingress 规则
//
// 可能返回的错误码:
//  INTERNALERROR_ACTIONREADTIMEOUT = "InternalError.ActionReadTimeout"
//  INTERNALERROR_DEFAULTINTERNALERROR = "InternalError.DefaultInternalError"
//  INTERNALERROR_DELETESERVICEERROR = "InternalError.DeleteServiceError"
//  INVALIDPARAMETERVALUE_NAMESPACENOTBELONGTOAPPID = "InvalidParameterValue.NamespaceNotBelongToAppid"
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteIngress require credential")
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
//  INVALIDPARAMETERVALUE_AUTOSCALERLARGERTHANONE = "InvalidParameterValue.AutoScalerLargerThanOne"
//  INVALIDPARAMETERVALUE_INVALIDDEPLOYVERSION = "InvalidParameterValue.InvalidDeployVersion"
//  INVALIDPARAMETERVALUE_TRAITSTRACINGNOTSUPPORTED = "InvalidParameterValue.TraitsTracingNotSupported"
//  INVALIDPARAMETERVALUE_VERSIONLOWERCASE = "InvalidParameterValue.VersionLowerCase"
//  MISSINGPARAMETER_DEPLOYVERSIONNULL = "MissingParameter.DeployVersionNull"
//  MISSINGPARAMETER_PKGNAMENULL = "MissingParameter.PkgNameNull"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  RESOURCENOTFOUND_SERVICENOTFOUND = "ResourceNotFound.ServiceNotFound"
//  RESOURCENOTFOUND_VERSIONNAMESPACENOTFOUND = "ResourceNotFound.VersionNamespaceNotFound"
//  RESOURCEUNAVAILABLE_WAITFORKRUISE = "ResourceUnavailable.WaitForKruise"
func (c *Client) DeployApplication(request *DeployApplicationRequest) (response *DeployApplicationResponse, err error) {
    return c.DeployApplicationWithContext(context.Background(), request)
}

// DeployApplication
// 应用部署
//
// 可能返回的错误码:
//  INTERNALERROR_CREATEAPMRESOURCEERROR = "InternalError.CreateApmResourceError"
//  INTERNALERROR_DEFAULTINTERNALERROR = "InternalError.DefaultInternalError"
//  INTERNALERROR_DEPLOYVERSIONERROR = "InternalError.DeployVersionError"
//  INVALIDPARAMETERVALUE_AUTOSCALERLARGERTHANONE = "InvalidParameterValue.AutoScalerLargerThanOne"
//  INVALIDPARAMETERVALUE_INVALIDDEPLOYVERSION = "InvalidParameterValue.InvalidDeployVersion"
//  INVALIDPARAMETERVALUE_TRAITSTRACINGNOTSUPPORTED = "InvalidParameterValue.TraitsTracingNotSupported"
//  INVALIDPARAMETERVALUE_VERSIONLOWERCASE = "InvalidParameterValue.VersionLowerCase"
//  MISSINGPARAMETER_DEPLOYVERSIONNULL = "MissingParameter.DeployVersionNull"
//  MISSINGPARAMETER_PKGNAMENULL = "MissingParameter.PkgNameNull"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  RESOURCENOTFOUND_SERVICENOTFOUND = "ResourceNotFound.ServiceNotFound"
//  RESOURCENOTFOUND_VERSIONNAMESPACENOTFOUND = "ResourceNotFound.VersionNamespaceNotFound"
//  RESOURCEUNAVAILABLE_WAITFORKRUISE = "ResourceUnavailable.WaitForKruise"
func (c *Client) DeployApplicationWithContext(ctx context.Context, request *DeployApplicationRequest) (response *DeployApplicationResponse, err error) {
    if request == nil {
        request = NewDeployApplicationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeployApplication require credential")
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
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  RESOURCENOTFOUND_SERVICENOTFOUND = "ResourceNotFound.ServiceNotFound"
//  RESOURCENOTFOUND_SERVICERUNNINGVERSIONNOTFOUND = "ResourceNotFound.ServiceRunningVersionNotFound"
//  RESOURCENOTFOUND_VERSIONNAMESPACENOTFOUND = "ResourceNotFound.VersionNamespaceNotFound"
func (c *Client) DescribeApplicationPods(request *DescribeApplicationPodsRequest) (response *DescribeApplicationPodsResponse, err error) {
    return c.DescribeApplicationPodsWithContext(context.Background(), request)
}

// DescribeApplicationPods
// 获取应用实例列表
//
// 可能返回的错误码:
//  INTERNALERROR_DEFAULTINTERNALERROR = "InternalError.DefaultInternalError"
//  INTERNALERROR_DESCRIBERUNPODLISTERROR = "InternalError.DescribeRunPodListError"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  RESOURCENOTFOUND_SERVICENOTFOUND = "ResourceNotFound.ServiceNotFound"
//  RESOURCENOTFOUND_SERVICERUNNINGVERSIONNOTFOUND = "ResourceNotFound.ServiceRunningVersionNotFound"
//  RESOURCENOTFOUND_VERSIONNAMESPACENOTFOUND = "ResourceNotFound.VersionNamespaceNotFound"
func (c *Client) DescribeApplicationPodsWithContext(ctx context.Context, request *DescribeApplicationPodsRequest) (response *DescribeApplicationPodsResponse, err error) {
    if request == nil {
        request = NewDescribeApplicationPodsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeApplicationPods require credential")
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
//  RESOURCENOTFOUND_VERSIONSERVICENOTFOUND = "ResourceNotFound.VersionServiceNotFound"
func (c *Client) DescribeDeployApplicationDetail(request *DescribeDeployApplicationDetailRequest) (response *DescribeDeployApplicationDetailResponse, err error) {
    return c.DescribeDeployApplicationDetailWithContext(context.Background(), request)
}

// DescribeDeployApplicationDetail
// 获取分批发布详情
//
// 可能返回的错误码:
//  INTERNALERROR_DESCRIBERUNPODLISTERROR = "InternalError.DescribeRunPodListError"
//  MISSINGPARAMETER_NAMESPACEIDNULL = "MissingParameter.NamespaceIdNull"
//  RESOURCENOTFOUND_SERVICENOTFOUND = "ResourceNotFound.ServiceNotFound"
//  RESOURCENOTFOUND_VERSIONNAMESPACENOTFOUND = "ResourceNotFound.VersionNamespaceNotFound"
//  RESOURCENOTFOUND_VERSIONSERVICENOTFOUND = "ResourceNotFound.VersionServiceNotFound"
func (c *Client) DescribeDeployApplicationDetailWithContext(ctx context.Context, request *DescribeDeployApplicationDetailRequest) (response *DescribeDeployApplicationDetailResponse, err error) {
    if request == nil {
        request = NewDescribeDeployApplicationDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDeployApplicationDetail require credential")
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
    return c.DescribeEnvironmentsWithContext(context.Background(), request)
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeEnvironments require credential")
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
//  RESOURCENOTFOUND_VERSIONNAMESPACENOTFOUND = "ResourceNotFound.VersionNamespaceNotFound"
func (c *Client) DescribeIngress(request *DescribeIngressRequest) (response *DescribeIngressResponse, err error) {
    return c.DescribeIngressWithContext(context.Background(), request)
}

// DescribeIngress
// 查询 Ingress 规则
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_VERSIONNAMESPACENOTFOUND = "ResourceNotFound.VersionNamespaceNotFound"
func (c *Client) DescribeIngressWithContext(ctx context.Context, request *DescribeIngressRequest) (response *DescribeIngressResponse, err error) {
    if request == nil {
        request = NewDescribeIngressRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeIngress require credential")
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
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  RESOURCENOTFOUND_VERSIONNAMESPACENOTFOUND = "ResourceNotFound.VersionNamespaceNotFound"
func (c *Client) DescribeIngresses(request *DescribeIngressesRequest) (response *DescribeIngressesResponse, err error) {
    return c.DescribeIngressesWithContext(context.Background(), request)
}

// DescribeIngresses
// 查询 Ingress 规则列表
//
// 可能返回的错误码:
//  MISSINGPARAMETER_NAMESPACEIDNULL = "MissingParameter.NamespaceIdNull"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  RESOURCENOTFOUND_VERSIONNAMESPACENOTFOUND = "ResourceNotFound.VersionNamespaceNotFound"
func (c *Client) DescribeIngressesWithContext(ctx context.Context, request *DescribeIngressesRequest) (response *DescribeIngressesResponse, err error) {
    if request == nil {
        request = NewDescribeIngressesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeIngresses require credential")
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
//  INTERNALERROR_DESCRIBESERVICEINGRESSERROR = "InternalError.DescribeServiceIngressError"
//  INVALIDPARAMETERVALUE_NAMESPACENOTBELONGTOAPPID = "InvalidParameterValue.NamespaceNotBelongToAppid"
//  MISSINGPARAMETER_NAMESPACEIDNULL = "MissingParameter.NamespaceIdNull"
//  MISSINGPARAMETER_SERVICEIDNULL = "MissingParameter.ServiceIdNull"
//  RESOURCENOTFOUND_SERVICENOTFOUND = "ResourceNotFound.ServiceNotFound"
//  RESOURCENOTFOUND_VERSIONNAMESPACENOTFOUND = "ResourceNotFound.VersionNamespaceNotFound"
func (c *Client) DescribeRelatedIngresses(request *DescribeRelatedIngressesRequest) (response *DescribeRelatedIngressesResponse, err error) {
    return c.DescribeRelatedIngressesWithContext(context.Background(), request)
}

// DescribeRelatedIngresses
// 查询应用关联的 Ingress 规则列表
//
// 可能返回的错误码:
//  INTERNALERROR_DESCRIBESERVICEINGRESSERROR = "InternalError.DescribeServiceIngressError"
//  INVALIDPARAMETERVALUE_NAMESPACENOTBELONGTOAPPID = "InvalidParameterValue.NamespaceNotBelongToAppid"
//  MISSINGPARAMETER_NAMESPACEIDNULL = "MissingParameter.NamespaceIdNull"
//  MISSINGPARAMETER_SERVICEIDNULL = "MissingParameter.ServiceIdNull"
//  RESOURCENOTFOUND_SERVICENOTFOUND = "ResourceNotFound.ServiceNotFound"
//  RESOURCENOTFOUND_VERSIONNAMESPACENOTFOUND = "ResourceNotFound.VersionNamespaceNotFound"
func (c *Client) DescribeRelatedIngressesWithContext(ctx context.Context, request *DescribeRelatedIngressesRequest) (response *DescribeRelatedIngressesResponse, err error) {
    if request == nil {
        request = NewDescribeRelatedIngressesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRelatedIngresses require credential")
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
//  INTERNALERROR_DESCRIBESERVICEINGRESSERROR = "InternalError.DescribeServiceIngressError"
//  INVALIDPARAMETERVALUE_NAMESPACENOTBELONGTOAPPID = "InvalidParameterValue.NamespaceNotBelongToAppid"
//  MISSINGPARAMETER_NAMESPACEIDNULL = "MissingParameter.NamespaceIdNull"
//  MISSINGPARAMETER_SERVICEIDNULL = "MissingParameter.ServiceIdNull"
//  RESOURCENOTFOUND_SERVICENOTFOUND = "ResourceNotFound.ServiceNotFound"
//  RESOURCENOTFOUND_VERSIONNAMESPACENOTFOUND = "ResourceNotFound.VersionNamespaceNotFound"
func (c *Client) GenerateApplicationPackageDownloadUrl(request *GenerateApplicationPackageDownloadUrlRequest) (response *GenerateApplicationPackageDownloadUrlResponse, err error) {
    return c.GenerateApplicationPackageDownloadUrlWithContext(context.Background(), request)
}

// GenerateApplicationPackageDownloadUrl
// 生成应用程序包预签名下载链接
//
// 可能返回的错误码:
//  INTERNALERROR_DESCRIBESERVICEINGRESSERROR = "InternalError.DescribeServiceIngressError"
//  INVALIDPARAMETERVALUE_NAMESPACENOTBELONGTOAPPID = "InvalidParameterValue.NamespaceNotBelongToAppid"
//  MISSINGPARAMETER_NAMESPACEIDNULL = "MissingParameter.NamespaceIdNull"
//  MISSINGPARAMETER_SERVICEIDNULL = "MissingParameter.ServiceIdNull"
//  RESOURCENOTFOUND_SERVICENOTFOUND = "ResourceNotFound.ServiceNotFound"
//  RESOURCENOTFOUND_VERSIONNAMESPACENOTFOUND = "ResourceNotFound.VersionNamespaceNotFound"
func (c *Client) GenerateApplicationPackageDownloadUrlWithContext(ctx context.Context, request *GenerateApplicationPackageDownloadUrlRequest) (response *GenerateApplicationPackageDownloadUrlResponse, err error) {
    if request == nil {
        request = NewGenerateApplicationPackageDownloadUrlRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GenerateApplicationPackageDownloadUrl require credential")
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
    return c.ModifyApplicationInfoWithContext(context.Background(), request)
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyApplicationInfo require credential")
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
    return c.ModifyApplicationReplicasWithContext(context.Background(), request)
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyApplicationReplicas require credential")
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
    return c.ModifyEnvironmentWithContext(context.Background(), request)
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyEnvironment require credential")
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
    return c.ModifyIngressWithContext(context.Background(), request)
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyIngress require credential")
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
//  RESOURCENOTFOUND_SERVICERUNNINGVERSIONNOTFOUND = "ResourceNotFound.ServiceRunningVersionNotFound"
//  RESOURCENOTFOUND_VERSIONNAMESPACENOTFOUND = "ResourceNotFound.VersionNamespaceNotFound"
func (c *Client) RestartApplication(request *RestartApplicationRequest) (response *RestartApplicationResponse, err error) {
    return c.RestartApplicationWithContext(context.Background(), request)
}

// RestartApplication
// 服务重启
//
// 可能返回的错误码:
//  INTERNALERROR_RESTARTAPPLICATIONERROR = "InternalError.RestartApplicationError"
//  MISSINGPARAMETER_NAMESPACEIDNULL = "MissingParameter.NamespaceIdNull"
//  RESOURCENOTFOUND_SERVICERUNNINGVERSIONNOTFOUND = "ResourceNotFound.ServiceRunningVersionNotFound"
//  RESOURCENOTFOUND_VERSIONNAMESPACENOTFOUND = "ResourceNotFound.VersionNamespaceNotFound"
func (c *Client) RestartApplicationWithContext(ctx context.Context, request *RestartApplicationRequest) (response *RestartApplicationResponse, err error) {
    if request == nil {
        request = NewRestartApplicationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RestartApplication require credential")
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
//  RESOURCENOTFOUND_SERVICERUNNINGVERSIONNOTFOUND = "ResourceNotFound.ServiceRunningVersionNotFound"
//  RESOURCENOTFOUND_VERSIONNAMESPACENOTFOUND = "ResourceNotFound.VersionNamespaceNotFound"
func (c *Client) RestartApplicationPod(request *RestartApplicationPodRequest) (response *RestartApplicationPodResponse, err error) {
    return c.RestartApplicationPodWithContext(context.Background(), request)
}

// RestartApplicationPod
// 重启应用实例
//
// 可能返回的错误码:
//  INTERNALERROR_RESTARTAPPLICATIONERROR = "InternalError.RestartApplicationError"
//  MISSINGPARAMETER_NAMESPACEIDNULL = "MissingParameter.NamespaceIdNull"
//  RESOURCENOTFOUND_SERVICERUNNINGVERSIONNOTFOUND = "ResourceNotFound.ServiceRunningVersionNotFound"
//  RESOURCENOTFOUND_VERSIONNAMESPACENOTFOUND = "ResourceNotFound.VersionNamespaceNotFound"
func (c *Client) RestartApplicationPodWithContext(ctx context.Context, request *RestartApplicationPodRequest) (response *RestartApplicationPodResponse, err error) {
    if request == nil {
        request = NewRestartApplicationPodRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RestartApplicationPod require credential")
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
//  MISSINGPARAMETER_NAMESPACEIDNULL = "MissingParameter.NamespaceIdNull"
func (c *Client) ResumeDeployApplication(request *ResumeDeployApplicationRequest) (response *ResumeDeployApplicationResponse, err error) {
    return c.ResumeDeployApplicationWithContext(context.Background(), request)
}

// ResumeDeployApplication
// 开始下一批次发布
//
// 可能返回的错误码:
//  MISSINGPARAMETER_NAMESPACEIDNULL = "MissingParameter.NamespaceIdNull"
func (c *Client) ResumeDeployApplicationWithContext(ctx context.Context, request *ResumeDeployApplicationRequest) (response *ResumeDeployApplicationResponse, err error) {
    if request == nil {
        request = NewResumeDeployApplicationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ResumeDeployApplication require credential")
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
//  RESOURCENOTFOUND_VERSIONNAMESPACENOTFOUND = "ResourceNotFound.VersionNamespaceNotFound"
func (c *Client) RevertDeployApplication(request *RevertDeployApplicationRequest) (response *RevertDeployApplicationResponse, err error) {
    return c.RevertDeployApplicationWithContext(context.Background(), request)
}

// RevertDeployApplication
// 回滚分批发布
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_VERSIONNAMESPACENOTFOUND = "ResourceNotFound.VersionNamespaceNotFound"
func (c *Client) RevertDeployApplicationWithContext(ctx context.Context, request *RevertDeployApplicationRequest) (response *RevertDeployApplicationResponse, err error) {
    if request == nil {
        request = NewRevertDeployApplicationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RevertDeployApplication require credential")
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
//  INVALIDPARAMETERVALUE_NAMESPACENOTBELONGTOAPPID = "InvalidParameterValue.NamespaceNotBelongToAppid"
//  INVALIDPARAMETERVALUE_VERSIONLOWERCASE = "InvalidParameterValue.VersionLowerCase"
//  RESOURCENOTFOUND_SERVICENOTFOUND = "ResourceNotFound.ServiceNotFound"
//  RESOURCENOTFOUND_VERSIONNAMESPACENOTFOUND = "ResourceNotFound.VersionNamespaceNotFound"
//  RESOURCEUNAVAILABLE_WAITFORKRUISE = "ResourceUnavailable.WaitForKruise"
func (c *Client) RollingUpdateApplicationByVersion(request *RollingUpdateApplicationByVersionRequest) (response *RollingUpdateApplicationByVersionResponse, err error) {
    return c.RollingUpdateApplicationByVersionWithContext(context.Background(), request)
}

// RollingUpdateApplicationByVersion
// 更新应用部署版本
//
// 可能返回的错误码:
//  INTERNALERROR_DEFAULTINTERNALERROR = "InternalError.DefaultInternalError"
//  INTERNALERROR_DEPLOYVERSIONERROR = "InternalError.DeployVersionError"
//  INVALIDPARAMETERVALUE_INVALIDDEPLOYVERSION = "InvalidParameterValue.InvalidDeployVersion"
//  INVALIDPARAMETERVALUE_NAMESPACENOTBELONGTOAPPID = "InvalidParameterValue.NamespaceNotBelongToAppid"
//  INVALIDPARAMETERVALUE_VERSIONLOWERCASE = "InvalidParameterValue.VersionLowerCase"
//  RESOURCENOTFOUND_SERVICENOTFOUND = "ResourceNotFound.ServiceNotFound"
//  RESOURCENOTFOUND_VERSIONNAMESPACENOTFOUND = "ResourceNotFound.VersionNamespaceNotFound"
//  RESOURCEUNAVAILABLE_WAITFORKRUISE = "ResourceUnavailable.WaitForKruise"
func (c *Client) RollingUpdateApplicationByVersionWithContext(ctx context.Context, request *RollingUpdateApplicationByVersionRequest) (response *RollingUpdateApplicationByVersionResponse, err error) {
    if request == nil {
        request = NewRollingUpdateApplicationByVersionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RollingUpdateApplicationByVersion require credential")
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
    return c.StopApplicationWithContext(context.Background(), request)
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("StopApplication require credential")
    }

    request.SetContext(ctx)
    
    response = NewStopApplicationResponse()
    err = c.Send(request, response)
    return
}
