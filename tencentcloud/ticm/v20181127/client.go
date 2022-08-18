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

package v20181127

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2018-11-27"

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


func NewDescribeVideoTaskRequest() (request *DescribeVideoTaskRequest) {
    request = &DescribeVideoTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ticm", APIVersion, "DescribeVideoTask")
    
    
    return
}

func NewDescribeVideoTaskResponse() (response *DescribeVideoTaskResponse) {
    response = &DescribeVideoTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeVideoTask
// 提交完视频审核任务后，可以通过本接口来获取当前处理的进度和结果
//
// 可能返回的错误码:
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
func (c *Client) DescribeVideoTask(request *DescribeVideoTaskRequest) (response *DescribeVideoTaskResponse, err error) {
    return c.DescribeVideoTaskWithContext(context.Background(), request)
}

// DescribeVideoTask
// 提交完视频审核任务后，可以通过本接口来获取当前处理的进度和结果
//
// 可能返回的错误码:
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
func (c *Client) DescribeVideoTaskWithContext(ctx context.Context, request *DescribeVideoTaskRequest) (response *DescribeVideoTaskResponse, err error) {
    if request == nil {
        request = NewDescribeVideoTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVideoTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVideoTaskResponse()
    err = c.Send(request, response)
    return
}

func NewImageModerationRequest() (request *ImageModerationRequest) {
    request = &ImageModerationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ticm", APIVersion, "ImageModeration")
    
    
    return
}

func NewImageModerationResponse() (response *ImageModerationResponse) {
    response = &ImageModerationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ImageModeration
// 本接口提供多种维度的图像审核能力，支持色情和性感内容识别，政治人物和涉政敏感场景识别，以及暴恐人物、场景、旗帜标识等违禁内容的识别。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_INVOKECHARGEERROR = "FailedOperation.InvokeChargeError"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) ImageModeration(request *ImageModerationRequest) (response *ImageModerationResponse, err error) {
    return c.ImageModerationWithContext(context.Background(), request)
}

// ImageModeration
// 本接口提供多种维度的图像审核能力，支持色情和性感内容识别，政治人物和涉政敏感场景识别，以及暴恐人物、场景、旗帜标识等违禁内容的识别。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_INVOKECHARGEERROR = "FailedOperation.InvokeChargeError"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) ImageModerationWithContext(ctx context.Context, request *ImageModerationRequest) (response *ImageModerationResponse, err error) {
    if request == nil {
        request = NewImageModerationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ImageModeration require credential")
    }

    request.SetContext(ctx)
    
    response = NewImageModerationResponse()
    err = c.Send(request, response)
    return
}

func NewVideoModerationRequest() (request *VideoModerationRequest) {
    request = &VideoModerationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ticm", APIVersion, "VideoModeration")
    
    
    return
}

func NewVideoModerationResponse() (response *VideoModerationResponse) {
    response = &VideoModerationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// VideoModeration
// 本接口提供多种维度的视频审核能力，支持色情和性感内容识别，政治人物和涉政敏感场景识别，以及暴恐人物、场景、旗帜标识等违禁内容的识别。
//
// 可能返回的错误码:
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
func (c *Client) VideoModeration(request *VideoModerationRequest) (response *VideoModerationResponse, err error) {
    return c.VideoModerationWithContext(context.Background(), request)
}

// VideoModeration
// 本接口提供多种维度的视频审核能力，支持色情和性感内容识别，政治人物和涉政敏感场景识别，以及暴恐人物、场景、旗帜标识等违禁内容的识别。
//
// 可能返回的错误码:
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
func (c *Client) VideoModerationWithContext(ctx context.Context, request *VideoModerationRequest) (response *VideoModerationResponse, err error) {
    if request == nil {
        request = NewVideoModerationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("VideoModeration require credential")
    }

    request.SetContext(ctx)
    
    response = NewVideoModerationResponse()
    err = c.Send(request, response)
    return
}
