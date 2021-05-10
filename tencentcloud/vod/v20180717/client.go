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

package v20180717

import (
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2018-07-17"

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


func NewApplyUploadRequest() (request *ApplyUploadRequest) {
    request = &ApplyUploadRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vod", APIVersion, "ApplyUpload")
    return
}

func NewApplyUploadResponse() (response *ApplyUploadResponse) {
    response = &ApplyUploadResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// * 我们强烈建议您使用云点播提供的 [服务端上传 SDK](/document/product/266/9759#1.-.E5.8F.91.E8.B5.B7.E4.B8.8A.E4.BC.A0) 来上传文件。直接调用 API 进行上传的难度和工作量都显著大于使用 SDK。
// * 该接口用于申请媒体文件（和封面文件）的上传，获取文件上传到云点播的元信息（包括上传路径、上传签名等），用于后续上传接口。
// * 上传流程请参考 [服务端上传综述](/document/product/266/9759)。
func (c *Client) ApplyUpload(request *ApplyUploadRequest) (response *ApplyUploadResponse, err error) {
    if request == nil {
        request = NewApplyUploadRequest()
    }
    response = NewApplyUploadResponse()
    err = c.Send(request, response)
    return
}

func NewAttachMediaSubtitlesRequest() (request *AttachMediaSubtitlesRequest) {
    request = &AttachMediaSubtitlesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vod", APIVersion, "AttachMediaSubtitles")
    return
}

func NewAttachMediaSubtitlesResponse() (response *AttachMediaSubtitlesResponse) {
    response = &AttachMediaSubtitlesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 关联媒资字幕，将指定的字幕关联到转自适应码流模板号对应的媒体输出文件中（或解除关联）。
func (c *Client) AttachMediaSubtitles(request *AttachMediaSubtitlesRequest) (response *AttachMediaSubtitlesResponse, err error) {
    if request == nil {
        request = NewAttachMediaSubtitlesRequest()
    }
    response = NewAttachMediaSubtitlesResponse()
    err = c.Send(request, response)
    return
}

func NewCommitUploadRequest() (request *CommitUploadRequest) {
    request = &CommitUploadRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vod", APIVersion, "CommitUpload")
    return
}

func NewCommitUploadResponse() (response *CommitUploadResponse) {
    response = &CommitUploadResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 该接口用于确认媒体文件（和封面文件）上传到腾讯云点播的结果，并存储媒体信息，返回文件的播放地址和文件 ID。
func (c *Client) CommitUpload(request *CommitUploadRequest) (response *CommitUploadResponse, err error) {
    if request == nil {
        request = NewCommitUploadRequest()
    }
    response = NewCommitUploadResponse()
    err = c.Send(request, response)
    return
}

func NewComposeMediaRequest() (request *ComposeMediaRequest) {
    request = &ComposeMediaRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vod", APIVersion, "ComposeMedia")
    return
}

func NewComposeMediaResponse() (response *ComposeMediaResponse) {
    response = &ComposeMediaResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 该接口用于制作媒体文件，可以
// 
// 1. 对一个媒体文件进行剪辑，生成一个新的媒体文件；
// 2. 对多个媒体文件进行裁剪拼接，生成一个新的媒体文件；
// 3. 对多个媒体文件的媒体流进行裁剪拼接，生成一个新的媒体文件。
// 
// 如使用事件通知，事件通知的类型为 [视频合成完成](https://cloud.tencent.com/document/product/266/43000)。
func (c *Client) ComposeMedia(request *ComposeMediaRequest) (response *ComposeMediaResponse, err error) {
    if request == nil {
        request = NewComposeMediaRequest()
    }
    response = NewComposeMediaResponse()
    err = c.Send(request, response)
    return
}

func NewConfirmEventsRequest() (request *ConfirmEventsRequest) {
    request = &ConfirmEventsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vod", APIVersion, "ConfirmEvents")
    return
}

func NewConfirmEventsResponse() (response *ConfirmEventsResponse) {
    response = &ConfirmEventsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// * 开发者调用拉取事件通知，获取到事件后，必须调用该接口来确认消息已经收到；
// * 开发者获取到事件句柄后，等待确认的有效时间为 30 秒，超出 30 秒会报参数错误（4000）；
// * 更多参考事件通知的[可靠回调](https://cloud.tencent.com/document/product/266/33779#.E5.8F.AF.E9.9D.A0.E5.9B.9E.E8.B0.83)。
func (c *Client) ConfirmEvents(request *ConfirmEventsRequest) (response *ConfirmEventsResponse, err error) {
    if request == nil {
        request = NewConfirmEventsRequest()
    }
    response = NewConfirmEventsResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAIAnalysisTemplateRequest() (request *CreateAIAnalysisTemplateRequest) {
    request = &CreateAIAnalysisTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vod", APIVersion, "CreateAIAnalysisTemplate")
    return
}

func NewCreateAIAnalysisTemplateResponse() (response *CreateAIAnalysisTemplateResponse) {
    response = &CreateAIAnalysisTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建用户自定义视频内容分析模板，数量上限：50。
func (c *Client) CreateAIAnalysisTemplate(request *CreateAIAnalysisTemplateRequest) (response *CreateAIAnalysisTemplateResponse, err error) {
    if request == nil {
        request = NewCreateAIAnalysisTemplateRequest()
    }
    response = NewCreateAIAnalysisTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAIRecognitionTemplateRequest() (request *CreateAIRecognitionTemplateRequest) {
    request = &CreateAIRecognitionTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vod", APIVersion, "CreateAIRecognitionTemplate")
    return
}

func NewCreateAIRecognitionTemplateResponse() (response *CreateAIRecognitionTemplateResponse) {
    response = &CreateAIRecognitionTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建用户自定义视频内容识别模板，数量上限：50。
func (c *Client) CreateAIRecognitionTemplate(request *CreateAIRecognitionTemplateRequest) (response *CreateAIRecognitionTemplateResponse, err error) {
    if request == nil {
        request = NewCreateAIRecognitionTemplateRequest()
    }
    response = NewCreateAIRecognitionTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAdaptiveDynamicStreamingTemplateRequest() (request *CreateAdaptiveDynamicStreamingTemplateRequest) {
    request = &CreateAdaptiveDynamicStreamingTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vod", APIVersion, "CreateAdaptiveDynamicStreamingTemplate")
    return
}

func NewCreateAdaptiveDynamicStreamingTemplateResponse() (response *CreateAdaptiveDynamicStreamingTemplateResponse) {
    response = &CreateAdaptiveDynamicStreamingTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建转自适应码流模板，数量上限：100。
func (c *Client) CreateAdaptiveDynamicStreamingTemplate(request *CreateAdaptiveDynamicStreamingTemplateRequest) (response *CreateAdaptiveDynamicStreamingTemplateResponse, err error) {
    if request == nil {
        request = NewCreateAdaptiveDynamicStreamingTemplateRequest()
    }
    response = NewCreateAdaptiveDynamicStreamingTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAnimatedGraphicsTemplateRequest() (request *CreateAnimatedGraphicsTemplateRequest) {
    request = &CreateAnimatedGraphicsTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vod", APIVersion, "CreateAnimatedGraphicsTemplate")
    return
}

func NewCreateAnimatedGraphicsTemplateResponse() (response *CreateAnimatedGraphicsTemplateResponse) {
    response = &CreateAnimatedGraphicsTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建用户自定义转动图模板，数量上限：16。
func (c *Client) CreateAnimatedGraphicsTemplate(request *CreateAnimatedGraphicsTemplateRequest) (response *CreateAnimatedGraphicsTemplateResponse, err error) {
    if request == nil {
        request = NewCreateAnimatedGraphicsTemplateRequest()
    }
    response = NewCreateAnimatedGraphicsTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewCreateClassRequest() (request *CreateClassRequest) {
    request = &CreateClassRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vod", APIVersion, "CreateClass")
    return
}

func NewCreateClassResponse() (response *CreateClassResponse) {
    response = &CreateClassResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// * 用于对媒体进行分类管理；
// * 该接口不影响既有媒体的分类，如需修改媒体分类，请调用[修改媒体文件属性](/document/product/266/31762)接口。
// * 分类层次不可超过 4 层。
// * 每个分类的子类数量不可超过 500 个。
func (c *Client) CreateClass(request *CreateClassRequest) (response *CreateClassResponse, err error) {
    if request == nil {
        request = NewCreateClassRequest()
    }
    response = NewCreateClassResponse()
    err = c.Send(request, response)
    return
}

func NewCreateContentReviewTemplateRequest() (request *CreateContentReviewTemplateRequest) {
    request = &CreateContentReviewTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vod", APIVersion, "CreateContentReviewTemplate")
    return
}

func NewCreateContentReviewTemplateResponse() (response *CreateContentReviewTemplateResponse) {
    response = &CreateContentReviewTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建用户自定义视频内容智能识别模板，数量上限：50。
func (c *Client) CreateContentReviewTemplate(request *CreateContentReviewTemplateRequest) (response *CreateContentReviewTemplateResponse, err error) {
    if request == nil {
        request = NewCreateContentReviewTemplateRequest()
    }
    response = NewCreateContentReviewTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewCreateHeadTailTemplateRequest() (request *CreateHeadTailTemplateRequest) {
    request = &CreateHeadTailTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vod", APIVersion, "CreateHeadTailTemplate")
    return
}

func NewCreateHeadTailTemplateResponse() (response *CreateHeadTailTemplateResponse) {
    response = &CreateHeadTailTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建片头片尾模板。
func (c *Client) CreateHeadTailTemplate(request *CreateHeadTailTemplateRequest) (response *CreateHeadTailTemplateResponse, err error) {
    if request == nil {
        request = NewCreateHeadTailTemplateRequest()
    }
    response = NewCreateHeadTailTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewCreateImageProcessingTemplateRequest() (request *CreateImageProcessingTemplateRequest) {
    request = &CreateImageProcessingTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vod", APIVersion, "CreateImageProcessingTemplate")
    return
}

func NewCreateImageProcessingTemplateResponse() (response *CreateImageProcessingTemplateResponse) {
    response = &CreateImageProcessingTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建一个用户自定义的图片处理模板，数量上限：16。最多支持三次操作，例如：裁剪-缩略-裁剪。
func (c *Client) CreateImageProcessingTemplate(request *CreateImageProcessingTemplateRequest) (response *CreateImageProcessingTemplateResponse, err error) {
    if request == nil {
        request = NewCreateImageProcessingTemplateRequest()
    }
    response = NewCreateImageProcessingTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewCreateImageSpriteTemplateRequest() (request *CreateImageSpriteTemplateRequest) {
    request = &CreateImageSpriteTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vod", APIVersion, "CreateImageSpriteTemplate")
    return
}

func NewCreateImageSpriteTemplateResponse() (response *CreateImageSpriteTemplateResponse) {
    response = &CreateImageSpriteTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建用户自定义雪碧图模板，数量上限：16。
func (c *Client) CreateImageSpriteTemplate(request *CreateImageSpriteTemplateRequest) (response *CreateImageSpriteTemplateResponse, err error) {
    if request == nil {
        request = NewCreateImageSpriteTemplateRequest()
    }
    response = NewCreateImageSpriteTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewCreatePersonSampleRequest() (request *CreatePersonSampleRequest) {
    request = &CreatePersonSampleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vod", APIVersion, "CreatePersonSample")
    return
}

func NewCreatePersonSampleResponse() (response *CreatePersonSampleResponse) {
    response = &CreatePersonSampleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 该接口用于创建素材样本，用于通过五官定位等技术，进行内容识别、不适宜视频识别等视频处理。
func (c *Client) CreatePersonSample(request *CreatePersonSampleRequest) (response *CreatePersonSampleResponse, err error) {
    if request == nil {
        request = NewCreatePersonSampleRequest()
    }
    response = NewCreatePersonSampleResponse()
    err = c.Send(request, response)
    return
}

func NewCreateProcedureTemplateRequest() (request *CreateProcedureTemplateRequest) {
    request = &CreateProcedureTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vod", APIVersion, "CreateProcedureTemplate")
    return
}

func NewCreateProcedureTemplateResponse() (response *CreateProcedureTemplateResponse) {
    response = &CreateProcedureTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建用户自定义的任务流模板，模板上限：50。
func (c *Client) CreateProcedureTemplate(request *CreateProcedureTemplateRequest) (response *CreateProcedureTemplateResponse, err error) {
    if request == nil {
        request = NewCreateProcedureTemplateRequest()
    }
    response = NewCreateProcedureTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewCreateSampleSnapshotTemplateRequest() (request *CreateSampleSnapshotTemplateRequest) {
    request = &CreateSampleSnapshotTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vod", APIVersion, "CreateSampleSnapshotTemplate")
    return
}

func NewCreateSampleSnapshotTemplateResponse() (response *CreateSampleSnapshotTemplateResponse) {
    response = &CreateSampleSnapshotTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建用户自定义采样截图模板，数量上限：16。
func (c *Client) CreateSampleSnapshotTemplate(request *CreateSampleSnapshotTemplateRequest) (response *CreateSampleSnapshotTemplateResponse, err error) {
    if request == nil {
        request = NewCreateSampleSnapshotTemplateRequest()
    }
    response = NewCreateSampleSnapshotTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewCreateSnapshotByTimeOffsetTemplateRequest() (request *CreateSnapshotByTimeOffsetTemplateRequest) {
    request = &CreateSnapshotByTimeOffsetTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vod", APIVersion, "CreateSnapshotByTimeOffsetTemplate")
    return
}

func NewCreateSnapshotByTimeOffsetTemplateResponse() (response *CreateSnapshotByTimeOffsetTemplateResponse) {
    response = &CreateSnapshotByTimeOffsetTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建用户自定义指定时间点截图模板，数量上限：16。
func (c *Client) CreateSnapshotByTimeOffsetTemplate(request *CreateSnapshotByTimeOffsetTemplateRequest) (response *CreateSnapshotByTimeOffsetTemplateResponse, err error) {
    if request == nil {
        request = NewCreateSnapshotByTimeOffsetTemplateRequest()
    }
    response = NewCreateSnapshotByTimeOffsetTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewCreateSubAppIdRequest() (request *CreateSubAppIdRequest) {
    request = &CreateSubAppIdRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vod", APIVersion, "CreateSubAppId")
    return
}

func NewCreateSubAppIdResponse() (response *CreateSubAppIdResponse) {
    response = &CreateSubAppIdResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 该接口用于创建点播子应用。
func (c *Client) CreateSubAppId(request *CreateSubAppIdRequest) (response *CreateSubAppIdResponse, err error) {
    if request == nil {
        request = NewCreateSubAppIdRequest()
    }
    response = NewCreateSubAppIdResponse()
    err = c.Send(request, response)
    return
}

func NewCreateSuperPlayerConfigRequest() (request *CreateSuperPlayerConfigRequest) {
    request = &CreateSuperPlayerConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vod", APIVersion, "CreateSuperPlayerConfig")
    return
}

func NewCreateSuperPlayerConfigResponse() (response *CreateSuperPlayerConfigResponse) {
    response = &CreateSuperPlayerConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建超级播放器配置，数量上限：100。
func (c *Client) CreateSuperPlayerConfig(request *CreateSuperPlayerConfigRequest) (response *CreateSuperPlayerConfigResponse, err error) {
    if request == nil {
        request = NewCreateSuperPlayerConfigRequest()
    }
    response = NewCreateSuperPlayerConfigResponse()
    err = c.Send(request, response)
    return
}

func NewCreateTranscodeTemplateRequest() (request *CreateTranscodeTemplateRequest) {
    request = &CreateTranscodeTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vod", APIVersion, "CreateTranscodeTemplate")
    return
}

func NewCreateTranscodeTemplateResponse() (response *CreateTranscodeTemplateResponse) {
    response = &CreateTranscodeTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建用户自定义转码模板，数量上限：100。
func (c *Client) CreateTranscodeTemplate(request *CreateTranscodeTemplateRequest) (response *CreateTranscodeTemplateResponse, err error) {
    if request == nil {
        request = NewCreateTranscodeTemplateRequest()
    }
    response = NewCreateTranscodeTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewCreateWatermarkTemplateRequest() (request *CreateWatermarkTemplateRequest) {
    request = &CreateWatermarkTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vod", APIVersion, "CreateWatermarkTemplate")
    return
}

func NewCreateWatermarkTemplateResponse() (response *CreateWatermarkTemplateResponse) {
    response = &CreateWatermarkTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建用户自定义水印模板，数量上限：1000。
func (c *Client) CreateWatermarkTemplate(request *CreateWatermarkTemplateRequest) (response *CreateWatermarkTemplateResponse, err error) {
    if request == nil {
        request = NewCreateWatermarkTemplateRequest()
    }
    response = NewCreateWatermarkTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewCreateWordSamplesRequest() (request *CreateWordSamplesRequest) {
    request = &CreateWordSamplesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vod", APIVersion, "CreateWordSamples")
    return
}

func NewCreateWordSamplesResponse() (response *CreateWordSamplesResponse) {
    response = &CreateWordSamplesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 该接口用于批量创建关键词样本，样本用于通过OCR、ASR技术，进行不适宜内容识别、内容识别等视频处理。
func (c *Client) CreateWordSamples(request *CreateWordSamplesRequest) (response *CreateWordSamplesResponse, err error) {
    if request == nil {
        request = NewCreateWordSamplesRequest()
    }
    response = NewCreateWordSamplesResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAIAnalysisTemplateRequest() (request *DeleteAIAnalysisTemplateRequest) {
    request = &DeleteAIAnalysisTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vod", APIVersion, "DeleteAIAnalysisTemplate")
    return
}

func NewDeleteAIAnalysisTemplateResponse() (response *DeleteAIAnalysisTemplateResponse) {
    response = &DeleteAIAnalysisTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除用户自定义视频内容分析模板。
// 
// 注意：模板 ID 为 10000 以下的为系统预置模板，不允许删除。
func (c *Client) DeleteAIAnalysisTemplate(request *DeleteAIAnalysisTemplateRequest) (response *DeleteAIAnalysisTemplateResponse, err error) {
    if request == nil {
        request = NewDeleteAIAnalysisTemplateRequest()
    }
    response = NewDeleteAIAnalysisTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAIRecognitionTemplateRequest() (request *DeleteAIRecognitionTemplateRequest) {
    request = &DeleteAIRecognitionTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vod", APIVersion, "DeleteAIRecognitionTemplate")
    return
}

func NewDeleteAIRecognitionTemplateResponse() (response *DeleteAIRecognitionTemplateResponse) {
    response = &DeleteAIRecognitionTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除用户自定义视频内容识别模板。
func (c *Client) DeleteAIRecognitionTemplate(request *DeleteAIRecognitionTemplateRequest) (response *DeleteAIRecognitionTemplateResponse, err error) {
    if request == nil {
        request = NewDeleteAIRecognitionTemplateRequest()
    }
    response = NewDeleteAIRecognitionTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAdaptiveDynamicStreamingTemplateRequest() (request *DeleteAdaptiveDynamicStreamingTemplateRequest) {
    request = &DeleteAdaptiveDynamicStreamingTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vod", APIVersion, "DeleteAdaptiveDynamicStreamingTemplate")
    return
}

func NewDeleteAdaptiveDynamicStreamingTemplateResponse() (response *DeleteAdaptiveDynamicStreamingTemplateResponse) {
    response = &DeleteAdaptiveDynamicStreamingTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除转自适应码流模板
func (c *Client) DeleteAdaptiveDynamicStreamingTemplate(request *DeleteAdaptiveDynamicStreamingTemplateRequest) (response *DeleteAdaptiveDynamicStreamingTemplateResponse, err error) {
    if request == nil {
        request = NewDeleteAdaptiveDynamicStreamingTemplateRequest()
    }
    response = NewDeleteAdaptiveDynamicStreamingTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAnimatedGraphicsTemplateRequest() (request *DeleteAnimatedGraphicsTemplateRequest) {
    request = &DeleteAnimatedGraphicsTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vod", APIVersion, "DeleteAnimatedGraphicsTemplate")
    return
}

func NewDeleteAnimatedGraphicsTemplateResponse() (response *DeleteAnimatedGraphicsTemplateResponse) {
    response = &DeleteAnimatedGraphicsTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除用户自定义转动图模板。
func (c *Client) DeleteAnimatedGraphicsTemplate(request *DeleteAnimatedGraphicsTemplateRequest) (response *DeleteAnimatedGraphicsTemplateResponse, err error) {
    if request == nil {
        request = NewDeleteAnimatedGraphicsTemplateRequest()
    }
    response = NewDeleteAnimatedGraphicsTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteClassRequest() (request *DeleteClassRequest) {
    request = &DeleteClassRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vod", APIVersion, "DeleteClass")
    return
}

func NewDeleteClassResponse() (response *DeleteClassResponse) {
    response = &DeleteClassResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// * 仅当待删分类无子分类且无媒体关联情况下，可删除分类；
// * 否则，请先执行[删除媒体](/document/product/266/31764)及子分类，再删除该分类；
func (c *Client) DeleteClass(request *DeleteClassRequest) (response *DeleteClassResponse, err error) {
    if request == nil {
        request = NewDeleteClassRequest()
    }
    response = NewDeleteClassResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteContentReviewTemplateRequest() (request *DeleteContentReviewTemplateRequest) {
    request = &DeleteContentReviewTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vod", APIVersion, "DeleteContentReviewTemplate")
    return
}

func NewDeleteContentReviewTemplateResponse() (response *DeleteContentReviewTemplateResponse) {
    response = &DeleteContentReviewTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除用户自定义视频内容智能识别模板。
func (c *Client) DeleteContentReviewTemplate(request *DeleteContentReviewTemplateRequest) (response *DeleteContentReviewTemplateResponse, err error) {
    if request == nil {
        request = NewDeleteContentReviewTemplateRequest()
    }
    response = NewDeleteContentReviewTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteHeadTailTemplateRequest() (request *DeleteHeadTailTemplateRequest) {
    request = &DeleteHeadTailTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vod", APIVersion, "DeleteHeadTailTemplate")
    return
}

func NewDeleteHeadTailTemplateResponse() (response *DeleteHeadTailTemplateResponse) {
    response = &DeleteHeadTailTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除片头片尾模板。
func (c *Client) DeleteHeadTailTemplate(request *DeleteHeadTailTemplateRequest) (response *DeleteHeadTailTemplateResponse, err error) {
    if request == nil {
        request = NewDeleteHeadTailTemplateRequest()
    }
    response = NewDeleteHeadTailTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteImageProcessingTemplateRequest() (request *DeleteImageProcessingTemplateRequest) {
    request = &DeleteImageProcessingTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vod", APIVersion, "DeleteImageProcessingTemplate")
    return
}

func NewDeleteImageProcessingTemplateResponse() (response *DeleteImageProcessingTemplateResponse) {
    response = &DeleteImageProcessingTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除用户自定义图片处理模板。
func (c *Client) DeleteImageProcessingTemplate(request *DeleteImageProcessingTemplateRequest) (response *DeleteImageProcessingTemplateResponse, err error) {
    if request == nil {
        request = NewDeleteImageProcessingTemplateRequest()
    }
    response = NewDeleteImageProcessingTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteImageSpriteTemplateRequest() (request *DeleteImageSpriteTemplateRequest) {
    request = &DeleteImageSpriteTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vod", APIVersion, "DeleteImageSpriteTemplate")
    return
}

func NewDeleteImageSpriteTemplateResponse() (response *DeleteImageSpriteTemplateResponse) {
    response = &DeleteImageSpriteTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除雪碧图模板。
func (c *Client) DeleteImageSpriteTemplate(request *DeleteImageSpriteTemplateRequest) (response *DeleteImageSpriteTemplateResponse, err error) {
    if request == nil {
        request = NewDeleteImageSpriteTemplateRequest()
    }
    response = NewDeleteImageSpriteTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteMediaRequest() (request *DeleteMediaRequest) {
    request = &DeleteMediaRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vod", APIVersion, "DeleteMedia")
    return
}

func NewDeleteMediaResponse() (response *DeleteMediaResponse) {
    response = &DeleteMediaResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// * 删除媒体及其对应的视频处理文件（如转码视频、雪碧图、截图、微信发布视频等）；
// * 可单独删除指定 ID 的视频文件下的转码，或者微信发布文件；
func (c *Client) DeleteMedia(request *DeleteMediaRequest) (response *DeleteMediaResponse, err error) {
    if request == nil {
        request = NewDeleteMediaRequest()
    }
    response = NewDeleteMediaResponse()
    err = c.Send(request, response)
    return
}

func NewDeletePersonSampleRequest() (request *DeletePersonSampleRequest) {
    request = &DeletePersonSampleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vod", APIVersion, "DeletePersonSample")
    return
}

func NewDeletePersonSampleResponse() (response *DeletePersonSampleResponse) {
    response = &DeletePersonSampleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 该接口用于根据人物 ID，删除素材样本。
func (c *Client) DeletePersonSample(request *DeletePersonSampleRequest) (response *DeletePersonSampleResponse, err error) {
    if request == nil {
        request = NewDeletePersonSampleRequest()
    }
    response = NewDeletePersonSampleResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteProcedureTemplateRequest() (request *DeleteProcedureTemplateRequest) {
    request = &DeleteProcedureTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vod", APIVersion, "DeleteProcedureTemplate")
    return
}

func NewDeleteProcedureTemplateResponse() (response *DeleteProcedureTemplateResponse) {
    response = &DeleteProcedureTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除用户自定义的任务流模板。  
func (c *Client) DeleteProcedureTemplate(request *DeleteProcedureTemplateRequest) (response *DeleteProcedureTemplateResponse, err error) {
    if request == nil {
        request = NewDeleteProcedureTemplateRequest()
    }
    response = NewDeleteProcedureTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteSampleSnapshotTemplateRequest() (request *DeleteSampleSnapshotTemplateRequest) {
    request = &DeleteSampleSnapshotTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vod", APIVersion, "DeleteSampleSnapshotTemplate")
    return
}

func NewDeleteSampleSnapshotTemplateResponse() (response *DeleteSampleSnapshotTemplateResponse) {
    response = &DeleteSampleSnapshotTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除用户自定义采样截图模板。
func (c *Client) DeleteSampleSnapshotTemplate(request *DeleteSampleSnapshotTemplateRequest) (response *DeleteSampleSnapshotTemplateResponse, err error) {
    if request == nil {
        request = NewDeleteSampleSnapshotTemplateRequest()
    }
    response = NewDeleteSampleSnapshotTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteSnapshotByTimeOffsetTemplateRequest() (request *DeleteSnapshotByTimeOffsetTemplateRequest) {
    request = &DeleteSnapshotByTimeOffsetTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vod", APIVersion, "DeleteSnapshotByTimeOffsetTemplate")
    return
}

func NewDeleteSnapshotByTimeOffsetTemplateResponse() (response *DeleteSnapshotByTimeOffsetTemplateResponse) {
    response = &DeleteSnapshotByTimeOffsetTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除用户自定义指定时间点截图模板。
func (c *Client) DeleteSnapshotByTimeOffsetTemplate(request *DeleteSnapshotByTimeOffsetTemplateRequest) (response *DeleteSnapshotByTimeOffsetTemplateResponse, err error) {
    if request == nil {
        request = NewDeleteSnapshotByTimeOffsetTemplateRequest()
    }
    response = NewDeleteSnapshotByTimeOffsetTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteSuperPlayerConfigRequest() (request *DeleteSuperPlayerConfigRequest) {
    request = &DeleteSuperPlayerConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vod", APIVersion, "DeleteSuperPlayerConfig")
    return
}

func NewDeleteSuperPlayerConfigResponse() (response *DeleteSuperPlayerConfigResponse) {
    response = &DeleteSuperPlayerConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除超级播放器配置。  
// *注：系统预置播放器配置不允许删除。*
func (c *Client) DeleteSuperPlayerConfig(request *DeleteSuperPlayerConfigRequest) (response *DeleteSuperPlayerConfigResponse, err error) {
    if request == nil {
        request = NewDeleteSuperPlayerConfigRequest()
    }
    response = NewDeleteSuperPlayerConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteTranscodeTemplateRequest() (request *DeleteTranscodeTemplateRequest) {
    request = &DeleteTranscodeTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vod", APIVersion, "DeleteTranscodeTemplate")
    return
}

func NewDeleteTranscodeTemplateResponse() (response *DeleteTranscodeTemplateResponse) {
    response = &DeleteTranscodeTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除用户自定义转码模板。
func (c *Client) DeleteTranscodeTemplate(request *DeleteTranscodeTemplateRequest) (response *DeleteTranscodeTemplateResponse, err error) {
    if request == nil {
        request = NewDeleteTranscodeTemplateRequest()
    }
    response = NewDeleteTranscodeTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteWatermarkTemplateRequest() (request *DeleteWatermarkTemplateRequest) {
    request = &DeleteWatermarkTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vod", APIVersion, "DeleteWatermarkTemplate")
    return
}

func NewDeleteWatermarkTemplateResponse() (response *DeleteWatermarkTemplateResponse) {
    response = &DeleteWatermarkTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除用户自定义水印模板。
func (c *Client) DeleteWatermarkTemplate(request *DeleteWatermarkTemplateRequest) (response *DeleteWatermarkTemplateResponse, err error) {
    if request == nil {
        request = NewDeleteWatermarkTemplateRequest()
    }
    response = NewDeleteWatermarkTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteWordSamplesRequest() (request *DeleteWordSamplesRequest) {
    request = &DeleteWordSamplesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vod", APIVersion, "DeleteWordSamples")
    return
}

func NewDeleteWordSamplesResponse() (response *DeleteWordSamplesResponse) {
    response = &DeleteWordSamplesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 该接口用于批量删除关键词样本。
func (c *Client) DeleteWordSamples(request *DeleteWordSamplesRequest) (response *DeleteWordSamplesResponse, err error) {
    if request == nil {
        request = NewDeleteWordSamplesRequest()
    }
    response = NewDeleteWordSamplesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAIAnalysisTemplatesRequest() (request *DescribeAIAnalysisTemplatesRequest) {
    request = &DescribeAIAnalysisTemplatesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vod", APIVersion, "DescribeAIAnalysisTemplates")
    return
}

func NewDescribeAIAnalysisTemplatesResponse() (response *DescribeAIAnalysisTemplatesResponse) {
    response = &DescribeAIAnalysisTemplatesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 根据视频内容分析模板唯一标识，获取视频内容分析模板详情列表。返回结果包含符合条件的所有用户自定义视频内容分析模板及[系统预置视频内容分析模板](https://cloud.tencent.com/document/product/266/33476#.E9.A2.84.E7.BD.AE.E8.A7.86.E9.A2.91.E5.86.85.E5.AE.B9.E5.88.86.E6.9E.90.E6.A8.A1.E6.9D.BF)。
func (c *Client) DescribeAIAnalysisTemplates(request *DescribeAIAnalysisTemplatesRequest) (response *DescribeAIAnalysisTemplatesResponse, err error) {
    if request == nil {
        request = NewDescribeAIAnalysisTemplatesRequest()
    }
    response = NewDescribeAIAnalysisTemplatesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAIRecognitionTemplatesRequest() (request *DescribeAIRecognitionTemplatesRequest) {
    request = &DescribeAIRecognitionTemplatesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vod", APIVersion, "DescribeAIRecognitionTemplates")
    return
}

func NewDescribeAIRecognitionTemplatesResponse() (response *DescribeAIRecognitionTemplatesResponse) {
    response = &DescribeAIRecognitionTemplatesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 根据视频内容识别模板唯一标识，获取视频内容识别模板详情列表。返回结果包含符合条件的所有用户自定义视频内容识别模板及[系统预置视频内容识别模板](https://cloud.tencent.com/document/product/266/33476#.E9.A2.84.E7.BD.AE.E8.A7.86.E9.A2.91.E5.86.85.E5.AE.B9.E8.AF.86.E5.88.AB.E6.A8.A1.E6.9D.BF)。
func (c *Client) DescribeAIRecognitionTemplates(request *DescribeAIRecognitionTemplatesRequest) (response *DescribeAIRecognitionTemplatesResponse, err error) {
    if request == nil {
        request = NewDescribeAIRecognitionTemplatesRequest()
    }
    response = NewDescribeAIRecognitionTemplatesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAdaptiveDynamicStreamingTemplatesRequest() (request *DescribeAdaptiveDynamicStreamingTemplatesRequest) {
    request = &DescribeAdaptiveDynamicStreamingTemplatesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vod", APIVersion, "DescribeAdaptiveDynamicStreamingTemplates")
    return
}

func NewDescribeAdaptiveDynamicStreamingTemplatesResponse() (response *DescribeAdaptiveDynamicStreamingTemplatesResponse) {
    response = &DescribeAdaptiveDynamicStreamingTemplatesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询转自适应码流模板，支持根据条件，分页查询。
func (c *Client) DescribeAdaptiveDynamicStreamingTemplates(request *DescribeAdaptiveDynamicStreamingTemplatesRequest) (response *DescribeAdaptiveDynamicStreamingTemplatesResponse, err error) {
    if request == nil {
        request = NewDescribeAdaptiveDynamicStreamingTemplatesRequest()
    }
    response = NewDescribeAdaptiveDynamicStreamingTemplatesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAllClassRequest() (request *DescribeAllClassRequest) {
    request = &DescribeAllClassRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vod", APIVersion, "DescribeAllClass")
    return
}

func NewDescribeAllClassResponse() (response *DescribeAllClassResponse) {
    response = &DescribeAllClassResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// * 获得用户的所有分类信息。
func (c *Client) DescribeAllClass(request *DescribeAllClassRequest) (response *DescribeAllClassResponse, err error) {
    if request == nil {
        request = NewDescribeAllClassRequest()
    }
    response = NewDescribeAllClassResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAnimatedGraphicsTemplatesRequest() (request *DescribeAnimatedGraphicsTemplatesRequest) {
    request = &DescribeAnimatedGraphicsTemplatesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vod", APIVersion, "DescribeAnimatedGraphicsTemplates")
    return
}

func NewDescribeAnimatedGraphicsTemplatesResponse() (response *DescribeAnimatedGraphicsTemplatesResponse) {
    response = &DescribeAnimatedGraphicsTemplatesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询转动图模板列表，支持根据条件，分页查询。
func (c *Client) DescribeAnimatedGraphicsTemplates(request *DescribeAnimatedGraphicsTemplatesRequest) (response *DescribeAnimatedGraphicsTemplatesResponse, err error) {
    if request == nil {
        request = NewDescribeAnimatedGraphicsTemplatesRequest()
    }
    response = NewDescribeAnimatedGraphicsTemplatesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCDNStatDetailsRequest() (request *DescribeCDNStatDetailsRequest) {
    request = &DescribeCDNStatDetailsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vod", APIVersion, "DescribeCDNStatDetails")
    return
}

func NewDescribeCDNStatDetailsResponse() (response *DescribeCDNStatDetailsResponse) {
    response = &DescribeCDNStatDetailsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 该接口用于查询点播域名的 CDN 带宽、流量等统计数据。
// * 查询的起始时间和结束时间跨度不超过90天。
// * 可以查询不同服务区域的数据。
// * 中国境内的数据支持查询指定地区、运营商的统计数据。
func (c *Client) DescribeCDNStatDetails(request *DescribeCDNStatDetailsRequest) (response *DescribeCDNStatDetailsResponse, err error) {
    if request == nil {
        request = NewDescribeCDNStatDetailsRequest()
    }
    response = NewDescribeCDNStatDetailsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCDNUsageDataRequest() (request *DescribeCDNUsageDataRequest) {
    request = &DescribeCDNUsageDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vod", APIVersion, "DescribeCDNUsageData")
    return
}

func NewDescribeCDNUsageDataResponse() (response *DescribeCDNUsageDataResponse) {
    response = &DescribeCDNUsageDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 该接口用于查询点播 CDN 的流量、带宽等统计数据。
//    1. 可以查询最近365天内的 CDN 用量数据。
//    2.  查询时间跨度不超过90天。
//    3. 可以指定用量数据的时间粒度，支持5分钟、1小时、1天的时间粒度。
//    4.  流量为查询时间粒度内的总流量，带宽为查询时间粒度内的峰值带宽。
func (c *Client) DescribeCDNUsageData(request *DescribeCDNUsageDataRequest) (response *DescribeCDNUsageDataResponse, err error) {
    if request == nil {
        request = NewDescribeCDNUsageDataRequest()
    }
    response = NewDescribeCDNUsageDataResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCdnLogsRequest() (request *DescribeCdnLogsRequest) {
    request = &DescribeCdnLogsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vod", APIVersion, "DescribeCdnLogs")
    return
}

func NewDescribeCdnLogsResponse() (response *DescribeCdnLogsResponse) {
    response = &DescribeCdnLogsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询点播域名的 CDN 访问日志的下载链接。
//     1. 可以查询最近30天内的 CDN 日志下载链接。
//     2. 默认情况下 CDN 每小时生成一个日志文件，如果某一个小时没有 CDN 访问，不会生成日志文件。    
//     3. CDN 日志下载链接的有效期为24小时。
func (c *Client) DescribeCdnLogs(request *DescribeCdnLogsRequest) (response *DescribeCdnLogsResponse, err error) {
    if request == nil {
        request = NewDescribeCdnLogsRequest()
    }
    response = NewDescribeCdnLogsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeContentReviewTemplatesRequest() (request *DescribeContentReviewTemplatesRequest) {
    request = &DescribeContentReviewTemplatesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vod", APIVersion, "DescribeContentReviewTemplates")
    return
}

func NewDescribeContentReviewTemplatesResponse() (response *DescribeContentReviewTemplatesResponse) {
    response = &DescribeContentReviewTemplatesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 根据视频内容智能识别模板唯一标识，获取视频内容智能识别模板详情列表。返回结果包含符合条件的所有用户自定义模板及[系统预置内容智能识别模板](https://cloud.tencent.com/document/product/266/33476#.E9.A2.84.E7.BD.AE.E8.A7.86.E9.A2.91.E5.86.85.E5.AE.B9.E5.AE.A1.E6.A0.B8.E6.A8.A1.E6.9D.BF)。
func (c *Client) DescribeContentReviewTemplates(request *DescribeContentReviewTemplatesRequest) (response *DescribeContentReviewTemplatesResponse, err error) {
    if request == nil {
        request = NewDescribeContentReviewTemplatesRequest()
    }
    response = NewDescribeContentReviewTemplatesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDailyMediaPlayStatRequest() (request *DescribeDailyMediaPlayStatRequest) {
    request = &DescribeDailyMediaPlayStatRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vod", APIVersion, "DescribeDailyMediaPlayStat")
    return
}

func NewDescribeDailyMediaPlayStatResponse() (response *DescribeDailyMediaPlayStatResponse) {
    response = &DescribeDailyMediaPlayStatResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 该接口用于查询指定日期范围内每天的播放统计数据。
// * 可以查询最近30天的播放统计数据。
func (c *Client) DescribeDailyMediaPlayStat(request *DescribeDailyMediaPlayStatRequest) (response *DescribeDailyMediaPlayStatResponse, err error) {
    if request == nil {
        request = NewDescribeDailyMediaPlayStatRequest()
    }
    response = NewDescribeDailyMediaPlayStatResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDailyMostPlayedStatRequest() (request *DescribeDailyMostPlayedStatRequest) {
    request = &DescribeDailyMostPlayedStatRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vod", APIVersion, "DescribeDailyMostPlayedStat")
    return
}

func NewDescribeDailyMostPlayedStatResponse() (response *DescribeDailyMostPlayedStatResponse) {
    response = &DescribeDailyMostPlayedStatResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 该接口用于查询每日播放Top100 的媒体文件的播放统计数据。
// * 可以查询最近30天的播放统计数据。
// * 可以按播放次数或者播放流量查询。
// * 播放次数统计说明：
//     1. HLS 文件：访问 M3U8 文件时统计播放次数；访问 TS 文件不统计播放次数。
//     2. 其它文件（如 MP4 文件）：播放请求带有 range 参数且 range 的 start 参数不等于0时不统计播放次数，其它情况统计播放次数。
func (c *Client) DescribeDailyMostPlayedStat(request *DescribeDailyMostPlayedStatRequest) (response *DescribeDailyMostPlayedStatResponse, err error) {
    if request == nil {
        request = NewDescribeDailyMostPlayedStatRequest()
    }
    response = NewDescribeDailyMostPlayedStatResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDailyPlayStatFileListRequest() (request *DescribeDailyPlayStatFileListRequest) {
    request = &DescribeDailyPlayStatFileListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vod", APIVersion, "DescribeDailyPlayStatFileList")
    return
}

func NewDescribeDailyPlayStatFileListResponse() (response *DescribeDailyPlayStatFileListResponse) {
    response = &DescribeDailyPlayStatFileListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 该接口用于查询播放统计文件的下载地址。
// * 可以查询最近30天的播放统计文件下载地址。
// * 云点播每天对前一天的 CDN 请求日志进行分析处理，生成播放统计文件。
// * 播放统计文件内容包含媒体文件的播放次数、播放流量等统计信息。
// * 播放次数统计说明：
//     1. HLS 文件：访问M3U8 文件时统计播放次数；访问TS 文件不统计播放次数。
//     2. 其它文件（如 MP4 文件）：播放请求带有 range 参数且 range 的 start 参数不等于0时不统计播放次数，其它情况统计播放次数。
// * 播放设备的统计：播放请求带了 UserAgent 参数，并且 UserAgent 包含 Android 或者 iPhone 等标识，会统计为移动端播放次数，否则统计为 PC 端播放次数。
func (c *Client) DescribeDailyPlayStatFileList(request *DescribeDailyPlayStatFileListRequest) (response *DescribeDailyPlayStatFileListResponse, err error) {
    if request == nil {
        request = NewDescribeDailyPlayStatFileListRequest()
    }
    response = NewDescribeDailyPlayStatFileListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDrmDataKeyRequest() (request *DescribeDrmDataKeyRequest) {
    request = &DescribeDrmDataKeyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vod", APIVersion, "DescribeDrmDataKey")
    return
}

func NewDescribeDrmDataKeyResponse() (response *DescribeDrmDataKeyResponse) {
    response = &DescribeDrmDataKeyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本 API 是 [旧版本加密](https://cloud.tencent.com/document/product/266/9638) 中[DescribeDrmDataKey 的 API 2017 接口](https://cloud.tencent.com/document/product/266/9643)的升级版本。
// 如果您是新接入点播加密的用户，不要使用该 API。请参考[视频加密综述](https://cloud.tencent.com/document/product/266/45552)使用。
func (c *Client) DescribeDrmDataKey(request *DescribeDrmDataKeyRequest) (response *DescribeDrmDataKeyResponse, err error) {
    if request == nil {
        request = NewDescribeDrmDataKeyRequest()
    }
    response = NewDescribeDrmDataKeyResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEventConfigRequest() (request *DescribeEventConfigRequest) {
    request = &DescribeEventConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vod", APIVersion, "DescribeEventConfig")
    return
}

func NewDescribeEventConfigResponse() (response *DescribeEventConfigResponse) {
    response = &DescribeEventConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 腾讯云点播为客户提供了媒体上传、媒体管理、媒体处理等等服务，在这些服务执行过程或执行结束时，腾讯云点播也提供各种对应的事件通知，方便开发者感知服务处理状态，并做下一步的业务操作。
// 
// 开发者可以通过本接口来查询当前配置事件通知的接收方式、接收地址以及哪些事件开启了接收回调通知。
// 
// 默认接口请求频率限制：100次/秒。
func (c *Client) DescribeEventConfig(request *DescribeEventConfigRequest) (response *DescribeEventConfigResponse, err error) {
    if request == nil {
        request = NewDescribeEventConfigRequest()
    }
    response = NewDescribeEventConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEventsStateRequest() (request *DescribeEventsStateRequest) {
    request = &DescribeEventsStateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vod", APIVersion, "DescribeEventsState")
    return
}

func NewDescribeEventsStateResponse() (response *DescribeEventsStateResponse) {
    response = &DescribeEventsStateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// * 该接口用于业务服务器获取 [可靠回调](https://cloud.tencent.com/document/product/266/33779#.E5.8F.AF.E9.9D.A0.E5.9B.9E.E8.B0.83) 事件通知的状态。
func (c *Client) DescribeEventsState(request *DescribeEventsStateRequest) (response *DescribeEventsStateResponse, err error) {
    if request == nil {
        request = NewDescribeEventsStateRequest()
    }
    response = NewDescribeEventsStateResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeHeadTailTemplatesRequest() (request *DescribeHeadTailTemplatesRequest) {
    request = &DescribeHeadTailTemplatesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vod", APIVersion, "DescribeHeadTailTemplates")
    return
}

func NewDescribeHeadTailTemplatesResponse() (response *DescribeHeadTailTemplatesResponse) {
    response = &DescribeHeadTailTemplatesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取片头片尾模板列表。
func (c *Client) DescribeHeadTailTemplates(request *DescribeHeadTailTemplatesRequest) (response *DescribeHeadTailTemplatesResponse, err error) {
    if request == nil {
        request = NewDescribeHeadTailTemplatesRequest()
    }
    response = NewDescribeHeadTailTemplatesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeImageProcessingTemplatesRequest() (request *DescribeImageProcessingTemplatesRequest) {
    request = &DescribeImageProcessingTemplatesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vod", APIVersion, "DescribeImageProcessingTemplates")
    return
}

func NewDescribeImageProcessingTemplatesResponse() (response *DescribeImageProcessingTemplatesResponse) {
    response = &DescribeImageProcessingTemplatesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取图片处理模板列表，支持根据条件，分页查询。
func (c *Client) DescribeImageProcessingTemplates(request *DescribeImageProcessingTemplatesRequest) (response *DescribeImageProcessingTemplatesResponse, err error) {
    if request == nil {
        request = NewDescribeImageProcessingTemplatesRequest()
    }
    response = NewDescribeImageProcessingTemplatesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeImageSpriteTemplatesRequest() (request *DescribeImageSpriteTemplatesRequest) {
    request = &DescribeImageSpriteTemplatesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vod", APIVersion, "DescribeImageSpriteTemplates")
    return
}

func NewDescribeImageSpriteTemplatesResponse() (response *DescribeImageSpriteTemplatesResponse) {
    response = &DescribeImageSpriteTemplatesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询雪碧图模板，支持根据条件，分页查询。
func (c *Client) DescribeImageSpriteTemplates(request *DescribeImageSpriteTemplatesRequest) (response *DescribeImageSpriteTemplatesResponse, err error) {
    if request == nil {
        request = NewDescribeImageSpriteTemplatesRequest()
    }
    response = NewDescribeImageSpriteTemplatesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMediaInfosRequest() (request *DescribeMediaInfosRequest) {
    request = &DescribeMediaInfosRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vod", APIVersion, "DescribeMediaInfos")
    return
}

func NewDescribeMediaInfosResponse() (response *DescribeMediaInfosResponse) {
    response = &DescribeMediaInfosResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 1. 该接口可以获取多个媒体文件的多种信息，包括：
//     1. 基础信息（basicInfo）：包括媒体名称、分类、播放地址、封面图片等。
//     2. 元信息（metaData）：包括大小、时长、视频流信息、音频流信息等。
//     3. 转码结果信息（transcodeInfo）：包括该媒体转码生成的各种规格的媒体地址、视频流参数、音频流参数等。
//     4. 转动图结果信息（animatedGraphicsInfo）：对视频转动图（如 gif）后的动图信息。
//     5. 采样截图信息（sampleSnapshotInfo）：对视频采样截图后的截图信息。
//     6. 雪碧图信息（imageSpriteInfo）：对视频截取雪碧图后的雪碧图信息。
//     7. 指定时间点截图信息（snapshotByTimeOffsetInfo）：对视频依照指定时间点截图后，的截图信息。
//     8. 视频打点信息（keyFrameDescInfo）：对视频设置的打点信息。
//     9. 转自适应码流信息（adaptiveDynamicStreamingInfo）：包括规格、加密类型、打包格式等相关信息。
// 2. 可以指定回包只返回部分信息。
func (c *Client) DescribeMediaInfos(request *DescribeMediaInfosRequest) (response *DescribeMediaInfosResponse, err error) {
    if request == nil {
        request = NewDescribeMediaInfosRequest()
    }
    response = NewDescribeMediaInfosResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMediaProcessUsageDataRequest() (request *DescribeMediaProcessUsageDataRequest) {
    request = &DescribeMediaProcessUsageDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vod", APIVersion, "DescribeMediaProcessUsageData")
    return
}

func NewDescribeMediaProcessUsageDataResponse() (response *DescribeMediaProcessUsageDataResponse) {
    response = &DescribeMediaProcessUsageDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 该接口返回查询时间范围内每天使用的视频处理用量信息。
//    1. 可以查询最近365天内的视频处理统计数据。
//    2. 查询时间跨度不超过90天。
func (c *Client) DescribeMediaProcessUsageData(request *DescribeMediaProcessUsageDataRequest) (response *DescribeMediaProcessUsageDataResponse, err error) {
    if request == nil {
        request = NewDescribeMediaProcessUsageDataRequest()
    }
    response = NewDescribeMediaProcessUsageDataResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePersonSamplesRequest() (request *DescribePersonSamplesRequest) {
    request = &DescribePersonSamplesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vod", APIVersion, "DescribePersonSamples")
    return
}

func NewDescribePersonSamplesResponse() (response *DescribePersonSamplesResponse) {
    response = &DescribePersonSamplesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 该接口用于查询素材样本信息，支持根据素材 ID、名称、标签，分页查询。
func (c *Client) DescribePersonSamples(request *DescribePersonSamplesRequest) (response *DescribePersonSamplesResponse, err error) {
    if request == nil {
        request = NewDescribePersonSamplesRequest()
    }
    response = NewDescribePersonSamplesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePrepaidProductsRequest() (request *DescribePrepaidProductsRequest) {
    request = &DescribePrepaidProductsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vod", APIVersion, "DescribePrepaidProducts")
    return
}

func NewDescribePrepaidProductsResponse() (response *DescribePrepaidProductsResponse) {
    response = &DescribePrepaidProductsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 该接口可以查询用户已经购买的预付费商品的信息，包括：
//     1. 商品的类型、生效和失效日期。
//     2. 商品中每种资源的额度和剩余额度。
func (c *Client) DescribePrepaidProducts(request *DescribePrepaidProductsRequest) (response *DescribePrepaidProductsResponse, err error) {
    if request == nil {
        request = NewDescribePrepaidProductsRequest()
    }
    response = NewDescribePrepaidProductsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeProcedureTemplatesRequest() (request *DescribeProcedureTemplatesRequest) {
    request = &DescribeProcedureTemplatesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vod", APIVersion, "DescribeProcedureTemplates")
    return
}

func NewDescribeProcedureTemplatesResponse() (response *DescribeProcedureTemplatesResponse) {
    response = &DescribeProcedureTemplatesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 根据任务流模板名字，获取任务流模板详情列表。
func (c *Client) DescribeProcedureTemplates(request *DescribeProcedureTemplatesRequest) (response *DescribeProcedureTemplatesResponse, err error) {
    if request == nil {
        request = NewDescribeProcedureTemplatesRequest()
    }
    response = NewDescribeProcedureTemplatesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeReviewDetailsRequest() (request *DescribeReviewDetailsRequest) {
    request = &DescribeReviewDetailsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vod", APIVersion, "DescribeReviewDetails")
    return
}

func NewDescribeReviewDetailsResponse() (response *DescribeReviewDetailsResponse) {
    response = &DescribeReviewDetailsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// <b>本接口已不推荐使用，用 [DescribeMediaProcessUsageData](/document/product/266/41464) 替代</b>
// 
// 该接口返回查询时间范围内每天使用的视频内容智能识别时长数据，单位： 秒。
// 
// 1. 可以查询最近365天内的视频内容智能识别时长统计数据。
// 2. 查询时间跨度不超过90天。
func (c *Client) DescribeReviewDetails(request *DescribeReviewDetailsRequest) (response *DescribeReviewDetailsResponse, err error) {
    if request == nil {
        request = NewDescribeReviewDetailsRequest()
    }
    response = NewDescribeReviewDetailsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSampleSnapshotTemplatesRequest() (request *DescribeSampleSnapshotTemplatesRequest) {
    request = &DescribeSampleSnapshotTemplatesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vod", APIVersion, "DescribeSampleSnapshotTemplates")
    return
}

func NewDescribeSampleSnapshotTemplatesResponse() (response *DescribeSampleSnapshotTemplatesResponse) {
    response = &DescribeSampleSnapshotTemplatesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询采样截图模板，支持根据条件，分页查询。
func (c *Client) DescribeSampleSnapshotTemplates(request *DescribeSampleSnapshotTemplatesRequest) (response *DescribeSampleSnapshotTemplatesResponse, err error) {
    if request == nil {
        request = NewDescribeSampleSnapshotTemplatesRequest()
    }
    response = NewDescribeSampleSnapshotTemplatesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSnapshotByTimeOffsetTemplatesRequest() (request *DescribeSnapshotByTimeOffsetTemplatesRequest) {
    request = &DescribeSnapshotByTimeOffsetTemplatesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vod", APIVersion, "DescribeSnapshotByTimeOffsetTemplates")
    return
}

func NewDescribeSnapshotByTimeOffsetTemplatesResponse() (response *DescribeSnapshotByTimeOffsetTemplatesResponse) {
    response = &DescribeSnapshotByTimeOffsetTemplatesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询指定时间点截图模板，支持根据条件，分页查询。
func (c *Client) DescribeSnapshotByTimeOffsetTemplates(request *DescribeSnapshotByTimeOffsetTemplatesRequest) (response *DescribeSnapshotByTimeOffsetTemplatesResponse, err error) {
    if request == nil {
        request = NewDescribeSnapshotByTimeOffsetTemplatesRequest()
    }
    response = NewDescribeSnapshotByTimeOffsetTemplatesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeStorageDataRequest() (request *DescribeStorageDataRequest) {
    request = &DescribeStorageDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vod", APIVersion, "DescribeStorageData")
    return
}

func NewDescribeStorageDataResponse() (response *DescribeStorageDataResponse) {
    response = &DescribeStorageDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询存储空间使用情况和文件数量。
func (c *Client) DescribeStorageData(request *DescribeStorageDataRequest) (response *DescribeStorageDataResponse, err error) {
    if request == nil {
        request = NewDescribeStorageDataRequest()
    }
    response = NewDescribeStorageDataResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeStorageDetailsRequest() (request *DescribeStorageDetailsRequest) {
    request = &DescribeStorageDetailsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vod", APIVersion, "DescribeStorageDetails")
    return
}

func NewDescribeStorageDetailsResponse() (response *DescribeStorageDetailsResponse) {
    response = &DescribeStorageDetailsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 该接口返回查询时间范围内使用的点播存储空间，单位：字节。
//     1. 可以查询最近365天内的存储空间数据；
//     2. 查询时间跨度不超过90天；
//     3. 分钟粒度查询跨度不超过7天；
func (c *Client) DescribeStorageDetails(request *DescribeStorageDetailsRequest) (response *DescribeStorageDetailsResponse, err error) {
    if request == nil {
        request = NewDescribeStorageDetailsRequest()
    }
    response = NewDescribeStorageDetailsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSubAppIdsRequest() (request *DescribeSubAppIdsRequest) {
    request = &DescribeSubAppIdsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vod", APIVersion, "DescribeSubAppIds")
    return
}

func NewDescribeSubAppIdsResponse() (response *DescribeSubAppIdsResponse) {
    response = &DescribeSubAppIdsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 该接口用于获取当前账号有权限的子应用列表，包含主应用。若尚未开通子应用功能，接口将返回 
//  FailedOperation。
func (c *Client) DescribeSubAppIds(request *DescribeSubAppIdsRequest) (response *DescribeSubAppIdsResponse, err error) {
    if request == nil {
        request = NewDescribeSubAppIdsRequest()
    }
    response = NewDescribeSubAppIdsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSuperPlayerConfigsRequest() (request *DescribeSuperPlayerConfigsRequest) {
    request = &DescribeSuperPlayerConfigsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vod", APIVersion, "DescribeSuperPlayerConfigs")
    return
}

func NewDescribeSuperPlayerConfigsResponse() (response *DescribeSuperPlayerConfigsResponse) {
    response = &DescribeSuperPlayerConfigsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询超级播放器配置，支持根据条件，分页查询。
func (c *Client) DescribeSuperPlayerConfigs(request *DescribeSuperPlayerConfigsRequest) (response *DescribeSuperPlayerConfigsResponse, err error) {
    if request == nil {
        request = NewDescribeSuperPlayerConfigsRequest()
    }
    response = NewDescribeSuperPlayerConfigsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTaskDetailRequest() (request *DescribeTaskDetailRequest) {
    request = &DescribeTaskDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vod", APIVersion, "DescribeTaskDetail")
    return
}

func NewDescribeTaskDetailResponse() (response *DescribeTaskDetailResponse) {
    response = &DescribeTaskDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 通过任务 ID 查询任务的执行状态和结果的详细信息（最多可以查询3天之内提交的任务）。
func (c *Client) DescribeTaskDetail(request *DescribeTaskDetailRequest) (response *DescribeTaskDetailResponse, err error) {
    if request == nil {
        request = NewDescribeTaskDetailRequest()
    }
    response = NewDescribeTaskDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTasksRequest() (request *DescribeTasksRequest) {
    request = &DescribeTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vod", APIVersion, "DescribeTasks")
    return
}

func NewDescribeTasksResponse() (response *DescribeTasksResponse) {
    response = &DescribeTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// * 该接口用于查询任务列表；
// * 当列表数据比较多时，单次接口调用无法拉取整个列表，可通过 ScrollToken 参数，分批拉取；
// * 只能查询到最近三天（72 小时）内的任务。
func (c *Client) DescribeTasks(request *DescribeTasksRequest) (response *DescribeTasksResponse, err error) {
    if request == nil {
        request = NewDescribeTasksRequest()
    }
    response = NewDescribeTasksResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTranscodeTemplatesRequest() (request *DescribeTranscodeTemplatesRequest) {
    request = &DescribeTranscodeTemplatesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vod", APIVersion, "DescribeTranscodeTemplates")
    return
}

func NewDescribeTranscodeTemplatesResponse() (response *DescribeTranscodeTemplatesResponse) {
    response = &DescribeTranscodeTemplatesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 根据转码模板唯一标识，获取转码模板详情列表。返回结果包含符合条件的所有用户自定义模板及[系统预置转码模板](https://cloud.tencent.com/document/product/266/33476#.E9.A2.84.E7.BD.AE.E8.BD.AC.E7.A0.81.E6.A8.A1.E6.9D.BF)。
func (c *Client) DescribeTranscodeTemplates(request *DescribeTranscodeTemplatesRequest) (response *DescribeTranscodeTemplatesResponse, err error) {
    if request == nil {
        request = NewDescribeTranscodeTemplatesRequest()
    }
    response = NewDescribeTranscodeTemplatesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVodDomainsRequest() (request *DescribeVodDomainsRequest) {
    request = &DescribeVodDomainsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vod", APIVersion, "DescribeVodDomains")
    return
}

func NewDescribeVodDomainsResponse() (response *DescribeVodDomainsResponse) {
    response = &DescribeVodDomainsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 该接口用于查询点播域名信息列表。
func (c *Client) DescribeVodDomains(request *DescribeVodDomainsRequest) (response *DescribeVodDomainsResponse, err error) {
    if request == nil {
        request = NewDescribeVodDomainsRequest()
    }
    response = NewDescribeVodDomainsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeWatermarkTemplatesRequest() (request *DescribeWatermarkTemplatesRequest) {
    request = &DescribeWatermarkTemplatesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vod", APIVersion, "DescribeWatermarkTemplates")
    return
}

func NewDescribeWatermarkTemplatesResponse() (response *DescribeWatermarkTemplatesResponse) {
    response = &DescribeWatermarkTemplatesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询用户自定义水印模板，支持根据条件，分页查询。
func (c *Client) DescribeWatermarkTemplates(request *DescribeWatermarkTemplatesRequest) (response *DescribeWatermarkTemplatesResponse, err error) {
    if request == nil {
        request = NewDescribeWatermarkTemplatesRequest()
    }
    response = NewDescribeWatermarkTemplatesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeWordSamplesRequest() (request *DescribeWordSamplesRequest) {
    request = &DescribeWordSamplesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vod", APIVersion, "DescribeWordSamples")
    return
}

func NewDescribeWordSamplesResponse() (response *DescribeWordSamplesResponse) {
    response = &DescribeWordSamplesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 该接口用于根据应用场景、关键词、标签，分页查询关键词样本信息。
func (c *Client) DescribeWordSamples(request *DescribeWordSamplesRequest) (response *DescribeWordSamplesResponse, err error) {
    if request == nil {
        request = NewDescribeWordSamplesRequest()
    }
    response = NewDescribeWordSamplesResponse()
    err = c.Send(request, response)
    return
}

func NewEditMediaRequest() (request *EditMediaRequest) {
    request = &EditMediaRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vod", APIVersion, "EditMedia")
    return
}

func NewEditMediaResponse() (response *EditMediaResponse) {
    response = &EditMediaResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 对视频进行编辑（剪辑、拼接等），生成一个新的点播视频。编辑的功能包括：
// 
// 1. 对点播中的一个文件进行剪辑，生成一个新的视频；
// 2. 对点播中的多个文件进行拼接，生成一个新的视频；
// 3. 对点播中的多个文件进行剪辑，然后再拼接，生成一个新的视频；
// 4. 对点播中的一个流，直接生成一个新的视频；
// 5. 对点播中的一个流进行剪辑，生成一个新的视频；
// 6. 对点播中的多个流进行拼接，生成一个新的视频；
// 7. 对点播中的多个流进行剪辑，然后拼接，生成一个新的视频。
// 
// 对于生成的新视频，还可以指定生成后的视频是否要执行任务流。
// 
// >当对直播流做剪辑、拼接等操作时，请确保流结束后再操作。否则生成的视频可能不完整。
// 
// 如使用事件通知，事件通知的类型为 [视频编辑完成](https://cloud.tencent.com/document/product/266/33794)。
func (c *Client) EditMedia(request *EditMediaRequest) (response *EditMediaResponse, err error) {
    if request == nil {
        request = NewEditMediaRequest()
    }
    response = NewEditMediaResponse()
    err = c.Send(request, response)
    return
}

func NewExecuteFunctionRequest() (request *ExecuteFunctionRequest) {
    request = &ExecuteFunctionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vod", APIVersion, "ExecuteFunction")
    return
}

func NewExecuteFunctionResponse() (response *ExecuteFunctionResponse) {
    response = &ExecuteFunctionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口仅用于定制开发的特殊场景，除非云点播客服人员主动告知您需要使用本接口，其它情况请勿调用。
func (c *Client) ExecuteFunction(request *ExecuteFunctionRequest) (response *ExecuteFunctionResponse, err error) {
    if request == nil {
        request = NewExecuteFunctionRequest()
    }
    response = NewExecuteFunctionResponse()
    err = c.Send(request, response)
    return
}

func NewForbidMediaDistributionRequest() (request *ForbidMediaDistributionRequest) {
    request = &ForbidMediaDistributionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vod", APIVersion, "ForbidMediaDistribution")
    return
}

func NewForbidMediaDistributionResponse() (response *ForbidMediaDistributionResponse) {
    response = &ForbidMediaDistributionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// * 对媒体禁播后，除了点播控制台预览，其他场景访问视频各种资源的 URL（原始文件、转码输出文件、截图等）均会返回 403。
//   禁播/解禁操作全网生效时间约 5~10 分钟。
func (c *Client) ForbidMediaDistribution(request *ForbidMediaDistributionRequest) (response *ForbidMediaDistributionResponse, err error) {
    if request == nil {
        request = NewForbidMediaDistributionRequest()
    }
    response = NewForbidMediaDistributionResponse()
    err = c.Send(request, response)
    return
}

func NewLiveRealTimeClipRequest() (request *LiveRealTimeClipRequest) {
    request = &LiveRealTimeClipRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vod", APIVersion, "LiveRealTimeClip")
    return
}

func NewLiveRealTimeClipResponse() (response *LiveRealTimeClipResponse) {
    response = &LiveRealTimeClipResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 直播即时剪辑，是指在直播过程中（即直播尚未结束时），客户可以在过往直播内容中选择一段，实时生成一个新的视频（HLS 格式），开发者可以将其立即分享出去，或者长久保存起来。
// 
// 腾讯云点播支持两种即时剪辑模式：
// - 剪辑固化：将剪辑出来的视频保存成独立的视频，拥有独立 FileId；适用于将精彩片段**长久保存**的场景；
// - 剪辑不固化：剪辑得到的视频附属于直播录制文件，没有独立 FileId；适用于将精彩片段**临时分享**的场景。
// 
// 注意：
// - 使用直播即时剪辑功能的前提是：目标直播流开启了[时移回看](https://cloud.tencent.com/document/product/267/32742)功能。
// - 直播即时剪辑是基于直播录制生成的 m3u8 文件进行的，故而其最小剪辑精度为一个 ts 切片，无法实现秒级或者更为精确的剪辑精度。
// - 由于直播过程中可能存在断流的情况，所以有可能导致剪辑生成的实际视频时长与期望不一致。例如剪辑某个直播流的时间区间为 2018-09-20T10:30:00Z 到 2018-09-20T10:40:00Z ，如果在该时间区间中发生过断流，那么返回的媒资文件的时长将少于 10 分钟，在这种情况下，可以通过输出参数 <a href="#p_segmentset">SegmentSet</a> 感知到。
// 
// ### 剪辑固化
// 所谓剪辑固化，是指将剪辑出来的视频是保存成一个独立的视频（拥有独立的 FileId）。其生命周期不受原始直播录制视频影响（即使原始录制视频被删除，剪辑结果也不会受到任何影响）；也可以对其进行转码、微信发布等二次处理。
// 
// 举例如下：一场完整的足球比赛，直播录制出来的原始视频可能长达 2 个小时，客户出于节省成本的目的可以对这个视频存储 2 个月，但对于直播即时剪辑的「精彩时刻」视频却可以指定存储更长时间，同时可以单独对「精彩时刻」视频进行转码、微信发布等额外的点播操作，这时候可以选择直播即时剪辑并且固化的方案。
// 
// 剪辑固化的优势在于其生命周期与原始录制视频相互独立，可以独立管理、长久保存。
// 
// ### 剪辑不固化
// 所谓剪辑不固化，是指剪辑所得到的结果（m3u8 文件）与直播录制视频共享相同的 ts 分片，新生成的视频不是一个独立完整的视频（没有独立 FileId，只有播放 URL），其有效期与直播录制的完整视频有效期是一致的。一旦直播录制出来的视频被删除，也会导致该片段无法播放。
// 
// 剪辑不固化，由于其剪辑结果不是一个独立的视频，因而也不会纳入点播媒资视频管理（例如控制台的视频总数不会统计这一片段）中，也无法单独针对这个片段做转码、微信发布等任何视频处理操作。
// 
// 剪辑不固化的优势在于其剪辑操作十分“轻量化”，不会产生额外的存储开销。但其不足之处在于生命周期与原始录制视频相同，且无法进一步进行转码等视频处理。
func (c *Client) LiveRealTimeClip(request *LiveRealTimeClipRequest) (response *LiveRealTimeClipResponse, err error) {
    if request == nil {
        request = NewLiveRealTimeClipRequest()
    }
    response = NewLiveRealTimeClipResponse()
    err = c.Send(request, response)
    return
}

func NewManageTaskRequest() (request *ManageTaskRequest) {
    request = &ManageTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vod", APIVersion, "ManageTask")
    return
}

func NewManageTaskResponse() (response *ManageTaskResponse) {
    response = &ManageTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 对已发起的任务进行管理。
func (c *Client) ManageTask(request *ManageTaskRequest) (response *ManageTaskResponse, err error) {
    if request == nil {
        request = NewManageTaskRequest()
    }
    response = NewManageTaskResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAIAnalysisTemplateRequest() (request *ModifyAIAnalysisTemplateRequest) {
    request = &ModifyAIAnalysisTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vod", APIVersion, "ModifyAIAnalysisTemplate")
    return
}

func NewModifyAIAnalysisTemplateResponse() (response *ModifyAIAnalysisTemplateResponse) {
    response = &ModifyAIAnalysisTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改用户自定义视频内容分析模板。
// 
// 注意：模板 ID 10000 以下的为系统预置模板，不允许修改。
func (c *Client) ModifyAIAnalysisTemplate(request *ModifyAIAnalysisTemplateRequest) (response *ModifyAIAnalysisTemplateResponse, err error) {
    if request == nil {
        request = NewModifyAIAnalysisTemplateRequest()
    }
    response = NewModifyAIAnalysisTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAIRecognitionTemplateRequest() (request *ModifyAIRecognitionTemplateRequest) {
    request = &ModifyAIRecognitionTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vod", APIVersion, "ModifyAIRecognitionTemplate")
    return
}

func NewModifyAIRecognitionTemplateResponse() (response *ModifyAIRecognitionTemplateResponse) {
    response = &ModifyAIRecognitionTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改用户自定义视频内容识别模板。
func (c *Client) ModifyAIRecognitionTemplate(request *ModifyAIRecognitionTemplateRequest) (response *ModifyAIRecognitionTemplateResponse, err error) {
    if request == nil {
        request = NewModifyAIRecognitionTemplateRequest()
    }
    response = NewModifyAIRecognitionTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAdaptiveDynamicStreamingTemplateRequest() (request *ModifyAdaptiveDynamicStreamingTemplateRequest) {
    request = &ModifyAdaptiveDynamicStreamingTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vod", APIVersion, "ModifyAdaptiveDynamicStreamingTemplate")
    return
}

func NewModifyAdaptiveDynamicStreamingTemplateResponse() (response *ModifyAdaptiveDynamicStreamingTemplateResponse) {
    response = &ModifyAdaptiveDynamicStreamingTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改转自适应码流模板
func (c *Client) ModifyAdaptiveDynamicStreamingTemplate(request *ModifyAdaptiveDynamicStreamingTemplateRequest) (response *ModifyAdaptiveDynamicStreamingTemplateResponse, err error) {
    if request == nil {
        request = NewModifyAdaptiveDynamicStreamingTemplateRequest()
    }
    response = NewModifyAdaptiveDynamicStreamingTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAnimatedGraphicsTemplateRequest() (request *ModifyAnimatedGraphicsTemplateRequest) {
    request = &ModifyAnimatedGraphicsTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vod", APIVersion, "ModifyAnimatedGraphicsTemplate")
    return
}

func NewModifyAnimatedGraphicsTemplateResponse() (response *ModifyAnimatedGraphicsTemplateResponse) {
    response = &ModifyAnimatedGraphicsTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改用户自定义转动图模板。
func (c *Client) ModifyAnimatedGraphicsTemplate(request *ModifyAnimatedGraphicsTemplateRequest) (response *ModifyAnimatedGraphicsTemplateResponse, err error) {
    if request == nil {
        request = NewModifyAnimatedGraphicsTemplateRequest()
    }
    response = NewModifyAnimatedGraphicsTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewModifyClassRequest() (request *ModifyClassRequest) {
    request = &ModifyClassRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vod", APIVersion, "ModifyClass")
    return
}

func NewModifyClassResponse() (response *ModifyClassResponse) {
    response = &ModifyClassResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改媒体分类属性。
func (c *Client) ModifyClass(request *ModifyClassRequest) (response *ModifyClassResponse, err error) {
    if request == nil {
        request = NewModifyClassRequest()
    }
    response = NewModifyClassResponse()
    err = c.Send(request, response)
    return
}

func NewModifyContentReviewTemplateRequest() (request *ModifyContentReviewTemplateRequest) {
    request = &ModifyContentReviewTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vod", APIVersion, "ModifyContentReviewTemplate")
    return
}

func NewModifyContentReviewTemplateResponse() (response *ModifyContentReviewTemplateResponse) {
    response = &ModifyContentReviewTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改用户自定义视频内容智能识别模板。
func (c *Client) ModifyContentReviewTemplate(request *ModifyContentReviewTemplateRequest) (response *ModifyContentReviewTemplateResponse, err error) {
    if request == nil {
        request = NewModifyContentReviewTemplateRequest()
    }
    response = NewModifyContentReviewTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewModifyEventConfigRequest() (request *ModifyEventConfigRequest) {
    request = &ModifyEventConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vod", APIVersion, "ModifyEventConfig")
    return
}

func NewModifyEventConfigResponse() (response *ModifyEventConfigResponse) {
    response = &ModifyEventConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 腾讯云点播为客户提供了媒体上传、媒体管理、媒体处理等等服务，在这些服务执行过程或执行结束时，腾讯云点播也提供各种对应的事件通知，方便开发者感知服务处理状态，并做下一步的业务操作。
// 
// 开发者可以通过调用本接口来实现：
// - 设置接收回调通知的类型，目前有[ HTTP 回调通知](https://cloud.tencent.com/document/product/266/33779) 和 [基于消息队列的可靠通知](https://cloud.tencent.com/document/product/266/33779) 两种类型。
// - 对于[ HTTP 回调通知](https://cloud.tencent.com/document/product/266/33779)，可设置 3.0 格式回调的地址。3.0 格式回调的说明参见 [历史格式回调](https://cloud.tencent.com/document/product/266/33796)。
// - 对具体事件服务的通知事件选择设置接收或者忽略。
func (c *Client) ModifyEventConfig(request *ModifyEventConfigRequest) (response *ModifyEventConfigResponse, err error) {
    if request == nil {
        request = NewModifyEventConfigRequest()
    }
    response = NewModifyEventConfigResponse()
    err = c.Send(request, response)
    return
}

func NewModifyHeadTailTemplateRequest() (request *ModifyHeadTailTemplateRequest) {
    request = &ModifyHeadTailTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vod", APIVersion, "ModifyHeadTailTemplate")
    return
}

func NewModifyHeadTailTemplateResponse() (response *ModifyHeadTailTemplateResponse) {
    response = &ModifyHeadTailTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改片头片尾模板。
func (c *Client) ModifyHeadTailTemplate(request *ModifyHeadTailTemplateRequest) (response *ModifyHeadTailTemplateResponse, err error) {
    if request == nil {
        request = NewModifyHeadTailTemplateRequest()
    }
    response = NewModifyHeadTailTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewModifyImageSpriteTemplateRequest() (request *ModifyImageSpriteTemplateRequest) {
    request = &ModifyImageSpriteTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vod", APIVersion, "ModifyImageSpriteTemplate")
    return
}

func NewModifyImageSpriteTemplateResponse() (response *ModifyImageSpriteTemplateResponse) {
    response = &ModifyImageSpriteTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改用户自定义雪碧图模板。
func (c *Client) ModifyImageSpriteTemplate(request *ModifyImageSpriteTemplateRequest) (response *ModifyImageSpriteTemplateResponse, err error) {
    if request == nil {
        request = NewModifyImageSpriteTemplateRequest()
    }
    response = NewModifyImageSpriteTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewModifyMediaInfoRequest() (request *ModifyMediaInfoRequest) {
    request = &ModifyMediaInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vod", APIVersion, "ModifyMediaInfo")
    return
}

func NewModifyMediaInfoResponse() (response *ModifyMediaInfoResponse) {
    response = &ModifyMediaInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改媒体文件的属性，包括分类、名称、描述、标签、过期时间、打点信息、视频封面、字幕信息等。
func (c *Client) ModifyMediaInfo(request *ModifyMediaInfoRequest) (response *ModifyMediaInfoResponse, err error) {
    if request == nil {
        request = NewModifyMediaInfoRequest()
    }
    response = NewModifyMediaInfoResponse()
    err = c.Send(request, response)
    return
}

func NewModifyPersonSampleRequest() (request *ModifyPersonSampleRequest) {
    request = &ModifyPersonSampleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vod", APIVersion, "ModifyPersonSample")
    return
}

func NewModifyPersonSampleResponse() (response *ModifyPersonSampleResponse) {
    response = &ModifyPersonSampleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 该接口用于根据素材 ID，修改素材样本信息，包括名称、描述的修改，以及五官、标签的添加、删除、重置操作。五官删除操作需保证至少剩余 1 张图片，否则，请使用重置操作。
func (c *Client) ModifyPersonSample(request *ModifyPersonSampleRequest) (response *ModifyPersonSampleResponse, err error) {
    if request == nil {
        request = NewModifyPersonSampleRequest()
    }
    response = NewModifyPersonSampleResponse()
    err = c.Send(request, response)
    return
}

func NewModifySampleSnapshotTemplateRequest() (request *ModifySampleSnapshotTemplateRequest) {
    request = &ModifySampleSnapshotTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vod", APIVersion, "ModifySampleSnapshotTemplate")
    return
}

func NewModifySampleSnapshotTemplateResponse() (response *ModifySampleSnapshotTemplateResponse) {
    response = &ModifySampleSnapshotTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改用户自定义采样截图模板。
func (c *Client) ModifySampleSnapshotTemplate(request *ModifySampleSnapshotTemplateRequest) (response *ModifySampleSnapshotTemplateResponse, err error) {
    if request == nil {
        request = NewModifySampleSnapshotTemplateRequest()
    }
    response = NewModifySampleSnapshotTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewModifySnapshotByTimeOffsetTemplateRequest() (request *ModifySnapshotByTimeOffsetTemplateRequest) {
    request = &ModifySnapshotByTimeOffsetTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vod", APIVersion, "ModifySnapshotByTimeOffsetTemplate")
    return
}

func NewModifySnapshotByTimeOffsetTemplateResponse() (response *ModifySnapshotByTimeOffsetTemplateResponse) {
    response = &ModifySnapshotByTimeOffsetTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改用户自定义指定时间点截图模板。
func (c *Client) ModifySnapshotByTimeOffsetTemplate(request *ModifySnapshotByTimeOffsetTemplateRequest) (response *ModifySnapshotByTimeOffsetTemplateResponse, err error) {
    if request == nil {
        request = NewModifySnapshotByTimeOffsetTemplateRequest()
    }
    response = NewModifySnapshotByTimeOffsetTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewModifySubAppIdInfoRequest() (request *ModifySubAppIdInfoRequest) {
    request = &ModifySubAppIdInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vod", APIVersion, "ModifySubAppIdInfo")
    return
}

func NewModifySubAppIdInfoResponse() (response *ModifySubAppIdInfoResponse) {
    response = &ModifySubAppIdInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 该接口用于修改子应用信息，但不允许修改主应用信息。
func (c *Client) ModifySubAppIdInfo(request *ModifySubAppIdInfoRequest) (response *ModifySubAppIdInfoResponse, err error) {
    if request == nil {
        request = NewModifySubAppIdInfoRequest()
    }
    response = NewModifySubAppIdInfoResponse()
    err = c.Send(request, response)
    return
}

func NewModifySubAppIdStatusRequest() (request *ModifySubAppIdStatusRequest) {
    request = &ModifySubAppIdStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vod", APIVersion, "ModifySubAppIdStatus")
    return
}

func NewModifySubAppIdStatusResponse() (response *ModifySubAppIdStatusResponse) {
    response = &ModifySubAppIdStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 该接口用于启用、停用子应用。被停用的子应用将封停对应域名，并限制控制台访问。
func (c *Client) ModifySubAppIdStatus(request *ModifySubAppIdStatusRequest) (response *ModifySubAppIdStatusResponse, err error) {
    if request == nil {
        request = NewModifySubAppIdStatusRequest()
    }
    response = NewModifySubAppIdStatusResponse()
    err = c.Send(request, response)
    return
}

func NewModifySuperPlayerConfigRequest() (request *ModifySuperPlayerConfigRequest) {
    request = &ModifySuperPlayerConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vod", APIVersion, "ModifySuperPlayerConfig")
    return
}

func NewModifySuperPlayerConfigResponse() (response *ModifySuperPlayerConfigResponse) {
    response = &ModifySuperPlayerConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改超级播放器配置。
func (c *Client) ModifySuperPlayerConfig(request *ModifySuperPlayerConfigRequest) (response *ModifySuperPlayerConfigResponse, err error) {
    if request == nil {
        request = NewModifySuperPlayerConfigRequest()
    }
    response = NewModifySuperPlayerConfigResponse()
    err = c.Send(request, response)
    return
}

func NewModifyTranscodeTemplateRequest() (request *ModifyTranscodeTemplateRequest) {
    request = &ModifyTranscodeTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vod", APIVersion, "ModifyTranscodeTemplate")
    return
}

func NewModifyTranscodeTemplateResponse() (response *ModifyTranscodeTemplateResponse) {
    response = &ModifyTranscodeTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改用户自定义转码模板信息。
func (c *Client) ModifyTranscodeTemplate(request *ModifyTranscodeTemplateRequest) (response *ModifyTranscodeTemplateResponse, err error) {
    if request == nil {
        request = NewModifyTranscodeTemplateRequest()
    }
    response = NewModifyTranscodeTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewModifyWatermarkTemplateRequest() (request *ModifyWatermarkTemplateRequest) {
    request = &ModifyWatermarkTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vod", APIVersion, "ModifyWatermarkTemplate")
    return
}

func NewModifyWatermarkTemplateResponse() (response *ModifyWatermarkTemplateResponse) {
    response = &ModifyWatermarkTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改用户自定义水印模板，水印类型不允许修改。
func (c *Client) ModifyWatermarkTemplate(request *ModifyWatermarkTemplateRequest) (response *ModifyWatermarkTemplateResponse, err error) {
    if request == nil {
        request = NewModifyWatermarkTemplateRequest()
    }
    response = NewModifyWatermarkTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewModifyWordSampleRequest() (request *ModifyWordSampleRequest) {
    request = &ModifyWordSampleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vod", APIVersion, "ModifyWordSample")
    return
}

func NewModifyWordSampleResponse() (response *ModifyWordSampleResponse) {
    response = &ModifyWordSampleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 该接口用于修改关键词的应用场景、标签，关键词本身不可修改，如需修改，可删除重建。
func (c *Client) ModifyWordSample(request *ModifyWordSampleRequest) (response *ModifyWordSampleResponse, err error) {
    if request == nil {
        request = NewModifyWordSampleRequest()
    }
    response = NewModifyWordSampleResponse()
    err = c.Send(request, response)
    return
}

func NewParseStreamingManifestRequest() (request *ParseStreamingManifestRequest) {
    request = &ParseStreamingManifestRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vod", APIVersion, "ParseStreamingManifest")
    return
}

func NewParseStreamingManifestResponse() (response *ParseStreamingManifestResponse) {
    response = &ParseStreamingManifestResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 上传 HLS 视频时，解析索引文件内容，返回待上传的分片文件列表。分片文件路径必须是当前目录或子目录的相对路径，不能是 URL，不能是绝对路径。
func (c *Client) ParseStreamingManifest(request *ParseStreamingManifestRequest) (response *ParseStreamingManifestResponse, err error) {
    if request == nil {
        request = NewParseStreamingManifestRequest()
    }
    response = NewParseStreamingManifestResponse()
    err = c.Send(request, response)
    return
}

func NewProcessMediaRequest() (request *ProcessMediaRequest) {
    request = &ProcessMediaRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vod", APIVersion, "ProcessMedia")
    return
}

func NewProcessMediaResponse() (response *ProcessMediaResponse) {
    response = &ProcessMediaResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 对点播中的音视频媒体发起处理任务，功能包括：
// 1. 视频转码（带水印）；
// 2. 视频转动图；
// 3. 对视频按指定时间点截图；
// 4. 对视频采样截图；
// 5. 对视频截图雪碧图；
// 6. 对视频截取一张图做封面；
// 7. 对视频转自适应码流（并加密）；
// 8. 智能内容审核（鉴黄、鉴恐、鉴政）；
// 9. 智能内容分析（标签、分类、封面、按帧标签）；
// 10. 智能内容识别（视频片头片尾、人脸、文本全文、文本关键词、语音全文、语音关键词、物体）。
// 
// 如使用事件通知，事件通知的类型为 [任务流状态变更](https://cloud.tencent.com/document/product/266/9636)。
func (c *Client) ProcessMedia(request *ProcessMediaRequest) (response *ProcessMediaResponse, err error) {
    if request == nil {
        request = NewProcessMediaRequest()
    }
    response = NewProcessMediaResponse()
    err = c.Send(request, response)
    return
}

func NewProcessMediaByProcedureRequest() (request *ProcessMediaByProcedureRequest) {
    request = &ProcessMediaByProcedureRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vod", APIVersion, "ProcessMediaByProcedure")
    return
}

func NewProcessMediaByProcedureResponse() (response *ProcessMediaByProcedureResponse) {
    response = &ProcessMediaByProcedureResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 使用任务流模板，对点播中的视频发起处理任务。
// 有两种方式创建任务流模板：
// 1. 在控制台上创建和修改任务流模板；
// 2. 通过任务流模板接口创建任务流模板。
// 
// 如使用事件通知，事件通知的类型为 [任务流状态变更](https://cloud.tencent.com/document/product/266/9636)。
func (c *Client) ProcessMediaByProcedure(request *ProcessMediaByProcedureRequest) (response *ProcessMediaByProcedureResponse, err error) {
    if request == nil {
        request = NewProcessMediaByProcedureRequest()
    }
    response = NewProcessMediaByProcedureResponse()
    err = c.Send(request, response)
    return
}

func NewProcessMediaByUrlRequest() (request *ProcessMediaByUrlRequest) {
    request = &ProcessMediaByUrlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vod", APIVersion, "ProcessMediaByUrl")
    return
}

func NewProcessMediaByUrlResponse() (response *ProcessMediaByUrlResponse) {
    response = &ProcessMediaByUrlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 该 API 已经<font color='red'>不再维护</font>，请使用 MPS 产品的 [ProcessMedia](https://cloud.tencent.com/document/product/862/37578) 接口，在入参 InputInfo.UrlInputInfo.Url 中指定视频 URL。
func (c *Client) ProcessMediaByUrl(request *ProcessMediaByUrlRequest) (response *ProcessMediaByUrlResponse, err error) {
    if request == nil {
        request = NewProcessMediaByUrlRequest()
    }
    response = NewProcessMediaByUrlResponse()
    err = c.Send(request, response)
    return
}

func NewPullEventsRequest() (request *PullEventsRequest) {
    request = &PullEventsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vod", APIVersion, "PullEvents")
    return
}

func NewPullEventsResponse() (response *PullEventsResponse) {
    response = &PullEventsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// * 该接口用于业务服务器以 [可靠回调](https://cloud.tencent.com/document/product/266/33779#.E5.8F.AF.E9.9D.A0.E5.9B.9E.E8.B0.83) 的方式获取事件通知；
// * 接口为长轮询模式，即：如果服务端存在未消费事件，则立即返回给请求方；如果服务端没有未消费事件，则后台会将请求挂起，直到有新的事件产生为止；
// * 请求最多挂起5秒，建议请求方将超时时间设置为10秒；
// * 未被拉取的事件通知最多保留4天，超过该时限的事件通知可能会被清除；
// * 若该接口有事件返回，调用方必须在<font color="red">30秒</font>内调用 [确认事件通知](https://cloud.tencent.com/document/product/266/33434) 接口，确认事件通知已经处理，否则该事件通知在<font color="red">30秒</font>后会再次被拉取到。
func (c *Client) PullEvents(request *PullEventsRequest) (response *PullEventsResponse, err error) {
    if request == nil {
        request = NewPullEventsRequest()
    }
    response = NewPullEventsResponse()
    err = c.Send(request, response)
    return
}

func NewPullUploadRequest() (request *PullUploadRequest) {
    request = &PullUploadRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vod", APIVersion, "PullUpload")
    return
}

func NewPullUploadResponse() (response *PullUploadResponse) {
    response = &PullUploadResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 该接口用于将一个网络上的视频拉取到云点播平台。
func (c *Client) PullUpload(request *PullUploadRequest) (response *PullUploadResponse, err error) {
    if request == nil {
        request = NewPullUploadRequest()
    }
    response = NewPullUploadResponse()
    err = c.Send(request, response)
    return
}

func NewPushUrlCacheRequest() (request *PushUrlCacheRequest) {
    request = &PushUrlCacheRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vod", APIVersion, "PushUrlCache")
    return
}

func NewPushUrlCacheResponse() (response *PushUrlCacheResponse) {
    response = &PushUrlCacheResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 1. 预热指定的 URL 列表。
// 2. URL 的域名必须已在云点播中注册。
// 3. 单次请求最多指定20个 URL。
func (c *Client) PushUrlCache(request *PushUrlCacheRequest) (response *PushUrlCacheResponse, err error) {
    if request == nil {
        request = NewPushUrlCacheRequest()
    }
    response = NewPushUrlCacheResponse()
    err = c.Send(request, response)
    return
}

func NewResetProcedureTemplateRequest() (request *ResetProcedureTemplateRequest) {
    request = &ResetProcedureTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vod", APIVersion, "ResetProcedureTemplate")
    return
}

func NewResetProcedureTemplateResponse() (response *ResetProcedureTemplateResponse) {
    response = &ResetProcedureTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 重新设置用户自定义任务流模板的内容。  
func (c *Client) ResetProcedureTemplate(request *ResetProcedureTemplateRequest) (response *ResetProcedureTemplateResponse, err error) {
    if request == nil {
        request = NewResetProcedureTemplateRequest()
    }
    response = NewResetProcedureTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewSearchMediaRequest() (request *SearchMediaRequest) {
    request = &SearchMediaRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vod", APIVersion, "SearchMedia")
    return
}

func NewSearchMediaResponse() (response *SearchMediaResponse) {
    response = &SearchMediaResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 搜索媒体信息，支持多种条件筛选，以及支持对返回结果排序、过滤等功能，具体包括：
// - 指定文件 ID 集合 FileIds ，返回匹配集合中任意 ID 的媒体。
// - 根据多个媒体文件名 Names 或描述信息 Descriptions 进行模糊搜索。
// - 根据多个文件名前缀 NamePrefixes 进行搜索。
// - 指定分类集合 ClassIds（见输入参数），返回满足集合中任意分类的媒体。例如：媒体分类有电影、电视剧、综艺等，其中电影分类下又有子分类历史片、动作片、言情片。如果 ClassIds 指定了电影、电视剧，那么电影和电视剧下的所有子分类都会返回；而如果 ClassIds 指定的是历史片、动作片，那么只有这2个子分类下的媒体才会返回。
// - 指定标签集合 Tags（见输入参数），返回满足集合中任意标签的媒体。例如：媒体标签有二次元、宫斗、鬼畜，如果 Tags 指定了二次元、鬼畜2个标签，那么只要符合这2个标签中任意一个的媒体都会被检索出来。
// - 指定文件类型集合 Categories（见输入参数），返回满足集合中任意类型的媒体。例如：文件类型有 Video（视频）、 Audio （音频）、 Image （图片）。如果Categories指定了 Video 和 Audio 2个文件类型，那么符合这些类型的媒体都会被检索出来。
// - 指定来源集合 SourceTypes（见输入参数），返回满足集合中任意来源的媒体。例如：媒体来源有 Record (直播录制)、Upload （上传）等。如果 SourceTypes 指定了 Record 和 Upload ，那么符合这些来源的媒体都会被检索出来。
// - 指定直播推流码集合 StreamIds（见输入参数）筛选直播录制的媒体。
// - 指定视频 ID 集合 Vids （见输入参数）筛选直播录制的媒体。
// - 指定媒体的创建时间范围筛选媒体。
// - （不推荐：应使用 Names、NamePrefixes 或 Descriptions 替代）指定单个文本 Text 对媒体文件名或描述信息进行模糊搜索。
// - （不推荐：应使用 SourceTypes 替代）指定单个媒体文件来源 SourceType 进行搜索。
// - （不推荐：应使用 StreamIds 替代）指定单个推流直播码 StreamId 进行搜索。
// - （不推荐：应使用 Vids 替代）指定单个视频 ID Vid 进行搜索。
// - （不推荐：应使用 CreateTime 替代）指定单个起始创建时间 StartTime 进行搜索。
// - （不推荐：应使用 CreateTime 替代）指定单个结尾创建时间 EndTime 进行搜索。
// - 以上参数之间可以任意组合进行检索。例如：筛选创建时间在2018年12月1日12:00:00到2018年12月8日12:00:00之间、分类为电影或电视剧、带有宫斗和悬疑标签的媒体。注意，任何支持数组输入的参数，其元素之间的搜索逻辑为‘或’。所有参数之间的逻辑关系为‘与’。
// 
// - 允许对结果根据创建时间进行排序并分页返回，通过 Offset 和 Limit （见输入参数）来控制分页。
// - 允许通过 Filters 控制返回的媒体信息种类（默认返回所有信息）。可选输入包括：
//     1. 基础信息（basicInfo）：包括媒体名称、分类、播放地址、封面图片等。
//     2. 元信息（metaData）：包括大小、时长、视频流信息、音频流信息等。
//     3. 转码结果信息（transcodeInfo）：包括该媒体转码生成的各种规格的媒体地址、视频流参数、音频流参数等。
//     4. 转动图结果信息（animatedGraphicsInfo）：对视频转动图（如 gif）后的动图信息。
//     5. 采样截图信息（sampleSnapshotInfo）：对视频采样截图后的截图信息。
//     6. 雪碧图信息（imageSpriteInfo）：对视频截取雪碧图后的雪碧图信息。
//     7. 指定时间点截图信息（snapshotByTimeOffsetInfo）：对视频依照指定时间点截图后，的截图信息。
//     8. 视频打点信息（keyFrameDescInfo）：对视频设置的打点信息。
//     9. 转自适应码流信息（adaptiveDynamicStreamingInfo）：包括规格、加密类型、打包格式等相关信息。
// 
// <div id="maxResultsDesc">接口返回结果数限制：</div>
// - <b><a href="#p_offset">Offset</a> 和 <a href="#p_limit">Limit</a> 两个参数影响单次分页查询结果数。特别注意：当这2个值都缺省时，本接口最多只返回10条查询结果。</b>
// - <b>最大支持返回5000条搜索结果，超出部分不再支持查询。如果搜索结果量太大，建议使用更精细的筛选条件来减少搜索结果。</b>
func (c *Client) SearchMedia(request *SearchMediaRequest) (response *SearchMediaResponse, err error) {
    if request == nil {
        request = NewSearchMediaRequest()
    }
    response = NewSearchMediaResponse()
    err = c.Send(request, response)
    return
}

func NewSimpleHlsClipRequest() (request *SimpleHlsClipRequest) {
    request = &SimpleHlsClipRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vod", APIVersion, "SimpleHlsClip")
    return
}

func NewSimpleHlsClipResponse() (response *SimpleHlsClipResponse) {
    response = &SimpleHlsClipResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 对 HLS 视频进行按时间段裁剪，实时生成一个新的视频（HLS 格式），开发者可以将其立即分享出去，或者长久保存起来。
// 
// 腾讯云点播支持两种剪辑模式：
// - 剪辑固化：将剪辑出来的视频保存成独立的视频，拥有独立 FileId；适用于将精彩片段长久保存的场景；
// - 剪辑不固化：剪辑得到的视频附属于输入文件，没有独立 FileId；适用于将精彩片段临时分享的场景。
// 
// 注意：
// - 剪辑是基于输入 m3u8 文件进行的，故而其最小剪辑精度为一个 ts 切片，无法实现秒级或者更为精确的剪辑精度。
// 
// 
// ###剪辑固化
// 所谓剪辑固化，是指将剪辑出来的视频保存成一个独立的视频（拥有独立的 FileId）。其生命周期不受原始输入视频影响（即使原始输入视频被删除，剪辑结果也不会受到任何影响）；也可以对其进行转码、微信发布等二次处理。
// 
// 举例如下：一场完整的足球比赛，原始视频可能长达 2 个小时，客户出于节省成本的目的可以对这个视频存储 2 个月，但对于剪辑的「精彩时刻」视频却可以指定存储更长时间，同时可以单独对「精彩时刻」视频进行转码、微信发布等额外的点播操作，这时候可以选择剪辑并且固化的方案。
// 
// 剪辑固化的优势在于其生命周期与原始输入视频相互独立，可以独立管理、长久保存。
// 
// ###剪辑不固化
// 所谓剪辑不固化，是指剪辑所得到的结果（m3u8 文件）与原始输入视频共享相同的 ts 分片，新生成的视频不是一个独立完整的视频（没有独立 FileId，只有播放 URL），其有效期与原始输入的完整视频有效期是一致的。一旦原始输入的视频被删除，也会导致该片段无法播放。
// 
// 剪辑不固化，由于其剪辑结果不是一个独立的视频，因而也不会纳入点播媒资视频管理（例如控制台的视频总数不会统计这一片段）中，也无法单独针对这个片段做转码、微信发布等任何视频处理操作。
// 
// 剪辑不固化的优势在于其剪辑操作十分“轻量化”，不会产生额外的存储开销。但其不足之处在于生命周期与原始录制视频相同，且无法进一步进行转码等视频处理。
func (c *Client) SimpleHlsClip(request *SimpleHlsClipRequest) (response *SimpleHlsClipResponse, err error) {
    if request == nil {
        request = NewSimpleHlsClipRequest()
    }
    response = NewSimpleHlsClipResponse()
    err = c.Send(request, response)
    return
}

func NewSplitMediaRequest() (request *SplitMediaRequest) {
    request = &SplitMediaRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vod", APIVersion, "SplitMedia")
    return
}

func NewSplitMediaResponse() (response *SplitMediaResponse) {
    response = &SplitMediaResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 对点播视频进行拆条，生成多个新的点播视频。
func (c *Client) SplitMedia(request *SplitMediaRequest) (response *SplitMediaResponse, err error) {
    if request == nil {
        request = NewSplitMediaRequest()
    }
    response = NewSplitMediaResponse()
    err = c.Send(request, response)
    return
}

func NewWeChatMiniProgramPublishRequest() (request *WeChatMiniProgramPublishRequest) {
    request = &WeChatMiniProgramPublishRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vod", APIVersion, "WeChatMiniProgramPublish")
    return
}

func NewWeChatMiniProgramPublishResponse() (response *WeChatMiniProgramPublishResponse) {
    response = &WeChatMiniProgramPublishResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 将点播视频发布到微信小程序，供微信小程序播放器播放。
func (c *Client) WeChatMiniProgramPublish(request *WeChatMiniProgramPublishRequest) (response *WeChatMiniProgramPublishResponse, err error) {
    if request == nil {
        request = NewWeChatMiniProgramPublishRequest()
    }
    response = NewWeChatMiniProgramPublishResponse()
    err = c.Send(request, response)
    return
}
