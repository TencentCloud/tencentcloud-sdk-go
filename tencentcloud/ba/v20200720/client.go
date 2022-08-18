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

package v20200720

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2020-07-20"

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


func NewCreateWeappQRUrlRequest() (request *CreateWeappQRUrlRequest) {
    request = &CreateWeappQRUrlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ba", APIVersion, "CreateWeappQRUrl")
    
    
    return
}

func NewCreateWeappQRUrlResponse() (response *CreateWeappQRUrlResponse) {
    response = &CreateWeappQRUrlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateWeappQRUrl
// 创建渠道备案小程序二维码
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) CreateWeappQRUrl(request *CreateWeappQRUrlRequest) (response *CreateWeappQRUrlResponse, err error) {
    return c.CreateWeappQRUrlWithContext(context.Background(), request)
}

// CreateWeappQRUrl
// 创建渠道备案小程序二维码
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) CreateWeappQRUrlWithContext(ctx context.Context, request *CreateWeappQRUrlRequest) (response *CreateWeappQRUrlResponse, err error) {
    if request == nil {
        request = NewCreateWeappQRUrlRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateWeappQRUrl require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateWeappQRUrlResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeGetAuthInfoRequest() (request *DescribeGetAuthInfoRequest) {
    request = &DescribeGetAuthInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ba", APIVersion, "DescribeGetAuthInfo")
    
    
    return
}

func NewDescribeGetAuthInfoResponse() (response *DescribeGetAuthInfoResponse) {
    response = &DescribeGetAuthInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeGetAuthInfo
// 获取实名认证信息
//
// 可能返回的错误码:
//  INTERNALERROR_RESOURCENOTFOUND = "InternalError.ResourceNotFound"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeGetAuthInfo(request *DescribeGetAuthInfoRequest) (response *DescribeGetAuthInfoResponse, err error) {
    return c.DescribeGetAuthInfoWithContext(context.Background(), request)
}

// DescribeGetAuthInfo
// 获取实名认证信息
//
// 可能返回的错误码:
//  INTERNALERROR_RESOURCENOTFOUND = "InternalError.ResourceNotFound"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeGetAuthInfoWithContext(ctx context.Context, request *DescribeGetAuthInfoRequest) (response *DescribeGetAuthInfoResponse, err error) {
    if request == nil {
        request = NewDescribeGetAuthInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeGetAuthInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeGetAuthInfoResponse()
    err = c.Send(request, response)
    return
}

func NewSyncIcpOrderWebInfoRequest() (request *SyncIcpOrderWebInfoRequest) {
    request = &SyncIcpOrderWebInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ba", APIVersion, "SyncIcpOrderWebInfo")
    
    
    return
}

func NewSyncIcpOrderWebInfoResponse() (response *SyncIcpOrderWebInfoResponse) {
    response = &SyncIcpOrderWebInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// SyncIcpOrderWebInfo
// 将备案ICP订单下的一个网站信息 同步给订单下其他网站，需要被同步的网站被检查通过(isCheck:true)；
//
// 只有指定的网站信息字段能被同步
//
// 可能返回的错误码:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) SyncIcpOrderWebInfo(request *SyncIcpOrderWebInfoRequest) (response *SyncIcpOrderWebInfoResponse, err error) {
    return c.SyncIcpOrderWebInfoWithContext(context.Background(), request)
}

// SyncIcpOrderWebInfo
// 将备案ICP订单下的一个网站信息 同步给订单下其他网站，需要被同步的网站被检查通过(isCheck:true)；
//
// 只有指定的网站信息字段能被同步
//
// 可能返回的错误码:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) SyncIcpOrderWebInfoWithContext(ctx context.Context, request *SyncIcpOrderWebInfoRequest) (response *SyncIcpOrderWebInfoResponse, err error) {
    if request == nil {
        request = NewSyncIcpOrderWebInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SyncIcpOrderWebInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewSyncIcpOrderWebInfoResponse()
    err = c.Send(request, response)
    return
}
