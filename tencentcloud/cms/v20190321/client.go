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


func NewAudioModerationRequest() (request *AudioModerationRequest) {
    request = &AudioModerationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cms", APIVersion, "AudioModeration")
    return
}

func NewAudioModerationResponse() (response *AudioModerationResponse) {
    response = &AudioModerationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 音频内容检测（Audio Moderation, AM）服务使用了波形分析、声纹分析等技术，能识别涉黄、涉政、涉恐等违规音频，同时支持用户配置音频黑库，打击自定义的违规内容。
// 
// <br>
// 接口返回值说明：调用本接口有两个返回值，一个是同步返回值，一个是识别完成后的异步回调返回值。
// 
// 音频识别结果存在于异步回调返回值中，异步回调返回值明细：
// 
// 参数名 | 类型 | 描述
// -|-|-
// SeqID | String | 请求seqId唯一标识
// EvilFlag | Integer | 是否恶意：0正常，1可疑（Homology模块下：0未匹配到，1恶意，2白样本）
// EvilType | Integer | 恶意类型：100正常，20001政治，20002色情，20007谩骂
// Duration | Integer | 音频时长（单位：毫秒）
// PornDetect | | 音频智能鉴黄
// PolityDetect | | 音频涉政识别
// CurseDetect | | 音频谩骂识别
// Homology | | 相似度识别
// HitFlag | Integer | 0正常，1可疑
// Score | Integer | 判断分值
// Keywords | Array of String | 关键词明细
// StartTime | Array of String | 恶意开始时间
// EndTime | Array of String | 恶意结束时间
// SeedUrl | String | 命中的种子URL
func (c *Client) AudioModeration(request *AudioModerationRequest) (response *AudioModerationResponse, err error) {
    if request == nil {
        request = NewAudioModerationRequest()
    }
    response = NewAudioModerationResponse()
    err = c.Send(request, response)
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

// 通过该接口可以将文件新增到样本库
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

// 新增文本类型样本库
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

// 删除文件样本库，支持批量删除，一次提交不超过20个
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

// 删除文字样本库，暂时只支持单个删除
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

// 查询文件样本库，支持批量查询
func (c *Client) DescribeFileSample(request *DescribeFileSampleRequest) (response *DescribeFileSampleResponse, err error) {
    if request == nil {
        request = NewDescribeFileSampleRequest()
    }
    response = NewDescribeFileSampleResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeModerationOverviewRequest() (request *DescribeModerationOverviewRequest) {
    request = &DescribeModerationOverviewRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cms", APIVersion, "DescribeModerationOverview")
    return
}

func NewDescribeModerationOverviewResponse() (response *DescribeModerationOverviewResponse) {
    response = &DescribeModerationOverviewResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 根据日期，渠道和服务类型查询识别结果概览数据
func (c *Client) DescribeModerationOverview(request *DescribeModerationOverviewRequest) (response *DescribeModerationOverviewResponse, err error) {
    if request == nil {
        request = NewDescribeModerationOverviewRequest()
    }
    response = NewDescribeModerationOverviewResponse()
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

// 支持批量查询文字样本库
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

func NewVideoModerationRequest() (request *VideoModerationRequest) {
    request = &VideoModerationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cms", APIVersion, "VideoModeration")
    return
}

func NewVideoModerationResponse() (response *VideoModerationResponse) {
    response = &VideoModerationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 视频内容检测（Video Moderation, VM）服务能识别涉黄、涉政、涉恐等违规视频，同时支持用户配置视频黑库，打击自定义的违规内容。
// 
// <br>
// 接口返回值说明：调用本接口有两个返回值，一个是同步返回值，一个是识别完成后的异步回调返回值。
// 
// 视频识别结果存在于异步回调返回值中，异步回调返回值明细：
// 
// 参数名 | 类型 | 描述
// -|-|-
// SeqID | String | 请求seqId唯一标识
// EvilFlag | Integer | 是否恶意：0正常，1可疑（Homology模块下：0未匹配到，1恶意，2白样本）
// EvilType | Integer | 恶意类型：100正常，20001政治，20002色情
// Duration | Integer | 视频时长（单位：秒）
// PornDetect |  | 视频智能鉴黄
// PolityDetect | | 视频涉政识别
// Homology | | 相似度识别
// HitFlag | Integer  | 0正常，1可疑
// Score | Integer | 判断分值
// SeedUrl | String | 命中的种子URL
func (c *Client) VideoModeration(request *VideoModerationRequest) (response *VideoModerationResponse, err error) {
    if request == nil {
        request = NewVideoModerationRequest()
    }
    response = NewVideoModerationResponse()
    err = c.Send(request, response)
    return
}
