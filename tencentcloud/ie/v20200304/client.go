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

package v20200304

import (
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2020-03-04"

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


func NewCreateEditingTaskRequest() (request *CreateEditingTaskRequest) {
    request = &CreateEditingTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ie", APIVersion, "CreateEditingTask")
    return
}

func NewCreateEditingTaskResponse() (response *CreateEditingTaskResponse) {
    response = &CreateEditingTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建编辑理解任务，可以同时选择视频标签识别、分类识别、智能拆条、智能集锦、智能封面和片头片尾识别中的一项或者多项能力。
func (c *Client) CreateEditingTask(request *CreateEditingTaskRequest) (response *CreateEditingTaskResponse, err error) {
    if request == nil {
        request = NewCreateEditingTaskRequest()
    }
    response = NewCreateEditingTaskResponse()
    err = c.Send(request, response)
    return
}

func NewCreateMediaProcessTaskRequest() (request *CreateMediaProcessTaskRequest) {
    request = &CreateMediaProcessTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ie", APIVersion, "CreateMediaProcessTask")
    return
}

func NewCreateMediaProcessTaskResponse() (response *CreateMediaProcessTaskResponse) {
    response = &CreateMediaProcessTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 用于创建编辑处理任务，如媒体截取、媒体编辑、媒体拼接、媒体字幕。
func (c *Client) CreateMediaProcessTask(request *CreateMediaProcessTaskRequest) (response *CreateMediaProcessTaskResponse, err error) {
    if request == nil {
        request = NewCreateMediaProcessTaskRequest()
    }
    response = NewCreateMediaProcessTaskResponse()
    err = c.Send(request, response)
    return
}

func NewCreateMediaQualityRestorationTaskRequest() (request *CreateMediaQualityRestorationTaskRequest) {
    request = &CreateMediaQualityRestorationTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ie", APIVersion, "CreateMediaQualityRestorationTask")
    return
}

func NewCreateMediaQualityRestorationTaskResponse() (response *CreateMediaQualityRestorationTaskResponse) {
    response = &CreateMediaQualityRestorationTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建画质重生任务，对视频进行转码、去噪、去划痕、去毛刺、超分、细节增强和色彩增强。
func (c *Client) CreateMediaQualityRestorationTask(request *CreateMediaQualityRestorationTaskRequest) (response *CreateMediaQualityRestorationTaskResponse, err error) {
    if request == nil {
        request = NewCreateMediaQualityRestorationTaskRequest()
    }
    response = NewCreateMediaQualityRestorationTaskResponse()
    err = c.Send(request, response)
    return
}

func NewCreateQualityControlTaskRequest() (request *CreateQualityControlTaskRequest) {
    request = &CreateQualityControlTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ie", APIVersion, "CreateQualityControlTask")
    return
}

func NewCreateQualityControlTaskResponse() (response *CreateQualityControlTaskResponse) {
    response = &CreateQualityControlTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 通过接口可以智能检测视频画面中抖动重影、模糊、低光照、过曝光、黑边、白边、黑屏、白屏、花屏、噪点、马赛克、二维码等在内的多个场景，还可以自动检测视频无音频异常、无声音片段。
func (c *Client) CreateQualityControlTask(request *CreateQualityControlTaskRequest) (response *CreateQualityControlTaskResponse, err error) {
    if request == nil {
        request = NewCreateQualityControlTaskRequest()
    }
    response = NewCreateQualityControlTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEditingTaskResultRequest() (request *DescribeEditingTaskResultRequest) {
    request = &DescribeEditingTaskResultRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ie", APIVersion, "DescribeEditingTaskResult")
    return
}

func NewDescribeEditingTaskResultResponse() (response *DescribeEditingTaskResultResponse) {
    response = &DescribeEditingTaskResultResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取编辑理解任务结果。
func (c *Client) DescribeEditingTaskResult(request *DescribeEditingTaskResultRequest) (response *DescribeEditingTaskResultResponse, err error) {
    if request == nil {
        request = NewDescribeEditingTaskResultRequest()
    }
    response = NewDescribeEditingTaskResultResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMediaProcessTaskResultRequest() (request *DescribeMediaProcessTaskResultRequest) {
    request = &DescribeMediaProcessTaskResultRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ie", APIVersion, "DescribeMediaProcessTaskResult")
    return
}

func NewDescribeMediaProcessTaskResultResponse() (response *DescribeMediaProcessTaskResultResponse) {
    response = &DescribeMediaProcessTaskResultResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 用于获取编辑处理任务的结果。
func (c *Client) DescribeMediaProcessTaskResult(request *DescribeMediaProcessTaskResultRequest) (response *DescribeMediaProcessTaskResultResponse, err error) {
    if request == nil {
        request = NewDescribeMediaProcessTaskResultRequest()
    }
    response = NewDescribeMediaProcessTaskResultResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMediaQualityRestorationTaskRusultRequest() (request *DescribeMediaQualityRestorationTaskRusultRequest) {
    request = &DescribeMediaQualityRestorationTaskRusultRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ie", APIVersion, "DescribeMediaQualityRestorationTaskRusult")
    return
}

func NewDescribeMediaQualityRestorationTaskRusultResponse() (response *DescribeMediaQualityRestorationTaskRusultResponse) {
    response = &DescribeMediaQualityRestorationTaskRusultResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取画质重生任务结果，查看结束后的文件信息
func (c *Client) DescribeMediaQualityRestorationTaskRusult(request *DescribeMediaQualityRestorationTaskRusultRequest) (response *DescribeMediaQualityRestorationTaskRusultResponse, err error) {
    if request == nil {
        request = NewDescribeMediaQualityRestorationTaskRusultRequest()
    }
    response = NewDescribeMediaQualityRestorationTaskRusultResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeQualityControlTaskResultRequest() (request *DescribeQualityControlTaskResultRequest) {
    request = &DescribeQualityControlTaskResultRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ie", APIVersion, "DescribeQualityControlTaskResult")
    return
}

func NewDescribeQualityControlTaskResultResponse() (response *DescribeQualityControlTaskResultResponse) {
    response = &DescribeQualityControlTaskResultResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取媒体质检任务结果
func (c *Client) DescribeQualityControlTaskResult(request *DescribeQualityControlTaskResultRequest) (response *DescribeQualityControlTaskResultResponse, err error) {
    if request == nil {
        request = NewDescribeQualityControlTaskResultRequest()
    }
    response = NewDescribeQualityControlTaskResultResponse()
    err = c.Send(request, response)
    return
}

func NewStopMediaProcessTaskRequest() (request *StopMediaProcessTaskRequest) {
    request = &StopMediaProcessTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ie", APIVersion, "StopMediaProcessTask")
    return
}

func NewStopMediaProcessTaskResponse() (response *StopMediaProcessTaskResponse) {
    response = &StopMediaProcessTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 用于停止正在进行中的编辑处理任务。
func (c *Client) StopMediaProcessTask(request *StopMediaProcessTaskRequest) (response *StopMediaProcessTaskResponse, err error) {
    if request == nil {
        request = NewStopMediaProcessTaskRequest()
    }
    response = NewStopMediaProcessTaskResponse()
    err = c.Send(request, response)
    return
}

func NewStopMediaQualityRestorationTaskRequest() (request *StopMediaQualityRestorationTaskRequest) {
    request = &StopMediaQualityRestorationTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ie", APIVersion, "StopMediaQualityRestorationTask")
    return
}

func NewStopMediaQualityRestorationTaskResponse() (response *StopMediaQualityRestorationTaskResponse) {
    response = &StopMediaQualityRestorationTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除正在进行的画质重生任务
func (c *Client) StopMediaQualityRestorationTask(request *StopMediaQualityRestorationTaskRequest) (response *StopMediaQualityRestorationTaskResponse, err error) {
    if request == nil {
        request = NewStopMediaQualityRestorationTaskRequest()
    }
    response = NewStopMediaQualityRestorationTaskResponse()
    err = c.Send(request, response)
    return
}
