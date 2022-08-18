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

package v20181115

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2018-11-15"

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


func NewDescribeDomainInfoRequest() (request *DescribeDomainInfoRequest) {
    request = &DescribeDomainInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tics", APIVersion, "DescribeDomainInfo")
    
    
    return
}

func NewDescribeDomainInfoResponse() (response *DescribeDomainInfoResponse) {
    response = &DescribeDomainInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDomainInfo
// 提供域名相关的基础信息以及与攻击事件（团伙、家族）、恶意文件等相关联信息。
//
// 可能返回的错误码:
//  INTERNALERROR_CACHEERR = "InternalError.CacheErr"
//  INTERNALERROR_LOCALERR = "InternalError.LocalErr"
//  INTERNALERROR_TIMEOUT = "InternalError.Timeout"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
func (c *Client) DescribeDomainInfo(request *DescribeDomainInfoRequest) (response *DescribeDomainInfoResponse, err error) {
    return c.DescribeDomainInfoWithContext(context.Background(), request)
}

// DescribeDomainInfo
// 提供域名相关的基础信息以及与攻击事件（团伙、家族）、恶意文件等相关联信息。
//
// 可能返回的错误码:
//  INTERNALERROR_CACHEERR = "InternalError.CacheErr"
//  INTERNALERROR_LOCALERR = "InternalError.LocalErr"
//  INTERNALERROR_TIMEOUT = "InternalError.Timeout"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
func (c *Client) DescribeDomainInfoWithContext(ctx context.Context, request *DescribeDomainInfoRequest) (response *DescribeDomainInfoResponse, err error) {
    if request == nil {
        request = NewDescribeDomainInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDomainInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDomainInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeFileInfoRequest() (request *DescribeFileInfoRequest) {
    request = &DescribeFileInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tics", APIVersion, "DescribeFileInfo")
    
    
    return
}

func NewDescribeFileInfoResponse() (response *DescribeFileInfoResponse) {
    response = &DescribeFileInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeFileInfo
// 提供文件相关的基础信息以及与攻击事件（团伙、家族）、恶意文件等相关联信息。
//
// 可能返回的错误码:
//  INTERNALERROR_CACHEERR = "InternalError.CacheErr"
//  INTERNALERROR_LOCALERR = "InternalError.LocalErr"
//  INTERNALERROR_TIMEOUT = "InternalError.Timeout"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
func (c *Client) DescribeFileInfo(request *DescribeFileInfoRequest) (response *DescribeFileInfoResponse, err error) {
    return c.DescribeFileInfoWithContext(context.Background(), request)
}

// DescribeFileInfo
// 提供文件相关的基础信息以及与攻击事件（团伙、家族）、恶意文件等相关联信息。
//
// 可能返回的错误码:
//  INTERNALERROR_CACHEERR = "InternalError.CacheErr"
//  INTERNALERROR_LOCALERR = "InternalError.LocalErr"
//  INTERNALERROR_TIMEOUT = "InternalError.Timeout"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
func (c *Client) DescribeFileInfoWithContext(ctx context.Context, request *DescribeFileInfoRequest) (response *DescribeFileInfoResponse, err error) {
    if request == nil {
        request = NewDescribeFileInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeFileInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeFileInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeIpInfoRequest() (request *DescribeIpInfoRequest) {
    request = &DescribeIpInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tics", APIVersion, "DescribeIpInfo")
    
    
    return
}

func NewDescribeIpInfoResponse() (response *DescribeIpInfoResponse) {
    response = &DescribeIpInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeIpInfo
// 提供IP相关的基础信息以及与攻击事件（团伙、家族）、恶意文件等相关联信息。
//
// 可能返回的错误码:
//  INTERNALERROR_CACHEERR = "InternalError.CacheErr"
//  INTERNALERROR_LOCALERR = "InternalError.LocalErr"
//  INTERNALERROR_TIMEOUT = "InternalError.Timeout"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
func (c *Client) DescribeIpInfo(request *DescribeIpInfoRequest) (response *DescribeIpInfoResponse, err error) {
    return c.DescribeIpInfoWithContext(context.Background(), request)
}

// DescribeIpInfo
// 提供IP相关的基础信息以及与攻击事件（团伙、家族）、恶意文件等相关联信息。
//
// 可能返回的错误码:
//  INTERNALERROR_CACHEERR = "InternalError.CacheErr"
//  INTERNALERROR_LOCALERR = "InternalError.LocalErr"
//  INTERNALERROR_TIMEOUT = "InternalError.Timeout"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
func (c *Client) DescribeIpInfoWithContext(ctx context.Context, request *DescribeIpInfoRequest) (response *DescribeIpInfoResponse, err error) {
    if request == nil {
        request = NewDescribeIpInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeIpInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeIpInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeThreatInfoRequest() (request *DescribeThreatInfoRequest) {
    request = &DescribeThreatInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tics", APIVersion, "DescribeThreatInfo")
    
    
    return
}

func NewDescribeThreatInfoResponse() (response *DescribeThreatInfoResponse) {
    response = &DescribeThreatInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeThreatInfo
// 提供IP和域名相关威胁情报信息查询，这些信息可以辅助检测失陷主机、帮助SIEM/SOC等系统做研判决策、帮助运营团队对设备报警的编排处理。
//
// 可能返回的错误码:
//  INTERNALERROR_CACHEERR = "InternalError.CacheErr"
//  INTERNALERROR_LOCALERR = "InternalError.LocalErr"
//  INTERNALERROR_TIMEOUT = "InternalError.Timeout"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
func (c *Client) DescribeThreatInfo(request *DescribeThreatInfoRequest) (response *DescribeThreatInfoResponse, err error) {
    return c.DescribeThreatInfoWithContext(context.Background(), request)
}

// DescribeThreatInfo
// 提供IP和域名相关威胁情报信息查询，这些信息可以辅助检测失陷主机、帮助SIEM/SOC等系统做研判决策、帮助运营团队对设备报警的编排处理。
//
// 可能返回的错误码:
//  INTERNALERROR_CACHEERR = "InternalError.CacheErr"
//  INTERNALERROR_LOCALERR = "InternalError.LocalErr"
//  INTERNALERROR_TIMEOUT = "InternalError.Timeout"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
func (c *Client) DescribeThreatInfoWithContext(ctx context.Context, request *DescribeThreatInfoRequest) (response *DescribeThreatInfoResponse, err error) {
    if request == nil {
        request = NewDescribeThreatInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeThreatInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeThreatInfoResponse()
    err = c.Send(request, response)
    return
}
