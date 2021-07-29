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

package v20201229

import (
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2020-12-29"

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


func NewCancelTaskRequest() (request *CancelTaskRequest) {
    request = &CancelTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ams", APIVersion, "CancelTask")
    return
}

func NewCancelTaskResponse() (response *CancelTaskResponse) {
    response = &CancelTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CancelTask
// 取消音频审核任务
//
// 
//
// <br>默认接口请求频率限制：**20次/秒**。
//
// 可能返回的错误码:
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CancelTask(request *CancelTaskRequest) (response *CancelTaskResponse, err error) {
    if request == nil {
        request = NewCancelTaskRequest()
    }
    response = NewCancelTaskResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAudioModerationSyncTaskRequest() (request *CreateAudioModerationSyncTaskRequest) {
    request = &CreateAudioModerationSyncTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ams", APIVersion, "CreateAudioModerationSyncTask")
    return
}

func NewCreateAudioModerationSyncTaskResponse() (response *CreateAudioModerationSyncTaskResponse) {
    response = &CreateAudioModerationSyncTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateAudioModerationSyncTask
// 本接口（CreateAudioModerationSyncTask） 用于提交短音频内容进行智能审核任务，使用前请您登录控制台开通音频内容安全服务。
//
// 
//
// 功能使用说明：
//
// 前往“内容安全控制台-音频内容安全”开启使用音频内容安全服务，首次开通可获得10小时免费调用时长；
//
// 
//
// 接口限制：
//
// - 音频文件大小支持：文件 < 5M;
//
// - 音频文件时长小于60s，超过60s音频调用则报错；
//
// - 音频码率类型支持：8Kbps - 16Kbps；
//
// - 音频文件支持格式：wav、mp3；
//
// - 接口仅限音频文件传入，视频文件传入请调用长音频异步接口；
//
// - 接口默认QPS为10，默认接口请求频率限制20次/秒，如需要更高的并发或请求频率，请工单咨询；
//
// - 接口超时为5s，每一次请求超过该时长会报错；
//
// 可能返回的错误码:
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_IMAGESIZETOOSMALL = "InvalidParameter.ImageSizeTooSmall"
//  INVALIDPARAMETER_INVALIDIMAGECONTENT = "InvalidParameter.InvalidImageContent"
//  INVALIDPARAMETER_PARAMETERERROR = "InvalidParameter.ParameterError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_EMPTYIMAGECONTENT = "InvalidParameterValue.EmptyImageContent"
//  INVALIDPARAMETERVALUE_IMAGESIZETOOSMALL = "InvalidParameterValue.ImageSizeTooSmall"
//  INVALIDPARAMETERVALUE_INVALIDCONTENT = "InvalidParameterValue.InvalidContent"
//  INVALIDPARAMETERVALUE_INVALIDDATAID = "InvalidParameterValue.InvalidDataId"
//  INVALIDPARAMETERVALUE_INVALIDFILECONTENTSIZE = "InvalidParameterValue.InvalidFileContentSize"
//  INVALIDPARAMETERVALUE_INVALIDIMAGECONTENT = "InvalidParameterValue.InvalidImageContent"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETER = "InvalidParameterValue.InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_INVALIDIMAGECONTENT = "ResourceUnavailable.InvalidImageContent"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateAudioModerationSyncTask(request *CreateAudioModerationSyncTaskRequest) (response *CreateAudioModerationSyncTaskResponse, err error) {
    if request == nil {
        request = NewCreateAudioModerationSyncTaskRequest()
    }
    response = NewCreateAudioModerationSyncTaskResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAudioModerationTaskRequest() (request *CreateAudioModerationTaskRequest) {
    request = &CreateAudioModerationTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ams", APIVersion, "CreateAudioModerationTask")
    return
}

func NewCreateAudioModerationTaskResponse() (response *CreateAudioModerationTaskResponse) {
    response = &CreateAudioModerationTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateAudioModerationTask
// 本接口（Audio Moderation）用于提交音频内容（包括音频文件或流地址）进行智能审核任务，使用前请您使用腾讯云主账号登录控制台[开通音频内容安全服务](https://console.cloud.tencent.com/cms/audio/package)并调整好对应的业务配置。<br>
//
// 
//
// ### 功能使用说明：
//
// - 前往“[内容安全控制台-音频内容安全](https://console.cloud.tencent.com/cms/audio/package)”开启使用音频内容安全服务，首次开通可获得**10小时**免费调用时长，有效期为1个月。
//
// - 默认接口请求频率限制：**20次/秒**；对于异步审核任务（点播音频），超出频率限制的请求会自动排入待审核队列，对于同步审核任务（直播音频），超出频率限制将会报错。
//
// 
//
// ### 接口功能说明：
//
// - 支持对音频流或音频文件进行检测，判断其中是否包含违规内容；
//
// - 支持设置回调地址 Callback 获取检测结果（对于已在审核的任务，最长回调时间为用户配置的**切片时长 + 2s**），或通过接口(查询音频检测结果)主动轮询获取检测结果；
//
// - 支持识别违规内容，包括：低俗、谩骂、色情、涉政、广告等场景；
//
// - 支持批量提交检测任务，检测任务列表**最多支持10个**。
//
// 
//
// ### 音频文件调用说明：
//
// - 音频文件大小支持：**文件 < 500M**；
//
// - 音频文件时长支持：**< 1小时**；
//
// - 音频码率类型支持：128 Kbps - 256 Kbps ；
//
// - 音频文件支持格式：wav、mp3、aac、flac、amr、3gp、 m4a、wma、ogg、ape；
//
// - （**当输入为视频文件时**）支持分离视频文件音轨，并对音频内容进行独立审核。
//
// 
//
// ### 音频流调用说明：
//
// - 音频流时长支持：**< 3小时**；
//
// - 音频码率类型支持：128 Kbps - 256 Kbps ；
//
// - 音频流支持的传输协议：RTMP、HTTP、HTTPS；
//
// - 音频流格式支持的类型：rtp、srtp、rtmp、rtmps、mmsh、 mmst、hls、http、tcp、https、m3u8；
//
// - （**当输入为视频流时**）支持提取视频流音轨，并对音频内容进行独立审核。
//
// 可能返回的错误码:
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateAudioModerationTask(request *CreateAudioModerationTaskRequest) (response *CreateAudioModerationTaskResponse, err error) {
    if request == nil {
        request = NewCreateAudioModerationTaskRequest()
    }
    response = NewCreateAudioModerationTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTaskDetailRequest() (request *DescribeTaskDetailRequest) {
    request = &DescribeTaskDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ams", APIVersion, "DescribeTaskDetail")
    return
}

func NewDescribeTaskDetailResponse() (response *DescribeTaskDetailResponse) {
    response = &DescribeTaskDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeTaskDetail
// 用于查看音频审核任务的详细信息
//
// 
//
// <br>默认接口请求频率限制：**100次/秒**。
//
// 可能返回的错误码:
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
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
    request.Init().WithApiInfo("ams", APIVersion, "DescribeTasks")
    return
}

func NewDescribeTasksResponse() (response *DescribeTasksResponse) {
    response = &DescribeTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeTasks
// 查看审核任务列表
//
// 
//
// <br>默认接口请求频率限制：**20次/秒**。
//
// 可能返回的错误码:
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeTasks(request *DescribeTasksRequest) (response *DescribeTasksResponse, err error) {
    if request == nil {
        request = NewDescribeTasksRequest()
    }
    response = NewDescribeTasksResponse()
    err = c.Send(request, response)
    return
}
