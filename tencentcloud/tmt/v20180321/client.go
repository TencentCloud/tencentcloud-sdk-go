// Copyright (c) 2017-2025 Tencent. All Rights Reserved.
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


func NewImageTranslateLLMRequest() (request *ImageTranslateLLMRequest) {
    request = &ImageTranslateLLMRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tmt", APIVersion, "ImageTranslateLLM")
    
    
    return
}

func NewImageTranslateLLMResponse() (response *ImageTranslateLLMResponse) {
    response = &ImageTranslateLLMResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ImageTranslateLLM
// 提供18种语言的图片翻译服务，可自动识别图片中的文本内容并翻译成目标语言，识别后的文本按行翻译，后续会提供可按段落翻译的版本。
//
// 
//
// - 输入图片格式：png、jpg、jpeg等常用图片格式，不支持gif动图。
//
// - 输出图片格式：jpg。
//
// 
//
// 提示：对于一般开发者，我们建议优先使用SDK接入简化开发。SDK使用介绍请直接查看 5. 开发者资源 部分。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DECODEERR = "FailedOperation.DecodeErr"
//  FAILEDOPERATION_DOWNLOADERR = "FailedOperation.DownloadErr"
//  FAILEDOPERATION_ERRORUSERAREA = "FailedOperation.ErrorUserArea"
//  FAILEDOPERATION_NOFREEAMOUNT = "FailedOperation.NoFreeAmount"
//  FAILEDOPERATION_SERVICEISOLATE = "FailedOperation.ServiceIsolate"
//  FAILEDOPERATION_STOPUSING = "FailedOperation.StopUsing"
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
//  UNSUPPORTEDOPERATION_UNSUPPORTEDTARGETLANGUAGE = "UnsupportedOperation.UnSupportedTargetLanguage"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDLANGUAGE = "UnsupportedOperation.UnsupportedLanguage"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDSOURCELANGUAGE = "UnsupportedOperation.UnsupportedSourceLanguage"
func (c *Client) ImageTranslateLLM(request *ImageTranslateLLMRequest) (response *ImageTranslateLLMResponse, err error) {
    return c.ImageTranslateLLMWithContext(context.Background(), request)
}

// ImageTranslateLLM
// 提供18种语言的图片翻译服务，可自动识别图片中的文本内容并翻译成目标语言，识别后的文本按行翻译，后续会提供可按段落翻译的版本。
//
// 
//
// - 输入图片格式：png、jpg、jpeg等常用图片格式，不支持gif动图。
//
// - 输出图片格式：jpg。
//
// 
//
// 提示：对于一般开发者，我们建议优先使用SDK接入简化开发。SDK使用介绍请直接查看 5. 开发者资源 部分。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DECODEERR = "FailedOperation.DecodeErr"
//  FAILEDOPERATION_DOWNLOADERR = "FailedOperation.DownloadErr"
//  FAILEDOPERATION_ERRORUSERAREA = "FailedOperation.ErrorUserArea"
//  FAILEDOPERATION_NOFREEAMOUNT = "FailedOperation.NoFreeAmount"
//  FAILEDOPERATION_SERVICEISOLATE = "FailedOperation.ServiceIsolate"
//  FAILEDOPERATION_STOPUSING = "FailedOperation.StopUsing"
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
//  UNSUPPORTEDOPERATION_UNSUPPORTEDTARGETLANGUAGE = "UnsupportedOperation.UnSupportedTargetLanguage"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDLANGUAGE = "UnsupportedOperation.UnsupportedLanguage"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDSOURCELANGUAGE = "UnsupportedOperation.UnsupportedSourceLanguage"
func (c *Client) ImageTranslateLLMWithContext(ctx context.Context, request *ImageTranslateLLMRequest) (response *ImageTranslateLLMResponse, err error) {
    if request == nil {
        request = NewImageTranslateLLMRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tmt", APIVersion, "ImageTranslateLLM")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ImageTranslateLLM require credential")
    }

    request.SetContext(ctx)
    
    response = NewImageTranslateLLMResponse()
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
// 腾讯翻译为合作伙伴提供文本翻译、文档翻译、交互翻译、AI同传等多种机器翻译服务，具有toB多行业解决方案。作为WMT世界机器翻译大赛冠军，翻译准确度值得信赖，其中，交互翻译能力是业界领先技术；腾讯同传是AI同传业界标杆。<br />
//
// 提示：对于一般开发者，我们建议优先使用SDK接入简化开发。SDK使用介绍请直接查看 5. 开发者资源 部分。
//
// 可能返回的错误码:
//  FAILEDOPERATION_ERRORUSERAREA = "FailedOperation.ErrorUserArea"
//  FAILEDOPERATION_LANGUAGERECOGNITIONERR = "FailedOperation.LanguageRecognitionErr"
//  FAILEDOPERATION_NOFREEAMOUNT = "FailedOperation.NoFreeAmount"
//  FAILEDOPERATION_REQUESTAILABERR = "FailedOperation.RequestAiLabErr"
//  FAILEDOPERATION_SERVICEISOLATE = "FailedOperation.ServiceIsolate"
//  FAILEDOPERATION_STOPUSING = "FailedOperation.StopUsing"
//  FAILEDOPERATION_USERNOTREGISTERED = "FailedOperation.UserNotRegistered"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_BACKENDTIMEOUT = "InternalError.BackendTimeout"
//  INTERNALERROR_ERRORGETROUTE = "InternalError.ErrorGetRoute"
//  INTERNALERROR_ERRORUNKNOWN = "InternalError.ErrorUnknown"
//  INTERNALERROR_REQUESTFAILED = "InternalError.RequestFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
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
// 腾讯翻译为合作伙伴提供文本翻译、文档翻译、交互翻译、AI同传等多种机器翻译服务，具有toB多行业解决方案。作为WMT世界机器翻译大赛冠军，翻译准确度值得信赖，其中，交互翻译能力是业界领先技术；腾讯同传是AI同传业界标杆。<br />
//
// 提示：对于一般开发者，我们建议优先使用SDK接入简化开发。SDK使用介绍请直接查看 5. 开发者资源 部分。
//
// 可能返回的错误码:
//  FAILEDOPERATION_ERRORUSERAREA = "FailedOperation.ErrorUserArea"
//  FAILEDOPERATION_LANGUAGERECOGNITIONERR = "FailedOperation.LanguageRecognitionErr"
//  FAILEDOPERATION_NOFREEAMOUNT = "FailedOperation.NoFreeAmount"
//  FAILEDOPERATION_REQUESTAILABERR = "FailedOperation.RequestAiLabErr"
//  FAILEDOPERATION_SERVICEISOLATE = "FailedOperation.ServiceIsolate"
//  FAILEDOPERATION_STOPUSING = "FailedOperation.StopUsing"
//  FAILEDOPERATION_USERNOTREGISTERED = "FailedOperation.UserNotRegistered"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_BACKENDTIMEOUT = "InternalError.BackendTimeout"
//  INTERNALERROR_ERRORGETROUTE = "InternalError.ErrorGetRoute"
//  INTERNALERROR_ERRORUNKNOWN = "InternalError.ErrorUnknown"
//  INTERNALERROR_REQUESTFAILED = "InternalError.RequestFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
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
    c.InitBaseRequest(&request.BaseRequest, "tmt", APIVersion, "TextTranslate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("TextTranslate require credential")
    }

    request.SetContext(ctx)
    
    response = NewTextTranslateResponse()
    err = c.Send(request, response)
    return
}
