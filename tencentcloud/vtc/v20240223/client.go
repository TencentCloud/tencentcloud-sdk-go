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

package v20240223

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2024-02-23"

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


func NewConfirmVideoTranslateJobRequest() (request *ConfirmVideoTranslateJobRequest) {
    request = &ConfirmVideoTranslateJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vtc", APIVersion, "ConfirmVideoTranslateJob")
    
    
    return
}

func NewConfirmVideoTranslateJobResponse() (response *ConfirmVideoTranslateJobResponse) {
    response = &ConfirmVideoTranslateJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ConfirmVideoTranslateJob
// 确认视频翻译结果
//
// 可能返回的错误码:
//  FAILEDOPERATION_AUDIOPROCESSNOTFINISHED = "FailedOperation.AudioProcessNotFinished"
//  FAILEDOPERATION_JOBNOTEXIST = "FailedOperation.JobNotExist"
//  FAILEDOPERATION_TRANSLATIONCONFIRMHASFINISHED = "FailedOperation.TranslationConfirmHasFinished"
//  FAILEDOPERATION_TRANSLATIONNOTNEEDCONFIRM = "FailedOperation.TranslationNotNeedConfirm"
//  REQUESTLIMITEXCEEDED_JOBNUMEXCEED = "RequestLimitExceeded.JobNumExceed"
func (c *Client) ConfirmVideoTranslateJob(request *ConfirmVideoTranslateJobRequest) (response *ConfirmVideoTranslateJobResponse, err error) {
    return c.ConfirmVideoTranslateJobWithContext(context.Background(), request)
}

// ConfirmVideoTranslateJob
// 确认视频翻译结果
//
// 可能返回的错误码:
//  FAILEDOPERATION_AUDIOPROCESSNOTFINISHED = "FailedOperation.AudioProcessNotFinished"
//  FAILEDOPERATION_JOBNOTEXIST = "FailedOperation.JobNotExist"
//  FAILEDOPERATION_TRANSLATIONCONFIRMHASFINISHED = "FailedOperation.TranslationConfirmHasFinished"
//  FAILEDOPERATION_TRANSLATIONNOTNEEDCONFIRM = "FailedOperation.TranslationNotNeedConfirm"
//  REQUESTLIMITEXCEEDED_JOBNUMEXCEED = "RequestLimitExceeded.JobNumExceed"
func (c *Client) ConfirmVideoTranslateJobWithContext(ctx context.Context, request *ConfirmVideoTranslateJobRequest) (response *ConfirmVideoTranslateJobResponse, err error) {
    if request == nil {
        request = NewConfirmVideoTranslateJobRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ConfirmVideoTranslateJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewConfirmVideoTranslateJobResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVideoTranslateJobRequest() (request *DescribeVideoTranslateJobRequest) {
    request = &DescribeVideoTranslateJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vtc", APIVersion, "DescribeVideoTranslateJob")
    
    
    return
}

func NewDescribeVideoTranslateJobResponse() (response *DescribeVideoTranslateJobResponse) {
    response = &DescribeVideoTranslateJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeVideoTranslateJob
// 查询视频翻译任务
//
// 可能返回的错误码:
//  FAILEDOPERATION_AUDIOPROCESSFAILED = "FailedOperation.AudioProcessFailed"
//  FAILEDOPERATION_JOBNOTEXIST = "FailedOperation.JobNotExist"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  REQUESTLIMITEXCEEDED_UINLIMITEXCEEDED = "RequestLimitExceeded.UinLimitExceeded"
//  RESOURCEUNAVAILABLE_DELIVERING = "ResourceUnavailable.Delivering"
//  RESOURCEUNAVAILABLE_ISOPENING = "ResourceUnavailable.IsOpening"
//  RESOURCEUNAVAILABLE_LOWBALANCE = "ResourceUnavailable.LowBalance"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCEUNAVAILABLE_NOTREADY = "ResourceUnavailable.NotReady"
//  RESOURCEUNAVAILABLE_RECOVER = "ResourceUnavailable.Recover"
//  RESOURCEUNAVAILABLE_UNKNOWNSTATUS = "ResourceUnavailable.UnknownStatus"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) DescribeVideoTranslateJob(request *DescribeVideoTranslateJobRequest) (response *DescribeVideoTranslateJobResponse, err error) {
    return c.DescribeVideoTranslateJobWithContext(context.Background(), request)
}

// DescribeVideoTranslateJob
// 查询视频翻译任务
//
// 可能返回的错误码:
//  FAILEDOPERATION_AUDIOPROCESSFAILED = "FailedOperation.AudioProcessFailed"
//  FAILEDOPERATION_JOBNOTEXIST = "FailedOperation.JobNotExist"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  REQUESTLIMITEXCEEDED_UINLIMITEXCEEDED = "RequestLimitExceeded.UinLimitExceeded"
//  RESOURCEUNAVAILABLE_DELIVERING = "ResourceUnavailable.Delivering"
//  RESOURCEUNAVAILABLE_ISOPENING = "ResourceUnavailable.IsOpening"
//  RESOURCEUNAVAILABLE_LOWBALANCE = "ResourceUnavailable.LowBalance"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCEUNAVAILABLE_NOTREADY = "ResourceUnavailable.NotReady"
//  RESOURCEUNAVAILABLE_RECOVER = "ResourceUnavailable.Recover"
//  RESOURCEUNAVAILABLE_UNKNOWNSTATUS = "ResourceUnavailable.UnknownStatus"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) DescribeVideoTranslateJobWithContext(ctx context.Context, request *DescribeVideoTranslateJobRequest) (response *DescribeVideoTranslateJobResponse, err error) {
    if request == nil {
        request = NewDescribeVideoTranslateJobRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVideoTranslateJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVideoTranslateJobResponse()
    err = c.Send(request, response)
    return
}

func NewSubmitVideoTranslateJobRequest() (request *SubmitVideoTranslateJobRequest) {
    request = &SubmitVideoTranslateJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vtc", APIVersion, "SubmitVideoTranslateJob")
    
    
    return
}

func NewSubmitVideoTranslateJobResponse() (response *SubmitVideoTranslateJobResponse) {
    response = &SubmitVideoTranslateJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SubmitVideoTranslateJob
// 提交视频转译任务
//
// 可能返回的错误码:
//  FAILEDOPERATION_AUDIODECODEFAILED = "FailedOperation.AudioDecodeFailed"
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownloadError"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"
//  FAILEDOPERATION_VIDEODURATIONEXCEED = "FailedOperation.VideoDurationExceed"
//  FAILEDOPERATION_VIDEOFPSEXCEED = "FailedOperation.VideoFpsExceed"
//  FAILEDOPERATION_VIDEORESOLUTIONEXCEED = "FailedOperation.VideoResolutionExceed"
//  FAILEDOPERATION_VIDEOSIZEEXCEED = "FailedOperation.VideoSizeExceed"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"
//  REQUESTLIMITEXCEEDED_JOBNUMEXCEED = "RequestLimitExceeded.JobNumExceed"
//  RESOURCEUNAVAILABLE_FREEZE = "ResourceUnavailable.Freeze"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_ISOPENING = "ResourceUnavailable.IsOpening"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCEUNAVAILABLE_NOTREADY = "ResourceUnavailable.NotReady"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) SubmitVideoTranslateJob(request *SubmitVideoTranslateJobRequest) (response *SubmitVideoTranslateJobResponse, err error) {
    return c.SubmitVideoTranslateJobWithContext(context.Background(), request)
}

// SubmitVideoTranslateJob
// 提交视频转译任务
//
// 可能返回的错误码:
//  FAILEDOPERATION_AUDIODECODEFAILED = "FailedOperation.AudioDecodeFailed"
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownloadError"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"
//  FAILEDOPERATION_VIDEODURATIONEXCEED = "FailedOperation.VideoDurationExceed"
//  FAILEDOPERATION_VIDEOFPSEXCEED = "FailedOperation.VideoFpsExceed"
//  FAILEDOPERATION_VIDEORESOLUTIONEXCEED = "FailedOperation.VideoResolutionExceed"
//  FAILEDOPERATION_VIDEOSIZEEXCEED = "FailedOperation.VideoSizeExceed"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"
//  REQUESTLIMITEXCEEDED_JOBNUMEXCEED = "RequestLimitExceeded.JobNumExceed"
//  RESOURCEUNAVAILABLE_FREEZE = "ResourceUnavailable.Freeze"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_ISOPENING = "ResourceUnavailable.IsOpening"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCEUNAVAILABLE_NOTREADY = "ResourceUnavailable.NotReady"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) SubmitVideoTranslateJobWithContext(ctx context.Context, request *SubmitVideoTranslateJobRequest) (response *SubmitVideoTranslateJobResponse, err error) {
    if request == nil {
        request = NewSubmitVideoTranslateJobRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SubmitVideoTranslateJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewSubmitVideoTranslateJobResponse()
    err = c.Send(request, response)
    return
}
