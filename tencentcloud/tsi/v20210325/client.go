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

package v20210325

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2021-03-25"

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


func NewTongChuanDisplayRequest() (request *TongChuanDisplayRequest) {
    request = &TongChuanDisplayRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsi", APIVersion, "TongChuanDisplay")
    
    
    return
}

func NewTongChuanDisplayResponse() (response *TongChuanDisplayResponse) {
    response = &TongChuanDisplayResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// TongChuanDisplay
// 获取同传结果。
//
// 可能返回的错误码:
//  UNSUPPORTEDOPERATION_ISNEW = "UnsupportedOperation.IsNew"
//  UNSUPPORTEDOPERATION_SEMAX = "UnsupportedOperation.SeMax"
func (c *Client) TongChuanDisplay(request *TongChuanDisplayRequest) (response *TongChuanDisplayResponse, err error) {
    return c.TongChuanDisplayWithContext(context.Background(), request)
}

// TongChuanDisplay
// 获取同传结果。
//
// 可能返回的错误码:
//  UNSUPPORTEDOPERATION_ISNEW = "UnsupportedOperation.IsNew"
//  UNSUPPORTEDOPERATION_SEMAX = "UnsupportedOperation.SeMax"
func (c *Client) TongChuanDisplayWithContext(ctx context.Context, request *TongChuanDisplayRequest) (response *TongChuanDisplayResponse, err error) {
    if request == nil {
        request = NewTongChuanDisplayRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("TongChuanDisplay require credential")
    }

    request.SetContext(ctx)
    
    response = NewTongChuanDisplayResponse()
    err = c.Send(request, response)
    return
}

func NewTongChuanRecognizeRequest() (request *TongChuanRecognizeRequest) {
    request = &TongChuanRecognizeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsi", APIVersion, "TongChuanRecognize")
    
    
    return
}

func NewTongChuanRecognizeResponse() (response *TongChuanRecognizeResponse) {
    response = &TongChuanRecognizeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// TongChuanRecognize
// 本接口提供上传音频，将音频进行语音识别并翻译成文本的服务，目前开放中英互译的同传服务。 待识别和翻译的音频文件格式是 pcm，pcm采样率要求16kHz、位深16bit、单声道、每个分片时长200ms~500ms，音频内语音清晰。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  UNSUPPORTEDOPERATION_AUDIOFORMAT = "UnsupportedOperation.AudioFormat"
//  UNSUPPORTEDOPERATION_ISEND = "UnsupportedOperation.IsEnd"
//  UNSUPPORTEDOPERATION_TRANSLATETIME = "UnsupportedOperation.TranslateTime"
func (c *Client) TongChuanRecognize(request *TongChuanRecognizeRequest) (response *TongChuanRecognizeResponse, err error) {
    return c.TongChuanRecognizeWithContext(context.Background(), request)
}

// TongChuanRecognize
// 本接口提供上传音频，将音频进行语音识别并翻译成文本的服务，目前开放中英互译的同传服务。 待识别和翻译的音频文件格式是 pcm，pcm采样率要求16kHz、位深16bit、单声道、每个分片时长200ms~500ms，音频内语音清晰。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  UNSUPPORTEDOPERATION_AUDIOFORMAT = "UnsupportedOperation.AudioFormat"
//  UNSUPPORTEDOPERATION_ISEND = "UnsupportedOperation.IsEnd"
//  UNSUPPORTEDOPERATION_TRANSLATETIME = "UnsupportedOperation.TranslateTime"
func (c *Client) TongChuanRecognizeWithContext(ctx context.Context, request *TongChuanRecognizeRequest) (response *TongChuanRecognizeResponse, err error) {
    if request == nil {
        request = NewTongChuanRecognizeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("TongChuanRecognize require credential")
    }

    request.SetContext(ctx)
    
    response = NewTongChuanRecognizeResponse()
    err = c.Send(request, response)
    return
}

func NewTongChuanSyncRequest() (request *TongChuanSyncRequest) {
    request = &TongChuanSyncRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsi", APIVersion, "TongChuanSync")
    
    
    return
}

func NewTongChuanSyncResponse() (response *TongChuanSyncResponse) {
    response = &TongChuanSyncResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// TongChuanSync
// 本接口提供上传音频，将音频进行语音识别并翻译成文本的服务，目前开放中英互译的同传服务。 待识别和翻译的音频文件格式是 pcm，pcm采样率要求16kHz、位深16bit、单声道、每个分片时长200ms~500ms，音频内语音清晰。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  UNSUPPORTEDOPERATION_AUDIOFORMAT = "UnsupportedOperation.AudioFormat"
//  UNSUPPORTEDOPERATION_ISEND = "UnsupportedOperation.IsEnd"
//  UNSUPPORTEDOPERATION_TRANSLATETIME = "UnsupportedOperation.TranslateTime"
func (c *Client) TongChuanSync(request *TongChuanSyncRequest) (response *TongChuanSyncResponse, err error) {
    return c.TongChuanSyncWithContext(context.Background(), request)
}

// TongChuanSync
// 本接口提供上传音频，将音频进行语音识别并翻译成文本的服务，目前开放中英互译的同传服务。 待识别和翻译的音频文件格式是 pcm，pcm采样率要求16kHz、位深16bit、单声道、每个分片时长200ms~500ms，音频内语音清晰。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  UNSUPPORTEDOPERATION_AUDIOFORMAT = "UnsupportedOperation.AudioFormat"
//  UNSUPPORTEDOPERATION_ISEND = "UnsupportedOperation.IsEnd"
//  UNSUPPORTEDOPERATION_TRANSLATETIME = "UnsupportedOperation.TranslateTime"
func (c *Client) TongChuanSyncWithContext(ctx context.Context, request *TongChuanSyncRequest) (response *TongChuanSyncResponse, err error) {
    if request == nil {
        request = NewTongChuanSyncRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("TongChuanSync require credential")
    }

    request.SetContext(ctx)
    
    response = NewTongChuanSyncResponse()
    err = c.Send(request, response)
    return
}
