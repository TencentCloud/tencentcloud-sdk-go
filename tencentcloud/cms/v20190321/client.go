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


func NewCreateKeywordsSamplesRequest() (request *CreateKeywordsSamplesRequest) {
    request = &CreateKeywordsSamplesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cms", APIVersion, "CreateKeywordsSamples")
    
    
    return
}

func NewCreateKeywordsSamplesResponse() (response *CreateKeywordsSamplesResponse) {
    response = &CreateKeywordsSamplesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateKeywordsSamples
// 创建关键词接口
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateKeywordsSamples(request *CreateKeywordsSamplesRequest) (response *CreateKeywordsSamplesResponse, err error) {
    return c.CreateKeywordsSamplesWithContext(context.Background(), request)
}

// CreateKeywordsSamples
// 创建关键词接口
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateKeywordsSamplesWithContext(ctx context.Context, request *CreateKeywordsSamplesRequest) (response *CreateKeywordsSamplesResponse, err error) {
    if request == nil {
        request = NewCreateKeywordsSamplesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateKeywordsSamples require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateKeywordsSamplesResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteLibSamplesRequest() (request *DeleteLibSamplesRequest) {
    request = &DeleteLibSamplesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cms", APIVersion, "DeleteLibSamples")
    
    
    return
}

func NewDeleteLibSamplesResponse() (response *DeleteLibSamplesResponse) {
    response = &DeleteLibSamplesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteLibSamples
// 删除关键词接口
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteLibSamples(request *DeleteLibSamplesRequest) (response *DeleteLibSamplesResponse, err error) {
    return c.DeleteLibSamplesWithContext(context.Background(), request)
}

// DeleteLibSamples
// 删除关键词接口
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteLibSamplesWithContext(ctx context.Context, request *DeleteLibSamplesRequest) (response *DeleteLibSamplesResponse, err error) {
    if request == nil {
        request = NewDeleteLibSamplesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteLibSamples require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteLibSamplesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeKeywordsLibsRequest() (request *DescribeKeywordsLibsRequest) {
    request = &DescribeKeywordsLibsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cms", APIVersion, "DescribeKeywordsLibs")
    
    
    return
}

func NewDescribeKeywordsLibsResponse() (response *DescribeKeywordsLibsResponse) {
    response = &DescribeKeywordsLibsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeKeywordsLibs
// 获取用户词库列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeKeywordsLibs(request *DescribeKeywordsLibsRequest) (response *DescribeKeywordsLibsResponse, err error) {
    return c.DescribeKeywordsLibsWithContext(context.Background(), request)
}

// DescribeKeywordsLibs
// 获取用户词库列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeKeywordsLibsWithContext(ctx context.Context, request *DescribeKeywordsLibsRequest) (response *DescribeKeywordsLibsResponse, err error) {
    if request == nil {
        request = NewDescribeKeywordsLibsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeKeywordsLibs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeKeywordsLibsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLibSamplesRequest() (request *DescribeLibSamplesRequest) {
    request = &DescribeLibSamplesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cms", APIVersion, "DescribeLibSamples")
    
    
    return
}

func NewDescribeLibSamplesResponse() (response *DescribeLibSamplesResponse) {
    response = &DescribeLibSamplesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeLibSamples
// 获取关键词接口
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeLibSamples(request *DescribeLibSamplesRequest) (response *DescribeLibSamplesResponse, err error) {
    return c.DescribeLibSamplesWithContext(context.Background(), request)
}

// DescribeLibSamples
// 获取关键词接口
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeLibSamplesWithContext(ctx context.Context, request *DescribeLibSamplesRequest) (response *DescribeLibSamplesResponse, err error) {
    if request == nil {
        request = NewDescribeLibSamplesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLibSamples require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLibSamplesResponse()
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
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_IMAGEASPECTRATIOTOOLARGE = "InvalidParameter.ImageAspectRatioTooLarge"
//  INVALIDPARAMETER_IMAGEDATATOOSMALL = "InvalidParameter.ImageDataTooSmall"
//  INVALIDPARAMETER_IMAGESIZETOOSMALL = "InvalidParameter.ImageSizeTooSmall"
//  INVALIDPARAMETER_INVALIDIMAGECONTENT = "InvalidParameter.InvalidImageContent"
//  INVALIDPARAMETER_PARAMETERERROR = "InvalidParameter.ParameterError"
//  INVALIDPARAMETERVALUE_ERRFILECONTENT = "InvalidParameterValue.ErrFileContent"
//  INVALIDPARAMETERVALUE_ERRIMAGESIZE = "InvalidParameterValue.ErrImageSize"
//  INVALIDPARAMETERVALUE_INVALIDCONTENT = "InvalidParameterValue.InvalidContent"
//  INVALIDPARAMETERVALUE_INVALIDFILECONTENTSIZE = "InvalidParameterValue.InvalidFileContentSize"
//  INVALIDPARAMETERVALUE_INVALIDIMAGECONTENT = "InvalidParameterValue.InvalidImageContent"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETER = "InvalidParameterValue.InvalidParameter"
//  MISSINGPARAMETER_ERRFILEURL = "MissingParameter.ErrFileUrl"
//  RESOURCENOTFOUND_ERRDOWDOWNINTERNALERROR = "ResourceNotFound.ErrDowdownInternalError"
//  RESOURCENOTFOUND_ERRDOWDOWNPARAMSERROR = "ResourceNotFound.ErrDowdownParamsError"
//  RESOURCENOTFOUND_ERRDOWDOWNSOURCEERROR = "ResourceNotFound.ErrDowdownSourceError"
//  RESOURCENOTFOUND_ERRDOWDOWNTIMEOUT = "ResourceNotFound.ErrDowdownTimeOut"
//  RESOURCEUNAVAILABLE_ERRIMAGETIMEOUT = "ResourceUnavailable.ErrImageTimeOut"
//  RESOURCEUNAVAILABLE_IMAGEDOWNLOADERROR = "ResourceUnavailable.ImageDownloadError"
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
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_IMAGEASPECTRATIOTOOLARGE = "InvalidParameter.ImageAspectRatioTooLarge"
//  INVALIDPARAMETER_IMAGEDATATOOSMALL = "InvalidParameter.ImageDataTooSmall"
//  INVALIDPARAMETER_IMAGESIZETOOSMALL = "InvalidParameter.ImageSizeTooSmall"
//  INVALIDPARAMETER_INVALIDIMAGECONTENT = "InvalidParameter.InvalidImageContent"
//  INVALIDPARAMETER_PARAMETERERROR = "InvalidParameter.ParameterError"
//  INVALIDPARAMETERVALUE_ERRFILECONTENT = "InvalidParameterValue.ErrFileContent"
//  INVALIDPARAMETERVALUE_ERRIMAGESIZE = "InvalidParameterValue.ErrImageSize"
//  INVALIDPARAMETERVALUE_INVALIDCONTENT = "InvalidParameterValue.InvalidContent"
//  INVALIDPARAMETERVALUE_INVALIDFILECONTENTSIZE = "InvalidParameterValue.InvalidFileContentSize"
//  INVALIDPARAMETERVALUE_INVALIDIMAGECONTENT = "InvalidParameterValue.InvalidImageContent"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETER = "InvalidParameterValue.InvalidParameter"
//  MISSINGPARAMETER_ERRFILEURL = "MissingParameter.ErrFileUrl"
//  RESOURCENOTFOUND_ERRDOWDOWNINTERNALERROR = "ResourceNotFound.ErrDowdownInternalError"
//  RESOURCENOTFOUND_ERRDOWDOWNPARAMSERROR = "ResourceNotFound.ErrDowdownParamsError"
//  RESOURCENOTFOUND_ERRDOWDOWNSOURCEERROR = "ResourceNotFound.ErrDowdownSourceError"
//  RESOURCENOTFOUND_ERRDOWDOWNTIMEOUT = "ResourceNotFound.ErrDowdownTimeOut"
//  RESOURCEUNAVAILABLE_ERRIMAGETIMEOUT = "ResourceUnavailable.ErrImageTimeOut"
//  RESOURCEUNAVAILABLE_IMAGEDOWNLOADERROR = "ResourceUnavailable.ImageDownloadError"
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
//  INVALIDPARAMETERVALUE_ERRFILECONTENT = "InvalidParameterValue.ErrFileContent"
//  INVALIDPARAMETERVALUE_ERRTEXTCONTENTLEN = "InvalidParameterValue.ErrTextContentLen"
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
//  INVALIDPARAMETERVALUE_ERRFILECONTENT = "InvalidParameterValue.ErrFileContent"
//  INVALIDPARAMETERVALUE_ERRTEXTCONTENTLEN = "InvalidParameterValue.ErrTextContentLen"
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
