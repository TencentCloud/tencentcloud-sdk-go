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
    "context"
    "errors"
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
// <li>提交长文本语音合成请求后，合成结果在3小时内完成，音频文件在服务端可保存24小时</li>
//
// 
//
// <p></p>
//
// 
//
// 长文本合成支持 SSML，语法详见 [SSML 标记语言](https://cloud.tencent.com/document/product/1073/49575)，使用时需满足如下使用规范：
//
// <li>使用 SSML 标签，需置于 speak 闭合标签内部；</li>
//
// <li>合成文本可包含多组 speak 闭合标签，且无数量限制；</li>
//
// <li>每个 speak 闭合标签内部，字符数不超过 150 字（标签字符本身不计算在内）；</li>
//
// <li>每个 speak 闭合标签内部，使用 break 标签数目最大为 10 个。如需要使用更多，可拆解到多个 speak 标签中；</li>
//
// 
//
// <p></p>
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
//  UNSUPPORTEDOPERATION_PKGEXHAUSTED = "UnsupportedOperation.PkgExhausted"
//  UNSUPPORTEDOPERATION_SERVERALREADYOPEN = "UnsupportedOperation.ServerAlreadyOpen"
//  UNSUPPORTEDOPERATION_SERVERDESTORYED = "UnsupportedOperation.ServerDestoryed"
//  UNSUPPORTEDOPERATION_SERVERNOTOPEN = "UnsupportedOperation.ServerNotOpen"
//  UNSUPPORTEDOPERATION_SERVERSTOPPED = "UnsupportedOperation.ServerStopped"
//  UNSUPPORTEDOPERATION_TEXTTOOLONG = "UnsupportedOperation.TextTooLong"
func (c *Client) CreateTtsTask(request *CreateTtsTaskRequest) (response *CreateTtsTaskResponse, err error) {
    return c.CreateTtsTaskWithContext(context.Background(), request)
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
// <li>提交长文本语音合成请求后，合成结果在3小时内完成，音频文件在服务端可保存24小时</li>
//
// 
//
// <p></p>
//
// 
//
// 长文本合成支持 SSML，语法详见 [SSML 标记语言](https://cloud.tencent.com/document/product/1073/49575)，使用时需满足如下使用规范：
//
// <li>使用 SSML 标签，需置于 speak 闭合标签内部；</li>
//
// <li>合成文本可包含多组 speak 闭合标签，且无数量限制；</li>
//
// <li>每个 speak 闭合标签内部，字符数不超过 150 字（标签字符本身不计算在内）；</li>
//
// <li>每个 speak 闭合标签内部，使用 break 标签数目最大为 10 个。如需要使用更多，可拆解到多个 speak 标签中；</li>
//
// 
//
// <p></p>
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
//  UNSUPPORTEDOPERATION_PKGEXHAUSTED = "UnsupportedOperation.PkgExhausted"
//  UNSUPPORTEDOPERATION_SERVERALREADYOPEN = "UnsupportedOperation.ServerAlreadyOpen"
//  UNSUPPORTEDOPERATION_SERVERDESTORYED = "UnsupportedOperation.ServerDestoryed"
//  UNSUPPORTEDOPERATION_SERVERNOTOPEN = "UnsupportedOperation.ServerNotOpen"
//  UNSUPPORTEDOPERATION_SERVERSTOPPED = "UnsupportedOperation.ServerStopped"
//  UNSUPPORTEDOPERATION_TEXTTOOLONG = "UnsupportedOperation.TextTooLong"
func (c *Client) CreateTtsTaskWithContext(ctx context.Context, request *CreateTtsTaskRequest) (response *CreateTtsTaskResponse, err error) {
    if request == nil {
        request = NewCreateTtsTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateTtsTask require credential")
    }

    request.SetContext(ctx)
    
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
    return c.DescribeTtsTaskStatusWithContext(context.Background(), request)
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
func (c *Client) DescribeTtsTaskStatusWithContext(ctx context.Context, request *DescribeTtsTaskStatusRequest) (response *DescribeTtsTaskStatusResponse, err error) {
    if request == nil {
        request = NewDescribeTtsTaskStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTtsTaskStatus require credential")
    }

    request.SetContext(ctx)
    
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
// 基础合成支持 SSML，语法详见 [SSML 标记语言](https://cloud.tencent.com/document/product/1073/49575)。
//
// 默认接口请求频率限制：20次/秒。
//
// 可能返回的错误码:
//  AUTHFAILURE_INVALIDAUTHORIZATION = "AuthFailure.InvalidAuthorization"
//  INTERNALERROR_ERRORGETROUTE = "InternalError.ErrorGetRoute"
//  INTERNALERROR_EXCEEDMAXLIMIT = "InternalError.ExceedMaxLimit"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDTEXT = "InvalidParameter.InvalidText"
//  INVALIDPARAMETERVALUE_APPID = "InvalidParameterValue.AppId"
//  INVALIDPARAMETERVALUE_APPIDNOTREGISTERED = "InvalidParameterValue.AppIdNotRegistered"
//  INVALIDPARAMETERVALUE_CODEC = "InvalidParameterValue.Codec"
//  INVALIDPARAMETERVALUE_ERRORCARDINALFORMAT = "InvalidParameterValue.ErrorCardinalFormat"
//  INVALIDPARAMETERVALUE_INVALIDTEXT = "InvalidParameterValue.InvalidText"
//  INVALIDPARAMETERVALUE_MISSPARAMETERS = "InvalidParameterValue.MissParameters"
//  INVALIDPARAMETERVALUE_PARTICIPLEERROR = "InvalidParameterValue.ParticipleError"
//  INVALIDPARAMETERVALUE_PRIMARYLANGUAGE = "InvalidParameterValue.PrimaryLanguage"
//  INVALIDPARAMETERVALUE_SSMLINVALID = "InvalidParameterValue.SSMLInvalid"
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
//  UNSUPPORTEDOPERATION_PKGEXHAUSTED = "UnsupportedOperation.PkgExhausted"
//  UNSUPPORTEDOPERATION_SERVERALREADYOPEN = "UnsupportedOperation.ServerAlreadyOpen"
//  UNSUPPORTEDOPERATION_SERVERDESTORYED = "UnsupportedOperation.ServerDestoryed"
//  UNSUPPORTEDOPERATION_SERVERNOTOPEN = "UnsupportedOperation.ServerNotOpen"
//  UNSUPPORTEDOPERATION_SERVERSTOPPED = "UnsupportedOperation.ServerStopped"
//  UNSUPPORTEDOPERATION_TEXTTOOLONG = "UnsupportedOperation.TextTooLong"
func (c *Client) TextToVoice(request *TextToVoiceRequest) (response *TextToVoiceResponse, err error) {
    return c.TextToVoiceWithContext(context.Background(), request)
}

// TextToVoice
// 腾讯云语音合成技术（TTS）可以将任意文本转化为语音，实现让机器和应用张口说话。
//
// 腾讯TTS技术可以应用到很多场景，比如，移动APP语音播报新闻；智能设备语音提醒；依靠网上现有节目或少量录音，快速合成明星语音，降低邀约成本；支持车载导航语音合成的个性化语音播报。
//
// 内测期间免费使用。
//
// 基础合成支持 SSML，语法详见 [SSML 标记语言](https://cloud.tencent.com/document/product/1073/49575)。
//
// 默认接口请求频率限制：20次/秒。
//
// 可能返回的错误码:
//  AUTHFAILURE_INVALIDAUTHORIZATION = "AuthFailure.InvalidAuthorization"
//  INTERNALERROR_ERRORGETROUTE = "InternalError.ErrorGetRoute"
//  INTERNALERROR_EXCEEDMAXLIMIT = "InternalError.ExceedMaxLimit"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDTEXT = "InvalidParameter.InvalidText"
//  INVALIDPARAMETERVALUE_APPID = "InvalidParameterValue.AppId"
//  INVALIDPARAMETERVALUE_APPIDNOTREGISTERED = "InvalidParameterValue.AppIdNotRegistered"
//  INVALIDPARAMETERVALUE_CODEC = "InvalidParameterValue.Codec"
//  INVALIDPARAMETERVALUE_ERRORCARDINALFORMAT = "InvalidParameterValue.ErrorCardinalFormat"
//  INVALIDPARAMETERVALUE_INVALIDTEXT = "InvalidParameterValue.InvalidText"
//  INVALIDPARAMETERVALUE_MISSPARAMETERS = "InvalidParameterValue.MissParameters"
//  INVALIDPARAMETERVALUE_PARTICIPLEERROR = "InvalidParameterValue.ParticipleError"
//  INVALIDPARAMETERVALUE_PRIMARYLANGUAGE = "InvalidParameterValue.PrimaryLanguage"
//  INVALIDPARAMETERVALUE_SSMLINVALID = "InvalidParameterValue.SSMLInvalid"
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
//  UNSUPPORTEDOPERATION_PKGEXHAUSTED = "UnsupportedOperation.PkgExhausted"
//  UNSUPPORTEDOPERATION_SERVERALREADYOPEN = "UnsupportedOperation.ServerAlreadyOpen"
//  UNSUPPORTEDOPERATION_SERVERDESTORYED = "UnsupportedOperation.ServerDestoryed"
//  UNSUPPORTEDOPERATION_SERVERNOTOPEN = "UnsupportedOperation.ServerNotOpen"
//  UNSUPPORTEDOPERATION_SERVERSTOPPED = "UnsupportedOperation.ServerStopped"
//  UNSUPPORTEDOPERATION_TEXTTOOLONG = "UnsupportedOperation.TextTooLong"
func (c *Client) TextToVoiceWithContext(ctx context.Context, request *TextToVoiceRequest) (response *TextToVoiceResponse, err error) {
    if request == nil {
        request = NewTextToVoiceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("TextToVoice require credential")
    }

    request.SetContext(ctx)
    
    response = NewTextToVoiceResponse()
    err = c.Send(request, response)
    return
}
