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

package v20190318

import (
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2019-03-18"

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


func NewAIAssistantRequest() (request *AIAssistantRequest) {
    request = &AIAssistantRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tci", APIVersion, "AIAssistant")
    return
}

func NewAIAssistantResponse() (response *AIAssistantResponse) {
    response = &AIAssistantResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 提供 AI 助教基础版本功能接口
func (c *Client) AIAssistant(request *AIAssistantRequest) (response *AIAssistantResponse, err error) {
    if request == nil {
        request = NewAIAssistantRequest()
    }
    response = NewAIAssistantResponse()
    err = c.Send(request, response)
    return
}

func NewCancelTaskRequest() (request *CancelTaskRequest) {
    request = &CancelTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tci", APIVersion, "CancelTask")
    return
}

func NewCancelTaskResponse() (response *CancelTaskResponse) {
    response = &CancelTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 用于取消已经提交的任务
func (c *Client) CancelTask(request *CancelTaskRequest) (response *CancelTaskResponse, err error) {
    if request == nil {
        request = NewCancelTaskRequest()
    }
    response = NewCancelTaskResponse()
    err = c.Send(request, response)
    return
}

func NewCheckAttendanceRequest() (request *CheckAttendanceRequest) {
    request = &CheckAttendanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tci", APIVersion, "CheckAttendance")
    return
}

func NewCheckAttendanceResponse() (response *CheckAttendanceResponse) {
    response = &CheckAttendanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 人员考勤
func (c *Client) CheckAttendance(request *CheckAttendanceRequest) (response *CheckAttendanceResponse, err error) {
    if request == nil {
        request = NewCheckAttendanceRequest()
    }
    response = NewCheckAttendanceResponse()
    err = c.Send(request, response)
    return
}

func NewCheckFacePhotoRequest() (request *CheckFacePhotoRequest) {
    request = &CheckFacePhotoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tci", APIVersion, "CheckFacePhoto")
    return
}

func NewCheckFacePhotoResponse() (response *CheckFacePhotoResponse) {
    response = &CheckFacePhotoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 检查人脸图片是否合法
func (c *Client) CheckFacePhoto(request *CheckFacePhotoRequest) (response *CheckFacePhotoResponse, err error) {
    if request == nil {
        request = NewCheckFacePhotoRequest()
    }
    response = NewCheckFacePhotoResponse()
    err = c.Send(request, response)
    return
}

func NewCreateFaceRequest() (request *CreateFaceRequest) {
    request = &CreateFaceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tci", APIVersion, "CreateFace")
    return
}

func NewCreateFaceResponse() (response *CreateFaceResponse) {
    response = &CreateFaceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建人脸
func (c *Client) CreateFace(request *CreateFaceRequest) (response *CreateFaceResponse, err error) {
    if request == nil {
        request = NewCreateFaceRequest()
    }
    response = NewCreateFaceResponse()
    err = c.Send(request, response)
    return
}

func NewCreateLibraryRequest() (request *CreateLibraryRequest) {
    request = &CreateLibraryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tci", APIVersion, "CreateLibrary")
    return
}

func NewCreateLibraryResponse() (response *CreateLibraryResponse) {
    response = &CreateLibraryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建人员库
func (c *Client) CreateLibrary(request *CreateLibraryRequest) (response *CreateLibraryResponse, err error) {
    if request == nil {
        request = NewCreateLibraryRequest()
    }
    response = NewCreateLibraryResponse()
    err = c.Send(request, response)
    return
}

func NewCreateVocabRequest() (request *CreateVocabRequest) {
    request = &CreateVocabRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tci", APIVersion, "CreateVocab")
    return
}

func NewCreateVocabResponse() (response *CreateVocabResponse) {
    response = &CreateVocabResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建词汇
func (c *Client) CreateVocab(request *CreateVocabRequest) (response *CreateVocabResponse, err error) {
    if request == nil {
        request = NewCreateVocabRequest()
    }
    response = NewCreateVocabResponse()
    err = c.Send(request, response)
    return
}

func NewCreateVocabLibRequest() (request *CreateVocabLibRequest) {
    request = &CreateVocabLibRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tci", APIVersion, "CreateVocabLib")
    return
}

func NewCreateVocabLibResponse() (response *CreateVocabLibResponse) {
    response = &CreateVocabLibResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 建立词汇库
func (c *Client) CreateVocabLib(request *CreateVocabLibRequest) (response *CreateVocabLibResponse, err error) {
    if request == nil {
        request = NewCreateVocabLibRequest()
    }
    response = NewCreateVocabLibResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteFaceRequest() (request *DeleteFaceRequest) {
    request = &DeleteFaceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tci", APIVersion, "DeleteFace")
    return
}

func NewDeleteFaceResponse() (response *DeleteFaceResponse) {
    response = &DeleteFaceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除人脸
func (c *Client) DeleteFace(request *DeleteFaceRequest) (response *DeleteFaceResponse, err error) {
    if request == nil {
        request = NewDeleteFaceRequest()
    }
    response = NewDeleteFaceResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteLibraryRequest() (request *DeleteLibraryRequest) {
    request = &DeleteLibraryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tci", APIVersion, "DeleteLibrary")
    return
}

func NewDeleteLibraryResponse() (response *DeleteLibraryResponse) {
    response = &DeleteLibraryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除人员库
func (c *Client) DeleteLibrary(request *DeleteLibraryRequest) (response *DeleteLibraryResponse, err error) {
    if request == nil {
        request = NewDeleteLibraryRequest()
    }
    response = NewDeleteLibraryResponse()
    err = c.Send(request, response)
    return
}

func NewDeletePersonRequest() (request *DeletePersonRequest) {
    request = &DeletePersonRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tci", APIVersion, "DeletePerson")
    return
}

func NewDeletePersonResponse() (response *DeletePersonResponse) {
    response = &DeletePersonResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除人员
func (c *Client) DeletePerson(request *DeletePersonRequest) (response *DeletePersonResponse, err error) {
    if request == nil {
        request = NewDeletePersonRequest()
    }
    response = NewDeletePersonResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteVocabRequest() (request *DeleteVocabRequest) {
    request = &DeleteVocabRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tci", APIVersion, "DeleteVocab")
    return
}

func NewDeleteVocabResponse() (response *DeleteVocabResponse) {
    response = &DeleteVocabResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除词汇
func (c *Client) DeleteVocab(request *DeleteVocabRequest) (response *DeleteVocabResponse, err error) {
    if request == nil {
        request = NewDeleteVocabRequest()
    }
    response = NewDeleteVocabResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteVocabLibRequest() (request *DeleteVocabLibRequest) {
    request = &DeleteVocabLibRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tci", APIVersion, "DeleteVocabLib")
    return
}

func NewDeleteVocabLibResponse() (response *DeleteVocabLibResponse) {
    response = &DeleteVocabLibResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除词汇库
func (c *Client) DeleteVocabLib(request *DeleteVocabLibRequest) (response *DeleteVocabLibResponse, err error) {
    if request == nil {
        request = NewDeleteVocabLibRequest()
    }
    response = NewDeleteVocabLibResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAITaskResultRequest() (request *DescribeAITaskResultRequest) {
    request = &DescribeAITaskResultRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tci", APIVersion, "DescribeAITaskResult")
    return
}

func NewDescribeAITaskResultResponse() (response *DescribeAITaskResultResponse) {
    response = &DescribeAITaskResultResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取标准化接口任务结果
func (c *Client) DescribeAITaskResult(request *DescribeAITaskResultRequest) (response *DescribeAITaskResultResponse, err error) {
    if request == nil {
        request = NewDescribeAITaskResultRequest()
    }
    response = NewDescribeAITaskResultResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAttendanceResultRequest() (request *DescribeAttendanceResultRequest) {
    request = &DescribeAttendanceResultRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tci", APIVersion, "DescribeAttendanceResult")
    return
}

func NewDescribeAttendanceResultResponse() (response *DescribeAttendanceResultResponse) {
    response = &DescribeAttendanceResultResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 人脸考勤查询结果
func (c *Client) DescribeAttendanceResult(request *DescribeAttendanceResultRequest) (response *DescribeAttendanceResultResponse, err error) {
    if request == nil {
        request = NewDescribeAttendanceResultRequest()
    }
    response = NewDescribeAttendanceResultResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAudioTaskRequest() (request *DescribeAudioTaskRequest) {
    request = &DescribeAudioTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tci", APIVersion, "DescribeAudioTask")
    return
}

func NewDescribeAudioTaskResponse() (response *DescribeAudioTaskResponse) {
    response = &DescribeAudioTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 音频评估任务信息查询接口，异步查询客户提交的请求的结果。
func (c *Client) DescribeAudioTask(request *DescribeAudioTaskRequest) (response *DescribeAudioTaskResponse, err error) {
    if request == nil {
        request = NewDescribeAudioTaskRequest()
    }
    response = NewDescribeAudioTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeConversationTaskRequest() (request *DescribeConversationTaskRequest) {
    request = &DescribeConversationTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tci", APIVersion, "DescribeConversationTask")
    return
}

func NewDescribeConversationTaskResponse() (response *DescribeConversationTaskResponse) {
    response = &DescribeConversationTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 音频对话任务评估任务信息查询接口，异步查询客户提交的请求的结果。
func (c *Client) DescribeConversationTask(request *DescribeConversationTaskRequest) (response *DescribeConversationTaskResponse, err error) {
    if request == nil {
        request = NewDescribeConversationTaskRequest()
    }
    response = NewDescribeConversationTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeHighlightResultRequest() (request *DescribeHighlightResultRequest) {
    request = &DescribeHighlightResultRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tci", APIVersion, "DescribeHighlightResult")
    return
}

func NewDescribeHighlightResultResponse() (response *DescribeHighlightResultResponse) {
    response = &DescribeHighlightResultResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 视频精彩集锦结果查询接口，异步查询客户提交的请求的结果。
func (c *Client) DescribeHighlightResult(request *DescribeHighlightResultRequest) (response *DescribeHighlightResultResponse, err error) {
    if request == nil {
        request = NewDescribeHighlightResultRequest()
    }
    response = NewDescribeHighlightResultResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeImageTaskRequest() (request *DescribeImageTaskRequest) {
    request = &DescribeImageTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tci", APIVersion, "DescribeImageTask")
    return
}

func NewDescribeImageTaskResponse() (response *DescribeImageTaskResponse) {
    response = &DescribeImageTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 拉取任务详情
func (c *Client) DescribeImageTask(request *DescribeImageTaskRequest) (response *DescribeImageTaskResponse, err error) {
    if request == nil {
        request = NewDescribeImageTaskRequest()
    }
    response = NewDescribeImageTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeImageTaskStatisticRequest() (request *DescribeImageTaskStatisticRequest) {
    request = &DescribeImageTaskStatisticRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tci", APIVersion, "DescribeImageTaskStatistic")
    return
}

func NewDescribeImageTaskStatisticResponse() (response *DescribeImageTaskStatisticResponse) {
    response = &DescribeImageTaskStatisticResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取图像任务统计信息
func (c *Client) DescribeImageTaskStatistic(request *DescribeImageTaskStatisticRequest) (response *DescribeImageTaskStatisticResponse, err error) {
    if request == nil {
        request = NewDescribeImageTaskStatisticRequest()
    }
    response = NewDescribeImageTaskStatisticResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLibrariesRequest() (request *DescribeLibrariesRequest) {
    request = &DescribeLibrariesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tci", APIVersion, "DescribeLibraries")
    return
}

func NewDescribeLibrariesResponse() (response *DescribeLibrariesResponse) {
    response = &DescribeLibrariesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取人员库列表
func (c *Client) DescribeLibraries(request *DescribeLibrariesRequest) (response *DescribeLibrariesResponse, err error) {
    if request == nil {
        request = NewDescribeLibrariesRequest()
    }
    response = NewDescribeLibrariesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePersonRequest() (request *DescribePersonRequest) {
    request = &DescribePersonRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tci", APIVersion, "DescribePerson")
    return
}

func NewDescribePersonResponse() (response *DescribePersonResponse) {
    response = &DescribePersonResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取人员详情
func (c *Client) DescribePerson(request *DescribePersonRequest) (response *DescribePersonResponse, err error) {
    if request == nil {
        request = NewDescribePersonRequest()
    }
    response = NewDescribePersonResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePersonsRequest() (request *DescribePersonsRequest) {
    request = &DescribePersonsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tci", APIVersion, "DescribePersons")
    return
}

func NewDescribePersonsResponse() (response *DescribePersonsResponse) {
    response = &DescribePersonsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 拉取人员列表
func (c *Client) DescribePersons(request *DescribePersonsRequest) (response *DescribePersonsResponse, err error) {
    if request == nil {
        request = NewDescribePersonsRequest()
    }
    response = NewDescribePersonsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVocabRequest() (request *DescribeVocabRequest) {
    request = &DescribeVocabRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tci", APIVersion, "DescribeVocab")
    return
}

func NewDescribeVocabResponse() (response *DescribeVocabResponse) {
    response = &DescribeVocabResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询词汇
func (c *Client) DescribeVocab(request *DescribeVocabRequest) (response *DescribeVocabResponse, err error) {
    if request == nil {
        request = NewDescribeVocabRequest()
    }
    response = NewDescribeVocabResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVocabLibRequest() (request *DescribeVocabLibRequest) {
    request = &DescribeVocabLibRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tci", APIVersion, "DescribeVocabLib")
    return
}

func NewDescribeVocabLibResponse() (response *DescribeVocabLibResponse) {
    response = &DescribeVocabLibResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询词汇库
func (c *Client) DescribeVocabLib(request *DescribeVocabLibRequest) (response *DescribeVocabLibResponse, err error) {
    if request == nil {
        request = NewDescribeVocabLibRequest()
    }
    response = NewDescribeVocabLibResponse()
    err = c.Send(request, response)
    return
}

func NewModifyLibraryRequest() (request *ModifyLibraryRequest) {
    request = &ModifyLibraryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tci", APIVersion, "ModifyLibrary")
    return
}

func NewModifyLibraryResponse() (response *ModifyLibraryResponse) {
    response = &ModifyLibraryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改人员库信息
func (c *Client) ModifyLibrary(request *ModifyLibraryRequest) (response *ModifyLibraryResponse, err error) {
    if request == nil {
        request = NewModifyLibraryRequest()
    }
    response = NewModifyLibraryResponse()
    err = c.Send(request, response)
    return
}

func NewModifyPersonRequest() (request *ModifyPersonRequest) {
    request = &ModifyPersonRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tci", APIVersion, "ModifyPerson")
    return
}

func NewModifyPersonResponse() (response *ModifyPersonResponse) {
    response = &ModifyPersonResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改人员信息
func (c *Client) ModifyPerson(request *ModifyPersonRequest) (response *ModifyPersonResponse, err error) {
    if request == nil {
        request = NewModifyPersonRequest()
    }
    response = NewModifyPersonResponse()
    err = c.Send(request, response)
    return
}

func NewSubmitAudioTaskRequest() (request *SubmitAudioTaskRequest) {
    request = &SubmitAudioTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tci", APIVersion, "SubmitAudioTask")
    return
}

func NewSubmitAudioTaskResponse() (response *SubmitAudioTaskResponse) {
    response = &SubmitAudioTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 音频任务提交接口
func (c *Client) SubmitAudioTask(request *SubmitAudioTaskRequest) (response *SubmitAudioTaskResponse, err error) {
    if request == nil {
        request = NewSubmitAudioTaskRequest()
    }
    response = NewSubmitAudioTaskResponse()
    err = c.Send(request, response)
    return
}

func NewSubmitCheckAttendanceTaskRequest() (request *SubmitCheckAttendanceTaskRequest) {
    request = &SubmitCheckAttendanceTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tci", APIVersion, "SubmitCheckAttendanceTask")
    return
}

func NewSubmitCheckAttendanceTaskResponse() (response *SubmitCheckAttendanceTaskResponse) {
    response = &SubmitCheckAttendanceTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 提交人员考勤任务
func (c *Client) SubmitCheckAttendanceTask(request *SubmitCheckAttendanceTaskRequest) (response *SubmitCheckAttendanceTaskResponse, err error) {
    if request == nil {
        request = NewSubmitCheckAttendanceTaskRequest()
    }
    response = NewSubmitCheckAttendanceTaskResponse()
    err = c.Send(request, response)
    return
}

func NewSubmitDoubleVideoHighlightsRequest() (request *SubmitDoubleVideoHighlightsRequest) {
    request = &SubmitDoubleVideoHighlightsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tci", APIVersion, "SubmitDoubleVideoHighlights")
    return
}

func NewSubmitDoubleVideoHighlightsResponse() (response *SubmitDoubleVideoHighlightsResponse) {
    response = &SubmitDoubleVideoHighlightsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 发起双路视频生成精彩集锦接口。该接口可以通过客户传入的学生音视频及老师视频两路Url，自动生成一堂课程的精彩集锦。需要通过SubmitDoubleVideoHighlights接口获取生成结果。
func (c *Client) SubmitDoubleVideoHighlights(request *SubmitDoubleVideoHighlightsRequest) (response *SubmitDoubleVideoHighlightsResponse, err error) {
    if request == nil {
        request = NewSubmitDoubleVideoHighlightsRequest()
    }
    response = NewSubmitDoubleVideoHighlightsResponse()
    err = c.Send(request, response)
    return
}

func NewSubmitHighlightsRequest() (request *SubmitHighlightsRequest) {
    request = &SubmitHighlightsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tci", APIVersion, "SubmitHighlights")
    return
}

func NewSubmitHighlightsResponse() (response *SubmitHighlightsResponse) {
    response = &SubmitHighlightsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 发起视频生成精彩集锦接口。该接口可以通过客户传入的课程音频数据及相关策略（如微笑抽取，专注抽取等），自动生成一堂课程的精彩集锦。需要通过QueryHighlightResult接口获取生成结果。
func (c *Client) SubmitHighlights(request *SubmitHighlightsRequest) (response *SubmitHighlightsResponse, err error) {
    if request == nil {
        request = NewSubmitHighlightsRequest()
    }
    response = NewSubmitHighlightsResponse()
    err = c.Send(request, response)
    return
}

func NewSubmitImageTaskRequest() (request *SubmitImageTaskRequest) {
    request = &SubmitImageTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tci", APIVersion, "SubmitImageTask")
    return
}

func NewSubmitImageTaskResponse() (response *SubmitImageTaskResponse) {
    response = &SubmitImageTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 提交图像分析任务
func (c *Client) SubmitImageTask(request *SubmitImageTaskRequest) (response *SubmitImageTaskResponse, err error) {
    if request == nil {
        request = NewSubmitImageTaskRequest()
    }
    response = NewSubmitImageTaskResponse()
    err = c.Send(request, response)
    return
}

func NewTransmitAudioStreamRequest() (request *TransmitAudioStreamRequest) {
    request = &TransmitAudioStreamRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tci", APIVersion, "TransmitAudioStream")
    return
}

func NewTransmitAudioStreamResponse() (response *TransmitAudioStreamResponse) {
    response = &TransmitAudioStreamResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 分析音频信息
func (c *Client) TransmitAudioStream(request *TransmitAudioStreamRequest) (response *TransmitAudioStreamResponse, err error) {
    if request == nil {
        request = NewTransmitAudioStreamRequest()
    }
    response = NewTransmitAudioStreamResponse()
    err = c.Send(request, response)
    return
}
