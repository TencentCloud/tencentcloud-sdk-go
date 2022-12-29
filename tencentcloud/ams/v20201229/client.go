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
    "context"
    "errors"
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
// 可使用该接口取消审核任务。请求成功后，接口返回RequestId则说明取消成功。<br>默认接口请求频率限制：**20次/秒**。
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
    return c.CancelTaskWithContext(context.Background(), request)
}

// CancelTask
// 可使用该接口取消审核任务。请求成功后，接口返回RequestId则说明取消成功。<br>默认接口请求频率限制：**20次/秒**。
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
// 本接口（CreateAudioModerationSyncTask） 用于提交短音频内容进行智能审核任务，使用前请您使用腾讯云主账号登录控制台 [开通音频内容安全服务](https://console.cloud.tencent.com/cms/audio/package) 并调整好对应的业务配置。
//
// 
//
// ### 接口使用说明：
//
// - 前往“[内容安全控制台-图片内容安全](https://console.cloud.tencent.com/cms/audio/package)”开启使用音频内容安全服务，首次开通服务的用户可免费领用试用套餐包，包含**10小时**免费调用时长，有效期为1个月。
//
// - 该接口为收费接口，计费方式敬请参见 [腾讯云音频内容安全定价](https://cloud.tencent.com/product/ams/pricing)。
//
// 
//
// ### 接口调用说明：
//
// - 音频文件大小支持：**文件 <= 4M**;
//
// - 音频文件**时长不超过60s**，超过60s音频调用则报错；
//
// - 音频文件支持格式：**wav (PCM编码)** 、**mp3**、**aac**、**m4a** (采样率：16kHz~48kHz，位深：16bit 小端，声道数：单声道/双声道，建议格式：**16kHz/16bit/单声道**)；
//
// - 接口仅限音频文件传入，视频文件传入请调用长音频异步接口；
//
// - 接口**默认QPS为20**，如需自定义配置并发或请求频率，请工单咨询；
//
// - 接口**默认超时为10s**，请求如超过该时长则接口会报错。
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
    return c.CreateAudioModerationSyncTaskWithContext(context.Background(), request)
}

// CreateAudioModerationSyncTask
// 本接口（CreateAudioModerationSyncTask） 用于提交短音频内容进行智能审核任务，使用前请您使用腾讯云主账号登录控制台 [开通音频内容安全服务](https://console.cloud.tencent.com/cms/audio/package) 并调整好对应的业务配置。
//
// 
//
// ### 接口使用说明：
//
// - 前往“[内容安全控制台-图片内容安全](https://console.cloud.tencent.com/cms/audio/package)”开启使用音频内容安全服务，首次开通服务的用户可免费领用试用套餐包，包含**10小时**免费调用时长，有效期为1个月。
//
// - 该接口为收费接口，计费方式敬请参见 [腾讯云音频内容安全定价](https://cloud.tencent.com/product/ams/pricing)。
//
// 
//
// ### 接口调用说明：
//
// - 音频文件大小支持：**文件 <= 4M**;
//
// - 音频文件**时长不超过60s**，超过60s音频调用则报错；
//
// - 音频文件支持格式：**wav (PCM编码)** 、**mp3**、**aac**、**m4a** (采样率：16kHz~48kHz，位深：16bit 小端，声道数：单声道/双声道，建议格式：**16kHz/16bit/单声道**)；
//
// - 接口仅限音频文件传入，视频文件传入请调用长音频异步接口；
//
// - 接口**默认QPS为20**，如需自定义配置并发或请求频率，请工单咨询；
//
// - 接口**默认超时为10s**，请求如超过该时长则接口会报错。
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
func (c *Client) CreateAudioModerationSyncTaskWithContext(ctx context.Context, request *CreateAudioModerationSyncTaskRequest) (response *CreateAudioModerationSyncTaskResponse, err error) {
    if request == nil {
        request = NewCreateAudioModerationSyncTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAudioModerationSyncTask require credential")
    }

    request.SetContext(ctx)
    
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
// - 支持识别违规内容，包括：低俗、谩骂、色情、广告等场景；
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
// 
//
// ### 直播断流处理说明：
//
// - 请确认已对接[取消任务](https://cloud.tencent.com/document/product/1219/53258)。
//
// - 如果直播任务取消/结束，则终止直播拉流并退出审核。
//
// - 如果直播任务没有取消/结束，直播视频推流因故中断，产品将在将在10分钟内持续拉流重试。如果10分钟检测到音频切片数据，则恢复正常审核，反之，则终止拉流并退出审核。在拉流终止后，用户如有审核需求，需重新送审。
//
// 
//
// 默认接口请求频率限制：20次/秒。
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
//  UNAUTHORIZEDOPERATION_UNAUTHORIZED = "UnauthorizedOperation.Unauthorized"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateAudioModerationTask(request *CreateAudioModerationTaskRequest) (response *CreateAudioModerationTaskResponse, err error) {
    return c.CreateAudioModerationTaskWithContext(context.Background(), request)
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
// - 支持识别违规内容，包括：低俗、谩骂、色情、广告等场景；
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
// 
//
// ### 直播断流处理说明：
//
// - 请确认已对接[取消任务](https://cloud.tencent.com/document/product/1219/53258)。
//
// - 如果直播任务取消/结束，则终止直播拉流并退出审核。
//
// - 如果直播任务没有取消/结束，直播视频推流因故中断，产品将在将在10分钟内持续拉流重试。如果10分钟检测到音频切片数据，则恢复正常审核，反之，则终止拉流并退出审核。在拉流终止后，用户如有审核需求，需重新送审。
//
// 
//
// 默认接口请求频率限制：20次/秒。
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
//  UNAUTHORIZEDOPERATION_UNAUTHORIZED = "UnauthorizedOperation.Unauthorized"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateAudioModerationTaskWithContext(ctx context.Context, request *CreateAudioModerationTaskRequest) (response *CreateAudioModerationTaskResponse, err error) {
    if request == nil {
        request = NewCreateAudioModerationTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAudioModerationTask require credential")
    }

    request.SetContext(ctx)
    
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
// 通过该接口可查看音频审核任务的详情信息，包括任务状态、检测结果、音频文件识别出的对应文本内容、检测结果所对应的恶意标签及推荐的后续操作等，具体输出内容可查看输出参数示例。<br>默认接口请求频率限制：**100次/秒**。
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
    return c.DescribeTaskDetailWithContext(context.Background(), request)
}

// DescribeTaskDetail
// 通过该接口可查看音频审核任务的详情信息，包括任务状态、检测结果、音频文件识别出的对应文本内容、检测结果所对应的恶意标签及推荐的后续操作等，具体输出内容可查看输出参数示例。<br>默认接口请求频率限制：**100次/秒**。
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
func (c *Client) DescribeTaskDetailWithContext(ctx context.Context, request *DescribeTaskDetailRequest) (response *DescribeTaskDetailResponse, err error) {
    if request == nil {
        request = NewDescribeTaskDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTaskDetail require credential")
    }

    request.SetContext(ctx)
    
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
// 通过该接口可查看审核任务列表；您也可根据多种业务信息（业务类型、审核结果、任务状态等）筛选审核任务列表。任务列表输出内容包括当前查询的任务总量、任务名称、任务状态、音频审核类型、基于检测结果的恶意标签及其后续操作等，具体输出内容可查看输出参数示例。<br>默认接口请求频率限制：**20次/秒**。
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
    return c.DescribeTasksWithContext(context.Background(), request)
}

// DescribeTasks
// 通过该接口可查看审核任务列表；您也可根据多种业务信息（业务类型、审核结果、任务状态等）筛选审核任务列表。任务列表输出内容包括当前查询的任务总量、任务名称、任务状态、音频审核类型、基于检测结果的恶意标签及其后续操作等，具体输出内容可查看输出参数示例。<br>默认接口请求频率限制：**20次/秒**。
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
func (c *Client) DescribeTasksWithContext(ctx context.Context, request *DescribeTasksRequest) (response *DescribeTasksResponse, err error) {
    if request == nil {
        request = NewDescribeTasksRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTasks require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTasksResponse()
    err = c.Send(request, response)
    return
}
