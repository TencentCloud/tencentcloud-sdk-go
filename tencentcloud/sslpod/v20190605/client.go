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

package v20190605

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2019-06-05"

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


func NewCreateDomainRequest() (request *CreateDomainRequest) {
    request = &CreateDomainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sslpod", APIVersion, "CreateDomain")
    
    
    return
}

func NewCreateDomainResponse() (response *CreateDomainResponse) {
    response = &CreateDomainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateDomain
// 通过域名端口添加监控
//
// 可能返回的错误码:
//  FAILEDOPERATION_REPETITIONADD = "FailedOperation.RepetitionAdd"
//  FAILEDOPERATION_RESOLVEDOMAINFAILED = "FailedOperation.ResolveDomainFailed"
//  INVALIDPARAMETER_INVALIDDOMAIN = "InvalidParameter.InvalidDomain"
//  INVALIDPARAMETER_INVALIDIP = "InvalidParameter.InvalidIP"
//  INVALIDPARAMETER_INVALIDPORT = "InvalidParameter.InvalidPort"
//  INVALIDPARAMETER_INVALIDSERVERTYPE = "InvalidParameter.InvalidServerType"
//  INVALIDPARAMETER_INVALIDTAGNAME = "InvalidParameter.InvalidTagName"
//  INVALIDPARAMETER_TOOMANYTAG = "InvalidParameter.TooManyTag"
//  INVALIDPARAMETERVALUE_INVALIDNOTICETYPE = "InvalidParameterValue.InvalidNoticeType"
//  INVALIDPARAMETERVALUE_INVALIDSEARCHTYPE = "InvalidParameterValue.InvalidSearchType"
//  LIMITEXCEEDED_ADDEXCEEDED = "LimitExceeded.AddExceeded"
//  LIMITEXCEEDED_MONITOREXCEEDED = "LimitExceeded.MonitorExceeded"
//  RESOURCENOTFOUND_PRODUCT = "ResourceNotFound.Product"
func (c *Client) CreateDomain(request *CreateDomainRequest) (response *CreateDomainResponse, err error) {
    return c.CreateDomainWithContext(context.Background(), request)
}

// CreateDomain
// 通过域名端口添加监控
//
// 可能返回的错误码:
//  FAILEDOPERATION_REPETITIONADD = "FailedOperation.RepetitionAdd"
//  FAILEDOPERATION_RESOLVEDOMAINFAILED = "FailedOperation.ResolveDomainFailed"
//  INVALIDPARAMETER_INVALIDDOMAIN = "InvalidParameter.InvalidDomain"
//  INVALIDPARAMETER_INVALIDIP = "InvalidParameter.InvalidIP"
//  INVALIDPARAMETER_INVALIDPORT = "InvalidParameter.InvalidPort"
//  INVALIDPARAMETER_INVALIDSERVERTYPE = "InvalidParameter.InvalidServerType"
//  INVALIDPARAMETER_INVALIDTAGNAME = "InvalidParameter.InvalidTagName"
//  INVALIDPARAMETER_TOOMANYTAG = "InvalidParameter.TooManyTag"
//  INVALIDPARAMETERVALUE_INVALIDNOTICETYPE = "InvalidParameterValue.InvalidNoticeType"
//  INVALIDPARAMETERVALUE_INVALIDSEARCHTYPE = "InvalidParameterValue.InvalidSearchType"
//  LIMITEXCEEDED_ADDEXCEEDED = "LimitExceeded.AddExceeded"
//  LIMITEXCEEDED_MONITOREXCEEDED = "LimitExceeded.MonitorExceeded"
//  RESOURCENOTFOUND_PRODUCT = "ResourceNotFound.Product"
func (c *Client) CreateDomainWithContext(ctx context.Context, request *CreateDomainRequest) (response *CreateDomainResponse, err error) {
    if request == nil {
        request = NewCreateDomainRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateDomain require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateDomainResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteDomainRequest() (request *DeleteDomainRequest) {
    request = &DeleteDomainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sslpod", APIVersion, "DeleteDomain")
    
    
    return
}

func NewDeleteDomainResponse() (response *DeleteDomainResponse) {
    response = &DeleteDomainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteDomain
// 通过域名ID删除监控的域名
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DeleteDomain(request *DeleteDomainRequest) (response *DeleteDomainResponse, err error) {
    return c.DeleteDomainWithContext(context.Background(), request)
}

// DeleteDomain
// 通过域名ID删除监控的域名
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DeleteDomainWithContext(ctx context.Context, request *DeleteDomainRequest) (response *DeleteDomainResponse, err error) {
    if request == nil {
        request = NewDeleteDomainRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteDomain require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteDomainResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDashboardRequest() (request *DescribeDashboardRequest) {
    request = &DescribeDashboardRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sslpod", APIVersion, "DescribeDashboard")
    
    
    return
}

func NewDescribeDashboardResponse() (response *DescribeDashboardResponse) {
    response = &DescribeDashboardResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDashboard
// 获取仪表盘数据
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeDashboard(request *DescribeDashboardRequest) (response *DescribeDashboardResponse, err error) {
    return c.DescribeDashboardWithContext(context.Background(), request)
}

// DescribeDashboard
// 获取仪表盘数据
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeDashboardWithContext(ctx context.Context, request *DescribeDashboardRequest) (response *DescribeDashboardResponse, err error) {
    if request == nil {
        request = NewDescribeDashboardRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDashboard require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDashboardResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDomainCertsRequest() (request *DescribeDomainCertsRequest) {
    request = &DescribeDomainCertsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sslpod", APIVersion, "DescribeDomainCerts")
    
    
    return
}

func NewDescribeDomainCertsResponse() (response *DescribeDomainCertsResponse) {
    response = &DescribeDomainCertsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDomainCerts
// 获取域名关联证书
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeDomainCerts(request *DescribeDomainCertsRequest) (response *DescribeDomainCertsResponse, err error) {
    return c.DescribeDomainCertsWithContext(context.Background(), request)
}

// DescribeDomainCerts
// 获取域名关联证书
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeDomainCertsWithContext(ctx context.Context, request *DescribeDomainCertsRequest) (response *DescribeDomainCertsResponse, err error) {
    if request == nil {
        request = NewDescribeDomainCertsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDomainCerts require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDomainCertsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDomainTagsRequest() (request *DescribeDomainTagsRequest) {
    request = &DescribeDomainTagsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sslpod", APIVersion, "DescribeDomainTags")
    
    
    return
}

func NewDescribeDomainTagsResponse() (response *DescribeDomainTagsResponse) {
    response = &DescribeDomainTagsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDomainTags
// 获取账号下所有tag
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeDomainTags(request *DescribeDomainTagsRequest) (response *DescribeDomainTagsResponse, err error) {
    return c.DescribeDomainTagsWithContext(context.Background(), request)
}

// DescribeDomainTags
// 获取账号下所有tag
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeDomainTagsWithContext(ctx context.Context, request *DescribeDomainTagsRequest) (response *DescribeDomainTagsResponse, err error) {
    if request == nil {
        request = NewDescribeDomainTagsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDomainTags require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDomainTagsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDomainsRequest() (request *DescribeDomainsRequest) {
    request = &DescribeDomainsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sslpod", APIVersion, "DescribeDomains")
    
    
    return
}

func NewDescribeDomainsResponse() (response *DescribeDomainsResponse) {
    response = &DescribeDomainsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDomains
// 通过searchType搜索已经添加的域名
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_INVALIDDOMAIN = "InvalidParameter.InvalidDomain"
//  INVALIDPARAMETERVALUE_INVALIDSEARCHTYPE = "InvalidParameterValue.InvalidSearchType"
func (c *Client) DescribeDomains(request *DescribeDomainsRequest) (response *DescribeDomainsResponse, err error) {
    return c.DescribeDomainsWithContext(context.Background(), request)
}

// DescribeDomains
// 通过searchType搜索已经添加的域名
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_INVALIDDOMAIN = "InvalidParameter.InvalidDomain"
//  INVALIDPARAMETERVALUE_INVALIDSEARCHTYPE = "InvalidParameterValue.InvalidSearchType"
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

func NewDescribeNoticeInfoRequest() (request *DescribeNoticeInfoRequest) {
    request = &DescribeNoticeInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sslpod", APIVersion, "DescribeNoticeInfo")
    
    
    return
}

func NewDescribeNoticeInfoResponse() (response *DescribeNoticeInfoResponse) {
    response = &DescribeNoticeInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeNoticeInfo
// 获取通知额度信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeNoticeInfo(request *DescribeNoticeInfoRequest) (response *DescribeNoticeInfoResponse, err error) {
    return c.DescribeNoticeInfoWithContext(context.Background(), request)
}

// DescribeNoticeInfo
// 获取通知额度信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeNoticeInfoWithContext(ctx context.Context, request *DescribeNoticeInfoRequest) (response *DescribeNoticeInfoResponse, err error) {
    if request == nil {
        request = NewDescribeNoticeInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeNoticeInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeNoticeInfoResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDomainTagsRequest() (request *ModifyDomainTagsRequest) {
    request = &ModifyDomainTagsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sslpod", APIVersion, "ModifyDomainTags")
    
    
    return
}

func NewModifyDomainTagsResponse() (response *ModifyDomainTagsResponse) {
    response = &ModifyDomainTagsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyDomainTags
// 修改域名tag
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDTAGNAME = "InvalidParameter.InvalidTagName"
//  INVALIDPARAMETER_TOOMANYTAG = "InvalidParameter.TooManyTag"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) ModifyDomainTags(request *ModifyDomainTagsRequest) (response *ModifyDomainTagsResponse, err error) {
    return c.ModifyDomainTagsWithContext(context.Background(), request)
}

// ModifyDomainTags
// 修改域名tag
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDTAGNAME = "InvalidParameter.InvalidTagName"
//  INVALIDPARAMETER_TOOMANYTAG = "InvalidParameter.TooManyTag"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) ModifyDomainTagsWithContext(ctx context.Context, request *ModifyDomainTagsRequest) (response *ModifyDomainTagsResponse, err error) {
    if request == nil {
        request = NewModifyDomainTagsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDomainTags require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDomainTagsResponse()
    err = c.Send(request, response)
    return
}

func NewRefreshDomainRequest() (request *RefreshDomainRequest) {
    request = &RefreshDomainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sslpod", APIVersion, "RefreshDomain")
    
    
    return
}

func NewRefreshDomainResponse() (response *RefreshDomainResponse) {
    response = &RefreshDomainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// RefreshDomain
// 强制重新检测域名
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) RefreshDomain(request *RefreshDomainRequest) (response *RefreshDomainResponse, err error) {
    return c.RefreshDomainWithContext(context.Background(), request)
}

// RefreshDomain
// 强制重新检测域名
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) RefreshDomainWithContext(ctx context.Context, request *RefreshDomainRequest) (response *RefreshDomainResponse, err error) {
    if request == nil {
        request = NewRefreshDomainRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RefreshDomain require credential")
    }

    request.SetContext(ctx)
    
    response = NewRefreshDomainResponse()
    err = c.Send(request, response)
    return
}

func NewResolveDomainRequest() (request *ResolveDomainRequest) {
    request = &ResolveDomainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sslpod", APIVersion, "ResolveDomain")
    
    
    return
}

func NewResolveDomainResponse() (response *ResolveDomainResponse) {
    response = &ResolveDomainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ResolveDomain
// 解析域名获得多个IP地址
//
// 可能返回的错误码:
//  FAILEDOPERATION_RESOLVEDOMAINFAILED = "FailedOperation.ResolveDomainFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDDOMAIN = "InvalidParameter.InvalidDomain"
func (c *Client) ResolveDomain(request *ResolveDomainRequest) (response *ResolveDomainResponse, err error) {
    return c.ResolveDomainWithContext(context.Background(), request)
}

// ResolveDomain
// 解析域名获得多个IP地址
//
// 可能返回的错误码:
//  FAILEDOPERATION_RESOLVEDOMAINFAILED = "FailedOperation.ResolveDomainFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDDOMAIN = "InvalidParameter.InvalidDomain"
func (c *Client) ResolveDomainWithContext(ctx context.Context, request *ResolveDomainRequest) (response *ResolveDomainResponse, err error) {
    if request == nil {
        request = NewResolveDomainRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ResolveDomain require credential")
    }

    request.SetContext(ctx)
    
    response = NewResolveDomainResponse()
    err = c.Send(request, response)
    return
}
