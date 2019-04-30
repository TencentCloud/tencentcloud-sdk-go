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

package v20190311

import (
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2019-03-11"

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


func NewPostAudioRequest() (request *PostAudioRequest) {
    request = &PostAudioRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbp", APIVersion, "PostAudio")
    return
}

func NewPostAudioResponse() (response *PostAudioResponse) {
    response = &PostAudioResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 机器人会话接口，接收音频信息，传递给后台机器人
func (c *Client) PostAudio(request *PostAudioRequest) (response *PostAudioResponse, err error) {
    if request == nil {
        request = NewPostAudioRequest()
    }
    response = NewPostAudioResponse()
    err = c.Send(request, response)
    return
}

func NewPostTextRequest() (request *PostTextRequest) {
    request = &PostTextRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbp", APIVersion, "PostText")
    return
}

func NewPostTextResponse() (response *PostTextResponse) {
    response = &PostTextResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 机器人会话接口，接收文本信息，传递给后台机器人
func (c *Client) PostText(request *PostTextRequest) (response *PostTextResponse, err error) {
    if request == nil {
        request = NewPostTextRequest()
    }
    response = NewPostTextResponse()
    err = c.Send(request, response)
    return
}

func NewResetRequest() (request *ResetRequest) {
    request = &ResetRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbp", APIVersion, "Reset")
    return
}

func NewResetResponse() (response *ResetResponse) {
    response = &ResetResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 对当前机器人的会话状态进行复位
func (c *Client) Reset(request *ResetRequest) (response *ResetResponse, err error) {
    if request == nil {
        request = NewResetRequest()
    }
    response = NewResetResponse()
    err = c.Send(request, response)
    return
}
