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

func NewClient(credential *common.Credential, region string, clientProfile *profile.ClientProfile) (client *Client, err error) {
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

// 企业号码认证
func (c *Client) CreateSmpnEpa(request *CreateSmpnEpaRequest) (response *CreateSmpnEpaResponse, err error) {
    if request == nil {
        request = NewCreateSmpnEpaRequest()
    }
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

// 查询号码的标记和标记次数
func (c *Client) DescribeSmpnChp(request *DescribeSmpnChpRequest) (response *DescribeSmpnChpResponse, err error) {
    if request == nil {
        request = NewDescribeSmpnChpRequest()
    }
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

// 虚假号码识别
func (c *Client) DescribeSmpnFnr(request *DescribeSmpnFnrRequest) (response *DescribeSmpnFnrResponse, err error) {
    if request == nil {
        request = NewDescribeSmpnFnrRequest()
    }
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

// 号码营销监控
func (c *Client) DescribeSmpnMhm(request *DescribeSmpnMhmRequest) (response *DescribeSmpnMhmResponse, err error) {
    if request == nil {
        request = NewDescribeSmpnMhmRequest()
    }
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

// 查询号码恶意标记等级
func (c *Client) DescribeSmpnMrl(request *DescribeSmpnMrlRequest) (response *DescribeSmpnMrlResponse, err error) {
    if request == nil {
        request = NewDescribeSmpnMrlRequest()
    }
    response = NewDescribeSmpnMrlResponse()
    err = c.Send(request, response)
    return
}
