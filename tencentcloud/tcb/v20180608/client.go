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

package v20180608

import (
    "context"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2018-06-08"

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


func NewBindEnvGatewayRequest() (request *BindEnvGatewayRequest) {
    request = &BindEnvGatewayRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcb", APIVersion, "BindEnvGateway")
    
    
    return
}

func NewBindEnvGatewayResponse() (response *BindEnvGatewayResponse) {
    response = &BindEnvGatewayResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// BindEnvGateway
// 绑定另外一个环境下的网关，callContainer请求可以访问到该网关
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) BindEnvGateway(request *BindEnvGatewayRequest) (response *BindEnvGatewayResponse, err error) {
    if request == nil {
        request = NewBindEnvGatewayRequest()
    }
    
    response = NewBindEnvGatewayResponse()
    err = c.Send(request, response)
    return
}

// BindEnvGateway
// 绑定另外一个环境下的网关，callContainer请求可以访问到该网关
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) BindEnvGatewayWithContext(ctx context.Context, request *BindEnvGatewayRequest) (response *BindEnvGatewayResponse, err error) {
    if request == nil {
        request = NewBindEnvGatewayRequest()
    }
    request.SetContext(ctx)
    
    response = NewBindEnvGatewayResponse()
    err = c.Send(request, response)
    return
}

func NewCheckTcbServiceRequest() (request *CheckTcbServiceRequest) {
    request = &CheckTcbServiceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcb", APIVersion, "CheckTcbService")
    
    
    return
}

func NewCheckTcbServiceResponse() (response *CheckTcbServiceResponse) {
    response = &CheckTcbServiceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CheckTcbService
// 检查是否开通Tcb服务
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_PARAM = "MissingParameter.Param"
func (c *Client) CheckTcbService(request *CheckTcbServiceRequest) (response *CheckTcbServiceResponse, err error) {
    if request == nil {
        request = NewCheckTcbServiceRequest()
    }
    
    response = NewCheckTcbServiceResponse()
    err = c.Send(request, response)
    return
}

// CheckTcbService
// 检查是否开通Tcb服务
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_PARAM = "MissingParameter.Param"
func (c *Client) CheckTcbServiceWithContext(ctx context.Context, request *CheckTcbServiceRequest) (response *CheckTcbServiceResponse, err error) {
    if request == nil {
        request = NewCheckTcbServiceRequest()
    }
    request.SetContext(ctx)
    
    response = NewCheckTcbServiceResponse()
    err = c.Send(request, response)
    return
}

func NewCommonServiceAPIRequest() (request *CommonServiceAPIRequest) {
    request = &CommonServiceAPIRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcb", APIVersion, "CommonServiceAPI")
    
    
    return
}

func NewCommonServiceAPIResponse() (response *CommonServiceAPIResponse) {
    response = &CommonServiceAPIResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CommonServiceAPI
// TCB云API统一入口
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_RESOURCENOTEXISTS = "InvalidParameter.ResourceNotExists"
func (c *Client) CommonServiceAPI(request *CommonServiceAPIRequest) (response *CommonServiceAPIResponse, err error) {
    if request == nil {
        request = NewCommonServiceAPIRequest()
    }
    
    response = NewCommonServiceAPIResponse()
    err = c.Send(request, response)
    return
}

// CommonServiceAPI
// TCB云API统一入口
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_RESOURCENOTEXISTS = "InvalidParameter.ResourceNotExists"
func (c *Client) CommonServiceAPIWithContext(ctx context.Context, request *CommonServiceAPIRequest) (response *CommonServiceAPIResponse, err error) {
    if request == nil {
        request = NewCommonServiceAPIRequest()
    }
    request.SetContext(ctx)
    
    response = NewCommonServiceAPIResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAndDeployCloudBaseProjectRequest() (request *CreateAndDeployCloudBaseProjectRequest) {
    request = &CreateAndDeployCloudBaseProjectRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcb", APIVersion, "CreateAndDeployCloudBaseProject")
    
    
    return
}

func NewCreateAndDeployCloudBaseProjectResponse() (response *CreateAndDeployCloudBaseProjectResponse) {
    response = &CreateAndDeployCloudBaseProjectResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateAndDeployCloudBaseProject
// 创建云开发项目
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER_PARAM = "MissingParameter.Param"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CODEOAUTHUNAUTHORIZED = "UnauthorizedOperation.CodeOAuthUnauthorized"
func (c *Client) CreateAndDeployCloudBaseProject(request *CreateAndDeployCloudBaseProjectRequest) (response *CreateAndDeployCloudBaseProjectResponse, err error) {
    if request == nil {
        request = NewCreateAndDeployCloudBaseProjectRequest()
    }
    
    response = NewCreateAndDeployCloudBaseProjectResponse()
    err = c.Send(request, response)
    return
}

// CreateAndDeployCloudBaseProject
// 创建云开发项目
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER_PARAM = "MissingParameter.Param"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CODEOAUTHUNAUTHORIZED = "UnauthorizedOperation.CodeOAuthUnauthorized"
func (c *Client) CreateAndDeployCloudBaseProjectWithContext(ctx context.Context, request *CreateAndDeployCloudBaseProjectRequest) (response *CreateAndDeployCloudBaseProjectResponse, err error) {
    if request == nil {
        request = NewCreateAndDeployCloudBaseProjectRequest()
    }
    request.SetContext(ctx)
    
    response = NewCreateAndDeployCloudBaseProjectResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAuthDomainRequest() (request *CreateAuthDomainRequest) {
    request = &CreateAuthDomainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcb", APIVersion, "CreateAuthDomain")
    
    
    return
}

func NewCreateAuthDomainResponse() (response *CreateAuthDomainResponse) {
    response = &CreateAuthDomainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateAuthDomain
// 增加安全域名
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER_PARAM = "MissingParameter.Param"
//  OPERATIONDENIED_RESOURCEFROZEN = "OperationDenied.ResourceFrozen"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) CreateAuthDomain(request *CreateAuthDomainRequest) (response *CreateAuthDomainResponse, err error) {
    if request == nil {
        request = NewCreateAuthDomainRequest()
    }
    
    response = NewCreateAuthDomainResponse()
    err = c.Send(request, response)
    return
}

// CreateAuthDomain
// 增加安全域名
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER_PARAM = "MissingParameter.Param"
//  OPERATIONDENIED_RESOURCEFROZEN = "OperationDenied.ResourceFrozen"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) CreateAuthDomainWithContext(ctx context.Context, request *CreateAuthDomainRequest) (response *CreateAuthDomainResponse, err error) {
    if request == nil {
        request = NewCreateAuthDomainRequest()
    }
    request.SetContext(ctx)
    
    response = NewCreateAuthDomainResponse()
    err = c.Send(request, response)
    return
}

func NewCreateCloudBaseRunResourceRequest() (request *CreateCloudBaseRunResourceRequest) {
    request = &CreateCloudBaseRunResourceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcb", APIVersion, "CreateCloudBaseRunResource")
    
    
    return
}

func NewCreateCloudBaseRunResourceResponse() (response *CreateCloudBaseRunResourceResponse) {
    response = &CreateCloudBaseRunResourceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateCloudBaseRunResource
// 开通容器托管的资源，包括集群创建，VPC配置，异步任务创建，镜像托管，Coding等，查看创建结果需要根据DescribeCloudBaseRunResource接口来查看
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED_ERRNAMESPACEMAXLIMIT = "LimitExceeded.ErrNamespaceMaxLimit"
//  LIMITEXCEEDED_ERRREPOMAXLIMIT = "LimitExceeded.ErrRepoMaxLimit"
func (c *Client) CreateCloudBaseRunResource(request *CreateCloudBaseRunResourceRequest) (response *CreateCloudBaseRunResourceResponse, err error) {
    if request == nil {
        request = NewCreateCloudBaseRunResourceRequest()
    }
    
    response = NewCreateCloudBaseRunResourceResponse()
    err = c.Send(request, response)
    return
}

// CreateCloudBaseRunResource
// 开通容器托管的资源，包括集群创建，VPC配置，异步任务创建，镜像托管，Coding等，查看创建结果需要根据DescribeCloudBaseRunResource接口来查看
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED_ERRNAMESPACEMAXLIMIT = "LimitExceeded.ErrNamespaceMaxLimit"
//  LIMITEXCEEDED_ERRREPOMAXLIMIT = "LimitExceeded.ErrRepoMaxLimit"
func (c *Client) CreateCloudBaseRunResourceWithContext(ctx context.Context, request *CreateCloudBaseRunResourceRequest) (response *CreateCloudBaseRunResourceResponse, err error) {
    if request == nil {
        request = NewCreateCloudBaseRunResourceRequest()
    }
    request.SetContext(ctx)
    
    response = NewCreateCloudBaseRunResourceResponse()
    err = c.Send(request, response)
    return
}

func NewCreateCloudBaseRunServerVersionRequest() (request *CreateCloudBaseRunServerVersionRequest) {
    request = &CreateCloudBaseRunServerVersionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcb", APIVersion, "CreateCloudBaseRunServerVersion")
    
    
    return
}

func NewCreateCloudBaseRunServerVersionResponse() (response *CreateCloudBaseRunServerVersionResponse) {
    response = &CreateCloudBaseRunServerVersionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateCloudBaseRunServerVersion
// 创建服务版本
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SERVICENOTEXIST = "InvalidParameter.ServiceNotExist"
//  LIMITEXCEEDED_ERRNAMESPACEMAXLIMIT = "LimitExceeded.ErrNamespaceMaxLimit"
//  LIMITEXCEEDED_ERRREPOMAXLIMIT = "LimitExceeded.ErrRepoMaxLimit"
func (c *Client) CreateCloudBaseRunServerVersion(request *CreateCloudBaseRunServerVersionRequest) (response *CreateCloudBaseRunServerVersionResponse, err error) {
    if request == nil {
        request = NewCreateCloudBaseRunServerVersionRequest()
    }
    
    response = NewCreateCloudBaseRunServerVersionResponse()
    err = c.Send(request, response)
    return
}

// CreateCloudBaseRunServerVersion
// 创建服务版本
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SERVICENOTEXIST = "InvalidParameter.ServiceNotExist"
//  LIMITEXCEEDED_ERRNAMESPACEMAXLIMIT = "LimitExceeded.ErrNamespaceMaxLimit"
//  LIMITEXCEEDED_ERRREPOMAXLIMIT = "LimitExceeded.ErrRepoMaxLimit"
func (c *Client) CreateCloudBaseRunServerVersionWithContext(ctx context.Context, request *CreateCloudBaseRunServerVersionRequest) (response *CreateCloudBaseRunServerVersionResponse, err error) {
    if request == nil {
        request = NewCreateCloudBaseRunServerVersionRequest()
    }
    request.SetContext(ctx)
    
    response = NewCreateCloudBaseRunServerVersionResponse()
    err = c.Send(request, response)
    return
}

func NewCreateHostingDomainRequest() (request *CreateHostingDomainRequest) {
    request = &CreateHostingDomainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcb", APIVersion, "CreateHostingDomain")
    
    
    return
}

func NewCreateHostingDomainResponse() (response *CreateHostingDomainResponse) {
    response = &CreateHostingDomainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateHostingDomain
// 创建托管域名
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_PARAM = "MissingParameter.Param"
//  OPERATIONDENIED_RESOURCEFROZEN = "OperationDenied.ResourceFrozen"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateHostingDomain(request *CreateHostingDomainRequest) (response *CreateHostingDomainResponse, err error) {
    if request == nil {
        request = NewCreateHostingDomainRequest()
    }
    
    response = NewCreateHostingDomainResponse()
    err = c.Send(request, response)
    return
}

// CreateHostingDomain
// 创建托管域名
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_PARAM = "MissingParameter.Param"
//  OPERATIONDENIED_RESOURCEFROZEN = "OperationDenied.ResourceFrozen"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateHostingDomainWithContext(ctx context.Context, request *CreateHostingDomainRequest) (response *CreateHostingDomainResponse, err error) {
    if request == nil {
        request = NewCreateHostingDomainRequest()
    }
    request.SetContext(ctx)
    
    response = NewCreateHostingDomainResponse()
    err = c.Send(request, response)
    return
}

func NewCreatePostpayPackageRequest() (request *CreatePostpayPackageRequest) {
    request = &CreatePostpayPackageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcb", APIVersion, "CreatePostpayPackage")
    
    
    return
}

func NewCreatePostpayPackageResponse() (response *CreatePostpayPackageResponse) {
    response = &CreatePostpayPackageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreatePostpayPackage
// 开通后付费资源
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE_BALANCENOTENOUGH = "ResourceUnavailable.BalanceNotEnough"
func (c *Client) CreatePostpayPackage(request *CreatePostpayPackageRequest) (response *CreatePostpayPackageResponse, err error) {
    if request == nil {
        request = NewCreatePostpayPackageRequest()
    }
    
    response = NewCreatePostpayPackageResponse()
    err = c.Send(request, response)
    return
}

// CreatePostpayPackage
// 开通后付费资源
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE_BALANCENOTENOUGH = "ResourceUnavailable.BalanceNotEnough"
func (c *Client) CreatePostpayPackageWithContext(ctx context.Context, request *CreatePostpayPackageRequest) (response *CreatePostpayPackageResponse, err error) {
    if request == nil {
        request = NewCreatePostpayPackageRequest()
    }
    request.SetContext(ctx)
    
    response = NewCreatePostpayPackageResponse()
    err = c.Send(request, response)
    return
}

func NewCreateStandaloneGatewayRequest() (request *CreateStandaloneGatewayRequest) {
    request = &CreateStandaloneGatewayRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcb", APIVersion, "CreateStandaloneGateway")
    
    
    return
}

func NewCreateStandaloneGatewayResponse() (response *CreateStandaloneGatewayResponse) {
    response = &CreateStandaloneGatewayResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateStandaloneGateway
// 本接口（CreateStandaloneGateway）用于创建独立网关。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_TIMEOUT = "InternalError.Timeout"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ACTION = "InvalidParameter.Action"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_PARAM = "MissingParameter.Param"
func (c *Client) CreateStandaloneGateway(request *CreateStandaloneGatewayRequest) (response *CreateStandaloneGatewayResponse, err error) {
    if request == nil {
        request = NewCreateStandaloneGatewayRequest()
    }
    
    response = NewCreateStandaloneGatewayResponse()
    err = c.Send(request, response)
    return
}

// CreateStandaloneGateway
// 本接口（CreateStandaloneGateway）用于创建独立网关。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_TIMEOUT = "InternalError.Timeout"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ACTION = "InvalidParameter.Action"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_PARAM = "MissingParameter.Param"
func (c *Client) CreateStandaloneGatewayWithContext(ctx context.Context, request *CreateStandaloneGatewayRequest) (response *CreateStandaloneGatewayResponse, err error) {
    if request == nil {
        request = NewCreateStandaloneGatewayRequest()
    }
    request.SetContext(ctx)
    
    response = NewCreateStandaloneGatewayResponse()
    err = c.Send(request, response)
    return
}

func NewCreateStaticStoreRequest() (request *CreateStaticStoreRequest) {
    request = &CreateStaticStoreRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcb", APIVersion, "CreateStaticStore")
    
    
    return
}

func NewCreateStaticStoreResponse() (response *CreateStaticStoreResponse) {
    response = &CreateStaticStoreResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateStaticStore
// 创建静态托管资源，包括COS和CDN，异步任务创建，查看创建结果需要根据DescribeStaticStore接口来查看
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateStaticStore(request *CreateStaticStoreRequest) (response *CreateStaticStoreResponse, err error) {
    if request == nil {
        request = NewCreateStaticStoreRequest()
    }
    
    response = NewCreateStaticStoreResponse()
    err = c.Send(request, response)
    return
}

// CreateStaticStore
// 创建静态托管资源，包括COS和CDN，异步任务创建，查看创建结果需要根据DescribeStaticStore接口来查看
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateStaticStoreWithContext(ctx context.Context, request *CreateStaticStoreRequest) (response *CreateStaticStoreResponse, err error) {
    if request == nil {
        request = NewCreateStaticStoreRequest()
    }
    request.SetContext(ctx)
    
    response = NewCreateStaticStoreResponse()
    err = c.Send(request, response)
    return
}

func NewCreateWxCloudBaseRunEnvRequest() (request *CreateWxCloudBaseRunEnvRequest) {
    request = &CreateWxCloudBaseRunEnvRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcb", APIVersion, "CreateWxCloudBaseRunEnv")
    
    
    return
}

func NewCreateWxCloudBaseRunEnvResponse() (response *CreateWxCloudBaseRunEnvResponse) {
    response = &CreateWxCloudBaseRunEnvResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateWxCloudBaseRunEnv
// 创建微信云托管
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SYSTEMFAIL = "InternalError.SystemFail"
//  INTERNALERROR_TIMEOUT = "InternalError.Timeout"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ACTION = "InvalidParameter.Action"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_PARAM = "MissingParameter.Param"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) CreateWxCloudBaseRunEnv(request *CreateWxCloudBaseRunEnvRequest) (response *CreateWxCloudBaseRunEnvResponse, err error) {
    if request == nil {
        request = NewCreateWxCloudBaseRunEnvRequest()
    }
    
    response = NewCreateWxCloudBaseRunEnvResponse()
    err = c.Send(request, response)
    return
}

// CreateWxCloudBaseRunEnv
// 创建微信云托管
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SYSTEMFAIL = "InternalError.SystemFail"
//  INTERNALERROR_TIMEOUT = "InternalError.Timeout"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ACTION = "InvalidParameter.Action"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_PARAM = "MissingParameter.Param"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) CreateWxCloudBaseRunEnvWithContext(ctx context.Context, request *CreateWxCloudBaseRunEnvRequest) (response *CreateWxCloudBaseRunEnvResponse, err error) {
    if request == nil {
        request = NewCreateWxCloudBaseRunEnvRequest()
    }
    request.SetContext(ctx)
    
    response = NewCreateWxCloudBaseRunEnvResponse()
    err = c.Send(request, response)
    return
}

func NewCreateWxCloudBaseRunServerDBClusterRequest() (request *CreateWxCloudBaseRunServerDBClusterRequest) {
    request = &CreateWxCloudBaseRunServerDBClusterRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcb", APIVersion, "CreateWxCloudBaseRunServerDBCluster")
    
    
    return
}

func NewCreateWxCloudBaseRunServerDBClusterResponse() (response *CreateWxCloudBaseRunServerDBClusterResponse) {
    response = &CreateWxCloudBaseRunServerDBClusterResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateWxCloudBaseRunServerDBCluster
// 开通微信云托管MySQL数据库服务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DATABASE = "InternalError.Database"
//  INTERNALERROR_SYSTEMFAIL = "InternalError.SystemFail"
//  INTERNALERROR_TIMEOUT = "InternalError.Timeout"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ENVID = "InvalidParameter.EnvId"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_REQUEST = "LimitExceeded.Request"
//  UNSUPPORTEDOPERATION_TASKEXISTED = "UnsupportedOperation.TaskExisted"
func (c *Client) CreateWxCloudBaseRunServerDBCluster(request *CreateWxCloudBaseRunServerDBClusterRequest) (response *CreateWxCloudBaseRunServerDBClusterResponse, err error) {
    if request == nil {
        request = NewCreateWxCloudBaseRunServerDBClusterRequest()
    }
    
    response = NewCreateWxCloudBaseRunServerDBClusterResponse()
    err = c.Send(request, response)
    return
}

// CreateWxCloudBaseRunServerDBCluster
// 开通微信云托管MySQL数据库服务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DATABASE = "InternalError.Database"
//  INTERNALERROR_SYSTEMFAIL = "InternalError.SystemFail"
//  INTERNALERROR_TIMEOUT = "InternalError.Timeout"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ENVID = "InvalidParameter.EnvId"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_REQUEST = "LimitExceeded.Request"
//  UNSUPPORTEDOPERATION_TASKEXISTED = "UnsupportedOperation.TaskExisted"
func (c *Client) CreateWxCloudBaseRunServerDBClusterWithContext(ctx context.Context, request *CreateWxCloudBaseRunServerDBClusterRequest) (response *CreateWxCloudBaseRunServerDBClusterResponse, err error) {
    if request == nil {
        request = NewCreateWxCloudBaseRunServerDBClusterRequest()
    }
    request.SetContext(ctx)
    
    response = NewCreateWxCloudBaseRunServerDBClusterResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteCloudBaseProjectLatestVersionRequest() (request *DeleteCloudBaseProjectLatestVersionRequest) {
    request = &DeleteCloudBaseProjectLatestVersionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcb", APIVersion, "DeleteCloudBaseProjectLatestVersion")
    
    
    return
}

func NewDeleteCloudBaseProjectLatestVersionResponse() (response *DeleteCloudBaseProjectLatestVersionResponse) {
    response = &DeleteCloudBaseProjectLatestVersionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteCloudBaseProjectLatestVersion
// 删除云项目
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER_PARAM = "MissingParameter.Param"
func (c *Client) DeleteCloudBaseProjectLatestVersion(request *DeleteCloudBaseProjectLatestVersionRequest) (response *DeleteCloudBaseProjectLatestVersionResponse, err error) {
    if request == nil {
        request = NewDeleteCloudBaseProjectLatestVersionRequest()
    }
    
    response = NewDeleteCloudBaseProjectLatestVersionResponse()
    err = c.Send(request, response)
    return
}

// DeleteCloudBaseProjectLatestVersion
// 删除云项目
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER_PARAM = "MissingParameter.Param"
func (c *Client) DeleteCloudBaseProjectLatestVersionWithContext(ctx context.Context, request *DeleteCloudBaseProjectLatestVersionRequest) (response *DeleteCloudBaseProjectLatestVersionResponse, err error) {
    if request == nil {
        request = NewDeleteCloudBaseProjectLatestVersionRequest()
    }
    request.SetContext(ctx)
    
    response = NewDeleteCloudBaseProjectLatestVersionResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteCloudBaseRunServerVersionRequest() (request *DeleteCloudBaseRunServerVersionRequest) {
    request = &DeleteCloudBaseRunServerVersionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcb", APIVersion, "DeleteCloudBaseRunServerVersion")
    
    
    return
}

func NewDeleteCloudBaseRunServerVersionResponse() (response *DeleteCloudBaseRunServerVersionResponse) {
    response = &DeleteCloudBaseRunServerVersionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteCloudBaseRunServerVersion
// 删除服务版本
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SERVICENOTEXIST = "InvalidParameter.ServiceNotExist"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEUNAVAILABLE_CDNFREEZED = "ResourceUnavailable.CDNFreezed"
func (c *Client) DeleteCloudBaseRunServerVersion(request *DeleteCloudBaseRunServerVersionRequest) (response *DeleteCloudBaseRunServerVersionResponse, err error) {
    if request == nil {
        request = NewDeleteCloudBaseRunServerVersionRequest()
    }
    
    response = NewDeleteCloudBaseRunServerVersionResponse()
    err = c.Send(request, response)
    return
}

// DeleteCloudBaseRunServerVersion
// 删除服务版本
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SERVICENOTEXIST = "InvalidParameter.ServiceNotExist"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEUNAVAILABLE_CDNFREEZED = "ResourceUnavailable.CDNFreezed"
func (c *Client) DeleteCloudBaseRunServerVersionWithContext(ctx context.Context, request *DeleteCloudBaseRunServerVersionRequest) (response *DeleteCloudBaseRunServerVersionResponse, err error) {
    if request == nil {
        request = NewDeleteCloudBaseRunServerVersionRequest()
    }
    request.SetContext(ctx)
    
    response = NewDeleteCloudBaseRunServerVersionResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteEndUserRequest() (request *DeleteEndUserRequest) {
    request = &DeleteEndUserRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcb", APIVersion, "DeleteEndUser")
    
    
    return
}

func NewDeleteEndUserResponse() (response *DeleteEndUserResponse) {
    response = &DeleteEndUserResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteEndUser
// 删除终端用户
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DeleteEndUser(request *DeleteEndUserRequest) (response *DeleteEndUserResponse, err error) {
    if request == nil {
        request = NewDeleteEndUserRequest()
    }
    
    response = NewDeleteEndUserResponse()
    err = c.Send(request, response)
    return
}

// DeleteEndUser
// 删除终端用户
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DeleteEndUserWithContext(ctx context.Context, request *DeleteEndUserRequest) (response *DeleteEndUserResponse, err error) {
    if request == nil {
        request = NewDeleteEndUserRequest()
    }
    request.SetContext(ctx)
    
    response = NewDeleteEndUserResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteWxGatewayRouteRequest() (request *DeleteWxGatewayRouteRequest) {
    request = &DeleteWxGatewayRouteRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcb", APIVersion, "DeleteWxGatewayRoute")
    
    
    return
}

func NewDeleteWxGatewayRouteResponse() (response *DeleteWxGatewayRouteResponse) {
    response = &DeleteWxGatewayRouteResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteWxGatewayRoute
// 删除安全网关路由
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED_ERRNAMESPACEMAXLIMIT = "LimitExceeded.ErrNamespaceMaxLimit"
//  LIMITEXCEEDED_ERRREPOMAXLIMIT = "LimitExceeded.ErrRepoMaxLimit"
//  RESOURCEINUSE = "ResourceInUse"
func (c *Client) DeleteWxGatewayRoute(request *DeleteWxGatewayRouteRequest) (response *DeleteWxGatewayRouteResponse, err error) {
    if request == nil {
        request = NewDeleteWxGatewayRouteRequest()
    }
    
    response = NewDeleteWxGatewayRouteResponse()
    err = c.Send(request, response)
    return
}

// DeleteWxGatewayRoute
// 删除安全网关路由
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED_ERRNAMESPACEMAXLIMIT = "LimitExceeded.ErrNamespaceMaxLimit"
//  LIMITEXCEEDED_ERRREPOMAXLIMIT = "LimitExceeded.ErrRepoMaxLimit"
//  RESOURCEINUSE = "ResourceInUse"
func (c *Client) DeleteWxGatewayRouteWithContext(ctx context.Context, request *DeleteWxGatewayRouteRequest) (response *DeleteWxGatewayRouteResponse, err error) {
    if request == nil {
        request = NewDeleteWxGatewayRouteRequest()
    }
    request.SetContext(ctx)
    
    response = NewDeleteWxGatewayRouteResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeActivityInfoRequest() (request *DescribeActivityInfoRequest) {
    request = &DescribeActivityInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcb", APIVersion, "DescribeActivityInfo")
    
    
    return
}

func NewDescribeActivityInfoResponse() (response *DescribeActivityInfoResponse) {
    response = &DescribeActivityInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeActivityInfo
// 查询活动信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_PARAM = "MissingParameter.Param"
func (c *Client) DescribeActivityInfo(request *DescribeActivityInfoRequest) (response *DescribeActivityInfoResponse, err error) {
    if request == nil {
        request = NewDescribeActivityInfoRequest()
    }
    
    response = NewDescribeActivityInfoResponse()
    err = c.Send(request, response)
    return
}

// DescribeActivityInfo
// 查询活动信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_PARAM = "MissingParameter.Param"
func (c *Client) DescribeActivityInfoWithContext(ctx context.Context, request *DescribeActivityInfoRequest) (response *DescribeActivityInfoResponse, err error) {
    if request == nil {
        request = NewDescribeActivityInfoRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeActivityInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeActivityRecordRequest() (request *DescribeActivityRecordRequest) {
    request = &DescribeActivityRecordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcb", APIVersion, "DescribeActivityRecord")
    
    
    return
}

func NewDescribeActivityRecordResponse() (response *DescribeActivityRecordResponse) {
    response = &DescribeActivityRecordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeActivityRecord
// 查询活动记录信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_PARAM = "MissingParameter.Param"
func (c *Client) DescribeActivityRecord(request *DescribeActivityRecordRequest) (response *DescribeActivityRecordResponse, err error) {
    if request == nil {
        request = NewDescribeActivityRecordRequest()
    }
    
    response = NewDescribeActivityRecordResponse()
    err = c.Send(request, response)
    return
}

// DescribeActivityRecord
// 查询活动记录信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_PARAM = "MissingParameter.Param"
func (c *Client) DescribeActivityRecordWithContext(ctx context.Context, request *DescribeActivityRecordRequest) (response *DescribeActivityRecordResponse, err error) {
    if request == nil {
        request = NewDescribeActivityRecordRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeActivityRecordResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAuthDomainsRequest() (request *DescribeAuthDomainsRequest) {
    request = &DescribeAuthDomainsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcb", APIVersion, "DescribeAuthDomains")
    
    
    return
}

func NewDescribeAuthDomainsResponse() (response *DescribeAuthDomainsResponse) {
    response = &DescribeAuthDomainsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAuthDomains
// 获取安全域名列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER_PARAM = "MissingParameter.Param"
func (c *Client) DescribeAuthDomains(request *DescribeAuthDomainsRequest) (response *DescribeAuthDomainsResponse, err error) {
    if request == nil {
        request = NewDescribeAuthDomainsRequest()
    }
    
    response = NewDescribeAuthDomainsResponse()
    err = c.Send(request, response)
    return
}

// DescribeAuthDomains
// 获取安全域名列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER_PARAM = "MissingParameter.Param"
func (c *Client) DescribeAuthDomainsWithContext(ctx context.Context, request *DescribeAuthDomainsRequest) (response *DescribeAuthDomainsResponse, err error) {
    if request == nil {
        request = NewDescribeAuthDomainsRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeAuthDomainsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCloudBaseBuildServiceRequest() (request *DescribeCloudBaseBuildServiceRequest) {
    request = &DescribeCloudBaseBuildServiceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcb", APIVersion, "DescribeCloudBaseBuildService")
    
    
    return
}

func NewDescribeCloudBaseBuildServiceResponse() (response *DescribeCloudBaseBuildServiceResponse) {
    response = &DescribeCloudBaseBuildServiceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCloudBaseBuildService
// 获取云托管代码上传url
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeCloudBaseBuildService(request *DescribeCloudBaseBuildServiceRequest) (response *DescribeCloudBaseBuildServiceResponse, err error) {
    if request == nil {
        request = NewDescribeCloudBaseBuildServiceRequest()
    }
    
    response = NewDescribeCloudBaseBuildServiceResponse()
    err = c.Send(request, response)
    return
}

// DescribeCloudBaseBuildService
// 获取云托管代码上传url
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeCloudBaseBuildServiceWithContext(ctx context.Context, request *DescribeCloudBaseBuildServiceRequest) (response *DescribeCloudBaseBuildServiceResponse, err error) {
    if request == nil {
        request = NewDescribeCloudBaseBuildServiceRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeCloudBaseBuildServiceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCloudBaseProjectLatestVersionListRequest() (request *DescribeCloudBaseProjectLatestVersionListRequest) {
    request = &DescribeCloudBaseProjectLatestVersionListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcb", APIVersion, "DescribeCloudBaseProjectLatestVersionList")
    
    
    return
}

func NewDescribeCloudBaseProjectLatestVersionListResponse() (response *DescribeCloudBaseProjectLatestVersionListResponse) {
    response = &DescribeCloudBaseProjectLatestVersionListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCloudBaseProjectLatestVersionList
// 获取云开发项目列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER_PARAM = "MissingParameter.Param"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeCloudBaseProjectLatestVersionList(request *DescribeCloudBaseProjectLatestVersionListRequest) (response *DescribeCloudBaseProjectLatestVersionListResponse, err error) {
    if request == nil {
        request = NewDescribeCloudBaseProjectLatestVersionListRequest()
    }
    
    response = NewDescribeCloudBaseProjectLatestVersionListResponse()
    err = c.Send(request, response)
    return
}

// DescribeCloudBaseProjectLatestVersionList
// 获取云开发项目列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER_PARAM = "MissingParameter.Param"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeCloudBaseProjectLatestVersionListWithContext(ctx context.Context, request *DescribeCloudBaseProjectLatestVersionListRequest) (response *DescribeCloudBaseProjectLatestVersionListResponse, err error) {
    if request == nil {
        request = NewDescribeCloudBaseProjectLatestVersionListRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeCloudBaseProjectLatestVersionListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCloudBaseProjectVersionListRequest() (request *DescribeCloudBaseProjectVersionListRequest) {
    request = &DescribeCloudBaseProjectVersionListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcb", APIVersion, "DescribeCloudBaseProjectVersionList")
    
    
    return
}

func NewDescribeCloudBaseProjectVersionListResponse() (response *DescribeCloudBaseProjectVersionListResponse) {
    response = &DescribeCloudBaseProjectVersionListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCloudBaseProjectVersionList
// 云项目部署列表
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeCloudBaseProjectVersionList(request *DescribeCloudBaseProjectVersionListRequest) (response *DescribeCloudBaseProjectVersionListResponse, err error) {
    if request == nil {
        request = NewDescribeCloudBaseProjectVersionListRequest()
    }
    
    response = NewDescribeCloudBaseProjectVersionListResponse()
    err = c.Send(request, response)
    return
}

// DescribeCloudBaseProjectVersionList
// 云项目部署列表
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeCloudBaseProjectVersionListWithContext(ctx context.Context, request *DescribeCloudBaseProjectVersionListRequest) (response *DescribeCloudBaseProjectVersionListResponse, err error) {
    if request == nil {
        request = NewDescribeCloudBaseProjectVersionListRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeCloudBaseProjectVersionListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCloudBaseRunAllVpcsRequest() (request *DescribeCloudBaseRunAllVpcsRequest) {
    request = &DescribeCloudBaseRunAllVpcsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcb", APIVersion, "DescribeCloudBaseRunAllVpcs")
    
    
    return
}

func NewDescribeCloudBaseRunAllVpcsResponse() (response *DescribeCloudBaseRunAllVpcsResponse) {
    response = &DescribeCloudBaseRunAllVpcsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCloudBaseRunAllVpcs
// 查询环境下所有的vpc列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeCloudBaseRunAllVpcs(request *DescribeCloudBaseRunAllVpcsRequest) (response *DescribeCloudBaseRunAllVpcsResponse, err error) {
    if request == nil {
        request = NewDescribeCloudBaseRunAllVpcsRequest()
    }
    
    response = NewDescribeCloudBaseRunAllVpcsResponse()
    err = c.Send(request, response)
    return
}

// DescribeCloudBaseRunAllVpcs
// 查询环境下所有的vpc列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeCloudBaseRunAllVpcsWithContext(ctx context.Context, request *DescribeCloudBaseRunAllVpcsRequest) (response *DescribeCloudBaseRunAllVpcsResponse, err error) {
    if request == nil {
        request = NewDescribeCloudBaseRunAllVpcsRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeCloudBaseRunAllVpcsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCloudBaseRunConfForGateWayRequest() (request *DescribeCloudBaseRunConfForGateWayRequest) {
    request = &DescribeCloudBaseRunConfForGateWayRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcb", APIVersion, "DescribeCloudBaseRunConfForGateWay")
    
    
    return
}

func NewDescribeCloudBaseRunConfForGateWayResponse() (response *DescribeCloudBaseRunConfForGateWayResponse) {
    response = &DescribeCloudBaseRunConfForGateWayResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCloudBaseRunConfForGateWay
// 独立网关中拉取云托管服务对应的配置信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_TIMEOUT = "InternalError.Timeout"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeCloudBaseRunConfForGateWay(request *DescribeCloudBaseRunConfForGateWayRequest) (response *DescribeCloudBaseRunConfForGateWayResponse, err error) {
    if request == nil {
        request = NewDescribeCloudBaseRunConfForGateWayRequest()
    }
    
    response = NewDescribeCloudBaseRunConfForGateWayResponse()
    err = c.Send(request, response)
    return
}

// DescribeCloudBaseRunConfForGateWay
// 独立网关中拉取云托管服务对应的配置信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_TIMEOUT = "InternalError.Timeout"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeCloudBaseRunConfForGateWayWithContext(ctx context.Context, request *DescribeCloudBaseRunConfForGateWayRequest) (response *DescribeCloudBaseRunConfForGateWayResponse, err error) {
    if request == nil {
        request = NewDescribeCloudBaseRunConfForGateWayRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeCloudBaseRunConfForGateWayResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCloudBaseRunOneClickTaskExternalRequest() (request *DescribeCloudBaseRunOneClickTaskExternalRequest) {
    request = &DescribeCloudBaseRunOneClickTaskExternalRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcb", APIVersion, "DescribeCloudBaseRunOneClickTaskExternal")
    
    
    return
}

func NewDescribeCloudBaseRunOneClickTaskExternalResponse() (response *DescribeCloudBaseRunOneClickTaskExternalResponse) {
    response = &DescribeCloudBaseRunOneClickTaskExternalResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCloudBaseRunOneClickTaskExternal
// 查询一键部署任务 （特定接口：外部查询使用）
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeCloudBaseRunOneClickTaskExternal(request *DescribeCloudBaseRunOneClickTaskExternalRequest) (response *DescribeCloudBaseRunOneClickTaskExternalResponse, err error) {
    if request == nil {
        request = NewDescribeCloudBaseRunOneClickTaskExternalRequest()
    }
    
    response = NewDescribeCloudBaseRunOneClickTaskExternalResponse()
    err = c.Send(request, response)
    return
}

// DescribeCloudBaseRunOneClickTaskExternal
// 查询一键部署任务 （特定接口：外部查询使用）
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeCloudBaseRunOneClickTaskExternalWithContext(ctx context.Context, request *DescribeCloudBaseRunOneClickTaskExternalRequest) (response *DescribeCloudBaseRunOneClickTaskExternalResponse, err error) {
    if request == nil {
        request = NewDescribeCloudBaseRunOneClickTaskExternalRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeCloudBaseRunOneClickTaskExternalResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCloudBaseRunOperationTypesRequest() (request *DescribeCloudBaseRunOperationTypesRequest) {
    request = &DescribeCloudBaseRunOperationTypesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcb", APIVersion, "DescribeCloudBaseRunOperationTypes")
    
    
    return
}

func NewDescribeCloudBaseRunOperationTypesResponse() (response *DescribeCloudBaseRunOperationTypesResponse) {
    response = &DescribeCloudBaseRunOperationTypesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCloudBaseRunOperationTypes
// 查询服务、版本和操作类型
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeCloudBaseRunOperationTypes(request *DescribeCloudBaseRunOperationTypesRequest) (response *DescribeCloudBaseRunOperationTypesResponse, err error) {
    if request == nil {
        request = NewDescribeCloudBaseRunOperationTypesRequest()
    }
    
    response = NewDescribeCloudBaseRunOperationTypesResponse()
    err = c.Send(request, response)
    return
}

// DescribeCloudBaseRunOperationTypes
// 查询服务、版本和操作类型
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeCloudBaseRunOperationTypesWithContext(ctx context.Context, request *DescribeCloudBaseRunOperationTypesRequest) (response *DescribeCloudBaseRunOperationTypesResponse, err error) {
    if request == nil {
        request = NewDescribeCloudBaseRunOperationTypesRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeCloudBaseRunOperationTypesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCloudBaseRunPodListRequest() (request *DescribeCloudBaseRunPodListRequest) {
    request = &DescribeCloudBaseRunPodListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcb", APIVersion, "DescribeCloudBaseRunPodList")
    
    
    return
}

func NewDescribeCloudBaseRunPodListResponse() (response *DescribeCloudBaseRunPodListResponse) {
    response = &DescribeCloudBaseRunPodListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCloudBaseRunPodList
// 查询云应用服务版本容器列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SERVICENOTEXIST = "InvalidParameter.ServiceNotExist"
//  MISSINGPARAMETER = "MissingParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeCloudBaseRunPodList(request *DescribeCloudBaseRunPodListRequest) (response *DescribeCloudBaseRunPodListResponse, err error) {
    if request == nil {
        request = NewDescribeCloudBaseRunPodListRequest()
    }
    
    response = NewDescribeCloudBaseRunPodListResponse()
    err = c.Send(request, response)
    return
}

// DescribeCloudBaseRunPodList
// 查询云应用服务版本容器列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SERVICENOTEXIST = "InvalidParameter.ServiceNotExist"
//  MISSINGPARAMETER = "MissingParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeCloudBaseRunPodListWithContext(ctx context.Context, request *DescribeCloudBaseRunPodListRequest) (response *DescribeCloudBaseRunPodListResponse, err error) {
    if request == nil {
        request = NewDescribeCloudBaseRunPodListRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeCloudBaseRunPodListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCloudBaseRunResourceRequest() (request *DescribeCloudBaseRunResourceRequest) {
    request = &DescribeCloudBaseRunResourceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcb", APIVersion, "DescribeCloudBaseRunResource")
    
    
    return
}

func NewDescribeCloudBaseRunResourceResponse() (response *DescribeCloudBaseRunResourceResponse) {
    response = &DescribeCloudBaseRunResourceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCloudBaseRunResource
// 查看容器托管的集群状态
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeCloudBaseRunResource(request *DescribeCloudBaseRunResourceRequest) (response *DescribeCloudBaseRunResourceResponse, err error) {
    if request == nil {
        request = NewDescribeCloudBaseRunResourceRequest()
    }
    
    response = NewDescribeCloudBaseRunResourceResponse()
    err = c.Send(request, response)
    return
}

// DescribeCloudBaseRunResource
// 查看容器托管的集群状态
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeCloudBaseRunResourceWithContext(ctx context.Context, request *DescribeCloudBaseRunResourceRequest) (response *DescribeCloudBaseRunResourceResponse, err error) {
    if request == nil {
        request = NewDescribeCloudBaseRunResourceRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeCloudBaseRunResourceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCloudBaseRunResourceForExtendRequest() (request *DescribeCloudBaseRunResourceForExtendRequest) {
    request = &DescribeCloudBaseRunResourceForExtendRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcb", APIVersion, "DescribeCloudBaseRunResourceForExtend")
    
    
    return
}

func NewDescribeCloudBaseRunResourceForExtendResponse() (response *DescribeCloudBaseRunResourceForExtendResponse) {
    response = &DescribeCloudBaseRunResourceForExtendResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCloudBaseRunResourceForExtend
// 查看容器托管的集群状态扩展使用
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeCloudBaseRunResourceForExtend(request *DescribeCloudBaseRunResourceForExtendRequest) (response *DescribeCloudBaseRunResourceForExtendResponse, err error) {
    if request == nil {
        request = NewDescribeCloudBaseRunResourceForExtendRequest()
    }
    
    response = NewDescribeCloudBaseRunResourceForExtendResponse()
    err = c.Send(request, response)
    return
}

// DescribeCloudBaseRunResourceForExtend
// 查看容器托管的集群状态扩展使用
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeCloudBaseRunResourceForExtendWithContext(ctx context.Context, request *DescribeCloudBaseRunResourceForExtendRequest) (response *DescribeCloudBaseRunResourceForExtendResponse, err error) {
    if request == nil {
        request = NewDescribeCloudBaseRunResourceForExtendRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeCloudBaseRunResourceForExtendResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCloudBaseRunServerRequest() (request *DescribeCloudBaseRunServerRequest) {
    request = &DescribeCloudBaseRunServerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcb", APIVersion, "DescribeCloudBaseRunServer")
    
    
    return
}

func NewDescribeCloudBaseRunServerResponse() (response *DescribeCloudBaseRunServerResponse) {
    response = &DescribeCloudBaseRunServerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCloudBaseRunServer
// 查询单个服务的详情，版本以及详情
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeCloudBaseRunServer(request *DescribeCloudBaseRunServerRequest) (response *DescribeCloudBaseRunServerResponse, err error) {
    if request == nil {
        request = NewDescribeCloudBaseRunServerRequest()
    }
    
    response = NewDescribeCloudBaseRunServerResponse()
    err = c.Send(request, response)
    return
}

// DescribeCloudBaseRunServer
// 查询单个服务的详情，版本以及详情
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeCloudBaseRunServerWithContext(ctx context.Context, request *DescribeCloudBaseRunServerRequest) (response *DescribeCloudBaseRunServerResponse, err error) {
    if request == nil {
        request = NewDescribeCloudBaseRunServerRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeCloudBaseRunServerResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCloudBaseRunServerDomainNameRequest() (request *DescribeCloudBaseRunServerDomainNameRequest) {
    request = &DescribeCloudBaseRunServerDomainNameRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcb", APIVersion, "DescribeCloudBaseRunServerDomainName")
    
    
    return
}

func NewDescribeCloudBaseRunServerDomainNameResponse() (response *DescribeCloudBaseRunServerDomainNameResponse) {
    response = &DescribeCloudBaseRunServerDomainNameResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCloudBaseRunServerDomainName
// 查询微信云托管服务域名
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
func (c *Client) DescribeCloudBaseRunServerDomainName(request *DescribeCloudBaseRunServerDomainNameRequest) (response *DescribeCloudBaseRunServerDomainNameResponse, err error) {
    if request == nil {
        request = NewDescribeCloudBaseRunServerDomainNameRequest()
    }
    
    response = NewDescribeCloudBaseRunServerDomainNameResponse()
    err = c.Send(request, response)
    return
}

// DescribeCloudBaseRunServerDomainName
// 查询微信云托管服务域名
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
func (c *Client) DescribeCloudBaseRunServerDomainNameWithContext(ctx context.Context, request *DescribeCloudBaseRunServerDomainNameRequest) (response *DescribeCloudBaseRunServerDomainNameResponse, err error) {
    if request == nil {
        request = NewDescribeCloudBaseRunServerDomainNameRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeCloudBaseRunServerDomainNameResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCloudBaseRunServerVersionRequest() (request *DescribeCloudBaseRunServerVersionRequest) {
    request = &DescribeCloudBaseRunServerVersionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcb", APIVersion, "DescribeCloudBaseRunServerVersion")
    
    
    return
}

func NewDescribeCloudBaseRunServerVersionResponse() (response *DescribeCloudBaseRunServerVersionResponse) {
    response = &DescribeCloudBaseRunServerVersionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCloudBaseRunServerVersion
// 查询服务版本的详情，CPU和MEM  请使用CPUSize和MemSize
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SERVICENOTEXIST = "InvalidParameter.ServiceNotExist"
//  LIMITEXCEEDED_ERRNAMESPACEMAXLIMIT = "LimitExceeded.ErrNamespaceMaxLimit"
//  LIMITEXCEEDED_ERRREPOMAXLIMIT = "LimitExceeded.ErrRepoMaxLimit"
func (c *Client) DescribeCloudBaseRunServerVersion(request *DescribeCloudBaseRunServerVersionRequest) (response *DescribeCloudBaseRunServerVersionResponse, err error) {
    if request == nil {
        request = NewDescribeCloudBaseRunServerVersionRequest()
    }
    
    response = NewDescribeCloudBaseRunServerVersionResponse()
    err = c.Send(request, response)
    return
}

// DescribeCloudBaseRunServerVersion
// 查询服务版本的详情，CPU和MEM  请使用CPUSize和MemSize
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SERVICENOTEXIST = "InvalidParameter.ServiceNotExist"
//  LIMITEXCEEDED_ERRNAMESPACEMAXLIMIT = "LimitExceeded.ErrNamespaceMaxLimit"
//  LIMITEXCEEDED_ERRREPOMAXLIMIT = "LimitExceeded.ErrRepoMaxLimit"
func (c *Client) DescribeCloudBaseRunServerVersionWithContext(ctx context.Context, request *DescribeCloudBaseRunServerVersionRequest) (response *DescribeCloudBaseRunServerVersionResponse, err error) {
    if request == nil {
        request = NewDescribeCloudBaseRunServerVersionRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeCloudBaseRunServerVersionResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCloudBaseRunVersionRequest() (request *DescribeCloudBaseRunVersionRequest) {
    request = &DescribeCloudBaseRunVersionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcb", APIVersion, "DescribeCloudBaseRunVersion")
    
    
    return
}

func NewDescribeCloudBaseRunVersionResponse() (response *DescribeCloudBaseRunVersionResponse) {
    response = &DescribeCloudBaseRunVersionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCloudBaseRunVersion
// 查询服务版本详情(新)
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SERVICENOTEXIST = "InvalidParameter.ServiceNotExist"
//  LIMITEXCEEDED_ERRNAMESPACEMAXLIMIT = "LimitExceeded.ErrNamespaceMaxLimit"
//  LIMITEXCEEDED_ERRREPOMAXLIMIT = "LimitExceeded.ErrRepoMaxLimit"
func (c *Client) DescribeCloudBaseRunVersion(request *DescribeCloudBaseRunVersionRequest) (response *DescribeCloudBaseRunVersionResponse, err error) {
    if request == nil {
        request = NewDescribeCloudBaseRunVersionRequest()
    }
    
    response = NewDescribeCloudBaseRunVersionResponse()
    err = c.Send(request, response)
    return
}

// DescribeCloudBaseRunVersion
// 查询服务版本详情(新)
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SERVICENOTEXIST = "InvalidParameter.ServiceNotExist"
//  LIMITEXCEEDED_ERRNAMESPACEMAXLIMIT = "LimitExceeded.ErrNamespaceMaxLimit"
//  LIMITEXCEEDED_ERRREPOMAXLIMIT = "LimitExceeded.ErrRepoMaxLimit"
func (c *Client) DescribeCloudBaseRunVersionWithContext(ctx context.Context, request *DescribeCloudBaseRunVersionRequest) (response *DescribeCloudBaseRunVersionResponse, err error) {
    if request == nil {
        request = NewDescribeCloudBaseRunVersionRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeCloudBaseRunVersionResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCloudBaseRunVersionRsByConditionRequest() (request *DescribeCloudBaseRunVersionRsByConditionRequest) {
    request = &DescribeCloudBaseRunVersionRsByConditionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcb", APIVersion, "DescribeCloudBaseRunVersionRsByCondition")
    
    
    return
}

func NewDescribeCloudBaseRunVersionRsByConditionResponse() (response *DescribeCloudBaseRunVersionRsByConditionResponse) {
    response = &DescribeCloudBaseRunVersionRsByConditionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCloudBaseRunVersionRsByCondition
// DescribeCloudBaseRunVersionRsByCondition 获取云托管详情
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) DescribeCloudBaseRunVersionRsByCondition(request *DescribeCloudBaseRunVersionRsByConditionRequest) (response *DescribeCloudBaseRunVersionRsByConditionResponse, err error) {
    if request == nil {
        request = NewDescribeCloudBaseRunVersionRsByConditionRequest()
    }
    
    response = NewDescribeCloudBaseRunVersionRsByConditionResponse()
    err = c.Send(request, response)
    return
}

// DescribeCloudBaseRunVersionRsByCondition
// DescribeCloudBaseRunVersionRsByCondition 获取云托管详情
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) DescribeCloudBaseRunVersionRsByConditionWithContext(ctx context.Context, request *DescribeCloudBaseRunVersionRsByConditionRequest) (response *DescribeCloudBaseRunVersionRsByConditionResponse, err error) {
    if request == nil {
        request = NewDescribeCloudBaseRunVersionRsByConditionRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeCloudBaseRunVersionRsByConditionResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCloudBaseRunVersionSnapshotRequest() (request *DescribeCloudBaseRunVersionSnapshotRequest) {
    request = &DescribeCloudBaseRunVersionSnapshotRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcb", APIVersion, "DescribeCloudBaseRunVersionSnapshot")
    
    
    return
}

func NewDescribeCloudBaseRunVersionSnapshotResponse() (response *DescribeCloudBaseRunVersionSnapshotResponse) {
    response = &DescribeCloudBaseRunVersionSnapshotResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCloudBaseRunVersionSnapshot
// 查询版本历史
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DATABASE = "InternalError.Database"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SERVICENOTEXIST = "InvalidParameter.ServiceNotExist"
func (c *Client) DescribeCloudBaseRunVersionSnapshot(request *DescribeCloudBaseRunVersionSnapshotRequest) (response *DescribeCloudBaseRunVersionSnapshotResponse, err error) {
    if request == nil {
        request = NewDescribeCloudBaseRunVersionSnapshotRequest()
    }
    
    response = NewDescribeCloudBaseRunVersionSnapshotResponse()
    err = c.Send(request, response)
    return
}

// DescribeCloudBaseRunVersionSnapshot
// 查询版本历史
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DATABASE = "InternalError.Database"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SERVICENOTEXIST = "InvalidParameter.ServiceNotExist"
func (c *Client) DescribeCloudBaseRunVersionSnapshotWithContext(ctx context.Context, request *DescribeCloudBaseRunVersionSnapshotRequest) (response *DescribeCloudBaseRunVersionSnapshotResponse, err error) {
    if request == nil {
        request = NewDescribeCloudBaseRunVersionSnapshotRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeCloudBaseRunVersionSnapshotResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCurveDataRequest() (request *DescribeCurveDataRequest) {
    request = &DescribeCurveDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcb", APIVersion, "DescribeCurveData")
    
    
    return
}

func NewDescribeCurveDataResponse() (response *DescribeCurveDataResponse) {
    response = &DescribeCurveDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCurveData
// 根据用户传入的指标, 拉取一段时间内的监控数据。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_PARAM = "MissingParameter.Param"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeCurveData(request *DescribeCurveDataRequest) (response *DescribeCurveDataResponse, err error) {
    if request == nil {
        request = NewDescribeCurveDataRequest()
    }
    
    response = NewDescribeCurveDataResponse()
    err = c.Send(request, response)
    return
}

// DescribeCurveData
// 根据用户传入的指标, 拉取一段时间内的监控数据。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_PARAM = "MissingParameter.Param"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeCurveDataWithContext(ctx context.Context, request *DescribeCurveDataRequest) (response *DescribeCurveDataResponse, err error) {
    if request == nil {
        request = NewDescribeCurveDataRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeCurveDataResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDatabaseACLRequest() (request *DescribeDatabaseACLRequest) {
    request = &DescribeDatabaseACLRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcb", APIVersion, "DescribeDatabaseACL")
    
    
    return
}

func NewDescribeDatabaseACLResponse() (response *DescribeDatabaseACLResponse) {
    response = &DescribeDatabaseACLResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDatabaseACL
// 获取数据库权限
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_PARAM = "MissingParameter.Param"
func (c *Client) DescribeDatabaseACL(request *DescribeDatabaseACLRequest) (response *DescribeDatabaseACLResponse, err error) {
    if request == nil {
        request = NewDescribeDatabaseACLRequest()
    }
    
    response = NewDescribeDatabaseACLResponse()
    err = c.Send(request, response)
    return
}

// DescribeDatabaseACL
// 获取数据库权限
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_PARAM = "MissingParameter.Param"
func (c *Client) DescribeDatabaseACLWithContext(ctx context.Context, request *DescribeDatabaseACLRequest) (response *DescribeDatabaseACLResponse, err error) {
    if request == nil {
        request = NewDescribeDatabaseACLRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeDatabaseACLResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDownloadFileRequest() (request *DescribeDownloadFileRequest) {
    request = &DescribeDownloadFileRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcb", APIVersion, "DescribeDownloadFile")
    
    
    return
}

func NewDescribeDownloadFileResponse() (response *DescribeDownloadFileResponse) {
    response = &DescribeDownloadFileResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDownloadFile
// 获取下载文件信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeDownloadFile(request *DescribeDownloadFileRequest) (response *DescribeDownloadFileResponse, err error) {
    if request == nil {
        request = NewDescribeDownloadFileRequest()
    }
    
    response = NewDescribeDownloadFileResponse()
    err = c.Send(request, response)
    return
}

// DescribeDownloadFile
// 获取下载文件信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeDownloadFileWithContext(ctx context.Context, request *DescribeDownloadFileRequest) (response *DescribeDownloadFileResponse, err error) {
    if request == nil {
        request = NewDescribeDownloadFileRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeDownloadFileResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEndUserLoginStatisticRequest() (request *DescribeEndUserLoginStatisticRequest) {
    request = &DescribeEndUserLoginStatisticRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcb", APIVersion, "DescribeEndUserLoginStatistic")
    
    
    return
}

func NewDescribeEndUserLoginStatisticResponse() (response *DescribeEndUserLoginStatisticResponse) {
    response = &DescribeEndUserLoginStatisticResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeEndUserLoginStatistic
// 获取环境终端用户新增与登录信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeEndUserLoginStatistic(request *DescribeEndUserLoginStatisticRequest) (response *DescribeEndUserLoginStatisticResponse, err error) {
    if request == nil {
        request = NewDescribeEndUserLoginStatisticRequest()
    }
    
    response = NewDescribeEndUserLoginStatisticResponse()
    err = c.Send(request, response)
    return
}

// DescribeEndUserLoginStatistic
// 获取环境终端用户新增与登录信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeEndUserLoginStatisticWithContext(ctx context.Context, request *DescribeEndUserLoginStatisticRequest) (response *DescribeEndUserLoginStatisticResponse, err error) {
    if request == nil {
        request = NewDescribeEndUserLoginStatisticRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeEndUserLoginStatisticResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEndUserStatisticRequest() (request *DescribeEndUserStatisticRequest) {
    request = &DescribeEndUserStatisticRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcb", APIVersion, "DescribeEndUserStatistic")
    
    
    return
}

func NewDescribeEndUserStatisticResponse() (response *DescribeEndUserStatisticResponse) {
    response = &DescribeEndUserStatisticResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeEndUserStatistic
// 获取终端用户总量与平台分布情况
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeEndUserStatistic(request *DescribeEndUserStatisticRequest) (response *DescribeEndUserStatisticResponse, err error) {
    if request == nil {
        request = NewDescribeEndUserStatisticRequest()
    }
    
    response = NewDescribeEndUserStatisticResponse()
    err = c.Send(request, response)
    return
}

// DescribeEndUserStatistic
// 获取终端用户总量与平台分布情况
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeEndUserStatisticWithContext(ctx context.Context, request *DescribeEndUserStatisticRequest) (response *DescribeEndUserStatisticResponse, err error) {
    if request == nil {
        request = NewDescribeEndUserStatisticRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeEndUserStatisticResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEndUsersRequest() (request *DescribeEndUsersRequest) {
    request = &DescribeEndUsersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcb", APIVersion, "DescribeEndUsers")
    
    
    return
}

func NewDescribeEndUsersResponse() (response *DescribeEndUsersResponse) {
    response = &DescribeEndUsersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeEndUsers
// 获取终端用户列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER_PARAM = "MissingParameter.Param"
func (c *Client) DescribeEndUsers(request *DescribeEndUsersRequest) (response *DescribeEndUsersResponse, err error) {
    if request == nil {
        request = NewDescribeEndUsersRequest()
    }
    
    response = NewDescribeEndUsersResponse()
    err = c.Send(request, response)
    return
}

// DescribeEndUsers
// 获取终端用户列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER_PARAM = "MissingParameter.Param"
func (c *Client) DescribeEndUsersWithContext(ctx context.Context, request *DescribeEndUsersRequest) (response *DescribeEndUsersResponse, err error) {
    if request == nil {
        request = NewDescribeEndUsersRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeEndUsersResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEnvDealRegionRequest() (request *DescribeEnvDealRegionRequest) {
    request = &DescribeEnvDealRegionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcb", APIVersion, "DescribeEnvDealRegion")
    
    
    return
}

func NewDescribeEnvDealRegionResponse() (response *DescribeEnvDealRegionResponse) {
    response = &DescribeEnvDealRegionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeEnvDealRegion
// 获取环境下单地域
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_CONCURRENT = "LimitExceeded.Concurrent"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_PARAM = "MissingParameter.Param"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeEnvDealRegion(request *DescribeEnvDealRegionRequest) (response *DescribeEnvDealRegionResponse, err error) {
    if request == nil {
        request = NewDescribeEnvDealRegionRequest()
    }
    
    response = NewDescribeEnvDealRegionResponse()
    err = c.Send(request, response)
    return
}

// DescribeEnvDealRegion
// 获取环境下单地域
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_CONCURRENT = "LimitExceeded.Concurrent"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_PARAM = "MissingParameter.Param"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeEnvDealRegionWithContext(ctx context.Context, request *DescribeEnvDealRegionRequest) (response *DescribeEnvDealRegionResponse, err error) {
    if request == nil {
        request = NewDescribeEnvDealRegionRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeEnvDealRegionResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEnvFreeQuotaRequest() (request *DescribeEnvFreeQuotaRequest) {
    request = &DescribeEnvFreeQuotaRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcb", APIVersion, "DescribeEnvFreeQuota")
    
    
    return
}

func NewDescribeEnvFreeQuotaResponse() (response *DescribeEnvFreeQuotaResponse) {
    response = &DescribeEnvFreeQuotaResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeEnvFreeQuota
// 查询后付费免费配额信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeEnvFreeQuota(request *DescribeEnvFreeQuotaRequest) (response *DescribeEnvFreeQuotaResponse, err error) {
    if request == nil {
        request = NewDescribeEnvFreeQuotaRequest()
    }
    
    response = NewDescribeEnvFreeQuotaResponse()
    err = c.Send(request, response)
    return
}

// DescribeEnvFreeQuota
// 查询后付费免费配额信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeEnvFreeQuotaWithContext(ctx context.Context, request *DescribeEnvFreeQuotaRequest) (response *DescribeEnvFreeQuotaResponse, err error) {
    if request == nil {
        request = NewDescribeEnvFreeQuotaRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeEnvFreeQuotaResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEnvLimitRequest() (request *DescribeEnvLimitRequest) {
    request = &DescribeEnvLimitRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcb", APIVersion, "DescribeEnvLimit")
    
    
    return
}

func NewDescribeEnvLimitResponse() (response *DescribeEnvLimitResponse) {
    response = &DescribeEnvLimitResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeEnvLimit
// 查询环境个数上限
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_PARAM = "MissingParameter.Param"
func (c *Client) DescribeEnvLimit(request *DescribeEnvLimitRequest) (response *DescribeEnvLimitResponse, err error) {
    if request == nil {
        request = NewDescribeEnvLimitRequest()
    }
    
    response = NewDescribeEnvLimitResponse()
    err = c.Send(request, response)
    return
}

// DescribeEnvLimit
// 查询环境个数上限
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_PARAM = "MissingParameter.Param"
func (c *Client) DescribeEnvLimitWithContext(ctx context.Context, request *DescribeEnvLimitRequest) (response *DescribeEnvLimitResponse, err error) {
    if request == nil {
        request = NewDescribeEnvLimitRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeEnvLimitResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEnvPostpaidDeductRequest() (request *DescribeEnvPostpaidDeductRequest) {
    request = &DescribeEnvPostpaidDeductRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcb", APIVersion, "DescribeEnvPostpaidDeduct")
    
    
    return
}

func NewDescribeEnvPostpaidDeductResponse() (response *DescribeEnvPostpaidDeductResponse) {
    response = &DescribeEnvPostpaidDeductResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeEnvPostpaidDeduct
// 查询环境后付费计费详情
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_PARAM = "MissingParameter.Param"
func (c *Client) DescribeEnvPostpaidDeduct(request *DescribeEnvPostpaidDeductRequest) (response *DescribeEnvPostpaidDeductResponse, err error) {
    if request == nil {
        request = NewDescribeEnvPostpaidDeductRequest()
    }
    
    response = NewDescribeEnvPostpaidDeductResponse()
    err = c.Send(request, response)
    return
}

// DescribeEnvPostpaidDeduct
// 查询环境后付费计费详情
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_PARAM = "MissingParameter.Param"
func (c *Client) DescribeEnvPostpaidDeductWithContext(ctx context.Context, request *DescribeEnvPostpaidDeductRequest) (response *DescribeEnvPostpaidDeductResponse, err error) {
    if request == nil {
        request = NewDescribeEnvPostpaidDeductRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeEnvPostpaidDeductResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEnvsRequest() (request *DescribeEnvsRequest) {
    request = &DescribeEnvsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcb", APIVersion, "DescribeEnvs")
    
    
    return
}

func NewDescribeEnvsResponse() (response *DescribeEnvsResponse) {
    response = &DescribeEnvsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeEnvs
// 获取环境列表，含环境下的各个资源信息。尤其是各资源的唯一标识，是请求各资源的关键参数
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ACTION = "InvalidParameter.Action"
//  INVALIDPARAMETER_ENVID = "InvalidParameter.EnvId"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_PARAM = "MissingParameter.Param"
//  RESOURCENOTFOUND_USERNOTEXISTS = "ResourceNotFound.UserNotExists"
func (c *Client) DescribeEnvs(request *DescribeEnvsRequest) (response *DescribeEnvsResponse, err error) {
    if request == nil {
        request = NewDescribeEnvsRequest()
    }
    
    response = NewDescribeEnvsResponse()
    err = c.Send(request, response)
    return
}

// DescribeEnvs
// 获取环境列表，含环境下的各个资源信息。尤其是各资源的唯一标识，是请求各资源的关键参数
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ACTION = "InvalidParameter.Action"
//  INVALIDPARAMETER_ENVID = "InvalidParameter.EnvId"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_PARAM = "MissingParameter.Param"
//  RESOURCENOTFOUND_USERNOTEXISTS = "ResourceNotFound.UserNotExists"
func (c *Client) DescribeEnvsWithContext(ctx context.Context, request *DescribeEnvsRequest) (response *DescribeEnvsResponse, err error) {
    if request == nil {
        request = NewDescribeEnvsRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeEnvsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeExtensionUploadInfoRequest() (request *DescribeExtensionUploadInfoRequest) {
    request = &DescribeExtensionUploadInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcb", APIVersion, "DescribeExtensionUploadInfo")
    
    
    return
}

func NewDescribeExtensionUploadInfoResponse() (response *DescribeExtensionUploadInfoResponse) {
    response = &DescribeExtensionUploadInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeExtensionUploadInfo
// 描述扩展上传文件信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeExtensionUploadInfo(request *DescribeExtensionUploadInfoRequest) (response *DescribeExtensionUploadInfoResponse, err error) {
    if request == nil {
        request = NewDescribeExtensionUploadInfoRequest()
    }
    
    response = NewDescribeExtensionUploadInfoResponse()
    err = c.Send(request, response)
    return
}

// DescribeExtensionUploadInfo
// 描述扩展上传文件信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeExtensionUploadInfoWithContext(ctx context.Context, request *DescribeExtensionUploadInfoRequest) (response *DescribeExtensionUploadInfoResponse, err error) {
    if request == nil {
        request = NewDescribeExtensionUploadInfoRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeExtensionUploadInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeExtraPkgBillingInfoRequest() (request *DescribeExtraPkgBillingInfoRequest) {
    request = &DescribeExtraPkgBillingInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcb", APIVersion, "DescribeExtraPkgBillingInfo")
    
    
    return
}

func NewDescribeExtraPkgBillingInfoResponse() (response *DescribeExtraPkgBillingInfoResponse) {
    response = &DescribeExtraPkgBillingInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeExtraPkgBillingInfo
// 获取增值包计费相关信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_PARAM = "MissingParameter.Param"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeExtraPkgBillingInfo(request *DescribeExtraPkgBillingInfoRequest) (response *DescribeExtraPkgBillingInfoResponse, err error) {
    if request == nil {
        request = NewDescribeExtraPkgBillingInfoRequest()
    }
    
    response = NewDescribeExtraPkgBillingInfoResponse()
    err = c.Send(request, response)
    return
}

// DescribeExtraPkgBillingInfo
// 获取增值包计费相关信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_PARAM = "MissingParameter.Param"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeExtraPkgBillingInfoWithContext(ctx context.Context, request *DescribeExtraPkgBillingInfoRequest) (response *DescribeExtraPkgBillingInfoResponse, err error) {
    if request == nil {
        request = NewDescribeExtraPkgBillingInfoRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeExtraPkgBillingInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeHostingDomainTaskRequest() (request *DescribeHostingDomainTaskRequest) {
    request = &DescribeHostingDomainTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcb", APIVersion, "DescribeHostingDomainTask")
    
    
    return
}

func NewDescribeHostingDomainTaskResponse() (response *DescribeHostingDomainTaskResponse) {
    response = &DescribeHostingDomainTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeHostingDomainTask
// 查询静态托管域名任务状态
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_CONCURRENT = "LimitExceeded.Concurrent"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeHostingDomainTask(request *DescribeHostingDomainTaskRequest) (response *DescribeHostingDomainTaskResponse, err error) {
    if request == nil {
        request = NewDescribeHostingDomainTaskRequest()
    }
    
    response = NewDescribeHostingDomainTaskResponse()
    err = c.Send(request, response)
    return
}

// DescribeHostingDomainTask
// 查询静态托管域名任务状态
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_CONCURRENT = "LimitExceeded.Concurrent"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeHostingDomainTaskWithContext(ctx context.Context, request *DescribeHostingDomainTaskRequest) (response *DescribeHostingDomainTaskResponse, err error) {
    if request == nil {
        request = NewDescribeHostingDomainTaskRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeHostingDomainTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePostpayFreeQuotasRequest() (request *DescribePostpayFreeQuotasRequest) {
    request = &DescribePostpayFreeQuotasRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcb", APIVersion, "DescribePostpayFreeQuotas")
    
    
    return
}

func NewDescribePostpayFreeQuotasResponse() (response *DescribePostpayFreeQuotasResponse) {
    response = &DescribePostpayFreeQuotasResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribePostpayFreeQuotas
// 查询后付费资源免费量
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_PARAM = "MissingParameter.Param"
func (c *Client) DescribePostpayFreeQuotas(request *DescribePostpayFreeQuotasRequest) (response *DescribePostpayFreeQuotasResponse, err error) {
    if request == nil {
        request = NewDescribePostpayFreeQuotasRequest()
    }
    
    response = NewDescribePostpayFreeQuotasResponse()
    err = c.Send(request, response)
    return
}

// DescribePostpayFreeQuotas
// 查询后付费资源免费量
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_PARAM = "MissingParameter.Param"
func (c *Client) DescribePostpayFreeQuotasWithContext(ctx context.Context, request *DescribePostpayFreeQuotasRequest) (response *DescribePostpayFreeQuotasResponse, err error) {
    if request == nil {
        request = NewDescribePostpayFreeQuotasRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribePostpayFreeQuotasResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePostpayPackageFreeQuotasRequest() (request *DescribePostpayPackageFreeQuotasRequest) {
    request = &DescribePostpayPackageFreeQuotasRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcb", APIVersion, "DescribePostpayPackageFreeQuotas")
    
    
    return
}

func NewDescribePostpayPackageFreeQuotasResponse() (response *DescribePostpayPackageFreeQuotasResponse) {
    response = &DescribePostpayPackageFreeQuotasResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribePostpayPackageFreeQuotas
// 获取后付费免费额度
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_PARAM = "MissingParameter.Param"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribePostpayPackageFreeQuotas(request *DescribePostpayPackageFreeQuotasRequest) (response *DescribePostpayPackageFreeQuotasResponse, err error) {
    if request == nil {
        request = NewDescribePostpayPackageFreeQuotasRequest()
    }
    
    response = NewDescribePostpayPackageFreeQuotasResponse()
    err = c.Send(request, response)
    return
}

// DescribePostpayPackageFreeQuotas
// 获取后付费免费额度
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_PARAM = "MissingParameter.Param"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribePostpayPackageFreeQuotasWithContext(ctx context.Context, request *DescribePostpayPackageFreeQuotasRequest) (response *DescribePostpayPackageFreeQuotasResponse, err error) {
    if request == nil {
        request = NewDescribePostpayPackageFreeQuotasRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribePostpayPackageFreeQuotasResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeQuotaDataRequest() (request *DescribeQuotaDataRequest) {
    request = &DescribeQuotaDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcb", APIVersion, "DescribeQuotaData")
    
    
    return
}

func NewDescribeQuotaDataResponse() (response *DescribeQuotaDataResponse) {
    response = &DescribeQuotaDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeQuotaData
// 查询指定指标的配额使用量
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_PARAM = "MissingParameter.Param"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeQuotaData(request *DescribeQuotaDataRequest) (response *DescribeQuotaDataResponse, err error) {
    if request == nil {
        request = NewDescribeQuotaDataRequest()
    }
    
    response = NewDescribeQuotaDataResponse()
    err = c.Send(request, response)
    return
}

// DescribeQuotaData
// 查询指定指标的配额使用量
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_PARAM = "MissingParameter.Param"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeQuotaDataWithContext(ctx context.Context, request *DescribeQuotaDataRequest) (response *DescribeQuotaDataResponse, err error) {
    if request == nil {
        request = NewDescribeQuotaDataRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeQuotaDataResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSmsQuotasRequest() (request *DescribeSmsQuotasRequest) {
    request = &DescribeSmsQuotasRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcb", APIVersion, "DescribeSmsQuotas")
    
    
    return
}

func NewDescribeSmsQuotasResponse() (response *DescribeSmsQuotasResponse) {
    response = &DescribeSmsQuotasResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeSmsQuotas
// 查询后付费短信资源量
//
// 1 有免费包的返回SmsFreeQuota结构所有字段
//
// 2 没有免费包，有付费包，付费返回复用SmsFreeQuota结构，其中只有 TodayUsedQuota 字段有效
//
// 3 都没有返回为空数组
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_PARAM = "MissingParameter.Param"
func (c *Client) DescribeSmsQuotas(request *DescribeSmsQuotasRequest) (response *DescribeSmsQuotasResponse, err error) {
    if request == nil {
        request = NewDescribeSmsQuotasRequest()
    }
    
    response = NewDescribeSmsQuotasResponse()
    err = c.Send(request, response)
    return
}

// DescribeSmsQuotas
// 查询后付费短信资源量
//
// 1 有免费包的返回SmsFreeQuota结构所有字段
//
// 2 没有免费包，有付费包，付费返回复用SmsFreeQuota结构，其中只有 TodayUsedQuota 字段有效
//
// 3 都没有返回为空数组
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_PARAM = "MissingParameter.Param"
func (c *Client) DescribeSmsQuotasWithContext(ctx context.Context, request *DescribeSmsQuotasRequest) (response *DescribeSmsQuotasResponse, err error) {
    if request == nil {
        request = NewDescribeSmsQuotasRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeSmsQuotasResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSpecialCostItemsRequest() (request *DescribeSpecialCostItemsRequest) {
    request = &DescribeSpecialCostItemsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcb", APIVersion, "DescribeSpecialCostItems")
    
    
    return
}

func NewDescribeSpecialCostItemsResponse() (response *DescribeSpecialCostItemsResponse) {
    response = &DescribeSpecialCostItemsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeSpecialCostItems
// 查询环境1分钱抵扣信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_PARAM = "MissingParameter.Param"
func (c *Client) DescribeSpecialCostItems(request *DescribeSpecialCostItemsRequest) (response *DescribeSpecialCostItemsResponse, err error) {
    if request == nil {
        request = NewDescribeSpecialCostItemsRequest()
    }
    
    response = NewDescribeSpecialCostItemsResponse()
    err = c.Send(request, response)
    return
}

// DescribeSpecialCostItems
// 查询环境1分钱抵扣信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_PARAM = "MissingParameter.Param"
func (c *Client) DescribeSpecialCostItemsWithContext(ctx context.Context, request *DescribeSpecialCostItemsRequest) (response *DescribeSpecialCostItemsResponse, err error) {
    if request == nil {
        request = NewDescribeSpecialCostItemsRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeSpecialCostItemsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeStandaloneGatewayRequest() (request *DescribeStandaloneGatewayRequest) {
    request = &DescribeStandaloneGatewayRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcb", APIVersion, "DescribeStandaloneGateway")
    
    
    return
}

func NewDescribeStandaloneGatewayResponse() (response *DescribeStandaloneGatewayResponse) {
    response = &DescribeStandaloneGatewayResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeStandaloneGateway
// 本接口（DescribeStandaloneGateway）查询小租户网关套餐信息。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ACTION = "InvalidParameter.Action"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeStandaloneGateway(request *DescribeStandaloneGatewayRequest) (response *DescribeStandaloneGatewayResponse, err error) {
    if request == nil {
        request = NewDescribeStandaloneGatewayRequest()
    }
    
    response = NewDescribeStandaloneGatewayResponse()
    err = c.Send(request, response)
    return
}

// DescribeStandaloneGateway
// 本接口（DescribeStandaloneGateway）查询小租户网关套餐信息。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ACTION = "InvalidParameter.Action"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeStandaloneGatewayWithContext(ctx context.Context, request *DescribeStandaloneGatewayRequest) (response *DescribeStandaloneGatewayResponse, err error) {
    if request == nil {
        request = NewDescribeStandaloneGatewayRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeStandaloneGatewayResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeStandaloneGatewayPackageRequest() (request *DescribeStandaloneGatewayPackageRequest) {
    request = &DescribeStandaloneGatewayPackageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcb", APIVersion, "DescribeStandaloneGatewayPackage")
    
    
    return
}

func NewDescribeStandaloneGatewayPackageResponse() (response *DescribeStandaloneGatewayPackageResponse) {
    response = &DescribeStandaloneGatewayPackageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeStandaloneGatewayPackage
// 本接口（DescribeStandaloneGatewayPackage）用于查询小租户网关套餐信息。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ACTION = "InvalidParameter.Action"
//  INVALIDPARAMETER_ENVID = "InvalidParameter.EnvId"
func (c *Client) DescribeStandaloneGatewayPackage(request *DescribeStandaloneGatewayPackageRequest) (response *DescribeStandaloneGatewayPackageResponse, err error) {
    if request == nil {
        request = NewDescribeStandaloneGatewayPackageRequest()
    }
    
    response = NewDescribeStandaloneGatewayPackageResponse()
    err = c.Send(request, response)
    return
}

// DescribeStandaloneGatewayPackage
// 本接口（DescribeStandaloneGatewayPackage）用于查询小租户网关套餐信息。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ACTION = "InvalidParameter.Action"
//  INVALIDPARAMETER_ENVID = "InvalidParameter.EnvId"
func (c *Client) DescribeStandaloneGatewayPackageWithContext(ctx context.Context, request *DescribeStandaloneGatewayPackageRequest) (response *DescribeStandaloneGatewayPackageResponse, err error) {
    if request == nil {
        request = NewDescribeStandaloneGatewayPackageRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeStandaloneGatewayPackageResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUserActivityInfoRequest() (request *DescribeUserActivityInfoRequest) {
    request = &DescribeUserActivityInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcb", APIVersion, "DescribeUserActivityInfo")
    
    
    return
}

func NewDescribeUserActivityInfoResponse() (response *DescribeUserActivityInfoResponse) {
    response = &DescribeUserActivityInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeUserActivityInfo
// 查询用户活动信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_PARAM = "MissingParameter.Param"
func (c *Client) DescribeUserActivityInfo(request *DescribeUserActivityInfoRequest) (response *DescribeUserActivityInfoResponse, err error) {
    if request == nil {
        request = NewDescribeUserActivityInfoRequest()
    }
    
    response = NewDescribeUserActivityInfoResponse()
    err = c.Send(request, response)
    return
}

// DescribeUserActivityInfo
// 查询用户活动信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_PARAM = "MissingParameter.Param"
func (c *Client) DescribeUserActivityInfoWithContext(ctx context.Context, request *DescribeUserActivityInfoRequest) (response *DescribeUserActivityInfoResponse, err error) {
    if request == nil {
        request = NewDescribeUserActivityInfoRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeUserActivityInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeWxCloudBaseRunEnvsRequest() (request *DescribeWxCloudBaseRunEnvsRequest) {
    request = &DescribeWxCloudBaseRunEnvsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcb", APIVersion, "DescribeWxCloudBaseRunEnvs")
    
    
    return
}

func NewDescribeWxCloudBaseRunEnvsResponse() (response *DescribeWxCloudBaseRunEnvsResponse) {
    response = &DescribeWxCloudBaseRunEnvsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeWxCloudBaseRunEnvs
// 查询微信云托管环境信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SYSTEMFAIL = "InternalError.SystemFail"
//  INTERNALERROR_TIMEOUT = "InternalError.Timeout"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ACTION = "InvalidParameter.Action"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_PARAM = "MissingParameter.Param"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeWxCloudBaseRunEnvs(request *DescribeWxCloudBaseRunEnvsRequest) (response *DescribeWxCloudBaseRunEnvsResponse, err error) {
    if request == nil {
        request = NewDescribeWxCloudBaseRunEnvsRequest()
    }
    
    response = NewDescribeWxCloudBaseRunEnvsResponse()
    err = c.Send(request, response)
    return
}

// DescribeWxCloudBaseRunEnvs
// 查询微信云托管环境信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SYSTEMFAIL = "InternalError.SystemFail"
//  INTERNALERROR_TIMEOUT = "InternalError.Timeout"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ACTION = "InvalidParameter.Action"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_PARAM = "MissingParameter.Param"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeWxCloudBaseRunEnvsWithContext(ctx context.Context, request *DescribeWxCloudBaseRunEnvsRequest) (response *DescribeWxCloudBaseRunEnvsResponse, err error) {
    if request == nil {
        request = NewDescribeWxCloudBaseRunEnvsRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeWxCloudBaseRunEnvsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeWxCloudBaseRunSubNetsRequest() (request *DescribeWxCloudBaseRunSubNetsRequest) {
    request = &DescribeWxCloudBaseRunSubNetsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcb", APIVersion, "DescribeWxCloudBaseRunSubNets")
    
    
    return
}

func NewDescribeWxCloudBaseRunSubNetsResponse() (response *DescribeWxCloudBaseRunSubNetsResponse) {
    response = &DescribeWxCloudBaseRunSubNetsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeWxCloudBaseRunSubNets
// 查询微信云托管子网
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SYSTEMFAIL = "InternalError.SystemFail"
//  INTERNALERROR_TIMEOUT = "InternalError.Timeout"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ACTION = "InvalidParameter.Action"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_PARAM = "MissingParameter.Param"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeWxCloudBaseRunSubNets(request *DescribeWxCloudBaseRunSubNetsRequest) (response *DescribeWxCloudBaseRunSubNetsResponse, err error) {
    if request == nil {
        request = NewDescribeWxCloudBaseRunSubNetsRequest()
    }
    
    response = NewDescribeWxCloudBaseRunSubNetsResponse()
    err = c.Send(request, response)
    return
}

// DescribeWxCloudBaseRunSubNets
// 查询微信云托管子网
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SYSTEMFAIL = "InternalError.SystemFail"
//  INTERNALERROR_TIMEOUT = "InternalError.Timeout"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ACTION = "InvalidParameter.Action"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_PARAM = "MissingParameter.Param"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeWxCloudBaseRunSubNetsWithContext(ctx context.Context, request *DescribeWxCloudBaseRunSubNetsRequest) (response *DescribeWxCloudBaseRunSubNetsResponse, err error) {
    if request == nil {
        request = NewDescribeWxCloudBaseRunSubNetsRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeWxCloudBaseRunSubNetsResponse()
    err = c.Send(request, response)
    return
}

func NewDestroyEnvRequest() (request *DestroyEnvRequest) {
    request = &DestroyEnvRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcb", APIVersion, "DestroyEnv")
    
    
    return
}

func NewDestroyEnvResponse() (response *DestroyEnvResponse) {
    response = &DestroyEnvResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DestroyEnv
// 销毁环境
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_PARAM = "MissingParameter.Param"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_INVOICEAMOUNTLACK = "ResourceUnavailable.InvoiceAmountLack"
//  RESOURCEUNAVAILABLE_RESOURCEOVERDUE = "ResourceUnavailable.ResourceOverdue"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DestroyEnv(request *DestroyEnvRequest) (response *DestroyEnvResponse, err error) {
    if request == nil {
        request = NewDestroyEnvRequest()
    }
    
    response = NewDestroyEnvResponse()
    err = c.Send(request, response)
    return
}

// DestroyEnv
// 销毁环境
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_PARAM = "MissingParameter.Param"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_INVOICEAMOUNTLACK = "ResourceUnavailable.InvoiceAmountLack"
//  RESOURCEUNAVAILABLE_RESOURCEOVERDUE = "ResourceUnavailable.ResourceOverdue"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DestroyEnvWithContext(ctx context.Context, request *DestroyEnvRequest) (response *DestroyEnvResponse, err error) {
    if request == nil {
        request = NewDestroyEnvRequest()
    }
    request.SetContext(ctx)
    
    response = NewDestroyEnvResponse()
    err = c.Send(request, response)
    return
}

func NewDestroyStandaloneGatewayRequest() (request *DestroyStandaloneGatewayRequest) {
    request = &DestroyStandaloneGatewayRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcb", APIVersion, "DestroyStandaloneGateway")
    
    
    return
}

func NewDestroyStandaloneGatewayResponse() (response *DestroyStandaloneGatewayResponse) {
    response = &DestroyStandaloneGatewayResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DestroyStandaloneGateway
// 本接口（DestroyStandaloneGateway）用于销毁小租户网关。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ACTION = "InvalidParameter.Action"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DestroyStandaloneGateway(request *DestroyStandaloneGatewayRequest) (response *DestroyStandaloneGatewayResponse, err error) {
    if request == nil {
        request = NewDestroyStandaloneGatewayRequest()
    }
    
    response = NewDestroyStandaloneGatewayResponse()
    err = c.Send(request, response)
    return
}

// DestroyStandaloneGateway
// 本接口（DestroyStandaloneGateway）用于销毁小租户网关。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ACTION = "InvalidParameter.Action"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DestroyStandaloneGatewayWithContext(ctx context.Context, request *DestroyStandaloneGatewayRequest) (response *DestroyStandaloneGatewayResponse, err error) {
    if request == nil {
        request = NewDestroyStandaloneGatewayRequest()
    }
    request.SetContext(ctx)
    
    response = NewDestroyStandaloneGatewayResponse()
    err = c.Send(request, response)
    return
}

func NewDestroyStaticStoreRequest() (request *DestroyStaticStoreRequest) {
    request = &DestroyStaticStoreRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcb", APIVersion, "DestroyStaticStore")
    
    
    return
}

func NewDestroyStaticStoreResponse() (response *DestroyStaticStoreResponse) {
    response = &DestroyStaticStoreResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DestroyStaticStore
// 销毁静态托管资源，该接口创建异步销毁任务，资源最终状态可从DestroyStaticStore接口查看
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DestroyStaticStore(request *DestroyStaticStoreRequest) (response *DestroyStaticStoreResponse, err error) {
    if request == nil {
        request = NewDestroyStaticStoreRequest()
    }
    
    response = NewDestroyStaticStoreResponse()
    err = c.Send(request, response)
    return
}

// DestroyStaticStore
// 销毁静态托管资源，该接口创建异步销毁任务，资源最终状态可从DestroyStaticStore接口查看
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DestroyStaticStoreWithContext(ctx context.Context, request *DestroyStaticStoreRequest) (response *DestroyStaticStoreResponse, err error) {
    if request == nil {
        request = NewDestroyStaticStoreRequest()
    }
    request.SetContext(ctx)
    
    response = NewDestroyStaticStoreResponse()
    err = c.Send(request, response)
    return
}

func NewEstablishCloudBaseRunServerRequest() (request *EstablishCloudBaseRunServerRequest) {
    request = &EstablishCloudBaseRunServerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcb", APIVersion, "EstablishCloudBaseRunServer")
    
    
    return
}

func NewEstablishCloudBaseRunServerResponse() (response *EstablishCloudBaseRunServerResponse) {
    response = &EstablishCloudBaseRunServerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// EstablishCloudBaseRunServer
// 创建云应用服务
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SERVICEEVIL = "InvalidParameter.ServiceEvil"
//  LIMITEXCEEDED_ERRNAMESPACEMAXLIMIT = "LimitExceeded.ErrNamespaceMaxLimit"
//  LIMITEXCEEDED_ERRREPOMAXLIMIT = "LimitExceeded.ErrRepoMaxLimit"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEUNAVAILABLE_CDNFREEZED = "ResourceUnavailable.CDNFreezed"
func (c *Client) EstablishCloudBaseRunServer(request *EstablishCloudBaseRunServerRequest) (response *EstablishCloudBaseRunServerResponse, err error) {
    if request == nil {
        request = NewEstablishCloudBaseRunServerRequest()
    }
    
    response = NewEstablishCloudBaseRunServerResponse()
    err = c.Send(request, response)
    return
}

// EstablishCloudBaseRunServer
// 创建云应用服务
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SERVICEEVIL = "InvalidParameter.ServiceEvil"
//  LIMITEXCEEDED_ERRNAMESPACEMAXLIMIT = "LimitExceeded.ErrNamespaceMaxLimit"
//  LIMITEXCEEDED_ERRREPOMAXLIMIT = "LimitExceeded.ErrRepoMaxLimit"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEUNAVAILABLE_CDNFREEZED = "ResourceUnavailable.CDNFreezed"
func (c *Client) EstablishCloudBaseRunServerWithContext(ctx context.Context, request *EstablishCloudBaseRunServerRequest) (response *EstablishCloudBaseRunServerResponse, err error) {
    if request == nil {
        request = NewEstablishCloudBaseRunServerRequest()
    }
    request.SetContext(ctx)
    
    response = NewEstablishCloudBaseRunServerResponse()
    err = c.Send(request, response)
    return
}

func NewEstablishWxGatewayRouteRequest() (request *EstablishWxGatewayRouteRequest) {
    request = &EstablishWxGatewayRouteRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcb", APIVersion, "EstablishWxGatewayRoute")
    
    
    return
}

func NewEstablishWxGatewayRouteResponse() (response *EstablishWxGatewayRouteResponse) {
    response = &EstablishWxGatewayRouteResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// EstablishWxGatewayRoute
// 创建或修改安全网关路由
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED_ERRNAMESPACEMAXLIMIT = "LimitExceeded.ErrNamespaceMaxLimit"
//  LIMITEXCEEDED_ERRREPOMAXLIMIT = "LimitExceeded.ErrRepoMaxLimit"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) EstablishWxGatewayRoute(request *EstablishWxGatewayRouteRequest) (response *EstablishWxGatewayRouteResponse, err error) {
    if request == nil {
        request = NewEstablishWxGatewayRouteRequest()
    }
    
    response = NewEstablishWxGatewayRouteResponse()
    err = c.Send(request, response)
    return
}

// EstablishWxGatewayRoute
// 创建或修改安全网关路由
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED_ERRNAMESPACEMAXLIMIT = "LimitExceeded.ErrNamespaceMaxLimit"
//  LIMITEXCEEDED_ERRREPOMAXLIMIT = "LimitExceeded.ErrRepoMaxLimit"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) EstablishWxGatewayRouteWithContext(ctx context.Context, request *EstablishWxGatewayRouteRequest) (response *EstablishWxGatewayRouteResponse, err error) {
    if request == nil {
        request = NewEstablishWxGatewayRouteRequest()
    }
    request.SetContext(ctx)
    
    response = NewEstablishWxGatewayRouteResponse()
    err = c.Send(request, response)
    return
}

func NewModifyCloudBaseRunServerFlowConfRequest() (request *ModifyCloudBaseRunServerFlowConfRequest) {
    request = &ModifyCloudBaseRunServerFlowConfRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcb", APIVersion, "ModifyCloudBaseRunServerFlowConf")
    
    
    return
}

func NewModifyCloudBaseRunServerFlowConfResponse() (response *ModifyCloudBaseRunServerFlowConfResponse) {
    response = &ModifyCloudBaseRunServerFlowConfResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyCloudBaseRunServerFlowConf
// 修改容器内的版本流量配置
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SERVICENOTEXIST = "InvalidParameter.ServiceNotExist"
func (c *Client) ModifyCloudBaseRunServerFlowConf(request *ModifyCloudBaseRunServerFlowConfRequest) (response *ModifyCloudBaseRunServerFlowConfResponse, err error) {
    if request == nil {
        request = NewModifyCloudBaseRunServerFlowConfRequest()
    }
    
    response = NewModifyCloudBaseRunServerFlowConfResponse()
    err = c.Send(request, response)
    return
}

// ModifyCloudBaseRunServerFlowConf
// 修改容器内的版本流量配置
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SERVICENOTEXIST = "InvalidParameter.ServiceNotExist"
func (c *Client) ModifyCloudBaseRunServerFlowConfWithContext(ctx context.Context, request *ModifyCloudBaseRunServerFlowConfRequest) (response *ModifyCloudBaseRunServerFlowConfResponse, err error) {
    if request == nil {
        request = NewModifyCloudBaseRunServerFlowConfRequest()
    }
    request.SetContext(ctx)
    
    response = NewModifyCloudBaseRunServerFlowConfResponse()
    err = c.Send(request, response)
    return
}

func NewModifyCloudBaseRunServerVersionRequest() (request *ModifyCloudBaseRunServerVersionRequest) {
    request = &ModifyCloudBaseRunServerVersionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcb", APIVersion, "ModifyCloudBaseRunServerVersion")
    
    
    return
}

func NewModifyCloudBaseRunServerVersionResponse() (response *ModifyCloudBaseRunServerVersionResponse) {
    response = &ModifyCloudBaseRunServerVersionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyCloudBaseRunServerVersion
// 修改服务版本的副本数，环境变量
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SERVICENOTEXIST = "InvalidParameter.ServiceNotExist"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyCloudBaseRunServerVersion(request *ModifyCloudBaseRunServerVersionRequest) (response *ModifyCloudBaseRunServerVersionResponse, err error) {
    if request == nil {
        request = NewModifyCloudBaseRunServerVersionRequest()
    }
    
    response = NewModifyCloudBaseRunServerVersionResponse()
    err = c.Send(request, response)
    return
}

// ModifyCloudBaseRunServerVersion
// 修改服务版本的副本数，环境变量
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SERVICENOTEXIST = "InvalidParameter.ServiceNotExist"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyCloudBaseRunServerVersionWithContext(ctx context.Context, request *ModifyCloudBaseRunServerVersionRequest) (response *ModifyCloudBaseRunServerVersionResponse, err error) {
    if request == nil {
        request = NewModifyCloudBaseRunServerVersionRequest()
    }
    request.SetContext(ctx)
    
    response = NewModifyCloudBaseRunServerVersionResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDatabaseACLRequest() (request *ModifyDatabaseACLRequest) {
    request = &ModifyDatabaseACLRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcb", APIVersion, "ModifyDatabaseACL")
    
    
    return
}

func NewModifyDatabaseACLResponse() (response *ModifyDatabaseACLResponse) {
    response = &ModifyDatabaseACLResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyDatabaseACL
// 修改数据库权限
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_PARAM = "MissingParameter.Param"
func (c *Client) ModifyDatabaseACL(request *ModifyDatabaseACLRequest) (response *ModifyDatabaseACLResponse, err error) {
    if request == nil {
        request = NewModifyDatabaseACLRequest()
    }
    
    response = NewModifyDatabaseACLResponse()
    err = c.Send(request, response)
    return
}

// ModifyDatabaseACL
// 修改数据库权限
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_PARAM = "MissingParameter.Param"
func (c *Client) ModifyDatabaseACLWithContext(ctx context.Context, request *ModifyDatabaseACLRequest) (response *ModifyDatabaseACLResponse, err error) {
    if request == nil {
        request = NewModifyDatabaseACLRequest()
    }
    request.SetContext(ctx)
    
    response = NewModifyDatabaseACLResponse()
    err = c.Send(request, response)
    return
}

func NewModifyEndUserRequest() (request *ModifyEndUserRequest) {
    request = &ModifyEndUserRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcb", APIVersion, "ModifyEndUser")
    
    
    return
}

func NewModifyEndUserResponse() (response *ModifyEndUserResponse) {
    response = &ModifyEndUserResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyEndUser
// 管理终端用户
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_PARAM = "MissingParameter.Param"
func (c *Client) ModifyEndUser(request *ModifyEndUserRequest) (response *ModifyEndUserResponse, err error) {
    if request == nil {
        request = NewModifyEndUserRequest()
    }
    
    response = NewModifyEndUserResponse()
    err = c.Send(request, response)
    return
}

// ModifyEndUser
// 管理终端用户
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_PARAM = "MissingParameter.Param"
func (c *Client) ModifyEndUserWithContext(ctx context.Context, request *ModifyEndUserRequest) (response *ModifyEndUserResponse, err error) {
    if request == nil {
        request = NewModifyEndUserRequest()
    }
    request.SetContext(ctx)
    
    response = NewModifyEndUserResponse()
    err = c.Send(request, response)
    return
}

func NewModifyEnvRequest() (request *ModifyEnvRequest) {
    request = &ModifyEnvRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcb", APIVersion, "ModifyEnv")
    
    
    return
}

func NewModifyEnvResponse() (response *ModifyEnvResponse) {
    response = &ModifyEnvResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyEnv
// 更新环境信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_PARAM = "MissingParameter.Param"
func (c *Client) ModifyEnv(request *ModifyEnvRequest) (response *ModifyEnvResponse, err error) {
    if request == nil {
        request = NewModifyEnvRequest()
    }
    
    response = NewModifyEnvResponse()
    err = c.Send(request, response)
    return
}

// ModifyEnv
// 更新环境信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_PARAM = "MissingParameter.Param"
func (c *Client) ModifyEnvWithContext(ctx context.Context, request *ModifyEnvRequest) (response *ModifyEnvResponse, err error) {
    if request == nil {
        request = NewModifyEnvRequest()
    }
    request.SetContext(ctx)
    
    response = NewModifyEnvResponse()
    err = c.Send(request, response)
    return
}

func NewReinstateEnvRequest() (request *ReinstateEnvRequest) {
    request = &ReinstateEnvRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcb", APIVersion, "ReinstateEnv")
    
    
    return
}

func NewReinstateEnvResponse() (response *ReinstateEnvResponse) {
    response = &ReinstateEnvResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ReinstateEnv
// 针对已隔离的免费环境，可以通过本接口将其恢复访问。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ENVID = "InvalidParameter.EnvId"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ReinstateEnv(request *ReinstateEnvRequest) (response *ReinstateEnvResponse, err error) {
    if request == nil {
        request = NewReinstateEnvRequest()
    }
    
    response = NewReinstateEnvResponse()
    err = c.Send(request, response)
    return
}

// ReinstateEnv
// 针对已隔离的免费环境，可以通过本接口将其恢复访问。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ENVID = "InvalidParameter.EnvId"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ReinstateEnvWithContext(ctx context.Context, request *ReinstateEnvRequest) (response *ReinstateEnvResponse, err error) {
    if request == nil {
        request = NewReinstateEnvRequest()
    }
    request.SetContext(ctx)
    
    response = NewReinstateEnvResponse()
    err = c.Send(request, response)
    return
}

func NewReplaceActivityRecordRequest() (request *ReplaceActivityRecordRequest) {
    request = &ReplaceActivityRecordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcb", APIVersion, "ReplaceActivityRecord")
    
    
    return
}

func NewReplaceActivityRecordResponse() (response *ReplaceActivityRecordResponse) {
    response = &ReplaceActivityRecordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ReplaceActivityRecord
// 更新活动详情
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_PARAM = "MissingParameter.Param"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) ReplaceActivityRecord(request *ReplaceActivityRecordRequest) (response *ReplaceActivityRecordResponse, err error) {
    if request == nil {
        request = NewReplaceActivityRecordRequest()
    }
    
    response = NewReplaceActivityRecordResponse()
    err = c.Send(request, response)
    return
}

// ReplaceActivityRecord
// 更新活动详情
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_PARAM = "MissingParameter.Param"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) ReplaceActivityRecordWithContext(ctx context.Context, request *ReplaceActivityRecordRequest) (response *ReplaceActivityRecordResponse, err error) {
    if request == nil {
        request = NewReplaceActivityRecordRequest()
    }
    request.SetContext(ctx)
    
    response = NewReplaceActivityRecordResponse()
    err = c.Send(request, response)
    return
}

func NewRollUpdateCloudBaseRunServerVersionRequest() (request *RollUpdateCloudBaseRunServerVersionRequest) {
    request = &RollUpdateCloudBaseRunServerVersionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcb", APIVersion, "RollUpdateCloudBaseRunServerVersion")
    
    
    return
}

func NewRollUpdateCloudBaseRunServerVersionResponse() (response *RollUpdateCloudBaseRunServerVersionResponse) {
    response = &RollUpdateCloudBaseRunServerVersionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// RollUpdateCloudBaseRunServerVersion
// 针对特定的版本，进行滚动更新
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SERVICENOTEXIST = "InvalidParameter.ServiceNotExist"
func (c *Client) RollUpdateCloudBaseRunServerVersion(request *RollUpdateCloudBaseRunServerVersionRequest) (response *RollUpdateCloudBaseRunServerVersionResponse, err error) {
    if request == nil {
        request = NewRollUpdateCloudBaseRunServerVersionRequest()
    }
    
    response = NewRollUpdateCloudBaseRunServerVersionResponse()
    err = c.Send(request, response)
    return
}

// RollUpdateCloudBaseRunServerVersion
// 针对特定的版本，进行滚动更新
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SERVICENOTEXIST = "InvalidParameter.ServiceNotExist"
func (c *Client) RollUpdateCloudBaseRunServerVersionWithContext(ctx context.Context, request *RollUpdateCloudBaseRunServerVersionRequest) (response *RollUpdateCloudBaseRunServerVersionResponse, err error) {
    if request == nil {
        request = NewRollUpdateCloudBaseRunServerVersionRequest()
    }
    request.SetContext(ctx)
    
    response = NewRollUpdateCloudBaseRunServerVersionResponse()
    err = c.Send(request, response)
    return
}

func NewSearchClsLogRequest() (request *SearchClsLogRequest) {
    request = &SearchClsLogRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcb", APIVersion, "SearchClsLog")
    
    
    return
}

func NewSearchClsLogResponse() (response *SearchClsLogResponse) {
    response = &SearchClsLogResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// SearchClsLog
// 搜索CLS日志，TCB角色秘钥访问
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) SearchClsLog(request *SearchClsLogRequest) (response *SearchClsLogResponse, err error) {
    if request == nil {
        request = NewSearchClsLogRequest()
    }
    
    response = NewSearchClsLogResponse()
    err = c.Send(request, response)
    return
}

// SearchClsLog
// 搜索CLS日志，TCB角色秘钥访问
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) SearchClsLogWithContext(ctx context.Context, request *SearchClsLogRequest) (response *SearchClsLogResponse, err error) {
    if request == nil {
        request = NewSearchClsLogRequest()
    }
    request.SetContext(ctx)
    
    response = NewSearchClsLogResponse()
    err = c.Send(request, response)
    return
}

func NewTurnOffStandaloneGatewayRequest() (request *TurnOffStandaloneGatewayRequest) {
    request = &TurnOffStandaloneGatewayRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcb", APIVersion, "TurnOffStandaloneGateway")
    
    
    return
}

func NewTurnOffStandaloneGatewayResponse() (response *TurnOffStandaloneGatewayResponse) {
    response = &TurnOffStandaloneGatewayResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// TurnOffStandaloneGateway
// 本接口（TurnOffStandaloneGateway）用于关闭小租户网关。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PARTIALFAILURE = "FailedOperation.PartialFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ACTION = "InvalidParameter.Action"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) TurnOffStandaloneGateway(request *TurnOffStandaloneGatewayRequest) (response *TurnOffStandaloneGatewayResponse, err error) {
    if request == nil {
        request = NewTurnOffStandaloneGatewayRequest()
    }
    
    response = NewTurnOffStandaloneGatewayResponse()
    err = c.Send(request, response)
    return
}

// TurnOffStandaloneGateway
// 本接口（TurnOffStandaloneGateway）用于关闭小租户网关。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PARTIALFAILURE = "FailedOperation.PartialFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ACTION = "InvalidParameter.Action"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) TurnOffStandaloneGatewayWithContext(ctx context.Context, request *TurnOffStandaloneGatewayRequest) (response *TurnOffStandaloneGatewayResponse, err error) {
    if request == nil {
        request = NewTurnOffStandaloneGatewayRequest()
    }
    request.SetContext(ctx)
    
    response = NewTurnOffStandaloneGatewayResponse()
    err = c.Send(request, response)
    return
}

func NewTurnOnStandaloneGatewayRequest() (request *TurnOnStandaloneGatewayRequest) {
    request = &TurnOnStandaloneGatewayRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcb", APIVersion, "TurnOnStandaloneGateway")
    
    
    return
}

func NewTurnOnStandaloneGatewayResponse() (response *TurnOnStandaloneGatewayResponse) {
    response = &TurnOnStandaloneGatewayResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// TurnOnStandaloneGateway
// 本接口（TurnOnStandaloneGateway）用于开启小租户网关。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ACTION = "InvalidParameter.Action"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) TurnOnStandaloneGateway(request *TurnOnStandaloneGatewayRequest) (response *TurnOnStandaloneGatewayResponse, err error) {
    if request == nil {
        request = NewTurnOnStandaloneGatewayRequest()
    }
    
    response = NewTurnOnStandaloneGatewayResponse()
    err = c.Send(request, response)
    return
}

// TurnOnStandaloneGateway
// 本接口（TurnOnStandaloneGateway）用于开启小租户网关。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ACTION = "InvalidParameter.Action"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) TurnOnStandaloneGatewayWithContext(ctx context.Context, request *TurnOnStandaloneGatewayRequest) (response *TurnOnStandaloneGatewayResponse, err error) {
    if request == nil {
        request = NewTurnOnStandaloneGatewayRequest()
    }
    request.SetContext(ctx)
    
    response = NewTurnOnStandaloneGatewayResponse()
    err = c.Send(request, response)
    return
}
