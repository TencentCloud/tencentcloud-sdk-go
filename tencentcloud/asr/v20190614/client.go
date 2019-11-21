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

package v20190614

import (
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2019-06-14"

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


func NewCreateRecTaskRequest() (request *CreateRecTaskRequest) {
    request = &CreateRecTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("asr", APIVersion, "CreateRecTask")
    return
}

func NewCreateRecTaskResponse() (response *CreateRecTaskResponse) {
    response = &CreateRecTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口服务对录音时长1小时以内的录音文件进行识别，异步返回识别全部结果。
// <br>• 接口是 HTTP RESTful 形式
// <br>• 接口支持wav、mp3、silk、amr、m4a等主流音频格式
// <br>• 支持语音 URL 和本地语音文件两种请求方式
// <br>• 本地语音文件上传的文件不能大于5MB，语音 URL的音频时长不能长于1小时
// <br>• 目前仅支持中文普通话
// <br>• 支持回调或轮询的方式获取结果，结果获取请参考[ 录音文件识别结果查询](https://cloud.tencent.com/document/product/1093/37822)。
func (c *Client) CreateRecTask(request *CreateRecTaskRequest) (response *CreateRecTaskResponse, err error) {
    if request == nil {
        request = NewCreateRecTaskRequest()
    }
    response = NewCreateRecTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTaskStatusRequest() (request *DescribeTaskStatusRequest) {
    request = &DescribeTaskStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("asr", APIVersion, "DescribeTaskStatus")
    return
}

func NewDescribeTaskStatusResponse() (response *DescribeTaskStatusResponse) {
    response = &DescribeTaskStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 在调用录音文件识别请求接口后，有回调和轮询两种方式获取识别结果。
// <br>• 当采用回调方式时，识别完成后会将结果通过 POST 请求的形式通知到用户在请求时填写的回调 URL，具体请参见[ 录音识别结果回调 ](https://cloud.tencent.com/document/product/1093/37139#callback)。
// <br>• 当采用轮询方式时，需要主动提交任务ID来轮询识别结果，共有任务成功、等待、执行中和失败四种结果，具体信息请参见下文说明。
func (c *Client) DescribeTaskStatus(request *DescribeTaskStatusRequest) (response *DescribeTaskStatusResponse, err error) {
    if request == nil {
        request = NewDescribeTaskStatusRequest()
    }
    response = NewDescribeTaskStatusResponse()
    err = c.Send(request, response)
    return
}

func NewSentenceRecognitionRequest() (request *SentenceRecognitionRequest) {
    request = &SentenceRecognitionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("asr", APIVersion, "SentenceRecognition")
    return
}

func NewSentenceRecognitionResponse() (response *SentenceRecognitionResponse) {
    response = &SentenceRecognitionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口用于对60秒之内的短音频文件进行识别。
// <br>•   支持中文普通话、英语、粤语。
// <br>•   支持本地语音文件上传和语音URL上传两种请求方式。
// <br>•   音频格式支持wav、mp3；采样率支持8000Hz或者16000Hz；采样精度支持16bits；声道支持单声道。
// <br>•   当音频文件通过请求中body内容上传时，请求大小不能超过600KB；当音频以URL方式传输时，音频时长不可超过60s。
// <br>•   所有请求参数放在POST请求的body中，编码类型采用x-www-form-urlencoded，参数进行urlencode编码后传输。
func (c *Client) SentenceRecognition(request *SentenceRecognitionRequest) (response *SentenceRecognitionResponse, err error) {
    if request == nil {
        request = NewSentenceRecognitionRequest()
    }
    response = NewSentenceRecognitionResponse()
    err = c.Send(request, response)
    return
}
