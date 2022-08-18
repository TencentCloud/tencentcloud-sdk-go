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

package v20190321

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2019-03-21"

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


func NewCreateFileSampleRequest() (request *CreateFileSampleRequest) {
    request = &CreateFileSampleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cms", APIVersion, "CreateFileSample")
    
    
    return
}

func NewCreateFileSampleResponse() (response *CreateFileSampleResponse) {
    response = &CreateFileSampleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateFileSample
// 本文档适用于图片内容安全、视频内容安全自定义识别库的管理。
//
// <br>
//
// 通过该接口可以将图片新增到样本库。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) CreateFileSample(request *CreateFileSampleRequest) (response *CreateFileSampleResponse, err error) {
    return c.CreateFileSampleWithContext(context.Background(), request)
}

// CreateFileSample
// 本文档适用于图片内容安全、视频内容安全自定义识别库的管理。
//
// <br>
//
// 通过该接口可以将图片新增到样本库。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) CreateFileSampleWithContext(ctx context.Context, request *CreateFileSampleRequest) (response *CreateFileSampleResponse, err error) {
    if request == nil {
        request = NewCreateFileSampleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateFileSample require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateFileSampleResponse()
    err = c.Send(request, response)
    return
}

func NewCreateTextSampleRequest() (request *CreateTextSampleRequest) {
    request = &CreateTextSampleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cms", APIVersion, "CreateTextSample")
    
    
    return
}

func NewCreateTextSampleResponse() (response *CreateTextSampleResponse) {
    response = &CreateTextSampleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateTextSample
// 本文档适用于文本内容安全、音频内容安全自定义识别库的管理。
//
// <br>
//
// 通过该接口可以将文本新增到样本库。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) CreateTextSample(request *CreateTextSampleRequest) (response *CreateTextSampleResponse, err error) {
    return c.CreateTextSampleWithContext(context.Background(), request)
}

// CreateTextSample
// 本文档适用于文本内容安全、音频内容安全自定义识别库的管理。
//
// <br>
//
// 通过该接口可以将文本新增到样本库。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) CreateTextSampleWithContext(ctx context.Context, request *CreateTextSampleRequest) (response *CreateTextSampleResponse, err error) {
    if request == nil {
        request = NewCreateTextSampleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateTextSample require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateTextSampleResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteFileSampleRequest() (request *DeleteFileSampleRequest) {
    request = &DeleteFileSampleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cms", APIVersion, "DeleteFileSample")
    
    
    return
}

func NewDeleteFileSampleResponse() (response *DeleteFileSampleResponse) {
    response = &DeleteFileSampleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteFileSample
// 本文档适用于图片内容安全、视频内容安全自定义识别库的管理。
//
// <br>
//
// 删除图片样本库，支持批量删除，一次提交不超过20个。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DeleteFileSample(request *DeleteFileSampleRequest) (response *DeleteFileSampleResponse, err error) {
    return c.DeleteFileSampleWithContext(context.Background(), request)
}

// DeleteFileSample
// 本文档适用于图片内容安全、视频内容安全自定义识别库的管理。
//
// <br>
//
// 删除图片样本库，支持批量删除，一次提交不超过20个。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DeleteFileSampleWithContext(ctx context.Context, request *DeleteFileSampleRequest) (response *DeleteFileSampleResponse, err error) {
    if request == nil {
        request = NewDeleteFileSampleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteFileSample require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteFileSampleResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteTextSampleRequest() (request *DeleteTextSampleRequest) {
    request = &DeleteTextSampleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cms", APIVersion, "DeleteTextSample")
    
    
    return
}

func NewDeleteTextSampleResponse() (response *DeleteTextSampleResponse) {
    response = &DeleteTextSampleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteTextSample
// 本文档适用于文本内容安全、音频内容安全自定义识别库的管理。
//
// <br>
//
// 删除文本样本库，暂时只支持单个删除。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DeleteTextSample(request *DeleteTextSampleRequest) (response *DeleteTextSampleResponse, err error) {
    return c.DeleteTextSampleWithContext(context.Background(), request)
}

// DeleteTextSample
// 本文档适用于文本内容安全、音频内容安全自定义识别库的管理。
//
// <br>
//
// 删除文本样本库，暂时只支持单个删除。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DeleteTextSampleWithContext(ctx context.Context, request *DeleteTextSampleRequest) (response *DeleteTextSampleResponse, err error) {
    if request == nil {
        request = NewDeleteTextSampleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteTextSample require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteTextSampleResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeFileSampleRequest() (request *DescribeFileSampleRequest) {
    request = &DescribeFileSampleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cms", APIVersion, "DescribeFileSample")
    
    
    return
}

func NewDescribeFileSampleResponse() (response *DescribeFileSampleResponse) {
    response = &DescribeFileSampleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeFileSample
// 本文档适用于图片内容安全、视频内容安全自定义识别库的管理。
//
// <br>
//
// 查询图片样本库，支持批量查询。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeFileSample(request *DescribeFileSampleRequest) (response *DescribeFileSampleResponse, err error) {
    return c.DescribeFileSampleWithContext(context.Background(), request)
}

// DescribeFileSample
// 本文档适用于图片内容安全、视频内容安全自定义识别库的管理。
//
// <br>
//
// 查询图片样本库，支持批量查询。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeFileSampleWithContext(ctx context.Context, request *DescribeFileSampleRequest) (response *DescribeFileSampleResponse, err error) {
    if request == nil {
        request = NewDescribeFileSampleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeFileSample require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeFileSampleResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTextSampleRequest() (request *DescribeTextSampleRequest) {
    request = &DescribeTextSampleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cms", APIVersion, "DescribeTextSample")
    
    
    return
}

func NewDescribeTextSampleResponse() (response *DescribeTextSampleResponse) {
    response = &DescribeTextSampleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeTextSample
// 本文档适用于文本内容安全、音频内容安全自定义识别库的管理。
//
// <br>
//
// 支持批量查询文本样本库。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeTextSample(request *DescribeTextSampleRequest) (response *DescribeTextSampleResponse, err error) {
    return c.DescribeTextSampleWithContext(context.Background(), request)
}

// DescribeTextSample
// 本文档适用于文本内容安全、音频内容安全自定义识别库的管理。
//
// <br>
//
// 支持批量查询文本样本库。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeTextSampleWithContext(ctx context.Context, request *DescribeTextSampleRequest) (response *DescribeTextSampleResponse, err error) {
    if request == nil {
        request = NewDescribeTextSampleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTextSample require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTextSampleResponse()
    err = c.Send(request, response)
    return
}

func NewImageModerationRequest() (request *ImageModerationRequest) {
    request = &ImageModerationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cms", APIVersion, "ImageModeration")
    
    
    return
}

func NewImageModerationResponse() (response *ImageModerationResponse) {
    response = &ImageModerationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ImageModeration
// 图片内容检测服务（Image Moderation, IM）能自动扫描图片，识别涉黄、涉恐、涉政、涉毒等有害内容，同时支持用户配置图片黑名单，打击自定义的违规图片。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_PARAMETERERROR = "InvalidParameter.ParameterError"
//  INVALIDPARAMETERVALUE_ERRFILECONTENT = "InvalidParameterValue.ErrFileContent"
//  INVALIDPARAMETERVALUE_INVALIDIMAGECONTENT = "InvalidParameterValue.InvalidImageContent"
//  MISSINGPARAMETER_ERRFILEURL = "MissingParameter.ErrFileUrl"
//  RESOURCENOTFOUND_ERRDOWDOWNINTERNALERROR = "ResourceNotFound.ErrDowdownInternalError"
//  RESOURCENOTFOUND_ERRDOWDOWNPARAMSERROR = "ResourceNotFound.ErrDowdownParamsError"
//  RESOURCENOTFOUND_ERRDOWDOWNSOURCEERROR = "ResourceNotFound.ErrDowdownSourceError"
//  RESOURCENOTFOUND_ERRDOWDOWNTIMEOUT = "ResourceNotFound.ErrDowdownTimeOut"
//  RESOURCEUNAVAILABLE_ERRIMAGETIMEOUT = "ResourceUnavailable.ErrImageTimeOut"
//  UNAUTHORIZEDOPERATION_ERRAUTH = "UnauthorizedOperation.ErrAuth"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZED = "UnauthorizedOperation.Unauthorized"
func (c *Client) ImageModeration(request *ImageModerationRequest) (response *ImageModerationResponse, err error) {
    return c.ImageModerationWithContext(context.Background(), request)
}

// ImageModeration
// 图片内容检测服务（Image Moderation, IM）能自动扫描图片，识别涉黄、涉恐、涉政、涉毒等有害内容，同时支持用户配置图片黑名单，打击自定义的违规图片。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_PARAMETERERROR = "InvalidParameter.ParameterError"
//  INVALIDPARAMETERVALUE_ERRFILECONTENT = "InvalidParameterValue.ErrFileContent"
//  INVALIDPARAMETERVALUE_INVALIDIMAGECONTENT = "InvalidParameterValue.InvalidImageContent"
//  MISSINGPARAMETER_ERRFILEURL = "MissingParameter.ErrFileUrl"
//  RESOURCENOTFOUND_ERRDOWDOWNINTERNALERROR = "ResourceNotFound.ErrDowdownInternalError"
//  RESOURCENOTFOUND_ERRDOWDOWNPARAMSERROR = "ResourceNotFound.ErrDowdownParamsError"
//  RESOURCENOTFOUND_ERRDOWDOWNSOURCEERROR = "ResourceNotFound.ErrDowdownSourceError"
//  RESOURCENOTFOUND_ERRDOWDOWNTIMEOUT = "ResourceNotFound.ErrDowdownTimeOut"
//  RESOURCEUNAVAILABLE_ERRIMAGETIMEOUT = "ResourceUnavailable.ErrImageTimeOut"
//  UNAUTHORIZEDOPERATION_ERRAUTH = "UnauthorizedOperation.ErrAuth"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZED = "UnauthorizedOperation.Unauthorized"
func (c *Client) ImageModerationWithContext(ctx context.Context, request *ImageModerationRequest) (response *ImageModerationResponse, err error) {
    if request == nil {
        request = NewImageModerationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ImageModeration require credential")
    }

    request.SetContext(ctx)
    
    response = NewImageModerationResponse()
    err = c.Send(request, response)
    return
}

func NewManualReviewRequest() (request *ManualReviewRequest) {
    request = &ManualReviewRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cms", APIVersion, "ManualReview")
    
    
    return
}

func NewManualReviewResponse() (response *ManualReviewResponse) {
    response = &ManualReviewResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ManualReview
// 人工审核对外接口
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_CUSTOMAPPID = "InvalidParameterValue.CustomAppId"
//  INVALIDPARAMETERVALUE_DUPLICATECONTENTID = "InvalidParameterValue.DuplicateContentID"
//  INVALIDPARAMETERVALUE_ERRACTION = "InvalidParameterValue.ErrAction"
//  INVALIDPARAMETERVALUE_ERRAPPID = "InvalidParameterValue.ErrAppId"
//  INVALIDPARAMETERVALUE_ERRREQUESTID = "InvalidParameterValue.ErrRequestID"
//  INVALIDPARAMETERVALUE_ERRREQUESTSOURCE = "InvalidParameterValue.ErrRequestSource"
//  INVALIDPARAMETERVALUE_ERRUIN = "InvalidParameterValue.ErrUin"
//  INVALIDPARAMETERVALUE_INVALIDBATCHID = "InvalidParameterValue.InvalidBatchId"
//  INVALIDPARAMETERVALUE_INVALIDCONTENT = "InvalidParameterValue.InvalidContent"
//  INVALIDPARAMETERVALUE_INVALIDCONTENTID = "InvalidParameterValue.InvalidContentID"
//  INVALIDPARAMETERVALUE_INVALIDCONTENTTYPE = "InvalidParameterValue.InvalidContentType"
//  INVALIDPARAMETERVALUE_INVALIDCUSTOMAPPID = "InvalidParameterValue.InvalidCustomAppId"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETER = "InvalidParameterValue.InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDPRIORITY = "InvalidParameterValue.InvalidPriority"
//  INVALIDPARAMETERVALUE_INVALIDTITLE = "InvalidParameterValue.InvalidTitle"
//  UNAUTHORIZEDOPERATION_ERRAUTH = "UnauthorizedOperation.ErrAuth"
func (c *Client) ManualReview(request *ManualReviewRequest) (response *ManualReviewResponse, err error) {
    return c.ManualReviewWithContext(context.Background(), request)
}

// ManualReview
// 人工审核对外接口
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_CUSTOMAPPID = "InvalidParameterValue.CustomAppId"
//  INVALIDPARAMETERVALUE_DUPLICATECONTENTID = "InvalidParameterValue.DuplicateContentID"
//  INVALIDPARAMETERVALUE_ERRACTION = "InvalidParameterValue.ErrAction"
//  INVALIDPARAMETERVALUE_ERRAPPID = "InvalidParameterValue.ErrAppId"
//  INVALIDPARAMETERVALUE_ERRREQUESTID = "InvalidParameterValue.ErrRequestID"
//  INVALIDPARAMETERVALUE_ERRREQUESTSOURCE = "InvalidParameterValue.ErrRequestSource"
//  INVALIDPARAMETERVALUE_ERRUIN = "InvalidParameterValue.ErrUin"
//  INVALIDPARAMETERVALUE_INVALIDBATCHID = "InvalidParameterValue.InvalidBatchId"
//  INVALIDPARAMETERVALUE_INVALIDCONTENT = "InvalidParameterValue.InvalidContent"
//  INVALIDPARAMETERVALUE_INVALIDCONTENTID = "InvalidParameterValue.InvalidContentID"
//  INVALIDPARAMETERVALUE_INVALIDCONTENTTYPE = "InvalidParameterValue.InvalidContentType"
//  INVALIDPARAMETERVALUE_INVALIDCUSTOMAPPID = "InvalidParameterValue.InvalidCustomAppId"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETER = "InvalidParameterValue.InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDPRIORITY = "InvalidParameterValue.InvalidPriority"
//  INVALIDPARAMETERVALUE_INVALIDTITLE = "InvalidParameterValue.InvalidTitle"
//  UNAUTHORIZEDOPERATION_ERRAUTH = "UnauthorizedOperation.ErrAuth"
func (c *Client) ManualReviewWithContext(ctx context.Context, request *ManualReviewRequest) (response *ManualReviewResponse, err error) {
    if request == nil {
        request = NewManualReviewRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ManualReview require credential")
    }

    request.SetContext(ctx)
    
    response = NewManualReviewResponse()
    err = c.Send(request, response)
    return
}

func NewTextModerationRequest() (request *TextModerationRequest) {
    request = &TextModerationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cms", APIVersion, "TextModeration")
    
    
    return
}

func NewTextModerationResponse() (response *TextModerationResponse) {
    response = &TextModerationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// TextModeration
// 文本内容检测（Text Moderation）服务使用了深度学习技术，识别涉黄、涉政、涉恐等有害内容，同时支持用户配置词库，打击自定义的违规文本。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_PARAMETERERROR = "InvalidParameter.ParameterError"
//  INVALIDPARAMETERVALUE_ERRTEXTCONTENTTYPE = "InvalidParameterValue.ErrTextContentType"
//  RESOURCEUNAVAILABLE_ERRTEXTTIMEOUT = "ResourceUnavailable.ErrTextTimeOut"
//  UNAUTHORIZEDOPERATION_ERRAUTH = "UnauthorizedOperation.ErrAuth"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZED = "UnauthorizedOperation.Unauthorized"
func (c *Client) TextModeration(request *TextModerationRequest) (response *TextModerationResponse, err error) {
    return c.TextModerationWithContext(context.Background(), request)
}

// TextModeration
// 文本内容检测（Text Moderation）服务使用了深度学习技术，识别涉黄、涉政、涉恐等有害内容，同时支持用户配置词库，打击自定义的违规文本。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_PARAMETERERROR = "InvalidParameter.ParameterError"
//  INVALIDPARAMETERVALUE_ERRTEXTCONTENTTYPE = "InvalidParameterValue.ErrTextContentType"
//  RESOURCEUNAVAILABLE_ERRTEXTTIMEOUT = "ResourceUnavailable.ErrTextTimeOut"
//  UNAUTHORIZEDOPERATION_ERRAUTH = "UnauthorizedOperation.ErrAuth"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZED = "UnauthorizedOperation.Unauthorized"
func (c *Client) TextModerationWithContext(ctx context.Context, request *TextModerationRequest) (response *TextModerationResponse, err error) {
    if request == nil {
        request = NewTextModerationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("TextModeration require credential")
    }

    request.SetContext(ctx)
    
    response = NewTextModerationResponse()
    err = c.Send(request, response)
    return
}
