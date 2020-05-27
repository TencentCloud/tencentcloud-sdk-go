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

package v20190321

import (
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2019-03-21"

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


func NewCreateFileSampleRequest() (request *CreateFileSampleRequest) {
    request = &CreateFileSampleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cms", APIVersion, "CreateFileSample")
    return
}

func NewCreateFileSampleResponse() (response *CreateFileSampleResponse) {
    response = &CreateFileSampleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本文档适用于图片内容安全、视频内容安全自定义识别库的管理。
// <br>
// 通过该接口可以将图片新增到样本库。
func (c *Client) CreateFileSample(request *CreateFileSampleRequest) (response *CreateFileSampleResponse, err error) {
    if request == nil {
        request = NewCreateFileSampleRequest()
    }
    response = NewCreateFileSampleResponse()
    err = c.Send(request, response)
    return
}

func NewCreateTextSampleRequest() (request *CreateTextSampleRequest) {
    request = &CreateTextSampleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cms", APIVersion, "CreateTextSample")
    return
}

func NewCreateTextSampleResponse() (response *CreateTextSampleResponse) {
    response = &CreateTextSampleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本文档适用于文本内容安全、音频内容安全自定义识别库的管理。
// <br>
// 通过该接口可以将文本新增到样本库。
func (c *Client) CreateTextSample(request *CreateTextSampleRequest) (response *CreateTextSampleResponse, err error) {
    if request == nil {
        request = NewCreateTextSampleRequest()
    }
    response = NewCreateTextSampleResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteFileSampleRequest() (request *DeleteFileSampleRequest) {
    request = &DeleteFileSampleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cms", APIVersion, "DeleteFileSample")
    return
}

func NewDeleteFileSampleResponse() (response *DeleteFileSampleResponse) {
    response = &DeleteFileSampleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本文档适用于图片内容安全、视频内容安全自定义识别库的管理。
// <br>
// 删除图片样本库，支持批量删除，一次提交不超过20个。
func (c *Client) DeleteFileSample(request *DeleteFileSampleRequest) (response *DeleteFileSampleResponse, err error) {
    if request == nil {
        request = NewDeleteFileSampleRequest()
    }
    response = NewDeleteFileSampleResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteTextSampleRequest() (request *DeleteTextSampleRequest) {
    request = &DeleteTextSampleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cms", APIVersion, "DeleteTextSample")
    return
}

func NewDeleteTextSampleResponse() (response *DeleteTextSampleResponse) {
    response = &DeleteTextSampleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本文档适用于文本内容安全、音频内容安全自定义识别库的管理。
// <br>
// 删除文本样本库，暂时只支持单个删除。
func (c *Client) DeleteTextSample(request *DeleteTextSampleRequest) (response *DeleteTextSampleResponse, err error) {
    if request == nil {
        request = NewDeleteTextSampleRequest()
    }
    response = NewDeleteTextSampleResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeFileSampleRequest() (request *DescribeFileSampleRequest) {
    request = &DescribeFileSampleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cms", APIVersion, "DescribeFileSample")
    return
}

func NewDescribeFileSampleResponse() (response *DescribeFileSampleResponse) {
    response = &DescribeFileSampleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本文档适用于图片内容安全、视频内容安全自定义识别库的管理。
// <br>
// 查询图片样本库，支持批量查询。
func (c *Client) DescribeFileSample(request *DescribeFileSampleRequest) (response *DescribeFileSampleResponse, err error) {
    if request == nil {
        request = NewDescribeFileSampleRequest()
    }
    response = NewDescribeFileSampleResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTextSampleRequest() (request *DescribeTextSampleRequest) {
    request = &DescribeTextSampleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cms", APIVersion, "DescribeTextSample")
    return
}

func NewDescribeTextSampleResponse() (response *DescribeTextSampleResponse) {
    response = &DescribeTextSampleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本文档适用于文本内容安全、音频内容安全自定义识别库的管理。
// <br>
// 支持批量查询文本样本库。
func (c *Client) DescribeTextSample(request *DescribeTextSampleRequest) (response *DescribeTextSampleResponse, err error) {
    if request == nil {
        request = NewDescribeTextSampleRequest()
    }
    response = NewDescribeTextSampleResponse()
    err = c.Send(request, response)
    return
}

func NewImageModerationRequest() (request *ImageModerationRequest) {
    request = &ImageModerationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cms", APIVersion, "ImageModeration")
    return
}

func NewImageModerationResponse() (response *ImageModerationResponse) {
    response = &ImageModerationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 图片内容检测服务（Image Moderation, IM）能自动扫描图片，识别涉黄、涉恐、涉政、涉毒等有害内容，同时支持用户配置图片黑名单，打击自定义的违规图片。
func (c *Client) ImageModeration(request *ImageModerationRequest) (response *ImageModerationResponse, err error) {
    if request == nil {
        request = NewImageModerationRequest()
    }
    response = NewImageModerationResponse()
    err = c.Send(request, response)
    return
}

func NewTextModerationRequest() (request *TextModerationRequest) {
    request = &TextModerationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cms", APIVersion, "TextModeration")
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
