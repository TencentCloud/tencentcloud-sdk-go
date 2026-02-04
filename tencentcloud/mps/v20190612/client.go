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

package v20190612

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2019-06-12"

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


func NewBatchDeleteStreamLinkFlowRequest() (request *BatchDeleteStreamLinkFlowRequest) {
    request = &BatchDeleteStreamLinkFlowRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "BatchDeleteStreamLinkFlow")
    
    
    return
}

func NewBatchDeleteStreamLinkFlowResponse() (response *BatchDeleteStreamLinkFlowResponse) {
    response = &BatchDeleteStreamLinkFlowResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// BatchDeleteStreamLinkFlow
// 批量删除媒体传输流。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_OUTPUTGROUPS = "InvalidParameter.OutputGroups"
//  INVALIDPARAMETER_STATE = "InvalidParameter.State"
func (c *Client) BatchDeleteStreamLinkFlow(request *BatchDeleteStreamLinkFlowRequest) (response *BatchDeleteStreamLinkFlowResponse, err error) {
    return c.BatchDeleteStreamLinkFlowWithContext(context.Background(), request)
}

// BatchDeleteStreamLinkFlow
// 批量删除媒体传输流。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_OUTPUTGROUPS = "InvalidParameter.OutputGroups"
//  INVALIDPARAMETER_STATE = "InvalidParameter.State"
func (c *Client) BatchDeleteStreamLinkFlowWithContext(ctx context.Context, request *BatchDeleteStreamLinkFlowRequest) (response *BatchDeleteStreamLinkFlowResponse, err error) {
    if request == nil {
        request = NewBatchDeleteStreamLinkFlowRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "BatchDeleteStreamLinkFlow")
    
    if c.GetCredential() == nil {
        return nil, errors.New("BatchDeleteStreamLinkFlow require credential")
    }

    request.SetContext(ctx)
    
    response = NewBatchDeleteStreamLinkFlowResponse()
    err = c.Send(request, response)
    return
}

func NewBatchProcessMediaRequest() (request *BatchProcessMediaRequest) {
    request = &BatchProcessMediaRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "BatchProcessMedia")
    
    
    return
}

func NewBatchProcessMediaResponse() (response *BatchProcessMediaResponse) {
    response = &BatchProcessMediaResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// BatchProcessMedia
// 对 URL视频链接批量发起处理任务，功能包括：
//
// 智能字幕（语音全文、语音热词、语音翻译）
//
// 可能返回的错误码:
//  FAILEDOPERATION_GENERATERESOURCE = "FailedOperation.GenerateResource"
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_SESSIONCONTEXTTOOLONG = "InvalidParameterValue.SessionContextTooLong"
//  INVALIDPARAMETERVALUE_SESSIONID = "InvalidParameterValue.SessionId"
//  INVALIDPARAMETERVALUE_SESSIONIDTOOLONG = "InvalidParameterValue.SessionIdTooLong"
func (c *Client) BatchProcessMedia(request *BatchProcessMediaRequest) (response *BatchProcessMediaResponse, err error) {
    return c.BatchProcessMediaWithContext(context.Background(), request)
}

// BatchProcessMedia
// 对 URL视频链接批量发起处理任务，功能包括：
//
// 智能字幕（语音全文、语音热词、语音翻译）
//
// 可能返回的错误码:
//  FAILEDOPERATION_GENERATERESOURCE = "FailedOperation.GenerateResource"
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_SESSIONCONTEXTTOOLONG = "InvalidParameterValue.SessionContextTooLong"
//  INVALIDPARAMETERVALUE_SESSIONID = "InvalidParameterValue.SessionId"
//  INVALIDPARAMETERVALUE_SESSIONIDTOOLONG = "InvalidParameterValue.SessionIdTooLong"
func (c *Client) BatchProcessMediaWithContext(ctx context.Context, request *BatchProcessMediaRequest) (response *BatchProcessMediaResponse, err error) {
    if request == nil {
        request = NewBatchProcessMediaRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "BatchProcessMedia")
    
    if c.GetCredential() == nil {
        return nil, errors.New("BatchProcessMedia require credential")
    }

    request.SetContext(ctx)
    
    response = NewBatchProcessMediaResponse()
    err = c.Send(request, response)
    return
}

func NewBatchStartStreamLinkFlowRequest() (request *BatchStartStreamLinkFlowRequest) {
    request = &BatchStartStreamLinkFlowRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "BatchStartStreamLinkFlow")
    
    
    return
}

func NewBatchStartStreamLinkFlowResponse() (response *BatchStartStreamLinkFlowResponse) {
    response = &BatchStartStreamLinkFlowResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// BatchStartStreamLinkFlow
// 批量启动媒体传输流。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_OUTPUTGROUPS = "InvalidParameter.OutputGroups"
//  INVALIDPARAMETER_STATE = "InvalidParameter.State"
func (c *Client) BatchStartStreamLinkFlow(request *BatchStartStreamLinkFlowRequest) (response *BatchStartStreamLinkFlowResponse, err error) {
    return c.BatchStartStreamLinkFlowWithContext(context.Background(), request)
}

// BatchStartStreamLinkFlow
// 批量启动媒体传输流。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_OUTPUTGROUPS = "InvalidParameter.OutputGroups"
//  INVALIDPARAMETER_STATE = "InvalidParameter.State"
func (c *Client) BatchStartStreamLinkFlowWithContext(ctx context.Context, request *BatchStartStreamLinkFlowRequest) (response *BatchStartStreamLinkFlowResponse, err error) {
    if request == nil {
        request = NewBatchStartStreamLinkFlowRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "BatchStartStreamLinkFlow")
    
    if c.GetCredential() == nil {
        return nil, errors.New("BatchStartStreamLinkFlow require credential")
    }

    request.SetContext(ctx)
    
    response = NewBatchStartStreamLinkFlowResponse()
    err = c.Send(request, response)
    return
}

func NewBatchStopStreamLinkFlowRequest() (request *BatchStopStreamLinkFlowRequest) {
    request = &BatchStopStreamLinkFlowRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "BatchStopStreamLinkFlow")
    
    
    return
}

func NewBatchStopStreamLinkFlowResponse() (response *BatchStopStreamLinkFlowResponse) {
    response = &BatchStopStreamLinkFlowResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// BatchStopStreamLinkFlow
// 批量停止媒体传输流。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_STATE = "InvalidParameter.State"
func (c *Client) BatchStopStreamLinkFlow(request *BatchStopStreamLinkFlowRequest) (response *BatchStopStreamLinkFlowResponse, err error) {
    return c.BatchStopStreamLinkFlowWithContext(context.Background(), request)
}

// BatchStopStreamLinkFlow
// 批量停止媒体传输流。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_STATE = "InvalidParameter.State"
func (c *Client) BatchStopStreamLinkFlowWithContext(ctx context.Context, request *BatchStopStreamLinkFlowRequest) (response *BatchStopStreamLinkFlowResponse, err error) {
    if request == nil {
        request = NewBatchStopStreamLinkFlowRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "BatchStopStreamLinkFlow")
    
    if c.GetCredential() == nil {
        return nil, errors.New("BatchStopStreamLinkFlow require credential")
    }

    request.SetContext(ctx)
    
    response = NewBatchStopStreamLinkFlowResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAIAnalysisTemplateRequest() (request *CreateAIAnalysisTemplateRequest) {
    request = &CreateAIAnalysisTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "CreateAIAnalysisTemplate")
    
    
    return
}

func NewCreateAIAnalysisTemplateResponse() (response *CreateAIAnalysisTemplateResponse) {
    response = &CreateAIAnalysisTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateAIAnalysisTemplate
// 创建用户自定义内容分析模板，数量上限：50。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_CLASSIFCATIONCONFIGURE = "InvalidParameterValue.ClassifcationConfigure"
//  INVALIDPARAMETERVALUE_COMMENT = "InvalidParameterValue.Comment"
//  INVALIDPARAMETERVALUE_COVERCONFIGURE = "InvalidParameterValue.CoverConfigure"
//  INVALIDPARAMETERVALUE_DEFINITION = "InvalidParameterValue.Definition"
//  INVALIDPARAMETERVALUE_FRAMETAGCONFIGURE = "InvalidParameterValue.FrameTagConfigure"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  INVALIDPARAMETERVALUE_TAGCONFIGURE = "InvalidParameterValue.TagConfigure"
//  LIMITEXCEEDED_TOOMUCHTEMPLATE = "LimitExceeded.TooMuchTemplate"
func (c *Client) CreateAIAnalysisTemplate(request *CreateAIAnalysisTemplateRequest) (response *CreateAIAnalysisTemplateResponse, err error) {
    return c.CreateAIAnalysisTemplateWithContext(context.Background(), request)
}

// CreateAIAnalysisTemplate
// 创建用户自定义内容分析模板，数量上限：50。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_CLASSIFCATIONCONFIGURE = "InvalidParameterValue.ClassifcationConfigure"
//  INVALIDPARAMETERVALUE_COMMENT = "InvalidParameterValue.Comment"
//  INVALIDPARAMETERVALUE_COVERCONFIGURE = "InvalidParameterValue.CoverConfigure"
//  INVALIDPARAMETERVALUE_DEFINITION = "InvalidParameterValue.Definition"
//  INVALIDPARAMETERVALUE_FRAMETAGCONFIGURE = "InvalidParameterValue.FrameTagConfigure"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  INVALIDPARAMETERVALUE_TAGCONFIGURE = "InvalidParameterValue.TagConfigure"
//  LIMITEXCEEDED_TOOMUCHTEMPLATE = "LimitExceeded.TooMuchTemplate"
func (c *Client) CreateAIAnalysisTemplateWithContext(ctx context.Context, request *CreateAIAnalysisTemplateRequest) (response *CreateAIAnalysisTemplateResponse, err error) {
    if request == nil {
        request = NewCreateAIAnalysisTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "CreateAIAnalysisTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAIAnalysisTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAIAnalysisTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAIRecognitionTemplateRequest() (request *CreateAIRecognitionTemplateRequest) {
    request = &CreateAIRecognitionTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "CreateAIRecognitionTemplate")
    
    
    return
}

func NewCreateAIRecognitionTemplateResponse() (response *CreateAIRecognitionTemplateResponse) {
    response = &CreateAIRecognitionTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateAIRecognitionTemplate
// 创建用户自定义内容识别模板，数量上限：50。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_COMMENT = "InvalidParameterValue.Comment"
//  INVALIDPARAMETERVALUE_DEFAULTLIBRARYLABELSET = "InvalidParameterValue.DefaultLibraryLabelSet"
//  INVALIDPARAMETERVALUE_DESTINATIONLANGUAGE = "InvalidParameterValue.DestinationLanguage"
//  INVALIDPARAMETERVALUE_FACELIBRARY = "InvalidParameterValue.FaceLibrary"
//  INVALIDPARAMETERVALUE_FACESCORE = "InvalidParameterValue.FaceScore"
//  INVALIDPARAMETERVALUE_LABELSET = "InvalidParameterValue.LabelSet"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  INVALIDPARAMETERVALUE_OBJECTLIBRARY = "InvalidParameterValue.ObjectLibrary"
//  INVALIDPARAMETERVALUE_SOURCELANGUAGE = "InvalidParameterValue.SourceLanguage"
//  INVALIDPARAMETERVALUE_SUBTITLEFORMAT = "InvalidParameterValue.SubtitleFormat"
//  INVALIDPARAMETERVALUE_SWITCH = "InvalidParameterValue.Switch"
//  INVALIDPARAMETERVALUE_USERDEFINELIBRARYLABELSET = "InvalidParameterValue.UserDefineLibraryLabelSet"
//  LIMITEXCEEDED_TOOMUCHTEMPLATE = "LimitExceeded.TooMuchTemplate"
func (c *Client) CreateAIRecognitionTemplate(request *CreateAIRecognitionTemplateRequest) (response *CreateAIRecognitionTemplateResponse, err error) {
    return c.CreateAIRecognitionTemplateWithContext(context.Background(), request)
}

// CreateAIRecognitionTemplate
// 创建用户自定义内容识别模板，数量上限：50。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_COMMENT = "InvalidParameterValue.Comment"
//  INVALIDPARAMETERVALUE_DEFAULTLIBRARYLABELSET = "InvalidParameterValue.DefaultLibraryLabelSet"
//  INVALIDPARAMETERVALUE_DESTINATIONLANGUAGE = "InvalidParameterValue.DestinationLanguage"
//  INVALIDPARAMETERVALUE_FACELIBRARY = "InvalidParameterValue.FaceLibrary"
//  INVALIDPARAMETERVALUE_FACESCORE = "InvalidParameterValue.FaceScore"
//  INVALIDPARAMETERVALUE_LABELSET = "InvalidParameterValue.LabelSet"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  INVALIDPARAMETERVALUE_OBJECTLIBRARY = "InvalidParameterValue.ObjectLibrary"
//  INVALIDPARAMETERVALUE_SOURCELANGUAGE = "InvalidParameterValue.SourceLanguage"
//  INVALIDPARAMETERVALUE_SUBTITLEFORMAT = "InvalidParameterValue.SubtitleFormat"
//  INVALIDPARAMETERVALUE_SWITCH = "InvalidParameterValue.Switch"
//  INVALIDPARAMETERVALUE_USERDEFINELIBRARYLABELSET = "InvalidParameterValue.UserDefineLibraryLabelSet"
//  LIMITEXCEEDED_TOOMUCHTEMPLATE = "LimitExceeded.TooMuchTemplate"
func (c *Client) CreateAIRecognitionTemplateWithContext(ctx context.Context, request *CreateAIRecognitionTemplateRequest) (response *CreateAIRecognitionTemplateResponse, err error) {
    if request == nil {
        request = NewCreateAIRecognitionTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "CreateAIRecognitionTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAIRecognitionTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAIRecognitionTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAdaptiveDynamicStreamingTemplateRequest() (request *CreateAdaptiveDynamicStreamingTemplateRequest) {
    request = &CreateAdaptiveDynamicStreamingTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "CreateAdaptiveDynamicStreamingTemplate")
    
    
    return
}

func NewCreateAdaptiveDynamicStreamingTemplateResponse() (response *CreateAdaptiveDynamicStreamingTemplateResponse) {
    response = &CreateAdaptiveDynamicStreamingTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateAdaptiveDynamicStreamingTemplate
// 创建转自适应码流模板，数量上限：100。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_BITRATE = "InvalidParameterValue.Bitrate"
//  INVALIDPARAMETERVALUE_CODEC = "InvalidParameterValue.Codec"
//  INVALIDPARAMETERVALUE_DISABLEHIGHERVIDEOBITRATE = "InvalidParameterValue.DisableHigherVideoBitrate"
//  INVALIDPARAMETERVALUE_DISABLEHIGHERVIDEORESOLUTION = "InvalidParameterValue.DisableHigherVideoResolution"
//  INVALIDPARAMETERVALUE_FILLTYPE = "InvalidParameterValue.FillType"
//  INVALIDPARAMETERVALUE_FORMAT = "InvalidParameterValue.Format"
//  INVALIDPARAMETERVALUE_FPS = "InvalidParameterValue.Fps"
//  INVALIDPARAMETERVALUE_GOP = "InvalidParameterValue.Gop"
//  INVALIDPARAMETERVALUE_HEIGHT = "InvalidParameterValue.Height"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  INVALIDPARAMETERVALUE_REMOVEAUDIO = "InvalidParameterValue.RemoveAudio"
//  INVALIDPARAMETERVALUE_REMOVEVIDEO = "InvalidParameterValue.RemoveVideo"
//  INVALIDPARAMETERVALUE_SAMPLERATE = "InvalidParameterValue.SampleRate"
//  INVALIDPARAMETERVALUE_SOUNDSYSTEM = "InvalidParameterValue.SoundSystem"
//  INVALIDPARAMETERVALUE_WIDTH = "InvalidParameterValue.Width"
//  LIMITEXCEEDED_TOOMUCHTEMPLATE = "LimitExceeded.TooMuchTemplate"
func (c *Client) CreateAdaptiveDynamicStreamingTemplate(request *CreateAdaptiveDynamicStreamingTemplateRequest) (response *CreateAdaptiveDynamicStreamingTemplateResponse, err error) {
    return c.CreateAdaptiveDynamicStreamingTemplateWithContext(context.Background(), request)
}

// CreateAdaptiveDynamicStreamingTemplate
// 创建转自适应码流模板，数量上限：100。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_BITRATE = "InvalidParameterValue.Bitrate"
//  INVALIDPARAMETERVALUE_CODEC = "InvalidParameterValue.Codec"
//  INVALIDPARAMETERVALUE_DISABLEHIGHERVIDEOBITRATE = "InvalidParameterValue.DisableHigherVideoBitrate"
//  INVALIDPARAMETERVALUE_DISABLEHIGHERVIDEORESOLUTION = "InvalidParameterValue.DisableHigherVideoResolution"
//  INVALIDPARAMETERVALUE_FILLTYPE = "InvalidParameterValue.FillType"
//  INVALIDPARAMETERVALUE_FORMAT = "InvalidParameterValue.Format"
//  INVALIDPARAMETERVALUE_FPS = "InvalidParameterValue.Fps"
//  INVALIDPARAMETERVALUE_GOP = "InvalidParameterValue.Gop"
//  INVALIDPARAMETERVALUE_HEIGHT = "InvalidParameterValue.Height"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  INVALIDPARAMETERVALUE_REMOVEAUDIO = "InvalidParameterValue.RemoveAudio"
//  INVALIDPARAMETERVALUE_REMOVEVIDEO = "InvalidParameterValue.RemoveVideo"
//  INVALIDPARAMETERVALUE_SAMPLERATE = "InvalidParameterValue.SampleRate"
//  INVALIDPARAMETERVALUE_SOUNDSYSTEM = "InvalidParameterValue.SoundSystem"
//  INVALIDPARAMETERVALUE_WIDTH = "InvalidParameterValue.Width"
//  LIMITEXCEEDED_TOOMUCHTEMPLATE = "LimitExceeded.TooMuchTemplate"
func (c *Client) CreateAdaptiveDynamicStreamingTemplateWithContext(ctx context.Context, request *CreateAdaptiveDynamicStreamingTemplateRequest) (response *CreateAdaptiveDynamicStreamingTemplateResponse, err error) {
    if request == nil {
        request = NewCreateAdaptiveDynamicStreamingTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "CreateAdaptiveDynamicStreamingTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAdaptiveDynamicStreamingTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAdaptiveDynamicStreamingTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAigcImageTaskRequest() (request *CreateAigcImageTaskRequest) {
    request = &CreateAigcImageTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "CreateAigcImageTask")
    
    
    return
}

func NewCreateAigcImageTaskResponse() (response *CreateAigcImageTaskResponse) {
    response = &CreateAigcImageTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateAigcImageTask
// 调用该接口用于创建AIGC生图片任务。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_BITRATE = "InvalidParameterValue.Bitrate"
//  INVALIDPARAMETERVALUE_CODEC = "InvalidParameterValue.Codec"
//  INVALIDPARAMETERVALUE_DISABLEHIGHERVIDEOBITRATE = "InvalidParameterValue.DisableHigherVideoBitrate"
//  INVALIDPARAMETERVALUE_DISABLEHIGHERVIDEORESOLUTION = "InvalidParameterValue.DisableHigherVideoResolution"
//  INVALIDPARAMETERVALUE_FILLTYPE = "InvalidParameterValue.FillType"
//  INVALIDPARAMETERVALUE_FORMAT = "InvalidParameterValue.Format"
//  INVALIDPARAMETERVALUE_FPS = "InvalidParameterValue.Fps"
//  INVALIDPARAMETERVALUE_GOP = "InvalidParameterValue.Gop"
//  INVALIDPARAMETERVALUE_HEIGHT = "InvalidParameterValue.Height"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  INVALIDPARAMETERVALUE_REMOVEAUDIO = "InvalidParameterValue.RemoveAudio"
//  INVALIDPARAMETERVALUE_REMOVEVIDEO = "InvalidParameterValue.RemoveVideo"
//  INVALIDPARAMETERVALUE_SAMPLERATE = "InvalidParameterValue.SampleRate"
//  INVALIDPARAMETERVALUE_SOUNDSYSTEM = "InvalidParameterValue.SoundSystem"
//  INVALIDPARAMETERVALUE_WIDTH = "InvalidParameterValue.Width"
//  LIMITEXCEEDED_TOOMUCHTEMPLATE = "LimitExceeded.TooMuchTemplate"
func (c *Client) CreateAigcImageTask(request *CreateAigcImageTaskRequest) (response *CreateAigcImageTaskResponse, err error) {
    return c.CreateAigcImageTaskWithContext(context.Background(), request)
}

// CreateAigcImageTask
// 调用该接口用于创建AIGC生图片任务。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_BITRATE = "InvalidParameterValue.Bitrate"
//  INVALIDPARAMETERVALUE_CODEC = "InvalidParameterValue.Codec"
//  INVALIDPARAMETERVALUE_DISABLEHIGHERVIDEOBITRATE = "InvalidParameterValue.DisableHigherVideoBitrate"
//  INVALIDPARAMETERVALUE_DISABLEHIGHERVIDEORESOLUTION = "InvalidParameterValue.DisableHigherVideoResolution"
//  INVALIDPARAMETERVALUE_FILLTYPE = "InvalidParameterValue.FillType"
//  INVALIDPARAMETERVALUE_FORMAT = "InvalidParameterValue.Format"
//  INVALIDPARAMETERVALUE_FPS = "InvalidParameterValue.Fps"
//  INVALIDPARAMETERVALUE_GOP = "InvalidParameterValue.Gop"
//  INVALIDPARAMETERVALUE_HEIGHT = "InvalidParameterValue.Height"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  INVALIDPARAMETERVALUE_REMOVEAUDIO = "InvalidParameterValue.RemoveAudio"
//  INVALIDPARAMETERVALUE_REMOVEVIDEO = "InvalidParameterValue.RemoveVideo"
//  INVALIDPARAMETERVALUE_SAMPLERATE = "InvalidParameterValue.SampleRate"
//  INVALIDPARAMETERVALUE_SOUNDSYSTEM = "InvalidParameterValue.SoundSystem"
//  INVALIDPARAMETERVALUE_WIDTH = "InvalidParameterValue.Width"
//  LIMITEXCEEDED_TOOMUCHTEMPLATE = "LimitExceeded.TooMuchTemplate"
func (c *Client) CreateAigcImageTaskWithContext(ctx context.Context, request *CreateAigcImageTaskRequest) (response *CreateAigcImageTaskResponse, err error) {
    if request == nil {
        request = NewCreateAigcImageTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "CreateAigcImageTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAigcImageTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAigcImageTaskResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAigcVideoTaskRequest() (request *CreateAigcVideoTaskRequest) {
    request = &CreateAigcVideoTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "CreateAigcVideoTask")
    
    
    return
}

func NewCreateAigcVideoTaskResponse() (response *CreateAigcVideoTaskResponse) {
    response = &CreateAigcVideoTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateAigcVideoTask
// 调用该接口，用于创建AI生视频任务。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_BITRATE = "InvalidParameterValue.Bitrate"
//  INVALIDPARAMETERVALUE_CODEC = "InvalidParameterValue.Codec"
//  INVALIDPARAMETERVALUE_DISABLEHIGHERVIDEOBITRATE = "InvalidParameterValue.DisableHigherVideoBitrate"
//  INVALIDPARAMETERVALUE_DISABLEHIGHERVIDEORESOLUTION = "InvalidParameterValue.DisableHigherVideoResolution"
//  INVALIDPARAMETERVALUE_FILLTYPE = "InvalidParameterValue.FillType"
//  INVALIDPARAMETERVALUE_FORMAT = "InvalidParameterValue.Format"
//  INVALIDPARAMETERVALUE_FPS = "InvalidParameterValue.Fps"
//  INVALIDPARAMETERVALUE_GOP = "InvalidParameterValue.Gop"
//  INVALIDPARAMETERVALUE_HEIGHT = "InvalidParameterValue.Height"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  INVALIDPARAMETERVALUE_REMOVEAUDIO = "InvalidParameterValue.RemoveAudio"
//  INVALIDPARAMETERVALUE_REMOVEVIDEO = "InvalidParameterValue.RemoveVideo"
//  INVALIDPARAMETERVALUE_SAMPLERATE = "InvalidParameterValue.SampleRate"
//  INVALIDPARAMETERVALUE_SOUNDSYSTEM = "InvalidParameterValue.SoundSystem"
//  INVALIDPARAMETERVALUE_WIDTH = "InvalidParameterValue.Width"
//  LIMITEXCEEDED_TOOMUCHTEMPLATE = "LimitExceeded.TooMuchTemplate"
func (c *Client) CreateAigcVideoTask(request *CreateAigcVideoTaskRequest) (response *CreateAigcVideoTaskResponse, err error) {
    return c.CreateAigcVideoTaskWithContext(context.Background(), request)
}

// CreateAigcVideoTask
// 调用该接口，用于创建AI生视频任务。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_BITRATE = "InvalidParameterValue.Bitrate"
//  INVALIDPARAMETERVALUE_CODEC = "InvalidParameterValue.Codec"
//  INVALIDPARAMETERVALUE_DISABLEHIGHERVIDEOBITRATE = "InvalidParameterValue.DisableHigherVideoBitrate"
//  INVALIDPARAMETERVALUE_DISABLEHIGHERVIDEORESOLUTION = "InvalidParameterValue.DisableHigherVideoResolution"
//  INVALIDPARAMETERVALUE_FILLTYPE = "InvalidParameterValue.FillType"
//  INVALIDPARAMETERVALUE_FORMAT = "InvalidParameterValue.Format"
//  INVALIDPARAMETERVALUE_FPS = "InvalidParameterValue.Fps"
//  INVALIDPARAMETERVALUE_GOP = "InvalidParameterValue.Gop"
//  INVALIDPARAMETERVALUE_HEIGHT = "InvalidParameterValue.Height"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  INVALIDPARAMETERVALUE_REMOVEAUDIO = "InvalidParameterValue.RemoveAudio"
//  INVALIDPARAMETERVALUE_REMOVEVIDEO = "InvalidParameterValue.RemoveVideo"
//  INVALIDPARAMETERVALUE_SAMPLERATE = "InvalidParameterValue.SampleRate"
//  INVALIDPARAMETERVALUE_SOUNDSYSTEM = "InvalidParameterValue.SoundSystem"
//  INVALIDPARAMETERVALUE_WIDTH = "InvalidParameterValue.Width"
//  LIMITEXCEEDED_TOOMUCHTEMPLATE = "LimitExceeded.TooMuchTemplate"
func (c *Client) CreateAigcVideoTaskWithContext(ctx context.Context, request *CreateAigcVideoTaskRequest) (response *CreateAigcVideoTaskResponse, err error) {
    if request == nil {
        request = NewCreateAigcVideoTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "CreateAigcVideoTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAigcVideoTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAigcVideoTaskResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAnimatedGraphicsTemplateRequest() (request *CreateAnimatedGraphicsTemplateRequest) {
    request = &CreateAnimatedGraphicsTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "CreateAnimatedGraphicsTemplate")
    
    
    return
}

func NewCreateAnimatedGraphicsTemplateResponse() (response *CreateAnimatedGraphicsTemplateResponse) {
    response = &CreateAnimatedGraphicsTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateAnimatedGraphicsTemplate
// 创建用户自定义转动图模板，数量上限：16。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_COMMENT = "InvalidParameterValue.Comment"
//  INVALIDPARAMETERVALUE_FORMAT = "InvalidParameterValue.Format"
//  INVALIDPARAMETERVALUE_FORMATWEBPLACKWIDTHANDHEIGHT = "InvalidParameterValue.FormatWebpLackWidthAndHeight"
//  INVALIDPARAMETERVALUE_FORMATWEBPWIDTHANDHEIGHTBOTHZERO = "InvalidParameterValue.FormatWebpWidthAndHeightBothZero"
//  INVALIDPARAMETERVALUE_FPS = "InvalidParameterValue.Fps"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  INVALIDPARAMETERVALUE_QUALITY = "InvalidParameterValue.Quality"
//  INVALIDPARAMETERVALUE_RESOLUTION = "InvalidParameterValue.Resolution"
//  LIMITEXCEEDED_TOOMUCHTEMPLATE = "LimitExceeded.TooMuchTemplate"
func (c *Client) CreateAnimatedGraphicsTemplate(request *CreateAnimatedGraphicsTemplateRequest) (response *CreateAnimatedGraphicsTemplateResponse, err error) {
    return c.CreateAnimatedGraphicsTemplateWithContext(context.Background(), request)
}

// CreateAnimatedGraphicsTemplate
// 创建用户自定义转动图模板，数量上限：16。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_COMMENT = "InvalidParameterValue.Comment"
//  INVALIDPARAMETERVALUE_FORMAT = "InvalidParameterValue.Format"
//  INVALIDPARAMETERVALUE_FORMATWEBPLACKWIDTHANDHEIGHT = "InvalidParameterValue.FormatWebpLackWidthAndHeight"
//  INVALIDPARAMETERVALUE_FORMATWEBPWIDTHANDHEIGHTBOTHZERO = "InvalidParameterValue.FormatWebpWidthAndHeightBothZero"
//  INVALIDPARAMETERVALUE_FPS = "InvalidParameterValue.Fps"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  INVALIDPARAMETERVALUE_QUALITY = "InvalidParameterValue.Quality"
//  INVALIDPARAMETERVALUE_RESOLUTION = "InvalidParameterValue.Resolution"
//  LIMITEXCEEDED_TOOMUCHTEMPLATE = "LimitExceeded.TooMuchTemplate"
func (c *Client) CreateAnimatedGraphicsTemplateWithContext(ctx context.Context, request *CreateAnimatedGraphicsTemplateRequest) (response *CreateAnimatedGraphicsTemplateResponse, err error) {
    if request == nil {
        request = NewCreateAnimatedGraphicsTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "CreateAnimatedGraphicsTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAnimatedGraphicsTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAnimatedGraphicsTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAsrHotwordsRequest() (request *CreateAsrHotwordsRequest) {
    request = &CreateAsrHotwordsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "CreateAsrHotwords")
    
    
    return
}

func NewCreateAsrHotwordsResponse() (response *CreateAsrHotwordsResponse) {
    response = &CreateAsrHotwordsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateAsrHotwords
// 智能字幕新建热词库接口
//
// 可能返回的错误码:
//  INTERNALERROR_ACCESSDBERROR = "InternalError.AccessDBError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_HOTWORDSFORMATERROR = "InvalidParameterValue.HotwordsFormatError"
//  INVALIDPARAMETERVALUE_INPUTINFO = "InvalidParameterValue.InputInfo"
//  LIMITEXCEEDED_TOOMUCHHOTWORDS = "LimitExceeded.TooMuchHotWords"
//  LIMITEXCEEDED_TOOMUCHLARGEHOTWORDS = "LimitExceeded.TooMuchLargeHotWords"
func (c *Client) CreateAsrHotwords(request *CreateAsrHotwordsRequest) (response *CreateAsrHotwordsResponse, err error) {
    return c.CreateAsrHotwordsWithContext(context.Background(), request)
}

// CreateAsrHotwords
// 智能字幕新建热词库接口
//
// 可能返回的错误码:
//  INTERNALERROR_ACCESSDBERROR = "InternalError.AccessDBError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_HOTWORDSFORMATERROR = "InvalidParameterValue.HotwordsFormatError"
//  INVALIDPARAMETERVALUE_INPUTINFO = "InvalidParameterValue.InputInfo"
//  LIMITEXCEEDED_TOOMUCHHOTWORDS = "LimitExceeded.TooMuchHotWords"
//  LIMITEXCEEDED_TOOMUCHLARGEHOTWORDS = "LimitExceeded.TooMuchLargeHotWords"
func (c *Client) CreateAsrHotwordsWithContext(ctx context.Context, request *CreateAsrHotwordsRequest) (response *CreateAsrHotwordsResponse, err error) {
    if request == nil {
        request = NewCreateAsrHotwordsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "CreateAsrHotwords")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAsrHotwords require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAsrHotwordsResponse()
    err = c.Send(request, response)
    return
}

func NewCreateBlindWatermarkTemplateRequest() (request *CreateBlindWatermarkTemplateRequest) {
    request = &CreateBlindWatermarkTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "CreateBlindWatermarkTemplate")
    
    
    return
}

func NewCreateBlindWatermarkTemplateResponse() (response *CreateBlindWatermarkTemplateResponse) {
    response = &CreateBlindWatermarkTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateBlindWatermarkTemplate
// 创建用户自定义数字水印模板。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GENDEFINITION = "InternalError.GenDefinition"
//  INVALIDPARAMETERVALUE_COMMENT = "InvalidParameterValue.Comment"
//  INVALIDPARAMETERVALUE_DUPLICATEDTEXTCONTENT = "InvalidParameterValue.DuplicatedTextContent"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  LIMITEXCEEDED_TOOMUCHTEMPLATE = "LimitExceeded.TooMuchTemplate"
func (c *Client) CreateBlindWatermarkTemplate(request *CreateBlindWatermarkTemplateRequest) (response *CreateBlindWatermarkTemplateResponse, err error) {
    return c.CreateBlindWatermarkTemplateWithContext(context.Background(), request)
}

// CreateBlindWatermarkTemplate
// 创建用户自定义数字水印模板。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GENDEFINITION = "InternalError.GenDefinition"
//  INVALIDPARAMETERVALUE_COMMENT = "InvalidParameterValue.Comment"
//  INVALIDPARAMETERVALUE_DUPLICATEDTEXTCONTENT = "InvalidParameterValue.DuplicatedTextContent"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  LIMITEXCEEDED_TOOMUCHTEMPLATE = "LimitExceeded.TooMuchTemplate"
func (c *Client) CreateBlindWatermarkTemplateWithContext(ctx context.Context, request *CreateBlindWatermarkTemplateRequest) (response *CreateBlindWatermarkTemplateResponse, err error) {
    if request == nil {
        request = NewCreateBlindWatermarkTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "CreateBlindWatermarkTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateBlindWatermarkTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateBlindWatermarkTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewCreateContentReviewTemplateRequest() (request *CreateContentReviewTemplateRequest) {
    request = &CreateContentReviewTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "CreateContentReviewTemplate")
    
    
    return
}

func NewCreateContentReviewTemplateResponse() (response *CreateContentReviewTemplateResponse) {
    response = &CreateContentReviewTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateContentReviewTemplate
// 创建用户自定义内容审核模板，数量上限：50。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GENDEFINITION = "InternalError.GenDefinition"
//  INVALIDPARAMETERVALUE_BLOCKCONFIDENCE = "InvalidParameterValue.BlockConfidence"
//  INVALIDPARAMETERVALUE_COMMENT = "InvalidParameterValue.Comment"
//  INVALIDPARAMETERVALUE_LABELSET = "InvalidParameterValue.LabelSet"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  INVALIDPARAMETERVALUE_REVIEWCONFIDENCE = "InvalidParameterValue.ReviewConfidence"
//  INVALIDPARAMETERVALUE_SWITCH = "InvalidParameterValue.Switch"
//  LIMITEXCEEDED_TOOMUCHTEMPLATE = "LimitExceeded.TooMuchTemplate"
func (c *Client) CreateContentReviewTemplate(request *CreateContentReviewTemplateRequest) (response *CreateContentReviewTemplateResponse, err error) {
    return c.CreateContentReviewTemplateWithContext(context.Background(), request)
}

// CreateContentReviewTemplate
// 创建用户自定义内容审核模板，数量上限：50。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GENDEFINITION = "InternalError.GenDefinition"
//  INVALIDPARAMETERVALUE_BLOCKCONFIDENCE = "InvalidParameterValue.BlockConfidence"
//  INVALIDPARAMETERVALUE_COMMENT = "InvalidParameterValue.Comment"
//  INVALIDPARAMETERVALUE_LABELSET = "InvalidParameterValue.LabelSet"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  INVALIDPARAMETERVALUE_REVIEWCONFIDENCE = "InvalidParameterValue.ReviewConfidence"
//  INVALIDPARAMETERVALUE_SWITCH = "InvalidParameterValue.Switch"
//  LIMITEXCEEDED_TOOMUCHTEMPLATE = "LimitExceeded.TooMuchTemplate"
func (c *Client) CreateContentReviewTemplateWithContext(ctx context.Context, request *CreateContentReviewTemplateRequest) (response *CreateContentReviewTemplateResponse, err error) {
    if request == nil {
        request = NewCreateContentReviewTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "CreateContentReviewTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateContentReviewTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateContentReviewTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewCreateImageSpriteTemplateRequest() (request *CreateImageSpriteTemplateRequest) {
    request = &CreateImageSpriteTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "CreateImageSpriteTemplate")
    
    
    return
}

func NewCreateImageSpriteTemplateResponse() (response *CreateImageSpriteTemplateResponse) {
    response = &CreateImageSpriteTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateImageSpriteTemplate
// 创建用户自定义雪碧图模板，数量上限：16。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_COLUMNCOUNT = "InvalidParameterValue.ColumnCount"
//  INVALIDPARAMETERVALUE_FORMAT = "InvalidParameterValue.Format"
//  INVALIDPARAMETERVALUE_HEIGHT = "InvalidParameterValue.Height"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  INVALIDPARAMETERVALUE_RESOLUTION = "InvalidParameterValue.Resolution"
//  INVALIDPARAMETERVALUE_ROWCOUNT = "InvalidParameterValue.RowCount"
//  INVALIDPARAMETERVALUE_SAMPLEINTERVAL = "InvalidParameterValue.SampleInterval"
//  INVALIDPARAMETERVALUE_SAMPLETYPE = "InvalidParameterValue.SampleType"
//  INVALIDPARAMETERVALUE_WIDTH = "InvalidParameterValue.Width"
//  LIMITEXCEEDED_TOOMUCHTEMPLATE = "LimitExceeded.TooMuchTemplate"
func (c *Client) CreateImageSpriteTemplate(request *CreateImageSpriteTemplateRequest) (response *CreateImageSpriteTemplateResponse, err error) {
    return c.CreateImageSpriteTemplateWithContext(context.Background(), request)
}

// CreateImageSpriteTemplate
// 创建用户自定义雪碧图模板，数量上限：16。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_COLUMNCOUNT = "InvalidParameterValue.ColumnCount"
//  INVALIDPARAMETERVALUE_FORMAT = "InvalidParameterValue.Format"
//  INVALIDPARAMETERVALUE_HEIGHT = "InvalidParameterValue.Height"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  INVALIDPARAMETERVALUE_RESOLUTION = "InvalidParameterValue.Resolution"
//  INVALIDPARAMETERVALUE_ROWCOUNT = "InvalidParameterValue.RowCount"
//  INVALIDPARAMETERVALUE_SAMPLEINTERVAL = "InvalidParameterValue.SampleInterval"
//  INVALIDPARAMETERVALUE_SAMPLETYPE = "InvalidParameterValue.SampleType"
//  INVALIDPARAMETERVALUE_WIDTH = "InvalidParameterValue.Width"
//  LIMITEXCEEDED_TOOMUCHTEMPLATE = "LimitExceeded.TooMuchTemplate"
func (c *Client) CreateImageSpriteTemplateWithContext(ctx context.Context, request *CreateImageSpriteTemplateRequest) (response *CreateImageSpriteTemplateResponse, err error) {
    if request == nil {
        request = NewCreateImageSpriteTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "CreateImageSpriteTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateImageSpriteTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateImageSpriteTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewCreateLiveRecordTemplateRequest() (request *CreateLiveRecordTemplateRequest) {
    request = &CreateLiveRecordTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "CreateLiveRecordTemplate")
    
    
    return
}

func NewCreateLiveRecordTemplateResponse() (response *CreateLiveRecordTemplateResponse) {
    response = &CreateLiveRecordTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateLiveRecordTemplate
// 创建直播录制模板
//
// 可能返回的错误码:
//  INTERNALERROR_ACCESSDBERROR = "InternalError.AccessDBError"
//  INTERNALERROR_GENDEFINITION = "InternalError.GenDefinition"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_TOOMUCHTEMPLATE = "LimitExceeded.TooMuchTemplate"
func (c *Client) CreateLiveRecordTemplate(request *CreateLiveRecordTemplateRequest) (response *CreateLiveRecordTemplateResponse, err error) {
    return c.CreateLiveRecordTemplateWithContext(context.Background(), request)
}

// CreateLiveRecordTemplate
// 创建直播录制模板
//
// 可能返回的错误码:
//  INTERNALERROR_ACCESSDBERROR = "InternalError.AccessDBError"
//  INTERNALERROR_GENDEFINITION = "InternalError.GenDefinition"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_TOOMUCHTEMPLATE = "LimitExceeded.TooMuchTemplate"
func (c *Client) CreateLiveRecordTemplateWithContext(ctx context.Context, request *CreateLiveRecordTemplateRequest) (response *CreateLiveRecordTemplateResponse, err error) {
    if request == nil {
        request = NewCreateLiveRecordTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "CreateLiveRecordTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateLiveRecordTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateLiveRecordTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewCreateMediaEvaluationRequest() (request *CreateMediaEvaluationRequest) {
    request = &CreateMediaEvaluationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "CreateMediaEvaluation")
    
    
    return
}

func NewCreateMediaEvaluationResponse() (response *CreateMediaEvaluationResponse) {
    response = &CreateMediaEvaluationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateMediaEvaluation
// 发起视频评测任务，功能包括：
//
// 
//
// 1. 对一个原视频和多个转码后的视频进行评分。
//
// 2. 计算不同转码方式的 BD-Rate。
//
// 可能返回的错误码:
//  FAILEDOPERATION_GENERATERESOURCE = "FailedOperation.GenerateResource"
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_SESSIONCONTEXTTOOLONG = "InvalidParameterValue.SessionContextTooLong"
//  INVALIDPARAMETERVALUE_SESSIONID = "InvalidParameterValue.SessionId"
//  INVALIDPARAMETERVALUE_SESSIONIDTOOLONG = "InvalidParameterValue.SessionIdTooLong"
func (c *Client) CreateMediaEvaluation(request *CreateMediaEvaluationRequest) (response *CreateMediaEvaluationResponse, err error) {
    return c.CreateMediaEvaluationWithContext(context.Background(), request)
}

// CreateMediaEvaluation
// 发起视频评测任务，功能包括：
//
// 
//
// 1. 对一个原视频和多个转码后的视频进行评分。
//
// 2. 计算不同转码方式的 BD-Rate。
//
// 可能返回的错误码:
//  FAILEDOPERATION_GENERATERESOURCE = "FailedOperation.GenerateResource"
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_SESSIONCONTEXTTOOLONG = "InvalidParameterValue.SessionContextTooLong"
//  INVALIDPARAMETERVALUE_SESSIONID = "InvalidParameterValue.SessionId"
//  INVALIDPARAMETERVALUE_SESSIONIDTOOLONG = "InvalidParameterValue.SessionIdTooLong"
func (c *Client) CreateMediaEvaluationWithContext(ctx context.Context, request *CreateMediaEvaluationRequest) (response *CreateMediaEvaluationResponse, err error) {
    if request == nil {
        request = NewCreateMediaEvaluationRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "CreateMediaEvaluation")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateMediaEvaluation require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateMediaEvaluationResponse()
    err = c.Send(request, response)
    return
}

func NewCreatePersonSampleRequest() (request *CreatePersonSampleRequest) {
    request = &CreatePersonSampleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "CreatePersonSample")
    
    
    return
}

func NewCreatePersonSampleResponse() (response *CreatePersonSampleResponse) {
    response = &CreatePersonSampleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreatePersonSample
// 该接口用于创建素材样本，用于通过五官定位等技术，进行内容识别、内容不适宜等视频处理。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_FACEDUPLICATE = "InvalidParameterValue.FaceDuplicate"
//  INVALIDPARAMETERVALUE_PICFORMATERROR = "InvalidParameterValue.PicFormatError"
func (c *Client) CreatePersonSample(request *CreatePersonSampleRequest) (response *CreatePersonSampleResponse, err error) {
    return c.CreatePersonSampleWithContext(context.Background(), request)
}

// CreatePersonSample
// 该接口用于创建素材样本，用于通过五官定位等技术，进行内容识别、内容不适宜等视频处理。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_FACEDUPLICATE = "InvalidParameterValue.FaceDuplicate"
//  INVALIDPARAMETERVALUE_PICFORMATERROR = "InvalidParameterValue.PicFormatError"
func (c *Client) CreatePersonSampleWithContext(ctx context.Context, request *CreatePersonSampleRequest) (response *CreatePersonSampleResponse, err error) {
    if request == nil {
        request = NewCreatePersonSampleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "CreatePersonSample")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreatePersonSample require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreatePersonSampleResponse()
    err = c.Send(request, response)
    return
}

func NewCreateProcessImageTemplateRequest() (request *CreateProcessImageTemplateRequest) {
    request = &CreateProcessImageTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "CreateProcessImageTemplate")
    
    
    return
}

func NewCreateProcessImageTemplateResponse() (response *CreateProcessImageTemplateResponse) {
    response = &CreateProcessImageTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateProcessImageTemplate
// 创建图片处理模板
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GENDEFINITION = "InternalError.GenDefinition"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) CreateProcessImageTemplate(request *CreateProcessImageTemplateRequest) (response *CreateProcessImageTemplateResponse, err error) {
    return c.CreateProcessImageTemplateWithContext(context.Background(), request)
}

// CreateProcessImageTemplate
// 创建图片处理模板
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GENDEFINITION = "InternalError.GenDefinition"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) CreateProcessImageTemplateWithContext(ctx context.Context, request *CreateProcessImageTemplateRequest) (response *CreateProcessImageTemplateResponse, err error) {
    if request == nil {
        request = NewCreateProcessImageTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "CreateProcessImageTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateProcessImageTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateProcessImageTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewCreateQualityControlTemplateRequest() (request *CreateQualityControlTemplateRequest) {
    request = &CreateQualityControlTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "CreateQualityControlTemplate")
    
    
    return
}

func NewCreateQualityControlTemplateResponse() (response *CreateQualityControlTemplateResponse) {
    response = &CreateQualityControlTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateQualityControlTemplate
// 创建媒体质检模板，数量上限：50。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GENDEFINITION = "InternalError.GenDefinition"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_EMPTYDETECTITEM = "InvalidParameterValue.EmptyDetectItem"
//  INVALIDPARAMETERVALUE_UNKNOWNCATEGORY = "InvalidParameterValue.UnknownCategory"
func (c *Client) CreateQualityControlTemplate(request *CreateQualityControlTemplateRequest) (response *CreateQualityControlTemplateResponse, err error) {
    return c.CreateQualityControlTemplateWithContext(context.Background(), request)
}

// CreateQualityControlTemplate
// 创建媒体质检模板，数量上限：50。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GENDEFINITION = "InternalError.GenDefinition"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_EMPTYDETECTITEM = "InvalidParameterValue.EmptyDetectItem"
//  INVALIDPARAMETERVALUE_UNKNOWNCATEGORY = "InvalidParameterValue.UnknownCategory"
func (c *Client) CreateQualityControlTemplateWithContext(ctx context.Context, request *CreateQualityControlTemplateRequest) (response *CreateQualityControlTemplateResponse, err error) {
    if request == nil {
        request = NewCreateQualityControlTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "CreateQualityControlTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateQualityControlTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateQualityControlTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewCreateSampleSnapshotTemplateRequest() (request *CreateSampleSnapshotTemplateRequest) {
    request = &CreateSampleSnapshotTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "CreateSampleSnapshotTemplate")
    
    
    return
}

func NewCreateSampleSnapshotTemplateResponse() (response *CreateSampleSnapshotTemplateResponse) {
    response = &CreateSampleSnapshotTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateSampleSnapshotTemplate
// 创建用户自定义采样截图模板，数量上限：16。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_COMMENT = "InvalidParameterValue.Comment"
//  INVALIDPARAMETERVALUE_FORMAT = "InvalidParameterValue.Format"
//  INVALIDPARAMETERVALUE_HEIGHT = "InvalidParameterValue.Height"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  INVALIDPARAMETERVALUE_RESOLUTION = "InvalidParameterValue.Resolution"
//  INVALIDPARAMETERVALUE_SAMPLEINTERVAL = "InvalidParameterValue.SampleInterval"
//  INVALIDPARAMETERVALUE_SAMPLETYPE = "InvalidParameterValue.SampleType"
//  INVALIDPARAMETERVALUE_WIDTH = "InvalidParameterValue.Width"
//  LIMITEXCEEDED_TOOMUCHTEMPLATE = "LimitExceeded.TooMuchTemplate"
func (c *Client) CreateSampleSnapshotTemplate(request *CreateSampleSnapshotTemplateRequest) (response *CreateSampleSnapshotTemplateResponse, err error) {
    return c.CreateSampleSnapshotTemplateWithContext(context.Background(), request)
}

// CreateSampleSnapshotTemplate
// 创建用户自定义采样截图模板，数量上限：16。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_COMMENT = "InvalidParameterValue.Comment"
//  INVALIDPARAMETERVALUE_FORMAT = "InvalidParameterValue.Format"
//  INVALIDPARAMETERVALUE_HEIGHT = "InvalidParameterValue.Height"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  INVALIDPARAMETERVALUE_RESOLUTION = "InvalidParameterValue.Resolution"
//  INVALIDPARAMETERVALUE_SAMPLEINTERVAL = "InvalidParameterValue.SampleInterval"
//  INVALIDPARAMETERVALUE_SAMPLETYPE = "InvalidParameterValue.SampleType"
//  INVALIDPARAMETERVALUE_WIDTH = "InvalidParameterValue.Width"
//  LIMITEXCEEDED_TOOMUCHTEMPLATE = "LimitExceeded.TooMuchTemplate"
func (c *Client) CreateSampleSnapshotTemplateWithContext(ctx context.Context, request *CreateSampleSnapshotTemplateRequest) (response *CreateSampleSnapshotTemplateResponse, err error) {
    if request == nil {
        request = NewCreateSampleSnapshotTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "CreateSampleSnapshotTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateSampleSnapshotTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateSampleSnapshotTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewCreateScheduleRequest() (request *CreateScheduleRequest) {
    request = &CreateScheduleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "CreateSchedule")
    
    
    return
}

func NewCreateScheduleResponse() (response *CreateScheduleResponse) {
    response = &CreateScheduleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateSchedule
// 对 COS 中指定 Bucket 的目录下上传的媒体文件，设置处理规则，包括：
//
// 1. 视频转码（带水印）；
//
// 2. 视频转动图；
//
// 3. 对视频按指定时间点截图；
//
// 4. 对视频采样截图；
//
// 5. 对视频截图雪碧图；
//
// 6. 对视频转自适应码流；
//
// 7. 智能内容审核（鉴黄、敏感信息检测）；
//
// 8. 智能内容分析（标签、分类、封面、按帧标签）；
//
// 9. 智能内容识别（人脸、文本全文、文本关键词、语音全文、语音关键词）。
//
// 10. 媒体质检（直播流格式诊断、音画内容检测（抖动、模糊、低光照、过曝光、黑边、白边、黑屏、白屏、花屏、噪点、马赛克、二维码等）、无参考打分）
//
// 11. 智能字幕（语音全文、语音热词、语音翻译）
//
// 12. 智能擦除（去水印、去字幕、隐私保护）；
//
// 
//
// 注意：创建编排成功后是禁用状态，需要手动启用。
//
// 可能返回的错误码:
//  FAILEDOPERATION_COSSTATUSINAVLID = "FailedOperation.CosStatusInavlid"
//  FAILEDOPERATION_GENERATERESOURCE = "FailedOperation.GenerateResource"
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ACCESSDBERROR = "InternalError.AccessDBError"
//  INTERNALERROR_GENDEFINITION = "InternalError.GenDefinition"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_TOOMUCHTEMPLATE = "LimitExceeded.TooMuchTemplate"
//  RESOURCENOTFOUND_COSBUCKETNAMEINVALID = "ResourceNotFound.CosBucketNameInvalid"
//  RESOURCENOTFOUND_COSBUCKETNOTEXIST = "ResourceNotFound.CosBucketNotExist"
func (c *Client) CreateSchedule(request *CreateScheduleRequest) (response *CreateScheduleResponse, err error) {
    return c.CreateScheduleWithContext(context.Background(), request)
}

// CreateSchedule
// 对 COS 中指定 Bucket 的目录下上传的媒体文件，设置处理规则，包括：
//
// 1. 视频转码（带水印）；
//
// 2. 视频转动图；
//
// 3. 对视频按指定时间点截图；
//
// 4. 对视频采样截图；
//
// 5. 对视频截图雪碧图；
//
// 6. 对视频转自适应码流；
//
// 7. 智能内容审核（鉴黄、敏感信息检测）；
//
// 8. 智能内容分析（标签、分类、封面、按帧标签）；
//
// 9. 智能内容识别（人脸、文本全文、文本关键词、语音全文、语音关键词）。
//
// 10. 媒体质检（直播流格式诊断、音画内容检测（抖动、模糊、低光照、过曝光、黑边、白边、黑屏、白屏、花屏、噪点、马赛克、二维码等）、无参考打分）
//
// 11. 智能字幕（语音全文、语音热词、语音翻译）
//
// 12. 智能擦除（去水印、去字幕、隐私保护）；
//
// 
//
// 注意：创建编排成功后是禁用状态，需要手动启用。
//
// 可能返回的错误码:
//  FAILEDOPERATION_COSSTATUSINAVLID = "FailedOperation.CosStatusInavlid"
//  FAILEDOPERATION_GENERATERESOURCE = "FailedOperation.GenerateResource"
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ACCESSDBERROR = "InternalError.AccessDBError"
//  INTERNALERROR_GENDEFINITION = "InternalError.GenDefinition"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_TOOMUCHTEMPLATE = "LimitExceeded.TooMuchTemplate"
//  RESOURCENOTFOUND_COSBUCKETNAMEINVALID = "ResourceNotFound.CosBucketNameInvalid"
//  RESOURCENOTFOUND_COSBUCKETNOTEXIST = "ResourceNotFound.CosBucketNotExist"
func (c *Client) CreateScheduleWithContext(ctx context.Context, request *CreateScheduleRequest) (response *CreateScheduleResponse, err error) {
    if request == nil {
        request = NewCreateScheduleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "CreateSchedule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateSchedule require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateScheduleResponse()
    err = c.Send(request, response)
    return
}

func NewCreateSmartEraseTemplateRequest() (request *CreateSmartEraseTemplateRequest) {
    request = &CreateSmartEraseTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "CreateSmartEraseTemplate")
    
    
    return
}

func NewCreateSmartEraseTemplateResponse() (response *CreateSmartEraseTemplateResponse) {
    response = &CreateSmartEraseTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateSmartEraseTemplate
// 创建自定义智能擦除模板
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_AUTOAREAS = "InvalidParameterValue.AutoAreas"
//  INVALIDPARAMETERVALUE_COMMENT = "InvalidParameterValue.Comment"
//  INVALIDPARAMETERVALUE_CUSTOMAREAS = "InvalidParameterValue.CustomAreas"
//  INVALIDPARAMETERVALUE_ERASEPRIVACYCONFIG = "InvalidParameterValue.ErasePrivacyConfig"
//  INVALIDPARAMETERVALUE_ERASESUBTITLECONFIG = "InvalidParameterValue.EraseSubtitleConfig"
//  INVALIDPARAMETERVALUE_ERASETYPE = "InvalidParameterValue.EraseType"
//  INVALIDPARAMETERVALUE_ERASEWATERMARKCONFIG = "InvalidParameterValue.EraseWatermarkConfig"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  INVALIDPARAMETERVALUE_OCRSWITCH = "InvalidParameterValue.OcrSwitch"
//  INVALIDPARAMETERVALUE_PRIVACYMODEL = "InvalidParameterValue.PrivacyModel"
//  INVALIDPARAMETERVALUE_PRIVACYTARGETS = "InvalidParameterValue.PrivacyTargets"
//  INVALIDPARAMETERVALUE_SUBTITLEERASEMETHOD = "InvalidParameterValue.SubtitleEraseMethod"
//  INVALIDPARAMETERVALUE_SUBTITLELANG = "InvalidParameterValue.SubtitleLang"
//  INVALIDPARAMETERVALUE_SUBTITLEMODEL = "InvalidParameterValue.SubtitleModel"
//  INVALIDPARAMETERVALUE_SWITCH = "InvalidParameterValue.Switch"
//  INVALIDPARAMETERVALUE_TRANSDSTLANG = "InvalidParameterValue.TransDstLang"
//  INVALIDPARAMETERVALUE_TRANSSWITCH = "InvalidParameterValue.TransSwitch"
//  INVALIDPARAMETERVALUE_WATERMARKERASEMETHOD = "InvalidParameterValue.WatermarkEraseMethod"
//  INVALIDPARAMETERVALUE_WATERMARKMODEL = "InvalidParameterValue.WatermarkModel"
func (c *Client) CreateSmartEraseTemplate(request *CreateSmartEraseTemplateRequest) (response *CreateSmartEraseTemplateResponse, err error) {
    return c.CreateSmartEraseTemplateWithContext(context.Background(), request)
}

// CreateSmartEraseTemplate
// 创建自定义智能擦除模板
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_AUTOAREAS = "InvalidParameterValue.AutoAreas"
//  INVALIDPARAMETERVALUE_COMMENT = "InvalidParameterValue.Comment"
//  INVALIDPARAMETERVALUE_CUSTOMAREAS = "InvalidParameterValue.CustomAreas"
//  INVALIDPARAMETERVALUE_ERASEPRIVACYCONFIG = "InvalidParameterValue.ErasePrivacyConfig"
//  INVALIDPARAMETERVALUE_ERASESUBTITLECONFIG = "InvalidParameterValue.EraseSubtitleConfig"
//  INVALIDPARAMETERVALUE_ERASETYPE = "InvalidParameterValue.EraseType"
//  INVALIDPARAMETERVALUE_ERASEWATERMARKCONFIG = "InvalidParameterValue.EraseWatermarkConfig"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  INVALIDPARAMETERVALUE_OCRSWITCH = "InvalidParameterValue.OcrSwitch"
//  INVALIDPARAMETERVALUE_PRIVACYMODEL = "InvalidParameterValue.PrivacyModel"
//  INVALIDPARAMETERVALUE_PRIVACYTARGETS = "InvalidParameterValue.PrivacyTargets"
//  INVALIDPARAMETERVALUE_SUBTITLEERASEMETHOD = "InvalidParameterValue.SubtitleEraseMethod"
//  INVALIDPARAMETERVALUE_SUBTITLELANG = "InvalidParameterValue.SubtitleLang"
//  INVALIDPARAMETERVALUE_SUBTITLEMODEL = "InvalidParameterValue.SubtitleModel"
//  INVALIDPARAMETERVALUE_SWITCH = "InvalidParameterValue.Switch"
//  INVALIDPARAMETERVALUE_TRANSDSTLANG = "InvalidParameterValue.TransDstLang"
//  INVALIDPARAMETERVALUE_TRANSSWITCH = "InvalidParameterValue.TransSwitch"
//  INVALIDPARAMETERVALUE_WATERMARKERASEMETHOD = "InvalidParameterValue.WatermarkEraseMethod"
//  INVALIDPARAMETERVALUE_WATERMARKMODEL = "InvalidParameterValue.WatermarkModel"
func (c *Client) CreateSmartEraseTemplateWithContext(ctx context.Context, request *CreateSmartEraseTemplateRequest) (response *CreateSmartEraseTemplateResponse, err error) {
    if request == nil {
        request = NewCreateSmartEraseTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "CreateSmartEraseTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateSmartEraseTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateSmartEraseTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewCreateSmartSubtitleTemplateRequest() (request *CreateSmartSubtitleTemplateRequest) {
    request = &CreateSmartSubtitleTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "CreateSmartSubtitleTemplate")
    
    
    return
}

func NewCreateSmartSubtitleTemplateResponse() (response *CreateSmartSubtitleTemplateResponse) {
    response = &CreateSmartSubtitleTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateSmartSubtitleTemplate
// 创建自定义智能字幕模板
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_ASRHOTWORDSCONFIGURE = "InvalidParameterValue.AsrHotWordsConfigure"
//  INVALIDPARAMETERVALUE_ASRHOTWORDSLIBRARYID = "InvalidParameterValue.AsrHotWordsLibraryId"
//  INVALIDPARAMETERVALUE_ASRHOTWORDSSWITCH = "InvalidParameterValue.AsrHotWordsSwitch"
//  INVALIDPARAMETERVALUE_COMMENT = "InvalidParameterValue.Comment"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  INVALIDPARAMETERVALUE_SUBTITLEFORMAT = "InvalidParameterValue.SubtitleFormat"
//  INVALIDPARAMETERVALUE_SUBTITLETYPE = "InvalidParameterValue.SubtitleType"
//  INVALIDPARAMETERVALUE_TRANSLATEDSTLANGUAGE = "InvalidParameterValue.TranslateDstLanguage"
//  INVALIDPARAMETERVALUE_TRANSLATESWITCH = "InvalidParameterValue.TranslateSwitch"
//  INVALIDPARAMETERVALUE_VIDEOSRCLANGUAGE = "InvalidParameterValue.VideoSrcLanguage"
func (c *Client) CreateSmartSubtitleTemplate(request *CreateSmartSubtitleTemplateRequest) (response *CreateSmartSubtitleTemplateResponse, err error) {
    return c.CreateSmartSubtitleTemplateWithContext(context.Background(), request)
}

// CreateSmartSubtitleTemplate
// 创建自定义智能字幕模板
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_ASRHOTWORDSCONFIGURE = "InvalidParameterValue.AsrHotWordsConfigure"
//  INVALIDPARAMETERVALUE_ASRHOTWORDSLIBRARYID = "InvalidParameterValue.AsrHotWordsLibraryId"
//  INVALIDPARAMETERVALUE_ASRHOTWORDSSWITCH = "InvalidParameterValue.AsrHotWordsSwitch"
//  INVALIDPARAMETERVALUE_COMMENT = "InvalidParameterValue.Comment"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  INVALIDPARAMETERVALUE_SUBTITLEFORMAT = "InvalidParameterValue.SubtitleFormat"
//  INVALIDPARAMETERVALUE_SUBTITLETYPE = "InvalidParameterValue.SubtitleType"
//  INVALIDPARAMETERVALUE_TRANSLATEDSTLANGUAGE = "InvalidParameterValue.TranslateDstLanguage"
//  INVALIDPARAMETERVALUE_TRANSLATESWITCH = "InvalidParameterValue.TranslateSwitch"
//  INVALIDPARAMETERVALUE_VIDEOSRCLANGUAGE = "InvalidParameterValue.VideoSrcLanguage"
func (c *Client) CreateSmartSubtitleTemplateWithContext(ctx context.Context, request *CreateSmartSubtitleTemplateRequest) (response *CreateSmartSubtitleTemplateResponse, err error) {
    if request == nil {
        request = NewCreateSmartSubtitleTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "CreateSmartSubtitleTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateSmartSubtitleTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateSmartSubtitleTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewCreateSnapshotByTimeOffsetTemplateRequest() (request *CreateSnapshotByTimeOffsetTemplateRequest) {
    request = &CreateSnapshotByTimeOffsetTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "CreateSnapshotByTimeOffsetTemplate")
    
    
    return
}

func NewCreateSnapshotByTimeOffsetTemplateResponse() (response *CreateSnapshotByTimeOffsetTemplateResponse) {
    response = &CreateSnapshotByTimeOffsetTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateSnapshotByTimeOffsetTemplate
// 创建用户自定义指定时间点截图模板，数量上限：16。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_COMMENT = "InvalidParameterValue.Comment"
//  INVALIDPARAMETERVALUE_FORMAT = "InvalidParameterValue.Format"
//  INVALIDPARAMETERVALUE_HEIGHT = "InvalidParameterValue.Height"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  INVALIDPARAMETERVALUE_RESOLUTION = "InvalidParameterValue.Resolution"
//  INVALIDPARAMETERVALUE_WIDTH = "InvalidParameterValue.Width"
//  LIMITEXCEEDED_TOOMUCHTEMPLATE = "LimitExceeded.TooMuchTemplate"
func (c *Client) CreateSnapshotByTimeOffsetTemplate(request *CreateSnapshotByTimeOffsetTemplateRequest) (response *CreateSnapshotByTimeOffsetTemplateResponse, err error) {
    return c.CreateSnapshotByTimeOffsetTemplateWithContext(context.Background(), request)
}

// CreateSnapshotByTimeOffsetTemplate
// 创建用户自定义指定时间点截图模板，数量上限：16。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_COMMENT = "InvalidParameterValue.Comment"
//  INVALIDPARAMETERVALUE_FORMAT = "InvalidParameterValue.Format"
//  INVALIDPARAMETERVALUE_HEIGHT = "InvalidParameterValue.Height"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  INVALIDPARAMETERVALUE_RESOLUTION = "InvalidParameterValue.Resolution"
//  INVALIDPARAMETERVALUE_WIDTH = "InvalidParameterValue.Width"
//  LIMITEXCEEDED_TOOMUCHTEMPLATE = "LimitExceeded.TooMuchTemplate"
func (c *Client) CreateSnapshotByTimeOffsetTemplateWithContext(ctx context.Context, request *CreateSnapshotByTimeOffsetTemplateRequest) (response *CreateSnapshotByTimeOffsetTemplateResponse, err error) {
    if request == nil {
        request = NewCreateSnapshotByTimeOffsetTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "CreateSnapshotByTimeOffsetTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateSnapshotByTimeOffsetTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateSnapshotByTimeOffsetTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewCreateStreamLinkEventRequest() (request *CreateStreamLinkEventRequest) {
    request = &CreateStreamLinkEventRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "CreateStreamLinkEvent")
    
    
    return
}

func NewCreateStreamLinkEventResponse() (response *CreateStreamLinkEventResponse) {
    response = &CreateStreamLinkEventResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateStreamLinkEvent
// 创建媒体传输的事件Event。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_EXCEEDEDQUANTITYLIMIT = "InvalidParameter.ExceededQuantityLimit"
//  INVALIDPARAMETER_INPUT = "InvalidParameter.Input"
//  INVALIDPARAMETER_MAXBANDWIDTH = "InvalidParameter.MaxBandwidth"
//  INVALIDPARAMETER_NAME = "InvalidParameter.Name"
func (c *Client) CreateStreamLinkEvent(request *CreateStreamLinkEventRequest) (response *CreateStreamLinkEventResponse, err error) {
    return c.CreateStreamLinkEventWithContext(context.Background(), request)
}

// CreateStreamLinkEvent
// 创建媒体传输的事件Event。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_EXCEEDEDQUANTITYLIMIT = "InvalidParameter.ExceededQuantityLimit"
//  INVALIDPARAMETER_INPUT = "InvalidParameter.Input"
//  INVALIDPARAMETER_MAXBANDWIDTH = "InvalidParameter.MaxBandwidth"
//  INVALIDPARAMETER_NAME = "InvalidParameter.Name"
func (c *Client) CreateStreamLinkEventWithContext(ctx context.Context, request *CreateStreamLinkEventRequest) (response *CreateStreamLinkEventResponse, err error) {
    if request == nil {
        request = NewCreateStreamLinkEventRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "CreateStreamLinkEvent")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateStreamLinkEvent require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateStreamLinkEventResponse()
    err = c.Send(request, response)
    return
}

func NewCreateStreamLinkFlowRequest() (request *CreateStreamLinkFlowRequest) {
    request = &CreateStreamLinkFlowRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "CreateStreamLinkFlow")
    
    
    return
}

func NewCreateStreamLinkFlowResponse() (response *CreateStreamLinkFlowResponse) {
    response = &CreateStreamLinkFlowResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateStreamLinkFlow
// 创建媒体传输的传输流配置。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_EXCEEDEDQUANTITYLIMIT = "InvalidParameter.ExceededQuantityLimit"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_INPUT = "InvalidParameter.Input"
//  INVALIDPARAMETER_MAXBANDWIDTH = "InvalidParameter.MaxBandwidth"
//  INVALIDPARAMETER_NAME = "InvalidParameter.Name"
func (c *Client) CreateStreamLinkFlow(request *CreateStreamLinkFlowRequest) (response *CreateStreamLinkFlowResponse, err error) {
    return c.CreateStreamLinkFlowWithContext(context.Background(), request)
}

// CreateStreamLinkFlow
// 创建媒体传输的传输流配置。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_EXCEEDEDQUANTITYLIMIT = "InvalidParameter.ExceededQuantityLimit"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_INPUT = "InvalidParameter.Input"
//  INVALIDPARAMETER_MAXBANDWIDTH = "InvalidParameter.MaxBandwidth"
//  INVALIDPARAMETER_NAME = "InvalidParameter.Name"
func (c *Client) CreateStreamLinkFlowWithContext(ctx context.Context, request *CreateStreamLinkFlowRequest) (response *CreateStreamLinkFlowResponse, err error) {
    if request == nil {
        request = NewCreateStreamLinkFlowRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "CreateStreamLinkFlow")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateStreamLinkFlow require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateStreamLinkFlowResponse()
    err = c.Send(request, response)
    return
}

func NewCreateStreamLinkInputRequest() (request *CreateStreamLinkInputRequest) {
    request = &CreateStreamLinkInputRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "CreateStreamLinkInput")
    
    
    return
}

func NewCreateStreamLinkInputResponse() (response *CreateStreamLinkInputResponse) {
    response = &CreateStreamLinkInputResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateStreamLinkInput
// 创建媒体传输的输入配置。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_EXCEEDEDQUANTITYLIMIT = "InvalidParameter.ExceededQuantityLimit"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_INPUT = "InvalidParameter.Input"
//  INVALIDPARAMETER_MAXBANDWIDTH = "InvalidParameter.MaxBandwidth"
//  INVALIDPARAMETER_NAME = "InvalidParameter.Name"
//  INVALIDPARAMETER_OUTPUT = "InvalidParameter.Output"
func (c *Client) CreateStreamLinkInput(request *CreateStreamLinkInputRequest) (response *CreateStreamLinkInputResponse, err error) {
    return c.CreateStreamLinkInputWithContext(context.Background(), request)
}

// CreateStreamLinkInput
// 创建媒体传输的输入配置。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_EXCEEDEDQUANTITYLIMIT = "InvalidParameter.ExceededQuantityLimit"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_INPUT = "InvalidParameter.Input"
//  INVALIDPARAMETER_MAXBANDWIDTH = "InvalidParameter.MaxBandwidth"
//  INVALIDPARAMETER_NAME = "InvalidParameter.Name"
//  INVALIDPARAMETER_OUTPUT = "InvalidParameter.Output"
func (c *Client) CreateStreamLinkInputWithContext(ctx context.Context, request *CreateStreamLinkInputRequest) (response *CreateStreamLinkInputResponse, err error) {
    if request == nil {
        request = NewCreateStreamLinkInputRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "CreateStreamLinkInput")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateStreamLinkInput require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateStreamLinkInputResponse()
    err = c.Send(request, response)
    return
}

func NewCreateStreamLinkOutputInfoRequest() (request *CreateStreamLinkOutputInfoRequest) {
    request = &CreateStreamLinkOutputInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "CreateStreamLinkOutputInfo")
    
    
    return
}

func NewCreateStreamLinkOutputInfoResponse() (response *CreateStreamLinkOutputInfoResponse) {
    response = &CreateStreamLinkOutputInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateStreamLinkOutputInfo
// 创建媒体传输流的输出信息。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_EXCEEDEDQUANTITYLIMIT = "InvalidParameter.ExceededQuantityLimit"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_NAME = "InvalidParameter.Name"
//  INVALIDPARAMETER_OUTPUT = "InvalidParameter.Output"
//  INVALIDPARAMETER_PROTOCOL = "InvalidParameter.Protocol"
//  INVALIDPARAMETER_STATE = "InvalidParameter.State"
func (c *Client) CreateStreamLinkOutputInfo(request *CreateStreamLinkOutputInfoRequest) (response *CreateStreamLinkOutputInfoResponse, err error) {
    return c.CreateStreamLinkOutputInfoWithContext(context.Background(), request)
}

// CreateStreamLinkOutputInfo
// 创建媒体传输流的输出信息。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_EXCEEDEDQUANTITYLIMIT = "InvalidParameter.ExceededQuantityLimit"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_NAME = "InvalidParameter.Name"
//  INVALIDPARAMETER_OUTPUT = "InvalidParameter.Output"
//  INVALIDPARAMETER_PROTOCOL = "InvalidParameter.Protocol"
//  INVALIDPARAMETER_STATE = "InvalidParameter.State"
func (c *Client) CreateStreamLinkOutputInfoWithContext(ctx context.Context, request *CreateStreamLinkOutputInfoRequest) (response *CreateStreamLinkOutputInfoResponse, err error) {
    if request == nil {
        request = NewCreateStreamLinkOutputInfoRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "CreateStreamLinkOutputInfo")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateStreamLinkOutputInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateStreamLinkOutputInfoResponse()
    err = c.Send(request, response)
    return
}

func NewCreateStreamLinkSecurityGroupRequest() (request *CreateStreamLinkSecurityGroupRequest) {
    request = &CreateStreamLinkSecurityGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "CreateStreamLinkSecurityGroup")
    
    
    return
}

func NewCreateStreamLinkSecurityGroupResponse() (response *CreateStreamLinkSecurityGroupResponse) {
    response = &CreateStreamLinkSecurityGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateStreamLinkSecurityGroup
// 创建安全组，数量限制5个。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_EXCEEDEDQUANTITYLIMIT = "InvalidParameter.ExceededQuantityLimit"
//  INVALIDPARAMETER_NAME = "InvalidParameter.Name"
//  INVALIDPARAMETER_WHITELIST = "InvalidParameter.Whitelist"
func (c *Client) CreateStreamLinkSecurityGroup(request *CreateStreamLinkSecurityGroupRequest) (response *CreateStreamLinkSecurityGroupResponse, err error) {
    return c.CreateStreamLinkSecurityGroupWithContext(context.Background(), request)
}

// CreateStreamLinkSecurityGroup
// 创建安全组，数量限制5个。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_EXCEEDEDQUANTITYLIMIT = "InvalidParameter.ExceededQuantityLimit"
//  INVALIDPARAMETER_NAME = "InvalidParameter.Name"
//  INVALIDPARAMETER_WHITELIST = "InvalidParameter.Whitelist"
func (c *Client) CreateStreamLinkSecurityGroupWithContext(ctx context.Context, request *CreateStreamLinkSecurityGroupRequest) (response *CreateStreamLinkSecurityGroupResponse, err error) {
    if request == nil {
        request = NewCreateStreamLinkSecurityGroupRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "CreateStreamLinkSecurityGroup")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateStreamLinkSecurityGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateStreamLinkSecurityGroupResponse()
    err = c.Send(request, response)
    return
}

func NewCreateTranscodeTemplateRequest() (request *CreateTranscodeTemplateRequest) {
    request = &CreateTranscodeTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "CreateTranscodeTemplate")
    
    
    return
}

func NewCreateTranscodeTemplateResponse() (response *CreateTranscodeTemplateResponse) {
    response = &CreateTranscodeTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateTranscodeTemplate
// 创建用户自定义转码模板，数量上限：1000
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ACCESSDBERROR = "InternalError.AccessDBError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_AUDIOBITRATE = "InvalidParameterValue.AudioBitrate"
//  INVALIDPARAMETERVALUE_AUDIOCHANNEL = "InvalidParameterValue.AudioChannel"
//  INVALIDPARAMETERVALUE_AUDIOCODEC = "InvalidParameterValue.AudioCodec"
//  INVALIDPARAMETERVALUE_AUDIOSAMPLERATE = "InvalidParameterValue.AudioSampleRate"
//  INVALIDPARAMETERVALUE_CONTAINER = "InvalidParameterValue.Container"
//  INVALIDPARAMETERVALUE_FPS = "InvalidParameterValue.Fps"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  INVALIDPARAMETERVALUE_REMOVEAUDIO = "InvalidParameterValue.RemoveAudio"
//  INVALIDPARAMETERVALUE_REMOVEVIDEO = "InvalidParameterValue.RemoveVideo"
//  INVALIDPARAMETERVALUE_RESOLUTION = "InvalidParameterValue.Resolution"
//  INVALIDPARAMETERVALUE_SAMPLERATE = "InvalidParameterValue.SampleRate"
//  INVALIDPARAMETERVALUE_TEHDTYPE = "InvalidParameterValue.TEHDType"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  INVALIDPARAMETERVALUE_VIDEOBITRATE = "InvalidParameterValue.VideoBitrate"
//  INVALIDPARAMETERVALUE_VIDEOCODEC = "InvalidParameterValue.VideoCodec"
//  LIMITEXCEEDED_TOOMUCHTEMPLATE = "LimitExceeded.TooMuchTemplate"
func (c *Client) CreateTranscodeTemplate(request *CreateTranscodeTemplateRequest) (response *CreateTranscodeTemplateResponse, err error) {
    return c.CreateTranscodeTemplateWithContext(context.Background(), request)
}

// CreateTranscodeTemplate
// 创建用户自定义转码模板，数量上限：1000
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ACCESSDBERROR = "InternalError.AccessDBError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_AUDIOBITRATE = "InvalidParameterValue.AudioBitrate"
//  INVALIDPARAMETERVALUE_AUDIOCHANNEL = "InvalidParameterValue.AudioChannel"
//  INVALIDPARAMETERVALUE_AUDIOCODEC = "InvalidParameterValue.AudioCodec"
//  INVALIDPARAMETERVALUE_AUDIOSAMPLERATE = "InvalidParameterValue.AudioSampleRate"
//  INVALIDPARAMETERVALUE_CONTAINER = "InvalidParameterValue.Container"
//  INVALIDPARAMETERVALUE_FPS = "InvalidParameterValue.Fps"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  INVALIDPARAMETERVALUE_REMOVEAUDIO = "InvalidParameterValue.RemoveAudio"
//  INVALIDPARAMETERVALUE_REMOVEVIDEO = "InvalidParameterValue.RemoveVideo"
//  INVALIDPARAMETERVALUE_RESOLUTION = "InvalidParameterValue.Resolution"
//  INVALIDPARAMETERVALUE_SAMPLERATE = "InvalidParameterValue.SampleRate"
//  INVALIDPARAMETERVALUE_TEHDTYPE = "InvalidParameterValue.TEHDType"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  INVALIDPARAMETERVALUE_VIDEOBITRATE = "InvalidParameterValue.VideoBitrate"
//  INVALIDPARAMETERVALUE_VIDEOCODEC = "InvalidParameterValue.VideoCodec"
//  LIMITEXCEEDED_TOOMUCHTEMPLATE = "LimitExceeded.TooMuchTemplate"
func (c *Client) CreateTranscodeTemplateWithContext(ctx context.Context, request *CreateTranscodeTemplateRequest) (response *CreateTranscodeTemplateResponse, err error) {
    if request == nil {
        request = NewCreateTranscodeTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "CreateTranscodeTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateTranscodeTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateTranscodeTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewCreateVideoDatabaseEntryTaskRequest() (request *CreateVideoDatabaseEntryTaskRequest) {
    request = &CreateVideoDatabaseEntryTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "CreateVideoDatabaseEntryTask")
    
    
    return
}

func NewCreateVideoDatabaseEntryTaskResponse() (response *CreateVideoDatabaseEntryTaskResponse) {
    response = &CreateVideoDatabaseEntryTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateVideoDatabaseEntryTask
// 对URL链接或COS中的视频发起入库任务。
//
// 可选在任务完成后向回调方发送任务完成状态信息。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CreateVideoDatabaseEntryTask(request *CreateVideoDatabaseEntryTaskRequest) (response *CreateVideoDatabaseEntryTaskResponse, err error) {
    return c.CreateVideoDatabaseEntryTaskWithContext(context.Background(), request)
}

// CreateVideoDatabaseEntryTask
// 对URL链接或COS中的视频发起入库任务。
//
// 可选在任务完成后向回调方发送任务完成状态信息。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CreateVideoDatabaseEntryTaskWithContext(ctx context.Context, request *CreateVideoDatabaseEntryTaskRequest) (response *CreateVideoDatabaseEntryTaskResponse, err error) {
    if request == nil {
        request = NewCreateVideoDatabaseEntryTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "CreateVideoDatabaseEntryTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateVideoDatabaseEntryTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateVideoDatabaseEntryTaskResponse()
    err = c.Send(request, response)
    return
}

func NewCreateVideoSearchTaskRequest() (request *CreateVideoSearchTaskRequest) {
    request = &CreateVideoSearchTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "CreateVideoSearchTask")
    
    
    return
}

func NewCreateVideoSearchTaskResponse() (response *CreateVideoSearchTaskResponse) {
    response = &CreateVideoSearchTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateVideoSearchTask
// 使用检索值检索库中最接近检索值的若干视频。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CreateVideoSearchTask(request *CreateVideoSearchTaskRequest) (response *CreateVideoSearchTaskResponse, err error) {
    return c.CreateVideoSearchTaskWithContext(context.Background(), request)
}

// CreateVideoSearchTask
// 使用检索值检索库中最接近检索值的若干视频。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CreateVideoSearchTaskWithContext(ctx context.Context, request *CreateVideoSearchTaskRequest) (response *CreateVideoSearchTaskResponse, err error) {
    if request == nil {
        request = NewCreateVideoSearchTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "CreateVideoSearchTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateVideoSearchTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateVideoSearchTaskResponse()
    err = c.Send(request, response)
    return
}

func NewCreateWatermarkTemplateRequest() (request *CreateWatermarkTemplateRequest) {
    request = &CreateWatermarkTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "CreateWatermarkTemplate")
    
    
    return
}

func NewCreateWatermarkTemplateResponse() (response *CreateWatermarkTemplateResponse) {
    response = &CreateWatermarkTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateWatermarkTemplate
// 创建用户自定义水印模板，数量上限：1000。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GENDEFINITION = "InternalError.GenDefinition"
//  INTERNALERROR_UPLOADWATERMARKERROR = "InternalError.UploadWatermarkError"
//  INVALIDPARAMETERVALUE_COMMENT = "InvalidParameterValue.Comment"
//  INVALIDPARAMETERVALUE_COORDINATEORIGIN = "InvalidParameterValue.CoordinateOrigin"
//  INVALIDPARAMETERVALUE_HEIGHT = "InvalidParameterValue.Height"
//  INVALIDPARAMETERVALUE_IMAGECONTENT = "InvalidParameterValue.ImageContent"
//  INVALIDPARAMETERVALUE_IMAGETEMPLATE = "InvalidParameterValue.ImageTemplate"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  INVALIDPARAMETERVALUE_REPEATTYPE = "InvalidParameterValue.RepeatType"
//  INVALIDPARAMETERVALUE_SVGTEMPLATE = "InvalidParameterValue.SvgTemplate"
//  INVALIDPARAMETERVALUE_SVGTEMPLATEHEIGHT = "InvalidParameterValue.SvgTemplateHeight"
//  INVALIDPARAMETERVALUE_SVGTEMPLATEWIDTH = "InvalidParameterValue.SvgTemplateWidth"
//  INVALIDPARAMETERVALUE_TEXTALPHA = "InvalidParameterValue.TextAlpha"
//  INVALIDPARAMETERVALUE_TEXTTEMPLATE = "InvalidParameterValue.TextTemplate"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  INVALIDPARAMETERVALUE_WIDTH = "InvalidParameterValue.Width"
//  INVALIDPARAMETERVALUE_XPOS = "InvalidParameterValue.XPos"
//  INVALIDPARAMETERVALUE_YPOS = "InvalidParameterValue.YPos"
//  LIMITEXCEEDED_TOOMUCHTEMPLATE = "LimitExceeded.TooMuchTemplate"
func (c *Client) CreateWatermarkTemplate(request *CreateWatermarkTemplateRequest) (response *CreateWatermarkTemplateResponse, err error) {
    return c.CreateWatermarkTemplateWithContext(context.Background(), request)
}

// CreateWatermarkTemplate
// 创建用户自定义水印模板，数量上限：1000。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GENDEFINITION = "InternalError.GenDefinition"
//  INTERNALERROR_UPLOADWATERMARKERROR = "InternalError.UploadWatermarkError"
//  INVALIDPARAMETERVALUE_COMMENT = "InvalidParameterValue.Comment"
//  INVALIDPARAMETERVALUE_COORDINATEORIGIN = "InvalidParameterValue.CoordinateOrigin"
//  INVALIDPARAMETERVALUE_HEIGHT = "InvalidParameterValue.Height"
//  INVALIDPARAMETERVALUE_IMAGECONTENT = "InvalidParameterValue.ImageContent"
//  INVALIDPARAMETERVALUE_IMAGETEMPLATE = "InvalidParameterValue.ImageTemplate"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  INVALIDPARAMETERVALUE_REPEATTYPE = "InvalidParameterValue.RepeatType"
//  INVALIDPARAMETERVALUE_SVGTEMPLATE = "InvalidParameterValue.SvgTemplate"
//  INVALIDPARAMETERVALUE_SVGTEMPLATEHEIGHT = "InvalidParameterValue.SvgTemplateHeight"
//  INVALIDPARAMETERVALUE_SVGTEMPLATEWIDTH = "InvalidParameterValue.SvgTemplateWidth"
//  INVALIDPARAMETERVALUE_TEXTALPHA = "InvalidParameterValue.TextAlpha"
//  INVALIDPARAMETERVALUE_TEXTTEMPLATE = "InvalidParameterValue.TextTemplate"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  INVALIDPARAMETERVALUE_WIDTH = "InvalidParameterValue.Width"
//  INVALIDPARAMETERVALUE_XPOS = "InvalidParameterValue.XPos"
//  INVALIDPARAMETERVALUE_YPOS = "InvalidParameterValue.YPos"
//  LIMITEXCEEDED_TOOMUCHTEMPLATE = "LimitExceeded.TooMuchTemplate"
func (c *Client) CreateWatermarkTemplateWithContext(ctx context.Context, request *CreateWatermarkTemplateRequest) (response *CreateWatermarkTemplateResponse, err error) {
    if request == nil {
        request = NewCreateWatermarkTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "CreateWatermarkTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateWatermarkTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateWatermarkTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewCreateWordSamplesRequest() (request *CreateWordSamplesRequest) {
    request = &CreateWordSamplesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "CreateWordSamples")
    
    
    return
}

func NewCreateWordSamplesResponse() (response *CreateWordSamplesResponse) {
    response = &CreateWordSamplesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateWordSamples
// 该接口用于批量创建关键词样本，样本用于通过OCR、ASR技术，进行不适宜内容识别、内容识别等视频处理。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) CreateWordSamples(request *CreateWordSamplesRequest) (response *CreateWordSamplesResponse, err error) {
    return c.CreateWordSamplesWithContext(context.Background(), request)
}

// CreateWordSamples
// 该接口用于批量创建关键词样本，样本用于通过OCR、ASR技术，进行不适宜内容识别、内容识别等视频处理。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) CreateWordSamplesWithContext(ctx context.Context, request *CreateWordSamplesRequest) (response *CreateWordSamplesResponse, err error) {
    if request == nil {
        request = NewCreateWordSamplesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "CreateWordSamples")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateWordSamples require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateWordSamplesResponse()
    err = c.Send(request, response)
    return
}

func NewCreateWorkflowRequest() (request *CreateWorkflowRequest) {
    request = &CreateWorkflowRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "CreateWorkflow")
    
    
    return
}

func NewCreateWorkflowResponse() (response *CreateWorkflowResponse) {
    response = &CreateWorkflowResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateWorkflow
// 对 COS 中指定 Bucket 的目录下上传的媒体文件，设置处理规则，包括：
//
// 1. 视频转码（带水印）；
//
// 2. 视频转动图；
//
// 3. 对视频按指定时间点截图；
//
// 4. 对视频采样截图；
//
// 5. 对视频截图雪碧图；
//
// 6. 对视频转自适应码流；
//
// 7. 智能内容审核（鉴黄、敏感信息检测）；
//
// 8. 智能内容分析（标签、分类、封面、按帧标签）；
//
// 9. 智能内容识别（人脸、文本全文、文本关键词、语音全文、语音关键词）。
//
// 
//
// 注意：创建工作流成功后是禁用状态，需要手动启用。
//
// 可能返回的错误码:
//  FAILEDOPERATION_COSSTATUSINAVLID = "FailedOperation.CosStatusInavlid"
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_TOOMUCHTEMPLATE = "LimitExceeded.TooMuchTemplate"
//  RESOURCENOTFOUND_COSBUCKETNAMEINVALID = "ResourceNotFound.CosBucketNameInvalid"
//  RESOURCENOTFOUND_COSBUCKETNOTEXIST = "ResourceNotFound.CosBucketNotExist"
func (c *Client) CreateWorkflow(request *CreateWorkflowRequest) (response *CreateWorkflowResponse, err error) {
    return c.CreateWorkflowWithContext(context.Background(), request)
}

// CreateWorkflow
// 对 COS 中指定 Bucket 的目录下上传的媒体文件，设置处理规则，包括：
//
// 1. 视频转码（带水印）；
//
// 2. 视频转动图；
//
// 3. 对视频按指定时间点截图；
//
// 4. 对视频采样截图；
//
// 5. 对视频截图雪碧图；
//
// 6. 对视频转自适应码流；
//
// 7. 智能内容审核（鉴黄、敏感信息检测）；
//
// 8. 智能内容分析（标签、分类、封面、按帧标签）；
//
// 9. 智能内容识别（人脸、文本全文、文本关键词、语音全文、语音关键词）。
//
// 
//
// 注意：创建工作流成功后是禁用状态，需要手动启用。
//
// 可能返回的错误码:
//  FAILEDOPERATION_COSSTATUSINAVLID = "FailedOperation.CosStatusInavlid"
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_TOOMUCHTEMPLATE = "LimitExceeded.TooMuchTemplate"
//  RESOURCENOTFOUND_COSBUCKETNAMEINVALID = "ResourceNotFound.CosBucketNameInvalid"
//  RESOURCENOTFOUND_COSBUCKETNOTEXIST = "ResourceNotFound.CosBucketNotExist"
func (c *Client) CreateWorkflowWithContext(ctx context.Context, request *CreateWorkflowRequest) (response *CreateWorkflowResponse, err error) {
    if request == nil {
        request = NewCreateWorkflowRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "CreateWorkflow")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateWorkflow require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateWorkflowResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAIAnalysisTemplateRequest() (request *DeleteAIAnalysisTemplateRequest) {
    request = &DeleteAIAnalysisTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "DeleteAIAnalysisTemplate")
    
    
    return
}

func NewDeleteAIAnalysisTemplateResponse() (response *DeleteAIAnalysisTemplateResponse) {
    response = &DeleteAIAnalysisTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteAIAnalysisTemplate
// 删除用户自定义内容分析模板。
//
// 
//
// 注意：模板 ID 为 10000 以下的为系统预置模板，不允许删除。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_DELETEDEFAULTTEMPLATE = "InvalidParameterValue.DeleteDefaultTemplate"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) DeleteAIAnalysisTemplate(request *DeleteAIAnalysisTemplateRequest) (response *DeleteAIAnalysisTemplateResponse, err error) {
    return c.DeleteAIAnalysisTemplateWithContext(context.Background(), request)
}

// DeleteAIAnalysisTemplate
// 删除用户自定义内容分析模板。
//
// 
//
// 注意：模板 ID 为 10000 以下的为系统预置模板，不允许删除。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_DELETEDEFAULTTEMPLATE = "InvalidParameterValue.DeleteDefaultTemplate"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) DeleteAIAnalysisTemplateWithContext(ctx context.Context, request *DeleteAIAnalysisTemplateRequest) (response *DeleteAIAnalysisTemplateResponse, err error) {
    if request == nil {
        request = NewDeleteAIAnalysisTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "DeleteAIAnalysisTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteAIAnalysisTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteAIAnalysisTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAIRecognitionTemplateRequest() (request *DeleteAIRecognitionTemplateRequest) {
    request = &DeleteAIRecognitionTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "DeleteAIRecognitionTemplate")
    
    
    return
}

func NewDeleteAIRecognitionTemplateResponse() (response *DeleteAIRecognitionTemplateResponse) {
    response = &DeleteAIRecognitionTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteAIRecognitionTemplate
// 删除用户自定义内容识别模板。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_DELETEDEFAULTTEMPLATE = "InvalidParameterValue.DeleteDefaultTemplate"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) DeleteAIRecognitionTemplate(request *DeleteAIRecognitionTemplateRequest) (response *DeleteAIRecognitionTemplateResponse, err error) {
    return c.DeleteAIRecognitionTemplateWithContext(context.Background(), request)
}

// DeleteAIRecognitionTemplate
// 删除用户自定义内容识别模板。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_DELETEDEFAULTTEMPLATE = "InvalidParameterValue.DeleteDefaultTemplate"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) DeleteAIRecognitionTemplateWithContext(ctx context.Context, request *DeleteAIRecognitionTemplateRequest) (response *DeleteAIRecognitionTemplateResponse, err error) {
    if request == nil {
        request = NewDeleteAIRecognitionTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "DeleteAIRecognitionTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteAIRecognitionTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteAIRecognitionTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAdaptiveDynamicStreamingTemplateRequest() (request *DeleteAdaptiveDynamicStreamingTemplateRequest) {
    request = &DeleteAdaptiveDynamicStreamingTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "DeleteAdaptiveDynamicStreamingTemplate")
    
    
    return
}

func NewDeleteAdaptiveDynamicStreamingTemplateResponse() (response *DeleteAdaptiveDynamicStreamingTemplateResponse) {
    response = &DeleteAdaptiveDynamicStreamingTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteAdaptiveDynamicStreamingTemplate
// 删除转自适应码流模板
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) DeleteAdaptiveDynamicStreamingTemplate(request *DeleteAdaptiveDynamicStreamingTemplateRequest) (response *DeleteAdaptiveDynamicStreamingTemplateResponse, err error) {
    return c.DeleteAdaptiveDynamicStreamingTemplateWithContext(context.Background(), request)
}

// DeleteAdaptiveDynamicStreamingTemplate
// 删除转自适应码流模板
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) DeleteAdaptiveDynamicStreamingTemplateWithContext(ctx context.Context, request *DeleteAdaptiveDynamicStreamingTemplateRequest) (response *DeleteAdaptiveDynamicStreamingTemplateResponse, err error) {
    if request == nil {
        request = NewDeleteAdaptiveDynamicStreamingTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "DeleteAdaptiveDynamicStreamingTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteAdaptiveDynamicStreamingTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteAdaptiveDynamicStreamingTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAnimatedGraphicsTemplateRequest() (request *DeleteAnimatedGraphicsTemplateRequest) {
    request = &DeleteAnimatedGraphicsTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "DeleteAnimatedGraphicsTemplate")
    
    
    return
}

func NewDeleteAnimatedGraphicsTemplateResponse() (response *DeleteAnimatedGraphicsTemplateResponse) {
    response = &DeleteAnimatedGraphicsTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteAnimatedGraphicsTemplate
// 删除用户自定义转动图模板。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) DeleteAnimatedGraphicsTemplate(request *DeleteAnimatedGraphicsTemplateRequest) (response *DeleteAnimatedGraphicsTemplateResponse, err error) {
    return c.DeleteAnimatedGraphicsTemplateWithContext(context.Background(), request)
}

// DeleteAnimatedGraphicsTemplate
// 删除用户自定义转动图模板。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) DeleteAnimatedGraphicsTemplateWithContext(ctx context.Context, request *DeleteAnimatedGraphicsTemplateRequest) (response *DeleteAnimatedGraphicsTemplateResponse, err error) {
    if request == nil {
        request = NewDeleteAnimatedGraphicsTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "DeleteAnimatedGraphicsTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteAnimatedGraphicsTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteAnimatedGraphicsTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAsrHotwordsRequest() (request *DeleteAsrHotwordsRequest) {
    request = &DeleteAsrHotwordsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "DeleteAsrHotwords")
    
    
    return
}

func NewDeleteAsrHotwordsResponse() (response *DeleteAsrHotwordsResponse) {
    response = &DeleteAsrHotwordsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteAsrHotwords
// 删除智能字幕热词库
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ACCESSDBERROR = "InternalError.AccessDBError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_HOTWORDSNOTEXIST = "InvalidParameterValue.HotWordsNotExist"
func (c *Client) DeleteAsrHotwords(request *DeleteAsrHotwordsRequest) (response *DeleteAsrHotwordsResponse, err error) {
    return c.DeleteAsrHotwordsWithContext(context.Background(), request)
}

// DeleteAsrHotwords
// 删除智能字幕热词库
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ACCESSDBERROR = "InternalError.AccessDBError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_HOTWORDSNOTEXIST = "InvalidParameterValue.HotWordsNotExist"
func (c *Client) DeleteAsrHotwordsWithContext(ctx context.Context, request *DeleteAsrHotwordsRequest) (response *DeleteAsrHotwordsResponse, err error) {
    if request == nil {
        request = NewDeleteAsrHotwordsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "DeleteAsrHotwords")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteAsrHotwords require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteAsrHotwordsResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteBlindWatermarkTemplateRequest() (request *DeleteBlindWatermarkTemplateRequest) {
    request = &DeleteBlindWatermarkTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "DeleteBlindWatermarkTemplate")
    
    
    return
}

func NewDeleteBlindWatermarkTemplateResponse() (response *DeleteBlindWatermarkTemplateResponse) {
    response = &DeleteBlindWatermarkTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteBlindWatermarkTemplate
// 删除用户自定义数字水印模板。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_DEFINITION = "InvalidParameterValue.Definition"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteBlindWatermarkTemplate(request *DeleteBlindWatermarkTemplateRequest) (response *DeleteBlindWatermarkTemplateResponse, err error) {
    return c.DeleteBlindWatermarkTemplateWithContext(context.Background(), request)
}

// DeleteBlindWatermarkTemplate
// 删除用户自定义数字水印模板。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_DEFINITION = "InvalidParameterValue.Definition"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteBlindWatermarkTemplateWithContext(ctx context.Context, request *DeleteBlindWatermarkTemplateRequest) (response *DeleteBlindWatermarkTemplateResponse, err error) {
    if request == nil {
        request = NewDeleteBlindWatermarkTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "DeleteBlindWatermarkTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteBlindWatermarkTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteBlindWatermarkTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteContentReviewTemplateRequest() (request *DeleteContentReviewTemplateRequest) {
    request = &DeleteContentReviewTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "DeleteContentReviewTemplate")
    
    
    return
}

func NewDeleteContentReviewTemplateResponse() (response *DeleteContentReviewTemplateResponse) {
    response = &DeleteContentReviewTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteContentReviewTemplate
// 删除用户自定义内容审核模板。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_DELETEDEFAULTTEMPLATE = "InvalidParameterValue.DeleteDefaultTemplate"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) DeleteContentReviewTemplate(request *DeleteContentReviewTemplateRequest) (response *DeleteContentReviewTemplateResponse, err error) {
    return c.DeleteContentReviewTemplateWithContext(context.Background(), request)
}

// DeleteContentReviewTemplate
// 删除用户自定义内容审核模板。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_DELETEDEFAULTTEMPLATE = "InvalidParameterValue.DeleteDefaultTemplate"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) DeleteContentReviewTemplateWithContext(ctx context.Context, request *DeleteContentReviewTemplateRequest) (response *DeleteContentReviewTemplateResponse, err error) {
    if request == nil {
        request = NewDeleteContentReviewTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "DeleteContentReviewTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteContentReviewTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteContentReviewTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteImageSpriteTemplateRequest() (request *DeleteImageSpriteTemplateRequest) {
    request = &DeleteImageSpriteTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "DeleteImageSpriteTemplate")
    
    
    return
}

func NewDeleteImageSpriteTemplateResponse() (response *DeleteImageSpriteTemplateResponse) {
    response = &DeleteImageSpriteTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteImageSpriteTemplate
// 删除雪碧图模板。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) DeleteImageSpriteTemplate(request *DeleteImageSpriteTemplateRequest) (response *DeleteImageSpriteTemplateResponse, err error) {
    return c.DeleteImageSpriteTemplateWithContext(context.Background(), request)
}

// DeleteImageSpriteTemplate
// 删除雪碧图模板。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) DeleteImageSpriteTemplateWithContext(ctx context.Context, request *DeleteImageSpriteTemplateRequest) (response *DeleteImageSpriteTemplateResponse, err error) {
    if request == nil {
        request = NewDeleteImageSpriteTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "DeleteImageSpriteTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteImageSpriteTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteImageSpriteTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteLiveRecordTemplateRequest() (request *DeleteLiveRecordTemplateRequest) {
    request = &DeleteLiveRecordTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "DeleteLiveRecordTemplate")
    
    
    return
}

func NewDeleteLiveRecordTemplateResponse() (response *DeleteLiveRecordTemplateResponse) {
    response = &DeleteLiveRecordTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteLiveRecordTemplate
// 删除直播录制模板
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ACCESSDBERROR = "InternalError.AccessDBError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) DeleteLiveRecordTemplate(request *DeleteLiveRecordTemplateRequest) (response *DeleteLiveRecordTemplateResponse, err error) {
    return c.DeleteLiveRecordTemplateWithContext(context.Background(), request)
}

// DeleteLiveRecordTemplate
// 删除直播录制模板
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ACCESSDBERROR = "InternalError.AccessDBError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) DeleteLiveRecordTemplateWithContext(ctx context.Context, request *DeleteLiveRecordTemplateRequest) (response *DeleteLiveRecordTemplateResponse, err error) {
    if request == nil {
        request = NewDeleteLiveRecordTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "DeleteLiveRecordTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteLiveRecordTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteLiveRecordTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewDeletePersonSampleRequest() (request *DeletePersonSampleRequest) {
    request = &DeletePersonSampleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "DeletePersonSample")
    
    
    return
}

func NewDeletePersonSampleResponse() (response *DeletePersonSampleResponse) {
    response = &DeletePersonSampleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeletePersonSample
// 该接口用于根据素材 ID，删除素材样本。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND_PERSON = "ResourceNotFound.Person"
func (c *Client) DeletePersonSample(request *DeletePersonSampleRequest) (response *DeletePersonSampleResponse, err error) {
    return c.DeletePersonSampleWithContext(context.Background(), request)
}

// DeletePersonSample
// 该接口用于根据素材 ID，删除素材样本。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND_PERSON = "ResourceNotFound.Person"
func (c *Client) DeletePersonSampleWithContext(ctx context.Context, request *DeletePersonSampleRequest) (response *DeletePersonSampleResponse, err error) {
    if request == nil {
        request = NewDeletePersonSampleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "DeletePersonSample")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeletePersonSample require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeletePersonSampleResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteProcessImageTemplateRequest() (request *DeleteProcessImageTemplateRequest) {
    request = &DeleteProcessImageTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "DeleteProcessImageTemplate")
    
    
    return
}

func NewDeleteProcessImageTemplateResponse() (response *DeleteProcessImageTemplateResponse) {
    response = &DeleteProcessImageTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteProcessImageTemplate
// 删除图片处理模板
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) DeleteProcessImageTemplate(request *DeleteProcessImageTemplateRequest) (response *DeleteProcessImageTemplateResponse, err error) {
    return c.DeleteProcessImageTemplateWithContext(context.Background(), request)
}

// DeleteProcessImageTemplate
// 删除图片处理模板
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) DeleteProcessImageTemplateWithContext(ctx context.Context, request *DeleteProcessImageTemplateRequest) (response *DeleteProcessImageTemplateResponse, err error) {
    if request == nil {
        request = NewDeleteProcessImageTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "DeleteProcessImageTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteProcessImageTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteProcessImageTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteQualityControlTemplateRequest() (request *DeleteQualityControlTemplateRequest) {
    request = &DeleteQualityControlTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "DeleteQualityControlTemplate")
    
    
    return
}

func NewDeleteQualityControlTemplateResponse() (response *DeleteQualityControlTemplateResponse) {
    response = &DeleteQualityControlTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteQualityControlTemplate
// 删除媒体质检模板
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) DeleteQualityControlTemplate(request *DeleteQualityControlTemplateRequest) (response *DeleteQualityControlTemplateResponse, err error) {
    return c.DeleteQualityControlTemplateWithContext(context.Background(), request)
}

// DeleteQualityControlTemplate
// 删除媒体质检模板
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) DeleteQualityControlTemplateWithContext(ctx context.Context, request *DeleteQualityControlTemplateRequest) (response *DeleteQualityControlTemplateResponse, err error) {
    if request == nil {
        request = NewDeleteQualityControlTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "DeleteQualityControlTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteQualityControlTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteQualityControlTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteSampleSnapshotTemplateRequest() (request *DeleteSampleSnapshotTemplateRequest) {
    request = &DeleteSampleSnapshotTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "DeleteSampleSnapshotTemplate")
    
    
    return
}

func NewDeleteSampleSnapshotTemplateResponse() (response *DeleteSampleSnapshotTemplateResponse) {
    response = &DeleteSampleSnapshotTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteSampleSnapshotTemplate
// 删除用户自定义采样截图模板。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) DeleteSampleSnapshotTemplate(request *DeleteSampleSnapshotTemplateRequest) (response *DeleteSampleSnapshotTemplateResponse, err error) {
    return c.DeleteSampleSnapshotTemplateWithContext(context.Background(), request)
}

// DeleteSampleSnapshotTemplate
// 删除用户自定义采样截图模板。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) DeleteSampleSnapshotTemplateWithContext(ctx context.Context, request *DeleteSampleSnapshotTemplateRequest) (response *DeleteSampleSnapshotTemplateResponse, err error) {
    if request == nil {
        request = NewDeleteSampleSnapshotTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "DeleteSampleSnapshotTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteSampleSnapshotTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteSampleSnapshotTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteScheduleRequest() (request *DeleteScheduleRequest) {
    request = &DeleteScheduleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "DeleteSchedule")
    
    
    return
}

func NewDeleteScheduleResponse() (response *DeleteScheduleResponse) {
    response = &DeleteScheduleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteSchedule
// 删除编排
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ACCESSDBERROR = "InternalError.AccessDBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) DeleteSchedule(request *DeleteScheduleRequest) (response *DeleteScheduleResponse, err error) {
    return c.DeleteScheduleWithContext(context.Background(), request)
}

// DeleteSchedule
// 删除编排
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ACCESSDBERROR = "InternalError.AccessDBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) DeleteScheduleWithContext(ctx context.Context, request *DeleteScheduleRequest) (response *DeleteScheduleResponse, err error) {
    if request == nil {
        request = NewDeleteScheduleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "DeleteSchedule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteSchedule require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteScheduleResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteSmartEraseTemplateRequest() (request *DeleteSmartEraseTemplateRequest) {
    request = &DeleteSmartEraseTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "DeleteSmartEraseTemplate")
    
    
    return
}

func NewDeleteSmartEraseTemplateResponse() (response *DeleteSmartEraseTemplateResponse) {
    response = &DeleteSmartEraseTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteSmartEraseTemplate
// 删除用户自定义智能擦除模板。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_DELETEDEFAULTTEMPLATE = "InvalidParameterValue.DeleteDefaultTemplate"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) DeleteSmartEraseTemplate(request *DeleteSmartEraseTemplateRequest) (response *DeleteSmartEraseTemplateResponse, err error) {
    return c.DeleteSmartEraseTemplateWithContext(context.Background(), request)
}

// DeleteSmartEraseTemplate
// 删除用户自定义智能擦除模板。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_DELETEDEFAULTTEMPLATE = "InvalidParameterValue.DeleteDefaultTemplate"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) DeleteSmartEraseTemplateWithContext(ctx context.Context, request *DeleteSmartEraseTemplateRequest) (response *DeleteSmartEraseTemplateResponse, err error) {
    if request == nil {
        request = NewDeleteSmartEraseTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "DeleteSmartEraseTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteSmartEraseTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteSmartEraseTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteSmartSubtitleTemplateRequest() (request *DeleteSmartSubtitleTemplateRequest) {
    request = &DeleteSmartSubtitleTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "DeleteSmartSubtitleTemplate")
    
    
    return
}

func NewDeleteSmartSubtitleTemplateResponse() (response *DeleteSmartSubtitleTemplateResponse) {
    response = &DeleteSmartSubtitleTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteSmartSubtitleTemplate
// 删除用户自定义智能字幕模板。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_DELETEDEFAULTTEMPLATE = "InvalidParameterValue.DeleteDefaultTemplate"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) DeleteSmartSubtitleTemplate(request *DeleteSmartSubtitleTemplateRequest) (response *DeleteSmartSubtitleTemplateResponse, err error) {
    return c.DeleteSmartSubtitleTemplateWithContext(context.Background(), request)
}

// DeleteSmartSubtitleTemplate
// 删除用户自定义智能字幕模板。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_DELETEDEFAULTTEMPLATE = "InvalidParameterValue.DeleteDefaultTemplate"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) DeleteSmartSubtitleTemplateWithContext(ctx context.Context, request *DeleteSmartSubtitleTemplateRequest) (response *DeleteSmartSubtitleTemplateResponse, err error) {
    if request == nil {
        request = NewDeleteSmartSubtitleTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "DeleteSmartSubtitleTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteSmartSubtitleTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteSmartSubtitleTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteSnapshotByTimeOffsetTemplateRequest() (request *DeleteSnapshotByTimeOffsetTemplateRequest) {
    request = &DeleteSnapshotByTimeOffsetTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "DeleteSnapshotByTimeOffsetTemplate")
    
    
    return
}

func NewDeleteSnapshotByTimeOffsetTemplateResponse() (response *DeleteSnapshotByTimeOffsetTemplateResponse) {
    response = &DeleteSnapshotByTimeOffsetTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteSnapshotByTimeOffsetTemplate
// 删除用户自定义指定时间点截图模板。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) DeleteSnapshotByTimeOffsetTemplate(request *DeleteSnapshotByTimeOffsetTemplateRequest) (response *DeleteSnapshotByTimeOffsetTemplateResponse, err error) {
    return c.DeleteSnapshotByTimeOffsetTemplateWithContext(context.Background(), request)
}

// DeleteSnapshotByTimeOffsetTemplate
// 删除用户自定义指定时间点截图模板。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) DeleteSnapshotByTimeOffsetTemplateWithContext(ctx context.Context, request *DeleteSnapshotByTimeOffsetTemplateRequest) (response *DeleteSnapshotByTimeOffsetTemplateResponse, err error) {
    if request == nil {
        request = NewDeleteSnapshotByTimeOffsetTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "DeleteSnapshotByTimeOffsetTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteSnapshotByTimeOffsetTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteSnapshotByTimeOffsetTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteStreamLinkEventRequest() (request *DeleteStreamLinkEventRequest) {
    request = &DeleteStreamLinkEventRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "DeleteStreamLinkEvent")
    
    
    return
}

func NewDeleteStreamLinkEventResponse() (response *DeleteStreamLinkEventResponse) {
    response = &DeleteStreamLinkEventResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteStreamLinkEvent
// 删除媒体传输的事件配置。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_STATE = "InvalidParameter.State"
func (c *Client) DeleteStreamLinkEvent(request *DeleteStreamLinkEventRequest) (response *DeleteStreamLinkEventResponse, err error) {
    return c.DeleteStreamLinkEventWithContext(context.Background(), request)
}

// DeleteStreamLinkEvent
// 删除媒体传输的事件配置。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_STATE = "InvalidParameter.State"
func (c *Client) DeleteStreamLinkEventWithContext(ctx context.Context, request *DeleteStreamLinkEventRequest) (response *DeleteStreamLinkEventResponse, err error) {
    if request == nil {
        request = NewDeleteStreamLinkEventRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "DeleteStreamLinkEvent")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteStreamLinkEvent require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteStreamLinkEventResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteStreamLinkFlowRequest() (request *DeleteStreamLinkFlowRequest) {
    request = &DeleteStreamLinkFlowRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "DeleteStreamLinkFlow")
    
    
    return
}

func NewDeleteStreamLinkFlowResponse() (response *DeleteStreamLinkFlowResponse) {
    response = &DeleteStreamLinkFlowResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteStreamLinkFlow
// 删除媒体传输的传输流配置。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_STATE = "InvalidParameter.State"
func (c *Client) DeleteStreamLinkFlow(request *DeleteStreamLinkFlowRequest) (response *DeleteStreamLinkFlowResponse, err error) {
    return c.DeleteStreamLinkFlowWithContext(context.Background(), request)
}

// DeleteStreamLinkFlow
// 删除媒体传输的传输流配置。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_STATE = "InvalidParameter.State"
func (c *Client) DeleteStreamLinkFlowWithContext(ctx context.Context, request *DeleteStreamLinkFlowRequest) (response *DeleteStreamLinkFlowResponse, err error) {
    if request == nil {
        request = NewDeleteStreamLinkFlowRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "DeleteStreamLinkFlow")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteStreamLinkFlow require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteStreamLinkFlowResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteStreamLinkOutputRequest() (request *DeleteStreamLinkOutputRequest) {
    request = &DeleteStreamLinkOutputRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "DeleteStreamLinkOutput")
    
    
    return
}

func NewDeleteStreamLinkOutputResponse() (response *DeleteStreamLinkOutputResponse) {
    response = &DeleteStreamLinkOutputResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteStreamLinkOutput
// 删除媒体传输流的输出配置。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_NOTFOUND = "InvalidParameter.NotFound"
//  INVALIDPARAMETER_OUTPUTID = "InvalidParameter.OutputId"
//  INVALIDPARAMETER_STATE = "InvalidParameter.State"
func (c *Client) DeleteStreamLinkOutput(request *DeleteStreamLinkOutputRequest) (response *DeleteStreamLinkOutputResponse, err error) {
    return c.DeleteStreamLinkOutputWithContext(context.Background(), request)
}

// DeleteStreamLinkOutput
// 删除媒体传输流的输出配置。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_NOTFOUND = "InvalidParameter.NotFound"
//  INVALIDPARAMETER_OUTPUTID = "InvalidParameter.OutputId"
//  INVALIDPARAMETER_STATE = "InvalidParameter.State"
func (c *Client) DeleteStreamLinkOutputWithContext(ctx context.Context, request *DeleteStreamLinkOutputRequest) (response *DeleteStreamLinkOutputResponse, err error) {
    if request == nil {
        request = NewDeleteStreamLinkOutputRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "DeleteStreamLinkOutput")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteStreamLinkOutput require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteStreamLinkOutputResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteStreamLinkSecurityGroupRequest() (request *DeleteStreamLinkSecurityGroupRequest) {
    request = &DeleteStreamLinkSecurityGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "DeleteStreamLinkSecurityGroup")
    
    
    return
}

func NewDeleteStreamLinkSecurityGroupResponse() (response *DeleteStreamLinkSecurityGroupResponse) {
    response = &DeleteStreamLinkSecurityGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteStreamLinkSecurityGroup
// 删除安全组。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ALREADYASSOCIATEDINPUT = "InvalidParameter.AlreadyAssociatedInput"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_NOTFOUND = "InvalidParameter.NotFound"
func (c *Client) DeleteStreamLinkSecurityGroup(request *DeleteStreamLinkSecurityGroupRequest) (response *DeleteStreamLinkSecurityGroupResponse, err error) {
    return c.DeleteStreamLinkSecurityGroupWithContext(context.Background(), request)
}

// DeleteStreamLinkSecurityGroup
// 删除安全组。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ALREADYASSOCIATEDINPUT = "InvalidParameter.AlreadyAssociatedInput"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_NOTFOUND = "InvalidParameter.NotFound"
func (c *Client) DeleteStreamLinkSecurityGroupWithContext(ctx context.Context, request *DeleteStreamLinkSecurityGroupRequest) (response *DeleteStreamLinkSecurityGroupResponse, err error) {
    if request == nil {
        request = NewDeleteStreamLinkSecurityGroupRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "DeleteStreamLinkSecurityGroup")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteStreamLinkSecurityGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteStreamLinkSecurityGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteTranscodeTemplateRequest() (request *DeleteTranscodeTemplateRequest) {
    request = &DeleteTranscodeTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "DeleteTranscodeTemplate")
    
    
    return
}

func NewDeleteTranscodeTemplateResponse() (response *DeleteTranscodeTemplateResponse) {
    response = &DeleteTranscodeTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteTranscodeTemplate
// 删除用户自定义转码模板。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ACCESSDBERROR = "InternalError.AccessDBError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DEFINITION = "InvalidParameterValue.Definition"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) DeleteTranscodeTemplate(request *DeleteTranscodeTemplateRequest) (response *DeleteTranscodeTemplateResponse, err error) {
    return c.DeleteTranscodeTemplateWithContext(context.Background(), request)
}

// DeleteTranscodeTemplate
// 删除用户自定义转码模板。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ACCESSDBERROR = "InternalError.AccessDBError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DEFINITION = "InvalidParameterValue.Definition"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) DeleteTranscodeTemplateWithContext(ctx context.Context, request *DeleteTranscodeTemplateRequest) (response *DeleteTranscodeTemplateResponse, err error) {
    if request == nil {
        request = NewDeleteTranscodeTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "DeleteTranscodeTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteTranscodeTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteTranscodeTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteWatermarkTemplateRequest() (request *DeleteWatermarkTemplateRequest) {
    request = &DeleteWatermarkTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "DeleteWatermarkTemplate")
    
    
    return
}

func NewDeleteWatermarkTemplateResponse() (response *DeleteWatermarkTemplateResponse) {
    response = &DeleteWatermarkTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteWatermarkTemplate
// 删除用户自定义水印模板。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_DEFINITION = "InvalidParameterValue.Definition"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteWatermarkTemplate(request *DeleteWatermarkTemplateRequest) (response *DeleteWatermarkTemplateResponse, err error) {
    return c.DeleteWatermarkTemplateWithContext(context.Background(), request)
}

// DeleteWatermarkTemplate
// 删除用户自定义水印模板。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_DEFINITION = "InvalidParameterValue.Definition"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteWatermarkTemplateWithContext(ctx context.Context, request *DeleteWatermarkTemplateRequest) (response *DeleteWatermarkTemplateResponse, err error) {
    if request == nil {
        request = NewDeleteWatermarkTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "DeleteWatermarkTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteWatermarkTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteWatermarkTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteWordSamplesRequest() (request *DeleteWordSamplesRequest) {
    request = &DeleteWordSamplesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "DeleteWordSamples")
    
    
    return
}

func NewDeleteWordSamplesResponse() (response *DeleteWordSamplesResponse) {
    response = &DeleteWordSamplesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteWordSamples
// 该接口用于批量删除关键词样本。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DeleteWordSamples(request *DeleteWordSamplesRequest) (response *DeleteWordSamplesResponse, err error) {
    return c.DeleteWordSamplesWithContext(context.Background(), request)
}

// DeleteWordSamples
// 该接口用于批量删除关键词样本。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DeleteWordSamplesWithContext(ctx context.Context, request *DeleteWordSamplesRequest) (response *DeleteWordSamplesResponse, err error) {
    if request == nil {
        request = NewDeleteWordSamplesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "DeleteWordSamples")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteWordSamples require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteWordSamplesResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteWorkflowRequest() (request *DeleteWorkflowRequest) {
    request = &DeleteWorkflowRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "DeleteWorkflow")
    
    
    return
}

func NewDeleteWorkflowResponse() (response *DeleteWorkflowResponse) {
    response = &DeleteWorkflowResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteWorkflow
// 删除工作流。对于已启用的工作流，需要禁用后才能删除。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) DeleteWorkflow(request *DeleteWorkflowRequest) (response *DeleteWorkflowResponse, err error) {
    return c.DeleteWorkflowWithContext(context.Background(), request)
}

// DeleteWorkflow
// 删除工作流。对于已启用的工作流，需要禁用后才能删除。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) DeleteWorkflowWithContext(ctx context.Context, request *DeleteWorkflowRequest) (response *DeleteWorkflowResponse, err error) {
    if request == nil {
        request = NewDeleteWorkflowRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "DeleteWorkflow")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteWorkflow require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteWorkflowResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAIAnalysisTemplatesRequest() (request *DescribeAIAnalysisTemplatesRequest) {
    request = &DescribeAIAnalysisTemplatesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "DescribeAIAnalysisTemplates")
    
    
    return
}

func NewDescribeAIAnalysisTemplatesResponse() (response *DescribeAIAnalysisTemplatesResponse) {
    response = &DescribeAIAnalysisTemplatesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAIAnalysisTemplates
// 根据内容分析模板唯一标识，获取内容分析模板详情列表。返回结果包含符合条件的所有用户自定义内容分析模板及系统预置视频内容分析模板
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED_TOOMUCHTEMPLATE = "LimitExceeded.TooMuchTemplate"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) DescribeAIAnalysisTemplates(request *DescribeAIAnalysisTemplatesRequest) (response *DescribeAIAnalysisTemplatesResponse, err error) {
    return c.DescribeAIAnalysisTemplatesWithContext(context.Background(), request)
}

// DescribeAIAnalysisTemplates
// 根据内容分析模板唯一标识，获取内容分析模板详情列表。返回结果包含符合条件的所有用户自定义内容分析模板及系统预置视频内容分析模板
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED_TOOMUCHTEMPLATE = "LimitExceeded.TooMuchTemplate"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) DescribeAIAnalysisTemplatesWithContext(ctx context.Context, request *DescribeAIAnalysisTemplatesRequest) (response *DescribeAIAnalysisTemplatesResponse, err error) {
    if request == nil {
        request = NewDescribeAIAnalysisTemplatesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "DescribeAIAnalysisTemplates")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAIAnalysisTemplates require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAIAnalysisTemplatesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAIRecognitionTemplatesRequest() (request *DescribeAIRecognitionTemplatesRequest) {
    request = &DescribeAIRecognitionTemplatesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "DescribeAIRecognitionTemplates")
    
    
    return
}

func NewDescribeAIRecognitionTemplatesResponse() (response *DescribeAIRecognitionTemplatesResponse) {
    response = &DescribeAIRecognitionTemplatesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAIRecognitionTemplates
// 根据内容识别模板唯一标识，获取内容识别模板详情列表。返回结果包含符合条件的所有用户自定义内容识别模板及系统预置视频内容识别模板
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_DEFINITIONS = "InvalidParameterValue.Definitions"
//  INVALIDPARAMETERVALUE_LIMIT = "InvalidParameterValue.Limit"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) DescribeAIRecognitionTemplates(request *DescribeAIRecognitionTemplatesRequest) (response *DescribeAIRecognitionTemplatesResponse, err error) {
    return c.DescribeAIRecognitionTemplatesWithContext(context.Background(), request)
}

// DescribeAIRecognitionTemplates
// 根据内容识别模板唯一标识，获取内容识别模板详情列表。返回结果包含符合条件的所有用户自定义内容识别模板及系统预置视频内容识别模板
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_DEFINITIONS = "InvalidParameterValue.Definitions"
//  INVALIDPARAMETERVALUE_LIMIT = "InvalidParameterValue.Limit"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) DescribeAIRecognitionTemplatesWithContext(ctx context.Context, request *DescribeAIRecognitionTemplatesRequest) (response *DescribeAIRecognitionTemplatesResponse, err error) {
    if request == nil {
        request = NewDescribeAIRecognitionTemplatesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "DescribeAIRecognitionTemplates")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAIRecognitionTemplates require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAIRecognitionTemplatesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAdaptiveDynamicStreamingTemplatesRequest() (request *DescribeAdaptiveDynamicStreamingTemplatesRequest) {
    request = &DescribeAdaptiveDynamicStreamingTemplatesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "DescribeAdaptiveDynamicStreamingTemplates")
    
    
    return
}

func NewDescribeAdaptiveDynamicStreamingTemplatesResponse() (response *DescribeAdaptiveDynamicStreamingTemplatesResponse) {
    response = &DescribeAdaptiveDynamicStreamingTemplatesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAdaptiveDynamicStreamingTemplates
// 查询转自适应码流模板，支持根据条件，分页查询。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_DEFINITIONS = "InvalidParameterValue.Definitions"
//  INVALIDPARAMETERVALUE_LIMIT = "InvalidParameterValue.Limit"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeAdaptiveDynamicStreamingTemplates(request *DescribeAdaptiveDynamicStreamingTemplatesRequest) (response *DescribeAdaptiveDynamicStreamingTemplatesResponse, err error) {
    return c.DescribeAdaptiveDynamicStreamingTemplatesWithContext(context.Background(), request)
}

// DescribeAdaptiveDynamicStreamingTemplates
// 查询转自适应码流模板，支持根据条件，分页查询。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_DEFINITIONS = "InvalidParameterValue.Definitions"
//  INVALIDPARAMETERVALUE_LIMIT = "InvalidParameterValue.Limit"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeAdaptiveDynamicStreamingTemplatesWithContext(ctx context.Context, request *DescribeAdaptiveDynamicStreamingTemplatesRequest) (response *DescribeAdaptiveDynamicStreamingTemplatesResponse, err error) {
    if request == nil {
        request = NewDescribeAdaptiveDynamicStreamingTemplatesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "DescribeAdaptiveDynamicStreamingTemplates")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAdaptiveDynamicStreamingTemplates require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAdaptiveDynamicStreamingTemplatesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAigcImageTaskRequest() (request *DescribeAigcImageTaskRequest) {
    request = &DescribeAigcImageTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "DescribeAigcImageTask")
    
    
    return
}

func NewDescribeAigcImageTaskResponse() (response *DescribeAigcImageTaskResponse) {
    response = &DescribeAigcImageTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAigcImageTask
// 调用该接口，查询AIGC生图片任务进度以及获取生成结果。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_DEFINITIONS = "InvalidParameterValue.Definitions"
//  INVALIDPARAMETERVALUE_LIMIT = "InvalidParameterValue.Limit"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeAigcImageTask(request *DescribeAigcImageTaskRequest) (response *DescribeAigcImageTaskResponse, err error) {
    return c.DescribeAigcImageTaskWithContext(context.Background(), request)
}

// DescribeAigcImageTask
// 调用该接口，查询AIGC生图片任务进度以及获取生成结果。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_DEFINITIONS = "InvalidParameterValue.Definitions"
//  INVALIDPARAMETERVALUE_LIMIT = "InvalidParameterValue.Limit"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeAigcImageTaskWithContext(ctx context.Context, request *DescribeAigcImageTaskRequest) (response *DescribeAigcImageTaskResponse, err error) {
    if request == nil {
        request = NewDescribeAigcImageTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "DescribeAigcImageTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAigcImageTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAigcImageTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAigcVideoTaskRequest() (request *DescribeAigcVideoTaskRequest) {
    request = &DescribeAigcVideoTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "DescribeAigcVideoTask")
    
    
    return
}

func NewDescribeAigcVideoTaskResponse() (response *DescribeAigcVideoTaskResponse) {
    response = &DescribeAigcVideoTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAigcVideoTask
// 调用该接口，用于查询AIGC生视频任务的进度以及获取生成结果。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_DEFINITIONS = "InvalidParameterValue.Definitions"
//  INVALIDPARAMETERVALUE_LIMIT = "InvalidParameterValue.Limit"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeAigcVideoTask(request *DescribeAigcVideoTaskRequest) (response *DescribeAigcVideoTaskResponse, err error) {
    return c.DescribeAigcVideoTaskWithContext(context.Background(), request)
}

// DescribeAigcVideoTask
// 调用该接口，用于查询AIGC生视频任务的进度以及获取生成结果。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_DEFINITIONS = "InvalidParameterValue.Definitions"
//  INVALIDPARAMETERVALUE_LIMIT = "InvalidParameterValue.Limit"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeAigcVideoTaskWithContext(ctx context.Context, request *DescribeAigcVideoTaskRequest) (response *DescribeAigcVideoTaskResponse, err error) {
    if request == nil {
        request = NewDescribeAigcVideoTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "DescribeAigcVideoTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAigcVideoTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAigcVideoTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAnimatedGraphicsTemplatesRequest() (request *DescribeAnimatedGraphicsTemplatesRequest) {
    request = &DescribeAnimatedGraphicsTemplatesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "DescribeAnimatedGraphicsTemplates")
    
    
    return
}

func NewDescribeAnimatedGraphicsTemplatesResponse() (response *DescribeAnimatedGraphicsTemplatesResponse) {
    response = &DescribeAnimatedGraphicsTemplatesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAnimatedGraphicsTemplates
// 查询转动图模板列表，支持根据条件，分页查询。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_LIMIT = "InvalidParameterValue.Limit"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
func (c *Client) DescribeAnimatedGraphicsTemplates(request *DescribeAnimatedGraphicsTemplatesRequest) (response *DescribeAnimatedGraphicsTemplatesResponse, err error) {
    return c.DescribeAnimatedGraphicsTemplatesWithContext(context.Background(), request)
}

// DescribeAnimatedGraphicsTemplates
// 查询转动图模板列表，支持根据条件，分页查询。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_LIMIT = "InvalidParameterValue.Limit"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
func (c *Client) DescribeAnimatedGraphicsTemplatesWithContext(ctx context.Context, request *DescribeAnimatedGraphicsTemplatesRequest) (response *DescribeAnimatedGraphicsTemplatesResponse, err error) {
    if request == nil {
        request = NewDescribeAnimatedGraphicsTemplatesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "DescribeAnimatedGraphicsTemplates")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAnimatedGraphicsTemplates require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAnimatedGraphicsTemplatesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAsrHotwordsRequest() (request *DescribeAsrHotwordsRequest) {
    request = &DescribeAsrHotwordsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "DescribeAsrHotwords")
    
    
    return
}

func NewDescribeAsrHotwordsResponse() (response *DescribeAsrHotwordsResponse) {
    response = &DescribeAsrHotwordsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAsrHotwords
// 查询智能字幕热词库
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ACCESSDBERROR = "InternalError.AccessDBError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeAsrHotwords(request *DescribeAsrHotwordsRequest) (response *DescribeAsrHotwordsResponse, err error) {
    return c.DescribeAsrHotwordsWithContext(context.Background(), request)
}

// DescribeAsrHotwords
// 查询智能字幕热词库
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ACCESSDBERROR = "InternalError.AccessDBError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeAsrHotwordsWithContext(ctx context.Context, request *DescribeAsrHotwordsRequest) (response *DescribeAsrHotwordsResponse, err error) {
    if request == nil {
        request = NewDescribeAsrHotwordsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "DescribeAsrHotwords")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAsrHotwords require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAsrHotwordsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAsrHotwordsListRequest() (request *DescribeAsrHotwordsListRequest) {
    request = &DescribeAsrHotwordsListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "DescribeAsrHotwordsList")
    
    
    return
}

func NewDescribeAsrHotwordsListResponse() (response *DescribeAsrHotwordsListResponse) {
    response = &DescribeAsrHotwordsListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAsrHotwordsList
// 获取热词库列表
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeAsrHotwordsList(request *DescribeAsrHotwordsListRequest) (response *DescribeAsrHotwordsListResponse, err error) {
    return c.DescribeAsrHotwordsListWithContext(context.Background(), request)
}

// DescribeAsrHotwordsList
// 获取热词库列表
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeAsrHotwordsListWithContext(ctx context.Context, request *DescribeAsrHotwordsListRequest) (response *DescribeAsrHotwordsListResponse, err error) {
    if request == nil {
        request = NewDescribeAsrHotwordsListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "DescribeAsrHotwordsList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAsrHotwordsList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAsrHotwordsListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBatchTaskDetailRequest() (request *DescribeBatchTaskDetailRequest) {
    request = &DescribeBatchTaskDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "DescribeBatchTaskDetail")
    
    
    return
}

func NewDescribeBatchTaskDetailResponse() (response *DescribeBatchTaskDetailResponse) {
    response = &DescribeBatchTaskDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeBatchTaskDetail
// 通过任务 ID 查询任务的执行状态和结果的详细信息（最多可以查询7天之内提交的任务）。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeBatchTaskDetail(request *DescribeBatchTaskDetailRequest) (response *DescribeBatchTaskDetailResponse, err error) {
    return c.DescribeBatchTaskDetailWithContext(context.Background(), request)
}

// DescribeBatchTaskDetail
// 通过任务 ID 查询任务的执行状态和结果的详细信息（最多可以查询7天之内提交的任务）。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeBatchTaskDetailWithContext(ctx context.Context, request *DescribeBatchTaskDetailRequest) (response *DescribeBatchTaskDetailResponse, err error) {
    if request == nil {
        request = NewDescribeBatchTaskDetailRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "DescribeBatchTaskDetail")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBatchTaskDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBatchTaskDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBlindWatermarkTemplatesRequest() (request *DescribeBlindWatermarkTemplatesRequest) {
    request = &DescribeBlindWatermarkTemplatesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "DescribeBlindWatermarkTemplates")
    
    
    return
}

func NewDescribeBlindWatermarkTemplatesResponse() (response *DescribeBlindWatermarkTemplatesResponse) {
    response = &DescribeBlindWatermarkTemplatesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeBlindWatermarkTemplates
// 查询用户自定义数字水印模板，支持根据条件，分页查询。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_DEFINITIONS = "InvalidParameterValue.Definitions"
//  INVALIDPARAMETERVALUE_LIMIT = "InvalidParameterValue.Limit"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) DescribeBlindWatermarkTemplates(request *DescribeBlindWatermarkTemplatesRequest) (response *DescribeBlindWatermarkTemplatesResponse, err error) {
    return c.DescribeBlindWatermarkTemplatesWithContext(context.Background(), request)
}

// DescribeBlindWatermarkTemplates
// 查询用户自定义数字水印模板，支持根据条件，分页查询。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_DEFINITIONS = "InvalidParameterValue.Definitions"
//  INVALIDPARAMETERVALUE_LIMIT = "InvalidParameterValue.Limit"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) DescribeBlindWatermarkTemplatesWithContext(ctx context.Context, request *DescribeBlindWatermarkTemplatesRequest) (response *DescribeBlindWatermarkTemplatesResponse, err error) {
    if request == nil {
        request = NewDescribeBlindWatermarkTemplatesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "DescribeBlindWatermarkTemplates")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBlindWatermarkTemplates require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBlindWatermarkTemplatesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeContentReviewTemplatesRequest() (request *DescribeContentReviewTemplatesRequest) {
    request = &DescribeContentReviewTemplatesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "DescribeContentReviewTemplates")
    
    
    return
}

func NewDescribeContentReviewTemplatesResponse() (response *DescribeContentReviewTemplatesResponse) {
    response = &DescribeContentReviewTemplatesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeContentReviewTemplates
// 根据智能审核模板唯一标识，获取智能审核模板详情列表。返回结果包含符合条件的所有用户自定义模板及系统预置智能审核模板。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_DEFINITIONS = "InvalidParameterValue.Definitions"
//  INVALIDPARAMETERVALUE_LIMIT = "InvalidParameterValue.Limit"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) DescribeContentReviewTemplates(request *DescribeContentReviewTemplatesRequest) (response *DescribeContentReviewTemplatesResponse, err error) {
    return c.DescribeContentReviewTemplatesWithContext(context.Background(), request)
}

// DescribeContentReviewTemplates
// 根据智能审核模板唯一标识，获取智能审核模板详情列表。返回结果包含符合条件的所有用户自定义模板及系统预置智能审核模板。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_DEFINITIONS = "InvalidParameterValue.Definitions"
//  INVALIDPARAMETERVALUE_LIMIT = "InvalidParameterValue.Limit"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) DescribeContentReviewTemplatesWithContext(ctx context.Context, request *DescribeContentReviewTemplatesRequest) (response *DescribeContentReviewTemplatesResponse, err error) {
    if request == nil {
        request = NewDescribeContentReviewTemplatesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "DescribeContentReviewTemplates")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeContentReviewTemplates require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeContentReviewTemplatesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeGroupAttachFlowsByIdRequest() (request *DescribeGroupAttachFlowsByIdRequest) {
    request = &DescribeGroupAttachFlowsByIdRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "DescribeGroupAttachFlowsById")
    
    
    return
}

func NewDescribeGroupAttachFlowsByIdResponse() (response *DescribeGroupAttachFlowsByIdResponse) {
    response = &DescribeGroupAttachFlowsByIdResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeGroupAttachFlowsById
// 根据安全组反差关联的Flow信息。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeGroupAttachFlowsById(request *DescribeGroupAttachFlowsByIdRequest) (response *DescribeGroupAttachFlowsByIdResponse, err error) {
    return c.DescribeGroupAttachFlowsByIdWithContext(context.Background(), request)
}

// DescribeGroupAttachFlowsById
// 根据安全组反差关联的Flow信息。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeGroupAttachFlowsByIdWithContext(ctx context.Context, request *DescribeGroupAttachFlowsByIdRequest) (response *DescribeGroupAttachFlowsByIdResponse, err error) {
    if request == nil {
        request = NewDescribeGroupAttachFlowsByIdRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "DescribeGroupAttachFlowsById")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeGroupAttachFlowsById require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeGroupAttachFlowsByIdResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeImageSpriteTemplatesRequest() (request *DescribeImageSpriteTemplatesRequest) {
    request = &DescribeImageSpriteTemplatesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "DescribeImageSpriteTemplates")
    
    
    return
}

func NewDescribeImageSpriteTemplatesResponse() (response *DescribeImageSpriteTemplatesResponse) {
    response = &DescribeImageSpriteTemplatesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeImageSpriteTemplates
// 查询雪碧图模板，支持根据条件，分页查询。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_LIMIT = "InvalidParameterValue.Limit"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
func (c *Client) DescribeImageSpriteTemplates(request *DescribeImageSpriteTemplatesRequest) (response *DescribeImageSpriteTemplatesResponse, err error) {
    return c.DescribeImageSpriteTemplatesWithContext(context.Background(), request)
}

// DescribeImageSpriteTemplates
// 查询雪碧图模板，支持根据条件，分页查询。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_LIMIT = "InvalidParameterValue.Limit"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
func (c *Client) DescribeImageSpriteTemplatesWithContext(ctx context.Context, request *DescribeImageSpriteTemplatesRequest) (response *DescribeImageSpriteTemplatesResponse, err error) {
    if request == nil {
        request = NewDescribeImageSpriteTemplatesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "DescribeImageSpriteTemplates")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeImageSpriteTemplates require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeImageSpriteTemplatesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeImageTaskDetailRequest() (request *DescribeImageTaskDetailRequest) {
    request = &DescribeImageTaskDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "DescribeImageTaskDetail")
    
    
    return
}

func NewDescribeImageTaskDetailResponse() (response *DescribeImageTaskDetailResponse) {
    response = &DescribeImageTaskDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeImageTaskDetail
// 通过任务 ID 查询任务的执行状态和结果的详细信息（最多可以查询7天之内提交的任务）。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeImageTaskDetail(request *DescribeImageTaskDetailRequest) (response *DescribeImageTaskDetailResponse, err error) {
    return c.DescribeImageTaskDetailWithContext(context.Background(), request)
}

// DescribeImageTaskDetail
// 通过任务 ID 查询任务的执行状态和结果的详细信息（最多可以查询7天之内提交的任务）。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeImageTaskDetailWithContext(ctx context.Context, request *DescribeImageTaskDetailRequest) (response *DescribeImageTaskDetailResponse, err error) {
    if request == nil {
        request = NewDescribeImageTaskDetailRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "DescribeImageTaskDetail")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeImageTaskDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeImageTaskDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLiveRecordTemplatesRequest() (request *DescribeLiveRecordTemplatesRequest) {
    request = &DescribeLiveRecordTemplatesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "DescribeLiveRecordTemplates")
    
    
    return
}

func NewDescribeLiveRecordTemplatesResponse() (response *DescribeLiveRecordTemplatesResponse) {
    response = &DescribeLiveRecordTemplatesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeLiveRecordTemplates
// 获取直播录制模板
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ACCESSDBERROR = "InternalError.AccessDBError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LIMIT = "InvalidParameterValue.Limit"
func (c *Client) DescribeLiveRecordTemplates(request *DescribeLiveRecordTemplatesRequest) (response *DescribeLiveRecordTemplatesResponse, err error) {
    return c.DescribeLiveRecordTemplatesWithContext(context.Background(), request)
}

// DescribeLiveRecordTemplates
// 获取直播录制模板
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ACCESSDBERROR = "InternalError.AccessDBError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LIMIT = "InvalidParameterValue.Limit"
func (c *Client) DescribeLiveRecordTemplatesWithContext(ctx context.Context, request *DescribeLiveRecordTemplatesRequest) (response *DescribeLiveRecordTemplatesResponse, err error) {
    if request == nil {
        request = NewDescribeLiveRecordTemplatesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "DescribeLiveRecordTemplates")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLiveRecordTemplates require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLiveRecordTemplatesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMediaMetaDataRequest() (request *DescribeMediaMetaDataRequest) {
    request = &DescribeMediaMetaDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "DescribeMediaMetaData")
    
    
    return
}

func NewDescribeMediaMetaDataResponse() (response *DescribeMediaMetaDataResponse) {
    response = &DescribeMediaMetaDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeMediaMetaData
// 获取媒体的元信息，包括视频画面宽、高、编码格式、时长、帧率等。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_SRCFILE = "InvalidParameterValue.SrcFile"
func (c *Client) DescribeMediaMetaData(request *DescribeMediaMetaDataRequest) (response *DescribeMediaMetaDataResponse, err error) {
    return c.DescribeMediaMetaDataWithContext(context.Background(), request)
}

// DescribeMediaMetaData
// 获取媒体的元信息，包括视频画面宽、高、编码格式、时长、帧率等。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_SRCFILE = "InvalidParameterValue.SrcFile"
func (c *Client) DescribeMediaMetaDataWithContext(ctx context.Context, request *DescribeMediaMetaDataRequest) (response *DescribeMediaMetaDataResponse, err error) {
    if request == nil {
        request = NewDescribeMediaMetaDataRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "DescribeMediaMetaData")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMediaMetaData require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMediaMetaDataResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePersonSamplesRequest() (request *DescribePersonSamplesRequest) {
    request = &DescribePersonSamplesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "DescribePersonSamples")
    
    
    return
}

func NewDescribePersonSamplesResponse() (response *DescribePersonSamplesResponse) {
    response = &DescribePersonSamplesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribePersonSamples
// 该接口用于查询素材样本信息，支持根据素材 ID、名称、标签，分页查询。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribePersonSamples(request *DescribePersonSamplesRequest) (response *DescribePersonSamplesResponse, err error) {
    return c.DescribePersonSamplesWithContext(context.Background(), request)
}

// DescribePersonSamples
// 该接口用于查询素材样本信息，支持根据素材 ID、名称、标签，分页查询。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribePersonSamplesWithContext(ctx context.Context, request *DescribePersonSamplesRequest) (response *DescribePersonSamplesResponse, err error) {
    if request == nil {
        request = NewDescribePersonSamplesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "DescribePersonSamples")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePersonSamples require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePersonSamplesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeProcessImageTemplatesRequest() (request *DescribeProcessImageTemplatesRequest) {
    request = &DescribeProcessImageTemplatesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "DescribeProcessImageTemplates")
    
    
    return
}

func NewDescribeProcessImageTemplatesResponse() (response *DescribeProcessImageTemplatesResponse) {
    response = &DescribeProcessImageTemplatesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeProcessImageTemplates
// 查询图片处理模板列表。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DEFINITIONS = "InvalidParameterValue.Definitions"
//  INVALIDPARAMETERVALUE_LIMIT = "InvalidParameterValue.Limit"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) DescribeProcessImageTemplates(request *DescribeProcessImageTemplatesRequest) (response *DescribeProcessImageTemplatesResponse, err error) {
    return c.DescribeProcessImageTemplatesWithContext(context.Background(), request)
}

// DescribeProcessImageTemplates
// 查询图片处理模板列表。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DEFINITIONS = "InvalidParameterValue.Definitions"
//  INVALIDPARAMETERVALUE_LIMIT = "InvalidParameterValue.Limit"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) DescribeProcessImageTemplatesWithContext(ctx context.Context, request *DescribeProcessImageTemplatesRequest) (response *DescribeProcessImageTemplatesResponse, err error) {
    if request == nil {
        request = NewDescribeProcessImageTemplatesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "DescribeProcessImageTemplates")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeProcessImageTemplates require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeProcessImageTemplatesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeQualityControlTemplatesRequest() (request *DescribeQualityControlTemplatesRequest) {
    request = &DescribeQualityControlTemplatesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "DescribeQualityControlTemplates")
    
    
    return
}

func NewDescribeQualityControlTemplatesResponse() (response *DescribeQualityControlTemplatesResponse) {
    response = &DescribeQualityControlTemplatesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeQualityControlTemplates
// 查询用户自定义媒体质检模板，支持根据条件，分页查询。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DEFINITIONS = "InvalidParameterValue.Definitions"
//  INVALIDPARAMETERVALUE_LIMIT = "InvalidParameterValue.Limit"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) DescribeQualityControlTemplates(request *DescribeQualityControlTemplatesRequest) (response *DescribeQualityControlTemplatesResponse, err error) {
    return c.DescribeQualityControlTemplatesWithContext(context.Background(), request)
}

// DescribeQualityControlTemplates
// 查询用户自定义媒体质检模板，支持根据条件，分页查询。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DEFINITIONS = "InvalidParameterValue.Definitions"
//  INVALIDPARAMETERVALUE_LIMIT = "InvalidParameterValue.Limit"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) DescribeQualityControlTemplatesWithContext(ctx context.Context, request *DescribeQualityControlTemplatesRequest) (response *DescribeQualityControlTemplatesResponse, err error) {
    if request == nil {
        request = NewDescribeQualityControlTemplatesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "DescribeQualityControlTemplates")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeQualityControlTemplates require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeQualityControlTemplatesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSampleSnapshotTemplatesRequest() (request *DescribeSampleSnapshotTemplatesRequest) {
    request = &DescribeSampleSnapshotTemplatesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "DescribeSampleSnapshotTemplates")
    
    
    return
}

func NewDescribeSampleSnapshotTemplatesResponse() (response *DescribeSampleSnapshotTemplatesResponse) {
    response = &DescribeSampleSnapshotTemplatesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSampleSnapshotTemplates
// 查询采样截图模板，支持根据条件，分页查询。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_LIMIT = "InvalidParameterValue.Limit"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
func (c *Client) DescribeSampleSnapshotTemplates(request *DescribeSampleSnapshotTemplatesRequest) (response *DescribeSampleSnapshotTemplatesResponse, err error) {
    return c.DescribeSampleSnapshotTemplatesWithContext(context.Background(), request)
}

// DescribeSampleSnapshotTemplates
// 查询采样截图模板，支持根据条件，分页查询。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_LIMIT = "InvalidParameterValue.Limit"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
func (c *Client) DescribeSampleSnapshotTemplatesWithContext(ctx context.Context, request *DescribeSampleSnapshotTemplatesRequest) (response *DescribeSampleSnapshotTemplatesResponse, err error) {
    if request == nil {
        request = NewDescribeSampleSnapshotTemplatesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "DescribeSampleSnapshotTemplates")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSampleSnapshotTemplates require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSampleSnapshotTemplatesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSchedulesRequest() (request *DescribeSchedulesRequest) {
    request = &DescribeSchedulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "DescribeSchedules")
    
    
    return
}

func NewDescribeSchedulesResponse() (response *DescribeSchedulesResponse) {
    response = &DescribeSchedulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSchedules
// 查询编排。
//
// 可能返回的错误码:
//  FAILEDOPERATION_GENERATERESOURCE = "FailedOperation.GenerateResource"
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  FAILEDOPERATION_INVALIDUSER = "FailedOperation.InvalidUser"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ACCESSDBERROR = "InternalError.AccessDBError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeSchedules(request *DescribeSchedulesRequest) (response *DescribeSchedulesResponse, err error) {
    return c.DescribeSchedulesWithContext(context.Background(), request)
}

// DescribeSchedules
// 查询编排。
//
// 可能返回的错误码:
//  FAILEDOPERATION_GENERATERESOURCE = "FailedOperation.GenerateResource"
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  FAILEDOPERATION_INVALIDUSER = "FailedOperation.InvalidUser"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ACCESSDBERROR = "InternalError.AccessDBError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeSchedulesWithContext(ctx context.Context, request *DescribeSchedulesRequest) (response *DescribeSchedulesResponse, err error) {
    if request == nil {
        request = NewDescribeSchedulesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "DescribeSchedules")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSchedules require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSchedulesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSmartEraseTemplatesRequest() (request *DescribeSmartEraseTemplatesRequest) {
    request = &DescribeSmartEraseTemplatesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "DescribeSmartEraseTemplates")
    
    
    return
}

func NewDescribeSmartEraseTemplatesResponse() (response *DescribeSmartEraseTemplatesResponse) {
    response = &DescribeSmartEraseTemplatesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSmartEraseTemplates
// 根据智能擦除模板唯一标识，获取智能擦除模板详情列表。返回结果包含符合条件的所有用户自定义智能擦除模板及系统预置智能擦除模板
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_DEFINITIONS = "InvalidParameterValue.Definitions"
//  INVALIDPARAMETERVALUE_ERASETYPE = "InvalidParameterValue.EraseType"
//  INVALIDPARAMETERVALUE_LIMIT = "InvalidParameterValue.Limit"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) DescribeSmartEraseTemplates(request *DescribeSmartEraseTemplatesRequest) (response *DescribeSmartEraseTemplatesResponse, err error) {
    return c.DescribeSmartEraseTemplatesWithContext(context.Background(), request)
}

// DescribeSmartEraseTemplates
// 根据智能擦除模板唯一标识，获取智能擦除模板详情列表。返回结果包含符合条件的所有用户自定义智能擦除模板及系统预置智能擦除模板
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_DEFINITIONS = "InvalidParameterValue.Definitions"
//  INVALIDPARAMETERVALUE_ERASETYPE = "InvalidParameterValue.EraseType"
//  INVALIDPARAMETERVALUE_LIMIT = "InvalidParameterValue.Limit"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) DescribeSmartEraseTemplatesWithContext(ctx context.Context, request *DescribeSmartEraseTemplatesRequest) (response *DescribeSmartEraseTemplatesResponse, err error) {
    if request == nil {
        request = NewDescribeSmartEraseTemplatesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "DescribeSmartEraseTemplates")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSmartEraseTemplates require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSmartEraseTemplatesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSmartSubtitleTemplatesRequest() (request *DescribeSmartSubtitleTemplatesRequest) {
    request = &DescribeSmartSubtitleTemplatesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "DescribeSmartSubtitleTemplates")
    
    
    return
}

func NewDescribeSmartSubtitleTemplatesResponse() (response *DescribeSmartSubtitleTemplatesResponse) {
    response = &DescribeSmartSubtitleTemplatesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSmartSubtitleTemplates
// 根据智能字幕 模板唯一标识，获取智能字幕模板详情列表。返回结果包含符合条件的所有用户自定义智能字幕模板及系统预置智能字幕模板
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_DEFINITIONS = "InvalidParameterValue.Definitions"
//  INVALIDPARAMETERVALUE_LIMIT = "InvalidParameterValue.Limit"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) DescribeSmartSubtitleTemplates(request *DescribeSmartSubtitleTemplatesRequest) (response *DescribeSmartSubtitleTemplatesResponse, err error) {
    return c.DescribeSmartSubtitleTemplatesWithContext(context.Background(), request)
}

// DescribeSmartSubtitleTemplates
// 根据智能字幕 模板唯一标识，获取智能字幕模板详情列表。返回结果包含符合条件的所有用户自定义智能字幕模板及系统预置智能字幕模板
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_DEFINITIONS = "InvalidParameterValue.Definitions"
//  INVALIDPARAMETERVALUE_LIMIT = "InvalidParameterValue.Limit"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) DescribeSmartSubtitleTemplatesWithContext(ctx context.Context, request *DescribeSmartSubtitleTemplatesRequest) (response *DescribeSmartSubtitleTemplatesResponse, err error) {
    if request == nil {
        request = NewDescribeSmartSubtitleTemplatesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "DescribeSmartSubtitleTemplates")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSmartSubtitleTemplates require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSmartSubtitleTemplatesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSnapshotByTimeOffsetTemplatesRequest() (request *DescribeSnapshotByTimeOffsetTemplatesRequest) {
    request = &DescribeSnapshotByTimeOffsetTemplatesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "DescribeSnapshotByTimeOffsetTemplates")
    
    
    return
}

func NewDescribeSnapshotByTimeOffsetTemplatesResponse() (response *DescribeSnapshotByTimeOffsetTemplatesResponse) {
    response = &DescribeSnapshotByTimeOffsetTemplatesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSnapshotByTimeOffsetTemplates
// 查询指定时间点截图模板，支持根据条件，分页查询。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_LIMIT = "InvalidParameterValue.Limit"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
func (c *Client) DescribeSnapshotByTimeOffsetTemplates(request *DescribeSnapshotByTimeOffsetTemplatesRequest) (response *DescribeSnapshotByTimeOffsetTemplatesResponse, err error) {
    return c.DescribeSnapshotByTimeOffsetTemplatesWithContext(context.Background(), request)
}

// DescribeSnapshotByTimeOffsetTemplates
// 查询指定时间点截图模板，支持根据条件，分页查询。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_LIMIT = "InvalidParameterValue.Limit"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
func (c *Client) DescribeSnapshotByTimeOffsetTemplatesWithContext(ctx context.Context, request *DescribeSnapshotByTimeOffsetTemplatesRequest) (response *DescribeSnapshotByTimeOffsetTemplatesResponse, err error) {
    if request == nil {
        request = NewDescribeSnapshotByTimeOffsetTemplatesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "DescribeSnapshotByTimeOffsetTemplates")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSnapshotByTimeOffsetTemplates require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSnapshotByTimeOffsetTemplatesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeStreamLinkActivateStateRequest() (request *DescribeStreamLinkActivateStateRequest) {
    request = &DescribeStreamLinkActivateStateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "DescribeStreamLinkActivateState")
    
    
    return
}

func NewDescribeStreamLinkActivateStateResponse() (response *DescribeStreamLinkActivateStateResponse) {
    response = &DescribeStreamLinkActivateStateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeStreamLinkActivateState
// 查询媒体传输开通状态。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeStreamLinkActivateState(request *DescribeStreamLinkActivateStateRequest) (response *DescribeStreamLinkActivateStateResponse, err error) {
    return c.DescribeStreamLinkActivateStateWithContext(context.Background(), request)
}

// DescribeStreamLinkActivateState
// 查询媒体传输开通状态。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeStreamLinkActivateStateWithContext(ctx context.Context, request *DescribeStreamLinkActivateStateRequest) (response *DescribeStreamLinkActivateStateResponse, err error) {
    if request == nil {
        request = NewDescribeStreamLinkActivateStateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "DescribeStreamLinkActivateState")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeStreamLinkActivateState require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeStreamLinkActivateStateResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeStreamLinkEventRequest() (request *DescribeStreamLinkEventRequest) {
    request = &DescribeStreamLinkEventRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "DescribeStreamLinkEvent")
    
    
    return
}

func NewDescribeStreamLinkEventResponse() (response *DescribeStreamLinkEventResponse) {
    response = &DescribeStreamLinkEventResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeStreamLinkEvent
// 查询媒体传输事件的配置信息。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_PAGENUM = "InvalidParameter.PageNum"
//  INVALIDPARAMETER_PAGESIZE = "InvalidParameter.PageSize"
func (c *Client) DescribeStreamLinkEvent(request *DescribeStreamLinkEventRequest) (response *DescribeStreamLinkEventResponse, err error) {
    return c.DescribeStreamLinkEventWithContext(context.Background(), request)
}

// DescribeStreamLinkEvent
// 查询媒体传输事件的配置信息。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_PAGENUM = "InvalidParameter.PageNum"
//  INVALIDPARAMETER_PAGESIZE = "InvalidParameter.PageSize"
func (c *Client) DescribeStreamLinkEventWithContext(ctx context.Context, request *DescribeStreamLinkEventRequest) (response *DescribeStreamLinkEventResponse, err error) {
    if request == nil {
        request = NewDescribeStreamLinkEventRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "DescribeStreamLinkEvent")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeStreamLinkEvent require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeStreamLinkEventResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeStreamLinkEventAttachedFlowsRequest() (request *DescribeStreamLinkEventAttachedFlowsRequest) {
    request = &DescribeStreamLinkEventAttachedFlowsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "DescribeStreamLinkEventAttachedFlows")
    
    
    return
}

func NewDescribeStreamLinkEventAttachedFlowsResponse() (response *DescribeStreamLinkEventAttachedFlowsResponse) {
    response = &DescribeStreamLinkEventAttachedFlowsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeStreamLinkEventAttachedFlows
// 查询媒体传输事件关联的所有媒体输入流的配置信息。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_PAGENUM = "InvalidParameter.PageNum"
//  INVALIDPARAMETER_PAGESIZE = "InvalidParameter.PageSize"
func (c *Client) DescribeStreamLinkEventAttachedFlows(request *DescribeStreamLinkEventAttachedFlowsRequest) (response *DescribeStreamLinkEventAttachedFlowsResponse, err error) {
    return c.DescribeStreamLinkEventAttachedFlowsWithContext(context.Background(), request)
}

// DescribeStreamLinkEventAttachedFlows
// 查询媒体传输事件关联的所有媒体输入流的配置信息。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_PAGENUM = "InvalidParameter.PageNum"
//  INVALIDPARAMETER_PAGESIZE = "InvalidParameter.PageSize"
func (c *Client) DescribeStreamLinkEventAttachedFlowsWithContext(ctx context.Context, request *DescribeStreamLinkEventAttachedFlowsRequest) (response *DescribeStreamLinkEventAttachedFlowsResponse, err error) {
    if request == nil {
        request = NewDescribeStreamLinkEventAttachedFlowsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "DescribeStreamLinkEventAttachedFlows")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeStreamLinkEventAttachedFlows require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeStreamLinkEventAttachedFlowsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeStreamLinkEventsRequest() (request *DescribeStreamLinkEventsRequest) {
    request = &DescribeStreamLinkEventsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "DescribeStreamLinkEvents")
    
    
    return
}

func NewDescribeStreamLinkEventsResponse() (response *DescribeStreamLinkEventsResponse) {
    response = &DescribeStreamLinkEventsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeStreamLinkEvents
// 批量查询媒体传输事件的配置信息。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_PAGENUM = "InvalidParameter.PageNum"
//  INVALIDPARAMETER_PAGESIZE = "InvalidParameter.PageSize"
func (c *Client) DescribeStreamLinkEvents(request *DescribeStreamLinkEventsRequest) (response *DescribeStreamLinkEventsResponse, err error) {
    return c.DescribeStreamLinkEventsWithContext(context.Background(), request)
}

// DescribeStreamLinkEvents
// 批量查询媒体传输事件的配置信息。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_PAGENUM = "InvalidParameter.PageNum"
//  INVALIDPARAMETER_PAGESIZE = "InvalidParameter.PageSize"
func (c *Client) DescribeStreamLinkEventsWithContext(ctx context.Context, request *DescribeStreamLinkEventsRequest) (response *DescribeStreamLinkEventsResponse, err error) {
    if request == nil {
        request = NewDescribeStreamLinkEventsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "DescribeStreamLinkEvents")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeStreamLinkEvents require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeStreamLinkEventsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeStreamLinkFlowRequest() (request *DescribeStreamLinkFlowRequest) {
    request = &DescribeStreamLinkFlowRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "DescribeStreamLinkFlow")
    
    
    return
}

func NewDescribeStreamLinkFlowResponse() (response *DescribeStreamLinkFlowResponse) {
    response = &DescribeStreamLinkFlowResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeStreamLinkFlow
// 查询媒体输入流的配置信息。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
func (c *Client) DescribeStreamLinkFlow(request *DescribeStreamLinkFlowRequest) (response *DescribeStreamLinkFlowResponse, err error) {
    return c.DescribeStreamLinkFlowWithContext(context.Background(), request)
}

// DescribeStreamLinkFlow
// 查询媒体输入流的配置信息。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
func (c *Client) DescribeStreamLinkFlowWithContext(ctx context.Context, request *DescribeStreamLinkFlowRequest) (response *DescribeStreamLinkFlowResponse, err error) {
    if request == nil {
        request = NewDescribeStreamLinkFlowRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "DescribeStreamLinkFlow")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeStreamLinkFlow require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeStreamLinkFlowResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeStreamLinkFlowLogsRequest() (request *DescribeStreamLinkFlowLogsRequest) {
    request = &DescribeStreamLinkFlowLogsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "DescribeStreamLinkFlowLogs")
    
    
    return
}

func NewDescribeStreamLinkFlowLogsResponse() (response *DescribeStreamLinkFlowLogsResponse) {
    response = &DescribeStreamLinkFlowLogsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeStreamLinkFlowLogs
// 查询媒体传输流的日志信息。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ENDTIME = "InvalidParameter.EndTime"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_NOTFOUND = "InvalidParameter.NotFound"
//  INVALIDPARAMETER_PAGENUM = "InvalidParameter.PageNum"
//  INVALIDPARAMETER_PAGESIZE = "InvalidParameter.PageSize"
//  INVALIDPARAMETER_PIPELINE = "InvalidParameter.Pipeline"
//  INVALIDPARAMETER_SORTTYPE = "InvalidParameter.SortType"
//  INVALIDPARAMETER_STARTTIME = "InvalidParameter.StartTime"
//  INVALIDPARAMETER_TYPE = "InvalidParameter.Type"
func (c *Client) DescribeStreamLinkFlowLogs(request *DescribeStreamLinkFlowLogsRequest) (response *DescribeStreamLinkFlowLogsResponse, err error) {
    return c.DescribeStreamLinkFlowLogsWithContext(context.Background(), request)
}

// DescribeStreamLinkFlowLogs
// 查询媒体传输流的日志信息。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ENDTIME = "InvalidParameter.EndTime"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_NOTFOUND = "InvalidParameter.NotFound"
//  INVALIDPARAMETER_PAGENUM = "InvalidParameter.PageNum"
//  INVALIDPARAMETER_PAGESIZE = "InvalidParameter.PageSize"
//  INVALIDPARAMETER_PIPELINE = "InvalidParameter.Pipeline"
//  INVALIDPARAMETER_SORTTYPE = "InvalidParameter.SortType"
//  INVALIDPARAMETER_STARTTIME = "InvalidParameter.StartTime"
//  INVALIDPARAMETER_TYPE = "InvalidParameter.Type"
func (c *Client) DescribeStreamLinkFlowLogsWithContext(ctx context.Context, request *DescribeStreamLinkFlowLogsRequest) (response *DescribeStreamLinkFlowLogsResponse, err error) {
    if request == nil {
        request = NewDescribeStreamLinkFlowLogsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "DescribeStreamLinkFlowLogs")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeStreamLinkFlowLogs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeStreamLinkFlowLogsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeStreamLinkFlowMediaStatisticsRequest() (request *DescribeStreamLinkFlowMediaStatisticsRequest) {
    request = &DescribeStreamLinkFlowMediaStatisticsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "DescribeStreamLinkFlowMediaStatistics")
    
    
    return
}

func NewDescribeStreamLinkFlowMediaStatisticsResponse() (response *DescribeStreamLinkFlowMediaStatisticsResponse) {
    response = &DescribeStreamLinkFlowMediaStatisticsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeStreamLinkFlowMediaStatistics
// 查询媒体传输流的媒体质量数据。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ENDTIME = "InvalidParameter.EndTime"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_INPUTOUTPUTID = "InvalidParameter.InputOutputId"
//  INVALIDPARAMETER_NOTFOUND = "InvalidParameter.NotFound"
//  INVALIDPARAMETER_PERIOD = "InvalidParameter.Period"
//  INVALIDPARAMETER_PIPELINE = "InvalidParameter.Pipeline"
//  INVALIDPARAMETER_STARTTIME = "InvalidParameter.StartTime"
//  INVALIDPARAMETER_TYPE = "InvalidParameter.Type"
func (c *Client) DescribeStreamLinkFlowMediaStatistics(request *DescribeStreamLinkFlowMediaStatisticsRequest) (response *DescribeStreamLinkFlowMediaStatisticsResponse, err error) {
    return c.DescribeStreamLinkFlowMediaStatisticsWithContext(context.Background(), request)
}

// DescribeStreamLinkFlowMediaStatistics
// 查询媒体传输流的媒体质量数据。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ENDTIME = "InvalidParameter.EndTime"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_INPUTOUTPUTID = "InvalidParameter.InputOutputId"
//  INVALIDPARAMETER_NOTFOUND = "InvalidParameter.NotFound"
//  INVALIDPARAMETER_PERIOD = "InvalidParameter.Period"
//  INVALIDPARAMETER_PIPELINE = "InvalidParameter.Pipeline"
//  INVALIDPARAMETER_STARTTIME = "InvalidParameter.StartTime"
//  INVALIDPARAMETER_TYPE = "InvalidParameter.Type"
func (c *Client) DescribeStreamLinkFlowMediaStatisticsWithContext(ctx context.Context, request *DescribeStreamLinkFlowMediaStatisticsRequest) (response *DescribeStreamLinkFlowMediaStatisticsResponse, err error) {
    if request == nil {
        request = NewDescribeStreamLinkFlowMediaStatisticsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "DescribeStreamLinkFlowMediaStatistics")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeStreamLinkFlowMediaStatistics require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeStreamLinkFlowMediaStatisticsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeStreamLinkFlowRealtimeStatusRequest() (request *DescribeStreamLinkFlowRealtimeStatusRequest) {
    request = &DescribeStreamLinkFlowRealtimeStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "DescribeStreamLinkFlowRealtimeStatus")
    
    
    return
}

func NewDescribeStreamLinkFlowRealtimeStatusResponse() (response *DescribeStreamLinkFlowRealtimeStatusResponse) {
    response = &DescribeStreamLinkFlowRealtimeStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeStreamLinkFlowRealtimeStatus
// 实时查询流的当前状态
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_NOTFOUND = "InvalidParameter.NotFound"
func (c *Client) DescribeStreamLinkFlowRealtimeStatus(request *DescribeStreamLinkFlowRealtimeStatusRequest) (response *DescribeStreamLinkFlowRealtimeStatusResponse, err error) {
    return c.DescribeStreamLinkFlowRealtimeStatusWithContext(context.Background(), request)
}

// DescribeStreamLinkFlowRealtimeStatus
// 实时查询流的当前状态
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_NOTFOUND = "InvalidParameter.NotFound"
func (c *Client) DescribeStreamLinkFlowRealtimeStatusWithContext(ctx context.Context, request *DescribeStreamLinkFlowRealtimeStatusRequest) (response *DescribeStreamLinkFlowRealtimeStatusResponse, err error) {
    if request == nil {
        request = NewDescribeStreamLinkFlowRealtimeStatusRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "DescribeStreamLinkFlowRealtimeStatus")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeStreamLinkFlowRealtimeStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeStreamLinkFlowRealtimeStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeStreamLinkFlowSRTStatisticsRequest() (request *DescribeStreamLinkFlowSRTStatisticsRequest) {
    request = &DescribeStreamLinkFlowSRTStatisticsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "DescribeStreamLinkFlowSRTStatistics")
    
    
    return
}

func NewDescribeStreamLinkFlowSRTStatisticsResponse() (response *DescribeStreamLinkFlowSRTStatisticsResponse) {
    response = &DescribeStreamLinkFlowSRTStatisticsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeStreamLinkFlowSRTStatistics
// 查询媒体传输流的SRT质量数据。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ENDTIME = "InvalidParameter.EndTime"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_INPUTOUTPUTID = "InvalidParameter.InputOutputId"
//  INVALIDPARAMETER_NOTFOUND = "InvalidParameter.NotFound"
//  INVALIDPARAMETER_PERIOD = "InvalidParameter.Period"
//  INVALIDPARAMETER_PIPELINE = "InvalidParameter.Pipeline"
//  INVALIDPARAMETER_STARTTIME = "InvalidParameter.StartTime"
//  INVALIDPARAMETER_TYPE = "InvalidParameter.Type"
func (c *Client) DescribeStreamLinkFlowSRTStatistics(request *DescribeStreamLinkFlowSRTStatisticsRequest) (response *DescribeStreamLinkFlowSRTStatisticsResponse, err error) {
    return c.DescribeStreamLinkFlowSRTStatisticsWithContext(context.Background(), request)
}

// DescribeStreamLinkFlowSRTStatistics
// 查询媒体传输流的SRT质量数据。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ENDTIME = "InvalidParameter.EndTime"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_INPUTOUTPUTID = "InvalidParameter.InputOutputId"
//  INVALIDPARAMETER_NOTFOUND = "InvalidParameter.NotFound"
//  INVALIDPARAMETER_PERIOD = "InvalidParameter.Period"
//  INVALIDPARAMETER_PIPELINE = "InvalidParameter.Pipeline"
//  INVALIDPARAMETER_STARTTIME = "InvalidParameter.StartTime"
//  INVALIDPARAMETER_TYPE = "InvalidParameter.Type"
func (c *Client) DescribeStreamLinkFlowSRTStatisticsWithContext(ctx context.Context, request *DescribeStreamLinkFlowSRTStatisticsRequest) (response *DescribeStreamLinkFlowSRTStatisticsResponse, err error) {
    if request == nil {
        request = NewDescribeStreamLinkFlowSRTStatisticsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "DescribeStreamLinkFlowSRTStatistics")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeStreamLinkFlowSRTStatistics require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeStreamLinkFlowSRTStatisticsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeStreamLinkFlowStatisticsRequest() (request *DescribeStreamLinkFlowStatisticsRequest) {
    request = &DescribeStreamLinkFlowStatisticsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "DescribeStreamLinkFlowStatistics")
    
    
    return
}

func NewDescribeStreamLinkFlowStatisticsResponse() (response *DescribeStreamLinkFlowStatisticsResponse) {
    response = &DescribeStreamLinkFlowStatisticsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeStreamLinkFlowStatistics
// 查询媒体传输流的媒体质量数据。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ENDTIME = "InvalidParameter.EndTime"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_INPUTOUTPUTID = "InvalidParameter.InputOutputId"
//  INVALIDPARAMETER_NOTFOUND = "InvalidParameter.NotFound"
//  INVALIDPARAMETER_PERIOD = "InvalidParameter.Period"
//  INVALIDPARAMETER_PIPELINE = "InvalidParameter.Pipeline"
//  INVALIDPARAMETER_STARTTIME = "InvalidParameter.StartTime"
//  INVALIDPARAMETER_TYPE = "InvalidParameter.Type"
func (c *Client) DescribeStreamLinkFlowStatistics(request *DescribeStreamLinkFlowStatisticsRequest) (response *DescribeStreamLinkFlowStatisticsResponse, err error) {
    return c.DescribeStreamLinkFlowStatisticsWithContext(context.Background(), request)
}

// DescribeStreamLinkFlowStatistics
// 查询媒体传输流的媒体质量数据。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ENDTIME = "InvalidParameter.EndTime"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_INPUTOUTPUTID = "InvalidParameter.InputOutputId"
//  INVALIDPARAMETER_NOTFOUND = "InvalidParameter.NotFound"
//  INVALIDPARAMETER_PERIOD = "InvalidParameter.Period"
//  INVALIDPARAMETER_PIPELINE = "InvalidParameter.Pipeline"
//  INVALIDPARAMETER_STARTTIME = "InvalidParameter.StartTime"
//  INVALIDPARAMETER_TYPE = "InvalidParameter.Type"
func (c *Client) DescribeStreamLinkFlowStatisticsWithContext(ctx context.Context, request *DescribeStreamLinkFlowStatisticsRequest) (response *DescribeStreamLinkFlowStatisticsResponse, err error) {
    if request == nil {
        request = NewDescribeStreamLinkFlowStatisticsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "DescribeStreamLinkFlowStatistics")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeStreamLinkFlowStatistics require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeStreamLinkFlowStatisticsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeStreamLinkFlowsRequest() (request *DescribeStreamLinkFlowsRequest) {
    request = &DescribeStreamLinkFlowsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "DescribeStreamLinkFlows")
    
    
    return
}

func NewDescribeStreamLinkFlowsResponse() (response *DescribeStreamLinkFlowsResponse) {
    response = &DescribeStreamLinkFlowsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeStreamLinkFlows
// 批量查询媒体输入流的配置信息。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_PAGENUM = "InvalidParameter.PageNum"
//  INVALIDPARAMETER_PAGESIZE = "InvalidParameter.PageSize"
func (c *Client) DescribeStreamLinkFlows(request *DescribeStreamLinkFlowsRequest) (response *DescribeStreamLinkFlowsResponse, err error) {
    return c.DescribeStreamLinkFlowsWithContext(context.Background(), request)
}

// DescribeStreamLinkFlows
// 批量查询媒体输入流的配置信息。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_PAGENUM = "InvalidParameter.PageNum"
//  INVALIDPARAMETER_PAGESIZE = "InvalidParameter.PageSize"
func (c *Client) DescribeStreamLinkFlowsWithContext(ctx context.Context, request *DescribeStreamLinkFlowsRequest) (response *DescribeStreamLinkFlowsResponse, err error) {
    if request == nil {
        request = NewDescribeStreamLinkFlowsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "DescribeStreamLinkFlows")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeStreamLinkFlows require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeStreamLinkFlowsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeStreamLinkRegionsRequest() (request *DescribeStreamLinkRegionsRequest) {
    request = &DescribeStreamLinkRegionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "DescribeStreamLinkRegions")
    
    
    return
}

func NewDescribeStreamLinkRegionsResponse() (response *DescribeStreamLinkRegionsResponse) {
    response = &DescribeStreamLinkRegionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeStreamLinkRegions
// 查询媒体传输所有地区。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeStreamLinkRegions(request *DescribeStreamLinkRegionsRequest) (response *DescribeStreamLinkRegionsResponse, err error) {
    return c.DescribeStreamLinkRegionsWithContext(context.Background(), request)
}

// DescribeStreamLinkRegions
// 查询媒体传输所有地区。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeStreamLinkRegionsWithContext(ctx context.Context, request *DescribeStreamLinkRegionsRequest) (response *DescribeStreamLinkRegionsResponse, err error) {
    if request == nil {
        request = NewDescribeStreamLinkRegionsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "DescribeStreamLinkRegions")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeStreamLinkRegions require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeStreamLinkRegionsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeStreamLinkSecurityGroupsRequest() (request *DescribeStreamLinkSecurityGroupsRequest) {
    request = &DescribeStreamLinkSecurityGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "DescribeStreamLinkSecurityGroups")
    
    
    return
}

func NewDescribeStreamLinkSecurityGroupsResponse() (response *DescribeStreamLinkSecurityGroupsResponse) {
    response = &DescribeStreamLinkSecurityGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeStreamLinkSecurityGroups
// 批量查询安全组信息。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeStreamLinkSecurityGroups(request *DescribeStreamLinkSecurityGroupsRequest) (response *DescribeStreamLinkSecurityGroupsResponse, err error) {
    return c.DescribeStreamLinkSecurityGroupsWithContext(context.Background(), request)
}

// DescribeStreamLinkSecurityGroups
// 批量查询安全组信息。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeStreamLinkSecurityGroupsWithContext(ctx context.Context, request *DescribeStreamLinkSecurityGroupsRequest) (response *DescribeStreamLinkSecurityGroupsResponse, err error) {
    if request == nil {
        request = NewDescribeStreamLinkSecurityGroupsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "DescribeStreamLinkSecurityGroups")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeStreamLinkSecurityGroups require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeStreamLinkSecurityGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTaskDetailRequest() (request *DescribeTaskDetailRequest) {
    request = &DescribeTaskDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "DescribeTaskDetail")
    
    
    return
}

func NewDescribeTaskDetailResponse() (response *DescribeTaskDetailResponse) {
    response = &DescribeTaskDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTaskDetail
// 通过任务 ID 查询任务的执行状态和结果的详细信息（最多可以查询7天之内提交的任务）。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeTaskDetail(request *DescribeTaskDetailRequest) (response *DescribeTaskDetailResponse, err error) {
    return c.DescribeTaskDetailWithContext(context.Background(), request)
}

// DescribeTaskDetail
// 通过任务 ID 查询任务的执行状态和结果的详细信息（最多可以查询7天之内提交的任务）。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeTaskDetailWithContext(ctx context.Context, request *DescribeTaskDetailRequest) (response *DescribeTaskDetailResponse, err error) {
    if request == nil {
        request = NewDescribeTaskDetailRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "DescribeTaskDetail")
    
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
    
    request.Init().WithApiInfo("mps", APIVersion, "DescribeTasks")
    
    
    return
}

func NewDescribeTasksResponse() (response *DescribeTasksResponse) {
    response = &DescribeTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTasks
// * 该接口用于查询任务列表；
//
// * 当列表数据比较多时，单次接口调用无法拉取整个列表，可通过 ScrollToken 参数，分批拉取；
//
// * 只能查询到最近七天（168小时）内的任务。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LIMIT = "InvalidParameterValue.Limit"
func (c *Client) DescribeTasks(request *DescribeTasksRequest) (response *DescribeTasksResponse, err error) {
    return c.DescribeTasksWithContext(context.Background(), request)
}

// DescribeTasks
// * 该接口用于查询任务列表；
//
// * 当列表数据比较多时，单次接口调用无法拉取整个列表，可通过 ScrollToken 参数，分批拉取；
//
// * 只能查询到最近七天（168小时）内的任务。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LIMIT = "InvalidParameterValue.Limit"
func (c *Client) DescribeTasksWithContext(ctx context.Context, request *DescribeTasksRequest) (response *DescribeTasksResponse, err error) {
    if request == nil {
        request = NewDescribeTasksRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "DescribeTasks")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTasks require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTasksResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTranscodeTemplatesRequest() (request *DescribeTranscodeTemplatesRequest) {
    request = &DescribeTranscodeTemplatesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "DescribeTranscodeTemplates")
    
    
    return
}

func NewDescribeTranscodeTemplatesResponse() (response *DescribeTranscodeTemplatesResponse) {
    response = &DescribeTranscodeTemplatesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTranscodeTemplates
// 根据转码模板唯一标识，获取转码模板详情列表。返回结果包含符合条件的所有用户自定义模板及[系统预置转码模板](https://cloud.tencent.com/document/product/266/33476#.E9.A2.84.E7.BD.AE.E8.BD.AC.E7.A0.81.E6.A8.A1.E6.9D.BF)。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ACCESSDBERROR = "InternalError.AccessDBError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_CONTAINERTYPE = "InvalidParameterValue.ContainerType"
//  INVALIDPARAMETERVALUE_DEFINITIONS = "InvalidParameterValue.Definitions"
//  INVALIDPARAMETERVALUE_LIMIT = "InvalidParameterValue.Limit"
//  INVALIDPARAMETERVALUE_TEHDTYPE = "InvalidParameterValue.TEHDType"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) DescribeTranscodeTemplates(request *DescribeTranscodeTemplatesRequest) (response *DescribeTranscodeTemplatesResponse, err error) {
    return c.DescribeTranscodeTemplatesWithContext(context.Background(), request)
}

// DescribeTranscodeTemplates
// 根据转码模板唯一标识，获取转码模板详情列表。返回结果包含符合条件的所有用户自定义模板及[系统预置转码模板](https://cloud.tencent.com/document/product/266/33476#.E9.A2.84.E7.BD.AE.E8.BD.AC.E7.A0.81.E6.A8.A1.E6.9D.BF)。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ACCESSDBERROR = "InternalError.AccessDBError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_CONTAINERTYPE = "InvalidParameterValue.ContainerType"
//  INVALIDPARAMETERVALUE_DEFINITIONS = "InvalidParameterValue.Definitions"
//  INVALIDPARAMETERVALUE_LIMIT = "InvalidParameterValue.Limit"
//  INVALIDPARAMETERVALUE_TEHDTYPE = "InvalidParameterValue.TEHDType"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) DescribeTranscodeTemplatesWithContext(ctx context.Context, request *DescribeTranscodeTemplatesRequest) (response *DescribeTranscodeTemplatesResponse, err error) {
    if request == nil {
        request = NewDescribeTranscodeTemplatesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "DescribeTranscodeTemplates")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTranscodeTemplates require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTranscodeTemplatesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUsageDataRequest() (request *DescribeUsageDataRequest) {
    request = &DescribeUsageDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "DescribeUsageData")
    
    
    return
}

func NewDescribeUsageDataResponse() (response *DescribeUsageDataResponse) {
    response = &DescribeUsageDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeUsageData
// 该接口返回查询时间范围内每天使用的媒体处理用量信息。
//
//    1. 可以查询最近365天内的媒体处理统计数据。
//
//    2. 查询时间跨度不超过90天。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  FAILEDOPERATION_NETWORKERROR = "FailedOperation.NetWorkError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_SERVICE = "InvalidParameterValue.Service"
func (c *Client) DescribeUsageData(request *DescribeUsageDataRequest) (response *DescribeUsageDataResponse, err error) {
    return c.DescribeUsageDataWithContext(context.Background(), request)
}

// DescribeUsageData
// 该接口返回查询时间范围内每天使用的媒体处理用量信息。
//
//    1. 可以查询最近365天内的媒体处理统计数据。
//
//    2. 查询时间跨度不超过90天。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  FAILEDOPERATION_NETWORKERROR = "FailedOperation.NetWorkError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_SERVICE = "InvalidParameterValue.Service"
func (c *Client) DescribeUsageDataWithContext(ctx context.Context, request *DescribeUsageDataRequest) (response *DescribeUsageDataResponse, err error) {
    if request == nil {
        request = NewDescribeUsageDataRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "DescribeUsageData")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUsageData require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeUsageDataResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVideoDatabaseEntryTaskDetailRequest() (request *DescribeVideoDatabaseEntryTaskDetailRequest) {
    request = &DescribeVideoDatabaseEntryTaskDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "DescribeVideoDatabaseEntryTaskDetail")
    
    
    return
}

func NewDescribeVideoDatabaseEntryTaskDetailResponse() (response *DescribeVideoDatabaseEntryTaskDetailResponse) {
    response = &DescribeVideoDatabaseEntryTaskDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeVideoDatabaseEntryTaskDetail
// 根据任务ID查询视频入库任务的状态。
//
// 可能返回的错误码:
//  FAILEDOPERATION_TASKNOTFOUND = "FailedOperation.TaskNotFound"
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeVideoDatabaseEntryTaskDetail(request *DescribeVideoDatabaseEntryTaskDetailRequest) (response *DescribeVideoDatabaseEntryTaskDetailResponse, err error) {
    return c.DescribeVideoDatabaseEntryTaskDetailWithContext(context.Background(), request)
}

// DescribeVideoDatabaseEntryTaskDetail
// 根据任务ID查询视频入库任务的状态。
//
// 可能返回的错误码:
//  FAILEDOPERATION_TASKNOTFOUND = "FailedOperation.TaskNotFound"
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeVideoDatabaseEntryTaskDetailWithContext(ctx context.Context, request *DescribeVideoDatabaseEntryTaskDetailRequest) (response *DescribeVideoDatabaseEntryTaskDetailResponse, err error) {
    if request == nil {
        request = NewDescribeVideoDatabaseEntryTaskDetailRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "DescribeVideoDatabaseEntryTaskDetail")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVideoDatabaseEntryTaskDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVideoDatabaseEntryTaskDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVideoSearchTaskDetailRequest() (request *DescribeVideoSearchTaskDetailRequest) {
    request = &DescribeVideoSearchTaskDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "DescribeVideoSearchTaskDetail")
    
    
    return
}

func NewDescribeVideoSearchTaskDetailResponse() (response *DescribeVideoSearchTaskDetailResponse) {
    response = &DescribeVideoSearchTaskDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeVideoSearchTaskDetail
// 根据任务ID查询视频检索任务的状态。
//
// 可能返回的错误码:
//  FAILEDOPERATION_TASKNOTFOUND = "FailedOperation.TaskNotFound"
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeVideoSearchTaskDetail(request *DescribeVideoSearchTaskDetailRequest) (response *DescribeVideoSearchTaskDetailResponse, err error) {
    return c.DescribeVideoSearchTaskDetailWithContext(context.Background(), request)
}

// DescribeVideoSearchTaskDetail
// 根据任务ID查询视频检索任务的状态。
//
// 可能返回的错误码:
//  FAILEDOPERATION_TASKNOTFOUND = "FailedOperation.TaskNotFound"
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeVideoSearchTaskDetailWithContext(ctx context.Context, request *DescribeVideoSearchTaskDetailRequest) (response *DescribeVideoSearchTaskDetailResponse, err error) {
    if request == nil {
        request = NewDescribeVideoSearchTaskDetailRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "DescribeVideoSearchTaskDetail")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVideoSearchTaskDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVideoSearchTaskDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeWatermarkTemplatesRequest() (request *DescribeWatermarkTemplatesRequest) {
    request = &DescribeWatermarkTemplatesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "DescribeWatermarkTemplates")
    
    
    return
}

func NewDescribeWatermarkTemplatesResponse() (response *DescribeWatermarkTemplatesResponse) {
    response = &DescribeWatermarkTemplatesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeWatermarkTemplates
// 查询用户自定义水印模板，支持根据条件，分页查询。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_DEFINITIONS = "InvalidParameterValue.Definitions"
//  INVALIDPARAMETERVALUE_LIMIT = "InvalidParameterValue.Limit"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) DescribeWatermarkTemplates(request *DescribeWatermarkTemplatesRequest) (response *DescribeWatermarkTemplatesResponse, err error) {
    return c.DescribeWatermarkTemplatesWithContext(context.Background(), request)
}

// DescribeWatermarkTemplates
// 查询用户自定义水印模板，支持根据条件，分页查询。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_DEFINITIONS = "InvalidParameterValue.Definitions"
//  INVALIDPARAMETERVALUE_LIMIT = "InvalidParameterValue.Limit"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) DescribeWatermarkTemplatesWithContext(ctx context.Context, request *DescribeWatermarkTemplatesRequest) (response *DescribeWatermarkTemplatesResponse, err error) {
    if request == nil {
        request = NewDescribeWatermarkTemplatesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "DescribeWatermarkTemplates")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeWatermarkTemplates require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeWatermarkTemplatesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeWordSamplesRequest() (request *DescribeWordSamplesRequest) {
    request = &DescribeWordSamplesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "DescribeWordSamples")
    
    
    return
}

func NewDescribeWordSamplesResponse() (response *DescribeWordSamplesResponse) {
    response = &DescribeWordSamplesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeWordSamples
// 该接口用于根据应用场景、关键词、标签，分页查询关键词样本信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeWordSamples(request *DescribeWordSamplesRequest) (response *DescribeWordSamplesResponse, err error) {
    return c.DescribeWordSamplesWithContext(context.Background(), request)
}

// DescribeWordSamples
// 该接口用于根据应用场景、关键词、标签，分页查询关键词样本信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeWordSamplesWithContext(ctx context.Context, request *DescribeWordSamplesRequest) (response *DescribeWordSamplesResponse, err error) {
    if request == nil {
        request = NewDescribeWordSamplesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "DescribeWordSamples")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeWordSamples require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeWordSamplesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeWorkflowsRequest() (request *DescribeWorkflowsRequest) {
    request = &DescribeWorkflowsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "DescribeWorkflows")
    
    
    return
}

func NewDescribeWorkflowsResponse() (response *DescribeWorkflowsResponse) {
    response = &DescribeWorkflowsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeWorkflows
// 根据工作流 ID，获取工作流详情列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeWorkflows(request *DescribeWorkflowsRequest) (response *DescribeWorkflowsResponse, err error) {
    return c.DescribeWorkflowsWithContext(context.Background(), request)
}

// DescribeWorkflows
// 根据工作流 ID，获取工作流详情列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeWorkflowsWithContext(ctx context.Context, request *DescribeWorkflowsRequest) (response *DescribeWorkflowsResponse, err error) {
    if request == nil {
        request = NewDescribeWorkflowsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "DescribeWorkflows")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeWorkflows require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeWorkflowsResponse()
    err = c.Send(request, response)
    return
}

func NewDisableScheduleRequest() (request *DisableScheduleRequest) {
    request = &DisableScheduleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "DisableSchedule")
    
    
    return
}

func NewDisableScheduleResponse() (response *DisableScheduleResponse) {
    response = &DisableScheduleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DisableSchedule
// 禁用自动化触发编排任务。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_BUCKETNOTIFYALREADYEXIST = "FailedOperation.BucketNotifyAlreadyExist"
//  FAILEDOPERATION_COSSTATUSINAVLID = "FailedOperation.CosStatusInavlid"
//  FAILEDOPERATION_GETSOURCENOTIFY = "FailedOperation.GetSourceNotify"
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  FAILEDOPERATION_INVALIDUSER = "FailedOperation.InvalidUser"
//  FAILEDOPERATION_SETSOURCENOTIFY = "FailedOperation.SetSourceNotify"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ACCESSDBERROR = "InternalError.AccessDBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_COSBUCKETNAMEINVALID = "ResourceNotFound.CosBucketNameInvalid"
//  RESOURCENOTFOUND_COSBUCKETNOTEXIST = "ResourceNotFound.CosBucketNotExist"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) DisableSchedule(request *DisableScheduleRequest) (response *DisableScheduleResponse, err error) {
    return c.DisableScheduleWithContext(context.Background(), request)
}

// DisableSchedule
// 禁用自动化触发编排任务。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_BUCKETNOTIFYALREADYEXIST = "FailedOperation.BucketNotifyAlreadyExist"
//  FAILEDOPERATION_COSSTATUSINAVLID = "FailedOperation.CosStatusInavlid"
//  FAILEDOPERATION_GETSOURCENOTIFY = "FailedOperation.GetSourceNotify"
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  FAILEDOPERATION_INVALIDUSER = "FailedOperation.InvalidUser"
//  FAILEDOPERATION_SETSOURCENOTIFY = "FailedOperation.SetSourceNotify"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ACCESSDBERROR = "InternalError.AccessDBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_COSBUCKETNAMEINVALID = "ResourceNotFound.CosBucketNameInvalid"
//  RESOURCENOTFOUND_COSBUCKETNOTEXIST = "ResourceNotFound.CosBucketNotExist"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) DisableScheduleWithContext(ctx context.Context, request *DisableScheduleRequest) (response *DisableScheduleResponse, err error) {
    if request == nil {
        request = NewDisableScheduleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "DisableSchedule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DisableSchedule require credential")
    }

    request.SetContext(ctx)
    
    response = NewDisableScheduleResponse()
    err = c.Send(request, response)
    return
}

func NewDisableWorkflowRequest() (request *DisableWorkflowRequest) {
    request = &DisableWorkflowRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "DisableWorkflow")
    
    
    return
}

func NewDisableWorkflowResponse() (response *DisableWorkflowResponse) {
    response = &DisableWorkflowResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DisableWorkflow
// 禁用工作流。
//
// 可能返回的错误码:
//  FAILEDOPERATION_BUCKETNOTIFYALREADYEXIST = "FailedOperation.BucketNotifyAlreadyExist"
//  FAILEDOPERATION_COSSTATUSINAVLID = "FailedOperation.CosStatusInavlid"
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_COSBUCKETNAMEINVALID = "ResourceNotFound.CosBucketNameInvalid"
//  RESOURCENOTFOUND_COSBUCKETNOTEXIST = "ResourceNotFound.CosBucketNotExist"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) DisableWorkflow(request *DisableWorkflowRequest) (response *DisableWorkflowResponse, err error) {
    return c.DisableWorkflowWithContext(context.Background(), request)
}

// DisableWorkflow
// 禁用工作流。
//
// 可能返回的错误码:
//  FAILEDOPERATION_BUCKETNOTIFYALREADYEXIST = "FailedOperation.BucketNotifyAlreadyExist"
//  FAILEDOPERATION_COSSTATUSINAVLID = "FailedOperation.CosStatusInavlid"
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_COSBUCKETNAMEINVALID = "ResourceNotFound.CosBucketNameInvalid"
//  RESOURCENOTFOUND_COSBUCKETNOTEXIST = "ResourceNotFound.CosBucketNotExist"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) DisableWorkflowWithContext(ctx context.Context, request *DisableWorkflowRequest) (response *DisableWorkflowResponse, err error) {
    if request == nil {
        request = NewDisableWorkflowRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "DisableWorkflow")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DisableWorkflow require credential")
    }

    request.SetContext(ctx)
    
    response = NewDisableWorkflowResponse()
    err = c.Send(request, response)
    return
}

func NewDisassociateSecurityGroupRequest() (request *DisassociateSecurityGroupRequest) {
    request = &DisassociateSecurityGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "DisassociateSecurityGroup")
    
    
    return
}

func NewDisassociateSecurityGroupResponse() (response *DisassociateSecurityGroupResponse) {
    response = &DisassociateSecurityGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DisassociateSecurityGroup
// 批量解绑安全组下面关联的输入输出。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DisassociateSecurityGroup(request *DisassociateSecurityGroupRequest) (response *DisassociateSecurityGroupResponse, err error) {
    return c.DisassociateSecurityGroupWithContext(context.Background(), request)
}

// DisassociateSecurityGroup
// 批量解绑安全组下面关联的输入输出。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DisassociateSecurityGroupWithContext(ctx context.Context, request *DisassociateSecurityGroupRequest) (response *DisassociateSecurityGroupResponse, err error) {
    if request == nil {
        request = NewDisassociateSecurityGroupRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "DisassociateSecurityGroup")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DisassociateSecurityGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewDisassociateSecurityGroupResponse()
    err = c.Send(request, response)
    return
}

func NewEditMediaRequest() (request *EditMediaRequest) {
    request = &EditMediaRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "EditMedia")
    
    
    return
}

func NewEditMediaResponse() (response *EditMediaResponse) {
    response = &EditMediaResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// EditMedia
// 对视频进行编辑，生成一个新的视频。编辑的功能包括：
//
//  
//
// 
//
// 一、**剪辑任务**：简单的视频剪辑，如剪辑、拼接等
//
// 1. 对一个文件进行剪辑，生成一个新的视频；
//
// 2. 对多个文件进行拼接，生成一个新的视频；
//
// 3. 对多个文件进行剪辑，然后再拼接，生成一个新的视频。
//
// 
//
// 二、**合成任务**：通过接口描述信息，合成一个新的视频。
//
// 1. 多轨道（视频、音频、字幕）、多类型元素（视频、图片、音频、文字、空）
//
// 2. 图像级别：贴图、缩放、任意角度旋转、镜像等
//
// 3. 音频级别：音量控制、淡入淡出、混音等
//
// 4. 视频级别：转场、倍数播放、拼接、剪切、字幕、画中画、音画分离、出入场动效等
//
// 可能返回的错误码:
//  FAILEDOPERATION_GENERATERESOURCE = "FailedOperation.GenerateResource"
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) EditMedia(request *EditMediaRequest) (response *EditMediaResponse, err error) {
    return c.EditMediaWithContext(context.Background(), request)
}

// EditMedia
// 对视频进行编辑，生成一个新的视频。编辑的功能包括：
//
//  
//
// 
//
// 一、**剪辑任务**：简单的视频剪辑，如剪辑、拼接等
//
// 1. 对一个文件进行剪辑，生成一个新的视频；
//
// 2. 对多个文件进行拼接，生成一个新的视频；
//
// 3. 对多个文件进行剪辑，然后再拼接，生成一个新的视频。
//
// 
//
// 二、**合成任务**：通过接口描述信息，合成一个新的视频。
//
// 1. 多轨道（视频、音频、字幕）、多类型元素（视频、图片、音频、文字、空）
//
// 2. 图像级别：贴图、缩放、任意角度旋转、镜像等
//
// 3. 音频级别：音量控制、淡入淡出、混音等
//
// 4. 视频级别：转场、倍数播放、拼接、剪切、字幕、画中画、音画分离、出入场动效等
//
// 可能返回的错误码:
//  FAILEDOPERATION_GENERATERESOURCE = "FailedOperation.GenerateResource"
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) EditMediaWithContext(ctx context.Context, request *EditMediaRequest) (response *EditMediaResponse, err error) {
    if request == nil {
        request = NewEditMediaRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "EditMedia")
    
    if c.GetCredential() == nil {
        return nil, errors.New("EditMedia require credential")
    }

    request.SetContext(ctx)
    
    response = NewEditMediaResponse()
    err = c.Send(request, response)
    return
}

func NewEnableScheduleRequest() (request *EnableScheduleRequest) {
    request = &EnableScheduleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "EnableSchedule")
    
    
    return
}

func NewEnableScheduleResponse() (response *EnableScheduleResponse) {
    response = &EnableScheduleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// EnableSchedule
// 启用自动化触发编排任务。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_BUCKETNOTIFYALREADYEXIST = "FailedOperation.BucketNotifyAlreadyExist"
//  FAILEDOPERATION_COSSTATUSINAVLID = "FailedOperation.CosStatusInavlid"
//  FAILEDOPERATION_GENERATERESOURCE = "FailedOperation.GenerateResource"
//  FAILEDOPERATION_GETSOURCENOTIFY = "FailedOperation.GetSourceNotify"
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  FAILEDOPERATION_INVALIDUSER = "FailedOperation.InvalidUser"
//  FAILEDOPERATION_SETSOURCENOTIFY = "FailedOperation.SetSourceNotify"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ACCESSDBERROR = "InternalError.AccessDBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_COSBUCKETNAMEINVALID = "ResourceNotFound.CosBucketNameInvalid"
//  RESOURCENOTFOUND_COSBUCKETNOTEXIST = "ResourceNotFound.CosBucketNotExist"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) EnableSchedule(request *EnableScheduleRequest) (response *EnableScheduleResponse, err error) {
    return c.EnableScheduleWithContext(context.Background(), request)
}

// EnableSchedule
// 启用自动化触发编排任务。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_BUCKETNOTIFYALREADYEXIST = "FailedOperation.BucketNotifyAlreadyExist"
//  FAILEDOPERATION_COSSTATUSINAVLID = "FailedOperation.CosStatusInavlid"
//  FAILEDOPERATION_GENERATERESOURCE = "FailedOperation.GenerateResource"
//  FAILEDOPERATION_GETSOURCENOTIFY = "FailedOperation.GetSourceNotify"
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  FAILEDOPERATION_INVALIDUSER = "FailedOperation.InvalidUser"
//  FAILEDOPERATION_SETSOURCENOTIFY = "FailedOperation.SetSourceNotify"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ACCESSDBERROR = "InternalError.AccessDBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_COSBUCKETNAMEINVALID = "ResourceNotFound.CosBucketNameInvalid"
//  RESOURCENOTFOUND_COSBUCKETNOTEXIST = "ResourceNotFound.CosBucketNotExist"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) EnableScheduleWithContext(ctx context.Context, request *EnableScheduleRequest) (response *EnableScheduleResponse, err error) {
    if request == nil {
        request = NewEnableScheduleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "EnableSchedule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("EnableSchedule require credential")
    }

    request.SetContext(ctx)
    
    response = NewEnableScheduleResponse()
    err = c.Send(request, response)
    return
}

func NewEnableWorkflowRequest() (request *EnableWorkflowRequest) {
    request = &EnableWorkflowRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "EnableWorkflow")
    
    
    return
}

func NewEnableWorkflowResponse() (response *EnableWorkflowResponse) {
    response = &EnableWorkflowResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// EnableWorkflow
// 启用工作流。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_BUCKETNOTIFYALREADYEXIST = "FailedOperation.BucketNotifyAlreadyExist"
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_COSBUCKETNAMEINVALID = "ResourceNotFound.CosBucketNameInvalid"
//  RESOURCENOTFOUND_COSBUCKETNOTEXIST = "ResourceNotFound.CosBucketNotExist"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) EnableWorkflow(request *EnableWorkflowRequest) (response *EnableWorkflowResponse, err error) {
    return c.EnableWorkflowWithContext(context.Background(), request)
}

// EnableWorkflow
// 启用工作流。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_BUCKETNOTIFYALREADYEXIST = "FailedOperation.BucketNotifyAlreadyExist"
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_COSBUCKETNAMEINVALID = "ResourceNotFound.CosBucketNameInvalid"
//  RESOURCENOTFOUND_COSBUCKETNOTEXIST = "ResourceNotFound.CosBucketNotExist"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) EnableWorkflowWithContext(ctx context.Context, request *EnableWorkflowRequest) (response *EnableWorkflowResponse, err error) {
    if request == nil {
        request = NewEnableWorkflowRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "EnableWorkflow")
    
    if c.GetCredential() == nil {
        return nil, errors.New("EnableWorkflow require credential")
    }

    request.SetContext(ctx)
    
    response = NewEnableWorkflowResponse()
    err = c.Send(request, response)
    return
}

func NewExecuteFunctionRequest() (request *ExecuteFunctionRequest) {
    request = &ExecuteFunctionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "ExecuteFunction")
    
    
    return
}

func NewExecuteFunctionResponse() (response *ExecuteFunctionResponse) {
    response = &ExecuteFunctionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ExecuteFunction
// 本接口仅用于定制开发的特殊场景，除非云媒体处理客服人员主动告知您需要使用本接口，其它情况请勿调用。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  FAILEDOPERATION_INVALIDUSER = "FailedOperation.InvalidUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_FUNCTIONARG = "InvalidParameterValue.FunctionArg"
//  INVALIDPARAMETERVALUE_FUNCTIONNAME = "InvalidParameterValue.FunctionName"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ExecuteFunction(request *ExecuteFunctionRequest) (response *ExecuteFunctionResponse, err error) {
    return c.ExecuteFunctionWithContext(context.Background(), request)
}

// ExecuteFunction
// 本接口仅用于定制开发的特殊场景，除非云媒体处理客服人员主动告知您需要使用本接口，其它情况请勿调用。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  FAILEDOPERATION_INVALIDUSER = "FailedOperation.InvalidUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_FUNCTIONARG = "InvalidParameterValue.FunctionArg"
//  INVALIDPARAMETERVALUE_FUNCTIONNAME = "InvalidParameterValue.FunctionName"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ExecuteFunctionWithContext(ctx context.Context, request *ExecuteFunctionRequest) (response *ExecuteFunctionResponse, err error) {
    if request == nil {
        request = NewExecuteFunctionRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "ExecuteFunction")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ExecuteFunction require credential")
    }

    request.SetContext(ctx)
    
    response = NewExecuteFunctionResponse()
    err = c.Send(request, response)
    return
}

func NewExtractBlindWatermarkRequest() (request *ExtractBlindWatermarkRequest) {
    request = &ExtractBlindWatermarkRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "ExtractBlindWatermark")
    
    
    return
}

func NewExtractBlindWatermarkResponse() (response *ExtractBlindWatermarkResponse) {
    response = &ExtractBlindWatermarkResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ExtractBlindWatermark
// 用于发起提取视频数字水印任务，提取结果可以通过DescribeTaskDetail查询。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_SRCFILE = "InvalidParameterValue.SrcFile"
func (c *Client) ExtractBlindWatermark(request *ExtractBlindWatermarkRequest) (response *ExtractBlindWatermarkResponse, err error) {
    return c.ExtractBlindWatermarkWithContext(context.Background(), request)
}

// ExtractBlindWatermark
// 用于发起提取视频数字水印任务，提取结果可以通过DescribeTaskDetail查询。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_SRCFILE = "InvalidParameterValue.SrcFile"
func (c *Client) ExtractBlindWatermarkWithContext(ctx context.Context, request *ExtractBlindWatermarkRequest) (response *ExtractBlindWatermarkResponse, err error) {
    if request == nil {
        request = NewExtractBlindWatermarkRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "ExtractBlindWatermark")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ExtractBlindWatermark require credential")
    }

    request.SetContext(ctx)
    
    response = NewExtractBlindWatermarkResponse()
    err = c.Send(request, response)
    return
}

func NewManageTaskRequest() (request *ManageTaskRequest) {
    request = &ManageTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "ManageTask")
    
    
    return
}

func NewManageTaskResponse() (response *ManageTaskResponse) {
    response = &ManageTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ManageTask
// 对已发起的任务进行管理。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDOPERATIONTYPE = "InvalidParameterValue.InvalidOperationType"
//  INVALIDPARAMETERVALUE_NOTPROCESSINGTASK = "InvalidParameterValue.NotProcessingTask"
//  INVALIDPARAMETERVALUE_TASKID = "InvalidParameterValue.TaskId"
func (c *Client) ManageTask(request *ManageTaskRequest) (response *ManageTaskResponse, err error) {
    return c.ManageTaskWithContext(context.Background(), request)
}

// ManageTask
// 对已发起的任务进行管理。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDOPERATIONTYPE = "InvalidParameterValue.InvalidOperationType"
//  INVALIDPARAMETERVALUE_NOTPROCESSINGTASK = "InvalidParameterValue.NotProcessingTask"
//  INVALIDPARAMETERVALUE_TASKID = "InvalidParameterValue.TaskId"
func (c *Client) ManageTaskWithContext(ctx context.Context, request *ManageTaskRequest) (response *ManageTaskResponse, err error) {
    if request == nil {
        request = NewManageTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "ManageTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ManageTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewManageTaskResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAIAnalysisTemplateRequest() (request *ModifyAIAnalysisTemplateRequest) {
    request = &ModifyAIAnalysisTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "ModifyAIAnalysisTemplate")
    
    
    return
}

func NewModifyAIAnalysisTemplateResponse() (response *ModifyAIAnalysisTemplateResponse) {
    response = &ModifyAIAnalysisTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyAIAnalysisTemplate
// 修改用户自定义内容分析模板。
//
// 
//
// 注意：模板 ID 10000 以下的为系统预置模板，不允许修改。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_CLASSIFCATIONCONFIGURE = "InvalidParameterValue.ClassifcationConfigure"
//  INVALIDPARAMETERVALUE_COMMENT = "InvalidParameterValue.Comment"
//  INVALIDPARAMETERVALUE_COVERCONFIGURE = "InvalidParameterValue.CoverConfigure"
//  INVALIDPARAMETERVALUE_FRAMETAGCONFIGURE = "InvalidParameterValue.FrameTagConfigure"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  INVALIDPARAMETERVALUE_TAGCONFIGURE = "InvalidParameterValue.TagConfigure"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) ModifyAIAnalysisTemplate(request *ModifyAIAnalysisTemplateRequest) (response *ModifyAIAnalysisTemplateResponse, err error) {
    return c.ModifyAIAnalysisTemplateWithContext(context.Background(), request)
}

// ModifyAIAnalysisTemplate
// 修改用户自定义内容分析模板。
//
// 
//
// 注意：模板 ID 10000 以下的为系统预置模板，不允许修改。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_CLASSIFCATIONCONFIGURE = "InvalidParameterValue.ClassifcationConfigure"
//  INVALIDPARAMETERVALUE_COMMENT = "InvalidParameterValue.Comment"
//  INVALIDPARAMETERVALUE_COVERCONFIGURE = "InvalidParameterValue.CoverConfigure"
//  INVALIDPARAMETERVALUE_FRAMETAGCONFIGURE = "InvalidParameterValue.FrameTagConfigure"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  INVALIDPARAMETERVALUE_TAGCONFIGURE = "InvalidParameterValue.TagConfigure"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) ModifyAIAnalysisTemplateWithContext(ctx context.Context, request *ModifyAIAnalysisTemplateRequest) (response *ModifyAIAnalysisTemplateResponse, err error) {
    if request == nil {
        request = NewModifyAIAnalysisTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "ModifyAIAnalysisTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAIAnalysisTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAIAnalysisTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAIRecognitionTemplateRequest() (request *ModifyAIRecognitionTemplateRequest) {
    request = &ModifyAIRecognitionTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "ModifyAIRecognitionTemplate")
    
    
    return
}

func NewModifyAIRecognitionTemplateResponse() (response *ModifyAIRecognitionTemplateResponse) {
    response = &ModifyAIRecognitionTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyAIRecognitionTemplate
// 修改用户自定义内容识别模板。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_COMMENT = "InvalidParameterValue.Comment"
//  INVALIDPARAMETERVALUE_DEFAULTLIBRARYLABELSET = "InvalidParameterValue.DefaultLibraryLabelSet"
//  INVALIDPARAMETERVALUE_FACELIBRARY = "InvalidParameterValue.FaceLibrary"
//  INVALIDPARAMETERVALUE_FACESCORE = "InvalidParameterValue.FaceScore"
//  INVALIDPARAMETERVALUE_LABELSET = "InvalidParameterValue.LabelSet"
//  INVALIDPARAMETERVALUE_MODIFYDEFAULTTEMPLATE = "InvalidParameterValue.ModifyDefaultTemplate"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  INVALIDPARAMETERVALUE_OBJECTLIBRARY = "InvalidParameterValue.ObjectLibrary"
//  INVALIDPARAMETERVALUE_SOURCELANGUAGE = "InvalidParameterValue.SourceLanguage"
//  INVALIDPARAMETERVALUE_SUBTITLEFORMAT = "InvalidParameterValue.SubtitleFormat"
//  INVALIDPARAMETERVALUE_SWITCH = "InvalidParameterValue.Switch"
//  INVALIDPARAMETERVALUE_USERDEFINELIBRARYLABELSET = "InvalidParameterValue.UserDefineLibraryLabelSet"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) ModifyAIRecognitionTemplate(request *ModifyAIRecognitionTemplateRequest) (response *ModifyAIRecognitionTemplateResponse, err error) {
    return c.ModifyAIRecognitionTemplateWithContext(context.Background(), request)
}

// ModifyAIRecognitionTemplate
// 修改用户自定义内容识别模板。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_COMMENT = "InvalidParameterValue.Comment"
//  INVALIDPARAMETERVALUE_DEFAULTLIBRARYLABELSET = "InvalidParameterValue.DefaultLibraryLabelSet"
//  INVALIDPARAMETERVALUE_FACELIBRARY = "InvalidParameterValue.FaceLibrary"
//  INVALIDPARAMETERVALUE_FACESCORE = "InvalidParameterValue.FaceScore"
//  INVALIDPARAMETERVALUE_LABELSET = "InvalidParameterValue.LabelSet"
//  INVALIDPARAMETERVALUE_MODIFYDEFAULTTEMPLATE = "InvalidParameterValue.ModifyDefaultTemplate"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  INVALIDPARAMETERVALUE_OBJECTLIBRARY = "InvalidParameterValue.ObjectLibrary"
//  INVALIDPARAMETERVALUE_SOURCELANGUAGE = "InvalidParameterValue.SourceLanguage"
//  INVALIDPARAMETERVALUE_SUBTITLEFORMAT = "InvalidParameterValue.SubtitleFormat"
//  INVALIDPARAMETERVALUE_SWITCH = "InvalidParameterValue.Switch"
//  INVALIDPARAMETERVALUE_USERDEFINELIBRARYLABELSET = "InvalidParameterValue.UserDefineLibraryLabelSet"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) ModifyAIRecognitionTemplateWithContext(ctx context.Context, request *ModifyAIRecognitionTemplateRequest) (response *ModifyAIRecognitionTemplateResponse, err error) {
    if request == nil {
        request = NewModifyAIRecognitionTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "ModifyAIRecognitionTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAIRecognitionTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAIRecognitionTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAdaptiveDynamicStreamingTemplateRequest() (request *ModifyAdaptiveDynamicStreamingTemplateRequest) {
    request = &ModifyAdaptiveDynamicStreamingTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "ModifyAdaptiveDynamicStreamingTemplate")
    
    
    return
}

func NewModifyAdaptiveDynamicStreamingTemplateResponse() (response *ModifyAdaptiveDynamicStreamingTemplateResponse) {
    response = &ModifyAdaptiveDynamicStreamingTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyAdaptiveDynamicStreamingTemplate
// 修改转自适应码流模板
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_AUDIOBITRATE = "InvalidParameterValue.AudioBitrate"
//  INVALIDPARAMETERVALUE_AUDIOCHANNEL = "InvalidParameterValue.AudioChannel"
//  INVALIDPARAMETERVALUE_AUDIOCODEC = "InvalidParameterValue.AudioCodec"
//  INVALIDPARAMETERVALUE_AUDIOSAMPLERATE = "InvalidParameterValue.AudioSampleRate"
//  INVALIDPARAMETERVALUE_BITRATE = "InvalidParameterValue.Bitrate"
//  INVALIDPARAMETERVALUE_CODEC = "InvalidParameterValue.Codec"
//  INVALIDPARAMETERVALUE_COMMENT = "InvalidParameterValue.Comment"
//  INVALIDPARAMETERVALUE_DEFINITION = "InvalidParameterValue.Definition"
//  INVALIDPARAMETERVALUE_DISABLEHIGHERVIDEOBITRATE = "InvalidParameterValue.DisableHigherVideoBitrate"
//  INVALIDPARAMETERVALUE_DISABLEHIGHERVIDEORESOLUTION = "InvalidParameterValue.DisableHigherVideoResolution"
//  INVALIDPARAMETERVALUE_FILLTYPE = "InvalidParameterValue.FillType"
//  INVALIDPARAMETERVALUE_FORMAT = "InvalidParameterValue.Format"
//  INVALIDPARAMETERVALUE_FPS = "InvalidParameterValue.Fps"
//  INVALIDPARAMETERVALUE_GOP = "InvalidParameterValue.Gop"
//  INVALIDPARAMETERVALUE_HEIGHT = "InvalidParameterValue.Height"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  INVALIDPARAMETERVALUE_REMOVEAUDIO = "InvalidParameterValue.RemoveAudio"
//  INVALIDPARAMETERVALUE_REMOVEVIDEO = "InvalidParameterValue.RemoveVideo"
//  INVALIDPARAMETERVALUE_SOUNDSYSTEM = "InvalidParameterValue.SoundSystem"
//  INVALIDPARAMETERVALUE_VIDEOBITRATE = "InvalidParameterValue.VideoBitrate"
//  INVALIDPARAMETERVALUE_VIDEOCODEC = "InvalidParameterValue.VideoCodec"
//  INVALIDPARAMETERVALUE_WIDTH = "InvalidParameterValue.Width"
func (c *Client) ModifyAdaptiveDynamicStreamingTemplate(request *ModifyAdaptiveDynamicStreamingTemplateRequest) (response *ModifyAdaptiveDynamicStreamingTemplateResponse, err error) {
    return c.ModifyAdaptiveDynamicStreamingTemplateWithContext(context.Background(), request)
}

// ModifyAdaptiveDynamicStreamingTemplate
// 修改转自适应码流模板
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_AUDIOBITRATE = "InvalidParameterValue.AudioBitrate"
//  INVALIDPARAMETERVALUE_AUDIOCHANNEL = "InvalidParameterValue.AudioChannel"
//  INVALIDPARAMETERVALUE_AUDIOCODEC = "InvalidParameterValue.AudioCodec"
//  INVALIDPARAMETERVALUE_AUDIOSAMPLERATE = "InvalidParameterValue.AudioSampleRate"
//  INVALIDPARAMETERVALUE_BITRATE = "InvalidParameterValue.Bitrate"
//  INVALIDPARAMETERVALUE_CODEC = "InvalidParameterValue.Codec"
//  INVALIDPARAMETERVALUE_COMMENT = "InvalidParameterValue.Comment"
//  INVALIDPARAMETERVALUE_DEFINITION = "InvalidParameterValue.Definition"
//  INVALIDPARAMETERVALUE_DISABLEHIGHERVIDEOBITRATE = "InvalidParameterValue.DisableHigherVideoBitrate"
//  INVALIDPARAMETERVALUE_DISABLEHIGHERVIDEORESOLUTION = "InvalidParameterValue.DisableHigherVideoResolution"
//  INVALIDPARAMETERVALUE_FILLTYPE = "InvalidParameterValue.FillType"
//  INVALIDPARAMETERVALUE_FORMAT = "InvalidParameterValue.Format"
//  INVALIDPARAMETERVALUE_FPS = "InvalidParameterValue.Fps"
//  INVALIDPARAMETERVALUE_GOP = "InvalidParameterValue.Gop"
//  INVALIDPARAMETERVALUE_HEIGHT = "InvalidParameterValue.Height"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  INVALIDPARAMETERVALUE_REMOVEAUDIO = "InvalidParameterValue.RemoveAudio"
//  INVALIDPARAMETERVALUE_REMOVEVIDEO = "InvalidParameterValue.RemoveVideo"
//  INVALIDPARAMETERVALUE_SOUNDSYSTEM = "InvalidParameterValue.SoundSystem"
//  INVALIDPARAMETERVALUE_VIDEOBITRATE = "InvalidParameterValue.VideoBitrate"
//  INVALIDPARAMETERVALUE_VIDEOCODEC = "InvalidParameterValue.VideoCodec"
//  INVALIDPARAMETERVALUE_WIDTH = "InvalidParameterValue.Width"
func (c *Client) ModifyAdaptiveDynamicStreamingTemplateWithContext(ctx context.Context, request *ModifyAdaptiveDynamicStreamingTemplateRequest) (response *ModifyAdaptiveDynamicStreamingTemplateResponse, err error) {
    if request == nil {
        request = NewModifyAdaptiveDynamicStreamingTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "ModifyAdaptiveDynamicStreamingTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAdaptiveDynamicStreamingTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAdaptiveDynamicStreamingTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAnimatedGraphicsTemplateRequest() (request *ModifyAnimatedGraphicsTemplateRequest) {
    request = &ModifyAnimatedGraphicsTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "ModifyAnimatedGraphicsTemplate")
    
    
    return
}

func NewModifyAnimatedGraphicsTemplateResponse() (response *ModifyAnimatedGraphicsTemplateResponse) {
    response = &ModifyAnimatedGraphicsTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyAnimatedGraphicsTemplate
// 修改用户自定义转动图模板。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_COMMENT = "InvalidParameterValue.Comment"
//  INVALIDPARAMETERVALUE_FORMAT = "InvalidParameterValue.Format"
//  INVALIDPARAMETERVALUE_FORMATWEBPLACKWIDTHANDHEIGHT = "InvalidParameterValue.FormatWebpLackWidthAndHeight"
//  INVALIDPARAMETERVALUE_FORMATWEBPWIDTHANDHEIGHTBOTHZERO = "InvalidParameterValue.FormatWebpWidthAndHeightBothZero"
//  INVALIDPARAMETERVALUE_FPS = "InvalidParameterValue.Fps"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  INVALIDPARAMETERVALUE_QUALITY = "InvalidParameterValue.Quality"
//  INVALIDPARAMETERVALUE_RESOLUTION = "InvalidParameterValue.Resolution"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) ModifyAnimatedGraphicsTemplate(request *ModifyAnimatedGraphicsTemplateRequest) (response *ModifyAnimatedGraphicsTemplateResponse, err error) {
    return c.ModifyAnimatedGraphicsTemplateWithContext(context.Background(), request)
}

// ModifyAnimatedGraphicsTemplate
// 修改用户自定义转动图模板。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_COMMENT = "InvalidParameterValue.Comment"
//  INVALIDPARAMETERVALUE_FORMAT = "InvalidParameterValue.Format"
//  INVALIDPARAMETERVALUE_FORMATWEBPLACKWIDTHANDHEIGHT = "InvalidParameterValue.FormatWebpLackWidthAndHeight"
//  INVALIDPARAMETERVALUE_FORMATWEBPWIDTHANDHEIGHTBOTHZERO = "InvalidParameterValue.FormatWebpWidthAndHeightBothZero"
//  INVALIDPARAMETERVALUE_FPS = "InvalidParameterValue.Fps"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  INVALIDPARAMETERVALUE_QUALITY = "InvalidParameterValue.Quality"
//  INVALIDPARAMETERVALUE_RESOLUTION = "InvalidParameterValue.Resolution"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) ModifyAnimatedGraphicsTemplateWithContext(ctx context.Context, request *ModifyAnimatedGraphicsTemplateRequest) (response *ModifyAnimatedGraphicsTemplateResponse, err error) {
    if request == nil {
        request = NewModifyAnimatedGraphicsTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "ModifyAnimatedGraphicsTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAnimatedGraphicsTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAnimatedGraphicsTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAsrHotwordsRequest() (request *ModifyAsrHotwordsRequest) {
    request = &ModifyAsrHotwordsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "ModifyAsrHotwords")
    
    
    return
}

func NewModifyAsrHotwordsResponse() (response *ModifyAsrHotwordsResponse) {
    response = &ModifyAsrHotwordsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyAsrHotwords
// 智能字幕更新热词库接口
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ACCESSDBERROR = "InternalError.AccessDBError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_HOTWORDSNOTEXIST = "InvalidParameterValue.HotWordsNotExist"
func (c *Client) ModifyAsrHotwords(request *ModifyAsrHotwordsRequest) (response *ModifyAsrHotwordsResponse, err error) {
    return c.ModifyAsrHotwordsWithContext(context.Background(), request)
}

// ModifyAsrHotwords
// 智能字幕更新热词库接口
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ACCESSDBERROR = "InternalError.AccessDBError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_HOTWORDSNOTEXIST = "InvalidParameterValue.HotWordsNotExist"
func (c *Client) ModifyAsrHotwordsWithContext(ctx context.Context, request *ModifyAsrHotwordsRequest) (response *ModifyAsrHotwordsResponse, err error) {
    if request == nil {
        request = NewModifyAsrHotwordsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "ModifyAsrHotwords")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAsrHotwords require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAsrHotwordsResponse()
    err = c.Send(request, response)
    return
}

func NewModifyBlindWatermarkTemplateRequest() (request *ModifyBlindWatermarkTemplateRequest) {
    request = &ModifyBlindWatermarkTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "ModifyBlindWatermarkTemplate")
    
    
    return
}

func NewModifyBlindWatermarkTemplateResponse() (response *ModifyBlindWatermarkTemplateResponse) {
    response = &ModifyBlindWatermarkTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyBlindWatermarkTemplate
// 修改用户自定义数字水印模板，数字水印类型不允许修改。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_COMMENT = "InvalidParameterValue.Comment"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) ModifyBlindWatermarkTemplate(request *ModifyBlindWatermarkTemplateRequest) (response *ModifyBlindWatermarkTemplateResponse, err error) {
    return c.ModifyBlindWatermarkTemplateWithContext(context.Background(), request)
}

// ModifyBlindWatermarkTemplate
// 修改用户自定义数字水印模板，数字水印类型不允许修改。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_COMMENT = "InvalidParameterValue.Comment"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) ModifyBlindWatermarkTemplateWithContext(ctx context.Context, request *ModifyBlindWatermarkTemplateRequest) (response *ModifyBlindWatermarkTemplateResponse, err error) {
    if request == nil {
        request = NewModifyBlindWatermarkTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "ModifyBlindWatermarkTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyBlindWatermarkTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyBlindWatermarkTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewModifyContentReviewTemplateRequest() (request *ModifyContentReviewTemplateRequest) {
    request = &ModifyContentReviewTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "ModifyContentReviewTemplate")
    
    
    return
}

func NewModifyContentReviewTemplateResponse() (response *ModifyContentReviewTemplateResponse) {
    response = &ModifyContentReviewTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyContentReviewTemplate
// 修改用户自定义内容审核模板。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_BLOCKCONFIDENCE = "InvalidParameterValue.BlockConfidence"
//  INVALIDPARAMETERVALUE_COMMENT = "InvalidParameterValue.Comment"
//  INVALIDPARAMETERVALUE_LABELSET = "InvalidParameterValue.LabelSet"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  INVALIDPARAMETERVALUE_REVIEWCONFIDENCE = "InvalidParameterValue.ReviewConfidence"
//  INVALIDPARAMETERVALUE_SWITCH = "InvalidParameterValue.Switch"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) ModifyContentReviewTemplate(request *ModifyContentReviewTemplateRequest) (response *ModifyContentReviewTemplateResponse, err error) {
    return c.ModifyContentReviewTemplateWithContext(context.Background(), request)
}

// ModifyContentReviewTemplate
// 修改用户自定义内容审核模板。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_BLOCKCONFIDENCE = "InvalidParameterValue.BlockConfidence"
//  INVALIDPARAMETERVALUE_COMMENT = "InvalidParameterValue.Comment"
//  INVALIDPARAMETERVALUE_LABELSET = "InvalidParameterValue.LabelSet"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  INVALIDPARAMETERVALUE_REVIEWCONFIDENCE = "InvalidParameterValue.ReviewConfidence"
//  INVALIDPARAMETERVALUE_SWITCH = "InvalidParameterValue.Switch"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) ModifyContentReviewTemplateWithContext(ctx context.Context, request *ModifyContentReviewTemplateRequest) (response *ModifyContentReviewTemplateResponse, err error) {
    if request == nil {
        request = NewModifyContentReviewTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "ModifyContentReviewTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyContentReviewTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyContentReviewTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewModifyImageSpriteTemplateRequest() (request *ModifyImageSpriteTemplateRequest) {
    request = &ModifyImageSpriteTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "ModifyImageSpriteTemplate")
    
    
    return
}

func NewModifyImageSpriteTemplateResponse() (response *ModifyImageSpriteTemplateResponse) {
    response = &ModifyImageSpriteTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyImageSpriteTemplate
// 修改用户自定义雪碧图模板。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_COLUMNCOUNT = "InvalidParameterValue.ColumnCount"
//  INVALIDPARAMETERVALUE_HEIGHT = "InvalidParameterValue.Height"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  INVALIDPARAMETERVALUE_RESOLUTION = "InvalidParameterValue.Resolution"
//  INVALIDPARAMETERVALUE_ROWCOUNT = "InvalidParameterValue.RowCount"
//  INVALIDPARAMETERVALUE_SAMPLEINTERVAL = "InvalidParameterValue.SampleInterval"
//  INVALIDPARAMETERVALUE_SAMPLETYPE = "InvalidParameterValue.SampleType"
//  INVALIDPARAMETERVALUE_WIDTH = "InvalidParameterValue.Width"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) ModifyImageSpriteTemplate(request *ModifyImageSpriteTemplateRequest) (response *ModifyImageSpriteTemplateResponse, err error) {
    return c.ModifyImageSpriteTemplateWithContext(context.Background(), request)
}

// ModifyImageSpriteTemplate
// 修改用户自定义雪碧图模板。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_COLUMNCOUNT = "InvalidParameterValue.ColumnCount"
//  INVALIDPARAMETERVALUE_HEIGHT = "InvalidParameterValue.Height"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  INVALIDPARAMETERVALUE_RESOLUTION = "InvalidParameterValue.Resolution"
//  INVALIDPARAMETERVALUE_ROWCOUNT = "InvalidParameterValue.RowCount"
//  INVALIDPARAMETERVALUE_SAMPLEINTERVAL = "InvalidParameterValue.SampleInterval"
//  INVALIDPARAMETERVALUE_SAMPLETYPE = "InvalidParameterValue.SampleType"
//  INVALIDPARAMETERVALUE_WIDTH = "InvalidParameterValue.Width"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) ModifyImageSpriteTemplateWithContext(ctx context.Context, request *ModifyImageSpriteTemplateRequest) (response *ModifyImageSpriteTemplateResponse, err error) {
    if request == nil {
        request = NewModifyImageSpriteTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "ModifyImageSpriteTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyImageSpriteTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyImageSpriteTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewModifyLiveRecordTemplateRequest() (request *ModifyLiveRecordTemplateRequest) {
    request = &ModifyLiveRecordTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "ModifyLiveRecordTemplate")
    
    
    return
}

func NewModifyLiveRecordTemplateResponse() (response *ModifyLiveRecordTemplateResponse) {
    response = &ModifyLiveRecordTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyLiveRecordTemplate
// 修改直播录制模板
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ACCESSDBERROR = "InternalError.AccessDBError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) ModifyLiveRecordTemplate(request *ModifyLiveRecordTemplateRequest) (response *ModifyLiveRecordTemplateResponse, err error) {
    return c.ModifyLiveRecordTemplateWithContext(context.Background(), request)
}

// ModifyLiveRecordTemplate
// 修改直播录制模板
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ACCESSDBERROR = "InternalError.AccessDBError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) ModifyLiveRecordTemplateWithContext(ctx context.Context, request *ModifyLiveRecordTemplateRequest) (response *ModifyLiveRecordTemplateResponse, err error) {
    if request == nil {
        request = NewModifyLiveRecordTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "ModifyLiveRecordTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyLiveRecordTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyLiveRecordTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewModifyPersonSampleRequest() (request *ModifyPersonSampleRequest) {
    request = &ModifyPersonSampleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "ModifyPersonSample")
    
    
    return
}

func NewModifyPersonSampleResponse() (response *ModifyPersonSampleResponse) {
    response = &ModifyPersonSampleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyPersonSample
// 该接口用于根据素材 ID，修改素材样本信息，包括名称、描述的修改，以及五官、标签的添加、删除、重置操作。五官删除操作需保证至少剩余 1 张图片，否则，请使用重置操作。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_FACEDUPLICATE = "InvalidParameterValue.FaceDuplicate"
//  INVALIDPARAMETERVALUE_PICFORMATERROR = "InvalidParameterValue.PicFormatError"
//  RESOURCENOTFOUND_PERSON = "ResourceNotFound.Person"
func (c *Client) ModifyPersonSample(request *ModifyPersonSampleRequest) (response *ModifyPersonSampleResponse, err error) {
    return c.ModifyPersonSampleWithContext(context.Background(), request)
}

// ModifyPersonSample
// 该接口用于根据素材 ID，修改素材样本信息，包括名称、描述的修改，以及五官、标签的添加、删除、重置操作。五官删除操作需保证至少剩余 1 张图片，否则，请使用重置操作。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_FACEDUPLICATE = "InvalidParameterValue.FaceDuplicate"
//  INVALIDPARAMETERVALUE_PICFORMATERROR = "InvalidParameterValue.PicFormatError"
//  RESOURCENOTFOUND_PERSON = "ResourceNotFound.Person"
func (c *Client) ModifyPersonSampleWithContext(ctx context.Context, request *ModifyPersonSampleRequest) (response *ModifyPersonSampleResponse, err error) {
    if request == nil {
        request = NewModifyPersonSampleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "ModifyPersonSample")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyPersonSample require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyPersonSampleResponse()
    err = c.Send(request, response)
    return
}

func NewModifyProcessImageTemplateRequest() (request *ModifyProcessImageTemplateRequest) {
    request = &ModifyProcessImageTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "ModifyProcessImageTemplate")
    
    
    return
}

func NewModifyProcessImageTemplateResponse() (response *ModifyProcessImageTemplateResponse) {
    response = &ModifyProcessImageTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyProcessImageTemplate
// 修改图片处理模板。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) ModifyProcessImageTemplate(request *ModifyProcessImageTemplateRequest) (response *ModifyProcessImageTemplateResponse, err error) {
    return c.ModifyProcessImageTemplateWithContext(context.Background(), request)
}

// ModifyProcessImageTemplate
// 修改图片处理模板。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) ModifyProcessImageTemplateWithContext(ctx context.Context, request *ModifyProcessImageTemplateRequest) (response *ModifyProcessImageTemplateResponse, err error) {
    if request == nil {
        request = NewModifyProcessImageTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "ModifyProcessImageTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyProcessImageTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyProcessImageTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewModifyQualityControlTemplateRequest() (request *ModifyQualityControlTemplateRequest) {
    request = &ModifyQualityControlTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "ModifyQualityControlTemplate")
    
    
    return
}

func NewModifyQualityControlTemplateResponse() (response *ModifyQualityControlTemplateResponse) {
    response = &ModifyQualityControlTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyQualityControlTemplate
// 修改媒体质检模板。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_EMPTYDETECTITEM = "InvalidParameterValue.EmptyDetectItem"
//  INVALIDPARAMETERVALUE_UNKNOWNCATEGORY = "InvalidParameterValue.UnknownCategory"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) ModifyQualityControlTemplate(request *ModifyQualityControlTemplateRequest) (response *ModifyQualityControlTemplateResponse, err error) {
    return c.ModifyQualityControlTemplateWithContext(context.Background(), request)
}

// ModifyQualityControlTemplate
// 修改媒体质检模板。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_EMPTYDETECTITEM = "InvalidParameterValue.EmptyDetectItem"
//  INVALIDPARAMETERVALUE_UNKNOWNCATEGORY = "InvalidParameterValue.UnknownCategory"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) ModifyQualityControlTemplateWithContext(ctx context.Context, request *ModifyQualityControlTemplateRequest) (response *ModifyQualityControlTemplateResponse, err error) {
    if request == nil {
        request = NewModifyQualityControlTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "ModifyQualityControlTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyQualityControlTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyQualityControlTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewModifySampleSnapshotTemplateRequest() (request *ModifySampleSnapshotTemplateRequest) {
    request = &ModifySampleSnapshotTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "ModifySampleSnapshotTemplate")
    
    
    return
}

func NewModifySampleSnapshotTemplateResponse() (response *ModifySampleSnapshotTemplateResponse) {
    response = &ModifySampleSnapshotTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifySampleSnapshotTemplate
// 修改用户自定义采样截图模板。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_COMMENT = "InvalidParameterValue.Comment"
//  INVALIDPARAMETERVALUE_FILLTYPE = "InvalidParameterValue.FillType"
//  INVALIDPARAMETERVALUE_FORMAT = "InvalidParameterValue.Format"
//  INVALIDPARAMETERVALUE_HEIGHT = "InvalidParameterValue.Height"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  INVALIDPARAMETERVALUE_RESOLUTION = "InvalidParameterValue.Resolution"
//  INVALIDPARAMETERVALUE_SAMPLEINTERVAL = "InvalidParameterValue.SampleInterval"
//  INVALIDPARAMETERVALUE_SAMPLETYPE = "InvalidParameterValue.SampleType"
//  INVALIDPARAMETERVALUE_WIDTH = "InvalidParameterValue.Width"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) ModifySampleSnapshotTemplate(request *ModifySampleSnapshotTemplateRequest) (response *ModifySampleSnapshotTemplateResponse, err error) {
    return c.ModifySampleSnapshotTemplateWithContext(context.Background(), request)
}

// ModifySampleSnapshotTemplate
// 修改用户自定义采样截图模板。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_COMMENT = "InvalidParameterValue.Comment"
//  INVALIDPARAMETERVALUE_FILLTYPE = "InvalidParameterValue.FillType"
//  INVALIDPARAMETERVALUE_FORMAT = "InvalidParameterValue.Format"
//  INVALIDPARAMETERVALUE_HEIGHT = "InvalidParameterValue.Height"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  INVALIDPARAMETERVALUE_RESOLUTION = "InvalidParameterValue.Resolution"
//  INVALIDPARAMETERVALUE_SAMPLEINTERVAL = "InvalidParameterValue.SampleInterval"
//  INVALIDPARAMETERVALUE_SAMPLETYPE = "InvalidParameterValue.SampleType"
//  INVALIDPARAMETERVALUE_WIDTH = "InvalidParameterValue.Width"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) ModifySampleSnapshotTemplateWithContext(ctx context.Context, request *ModifySampleSnapshotTemplateRequest) (response *ModifySampleSnapshotTemplateResponse, err error) {
    if request == nil {
        request = NewModifySampleSnapshotTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "ModifySampleSnapshotTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifySampleSnapshotTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifySampleSnapshotTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewModifyScheduleRequest() (request *ModifyScheduleRequest) {
    request = &ModifyScheduleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "ModifySchedule")
    
    
    return
}

func NewModifyScheduleResponse() (response *ModifyScheduleResponse) {
    response = &ModifyScheduleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifySchedule
// 修改编排
//
// 可能返回的错误码:
//  FAILEDOPERATION_GENERATERESOURCE = "FailedOperation.GenerateResource"
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  FAILEDOPERATION_INVALIDUSER = "FailedOperation.InvalidUser"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ACCESSDBERROR = "InternalError.AccessDBError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) ModifySchedule(request *ModifyScheduleRequest) (response *ModifyScheduleResponse, err error) {
    return c.ModifyScheduleWithContext(context.Background(), request)
}

// ModifySchedule
// 修改编排
//
// 可能返回的错误码:
//  FAILEDOPERATION_GENERATERESOURCE = "FailedOperation.GenerateResource"
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  FAILEDOPERATION_INVALIDUSER = "FailedOperation.InvalidUser"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ACCESSDBERROR = "InternalError.AccessDBError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) ModifyScheduleWithContext(ctx context.Context, request *ModifyScheduleRequest) (response *ModifyScheduleResponse, err error) {
    if request == nil {
        request = NewModifyScheduleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "ModifySchedule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifySchedule require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyScheduleResponse()
    err = c.Send(request, response)
    return
}

func NewModifySmartEraseTemplateRequest() (request *ModifySmartEraseTemplateRequest) {
    request = &ModifySmartEraseTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "ModifySmartEraseTemplate")
    
    
    return
}

func NewModifySmartEraseTemplateResponse() (response *ModifySmartEraseTemplateResponse) {
    response = &ModifySmartEraseTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifySmartEraseTemplate
// 修改用户自定义智能擦除模板。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_AUTOAREAS = "InvalidParameterValue.AutoAreas"
//  INVALIDPARAMETERVALUE_COMMENT = "InvalidParameterValue.Comment"
//  INVALIDPARAMETERVALUE_CUSTOMAREAS = "InvalidParameterValue.CustomAreas"
//  INVALIDPARAMETERVALUE_DEFINITION = "InvalidParameterValue.Definition"
//  INVALIDPARAMETERVALUE_ERASEPRIVACYCONFIG = "InvalidParameterValue.ErasePrivacyConfig"
//  INVALIDPARAMETERVALUE_ERASESUBTITLECONFIG = "InvalidParameterValue.EraseSubtitleConfig"
//  INVALIDPARAMETERVALUE_ERASETYPE = "InvalidParameterValue.EraseType"
//  INVALIDPARAMETERVALUE_ERASEWATERMARKCONFIG = "InvalidParameterValue.EraseWatermarkConfig"
//  INVALIDPARAMETERVALUE_MODIFYDEFAULTTEMPLATE = "InvalidParameterValue.ModifyDefaultTemplate"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  INVALIDPARAMETERVALUE_OCRSWITCH = "InvalidParameterValue.OcrSwitch"
//  INVALIDPARAMETERVALUE_PRIVACYMODEL = "InvalidParameterValue.PrivacyModel"
//  INVALIDPARAMETERVALUE_PRIVACYTARGETS = "InvalidParameterValue.PrivacyTargets"
//  INVALIDPARAMETERVALUE_SUBTITLEERASEMETHOD = "InvalidParameterValue.SubtitleEraseMethod"
//  INVALIDPARAMETERVALUE_SUBTITLEFORMAT = "InvalidParameterValue.SubtitleFormat"
//  INVALIDPARAMETERVALUE_SUBTITLELANG = "InvalidParameterValue.SubtitleLang"
//  INVALIDPARAMETERVALUE_SUBTITLEMODEL = "InvalidParameterValue.SubtitleModel"
//  INVALIDPARAMETERVALUE_TRANSDSTLANG = "InvalidParameterValue.TransDstLang"
//  INVALIDPARAMETERVALUE_TRANSSWITCH = "InvalidParameterValue.TransSwitch"
//  INVALIDPARAMETERVALUE_WATERMARKERASEMETHOD = "InvalidParameterValue.WatermarkEraseMethod"
//  INVALIDPARAMETERVALUE_WATERMARKMODEL = "InvalidParameterValue.WatermarkModel"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) ModifySmartEraseTemplate(request *ModifySmartEraseTemplateRequest) (response *ModifySmartEraseTemplateResponse, err error) {
    return c.ModifySmartEraseTemplateWithContext(context.Background(), request)
}

// ModifySmartEraseTemplate
// 修改用户自定义智能擦除模板。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_AUTOAREAS = "InvalidParameterValue.AutoAreas"
//  INVALIDPARAMETERVALUE_COMMENT = "InvalidParameterValue.Comment"
//  INVALIDPARAMETERVALUE_CUSTOMAREAS = "InvalidParameterValue.CustomAreas"
//  INVALIDPARAMETERVALUE_DEFINITION = "InvalidParameterValue.Definition"
//  INVALIDPARAMETERVALUE_ERASEPRIVACYCONFIG = "InvalidParameterValue.ErasePrivacyConfig"
//  INVALIDPARAMETERVALUE_ERASESUBTITLECONFIG = "InvalidParameterValue.EraseSubtitleConfig"
//  INVALIDPARAMETERVALUE_ERASETYPE = "InvalidParameterValue.EraseType"
//  INVALIDPARAMETERVALUE_ERASEWATERMARKCONFIG = "InvalidParameterValue.EraseWatermarkConfig"
//  INVALIDPARAMETERVALUE_MODIFYDEFAULTTEMPLATE = "InvalidParameterValue.ModifyDefaultTemplate"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  INVALIDPARAMETERVALUE_OCRSWITCH = "InvalidParameterValue.OcrSwitch"
//  INVALIDPARAMETERVALUE_PRIVACYMODEL = "InvalidParameterValue.PrivacyModel"
//  INVALIDPARAMETERVALUE_PRIVACYTARGETS = "InvalidParameterValue.PrivacyTargets"
//  INVALIDPARAMETERVALUE_SUBTITLEERASEMETHOD = "InvalidParameterValue.SubtitleEraseMethod"
//  INVALIDPARAMETERVALUE_SUBTITLEFORMAT = "InvalidParameterValue.SubtitleFormat"
//  INVALIDPARAMETERVALUE_SUBTITLELANG = "InvalidParameterValue.SubtitleLang"
//  INVALIDPARAMETERVALUE_SUBTITLEMODEL = "InvalidParameterValue.SubtitleModel"
//  INVALIDPARAMETERVALUE_TRANSDSTLANG = "InvalidParameterValue.TransDstLang"
//  INVALIDPARAMETERVALUE_TRANSSWITCH = "InvalidParameterValue.TransSwitch"
//  INVALIDPARAMETERVALUE_WATERMARKERASEMETHOD = "InvalidParameterValue.WatermarkEraseMethod"
//  INVALIDPARAMETERVALUE_WATERMARKMODEL = "InvalidParameterValue.WatermarkModel"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) ModifySmartEraseTemplateWithContext(ctx context.Context, request *ModifySmartEraseTemplateRequest) (response *ModifySmartEraseTemplateResponse, err error) {
    if request == nil {
        request = NewModifySmartEraseTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "ModifySmartEraseTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifySmartEraseTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifySmartEraseTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewModifySmartSubtitleTemplateRequest() (request *ModifySmartSubtitleTemplateRequest) {
    request = &ModifySmartSubtitleTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "ModifySmartSubtitleTemplate")
    
    
    return
}

func NewModifySmartSubtitleTemplateResponse() (response *ModifySmartSubtitleTemplateResponse) {
    response = &ModifySmartSubtitleTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifySmartSubtitleTemplate
// 修改用户自定义智能字幕模板。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_ASRHOTWORDSLIBRARYID = "InvalidParameterValue.AsrHotWordsLibraryId"
//  INVALIDPARAMETERVALUE_ASRHOTWORDSSWITCH = "InvalidParameterValue.AsrHotWordsSwitch"
//  INVALIDPARAMETERVALUE_COMMENT = "InvalidParameterValue.Comment"
//  INVALIDPARAMETERVALUE_MODIFYDEFAULTTEMPLATE = "InvalidParameterValue.ModifyDefaultTemplate"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  INVALIDPARAMETERVALUE_SUBTITLEFORMAT = "InvalidParameterValue.SubtitleFormat"
//  INVALIDPARAMETERVALUE_SUBTITLETYPE = "InvalidParameterValue.SubtitleType"
//  INVALIDPARAMETERVALUE_TRANSLATEDSTLANGUAGE = "InvalidParameterValue.TranslateDstLanguage"
//  INVALIDPARAMETERVALUE_TRANSLATESWITCH = "InvalidParameterValue.TranslateSwitch"
//  INVALIDPARAMETERVALUE_VIDEOSRCLANGUAGE = "InvalidParameterValue.VideoSrcLanguage"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) ModifySmartSubtitleTemplate(request *ModifySmartSubtitleTemplateRequest) (response *ModifySmartSubtitleTemplateResponse, err error) {
    return c.ModifySmartSubtitleTemplateWithContext(context.Background(), request)
}

// ModifySmartSubtitleTemplate
// 修改用户自定义智能字幕模板。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_ASRHOTWORDSLIBRARYID = "InvalidParameterValue.AsrHotWordsLibraryId"
//  INVALIDPARAMETERVALUE_ASRHOTWORDSSWITCH = "InvalidParameterValue.AsrHotWordsSwitch"
//  INVALIDPARAMETERVALUE_COMMENT = "InvalidParameterValue.Comment"
//  INVALIDPARAMETERVALUE_MODIFYDEFAULTTEMPLATE = "InvalidParameterValue.ModifyDefaultTemplate"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  INVALIDPARAMETERVALUE_SUBTITLEFORMAT = "InvalidParameterValue.SubtitleFormat"
//  INVALIDPARAMETERVALUE_SUBTITLETYPE = "InvalidParameterValue.SubtitleType"
//  INVALIDPARAMETERVALUE_TRANSLATEDSTLANGUAGE = "InvalidParameterValue.TranslateDstLanguage"
//  INVALIDPARAMETERVALUE_TRANSLATESWITCH = "InvalidParameterValue.TranslateSwitch"
//  INVALIDPARAMETERVALUE_VIDEOSRCLANGUAGE = "InvalidParameterValue.VideoSrcLanguage"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) ModifySmartSubtitleTemplateWithContext(ctx context.Context, request *ModifySmartSubtitleTemplateRequest) (response *ModifySmartSubtitleTemplateResponse, err error) {
    if request == nil {
        request = NewModifySmartSubtitleTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "ModifySmartSubtitleTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifySmartSubtitleTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifySmartSubtitleTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewModifySnapshotByTimeOffsetTemplateRequest() (request *ModifySnapshotByTimeOffsetTemplateRequest) {
    request = &ModifySnapshotByTimeOffsetTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "ModifySnapshotByTimeOffsetTemplate")
    
    
    return
}

func NewModifySnapshotByTimeOffsetTemplateResponse() (response *ModifySnapshotByTimeOffsetTemplateResponse) {
    response = &ModifySnapshotByTimeOffsetTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifySnapshotByTimeOffsetTemplate
// 修改用户自定义指定时间点截图模板。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_COMMENT = "InvalidParameterValue.Comment"
//  INVALIDPARAMETERVALUE_FILLTYPE = "InvalidParameterValue.FillType"
//  INVALIDPARAMETERVALUE_FORMAT = "InvalidParameterValue.Format"
//  INVALIDPARAMETERVALUE_HEIGHT = "InvalidParameterValue.Height"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  INVALIDPARAMETERVALUE_WIDTH = "InvalidParameterValue.Width"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) ModifySnapshotByTimeOffsetTemplate(request *ModifySnapshotByTimeOffsetTemplateRequest) (response *ModifySnapshotByTimeOffsetTemplateResponse, err error) {
    return c.ModifySnapshotByTimeOffsetTemplateWithContext(context.Background(), request)
}

// ModifySnapshotByTimeOffsetTemplate
// 修改用户自定义指定时间点截图模板。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_COMMENT = "InvalidParameterValue.Comment"
//  INVALIDPARAMETERVALUE_FILLTYPE = "InvalidParameterValue.FillType"
//  INVALIDPARAMETERVALUE_FORMAT = "InvalidParameterValue.Format"
//  INVALIDPARAMETERVALUE_HEIGHT = "InvalidParameterValue.Height"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  INVALIDPARAMETERVALUE_WIDTH = "InvalidParameterValue.Width"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) ModifySnapshotByTimeOffsetTemplateWithContext(ctx context.Context, request *ModifySnapshotByTimeOffsetTemplateRequest) (response *ModifySnapshotByTimeOffsetTemplateResponse, err error) {
    if request == nil {
        request = NewModifySnapshotByTimeOffsetTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "ModifySnapshotByTimeOffsetTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifySnapshotByTimeOffsetTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifySnapshotByTimeOffsetTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewModifyStreamLinkEventRequest() (request *ModifyStreamLinkEventRequest) {
    request = &ModifyStreamLinkEventRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "ModifyStreamLinkEvent")
    
    
    return
}

func NewModifyStreamLinkEventResponse() (response *ModifyStreamLinkEventResponse) {
    response = &ModifyStreamLinkEventResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyStreamLinkEvent
// 修改媒体传输的事件配置信息。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_NAME = "InvalidParameter.Name"
//  INVALIDPARAMETER_STATE = "InvalidParameter.State"
func (c *Client) ModifyStreamLinkEvent(request *ModifyStreamLinkEventRequest) (response *ModifyStreamLinkEventResponse, err error) {
    return c.ModifyStreamLinkEventWithContext(context.Background(), request)
}

// ModifyStreamLinkEvent
// 修改媒体传输的事件配置信息。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_NAME = "InvalidParameter.Name"
//  INVALIDPARAMETER_STATE = "InvalidParameter.State"
func (c *Client) ModifyStreamLinkEventWithContext(ctx context.Context, request *ModifyStreamLinkEventRequest) (response *ModifyStreamLinkEventResponse, err error) {
    if request == nil {
        request = NewModifyStreamLinkEventRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "ModifyStreamLinkEvent")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyStreamLinkEvent require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyStreamLinkEventResponse()
    err = c.Send(request, response)
    return
}

func NewModifyStreamLinkFlowRequest() (request *ModifyStreamLinkFlowRequest) {
    request = &ModifyStreamLinkFlowRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "ModifyStreamLinkFlow")
    
    
    return
}

func NewModifyStreamLinkFlowResponse() (response *ModifyStreamLinkFlowResponse) {
    response = &ModifyStreamLinkFlowResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyStreamLinkFlow
// 修改媒体传输的传输流配置信息。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_NAME = "InvalidParameter.Name"
//  INVALIDPARAMETER_STATE = "InvalidParameter.State"
func (c *Client) ModifyStreamLinkFlow(request *ModifyStreamLinkFlowRequest) (response *ModifyStreamLinkFlowResponse, err error) {
    return c.ModifyStreamLinkFlowWithContext(context.Background(), request)
}

// ModifyStreamLinkFlow
// 修改媒体传输的传输流配置信息。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_NAME = "InvalidParameter.Name"
//  INVALIDPARAMETER_STATE = "InvalidParameter.State"
func (c *Client) ModifyStreamLinkFlowWithContext(ctx context.Context, request *ModifyStreamLinkFlowRequest) (response *ModifyStreamLinkFlowResponse, err error) {
    if request == nil {
        request = NewModifyStreamLinkFlowRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "ModifyStreamLinkFlow")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyStreamLinkFlow require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyStreamLinkFlowResponse()
    err = c.Send(request, response)
    return
}

func NewModifyStreamLinkInputRequest() (request *ModifyStreamLinkInputRequest) {
    request = &ModifyStreamLinkInputRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "ModifyStreamLinkInput")
    
    
    return
}

func NewModifyStreamLinkInputResponse() (response *ModifyStreamLinkInputResponse) {
    response = &ModifyStreamLinkInputResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyStreamLinkInput
// 修改媒体传输流的输入信息。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_INPUT = "InvalidParameter.Input"
//  INVALIDPARAMETER_OUTPUT = "InvalidParameter.Output"
//  INVALIDPARAMETER_PROTOCOL = "InvalidParameter.Protocol"
//  INVALIDPARAMETER_STATE = "InvalidParameter.State"
func (c *Client) ModifyStreamLinkInput(request *ModifyStreamLinkInputRequest) (response *ModifyStreamLinkInputResponse, err error) {
    return c.ModifyStreamLinkInputWithContext(context.Background(), request)
}

// ModifyStreamLinkInput
// 修改媒体传输流的输入信息。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_INPUT = "InvalidParameter.Input"
//  INVALIDPARAMETER_OUTPUT = "InvalidParameter.Output"
//  INVALIDPARAMETER_PROTOCOL = "InvalidParameter.Protocol"
//  INVALIDPARAMETER_STATE = "InvalidParameter.State"
func (c *Client) ModifyStreamLinkInputWithContext(ctx context.Context, request *ModifyStreamLinkInputRequest) (response *ModifyStreamLinkInputResponse, err error) {
    if request == nil {
        request = NewModifyStreamLinkInputRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "ModifyStreamLinkInput")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyStreamLinkInput require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyStreamLinkInputResponse()
    err = c.Send(request, response)
    return
}

func NewModifyStreamLinkOutputInfoRequest() (request *ModifyStreamLinkOutputInfoRequest) {
    request = &ModifyStreamLinkOutputInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "ModifyStreamLinkOutputInfo")
    
    
    return
}

func NewModifyStreamLinkOutputInfoResponse() (response *ModifyStreamLinkOutputInfoResponse) {
    response = &ModifyStreamLinkOutputInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyStreamLinkOutputInfo
// 修改媒体传输流的输出配置。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_INPUT = "InvalidParameter.Input"
//  INVALIDPARAMETER_NAME = "InvalidParameter.Name"
//  INVALIDPARAMETER_OUTPUT = "InvalidParameter.Output"
//  INVALIDPARAMETER_OUTPUTID = "InvalidParameter.OutputId"
//  INVALIDPARAMETER_PROTOCOL = "InvalidParameter.Protocol"
//  INVALIDPARAMETER_STATE = "InvalidParameter.State"
func (c *Client) ModifyStreamLinkOutputInfo(request *ModifyStreamLinkOutputInfoRequest) (response *ModifyStreamLinkOutputInfoResponse, err error) {
    return c.ModifyStreamLinkOutputInfoWithContext(context.Background(), request)
}

// ModifyStreamLinkOutputInfo
// 修改媒体传输流的输出配置。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_INPUT = "InvalidParameter.Input"
//  INVALIDPARAMETER_NAME = "InvalidParameter.Name"
//  INVALIDPARAMETER_OUTPUT = "InvalidParameter.Output"
//  INVALIDPARAMETER_OUTPUTID = "InvalidParameter.OutputId"
//  INVALIDPARAMETER_PROTOCOL = "InvalidParameter.Protocol"
//  INVALIDPARAMETER_STATE = "InvalidParameter.State"
func (c *Client) ModifyStreamLinkOutputInfoWithContext(ctx context.Context, request *ModifyStreamLinkOutputInfoRequest) (response *ModifyStreamLinkOutputInfoResponse, err error) {
    if request == nil {
        request = NewModifyStreamLinkOutputInfoRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "ModifyStreamLinkOutputInfo")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyStreamLinkOutputInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyStreamLinkOutputInfoResponse()
    err = c.Send(request, response)
    return
}

func NewModifyStreamLinkSecurityGroupRequest() (request *ModifyStreamLinkSecurityGroupRequest) {
    request = &ModifyStreamLinkSecurityGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "ModifyStreamLinkSecurityGroup")
    
    
    return
}

func NewModifyStreamLinkSecurityGroupResponse() (response *ModifyStreamLinkSecurityGroupResponse) {
    response = &ModifyStreamLinkSecurityGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyStreamLinkSecurityGroup
// 更新安全组。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_NAME = "InvalidParameter.Name"
//  INVALIDPARAMETER_NOTFOUND = "InvalidParameter.NotFound"
//  INVALIDPARAMETER_WHITELIST = "InvalidParameter.Whitelist"
func (c *Client) ModifyStreamLinkSecurityGroup(request *ModifyStreamLinkSecurityGroupRequest) (response *ModifyStreamLinkSecurityGroupResponse, err error) {
    return c.ModifyStreamLinkSecurityGroupWithContext(context.Background(), request)
}

// ModifyStreamLinkSecurityGroup
// 更新安全组。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_NAME = "InvalidParameter.Name"
//  INVALIDPARAMETER_NOTFOUND = "InvalidParameter.NotFound"
//  INVALIDPARAMETER_WHITELIST = "InvalidParameter.Whitelist"
func (c *Client) ModifyStreamLinkSecurityGroupWithContext(ctx context.Context, request *ModifyStreamLinkSecurityGroupRequest) (response *ModifyStreamLinkSecurityGroupResponse, err error) {
    if request == nil {
        request = NewModifyStreamLinkSecurityGroupRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "ModifyStreamLinkSecurityGroup")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyStreamLinkSecurityGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyStreamLinkSecurityGroupResponse()
    err = c.Send(request, response)
    return
}

func NewModifyTranscodeTemplateRequest() (request *ModifyTranscodeTemplateRequest) {
    request = &ModifyTranscodeTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "ModifyTranscodeTemplate")
    
    
    return
}

func NewModifyTranscodeTemplateResponse() (response *ModifyTranscodeTemplateResponse) {
    response = &ModifyTranscodeTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyTranscodeTemplate
// 修改用户自定义转码模板信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ACCESSDBERROR = "InternalError.AccessDBError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_AUDIOBITRATE = "InvalidParameterValue.AudioBitrate"
//  INVALIDPARAMETERVALUE_AUDIOCHANNEL = "InvalidParameterValue.AudioChannel"
//  INVALIDPARAMETERVALUE_AUDIOCODEC = "InvalidParameterValue.AudioCodec"
//  INVALIDPARAMETERVALUE_AUDIOSAMPLERATE = "InvalidParameterValue.AudioSampleRate"
//  INVALIDPARAMETERVALUE_CONTAINER = "InvalidParameterValue.Container"
//  INVALIDPARAMETERVALUE_FPS = "InvalidParameterValue.Fps"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  INVALIDPARAMETERVALUE_REMOVEAUDIO = "InvalidParameterValue.RemoveAudio"
//  INVALIDPARAMETERVALUE_REMOVEVIDEO = "InvalidParameterValue.RemoveVideo"
//  INVALIDPARAMETERVALUE_RESOLUTION = "InvalidParameterValue.Resolution"
//  INVALIDPARAMETERVALUE_RESOLUTIONADAPTIVE = "InvalidParameterValue.ResolutionAdaptive"
//  INVALIDPARAMETERVALUE_SAMPLERATE = "InvalidParameterValue.SampleRate"
//  INVALIDPARAMETERVALUE_TEHDTYPE = "InvalidParameterValue.TEHDType"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  INVALIDPARAMETERVALUE_VIDEOBITRATE = "InvalidParameterValue.VideoBitrate"
//  INVALIDPARAMETERVALUE_VIDEOCODEC = "InvalidParameterValue.VideoCodec"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) ModifyTranscodeTemplate(request *ModifyTranscodeTemplateRequest) (response *ModifyTranscodeTemplateResponse, err error) {
    return c.ModifyTranscodeTemplateWithContext(context.Background(), request)
}

// ModifyTranscodeTemplate
// 修改用户自定义转码模板信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ACCESSDBERROR = "InternalError.AccessDBError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_AUDIOBITRATE = "InvalidParameterValue.AudioBitrate"
//  INVALIDPARAMETERVALUE_AUDIOCHANNEL = "InvalidParameterValue.AudioChannel"
//  INVALIDPARAMETERVALUE_AUDIOCODEC = "InvalidParameterValue.AudioCodec"
//  INVALIDPARAMETERVALUE_AUDIOSAMPLERATE = "InvalidParameterValue.AudioSampleRate"
//  INVALIDPARAMETERVALUE_CONTAINER = "InvalidParameterValue.Container"
//  INVALIDPARAMETERVALUE_FPS = "InvalidParameterValue.Fps"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  INVALIDPARAMETERVALUE_REMOVEAUDIO = "InvalidParameterValue.RemoveAudio"
//  INVALIDPARAMETERVALUE_REMOVEVIDEO = "InvalidParameterValue.RemoveVideo"
//  INVALIDPARAMETERVALUE_RESOLUTION = "InvalidParameterValue.Resolution"
//  INVALIDPARAMETERVALUE_RESOLUTIONADAPTIVE = "InvalidParameterValue.ResolutionAdaptive"
//  INVALIDPARAMETERVALUE_SAMPLERATE = "InvalidParameterValue.SampleRate"
//  INVALIDPARAMETERVALUE_TEHDTYPE = "InvalidParameterValue.TEHDType"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  INVALIDPARAMETERVALUE_VIDEOBITRATE = "InvalidParameterValue.VideoBitrate"
//  INVALIDPARAMETERVALUE_VIDEOCODEC = "InvalidParameterValue.VideoCodec"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) ModifyTranscodeTemplateWithContext(ctx context.Context, request *ModifyTranscodeTemplateRequest) (response *ModifyTranscodeTemplateResponse, err error) {
    if request == nil {
        request = NewModifyTranscodeTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "ModifyTranscodeTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyTranscodeTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyTranscodeTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewModifyWatermarkTemplateRequest() (request *ModifyWatermarkTemplateRequest) {
    request = &ModifyWatermarkTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "ModifyWatermarkTemplate")
    
    
    return
}

func NewModifyWatermarkTemplateResponse() (response *ModifyWatermarkTemplateResponse) {
    response = &ModifyWatermarkTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyWatermarkTemplate
// 修改用户自定义水印模板，水印类型不允许修改。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UPLOADWATERMARKERROR = "InternalError.UploadWatermarkError"
//  INVALIDPARAMETERVALUE_COMMENT = "InvalidParameterValue.Comment"
//  INVALIDPARAMETERVALUE_COORDINATEORIGIN = "InvalidParameterValue.CoordinateOrigin"
//  INVALIDPARAMETERVALUE_HEIGHT = "InvalidParameterValue.Height"
//  INVALIDPARAMETERVALUE_IMAGECONTENT = "InvalidParameterValue.ImageContent"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  INVALIDPARAMETERVALUE_REPEATTYPE = "InvalidParameterValue.RepeatType"
//  INVALIDPARAMETERVALUE_SVGTEMPLATEHEIGHT = "InvalidParameterValue.SvgTemplateHeight"
//  INVALIDPARAMETERVALUE_SVGTEMPLATEWIDTH = "InvalidParameterValue.SvgTemplateWidth"
//  INVALIDPARAMETERVALUE_TEXTALPHA = "InvalidParameterValue.TextAlpha"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  INVALIDPARAMETERVALUE_WIDTH = "InvalidParameterValue.Width"
//  INVALIDPARAMETERVALUE_XPOS = "InvalidParameterValue.XPos"
//  INVALIDPARAMETERVALUE_YPOS = "InvalidParameterValue.YPos"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) ModifyWatermarkTemplate(request *ModifyWatermarkTemplateRequest) (response *ModifyWatermarkTemplateResponse, err error) {
    return c.ModifyWatermarkTemplateWithContext(context.Background(), request)
}

// ModifyWatermarkTemplate
// 修改用户自定义水印模板，水印类型不允许修改。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UPLOADWATERMARKERROR = "InternalError.UploadWatermarkError"
//  INVALIDPARAMETERVALUE_COMMENT = "InvalidParameterValue.Comment"
//  INVALIDPARAMETERVALUE_COORDINATEORIGIN = "InvalidParameterValue.CoordinateOrigin"
//  INVALIDPARAMETERVALUE_HEIGHT = "InvalidParameterValue.Height"
//  INVALIDPARAMETERVALUE_IMAGECONTENT = "InvalidParameterValue.ImageContent"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  INVALIDPARAMETERVALUE_REPEATTYPE = "InvalidParameterValue.RepeatType"
//  INVALIDPARAMETERVALUE_SVGTEMPLATEHEIGHT = "InvalidParameterValue.SvgTemplateHeight"
//  INVALIDPARAMETERVALUE_SVGTEMPLATEWIDTH = "InvalidParameterValue.SvgTemplateWidth"
//  INVALIDPARAMETERVALUE_TEXTALPHA = "InvalidParameterValue.TextAlpha"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  INVALIDPARAMETERVALUE_WIDTH = "InvalidParameterValue.Width"
//  INVALIDPARAMETERVALUE_XPOS = "InvalidParameterValue.XPos"
//  INVALIDPARAMETERVALUE_YPOS = "InvalidParameterValue.YPos"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) ModifyWatermarkTemplateWithContext(ctx context.Context, request *ModifyWatermarkTemplateRequest) (response *ModifyWatermarkTemplateResponse, err error) {
    if request == nil {
        request = NewModifyWatermarkTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "ModifyWatermarkTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyWatermarkTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyWatermarkTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewModifyWordSampleRequest() (request *ModifyWordSampleRequest) {
    request = &ModifyWordSampleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "ModifyWordSample")
    
    
    return
}

func NewModifyWordSampleResponse() (response *ModifyWordSampleResponse) {
    response = &ModifyWordSampleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyWordSample
// 该接口用于修改关键词的应用场景、标签，关键词本身不可修改，如需修改，可删除重建。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_WORD = "ResourceNotFound.Word"
func (c *Client) ModifyWordSample(request *ModifyWordSampleRequest) (response *ModifyWordSampleResponse, err error) {
    return c.ModifyWordSampleWithContext(context.Background(), request)
}

// ModifyWordSample
// 该接口用于修改关键词的应用场景、标签，关键词本身不可修改，如需修改，可删除重建。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_WORD = "ResourceNotFound.Word"
func (c *Client) ModifyWordSampleWithContext(ctx context.Context, request *ModifyWordSampleRequest) (response *ModifyWordSampleResponse, err error) {
    if request == nil {
        request = NewModifyWordSampleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "ModifyWordSample")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyWordSample require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyWordSampleResponse()
    err = c.Send(request, response)
    return
}

func NewParseLiveStreamProcessNotificationRequest() (request *ParseLiveStreamProcessNotificationRequest) {
    request = &ParseLiveStreamProcessNotificationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "ParseLiveStreamProcessNotification")
    
    
    return
}

func NewParseLiveStreamProcessNotificationResponse() (response *ParseLiveStreamProcessNotificationResponse) {
    response = &ParseLiveStreamProcessNotificationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ParseLiveStreamProcessNotification
// 从 CMQ 获取到消息后，从消息的 msgBody 字段中解析出 MPS 直播流处理事件通知的内容。
//
// 该接口不用于发起网络调用，而是用来帮助生成各个语言平台的 SDK，您可以参考 SDK 中的解析实现事件通知的解析。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INVALIDPARAMETERVALUE_INVALIDCONTENT = "InvalidParameterValue.InvalidContent"
func (c *Client) ParseLiveStreamProcessNotification(request *ParseLiveStreamProcessNotificationRequest) (response *ParseLiveStreamProcessNotificationResponse, err error) {
    return c.ParseLiveStreamProcessNotificationWithContext(context.Background(), request)
}

// ParseLiveStreamProcessNotification
// 从 CMQ 获取到消息后，从消息的 msgBody 字段中解析出 MPS 直播流处理事件通知的内容。
//
// 该接口不用于发起网络调用，而是用来帮助生成各个语言平台的 SDK，您可以参考 SDK 中的解析实现事件通知的解析。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INVALIDPARAMETERVALUE_INVALIDCONTENT = "InvalidParameterValue.InvalidContent"
func (c *Client) ParseLiveStreamProcessNotificationWithContext(ctx context.Context, request *ParseLiveStreamProcessNotificationRequest) (response *ParseLiveStreamProcessNotificationResponse, err error) {
    if request == nil {
        request = NewParseLiveStreamProcessNotificationRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "ParseLiveStreamProcessNotification")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ParseLiveStreamProcessNotification require credential")
    }

    request.SetContext(ctx)
    
    response = NewParseLiveStreamProcessNotificationResponse()
    err = c.Send(request, response)
    return
}

func NewParseNotificationRequest() (request *ParseNotificationRequest) {
    request = &ParseNotificationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "ParseNotification")
    
    
    return
}

func NewParseNotificationResponse() (response *ParseNotificationResponse) {
    response = &ParseNotificationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ParseNotification
// 从 CMQ 获取到消息后，从消息的 msgBody 字段中解析出 MPS 事件通知的内容。
//
// 该接口不用于发起网络调用，而是用来帮助生成各个语言平台的 SDK，您可以参考 SDK 中的解析函数，实现事件通知的解析。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDCONTENT = "InvalidParameterValue.InvalidContent"
func (c *Client) ParseNotification(request *ParseNotificationRequest) (response *ParseNotificationResponse, err error) {
    return c.ParseNotificationWithContext(context.Background(), request)
}

// ParseNotification
// 从 CMQ 获取到消息后，从消息的 msgBody 字段中解析出 MPS 事件通知的内容。
//
// 该接口不用于发起网络调用，而是用来帮助生成各个语言平台的 SDK，您可以参考 SDK 中的解析函数，实现事件通知的解析。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDCONTENT = "InvalidParameterValue.InvalidContent"
func (c *Client) ParseNotificationWithContext(ctx context.Context, request *ParseNotificationRequest) (response *ParseNotificationResponse, err error) {
    if request == nil {
        request = NewParseNotificationRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "ParseNotification")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ParseNotification require credential")
    }

    request.SetContext(ctx)
    
    response = NewParseNotificationResponse()
    err = c.Send(request, response)
    return
}

func NewProcessImageRequest() (request *ProcessImageRequest) {
    request = &ProcessImageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "ProcessImage")
    
    
    return
}

func NewProcessImageResponse() (response *ProcessImageResponse) {
    response = &ProcessImageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ProcessImage
// 发起图片处理，功能包括：
//
// 1. 格式转换；
//
// 2. 图像增强；
//
// 3. 图像擦除;
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_SESSIONCONTEXTTOOLONG = "InvalidParameterValue.SessionContextTooLong"
//  INVALIDPARAMETERVALUE_SESSIONID = "InvalidParameterValue.SessionId"
//  INVALIDPARAMETERVALUE_SESSIONIDTOOLONG = "InvalidParameterValue.SessionIdTooLong"
func (c *Client) ProcessImage(request *ProcessImageRequest) (response *ProcessImageResponse, err error) {
    return c.ProcessImageWithContext(context.Background(), request)
}

// ProcessImage
// 发起图片处理，功能包括：
//
// 1. 格式转换；
//
// 2. 图像增强；
//
// 3. 图像擦除;
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_SESSIONCONTEXTTOOLONG = "InvalidParameterValue.SessionContextTooLong"
//  INVALIDPARAMETERVALUE_SESSIONID = "InvalidParameterValue.SessionId"
//  INVALIDPARAMETERVALUE_SESSIONIDTOOLONG = "InvalidParameterValue.SessionIdTooLong"
func (c *Client) ProcessImageWithContext(ctx context.Context, request *ProcessImageRequest) (response *ProcessImageResponse, err error) {
    if request == nil {
        request = NewProcessImageRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "ProcessImage")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ProcessImage require credential")
    }

    request.SetContext(ctx)
    
    response = NewProcessImageResponse()
    err = c.Send(request, response)
    return
}

func NewProcessLiveStreamRequest() (request *ProcessLiveStreamRequest) {
    request = &ProcessLiveStreamRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "ProcessLiveStream")
    
    
    return
}

func NewProcessLiveStreamResponse() (response *ProcessLiveStreamResponse) {
    response = &ProcessLiveStreamResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ProcessLiveStream
// 对直播流媒体发起处理任务，功能包括：
//
// 
//
// * 智能内容审核（画面鉴黄、敏感信息检测、声音鉴黄）；
//
// * 智能内容识别（人脸、文本全文、文本关键词、语音全文、语音关键词、语音实时翻译、物体识别、游戏打点）。
//
// * 智能内容分析（拆条，集锦）。
//
// * 质检（直播流格式诊断、音画内容检测（抖动、模糊、低光照、过曝光、黑边、白边、黑屏、白屏、花屏、噪点、马赛克、二维码等）、无参考打分）。
//
// * 录制
//
// 
//
// 直播流处理事件通知支持HTTP回调，也支持实时写入用户指定的消息队列 CMQ 中，用户从消息队列 CMQ 中获取事件通知结果，同时处理过程中存在输出文件的，会写入用户指定的输出文件的目标存储中。
//
// 可能返回的错误码:
//  FAILEDOPERATION_GENERATERESOURCE = "FailedOperation.GenerateResource"
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_DEFINITION = "InvalidParameterValue.Definition"
//  INVALIDPARAMETERVALUE_DEFINITIONS = "InvalidParameterValue.Definitions"
//  INVALIDPARAMETERVALUE_SESSIONID = "InvalidParameterValue.SessionId"
func (c *Client) ProcessLiveStream(request *ProcessLiveStreamRequest) (response *ProcessLiveStreamResponse, err error) {
    return c.ProcessLiveStreamWithContext(context.Background(), request)
}

// ProcessLiveStream
// 对直播流媒体发起处理任务，功能包括：
//
// 
//
// * 智能内容审核（画面鉴黄、敏感信息检测、声音鉴黄）；
//
// * 智能内容识别（人脸、文本全文、文本关键词、语音全文、语音关键词、语音实时翻译、物体识别、游戏打点）。
//
// * 智能内容分析（拆条，集锦）。
//
// * 质检（直播流格式诊断、音画内容检测（抖动、模糊、低光照、过曝光、黑边、白边、黑屏、白屏、花屏、噪点、马赛克、二维码等）、无参考打分）。
//
// * 录制
//
// 
//
// 直播流处理事件通知支持HTTP回调，也支持实时写入用户指定的消息队列 CMQ 中，用户从消息队列 CMQ 中获取事件通知结果，同时处理过程中存在输出文件的，会写入用户指定的输出文件的目标存储中。
//
// 可能返回的错误码:
//  FAILEDOPERATION_GENERATERESOURCE = "FailedOperation.GenerateResource"
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_DEFINITION = "InvalidParameterValue.Definition"
//  INVALIDPARAMETERVALUE_DEFINITIONS = "InvalidParameterValue.Definitions"
//  INVALIDPARAMETERVALUE_SESSIONID = "InvalidParameterValue.SessionId"
func (c *Client) ProcessLiveStreamWithContext(ctx context.Context, request *ProcessLiveStreamRequest) (response *ProcessLiveStreamResponse, err error) {
    if request == nil {
        request = NewProcessLiveStreamRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "ProcessLiveStream")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ProcessLiveStream require credential")
    }

    request.SetContext(ctx)
    
    response = NewProcessLiveStreamResponse()
    err = c.Send(request, response)
    return
}

func NewProcessMediaRequest() (request *ProcessMediaRequest) {
    request = &ProcessMediaRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "ProcessMedia")
    
    
    return
}

func NewProcessMediaResponse() (response *ProcessMediaResponse) {
    response = &ProcessMediaResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ProcessMedia
// 对 URL视频链接 或 COS 中的媒体文件发起处理任务，功能包括：
//
// - 音视频转码（例如普通转码、极速高清转码、音视频增强、添加明水印、添加数字水印）；
//
// - 音视频转自适应码流；
//
// - 视频转动图；
//
// - 对视频按指定时间点截图；
//
// - 对视频采样截图；
//
// - 对视频截图雪碧图；
//
// - 媒体质检（例如媒体格式诊断、音画内容检测、无参考打分，其中音画内容检测主要针对抖动、模糊、低光照、过曝光、花屏、噪点、马赛克、二维码等问题）;
//
// - 智能字幕（例如生成字幕并翻译）；
//
// - 智能擦除（例如去水印、去字幕、隐私保护）；
//
// - 智能内容审核（例如鉴黄、敏感信息检测）；
//
// - 智能内容分析（例如标签、分类、封面、按帧标签、拆条、集锦、片头片尾、游戏打点）；
//
// - 智能内容识别（例如人脸、文本全文、文本关键词、语音全文、语音关键词、语音翻译、物体识别）；
//
// 可能返回的错误码:
//  FAILEDOPERATION_GENERATERESOURCE = "FailedOperation.GenerateResource"
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_SESSIONCONTEXTTOOLONG = "InvalidParameterValue.SessionContextTooLong"
//  INVALIDPARAMETERVALUE_SESSIONID = "InvalidParameterValue.SessionId"
//  INVALIDPARAMETERVALUE_SESSIONIDTOOLONG = "InvalidParameterValue.SessionIdTooLong"
func (c *Client) ProcessMedia(request *ProcessMediaRequest) (response *ProcessMediaResponse, err error) {
    return c.ProcessMediaWithContext(context.Background(), request)
}

// ProcessMedia
// 对 URL视频链接 或 COS 中的媒体文件发起处理任务，功能包括：
//
// - 音视频转码（例如普通转码、极速高清转码、音视频增强、添加明水印、添加数字水印）；
//
// - 音视频转自适应码流；
//
// - 视频转动图；
//
// - 对视频按指定时间点截图；
//
// - 对视频采样截图；
//
// - 对视频截图雪碧图；
//
// - 媒体质检（例如媒体格式诊断、音画内容检测、无参考打分，其中音画内容检测主要针对抖动、模糊、低光照、过曝光、花屏、噪点、马赛克、二维码等问题）;
//
// - 智能字幕（例如生成字幕并翻译）；
//
// - 智能擦除（例如去水印、去字幕、隐私保护）；
//
// - 智能内容审核（例如鉴黄、敏感信息检测）；
//
// - 智能内容分析（例如标签、分类、封面、按帧标签、拆条、集锦、片头片尾、游戏打点）；
//
// - 智能内容识别（例如人脸、文本全文、文本关键词、语音全文、语音关键词、语音翻译、物体识别）；
//
// 可能返回的错误码:
//  FAILEDOPERATION_GENERATERESOURCE = "FailedOperation.GenerateResource"
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_SESSIONCONTEXTTOOLONG = "InvalidParameterValue.SessionContextTooLong"
//  INVALIDPARAMETERVALUE_SESSIONID = "InvalidParameterValue.SessionId"
//  INVALIDPARAMETERVALUE_SESSIONIDTOOLONG = "InvalidParameterValue.SessionIdTooLong"
func (c *Client) ProcessMediaWithContext(ctx context.Context, request *ProcessMediaRequest) (response *ProcessMediaResponse, err error) {
    if request == nil {
        request = NewProcessMediaRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "ProcessMedia")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ProcessMedia require credential")
    }

    request.SetContext(ctx)
    
    response = NewProcessMediaResponse()
    err = c.Send(request, response)
    return
}

func NewRecognizeAudioRequest() (request *RecognizeAudioRequest) {
    request = &RecognizeAudioRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "RecognizeAudio")
    
    
    return
}

func NewRecognizeAudioResponse() (response *RecognizeAudioResponse) {
    response = &RecognizeAudioResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RecognizeAudio
// 同步接口，返回语音识别结果
//
// 可能返回的错误码:
//  INTERNALERROR_RECOGNITIONERROR = "InternalError.RecognitionError"
//  INVALIDPARAMETERVALUE_AUDIODATA = "InvalidParameterValue.AudioData"
//  INVALIDPARAMETERVALUE_AUDIODATATOOLONG = "InvalidParameterValue.AudioDataTooLong"
//  INVALIDPARAMETERVALUE_AUDIOFORMAT = "InvalidParameterValue.AudioFormat"
//  INVALIDPARAMETERVALUE_SAMPLERATE = "InvalidParameterValue.SampleRate"
//  INVALIDPARAMETERVALUE_SOURCELANGUAGE = "InvalidParameterValue.SourceLanguage"
//  RESOURCENOTFOUND_USERUNREGISTER = "ResourceNotFound.UserUnregister"
func (c *Client) RecognizeAudio(request *RecognizeAudioRequest) (response *RecognizeAudioResponse, err error) {
    return c.RecognizeAudioWithContext(context.Background(), request)
}

// RecognizeAudio
// 同步接口，返回语音识别结果
//
// 可能返回的错误码:
//  INTERNALERROR_RECOGNITIONERROR = "InternalError.RecognitionError"
//  INVALIDPARAMETERVALUE_AUDIODATA = "InvalidParameterValue.AudioData"
//  INVALIDPARAMETERVALUE_AUDIODATATOOLONG = "InvalidParameterValue.AudioDataTooLong"
//  INVALIDPARAMETERVALUE_AUDIOFORMAT = "InvalidParameterValue.AudioFormat"
//  INVALIDPARAMETERVALUE_SAMPLERATE = "InvalidParameterValue.SampleRate"
//  INVALIDPARAMETERVALUE_SOURCELANGUAGE = "InvalidParameterValue.SourceLanguage"
//  RESOURCENOTFOUND_USERUNREGISTER = "ResourceNotFound.UserUnregister"
func (c *Client) RecognizeAudioWithContext(ctx context.Context, request *RecognizeAudioRequest) (response *RecognizeAudioResponse, err error) {
    if request == nil {
        request = NewRecognizeAudioRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "RecognizeAudio")
    
    if c.GetCredential() == nil {
        return nil, errors.New("RecognizeAudio require credential")
    }

    request.SetContext(ctx)
    
    response = NewRecognizeAudioResponse()
    err = c.Send(request, response)
    return
}

func NewRecognizeMediaForZhiXueRequest() (request *RecognizeMediaForZhiXueRequest) {
    request = &RecognizeMediaForZhiXueRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "RecognizeMediaForZhiXue")
    
    
    return
}

func NewRecognizeMediaForZhiXueResponse() (response *RecognizeMediaForZhiXueResponse) {
    response = &RecognizeMediaForZhiXueResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RecognizeMediaForZhiXue
// 智能媒体识别，包含表情和动作识别。仅用于智学，其他调用无效。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) RecognizeMediaForZhiXue(request *RecognizeMediaForZhiXueRequest) (response *RecognizeMediaForZhiXueResponse, err error) {
    return c.RecognizeMediaForZhiXueWithContext(context.Background(), request)
}

// RecognizeMediaForZhiXue
// 智能媒体识别，包含表情和动作识别。仅用于智学，其他调用无效。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) RecognizeMediaForZhiXueWithContext(ctx context.Context, request *RecognizeMediaForZhiXueRequest) (response *RecognizeMediaForZhiXueResponse, err error) {
    if request == nil {
        request = NewRecognizeMediaForZhiXueRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "RecognizeMediaForZhiXue")
    
    if c.GetCredential() == nil {
        return nil, errors.New("RecognizeMediaForZhiXue require credential")
    }

    request.SetContext(ctx)
    
    response = NewRecognizeMediaForZhiXueResponse()
    err = c.Send(request, response)
    return
}

func NewResetWorkflowRequest() (request *ResetWorkflowRequest) {
    request = &ResetWorkflowRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "ResetWorkflow")
    
    
    return
}

func NewResetWorkflowResponse() (response *ResetWorkflowResponse) {
    response = &ResetWorkflowResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ResetWorkflow
// 重新设置一个已经存在且处于禁用状态的工作流。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) ResetWorkflow(request *ResetWorkflowRequest) (response *ResetWorkflowResponse, err error) {
    return c.ResetWorkflowWithContext(context.Background(), request)
}

// ResetWorkflow
// 重新设置一个已经存在且处于禁用状态的工作流。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) ResetWorkflowWithContext(ctx context.Context, request *ResetWorkflowRequest) (response *ResetWorkflowResponse, err error) {
    if request == nil {
        request = NewResetWorkflowRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "ResetWorkflow")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ResetWorkflow require credential")
    }

    request.SetContext(ctx)
    
    response = NewResetWorkflowResponse()
    err = c.Send(request, response)
    return
}

func NewStartStreamLinkFlowRequest() (request *StartStreamLinkFlowRequest) {
    request = &StartStreamLinkFlowRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "StartStreamLinkFlow")
    
    
    return
}

func NewStartStreamLinkFlowResponse() (response *StartStreamLinkFlowResponse) {
    response = &StartStreamLinkFlowResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// StartStreamLinkFlow
// 启动媒体传输流。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_OUTPUTGROUPS = "InvalidParameter.OutputGroups"
//  INVALIDPARAMETER_STATE = "InvalidParameter.State"
func (c *Client) StartStreamLinkFlow(request *StartStreamLinkFlowRequest) (response *StartStreamLinkFlowResponse, err error) {
    return c.StartStreamLinkFlowWithContext(context.Background(), request)
}

// StartStreamLinkFlow
// 启动媒体传输流。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_OUTPUTGROUPS = "InvalidParameter.OutputGroups"
//  INVALIDPARAMETER_STATE = "InvalidParameter.State"
func (c *Client) StartStreamLinkFlowWithContext(ctx context.Context, request *StartStreamLinkFlowRequest) (response *StartStreamLinkFlowResponse, err error) {
    if request == nil {
        request = NewStartStreamLinkFlowRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "StartStreamLinkFlow")
    
    if c.GetCredential() == nil {
        return nil, errors.New("StartStreamLinkFlow require credential")
    }

    request.SetContext(ctx)
    
    response = NewStartStreamLinkFlowResponse()
    err = c.Send(request, response)
    return
}

func NewStopStreamLinkFlowRequest() (request *StopStreamLinkFlowRequest) {
    request = &StopStreamLinkFlowRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "StopStreamLinkFlow")
    
    
    return
}

func NewStopStreamLinkFlowResponse() (response *StopStreamLinkFlowResponse) {
    response = &StopStreamLinkFlowResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// StopStreamLinkFlow
// 停止媒体传输流。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_STATE = "InvalidParameter.State"
func (c *Client) StopStreamLinkFlow(request *StopStreamLinkFlowRequest) (response *StopStreamLinkFlowResponse, err error) {
    return c.StopStreamLinkFlowWithContext(context.Background(), request)
}

// StopStreamLinkFlow
// 停止媒体传输流。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_STATE = "InvalidParameter.State"
func (c *Client) StopStreamLinkFlowWithContext(ctx context.Context, request *StopStreamLinkFlowRequest) (response *StopStreamLinkFlowResponse, err error) {
    if request == nil {
        request = NewStopStreamLinkFlowRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "StopStreamLinkFlow")
    
    if c.GetCredential() == nil {
        return nil, errors.New("StopStreamLinkFlow require credential")
    }

    request.SetContext(ctx)
    
    response = NewStopStreamLinkFlowResponse()
    err = c.Send(request, response)
    return
}

func NewSyncDubbingRequest() (request *SyncDubbingRequest) {
    request = &SyncDubbingRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "SyncDubbing")
    
    
    return
}

func NewSyncDubbingResponse() (response *SyncDubbingResponse) {
    response = &SyncDubbingResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SyncDubbing
// 同步接口，返回克隆音色Id或合成音频结果
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_STATE = "InvalidParameter.State"
func (c *Client) SyncDubbing(request *SyncDubbingRequest) (response *SyncDubbingResponse, err error) {
    return c.SyncDubbingWithContext(context.Background(), request)
}

// SyncDubbing
// 同步接口，返回克隆音色Id或合成音频结果
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_STATE = "InvalidParameter.State"
func (c *Client) SyncDubbingWithContext(ctx context.Context, request *SyncDubbingRequest) (response *SyncDubbingResponse, err error) {
    if request == nil {
        request = NewSyncDubbingRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "SyncDubbing")
    
    if c.GetCredential() == nil {
        return nil, errors.New("SyncDubbing require credential")
    }

    request.SetContext(ctx)
    
    response = NewSyncDubbingResponse()
    err = c.Send(request, response)
    return
}

func NewTextTranslationRequest() (request *TextTranslationRequest) {
    request = &TextTranslationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "TextTranslation")
    
    
    return
}

func NewTextTranslationResponse() (response *TextTranslationResponse) {
    response = &TextTranslationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// TextTranslation
// 文本翻译
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_SOURCELANGUAGE = "InvalidParameterValue.SourceLanguage"
//  INVALIDPARAMETERVALUE_SOURCETEXT = "InvalidParameterValue.SourceText"
//  INVALIDPARAMETERVALUE_TEXTCONTENT = "InvalidParameterValue.TextContent"
//  INVALIDPARAMETERVALUE_TRANSLATEDSTLANGUAGE = "InvalidParameterValue.TranslateDstLanguage"
//  RESOURCENOTFOUND_USERUNREGISTER = "ResourceNotFound.UserUnregister"
//  UNSUPPORTEDOPERATION_TEXTTOOLONG = "UnsupportedOperation.TextTooLong"
func (c *Client) TextTranslation(request *TextTranslationRequest) (response *TextTranslationResponse, err error) {
    return c.TextTranslationWithContext(context.Background(), request)
}

// TextTranslation
// 文本翻译
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_SOURCELANGUAGE = "InvalidParameterValue.SourceLanguage"
//  INVALIDPARAMETERVALUE_SOURCETEXT = "InvalidParameterValue.SourceText"
//  INVALIDPARAMETERVALUE_TEXTCONTENT = "InvalidParameterValue.TextContent"
//  INVALIDPARAMETERVALUE_TRANSLATEDSTLANGUAGE = "InvalidParameterValue.TranslateDstLanguage"
//  RESOURCENOTFOUND_USERUNREGISTER = "ResourceNotFound.UserUnregister"
//  UNSUPPORTEDOPERATION_TEXTTOOLONG = "UnsupportedOperation.TextTooLong"
func (c *Client) TextTranslationWithContext(ctx context.Context, request *TextTranslationRequest) (response *TextTranslationResponse, err error) {
    if request == nil {
        request = NewTextTranslationRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "TextTranslation")
    
    if c.GetCredential() == nil {
        return nil, errors.New("TextTranslation require credential")
    }

    request.SetContext(ctx)
    
    response = NewTextTranslationResponse()
    err = c.Send(request, response)
    return
}

func NewWithdrawsWatermarkRequest() (request *WithdrawsWatermarkRequest) {
    request = &WithdrawsWatermarkRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "WithdrawsWatermark")
    
    
    return
}

func NewWithdrawsWatermarkResponse() (response *WithdrawsWatermarkResponse) {
    response = &WithdrawsWatermarkResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// WithdrawsWatermark
// 提取视频中的盲水印。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) WithdrawsWatermark(request *WithdrawsWatermarkRequest) (response *WithdrawsWatermarkResponse, err error) {
    return c.WithdrawsWatermarkWithContext(context.Background(), request)
}

// WithdrawsWatermark
// 提取视频中的盲水印。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) WithdrawsWatermarkWithContext(ctx context.Context, request *WithdrawsWatermarkRequest) (response *WithdrawsWatermarkResponse, err error) {
    if request == nil {
        request = NewWithdrawsWatermarkRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "WithdrawsWatermark")
    
    if c.GetCredential() == nil {
        return nil, errors.New("WithdrawsWatermark require credential")
    }

    request.SetContext(ctx)
    
    response = NewWithdrawsWatermarkResponse()
    err = c.Send(request, response)
    return
}
