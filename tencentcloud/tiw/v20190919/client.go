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

package v20190919

import (
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2019-09-19"

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


func NewCreateTranscodeRequest() (request *CreateTranscodeRequest) {
    request = &CreateTranscodeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tiw", APIVersion, "CreateTranscode")
    return
}

func NewCreateTranscodeResponse() (response *CreateTranscodeResponse) {
    response = &CreateTranscodeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建一个文档转码任务
func (c *Client) CreateTranscode(request *CreateTranscodeRequest) (response *CreateTranscodeResponse, err error) {
    if request == nil {
        request = NewCreateTranscodeRequest()
    }
    response = NewCreateTranscodeResponse()
    err = c.Send(request, response)
    return
}

func NewCreateVideoGenerationTaskRequest() (request *CreateVideoGenerationTaskRequest) {
    request = &CreateVideoGenerationTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tiw", APIVersion, "CreateVideoGenerationTask")
    return
}

func NewCreateVideoGenerationTaskResponse() (response *CreateVideoGenerationTaskResponse) {
    response = &CreateVideoGenerationTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建视频生成任务
func (c *Client) CreateVideoGenerationTask(request *CreateVideoGenerationTaskRequest) (response *CreateVideoGenerationTaskResponse, err error) {
    if request == nil {
        request = NewCreateVideoGenerationTaskRequest()
    }
    response = NewCreateVideoGenerationTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeOnlineRecordRequest() (request *DescribeOnlineRecordRequest) {
    request = &DescribeOnlineRecordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tiw", APIVersion, "DescribeOnlineRecord")
    return
}

func NewDescribeOnlineRecordResponse() (response *DescribeOnlineRecordResponse) {
    response = &DescribeOnlineRecordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询录制任务状态与结果
func (c *Client) DescribeOnlineRecord(request *DescribeOnlineRecordRequest) (response *DescribeOnlineRecordResponse, err error) {
    if request == nil {
        request = NewDescribeOnlineRecordRequest()
    }
    response = NewDescribeOnlineRecordResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeOnlineRecordCallbackRequest() (request *DescribeOnlineRecordCallbackRequest) {
    request = &DescribeOnlineRecordCallbackRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tiw", APIVersion, "DescribeOnlineRecordCallback")
    return
}

func NewDescribeOnlineRecordCallbackResponse() (response *DescribeOnlineRecordCallbackResponse) {
    response = &DescribeOnlineRecordCallbackResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询实时录制回调地址
func (c *Client) DescribeOnlineRecordCallback(request *DescribeOnlineRecordCallbackRequest) (response *DescribeOnlineRecordCallbackResponse, err error) {
    if request == nil {
        request = NewDescribeOnlineRecordCallbackRequest()
    }
    response = NewDescribeOnlineRecordCallbackResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeQualityMetricsRequest() (request *DescribeQualityMetricsRequest) {
    request = &DescribeQualityMetricsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tiw", APIVersion, "DescribeQualityMetrics")
    return
}

func NewDescribeQualityMetricsResponse() (response *DescribeQualityMetricsResponse) {
    response = &DescribeQualityMetricsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询互动白板质量数据
func (c *Client) DescribeQualityMetrics(request *DescribeQualityMetricsRequest) (response *DescribeQualityMetricsResponse, err error) {
    if request == nil {
        request = NewDescribeQualityMetricsRequest()
    }
    response = NewDescribeQualityMetricsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTranscodeRequest() (request *DescribeTranscodeRequest) {
    request = &DescribeTranscodeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tiw", APIVersion, "DescribeTranscode")
    return
}

func NewDescribeTranscodeResponse() (response *DescribeTranscodeResponse) {
    response = &DescribeTranscodeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询文档转码任务的执行进度与转码结果
func (c *Client) DescribeTranscode(request *DescribeTranscodeRequest) (response *DescribeTranscodeResponse, err error) {
    if request == nil {
        request = NewDescribeTranscodeRequest()
    }
    response = NewDescribeTranscodeResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTranscodeCallbackRequest() (request *DescribeTranscodeCallbackRequest) {
    request = &DescribeTranscodeCallbackRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tiw", APIVersion, "DescribeTranscodeCallback")
    return
}

func NewDescribeTranscodeCallbackResponse() (response *DescribeTranscodeCallbackResponse) {
    response = &DescribeTranscodeCallbackResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询文档转码回调地址
func (c *Client) DescribeTranscodeCallback(request *DescribeTranscodeCallbackRequest) (response *DescribeTranscodeCallbackResponse, err error) {
    if request == nil {
        request = NewDescribeTranscodeCallbackRequest()
    }
    response = NewDescribeTranscodeCallbackResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVideoGenerationTaskRequest() (request *DescribeVideoGenerationTaskRequest) {
    request = &DescribeVideoGenerationTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tiw", APIVersion, "DescribeVideoGenerationTask")
    return
}

func NewDescribeVideoGenerationTaskResponse() (response *DescribeVideoGenerationTaskResponse) {
    response = &DescribeVideoGenerationTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询录制视频生成任务状态与结果
func (c *Client) DescribeVideoGenerationTask(request *DescribeVideoGenerationTaskRequest) (response *DescribeVideoGenerationTaskResponse, err error) {
    if request == nil {
        request = NewDescribeVideoGenerationTaskRequest()
    }
    response = NewDescribeVideoGenerationTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVideoGenerationTaskCallbackRequest() (request *DescribeVideoGenerationTaskCallbackRequest) {
    request = &DescribeVideoGenerationTaskCallbackRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tiw", APIVersion, "DescribeVideoGenerationTaskCallback")
    return
}

func NewDescribeVideoGenerationTaskCallbackResponse() (response *DescribeVideoGenerationTaskCallbackResponse) {
    response = &DescribeVideoGenerationTaskCallbackResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询录制视频生成回调地址
func (c *Client) DescribeVideoGenerationTaskCallback(request *DescribeVideoGenerationTaskCallbackRequest) (response *DescribeVideoGenerationTaskCallbackResponse, err error) {
    if request == nil {
        request = NewDescribeVideoGenerationTaskCallbackRequest()
    }
    response = NewDescribeVideoGenerationTaskCallbackResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeWhiteboardPushRequest() (request *DescribeWhiteboardPushRequest) {
    request = &DescribeWhiteboardPushRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tiw", APIVersion, "DescribeWhiteboardPush")
    return
}

func NewDescribeWhiteboardPushResponse() (response *DescribeWhiteboardPushResponse) {
    response = &DescribeWhiteboardPushResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询推流任务状态与结果
func (c *Client) DescribeWhiteboardPush(request *DescribeWhiteboardPushRequest) (response *DescribeWhiteboardPushResponse, err error) {
    if request == nil {
        request = NewDescribeWhiteboardPushRequest()
    }
    response = NewDescribeWhiteboardPushResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeWhiteboardPushCallbackRequest() (request *DescribeWhiteboardPushCallbackRequest) {
    request = &DescribeWhiteboardPushCallbackRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tiw", APIVersion, "DescribeWhiteboardPushCallback")
    return
}

func NewDescribeWhiteboardPushCallbackResponse() (response *DescribeWhiteboardPushCallbackResponse) {
    response = &DescribeWhiteboardPushCallbackResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询白板推流回调地址
func (c *Client) DescribeWhiteboardPushCallback(request *DescribeWhiteboardPushCallbackRequest) (response *DescribeWhiteboardPushCallbackResponse, err error) {
    if request == nil {
        request = NewDescribeWhiteboardPushCallbackRequest()
    }
    response = NewDescribeWhiteboardPushCallbackResponse()
    err = c.Send(request, response)
    return
}

func NewPauseOnlineRecordRequest() (request *PauseOnlineRecordRequest) {
    request = &PauseOnlineRecordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tiw", APIVersion, "PauseOnlineRecord")
    return
}

func NewPauseOnlineRecordResponse() (response *PauseOnlineRecordResponse) {
    response = &PauseOnlineRecordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 暂停实时录制
func (c *Client) PauseOnlineRecord(request *PauseOnlineRecordRequest) (response *PauseOnlineRecordResponse, err error) {
    if request == nil {
        request = NewPauseOnlineRecordRequest()
    }
    response = NewPauseOnlineRecordResponse()
    err = c.Send(request, response)
    return
}

func NewResumeOnlineRecordRequest() (request *ResumeOnlineRecordRequest) {
    request = &ResumeOnlineRecordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tiw", APIVersion, "ResumeOnlineRecord")
    return
}

func NewResumeOnlineRecordResponse() (response *ResumeOnlineRecordResponse) {
    response = &ResumeOnlineRecordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 恢复实时录制
func (c *Client) ResumeOnlineRecord(request *ResumeOnlineRecordRequest) (response *ResumeOnlineRecordResponse, err error) {
    if request == nil {
        request = NewResumeOnlineRecordRequest()
    }
    response = NewResumeOnlineRecordResponse()
    err = c.Send(request, response)
    return
}

func NewSetOnlineRecordCallbackRequest() (request *SetOnlineRecordCallbackRequest) {
    request = &SetOnlineRecordCallbackRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tiw", APIVersion, "SetOnlineRecordCallback")
    return
}

func NewSetOnlineRecordCallbackResponse() (response *SetOnlineRecordCallbackResponse) {
    response = &SetOnlineRecordCallbackResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 设置实时录制回调地址，回调数据格式请参考文档：https://cloud.tencent.com/document/product/1137/40258
func (c *Client) SetOnlineRecordCallback(request *SetOnlineRecordCallbackRequest) (response *SetOnlineRecordCallbackResponse, err error) {
    if request == nil {
        request = NewSetOnlineRecordCallbackRequest()
    }
    response = NewSetOnlineRecordCallbackResponse()
    err = c.Send(request, response)
    return
}

func NewSetOnlineRecordCallbackKeyRequest() (request *SetOnlineRecordCallbackKeyRequest) {
    request = &SetOnlineRecordCallbackKeyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tiw", APIVersion, "SetOnlineRecordCallbackKey")
    return
}

func NewSetOnlineRecordCallbackKeyResponse() (response *SetOnlineRecordCallbackKeyResponse) {
    response = &SetOnlineRecordCallbackKeyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 设置实时录制回调鉴权密钥，回调鉴权方式请参考文档：https://cloud.tencent.com/document/product/1137/40257
func (c *Client) SetOnlineRecordCallbackKey(request *SetOnlineRecordCallbackKeyRequest) (response *SetOnlineRecordCallbackKeyResponse, err error) {
    if request == nil {
        request = NewSetOnlineRecordCallbackKeyRequest()
    }
    response = NewSetOnlineRecordCallbackKeyResponse()
    err = c.Send(request, response)
    return
}

func NewSetTranscodeCallbackRequest() (request *SetTranscodeCallbackRequest) {
    request = &SetTranscodeCallbackRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tiw", APIVersion, "SetTranscodeCallback")
    return
}

func NewSetTranscodeCallbackResponse() (response *SetTranscodeCallbackResponse) {
    response = &SetTranscodeCallbackResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 设置文档转码回调地址，回调数据格式请参考文档：https://cloud.tencent.com/document/product/1137/40260
func (c *Client) SetTranscodeCallback(request *SetTranscodeCallbackRequest) (response *SetTranscodeCallbackResponse, err error) {
    if request == nil {
        request = NewSetTranscodeCallbackRequest()
    }
    response = NewSetTranscodeCallbackResponse()
    err = c.Send(request, response)
    return
}

func NewSetTranscodeCallbackKeyRequest() (request *SetTranscodeCallbackKeyRequest) {
    request = &SetTranscodeCallbackKeyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tiw", APIVersion, "SetTranscodeCallbackKey")
    return
}

func NewSetTranscodeCallbackKeyResponse() (response *SetTranscodeCallbackKeyResponse) {
    response = &SetTranscodeCallbackKeyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 设置文档转码回调鉴权密钥，回调鉴权方式请参考文档：https://cloud.tencent.com/document/product/1137/40257
func (c *Client) SetTranscodeCallbackKey(request *SetTranscodeCallbackKeyRequest) (response *SetTranscodeCallbackKeyResponse, err error) {
    if request == nil {
        request = NewSetTranscodeCallbackKeyRequest()
    }
    response = NewSetTranscodeCallbackKeyResponse()
    err = c.Send(request, response)
    return
}

func NewSetVideoGenerationTaskCallbackRequest() (request *SetVideoGenerationTaskCallbackRequest) {
    request = &SetVideoGenerationTaskCallbackRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tiw", APIVersion, "SetVideoGenerationTaskCallback")
    return
}

func NewSetVideoGenerationTaskCallbackResponse() (response *SetVideoGenerationTaskCallbackResponse) {
    response = &SetVideoGenerationTaskCallbackResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 设置录制视频生成回调地址
func (c *Client) SetVideoGenerationTaskCallback(request *SetVideoGenerationTaskCallbackRequest) (response *SetVideoGenerationTaskCallbackResponse, err error) {
    if request == nil {
        request = NewSetVideoGenerationTaskCallbackRequest()
    }
    response = NewSetVideoGenerationTaskCallbackResponse()
    err = c.Send(request, response)
    return
}

func NewSetVideoGenerationTaskCallbackKeyRequest() (request *SetVideoGenerationTaskCallbackKeyRequest) {
    request = &SetVideoGenerationTaskCallbackKeyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tiw", APIVersion, "SetVideoGenerationTaskCallbackKey")
    return
}

func NewSetVideoGenerationTaskCallbackKeyResponse() (response *SetVideoGenerationTaskCallbackKeyResponse) {
    response = &SetVideoGenerationTaskCallbackKeyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 设置视频生成回调鉴权密钥
func (c *Client) SetVideoGenerationTaskCallbackKey(request *SetVideoGenerationTaskCallbackKeyRequest) (response *SetVideoGenerationTaskCallbackKeyResponse, err error) {
    if request == nil {
        request = NewSetVideoGenerationTaskCallbackKeyRequest()
    }
    response = NewSetVideoGenerationTaskCallbackKeyResponse()
    err = c.Send(request, response)
    return
}

func NewSetWhiteboardPushCallbackRequest() (request *SetWhiteboardPushCallbackRequest) {
    request = &SetWhiteboardPushCallbackRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tiw", APIVersion, "SetWhiteboardPushCallback")
    return
}

func NewSetWhiteboardPushCallbackResponse() (response *SetWhiteboardPushCallbackResponse) {
    response = &SetWhiteboardPushCallbackResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 设置白板推流回调地址，回调数据格式请参考文档：https://cloud.tencent.com/document/product/1137/40257
func (c *Client) SetWhiteboardPushCallback(request *SetWhiteboardPushCallbackRequest) (response *SetWhiteboardPushCallbackResponse, err error) {
    if request == nil {
        request = NewSetWhiteboardPushCallbackRequest()
    }
    response = NewSetWhiteboardPushCallbackResponse()
    err = c.Send(request, response)
    return
}

func NewSetWhiteboardPushCallbackKeyRequest() (request *SetWhiteboardPushCallbackKeyRequest) {
    request = &SetWhiteboardPushCallbackKeyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tiw", APIVersion, "SetWhiteboardPushCallbackKey")
    return
}

func NewSetWhiteboardPushCallbackKeyResponse() (response *SetWhiteboardPushCallbackKeyResponse) {
    response = &SetWhiteboardPushCallbackKeyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 设置白板推流回调鉴权密钥，回调鉴权方式请参考文档：https://cloud.tencent.com/document/product/1137/40257
func (c *Client) SetWhiteboardPushCallbackKey(request *SetWhiteboardPushCallbackKeyRequest) (response *SetWhiteboardPushCallbackKeyResponse, err error) {
    if request == nil {
        request = NewSetWhiteboardPushCallbackKeyRequest()
    }
    response = NewSetWhiteboardPushCallbackKeyResponse()
    err = c.Send(request, response)
    return
}

func NewStartOnlineRecordRequest() (request *StartOnlineRecordRequest) {
    request = &StartOnlineRecordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tiw", APIVersion, "StartOnlineRecord")
    return
}

func NewStartOnlineRecordResponse() (response *StartOnlineRecordResponse) {
    response = &StartOnlineRecordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 发起一个实时录制任务
func (c *Client) StartOnlineRecord(request *StartOnlineRecordRequest) (response *StartOnlineRecordResponse, err error) {
    if request == nil {
        request = NewStartOnlineRecordRequest()
    }
    response = NewStartOnlineRecordResponse()
    err = c.Send(request, response)
    return
}

func NewStartWhiteboardPushRequest() (request *StartWhiteboardPushRequest) {
    request = &StartWhiteboardPushRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tiw", APIVersion, "StartWhiteboardPush")
    return
}

func NewStartWhiteboardPushResponse() (response *StartWhiteboardPushResponse) {
    response = &StartWhiteboardPushResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 发起一个白板推流任务
func (c *Client) StartWhiteboardPush(request *StartWhiteboardPushRequest) (response *StartWhiteboardPushResponse, err error) {
    if request == nil {
        request = NewStartWhiteboardPushRequest()
    }
    response = NewStartWhiteboardPushResponse()
    err = c.Send(request, response)
    return
}

func NewStopOnlineRecordRequest() (request *StopOnlineRecordRequest) {
    request = &StopOnlineRecordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tiw", APIVersion, "StopOnlineRecord")
    return
}

func NewStopOnlineRecordResponse() (response *StopOnlineRecordResponse) {
    response = &StopOnlineRecordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 停止实时录制
func (c *Client) StopOnlineRecord(request *StopOnlineRecordRequest) (response *StopOnlineRecordResponse, err error) {
    if request == nil {
        request = NewStopOnlineRecordRequest()
    }
    response = NewStopOnlineRecordResponse()
    err = c.Send(request, response)
    return
}

func NewStopWhiteboardPushRequest() (request *StopWhiteboardPushRequest) {
    request = &StopWhiteboardPushRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tiw", APIVersion, "StopWhiteboardPush")
    return
}

func NewStopWhiteboardPushResponse() (response *StopWhiteboardPushResponse) {
    response = &StopWhiteboardPushResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 停止白板推流任务
func (c *Client) StopWhiteboardPush(request *StopWhiteboardPushRequest) (response *StopWhiteboardPushResponse, err error) {
    if request == nil {
        request = NewStopWhiteboardPushRequest()
    }
    response = NewStopWhiteboardPushResponse()
    err = c.Send(request, response)
    return
}
