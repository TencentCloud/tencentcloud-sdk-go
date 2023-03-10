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
//  INVALIDPARAMETERVALUE_COMMENT = "InvalidParameterValue.Comment"
//  INVALIDPARAMETERVALUE_DEFAULTLIBRARYLABELSET = "InvalidParameterValue.DefaultLibraryLabelSet"
//  INVALIDPARAMETERVALUE_FACELIBRARY = "InvalidParameterValue.FaceLibrary"
//  INVALIDPARAMETERVALUE_FACESCORE = "InvalidParameterValue.FaceScore"
//  INVALIDPARAMETERVALUE_LABELSET = "InvalidParameterValue.LabelSet"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  INVALIDPARAMETERVALUE_OBJECTLIBRARY = "InvalidParameterValue.ObjectLibrary"
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
//  INVALIDPARAMETERVALUE_COMMENT = "InvalidParameterValue.Comment"
//  INVALIDPARAMETERVALUE_DEFAULTLIBRARYLABELSET = "InvalidParameterValue.DefaultLibraryLabelSet"
//  INVALIDPARAMETERVALUE_FACELIBRARY = "InvalidParameterValue.FaceLibrary"
//  INVALIDPARAMETERVALUE_FACESCORE = "InvalidParameterValue.FaceScore"
//  INVALIDPARAMETERVALUE_LABELSET = "InvalidParameterValue.LabelSet"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  INVALIDPARAMETERVALUE_OBJECTLIBRARY = "InvalidParameterValue.ObjectLibrary"
//  INVALIDPARAMETERVALUE_SUBTITLEFORMAT = "InvalidParameterValue.SubtitleFormat"
//  INVALIDPARAMETERVALUE_SWITCH = "InvalidParameterValue.Switch"
//  INVALIDPARAMETERVALUE_USERDEFINELIBRARYLABELSET = "InvalidParameterValue.UserDefineLibraryLabelSet"
//  LIMITEXCEEDED_TOOMUCHTEMPLATE = "LimitExceeded.TooMuchTemplate"
func (c *Client) CreateAIRecognitionTemplateWithContext(ctx context.Context, request *CreateAIRecognitionTemplateRequest) (response *CreateAIRecognitionTemplateResponse, err error) {
    if request == nil {
        request = NewCreateAIRecognitionTemplateRequest()
    }
    
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAdaptiveDynamicStreamingTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAdaptiveDynamicStreamingTemplateResponse()
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAnimatedGraphicsTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAnimatedGraphicsTemplateResponse()
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateImageSpriteTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateImageSpriteTemplateResponse()
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreatePersonSample require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreatePersonSampleResponse()
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
// 
//
// 注意：创建编排成功后是禁用状态，需要手动启用。
//
// 可能返回的错误码:
//  FAILEDOPERATION_COSSTATUSINAVLID = "FailedOperation.CosStatusInavlid"
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
// 
//
// 注意：创建编排成功后是禁用状态，需要手动启用。
//
// 可能返回的错误码:
//  FAILEDOPERATION_COSSTATUSINAVLID = "FailedOperation.CosStatusInavlid"
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateSchedule require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateScheduleResponse()
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateSnapshotByTimeOffsetTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateSnapshotByTimeOffsetTemplateResponse()
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
//  INVALIDPARAMETER_INPUT = "InvalidParameter.Input"
//  INVALIDPARAMETER_MAXBANDWIDTH = "InvalidParameter.MaxBandwidth"
//  INVALIDPARAMETER_NAME = "InvalidParameter.Name"
func (c *Client) CreateStreamLinkFlowWithContext(ctx context.Context, request *CreateStreamLinkFlowRequest) (response *CreateStreamLinkFlowResponse, err error) {
    if request == nil {
        request = NewCreateStreamLinkFlowRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateStreamLinkFlow require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateStreamLinkFlowResponse()
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateStreamLinkOutputInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateStreamLinkOutputInfoResponse()
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
// 创建用户自定义转码模板，数量上限：1000。
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
// 创建用户自定义转码模板，数量上限：1000。
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateTranscodeTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateTranscodeTemplateResponse()
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteAnimatedGraphicsTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteAnimatedGraphicsTemplateResponse()
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteImageSpriteTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteImageSpriteTemplateResponse()
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeletePersonSample require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeletePersonSampleResponse()
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteSchedule require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteScheduleResponse()
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteSnapshotByTimeOffsetTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteSnapshotByTimeOffsetTemplateResponse()
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteStreamLinkOutput require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteStreamLinkOutputResponse()
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAdaptiveDynamicStreamingTemplates require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAdaptiveDynamicStreamingTemplatesResponse()
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAnimatedGraphicsTemplates require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAnimatedGraphicsTemplatesResponse()
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeContentReviewTemplates require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeContentReviewTemplatesResponse()
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeImageSpriteTemplates require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeImageSpriteTemplatesResponse()
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePersonSamples require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePersonSamplesResponse()
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
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  FAILEDOPERATION_INVALIDUSER = "FailedOperation.InvalidUser"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ACCESSDBERROR = "InternalError.AccessDBError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeSchedulesWithContext(ctx context.Context, request *DescribeSchedulesRequest) (response *DescribeSchedulesResponse, err error) {
    if request == nil {
        request = NewDescribeSchedulesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSchedules require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSchedulesResponse()
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeStreamLinkActivateState require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeStreamLinkActivateStateResponse()
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeStreamLinkRegions require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeStreamLinkRegionsResponse()
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
// 通过任务 ID 查询任务的执行状态和结果的详细信息（最多可以查询3天之内提交的任务）。
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
// 通过任务 ID 查询任务的执行状态和结果的详细信息（最多可以查询3天之内提交的任务）。
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTranscodeTemplates require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTranscodeTemplatesResponse()
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DisableWorkflow require credential")
    }

    request.SetContext(ctx)
    
    response = NewDisableWorkflowResponse()
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
// 对视频进行编辑（剪辑、拼接等），生成一个新的点播视频。编辑的功能包括：
//
// 
//
// 1. 对一个文件进行剪辑，生成一个新的视频；
//
// 2. 对多个文件进行拼接，生成一个新的视频；
//
// 3. 对多个文件进行剪辑，然后再拼接，生成一个新的视频。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) EditMedia(request *EditMediaRequest) (response *EditMediaResponse, err error) {
    return c.EditMediaWithContext(context.Background(), request)
}

// EditMedia
// 对视频进行编辑（剪辑、拼接等），生成一个新的点播视频。编辑的功能包括：
//
// 
//
// 1. 对一个文件进行剪辑，生成一个新的视频；
//
// 2. 对多个文件进行拼接，生成一个新的视频；
//
// 3. 对多个文件进行剪辑，然后再拼接，生成一个新的视频。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) EditMediaWithContext(ctx context.Context, request *EditMediaRequest) (response *EditMediaResponse, err error) {
    if request == nil {
        request = NewEditMediaRequest()
    }
    
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
func (c *Client) EnableSchedule(request *EnableScheduleRequest) (response *EnableScheduleResponse, err error) {
    return c.EnableScheduleWithContext(context.Background(), request)
}

// EnableSchedule
// 启用自动化触发编排任务。
//
// 可能返回的错误码:
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
func (c *Client) EnableScheduleWithContext(ctx context.Context, request *EnableScheduleRequest) (response *EnableScheduleResponse, err error) {
    if request == nil {
        request = NewEnableScheduleRequest()
    }
    
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("ExecuteFunction require credential")
    }

    request.SetContext(ctx)
    
    response = NewExecuteFunctionResponse()
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
//  INVALIDPARAMETERVALUE_SUBTITLEFORMAT = "InvalidParameterValue.SubtitleFormat"
//  INVALIDPARAMETERVALUE_SWITCH = "InvalidParameterValue.Switch"
//  INVALIDPARAMETERVALUE_USERDEFINELIBRARYLABELSET = "InvalidParameterValue.UserDefineLibraryLabelSet"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) ModifyAIRecognitionTemplateWithContext(ctx context.Context, request *ModifyAIRecognitionTemplateRequest) (response *ModifyAIRecognitionTemplateResponse, err error) {
    if request == nil {
        request = NewModifyAIRecognitionTemplateRequest()
    }
    
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAnimatedGraphicsTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAnimatedGraphicsTemplateResponse()
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyImageSpriteTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyImageSpriteTemplateResponse()
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyPersonSample require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyPersonSampleResponse()
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifySchedule require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyScheduleResponse()
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifySnapshotByTimeOffsetTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifySnapshotByTimeOffsetTemplateResponse()
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
//  INVALIDPARAMETER_PROTOCOL = "InvalidParameter.Protocol"
//  INVALIDPARAMETER_STATE = "InvalidParameter.State"
func (c *Client) ModifyStreamLinkInputWithContext(ctx context.Context, request *ModifyStreamLinkInputRequest) (response *ModifyStreamLinkInputResponse, err error) {
    if request == nil {
        request = NewModifyStreamLinkInputRequest()
    }
    
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyStreamLinkOutputInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyStreamLinkOutputInfoResponse()
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
// 该接口不用于发起网络调用，而是用来帮助生成各个语言平台的 SDK，您可以参考 SDK 的中解析函数的实现事件通知的解析。
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
// 该接口不用于发起网络调用，而是用来帮助生成各个语言平台的 SDK，您可以参考 SDK 的中解析函数的实现事件通知的解析。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INVALIDPARAMETERVALUE_INVALIDCONTENT = "InvalidParameterValue.InvalidContent"
func (c *Client) ParseLiveStreamProcessNotificationWithContext(ctx context.Context, request *ParseLiveStreamProcessNotificationRequest) (response *ParseLiveStreamProcessNotificationResponse, err error) {
    if request == nil {
        request = NewParseLiveStreamProcessNotificationRequest()
    }
    
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
// 该接口不用于发起网络调用，而是用来帮助生成各个语言平台的 SDK，您可以参考 SDK 的中解析函数的实现事件通知的解析。
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
// 该接口不用于发起网络调用，而是用来帮助生成各个语言平台的 SDK，您可以参考 SDK 的中解析函数的实现事件通知的解析。
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("ParseNotification require credential")
    }

    request.SetContext(ctx)
    
    response = NewParseNotificationResponse()
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
// * 智能内容识别（人脸、文本全文、文本关键词、语音全文、语音关键词、语音实时翻译）。
//
// 
//
// 直播流处理事件通知实时写入用户指定的消息队列 CMQ 中，用户需要从消息队列 CMQ 中获取事件通知结果，同时处理过程中存在输出文件的，会写入用户指定的输出文件的目标存储中。
//
// 可能返回的错误码:
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
// * 智能内容识别（人脸、文本全文、文本关键词、语音全文、语音关键词、语音实时翻译）。
//
// 
//
// 直播流处理事件通知实时写入用户指定的消息队列 CMQ 中，用户需要从消息队列 CMQ 中获取事件通知结果，同时处理过程中存在输出文件的，会写入用户指定的输出文件的目标存储中。
//
// 可能返回的错误码:
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
// 1. 视频转码（普通转码、极速高清转码、音视频增强）；
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
// 8. 智能内容分析（标签、分类、封面、按帧标签、拆条、集锦、片头片尾）；
//
// 9. 智能内容识别（人脸、文本全文、文本关键词、语音全文、语音关键词、语音翻译、物体识别）。
//
// 可能返回的错误码:
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
// 1. 视频转码（普通转码、极速高清转码、音视频增强）；
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
// 8. 智能内容分析（标签、分类、封面、按帧标签、拆条、集锦、片头片尾）；
//
// 9. 智能内容识别（人脸、文本全文、文本关键词、语音全文、语音关键词、语音翻译、物体识别）。
//
// 可能返回的错误码:
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("ProcessMedia require credential")
    }

    request.SetContext(ctx)
    
    response = NewProcessMediaResponse()
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("StopStreamLinkFlow require credential")
    }

    request.SetContext(ctx)
    
    response = NewStopStreamLinkFlowResponse()
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("WithdrawsWatermark require credential")
    }

    request.SetContext(ctx)
    
    response = NewWithdrawsWatermarkResponse()
    err = c.Send(request, response)
    return
}
