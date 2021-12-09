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

package v20210903

import (
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2021-09-03"

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


func NewCreateTaskRequest() (request *CreateTaskRequest) {
    request = &CreateTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ivld", APIVersion, "CreateTask")
    
    
    return
}

func NewCreateTaskResponse() (response *CreateTaskResponse) {
    response = &CreateTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateTask
// 创建智能标签任务。
//
// 
//
// 请注意，本接口为异步接口，**返回TaskId只代表任务创建成功，不代表任务执行成功**。
//
// 可能返回的错误码:
//  AUTHFAILURE_USERNOTFOUND = "AuthFailure.UserNotFound"
//  FAILEDOPERATION_AITEMPLATENOTEXIST = "FailedOperation.AiTemplateNotExist"
//  FAILEDOPERATION_MEDIANOTREADY = "FailedOperation.MediaNotReady"
//  FAILEDOPERATION_TASKALREADYEXIST = "FailedOperation.TaskAlreadyExist"
//  FAILEDOPERATION_TASKNOTFINISHED = "FailedOperation.TaskNotFinished"
//  INVALIDPARAMETER_INVALIDMEDIAID = "InvalidParameter.InvalidMediaId"
//  INVALIDPARAMETER_INVALIDNAME = "InvalidParameter.InvalidName"
//  INVALIDPARAMETER_INVALIDURL = "InvalidParameter.InvalidURL"
//  RESOURCENOTFOUND_MEDIANOTFOUND = "ResourceNotFound.MediaNotFound"
func (c *Client) CreateTask(request *CreateTaskRequest) (response *CreateTaskResponse, err error) {
    if request == nil {
        request = NewCreateTaskRequest()
    }
    
    response = NewCreateTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteMediaRequest() (request *DeleteMediaRequest) {
    request = &DeleteMediaRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ivld", APIVersion, "DeleteMedia")
    
    
    return
}

func NewDeleteMediaResponse() (response *DeleteMediaResponse) {
    response = &DeleteMediaResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteMedia
// 将MediaId对应的媒资文件从系统中删除。
//
// 
//
// **请注意，本接口仅删除媒资文件，媒资文件对应的视频分析结果不会被删除**。如您需要删除结构化分析结果，请调用DeleteTask接口。
//
// 可能返回的错误码:
//  FAILEDOPERATION_MEDIAALREADYEXIST = "FailedOperation.MediaAlreadyExist"
//  FAILEDOPERATION_MEDIAINUSE = "FailedOperation.MediaInUse"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDMEDIAID = "InvalidParameter.InvalidMediaId"
//  INVALIDPARAMETER_INVALIDMEDIASTATUS = "InvalidParameter.InvalidMediaStatus"
//  RESOURCENOTFOUND_MEDIANOTFOUND = "ResourceNotFound.MediaNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_MEDIANOTACCESSIBLE = "UnsupportedOperation.MediaNotAccessible"
func (c *Client) DeleteMedia(request *DeleteMediaRequest) (response *DeleteMediaResponse, err error) {
    if request == nil {
        request = NewDeleteMediaRequest()
    }
    
    response = NewDeleteMediaResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMediaRequest() (request *DescribeMediaRequest) {
    request = &DescribeMediaRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ivld", APIVersion, "DescribeMedia")
    
    
    return
}

func NewDescribeMediaResponse() (response *DescribeMediaResponse) {
    response = &DescribeMediaResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeMedia
// 描述媒资文件信息，包括媒资状态，分辨率，帧率等。
//
// 
//
// 如果媒资文件未完成导入，本接口将仅输出媒资文件的状态信息；导入完成后，本接口还将输出媒资文件的其他元信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADFAILED = "FailedOperation.DownloadFailed"
//  FAILEDOPERATION_MD5MISMATCH = "FailedOperation.MD5Mismatch"
//  FAILEDOPERATION_MEDIANOTREADY = "FailedOperation.MediaNotReady"
//  FAILEDOPERATION_TRANSCODEFAILED = "FailedOperation.TranscodeFailed"
//  INVALIDPARAMETER_INVALIDMEDIAID = "InvalidParameter.InvalidMediaId"
//  RESOURCENOTFOUND_MEDIANOTFOUND = "ResourceNotFound.MediaNotFound"
func (c *Client) DescribeMedia(request *DescribeMediaRequest) (response *DescribeMediaResponse, err error) {
    if request == nil {
        request = NewDescribeMediaRequest()
    }
    
    response = NewDescribeMediaResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMediasRequest() (request *DescribeMediasRequest) {
    request = &DescribeMediasRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ivld", APIVersion, "DescribeMedias")
    
    
    return
}

func NewDescribeMediasResponse() (response *DescribeMediasResponse) {
    response = &DescribeMediasResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeMedias
// 依照输入条件，描述命中的媒资文件信息，包括媒资状态，分辨率，帧率等。
//
// 
//
// 请注意，本接口最多支持同时描述**50**个媒资文件
//
// 
//
// 如果媒资文件未完成导入，本接口将仅输出媒资文件的状态信息；导入完成后，本接口还将输出媒资文件的其他元信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION_MEDIAALREADYEXIST = "FailedOperation.MediaAlreadyExist"
//  INVALIDPARAMETER_INVALIDMD5 = "InvalidParameter.InvalidMD5"
//  INVALIDPARAMETER_INVALIDMEDIAID = "InvalidParameter.InvalidMediaId"
//  INVALIDPARAMETER_INVALIDMEDIANAME = "InvalidParameter.InvalidMediaName"
//  INVALIDPARAMETER_INVALIDMEDIASTATUS = "InvalidParameter.InvalidMediaStatus"
//  INVALIDPARAMETER_INVALIDNAME = "InvalidParameter.InvalidName"
//  INVALIDPARAMETER_INVALIDPAGENUMBER = "InvalidParameter.InvalidPageNumber"
//  INVALIDPARAMETER_INVALIDPAGESIZE = "InvalidParameter.InvalidPageSize"
//  INVALIDPARAMETER_INVALIDSORTBY = "InvalidParameter.InvalidSortBy"
//  INVALIDPARAMETER_INVALIDSORTORDER = "InvalidParameter.InvalidSortOrder"
//  INVALIDPARAMETER_INVALIDURL = "InvalidParameter.InvalidURL"
//  INVALIDPARAMETER_INVALIDUIN = "InvalidParameter.InvalidUin"
//  INVALIDPARAMETER_NAMETOOLONG = "InvalidParameter.NameTooLong"
//  INVALIDPARAMETER_UNSUPPORTURL = "InvalidParameter.UnsupportURL"
func (c *Client) DescribeMedias(request *DescribeMediasRequest) (response *DescribeMediasResponse, err error) {
    if request == nil {
        request = NewDescribeMediasRequest()
    }
    
    response = NewDescribeMediasResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTaskRequest() (request *DescribeTaskRequest) {
    request = &DescribeTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ivld", APIVersion, "DescribeTask")
    
    
    return
}

func NewDescribeTaskResponse() (response *DescribeTaskResponse) {
    response = &DescribeTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeTask
// 描述智能标签任务进度。
//
// 
//
// 请注意，**此接口仅返回任务执行状态信息，不返回任务执行结果**
//
// 
//
// 可能返回的错误码:
//  AUTHFAILURE_USERNOTFOUND = "AuthFailure.UserNotFound"
//  INVALIDPARAMETER_INVALIDTASKID = "InvalidParameter.InvalidTaskId"
//  RESOURCENOTFOUND_TASKNOTFOUND = "ResourceNotFound.TaskNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDPRODUCT = "UnauthorizedOperation.UnauthorizedProduct"
func (c *Client) DescribeTask(request *DescribeTaskRequest) (response *DescribeTaskResponse, err error) {
    if request == nil {
        request = NewDescribeTaskRequest()
    }
    
    response = NewDescribeTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTaskDetailRequest() (request *DescribeTaskDetailRequest) {
    request = &DescribeTaskDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ivld", APIVersion, "DescribeTaskDetail")
    
    
    return
}

func NewDescribeTaskDetailResponse() (response *DescribeTaskDetailResponse) {
    response = &DescribeTaskDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeTaskDetail
// 描述任务信息，如果任务成功完成，还将返回任务结果
//
// 可能返回的错误码:
//  AUTHFAILURE_USERNOTFOUND = "AuthFailure.UserNotFound"
//  FAILEDOPERATION_TASKNOTFINISHED = "FailedOperation.TaskNotFinished"
//  INVALIDPARAMETER_INVALIDTASKID = "InvalidParameter.InvalidTaskId"
//  RESOURCENOTFOUND_TASKNOTFOUND = "ResourceNotFound.TaskNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDPRODUCT = "UnauthorizedOperation.UnauthorizedProduct"
func (c *Client) DescribeTaskDetail(request *DescribeTaskDetailRequest) (response *DescribeTaskDetailResponse, err error) {
    if request == nil {
        request = NewDescribeTaskDetailRequest()
    }
    
    response = NewDescribeTaskDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTasksRequest() (request *DescribeTasksRequest) {
    request = &DescribeTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ivld", APIVersion, "DescribeTasks")
    
    
    return
}

func NewDescribeTasksResponse() (response *DescribeTasksResponse) {
    response = &DescribeTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeTasks
// 依照输入条件，描述命中的任务信息，包括任务创建时间，处理时间信息等。
//
// 
//
// 请注意，本接口最多支持同时描述**50**个任务信息
//
// 可能返回的错误码:
//  AUTHFAILURE_USERNOTFOUND = "AuthFailure.UserNotFound"
//  FAILEDOPERATION_GETTASKLISTFAILED = "FailedOperation.GetTaskListFailed"
//  INVALIDPARAMETER_INVALIDPAGENUMBER = "InvalidParameter.InvalidPageNumber"
//  INVALIDPARAMETER_INVALIDPAGESIZE = "InvalidParameter.InvalidPageSize"
//  INVALIDPARAMETER_INVALIDSORTBY = "InvalidParameter.InvalidSortBy"
//  INVALIDPARAMETER_INVALIDTASKID = "InvalidParameter.InvalidTaskId"
//  INVALIDPARAMETER_INVALIDTASKNAME = "InvalidParameter.InvalidTaskName"
//  INVALIDPARAMETER_INVALIDTASKSTATUS = "InvalidParameter.InvalidTaskStatus"
//  INVALIDPARAMETER_INVALIDUIN = "InvalidParameter.InvalidUin"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDPRODUCT = "UnauthorizedOperation.UnauthorizedProduct"
func (c *Client) DescribeTasks(request *DescribeTasksRequest) (response *DescribeTasksResponse, err error) {
    if request == nil {
        request = NewDescribeTasksRequest()
    }
    
    response = NewDescribeTasksResponse()
    err = c.Send(request, response)
    return
}

func NewImportMediaRequest() (request *ImportMediaRequest) {
    request = &ImportMediaRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ivld", APIVersion, "ImportMedia")
    
    
    return
}

func NewImportMediaResponse() (response *ImportMediaResponse) {
    response = &ImportMediaResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ImportMedia
// 将URL指向的媒资视频文件导入系统之中。
//
// 
//
// **请注意，本接口为异步接口**。接口返回MediaId仅代表导入视频任务发起，不代表任务完成，您可调用读接口(DescribeMedia/DescribeMedias)接口查询MediaId对应的媒资文件的状态。
//
// 
//
// 当前URL只支持COS地址，其形式为`https://${Bucket}-${AppId}.cos.${Region}.myqcloud.com/${ObjectKey}`，其中`${Bucket}`为您的COS桶名称，Region为COS桶所在[可用区](https://cloud.tencent.com/document/product/213/6091)，`${ObjectKey}`为指向存储在COS桶内的待分析的视频的[ObjectKey](https://cloud.tencent.com/document/product/436/13324)
//
// 
//
// 分析完成后，本产品将在您的`${Bucket}`桶内创建名为`${ObjectKey}-${task-start-time}`的目录(`task-start-time`形式为1970-01-01T08:08:08)并将分析结果将回传回该目录，也即，结构化分析结果(包括图片，JSON等数据)将会写回`https://${Bucket}-${AppId}.cos.${Region}.myqcloud.com/${ObjectKey}-${task-start-time}`目录
//
// 可能返回的错误码:
//  AUTHFAILURE_USERNOTFOUND = "AuthFailure.UserNotFound"
//  FAILEDOPERATION_DBCONNECTIONERROR = "FailedOperation.DBConnectionError"
//  FAILEDOPERATION_DBOPERATIONERROR = "FailedOperation.DBOperationError"
//  FAILEDOPERATION_GETVIDEOMETADATAFAILED = "FailedOperation.GetVideoMetadataFailed"
//  FAILEDOPERATION_MD5MISMATCH = "FailedOperation.MD5Mismatch"
//  FAILEDOPERATION_MEDIAALREADYEXIST = "FailedOperation.MediaAlreadyExist"
//  FAILEDOPERATION_MEDIAINUSE = "FailedOperation.MediaInUse"
//  FAILEDOPERATION_MEDIANOTREADY = "FailedOperation.MediaNotReady"
//  FAILEDOPERATION_TASKALREADYEXIST = "FailedOperation.TaskAlreadyExist"
//  INTERNALERROR_INNERERROR = "InternalError.InnerError"
//  INVALIDPARAMETER_INVALIDFILEPATH = "InvalidParameter.InvalidFilePath"
//  INVALIDPARAMETER_INVALIDMD5 = "InvalidParameter.InvalidMD5"
//  INVALIDPARAMETER_INVALIDMEDIAID = "InvalidParameter.InvalidMediaId"
//  INVALIDPARAMETER_INVALIDMEDIANAME = "InvalidParameter.InvalidMediaName"
//  INVALIDPARAMETER_INVALIDMEDIASTATUS = "InvalidParameter.InvalidMediaStatus"
//  INVALIDPARAMETER_INVALIDNAME = "InvalidParameter.InvalidName"
//  INVALIDPARAMETER_INVALIDURL = "InvalidParameter.InvalidURL"
//  INVALIDPARAMETER_INVALIDUIN = "InvalidParameter.InvalidUin"
//  INVALIDPARAMETER_NAMETOOLONG = "InvalidParameter.NameTooLong"
//  INVALIDPARAMETER_UNSUPPORTURL = "InvalidParameter.UnsupportURL"
//  LIMITEXCEEDED_USAGELIMITEXCEEDED = "LimitExceeded.UsageLimitExceeded"
//  REQUESTLIMITEXCEEDED_CONCURRENCYOVERFLOW = "RequestLimitExceeded.ConcurrencyOverflow"
//  RESOURCENOTFOUND_MEDIANOTFOUND = "ResourceNotFound.MediaNotFound"
func (c *Client) ImportMedia(request *ImportMediaRequest) (response *ImportMediaResponse, err error) {
    if request == nil {
        request = NewImportMediaRequest()
    }
    
    response = NewImportMediaResponse()
    err = c.Send(request, response)
    return
}
