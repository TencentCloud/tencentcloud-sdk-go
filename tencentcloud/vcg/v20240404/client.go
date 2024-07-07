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

package v20240404

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2024-04-04"

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


func NewDescribeVideoStylizationJobRequest() (request *DescribeVideoStylizationJobRequest) {
    request = &DescribeVideoStylizationJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vcg", APIVersion, "DescribeVideoStylizationJob")
    
    
    return
}

func NewDescribeVideoStylizationJobResponse() (response *DescribeVideoStylizationJobResponse) {
    response = &DescribeVideoStylizationJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeVideoStylizationJob
// 用于查询视频风格化任务。视频风格化支持将输入视频生成特定风格的视频。生成后的视频画面风格多样、流畅自然，能够满足社交娱乐、互动营销、视频素材制作等场景的需求。
//
// 可能返回的错误码:
//  FAILEDOPERATION_TASKNOTEXIST = "FailedOperation.TaskNotExist"
//  FAILEDOPERATION_TASKSTATUSERROR = "FailedOperation.TaskStatusError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_PARAMETERVALUEERROR = "InvalidParameterValue.ParameterValueError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  REQUESTLIMITEXCEEDED_JOBNUMEXCEED = "RequestLimitExceeded.JobNumExceed"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
func (c *Client) DescribeVideoStylizationJob(request *DescribeVideoStylizationJobRequest) (response *DescribeVideoStylizationJobResponse, err error) {
    return c.DescribeVideoStylizationJobWithContext(context.Background(), request)
}

// DescribeVideoStylizationJob
// 用于查询视频风格化任务。视频风格化支持将输入视频生成特定风格的视频。生成后的视频画面风格多样、流畅自然，能够满足社交娱乐、互动营销、视频素材制作等场景的需求。
//
// 可能返回的错误码:
//  FAILEDOPERATION_TASKNOTEXIST = "FailedOperation.TaskNotExist"
//  FAILEDOPERATION_TASKSTATUSERROR = "FailedOperation.TaskStatusError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_PARAMETERVALUEERROR = "InvalidParameterValue.ParameterValueError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  REQUESTLIMITEXCEEDED_JOBNUMEXCEED = "RequestLimitExceeded.JobNumExceed"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
func (c *Client) DescribeVideoStylizationJobWithContext(ctx context.Context, request *DescribeVideoStylizationJobRequest) (response *DescribeVideoStylizationJobResponse, err error) {
    if request == nil {
        request = NewDescribeVideoStylizationJobRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVideoStylizationJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVideoStylizationJobResponse()
    err = c.Send(request, response)
    return
}

func NewSubmitVideoStylizationJobRequest() (request *SubmitVideoStylizationJobRequest) {
    request = &SubmitVideoStylizationJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vcg", APIVersion, "SubmitVideoStylizationJob")
    
    
    return
}

func NewSubmitVideoStylizationJobResponse() (response *SubmitVideoStylizationJobResponse) {
    response = &SubmitVideoStylizationJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SubmitVideoStylizationJob
// 用于提交视频风格化任务。支持将输入视频生成特定风格的视频。生成后的视频画面风格多样、流畅自然，能够满足社交娱乐、互动营销、视频素材制作等场景的需求。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownloadError"
//  FAILEDOPERATION_VIDEODECODEFAILED = "FailedOperation.VideoDecodeFailed"
//  INVALIDPARAMETERVALUE_INVALIDVIDEOASPECTRATIO = "InvalidParameterValue.InvalidVideoAspectRatio"
//  INVALIDPARAMETERVALUE_INVALIDVIDEODURATION = "InvalidParameterValue.InvalidVideoDuration"
//  INVALIDPARAMETERVALUE_INVALIDVIDEOFPS = "InvalidParameterValue.InvalidVideoFPS"
//  INVALIDPARAMETERVALUE_INVALIDVIDEOFORMAT = "InvalidParameterValue.InvalidVideoFormat"
//  INVALIDPARAMETERVALUE_INVALIDVIDEORESOLUTION = "InvalidParameterValue.InvalidVideoResolution"
//  INVALIDPARAMETERVALUE_PARAMETERVALUEERROR = "InvalidParameterValue.ParameterValueError"
//  INVALIDPARAMETERVALUE_STYLENOTEXIST = "InvalidParameterValue.StyleNotExist"
//  INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"
//  INVALIDPARAMETERVALUE_VIDEOSIZEEXCEED = "InvalidParameterValue.VideoSizeExceed"
//  REQUESTLIMITEXCEEDED_JOBNUMEXCEED = "RequestLimitExceeded.JobNumExceed"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
func (c *Client) SubmitVideoStylizationJob(request *SubmitVideoStylizationJobRequest) (response *SubmitVideoStylizationJobResponse, err error) {
    return c.SubmitVideoStylizationJobWithContext(context.Background(), request)
}

// SubmitVideoStylizationJob
// 用于提交视频风格化任务。支持将输入视频生成特定风格的视频。生成后的视频画面风格多样、流畅自然，能够满足社交娱乐、互动营销、视频素材制作等场景的需求。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownloadError"
//  FAILEDOPERATION_VIDEODECODEFAILED = "FailedOperation.VideoDecodeFailed"
//  INVALIDPARAMETERVALUE_INVALIDVIDEOASPECTRATIO = "InvalidParameterValue.InvalidVideoAspectRatio"
//  INVALIDPARAMETERVALUE_INVALIDVIDEODURATION = "InvalidParameterValue.InvalidVideoDuration"
//  INVALIDPARAMETERVALUE_INVALIDVIDEOFPS = "InvalidParameterValue.InvalidVideoFPS"
//  INVALIDPARAMETERVALUE_INVALIDVIDEOFORMAT = "InvalidParameterValue.InvalidVideoFormat"
//  INVALIDPARAMETERVALUE_INVALIDVIDEORESOLUTION = "InvalidParameterValue.InvalidVideoResolution"
//  INVALIDPARAMETERVALUE_PARAMETERVALUEERROR = "InvalidParameterValue.ParameterValueError"
//  INVALIDPARAMETERVALUE_STYLENOTEXIST = "InvalidParameterValue.StyleNotExist"
//  INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"
//  INVALIDPARAMETERVALUE_VIDEOSIZEEXCEED = "InvalidParameterValue.VideoSizeExceed"
//  REQUESTLIMITEXCEEDED_JOBNUMEXCEED = "RequestLimitExceeded.JobNumExceed"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
func (c *Client) SubmitVideoStylizationJobWithContext(ctx context.Context, request *SubmitVideoStylizationJobRequest) (response *SubmitVideoStylizationJobResponse, err error) {
    if request == nil {
        request = NewSubmitVideoStylizationJobRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SubmitVideoStylizationJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewSubmitVideoStylizationJobResponse()
    err = c.Send(request, response)
    return
}
