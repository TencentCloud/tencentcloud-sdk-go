// Copyright (c) 2017-2025 Tencent. All Rights Reserved.
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

package v20241024

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2024-10-24"

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


func NewAcceptTccVpcEndPointConnectRequest() (request *AcceptTccVpcEndPointConnectRequest) {
    request = &AcceptTccVpcEndPointConnectRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tccatalog", APIVersion, "AcceptTccVpcEndPointConnect")
    
    
    return
}

func NewAcceptTccVpcEndPointConnectResponse() (response *AcceptTccVpcEndPointConnectResponse) {
    response = &AcceptTccVpcEndPointConnectResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AcceptTccVpcEndPointConnect
// 接受终端节点连接
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION_USERNOTALLOWOPERATION = "UnauthorizedOperation.UserNotAllowOperation"
func (c *Client) AcceptTccVpcEndPointConnect(request *AcceptTccVpcEndPointConnectRequest) (response *AcceptTccVpcEndPointConnectResponse, err error) {
    return c.AcceptTccVpcEndPointConnectWithContext(context.Background(), request)
}

// AcceptTccVpcEndPointConnect
// 接受终端节点连接
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION_USERNOTALLOWOPERATION = "UnauthorizedOperation.UserNotAllowOperation"
func (c *Client) AcceptTccVpcEndPointConnectWithContext(ctx context.Context, request *AcceptTccVpcEndPointConnectRequest) (response *AcceptTccVpcEndPointConnectResponse, err error) {
    if request == nil {
        request = NewAcceptTccVpcEndPointConnectRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tccatalog", APIVersion, "AcceptTccVpcEndPointConnect")
    
    if c.GetCredential() == nil {
        return nil, errors.New("AcceptTccVpcEndPointConnect require credential")
    }

    request.SetContext(ctx)
    
    response = NewAcceptTccVpcEndPointConnectResponse()
    err = c.Send(request, response)
    return
}

func NewBindTccVpcEndPointServiceWhiteListRequest() (request *BindTccVpcEndPointServiceWhiteListRequest) {
    request = &BindTccVpcEndPointServiceWhiteListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tccatalog", APIVersion, "BindTccVpcEndPointServiceWhiteList")
    
    
    return
}

func NewBindTccVpcEndPointServiceWhiteListResponse() (response *BindTccVpcEndPointServiceWhiteListResponse) {
    response = &BindTccVpcEndPointServiceWhiteListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// BindTccVpcEndPointServiceWhiteList
// 绑定终端节点服务白名单用户
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  RESOURCEUNAVAILABLE_CLUSTERUNAVAILABLE = "ResourceUnavailable.ClusterUnavailable"
//  UNAUTHORIZEDOPERATION_USERNOTALLOWOPERATION = "UnauthorizedOperation.UserNotAllowOperation"
func (c *Client) BindTccVpcEndPointServiceWhiteList(request *BindTccVpcEndPointServiceWhiteListRequest) (response *BindTccVpcEndPointServiceWhiteListResponse, err error) {
    return c.BindTccVpcEndPointServiceWhiteListWithContext(context.Background(), request)
}

// BindTccVpcEndPointServiceWhiteList
// 绑定终端节点服务白名单用户
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  RESOURCEUNAVAILABLE_CLUSTERUNAVAILABLE = "ResourceUnavailable.ClusterUnavailable"
//  UNAUTHORIZEDOPERATION_USERNOTALLOWOPERATION = "UnauthorizedOperation.UserNotAllowOperation"
func (c *Client) BindTccVpcEndPointServiceWhiteListWithContext(ctx context.Context, request *BindTccVpcEndPointServiceWhiteListRequest) (response *BindTccVpcEndPointServiceWhiteListResponse, err error) {
    if request == nil {
        request = NewBindTccVpcEndPointServiceWhiteListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tccatalog", APIVersion, "BindTccVpcEndPointServiceWhiteList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("BindTccVpcEndPointServiceWhiteList require credential")
    }

    request.SetContext(ctx)
    
    response = NewBindTccVpcEndPointServiceWhiteListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTccCatalogRequest() (request *DescribeTccCatalogRequest) {
    request = &DescribeTccCatalogRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tccatalog", APIVersion, "DescribeTccCatalog")
    
    
    return
}

func NewDescribeTccCatalogResponse() (response *DescribeTccCatalogResponse) {
    response = &DescribeTccCatalogResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTccCatalog
// 获取Tcc数据目录详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_INVALIDTIMEFORMAT = "InvalidParameter.InvalidTimeFormat"
//  INVALIDPARAMETER_INVALIDTIMEPARAMETER = "InvalidParameter.InvalidTimeParameter"
func (c *Client) DescribeTccCatalog(request *DescribeTccCatalogRequest) (response *DescribeTccCatalogResponse, err error) {
    return c.DescribeTccCatalogWithContext(context.Background(), request)
}

// DescribeTccCatalog
// 获取Tcc数据目录详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_INVALIDTIMEFORMAT = "InvalidParameter.InvalidTimeFormat"
//  INVALIDPARAMETER_INVALIDTIMEPARAMETER = "InvalidParameter.InvalidTimeParameter"
func (c *Client) DescribeTccCatalogWithContext(ctx context.Context, request *DescribeTccCatalogRequest) (response *DescribeTccCatalogResponse, err error) {
    if request == nil {
        request = NewDescribeTccCatalogRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tccatalog", APIVersion, "DescribeTccCatalog")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTccCatalog require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTccCatalogResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTccCatalogsRequest() (request *DescribeTccCatalogsRequest) {
    request = &DescribeTccCatalogsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tccatalog", APIVersion, "DescribeTccCatalogs")
    
    
    return
}

func NewDescribeTccCatalogsResponse() (response *DescribeTccCatalogsResponse) {
    response = &DescribeTccCatalogsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTccCatalogs
// 获取Tcc数据目录列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_INVALIDTIMEFORMAT = "InvalidParameter.InvalidTimeFormat"
//  INVALIDPARAMETER_INVALIDTIMEPARAMETER = "InvalidParameter.InvalidTimeParameter"
func (c *Client) DescribeTccCatalogs(request *DescribeTccCatalogsRequest) (response *DescribeTccCatalogsResponse, err error) {
    return c.DescribeTccCatalogsWithContext(context.Background(), request)
}

// DescribeTccCatalogs
// 获取Tcc数据目录列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_INVALIDTIMEFORMAT = "InvalidParameter.InvalidTimeFormat"
//  INVALIDPARAMETER_INVALIDTIMEPARAMETER = "InvalidParameter.InvalidTimeParameter"
func (c *Client) DescribeTccCatalogsWithContext(ctx context.Context, request *DescribeTccCatalogsRequest) (response *DescribeTccCatalogsResponse, err error) {
    if request == nil {
        request = NewDescribeTccCatalogsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tccatalog", APIVersion, "DescribeTccCatalogs")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTccCatalogs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTccCatalogsResponse()
    err = c.Send(request, response)
    return
}
