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

package v20180321

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2018-03-21"

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


func NewFileTranslateRequest() (request *FileTranslateRequest) {
    request = &FileTranslateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tmt", APIVersion, "FileTranslate")
    
    
    return
}

func NewFileTranslateResponse() (response *FileTranslateResponse) {
    response = &FileTranslateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// FileTranslate
// 提交文档原文内容，输出任务ID， 支持原文为单一语种文档（如出现多语言文档，仅支持以选定的源语言相关内容翻译）,文件格式有pdf、docx、pptx、xlsx，支持的文本格式有txt、xml、html、markdown、properties。任务翻译数据可保存7天，7天后不再返回任务数据。请注意保存。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INSERTERR = "FailedOperation.InsertErr"
//  FAILEDOPERATION_REQUESTAILABERR = "FailedOperation.RequestAiLabErr"
//  FAILEDOPERATION_USERNOTREGISTERED = "FailedOperation.UserNotRegistered"
func (c *Client) FileTranslate(request *FileTranslateRequest) (response *FileTranslateResponse, err error) {
    return c.FileTranslateWithContext(context.Background(), request)
}

// FileTranslate
// 提交文档原文内容，输出任务ID， 支持原文为单一语种文档（如出现多语言文档，仅支持以选定的源语言相关内容翻译）,文件格式有pdf、docx、pptx、xlsx，支持的文本格式有txt、xml、html、markdown、properties。任务翻译数据可保存7天，7天后不再返回任务数据。请注意保存。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INSERTERR = "FailedOperation.InsertErr"
//  FAILEDOPERATION_REQUESTAILABERR = "FailedOperation.RequestAiLabErr"
//  FAILEDOPERATION_USERNOTREGISTERED = "FailedOperation.UserNotRegistered"
func (c *Client) FileTranslateWithContext(ctx context.Context, request *FileTranslateRequest) (response *FileTranslateResponse, err error) {
    if request == nil {
        request = NewFileTranslateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("FileTranslate require credential")
    }

    request.SetContext(ctx)
    
    response = NewFileTranslateResponse()
    err = c.Send(request, response)
    return
}

func NewGetFileTranslateRequest() (request *GetFileTranslateRequest) {
    request = &GetFileTranslateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tmt", APIVersion, "GetFileTranslate")
    
    
    return
}

func NewGetFileTranslateResponse() (response *GetFileTranslateResponse) {
    response = &GetFileTranslateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetFileTranslate
// 在调用文档翻译请求接口后，有回调和轮询两种方式获取识别结果。
//
// •当采用回调方式时，翻译完成后会将结果通过 POST 请求的形式通知到用户在请求时填写的回调 URL，具体请参见文档翻译结果回调 。
//
// • 当采用轮询方式时，需要主动提交任务ID来轮询识别结果，共有任务成功、等待、执行中和失败四种结果，具体信息请参见参数说明。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) GetFileTranslate(request *GetFileTranslateRequest) (response *GetFileTranslateResponse, err error) {
    return c.GetFileTranslateWithContext(context.Background(), request)
}

// GetFileTranslate
// 在调用文档翻译请求接口后，有回调和轮询两种方式获取识别结果。
//
// •当采用回调方式时，翻译完成后会将结果通过 POST 请求的形式通知到用户在请求时填写的回调 URL，具体请参见文档翻译结果回调 。
//
// • 当采用轮询方式时，需要主动提交任务ID来轮询识别结果，共有任务成功、等待、执行中和失败四种结果，具体信息请参见参数说明。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) GetFileTranslateWithContext(ctx context.Context, request *GetFileTranslateRequest) (response *GetFileTranslateResponse, err error) {
    if request == nil {
        request = NewGetFileTranslateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetFileTranslate require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetFileTranslateResponse()
    err = c.Send(request, response)
    return
}

func NewImageTranslateRequest() (request *ImageTranslateRequest) {
    request = &ImageTranslateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tmt", APIVersion, "ImageTranslate")
    
    
    return
}

func NewImageTranslateResponse() (response *ImageTranslateResponse) {
    response = &ImageTranslateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ImageTranslate
// 提供13种语言的图片翻译服务，可自动识别图片中的文本内容并翻译成目标语言，识别后的文本按行翻译，后续会提供可按段落翻译的版本。<br />
//
// 提示：对于一般开发者，我们建议优先使用SDK接入简化开发。SDK使用介绍请直接查看 5. 开发者资源 部分。
//
// 可能返回的错误码:
//  FAILEDOPERATION_NOFREEAMOUNT = "FailedOperation.NoFreeAmount"
//  FAILEDOPERATION_SERVICEISOLATE = "FailedOperation.ServiceIsolate"
//  FAILEDOPERATION_USERNOTREGISTERED = "FailedOperation.UserNotRegistered"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_BACKENDTIMEOUT = "InternalError.BackendTimeout"
//  INTERNALERROR_ERRORUNKNOWN = "InternalError.ErrorUnknown"
//  INTERNALERROR_REQUESTFAILED = "InternalError.RequestFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION_ACTIONNOTFOUND = "UnauthorizedOperation.ActionNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_TEXTTOOLONG = "UnsupportedOperation.TextTooLong"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDTARGETLANGUAGE = "UnsupportedOperation.UnSupportedTargetLanguage"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDLANGUAGE = "UnsupportedOperation.UnsupportedLanguage"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDSOURCELANGUAGE = "UnsupportedOperation.UnsupportedSourceLanguage"
func (c *Client) ImageTranslate(request *ImageTranslateRequest) (response *ImageTranslateResponse, err error) {
    return c.ImageTranslateWithContext(context.Background(), request)
}

// ImageTranslate
// 提供13种语言的图片翻译服务，可自动识别图片中的文本内容并翻译成目标语言，识别后的文本按行翻译，后续会提供可按段落翻译的版本。<br />
//
// 提示：对于一般开发者，我们建议优先使用SDK接入简化开发。SDK使用介绍请直接查看 5. 开发者资源 部分。
//
// 可能返回的错误码:
//  FAILEDOPERATION_NOFREEAMOUNT = "FailedOperation.NoFreeAmount"
//  FAILEDOPERATION_SERVICEISOLATE = "FailedOperation.ServiceIsolate"
//  FAILEDOPERATION_USERNOTREGISTERED = "FailedOperation.UserNotRegistered"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_BACKENDTIMEOUT = "InternalError.BackendTimeout"
//  INTERNALERROR_ERRORUNKNOWN = "InternalError.ErrorUnknown"
//  INTERNALERROR_REQUESTFAILED = "InternalError.RequestFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION_ACTIONNOTFOUND = "UnauthorizedOperation.ActionNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_TEXTTOOLONG = "UnsupportedOperation.TextTooLong"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDTARGETLANGUAGE = "UnsupportedOperation.UnSupportedTargetLanguage"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDLANGUAGE = "UnsupportedOperation.UnsupportedLanguage"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDSOURCELANGUAGE = "UnsupportedOperation.UnsupportedSourceLanguage"
func (c *Client) ImageTranslateWithContext(ctx context.Context, request *ImageTranslateRequest) (response *ImageTranslateResponse, err error) {
    if request == nil {
        request = NewImageTranslateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ImageTranslate require credential")
    }

    request.SetContext(ctx)
    
    response = NewImageTranslateResponse()
    err = c.Send(request, response)
    return
}

func NewLanguageDetectRequest() (request *LanguageDetectRequest) {
    request = &LanguageDetectRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tmt", APIVersion, "LanguageDetect")
    
    
    return
}

func NewLanguageDetectResponse() (response *LanguageDetectResponse) {
    response = &LanguageDetectResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// LanguageDetect
// 可自动识别文本内容的语言种类，轻量高效，无需额外实现判断方式，使面向客户的服务体验更佳。 <br />
//
// 提示：对于一般开发者，我们建议优先使用SDK接入简化开发。SDK使用介绍请直接查看 5. 开发者资源 部分。
//
// 可能返回的错误码:
//  FAILEDOPERATION_NOFREEAMOUNT = "FailedOperation.NoFreeAmount"
//  FAILEDOPERATION_SERVICEISOLATE = "FailedOperation.ServiceIsolate"
//  FAILEDOPERATION_USERNOTREGISTERED = "FailedOperation.UserNotRegistered"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_BACKENDTIMEOUT = "InternalError.BackendTimeout"
//  INTERNALERROR_ERRORUNKNOWN = "InternalError.ErrorUnknown"
//  INTERNALERROR_REQUESTFAILED = "InternalError.RequestFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_LIMITEDACCESSFREQUENCY = "LimitExceeded.LimitedAccessFrequency"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION_ACTIONNOTFOUND = "UnauthorizedOperation.ActionNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_TEXTTOOLONG = "UnsupportedOperation.TextTooLong"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDTARGETLANGUAGE = "UnsupportedOperation.UnSupportedTargetLanguage"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDLANGUAGE = "UnsupportedOperation.UnsupportedLanguage"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDSOURCELANGUAGE = "UnsupportedOperation.UnsupportedSourceLanguage"
func (c *Client) LanguageDetect(request *LanguageDetectRequest) (response *LanguageDetectResponse, err error) {
    return c.LanguageDetectWithContext(context.Background(), request)
}

// LanguageDetect
// 可自动识别文本内容的语言种类，轻量高效，无需额外实现判断方式，使面向客户的服务体验更佳。 <br />
//
// 提示：对于一般开发者，我们建议优先使用SDK接入简化开发。SDK使用介绍请直接查看 5. 开发者资源 部分。
//
// 可能返回的错误码:
//  FAILEDOPERATION_NOFREEAMOUNT = "FailedOperation.NoFreeAmount"
//  FAILEDOPERATION_SERVICEISOLATE = "FailedOperation.ServiceIsolate"
//  FAILEDOPERATION_USERNOTREGISTERED = "FailedOperation.UserNotRegistered"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_BACKENDTIMEOUT = "InternalError.BackendTimeout"
//  INTERNALERROR_ERRORUNKNOWN = "InternalError.ErrorUnknown"
//  INTERNALERROR_REQUESTFAILED = "InternalError.RequestFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_LIMITEDACCESSFREQUENCY = "LimitExceeded.LimitedAccessFrequency"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION_ACTIONNOTFOUND = "UnauthorizedOperation.ActionNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_TEXTTOOLONG = "UnsupportedOperation.TextTooLong"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDTARGETLANGUAGE = "UnsupportedOperation.UnSupportedTargetLanguage"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDLANGUAGE = "UnsupportedOperation.UnsupportedLanguage"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDSOURCELANGUAGE = "UnsupportedOperation.UnsupportedSourceLanguage"
func (c *Client) LanguageDetectWithContext(ctx context.Context, request *LanguageDetectRequest) (response *LanguageDetectResponse, err error) {
    if request == nil {
        request = NewLanguageDetectRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("LanguageDetect require credential")
    }

    request.SetContext(ctx)
    
    response = NewLanguageDetectResponse()
    err = c.Send(request, response)
    return
}

func NewSpeechTranslateRequest() (request *SpeechTranslateRequest) {
    request = &SpeechTranslateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tmt", APIVersion, "SpeechTranslate")
    
    
    return
}

func NewSpeechTranslateResponse() (response *SpeechTranslateResponse) {
    response = &SpeechTranslateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// SpeechTranslate
// 本接口提供上传音频，将音频进行语音识别并翻译成文本的服务，目前开放中英互译的语音翻译服务。
//
// 待识别和翻译的音频文件可以是 pcm、mp3和speex 格式，pcm采样率要求16kHz、位深16bit、单声道，音频内语音清晰。<br/>
//
// 如果采用流式传输的方式，要求每个分片时长200ms~500ms；如果采用非流式的传输方式，要求音频时长不超过8s。注意最后一个分片的IsEnd参数设置为1。<br />
//
// 提示：对于一般开发者，我们建议优先使用SDK接入简化开发。SDK使用介绍请直接查看 5. 开发者资源部分。
//
// 可能返回的错误码:
//  FAILEDOPERATION_NOFREEAMOUNT = "FailedOperation.NoFreeAmount"
//  FAILEDOPERATION_SERVICEISOLATE = "FailedOperation.ServiceIsolate"
//  FAILEDOPERATION_USERNOTREGISTERED = "FailedOperation.UserNotRegistered"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_BACKENDTIMEOUT = "InternalError.BackendTimeout"
//  INTERNALERROR_ERRORUNKNOWN = "InternalError.ErrorUnknown"
//  INTERNALERROR_REQUESTFAILED = "InternalError.RequestFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DUPLICATEDSESSIONIDANDSEQ = "InvalidParameter.DuplicatedSessionIdAndSeq"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_SEQINTERVALTOOLARGE = "InvalidParameter.SeqIntervalTooLarge"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_LIMITEDACCESSFREQUENCY = "LimitExceeded.LimitedAccessFrequency"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION_ACTIONNOTFOUND = "UnauthorizedOperation.ActionNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_AUDIODURATIONEXCEED = "UnsupportedOperation.AudioDurationExceed"
//  UNSUPPORTEDOPERATION_TEXTTOOLONG = "UnsupportedOperation.TextTooLong"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDTARGETLANGUAGE = "UnsupportedOperation.UnSupportedTargetLanguage"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDLANGUAGE = "UnsupportedOperation.UnsupportedLanguage"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDSOURCELANGUAGE = "UnsupportedOperation.UnsupportedSourceLanguage"
func (c *Client) SpeechTranslate(request *SpeechTranslateRequest) (response *SpeechTranslateResponse, err error) {
    return c.SpeechTranslateWithContext(context.Background(), request)
}

// SpeechTranslate
// 本接口提供上传音频，将音频进行语音识别并翻译成文本的服务，目前开放中英互译的语音翻译服务。
//
// 待识别和翻译的音频文件可以是 pcm、mp3和speex 格式，pcm采样率要求16kHz、位深16bit、单声道，音频内语音清晰。<br/>
//
// 如果采用流式传输的方式，要求每个分片时长200ms~500ms；如果采用非流式的传输方式，要求音频时长不超过8s。注意最后一个分片的IsEnd参数设置为1。<br />
//
// 提示：对于一般开发者，我们建议优先使用SDK接入简化开发。SDK使用介绍请直接查看 5. 开发者资源部分。
//
// 可能返回的错误码:
//  FAILEDOPERATION_NOFREEAMOUNT = "FailedOperation.NoFreeAmount"
//  FAILEDOPERATION_SERVICEISOLATE = "FailedOperation.ServiceIsolate"
//  FAILEDOPERATION_USERNOTREGISTERED = "FailedOperation.UserNotRegistered"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_BACKENDTIMEOUT = "InternalError.BackendTimeout"
//  INTERNALERROR_ERRORUNKNOWN = "InternalError.ErrorUnknown"
//  INTERNALERROR_REQUESTFAILED = "InternalError.RequestFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DUPLICATEDSESSIONIDANDSEQ = "InvalidParameter.DuplicatedSessionIdAndSeq"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_SEQINTERVALTOOLARGE = "InvalidParameter.SeqIntervalTooLarge"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_LIMITEDACCESSFREQUENCY = "LimitExceeded.LimitedAccessFrequency"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION_ACTIONNOTFOUND = "UnauthorizedOperation.ActionNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_AUDIODURATIONEXCEED = "UnsupportedOperation.AudioDurationExceed"
//  UNSUPPORTEDOPERATION_TEXTTOOLONG = "UnsupportedOperation.TextTooLong"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDTARGETLANGUAGE = "UnsupportedOperation.UnSupportedTargetLanguage"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDLANGUAGE = "UnsupportedOperation.UnsupportedLanguage"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDSOURCELANGUAGE = "UnsupportedOperation.UnsupportedSourceLanguage"
func (c *Client) SpeechTranslateWithContext(ctx context.Context, request *SpeechTranslateRequest) (response *SpeechTranslateResponse, err error) {
    if request == nil {
        request = NewSpeechTranslateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SpeechTranslate require credential")
    }

    request.SetContext(ctx)
    
    response = NewSpeechTranslateResponse()
    err = c.Send(request, response)
    return
}

func NewTextTranslateRequest() (request *TextTranslateRequest) {
    request = &TextTranslateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tmt", APIVersion, "TextTranslate")
    
    
    return
}

func NewTextTranslateResponse() (response *TextTranslateResponse) {
    response = &TextTranslateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// TextTranslate
// 提供中文到英文、英文到中文的等多种语言的文本内容翻译服务， 经过大数据语料库、多种解码算法、翻译引擎深度优化，在新闻文章、生活口语等不同语言场景中都有深厚积累，翻译结果专业评价处于行业领先水平。<br />
//
// 提示：对于一般开发者，我们建议优先使用SDK接入简化开发。SDK使用介绍请直接查看 5. 开发者资源 部分。
//
// 可能返回的错误码:
//  FAILEDOPERATION_NOFREEAMOUNT = "FailedOperation.NoFreeAmount"
//  FAILEDOPERATION_SERVICEISOLATE = "FailedOperation.ServiceIsolate"
//  FAILEDOPERATION_USERNOTREGISTERED = "FailedOperation.UserNotRegistered"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_BACKENDTIMEOUT = "InternalError.BackendTimeout"
//  INTERNALERROR_ERRORGETROUTE = "InternalError.ErrorGetRoute"
//  INTERNALERROR_ERRORUNKNOWN = "InternalError.ErrorUnknown"
//  INTERNALERROR_REQUESTFAILED = "InternalError.RequestFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_LIMITEDACCESSFREQUENCY = "LimitExceeded.LimitedAccessFrequency"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED_UINLIMITEXCEEDED = "RequestLimitExceeded.UinLimitExceeded"
//  UNAUTHORIZEDOPERATION_ACTIONNOTFOUND = "UnauthorizedOperation.ActionNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_TEXTTOOLONG = "UnsupportedOperation.TextTooLong"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDTARGETLANGUAGE = "UnsupportedOperation.UnSupportedTargetLanguage"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDLANGUAGE = "UnsupportedOperation.UnsupportedLanguage"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDSOURCELANGUAGE = "UnsupportedOperation.UnsupportedSourceLanguage"
func (c *Client) TextTranslate(request *TextTranslateRequest) (response *TextTranslateResponse, err error) {
    return c.TextTranslateWithContext(context.Background(), request)
}

// TextTranslate
// 提供中文到英文、英文到中文的等多种语言的文本内容翻译服务， 经过大数据语料库、多种解码算法、翻译引擎深度优化，在新闻文章、生活口语等不同语言场景中都有深厚积累，翻译结果专业评价处于行业领先水平。<br />
//
// 提示：对于一般开发者，我们建议优先使用SDK接入简化开发。SDK使用介绍请直接查看 5. 开发者资源 部分。
//
// 可能返回的错误码:
//  FAILEDOPERATION_NOFREEAMOUNT = "FailedOperation.NoFreeAmount"
//  FAILEDOPERATION_SERVICEISOLATE = "FailedOperation.ServiceIsolate"
//  FAILEDOPERATION_USERNOTREGISTERED = "FailedOperation.UserNotRegistered"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_BACKENDTIMEOUT = "InternalError.BackendTimeout"
//  INTERNALERROR_ERRORGETROUTE = "InternalError.ErrorGetRoute"
//  INTERNALERROR_ERRORUNKNOWN = "InternalError.ErrorUnknown"
//  INTERNALERROR_REQUESTFAILED = "InternalError.RequestFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_LIMITEDACCESSFREQUENCY = "LimitExceeded.LimitedAccessFrequency"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED_UINLIMITEXCEEDED = "RequestLimitExceeded.UinLimitExceeded"
//  UNAUTHORIZEDOPERATION_ACTIONNOTFOUND = "UnauthorizedOperation.ActionNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_TEXTTOOLONG = "UnsupportedOperation.TextTooLong"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDTARGETLANGUAGE = "UnsupportedOperation.UnSupportedTargetLanguage"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDLANGUAGE = "UnsupportedOperation.UnsupportedLanguage"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDSOURCELANGUAGE = "UnsupportedOperation.UnsupportedSourceLanguage"
func (c *Client) TextTranslateWithContext(ctx context.Context, request *TextTranslateRequest) (response *TextTranslateResponse, err error) {
    if request == nil {
        request = NewTextTranslateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("TextTranslate require credential")
    }

    request.SetContext(ctx)
    
    response = NewTextTranslateResponse()
    err = c.Send(request, response)
    return
}

func NewTextTranslateBatchRequest() (request *TextTranslateBatchRequest) {
    request = &TextTranslateBatchRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tmt", APIVersion, "TextTranslateBatch")
    
    
    return
}

func NewTextTranslateBatchResponse() (response *TextTranslateBatchResponse) {
    response = &TextTranslateBatchResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// TextTranslateBatch
// 文本翻译的批量接口
//
// 可能返回的错误码:
//  FAILEDOPERATION_NOFREEAMOUNT = "FailedOperation.NoFreeAmount"
//  FAILEDOPERATION_SERVICEISOLATE = "FailedOperation.ServiceIsolate"
//  FAILEDOPERATION_USERNOTREGISTERED = "FailedOperation.UserNotRegistered"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_BACKENDTIMEOUT = "InternalError.BackendTimeout"
//  INTERNALERROR_ERRORUNKNOWN = "InternalError.ErrorUnknown"
//  INTERNALERROR_REQUESTFAILED = "InternalError.RequestFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_LIMITEDACCESSFREQUENCY = "LimitExceeded.LimitedAccessFrequency"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED_UINLIMITEXCEEDED = "RequestLimitExceeded.UinLimitExceeded"
//  UNAUTHORIZEDOPERATION_ACTIONNOTFOUND = "UnauthorizedOperation.ActionNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_TEXTTOOLONG = "UnsupportedOperation.TextTooLong"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDTARGETLANGUAGE = "UnsupportedOperation.UnSupportedTargetLanguage"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDLANGUAGE = "UnsupportedOperation.UnsupportedLanguage"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDSOURCELANGUAGE = "UnsupportedOperation.UnsupportedSourceLanguage"
func (c *Client) TextTranslateBatch(request *TextTranslateBatchRequest) (response *TextTranslateBatchResponse, err error) {
    return c.TextTranslateBatchWithContext(context.Background(), request)
}

// TextTranslateBatch
// 文本翻译的批量接口
//
// 可能返回的错误码:
//  FAILEDOPERATION_NOFREEAMOUNT = "FailedOperation.NoFreeAmount"
//  FAILEDOPERATION_SERVICEISOLATE = "FailedOperation.ServiceIsolate"
//  FAILEDOPERATION_USERNOTREGISTERED = "FailedOperation.UserNotRegistered"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_BACKENDTIMEOUT = "InternalError.BackendTimeout"
//  INTERNALERROR_ERRORUNKNOWN = "InternalError.ErrorUnknown"
//  INTERNALERROR_REQUESTFAILED = "InternalError.RequestFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_LIMITEDACCESSFREQUENCY = "LimitExceeded.LimitedAccessFrequency"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED_UINLIMITEXCEEDED = "RequestLimitExceeded.UinLimitExceeded"
//  UNAUTHORIZEDOPERATION_ACTIONNOTFOUND = "UnauthorizedOperation.ActionNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_TEXTTOOLONG = "UnsupportedOperation.TextTooLong"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDTARGETLANGUAGE = "UnsupportedOperation.UnSupportedTargetLanguage"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDLANGUAGE = "UnsupportedOperation.UnsupportedLanguage"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDSOURCELANGUAGE = "UnsupportedOperation.UnsupportedSourceLanguage"
func (c *Client) TextTranslateBatchWithContext(ctx context.Context, request *TextTranslateBatchRequest) (response *TextTranslateBatchResponse, err error) {
    if request == nil {
        request = NewTextTranslateBatchRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("TextTranslateBatch require credential")
    }

    request.SetContext(ctx)
    
    response = NewTextTranslateBatchResponse()
    err = c.Send(request, response)
    return
}
