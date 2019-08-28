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

// 用于取消已经提交的任务，目前只支持图像任务。
func (c *Client) CancelTask(request *CancelTaskRequest) (response *CancelTaskResponse, err error) {
    if request == nil {
        request = NewCancelTaskRequest()
    }
    response = NewCancelTaskResponse()
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

func NewCreatePersonRequest() (request *CreatePersonRequest) {
    request = &CreatePersonRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tci", APIVersion, "CreatePerson")
    return
}

func NewCreatePersonResponse() (response *CreatePersonResponse) {
    response = &CreatePersonResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建人员
func (c *Client) CreatePerson(request *CreatePersonRequest) (response *CreatePersonResponse, err error) {
    if request == nil {
        request = NewCreatePersonRequest()
    }
    response = NewCreatePersonResponse()
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

// 提交人员考勤任务，支持包括点播和直播资源；支持通过DescribeAttendanceResult查询结果，也支持通过NoticeUrl设置考勤回调结果，回调结果结构如下：
// ##### 回调事件结构
//  | 参数名称 | 类型 | 描述 | 
//  | ----  | ---  | ------  |
//  | jobid | Integer | 任务ID | 
//  | person_info | array of PersonInfo | 识别到的人员列表 | 
// #####子结构PersonInfo
//  | 参数名称 | 类型 | 描述 | 
//  | ----  | ---  | ------  |
//  | traceid | String | 可用于区分同一路视频流下的不同陌生人 | 
//  | personid | String | 识别到的人员ID，如果是陌生人则返回空串 | 
//  | libid | String | 识别到的人员所在的库ID，如果是陌生人则返回空串 | 
//  | timestamp | uint64 | 识别到人脸的绝对时间戳，单位ms | 
//  | image_url | string | 识别到人脸的事件抓图的下载地址，不长期保存，需要请及时下载 | 
func (c *Client) SubmitCheckAttendanceTask(request *SubmitCheckAttendanceTaskRequest) (response *SubmitCheckAttendanceTaskResponse, err error) {
    if request == nil {
        request = NewSubmitCheckAttendanceTaskRequest()
    }
    response = NewSubmitCheckAttendanceTaskResponse()
    err = c.Send(request, response)
    return
}

func NewSubmitConversationTaskRequest() (request *SubmitConversationTaskRequest) {
    request = &SubmitConversationTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tci", APIVersion, "SubmitConversationTask")
    return
}

func NewSubmitConversationTaskResponse() (response *SubmitConversationTaskResponse) {
    response = &SubmitConversationTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 对话任务分析接口
func (c *Client) SubmitConversationTask(request *SubmitConversationTaskRequest) (response *SubmitConversationTaskResponse, err error) {
    if request == nil {
        request = NewSubmitConversationTaskRequest()
    }
    response = NewSubmitConversationTaskResponse()
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

// 发起双路视频生成精彩集锦接口。该接口可以通过客户传入的学生音视频及老师视频两路Url，自动生成一堂课程的精彩集锦。需要通过DescribeHighlightResult
// 接口获取生成结果。
func (c *Client) SubmitDoubleVideoHighlights(request *SubmitDoubleVideoHighlightsRequest) (response *SubmitDoubleVideoHighlightsResponse, err error) {
    if request == nil {
        request = NewSubmitDoubleVideoHighlightsRequest()
    }
    response = NewSubmitDoubleVideoHighlightsResponse()
    err = c.Send(request, response)
    return
}

func NewSubmitFullBodyClassTaskRequest() (request *SubmitFullBodyClassTaskRequest) {
    request = &SubmitFullBodyClassTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tci", APIVersion, "SubmitFullBodyClassTask")
    return
}

func NewSubmitFullBodyClassTaskResponse() (response *SubmitFullBodyClassTaskResponse) {
    response = &SubmitFullBodyClassTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// **传统课堂授课任务**：在此场景中，老师为站立授课，有白板或投影供老师展示课程内容，摄像头可以拍摄到老师的半身或者全身。拍摄视频为一路全局画面，且背景不动，要求画面稳定清晰。通过此接口可分析老师授课的行为及语音，以支持AI评教。  
//   
// **提供的功能接口有：**老师人脸识别、老师表情识别、老师肢体动作识别、语音识别。  可分析的指标维度包括：身份识别、正脸、侧脸、人脸坐标、人脸尺寸、高兴、中性、高兴、中性、惊讶、厌恶、恐惧、愤怒、蔑视、悲伤、正面讲解、写板书、指黑板、语音转文字、发音时长、非发音时长、音量、语速、指定关键词的使用等
//   
// **对场景的要求为：**真实场景老师1人出现在画面中，全局画面且背景不动；人脸上下角度在20度以内，左右角度在15度以内，歪头角度在15度以内；光照均匀，无遮挡，人脸清晰可见；像素最好在 100X100 像素以上，但是图像整体质量不能超过1080p。
//     
// **结果查询方式：**图像任务直接返回结果，点播及直播任务通过DescribeAITaskResult查询结果。
func (c *Client) SubmitFullBodyClassTask(request *SubmitFullBodyClassTaskRequest) (response *SubmitFullBodyClassTaskResponse, err error) {
    if request == nil {
        request = NewSubmitFullBodyClassTaskRequest()
    }
    response = NewSubmitFullBodyClassTaskResponse()
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

func NewSubmitOneByOneClassTaskRequest() (request *SubmitOneByOneClassTaskRequest) {
    request = &SubmitOneByOneClassTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tci", APIVersion, "SubmitOneByOneClassTask")
    return
}

func NewSubmitOneByOneClassTaskResponse() (response *SubmitOneByOneClassTaskResponse) {
    response = &SubmitOneByOneClassTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// **提交在线1对1课堂任务**  
// 对于在线1对1课堂，老师通过视频向学生授课，并且学生人数为1人。通过上传学生端的图像信息，可以获取学生的听课情况分析。 具体指一路全局画面且背景不动，有1位学生的头像或上半身的画面，要求画面稳定清晰。
//   
// **提供的功能接口有：**学生人脸识别、学生表情识别、语音识别。可分析的指标维度包括：学生身份识别、正脸、侧脸、抬头、低头、人脸坐标、人脸尺寸、高兴、中性、高兴、中性、惊讶、厌恶、恐惧、愤怒、蔑视、悲伤、语音转文字、发音时长、非发音时长、音量、语速等。
//   
// **对场景的要求为：**真实常规1v1授课场景，学生2人以下，全局画面且背景不动；人脸上下角度在20度以内，左右角度在15度以内，歪头角度在15度以内；光照均匀，无遮挡，人脸清晰可见；像素最好在 100X100 像素以上，但是图像整体质量不能超过1080p。
//     
// **结果查询方式：**图像任务直接返回结果，点播及直播任务通过DescribeAITaskResult查询结果。
func (c *Client) SubmitOneByOneClassTask(request *SubmitOneByOneClassTaskRequest) (response *SubmitOneByOneClassTaskResponse, err error) {
    if request == nil {
        request = NewSubmitOneByOneClassTaskRequest()
    }
    response = NewSubmitOneByOneClassTaskResponse()
    err = c.Send(request, response)
    return
}

func NewSubmitOpenClassTaskRequest() (request *SubmitOpenClassTaskRequest) {
    request = &SubmitOpenClassTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tci", APIVersion, "SubmitOpenClassTask")
    return
}

func NewSubmitOpenClassTaskResponse() (response *SubmitOpenClassTaskResponse) {
    response = &SubmitOpenClassTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// **提交线下小班（无课桌）课任务**  
// 线下小班课是指有学生无课桌的课堂，满座15人以下，全局画面且背景不动，能清晰看到。  
//   
// **提供的功能接口有：**学生人脸识别、学生表情识别、学生肢体动作识别。  可分析的指标维度包括：身份识别、正脸、侧脸、抬头、低头、高兴、中性、高兴、中性、惊讶、厌恶、恐惧、愤怒、蔑视、悲伤、站立、举手、坐着等。
//   
// **对场景的要求为：**真实常规教室，满座15人以下，全局画面且背景不动；人脸上下角度在20度以内，左右角度在15度以内，歪头角度在15度以内；光照均匀，无遮挡，人脸清晰可见；像素最好在 100X100 像素以上但是图像整体质量不能超过1080p。
//     
// **结果查询方式：**图像任务直接返回结果，点播及直播任务通过DescribeAITaskResult查询结果。
func (c *Client) SubmitOpenClassTask(request *SubmitOpenClassTaskRequest) (response *SubmitOpenClassTaskResponse, err error) {
    if request == nil {
        request = NewSubmitOpenClassTaskRequest()
    }
    response = NewSubmitOpenClassTaskResponse()
    err = c.Send(request, response)
    return
}

func NewSubmitPartialBodyClassTaskRequest() (request *SubmitPartialBodyClassTaskRequest) {
    request = &SubmitPartialBodyClassTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tci", APIVersion, "SubmitPartialBodyClassTask")
    return
}

func NewSubmitPartialBodyClassTaskResponse() (response *SubmitPartialBodyClassTaskResponse) {
    response = &SubmitPartialBodyClassTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// **在线小班课任务**：此场景是在线授课场景，老师一般为坐着授课，摄像头可以拍摄到老师的头部及上半身。拍摄视频为一路全局画面，且背景不动，要求画面稳定清晰。通过此接口可分析老师授课的行为及语音，以支持AI评教。    
//   
// **提供的功能接口有：**老师人脸识别、老师表情识别、老师手势识别、光线识别、语音识别。 可分析的指标维度包括：身份识别、正脸、侧脸、人脸坐标、人脸尺寸、高兴、中性、高兴、中性、惊讶、厌恶、恐惧、愤怒、蔑视、悲伤、点赞手势、听你说手势、听我说手势、拿教具行为、语音转文字、发音时长、非发音时长、音量、语速、指定关键词的使用等 
//   
// **对场景的要求为：**在线常规授课场景，全局画面且背景不动；人脸上下角度在20度以内，左右角度在15度以内，歪头角度在15度以内；光照均匀，无遮挡，人脸清晰可见；像素最好在 100X100 像素以上，但是图像整体质量不能超过1080p。
//     
// **结果查询方式：**图像任务直接返回结果，点播及直播任务通过DescribeAITaskResult查询结果。
func (c *Client) SubmitPartialBodyClassTask(request *SubmitPartialBodyClassTaskRequest) (response *SubmitPartialBodyClassTaskResponse, err error) {
    if request == nil {
        request = NewSubmitPartialBodyClassTaskRequest()
    }
    response = NewSubmitPartialBodyClassTaskResponse()
    err = c.Send(request, response)
    return
}

func NewSubmitTraditionalClassTaskRequest() (request *SubmitTraditionalClassTaskRequest) {
    request = &SubmitTraditionalClassTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tci", APIVersion, "SubmitTraditionalClassTask")
    return
}

func NewSubmitTraditionalClassTaskResponse() (response *SubmitTraditionalClassTaskResponse) {
    response = &SubmitTraditionalClassTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// **提交线下传统面授大班课（含课桌）任务。**  
// 传统教室课堂是指有学生课堂有课桌的课堂，满座20-50人，全局画面且背景不动。  
//   
// **提供的功能接口有：**学生人脸识别、学生表情识别、学生肢体动作识别。可分析的指标维度包括：学生身份识别、正脸、侧脸、抬头、低头、高兴、中性、高兴、中性、惊讶、厌恶、恐惧、愤怒、蔑视、悲伤、举手、站立、坐着、趴桌子、玩手机等  
//   
// **对场景的要求为：**传统的学生上课教室，满座20-50人，全局画面且背景不动；人脸上下角度在20度以内，左右角度在15度以内，歪头角度在15度以内；光照均匀，无遮挡，人脸清晰可见；像素最好在 100X100 像素以上，但是图像整体质量不能超过1080p。
//     
// **结果查询方式：**图像任务直接返回结果，点播及直播任务通过DescribeAITaskResult查询结果。
//   
func (c *Client) SubmitTraditionalClassTask(request *SubmitTraditionalClassTaskRequest) (response *SubmitTraditionalClassTaskResponse, err error) {
    if request == nil {
        request = NewSubmitTraditionalClassTaskRequest()
    }
    response = NewSubmitTraditionalClassTaskResponse()
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
