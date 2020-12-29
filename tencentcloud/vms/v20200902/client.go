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

package v20200902

import (
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2020-09-02"

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


func NewSendCodeVoiceRequest() (request *SendCodeVoiceRequest) {
    request = &SendCodeVoiceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vms", APIVersion, "SendCodeVoice")
    return
}

func NewSendCodeVoiceResponse() (response *SendCodeVoiceResponse) {
    response = &SendCodeVoiceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 给用户发语音验证码（仅支持数字）。
func (c *Client) SendCodeVoice(request *SendCodeVoiceRequest) (response *SendCodeVoiceResponse, err error) {
    if request == nil {
        request = NewSendCodeVoiceRequest()
    }
    response = NewSendCodeVoiceResponse()
    err = c.Send(request, response)
    return
}

func NewSendTtsVoiceRequest() (request *SendTtsVoiceRequest) {
    request = &SendTtsVoiceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vms", APIVersion, "SendTtsVoice")
    return
}

func NewSendTtsVoiceResponse() (response *SendTtsVoiceResponse) {
    response = &SendTtsVoiceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 给用户发送指定模板的语音通知。
func (c *Client) SendTtsVoice(request *SendTtsVoiceRequest) (response *SendTtsVoiceResponse, err error) {
    if request == nil {
        request = NewSendTtsVoiceRequest()
    }
    response = NewSendTtsVoiceResponse()
    err = c.Send(request, response)
    return
}
