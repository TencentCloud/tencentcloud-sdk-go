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

func NewClient(credential *common.Credential, region string, clientProfile *profile.ClientProfile) (client *Client, err error) {
    client = &Client{}
    client.Init(region).
        WithCredential(credential).
        WithProfile(clientProfile)
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
//  UNSUPPORTEDOPERATION_NOBANLANCE = "UnsupportedOperation.NoBanlance"
//  UNSUPPORTEDOPERATION_NOFREEACCOUNT = "UnsupportedOperation.NoFreeAccount"
//  UNSUPPORTEDOPERATION_SERVERALREADYOPEN = "UnsupportedOperation.ServerAlreadyOpen"
//  UNSUPPORTEDOPERATION_SERVERNOTOPEN = "UnsupportedOperation.ServerNotOpen"
//  UNSUPPORTEDOPERATION_TEXTTOOLONG = "UnsupportedOperation.TextTooLong"
func (c *Client) TextToVoice(request *TextToVoiceRequest) (response *TextToVoiceResponse, err error) {
    if request == nil {
        request = NewTextToVoiceRequest()
    }
    response = NewTextToVoiceResponse()
    err = c.Send(request, response)
    return
}
