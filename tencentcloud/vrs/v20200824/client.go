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

package v20200824

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2020-08-24"

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


func NewCancelVRSTaskRequest() (request *CancelVRSTaskRequest) {
    request = &CancelVRSTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vrs", APIVersion, "CancelVRSTask")
    
    
    return
}

func NewCancelVRSTaskResponse() (response *CancelVRSTaskResponse) {
    response = &CancelVRSTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CancelVRSTask
// 声音复刻取消任务接口
//
// 可能返回的错误码:
//  FAILEDOPERATION_ERRORINVALIDTASKSTATUS = "FailedOperation.ErrorInvalidTaskStatus"
//  FAILEDOPERATION_NOSUCHTASK = "FailedOperation.NoSuchTask"
//  INTERNALERROR_FAILACCESSDATABASE = "InternalError.FailAccessDatabase"
func (c *Client) CancelVRSTask(request *CancelVRSTaskRequest) (response *CancelVRSTaskResponse, err error) {
    return c.CancelVRSTaskWithContext(context.Background(), request)
}

// CancelVRSTask
// 声音复刻取消任务接口
//
// 可能返回的错误码:
//  FAILEDOPERATION_ERRORINVALIDTASKSTATUS = "FailedOperation.ErrorInvalidTaskStatus"
//  FAILEDOPERATION_NOSUCHTASK = "FailedOperation.NoSuchTask"
//  INTERNALERROR_FAILACCESSDATABASE = "InternalError.FailAccessDatabase"
func (c *Client) CancelVRSTaskWithContext(ctx context.Context, request *CancelVRSTaskRequest) (response *CancelVRSTaskResponse, err error) {
    if request == nil {
        request = NewCancelVRSTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CancelVRSTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCancelVRSTaskResponse()
    err = c.Send(request, response)
    return
}

func NewCreateVRSTaskRequest() (request *CreateVRSTaskRequest) {
    request = &CreateVRSTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vrs", APIVersion, "CreateVRSTask")
    
    
    return
}

func NewCreateVRSTaskResponse() (response *CreateVRSTaskResponse) {
    response = &CreateVRSTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateVRSTask
// 本接口服务对提交音频进行声音复刻任务创建接口，异步返回复刻结果。
//
// • 请求方法为 HTTP POST , Content-Type为"application/json; charset=utf-8"
//
// • 签名方法参考 公共参数 中签名方法v3。
//
// 可能返回的错误码:
//  INTERNALERROR_FAILACCESSDATABASE = "InternalError.FailAccessDatabase"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_CODEC = "InvalidParameterValue.Codec"
//  INVALIDPARAMETERVALUE_SAMPLERATE = "InvalidParameterValue.SampleRate"
//  INVALIDPARAMETERVALUE_VOICEGENDER = "InvalidParameterValue.VoiceGender"
//  INVALIDPARAMETERVALUE_VOICELANGUAGE = "InvalidParameterValue.VoiceLanguage"
//  INVALIDPARAMETERVALUE_VOICENAME = "InvalidParameterValue.VoiceName"
//  UNSUPPORTEDOPERATION_VRSQUOTAEXHAUSTED = "UnsupportedOperation.VRSQuotaExhausted"
func (c *Client) CreateVRSTask(request *CreateVRSTaskRequest) (response *CreateVRSTaskResponse, err error) {
    return c.CreateVRSTaskWithContext(context.Background(), request)
}

// CreateVRSTask
// 本接口服务对提交音频进行声音复刻任务创建接口，异步返回复刻结果。
//
// • 请求方法为 HTTP POST , Content-Type为"application/json; charset=utf-8"
//
// • 签名方法参考 公共参数 中签名方法v3。
//
// 可能返回的错误码:
//  INTERNALERROR_FAILACCESSDATABASE = "InternalError.FailAccessDatabase"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_CODEC = "InvalidParameterValue.Codec"
//  INVALIDPARAMETERVALUE_SAMPLERATE = "InvalidParameterValue.SampleRate"
//  INVALIDPARAMETERVALUE_VOICEGENDER = "InvalidParameterValue.VoiceGender"
//  INVALIDPARAMETERVALUE_VOICELANGUAGE = "InvalidParameterValue.VoiceLanguage"
//  INVALIDPARAMETERVALUE_VOICENAME = "InvalidParameterValue.VoiceName"
//  UNSUPPORTEDOPERATION_VRSQUOTAEXHAUSTED = "UnsupportedOperation.VRSQuotaExhausted"
func (c *Client) CreateVRSTaskWithContext(ctx context.Context, request *CreateVRSTaskRequest) (response *CreateVRSTaskResponse, err error) {
    if request == nil {
        request = NewCreateVRSTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateVRSTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateVRSTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVRSTaskStatusRequest() (request *DescribeVRSTaskStatusRequest) {
    request = &DescribeVRSTaskStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vrs", APIVersion, "DescribeVRSTaskStatus")
    
    
    return
}

func NewDescribeVRSTaskStatusResponse() (response *DescribeVRSTaskStatusResponse) {
    response = &DescribeVRSTaskStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeVRSTaskStatus
// 在调用声音复刻创建任务请求接口后，有回调和轮询两种方式获取识别结果。
//
// • 当采用回调方式时，识别完成后会将结果通过 POST 请求的形式通知到用户在请求时填写的回调 URL，具体请参见 声音复刻结果回调 。
//
// • 当采用轮询方式时，需要主动提交任务ID来轮询识别结果，共有任务成功、等待、执行中和失败四种结果，具体信息请参见下文说明。
//
// • 请求方法为 HTTP POST , Content-Type为"application/json; charset=utf-8"
//
// • 签名方法参考 公共参数 中签名方法v3。
//
// 可能返回的错误码:
//  FAILEDOPERATION_NOSUCHTASK = "FailedOperation.NoSuchTask"
func (c *Client) DescribeVRSTaskStatus(request *DescribeVRSTaskStatusRequest) (response *DescribeVRSTaskStatusResponse, err error) {
    return c.DescribeVRSTaskStatusWithContext(context.Background(), request)
}

// DescribeVRSTaskStatus
// 在调用声音复刻创建任务请求接口后，有回调和轮询两种方式获取识别结果。
//
// • 当采用回调方式时，识别完成后会将结果通过 POST 请求的形式通知到用户在请求时填写的回调 URL，具体请参见 声音复刻结果回调 。
//
// • 当采用轮询方式时，需要主动提交任务ID来轮询识别结果，共有任务成功、等待、执行中和失败四种结果，具体信息请参见下文说明。
//
// • 请求方法为 HTTP POST , Content-Type为"application/json; charset=utf-8"
//
// • 签名方法参考 公共参数 中签名方法v3。
//
// 可能返回的错误码:
//  FAILEDOPERATION_NOSUCHTASK = "FailedOperation.NoSuchTask"
func (c *Client) DescribeVRSTaskStatusWithContext(ctx context.Context, request *DescribeVRSTaskStatusRequest) (response *DescribeVRSTaskStatusResponse, err error) {
    if request == nil {
        request = NewDescribeVRSTaskStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVRSTaskStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVRSTaskStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDetectEnvAndSoundQualityRequest() (request *DetectEnvAndSoundQualityRequest) {
    request = &DetectEnvAndSoundQualityRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vrs", APIVersion, "DetectEnvAndSoundQuality")
    
    
    return
}

func NewDetectEnvAndSoundQualityResponse() (response *DetectEnvAndSoundQualityResponse) {
    response = &DetectEnvAndSoundQualityResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DetectEnvAndSoundQuality
// 本接口用于检测音频的环境和音频质量。
//
// 对于一句话声音复刻，音频时长需大于3s，小于15s，文件大小不能超过2MB，音频需为单声道，位深为16bit。建议格式：wav、单声道、采样率48kHz或24kHz 
//
// • 请求方法为 HTTP POST , Content-Type为"application/json; charset=utf-8"
//
// • 签名方法参考 公共参数 中签名方法v3。
//
// 可能返回的错误码:
//  FAILEDOPERATION_VOICEEVALUATEFAILED = "FailedOperation.VoiceEvaluateFailed"
//  FAILEDOPERATION_VOICENOTQUALIFIED = "FailedOperation.VoiceNotQualified"
//  INVALIDPARAMETERVALUE_AUDIODATA = "InvalidParameterValue.AudioData"
//  INVALIDPARAMETERVALUE_AUDIODURATIONEXCEEDSLIMIT = "InvalidParameterValue.AudioDurationExceedsLimit"
//  INVALIDPARAMETERVALUE_CODEC = "InvalidParameterValue.Codec"
//  INVALIDPARAMETERVALUE_SAMPLERATE = "InvalidParameterValue.SampleRate"
func (c *Client) DetectEnvAndSoundQuality(request *DetectEnvAndSoundQualityRequest) (response *DetectEnvAndSoundQualityResponse, err error) {
    return c.DetectEnvAndSoundQualityWithContext(context.Background(), request)
}

// DetectEnvAndSoundQuality
// 本接口用于检测音频的环境和音频质量。
//
// 对于一句话声音复刻，音频时长需大于3s，小于15s，文件大小不能超过2MB，音频需为单声道，位深为16bit。建议格式：wav、单声道、采样率48kHz或24kHz 
//
// • 请求方法为 HTTP POST , Content-Type为"application/json; charset=utf-8"
//
// • 签名方法参考 公共参数 中签名方法v3。
//
// 可能返回的错误码:
//  FAILEDOPERATION_VOICEEVALUATEFAILED = "FailedOperation.VoiceEvaluateFailed"
//  FAILEDOPERATION_VOICENOTQUALIFIED = "FailedOperation.VoiceNotQualified"
//  INVALIDPARAMETERVALUE_AUDIODATA = "InvalidParameterValue.AudioData"
//  INVALIDPARAMETERVALUE_AUDIODURATIONEXCEEDSLIMIT = "InvalidParameterValue.AudioDurationExceedsLimit"
//  INVALIDPARAMETERVALUE_CODEC = "InvalidParameterValue.Codec"
//  INVALIDPARAMETERVALUE_SAMPLERATE = "InvalidParameterValue.SampleRate"
func (c *Client) DetectEnvAndSoundQualityWithContext(ctx context.Context, request *DetectEnvAndSoundQualityRequest) (response *DetectEnvAndSoundQualityResponse, err error) {
    if request == nil {
        request = NewDetectEnvAndSoundQualityRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DetectEnvAndSoundQuality require credential")
    }

    request.SetContext(ctx)
    
    response = NewDetectEnvAndSoundQualityResponse()
    err = c.Send(request, response)
    return
}

func NewDownloadVRSModelRequest() (request *DownloadVRSModelRequest) {
    request = &DownloadVRSModelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vrs", APIVersion, "DownloadVRSModel")
    
    
    return
}

func NewDownloadVRSModelResponse() (response *DownloadVRSModelResponse) {
    response = &DownloadVRSModelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DownloadVRSModel
// 下载声音复刻离线模型
//
// 可能返回的错误码:
//  FAILEDOPERATION_ERRORINVALIDTASKSTATUS = "FailedOperation.ErrorInvalidTaskStatus"
//  FAILEDOPERATION_NOSUCHTASK = "FailedOperation.NoSuchTask"
//  INTERNALERROR_FAILACCESSDATABASE = "InternalError.FailAccessDatabase"
//  INVALIDPARAMETERVALUE_ERRORINVALIDTASKID = "InvalidParameterValue.ErrorInvalidTaskId"
func (c *Client) DownloadVRSModel(request *DownloadVRSModelRequest) (response *DownloadVRSModelResponse, err error) {
    return c.DownloadVRSModelWithContext(context.Background(), request)
}

// DownloadVRSModel
// 下载声音复刻离线模型
//
// 可能返回的错误码:
//  FAILEDOPERATION_ERRORINVALIDTASKSTATUS = "FailedOperation.ErrorInvalidTaskStatus"
//  FAILEDOPERATION_NOSUCHTASK = "FailedOperation.NoSuchTask"
//  INTERNALERROR_FAILACCESSDATABASE = "InternalError.FailAccessDatabase"
//  INVALIDPARAMETERVALUE_ERRORINVALIDTASKID = "InvalidParameterValue.ErrorInvalidTaskId"
func (c *Client) DownloadVRSModelWithContext(ctx context.Context, request *DownloadVRSModelRequest) (response *DownloadVRSModelResponse, err error) {
    if request == nil {
        request = NewDownloadVRSModelRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DownloadVRSModel require credential")
    }

    request.SetContext(ctx)
    
    response = NewDownloadVRSModelResponse()
    err = c.Send(request, response)
    return
}

func NewGetTrainingTextRequest() (request *GetTrainingTextRequest) {
    request = &GetTrainingTextRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vrs", APIVersion, "GetTrainingText")
    
    
    return
}

func NewGetTrainingTextResponse() (response *GetTrainingTextResponse) {
    response = &GetTrainingTextResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetTrainingText
// 本接口用于获取声音复刻训练文本信息。
//
//  请求方法为 HTTP POST , Content-Type为"application/json; charset=utf-8"
//
// • 签名方法参考 公共参数 中签名方法v3。
//
// • 当复刻类型为一句话声音复刻时，生成的TextId有效期为7天，且在成功创建一次复刻任务后失效。
//
// 可能返回的错误码:
//  FAILEDOPERATION_ERRORINVALIDTASKSTATUS = "FailedOperation.ErrorInvalidTaskStatus"
//  FAILEDOPERATION_NOSUCHTASK = "FailedOperation.NoSuchTask"
//  INTERNALERROR_FAILACCESSDATABASE = "InternalError.FailAccessDatabase"
//  INVALIDPARAMETERVALUE_ERRORINVALIDTASKID = "InvalidParameterValue.ErrorInvalidTaskId"
func (c *Client) GetTrainingText(request *GetTrainingTextRequest) (response *GetTrainingTextResponse, err error) {
    return c.GetTrainingTextWithContext(context.Background(), request)
}

// GetTrainingText
// 本接口用于获取声音复刻训练文本信息。
//
//  请求方法为 HTTP POST , Content-Type为"application/json; charset=utf-8"
//
// • 签名方法参考 公共参数 中签名方法v3。
//
// • 当复刻类型为一句话声音复刻时，生成的TextId有效期为7天，且在成功创建一次复刻任务后失效。
//
// 可能返回的错误码:
//  FAILEDOPERATION_ERRORINVALIDTASKSTATUS = "FailedOperation.ErrorInvalidTaskStatus"
//  FAILEDOPERATION_NOSUCHTASK = "FailedOperation.NoSuchTask"
//  INTERNALERROR_FAILACCESSDATABASE = "InternalError.FailAccessDatabase"
//  INVALIDPARAMETERVALUE_ERRORINVALIDTASKID = "InvalidParameterValue.ErrorInvalidTaskId"
func (c *Client) GetTrainingTextWithContext(ctx context.Context, request *GetTrainingTextRequest) (response *GetTrainingTextResponse, err error) {
    if request == nil {
        request = NewGetTrainingTextRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetTrainingText require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetTrainingTextResponse()
    err = c.Send(request, response)
    return
}

func NewGetVRSVoiceTypeInfoRequest() (request *GetVRSVoiceTypeInfoRequest) {
    request = &GetVRSVoiceTypeInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vrs", APIVersion, "GetVRSVoiceTypeInfo")
    
    
    return
}

func NewGetVRSVoiceTypeInfoResponse() (response *GetVRSVoiceTypeInfoResponse) {
    response = &GetVRSVoiceTypeInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetVRSVoiceTypeInfo
// 该接口用于查询复刻音色详细信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION_ERRORINVALIDTASKSTATUS = "FailedOperation.ErrorInvalidTaskStatus"
//  FAILEDOPERATION_NOSUCHTASK = "FailedOperation.NoSuchTask"
//  INTERNALERROR_FAILACCESSDATABASE = "InternalError.FailAccessDatabase"
//  INVALIDPARAMETERVALUE_ERRORINVALIDTASKID = "InvalidParameterValue.ErrorInvalidTaskId"
func (c *Client) GetVRSVoiceTypeInfo(request *GetVRSVoiceTypeInfoRequest) (response *GetVRSVoiceTypeInfoResponse, err error) {
    return c.GetVRSVoiceTypeInfoWithContext(context.Background(), request)
}

// GetVRSVoiceTypeInfo
// 该接口用于查询复刻音色详细信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION_ERRORINVALIDTASKSTATUS = "FailedOperation.ErrorInvalidTaskStatus"
//  FAILEDOPERATION_NOSUCHTASK = "FailedOperation.NoSuchTask"
//  INTERNALERROR_FAILACCESSDATABASE = "InternalError.FailAccessDatabase"
//  INVALIDPARAMETERVALUE_ERRORINVALIDTASKID = "InvalidParameterValue.ErrorInvalidTaskId"
func (c *Client) GetVRSVoiceTypeInfoWithContext(ctx context.Context, request *GetVRSVoiceTypeInfoRequest) (response *GetVRSVoiceTypeInfoResponse, err error) {
    if request == nil {
        request = NewGetVRSVoiceTypeInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetVRSVoiceTypeInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetVRSVoiceTypeInfoResponse()
    err = c.Send(request, response)
    return
}

func NewGetVRSVoiceTypesRequest() (request *GetVRSVoiceTypesRequest) {
    request = &GetVRSVoiceTypesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vrs", APIVersion, "GetVRSVoiceTypes")
    
    
    return
}

func NewGetVRSVoiceTypesResponse() (response *GetVRSVoiceTypesResponse) {
    response = &GetVRSVoiceTypesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetVRSVoiceTypes
// 查询复刻音色
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  INTERNALERROR_FAILACCESSDATABASE = "InternalError.FailAccessDatabase"
func (c *Client) GetVRSVoiceTypes(request *GetVRSVoiceTypesRequest) (response *GetVRSVoiceTypesResponse, err error) {
    return c.GetVRSVoiceTypesWithContext(context.Background(), request)
}

// GetVRSVoiceTypes
// 查询复刻音色
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  INTERNALERROR_FAILACCESSDATABASE = "InternalError.FailAccessDatabase"
func (c *Client) GetVRSVoiceTypesWithContext(ctx context.Context, request *GetVRSVoiceTypesRequest) (response *GetVRSVoiceTypesResponse, err error) {
    if request == nil {
        request = NewGetVRSVoiceTypesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetVRSVoiceTypes require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetVRSVoiceTypesResponse()
    err = c.Send(request, response)
    return
}
