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
func (c *Client) CommonServiceAPI(request *CommonServiceAPIRequest) (response *CommonServiceAPIResponse, err error) {
    if request == nil {
        request = NewCommonServiceAPIRequest()
    }
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
//  RESOURCEINUSE = "ResourceInUse"
func (c *Client) CreateAuthDomain(request *CreateAuthDomainRequest) (response *CreateAuthDomainResponse, err error) {
    if request == nil {
        request = NewCreateAuthDomainRequest()
    }
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
func (c *Client) DescribeHostingDomainTask(request *DescribeHostingDomainTaskRequest) (response *DescribeHostingDomainTaskResponse, err error) {
    if request == nil {
        request = NewDescribeHostingDomainTaskRequest()
    }
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
func (c *Client) ReplaceActivityRecord(request *ReplaceActivityRecordRequest) (response *ReplaceActivityRecordResponse, err error) {
    if request == nil {
        request = NewReplaceActivityRecordRequest()
    }
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
