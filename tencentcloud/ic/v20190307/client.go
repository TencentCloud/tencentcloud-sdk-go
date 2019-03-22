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

package v20190307

import (
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2019-03-07"

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


func NewDescribeAppRequest() (request *DescribeAppRequest) {
    request = &DescribeAppRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ic", APIVersion, "DescribeApp")
    return
}

func NewDescribeAppResponse() (response *DescribeAppResponse) {
    response = &DescribeAppResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 根据应用id查询物联卡应用详情
func (c *Client) DescribeApp(request *DescribeAppRequest) (response *DescribeAppResponse, err error) {
    if request == nil {
        request = NewDescribeAppRequest()
    }
    response = NewDescribeAppResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCardRequest() (request *DescribeCardRequest) {
    request = &DescribeCardRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ic", APIVersion, "DescribeCard")
    return
}

func NewDescribeCardResponse() (response *DescribeCardResponse) {
    response = &DescribeCardResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询卡片详细信息
func (c *Client) DescribeCard(request *DescribeCardRequest) (response *DescribeCardResponse, err error) {
    if request == nil {
        request = NewDescribeCardRequest()
    }
    response = NewDescribeCardResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCardsRequest() (request *DescribeCardsRequest) {
    request = &DescribeCardsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ic", APIVersion, "DescribeCards")
    return
}

func NewDescribeCardsResponse() (response *DescribeCardsResponse) {
    response = &DescribeCardsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询卡片列表信息
func (c *Client) DescribeCards(request *DescribeCardsRequest) (response *DescribeCardsResponse, err error) {
    if request == nil {
        request = NewDescribeCardsRequest()
    }
    response = NewDescribeCardsResponse()
    err = c.Send(request, response)
    return
}

func NewSendMultiSmsRequest() (request *SendMultiSmsRequest) {
    request = &SendMultiSmsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ic", APIVersion, "SendMultiSms")
    return
}

func NewSendMultiSmsResponse() (response *SendMultiSmsResponse) {
    response = &SendMultiSmsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 群发短信
func (c *Client) SendMultiSms(request *SendMultiSmsRequest) (response *SendMultiSmsResponse, err error) {
    if request == nil {
        request = NewSendMultiSmsRequest()
    }
    response = NewSendMultiSmsResponse()
    err = c.Send(request, response)
    return
}

func NewSendSmsRequest() (request *SendSmsRequest) {
    request = &SendSmsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ic", APIVersion, "SendSms")
    return
}

func NewSendSmsResponse() (response *SendSmsResponse) {
    response = &SendSmsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 发送短信息接口
func (c *Client) SendSms(request *SendSmsRequest) (response *SendSmsResponse, err error) {
    if request == nil {
        request = NewSendSmsRequest()
    }
    response = NewSendSmsResponse()
    err = c.Send(request, response)
    return
}
