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

package v20200713

import (
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2020-07-13"

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


func NewAccountTipoffAccessRequest() (request *AccountTipoffAccessRequest) {
    request = &AccountTipoffAccessRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tms", APIVersion, "AccountTipoffAccess")
    return
}

func NewAccountTipoffAccessResponse() (response *AccountTipoffAccessResponse) {
    response = &AccountTipoffAccessResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 举报恶意账号
func (c *Client) AccountTipoffAccess(request *AccountTipoffAccessRequest) (response *AccountTipoffAccessResponse, err error) {
    if request == nil {
        request = NewAccountTipoffAccessRequest()
    }
    response = NewAccountTipoffAccessResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTextLibRequest() (request *DescribeTextLibRequest) {
    request = &DescribeTextLibRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tms", APIVersion, "DescribeTextLib")
    return
}

func NewDescribeTextLibResponse() (response *DescribeTextLibResponse) {
    response = &DescribeTextLibResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 控制台获取用户词库列表
func (c *Client) DescribeTextLib(request *DescribeTextLibRequest) (response *DescribeTextLibResponse, err error) {
    if request == nil {
        request = NewDescribeTextLibRequest()
    }
    response = NewDescribeTextLibResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTextStatRequest() (request *DescribeTextStatRequest) {
    request = &DescribeTextStatRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tms", APIVersion, "DescribeTextStat")
    return
}

func NewDescribeTextStatResponse() (response *DescribeTextStatResponse) {
    response = &DescribeTextStatResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 控制台识别统计
func (c *Client) DescribeTextStat(request *DescribeTextStatRequest) (response *DescribeTextStatResponse, err error) {
    if request == nil {
        request = NewDescribeTextStatRequest()
    }
    response = NewDescribeTextStatResponse()
    err = c.Send(request, response)
    return
}

func NewTextModerationRequest() (request *TextModerationRequest) {
    request = &TextModerationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tms", APIVersion, "TextModeration")
    return
}

func NewTextModerationResponse() (response *TextModerationResponse) {
    response = &TextModerationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 文本内容检测（Text Moderation）服务使用了深度学习技术，识别涉黄、涉政、涉恐等有害内容，同时支持用户配置词库，打击自定义的违规文本。
func (c *Client) TextModeration(request *TextModerationRequest) (response *TextModerationResponse, err error) {
    if request == nil {
        request = NewTextModerationRequest()
    }
    response = NewTextModerationResponse()
    err = c.Send(request, response)
    return
}
