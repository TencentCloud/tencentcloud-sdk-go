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

package v20180522

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2018-05-22"

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


func NewChatRequest() (request *ChatRequest) {
    request = &ChatRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("aai", APIVersion, "Chat")
    
    
    return
}

func NewChatResponse() (response *ChatResponse) {
    response = &ChatResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// Chat
// 提供基于文本的基础聊天能力，可以让您的应用快速拥有具备深度语义理解的机器聊天功能。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ERRORCONFIGURE = "InternalError.ErrorConfigure"
//  INTERNALERROR_ERRORCREATELOG = "InternalError.ErrorCreateLog"
//  INTERNALERROR_ERRORGETROUTE = "InternalError.ErrorGetRoute"
//  INTERNALERROR_ERRORMAKELOGPATH = "InternalError.ErrorMakeLogpath"
//  INVALIDPARAMETER_ERRORCHATTEXT = "InvalidParameter.ErrorChatText"
//  INVALIDPARAMETER_ERRORCHATUSER = "InvalidParameter.ErrorChatUser"
//  INVALIDPARAMETER_ERRORCONTENTLENGTH = "InvalidParameter.ErrorContentlength"
//  INVALIDPARAMETER_ERRORPARAMSMISSING = "InvalidParameter.ErrorParamsMissing"
//  INVALIDPARAMETER_ERRORPARSEQUEST = "InvalidParameter.ErrorParsequest"
//  INVALIDPARAMETERVALUE_ERRORINVALIDAPPID = "InvalidParameterValue.ErrorInvalidAppid"
//  INVALIDPARAMETERVALUE_ERRORINVALIDCLIENTIP = "InvalidParameterValue.ErrorInvalidClientip"
//  INVALIDPARAMETERVALUE_ERRORINVALIDPROJECTID = "InvalidParameterValue.ErrorInvalidProjectid"
//  INVALIDPARAMETERVALUE_ERRORINVALIDREQUESTID = "InvalidParameterValue.ErrorInvalidRequestid"
func (c *Client) Chat(request *ChatRequest) (response *ChatResponse, err error) {
    return c.ChatWithContext(context.Background(), request)
}

// Chat
// 提供基于文本的基础聊天能力，可以让您的应用快速拥有具备深度语义理解的机器聊天功能。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ERRORCONFIGURE = "InternalError.ErrorConfigure"
//  INTERNALERROR_ERRORCREATELOG = "InternalError.ErrorCreateLog"
//  INTERNALERROR_ERRORGETROUTE = "InternalError.ErrorGetRoute"
//  INTERNALERROR_ERRORMAKELOGPATH = "InternalError.ErrorMakeLogpath"
//  INVALIDPARAMETER_ERRORCHATTEXT = "InvalidParameter.ErrorChatText"
//  INVALIDPARAMETER_ERRORCHATUSER = "InvalidParameter.ErrorChatUser"
//  INVALIDPARAMETER_ERRORCONTENTLENGTH = "InvalidParameter.ErrorContentlength"
//  INVALIDPARAMETER_ERRORPARAMSMISSING = "InvalidParameter.ErrorParamsMissing"
//  INVALIDPARAMETER_ERRORPARSEQUEST = "InvalidParameter.ErrorParsequest"
//  INVALIDPARAMETERVALUE_ERRORINVALIDAPPID = "InvalidParameterValue.ErrorInvalidAppid"
//  INVALIDPARAMETERVALUE_ERRORINVALIDCLIENTIP = "InvalidParameterValue.ErrorInvalidClientip"
//  INVALIDPARAMETERVALUE_ERRORINVALIDPROJECTID = "InvalidParameterValue.ErrorInvalidProjectid"
//  INVALIDPARAMETERVALUE_ERRORINVALIDREQUESTID = "InvalidParameterValue.ErrorInvalidRequestid"
func (c *Client) ChatWithContext(ctx context.Context, request *ChatRequest) (response *ChatResponse, err error) {
    if request == nil {
        request = NewChatRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("Chat require credential")
    }

    request.SetContext(ctx)
    
    response = NewChatResponse()
    err = c.Send(request, response)
    return
}

func NewSentenceRecognitionRequest() (request *SentenceRecognitionRequest) {
    request = &SentenceRecognitionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("aai", APIVersion, "SentenceRecognition")
    
    
    return
}

func NewSentenceRecognitionResponse() (response *SentenceRecognitionResponse) {
    response = &SentenceRecognitionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// SentenceRecognition
// 识别60s内的短语音，当音频放在请求body中传输时整个请求大小不能超过600KB，当音频以url方式传输时，音频时长不可超过60s。所有请求参数放在post的body中采用x-www-form-urlencoded（数据转换成一个字符串（name1=value1&name2=value2…）进行urlencode后）编码传输。现暂只支持中文普通话识别，支持识别8k(16k)的16bit的mp3或者wav音频。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SERVICEISOLATE = "FailedOperation.ServiceIsolate"
//  FAILEDOPERATION_USERHASNOFREEAMOUNT = "FailedOperation.UserHasNoFreeAmount"
//  FAILEDOPERATION_USERNOTREGISTERED = "FailedOperation.UserNotRegistered"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ERRORCONFIGURE = "InternalError.ErrorConfigure"
//  INTERNALERROR_ERRORCREATELOG = "InternalError.ErrorCreateLog"
//  INTERNALERROR_ERRORDOWNFILE = "InternalError.ErrorDownFile"
//  INTERNALERROR_ERRORFAILNEWPREQUEST = "InternalError.ErrorFailNewprequest"
//  INTERNALERROR_ERRORFAILWRITETODB = "InternalError.ErrorFailWritetodb"
//  INTERNALERROR_ERRORFILECANNOTOPEN = "InternalError.ErrorFileCannotopen"
//  INTERNALERROR_ERRORGETROUTE = "InternalError.ErrorGetRoute"
//  INTERNALERROR_ERRORMAKELOGPATH = "InternalError.ErrorMakeLogpath"
//  INTERNALERROR_ERRORRECOGNIZE = "InternalError.ErrorRecognize"
//  INVALIDPARAMETER_ERRORCONTENTLENGTH = "InvalidParameter.ErrorContentlength"
//  INVALIDPARAMETER_ERRORPARAMSMISSING = "InvalidParameter.ErrorParamsMissing"
//  INVALIDPARAMETER_ERRORPARSEQUEST = "InvalidParameter.ErrorParsequest"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ERRORINVALIDAPPID = "InvalidParameterValue.ErrorInvalidAppid"
//  INVALIDPARAMETERVALUE_ERRORINVALIDCLIENTIP = "InvalidParameterValue.ErrorInvalidClientip"
//  INVALIDPARAMETERVALUE_ERRORINVALIDENGSERVICE = "InvalidParameterValue.ErrorInvalidEngservice"
//  INVALIDPARAMETERVALUE_ERRORINVALIDPROJECTID = "InvalidParameterValue.ErrorInvalidProjectid"
//  INVALIDPARAMETERVALUE_ERRORINVALIDREQUESTID = "InvalidParameterValue.ErrorInvalidRequestid"
//  INVALIDPARAMETERVALUE_ERRORINVALIDSOURCETYPE = "InvalidParameterValue.ErrorInvalidSourcetype"
//  INVALIDPARAMETERVALUE_ERRORINVALIDSUBSERVICETYPE = "InvalidParameterValue.ErrorInvalidSubservicetype"
//  INVALIDPARAMETERVALUE_ERRORINVALIDURL = "InvalidParameterValue.ErrorInvalidUrl"
//  INVALIDPARAMETERVALUE_ERRORINVALIDUSERAUDIOKEY = "InvalidParameterValue.ErrorInvalidUseraudiokey"
//  INVALIDPARAMETERVALUE_ERRORINVALIDVOICEFORMAT = "InvalidParameterValue.ErrorInvalidVoiceFormat"
//  INVALIDPARAMETERVALUE_ERRORINVALIDVOICEDATA = "InvalidParameterValue.ErrorInvalidVoicedata"
func (c *Client) SentenceRecognition(request *SentenceRecognitionRequest) (response *SentenceRecognitionResponse, err error) {
    return c.SentenceRecognitionWithContext(context.Background(), request)
}

// SentenceRecognition
// 识别60s内的短语音，当音频放在请求body中传输时整个请求大小不能超过600KB，当音频以url方式传输时，音频时长不可超过60s。所有请求参数放在post的body中采用x-www-form-urlencoded（数据转换成一个字符串（name1=value1&name2=value2…）进行urlencode后）编码传输。现暂只支持中文普通话识别，支持识别8k(16k)的16bit的mp3或者wav音频。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SERVICEISOLATE = "FailedOperation.ServiceIsolate"
//  FAILEDOPERATION_USERHASNOFREEAMOUNT = "FailedOperation.UserHasNoFreeAmount"
//  FAILEDOPERATION_USERNOTREGISTERED = "FailedOperation.UserNotRegistered"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ERRORCONFIGURE = "InternalError.ErrorConfigure"
//  INTERNALERROR_ERRORCREATELOG = "InternalError.ErrorCreateLog"
//  INTERNALERROR_ERRORDOWNFILE = "InternalError.ErrorDownFile"
//  INTERNALERROR_ERRORFAILNEWPREQUEST = "InternalError.ErrorFailNewprequest"
//  INTERNALERROR_ERRORFAILWRITETODB = "InternalError.ErrorFailWritetodb"
//  INTERNALERROR_ERRORFILECANNOTOPEN = "InternalError.ErrorFileCannotopen"
//  INTERNALERROR_ERRORGETROUTE = "InternalError.ErrorGetRoute"
//  INTERNALERROR_ERRORMAKELOGPATH = "InternalError.ErrorMakeLogpath"
//  INTERNALERROR_ERRORRECOGNIZE = "InternalError.ErrorRecognize"
//  INVALIDPARAMETER_ERRORCONTENTLENGTH = "InvalidParameter.ErrorContentlength"
//  INVALIDPARAMETER_ERRORPARAMSMISSING = "InvalidParameter.ErrorParamsMissing"
//  INVALIDPARAMETER_ERRORPARSEQUEST = "InvalidParameter.ErrorParsequest"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ERRORINVALIDAPPID = "InvalidParameterValue.ErrorInvalidAppid"
//  INVALIDPARAMETERVALUE_ERRORINVALIDCLIENTIP = "InvalidParameterValue.ErrorInvalidClientip"
//  INVALIDPARAMETERVALUE_ERRORINVALIDENGSERVICE = "InvalidParameterValue.ErrorInvalidEngservice"
//  INVALIDPARAMETERVALUE_ERRORINVALIDPROJECTID = "InvalidParameterValue.ErrorInvalidProjectid"
//  INVALIDPARAMETERVALUE_ERRORINVALIDREQUESTID = "InvalidParameterValue.ErrorInvalidRequestid"
//  INVALIDPARAMETERVALUE_ERRORINVALIDSOURCETYPE = "InvalidParameterValue.ErrorInvalidSourcetype"
//  INVALIDPARAMETERVALUE_ERRORINVALIDSUBSERVICETYPE = "InvalidParameterValue.ErrorInvalidSubservicetype"
//  INVALIDPARAMETERVALUE_ERRORINVALIDURL = "InvalidParameterValue.ErrorInvalidUrl"
//  INVALIDPARAMETERVALUE_ERRORINVALIDUSERAUDIOKEY = "InvalidParameterValue.ErrorInvalidUseraudiokey"
//  INVALIDPARAMETERVALUE_ERRORINVALIDVOICEFORMAT = "InvalidParameterValue.ErrorInvalidVoiceFormat"
//  INVALIDPARAMETERVALUE_ERRORINVALIDVOICEDATA = "InvalidParameterValue.ErrorInvalidVoicedata"
func (c *Client) SentenceRecognitionWithContext(ctx context.Context, request *SentenceRecognitionRequest) (response *SentenceRecognitionResponse, err error) {
    if request == nil {
        request = NewSentenceRecognitionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SentenceRecognition require credential")
    }

    request.SetContext(ctx)
    
    response = NewSentenceRecognitionResponse()
    err = c.Send(request, response)
    return
}

func NewSimultaneousInterpretingRequest() (request *SimultaneousInterpretingRequest) {
    request = &SimultaneousInterpretingRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("aai", APIVersion, "SimultaneousInterpreting")
    
    
    return
}

func NewSimultaneousInterpretingResponse() (response *SimultaneousInterpretingResponse) {
    response = &SimultaneousInterpretingResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// SimultaneousInterpreting
// 该接口是实时流式识别，可同时返回语音识别文本及翻译文本，当前仅支持中文和英文。该接口可配合同传windows客户端，提供会议现场同传服务。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ERRORCONFIGURE = "InternalError.ErrorConfigure"
//  INTERNALERROR_ERRORCREATELOG = "InternalError.ErrorCreateLog"
//  INTERNALERROR_ERRORFAILWRITETODB = "InternalError.ErrorFailWritetodb"
//  INTERNALERROR_ERRORGETROUTE = "InternalError.ErrorGetRoute"
//  INTERNALERROR_ERRORMAKELOGPATH = "InternalError.ErrorMakeLogpath"
//  INTERNALERROR_ERRORRECOGNIZE = "InternalError.ErrorRecognize"
//  INTERNALERROR_ERRORTRANSLATE = "InternalError.ErrorTranslate"
//  INVALIDPARAMETER_ERRORCONTENTLENGTH = "InvalidParameter.ErrorContentlength"
//  INVALIDPARAMETER_ERRORPARAMSMISSING = "InvalidParameter.ErrorParamsMissing"
//  INVALIDPARAMETER_ERRORPARSEQUEST = "InvalidParameter.ErrorParsequest"
//  INVALIDPARAMETERVALUE_ERRORINVALIDAPPID = "InvalidParameterValue.ErrorInvalidAppid"
//  INVALIDPARAMETERVALUE_ERRORINVALIDCLIENTIP = "InvalidParameterValue.ErrorInvalidClientip"
//  INVALIDPARAMETERVALUE_ERRORINVALIDENGSERVICE = "InvalidParameterValue.ErrorInvalidEngservice"
//  INVALIDPARAMETERVALUE_ERRORINVALIDOPENTRANSLATE = "InvalidParameterValue.ErrorInvalidOpenTranslate"
//  INVALIDPARAMETERVALUE_ERRORINVALIDPROJECTID = "InvalidParameterValue.ErrorInvalidProjectid"
//  INVALIDPARAMETERVALUE_ERRORINVALIDREQUESTID = "InvalidParameterValue.ErrorInvalidRequestid"
//  INVALIDPARAMETERVALUE_ERRORINVALIDSOURCELANGUAGE = "InvalidParameterValue.ErrorInvalidSourceLanguage"
//  INVALIDPARAMETERVALUE_ERRORINVALIDSUBSERVICETYPE = "InvalidParameterValue.ErrorInvalidSubservicetype"
//  INVALIDPARAMETERVALUE_ERRORINVALIDTARGETLANGUAGE = "InvalidParameterValue.ErrorInvalidTargetLanguage"
//  INVALIDPARAMETERVALUE_ERRORINVALIDVOICEFORMAT = "InvalidParameterValue.ErrorInvalidVoiceFormat"
func (c *Client) SimultaneousInterpreting(request *SimultaneousInterpretingRequest) (response *SimultaneousInterpretingResponse, err error) {
    return c.SimultaneousInterpretingWithContext(context.Background(), request)
}

// SimultaneousInterpreting
// 该接口是实时流式识别，可同时返回语音识别文本及翻译文本，当前仅支持中文和英文。该接口可配合同传windows客户端，提供会议现场同传服务。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ERRORCONFIGURE = "InternalError.ErrorConfigure"
//  INTERNALERROR_ERRORCREATELOG = "InternalError.ErrorCreateLog"
//  INTERNALERROR_ERRORFAILWRITETODB = "InternalError.ErrorFailWritetodb"
//  INTERNALERROR_ERRORGETROUTE = "InternalError.ErrorGetRoute"
//  INTERNALERROR_ERRORMAKELOGPATH = "InternalError.ErrorMakeLogpath"
//  INTERNALERROR_ERRORRECOGNIZE = "InternalError.ErrorRecognize"
//  INTERNALERROR_ERRORTRANSLATE = "InternalError.ErrorTranslate"
//  INVALIDPARAMETER_ERRORCONTENTLENGTH = "InvalidParameter.ErrorContentlength"
//  INVALIDPARAMETER_ERRORPARAMSMISSING = "InvalidParameter.ErrorParamsMissing"
//  INVALIDPARAMETER_ERRORPARSEQUEST = "InvalidParameter.ErrorParsequest"
//  INVALIDPARAMETERVALUE_ERRORINVALIDAPPID = "InvalidParameterValue.ErrorInvalidAppid"
//  INVALIDPARAMETERVALUE_ERRORINVALIDCLIENTIP = "InvalidParameterValue.ErrorInvalidClientip"
//  INVALIDPARAMETERVALUE_ERRORINVALIDENGSERVICE = "InvalidParameterValue.ErrorInvalidEngservice"
//  INVALIDPARAMETERVALUE_ERRORINVALIDOPENTRANSLATE = "InvalidParameterValue.ErrorInvalidOpenTranslate"
//  INVALIDPARAMETERVALUE_ERRORINVALIDPROJECTID = "InvalidParameterValue.ErrorInvalidProjectid"
//  INVALIDPARAMETERVALUE_ERRORINVALIDREQUESTID = "InvalidParameterValue.ErrorInvalidRequestid"
//  INVALIDPARAMETERVALUE_ERRORINVALIDSOURCELANGUAGE = "InvalidParameterValue.ErrorInvalidSourceLanguage"
//  INVALIDPARAMETERVALUE_ERRORINVALIDSUBSERVICETYPE = "InvalidParameterValue.ErrorInvalidSubservicetype"
//  INVALIDPARAMETERVALUE_ERRORINVALIDTARGETLANGUAGE = "InvalidParameterValue.ErrorInvalidTargetLanguage"
//  INVALIDPARAMETERVALUE_ERRORINVALIDVOICEFORMAT = "InvalidParameterValue.ErrorInvalidVoiceFormat"
func (c *Client) SimultaneousInterpretingWithContext(ctx context.Context, request *SimultaneousInterpretingRequest) (response *SimultaneousInterpretingResponse, err error) {
    if request == nil {
        request = NewSimultaneousInterpretingRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SimultaneousInterpreting require credential")
    }

    request.SetContext(ctx)
    
    response = NewSimultaneousInterpretingResponse()
    err = c.Send(request, response)
    return
}

func NewTextToVoiceRequest() (request *TextToVoiceRequest) {
    request = &TextToVoiceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("aai", APIVersion, "TextToVoice")
    
    
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
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_TEXTCONVERTFAILED = "InternalError.TextConvertFailed"
//  INVALIDPARAMETER_ERRORNOBODYDATA = "InvalidParameter.ErrorNoBodydata"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_APPIDNOTREGISTERED = "InvalidParameterValue.AppIdNotRegistered"
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
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_TEXTCONVERTFAILED = "InternalError.TextConvertFailed"
//  INVALIDPARAMETER_ERRORNOBODYDATA = "InvalidParameter.ErrorNoBodydata"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_APPIDNOTREGISTERED = "InvalidParameterValue.AppIdNotRegistered"
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
