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

package v20191012

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2019-10-12"

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


func NewDescribeDomainsRequest() (request *DescribeDomainsRequest) {
    request = &DescribeDomainsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ecdn", APIVersion, "DescribeDomains")
    
    
    return
}

func NewDescribeDomainsResponse() (response *DescribeDomainsResponse) {
    response = &DescribeDomainsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDomains
// 本接口（DescribeDomains）用于查询CDN域名基本信息，包括项目id，状态，业务类型，创建时间，更新时间等。
//
// 
//
// >?  若您的业务已迁移至 CDN 控制台，请参考<a href="https://cloud.tencent.com/document/api/228/41118"> CDN 接口文档</a>，使用  CDN 相关API 进行操作。
//
// 可能返回的错误码:
//  INTERNALERROR_CAMSYSTEMERROR = "InternalError.CamSystemError"
//  INTERNALERROR_ECDNCONFIGERROR = "InternalError.EcdnConfigError"
//  INTERNALERROR_ECDNDBERROR = "InternalError.EcdnDbError"
//  INTERNALERROR_ECDNQUERYSYSTEMERROR = "InternalError.EcdnQuerySystemError"
//  INTERNALERROR_ECDNSYSTEMERROR = "InternalError.EcdnSystemError"
//  INTERNALERROR_ERROR = "InternalError.Error"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INVALIDPARAMETER_ECDNINTERFACEERROR = "InvalidParameter.EcdnInterfaceError"
//  INVALIDPARAMETER_ECDNPARAMERROR = "InvalidParameter.EcdnParamError"
//  LIMITEXCEEDED_ECDNDOMAINOPTOOOFTEN = "LimitExceeded.EcdnDomainOpTooOften"
//  RESOURCENOTFOUND_ECDNDOMAINNOTEXISTS = "ResourceNotFound.EcdnDomainNotExists"
//  RESOURCENOTFOUND_ECDNUSERNOTEXISTS = "ResourceNotFound.EcdnUserNotExists"
//  UNAUTHORIZEDOPERATION_ECDNCAMUNAUTHORIZED = "UnauthorizedOperation.EcdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
func (c *Client) DescribeDomains(request *DescribeDomainsRequest) (response *DescribeDomainsResponse, err error) {
    return c.DescribeDomainsWithContext(context.Background(), request)
}

// DescribeDomains
// 本接口（DescribeDomains）用于查询CDN域名基本信息，包括项目id，状态，业务类型，创建时间，更新时间等。
//
// 
//
// >?  若您的业务已迁移至 CDN 控制台，请参考<a href="https://cloud.tencent.com/document/api/228/41118"> CDN 接口文档</a>，使用  CDN 相关API 进行操作。
//
// 可能返回的错误码:
//  INTERNALERROR_CAMSYSTEMERROR = "InternalError.CamSystemError"
//  INTERNALERROR_ECDNCONFIGERROR = "InternalError.EcdnConfigError"
//  INTERNALERROR_ECDNDBERROR = "InternalError.EcdnDbError"
//  INTERNALERROR_ECDNQUERYSYSTEMERROR = "InternalError.EcdnQuerySystemError"
//  INTERNALERROR_ECDNSYSTEMERROR = "InternalError.EcdnSystemError"
//  INTERNALERROR_ERROR = "InternalError.Error"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INVALIDPARAMETER_ECDNINTERFACEERROR = "InvalidParameter.EcdnInterfaceError"
//  INVALIDPARAMETER_ECDNPARAMERROR = "InvalidParameter.EcdnParamError"
//  LIMITEXCEEDED_ECDNDOMAINOPTOOOFTEN = "LimitExceeded.EcdnDomainOpTooOften"
//  RESOURCENOTFOUND_ECDNDOMAINNOTEXISTS = "ResourceNotFound.EcdnDomainNotExists"
//  RESOURCENOTFOUND_ECDNUSERNOTEXISTS = "ResourceNotFound.EcdnUserNotExists"
//  UNAUTHORIZEDOPERATION_ECDNCAMUNAUTHORIZED = "UnauthorizedOperation.EcdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
func (c *Client) DescribeDomainsWithContext(ctx context.Context, request *DescribeDomainsRequest) (response *DescribeDomainsResponse, err error) {
    if request == nil {
        request = NewDescribeDomainsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDomains require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDomainsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDomainsConfigRequest() (request *DescribeDomainsConfigRequest) {
    request = &DescribeDomainsConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ecdn", APIVersion, "DescribeDomainsConfig")
    
    
    return
}

func NewDescribeDomainsConfigResponse() (response *DescribeDomainsConfigResponse) {
    response = &DescribeDomainsConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDomainsConfig
// 本接口（DescribeDomainsConfig）用于查询CDN加速域名详细配置信息。
//
// 
//
// >?  若您的业务已迁移至 CDN 控制台，请参考<a href="https://cloud.tencent.com/document/api/228/41117"> CDN 接口文档</a>，使用  CDN 相关API 进行操作。
//
// 可能返回的错误码:
//  FAILEDOPERATION_ECDNCONFIGERROR = "FailedOperation.EcdnConfigError"
//  INTERNALERROR_CAMSYSTEMERROR = "InternalError.CamSystemError"
//  INTERNALERROR_ECDNCONFIGERROR = "InternalError.EcdnConfigError"
//  INTERNALERROR_ECDNDBERROR = "InternalError.EcdnDbError"
//  INTERNALERROR_ECDNQUERYSYSTEMERROR = "InternalError.EcdnQuerySystemError"
//  INTERNALERROR_ECDNSYSTEMERROR = "InternalError.EcdnSystemError"
//  INTERNALERROR_ERROR = "InternalError.Error"
//  INVALIDPARAMETER_ECDNCERTNOCERTINFO = "InvalidParameter.EcdnCertNoCertInfo"
//  INVALIDPARAMETER_ECDNCONFIGINVALIDCACHE = "InvalidParameter.EcdnConfigInvalidCache"
//  INVALIDPARAMETER_ECDNINTERFACEERROR = "InvalidParameter.EcdnInterfaceError"
//  INVALIDPARAMETER_ECDNPARAMERROR = "InvalidParameter.EcdnParamError"
//  LIMITEXCEEDED_ECDNDOMAINOPTOOOFTEN = "LimitExceeded.EcdnDomainOpTooOften"
//  RESOURCENOTFOUND_ECDNDOMAINNOTEXISTS = "ResourceNotFound.EcdnDomainNotExists"
//  RESOURCENOTFOUND_ECDNUSERNOTEXISTS = "ResourceNotFound.EcdnUserNotExists"
//  UNAUTHORIZEDOPERATION_ECDNCAMUNAUTHORIZED = "UnauthorizedOperation.EcdnCamUnauthorized"
func (c *Client) DescribeDomainsConfig(request *DescribeDomainsConfigRequest) (response *DescribeDomainsConfigResponse, err error) {
    return c.DescribeDomainsConfigWithContext(context.Background(), request)
}

// DescribeDomainsConfig
// 本接口（DescribeDomainsConfig）用于查询CDN加速域名详细配置信息。
//
// 
//
// >?  若您的业务已迁移至 CDN 控制台，请参考<a href="https://cloud.tencent.com/document/api/228/41117"> CDN 接口文档</a>，使用  CDN 相关API 进行操作。
//
// 可能返回的错误码:
//  FAILEDOPERATION_ECDNCONFIGERROR = "FailedOperation.EcdnConfigError"
//  INTERNALERROR_CAMSYSTEMERROR = "InternalError.CamSystemError"
//  INTERNALERROR_ECDNCONFIGERROR = "InternalError.EcdnConfigError"
//  INTERNALERROR_ECDNDBERROR = "InternalError.EcdnDbError"
//  INTERNALERROR_ECDNQUERYSYSTEMERROR = "InternalError.EcdnQuerySystemError"
//  INTERNALERROR_ECDNSYSTEMERROR = "InternalError.EcdnSystemError"
//  INTERNALERROR_ERROR = "InternalError.Error"
//  INVALIDPARAMETER_ECDNCERTNOCERTINFO = "InvalidParameter.EcdnCertNoCertInfo"
//  INVALIDPARAMETER_ECDNCONFIGINVALIDCACHE = "InvalidParameter.EcdnConfigInvalidCache"
//  INVALIDPARAMETER_ECDNINTERFACEERROR = "InvalidParameter.EcdnInterfaceError"
//  INVALIDPARAMETER_ECDNPARAMERROR = "InvalidParameter.EcdnParamError"
//  LIMITEXCEEDED_ECDNDOMAINOPTOOOFTEN = "LimitExceeded.EcdnDomainOpTooOften"
//  RESOURCENOTFOUND_ECDNDOMAINNOTEXISTS = "ResourceNotFound.EcdnDomainNotExists"
//  RESOURCENOTFOUND_ECDNUSERNOTEXISTS = "ResourceNotFound.EcdnUserNotExists"
//  UNAUTHORIZEDOPERATION_ECDNCAMUNAUTHORIZED = "UnauthorizedOperation.EcdnCamUnauthorized"
func (c *Client) DescribeDomainsConfigWithContext(ctx context.Context, request *DescribeDomainsConfigRequest) (response *DescribeDomainsConfigResponse, err error) {
    if request == nil {
        request = NewDescribeDomainsConfigRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDomainsConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDomainsConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEcdnDomainLogsRequest() (request *DescribeEcdnDomainLogsRequest) {
    request = &DescribeEcdnDomainLogsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ecdn", APIVersion, "DescribeEcdnDomainLogs")
    
    
    return
}

func NewDescribeEcdnDomainLogsResponse() (response *DescribeEcdnDomainLogsResponse) {
    response = &DescribeEcdnDomainLogsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeEcdnDomainLogs
// 本接口（DescribeEcdnDomainLogs）用于查询域名的访问日志下载地址。
//
// 可能返回的错误码:
//  INTERNALERROR_ECDNCONFIGERROR = "InternalError.EcdnConfigError"
//  INTERNALERROR_ECDNDBERROR = "InternalError.EcdnDbError"
//  INVALIDPARAMETER_ECDNPARAMERROR = "InvalidParameter.EcdnParamError"
//  INVALIDPARAMETER_ECDNSTATINVALIDDATE = "InvalidParameter.EcdnStatInvalidDate"
//  LIMITEXCEEDED_ECDNDOMAINOPTOOOFTEN = "LimitExceeded.EcdnDomainOpTooOften"
//  RESOURCENOTFOUND_ECDNDOMAINNOTEXISTS = "ResourceNotFound.EcdnDomainNotExists"
//  RESOURCENOTFOUND_ECDNHOSTNOTEXISTS = "ResourceNotFound.EcdnHostNotExists"
//  RESOURCENOTFOUND_ECDNUSERNOTEXISTS = "ResourceNotFound.EcdnUserNotExists"
//  UNAUTHORIZEDOPERATION_ECDNACCOUNTUNAUTHORIZED = "UnauthorizedOperation.EcdnAccountUnauthorized"
//  UNAUTHORIZEDOPERATION_ECDNCAMUNAUTHORIZED = "UnauthorizedOperation.EcdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_ECDNDOMAINUNAUTHORIZED = "UnauthorizedOperation.EcdnDomainUnauthorized"
//  UNAUTHORIZEDOPERATION_ECDNHOSTUNAUTHORIZED = "UnauthorizedOperation.EcdnHostUnauthorized"
//  UNAUTHORIZEDOPERATION_ECDNNODOMAINUNAUTHORIZED = "UnauthorizedOperation.EcdnNoDomainUnauthorized"
func (c *Client) DescribeEcdnDomainLogs(request *DescribeEcdnDomainLogsRequest) (response *DescribeEcdnDomainLogsResponse, err error) {
    return c.DescribeEcdnDomainLogsWithContext(context.Background(), request)
}

// DescribeEcdnDomainLogs
// 本接口（DescribeEcdnDomainLogs）用于查询域名的访问日志下载地址。
//
// 可能返回的错误码:
//  INTERNALERROR_ECDNCONFIGERROR = "InternalError.EcdnConfigError"
//  INTERNALERROR_ECDNDBERROR = "InternalError.EcdnDbError"
//  INVALIDPARAMETER_ECDNPARAMERROR = "InvalidParameter.EcdnParamError"
//  INVALIDPARAMETER_ECDNSTATINVALIDDATE = "InvalidParameter.EcdnStatInvalidDate"
//  LIMITEXCEEDED_ECDNDOMAINOPTOOOFTEN = "LimitExceeded.EcdnDomainOpTooOften"
//  RESOURCENOTFOUND_ECDNDOMAINNOTEXISTS = "ResourceNotFound.EcdnDomainNotExists"
//  RESOURCENOTFOUND_ECDNHOSTNOTEXISTS = "ResourceNotFound.EcdnHostNotExists"
//  RESOURCENOTFOUND_ECDNUSERNOTEXISTS = "ResourceNotFound.EcdnUserNotExists"
//  UNAUTHORIZEDOPERATION_ECDNACCOUNTUNAUTHORIZED = "UnauthorizedOperation.EcdnAccountUnauthorized"
//  UNAUTHORIZEDOPERATION_ECDNCAMUNAUTHORIZED = "UnauthorizedOperation.EcdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_ECDNDOMAINUNAUTHORIZED = "UnauthorizedOperation.EcdnDomainUnauthorized"
//  UNAUTHORIZEDOPERATION_ECDNHOSTUNAUTHORIZED = "UnauthorizedOperation.EcdnHostUnauthorized"
//  UNAUTHORIZEDOPERATION_ECDNNODOMAINUNAUTHORIZED = "UnauthorizedOperation.EcdnNoDomainUnauthorized"
func (c *Client) DescribeEcdnDomainLogsWithContext(ctx context.Context, request *DescribeEcdnDomainLogsRequest) (response *DescribeEcdnDomainLogsResponse, err error) {
    if request == nil {
        request = NewDescribeEcdnDomainLogsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeEcdnDomainLogs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeEcdnDomainLogsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEcdnDomainStatisticsRequest() (request *DescribeEcdnDomainStatisticsRequest) {
    request = &DescribeEcdnDomainStatisticsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ecdn", APIVersion, "DescribeEcdnDomainStatistics")
    
    
    return
}

func NewDescribeEcdnDomainStatisticsResponse() (response *DescribeEcdnDomainStatisticsResponse) {
    response = &DescribeEcdnDomainStatisticsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeEcdnDomainStatistics
// 本接口（DescribeEcdnDomainStatistics）用于查询指定时间段内的域名访问统计指标。
//
// 
//
// >?  若您的业务已迁移至 CDN 控制台，请参考<a href="https://cloud.tencent.com/document/api/228/30986"> CDN 接口文档</a>，使用  CDN 相关API 进行操作。
//
// 可能返回的错误码:
//  INTERNALERROR_ECDNDBERROR = "InternalError.EcdnDbError"
//  INTERNALERROR_ECDNSYSTEMERROR = "InternalError.EcdnSystemError"
//  INVALIDPARAMETER_ECDNPARAMERROR = "InvalidParameter.EcdnParamError"
//  INVALIDPARAMETER_ECDNSTATINVALIDDATE = "InvalidParameter.EcdnStatInvalidDate"
//  INVALIDPARAMETER_ECDNSTATINVALIDMETRIC = "InvalidParameter.EcdnStatInvalidMetric"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  LIMITEXCEEDED_ECDNDOMAINOPTOOOFTEN = "LimitExceeded.EcdnDomainOpTooOften"
//  RESOURCENOTFOUND_ECDNDOMAINNOTEXISTS = "ResourceNotFound.EcdnDomainNotExists"
//  RESOURCENOTFOUND_ECDNPROJECTNOTEXISTS = "ResourceNotFound.EcdnProjectNotExists"
//  RESOURCENOTFOUND_ECDNUSERNOTEXISTS = "ResourceNotFound.EcdnUserNotExists"
//  UNAUTHORIZEDOPERATION_CDNACCOUNTUNAUTHORIZED = "UnauthorizedOperation.CdnAccountUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNDOMAINUNAUTHORIZED = "UnauthorizedOperation.CdnDomainUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNHOSTUNAUTHORIZED = "UnauthorizedOperation.CdnHostUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNNODOMAINUNAUTHORIZED = "UnauthorizedOperation.CdnNoDomainUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNPROJECTUNAUTHORIZED = "UnauthorizedOperation.CdnProjectUnauthorized"
//  UNAUTHORIZEDOPERATION_DOMAINNOPERMISSION = "UnauthorizedOperation.DomainNoPermission"
//  UNAUTHORIZEDOPERATION_DOMAINSNOPERMISSION = "UnauthorizedOperation.DomainsNoPermission"
//  UNAUTHORIZEDOPERATION_ECDNACCOUNTUNAUTHORIZED = "UnauthorizedOperation.EcdnAccountUnauthorized"
//  UNAUTHORIZEDOPERATION_ECDNCAMUNAUTHORIZED = "UnauthorizedOperation.EcdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_ECDNDOMAINUNAUTHORIZED = "UnauthorizedOperation.EcdnDomainUnauthorized"
//  UNAUTHORIZEDOPERATION_ECDNHOSTUNAUTHORIZED = "UnauthorizedOperation.EcdnHostUnauthorized"
//  UNAUTHORIZEDOPERATION_ECDNNODOMAINUNAUTHORIZED = "UnauthorizedOperation.EcdnNoDomainUnauthorized"
//  UNAUTHORIZEDOPERATION_ECDNPROJECTUNAUTHORIZED = "UnauthorizedOperation.EcdnProjectUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNAUTHORIZEDOPERATION_PROJECTNOPERMISSION = "UnauthorizedOperation.ProjectNoPermission"
//  UNAUTHORIZEDOPERATION_PROJECTSNOPERMISSION = "UnauthorizedOperation.ProjectsNoPermission"
func (c *Client) DescribeEcdnDomainStatistics(request *DescribeEcdnDomainStatisticsRequest) (response *DescribeEcdnDomainStatisticsResponse, err error) {
    return c.DescribeEcdnDomainStatisticsWithContext(context.Background(), request)
}

// DescribeEcdnDomainStatistics
// 本接口（DescribeEcdnDomainStatistics）用于查询指定时间段内的域名访问统计指标。
//
// 
//
// >?  若您的业务已迁移至 CDN 控制台，请参考<a href="https://cloud.tencent.com/document/api/228/30986"> CDN 接口文档</a>，使用  CDN 相关API 进行操作。
//
// 可能返回的错误码:
//  INTERNALERROR_ECDNDBERROR = "InternalError.EcdnDbError"
//  INTERNALERROR_ECDNSYSTEMERROR = "InternalError.EcdnSystemError"
//  INVALIDPARAMETER_ECDNPARAMERROR = "InvalidParameter.EcdnParamError"
//  INVALIDPARAMETER_ECDNSTATINVALIDDATE = "InvalidParameter.EcdnStatInvalidDate"
//  INVALIDPARAMETER_ECDNSTATINVALIDMETRIC = "InvalidParameter.EcdnStatInvalidMetric"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  LIMITEXCEEDED_ECDNDOMAINOPTOOOFTEN = "LimitExceeded.EcdnDomainOpTooOften"
//  RESOURCENOTFOUND_ECDNDOMAINNOTEXISTS = "ResourceNotFound.EcdnDomainNotExists"
//  RESOURCENOTFOUND_ECDNPROJECTNOTEXISTS = "ResourceNotFound.EcdnProjectNotExists"
//  RESOURCENOTFOUND_ECDNUSERNOTEXISTS = "ResourceNotFound.EcdnUserNotExists"
//  UNAUTHORIZEDOPERATION_CDNACCOUNTUNAUTHORIZED = "UnauthorizedOperation.CdnAccountUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNDOMAINUNAUTHORIZED = "UnauthorizedOperation.CdnDomainUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNHOSTUNAUTHORIZED = "UnauthorizedOperation.CdnHostUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNNODOMAINUNAUTHORIZED = "UnauthorizedOperation.CdnNoDomainUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNPROJECTUNAUTHORIZED = "UnauthorizedOperation.CdnProjectUnauthorized"
//  UNAUTHORIZEDOPERATION_DOMAINNOPERMISSION = "UnauthorizedOperation.DomainNoPermission"
//  UNAUTHORIZEDOPERATION_DOMAINSNOPERMISSION = "UnauthorizedOperation.DomainsNoPermission"
//  UNAUTHORIZEDOPERATION_ECDNACCOUNTUNAUTHORIZED = "UnauthorizedOperation.EcdnAccountUnauthorized"
//  UNAUTHORIZEDOPERATION_ECDNCAMUNAUTHORIZED = "UnauthorizedOperation.EcdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_ECDNDOMAINUNAUTHORIZED = "UnauthorizedOperation.EcdnDomainUnauthorized"
//  UNAUTHORIZEDOPERATION_ECDNHOSTUNAUTHORIZED = "UnauthorizedOperation.EcdnHostUnauthorized"
//  UNAUTHORIZEDOPERATION_ECDNNODOMAINUNAUTHORIZED = "UnauthorizedOperation.EcdnNoDomainUnauthorized"
//  UNAUTHORIZEDOPERATION_ECDNPROJECTUNAUTHORIZED = "UnauthorizedOperation.EcdnProjectUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNAUTHORIZEDOPERATION_PROJECTNOPERMISSION = "UnauthorizedOperation.ProjectNoPermission"
//  UNAUTHORIZEDOPERATION_PROJECTSNOPERMISSION = "UnauthorizedOperation.ProjectsNoPermission"
func (c *Client) DescribeEcdnDomainStatisticsWithContext(ctx context.Context, request *DescribeEcdnDomainStatisticsRequest) (response *DescribeEcdnDomainStatisticsResponse, err error) {
    if request == nil {
        request = NewDescribeEcdnDomainStatisticsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeEcdnDomainStatistics require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeEcdnDomainStatisticsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEcdnStatisticsRequest() (request *DescribeEcdnStatisticsRequest) {
    request = &DescribeEcdnStatisticsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ecdn", APIVersion, "DescribeEcdnStatistics")
    
    
    return
}

func NewDescribeEcdnStatisticsResponse() (response *DescribeEcdnStatisticsResponse) {
    response = &DescribeEcdnStatisticsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeEcdnStatistics
// DescribeEcdnStatistics用于查询 ECDN 实时访问监控数据，支持以下指标查询：
//
// 
//
// + 流量（单位为 byte）
//
// + 带宽（单位为 bps）
//
// + 请求数（单位为 次）
//
// + 状态码 2xx 汇总及各 2 开头状态码明细（单位为 个）
//
// + 状态码 3xx 汇总及各 3 开头状态码明细（单位为 个）
//
// + 状态码 4xx 汇总及各 4 开头状态码明细（单位为 个）
//
// + 状态码 5xx 汇总及各 5 开头状态码明细（单位为 个）
//
// 可能返回的错误码:
//  INTERNALERROR_ECDNCONFIGERROR = "InternalError.EcdnConfigError"
//  INTERNALERROR_ECDNDBERROR = "InternalError.EcdnDbError"
//  INTERNALERROR_ECDNSYSTEMERROR = "InternalError.EcdnSystemError"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INVALIDPARAMETER_ECDNINVALIDPARAMINTERVAL = "InvalidParameter.EcdnInvalidParamInterval"
//  INVALIDPARAMETER_ECDNPARAMERROR = "InvalidParameter.EcdnParamError"
//  INVALIDPARAMETER_ECDNSTATINVALIDDATE = "InvalidParameter.EcdnStatInvalidDate"
//  INVALIDPARAMETER_ECDNSTATINVALIDMETRIC = "InvalidParameter.EcdnStatInvalidMetric"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  LIMITEXCEEDED_ECDNDOMAINOPTOOOFTEN = "LimitExceeded.EcdnDomainOpTooOften"
//  RESOURCENOTFOUND_ECDNDOMAINNOTEXISTS = "ResourceNotFound.EcdnDomainNotExists"
//  RESOURCENOTFOUND_ECDNHOSTNOTEXISTS = "ResourceNotFound.EcdnHostNotExists"
//  RESOURCENOTFOUND_ECDNPROJECTNOTEXISTS = "ResourceNotFound.EcdnProjectNotExists"
//  RESOURCENOTFOUND_ECDNUSERNOTEXISTS = "ResourceNotFound.EcdnUserNotExists"
//  UNAUTHORIZEDOPERATION_CDNACCOUNTUNAUTHORIZED = "UnauthorizedOperation.CdnAccountUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNDOMAINUNAUTHORIZED = "UnauthorizedOperation.CdnDomainUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNHOSTUNAUTHORIZED = "UnauthorizedOperation.CdnHostUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNNODOMAINUNAUTHORIZED = "UnauthorizedOperation.CdnNoDomainUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNPROJECTUNAUTHORIZED = "UnauthorizedOperation.CdnProjectUnauthorized"
//  UNAUTHORIZEDOPERATION_DOMAINNOPERMISSION = "UnauthorizedOperation.DomainNoPermission"
//  UNAUTHORIZEDOPERATION_DOMAINSNOPERMISSION = "UnauthorizedOperation.DomainsNoPermission"
//  UNAUTHORIZEDOPERATION_ECDNACCOUNTUNAUTHORIZED = "UnauthorizedOperation.EcdnAccountUnauthorized"
//  UNAUTHORIZEDOPERATION_ECDNCAMUNAUTHORIZED = "UnauthorizedOperation.EcdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_ECDNDOMAINUNAUTHORIZED = "UnauthorizedOperation.EcdnDomainUnauthorized"
//  UNAUTHORIZEDOPERATION_ECDNHOSTUNAUTHORIZED = "UnauthorizedOperation.EcdnHostUnauthorized"
//  UNAUTHORIZEDOPERATION_ECDNNODOMAINUNAUTHORIZED = "UnauthorizedOperation.EcdnNoDomainUnauthorized"
//  UNAUTHORIZEDOPERATION_ECDNPROJECTUNAUTHORIZED = "UnauthorizedOperation.EcdnProjectUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNAUTHORIZEDOPERATION_PROJECTNOPERMISSION = "UnauthorizedOperation.ProjectNoPermission"
//  UNAUTHORIZEDOPERATION_PROJECTSNOPERMISSION = "UnauthorizedOperation.ProjectsNoPermission"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) DescribeEcdnStatistics(request *DescribeEcdnStatisticsRequest) (response *DescribeEcdnStatisticsResponse, err error) {
    return c.DescribeEcdnStatisticsWithContext(context.Background(), request)
}

// DescribeEcdnStatistics
// DescribeEcdnStatistics用于查询 ECDN 实时访问监控数据，支持以下指标查询：
//
// 
//
// + 流量（单位为 byte）
//
// + 带宽（单位为 bps）
//
// + 请求数（单位为 次）
//
// + 状态码 2xx 汇总及各 2 开头状态码明细（单位为 个）
//
// + 状态码 3xx 汇总及各 3 开头状态码明细（单位为 个）
//
// + 状态码 4xx 汇总及各 4 开头状态码明细（单位为 个）
//
// + 状态码 5xx 汇总及各 5 开头状态码明细（单位为 个）
//
// 可能返回的错误码:
//  INTERNALERROR_ECDNCONFIGERROR = "InternalError.EcdnConfigError"
//  INTERNALERROR_ECDNDBERROR = "InternalError.EcdnDbError"
//  INTERNALERROR_ECDNSYSTEMERROR = "InternalError.EcdnSystemError"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INVALIDPARAMETER_ECDNINVALIDPARAMINTERVAL = "InvalidParameter.EcdnInvalidParamInterval"
//  INVALIDPARAMETER_ECDNPARAMERROR = "InvalidParameter.EcdnParamError"
//  INVALIDPARAMETER_ECDNSTATINVALIDDATE = "InvalidParameter.EcdnStatInvalidDate"
//  INVALIDPARAMETER_ECDNSTATINVALIDMETRIC = "InvalidParameter.EcdnStatInvalidMetric"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  LIMITEXCEEDED_ECDNDOMAINOPTOOOFTEN = "LimitExceeded.EcdnDomainOpTooOften"
//  RESOURCENOTFOUND_ECDNDOMAINNOTEXISTS = "ResourceNotFound.EcdnDomainNotExists"
//  RESOURCENOTFOUND_ECDNHOSTNOTEXISTS = "ResourceNotFound.EcdnHostNotExists"
//  RESOURCENOTFOUND_ECDNPROJECTNOTEXISTS = "ResourceNotFound.EcdnProjectNotExists"
//  RESOURCENOTFOUND_ECDNUSERNOTEXISTS = "ResourceNotFound.EcdnUserNotExists"
//  UNAUTHORIZEDOPERATION_CDNACCOUNTUNAUTHORIZED = "UnauthorizedOperation.CdnAccountUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNDOMAINUNAUTHORIZED = "UnauthorizedOperation.CdnDomainUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNHOSTUNAUTHORIZED = "UnauthorizedOperation.CdnHostUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNNODOMAINUNAUTHORIZED = "UnauthorizedOperation.CdnNoDomainUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNPROJECTUNAUTHORIZED = "UnauthorizedOperation.CdnProjectUnauthorized"
//  UNAUTHORIZEDOPERATION_DOMAINNOPERMISSION = "UnauthorizedOperation.DomainNoPermission"
//  UNAUTHORIZEDOPERATION_DOMAINSNOPERMISSION = "UnauthorizedOperation.DomainsNoPermission"
//  UNAUTHORIZEDOPERATION_ECDNACCOUNTUNAUTHORIZED = "UnauthorizedOperation.EcdnAccountUnauthorized"
//  UNAUTHORIZEDOPERATION_ECDNCAMUNAUTHORIZED = "UnauthorizedOperation.EcdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_ECDNDOMAINUNAUTHORIZED = "UnauthorizedOperation.EcdnDomainUnauthorized"
//  UNAUTHORIZEDOPERATION_ECDNHOSTUNAUTHORIZED = "UnauthorizedOperation.EcdnHostUnauthorized"
//  UNAUTHORIZEDOPERATION_ECDNNODOMAINUNAUTHORIZED = "UnauthorizedOperation.EcdnNoDomainUnauthorized"
//  UNAUTHORIZEDOPERATION_ECDNPROJECTUNAUTHORIZED = "UnauthorizedOperation.EcdnProjectUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNAUTHORIZEDOPERATION_PROJECTNOPERMISSION = "UnauthorizedOperation.ProjectNoPermission"
//  UNAUTHORIZEDOPERATION_PROJECTSNOPERMISSION = "UnauthorizedOperation.ProjectsNoPermission"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) DescribeEcdnStatisticsWithContext(ctx context.Context, request *DescribeEcdnStatisticsRequest) (response *DescribeEcdnStatisticsResponse, err error) {
    if request == nil {
        request = NewDescribeEcdnStatisticsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeEcdnStatistics require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeEcdnStatisticsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeIpStatusRequest() (request *DescribeIpStatusRequest) {
    request = &DescribeIpStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ecdn", APIVersion, "DescribeIpStatus")
    
    
    return
}

func NewDescribeIpStatusResponse() (response *DescribeIpStatusResponse) {
    response = &DescribeIpStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeIpStatus
// DescribeIpStatus 用于查询域名所在加速平台的所有节点信息, 如果您的源站有白名单设置,可以通过本接口获取ECDN服务的节点IP进行加白, 本接口为内测接口,请联系腾讯云工程师开白。
//
// 
//
// 由于产品服务节点常有更新，对于源站开白的使用场景，请定期调用接口获取最新节点信息，若新增服务节点发布7日后您尚未更新加白导致回源失败等问题，ECDN侧不对此承担责任。
//
// 可能返回的错误码:
//  INTERNALERROR_ECDNCONFIGERROR = "InternalError.EcdnConfigError"
//  INTERNALERROR_ECDNDBERROR = "InternalError.EcdnDbError"
//  INTERNALERROR_ECDNSYSTEMERROR = "InternalError.EcdnSystemError"
//  INVALIDPARAMETER_ECDNPARAMERROR = "InvalidParameter.EcdnParamError"
//  RESOURCENOTFOUND_ECDNDOMAINNOTEXISTS = "ResourceNotFound.EcdnDomainNotExists"
//  RESOURCENOTFOUND_ECDNHOSTNOTEXISTS = "ResourceNotFound.EcdnHostNotExists"
//  RESOURCENOTFOUND_ECDNUSERNOTEXISTS = "ResourceNotFound.EcdnUserNotExists"
//  UNAUTHORIZEDOPERATION_ECDNCAMUNAUTHORIZED = "UnauthorizedOperation.EcdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_ECDNDOMAINUNAUTHORIZED = "UnauthorizedOperation.EcdnDomainUnauthorized"
//  UNAUTHORIZEDOPERATION_ECDNHOSTUNAUTHORIZED = "UnauthorizedOperation.EcdnHostUnauthorized"
//  UNAUTHORIZEDOPERATION_ECDNUSERNOWHITELIST = "UnauthorizedOperation.EcdnUserNoWhitelist"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
func (c *Client) DescribeIpStatus(request *DescribeIpStatusRequest) (response *DescribeIpStatusResponse, err error) {
    return c.DescribeIpStatusWithContext(context.Background(), request)
}

// DescribeIpStatus
// DescribeIpStatus 用于查询域名所在加速平台的所有节点信息, 如果您的源站有白名单设置,可以通过本接口获取ECDN服务的节点IP进行加白, 本接口为内测接口,请联系腾讯云工程师开白。
//
// 
//
// 由于产品服务节点常有更新，对于源站开白的使用场景，请定期调用接口获取最新节点信息，若新增服务节点发布7日后您尚未更新加白导致回源失败等问题，ECDN侧不对此承担责任。
//
// 可能返回的错误码:
//  INTERNALERROR_ECDNCONFIGERROR = "InternalError.EcdnConfigError"
//  INTERNALERROR_ECDNDBERROR = "InternalError.EcdnDbError"
//  INTERNALERROR_ECDNSYSTEMERROR = "InternalError.EcdnSystemError"
//  INVALIDPARAMETER_ECDNPARAMERROR = "InvalidParameter.EcdnParamError"
//  RESOURCENOTFOUND_ECDNDOMAINNOTEXISTS = "ResourceNotFound.EcdnDomainNotExists"
//  RESOURCENOTFOUND_ECDNHOSTNOTEXISTS = "ResourceNotFound.EcdnHostNotExists"
//  RESOURCENOTFOUND_ECDNUSERNOTEXISTS = "ResourceNotFound.EcdnUserNotExists"
//  UNAUTHORIZEDOPERATION_ECDNCAMUNAUTHORIZED = "UnauthorizedOperation.EcdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_ECDNDOMAINUNAUTHORIZED = "UnauthorizedOperation.EcdnDomainUnauthorized"
//  UNAUTHORIZEDOPERATION_ECDNHOSTUNAUTHORIZED = "UnauthorizedOperation.EcdnHostUnauthorized"
//  UNAUTHORIZEDOPERATION_ECDNUSERNOWHITELIST = "UnauthorizedOperation.EcdnUserNoWhitelist"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
func (c *Client) DescribeIpStatusWithContext(ctx context.Context, request *DescribeIpStatusRequest) (response *DescribeIpStatusResponse, err error) {
    if request == nil {
        request = NewDescribeIpStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeIpStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeIpStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePurgeTasksRequest() (request *DescribePurgeTasksRequest) {
    request = &DescribePurgeTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ecdn", APIVersion, "DescribePurgeTasks")
    
    
    return
}

func NewDescribePurgeTasksResponse() (response *DescribePurgeTasksResponse) {
    response = &DescribePurgeTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribePurgeTasks
// ECDN即将下线，如需要动态加速请使用EdgeOne
//
// 
//
// DescribePurgeTasks 用于查询刷新任务提交历史记录及执行进度。
//
// 
//
// >?  若您的业务已迁移至 CDN 控制台，请参考<a href="https://cloud.tencent.com/document/api/228/37873"> CDN 接口文档</a>，使用  CDN 相关API 进行操作。
//
// 可能返回的错误码:
//  INTERNALERROR_ECDNSYSTEMERROR = "InternalError.EcdnSystemError"
//  INVALIDPARAMETER_ECDNINTERFACEERROR = "InvalidParameter.EcdnInterfaceError"
//  INVALIDPARAMETER_ECDNPARAMERROR = "InvalidParameter.EcdnParamError"
//  RESOURCENOTFOUND_ECDNDOMAINNOTEXISTS = "ResourceNotFound.EcdnDomainNotExists"
//  RESOURCENOTFOUND_ECDNUSERNOTEXISTS = "ResourceNotFound.EcdnUserNotExists"
//  UNAUTHORIZEDOPERATION_ECDNCAMUNAUTHORIZED = "UnauthorizedOperation.EcdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_ECDNMIGRATEDCDN = "UnauthorizedOperation.EcdnMigratedCdn"
//  UNAUTHORIZEDOPERATION_ECDNUSERISSUSPENDED = "UnauthorizedOperation.EcdnUserIsSuspended"
func (c *Client) DescribePurgeTasks(request *DescribePurgeTasksRequest) (response *DescribePurgeTasksResponse, err error) {
    return c.DescribePurgeTasksWithContext(context.Background(), request)
}

// DescribePurgeTasks
// ECDN即将下线，如需要动态加速请使用EdgeOne
//
// 
//
// DescribePurgeTasks 用于查询刷新任务提交历史记录及执行进度。
//
// 
//
// >?  若您的业务已迁移至 CDN 控制台，请参考<a href="https://cloud.tencent.com/document/api/228/37873"> CDN 接口文档</a>，使用  CDN 相关API 进行操作。
//
// 可能返回的错误码:
//  INTERNALERROR_ECDNSYSTEMERROR = "InternalError.EcdnSystemError"
//  INVALIDPARAMETER_ECDNINTERFACEERROR = "InvalidParameter.EcdnInterfaceError"
//  INVALIDPARAMETER_ECDNPARAMERROR = "InvalidParameter.EcdnParamError"
//  RESOURCENOTFOUND_ECDNDOMAINNOTEXISTS = "ResourceNotFound.EcdnDomainNotExists"
//  RESOURCENOTFOUND_ECDNUSERNOTEXISTS = "ResourceNotFound.EcdnUserNotExists"
//  UNAUTHORIZEDOPERATION_ECDNCAMUNAUTHORIZED = "UnauthorizedOperation.EcdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_ECDNMIGRATEDCDN = "UnauthorizedOperation.EcdnMigratedCdn"
//  UNAUTHORIZEDOPERATION_ECDNUSERISSUSPENDED = "UnauthorizedOperation.EcdnUserIsSuspended"
func (c *Client) DescribePurgeTasksWithContext(ctx context.Context, request *DescribePurgeTasksRequest) (response *DescribePurgeTasksResponse, err error) {
    if request == nil {
        request = NewDescribePurgeTasksRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePurgeTasks require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePurgeTasksResponse()
    err = c.Send(request, response)
    return
}

func NewPurgeUrlsCacheRequest() (request *PurgeUrlsCacheRequest) {
    request = &PurgeUrlsCacheRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ecdn", APIVersion, "PurgeUrlsCache")
    
    
    return
}

func NewPurgeUrlsCacheResponse() (response *PurgeUrlsCacheResponse) {
    response = &PurgeUrlsCacheResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// PurgeUrlsCache
// ECDN即将下线，如需要动态加速请使用EdgeOne
//
// 
//
// PurgeUrlsCache 用于批量刷新Url，一次提交将返回一个刷新任务id。
//
// 
//
// >?  若您的业务已迁移至 CDN 控制台，请参考<a href="https://cloud.tencent.com/document/api/228/37870"> CDN 接口文档</a>，使用  CDN 相关API 进行操作。
//
// 可能返回的错误码:
//  FAILEDOPERATION_ECDNCONFIGERROR = "FailedOperation.EcdnConfigError"
//  INTERNALERROR_ECDNCONFIGERROR = "InternalError.EcdnConfigError"
//  INTERNALERROR_ECDNDBERROR = "InternalError.EcdnDbError"
//  INTERNALERROR_ECDNSYSTEMERROR = "InternalError.EcdnSystemError"
//  INVALIDPARAMETER_ECDNDOMAININVALIDSTATUS = "InvalidParameter.EcdnDomainInvalidStatus"
//  INVALIDPARAMETER_ECDNINTERFACEERROR = "InvalidParameter.EcdnInterfaceError"
//  INVALIDPARAMETER_ECDNPARAMERROR = "InvalidParameter.EcdnParamError"
//  INVALIDPARAMETER_ECDNPURGEWILDCARDNOTALLOWED = "InvalidParameter.EcdnPurgeWildcardNotAllowed"
//  INVALIDPARAMETER_ECDNURLEXCEEDLENGTH = "InvalidParameter.EcdnUrlExceedLength"
//  LIMITEXCEEDED_ECDNPURGEURLEXCEEDBATCHLIMIT = "LimitExceeded.EcdnPurgeUrlExceedBatchLimit"
//  LIMITEXCEEDED_ECDNPURGEURLEXCEEDDAYLIMIT = "LimitExceeded.EcdnPurgeUrlExceedDayLimit"
//  RESOURCENOTFOUND_ECDNDOMAINNOTEXISTS = "ResourceNotFound.EcdnDomainNotExists"
//  RESOURCENOTFOUND_ECDNUSERNOTEXISTS = "ResourceNotFound.EcdnUserNotExists"
//  UNAUTHORIZEDOPERATION_ECDNCAMUNAUTHORIZED = "UnauthorizedOperation.EcdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_ECDNDOMAINUNAUTHORIZED = "UnauthorizedOperation.EcdnDomainUnauthorized"
//  UNAUTHORIZEDOPERATION_ECDNMIGRATEDCDN = "UnauthorizedOperation.EcdnMigratedCdn"
//  UNAUTHORIZEDOPERATION_ECDNUSERISSUSPENDED = "UnauthorizedOperation.EcdnUserIsSuspended"
func (c *Client) PurgeUrlsCache(request *PurgeUrlsCacheRequest) (response *PurgeUrlsCacheResponse, err error) {
    return c.PurgeUrlsCacheWithContext(context.Background(), request)
}

// PurgeUrlsCache
// ECDN即将下线，如需要动态加速请使用EdgeOne
//
// 
//
// PurgeUrlsCache 用于批量刷新Url，一次提交将返回一个刷新任务id。
//
// 
//
// >?  若您的业务已迁移至 CDN 控制台，请参考<a href="https://cloud.tencent.com/document/api/228/37870"> CDN 接口文档</a>，使用  CDN 相关API 进行操作。
//
// 可能返回的错误码:
//  FAILEDOPERATION_ECDNCONFIGERROR = "FailedOperation.EcdnConfigError"
//  INTERNALERROR_ECDNCONFIGERROR = "InternalError.EcdnConfigError"
//  INTERNALERROR_ECDNDBERROR = "InternalError.EcdnDbError"
//  INTERNALERROR_ECDNSYSTEMERROR = "InternalError.EcdnSystemError"
//  INVALIDPARAMETER_ECDNDOMAININVALIDSTATUS = "InvalidParameter.EcdnDomainInvalidStatus"
//  INVALIDPARAMETER_ECDNINTERFACEERROR = "InvalidParameter.EcdnInterfaceError"
//  INVALIDPARAMETER_ECDNPARAMERROR = "InvalidParameter.EcdnParamError"
//  INVALIDPARAMETER_ECDNPURGEWILDCARDNOTALLOWED = "InvalidParameter.EcdnPurgeWildcardNotAllowed"
//  INVALIDPARAMETER_ECDNURLEXCEEDLENGTH = "InvalidParameter.EcdnUrlExceedLength"
//  LIMITEXCEEDED_ECDNPURGEURLEXCEEDBATCHLIMIT = "LimitExceeded.EcdnPurgeUrlExceedBatchLimit"
//  LIMITEXCEEDED_ECDNPURGEURLEXCEEDDAYLIMIT = "LimitExceeded.EcdnPurgeUrlExceedDayLimit"
//  RESOURCENOTFOUND_ECDNDOMAINNOTEXISTS = "ResourceNotFound.EcdnDomainNotExists"
//  RESOURCENOTFOUND_ECDNUSERNOTEXISTS = "ResourceNotFound.EcdnUserNotExists"
//  UNAUTHORIZEDOPERATION_ECDNCAMUNAUTHORIZED = "UnauthorizedOperation.EcdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_ECDNDOMAINUNAUTHORIZED = "UnauthorizedOperation.EcdnDomainUnauthorized"
//  UNAUTHORIZEDOPERATION_ECDNMIGRATEDCDN = "UnauthorizedOperation.EcdnMigratedCdn"
//  UNAUTHORIZEDOPERATION_ECDNUSERISSUSPENDED = "UnauthorizedOperation.EcdnUserIsSuspended"
func (c *Client) PurgeUrlsCacheWithContext(ctx context.Context, request *PurgeUrlsCacheRequest) (response *PurgeUrlsCacheResponse, err error) {
    if request == nil {
        request = NewPurgeUrlsCacheRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("PurgeUrlsCache require credential")
    }

    request.SetContext(ctx)
    
    response = NewPurgeUrlsCacheResponse()
    err = c.Send(request, response)
    return
}
