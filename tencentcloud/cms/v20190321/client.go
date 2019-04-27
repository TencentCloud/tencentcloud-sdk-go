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
// 通过API直接上传音频即可进行检测，对于高危部分直接屏蔽，可疑部分人工复审，从而节省审核人力，释放业务风险。
func (c *Client) AudioModeration(request *AudioModerationRequest) (response *AudioModerationResponse, err error) {
    if request == nil {
        request = NewAudioModerationRequest()
    }
    response = NewAudioModerationResponse()
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
// 通过API获取检测的标签及置信度，可直接采信高置信度的结果，人工复审低置信度的结果，从而降低人工成本，提高审核效率。
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
// 通过API接口，能检测内容的危险等级，对于高危部分直接过滤，可疑部分人工复审，从而节省审核人力，释放业务风险。
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
// 通过API直接上传视频即可进行检测，对于高危部分直接过滤，可疑部分人工复审，从而节省审核人力，释放业务风险。
func (c *Client) VideoModeration(request *VideoModerationRequest) (response *VideoModerationResponse, err error) {
    if request == nil {
        request = NewVideoModerationRequest()
    }
    response = NewVideoModerationResponse()
    err = c.Send(request, response)
    return
}
