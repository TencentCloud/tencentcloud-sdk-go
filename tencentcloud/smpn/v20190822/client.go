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

package v20190822

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2019-08-22"

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


func NewCreateSmpnEpaRequest() (request *CreateSmpnEpaRequest) {
    request = &CreateSmpnEpaRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("smpn", APIVersion, "CreateSmpnEpa")
    
    
    return
}

func NewCreateSmpnEpaResponse() (response *CreateSmpnEpaResponse) {
    response = &CreateSmpnEpaResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateSmpnEpa
// 企业号码认证
//
// 可能返回的错误码:
//  INTERNALERROR_TIMEOUT = "InternalError.Timeout"
//  INVALIDPARAMETER_NAME = "InvalidParameter.Name"
//  INVALIDPARAMETER_PHONENUMBER = "InvalidParameter.PhoneNumber"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) CreateSmpnEpa(request *CreateSmpnEpaRequest) (response *CreateSmpnEpaResponse, err error) {
    return c.CreateSmpnEpaWithContext(context.Background(), request)
}

// CreateSmpnEpa
// 企业号码认证
//
// 可能返回的错误码:
//  INTERNALERROR_TIMEOUT = "InternalError.Timeout"
//  INVALIDPARAMETER_NAME = "InvalidParameter.Name"
//  INVALIDPARAMETER_PHONENUMBER = "InvalidParameter.PhoneNumber"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) CreateSmpnEpaWithContext(ctx context.Context, request *CreateSmpnEpaRequest) (response *CreateSmpnEpaResponse, err error) {
    if request == nil {
        request = NewCreateSmpnEpaRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateSmpnEpa require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateSmpnEpaResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSmpnChpRequest() (request *DescribeSmpnChpRequest) {
    request = &DescribeSmpnChpRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("smpn", APIVersion, "DescribeSmpnChp")
    
    
    return
}

func NewDescribeSmpnChpResponse() (response *DescribeSmpnChpResponse) {
    response = &DescribeSmpnChpResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeSmpnChp
// 查询号码的标记和标记次数
//
// 可能返回的错误码:
//  INTERNALERROR_TIMEOUT = "InternalError.Timeout"
//  INVALIDPARAMETER_PHONENUMBER = "InvalidParameter.PhoneNumber"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeSmpnChp(request *DescribeSmpnChpRequest) (response *DescribeSmpnChpResponse, err error) {
    return c.DescribeSmpnChpWithContext(context.Background(), request)
}

// DescribeSmpnChp
// 查询号码的标记和标记次数
//
// 可能返回的错误码:
//  INTERNALERROR_TIMEOUT = "InternalError.Timeout"
//  INVALIDPARAMETER_PHONENUMBER = "InvalidParameter.PhoneNumber"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeSmpnChpWithContext(ctx context.Context, request *DescribeSmpnChpRequest) (response *DescribeSmpnChpResponse, err error) {
    if request == nil {
        request = NewDescribeSmpnChpRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSmpnChp require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSmpnChpResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSmpnFnrRequest() (request *DescribeSmpnFnrRequest) {
    request = &DescribeSmpnFnrRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("smpn", APIVersion, "DescribeSmpnFnr")
    
    
    return
}

func NewDescribeSmpnFnrResponse() (response *DescribeSmpnFnrResponse) {
    response = &DescribeSmpnFnrResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeSmpnFnr
// 虚假号码识别
//
// 可能返回的错误码:
//  INTERNALERROR_TIMEOUT = "InternalError.Timeout"
//  INVALIDPARAMETER_PHONENUMBER = "InvalidParameter.PhoneNumber"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeSmpnFnr(request *DescribeSmpnFnrRequest) (response *DescribeSmpnFnrResponse, err error) {
    return c.DescribeSmpnFnrWithContext(context.Background(), request)
}

// DescribeSmpnFnr
// 虚假号码识别
//
// 可能返回的错误码:
//  INTERNALERROR_TIMEOUT = "InternalError.Timeout"
//  INVALIDPARAMETER_PHONENUMBER = "InvalidParameter.PhoneNumber"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeSmpnFnrWithContext(ctx context.Context, request *DescribeSmpnFnrRequest) (response *DescribeSmpnFnrResponse, err error) {
    if request == nil {
        request = NewDescribeSmpnFnrRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSmpnFnr require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSmpnFnrResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSmpnMhmRequest() (request *DescribeSmpnMhmRequest) {
    request = &DescribeSmpnMhmRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("smpn", APIVersion, "DescribeSmpnMhm")
    
    
    return
}

func NewDescribeSmpnMhmResponse() (response *DescribeSmpnMhmResponse) {
    response = &DescribeSmpnMhmResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeSmpnMhm
// 号码营销监控
//
// 可能返回的错误码:
//  INTERNALERROR_TIMEOUT = "InternalError.Timeout"
//  INVALIDPARAMETER_PHONENUMBER = "InvalidParameter.PhoneNumber"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeSmpnMhm(request *DescribeSmpnMhmRequest) (response *DescribeSmpnMhmResponse, err error) {
    return c.DescribeSmpnMhmWithContext(context.Background(), request)
}

// DescribeSmpnMhm
// 号码营销监控
//
// 可能返回的错误码:
//  INTERNALERROR_TIMEOUT = "InternalError.Timeout"
//  INVALIDPARAMETER_PHONENUMBER = "InvalidParameter.PhoneNumber"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeSmpnMhmWithContext(ctx context.Context, request *DescribeSmpnMhmRequest) (response *DescribeSmpnMhmResponse, err error) {
    if request == nil {
        request = NewDescribeSmpnMhmRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSmpnMhm require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSmpnMhmResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSmpnMrlRequest() (request *DescribeSmpnMrlRequest) {
    request = &DescribeSmpnMrlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("smpn", APIVersion, "DescribeSmpnMrl")
    
    
    return
}

func NewDescribeSmpnMrlResponse() (response *DescribeSmpnMrlResponse) {
    response = &DescribeSmpnMrlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeSmpnMrl
// 查询号码恶意标记等级
//
// 可能返回的错误码:
//  INTERNALERROR_TIMEOUT = "InternalError.Timeout"
//  INVALIDPARAMETER_PHONENUMBER = "InvalidParameter.PhoneNumber"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeSmpnMrl(request *DescribeSmpnMrlRequest) (response *DescribeSmpnMrlResponse, err error) {
    return c.DescribeSmpnMrlWithContext(context.Background(), request)
}

// DescribeSmpnMrl
// 查询号码恶意标记等级
//
// 可能返回的错误码:
//  INTERNALERROR_TIMEOUT = "InternalError.Timeout"
//  INVALIDPARAMETER_PHONENUMBER = "InvalidParameter.PhoneNumber"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeSmpnMrlWithContext(ctx context.Context, request *DescribeSmpnMrlRequest) (response *DescribeSmpnMrlResponse, err error) {
    if request == nil {
        request = NewDescribeSmpnMrlRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSmpnMrl require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSmpnMrlResponse()
    err = c.Send(request, response)
    return
}
