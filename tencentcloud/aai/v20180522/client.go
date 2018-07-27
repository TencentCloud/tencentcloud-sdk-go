// Copyright 1999-2018 Tencent Ltd.
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

package v20180522

import (
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2018-05-22"

type Client struct {
    common.Client
}

func NewClientWithSecretId(secretId, secretKey, region string) (client *Client, err error) {
    client = &Client{}
    client.Init(region).WithSecretId(secretId, secretKey)
    return
}

func NewClient(credential *common.Credential, region string, clientProfile *profile.ClientProfile) (client *Client, err error) {
    client = &Client{}
    client.Init(region).
        WithSecretId(credential.SecretId, credential.SecretKey).
        WithProfile(clientProfile)
    return
}


func NewSentenceRecognitionRequest() (request *SentenceRecognitionRequest) {
    request = &SentenceRecognitionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("aai", APIVersion, "SentenceRecognition")
    return
}

func NewSentenceRecognitionResponse() (response *SentenceRecognitionResponse) {
    response = &SentenceRecognitionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 识别60s内的短语音，当音频放在请求body中传输时整个请求大小不能超过1M，当音频以url方式传输时，音频时长不可超过60s。所有请求参数放在post的body中采用x-www-form-urlencoded（数据转换成一个字串（name1=value1&name2=value2…）进行urlencode后传输）编码传输。
func (c *Client) SentenceRecognition(request *SentenceRecognitionRequest) (response *SentenceRecognitionResponse, err error) {
    if request == nil {
        request = NewSentenceRecognitionRequest()
    }
    response = NewSentenceRecognitionResponse()
    err = c.Send(request, response)
    return
}

func NewTextToVoiceRequest() (request *TextToVoiceRequest) {
    request = &TextToVoiceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("aai", APIVersion, "TextToVoice")
    return
}

func NewTextToVoiceResponse() (response *TextToVoiceResponse) {
    response = &TextToVoiceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 腾讯云语音合成技术（TTS）可以将任意文本转化为语音，实现让机器和应用张口说话。
// 腾讯TTS技术可以应用到很多场景，比如，移动APP语音播报新闻；智能设备语音提醒；依靠网上现有节目或少量录音，快速合成明星语音，降低邀约成本；支持车载导航语音合成的个性化语音播报。
// 内测期间免费使用。
func (c *Client) TextToVoice(request *TextToVoiceRequest) (response *TextToVoiceResponse, err error) {
    if request == nil {
        request = NewTextToVoiceRequest()
    }
    response = NewTextToVoiceResponse()
    err = c.Send(request, response)
    return
}
