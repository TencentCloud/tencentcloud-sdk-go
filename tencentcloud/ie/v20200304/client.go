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
    "context"
    "errors"
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

func NewClient(credential common.CredentialIface, region string, clientProfile *profile.ClientProfile) (client *Client, err error) {
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

// CreateEditingTask
// 创建编辑理解任务，可以同时选择视频标签识别、分类识别、智能拆条、智能集锦、智能封面和片头片尾识别中的一项或者多项能力。
//
// 可能返回的错误码:
//  FAILEDOPERATION_COSSTORAGEERROR = "FailedOperation.CosStorageError"
//  FAILEDOPERATION_RUNNINGTASKEXCEED = "FailedOperation.RunningTaskExceed"
//  FAILEDOPERATION_SERVERBUSY = "FailedOperation.ServerBusy"
//  FAILEDOPERATION_TASKRESUBMIT = "FailedOperation.TaskResubmit"
//  FAILEDOPERATION_VIDEODOWNLOADERROR = "FailedOperation.VideoDownloadError"
//  FAILEDOPERATION_VIDEOPARSEERROR = "FailedOperation.VideoParseError"
//  FAILEDOPERATION_VIDEOSIZEEXCEED = "FailedOperation.VideoSizeExceed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_LIVESTREAMNOTSUPPORT = "InvalidParameter.LiveStreamNotSupport"
//  INVALIDPARAMETERVALUE_CALLBACKURLERROR = "InvalidParameterValue.CallbackUrlError"
//  INVALIDPARAMETERVALUE_CALLBACKURLNOTEXIST = "InvalidParameterValue.CallbackUrlNotExist"
//  INVALIDPARAMETERVALUE_COSAUTHMODEERROR = "InvalidParameterValue.CosAuthModeError"
//  INVALIDPARAMETERVALUE_COSHOSTEDIDNOTEXIST = "InvalidParameterValue.CosHostedIdNotExist"
//  INVALIDPARAMETERVALUE_DOWNINFOFORMATWRONG = "InvalidParameterValue.DownInfoFormatWrong"
//  INVALIDPARAMETERVALUE_DOWNINFOTYPEWRONG = "InvalidParameterValue.DownInfoTypeWrong"
//  INVALIDPARAMETERVALUE_SAVEINFONOTEXIST = "InvalidParameterValue.SaveInfoNotExist"
//  INVALIDPARAMETERVALUE_URLINFOURLERROR = "InvalidParameterValue.UrlInfoUrlError"
//  INVALIDPARAMETERVALUE_VIDEOFORMATERROR = "InvalidParameterValue.VideoFormatError"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateEditingTask(request *CreateEditingTaskRequest) (response *CreateEditingTaskResponse, err error) {
    return c.CreateEditingTaskWithContext(context.Background(), request)
}

// CreateEditingTask
// 创建编辑理解任务，可以同时选择视频标签识别、分类识别、智能拆条、智能集锦、智能封面和片头片尾识别中的一项或者多项能力。
//
// 可能返回的错误码:
//  FAILEDOPERATION_COSSTORAGEERROR = "FailedOperation.CosStorageError"
//  FAILEDOPERATION_RUNNINGTASKEXCEED = "FailedOperation.RunningTaskExceed"
//  FAILEDOPERATION_SERVERBUSY = "FailedOperation.ServerBusy"
//  FAILEDOPERATION_TASKRESUBMIT = "FailedOperation.TaskResubmit"
//  FAILEDOPERATION_VIDEODOWNLOADERROR = "FailedOperation.VideoDownloadError"
//  FAILEDOPERATION_VIDEOPARSEERROR = "FailedOperation.VideoParseError"
//  FAILEDOPERATION_VIDEOSIZEEXCEED = "FailedOperation.VideoSizeExceed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_LIVESTREAMNOTSUPPORT = "InvalidParameter.LiveStreamNotSupport"
//  INVALIDPARAMETERVALUE_CALLBACKURLERROR = "InvalidParameterValue.CallbackUrlError"
//  INVALIDPARAMETERVALUE_CALLBACKURLNOTEXIST = "InvalidParameterValue.CallbackUrlNotExist"
//  INVALIDPARAMETERVALUE_COSAUTHMODEERROR = "InvalidParameterValue.CosAuthModeError"
//  INVALIDPARAMETERVALUE_COSHOSTEDIDNOTEXIST = "InvalidParameterValue.CosHostedIdNotExist"
//  INVALIDPARAMETERVALUE_DOWNINFOFORMATWRONG = "InvalidParameterValue.DownInfoFormatWrong"
//  INVALIDPARAMETERVALUE_DOWNINFOTYPEWRONG = "InvalidParameterValue.DownInfoTypeWrong"
//  INVALIDPARAMETERVALUE_SAVEINFONOTEXIST = "InvalidParameterValue.SaveInfoNotExist"
//  INVALIDPARAMETERVALUE_URLINFOURLERROR = "InvalidParameterValue.UrlInfoUrlError"
//  INVALIDPARAMETERVALUE_VIDEOFORMATERROR = "InvalidParameterValue.VideoFormatError"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateEditingTaskWithContext(ctx context.Context, request *CreateEditingTaskRequest) (response *CreateEditingTaskResponse, err error) {
    if request == nil {
        request = NewCreateEditingTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateEditingTask require credential")
    }

    request.SetContext(ctx)
    
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

// CreateMediaProcessTask
// 用于创建编辑处理任务，如媒体截取、媒体编辑、媒体拼接、媒体字幕。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_RUNNINGTASKEXCEED = "FailedOperation.RunningTaskExceed"
//  FAILEDOPERATION_SERVERBUSY = "FailedOperation.ServerBusy"
//  FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"
//  FAILEDOPERATION_VIDEODOWNLOADERROR = "FailedOperation.VideoDownloadError"
//  FAILEDOPERATION_VIDEOPARSEERROR = "FailedOperation.VideoParseError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_CALLBACKURLERROR = "InvalidParameterValue.CallbackUrlError"
//  INVALIDPARAMETERVALUE_DOWNINFOFORMATWRONG = "InvalidParameterValue.DownInfoFormatWrong"
//  INVALIDPARAMETERVALUE_DOWNINFOTYPEWRONG = "InvalidParameterValue.DownInfoTypeWrong"
//  INVALIDPARAMETERVALUE_LIVESOURCENOTSUPPORT = "InvalidParameterValue.LiveSourceNotSupport"
//  INVALIDPARAMETERVALUE_TASKIDNOTEXIST = "InvalidParameterValue.TaskIdNotExist"
//  INVALIDPARAMETERVALUE_URLINFOURLERROR = "InvalidParameterValue.UrlInfoUrlError"
//  INVALIDPARAMETERVALUE_VIDEOFORMATERROR = "InvalidParameterValue.VideoFormatError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
func (c *Client) CreateMediaProcessTask(request *CreateMediaProcessTaskRequest) (response *CreateMediaProcessTaskResponse, err error) {
    return c.CreateMediaProcessTaskWithContext(context.Background(), request)
}

// CreateMediaProcessTask
// 用于创建编辑处理任务，如媒体截取、媒体编辑、媒体拼接、媒体字幕。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_RUNNINGTASKEXCEED = "FailedOperation.RunningTaskExceed"
//  FAILEDOPERATION_SERVERBUSY = "FailedOperation.ServerBusy"
//  FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"
//  FAILEDOPERATION_VIDEODOWNLOADERROR = "FailedOperation.VideoDownloadError"
//  FAILEDOPERATION_VIDEOPARSEERROR = "FailedOperation.VideoParseError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_CALLBACKURLERROR = "InvalidParameterValue.CallbackUrlError"
//  INVALIDPARAMETERVALUE_DOWNINFOFORMATWRONG = "InvalidParameterValue.DownInfoFormatWrong"
//  INVALIDPARAMETERVALUE_DOWNINFOTYPEWRONG = "InvalidParameterValue.DownInfoTypeWrong"
//  INVALIDPARAMETERVALUE_LIVESOURCENOTSUPPORT = "InvalidParameterValue.LiveSourceNotSupport"
//  INVALIDPARAMETERVALUE_TASKIDNOTEXIST = "InvalidParameterValue.TaskIdNotExist"
//  INVALIDPARAMETERVALUE_URLINFOURLERROR = "InvalidParameterValue.UrlInfoUrlError"
//  INVALIDPARAMETERVALUE_VIDEOFORMATERROR = "InvalidParameterValue.VideoFormatError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
func (c *Client) CreateMediaProcessTaskWithContext(ctx context.Context, request *CreateMediaProcessTaskRequest) (response *CreateMediaProcessTaskResponse, err error) {
    if request == nil {
        request = NewCreateMediaProcessTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateMediaProcessTask require credential")
    }

    request.SetContext(ctx)
    
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

// CreateMediaQualityRestorationTask
// 创建画质重生任务，对视频进行转码、去噪、去划痕、去毛刺、超分、细节增强和色彩增强。
//
// 可能返回的错误码:
//  FAILEDOPERATION_RUNNINGTASKEXCEED = "FailedOperation.RunningTaskExceed"
//  FAILEDOPERATION_SERVERBUSY = "FailedOperation.ServerBusy"
//  FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_CALLBACKURLERROR = "InvalidParameterValue.CallbackUrlError"
//  INVALIDPARAMETERVALUE_DOWNINFOTYPEWRONG = "InvalidParameterValue.DownInfoTypeWrong"
//  INVALIDPARAMETERVALUE_TASKALREADYDONE = "InvalidParameterValue.TaskAlreadyDone"
//  INVALIDPARAMETERVALUE_TASKDELETED = "InvalidParameterValue.TaskDeleted"
//  INVALIDPARAMETERVALUE_TASKIDNOTEXIST = "InvalidParameterValue.TaskIdNotExist"
//  INVALIDPARAMETERVALUE_URIERROR = "InvalidParameterValue.UriError"
//  INVALIDPARAMETERVALUE_URLINFOURLERROR = "InvalidParameterValue.UrlInfoUrlError"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateMediaQualityRestorationTask(request *CreateMediaQualityRestorationTaskRequest) (response *CreateMediaQualityRestorationTaskResponse, err error) {
    return c.CreateMediaQualityRestorationTaskWithContext(context.Background(), request)
}

// CreateMediaQualityRestorationTask
// 创建画质重生任务，对视频进行转码、去噪、去划痕、去毛刺、超分、细节增强和色彩增强。
//
// 可能返回的错误码:
//  FAILEDOPERATION_RUNNINGTASKEXCEED = "FailedOperation.RunningTaskExceed"
//  FAILEDOPERATION_SERVERBUSY = "FailedOperation.ServerBusy"
//  FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_CALLBACKURLERROR = "InvalidParameterValue.CallbackUrlError"
//  INVALIDPARAMETERVALUE_DOWNINFOTYPEWRONG = "InvalidParameterValue.DownInfoTypeWrong"
//  INVALIDPARAMETERVALUE_TASKALREADYDONE = "InvalidParameterValue.TaskAlreadyDone"
//  INVALIDPARAMETERVALUE_TASKDELETED = "InvalidParameterValue.TaskDeleted"
//  INVALIDPARAMETERVALUE_TASKIDNOTEXIST = "InvalidParameterValue.TaskIdNotExist"
//  INVALIDPARAMETERVALUE_URIERROR = "InvalidParameterValue.UriError"
//  INVALIDPARAMETERVALUE_URLINFOURLERROR = "InvalidParameterValue.UrlInfoUrlError"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateMediaQualityRestorationTaskWithContext(ctx context.Context, request *CreateMediaQualityRestorationTaskRequest) (response *CreateMediaQualityRestorationTaskResponse, err error) {
    if request == nil {
        request = NewCreateMediaQualityRestorationTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateMediaQualityRestorationTask require credential")
    }

    request.SetContext(ctx)
    
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

// CreateQualityControlTask
// 通过接口可以智能检测视频画面中抖动重影、模糊、低光照、过曝光、黑边、白边、黑屏、白屏、花屏、噪点、马赛克、二维码等在内的多个场景，还可以自动检测视频无音频异常、无声音片段。
//
// 可能返回的错误码:
//  FAILEDOPERATION_RUNNINGTASKEXCEED = "FailedOperation.RunningTaskExceed"
//  FAILEDOPERATION_SERVERBUSY = "FailedOperation.ServerBusy"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_CALLBACKURLERROR = "InvalidParameterValue.CallbackUrlError"
//  INVALIDPARAMETERVALUE_COSAUTHMODEERROR = "InvalidParameterValue.CosAuthModeError"
//  INVALIDPARAMETERVALUE_COSHOSTEDIDNOTEXIST = "InvalidParameterValue.CosHostedIdNotExist"
//  INVALIDPARAMETERVALUE_DOWNINFOFORMATWRONG = "InvalidParameterValue.DownInfoFormatWrong"
//  INVALIDPARAMETERVALUE_DOWNINFOTYPEWRONG = "InvalidParameterValue.DownInfoTypeWrong"
//  INVALIDPARAMETERVALUE_URLINFOURLERROR = "InvalidParameterValue.UrlInfoUrlError"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateQualityControlTask(request *CreateQualityControlTaskRequest) (response *CreateQualityControlTaskResponse, err error) {
    return c.CreateQualityControlTaskWithContext(context.Background(), request)
}

// CreateQualityControlTask
// 通过接口可以智能检测视频画面中抖动重影、模糊、低光照、过曝光、黑边、白边、黑屏、白屏、花屏、噪点、马赛克、二维码等在内的多个场景，还可以自动检测视频无音频异常、无声音片段。
//
// 可能返回的错误码:
//  FAILEDOPERATION_RUNNINGTASKEXCEED = "FailedOperation.RunningTaskExceed"
//  FAILEDOPERATION_SERVERBUSY = "FailedOperation.ServerBusy"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_CALLBACKURLERROR = "InvalidParameterValue.CallbackUrlError"
//  INVALIDPARAMETERVALUE_COSAUTHMODEERROR = "InvalidParameterValue.CosAuthModeError"
//  INVALIDPARAMETERVALUE_COSHOSTEDIDNOTEXIST = "InvalidParameterValue.CosHostedIdNotExist"
//  INVALIDPARAMETERVALUE_DOWNINFOFORMATWRONG = "InvalidParameterValue.DownInfoFormatWrong"
//  INVALIDPARAMETERVALUE_DOWNINFOTYPEWRONG = "InvalidParameterValue.DownInfoTypeWrong"
//  INVALIDPARAMETERVALUE_URLINFOURLERROR = "InvalidParameterValue.UrlInfoUrlError"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateQualityControlTaskWithContext(ctx context.Context, request *CreateQualityControlTaskRequest) (response *CreateQualityControlTaskResponse, err error) {
    if request == nil {
        request = NewCreateQualityControlTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateQualityControlTask require credential")
    }

    request.SetContext(ctx)
    
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

// DescribeEditingTaskResult
// 获取编辑理解任务结果。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_TASKIDNOTEXIST = "InvalidParameterValue.TaskIdNotExist"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeEditingTaskResult(request *DescribeEditingTaskResultRequest) (response *DescribeEditingTaskResultResponse, err error) {
    return c.DescribeEditingTaskResultWithContext(context.Background(), request)
}

// DescribeEditingTaskResult
// 获取编辑理解任务结果。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_TASKIDNOTEXIST = "InvalidParameterValue.TaskIdNotExist"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeEditingTaskResultWithContext(ctx context.Context, request *DescribeEditingTaskResultRequest) (response *DescribeEditingTaskResultResponse, err error) {
    if request == nil {
        request = NewDescribeEditingTaskResultRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeEditingTaskResult require credential")
    }

    request.SetContext(ctx)
    
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

// DescribeMediaProcessTaskResult
// 用于获取编辑处理任务的结果。
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_TASKIDNOTEXIST = "InvalidParameterValue.TaskIdNotExist"
func (c *Client) DescribeMediaProcessTaskResult(request *DescribeMediaProcessTaskResultRequest) (response *DescribeMediaProcessTaskResultResponse, err error) {
    return c.DescribeMediaProcessTaskResultWithContext(context.Background(), request)
}

// DescribeMediaProcessTaskResult
// 用于获取编辑处理任务的结果。
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_TASKIDNOTEXIST = "InvalidParameterValue.TaskIdNotExist"
func (c *Client) DescribeMediaProcessTaskResultWithContext(ctx context.Context, request *DescribeMediaProcessTaskResultRequest) (response *DescribeMediaProcessTaskResultResponse, err error) {
    if request == nil {
        request = NewDescribeMediaProcessTaskResultRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMediaProcessTaskResult require credential")
    }

    request.SetContext(ctx)
    
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

// DescribeMediaQualityRestorationTaskRusult
// 获取画质重生任务结果，查看结束后的文件信息
//
// 可能返回的错误码:
//  FAILEDOPERATION_COSSTORAGEERROR = "FailedOperation.CosStorageError"
//  FAILEDOPERATION_EDITERROR = "FailedOperation.EditError"
//  FAILEDOPERATION_ENCODEFORMATERROR = "FailedOperation.EncodeFormatError"
//  FAILEDOPERATION_SEGMENTERROR = "FailedOperation.SegmentError"
//  FAILEDOPERATION_SERVERBUSY = "FailedOperation.ServerBusy"
//  FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"
//  FAILEDOPERATION_TRANSCODEERROR = "FailedOperation.TranscodeError"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnknowError"
//  FAILEDOPERATION_VIDEODOWNLOADERROR = "FailedOperation.VideoDownloadError"
//  FAILEDOPERATION_VIDEOPARSEERROR = "FailedOperation.VideoParseError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_TASKALREADYDONE = "InvalidParameterValue.TaskAlreadyDone"
//  INVALIDPARAMETERVALUE_TASKDELETED = "InvalidParameterValue.TaskDeleted"
//  INVALIDPARAMETERVALUE_TASKIDNOTEXIST = "InvalidParameterValue.TaskIdNotExist"
//  INVALIDPARAMETERVALUE_URIERROR = "InvalidParameterValue.UriError"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeMediaQualityRestorationTaskRusult(request *DescribeMediaQualityRestorationTaskRusultRequest) (response *DescribeMediaQualityRestorationTaskRusultResponse, err error) {
    return c.DescribeMediaQualityRestorationTaskRusultWithContext(context.Background(), request)
}

// DescribeMediaQualityRestorationTaskRusult
// 获取画质重生任务结果，查看结束后的文件信息
//
// 可能返回的错误码:
//  FAILEDOPERATION_COSSTORAGEERROR = "FailedOperation.CosStorageError"
//  FAILEDOPERATION_EDITERROR = "FailedOperation.EditError"
//  FAILEDOPERATION_ENCODEFORMATERROR = "FailedOperation.EncodeFormatError"
//  FAILEDOPERATION_SEGMENTERROR = "FailedOperation.SegmentError"
//  FAILEDOPERATION_SERVERBUSY = "FailedOperation.ServerBusy"
//  FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"
//  FAILEDOPERATION_TRANSCODEERROR = "FailedOperation.TranscodeError"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnknowError"
//  FAILEDOPERATION_VIDEODOWNLOADERROR = "FailedOperation.VideoDownloadError"
//  FAILEDOPERATION_VIDEOPARSEERROR = "FailedOperation.VideoParseError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_TASKALREADYDONE = "InvalidParameterValue.TaskAlreadyDone"
//  INVALIDPARAMETERVALUE_TASKDELETED = "InvalidParameterValue.TaskDeleted"
//  INVALIDPARAMETERVALUE_TASKIDNOTEXIST = "InvalidParameterValue.TaskIdNotExist"
//  INVALIDPARAMETERVALUE_URIERROR = "InvalidParameterValue.UriError"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeMediaQualityRestorationTaskRusultWithContext(ctx context.Context, request *DescribeMediaQualityRestorationTaskRusultRequest) (response *DescribeMediaQualityRestorationTaskRusultResponse, err error) {
    if request == nil {
        request = NewDescribeMediaQualityRestorationTaskRusultRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMediaQualityRestorationTaskRusult require credential")
    }

    request.SetContext(ctx)
    
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

// DescribeQualityControlTaskResult
// 获取媒体质检任务结果
//
// 可能返回的错误码:
//  FAILEDOPERATION_SERVERBUSY = "FailedOperation.ServerBusy"
//  FAILEDOPERATION_VIDEODOWNLOADERROR = "FailedOperation.VideoDownloadError"
//  FAILEDOPERATION_VIDEOPARSEERROR = "FailedOperation.VideoParseError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_CALLBACKURLNOTEXIST = "InvalidParameterValue.CallbackUrlNotExist"
//  INVALIDPARAMETERVALUE_COSAUTHMODEERROR = "InvalidParameterValue.CosAuthModeError"
//  INVALIDPARAMETERVALUE_COSHOSTEDIDNOTEXIST = "InvalidParameterValue.CosHostedIdNotExist"
//  INVALIDPARAMETERVALUE_DOWNINFOFORMATWRONG = "InvalidParameterValue.DownInfoFormatWrong"
//  INVALIDPARAMETERVALUE_DOWNINFOTYPEWRONG = "InvalidParameterValue.DownInfoTypeWrong"
//  INVALIDPARAMETERVALUE_TASKIDNOTEXIST = "InvalidParameterValue.TaskIdNotExist"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeQualityControlTaskResult(request *DescribeQualityControlTaskResultRequest) (response *DescribeQualityControlTaskResultResponse, err error) {
    return c.DescribeQualityControlTaskResultWithContext(context.Background(), request)
}

// DescribeQualityControlTaskResult
// 获取媒体质检任务结果
//
// 可能返回的错误码:
//  FAILEDOPERATION_SERVERBUSY = "FailedOperation.ServerBusy"
//  FAILEDOPERATION_VIDEODOWNLOADERROR = "FailedOperation.VideoDownloadError"
//  FAILEDOPERATION_VIDEOPARSEERROR = "FailedOperation.VideoParseError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_CALLBACKURLNOTEXIST = "InvalidParameterValue.CallbackUrlNotExist"
//  INVALIDPARAMETERVALUE_COSAUTHMODEERROR = "InvalidParameterValue.CosAuthModeError"
//  INVALIDPARAMETERVALUE_COSHOSTEDIDNOTEXIST = "InvalidParameterValue.CosHostedIdNotExist"
//  INVALIDPARAMETERVALUE_DOWNINFOFORMATWRONG = "InvalidParameterValue.DownInfoFormatWrong"
//  INVALIDPARAMETERVALUE_DOWNINFOTYPEWRONG = "InvalidParameterValue.DownInfoTypeWrong"
//  INVALIDPARAMETERVALUE_TASKIDNOTEXIST = "InvalidParameterValue.TaskIdNotExist"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeQualityControlTaskResultWithContext(ctx context.Context, request *DescribeQualityControlTaskResultRequest) (response *DescribeQualityControlTaskResultResponse, err error) {
    if request == nil {
        request = NewDescribeQualityControlTaskResultRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeQualityControlTaskResult require credential")
    }

    request.SetContext(ctx)
    
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

// StopMediaProcessTask
// 用于停止正在进行中的编辑处理任务。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_ACTIONNOTSUPPORT = "InvalidParameterValue.ActionNotSupport"
//  INVALIDPARAMETERVALUE_TASKIDNOTEXIST = "InvalidParameterValue.TaskIdNotExist"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
func (c *Client) StopMediaProcessTask(request *StopMediaProcessTaskRequest) (response *StopMediaProcessTaskResponse, err error) {
    return c.StopMediaProcessTaskWithContext(context.Background(), request)
}

// StopMediaProcessTask
// 用于停止正在进行中的编辑处理任务。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_ACTIONNOTSUPPORT = "InvalidParameterValue.ActionNotSupport"
//  INVALIDPARAMETERVALUE_TASKIDNOTEXIST = "InvalidParameterValue.TaskIdNotExist"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
func (c *Client) StopMediaProcessTaskWithContext(ctx context.Context, request *StopMediaProcessTaskRequest) (response *StopMediaProcessTaskResponse, err error) {
    if request == nil {
        request = NewStopMediaProcessTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("StopMediaProcessTask require credential")
    }

    request.SetContext(ctx)
    
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

// StopMediaQualityRestorationTask
// 删除正在进行的画质重生任务
//
// 可能返回的错误码:
//  FAILEDOPERATION_SERVERBUSY = "FailedOperation.ServerBusy"
//  FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_TASKALREADYDONE = "InvalidParameterValue.TaskAlreadyDone"
//  INVALIDPARAMETERVALUE_TASKDELETED = "InvalidParameterValue.TaskDeleted"
//  INVALIDPARAMETERVALUE_TASKIDNOTEXIST = "InvalidParameterValue.TaskIdNotExist"
//  INVALIDPARAMETERVALUE_URIERROR = "InvalidParameterValue.UriError"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) StopMediaQualityRestorationTask(request *StopMediaQualityRestorationTaskRequest) (response *StopMediaQualityRestorationTaskResponse, err error) {
    return c.StopMediaQualityRestorationTaskWithContext(context.Background(), request)
}

// StopMediaQualityRestorationTask
// 删除正在进行的画质重生任务
//
// 可能返回的错误码:
//  FAILEDOPERATION_SERVERBUSY = "FailedOperation.ServerBusy"
//  FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_TASKALREADYDONE = "InvalidParameterValue.TaskAlreadyDone"
//  INVALIDPARAMETERVALUE_TASKDELETED = "InvalidParameterValue.TaskDeleted"
//  INVALIDPARAMETERVALUE_TASKIDNOTEXIST = "InvalidParameterValue.TaskIdNotExist"
//  INVALIDPARAMETERVALUE_URIERROR = "InvalidParameterValue.UriError"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) StopMediaQualityRestorationTaskWithContext(ctx context.Context, request *StopMediaQualityRestorationTaskRequest) (response *StopMediaQualityRestorationTaskResponse, err error) {
    if request == nil {
        request = NewStopMediaQualityRestorationTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("StopMediaQualityRestorationTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewStopMediaQualityRestorationTaskResponse()
    err = c.Send(request, response)
    return
}
