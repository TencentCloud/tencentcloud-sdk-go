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

package v20190823

import (
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2019-08-23"

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


func NewCreateTtsTaskRequest() (request *CreateTtsTaskRequest) {
    request = &CreateTtsTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tts", APIVersion, "CreateTtsTask")
    
    return
}

func NewCreateTtsTaskResponse() (response *CreateTtsTaskResponse) {
    response = &CreateTtsTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateTtsTask
// 本接口服务对10万字符以内的文本进行语音合成，异步返回音频结果。满足一次性合成较长文本的客户需求，如阅读播报、新闻媒体等场景。
//
// 
//
// <li>支持音频格式：mp3,wav,pcm</li>
//
// <li>支持音频采样率：16000 Hz, 8000 Hz</li>
//
// <li>支持中文普通话、英文、中英文混读、粤语合成</li>
//
// <li>支持语速、音量设置</li>
//
// <li>支持回调或轮询的方式获取结果，结果获取请参考 长文本语音合成结果查询。</li>
//
// <li>长文本语音合成任务完成后，合成音频结果在服务端可保存24小时</li>
//
// 可能返回的错误码:
//  AUTHFAILURE_INVALIDAUTHORIZATION = "AuthFailure.InvalidAuthorization"
//  INTERNALERROR_ERRORGETROUTE = "InternalError.ErrorGetRoute"
//  INTERNALERROR_FAILACCESSDATABASE = "InternalError.FailAccessDatabase"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDTEXT = "InvalidParameter.InvalidText"
//  INVALIDPARAMETER_STATUS = "InvalidParameter.Status"
//  INVALIDPARAMETERVALUE_APPID = "InvalidParameterValue.AppId"
//  INVALIDPARAMETERVALUE_APPIDNOTREGISTERED = "InvalidParameterValue.AppIdNotRegistered"
//  INVALIDPARAMETERVALUE_CALLBACKURL = "InvalidParameterValue.CallbackUrl"
//  INVALIDPARAMETERVALUE_CODEC = "InvalidParameterValue.Codec"
//  INVALIDPARAMETERVALUE_ERRORCARDINALFORMAT = "InvalidParameterValue.ErrorCardinalFormat"
//  INVALIDPARAMETERVALUE_INVALIDTEXT = "InvalidParameterValue.InvalidText"
//  INVALIDPARAMETERVALUE_MISSPARAMETERS = "InvalidParameterValue.MissParameters"
//  INVALIDPARAMETERVALUE_MODELTYPE = "InvalidParameterValue.ModelType"
//  INVALIDPARAMETERVALUE_PRIMARYLANGUAGE = "InvalidParameterValue.PrimaryLanguage"
//  INVALIDPARAMETERVALUE_SAMPLERATE = "InvalidParameterValue.SampleRate"
//  INVALIDPARAMETERVALUE_SPEED = "InvalidParameterValue.Speed"
//  INVALIDPARAMETERVALUE_TEXT = "InvalidParameterValue.Text"
//  INVALIDPARAMETERVALUE_TEXTEMPTY = "InvalidParameterValue.TextEmpty"
//  INVALIDPARAMETERVALUE_TEXTNOTUTF8 = "InvalidParameterValue.TextNotUtf8"
//  INVALIDPARAMETERVALUE_TEXTSSMLPARSEERROR = "InvalidParameterValue.TextSsmlParseError"
//  INVALIDPARAMETERVALUE_TEXTTOOLONG = "InvalidParameterValue.TextTooLong"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  INVALIDPARAMETERVALUE_VOICETYPE = "InvalidParameterValue.VoiceType"
//  INVALIDPARAMETERVALUE_VOLUME = "InvalidParameterValue.Volume"
//  LIMITEXCEEDED_ACCESSLIMIT = "LimitExceeded.AccessLimit"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_ACCOUNTARREARS = "UnsupportedOperation.AccountArrears"
//  UNSUPPORTEDOPERATION_AUTHORIZATIONEXPIRED = "UnsupportedOperation.AuthorizationExpired"
//  UNSUPPORTEDOPERATION_AUTHORIZATIONFAILED = "UnsupportedOperation.AuthorizationFailed"
//  UNSUPPORTEDOPERATION_FORBIDDENUSE = "UnsupportedOperation.ForbiddenUse"
//  UNSUPPORTEDOPERATION_NOBANLANCE = "UnsupportedOperation.NoBanlance"
//  UNSUPPORTEDOPERATION_NOFREEACCOUNT = "UnsupportedOperation.NoFreeAccount"
//  UNSUPPORTEDOPERATION_SERVERALREADYOPEN = "UnsupportedOperation.ServerAlreadyOpen"
//  UNSUPPORTEDOPERATION_SERVERDESTORYED = "UnsupportedOperation.ServerDestoryed"
//  UNSUPPORTEDOPERATION_SERVERNOTOPEN = "UnsupportedOperation.ServerNotOpen"
//  UNSUPPORTEDOPERATION_SERVERSTOPPED = "UnsupportedOperation.ServerStopped"
//  UNSUPPORTEDOPERATION_TEXTTOOLONG = "UnsupportedOperation.TextTooLong"
func (c *Client) CreateTtsTask(request *CreateTtsTaskRequest) (response *CreateTtsTaskResponse, err error) {
    if request == nil {
        request = NewCreateTtsTaskRequest()
    }
    response = NewCreateTtsTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTtsTaskStatusRequest() (request *DescribeTtsTaskStatusRequest) {
    request = &DescribeTtsTaskStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tts", APIVersion, "DescribeTtsTaskStatus")
    
    return
}

func NewDescribeTtsTaskStatusResponse() (response *DescribeTtsTaskStatusResponse) {
    response = &DescribeTtsTaskStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeTtsTaskStatus
// 在调用长文本语音合成请求接口后，有回调和轮询两种方式获取识别结果。
//
// 
//
// <li>当采用回调方式时，合成完毕后会将结果通过 POST 请求的形式通知到用户在请求时填写的回调 URL，具体请参见 长文本语音合成结果查询 。</li>
//
// <li>当采用轮询方式时，需要主动提交任务ID来轮询识别结果，共有任务成功、等待、执行中和失败四种结果，具体信息请参见下文说明。</li>
//
// 可能返回的错误码:
//  AUTHFAILURE_INVALIDAUTHORIZATION = "AuthFailure.InvalidAuthorization"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_NOSUCHTASK = "FailedOperation.NoSuchTask"
//  INTERNALERROR_ERRORGETROUTE = "InternalError.ErrorGetRoute"
//  INTERNALERROR_FAILACCESSDATABASE = "InternalError.FailAccessDatabase"
//  INTERNALERROR_FAILACCESSREDIS = "InternalError.FailAccessRedis"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDTEXT = "InvalidParameter.InvalidText"
//  INVALIDPARAMETER_STATUS = "InvalidParameter.Status"
//  INVALIDPARAMETERVALUE_APPID = "InvalidParameterValue.AppId"
//  INVALIDPARAMETERVALUE_APPIDNOTREGISTERED = "InvalidParameterValue.AppIdNotRegistered"
//  INVALIDPARAMETERVALUE_CODEC = "InvalidParameterValue.Codec"
//  INVALIDPARAMETERVALUE_ERRORCARDINALFORMAT = "InvalidParameterValue.ErrorCardinalFormat"
//  INVALIDPARAMETERVALUE_INVALIDTEXT = "InvalidParameterValue.InvalidText"
//  INVALIDPARAMETERVALUE_MISSPARAMETERS = "InvalidParameterValue.MissParameters"
//  INVALIDPARAMETERVALUE_PRIMARYLANGUAGE = "InvalidParameterValue.PrimaryLanguage"
//  INVALIDPARAMETERVALUE_SAMPLERATE = "InvalidParameterValue.SampleRate"
//  INVALIDPARAMETERVALUE_SESSIONID = "InvalidParameterValue.SessionId"
//  INVALIDPARAMETERVALUE_SPEED = "InvalidParameterValue.Speed"
//  INVALIDPARAMETERVALUE_TEXT = "InvalidParameterValue.Text"
//  INVALIDPARAMETERVALUE_TEXTEMPTY = "InvalidParameterValue.TextEmpty"
//  INVALIDPARAMETERVALUE_TEXTNOTUTF8 = "InvalidParameterValue.TextNotUtf8"
//  INVALIDPARAMETERVALUE_TEXTSSMLPARSEERROR = "InvalidParameterValue.TextSsmlParseError"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  INVALIDPARAMETERVALUE_VOICETYPE = "InvalidParameterValue.VoiceType"
//  INVALIDPARAMETERVALUE_VOLUME = "InvalidParameterValue.Volume"
//  LIMITEXCEEDED_ACCESSLIMIT = "LimitExceeded.AccessLimit"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_ACCOUNTARREARS = "UnsupportedOperation.AccountArrears"
//  UNSUPPORTEDOPERATION_AUTHORIZATIONEXPIRED = "UnsupportedOperation.AuthorizationExpired"
//  UNSUPPORTEDOPERATION_AUTHORIZATIONFAILED = "UnsupportedOperation.AuthorizationFailed"
//  UNSUPPORTEDOPERATION_FORBIDDENUSE = "UnsupportedOperation.ForbiddenUse"
//  UNSUPPORTEDOPERATION_NOBANLANCE = "UnsupportedOperation.NoBanlance"
//  UNSUPPORTEDOPERATION_NOFREEACCOUNT = "UnsupportedOperation.NoFreeAccount"
//  UNSUPPORTEDOPERATION_SERVERALREADYOPEN = "UnsupportedOperation.ServerAlreadyOpen"
//  UNSUPPORTEDOPERATION_SERVERDESTORYED = "UnsupportedOperation.ServerDestoryed"
//  UNSUPPORTEDOPERATION_SERVERNOTOPEN = "UnsupportedOperation.ServerNotOpen"
//  UNSUPPORTEDOPERATION_SERVERSTOPPED = "UnsupportedOperation.ServerStopped"
//  UNSUPPORTEDOPERATION_TEXTTOOLONG = "UnsupportedOperation.TextTooLong"
func (c *Client) DescribeTtsTaskStatus(request *DescribeTtsTaskStatusRequest) (response *DescribeTtsTaskStatusResponse, err error) {
    if request == nil {
        request = NewDescribeTtsTaskStatusRequest()
    }
    response = NewDescribeTtsTaskStatusResponse()
    err = c.Send(request, response)
    return
}

func NewTextToVoiceRequest() (request *TextToVoiceRequest) {
    request = &TextToVoiceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tts", APIVersion, "TextToVoice")
    
    return
}

func NewTextToVoiceResponse() (response *TextToVoiceResponse) {
    response = &TextToVoiceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// TextToVoice
// 腾讯云语音合成技术（TTS）可以将任意文本转化为语音，实现让机器和应用张口说话。
//
// 腾讯TTS技术可以应用到很多场景，比如，移动APP语音播报新闻；智能设备语音提醒；依靠网上现有节目或少量录音，快速合成明星语音，降低邀约成本；支持车载导航语音合成的个性化语音播报。
//
// 内测期间免费使用。
//
// 可能返回的错误码:
//  AUTHFAILURE_INVALIDAUTHORIZATION = "AuthFailure.InvalidAuthorization"
//  INTERNALERROR_ERRORGETROUTE = "InternalError.ErrorGetRoute"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDTEXT = "InvalidParameter.InvalidText"
//  INVALIDPARAMETER_STATUS = "InvalidParameter.Status"
//  INVALIDPARAMETERVALUE_APPID = "InvalidParameterValue.AppId"
//  INVALIDPARAMETERVALUE_APPIDNOTREGISTERED = "InvalidParameterValue.AppIdNotRegistered"
//  INVALIDPARAMETERVALUE_CODEC = "InvalidParameterValue.Codec"
//  INVALIDPARAMETERVALUE_INVALIDTEXT = "InvalidParameterValue.InvalidText"
//  INVALIDPARAMETERVALUE_MISSPARAMETERS = "InvalidParameterValue.MissParameters"
//  INVALIDPARAMETERVALUE_PRIMARYLANGUAGE = "InvalidParameterValue.PrimaryLanguage"
//  INVALIDPARAMETERVALUE_SAMPLERATE = "InvalidParameterValue.SampleRate"
//  INVALIDPARAMETERVALUE_SESSIONID = "InvalidParameterValue.SessionId"
//  INVALIDPARAMETERVALUE_SPEED = "InvalidParameterValue.Speed"
//  INVALIDPARAMETERVALUE_TEXT = "InvalidParameterValue.Text"
//  INVALIDPARAMETERVALUE_TEXTEMPTY = "InvalidParameterValue.TextEmpty"
//  INVALIDPARAMETERVALUE_TEXTNOTUTF8 = "InvalidParameterValue.TextNotUtf8"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  INVALIDPARAMETERVALUE_VOICETYPE = "InvalidParameterValue.VoiceType"
//  INVALIDPARAMETERVALUE_VOLUME = "InvalidParameterValue.Volume"
//  LIMITEXCEEDED_ACCESSLIMIT = "LimitExceeded.AccessLimit"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_ACCOUNTARREARS = "UnsupportedOperation.AccountArrears"
//  UNSUPPORTEDOPERATION_AUTHORIZATIONEXPIRED = "UnsupportedOperation.AuthorizationExpired"
//  UNSUPPORTEDOPERATION_AUTHORIZATIONFAILED = "UnsupportedOperation.AuthorizationFailed"
//  UNSUPPORTEDOPERATION_FORBIDDENUSE = "UnsupportedOperation.ForbiddenUse"
//  UNSUPPORTEDOPERATION_NOFREEACCOUNT = "UnsupportedOperation.NoFreeAccount"
//  UNSUPPORTEDOPERATION_SERVERALREADYOPEN = "UnsupportedOperation.ServerAlreadyOpen"
//  UNSUPPORTEDOPERATION_SERVERNOTOPEN = "UnsupportedOperation.ServerNotOpen"
//  UNSUPPORTEDOPERATION_SERVERSTOPPED = "UnsupportedOperation.ServerStopped"
//  UNSUPPORTEDOPERATION_TEXTTOOLONG = "UnsupportedOperation.TextTooLong"
func (c *Client) TextToVoice(request *TextToVoiceRequest) (response *TextToVoiceResponse, err error) {
    if request == nil {
        request = NewTextToVoiceRequest()
    }
    response = NewTextToVoiceResponse()
    err = c.Send(request, response)
    return
}
