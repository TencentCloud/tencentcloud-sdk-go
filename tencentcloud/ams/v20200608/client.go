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

package v20200608

import (
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2020-06-08"

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


func NewCreateAudioModerationTaskRequest() (request *CreateAudioModerationTaskRequest) {
    request = &CreateAudioModerationTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ams", APIVersion, "CreateAudioModerationTask")
    return
}

func NewCreateAudioModerationTaskResponse() (response *CreateAudioModerationTaskResponse) {
    response = &CreateAudioModerationTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（Audio Moderation）用于提交音频内容（包括音频文件或流地址）进行智能审核任务，使用前请您登陆控制台开通音频内容安全服务。
// 
// ### 功能使用说明：
// - 前往“内容安全控制台-音频内容安全”开启使用音频内容安全服务，首次开通可获得20小时免费调用时长
// 
// ### 接口功能说明：
// - 支持对音频流或音频文件进行检测，判断其中是否包含违规内容；
// - 支持设置回调地址 Callback 获取检测结果，或通过接口(查询音频检测结果)主动轮询获取检测结果；
// - 支持识别违规内容，包括：低俗、谩骂、色情、涉政、广告等场景；
// - 支持批量提交检测任务。检测任务列表最多支持10个；
// 
// ### 音频文件调用说明：
// - 音频文件大小支持：文件 < 500M；
// - 音频文件时长支持：< 1小时；
// - 音频码率类型支持：128 Kbps - 256 Kbps ；
// - 音频文件支持格式：wav、mp3、aac、flac、amr、3gp、 m4a、wma、ogg、ape；
// - 支持音视频文件分离并对音频文件进行独立识别；
// 
// ### 音频流调用说明：
// - 音频流时长支持：< 3小时；
// - 音频码率类型支持：128 Kbps - 256 Kbps ；
// - 音频流支持的传输协议：RTMP、HTTP、HTTPS；
// - 音频流格式支持的类型：rtp、srtp、rtmp、rtmps、mmsh、 mmst、hls、http、tcp、https、m3u8；
// - 支持音视频流分离并对音频流进行独立识别；
func (c *Client) CreateAudioModerationTask(request *CreateAudioModerationTaskRequest) (response *CreateAudioModerationTaskResponse, err error) {
    if request == nil {
        request = NewCreateAudioModerationTaskRequest()
    }
    response = NewCreateAudioModerationTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTaskDetailRequest() (request *DescribeTaskDetailRequest) {
    request = &DescribeTaskDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ams", APIVersion, "DescribeTaskDetail")
    return
}

func NewDescribeTaskDetailResponse() (response *DescribeTaskDetailResponse) {
    response = &DescribeTaskDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查看任务详情
func (c *Client) DescribeTaskDetail(request *DescribeTaskDetailRequest) (response *DescribeTaskDetailResponse, err error) {
    if request == nil {
        request = NewDescribeTaskDetailRequest()
    }
    response = NewDescribeTaskDetailResponse()
    err = c.Send(request, response)
    return
}
