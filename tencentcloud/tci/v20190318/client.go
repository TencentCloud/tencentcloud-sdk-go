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
    "context"
    "errors"
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

func NewClient(credential common.CredentialIface, region string, clientProfile *profile.ClientProfile) (client *Client, err error) {
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

// AIAssistant
// 提供 AI 助教基础版本功能接口
//
// 可能返回的错误码:
//  INVALIDPARAMETER_LIBHAVENOPERSON = "InvalidParameter.LibHaveNoPerson"
//  INVALIDPARAMETER_LIBISEMPTY = "InvalidParameter.LibIsEmpty"
//  INVALIDPARAMETER_PROCESSUNREADY = "InvalidParameter.ProcessUnReady"
//  INVALIDPARAMETER_STATUSJOBFAILED = "InvalidParameter.StatusJobFailed"
//  INVALIDPARAMETER_STATUSJOBUNFINISHED = "InvalidParameter.StatusJobUnfinished"
func (c *Client) AIAssistant(request *AIAssistantRequest) (response *AIAssistantResponse, err error) {
    return c.AIAssistantWithContext(context.Background(), request)
}

// AIAssistant
// 提供 AI 助教基础版本功能接口
//
// 可能返回的错误码:
//  INVALIDPARAMETER_LIBHAVENOPERSON = "InvalidParameter.LibHaveNoPerson"
//  INVALIDPARAMETER_LIBISEMPTY = "InvalidParameter.LibIsEmpty"
//  INVALIDPARAMETER_PROCESSUNREADY = "InvalidParameter.ProcessUnReady"
//  INVALIDPARAMETER_STATUSJOBFAILED = "InvalidParameter.StatusJobFailed"
//  INVALIDPARAMETER_STATUSJOBUNFINISHED = "InvalidParameter.StatusJobUnfinished"
func (c *Client) AIAssistantWithContext(ctx context.Context, request *AIAssistantRequest) (response *AIAssistantResponse, err error) {
    if request == nil {
        request = NewAIAssistantRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AIAssistant require credential")
    }

    request.SetContext(ctx)
    
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

// CancelTask
// 用于取消已经提交的任务，目前只支持图像任务。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_STATUSJOBFAILED = "InvalidParameter.StatusJobFailed"
//  RESOURCENOTFOUND_HLJOBIDNOTFOUND = "ResourceNotFound.HLJobIdNotFound"
func (c *Client) CancelTask(request *CancelTaskRequest) (response *CancelTaskResponse, err error) {
    return c.CancelTaskWithContext(context.Background(), request)
}

// CancelTask
// 用于取消已经提交的任务，目前只支持图像任务。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_STATUSJOBFAILED = "InvalidParameter.StatusJobFailed"
//  RESOURCENOTFOUND_HLJOBIDNOTFOUND = "ResourceNotFound.HLJobIdNotFound"
func (c *Client) CancelTaskWithContext(ctx context.Context, request *CancelTaskRequest) (response *CancelTaskResponse, err error) {
    if request == nil {
        request = NewCancelTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CancelTask require credential")
    }

    request.SetContext(ctx)
    
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

// CheckFacePhoto
// 检查人脸图片是否合法
//
// 可能返回的错误码:
//  INVALIDPARAMETER_STATUSJOBFAILED = "InvalidParameter.StatusJobFailed"
//  RESOURCENOTFOUND_HLJOBIDNOTFOUND = "ResourceNotFound.HLJobIdNotFound"
func (c *Client) CheckFacePhoto(request *CheckFacePhotoRequest) (response *CheckFacePhotoResponse, err error) {
    return c.CheckFacePhotoWithContext(context.Background(), request)
}

// CheckFacePhoto
// 检查人脸图片是否合法
//
// 可能返回的错误码:
//  INVALIDPARAMETER_STATUSJOBFAILED = "InvalidParameter.StatusJobFailed"
//  RESOURCENOTFOUND_HLJOBIDNOTFOUND = "ResourceNotFound.HLJobIdNotFound"
func (c *Client) CheckFacePhotoWithContext(ctx context.Context, request *CheckFacePhotoRequest) (response *CheckFacePhotoResponse, err error) {
    if request == nil {
        request = NewCheckFacePhotoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CheckFacePhoto require credential")
    }

    request.SetContext(ctx)
    
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

// CreateFace
// 创建人脸
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CreateFace(request *CreateFaceRequest) (response *CreateFaceResponse, err error) {
    return c.CreateFaceWithContext(context.Background(), request)
}

// CreateFace
// 创建人脸
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CreateFaceWithContext(ctx context.Context, request *CreateFaceRequest) (response *CreateFaceResponse, err error) {
    if request == nil {
        request = NewCreateFaceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateFace require credential")
    }

    request.SetContext(ctx)
    
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

// CreateLibrary
// 创建人员库
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CreateLibrary(request *CreateLibraryRequest) (response *CreateLibraryResponse, err error) {
    return c.CreateLibraryWithContext(context.Background(), request)
}

// CreateLibrary
// 创建人员库
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CreateLibraryWithContext(ctx context.Context, request *CreateLibraryRequest) (response *CreateLibraryResponse, err error) {
    if request == nil {
        request = NewCreateLibraryRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateLibrary require credential")
    }

    request.SetContext(ctx)
    
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

// CreatePerson
// 创建人员
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CreatePerson(request *CreatePersonRequest) (response *CreatePersonResponse, err error) {
    return c.CreatePersonWithContext(context.Background(), request)
}

// CreatePerson
// 创建人员
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CreatePersonWithContext(ctx context.Context, request *CreatePersonRequest) (response *CreatePersonResponse, err error) {
    if request == nil {
        request = NewCreatePersonRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreatePerson require credential")
    }

    request.SetContext(ctx)
    
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

// CreateVocab
// 创建词汇
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_RESOURCENAMEDUPLICATE = "InvalidParameter.ResourceNameDuplicate"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateVocab(request *CreateVocabRequest) (response *CreateVocabResponse, err error) {
    return c.CreateVocabWithContext(context.Background(), request)
}

// CreateVocab
// 创建词汇
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_RESOURCENAMEDUPLICATE = "InvalidParameter.ResourceNameDuplicate"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateVocabWithContext(ctx context.Context, request *CreateVocabRequest) (response *CreateVocabResponse, err error) {
    if request == nil {
        request = NewCreateVocabRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateVocab require credential")
    }

    request.SetContext(ctx)
    
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

// CreateVocabLib
// 建立词汇库
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_RESOURCENAMEDUPLICATE = "InvalidParameter.ResourceNameDuplicate"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateVocabLib(request *CreateVocabLibRequest) (response *CreateVocabLibResponse, err error) {
    return c.CreateVocabLibWithContext(context.Background(), request)
}

// CreateVocabLib
// 建立词汇库
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_RESOURCENAMEDUPLICATE = "InvalidParameter.ResourceNameDuplicate"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateVocabLibWithContext(ctx context.Context, request *CreateVocabLibRequest) (response *CreateVocabLibResponse, err error) {
    if request == nil {
        request = NewCreateVocabLibRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateVocabLib require credential")
    }

    request.SetContext(ctx)
    
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

// DeleteFace
// 删除人脸
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteFace(request *DeleteFaceRequest) (response *DeleteFaceResponse, err error) {
    return c.DeleteFaceWithContext(context.Background(), request)
}

// DeleteFace
// 删除人脸
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteFaceWithContext(ctx context.Context, request *DeleteFaceRequest) (response *DeleteFaceResponse, err error) {
    if request == nil {
        request = NewDeleteFaceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteFace require credential")
    }

    request.SetContext(ctx)
    
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

// DeleteLibrary
// 删除人员库
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteLibrary(request *DeleteLibraryRequest) (response *DeleteLibraryResponse, err error) {
    return c.DeleteLibraryWithContext(context.Background(), request)
}

// DeleteLibrary
// 删除人员库
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteLibraryWithContext(ctx context.Context, request *DeleteLibraryRequest) (response *DeleteLibraryResponse, err error) {
    if request == nil {
        request = NewDeleteLibraryRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteLibrary require credential")
    }

    request.SetContext(ctx)
    
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

// DeletePerson
// 删除人员
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeletePerson(request *DeletePersonRequest) (response *DeletePersonResponse, err error) {
    return c.DeletePersonWithContext(context.Background(), request)
}

// DeletePerson
// 删除人员
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeletePersonWithContext(ctx context.Context, request *DeletePersonRequest) (response *DeletePersonResponse, err error) {
    if request == nil {
        request = NewDeletePersonRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeletePerson require credential")
    }

    request.SetContext(ctx)
    
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

// DeleteVocab
// 删除词汇
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_RESOURCENAMEDUPLICATE = "InvalidParameter.ResourceNameDuplicate"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteVocab(request *DeleteVocabRequest) (response *DeleteVocabResponse, err error) {
    return c.DeleteVocabWithContext(context.Background(), request)
}

// DeleteVocab
// 删除词汇
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_RESOURCENAMEDUPLICATE = "InvalidParameter.ResourceNameDuplicate"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteVocabWithContext(ctx context.Context, request *DeleteVocabRequest) (response *DeleteVocabResponse, err error) {
    if request == nil {
        request = NewDeleteVocabRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteVocab require credential")
    }

    request.SetContext(ctx)
    
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

// DeleteVocabLib
// 删除词汇库
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_RESOURCENAMEDUPLICATE = "InvalidParameter.ResourceNameDuplicate"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteVocabLib(request *DeleteVocabLibRequest) (response *DeleteVocabLibResponse, err error) {
    return c.DeleteVocabLibWithContext(context.Background(), request)
}

// DeleteVocabLib
// 删除词汇库
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_RESOURCENAMEDUPLICATE = "InvalidParameter.ResourceNameDuplicate"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteVocabLibWithContext(ctx context.Context, request *DeleteVocabLibRequest) (response *DeleteVocabLibResponse, err error) {
    if request == nil {
        request = NewDeleteVocabLibRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteVocabLib require credential")
    }

    request.SetContext(ctx)
    
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

// DescribeAITaskResult
// 获取标准化接口任务结果
//
// 可能返回的错误码:
//  INTERNALERROR_STATUSQUERYFACELIBFAILED = "InternalError.StatusQueryFaceLibFailed"
//  INVALIDPARAMETER_STATUSJOBFAILED = "InvalidParameter.StatusJobFailed"
//  INVALIDPARAMETER_STATUSJOBUNFINISHED = "InvalidParameter.StatusJobUnfinished"
//  RESOURCENOTFOUND_HLJOBIDNOTFOUND = "ResourceNotFound.HLJobIdNotFound"
func (c *Client) DescribeAITaskResult(request *DescribeAITaskResultRequest) (response *DescribeAITaskResultResponse, err error) {
    return c.DescribeAITaskResultWithContext(context.Background(), request)
}

// DescribeAITaskResult
// 获取标准化接口任务结果
//
// 可能返回的错误码:
//  INTERNALERROR_STATUSQUERYFACELIBFAILED = "InternalError.StatusQueryFaceLibFailed"
//  INVALIDPARAMETER_STATUSJOBFAILED = "InvalidParameter.StatusJobFailed"
//  INVALIDPARAMETER_STATUSJOBUNFINISHED = "InvalidParameter.StatusJobUnfinished"
//  RESOURCENOTFOUND_HLJOBIDNOTFOUND = "ResourceNotFound.HLJobIdNotFound"
func (c *Client) DescribeAITaskResultWithContext(ctx context.Context, request *DescribeAITaskResultRequest) (response *DescribeAITaskResultResponse, err error) {
    if request == nil {
        request = NewDescribeAITaskResultRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAITaskResult require credential")
    }

    request.SetContext(ctx)
    
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

// DescribeAttendanceResult
// 人脸考勤查询结果
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FACENOTEXISTED = "InvalidParameter.FaceNotExisted"
//  INVALIDPARAMETER_FILECONTENTEMPTY = "InvalidParameter.FileContentEmpty"
//  INVALIDPARAMETER_INVALIDRESOURCENAME = "InvalidParameter.InvalidResourceName"
//  INVALIDPARAMETER_JOBNOTVALID = "InvalidParameter.JobNotValid"
//  INVALIDPARAMETER_NOTATTENDANCEJOB = "InvalidParameter.NotAttendanceJob"
//  INVALIDPARAMETER_PERSONNOTREGISTERED = "InvalidParameter.PersonNotRegistered"
//  INVALIDPARAMETER_RESOURCENAMEDUPLICATE = "InvalidParameter.ResourceNameDuplicate"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAttendanceResult(request *DescribeAttendanceResultRequest) (response *DescribeAttendanceResultResponse, err error) {
    return c.DescribeAttendanceResultWithContext(context.Background(), request)
}

// DescribeAttendanceResult
// 人脸考勤查询结果
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FACENOTEXISTED = "InvalidParameter.FaceNotExisted"
//  INVALIDPARAMETER_FILECONTENTEMPTY = "InvalidParameter.FileContentEmpty"
//  INVALIDPARAMETER_INVALIDRESOURCENAME = "InvalidParameter.InvalidResourceName"
//  INVALIDPARAMETER_JOBNOTVALID = "InvalidParameter.JobNotValid"
//  INVALIDPARAMETER_NOTATTENDANCEJOB = "InvalidParameter.NotAttendanceJob"
//  INVALIDPARAMETER_PERSONNOTREGISTERED = "InvalidParameter.PersonNotRegistered"
//  INVALIDPARAMETER_RESOURCENAMEDUPLICATE = "InvalidParameter.ResourceNameDuplicate"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAttendanceResultWithContext(ctx context.Context, request *DescribeAttendanceResultRequest) (response *DescribeAttendanceResultResponse, err error) {
    if request == nil {
        request = NewDescribeAttendanceResultRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAttendanceResult require credential")
    }

    request.SetContext(ctx)
    
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

// DescribeAudioTask
// 音频评估任务信息查询接口，异步查询客户提交的请求的结果。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAudioTask(request *DescribeAudioTaskRequest) (response *DescribeAudioTaskResponse, err error) {
    return c.DescribeAudioTaskWithContext(context.Background(), request)
}

// DescribeAudioTask
// 音频评估任务信息查询接口，异步查询客户提交的请求的结果。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAudioTaskWithContext(ctx context.Context, request *DescribeAudioTaskRequest) (response *DescribeAudioTaskResponse, err error) {
    if request == nil {
        request = NewDescribeAudioTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAudioTask require credential")
    }

    request.SetContext(ctx)
    
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

// DescribeConversationTask
// 音频对话任务评估任务信息查询接口，异步查询客户提交的请求的结果。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeConversationTask(request *DescribeConversationTaskRequest) (response *DescribeConversationTaskResponse, err error) {
    return c.DescribeConversationTaskWithContext(context.Background(), request)
}

// DescribeConversationTask
// 音频对话任务评估任务信息查询接口，异步查询客户提交的请求的结果。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeConversationTaskWithContext(ctx context.Context, request *DescribeConversationTaskRequest) (response *DescribeConversationTaskResponse, err error) {
    if request == nil {
        request = NewDescribeConversationTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeConversationTask require credential")
    }

    request.SetContext(ctx)
    
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

// DescribeHighlightResult
// 视频精彩集锦结果查询接口，异步查询客户提交的请求的结果。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_JOBCANCELED = "InternalError.JobCanceled"
//  INTERNALERROR_JOBREACHMAXRETRYTIMES = "InternalError.JobReachMaxRetryTimes"
//  INTERNALERROR_STATUSQUERYFACELIBFAILED = "InternalError.StatusQueryFaceLibFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_BODYRETRIEVALFAILED = "InvalidParameter.BodyRetrievalFailed"
//  INVALIDPARAMETER_CANNOTFINDFACEINPICTURE = "InvalidParameter.CannotFindFaceInPicture"
//  INVALIDPARAMETER_CLIPPINGFRAMEFAILED = "InvalidParameter.ClippingFrameFailed"
//  INVALIDPARAMETER_COMPUTATIONALSIMILARITYFAILED = "InvalidParameter.ComputationalSimilarityFailed"
//  INVALIDPARAMETER_CREATEFACEFAILED = "InvalidParameter.CreateFaceFailed"
//  INVALIDPARAMETER_CREATEPERSONFAILED = "InvalidParameter.CreatePersonFailed"
//  INVALIDPARAMETER_FACEDETECTEDFAILED = "InvalidParameter.FaceDetectedFailed"
//  INVALIDPARAMETER_FACENOTEXIST = "InvalidParameter.FaceNotExist"
//  INVALIDPARAMETER_FACENOTEXISTED = "InvalidParameter.FaceNotExisted"
//  INVALIDPARAMETER_FACENUMBEREXCEEDLIMITOFSINGLEDETECTION = "InvalidParameter.FaceNumberExceedLimitOfSingleDetection"
//  INVALIDPARAMETER_FACENUMBERLIMIT = "InvalidParameter.FaceNumberLimit"
//  INVALIDPARAMETER_FACEREGISTRATIONFAILED = "InvalidParameter.FaceRegistrationFailed"
//  INVALIDPARAMETER_FACERETRIEVALFAILED = "InvalidParameter.FaceRetrievalFailed"
//  INVALIDPARAMETER_FACETOOSMALL = "InvalidParameter.FaceTooSmall"
//  INVALIDPARAMETER_FAILEDTODECODEIMAGE = "InvalidParameter.FailedToDecodeImage"
//  INVALIDPARAMETER_FEATUREEXTRACTIONFAILED = "InvalidParameter.FeatureExtractionFailed"
//  INVALIDPARAMETER_FILECONTENTEMPTY = "InvalidParameter.FileContentEmpty"
//  INVALIDPARAMETER_IMAGEDECODEFAILED = "InvalidParameter.ImageDecodeFailed"
//  INVALIDPARAMETER_IMAGETOOLARGE = "InvalidParameter.ImageTooLarge"
//  INVALIDPARAMETER_IMAGETOOSMALL = "InvalidParameter.ImageTooSmall"
//  INVALIDPARAMETER_INVALIDFILETYPE = "InvalidParameter.InvalidFileType"
//  INVALIDPARAMETER_INVALIDLANG = "InvalidParameter.InvalidLang"
//  INVALIDPARAMETER_INVALIDRESOURCENAME = "InvalidParameter.InvalidResourceName"
//  INVALIDPARAMETER_INVALIDSEQID = "InvalidParameter.InvalidSeqId"
//  INVALIDPARAMETER_INVALIDURL = "InvalidParameter.InvalidUrl"
//  INVALIDPARAMETER_JOBNOTVALID = "InvalidParameter.JobNotValid"
//  INVALIDPARAMETER_LIBRARYNOTEXISTED = "InvalidParameter.LibraryNotExisted"
//  INVALIDPARAMETER_LOWQUALITYPICTURE = "InvalidParameter.LowQualityPicture"
//  INVALIDPARAMETER_NOSUCHSERVICE = "InvalidParameter.NoSuchService"
//  INVALIDPARAMETER_NOTATTENDANCEJOB = "InvalidParameter.NotAttendanceJob"
//  INVALIDPARAMETER_NOTVALIDPITCH = "InvalidParameter.NotValidPitch"
//  INVALIDPARAMETER_NOTVALIDROW = "InvalidParameter.NotValidRow"
//  INVALIDPARAMETER_NOTVALIDYAW = "InvalidParameter.NotValidYaw"
//  INVALIDPARAMETER_PERSONEXISTED = "InvalidParameter.PersonExisted"
//  INVALIDPARAMETER_PERSONIDNOTVALID = "InvalidParameter.PersonIdNotValid"
//  INVALIDPARAMETER_PERSONNOTEXIST = "InvalidParameter.PersonNotExist"
//  INVALIDPARAMETER_PERSONNOTEXISTED = "InvalidParameter.PersonNotExisted"
//  INVALIDPARAMETER_PERSONNOTREGISTERED = "InvalidParameter.PersonNotRegistered"
//  INVALIDPARAMETER_PROCESSUNREADY = "InvalidParameter.ProcessUnReady"
//  INVALIDPARAMETER_RESOURCENAMEDUPLICATE = "InvalidParameter.ResourceNameDuplicate"
//  INVALIDPARAMETER_STATUSJOBFAILED = "InvalidParameter.StatusJobFailed"
//  INVALIDPARAMETER_STATUSJOBNOTFOUND = "InvalidParameter.StatusJobNotFound"
//  INVALIDPARAMETER_STATUSJOBUNFINISHED = "InvalidParameter.StatusJobUnfinished"
//  INVALIDPARAMETER_STATUSNOCAMERAFOUND = "InvalidParameter.StatusNoCameraFound"
//  INVALIDPARAMETER_TOOMANYFACE = "InvalidParameter.TooManyFace"
//  INVALIDPARAMETER_TRACKINGFAILED = "InvalidParameter.TrackingFailed"
//  LIMITEXCEEDED_STATUSCAMERACOUNTOVERMAX = "LimitExceeded.StatusCameraCountOverMax"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_HLJOBIDNOTFOUND = "ResourceNotFound.HLJobIdNotFound"
//  RESOURCEUNAVAILABLE_STATUSQUERYSTREAMFAILED = "ResourceUnavailable.StatusQueryStreamFailed"
func (c *Client) DescribeHighlightResult(request *DescribeHighlightResultRequest) (response *DescribeHighlightResultResponse, err error) {
    return c.DescribeHighlightResultWithContext(context.Background(), request)
}

// DescribeHighlightResult
// 视频精彩集锦结果查询接口，异步查询客户提交的请求的结果。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_JOBCANCELED = "InternalError.JobCanceled"
//  INTERNALERROR_JOBREACHMAXRETRYTIMES = "InternalError.JobReachMaxRetryTimes"
//  INTERNALERROR_STATUSQUERYFACELIBFAILED = "InternalError.StatusQueryFaceLibFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_BODYRETRIEVALFAILED = "InvalidParameter.BodyRetrievalFailed"
//  INVALIDPARAMETER_CANNOTFINDFACEINPICTURE = "InvalidParameter.CannotFindFaceInPicture"
//  INVALIDPARAMETER_CLIPPINGFRAMEFAILED = "InvalidParameter.ClippingFrameFailed"
//  INVALIDPARAMETER_COMPUTATIONALSIMILARITYFAILED = "InvalidParameter.ComputationalSimilarityFailed"
//  INVALIDPARAMETER_CREATEFACEFAILED = "InvalidParameter.CreateFaceFailed"
//  INVALIDPARAMETER_CREATEPERSONFAILED = "InvalidParameter.CreatePersonFailed"
//  INVALIDPARAMETER_FACEDETECTEDFAILED = "InvalidParameter.FaceDetectedFailed"
//  INVALIDPARAMETER_FACENOTEXIST = "InvalidParameter.FaceNotExist"
//  INVALIDPARAMETER_FACENOTEXISTED = "InvalidParameter.FaceNotExisted"
//  INVALIDPARAMETER_FACENUMBEREXCEEDLIMITOFSINGLEDETECTION = "InvalidParameter.FaceNumberExceedLimitOfSingleDetection"
//  INVALIDPARAMETER_FACENUMBERLIMIT = "InvalidParameter.FaceNumberLimit"
//  INVALIDPARAMETER_FACEREGISTRATIONFAILED = "InvalidParameter.FaceRegistrationFailed"
//  INVALIDPARAMETER_FACERETRIEVALFAILED = "InvalidParameter.FaceRetrievalFailed"
//  INVALIDPARAMETER_FACETOOSMALL = "InvalidParameter.FaceTooSmall"
//  INVALIDPARAMETER_FAILEDTODECODEIMAGE = "InvalidParameter.FailedToDecodeImage"
//  INVALIDPARAMETER_FEATUREEXTRACTIONFAILED = "InvalidParameter.FeatureExtractionFailed"
//  INVALIDPARAMETER_FILECONTENTEMPTY = "InvalidParameter.FileContentEmpty"
//  INVALIDPARAMETER_IMAGEDECODEFAILED = "InvalidParameter.ImageDecodeFailed"
//  INVALIDPARAMETER_IMAGETOOLARGE = "InvalidParameter.ImageTooLarge"
//  INVALIDPARAMETER_IMAGETOOSMALL = "InvalidParameter.ImageTooSmall"
//  INVALIDPARAMETER_INVALIDFILETYPE = "InvalidParameter.InvalidFileType"
//  INVALIDPARAMETER_INVALIDLANG = "InvalidParameter.InvalidLang"
//  INVALIDPARAMETER_INVALIDRESOURCENAME = "InvalidParameter.InvalidResourceName"
//  INVALIDPARAMETER_INVALIDSEQID = "InvalidParameter.InvalidSeqId"
//  INVALIDPARAMETER_INVALIDURL = "InvalidParameter.InvalidUrl"
//  INVALIDPARAMETER_JOBNOTVALID = "InvalidParameter.JobNotValid"
//  INVALIDPARAMETER_LIBRARYNOTEXISTED = "InvalidParameter.LibraryNotExisted"
//  INVALIDPARAMETER_LOWQUALITYPICTURE = "InvalidParameter.LowQualityPicture"
//  INVALIDPARAMETER_NOSUCHSERVICE = "InvalidParameter.NoSuchService"
//  INVALIDPARAMETER_NOTATTENDANCEJOB = "InvalidParameter.NotAttendanceJob"
//  INVALIDPARAMETER_NOTVALIDPITCH = "InvalidParameter.NotValidPitch"
//  INVALIDPARAMETER_NOTVALIDROW = "InvalidParameter.NotValidRow"
//  INVALIDPARAMETER_NOTVALIDYAW = "InvalidParameter.NotValidYaw"
//  INVALIDPARAMETER_PERSONEXISTED = "InvalidParameter.PersonExisted"
//  INVALIDPARAMETER_PERSONIDNOTVALID = "InvalidParameter.PersonIdNotValid"
//  INVALIDPARAMETER_PERSONNOTEXIST = "InvalidParameter.PersonNotExist"
//  INVALIDPARAMETER_PERSONNOTEXISTED = "InvalidParameter.PersonNotExisted"
//  INVALIDPARAMETER_PERSONNOTREGISTERED = "InvalidParameter.PersonNotRegistered"
//  INVALIDPARAMETER_PROCESSUNREADY = "InvalidParameter.ProcessUnReady"
//  INVALIDPARAMETER_RESOURCENAMEDUPLICATE = "InvalidParameter.ResourceNameDuplicate"
//  INVALIDPARAMETER_STATUSJOBFAILED = "InvalidParameter.StatusJobFailed"
//  INVALIDPARAMETER_STATUSJOBNOTFOUND = "InvalidParameter.StatusJobNotFound"
//  INVALIDPARAMETER_STATUSJOBUNFINISHED = "InvalidParameter.StatusJobUnfinished"
//  INVALIDPARAMETER_STATUSNOCAMERAFOUND = "InvalidParameter.StatusNoCameraFound"
//  INVALIDPARAMETER_TOOMANYFACE = "InvalidParameter.TooManyFace"
//  INVALIDPARAMETER_TRACKINGFAILED = "InvalidParameter.TrackingFailed"
//  LIMITEXCEEDED_STATUSCAMERACOUNTOVERMAX = "LimitExceeded.StatusCameraCountOverMax"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_HLJOBIDNOTFOUND = "ResourceNotFound.HLJobIdNotFound"
//  RESOURCEUNAVAILABLE_STATUSQUERYSTREAMFAILED = "ResourceUnavailable.StatusQueryStreamFailed"
func (c *Client) DescribeHighlightResultWithContext(ctx context.Context, request *DescribeHighlightResultRequest) (response *DescribeHighlightResultResponse, err error) {
    if request == nil {
        request = NewDescribeHighlightResultRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeHighlightResult require credential")
    }

    request.SetContext(ctx)
    
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

// DescribeImageTask
// 拉取任务详情
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeImageTask(request *DescribeImageTaskRequest) (response *DescribeImageTaskResponse, err error) {
    return c.DescribeImageTaskWithContext(context.Background(), request)
}

// DescribeImageTask
// 拉取任务详情
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeImageTaskWithContext(ctx context.Context, request *DescribeImageTaskRequest) (response *DescribeImageTaskResponse, err error) {
    if request == nil {
        request = NewDescribeImageTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeImageTask require credential")
    }

    request.SetContext(ctx)
    
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

// DescribeImageTaskStatistic
// 获取图像任务统计信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeImageTaskStatistic(request *DescribeImageTaskStatisticRequest) (response *DescribeImageTaskStatisticResponse, err error) {
    return c.DescribeImageTaskStatisticWithContext(context.Background(), request)
}

// DescribeImageTaskStatistic
// 获取图像任务统计信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeImageTaskStatisticWithContext(ctx context.Context, request *DescribeImageTaskStatisticRequest) (response *DescribeImageTaskStatisticResponse, err error) {
    if request == nil {
        request = NewDescribeImageTaskStatisticRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeImageTaskStatistic require credential")
    }

    request.SetContext(ctx)
    
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

// DescribeLibraries
// 获取人员库列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeLibraries(request *DescribeLibrariesRequest) (response *DescribeLibrariesResponse, err error) {
    return c.DescribeLibrariesWithContext(context.Background(), request)
}

// DescribeLibraries
// 获取人员库列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeLibrariesWithContext(ctx context.Context, request *DescribeLibrariesRequest) (response *DescribeLibrariesResponse, err error) {
    if request == nil {
        request = NewDescribeLibrariesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLibraries require credential")
    }

    request.SetContext(ctx)
    
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

// DescribePerson
// 获取人员详情
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribePerson(request *DescribePersonRequest) (response *DescribePersonResponse, err error) {
    return c.DescribePersonWithContext(context.Background(), request)
}

// DescribePerson
// 获取人员详情
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribePersonWithContext(ctx context.Context, request *DescribePersonRequest) (response *DescribePersonResponse, err error) {
    if request == nil {
        request = NewDescribePersonRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePerson require credential")
    }

    request.SetContext(ctx)
    
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

// DescribePersons
// 拉取人员列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribePersons(request *DescribePersonsRequest) (response *DescribePersonsResponse, err error) {
    return c.DescribePersonsWithContext(context.Background(), request)
}

// DescribePersons
// 拉取人员列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribePersonsWithContext(ctx context.Context, request *DescribePersonsRequest) (response *DescribePersonsResponse, err error) {
    if request == nil {
        request = NewDescribePersonsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePersons require credential")
    }

    request.SetContext(ctx)
    
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

// DescribeVocab
// 查询词汇
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_RESOURCENAMEDUPLICATE = "InvalidParameter.ResourceNameDuplicate"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeVocab(request *DescribeVocabRequest) (response *DescribeVocabResponse, err error) {
    return c.DescribeVocabWithContext(context.Background(), request)
}

// DescribeVocab
// 查询词汇
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_RESOURCENAMEDUPLICATE = "InvalidParameter.ResourceNameDuplicate"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeVocabWithContext(ctx context.Context, request *DescribeVocabRequest) (response *DescribeVocabResponse, err error) {
    if request == nil {
        request = NewDescribeVocabRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVocab require credential")
    }

    request.SetContext(ctx)
    
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

// DescribeVocabLib
// 查询词汇库
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_RESOURCENAMEDUPLICATE = "InvalidParameter.ResourceNameDuplicate"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeVocabLib(request *DescribeVocabLibRequest) (response *DescribeVocabLibResponse, err error) {
    return c.DescribeVocabLibWithContext(context.Background(), request)
}

// DescribeVocabLib
// 查询词汇库
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_RESOURCENAMEDUPLICATE = "InvalidParameter.ResourceNameDuplicate"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeVocabLibWithContext(ctx context.Context, request *DescribeVocabLibRequest) (response *DescribeVocabLibResponse, err error) {
    if request == nil {
        request = NewDescribeVocabLibRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVocabLib require credential")
    }

    request.SetContext(ctx)
    
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

// ModifyLibrary
// 修改人员库信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyLibrary(request *ModifyLibraryRequest) (response *ModifyLibraryResponse, err error) {
    return c.ModifyLibraryWithContext(context.Background(), request)
}

// ModifyLibrary
// 修改人员库信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyLibraryWithContext(ctx context.Context, request *ModifyLibraryRequest) (response *ModifyLibraryResponse, err error) {
    if request == nil {
        request = NewModifyLibraryRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyLibrary require credential")
    }

    request.SetContext(ctx)
    
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

// ModifyPerson
// 修改人员信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyPerson(request *ModifyPersonRequest) (response *ModifyPersonResponse, err error) {
    return c.ModifyPersonWithContext(context.Background(), request)
}

// ModifyPerson
// 修改人员信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyPersonWithContext(ctx context.Context, request *ModifyPersonRequest) (response *ModifyPersonResponse, err error) {
    if request == nil {
        request = NewModifyPersonRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyPerson require credential")
    }

    request.SetContext(ctx)
    
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

// SubmitAudioTask
// 音频任务提交接口
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) SubmitAudioTask(request *SubmitAudioTaskRequest) (response *SubmitAudioTaskResponse, err error) {
    return c.SubmitAudioTaskWithContext(context.Background(), request)
}

// SubmitAudioTask
// 音频任务提交接口
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) SubmitAudioTaskWithContext(ctx context.Context, request *SubmitAudioTaskRequest) (response *SubmitAudioTaskResponse, err error) {
    if request == nil {
        request = NewSubmitAudioTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SubmitAudioTask require credential")
    }

    request.SetContext(ctx)
    
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

// SubmitCheckAttendanceTask
// 提交人员考勤任务，支持包括点播和直播资源；支持通过DescribeAttendanceResult查询结果，也支持通过NoticeUrl设置考勤回调结果，回调结果结构如下：
//
// ##### 回调事件结构
//
//  | 参数名称 | 类型 | 描述 | 
//
//  | ----  | ---  | ------  |
//
//  | jobid | Integer | 任务ID | 
//
//  | person_info | array of PersonInfo | 识别到的人员列表 | 
//
// #####子结构PersonInfo
//
//  | 参数名称 | 类型 | 描述 | 
//
//  | ----  | ---  | ------  |
//
//  | traceid | String | 可用于区分同一路视频流下的不同陌生人 | 
//
//  | personid | String | 识别到的人员ID，如果是陌生人则返回空串 | 
//
//  | libid | String | 识别到的人员所在的库ID，如果是陌生人则返回空串 | 
//
//  | timestamp | uint64 | 识别到人脸的绝对时间戳，单位ms | 
//
//  | image_url | string | 识别到人脸的事件抓图的下载地址，不长期保存，需要请及时下载 | 
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FACENOTEXISTED = "InvalidParameter.FaceNotExisted"
//  INVALIDPARAMETER_FILECONTENTEMPTY = "InvalidParameter.FileContentEmpty"
//  INVALIDPARAMETER_INVALIDRESOURCENAME = "InvalidParameter.InvalidResourceName"
//  INVALIDPARAMETER_JOBNOTVALID = "InvalidParameter.JobNotValid"
//  INVALIDPARAMETER_NOTATTENDANCEJOB = "InvalidParameter.NotAttendanceJob"
//  INVALIDPARAMETER_PERSONNOTREGISTERED = "InvalidParameter.PersonNotRegistered"
//  INVALIDPARAMETER_RESOURCENAMEDUPLICATE = "InvalidParameter.ResourceNameDuplicate"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) SubmitCheckAttendanceTask(request *SubmitCheckAttendanceTaskRequest) (response *SubmitCheckAttendanceTaskResponse, err error) {
    return c.SubmitCheckAttendanceTaskWithContext(context.Background(), request)
}

// SubmitCheckAttendanceTask
// 提交人员考勤任务，支持包括点播和直播资源；支持通过DescribeAttendanceResult查询结果，也支持通过NoticeUrl设置考勤回调结果，回调结果结构如下：
//
// ##### 回调事件结构
//
//  | 参数名称 | 类型 | 描述 | 
//
//  | ----  | ---  | ------  |
//
//  | jobid | Integer | 任务ID | 
//
//  | person_info | array of PersonInfo | 识别到的人员列表 | 
//
// #####子结构PersonInfo
//
//  | 参数名称 | 类型 | 描述 | 
//
//  | ----  | ---  | ------  |
//
//  | traceid | String | 可用于区分同一路视频流下的不同陌生人 | 
//
//  | personid | String | 识别到的人员ID，如果是陌生人则返回空串 | 
//
//  | libid | String | 识别到的人员所在的库ID，如果是陌生人则返回空串 | 
//
//  | timestamp | uint64 | 识别到人脸的绝对时间戳，单位ms | 
//
//  | image_url | string | 识别到人脸的事件抓图的下载地址，不长期保存，需要请及时下载 | 
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FACENOTEXISTED = "InvalidParameter.FaceNotExisted"
//  INVALIDPARAMETER_FILECONTENTEMPTY = "InvalidParameter.FileContentEmpty"
//  INVALIDPARAMETER_INVALIDRESOURCENAME = "InvalidParameter.InvalidResourceName"
//  INVALIDPARAMETER_JOBNOTVALID = "InvalidParameter.JobNotValid"
//  INVALIDPARAMETER_NOTATTENDANCEJOB = "InvalidParameter.NotAttendanceJob"
//  INVALIDPARAMETER_PERSONNOTREGISTERED = "InvalidParameter.PersonNotRegistered"
//  INVALIDPARAMETER_RESOURCENAMEDUPLICATE = "InvalidParameter.ResourceNameDuplicate"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) SubmitCheckAttendanceTaskWithContext(ctx context.Context, request *SubmitCheckAttendanceTaskRequest) (response *SubmitCheckAttendanceTaskResponse, err error) {
    if request == nil {
        request = NewSubmitCheckAttendanceTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SubmitCheckAttendanceTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewSubmitCheckAttendanceTaskResponse()
    err = c.Send(request, response)
    return
}

func NewSubmitCheckAttendanceTaskPlusRequest() (request *SubmitCheckAttendanceTaskPlusRequest) {
    request = &SubmitCheckAttendanceTaskPlusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tci", APIVersion, "SubmitCheckAttendanceTaskPlus")
    
    
    return
}

func NewSubmitCheckAttendanceTaskPlusResponse() (response *SubmitCheckAttendanceTaskPlusResponse) {
    response = &SubmitCheckAttendanceTaskPlusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// SubmitCheckAttendanceTaskPlus
// 支持多路视频流，提交高级人员考勤任务
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_FACENOTEXISTED = "InvalidParameter.FaceNotExisted"
//  INVALIDPARAMETER_FILECONTENTEMPTY = "InvalidParameter.FileContentEmpty"
//  INVALIDPARAMETER_INVALIDRESOURCENAME = "InvalidParameter.InvalidResourceName"
//  INVALIDPARAMETER_JOBNOTVALID = "InvalidParameter.JobNotValid"
//  INVALIDPARAMETER_NOTATTENDANCEJOB = "InvalidParameter.NotAttendanceJob"
//  INVALIDPARAMETER_PERSONNOTREGISTERED = "InvalidParameter.PersonNotRegistered"
//  INVALIDPARAMETER_RESOURCENAMEDUPLICATE = "InvalidParameter.ResourceNameDuplicate"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) SubmitCheckAttendanceTaskPlus(request *SubmitCheckAttendanceTaskPlusRequest) (response *SubmitCheckAttendanceTaskPlusResponse, err error) {
    return c.SubmitCheckAttendanceTaskPlusWithContext(context.Background(), request)
}

// SubmitCheckAttendanceTaskPlus
// 支持多路视频流，提交高级人员考勤任务
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_FACENOTEXISTED = "InvalidParameter.FaceNotExisted"
//  INVALIDPARAMETER_FILECONTENTEMPTY = "InvalidParameter.FileContentEmpty"
//  INVALIDPARAMETER_INVALIDRESOURCENAME = "InvalidParameter.InvalidResourceName"
//  INVALIDPARAMETER_JOBNOTVALID = "InvalidParameter.JobNotValid"
//  INVALIDPARAMETER_NOTATTENDANCEJOB = "InvalidParameter.NotAttendanceJob"
//  INVALIDPARAMETER_PERSONNOTREGISTERED = "InvalidParameter.PersonNotRegistered"
//  INVALIDPARAMETER_RESOURCENAMEDUPLICATE = "InvalidParameter.ResourceNameDuplicate"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) SubmitCheckAttendanceTaskPlusWithContext(ctx context.Context, request *SubmitCheckAttendanceTaskPlusRequest) (response *SubmitCheckAttendanceTaskPlusResponse, err error) {
    if request == nil {
        request = NewSubmitCheckAttendanceTaskPlusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SubmitCheckAttendanceTaskPlus require credential")
    }

    request.SetContext(ctx)
    
    response = NewSubmitCheckAttendanceTaskPlusResponse()
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

// SubmitConversationTask
// 对话任务分析接口
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) SubmitConversationTask(request *SubmitConversationTaskRequest) (response *SubmitConversationTaskResponse, err error) {
    return c.SubmitConversationTaskWithContext(context.Background(), request)
}

// SubmitConversationTask
// 对话任务分析接口
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) SubmitConversationTaskWithContext(ctx context.Context, request *SubmitConversationTaskRequest) (response *SubmitConversationTaskResponse, err error) {
    if request == nil {
        request = NewSubmitConversationTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SubmitConversationTask require credential")
    }

    request.SetContext(ctx)
    
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

// SubmitDoubleVideoHighlights
// 发起双路视频生成精彩集锦接口。该接口可以通过客户传入的学生音视频及老师视频两路Url，自动生成一堂课程的精彩集锦。需要通过DescribeHighlightResult
//
// 接口获取生成结果。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_CLIPPINGFRAMEFAILED = "InvalidParameter.ClippingFrameFailed"
//  INVALIDPARAMETER_FACEDETECTEDFAILED = "InvalidParameter.FaceDetectedFailed"
//  INVALIDPARAMETER_FACENOTEXISTED = "InvalidParameter.FaceNotExisted"
//  INVALIDPARAMETER_FACEREGISTRATIONFAILED = "InvalidParameter.FaceRegistrationFailed"
//  INVALIDPARAMETER_FAILEDTODECODEIMAGE = "InvalidParameter.FailedToDecodeImage"
//  INVALIDPARAMETER_FEATUREEXTRACTIONFAILED = "InvalidParameter.FeatureExtractionFailed"
//  INVALIDPARAMETER_FILECONTENTEMPTY = "InvalidParameter.FileContentEmpty"
//  INVALIDPARAMETER_IMAGEDECODEFAILED = "InvalidParameter.ImageDecodeFailed"
//  INVALIDPARAMETER_INVALIDRESOURCENAME = "InvalidParameter.InvalidResourceName"
//  INVALIDPARAMETER_JOBNOTVALID = "InvalidParameter.JobNotValid"
//  INVALIDPARAMETER_NOSUCHSERVICE = "InvalidParameter.NoSuchService"
//  INVALIDPARAMETER_NOTATTENDANCEJOB = "InvalidParameter.NotAttendanceJob"
//  INVALIDPARAMETER_PERSONNOTREGISTERED = "InvalidParameter.PersonNotRegistered"
//  INVALIDPARAMETER_RESOURCENAMEDUPLICATE = "InvalidParameter.ResourceNameDuplicate"
func (c *Client) SubmitDoubleVideoHighlights(request *SubmitDoubleVideoHighlightsRequest) (response *SubmitDoubleVideoHighlightsResponse, err error) {
    return c.SubmitDoubleVideoHighlightsWithContext(context.Background(), request)
}

// SubmitDoubleVideoHighlights
// 发起双路视频生成精彩集锦接口。该接口可以通过客户传入的学生音视频及老师视频两路Url，自动生成一堂课程的精彩集锦。需要通过DescribeHighlightResult
//
// 接口获取生成结果。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_CLIPPINGFRAMEFAILED = "InvalidParameter.ClippingFrameFailed"
//  INVALIDPARAMETER_FACEDETECTEDFAILED = "InvalidParameter.FaceDetectedFailed"
//  INVALIDPARAMETER_FACENOTEXISTED = "InvalidParameter.FaceNotExisted"
//  INVALIDPARAMETER_FACEREGISTRATIONFAILED = "InvalidParameter.FaceRegistrationFailed"
//  INVALIDPARAMETER_FAILEDTODECODEIMAGE = "InvalidParameter.FailedToDecodeImage"
//  INVALIDPARAMETER_FEATUREEXTRACTIONFAILED = "InvalidParameter.FeatureExtractionFailed"
//  INVALIDPARAMETER_FILECONTENTEMPTY = "InvalidParameter.FileContentEmpty"
//  INVALIDPARAMETER_IMAGEDECODEFAILED = "InvalidParameter.ImageDecodeFailed"
//  INVALIDPARAMETER_INVALIDRESOURCENAME = "InvalidParameter.InvalidResourceName"
//  INVALIDPARAMETER_JOBNOTVALID = "InvalidParameter.JobNotValid"
//  INVALIDPARAMETER_NOSUCHSERVICE = "InvalidParameter.NoSuchService"
//  INVALIDPARAMETER_NOTATTENDANCEJOB = "InvalidParameter.NotAttendanceJob"
//  INVALIDPARAMETER_PERSONNOTREGISTERED = "InvalidParameter.PersonNotRegistered"
//  INVALIDPARAMETER_RESOURCENAMEDUPLICATE = "InvalidParameter.ResourceNameDuplicate"
func (c *Client) SubmitDoubleVideoHighlightsWithContext(ctx context.Context, request *SubmitDoubleVideoHighlightsRequest) (response *SubmitDoubleVideoHighlightsResponse, err error) {
    if request == nil {
        request = NewSubmitDoubleVideoHighlightsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SubmitDoubleVideoHighlights require credential")
    }

    request.SetContext(ctx)
    
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

// SubmitFullBodyClassTask
// **传统课堂授课任务**：在此场景中，老师为站立授课，有白板或投影供老师展示课程内容，摄像头可以拍摄到老师的半身或者全身。拍摄视频为一路全局画面，且背景不动，要求画面稳定清晰。通过此接口可分析老师授课的行为及语音，以支持AI评教。  
//
//   
//
// **提供的功能接口有：**老师人脸识别、老师表情识别、老师肢体动作识别、语音识别。  可分析的指标维度包括：身份识别、正脸、侧脸、人脸坐标、人脸尺寸、高兴、中性、高兴、中性、惊讶、厌恶、恐惧、愤怒、蔑视、悲伤、正面讲解、写板书、指黑板、语音转文字、发音时长、非发音时长、音量、语速、指定关键词的使用等
//
//   
//
// **对场景的要求为：**真实场景老师1人出现在画面中，全局画面且背景不动；人脸上下角度在20度以内，左右角度在15度以内，歪头角度在15度以内；光照均匀，无遮挡，人脸清晰可见；像素最好在 100X100 像素以上，但是图像整体质量不能超过1080p。
//
//     
//
// **结果查询方式：**图像任务直接返回结果，点播及直播任务通过DescribeAITaskResult查询结果。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_JOBCANCELED = "InternalError.JobCanceled"
//  INTERNALERROR_JOBREACHMAXRETRYTIMES = "InternalError.JobReachMaxRetryTimes"
//  INTERNALERROR_STATUSQUERYFACELIBFAILED = "InternalError.StatusQueryFaceLibFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_BODYRETRIEVALFAILED = "InvalidParameter.BodyRetrievalFailed"
//  INVALIDPARAMETER_CANNOTFINDFACEINPICTURE = "InvalidParameter.CannotFindFaceInPicture"
//  INVALIDPARAMETER_CLIPPINGFRAMEFAILED = "InvalidParameter.ClippingFrameFailed"
//  INVALIDPARAMETER_COMPUTATIONALSIMILARITYFAILED = "InvalidParameter.ComputationalSimilarityFailed"
//  INVALIDPARAMETER_CREATEFACEFAILED = "InvalidParameter.CreateFaceFailed"
//  INVALIDPARAMETER_CREATEPERSONFAILED = "InvalidParameter.CreatePersonFailed"
//  INVALIDPARAMETER_FACEDETECTEDFAILED = "InvalidParameter.FaceDetectedFailed"
//  INVALIDPARAMETER_FACENOTEXIST = "InvalidParameter.FaceNotExist"
//  INVALIDPARAMETER_FACENOTEXISTED = "InvalidParameter.FaceNotExisted"
//  INVALIDPARAMETER_FACENUMBEREXCEEDLIMITOFSINGLEDETECTION = "InvalidParameter.FaceNumberExceedLimitOfSingleDetection"
//  INVALIDPARAMETER_FACENUMBERLIMIT = "InvalidParameter.FaceNumberLimit"
//  INVALIDPARAMETER_FACEREGISTRATIONFAILED = "InvalidParameter.FaceRegistrationFailed"
//  INVALIDPARAMETER_FACERETRIEVALFAILED = "InvalidParameter.FaceRetrievalFailed"
//  INVALIDPARAMETER_FACETOOSMALL = "InvalidParameter.FaceTooSmall"
//  INVALIDPARAMETER_FAILEDTODECODEIMAGE = "InvalidParameter.FailedToDecodeImage"
//  INVALIDPARAMETER_FEATUREEXTRACTIONFAILED = "InvalidParameter.FeatureExtractionFailed"
//  INVALIDPARAMETER_FILECONTENTEMPTY = "InvalidParameter.FileContentEmpty"
//  INVALIDPARAMETER_IMAGEDECODEFAILED = "InvalidParameter.ImageDecodeFailed"
//  INVALIDPARAMETER_IMAGETOOLARGE = "InvalidParameter.ImageTooLarge"
//  INVALIDPARAMETER_IMAGETOOSMALL = "InvalidParameter.ImageTooSmall"
//  INVALIDPARAMETER_INVALIDFILETYPE = "InvalidParameter.InvalidFileType"
//  INVALIDPARAMETER_INVALIDLANG = "InvalidParameter.InvalidLang"
//  INVALIDPARAMETER_INVALIDRESOURCEIDENTIFY = "InvalidParameter.InvalidResourceIdentify"
//  INVALIDPARAMETER_INVALIDRESOURCENAME = "InvalidParameter.InvalidResourceName"
//  INVALIDPARAMETER_INVALIDSEQID = "InvalidParameter.InvalidSeqId"
//  INVALIDPARAMETER_INVALIDURL = "InvalidParameter.InvalidUrl"
//  INVALIDPARAMETER_JOBNOTVALID = "InvalidParameter.JobNotValid"
//  INVALIDPARAMETER_LIBRARYNOTEXISTED = "InvalidParameter.LibraryNotExisted"
//  INVALIDPARAMETER_LOWQUALITYPICTURE = "InvalidParameter.LowQualityPicture"
//  INVALIDPARAMETER_NOSUCHSERVICE = "InvalidParameter.NoSuchService"
//  INVALIDPARAMETER_NOTATTENDANCEJOB = "InvalidParameter.NotAttendanceJob"
//  INVALIDPARAMETER_NOTVALIDPITCH = "InvalidParameter.NotValidPitch"
//  INVALIDPARAMETER_NOTVALIDROW = "InvalidParameter.NotValidRow"
//  INVALIDPARAMETER_NOTVALIDYAW = "InvalidParameter.NotValidYaw"
//  INVALIDPARAMETER_PERSONEXISTED = "InvalidParameter.PersonExisted"
//  INVALIDPARAMETER_PERSONIDNOTVALID = "InvalidParameter.PersonIdNotValid"
//  INVALIDPARAMETER_PERSONNOTEXIST = "InvalidParameter.PersonNotExist"
//  INVALIDPARAMETER_PERSONNOTEXISTED = "InvalidParameter.PersonNotExisted"
//  INVALIDPARAMETER_PERSONNOTREGISTERED = "InvalidParameter.PersonNotRegistered"
//  INVALIDPARAMETER_PROCESSUNREADY = "InvalidParameter.ProcessUnReady"
//  INVALIDPARAMETER_RESOURCENAMEDUPLICATE = "InvalidParameter.ResourceNameDuplicate"
//  INVALIDPARAMETER_STATUSJOBFAILED = "InvalidParameter.StatusJobFailed"
//  INVALIDPARAMETER_STATUSJOBNOTFOUND = "InvalidParameter.StatusJobNotFound"
//  INVALIDPARAMETER_STATUSJOBUNFINISHED = "InvalidParameter.StatusJobUnfinished"
//  INVALIDPARAMETER_STATUSNOCAMERAFOUND = "InvalidParameter.StatusNoCameraFound"
//  INVALIDPARAMETER_TOOMANYFACE = "InvalidParameter.TooManyFace"
//  INVALIDPARAMETER_TRACKINGFAILED = "InvalidParameter.TrackingFailed"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSEXCEED = "InvalidParameterValue.GroupExDescriptionsExceed"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMEIDENTICAL = "InvalidParameterValue.GroupExDescriptionsNameIdentical"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMEILLEGAL = "InvalidParameterValue.GroupExDescriptionsNameIllegal"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMETOOLONG = "InvalidParameterValue.GroupExDescriptionsNameTooLong"
//  LIMITEXCEEDED_STATUSCAMERACOUNTOVERMAX = "LimitExceeded.StatusCameraCountOverMax"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_HLJOBIDNOTFOUND = "ResourceNotFound.HLJobIdNotFound"
//  RESOURCEUNAVAILABLE_STATUSQUERYSTREAMFAILED = "ResourceUnavailable.StatusQueryStreamFailed"
func (c *Client) SubmitFullBodyClassTask(request *SubmitFullBodyClassTaskRequest) (response *SubmitFullBodyClassTaskResponse, err error) {
    return c.SubmitFullBodyClassTaskWithContext(context.Background(), request)
}

// SubmitFullBodyClassTask
// **传统课堂授课任务**：在此场景中，老师为站立授课，有白板或投影供老师展示课程内容，摄像头可以拍摄到老师的半身或者全身。拍摄视频为一路全局画面，且背景不动，要求画面稳定清晰。通过此接口可分析老师授课的行为及语音，以支持AI评教。  
//
//   
//
// **提供的功能接口有：**老师人脸识别、老师表情识别、老师肢体动作识别、语音识别。  可分析的指标维度包括：身份识别、正脸、侧脸、人脸坐标、人脸尺寸、高兴、中性、高兴、中性、惊讶、厌恶、恐惧、愤怒、蔑视、悲伤、正面讲解、写板书、指黑板、语音转文字、发音时长、非发音时长、音量、语速、指定关键词的使用等
//
//   
//
// **对场景的要求为：**真实场景老师1人出现在画面中，全局画面且背景不动；人脸上下角度在20度以内，左右角度在15度以内，歪头角度在15度以内；光照均匀，无遮挡，人脸清晰可见；像素最好在 100X100 像素以上，但是图像整体质量不能超过1080p。
//
//     
//
// **结果查询方式：**图像任务直接返回结果，点播及直播任务通过DescribeAITaskResult查询结果。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_JOBCANCELED = "InternalError.JobCanceled"
//  INTERNALERROR_JOBREACHMAXRETRYTIMES = "InternalError.JobReachMaxRetryTimes"
//  INTERNALERROR_STATUSQUERYFACELIBFAILED = "InternalError.StatusQueryFaceLibFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_BODYRETRIEVALFAILED = "InvalidParameter.BodyRetrievalFailed"
//  INVALIDPARAMETER_CANNOTFINDFACEINPICTURE = "InvalidParameter.CannotFindFaceInPicture"
//  INVALIDPARAMETER_CLIPPINGFRAMEFAILED = "InvalidParameter.ClippingFrameFailed"
//  INVALIDPARAMETER_COMPUTATIONALSIMILARITYFAILED = "InvalidParameter.ComputationalSimilarityFailed"
//  INVALIDPARAMETER_CREATEFACEFAILED = "InvalidParameter.CreateFaceFailed"
//  INVALIDPARAMETER_CREATEPERSONFAILED = "InvalidParameter.CreatePersonFailed"
//  INVALIDPARAMETER_FACEDETECTEDFAILED = "InvalidParameter.FaceDetectedFailed"
//  INVALIDPARAMETER_FACENOTEXIST = "InvalidParameter.FaceNotExist"
//  INVALIDPARAMETER_FACENOTEXISTED = "InvalidParameter.FaceNotExisted"
//  INVALIDPARAMETER_FACENUMBEREXCEEDLIMITOFSINGLEDETECTION = "InvalidParameter.FaceNumberExceedLimitOfSingleDetection"
//  INVALIDPARAMETER_FACENUMBERLIMIT = "InvalidParameter.FaceNumberLimit"
//  INVALIDPARAMETER_FACEREGISTRATIONFAILED = "InvalidParameter.FaceRegistrationFailed"
//  INVALIDPARAMETER_FACERETRIEVALFAILED = "InvalidParameter.FaceRetrievalFailed"
//  INVALIDPARAMETER_FACETOOSMALL = "InvalidParameter.FaceTooSmall"
//  INVALIDPARAMETER_FAILEDTODECODEIMAGE = "InvalidParameter.FailedToDecodeImage"
//  INVALIDPARAMETER_FEATUREEXTRACTIONFAILED = "InvalidParameter.FeatureExtractionFailed"
//  INVALIDPARAMETER_FILECONTENTEMPTY = "InvalidParameter.FileContentEmpty"
//  INVALIDPARAMETER_IMAGEDECODEFAILED = "InvalidParameter.ImageDecodeFailed"
//  INVALIDPARAMETER_IMAGETOOLARGE = "InvalidParameter.ImageTooLarge"
//  INVALIDPARAMETER_IMAGETOOSMALL = "InvalidParameter.ImageTooSmall"
//  INVALIDPARAMETER_INVALIDFILETYPE = "InvalidParameter.InvalidFileType"
//  INVALIDPARAMETER_INVALIDLANG = "InvalidParameter.InvalidLang"
//  INVALIDPARAMETER_INVALIDRESOURCEIDENTIFY = "InvalidParameter.InvalidResourceIdentify"
//  INVALIDPARAMETER_INVALIDRESOURCENAME = "InvalidParameter.InvalidResourceName"
//  INVALIDPARAMETER_INVALIDSEQID = "InvalidParameter.InvalidSeqId"
//  INVALIDPARAMETER_INVALIDURL = "InvalidParameter.InvalidUrl"
//  INVALIDPARAMETER_JOBNOTVALID = "InvalidParameter.JobNotValid"
//  INVALIDPARAMETER_LIBRARYNOTEXISTED = "InvalidParameter.LibraryNotExisted"
//  INVALIDPARAMETER_LOWQUALITYPICTURE = "InvalidParameter.LowQualityPicture"
//  INVALIDPARAMETER_NOSUCHSERVICE = "InvalidParameter.NoSuchService"
//  INVALIDPARAMETER_NOTATTENDANCEJOB = "InvalidParameter.NotAttendanceJob"
//  INVALIDPARAMETER_NOTVALIDPITCH = "InvalidParameter.NotValidPitch"
//  INVALIDPARAMETER_NOTVALIDROW = "InvalidParameter.NotValidRow"
//  INVALIDPARAMETER_NOTVALIDYAW = "InvalidParameter.NotValidYaw"
//  INVALIDPARAMETER_PERSONEXISTED = "InvalidParameter.PersonExisted"
//  INVALIDPARAMETER_PERSONIDNOTVALID = "InvalidParameter.PersonIdNotValid"
//  INVALIDPARAMETER_PERSONNOTEXIST = "InvalidParameter.PersonNotExist"
//  INVALIDPARAMETER_PERSONNOTEXISTED = "InvalidParameter.PersonNotExisted"
//  INVALIDPARAMETER_PERSONNOTREGISTERED = "InvalidParameter.PersonNotRegistered"
//  INVALIDPARAMETER_PROCESSUNREADY = "InvalidParameter.ProcessUnReady"
//  INVALIDPARAMETER_RESOURCENAMEDUPLICATE = "InvalidParameter.ResourceNameDuplicate"
//  INVALIDPARAMETER_STATUSJOBFAILED = "InvalidParameter.StatusJobFailed"
//  INVALIDPARAMETER_STATUSJOBNOTFOUND = "InvalidParameter.StatusJobNotFound"
//  INVALIDPARAMETER_STATUSJOBUNFINISHED = "InvalidParameter.StatusJobUnfinished"
//  INVALIDPARAMETER_STATUSNOCAMERAFOUND = "InvalidParameter.StatusNoCameraFound"
//  INVALIDPARAMETER_TOOMANYFACE = "InvalidParameter.TooManyFace"
//  INVALIDPARAMETER_TRACKINGFAILED = "InvalidParameter.TrackingFailed"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSEXCEED = "InvalidParameterValue.GroupExDescriptionsExceed"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMEIDENTICAL = "InvalidParameterValue.GroupExDescriptionsNameIdentical"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMEILLEGAL = "InvalidParameterValue.GroupExDescriptionsNameIllegal"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMETOOLONG = "InvalidParameterValue.GroupExDescriptionsNameTooLong"
//  LIMITEXCEEDED_STATUSCAMERACOUNTOVERMAX = "LimitExceeded.StatusCameraCountOverMax"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_HLJOBIDNOTFOUND = "ResourceNotFound.HLJobIdNotFound"
//  RESOURCEUNAVAILABLE_STATUSQUERYSTREAMFAILED = "ResourceUnavailable.StatusQueryStreamFailed"
func (c *Client) SubmitFullBodyClassTaskWithContext(ctx context.Context, request *SubmitFullBodyClassTaskRequest) (response *SubmitFullBodyClassTaskResponse, err error) {
    if request == nil {
        request = NewSubmitFullBodyClassTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SubmitFullBodyClassTask require credential")
    }

    request.SetContext(ctx)
    
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

// SubmitHighlights
// 发起视频生成精彩集锦接口。该接口可以通过客户传入的课程音频数据及相关策略（如微笑抽取，专注抽取等），自动生成一堂课程的精彩集锦。需要通过QueryHighlightResult接口获取生成结果。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_JOBCANCELED = "InternalError.JobCanceled"
//  INTERNALERROR_JOBREACHMAXRETRYTIMES = "InternalError.JobReachMaxRetryTimes"
//  INTERNALERROR_STATUSQUERYFACELIBFAILED = "InternalError.StatusQueryFaceLibFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_BODYRETRIEVALFAILED = "InvalidParameter.BodyRetrievalFailed"
//  INVALIDPARAMETER_CANNOTFINDFACEINPICTURE = "InvalidParameter.CannotFindFaceInPicture"
//  INVALIDPARAMETER_CLIPPINGFRAMEFAILED = "InvalidParameter.ClippingFrameFailed"
//  INVALIDPARAMETER_COMPUTATIONALSIMILARITYFAILED = "InvalidParameter.ComputationalSimilarityFailed"
//  INVALIDPARAMETER_CREATEFACEFAILED = "InvalidParameter.CreateFaceFailed"
//  INVALIDPARAMETER_CREATEPERSONFAILED = "InvalidParameter.CreatePersonFailed"
//  INVALIDPARAMETER_FACEDETECTEDFAILED = "InvalidParameter.FaceDetectedFailed"
//  INVALIDPARAMETER_FACENOTEXIST = "InvalidParameter.FaceNotExist"
//  INVALIDPARAMETER_FACENOTEXISTED = "InvalidParameter.FaceNotExisted"
//  INVALIDPARAMETER_FACENUMBEREXCEEDLIMITOFSINGLEDETECTION = "InvalidParameter.FaceNumberExceedLimitOfSingleDetection"
//  INVALIDPARAMETER_FACENUMBERLIMIT = "InvalidParameter.FaceNumberLimit"
//  INVALIDPARAMETER_FACEREGISTRATIONFAILED = "InvalidParameter.FaceRegistrationFailed"
//  INVALIDPARAMETER_FACERETRIEVALFAILED = "InvalidParameter.FaceRetrievalFailed"
//  INVALIDPARAMETER_FACETOOSMALL = "InvalidParameter.FaceTooSmall"
//  INVALIDPARAMETER_FAILEDTODECODEIMAGE = "InvalidParameter.FailedToDecodeImage"
//  INVALIDPARAMETER_FEATUREEXTRACTIONFAILED = "InvalidParameter.FeatureExtractionFailed"
//  INVALIDPARAMETER_FILECONTENTEMPTY = "InvalidParameter.FileContentEmpty"
//  INVALIDPARAMETER_IMAGEDECODEFAILED = "InvalidParameter.ImageDecodeFailed"
//  INVALIDPARAMETER_IMAGETOOLARGE = "InvalidParameter.ImageTooLarge"
//  INVALIDPARAMETER_IMAGETOOSMALL = "InvalidParameter.ImageTooSmall"
//  INVALIDPARAMETER_INVALIDFILETYPE = "InvalidParameter.InvalidFileType"
//  INVALIDPARAMETER_INVALIDLANG = "InvalidParameter.InvalidLang"
//  INVALIDPARAMETER_INVALIDRESOURCENAME = "InvalidParameter.InvalidResourceName"
//  INVALIDPARAMETER_INVALIDURL = "InvalidParameter.InvalidUrl"
//  INVALIDPARAMETER_JOBNOTVALID = "InvalidParameter.JobNotValid"
//  INVALIDPARAMETER_LIBRARYNOTEXISTED = "InvalidParameter.LibraryNotExisted"
//  INVALIDPARAMETER_LOWQUALITYPICTURE = "InvalidParameter.LowQualityPicture"
//  INVALIDPARAMETER_NOSUCHSERVICE = "InvalidParameter.NoSuchService"
//  INVALIDPARAMETER_NOTATTENDANCEJOB = "InvalidParameter.NotAttendanceJob"
//  INVALIDPARAMETER_NOTVALIDPITCH = "InvalidParameter.NotValidPitch"
//  INVALIDPARAMETER_NOTVALIDROW = "InvalidParameter.NotValidRow"
//  INVALIDPARAMETER_NOTVALIDYAW = "InvalidParameter.NotValidYaw"
//  INVALIDPARAMETER_PERSONEXISTED = "InvalidParameter.PersonExisted"
//  INVALIDPARAMETER_PERSONIDNOTVALID = "InvalidParameter.PersonIdNotValid"
//  INVALIDPARAMETER_PERSONNOTEXIST = "InvalidParameter.PersonNotExist"
//  INVALIDPARAMETER_PERSONNOTEXISTED = "InvalidParameter.PersonNotExisted"
//  INVALIDPARAMETER_PERSONNOTREGISTERED = "InvalidParameter.PersonNotRegistered"
//  INVALIDPARAMETER_PROCESSUNREADY = "InvalidParameter.ProcessUnReady"
//  INVALIDPARAMETER_RESOURCENAMEDUPLICATE = "InvalidParameter.ResourceNameDuplicate"
//  INVALIDPARAMETER_STATUSJOBFAILED = "InvalidParameter.StatusJobFailed"
//  INVALIDPARAMETER_STATUSJOBNOTFOUND = "InvalidParameter.StatusJobNotFound"
//  INVALIDPARAMETER_STATUSJOBUNFINISHED = "InvalidParameter.StatusJobUnfinished"
//  INVALIDPARAMETER_STATUSNOCAMERAFOUND = "InvalidParameter.StatusNoCameraFound"
//  INVALIDPARAMETER_TOOMANYFACE = "InvalidParameter.TooManyFace"
//  INVALIDPARAMETER_TRACKINGFAILED = "InvalidParameter.TrackingFailed"
//  LIMITEXCEEDED_STATUSCAMERACOUNTOVERMAX = "LimitExceeded.StatusCameraCountOverMax"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE_STATUSQUERYSTREAMFAILED = "ResourceUnavailable.StatusQueryStreamFailed"
func (c *Client) SubmitHighlights(request *SubmitHighlightsRequest) (response *SubmitHighlightsResponse, err error) {
    return c.SubmitHighlightsWithContext(context.Background(), request)
}

// SubmitHighlights
// 发起视频生成精彩集锦接口。该接口可以通过客户传入的课程音频数据及相关策略（如微笑抽取，专注抽取等），自动生成一堂课程的精彩集锦。需要通过QueryHighlightResult接口获取生成结果。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_JOBCANCELED = "InternalError.JobCanceled"
//  INTERNALERROR_JOBREACHMAXRETRYTIMES = "InternalError.JobReachMaxRetryTimes"
//  INTERNALERROR_STATUSQUERYFACELIBFAILED = "InternalError.StatusQueryFaceLibFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_BODYRETRIEVALFAILED = "InvalidParameter.BodyRetrievalFailed"
//  INVALIDPARAMETER_CANNOTFINDFACEINPICTURE = "InvalidParameter.CannotFindFaceInPicture"
//  INVALIDPARAMETER_CLIPPINGFRAMEFAILED = "InvalidParameter.ClippingFrameFailed"
//  INVALIDPARAMETER_COMPUTATIONALSIMILARITYFAILED = "InvalidParameter.ComputationalSimilarityFailed"
//  INVALIDPARAMETER_CREATEFACEFAILED = "InvalidParameter.CreateFaceFailed"
//  INVALIDPARAMETER_CREATEPERSONFAILED = "InvalidParameter.CreatePersonFailed"
//  INVALIDPARAMETER_FACEDETECTEDFAILED = "InvalidParameter.FaceDetectedFailed"
//  INVALIDPARAMETER_FACENOTEXIST = "InvalidParameter.FaceNotExist"
//  INVALIDPARAMETER_FACENOTEXISTED = "InvalidParameter.FaceNotExisted"
//  INVALIDPARAMETER_FACENUMBEREXCEEDLIMITOFSINGLEDETECTION = "InvalidParameter.FaceNumberExceedLimitOfSingleDetection"
//  INVALIDPARAMETER_FACENUMBERLIMIT = "InvalidParameter.FaceNumberLimit"
//  INVALIDPARAMETER_FACEREGISTRATIONFAILED = "InvalidParameter.FaceRegistrationFailed"
//  INVALIDPARAMETER_FACERETRIEVALFAILED = "InvalidParameter.FaceRetrievalFailed"
//  INVALIDPARAMETER_FACETOOSMALL = "InvalidParameter.FaceTooSmall"
//  INVALIDPARAMETER_FAILEDTODECODEIMAGE = "InvalidParameter.FailedToDecodeImage"
//  INVALIDPARAMETER_FEATUREEXTRACTIONFAILED = "InvalidParameter.FeatureExtractionFailed"
//  INVALIDPARAMETER_FILECONTENTEMPTY = "InvalidParameter.FileContentEmpty"
//  INVALIDPARAMETER_IMAGEDECODEFAILED = "InvalidParameter.ImageDecodeFailed"
//  INVALIDPARAMETER_IMAGETOOLARGE = "InvalidParameter.ImageTooLarge"
//  INVALIDPARAMETER_IMAGETOOSMALL = "InvalidParameter.ImageTooSmall"
//  INVALIDPARAMETER_INVALIDFILETYPE = "InvalidParameter.InvalidFileType"
//  INVALIDPARAMETER_INVALIDLANG = "InvalidParameter.InvalidLang"
//  INVALIDPARAMETER_INVALIDRESOURCENAME = "InvalidParameter.InvalidResourceName"
//  INVALIDPARAMETER_INVALIDURL = "InvalidParameter.InvalidUrl"
//  INVALIDPARAMETER_JOBNOTVALID = "InvalidParameter.JobNotValid"
//  INVALIDPARAMETER_LIBRARYNOTEXISTED = "InvalidParameter.LibraryNotExisted"
//  INVALIDPARAMETER_LOWQUALITYPICTURE = "InvalidParameter.LowQualityPicture"
//  INVALIDPARAMETER_NOSUCHSERVICE = "InvalidParameter.NoSuchService"
//  INVALIDPARAMETER_NOTATTENDANCEJOB = "InvalidParameter.NotAttendanceJob"
//  INVALIDPARAMETER_NOTVALIDPITCH = "InvalidParameter.NotValidPitch"
//  INVALIDPARAMETER_NOTVALIDROW = "InvalidParameter.NotValidRow"
//  INVALIDPARAMETER_NOTVALIDYAW = "InvalidParameter.NotValidYaw"
//  INVALIDPARAMETER_PERSONEXISTED = "InvalidParameter.PersonExisted"
//  INVALIDPARAMETER_PERSONIDNOTVALID = "InvalidParameter.PersonIdNotValid"
//  INVALIDPARAMETER_PERSONNOTEXIST = "InvalidParameter.PersonNotExist"
//  INVALIDPARAMETER_PERSONNOTEXISTED = "InvalidParameter.PersonNotExisted"
//  INVALIDPARAMETER_PERSONNOTREGISTERED = "InvalidParameter.PersonNotRegistered"
//  INVALIDPARAMETER_PROCESSUNREADY = "InvalidParameter.ProcessUnReady"
//  INVALIDPARAMETER_RESOURCENAMEDUPLICATE = "InvalidParameter.ResourceNameDuplicate"
//  INVALIDPARAMETER_STATUSJOBFAILED = "InvalidParameter.StatusJobFailed"
//  INVALIDPARAMETER_STATUSJOBNOTFOUND = "InvalidParameter.StatusJobNotFound"
//  INVALIDPARAMETER_STATUSJOBUNFINISHED = "InvalidParameter.StatusJobUnfinished"
//  INVALIDPARAMETER_STATUSNOCAMERAFOUND = "InvalidParameter.StatusNoCameraFound"
//  INVALIDPARAMETER_TOOMANYFACE = "InvalidParameter.TooManyFace"
//  INVALIDPARAMETER_TRACKINGFAILED = "InvalidParameter.TrackingFailed"
//  LIMITEXCEEDED_STATUSCAMERACOUNTOVERMAX = "LimitExceeded.StatusCameraCountOverMax"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE_STATUSQUERYSTREAMFAILED = "ResourceUnavailable.StatusQueryStreamFailed"
func (c *Client) SubmitHighlightsWithContext(ctx context.Context, request *SubmitHighlightsRequest) (response *SubmitHighlightsResponse, err error) {
    if request == nil {
        request = NewSubmitHighlightsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SubmitHighlights require credential")
    }

    request.SetContext(ctx)
    
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

// SubmitImageTask
// 提交图像分析任务
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_UNSUPPORTEDFILETYPEMOV = "InvalidParameter.UnsupportedFileTypeMov"
//  INVALIDPARAMETER_UNSUPPORTEDVIDEOSIZE = "InvalidParameter.UnsupportedVideoSize"
func (c *Client) SubmitImageTask(request *SubmitImageTaskRequest) (response *SubmitImageTaskResponse, err error) {
    return c.SubmitImageTaskWithContext(context.Background(), request)
}

// SubmitImageTask
// 提交图像分析任务
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_UNSUPPORTEDFILETYPEMOV = "InvalidParameter.UnsupportedFileTypeMov"
//  INVALIDPARAMETER_UNSUPPORTEDVIDEOSIZE = "InvalidParameter.UnsupportedVideoSize"
func (c *Client) SubmitImageTaskWithContext(ctx context.Context, request *SubmitImageTaskRequest) (response *SubmitImageTaskResponse, err error) {
    if request == nil {
        request = NewSubmitImageTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SubmitImageTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewSubmitImageTaskResponse()
    err = c.Send(request, response)
    return
}

func NewSubmitImageTaskPlusRequest() (request *SubmitImageTaskPlusRequest) {
    request = &SubmitImageTaskPlusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tci", APIVersion, "SubmitImageTaskPlus")
    
    
    return
}

func NewSubmitImageTaskPlusResponse() (response *SubmitImageTaskPlusResponse) {
    response = &SubmitImageTaskPlusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// SubmitImageTaskPlus
// 高级图像分析任务，开放了图像任务里的所有开关，可以根据场景深度定制图像分析任务。支持的图像类别有，图片链接、图片二进制数据、点播链接和直播链接。
//
// 可能返回的错误码:
//  FAILEDOPERATION_NOTSUPPORTEDFUNCTIONERROR = "FailedOperation.NotSupportedFunctionError"
//  INTERNALERROR_JOBCANCELED = "InternalError.JobCanceled"
//  INTERNALERROR_JOBREACHMAXRETRYTIMES = "InternalError.JobReachMaxRetryTimes"
//  INTERNALERROR_STATUSQUERYFACELIBFAILED = "InternalError.StatusQueryFaceLibFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_BODYRETRIEVALFAILED = "InvalidParameter.BodyRetrievalFailed"
//  INVALIDPARAMETER_CANNOTFINDFACEERROR = "InvalidParameter.CannotFindFaceError"
//  INVALIDPARAMETER_CANNOTFINDFACEINPICTURE = "InvalidParameter.CannotFindFaceInPicture"
//  INVALIDPARAMETER_CANNOTREADVIDEOFROMURLERROR = "InvalidParameter.CannotReadVideoFromUrlError"
//  INVALIDPARAMETER_CLIPPINGFRAMEFAILED = "InvalidParameter.ClippingFrameFailed"
//  INVALIDPARAMETER_COMPUTATIONALSIMILARITYFAILED = "InvalidParameter.ComputationalSimilarityFailed"
//  INVALIDPARAMETER_CREATEFACEFAILED = "InvalidParameter.CreateFaceFailed"
//  INVALIDPARAMETER_CREATEPERSONFAILED = "InvalidParameter.CreatePersonFailed"
//  INVALIDPARAMETER_FACEDETECTEDFAILED = "InvalidParameter.FaceDetectedFailed"
//  INVALIDPARAMETER_FACENOTEXIST = "InvalidParameter.FaceNotExist"
//  INVALIDPARAMETER_FACENOTEXISTED = "InvalidParameter.FaceNotExisted"
//  INVALIDPARAMETER_FACENUMBEREXCEEDLIMITOFSINGLEDETECTION = "InvalidParameter.FaceNumberExceedLimitOfSingleDetection"
//  INVALIDPARAMETER_FACENUMBERLIMIT = "InvalidParameter.FaceNumberLimit"
//  INVALIDPARAMETER_FACEREGISTRATIONFAILED = "InvalidParameter.FaceRegistrationFailed"
//  INVALIDPARAMETER_FACERETRIEVALFAILED = "InvalidParameter.FaceRetrievalFailed"
//  INVALIDPARAMETER_FACETOOSMALL = "InvalidParameter.FaceTooSmall"
//  INVALIDPARAMETER_FAILEDTODECODEIMAGE = "InvalidParameter.FailedToDecodeImage"
//  INVALIDPARAMETER_FEATUREEXTRACTIONFAILED = "InvalidParameter.FeatureExtractionFailed"
//  INVALIDPARAMETER_FILECONTENTEMPTY = "InvalidParameter.FileContentEmpty"
//  INVALIDPARAMETER_IMAGEDECODEFAILED = "InvalidParameter.ImageDecodeFailed"
//  INVALIDPARAMETER_IMAGETOOLARGE = "InvalidParameter.ImageTooLarge"
//  INVALIDPARAMETER_IMAGETOOSMALL = "InvalidParameter.ImageTooSmall"
//  INVALIDPARAMETER_INVALIDFILETYPE = "InvalidParameter.InvalidFileType"
//  INVALIDPARAMETER_INVALIDLANG = "InvalidParameter.InvalidLang"
//  INVALIDPARAMETER_INVALIDRESOURCEIDENTIFY = "InvalidParameter.InvalidResourceIdentify"
//  INVALIDPARAMETER_INVALIDRESOURCENAME = "InvalidParameter.InvalidResourceName"
//  INVALIDPARAMETER_INVALIDSEQID = "InvalidParameter.InvalidSeqId"
//  INVALIDPARAMETER_INVALIDURL = "InvalidParameter.InvalidUrl"
//  INVALIDPARAMETER_JOBNOTVALID = "InvalidParameter.JobNotValid"
//  INVALIDPARAMETER_LIBHAVENOPERSON = "InvalidParameter.LibHaveNoPerson"
//  INVALIDPARAMETER_LIBISEMPTY = "InvalidParameter.LibIsEmpty"
//  INVALIDPARAMETER_LIBRARYNOTEXISTED = "InvalidParameter.LibraryNotExisted"
//  INVALIDPARAMETER_LOWQUALITYPICTURE = "InvalidParameter.LowQualityPicture"
//  INVALIDPARAMETER_NOSUCHSERVICE = "InvalidParameter.NoSuchService"
//  INVALIDPARAMETER_NOTATTENDANCEJOB = "InvalidParameter.NotAttendanceJob"
//  INVALIDPARAMETER_NOTVALIDPITCH = "InvalidParameter.NotValidPitch"
//  INVALIDPARAMETER_NOTVALIDROW = "InvalidParameter.NotValidRow"
//  INVALIDPARAMETER_NOTVALIDYAW = "InvalidParameter.NotValidYaw"
//  INVALIDPARAMETER_PERSONEXISTED = "InvalidParameter.PersonExisted"
//  INVALIDPARAMETER_PERSONIDNOTVALID = "InvalidParameter.PersonIdNotValid"
//  INVALIDPARAMETER_PERSONNOTEXIST = "InvalidParameter.PersonNotExist"
//  INVALIDPARAMETER_PERSONNOTEXISTED = "InvalidParameter.PersonNotExisted"
//  INVALIDPARAMETER_PERSONNOTREGISTERED = "InvalidParameter.PersonNotRegistered"
//  INVALIDPARAMETER_PROCESSUNREADY = "InvalidParameter.ProcessUnReady"
//  INVALIDPARAMETER_RESOURCENAMEDUPLICATE = "InvalidParameter.ResourceNameDuplicate"
//  INVALIDPARAMETER_STATUSJOBFAILED = "InvalidParameter.StatusJobFailed"
//  INVALIDPARAMETER_STATUSJOBNOTFOUND = "InvalidParameter.StatusJobNotFound"
//  INVALIDPARAMETER_STATUSJOBUNFINISHED = "InvalidParameter.StatusJobUnfinished"
//  INVALIDPARAMETER_STATUSNOCAMERAFOUND = "InvalidParameter.StatusNoCameraFound"
//  INVALIDPARAMETER_TOOMANYFACE = "InvalidParameter.TooManyFace"
//  INVALIDPARAMETER_TRACKINGFAILED = "InvalidParameter.TrackingFailed"
//  INVALIDPARAMETER_VIDEOALREDYPROCESSEDERROR = "InvalidParameter.VideoAlredyProcessedError"
//  INVALIDPARAMETERVALUE_GETHTTPBODYERROR = "InvalidParameterValue.GetHttpBodyError"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSEXCEED = "InvalidParameterValue.GroupExDescriptionsExceed"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMEIDENTICAL = "InvalidParameterValue.GroupExDescriptionsNameIdentical"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMEILLEGAL = "InvalidParameterValue.GroupExDescriptionsNameIllegal"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMETOOLONG = "InvalidParameterValue.GroupExDescriptionsNameTooLong"
//  INVALIDPARAMETERVALUE_IMAGEILLEGAL = "InvalidParameterValue.ImageIllegal"
//  INVALIDPARAMETERVALUE_INVALIDSEQID = "InvalidParameterValue.InvalidSeqId"
//  INVALIDPARAMETERVALUE_NOTSUPPORTEDFUNCTIONERROR = "InvalidParameterValue.NotSupportedFunctionError"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONINFOSEXCEED = "InvalidParameterValue.PersonExDescriptionInfosExceed"
//  LIMITEXCEEDED_STATUSCAMERACOUNTOVERMAX = "LimitExceeded.StatusCameraCountOverMax"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_HLJOBIDNOTFOUND = "ResourceNotFound.HLJobIdNotFound"
//  RESOURCEUNAVAILABLE_CANNOTFINDSESSION = "ResourceUnavailable.CannotFindSession"
//  RESOURCEUNAVAILABLE_STATUSQUERYSTREAMFAILED = "ResourceUnavailable.StatusQueryStreamFailed"
func (c *Client) SubmitImageTaskPlus(request *SubmitImageTaskPlusRequest) (response *SubmitImageTaskPlusResponse, err error) {
    return c.SubmitImageTaskPlusWithContext(context.Background(), request)
}

// SubmitImageTaskPlus
// 高级图像分析任务，开放了图像任务里的所有开关，可以根据场景深度定制图像分析任务。支持的图像类别有，图片链接、图片二进制数据、点播链接和直播链接。
//
// 可能返回的错误码:
//  FAILEDOPERATION_NOTSUPPORTEDFUNCTIONERROR = "FailedOperation.NotSupportedFunctionError"
//  INTERNALERROR_JOBCANCELED = "InternalError.JobCanceled"
//  INTERNALERROR_JOBREACHMAXRETRYTIMES = "InternalError.JobReachMaxRetryTimes"
//  INTERNALERROR_STATUSQUERYFACELIBFAILED = "InternalError.StatusQueryFaceLibFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_BODYRETRIEVALFAILED = "InvalidParameter.BodyRetrievalFailed"
//  INVALIDPARAMETER_CANNOTFINDFACEERROR = "InvalidParameter.CannotFindFaceError"
//  INVALIDPARAMETER_CANNOTFINDFACEINPICTURE = "InvalidParameter.CannotFindFaceInPicture"
//  INVALIDPARAMETER_CANNOTREADVIDEOFROMURLERROR = "InvalidParameter.CannotReadVideoFromUrlError"
//  INVALIDPARAMETER_CLIPPINGFRAMEFAILED = "InvalidParameter.ClippingFrameFailed"
//  INVALIDPARAMETER_COMPUTATIONALSIMILARITYFAILED = "InvalidParameter.ComputationalSimilarityFailed"
//  INVALIDPARAMETER_CREATEFACEFAILED = "InvalidParameter.CreateFaceFailed"
//  INVALIDPARAMETER_CREATEPERSONFAILED = "InvalidParameter.CreatePersonFailed"
//  INVALIDPARAMETER_FACEDETECTEDFAILED = "InvalidParameter.FaceDetectedFailed"
//  INVALIDPARAMETER_FACENOTEXIST = "InvalidParameter.FaceNotExist"
//  INVALIDPARAMETER_FACENOTEXISTED = "InvalidParameter.FaceNotExisted"
//  INVALIDPARAMETER_FACENUMBEREXCEEDLIMITOFSINGLEDETECTION = "InvalidParameter.FaceNumberExceedLimitOfSingleDetection"
//  INVALIDPARAMETER_FACENUMBERLIMIT = "InvalidParameter.FaceNumberLimit"
//  INVALIDPARAMETER_FACEREGISTRATIONFAILED = "InvalidParameter.FaceRegistrationFailed"
//  INVALIDPARAMETER_FACERETRIEVALFAILED = "InvalidParameter.FaceRetrievalFailed"
//  INVALIDPARAMETER_FACETOOSMALL = "InvalidParameter.FaceTooSmall"
//  INVALIDPARAMETER_FAILEDTODECODEIMAGE = "InvalidParameter.FailedToDecodeImage"
//  INVALIDPARAMETER_FEATUREEXTRACTIONFAILED = "InvalidParameter.FeatureExtractionFailed"
//  INVALIDPARAMETER_FILECONTENTEMPTY = "InvalidParameter.FileContentEmpty"
//  INVALIDPARAMETER_IMAGEDECODEFAILED = "InvalidParameter.ImageDecodeFailed"
//  INVALIDPARAMETER_IMAGETOOLARGE = "InvalidParameter.ImageTooLarge"
//  INVALIDPARAMETER_IMAGETOOSMALL = "InvalidParameter.ImageTooSmall"
//  INVALIDPARAMETER_INVALIDFILETYPE = "InvalidParameter.InvalidFileType"
//  INVALIDPARAMETER_INVALIDLANG = "InvalidParameter.InvalidLang"
//  INVALIDPARAMETER_INVALIDRESOURCEIDENTIFY = "InvalidParameter.InvalidResourceIdentify"
//  INVALIDPARAMETER_INVALIDRESOURCENAME = "InvalidParameter.InvalidResourceName"
//  INVALIDPARAMETER_INVALIDSEQID = "InvalidParameter.InvalidSeqId"
//  INVALIDPARAMETER_INVALIDURL = "InvalidParameter.InvalidUrl"
//  INVALIDPARAMETER_JOBNOTVALID = "InvalidParameter.JobNotValid"
//  INVALIDPARAMETER_LIBHAVENOPERSON = "InvalidParameter.LibHaveNoPerson"
//  INVALIDPARAMETER_LIBISEMPTY = "InvalidParameter.LibIsEmpty"
//  INVALIDPARAMETER_LIBRARYNOTEXISTED = "InvalidParameter.LibraryNotExisted"
//  INVALIDPARAMETER_LOWQUALITYPICTURE = "InvalidParameter.LowQualityPicture"
//  INVALIDPARAMETER_NOSUCHSERVICE = "InvalidParameter.NoSuchService"
//  INVALIDPARAMETER_NOTATTENDANCEJOB = "InvalidParameter.NotAttendanceJob"
//  INVALIDPARAMETER_NOTVALIDPITCH = "InvalidParameter.NotValidPitch"
//  INVALIDPARAMETER_NOTVALIDROW = "InvalidParameter.NotValidRow"
//  INVALIDPARAMETER_NOTVALIDYAW = "InvalidParameter.NotValidYaw"
//  INVALIDPARAMETER_PERSONEXISTED = "InvalidParameter.PersonExisted"
//  INVALIDPARAMETER_PERSONIDNOTVALID = "InvalidParameter.PersonIdNotValid"
//  INVALIDPARAMETER_PERSONNOTEXIST = "InvalidParameter.PersonNotExist"
//  INVALIDPARAMETER_PERSONNOTEXISTED = "InvalidParameter.PersonNotExisted"
//  INVALIDPARAMETER_PERSONNOTREGISTERED = "InvalidParameter.PersonNotRegistered"
//  INVALIDPARAMETER_PROCESSUNREADY = "InvalidParameter.ProcessUnReady"
//  INVALIDPARAMETER_RESOURCENAMEDUPLICATE = "InvalidParameter.ResourceNameDuplicate"
//  INVALIDPARAMETER_STATUSJOBFAILED = "InvalidParameter.StatusJobFailed"
//  INVALIDPARAMETER_STATUSJOBNOTFOUND = "InvalidParameter.StatusJobNotFound"
//  INVALIDPARAMETER_STATUSJOBUNFINISHED = "InvalidParameter.StatusJobUnfinished"
//  INVALIDPARAMETER_STATUSNOCAMERAFOUND = "InvalidParameter.StatusNoCameraFound"
//  INVALIDPARAMETER_TOOMANYFACE = "InvalidParameter.TooManyFace"
//  INVALIDPARAMETER_TRACKINGFAILED = "InvalidParameter.TrackingFailed"
//  INVALIDPARAMETER_VIDEOALREDYPROCESSEDERROR = "InvalidParameter.VideoAlredyProcessedError"
//  INVALIDPARAMETERVALUE_GETHTTPBODYERROR = "InvalidParameterValue.GetHttpBodyError"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSEXCEED = "InvalidParameterValue.GroupExDescriptionsExceed"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMEIDENTICAL = "InvalidParameterValue.GroupExDescriptionsNameIdentical"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMEILLEGAL = "InvalidParameterValue.GroupExDescriptionsNameIllegal"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMETOOLONG = "InvalidParameterValue.GroupExDescriptionsNameTooLong"
//  INVALIDPARAMETERVALUE_IMAGEILLEGAL = "InvalidParameterValue.ImageIllegal"
//  INVALIDPARAMETERVALUE_INVALIDSEQID = "InvalidParameterValue.InvalidSeqId"
//  INVALIDPARAMETERVALUE_NOTSUPPORTEDFUNCTIONERROR = "InvalidParameterValue.NotSupportedFunctionError"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONINFOSEXCEED = "InvalidParameterValue.PersonExDescriptionInfosExceed"
//  LIMITEXCEEDED_STATUSCAMERACOUNTOVERMAX = "LimitExceeded.StatusCameraCountOverMax"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_HLJOBIDNOTFOUND = "ResourceNotFound.HLJobIdNotFound"
//  RESOURCEUNAVAILABLE_CANNOTFINDSESSION = "ResourceUnavailable.CannotFindSession"
//  RESOURCEUNAVAILABLE_STATUSQUERYSTREAMFAILED = "ResourceUnavailable.StatusQueryStreamFailed"
func (c *Client) SubmitImageTaskPlusWithContext(ctx context.Context, request *SubmitImageTaskPlusRequest) (response *SubmitImageTaskPlusResponse, err error) {
    if request == nil {
        request = NewSubmitImageTaskPlusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SubmitImageTaskPlus require credential")
    }

    request.SetContext(ctx)
    
    response = NewSubmitImageTaskPlusResponse()
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

// SubmitOneByOneClassTask
// **提交在线1对1课堂任务**  
//
// 对于在线1对1课堂，老师通过视频向学生授课，并且学生人数为1人。通过上传学生端的图像信息，可以获取学生的听课情况分析。 具体指一路全局画面且背景不动，有1位学生的头像或上半身的画面，要求画面稳定清晰。
//
//   
//
// **提供的功能接口有：**学生人脸识别、学生表情识别、语音识别。可分析的指标维度包括：学生身份识别、正脸、侧脸、抬头、低头、人脸坐标、人脸尺寸、高兴、中性、高兴、中性、惊讶、厌恶、恐惧、愤怒、蔑视、悲伤、语音转文字、发音时长、非发音时长、音量、语速等。
//
//   
//
// **对场景的要求为：**真实常规1v1授课场景，学生2人以下，全局画面且背景不动；人脸上下角度在20度以内，左右角度在15度以内，歪头角度在15度以内；光照均匀，无遮挡，人脸清晰可见；像素最好在 100X100 像素以上，但是图像整体质量不能超过1080p。
//
//     
//
// **结果查询方式：**图像任务直接返回结果，点播及直播任务通过DescribeAITaskResult查询结果。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_JOBCANCELED = "InternalError.JobCanceled"
//  INTERNALERROR_JOBREACHMAXRETRYTIMES = "InternalError.JobReachMaxRetryTimes"
//  INTERNALERROR_STATUSQUERYFACELIBFAILED = "InternalError.StatusQueryFaceLibFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_BODYRETRIEVALFAILED = "InvalidParameter.BodyRetrievalFailed"
//  INVALIDPARAMETER_CANNOTFINDFACEINPICTURE = "InvalidParameter.CannotFindFaceInPicture"
//  INVALIDPARAMETER_CLIPPINGFRAMEFAILED = "InvalidParameter.ClippingFrameFailed"
//  INVALIDPARAMETER_COMPUTATIONALSIMILARITYFAILED = "InvalidParameter.ComputationalSimilarityFailed"
//  INVALIDPARAMETER_CREATEFACEFAILED = "InvalidParameter.CreateFaceFailed"
//  INVALIDPARAMETER_CREATEPERSONFAILED = "InvalidParameter.CreatePersonFailed"
//  INVALIDPARAMETER_FACEDETECTEDFAILED = "InvalidParameter.FaceDetectedFailed"
//  INVALIDPARAMETER_FACENOTEXIST = "InvalidParameter.FaceNotExist"
//  INVALIDPARAMETER_FACENOTEXISTED = "InvalidParameter.FaceNotExisted"
//  INVALIDPARAMETER_FACENUMBEREXCEEDLIMITOFSINGLEDETECTION = "InvalidParameter.FaceNumberExceedLimitOfSingleDetection"
//  INVALIDPARAMETER_FACENUMBERLIMIT = "InvalidParameter.FaceNumberLimit"
//  INVALIDPARAMETER_FACEREGISTRATIONFAILED = "InvalidParameter.FaceRegistrationFailed"
//  INVALIDPARAMETER_FACERETRIEVALFAILED = "InvalidParameter.FaceRetrievalFailed"
//  INVALIDPARAMETER_FACETOOSMALL = "InvalidParameter.FaceTooSmall"
//  INVALIDPARAMETER_FAILEDTODECODEIMAGE = "InvalidParameter.FailedToDecodeImage"
//  INVALIDPARAMETER_FEATUREEXTRACTIONFAILED = "InvalidParameter.FeatureExtractionFailed"
//  INVALIDPARAMETER_FILECONTENTEMPTY = "InvalidParameter.FileContentEmpty"
//  INVALIDPARAMETER_IMAGEDECODEFAILED = "InvalidParameter.ImageDecodeFailed"
//  INVALIDPARAMETER_IMAGETOOLARGE = "InvalidParameter.ImageTooLarge"
//  INVALIDPARAMETER_IMAGETOOSMALL = "InvalidParameter.ImageTooSmall"
//  INVALIDPARAMETER_INVALIDFILETYPE = "InvalidParameter.InvalidFileType"
//  INVALIDPARAMETER_INVALIDLANG = "InvalidParameter.InvalidLang"
//  INVALIDPARAMETER_INVALIDRESOURCEIDENTIFY = "InvalidParameter.InvalidResourceIdentify"
//  INVALIDPARAMETER_INVALIDRESOURCENAME = "InvalidParameter.InvalidResourceName"
//  INVALIDPARAMETER_INVALIDSEQID = "InvalidParameter.InvalidSeqId"
//  INVALIDPARAMETER_INVALIDURL = "InvalidParameter.InvalidUrl"
//  INVALIDPARAMETER_JOBNOTVALID = "InvalidParameter.JobNotValid"
//  INVALIDPARAMETER_LIBRARYNOTEXISTED = "InvalidParameter.LibraryNotExisted"
//  INVALIDPARAMETER_LOWQUALITYPICTURE = "InvalidParameter.LowQualityPicture"
//  INVALIDPARAMETER_NOSUCHSERVICE = "InvalidParameter.NoSuchService"
//  INVALIDPARAMETER_NOTATTENDANCEJOB = "InvalidParameter.NotAttendanceJob"
//  INVALIDPARAMETER_NOTVALIDPITCH = "InvalidParameter.NotValidPitch"
//  INVALIDPARAMETER_NOTVALIDROW = "InvalidParameter.NotValidRow"
//  INVALIDPARAMETER_NOTVALIDYAW = "InvalidParameter.NotValidYaw"
//  INVALIDPARAMETER_PERSONEXISTED = "InvalidParameter.PersonExisted"
//  INVALIDPARAMETER_PERSONIDNOTVALID = "InvalidParameter.PersonIdNotValid"
//  INVALIDPARAMETER_PERSONNOTEXIST = "InvalidParameter.PersonNotExist"
//  INVALIDPARAMETER_PERSONNOTEXISTED = "InvalidParameter.PersonNotExisted"
//  INVALIDPARAMETER_PERSONNOTREGISTERED = "InvalidParameter.PersonNotRegistered"
//  INVALIDPARAMETER_PROCESSUNREADY = "InvalidParameter.ProcessUnReady"
//  INVALIDPARAMETER_RESOURCENAMEDUPLICATE = "InvalidParameter.ResourceNameDuplicate"
//  INVALIDPARAMETER_STATUSJOBFAILED = "InvalidParameter.StatusJobFailed"
//  INVALIDPARAMETER_STATUSJOBNOTFOUND = "InvalidParameter.StatusJobNotFound"
//  INVALIDPARAMETER_STATUSJOBUNFINISHED = "InvalidParameter.StatusJobUnfinished"
//  INVALIDPARAMETER_STATUSNOCAMERAFOUND = "InvalidParameter.StatusNoCameraFound"
//  INVALIDPARAMETER_TOOMANYFACE = "InvalidParameter.TooManyFace"
//  INVALIDPARAMETER_TRACKINGFAILED = "InvalidParameter.TrackingFailed"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSEXCEED = "InvalidParameterValue.GroupExDescriptionsExceed"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMEIDENTICAL = "InvalidParameterValue.GroupExDescriptionsNameIdentical"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMEILLEGAL = "InvalidParameterValue.GroupExDescriptionsNameIllegal"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMETOOLONG = "InvalidParameterValue.GroupExDescriptionsNameTooLong"
//  LIMITEXCEEDED_STATUSCAMERACOUNTOVERMAX = "LimitExceeded.StatusCameraCountOverMax"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_HLJOBIDNOTFOUND = "ResourceNotFound.HLJobIdNotFound"
//  RESOURCEUNAVAILABLE_STATUSQUERYSTREAMFAILED = "ResourceUnavailable.StatusQueryStreamFailed"
func (c *Client) SubmitOneByOneClassTask(request *SubmitOneByOneClassTaskRequest) (response *SubmitOneByOneClassTaskResponse, err error) {
    return c.SubmitOneByOneClassTaskWithContext(context.Background(), request)
}

// SubmitOneByOneClassTask
// **提交在线1对1课堂任务**  
//
// 对于在线1对1课堂，老师通过视频向学生授课，并且学生人数为1人。通过上传学生端的图像信息，可以获取学生的听课情况分析。 具体指一路全局画面且背景不动，有1位学生的头像或上半身的画面，要求画面稳定清晰。
//
//   
//
// **提供的功能接口有：**学生人脸识别、学生表情识别、语音识别。可分析的指标维度包括：学生身份识别、正脸、侧脸、抬头、低头、人脸坐标、人脸尺寸、高兴、中性、高兴、中性、惊讶、厌恶、恐惧、愤怒、蔑视、悲伤、语音转文字、发音时长、非发音时长、音量、语速等。
//
//   
//
// **对场景的要求为：**真实常规1v1授课场景，学生2人以下，全局画面且背景不动；人脸上下角度在20度以内，左右角度在15度以内，歪头角度在15度以内；光照均匀，无遮挡，人脸清晰可见；像素最好在 100X100 像素以上，但是图像整体质量不能超过1080p。
//
//     
//
// **结果查询方式：**图像任务直接返回结果，点播及直播任务通过DescribeAITaskResult查询结果。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_JOBCANCELED = "InternalError.JobCanceled"
//  INTERNALERROR_JOBREACHMAXRETRYTIMES = "InternalError.JobReachMaxRetryTimes"
//  INTERNALERROR_STATUSQUERYFACELIBFAILED = "InternalError.StatusQueryFaceLibFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_BODYRETRIEVALFAILED = "InvalidParameter.BodyRetrievalFailed"
//  INVALIDPARAMETER_CANNOTFINDFACEINPICTURE = "InvalidParameter.CannotFindFaceInPicture"
//  INVALIDPARAMETER_CLIPPINGFRAMEFAILED = "InvalidParameter.ClippingFrameFailed"
//  INVALIDPARAMETER_COMPUTATIONALSIMILARITYFAILED = "InvalidParameter.ComputationalSimilarityFailed"
//  INVALIDPARAMETER_CREATEFACEFAILED = "InvalidParameter.CreateFaceFailed"
//  INVALIDPARAMETER_CREATEPERSONFAILED = "InvalidParameter.CreatePersonFailed"
//  INVALIDPARAMETER_FACEDETECTEDFAILED = "InvalidParameter.FaceDetectedFailed"
//  INVALIDPARAMETER_FACENOTEXIST = "InvalidParameter.FaceNotExist"
//  INVALIDPARAMETER_FACENOTEXISTED = "InvalidParameter.FaceNotExisted"
//  INVALIDPARAMETER_FACENUMBEREXCEEDLIMITOFSINGLEDETECTION = "InvalidParameter.FaceNumberExceedLimitOfSingleDetection"
//  INVALIDPARAMETER_FACENUMBERLIMIT = "InvalidParameter.FaceNumberLimit"
//  INVALIDPARAMETER_FACEREGISTRATIONFAILED = "InvalidParameter.FaceRegistrationFailed"
//  INVALIDPARAMETER_FACERETRIEVALFAILED = "InvalidParameter.FaceRetrievalFailed"
//  INVALIDPARAMETER_FACETOOSMALL = "InvalidParameter.FaceTooSmall"
//  INVALIDPARAMETER_FAILEDTODECODEIMAGE = "InvalidParameter.FailedToDecodeImage"
//  INVALIDPARAMETER_FEATUREEXTRACTIONFAILED = "InvalidParameter.FeatureExtractionFailed"
//  INVALIDPARAMETER_FILECONTENTEMPTY = "InvalidParameter.FileContentEmpty"
//  INVALIDPARAMETER_IMAGEDECODEFAILED = "InvalidParameter.ImageDecodeFailed"
//  INVALIDPARAMETER_IMAGETOOLARGE = "InvalidParameter.ImageTooLarge"
//  INVALIDPARAMETER_IMAGETOOSMALL = "InvalidParameter.ImageTooSmall"
//  INVALIDPARAMETER_INVALIDFILETYPE = "InvalidParameter.InvalidFileType"
//  INVALIDPARAMETER_INVALIDLANG = "InvalidParameter.InvalidLang"
//  INVALIDPARAMETER_INVALIDRESOURCEIDENTIFY = "InvalidParameter.InvalidResourceIdentify"
//  INVALIDPARAMETER_INVALIDRESOURCENAME = "InvalidParameter.InvalidResourceName"
//  INVALIDPARAMETER_INVALIDSEQID = "InvalidParameter.InvalidSeqId"
//  INVALIDPARAMETER_INVALIDURL = "InvalidParameter.InvalidUrl"
//  INVALIDPARAMETER_JOBNOTVALID = "InvalidParameter.JobNotValid"
//  INVALIDPARAMETER_LIBRARYNOTEXISTED = "InvalidParameter.LibraryNotExisted"
//  INVALIDPARAMETER_LOWQUALITYPICTURE = "InvalidParameter.LowQualityPicture"
//  INVALIDPARAMETER_NOSUCHSERVICE = "InvalidParameter.NoSuchService"
//  INVALIDPARAMETER_NOTATTENDANCEJOB = "InvalidParameter.NotAttendanceJob"
//  INVALIDPARAMETER_NOTVALIDPITCH = "InvalidParameter.NotValidPitch"
//  INVALIDPARAMETER_NOTVALIDROW = "InvalidParameter.NotValidRow"
//  INVALIDPARAMETER_NOTVALIDYAW = "InvalidParameter.NotValidYaw"
//  INVALIDPARAMETER_PERSONEXISTED = "InvalidParameter.PersonExisted"
//  INVALIDPARAMETER_PERSONIDNOTVALID = "InvalidParameter.PersonIdNotValid"
//  INVALIDPARAMETER_PERSONNOTEXIST = "InvalidParameter.PersonNotExist"
//  INVALIDPARAMETER_PERSONNOTEXISTED = "InvalidParameter.PersonNotExisted"
//  INVALIDPARAMETER_PERSONNOTREGISTERED = "InvalidParameter.PersonNotRegistered"
//  INVALIDPARAMETER_PROCESSUNREADY = "InvalidParameter.ProcessUnReady"
//  INVALIDPARAMETER_RESOURCENAMEDUPLICATE = "InvalidParameter.ResourceNameDuplicate"
//  INVALIDPARAMETER_STATUSJOBFAILED = "InvalidParameter.StatusJobFailed"
//  INVALIDPARAMETER_STATUSJOBNOTFOUND = "InvalidParameter.StatusJobNotFound"
//  INVALIDPARAMETER_STATUSJOBUNFINISHED = "InvalidParameter.StatusJobUnfinished"
//  INVALIDPARAMETER_STATUSNOCAMERAFOUND = "InvalidParameter.StatusNoCameraFound"
//  INVALIDPARAMETER_TOOMANYFACE = "InvalidParameter.TooManyFace"
//  INVALIDPARAMETER_TRACKINGFAILED = "InvalidParameter.TrackingFailed"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSEXCEED = "InvalidParameterValue.GroupExDescriptionsExceed"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMEIDENTICAL = "InvalidParameterValue.GroupExDescriptionsNameIdentical"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMEILLEGAL = "InvalidParameterValue.GroupExDescriptionsNameIllegal"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMETOOLONG = "InvalidParameterValue.GroupExDescriptionsNameTooLong"
//  LIMITEXCEEDED_STATUSCAMERACOUNTOVERMAX = "LimitExceeded.StatusCameraCountOverMax"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_HLJOBIDNOTFOUND = "ResourceNotFound.HLJobIdNotFound"
//  RESOURCEUNAVAILABLE_STATUSQUERYSTREAMFAILED = "ResourceUnavailable.StatusQueryStreamFailed"
func (c *Client) SubmitOneByOneClassTaskWithContext(ctx context.Context, request *SubmitOneByOneClassTaskRequest) (response *SubmitOneByOneClassTaskResponse, err error) {
    if request == nil {
        request = NewSubmitOneByOneClassTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SubmitOneByOneClassTask require credential")
    }

    request.SetContext(ctx)
    
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

// SubmitOpenClassTask
// **提交线下小班（无课桌）课任务**  
//
// 线下小班课是指有学生无课桌的课堂，满座15人以下，全局画面且背景不动，能清晰看到。  
//
//   
//
// **提供的功能接口有：**学生人脸识别、学生表情识别、学生肢体动作识别。  可分析的指标维度包括：身份识别、正脸、侧脸、抬头、低头、高兴、中性、高兴、中性、惊讶、厌恶、恐惧、愤怒、蔑视、悲伤、站立、举手、坐着等。
//
//   
//
// **对场景的要求为：**真实常规教室，满座15人以下，全局画面且背景不动；人脸上下角度在20度以内，左右角度在15度以内，歪头角度在15度以内；光照均匀，无遮挡，人脸清晰可见；像素最好在 100X100 像素以上但是图像整体质量不能超过1080p。
//
//     
//
// **结果查询方式：**图像任务直接返回结果，点播及直播任务通过DescribeAITaskResult查询结果。
//
// 可能返回的错误码:
//  INTERNALERROR_JOBCANCELED = "InternalError.JobCanceled"
//  INTERNALERROR_JOBREACHMAXRETRYTIMES = "InternalError.JobReachMaxRetryTimes"
//  INTERNALERROR_STATUSQUERYFACELIBFAILED = "InternalError.StatusQueryFaceLibFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_BODYRETRIEVALFAILED = "InvalidParameter.BodyRetrievalFailed"
//  INVALIDPARAMETER_CANNOTFINDFACEINPICTURE = "InvalidParameter.CannotFindFaceInPicture"
//  INVALIDPARAMETER_CLIPPINGFRAMEFAILED = "InvalidParameter.ClippingFrameFailed"
//  INVALIDPARAMETER_COMPUTATIONALSIMILARITYFAILED = "InvalidParameter.ComputationalSimilarityFailed"
//  INVALIDPARAMETER_CREATEFACEFAILED = "InvalidParameter.CreateFaceFailed"
//  INVALIDPARAMETER_CREATEPERSONFAILED = "InvalidParameter.CreatePersonFailed"
//  INVALIDPARAMETER_FACEDETECTEDFAILED = "InvalidParameter.FaceDetectedFailed"
//  INVALIDPARAMETER_FACENOTEXIST = "InvalidParameter.FaceNotExist"
//  INVALIDPARAMETER_FACENOTEXISTED = "InvalidParameter.FaceNotExisted"
//  INVALIDPARAMETER_FACENUMBEREXCEEDLIMITOFSINGLEDETECTION = "InvalidParameter.FaceNumberExceedLimitOfSingleDetection"
//  INVALIDPARAMETER_FACENUMBERLIMIT = "InvalidParameter.FaceNumberLimit"
//  INVALIDPARAMETER_FACEREGISTRATIONFAILED = "InvalidParameter.FaceRegistrationFailed"
//  INVALIDPARAMETER_FACERETRIEVALFAILED = "InvalidParameter.FaceRetrievalFailed"
//  INVALIDPARAMETER_FACETOOSMALL = "InvalidParameter.FaceTooSmall"
//  INVALIDPARAMETER_FAILEDTODECODEIMAGE = "InvalidParameter.FailedToDecodeImage"
//  INVALIDPARAMETER_FEATUREEXTRACTIONFAILED = "InvalidParameter.FeatureExtractionFailed"
//  INVALIDPARAMETER_FILECONTENTEMPTY = "InvalidParameter.FileContentEmpty"
//  INVALIDPARAMETER_IMAGEDECODEFAILED = "InvalidParameter.ImageDecodeFailed"
//  INVALIDPARAMETER_IMAGETOOLARGE = "InvalidParameter.ImageTooLarge"
//  INVALIDPARAMETER_IMAGETOOSMALL = "InvalidParameter.ImageTooSmall"
//  INVALIDPARAMETER_INVALIDFILETYPE = "InvalidParameter.InvalidFileType"
//  INVALIDPARAMETER_INVALIDLANG = "InvalidParameter.InvalidLang"
//  INVALIDPARAMETER_INVALIDRESOURCEIDENTIFY = "InvalidParameter.InvalidResourceIdentify"
//  INVALIDPARAMETER_INVALIDRESOURCENAME = "InvalidParameter.InvalidResourceName"
//  INVALIDPARAMETER_INVALIDSEQID = "InvalidParameter.InvalidSeqId"
//  INVALIDPARAMETER_INVALIDURL = "InvalidParameter.InvalidUrl"
//  INVALIDPARAMETER_JOBNOTVALID = "InvalidParameter.JobNotValid"
//  INVALIDPARAMETER_LIBRARYNOTEXISTED = "InvalidParameter.LibraryNotExisted"
//  INVALIDPARAMETER_LOWQUALITYPICTURE = "InvalidParameter.LowQualityPicture"
//  INVALIDPARAMETER_NOSUCHSERVICE = "InvalidParameter.NoSuchService"
//  INVALIDPARAMETER_NOTATTENDANCEJOB = "InvalidParameter.NotAttendanceJob"
//  INVALIDPARAMETER_NOTVALIDPITCH = "InvalidParameter.NotValidPitch"
//  INVALIDPARAMETER_NOTVALIDROW = "InvalidParameter.NotValidRow"
//  INVALIDPARAMETER_NOTVALIDYAW = "InvalidParameter.NotValidYaw"
//  INVALIDPARAMETER_PERSONEXISTED = "InvalidParameter.PersonExisted"
//  INVALIDPARAMETER_PERSONIDNOTVALID = "InvalidParameter.PersonIdNotValid"
//  INVALIDPARAMETER_PERSONNOTEXIST = "InvalidParameter.PersonNotExist"
//  INVALIDPARAMETER_PERSONNOTEXISTED = "InvalidParameter.PersonNotExisted"
//  INVALIDPARAMETER_PERSONNOTREGISTERED = "InvalidParameter.PersonNotRegistered"
//  INVALIDPARAMETER_PROCESSUNREADY = "InvalidParameter.ProcessUnReady"
//  INVALIDPARAMETER_RESOURCENAMEDUPLICATE = "InvalidParameter.ResourceNameDuplicate"
//  INVALIDPARAMETER_STATUSJOBFAILED = "InvalidParameter.StatusJobFailed"
//  INVALIDPARAMETER_STATUSJOBNOTFOUND = "InvalidParameter.StatusJobNotFound"
//  INVALIDPARAMETER_STATUSJOBUNFINISHED = "InvalidParameter.StatusJobUnfinished"
//  INVALIDPARAMETER_STATUSNOCAMERAFOUND = "InvalidParameter.StatusNoCameraFound"
//  INVALIDPARAMETER_TOOMANYFACE = "InvalidParameter.TooManyFace"
//  INVALIDPARAMETER_TRACKINGFAILED = "InvalidParameter.TrackingFailed"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSEXCEED = "InvalidParameterValue.GroupExDescriptionsExceed"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMEIDENTICAL = "InvalidParameterValue.GroupExDescriptionsNameIdentical"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMEILLEGAL = "InvalidParameterValue.GroupExDescriptionsNameIllegal"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMETOOLONG = "InvalidParameterValue.GroupExDescriptionsNameTooLong"
//  LIMITEXCEEDED_STATUSCAMERACOUNTOVERMAX = "LimitExceeded.StatusCameraCountOverMax"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_HLJOBIDNOTFOUND = "ResourceNotFound.HLJobIdNotFound"
//  RESOURCEUNAVAILABLE_STATUSQUERYSTREAMFAILED = "ResourceUnavailable.StatusQueryStreamFailed"
func (c *Client) SubmitOpenClassTask(request *SubmitOpenClassTaskRequest) (response *SubmitOpenClassTaskResponse, err error) {
    return c.SubmitOpenClassTaskWithContext(context.Background(), request)
}

// SubmitOpenClassTask
// **提交线下小班（无课桌）课任务**  
//
// 线下小班课是指有学生无课桌的课堂，满座15人以下，全局画面且背景不动，能清晰看到。  
//
//   
//
// **提供的功能接口有：**学生人脸识别、学生表情识别、学生肢体动作识别。  可分析的指标维度包括：身份识别、正脸、侧脸、抬头、低头、高兴、中性、高兴、中性、惊讶、厌恶、恐惧、愤怒、蔑视、悲伤、站立、举手、坐着等。
//
//   
//
// **对场景的要求为：**真实常规教室，满座15人以下，全局画面且背景不动；人脸上下角度在20度以内，左右角度在15度以内，歪头角度在15度以内；光照均匀，无遮挡，人脸清晰可见；像素最好在 100X100 像素以上但是图像整体质量不能超过1080p。
//
//     
//
// **结果查询方式：**图像任务直接返回结果，点播及直播任务通过DescribeAITaskResult查询结果。
//
// 可能返回的错误码:
//  INTERNALERROR_JOBCANCELED = "InternalError.JobCanceled"
//  INTERNALERROR_JOBREACHMAXRETRYTIMES = "InternalError.JobReachMaxRetryTimes"
//  INTERNALERROR_STATUSQUERYFACELIBFAILED = "InternalError.StatusQueryFaceLibFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_BODYRETRIEVALFAILED = "InvalidParameter.BodyRetrievalFailed"
//  INVALIDPARAMETER_CANNOTFINDFACEINPICTURE = "InvalidParameter.CannotFindFaceInPicture"
//  INVALIDPARAMETER_CLIPPINGFRAMEFAILED = "InvalidParameter.ClippingFrameFailed"
//  INVALIDPARAMETER_COMPUTATIONALSIMILARITYFAILED = "InvalidParameter.ComputationalSimilarityFailed"
//  INVALIDPARAMETER_CREATEFACEFAILED = "InvalidParameter.CreateFaceFailed"
//  INVALIDPARAMETER_CREATEPERSONFAILED = "InvalidParameter.CreatePersonFailed"
//  INVALIDPARAMETER_FACEDETECTEDFAILED = "InvalidParameter.FaceDetectedFailed"
//  INVALIDPARAMETER_FACENOTEXIST = "InvalidParameter.FaceNotExist"
//  INVALIDPARAMETER_FACENOTEXISTED = "InvalidParameter.FaceNotExisted"
//  INVALIDPARAMETER_FACENUMBEREXCEEDLIMITOFSINGLEDETECTION = "InvalidParameter.FaceNumberExceedLimitOfSingleDetection"
//  INVALIDPARAMETER_FACENUMBERLIMIT = "InvalidParameter.FaceNumberLimit"
//  INVALIDPARAMETER_FACEREGISTRATIONFAILED = "InvalidParameter.FaceRegistrationFailed"
//  INVALIDPARAMETER_FACERETRIEVALFAILED = "InvalidParameter.FaceRetrievalFailed"
//  INVALIDPARAMETER_FACETOOSMALL = "InvalidParameter.FaceTooSmall"
//  INVALIDPARAMETER_FAILEDTODECODEIMAGE = "InvalidParameter.FailedToDecodeImage"
//  INVALIDPARAMETER_FEATUREEXTRACTIONFAILED = "InvalidParameter.FeatureExtractionFailed"
//  INVALIDPARAMETER_FILECONTENTEMPTY = "InvalidParameter.FileContentEmpty"
//  INVALIDPARAMETER_IMAGEDECODEFAILED = "InvalidParameter.ImageDecodeFailed"
//  INVALIDPARAMETER_IMAGETOOLARGE = "InvalidParameter.ImageTooLarge"
//  INVALIDPARAMETER_IMAGETOOSMALL = "InvalidParameter.ImageTooSmall"
//  INVALIDPARAMETER_INVALIDFILETYPE = "InvalidParameter.InvalidFileType"
//  INVALIDPARAMETER_INVALIDLANG = "InvalidParameter.InvalidLang"
//  INVALIDPARAMETER_INVALIDRESOURCEIDENTIFY = "InvalidParameter.InvalidResourceIdentify"
//  INVALIDPARAMETER_INVALIDRESOURCENAME = "InvalidParameter.InvalidResourceName"
//  INVALIDPARAMETER_INVALIDSEQID = "InvalidParameter.InvalidSeqId"
//  INVALIDPARAMETER_INVALIDURL = "InvalidParameter.InvalidUrl"
//  INVALIDPARAMETER_JOBNOTVALID = "InvalidParameter.JobNotValid"
//  INVALIDPARAMETER_LIBRARYNOTEXISTED = "InvalidParameter.LibraryNotExisted"
//  INVALIDPARAMETER_LOWQUALITYPICTURE = "InvalidParameter.LowQualityPicture"
//  INVALIDPARAMETER_NOSUCHSERVICE = "InvalidParameter.NoSuchService"
//  INVALIDPARAMETER_NOTATTENDANCEJOB = "InvalidParameter.NotAttendanceJob"
//  INVALIDPARAMETER_NOTVALIDPITCH = "InvalidParameter.NotValidPitch"
//  INVALIDPARAMETER_NOTVALIDROW = "InvalidParameter.NotValidRow"
//  INVALIDPARAMETER_NOTVALIDYAW = "InvalidParameter.NotValidYaw"
//  INVALIDPARAMETER_PERSONEXISTED = "InvalidParameter.PersonExisted"
//  INVALIDPARAMETER_PERSONIDNOTVALID = "InvalidParameter.PersonIdNotValid"
//  INVALIDPARAMETER_PERSONNOTEXIST = "InvalidParameter.PersonNotExist"
//  INVALIDPARAMETER_PERSONNOTEXISTED = "InvalidParameter.PersonNotExisted"
//  INVALIDPARAMETER_PERSONNOTREGISTERED = "InvalidParameter.PersonNotRegistered"
//  INVALIDPARAMETER_PROCESSUNREADY = "InvalidParameter.ProcessUnReady"
//  INVALIDPARAMETER_RESOURCENAMEDUPLICATE = "InvalidParameter.ResourceNameDuplicate"
//  INVALIDPARAMETER_STATUSJOBFAILED = "InvalidParameter.StatusJobFailed"
//  INVALIDPARAMETER_STATUSJOBNOTFOUND = "InvalidParameter.StatusJobNotFound"
//  INVALIDPARAMETER_STATUSJOBUNFINISHED = "InvalidParameter.StatusJobUnfinished"
//  INVALIDPARAMETER_STATUSNOCAMERAFOUND = "InvalidParameter.StatusNoCameraFound"
//  INVALIDPARAMETER_TOOMANYFACE = "InvalidParameter.TooManyFace"
//  INVALIDPARAMETER_TRACKINGFAILED = "InvalidParameter.TrackingFailed"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSEXCEED = "InvalidParameterValue.GroupExDescriptionsExceed"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMEIDENTICAL = "InvalidParameterValue.GroupExDescriptionsNameIdentical"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMEILLEGAL = "InvalidParameterValue.GroupExDescriptionsNameIllegal"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMETOOLONG = "InvalidParameterValue.GroupExDescriptionsNameTooLong"
//  LIMITEXCEEDED_STATUSCAMERACOUNTOVERMAX = "LimitExceeded.StatusCameraCountOverMax"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_HLJOBIDNOTFOUND = "ResourceNotFound.HLJobIdNotFound"
//  RESOURCEUNAVAILABLE_STATUSQUERYSTREAMFAILED = "ResourceUnavailable.StatusQueryStreamFailed"
func (c *Client) SubmitOpenClassTaskWithContext(ctx context.Context, request *SubmitOpenClassTaskRequest) (response *SubmitOpenClassTaskResponse, err error) {
    if request == nil {
        request = NewSubmitOpenClassTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SubmitOpenClassTask require credential")
    }

    request.SetContext(ctx)
    
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

// SubmitPartialBodyClassTask
// **在线小班课任务**：此场景是在线授课场景，老师一般为坐着授课，摄像头可以拍摄到老师的头部及上半身。拍摄视频为一路全局画面，且背景不动，要求画面稳定清晰。通过此接口可分析老师授课的行为及语音，以支持AI评教。    
//
//   
//
// **提供的功能接口有：**老师人脸识别、老师表情识别、老师手势识别、光线识别、语音识别。 可分析的指标维度包括：身份识别、正脸、侧脸、人脸坐标、人脸尺寸、高兴、中性、高兴、中性、惊讶、厌恶、恐惧、愤怒、蔑视、悲伤、点赞手势、听你说手势、听我说手势、拿教具行为、语音转文字、发音时长、非发音时长、音量、语速、指定关键词的使用等 
//
//   
//
// **对场景的要求为：**在线常规授课场景，全局画面且背景不动；人脸上下角度在20度以内，左右角度在15度以内，歪头角度在15度以内；光照均匀，无遮挡，人脸清晰可见；像素最好在 100X100 像素以上，但是图像整体质量不能超过1080p。
//
//     
//
// **结果查询方式：**图像任务直接返回结果，点播及直播任务通过DescribeAITaskResult查询结果。
//
// 可能返回的错误码:
//  INTERNALERROR_JOBCANCELED = "InternalError.JobCanceled"
//  INTERNALERROR_JOBREACHMAXRETRYTIMES = "InternalError.JobReachMaxRetryTimes"
//  INTERNALERROR_STATUSQUERYFACELIBFAILED = "InternalError.StatusQueryFaceLibFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_BODYRETRIEVALFAILED = "InvalidParameter.BodyRetrievalFailed"
//  INVALIDPARAMETER_CANNOTFINDFACEINPICTURE = "InvalidParameter.CannotFindFaceInPicture"
//  INVALIDPARAMETER_CLIPPINGFRAMEFAILED = "InvalidParameter.ClippingFrameFailed"
//  INVALIDPARAMETER_COMPUTATIONALSIMILARITYFAILED = "InvalidParameter.ComputationalSimilarityFailed"
//  INVALIDPARAMETER_CREATEFACEFAILED = "InvalidParameter.CreateFaceFailed"
//  INVALIDPARAMETER_CREATEPERSONFAILED = "InvalidParameter.CreatePersonFailed"
//  INVALIDPARAMETER_FACEDETECTEDFAILED = "InvalidParameter.FaceDetectedFailed"
//  INVALIDPARAMETER_FACENOTEXIST = "InvalidParameter.FaceNotExist"
//  INVALIDPARAMETER_FACENOTEXISTED = "InvalidParameter.FaceNotExisted"
//  INVALIDPARAMETER_FACENUMBEREXCEEDLIMITOFSINGLEDETECTION = "InvalidParameter.FaceNumberExceedLimitOfSingleDetection"
//  INVALIDPARAMETER_FACENUMBERLIMIT = "InvalidParameter.FaceNumberLimit"
//  INVALIDPARAMETER_FACEREGISTRATIONFAILED = "InvalidParameter.FaceRegistrationFailed"
//  INVALIDPARAMETER_FACERETRIEVALFAILED = "InvalidParameter.FaceRetrievalFailed"
//  INVALIDPARAMETER_FACETOOSMALL = "InvalidParameter.FaceTooSmall"
//  INVALIDPARAMETER_FAILEDTODECODEIMAGE = "InvalidParameter.FailedToDecodeImage"
//  INVALIDPARAMETER_FEATUREEXTRACTIONFAILED = "InvalidParameter.FeatureExtractionFailed"
//  INVALIDPARAMETER_FILECONTENTEMPTY = "InvalidParameter.FileContentEmpty"
//  INVALIDPARAMETER_IMAGEDECODEFAILED = "InvalidParameter.ImageDecodeFailed"
//  INVALIDPARAMETER_IMAGETOOLARGE = "InvalidParameter.ImageTooLarge"
//  INVALIDPARAMETER_IMAGETOOSMALL = "InvalidParameter.ImageTooSmall"
//  INVALIDPARAMETER_INVALIDFILETYPE = "InvalidParameter.InvalidFileType"
//  INVALIDPARAMETER_INVALIDLANG = "InvalidParameter.InvalidLang"
//  INVALIDPARAMETER_INVALIDRESOURCEIDENTIFY = "InvalidParameter.InvalidResourceIdentify"
//  INVALIDPARAMETER_INVALIDRESOURCENAME = "InvalidParameter.InvalidResourceName"
//  INVALIDPARAMETER_INVALIDSEQID = "InvalidParameter.InvalidSeqId"
//  INVALIDPARAMETER_INVALIDURL = "InvalidParameter.InvalidUrl"
//  INVALIDPARAMETER_JOBNOTVALID = "InvalidParameter.JobNotValid"
//  INVALIDPARAMETER_LIBRARYNOTEXISTED = "InvalidParameter.LibraryNotExisted"
//  INVALIDPARAMETER_LOWQUALITYPICTURE = "InvalidParameter.LowQualityPicture"
//  INVALIDPARAMETER_NOSUCHSERVICE = "InvalidParameter.NoSuchService"
//  INVALIDPARAMETER_NOTATTENDANCEJOB = "InvalidParameter.NotAttendanceJob"
//  INVALIDPARAMETER_NOTVALIDPITCH = "InvalidParameter.NotValidPitch"
//  INVALIDPARAMETER_NOTVALIDROW = "InvalidParameter.NotValidRow"
//  INVALIDPARAMETER_NOTVALIDYAW = "InvalidParameter.NotValidYaw"
//  INVALIDPARAMETER_PERSONEXISTED = "InvalidParameter.PersonExisted"
//  INVALIDPARAMETER_PERSONIDNOTVALID = "InvalidParameter.PersonIdNotValid"
//  INVALIDPARAMETER_PERSONNOTEXIST = "InvalidParameter.PersonNotExist"
//  INVALIDPARAMETER_PERSONNOTEXISTED = "InvalidParameter.PersonNotExisted"
//  INVALIDPARAMETER_PERSONNOTREGISTERED = "InvalidParameter.PersonNotRegistered"
//  INVALIDPARAMETER_PROCESSUNREADY = "InvalidParameter.ProcessUnReady"
//  INVALIDPARAMETER_RESOURCENAMEDUPLICATE = "InvalidParameter.ResourceNameDuplicate"
//  INVALIDPARAMETER_STATUSJOBFAILED = "InvalidParameter.StatusJobFailed"
//  INVALIDPARAMETER_STATUSJOBNOTFOUND = "InvalidParameter.StatusJobNotFound"
//  INVALIDPARAMETER_STATUSJOBUNFINISHED = "InvalidParameter.StatusJobUnfinished"
//  INVALIDPARAMETER_STATUSNOCAMERAFOUND = "InvalidParameter.StatusNoCameraFound"
//  INVALIDPARAMETER_TOOMANYFACE = "InvalidParameter.TooManyFace"
//  INVALIDPARAMETER_TRACKINGFAILED = "InvalidParameter.TrackingFailed"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSEXCEED = "InvalidParameterValue.GroupExDescriptionsExceed"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMEIDENTICAL = "InvalidParameterValue.GroupExDescriptionsNameIdentical"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMEILLEGAL = "InvalidParameterValue.GroupExDescriptionsNameIllegal"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMETOOLONG = "InvalidParameterValue.GroupExDescriptionsNameTooLong"
//  INVALIDPARAMETERVALUE_IMAGEILLEGAL = "InvalidParameterValue.ImageIllegal"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONINFOSEXCEED = "InvalidParameterValue.PersonExDescriptionInfosExceed"
//  LIMITEXCEEDED_STATUSCAMERACOUNTOVERMAX = "LimitExceeded.StatusCameraCountOverMax"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_HLJOBIDNOTFOUND = "ResourceNotFound.HLJobIdNotFound"
//  RESOURCEUNAVAILABLE_STATUSQUERYSTREAMFAILED = "ResourceUnavailable.StatusQueryStreamFailed"
func (c *Client) SubmitPartialBodyClassTask(request *SubmitPartialBodyClassTaskRequest) (response *SubmitPartialBodyClassTaskResponse, err error) {
    return c.SubmitPartialBodyClassTaskWithContext(context.Background(), request)
}

// SubmitPartialBodyClassTask
// **在线小班课任务**：此场景是在线授课场景，老师一般为坐着授课，摄像头可以拍摄到老师的头部及上半身。拍摄视频为一路全局画面，且背景不动，要求画面稳定清晰。通过此接口可分析老师授课的行为及语音，以支持AI评教。    
//
//   
//
// **提供的功能接口有：**老师人脸识别、老师表情识别、老师手势识别、光线识别、语音识别。 可分析的指标维度包括：身份识别、正脸、侧脸、人脸坐标、人脸尺寸、高兴、中性、高兴、中性、惊讶、厌恶、恐惧、愤怒、蔑视、悲伤、点赞手势、听你说手势、听我说手势、拿教具行为、语音转文字、发音时长、非发音时长、音量、语速、指定关键词的使用等 
//
//   
//
// **对场景的要求为：**在线常规授课场景，全局画面且背景不动；人脸上下角度在20度以内，左右角度在15度以内，歪头角度在15度以内；光照均匀，无遮挡，人脸清晰可见；像素最好在 100X100 像素以上，但是图像整体质量不能超过1080p。
//
//     
//
// **结果查询方式：**图像任务直接返回结果，点播及直播任务通过DescribeAITaskResult查询结果。
//
// 可能返回的错误码:
//  INTERNALERROR_JOBCANCELED = "InternalError.JobCanceled"
//  INTERNALERROR_JOBREACHMAXRETRYTIMES = "InternalError.JobReachMaxRetryTimes"
//  INTERNALERROR_STATUSQUERYFACELIBFAILED = "InternalError.StatusQueryFaceLibFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_BODYRETRIEVALFAILED = "InvalidParameter.BodyRetrievalFailed"
//  INVALIDPARAMETER_CANNOTFINDFACEINPICTURE = "InvalidParameter.CannotFindFaceInPicture"
//  INVALIDPARAMETER_CLIPPINGFRAMEFAILED = "InvalidParameter.ClippingFrameFailed"
//  INVALIDPARAMETER_COMPUTATIONALSIMILARITYFAILED = "InvalidParameter.ComputationalSimilarityFailed"
//  INVALIDPARAMETER_CREATEFACEFAILED = "InvalidParameter.CreateFaceFailed"
//  INVALIDPARAMETER_CREATEPERSONFAILED = "InvalidParameter.CreatePersonFailed"
//  INVALIDPARAMETER_FACEDETECTEDFAILED = "InvalidParameter.FaceDetectedFailed"
//  INVALIDPARAMETER_FACENOTEXIST = "InvalidParameter.FaceNotExist"
//  INVALIDPARAMETER_FACENOTEXISTED = "InvalidParameter.FaceNotExisted"
//  INVALIDPARAMETER_FACENUMBEREXCEEDLIMITOFSINGLEDETECTION = "InvalidParameter.FaceNumberExceedLimitOfSingleDetection"
//  INVALIDPARAMETER_FACENUMBERLIMIT = "InvalidParameter.FaceNumberLimit"
//  INVALIDPARAMETER_FACEREGISTRATIONFAILED = "InvalidParameter.FaceRegistrationFailed"
//  INVALIDPARAMETER_FACERETRIEVALFAILED = "InvalidParameter.FaceRetrievalFailed"
//  INVALIDPARAMETER_FACETOOSMALL = "InvalidParameter.FaceTooSmall"
//  INVALIDPARAMETER_FAILEDTODECODEIMAGE = "InvalidParameter.FailedToDecodeImage"
//  INVALIDPARAMETER_FEATUREEXTRACTIONFAILED = "InvalidParameter.FeatureExtractionFailed"
//  INVALIDPARAMETER_FILECONTENTEMPTY = "InvalidParameter.FileContentEmpty"
//  INVALIDPARAMETER_IMAGEDECODEFAILED = "InvalidParameter.ImageDecodeFailed"
//  INVALIDPARAMETER_IMAGETOOLARGE = "InvalidParameter.ImageTooLarge"
//  INVALIDPARAMETER_IMAGETOOSMALL = "InvalidParameter.ImageTooSmall"
//  INVALIDPARAMETER_INVALIDFILETYPE = "InvalidParameter.InvalidFileType"
//  INVALIDPARAMETER_INVALIDLANG = "InvalidParameter.InvalidLang"
//  INVALIDPARAMETER_INVALIDRESOURCEIDENTIFY = "InvalidParameter.InvalidResourceIdentify"
//  INVALIDPARAMETER_INVALIDRESOURCENAME = "InvalidParameter.InvalidResourceName"
//  INVALIDPARAMETER_INVALIDSEQID = "InvalidParameter.InvalidSeqId"
//  INVALIDPARAMETER_INVALIDURL = "InvalidParameter.InvalidUrl"
//  INVALIDPARAMETER_JOBNOTVALID = "InvalidParameter.JobNotValid"
//  INVALIDPARAMETER_LIBRARYNOTEXISTED = "InvalidParameter.LibraryNotExisted"
//  INVALIDPARAMETER_LOWQUALITYPICTURE = "InvalidParameter.LowQualityPicture"
//  INVALIDPARAMETER_NOSUCHSERVICE = "InvalidParameter.NoSuchService"
//  INVALIDPARAMETER_NOTATTENDANCEJOB = "InvalidParameter.NotAttendanceJob"
//  INVALIDPARAMETER_NOTVALIDPITCH = "InvalidParameter.NotValidPitch"
//  INVALIDPARAMETER_NOTVALIDROW = "InvalidParameter.NotValidRow"
//  INVALIDPARAMETER_NOTVALIDYAW = "InvalidParameter.NotValidYaw"
//  INVALIDPARAMETER_PERSONEXISTED = "InvalidParameter.PersonExisted"
//  INVALIDPARAMETER_PERSONIDNOTVALID = "InvalidParameter.PersonIdNotValid"
//  INVALIDPARAMETER_PERSONNOTEXIST = "InvalidParameter.PersonNotExist"
//  INVALIDPARAMETER_PERSONNOTEXISTED = "InvalidParameter.PersonNotExisted"
//  INVALIDPARAMETER_PERSONNOTREGISTERED = "InvalidParameter.PersonNotRegistered"
//  INVALIDPARAMETER_PROCESSUNREADY = "InvalidParameter.ProcessUnReady"
//  INVALIDPARAMETER_RESOURCENAMEDUPLICATE = "InvalidParameter.ResourceNameDuplicate"
//  INVALIDPARAMETER_STATUSJOBFAILED = "InvalidParameter.StatusJobFailed"
//  INVALIDPARAMETER_STATUSJOBNOTFOUND = "InvalidParameter.StatusJobNotFound"
//  INVALIDPARAMETER_STATUSJOBUNFINISHED = "InvalidParameter.StatusJobUnfinished"
//  INVALIDPARAMETER_STATUSNOCAMERAFOUND = "InvalidParameter.StatusNoCameraFound"
//  INVALIDPARAMETER_TOOMANYFACE = "InvalidParameter.TooManyFace"
//  INVALIDPARAMETER_TRACKINGFAILED = "InvalidParameter.TrackingFailed"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSEXCEED = "InvalidParameterValue.GroupExDescriptionsExceed"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMEIDENTICAL = "InvalidParameterValue.GroupExDescriptionsNameIdentical"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMEILLEGAL = "InvalidParameterValue.GroupExDescriptionsNameIllegal"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMETOOLONG = "InvalidParameterValue.GroupExDescriptionsNameTooLong"
//  INVALIDPARAMETERVALUE_IMAGEILLEGAL = "InvalidParameterValue.ImageIllegal"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONINFOSEXCEED = "InvalidParameterValue.PersonExDescriptionInfosExceed"
//  LIMITEXCEEDED_STATUSCAMERACOUNTOVERMAX = "LimitExceeded.StatusCameraCountOverMax"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_HLJOBIDNOTFOUND = "ResourceNotFound.HLJobIdNotFound"
//  RESOURCEUNAVAILABLE_STATUSQUERYSTREAMFAILED = "ResourceUnavailable.StatusQueryStreamFailed"
func (c *Client) SubmitPartialBodyClassTaskWithContext(ctx context.Context, request *SubmitPartialBodyClassTaskRequest) (response *SubmitPartialBodyClassTaskResponse, err error) {
    if request == nil {
        request = NewSubmitPartialBodyClassTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SubmitPartialBodyClassTask require credential")
    }

    request.SetContext(ctx)
    
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

// SubmitTraditionalClassTask
// **提交线下传统面授大班课（含课桌）任务。**  
//
// 传统教室课堂是指有学生课堂有课桌的课堂，满座20-50人，全局画面且背景不动。  
//
//   
//
// **提供的功能接口有：**学生人脸识别、学生表情识别、学生肢体动作识别。可分析的指标维度包括：学生身份识别、正脸、侧脸、抬头、低头、高兴、中性、高兴、中性、惊讶、厌恶、恐惧、愤怒、蔑视、悲伤、举手、站立、坐着、趴桌子、玩手机等  
//
//   
//
// **对场景的要求为：**传统的学生上课教室，满座20-50人，全局画面且背景不动；人脸上下角度在20度以内，左右角度在15度以内，歪头角度在15度以内；光照均匀，无遮挡，人脸清晰可见；像素最好在 100X100 像素以上，但是图像整体质量不能超过1080p。
//
//     
//
// **结果查询方式：**图像任务直接返回结果，点播及直播任务通过DescribeAITaskResult查询结果。
//
//   
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_JOBCANCELED = "InternalError.JobCanceled"
//  INTERNALERROR_JOBREACHMAXRETRYTIMES = "InternalError.JobReachMaxRetryTimes"
//  INTERNALERROR_STATUSQUERYFACELIBFAILED = "InternalError.StatusQueryFaceLibFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_BODYRETRIEVALFAILED = "InvalidParameter.BodyRetrievalFailed"
//  INVALIDPARAMETER_CANNOTFINDFACEINPICTURE = "InvalidParameter.CannotFindFaceInPicture"
//  INVALIDPARAMETER_CLIPPINGFRAMEFAILED = "InvalidParameter.ClippingFrameFailed"
//  INVALIDPARAMETER_COMPUTATIONALSIMILARITYFAILED = "InvalidParameter.ComputationalSimilarityFailed"
//  INVALIDPARAMETER_CREATEFACEFAILED = "InvalidParameter.CreateFaceFailed"
//  INVALIDPARAMETER_CREATEPERSONFAILED = "InvalidParameter.CreatePersonFailed"
//  INVALIDPARAMETER_FACEDETECTEDFAILED = "InvalidParameter.FaceDetectedFailed"
//  INVALIDPARAMETER_FACENOTEXIST = "InvalidParameter.FaceNotExist"
//  INVALIDPARAMETER_FACENOTEXISTED = "InvalidParameter.FaceNotExisted"
//  INVALIDPARAMETER_FACENUMBEREXCEEDLIMITOFSINGLEDETECTION = "InvalidParameter.FaceNumberExceedLimitOfSingleDetection"
//  INVALIDPARAMETER_FACENUMBERLIMIT = "InvalidParameter.FaceNumberLimit"
//  INVALIDPARAMETER_FACEREGISTRATIONFAILED = "InvalidParameter.FaceRegistrationFailed"
//  INVALIDPARAMETER_FACERETRIEVALFAILED = "InvalidParameter.FaceRetrievalFailed"
//  INVALIDPARAMETER_FACETOOSMALL = "InvalidParameter.FaceTooSmall"
//  INVALIDPARAMETER_FAILEDTODECODEIMAGE = "InvalidParameter.FailedToDecodeImage"
//  INVALIDPARAMETER_FEATUREEXTRACTIONFAILED = "InvalidParameter.FeatureExtractionFailed"
//  INVALIDPARAMETER_FILECONTENTEMPTY = "InvalidParameter.FileContentEmpty"
//  INVALIDPARAMETER_IMAGEDECODEFAILED = "InvalidParameter.ImageDecodeFailed"
//  INVALIDPARAMETER_IMAGETOOLARGE = "InvalidParameter.ImageTooLarge"
//  INVALIDPARAMETER_IMAGETOOSMALL = "InvalidParameter.ImageTooSmall"
//  INVALIDPARAMETER_INVALIDFILETYPE = "InvalidParameter.InvalidFileType"
//  INVALIDPARAMETER_INVALIDLANG = "InvalidParameter.InvalidLang"
//  INVALIDPARAMETER_INVALIDRESOURCEIDENTIFY = "InvalidParameter.InvalidResourceIdentify"
//  INVALIDPARAMETER_INVALIDRESOURCENAME = "InvalidParameter.InvalidResourceName"
//  INVALIDPARAMETER_INVALIDSEQID = "InvalidParameter.InvalidSeqId"
//  INVALIDPARAMETER_INVALIDURL = "InvalidParameter.InvalidUrl"
//  INVALIDPARAMETER_JOBNOTVALID = "InvalidParameter.JobNotValid"
//  INVALIDPARAMETER_LIBRARYNOTEXISTED = "InvalidParameter.LibraryNotExisted"
//  INVALIDPARAMETER_LOWQUALITYPICTURE = "InvalidParameter.LowQualityPicture"
//  INVALIDPARAMETER_NOSUCHSERVICE = "InvalidParameter.NoSuchService"
//  INVALIDPARAMETER_NOTATTENDANCEJOB = "InvalidParameter.NotAttendanceJob"
//  INVALIDPARAMETER_NOTVALIDPITCH = "InvalidParameter.NotValidPitch"
//  INVALIDPARAMETER_NOTVALIDROW = "InvalidParameter.NotValidRow"
//  INVALIDPARAMETER_NOTVALIDYAW = "InvalidParameter.NotValidYaw"
//  INVALIDPARAMETER_PERSONEXISTED = "InvalidParameter.PersonExisted"
//  INVALIDPARAMETER_PERSONIDNOTVALID = "InvalidParameter.PersonIdNotValid"
//  INVALIDPARAMETER_PERSONNOTEXIST = "InvalidParameter.PersonNotExist"
//  INVALIDPARAMETER_PERSONNOTEXISTED = "InvalidParameter.PersonNotExisted"
//  INVALIDPARAMETER_PERSONNOTREGISTERED = "InvalidParameter.PersonNotRegistered"
//  INVALIDPARAMETER_PROCESSUNREADY = "InvalidParameter.ProcessUnReady"
//  INVALIDPARAMETER_RESOURCENAMEDUPLICATE = "InvalidParameter.ResourceNameDuplicate"
//  INVALIDPARAMETER_STATUSJOBFAILED = "InvalidParameter.StatusJobFailed"
//  INVALIDPARAMETER_STATUSJOBNOTFOUND = "InvalidParameter.StatusJobNotFound"
//  INVALIDPARAMETER_STATUSJOBUNFINISHED = "InvalidParameter.StatusJobUnfinished"
//  INVALIDPARAMETER_STATUSNOCAMERAFOUND = "InvalidParameter.StatusNoCameraFound"
//  INVALIDPARAMETER_TOOMANYFACE = "InvalidParameter.TooManyFace"
//  INVALIDPARAMETER_TRACKINGFAILED = "InvalidParameter.TrackingFailed"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSEXCEED = "InvalidParameterValue.GroupExDescriptionsExceed"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMEIDENTICAL = "InvalidParameterValue.GroupExDescriptionsNameIdentical"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMEILLEGAL = "InvalidParameterValue.GroupExDescriptionsNameIllegal"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMETOOLONG = "InvalidParameterValue.GroupExDescriptionsNameTooLong"
//  LIMITEXCEEDED_STATUSCAMERACOUNTOVERMAX = "LimitExceeded.StatusCameraCountOverMax"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_HLJOBIDNOTFOUND = "ResourceNotFound.HLJobIdNotFound"
//  RESOURCEUNAVAILABLE_STATUSQUERYSTREAMFAILED = "ResourceUnavailable.StatusQueryStreamFailed"
func (c *Client) SubmitTraditionalClassTask(request *SubmitTraditionalClassTaskRequest) (response *SubmitTraditionalClassTaskResponse, err error) {
    return c.SubmitTraditionalClassTaskWithContext(context.Background(), request)
}

// SubmitTraditionalClassTask
// **提交线下传统面授大班课（含课桌）任务。**  
//
// 传统教室课堂是指有学生课堂有课桌的课堂，满座20-50人，全局画面且背景不动。  
//
//   
//
// **提供的功能接口有：**学生人脸识别、学生表情识别、学生肢体动作识别。可分析的指标维度包括：学生身份识别、正脸、侧脸、抬头、低头、高兴、中性、高兴、中性、惊讶、厌恶、恐惧、愤怒、蔑视、悲伤、举手、站立、坐着、趴桌子、玩手机等  
//
//   
//
// **对场景的要求为：**传统的学生上课教室，满座20-50人，全局画面且背景不动；人脸上下角度在20度以内，左右角度在15度以内，歪头角度在15度以内；光照均匀，无遮挡，人脸清晰可见；像素最好在 100X100 像素以上，但是图像整体质量不能超过1080p。
//
//     
//
// **结果查询方式：**图像任务直接返回结果，点播及直播任务通过DescribeAITaskResult查询结果。
//
//   
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_JOBCANCELED = "InternalError.JobCanceled"
//  INTERNALERROR_JOBREACHMAXRETRYTIMES = "InternalError.JobReachMaxRetryTimes"
//  INTERNALERROR_STATUSQUERYFACELIBFAILED = "InternalError.StatusQueryFaceLibFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_BODYRETRIEVALFAILED = "InvalidParameter.BodyRetrievalFailed"
//  INVALIDPARAMETER_CANNOTFINDFACEINPICTURE = "InvalidParameter.CannotFindFaceInPicture"
//  INVALIDPARAMETER_CLIPPINGFRAMEFAILED = "InvalidParameter.ClippingFrameFailed"
//  INVALIDPARAMETER_COMPUTATIONALSIMILARITYFAILED = "InvalidParameter.ComputationalSimilarityFailed"
//  INVALIDPARAMETER_CREATEFACEFAILED = "InvalidParameter.CreateFaceFailed"
//  INVALIDPARAMETER_CREATEPERSONFAILED = "InvalidParameter.CreatePersonFailed"
//  INVALIDPARAMETER_FACEDETECTEDFAILED = "InvalidParameter.FaceDetectedFailed"
//  INVALIDPARAMETER_FACENOTEXIST = "InvalidParameter.FaceNotExist"
//  INVALIDPARAMETER_FACENOTEXISTED = "InvalidParameter.FaceNotExisted"
//  INVALIDPARAMETER_FACENUMBEREXCEEDLIMITOFSINGLEDETECTION = "InvalidParameter.FaceNumberExceedLimitOfSingleDetection"
//  INVALIDPARAMETER_FACENUMBERLIMIT = "InvalidParameter.FaceNumberLimit"
//  INVALIDPARAMETER_FACEREGISTRATIONFAILED = "InvalidParameter.FaceRegistrationFailed"
//  INVALIDPARAMETER_FACERETRIEVALFAILED = "InvalidParameter.FaceRetrievalFailed"
//  INVALIDPARAMETER_FACETOOSMALL = "InvalidParameter.FaceTooSmall"
//  INVALIDPARAMETER_FAILEDTODECODEIMAGE = "InvalidParameter.FailedToDecodeImage"
//  INVALIDPARAMETER_FEATUREEXTRACTIONFAILED = "InvalidParameter.FeatureExtractionFailed"
//  INVALIDPARAMETER_FILECONTENTEMPTY = "InvalidParameter.FileContentEmpty"
//  INVALIDPARAMETER_IMAGEDECODEFAILED = "InvalidParameter.ImageDecodeFailed"
//  INVALIDPARAMETER_IMAGETOOLARGE = "InvalidParameter.ImageTooLarge"
//  INVALIDPARAMETER_IMAGETOOSMALL = "InvalidParameter.ImageTooSmall"
//  INVALIDPARAMETER_INVALIDFILETYPE = "InvalidParameter.InvalidFileType"
//  INVALIDPARAMETER_INVALIDLANG = "InvalidParameter.InvalidLang"
//  INVALIDPARAMETER_INVALIDRESOURCEIDENTIFY = "InvalidParameter.InvalidResourceIdentify"
//  INVALIDPARAMETER_INVALIDRESOURCENAME = "InvalidParameter.InvalidResourceName"
//  INVALIDPARAMETER_INVALIDSEQID = "InvalidParameter.InvalidSeqId"
//  INVALIDPARAMETER_INVALIDURL = "InvalidParameter.InvalidUrl"
//  INVALIDPARAMETER_JOBNOTVALID = "InvalidParameter.JobNotValid"
//  INVALIDPARAMETER_LIBRARYNOTEXISTED = "InvalidParameter.LibraryNotExisted"
//  INVALIDPARAMETER_LOWQUALITYPICTURE = "InvalidParameter.LowQualityPicture"
//  INVALIDPARAMETER_NOSUCHSERVICE = "InvalidParameter.NoSuchService"
//  INVALIDPARAMETER_NOTATTENDANCEJOB = "InvalidParameter.NotAttendanceJob"
//  INVALIDPARAMETER_NOTVALIDPITCH = "InvalidParameter.NotValidPitch"
//  INVALIDPARAMETER_NOTVALIDROW = "InvalidParameter.NotValidRow"
//  INVALIDPARAMETER_NOTVALIDYAW = "InvalidParameter.NotValidYaw"
//  INVALIDPARAMETER_PERSONEXISTED = "InvalidParameter.PersonExisted"
//  INVALIDPARAMETER_PERSONIDNOTVALID = "InvalidParameter.PersonIdNotValid"
//  INVALIDPARAMETER_PERSONNOTEXIST = "InvalidParameter.PersonNotExist"
//  INVALIDPARAMETER_PERSONNOTEXISTED = "InvalidParameter.PersonNotExisted"
//  INVALIDPARAMETER_PERSONNOTREGISTERED = "InvalidParameter.PersonNotRegistered"
//  INVALIDPARAMETER_PROCESSUNREADY = "InvalidParameter.ProcessUnReady"
//  INVALIDPARAMETER_RESOURCENAMEDUPLICATE = "InvalidParameter.ResourceNameDuplicate"
//  INVALIDPARAMETER_STATUSJOBFAILED = "InvalidParameter.StatusJobFailed"
//  INVALIDPARAMETER_STATUSJOBNOTFOUND = "InvalidParameter.StatusJobNotFound"
//  INVALIDPARAMETER_STATUSJOBUNFINISHED = "InvalidParameter.StatusJobUnfinished"
//  INVALIDPARAMETER_STATUSNOCAMERAFOUND = "InvalidParameter.StatusNoCameraFound"
//  INVALIDPARAMETER_TOOMANYFACE = "InvalidParameter.TooManyFace"
//  INVALIDPARAMETER_TRACKINGFAILED = "InvalidParameter.TrackingFailed"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSEXCEED = "InvalidParameterValue.GroupExDescriptionsExceed"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMEIDENTICAL = "InvalidParameterValue.GroupExDescriptionsNameIdentical"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMEILLEGAL = "InvalidParameterValue.GroupExDescriptionsNameIllegal"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMETOOLONG = "InvalidParameterValue.GroupExDescriptionsNameTooLong"
//  LIMITEXCEEDED_STATUSCAMERACOUNTOVERMAX = "LimitExceeded.StatusCameraCountOverMax"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_HLJOBIDNOTFOUND = "ResourceNotFound.HLJobIdNotFound"
//  RESOURCEUNAVAILABLE_STATUSQUERYSTREAMFAILED = "ResourceUnavailable.StatusQueryStreamFailed"
func (c *Client) SubmitTraditionalClassTaskWithContext(ctx context.Context, request *SubmitTraditionalClassTaskRequest) (response *SubmitTraditionalClassTaskResponse, err error) {
    if request == nil {
        request = NewSubmitTraditionalClassTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SubmitTraditionalClassTask require credential")
    }

    request.SetContext(ctx)
    
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

// TransmitAudioStream
// 分析音频信息
//
// 可能返回的错误码:
//  INTERNALERROR_STATUSQUERYFACELIBFAILED = "InternalError.StatusQueryFaceLibFailed"
//  INVALIDPARAMETER_INVALIDLANG = "InvalidParameter.InvalidLang"
//  INVALIDPARAMETER_INVALIDSEQID = "InvalidParameter.InvalidSeqId"
//  INVALIDPARAMETER_PROCESSUNREADY = "InvalidParameter.ProcessUnReady"
//  INVALIDPARAMETER_STATUSJOBFAILED = "InvalidParameter.StatusJobFailed"
//  INVALIDPARAMETER_STATUSJOBNOTFOUND = "InvalidParameter.StatusJobNotFound"
//  INVALIDPARAMETER_STATUSJOBUNFINISHED = "InvalidParameter.StatusJobUnfinished"
//  INVALIDPARAMETER_STATUSNOCAMERAFOUND = "InvalidParameter.StatusNoCameraFound"
//  LIMITEXCEEDED_STATUSCAMERACOUNTOVERMAX = "LimitExceeded.StatusCameraCountOverMax"
//  RESOURCENOTFOUND_HLJOBIDNOTFOUND = "ResourceNotFound.HLJobIdNotFound"
//  RESOURCEUNAVAILABLE_STATUSQUERYSTREAMFAILED = "ResourceUnavailable.StatusQueryStreamFailed"
func (c *Client) TransmitAudioStream(request *TransmitAudioStreamRequest) (response *TransmitAudioStreamResponse, err error) {
    return c.TransmitAudioStreamWithContext(context.Background(), request)
}

// TransmitAudioStream
// 分析音频信息
//
// 可能返回的错误码:
//  INTERNALERROR_STATUSQUERYFACELIBFAILED = "InternalError.StatusQueryFaceLibFailed"
//  INVALIDPARAMETER_INVALIDLANG = "InvalidParameter.InvalidLang"
//  INVALIDPARAMETER_INVALIDSEQID = "InvalidParameter.InvalidSeqId"
//  INVALIDPARAMETER_PROCESSUNREADY = "InvalidParameter.ProcessUnReady"
//  INVALIDPARAMETER_STATUSJOBFAILED = "InvalidParameter.StatusJobFailed"
//  INVALIDPARAMETER_STATUSJOBNOTFOUND = "InvalidParameter.StatusJobNotFound"
//  INVALIDPARAMETER_STATUSJOBUNFINISHED = "InvalidParameter.StatusJobUnfinished"
//  INVALIDPARAMETER_STATUSNOCAMERAFOUND = "InvalidParameter.StatusNoCameraFound"
//  LIMITEXCEEDED_STATUSCAMERACOUNTOVERMAX = "LimitExceeded.StatusCameraCountOverMax"
//  RESOURCENOTFOUND_HLJOBIDNOTFOUND = "ResourceNotFound.HLJobIdNotFound"
//  RESOURCEUNAVAILABLE_STATUSQUERYSTREAMFAILED = "ResourceUnavailable.StatusQueryStreamFailed"
func (c *Client) TransmitAudioStreamWithContext(ctx context.Context, request *TransmitAudioStreamRequest) (response *TransmitAudioStreamResponse, err error) {
    if request == nil {
        request = NewTransmitAudioStreamRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("TransmitAudioStream require credential")
    }

    request.SetContext(ctx)
    
    response = NewTransmitAudioStreamResponse()
    err = c.Send(request, response)
    return
}
